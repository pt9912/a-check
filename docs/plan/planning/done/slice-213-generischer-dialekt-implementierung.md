# slice-213 — `shapes`-Dialekt `gomod` und einzeilige Meldung: Implementierung

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `git mv` (`make slice-mv`).

**Welle:** welle-17 — [Welle-Plan](welle-17-shapes-generischer-dialekt.md).

**Bezug:** [AC-FA-RULE-012](../../../../spec/lastenheft.md#ac-fa-rule-012) (Lastenheft 0.29.0),
[ADR-0042](../../adr/0042-shapes-benannte-dialekte-gomod-json.md).

**Berührte Spec-Stellen:** `spezifikation.md`
§[SPEC-CONF-001](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema),
§[SPEC-EXTRACT-001](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion),
§[SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung)
(in slice-212 geschrieben, hier umgesetzt).

**Verantwortlich:** Claude — gesetzt beim Übergang nach `next/` (Maintainer: „ja", 2026-10-07).

**Autor:** Claude. **Datum:** 2026-10-07.

**Lerneintrag — Form:** geschärfte Regel.

---

## 1. Ziel und Abgrenzung

**Ziel:** `shapes` prüft `go.mod` mit `dialect: gomod` nach dem in slice-212 abgenommenen
Vertrag, und die Meldung jedes `shape-*`-Befunds wird einzeilig und umkehrbar (mit Art-Präfix bei
`shape-unused`); jede Gegenprobe ist ein Test, jeder Fix hat eine Mutations-Gegenprobe.

**Plan-Änderung 2026-10-07 (Review F-5/F-11, vor dem Nachlauf-Code):** zwei Klammer-Fälle,
die der Code schon fail-closed behandelt, nennt [SPEC-EXTRACT-001](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion) nicht ausdrücklich — ein `(`/`)`
als Token mitten in einer Zeile und ein leerer Block ohne Kopf (eine Zeile nur aus `()`). Der Slice
präzisiert die Spezifikation dafür (0.37.0); keine Zusage ändert sich.

**Plan-Änderung 2026-10-07 (vor dem Start):** der Vertrag trägt zwei Formate mit unabhängiger
Lexik — nach dem Größen-Vorbehalt unten geteilt; `json` übernimmt slice-214.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Vertragsänderungen.** *Schicht-Abgrenzung*: der Vertrag steht nach slice-212; eine Lücke geht
  als Plan-Änderung vor den Code, wie in welle-16 dreimal geübt.
- **Der Dialekt `json`.** *Ein Folge-Slice übernimmt es*: slice-214 (Teilung nach dem
  Größen-Vorbehalt).
- **Ein zweiter Dateisuch-Pfad.** *Bestand bleibt bewusst stehen*: die Dateisuche samt
  Symlink-Prüfung aus welle-16 wird wiederverwendet, nicht neu gebaut — sie ist reviewt.

**Größen-Vorbehalt:** Trägt der abgenommene Vertrag zwei unabhängige Formate mit eigener Lexik,
wird dieser Slice vor dem Start je Format geteilt (Rückführung `next`).

## 2. Definition of Done

Regeln dieser Sektion: Baseline-Regelwerk `modul-05-planning-harness.md`
§Ziel-Form: Slice — ≤ 3 Liefer-Punkte.

- [x] Lexik und Zerlegung von `gomod` nach [SPEC-EXTRACT-001](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion), mit Tests je Lexik-Form und
      Mutations-Gegenprobe; nicht zerlegbare Dateien Exit 2.
- [x] Einzeilige, umkehrbare Meldung aller `shape-*`-Befunde und `shape-unused`-Präfix nach [SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung); Konfiguration (Schema aus [SPEC-CONF-001](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema), Befundform aus [SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung)), die
      Gegenprobe-Fälle „Boundary (`gomod`)" als End-to-End-Tests an einem `go.mod` in realer Form.
- [x] `--print-config`, Benutzerhandbuch, CHANGELOG `[Unreleased]` (Vermerk „noch nicht
      implementiert" für `gomod` und die Meldung umschreiben).
- [x] Unabhängiger Review, Report unter [`docs/reviews/`](../../../reviews/README.md).
- [x] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [x] Beobachtungs-Register fortgeschrieben (oder „keine Beobachtung" notiert).
- [x] Jedes Risiko aus §6 trägt einen Ausgang.

`make gates` grün.

## 3. Plan (vor Code)

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| `internal/adapter/driven/extract/` (Dialekt-Registry, neue Normalisierer) | neu / update | Lexik je Format, neben `kotlin` |
| `internal/adapter/driven/config/`, `internal/hexagon/core/shapes.go` | update | Schema und ggf. Zerlegungstiefe |
| Tests (Normalisierer, Konfiguration, End-to-End mit realen Manifest-Formen) | neu | Gegenprobe-Fälle je Format |
| `internal/cli/cli.go`, `docs/user/benutzerhandbuch.md`, `CHANGELOG.md` | update | öffentlicher Vertrag |

## 4. Trigger

**Start** (`next` → `in-progress`): slice-212 in `done/`, Folge-ADR `Accepted`, WIP-Limit frei.

**Rückführungen — vorab benannt:**

- `in-progress` → `next` (zu groß): zwei Formate mit unabhängiger Lexik — Teilung je Format.
- `in-progress` → `open` (blockiert): eine Gegenprobe ist mit dem Vertrag nicht erfüllbar —
  zurück an den Planner.

## 5. Closure-Trigger

DoD vollständig; jede Testbehauptung mit Mutations-Gegenprobe; `make gates` Exit 0;
Closure-Notiz mit Lerneintrag.

## 6. Risiken und offene Punkte

- **Lexik-Lücke im Code trotz abgeleitetem Vertrag** — die nicht fail-safe Richtung (Code als
  Kommentar verschluckt). Für JSON entfällt sie strukturell (JSON kennt keine Kommentare), für
  `go.mod` nicht. — **Ausgang:** *entfallen* — gestrichen mit Begründung: in `gomod` beginnt ein
  Kommentar nur am Token-Anfang und läuft bis zum Zeilenende, Block-Kommentare sind Exit 2 — es
  gibt keine Form, die über eine Zeile hinaus Code verschlucken könnte. Der Review fand keinen
  Fall; das Image zerlegte alle 153 `go.mod` unter `/Development` ohne Fehler. Für `json` trägt
  slice-214 das Risiko selbst.

## 7. Closure-Notiz

**Lerneintrag — Form: geschärfte Regel.** Dreimal in drei Slices sagte ein Testkommentar mehr zu,
als seine Assertion prüfte; mit diesem Slice erreicht
[`BEO-GATE/testbeschreibung-weiter-als-assertion`](../observations/BEO-GATE/testbeschreibung-weiter-als-assertion/observation.md)
3×. Ausgang *geplant*: slice-215 verkörpert die Regel als vierte Mess-Regel. Ein Sensor ist nicht
möglich — ob ein Kommentar mehr behauptet als seine Assertion, ist ein Urteil.

**Geliefert:** Dialekt `gomod`
([AC-FA-RULE-012](../../../../spec/lastenheft.md#ac-fa-rule-012),
[ADR-0042](../../adr/0042-shapes-benannte-dialekte-gomod-json.md)) und die einzeilige, umkehrbare
Meldung aller `shape-*`-Befunde mit Art-Präfix bei `shape-unused`; Spezifikation 0.37.0 (zwei
Klammer-Fälle ausdrücklich), Handbuch 1.44, Gerüst, CHANGELOG. 8 reale `go.mod` der Konsumenten
und — im Review — 153 `go.mod` ohne einen Exit 2.

**Was hat funktioniert:** Die Lexik aus der Referenz abzuleiten hat getragen: Der Review fand im
Code **keinen** Durchlass, **keinen** Fehlalarm und keine ungewollte Normalform-Kollision —
erstmals in dieser Regelart ein Slice ohne Befund im Kern der Sache. Die Mutations-Gegenprobe
fand je Fix den richtigen roten Grund.

**Was ging anders als geplant:** Die Befunde lagen in Tests und Doku: drei Testkommentare
behaupteten mehr, als sie prüften, und ein allgemeiner Handbuch-Satz („Umbrüche zählen nicht“)
wurde durch diesen Slice für `gomod` falsch. Zwei schon fail-closed behandelte Klammer-Fälle
fehlten in der Fehlerliste der Spezifikation — als Plan-Änderung nachgezogen.

**Steering-Loop-Eintrag:** geschärfte Regel, *geplant* — liegt bei Closure in slice-215, nicht
hier; gezählt, noch nicht verkörpert.

**Beobachtungs-Register (`../observations/`):** `testbeschreibung-weiter-als-assertion` → 3×,
Ausgang *geplant* (slice-215).

**Folge-Slices:** slice-214 (`json`, Welle) und slice-215 (Mess-Regel, wellenlos) — beide in
`open/`.

**Risiken aus §6:** das eine Risiko trägt seinen Ausgang (*entfallen* mit Begründung).

**Drei Paarungen:** getragen von der Closure von welle-17.

**Trigger-Audit der aktiven MR:** [`MR-016`](../../../../harness/conventions.md#mr-016) [`MR-019`](../../../../harness/conventions.md#mr-019) [`MR-025`](../../../../harness/conventions.md#mr-025) [`MR-027`](../../../../harness/conventions.md#mr-027) [`MR-028`](../../../../harness/conventions.md#mr-028) [`MR-029`](../../../../harness/conventions.md#mr-029) [`MR-030`](../../../../harness/conventions.md#mr-030) — 0 offen (geprüft 2026-10-07).

## 8. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** `KERN` (1, 2, 3 ✓) · `ADAPT` (2, 3 ✓) · `USER` (2, 3 ✓).

**Vorgelagert — offene Beobachtungen sichten** (2026-10-07):
`BEO-ADAPT/hermetik-lexikalisch-statt-am-dateisystem` (1×) — die Dateisuche wird wiederverwendet;
`BEO-GATE/testbeschreibung-weiter-als-assertion` (2×) — ein dritter Fall wäre eine Lücke: jeder
Testkommentar wird beim Schreiben gegen seine Assertion gelesen. Beim Übergang nach `next/`
erneut sichten.

**Modus-Begründungsblock:** alle berührten Sub-Areas GF.
