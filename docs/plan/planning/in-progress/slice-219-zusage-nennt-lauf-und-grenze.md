# slice-219 — Regel: eine Zusage über eine Prüfung nennt den Lauf, der sie hält, und seine Grenze

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `git mv` (`make slice-mv`).

**Welle:** ohne Welle — die Closure-Bedingung wäre die eigene DoD.

**Bezug:** Lese-Schritt der welle-18-Closure für
[`BEO-GATE/zusage-weiter-als-ihre-durchsetzung`](../observations/BEO-GATE/zusage-weiter-als-ihre-durchsetzung/observation.md)
(slice-186, slice-187, slice-216, slice-217, slice-218, slice-220 — 6×) — Ausgang *geplant* mit
diesem Slice.

**Berührte Spec-Stellen:** — · Der Slice berührt kein Spec-Stratum.

**Verantwortlich:** Claude — gesetzt beim Übergang nach `next/` (Maintainer: „machen wir weiter“, 2026-10-07).

**Autor:** Claude. **Datum:** 2026-10-07.

**Lerneintrag — Form:** wird bei Closure benannt (erwartet: geschärfte Regel).

---

## 1. Ziel und Abgrenzung

**Ziel:** Die Klasse „ein Dokument oder Kommentar sagt eine Prüfung zu, die der Prüfer nicht oder
nur teilweise hält" wird als Regel verkörpert, an die sich bindet, wer die Zusage schreibt: Eine
Zusage über eine Prüfung nennt den Lauf, der sie hält, und seine Grenze **am Satz** — nicht
*„wird geprüft"*, sondern *„geprüft von X; Y nicht, weil …"*. Ort: als fünfte Regel in
[`harness/rules/mess-regeln.md`](../../../../harness/rules/mess-regeln.md) oder als eigene Regel —
das entscheidet der Plan beim Start; Herleitung mit den sechs Fällen im Reviewer-Skill.

**Plan-Änderung 2026-10-07 (beim Start):** Ort ist **Regel 5** in
[`harness/rules/mess-regeln.md`](../../../../harness/rules/mess-regeln.md) — eine Zusage über eine
Prüfung ist eine Aussage über einen Beleg, und die Mess-Regeln binden jeden, der einen schreibt.
Zählstellen, mit drei verschieden gebauten Suchen gefunden: [`AGENTS.md`](../../../../AGENTS.md)
§5 Zeile 15; `mess-regeln.md` Titel, Einleitung und „alle vier"; Kopf und „Alle vier" in
§Mess-Regeln des Reviewer-Skills; der Zustand von
`BEO-GATE/testbeschreibung-weiter-als-assertion` („den drei übrigen") — er verliert die Zahl.

**Plan-Änderung 2026-10-07 (Review F-1/F-2/F-8, vor dem Fix):** Zwei der sechs Belege sagen
keine Prüfung zu, sondern eine **Eigenschaft** (slice-216: derselbe Digest, gleiche Ausgabe;
slice-218: woraus die Pipeline besteht). Die Regel fasst darum jede Zusage, dass etwas gilt oder
geschieht, nicht nur die über eine Prüfung; ihr Titel und der Zeiger in `AGENTS.md` §5 nennen
neben dem Beleg die Zusage. Ein **Zeiger am Satz** auf die Stelle, die die Grenze trägt (der
Gate-Index: „Grenzen: siehe Datei", `AGENTS.md` §4), erfüllt die Regel; die Ausprägung „Grenze
steht woanders" meint die Grenze **ohne** Zeiger.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Ein Sensor.** *Bestand bleibt bewusst stehen*: ob ein Prosa-Satz den Geltungsbereich seines
  Prüfers trifft, ist ein Urteil über zwei Formulierungen (AGENTS §3.7); der Register-Eintrag
  begründet das.
- **Nachzug alter Zusagen.** *Ein anderer Vorgang*: die sechs Fälle sind behoben; ein Durchgang über
  alle Dokumente wäre eine Kampagne.

## 2. Definition of Done

- [ ] Die Regel mit Herkunfts-Anker `seit slice-219` an ihrem Ort; Zählstellen nachgezogen.
- [ ] Herleitung mit den sechs Fällen im Reviewer-Skill; Register-Stand *verkörpert*.
- [ ] Unabhängiger Review, Report unter [`docs/reviews/`](../../../reviews/README.md).
- [ ] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [ ] Beobachtungs-Register fortgeschrieben (oder „keine Beobachtung" notiert).
- [ ] Jedes Risiko aus §6 trägt einen Ausgang.

`make gates` und `make verify` grün.

## 3. Plan (vor Code)

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| `harness/rules/mess-regeln.md`, `AGENTS.md` §5 Zeile 15 | update | Regel 5, Zählstellen |
| `docs/plan/planning/observations/BEO-GATE/…` (zwei `state.md`) | update | Ausgang *verkörpert*; Zahl aus dem Nachbar-Zustand |
| `.harness/skills/reviewer.md` | update | Herleitung |

## 4. Trigger

**Start** (`open` → `in-progress`): WIP-Limit frei.

**Rückführungen:** `in-progress` → `next` (zu groß): entfällt — eine Regel samt Herleitung.
`in-progress` → `open` (blockiert): entfällt.

## 5. Closure-Trigger

DoD vollständig, `make gates` und `make verify` grün, Closure-Notiz mit Lerneintrag.

## 6. Risiken und offene Punkte

- **Prosa ohne Wirkung:** die Regel wird gelesen oder nicht; fünf der sechs Fälle fing der Review,
  nicht der Schreibende. — **Ausgang:** *(bei Closure zuzuweisen: eingetreten / entfallen / weiter offen)*

## 7. Closure-Notiz

*(folgt bei Closure)*

## 8. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** `HARNESS` (Achsen 1, 2, 3 ✓).

**Vorgelagert — offene Beobachtungen sichten** (2026-10-07):
`BEO-GATE/zusage-weiter-als-ihre-durchsetzung` (6×) ist dieser Slice selbst;
`BEO-GATE/umbau-verliert-pruefung-still` (2×) ist eine Nachbarklasse — ein dritter Beleg wäre
eine eigene Lücke, kein Teil dieses Slice.

**Modus-Begründungsblock:** alle berührten Sub-Areas GF.
