# slice-193 — Der Gate-Index steht einmal

**Welle:** ohne Welle.

**Bezug:** Folge-Slice aus
[slice-192](../done/wellenlos/slice-192-baseline-v660-vendoring.md) §1 — die eine
inhaltliche Neuerung des Sprungs auf `v6.6.0`.
[`AC-QA-02`](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze).

**Berührte Spec-Stellen:** — · Der Slice berührt kein Spec-Stratum.

**Verantwortlich:** Claude — gesetzt beim Übergang nach `in-progress/`.

**Autor:** Claude. **Datum:** 2026-09-08.

**Lerneintrag — Form:** geschärfte Regel.

---

## 1. Ziel und Abgrenzung

**Ziel:** Der Gate-Index steht **einmal** — in
[`harness/README.md`](../../../../harness/README.md) §Sensors.
[`AGENTS.md`](../../../../AGENTS.md) §4 trägt die **Regel** (kein behauptetes
Gate ohne Deckung) und den **Zeiger**, nicht die Liste. Die drei Prüfer, die
heute auf zwei Indizes zeigen, zeigen danach auf einen.

**Der Grund ist nicht Ordnungsliebe** — er steht in `grundlagen-harness-dateien.md`
(`v6.6.0`): Beide Dateien liegen in **jedem** Lauf-Kontext, ein zweiter Index
wird also pro Lauf **zweimal bezahlt**, und er läuft auseinander, weil die
Pflicht, ihn nachzuziehen, nirgends steht.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Neue Gates oder geänderte Verträge.** *Schicht-Abgrenzung:* Der Slice
  bewegt eine Deklaration; welche Zusage ein Target trägt, entscheidet er nicht.
- **Die Aufteilung Gates / Nicht-Gates.** *Bestand bleibt bewusst stehen:*
  `harness/README.md` führt beide Blöcke bereits mit dem Kriterium *urteilen*
  gegen *bewegen · messen · sagen*. Was aus §4 kommt, wird dort **eingeordnet**,
  nicht neu sortiert.
- **Der Voll-Abgleich der Ziel-Formen.** *Ein Folge-Slice übernimmt es:*
  [slice-188](../open/slice-188-voll-abgleich-gate-und-skill.md) /
  [slice-189](../open/slice-189-voll-abgleich-spec-straten.md).

## 2. Ausgangsmessung (2026-09-08)

| Gegenstand | Ist-Stand |
|---|---|
| [`AGENTS.md`](../../../../AGENTS.md) §4 | 7094 Zeichen, **41** Tabellenzeilen |
| [`harness/README.md`](../../../../harness/README.md) §Sensors | 9512 Zeichen, **28** Tabellenzeilen, dazu §Nicht-Gates |
| `targets` in [`.d-check.yml`](../../../../.d-check.yml) | `doc-tables: [AGENTS.md, harness/README.md]`, `authority: AGENTS.md` |
| `mentions` in [`.d-check.yml`](../../../../.d-check.yml) | `documents: ["AGENTS.md"]` — jede Sensor-Datei ist dort genannt |
| `structure`-Zellengrenze | `cell-max-chars` auf der Spalte *Zweck* in `AGENTS.md` §4 |
| Target-Zellen mit Argument in der Code-Span | **null**, in beiden Tabellen |

**Die Differenz 41 gegen 28 ist der eigentliche Umfang:** `AGENTS.md` §4 führt
Targets, die in `harness/README.md` gar nicht stehen. Welche das sind, ist die
erste Handlung des Slice — und sie ist maschinell zu erheben, nicht zu schätzen.

## 3. Umsetzung

### 3.1 Die Differenz-Messung — und warum die erste falsch war

Die Zahl aus §2 (41 gegen 28) ist eine **Zeilen**-Zahl, keine Target-Zahl. Der
erste Abgleich zählte je Zeile **ein** `make X` und meldete acht Targets, die
nur in [`AGENTS.md`](../../../../AGENTS.md) §4 stünden. Falsch: Der Index führt
eine Sammelzeile mit **vier** advisory-Targets in einer Zelle, und der erste
Zähler sah davon nur das erste. **Dieselbe Klasse, die der Bestand als
[`kandidaten-klassifikation-groeber-als-der-kandidat`](../observations/BEO-PLAN/kandidaten-klassifikation-groeber-als-der-kandidat/observation.md)
führt** — hier an der Messung selbst, nicht an ihrer Beschreibung.

