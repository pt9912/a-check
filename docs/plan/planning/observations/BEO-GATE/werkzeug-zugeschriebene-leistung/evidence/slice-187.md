**Vorgang:** slice-187

**Nachzug.** Dieser Beleg wird bei der Closure von slice-188 geschrieben, nicht bei der von
slice-187. Er zitiert einen **bestehenden, unveränderlichen** Beleg — den Review-Report
`2026-09-08-slice-187-voll-abgleich-erstdurchgang-rest.md`, F-13, archiviert in
`done/wellenlos/slice-187-archiv.zip` — und erzählt nichts nach. Ohne ihn hätte die Klasse
ihren ersten Vorgang verloren, weil sie damals keinen Eintrag bekam.

**Fund:** F-13 trug zwei Stellen derselben Klasse:

- (a) `AGENTS.md` §3.3 sagt zum zweiten Lifecycle-Fall (*erst der Inhalt, dann der `git mv`*):
  *„`make slice-mv` fährt genau ihn"*. Das Werkzeug fährt nur den Move; die **Reihenfolge** ist
  eine Entscheidung des Laufs.
- (b) Der Commit `bd2ba58` trägt den Betreff `(make slice-mv)` und ändert daneben die
  **Roadmap** — die Streichung des Ruhe-Markers *„Nichts in Arbeit"*.

**Warum die Zuschreibung falsch ist:** `tools/slice-mv.sh` sagt in seinem Kopfkommentar
§NICHT BEHANDELT (1) *SEMANTIK* ausdrücklich, dass es **Pfade** nachzieht und **keine
Aussagen** — ein Zustandssatz ist genau das, was es nicht anfasst.

**Wie es auffiel:** im unabhängigen Review (slice-187), nicht im Lauf selbst — und **ohne**
Register-Eintrag, weshalb die Klasse beim nächsten Auftreten als erste Gelegenheit zählte.

**Folge:** (a) ist mit slice-187 behoben worden; (b) ist die Form, die slice-188 erneut zeigt
(`evidence/slice-188.md`).
