**Vorgang:** slice-181

**Fund:** Die Vorab-Messung der Zellenlängen splittete Tabellenzeilen an `|` und übersah, dass
eine Zelle **escapte** Pipes tragen kann (`TO=<open\|next\|in-progress\|done>`). Dieselbe Zelle
erschien dadurch als mehrere kurze und blieb unter jeder Schwelle.

**Folgen, beide gemessen:** Die Zelle (`make slice-mv` in `AGENTS.md` §4) fehlte in der Liste der
aufzulösenden — **der Sensor fand sie, die Zählung nicht**. Und die Kalibrierungs-Zahlen des
Slice standen auf derselben Grundlage: „bei 200 wären es 17, bei 250 acht" — richtig sind **18**
und **9**, jeweils genau um diese eine Zelle daneben.

**Warum es hierher gehört:** Das Split-Muster war an den Zellen kalibriert, die es beim Schreiben
gab — keine davon trug einen escapten Pipe. Dieselbe Klasse wie beim `slice-mv`-Muster
(slice-180), das nur die zwei häufigen Verweis-Formen kannte.

**Wie es auffiel:** durch den Sensor selbst (er meldete eine Zelle, die die Zählung nicht hatte)
und durch den unabhängigen Review (F-1, die Zahlen).
