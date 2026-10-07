# slice-215 — Regel: ein Testkommentar sagt nicht mehr zu, als seine Assertion prüft

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `git mv` (`make slice-mv`).

**Welle:** ohne Welle — die Closure-Bedingung wäre die eigene DoD.

**Bezug:** 3. und 4. Auflage von
[`BEO-GATE/testbeschreibung-weiter-als-assertion`](../observations/BEO-GATE/testbeschreibung-weiter-als-assertion/observation.md)
(slice-210, slice-211, slice-213, slice-214) — Ausgang *geplant* mit diesem Slice.

**Berührte Spec-Stellen:** — · Der Slice berührt kein Spec-Stratum.

**Verantwortlich:** Claude — gesetzt beim Übergang nach `next/` (Maintainer: „machen wir weiter“, 2026-10-07).

**Autor:** Claude. **Datum:** 2026-10-07.

**Lerneintrag — Form:** wird bei Closure benannt (erwartet: geschärfte Regel).

---

## 1. Ziel und Abgrenzung

**Ziel:** Die Klasse „ein Testkommentar sagt eine Eigenschaft zu, die keine Assertion prüft"
wird als Regel verkörpert, an die sich bindet, wer einen Test schreibt — als vierte Mess-Regel
neben „eine Mutations-Probe belegt erst, wenn sie rot war" in
[`harness/rules/mess-regeln.md`](../../../../harness/rules/mess-regeln.md), samt Herleitung im
Reviewer-Skill.

**Plan-Änderung 2026-10-07 (vor der Arbeit):** Die Zahl „drei Mess-Regeln“ steht auch in
[`AGENTS.md`](../../../../AGENTS.md) §5 Zeile 15 und im Titel der Regel-Datei; der Kopf von
§Mess-Regeln im Reviewer-Skill zählt „zwei“ und verortet die Zusage in `AGENTS.md` §5, obwohl sie
seit slice-201 in der Regel-Datei steht. Beides zieht dieser Slice mit nach — eine vierte Regel,
neben der „drei“ stehen bleibt, wäre die Aufzählungs-Drift, die die Regel selbst meidet. Ein
CHANGELOG-Eintrag entfällt wie bei den drei Vorgänger-Regeln: Mess-Regeln binden den
Harness-Lauf, nicht den Konsumenten des Werkzeugs.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Ein Sensor.** *Bestand bleibt bewusst stehen*: ob ein Kommentar mehr behauptet als seine
  Assertion prüft, ist ein Urteil über den Text (AGENTS §3.7) — dieselbe Grenze wie bei den drei
  bestehenden Mess-Regeln.
- **Nachzug alter Testkommentare.** *Ein anderer Vorgang*: die vier Fälle sind behoben; ein
  Durchgang über den ganzen Testbestand wäre eine Kampagne, kein Regel-Slice.

## 2. Definition of Done

- [ ] Vierte Mess-Regel in `harness/rules/mess-regeln.md` (Wortlaut: jede Eigenschaft, die ein
      Testkommentar nennt, belegt eine Assertion; eine Mutation, die genau diese Eigenschaft
      bricht, macht den Test rot) mit Herkunfts-Anker `seit slice-215`.
- [ ] Herleitung mit den vier Fällen im Reviewer-Skill §Mess-Regeln; Register-Stand *verkörpert*.
- [ ] Unabhängiger Review, Report unter [`docs/reviews/`](../../../reviews/README.md).
- [ ] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [ ] Beobachtungs-Register fortgeschrieben (oder „keine Beobachtung" notiert).
- [ ] Jedes Risiko aus §6 trägt einen Ausgang.

`make gates` und `make verify` grün.

## 3. Plan (vor Code)

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| `harness/rules/mess-regeln.md` | update | vierte Regel |
| `.harness/skills/reviewer.md` | update | Herleitung, damit der Review sie prüft; Kopf zählt und verortet richtig |
| `AGENTS.md` §5 Zeile 15 | update | Zahl der Mess-Regeln |
| `docs/plan/planning/observations/BEO-GATE/testbeschreibung-weiter-als-assertion/state.md` | update | Ausgang *verkörpert* mit Herkunfts-Anker |

## 4. Trigger

**Start** (`open` → `in-progress`): WIP-Limit frei.

**Rückführungen:** `in-progress` → `next` (zu groß): entfällt — eine Regel samt Herleitung.
`in-progress` → `open` (blockiert): entfällt.

## 5. Closure-Trigger

DoD vollständig, `make gates` und `make verify` grün, Closure-Notiz mit Lerneintrag.

## 6. Risiken und offene Punkte

- **Prosa ohne Wirkung:** eine vierte Mess-Regel wird gelesen wie die anderen drei — oder nicht.
  — **Ausgang:** *(bei Closure zuzuweisen: eingetreten / entfallen / weiter offen)*

## 7. Closure-Notiz

*(folgt bei Closure)*

## 8. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** `HARNESS` (Achsen 1, 2, 3 ✓).

**Vorgelagert — offene Beobachtungen sichten** (2026-10-07):
`BEO-GATE/testbeschreibung-weiter-als-assertion` (4×) ist dieser Slice selbst. Der vierte
Beleg kam vor der Verkörperung — die Prosa-Form ist damit nicht ausgeschöpft, sondern noch nicht
geschrieben; ein Sensor bleibt ausgeschlossen (§1). Weitere Treffer in `HARNESS`: keine offenen
Einträge, die dieser Slice berührt.

**Modus-Begründungsblock:** alle berührten Sub-Areas GF.
