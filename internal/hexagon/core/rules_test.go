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
