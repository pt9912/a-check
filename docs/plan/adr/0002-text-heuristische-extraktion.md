# ADR-0002 — Text-heuristische Import-Extraktion (kein Sprach-AST)

- **Status:** Proposed
- **Datum:** 2026-06-21
- **Bezug:** [AC-FA-EXTRACT-001](../../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion), [AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze), [AC-QA-01](../../../spec/lastenheft.md#ac-qa-01--determinismus)
- **Schärft:** — (`spec/spezifikation.md` entsteht mit slice-002; siehe [ADR-Index](README.md))
- **Supersedes:** —

## Kontext

`a-check` muss vier Sprachen abdecken — C++ (`#include`), Go (`import`),
Rust (`use`/`extern crate`), Kotlin (`import`) — und konsolidiert vier
divergente `arch-check.sh` (Lastenheft §1). Pro Sprache liefert ein Backend
die Menge „welche Symbole/Module importiert diese Datei"
([AC-FA-EXTRACT-001](../../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion)).
Diese Menge ist die Eingabe der fünf — dieser ADR *nachgelagerten* —
Hexagon-Regeln (`AC-FA-RULE-*`); die begründende Anforderung dieser ADR ist
[AC-FA-EXTRACT-001](../../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion).
Der Scan muss text-basiert, netzlos und distroless bleiben
([AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)).

## Optionen

1. **Text-/Regex-Heuristik** über konfigurierbare Muster je Sprache.
   Trade-off: unvollständig (Import-ähnliche Zeilen in Kommentaren/Strings,
   gleichnamige framework-fremde Symbole) — aber die Grenze ist
   *dokumentierbar* statt verschwiegen (ehrliche Heuristik-Grenze,
   [AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)),
   und der Lauf bleibt toolchain-frei und netzlos.
2. **Sprach-Toolchains** (`go list`, clang-AST, rust-analyzer, `kotlinc`).
   Trade-off: präzise, aber je Sprache eine schwere Runtime im Image →
   bricht distroless/static + netzlos, nicht-deterministisch über
   Toolchain-Versionen. Im Lastenheft als opt-in-Re-Eval ausgewiesen
   ([AC-FA-EXTRACT-001](../../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion)
   Out-of-Scope), nicht 0.1.0.
3. **`tree-sitter`** (eine Engine, viele Grammars). Trade-off: Mittelweg,
   aber CGo/Native-Deps und ein größeres Image als die Regel-Klasse braucht.

## Entscheidung

**Option 1 — text-heuristische Extraktion über Config-Muster.** Die fünf
Hexagon-Regeln (`AC-FA-RULE-*`) arbeiten auf der Menge der importierten
Symbole je Datei; diese Menge ist text-heuristisch hinreichend gewinnbar.
Toolchain-gestützte Backends bleiben ein dokumentiertes opt-in-Re-Eval,
nicht 0.1.0.

## Konsequenzen

- Die **ehrliche Heuristik-Grenze** ist Pflicht-Output, nicht verschwiegen
  ([AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)):
  eine Allowlist/Marker-Ausnahme ist konfigurierbar (siehe
  [ADR-0003](0003-config-modell-a-check-yml.md)).
- **Fitness Function / Gate** (slice-003): `make test` mit
  Happy/Boundary/Negative je Backend
  ([AC-FA-EXTRACT-001](../../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion)),
  inkl. Negative „Import in Kommentar/String wird nicht gewertet".
- Determinismus ([AC-QA-01](../../../spec/lastenheft.md#ac-qa-01--determinismus)):
  stabile, sortierte Extraktions- und Befundreihenfolge.
