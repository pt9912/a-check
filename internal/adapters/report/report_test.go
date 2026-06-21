package report_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/pt9912/a-check/internal/adapters/report"
	"github.com/pt9912/a-check/internal/core"
)

func TestReportFindings(t *testing.T) {
	var out, errb bytes.Buffer
	code := report.New(&out, &errb).Report([]core.Finding{
		{Path: "a.go", Line: 3, Rule: "core-impurity", Msg: "x"},
		{Path: "b.go", Line: 1, Rule: "tech-leak", Msg: "y"},
	})
	if code != 1 {
		t.Fatalf("expected exit 1, got %d", code)
	}
	if !strings.Contains(out.String(), "a.go:3: core-impurity: x") {
		t.Fatalf("stdout: %q", out.String())
	}
	if !strings.Contains(errb.String(), "gesamt: 2 Befund(e)") {
		t.Fatalf("stderr: %q", errb.String())
	}
}

func TestReportClean(t *testing.T) {
	var out, errb bytes.Buffer
	code := report.New(&out, &errb).Report(nil)
	if code != 0 {
		t.Fatalf("expected exit 0, got %d", code)
	}
	if out.String() != "" {
		t.Fatalf("expected empty stdout, got %q", out.String())
	}
	if !strings.Contains(errb.String(), "gesamt: 0 Befund(e)") {
		t.Fatalf("stderr: %q", errb.String())
	}
}
