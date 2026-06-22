package core

import "testing"

func testModel() Model {
	return Model{
		Layers: []Layer{
			{Name: "core", Globs: []string{"core/**"}},
			{Name: "ports", Globs: []string{"ports/**"}},
			{Name: "adapters", Globs: []string{"adapters/**"}},
		},
		Edges:           []Edge{{From: "adapters", To: "ports"}, {From: "ports", To: "core"}},
		AdapterSink:     "driver-common",
		Techs:           []Tech{{Pattern: "net/http", Adapter: "adapters/http"}},
		CompositionRoot: []string{"cmd/**"},
		Forbidden:       map[string][]string{"ports": {"impl "}},
	}
}

func hasRule(fs []Finding, rule string) bool {
	for _, f := range fs {
		if f.Rule == rule {
			return true
		}
	}
	return false
}

func TestCoreImpurity(t *testing.T) { // AC-FA-RULE-001 negative
	fs := Evaluate(testModel(), []FileImports{
		{Path: "core/svc.go", Layer: "core", Imports: []Import{{Symbol: "adapters/http", Line: 3}}},
	})
	if !hasRule(fs, "core-impurity") {
		t.Fatalf("expected core-impurity, got %v", fs)
	}
}

func TestCoreClean(t *testing.T) { // AC-FA-RULE-001 happy
	fs := Evaluate(testModel(), []FileImports{
		{Path: "core/svc.go", Layer: "core", Imports: []Import{{Symbol: "core/util", Line: 3}}},
	})
	if len(fs) != 0 {
		t.Fatalf("expected clean, got %v", fs)
	}
}

func TestLateralAdapter(t *testing.T) { // AC-FA-RULE-002 negative
	fs := Evaluate(testModel(), []FileImports{
		{Path: "adapters/a/x.go", Layer: "adapters", Imports: []Import{{Symbol: "adapters/b/y", Line: 5}}},
	})
	if !hasRule(fs, "lateral-adapter") {
		t.Fatalf("expected lateral-adapter, got %v", fs)
	}
}

func TestLateralSinkAllowed(t *testing.T) { // AC-FA-RULE-002 boundary
	fs := Evaluate(testModel(), []FileImports{
		{Path: "adapters/a/x.go", Layer: "adapters", Imports: []Import{{Symbol: "adapters/driver-common/z", Line: 5}}},
	})
	if len(fs) != 0 {
		t.Fatalf("shared sink must be allowed, got %v", fs)
	}
}

func TestTechLeak(t *testing.T) { // AC-FA-RULE-003 negative
	fs := Evaluate(testModel(), []FileImports{
		{Path: "adapters/persistence/db.go", Layer: "adapters", Imports: []Import{{Symbol: "net/http", Line: 7}}},
	})
	if !hasRule(fs, "tech-leak") {
		t.Fatalf("expected tech-leak, got %v", fs)
	}
}

func TestTechInOwnAdapter(t *testing.T) { // AC-FA-RULE-003 boundary
	fs := Evaluate(testModel(), []FileImports{
		{Path: "adapters/http/client.go", Layer: "adapters", Imports: []Import{{Symbol: "net/http", Line: 7}}},
	})
	if len(fs) != 0 {
		t.Fatalf("tech in its own adapter must be allowed, got %v", fs)
	}
}

func TestPortDomainAllowed(t *testing.T) { // AC-FA-RULE-004 happy: Ports dürfen die Domäne referenzieren
	fs := Evaluate(testModel(), []FileImports{
		{Path: "ports/p.go", Layer: "ports", Imports: []Import{{Symbol: "core/x", Line: 2}}},
	})
	if len(fs) != 0 {
		t.Fatalf("ports may reference the domain (edge ports->core declared), got %v", fs)
	}
}

