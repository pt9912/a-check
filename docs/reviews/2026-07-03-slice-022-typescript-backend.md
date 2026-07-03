# Review-Synthese — slice-022: TypeScript-Backend + `relative`-Auflösungs-Modus

**Datum:** 2026-07-03 · **Gegenstand:** [slice-022](../plan/planning/done/slice-022-typescript-backend.md)
([ADR-0017](../plan/adr/0017-relative-resolution-modus.md),
[AC-FA-EXTRACT-001](../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion)/[AC-FA-CONF-001](../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml) 0.13.0)
· **Form:** adversarisches Multi-Linsen-Review, 4 read-only Linsen + Delta-Re-Review des Fix-Deltas.

## Ablauf

1. **Entwurfs-Review (Maintainer, vor Abnahme):** 1 BLOCKER (Mehrzeilen-Imports → Entscheid G
   Fortsetzungs-Regex), 6 MAJOR, 8 MINOR, 5 NIT — vollständig in den Entwurf eingearbeitet
   (Commit `d3d285c`), danach Abnahme A–G.
2. **Implementierungs-Review R1 (4 Linsen, parallel, read-only)** über `main...HEAD`
   nach Commit `721cc31`.
3. **Fixes** (Commit `926d38d`) + **Delta-Re-Review** über `721cc31..926d38d`.
4. Rest-Befunde des Delta-Re-Reviews (1 MINOR, 2 NIT) im Closure-Commit behoben.

## R1-Befunde und Auflösung

