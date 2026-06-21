// Package extract is the extraction adapter (ARC-003): it walks the source tree
// and yields per-file imports text-heuristically (SPEC-EXTRACT-001), plus
// forbidden-construct hits for port-impurity. It is a heuristic, not a parser:
// import-like lines in comments are stripped; the boundary is documented
// (AC-QA-02), and `markers.ignore_symbols` provides an allowlist.
package extract

import (
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/pt9912/a-check/internal/core"
)

// Adapter implements core.ExtractionPort. Its compiled patterns live on the
// value (not as package globals) to satisfy the lint profile (ADR-0005).
type Adapter struct {
	goSingle, goBlock, goQuoted     *regexp.Regexp
	cppInclude                      *regexp.Regexp
	rustUse, rustCrate              *regexp.Regexp
	kotlinImp                       *regexp.Regexp
}

// New returns an extraction adapter.
func New() core.ExtractionPort { return newAdapter() }

func newAdapter() Adapter {
	return Adapter{
		goSingle:   regexp.MustCompile(`^\s*import\s+(?:[\w.]+\s+)?"([^"]+)"`),
		goBlock:    regexp.MustCompile(`^\s*import\s*\(\s*$`),
		goQuoted:   regexp.MustCompile(`^\s*(?:[\w.]+\s+)?"([^"]+)"`),
		cppInclude: regexp.MustCompile(`^\s*#\s*include\s*[<"]([^>"]+)[>"]`),
		rustUse:    regexp.MustCompile(`^\s*use\s+([A-Za-z_][A-Za-z0-9_]*)`),
		rustCrate:  regexp.MustCompile(`^\s*extern\s+crate\s+([A-Za-z_][A-Za-z0-9_]*)`),
		kotlinImp:  regexp.MustCompile(`^\s*import\s+([A-Za-z_][A-Za-z0-9_.]*)`),
	}
}

// Extract walks root and returns the imports per source file, stably ordered.
func (a Adapter) Extract(root string, m core.Model) ([]core.FileImports, error) {
	var out []core.FileImports
	err := filepath.WalkDir(root, func(p string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if d.IsDir() {
			if d.Name() == ".git" {
				return filepath.SkipDir
			}
			return nil
		}
		rel, relErr := filepath.Rel(root, p)
		if relErr != nil {
			return relErr
		}
		rel = filepath.ToSlash(rel)
		lang := langFor(rel, m.Languages)
		if lang == "" {
			return nil
		}
		data, readErr := os.ReadFile(p)
		if readErr != nil {
			return readErr
		}
		src := stripComments(string(data))
		fi := core.FileImports{
			Path:    rel,
			Layer:   core.LayerOf(rel, m.Layers),
			Imports: filterIgnored(a.importsFromSource(lang, src), m.IgnoreSymbols),
		}
		if pats := m.Forbidden[fi.Layer]; len(pats) > 0 {
			fi.Constructs = constructsFromSource(src, pats)
		}
		out = append(out, fi)
		return nil
	})
	if err != nil {
		return nil, err
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Path < out[j].Path })
	return out, nil
}

func langFor(rel string, langs map[string][]string) string {
	names := make([]string, 0, len(langs))
	for n := range langs {
		names = append(names, n)
	}
	sort.Strings(names)
	for _, n := range names {
		if core.MatchGlobs(rel, langs[n]) {
			return n
		}
	}
	return ""
}

func (a Adapter) importsFromSource(lang, src string) []core.Import {
	switch lang {
	case "go":
		return dedupeSort(a.goImports(src))
	case "cpp":
		return dedupeSort(lineMatches(src, a.cppInclude))
	case "rust":
		return dedupeSort(lineMatches(src, a.rustUse, a.rustCrate))
	case "kotlin":
		return dedupeSort(lineMatches(src, a.kotlinImp))
	default:
		return nil
	}
}

func (a Adapter) goImports(src string) []core.Import {
	var imps []core.Import
	inBlock := false
	for i, ln := range strings.Split(src, "\n") {
		switch {
		case inBlock:
			if strings.Contains(ln, ")") {
				inBlock = false
			} else if mm := a.goQuoted.FindStringSubmatch(ln); mm != nil {
				imps = append(imps, core.Import{Symbol: mm[1], Line: i + 1})
			}
		case a.goBlock.MatchString(ln):
			inBlock = true
		default:
			if mm := a.goSingle.FindStringSubmatch(ln); mm != nil {
				imps = append(imps, core.Import{Symbol: mm[1], Line: i + 1})
			}
		}
	}
	return imps
}

func lineMatches(src string, res ...*regexp.Regexp) []core.Import {
	var imps []core.Import
	for i, ln := range strings.Split(src, "\n") {
		for _, re := range res {
			if mm := re.FindStringSubmatch(ln); mm != nil {
				imps = append(imps, core.Import{Symbol: mm[1], Line: i + 1})
			}
		}
	}
	return imps
}

func constructsFromSource(src string, pats []string) []core.Import {
	var cs []core.Import
	for i, ln := range strings.Split(src, "\n") {
		for _, p := range pats {
			if p != "" && strings.Contains(ln, p) {
				cs = append(cs, core.Import{Symbol: p, Line: i + 1})
			}
		}
	}
	return dedupeSort(cs)
}

func filterIgnored(imps []core.Import, ignore []string) []core.Import {
	if len(ignore) == 0 {
		return imps
	}
	var out []core.Import
	for _, imp := range imps {
		if !ignored(imp.Symbol, ignore) {
			out = append(out, imp)
		}
	}
	return out
}

func ignored(sym string, ignore []string) bool {
	for _, ig := range ignore {
		if ig != "" && strings.Contains(sym, ig) {
			return true
		}
	}
	return false
}

// dedupeSort removes duplicate symbols (keeping the first line) and sorts by
// symbol then line for a deterministic order (SPEC-DET-001).
func dedupeSort(in []core.Import) []core.Import {
	seen := map[string]bool{}
	var out []core.Import
	for _, imp := range in {
		if seen[imp.Symbol] {
			continue
		}
		seen[imp.Symbol] = true
		out = append(out, imp)
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].Symbol != out[j].Symbol {
			return out[i].Symbol < out[j].Symbol
		}
		return out[i].Line < out[j].Line
	})
	return out
}

// stripComments removes // line and /* */ block comments while preserving
// newlines so source line numbers stay aligned.
func stripComments(src string) string {
	var b strings.Builder
	for i := 0; i < len(src); {
		switch {
		case peek2(src, i, "/*"):
			i = skipBlock(src, i, &b)
		case peek2(src, i, "//"):
			i = skipLine(src, i, &b)
		default:
			b.WriteByte(src[i])
			i++
		}
	}
	return b.String()
}

func peek2(s string, i int, tok string) bool {
	return i+1 < len(s) && s[i] == tok[0] && s[i+1] == tok[1]
}

// skipBlock consumes a /* */ comment from i, preserving inner newlines.
func skipBlock(s string, i int, b *strings.Builder) int {
	for i += 2; i < len(s) && !peek2(s, i, "*/"); i++ {
		if s[i] == '\n' {
			b.WriteByte('\n')
		}
	}
	return i + 2
}

// skipLine consumes a // comment up to and including the newline.
func skipLine(s string, i int, b *strings.Builder) int {
	for i < len(s) && s[i] != '\n' {
		i++
	}
	if i < len(s) {
		b.WriteByte('\n')
		i++
	}
	return i
}
