# slice-196 — Der `[Unreleased]`-Abschnitt holt 58 Slices nach

**Welle:** ohne Welle.

**Bezug:** Ausgang von
[`BEO-PLAN/changelog-unreleased-ungepflegt`](../observations/BEO-PLAN/changelog-unreleased-ungepflegt/observation.md)
bei **3×**. [`AC-QA-02`](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze).

**Berührte Spec-Stellen:** — (der Slice berührt kein Spec-Stratum; er beschreibt, was gelandet ist).

**Verantwortlich:** Claude — gesetzt beim Übergang nach `in-progress/`.

**Autor:** Claude, im Auftrag des Maintainers. **Datum:** 2026-09-19.

**Lerneintrag — Form:** wird bei Closure benannt.

---

## 1. Ziel und Abgrenzung

**Ziel:** Der Abschnitt `[Unreleased]` in [`CHANGELOG.md`](../../../../CHANGELOG.md) beschreibt, was
seit dem letzten Release gelandet ist — die **58** Slices von 135 bis 195. Und die Finding-Klasse
*dahinter* bekommt ihren Ausgang, den sie bei 3× verlangt.

**Ausgangslage (gemessen):** `v0.19.0..HEAD` trägt **265** Commits und **58** Slices. Der
`[Unreleased]`-Abschnitt trug bis zu diesem Slice die vier Einträge der Slices 194/195 und sonst
nichts. **Die konsumenten-sichtbare Änderung ist dabei genau eine:** `git log v0.19.0..HEAD --
internal/` nennt nur die Commits von slice-194; `spec/` änderte sich in slice-194 und slice-154
(Form-Nachzug), das Lastenheft gar nicht. Die übrigen 56 Slices sind Harness-Arbeit — für sie
führt der CHANGELOG den Abschnitt *„Harness (nicht anwender-sichtbar)"*.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Der Schnitt des Abschnitts unter eine Versionsnummer.** *Schicht-Abgrenzung:* `releasing.md`
  bindet den Schnitt an den **Re-Pin nach dem Publish** — ein Schnitt davor macht den Pin-Check in
  der Pipeline rot. Dieser Slice **füllt**, er schneidet nicht.
- **Die drei `doc-immutable`-Befunde über die Release-Range** (ADR-0017, ADR-0018, ADR-0038).
  *Es ist ein anderer Vorgang:* dort geht es um den Kern `Accepted`er ADRs, hier um den CHANGELOG.
  Eigener Folge-Slice, siehe §7.
- **Ein Eintrag je Slice.** *Es wäre ein anderes Erzeugnis:* Der CHANGELOG ist die **kuratierte
  Begründung** eines Releases, kein Slice-Protokoll; die Prosa je Slice steht in ihrer
  Closure-Notiz und in `git`.

## 2. Ausgangsmessung

| Zähler | Bau | Ergebnis |
|---|---|---|
| Slices seit dem Release | `find done -name 'slice-*.md'` ≥ 135, beide Ablageorte | **58** |
| Commits seit dem Release | `git log --oneline v0.19.0..HEAD` | **265** |
| davon mit `internal/`-Änderung | `git log v0.19.0..HEAD -- internal/` | **2 Commits, ein Slice** (194) |
| Einträge in `[Unreleased]` vorher | `sed -n '/Unreleased/,/0.19.0/p'` | **3** (die von 194/195) |

**Geltungsbereich:** die Ablageorte `done/**` dieses Repos. Slices, die noch in `open/` liegen
(189–191), zählen **nicht** — sie sind nicht gelandet.

## 3. Umsetzung

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| [`CHANGELOG.md`](../../../../CHANGELOG.md) | update | der `[Unreleased]`-Abschnitt: die konsumenten-sichtbaren Einträge bleiben; ein Block *„Harness"* beschreibt die übrigen Slices in Gruppen |
| die Modus-Deklaration oder [`AGENTS.md`](../../../../AGENTS.md) §5 | update | der **Ausgang** der Klasse bei 3× — Guide oder Sensor (Entscheidung beim Bau, §7) |

