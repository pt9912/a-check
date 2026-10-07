# slice-214 — `shapes`-Dialekt `json`: Implementierung

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `git mv` (`make slice-mv`).

**Welle:** welle-17 — [Welle-Plan](../welle-17-shapes-generischer-dialekt.md).

**Bezug:** [AC-FA-RULE-012](../../../../spec/lastenheft.md#ac-fa-rule-012) (Lastenheft 0.29.0),
[ADR-0042](../../adr/0042-shapes-benannte-dialekte-gomod-json.md).

**Berührte Spec-Stellen:** `spezifikation.md`
§[SPEC-EXTRACT-001](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion)
(Absatz „Dialekt `json`", `literal`-Einträge in Quellform) — umgesetzt; Schritt 2 „Gültigkeit“
geändert durch die Plan-Änderung in §1 (Spezifikation 0.38.0).

**Verantwortlich:** Claude — gesetzt beim Übergang nach `next/` (Maintainer: „ja", 2026-10-07).

**Autor:** Claude. **Datum:** 2026-10-07.

**Lerneintrag — Form:** benannte Spec-Lücke.

---

## 1. Ziel und Abgrenzung

**Ziel:** `shapes` prüft JSON-Dateien wie `package.json` mit `dialect: json` nach [SPEC-EXTRACT-001](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion):
Lexik nach RFC 8259, Anweisung = Mitglied des Wurzel-Objekts, `literal`-Einträge als Mitglied in
Quellform (in `{ }` eingeschlossen) — jede Gegenprobe und jeder Exit-2-Fall als Test.

**Plan-Änderung 2026-10-07 (Review F-1/F-3/F-5, vor dem Nachlauf-Code):** Die Spezifikation
zählte nur einzelne Fehlerfälle auf; Formen, die RFC 8259 verbietet (`tru`, `01`, `"\x"`,
fehlendes Komma zwischen Mitgliedern, abschließendes Komma in einem verschachtelten Objekt,
ungültiges UTF-8), wurden angenommen — gegen [ADR-0042](../../adr/0042-shapes-benannte-dialekte-gomod-json.md)
Entscheidung 2 („fail-closed, wo die Quelle eine Form verbietet“). Der Slice nimmt die
Präzisierung mit: die Datei muss **gültiges JSON** sein (Spezifikation 0.38.0), geprüft von einem
Grammatik-Prüfer über die Tokens.

**Übernimmt:** den `json`-Teil von slice-213 (Teilung vor dem Start, weil die beiden Formate
unabhängige Lexik tragen — §1 von slice-213 sah das vor).

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Die Ausgabe-Regel (einzeilige, umkehrbare Meldung, `shape-unused`-Präfix).** *Ein anderer
  Slice übernimmt es*: slice-213, weil sie dialekt-übergreifend ist und vor diesem Slice liegt.
- **Vertragsänderungen über die Plan-Änderung oben hinaus.** *Schicht-Abgrenzung*: der Vertrag
  steht; eine weitere Lücke geht als eigene Plan-Änderung vor den Code.

## 2. Definition of Done

Regeln dieser Sektion: Baseline-Regelwerk `modul-05-planning-harness.md`
§Ziel-Form: Slice — ≤ 3 Liefer-Punkte.

- [x] Normalisierer `json` (Lexik, Zerlegung in Wurzel-Mitglieder, alle Exit-2-Fälle der
      Spezifikation) mit Tests je Fall und Mutations-Gegenprobe.
- [x] `literal`-Einträge in Quellform (Einschluss in `{ }`), End-to-End-Tests mit den
      Gegenprobe-Fällen des Akzeptanzkriteriums „Boundary (`json`)" an einer `package.json` in
      realer Form.
- [x] `--print-config`, Benutzerhandbuch, CHANGELOG `[Unreleased]` (den Vermerk „noch nicht
      implementiert" für `json` umschreiben).
- [x] Unabhängiger Review, Report unter [`docs/reviews/`](../../../reviews/README.md).
- [x] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [x] Beobachtungs-Register fortgeschrieben (oder „keine Beobachtung" notiert).
- [x] Jedes Risiko aus §6 trägt einen Ausgang.

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
  Fehlalarm an realen Dateien. — **Ausgang:** *entfallen* — gestrichen mit Begründung: die
  Zerlegung läuft erst nach dem Grammatik-Prüfer, Kommas in Zeichenketten und verschachtelten
  Werten sind getestet (`TestJSONStringsAndNesting`), und 233 reale `package.json` unter
  `/Development` wurden ohne einen Fehlalarm zerlegt — 231 mit Befunden gegen eine leere Liste,
  2 mit Exit 2 wegen einer Zeichenkette als Wurzel (gültiges JSON, vom Dialekt ausgeschlossen);
  Pythons `json` hielt alle 233 für gültig. Der Delta-Review maß dasselbe und fuzzte 4·10^6
  ASCII-Eingaben gegen `encoding/json` ohne ungewollte Abweichung.

## 7. Closure-Notiz

**Lerneintrag — Form: benannte Spec-Lücke.** Der Vertrag des Dialekts zählte Fehlerfälle auf,
statt Gültigkeit nach der Format-Quelle zu verlangen; was die Aufzählung nicht traf, nahm der Code
an — gegen [ADR-0042](../../adr/0042-shapes-benannte-dialekte-gomod-json.md) Entscheidung 2.
[SPEC-EXTRACT-001](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion) verlangt
seit 0.38.0 gültiges JSON nach RFC 8259; die bisherigen Fälle sind Beispiele dieser Regel. Neu im
Register als `BEO-SPEC/fehlerfaelle-statt-grammatik-der-quelle` (1×).

**Geliefert:** Dialekt `json`
([AC-FA-RULE-012](../../../../spec/lastenheft.md#ac-fa-rule-012)): Lexik, Gültigkeitsprüfung,
Zerlegung in Wurzel-Mitglieder, `literal`-Einträge in Quellform; Spezifikation 0.38.0,
Handbuch 1.45, Gerüst, CHANGELOG `[Unreleased]`. Zwei Review-Läufe plus Delta-Bestätigung.

**Was hat funktioniert:** Die Plan-Änderung vor dem Code statt eines Rückbaus der Zusagen: Statt
CHANGELOG und Handbuch auf die Aufzählung zu kürzen, wurde der Vertrag auf die Quelle gehoben —
damit wurden sechs Findings mit einer Regel erledigt, und die Doku-Sätze stimmen jetzt, weil das
Verhalten ihnen folgt.

**Was ging anders als geplant:** Die Abgrenzung „der Vertrag steht" hielt nicht — die erste
Fassung des Dialekts war fail-safe, aber nicht fail-closed, und das merkte erst der Review. Der
Nachlauf brachte zwei neue Wortlaut-Findings (Kopf des Plans, Zuordnung „Wurzel kein Objekt" zur
RFC-Regel) und einen Testkommentar, der wieder mehr zusagte als prüfte. Die frühen Prüfungen
(Mitglied ohne `:`, Leerraum zwischen Werten) sind nach dem Grammatik-Prüfer unerreichbar; ihre
Mutationen sind äquivalent (Delta-Review D-6) — die Mutations-Gegenprobe der DoD gilt für den
Prüfer, nicht für diese Zweige.

**Steering-Loop-Eintrag:** benannte Spec-Lücke — verkörpert in Spezifikation 0.38.0 (Schritt 2
„Gültigkeit"), nicht an einem Zielort mit Herkunfts-Anker.

**Beobachtungs-Register (`../observations/`):**
`BEO-GATE/testbeschreibung-weiter-als-assertion` → 4×, Ausgang bleibt *geplant* (slice-215);
`BEO-SPEC/fehlerfaelle-statt-grammatik-der-quelle` neu (1×);
`BEO-ADAPT/rekursiver-pruefer-ohne-tiefengrenze` neu (1×, Delta-Review D-4 — eine Tiefengrenze wäre
eine Vertragsänderung, nicht Teil dieses Slice).

**Folge-Slices:** keine neuen; slice-215 liegt in `open/`.

**Risiken aus §6:** das eine Risiko trägt seinen Ausgang (*entfallen* mit Begründung).

**Drei Paarungen:** getragen von der Closure von welle-17.

**Trigger-Audit der aktiven MR:** [`MR-016`](../../../../harness/conventions.md#mr-016) [`MR-019`](../../../../harness/conventions.md#mr-019) [`MR-025`](../../../../harness/conventions.md#mr-025) [`MR-027`](../../../../harness/conventions.md#mr-027) [`MR-028`](../../../../harness/conventions.md#mr-028) [`MR-029`](../../../../harness/conventions.md#mr-029) [`MR-030`](../../../../harness/conventions.md#mr-030) — 0 offen (geprüft 2026-10-07).

## 8. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** `ADAPT` (2, 3 ✓) · `USER` (2, 3 ✓).

**Vorgelagert — offene Beobachtungen sichten** (2026-10-07): beim Übergang nach `next/` erneut;
heute relevant `BEO-GATE/testbeschreibung-weiter-als-assertion` (2×) — jeder Testkommentar wird
gegen seine Assertion gelesen.

**Modus-Begründungsblock:** alle berührten Sub-Areas GF.
