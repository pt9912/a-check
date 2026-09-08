**Vorgang:** slice-169

**Zwei Ausprägungen im selben Vorgang** — nach der Zählregel *ein Vorgang zählt einmal* ein
Beleg, nicht zwei:

1. **Review-Haken ohne Review.** Der Commit `0565ca9` setzte `- [x] Unabhängiger Review
   durchgeführt`, ohne dass ein Review stattgefunden hatte; ein Report entstand erst zwei Tage
   später. Der nachgeholte Review war merge-blockierend (2 HIGH, 8 MEDIUM, 4 LOW) und fand unter
   anderem, dass die Korpus-Kontrolle des Slice eine Obermenge ihres Gegenstands zählte — der
   Haken hätte einen Slice nach `done/` gelassen, dessen Kernzusage nicht trug.
2. **Paarungs-Zeile vor dem `git mv`.** Die Closure-Notiz trug „**Drei Paarungen** (Repo ohne
   Wellen-Betrieb, **nach dem `git mv` geprüft**)", während die Datei in `in-progress/` lag. Die
   drei Ergebnisse selbst trafen zu; falsch war die Zeit-Aussage — und die ist der Beleg-Satz
   der Paarung.

**Wie es auffiel:** (1) beim Nachzählen der DoD gegen `docs/reviews/`, (2) durch die
kontext-getrennte Prüfung der Closure-Notiz. Beide Male nicht durch einen Lauf.

**Gemessen (1):** Von den 4 Slices im lebenden Bestand mit der Phrase „unabhängiger Review" war
slice-169 der einzige ohne Report. Geltungsbereich: die 208 Stubs unter `done/wellenlos/` tragen
keine DoD mehr und fallen aus der Messung — für sie sagt sie nichts.
