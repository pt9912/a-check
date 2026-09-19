# Ein lebendes Dokument führt zwei Vokabulare für dieselbe Rolle

**Sub-Area:** Benutzer-Doku

Eine Rolle trägt **ein** Vokabular, und das Dokument, das ihre Benutzung erklärt, führt daneben ein
zweites — aus einer früheren Fassung, die dieselbe Rolle anders benannte. Beide Wörter stehen
lesbar nebeneinander; der Leser wählt das falsche und bekommt vom Werkzeug eine Fehlermeldung, die
ihm nicht sagt, woher der Widerspruch kommt.

**Warum das nicht bloß eine Ungenauigkeit ist:** Das Dokument ist für die Rolle die *einzige*
Anleitung. Wer `internal/adapters/{inbound,outbound}/…` als Beispielstruktur liest und danach
baut, hat eine Struktur, die die Regel-Tabelle derselben Datei später anders benennt — und die
`direction`-Werte, die er dort einträgt, sind dann die falschen. Der Widerspruch liegt **im selben
Dokument**, oft nur Abschnitte auseinander.

**Woran es auffällt:** an einem Leser, der beide Stellen sieht — nicht an einem Lauf. Ein
Vokabular ist Prosa; die Konsistenz zweier Benennungen desselben Gegenstands ist ein Urteil
([`AGENTS.md`](../../../../../../AGENTS.md) §3.7). Was mechanisch greift, ist ein **Zähler in
zwei verschieden gebauten Fassungen** (`grep` über die alte Wendung, `awk` über Zeilen mit `dapter`
**und** dem alten Wert) — und der trennt „eine Stelle" von „ein Muster".

**Abgrenzung zu [`BEO-KERN/dirvocab-portfor-auseinander`](../../BEO-KERN/dirvocab-portfor-auseinander/observation.md):**
Dort laufen zwei **Code**-Pakete auseinander, die dieselbe Zuordnung führen; hier widerspricht sich
ein **Text** selbst. Die eine Klasse ist maschinell bewachbar, die andere nicht.
