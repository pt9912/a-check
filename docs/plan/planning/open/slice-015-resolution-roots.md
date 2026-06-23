# slice-015 — Resolution-Roots: Import-Auflösung gegen konfigurierbare Wurzeln (Backlog, gated)

**Status:** open (Backlog — gated auf JVM-Konsumenten-Adoption).
**Bezug:** setzt [ADR-0014](../../adr/0014-resolution-roots.md) um; erweitert
[AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml)
(Schema) + [AC-FA-EXTRACT-001](../../../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion);
ehrliche Heuristik-Grenze [AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze).
[Roadmap](../in-progress/roadmap.md). **Evidenz:** x-wal/b-cad-Pilot (JVM-Auflösungslücke).

> **Backlog-Stub.** Kein Entwurf zur Abnahme — getrackte Folge-Arbeit zu
> [ADR-0014](../../adr/0014-resolution-roots.md), damit die JVM-Auflösungsgrenze nicht stille
> Annahme bleibt. Wird zum Slice ausgearbeitet, sobald der Trigger (§1) feuert.

## 1. Auslöser (Gate)

Die heutige Auflösung nimmt „Import = wurzel-relativer Pfad" an — bricht für JVM (gepunktete
Pakete, [ADR-0014](../../adr/0014-resolution-roots.md) Kontext). Bedarf wird real, sobald ein
JVM-Konsument (x-wal o. a.) a-check adoptiert und die paket-feinen Rules (`driving`/`driven`,
`port-impurity`) braucht. **C++ (b-cad) funktioniert schon ohne** (pfad-artige Includes) — der
Trigger ist primär JVM.

## 2. Geplanter Umfang (a-check-seitig)

1. **`resolution`-Block** in `.a-check.yml`
   ([SPEC-CONF-001](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema)):
   `roots` (Auflösungs-Wurzeln) + `package_base` (gepunktete Pakete normalisieren); strict-decode.
2. **Separator-agnostische Auflösung** in `rules.go` (`targetLayer`/`segIndex`): `.` wie `/`,
   wurzel-relativ; **Default unverändert** (Import-als-Pfad) ohne `resolution`-Block.
3. Tests: JVM-Paket→Layer; C++ `src`-Root; Default-Rückwärtskompat (Go/C++).

## 3. Vor der Umsetzung zu klären

- Manifest-Ableitung (CMake/gradle automatisch lesen) — dieses Inkrement oder später?
  ([ADR-0014](../../adr/0014-resolution-roots.md) Re-Eval-Trigger.)
- x-wals Paket↔Verzeichnis-Divergenz (`driving`/`driven` im *Verzeichnis*, nicht im Paket) —
  wie mappt der `resolution`-Block das?
