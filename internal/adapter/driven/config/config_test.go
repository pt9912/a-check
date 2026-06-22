package config

import (
	"os"
	"path/filepath"
	"testing"
)

const valid = `version: 1
languages:
  go: ["**/*.go"]
layers:
  core: ["core/**"]
  adapters: ["adapters/**"]
edges:
  - {from: adapters, to: core}
adapter_sink: driver-common
tech:
  - {pattern: "net/http", adapter: "adapters/http"}
markers:
  ignore_symbols: ["Queue.h"]
`

func write(t *testing.T, body string) string {
	t.Helper()
	p := filepath.Join(t.TempDir(), ".a-check.yml")
	if err := os.WriteFile(p, []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
	return p
}

func TestValidLoad(t *testing.T) { // AC-FA-CONF-001 happy
	m, err := New().Load(write(t, valid))
	if err != nil {
		t.Fatalf("valid config failed: %v", err)
	}
	if len(m.Layers) != 2 || len(m.Edges) != 1 || m.AdapterSink != "driver-common" {
		t.Fatalf("model not decoded as expected: %+v", m)
	}
	if len(m.IgnoreSymbols) != 1 || m.IgnoreSymbols[0] != "Queue.h" {
		t.Fatalf("markers not decoded: %+v", m.IgnoreSymbols)
	}
}

func TestUnknownKeyFailsClosed(t *testing.T) { // AC-FA-CONF-001 negative (strict decode)
	if _, err := New().Load(write(t, valid+"bogus: 1\n")); err == nil {
		t.Fatal("expected error on unknown key (fail-closed)")
	}
}

func TestMissingRequiredBlock(t *testing.T) { // AC-FA-CONF-001 boundary
	body := "version: 1\nlanguages:\n  go: [\"**/*.go\"]\nlayers:\n  core: [\"core/**\"]\n"
	if _, err := New().Load(write(t, body)); err == nil {
		t.Fatal("expected error: 'edges' missing")
	}
}

func TestWrongVersion(t *testing.T) {
	if _, err := New().Load(write(t, "version: 2\nlanguages:\n  go: [\"a\"]\nlayers:\n  c: [\"c\"]\nedges:\n  - {from: c, to: c}\n")); err == nil {
		t.Fatal("expected error on unsupported version")
	}
}

func TestLayerObjectFormWithRole(t *testing.T) { // AC-FA-RULE-006 / AC-FA-CONF-001: Objektform mit role
	body := "version: 1\nlanguages:\n  go: [\"**/*.go\"]\nlayers:\n  geometry: {globs: [\"geometry/**\"], role: adapter}\n  core: [\"core/**\"]\nedges:\n  - {from: geometry, to: core}\n"
	m, err := New().Load(write(t, body))
	if err != nil {
		t.Fatalf("object-form layer failed: %v", err)
	}
	got := ""
	for _, l := range m.Layers {
		if l.Name == "geometry" {
			got = l.Role
		}
	}
	if got != "adapter" {
		t.Fatalf("expected geometry role 'adapter', got %q (%+v)", got, m.Layers)
	}
}

func TestLayerObjectUnknownKeyFailsClosed(t *testing.T) { // SPEC-CONF-001: strict auch im Objekt (yaml.Node-Gotcha)
	body := "version: 1\nlanguages:\n  go: [\"**/*.go\"]\nlayers:\n  core: {globs: [\"core/**\"], bogus: 1}\nedges:\n  - {from: core, to: core}\n"
	if _, err := New().Load(write(t, body)); err == nil {
		t.Fatal("expected error on unknown key inside a layer object (fail-closed)")
	}
}

func TestLayerInvalidRoleFailsClosed(t *testing.T) { // AC-FA-RULE-006: role nur domain|port|adapter
	body := "version: 1\nlanguages:\n  go: [\"**/*.go\"]\nlayers:\n  core: {globs: [\"core/**\"], role: domainx}\nedges:\n  - {from: core, to: core}\n"
	if _, err := New().Load(write(t, body)); err == nil {
		t.Fatal("expected error on invalid role")
	}
}

func TestLayerScalarFailsClosed(t *testing.T) { // SPEC-CONF-001: Scalar-Layer ist weder Glob-Liste noch Objekt
	body := "version: 1\nlanguages:\n  go: [\"**/*.go\"]\nlayers:\n  core: \"core/**\"\nedges:\n  - {from: core, to: core}\n"
	if _, err := New().Load(write(t, body)); err == nil {
		t.Fatal("expected error for scalar layer value")
	}
}

func TestLayerSeqBadElementFailsClosed(t *testing.T) { // Glob-Liste mit Nicht-String-Element
	body := "version: 1\nlanguages:\n  go: [\"**/*.go\"]\nlayers:\n  core: [{x: 1}]\nedges:\n  - {from: core, to: core}\n"
	if _, err := New().Load(write(t, body)); err == nil {
		t.Fatal("expected error for non-string glob element")
	}
}

func TestLayerObjectBadGlobsTypeFailsClosed(t *testing.T) { // globs im Objekt muss eine Liste sein
	body := "version: 1\nlanguages:\n  go: [\"**/*.go\"]\nlayers:\n  core: {globs: \"core/**\", role: domain}\nedges:\n  - {from: core, to: core}\n"
	if _, err := New().Load(write(t, body)); err == nil {
		t.Fatal("expected error for non-list globs in object form")
	}
}

func TestLayerRoleAppAccepted(t *testing.T) { // AC-FA-RULE-007: role: app wird akzeptiert (Positiv-Test)
	body := "version: 1\nlanguages:\n  go: [\"**/*.go\"]\nlayers:\n  application: {globs: [\"application/**\"], role: app}\n  core: [\"core/**\"]\nedges:\n  - {from: application, to: core}\n"
	m, err := New().Load(write(t, body))
	if err != nil {
		t.Fatalf("role: app sollte akzeptiert werden, got %v", err)
	}
	got := ""
	for _, l := range m.Layers {
		if l.Name == "application" {
			got = l.Role
		}
	}
	if got != "app" {
		t.Fatalf("expected application role 'app', got %q (%+v)", got, m.Layers)
	}
}
