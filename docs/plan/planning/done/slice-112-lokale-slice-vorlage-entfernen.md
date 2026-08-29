# slice-112 — Die lokale Slice-Vorlage entfällt

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.
**Autor:** Claude (Opus 5), im Auftrag des Maintainers. **Datum:** 2026-08-29.
**Berührte Spec-Stellen:** — *(keine)*
**Deckt:** keine `AC-*`/`ADR-*`.
**Bezug:** Maintainer-Vorgabe 2026-08-29, zweimal gestellt.
[Roadmap](../in-progress/roadmap.md).

---

## 1. Ziel

`docs/plan/planning/slice.template.md` entfällt; die Ziel-Form steht netzlos unter
`.harness/baseline/v5.12.0/templates/docs/plan/planning/slice.template.md`.

**Zur Vorgeschichte, weil sie den Schnitt bestimmt:** Ich hatte gegen die Entfernung argumentiert
und die Vorlagen-README zitiert — *„wiederkehrend ⇒ als `.template.md` co-located im Repo
behalten"*. Der **Zweck** dieser Regel ist, dass eine Ziel-Form dort verfügbar ist, wo Instanzen
entstehen. Das ist sie: der vendored Baum liegt committet im Repo und ist netzlos lesbar. Was die
lokale Kopie zusätzlich leistete, war Drift — sie stand bis slice-107 auf der `v3.5.2`-Gliederung,
während der Stand längst `v5.12.0` war.

## 2. Definition of Done

- [x] Die Datei ist entfernt; die **drei** lebenden Verweise zeigen auf den vendored Pfad —
      Beleg: `grep` und `doc-check`.
- [x] An beiden Doku-Stellen steht, was beim Kopieren **anzupassen** ist: die Zeile
      `Lerneintrag — Form: …` und die vier Felder, die a-check nicht führt — Beleg: Diff.
- [x] [`AGENTS.md`](../../../../AGENTS.md) §4 beschreibt `verify-slice-form` korrekt; die Aussage
      „kein Gate-Lauf als DoD-Punkt" ist seit slice-107 falsch — Beleg: Diff.

- [x] `make gates` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.
- [x] Closure-Notiz mit benanntem Lerneintrag geschrieben (§7).
- [x] Beobachtungs-Register fortgeschrieben.
- [x] Jedes Risiko aus §6 trägt genau einen Ausgang.

## 3. Plan (vor Code)

| Datei | Änderungs-Art | Begründung |
|---|---|---|
| `docs/plan/planning/slice.template.md` | entfällt | Ziel-Form liegt vendored |
| [`AGENTS.md`](../../../../AGENTS.md) §5 | update | Verweis + Kopier-Hinweis |
| [`AGENTS.md`](../../../../AGENTS.md) §4 | update | falsche Sensor-Beschreibung |
| `docs/plan/planning/README.md` | update | Verweis + Kopier-Hinweis |
| `tools/verify-slice-form.sh` | update | Pfad in der FAIL-Meldung |

**Auszuführende Gates:** `make gates` — tragend `doc-check` (drei Verweise ziehen um) und
`doc-targets` (§4 wird angefasst). Zum Abschluss `make verify`.

## 4. Trigger

**Start:** unmittelbar; die Vorgabe steht zweimal.

**Rückführungen:** `in-progress` → `open`, falls sich beim Umhängen zeigt, dass ein Sensor die
lokale Datei **braucht** statt sie nur zu nennen.

## 5. Closure-Trigger

Datei weg, drei Verweise umgehängt, Kopier-Hinweis steht, §4 korrigiert, Gates grün.

**Was bewusst nicht getan wird:** Die **Carveout-Vorlage** bleibt. Sie ist strukturell unverändert
gegenüber der Ziel-Form (slice-099 gemessen) und driftet damit nicht; ihr Fall wäre eine eigene
Entscheidung mit eigener Begründung, nicht ein Analogieschluss.

## 6. Risiken und offene Punkte

- *Der vendored Pfad ist tag-gescopt* — jeder Baseline-Sprung verschiebt drei Verweise, die die
  lokale Kopie nicht hatte. **Ausgang:** weiter offen, `BEO-020` im Beobachtungs-Register. Das ist
  der Preis, den die co-located-Regel vermeiden wollte; er ist klein und jetzt benannt.
