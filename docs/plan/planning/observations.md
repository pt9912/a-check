# Beobachtungs-Register

Regeln dieses Registers: Baseline-Regelwerk `modul-06-roadmap.md`
§Das Beobachtungs-Register — wer schreibt, wer liest, wann gestrichen wird, welche Form ein Beleg
hat, und dass eine leere Tabelle `— keine —` trägt statt zu verschwinden.

**Wer schreibt:** die **Slice-Closure** — neue Kennung vergeben oder Zähler erhöhen und Beleg
ergänzen. Damit läuft der Zähler mit jedem geschlossenen Slice und nicht mit der Welle; in diesem
Repo ist das wesentlich, weil gerade keine Welle offen ist.
**Wer liest:** die Welle-Closure, was **3×** erreicht hat — und die Slice-Planung (§8 des Plans),
was darunter steht. Wer nur den ersten Schritt kennt, sieht alles unter 3× nie wieder an.

**Belege sind formgebunden:** Slice-Kennung `slice-<NNN>`, kein Freitext; so viele, wie der Zähler
sagt. Angelegt mit slice-101.

| Kennung | Beobachtung | Sub-Area | Zähler | Belege | Stand |
|---|---|---|---|---|---|
| BEO-001 | Die Harness-Einstiegs- und Konventions-Dateien haben keine Zeile in der Modus-Deklaration pro Sub-Area | Harness-Einstieg | 4× | slice-091, slice-095, slice-096, slice-097 | **verkörpert** in [`harness/conventions.md` §Modus-Deklaration pro Sub-Area](../../../harness/conventions.md#modus-deklaration-pro-sub-area) (`seit slice-101`) — Schwelle beim Erstauftreten im Register bereits überschritten |
| BEO-002 | „Höchstens zwei Schichten" der Größen-Regel bleibt maschinell ungeprüft; was eine Schicht ist, ist Ermessen über Modul-Grenzen | Planungs-Harness | 1× | slice-098 | offen — ein Zähler darüber wäre Schein-Genauigkeit; bleibt Sache des Reviews |
| BEO-003 | Ein Rückbau-Eintrag, der erklärt, dass **nicht** mehr abgewichen wird, liegt von Beginn an in `conventions/done/` statt in der aktiven Tabelle — Auslegung, keine Vorschrift | Harness-Einstieg | 1× | slice-097 | offen — überstimmbar; die Begründung steht in slice-097 §3 |
| BEO-004 | Ein Adaptions-Eintrag korrigiert eine Repo-Aussage statt einer Baseline-Regel und ist damit unter dem Fork-Test ein Rückbau-Kandidat | Harness-Einstieg | 1× | slice-097 | offen — auflösbar erst mit der Überarbeitung der ID-Schema-Deklaration; der Eintrag trägt den Trigger selbst |
| BEO-005 | Die Platzhalter-Liste von `verify-closure-notes` enthält eine Wendung, die **unzitiert** legitim in einem Risiko-Block steht | Gate-/Werkzeug-Schicht | 2× | slice-099, slice-100 | offen — Streichen wäre eine Schwellen-Senkung und braucht nach [`AGENTS.md`](../../../AGENTS.md) §3.6 eine ADR; entschieden wird am nächsten unzitierten Vorfall |
| BEO-006 | Die vorgeschriebene Reihenfolge kann die Prüfung nicht abnehmen, die sie erfüllen soll: `verify-closure-notes` greift nur in `done/`, der Workflow fährt `make verify` in Schritt 8 und den `git mv` erst in Schritt 9 | Gate-/Werkzeug-Schicht | 1× | slice-099 | offen — die Closure-Notiz eines Slice wird frühestens beim **nächsten** geprüft |
| BEO-007 | Der Risiko-Ausgangs-Sensor prüft **innerhalb** eines vorhandenen Blocks, nicht seine Existenz — wer den Block weglässt, wird nicht erwischt | Gate-/Werkzeug-Schicht | 1× | slice-102 | offen — die Existenz einzufordern verlangte ein Urteil (gab es hier überhaupt Risiken?) und erzeugte auf jedem älteren Slice Rauschen |
| BEO-008 | `verify-slice-links` nimmt `done/` als Endzustand aus — ein **done**-Slice, der auf einen noch wandernden zeigt, bleibt ungeprüft und bricht, wenn dieser wandert | Planungs-Harness | 1× | slice-093 | offen — die Invariante ist auf wandernde Slices gerichtet, nicht auf Verweise **auf** sie; gefangen hat es `doc-check`, nicht der zuständige Sensor |

## Gestrichene Einträge

Regeln dieser Sektion: Baseline-Regelwerk `modul-06-roadmap.md`
§Das Beobachtungs-Register — wer eine Zeile still löscht, macht sie ununterscheidbar von einer,
die es nie gab.

| Kennung | Beobachtung | Gestrichen am | Warum sie nicht mehr auftreten kann |
|---|---|---|---|
| — keine — | | | |
