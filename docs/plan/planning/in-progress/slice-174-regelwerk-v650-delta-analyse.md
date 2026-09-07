# slice-174 — Regelwerk-Migration `v6.2.0` → `v6.5.0`: Delta-Analyse

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `make slice-mv` ([`AGENTS.md`](../../../../AGENTS.md)
§3.3/§5).

**Welle:** [welle-15](../welle-15-regelwerk-v650-migration.md)

**Bezug:** Maintainer-Hinweis 2026-09-07 („Neues Regelwerk-Release
`v6.5.0`"). Vorbilder derselben Form:
[slice-161](../done/slice-161-regelwerk-v610-delta-analyse.md) und
[slice-164](../done/slice-164-regelwerk-v620-delta-analyse.md).

**Berührte Spec-Stellen:** — *(keine)* — reine Ist-Messung ohne
Vertragsberührung.

**Verantwortlich:** — *(noch nicht priorisiert)*

**Autor:** Claude (Opus 5), im Auftrag des Maintainers. **Datum:**
2026-09-07.

---

## 1. Ziel

Das Delta `v6.2.0` → `v6.5.0` ist vermessen und je Änderung beantwortet, ob
sie a-check berührt — und wenn ja, mit welchem Folge-Slice. Kein Vendoring,
kein Nachzug: dieser Slice liefert die Grundlage, auf der die übrigen Etappen
von [welle-15](../welle-15-regelwerk-v650-migration.md) geschnitten werden.

## 2. Ausgangsmessung (vor der Analyse)

Am 2026-09-07 im frischen Klon von `pt9912/ai-harness-course`
(`git diff --stat v6.2.0 v6.5.0 -- lab/regelwerk lab/templates`):

| Kennzahl | Wert |
|---|---|
| Dateien | **30** |
| Zeilen | **+572 / −149** |
| Neu | `lab/templates/harness/sensors/gate.template.md` (+81) |
| Gelöscht | keine |
| Übersprungene Releases | `v6.3.0`, `v6.3.1`, `v6.4.0` |

**Größenordnung, nicht Increment.** Der vorige Sprung `v6.0.0` → `v6.2.0`
umfasste 9 Dateien und `+53/−5`
([slice-164](../done/slice-164-regelwerk-v620-delta-analyse.md) §2). Dieser ist
rund zehnmal so groß; die Analyse ist entsprechend nicht in einem Absatz zu
erledigen, und der Slice liefert bewusst **nur** sie.

Die vier größten Posten, aus denen die Etappen-Schnitte absehbar folgen:

| Datei | Δ | warum a-check betroffen ist |
|---|---|---|
| `templates/harness/sensors/gate.template.md` | **neu** | Ziel-Form je Sensor, „sobald ein Target mehr braucht als einen Satz". a-checks Sensor-Verträge stehen in **Tabellenzellen** ([`AGENTS.md`](../../../../AGENTS.md) §4, [`harness/README.md`](../../../../harness/README.md) §Sensors) — die Vorlage nennt genau diesen Überhang |
| `modul-13-quality-gates.md` | +50 | Gate-Regeln; a-check führt 20+ Targets |
| `modul-05-planning-harness.md` | +47 | Slice-Schnitt und Lifecycle |
| `templates/docs/plan/planning/slice.template.md` | +46 | neu `## 1. Ziel und Abgrenzung`, `## 8` umbenannt — trifft die Kopieranleitung in [`AGENTS.md`](../../../../AGENTS.md) §5 |

## 3. Umsetzung

*(offen — entsteht mit der Analyse)*

## 4. Definition of Done

- [ ] Jede der 30 geänderten Dateien ist eingeordnet: berührt a-check / berührt
      nicht / berührt nur eine Stelle, die a-check ohnehin anders löst — mit
      Begründung je Zeile, nicht als Sammelurteil.
- [ ] Für jede Berührung steht ein Vorschlag da: Folge-Slice (mit Titel und
      Etappe) oder ausdrückliche Nicht-Handlung mit Grund.
- [ ] Unabhängiger Review durchgeführt (Report unter [`docs/reviews/`](../../../reviews/README.md)).
- [ ] `make gates` grün.
- [ ] `make verify` grün.
- [ ] Jedes Risiko trägt einen Ausgang.

## 5. Trigger

**Start** (`open` → `in-progress`): Maintainer-Freigabe und WIP-Limit frei.

**Rückführungen:** stellt sich heraus, dass die Analyse selbst in Etappen
zerfällt — etwa weil `modul-13` und die `sensors/`-Ziel-Form getrennt zu
bewerten sind —, zurück nach `next/` zur Zerlegung. Wird ein weiteres Release
veröffentlicht, bevor die Analyse steht: zurück nach `open/`, weil dann das
Ziel der Welle neu zu setzen ist (Präzedenz: der Retarget von `welle-14`).

## 6. Closure-Trigger

DoD vollständig, `make gates`/`make verify` grün, Closure-Notiz geschrieben.
Der Slice trägt ein `**Welle:**`-Feld und archiviert **mit seiner Welle**,
nicht einzeln ([`AGENTS.md`](../../../../AGENTS.md) §6).

## 7. Risiken und offene Punkte

- *Die Analyse misst den Diff, nicht die Wirkung — eine Änderung kann klein
  aussehen und einen Vertrag brechen, wie umgekehrt* — Ausgang bei Closure.
- *Ein weiteres Release erscheint während der Welle; `v6.2.0` und `v6.5.0`
  lagen nur zwei Tage auseinander* — Ausgang bei Closure; der Präzedenzfall ist
  der Retarget von `welle-14` (`v6.1.0` → `v6.2.0` zwei Stunden nach der
  Eröffnung).
- *Die neue `sensors/`-Ziel-Form könnte a-checks Tabellen-Praxis nicht nur
  ergänzen, sondern ihr widersprechen — dann ist es keine Nachzugs-, sondern
  eine Adaptions-Frage* — Ausgang bei Closure.

## 8. Closure-Notiz

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice;
Lerneintrag — Form: wird dort benannt.)_

