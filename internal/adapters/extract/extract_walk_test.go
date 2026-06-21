package extract

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/pt9912/a-check/internal/core"
)

func TestExtractWalk(t *testing.T) { // SPEC-EXTRACT-001: Walker + Layer + Konstrukte
	dir := t.TempDir()
	if err := os.MkdirAll(filepath.Join(dir, "internal/core"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "internal/core/p.go"),
		[]byte("package core\nimport \"fmt\"\nimpl X\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	m := core.Model{
		Languages:     map[string][]string{"go": {"**/*.go"}},
		Layers:        []core.Layer{{Name: "core", Globs: []string{"internal/core/**"}}},
		Forbidden:     map[string][]string{"core": {"impl "}},
		IgnoreSymbols: []string{"never-matches"},
	}
	files, err := New().Extract(dir, m)
	if err != nil {
		t.Fatal(err)
	}
	if len(files) != 1 || files[0].Layer != "core" {
		t.Fatalf("unexpected files: %+v", files)
	}
	if len(files[0].Constructs) == 0 {
		t.Fatalf("expected forbidden-construct hit for 'impl '")
	}
}
