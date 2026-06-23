# ADR-0014 — Resolution-Roots: Import-Auflösung gegen konfigurierbare Wurzeln

- **Status:** Proposed
- **Datum:** 2026-06-23
- **Autor:** pt9912
- **Bezug:** [AC-FA-EXTRACT-001](../../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion), [AC-FA-CONF-001](../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml), [AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze) — **Re-Evaluierung** von [ADR-0002](0002-text-heuristische-extraktion.md) (erweitert, kein Supersede).
- **Schärft:** [SPEC-RULE-001](../../../spec/spezifikation.md#spec-rule-001--regel-auswertung) + [SPEC-CONF-001](../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema) — die Symbol→Layer-Auflösung gegen explizite Wurzeln.
- **Supersedes:** —

## Kontext

[ADR-0002](0002-text-heuristische-extraktion.md) wählte Text-Heuristik (kein AST). Eine
dabei **implizite Annahme** wurde nie ausgeschrieben: *der Import-String **ist** der
wurzel-relative Pfad.* Sie hält für **Go** (der `go.mod`-Modulpfad steckt im Import,
`github.com/x/internal/core`) und **C++** (pfad-artige `#include "hexagon/model/room.h"` —
*wenn* die Scan-Wurzel = Include-Root ist). Sie **bricht für JVM:** das Kotlin/Java-Paket
`com.xwal.domain.port.input` ist *gepunktet* und **nicht** das Verzeichnis
`hexagon/ports/.../input` — `targetLayer`/`segIndex` (auf `/`) trifft es nie.

Belegt durch den **x-wal/b-cad-Pilot** (welle-06, [Roadmap](../planning/in-progress/roadmap.md)):
ein Kotlin-Hexagon-Repo, dessen Imports gepunktete Pakete sind → a-checks Layer-Auflösung
greift gar nicht, obwohl Kotlin als Sprache geführt wird.

Das Lastenheft (§1 *Out of Scope (Produkt)*) rahmt a-checks Rolle bereits: a-check **ersetzt
keine** compile-time durchgesetzte Modulgrenze (Gradle-Module), **sondern ergänzt sie um die
*feingranularen* Fitness-Functions**, die der Compiler nicht abdeckt. Konkret: x-wals Ports
sind **ein** Gradle-Modul `hexagon:ports`; die `driving`/`driven`-Trennung lebt als Pakete
*unterhalb* der Modulgrenze — genau a-checks Feld. Um sie zu liefern, muss a-check den
gepunkteten Import **auf Paket-Ebene** auflösen.

## Optionen

| Weg | Idee | Bewertung |
|---|---|---|
| **A — Resolution-Roots** | Import gegen *konfigurierbare* Wurzeln auflösen (deklariert in `.a-check.yml`, optional aus dem Build-Manifest); separator-agnostisch (`.`/`/`). | **Gewählt.** Bleibt text-heuristisch (ADR-0002-treu), löst den JVM-Fall, ist klein. |
| **B — `tree-sitter`** (Parser) | Sprach-Grammars, exakter AST. | Verworfen: schon in [ADR-0002](0002-text-heuristische-extraktion.md) (CGo/Native-Deps, Image↑); löst zudem die **semantische** Paket↔Layer-Abbildung *nicht* (nur Syntax); überdimensioniert. |
| **C — Build-Modul-Graph-Backend** | Gradle/CMake-Deps (`project(...)`, `target_link_libraries`) parsen. | Verworfen: liefert nur die **groben** Kanten, die der Compiler schon erzwingt — *dupliziert* die Modulgrenze statt sie zu ergänzen; zu grob für die paket-feine `driving`/`driven`-Trennung. |

## Entscheidung

**Weg A.** Die Symbol→Layer-Auflösung läuft gegen **Resolution-Roots**:

1. **Default rückwärtskompatibel:** Import-als-Pfad (heutiges Verhalten) — kein Bruch für Go/C++.
2. **Deklarierbar** in `.a-check.yml` (z. B. `resolution: {roots: ["src"], package_base: "com.xwal"}`):
   ein `roots`-Eintrag setzt die Wurzel, gegen die Import-Präfixe matchen; `package_base`
   normalisiert gepunktete Pakete (separator-agnostisch `.`↔`/`).
3. **Optional manifest-gestützt:** die Wurzeln *können* aus dem Build-Manifest abgeleitet
   werden (Hinweis, **nie** Regel-Backend). Jede Sprache deklariert ihre Wurzeln dort:

| Sprache | Manifest | Compile-time-Grenze (a-check *ergänzt*) | Resolution-Root-Hinweis |
|---|---|---|---|
| C++ | `CMakeLists.txt` | `target_link_libraries` (Target-Deps) | `target_include_directories` → **Include-Root** |
| JVM | `settings.gradle` / `pom.xml` | Gradle-/Maven-Module | `group` / Modul-Root → **Basis-Paket** |
| Rust | `Cargo.toml` | Crate-Grenzen | Crate-Roots |
| Go | `go.mod` | (Imports sind schon Modulpfad) | `module`-Pfad |

Das Build-System bleibt die **grobe** Grenze; a-check die **paket-/pfad-feine** Schicht darunter.

## Konsequenzen

- [ADR-0002](0002-text-heuristische-extraktion.md) **bleibt gültig** — a-check bleibt
  text-heuristisch, kein Parser. ADR-0014 entkräftet nur die *implizite* „Import = Pfad"-Annahme
  (die ADR-0002 nie ausschrieb) und ersetzt sie durch explizite Wurzeln.
- **Schema** ([AC-FA-CONF-001](../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml)/[SPEC-CONF-001](../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema)):
  optionaler `resolution`-Block; strict-decode (Exit 2 sonst).
- **JVM-Konsumenten werden prüfbar** (x-wal); b-cads `src/`-Include-Root wird *deklariert/abgeleitet*
  statt im `.a-check.yml` geraten.
- Build-Manifest (CMake/gradle/pom/Cargo) = **optionaler Resolution-Hint**, nie Regel-Backend
  (bleibt Compiler-/Build-System-Job, Lastenheft §1).

## Fitness Function

- `make test`: JVM-Paket→Layer-Auflösung über `package_base`; C++ `src`-gerootete Includes;
  **Default (Import-als-Pfad) unverändert** für Go/C++ ohne `resolution`-Block.
- `make arch-check` (Dogfooding, [AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)):
  unverändert 0 (a-check deklariert keine `resolution` → Default).

## Re-Evaluierungs-Trigger

- `tree-sitter`, falls Extraktions-*Genauigkeit* über alle Import-Formen je zum Engpass wird
  (ADR-0002 Option 3).
- **Manifest-Ableitung** (CMake-/Gradle-Wurzeln automatisch lesen) als eigenes Inkrement,
  falls die manuelle `roots`-Deklaration zu mühsam wird.

## Geschichte

| Datum | Ereignis |
|---|---|
| 2026-06-23 | Proposed — welle-06, aus dem x-wal/b-cad-Pilot (JVM-Auflösungslücke); Weg A (Resolution-Roots) gegen tree-sitter/Modul-Graph. |
