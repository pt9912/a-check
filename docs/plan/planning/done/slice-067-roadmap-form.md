# slice-067 — Die Roadmap ist eine Wellen-Sequenz, keine Slice-Chronik

**Status:** der Zustand ist das **Verzeichnis** dieser Datei
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5) — dieses Feld führt ihn bewusst **nicht** doppelt.
**Deckt:** Form-Verstoß der Roadmap gegen `modul-06` §Roadmap-Struktur, gemeldet vom Maintainer am
2026-07-26.
[Roadmap](../in-progress/roadmap.md).

---

## 1. Auslöser

`modul-06` gibt **fünf Abschnitte** vor (*Aktuelle Welle · Nächste Wellen · Meilensteine ·
Abgeschlossene Wellen · Historische Trigger-Verschiebungen*) und den Satz, der die Form trägt:
*„Eine Roadmap ist nicht ‚wann?', sondern ‚in welcher Reihenfolge wovon?'"*. Gemessen an 332
Zeilen:

| Stelle | Umfang | Verstoß |
|---|---|---|
| Kopf-Blockquote „Wiedereinstieg (Stand 2026-07-25)" | ~103 Z. | kein Pflicht-Abschnitt; Slice-für-Slice-Chronik von slice-042 bis slice-057 |
| `## Stand 2026-07-26` | 71 Z. | dito — **in slice-063/066 von mir selbst angelegt**, statt den vorhandenen Wildwuchs zu beheben |
| `## Aktuelle Welle` | 69 Z. | führt eine **abgeschlossene** Welle (`welle-10`), die zugleich im Closure-Log steht; **keiner** der drei Pflicht-Bestandteile (Slice-IDs · Trigger · Closure-Kriterien); ~60 Z. Fließtext-Chronik slice-016 … slice-036 |
| `## Nächste Wellen` | 11 Z. | enthält drei **abgeschlossene** Wellen (05, 06, 11); Spalte `Status` statt des vorgeschriebenen geschätzten Aufwands (S/M/L) |

**~243 von 332 Zeilen sind Slice-Chronik statt Wellen-Sequenz.** Konform sind *Meilensteine*,
*Abgeschlossene Wellen* und das *Drift-Log*.

**Die Ursache ist Duplikation.** Jede dieser Chroniken wiederholt, was in der Closure-Notiz des
jeweiligen Slice unter `done/` bereits steht — und driftet dann davon weg. Genau deshalb konnte
der Kopf-Block „Etappe F läuft" behaupten, während F abgeschlossen war: ich habe den
Abschluss-Absatz in einen Blockquote gesetzt, der **„Stand 2026-07-25"** überschrieben ist.

## 2. Betroffene Module

- [`docs/plan/planning/in-progress/roadmap.md`](../in-progress/roadmap.md) — die vier
  nicht-konformen Stellen.

**Eine Schicht** (Planungs-Doku). Kein Code, kein Vertrag, kein Sensor.

## 3. Auszuführende Gates

`make gates` und `make verify`, Ausgabe je in eine Datei, Exit-Code getrennt geprüft. Kein
Sensor für die Roadmap-Form — `make doc-planning` prüft Lifecycle-Konsistenz, nicht
Abschnitts-Struktur. Als Lücke benannt, nicht behauptet (§4).

## 4. Was bewusst nicht getan wird

- **Kein Informationsverlust durch Löschen.** Alles Slice-Bezogene aus den Chroniken steht bereits
  in den Closure-Notizen unter `done/`; die Roadmap **verweist** künftig darauf, statt es zu
  wiederholen. Was nur in der Roadmap stand — Wellen-Zuordnung, Meilensteine, Drift — bleibt.
- **Kein Sensor für die Abschnitts-Struktur.** Er wäre baubar (fünf Überschriften prüfen), aber
  die Form-Verstöße lagen nicht an fehlenden Überschriften, sondern am **Inhalt** der Abschnitte —
  das prüft kein Regex. Ein Gate, das die Überschriften zählt und den Wildwuchs darunter
  durchwinkt, wäre ein grünes Gate ohne Aussage.
