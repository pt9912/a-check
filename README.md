# a-check

**English** · [Deutsch](README.de.md)

[![CI](https://github.com/pt9912/a-check/actions/workflows/ci.yml/badge.svg)](https://github.com/pt9912/a-check/actions/workflows/ci.yml)

Cross-language hexagonal-architecture checker — deterministic, side-effect-free,
text-heuristic, shipped as a container image.

## What is a-check?

**a-check** enforces a repository's hexagonal layered architecture **across languages**, driven
by a single config file. Seven universal rules, each one requirement in the
[requirements spec](spec/lastenheft.md):

- `core-impurity` — the core (domain, innermost layer) imports neither a port, an adapter, nor a framework/tech
  ([AC-FA-RULE-001](spec/lastenheft.md#ac-fa-rule-001--kern-reinheit-regel-core-impurity))
- `lateral-adapter` — an adapter imports no other adapter (except the shared sink)
  ([AC-FA-RULE-002](spec/lastenheft.md#ac-fa-rule-002--keine-lateralen-adapter-kanten-regel-lateral-adapter))
- `tech-leak` — a framework/tech appears only in its adapter
  ([AC-FA-RULE-003](spec/lastenheft.md#ac-fa-rule-003--tech-kapselung-regel-tech-leak))
- `port-impurity` — a port imports no adapter and no framework/tech (it may reference the core's domain types)
  ([AC-FA-RULE-004](spec/lastenheft.md#ac-fa-rule-004--port-disziplin-regel-port-impurity))
- `wrong-direction` — layer edges are one-way
  ([AC-FA-RULE-005](spec/lastenheft.md#ac-fa-rule-005--schicht-richtung-regel-wrong-direction))
- `app-impurity` — the application layer imports no adapter and no framework/tech (it may use the domain and ports)
  ([AC-FA-RULE-007](spec/lastenheft.md#ac-fa-rule-007--rolle-app-und-strenge-domain))
- `port-direction-mismatch` — an adapter talks only to ports of **its** direction (optional `driving`/`driven` dimension, orthogonal to the role)
  ([AC-FA-RULE-008](spec/lastenheft.md#ac-fa-rule-008--driving-driven-port-richtung-regel-port-direction-mismatch))

Imports are extracted **text-heuristically** per language (C++/Go/Rust/Kotlin/Java/Python/C#/TypeScript)
([AC-FA-EXTRACT-001](spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion)).
Every finding names the file, line, rule and reason; exit codes: `0` clean, `1` findings,
`2` usage/config error
([AC-FA-CLI-001](spec/lastenheft.md#ac-fa-cli-001--aufruf-scan-wurzel-und-exit-codes)).

## Why a-check?

Four functionally overlapping `arch-check.sh` variants grew across the sibling repositories —
C++ via `#include` heuristics (`b-cad`), Go via `go list` (`d-check`), Rust via `use` heuristics
(`grid-guide`), Kotlin via Gradle module boundaries (`d-migrate`): four languages, four
mechanisms, the same seven rules. a-check replaces them with **one** tool:

- **Configuration instead of fork:** repo-specific layer/tech rules live declaratively in
  `.a-check.yml`
  ([AC-FA-CONF-001](spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml)),
  not in copied scripts.
- **One distribution path:** a digest-pinned container image plus a shipped `a-check.mk`
  instead of N maintained copies
  ([AC-FA-DIST-001](spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk)).

It is the **architecture counterpart to `d-check`** (doc references): the same founding logic
(replace a family of drifting scripts with one tool), one abstraction level higher.

## Core idea

**Architecture is an import graph with checkable invariants.** Whether the core stays pure, an
adapter imports laterally, or a layer edge runs against its direction is machine-decidable —
a-check turns these invariants into a gate instead of a review opinion.

The rule is **report, never repair**: a-check is a pure read-only tool. The **heuristic boundary**
(text-based, not a full parser per language) is disclosed, not hidden; an allowlist/marker
exception is configurable
([AC-QA-02](spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)).

## What makes it trustworthy?

- **Determinism:** identical input ⇒ byte-identical, stably sorted output
  ([AC-QA-01](spec/lastenheft.md#ac-qa-01--determinismus)).
- **Hermetic & network-less:** never writes into the checked repo, runs with `--network none`
  on distroless/static
  ([AC-QA-02](spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)).
- **No silent defaults:** any invalid `.a-check.yml` aborts with exit 2 (strict decode,
  [AC-FA-CONF-001](spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml)).
- **Reproducible:** the image and `a-check.mk` reference a `@sha256:` digest
  ([AC-QA-03](spec/lastenheft.md#ac-qa-03--reproduzierbarkeit)).
- **Dogfooding:** a-check checks its own hexagonal architecture on every `make arch-check` —
  against the [self-configuration](.a-check.yml), 0 findings.

## Usage

Against the published image (digest-pinned, network-less, read-only):

```bash
docker run --rm --network none -v "$PWD:/src:ro" \
  ghcr.io/pt9912/a-check@sha256:aef28cfe25bb054b1b0eb28420222a45b9f6ce9425b7ffd0f55e6ae56f295b56 /src
```

Consumers wire a-check in as a `make a-check` gate — **without a script copy**: the shipped
[`a-check.mk`](a-check.mk) (produced by `a-check --print-mk`) is `include`-d, plus an
`.a-check.yml`. `A_CHECK_IMAGE` is pinned to the release's `@sha256:` digest
([AC-FA-DIST-001](spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk),
[AC-QA-03](spec/lastenheft.md#ac-qa-03--reproduzierbarkeit)); raising the pin is a deliberate
commit. The release process is described in [`docs/user/releasing.md`](docs/user/releasing.md).

Locally (dogfooding, no pull):

```bash
make build        # build the static/distroless image
make arch-check   # a-check checks itself (network-less, read-only)
```

## Configuration (`.a-check.yml`)

In the repo root; strictly decoded (unknown key ⇒ exit 2):

```yaml
version: 1
languages:
  go: ["**/*.go"]
layers:
  core:     ["internal/core/**"]
  ports:    ["internal/ports/**"]
  adapters: ["internal/adapters/**"]
edges:
  - {from: adapters, to: ports}
  - {from: ports,    to: core}     # ports may reference domain types
  # - {from: adapters, to: core}   # if adapters reference domain types directly
tech:
  - {pattern: "net/http", adapter: adapters/http}                 # substring (default)
  # - {pattern: "Q[A-Za-z]", adapter: adapters/ui, match: regex}  # RE2 regex (e.g. Qt)
```

The full schema is in the
[specification §SPEC-CONF-001](spec/spezifikation.md#spec-conf-001--konfigurationsschema);
a living example is [this repo's self-configuration](.a-check.yml).
`a-check --print-config` prints a commented skeleton.

## Getting started

| Document | Content |
|---|---|
| [`docs/user/benutzerhandbuch.md`](docs/user/benutzerhandbuch.md) | **User manual** (German) — installation, usage, `.a-check.yml`, troubleshooting |
| [`docs/user/releasing.md`](docs/user/releasing.md) | Release process — tagging, GHCR, digest pin |
| [`spec/lastenheft.md`](spec/lastenheft.md) | Requirements (`AC-FA-*`, `AC-QA-*`), acceptance criteria |
| [`spec/spezifikation.md`](spec/spezifikation.md) | Algorithms, `.a-check.yml` schema, exit codes (`SPEC-*`) |
| [`spec/architecture.md`](spec/architecture.md) | Hexagon components and roles (`ARC-*`) |
| [`docs/plan/adr/README.md`](docs/plan/adr/README.md) | Architecture decisions (ADR index) |
| [`harness/README.md`](harness/README.md) | Harness entry: source precedence, guides, sensors |
| [`AGENTS.md`](AGENTS.md) | Briefing for AI coding agents, hard rules |
| [`CHANGELOG.md`](CHANGELOG.md) | Change history |

> The specification, requirements and user manual are maintained in German. The `AC-*`/`SPEC-*`/`ARC-*`
> anchors above therefore point to German headings.

## Development

The host needs only `git`, GNU `make`, `bash` and Docker
([`AGENTS.md`](AGENTS.md) §3.1).

```bash
make help     # available targets
make gates    # all inner gates (lint/test/coverage-gate/arch-check/doc-check/gate-consistency/guard-selftest)
make ci       # gates + image-test (CI-equivalent); make trace-check checks commit IDs
```

## License

This project is licensed under the [MIT License](LICENSE).
