# Changelog

Alle nennenswerten Änderungen an diesem Projekt werden in dieser Datei
dokumentiert. Das Format folgt [Keep a Changelog](https://keepachangelog.com/de/1.1.0/),
die Versionierung folgt [SemVer](https://semver.org/lang/de/).

## [Unreleased]

## [0.3.0] - 2026-06-23

Dritte Welle: `welle-10b/b2b` — Driving/Driven-Port-Richtung + `LayerOf`-Angleichung an
`targetLayer`. Lastenheft/Spezifikation 0.5.0 → 0.6.0; **sieben** Regeln.

### Added

- **`AC-FA-RULE-008` (Lastenheft 0.5.0→0.6.0):** Driving/Driven-Port-Richtung —
  optionale Schicht-`direction` (`driving`/`driven`), **orthogonal** zur Rolle; neue
  Regel `port-direction-mismatch` (ein Adapter spricht nur Ports **seiner** Richtung),
  **kategorisch** (über `edges`/`allow` nicht aufhebbar). Ohne `direction` keine Prüfung
  (rückwärtskompatibel). `layers`-Objektform um `direction` (und das in 0.5.0 fehlende
  `app`) erweitert. ADR-0012; slice-012.

### Changed

- **`LayerOf` (ADR-0013):** die Schicht-Zuordnung einer Datei nimmt den spezifischsten/
  längsten **literalen** Glob-Präfix (`litPrefixLen`, Angleichung an `targetLayer`) statt
  des Erst-Treffers — Verhaltensänderung nur bei verschachtelten Schicht-Globs. slice-012.
- `--print-mk`/`a-check.mk` und der `aCheckImage`-Default sind auf den
  v0.2.0-Release **digest-gepinnt** (`ghcr.io/pt9912/a-check@sha256:4132a7af…`) —
  Pin-Hebung nach dem Release (AC-QA-03, ADR-0004/ADR-0007).

## [0.2.0] - 2026-06-22

Zweite Welle: das Regel-Modell dispatcht über Layer-**Rollen** statt -Namen und ist
auf vier Schichten (`domain`/`app`/`port`/`adapter`) ausgebaut; Ports dürfen
Domänentypen referenzieren. Lastenheft 0.1.0 → 0.5.0.

### Added

- **`AC-FA-RULE-006` (Lastenheft 0.2.0→0.4.0):** Schicht-**Rollen** — die
  Reinheits-Regeln dispatchen über eine Layer-Rolle (`domain`/`port`/`adapter`, aus
  `role:` oder Namens-Inferenz) statt über die Namen `core`/`ports`/`adapters`; fremd
  benannte Schichten sind voll prüfbar. `layers`-Eintrag als Glob-Liste **oder**
  `{globs, role}`; `lateral-adapter` cross-layer + kategorisch. ADR-0009; b1 (ADR-0010)
  macht `adapterSeg`/`targetLayer` vollständig namensunabhängig (längster,
  segment-bewusster Präfix). welle-10a/b1.
- **`AC-FA-RULE-007` (Lastenheft 0.4.0→0.5.0):** neue Schicht-Rolle `app`
  (Application-/Use-Case-Schicht) — darf `domain`+`port` referenzieren, aber keinen
  Adapter/Tech: neuer Befund `app-impurity`. Zugleich `domain` verschärft (Import auf
  `app`/`port`/`adapter`/Tech ⇒ `core-impurity`, kategorisch — „Domäne kennt keine
  Ports"); `role`-Schema um `app`. ADR-0011. **Breaking für geprüfte Repos:** eine
  `role: domain`-Schicht, die einen `port`/`app`-Layer importiert, wird jetzt rot
  (vorher per deklarierter Kante grün) — Migration: Port-/Use-Case-Nutzung in eine
  `role: app`-Schicht heben. a-checks Eigen-Dogfooding bleibt unverändert grün;
  Multi-Linsen-Review.
- Benutzerhandbuch 1.6: die Schicht-`role` dokumentiert (Objektform `{globs, role}`,
  Rollen, Namens-Inferenz, Vorrang, Vier-Schichten-`app`-Modell).

### Changed

- **`AC-FA-RULE-004` (Lastenheft 0.1.0→0.2.0):** Ports dürfen jetzt Domänen-/
  Kern-Typen referenzieren (`ports → core` per deklarierter Kante); `port-impurity`
  feuert nur noch bei Adapter-/Tech-Import, nicht mehr bei Kern-Import. Motiviert
  durch die Vier-Repo-Evidenz (b-cad/d-migrate-Ports referenzieren die Domäne);
  ADR-0008 (Accepted). a-check selbst auf eine echte `ports`-Schicht umgebaut
  (`internal/hexagon/{core,port}`, `internal/adapter/driven/*`), Dogfooding grün
  (AC-QA-02); Multi-Linsen-Review (`docs/reviews/2026-06-22-…`).
- `--print-mk`/`a-check.mk` und der `aCheckImage`-Default sind auf den
  v0.1.0-Release **digest-gepinnt**
  (`ghcr.io/pt9912/a-check@sha256:13459f44…`) statt auf die Tag-Form — Pin-Hebung
  nach dem ersten Release (AC-QA-03, ADR-0004/ADR-0007).

## [0.1.0] - 2026-06-21

Erstes Release: a-check als sprach-agnostischer Hexagonal-Architektur-Checker
(text-heuristisch, netzlos, distroless/static) inkl. Harness, Quality-Gates,
Durchsetzungsschicht und CI-/Release-Pipeline. Distribution als digest-gepinntes
GHCR-Image + `--print-mk`/`a-check.mk`.

### Added

- Bootstrap — Harness-Gerüst (AGENTS.md, harness/-Trias, Lastenheft 0.1.0)
  und das Doku-Gate `make doc-check` via Schwester-Tool d-check
  (digest-gepinnt, netzlos, read-only).
- slice-001 — Fundament-ADRs ADR-0001..0004 (Go als Implementierungssprache;
  text-heuristische Extraktion; Config-Modell `.a-check.yml`; Distribution
  inkl. `--print-mk`/`a-check.mk`); Status Accepted.
- slice-002 — Technik-Stratum `spec/spezifikation.md`
  (SPEC-CONF/EXTRACT/RULE/CLI/DET/DIST-001) und Sicht-Stratum
  `spec/architecture.md` (ARC-001..006); Spec-Strata in `harness/conventions.md`
  (MR-004) deklariert.
- slice-003 — Go-Implementierung (Hexagon: `internal/core`/`adapters`/`cli`,
  `cmd/a-check`): fünf Regeln AC-FA-RULE-001..005, text-heuristische
  Extraktion C++/Go/Rust/Kotlin (AC-FA-EXTRACT-001), strict-decode
  `.a-check.yml` (AC-FA-CONF-001), CLI/Exit-Codes (AC-FA-CLI-001),
  `--print-config`/`--print-mk` (AC-FA-DIST-001), Determinismus (AC-QA-01).
  Multi-Stage-Dockerfile (static/distroless, digest-gepinnte Bases, AC-QA-02/03).
- slice-003 — Quality-Gates `make lint`/`test`/`coverage-gate`/`arch-check`/`gates`
  (Dockerfile-Stages, Muster d-check/u-boot); `a-check.mk` via `--print-mk`.
  Lint-Profil golangci-lint v2 (ADR-0005); Coverage-Gate 90 % (ADR-0006, Ist 92,6 %).
  Dogfooding: a-check prüft seine eigene Hexagon-Architektur (AC-QA-02), 0 Befunde.
- slice-004 — Durchsetzungsschicht: Meta-Gates `make gate-consistency`
  (dokumentierte Targets ↔ Makefile + `.d-check.yml`-Module; Schutz gegen
  Harness-Lügen, schützt die doc-check-Beweisaussage AC-QA-02) und
  `make record-gates` (inhaltsbasierter Working-Tree-Hash-Nachweis) plus
  `.claude`-Stop-Hook als Handoff-Gate (fail-closed, loop-guarded, bootstrap-aware).
- slice-005 — Durchsetzungsschicht vollständig: PreToolUse-Command-Guard
  (`.claude/hooks/pretooluse-command-guard.sh`) lehnt Host-Toolchain/-Paketmanager
  (go/golangci-lint/pip/npm/cargo/apt/brew/…) vor der Ausführung fail-closed ab
  (Tool-Call-Gate, AGENTS §3.1); Selbsttest `make guard-selftest` (in `make gates`).
- slice-006 — CI: PR-/Push-Workflow `.github/workflows/ci.yml` (SHA-gepinnt,
  `permissions: {}`, Tags ausgenommen) fährt `make ci` (= `gates` + `make image-test`:
  AC-FA-DIST-001 `--print-mk`/`--print-config`/unbekanntes Flag + nativ==Container-
  Determinismus, AC-QA-02) und `make trace-check` (AC-/ADR-/MR-/slice-ID je Commit,
  AGENTS §5). Dockerfile-OCI-Labels (`org.opencontainers.image.*`) + `VERSION`-Build-Arg.
- slice-007 — Release-Pipeline `.github/workflows/release.yml` (auf `v*`-Tags,
  SHA-gepinnt): SemVer-Validate → `make ci VERSION=` → GHCR-Login → Tag (`:latest`
  nur stabil, ADR-0007) → OCI-Label-Verify → Push → GitHub-Release mit Digest-Pin.
  `:latest`-Tag-Politik in ADR-0007 (Accepted); `releasing.md` auf die reale
  Pipeline aktualisiert.
- slice-008 — lokaler `commit-msg`-Hook (`.githooks/commit-msg` + `make hooks`):
  ruft `trace-check` vor dem Commit (AGENTS §5), opt-in pro Klon; dieselbe
  Wahrheit wie CI/`make trace-check`.
