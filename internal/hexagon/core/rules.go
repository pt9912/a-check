package core

import (
	"fmt"
	"path"
	"regexp"
	"sort"
	"strings"
)

// Evaluate runs the nine hexagon rules (SPEC-RULE-001) on the extracted files
// against the model and returns a stably sorted finding list (SPEC-DET-001).
// Per (file, import) the most specific rule wins (first match), so an import is
// reported once. A non-nil error is a fail-closed resolution ambiguity
// (*AmbiguousResolution, ADR-0022): the caller maps it to exit code 2 and no
// findings are reported.
func Evaluate(m Model, files []FileImports) ([]Finding, error) {
	idx := newFileIndex(files)
	var fs []Finding
	for _, f := range files {
		if matchesAny(f.Path, m.CompositionRoot) {
			fs = append(fs, compositionRootFindings(m, f)...)
			continue
		}
		for _, imp := range f.Imports {
			find, ok, err := ruleFor(m, f, imp, idx)
			if err != nil {
				return nil, err
			}
			if ok {
				fs = append(fs, find)
			}
		}
		if roleOf(f.Layer, m) == "port" {
			for _, c := range f.ForbiddenHits {
				fs = append(fs, Finding{Path: f.Path, Line: c.Line, Rule: "port-impurity", Msg: "verbotenes Konstrukt: " + c.Symbol})
			}
		}
		fs = append(fs, constructFindings(m, f, false)...)
	}
	sortFindings(fs)
	return fs, nil
}

// compositionRootFindings checks a composition-root file: it wires everything,
// so the layer rules are exempt — but tech-leak and construct-leak stay active
// for entries with `composition_root: forbid` (AC-FA-RULE-003 0.14.0 /
// AC-FA-RULE-011); entries with the allow default keep the exemption.
func compositionRootFindings(m Model, f FileImports) []Finding {
	var fs []Finding
	for _, imp := range f.Imports {
		if tech, isTech := matchTech(imp.Symbol, m.Techs); isTech && tech.ForbidCompositionRoot && !tech.inAdapter(f.Path) {
			fs = append(fs, Finding{f.Path, imp.Line, "tech-leak", "Tech " + tech.Pattern + " außerhalb " + tech.adapterLabel() + " (composition_root: forbid)"})
		}
	}
	return append(fs, constructFindings(m, f, true)...)
}

// constructFindings reports every raw-text construct hit that sits outside its
// entry's zone(s) — the construct-leak rule (AC-FA-RULE-011). It is FILE-, not
// import-bound and therefore outside the per-import first-match chain: two
// patterns hitting one line yield two findings (the total finding order keeps
// them deterministic, SPEC-DET-001). inCompositionRoot narrows the check to
// entries that switched the wiring exemption off (`composition_root: forbid`),
// mirroring tech-leak. The hit carries the entry INDEX; a stale index would
// panic loudly rather than silently skip a violation (false-green).
func constructFindings(m Model, f FileImports, inCompositionRoot bool) []Finding {
	var fs []Finding
	for _, h := range f.ConstructHits {
		c := m.Constructs[h.Entry]
		if inCompositionRoot && !c.ForbidCompositionRoot {
			continue
		}
		if c.inZone(f.Path) {
			continue
		}
		msg := "Konstrukt " + c.Pattern + " außerhalb " + c.zoneLabel()
		if inCompositionRoot {
			msg += " (composition_root: forbid)"
		}
		fs = append(fs, Finding{f.Path, h.Line, "construct-leak", msg})
	}
	return fs
}

// ruleFor returns the most specific rule violation for one import (first match),
// or ok=false if the import is clean. The purity rules dispatch on the layer's
// ROLE, not its name (AC-FA-RULE-006/007). A non-nil error is a fail-closed
// resolution ambiguity (ADR-0022) propagated from targetLayer.
func ruleFor(m Model, f FileImports, imp Import, idx fileIndex) (Finding, bool, error) {
	tl, cand, err := targetLayer(imp.Symbol, f.Path, m.Layers, m.Resolution[f.Language], idx)
	if err != nil {
		return Finding{}, false, err
	}
	srcRole := roleOf(f.Layer, m)
	tgtRole := roleOf(tl, m)
	tech, isTech := matchTech(imp.Symbol, m.Techs)
	if find, ok := impurityFinding(f, imp, srcRole, tgtRole, isTech); ok {
		return find, true, nil // core-/app-/port-impurity (domain-seitig, kategorisch)
	}
	if find := connectivityRule(m, f, imp, srcRole, tgtRole, tech, isTech, cand, tl); find.Rule != "" {
		return find, true, nil
	}
	return Finding{}, false, nil
}

// connectivityRule applies the connectivity rules in deterministic first-match
// order (SPEC-RULE-001): lateral-adapter → lateral-slice → tech-leak →
// port-direction-mismatch → port-locality → wrong-direction. Split into two
// halves (lateral/tech, then direction/locality) to keep each under the
// cyclomatic budget; the order across the split is preserved. Returns the zero
// Finding when none fires; ruleFor turns a non-zero rule name into ok=true.
func connectivityRule(m Model, f FileImports, imp Import, srcRole, tgtRole string, tech Tech, isTech bool, cand, tl string) Finding {
	if find := lateralRule(m, f, imp, srcRole, tgtRole, tech, isTech, cand, tl); find.Rule != "" {
		return find
	}
	return directionRule(m, f, imp, srcRole, tgtRole, cand, tl)
}

// lateralRule: lateral-adapter → lateral-slice → tech-leak (first-match).
func lateralRule(m Model, f FileImports, imp Import, srcRole, tgtRole string, tech Tech, isTech bool, cand, tl string) Finding {
	switch {
	case srcRole == "adapter" && tgtRole == "adapter" && lateral(m, f, cand, tl):
		return Finding{f.Path, imp.Line, "lateral-adapter", "Adapter importiert anderen Adapter " + imp.Symbol}
	case srcRole == "app" && tl == f.Layer && lateralSlice(m, f.Path, cand):
		return Finding{f.Path, imp.Line, "lateral-slice", "app-Slice importiert fremde Slice " + imp.Symbol}
	case isTech && !tech.inAdapter(f.Path):
		return Finding{f.Path, imp.Line, "tech-leak", "Tech " + tech.Pattern + " außerhalb " + tech.adapterLabel()}
	}
	return Finding{}
}

