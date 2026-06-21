// Package core is the dependency-free kernel of a-check (ARC-001): the
// architecture model, the port interfaces (ARC-002, realized as Go interfaces
// co-located with the domain) and the rule engine (SPEC-RULE-001). It imports
// nothing outside the standard library — its purity is what `make arch-check`
// (Dogfooding, AC-QA-02) enforces on a-check itself.
package core

// Import is one extracted import or construct hit with its source line.
type Import struct {
	Symbol string
	Line   int
}

// FileImports are the imports (and forbidden-construct hits) extracted from one
// source file, plus the architectural layer the file resolves to.
type FileImports struct {
	Path       string
	Layer      string
	Imports    []Import
	Constructs []Import
}

// Layer is a named architectural layer with repo-relative path globs.
type Layer struct {
	Name  string
	Globs []string
}

// Edge is a directed, allowed dependency between layers (from imports to).
type Edge struct {
	From string
	To   string
}

// Tech maps a framework/tech pattern to the path fragment of its owning adapter.
type Tech struct {
	Pattern string
	Adapter string
}

// Model is the resolved architecture model decoded from `.a-check.yml`.
type Model struct {
	Languages       map[string][]string // language -> file globs
	Layers          []Layer
	Edges           []Edge // allowed directed edges
	Allow           []Edge // explicit extra allowed edges
	AdapterSink     string // shared adapter sink (path fragment), optional
	Techs           []Tech
	CompositionRoot []string            // globs, exempt from layering + tech-leak
	Forbidden       map[string][]string // layer name -> forbidden text patterns
	IgnoreSymbols   []string            // heuristic-boundary allowlist (markers)
}

// Finding is one rule violation. Its fields define the stable sort order
// (SPEC-DET-001): Path, then Line, then Rule.
type Finding struct {
	Path string
	Line int
	Rule string
	Msg  string
}

// ConfigPort loads and strictly decodes the configuration (ARC-004).
type ConfigPort interface {
	Load(path string) (Model, error)
}

// ExtractionPort yields the imports per source file under root (ARC-003).
type ExtractionPort interface {
	Extract(root string, m Model) ([]FileImports, error)
}

// ReportPort renders findings and yields the finding exit code 0/1 (ARC-005).
type ReportPort interface {
	Report(findings []Finding) int
}
