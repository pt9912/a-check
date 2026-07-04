// Package core is the dependency-free kernel of a-check (ARC-001): the
// architecture model and the rule engine (SPEC-RULE-001). It imports nothing
// outside the standard library — its purity is what `make arch-check`
// (Dogfooding, AC-QA-02) enforces on a-check itself. The driven port interfaces
// (ARC-002) live in the sibling package `port`, which references these domain
// types.
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
	Language   string // source-language of this file (drives resolution, ADR-0016)
	Imports    []Import
	Constructs []Import
}

// ResolutionConfig is how one language's import symbols resolve to layers
// (ADR-0016). Mode "" / "path" leaves the import unchanged (Import = Pfad);
// "fixed-root" strips PackageBase, maps "." to "/" and roots the import;
// "relative" resolves ./-style specifiers against the importing file's
// directory and yields no candidate for bare imports or root escapes
// (ADR-0017 — Roots/PackageBase are rejected for it by the config adapter).
type ResolutionConfig struct {
	Mode        string
	Roots       []string
	PackageBase string
}

// PhantomRootConflict witnesses an ambiguous multi-root fixed-root resolution
// (ADR-0020): two distinct roots each lie fully inside a DIFFERENT layer, so
// every import spawns one candidate per root that matches its layer by the root
// prefix alone (phantom candidates across layer boundaries) and targetLayer's
// longest-prefix pick decides the layer instead of the symbol — a silent
// false-negative. The config adapter turns a non-nil conflict into exit code 2.
type PhantomRootConflict struct {
	RootA, LayerA string
	RootB, LayerB string
}

// Layer is a named architectural layer with repo-relative path globs and an
// optional role (domain|app|port|adapter, AC-FA-RULE-006/007) that drives the
// purity rules; a blank role falls back to name inference
// (core/ports/adapters/application), and a layer resolving to no role is only
// edge-checked. Direction (driving|driven, AC-FA-RULE-008) is an OPTIONAL
// dimension ORTHOGONAL to the role: it governs only port-direction-mismatch and
// is never inferred from the name; a blank direction opts the layer out.
type Layer struct {
	Name      string
	Globs     []string
	Role      string
	Direction string
}

// Edge is a directed, allowed dependency between layers (from imports to).
type Edge struct {
	From string
	To   string
}

// Tech maps a framework/tech pattern to the path fragment(s) of its owning
// adapter(s) — the symbol is allowed in EVERY listed adapter (AC-FA-RULE-003
// 0.14.0; a scalar config adapter becomes a one-element list). The pattern
// matches an imported symbol as a substring (default) or, when built via
// NewTech with match=="regex", as an unanchored RE2 regexp (ADR-0015). A
// literal Tech (zero match) matches as substring — backward compatible.
// ForbidCompositionRoot switches off ONLY the composition-root exemption of
// tech-leak for this entry (`composition_root: forbid`); the layer-rule
// exemption of the composition root is untouched.
type Tech struct {
	Pattern               string
	Adapters              []string
	ForbidCompositionRoot bool
	match                 func(string) bool // compiled matcher; nil ⇒ substring on Pattern
}

// Model is the resolved architecture model decoded from `.a-check.yml`.
type Model struct {
	Languages       map[string][]string // language -> file globs
	Layers          []Layer
	Edges           []Edge // allowed directed edges
	Allow           []Edge // explicit extra allowed edges
	AdapterSink     string // shared adapter sink (path fragment), optional
	Techs           []Tech
	CompositionRoot []string                    // globs, exempt from layering (+ tech-leak per entry, AC-FA-RULE-003)
	Forbidden       map[string][]string         // layer name -> forbidden text patterns
	IgnoreSymbols   []string                    // heuristic-boundary allowlist (markers)
	Resolution      map[string]ResolutionConfig // language -> import resolution (ADR-0016)
	Exclude         []string                    // file globs removed from the scan before extraction (ADR-0018)
}

// Finding is one rule violation. Its fields define the stable sort order
// (SPEC-DET-001): Path, then Line, then Rule.
type Finding struct {
	Path string
	Line int
	Rule string
	Msg  string
}
