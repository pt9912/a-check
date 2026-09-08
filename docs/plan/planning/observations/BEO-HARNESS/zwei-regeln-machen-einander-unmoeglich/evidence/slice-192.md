**Vorgang:** slice-192

**Fund:** [`harness/conventions.md`](../../../../../../../harness/conventions.md) §Baseline trägt
zwei Zusagen: *genau **ein** Stand liegt vendored* und *aufgelöste Einträge sind ausgenommen — ihr
Zeiger **bleibt** auf dem Stand, gegen den sie damals formuliert wurden*. Beide zusammen ergeben
einen garantiert toten Link bei **jedem** Baseline-Sprung.

**Gemessen:** Zwei aufgelöste Adaptionen —
[`MR-011`](../../../../../../../harness/conventions.md#mr-011) und
[`MR-021`](../../../../../../../harness/conventions.md#mr-021) — trugen ihren
`Ersetzt-Baseline-Regel`-Zeiger als Markdown-Link in den vendorten Baum. Mit dem Entfernen des
alten Standes wären beide rot geworden.

**Warum es niemand sah:** Der Abschnitt erklärt die Kollision ausdrücklich für *„bezahlt"* und
belegt das mit den **null Nachzügen** des vorigen Sprungs. Diese Messung galt den vier
einfrierenden Ziel-Formen (Review-Report, Ergebnisnotiz, zwei Archiv-Stubs) —
`harness/conventions/done/` war nie darin, und der Satz nennt seinen Geltungsbereich nicht. Eine
Erfolgsmeldung ohne Datei-Klasse deckt beim nächsten Mal etwas, das sie nie angesehen hat.

**Behoben** durch dieselbe Antwort, die für jene vier gilt: Kennung statt Adresse. Der genannte
Stand bleibt derselbe, nur die Adresse fällt weg — strikt weniger Eingriff als ein mitwandernder
Zeiger, und die Aussage des Eintrags bleibt wortgleich.
