# `make dcheck-phrase-selftest` — Kalibrierung phrasen-basierter Modul-Konfigurationen

## Vertrag

**Zwei Hälften**, weil eine phrasen-basierte Konfiguration auf zwei Weisen
ausfallen kann:

| Hälfte | Frage | Kontrollen |
|---|---|---|
| **Werkzeug** | Reagiert `d-check` noch auf die Formulierung? | 4 — zwei Muster × Positiv/Negativ, gegen eigene Fixtures und den gepinnten Digest |
| **Korpus** | Trägt a-checks eigener `done/`-Bestand sie noch? | 2 — je Muster eine **Nichtleerheits**-Prüfung |

Betroffen sind die `reviews`-Trigger-Phrase „unabhängiger Review" und
`structure`s `tasks-ignore-pattern`. Zusammen verhindern beide Hälften das
**stille Wegdriften einer bereits getroffenen Formulierungs-Wahl** — die
Werkzeug-Seite, wenn das Muster nicht mehr greift; die Korpus-Seite, wenn es
nichts mehr zu greifen gibt.

Die Korpus-Seite liest das Muster **aus `.d-check.yml`**, nicht aus einer Kopie:
Eine Kopie neben dem Original bleibt grün, nachdem das Original gebrochen wurde
(belegt, Review zu slice-168). Der Auszug ist **fail-closed** — findet er das
Feld nicht, bricht der Lauf ab.

## Grenze — was das Grün nicht abdeckt

1. **Ob eine künftig geänderte Formulierung ebenfalls greifen würde** — geprüft
   ist nur die aktuell empfohlene. Permanent: Der Test kann nicht wissen, was
   jemand morgen schreibt.
2. **Eine Erwartungszahl** — geprüft ist **Nichtleerheit**, nicht eine Größe.
   Die belegte Ausfallart ist „die Menge wird leer"; eine feste Zahl bräche bei
   jedem neuen Slice. Permanent, und Absicht.
   **Wie groß die Mengen heute sind, sagt der Lauf**, nicht diese Datei — die
   Erfolgs-Zeile nennt beide.
3. **Zwölf der vierzehn phrasen-basierten Felder** in `.d-check.yml` — gedeckt
   sind die zwei mit **belegtem** Ausfall. Für die übrigen gibt es keinen
   Vorfall, und ein Sensor ohne Anlass ist selbst eine Behauptung.
   `versions.current-from` wäre der nächste Kandidat (slice-179 belegt ihn als
   drittes Muster), fällt aber **laut** aus: `d-check` bricht bei nicht
   auflösbarem Anker ab, statt grün zu melden. Ihm fehlt die Gefährlichkeit der
   leeren Prüfmenge.

## Sperren

- `Feld '<name>' in .d-check.yml nicht lesbar` — der Auszug findet das
  Muster-Feld nicht (Umbenennung, YAML-Umformatierung) → Feldnamen prüfen; der
  Lauf rät nicht.

## Bindung

Harness-Prozess · Antwort auf
[`BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf`](../../docs/plan/planning/observations/BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf/observation.md)
bei 3× · slice-168.
