// Package extract is the extraction adapter (ARC-003): it walks the source tree
// and yields per-file imports text-heuristically (SPEC-EXTRACT-001), plus
// forbidden-construct hits for port-impurity. It is a heuristic, not a parser:
// import-like lines in comments are stripped; the boundary is documented
// (AC-QA-02), and `markers.ignore_symbols` provides an allowlist.
package extract

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/pt9912/a-check/internal/hexagon/core"
	"github.com/pt9912/a-check/internal/hexagon/port"
)

// extractFn yields a file's imports for one language backend.
type extractFn func(src string) []core.Import

// Adapter implements port.ExtractionPort. Its compiled patterns live on the
// value (not as package globals) to satisfy the lint profile (ADR-0005).
type Adapter struct {
	goSingle, goBlock, goQuoted       *regexp.Regexp
	cppInclude                        *regexp.Regexp
	rustUse, rustCrate                *regexp.Regexp
	kotlinImp, javaImp                *regexp.Regexp
	pyImp, pyFrom                     *regexp.Regexp
	csUsing                           *regexp.Regexp
	tsFrom, tsSide, tsRequire, tsCont *regexp.Regexp
	// ktDecls extracts Kotlin top-level declaration names for the
	// declaration-aware fixed-root resolution (ADR-0023); the sole
	// declaration-aware backend in 0.18.0.
	ktDecls []*regexp.Regexp
	// backends maps a language to its extractor; its keys are the single
	// source of the supported-backend set (SPEC-EXTRACT-001). A new backend is
	// one entry — dispatch and language validation share this one map.
	backends map[string]extractFn
}

// New returns an extraction adapter.
func New() port.ExtractionPort { return newAdapter() }

func newAdapter() Adapter {
	a := Adapter{
		goSingle:   regexp.MustCompile(`^\s*import\s+(?:[\w.]+\s+)?"([^"]+)"`),
		goBlock:    regexp.MustCompile(`^\s*import\s*\(\s*$`),
		goQuoted:   regexp.MustCompile(`^\s*(?:[\w.]+\s+)?"([^"]+)"`),
		cppInclude: regexp.MustCompile(`^\s*#\s*include\s*[<"]([^>"]+)[>"]`),
		rustUse:    regexp.MustCompile(`^\s*use\s+([A-Za-z_][A-Za-z0-9_]*)`),
		rustCrate:  regexp.MustCompile(`^\s*extern\s+crate\s+([A-Za-z_][A-Za-z0-9_]*)`),
		kotlinImp:  regexp.MustCompile(`^\s*import\s+([A-Za-z_][A-Za-z0-9_.]*)`),
		javaImp:    regexp.MustCompile(`^\s*import\s+(?:static\s+)?([A-Za-z_][A-Za-z0-9_.]*)`),
		// Python: both forms yield the dotted module path; relative imports
		// (leading dot) never match [A-Za-z_] — a documented boundary of the
		// Python extraction (SPEC-EXTRACT-001): they stay unextracted even
		// though the `relative` resolution mode exists for specifier
		// languages like TypeScript (ADR-0017).
		pyImp:  regexp.MustCompile(`^\s*import\s+([A-Za-z_][A-Za-z0-9_.]*)`),
		pyFrom: regexp.MustCompile(`^\s*from\s+([A-Za-z_][A-Za-z0-9_.]*)\s+import\b`),
		// C#: using DIRECTIVES only — `global`/`static` skipped, the alias form
		// yields its target (right-hand side). The mandatory `;` right after the
		// dotted name keeps using STATEMENTS (`using var x = …;`, `using (…)`,
		// `using T x = …;`) from ever matching (SPEC-EXTRACT-001).
		csUsing: regexp.MustCompile(`^\s*(?:global\s+)?using\s+(?:static\s+)?(?:[A-Za-z_][A-Za-z0-9_]*\s*=\s*)?([A-Za-z_][A-Za-z0-9_.]*)\s*;`),
		// TypeScript: the module SPECIFIER (single or double quotes, semicolon
		// optional/ASI) — never the names left of `from`. The middle part is
		// restricted to import-clause characters (identifiers, { } * , and
		// whitespace — no `=`, `(`, `.`, no quotes), so expression lines like
		// `export const q = knex.from('users')` and dynamic `import(…)`/
		// `require(…)` never match (SPEC-EXTRACT-001). tsCont catches the
		// closing `} from '…'` line of a multi-line (Prettier-wrapped) import;
		// call chains like `db.from('x')` never lead a line with `}`.
		tsFrom:    regexp.MustCompile(`^\s*(?:import|export)\s+[\w$*{},\s]*?\bfrom\s*['"]([^'"]+)['"]`),
		tsSide:    regexp.MustCompile(`^\s*import\s*['"]([^'"]+)['"]`),
		tsRequire: regexp.MustCompile(`^\s*import\s+[\w$]+\s*=\s*require\s*\(\s*['"]([^'"]+)['"]`),
		tsCont:    regexp.MustCompile(`^\s*\}\s*from\s*['"]([^'"]+)['"]`),
	}
	a.ktDecls = kotlinDeclPatterns()
	a.backends = map[string]extractFn{
		"go":     func(src string) []core.Import { return dedupeSort(a.goImports(src)) },
		"cpp":    func(src string) []core.Import { return dedupeSort(lineMatches(src, a.cppInclude)) },
		"rust":   func(src string) []core.Import { return dedupeSort(lineMatches(src, a.rustUse, a.rustCrate)) },
		"kotlin": func(src string) []core.Import { return dedupeSort(lineMatches(src, a.kotlinImp)) },
		"java":   func(src string) []core.Import { return dedupeSort(lineMatches(src, a.javaImp)) },
		"python": func(src string) []core.Import { return dedupeSort(lineMatches(src, a.pyImp, a.pyFrom)) },
		"csharp": func(src string) []core.Import { return dedupeSort(lineMatches(src, a.csUsing)) },
		"typescript": func(src string) []core.Import {
			return dedupeSort(lineMatches(src, a.tsFrom, a.tsSide, a.tsRequire, a.tsCont))
		},
	}
	return a
}

