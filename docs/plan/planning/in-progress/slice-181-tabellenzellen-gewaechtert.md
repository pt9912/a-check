# slice-181 — Tabellenzellen maschinell gewächtert

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `make slice-mv` ([`AGENTS.md`](../../../../AGENTS.md)
§3.3/§5).

**Welle:** ohne Welle (der Closure-Trigger wäre die eigene DoD — kein
repo-weites Mehr).

**Bezug:** Maintainer-Hinweis 2026-09-07 auf die `table.column`-Fähigkeit des
`structure`-Moduls. Sie liegt seit dem Pin `v0.74.1` im Werkzeug und ist nie
konfiguriert worden —
[`BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf`](../observations/BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf/observation.md)
in einer Ausprägung, die der Eintrag bisher nicht kennt: nicht ein Muster ohne
Gegenstand, sondern ein **ganzes Feld ohne Konfiguration**.

**Berührte Spec-Stellen:** — *(keine)* — Gate-Konfiguration ohne
Vertragsberührung.

**Verantwortlich:** Claude — gesetzt beim Übergang nach `in-progress/`. *(noch nicht priorisiert)*

**Autor:** Claude (Opus 5), im Auftrag des Maintainers. **Datum:**
2026-09-07.

---

## 1. Ziel und Abgrenzung

Eine Tabellenzelle, die zum Absatz geworden ist, meldet `make doc-structure` —
statt dass jemand sie zählt. `table.column` mit `cell-max-chars` und
`cell-min-chars` ist für [`harness/README.md`](../../../../harness/README.md)
§Sensors und [`harness/conventions.md`](../../../../harness/conventions.md)
§Aktive Adaptionen konfiguriert, und die dabei gefundenen Zellen sind aufgelöst.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **[`AGENTS.md`](../../../../AGENTS.md) §4** — es wäre ein **anderer Vorgang**:
  Dort steht keine `Vertrag`-Spalte, sondern `Zweck`, und die vier verbliebenen
  überlangen Zellen tragen Historie statt Vertrag
  ([slice-177](../done/welle-15/slice-177-sensors-struktur-zwei-tabellen.md) §3.1).
  Ihre Antwort ist Kürzen nach §3.7, nicht eine Grenze.
- **Eine repo-weite Zellen-Grenze** — Bestand bleibt bewusst stehen: Der
  Roadmap-Bestand und die Beobachtungs-Register-Tabellen tragen bewusst lange
  Zellen, und eine Regel, die den Bestand massenhaft bricht, wird abgeschaltet
  statt befolgt ([`AGENTS.md`](../../../../AGENTS.md) §5).
- **Die `order`-Hälfte von `table`** (Chronologie-Monotonie) — es wäre ein
  anderer Vorgang mit eigenem Anlass; hier gibt es keinen.

**Eine benannte Plan-Änderung** ([`AGENTS.md`](../../../../AGENTS.md) §6) —
**nachträglich benannt, und das ist selbst der Befund:** Die Regel verlangt sie
*vor* dem Code; hier fand sie der Review (F-7).

**[`AGENTS.md`](../../../../AGENTS.md) §4 ist doch konfiguriert.** Der
Ausschluss oben begründete sich damit, die vier überlangen Zellen dort trügen
*„Historie statt Vertrag"*, und ihre Antwort sei Kürzen. Gemessen trifft das
auf **eine** zu (`doc-targets`); die anderen trugen Deckungsgrenzen, die in
Sensor-Dateien gehören — dieselbe Klasse wie in
[`harness/README.md`](../../../../harness/README.md). Die Tabelle hat dieselbe
Funktion, dieselbe Fehlerart und verdient dieselbe Grenze; nur heißt ihre
Spalte `Zweck`.
**Die zugesagte dritte Tabelle ist ebenfalls konfiguriert** —
[`harness/conventions.md`](../../../../harness/conventions.md) §Aktive
Adaptionen stand in §1 als Zusage und fehlte im ersten Anlauf. Beides zusammen
heißt: Der Slice deckt jetzt drei Tabellen statt der geplanten zwei, und die
Liefer-Punkte bleiben bei drei — die Konfiguration ist **einer**, unabhängig
davon, wie viele Tabellen sie adressiert.