**Was der Harness-Block trägt und was nicht.** Er nennt **Zustände**, die heute im Repo prüfbar
sind — der vendored Stand, das Register auf Verzeichnisform, die neuen Sensoren, das
Archiv-Werkzeug —, und führt die Slice-Kennungen als **Herkunft** daneben. Er ist **nicht** die
Nacherzählung von 58 Closure-Notizen: eine Zusammenfassung, die ihren Gegenstand nicht gelesen hat,
ist der Fehler, den [`BEO-PLAN/dateiinhalt-aus-gedaechtnis-zitiert`](../observations/BEO-PLAN/dateiinhalt-aus-gedaechtnis-zitiert/observation.md)
beschreibt.

**Auszuführende Gates:** `make gates`, zum Abschluss `make verify`.

## 4. Definition of Done

- [ ] Der `[Unreleased]`-Abschnitt nennt jede der **58** Slices — die sichtbaren einzeln, die
      übrigen in benannten Gruppen, deren Zustand im Repo prüfbar ist.
- [ ] Die Klasse aus `BEO-PLAN/changelog-unreleased-ungepflegt` hat ihren **Ausgang** (Guide oder
      Sensor) — bei 3× ist das Pflicht, nicht Kür.

- [ ] Unabhängiger Review durchgeführt (Report unter [`docs/reviews/`](../../../reviews/README.md)).
- [ ] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [ ] Beobachtungs-Register fortgeschrieben.
- [ ] Jedes Risiko aus §7 trägt einen Ausgang.

`make gates` und `make verify` grün.

## 5. Trigger

**Start** (`open` → `in-progress`): das WIP-Limit ist frei.

**Rückführungen — vorab benannt:**

- `in-progress` → `next` (zu groß): Zeigt sich, dass der Harness-Block ohne Lesen der 58
  Closure-Notizen nicht tragfähig zu schreiben ist, wird er geteilt — die Baseline-Migrationen und
  das Archiv-Werkzeug sind zwei Vorgänge, kein einer.
- `in-progress` → `open` (blockiert): Erweist sich der Ausgang als Sensor, der einen eigenen
  Vertrag braucht (etwa ein Modul, das der Pin nicht trägt), ist das ein Pin-Vorgang.

## 6. Closure-Trigger

DoD vollständig, `make gates` und `make verify` grün, Closure-Notiz mit Lerneintrag.

## 7. Risiken und offene Punkte

- **Der Block behauptet einen Zustand, den er nicht geprüft hat.** Genau die Klasse, die dieser
  Slice im CHANGELOG *nicht* wiederholen darf; dagegen steht §3. — **Ausgang:** <offen bis Closure>
- **Der Ausgang wird zur Attrappe.** Ein Guide, der nur sagt „besser aufpassen", ist keiner; ein
  Sensor, der den Gegenstand nicht erreicht, meldet grün. — **Ausgang:** <offen bis Closure>
- **Die drei ADR-Kern-Drifts bleiben liegen.** Sie sind Item 2 der Release-Checkliste und
  **vorbestehend** (gemessen `v0.19.0..89fc7dc`, Exit 2). Dieser Slice schließt sie ausdrücklich
  aus — sie sind ein eigener Vorgang mit eigener Entscheidung (Folge-ADR oder gewollter
  Pfad-Nachzug). — **Ausgang:** <offen bis Closure>

## 8. Closure-Notiz

*(bei Closure auszufüllen)*

## 9. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** berührt sind `PLAN` (der CHANGELOG ist ein
Planungs-Artefakt: `releasing.md` führt ihn als Versionsquelle) und, je nach Ausgang, `HARNESS`
(Guide) oder `GATE` (Sensor).

**Vorgelagert — offene Beobachtungen sichten:** Register über die berührten Kürzel gelesen —
`BEO-PLAN` führt den **auslösenden** Eintrag (3×, Ausgang fällig); `BEO-PLAN/dateiinhalt-aus-gedaechtnis-zitiert`
steht als Risiko in §7.

**Modus-Begründungsblock:** alle berührten Sub-Areas GF. Kein Block je Sub-Area nötig.
