package core

import (
	"fmt"
	"path"
	"regexp"
	"strings"
)

// ShapeAllowSpec is one `allow` entry as decoded from the config, before
// validation (AC-FA-RULE-012, SPEC-CONF-001). Match is "" or "literal" for a
// literal entry, "regex" for an RE2 entry. Line is the entry's line in
// `.a-check.yml` — the location a `shape-unused` finding points at.
type ShapeAllowSpec struct {
	Pattern string
	Match   string
	Line    int
}

// ShapeAllow is a validated `allow` entry. A literal entry is compared after the
// extraction has normalized it like the checked file (ADR-0041 point 5); a regex
// entry is matched FULLY ANCHORED against the whole normalized statement — the
// pattern compiled on its own AND wrapped as ^(?:…)$ (SPEC-CONF-001), so neither
// an unanchored `dependencies\{.*` nor an only-wrapped-compiling `a)|(.*` gets
// through.
type ShapeAllow struct {
	Pattern string
	Regex   bool
	Line    int
	match   func(string) bool // compiled, anchored; nil for a literal entry
}

// Shape is one validated `shapes` entry (AC-FA-RULE-012, ADR-0041): the files
// whose statements must all be listed in Allow. Dialect is validated against the
// extraction's dialect registry, not here — the core knows no dialect, the same
// split as `languages` (SPEC-EXTRACT-001).
type Shape struct {
	Files   []string
	Dialect string
	Mode    string
	Allow   []ShapeAllow
	Unused  bool
	Expect  string
}

// Statement is one top-level statement of a shapes file in its normalized form,
// with the original line it starts on (SPEC-EXTRACT-001).
type Statement struct {
	Text string
	Line int
}

// ShapeFile is the statement sequence of one file matched by shapes entry Entry
// (the INDEX into Model.Shapes, like ConstructHit — two entries naming the same
// file stay distinguishable). Lines is the file's line count, at least 1.
type ShapeFile struct {
	Entry      int
	Path       string
	Statements []Statement
	Lines      int
}

// ShapeScan is what the extraction delivers for the shapes block: the statement
// sequences of all matched files and, per entry, the normalized form of each
// LITERAL allow entry (Literals[e][i]; "" for a regex entry). The core reads
// nothing itself (ADR-0041 point 9).
type ShapeScan struct {
	Files    []ShapeFile
	Literals [][]string
	// Expected[e] is the normalized statement sequence of entry e's `expect`
	// file (mode exact); nil for an allow-statements entry.
	Expected [][]Statement
}

// NewShape validates one shapes entry fail-closed (exit 2, SPEC-CONF-001): the
// keys the config adapter could decode are checked here for VALUES and
// combinations. Dialect membership and the parseability of literal entries are
// the extraction's part (its Validate), because only it knows the dialects.
func NewShape(files []string, dialect, mode string, allow []ShapeAllowSpec, unused, expect string) (Shape, error) {
	files = trimDotSlash(files)
	if err := validateShapeFiles(files); err != nil {
		return Shape{}, err
	}
	if dialect == "" {
		return Shape{}, fmt.Errorf("shapes: dialect fehlt")
	}
	switch mode {
	case "allow-statements":
		return newAllowShape(files, dialect, allow, unused, expect)
	case "exact":
		return newExactShape(files, dialect, allow, unused, expect)
	default:
		return Shape{}, fmt.Errorf("shapes: ungültiger mode %q (allow-statements|exact)", mode)
	}
}

// newAllowShape: allow is mandatory and non-empty, unused is "" or "fail",
// expect belongs to exact.
func newAllowShape(files []string, dialect string, allow []ShapeAllowSpec, unused, expect string) (Shape, error) {
	if expect != "" {
		return Shape{}, fmt.Errorf("shapes: expect gehört zu mode exact, nicht zu allow-statements")
	}
	if unused != "" && unused != "fail" {
		return Shape{}, fmt.Errorf("shapes: ungültiger unused-Wert %q (fail)", unused)
	}
	if len(allow) == 0 {
		return Shape{}, fmt.Errorf("shapes: allow fehlt oder ist leer (mode allow-statements)")
	}
	out := make([]ShapeAllow, 0, len(allow))
	for _, a := range allow {
		sa, err := newShapeAllow(a)
		if err != nil {
			return Shape{}, err
		}
		out = append(out, sa)
	}
	return Shape{Files: files, Dialect: dialect, Mode: "allow-statements", Allow: out, Unused: unused == "fail"}, nil
}

