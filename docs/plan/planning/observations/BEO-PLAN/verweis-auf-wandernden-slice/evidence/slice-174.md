**Vorgang:** slice-174
**Fund:** Die Verkörperung dieser Beobachtung — `make slice-mv`, das die Verweise **auf** einen
wandernden Slice nachzieht — hat eine Form-Lücke, und sie ist mit `welle-15` erstmals aufgetreten.

Beim Übergang `in-progress/` → `done/` zog das Werkzeug **sechs** Verweise nach und ließ einen
siebten stehen: die **flach** unter `docs/plan/planning/` liegende Welle-Datei verweist als
`[slice-174](in-progress/slice-174-…md)` — mit Verzeichnis-Präfix, aber **ohne** `../`, weil sie
selbst eine Ebene höher liegt als die Lifecycle-Verzeichnisse.
[`AGENTS.md`](../../../../../../../AGENTS.md) §4 sagt zum Werkzeug „in beiden im Bestand
vorkommenden Formen"; diese dritte kam bis jetzt nicht vor — `welle-14` lag bereits in `done/`,
als ihre Slices wanderten, und wellenlose Slices haben gar keine Welle-Datei, die auf sie zeigt.

**Kein stiller Ausfall:** `make doc-check` meldete den toten Verweis (`target-missing`), die
Korrektur war ein `sed`. Der Zähler bewegt sich trotzdem — die Beobachtung sagt, dass Verweise auf
wandernde Slices brechen, und genau das ist passiert; dass ein zweiter Sensor es fing, macht die
Lücke im ersten nicht kleiner.

**Für den Ausgang:** Der Eintrag steht auf *verkörpert*. Diese Instanz widerlegt das nicht, sie
begrenzt es — die Verkörperung deckt zwei von drei Formen. Ob das Werkzeug die dritte lernt, ist
eine Entscheidung für die Closure von [`welle-15`](../../../../welle-15-regelwerk-v650-migration.md),
die noch drei weitere Slice-Übergänge fahren wird und die Form damit dreimal wiedersieht.
