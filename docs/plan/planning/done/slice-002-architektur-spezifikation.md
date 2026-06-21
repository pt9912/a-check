# Slice slice-002: Technik- und Sicht-Stratum (Spezifikation + Architektur)

**Status:** done.

**Welle:** welle-02-spec.

**Bezug:** [`AC-FA-CONF-001`](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml),
[`AC-FA-EXTRACT-001`](../../../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion),
[`AC-FA-CLI-001`](../../../../spec/lastenheft.md#ac-fa-cli-001--aufruf-scan-wurzel-und-exit-codes),
[`AC-FA-DIST-001`](../../../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk),
[`AC-QA-01`](../../../../spec/lastenheft.md#ac-qa-01--determinismus).

**Autor:** pt9912. **Datum:** 2026-06-21.

> Retroaktiv angelegt (Planning-Harness-Nachzug, Regelwerk Modul 5).

---

## 1. Ziel

Die Spec-Stratifizierung ist vollständig: ein Technik-Stratum
(`spec/spezifikation.md`) präzisiert das Lastenheft, ein Sicht-Stratum
(`spec/architecture.md`) visualisiert die Komponenten — beide sprach-/
meilensteinfrei.

## 2. Definition of Done

- [x] `spec/spezifikation.md` mit [`SPEC-CONF-001`](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema) u. a. (`SPEC-EXTRACT/RULE/CLI/DET/DIST-001`) — präzisiert, erweitert nie.
- [x] `spec/architecture.md` mit den Hexagon-Komponenten (`ARC-*`: Kern/Ports/Adapter/Composition-Root), Schicht-Richtung, Scan-Sequenz.
- [x] Strata + ID-Schemata in [`harness/conventions.md`](../../../../harness/conventions.md) deklariert ([`MR-004`](../../../../harness/conventions.md#mr-004--spezifikation-und-architektur-strata-und-id-schemata)) — keine stille Setzung.
- [x] Source-Precedence Rang 2–3 in [`AGENTS.md`](../../../../AGENTS.md)/[`harness/README.md`](../../../../harness/README.md) real + verlinkt.
- [x] ADR-`Schärft` aufwärts gefüllt ([ADR-0002](../../adr/0002-text-heuristische-extraktion.md)/[ADR-0003](../../adr/0003-config-modell-a-check-yml.md)/[ADR-0004](../../adr/0004-distribution-image-mk.md)); [ADR-0001](../../adr/0001-go-impl-sprache.md) bleibt `—`.
- [x] `make doc-check` grün.

## 3. Plan (vor Code)

Technik-Stratum aus den Fundament-ADRs ableiten (SPEC präzisiert AC), Sicht aus
SPEC visualisieren; Referenz-Richtung strikt aufwärts; keine Abwärts-Verweise
auf ADR/Slice (matrix-Gate).

## 4. Trigger

Fundament-ADRs (slice-001) `Accepted` → das Technik-Stratum konnte aus ihnen
formalisiert werden.

## 5. Closure-Trigger

Beide Strata angelegt, `Schärft` gefüllt, Review abgeschlossen, `make doc-check` grün.

## 6. Risiken und offene Punkte

ID-Linkpflicht für `SPEC-*`/`ARC-*` zunächst noch nicht im Gate (im
Adaptions-Block als Nachzug deklariert) — mit slice-003 in `.d-check.yml` aktiviert.

## 7. Closure-Notiz (nach `done/`)

**Belege:** [Review](../../../reviews/2026-06-21-slice-002-spec-architektur.md)
(0 HIGH; 4 MEDIUM behoben), `make doc-check` grün.

**Lerneintrag (Steering-Loop):**

- *Geschärfte Regel:* Ein Spec-Dokument ohne deklariertes Stratum/ID-Schema ist
  eine stille Setzung — [`MR-004`](../../../../harness/conventions.md#mr-004--spezifikation-und-architektur-strata-und-id-schemata) deklariert Technik/Sicht + `SPEC-*`/`ARC-*`
  explizit, bevor sie normativ zitiert werden.
- *Benannte Spec-Lücke (aus dem Review):* Die Regel `port-impurity` verlangte
  „verbotene Konstrukte", aber [`SPEC-CONF-001`](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema)
  hatte keinen Config-Schlüssel dafür — im Review als Lücke erkannt und um
  `forbidden_constructs` geschlossen (Eingang in slice-003).

## 8. Sub-Area-Modus-Begründung

### Sub-Area: Spec-Schreibung (Technik + Sicht)

- **Modus:** GF
- **Konventionen-Dichte:** hoch (`harness/conventions.md` `MR-*`-Adaptionen).
- **Phase-Reife:** Phase 4 — Strata kohärent, von außen zitierbar.
- **Evidenz-/Diskrepanz-Risiko:** niedrig (Greenfield).
- **Reconciliation-Aufwand:** keiner.
