# Review-Synthese — slice-026 Stufe 1: Mehr-Wurzel-Phantom-Guard

**Datum:** 2026-07-04 · **Gegenstand:** [slice-026](../plan/planning/in-progress/slice-026-kmp-mehr-root-phantom.md)
([ADR-0020](../plan/adr/0020-mehr-wurzel-phantom-guard.md),
[AC-FA-CONF-001](../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml) 0.16.0)
· **Form:** adversarisches Multi-Linsen-Review, 4 read-only Linsen (Code-Korrektheit ·
Vertrag/Spec-Konsistenz · Test-Abdeckung · Regelwerk/Konvention) + Delta-Re-Review des Fix-Deltas.

## Ablauf

1. **Umsetzung** (spec-first): Lastenheft/Spez 0.16.0 + ADR-0020 Proposed (`63799a7`),
   Code + Tests (`17ef852`), verifiziert gegen das KMP-Repro.
2. **R1 — 4 Linsen parallel, read-only** über `main...HEAD`.
3. **Fixes** (dieser Commit) + **Delta-Re-Review** des Fix-Deltas (unten).

## R1-Befunde und Auflösung

| # | Linse | Severity | Befund | Auflösung |
|---|---|---|---|---|
| C-1 | Code-Korrektheit | **BLOCKER** | Das Enthaltungs-Prädikat `layerContainsRoot` bildete die reale Auflösung nicht ab: es verzeichnete **jede** enthaltende Schicht (nicht die dominante) und nutzte `segIndex == 0` statt `>= 0`. Folge zweifach: **Falsch-Positiv** bei verschachtelten/überlappenden Schichten (beide Roots lösen eindeutig in die tiefere Schicht auf, der Guard meldete dennoch Konflikt — Meldung faktisch falsch), **Falsch-Negativ** bei Layer-Präfix als **innerem** Segment des Roots (`build/gen/core` unter `core/**`), das `targetLayer` matcht, der Guard aber verfehlte. | Prädikat auf die **erzwungene Schicht** umgestellt = längster passender Glob-Präfix am Root via `segIndex >= 0`, erste Schicht bei Gleichstand — **exakt `targetLayer`**. `PhantomRootConflictIn` vergleicht die pro-Root erzwungenen Schichten; Konflikt nur bei zwei verschiedenen, nicht-leeren. `rootForcedLayer` ersetzt `layerContainsRoot`. Lastenheft/[SPEC-CONF-001](../../spec/spezifikation.md#spec-conf-001--konfigurationsschema)/ADR-0020 (Entscheidung + Geschichte) nachgezogen. |
| C-2 | Code-Korrektheit | NIEDRIG | Asymmetrisches Phantom: nur **ein** Root erzwingt eine Schicht (Catch-all-Layer über einem Teilbaum), der andere keine — bleibt ungefangen. | Bewusst außerhalb Stufe 1: als **Residual** in ADR-0020 §Abgrenzung + [AC-QA-02](../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)-Grenze dokumentiert; die datei-mengen-bewusste Auflösung (Weg B, slice-027) schließt es mit. Kein Code-Delta. |
| T-1 | Test-Abdeckung | **HOCH** | Positiv-Test `…DeepGlobsValid` prüfte nur „lädt", nicht die von der Fitness Function geforderte Exit-1-Hälfte (Rezept **findet** die Verletzung). | `TestPhantomDeepGlobsResolvesCorrectly` (core): paket-tiefe Globs → `core → adapter` = `core-impurity` über `Evaluate`. |
| T-2 | Test-Abdeckung | **HOCH/MITTEL** | Kein Regressions-Anker für den Falsch-Negativ-**Mechanismus** (dass ohne Guard das Phantom entstünde). | `TestPhantomFlatGlobsMisresolves` pinnt `targetLayer` → falsch `core` bei flachen Globs; bricht sichtbar, falls Stufe 2 die Auflösung ändert. |
| T-3 | Test-Abdeckung | MITTEL | Cross-Sprachen-Determinismus (`sort.Strings`) ungetestet — ein Refactor, der die Sortierung entfernt, bliebe grün, würde produktiv flaky. | `TestResolutionMultiRootConflictLanguageDeterministic`: zwei konfligierende Sprachen → `go` (sortiert-erste) gemeldet. |
| T-4 | Test-Abdeckung | MITTEL | Zeugen-Determinismus bei mehreren Konflikten ungeübt (alle Fälle hatten genau ein Paar). | Core-Fall `three-roots-earliest-witness` (3 Roots/3 Schichten) prüft das erste (i,j)-Paar. |
| T-5 | Test-Abdeckung | NIEDRIG/MITTEL | Meldung verifizierte nur Roots, nicht Schicht-Namen + Rezept-Hinweis. | Assertion um `core`/`adapters` + `paket-spezifische Globs` erweitert. |
| C/T (Befund 1) | Code + Test | — | Die C-1-Korrektur braucht eigene Anker. | Core-Fälle `nested-layers-no-conflict` (Falsch-Positiv-Anker) + `interior-segment-conflict` (Falsch-Negativ-Anker) ergänzt. |
| V-1 / R-2 | Vertrag / Regelwerk | NIEDRIG | ADR-0020 `[slice-017]`/`[slice-023]` als Klammer-Text ohne Link-Ziel (rendern literal). | Als echte Links auf die done/-Slices gesetzt. |
| R-1 / M1 | Regelwerk | **MITTEL** | Lifecycle: Slice blieb in `open/` mit Status „noch kein Code" (nach Umsetzung faktisch falsch). | Reiner `git mv` open → in-progress (`14b1b80`), Status auf „in-progress, Stufe 1 umgesetzt"; interner Roadmap-Link nachgezogen. |
| L1 | Regelwerk | NIEDRIG | Neue AC war eine Negative-Zeile mit eingebettetem Boundary; die Prädikat-Kante (Root == Glob-Präfix) fehlte als Lastenheft-Kriterium. | Eigene **Boundary**-Zeile (paket-tiefe/verschachtelte Globs laden + finden; Root==Präfix-Kante) in AC-FA-CONF-001 ergänzt. |

## Grün bestätigt (keine Änderung nötig)

- **Vertrag/Spec-Konsistenz (Linse 2):** Prädikat Spec↔Code deckungsgleich (nach C-1-Fix
  erneut); §3.4 gewahrt (kein ADR/Slice im Spec-Body, nur Historie); ADR-0020 korrekt
  Proposed; alle Anker/Links verifiziert (`make doc-check` real: 0 Befunde); Versions-/
  Historie-Zeilen konsistent 0.16.0.
- **Regelwerk (Linse 4):** Schärfung nur im Lastenheft (nicht per ADR); Kern-Reinheit
  gewahrt (`core` nur Stdlib; `rootForcedLayer`/`PhantomRootConflictIn` nutzen nur
  `globPrefix`/`segIndex`); Schichtung sinnvoll (Prädikat = Auflösungs-Fakt → Kern, Policy
  + dt. Meldung → Config-Adapter); ID-Vergabe/Index/Traceability korrekt.
- **Determinismus:** `resolveAndCheck` sortiert Sprachen; `rootForcedLayer` in Root-/
  Layer-Deklarationsreihenfolge → Zeuge stabil (jetzt zusätzlich per Test gepinnt).

## Delta-Re-Review (Fix-Delta)

Eine Linse (Code-Korrektheit) über das C-1-Fix-Delta, read-only. **Ergebnis: Fix sauber**
— `rootForcedLayer` spiegelt `targetLayer` Zeile für Zeile (gleiches `globPrefix`/`p != ""`/
`segIndex >= 0`/`len(p) > bestLen`), die Konflikt-Schleife ist deterministisch (kein
Selbstpaar, keine leere Schicht, Zeuge = erstes (i,j)-Paar), und die neuen Tests **brechen
unter dem alten Code** (`interior-segment-conflict` und `nested-layers-no-conflict` sind
echte Anker, nicht tautologisch). Code ↔ ADR-0020 ↔ SPEC-CONF-001 ↔ AC-FA-CONF-001
wortgleich.

| # | Severity | Beobachtung | Auflösung |
|---|---|---|---|
| B1 | NIEDRIG | `nested-layers-no-conflict` unterscheidet „längster Präfix" nicht von „erst-deklariert" (beide Roots fallen ohnehin in dieselbe Schicht) — die Längster-Präfix-Wahl von `rootForcedLayer` bei mehreren matchenden Schichten war nur indirekt (über geteilten Code + `targetLayer`-Tests) gepinnt. | Anker `forced-layer-longest-prefix` ergänzt: ein Root matcht `broad` (zuerst deklariert) **und** `appL` (danach); der Zeuge muss `appL` (längster Präfix) nennen, nicht `broad`. |
| B2 | NIEDRIG | Konservatives Falsch-Positiv (nicht vom Fix eingeführt): eine Schicht mit *flachem* **und** *tiefem* Glob lässt die flachen Globs den Konflikt erzwingen, obwohl die tiefen korrekt auflösten. | Als Re-Eval-Trigger in ADR-0020 dokumentiert (fail-closed-konform, degenerierte/redundante Form, bei keiner realen Config). |

Verbliebene, korrekt dokumentierte Grenze: die **asymmetrische** Phantom-Form (nur ein Root
erzwingt eine Schicht) — Residual für Stufe 2 ([AC-QA-02](../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze), ADR-0020 §Abgrenzung).

## Prozess-Nachtrag (MR-006 / d-check-Matrix)

ADR-0020 zitierte in der Optionen-Tabelle zunächst `slice-017`/`slice-023` als
Entscheidungs-**Beleg** — das MR-006-Anti-Muster (ADRs argumentieren aus Spec/Verhalten,
nicht aus Planungs-Artefakten; d-check erzwingt das maschinell via `matrix.adr → slice` +
Token-Erkennung). a-checks Matrix-Modul kennt diese Regel noch nicht (ältere Variante);
der Verweis wurde dennoch auf das **Verhalten** ([AC-FA-CONF-001](../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml):
unbekannter `languages`-Schlüssel / leerer `tech.adapter` → Exit 2) umgestellt. Die
Übernahme der strengeren d-check-Matrix (intra-Spec `order`/`direction`, `adr → slice`,
Provenance-Marker, Grandfathering) ist als eigene Harness-Konvergenz vorgemerkt.
