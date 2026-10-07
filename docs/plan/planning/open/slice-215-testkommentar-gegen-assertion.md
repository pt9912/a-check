# slice-215 — Regel: ein Testkommentar sagt nicht mehr zu, als seine Assertion prüft

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `git mv` (`make slice-mv`).

**Welle:** ohne Welle — die Closure-Bedingung wäre die eigene DoD.

**Bezug:** 3. Auflage von
[`BEO-GATE/testbeschreibung-weiter-als-assertion`](../observations/BEO-GATE/testbeschreibung-weiter-als-assertion/observation.md)
(slice-210, slice-211, slice-213) — Ausgang *geplant* mit diesem Slice.

**Berührte Spec-Stellen:** — · Der Slice berührt kein Spec-Stratum.

**Verantwortlich:** —

**Autor:** Claude. **Datum:** 2026-10-07.

**Lerneintrag — Form:** wird bei Closure benannt (erwartet: geschärfte Regel).

---

## 1. Ziel und Abgrenzung

**Ziel:** Die Klasse „ein Testkommentar sagt eine Eigenschaft zu, die keine Assertion prüft"
wird als Regel verkörpert, an die sich bindet, wer einen Test schreibt — als vierte Mess-Regel
neben „eine Mutations-Probe belegt erst, wenn sie rot war" in
[`harness/rules/mess-regeln.md`](../../../../harness/rules/mess-regeln.md), samt Herleitung im
Reviewer-Skill.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Ein Sensor.** *Bestand bleibt bewusst stehen*: ob ein Kommentar mehr behauptet als seine
  Assertion prüft, ist ein Urteil über den Text (AGENTS §3.7) — dieselbe Grenze wie bei den drei
  bestehenden Mess-Regeln.
- **Nachzug alter Testkommentare.** *Ein anderer Vorgang*: die drei Fälle sind behoben; ein
  Durchgang über den ganzen Testbestand wäre eine Kampagne, kein Regel-Slice.

## 2. Definition of Done

- [ ] Vierte Mess-Regel in `harness/rules/mess-regeln.md` (Wortlaut: jede Eigenschaft, die ein
      Testkommentar nennt, belegt eine Assertion; eine Mutation, die genau diese Eigenschaft
      bricht, macht den Test rot) mit Herkunfts-Anker `seit slice-215`.
- [ ] Herleitung mit den drei Fällen im Reviewer-Skill §Mess-Regeln; Register-Stand *verkörpert*.
- [ ] Unabhängiger Review, Report unter [`docs/reviews/`](../../../reviews/README.md).
- [ ] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [ ] Beobachtungs-Register fortgeschrieben (oder „keine Beobachtung" notiert).
- [ ] Jedes Risiko aus §6 trägt einen Ausgang.

`make gates` und `make verify` grün.

## 3. Plan (vor Code)

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| `harness/rules/mess-regeln.md` | update | vierte Regel |
| `.harness/skills/reviewer.md` | update | Herleitung, damit der Review sie prüft |

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
`BEO-GATE/testbeschreibung-weiter-als-assertion` (3×) ist dieser Slice selbst.

**Modus-Begründungsblock:** alle berührten Sub-Areas GF.
