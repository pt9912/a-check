package core

import (
	"regexp"
	"sort"
	"strings"
)

// Evaluate runs the five hexagon rules (SPEC-RULE-001) on the extracted files
// against the model and returns a stably sorted finding list (SPEC-DET-001).
// Per (file, import) the most specific rule wins (first match), so an import is
// reported once.
func Evaluate(m Model, files []FileImports) []Finding {
	var fs []Finding
	for _, f := range files {
		if matchesAny(f.Path, m.CompositionRoot) {
			continue // composition root wires everything; exempt from layering + tech-leak
		}
		for _, imp := range f.Imports {
			if find, ok := ruleFor(m, f, imp); ok {
				fs = append(fs, find)
			}
		}
		if roleOf(f.Layer, m) == "port" {
			for _, c := range f.Constructs {
				fs = append(fs, Finding{Path: f.Path, Line: c.Line, Rule: "port-impurity", Msg: "verbotenes Konstrukt: " + c.Symbol})
			}
		}
	}
	sortFindings(fs)
	return fs
}

// ruleFor returns the most specific rule violation for one import (first match),
// or ok=false if the import is clean. The purity rules dispatch on the layer's
// ROLE, not its name (AC-FA-RULE-006).
func ruleFor(m Model, f FileImports, imp Import) (Finding, bool) {
	tl := targetLayer(imp.Symbol, m.Layers)
	srcRole := roleOf(f.Layer, m)
	tgtRole := roleOf(tl, m)
	tech, isTech := matchTech(imp.Symbol, m.Techs)
	switch {
	// domain/port reference the domain freely but never adapters or tech; ports->core
	// is edge-governed (ADR-0008) and falls to wrong-direction below.
	case srcRole == "domain" && (tgtRole == "adapter" || isTech):
		return Finding{f.Path, imp.Line, "core-impurity", "Kern importiert " + imp.Symbol}, true
	case srcRole == "port" && (tgtRole == "adapter" || isTech):
		return Finding{f.Path, imp.Line, "port-impurity", "Port importiert " + imp.Symbol}, true
	case srcRole == "adapter" && tgtRole == "adapter" && lateral(m, f, imp, tl):
		return Finding{f.Path, imp.Line, "lateral-adapter", "Adapter importiert anderen Adapter " + imp.Symbol}, true
	case isTech && !strings.Contains(f.Path, tech.Adapter):
		return Finding{f.Path, imp.Line, "tech-leak", "Tech " + tech.Pattern + " außerhalb " + tech.Adapter}, true
	case tl != "" && wrongDirection(m, f, tl):
		return Finding{f.Path, imp.Line, "wrong-direction", f.Layer + " -> " + tl + " (" + imp.Symbol + ")"}, true
	}
	return Finding{}, false
}

// roleOf returns a layer's role: the explicit role: (AC-FA-RULE-006), else the
// name inference, else "" (unknown layer / only edge-checked).
func roleOf(name string, m Model) string {
	switch l := layerByName(name, m); {
	case l.Name == "":
		return ""
	case l.Role != "":
		return l.Role
	default:
		return inferRole(name)
	}
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
// ADR-0010). The caller guarantees both ends resolve to role adapter.
func lateral(m Model, f FileImports, imp Import, tl string) bool {
	if contains(imp.Symbol, m.AdapterSink) {
		return false
	}
	if tl != f.Layer {
		return true
	}
	layer := layerByName(f.Layer, m)
	return adapterSeg(f.Path, layer) != adapterSeg(imp.Symbol, layer)
}

func wrongDirection(m Model, f FileImports, tl string) bool {
	return tl != f.Layer && !edgeAllowed(f.Layer, tl, m)
}

func sortFindings(fs []Finding) {
	sort.Slice(fs, func(i, j int) bool {
		if fs[i].Path != fs[j].Path {
			return fs[i].Path < fs[j].Path
		}
		if fs[i].Line != fs[j].Line {
			return fs[i].Line < fs[j].Line
		}
		return fs[i].Rule < fs[j].Rule
	})
}

// MatchGlobs reports whether the repo-relative path matches any of the globs.
func MatchGlobs(path string, globs []string) bool { return matchesAny(path, globs) }

// LayerOf returns the name of the first layer whose glob matches the
// repo-relative path, or "" if none.
func LayerOf(relPath string, layers []Layer) string {
	for _, l := range layers {
		if matchesAny(relPath, l.Globs) {
			return l.Name
		}
	}
	return ""
}

// targetLayer resolves an import string to a layer by testing whether a layer
// glob's path prefix occurs in the import (handles module-qualified paths such
// as github.com/x/internal/core). The most specific (longest) matching prefix
// wins — the first declared layer on an equal-length tie — so nested layers
// resolve correctly (ADR-0010).
func targetLayer(imp string, layers []Layer) string {
	best, bestLen := "", -1
	for _, l := range layers {
		for _, g := range l.Globs {
			if p := globPrefix(g); p != "" && segIndex(imp, p) >= 0 && len(p) > bestLen {
				best, bestLen = l.Name, len(p)
			}
		}
	}
	return best
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

func matchTech(imp string, techs []Tech) (Tech, bool) {
	for _, t := range techs {
		if t.Pattern != "" && strings.Contains(imp, t.Pattern) {
			return t, true
		}
	}
	return Tech{}, false
}

// adapterSeg returns an adapter's sub-unit within its layer: the first path
// segment after the layer's matching glob prefix — the longest matching prefix
// when a layer has several globs, mirroring targetLayer (ADR-0010). It tells two
// adapters apart inside one layer for any name, e.g. src/geometry/step vs
// src/geometry/io under a "geometry" layer.
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
