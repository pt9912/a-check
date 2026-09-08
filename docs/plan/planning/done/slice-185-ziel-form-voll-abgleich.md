# slice-185 — Voll-Abgleich Ziel-Form ↔ Gegenstück: was vier Delta-Analysen nicht finden konnten

**Welle:** ohne Welle — die Closure-Bedingung wäre die eigene DoD
(Baseline-Regelwerk `modul-06-roadmap.md` §Wann Arbeit eine Welle braucht).

**Bezug:** [`AC-QA-02`](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)
(Harness-Integrität). Keine aktive ADR wird berührt.

**Berührte Spec-Stellen:** — (Harness-Artefakte ohne Vertragsberührung).

**Verantwortlich:** Claude — gesetzt beim Übergang nach `in-progress/`.

**Autor:** Claude. **Datum:** 2026-09-08.

**Lerneintrag — Form:** geschärfte Regel.

---

## 1. Ziel und Abgrenzung

**Ziel:** Jede vendored Ziel-Form mit **genau einem** Gegenstück im Repo ist
einmal vollständig gegen dieses Gegenstück abgeglichen — und der Abgleich ist
als wiederkehrender Schritt beim Baseline-Sprung verankert.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Die zwölf Instanz-Vorlagen** (Slice, ADR, Carveout, Review-Report,
  Beobachtung, Welle, Archiv-Stubs, …). Sie haben kein *einzelnes* Gegenstück,
  sondern viele; ihre Form prüft `make doc-structure` bereits über Muster. Ein
  Abgleich „Vorlage gegen alle Instanzen" wäre ein anderer Vorgang mit einem
  anderen Werkzeug.
- **Die Ziel-Formen ohne Gegenstück** (`reconciliation.template.md`,
  `project-readme.template.md`, `gate.template.md`, …). a-check führt sie
  bewusst nicht — `reconciliation.md` etwa gibt es nicht, weil es keinen
  Brownfield-Bootstrap gab ([`AGENTS.md`](../../../../AGENTS.md) §5,
  Kopieranleitung Punkt 2). Bestand, der bewusst stehen bleibt.
- **Ein Sensor auf Satz-Deckung.** *„Trägt das Gegenstück diese Aussage?"* ist
  ein Urteil über zwei Formulierungen, kein Match — dieselbe Grenze wie
  [`AGENTS.md`](../../../../AGENTS.md) §3.7. Ein Grob-Vergleich per Wortfolge
  liefert überwiegend Rauschen; das ist in §2.3 gemessen, nicht vermutet.
- **Produkt-Code.** Der Slice rührt `internal/` nicht an.

## 2. Ausgangsmessung (2026-09-08)

### 2.1 Der Gegenstand

| | |
|---|---|
| Dateien unter `.harness/baseline/v6.5.0/templates/` | **28** |
| davon **Ziel-Formen** (ohne `templates/README.md`, den Verzeichnis-Index) | **27** |
| davon mit **genau einem** Gegenstück im Repo | **15** |
| Normtext (ohne Bedienhinweise und Platzhalter), gemessen über 14 davon | **66 741** Zeichen |

**Korrigiert nach dem Review (F-7):** Die erste Fassung nannte 26 und 14. Die
26 waren `find -name '*.md'` — sie zählte `Makefile` und `.d-check.yml` **nicht**
und den Verzeichnis-Index **mit**. Und `roadmap.template.md` hat sehr wohl ein
Gegenstück: [`docs/plan/planning/in-progress/roadmap.md`](../in-progress/roadmap.md).
Die Paarbildung suchte es unter `docs/plan/planning/roadmap.md` und fand nichts —
**ein Pfad-Match ist keine Zuordnung.** Die Normtext-Zahl bezieht sich weiter auf
die 14 damals gebildeten Paare; sie ist damit eine **Untergrenze**.

**Ein Gegenstück ist kürzer als seine Vorlage** — bei einem ausgefüllten
Dokument der schärfste Einstiegspunkt:

| Paar | Ziel-Form | Gegenstück |
|---|---|---|
| `.harness/skills/reviewer.md` | 5294 | **3835** |