func TestPortImpurityAdapter(t *testing.T) { // AC-FA-RULE-004 negative: Port importiert Adapter
	fs := Evaluate(testModel(), []FileImports{
		{Path: "ports/p.go", Layer: "ports", Imports: []Import{{Symbol: "adapters/http/client", Line: 2}}},
	})
	if !hasRule(fs, "port-impurity") {
		t.Fatalf("expected port-impurity (port imports adapter), got %v", fs)
	}
}

func TestPortImpurityTech(t *testing.T) { // AC-FA-RULE-004 negative: Port importiert Tech/Framework
	fs := Evaluate(testModel(), []FileImports{
		{Path: "ports/p.go", Layer: "ports", Imports: []Import{{Symbol: "net/http", Line: 2}}},
	})
	if !hasRule(fs, "port-impurity") {
		t.Fatalf("expected port-impurity (port imports tech), got %v", fs)
	}
}

func TestPortImpurityConstruct(t *testing.T) { // AC-FA-RULE-004 negative (construct)
	fs := Evaluate(testModel(), []FileImports{
		{Path: "ports/p.go", Layer: "ports", Constructs: []Import{{Symbol: "impl ", Line: 4}}},
	})
	if !hasRule(fs, "port-impurity") {
		t.Fatalf("expected port-impurity from forbidden construct, got %v", fs)
	}
}

func TestPortToCoreWithoutEdge(t *testing.T) { // AC-FA-RULE-004/005: ports->core ist edge-regiert, nicht port-impurity
	m := testModel()
	m.Edges = []Edge{{From: "adapters", To: "ports"}} // {ports->core}-Kante entfernt
	fs := Evaluate(m, []FileImports{
		{Path: "ports/p.go", Layer: "ports", Imports: []Import{{Symbol: "core/x", Line: 2}}},
	})
	if hasRule(fs, "port-impurity") {
		t.Fatalf("ports->core darf NIE port-impurity sein (Kern-Referenz erlaubt), got %v", fs)
	}
	if !hasRule(fs, "wrong-direction") {
		t.Fatalf("ports->core ohne deklarierte Kante muss wrong-direction sein, got %v", fs)
	}
}

func TestWrongDirection(t *testing.T) { // AC-FA-RULE-005 negative
	fs := Evaluate(testModel(), []FileImports{
		{Path: "adapters/a/x.go", Layer: "adapters", Imports: []Import{{Symbol: "core/x", Line: 9}}},
	})
	if !hasRule(fs, "wrong-direction") {
		t.Fatalf("expected wrong-direction (adapters->core not in edges), got %v", fs)
	}
}

func TestCompositionRootExempt(t *testing.T) { // composition root wires everything
	fs := Evaluate(testModel(), []FileImports{
		{Path: "cmd/main.go", Layer: "", Imports: []Import{{Symbol: "adapters/http", Line: 1}, {Symbol: "core/x", Line: 2}}},
	})
	if len(fs) != 0 {
		t.Fatalf("composition root must be exempt, got %v", fs)
	}
}

func TestGlobAndLayerHelpers(t *testing.T) {
	layers := []Layer{
		{Name: "core", Globs: []string{"core/**"}},
		{Name: "x", Globs: []string{"a/*/b.go"}},
	}
	if LayerOf("core/deep/f.go", layers) != "core" {
		t.Fatal("LayerOf ** failed")
	}
	if LayerOf("nope/x.go", layers) != "" {
		t.Fatal("LayerOf none failed")
	}
	if !MatchGlobs("a/z/b.go", []string{"a/*/b.go"}) {
		t.Fatal("MatchGlobs * failed")
	}
	if MatchGlobs("a/zz/b.go", []string{"a/?/b.go"}) {
		t.Fatal("MatchGlobs ? must not match two chars")
	}
}

func TestAllowEdge(t *testing.T) { // edgeAllowed via Allow-Liste
	m := testModel()
	m.Allow = []Edge{{From: "adapters", To: "core"}}
	fs := Evaluate(m, []FileImports{
		{Path: "adapters/a/x.go", Layer: "adapters", Imports: []Import{{Symbol: "core/x", Line: 1}}},
	})
	if hasRule(fs, "wrong-direction") {
		t.Fatalf("Allow should permit adapters->core: %v", fs)
	}
}

