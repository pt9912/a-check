// Package port holds the hexagon's driven ports (ARC-002): pure abstractions
// the composition root wires to concrete adapters. They speak the domain's
// language — referencing core types (Model, FileImports, Finding) — but import
// no adapter and no tech (port-impurity, AC-FA-RULE-004). The ports layer
// declares a {from: ports, to: core} edge in .a-check.yml so the dogfooding
// check permits exactly these core references.
package port

import "github.com/pt9912/a-check/internal/hexagon/core"

// ConfigPort loads and strictly decodes the configuration (ARC-004).
type ConfigPort interface {
	Load(path string) (core.Model, error)
}

// ExtractionPort yields the imports per source file under root (ARC-003).
type ExtractionPort interface {
	Extract(root string, m core.Model) ([]core.FileImports, error)
}

// ReportPort renders findings and yields the finding exit code 0/1 (ARC-005).
type ReportPort interface {
	Report(findings []core.Finding) int
}
