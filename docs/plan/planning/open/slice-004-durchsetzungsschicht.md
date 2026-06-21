# slice-004 — Durchsetzungsschicht: gate-consistency + record-gates

**Status:** open (Backlog; wartet auf Trigger/Priorisierung).
**Welle:** welle-04-durchsetzungsschicht.
**Bezug:** [`AGENTS.md`](../../../../AGENTS.md) §3.6 (Gate-Disziplin) + §4 (`make gates`);
[ADR-0005](../../adr/0005-lint-profil.md)/[ADR-0006](../../adr/0006-coverage-gate.md)
(bestehende Gates); Stack-Vorbild `d-check` (`gate-consistency`/`record-gates`).

## Ziel

Die im Review zu [slice-003](../../../reviews/2026-06-21-slice-003-impl-gates.md)
(Befund B-M4) aufgedeckte Stack-Paritäts-Lücke schließen — die zwei
computational-Bindepunkte der Durchsetzungsschicht (Regelwerk-Grundlagen
§Durchsetzungsschicht), die a-check ggü. `d-check` noch fehlen:

- **gate-consistency** (Meta-Gate, *computational feedback*): prüft maschinell,
  dass die in [`AGENTS.md`](../../../../AGENTS.md) §4 und
  [`harness/README.md`](../../../../harness/README.md) §Sensors dokumentierten
  Targets ↔ Makefile übereinstimmen (Schutz gegen Doku-/Gate-Drift — genau die
  Klasse, die im slice-003-Review als Reviewer-Halluzination *und* als reale
  „real/geplant"-Drift auftrat).
- **record-gates** (Handoff-Gate-Nachweis): inhaltsbasierter
  Working-Tree-Hash-Nachweis, dass die Gates auf genau diesem Stand liefen —
  Grundlage für einen `.claude`-Stop-Hook; fail-closed, loop-guarded,
  bootstrap-aware (Regelwerk §Durchsetzungsschicht, vier Design-Eigenschaften).

## Definition of Done

- `make gate-consistency` existiert und bricht bei Drift zwischen dokumentierten
  Targets und Makefile ab.
- `make record-gates` erzeugt den Nachweis nach grünen Gates; als letzter
  Schritt in `make gates` eingehängt (Reihenfolge via `.NOTPARALLEL` gesichert).
- Nachweis-Ablage `.harness/state/` (bereits in `.gitignore` vorgesehen).
- Optional: `.claude`-Stop-Hook, der den Nachweis vor „fertig" erzwingt.
- [`harness/README.md`](../../../../harness/README.md) §Sensors +
  [`AGENTS.md`](../../../../AGENTS.md) §4 um die neuen Gates ergänzt.
- Beleg: `make gates` grün inkl. der neuen Gates.

## Offene Fragen

- `record-gates`: eigenes Skript unter `tools/` vs. `.claude`-Hook-Verdrahtung —
  Reihenfolge der Einführung.
- `gate-consistency`: Markdown-Parser für die Sensors-Tabellen vs. einfache
  Target-Listen-Diff.
- Ob ein eigener `MR-*`-Adaptionseintrag nötig ist (analog d-checks
  Parallelitäts-Regel für den Nachweis).