// directionRule: port-direction-mismatch → port-locality → wrong-direction
// (first-match). port-locality precedes wrong-direction so an allowed app→port
// edge cannot mask the locality violation (kategorisch, SPEC-RULE-001).
func directionRule(m Model, f FileImports, imp Import, srcRole, tgtRole, cand, tl string) Finding {
	switch {
	case srcRole == "adapter" && tgtRole == "port" && directionMismatch(m, f.Layer, tl):
		return Finding{f.Path, imp.Line, "port-direction-mismatch", f.Layer + " (" + dirOf(f.Layer, m) + ") -> " + tl + " (" + dirOf(tl, m) + "): " + imp.Symbol}
	case srcRole == "app" && tgtRole == "port" && portLocality(m, f.Path, cand):
		return Finding{f.Path, imp.Line, "port-locality", "app außerhalb Port-Scope " + portScope(cand, m) + ": " + imp.Symbol}
	case tl != "" && wrongDirection(m, f, tl):
		return Finding{f.Path, imp.Line, "wrong-direction", layerLabel(f.Layer) + " -> " + tl + " (" + imp.Symbol + ")"}
	}
	return Finding{}
}

// impurityFinding reports a purity violation for a domain/app/port source (the
// domain-side roles), or ok=false. domain is innermost — importing app/port/adapter
// or a tech is core-impurity (AC-FA-RULE-007). app may use domain+port but no
// adapter/tech; port may use domain but no adapter/tech. All categorical; the
// direction (port->core, app->port) is edge-governed (ADR-0008) and falls to
// wrong-direction in ruleFor.
func impurityFinding(f FileImports, imp Import, srcRole, tgtRole string, isTech bool) (Finding, bool) {
	var rule, who string
	switch srcRole {
	case "domain":
		if tgtRole == "app" || tgtRole == "port" || tgtRole == "adapter" || isTech {
			rule, who = "core-impurity", "Kern importiert "
		}
	case "app":
		if tgtRole == "adapter" || isTech {
			rule, who = "app-impurity", "Application importiert "
		}
	case "port":
		if tgtRole == "adapter" || isTech {
			rule, who = "port-impurity", "Port importiert "
		}
	}
	if rule == "" {
		return Finding{}, false
	}
	return Finding{f.Path, imp.Line, rule, who + imp.Symbol}, true
}

// EffectiveRole returns a layer's effective role: the explicit role: takes
// precedence (AC-FA-RULE-006), else the conventional name inference
// (core→domain, ports→port, adapters→adapter, application/app→app), else ""
// (only edge-checked). The rule engine and the graph renderer (SPEC-CLI-002)
// share this one resolver so the Mermaid colouring never drifts from the role
// the engine actually enforces.
func EffectiveRole(l Layer) string {
	if l.Role != "" {
		return l.Role
	}
	return inferRole(l.Name)
}

// roleOf returns a layer's role by name: the explicit role: (AC-FA-RULE-006), else
// the name inference, else "" (unknown layer / only edge-checked).
func roleOf(name string, m Model) string {
	l := layerByName(name, m)
	if l.Name == "" {
		return ""
	}
	return EffectiveRole(l)
}

// layerLabel names a layer for a finding message. A scanned file in NO layer has
// an empty name, which rendered as a hole ("wrong-direction:  -> ui"); it is the
// source-side symptom of the coverage gap UncoveredFiles reports, so it says so
// instead of showing nothing (ADR-0029). Only wrong-direction can see an empty
// source layer — every other rule requires a role, which requires a layer.
func layerLabel(name string) string {
	if name == "" {
		return "(ohne Schicht)"
	}
	return name
}

// dirOf returns a layer's explicit direction (driving|driven) or "". Unlike
// roleOf there is NO name inference — direction is declared only (AC-FA-RULE-008).
func dirOf(name string, m Model) string {
	return layerByName(name, m).Direction
}

// layerByName returns the layer with the given name, or a zero Layer if none.
func layerByName(name string, m Model) Layer {
	for _, l := range m.Layers {
		if l.Name == name {
			return l
		}
	}
	return Layer{}
}

// inferRole maps the conventional layer names to roles (Rückwärtskompatibilität).
func inferRole(name string) string {
	switch name {
	case "core":
		return "domain"
	case "application", "app":
		return "app"
	case "ports":
		return "port"
	case "adapters":
		return "adapter"
	default:
		return ""
	}
}

// lateral reports a forbidden adapter->adapter import (AC-FA-RULE-006). It is
// categorical — only adapter_sink exempts, not edges/allow — and fires across
// different adapter layers (layer identity) or, within one layer, across adapter
// sub-units distinguished relative to the layer's glob prefix (name-independent,
// ADR-0010). cand is the RESOLVED candidate targetLayer matched (ADR-0017):
// sink containment and sub-unit discrimination live in the layer-glob namespace
// — a raw relative specifier like "./helper" never carries the layer prefix and
// would misreport every same-sub-unit import. In path mode cand IS the raw
// import, so the pre-resolution behavior is unchanged; under dotted fixed-root
// the sink/sub-unit checks now see the resolved PATH (adapter_sink is a path
// fragment with slashes there, SPEC-RULE-001). The caller guarantees both ends
// resolve to role adapter.
func lateral(m Model, f FileImports, cand, tl string) bool {
	if contains(cand, m.AdapterSink) {
		return false
	}
	if tl != f.Layer {
		return true
	}
	layer := layerByName(f.Layer, m)
	return adapterSeg(f.Path, layer) != adapterSeg(cand, layer)
}

func wrongDirection(m Model, f FileImports, tl string) bool {
	return tl != f.Layer && !edgeAllowed(f.Layer, tl, m)
}

// directionMismatch reports an adapter->port import across opposite directions
// when BOTH sides declare one — categorical and edge-independent (AC-FA-RULE-008,
// it sits before wrong-direction in ruleFor). The caller guarantees src role
// adapter and target role port.
func directionMismatch(m Model, srcLayer, tgtLayer string) bool {
	sd, td := dirOf(srcLayer, m), dirOf(tgtLayer, m)
	return sd != "" && td != "" && sd != td
}

