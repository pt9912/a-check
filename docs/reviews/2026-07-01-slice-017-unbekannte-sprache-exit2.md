# Review — slice-017 Unbekannte Sprache → Exit 2 (falsch-grün-Falle)

**Datum:** 2026-07-01 · **Slice:** [slice-017](../plan/planning/done/slice-017-unbekannte-sprache-exit2.md) ·
**Anforderung:** [AC-FA-CONF-001](../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml)
(neue Negative-AC), Backend-Menge aus [AC-FA-EXTRACT-001](../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion)
(Lastenheft/Spezifikation 0.9.0). Kein neuer ADR (strict-config, [ADR-0003](../plan/adr/0003-config-modell-a-check-yml.md)).

**Methode:** drei perspektiven-diverse, adversarische Agent-Linsen parallel (read-only) —
*Code-Korrektheit* · *Vertrag/Spec-Konsistenz* · *Test-Abdeckung* — gegen `git diff main`; Grundwahrheit
war das reale Verhalten (die Registry-Keys). Danach Fixes + **Delta-Re-Review** via `make gates`.
Der Slice-**Plan** hatte zuvor bereits zwei Maintainer-Review-Runden (Registry-Map als Single Source,
SPEC-EXTRACT-001-Owner) durchlaufen.

**Gesamtbewertung:** **Kein BLOCKER.** Code-Linse 0 Befunde (Closure-Capture, kein stilles Netz,
Determinismus, Exit-Mapping, Regression alle sauber). Test- und Vertrag/Spec-Linse lieferten
Härtungen — alle vor Closure eingearbeitet.

## Befunde

| # | Linse | Schwere | Befund | Status |
|---|---|---|---|---|
| C | Code | — | keine korrektheitsrelevanten Fehler; nil-Panic nach `checkLanguages` unerreichbar, Registry-Closures == alte switch-Arme | ✅ bestätigt |
| T1 | Test | MED | keine Case-/Tippfehler-Kante (`Go`) — ein späterer `ToLower`-„Fix" bliebe unentdeckt | ✅ `TestCheckLanguagesCaseSensitive` |
| T2 | Test | LOW | Fehlermeldung nie als exaktes Literal gepinnt (Klammerung/Reihenfolge könnte driften) | ✅ `TestCheckLanguagesUnknown` prüft jetzt `==` gegen das volle Literal |
| T3 | Test | LOW | „stderr, nicht stdout" nur halb belegt | ✅ `out.Len() != 0`-Check in `TestUnknownLanguageExit2` |
| T4 | Test | LOW | fehlender `languages`-Pflichtblock (vorbestehender Pfad) ohne Test | ✅ `TestMissingLanguagesBlock` |
| T5 | Test | LOW | Mono-Repo-Mixed-Test nur positionsabhängig (unsupported sortiert *nach* supported) | ✅ beidseitig: `typescript` (nach `go`) **und** `csharp` (vor `go`) + „unbekannte Sprache"-Prüfung |
| V1 | Vertrag/Spec | LOW/MED | SPEC-CLI-001/AC-FA-CLI-001 versprechen „Exit 2 **mit Zeilenangabe**" — der neue (und andere) Config-Fehler tragen keine Zeile | ✅ auf „mit Zeilenangabe, **wo die Fehlerquelle eine Zeile hat**" abgeschwächt |
| V2 | Vertrag/Spec | LOW | Menge trotz „kein Duplikat"-Anspruch inline in SPEC-CONF-001 + Negative-AC wiederholt (Drift bei 6. Backend) | ✅ Literale gestrichen, nur Verweis auf den Owner (SPEC-/AC-FA-EXTRACT-001) |

## Sauber bestätigt (knapp)

- **Registry = echte Single Source** (Code-Linse): `default: return nil` vollständig weg; eine
  unbekannte Sprache erreicht den Dispatch nach `checkLanguages` nicht (nil-Panic wäre laut, nicht
  still). Test-fixiert (`TestBackendRegistrySet`, genau 5).
- **Mono-Repo-tauglich** (Test-Linse): jeder `languages`-Key wird einzeln geprüft — eine unterstützte
  Sprache „rettet" die Config nicht; die 5 Backends extrahieren nach dem switch→Registry-Umbau
  unverändert (Bestandstests exerzieren den neuen Dispatch).
- **Vertrag konsistent** (Vertrag/Spec-Linse): Versions-Bumps (0.9.0 je einmal), Anker gültig,
  SPEC-EXTRACT-001 Owner / SPEC-CONF-001 Verweis, Spec beschreibt „Validierung vor dem Walk" = Code.

## Delta-Re-Review

Nach den Fixes: `make gates` erneut grün — Dogfooding `0`, d-check `0`, alle Test-Pakete `ok`
inkl. der neuen Härtungs-Tests. Die Fixes sind additiv (Tests + Spec-Wortlaut), kein Eingriff in
die bereits verifizierte Registry-/Validierungs-Logik.
