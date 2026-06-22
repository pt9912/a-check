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
		if f.Layer == "ports" {
			for _, c := range f.Constructs {
				fs = append(fs, Finding{Path: f.Path, Line: c.Line, Rule: "port-impurity", Msg: "verbotenes Konstrukt: " + c.Symbol})
			}
		}
	}
	sortFindings(fs)
	return fs
}

// ruleFor returns the most specific rule violation for one import (first match),
// or ok=false if the import is clean.
func ruleFor(m Model, f FileImports, imp Import) (Finding, bool) {
	tl := targetLayer(imp.Symbol, m.Layers)
	tech, isTech := matchTech(imp.Symbol, m.Techs)
	switch {
	case f.Layer == "core" && (tl == "adapters" || isTech):
		return Finding{f.Path, imp.Line, "core-impurity", "Kern importiert " + imp.Symbol}, true
	// port-impurity: ports speak the domain's language — core references are fine and
	// edge-governed (ADR-0008); ports must not import adapters or tech. Symmetric to
	// core-impurity above.
	case f.Layer == "ports" && (tl == "adapters" || isTech):
		return Finding{f.Path, imp.Line, "port-impurity", "Port importiert " + imp.Symbol}, true
	case f.Layer == "adapters" && tl == "adapters" && lateral(m, f, imp):
		return Finding{f.Path, imp.Line, "lateral-adapter", "Adapter importiert anderen Adapter " + imp.Symbol}, true
	case isTech && !strings.Contains(f.Path, tech.Adapter):
		return Finding{f.Path, imp.Line, "tech-leak", "Tech " + tech.Pattern + " außerhalb " + tech.Adapter}, true
	case tl != "" && wrongDirection(m, f, tl):
		return Finding{f.Path, imp.Line, "wrong-direction", f.Layer + " -> " + tl + " (" + imp.Symbol + ")"}, true
	}
	return Finding{}, false
}

func lateral(m Model, f FileImports, imp Import) bool {
	return adapterSeg(f.Path) != adapterSeg(imp.Symbol) && !contains(imp.Symbol, m.AdapterSink)
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
// as github.com/x/internal/core).
func targetLayer(imp string, layers []Layer) string {
	for _, l := range layers {
		for _, g := range l.Globs {
			p := strings.TrimSuffix(g, "/**")
			p = strings.TrimSuffix(p, "/*")
			if p != "" && p != "**" && strings.Contains(imp, p) {
				return l.Name
			}
		}
	}
	return ""
}

func matchTech(imp string, techs []Tech) (Tech, bool) {
	for _, t := range techs {
		if t.Pattern != "" && strings.Contains(imp, t.Pattern) {
			return t, true
		}
	}
	return Tech{}, false
}

// adapterSeg returns the path segment immediately after "adapters", used to tell
// two adapters apart for the lateral-adapter rule.
func adapterSeg(s string) string {
	parts := strings.Split(s, "/")
	for i, p := range parts {
		if p == "adapters" && i+1 < len(parts) {
			return parts[i+1]
		}
	}
	return ""
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