// sortFindings imposes the TOTAL order of SPEC-DET-001: path, line, rule, msg.
// The message is the last key on purpose — one line can carry several findings
// of the same rule (two constructs / two forbidden_constructs patterns), and
// with a non-total key the unstable sort.Slice would order those siblings by
// chance (AC-QA-01).
func sortFindings(fs []Finding) {
	sort.Slice(fs, func(i, j int) bool {
		if fs[i].Path != fs[j].Path {
			return fs[i].Path < fs[j].Path
		}
		if fs[i].Line != fs[j].Line {
			return fs[i].Line < fs[j].Line
		}
		if fs[i].Rule != fs[j].Rule {
			return fs[i].Rule < fs[j].Rule
		}
		return fs[i].Msg < fs[j].Msg
	})
}

// MatchGlobs reports whether the repo-relative path matches any of the globs.
func MatchGlobs(path string, globs []string) bool { return matchesAny(path, globs) }

// UncoveredFiles returns the scanned files that lie in NO layer glob, stably
// sorted by path (ADR-0029). Those files exist for the scan but for no layer
// rule — a deliberate fail-open boundary (AC-QA-02) that stays invisible unless
// it is reported: a green gate over a partly unchecked tree looks exactly like a
// green gate over a checked one.
//
// composition_root files are NOT counted: they are layer-less by design (the
// wiring point is exempt from the layer rules anyway). exclude files never reach
// this function — they are dropped before extraction (ADR-0018/ADR-0025).
//
// Only the SOURCE side is reported. An import TARGET that resolves to no layer
// is the same class of gap, but there it cannot be told apart from repo-EXTERNAL
// code (a third-party library) — reporting it would be noise. The cause is fixed
// from the source side anyway: declare the zone, and both sides resolve.
func UncoveredFiles(m Model, files []FileImports) []string {
	var out []string
	for _, f := range files {
		if f.Layer != "" || matchesAny(f.Path, m.CompositionRoot) {
			continue
		}
		out = append(out, f.Path)
	}
	sort.Strings(out) // stable, path-ordered (SPEC-DET-001)
	return out
}

// LayerResolution is how much of ONE layer's extracted import symbols actually
// resolve to a layer — the raw counts behind the resolution diagnosis (slice-085).
type LayerResolution struct {
	Layer    string
	Files    int
	Symbols  int
	Resolved int
}

// LayerResolutions counts, per layer, how many extracted symbols resolve to any
// layer, stably sorted by layer name. A layer whose symbols resolve to NOTHING
// while symbols were extracted is the dangerous shape: every file sits in a
// layer, every import is extracted, and still no edge is ever judged — fully
// green, fully blind.
//
// composition_root files are excluded (they are exempt from the layer rules,
// like in UncoveredFiles); files in no layer belong to the coverage diagnosis.
// A resolution error cannot occur here — the CLI only reaches the diagnosis
// after Evaluate succeeded — and is counted as unresolved if it ever did.
func LayerResolutions(m Model, files []FileImports) []LayerResolution {
	idx := newFileIndex(files)
	byLayer := map[string]*LayerResolution{}
	for _, f := range files {
		if f.Layer == "" || matchesAny(f.Path, m.CompositionRoot) {
			continue
		}
		st, ok := byLayer[f.Layer]
		if !ok {
			st = &LayerResolution{Layer: f.Layer}
			byLayer[f.Layer] = st
		}
		st.Files++
		for _, imp := range f.Imports {
			st.Symbols++
			if l, _, err := targetLayer(imp.Symbol, f.Path, m.Layers, m.Resolution[f.Language], idx); err == nil && l != "" {
				st.Resolved++
			}
		}
	}
	out := make([]LayerResolution, 0, len(byLayer))
	for _, st := range byLayer {
		out = append(out, *st)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Layer < out[j].Layer })
	return out
}

// formUnresolvableRelative names the second limit class in the diagnosis. It is
// part of the CLI contract (SPEC-CLI-001), so it lives as a constant next to the
// logic that emits it rather than being spelled out at the call site.
const formUnresolvableRelative = "relativer Pfad, den der Auflösungs-Modus %q nicht auflöst"

// HeuristicLimits returns every import line whose SPELLING keeps it from
// becoming a judged edge, stably sorted by (path, line) — the scan-wide view the
// CLI prints (ADR-0031, SPEC-CLI-001). It joins the two classes:
//
//  1. NOT EXTRACTED — the extractor already flagged these per file (Limits): a
//     relative Python import, a second directive on one line. No config knowledge
//     is needed to see them, so the adapter finds them while it reads the source.
//
//  2. EXTRACTED BUT UNRESOLVABLE UNDER THIS CONFIG — a "./" or "../" symbol
//     under any mode other than "relative" WHOSE CANDIDATES match no layer glob
//     prefix. Both halves count (ADR-0035): the spelling alone does not decide
//     it. "path" passes the symbol through with its dots, and layerOfCand looks
//     for the glob prefix on segment boundaries ANYWHERE in the candidate — so
//     "../adapters/db/x.h" against "adapters/**" DOES resolve and must not be
//     reported. This class needs the mode AND the globs, both config, which is
//     why it is derived HERE and not in the extractor.
//
// What is deliberately NOT reported: a symbol that would resolve syntactically
// and merely finds no target in this tree. That is indistinguishable from
// repo-EXTERNAL code — the same boundary at which UncoveredFiles leaves out the
// target side (AC-QA-02). Mixing a certain statement with an uncertain one would
// cost the diagnosis its credibility.
//
// The diagnosis stays TREE-FREE: it reads the mode and the globs, never the file
// index. That is what keeps class 2 on the certain side of the boundary.
func HeuristicLimits(m Model, files []FileImports) []LimitNote {
	var out []LimitNote
	for _, f := range files {
		for _, l := range f.Limits {
			out = append(out, LimitNote{Path: f.Path, Line: l.Line, Form: l.Form})
		}
		mode := m.Resolution[f.Language].Mode
		if mode == "relative" {
			continue // there the relative spelling is the resolving one
		}
		res := m.Resolution[f.Language]
		for _, imp := range f.Imports {
			if relativeSpecifier(imp.Symbol) && !anyCandHitsGlob(imp.Symbol, f.Path, res, m.Layers) {
				out = append(out, LimitNote{
					Path: f.Path,
					Line: imp.Line,
					Form: fmt.Sprintf(formUnresolvableRelative, modeName(mode)),
				})
			}
		}
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].Path != out[j].Path {
			return out[i].Path < out[j].Path
		}
		if out[i].Line != out[j].Line {
			return out[i].Line < out[j].Line
		}
		return out[i].Form < out[j].Form // total order (SPEC-DET-001)
	})
	return out
}