func TestDeterministicOrder(t *testing.T) { // AC-QA-01: stable sort by path, line, rule
	files := []FileImports{
		{Path: "core/b.go", Layer: "core", Imports: []Import{{Symbol: "adapters/x", Line: 2}}},
		{Path: "core/a.go", Layer: "core", Imports: []Import{{Symbol: "adapters/x", Line: 9}}},
	}
	fs := Evaluate(testModel(), files)
	if len(fs) != 2 || fs[0].Path != "core/a.go" || fs[1].Path != "core/b.go" {
		t.Fatalf("findings not stably sorted: %v", fs)
	}
}

// roleModel uses FOREIGN layer names with explicit roles (AC-FA-RULE-006) to
// prove the dispatch is name-independent.
func roleModel() Model {
	return Model{
		Layers: []Layer{
			{Name: "domainx", Globs: []string{"domainx/**"}, Role: "domain"},
			{Name: "geometry", Globs: []string{"geometry/**"}, Role: "adapter"},
			{Name: "persistence", Globs: []string{"persistence/**"}, Role: "adapter"},
		},
		Edges: []Edge{
			{From: "geometry", To: "domainx"},
			{From: "persistence", To: "domainx"},
		},
		AdapterSink: "driver-common",
	}
}

func TestRoleCrossLayerLateral(t *testing.T) { // AC-FA-RULE-006 happy: fremde Namen, role:adapter -> role:adapter
	fs := Evaluate(roleModel(), []FileImports{
		{Path: "geometry/g.go", Layer: "geometry", Imports: []Import{{Symbol: "persistence/p", Line: 3}}},
	})
	if !hasRule(fs, "lateral-adapter") {
		t.Fatalf("expected lateral-adapter across different role:adapter layers, got %v", fs)
	}
}

func TestRoleCrossLayerLateralCategorical(t *testing.T) { // AC-FA-RULE-006: kategorisch — weder allow NOCH edges heben auf
	// Differenzial: dieselbe Kante als allow UND als edge darf lateral nicht
	// unterdrücken. Wäre lateral edge-regiert, fiele der Import wegen der
	// erlaubten Kante auf KEINEN Befund (len 0) -> der len==1-Assert würde rot.
	for _, tc := range []struct {
		name string
		mod  func(*Model)
	}{
		{"allow", func(m *Model) { m.Allow = []Edge{{From: "geometry", To: "persistence"}} }},
		{"edge", func(m *Model) { m.Edges = append(m.Edges, Edge{From: "geometry", To: "persistence"}) }},
	} {
		t.Run(tc.name, func(t *testing.T) {
			m := roleModel()
			tc.mod(&m)
			fs := Evaluate(m, []FileImports{
				{Path: "geometry/g.go", Layer: "geometry", Imports: []Import{{Symbol: "persistence/p", Line: 3}}},
			})
			if len(fs) != 1 || fs[0].Rule != "lateral-adapter" {
				t.Fatalf("lateral ist kategorisch (%s darf nicht aufheben; genau ein lateral-adapter erwartet), got %v", tc.name, fs)
			}
		})
	}
}

func TestRoleDomainImportsAdapter(t *testing.T) { // AC-FA-RULE-006 negative (a): role:domain -> role:adapter
	fs := Evaluate(roleModel(), []FileImports{
		{Path: "domainx/d.go", Layer: "domainx", Imports: []Import{{Symbol: "geometry/g", Line: 2}}},
	})
	if !hasRule(fs, "core-impurity") {
		t.Fatalf("expected core-impurity (role:domain imports role:adapter), got %v", fs)
	}
}

