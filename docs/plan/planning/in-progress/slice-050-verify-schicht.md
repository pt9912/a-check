# slice-050 — Etappe E (2/3): Verifikations-Schicht einziehen

**Status:** in-progress — zweiter Schnitt der **Etappe E (Mechanik)** aus
[slice-048 §5](../done/slice-048-modul-delta-lesen.md).
**Deckt:** Fund **B-3** (kein `verify`-Target) und **B-4** (`closure-note-reviewer` fehlt).
**Nicht in diesem Slice:** B-16 (Workflow-Skelett) und B-20 (Freigabe-Checkliste) — sie folgen als
3/3. [Roadmap](roadmap.md).

---

## 1. Was gefehlt hat — mehr als das Target

Regelwerk Modul 11 trennt zwei Schichten: `make gates` beantwortet **Code-/Architektur**-Fragen,
`make verify` die **DoD-/Closure**-Fragen. a-check hatte die zweite Schicht nicht; die DoD-Prüfung
lief als Prosa im Slice-Dokument, also als Behauptung ohne Sensor.

Beim Bau fiel eine **Vorbedingung** auf, die slice-048 nicht gesehen hatte: die **Closure-Pflicht
war im Repo nirgends normativ verankert**. 45 Slices in `done/` praktizierten sie, aber keine Regel
forderte sie — eine stille Setzung. Ein Sensor ohne deklarierte Regel hätte nichts durchgesetzt,
sondern eine erfunden. Darum entsteht hier zuerst der Anker
([`AGENTS.md`](../../../../AGENTS.md) §5), dann das Target.

Bewusst **nicht** mitgenommen: die Pflicht, den Lerneintrag einer der **drei Baseline-Formen**
zuzuordnen (Fund B-5). Das ist Form-Arbeit und gehört nach Etappe D; §5 fordert hier vorerst
*einen* von drei Inhalten, nicht dessen deklarierte Form.

## 2. Was der Sensor beim ersten Lauf gefunden hat

Der Bestand war besser als befürchtet, aber nicht sauber. Wichtiger: **zwei der ersten
Befund-Wellen waren Fehlalarme meines eigenen Musters** — der Sensor musste gegen die Realität
kalibriert werden, nicht die Realität gegen den Sensor:

| Lauf | Meldung | Bewertung |
|---|---|---|
| 1 | slice-001/002/003 mit „2 Closure-Abschnitte" | **Fehlalarm** — `## Closure-Trigger` ist ein legitimer eigener Abschnitt, keine Notiz |
| 1 | slice-040/041 „Platzhalter" | **Fehlalarm** — der Regex `^_\(.*\)_$` traf jede kursive Klammer, auch substanziellen Text |
| 2 | slice-018/027/029/030/031 „kein Closure-Abschnitt" | **Fehlalarm** — die Bestandsform heißt mehrheitlich schlicht `## N. Closure`, nicht `Closure-Notiz`; mein verschärfter Regex war zu eng |
| 3 | die fünf unten | **echt** |

Endstand des Musters: Überschrift enthält `Closure`, ausgenommen `Closure-Trigger` und
`Closure-Kriterien`; Platzhalter sind benannte Wendungen statt jeder Kursivklammer.

### Die fünf echten Befunde — und wie sie behoben wurden