// anyCandHitsGlob reports whether ANY resolution candidate of imp carries a layer
// glob prefix on a segment boundary (ADR-0035). Config only — mode plus globs; the
// file index is never consulted, so the answer holds for every tree.
//
// It deliberately uses layerOfCand, the same attribution the rules use: a second,
// simpler implementation here would be a second truth about "does this resolve",
// and the two would drift.
func anyCandHitsGlob(imp, srcPath string, res ResolutionConfig, layers []Layer) bool {
	for _, cand := range resolveImport(imp, srcPath, res) {
		if l, _ := layerOfCand(cand, layers); l != "" {
			return true
		}
	}
	return false
}

// modeName spells the effective resolution mode for the message. An empty mode
// IS "path" (the default, SPEC-CONF-001); printing "" would leave the reader
// guessing at exactly the value that explains the finding.
func modeName(mode string) string {
	if mode == "" {
		return "path"
	}
	return mode
}

// sliceOf returns the vertical-slice identity of a path (AC-FA-RULE-009): the
// longest app-role layer-glob LITERAL prefix that matches path as a segment run
// (segIndex, module-prefix tolerant like the target resolution). The returned
// value is the glob prefix itself — canonical, so a source file and a
// module-qualified candidate in the SAME slice compare equal. "" if no app glob
// with a clean literal prefix matches (a `*.go`/`**/…` glob has none and never
// resolves a slice — the documented AC-QA-02 boundary; slice-isolation is opt-in
// via per-slice `**` globs).
func sliceOf(path string, m Model) string {
	best, bestLen := "", -1
	for _, l := range m.Layers {
		if EffectiveRole(l) != "app" {
			continue
		}
		for _, g := range l.Globs {
			if p := globPrefix(g); p != "" && segIndex(path, p) >= 0 && len(p) > bestLen {
				best, bestLen = p, len(p)
			}
		}
	}
	return best
}

// portScope returns the scope directory of a port candidate (AC-FA-RULE-010):
// the longest matching port-role glob literal prefix MINUS its last path segment
// (the port-dir marker, typically "ports"), so `…/createorder/ports/**` scopes
// to `…/createorder`, `…/order/ports/**` to `…/order` (business area), and
// `…/application/ports/**` to `…/application` (app-wide). "" if no port glob
// matches or the prefix has no parent segment (then port-locality does not fire).
func portScope(cand string, m Model) string {
	best, bestLen := "", -1
	for _, l := range m.Layers {
		if EffectiveRole(l) != "port" {
			continue
		}
		for _, g := range l.Globs {
			if p := globPrefix(g); p != "" && segIndex(cand, p) >= 0 && len(p) > bestLen {
				best, bestLen = p, len(p)
			}
		}
	}
	if i := strings.LastIndexByte(best, '/'); i >= 0 {
		return best[:i]
	}
	return ""
}

// lateralSlice reports a forbidden cross-slice app import (AC-FA-RULE-009). It
// fires only when BOTH source and candidate resolve to a (non-empty) slice and
// the slices differ — an unresolvable end (empty sliceOf) yields no finding
// (AC-QA-02 boundary, not a false positive). The caller guarantees the import
// stays WITHIN one app layer (tl == f.Layer): distinct app LAYERS are governed
// by edges (a declared services→services_geo edge is legitimate, b-cad), so
// slice isolation applies to the per-slice globs inside a single app layer.
// Kategorisch within that layer: not liftable via edges/allow.
func lateralSlice(m Model, srcPath, cand string) bool {
	ss, sc := sliceOf(srcPath, m), sliceOf(cand, m)
	return ss != "" && sc != "" && ss != sc
}

// portLocality reports an app import of a port outside its scope directory
// (AC-FA-RULE-010): the importer's path does not contain the port's scope as a
// segment run. It fires ONLY for ports NESTED within the application tree (the
// scope is an ancestor of some app slice) — classic hexagonal, where ports are a
// SIBLING of the app layer (`hexagon/ports` next to `hexagon/services`), has no
// per-slice locality and must not report (b-cad regression). No finding when the
// scope is undeterminable (empty). Kategorisch. The caller guarantees source role
// app and target role port; adapters that implement a port are never routed here.
func portLocality(m Model, srcPath, cand string) bool {
	sc := portScope(cand, m)
	return sc != "" && appTreeContains(m, sc) && segIndex(srcPath, sc) < 0
}

// appTreeContains reports whether scope is an ancestor of (or equal to) some
// app-role layer-glob prefix — i.e. the directory lies within the application
// tree, so a nested port's locality is meaningful (AC-FA-RULE-010). For sibling
// ports (scope outside every app subtree) it returns false and port-locality
// stays inert.
func appTreeContains(m Model, scope string) bool {
	for _, l := range m.Layers {
		if EffectiveRole(l) != "app" {
			continue
		}
		for _, g := range l.Globs {
			if p := globPrefix(g); p != "" && segIndex(p, scope) >= 0 {
				return true
			}
		}
	}
	return false
}

// LayerOf returns the name of the most specific layer whose glob matches the
// repo-relative path: the longest matching glob prefix wins (consistent with
// targetLayer, ADR-0013), the first declared layer on an equal-length tie, or
// "" if none. The match stays full-glob (matchesAny semantics, inner ** ok);
// only the choice among several matching layers switched from first-match to
// longest-prefix.
func LayerOf(relPath string, layers []Layer) string {
	best, bestLen := "", -1
	for _, l := range layers {
		if n, ok := matchSpecificity(relPath, l.Globs); ok && n > bestLen {
			best, bestLen = l.Name, n
		}
	}
	return best
}

// matchSpecificity reports whether any of the globs matches the path and, if so,
// the longest LITERAL prefix length among the MATCHING globs — the per-glob
// specificity score that mirrors targetLayer's glob loop (ADR-0013).
func matchSpecificity(path string, globs []string) (int, bool) {
	best, matched := -1, false
	for _, g := range globs {
		if globToRegexp(g).MatchString(path) {
			matched = true
			if n := litPrefixLen(g); n > best {
				best = n
			}
		}
	}
	return best, matched
}

