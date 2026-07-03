# Review-Synthese — slice-024: Root-Sub-Einheit für `adapterSeg` (Blatt-Klassifikation)

**Datum:** 2026-07-03 · **Gegenstand:** [slice-024](../plan/planning/done/slice-024-adapterseg-root-subeinheit.md)
([ADR-0019](../plan/adr/0019-adapterseg-root-subeinheit.md) Proposed — A–C abgenommen, D ausstehend;
[AC-FA-RULE-002](../../spec/lastenheft.md#ac-fa-rule-002--keine-lateralen-adapter-kanten-regel-lateral-adapter) 0.15.0)
· **Form:** adversarisches Multi-Linsen-Review, 3 read-only Linsen (Code · Vertrag/Spec · Tests).

## Besonderheit des Slices: deklarierter Umsetzungs-Fund (Entscheid D)

Die abgenommene Entwurfs-Semantik (reine Root-Regel) erwies sich im **Dogfooding** als
unvollständig: `report_test.go` (externes Go-Testpaket) importiert sein eigenes Paket —
der Kandidat endet auf dem Paket-**Verzeichnis**, die Root-Deutung erzeugte einen neuen
Falsch-Positiv und hätte Cross-Paket-Lateral in Go-Repos geblendet. Umgesetzt wurde die
**Blatt-Klassifikation** (datei-förmiges Blatt `.` → Root `''`; verzeichnis-förmiges
Blatt → ist die Sub-Einheit) als **Entscheid D zur Abnahme**; ADR-0019 blieb dafür
`Proposed` (Lerneintrag [slice-023](../plan/planning/done/slice-023-dcheck-pilot-deltas.md):
Abweichung deklariert statt still; Alt-Code-Semantik am `main`-Stand verifiziert).

## Linsen-Ergebnisse

**Kein BLOCKER, kein MAJOR — in keiner Linse.**

| # | Linse | Severity | Befund | Auflösung |
|---|---|---|---|---|
| C-M1 | Code | MINOR | Gegenrichtung der Punkt-Ambiguität undokumentiert: ein Verzeichnis-Blatt **mit** Punkt (`yaml.v2`, `config.v1`) gilt datei-förmig → Root; Lateral zwischen solchen Sub-Einheiten würde geblendet (falsch-negativ, real selten) | Als gepinnte [AC-QA-02](../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)-Grenze dokumentiert (ADR-Re-Eval-Trigger) + Test-Pin (`adapterSeg("…/yaml.v2") == ""`) |
| T-1 | Tests | LOW | Go-Paket-Blatt end-to-end nur via `make arch-check` bewacht, nicht im reinen `go test` | `TestGoPackageRootSubunitE2E` (Mini-Go-Repo: Eigen-Paket-Import befundfrei, Fremd-Paket-Import genau 1 lateral) |
| S-B1 | Spec | LOW | §3-Fence des Slice-Docs ist Vor-D-Stand | Hinweis-Box nach slice-022-Muster auf „maßgeblich 0.15.0 inkl. Entscheid D" umgestellt (Fence bleibt historischer Entwurf, Bestands-Praxis) |
| S-B2 | Spec | LOW | Plan-Text suggerierte „Sign-off ⇒ Accepted" — seit D unvollständig | §2/§4.1 auf „Accepted erst nach Abnahme inkl. D" präzisiert |
| S-B3/B4/B5 | Spec | INFO | AC-FA-RULE-006-Bezug fehlte in §2; „39/40"-Zahlen nur in ADR-Geschichte; Glossar ohne Endungslos-Halbsatz | §2-Bezug ergänzt; Geschichte präzisiert (V3b: 39, V4: 40); Glossar-Halbsatz ergänzt |
| C-M2/NIT, T-2/T-3 | Code/Tests | INFO | „irgendein Punkt ≠ echte Endung" (deckungsgleich mit C-M1); endungslose Sprach-Globs als Annahme; Symmetrie-Doktrin-Hinweis; ADR-Status vs. Test-Pins = der deklarierte D-Zustand | Keine Änderung (dokumentiert bzw. beabsichtigt) |

**Explizit bestätigt:** byte-identisches Verhalten für alle Nicht-Blatt-Fälle (inkl. leerer
Rest = Import des Layer-Verzeichnisses); Sink-/Cross-Layer-Zweige unberührt; **kein
überlebender Mutant** (8 geprüfte Mutationen, jede mit benanntem brechendem Test); der
Root↔Root-Pin ist am `main`-Stand verifiziert **rot**; kein Bestands-Pin verliert Schärfe
(alle Alt-Fixtures laufen im unveränderten Multi-Segment-Zweig); alle Straten
deckungsgleich (Lastenheft/Spez 0.15.0, ADR, Handbuch 1.22, CHANGELOG); AGENTS §3.4/§3.5
konform.

## Gate-/Pilot-Beleg (nach allen Fixes)

`make gates`/`make ci` grün — `lint` 0 issues, alle Test-Pakete `ok`, `coverage-gate`
96,50 % (≥ 90 %), `arch-check` 0 (Dogfooding — der aufdeckende Fall selbst), `doc-check`
0 Befunde, `image-test` 4/4. **Pilot-Beleg:** die verifizierte b-cad-Vollrichtungs-Config
(V4) gegen `a-check:dev` ⇒ **0 Befunde** (vorher 40 Falsch-Positive); adversarische
Gegenprobe (io → ui/view) feuert weiter.