## 2. Ausgangsmessung (2026-09-07)

Die Fähigkeit **existiert im Pin** und ist im Startgerüst dokumentiert
(`d-check --print-config`, `structure`-Block): `table.column` adressiert eine
Spalte über ihren **Kopfzeilen-Namen**, `cell-max-chars` meldet
`section-cell-oversized` auf **ihrer** Zeile, `cell-min-chars`
`section-cell-undersized`; eine nicht adressierbare Spalte meldet
`section-column-missing`.

Probeweise eingesetzt und wieder zurückgenommen — **8 Befunde**:

| Ort | Zeichen | Spalte |
|---|---|---|
| `harness/README.md`:78 (`gate-consistency`) | **356** | Vertrag |
| :95 (`verify-observations`) | 301 | Vertrag |
| :91 (`gates`) | 288 | Vertrag |
| :89 (`doc-workflows`) | 260 | Vertrag |
| :92 · :80 · :79 | 244 · 232 · 205 | Vertrag |
| `harness/conventions.md`:128 ([`MR-019`](../../../../harness/conventions.md#mr-019)) | 333 | Ersetzt-Baseline-Regel |

**Warum das zählt:** [slice-177](../done/welle-15/slice-177-sensors-struktur-zwei-tabellen.md)
hat 16 solche Zellen **von Hand** vermessen und den Schnitt selbst begründet —
der Review nannte die Begründung eine Selbstrechtfertigung statt einer Messung.
Mit `cell-max-chars` stellt sich die Frage nicht: Der Lauf zählt.

**Zu klären vor der Umsetzung:**

1. **Welche Obergrenze?** 200 Zeichen ist der Wert aus dem Hinweis. Gegen den
   Bestand zu messen, nicht zu setzen: Bei welcher Schwelle bleiben die Zeilen
   übrig, die wirklich Absätze sind?
2. ~~Trägt `cell-min-chars` auf der Bindung-Spalte?~~ **Beantwortet, teils aus
   der Ziel-Form, teils gemessen.**

   *Warum keine Obergrenze:* `v6.5.0` · `templates/harness/README.template.md`
   beschreibt die Spalte als *„strukturelle Referenzen — Carveout-ID
   (`CO-<NNN>`), Slice-ID, Schwelle, Image-Hash, ADR-ID"*. Das ist eine
   **Kennungs-Liste**, keine Prosa; ihre Länge wächst mit der Zahl der
   Bindungen, nicht mit der Geschwätzigkeit. Eine Zeichen-Obergrenze bestrafte
   dort das Gegenteil von dem, was sie meint. Gemessen stützt das die Form: Ø
   **114** Zeichen gegen Ø **152** in der Vertrags-Spalte — die Bindung ist im
   Schnitt *kürzer*, und die Ausreißer sind Listen.

   *Warum trotzdem eine Untergrenze:* Sie bindet die Spalte überhaupt erst.
   Ohne sie meldet eine Zeile, der die Spalte **ganz fehlt**, nichts —
   `section-column-missing` greift nur auf adressierten Spalten. **Bei uns ist
   das heute gegenstandslos** (alle 31 Gate-Zeilen führen drei Spalten,
   gemessen), aber die Untergrenze kostet nichts und deckt den Fall, bevor er
   eintritt. Schwelle **1**: `—` ist die kürzeste zulässige Form, leer ist
   keine.
3. **Sechs der acht sind Vertrags-Zellen ohne Sensor-Datei.** Kürzen oder
   auslagern? Das Kriterium steht in
   [slice-177](../done/welle-15/slice-177-sensors-struktur-zwei-tabellen.md) §3.1
   und ist anzuwenden, nicht neu zu erfinden.
4. **Drei Abweichungen zur Ziel-Form, die die Konfiguration betreffen** —
   gemessen gegen `v6.5.0` · `templates/harness/README.template.md`:

   | | Ziel-Form | a-check | Folge für `table.column` |
   |---|---|---|---|
   | Zweite Tabelle | fette Zeile *„Werkzeuge — genannt, weil der Lauf sie braucht, aber kein Gate:"* | eigene Überschrift `### Nicht-Gates` | Die `###` erlaubt, die zweite Tabelle **separat** zu adressieren — ein Vorteil, aber eine Abweichung |
   | Spaltenname | `Tut was` | `Was es tut` | **`name: Tut was` träfe bei uns nichts** → `section-column-missing`. Entweder Spalte umbenennen oder Konfiguration anpassen |
   | Abschnitt `## Rollen und ihre Übergabe-Artefakte` | kennt sie nicht | seit slice-066 | berührt `table.column` nicht, gehört aber in die Bestandsaufnahme |

5. **Nicht die Zellen — die Prosa daneben.** Derselbe Überhang sitzt auch
   **über** und **unter** den Tabellen, wo kein `table.column` hinreicht
   (§Sensors: 1930 Zeichen sichtbare Prosa gegen 346 der Ziel-Form). Das ist
   eine **Inhalts**-Entscheidung ohne Sensor und damit ein eigener Vorgang:
   [slice-182](../done/wellenlos/slice-182-readme-verweist-statt-wiederholt.md) — dort die
   Messung, hier nur der Zeiger. Beide Slices fassen `harness/README.md` an;
   wer zuerst läuft, gibt dem anderen den neuen Bestand vor.

   **Zu entscheiden (Punkt 4):** angleichen oder abweichen. Die Überschrift hat einen
   messbaren Vorteil (separate Adressierbarkeit); der Spaltenname hat keinen —
   dort ist Angleichen billiger als eine abweichende Konfiguration, die beim
   nächsten Vorlagen-Vergleich wieder auffällt.

## 3. Umsetzung

*(offen)*

## 4. Definition of Done

- [ ] `table.column` ist für beide Tabellen in [`.d-check.yml`](../../../../.d-check.yml)
      konfiguriert; die Schwelle ist **am Bestand gemessen**, nicht übernommen.
- [ ] Die gefundenen Zellen sind aufgelöst — je Zelle nach dem Kriterium aus
      slice-177 §3.1 (Vertrag mit Ausgängen ⇒ Sensor-Datei; Historie ⇒ kürzen).
- [ ] Mutations-Probe in beide Richtungen: eine künstlich verlängerte Zelle
      meldet, eine konforme nicht.
- [ ] Unabhängiger Review durchgeführt (Report unter [`docs/reviews/`](../../../reviews/README.md)).
- [ ] `make gates` grün.
- [ ] `make verify` grün.
- [ ] Jedes Risiko trägt einen Ausgang.

## 5. Trigger

**Start** (`open` → `in-progress`): Maintainer-Freigabe und WIP-Limit frei.

**Rückführungen:** bricht die gemessene Schwelle mehr als eine Handvoll Zeilen,
zurück nach `next/` — dann ist es kein Nachzug, sondern ein Umbau der beiden
Tabellen und gehört getrennt geschnitten.

## 6. Closure-Trigger

DoD vollständig, `make gates`/`make verify` grün, Closure-Notiz geschrieben.
Danach Archivierung als wellenloser Slice ([`AGENTS.md`](../../../../AGENTS.md) §6).

## 7. Risiken und offene Punkte

- *Eine Zeichen-Obergrenze misst Länge, nicht Prosa — eine lange Zelle aus
  Kennungen und Links wäre ein Fehlalarm* — **Ausgang:** *weiter offen* →
  Beobachtungs-Register,
  [`BEO-GATE/muster-trifft-nur-die-haeufige-schreibweise`](../observations/BEO-GATE/muster-trifft-nur-die-haeufige-schreibweise/observation.md).
  Für die Bindung-Spalte ist es entschieden (§2.2: keine Obergrenze, weil sie
  eine Kennungs-Liste trägt). Für die Vertrags-Spalte ist das Risiko **kleiner
  als gedacht und größer zugleich**: Ein Fehlalarm ist bisher nicht aufgetreten
  (alle neun Befunde waren echte Absätze), aber der umgekehrte Fall ist
  gemessen — **16** der 66 Zellen tragen mehr als einen Satz, die meisten unter
  der Schwelle. Der Proxy meldet zu wenig, nicht zu viel; als Grenze steht das
  in [`.d-check.yml`](../../../../.d-check.yml) und in
  `harness/sensors/doc-structure.md`.
- *Die Grenze deckt zwei Tabellen; die dritte (`AGENTS.md` §4) bleibt
  ungewächtert, und dort stehen die längsten Zellen* — **Ausgang:** *entfallen*,
  gestrichen mit Begründung: Es ist **eingetreten und im selben Slice
  aufgelöst** — `AGENTS.md` §4 ist konfiguriert
  (§1, Plan-Änderung), ebenso die in §1 zugesagte dritte Tabelle
  [`harness/conventions.md`](../../../../harness/conventions.md) §Aktive
  Adaptionen. Das Risiko hat sich als richtig erwiesen — nur nicht als Risiko,
  sondern als Lücke im eigenen Ausschluss.
  *(Die geschlossene Dreier-Menge kennt für „eingetreten" nur Carveout oder
  Folge-Slice; ein Risiko, das im selben Slice behoben wird, besteht nicht mehr
  und ist damit* entfallen — *die Begründung nennt beides, damit die Chronologie
  nicht verlorengeht.)*

## 8. Closure-Notiz

**Lerneintrag — Form: geschärfte Regel** (Mutations-Proben, `AGENTS.md` §5).

- **Was hat funktioniert:** Die Schwelle **am Bestand gemessen** statt gesetzt.
  200 stand als Vorschlag im Hinweis; die Messung zeigte 18 Befunde dort und 9
  bei 250 — und die 9 waren genau die unlesbaren. Ohne die Messung wäre 200
  übernommen worden, und eine Regel, die den Bestand massenhaft bricht, wird
  abgeschaltet statt befolgt.

- **Was ging anders als geplant — der Review fand sieben HIGH, und alle sieben
  betrafen Aussagen, nicht Mechanik.** Sein Verdikt trennt das ausdrücklich:
  *„Die Mechanik trägt … blockierend ist die Aussagenschicht darum herum."*
  Konfiguration, Schwelle, Untergrenze und alle vier Proben waren
  reproduzierbar; falsch waren die Zahlen daneben, die Beschreibung der
  Sensor-Datei und die Zuordnung zum Plan.

- **Die eine Ursache hinter drei Findings:** Ein Split an `|` übersah **escapte**
  Pipes. Daraus folgten (a) eine übersehene Zelle, die erst der Sensor fand,
  (b) die falschen Kalibrierungs-Zahlen 17/8 statt 18/9, (c) die falsche
  Behauptung, nur eine der 66 Zellen trage mehr als einen Satz — es sind 16,
  und mein Zähler zählte Satz*zeichen* statt Sätze. **Drei Findings, ein
  defektes Muster.**

- **Steering-Loop-Eintrag — geschärfte Regel:** *Eine Mutations-Probe belegt
  erst, wenn sie rot war — und die rote Richtung ist die, die zählt.* Grün
  beweist nichts: Ein Prüfer, der seinen Gegenstand nicht erreicht, ist grün,
  und eine Probe, die ihn verfehlt, ebenso. — liegt in
  [`AGENTS.md`](../../../../AGENTS.md) §5 (`seit slice-181` dort).
  Auslöser: [`BEO-GATE/probe-liefert-den-gegenstand-mit`](../observations/BEO-GATE/probe-liefert-den-gegenstand-mit/observation.md)
  (slice-169, slice-180, slice-181 — **3×**, Schwelle erreicht).
  **Der Beleg dieses Slice ist der erste, den die Sitzung selbst fand** statt
  der Review — weil das Ergebnis nicht zur Erwartung passte. Genau das ist die
  Frage, die die Regel festhält.

- **Beobachtungs-Register (`../observations/`):** zwei Belege.
  `probe-liefert-den-gegenstand-mit` erreicht **3×** und wechselt auf
  *verkörpert*;
  [`muster-trifft-nur-die-haeufige-schreibweise`](../observations/BEO-GATE/muster-trifft-nur-die-haeufige-schreibweise/observation.md)
  steht mit dem escapten Pipe bei **2×** und bleibt offen — ein allgemeines
  Gegenmittel ist nicht in Sicht, weil „welche Schreibweisen hat mein
  Gegenstand?" mit einem *weiteren* Muster beantwortet würde.

- **Folge-Slices:** keiner neu. Die Inhaltsspalte der Nicht-Gates-Tabelle
  (`Was es tut`) bleibt unadressiert und ist als Grenze in
  `harness/sensors/doc-structure.md` benannt; sie bekommt einen Slice, wenn dort
  eine Zelle auffällt.

- **Risiken aus §7:** siehe dort, jedes mit genau einem Ausgang.

- **Drei Paarungen** (Repo ohne Wellen-Betrieb) — geprüft **nach** dem `git mv`
  nach `done/`, weil sie dort suchen; eingetragen im dritten Closure-Commit.

## 9. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** zwei Sub-Areas berührt.
**Gate-/Werkzeug-Schicht** `GATE` ([`.d-check.yml`](../../../../.d-check.yml))
und **Harness-Einstieg** `HARNESS` ([`AGENTS.md`](../../../../AGENTS.md) §4,
[`harness/README.md`](../../../../harness/README.md) §Sensors,
[`harness/conventions.md`](../../../../harness/conventions.md),
`harness/sensors/`) — beide Achsen 1,2,3.

**Vorgelagert — offene Beobachtungen sichten:** gesichtet am 2026-09-08.
`BEO-GATE/` führt **21** Einträge; drei sind einschlägig, und einer davon
erreicht mit diesem Slice die Schwelle.

- [`BEO-GATE/probe-liefert-den-gegenstand-mit`](../observations/BEO-GATE/probe-liefert-den-gegenstand-mit/observation.md)
  — **erreicht 3×**. Die erste Mutations-Probe dieses Slice setzte ihren
  Fülltext hinter das schließende `|`, traf die Zelle also nicht und meldete
  grün. Der Eintrag entstand in slice-180; im unmittelbar folgenden Slice trat
  er wieder auf. **Unterschied zu den ersten beiden Malen:** Diesmal fiel es in
  der Sitzung auf, weil das Ergebnis nicht zur Erwartung passte — nicht erst im
  Review. Bei 3× ist es eine Lücke und braucht einen Ausgang.
- [`BEO-GATE/muster-trifft-nur-die-haeufige-schreibweise`](../observations/BEO-GATE/muster-trifft-nur-die-haeufige-schreibweise/observation.md)
  — **zweites Vorkommen**. Die Vorab-Messung splittete Tabellenzeilen an `|` und
  übersah **escapte** Pipes (`\|`); dieselbe Zelle erschien dadurch als mehrere
  kurze. Der Sensor fand sie, meine Zählung nicht — und dieselbe Ursache trägt
  die falschen Zahlen aus Review-Finding F-1.
- [`BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf`](../observations/BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf/observation.md)
  — **kein Beleg.** Der Verdacht lag nahe (eine Konfiguration, die ihre Tabelle
  nicht trifft), und genau die Probe hat ihn ausgeschlossen: Die Regel greift
  auf allen drei Tabellen, mutations-belegt. Keine Treffer sind auch eine
  Antwort.

**Alle berührten Sub-Areas GF** — kein Begründungsblock nötig.

**Alle berührten Sub-Areas GF** — kein Begründungsblock nötig.
