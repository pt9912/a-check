package extract

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"

	"github.com/pt9912/a-check/internal/hexagon/core"
)

// dialectFn normalizes one source of a shapes dialect into its top-level
// statements and line count (SPEC-EXTRACT-001), or fails if it cannot be split.
type dialectFn func(src string) ([]core.Statement, int, error)

// dialects is the registry of shapes dialects — the single source of the
// supported set, like backends for languages. A new dialect is one entry.
func dialects() map[string]dialectFn {
	return map[string]dialectFn{"kotlin": normalizeKotlin}
}

// checkShapes validates the shapes block without reading a file: every dialect
// is registered, and every literal allow entry normalizes to exactly ONE
// statement (SPEC-CONF-001). It runs in Validate, so the no-scan --print-graph
// path fails on a broken entry exactly as a scan does.
func (a Adapter) checkShapes(m core.Model) error {
	_, err := a.shapeLiterals(m)
	return err
}

// shapeLiterals normalizes the literal allow entries, per shapes entry in
// declaration order; a regex entry keeps "" (it is matched, not compared).
func (a Adapter) shapeLiterals(m core.Model) ([][]string, error) {
	out := make([][]string, len(m.Shapes))
	for e, sh := range m.Shapes {
		norm, ok := a.dialects[sh.Dialect]
		if !ok {
			return nil, fmt.Errorf("shapes[%d]: unbekannter dialect %q (%s)", e, sh.Dialect, a.dialectList())
		}
		out[e] = make([]string, len(sh.Allow))
		for i, al := range sh.Allow {
			if al.Regex {
				continue
			}
			stmts, _, err := norm(al.Pattern)
			if err != nil {
				return nil, fmt.Errorf("shapes[%d]: allow %q lässt sich nicht zerlegen: %w", e, al.Pattern, err)
			}
			if len(stmts) != 1 {
				return nil, fmt.Errorf("shapes[%d]: allow %q ergibt %d Anweisungen, erwartet genau eine", e, al.Pattern, len(stmts))
			}
			out[e][i] = stmts[0].Text
		}
	}
	return out, nil
}

func (a Adapter) dialectList() string {
	names := make([]string, 0, len(a.dialects))
	for n := range a.dialects {
		names = append(names, n)
	}
	sort.Strings(names)
	return strings.Join(names, "|")
}

// Shapes reads every file a shapes entry names and returns their statement
// sequences plus the normalized literal entries (AC-FA-RULE-012). It is
// independent of the languages/layers walk. Fail-closed (exit 2): a glob that
// matches no file, a matched file that exclude removes, a file that cannot be
// split, and the exact-mode cases of the Sollform file — checked per entry in
// declaration order (SPEC-DET-001).
func (a Adapter) Shapes(root string, m core.Model) (core.ShapeScan, error) {
	lits, err := a.shapeLiterals(m)
	if err != nil {
		return core.ShapeScan{}, err
	}
	scan := core.ShapeScan{Literals: lits, Expected: make([][]core.Statement, len(m.Shapes))}
	for e, sh := range m.Shapes {
		if sh.Mode == "exact" {
			want, err := a.expected(e, root, sh)
			if err != nil {
				return core.ShapeScan{}, err
			}
			scan.Expected[e] = want
		}
		files, err := a.entryFiles(e, root, sh, m.Exclude)
		if err != nil {
			return core.ShapeScan{}, err
		}
		scan.Files = append(scan.Files, files...)
	}
	return scan, nil
}

// regularNoSymlink checks every component of rel below root with Lstat: no
// component may be a symlink — a link anywhere in the path could lead out of
// the scan root (Review slice-211 N-1) — and the last one must be a regular
// file.
func regularNoSymlink(root, rel string) error {
	cur := root
	segs := strings.Split(rel, "/")
	for i, s := range segs {
		cur = filepath.Join(cur, s)
		info, err := os.Lstat(cur)
		if err != nil {
			return fmt.Errorf("fehlt oder ist nicht lesbar: %w", err)
		}
		if info.Mode()&os.ModeSymlink != 0 {
			return fmt.Errorf("der Pfad enthält einen Symlink (%s) — er wird nicht verfolgt", path.Join(segs[:i+1]...))
		}
		if i == len(segs)-1 && !info.Mode().IsRegular() {
			return fmt.Errorf("ist keine reguläre Datei")
		}
	}
	return nil
}

