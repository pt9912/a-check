package graph_test

import (
	"strings"
	"testing"

	"github.com/pt9912/a-check/internal/adapter/driven/graph"
	"github.com/pt9912/a-check/internal/hexagon/core"
)

func render(m core.Model) string { return graph.New().Render(m) }

// TestRenderRolesEdgesHeader: klassische Config ohne role: → effektive Rollen
// via Kern-Inferenz; Kanten referenzieren interne IDs, nie Roh-Namen
// (SPEC-CLI-002, AC-FA-CLI-002 effektive Rollen).
func TestRenderRolesEdgesHeader(t *testing.T) {
	m := core.Model{
		Layers: []core.Layer{
			{Name: "adapters", Globs: []string{"internal/adapters/**"}},
			{Name: "core", Globs: []string{"internal/core/**"}},
			{Name: "ports", Globs: []string{"internal/ports/**"}},
		},
		Edges: []core.Edge{{From: "adapters", To: "core"}, {From: "ports", To: "core"}},
	}
	out := render(m)
	for _, want := range []string{
		"flowchart TB",
		`["adapters<br/>internal/adapters/**"]:::adapter`,
		`["core<br/>internal/core/**"]:::domain`,
		`["ports<br/>internal/ports/**"]:::port`,
		"classDef domain",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("Render fehlt %q\n%s", want, out)
		}
	}
	// Sortiert nach Name: adapters=L0, core=L1, ports=L2. Kanten via IDs.
	if !strings.Contains(out, "    L0 --> L1\n") || !strings.Contains(out, "    L2 --> L1\n") {
		t.Errorf("Kanten falsch aufgelöst:\n%s", out)
	}
	if strings.Contains(out, "adapters --> core") {
		t.Errorf("Roh-Name statt interner ID in einer Kante:\n%s", out)
	}
}

// TestRenderAllowDashed: allow-Kante wird gestrichelt und gelabelt gerendert.
func TestRenderAllowDashed(t *testing.T) {
	m := core.Model{
		Layers: []core.Layer{{Name: "ports"}},
		Allow:  []core.Edge{{From: "ports", To: "ports"}},
	}
	if out := render(m); !strings.Contains(out, "-.->|allow|") {
		t.Errorf("allow-Kante nicht gestrichelt/gelabelt:\n%s", out)
	}
}

// TestRenderDirectionSubgraphs: driving/driven → stabile Subgraphs; ein Layer
// ohne direction bleibt außerhalb (AC-FA-CLI-002 direction-Gruppierung).
func TestRenderDirectionSubgraphs(t *testing.T) {
	m := core.Model{Layers: []core.Layer{
		{Name: "driver", Role: "adapter", Direction: "driving"},
		{Name: "sink", Role: "adapter", Direction: "driven"},
		{Name: "core", Role: "domain"},
	}}
	out := render(m)
	if !strings.Contains(out, `subgraph dir_driving["driving"]`) ||
		!strings.Contains(out, `subgraph dir_driven["driven"]`) {
		t.Errorf("Direction-Subgraphs fehlen:\n%s", out)
	}
	if n := strings.Count(out, "    end\n"); n != 2 {
		t.Errorf("erwartet 2 subgraph-Enden, got %d:\n%s", n, out)
	}
}

// TestRenderCompositionRootIsolated: composition_root → EIN isolierter
// Notizknoten, keine All-to-All-Kanten (SPEC-CLI-002).
func TestRenderCompositionRootIsolated(t *testing.T) {
	m := core.Model{
		Layers:          []core.Layer{{Name: "core"}, {Name: "adapters"}},
		CompositionRoot: []string{"cmd/**", "internal/cli/**"},
	}
	out := render(m)
	if !strings.Contains(out, `C0["composition_root<br/>cmd/**<br/>internal/cli/**"]:::exempt`) {
		t.Errorf("composition_root-Notizknoten fehlt/falsch:\n%s", out)
	}
	if strings.Contains(out, "C0 -->") || strings.Contains(out, "--> C0") || strings.Contains(out, "C0 -.") {
		t.Errorf("composition_root darf keine Kanten haben:\n%s", out)
	}
}

// TestRenderAdapterSinkIsolated: adapter_sink → isolierter Ausnahmeknoten ohne
// Layer-Klasse und ohne Kanten (SPEC-CLI-002).
func TestRenderAdapterSinkIsolated(t *testing.T) {
	m := core.Model{Layers: []core.Layer{{Name: "adapters"}}, AdapterSink: "driver-common"}
	out := render(m)
	if !strings.Contains(out, `S0["adapter_sink<br/>driver-common"]:::exempt`) {
		t.Errorf("adapter_sink-Knoten fehlt:\n%s", out)
	}
	if strings.Contains(out, "S0 -->") || strings.Contains(out, "--> S0") {
		t.Errorf("adapter_sink darf keine Kanten haben:\n%s", out)
	}
}

// TestRenderMinimal: minimale Config ohne optionale Blöcke → keine leeren
// Sonderknoten, keine Subgraphs (AC-FA-CLI-002 Boundary).
func TestRenderMinimal(t *testing.T) {
	m := core.Model{
		Layers: []core.Layer{{Name: "core"}, {Name: "adapters"}},
		Edges:  []core.Edge{{From: "adapters", To: "core"}},
	}
	out := render(m)
	for _, unwanted := range []string{"C0[", "S0[", ":::dangling", "subgraph"} {
		if strings.Contains(out, unwanted) {
			t.Errorf("minimale Config: unerwartetes %q:\n%s", unwanted, out)
		}
	}
}

