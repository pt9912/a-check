# `make dcheck-phrase-selftest` — Kalibrierung phrasen-basierter Modul-Konfigurationen

## Vertrag

**Zwei Hälften**, weil eine phrasen-basierte Konfiguration auf zwei Weisen
ausfallen kann:

| Hälfte | Frage | Kontrollen |
|---|---|---|
| **Werkzeug** | Reagiert `d-check` noch auf die Formulierung? | 4 — zwei Muster × Positiv/Negativ, gegen eigene Fixtures und den gepinnten Digest |
| **Korpus** | Trägt a-checks eigener `done/`-Bestand sie noch? | 1 — **Nichtleerheit** der `reviews`-Kandidatenmenge |

Die Werkzeug-Seite prüft **zwei** Muster: die `reviews`-Trigger-Phrase
„unabhängiger Review" und `structure`s `tasks-ignore-pattern`. Die Korpus-Seite
prüft **eines** — nur die `reviews`-Phrase. Der Unterschied ist die Ausfall-Art:
Verliert `reviews` seine Kandidaten, meldet `doc-reviews` **grün, ohne etwas zu
prüfen**. Verliert `tasks-ignore-pattern` seine Treffer, zählen die konstanten
DoD-Posten mit und `doc-structure` meldet `section-oversized` — **laut**, und
damit ohne Bedarf für einen Wächter (Review slice-169, F-6).

Zusammen verhindern die Hälften das **stille Wegdriften einer bereits
getroffenen Formulierungs-Wahl** — die Werkzeug-Seite, wenn das Muster nicht
mehr greift; die Korpus-Seite, wenn es nichts mehr zu greifen gibt.

**Gezählt wird die Menge des Moduls, nicht eine Obermenge.** Das Modul `reviews`
sieht einen **DoD-Haken** in einem **flachen** `done/`-Slice; eine Nennung in
Prosa, in einer Tabelle oder in einer Wellen-Ergebnisnotiz sieht es nicht, und
archivierte Stubs tragen keine DoD mehr. Eine weiter gefasste Zählung meldet
grün, während die echte Menge leer ist — gemessen an der ersten Fassung dieser
Kontrolle, die bei entwerteten Kandidaten Exit 0 meldete (Review slice-169,
F-1). Das Kandidatenverzeichnis kommt **aus `.d-check.yml`** (`done-dir`), nicht
aus einer Kopie: Eine Kopie neben dem Original bleibt grün, nachdem das Original
umgezogen ist (belegt, Review zu slice-168). Der Auszug ist **fail-closed** —
findet er das Feld nicht oder zeigt es ins Leere, bricht der Lauf ab.

## Grenze — was das Grün nicht abdeckt

1. **Ob eine künftig geänderte Formulierung ebenfalls greifen würde** — geprüft
   ist nur die aktuell empfohlene. Permanent: Der Test kann nicht wissen, was
   jemand morgen schreibt.
2. **Eine Erwartungszahl** — geprüft ist **Nichtleerheit**, nicht eine Größe.
   Der Lauf ist rot bei leerer Menge und grün bei jeder Größe darüber; die Zahl
   in der Erfolgs-Zeile ist Ausgabe, keine Zusage. Permanent, und Absicht.
   **Wie groß die Menge heute ist, sagt der Lauf**, nicht diese Datei — die
   Erfolgs-Zeile nennt sie.
3. **Die übrigen phrasen-basierten Felder** in `.d-check.yml` — gedeckt ist
   das eine mit **belegtem** Ausfall (die `reviews`-Trigger-Phrase). Für die
   anderen gibt es keinen Vorfall, und ein Sensor ohne Anlass ist selbst eine
   Behauptung. **Keine Gesamtzahl:** sie hinge an einer Zählregel für
   „phrasen-basiert", die nirgends festgelegt ist.
   Zwei Kandidaten sind ausdrücklich draußen, beide weil sie **laut** ausfallen:
   `versions.current-from` (Beleg `evidence/slice-173.md` im Register) — `d-check`
   bricht bei nicht auflösbarem Anker ab; und `structure.tasks-ignore-pattern` —
   trifft es nichts, zählen die konstanten DoD-Posten mit und `doc-structure`
   meldet `section-oversized`. Beiden fehlt die Gefährlichkeit der leeren
   Prüfmenge, die grün meldet.

## Sperren

- `done-dir aus .d-check.yml nicht lesbar oder kein Verzeichnis` — der Auszug
  findet das Feld nicht (Umbenennung, YAML-Umformatierung) oder es zeigt ins
  Leere → Feld und Pfad prüfen; der Lauf rät nicht.
- `Kandidatenmenge des reviews-Moduls ist LEER` — kein flacher `done/`-Slice
  trägt die Phrase auf einem DoD-Haken → `make doc-reviews` prüft nichts mehr;
  die Phrase gehört in die DoD neuer Slices (`AGENTS.md` §5, Kopieranleitung
  Punkt 7).

## Bindung

Harness-Prozess · Antwort auf
[`BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf`](../../docs/plan/planning/observations/BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf/observation.md)
bei 3× · slice-168 (Werkzeug-Hälfte) · slice-169 (Korpus-Hälfte).
