# slice-075 — Sensoren messen die Größe, die sie nennen

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Deckt:** `F-4` und `F-13` aus dem
[Review-Report `welle-12`](../../../reviews/2026-08-09-welle-12-unabhaengig.md) (Gruppe B).
**Bezug:** `SL-002`; Roadmap-Zeile *Aktuelle Welle* in der
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

- [x] `verify-closure-notes` zählt Sätze, nicht Satzzeichen. Beleg: „Geprueft via foo.go." als
      einzige Zeile → Exit 1, *„trägt weniger als zwei Sätze"*; „Erster Satz via foo.go. Zweiter
      Satz hier." → Exit 0. Beide Richtungen zusätzlich im Selbsttest verankert.
- [x] `verify-slice-links` sieht Referenz-Links. Beleg: `[rm]: roadmap.md` → Exit 1 mit demselben
      `SL-002`-Befund wie ein Inline-Link; `[rm]: ../in-progress/roadmap.md` → Exit 0. Ebenfalls
      im Selbsttest.
- [x] `make gates` und `make verify` grün — **Ausgabe in eine Datei**, Exit-Code getrennt geprüft,
      nie in eine Pipe.

## 6. Closure-Notiz

**Geliefert:** [`verify-closure-notes.sh`](../../../../tools/verify-closure-notes.sh) zählt
Satzzeichen nur noch, wenn Whitespace oder Zeilenende folgt;
[`verify-slice-links.sh`](../../../../tools/verify-slice-links.sh) erfasst zusätzlich
Referenz-**Definitionen**. Beide Korrekturen sind mit je zwei Selbsttest-Fixturen gegen Rückfall
gesichert.

**Dritter Fund beim Bauen, nicht im Report:** die Satzzählung filterte Code-Blöcke über
`grep -v '^\s*```'` — das entfernt nur die **Fence-Zeilen**, nicht den Inhalt, obwohl der Kommentar
daneben „ausserhalb von Code-Zeilen" behauptete. Ein Code-Block in einer Closure-Notiz zählte also
mit. Behoben mit demselben `sed`-Range, den `verify-slice-links` längst nutzte.

**Lerneintrag — Form: geschärfte Regel.** Als Prüfsatz: *Wenn eine Prüfung leichter zu schreiben
ist als die Frage, die sie beantworten soll, misst sie fast sicher die leichtere Frage.*

**Die Ursache** ist in allen drei Fällen dieselbe: Markdown-**Struktur** wurde mit
**zeilenbasierten** Mitteln nachgebaut. `grep -oE '[.!?]'` ist Satzzeichen-Zählung statt
Satz-Zählung; `grep -oE '\]\([^)]*\)'` ist Inline-Link-Erkennung statt Link-Erkennung;
`grep -v '^\s*```'` ist Fence-Zeilen-Filterung statt Code-Block-Filterung. Jede dieser Näherungen
war beim Schreiben plausibel und ist im Normalfall richtig — sie bricht erst am Sonderfall, und
der taucht in einem gewachsenen Bestand irgendwann auf. Genau diese Klasse hat
[slice-073](../done/slice-073-dcheck-statt-eigenbau.md) als CR 1/CR 2 an d-check formuliert; dieser
Slice hat **nicht** darauf gewartet, weil der Bestand bis dahin falsch misst.

**Zwei beobachtbare Closure-Kriterien:**

1. Die vier Fixturen aus §3 verhalten sich wie erwartet — je eine rot, je eine grün — und alle vier
   liegen als Selbsttest im jeweiligen Sensor, nicht nur in dieser Notiz.
2. Der Bestand bleibt unberührt: 72 Closure-Notizen in `done/` bestehen die **schärfere**
   Satzzählung, 6 wandernde Slices die erweiterte Link-Erkennung.

**Folge-Slices:** keine. [slice-076](../done/slice-076-vertrag-und-sensor.md) behandelt die
verbleibenden zwei Findings der Gruppe B — dort klaffen Vertrag und Sensor, was je eine
Entscheidung braucht.

## 7. Sub-Area-Modus

Alle berührten Sub-Areas GF.
