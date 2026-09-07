**Vorgang:** slice-174
**Fund:** Die Delta-Analyse bereinigte den Roh-Diff um **eine** Formatierungs-Klasse
(Tabellen-Trennzeilen, 36 Zeilen, korrekt gezählt) und leitete daraus die Umfangs-Beurteilung ab:
„8 Dateien tragen 436 von 536 substanziellen Zeilen, 9 tragen null". Der Kurs hatte in `v6.5.0`
zusätzlich die **Zell-Innenabstände** ganzer Tabellen normalisiert. Mit `git diff --numstat -w`
gemessen: **16** Dateien, **+452/−29** — 84 weitere `+`-Zeilen sind inhaltsgleich, und **14**
statt 9 Dateien haben null Inhaltsänderung.

**Warum es dieselbe Klasse ist:** Der Geltungsbereich der Messung war begründet — die
Trennzeilen-Bereinigung ist richtig und wurde belegt — und trotzdem zu eng gezogen; die zweite
Klasse wurde nicht gesucht. Gefunden hat es der unabhängige Review, nicht der Messende.

**Verschärfend, und darum steht es hier statt in einer neuen Beobachtung:** Der **Lerneintrag
desselben Slice** lautet *„Ein Roh-Diff wird vor der Umfangs-Beurteilung um seine formatierenden
Anteile bereinigt; die Bereinigung ist zu messen, nicht zu schätzen."* Die Regel war richtig
formuliert und halb ausgeführt. Ein Lerneintrag schützt nicht davor, im selben Slice gegen ihn zu
verstoßen.

**Folgen im Bestand, nicht nur in der Zahl:** Fünf Dateien standen unter §3.3 als „Tabellen-Ausbau
und Verweise, zu prüfen im Adaptions-Durchgang" — darunter
`grundlagen-referenz-richtung.md` und `grundlagen-source-precedence.md`, die die
`Ersetzt-Baseline-Regel`-Anker von [`MR-011`](../../../../../../../harness/conventions.md#mr-011) und [`MR-012`](../../../../../../../harness/conventions.md#mr-012) tragen. Sie sind **unverändert**; der
Durchgang hätte dort gegen nichts geprüft.

**Dritte Instanz** (slice-105 · slice-172 · dieser Slice) — Schwelle erreicht.
