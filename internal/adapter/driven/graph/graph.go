// Package graph renders the resolved architecture model as a Mermaid flowchart
// (ARC-007): the presentation adapter behind the no-scan --print-graph mode
// (SPEC-CLI-002). It is PURE — it imports only core (and the driven port) and
// returns a string, no I/O — so a-check's own arch-check keeps it adapter-clean
// (it never imports a sibling adapter). The composition root writes the
// returned string to stdout.
package graph

import (
	"sort"
	"strconv"
	"strings"

	"github.com/pt9912/a-check/internal/hexagon/core"
	"github.com/pt9912/a-check/internal/hexagon/port"
)

// Adapter is the stateless Mermaid renderer.
type Adapter struct{}

// New returns the graph presentation adapter as a GraphPort.
func New() port.GraphPort { return Adapter{} }

// classDefs are emitted in a FIXED order (determinism, SPEC-CLI-002): one class
// per effective role plus the special classes. Unused classes are harmless.
const classDefs = `    classDef domain fill:#fff4d6,stroke:#d4a017,color:#000
    classDef app fill:#e6e0ff,stroke:#6a4bbf,color:#000
    classDef port fill:#e0f0e0,stroke:#3a8a3a,color:#000
    classDef adapter fill:#e0ecff,stroke:#2a6ad4,color:#000
    classDef exempt fill:#f0f0f0,stroke:#999999,stroke-dasharray:4 3,color:#000
    classDef dangling fill:#ffe0e0,stroke:#cc3333,stroke-dasharray:2 2,color:#000
    classDef legend fill:#fafafa,stroke:#bbbbbb,color:#000
`

// Render maps the config model to a DETERMINISTIC Mermaid flowchart string
// (SPEC-CLI-002): stable internal node IDs (Lᵢ/Dⱼ/C0/S0), the raw layer names
// only as escaped labels, edges/allow as (dashed, labelled) links, one classDef
// per effective role, driving/driven grouped into subgraphs, and the implicit
// categorical rules as a legend note (never as an edge, AC-QA-02).
func (Adapter) Render(m core.Model) string {
	layers := sortedLayers(m.Layers)
	id := layerIDs(layers)
	dangling, did := danglingNodes(m, id)
	nodeID := func(name string) string {
		if v, ok := id[name]; ok {
			return v
		}
		return did[name]
	}

	var b strings.Builder
	b.WriteString("flowchart TB\n")
	writeLayerNodes(&b, layers, id)
	writeSpecialNodes(&b, m, dangling, did)
	writeLinks(&b, m.Edges, nodeID, " --> ")
	writeLinks(&b, m.Allow, nodeID, " -.->|allow| ")
	writeLegend(&b)
	b.WriteString(classDefs)
	return b.String()
}

// sortedLayers returns a copy of the layers sorted by name (stable IDs).
func sortedLayers(in []core.Layer) []core.Layer {
	out := append([]core.Layer(nil), in...)
	sort.Slice(out, func(i, j int) bool { return out[i].Name < out[j].Name })
	return out
}

// layerIDs assigns the stable internal ID L0, L1, … in sorted order.
func layerIDs(layers []core.Layer) map[string]string {
	id := make(map[string]string, len(layers))
	for i, l := range layers {
		id[l.Name] = "L" + strconv.Itoa(i)
	}
	return id
}

// danglingNodes collects edge/allow endpoints that name no declared layer and
// assigns them stable D-IDs (SPEC-CLI-002 — shown as notes, never a new exit 2).
func danglingNodes(m core.Model, id map[string]string) ([]string, map[string]string) {
	set := map[string]bool{}
	mark := func(name string) {
		if _, ok := id[name]; !ok {
			set[name] = true
		}
	}
	for _, e := range m.Edges {
		mark(e.From)
		mark(e.To)
	}
	for _, e := range m.Allow {
		mark(e.From)
		mark(e.To)
	}
	names := make([]string, 0, len(set))
	for name := range set {
		names = append(names, name)
	}
	sort.Strings(names)
	did := make(map[string]string, len(names))
	for i, name := range names {
		did[name] = "D" + strconv.Itoa(i)
	}
	return names, did
}

// writeLayerNodes emits the layer nodes: driving/driven into stable subgraphs,
// EVERY other layer (including any unexpected direction value) at top level, so
// no node an edge references is ever left undefined — the renderer does not rely
// on the config validator's direction check for this invariant (SPEC-CLI-002).
func writeLayerNodes(b *strings.Builder, layers []core.Layer, id map[string]string) {
	for _, l := range layers {
		if l.Direction != "driving" && l.Direction != "driven" {
			writeLayerNode(b, l, id)
		}
	}
	writeDirGroup(b, layers, id, "driving")
	writeDirGroup(b, layers, id, "driven")
}