Richtig gezählt, über **alle** `make X` je Zelle und zusätzlich gegen die Prosa:

| Menge | Zahl |
|---|---|
| Target-Zellen in `AGENTS.md` §4 | 41 |
| Target-Nennungen im Index, inkl. Prosa | 39 |
| **nur in `AGENTS.md`, auch gegen die Prosa** | **3** — `archive-wave-test`, `commit-scope-check`, `verify` |
| im Index nur in **Prosa**, nicht als Zelle | 2 — `doc-commits`, `doc-tracked` |
| nur im Index | **0** |

**Fünf Targets brauchten eine Zeile**, nicht acht und nicht drei: die drei
fehlenden plus die zwei, die im Index nur im Fließtext standen — für den
`targets`-Sensor ist Prosa kein Eintrag, und mit dem Wegfall der zweiten Tabelle
wären sie `gate-undocumented` geworden.

### 3.2 Der Umzug

[`AGENTS.md`](../../../../AGENTS.md) §4: **7094 → 1170 Zeichen**. Was bleibt, ist
die Regel (*kein Target nennen, das im Makefile nicht existiert — auch nicht in
Prosa*), der Zeiger auf den Index, die Nennung der maschinellen Hälfte
(`make doc-targets`, beide Richtungen, **eine** Autoritäts-Doku) und die
Aussage, was *mandatory* heißt. Die Liste ist weg.

[`harness/README.md`](../../../../harness/README.md) §Sensors trägt die fünf
neuen Zeilen. Der Prosa-Absatz *„Nicht hier, obwohl sie in keinem Aggregat
hängen"* nennt jetzt fünf Targets und sagt, dass sie **oben in der Tabelle**
stehen — vorher behauptete er das Gegenteil für zwei von ihnen.

Drei Konfigurationen zeigen auf einen Index: `targets.doc-tables`,
`targets.authority`, `mentions.documents`.

### 3.3 `mentions` konnte dem Index folgen — die Unmöglichkeit war ungemessen

[`harness/sensors/doc-mentions.md`](../../../../harness/sensors/doc-mentions.md)
führte seit slice-184 als Grenze 1: Das Modul sucht den **vollen repo-relativen
Pfad**, `harness/README.md` verlinkt geschwister-relativ und **könne das nicht
ändern**, ohne seine Links zu brechen. Die erste Hälfte stimmt und ist gemessen
(`resolve-from`, `match-basename`, `strip-prefix`, `paths-relative-to` sind alle
**unbekannte Felder** — vier Fehlversuche gegen das Modul). Die zweite war eine
**Annahme**.

**Gegenprobe in beide Richtungen**, an einem Wegwerf-Repo mit einem Artefakt und
einem Dokument: Link als `sensors/alpha.md` ⇒ *0 von 1 erwähnt*, Exit 2. Link als
`../harness/sensors/alpha.md` ⇒ *1 von 1*, Exit 0 — **und der Link löst weiter
auf**, weil er über die Repo-Wurzel geht. Im Bestand: 16 Links umgestellt,
`mentions: 16 von 16`.

**Der Preis steht in der Sensor-Datei**, weil er sonst niemandem auffällt: Wer
die Links auf die kürzere Form zurücksetzt, macht den Sensor blind, **ohne dass
ein Link bricht**. Und Grenze 4 (die 39 ADRs) ist damit kein technisches
Hindernis mehr, sondern eine Abwägung.

### 3.4 Die Zellengrenze entfällt — und das Modul fällt dabei fail-closed aus

Die `structure`-Zellengrenze auf der *Zweck*-Spalte von §4 verlor mit der Tabelle
ihren Gegenstand und ist entfernt; die Zellengrenze der Sensors-Tabelle deckt ihn
allein.