// entryFiles reads and normalizes the files of one entry. In exact mode the
// Sollform file must not be one of them — the entry could never report.
func (a Adapter) entryFiles(e int, root string, sh core.Shape, exclude []string) ([]core.ShapeFile, error) {
	paths, err := shapePaths(e, root, sh.Files, exclude)
	if err != nil {
		return nil, err
	}
	out := make([]core.ShapeFile, 0, len(paths))
	for _, rel := range paths {
		if sh.Mode == "exact" && sameFile(root, rel, sh.Expect) {
			return nil, fmt.Errorf("shapes[%d]: expect-Datei %s ist zugleich die geprüfte Datei %s — der Eintrag könnte nie melden", e, sh.Expect, rel)
		}
		data, rerr := os.ReadFile(filepath.Join(root, filepath.FromSlash(rel)))
		if rerr != nil {
			return nil, rerr
		}
		stmts, lines, nerr := a.dialects[sh.Dialect](string(data))
		if nerr != nil {
			return nil, fmt.Errorf("shapes[%d]: %s lässt sich nicht zerlegen: %w", e, rel, nerr)
		}
		out = append(out, core.ShapeFile{Entry: e, Path: rel, Statements: stmts, Lines: lines})
	}
	return out, nil
}

// expected reads and normalizes the Sollform file of an exact entry. A missing
// file, one that is no regular file and an unsplittable one are exit 2
// (SPEC-CONF-001).
func (a Adapter) expected(e int, root string, sh core.Shape) ([]core.Statement, error) {
	p := filepath.Join(root, filepath.FromSlash(sh.Expect))
	if err := regularNoSymlink(root, sh.Expect); err != nil {
		return nil, fmt.Errorf("shapes[%d]: expect-Datei %q: %w", e, sh.Expect, err)
	}
	data, err := os.ReadFile(p)
	if err != nil {
		return nil, fmt.Errorf("shapes[%d]: expect-Datei %q fehlt oder ist nicht lesbar: %w", e, sh.Expect, err)
	}
	want, _, nerr := a.dialects[sh.Dialect](string(data))
	if nerr != nil {
		return nil, fmt.Errorf("shapes[%d]: expect-Datei %s lässt sich nicht zerlegen: %w", e, sh.Expect, nerr)
	}
	return want, nil
}

// sameFile reports whether two paths below root name the same file — by
// name, and by identity (os.SameFile), so a hard link of the checked file as
// Sollform is caught too (Review slice-211 N-5).
func sameFile(root, a, b string) bool {
	if a == b {
		return true
	}
	ia, ea := os.Stat(filepath.Join(root, filepath.FromSlash(a)))
	ib, eb := os.Stat(filepath.Join(root, filepath.FromSlash(b)))
	return ea == nil && eb == nil && os.SameFile(ia, ib)
}

// shapePaths resolves one entry's globs to the sorted, deduplicated set of
// matching files. Each glob must match at least one file, and no match may lie
// under exclude — a file named in shapes and removed by exclude is a
// contradiction of the configuration, not a silent skip.
func shapePaths(entry int, root string, globs, exclude []string) ([]string, error) {
	seen := map[string]bool{}
	var out []string
	for _, g := range globs {
		matches, err := globFiles(root, g)
		if err != nil {
			return nil, err
		}
		if len(matches) == 0 {
			return nil, fmt.Errorf("shapes[%d]: files-Glob %q trifft keine Datei", entry, g)
		}
		for _, rel := range matches {
			if core.MatchGlobs(rel, exclude) {
				return nil, fmt.Errorf("shapes[%d]: %s ist in shapes genannt und durch exclude ausgenommen (Widerspruch)", entry, rel)
			}
			if !seen[rel] {
				seen[rel] = true
				out = append(out, rel)
			}
		}
	}
	sort.Strings(out)
	return out, nil
}

// globFiles lists the regular files matching glob g, walking only the glob's
// literal directory prefix (the segments before the first wildcard). It does not
// prune by exclude: an excluded match must be SEEN to be reported as a
// contradiction. .git is skipped.
func globFiles(root, g string) ([]string, error) {
	base := literalPrefixDir(g)
	start := filepath.Join(root, filepath.FromSlash(base))
	info, err := os.Stat(start)
	if err != nil || !info.IsDir() {
		// no such directory: the glob matches nothing, which the caller
		// reports as exit 2 — not an I/O error of its own
		return nil, nil
	}
	var out []string
	werr := filepath.WalkDir(start, func(p string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if d.IsDir() {
			if d.Name() == ".git" {
				return filepath.SkipDir
			}
			return nil
		}
		rel, rerr := filepath.Rel(root, p)
		if rerr != nil {
			return rerr
		}
		rel = filepath.ToSlash(rel)
		if d.Type().IsRegular() && core.MatchGlobs(rel, []string{g}) {
			out = append(out, rel)
		}
		return nil
	})
	return out, werr
}

// literalPrefixDir returns the directory part of g before its first wildcard
// segment ("." when the first segment already holds one).
func literalPrefixDir(g string) string {
	segs := strings.Split(g, "/")
	var lit []string
	for _, s := range segs[:len(segs)-1] {
		if strings.ContainsAny(s, "*?") {
			break
		}
		lit = append(lit, s)
	}
	if len(lit) == 0 {
		return "."
	}
	return path.Join(lit...)
}
