**Vorgang:** slice-173
**Fund:** Die Lücke ist erstmals **gemessen** statt beobachtet.
`.claude/rules/modul-05-planning-harness.md` wurde auf einen entfernten Baseline-Stand umgebogen
(`../../.harness/baseline/v6.1.0/regelwerk/modul-05-planning-harness.md`, Ziel existiert seit
[slice-172](../../../../done/wellenlos/slice-172-baseline-v600-entfernen.md) nicht mehr) — `make
doc-check` meldete **0 Befunde**, Exit 0. Der Symlink zeigte ins Leere, und das vollständige
Doku-Gate schwieg.

Damit ist die Aussage des Eintrags („kein Gate hätte den Bruch gefangen") nicht mehr nur plausibel,
sondern belegt. Sie gilt auch, nachdem `versions` in diesem Slice die Baseline-Pins übernommen hat:
das neue Muster liest **Dateiinhalt**; ein Symlink hat für `d-check` keinen Linkpfad, es liest den
**Zielinhalt**.

**Dritte Instanz** (slice-142 · slice-167 · dieser Slice) — Schwelle erreicht, und damit keine
Notiz mehr. Verkörpert im selben Slice, weil dieselbe Schicht und derselbe Gegenstand betroffen
sind: `make symlink-check`, im `gates`-Aggregat.

**Nachtrag aus dem unabhängigen Review dieses Slice:** Die hier gemessene Form ist **nicht** die
von slice-167. Dort zeigten die Symlinks auf `v6.0.0`, das während der Migration daneben vendored
blieb — sie lösten auf. Ein Sensor, der nur *„Ziel existiert nicht"* prüft (die erste Fassung),
hätte jene Instanz grün gemeldet. Der Sensor prüft darum zusätzlich, ob ein Baseline-Ziel den
**adoptierten** Stand trägt. Ohne diesen Nachtrag hätte der Eintrag als verkörpert gegolten,
während eine der drei gezählten Instanzen weiter durchgefallen wäre.
