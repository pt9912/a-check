# slice-119 — Guide für CR-Texte an ein fremdes Werkzeug

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese Datei liegt — eines von
`open/`, `next/`, `in-progress/`, `done/`. Er wechselt nur durch `make slice-mv`
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).

**Bezug:** [`BEO-022`](../observations.md) — bei **3×**, Schwelle überschritten.

**Berührte Spec-Stellen:** — *(keine)*

**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.

**Autor:** Claude (Opus 5), im Auftrag des Maintainers. **Datum:** 2026-08-30.

---

## 1. Ziel

Ein CR an ein fremdes Werkzeug bekommt einen Prüf-Durchgang, bevor er hinausgeht — dieselbe
Stelle, an der die letzten drei ihre Fehler hatten.

## 2. Definition of Done

- [ ] Ein Skill `.harness/skills/cr-text-reviewer.md` prüft einen CR-Text gegen die **zwei
      belegten Ausprägungen**: eine messbare Behauptung gar nicht gemessen, oder die **eigene**
      Menge gemessen und über die **fremde** ausgesagt. Er trägt die vier belegten Fälle als
      Fixtures — drei eigene, einer vom Empfänger.
- [ ] Der Skill nennt die **Prüf-Frage je Behauptungs-Klasse** (eigener Bestand · fremder Vertrag ·
      eigener Gate-Pfad) mit dem Handgriff, der sie beantwortet — nicht „sorgfältig prüfen".
- [ ] Das Workflow-Skelett verweist an der Stelle darauf, an der CR-Texte entstehen; der
      Skill-Index in [`AGENTS.md`](../../../../AGENTS.md) §5 nennt ihn neben dem
      Closure-Note-Reviewer.

- [ ] `make gates` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.
- [ ] Closure-Notiz mit benanntem Lerneintrag geschrieben (§7).
- [ ] Beobachtungs-Register fortgeschrieben.
- [ ] Jedes Risiko aus §6 trägt genau einen Ausgang.

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
([`SL-002`](../observations.md), zwei Vorfälle nach dem Guide). Für CR-Texte gibt es noch keinen —
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

- *Ein Guide, den niemand aufruft, ist wirkungslos — genau der Befund aus slice-118* —
  **Ausgang:** <bei Closure>
- *Die Fixtures sind vier Fälle aus zwei Repos; eine fünfte Klasse könnte danebenliegen* —
  **Ausgang:** <bei Closure>

## 7. Closure-Notiz

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice,
[`AGENTS.md`](../../../../AGENTS.md) §5.)_

**Lerneintrag — Form: <geschärfte Regel | neuer Sensor | benannte Spec-Lücke>**

## 8. Sub-Area-Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** berührt wird der **Harness-Einstieg** (Skill, AGENTS §5,
Workflow-Skelett). Kein Code, kein Gate.

**Vorgelagert — offene Beobachtungen sichten:** [`BEO-022`](../observations.md) ist der Anlass
(3×). [`BEO-009`](../observations.md) (Chronik in Dateien, die jeder Lauf liest) betrifft dieselbe
Schicht und bleibt offen — dieser Slice schreibt einen Skill, den nur der CR-Fall lädt.

Alle berührten Sub-Areas GF.
