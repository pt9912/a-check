**Stand:** verkörpert in [`make symlink-check`](../../../../../../Makefile)
(`tools/symlink-check.sh`, im `gates`-Aggregat) `seit slice-173`.

**Zwei** Prüfungen, weil die drei gezählten Instanzen zwei Formen haben — der Sensor prüfte
zunächst nur die erste, und der unabhängige Review hat das gefangen (F-2):

| Prüfung | Form | Instanz |
|---|---|---|
| Ziel existiert | der Stand wurde entfernt, der Symlink zeigt ins Leere | slice-173 |
| Baseline-Ziel trägt den adoptierten Stand | der alte Stand liegt noch daneben — der Symlink löst auf und ist trotzdem falsch | slice-167 |

Die zweite ist die wichtigere: `harness/conventions.md` §Baseline erlaubt zwei parallele Stände
**während einer Migration**, also genau im Fenster, das diese Beobachtung beschreibt. Ein Sensor
mit nur der ersten Prüfung hätte slice-167 grün gemeldet und diesen Eintrag fälschlich als
verkörpert ausgewiesen.

**Nicht** geprüft und damit offen geblieben: ob ein Symlink **außerhalb** der vendored Baseline auf
das inhaltlich richtige Ziel zeigt. Das wäre ein Urteil über Absicht, kein Match.
