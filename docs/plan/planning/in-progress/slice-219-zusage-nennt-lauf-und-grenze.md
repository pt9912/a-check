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

**Lerneintrag — Form:** geschärfte Regel.

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

**Plan-Änderung 2026-10-07 (Delta-Review D-1, vor dem Fix):** In den Spec-Straten (Lastenheft,
Spezifikation, Architektur) ist eine Zusage der Vertrag selbst; ihren Lauf nennt der umsetzende
Slice (Beleg-Pflicht, `harness/conventions.md` §Anforderungs-Anlege-Prozess), nicht der Satz —
`spec/architecture.md` bleibt technologiefrei (`AGENTS.md` §3.4). Die Regel bindet dort den, der
die Zusage hineinschreibt: keine Eigenschaft in den Vertrag, für die kein Lauf geplant ist (der Fall
slice-216).

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Ein Sensor.** *Bestand bleibt bewusst stehen*: ob ein Prosa-Satz den Geltungsbereich seines
  Prüfers trifft, ist ein Urteil über zwei Formulierungen (AGENTS §3.7); der Register-Eintrag
  begründet das.
- **Nachzug alter Zusagen.** *Ein anderer Vorgang*: die sechs Fälle sind behoben; ein Durchgang über
  alle Dokumente wäre eine Kampagne.

## 2. Definition of Done

- [x] Die Regel mit Herkunfts-Anker `seit slice-219` an ihrem Ort; Zählstellen nachgezogen.
- [x] Herleitung mit den sechs Fällen im Reviewer-Skill; Register-Stand *verkörpert*.
- [x] Unabhängiger Review, Report unter [`docs/reviews/`](../../../reviews/README.md).
- [x] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [x] Beobachtungs-Register fortgeschrieben (oder „keine Beobachtung" notiert).
- [x] Jedes Risiko aus §6 trägt einen Ausgang.

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
  nicht der Schreibende. — **Ausgang:** *weiter offen* — ob die Regel wirkt, zeigt die nächste
  Zusage, die nach ihr geschrieben wird; ein neuer Beleg unter
  `BEO-GATE/zusage-weiter-als-ihre-durchsetzung` im Beobachtungs-Register nach slice-219 ist genau
  dieser Fall.

## 7. Closure-Notiz

**Lerneintrag — Form: geschärfte Regel.** Sechsmal sagte ein Dokument oder Kommentar eine Prüfung
oder Eigenschaft zu, die kein Lauf hielt. Die Regel steht als fünfte Mess-Regel in
[`harness/rules/mess-regeln.md`](../../../../harness/rules/mess-regeln.md) (`seit slice-219`): eine
Zusage nennt den Lauf, der sie hält, und seine Grenze am Satz — als Text oder Zeiger; ohne Lauf
sagt der Satz *geplant* oder *vom Review getragen*; in den Spec-Straten nennt den Lauf der
umsetzende Slice, und keine Eigenschaft kommt ohne geplanten Lauf in den Vertrag. Herleitung mit
drei Ausprägungen im Reviewer-Skill. Kein Sensor: Urteil über Text.

**Geliefert:** Regel 5, Herleitung, Zählstellen (`AGENTS.md` §5, Titel und Einleitung der
Regel-Datei, Skill-Kopf), Register *verkörpert*; zwei `state.md` ohne geführten Zähler.

**Was hat funktioniert:** Drei verschieden gebaute Suchen nach den Zählstellen — die dritte fand
den Register-Zustand, den die ersten zwei nicht trafen. Und die Gegenzählung der eigenen Zahl
(„vier der sechs Fälle fand der Review" → fünf) vor dem Commit.

**Was ging anders als geplant:** Die Reichweite der Regel brauchte zwei Anläufe: erst zu eng (nur
Prüf-Zusagen, zwei Belege fielen heraus), dann zu weit (alle Eigenschafts-Zusagen, auch in den
Spec-Straten). Und die Herleitung paraphrasierte Belege in Zitatzeichen — wie schon bei slice-215.
Beide Klassen stehen jetzt im Register.

**Kein CHANGELOG-Eintrag:** die Regel bindet den Harness-Lauf, nicht den Konsumenten des Werkzeugs
— kein öffentlicher Vertrag im Sinn von `AGENTS.md` §6 Schritt 7 (Review F-10).

**Steering-Loop-Eintrag:** geschärfte Regel — liegt in `harness/rules/mess-regeln.md` Regel 5.
Auslöser: `BEO-GATE/zusage-weiter-als-ihre-durchsetzung` (slice-186, slice-187, slice-216,
slice-217, slice-218, slice-220 — 6×).

**Beobachtungs-Register (`../observations/`):** `BEO-GATE/zusage-weiter-als-ihre-durchsetzung` →
*verkörpert*; neu `BEO-HARNESS/herleitung-paraphrasiert-ihre-belege` (2×: slice-215, slice-219) und
`BEO-HARNESS/regel-reichweite-ungleich-ihren-belegen` (2×: slice-215, slice-219) — ein dritter Beleg
wäre jeweils eine Lücke.

**Folge-Slices:** keine.

**Risiken aus §6:** das eine Risiko trägt seinen Ausgang (*weiter offen*, Beobachtungs-Register).

**Drei Paarungen:** Anker — `seit slice-219` steht im Zielort (Regel-Datei) und im Skill ·
Folge-Slice — keiner · Register — alle drei genannten Pfade existieren mit nicht leerem `evidence/`.

**Trigger-Audit der aktiven MR:** [`MR-016`](../../../../harness/conventions.md#mr-016) [`MR-019`](../../../../harness/conventions.md#mr-019) [`MR-025`](../../../../harness/conventions.md#mr-025) [`MR-027`](../../../../harness/conventions.md#mr-027) [`MR-028`](../../../../harness/conventions.md#mr-028) [`MR-029`](../../../../harness/conventions.md#mr-029) [`MR-030`](../../../../harness/conventions.md#mr-030) — 0 offen (geprüft 2026-10-07).

## 8. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** `HARNESS` (Achsen 1, 2, 3 ✓).

**Vorgelagert — offene Beobachtungen sichten** (2026-10-07):
`BEO-GATE/zusage-weiter-als-ihre-durchsetzung` (6×) ist dieser Slice selbst;
`BEO-GATE/umbau-verliert-pruefung-still` (2×) ist eine Nachbarklasse — ein dritter Beleg wäre
eine eigene Lücke, kein Teil dieses Slice.

**Modus-Begründungsblock:** alle berührten Sub-Areas GF.
