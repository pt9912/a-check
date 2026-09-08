# slice-182 — `harness/README.md` verweist, statt Regelwerk-Begründung nachzuschreiben

**Welle:** ohne Welle — die Closure-Bedingung wäre die DoD dieses Slice, und
`in-progress/` trägt keinen Welle-Bezug (Baseline-Regelwerk
`modul-06-roadmap.md` §Wann Arbeit eine Welle braucht).

**Bezug:** [`AC-QA-02`](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze) (Harness-Integrität), keine aktive ADR.

**Berührte Spec-Stellen:** — (der Slice berührt nur Rang 9 der Source Precedence).

**Verantwortlich:** Claude — gesetzt beim Übergang nach `in-progress/`.

**Autor:** Claude. **Datum:** 2026-09-07.

**Lerneintrag — Form:** geschärfte Regel.

---

## 1. Ziel und Abgrenzung

**Ziel:** [`harness/README.md`](../../../../harness/README.md) trägt für jede
Regel, die im vendorten Regelwerk steht, einen **Verweis** statt einer
nachgeschriebenen Begründung — die Datei ist Schritt 1 des Minimal Agent
Workflow und wird von jedem Lauf ganz gelesen.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Die Tabellen selbst.** Sie wachsen mit der Zahl der Gates, nicht mit der
  Geschwätzigkeit; ihre Zellen wächtert
  [slice-181](../open/slice-181-tabellenzellen-gewaechtert.md) maschinell. Hier geht es
  um die Prosa **um** sie herum, die kein `table.column` erreicht.
