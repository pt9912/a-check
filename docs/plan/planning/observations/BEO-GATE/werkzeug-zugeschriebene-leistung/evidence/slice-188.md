**Vorgang:** slice-188

**Fund:** F-11 des unabhängigen Reviews. Der Lifecycle-Commit `fd301de` (`open` →
`in-progress`) trägt den Betreff *„slice-188 open -> in-progress (make slice-mv)"*; derselbe
Commit streicht in `docs/plan/planning/in-progress/roadmap.md` den Ruhe-Marker
*„Nichts in Arbeit."*.

**Beleg, dass das Werkzeug ihn nicht anfasst:** `grep -n "Marker" tools/slice-mv.sh` findet
weder das Wort noch den Satz. Das Skript erklärt seine Grenze selbst (§NICHT BEHANDELT (1)
*SEMANTIK*: *„Das Werkzeug zieht PFADE nach, keine Aussagen"*).

**Zweites Vorkommen derselben Klasse.** Das erste steht in `evidence/slice-187.md` — dort als
Teil (b) desselben Fundes. Die beiden Vorkommen liegen in zwei aufeinanderfolgenden Vorgängen
und betreffen denselben Betreff-Typ.

**Nicht rückwirkend korrigierbar.** Ein Commit-Betreff steht in `git`; ihn zu ändern hieße,
die Historie umzuschreiben, und das wäre teurer als der Fehler. Der Vorgang ist darum
**benannt**, nicht behoben — was hier zählt, ist die Frage beim nächsten Lifecycle-Commit.