| Slice | Befund | Behebung |
|---|---|---|
| **slice-037** | Platzhalter-§9 neben der echten Notiz in §10 | Platzhalter entfernt — **heute selbst hinterlassen**, gleiche Klasse wie slice-044 |
| **slice-044** | dieselbe Doppelung | Platzhalter entfernt (vor dem Sensor-Lauf bemerkt) |
| **slice-038** | Abschluss-Inhalt existiert, heißt aber `Verifikations-Notiz` | Überschrift auf `## 8. Closure (Verifikations-Notiz, …)` — reine Umbenennung, kein neuer Text |
| **slice-040** | einsätzige Notiz ohne Abschluss-Beleg | Release-Fakt ergänzt (**v0.15.0**, aus `CHANGELOG.md` belegt) |
| **slice-041** | einsätzig **und** überholt („noch unveröffentlicht") | korrigiert: **ausgeliefert mit v0.16.0**, mit Hinweis auf die Futur-Falle |
| **slice-026** | gar kein Abschluss-Abschnitt | Closure nachgetragen, **aus dem Repo-Stand belegt** (§4 des Slice, slice-027 §9, [ADR-0020](../../adr/0020-mehr-wurzel-phantom-guard.md)/[ADR-0022](../../adr/0022-datei-mengen-bewusste-mehr-wurzel-aufloesung.md)), nicht rekonstruiert |

Jeder Nachtrag trägt den Vermerk, dass er in slice-050 entstand. Das ist Absicht: eine still
nachgetragene Historie wäre schlimmer als die Lücke.

## 3. Bewusst nicht geprüft

- **Offene DoD-Haken in `done/`.** Naiv geprüft gäbe das zwei False-Positives: slice-017 führt
  einen **Dauer-Merker** als bewusst offenen Punkt, slice-039 einen Vorgang in einem fremden Repo.
  Ein belastbarer Check braucht erst einen deklarierten Marker für solche Punkte — Form-Arbeit,
  Etappe D. Ein Sensor, der sofort Rauschen produziert, wird abgeschaltet und nicht repariert.
- **Inhalt vs. Floskel.** Semantisch, gehört dem Skill (B-4). Das Struktur-Gate kennt nur
  Überschrift, Platzhalter-Wendungen, Floskel-Liste und Satzzahl — und sagt das in seinem Kopf.
- **`check-references`.** Die Baseline hängt es in `verify` ein. In a-check ist es **bereits**
  vorhanden: das `matrix`-Modul von `doc-check` erzwingt `spec-straten → adr/slice: false`,
  `adr → slice: false` und `status.forbidden: [superseded, deprecated]`
  ([MR-005](../../../../harness/conventions.md#mr-005--referenzmatrix-intra-spec-richtung--adrslice-disziplin-d-check-angleichung)).
  Ein zweiter Prüfer wäre Doppelung. **Korrektur an slice-048:** dessen §4 behauptete, „die
  maschinelle Hälfte — `check-references` — fehlt". Das war falsch; sie sitzt nur in `gates`
  statt in `verify`.

## 4. Betroffene Module

- [`AGENTS.md`](../../../../AGENTS.md) — §5 Closure-Pflicht (neuer Anker), §4 zwei Targets,
  §6 Abschluss-Schritt.
- `tools/verify-closure-notes.sh` + [`Makefile`](../../../../Makefile) — Struktur-Gate und die
  `verify`-Aggregation.
- [`.harness/skills/closure-note-reviewer.md`](../../../../.harness/skills/closure-note-reviewer.md)
  — semantische Hälfte, aus der vendored Vorlage **übersetzt** statt kopiert (die Vorlage bindet
  an eine Kurs-ADR-Nummer und an Python-Tooling; beides gibt es hier nicht).
- Sechs Slices in `done/` — Nachträge aus §2.

## 5. DoD

- [x] Closure-Pflicht in [`AGENTS.md`](../../../../AGENTS.md) §5 verankert; `make verify` existiert
      als eigene Schicht neben `gates` und prüft sie strukturell mit Selbsttest gegen ein totes
      Muster (B-3).
- [x] [`.harness/skills/closure-note-reviewer.md`](../../../../.harness/skills/closure-note-reviewer.md)
      angelegt, auf a-checks reale Anker übersetzt, mit ausgewiesener Abweichung von der Vorlage (B-4).
- [x] `make verify` **und** `make gates` grün; die fünf echten Befunde aus §2 behoben, jeder
      Nachtrag als solcher gekennzeichnet.

## 6. Closure-Notiz

**Geliefert:** `make verify` als eigene Schicht neben `gates`, die Closure-Pflicht als normativer
Anker in [`AGENTS.md`](../../../../AGENTS.md) §5, der übersetzte
[`closure-note-reviewer`](../../../../.harness/skills/closure-note-reviewer.md)-Skill und sechs
behobene Befunde im Bestand.

**Lerneintrag — Form: geschärfte Regel.**
> **Ein neuer Sensor über einem gewachsenen Bestand meldet zuerst sich selbst.** Von vier
> Befund-Wellen waren drei Fehlalarme des eigenen Musters — die Bestandsform hieß anders, als ich
> annahm (`## N. Closure`, nicht `Closure-Notiz`), und ein Nachbar-Abschnitt (`Closure-Trigger`)
> sah aus wie eine Doppelung. Wer beim ersten roten Lauf den Bestand korrigiert statt das Muster,
> schreibt die Realität auf den Sensor um. Prüfregel: *jede Befundklasse eines neuen Sensors
> einmal einzeln ansehen, bevor irgendetwas am Prüfgegenstand geändert wird* — hier hätte die
> Abkürzung fünf Slices fälschlich „repariert".

**Zwei beobachtbare Closure-Kriterien:**

1. `make verify` und `make gates` grün auf demselben Stand (je Exit 0) — belegt.
2. Alle 45 Slices in `done/` tragen genau einen ausgefüllten Closure-Abschnitt; jeder der sechs
   Nachträge ist im Ziel-Slice als Nachtrag gekennzeichnet und aus dem Repo-Stand belegt.

**Benannte Spec-Lücke (zweiter Lerneintrag):** die Closure-Pflicht war 45 Slices lang gelebte
Praxis ohne Regel. Solche stillen Setzungen fallen erst auf, wenn jemand einen Sensor dafür bauen
will — der Sensor ist damit auch ein Aufdecker fehlender Normen, nicht nur ihr Vollstrecker.

**Folge-Slice:** [slice-051](../in-progress/slice-051-workflow-und-freigabe.md) (Etappe E, 3/3).
