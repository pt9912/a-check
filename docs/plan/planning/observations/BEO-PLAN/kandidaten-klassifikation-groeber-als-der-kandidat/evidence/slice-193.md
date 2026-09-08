**Vorgang:** slice-193

**Fund:** Die Differenz-Messung *„welche Targets stehen nur in `AGENTS.md` §4?"* las je
Tabellenzeile **ein** `make X` und meldete **acht**. Der Gate-Index führt aber eine **Sammelzelle**
mit vier advisory-Targets (`doc-trace` · `doc-doctor` · `doc-usage` · `doc-help`), und der Zähler
sah davon nur das erste. Die Klassifikation war gröber als der Kandidat — hier an der **Messung**
selbst, nicht an ihrer Beschreibung.

**Was die falsche Zahl gekostet hätte:** Drei Targets hätten eine überflüssige Zeile bekommen. Und
zwei — `doc-commits`, `doc-tracked` — hätten **keine**, obwohl sie eine brauchen: Sie standen im
Index nur in **Prosa**, und für den `targets`-Sensor ist Prosa kein Eintrag. Mit dem Wegfall der
zweiten Tabelle wären beide `gate-undocumented` geworden — ein rotes Gate direkt nach dem Umzug.

**Richtig gezählt** (alle `make X` je Zelle, zusätzlich gegen die Prosa): **drei** fehlen ganz,
**zwei** stehen nur in Prosa, **fünf** brauchen eine Zeile, **null** stehen nur im Index.

**Drittes Mal dieselbe Klasse.** Aufgefallen ist sie diesmal **beim Messen selbst** — die zweite
Zählung entstand aus der Frage, warum `doc-doctor` und `doc-help` in der Liste standen, obwohl der
Index sie sichtbar führt. Das ist die Gegenprobe am Fund, die
[`muster-trifft-nur-die-haeufige-schreibweise`](../../../BEO-GATE/muster-trifft-nur-die-haeufige-schreibweise/observation.md)
als billigstes Mittel nennt.

**Nachtrag aus dem Review:** Die Closure hatte den Ausgang zunächst auf *„den nächsten
Lese-Schritt"* vertagt. In einem Repo ohne Wellen-Betrieb **ist** die Slice-Closure der
Lese-Schritt; der Eintrag hätte sie ohne Ausgang überstanden, und `make verify` hätte es nicht
gemeldet — `verify-observations` prüft Deckung, nicht Ausgänge.
