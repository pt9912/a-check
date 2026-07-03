# Review-Synthese — slice-023: d-check-Pilot-Deltas

**Datum:** 2026-07-03 · **Gegenstand:** [slice-023](../plan/planning/open/slice-023-dcheck-pilot-deltas.md)
([ADR-0018](../plan/adr/0018-exclude-scan-scope.md),
[AC-FA-RULE-003](../../spec/lastenheft.md#ac-fa-rule-003--tech-kapselung-regel-tech-leak)/[AC-FA-CONF-001](../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml)
0.14.0: `tech.adapter`-Liste · `composition_root: allow|forbid` · `exclude`)
· **Form:** adversarisches Multi-Linsen-Review, 4 read-only Linsen + Delta-Re-Review des Fix-Deltas.

## Ablauf

1. CR 0.14.0 spec-first vom Maintainer (Lastenheft, auf `main`); Umsetzung auf dem Branch in der
   Reihenfolge ADR-0018 (Proposed) → Spezifikation 0.14.0 → Code → Tests → Handbuch 1.19
   (Commit `17115c6`).
2. **R1: 4 Linsen parallel** über `main...HEAD`.
3. **Fixes** (Commit `831d66d`) + **Delta-Re-Review**; Delta-Reste im Closure-Commit.

## R1-Befunde und Auflösung

| # | Linse | Severity | Befund | Auflösung |
|---|---|---|---|---|
| C-B1 / S-M1 | Code + Spec (deckungsgleich) | **BLOCKER** | Leerer/fehlender Skalar-`adapter` **invertierte** das Verhalten: Alt-Semantik war ein stiller **Never-Leak** (`strings.Contains(pfad, "")` immer wahr → das Muster meldete nie — toter Eintrag, falsch-grün); die Erst-Umsetzung kippte via `contains`-Helper auf Always-Leak; der Code-Kommentar beschrieb die Alt-Semantik falsch | **Entscheid fail-closed** (statt Konservierung des stillen No-Ops): leerer wie fehlender `adapter` (auch `null`) → Exit 2 — Ethos-Linie [slice-017](../plan/planning/done/slice-017-unbekannte-sprache-exit2.md)/leerer `resolution`-Root; nicht-leerer Skalar byte-identisch (golden gepinnt). Lastenheft-0.14.0-CR-Zeile, [SPEC-CONF-001](../../spec/spezifikation.md#spec-conf-001--konfigurationsschema), CHANGELOG, Handbuch und Slice-§3(d) entsprechend präzisiert |
| C-M1 | Code | MINOR | YAML-Alias (`&anchor`/`*ref`) auf `adapter` wurde fälschlich abgelehnt (`yaml.Node` behält `Kind==AliasNode`) | Alias-Dereferenzierung vor dem Kind-Switch; `TestTechAdapterAliasResolved` (Alias-auf-Liste durchläuft die volle Listen-Validierung — im Delta-Review verifiziert) |
| T-4/T-l | Tests | MINOR | Geordnete Mischung aus Composition-Root- und Normal-Befunden nicht determinismus-gepinnt | `TestCompositionRootFindingsSortedWithOthers` (Append-Reihenfolge weicht bewusst von der sortierten Ordnung ab, exakte Pfad-Folge assertiert, [AC-QA-01](../../spec/lastenheft.md#ac-qa-01--determinismus)) |
| T-3 | Tests | MINOR | Schicht-Ausnahme der Composition Root nur **inzidentell** bewacht (hing an einer zufälligen `""→core`-Kante) | Fixture kategorisch geschärft: CR-Datei liegt im Domain-Layer und importiert ein Adapter-Symbol — ohne Ausnahme gäbe es kategorisches `core-impurity` |
| T-i/T-Mapping/T-k | Tests | MINOR | Fehlender/leerer Skalar, `adapter: {…}` (Mapping) und der `NewTech`-Kern-Wächter ungetestet (Coverage-Drop-Ursache) | `TestTechAdapterAbsentFailsClosed`, `…EmptyScalarFailsClosed` (inkl. `null`), `…AsMappingFailsClosed`, `TestNewTechEmptyAdapterListFails` — Coverage zurück auf 96,10 % |
| T-6 | Tests | NIT | Byte-Identität der Einzel-Adapter-Meldung nicht golden gepinnt | `TestTechLeakSingleAdapterGoldenMessage` (`Tech net/http außerhalb adapters/http`, im Delta-Review zeichengenau gegen das Alt-Format verifiziert) |
| S-N1/S-N2 | Spec | NIT | Stale YAML-Inline-Kommentare zu `composition_root` (Spez „Ausnahme für tech-leak", Handbuch „von Regeln ausgenommen") | Beide präzisiert (Schicht-Regeln + per-Eintrag-abschaltbares `tech-leak`) |
| C-N1 | Code | NIT (Grenze) | `inAdapter` ist Teilstring-, nicht segmentgrenzen-bewusst (`adapters/config` matcht `adapters/configurator`) — Alt-Verhalten, bei Listen nur breiter anwendbar | Als dokumentierte Grenze in SPEC-CONF-001 + Handbuch ausgewiesen |
| C-N2 | Code | NIT (Grenze) | Erst-Treffer-Shadowing: breites `allow`-Tech vor spezifischerem `forbid`-Tech überdeckt letzteres in der Composition Root | Konsistent zur ADR-0015-Erst-Treffer-Kette; als Beobachtung notiert (kein Delta) |
| C-N3 | Code | NIT (Perf) | `exclude` prunt keine Verzeichnis-Traversierung (Walk steigt in `node_modules` ab, filtert je Datei) | Korrektheit gewahrt; `SkipDir`-Optimierung als möglicher Folge-Punkt notiert |
| R-B4 | Regelwerk | NIT | CHANGELOG-Intro suggerierte, auch das Lastenheft ziehe erst jetzt auf 0.14.0 | Präzisiert („Spezifikation … folgt dem Lastenheft-CR 0.14.0") |
| T-L1/T-e | Tests | Notiz | „ungültiger `exclude`-Glob" ≡ leerer Glob (die Glob-Engine ist total); „nie gelesen" ist nur strukturell garantiert (Code-Position vor `os.ReadFile`, im Review verifiziert — ein chmod-Fixture scheitert an Root-Rechten der Test-Umgebung) | Beides dokumentiert belassen |

**Regelwerk-Linse:** keine Verstöße (§3.2/§3.4/§3.5/§3.6/Traceability/Linkpflicht/Index alle konform);
einzige Merge-Vorbedingung: [ADR-0018](../plan/adr/0018-exclude-scan-scope.md) Sign-off
(Proposed → Accepted).

## Delta-Re-Review (über `17115c6..831d66d`)

**Kein BLOCKER, kein MAJOR.** Alle `tech.adapter`-Eingabewege simuliert (fehlt/`null`/leer/Skalar/
Liste/Alias-auf-Skalar/Alias-auf-Liste): keiner erreicht mehr `[""]` oder leere Einträge; die
Golden-Meldung ist byte-identisch zum Alt-Format; das kategorische CR-Fixture bricht nachweislich
bei gefallener Schicht-Ausnahme (zeichengenau geprüft: `adapters/http/client` enthält `net/http`
nicht); die Sortier-Mischung ist korrekt. Zwei Reste — Lastenheft-Wortlaut (Skalar-Rückwärtskompat
ohne die neue leer/fehlend-Ausnahme) und der `NewTech`-Wächter enger als sein Kommentar — wurden im
Closure-Commit behoben (Lastenheft-Präzisierung der 0.14.0-CR-Zeile + Negativ-AC; `NewTech` weist
jetzt auch leere Einträge ab).

## Gate-Beleg (nach allen Fixes)

`make gates`/`make ci` grün — `lint` 0 issues, alle Test-Pakete `ok`, `coverage-gate` 96,10 %
(≥ 90 %), `arch-check` 0 (Dogfooding, Eigen-Config byte-identisch), `doc-check` 0 Befunde,
`image-test` 4/4.