- *Wer aus der vendored Vorlage kopiert, bekommt DoD-Punkte für Artefakte, die a-check nicht
  führt* (Reconciliation-Register, drei Paarungen) und **keine** Zeile `Lerneintrag — Form: …`,
  die `verify-closure-notes` verlangt. **Ausgang:** entfallen — gestrichen mit Begründung: der
  Kopier-Hinweis an beiden Verweis-Stellen nennt genau diese zwei Anpassungen, und die drei Formen
  stehen ohnehin in [`AGENTS.md`](../../../../AGENTS.md) §5.

## 7. Closure-Notiz

**Geliefert:** die lokale Vorlage ist weg; drei lebende Verweise zeigen auf die vendored Ziel-Form
und nennen die zwei Anpassungen, die beim Kopieren nötig sind. Dazu eine falsche Sensor-Beschreibung
in `AGENTS.md` §4 korrigiert.

**Lerneintrag — Form: geschärfte Regel.** *Eine Regel wird an ihrem Zweck geprüft, nicht an ihrem
Wortlaut — und wer sie zitiert, trägt die Beweislast für beides.* Ich hatte die Entfernung mit
*„wiederkehrend ⇒ co-located behalten"* abgelehnt. Der Wortlaut stimmte; der **Zweck** — eine
Ziel-Form dort, wo Instanzen entstehen — war längst erfüllt, weil der vendored Baum committet und
netzlos im Repo liegt. *Weil* ein wörtliches Zitat wie ein Argument aussieht, auch wenn es die
Frage nicht beantwortet, ob die Bedingung der Regel schon anders erfüllt ist. Die lokale Kopie hat
in der Zwischenzeit genau das getan, wovor die Baseline sonst warnt: sie ist gedriftet und stand
bis slice-107 auf einer Gliederung, die es nicht mehr gab.

**Zwei beobachtbare Closure-Kriterien:**

1. `doc-check` 226 Dateien 0 Befunde **nach** dem Löschen — die drei toten Zeiger in `done/` und
   in zwei Review-Reports sind entlinkt, ihre **Aussage** aber unverändert: Hard Rule 3.7 verbietet
   das Umschreiben von Gewesenem, nicht das Entfernen eines Pfades, dessen Ziel es nicht mehr gibt.
2. `AGENTS.md` §4 beschreibt `verify-slice-form` wieder korrekt; die Behauptung „kein Gate-Lauf als
   DoD-Punkt" stand dort noch, obwohl slice-107 die Prüfung zurückgebaut hat.

**Offene Risiken und ihr Ausgang:**

- *Die Ziel-Form liegt nur noch tag-gescopt* — jeder Baseline-Sprung verschiebt drei Verweise.
  **Ausgang:** weiter offen, `BEO-020` im Register. Das ist der Preis, den die co-located-Regel
  vermeiden wollte; er ist klein und jetzt benannt statt übersehen.
- *Wer aus der vendored Vorlage kopiert, bekommt Felder für Artefakte, die a-check nicht führt* —
  **Ausgang:** gestrichen mit Begründung: beide Verweis-Stellen nennen die zwei Anpassungen, und
  die drei Lerneintrag-Formen stehen ohnehin in `AGENTS.md` §5.

**Beobachtungs-Register:** `BEO-020` neu.

**Folge-Slices:** keine.

## 8. Sub-Area-Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** berührt werden **Planungs-Harness**, **Harness-Einstieg**
und die **Gate-/Werkzeug-Schicht** — drei, aber nur eine trägt Substanz; die anderen zwei je einen
Verweis. Ausdifferenzierung wäre Zeremonie ohne Erkenntnis.

**Vorgelagert — offene Beobachtungen sichten:** `BEO-015` (vier nicht adoptierte Felder) betrifft
denselben Gegenstand und bleibt offen; `BEO-002` (Schichten-Zahl) ist hier einschlägig und
ausdrücklich Review-Sache.

Alle berührten Sub-Areas GF.