// Extract walks root and returns the imports per source file, stably ordered.
func (a Adapter) Extract(root string, m core.Model) ([]core.FileImports, error) {
	if err := a.checkLanguages(m.Languages); err != nil {
		return nil, err
	}
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
		if core.MatchGlobs(rel, m.Exclude) {
			// exclude removes the file from the scan BEFORE extraction: it is
			// never read and yields neither imports nor forbidden-construct
			// hits (ADR-0018, SPEC-EXTRACT-001).
			return nil
		}
		lang := langFor(rel, m.Languages)
		if lang == "" {
			return nil
		}
		data, readErr := os.ReadFile(p)
		if readErr != nil {
			return readErr
		}
		src := prepSource(lang, string(data))
		fi := core.FileImports{
			Path:         rel,
			Layer:        core.LayerOf(rel, m.Layers),
			Language:     lang,
			Imports:      filterIgnored(a.importsFromSource(lang, src), m.IgnoreSymbols),
			Declarations: a.declarationsFor(lang, src),
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

// checkLanguages rejects a `languages` key outside the backend registry with a
// config error (SPEC-EXTRACT-001; the CLI maps it to exit code 2). This closes
// the silent no-op — every declared language must resolve to a real backend, so
// an unsupported/typo'd language (e.g. `ruby`, `pythn`) fails loudly instead
// of extracting nothing (false-green). Deterministic order for a stable message.
func (a Adapter) checkLanguages(langs map[string][]string) error {
	names := make([]string, 0, len(langs))
	for n := range langs {
		names = append(names, n)
	}
	sort.Strings(names)
	for _, name := range names {
		if _, ok := a.backends[name]; !ok {
			return fmt.Errorf("unbekannte Sprache %q (%s)", name, a.supportedList())
		}
	}
	return nil
}

// supportedList is the sorted backend set from the registry keys — the single
// source, so the message never drifts from the actual dispatch.
func (a Adapter) supportedList() string {
	names := make([]string, 0, len(a.backends))
	for n := range a.backends {
		names = append(names, n)
	}
	sort.Strings(names)
	return strings.Join(names, "|")
}

// importsFromSource dispatches to the language backend via the registry. There
// is no silent default: after checkLanguages, lang is always registered; an
// unregistered lang would nil-panic (loud), never extract nothing (false-green).
func (a Adapter) importsFromSource(lang, src string) []core.Import {
	return a.backends[lang](src)
}

// declarationsFor yields the file's top-level declaration names for a
// declaration-aware backend (ADR-0023). 0.18.0 has exactly one: Kotlin; every
// other language returns nil (no-op — its resolution stays at package==directory,
// SPEC-EXTRACT-001). Java is a syntax-near follow-up.
func (a Adapter) declarationsFor(lang, src string) []string {
	if lang == "kotlin" {
		return a.declarations(src)
	}
	return nil
}

// declarations runs the Kotlin top-level declaration patterns line by line and
// returns the deduplicated, sorted set of declared names (SPEC-DET-001). src is
// already comment-stripped (prepSource), so a declaration in a comment never
// counts; a declaration-like line inside a multi-line string literal is the
// documented heuristic boundary (AC-QA-02), like the import extraction.
func (a Adapter) declarations(src string) []string {
	seen := map[string]bool{}
	var out []string
	for _, ln := range strings.Split(src, "\n") {
		for _, re := range a.ktDecls {
			if mm := re.FindStringSubmatch(ln); mm != nil && mm[1] != "" && !seen[mm[1]] {
				seen[mm[1]] = true
				out = append(out, mm[1])
			}
		}
	}
	sort.Strings(out)
	return out
}

// kotlinDeclPatterns compiles the top-level declaration patterns (ADR-0023),
// each capturing the declared NAME in group 1. They anchor at line start with NO
// leading whitespace so indented members never match (top-level declarations sit
// at column 0 by convention). The building blocks: mods = an optional
// modifier/annotation prefix; gen = an optional generic parameter list; recv = an
// optional extension receiver (fun R.name captures name); nm = the declared name.
// Text-heuristic, not a parser (ADR-0002): exotic formatting is a documented
// boundary (AC-QA-02).
func kotlinDeclPatterns() []*regexp.Regexp {
	mods := `(?:@[\w.]+(?:\([^)]*\))?\s+|(?:public|private|internal|protected|open|final|abstract|sealed|data|enum|annotation|value|inner|companion|const|lateinit|inline|noinline|crossinline|external|expect|actual|suspend|operator|infix|tailrec)\s+)*`
	gen := `(?:<[^>]*>\s*)?`
	recv := `(?:[A-Za-z_][\w.]*\.)?`
	nm := `([A-Za-z_][A-Za-z0-9_]*)`
	return []*regexp.Regexp{
		regexp.MustCompile(`^` + mods + `fun\s+` + gen + recv + nm + `\s*[(<]`),
		regexp.MustCompile(`^` + mods + `(?:val|var)\s+` + gen + recv + nm + `\b`),
		regexp.MustCompile(`^` + mods + `(?:class|interface|object)\s+` + nm + `\b`),
		regexp.MustCompile(`^` + mods + `typealias\s+` + nm + `\b`),
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

// prepSource neutralizes comments per language family: the C-syntax languages
// get // and /* */ stripped. Python is NOT C-stripped — its # comment lines
// never match the line-anchored patterns anyway, and a /*-like byte sequence
// inside a Python string literal (e.g. the glob "**/*.py") would otherwise
// swallow every real import up to the next */ — a silent false-green
// (SPEC-EXTRACT-001, AC-QA-02).
func prepSource(lang, raw string) string {
	if lang == "python" {
		return raw
	}
	return stripComments(raw)
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
