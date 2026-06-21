// Package cli is the composition root and CLI logic (ARC-006): it parses flags,
// wires the adapters to the core rule engine and owns the usage/config exit
// code 2 (SPEC-CLI-001, SPEC-DIST-001). It lives under internal/ so its
// contract is black-box testable (package cli_test); cmd/a-check is the thin
// os.Exit entrypoint.
package cli

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"path/filepath"

	"github.com/pt9912/a-check/internal/adapters/config"
	"github.com/pt9912/a-check/internal/adapters/extract"
	"github.com/pt9912/a-check/internal/adapters/report"
	"github.com/pt9912/a-check/internal/core"
)

// Run parses args, runs the architecture check and returns the process exit
// code: 0 (no finding), 1 (findings), 2 (usage/config error).
func Run(args []string, out, errw io.Writer) int {
	fs := flag.NewFlagSet("a-check", flag.ContinueOnError)
	fs.SetOutput(errw)
	printConfig := fs.Bool("print-config", false, "kommentiertes .a-check.yml-Gerüst ausgeben (read-only)")
	printMk := fs.Bool("print-mk", false, "includebares a-check.mk ausgeben (read-only)")
	if err := fs.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return 0
		}
		return 2 // unbekanntes Flag o. Ä. (SPEC-CLI-001)
	}
	if *printConfig {
		_, _ = fmt.Fprint(out, sampleConfig)
		return 0
	}
	if *printMk {
		_, _ = fmt.Fprint(out, mkFragment)
		return 0
	}

	root := "/src"
	if fs.NArg() > 0 {
		root = fs.Arg(0)
	}
	m, err := config.New().Load(filepath.Join(root, ".a-check.yml"))
	if err != nil {
		_, _ = fmt.Fprintf(errw, "a-check: %v\n", err)
		return 2
	}
	files, err := extract.New().Extract(root, m)
	if err != nil {
		_, _ = fmt.Fprintf(errw, "a-check: %v\n", err)
		return 2
	}
	return report.New(out, errw).Report(core.Evaluate(m, files))
}

// aCheckImage is the distributed image reference. The digest pin (`@sha256:…`)
// is set at release (AC-QA-03); the tag form is the bootstrap placeholder.
const aCheckImage = "ghcr.io/pt9912/a-check:0.1.0"

const mkFragment = `# a-check.mk — Architektur-Gate via a-check, zum ` + "`include`" + ` in das
# Makefile des konsumierenden Repos. Erzeugt von ` + "`a-check --print-mk`" + `.
#
# A_CHECK_IMAGE wird beim Release auf ` + "`@sha256:…`" + ` digest-gepinnt (AC-QA-03).
A_CHECK_IMAGE ?= ` + aCheckImage + `

.PHONY: a-check
a-check: ## Architektur: Hexagon-Regeln via a-check (netzlos, read-only).
	docker run --rm --network none -v "$(CURDIR)":/src:ro $(A_CHECK_IMAGE) /src
`

const sampleConfig = `# .a-check.yml — Architektur-Regeln für a-check (Gerüst, ` + "`a-check --print-config`" + `).
version: 1
languages:
  go: ["**/*.go"]
layers:
  core:     ["internal/core/**"]
  adapters: ["internal/adapters/**"]
edges:
  - {from: adapters, to: core}
adapter_sink: driver-common
tech:
  - {pattern: "gopkg.in/yaml", adapter: "adapters/config"}
composition_root: ["cmd/**", "internal/cli/**"]
forbidden_constructs:
  ports: ["impl "]
markers:
  ignore_symbols: []
`
