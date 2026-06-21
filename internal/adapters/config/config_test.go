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
