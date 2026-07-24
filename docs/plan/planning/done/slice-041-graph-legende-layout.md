# slice-041 — Graph-Legende: lesbares Zeilen-Layout

**Status:** **done** (2026-07-24) — reine Präsentation, kein Vertrag/ADR/Spec-Bump.
**Auslöser:** Nutzer-Feedback am gerenderten Graphen — der Legenden-Block war unübersichtlich: die lange Regel-Zeile lief an Mermaids maximale Knotenbreite und brach **mitten im Wort** um (`port-<br>direction-mismatch`).
**Bezug:** [AC-FA-CLI-002](../../../../spec/lastenheft.md#ac-fa-cli-002--architektur-graph-ausgabe) / [SPEC-CLI-002](../../../../spec/spezifikation.md#spec-cli-002--graph-renderer-vertrag) (Legende als Notiz — Layout ist nicht spec-gepinnt). Folge zu [slice-040](slice-040-graph-legende-vertical-slice-regeln.md).

---

## 1. Umsetzung

`graph.go` `writeLegend`: statt zweier langer, auto-umbrechender Zeilen jetzt **eine Regel pro Zeile**
(explizite `<br/>`), Kopf in zwei Zeilen, eine **Leerzeile** (`<br/><br/>`) als Trenner zu den
Stil-Hinweisen. Kurze Zeilen ⇒ Mermaid bricht **nie mitten im Wort** um.

```text
Legende — implizite Regeln
(nie als Kante gezeichnet):
core-impurity
lateral-adapter
lateral-slice
port-direction-mismatch
port-locality

durchgezogen = edges
gestrichelt = allow
Farbe = effektive Rolle
```

**Kein** Verhaltens-/Vertragswechsel: dieselben fünf kategorischen Regeln + Stil-Hinweise, nur das
Zeilen-Layout des einen Legenden-Knotens; weiterhin **pur**/deterministisch ([SPEC-CLI-002](../../../../spec/spezifikation.md#spec-cli-002--graph-renderer-vertrag)).

**Nachtrag (Nutzer-Feedback: zentriert statt linksbündig):** Mermaid zentriert Knoten-Text per Default;
ein portabel zuverlässiges „linksbündig" gibt es nicht. Der Legenden-Text ist jetzt in ein
`<div style='text-align:left'>` gewickelt (renderer-feste HTML-Ebene wie die `<br/>`, **nicht** durch den
Escaper). Das richtet in **freizügigen** Renderern (mermaid.live, VS Code, meist GitHub) linksbündig aus
und wird von **strengen** Sanitizern höchstens ignoriert (fällt dann auf zentriert zurück — harmlos, keine
Syntax). `image-test` (2b) bestätigt weiter dekodierbares Mermaid.

## 2. Definition of Done

- [x] `graph.go` `writeLegend` umgestellt; `TestRenderLegendListsCategoricalRules` grün (fünf Regeln weiter vorhanden).
- [x] `make gates`/`make ci` grün; `--print-graph` gegen das HexSlice-Beispiel gesichtet (Block lesbar, kein Wort-Umbruch).

## 3. Closure-Notiz

_(reine Legenden-Formatierung; noch unveröffentlicht — folgt mit dem nächsten Release.)_
