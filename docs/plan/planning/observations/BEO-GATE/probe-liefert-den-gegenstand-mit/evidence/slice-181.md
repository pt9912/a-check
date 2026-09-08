**Vorgang:** slice-181

**Fund:** Die erste Mutations-Probe für die neue Zellengrenze setzte ihren
Fülltext an das Zeilenende — also **hinter** das schließende `|` der Tabellenzeile und damit
außerhalb der Zelle. Der Lauf meldete grün, und für einen Moment sah es so aus, als griffe die
Konfiguration auf `harness/README.md` nicht.

**Wie es auffiel:** in derselben Sitzung, weil das Ergebnis nicht zur Erwartung passte — die
Probe *sollte* rot sein. Das ist der Unterschied zu slice-169 und slice-180, wo es erst der
unabhängige Review fand.

**Dritte Ausprägung derselben Klasse:** slice-169 mutierte das *Muster* statt der
Kandidatenmenge; slice-180 lieferte den Gegenstand mit (alle drei Verweis-Formen in einer Datei);
hier traf die Mutation daneben. Gemeinsam ist: Die Probe war grün, und ihr Grün bedeutete nichts.

**Mit diesem Beleg bei 3×** — verkörpert als Regel in `AGENTS.md` §5.
