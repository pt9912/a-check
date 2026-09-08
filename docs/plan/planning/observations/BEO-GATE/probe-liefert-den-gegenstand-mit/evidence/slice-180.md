**Vorgang:** slice-180

**Fund:** Die End-to-End-Probe legte eine flache Datei mit **allen drei** Verweis-Formen an und
belegte damit, dass alle drei nachgezogen werden. Tatsächlich brachten die zwei **bereits
gedeckten** Formen die Datei in die Kandidatenmenge; die neu ergänzte dritte wurde von der
Auswahl nie gefunden. Eine Datei mit ausschließlich Form 3 blieb unberührt — das Werkzeug meldete
„0 Datei(en) nachgezogen" mit Exit 0.

**Wie es auffiel:** unabhängiger Review, F-1. Er las die Kandidaten-Auswahl statt der Probe.

**Behandlung:** Auswahl und Ersetzung teilen jetzt eine Muster-Quelle (`match_patterns`), und der
Selbsttest fährt **beide** Schritte. Zweiter End-to-End-Anlauf mit einer Datei, die **nur** Form 3
trägt.
