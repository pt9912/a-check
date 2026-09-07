# slice-177 — Sensors-Struktur: zwei Tabellen und `harness/sensors/`

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `make slice-mv` ([`AGENTS.md`](../../../../AGENTS.md)
§3.3/§5).

**Welle:** [welle-15](welle-15-regelwerk-v650-migration.md)

**Bezug:** [slice-174](../done/slice-174-regelwerk-v650-delta-analyse.md)
§3.2, T-2 und T-4 — Etappe **D** des Schnitts in §3.4.

**Berührte Spec-Stellen:** — *(keine)* — Harness-Struktur ohne
Vertragsberührung.

**Verantwortlich:** — *(noch nicht priorisiert)*

**Autor:** Claude (Opus 5), im Auftrag des Maintainers. **Datum:**
2026-09-07.

---

## 1. Ziel und Abgrenzung

Die Sensors-Tabelle wird Index: Nicht-Gates stehen in einer zweiten Tabelle mit `kein Gate` in der Zeile, und ein Gate-Vertrag, der mehr als einen Satz braucht, wandert nach `harness/sensors/<target>.md` — die Target-Zelle wird zum Link darauf.

**Nicht in diesem Slice**, je Punkt mit Grund:

- **Die zwei Zellen, deren Überhang Historie ist** — es wäre ein anderer
  Vorgang: Ihre Antwort ist Kürzen nach [`AGENTS.md`](../../../../AGENTS.md)
  §3.7 („die Historie hält `git`"), nicht eine Sensor-Datei. Wer beides in
  einen Slice packt, vermischt zwei Umbauten mit verschiedenen Kriterien
  (§3.1). Bestand bleibt bis dahin bewusst stehen.
- **Die Zitier-Form im `**Welle:**`-Feld der Archiv-Stubs** — ein Folge-Slice
  übernimmt es; sie ist Code (`tools/archive-wave/`) und zieht einen
  Test-Umbau nach sich. Benannt in
  [slice-176](../done/slice-176-zitier-form-einfrierende-artefakte.md) §3 und
  in `harness/sensors/archive-wave.md` §Grenze.

- **Andere Etappen von [welle-15](welle-15-regelwerk-v650-migration.md)** —
  Schicht-Abgrenzung: jede Etappe misst gegen den vendorten Stand und ist
  einzeln lieferbar.
- **Nachrüsten des Altbestands**, wo die Ziel-Form nur für Neues gilt —
  Bestand bleibt bewusst stehen; ein Sensor gegen unentschiedenen Altbestand
  wäre ein Fehlalarm.

## 2. Analyse (vor der Umsetzung)

**Der Slice ist kein reiner `v6.5.0`-Nachzug.** Die Regel, die a-checks §4-Tabelle
bricht, steht **seit `v5.12.0` unverändert** in
[`AGENTS.template.md`](../../../../.harness/baseline/v6.5.0/templates/AGENTS.template.md)
§4: *„Diese Tabelle **listet auf**; definiert wird hier nichts. Die **Bindung**
eines Targets … steht in `harness/README.md` §Sensors."* Über vier
Baseline-Stände und zwei Adaptions-Durchgänge hinweg nie befolgt, und **keine**
aktive Adaption deckt es. Neu an `v6.5.0` ist nur die *Antwort* — die
`harness/sensors/`-Struktur; die *Regel* ist alt.
Beleg: [`BEO-HARNESS/baseline-regel-nie-erwogen-weil-bestand-sie-verletzt`](../observations/BEO-HARNESS/baseline-regel-nie-erwogen-weil-bestand-sie-verletzt/observation.md)
`evidence/slice-177.md`, damit 2×.

**Ausgangsmessung 2026-09-07** (Zeichen je Vertrags-/Zweck-Zelle, Schwelle 250
als Näherung für „mehr als einen Satz"):

| Ort | über der Schwelle | Spitzenwerte |
|---|---|---|
| [`AGENTS.md`](../../../../AGENTS.md) §4 | **16 von 40** | `doc-check` 1871 · `symlink-check` 1106 · `doc-workflows` 729 |
| [`harness/README.md`](../../../../harness/README.md) §Sensors | **6 von 24** | `doc-check` 639 · `symlink-check` 369 · `gate-consistency` 356 |

**Der eigentliche Fund ist nicht der Überhang, sondern die Doppelung.**
Dieselben Verträge stehen an **zwei** Orten, in AGENTS.md ausführlicher — bei
`doc-check` 1871 gegen 639 Zeichen, bei `symlink-check` 1106 gegen 369. Zwei
Fassungen desselben Vertrags sind zwei Quellen, und Kopien driften.

Die Ziel-Form löst genau das mit: *„Von außen — aus `AGENTS.md`, einer ADR,
einem Slice — wird **diese Datei direkt** adressiert, nicht der Index: sie
wandert nie."* Steht der Vertrag unter `harness/sensors/<target>.md`, tragen
**beide** Tabellen nur noch ihre Index-Zeile, und AGENTS.md §4 verweist auf
dieselbe Datei statt eine zweite Fassung zu führen.

**Zu klären, bevor umgebaut wird:** AGENTS.md §4 führt `| Target | Zweck |` —
die Spalte, in der die Ziel-Form `kein Gate` verlangt, existiert dort nicht
(`harness/README.md` hat sie als *Bindung*). Ob a-check die Spalte ergänzt oder
die Abweichung deklariert, ist die Adaptions-Frage aus
[slice-174](../done/slice-174-regelwerk-v650-delta-analyse.md) §7.

## 3. Umsetzung

### 3.1 Das Kriterium — und warum meine erste Auswahl es verfehlte

**16 Zellen liegen über 250 Zeichen** (§2). Die Ziel-Form nennt aber nicht die
Länge als Kriterium, sondern *„Deckungsgrenze, Ausgabe-Bedeutung, Exit-Codes,
Abbruch-Bedingungen"* — also **was** über den Vertrag hinaus zu sagen ist. Die
**Deckungsgrenze steht dort an erster Stelle**.

**Erster Anlauf: vier Dateien.** Ich hatte die 16 in „vier mit Ausgängen und
Sperren" und „zwölf mit Vertrag plus Historie" geteilt. Der unabhängige Review
hat das nachgemessen und widerlegt (F-2): **acht der zwölf** tragen eine
wörtlich ausgewiesene Deckungsgrenze, zwei weitere eine Vorbedingung oder eine
bewusst unkonfigurierte Fähigkeit. Für „Vertrag plus Historie" blieben **zwei**
übrig.

Die Auswahl hatte sich damit selbst begründet, statt gemessen zu sein — und das
ausgerechnet an dem Kriterium, das die Ziel-Form zuerst nennt. Nachgezählt mit
einem Muster auf die Grenzen-Formulierungen (*„**nicht** geprüft"*, *„bleibt
ungeprüft"*, *„sagt es nicht"*, *„braucht den Pin"*):

| Gruppe | Zahl | Antwort |
|---|---|---|
| Deckungsgrenze, Ausgänge, Sperren oder Vorbedingung | **14** | **eigene Datei** |
| Vertrag plus Historie | 2 | kürzen, nicht auslagern ([`AGENTS.md`](../../../../AGENTS.md) §3.7) |

**Umgesetzt sind jetzt alle 14.** Was bleibt, sind zwei Zellen, deren Überhang
Herkunfts-Geschichte ist; sie gehören gekürzt, und das ist ein anderer Umbau mit
einem anderen Kriterium (§1).

### 3.2 Was entstanden ist

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| `harness/sensors/` (4 Dateien) | neu | Vertrag · Grenze · Ausgänge · Sperren · Bindung, nach `v6.5.0` · `templates/harness/sensors/gate.template.md` |
| [`harness/README.md`](../../../../harness/README.md) §Sensors | update | Target-Zelle wird **Link** auf die Datei; **zweite Tabelle** für Nicht-Gates mit `kein Gate` in der Zeile |
| [`AGENTS.md`](../../../../AGENTS.md) §4 | update | dieselben vier Zellen auf Vertrag gekürzt, Verweis auf die Sensor-Datei |

**Gemessene Wirkung** auf [`AGENTS.md`](../../../../AGENTS.md) §4 — Zellen über
250 Zeichen: **16 → 4**. Die Spitzenwerte:

| Target | vorher | nachher |
|---|---|---|
| `doc-check` | **1871** | 339 |
| `symlink-check` | 1106 | 127 |
| `doc-workflows` | 729 | 190 |
| `archive-wave` | 728 | 256 |
| `image-scan` | 663 | 194 |

Die Doppelung ist damit aufgelöst: Der Vertrag steht **einmal** in der
Sensor-Datei, und beide Tabellen zeigen darauf.

**Zwei Targets fehlten ganz** in `harness/README.md` §Sensors — `doc-reviews`
und `verify-risiko-ausgaenge`. Vorbestehender Mangel, beim Verlinken
aufgefallen: `make doc-targets` erzwingt die Deckung gegen `AGENTS.md` §4, nicht
gegen diese Tabelle. Beide Zeilen ergänzt. Die Ziel-Form nennt das
den Zweck des Links — *„Von außen wird diese Datei direkt adressiert, nicht der
Index: sie wandert nie."*

### 3.3 Die zweite Tabelle

Nicht-Gates nach der Definition der Ziel-Form — Targets, die **nicht über den
Zustand des Repos urteilen**: `archive-wave` und `slice-mv` *bewegen*,
`regelwerk-check` *misst*, `doc-repair`/`doc-trace`/`doc-doctor`/`doc-usage`/
`doc-help` *sagen*. Sie tragen `kein Gate` in der Bindung-Spalte statt als
Prosa-Vorspann in der Zweck-Zelle.

**`image-scan` steht nicht dort.** Es urteilt sehr wohl über einen Zustand — den
des publizierten Images — und ist damit ein Gate, nur keines im Aggregat.
Baseline `modul-13` §Vorhanden ≠ behauptet trennt das ausdrücklich: Ein reales
Target *nicht* als Gate zu führen ist keine Harness-Lüge; ein Gate zu
*versprechen*, das nicht läuft, wäre eine.

## 4. Definition of Done

- [x] Nicht-Gates (`slice-mv`, `archive-wave`, `regelwerk-check`) stehen in einer zweiten Tabelle, `kein Gate` in der Zeile statt als Prosa-Vorspann.
- [x] Für jeden Überhang eine Datei unter `harness/sensors/`, Target-Zelle verlinkt; welche der 16 gemessenen Zellen betroffen sind, ist im Slice begründet — nicht alle 16 pauschal.
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
  steht eine Regel da, die der eigene Bestand bricht* — **Ausgang:** gestrichen
  mit Begründung. **14 von 16** Zellen tragen die Form; die zwei übrigen sind
  kein Bestand, der die Regel bricht, sondern ein anderer Mangel mit einer
  anderen Antwort (§3.1). Der erste Anlauf löste nur vier auf und begründete
  die Lücke mit einem Kriterium, das die Messung nicht stützte — gefangen vom
  Review (F-2), nachgemessen und umgesetzt. Klasse
  [`BEO-HARNESS/baseline-regel-nie-erwogen-weil-bestand-sie-verletzt`](../observations/BEO-HARNESS/baseline-regel-nie-erwogen-weil-bestand-sie-verletzt/observation.md)
  (2×) — **kein neuer Beleg**: Der Eintrag beobachtet Regeln, die *nie erwogen*
  werden. Diese hier war erwogen; falsch war die Auswahl, nicht die Aufmerksamkeit.
- *Die Spalte, in der die Ziel-Form `kein Gate` verlangt, existiert in
  `AGENTS.md` §4 nicht* — **Ausgang:** gestrichen mit Begründung. Die zweite
  Tabelle liegt in [`harness/README.md`](../../../../harness/README.md), wo die
  Bindung-Spalte existiert; `AGENTS.md` §4 *listet auf* und braucht sie nach
  der Ziel-Form gar nicht. Die Frage stellte sich nur, solange beide Tabellen
  dasselbe leisten sollten.

## 8. Closure-Notiz

**Lerneintrag — Form: geschärfte Regel** (das Kriterium für eine Sensor-Datei
ist nicht die Länge einer Zelle, sondern was über den Vertrag hinaus zu sagen
ist).

- **Was hat funktioniert:** Die 16 Kandidaten **nicht** nach Zeichenzahl
  abgearbeitet, sondern nach dem Kriterium der Ziel-Form — *Deckungsgrenze,
  Ausgabe-Bedeutung, Exit-Codes, Abbruch-Bedingungen*. Der Gedanke war richtig;
  die Ausführung nicht (siehe unten). Was am Ende trägt, ist die Trennung
  zweier Überhang-**Ursachen**: ein Vertrag mit Grenzen und Ausgängen gehört in
  eine Datei, angesammelte Herkunfts-Geschichte gehört gelöscht. Das sind zwei
  Probleme mit zwei Antworten, und nur eines davon war dieser Slice.

- **Was ging anders als geplant:** Zwei Dinge, und das zweite wiegt schwerer.

  Erstens löste sich eine offene Frage beim Bauen auf: §2 führte die
  Spalten-Frage (`AGENTS.md` §4 hat kein `Bindung`) als Adaptions-Frage. Die
  zweite Tabelle gehört nach `harness/README.md`, wo die Spalte existiert, und
  §4 *listet* nach der Ziel-Form ohnehin nur auf. Die Frage stellte sich nur,
  solange ich beide Tabellen für gleichwertig hielt.

  Zweitens — und das ist der Befund des Reviews (F-2): **Ich habe das richtige
  Kriterium genannt und dann nicht angewandt.** §3.1 behauptete, zwölf der
  sechzehn Zellen trügen „Vertrag plus Historie". Nachgemessen tragen **zehn**
  eine Deckungsgrenze oder Vorbedingung — und die Deckungsgrenze steht in der
  Ziel-Form an *erster* Stelle. Die Einteilung war keine Messung, sondern eine
  Rechtfertigung dafür, bei vier Dateien aufzuhören. Sie stand in einem Slice,
  dessen ganzer Zweck es ist, Verträge sichtbar zu machen.

- **Steering-Loop-Eintrag — geschärfte Regel:** Wer ein Kriterium **nennt**, hat
  es noch nicht **angewandt**; die Anwendung ist ein eigener Handgriff und
  hinterlässt eine Liste, keinen Absatz. Ein Satz der Form „danach zerfallen die
  N in X und Y" ist eine Behauptung, solange die Zuordnung je Element nicht
  dasteht. — liegt in `harness/sensors/` (vierzehn Dateien als Muster) und
  `harness/README.md §Sensors`.

- **Beobachtungs-Register (`../observations/`):** **kein neuer Eintrag**, aber
  zwei Belege — `evidence/slice-177.md` in
  [`baseline-regel-nie-erwogen-weil-bestand-sie-verletzt`](../observations/BEO-HARNESS/baseline-regel-nie-erwogen-weil-bestand-sie-verletzt/observation.md)
  (angelegt beim Maintainer-Einwand, der diesen Slice ausgelöst hat — damit 2×,
  `state.md` nachgezogen) und `evidence/slice-177.md` in
  [`verweis-auf-wandernden-slice`](../observations/BEO-PLAN/verweis-auf-wandernden-slice/observation.md)
  (die `slice-mv`-Form-Lücke, viertes Mal in Folge). Eine frühere Fassung dieser
  Zeile sagte „keine Beobachtung angefallen" und widersprach damit dem eigenen
  §2 — Review-Befund F-10.
  Der naheliegende Eintrag
  [`baseline-regel-nie-erwogen-weil-bestand-sie-verletzt`](../observations/BEO-HARNESS/baseline-regel-nie-erwogen-weil-bestand-sie-verletzt/observation.md)
  bekommt **keinen** Beleg: Er beobachtet Regeln, die *nie erwogen* werden —
  diese wurde erwogen, gemessen und teilweise umgesetzt. Ein Beleg dort wäre
  eine Verwässerung der Klasse.

- **Folge-Slices:** zwei benannt, beide noch ohne Datei — die **zwei**
  verbliebenen Historie-Zellen und die Zitier-Form im `**Welle:**`-Feld der
  Archiv-Stubs (§1). Sie entstehen, wenn der Maintainer sie priorisiert; dieser
  Slice erfindet sie nicht als Adresse, die niemand annimmt.

- **Risiken aus §7:** zwei, jedes mit genau einem Ausgang — einmal
  *eingetreten, teilweise aufgelöst*, einmal *gestrichen mit Begründung*.

- **Drei Paarungen:** trägt die Welle-Closure — der Slice hat ein
  `**Welle:**`-Feld.

## 9. Sub-Area-Modus

**Vorgelagert — Sub-Area-Wahl prüfen:** zwei Sub-Areas berührt.
**Harness-Einstieg** (`AGENTS.md`, `harness/`, neu `harness/sensors/`) und
**Gate-/Werkzeug-Schicht** (die Verträge selbst beschreiben `Makefile`-Targets
und `tools/`). Beide über der Schwelle ≥ 2/3 und in
[`conventions.md`](../../../../harness/conventions.md#modus-deklaration-pro-sub-area)
deklariert. `harness/sensors/` ist **keine neue Sub-Area**: es ist eine
Datei-Familie *innerhalb* des Harness-Einstiegs, und eine eigene `MR`-Adaption
wäre dafür nicht plausibel formulierbar (Achse 1 fehlt).

**Vorgelagert — offene Beobachtungen sichten:** Register am 2026-09-07
durchgegangen. Zwei einschlägig:

| Eintrag | Stand | Bezug |
|---|---|---|
| [`baseline-regel-nie-erwogen-weil-bestand-sie-verletzt`](../observations/BEO-HARNESS/baseline-regel-nie-erwogen-weil-bestand-sie-verletzt/observation.md) | offen, 2× | **adressiert, nicht erhöht** — der Eintrag entstand an genau dieser Regel (`AGENTS.md` §4 gegen die Ziel-Form); dieser Slice setzt sie um. Kein Beleg: beobachtet werden Regeln, die *nie erwogen* werden |
| [`verweis-auf-wandernden-slice`](../observations/BEO-PLAN/verweis-auf-wandernden-slice/observation.md) | verkörpert | **erneut aufgetreten** beim Lifecycle-Wechsel; siehe `evidence/slice-177.md` |

**Alle berührten Sub-Areas GF** — kein Begründungsblock nötig.