// litPrefixLen is the byte length of a glob's literal path prefix — the part
// before the first segment that contains a wildcard (* or ?). It is the
// specificity yardstick that keeps LayerOf consistent with targetLayer, which
// can only resolve literal prefixes (segIndex): a wildcard-bearing prefix like
// "src/*/x" scores as its literal head "src", never its raw string length, and
// "**/foo" scores 0.
func litPrefixLen(g string) int { return len(literalHead(g)) }

// literalHead is a glob's literal path prefix: everything before the first
// segment carrying a wildcard. "" when the very first segment already has one.
func literalHead(g string) string {
	p := globPrefix(g)
	if i := strings.IndexAny(p, "*?"); i >= 0 {
		if j := strings.LastIndexByte(p[:i], '/'); j >= 0 {
			p = p[:j]
		} else {
			p = ""
		}
	}
	return p
}

// literalTail is the literal segment run AFTER a glob prefix's last wildcard
// segment — the marker that says which directory the glob really aims at
// ("src/hexagon/application/**/ports" → "ports"). "" when the prefix ends on the
// wildcard segment itself (then the glob names no target marker at all).
func literalTail(g string) string {
	p := globPrefix(g)
	i := strings.LastIndexAny(p, "*?")
	if i < 0 {
		return ""
	}
	j := strings.IndexByte(p[i:], '/')
	if j < 0 {
		return ""
	}
	return p[i+j+1:]
}

// fileIndex is the real scanned file set, normalized for the extensionless,
// package==directory candidate matching of the file-set-aware fixed-root
// resolution (ADR-0022, slice-027). Set membership is order-free (SPEC-DET-001);
// it is built once per Evaluate from the already-extracted file list — no I/O.
type fileIndex struct {
	// stripped holds each scanned path with its file extension stripped.
	stripped map[string]struct{}
	// dirs holds every ancestor directory of every scanned path.
	dirs map[string]struct{}
	// decls maps a directory to the set of top-level declaration names of the
	// files scanned there (ADR-0023); only a declaration-aware backend fills it.
	decls map[string]map[string]struct{}
}

func newFileIndex(files []FileImports) fileIndex {
	idx := fileIndex{stripped: map[string]struct{}{}, dirs: map[string]struct{}{}, decls: map[string]map[string]struct{}{}}
	for _, f := range files {
		idx.stripped[stripExt(f.Path)] = struct{}{}
		for d := path.Dir(f.Path); d != "." && d != "/" && d != ""; d = path.Dir(d) {
			idx.dirs[d] = struct{}{}
		}
		if len(f.Declarations) > 0 {
			dir := path.Dir(f.Path)
			names := idx.decls[dir]
			if names == nil {
				names = map[string]struct{}{}
				idx.decls[dir] = names
			}
			for _, n := range f.Declarations {
				names[n] = struct{}{}
			}
		}
	}
	return idx
}

// strength scores how strongly a resolved candidate is backed by the scan
// (extension-agnostic, package==directory — the AC-QA-02 heuristic boundary):
//   - **3 (declared):** a scanned file in the candidate's PARENT package directory
//     declares the last segment as a top-level symbol, regardless of the file name
//     (ADR-0023). This is the real evidence a declaration-aware backend (Kotlin)
//     supplies; for others idx.decls is empty and this tier never fires.
//   - **2 (exact):** a real file with the extension stripped (…/B ↔ …/B.kt). A mere
//     same-named file that does NOT declare the symbol counts only here — the file
//     name alone is not proof (ADR-0023).
//   - **1 (weak):** only the candidate's PARENT package directory exists — a symbol
//     declared in a differently-named file the declaration index missed, a
//     non-declaration-aware backend (class≠file fallback), or a package/wildcard
//     import (`a.b.*` → …/b/) that matches a whole package directory (no single
//     symbol to pin, so it fails open on a cross-layer split).
//   - **0:** phantom (nothing real).
// filterReal keeps the strongest tier present, so a real declaration beats a bare
// file-name match (the AC3 critical case), which beats a parent-dir-only candidate.
// A nested-class import resolves one level too deep and stays external (boundary).
func (idx fileIndex) strength(cand string) int {
	if strings.HasSuffix(cand, "/") {
		// A wildcard/package import (`a.b.*` → …/b/) matches a whole package
		// DIRECTORY, not a specific symbol — package-directory evidence (weak, tier
		// 1). A split package imported by wildcard therefore fails OPEN (extern) on a
		// cross-layer clash instead of breaking the scan (ADR-0023): the wildcard
		// pulls a whole package and cannot be pinned to one layer.
		if _, ok := idx.dirs[path.Clean(cand)]; ok {
			return 1
		}
		return 0
	}
	c := path.Clean(cand)
	if names, ok := idx.decls[path.Dir(c)]; ok {
		if _, ok := names[path.Base(c)]; ok {
			return 3 // declared: a real top-level declaration, file name notwithstanding
		}
	}
	if _, ok := idx.stripped[c]; ok {
		return 2
	}
	if _, ok := idx.dirs[path.Dir(c)]; ok {
		return 1
	}
	return 0
}

// denotes reports whether a candidate is backed by the scan at all (any tier).
func (idx fileIndex) denotes(cand string) bool { return idx.strength(cand) > 0 }

// stripExt drops a path's file extension (last segment only), so an extensionless
// resolved candidate can match a real scanned file.
func stripExt(p string) string {
	if e := path.Ext(p); e != "" {
		return p[:len(p)-len(e)]
	}
	return p
}

// AmbiguousResolution is the fail-closed error a file-set-aware fixed-root
// resolution raises when the SAME import resolves to real files under ≥2 roots
// that fall into DIFFERENT layers (ADR-0022): a FQN must map to at most one layer.
// The CLI maps it to exit code 2 (like a load/extract error). expect/actual in the
// SAME layer resolves cleanly and is NOT an error (the distinct-layer refinement
// of the residual left open by ADR-0020).
type AmbiguousResolution struct {
	Symbol, Src   string
	LayerA, CandA string
	LayerB, CandB string
}

func (e *AmbiguousResolution) Error() string {
	return fmt.Sprintf("mehrdeutige Mehr-Wurzel-Auflösung: %q (in %s) existiert real unter mehreren Wurzeln und fällt in verschiedene Schichten — %q (%s) vs. %q (%s); ein FQN muss in höchstens eine Schicht auflösen. Nutze disjunkte Paket-Namespaces oder root-spezifische Globs",
		e.Symbol, e.Src, e.LayerA, e.CandA, e.LayerB, e.CandB)
}

