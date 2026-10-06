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
}

// shapeModeKnown classifies a `mode` value against the closed set. allow-statements is the one
// implemented mode; exact follows with its own implementation step and is
// rejected until then, loudly instead of as a silent no-op.
func shapeModeKnown(mode string) (implemented, known bool) {
	switch mode {
	case "allow-statements":
		return true, true
	case "exact":
		return false, true
	default:
		return false, false
	}
}

// NewShape validates one shapes entry fail-closed (exit 2, SPEC-CONF-001): the
// keys the config adapter could decode are checked here for VALUES and
// combinations. Dialect membership and the parseability of literal entries are
// the extraction's part (its Validate), because only it knows the dialects.
func NewShape(files []string, dialect, mode string, allow []ShapeAllowSpec, unused, expect string) (Shape, error) {
	if err := validateShapeFiles(files); err != nil {
		return Shape{}, err
	}
	if dialect == "" {
		return Shape{}, fmt.Errorf("shapes: dialect fehlt")
	}
	implemented, known := shapeModeKnown(mode)
	if !known {
		return Shape{}, fmt.Errorf("shapes: ungültiger mode %q (allow-statements|exact)", mode)
	}
	if !implemented {
		return Shape{}, fmt.Errorf("shapes: mode %q ist in dieser Version noch nicht implementiert", mode)
	}
	if expect != "" {
		return Shape{}, fmt.Errorf("shapes: expect gehört zu mode exact, nicht zu %q", mode)
	}
	if unused != "" {
		return Shape{}, fmt.Errorf("shapes: unused ist in dieser Version noch nicht implementiert")
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
	return Shape{Files: files, Dialect: dialect, Mode: mode, Allow: out}, nil
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
// order of SPEC-DET-001. It stands outside the per-import first-match chain:
// shape-unlisted judges statements of named files, independent of layer,
// language and composition root. Byte-identical findings — the same file matched
// by two entries — are reported once (SPEC-CONF-001).
func EvaluateShapes(m Model, s ShapeScan) []Finding {
	var fs []Finding
	for _, f := range s.Files {
		sh := m.Shapes[f.Entry]
		for _, st := range f.Statements {
			if !allowed(sh, s.Literals[f.Entry], st.Text) {
				fs = append(fs, Finding{Path: f.Path, Line: st.Line, Rule: "shape-unlisted", Msg: st.Text})
			}
		}
	}
	sortFindings(fs)
	return dedupeSorted(fs)
}

// allowed reports whether a normalized statement equals a literal entry or fully
// matches a regex entry. Set semantics: order and repetition do not matter.
func allowed(sh Shape, literals []string, text string) bool {
	for i, a := range sh.Allow {
		if a.Regex {
			if a.match(text) {
				return true
			}
			continue
		}
		if literals[i] == text {
			return true
		}
	}
	return false
}

// SortFindings puts findings into the total order of SPEC-DET-001. The CLI uses
// it after merging the shapes findings into the import-rule findings; it
// collapses nothing — two identical constructs entries legitimately report twice.
func SortFindings(fs []Finding) { sortFindings(fs) }

// dedupeSorted collapses byte-identical neighbours of a SORTED list — only for
// the shapes findings, where one file matched by two entries yields the same
// line twice (SPEC-CONF-001).
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
