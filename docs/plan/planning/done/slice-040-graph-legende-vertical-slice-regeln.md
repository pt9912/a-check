# slice-040 — Graph-Legende: `lateral-slice` + `port-locality` ergänzen

**Status:** **done** (2026-07-24) — kleiner Nachzug zu [slice-039](slice-039-hexslice-vertical-slice-regeln.md); Renderer + Spec + Handbuch + Test.
**Auslöser:** [slice-039 §6](slice-039-hexslice-vertical-slice-regeln.md#6-offen--grenzen) offene Folge — die `--print-graph`-Legende nannte die zwei neuen kategorischen Regeln noch nicht.
**Bezug:** [AC-FA-CLI-002](../../../../spec/lastenheft.md#ac-fa-cli-002--architektur-graph-ausgabe) / [SPEC-CLI-002](../../../../spec/spezifikation.md#spec-cli-002--graph-renderer-vertrag); ehrliche Ausgabe [AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze). [Roadmap](../in-progress/roadmap.md).

---

## 1. Ziel & Umsetzung

Die Legende der Mermaid-Ausgabe (`writeLegend`) listete `core-impurity · lateral-adapter ·
port-direction-mismatch`. Mit [slice-039](slice-039-hexslice-vertical-slice-regeln.md) gibt es zwei
weitere **kategorische** Regeln — `lateral-slice` und `port-locality` — die als Legenden-Notiz gehören
(nicht als gezeichnete Kante, [AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze):
keine Semantik-Behauptung über den realen Code).

- **Code:** `graph.go` `writeLegend` — beide Regeln in die Legenden-Zeile (rein additiv, pure Funktion, deterministisch).
- **Spec:** [SPEC-CLI-002](../../../../spec/spezifikation.md#spec-cli-002--graph-renderer-vertrag) Legenden-Aufzählung, Spez **0.23.0**.
- **Handbuch:** §3.6 Legenden-Satz, **1.32**.
- **Test:** `TestRenderLegendListsCategoricalRules` (alle fünf kategorischen Regeln im Output).

**Kein** neuer Vertrag/ADR/Lastenheft-Bump (reiner Legenden-Nachzug zu bestehendem [AC-FA-CLI-002](../../../../spec/lastenheft.md#ac-fa-cli-002--architektur-graph-ausgabe)).

## 2. Definition of Done

- [x] `graph.go` Legende + Test grün; `make gates`/`make ci` grün; `arch-check` 0; Determinismus unberührt (image-test).
- [x] [SPEC-CLI-002](../../../../spec/spezifikation.md#spec-cli-002--graph-renderer-vertrag) + Handbuch §3.6 nachgezogen; Spez 0.23.0, Handbuch 1.32.

## 3. Closure-Notiz

_(Gate-Beleg beim Merge; reine Legenden-Erweiterung, kein Verhaltenswechsel am Scan.)_
