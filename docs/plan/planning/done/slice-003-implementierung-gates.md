# Slice slice-003: Go-Implementierung + Quality-Gates

**Status:** done.

**Welle:** welle-03-implementierung.

**Bezug:** [`AC-FA-RULE-001`](../../../../spec/lastenheft.md#ac-fa-rule-001--kern-reinheit-regel-core-impurity),
[`AC-FA-RULE-002`](../../../../spec/lastenheft.md#ac-fa-rule-002--keine-lateralen-adapter-kanten-regel-lateral-adapter),
[`AC-FA-RULE-003`](../../../../spec/lastenheft.md#ac-fa-rule-003--tech-kapselung-regel-tech-leak),
[`AC-FA-RULE-004`](../../../../spec/lastenheft.md#ac-fa-rule-004--port-disziplin-regel-port-impurity),
[`AC-FA-RULE-005`](../../../../spec/lastenheft.md#ac-fa-rule-005--schicht-richtung-regel-wrong-direction),
[`AC-FA-EXTRACT-001`](../../../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion),
[`AC-FA-CONF-001`](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml),
[`AC-FA-CLI-001`](../../../../spec/lastenheft.md#ac-fa-cli-001--aufruf-scan-wurzel-und-exit-codes),
[`AC-FA-DIST-001`](../../../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk),
[`AC-QA-01`](../../../../spec/lastenheft.md#ac-qa-01--determinismus),
[`AC-QA-02`](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze),
[`AC-QA-03`](../../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit).

**Autor:** pt9912. **Datum:** 2026-06-21.

> Retroaktiv angelegt (Planning-Harness-Nachzug, Regelwerk Modul 5).

---

## 1. Ziel

Das a-check-Binary (Go 1.26, hexagonaler Schnitt: Kern/Ports/Adapter/Composition-Root) prüft fremde Repos
gegen die fünf Regeln; alle inneren Gates sind real und grün, distroless/static
verteilt, mit Eigen-Architektur-Dogfooding.

## 2. Definition of Done

- [x] Go-Implementierung: Kern ([`ARC-001`](../../../../spec/architecture.md)) + Adapter + CLI; fünf Regeln nach [`SPEC-RULE-001`](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung), Extraktion nach [`SPEC-EXTRACT-001`](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion).
- [x] strict-decode `.a-check.yml`, Exit-Codes 0/1/2, `--print-config`/`--print-mk`.
- [x] Multi-Stage-Dockerfile (static/distroless, digest-gepinnte Bases, [`AC-QA-02`](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)/[`AC-QA-03`](../../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit)).
- [x] Gates `make lint`/`test`/`coverage-gate`/`arch-check`/`gates` real; Coverage ≥ 90 % ([ADR-0006](../../adr/0006-coverage-gate.md), Ist 92,60 %), Lint-Profil ([ADR-0005](../../adr/0005-lint-profil.md)) `Accepted`.
- [x] Dogfooding: `make arch-check` prüft a-check selbst — 0 Befunde ([`AC-QA-02`](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)).
- [x] `make gates` grün (lint/test/coverage-gate/arch-check/doc-check).

## 3. Plan (vor Code)

Hexagon-Pakete (`internal/core` rein, Adapter implementieren Ports, `cmd`/
`internal/cli` als Composition Root), Gates als Dockerfile-Stages (Muster
d-check/u-boot), Tests AC-gebunden.

## 4. Trigger

Spec-Straten (slice-002) standen; die Implementierung konnte spec-treu folgen.

## 5. Closure-Trigger

`make gates` grün, Review abgeschlossen, [ADR-0005](../../adr/0005-lint-profil.md)/[ADR-0006](../../adr/0006-coverage-gate.md) `Accepted`.

## 6. Risiken und offene Punkte

Heuristik-Grenzen (Text-Extraktion, kein Parser) bewusst, [`AC-QA-02`](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)-ausgewiesen;
Stack-Paritäts-Meta-Gates offen → [slice-004](../open/slice-004-durchsetzungsschicht.md).

## 7. Closure-Notiz (nach `done/`)

**Belege:** [Review](../../../reviews/2026-06-21-slice-003-impl-gates.md)
(0 HIGH nach Verifikation), `make gates` grün — coverage-gate 92,60 %,
arch-check 0 Befunde, doc-check grün.

**Lerneintrag (Steering-Loop):**

- *Neuer Sensor:* zwei reale Maintainability-Gates eingeführt — Lint-Profil
  ([ADR-0005](../../adr/0005-lint-profil.md)) und Coverage-Gate 90 %
  ([ADR-0006](../../adr/0006-coverage-gate.md)); die zuvor als „geplant"
  geführten Gates wurden auf real promotet (Promotion-Trigger).
- *Geschärfte Regel:* Die Port-Interfaces ([`ARC-002`](../../../../spec/architecture.md))
  sind als Go-Interfaces **im Kern co-lokiert** (Go-Idiom) — die
  `port-impurity`-Eigenkollision aufgelöst; `spec/architecture.md` nachgezogen.
- *Benannte Spec-Lücke / Folge-Slice:* Stack-Paritäts-Meta-Gates
  `gate-consistency` (Doku ↔ Makefile) und `record-gates` (Working-Tree-Hash)
  fehlen → Backlog [slice-004](../open/slice-004-durchsetzungsschicht.md).
- *Prozess-Lücke (meta):* Bis hierher liefen die Slices **ohne
  Planning-Harness-Dateien** (Regelwerk-Verstoß gegen Modul 5 / [`AGENTS.md`](../../../../AGENTS.md)
  §5). Retroaktiv behoben (diese `done/`-Dateien mit Lerneinträgen);
  **Forward-Regel:** jeder Slice bekommt ab jetzt eine Datei und läuft den
  Lifecycle (slice-004 bereits in `open/`). Reviewer-Drift (halluzinierte
  Datei-Existenz) ist ein Kandidat für das künftige `gate-consistency`-Gate.

## 8. Sub-Area-Modus-Begründung

### Sub-Area: Implementierung (Go) + Gates

- **Modus:** GF
- **Konventionen-Dichte:** hoch (Spec-Straten + [ADR-0005](../../adr/0005-lint-profil.md)/[ADR-0006](../../adr/0006-coverage-gate.md) + `harness/conventions.md`).
- **Phase-Reife:** Phase 5 — Code an Spec gemessen, Gates erzwingen.
- **Evidenz-/Diskrepanz-Risiko:** niedrig (Doc führte, Code folgte; Dogfooding grün).
- **Reconciliation-Aufwand:** keiner; Folge-Arbeit als slice-004 (Backlog).
