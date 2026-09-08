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
Baseline-Verzeichnis unter seinem Tag, `v6.6.0` ist entfernt, alle **46**
Pin-Vorkommen in
den **14** lebenden Dateien zeigen auf den neuen Stand, und für jeden aktiven
`MR`-Eintrag ist **gemessen**, ob sein `Ersetzt-Baseline-Regel`-Zeiger mitwandern
darf.

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

## 2. Delta-Analyse `v6.6.0` → `v6.6.0` (2026-09-08)

**Gemessen** gegen den lokalen Kurs-Checkout, `git diff v6.6.0 v6.6.0 --
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
`harness/README.md` §Sensors"* nach §4 übernommen — aus der `v6.6.0`-Ziel-Form,
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

| Gegenstand | Zahl | Behandlung |
|---|---|---|
| Lebende Dateien mit Pin | **16** | gebumpt |
| Baseline-Symlinks unter `.claude/rules/` | **4** | umgehängt; `make symlink-check` grün |
| Stand-Zeile in [`harness/conventions.md`](../../../../harness/conventions.md) §Baseline | 1 | Kurs-Welle **129 · 2026-09-08**, aus dem Kopf des vendorten `regelwerk/README.md` übernommen |
| Eingefrorene Artefakte, die den alten Stand **nennen** | viele | **unangetastet** — sie beschreiben, was damals galt; `versions` nimmt sie über `exempt-paths` aus |

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

### 3.5 Die Delta-Zahl präzisiert

§2 nennt **10** geänderte Dateien nach `git diff`. Normalisiert (ohne die
Tag-Zeile) sind es **neun echte** plus eine, die sich **nur** im gepinnten
Release-Asset-Link unterscheidet: `templates/harness/conventions.template.md`.
Für slice-188/189 ist das relevant — sie messen gegen die Vorlagen, und eine
Vorlage, die sich nur im Tag unterscheidet, ist kein Abgleich-Gegenstand.

## 4. Definition of Done

- [x] Das Bundle (`regelwerk/` + `templates/`) liegt vendored unter dem Tag
      `v6.6.0` mit `SHA256SUMS`; `v6.6.0` ist entfernt; `make regelwerk-check`
      grün.
- [x] Alle **46** Pin-Vorkommen in den **14** lebenden Dateien und die vier
      Baseline-Symlinks unter `.claude/rules/` zeigen auf `v6.6.0`;
      `make symlink-check` und `make doc-check` grün.
- [x] Für **jeden** aktiven `MR`-Eintrag mit Baseline-Zeiger ist **gemessen**,
      ob der referenzierte Abschnitt im neuen Stand wortgleich ist — nicht
      angenommen. Wo nicht, trägt die Stelle die Abweichung sichtbar.
- [ ] Unabhängiger Review durchgeführt (Report unter [`docs/reviews/`](../../../reviews/README.md)).
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
- **Der Pin-Bump ist mechanisch und trifft 46 Stellen.** Ein übersehener Pin
  zeigt ins Leere, sobald der alte Stand weg ist; `make doc-check` fängt tote
  Links, aber ein Pin **im Fließtext** ohne Link bleibt stehen.
  — **Ausgang:** *entfallen*, gestrichen mit Begründung: Beide Richtungen sind
  gemessen. Vorwärts: `grep` auf den alten Stand liefert außerhalb der
  eingefrorenen Artefakte **null** Treffer. Rückwärts: `make doc-check`,
  `make symlink-check` und `make regelwerk-check` sind grün, und der alte Baum
  ist entfernt — ein übersehener Link hätte kein Ziel mehr und wäre rot. Der
  Fließtext-Fall ist damit ebenfalls abgedeckt, weil der `grep` nicht nach Links
  sucht, sondern nach der Zeichenfolge.
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

**Beobachtungs-Register.** Erhöht:
[`BEO-GATE/versions-sensor-trifft-planungs-vorgriff`](../observations/BEO-GATE/versions-sensor-trifft-planungs-vorgriff/observation.md)
auf **3×** — der Migrationsplan selbst nannte den kommenden Stand als Pfad, und
`versions` meldete ihn zu Recht. Der Ausgang stand seit zwei Belegen im
`state.md` und ist jetzt gezogen: **verkörpert**, Planungs-Dokumente übernehmen
die Zitier-Form ([`AGENTS.md`](../../../../AGENTS.md) §5, `seit slice-192`).
Neu bei 1×:
[`BEO-HARNESS/zwei-regeln-machen-einander-unmoeglich`](../observations/BEO-HARNESS/zwei-regeln-machen-einander-unmoeglich/observation.md)
— der Fund aus §3.3.

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

**Keine Treffer sind ebenfalls eine Antwort:** Zu `SPEC`, `PLAN`, `KERN`,
`ADAPT` und `REVIEW` steht nichts Offenes an, das dieser Slice berührt.
