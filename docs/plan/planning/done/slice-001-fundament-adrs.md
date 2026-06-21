# Slice slice-001: Fundament-ADRs — Sprache, Extraktion, Config, Distribution

**Status:** done.

**Welle:** welle-01-fundament.

**Bezug:** [`AC-QA-01`](../../../../spec/lastenheft.md#ac-qa-01--determinismus),
[`AC-QA-02`](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze),
[`AC-QA-03`](../../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit),
[`AC-FA-EXTRACT-001`](../../../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion),
[`AC-FA-CONF-001`](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml),
[`AC-FA-DIST-001`](../../../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk).

**Autor:** pt9912. **Datum:** 2026-06-21.

> Retroaktiv angelegt (Planning-Harness-Nachzug, Regelwerk Modul 5): die
> inhaltliche Arbeit war abgeschlossen, die Slice-Datei fehlte. Closure-Belege
> verweisen auf die real entstandenen Artefakte.

---

## 1. Ziel

Die vier Fundament-Entscheidungen liegen als akzeptierte ADRs vor:
Implementierungssprache, Import-Extraktions-Ansatz, Config-Modell und
Distributionsweg.

## 2. Definition of Done

- [x] [ADR-0001](../../adr/0001-go-impl-sprache.md) Go als Implementierungssprache: `Accepted`, ≥ 3 Alternativen abgewogen.
- [x] [ADR-0002](../../adr/0002-text-heuristische-extraktion.md) text-heuristische Extraktion (kein AST): `Accepted`.
- [x] [ADR-0003](../../adr/0003-config-modell-a-check-yml.md) Config-Modell `.a-check.yml` (strict-decode): `Accepted`.
- [x] [ADR-0004](../../adr/0004-distribution-image-mk.md) Distribution (Image + `--print-mk`/`a-check.mk`): `Accepted`.
- [x] [ADR-Index](../../adr/README.md) aktualisiert; jede ADR trägt `Bezug`/`Schärft` + Fitness-Function-Anker.
- [x] `make doc-check` grün.

## 3. Plan (vor Code)

MADR je Entscheidung mit Optionen/Trade-offs, `Bezug` aufwärts auf die AC,
`Schärft` aufwärts auf das (noch entstehende) Technik-Stratum, je ein
Fitness-Function-Anker für slice-003.

## 4. Trigger

Harness-Bootstrap abgeschlossen; die technische Basis musste vor der
Spezifikation (slice-002) entschieden werden.

## 5. Closure-Trigger

Alle vier ADRs `Accepted`, Review + Re-Review abgeschlossen, `make doc-check` grün.

## 6. Risiken und offene Punkte

`Schärft` blieb bis slice-002 (`spec/spezifikation.md`) offen;
[ADR-0001](../../adr/0001-go-impl-sprache.md) koppelt sprachneutral an keine
Spec-§ (bewusst `—`).

## 7. Closure-Notiz (nach `done/`)

**Belege:** [Review](../../../reviews/2026-06-21-adr-fundament-slice-001.md) +
[Re-Review](../../../reviews/2026-06-21-adr-fundament-slice-001-rereview.md)
(0 HIGH), `make doc-check` grün.

**Lerneintrag (Steering-Loop):**

- *Geschärfte Regel:* Eine ADR über die **Implementierungssprache** koppelt an
  **keine** Spec-§, weil die Spezifikation sprachneutral ist
  ([`AGENTS.md`](../../../../AGENTS.md) §3.4) — `Schärft: —` ist hier korrekt,
  nicht eine Lücke. Diese Einordnung wurde in den ADR-Index übernommen.
- *Neuer Sensor/Prozess:* Faktenbehauptungen über das Schwester-Tool `d-check`
  wurden im Review überzogen (z. B. „strict-decode konsistent mit d-check");
  daraus wurde die Reviewer-Disziplin **adversarische Verifikation gegen
  Repo-Artefakte** ([`.harness/skills/reviewer.md`](../../../../.harness/skills/reviewer.md)).

## 8. Sub-Area-Modus-Begründung

### Sub-Area: Spec-/ADR-Schreibung

- **Modus:** GF
- **Konventionen-Dichte:** hoch (ADR-Vorlage + `harness/conventions.md` `MR-*`-Adaptionen).
- **Phase-Reife:** Phase 4 — Doc führt, Code folgt.
- **Evidenz-/Diskrepanz-Risiko:** niedrig (Greenfield).
- **Reconciliation-Aufwand:** keiner.
