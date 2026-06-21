# Review-Report — slice-003 (Go-Implementierung + Quality-Gates)

- **Review-Art:** Code-/Build-Review vor ADR-Acceptance (ADR-0005/0006)
- **Gegenstand:** `internal/**`, `cmd/a-check`, Dockerfile, Makefile, `.golangci.yml`, `tools/coverage-gate.sh`, `.a-check.yml`, `.d-check.yml`, ADR-0005/0006, Sensor-Doku
- **Stand:** Commit `abe6122` (+ Doku `db62ba3`) · **Datum:** 2026-06-21 · **Modell:** Opus 4.8
- **Methode:** Drei unabhängige Reviewer-Agenten (frischer Kontext, perspektiven-divers: Go-Korrektheit · Gates/Build · Konventionen/Dogfooding), Synthese mit adversarischer Verifikation der HIGH-/Existenz-Befunde gegen die realen Artefakte (Modul 11).

## Adversarisch verifiziert / verworfen

| Gemeldet | Verifikation | Ergebnis |
|---|---|---|
| Reviewer-B M-1: ADR-0006 zitiert nicht-existenten `cli_test` | `find internal -name '*_test.go'` → `internal/cli/cli_test.go` u. a. existieren | **verworfen** (Halluzination); ADR-0006-Aussage ist korrekt |
| Reviewer-B M-2: `arch-check`-Dogfooding unbelegbar, keine `.a-check.yml` | `ls .a-check.yml` (635 B) vorhanden; `make arch-check` → 0 Befunde | **verworfen** (Halluzination) |

Steering-Hinweis: erneute Reviewer-Drift (Datei-Existenz halluziniert) — die Reviewer-Skill sollte den Datei-Kontext per `ls`/`find` verankern, bevor sie „existiert nicht" behauptet.

## Bestätigte Befunde + Disposition

| ID | Kat. | Befund | Aktion |
|---|---|---|---|
| C-B1 | HIGH | `architecture.md` ARC-002 führt Ports als eigene Schicht („importieren weder Kern noch Adapter"), Code co-lokiert sie aber im Kern; `.a-check.yml` hat keine `ports`-Schicht (Sicht ↔ Code-Inkonsistenz). | **behoben** — ARC-002 stellt klar: Ports sind Go-Interfaces **im Kern-Paket co-lokiert** (Go-Idiom), im Eigen-`.a-check.yml` Teil der `core`-Schicht; separates Ports-Paket ⇒ eigene `ports`-Schicht. |
| A-H1 | HIGH→MED | `ruleFor`-Erst-Treffer priorisiert `lateral-adapter` vor `tech-leak` ohne deklarierte Präzedenz. | **behoben** — SPEC-RULE-001 deklariert die deterministische Erst-Treffer-Reihenfolge explizit. |
| A-H2 | HIGH→MED | `composition_root`-Ausnahme wirkt datei-global (alle Regeln), Spec formulierte nur `tech-leak`. | **behoben** — SPEC-RULE-001 deklariert die Ausnahme als „alle Schicht-Regeln + tech-leak" (Verdrahtungspunkt). |
| A-L2 | LOW | `stripComments` behandelt nur Kommentare, nicht String-Literale; SPEC-EXTRACT-001 überzeichnete „String-Literale … nicht gewertet". | **behoben** — SPEC-EXTRACT-001 weist String-Literale als ehrliche Heuristik-Grenze (0.1.0) aus. |
| C-B4 | MED | `lint`/`coverage-gate` als „real & grün" geführt, ADR-0005/0006 aber `Proposed`. | **aufgelöst** durch das Acceptance-Sign-off dieses Laufs. |
| B-L1 | LOW | Dockerfile ohne `# syntax`-Frontend-Pin. | **behoben** — `# syntax=docker/dockerfile:1.7`. |
| B-M5 | LOW | kein `.NOTPARALLEL` (Reihenfolge-Garantie unter `make -j`). | **behoben** — `.NOTPARALLEL` ergänzt. |
| C-B3 | LOW | `sampleConfig` (`--print-config`) zeigt `ports`/`forbidden_constructs`, a-check selbst nicht. | **bewusst beibehalten** — generisches Schema-Gerüst für beliebige Konsumenten (die ein Ports-Paket haben können); `mkFragment` ist byte-konsistent zu `a-check.mk`. |
| A-M1..M3, A-L1/L3, C-B2 | LOW/INFO | Heuristik-Grenzen (`strings.Contains`-Pfad-/Senken-Auflösung, `adapterSeg`-Literal, Go-Block-Ende, Dogfooding deckt v. a. core-impurity + tech-leak). | **bewusst beibehalten** — AC-QA-02 (ehrliche Heuristik-Grenze); 0.1.0-Stand, dokumentiert. |
| B-M4, C-B5..B7 | LOW | Stack-Paritäts-Lücken: kein `gate-consistency`/`record-gates`-Meta-Gate; `.PHONY`/`.d-check.yml`-Kopfkommentar hinken hinterher. | **Follow-up-Backlog** — d-check-Parität (Meta-Gates) als eigener Slice; nicht acceptance-blockierend. |

## Negativbefunde (geprüft, ohne Befund)

- **Hexagon-Reinheit:** Kern importiert nur stdlib (Reviewer A I-2); kein Adapter importiert einen anderen.
- **Determinismus (AC-QA-01):** alle Ausgaben stabil sortiert; keine unsortierte Map-Iteration (Reviewer A).
- **Gate-Mechanik (AC-QA-02/03):** alle `FROM` digest-gepinnt (byte-identisch zu d-check), `-coverpkg`/`covermode=atomic`/`pipefail`, `--no-cache-filter` auf den Build-Stage-Gates, robustes `coverage-gate.sh`, THRESHOLD-Verdrahtung end-to-end (Reviewer B N-1..N-9).
- **ADR-0005/0006 ↔ Realität:** zitierte Werte korrekt (v2.12.2, 90 %, Ist 92,60 %, Exclusions, `-coverpkg`); fail-closed strict-decode; Coverage-Historie konsistent geführt (Reviewer B N-7/N-8, Reviewer C).

## Kategorie-Summary (nach Verifikation)

| Kategorie | Anzahl | Status |
|---|---|---|
| HIGH | 1 (C-B1) | behoben |
| MEDIUM | 3 (A-H1, A-H2, C-B4) | 2 behoben, 1 durch Acceptance aufgelöst |
| LOW | ~9 | 3 behoben, Rest bewusst beibehalten / Follow-up |
| verworfen (Halluzination) | 2 (B-M1, B-M2) | — |

## Verdikt

Kein offener blockierender Befund. Der eine reale HIGH (Ports-Sicht ↔ Code) ist behoben; die ADR-bezogenen Validierungen (ADR-0005 Lint-Profil, ADR-0006 Coverage-Gate) sind sauber — beide **acceptance-reif**. Stack-Paritäts-Lücken (`gate-consistency`/`record-gates`) bleiben als Follow-up-Backlog.

## Disposition-Beleg

`make gates` grün nach den Fixes; `make doc-check` grün inkl. dieses Reports.
