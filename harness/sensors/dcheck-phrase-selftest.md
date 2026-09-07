# `make dcheck-phrase-selftest` — Kalibrierung phrasen-basierter Modul-Konfigurationen

## Vertrag

Vier Kontrollen (zwei Muster × Positiv/Negativ) gegen eigene Fixtures und den
gepinnten Digest: Reagiert `d-check` noch auf die Formulierungen, an denen zwei
Modul-Konfigurationen hängen — die `reviews`-Trigger-Phrase „unabhängiger
Review" und `structure`s `tasks-ignore-pattern`? Er verhindert das **stille
Wegdriften einer bereits getroffenen Formulierungs-Wahl**.

## Grenze — was das Grün nicht abdeckt

1. **Ob eine künftig geänderte Formulierung ebenfalls greifen würde** — geprüft
   ist nur die aktuell empfohlene. Permanent: Der Test kann nicht wissen, was
   jemand morgen schreibt.
2. **Die Korpus-Seite** — ob a-checks eigene Dokumente die Phrase noch tragen,
   sagt der Lauf nicht. Er prüft das *Werkzeug*, nicht den *Bestand*. Genau das
   fiel zweimal aus (slice-120, slice-165); heilbar, geplant als slice-169.
3. **Ein drittes phrasen-basiertes Muster** — `versions.current-from` liest den
   erwarteten Baseline-Stand aus einer Prosa-Zeile in `harness/conventions.md`
   und steht **nicht** in diesen vier Kontrollen. Heilbar, benannt in
   [`BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf`](../../docs/plan/planning/observations/BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf/observation.md).

## Bindung

Harness-Prozess · Antwort auf
[`BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf`](../../docs/plan/planning/observations/BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf/observation.md)
bei 3× · slice-168.
