# slice-119 — Guide für CR-Texte an ein fremdes Werkzeug

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese Datei liegt — eines von
`open/`, `next/`, `in-progress/`, `done/`. Er wechselt nur durch `make slice-mv`
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).

**Bezug:** [`BEO-022`](../../../../docs/plan/planning/observations/BEO-GATE/cr-text-behauptet-statt-gemessen/observation.md) — bei **3×**, Schwelle überschritten.

**Berührte Spec-Stellen:** — *(keine)*

**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.

**Autor:** Claude (Opus 5), im Auftrag des Maintainers. **Datum:** 2026-08-30.

---

## 1. Ziel

Ein CR an ein fremdes Werkzeug bekommt einen Prüf-Durchgang, bevor er hinausgeht — dieselbe
Stelle, an der die letzten drei ihre Fehler hatten.

## 2. Definition of Done

- [x] Ein Skill `.harness/skills/cr-text-reviewer.md` prüft einen CR-Text gegen die **zwei
      belegten Ausprägungen**: eine messbare Behauptung gar nicht gemessen, oder die **eigene**
      Menge gemessen und über die **fremde** ausgesagt. Er trägt die vier belegten Fälle als
      Fixtures — drei eigene, einer vom Empfänger.
- [x] Der Skill nennt die **Prüf-Frage je Behauptungs-Klasse** (eigener Bestand · fremder Vertrag ·
      eigener Gate-Pfad) mit dem Handgriff, der sie beantwortet — nicht „sorgfältig prüfen".
- [x] Das Workflow-Skelett verweist an der Stelle darauf, an der CR-Texte entstehen; der
      Skill-Index in [`AGENTS.md`](../../../../AGENTS.md) §5 nennt ihn neben dem
      Closure-Note-Reviewer.

- [x] `make gates` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.
- [x] Closure-Notiz mit benanntem Lerneintrag geschrieben (§7).
- [x] Beobachtungs-Register fortgeschrieben.
- [x] Jedes Risiko aus §6 trägt genau einen Ausgang.

## 3. Plan (vor Code)

| Datei | Änderungs-Art | Begründung |
|---|---|---|
| `.harness/skills/cr-text-reviewer.md` | neu | der Guide; Form nach `closure-note-reviewer.md` |
| [`AGENTS.md`](../../../../AGENTS.md) §5 | update | Skill nennen, wo die Closure-Pflicht ihren nennt |
| `.claude/commands/slice.md` | update | Verweis, wo CR-Texte entstehen |

**Auszuführende Gates:** `make gates` (tragend `doc-check` — der Skill verlinkt Kennungen) und
`make verify`.

### Warum ein Guide und nicht ein Werkzeug

[slice-118](../done/slice-118-lifecycle-wechsel-werkzeug.md) hat gerade das Gegenteil entschieden,
und die Begründung trägt hier **nicht**: dort war ein Guide **schon einmal gescheitert**
([`SL-002`](../../../../docs/plan/planning/observations/README.md), zwei Vorfälle nach dem Guide). Für CR-Texte gibt es noch keinen —
ein Guide ist die richtige **erste** Antwort, nicht die zweite.

Ein Sensor ist zudem nicht baubar: *„ist diese Behauptung gemessen?"* ist ein Urteil über den
Entstehungsweg eines Satzes, kein Match. Dieselbe Grenze, die
[`AGENTS.md`](../../../../AGENTS.md) §3.7 für die Kommentar-Regel ausdrücklich benennt.

### Die zweite Ausprägung ist die gefährlichere

`BEO-022` zählt bislang drei Fälle, in denen **gar nicht** gemessen wurde. Der Empfänger hat in
derselben Runde die verwandte Form geliefert und sie bei sich als `BEO-020` verbucht: *„gemessen
wird die eigene Menge, ausgesagt wird über die fremde"*. Seine Antwort trug eine **Tabelle mit
Messwerten**, deren vierte Zeile ungemessen war — und sie traf im Ergebnis sogar zu.

Das ist die schärfere Klasse, denn sie **sieht aus wie ein Beleg**. Der Skill muss sie darum als
eigene Frage führen, nicht als Unterfall: nicht *„hast du gemessen?"*, sondern *„hast du **das**
gemessen, worüber du redest?"*.

## 4. Trigger

**Start:** eingetreten — `BEO-022` bei 3×.

**Rückführungen:**

- `in-progress` → `next`: falls sich zeigt, dass die Prüf-Fragen nur den bekannten drei Fällen
  nachgebildet sind und keine vierte Behauptungs-Klasse tragen. Dann ist es eine Fallsammlung,
  kein Guide.

## 5. Closure-Trigger

Skill steht mit Fixtures und Prüf-Fragen, beide Verweise gesetzt, Gates grün.

**Was bewusst nicht getan wird:** die vier bestehenden CR-Texte **nachträglich prüfen**. Ihre
Fehler sind bekannt, benannt und beantwortet; ein Nachlauf träfe die Form statt der Substanz.
Ebenso wenig entsteht ein CR-Verzeichnis analog zum Empfänger — a-check führt CR-Texte im Slice,
der sie erzeugt, und das hat bei vier Stück getragen.