| # | Linse | Severity | Befund | Auflösung |
|---|---|---|---|---|
| C-1 | Code-Korrektheit | **MAJOR** (BLOCKER-Tendenz) | `lateral`/`adapterSeg`/`adapter_sink` arbeiteten am **Roh-Specifier** — jeder gleich-schichtige relative Import (`./helper` in `src/adapters/http/`) wurde falsch als `lateral-adapter` gemeldet (`adapterSeg("./helper") == ""`); Alltagsfall jedes TS-Adapters | `targetLayer` liefert zusätzlich den **gematchten Kandidaten**; `lateral` prüft Sink + Sub-Einheit auf dem Kandidaten (path-Modus byte-identisch, gepunktetes fixed-root wird mit-korrigiert). [SPEC-RULE-001](../../spec/spezifikation.md#spec-rule-001--regel-auswertung)-lateral-Zeile präzisiert; Tests `TestRelativeIntraSubunitNoLateral`, `TestRelativeAdapterSinkOnCandidate`; Handbuch-Hinweis „`adapter_sink` unter `resolution` als Pfad-Fragment" |
| T-1 | Test-Abdeckung | **MAJOR** | Mittelteil-Klasse von `tsFrom` unbewacht: die `knex.from(…)`-Negative wurden vom `from(`-Anker geblockt, nicht von der Klasse — der Mutant „Klasse um `=`/`.` erweitern" überlebte die Suite; Fixture-Kommentare überklagten | Mutantentötende Negative `export const x = from './adapters/db5';` und `export obj.prop from './adapters/db7';`; CLI-Fixture-Zeile 3 auf `export const q = from '../adapters/db2';` umgestellt (pinnt die Klasse end-to-end); Kommentare richtiggestellt |
| S-1/S-2 | Spec-Konsistenz | **MAJOR** ×2 | „reservierter `relative`-Modus" stale in Lastenheft-Out-of-Scope und [SPEC-EXTRACT-001](../../spec/spezifikation.md#spec-extract-001--import-extraktion)-Python-Bullet — Selbstwiderspruch zur gültigen `mode`-Menge | Beide Stellen auf „dokumentierte Grenze der Python-Extraktion, unabhängig vom gültigen `relative`-Modus" umformuliert; ebenso der `extract.go`-Kommentar (S-3 MINOR) und — als Delta-Rest — der `extract_test.go`-Kommentar |
| C-2 | Code-Korrektheit | MINOR | TS-Specifier mit `//` (URL-Importe) fallen dem C-Kommentar-Strip zu — stilles Falsch-Grün, nicht dokumentiert | Als [AC-QA-02](../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)-Grenze in Lastenheft-Out-of-Scope + SPEC-EXTRACT-001 ausgewiesen (URL-ESM/Deno nicht unterstützt); kein Code-Delta (String-bewusstes Stripping = eigenes, sprachübergreifendes Inkrement) |
| T-2 | Test-Abdeckung | MINOR | AC-Klausel „`tech` greift am Roh-Symbol" im relative-Modus ungedeckt | `TestRelativeTechAtRawSymbol` (Bare-Import `@nestjs/core` → `tech-leak` trotz leerer Kandidatenmenge) |
| T-3 | Test-Abdeckung | MINOR | `.js`-Specifier → Schicht nicht end-to-end gepinnt | `targetLayer`-Assertion (`../adapters/db.js` → `adapters`) ergänzt |
| T-4 | Test-Abdeckung | MINOR | Escape-Ast `cand == ".."` untested | `resolveImport("..", "main.ts")` → leer, ergänzt |
| S-4 | Spec-Konsistenz | MINOR | Roadmap führte ADR-0017 als „(Proposed) in Abnahme" trotz Accepted | Auf „(Accepted) in Umsetzung" korrigiert (Regelwerk-Linse hatte die nachlaufende Roadmap als Bestands-Praxis eingeordnet; das faktisch falsche Status-Label wurde dennoch gefixt) |
| R-1 | Regelwerk | MINOR | slice-022-Dokument führte ADR-0017 zugleich als Accepted (Kopf) und Proposed (Bezug/Hinweis-Box) | Bezug + Hinweis-Box auf den erteilten Sign-off umgestellt |
| C-3/C-4 | Code-Korrektheit | MINOR/NIT | Kompakte Form `import{A}from'./b'` und nacktes `from '…'` auf eigener Zeile = Falsch-Negative | Als dokumentierte Grenzen ausgewiesen (Formatter-Konvention; das Pflicht-Whitespace schützt den Keyword-Präfix-Ausschluss) |
| T-5/T-6 | Test-Abdeckung | NIT | `exportX`-Präfix (versprochen) und Backtick-Negativ fehlten | Beide Negativ-Zeilen ergänzt |
| R-2 | Regelwerk | NIT | Struct-Alignment `extract.go` um eine Spalte unter-gepolstert (kein Gate erzwingt gofmt — Lauf-Status-Behauptung dadurch nicht falsifiziert) | Ausgerichtet |
| S-6 | Spec-Konsistenz | NIT | `python: {mode: relative}` = stiller No-Op (Extraktion liefert nie relative Specifier) undokumentiert | Als Notiz im Slice §7 ausgewiesen (Modus sprach-parametrisch generisch; Python-Grenze bleibt Extraktions-Grenze) |

## Delta-Re-Review (über `721cc31..926d38d`)

**Ergebnis: kein BLOCKER, kein MAJOR.** Verifiziert: path-Modus nach dem lateral-Fix
byte-identisch (Kandidat == Roh-Import; `lateral` nur bei aufgelöstem Ziel erreichbar);
kein Bestands-Test/Vertrag pinnt das alte fixed-root-Roh-Symbol-Verhalten (die Korrektur
dort heilt ein latentes Dauer-Falsch-Positiv); Kandidaten-Auswahl deterministisch
(roots-Reihenfolge, strikte Längen-Ungleichung); alle neuen Mutanten-Negative simuliert
und als tödlich bestätigt (inkl. `TestRelativeAdapterSinkOnCandidate`: Sink-Fragment
`adapters/driver-common` steckt bewusst nicht im Roh-Specifier); „reserviert"-Audit über
`spec/` + `internal/` — alle verbliebenen Fundstellen betreffen korrekt `namespace` oder
Punkt-in-Zeit-Historie. Rest-Befunde (MINOR: stale Kommentar `extract_test.go`;
NIT: `lateral`-Kommentar ohne fixed-root-Hinweis, Kandidaten-Tie-Break undokumentiert;
NIT: Handbuch-Migrationshinweis `adapter_sink` als Pfad-Fragment) — alle im
Closure-Commit behoben.

## Gate-Beleg (nach allen Fixes)

`make gates` grün — `lint` 0 issues, alle Test-Pakete `ok`, `coverage-gate` 96,00 %
(≥ 90 %), `arch-check` 0 (Dogfooding), `doc-check` 0 Befunde, Meta-Gates ok;
`make ci` (inkl. `image-test` 4/4) grün.
