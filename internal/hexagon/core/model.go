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

// ConstructHit is one raw-source-text match of a `constructs` entry: the INDEX
// of the entry in Model.Constructs plus the source line (ADR-0027). It carries
// the index, not the pattern text, so the rule engine decides the zone on the
// entry itself — two entries with the SAME pattern but different zones stay
// distinguishable (a pattern-text lookup would silently pick the first).
type ConstructHit struct {
	Entry int
	Line  int
}

// Limit is one import line whose SPELLING keeps it from becoming a judged edge,
// with the human-readable form that names the reason (ADR-0031). Form is a fixed
// phrase from a small set, not free text — it is the diagnostic output itself,
// so a drifting wording would drift the CLI contract.
type Limit struct {
	Line int
	Form string
}

// LimitNote is a diagnosed limit with its file — the flattened, scan-wide view
// the CLI prints (ADR-0031).
type LimitNote struct {
	Path string
	Line int
	Form string
}

// FileImports are the imports (and text-pattern hits) extracted from one source
// file, plus the architectural layer the file resolves to.
type FileImports struct {
	Path     string
	Layer    string
	Language string // source-language of this file (drives resolution, ADR-0016)
	Imports  []Import
	// ForbiddenHits are the LAYER-bound `forbidden_constructs` matches feeding
	// port-impurity (AC-FA-RULE-004). Named apart from ConstructHits/Constructs
	// on purpose: those are the ZONE-bound `constructs` block (ADR-0027) — one
	// name for both would blur exactly the distinction this code must keep.
	ForbiddenHits []Import
	// ConstructHits are the scan-wide `constructs` matches feeding construct-leak
	// (AC-FA-RULE-011); empty when no constructs block is configured.
	ConstructHits []ConstructHit
	// Limits are import lines the language backend did NOT extract because their
	// SPELLING falls outside its patterns (ADR-0031) — a relative Python import,
	// a second directive on one line. They are diagnostic only: they never become
	// imports and never reach a rule. The second limit class (extracted but
	// structurally unresolvable) is derived in HeuristicLimits from Imports plus
	// the resolution mode, so the extractor stays free of config knowledge.
	Limits []Limit
	// Declarations are the file's top-level declaration names (Kotlin: fun,
	// val/var, class/object/interface/typealias, ADR-0023). Only a
	// declaration-aware backend fills them (0.18.0: Kotlin); every other backend
	// leaves them nil (no-op). They feed the declaration-aware fixed-root
	// resolution so a top-level symbol whose file ≠ its name (extension function,
	// second class in a file) resolves to the module that really declares it.
	Declarations []string
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

// Layer is a named architectural layer with repo-relative path globs and an
// optional role (domain|app|port|adapter, AC-FA-RULE-006/007) that drives the
// purity rules; a blank role falls back to name inference
// (core/ports/adapters/application), and a layer resolving to no role is only
// edge-checked. Direction (AC-FA-RULE-008, ADR-0036) is an OPTIONAL
// dimension ORTHOGONAL to the role, and its VALUE SET DEPENDS ON THE ROLE:
// inbound|outbound on a port, driving|driven on an adapter — a port does not
// drive, it is used. It governs only port-direction-mismatch, is never inferred
// from the name, and a blank direction opts the layer out.
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

// Construct is a RAW-SOURCE-TEXT pattern that may appear ONLY inside its owning
// zone(s) — the `constructs` block (AC-FA-RULE-011, ADR-0027). It shares the
// scoping mechanic of Tech (zone as path fragment or list, substring default or
// unanchored RE2 via NewConstruct's match=="regex", per-entry composition-root
// exemption) but matches a source LINE instead of an extracted import symbol, so
// it catches constructs that are no import at all (a `dlopen(` call reachable
// through a transitive header). Zones is never empty and Pattern never blank —
// NewConstruct rejects both fail-closed.
type Construct struct {
	Pattern               string
	Zones                 []string
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
	Constructs      []Construct                 // raw-text zone monopolies (AC-FA-RULE-011)
	CompositionRoot []string                    // globs, exempt from layering (+ tech-/construct-leak per entry, AC-FA-RULE-003/011)
	Forbidden       map[string][]string         // layer name -> forbidden text patterns
	IgnoreSymbols   []string                    // heuristic-boundary allowlist (markers)
	Resolution      map[string]ResolutionConfig // language -> import resolution (ADR-0016)
	Exclude         []string                    // file globs removed from the scan before extraction (ADR-0018)
}

// Finding is one rule violation. Its fields define the TOTAL sort order
// (SPEC-DET-001): Path, then Line, then Rule, then Msg. Msg is what makes the
// order total — one file line can carry several findings of the SAME rule (two
// constructs or two forbidden_constructs patterns matching one line), and
// without it their order would depend on the input ordering (AC-QA-01).
type Finding struct {
	Path string
	Line int
	Rule string
	Msg  string
}