func TestRolePortConstructForeignName(t *testing.T) { // AC-FA-RULE-006 negative (b): role:port, fremder Name, Konstrukt
	m := Model{Layers: []Layer{{Name: "api", Globs: []string{"api/**"}, Role: "port"}}}
	fs := Evaluate(m, []FileImports{
		{Path: "api/p.go", Layer: "api", Constructs: []Import{{Symbol: "impl ", Line: 4}}},
	})
	if !hasRule(fs, "port-impurity") {
		t.Fatalf("expected port-impurity (role:port construct, foreign name), got %v", fs)
	}
}

func TestRolelessLayerEdgeOnly(t *testing.T) { // AC-FA-RULE-006: ohne Rolle nur kanten-geprüft
	m := Model{Layers: []Layer{
		{Name: "alpha", Globs: []string{"alpha/**"}},
		{Name: "beta", Globs: []string{"beta/**"}},
	}}
	fs := Evaluate(m, []FileImports{
		{Path: "alpha/a.go", Layer: "alpha", Imports: []Import{{Symbol: "beta/b", Line: 1}}},
	})
	if hasRule(fs, "core-impurity") || hasRule(fs, "port-impurity") || hasRule(fs, "lateral-adapter") {
		t.Fatalf("roleless layers must not trigger purity rules, got %v", fs)
	}
	if !hasRule(fs, "wrong-direction") {
		t.Fatalf("roleless cross-layer import without edge must be wrong-direction, got %v", fs)
	}
}

func TestExplicitRoleOverridesInference(t *testing.T) { // AC-FA-RULE-006: explizite role: schlägt Namens-Inferenz
	m := Model{
		Layers: []Layer{
			{Name: "core", Globs: []string{"core/**"}, Role: "adapter"}, // Name core, aber explizit adapter
			{Name: "other", Globs: []string{"other/**"}, Role: "adapter"},
		},
		Edges: []Edge{{From: "core", To: "other"}},
	}
	fs := Evaluate(m, []FileImports{
		{Path: "core/c.go", Layer: "core", Imports: []Import{{Symbol: "other/o", Line: 1}}},
	})
	if !hasRule(fs, "lateral-adapter") {
		t.Fatalf("explicit role:adapter must override name inference (core->domain), got %v", fs)
	}
}

func TestInferenceBoundaryClassicNames(t *testing.T) { // AC-FA-RULE-006 boundary: klassische Namen OHNE role == Verhalten 0.2.0
	m := Model{ // wie 0.2.0: Namen core/ports/adapters, KEINE Role -> Inferenz
		Layers: []Layer{
			{Name: "core", Globs: []string{"core/**"}},
			{Name: "ports", Globs: []string{"ports/**"}},
			{Name: "adapters", Globs: []string{"adapters/**"}},
		},
		Edges: []Edge{{From: "adapters", To: "ports"}, {From: "ports", To: "core"}},
	}
	coreToAdapter := Evaluate(m, []FileImports{
		{Path: "core/s.go", Layer: "core", Imports: []Import{{Symbol: "adapters/x", Line: 1}}},
	})
	if !hasRule(coreToAdapter, "core-impurity") {
		t.Fatalf("core->adapter via inference must be core-impurity, got %v", coreToAdapter)
	}
	portConstruct := Evaluate(m, []FileImports{
		{Path: "ports/p.go", Layer: "ports", Constructs: []Import{{Symbol: "impl ", Line: 2}}},
	})
	if !hasRule(portConstruct, "port-impurity") {
		t.Fatalf("ports construct via inference must be port-impurity, got %v", portConstruct)
	}
	intraLateral := Evaluate(m, []FileImports{
		{Path: "adapters/a/x.go", Layer: "adapters", Imports: []Import{{Symbol: "adapters/b/y", Line: 3}}},
	})
	if !hasRule(intraLateral, "lateral-adapter") {
		t.Fatalf("intra-adapters lateral via inference must be lateral-adapter, got %v", intraLateral)
	}
}

