# slice-075 — Sensoren messen die Größe, die sie nennen

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Deckt:** `F-4` und `F-13` aus dem
[Review-Report `welle-12`](../../../reviews/2026-08-09-welle-12-unabhaengig.md) (Gruppe B).
**Bezug:** [`SL-002`](../../steering-loop.md); Roadmap-Zeile *Aktuelle Welle* in der
[Roadmap](../in-progress/roadmap.md).

---

## 1. Auslöser

**Mechanismus: der Sensor misst eine Ersatzgröße, die der echten nur ähnlich sieht.** Er ist weder
falsch-grün noch zu spät — er misst schlicht etwas anderes, als sein Name und seine Meldung sagen.

**`F-13` — [`tools/verify-closure-notes.sh:80`](../../../../tools/verify-closure-notes.sh)**
zählt Sätze über `grep -oE '[.!?]'`. Gemessen: die Zeile „Geprueft via `foo.go`." ergibt **2** —
der Punkt im Dateinamen zählt als Satzende. Eine einzeilige Closure-Notiz besteht damit die
Mindestzahl von zwei Sätzen.

**`F-4` — [`tools/verify-slice-links.sh`](../../../../tools/verify-slice-links.sh)** extrahiert
Links über `grep -oE '\]\([^)]*\)'` und sieht damit **nur Inline-Links**. Ein Referenz-Link
(`[Roadmap][rm]` mit `[rm]: roadmap.md`) bricht beim Lifecycle-Wechsel genauso, wird aber nie
geprüft.

Beide Ersatzgrößen stammen aus derselben Quelle: Markdown-Struktur mit zeilenbasierten Mitteln
nachgebaut. Das ist der Befund, den [slice-073](../done/slice-073-dcheck-statt-eigenbau.md) als
CR 1/CR 2 an d-check formuliert hat — **dieser Slice wartet nicht darauf**, sondern repariert die
Messgröße im Bestand.

## 2. Betroffene Module

Eine Schicht: **`tools/`** — [`verify-closure-notes.sh`](../../../../tools/verify-closure-notes.sh)
und [`verify-slice-links.sh`](../../../../tools/verify-slice-links.sh).

## 3. Auszuführende Gates

`make verify`, `make gates`.

**Negativ-Proben:**

| Probe | Erwartung |
|---|---|
| Closure-Notiz aus **einer** Zeile mit Punkt im Dateinamen | rot — zählt als **ein** Satz |
| Closure-Notiz aus zwei echten Sätzen | grün |
| wandernder Slice mit **Referenz**-Link auf einen Nachbarn | rot mit demselben `SL-002`-Befund wie beim Inline-Link |
| derselbe Slice mit zustandsunabhängigem Referenz-Link | grün |

Die jeweils zweite Zeile ist Pflicht: eine Satzzählung, die auch echte Sätze verwirft, und eine
Link-Extraktion, die auch korrekte Verweise meldet, wären kein Fortschritt.

## 4. Was bewusst nicht getan wird

- **Auf CR 1/CR 2 an d-check warten.** Deren Umsetzung ist ungewiss und liegt in einem Fremdrepo;
  bis dahin misst der Bestand falsch. Fällt die Prüfung später an d-check, ist dieser Slice der
  Paritäts-Maßstab.
- **`mq` oder ein anderer Markdown-Parser als Abhängigkeit.** Hermetik-Regel
  [`AGENTS.md`](../../../../AGENTS.md) §3.1: kein Host-Toolchain-Aufruf. Die Korrektur bleibt in
  Bash.
- **`F-3` und `F-6`** — dort klaffen Vertrag und Sensor, was je eine Entscheidung braucht; eigener
  Slice.

## 5. DoD

- [ ] `verify-closure-notes` zählt Sätze, nicht Satzzeichen. Beleg: die ersten beiden Proben aus
      §3, vorher grün, nachher rot bzw. unverändert grün.
- [ ] `verify-slice-links` sieht Referenz-Links. Beleg: die letzten beiden Proben aus §3.
- [ ] `make gates` und `make verify` grün — **Ausgabe in eine Datei**, Exit-Code getrennt geprüft,
      nie in eine Pipe.

## 6. Closure-Notiz

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice,
[`AGENTS.md`](../../../../AGENTS.md) §5; `make verify` prüft das.)_

## 7. Sub-Area-Modus

Alle berührten Sub-Areas GF.
