package extract

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/pt9912/a-check/internal/hexagon/core"
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
	if len(files[0].ForbiddenHits) == 0 {
		t.Fatalf("expected forbidden-construct hit for 'impl '")
	}
}

// TestConstructHitsScanWide covers the extraction half of AC-FA-RULE-011
// (ADR-0027): the raw-text hits are collected for EVERY scanned file — layered
// or not — carry the ENTRY INDEX, and are matched against the comment-stripped
// source (a hit living only in a comment never fires).
func TestConstructHitsScanWide(t *testing.T) {
	dir := t.TempDir()
	files := map[string]string{
		"src/adapters/plugin/loader.cpp": "void load() { dlopen(p); }\n",
		"src/adapters/io/reader.cpp":     "void read() {\n  dlopen(p); // erlaubt?\n}\n",
		"src/main.cpp":                   "int main() { dlopen(p); }\n",  // in keinem Layer-Glob
		"src/adapters/io/doc.cpp":        "// dlopen(p) nur im Kommentar\n/* auch dlsym(p) */\n",
	}
	for rel, body := range files {
		p := filepath.Join(dir, rel)
		if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(p, []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	dlopen, err := core.NewConstruct("dlopen(", []string{"adapters/plugin"}, "", "")
	if err != nil {
		t.Fatal(err)
	}
	dlsym, err := core.NewConstruct("dlsym(", []string{"adapters/plugin"}, "", "")
	if err != nil {
		t.Fatal(err)
	}
	m := core.Model{
		Languages:  map[string][]string{"cpp": {"**/*.cpp"}},
		Layers:     []core.Layer{{Name: "adapters", Globs: []string{"src/adapters/**"}}},
		Constructs: []core.Construct{dlopen, dlsym},
	}
	out, err := New().Extract(dir, m)
	if err != nil {
		t.Fatal(err)
	}
	hits := map[string][]core.ConstructHit{}
	for _, f := range out {
		hits[f.Path] = f.ConstructHits
	}
	if len(hits) != 4 {
		t.Fatalf("erwarte 4 gescannte Dateien, got %v", hits)
	}
	if got := hits["src/main.cpp"]; len(got) != 1 || got[0] != (core.ConstructHit{Entry: 0, Line: 1}) {
		t.Fatalf("schichtlose Datei muss Treffer liefern (scan-weit): %v", got)
	}
	if got := hits["src/adapters/io/reader.cpp"]; len(got) != 1 || got[0].Line != 2 {
		t.Fatalf("Treffer mit Zeilennummer erwartet: %v", got)
	}
	if got := hits["src/adapters/plugin/loader.cpp"]; len(got) != 1 {
		t.Fatalf("die Zone liefert den Treffer ebenfalls (die Zonen-Wertung liegt im Kern): %v", got)
	}
	if got := hits["src/adapters/io/doc.cpp"]; len(got) != 0 {
		t.Fatalf("Treffer nur im Kommentar dürfen nicht entstehen (ausgewiesene grep-Divergenz): %v", got)
	}
}

// TestConstructHitsExcludedFile: exclude greift VOR der Extraktion, also auch
// vor der Konstrukt-Erkennung (SPEC-EXTRACT-001).
func TestConstructHitsExcludedFile(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "gen.cpp"), []byte("dlopen(p);\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	c, err := core.NewConstruct("dlopen(", []string{"adapters/plugin"}, "", "")
	if err != nil {
		t.Fatal(err)
	}
	m := core.Model{
		Languages:  map[string][]string{"cpp": {"**/*.cpp"}},
		Layers:     []core.Layer{{Name: "adapters", Globs: []string{"adapters/**"}}},
		Constructs: []core.Construct{c},
		Exclude:    []string{"gen.cpp"},
	}
	out, err := New().Extract(dir, m)
	if err != nil {
		t.Fatal(err)
	}
	if len(out) != 0 {
		t.Fatalf("ausgeschlossene Datei darf keine Konstrukt-Treffer liefern: %+v", out)
	}
}
