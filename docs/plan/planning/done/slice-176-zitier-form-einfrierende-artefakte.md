# slice-176 — Zitier-Form für einfrierende Artefakte übernehmen

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `make slice-mv` ([`AGENTS.md`](../../../../AGENTS.md)
§3.3/§5).

**Welle:** [welle-15](../welle-15-regelwerk-v650-migration.md)

**Bezug:** [slice-174](../done/slice-174-regelwerk-v650-delta-analyse.md)
§3.2, T-1 — Etappe **C** des Schnitts in §3.4.

**Berührte Spec-Stellen:** — *(keine)* — Harness-Struktur ohne
Vertragsberührung.

**Verantwortlich:** — *(noch nicht priorisiert)*

**Autor:** Claude (Opus 5), im Auftrag des Maintainers. **Datum:**
2026-09-07.

---

## 1. Ziel und Abgrenzung

Die vier einfrierenden Artefaktklassen zitieren die Baseline als Kennung statt als Adresse — `v<X.Y.Z>` · `regelwerk/<datei>.md` §Abschnitt statt als Link.

**Der Slice trägt zusätzlich den Abschluss von Etappe A.** Der neue und der vorige Stand liegen seit [slice-175](../done/slice-175-etappe-a-vendoring-v650.md)
nebeneinander — zulässig während einer Migration
([`conventions.md`](../../../../harness/conventions.md#baseline) §Baseline),
aber nur bis hierher. Das Löschen hängt an **diesem** Slice und an keinem
anderen, weil die Zitier-Form die Bedingung dafür schafft: Erst wenn die
16 Zeitdokumente die Baseline als Kennung statt als Link zitieren, bricht beim
Entfernen nichts.

**Nicht in diesem Slice**, je Punkt mit Grund:

- **Andere Etappen von [welle-15](../welle-15-regelwerk-v650-migration.md)** —
  Schicht-Abgrenzung: jede Etappe misst gegen den vendorten Stand und ist
  einzeln lieferbar.
- **Nachrüsten des Altbestands**, wo die Ziel-Form nur für Neues gilt —
  Bestand bleibt bewusst stehen; ein Sensor gegen unentschiedenen Altbestand
  wäre ein Fehlalarm.

## 2. Analyse

### 2.1 Der Bestand: 16 Links in 14 einfrierenden Artefakten

Gemessen am 2026-09-07 — Markdown-**Links** auf einen vendorten Baseline-Pfad,
in den Klassen, die `.d-check.yml` als Zeitdokumente führt:

| Klasse | Dateien | Links |
|---|---|---|
| `docs/plan/planning/done/` | 8 | 10 |
| `docs/plan/planning/observations/` (`observation.md` + `evidence/`) | 5 | 5 |
| `harness/conventions/done/` | 1 | 1 |
| `docs/reviews/` | 0 | 0 |

**Alle 16 zeigten auf `v6.2.0`** — also genau auf den Stand, der weichen soll.
Dass `docs/reviews/` leer ausgeht, ist kein Zufall: Die Reports zitieren die
Baseline schon heute als Text.

### 2.2 Umgestellt: Kennung statt Adresse

Die Ziel-Form
(`v6.5.0` · `templates/docs/reviews/review-report.template.md`, Block
*Zitier-Form*) verlangt eine Baseline-Stelle als **Tag + Pfad in Inline-Code**
statt als Link: `` `v<X.Y.Z>` · `regelwerk/<datei>.md` §<Abschnitt> ``.
Begründung dort wörtlich: *„Der vendored Baum trägt genau einen Tag; der Sprung
löscht den alten, und ein Link darauf färbt beim nächsten Bump ein Artefakt rot,
das niemand mehr anfassen darf."*

Alle 16 Links skriptgesteuert umgestellt. Zwei Nachbesserungen von Hand: eine
Stelle nannte den Abschnitt danach doppelt, eine andere brauchte eine
Klammer-Prüfung.

**Ändert das die Aussage der Zeitdokumente?** Nein — im Gegenteil: Der Tag wird
Teil des Textes. `[…](…/v6.2.0/regelwerk/modul-08-agentenrollen.md#…)` sagte
*„diese Stelle"* und ließ den Stand im Pfad verschwinden;
`` `v6.2.0` · `regelwerk/modul-08-agentenrollen.md` §Rollen-Sequenz für eine Welle ``
sagt dasselbe und nennt den Stand ausdrücklich. Das ist eine **Form**-Änderung
an unveränderlichen Dateien, keine Inhalts-Änderung — dieselbe Klasse wie der
Zeiger-Bump in
[slice-172](../done/wellenlos/slice-172-baseline-v600-entfernen.md) §2.2, nur
mit dem besseren Ergebnis: Der Verweis überlebt jeden künftigen Sprung.

### 2.3 Der Beweis: Löschen ohne Kollateralschaden

Der vorige vendorte Stand entfernt, `make gates` unmittelbar danach: **Exit 0**.
Kein `target-missing`, kein Nachziehen, **kein eingefrorenes Artefakt angefasst**.

Zum Vergleich derselbe Vorgang ohne Zitier-Form:
[slice-172](../done/wellenlos/slice-172-baseline-v600-entfernen.md) musste beim
Entfernen von `v6.0.0` **22** Links nachziehen, davon neun in `done/`-Slices und
einen mit einer Korrekturfußnote in einem abgeschlossenen Zeitdokument. Die
Klemme, die
[`conventions.md`](../../../../harness/conventions.md#baseline) §Baseline als
*„Preis des Löschens"* beschreibt, ist damit **bezahlt und nicht mehr fällig**.

### 2.4 `exempt-paths` kann nicht schrumpfen — gemessen

Die DoD verlangte die Prüfung „mit Messung, nicht als Vermutung". Probe:
`exempt-paths` aus [`.d-check.yml`](../../../../.d-check.yml) entfernt,
`make doc-check` gefahren → **28 `version-stale`-Befunde**, verteilt auf
`done/`-Slices, Welle-Dateien und Review-Reports. Danach zurückgesetzt, Exit 0.

**Warum:** Die Zitier-Form entfernt **Links**; der `versions`-Sensor prüft
**Text**. Ein Zeitdokument, das `` `v6.0.0` · `regelwerk/…` `` schreibt, nennt
den alten Stand weiterhin — und *soll* das, denn damals galt er.

Damit ist empirisch bestätigt, was
[slice-174](../done/slice-174-regelwerk-v650-delta-analyse.md) §3.2 T-1 nach dem
Review herausgearbeitet hat: **zwei Sensoren, zwei Fragen.** Die Zitier-Form
beantwortet die Link-Auflösbarkeit, `exempt-paths` die Versions-Staleness. Keine
ersetzt die andere.

## 3. Umsetzung

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| 14 einfrierende Artefakte, 16 Links | update | Kennung statt Adresse (§2.2) |
| [`.harness/skills/reviewer.md`](../../../../.harness/skills/reviewer.md) §Output-Schema | update | der Ort, an dem a-check Review-Reports schreibt — ein Report pro Lauf, und er friert ein |
| [`AGENTS.md`](../../../../AGENTS.md) §5 | update | die Regel für alle vier Artefaktklassen, mit dem gemessenen Beleg statt einer Behauptung |
| vorheriger vendorter Stand | löschen | Abschluss von Etappe A; der Beweis in §2.3 ist der Zweck dieses Slice |

**Nicht angefasst:** `tools/archive-wave/`. Der Grund dafür stand zuerst falsch
da — Review-Befund F-3, und der Befund ist präziser als meine Behauptung.

Gemessen: `grep -rn "baseline" tools/archive-wave/*.go` ist **leer**. Das
Werkzeug zitiert nie eine Baseline-Stelle, der von mir benannte Spalt existiert
nicht. Was es **sehr wohl** erzeugt, ist ein Markdown-**Link** im
`**Welle:**`-Feld des Stubs, samt eigener Nachzieh-Mechanik
(`RewriteFieldForMove` in `tools/archive-wave/archive.go`) — und genau das
verbietet die Zitier-Form mit *„`slice-NNN` statt seines Lifecycle-Pfads"*.
Die Mechanik ist die Antwort auf ein Problem, das die Form gar nicht erst
entstehen lässt.

Das bleibt außerhalb dieses Slice, aber aus einem anderen Grund als dem
zuerst genannten: Es ist **Code**, nicht Doku, und die Umstellung zöge einen
Test-Umbau nach sich (`archive_test.go` prüft die Nachzieh-Mechanik). Ein
eigener Vorgang — mit einem realen Gegenstand statt eines erfundenen.

## 4. Definition of Done

- [x] Die Zitier-Form steht dort, wo a-check sie beim Schreiben liest (Kopier-Hinweise, Skills), und der Bestand ist daran gemessen.
- [x] Geprüft, ob `exempt-paths` in [`.d-check.yml`](../../../../.d-check.yml) danach schrumpfen kann — mit Messung, nicht als Vermutung.
- [x] **Der vorige vendorte Stand ist entfernt** (Verzeichnis-Literal hier
      vermieden — `versions` meldete es sonst als `version-stale`), und zwar *ohne* dass ein
      eingefrorenes Artefakt dafür editiert werden musste — das ist der
      Nachweis, dass die Zitier-Form trägt. `make regelwerk-check` meldet
      danach **einen** Stand ohne „ungeprüft"-Hinweis.
      Übernommen aus [slice-175](../done/slice-175-etappe-a-vendoring-v650.md) §1
      (Maintainer-Entscheidung 2026-09-07).
- [x] Unabhängiger Review durchgeführt (Report unter [`docs/reviews/`](../../../reviews/README.md)).
- [x] `make gates` grün.
- [x] `make verify` grün.
- [x] Jedes Risiko trägt einen Ausgang.

## 5. Trigger

**Start** (`open` → `in-progress`): Etappe A (Vendoring) liegt in `done/`,
Maintainer-Freigabe, WIP-Limit frei.

**Rückführungen:** wächst der Umfang über die DoD hinaus, zurück nach `next/`
zur Zerlegung. Ändert sich der adoptierte Stand erneut, zurück nach `open/`.

## 6. Closure-Trigger

DoD vollständig, `make gates`/`make verify` grün, Closure-Notiz geschrieben.
Der Slice trägt ein `**Welle:**`-Feld und archiviert **mit seiner Welle**
([`AGENTS.md`](../../../../AGENTS.md) §6).

## 7. Risiken und offene Punkte

- *Die Ziel-Form wird übernommen, ohne dass a-checks Bestand sie trägt — dann
  steht eine Regel da, die der eigene Bestand bricht* — **Ausgang:** entfallen,
  gestrichen mit Begründung. Der Bestand wurde **zuerst** umgestellt (16 Links)
  und die Regel **danach** geschrieben; der Beweis in §2.3 ist genau die Probe
  darauf, dass sie trägt. Klasse
  [`BEO-HARNESS/baseline-regel-nie-erwogen-weil-bestand-sie-verletzt`](../observations/BEO-HARNESS/baseline-regel-nie-erwogen-weil-bestand-sie-verletzt/observation.md)
  — hier bewusst in der Gegenrichtung gefahren.
- *Die Umstellung ändert Zeitdokumente, die als unveränderlich geführt werden* —
  **Ausgang:** gestrichen mit Begründung. Es ist eine **Form**-Änderung, keine
  Inhalts-Änderung (§2.2), und sie macht die Aussage präziser statt schwächer:
  der Tag steht danach im Text statt im Pfad. **Einmal traf das nicht zu** und
  der Review hat es gefangen (F-1): In einem Beleg stand das Zitat eines Links
  in einem Code-Span — dort war die Umstellung sehr wohl eine Inhalts-Änderung
  und machte den Satz falsch. Zurückgestellt. Das Risiko ist damit nicht
  theoretisch geblieben, aber sein Anlass ist beseitigt.
- *`exempt-paths` bleibt trotz Zitier-Form nötig, und niemand weiß es* —
  **Ausgang:** gestrichen mit Begründung — die Sorge war, niemand wisse es;
  gemessen (§2.4: 28 Befunde ohne die Ausnahmen) und im Slice aufgeschrieben,
  ist sie gegenstandslos. Beim nächsten Sprung ist die Antwort nachlesbar,
  statt erneut ausprobiert zu werden.

## 8. Closure-Notiz

**Lerneintrag — Form: geschärfte Regel** (Zitier-Form für einfrierende
Artefakte, in [`AGENTS.md`](../../../../AGENTS.md) §5 und im Reviewer-Skill
verankert).

- **Was hat funktioniert:** Erst den Bestand umstellen, dann die Regel
  schreiben — und dazwischen die Probe. Der Ablauf war: 16 Links umstellen,
  vorigen Stand löschen, `make gates`. **Exit 0.** Damit ist die Regel nicht
  behauptet, sondern belegt, und der Beleg ist eine Zahl: slice-172 brauchte an
  derselben Stelle **22** Nachzüge in eingefrorenen Dateien, dieser Slice
  **null**.

- **Was ging anders als geplant:** Die DoD fragte, ob `exempt-paths` nach der
  Zitier-Form schrumpfen kann. Die Erwartung war *ja* — beide adressieren
  denselben Ärger. Gemessen: **nein**, 28 Befunde ohne die Ausnahmen. Die
  Zitier-Form entfernt **Links**, `versions` prüft **Text**, und ein
  Zeitdokument nennt den alten Stand weiterhin, weil damals er galt. Das
  bestätigt empirisch, was
  [slice-174](../done/slice-174-regelwerk-v650-delta-analyse.md) §3.2 T-1 erst
  nach einem Review-Befund richtig getrennt hatte: zwei Sensoren, zwei Fragen.

- **Steering-Loop-Eintrag — geschärfte Regel:** Was einfriert, zitiert **Kennung
  statt Adresse**; eine Baseline-Stelle als Tag plus Pfad in Inline-Code, nie
  als Link. — liegt in `AGENTS.md §5` und
  `.harness/skills/reviewer.md §Output-Schema`.
  Auslöser: [`BEO-HARNESS/zwei-baseline-staende-nach-migrationsende`](../observations/BEO-HARNESS/zwei-baseline-staende-nach-migrationsende/observation.md),
  aufgelöst statt weitergezählt.

- **Beobachtungs-Register (`../observations/`):** kein neuer *Eintrag*; ein
  Beleg, den der Review nachtragen musste — `evidence/slice-176.md` in
  [`BEO-PLAN/verweis-auf-wandernden-slice`](../observations/BEO-PLAN/verweis-auf-wandernden-slice/observation.md)
  (F-4). Die Form-Lücke von `slice-mv` trat zum dritten Mal in Folge auf, der
  Eintrag hatte genau diese Etappe vorhergesagt, und ich habe sie beim Sichten
  übersehen, obwohl die eigene Commit-Message sie nennt. `state.md` von
  [`zwei-baseline-staende-nach-migrationsende`](../observations/BEO-HARNESS/zwei-baseline-staende-nach-migrationsende/observation.md)
  auf *verkörpert* gesetzt — der Doppelstand ist beendet, und diesmal nicht nur
  der Zustand, sondern auch die Ursache: Es gibt jetzt eine Form, die das
  Löschen eines Standes folgenlos macht. Der Beleg in
  [`versions-sensor-trifft-planungs-vorgriff`](../observations/BEO-GATE/versions-sensor-trifft-planungs-vorgriff/observation.md)
  ist um den dritten Fund geschärft, ohne den Zähler zu bewegen (derselbe
  Vorgang).

- **Folge-Slices:** keine neuen. Offen bleiben
  [slice-177](../open/slice-177-sensors-struktur-zwei-tabellen.md) (Etappe D)
  und [slice-178](../open/slice-178-slice-form-ziel-und-abgrenzung.md) (E);
  Etappe B braucht noch ihren Slice.

- **Risiken aus §7:** drei, jedes mit genau einem Ausgang — einmal *entfallen*,
  zweimal *eingetreten und im Slice aufgefangen*.

- **Drei Paarungen:** trägt die Welle-Closure — der Slice hat ein
  `**Welle:**`-Feld.

## 9. Sub-Area-Modus

**Vorgelagert — Sub-Area-Wahl prüfen:** zwei Sub-Areas berührt.
**Planungs-Harness** (`docs/plan/planning/`, 13 der 14 umgestellten Dateien) und
**Harness-Einstieg** (`AGENTS.md`, `.harness/skills/`, `.harness/baseline/`).
Beide über der Schwelle ≥ 2/3 und in
[`conventions.md`](../../../../harness/conventions.md#modus-deklaration-pro-sub-area)
deklariert. `docs/reviews/` (Review-Harness) wäre die dritte, fiel aber leer aus
— dort gab es keinen Link umzustellen.

**Vorgelagert — offene Beobachtungen sichten:** Register am 2026-09-07
durchgegangen. Drei einschlägig:

| Eintrag | Stand | Bezug |
|---|---|---|
| [`zwei-baseline-staende-nach-migrationsende`](../observations/BEO-HARNESS/zwei-baseline-staende-nach-migrationsende/observation.md) | offen, 1× | **aufgelöst** — dieser Slice beendet den Doppelstand, und zwar mit der Entscheidung, die der Eintrag vermisste. `state.md` nachgezogen |
| [`versions-sensor-trifft-planungs-vorgriff`](../observations/BEO-GATE/versions-sensor-trifft-planungs-vorgriff/observation.md) | offen, 2× | **dritter Fund, kein neuer Beleg** — derselbe Vorgang; der bestehende Beleg ist geschärft (§2.3 nannte den Pfad, den der Slice gerade löschen lehrt) |
| [`baseline-regel-nie-erwogen-weil-bestand-sie-verletzt`](../observations/BEO-HARNESS/baseline-regel-nie-erwogen-weil-bestand-sie-verletzt/observation.md) | offen, 2× | **nicht erhöht** — hier bewusst in der Gegenrichtung gefahren: erst der Bestand, dann die Regel (§7) |
| [`verweis-auf-wandernden-slice`](../observations/BEO-PLAN/verweis-auf-wandernden-slice/observation.md) | verkörpert | **erneut aufgetreten**, Beleg `evidence/slice-176.md` — dritter Vorgang in Folge, fünf Vorkommen. Der Eintrag hatte Etappe C vorhergesagt. Beim Sichten übersehen und vom Review nachgetragen (F-4) |

**Alle berührten Sub-Areas GF** — kein Begründungsblock nötig.
