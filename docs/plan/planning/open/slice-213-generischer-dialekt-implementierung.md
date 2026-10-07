# slice-213 — Generischer `shapes`-Dialekt: Implementierung

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `git mv` (`make slice-mv`).

**Welle:** welle-17 — [Welle-Plan](../welle-17-shapes-generischer-dialekt.md).

**Bezug:** [AC-FA-RULE-012](../../../../spec/lastenheft.md#ac-fa-rule-012) in der Fassung aus
slice-212, die dort entstehende Folge-ADR zu
[ADR-0041](../../adr/0041-shapes-sollform-je-datei.md) (Kennung wird beim Übergang nach `next/`
eingetragen).

**Berührte Spec-Stellen:** `spezifikation.md`
§[SPEC-CONF-001](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema),
§[SPEC-EXTRACT-001](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion),
§[SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung)
(in slice-212 geschrieben, hier umgesetzt).

**Verantwortlich:** —

**Autor:** Claude. **Datum:** 2026-10-07.

**Lerneintrag — Form:** wird bei Closure benannt.

---

## 1. Ziel und Abgrenzung

**Ziel:** `shapes` prüft `go.mod` und `package.json` nach dem in slice-212 abgenommenen Vertrag;
jede Gegenprobe der Anforderung ist ein Test, jeder Fix hat eine Mutations-Gegenprobe.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Vertragsänderungen.** *Schicht-Abgrenzung*: der Vertrag steht nach slice-212; eine Lücke geht
  als Plan-Änderung vor den Code, wie in welle-16 dreimal geübt.
- **Ein zweiter Dateisuch-Pfad.** *Bestand bleibt bewusst stehen*: die Dateisuche samt
  Symlink-Prüfung aus welle-16 wird wiederverwendet, nicht neu gebaut — sie ist reviewt.

**Größen-Vorbehalt:** Trägt der abgenommene Vertrag zwei unabhängige Formate mit eigener Lexik,
wird dieser Slice vor dem Start je Format geteilt (Rückführung `next`).

## 2. Definition of Done

Regeln dieser Sektion: Baseline-Regelwerk `modul-05-planning-harness.md`
§Ziel-Form: Slice — ≤ 3 Liefer-Punkte.

- [ ] Lexik und Zerlegung der neuen Formate nach [SPEC-EXTRACT-001](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion), mit Tests je Lexik-Form und
      Mutations-Gegenprobe; nicht zerlegbare Dateien Exit 2.
- [ ] Konfiguration und Auswertung (Schema aus [SPEC-CONF-001](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema), Befundform aus [SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung)), die
      Gegenprobe-Fälle der Anforderung je Format als End-to-End-Tests.
- [ ] `--print-config`, Benutzerhandbuch, CHANGELOG `[Unreleased]`.
- [ ] Unabhängiger Review, Report unter [`docs/reviews/`](../../../reviews/README.md).
- [ ] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [ ] Beobachtungs-Register fortgeschrieben (oder „keine Beobachtung" notiert).
- [ ] Jedes Risiko aus §6 trägt einen Ausgang.

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
  `go.mod` nicht. — **Ausgang:** *(bei Closure zuzuweisen: eingetreten / entfallen / weiter offen)*

## 7. Closure-Notiz

*(folgt bei Closure)*

## 8. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** `KERN` (1, 2, 3 ✓) · `ADAPT` (2, 3 ✓) · `USER` (2, 3 ✓).

**Vorgelagert — offene Beobachtungen sichten** (2026-10-07):
`BEO-ADAPT/hermetik-lexikalisch-statt-am-dateisystem` (1×) — die Dateisuche wird wiederverwendet;
`BEO-GATE/testbeschreibung-weiter-als-assertion` (2×) — ein dritter Fall wäre eine Lücke: jeder
Testkommentar wird beim Schreiben gegen seine Assertion gelesen. Beim Übergang nach `next/`
erneut sichten.

**Modus-Begründungsblock:** alle berührten Sub-Areas GF.
