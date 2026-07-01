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

// assertCategorical proves a rule is edge-independent: neither an allow nor an
// edge for the (from,to) pair may suppress it — exactly one finding of wantRule
// must remain. Wäre die Regel edge-regiert, fiele der Import mit der erlaubten
// Kante auf KEINEN Befund und der len==1-Assert würde rot.
func assertCategorical(t *testing.T, base func() Model, from, to string, file FileImports, wantRule string) {
	t.Helper()
	for _, tc := range []struct {
		name string
		mod  func(*Model)
	}{
		{"allow", func(m *Model) { m.Allow = []Edge{{From: from, To: to}} }},
		{"edge", func(m *Model) { m.Edges = append(m.Edges, Edge{From: from, To: to}) }},
	} {
		t.Run(tc.name, func(t *testing.T) {
			m := base()
			tc.mod(&m)
			fs := Evaluate(m, []FileImports{file})
			if len(fs) != 1 || fs[0].Rule != wantRule {
				t.Fatalf("%s ist kategorisch (%s darf nicht aufheben), got %v", wantRule, tc.name, fs)
			}
		})
	}
}

func TestRoleCrossLayerLateralCategorical(t *testing.T) { // AC-FA-RULE-006: kategorisch — weder allow NOCH edges heben auf
	assertCategorical(t, roleModel, "geometry", "persistence",
		FileImports{Path: "geometry/g.go", Layer: "geometry", Imports: []Import{{Symbol: "persistence/p", Line: 3}}},
		"lateral-adapter")
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

// --- AC-FA-RULE-007 / ADR-0011: Rolle app + strenge domain (welle-10b/b2a) ---

// appModel: domain/app/port/adapter mit FREMDEN Namen + expliziten Rollen und den
// Kanten app->dom, app->prt, prt->dom (namensunabhängig).
func appModel() Model {
	return Model{
		Layers: []Layer{
			{Name: "dom", Globs: []string{"dom/**"}, Role: "domain"},
			{Name: "app", Globs: []string{"app/**"}, Role: "app"},
			{Name: "prt", Globs: []string{"prt/**"}, Role: "port"},
			{Name: "adp", Globs: []string{"adp/**"}, Role: "adapter"},
		},
		Edges: []Edge{
			{From: "app", To: "dom"},
			{From: "app", To: "prt"},
			{From: "prt", To: "dom"},
		},
		Techs: []Tech{{Pattern: "net/http", Adapter: "adp"}},
	}
}

func TestAppHappyDomainAndPort(t *testing.T) { // AC-FA-RULE-007 happy: app darf domain+port
	fs := Evaluate(appModel(), []FileImports{
		{Path: "app/u.go", Layer: "app", Imports: []Import{
			{Symbol: "dom/entity", Line: 1},
			{Symbol: "prt/repo", Line: 2},
		}},
	})
	if len(fs) != 0 {
		t.Fatalf("app darf domain+port importieren (kein Befund), got %v", fs)
	}
}

func TestAppImportsAdapterCategorical(t *testing.T) { // AC-FA-RULE-007 negative (app): app->adapter, kategorisch
	m := appModel()
	m.Edges = append(m.Edges, Edge{From: "app", To: "adp"}) // sogar mit erlaubter Kante
	fs := Evaluate(m, []FileImports{
		{Path: "app/u.go", Layer: "app", Imports: []Import{{Symbol: "adp/sql", Line: 3}}},
	})
	if len(fs) != 1 || fs[0].Rule != "app-impurity" {
		t.Fatalf("app->adapter ist app-impurity, kategorisch (genau ein Befund trotz Kante), got %v", fs)
	}
}

func TestAppImportsTech(t *testing.T) { // AC-FA-RULE-007 negative (app): zweiter Arm app->tech (vor tech-leak)
	fs := Evaluate(appModel(), []FileImports{
		{Path: "app/u.go", Layer: "app", Imports: []Import{{Symbol: "net/http", Line: 4}}},
	})
	if len(fs) != 1 || fs[0].Rule != "app-impurity" {
		t.Fatalf("app->tech ist app-impurity (genau ein Befund), got %v", fs)
	}
}

func TestDomainImportsPortCategorical(t *testing.T) { // AC-FA-RULE-007 negative (domain): domain->port, kategorisch
	m := appModel()
	m.Edges = append(m.Edges, Edge{From: "dom", To: "prt"}) // Kante hebt nicht auf
	fs := Evaluate(m, []FileImports{
		{Path: "dom/e.go", Layer: "dom", Imports: []Import{{Symbol: "prt/repo", Line: 5}}},
	})
	if len(fs) != 1 || fs[0].Rule != "core-impurity" {
		t.Fatalf("domain->port ist core-impurity, kategorisch (genau ein Befund trotz Kante), got %v", fs)
	}
}

func TestDomainImportsAppCategorical(t *testing.T) { // AC-FA-RULE-007: Schärfung deckt app, nicht nur port
	m := appModel()
	m.Edges = append(m.Edges, Edge{From: "dom", To: "app"})
	fs := Evaluate(m, []FileImports{
		{Path: "dom/e.go", Layer: "dom", Imports: []Import{{Symbol: "app/u", Line: 6}}},
	})
	if len(fs) != 1 || fs[0].Rule != "core-impurity" {
		t.Fatalf("domain->app ist core-impurity, kategorisch, got %v", fs)
	}
}

func TestInferAppRole(t *testing.T) { // AC-FA-RULE-007: Namens-Inferenz application/app -> app
	for _, name := range []string{"application", "app"} {
		t.Run(name, func(t *testing.T) {
			m := Model{Layers: []Layer{
				{Name: name, Globs: []string{name + "/**"}}, // KEINE role -> Inferenz
				{Name: "adp", Globs: []string{"adp/**"}, Role: "adapter"},
			}}
			fs := Evaluate(m, []FileImports{
				{Path: name + "/u.go", Layer: name, Imports: []Import{{Symbol: "adp/x", Line: 1}}},
			})
			if !hasRule(fs, "app-impurity") {
				t.Fatalf("%q sollte zu role app inferieren (app-impurity erwartet), got %v", name, fs)
			}
		})
	}
}

func TestExplicitRoleBeatsAppInference(t *testing.T) { // AC-FA-RULE-007: explizite role: schlägt Inferenz
	m := Model{Layers: []Layer{
		{Name: "app", Globs: []string{"app/**"}, Role: "domain"}, // Name app, aber role domain
		{Name: "prt", Globs: []string{"prt/**"}, Role: "port"},
	}}
	fs := Evaluate(m, []FileImports{
		{Path: "app/d.go", Layer: "app", Imports: []Import{{Symbol: "prt/p", Line: 1}}},
	})
	if !hasRule(fs, "core-impurity") || hasRule(fs, "app-impurity") {
		t.Fatalf("role: domain schlägt app-Inferenz -> core-impurity (kein app-impurity), got %v", fs)
	}
}

func TestDomainImportsAdapterCategorical(t *testing.T) { // AC-FA-RULE-007: domain->adapter kategorisch (Pin mit Kante)
	m := appModel()
	m.Edges = append(m.Edges, Edge{From: "dom", To: "adp"}) // Kante hebt nicht auf
	fs := Evaluate(m, []FileImports{
		{Path: "dom/e.go", Layer: "dom", Imports: []Import{{Symbol: "adp/sql", Line: 7}}},
	})
	if len(fs) != 1 || fs[0].Rule != "core-impurity" {
		t.Fatalf("domain->adapter ist core-impurity, kategorisch (genau ein Befund trotz Kante), got %v", fs)
	}
}

func TestDomainImportsTech(t *testing.T) { // AC-FA-RULE-007: domain->tech ist core-impurity (vor tech-leak)
	fs := Evaluate(appModel(), []FileImports{
		{Path: "dom/e.go", Layer: "dom", Imports: []Import{{Symbol: "net/http", Line: 8}}},
	})
	if len(fs) != 1 || fs[0].Rule != "core-impurity" {
		t.Fatalf("domain->tech ist core-impurity (genau ein Befund, vor tech-leak), got %v", fs)
	}
}

// --- AC-FA-RULE-008 / ADR-0012: Driving/Driven-Port-Richtung (welle-10b/b2b) ---

// dirModel: getrennte driving/driven Adapter- und Port-Schichten mit expliziten
// Richtungen. Die Kante cli->api erlaubt den happy-Fall, ohne wrong-direction.
func dirModel() Model {
	return Model{
		Layers: []Layer{
			{Name: "cli", Globs: []string{"cli/**"}, Role: "adapter", Direction: "driving"},
			{Name: "api", Globs: []string{"api/**"}, Role: "port", Direction: "driving"},
			{Name: "store", Globs: []string{"store/**"}, Role: "port", Direction: "driven"},
		},
		Edges: []Edge{{From: "cli", To: "api"}},
	}
}

func TestPortDirectionHappy(t *testing.T) { // AC-FA-RULE-008 happy: driving-Adapter -> driving-Port
	fs := Evaluate(dirModel(), []FileImports{
		{Path: "cli/c.go", Layer: "cli", Imports: []Import{{Symbol: "api/u", Line: 1}}},
	})
	if len(fs) != 0 {
		t.Fatalf("driving-Adapter -> driving-Port: kein Befund erwartet, got %v", fs)
	}
}

func TestPortDirectionMismatch(t *testing.T) { // AC-FA-RULE-008 negative: driving-Adapter -> driven-Port
	fs := Evaluate(dirModel(), []FileImports{
		{Path: "cli/c.go", Layer: "cli", Imports: []Import{{Symbol: "store/s", Line: 2}}},
	})
	if len(fs) != 1 || fs[0].Rule != "port-direction-mismatch" {
		t.Fatalf("driving-Adapter -> driven-Port: genau ein port-direction-mismatch erwartet, got %v", fs)
	}
}

func TestPortDirectionMismatchCategorical(t *testing.T) { // AC-FA-RULE-008: kategorisch — edges/allow heben nicht auf
	assertCategorical(t, dirModel, "cli", "store",
		FileImports{Path: "cli/c.go", Layer: "cli", Imports: []Import{{Symbol: "store/s", Line: 2}}},
		"port-direction-mismatch")
}

func TestPortDirectionOnlyOneSide(t *testing.T) { // AC-FA-RULE-008: nur EINE Seite mit direction -> keine Prüfung
	m := dirModel()
	for i := range m.Layers { // store verliert die Richtung -> tgt ohne direction
		if m.Layers[i].Name == "store" {
			m.Layers[i].Direction = ""
		}
	}
	m.Edges = append(m.Edges, Edge{From: "cli", To: "store"}) // Kante: auch wrong-direction schweigt
	fs := Evaluate(m, []FileImports{
		{Path: "cli/c.go", Layer: "cli", Imports: []Import{{Symbol: "store/s", Line: 2}}},
	})
	if hasRule(fs, "port-direction-mismatch") {
		t.Fatalf("nur eine Seite mit direction -> kein port-direction-mismatch, got %v", fs)
	}
}

func TestPortDirectionBoundaryNoDirection(t *testing.T) { // AC-FA-RULE-008 boundary: ohne direction == Verhalten 0.5.0
	m := Model{ // klassische role:adapter/port OHNE direction, Kante erlaubt
		Layers: []Layer{
			{Name: "adp", Globs: []string{"adp/**"}, Role: "adapter"},
			{Name: "prt", Globs: []string{"prt/**"}, Role: "port"},
		},
		Edges: []Edge{{From: "adp", To: "prt"}},
	}
	fs := Evaluate(m, []FileImports{
		{Path: "adp/a.go", Layer: "adp", Imports: []Import{{Symbol: "prt/p", Line: 1}}},
	})
	if len(fs) != 0 {
		t.Fatalf("ohne direction: adapter->port mit Kante ist kein Befund (wie 0.5.0), got %v", fs)
	}
}

// --- ADR-0013: LayerOf längster-Präfix (Angleichung an targetLayer) ---

func TestLayerOfLongestPrefix(t *testing.T) { // ADR-0013: spezifischster (längster) Glob-Präfix gewinnt
	layers := []Layer{
		{Name: "all", Globs: []string{"src/**"}},
		{Name: "app", Globs: []string{"src/app/**"}},
	}
	if got := LayerOf("src/app/u.go", layers); got != "app" {
		t.Fatalf("verschachtelt: erwarte spezifischste Schicht 'app', got %q", got)
	}
	if got := LayerOf("src/other/x.go", layers); got != "all" {
		t.Fatalf("nur src/** matcht -> 'all', got %q", got)
	}
	rev := []Layer{ // reihenfolge-unabhängig: app vor all
		{Name: "app", Globs: []string{"src/app/**"}},
		{Name: "all", Globs: []string{"src/**"}},
	}
	if got := LayerOf("src/app/u.go", rev); got != "app" {
		t.Fatalf("längster-Präfix muss reihenfolge-unabhängig sein, got %q", got)
	}
}

func TestLayerOfTieBreakFirstDeclared(t *testing.T) { // ADR-0013: Gleichstand -> zuerst deklariert
	layers := []Layer{
		{Name: "first", Globs: []string{"a/b/**"}},
		{Name: "second", Globs: []string{"a/b/**"}},
	}
	if got := LayerOf("a/b/c.go", layers); got != "first" {
		t.Fatalf("Tie-Break: zuerst deklarierter gewinnt, erwarte 'first', got %q", got)
	}
}

func TestLayerOfMultiGlobLayer(t *testing.T) { // ADR-0013: Auswahl über den längsten MATCHENDEN Glob je Layer
	// "broad" hat zwei Globs; für src/app/special/f.go matcht der längere
	// (src/app/special) -> "broad" schlägt das zuerst deklarierte "app" — die
	// per-Glob-Spezifität spiegelt targetLayers Glob-Schleife (nicht per-Layer).
	layers := []Layer{
		{Name: "app", Globs: []string{"src/app/**"}},
		{Name: "broad", Globs: []string{"src/**", "src/app/special/**"}},
	}
	if got := LayerOf("src/app/special/f.go", layers); got != "broad" {
		t.Fatalf("Mehr-Glob: längster matchender Glob (src/app/special) gewinnt -> 'broad', got %q", got)
	}
	if got := LayerOf("src/app/other/f.go", layers); got != "app" { // hier matcht in broad nur src/**
		t.Fatalf("nur src/** matcht in broad -> 'app' (src/app) spezifischer, got %q", got)
	}
}

func TestLayerOfLiteralBeatsWildcardPrefix(t *testing.T) { // ADR-0013: literale Segment-Tiefe schlägt Wildcard-Präfix (litPrefixLen)
	// src/*/handlers (litPräfix "src", Tiefe 3) darf den literalen src/app (7)
	// NICHT überstimmen — sonst mäße matchSpecificity rohe Stringlänge (14 > 7).
	layers := []Layer{
		{Name: "wild", Globs: []string{"src/*/handlers/**"}},
		{Name: "app", Globs: []string{"src/app/**"}},
	}
	if got := LayerOf("src/app/handlers/x.go", layers); got != "app" {
		t.Fatalf("literaler Präfix (src/app) muss Wildcard-Präfix (src/*/handlers) schlagen, got %q", got)
	}
	// **/foo hat Spezifität 0 und verliert gegen jeden literalen Präfix.
	star := []Layer{
		{Name: "anyfoo", Globs: []string{"**/foo/**"}},
		{Name: "src", Globs: []string{"src/**"}},
	}
	if got := LayerOf("src/foo/x.go", star); got != "src" {
		t.Fatalf("**/foo (Spezifität 0) darf src/** (3) nicht schlagen, got %q", got)
	}
}

func TestPortDirectionPortToPortNotCaught(t *testing.T) { // AC-FA-RULE-008: port->port (Gegenrichtung) ist OUT-OF-SCOPE (Rollen-Guard)
	// api(driving) und store(driven) sind BEIDE role:port. Ein port->port-Import
	// mit Gegenrichtung erfüllt zwar directionMismatch, aber der Rollen-Guard
	// (srcRole==adapter ∧ tgtRole==port) schließt ihn aus -> kein Befund.
	m := dirModel()
	m.Edges = append(m.Edges, Edge{From: "api", To: "store"}) // damit auch wrong-direction schweigt
	fs := Evaluate(m, []FileImports{
		{Path: "api/p.go", Layer: "api", Imports: []Import{{Symbol: "store/s", Line: 1}}},
	})
	if hasRule(fs, "port-direction-mismatch") {
		t.Fatalf("port->port ist out-of-scope (Rollen-Guard) -> kein port-direction-mismatch, got %v", fs)
	}
}

func TestPortDirectionSymmetric(t *testing.T) { // AC-FA-RULE-008: driven-Adapter -> driving-Port feuert spiegelbildlich
	m := dirModel()
	m.Layers = append(m.Layers, Layer{Name: "repo", Globs: []string{"repo/**"}, Role: "adapter", Direction: "driven"})
	fs := Evaluate(m, []FileImports{
		{Path: "repo/r.go", Layer: "repo", Imports: []Import{{Symbol: "api/u", Line: 1}}}, // driven-Adapter -> driving-Port
	})
	if len(fs) != 1 || fs[0].Rule != "port-direction-mismatch" {
		t.Fatalf("driven-Adapter -> driving-Port: symmetrisch genau ein port-direction-mismatch, got %v", fs)
	}
}

func TestPortDirectionSourceNoDirection(t *testing.T) { // AC-FA-RULE-008: nur das ZIEL trägt direction (sd=="") -> keine Prüfung
	m := dirModel()
	for i := range m.Layers { // cli(driving) verliert die Richtung; store(driven) behält
		if m.Layers[i].Name == "cli" {
			m.Layers[i].Direction = ""
		}
	}
	m.Edges = append(m.Edges, Edge{From: "cli", To: "store"}) // Kante: auch wrong-direction schweigt
	fs := Evaluate(m, []FileImports{
		{Path: "cli/c.go", Layer: "cli", Imports: []Import{{Symbol: "store/s", Line: 2}}},
	})
	if hasRule(fs, "port-direction-mismatch") {
		t.Fatalf("Quelle ohne direction (sd==\"\") -> kein port-direction-mismatch, got %v", fs)
	}
}

func TestPortDirectionTechLeakPrecedence(t *testing.T) { // AC-FA-RULE-008 / SPEC-RULE-001: tech-leak steht VOR port-direction-mismatch
	// Ein driving-Adapter importiert einen driven-Port, dessen Pfad zufällig ein
	// tech-Muster trägt und außerhalb des Tech-Adapters liegt: die dokumentierte
	// Erst-Treffer-Kette meldet bewusst tech-leak (nicht port-direction-mismatch).
	m := dirModel()
	m.Techs = []Tech{{Pattern: "store/grpc", Adapter: "store/grpc-adapter"}}
	fs := Evaluate(m, []FileImports{
		{Path: "cli/c.go", Layer: "cli", Imports: []Import{{Symbol: "store/grpc/client", Line: 1}}},
	})
	if len(fs) != 1 || fs[0].Rule != "tech-leak" {
		t.Fatalf("tech-leak hat Präzedenz vor port-direction-mismatch (SPEC-RULE-001-Kette), got %v", fs)
	}
}

func TestDeterministicOrderWithDirection(t *testing.T) { // AC-QA-01: der neue Befund fügt sich in die stabile Sortierung
	files := []FileImports{
		{Path: "cli/b.go", Layer: "cli", Imports: []Import{{Symbol: "store/s", Line: 2}}},
		{Path: "cli/a.go", Layer: "cli", Imports: []Import{{Symbol: "store/s", Line: 9}}},
	}
	fs := Evaluate(dirModel(), files)
	if len(fs) != 2 || fs[0].Path != "cli/a.go" || fs[1].Path != "cli/b.go" {
		t.Fatalf("port-direction-mismatch-Befunde nicht stabil nach Pfad sortiert: %v", fs)
	}
}

// regexTechModel: zwei Adapter (ui/geometry) + ein Qt-Muster als RE2-Regex auf ui.
func regexTechModel(t *testing.T) Model {
	t.Helper()
	qt, err := NewTech("Q[A-Za-z]", "adapters/ui", "regex")
	if err != nil {
		t.Fatal(err)
	}
	return Model{
		Layers: []Layer{
			{Name: "ui", Globs: []string{"adapters/ui/**"}, Role: "adapter"},
			{Name: "geo", Globs: []string{"adapters/geometry/**"}, Role: "adapter"},
		},
		Techs:           []Tech{qt},
		CompositionRoot: []string{"main.cpp"},
	}
}

func TestTechLeakRegexNegative(t *testing.T) { // AC-FA-RULE-003 / ADR-0015: match: regex meldet außerhalb des Adapters
	fs := Evaluate(regexTechModel(t), []FileImports{
		{Path: "adapters/geometry/g.cpp", Layer: "geo", Imports: []Import{{Symbol: "QWidget", Line: 3}}},
	})
	if len(fs) != 1 || fs[0].Rule != "tech-leak" {
		t.Fatalf("regex Q[A-Za-z] soll QWidget außerhalb adapters/ui als tech-leak melden, got %v", fs)
	}
}

func TestTechLeakRegexHappy(t *testing.T) { // AC-FA-RULE-003 / ADR-0015: Qt im eigenen Adapter erlaubt
	fs := Evaluate(regexTechModel(t), []FileImports{
		{Path: "adapters/ui/w.cpp", Layer: "ui", Imports: []Import{{Symbol: "QString", Line: 3}}},
	})
	if len(fs) != 0 {
		t.Fatalf("Qt im eigenen ui-Adapter erlaubt, got %v", fs)
	}
}

func TestTechLeakRegexComposition(t *testing.T) { // AC-FA-RULE-003 / ADR-0015: Composition Root ausgenommen
	fs := Evaluate(regexTechModel(t), []FileImports{
		{Path: "main.cpp", Layer: "", Imports: []Import{{Symbol: "QApplication", Line: 1}}},
	})
	if len(fs) != 0 {
		t.Fatalf("Qt in der Composition Root ausgenommen, got %v", fs)
	}
}

func TestTechPrecedenceDeclarationOrder(t *testing.T) { // AC-FA-RULE-003 / ADR-0015: Erst-Treffer in Deklarationsreihenfolge
	first, err := NewTech("Q[A-Za-z]", "adapters/ui", "regex")
	if err != nil {
		t.Fatal(err)
	}
	second, err := NewTech("Queue", "adapters/persistence", "substring")
	if err != nil {
		t.Fatal(err)
	}
	m := Model{Layers: []Layer{{Name: "ui", Globs: []string{"adapters/ui/**"}, Role: "adapter"}}}
	file := []FileImports{{Path: "adapters/ui/x.cpp", Layer: "ui", Imports: []Import{{Symbol: "Queue.h", Line: 1}}}}

	// "Queue.h" trifft beide Muster; der erste (regex → adapters/ui) gewinnt → kein Befund
	// (die Datei liegt in adapters/ui). Griffe der zweite, wäre es ein tech-leak.
	m.Techs = []Tech{first, second}
	if fs := Evaluate(m, file); len(fs) != 0 {
		t.Fatalf("Erst-Treffer (regex → adapters/ui) muss gewinnen → kein Befund, got %v", fs)
	}
	// Umgekehrte Reihenfolge: der substring-Eintrag (→ persistence) greift zuerst → tech-leak.
	m.Techs = []Tech{second, first}
	if fs := Evaluate(m, file); len(fs) != 1 || fs[0].Rule != "tech-leak" {
		t.Fatalf("bei umgekehrter Reihenfolge greift substring→persistence → tech-leak, got %v", fs)
	}
}

func TestNewTechBackCompatSubstring(t *testing.T) { // AC-FA-RULE-003 / ADR-0015: NewTech(p,a,"") verhält sich wie das Literal Tech (Substring)
	built, err := NewTech("net/http", "adapters/http", "")
	if err != nil {
		t.Fatal(err)
	}
	literal := Tech{Pattern: "net/http", Adapter: "adapters/http"}
	if !built.matches("net/http/client") || !literal.matches("net/http/client") {
		t.Fatalf("beide müssen den Substring-Treffer melden")
	}
	if built.matches("os/exec") || literal.matches("os/exec") {
		t.Fatalf("Nicht-Treffer muss für beide false sein")
	}
}

func TestNewTechSubstringNotRegex(t *testing.T) { // AC-FA-RULE-003 / ADR-0015: match: substring nimmt das Muster wörtlich, nicht als Regex
	sub, err := NewTech("Q[A-Za-z]", "adapters/ui", "substring")
	if err != nil {
		t.Fatal(err)
	}
	if sub.matches("QWidget") {
		t.Fatalf("match: substring darf Q[A-Za-z] NICHT als Regex behandeln (QWidget ist kein Substring von \"Q[A-Za-z]\")")
	}
	if !sub.matches("x Q[A-Za-z] y") {
		t.Fatalf("match: substring muss das wörtliche Muster als Teilstring treffen")
	}
}

func TestTechLeakRegexDeterministicOrder(t *testing.T) { // AC-QA-01: ≥2 regex-tech-leak-Befunde stabil nach Pfad sortiert
	files := []FileImports{
		{Path: "adapters/geometry/b.cpp", Layer: "geo", Imports: []Import{{Symbol: "QWidget", Line: 2}}},
		{Path: "adapters/geometry/a.cpp", Layer: "geo", Imports: []Import{{Symbol: "QString", Line: 9}}},
	}
	fs := Evaluate(regexTechModel(t), files)
	if len(fs) != 2 || fs[0].Path != "adapters/geometry/a.cpp" || fs[1].Path != "adapters/geometry/b.cpp" {
		t.Fatalf("regex-tech-leak-Befunde nicht stabil nach Pfad sortiert: %v", fs)
	}
	if fs[0].Rule != "tech-leak" || fs[1].Rule != "tech-leak" {
		t.Fatalf("erwarte zwei tech-leak, got %v", fs)
	}
}
