package cli_test

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/pt9912/a-check/internal/cli"
)

const cfg = `version: 1
languages:
  go: ["**/*.go"]
layers:
  core: ["internal/core/**"]
  adapters: ["internal/adapters/**"]
edges:
  - {from: adapters, to: core}
`

func writeRepo(t *testing.T, files map[string]string) string {
	t.Helper()
	dir := t.TempDir()
	for rel, body := range files {
		p := filepath.Join(dir, rel)
		if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(p, []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	return dir
}

func TestPrintConfig(t *testing.T) { // SPEC-DIST-001: read-only Gerüst
	var out, errb bytes.Buffer
	if code := cli.Run([]string{"--print-config"}, &out, &errb); code != 0 {
		t.Fatalf("exit %d", code)
	}
	if !strings.Contains(out.String(), "version: 1") {
		t.Fatalf("print-config: %q", out.String())
	}
}

func TestPrintMk(t *testing.T) { // SPEC-DIST-001: includebares Fragment
	var out, errb bytes.Buffer
	if code := cli.Run([]string{"--print-mk"}, &out, &errb); code != 0 {
		t.Fatalf("exit %d", code)
	}
	o := out.String()
	if !strings.Contains(o, "A_CHECK_IMAGE") || !strings.Contains(o, "a-check:") {
		t.Fatalf("print-mk: %q", o)
	}
}

func TestUnknownFlag(t *testing.T) { // SPEC-CLI-001: Exit 2
	var out, errb bytes.Buffer
	if code := cli.Run([]string{"--bogus"}, &out, &errb); code != 2 {
		t.Fatalf("expected 2, got %d", code)
	}
}

func TestScanClean(t *testing.T) { // SPEC-CLI-001 happy: Exit 0
	dir := writeRepo(t, map[string]string{
		".a-check.yml":           cfg,
		"internal/core/svc.go":   "package core\nimport \"fmt\"\n\nvar _ = fmt.Sprint\n",
		"internal/adapters/a.go": "package adapters\nimport \"x/internal/core\"\n",
	})
	var out, errb bytes.Buffer
	if code := cli.Run([]string{dir}, &out, &errb); code != 0 {
		t.Fatalf("expected 0, got %d (out=%q)", code, out.String())
	}
}

func TestScanViolation(t *testing.T) { // SPEC-CLI-001 negative: Exit 1
	dir := writeRepo(t, map[string]string{
		".a-check.yml":          cfg,
		"internal/core/bad.go":  "package core\nimport \"x/internal/adapters/http\"\n",
	})
	var out, errb bytes.Buffer
	if code := cli.Run([]string{dir}, &out, &errb); code != 1 {
		t.Fatalf("expected 1, got %d", code)
	}
	if !strings.Contains(out.String(), "core-impurity") {
		t.Fatalf("expected core-impurity: %q", out.String())
	}
}

func TestMissingConfig(t *testing.T) { // SPEC-CLI-001: Exit 2 bei fehlender Config
	var out, errb bytes.Buffer
	if code := cli.Run([]string{t.TempDir()}, &out, &errb); code != 2 {
		t.Fatalf("expected 2, got %d", code)
	}
}

// cfgRegexTech: C++-Repo mit Qt-Muster als RE2-Regex auf dem ui-Adapter (ADR-0015).
const cfgRegexTech = `version: 1
languages:
  cpp: ["**/*.cpp", "**/*.h"]
layers:
  ui:  {globs: ["adapters/ui/**"], role: adapter}
  geo: {globs: ["adapters/geometry/**"], role: adapter}
edges:
  - {from: geo, to: ui}
tech:
  - {pattern: "Q[A-Za-z]", adapter: "adapters/ui", match: regex}
`

func TestTechRegexLeakExit1(t *testing.T) { // AC-FA-RULE-003 / ADR-0015: regex-tech-leak erreicht Exit-Code 1
	dir := writeRepo(t, map[string]string{
		".a-check.yml":            cfgRegexTech,
		"adapters/geometry/g.cpp": "#include <QWidget>\n",
	})
	var out, errb bytes.Buffer
	if code := cli.Run([]string{dir}, &out, &errb); code != 1 {
		t.Fatalf("expected 1, got %d (out=%q)", code, out.String())
	}
	if !strings.Contains(out.String(), "tech-leak") {
		t.Fatalf("expected tech-leak in output: %q", out.String())
	}
}

func TestTechRegexInvalidMatchExit2(t *testing.T) { // AC-FA-CONF-001 / ADR-0015: ungültiges match erreicht Exit-Code 2
	badCfg := strings.Replace(cfgRegexTech, "match: regex", "match: glob", 1)
	dir := writeRepo(t, map[string]string{
		".a-check.yml":      badCfg,
		"adapters/ui/w.cpp": "#include <QWidget>\n",
	})
	var out, errb bytes.Buffer
	if code := cli.Run([]string{dir}, &out, &errb); code != 2 {
		t.Fatalf("expected 2, got %d", code)
	}
}

func TestTechRegexIgnoreSymbols(t *testing.T) { // AC-QA-02 / ADR-0015: markers.ignore_symbols unterdrückt den Q[A-Za-z]/Queue.h-False-Positive
	// Q[A-Za-z] trifft "Queue.h" (der dokumentierte FP); der Marker unterdrückt ihn.
	withMarker := map[string]string{
		".a-check.yml":            cfgRegexTech + "markers:\n  ignore_symbols: [\"Queue.h\"]\n",
		"adapters/geometry/g.cpp": "#include \"Queue.h\"\n",
	}
	var out, errb bytes.Buffer
	if code := cli.Run([]string{writeRepo(t, withMarker)}, &out, &errb); code != 0 {
		t.Fatalf("ignore_symbols muss den Queue.h-FP unterdrücken (Exit 0), got %d (out=%q)", code, out.String())
	}
	// Ohne den Marker schlägt derselbe FP als tech-leak an (Exit 1) — die Heuristik-Grenze ist ausgewiesen, nicht verschwiegen.
	withoutMarker := map[string]string{
		".a-check.yml":            cfgRegexTech,
		"adapters/geometry/g.cpp": "#include \"Queue.h\"\n",
	}
	out.Reset()
	errb.Reset()
	if code := cli.Run([]string{writeRepo(t, withoutMarker)}, &out, &errb); code != 1 {
		t.Fatalf("ohne ignore_symbols erwarte tech-leak (Exit 1), got %d (out=%q)", code, out.String())
	}
}

func TestUnknownLanguageExit2(t *testing.T) { // AC-FA-CONF-001 / slice-017: unbekannte Sprache -> Exit 2 statt still falsch-grün
	cfg := "version: 1\nlanguages:\n  ruby: [\"**/*.rb\"]\nlayers:\n  core: [\"core/**\"]\nedges:\n  - {from: core, to: core}\n"
	dir := writeRepo(t, map[string]string{".a-check.yml": cfg, "core/x.rb": "require 'json'\n"})
	var out, errb bytes.Buffer
	if code := cli.Run([]string{dir}, &out, &errb); code != 2 {
		t.Fatalf("unbekannte Sprache muss Exit 2 liefern, got %d", code)
	}
	if !strings.Contains(errb.String(), "unbekannte Sprache") || !strings.Contains(errb.String(), "ruby") {
		t.Fatalf("Meldung soll die Sprache nennen: %q", errb.String())
	}
	if out.Len() != 0 {
		t.Fatalf("Config-Fehler gehört auf stderr, stdout muss leer sein: %q", out.String())
	}
}

func TestPythonFixedRootResolution(t *testing.T) { // AC-FA-EXTRACT-001 (Python) + AC-FA-CONF-001 Happy-Auflösung: Backend + fixed-root-Rezept (slice-020 §3.3) greifen zusammen
	pyCfg := `version: 1
languages:
  python: ["**/*.py"]
layers:
  core:     ["src/myapp/domain/**"]
  adapters: ["src/myapp/adapters/**"]
edges:
  - {from: adapters, to: core}
resolution:
  python: {mode: fixed-root, roots: ["src/myapp"], package_base: "myapp"}
`
	dir := writeRepo(t, map[string]string{
		".a-check.yml":              pyCfg,
		"src/myapp/domain/model.py": "from myapp.adapters import db\n",
		"src/myapp/adapters/db.py":  "import json\n",
	})
	var out, errb bytes.Buffer
	if code := cli.Run([]string{dir}, &out, &errb); code != 1 {
		t.Fatalf("Python-Domäne importiert Adapter-Modul: erwarte Exit 1, got %d (out=%q err=%q)", code, out.String(), errb.String())
	}
	if !strings.Contains(out.String(), "core-impurity") || !strings.Contains(out.String(), "src/myapp/domain/model.py") {
		t.Fatalf("erwarte core-impurity-Befund für die Domänen-Datei: %q", out.String())
	}
}

func TestMonoRepoMixedUnsupportedExit2(t *testing.T) { // slice-017: Mono-Repo go+typescript(unsupported) -> Exit 2, go rettet nicht
	cfg := "version: 1\nlanguages:\n  go: [\"**/*.go\"]\n  typescript: [\"**/*.ts\"]\nlayers:\n  core: [\"core/**\"]\nedges:\n  - {from: core, to: core}\n"
	dir := writeRepo(t, map[string]string{".a-check.yml": cfg, "core/x.go": "package core\n"})
	var out, errb bytes.Buffer
	if code := cli.Run([]string{dir}, &out, &errb); code != 2 {
		t.Fatalf("gemischte Sprachen mit unsupported -> Exit 2, got %d", code)
	}
	if !strings.Contains(errb.String(), "typescript") {
		t.Fatalf("Meldung soll die unsupported Sprache nennen: %q", errb.String())
	}
}

func TestMonoRepoMultiSupportedRuns(t *testing.T) { // slice-017: Mono-Repo mit nur unterstützten Sprachen (go+cpp) läuft
	cfg := "version: 1\nlanguages:\n  go: [\"**/*.go\"]\n  cpp: [\"**/*.h\", \"**/*.cpp\"]\nlayers:\n  core: [\"core/**\"]\nedges:\n  - {from: core, to: core}\n"
	dir := writeRepo(t, map[string]string{".a-check.yml": cfg, "core/x.go": "package core\n"})
	var out, errb bytes.Buffer
	if code := cli.Run([]string{dir}, &out, &errb); code != 0 {
		t.Fatalf("go+cpp (beide unterstützt) muss laufen (Exit 0), got %d (err=%q)", code, errb.String())
	}
}
