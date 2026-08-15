// Package extract is the extraction adapter (ARC-003): it walks the source tree
// and yields per-file imports text-heuristically (SPEC-EXTRACT-001), plus
// forbidden-construct hits for port-impurity and raw-text construct hits for
// construct-leak (ADR-0027). It is a heuristic, not a parser:
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
	// limits maps a language to the counter-patterns that DIAGNOSE its
	// unextracted spellings (ADR-0031). Deliberately a separate map from
	// backends: a language may have a backend and no counter-pattern (C++,
	// whose limit is resolution-side and derived in core.HeuristicLimits).
	limits map[string][]limitPattern
}

// limitPattern is one counter-pattern naming a heuristic limit (ADR-0031). It is
// broader than the extraction patterns on purpose: it matches exactly the
// spellings the backend does NOT turn into an import, so the scan can say where
// it is blind instead of leaving the consumer to find out by planting a
// violation (AC-QA-02).
type limitPattern struct {
	re   *regexp.Regexp
	form string
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
	a.limits = limitPatterns()
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

// Validate checks the config's language backends against the registry without
// walking any file (SPEC-CLI-002): the no-scan --print-graph path reuses the
// exact backend validation of a scan without reading sources. Extract runs the
// same check, so there is no duplicated language logic.
func (a Adapter) Validate(m core.Model) error {
	return a.checkLanguages(m.Languages)
}

// Extract walks root and returns the imports per source file, stably ordered.
func (a Adapter) Extract(root string, m core.Model) ([]core.FileImports, error) {
	if err := a.Validate(m); err != nil {
		return nil, err
	}
	var out []core.FileImports
	err := filepath.WalkDir(root, func(p string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		rel, relErr := filepath.Rel(root, p)
		if relErr != nil {
			return relErr
		}
		rel = filepath.ToSlash(rel)
		if d.IsDir() {
			return dirAction(rel, d.Name(), m.Exclude)
		}
		if core.MatchGlobs(rel, m.Exclude) {
			// exclude removes the file from the scan BEFORE extraction: it is
			// never read and yields neither imports nor forbidden-construct
			// hits (ADR-0018, SPEC-EXTRACT-001).
			return nil
		}
		fi, keep, fErr := a.fileImports(p, rel, m)
		if fErr != nil {
			return fErr
		}
		if keep {
			out = append(out, fi)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Path < out[j].Path })
	return out, nil
}

// dirAction decides whether the walk descends into a directory or prunes it.
// .git is always skipped. `exclude` prunes the directory WALK, not just file
// extraction (ADR-0025, SPEC-EXTRACT-001): a directory whose WHOLE subtree is
// excluded is not descended into. Pruning happens before ReadDir (WalkDir
// contract), so an unreadable or huge excluded subtree no longer aborts the
// scan; a NON-excluded unreadable directory still fails fail-closed via the
// walkErr path (AC-QA-02). The root entry (rel == ".") is never pruned; single
// files are still excluded by the caller (ADR-0018).
func dirAction(rel, name string, exclude []string) error {
	if name == ".git" {
		return filepath.SkipDir
	}
	if rel != "." && dirExcluded(rel, exclude) {
		return filepath.SkipDir
	}
	return nil
}

// dirExcluded reports whether EVERY path under dir is covered by an exclude
// glob — the only sound condition for pruning the walk. A glob prunes dir iff it
// is a recursive-subtree pattern: "**" (everything) or "<prefix>/**" whose
// prefix matches dir exactly. Only "**" spans directory separators in the glob
// engine (core.globToRegexp → ".*"); a single-segment "<dir>/*", a trailing-slash
// "<dir>/", or a file pattern ("<dir>/*.go") covers only PART of the subtree, so
// pruning on those would silently drop files the glob does not match — a
// false-green (AC-QA-02). Restricting the prune to "<prefix>/**" keeps it
// provably output-equivalent to the per-file exclude: every file under a pruned
// dir already matches "<prefix>/.*" and would be excluded anyway; the only
// observable effect is not descending (no abort on an unreadable/huge excluded
// subtree, faster). Only the canonical spellings "**" and "<prefix>/**" prune;
// non-idiomatic equivalents ("**/", "foo/**/", "a/**/**") fall back to file-level
// exclusion — output stays correct, only the prune optimisation is skipped (a
// documented heuristic boundary, AC-QA-02).
func dirExcluded(dir string, exclude []string) bool {
	for _, g := range exclude {
		if g == "**" {
			return true
		}
		if prefix, ok := strings.CutSuffix(g, "/**"); ok && core.MatchGlobs(dir, []string{prefix}) {
			return true
		}
	}
	return false
}

// fileImports extracts one source file's imports and forbidden-construct hits.
// keep is false when the path matches no language backend — skipped, not an
// error (SPEC-EXTRACT-001).
func (a Adapter) fileImports(p, rel string, m core.Model) (core.FileImports, bool, error) {
	lang := langFor(rel, m.Languages)
	if lang == "" {
		return core.FileImports{}, false, nil
	}
	data, err := os.ReadFile(p)
	if err != nil {
		return core.FileImports{}, false, err
	}
	src := prepSource(lang, string(data))
	fi := core.FileImports{
		Path:         rel,
		Layer:        core.LayerOf(rel, m.Layers),
		Language:     lang,
		Imports:      filterIgnored(a.importsFromSource(lang, src), m.IgnoreSymbols),
		Limits:       a.limitsFromSource(lang, src),
		Declarations: a.declarationsFor(lang, src),
	}
	if pats := m.Forbidden[fi.Layer]; len(pats) > 0 {
		fi.ForbiddenHits = constructsFromSource(src, pats)
	}
	fi.ConstructHits = constructHits(src, m.Constructs)
	return fi, true, nil
}

// constructHits reports every `constructs` match in the PREPARED source — the
// same comment-stripped text the imports and forbidden_constructs see, so a hit
// living only in a comment never fires (a declared divergence from a grep
// reference, AC-FA-RULE-011/ADR-0027). It runs for EVERY scanned file, layer or
// not: the monopoly is a statement about the whole tree. Each hit carries the
// ENTRY INDEX (not the pattern text) so the rule engine decides the zone on the
// entry itself; entry-major iteration yields the (entry, line) order of
// SPEC-EXTRACT-001 without a sort.
func constructHits(src string, cs []core.Construct) []core.ConstructHit {
	if len(cs) == 0 {
		return nil
	}
	lines := strings.Split(src, "\n")
	var out []core.ConstructHit
	for e := range cs {
		for i, ln := range lines {
			if cs[e].Matches(ln) {
				out = append(out, core.ConstructHit{Entry: e, Line: i + 1})
			}
		}
	}
	return out
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

// The two diagnosed forms. They are the CLI's output text (SPEC-CLI-001), so
// they are constants: a reworded string here silently reworders the contract.
const (
	formRelativeImport  = "relativer Import — von diesem Backend nicht extrahiert"
	formSecondDirective = "zweite Direktive auf derselben Zeile — nur die erste wird extrahiert"
)

// limitPatterns compiles the counter-patterns per language (ADR-0031).
//
// The second-directive patterns are the backend's own pattern RE-ANCHORED at a
// preceding `;` instead of the line start — the narrowest shape that still finds
// the case. A loose `;.*import` would fire on any semicolon-plus-word line; this
// form requires a complete second directive, so a string literal or an unrelated
// statement never matches. C# `using var x = 1; using var y = 2;` (using
// STATEMENTS) stays clear for the same reason the extraction pattern does: the
// mandatory `;` right after a dotted name.
//
// Go has no entry: several `import` directives on one line are not legal Go, so
// there is nothing to miss. C++ has none either — its diagnosed limit is
// resolution-side (a `../` include under a non-`relative` mode) and is derived in
// core.HeuristicLimits from the EXTRACTED symbol, not from a source line.
//
// This map is the maintenance cost ADR-0031 accepted with open eyes: a limit
// without an entry here stays invisible. It is measured against the Out-of-Scope
// paragraph of AC-FA-EXTRACT-001.
func limitPatterns() map[string][]limitPattern {
	second := func(re string) limitPattern {
		return limitPattern{re: regexp.MustCompile(re), form: formSecondDirective}
	}
	return map[string][]limitPattern{
		"python": {
			// leading dot: `from . import x`, `from ..pkg import y` — never
			// matched by pyFrom's [A-Za-z_] start (a documented boundary).
			{re: regexp.MustCompile(`^\s*from\s+\.`), form: formRelativeImport},
			// `import a, b` — idiomatic Python and the practically relevant case.
			second(`^\s*import\s+[A-Za-z_][A-Za-z0-9_.]*(?:\s+as\s+[A-Za-z_][A-Za-z0-9_]*)?\s*,`),
		},
		"csharp":     {second(`;\s*(?:global\s+)?using\s+(?:static\s+)?(?:[A-Za-z_][A-Za-z0-9_]*\s*=\s*)?[A-Za-z_][A-Za-z0-9_.]*\s*;`)},
		"java":       {second(`;\s*import\s+(?:static\s+)?[A-Za-z_][A-Za-z0-9_.]*\s*;`)},
		"kotlin":     {second(`;\s*import\s+[A-Za-z_][A-Za-z0-9_.]*`)},
		"rust":       {second(`;\s*(?:use\s+|extern\s+crate\s+)[A-Za-z_][A-Za-z0-9_]*`)},
		"typescript": {second(`;\s*(?:import|export)\s+[\w$*{},\s]*?\bfrom\s*['"]`)},
	}
}

// limitsFromSource runs the language's counter-patterns over the PREPARED source
// — the same comment-stripped text the extraction sees, so a commented-out
// import never produces a diagnosis. A language without counter-patterns yields
// nil (no diagnosis, not an error).
//
// The counter-patterns inherit the extraction's boundary exactly: Python is not
// comment-stripped (prepSource), so an import-shaped line inside a triple-quoted
// string can produce a diagnosis — the same way it already produces a phantom
// IMPORT today (AC-QA-02). The diagnosis is no more heuristic than what it
// diagnoses; making it stricter than the extraction would report a limit that is
// not there.
func (a Adapter) limitsFromSource(lang, src string) []core.Limit {
	pats := a.limits[lang]
	if len(pats) == 0 {
		return nil
	}
	var out []core.Limit
	for i, ln := range strings.Split(src, "\n") {
		for _, p := range pats {
			if p.re.MatchString(ln) {
				out = append(out, core.Limit{Line: i + 1, Form: p.form})
			}
		}
	}
	return out
}

// lineMatches greift je Zeile und Regex den ERSTEN Treffer.
//
// HEURISTIK-GRENZE (AC-QA-02, dokumentiert in AC-FA-EXTRACT-001 §Out-of-Scope):
// stehen MEHRERE Direktiven auf einer Zeile, wird nur die erste extrahiert —
// `import a, b` liefert `a`, `using A; using B;` liefert `A`. Die Regexe sind
// `^\s*…`-verankert, FindStringSubmatch nimmt den ersten Treffer; ein
// FindAllStringSubmatch würde daran nichts ändern, solange die Verankerung
// steht. Betroffen sind alle Backends ausser C++ (Präprozessor-Direktiven sind
// ohnehin zeilenweise); praktisch relevant vor allem Python, wo `import a, b`
// idiomatisch ist.
//
// Warum hier nur ein Kommentar und kein Fix: die Grenze zu SCHLIESSEN ist eine
// Verhaltensänderung mit Vertragsbezug. slice-084 hat sie sichtbar gemacht —
// im Benutzerhandbuch §6 neben den übrigen Formen, die beim Konfigurieren
// auffallen — statt sie stillschweigend zu verschieben.
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
		case src[i] == '"' || src[i] == '\'' || src[i] == '`':
			// ZUERST: ein Literal wird VERBATIM uebernommen, damit ein `/*` oder
			// `//` darin keinen Phantom-Kommentar oeffnet (slice-090, ADR-0034).
			i = copyLiteral(src, i, &b)
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

// copyLiteral copies a string/char literal verbatim, including its delimiters,
// and returns the index just past it. Everything inside is opaque: a `/*` there
// is text, not a comment opener (ADR-0034).
//
// Quote-relative rules, deliberately coarse — this is a comment stripper, not a
// lexer (ADR-0002):
//
//   - Backslash escapes the next byte, EXCEPT in backtick literals (Go raw
//     strings and TS templates have no escapes).
//   - A `"` or `'` literal also ends at a NEWLINE. Neither is multi-line in any
//     supported language, so an unbalanced quote costs at most its own line —
//     without this, a stray apostrophe would make the rest of the file opaque
//     and stop comment stripping entirely. Backtick literals ARE multi-line and
//     run to their closing delimiter.
//   - Unterminated at EOF: the rest is copied verbatim. That errs toward
//     KEEPING text, never toward swallowing it — the direction this slice exists
//     to enforce.
//
// Known boundary (AC-QA-02): a Rust lifetime (`&'a str`) is an unbalanced
// apostrophe, so the remainder of THAT line is treated as a literal and its
// comments stay in. The failure direction is a visible false positive, not a
// silent false negative.
func copyLiteral(s string, i int, b *strings.Builder) int {
	q := s[i]
	b.WriteByte(q)
	for i++; i < len(s); i++ {
		if s[i] == '\\' && q != '`' && i+1 < len(s) {
			b.WriteByte(s[i])
			b.WriteByte(s[i+1])
			i++
			continue
		}
		b.WriteByte(s[i])
		if s[i] == q {
			return i + 1
		}
		if s[i] == '\n' && q != '`' {
			return i + 1 // unterminiert: endet an der Zeile
		}
	}
	return i
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