**Die erste Fassung dieses Abschnitts behauptete, die Regel sei danach grün
geblieben — das war falsch, und der Review hat es widerlegt** (F-1). Testweise
wortgleich wieder eingesetzt, meldet der Lauf
`AGENTS.md:166 :: Spalte Zweck section-column-missing`, **Exit 2**. Grün war
nicht der Prüfer, sondern **mein Lauf**: `doc-structure` hängt in `verify`, nicht
in `gates` — ich hatte nach dem Streichen der Tabelle `make gates` gefahren und
das Ergebnis auf den falschen Prüfer bezogen.

**Die Lehre ist die unangenehmere Fassung derselben Regel, die §3.1 trägt:** Eine
Messung nennt nicht nur ihr Ergebnis, sondern **ihren Aufruf**. Ein grüner Lauf
belegt nur, was in ihm hing.

### 3.5 Drei Mutations-Proben, alle rot

| Probe | Meldung |
|---|---|
| Zeile für `commit-scope-check` aus dem Index entfernt | `Makefile:113 commit-scope-check gate-undocumented — Makefile-Regel ohne Deklaration in der Autoritäts-Doku harness/README.md` |
| Zeile `make gibtsnicht` in den Index eingefügt | `harness/README.md:101 gibtsnicht gate-phantom — dokumentiertes Target ohne Makefile-Regel` |
| einen Sensor-Link auf die Geschwister-Form zurückgesetzt | `harness/sensors/doc-planning.md:1 artifact-unmentioned` · `mentions: 15 von 16` |

Alle drei mit **genannter Meldung**, nicht nur mit Exit-Code; danach
zurückgesetzt und `make gates` erneut grün.

### 3.6 Was der unabhängige Review verändert hat

Der Report
([`2026-09-08-slice-193-…`](../../../reviews/2026-09-08-slice-193-gate-index-steht-einmal.md))
trug **6 HIGH · 5 MEDIUM · 4 LOW · 1 INFO**. Er hat den Umzug selbst
**bestätigt** — Differenz beidseitig eigenständig nachgemessen, null verloren,
genau die fünf neuen Zeilen; die drei Mutations-Proben echt; der `mentions`-Umweg
in beide Richtungen belegt; §4 ohne Auslassung die Ziel-Form. Und er hat vier
merge-blockierende Dinge gefunden:

| Befund | Kern | Behebung |
|---|---|---|
| **F-1** Der zentrale „Fund" war falsch | Die entfallene Zellengrenze meldet **nicht** grün, sondern `section-column-missing`, Exit 2. Grün war `make gates`, in dem `doc-structure` gar nicht hängt | §3.4, §8, §9 und der `.d-check.yml`-Kommentar richtiggestellt — die Konfigurations-Zeile war die langlebigste der drei Stellen |
| **F-2/F-4/F-5** Der Umzug war gemacht, die Verweise nicht | `Makefile` Z. 133 und die `make help`-Zeile, `harness/conventions.md`, `docs/user/releasing.md` führten `AGENTS.md` §4 weiter als Deklarationsort; die **neu geschriebene** Grenze 4 in `doc-mentions.md` nannte für den ADR-Index `../../docs/plan/adr/…` — das löst nach `docs/docs/…` auf | alle vier nachgezogen; die ADR-Form auf `../../../` korrigiert und **in beide Richtungen gegengeprobt** (`../../` ⇒ `target-missing`, `../../../` ⇒ Exit 0) |
| **F-3** Zwei überholte Begründungsblöcke | `.d-check.yml` trug unmittelbar über den geänderten Zeilen die alte Begründung (*„BEIDE Doku-Tabellen"*, *„WARUM NUR AGENTS.md"*) | beide ersetzt |
| **F-6** Ein 3×-Eintrag ohne Ausgang | Der Ausgang war auf *„den nächsten Lese-Schritt"* vertagt — in einem Repo ohne Wellen **ist** die Slice-Closure der Lese-Schritt | Ausgang gezogen: *verkörpert*, dritte Mess-Regel (§8) |

