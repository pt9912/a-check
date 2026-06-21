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