- **[`AGENTS.md`](../../../../AGENTS.md).** Dieselbe Klasse, andere Datei — der
  Register-Eintrag nennt sie ausdrücklich (*„`AGENTS.md` §1 trägt dieselbe
  Klausel weiterhin im Rumpf"*). Ein Vorgang, der beide anfasst, berührt zwei
  Rang-Stufen der Source Precedence und ist in einer Review-Sitzung nicht mehr
  prüfbar. **Folge-Slice**, wenn dieser hier die Kriterien belegt hat.
- **Ein Sensor.** *„Ist dieser Satz im Regelwerk schon gesagt?"* ist ein Urteil,
  kein Match — dieselbe Grenze, die [`AGENTS.md`](../../../../AGENTS.md) §3.7 für
  sich selbst benennt. Was maschinell geht, ist die Zeichen-Obergrenze in
  Tabellenzellen, und die trägt [slice-181](../open/slice-181-tabellenzellen-gewaechtert.md).

**Zwei benannte Plan-Änderungen** ([`AGENTS.md`](../../../../AGENTS.md) §6: „Nimmt
der Lauf etwas mit, das §1 ausschließt, ist das eine **Plan-Änderung**"):

1. **`AGENTS.md` wird doch angefasst** — aber nur um den Zielort des
   Lerneintrags. Der Ausschluss oben zielt auf `AGENTS.md` als *Gegenstand* des
   Aufräumens; eine geschärfte Regel braucht dort ihren Herkunfts-Anker, sonst
   ist der Lerneintrag *gezählt, nicht verkörpert*. Der Gegenstands-Ausschluss
   bleibt: kein Absatz in `AGENTS.md` wurde auf Duplikate geprüft oder gekürzt.
2. **§Source precedence kommt hinzu.** §3 listete drei Abschnitte; der Review
   fand in einem vierten eine Chronik-Zeile derselben Klasse
   (*„der zuvor ausgelassene Rang ist mit [`MR-003`](../../../../harness/conventions/done/MR-003-source-precedence-ohne-docs-user.md) (aufgelöst) eingefügt"*).
   Eine Zeile in derselben Datei, dieselbe Regel — mitgenommen und hier
   benannt, statt sie für einen Folge-Slice liegen zu lassen.
3. **Zwei Tatsachen-Korrekturen, die die Kürzung sichtbar machte.**
   [`harness/README.md`](../../../../harness/README.md) §Rollen nannte in zwei
   Tabellenzeilen [`MR-009`](../../../../harness/conventions/done/MR-009-validator-unbesetzt.md)
   — aufgelöst durch [`MR-016`](../../../../harness/conventions/MR-016-validator-unbesetzt.md), auf das der neue Absatz daneben zeigt; und
   [`AGENTS.md`](../../../../AGENTS.md) §4 nannte `make gate-consistency` als
   Erzwinger der Doku-↔-Makefile-Konsistenz, obwohl `doc-targets` das seit
   slice-079 tut. Der zweite Widerspruch ist **älter** als dieser Slice; der
   gestrichene Chronik-Halbsatz war die einzige Prosa-Stelle, die beide
   Fassungen zusammenhielt. Beide korrigiert, statt sie als Kollateralschaden
   der Kürzung stehen zu lassen.

## 2. Ausgangsmessung (2026-09-07)

**Sichtbare Prosa** je Abschnitt, gemessen gegen
`v6.5.0` · `templates/harness/README.template.md`.

**Methode:** Abschnitt = von einer `##`-Überschrift bis zur nächsten; daraus
HTML-Kommentare (`<!-- … -->`) und alle mit `|` beginnenden Zeilen entfernt;
gezählt werden die verbleibenden Zeichen einschließlich Zeilenumbrüchen.
**Geltungsbereich:** **fünf** der **neun** `##`-Abschnitte — die vier übrigen
(§Purpose, §Guides, §Source precedence, §Minimal agent workflow) liegen bei
±241 Zeichen gegen die Ziel-Form und wurden deshalb nicht einzeln vermessen.
Das ist eine Auswahl nach Größe, und sie hat einen Fund verfehlt: In
§Source precedence stand eine Chronik-Zeile, die erst der Review fand
(§1, Plan-Änderung 2). Eine Größen-Auswahl sieht Duplikate nur, wenn sie groß
sind.

| Abschnitt | a-check | Ziel-Form | Faktor |
|---|---|---|---|
| §Sensors | 1930 | 346 | **5,6 ×** |
| §Rollen und ihre Übergabe-Artefakte | 1542 | **0** | Abschnitt existiert dort nicht |
| §Safety and scope boundaries | 1055 | 47 | 22 × |
| §Leseordnung | 843 | 464 | 1,8 × |
| §Traceability rules | 753 | 385 | 2,0 × |

Gesamt **21 426** gegen **10 436** Zeichen — Stand bei Slice-Beginn; die 21 374 der ersten Fassung waren der Stand zwei Commits davor (Review F-3). Der Faktor allein ist **kein**
Befund — ein ausgefülltes Dokument ist länger als seine Vorlage. Befund ist,
**wo derselbe Text schon woanders steht**; drei Stellen sind belegt:

1. **Doppeltes Zitat in derselben Datei.** `### Nicht-Gates` (500 Zeichen) und
   der Absatz *„Nicht hier, obwohl sie in keinem Aggregat hängen"* (553)
   paraphrasieren beide `modul-13` §Vorhanden ≠ behauptet — der Satz *„ein
   reales Target nicht als Gate zu führen ist keine Harness-Lüge"* steht
   zweimal, 22 Zeilen auseinander.
2. **Die Ziel-Form verbietet genau diese Stelle.** Ihr Bedienhinweis: `kein
   Gate` gehört *„IN DER ZEILE SELBST … **nicht in Prosa daneben**"*, und
   *„WÄCHST DIE SEKTION: … Ob der Überhang schon unter der Tabelle steht oder
   in die Zelle gedrängt wurde, **ist dieselbe Sache**."* An dieselbe Stelle
   setzt sie **68 Zeichen**: eine fette Zeile, keine Überschrift, kein Absatz.
   [slice-177](../done/welle-15/slice-177-sensors-struktur-zwei-tabellen.md) hat
   den Überhang aus den Zellen geholt und ihn darüber wieder aufgebaut.
3. **Dreifacher Satz im Repo.** *„Verifikation grün, Validation rot"* steht in
   `modul-08`:142, in [`MR-016`](../../../../harness/conventions/MR-016-validator-unbesetzt.md):19
   und in `harness/README.md`:189.

Im Rollen-Abschnitt trägt die **Tabelle** (11 Zeilen) die a-check-eigene
Zuordnung *Übergabe → Artefakt* — sie steht nirgends sonst und bleibt. Die 25
Prosa-Zeilen darum sind `modul-08` §Die neun Übergaben, §Rollen-Regeln und
[`MR-016`](../../../../harness/conventions/MR-016-validator-unbesetzt.md); der Schluss-Absatz über die Review-Serie vom 2026-07-26 ist zusätzlich
Chronik in einer gelesenen Datei ([`AGENTS.md`](../../../../AGENTS.md) §3.7).

**Beide Dateien verbieten das selbst:** [`AGENTS.md`](../../../../AGENTS.md) §1
(*„sie dupliziert deren Inhalt nicht; sonst entsteht Drift"*),
[`harness/README.md`](../../../../harness/README.md) §Purpose (*„Diese Datei
dupliziert sie nicht"*).

## 3. Plan (vor Code)

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| `harness/README.md` §Sensors, Nicht-Gates-Einleitung | refactor | auf die Ziel-Form-Zeile; das zweite Zitat entfällt |
| `harness/README.md` §Rollen und ihre Übergabe-Artefakte | refactor | Tabelle bleibt, Prosa auf zwei Zeiger, Chronik-Absatz raus |
| `harness/README.md` §Safety · §Leseordnung · §Traceability | update | je Satz geprüft: im Regelwerk ⇒ Verweis, a-check-eigen ⇒ bleibt |

**Kriterium je Satz** — nicht die Länge, sondern die Herkunft: Steht die
Aussage im vendorten Regelwerk, wird sie ein Verweis. Ist sie a-check-eigen
(eine Zuordnung, eine Schwelle, eine benannte Lücke), bleibt sie.

## 4. Definition of Done

- [ ] §Sensors trägt die Nicht-Gates-Kennzeichnung in der Form der Ziel-Form;
      das doppelte `modul-13`-Zitat ist auf **einen** Verweis reduziert.
- [ ] §Rollen trägt die Zuordnungs-Tabelle unverändert und die Prosa als
      Zeiger auf `modul-08` und [`MR-016`](../../../../harness/conventions/MR-016-validator-unbesetzt.md);
      der Chronik-Absatz ist fort.
- [ ] Die übrigen Abschnitte sind **je Satz** gegen das Regelwerk geprüft, und
      jede Streichung wie jedes Bleiben ist im Review-Report begründet.
- [ ] Unabhängiger Review durchgeführt (Report unter [`docs/reviews/`](../../../reviews/README.md)).
- [ ] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [ ] Beobachtungs-Register fortgeschrieben.
- [ ] Jedes Risiko aus §7 trägt einen Ausgang.

`make gates` und `make verify` grün. Die drei Paarungen (Anker · Folge-Slice ·
Register) trägt die Closure dieses Slice selbst — wellenloser Betrieb. Ein
öffentlicher Vertrag ist nicht berührt: `harness/README.md` ist Rang 9.

**Nebenbefund zur Zähl-Kalibrierung:** `tasks-ignore-pattern` in
[`.d-check.yml`](../../../../.d-check.yml) kennt *„Die drei Paarungen"* und
*„Doku-Update"* nicht, obwohl `modul-05` §Ziel-Form: Slice beide unter den
**konstanten** Posten führt. Beide stehen hier deshalb als feste Zeile statt als
Häkchen — dieselbe Auflösung, die [`AGENTS.md`](../../../../AGENTS.md) §5 für
den Gate-Lauf vorschreibt. Ob das Muster nachzuziehen ist, entscheidet
[slice-169](../done/wellenlos/slice-169-korpus-seitige-kalibrierung.md) oder ein
Folge-Slice, nicht dieser.

## 5. Trigger

**Start** (`next` → `in-progress`): [slice-169](../done/wellenlos/slice-169-korpus-seitige-kalibrierung.md)
liegt in `done/` (WIP-Limit 1).

**Rückführungen — vorab benannt:**

- `in-progress` → `next` (zu groß): Zeigt die Satz-für-Satz-Prüfung, dass mehr
  als die drei geplanten Abschnitte betroffen sind, wird nach Abschnitten
  zerlegt statt in einem Lauf durchgezogen.
- `in-progress` → `open` (blockiert): Stellt sich beim Kürzen heraus, dass eine
  Aussage **nur** hier steht und im Regelwerk fehlt, ist das eine benannte
  Spec-Lücke — sie wird eingetragen, und der Slice wartet auf ihre Auflösung.

## 6. Closure-Trigger

DoD vollständig, `make gates` und `make verify` grün, Closure-Notiz mit
Lerneintrag geschrieben — und die **gemessene** Gegenprobe: die sichtbare Prosa
je geänderten Abschnitts erneut gezählt, gegen §2 gehalten.

## 7. Risiken und offene Punkte

- **Kürzen entfernt eine Aussage, die nur hier steht** — der Kontext-Trennungs-Absatz
  über die Review-Serie 2026-07-26 trägt eine Selbst-Einschätzung, kein Regelzitat.
  — **Ausgang:** *entfallen*, gestrichen mit Begründung — die Aussage geht
  nicht verloren: Der Absatz sagte
  selbst, *„die Reports weisen sich darum ausdrücklich als Selbst-Review aus"*.
  Sie steht also in den Reports, die sie beschreibt, und die frieren ein. Was
  hier stand, war ihre Wiederholung an einer Stelle, die jeder Lauf liest.
- **Der Verweis altert anders als der Text.** Ein Zeiger auf `modul-08` §X bricht,
  wenn die nächste Baseline den Abschnitt umbenennt; ein nachgeschriebener Satz
  nicht. — **Ausgang:** *weiter offen* → Beobachtungs-Register,
  [`BEO-PLAN/ziel-form-tag-gescopt`](../observations/BEO-PLAN/ziel-form-tag-gescopt/observation.md)
  trägt dieselbe Mechanik für die Ziel-Formen. **Gemessen, nicht angenommen:**
  Die neun in diesem Slice gesetzten Verweise nennen den Abschnitt in
  **Inline-Code** (`` `modul-13` §Vorhanden ≠ behauptet ``), nicht als
  Markdown-Link — `doc-check` sieht sie deshalb nicht, und der `versions`-Sensor
  auch nicht. Der Bruch fiele beim Lesen auf, nicht beim Lauf; das ist die
  benannte Grenze, kein stiller Ausfall.
- **Kein Sensor hält das Ergebnis.** Nach dem Kürzen wächst die Prosa beim
  nächsten Slice wieder — genau so ist sie entstanden. — **Ausgang:** *weiter
  offen* → Beobachtungs-Register,
  [`BEO-HARNESS/baseline-normtext-nachgeschrieben`](../observations/BEO-HARNESS/baseline-normtext-nachgeschrieben/observation.md)
  (mit diesem Slice 2×). Ein Sensor dafür wäre ein Urteil über die Herkunft
  eines Satzes und damit keiner ([`AGENTS.md`](../../../../AGENTS.md) §3.7);
  was maschinell geht, ist die Zeichen-Obergrenze in Tabellenzellen, und die
  trägt [slice-181](../open/slice-181-tabellenzellen-gewaechtert.md).

## 8. Closure-Notiz

**Lerneintrag — Form: geschärfte Regel.**

- **Was hat funktioniert:** Die Ausgangsmessung nach **sichtbarer** Prosa zu
  trennen — HTML-Bedienhinweise der Vorlage und Tabellenzeilen herausgerechnet.
  Ohne diese Trennung hätte §Sensors mit +8001 Zeichen als größter Befund
  dagestanden; tatsächlich hat a-check dort **weniger** Prosa als die Vorlage
  Bedienhinweis, und die Tabelle wächst legitim mit der Zahl der Gates.

- **Was ging anders als geplant — drei der fünf Abschnitte blieben unangetastet.**
  §3 plante, §Safety, §Leseordnung und §Traceability „je Satz zu prüfen und zu
  entscheiden". Die Prüfung entschied gegen das Kürzen: §Traceability ist die
  **ausgefüllte** Ziel-Form (vier Punkte, gleiche Struktur, konkrete Targets),
  §Safety ebenso — dort führt die Vorlage zwei Platzhalter-Punkte und den
  Bedienhinweis *„repo-spezifisch formulieren"* —, und §Leseordnung **zitiert**
  die Regel bereits mit Verweis, statt sie nachzuschreiben; die Form, die dieser
  Slice herstellen wollte, stand dort schon. Der Faktor gegen die Ziel-Form war
  bei §Safety **22×** und trotzdem kein Befund.

  **Korrigiert nach dem Review (F-1):** Die erste Fassung begründete §Safety mit
  *„fünf AC-gebundene Zusagen, die nirgends sonst stehen"*. Das ist falsch —
  dieselben Zusagen stehen in den beiden `AC-QA-*`, auf die sie verlinken, in
  [`README.md`](../../../../README.md):90,
  [`spec/architecture.md`](../../../../spec/architecture.md):106 und wörtlich im
  [Benutzerhandbuch](../../../../docs/user/benutzerhandbuch.md):744. Der
  Abschnitt bleibt trotzdem stehen, aber aus einem anderen Grund: Eine
  `AC-*`-Zusammenfassung mit Link zeigt **nach oben** in der Source Precedence,
  und das ist die Aufgabe eines Einstiegspunkts. Gegenstand dieses Slice ist
  nachgeschriebener **Baseline**-Normtext, nicht die eigene Spec — die
  Unterscheidung fehlte der ersten Fassung, und sie steht jetzt in der
  geschärften Regel.

- **Steering-Loop-Eintrag — geschärfte Regel:** Der Größen-*Vergleich* mit einer
  Ziel-Form ist kein Befund; Befund ist, **wo derselbe Text schon woanders
  steht**. Eine Vorlage ist kürzer als jedes ausgefüllte Dokument, und ein
  Bedienhinweis, der beim Kopieren gelöscht wird, zählt in ihr mit. Wer nach
  Zeichen misst, kürzt am Ende die ausgefüllten Stellen und lässt die
  duplizierten stehen. — liegt in
  [`AGENTS.md`](../../../../AGENTS.md) §5, *Geltungsbereich einer Messung*
  (`seit slice-182` dort ergänzt).
  Auslöser: [`BEO-HARNESS/baseline-normtext-nachgeschrieben`](../observations/BEO-HARNESS/baseline-normtext-nachgeschrieben/observation.md)
  (slice-103, slice-182 — 2×).

- **Beobachtungs-Register (`../observations/`):** zwei Belege ergänzt —
  [`baseline-normtext-nachgeschrieben`](../observations/BEO-HARNESS/baseline-normtext-nachgeschrieben/observation.md)
  (2×, `evidence/slice-182.md`) und
  [`chronik-in-gelesenen-dateien`](../observations/BEO-HARNESS/chronik-in-gelesenen-dateien/observation.md)
  (2×, `evidence/slice-182.md`). Beide bleiben unter der Schwelle.

- **Was der Review fand — und warum es zu diesem Slice gehört.** Der
  unabhängige Lauf war **merge-blockierend**: 3 HIGH, 4 MEDIUM, 6 LOW, 1 INFO.
  Sein Verdikt trennt sauber: *„Die Kürzung selbst trägt … Blockierend ist die
  Beleg-Schicht darum herum."* Alle drei HIGH betrafen **Behauptungen über den
  Bestand**, nicht die Arbeit am Bestand:
  eine Alleinstellungs-Aussage, die gegen [`spec/lastenheft.md`](../../../../spec/lastenheft.md)
  falsch war und als Beleg nach [`AGENTS.md`](../../../../AGENTS.md) §5 gewandert
  ist (F-1); ein `state.md`, das nach dem Commit einen Zustand behauptete, den
  derselbe Commit beseitigt hatte (F-2); ein Vorher-Wert, der zwei Commits alt
  war (F-3). Der Reviewer benennt die Ironie: *„Befund ist, wo derselbe Text
  schon woanders steht — das setzt voraus, dass man weiß, wo er sonst noch
  steht."*

- **Folge-Slice:** [slice-183](../open/slice-183-agents-md-verweist-statt-wiederholt.md)
  — dieselbe Frage für [`AGENTS.md`](../../../../AGENTS.md), die §1 hier
  ausgeschlossen hat. Die Kriterien tragen jetzt: Der Slice hat gezeigt, dass
  der Größen-Vergleich **nicht** entscheidet (§Safety blieb bei Faktor 22×
  stehen) und die Herkunft eines Satzes schon (§Sensors trug dasselbe Zitat
  zweimal). Erste Messung an `AGENTS.md`: **34 690** gegen **10 885** Zeichen
  der Ziel-Form, §5 allein **15 464** gegen **726** — und dort ist die Ziel-Form
  **ausgeschrieben**, der Faktor also kein Platzhalter-Artefakt.

- **Risiken aus §7:** drei, jedes mit genau einem Ausgang — einmal *gestrichen
  mit Begründung*, zweimal *weiter offen* → Register.

- **Drei Paarungen** (Repo ohne Wellen-Betrieb) — geprüft nach dem `git mv`
  nach `done/`; eingetragen im dritten Closure-Commit.

### Messung nach der Umsetzung

| Abschnitt | vorher | jetzt | Ziel-Form | Entscheidung |
|---|---|---|---|---|
| §Sensors | 1930 | **1278** | 346 | gekürzt: doppeltes `modul-13`-Zitat auf einen Verweis, Chronik-Halbsatz raus |
| §Rollen (Prosa) | 1542 | **438** | *(kennt sie nicht)* | gekürzt: Tabelle bleibt, Prosa auf zwei Zeiger, Chronik-Absatz raus, wiederholter Trigger raus |
| §Source precedence | 212 | **81** | 385 | gekürzt (Plan-Änderung 2): Chronik-Zeile zu [`MR-003`](../../../../harness/conventions/done/MR-003-source-precedence-ohne-docs-user.md) |
| §Traceability rules | 753 | 753 | 385 | **unverändert** — ausgefüllte Ziel-Form |
| §Safety and scope | 1055 | 1055 | 47 | **unverändert** — ausgefüllte Ziel-Form; die Vorlage führt dort zwei Platzhalter-Punkte |
| §Leseordnung | 843 | 843 | 464 | **unverändert** — zitiert die Regel bereits mit Verweis |

Datei gesamt **21 426 → 19 409** Zeichen. Die Ziel-Form liegt bei 10 436; der
Abstand ist überwiegend die Gate-Tabelle mit 31 Zeilen, und die wächst mit der
Zahl der Gates, nicht mit der Geschwätzigkeit.

## 9. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** berührt ist allein **Harness-Einstieg**
(`HARNESS`; Achsen 1,2,3 laut
[`harness/conventions.md`](../../../../harness/conventions.md#modus-deklaration-pro-sub-area)).
Die Gate-/Werkzeug-Schicht ist **nicht** berührt: kein Target, kein Skript, kein
Workflow ändert sich.

**Vorgelagert — offene Beobachtungen sichten:** zwei Treffer in `HARNESS`.

- [`baseline-normtext-nachgeschrieben`](../observations/BEO-HARNESS/baseline-normtext-nachgeschrieben/observation.md)
  — bisher **1×** (slice-103). Dieser Slice ist das zweite Vorkommen und hebt
  ihn auf **2×**: Symptom, noch keine Lücke.
- [`chronik-in-gelesenen-dateien`](../observations/BEO-HARNESS/chronik-in-gelesenen-dateien/observation.md)
  — bisher **1×**. Der Absatz über die Review-Serie 2026-07-26 ist das zweite
  Vorkommen, ebenfalls **2×**.

Erreicht einer der beiden mit einem *weiteren* Slice 3×, ist er eine Lücke und
braucht einen eigenen Folge-Slice — für eine inferentielle Regel heißt das einen
Skill oder eine Hard Rule, keinen Sensor ([`AGENTS.md`](../../../../AGENTS.md) §3.7).

**Modus-Begründungsblock:** alle berührten Sub-Areas GF.