// newExactShape: expect is mandatory and inside the scan root; allow and unused
// belong to allow-statements.
func newExactShape(files []string, dialect string, allow []ShapeAllowSpec, unused, expect string) (Shape, error) {
	if len(allow) > 0 {
		return Shape{}, fmt.Errorf("shapes: allow gehört zu mode allow-statements, nicht zu exact")
	}
	if unused != "" {
		return Shape{}, fmt.Errorf("shapes: unused gehört zu mode allow-statements, nicht zu exact")
	}
	if expect == "" {
		return Shape{}, fmt.Errorf("shapes: expect fehlt (mode exact)")
	}
	if err := InsideRoot("expect", expect); err != nil {
		return Shape{}, err
	}
	// lexically cleaned, so `d//b.kts`, `x/../d/b.kts` and `./d/b.kts` name the
	// same file as the scan paths do (Review slice-211 N-2)
	expect = path.Clean(expect)
	return Shape{Files: files, Dialect: dialect, Mode: "exact", Expect: expect}, nil
}

// validateShapeFiles rejects an empty file list, an empty glob and every glob
// that points outside the scan root — absolute, or starting with `..` after
// lexical cleaning (hermeticity, AC-QA-02).
func validateShapeFiles(files []string) error {
	if len(files) == 0 {
		return fmt.Errorf("shapes: files fehlt oder ist leer")
	}
	for _, g := range files {
		if err := InsideRoot("files", g); err != nil {
			return err
		}
	}
	return nil
}

// trimDotSlash removes a leading `./` from each glob: scan paths are relative
// without it, so `./a` would never match and end in a misleading "trifft keine
// Datei" (SPEC-CONF-001).
func trimDotSlash(globs []string) []string {
	out := make([]string, len(globs))
	for i, g := range globs {
		for strings.HasPrefix(g, "./") {
			g = g[2:]
		}
		out[i] = g
	}
	return out
}

// InsideRoot rejects an empty path/glob and one that leaves the scan root:
// absolute, or `..`-leading after lexical cleaning (SPEC-CONF-001). Exported for
// the config adapter, which checks `expect` with the same rule.
func InsideRoot(key, p string) error {
	if p == "" {
		return fmt.Errorf("shapes: leerer %s-Eintrag unzulässig", key)
	}
	c := path.Clean(p)
	if strings.HasPrefix(p, "/") || c == ".." || strings.HasPrefix(c, "../") {
		return fmt.Errorf("shapes: %s %q zeigt aus der Scan-Wurzel hinaus", key, p)
	}
	return nil
}

// newShapeAllow validates one allow entry: a non-empty pattern; match literal or
// regex; a regex must compile on its own AND wrapped (SPEC-CONF-001).
func newShapeAllow(a ShapeAllowSpec) (ShapeAllow, error) {
	if a.Pattern == "" {
		return ShapeAllow{}, fmt.Errorf("shapes: leeres allow-pattern unzulässig")
	}
	switch a.Match {
	case "", "literal":
		return ShapeAllow{Pattern: a.Pattern, Line: a.Line}, nil
	case "regex":
		if _, err := regexp.Compile(a.Pattern); err != nil {
			return ShapeAllow{}, fmt.Errorf("shapes: allow %q: ungültige Regex: %w", a.Pattern, err)
		}
		re, err := regexp.Compile("^(?:" + a.Pattern + ")$")
		if err != nil {
			return ShapeAllow{}, fmt.Errorf("shapes: allow %q: Regex kompiliert verankert nicht: %w", a.Pattern, err)
		}
		return ShapeAllow{Pattern: a.Pattern, Regex: true, Line: a.Line, match: re.MatchString}, nil
	default:
		return ShapeAllow{}, fmt.Errorf("shapes: allow %q: ungültiges match %q (literal|regex)", a.Pattern, a.Match)
	}
}