## 9. Sub-Area-Modus

**Vorgelagert — Sub-Area-Wahl prüfen:** eine Sub-Area berührt —
**Vendored Baseline** (`.harness/baseline/`) führt *keinen* Modus (externer
Fremdtext), und dieser Slice ändert dort nichts. Gemessen und bewertet wird
über sie hinweg; die Sub-Areas, die **Folge**-Slices berühren werden
(`HARNESS`, `GATE`, `PLAN`), gehören in deren §9, nicht hierher.

**Vorgelagert — offene Beobachtungen sichten:** Register am 2026-09-07
durchgegangen, **31** offene Einträge in `BEO-HARNESS`/`BEO-GATE`/`BEO-PLAN`,
keiner an der Schwelle. Vier einschlägig:

| Eintrag | Stand | Bezug |
|---|---|---|
| [`zwei-baseline-staende-nach-migrationsende`](../observations/BEO-HARNESS/zwei-baseline-staende-nach-migrationsende/observation.md) | offen, 1× | genau dieses Fenster — während einer Migration liegen zwei Stände zulässig nebeneinander; die Welle muss ihn beim Abschluss wieder auf einen bringen |
| [`adaption-korrigiert-repo-aussage`](../observations/BEO-HARNESS/adaption-korrigiert-repo-aussage/observation.md) | offen, 2× | der Adaptions-Durchgang wird `MR`-Einträge anfassen; entsteht dabei ein weiterer Repo-Aussage-Korrektur-Eintrag, ist die Schwelle erreicht |
| [`form-vergleich-sprachblind`](../observations/BEO-PLAN/form-vergleich-sprachblind/observation.md) | offen, 1× | die Analyse ist ein Form-Vergleich gegen Ziel-Formen — genau die Tätigkeit, bei der der Eintrag entstand |
| [`slice-form-vier-begriffe-fehlen`](../observations/BEO-PLAN/slice-form-vier-begriffe-fehlen/observation.md) | offen, 1× | `slice.template.md` ändert sich (+46); die a-check-Kopieranleitung in `AGENTS.md` §5 ist daran zu messen |

**Alle berührten Sub-Areas GF** — kein Begründungsblock nötig.