// filterReal keeps the real (scanned-file-backed) candidates of a fixed-root
// resolution with ≥2 roots (ADR-0022), keeping only the STRONGEST evidence tier
// present (declared > file-exact > parent-dir), so a real declaration beats a bare
// file-name match and a phantom whose parent is a shared/base package dir cannot
// fabricate an ambiguity (ADR-0023). It also reports whether that surviving tier is
// the WEAK one (parent-dir-only, strength 1): a clash there is fail-OPEN (extern),
// because without a declaration the layer is not discriminable — whereas a clash at
// the declared/file-exact tier is a genuine, fail-CLOSED ambiguity. Every other mode
// (path/relative/1-root) passes through unchanged (weak=false). Root order is
// preserved within the tier (deterministic, SPEC-DET-001).
func filterReal(cands []string, res ResolutionConfig, idx fileIndex) ([]string, bool) {
	if res.Mode != "fixed-root" || len(res.Roots) < 2 {
		return cands, false
	}
	best := 0
	for _, c := range cands {
		if e := idx.strength(c); e > best {
			best = e
		}
	}
	if best == 0 {
		return nil, false
	}
	var kept []string
	for _, c := range cands {
		if idx.strength(c) == best {
			kept = append(kept, c)
		}
	}
	return kept, best == 1
}

// targetLayer resolves an import string to a layer by testing whether a layer
// glob's path prefix occurs in the resolved candidate (handles module-qualified
// paths such as github.com/x/internal/core). For fixed-root with ≥2 roots the
// candidates are first filtered to the ones backed by a REAL scanned file
// (filterReal, ADR-0022), so a phantom root/… candidate no longer wins by prefix
// length — the layer is decided by the real target, not the root. Per candidate
// the most specific (longest) matching prefix wins (first declared layer on a
// length tie); across REAL candidates the earliest wins a same-layer tie
// (deterministic, SPEC-DET-001). Two candidates in DIFFERENT layers clash: at the
// declared/file-exact tier that is a genuine ambiguity → *AmbiguousResolution
// (exit 2); at the weak parent-dir-only tier there is no declaration to
// discriminate the layer, so it fails OPEN (extern, ADR-0023). It also returns the resolved
// candidate that produced the match — the layer-glob-namespace path that
// downstream path work (lateral's sub-unit/sink checks) must use instead of the
// raw specifier; in path mode it is the raw import.
func targetLayer(imp, srcPath string, layers []Layer, res ResolutionConfig, idx fileIndex) (string, string, error) {
	cands, weak := filterReal(resolveImport(imp, srcPath, res), res, idx)
	best, bestCand, bestLen := "", "", -1
	for _, cand := range cands {
		l, n := layerOfCand(cand, layers)
		switch {
		case l == "":
			// candidate matches no layer: external, contributes nothing
		case best == "":
			best, bestCand, bestLen = l, cand, n
		case l != best:
			if weak {
				// parent-dir-only ambiguity: no declaration discriminates the layer,
				// so fail OPEN (extern) instead of breaking the whole scan (ADR-0023).
				return "", "", nil
			}
			return "", "", &AmbiguousResolution{Symbol: imp, Src: srcPath, LayerA: best, CandA: bestCand, LayerB: l, CandB: cand}
		case n > bestLen:
			bestCand, bestLen = cand, n
		}
	}
	return best, bestCand, nil
}

// layerOfCand returns the layer of ONE resolved candidate and its literal glob
// prefix length: the longest layer-glob prefix occurring in cand on a segment
// boundary (segIndex>=0), the first declared layer on a length tie, or "" / -1 if
// none. Mirrors the inner loop the pre-ADR-0022 targetLayer ran over every candidate.
func layerOfCand(cand string, layers []Layer) (string, int) {
	best, bestLen := "", -1
	for _, l := range layers {
		for _, g := range l.Globs {
			if p := globPrefix(g); p != "" && segIndex(cand, p) >= 0 && len(p) > bestLen {
				best, bestLen = l.Name, len(p)
			}
		}
	}
	if shadowedByWildcardGlob(cand, best, bestLen, layers) {
		return "", -1 // undecidable ⇒ external (fail-open), never the enclosing layer
	}
	return best, bestLen
}

// shadowedByWildcardGlob reports whether the attribution to `best` must be
// WITHDRAWN because another layer owns a glob with an INNER wildcard that could
// cover the candidate just as well — a glob segIndex can never match, so it
// silently contributes nothing and the candidate falls through to the enclosing
// layer (ADR-0028).
//
// The classic shape is a nested-port glob `…/application/**/ports/**` next to the
// enclosing `…/application/**`: an adapter importing that port resolved to
// `application` and produced a wrong-direction finding on a legitimate, declared
// `adapter → ports` edge — a false positive whose obvious "fix" (declaring the
// wrong edge) permanently masks real violations. Failing OPEN instead is the
// same line ADR-0023 draws for the weak evidence tier: where the layer is not
// discriminable, do not guess.
//
// The withdrawal is deliberately narrow — three conditions, or the enclosing
// attribution stands:
//   - the other glob's literal HEAD occurs in the candidate (it reaches this region),
//   - its literal TAIL — the marker after the wildcard, e.g. `ports` — occurs too
//     (without it, `…/application/**/ports/**` would swallow EVERY application
//     import and silently disable the app layer),
//   - the head is at least as specific as the winning prefix (a more specific
//     literal glob keeps winning).
//
// A glob whose prefix ends on the wildcard segment itself (no tail marker, e.g.
// `a/**/**`) names no target and never withdraws — non-idiomatic spellings keep
// today's behavior, as with the exclude prune (ADR-0025).
func shadowedByWildcardGlob(cand, best string, bestLen int, layers []Layer) bool {
	if best == "" {
		return false // nothing attributed, nothing to withdraw
	}
	for _, l := range layers {
		if l.Name == best {
			continue
		}
		for _, g := range l.Globs {
			if !strings.ContainsAny(globPrefix(g), "*?") {
				continue // resolvable as a target: already weighed above
			}
			head, tail := literalHead(g), literalTail(g)
			if head == "" || tail == "" {
				continue
			}
			if segIndex(cand, head) >= 0 && segIndex(cand, tail) >= 0 && len(head) >= bestLen {
				return true
			}
		}
	}
	return false
}

