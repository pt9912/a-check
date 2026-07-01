# slice-015 — Resolution-Roots: Import-Auflösung gegen konfigurierbare Wurzeln (Backlog, gated)

**Status:** open (Backlog — gated; Trigger **je Konsument**, nicht mehr JVM-only, §1).
**Bezug:** setzt [ADR-0014](../../adr/0014-resolution-roots.md) um; erweitert
[AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml)
(Schema) + [AC-FA-EXTRACT-001](../../../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion);
ehrliche Heuristik-Grenze [AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze).
[Roadmap](../in-progress/roadmap.md). **Evidenz:** b-cad-Pilot (C++ — Scan-Wurzel muss = Include-Root
`src/`, manuell/fehleranfällig, kollidiert mit dem `make a-check`-Gate) + x-wal (JVM, gepunktete
Pakete) + Polyglot-Bestand (Go ✓, Python/TypeScript/C# offen).

> **Backlog-Stub.** Kein Entwurf zur Abnahme — getrackte Folge-Arbeit zu
> [ADR-0014](../../adr/0014-resolution-roots.md), damit die Auflösungsgrenze (über **alle** Sprachen)
> nicht stille Annahme bleibt. Wird zum Slice ausgearbeitet, sobald der Trigger (§1) feuert.

## 1. Auslöser (Gate)

Die heutige Auflösung nimmt „Import = wurzel-relativer Pfad" an. Das hält **nur** für Sprachen, deren
Import *ist* der wurzel-relative Pfad (Go: Modulpfad) — und bricht in drei verschiedenen Formen:

- **Fester-Wurzel-dotted** (JVM, Python): `com.x.Y` / `a.b.c` sind gepunktet, kein `/`-Pfad — braucht
  Wurzel + Separator-Normalisierung ([ADR-0014](../../adr/0014-resolution-roots.md) Kontext).
- **Wurzel ≠ Scan-Wurzel** (C++): b-cads Includes sind `src/`-gewurzelt (`#include "hexagon/…"`); der
  **b-cad-Pilot** zeigte, dass „C++ funktioniert schon ohne" nur gilt, wenn man Scan-Wurzel = `src/`
  **von Hand** setzt — fehleranfällig und unvereinbar mit dem `make a-check`-Gate, das Repo-Root mountet.
  Ein deklariertes `roots: ["src"]` löst das, ohne den Datei-Baum zu verrenken.
- **Relativ zum importierenden File** (TypeScript: `./x`, `../lib/y`; C/C++-quoted): löst gegen den
  *Ort des Files* auf, nicht die Wurzel — **anderes Modell** als oben.
- **Namespace-entkoppelt** (C#): `using Foo.Bar;` bindet an einen Namespace, den jede Datei frei
  deklariert — kein Pfad-Bezug; braucht einen Namespace→Datei-Index.

Der Trigger feuert also **pro Konsument**, sobald dessen Import-Form nicht „Pfad = Wurzel-relativ" ist.

## 2. Geplanter Umfang (a-check-seitig) — sprach-parametrisch

1. **`resolution`-Block** in `.a-check.yml`
   ([SPEC-CONF-001](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema)):
   `roots` (Auflösungs-Wurzeln) + `package_base` (gepunktete Pakete normalisieren); strict-decode.
   **Als Map Sprache → Config** (Mono-Repo-tauglich, §3), nicht ein globaler Modus. Deckt
   **Modus fester-Wurzel** ab (Go bleibt Default, JVM/Python/C++-`src` via `roots`/`package_base`).
2. **Separator-agnostische Auflösung** in `rules.go` (`targetLayer`/`segIndex`): `.` wie `/`,
   wurzel-relativ; **Default unverändert** (Import-als-Pfad) ohne `resolution`-Block.
3. Tests: JVM-Paket→Layer; C++ `src`-Root; Default-Rückwärtskompat (Go/C++).

## 3. Vor der Umsetzung zu klären

- **Mono-Repos (mehrere Sprachen in einem Repo, z. B. Go + TypeScript):** `resolution` muss **pro
  Sprache** wählbar sein — Go löst über den Modulpfad auf, TypeScript relativ-zum-File. Der Block ist
  also eine **Map Sprache → Auflösungs-Config** (analog zum `languages`-Map-Muster), nicht ein
  globaler Modus. Die *Deklaration* mehrerer Sprachen + je-Key-Validierung ist bereits erledigt
  ([slice-017](../done/slice-017-unbekannte-sprache-exit2.md)); offen ist allein die per-Sprache-*Auflösung*.
- **Zwei weitere Auflösungs-Modi jenseits von [ADR-0014](../../adr/0014-resolution-roots.md)** —
  brauchen vermutlich je einen **Folge-ADR**, wenn ihr Pilot feuert: (a) *relativ-zum-File* (TypeScript,
  quoted C++), (b) *Namespace-Index* (C#). Beide sind kein reines `roots`/`package_base`.
- Manifest-Ableitung (CMake/gradle automatisch lesen) — dieses Inkrement oder später?
  ([ADR-0014](../../adr/0014-resolution-roots.md) Re-Eval-Trigger.)
- x-wals Paket↔Verzeichnis-Divergenz (`driving`/`driven` im *Verzeichnis*, nicht im Paket) —
  wie mappt der `resolution`-Block das?
