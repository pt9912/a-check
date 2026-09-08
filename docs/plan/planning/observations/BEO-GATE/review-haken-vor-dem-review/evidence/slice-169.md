**Vorgang:** slice-169

**Fund:** Der Commit `0565ca9` setzte `- [x] Unabhängiger Review durchgeführt`, ohne dass ein
Review stattgefunden hatte; ein Report entstand erst zwei Tage später. Der anschließende Review
war merge-blockierend (zwei HIGH) und fand unter anderem, dass die Korpus-Kontrolle des Slice
eine Obermenge ihres Gegenstands zählte — der Haken hätte einen Slice nach `done/` gelassen,
dessen Kernzusage nicht trug.

**Wie es auffiel:** beim Aufnehmen der Arbeit am Slice, durch Nachzählen der DoD gegen
`docs/reviews/`. Nicht durch einen Lauf.

**Gemessen:** Von den 4 Slices im lebenden Bestand mit der Phrase „unabhängiger Review" war
slice-169 der einzige ohne Report. Geltungsbereich: die 208 Stubs unter `done/wellenlos/` tragen
keine DoD mehr und fallen aus der Messung — für sie sagt sie nichts.
