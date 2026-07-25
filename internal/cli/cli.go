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

	"github.com/pt9912/a-check/internal/adapter/driven/config"
	"github.com/pt9912/a-check/internal/adapter/driven/extract"
	"github.com/pt9912/a-check/internal/adapter/driven/graph"
	"github.com/pt9912/a-check/internal/adapter/driven/report"
	"github.com/pt9912/a-check/internal/hexagon/core"
)

// Run parses args, runs the architecture check and returns the process exit
// code: 0 (no finding), 1 (findings), 2 (usage/config error).
func Run(args []string, out, errw io.Writer) int {
	fs := flag.NewFlagSet("a-check", flag.ContinueOnError)
	fs.SetOutput(errw)
	printConfig := fs.Bool("print-config", false, "kommentiertes .a-check.yml-Gerüst ausgeben (read-only)")
	printMk := fs.Bool("print-mk", false, "includebares a-check.mk ausgeben (read-only)")
	printGraph := fs.Bool("print-graph", false, "Architektur-Graph (Mermaid) aus .a-check.yml ausgeben (read-only, kein Scan)")
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
	if *printGraph {
		// no-scan mode (SPEC-CLI-002): at most one positional argument (the
		// path). Go's flag stops parsing after the first positional, so a stray
		// `--print-graph <pfad> --bogus` would silently become a second
		// positional — reject it as a usage error (exit 2) instead.
		if fs.NArg() > 1 {
			_, _ = fmt.Fprintln(errw, "a-check: --print-graph nimmt höchstens einen Pfad-Parameter")
			return 2
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
		// load-time/config-validation parity to a scan, WITHOUT a file walk:
		// an unknown language fails here exactly as it would in a scan.
		if err := extract.New().Validate(m); err != nil {
			_, _ = fmt.Fprintf(errw, "a-check: %v\n", err)
			return 2
		}
		_, _ = fmt.Fprint(out, graph.New().Render(m))
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
	findings, err := core.Evaluate(m, files)
	if err != nil {
		_, _ = fmt.Fprintf(errw, "a-check: %v\n", err)
		return 2
	}
	return report.New(out, errw).Report(findings)
}

// aCheckImage is the distributed image reference, digest-pinned to the v0.14.0
// release (AC-QA-03, ADR-0007): consumers pin the immutable digest, not a
// moving tag. Pin-Hebung is a conscious commit (ADR-0004).
const aCheckImage = "ghcr.io/pt9912/a-check@sha256:6425c93a9a4359ef28c4da231a2d1db6f421fdaa8f96877ac89d201827c42d09"

const mkFragment = `# a-check.mk — Architektur-Gate via a-check, zum ` + "`include`" + ` in das
# Makefile des konsumierenden Repos. Erzeugt von ` + "`a-check --print-mk`" + `.
#
# A_CHECK_IMAGE wird beim Release auf ` + "`@sha256:…`" + ` digest-gepinnt.
A_CHECK_IMAGE ?= ` + aCheckImage + `

.PHONY: a-check a-check-graph
a-check: ## Architektur: Hexagon-Regeln via a-check (netzlos, read-only).
	docker run --rm --network none -v "$(CURDIR)":/src:ro $(A_CHECK_IMAGE) /src

a-check-graph: ## Architektur-Graph (Mermaid) aus .a-check.yml auf stdout (read-only, kein Scan).
	docker run --rm --network none -v "$(CURDIR)":/src:ro $(A_CHECK_IMAGE) --print-graph /src
`

const sampleConfig = `# .a-check.yml — Architektur-Regeln für a-check (Gerüst, ` + "`a-check --print-config`" + `).
version: 1
languages:
  go: ["**/*.go"]
layers:
  core:     ["internal/core/**"]
  ports:    ["internal/ports/**"]
  adapters: ["internal/adapters/**"]
edges:
  - {from: adapters, to: ports}
  - {from: ports,    to: core}
  # - {from: adapters, to: core}   # falls Adapter Domänentypen direkt referenzieren
adapter_sink: driver-common
tech:
  - {pattern: "gopkg.in/yaml", adapter: "adapters/config"}
  # - {pattern: "Q[A-Za-z]", adapter: "adapters/ui", match: regex}  # RE2 statt Substring (Default: substring)
composition_root: ["cmd/**", "internal/cli/**"]
forbidden_constructs:
  ports: ["impl "]
# constructs:                     # optional: Roh-Text-Monopol (AC-FA-RULE-011/ADR-0027) —
#   - pattern: 'dlopen\s*\('      #   Muster nur in der Zone erlaubt, jedes Vorkommen
#     match: regex                #   ausserhalb ist ein Befund construct-leak. Greift auch
#     adapter: adapters/plugin    #   in Dateien ohne Schicht; Kommentare zaehlen nicht.
#     composition_root: forbid    #   Default allow (Composition Root ausgenommen).
markers:
  ignore_symbols: []
# resolution:                     # optional: Import-Symbol -> Schicht je Sprache (ADR-0016/ADR-0023)
#   kotlin:                       # Multi-Modul (KMP/Gradle): mehrere Module, geteiltes package_base;
#     mode: fixed-root            #   der interne FQN wird gegen die REALEN Dateien unter roots aufgeloest
#     package_base: dev.example   #   (nicht am Wurzel-Praefix). Split-Package (dasselbe Paket ueber zwei
#     roots:                      #   Schicht-Module): ein Top-Level-Symbol (Datei != Name, Kotlin) loest
#       - hexagon/domain/src/commonMain/kotlin/dev/example       #   ueber die reale DEKLARATION auf.
#       - hexagon/application/src/commonMain/kotlin/dev/example  #   FQN real in >=2 Schichten -> Exit 2.
`