- **Keine rückwirkende Wellen-Ergebnisnotiz.** Die Migration wandert als `welle-12` ins
  Closure-Log mit Verweis auf ihre Slices; eine `done/welle-12-results.md` entsteht **nicht** —
  die Fünf-Schritt-Prozedur gilt laut [slice-066](../done/slice-066-wellen-closure-und-rollen.md)
  ab der **nächsten** Welle.

## 5. DoD

- [x] Die Roadmap führt **genau** die fünf Pflicht-Abschnitte (plus den Abhängigkeitsgraphen als
      deklariertes Extra); die beiden Status-Blöcke sind entfernt.
- [x] *Aktuelle Welle* nennt die drei Pflicht-Bestandteile oder weist ausdrücklich aus, dass keine
      Welle läuft; *Nächste Wellen* enthält nur offene Wellen mit Trigger und Aufwand (S/M/L),
      abgeschlossene stehen ausschließlich im Closure-Log.
- [x] `make gates` und `make verify` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft,
      nie in eine Pipe.

## 6. Closure-Notiz

**Geliefert:** die Roadmap führt wieder **genau** die fünf Pflicht-Abschnitte plus den
Abhängigkeitsgraphen. Beide Status-Blöcke sind entfernt, *Aktuelle Welle* weist aus, dass keine
Welle läuft (mit den drei Pflicht-Bestandteilen für den Fall, dass eine gezogen wird),
*Nächste Wellen* enthält nur offene Wellen mit beobachtbarem Trigger und Aufwand (S/M/L). Die
Migration steht als `welle-12-regelwerk-migration` im Closure-Log. **332 → 109 Zeilen**, ohne
Informationsverlust: alles Slice-Bezogene stand bereits in den Closure-Notizen.

**Lerneintrag — Form: geschärfte Regel.**
> **Ein Dokument, das den Stand *wiederholt*, statt auf ihn zu *zeigen*, driftet — und zwar
> unbemerkt, weil beide Fassungen plausibel aussehen.** Die Roadmap trug zuletzt ~243 von 332
> Zeilen Slice-Chronik, die jede Closure-Notiz ein zweites Mal erzählte. Der Beweis für die Drift
> steht in meinem eigenen Beitrag: ich habe einen Absatz „**Etappe F abgeschlossen (2026-07-26)**"
> in einen Blockquote gesetzt, der „**Wiedereinstieg (Stand 2026-07-25)**" überschrieben ist —
> und zwei Wochen später hätte niemand mehr sagen können, welche der beiden Datumsangaben gilt.
> *Weil* die Information an zwei Orten lag, war die Frage „welcher stimmt?" überhaupt möglich.
> Prüfsatz: *bevor ein Status in ein zweites Dokument geschrieben wird, prüfen, ob ein Verweis
> genügt — wo er genügt, ist die Kopie bereits die Drift.*

**Zwei beobachtbare Closure-Kriterien:**

1. `make gates` und `make verify` grün auf demselben Stand (je Exit 0) — belegt.
2. `grep -cE '^## (Aktuelle Welle|Nächste Wellen|Meilensteine|Abgeschlossene Wellen|Historische)'`
   liefert **5**; keiner der fünf Abschnitte kommt doppelt vor, und weder *Aktuelle Welle* noch
   *Nächste Wellen* enthält eine abgeschlossene Welle.

**Was der Umbau nicht leistet, ausdrücklich:** es gibt **keinen Sensor** für die Roadmap-Form.
`make doc-planning` prüft Lifecycle-Konsistenz, nicht Abschnitts-Struktur. Ein Gate, das nur die
fünf Überschriften zählt, wäre grün gewesen — der Wildwuchs lag **unter** korrekten Überschriften.
Was hier gefehlt hat, war kein Regex, sondern ein Leser; gefunden hat es der Maintainer.

**Eigener Anteil, unbequem:** zwei der vier Verstöße stammen aus dieser Session — der Abschnitt
`## Stand 2026-07-26` und der fehlplatzierte Etappe-F-Absatz. Ich habe beim Nachziehen der
Roadmap den vorhandenen Wildwuchs nicht als Verstoß erkannt, sondern **fortgeschrieben**. Das ist
die vierte Selbstanwendung dieser Serie und die einzige, die kein Sensor gefangen hat.

**Folge-Slices:** keine.

## 7. Sub-Area-Modus

Alle berührten Sub-Areas GF.