**Korrigiert nach dem Review (F-1):** Die erste Fassung nannte hier ein zweites
Paar, `README.md` mit 13 433 gegen 9230. Die 13 433 gehören zu
`v6.5.0` · `templates/README.md` — dem **Index des Vorlagen-Verzeichnisses**,
der keine Ziel-Form ist und sich selbst als Übersicht beschreibt. Die Ziel-Form
eines Projekt-`README.md` ist `project-readme.template.md` mit **2076** Zeichen;
a-checks `README.md` (9274) ist ihr gegenüber **länger**.
**Ursache:** Die automatische Paarbildung mappte jeden Vorlagen-Pfad per
`.template`-Streichung auf ein Gegenstück — bei `templates/README.md` ergab das
`README.md`, und zwei Dateien mit demselben Basisnamen wurden zu einem Paar.
Ein Namens-Match ist keine Zuordnung; die richtige steht in der Übersichts-
Tabelle von `templates/README.md` selbst (*„`project-readme.template.md` →
Projekt-Root-`README.md`"*).

### 2.2 Der erste Treffer trägt sofort

`.harness/skills/reviewer.md` fehlen **drei HIGH-Kategorien**, die die Ziel-Form
führt:

| HIGH-Kategorie | Vorkommen in a-checks Skill |
|---|---|
| **Norm nur im Template-Kommentar** — eine Regel steht im `<!-- -->`-Block und nirgends sonst | **0** |
| **Kommentar trägt keine der Kommentar-Klassen** — verworfene Alternative, abwesender Text, abgebrochener Satz | **0** |
| **Zustandsfeld trägt Chronik** — eine `Stand`/`Status`-Zelle erzählt die Entstehung | **0** |

Jede trägt in der Ziel-Form denselben Nachsatz: *„Kein Gate fängt das."* Sie
stehen dort, **weil** kein Sensor sie sieht.

**Das ist bereits teuer geworden.** Zwei Register-Einträge zählen genau diese
Klassen —
[`chronik-in-gelesenen-dateien`](../observations/BEO-HARNESS/chronik-in-gelesenen-dateien/observation.md)
(2×) und
[`hard-rule-37-ohne-sensor`](../observations/BEO-HARNESS/hard-rule-37-ohne-sensor/observation.md)
(2×). Gefunden hat sie jedes Mal ein Review; **gesucht hat keines danach**,
weil der Skill die Kategorien nicht führt. Sie fielen als Nebenprodukt an.

### 2.3 Warum es vier Migrationen nicht fanden

Jede Baseline-Migration ist eine **Delta**-Analyse: slice-135
(`v5.12.0`→`v6.0.0`), slice-161, slice-164, slice-174 (`v6.2.0`→`v6.5.0`). Ein
Delta findet, was sich **ändert**.

**Gemessen an einem Beispielsatz** — dem Arbeitsteilungs-Satz für `AGENTS.md` §4
(*„Diese Tabelle listet auf; definiert wird hier nichts"*): Er steht **seit
`v5.12.0`** unverändert in der Ziel-Form. In keinem der vier Deltas taucht er
auf, und a-check hat ihn nie übernommen.

**Ein Voll-Abgleich Vorlage-gegen-Gegenstück hat es nie gegeben.** Das ist die
Lücke — nicht ein Versehen einer einzelnen Migration.

**Grenze des maschinellen Vergleichs, gemessen:** Ein Grob-Abgleich per
Wortfolge über `AGENTS.md` meldet **30** Sätze ohne Entsprechung, überwiegend
Rauschen — Template-Hinweise, Platzhalter und Bedienkommentare, die beim
Kopieren bestimmungsgemäß verschwinden. Der Abgleich ist damit **Lese-Arbeit
mit maschineller Vorauswahl**, kein Lauf.

## 3. Umsetzung

**Fünf der 15 Paare abgeglichen** — je Paar mit einem der drei Ausgänge:

| Paar | Kandidaten | Ausgang |
|---|---|---|
| `.harness/skills/reviewer.md` | — | **übernommen:** drei HIGH-Kategorien (§2.2) |
| `docs/plan/adr/README.md` | 2 | **übernommen:** der *Derivativ*-Hinweis |
| `docs/plan/carveouts/README.md` | 2 | **übernommen:** derselbe Hinweis |
| `spec/architecture.md` | 4 | **teils bewusst abweichend, teils offen** — zwei von drei Klauseln eingehalten, die dritte an slice-186 |
| `Makefile` | 6 | **ohne Befund** — alle sechs sind Bootstrap-Bedienhinweise der Vorlage |

**Der `Derivativ`-Hinweis hatte null Treffer im ganzen Repo.** Er sagt, was ein
Index *ist*: Quelle der Wahrheit sind die Dateien, der Index ist eine
Bequemlichkeits-Sicht und wird mitgezogen. **`make gate-consistency` prüft genau
das** (ADR-Index-Vollständigkeit, slice-087) — die Begründung des Sensors stand
im Sensor, nicht im Index. Beide Index-Dateien tragen sie jetzt, mit dem
Verweis auf die zwei Richtungen (`gate-consistency` für „Datei ohne Zeile",
`doc-check` für „Zeile ohne Datei").

**`spec/architecture.md` — teils bewusst abweichend, teils offener Punkt.** Die
Ziel-Form setzt dort eine Hard Rule **in die Datei**, und sie hat **drei**
Klauseln. Der Review (F-2) hat gezeigt, dass die erste Fassung dieses Absatzes
nur die zweite geprüft hat:

| Klausel der Ziel-Form | a-check |
|---|---|
| keine Wellen, Slices, Commit-Hashes, Closure-Daten | **eingehalten** — je 0 Treffer, gemessen |
| keine ADR-Bezüge | **eingehalten** — 0 Treffer |
| **keine Historie** — *„`Letzte Änderung` oben ist ein Frische-Marker, kein Protokoll"* | **verletzt** — §8 trägt eine Versions-Tabelle mit vier Einträgen |

**Die ersten zwei Klauseln:** Die Regel steht in
[`AGENTS.md`](../../../../AGENTS.md) §3.4 und wird eingehalten; sie zusätzlich
in die Datei zu schreiben wäre die Doppelung, die
[slice-183](../done/wellenlos/slice-183-agents-md-verweist-statt-wiederholt.md)
gerade aufgelöst hat. **Bewusst abweichend**, und dieser Fall bleibt der Beleg
dafür, dass ein Abgleich, der jede Differenz als Lücke liest, stillen Rückbau
erzeugt (§7, Risiko 2).

**Die dritte ist ein offener Punkt und geht an
[slice-186](../open/slice-186-voll-abgleich-restliche-paare.md).** Sie ist
nicht durch §3.4 gedeckt — dort steht nichts über Historie-Abschnitte — und
die Entscheidung ist keine dieses Slice: `lastenheft.template.md` **sieht** eine
Historie vor (drei Nennungen), `architecture.template.md` **verbietet** sie
(eine Nennung, im Verbot). Ob a-check den Abschnitt streicht und auf `git`
verweist oder die Abweichung als `MR` deklariert, ist eine Entscheidung über
ein Spec-Stratum.

**Was der Review daran zeigt:** Eine Teil-Messung als Deckungs-Nachweis für eine
**mehrteilige** Zusage auszugeben, ist derselbe Fehler wie eine Probe, die ihren
Gegenstand verfehlt — nur auf der Lese-Seite. Die Zusage hatte drei Klauseln,
geprüft war eine.

**Verankert** ist der Abgleich in
[`harness/conventions.md`](../../../../harness/conventions.md) §Baseline: Wer
den Stand hebt, fährt **Delta und Voll-Abgleich** — mit der Messung, die zeigt,
warum das Delta allein nicht reicht.

## 4. Definition of Done

- [ ] Die drei HIGH-Kategorien stehen in
      [`.harness/skills/reviewer.md`](../../../../.harness/skills/reviewer.md),
      in a-checks Sprache und mit den Fundstellen, die sie im Bestand hatten.
- [x] **Fünf** Paare sind abgeglichen; je Paar steht das Ergebnis — übernommen,
      bewusst abweichend (mit Begründung) oder ohne Befund.
      **Plan-Änderung, benannt statt stillschweigend:** Die DoD verlangte *alle*
      13. Gemessen sind es **169** Kandidaten über 13 Paare — Risiko 3 aus §7 ist
      damit eingetreten, und der Ausgang ist ein Folge-Slice mit Kennung
      ([slice-186](../open/slice-186-voll-abgleich-restliche-paare.md)), nicht
      ein gedehnter Slice.
- [ ] Der Abgleich ist als **wiederkehrender Schritt** verankert: Wer die
      Baseline hebt, fährt ihn — nicht nur das Delta.
- [ ] Unabhängiger Review durchgeführt (Report unter [`docs/reviews/`](../../../reviews/README.md)).
- [ ] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [ ] Beobachtungs-Register fortgeschrieben.
- [ ] Jedes Risiko aus §7 trägt einen Ausgang.

`make gates` und `make verify` grün. Ein öffentlicher Vertrag ist berührt:
`AGENTS.md` ist Rang 8, der Reviewer-Skill steuert jeden Review.

## 5. Trigger

**Start** (`open` → `in-progress`): WIP-Limit frei.

**Rückführungen — vorab benannt:**

- `in-progress` → `next` (zu groß): Zeigen die ersten Paare, dass der Abgleich
  je Datei einen eigenen Durchgang braucht, wird nach Dateien zerlegt — der
  Reviewer-Skill zuerst, weil sein Befund schon steht.
- `in-progress` → `open` (blockiert): Findet der Abgleich eine Ziel-Form-Regel,
  die a-check bewusst **nicht** übernehmen will, ist das eine
  `MR`-Adaption und kein Nachtrag; der Slice wartet auf diese Entscheidung.

## 6. Closure-Trigger

DoD vollständig, `make gates` und `make verify` grün, Closure-Notiz mit
Lerneintrag. Danach Archivierung als wellenloser Slice
([`AGENTS.md`](../../../../AGENTS.md) §6).

## 7. Risiken und offene Punkte

- **Der Abgleich ist Lese-Arbeit und damit unvollständig prüfbar.** Was ein
  Leser übersieht, sieht auch der nächste Lauf nicht — dieselbe Grenze, die
  §3.7 für sich benennt. — **Ausgang:** *weiter offen* → Beobachtungs-Register,
  [`BEO-HARNESS/hard-rule-37-ohne-sensor`](../observations/BEO-HARNESS/hard-rule-37-ohne-sensor/observation.md).
  Der Slice verkleinert das Risiko messbar, statt es zu schließen: Die
  maschinelle Vorauswahl **hat** die drei HIGH-Kategorien und den
  Derivativ-Hinweis gefunden — was ein Leser allein wohl übersehen hätte. Sie
  ersetzt das Lesen nicht, aber sie richtet es aus.
- **Ein Nachtrag kann eine bewusste Abweichung überschreiben.** a-check hat
  Ziel-Form-Regeln begründet nicht übernommen; ein Abgleich, der jede Differenz
  als Lücke liest, macht daraus stillen Rückbau. — **Ausgang:** *eingetreten*,
  aufgefangen **im ersten Durchgang**: `spec/architecture.md` (§3). Die Ziel-Form
  will die Hard Rule in der Datei, a-check trägt sie in
  [`AGENTS.md`](../../../../AGENTS.md) §3.4 und hält sie ein. Ein blinder
  Nachtrag hätte die Doppelung erzeugt, die
  [slice-183](../done/wellenlos/slice-183-agents-md-verweist-statt-wiederholt.md)
  gerade aufgelöst hat. **Das Risiko ist real und der Ausgang ist die
  Arbeitsweise:** je Paar wird der Ausgang genannt, nicht nur die Differenz.
  Ein Carveout braucht es dafür nicht — die Disziplin steht in der DoD.
- **13 Paare in einem Slice könnten die Größen-Regel sprengen.** Die
  Rückführung ist in §5 benannt, kostet aber einen Durchgang.
  — **Ausgang:** *eingetreten* → Folge-Slice
  [slice-186](../open/slice-186-voll-abgleich-restliche-paare.md). Gemessen sind
  es **169** Kandidaten über 13 Paare; dieser Slice trägt fünf. Die Rückführung
  nach `next/` wurde **nicht** gezogen, weil der Slice seine anderen zwei
  Liefer-Punkte vollständig erbracht hat — der Rest ist ein eigener Vorgang mit
  eigenem Gegenstand, kein unfertiger Teil dieses.

## 8. Closure-Notiz

**Lerneintrag — Form: geschärfte Regel** (Voll-Abgleich neben dem Delta,
verankert in `harness/conventions.md` §Baseline).

- **Was hat funktioniert:** Die **Kürzer-als-die-Vorlage**-Heuristik. Von 14
  Paaren waren zwei kürzer als ihre Ziel-Form; das erste hat beim ersten
  Hinsehen drei fehlende HIGH-Kategorien geliefert. Bei einem *ausgefüllten*
  Dokument ist „kürzer als die Vorlage" ein starkes Signal — und es ist eine
  Zahl, kein Urteil.

- **Was der Slice gefunden hat, ist unmittelbar teuer gewesen.** Die drei
  fehlenden HIGH-Kategorien sind genau die Klassen, die zwei Register-Einträge
  bei je 2× zählen. **Gefunden hat sie jedes Mal ein Review — gesucht hat
  keines danach.** Sie fielen als Nebenprodukt an, und ob sie auffielen, hing
  am Zufall statt am Skill.

- **Was ging anders als geplant:** Der Slice trägt **fünf** der 14 Paare, nicht
  alle. 169 Kandidaten über 13 Paare — Risiko 3 ist eingetreten, und der
  Ausgang ist ein Folge-Slice mit Kennung statt eines gedehnten Slice. Die
  Plan-Änderung steht in der DoD, nicht im Bericht danach.

- **Der wichtigste Einzelbefund ist ein Nicht-Befund:** `spec/architecture.md`
  weicht von seiner Ziel-Form ab und soll es. Die Hard Rule steht in
  [`AGENTS.md`](../../../../AGENTS.md) §3.4 und wird eingehalten; sie in die
  Datei zu kopieren wäre die Doppelung, die slice-183 gerade aufgelöst hat.
  **Ein Abgleich, der jede Differenz als Lücke liest, erzeugt stillen Rückbau** —
  deshalb verlangt die DoD je Paar einen *Ausgang*, nicht eine Differenzliste.

- **Was der Review fand — vier HIGH, und alle vier an derselben Stelle.** Sein
  Verdikt trennt sauber: *„Die Idee des Slice trägt … die Ausführung der Messung
  trägt an vier Stellen nicht — und dieser Slice macht Messgenauigkeit zu seinem
  Gegenstand."* Der **Kern-Beleg hält** (unabhängig nachgemessen: der
  Arbeitsteilungs-Satz steht wortgleich in `v5.12.0`, `v6.0.0`, `v6.2.0`,
  `v6.5.0` und fehlt in `v3.5.2`). Falsch waren die Messungen darum herum:

  | Finding | was falsch war |
  |---|---|
  | F-1 | `README.md` gegen den **Verzeichnis-Index** gemessen statt gegen `project-readme.template.md` — a-checks Datei ist **länger**, nicht kürzer |
  | F-2 | Nicht-Befund zu `spec/architecture.md` deckte **eine von drei** Klauseln; die dritte ist verletzt |
  | F-3 | *„a-check kopiert die Ziel-Formen nicht"* widerspricht `AGENTS.md` §5 — und der Fall ist besetzt |
  | F-4 | Feld `klasse` fehlte im Output-Schema, obwohl alle neun Reports es führen |

- **Zwei Fehler, eine Ursache: ein Match ist keine Zuordnung.** F-1 kam von der
  Paarbildung per **Basisnamen** (`templates/README.md` → `README.md`), F-7 vom
  Suchen per **Pfad** (`roadmap.template.md` → `docs/plan/planning/roadmap.md`,
  während die Datei in `in-progress/` liegt). Beide Male hat der Automatismus
  eine Zuordnung *hergestellt* statt sie nachzuschlagen — und die richtige stand
  daneben: `templates/README.md` führt eine **Übersichts-Tabelle**, die jeder
  Vorlage ihr Ziel zuweist.

- **Und ein dritter derselben Familie, auf der Lese-Seite:** F-2. Eine Zusage
  mit drei Klauseln, eine geprüft, das Ergebnis als Deckung ausgegeben. Der
  Reviewer nennt es *„Teil-Messung als Deckungs-Nachweis für eine mehrteilige
  Zusage"* — die Lese-Variante der Probe, die ihren Gegenstand verfehlt
  ([`AGENTS.md`](../../../../AGENTS.md) §5, `seit slice-181`).

- **Steering-Loop-Eintrag — geschärfte Regel:** *Zu jedem Baseline-Sprung gehört
  neben dem Delta ein **Voll-Abgleich** der Ziel-Formen gegen ihr Gegenstück.*
  Ein Delta findet, was sich ändert; was seit der Adoption fehlt, findet nur der
  Voll-Abgleich. Gemessen am Arbeitsteilungs-Satz für `AGENTS.md` §4: seit
  `v5.12.0` unverändert in der Vorlage, in **keinem** der vier Deltas, nie
  übernommen. — liegt in
  [`harness/conventions.md`](../../../../harness/conventions.md) §Baseline
  (`seit slice-185` dort, mit der Messung und der Abgrenzung der zwölf
  Instanz-Vorlagen).

- **Beobachtungs-Register (`../observations/`):** kein neuer Eintrag, kein
  neuer Beleg. Die zwei einschlägigen Einträge **bedient** der Slice, statt sie
  auszulösen (§9) — er richtet die Suche nach ihren Klassen ein.

- **Folge-Slices:** [slice-186](../open/slice-186-voll-abgleich-restliche-paare.md)
  (die neun verbleibenden Paare) — er nimmt den Punkt an, den §7 Risiko 3 ihm
  übergibt, und beginnt mit `README.md`, dem zweiten Gegenstück, das kürzer ist
  als seine Vorlage.

- **Risiken aus §7:** drei, jedes mit genau einem Ausgang — einmal *weiter
  offen* → Register, zweimal *eingetreten*, einmal davon mit Folge-Slice.

- **Drei Paarungen** (Repo ohne Wellen-Betrieb) — geprüft **nach** dem `git mv`
  nach `done/`, weil sie dort suchen; eingetragen im dritten Closure-Commit.

## 9. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** eine Sub-Area berührt —
**Harness-Einstieg** `HARNESS` (`.harness/skills/`, `harness/conventions.md`,
`docs/plan/adr/README.md`, `docs/plan/carveouts/README.md`), Achsen 1,2,3.
Die Spec-Straten sind **nicht** berührt: `spec/architecture.md` wurde
**gelesen** und ausdrücklich nicht geändert — eine Lese-Berührung ist keine.

**Vorgelagert — offene Beobachtungen sichten:** gesichtet am 2026-09-08. Zwei
Treffer in `HARNESS`, beide **ohne neuen Beleg** und beide aus demselben Grund:

- [`hard-rule-37-ohne-sensor`](../observations/BEO-HARNESS/hard-rule-37-ohne-sensor/observation.md)
  (**2×**) und
  [`chronik-in-gelesenen-dateien`](../observations/BEO-HARNESS/chronik-in-gelesenen-dateien/observation.md)
  (**2×**) sind die zwei Klassen, die dieser Slice als **HIGH-Kategorien in den
  Reviewer-Skill** aufnimmt. Der Slice **bedient** sie, statt sie erneut
  auszulösen — er hat keinen neuen Verstoß produziert, sondern die Suche danach
  eingerichtet.
- [`baseline-normtext-nachgeschrieben`](../observations/BEO-HARNESS/baseline-normtext-nachgeschrieben/observation.md)
  (**3×**, verkörpert) berührt die Gegenrichtung: Dort ging es um zu viel
  Baseline-Text im Repo, hier um zu wenig. **Kein Beleg** — die Klasse ist eine
  andere, und `spec/architecture.md` zeigt, dass beide Richtungen dieselbe
  Prüf-Frage brauchen: *trägt der Zielort die Aussage schon?*

**Keine weiteren Treffer.**

**Alle berührten Sub-Areas GF** — kein Begründungsblock nötig.
