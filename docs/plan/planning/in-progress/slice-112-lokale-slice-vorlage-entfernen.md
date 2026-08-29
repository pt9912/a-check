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

- [ ] Die Datei ist entfernt; die **drei** lebenden Verweise zeigen auf den vendored Pfad —
      Beleg: `grep` und `doc-check`.
- [ ] An beiden Doku-Stellen steht, was beim Kopieren **anzupassen** ist: die Zeile
      `Lerneintrag — Form: …` und die vier Felder, die a-check nicht führt — Beleg: Diff.
- [ ] [`AGENTS.md`](../../../../AGENTS.md) §4 beschreibt `verify-slice-form` korrekt; die Aussage
      „kein Gate-Lauf als DoD-Punkt" ist seit slice-107 falsch — Beleg: Diff.

- [ ] `make gates` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.
- [ ] Closure-Notiz mit benanntem Lerneintrag geschrieben (§7).
- [ ] Beobachtungs-Register fortgeschrieben.
- [ ] Jedes Risiko aus §6 trägt genau einen Ausgang.

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

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice,
[`AGENTS.md`](../../../../AGENTS.md) §5; `make verify` prüft das.)_

## 8. Sub-Area-Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** berührt werden **Planungs-Harness**, **Harness-Einstieg**
und die **Gate-/Werkzeug-Schicht** — drei, aber nur eine trägt Substanz; die anderen zwei je einen
Verweis. Ausdifferenzierung wäre Zeremonie ohne Erkenntnis.

**Vorgelagert — offene Beobachtungen sichten:** `BEO-015` (vier nicht adoptierte Felder) betrifft
denselben Gegenstand und bleibt offen; `BEO-002` (Schichten-Zahl) ist hier einschlägig und
ausdrücklich Review-Sache.

Alle berührten Sub-Areas GF.