// EvaluateShapes applies the shapes rules (SPEC-RULE-001) to the extraction's
// statement sequences and returns the findings, deduplicated and in the total
// order of SPEC-DET-001. It stands outside the per-import first-match chain: it
// judges statements of named files, independent of layer, language and
// composition root. Byte-identical findings — the same file matched by two
// entries, or two equal statements on one line — are reported once
// (SPEC-CONF-001). configPath locates shape-unused, which points at the entry
// in the configuration, not at a checked file.
func EvaluateShapes(m Model, s ShapeScan, configPath string) []Finding {
	var fs []Finding
	hit := make([][]bool, len(m.Shapes))
	for e, sh := range m.Shapes {
		hit[e] = make([]bool, len(sh.Allow))
	}
	for _, f := range s.Files {
		sh := m.Shapes[f.Entry]
		if sh.Mode == "exact" {
			if d, ok := firstDifference(f, s.Expected[f.Entry]); ok {
				fs = append(fs, d)
			}
			continue
		}
		for _, st := range f.Statements {
			if !markAllowed(sh, s.Literals[f.Entry], st.Text, hit[f.Entry]) {
				fs = append(fs, Finding{Path: f.Path, Line: st.Line, Rule: "shape-unlisted", Msg: oneLine(st.Text)})
			}
		}
	}
	fs = append(fs, unusedFindings(m, hit, configPath)...)
	sortFindings(fs)
	return dedupeSorted(fs)
}

// markAllowed reports whether a normalized statement equals a literal entry or
// fully matches a regex entry, and marks EVERY entry it matches as hit — so a
// second entry that also matches is not reported as unused. Set semantics:
// order and repetition do not matter.
func markAllowed(sh Shape, literals []string, text string, hit []bool) bool {
	ok := false
	for i, a := range sh.Allow {
		if (a.Regex && a.match(text)) || (!a.Regex && literals[i] == text) {
			hit[i] = true
			ok = true
		}
	}
	return ok
}

// firstDifference compares a file's statements with the expected sequence of
// its exact entry and reports the FIRST position where they differ, or ok=false
// if they are equal (SPEC-RULE-001 shape-differs).
func firstDifference(f ShapeFile, want []Statement) (Finding, bool) {
	got := f.Statements
	for i := 0; i < len(got) || i < len(want); i++ {
		switch {
		case i >= len(got):
			return Finding{Path: f.Path, Line: f.Lines, Rule: "shape-differs", Msg: "fehlt: " + oneLine(want[i].Text)}, true
		case i >= len(want):
			return Finding{Path: f.Path, Line: got[i].Line, Rule: "shape-differs", Msg: oneLine(got[i].Text) + " (nicht in der Sollform)"}, true
		case got[i].Text != want[i].Text:
			return Finding{Path: f.Path, Line: got[i].Line, Rule: "shape-differs", Msg: oneLine(got[i].Text) + " (erwartet: " + oneLine(want[i].Text) + ")"}, true
		}
	}
	return Finding{}, false
}

// unusedFindings reports every allow entry of an `unused: fail` entry that
// matched in none of its files (Opt-in, no warn level — ADR-0041 point 7). The
// message is `literal: ` or `regex: ` plus the entry in its declared form,
// trailing line ends dropped, written by oneLine (SPEC-RULE-001). The prefix
// keeps the kind unambiguous for every literal text (ADR-0042 point 5).
func unusedFindings(m Model, hit [][]bool, configPath string) []Finding {
	var fs []Finding
	for e, sh := range m.Shapes {
		if !sh.Unused {
			continue
		}
		for i, a := range sh.Allow {
			if hit[e][i] {
				continue
			}
			kind := "literal: "
			if a.Regex {
				kind = "regex: "
			}
			msg := kind + oneLine(strings.TrimRight(a.Pattern, "\r\n"))
			fs = append(fs, Finding{Path: configPath, Line: a.Line, Rule: "shape-unused", Msg: msg})
		}
	}
	return fs
}

// SortFindings puts findings into the total order of SPEC-DET-001. The CLI uses
// it after merging the shapes findings into the import-rule findings; it
// collapses nothing — two identical constructs entries legitimately report twice.
func SortFindings(fs []Finding) { sortFindings(fs) }

// dedupeSorted collapses byte-identical neighbours of a SORTED list — only for
// the shapes findings, where two entries for one file or two equal statements on
// one line yield the same line twice; the second line carries nothing new
// (SPEC-CONF-001).
func dedupeSorted(fs []Finding) []Finding {
	out := fs[:0]
	for i, f := range fs {
		if i > 0 && f == fs[i-1] {
			continue
		}
		out = append(out, f)
	}
	return out
}

// oneLine writes a statement or entry for a finding message on ONE line, and
// reversibly: backslash as `\\`, LF as `\n`, CR as `\r` (SPEC-RULE-001). The
// mapping is injective — two different texts never share a message —, which the
// dedupe of byte-identical findings relies on. Only the output is escaped;
// comparisons use the raw text.
func oneLine(s string) string {
	return strings.NewReplacer(`\`, `\\`, "\n", `\n`, "\r", `\r`).Replace(s)
}