func TestForeignAdapterIntraLateral(t *testing.T) { // ADR-0010 (b1): adapterSeg layer-relativ -> Intra-Unterscheidung fremder Namen
	// Ein einziger role:adapter-Layer mit fremdem Namen (kein "adapters"-Segment):
	// die Sub-Einheit wird relativ zum Glob-Präfix der Schicht bestimmt, also greift
	// die Intra-Unterscheidung jetzt namensunabhängig (kehrt den 10a-Regression-Pin um).
	m := Model{Layers: []Layer{{Name: "io", Globs: []string{"io/**"}, Role: "adapter"}}}
	fs := Evaluate(m, []FileImports{
		{Path: "io/a/x.go", Layer: "io", Imports: []Import{{Symbol: "io/b/y", Line: 1}}},
	})
	if !hasRule(fs, "lateral-adapter") {
		t.Fatalf("b1: fremd benannte Intra-Adapter-Sub-Einheiten müssen lateral-adapter sein, got %v", fs)
	}
}

func TestTargetLayerLongestPrefix(t *testing.T) { // ADR-0010 (b1): spezifischster (längster) Glob-Präfix gewinnt
	layers := []Layer{
		{Name: "core", Globs: []string{"internal/core/**"}},
		{Name: "legacy", Globs: []string{"internal/core/legacy/**"}},
	}
	if got := targetLayer("x/internal/core/legacy/db", layers); got != "legacy" {
		t.Fatalf("expected longest-prefix 'legacy', got %q", got)
	}
	if got := targetLayer("x/internal/core/svc", layers); got != "core" {
		t.Fatalf("expected 'core', got %q", got)
	}
	// Reihenfolge-unabhängig: legacy vor core deklariert.
	rev := []Layer{
		{Name: "legacy", Globs: []string{"internal/core/legacy/**"}},
		{Name: "core", Globs: []string{"internal/core/**"}},
	}
	if got := targetLayer("x/internal/core/legacy/db", rev); got != "legacy" {
		t.Fatalf("longest-prefix muss reihenfolge-unabhängig sein, got %q", got)
	}
	// Segment-bewusst: 'io'-Präfix matcht nicht in 'audio'.
	if got := targetLayer("audio/codec", []Layer{{Name: "io", Globs: []string{"io/**"}}}); got != "" {
		t.Fatalf("segment-bewusst: 'io' darf nicht in 'audio' matchen, got %q", got)
	}
	// Kernzweck: modul-qualifizierter Import, Präfix mitten im String.
	mod := []Layer{{Name: "core", Globs: []string{"internal/hexagon/core/**"}}}
	if got := targetLayer("github.com/x/a-check/internal/hexagon/core/model", mod); got != "core" {
		t.Fatalf("modul-qualifiziert: erwarte 'core', got %q", got)
	}
	// Präfix am Pfadende.
	if got := targetLayer("github.com/x/a-check/internal/hexagon/core", mod); got != "core" {
		t.Fatalf("Präfix am Pfadende: erwarte 'core', got %q", got)
	}
	// Tie-Break: bei gleichlangem Präfix gewinnt der zuerst deklarierte Layer.
	tie := []Layer{
		{Name: "first", Globs: []string{"a/b/**"}},
		{Name: "second", Globs: []string{"a/b/**"}},
	}
	if got := targetLayer("a/b/c", tie); got != "first" {
		t.Fatalf("Tie-Break: zuerst deklarierter gewinnt, erwarte 'first', got %q", got)
	}
}

func TestSameAdapterSubunitNoLateral(t *testing.T) { // ADR-0010 (b1): gleiche Sub-Einheit -> kein lateral
	m := Model{Layers: []Layer{{Name: "io", Globs: []string{"io/**"}, Role: "adapter"}}}
	fs := Evaluate(m, []FileImports{
		{Path: "io/a/x.go", Layer: "io", Imports: []Import{{Symbol: "io/a/z", Line: 1}}},
	})
	if hasRule(fs, "lateral-adapter") {
		t.Fatalf("gleiche Sub-Einheit darf kein lateral-adapter sein, got %v", fs)
	}
}
