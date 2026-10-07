**Vorgang:** slice-212

**Fund (Review F-4/F-5):** Der erste Vertragsentwurf machte die Meldung einzeilig, indem er ein
Zeilenende als `\n` schrieb — ohne den Backslash selbst zu maskieren. Zwei verschiedene
Anweisungen hätten dieselbe Meldung getragen (`shape-differs` zeigte `X (erwartet: X)`, die
Zusammenfassung byte-gleicher Befunde verschluckte einen), und zwei Stellen der Spezifikation
widersprachen sich für `shape-unused`. Dazu D-2: der Zusatz ` (regex)` war von einem Literal mit
demselben Text nicht zu unterscheiden.

**Entschieden** in [SPEC-RULE-001](../../../../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung)
(Spezifikation 0.36.0): Meldung einzeilig **und** umkehrbar (`\\`, `\n`, `\r`), `shape-unused` mit
Präfix `literal: `/`regex: `.