// resolveImport normalizes an import symbol into the layer-glob namespace per
// the source language's resolution (ADR-0016). "path"/"" (Default) leaves it
// unchanged. "fixed-root" prepends each root (one candidate per root); a set
// PackageBase marks a DOTTED language — its prefix is stripped and "." mapped to
// "/" (a path language like C++ keeps its "." as file extensions). "relative"
// resolves relative specifiers against the importing file's directory
// (ADR-0017); non-relative specifiers and root escapes yield an EMPTY candidate
// set — passing the raw symbol through (like the path default) would let a bare
// import such as "@actions/core" ghost-match a "core/**" glob on a segment
// boundary. Reserved/unknown modes never reach here — the config adapter
// rejects them (Exit 2).
func resolveImport(imp, srcPath string, res ResolutionConfig) []string {
	switch res.Mode {
	case "fixed-root":
		s := imp
		if res.PackageBase != "" { // dotted language (JVM/Python): package -> path
			s = strings.TrimPrefix(s, res.PackageBase+".")
			s = strings.ReplaceAll(s, ".", "/")
		}
		if len(res.Roots) == 0 {
			return []string{s}
		}
		out := make([]string, 0, len(res.Roots))
		for _, r := range res.Roots {
			out = append(out, r+"/"+s)
		}
		return out
	case "relative":
		if !relativeSpecifier(imp) {
			return nil // bare import: no path candidate (ADR-0017, AC-QA-02)
		}
		cand := path.Clean(path.Dir(srcPath) + "/" + imp)
		if cand == ".." || strings.HasPrefix(cand, "../") {
			return nil // escapes the scan root: documented boundary (AC-QA-02)
		}
		return []string{cand}
	default:
		return []string{imp}
	}
}

// relativeSpecifier reports whether a module specifier is relative per
// ADR-0017: exactly "." or ".." (barrel imports) or a "./" / "../" prefix.
func relativeSpecifier(imp string) bool {
	return imp == "." || imp == ".." ||
		strings.HasPrefix(imp, "./") || strings.HasPrefix(imp, "../")
}

// globPrefix is the literal path prefix of a glob (before a trailing /** or /*),
// or "" for a bare wildcard.
func globPrefix(g string) string {
	p := strings.TrimSuffix(strings.TrimSuffix(g, "/**"), "/*")
	if p == "**" {
		return ""
	}
	return p
}

// segIndex returns the index at which prefix p occurs in s on path-segment
// boundaries (p starts at s[0] or right after '/', and ends at '/' or end of s),
// or -1. Segment-aware, so e.g. "io" never matches inside "audio" (ADR-0010).
func segIndex(s, p string) int {
	if p == "" {
		return -1
	}
	for from := 0; from < len(s); {
		rel := strings.Index(s[from:], p)
		if rel < 0 {
			return -1
		}
		i := from + rel
		end := i + len(p)
		startOK := i == 0 || s[i-1] == '/'
		endOK := end == len(s) || s[end] == '/'
		if startOK && endOK {
			return i
		}
		from = i + 1
	}
	return -1
}

// NewTech builds a Tech, compiling pattern as an unanchored RE2 regexp when
// match=="regex". An empty match defaults to substring. adapters carries one or
// more owning path fragments (an empty list is rejected — AC-FA-RULE-003
// 0.14.0); compositionRoot is "allow" (default, "") or "forbid". It returns an
// error for an unknown match mode, an uncompilable regex, an empty adapter list
// or an unknown compositionRoot value, which the config adapter maps to exit
// code 2 (SPEC-CONF-001 / ADR-0015).
func NewTech(pattern string, adapters []string, match, compositionRoot string) (Tech, error) {
	const kind = "tech-Muster"
	if err := validateZones(kind, pattern, adapters); err != nil {
		return Tech{}, err
	}
	forbid, err := parseCompositionRoot(kind, pattern, compositionRoot)
	if err != nil {
		return Tech{}, err
	}
	if match == "regex" && pattern == "" {
		return Tech{}, fmt.Errorf("tech-Muster: leeres regex-Pattern unzulässig (match: regex würde jeden Import treffen)")
	}
	matcher, err := compileMatch(kind, pattern, match)
	if err != nil {
		return Tech{}, err
	}
	return Tech{Pattern: pattern, Adapters: adapters, ForbidCompositionRoot: forbid, match: matcher}, nil
}

// NewConstruct builds a Construct for the `constructs` block (AC-FA-RULE-011,
// ADR-0027): same scoping mechanic as NewTech (zone list, match mode,
// composition-root switch), but matched against raw source LINES. Unlike tech it
// rejects an EMPTY pattern in BOTH match modes — a blank substring pattern would
// be a silent never-match (false-green, the class 0.14.0 closed for the empty
// adapter). Errors map to exit code 2 (SPEC-CONF-001).
func NewConstruct(pattern string, zones []string, match, compositionRoot string) (Construct, error) {
	const kind = "constructs-Muster"
	if pattern == "" {
		return Construct{}, fmt.Errorf("%s: leeres pattern unzulässig (es würde nie melden)", kind)
	}
	if err := validateZones(kind, pattern, zones); err != nil {
		return Construct{}, err
	}
	forbid, err := parseCompositionRoot(kind, pattern, compositionRoot)
	if err != nil {
		return Construct{}, err
	}
	matcher, err := compileMatch(kind, pattern, match)
	if err != nil {
		return Construct{}, err
	}
	return Construct{Pattern: pattern, Zones: zones, ForbidCompositionRoot: forbid, match: matcher}, nil
}

// validateZones rejects an empty zone list or a blank zone entry — both were
// silent never-leak entries before 0.14.0 (`strings.Contains(path, "")` is always
// true), the false-green class this project fails closed on. Both blocks spell
// the zone key `adapter`, so the message is the same for tech and constructs.
func validateZones(kind, pattern string, zones []string) error {
	if len(zones) == 0 {
		return fmt.Errorf("%s %q: leere adapter-Liste unzulässig", kind, pattern)
	}
	for _, z := range zones {
		if z == "" {
			return fmt.Errorf("%s %q: leerer adapter-Eintrag unzulässig", kind, pattern)
		}
	}
	return nil
}