**F-1 ist der Befund über diesen Slice.** Er hatte einen eleganten Fund zu
erzählen — *ein Prüfer verliert seinen Gegenstand und meldet weiter grün* — und
die Belegkette dafür nie gefahren. Was grün war, war ein Aggregat, in dem der
Prüfer nicht hängt. **Der eigene Skill nennt genau diese Frage** (*„war sie rot,
und woran?"*), und sie war auf die Behauptung nicht angewandt, weil die
Behauptung eine über einen **grünen** Lauf war.

**F-2 ist die Wiederholung von slice-192.** Dort hieß der Befund *„eine
Massen-Ersetzung trifft auch die Aussage, die stehenbleiben muss"*; hier ist es
die Gegenrichtung — **eine Aussage, die mitgezogen werden musste und liegenblieb**.
Beide Male ist der Handgriff derselbe: nach dem Umbau `grep` auf die alte
Adresse, und zwar über **alle** Dateiarten, nicht nur die Doku. **Genau so
gefahren fand er zwei weitere**, die auch der Review nicht als HIGH führte:
`.claude/hooks/pretooluse-command-guard.sh` und `tools/gate-consistency.sh`
nannten `AGENTS.md` §4 in ihren Kopfkommentaren als Deklarationsort. Beide
nachgezogen — der `grep` über `*.sh` war der Unterschied.


## 4. Definition of Done

- [x] [`AGENTS.md`](../../../../AGENTS.md) §4 trägt **keine Tabelle** mehr,
      sondern die Regel und den Zeiger; jedes dort bisher gelistete Target ist
      in [`harness/README.md`](../../../../harness/README.md) eingeordnet —
      **belegt durch eine Differenz-Messung vorher/nachher**, nicht durch
      Augenschein.
- [x] Die drei Konfigurationen zeigen auf **einen** Index: `targets.doc-tables`,
      `targets.authority`, `mentions.documents`. Die Zellengrenze auf der
      *Zweck*-Spalte entfällt oder wandert mit.
- [x] Eine **Mutations-Probe** je umgezogenem Prüfer war **rot**, mit genannter
      Meldung: ein Target ohne Eintrag und ein Eintrag ohne Target.
- [x] Unabhängiger Review durchgeführt (Report unter [`docs/reviews/`](../../../reviews/README.md)) — 16 Findings, alle abgearbeitet (§3.6).
- [x] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [x] Beobachtungs-Register fortgeschrieben.
- [x] Jedes Risiko aus §7 trägt einen Ausgang.

`make gates` und `make verify` grün. Zwei öffentliche Verträge sind berührt:
`AGENTS.md` ist Rang 8, `harness/README.md` Rang 9.

## 5. Trigger

**Start** (`open` → `in-progress`):
[slice-192](../done/wellenlos/slice-192-baseline-v660-vendoring.md) liegt in `done/` und
das WIP-Limit ist frei.

**Rückführungen — vorab benannt:**

- `in-progress` → `next` (zu groß): Zeigt die Differenz-Messung mehr als ein
  Dutzend Targets ohne Gegenstück, wird der Umzug vom Konfigurations-Umbau
  getrennt.
- `in-progress` → `open` (blockiert): Verlangt `targets` zwei Autoritäts-Dateien
  und kann nur eine, ist die Ablösung ein CR an `d-check` — dann bleibt die
  Doppelung stehen, sichtbar als benannte Abweichung.

## 6. Closure-Trigger

DoD vollständig, `make gates` und `make verify` grün, Closure-Notiz mit
Lerneintrag.

## 7. Risiken und offene Punkte

- **Ein Target kann beim Umzug verschwinden.** 41 gegen 28 Zeilen: Wer die
  Tabelle streicht, ohne die Differenz zu messen, verliert Deklarationen still —
  und `make doc-targets` prüft danach gegen einen Index, in dem sie fehlen.
  — **Ausgang:** *eingetreten* → im Slice aufgelöst, und die **Klasse** liegt im
  Beobachtungs-Register. Es traf **die Messung**, nicht den Umzug: Der erste
  Zähler las je Zelle nur ein `make X` und übersah eine Sammelzeile mit vier
  advisory-Targets (§3.1). Wäre die Zahl acht geblieben, hätten drei Targets eine
  überflüssige Zeile bekommen und zwei — `doc-commits`, `doc-tracked` — **keine**,
  weil sie im Index nur in Prosa standen. Die Klasse ist
  [`BEO-PLAN/kandidaten-klassifikation-groeber-als-der-kandidat`](../observations/BEO-PLAN/kandidaten-klassifikation-groeber-als-der-kandidat/observation.md);
  Beleg dort, damit sie beim nächsten Mal gezählt ist — und der Eintrag erreicht
  damit 3× und trägt seinen Ausgang aus dieser Closure (§8).
- **`mentions` hängt an `AGENTS.md`.** Zieht die Konfiguration um, ohne dass die
  Sensor-Dateien in `harness/README.md` genannt sind, meldet der Lauf gegen den
  Bestand — oder, schlimmer, er meldet nichts, weil die Prüfmenge leer wird.
  — **Ausgang:** *entfallen*, gestrichen mit Begründung: Der Umzug ist gelungen,
  und die Sperre gegen die leere Prüfmenge ist im Modul selbst fail-closed
  (`mentions.documents […] trifft kein Dokument`). Der Weg dorthin war ein
  **Fund**, kein Risiko-Fall: Die Sensor-Datei erklärte die Kopplung seit
  slice-184 für unmöglich, und die Gegenprobe zeigt sie in zwei Zeilen (§3.3).
  Was bleibt, steht als **Preis** in der Sensor-Datei: Die kürzere Link-Form
  macht den Sensor blind, ohne einen Link zu brechen.
- **Die Zusage „ein Index" ist selbst eine Deklaration ohne Sensor.** Nichts
  hindert einen künftigen Lauf daran, wieder eine Tabelle in `AGENTS.md` zu
  schreiben. — **Ausgang:** *weiter offen* → **Beobachtungs-Register**,
  [`BEO-GATE/zusage-weiter-als-ihre-durchsetzung`](../observations/BEO-GATE/zusage-weiter-als-ihre-durchsetzung/observation.md)
  (2×, kein neuer Beleg — siehe unten). **Teilweise gedeckt, und die Grenze ist
  benennbar:** `targets.authority` kennt genau **eine** Datei; eine zweite
  Gate-Tabelle in `AGENTS.md` würde von `doc-targets` **nicht** gelesen und
  liefe damit sofort auseinander — sichtbar würde das aber erst, wenn jemand
  hinsieht. Ein Sensor auf *„dieses Dokument enthält keine Gate-Tabelle"* wäre
  ein Muster auf eine Abwesenheit; ob er trägt, ist eine eigene Entscheidung.
  **Kein Beleg an dem Eintrag**, weil hier keine Zusage weiter reicht als ihr
  Prüfer — die Zusage *hat* hier gar keinen, und das steht so da.

## 8. Closure-Notiz

**Lerneintrag — Form: geschärfte Regel.** *Eine Doppelung kostet nicht doppelt
so viel wie eine Stelle — sie kostet die Stelle, plus jeden Lauf, der beide
liest, plus den Tag, an dem sie auseinanderlaufen.* Der Gate-Index stand in zwei
Dateien, die **beide** in jedem Lauf-Kontext liegen. Die Baseline `v6.6.0` nennt
genau das als Grund, und der Bestand belegt ihn: `AGENTS.md` §4 führte **41**
Zeilen, der Index **39** Nennungen — und die Differenz war kein Zufall, sondern
Drift, für deren Nachziehen keine Pflicht existierte. §4 misst jetzt **1170**
statt 7094 Zeichen; keine Deklaration ist verloren gegangen.

**Der zweite Lerneintrag ist der wertvollere, und er widerlegt eine eigene
Behauptung:** *Eine für unmöglich erklärte Kopplung gehört gemessen, bevor sie
zur Grenze wird.*
[`harness/sensors/doc-mentions.md`](../../../../harness/sensors/doc-mentions.md)
führte seit slice-184, `harness/README.md` **könne** die Sensor-Dateien nicht so
nennen, wie das Modul sie sucht, ohne seine Links zu brechen. Die Prämisse war
richtig (das Modul kennt kein `resolve-from` — vier Feldnamen durchprobiert, alle
unbekannt), die Folgerung nicht: Die Link-Form `../harness/sensors/<name>.md`
löst von `harness/README.md` **gleich auf** und enthält den vollen Pfad als
Teilzeichenkette. Zwei Läufe an einem Wegwerf-Repo hätten das 2026 schon
gezeigt — *0 von 1* gegen *1 von 1*. **Die Grenze war nicht falsch belegt,
sondern ungemessen**, und sie stand ein halbes Dutzend Slices lang als Tatsache
in einer Sensor-Datei.

**Zwei beobachtbare Closure-Kriterien.** (1) **Drei Mutations-Proben, alle rot,
je mit genannter Meldung** (§3.5): `gate-undocumented`, `gate-phantom`,
`artifact-unmentioned` — danach zurückgesetzt, `make gates` grün. Das ist der
Beleg, dass der umgezogene Index tatsächlich gewächtert wird und nicht nur
grün meldet. (2) Die Differenz-Messung ist **beidseitig** null: kein Target steht
nur in `AGENTS.md`, keines nur im Index; `mentions: 16 von 16`.

**Und ein dritter Lerneintrag, den der Review erzwungen hat** (F-1): *Ein grüner
Lauf belegt nur, was in ihm hing.* Dieser Slice behauptete, die entfallene
Zellengrenze habe nach dem Wegfall ihres Gegenstands **grün** gemeldet — ein
schönes Beispiel für einen unkalibrierten Prüfer, und **falsch**. Sie fällt
fail-closed aus (`section-column-missing`, Exit 2); grün war `make gates`, in dem
`doc-structure` gar nicht hängt. Der Satz stand dreifach: im Plan, in der
Closure-Notiz und **dauerhaft als Kommentar in `.d-check.yml`** — die
langlebigste der drei Stellen. Alle drei sind korrigiert. **Eine Messung nennt
ihren Aufruf, nicht nur ihr Ergebnis**; genau das verlangt die Mess-Regel
*Geltungsbereich*, und sie war hier auf den eigenen Lauf nicht angewandt.

**Beobachtungs-Register.** Erhöht:
[`BEO-PLAN/kandidaten-klassifikation-groeber-als-der-kandidat`](../observations/BEO-PLAN/kandidaten-klassifikation-groeber-als-der-kandidat/observation.md)
auf **3×** — die Differenz-Messung selbst zählte eine Sammelzelle mit vier
Targets als eines und hätte zwei Targets ohne Index-Zeile gelassen (§3.1).
**Der Ausgang ist gezogen, hier und jetzt: *verkörpert*.** Die Zusage *„wer eine
Menge zählt, zählt sie zweimal verschieden"* steht als dritte Mess-Regel in
[`AGENTS.md`](../../../../AGENTS.md) §5, die Herleitung mit allen drei
Ausprägungen im Reviewer-Skill — dieselbe Teilung wie bei den zwei älteren.
**Das war zunächst auf „den nächsten Lese-Schritt" vertagt; der Review hat das
korrigiert** (F-6): a-check fährt wellenlos, und dann **ist** die Slice-Closure
der Lese-Schritt (`modul-06`, Tabelle *Träger im Repo ohne Wellen*). Ein Eintrag,
der sie ohne Ausgang übersteht, ist genau das, was die Regel verbietet — und
`make verify` hätte es nicht gemeldet, weil `verify-observations` nur Deckung
prüft.

## 9. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen.** Zugeschnitten **nach dem, was der Diff
anfasst** (`git diff --name-only`), nicht nach der Deklaration — der Handgriff
aus [slice-192](../done/wellenlos/slice-192-baseline-v660-vendoring.md) §3.6,
F-3. Berührte Verzeichnisse: `AGENTS.md`, `.d-check.yml`, `harness/`,
`harness/sensors/`, `tools/archive-wave/`, `docs/plan/carveouts/`,
`docs/plan/planning/`. Daraus:

- **`HARNESS`** (`AGENTS.md`, `harness/`) — Achsen 1,2,3. Der Kern.
- **`GATE`** (`.d-check.yml`, `tools/`) — Achsen 1,2,3. Drei Konfigurationen und
  eine entfallene Regel.
- **`PLAN`** (`docs/plan/planning/`) — Achsen 1,2,3. Plan und Register.
- **`ADR`** ist **nicht** berührt: `docs/plan/carveouts/README.md` liegt nicht in
  der ADR-Familie; die eine Zeile dort zieht einen Zeiger nach.

**Modus:** alle drei Greenfield; ein Begründungsblock entfällt.

**Vorgelagert — offene Beobachtungen sichten.** **Zwei Quellen**, Register nach
**allen** drei Kürzeln:

| Eintrag | Kürzel | Berührung |
|---|---|---|
| [`kandidaten-klassifikation-groeber-als-der-kandidat`](../observations/BEO-PLAN/kandidaten-klassifikation-groeber-als-der-kandidat/observation.md) | `PLAN` | **3× erreicht, Ausgang gezogen** — *verkörpert* als dritte Mess-Regel (§8); der Lese-Schritt ist diese Closure |
| [`zusage-weiter-als-ihre-durchsetzung`](../observations/BEO-GATE/zusage-weiter-als-ihre-durchsetzung/observation.md) | `GATE` | **berührt, kein Beleg** — die Zusage *„ein Index"* hat gar keinen Prüfer, statt einen zu knappen; siehe §7 |
| [`pruefer-ohne-gegenstand-oder-aufruf`](../observations/BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf/observation.md) | `GATE` | **kein Treffer, entgegen der ersten Buchung** — die Zellengrenze fällt fail-closed aus, nicht grün (§3.4, Review F-1). Die Buchung stützte sich auf einen Lauf, in dem der Prüfer gar nicht hing |
| [`muster-trifft-nur-die-haeufige-schreibweise`](../observations/BEO-GATE/muster-trifft-nur-die-haeufige-schreibweise/observation.md) | `GATE` | **bedient** — die Gegenprobe am Fund ist genau das, was §3.1 gerettet hat |
| [`vollstaendigkeits-haken-ohne-erschoepften-gegenstand`](../observations/BEO-PLAN/vollstaendigkeits-haken-ohne-erschoepften-gegenstand/observation.md) | `PLAN` | **bedient** — jeder Ausgang nennt seine Ebene, und die DoD verlangt die Differenz-Messung statt Augenschein |
| [`messung-ohne-reproduzierbares-instrument`](../observations/BEO-PLAN/messung-ohne-reproduzierbares-instrument/observation.md) | `PLAN` | **bedient** — §3.1 und §3.5 nennen Zählweise und Meldungen, nicht nur Ergebnisse |
| [`chronik-in-gelesenen-dateien`](../observations/BEO-HARNESS/chronik-in-gelesenen-dateien/observation.md) | `HARNESS` | **berührt, nicht erhöht** — der neue §4-Text nennt keine vorige Fassung; die Chronik steht hier im Plan |
| [`massen-ersetzung-trifft-die-historische-aussage`](../observations/BEO-HARNESS/massen-ersetzung-trifft-die-historische-aussage/observation.md) | `HARNESS` | **bedient** — die drei `AGENTS.md`-§4-Verweise außerhalb wurden einzeln gelesen, nicht ersetzt |

**Nur einer erreicht 3×**, und sein Ausgang ist **in dieser Closure gezogen**,
nicht vertagt (§8). Die übrigen sind **bedient oder berührt**, nicht neu belegt.

**Eine Zeile dieser Tabelle war falsch und ist korrigiert** (Review F-1): Die
Buchung zu `pruefer-ohne-gegenstand-oder-aufruf` stützte sich auf eine Messung,
die den Prüfer nie aufgerufen hatte. Sie steht hier stehen gelassen und
richtiggestellt statt still entfernt — eine gestrichene Zeile sähe aus, als wäre
die Frage nie gestellt worden.

**(2) Report des Vorgängers** — zu slice-192, mit ihm archiviert. Seine schärfste
Klasse für diesen Slice war *„eine Regel zu zitieren ist nicht, sie anzuwenden"*:
Deshalb ist die Sub-Area-Wahl oben **aus `git diff --name-only` abgeleitet** und
nicht aus der Deklaration — der Handgriff, den jener Report als fehlend benannt
hat.

**Keine Treffer sind ebenfalls eine Antwort:** Zu `SPEC`, `KERN`, `ADAPT` und
`REVIEW` steht nichts Offenes an, das dieser Slice berührt.