// TestRenderDangling: ein unbekannter Kanten-Endpunkt wird als Dangling-Knoten
// gerendert, kein Absturz (SPEC-CLI-002 [Medium-Dangling-Fix]).
func TestRenderDangling(t *testing.T) {
	m := core.Model{
		Layers: []core.Layer{{Name: "core"}},
		Edges:  []core.Edge{{From: "ghost", To: "core"}},
	}
	out := render(m)
	if !strings.Contains(out, `["ghost"]:::dangling`) {
		t.Errorf("Dangling-Knoten fehlt:\n%s", out)
	}
	if !strings.Contains(out, "D0 --> L0") {
		t.Errorf("Dangling-Endpunkt nicht als D-ID in der Kante:\n%s", out)
	}
}

// TestRenderEscaping: Mermaid-heikle Zeichen in Layer-Namen werden nach dem
// Escaping-Vertrag kodiert; kein roher Backtick/kein roher Name im Output
// (SPEC-CLI-002 [Medium-Fix]).
func TestRenderEscaping(t *testing.T) {
	m := core.Model{Layers: []core.Layer{
		{Name: "br[a]ck|et\\sl"},
		{Name: "lt<gt>&q\"tk`x"},
	}}
	out := render(m)
	for _, want := range []string{"&#91;", "&#93;", "&#124;", "&#92;", "&lt;", "&gt;", "&amp;", "&quot;", "&#96;"} {
		if !strings.Contains(out, want) {
			t.Errorf("Escaping fehlt %q:\n%s", want, out)
		}
	}
	if strings.Contains(out, "`") {
		t.Errorf("roher Backtick im Output (Syntax-Ausbruch möglich):\n%s", out)
	}
	if strings.Contains(out, "br[a]ck") {
		t.Errorf("roher unsafe Name im Label (nicht escaped):\n%s", out)
	}
}

// TestRenderControlChars: CR/LF/TAB werden zu Leerzeichen, übrige Steuerzeichen
// zu '?' — der Renderer bleibt einzeilig je Label (SPEC-CLI-002).
func TestRenderControlChars(t *testing.T) {
	m := core.Model{Layers: []core.Layer{{Name: "a\r\n\tb\x01c"}}}
	out := render(m)
	if strings.Contains(out, "\r") || strings.Contains(out, "\x01") {
		t.Errorf("Steuerzeichen nicht normalisiert:\n%q", out)
	}
	if !strings.Contains(out, "a   b?c") { // CR,LF,TAB → je ein Space; \x01 → ?
		t.Errorf("Steuerzeichen-Normalisierung falsch:\n%q", out)
	}
}

// TestRenderStableRegardlessOfInputOrder: gemischte Eingabe-Reihenfolge →
// byte-identische Ausgabe (Determinismus, AC-QA-01).
func TestRenderStableRegardlessOfInputOrder(t *testing.T) {
	a := core.Model{
		Layers: []core.Layer{{Name: "core"}, {Name: "adapters"}, {Name: "ports"}},
		Edges:  []core.Edge{{From: "adapters", To: "core"}, {From: "ports", To: "core"}},
	}
	b := core.Model{
		Layers: []core.Layer{{Name: "ports"}, {Name: "core"}, {Name: "adapters"}},
		Edges:  []core.Edge{{From: "ports", To: "core"}, {From: "adapters", To: "core"}},
	}
	if render(a) != render(b) {
		t.Errorf("Render nicht ordnungs-stabil:\n%s\n---\n%s", render(a), render(b))
	}
}

// TestRenderExplicitRoleWins: explizites role: hat Vorrang vor der Namens-Inferenz.
func TestRenderExplicitRoleWins(t *testing.T) {
	m := core.Model{Layers: []core.Layer{{Name: "core", Role: "adapter"}}}
	if out := render(m); !strings.Contains(out, `["core"]:::adapter`) {
		t.Errorf("explizites role: nicht angewandt:\n%s", out)
	}
}

// TestRenderUnexpectedDirectionNotDropped: ein Layer mit einer nicht erwarteten
// direction (jenseits ""/driving/driven) wird dennoch als Knoten DEFINIERT —
// sonst referenzierte eine Kante eine undefinierte Mermaid-ID (Review-Fix A1;
// der Renderer verlässt sich nicht auf die Config-Validierung der direction).
func TestRenderUnexpectedDirectionNotDropped(t *testing.T) {
	m := core.Model{
		Layers: []core.Layer{{Name: "x", Direction: "lateral"}, {Name: "core"}},
		Edges:  []core.Edge{{From: "x", To: "core"}},
	}
	out := render(m)
	if !strings.Contains(out, `["x"]`) {
		t.Errorf("Layer mit unerwarteter direction gedroppt (undefinierte Kanten-ID):\n%s", out)
	}
}

func TestRenderLegendListsCategoricalRules(t *testing.T) { // slice-040: alle kategorischen Regeln in der Legende
	m := core.Model{Layers: []core.Layer{{Name: "core", Globs: []string{"core/**"}, Role: "domain"}}}
	out := render(m)
	for _, want := range []string{"core-impurity", "lateral-adapter", "lateral-slice", "port-direction-mismatch", "port-locality"} {
		if !strings.Contains(out, want) {
			t.Errorf("Legende fehlt kategorische Regel %q:\n%s", want, out)
		}
	}
}
