**Vorgang:** slice-184

**Fund:** Der Slice legte `make doc-mentions` an und hängte es ins `gates`-Aggregat. Die
abschließende Aufzählung in `AGENTS.md` §4 nannte es nicht — und die Zeile in
`harness/README.md` sagte gleichzeitig „im `gates`-Aggregat". Zwei Aussagen derselben Datei-Familie
widersprachen sich.

**Beim Nachmessen kam ein zweiter Fehler ans Licht**, älter als dieser Slice: Der Absatz führte
`doc-immutable` als „in `gates`", obwohl es dort **nicht** hängt (0 Treffer in der `gates`-Zeile
des Makefile) — es ist CI-durchgesetzt über die Commit-Range, was etwas anderes ist. Und
`doc-reviews` fehlte seit slice-160, dem Slice, der es anlegte.

**Gemessen nach der Korrektur:** genannte Menge und reale Menge sind deckungsgleich, acht Targets,
Differenz null.

**Wie es auffiel:** unabhängiger Review (F-1). Kein Lauf deckt die Stelle.
