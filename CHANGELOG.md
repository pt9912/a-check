# Changelog

Alle nennenswerten Änderungen an diesem Projekt werden in dieser Datei
dokumentiert. Das Format folgt [Keep a Changelog](https://keepachangelog.com/de/1.1.0/),
die Versionierung folgt [SemVer](https://semver.org/lang/de/).

## [Unreleased]

Noch kein getaggtes Release; das GHCR-Image folgt. Das Lastenheft steht bei
0.1.0; die folgenden Inkremente sind im Repo abgeschlossen.

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