## 6. Risiken und offene Punkte

- *Ein Guide, den niemand aufruft, ist wirkungslos — genau der Befund aus
  [slice-118](../done/slice-118-lifecycle-wechsel-werkzeug.md)* — **Ausgang:** weiter offen im
  **Beobachtungs-Register**. Der Verweis steht an **zwei** Stellen im Pfad
  ([`AGENTS.md`](../../../../AGENTS.md) §5 und Workflow-Skelett bei „Spec-first"), aber das ist
  eine Verbesserung der Chance, kein Beleg für Wirkung. Der nächste CR misst es.
- *Die Fixtures sind vier Fälle aus zwei Repos; eine fünfte Klasse könnte danebenliegen* —
  **Ausgang:** gestrichen mit Begründung: der Skill führt die Klassen ausdrücklich als
  **„bisher belegte"** und fordert das Nachtragen einer weiteren samt Zähler-Erhöhung. Eine
  geschlossene Liste zu behaupten, wäre der Fehler gewesen — die offene ist die Antwort.

## 7. Closure-Notiz

**Geliefert:** [`.harness/skills/cr-text-reviewer.md`](../../../../.harness/skills/cr-text-reviewer.md)
— der Prüf-Durchgang für einen CR-Text, bevor er den Slice verlässt. Er markiert jeden Satz, der
eine **Tatsache** über ein System behauptet, und nennt je Klasse den Handgriff, der sie belegt:
`grep -c` über den eigenen Bestand, den Abschnitt im fremden Lastenheft, `grep` über die
Gate-Pfade. Vier Fixtures aus zwei Repos, davon drei eigene Fehler.

**Lerneintrag — Form: geschärfte Regel.** *Ein Beleg wird nicht daran geprüft, **ob** gemessen
wurde, sondern **ob die gemessene Menge die ist**, über die der Satz redet.* Die drei eigenen
Fälle waren offene Annahmen und damit leicht zu sehen, sobald man hinschaute. Der vierte — vom
Empfänger, und von ihm selbst gemeldet — ist die schärfere Form: eine **Tabelle mit vier
Messzeilen**, deren vierte ungemessen war. Er hatte seine eigenen Gate-Dateien gemessen und über
unser Fragment ausgesagt. *Weil* die Aussage im Ergebnis sogar zutraf, wäre sie ohne die Prüfung
nie aufgefallen — und wäre trotzdem kein Beleg gewesen. Der Empfänger hat das selbst so eingeordnet
(*„dass unsere unbelegte Behauptung im Ergebnis zutraf, macht sie nicht zu einem Beleg"*) und
seinen Zähler stehen lassen. Genau darum ist die Prüf-Frage nicht „hast du gemessen?" — diese
Frage hätte der vierte Fall mit **ja** beantwortet.

**Zwei beobachtbare Closure-Kriterien:**

1. Der Skill trägt die vier Fälle als **benannte Fixtures** mit Klasse und Ausprägung, nicht als
   Prosa — wer ihn prüfen will, hat vier Eingaben mit erwartetem Ergebnis.
2. Die Klassen-Liste ist ausdrücklich **offen** („die drei bisher belegten") und fordert das
   Nachtragen einer weiteren samt Zähler-Erhöhung. Eine geschlossene Liste hätte denselben Fehler
   wiederholt, den der Skill fangen soll: eine Vollständigkeit behaupten, die nicht gemessen ist.

**Offene Risiken und ihr Ausgang:** der erste weiter offen im Register, der zweite gestrichen mit
Begründung.

**Beobachtungs-Register:** [`BEO-022`](../../../../docs/plan/planning/observations/BEO-GATE/cr-text-behauptet-statt-gemessen/observation.md) ist **verkörpert** im Skill; der Zähler
bleibt bei 3×, sein Stand nennt jetzt den Ort. Die zweite Ausprägung („die falsche Menge
gemessen") ist in der Beschreibung ergänzt — sie ist bei uns **nicht** aufgetreten und erhöht
darum nichts; der Empfänger führt sie bei sich als eigene Klasse mit eigenem Zähler.

**Folge-Slices:** [slice-117](../done/slice-117-handbuch-verweis-cli.md) (Handbuch-Verweis) bleibt
der einzige offene; beide Schwellen dieser Sitzung sind damit beantwortet — `BEO-008` mit einem
Werkzeug, `BEO-022` mit einem Guide.
## 8. Sub-Area-Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** berührt wird der **Harness-Einstieg** (Skill, AGENTS §5,
Workflow-Skelett). Kein Code, kein Gate.

**Vorgelagert — offene Beobachtungen sichten:** [`BEO-022`](../../../../docs/plan/planning/observations/BEO-GATE/cr-text-behauptet-statt-gemessen/observation.md) ist der Anlass
(3×). [`BEO-009`](../../../../docs/plan/planning/observations/BEO-HARNESS/chronik-in-gelesenen-dateien/observation.md) (Chronik in Dateien, die jeder Lauf liest) betrifft dieselbe
Schicht und bleibt offen — dieser Slice schreibt einen Skill, den nur der CR-Fall lädt.

Alle berührten Sub-Areas GF.