// writeLayerNode emits one layer node: internal ID, escaped name + glob sublines,
// and the effective-role class (shared core resolver, no copied inference).
func writeLayerNode(b *strings.Builder, l core.Layer, id map[string]string) {
	b.WriteString("    " + id[l.Name] + `["` + labelWithGlobs(escapeLabel(l.Name), l.Globs) + `"]`)
	if role := core.EffectiveRole(l); role != "" {
		b.WriteString(":::" + role)
	}
	b.WriteString("\n")
}

// labelWithGlobs appends each glob as an escaped <br/> subline to an already-safe
// head (SPEC-CLI-002 label structure) — one place for the layer, composition_root
// and adapter_sink note labels.
func labelWithGlobs(head string, globs []string) string {
	label := head
	for _, g := range globs {
		label += "<br/>" + escapeLabel(g)
	}
	return label
}

// writeDirGroup emits one direction subgraph if it holds any layer.
func writeDirGroup(b *strings.Builder, layers []core.Layer, id map[string]string, dir string) {
	var group []core.Layer
	for _, l := range layers {
		if l.Direction == dir {
			group = append(group, l)
		}
	}
	if len(group) == 0 {
		return
	}
	b.WriteString("    subgraph dir_" + dir + `["` + dir + `"]` + "\n")
	for _, l := range group {
		writeLayerNode(b, l, id)
	}
	b.WriteString("    end\n")
}

// writeSpecialNodes emits the composition_root / adapter_sink note nodes (no
// wiring edges) and the sorted dangling notes (SPEC-CLI-002).
func writeSpecialNodes(b *strings.Builder, m core.Model, dangling []string, did map[string]string) {
	if len(m.CompositionRoot) > 0 {
		b.WriteString("    C0[\"" + labelWithGlobs("composition_root", m.CompositionRoot) + "\"]:::exempt\n")
	}
	if m.AdapterSink != "" {
		b.WriteString("    S0[\"" + labelWithGlobs("adapter_sink", []string{m.AdapterSink}) + "\"]:::exempt\n")
	}
	for _, name := range dangling {
		b.WriteString("    " + did[name] + `["` + escapeLabel(name) + `"]:::dangling` + "\n")
	}
}

// writeLegend emits the legend note: the implicit, categorical rules are NEVER
// edges (AC-QA-02 — the graph asserts no semantics about the real code).
func writeLegend(b *strings.Builder) {
	// The label wraps a left-aligning <div>: Mermaid centers node text by default;
	// the inline text-align works in permissive renderers and is harmlessly ignored
	// (falls back to centered) where a strict sanitizer strips the style. Fixed
	// renderer text — like the <br/> breaks, it is NOT run through the escaper.
	b.WriteString("    LEGEND[\"<div style='text-align:left'>Legende — implizite Regeln<br/>(nie als Kante gezeichnet):<br/>" +
		"core-impurity<br/>" +
		"lateral-adapter<br/>" +
		"lateral-slice<br/>" +
		"port-direction-mismatch<br/>" +
		"port-locality<br/><br/>" +
		"durchgezogen = edges<br/>" +
		"gestrichelt = allow<br/>" +
		"Farbe = effektive Rolle</div>\"]:::legend\n")
}

// writeLinks emits one link per edge (stably sorted, SPEC-DET-001) using the
// given connector (" --> " for edges, " -.->|allow| " for the abgesetzte
// allow-Kante), always via internal node IDs — never raw names.
func writeLinks(b *strings.Builder, edges []core.Edge, nodeID func(string) string, connector string) {
	for _, e := range sortEdges(edges) {
		b.WriteString("    " + nodeID(e.From) + connector + nodeID(e.To) + "\n")
	}
}

// sortEdges returns a copy sorted by (from, to) over the raw names (SPEC-DET-001).
func sortEdges(edges []core.Edge) []core.Edge {
	out := append([]core.Edge(nil), edges...)
	sort.SliceStable(out, func(i, j int) bool {
		if out[i].From != out[j].From {
			return out[i].From < out[j].From
		}
		return out[i].To < out[j].To
	})
	return out
}

// escapeLabel implements the SPEC-CLI-002 escaping contract: user-controlled
// text is first single-lined (\r/\n/\t → space, other control chars → ?) and
// then encoded for Mermaid/HTML delimiters in a fixed order so user text can
// never produce Mermaid syntax. Encoding on the ORIGINAL rune (one pass) avoids
// double-encoding: the `&` written for `&amp;` is never re-processed.
func escapeLabel(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	for _, r := range s {
		switch r {
		case '\r', '\n', '\t':
			b.WriteByte(' ')
		case '&':
			b.WriteString("&amp;")
		case '"':
			b.WriteString("&quot;")
		case '<':
			b.WriteString("&lt;")
		case '>':
			b.WriteString("&gt;")
		case '[':
			b.WriteString("&#91;")
		case ']':
			b.WriteString("&#93;")
		case '|':
			b.WriteString("&#124;")
		case '\\':
			b.WriteString("&#92;")
		case '`':
			b.WriteString("&#96;")
		default:
			if r < 0x20 || r == 0x7f {
				b.WriteByte('?')
			} else {
				b.WriteRune(r)
			}
		}
	}
	return b.String()
}