// parseCompositionRoot maps the per-entry composition_root switch: "" / "allow"
// keeps the wiring exemption, "forbid" drops it; anything else fails closed.
func parseCompositionRoot(kind, pattern, v string) (bool, error) {
	switch v {
	case "", "allow":
		return false, nil
	case "forbid":
		return true, nil
	default:
		return false, fmt.Errorf("%s %q: ungültiges composition_root %q (allow|forbid)", kind, pattern, v)
	}
}

// compileMatch returns the compiled matcher for a match mode: nil for the
// substring default (the caller falls back to a Contains on the pattern), an
// unanchored RE2 for "regex" (ADR-0015 — linear, no backtracking), an error for
// an unknown mode or an uncompilable expression.
func compileMatch(kind, pattern, match string) (func(string) bool, error) {
	switch match {
	case "", "substring":
		return nil, nil
	case "regex":
		re, err := regexp.Compile(pattern)
		if err != nil {
			return nil, fmt.Errorf("%s %q: ungültige Regex: %w", kind, pattern, err)
		}
		return re.MatchString, nil
	default:
		return nil, fmt.Errorf("%s %q: ungültiges match %q (substring|regex)", kind, pattern, match)
	}
}

// inAdapter reports whether the file path lies inside ANY of the tech's owning
// adapters (the symbol is allowed in each listed adapter, AC-FA-RULE-003).
func (t Tech) inAdapter(filePath string) bool { return inAnyZone(filePath, t.Adapters) }

// adapterLabel names all owning adapters in declaration order for the finding
// message (deterministic, SPEC-RULE-001).
func (t Tech) adapterLabel() string { return strings.Join(t.Adapters, "|") }

// inZone reports whether the file lies inside ANY of the construct's zones — the
// same substring-on-path comparison tech uses (AC-FA-RULE-011).
func (c Construct) inZone(filePath string) bool { return inAnyZone(filePath, c.Zones) }

// zoneLabel names all allowed zones in declaration order for the finding message.
func (c Construct) zoneLabel() string { return strings.Join(c.Zones, "|") }

// Matches reports whether a raw source line hits this construct's pattern: the
// compiled regexp when set, otherwise a substring test. Exported because the
// extraction adapter does the line scan (SPEC-EXTRACT-001) while the pattern
// semantics stay here, next to NewConstruct.
func (c Construct) Matches(line string) bool {
	if c.match != nil {
		return c.match(line)
	}
	return c.Pattern != "" && strings.Contains(line, c.Pattern)
}

// inAnyZone reports whether the path contains any of the zone fragments — the
// shared zone test of tech-leak and construct-leak (substring on the file path,
// not segment-aware; documented since 0.14.0).
func inAnyZone(filePath string, zones []string) bool {
	for _, z := range zones {
		if contains(filePath, z) {
			return true
		}
	}
	return false
}

// matches reports whether imp hits this tech pattern: the compiled regexp when
// set, otherwise a substring test on Pattern (default / literal Tech).
func (t Tech) matches(imp string) bool {
	if t.match != nil {
		return t.match(imp)
	}
	return t.Pattern != "" && strings.Contains(imp, t.Pattern)
}

// matchTech returns the first tech (in declaration order, ADR-0015) whose
// pattern matches imp; declaration order is the tie-breaker, not longest prefix.
func matchTech(imp string, techs []Tech) (Tech, bool) {
	for _, t := range techs {
		if t.matches(imp) {
			return t, true
		}
	}
	return Tech{}, false
}

// adapterSeg returns an adapter's sub-unit within its layer: the first path
// segment after the layer's matching glob prefix — the longest matching prefix
// when a layer has several globs, mirroring targetLayer (ADR-0010). It tells two
// adapters apart inside one layer for any name, e.g. src/geometry/step vs
// src/geometry/io under a "geometry" layer. A rest WITHOUT a further directory
// (file directly at the layer root) belongs to the root sub-unit "" — sub-units
// are directories like layers, never file names; otherwise every same-directory
// x.cpp -> x.h include of a per-adapter layer misreports as lateral (ADR-0019,
// b-cad evidence: 40 false positives, 0 real). The no-match case also yields ""
// (pre-existing: no lateral without a resolvable prefix) — both collapse into
// "no sub-unit boundary crossed", which is the intended semantics.
func adapterSeg(s string, layer Layer) string {
	bestEnd, bestLen := -1, -1
	for _, g := range layer.Globs {
		p := globPrefix(g)
		if p == "" || len(p) <= bestLen {
			continue
		}
		if i := segIndex(s, p); i >= 0 {
			bestEnd, bestLen = i+len(p), len(p)
		}
	}
	if bestEnd < 0 {
		return ""
	}
	rest := strings.TrimPrefix(s[bestEnd:], "/")
	if j := strings.IndexByte(rest, '/'); j >= 0 {
		return rest[:j]
	}
	if strings.Contains(rest, ".") {
		return "" // file-shaped leaf at the layer root: root sub-unit (ADR-0019)
	}
	// Directory-shaped leaf (no dot): the leaf IS the sub-unit — a Go package
	// import ends on the package DIRECTORY (…/driven/report), and treating it
	// as root would blind cross-package lateral detection (Dogfooding-Fund
	// slice-024). Extensionless file imports (TS `./b`) share this shape and
	// stay a documented heuristic boundary (AC-QA-02, ADR-0019 Re-Eval).
	return rest
}

func edgeAllowed(from, to string, m Model) bool {
	if from == to {
		return true
	}
	for _, e := range append(append([]Edge{}, m.Edges...), m.Allow...) {
		if e.From == from && e.To == to {
			return true
		}
	}
	return false
}

func contains(s, frag string) bool { return frag != "" && strings.Contains(s, frag) }

func matchesAny(path string, globs []string) bool {
	for _, g := range globs {
		if globToRegexp(g).MatchString(path) {
			return true
		}
	}
	return false
}

func globToRegexp(glob string) *regexp.Regexp {
	var b strings.Builder
	b.WriteString("^")
	for i := 0; i < len(glob); i++ {
		switch c := glob[i]; {
		case c == '*' && i+1 < len(glob) && glob[i+1] == '*':
			b.WriteString(".*")
			i++
			if i+1 < len(glob) && glob[i+1] == '/' {
				i++
			}
		case c == '*':
			b.WriteString("[^/]*")
		case c == '?':
			b.WriteString("[^/]")
		default:
			b.WriteString(regexp.QuoteMeta(string(c)))
		}
	}
	b.WriteString("$")
	return regexp.MustCompile(b.String())
}
