# slice-192 — Baseline auf `v6.6.0` heben: Vendoring, Pins, MR-Zeiger

**Welle:** ohne Welle — die Closure-Bedingung wäre die eigene DoD.

**Bezug:** Maintainer-Meldung „es gibt eine neue Regelwerks-Version `v6.6.0`".
[`AC-QA-02`](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze),
[`AC-QA-03`](../../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit).

**Berührte Spec-Stellen:** — · Der Slice berührt kein Spec-Stratum.

**Verantwortlich:** Claude — gesetzt beim Übergang nach `in-progress/`.

**Autor:** Claude. **Datum:** 2026-09-08.

**Lerneintrag — Form:** geschärfte Regel.

---

## 1. Ziel und Abgrenzung

**Ziel:** Der adoptierte Stand ist `v6.6.0`. Das Bundle liegt vendored im
Baseline-Verzeichnis unter seinem Tag, `v6.5.0` ist entfernt, **jede** Nennung
des alten Standes außerhalb der eingefrorenen Artefakte ist behandelt (§3.2), und
für jeden aktiven `MR`-Eintrag ist **gemessen**, ob sein
`Ersetzt-Baseline-Regel`-Zeiger mitwandern darf.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Die Adoption der einen inhaltlichen Neuerung** — *„Der Gate-Index steht
  einmal"* (§2). *Ein Folge-Slice übernimmt es* mit Kennung
  ([slice-193](../open/slice-193-gate-index-steht-einmal.md)): Sie streicht eine
  41-zeilige Tabelle aus [`AGENTS.md`](../../../../AGENTS.md) §4 und zieht drei
  Konfigurationen um. Das ist ein Umbau, kein Pin-Bump; zusammen wäre der Diff
  in einer Review-Sitzung nicht prüfbar.
- **Der Voll-Abgleich der 15 Ziel-Form-Paare.** *Ein Folge-Slice übernimmt es:*
  [slice-188](../open/slice-188-voll-abgleich-gate-und-skill.md) und
  [slice-189](../open/slice-189-voll-abgleich-spec-straten.md) tragen ihn
  ohnehin und laufen nach diesem Slice gegen `v6.6.0` statt zweimal.
  [`harness/conventions.md`](../../../../harness/conventions.md) §Baseline
  verlangt beim Sprung **Delta *und* Voll-Abgleich**; das Delta steht in §2, der
  Voll-Abgleich hat seine Adresse.
- **Zwei vendored Stände nebeneinander stehen lassen.** *Bestand bleibt nicht
  stehen:* §Baseline sagt, genau **ein** Stand liegt vendored; mehrere sind nur
  während einer Migration zulässig. Der alte geht mit diesem Slice.

## 2. Delta-Analyse `v6.5.0` → `v6.6.0` (2026-09-08)

**Gemessen** gegen den lokalen Kurs-Checkout, `git diff v6.5.0 v6.6.0 --
lab/regelwerk lab/templates`: **10 Dateien, +87/−32**. Kurs-Welle 129,
2026-09-08. Thema: *„Der Gate-Index steht einmal"*.

