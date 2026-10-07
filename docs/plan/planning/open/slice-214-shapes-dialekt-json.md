# slice-214 — `shapes`-Dialekt `json`: Implementierung

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `git mv` (`make slice-mv`).

**Welle:** welle-17 — [Welle-Plan](../welle-17-shapes-generischer-dialekt.md).

**Bezug:** [AC-FA-RULE-012](../../../../spec/lastenheft.md#ac-fa-rule-012) (Lastenheft 0.29.0),
[ADR-0042](../../adr/0042-shapes-benannte-dialekte-gomod-json.md).

**Berührte Spec-Stellen:** `spezifikation.md`
§[SPEC-EXTRACT-001](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion)
(Absatz „Dialekt `json`", `literal`-Einträge in Quellform) — umgesetzt, nicht geändert.

**Verantwortlich:** —

**Autor:** Claude. **Datum:** 2026-10-07.

**Lerneintrag — Form:** wird bei Closure benannt.

---

## 1. Ziel und Abgrenzung

**Ziel:** `shapes` prüft JSON-Dateien wie `package.json` mit `dialect: json` nach [SPEC-EXTRACT-001](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion):
Lexik nach RFC 8259, Anweisung = Mitglied des Wurzel-Objekts, `literal`-Einträge als Mitglied in
Quellform (in `{ }` eingeschlossen) — jede Gegenprobe und jeder Exit-2-Fall als Test.

**Übernimmt:** den `json`-Teil von slice-213 (Teilung vor dem Start, weil die beiden Formate
unabhängige Lexik tragen — §1 von slice-213 sah das vor).

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Die Ausgabe-Regel (einzeilige, umkehrbare Meldung, `shape-unused`-Präfix).** *Ein anderer
  Slice übernimmt es*: slice-213, weil sie dialekt-übergreifend ist und vor diesem Slice liegt.
- **Vertragsänderungen.** *Schicht-Abgrenzung*: der Vertrag steht; eine Lücke geht als
  Plan-Änderung vor den Code.

## 2. Definition of Done

Regeln dieser Sektion: Baseline-Regelwerk `modul-05-planning-harness.md`
§Ziel-Form: Slice — ≤ 3 Liefer-Punkte.

- [ ] Normalisierer `json` (Lexik, Zerlegung in Wurzel-Mitglieder, alle Exit-2-Fälle der
      Spezifikation) mit Tests je Fall und Mutations-Gegenprobe.
- [ ] `literal`-Einträge in Quellform (Einschluss in `{ }`), End-to-End-Tests mit den
      Gegenprobe-Fällen des Akzeptanzkriteriums „Boundary (`json`)" an einer `package.json` in
      realer Form.
- [ ] `--print-config`, Benutzerhandbuch, CHANGELOG `[Unreleased]` (den Vermerk „noch nicht
      implementiert" für `json` umschreiben).
- [ ] Unabhängiger Review, Report unter [`docs/reviews/`](../../../reviews/README.md).
- [ ] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [ ] Beobachtungs-Register fortgeschrieben (oder „keine Beobachtung" notiert).
- [ ] Jedes Risiko aus §6 trägt einen Ausgang.

`make gates` grün.

## 3. Plan (vor Code)

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| `internal/adapter/driven/extract/json_shape.go` + Registry | neu | Dialekt `json` |
| Tests (Normalisierer, Konfiguration, End-to-End) | neu | Gegenprobe-Fälle, Exit-2-Fälle |
| `internal/cli/cli.go`, `docs/user/benutzerhandbuch.md`, `CHANGELOG.md` | update | öffentlicher Vertrag |

## 4. Trigger

**Start** (`next` → `in-progress`): slice-213 in `done/`, WIP-Limit frei.

**Rückführungen:** `in-progress` → `next` (zu groß): entfällt voraussichtlich — ein Format mit
kleiner Grammatik. `in-progress` → `open` (blockiert): ein Gegenprobe-Fall ist mit dem Vertrag
nicht erfüllbar — zurück an den Planner.

## 5. Closure-Trigger

DoD vollständig; jede Testbehauptung mit Mutations-Gegenprobe; `make gates` Exit 0;
Closure-Notiz mit Lerneintrag.

## 6. Risiken und offene Punkte

- **Mitglieds-Grenze verrutscht:** ein Komma in einer Zeichenkette oder in einem verschachtelten
  Objekt als Mitglieds-Trenner gelesen, würde ein Mitglied zerteilen — fail-safe (rot), aber ein
  Fehlalarm an realen Dateien. — **Ausgang:** *(bei Closure zuzuweisen: eingetreten / entfallen /
  weiter offen)*

## 7. Closure-Notiz

*(folgt bei Closure)*

## 8. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** `ADAPT` (2, 3 ✓) · `USER` (2, 3 ✓).

**Vorgelagert — offene Beobachtungen sichten** (2026-10-07): beim Übergang nach `next/` erneut;
heute relevant `BEO-GATE/testbeschreibung-weiter-als-assertion` (2×) — jeder Testkommentar wird
gegen seine Assertion gelesen.

**Modus-Begründungsblock:** alle berührten Sub-Areas GF.
