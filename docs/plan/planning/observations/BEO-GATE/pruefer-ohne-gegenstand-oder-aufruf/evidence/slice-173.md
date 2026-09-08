**Vorgang:** slice-173
**Fund:** Die Korpus-Ausprägung hat ein **drittes** Muster bekommen, und es entstand mit diesem
Slice. Die neue `versions`-Konfiguration liest den erwarteten Baseline-Stand über
`current-from: harness/conventions.md#baseline` — also aus einer **Prosa-Zeile in einer anderen
Datei** (`- **Stand:** [v6.2.0](…)`). Ändert jemand deren Wortlaut oder verschiebt den Abschnitt,
verliert der Sensor seinen Erwartungswert.

**Warum das hierher gehört und nicht in einen eigenen Eintrag:** identisch zur bereits
registrierten Form — eine Sensor-Konfiguration hängt an einem Wortlaut, den ein anderes Dokument
führt, und ob er noch trägt, sagt kein Lauf.
[`make dcheck-phrase-selftest`](../../../../../../../Makefile) deckt genau diese Frage ab, aber nur
für zwei Muster (`reviews`-Trigger-Phrase, `structure` `tasks-ignore-pattern`); `current-from`
steht nicht darin.

**Unterschied zu den drei vorigen Belegen — er schwächt den Eintrag nicht, er schärft ihn:** hier
bräche der Sensor **laut**, nicht still. `d-check` bricht bei nicht auflösbarem `current-from`
fail-closed ab, statt grün zu melden. Das nimmt dem Fall die Gefährlichkeit der leeren Prüfmenge
und lässt die Diagnose-Kosten übrig: Die Ursache läge in einer Datei, die mit dem Sensor nichts zu
tun zu haben scheint.

**Konsequenz für den geplanten Ausgang:** [`slice-169`](../../../../done/wellenlos/slice-169-korpus-seitige-kalibrierung.md)
trägt bereits die Korpus-Seite. Sein Gegenstand ist damit um dieses dritte Muster erweitert —
gezählt, nicht separat verplant.
