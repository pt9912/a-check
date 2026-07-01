# Review — slice-016 Regex-fähige `tech`-Muster (`match: substring|regex`)

**Datum:** 2026-07-01 · **Slice:** [slice-016](../plan/planning/done/slice-016-regex-tech-muster.md) ·
**Anforderung:** [AC-FA-RULE-003](../../spec/lastenheft.md#ac-fa-rule-003--tech-kapselung-regel-tech-leak)
+ [AC-FA-CONF-001](../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml)
(Lastenheft/Spezifikation 0.8.0) · [ADR-0015](../plan/adr/0015-regex-tech-muster.md)
(Re-Eval [ADR-0003](../plan/adr/0003-config-modell-a-check-yml.md)).

**Methode:** vier perspektiven-diverse, adversarische Agent-Linsen parallel (read-only) —
*Code-Korrektheit* · *Vertrag/Spec-Konsistenz* · *Test-Abdeckung* · *Regelwerk/Konvention* —
gegen `git diff main` (12 Dateien) + die zwei neuen Docs; Grundwahrheit war das reale
Code-Verhalten (`matchTech`). Anschließend Fixes + **Delta-Re-Review** via `make gates`.

**Gesamtbewertung:** **Kein BLOCKER.** Der Vertrag ist widerspruchsfrei (Vertrag/Spec: 0 Befunde);
Code + Tests waren im Kern korrekt. Zwei Linsen fanden **unabhängig** denselben realen Fehler
(leeres Regex-Pattern trifft alles) — vor Closure gefixt (fail-closed). Die Konventions-Linse
schloss den Prozess-Rahmen (Nutzerdoku-Historie, Slice-Lifecycle, Roadmap).

## Befunde

| # | Linse | Schwere | Befund | Status |
|---|---|---|---|---|
| C1 | Code + Test | MAJOR | Leeres `pattern` mit `match: regex` → `regexp.Compile("")` matcht **jeden** Import (FP-Flut; schattet als Erst-Treffer nachfolgende `tech`-Regeln); Substring-Seite guardet leeres Pattern, Regex-Seite nicht | ✅ `NewTech` guardet leeres Regex-Pattern → Exit 2; AC/Spec + Test (`TestTechMatchEmptyRegexFailsClosed`) |
| T2 | Test | MAJOR | Exit-Code-Propagation nur als Go-`err` geprüft, die neuen Pfade nie durch `cli.Run` | ✅ `TestTechRegexLeakExit1` (→1), `TestTechRegexInvalidMatchExit2` (→2) |
| T3 | Test | MITTEL | dokumentierte `Q[A-Za-z]`/`Queue.h`-False-Positive-Semantik + `ignore_symbols`-Unterdrückung ungetestet | ✅ `TestTechRegexIgnoreSymbols` (mit Marker Exit 0, ohne Exit 1) |
| T4 | Test | MITTEL | kein Determinismus-Test mit ≥2 regex-`tech-leak`-Befunden | ✅ `TestTechLeakRegexDeterministicOrder` |
| T5/T6 | Test | NIEDRIG | Rückwärtskompat + Modus-Trennung (`match: substring` nicht als Regex) nur implizit | ✅ `TestNewTechBackCompatSubstring`, `TestNewTechSubstringNotRegex` |
| K1 | Konvention | MAJOR | Benutzerhandbuch-§10-Historie + Version nicht nachgezogen | ✅ 1.10→1.11, Stand 2026-07-01, §10-Zeile |
| K2 | Konvention | MAJOR | Slice-Lifecycle: slice-016 hing in `open/` trotz Accepted-ADR/Umsetzung | ✅ Status→done, §7 DoD + §8 Closure; reiner `git mv` nach `done/` |
| K3 | Konvention | MAJOR | Roadmap-Eintrag fehlte (hängender ADR→Roadmap-Verweis) + Datum veraltet | ✅ Roadmap-Zeile slice-016/ADR-0015 + Datum |
| S1 | Vertrag/Spec | INFO | Slice §4-Code kündigte „exported `Match` + `*regexp.Regexp`" an; Code nutzt sauberer ein Closure-Feld | ✅ Slice-Wortlaut an die Umsetzung angeglichen |

## Sauber bestätigt (knapp)

- **Präzedenz-Test echt, nicht tautologisch** (Test-Linse): `TestTechPrecedenceDeclarationOrder`
  bricht die Symmetrie über *verschiedene* Ziel-Adapter — einzige Variable ist die Slice-Reihenfolge,
  beweist damit Erst-Treffer in Deklarationsreihenfolge.
- **Vertrag konsistent** (Vertrag/Spec-Linse): Versions-Bumps (Lastenheft/Spec/CHANGELOG/ADR-Index),
  Präzedenz-Aussage (Lastenheft ↔ Spec ↔ ADR ↔ Code) und RE2/unverankert/Exit-2 deckungsgleich;
  keine toten Anker; ADR-Status/Bezug/`Supersedes: —` korrekt.
- **Rückwärtskompat / nil-Panik-Freiheit** (Code-Linse): Literal-`Tech` (`match == nil`) fällt in
  `matches` auf Substring zurück — kein Panic; Substring-Verhalten unverändert.
- **Determinismus/Hermetik** (Code-Linse): RE2 via `regexp.Compile` (kein `MustCompile`-Panic-Pfad),
  `MatchString` unverankert, kein globaler State.
- **Präzedenz-Matrix eingehalten** (Konvention-Linse): `spec/*` referenzieren **kein** ADR-0015
  (downstream) — die zunächst gesetzten Spec→ADR-Links wurden korrekt entfernt.

## Delta-Re-Review

Nach den Fixes: `make gates` erneut grün — Dogfooding `0`, d-check `0`, alle Test-Pakete `ok`
inkl. der neuen Kern-/Config-/CLI-Tests. Die Fixes sind additiv (ein Guard + Tests + Doku/Prozess),
kein Eingriff in die bereits verifizierte Match-/Präzedenz-Logik.
