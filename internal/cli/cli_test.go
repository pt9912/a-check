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