| Datei | Änderung | Trifft a-check |
|---|---|---|
| `templates/AGENTS.template.md` | §4 **verliert die Tabelle**; nur noch Regel + Zeiger auf `harness/README.md` §Sensors | **ja, groß** — §4 trägt 41 Zeilen, 7094 Zeichen |
| `templates/harness/README.template.md` | *„DIES IST DER EINZIGE GATE-INDEX"*; Target-Zelle = **nackter Name** (Argument in die Nachbarspalte, Link erlaubt) | **ja** für den Index-Anspruch; **nein** für die Zellen-Form — gemessen: null Zellen mit Argument in der Code-Span, 16 verlinkte je Tabelle |
| `templates/.d-check.yml` | `targets`-Beispiel mit `doc-tables: [harness/README.md]` und `authority: harness/README.md`; dazu *„Aktivieren heißt zwei Schritte"* | **ja** — a-check führt `doc-tables: [AGENTS.md, harness/README.md]`, `authority: AGENTS.md` |
| `templates/Makefile` | Kopfkommentar nennt nur noch `harness/README.md` §Sensors | ja, klein |
| `regelwerk/grundlagen-harness-dateien.md` | +17: der Gate-Index steht einmal, **mit Begründung** (beide Dateien liegen in jedem Lauf-Kontext, ein zweiter Index wird pro Lauf zweimal bezahlt); Target-Zelle nackt | ja — die Begründung ist die tragende |
| `regelwerk/modul-13-quality-gates.md` | +20: die Hard Rule hat eine **maschinelle Hälfte** (Deklarations-Sensor, beide Richtungen); Ausnahmeliste **namentlich, nie Glob** | teils erfüllt — `make doc-targets` läuft, `exempt-targets` ist namentlich |
| `regelwerk/modul-09-implementierung.md` | Ziel-Form-Satz: „Gate-**Regel** und Zeiger" statt „Gate-Tabelle" | Folge von oben |
| `regelwerk/modul-02-harness-bootstrap.md` | Bootstrap-Tabelle: `AGENTS.md` §4 fällt als Phasen-Träger weg | nein — Bootstrap ist durch |
| `regelwerk/modul-15-observability.md` | Drift-Regeln reden vom Gate-Index statt von `AGENTS.md` | nein — [`MR-014`](../../../../harness/conventions.md#mr-014) nimmt das Modul aus |
| `regelwerk/README.md` | Stand-Zeile | ja, mechanisch |

**Ein einziger inhaltlicher Punkt trifft a-check**, und er trifft hart: der
Gate-Index. Alles andere ist Pin-Arbeit. **Die Ironie gehört benannt:**
[slice-187](../done/wellenlos/slice-187-voll-abgleich-erstdurchgang-rest.md)
hat gerade den Satz *„Diese Tabelle listet auf; die Bindung steht in
`harness/README.md` §Sensors"* nach §4 übernommen — aus der `v6.5.0`-Ziel-Form,
in der er seit `v5.12.0` stand. `v6.6.0` löscht ihn zusammen mit der Tabelle.
Die Übernahme war **richtig gegen den adoptierten Stand** und ist mit dem
nächsten überholt; das ist kein Fehler des Abgleichs, sondern der Normalfall
eines Repos, das einer Baseline folgt.

## 3. Umsetzung

### 3.1 Das Bundle entsteht netzlos aus dem Tag — und das Verfahren ist belegt

Der Kurs liegt als lokaler Checkout vor. Gebaut wird **nicht** aus dessen
Arbeitsbaum (er trägt eine uncommittete Änderung, §7), sondern aus einem
`git archive` des Tags in ein Wegwerf-Verzeichnis; dort läuft
`tools/build-bundle.sh <ziel> <tag>` des Kurses — dasselbe Skript, das der
Release-Workflow fährt.

**Die Gegenprobe ist der Beleg, nicht die Annahme:** Dasselbe Verfahren, auf den
**alten** Tag angewandt, reproduziert den bisher vendorten Baum **byte-gleich** —
`diff -rq` liefert null Unterschiede, und das nachgebaute `SHA256SUMS` stimmt mit
dem committeten überein. Erst dadurch ist belegt, dass der neue Baum auf demselben
Weg entstanden ist wie der alte, und nicht auf einem ähnlichen.

**Das Manifest** entsteht mit demselben Aufruf, mit dem `make regelwerk-check` es
prüft (`find … | sort | xargs sha256sum`, `SHA256SUMS` ausgenommen) — auch das an
`v6.5.0` gegengeprobt: byte-gleich. **54** Dateien.

### 3.2 Pins, Symlinks, und was *nicht* mitgezogen wurde

**Gemessen gegen den Stand vor dem Swap** (`git show <swap>^:<datei> | grep -c`
über die im Swap-Commit berührten Nicht-Baseline-Dateien), nicht geschätzt:

| Gegenstand | Dateien | Nennungen | Behandlung |
|---|---|---|---|
| trugen den alten Pin | **17** | **51** | — |
| davon **gebumpt** | 15 | 48 | auf den neuen Stand |
| davon auf **Kennung** umgestellt | 2 | 3 | eingefroren, §3.3 |
| Baseline-Symlinks unter `.claude/rules/` | 4 | 4 | umgehängt; `make symlink-check` grün |
| Stand-Zeile in [`harness/conventions.md`](../../../../harness/conventions.md) §Baseline | 1 | 1 | Kurs-Welle **129 · 2026-09-08**, aus dem Kopf des vendorten `regelwerk/README.md` |
| eingefrorene Artefakte, die den alten Stand **nennen** | viele | — | **unangetastet** — sie beschreiben, was damals galt; `versions` nimmt sie über `exempt-paths` aus |

**Zwei eingefrorene Einträge trugen einen *Link*, keine Kennung** — und das ist
der Fund dieses Slice (§3.3).

### 3.3 Der Fund: „der Zeiger bleibt" und „genau ein Stand" schließen einander aus

[`harness/conventions.md`](../../../../harness/conventions.md) §Baseline sagt
zweierlei, das zusammen nicht geht:

1. *Genau **ein** Stand liegt vendored.*
2. *Aufgelöste Einträge sind ausgenommen — ihr Zeiger **bleibt** auf dem Stand,
   gegen den sie damals formuliert wurden.*

Gemessen: **zwei** aufgelöste Adaptionen — [`MR-011`](../../../../harness/conventions.md#mr-011)
und [`MR-021`](../../../../harness/conventions.md#mr-021) — trugen ihren
`Ersetzt-Baseline-Regel`-Zeiger als **Markdown-Link** in den vendorten Baum. Beide
Aussagen zugleich befolgt heißt: ein garantiert toter Link, sobald der alte Stand
geht. Der Abschnitt behauptet daneben, die Kollision sei *„bezahlt"* — belegt mit
den **null Nachzügen** beim vorigen Sprung. Diese Messung galt den vier
einfrierenden Ziel-Formen (Review-Report, Ergebnisnotiz, zwei Archiv-Stubs);
**`conventions/done/` war nie darin.**

**Aufgelöst wie bei jenen vieren: Kennung statt Adresse.** Beide Zeiger lauten
jetzt `` `v6.5.0` · `regelwerk/grundlagen-source-precedence.md` §ID-Schema als
Klammer `` — der genannte Stand bleibt **derselbe**, nur die Adresse fällt weg.
Das ist strikt weniger Eingriff als ein mitwandernder Zeiger: Die Aussage des
Eintrags — *gegen diesen Stand wurde formuliert* — bleibt wortgleich erhalten,
und die Immutabilität ist nicht berührt, weil kein Inhalt geändert wird.

### 3.4 Die fünf aktiven MR-Zeiger — gemessen, nicht angenommen

§Baseline verlangt, dass ein mitwandernder Zeiger nur dann wandert, wenn der
referenzierte Abschnitt im neuen Stand **wortgleich** ist. Gemessen wurde in zwei
Stufen, beide ohne die `<!-- Quelle: -->`-Zeile (sie trägt den Tag und ändert sich
bei **jedem** Sprung, ohne dass ein Wort anders wird):

| MR | Zielabschnitt | Datei im Delta? | Abschnitt |
|---|---|---|---|
| [`MR-012`](../../../../harness/conventions.md#mr-012) | `grundlagen-referenz-richtung.md` §Referenz-Richtung (SDP) | nein | wortgleich |
| [`MR-014`](../../../../harness/conventions.md#mr-014) | `modul-15-observability.md` §Kernidee | **ja** | **wortgleich** (554 Zeichen, unverändert) |
| [`MR-015`](../../../../harness/conventions.md#mr-015) | `modul-06-roadmap.md` §Wellen-Closure-Prozedur | nein | wortgleich (10 341 Zeichen) |
| [`MR-016`](../../../../harness/conventions.md#mr-016) | `modul-08-agentenrollen.md` §Die neun Übergaben | nein | wortgleich (1372 Zeichen) |
| [`MR-022`](../../../../harness/conventions.md#mr-022) | `grundlagen-source-precedence.md` §ID-Schema als Klammer | nein | wortgleich (4593 Zeichen) |

**Alle fünf dürfen wandern.** Der einzige Zeiger in eine geänderte Datei ist
[`MR-014`](../../../../harness/conventions.md#mr-014), und die Änderung liegt
in §Doku-Konsistenz-Drift-Regeln, nicht in
§Kernidee — genau die Unterscheidung, die die Regel verlangt und die eine
Datei-Ebenen-Messung verfehlt hätte.

### 3.5 Die Delta-Zahl präzisiert — und was sie für slice-188 bedeutet

§2 nennt **10** geänderte Dateien nach `git diff`. Normalisiert (ohne die
Tag-Zeile) sind es **neun echte**; die zehnte,
`templates/harness/conventions.template.md`, unterscheidet sich **nur** im
gepinnten Release-Asset-Link.

**Folge für die laufende Abgleich-Kette, und sie ist konkret:** Von den neun
echten Änderungen liegen **drei in Vorlagen, die slice-188/189 als Paar
führen** — `templates/.d-check.yml` (+24 Zeilen: der `targets`-Block mit
*einem* Gate-Index und der Hinweis *„Aktivieren heißt zwei Schritte"*),
`templates/AGENTS.template.md` und `templates/harness/README.template.md`.
Damit ist **slice-188s Ausgangsmessung überholt**: Die dort genannten *42
Kandidaten* für `.d-check.yml` sind gegen den alten Stand gemessen. Der Slice
muss neu messen, bevor er liest; sein §2 zitiert das Instrument aus slice-187,
also ist das ein Aufruf, kein Umbau.

### 3.6 Was der unabhängige Review verändert hat

Der Report
([`2026-09-08-slice-192-…`](../../../reviews/2026-09-08-slice-192-baseline-v660-vendoring.md))
trug **3 HIGH · 5 MEDIUM · 2 LOW · 2 INFO**. Er hat den **Bau-Weg vollständig
nachgefahren** — beide Bundles neu gebaut, beide byte-gleich, beide Manifeste
reproduziert — und die zweistufige Zeiger-Messung bestätigt. Gefunden hat er
drei Dinge, die kein Lauf fängt:

| Befund | Kern | Behebung |
|---|---|---|
| **F-1** Historische Nennungen mitgehoben | Der Massen-`sed` traf **sieben** Stellen im Plan, an denen der alte Stand eine **Tatsache** ist: §1 behauptete *„`v6.6.0` ist entfernt"*, §2 hieß *„Delta-Analyse `v6.6.0` → `v6.6.0`"*, und das dort genannte Instrument `git diff v6.6.0 v6.6.0` liefert leere Ausgabe | alle sieben zurückgesetzt; neu im Register: [`BEO-HARNESS/massen-ersetzung-trifft-die-historische-aussage`](../observations/BEO-HARNESS/massen-ersetzung-trifft-die-historische-aussage/observation.md) |
| **F-2** Das Closure-Log verfälscht | Dieselbe Ersetzung machte aus `welle-15`s Migration `v6.2.0`→`v6.5.0` eine `v6.2.0`→`v6.6.0` — im Widerspruch zur Ergebnisnotiz, auf die die Zeile zeigt | zurückgesetzt |
| **F-3** §9 erklärte `PLAN` und `REVIEW` für unberührt | ...obwohl der Diff in beiden Verzeichnissen arbeitet und §9 zwölf Zeilen höher selbst eine offene `BEO-PLAN`-Klasse als eingebaut führt — **die Wiederholung von slice-187 F-2**, aus dem dieser Slice gelernt haben wollte | §9 korrigiert |

**F-1 und F-2 sind derselbe Handgriff, und er ist die Lehre dieses Slice:** Ein
Versions-Wechsel ist eine Ersetzung von **Zeigern**, nicht von **Kennungen**.
Wer mit der nackten Kennung ersetzt, trifft jeden Satz, in dem der alte Stand
eine Tatsache ist — und am dichtesten stehen solche Sätze in dem Dokument, das
den Wechsel beschreibt. **Die billige Vorbeugung stand die ganze Zeit da:**
Ersetzt man mit dem **Muster des Versions-Sensors** (nur `<baseline>/<tag>/`),
ist keine der sieben Stellen betroffen.

**F-3 ist die unbequemste.** Der Slice hat die Lehre aus slice-187 in §9
**zitiert** — *„das Register nach allen berührten Kürzeln"* — und beim Ziehen der
Konsequenz dieselbe Grenze gezogen wie sein Vorgänger: nach Sub-Area-Deklaration
statt nach dem, was der Diff anfasst. **Eine Regel zu zitieren ist nicht,
sie anzuwenden**; das ist der zweite Lerneintrag (§8).

**Die übrigen neun** in Kürze: drei widersprüchliche Pin-Zahlen (46/14, 16,
gemessen 17/51 — §3.2 trägt jetzt die Messung samt Aufruf) · §3.5 rechnete die
Delta-Dateien falsch auf und übersah, dass `templates/.d-check.yml` slice-188s
erstes Paar ist (oben behoben) · der Rest ist LOW/INFO, darunter der Hinweis,
dass der Kurs-Checkout inzwischen sauber ist und Risiko 4 damit retrospektiv
nicht mehr an seiner Prämisse hängt — der **Ausgang** hängt nicht daran, weil er
über die Gegenprobe belegt ist, nicht über den Zustand des fremden Baums.

## 4. Definition of Done

- [x] Das Bundle (`regelwerk/` + `templates/`) liegt vendored unter dem Tag
      `v6.6.0` mit `SHA256SUMS`; `v6.5.0` ist entfernt; `make regelwerk-check`
      grün.
- [x] **Jede** Nennung des alten Standes außerhalb der eingefrorenen Artefakte
      ist behandelt — 17 Dateien, 51 Nennungen, davon 15/48 gebumpt und 2/3 auf
      Kennung umgestellt (§3.2) — und die vier Baseline-Symlinks zeigen auf
      `v6.6.0`; `make symlink-check` und `make doc-check` grün.
- [x] Für **jeden** aktiven `MR`-Eintrag mit Baseline-Zeiger ist **gemessen**,
      ob der referenzierte Abschnitt im neuen Stand wortgleich ist — nicht
      angenommen. Wo nicht, trägt die Stelle die Abweichung sichtbar.
- [x] Unabhängiger Review durchgeführt (Report unter [`docs/reviews/`](../../../reviews/README.md)) — 12 Findings, alle abgearbeitet (§3.6).
- [x] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [x] Beobachtungs-Register fortgeschrieben.
- [x] Jedes Risiko aus §7 trägt einen Ausgang.

`make gates` und `make verify` grün. Ein öffentlicher Vertrag ist berührt:
`harness/conventions.md` §Baseline deklariert den adoptierten Stand.

## 5. Trigger

**Start** (`open` → `in-progress`): das WIP-Limit ist frei.

**Rückführungen — vorab benannt:**

- `in-progress` → `next` (zu groß): Zeigt die Zeiger-Messung, dass mehr als zwei
  `MR`-Einträge eine Abweichung tragen, wird der Adaptions-Durchgang abgetrennt.
- `in-progress` → `open` (blockiert): Lässt sich das Bundle nicht netzlos
  erzeugen, ist die Beschaffung ein eigener Vorgang — dieses Repo baut nicht
  gegen das Netz.

## 6. Closure-Trigger

DoD vollständig, `make gates` und `make verify` grün, Closure-Notiz mit
Lerneintrag.

## 7. Risiken und offene Punkte

- **Zwischen diesem Slice und [slice-193](../open/slice-193-gate-index-steht-einmal.md)
  widerspricht [`AGENTS.md`](../../../../AGENTS.md) §4 dem adoptierten Stand.**
  Der Gate-Index steht dann an zwei Orten, obwohl die Baseline einen verlangt.
  — **Ausgang:** *eingetreten* → Folge-Slice
  [slice-193](../open/slice-193-gate-index-steht-einmal.md), der genau diesen
  Umbau trägt. Der Widerspruch ist **benannt und befristet**, nicht still: Er
  steht hier, in §1 als Abgrenzung und in slice-193 §1 als Ziel. Eine gestaffelte
  Migration erzeugt ihn zwangsläufig — die Alternative wäre ein Diff, den keine
  Review-Sitzung prüfen kann.
- **Der Pin-Bump ist mechanisch und trifft 51 Stellen.** Ein übersehener Pin
  zeigt ins Leere, sobald der alte Stand weg ist; `make doc-check` fängt tote
  Links, aber ein Pin **im Fließtext** ohne Link bleibt stehen.
  — **Ausgang:** *weiter offen* → **Beobachtungs-Register**,
  [`BEO-HARNESS/massen-ersetzung-trifft-die-historische-aussage`](../observations/BEO-HARNESS/massen-ersetzung-trifft-die-historische-aussage/observation.md)
  (neu, 1×). Der **Einzelfall** ist behoben (§3.6), die **Klasse** nicht: Jeder
  künftige Versions-Wechsel trifft dieselbe Falle, solange mit der nackten
  Kennung statt mit dem Pfad-Muster ersetzt wird. Das Risiko traf **die andere Richtung** als erwartet: Nicht ein
  übersehener Pin blieb stehen, sondern eine Massen-Ersetzung hob **historische**
  Nennungen mit, die stehenbleiben mussten — im Plan selbst und im Closure-Log
  der Roadmap. Kein Lauf fängt das: `versions` bindet an Pfad-Pins, eine nackte
  Kennung im Fließtext sieht es nicht. Gefunden hat es der unabhängige Review.
  **Nach der Behebung gemessen:** Außerhalb der eingefrorenen Artefakte nennt den
  alten Stand nur noch, was ihn nennen **muss** — die sieben historischen Stellen
  im Plan und die zwei Kennung-Zeiger in `conventions/done/`. `make doc-check`,
  `make symlink-check` und `make regelwerk-check` sind grün, und der alte Baum ist
  entfernt: Ein übersehener **Link** hätte kein Ziel mehr und wäre rot.
- **Die MR-Zeiger dürfen nur wandern, wenn der Zielabschnitt wortgleich ist.**
  — **Ausgang:** *entfallen*, gestrichen mit Begründung: Für alle fünf gemessen
  (§3.4), alle fünf wortgleich. Die Messung lief **zweistufig** — Datei im Delta,
  dann Abschnitt —, und die zweite Stufe war nötig:
  [`MR-014`](../../../../harness/conventions.md#mr-014) zeigt in eine
  geänderte Datei, aber in einen unveränderten Abschnitt. Eine Messung auf
  Datei-Ebene hätte hier eine Abweichung behauptet, die es nicht gibt.
- **Das Bundle entsteht aus einem fremden Arbeitsbaum.** Der lokale
  Kurs-Checkout trägt eine uncommittete Änderung; ein Build daraus wäre nicht
  der Tag. — **Ausgang:** *entfallen*, gestrichen mit Begründung: Gebaut wurde
  aus `git archive <tag>` in ein Wegwerf-Verzeichnis, der Arbeitsbaum also nie
  gelesen **und nie angefasst**. Belegt ist das nicht durch die Absicht, sondern
  durch die Gegenprobe (§3.1): Derselbe Weg auf den alten Tag reproduziert den
  bisher vendorten Baum byte-gleich — was aus einem verschmutzten Arbeitsbaum
  nicht folgen könnte.

## 8. Closure-Notiz

**Lerneintrag — Form: geschärfte Regel.** *Zwei Regeln, die einzeln richtig
sind, können sich gegenseitig unmöglich machen — und man merkt es erst an dem
Lauf, der beide zugleich befolgt.* [`harness/conventions.md`](../../../../harness/conventions.md)
§Baseline verlangt **genau einen** vendorten Stand und sagt zugleich, der Zeiger
einer aufgelösten Adaption **bleibe** auf ihrem damaligen Stand. Beides zusammen
heißt: ein garantiert toter Link, jedes Mal. Zwei Einträge trugen ihn als Link
und wären beim Entfernen des alten Standes rot geworden.

**Warum der Abschnitt trotzdem behauptete, das sei bezahlt:** Er beruft sich auf
die **null Nachzüge** des vorigen Sprungs. Diese Messung galt den vier
einfrierenden Ziel-Formen — Review-Report, Ergebnisnotiz, zwei Archiv-Stubs.
`conventions/done/` war nie darin, und der Satz sagte es nicht. **Die Regel, die
daraus wird:** Eine Messung, die eine Kollision für erledigt erklärt, nennt die
Datei-Klasse, über die sie gemessen hat — sonst deckt sie beim nächsten Mal
etwas, das sie nie angesehen hat. Das ist die Mess-Regel *Geltungsbereich* auf
ihre eigene Erfolgsmeldung angewandt.

**Zwei beobachtbare Closure-Kriterien.** (1) `make gates`, `make verify`,
`make regelwerk-check` und `make symlink-check` grün auf dem Stand, der nach
`done/` geht — der alte Baum ist entfernt, ein übersehener Zeiger hätte kein
Ziel. (2) **Die Gegenprobe des Bau-Verfahrens:** Derselbe Weg, auf den alten Tag
angewandt, reproduziert den bisher vendorten Baum byte-gleich inklusive Manifest.
Das ist der Beleg, dass der neue Baum auf demselben Weg entstanden ist wie der
alte — nachrechenbar, nicht behauptet.

**Was der Sprung inhaltlich bringt, ist genau ein Punkt** — *„Der Gate-Index
steht einmal"* —, und er ist **nicht** in diesem Slice. Er streicht eine
41-zeilige Tabelle aus [`AGENTS.md`](../../../../AGENTS.md) §4 und zieht drei
Konfigurationen um; das ist ein Umbau, kein Pin-Bump. Adresse:
[slice-193](../open/slice-193-gate-index-steht-einmal.md). Der Widerspruch
dazwischen ist benannt und befristet (§7).

**Und eine Ironie, die zur Sache gehört:**
[slice-187](../done/wellenlos/slice-187-voll-abgleich-erstdurchgang-rest.md) hat
soeben den Satz *„Diese Tabelle listet auf; die Bindung steht in
`harness/README.md` §Sensors"* nach §4 übernommen — korrekt gegen den damals
adoptierten Stand, in dem er seit `v5.12.0` stand. Der neue Stand löscht ihn
mitsamt der Tabelle. **Das ist kein Fehler des Abgleichs**, sondern der Normalfall
eines Repos, das einer Baseline folgt: Der Voll-Abgleich misst gegen den
**adoptierten** Stand, und der ist per Definition der von gestern, sobald ein
neuer erscheint. Wer daraus schließt, man solle mit dem Abgleich warten, hat die
Reihenfolge falsch herum: slice-187s Übernahme hat den Bestand für **diesen**
Sprung erst lesbar gemacht.

**Der zweite Lerneintrag kommt vom Review, und er ist unbequem: *Eine Regel zu
zitieren ist nicht, sie anzuwenden.*** §9 dieses Slice trägt die Lehre aus
slice-187 im Wortlaut — *„das Register nach **allen** berührten Kürzeln lesen"* —
und zog beim Ziehen der Konsequenz **dieselbe** Grenze wie sein Vorgänger: nach
der Sub-Area-Deklaration statt nach dem, was der Diff anfasst. Die Regel stand
zwölf Zeilen über der Stelle, an der sie verletzt wurde. **Was daraus folgt:**
Eine Lehre, die als Satz in den nächsten Plan wandert, ist noch kein Handgriff.
Der Handgriff wäre `git diff --name-only` gegen die Kürzel-Tabelle — mechanisch,
nicht als Erinnerung.

**Und der erste Lerneintrag hat eine Schwester, die ihn erklärt:** Beide Fehler
dieses Slice — die mitgehobenen historischen Nennungen und die zu enge Sichtung —
sind Fälle, in denen **das richtige Werkzeug danebenlag**. Für die Ersetzung war
es das Muster des Versions-Sensors, für die Sichtung `git diff --name-only`.
Gegriffen wurde beide Male zum Naheliegenden: der nackten Kennung und der
Deklaration.

**Beobachtungs-Register.** Erhöht:
[`BEO-GATE/versions-sensor-trifft-planungs-vorgriff`](../observations/BEO-GATE/versions-sensor-trifft-planungs-vorgriff/observation.md)
auf **3×** — der Migrationsplan selbst nannte den kommenden Stand als Pfad, und
`versions` meldete ihn zu Recht. Der Ausgang stand seit zwei Belegen im
`state.md` und ist jetzt gezogen: **verkörpert**, Planungs-Dokumente übernehmen
die Zitier-Form ([`AGENTS.md`](../../../../AGENTS.md) §5, `seit slice-192`).
Neu bei 1×:
[`BEO-HARNESS/zwei-regeln-machen-einander-unmoeglich`](../observations/BEO-HARNESS/zwei-regeln-machen-einander-unmoeglich/observation.md)
(§3.3) und
[`BEO-HARNESS/massen-ersetzung-trifft-die-historische-aussage`](../observations/BEO-HARNESS/massen-ersetzung-trifft-die-historische-aussage/observation.md)
(§3.6). **Nicht** erhöht wurden die vier `PLAN`-Einträge bei 2×: Sie sind in
diesem Vorgang **bedient oder behoben**, und ein Fund, den der Review findet und
derselbe Slice schließt, ist kein zweites Auftreten in einem anderen Vorgang
(§9).

## 9. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen.** Zwei Sub-Areas, beide in
[`harness/conventions.md`](../../../../harness/conventions.md)
§Modus-Deklaration geführt:

- **`HARNESS`** (`AGENTS.md`, `harness/`) — Achsen 1,2,3. Die Pins, die
  Stand-Zeile, die MR-Zeiger.
- **Vendored Baseline** (`.harness/baseline/`) — Achsen 1,3, **kein Modus**:
  externer, unveränderter Fremdtext. Die Deklaration nennt für sie ausdrücklich
  *„Aktualisierung nur als Migrations-Slice"* — dieser Slice ist genau der Fall,
  für den die Zeile geschrieben wurde.

**Modus:** `HARNESS` ist Greenfield, die Baseline trägt keinen — ein
Begründungsblock entfällt für beide.

**Randberührung, benannt statt übergangen:** `GATE` ist berührt, weil vier
Symlinks unter `.claude/rules/` umhängen und `make symlink-check` sie prüft. Es
ist keine Konventions-Frage, sondern eine Ziel-Anpassung; deshalb keine eigene
Zeile in der Prüfung, aber genannt.

**Vorgelagert — offene Beobachtungen sichten.** **Zwei Quellen**, und das
Register nach **allen** berührten Kürzeln — die Lehre aus
[slice-187](../done/wellenlos/slice-187-voll-abgleich-erstdurchgang-rest.md) §9,
wo die Sichtung nur `BEO-PLAN` las und zwei `BEO-HARNESS`-Einträge übersah, die
mit jenem Slice 3× erreichten.

**(1) Register — sechs Einträge betreffen diesen Slice:**

| Eintrag | Kürzel | Berührung |
|---|---|---|
| [`versions-sensor-trifft-planungs-vorgriff`](../observations/BEO-GATE/versions-sensor-trifft-planungs-vorgriff/observation.md) | `GATE` | **3× erreicht** — der Plan selbst löste ihn aus; Ausgang *verkörpert* (§8) |
| [`zwei-baseline-staende-nach-migrationsende`](../observations/BEO-HARNESS/zwei-baseline-staende-nach-migrationsende/observation.md) | `HARNESS` | **bedient, nicht erhöht** — der alte Stand geht mit diesem Slice; genau das, wogegen der Eintrag steht |
| [`rueckbau-kandidat-ueberlebt-baseline-migration`](../observations/BEO-HARNESS/rueckbau-kandidat-ueberlebt-baseline-migration/observation.md) | `HARNESS` | **geprüft, kein Treffer** — kein aktiver `MR` wird durch diesen Sprung entbehrlich; die fünf Zeiger wandern alle |
| [`adaption-korrigiert-repo-aussage`](../observations/BEO-HARNESS/adaption-korrigiert-repo-aussage/observation.md) | `HARNESS` | **nicht erhöht** — dieser Slice legt keinen `MR` an; sein Ausgang ist bereits *geplant* (slice-190) |
| [`baseline-regel-nie-erwogen-weil-bestand-sie-verletzt`](../observations/BEO-HARNESS/baseline-regel-nie-erwogen-weil-bestand-sie-verletzt/observation.md) | `HARNESS` | **berührt, nicht erhöht** — die neue Gate-Index-Regel *wird* erwogen, und zwar in slice-193 statt still übergangen |
| [`symlink-ziel-nach-baseline-bump-ungeprueft`](../observations/BEO-GATE/symlink-ziel-nach-baseline-bump-ungeprueft/observation.md) | `GATE` | **bedient** — `make symlink-check` ist genau dafür da und lief grün |

**Neu angelegt** (§3.3):
[`BEO-HARNESS/zwei-regeln-machen-einander-unmoeglich`](../observations/BEO-HARNESS/zwei-regeln-machen-einander-unmoeglich/observation.md)
bei 1×.

**(2) Report des Vorgängers** — zu slice-187, mit ihm archiviert. Zwei seiner
Finding-Klassen greifen hier unmittelbar und sind in die Arbeit eingebaut:
*Messung ohne reproduzierbares Instrument* → §3.1 nennt Bau-Weg und Gegenprobe
statt nur das Ergebnis; *Vollständigkeits-Haken ohne erschöpften Gegenstand* →
die MR-Zeiger sind **zweistufig** gemessen (Datei, dann Abschnitt), weil die
Datei-Ebene für [`MR-014`](../../../../harness/conventions.md#mr-014) das falsche Ergebnis geliefert hätte.

**`PLAN` und `REVIEW` sind berührt — korrigiert nach dem Review (F-3).** Der
Diff arbeitet in `docs/plan/planning/` (Plan, Roadmap, Register) und in
`docs/reviews/` (der Report). Die erste Fassung dieses Abschnitts erklärte beide
für unberührt, weil sie nach der **Sub-Area-Deklaration** zuschnitt statt nach
dem, was der Diff anfasst — dieselbe Grenze, die slice-187 gezogen hatte und
deren Lehre zwölf Zeilen weiter oben zitiert steht.

| Eintrag | Kürzel | Berührung |
|---|---|---|
| [`zielsatz-nach-plan-aenderung-nicht-nachgezogen`](../observations/BEO-PLAN/zielsatz-nach-plan-aenderung-nicht-nachgezogen/observation.md) | `PLAN` | **berührt, nicht erhöht** — der Umfang dieses Slice hat sich nicht geändert; Titel und §1 tragen dieselbe Zusage wie die DoD |
| [`messung-ohne-reproduzierbares-instrument`](../observations/BEO-PLAN/messung-ohne-reproduzierbares-instrument/observation.md) | `PLAN` | **bedient** — §3.1 und §3.2 nennen Bau-Weg, Gegenprobe und den Mess-Aufruf, nicht nur das Ergebnis |
| [`vollstaendigkeits-haken-ohne-erschoepften-gegenstand`](../observations/BEO-PLAN/vollstaendigkeits-haken-ohne-erschoepften-gegenstand/observation.md) | `PLAN` | **eingetreten und behoben** — die DoD-Zeile zu den Pins nannte eine Zahl, die den Gegenstand nicht deckte; jetzt steht die Messung mit ihrem Aufruf daneben |
| [`kandidaten-klassifikation-groeber-als-der-kandidat`](../observations/BEO-PLAN/kandidaten-klassifikation-groeber-als-der-kandidat/observation.md) | `PLAN` | **eingetreten und behoben** — *„10 geänderte Dateien"* deckte die zehnte nicht, die sich nur im Tag unterscheidet (§3.5) |
| [`review-geltungsbereich-zu-eng`](../observations/BEO-PLAN/review-geltungsbereich-zu-eng/observation.md) | `PLAN` | **berührt** — genau die Klasse, die §8 als zweiten Lerneintrag trägt |

**Keiner der fünf erreicht mit diesem Slice 3×**: Die vier `PLAN`-Einträge bei 2×
sind hier **bedient oder behoben**, nicht neu belegt — ein Fund, den der Review
findet und der Slice im selben Vorgang schließt, ist kein zweites Auftreten der
Klasse in einem *anderen* Vorgang. Zu `REVIEW` steht nichts Offenes an; zu
`SPEC`, `KERN` und `ADAPT` ebenfalls nichts.

**Neu angelegt, zusätzlich zu §3.3:**
[`BEO-HARNESS/massen-ersetzung-trifft-die-historische-aussage`](../observations/BEO-HARNESS/massen-ersetzung-trifft-die-historische-aussage/observation.md)
bei 1× (§3.6, F-1/F-2).
