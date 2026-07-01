# Review — slice-015 Resolution-Roots (sprach-parametrisch)

**Datum:** 2026-07-01 · **Slice:** [slice-015](../plan/planning/done/slice-015-resolution-roots.md) ·
**Anforderung:** [AC-FA-CONF-001](../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml)
+ [AC-FA-EXTRACT-001](../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion)
(Lastenheft/Spezifikation 0.10.0) · [ADR-0016](../plan/adr/0016-resolution-sprach-parametrisch.md)
(erweitert [ADR-0014](../plan/adr/0014-resolution-roots.md)).

**Methode:** drei perspektiven-diverse, adversarische Agent-Linsen parallel (read-only) —
*Code-Korrektheit* · *Vertrag/Spec-Konsistenz* · *Test-Abdeckung* — gegen `git diff main`; Grundwahrheit
war das reale Verhalten (`resolveImport`/`decodeResolution`). Danach Fixes + **Delta-Re-Review** via
`make gates`. Der Plan hatte zuvor eine Maintainer-Review-Runde (3 Blocker: Folge-ADR, Python≠Backend,
x-wal-Grenze) durchlaufen.

**Gesamtbewertung:** **Kein BLOCKER.** Keine HIGH-Code-Befunde (Closure-Capture, Threading,
Rückwärtskompat, Determinismus, Regression sauber). Ein wichtiger MED-Code-Fund (stiller No-Op bei
Sprach-Tippfehler) + ein HIGH-Test-Fund (Mono-Repo-Test bewies den per-Sprache-Dispatch nicht) — beide
vor Closure gefixt.

## Befunde

| # | Linse | Schwere | Befund | Status |
|---|---|---|---|---|
| C-MED-2 | Code | MED | `resolution`-Key ungeprüft gegen `languages` → Tippfehler (`koltin`) = stiller No-Op/false-green | ✅ `resolutionEntry` lehnt nicht-deklarierte Sprache ab (Exit 2) |
| C-MED-1 | Code | MED | `yamlConfig` nicht gofmt-clean (Tag-Ausrichtung) | ✅ Block auf `map[string]yamlResolution` ausgerichtet |
| C-LOW-3/4 | Code | LOW | degeneriertes `fixed-root` (ohne roots+package_base) / `path` mit roots = stiller No-Op | ✅ beide → Exit 2 (fail-closed) |
| C-LOW-5 | Code | LOW | `roots` nicht normalisiert (`"src/"` → `"src//x"`, `""` → `"/x"`) | ✅ `cleanRoots`: Trailing-Slash-Trim + leerer root → Exit 2 |
| T-F1 | Test | HOCH | Mono-Repo-Test nur 1 Sprache → per-Sprache-Dispatch unbewiesen | ✅ 2 Sprachen (kotlin+cpp, je eigener Modus) + Kontrolle „Sprache ohne resolution → unaufgelöst" |
| T-F2…F7 | Test | MED/LOW | strict-decode-Key, mehrere `roots`, `package_base`+`roots`, Extract-`Language`, reserved-Meldung, no-match-Präfix | ✅ je ein Test |
| V-LOW | Vertrag/Spec | LOW | Spec-Historie nicht monoton (0.10.0 vor 0.9.0) | ✅ hinter 0.9.0 verschoben |

## Sauber bestätigt (knapp)

- **`.`→`/` an `package_base` gebunden** (alle 3 Linsen): Code ersetzt Punkte **nur** bei gepunkteter
  Sprache; C++ behält `.`-Endungen. SPEC-CONF-001/SPEC-RULE-001/ADR-0016 formulieren identisch. Der C++-Test
  (`room.h` bleibt) und der Mono-Repo-Test (C++ `io/writer.h` → `src/io/writer.h`) nageln es fest.
- **`mode`-Enum & Governance** (Vertrag/Spec): `path`/`fixed-root` live, `relative`/`namespace` reserviert
  → Exit 2; ADR-0016 Accepted, Sitz als **Erweiterung** von ADR-0014 (`Supersedes: —`), spec verweist
  nirgends auf ADR-0016 (Präzedenz-Matrix).
- **Rückwärtskompat/Threading** (Code): fehlende Sprache/`Resolution == nil` → Zero-Config → Import
  unverändert (heutiges Verhalten). Dogfooding 0. Die 5 Backends + Bestandstests (`targetLayer` jetzt
  3-arg) unverändert.

## Delta-Re-Review

Nach den Fixes: `make gates` erneut grün — Dogfooding `0`, d-check `0`, alle Test-Pakete `ok`
(inkl. Mono-Repo-2-Sprach, Validierungs- und resolveImport-Kanten). Die Fixes sind additive Validierung
+ Tests + Refactor (Helfer gegen gocognit), kein Eingriff in die verifizierte Auflösungs-Logik.
