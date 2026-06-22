// Package report is the reporting adapter (ARC-005): it renders findings on
// stdout and a per-rule summary on stderr, and yields the finding exit code
// (0 = none, 1 = at least one finding) per SPEC-CLI-001. Exit code 2
// (usage/config error) is owned by the composition root (ARC-006).
package report

import (
	"fmt"
	"io"
	"sort"

	"github.com/pt9912/a-check/internal/hexagon/core"
	"github.com/pt9912/a-check/internal/hexagon/port"
)

// Adapter implements port.ReportPort.
type Adapter struct {
	out io.Writer
	err io.Writer
}

// New returns a reporting adapter writing to out (findings) and err (summary).
func New(out, err io.Writer) port.ReportPort { return Adapter{out: out, err: err} }

// Report prints findings deterministically and returns the exit code 0/1.
func (a Adapter) Report(findings []core.Finding) int {
	for _, f := range findings {
		writef(a.out, "%s:%d: %s: %s\n", f.Path, f.Line, f.Rule, f.Msg)
	}
	counts := map[string]int{}
	for _, f := range findings {
		counts[f.Rule]++
	}
	rules := make([]string, 0, len(counts))
	for r := range counts {
		rules = append(rules, r)
	}
	sort.Strings(rules)
	for _, r := range rules {
		writef(a.err, "%s: %d\n", r, counts[r])
	}
	writef(a.err, "gesamt: %d Befund(e)\n", len(findings))
	if len(findings) > 0 {
		return 1
	}
	return 0
}

// writef writes to a CLI stream; stream write errors are intentionally ignored
// (a broken pipe is handled by the OS) — kept explicit, no inline suppression
// (AGENTS.md §3.2).
func writef(w io.Writer, format string, a ...any) { _, _ = fmt.Fprintf(w, format, a...) }
