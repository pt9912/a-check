# Review — slice-009 Rollen-Dispatch ([AC-FA-RULE-006](../../spec/lastenheft.md#ac-fa-rule-006--schicht-rollen-generische-regel-anwendung), welle-10a)

**Datum:** 2026-06-22

**Gegenstand:** Umsetzung des rollen-basierten Regel-Dispatch — Lastenheft 0.3.0 (`AC-FA-RULE-006`) + [ADR-0009](../plan/adr/0009-rollen-basierter-regel-dispatch.md) + Spezifikation + Engine (`config.go`/`model.go`/`rules.go`) + Tests. Branch `feat/slice-009-rollen-dispatch`.

**Methode:** vier perspektiven-diverse adversarische Linsen (read-only) — Code-Korrektheit · Vertrag/Spec-Konsistenz · Test-Abdeckung · Regelwerk/Konvention.

**Gesamtbewertung:** **Kein Blocker.** Code-, Vertrag/Spec- und Regelwerk-Linse bestätigen die Engine korrekt und widerspruchsfrei. Die Test-Linse deckte reale Härtungs-Lücken auf — alle vor Merge geschlossen.

## Befunde & Resolution

| ID | Linse | Schwere | Befund | Resolution |
|---|---|---|---|---|
| **T1** | Test | HOCH | Boundary-AC (klassische Namen `core`/`ports`/`adapters` = Verhalten 0.2.0) nur implizit über Alttests | ✅ `TestInferenceBoundaryClassicNames` (core→adapter, ports-Konstrukt, intra-`adapters`-lateral via Inferenz). |
| **T2** | Test | HOCH | `TestRoleCrossLayerLateralCategorical` nicht differenzial (`hasRule` statt exakt; nur `allow`, nicht `edges`) | ✅ table-driven `allow`+`edge`, `len(fs)==1 && lateral-adapter` — würde bei edge-regiertem `lateral` rot. |
| **T3** | Test | MITTEL | `decodeLayer` Default- (Scalar) + Decode-Fehler-Zweige ungetestet | ✅ `TestLayerScalarFailsClosed`, `…SeqBadElement…`, `…ObjectBadGlobsType…`. |
| **T4** | Code | MAJOR *(10b)* | `adapterSeg==""`-Intra-Falsch-Negativ für fremd benannte Einzel-`adapter`-Layer (das dokumentierte R6-Loch) | ✅ Regression-Pin `TestForeignAdapterIntraNoLateral10a` + `lateral`-Doc-Kommentar geschärft. Code-Fix bleibt späteres Inkrement. |
| **K1** | Regelwerk | NIEDRIG | `(welle-10b)` im **normativen** Out-of-Scope von `AC-FA-RULE-006` | ✅ → „späteres Inkrement". |
| — | Code | MINOR | YAML-Merge-Key `<<` im Layer-Objekt fail-closed abgelehnt (sicher, aber inkonsistent zur Top-Level-Ebene) | bewusst offen — dokumentierte fail-closed-Grenze. |

## Pro Linse (Kurzfazit)

- **Code:** Rollen-Dispatch + Mapping, kategorisches `lateral` (kein `edgeAllowed`/`allow`-Check, verifiziert), Erst-Treffer-Reihenfolge, `role:` > Inferenz, Konstrukt-`port-impurity` rollen-basiert, yaml-Gotcha-Handling — alle bestätigt. Einziger Substanzpunkt: das auf ein späteres Inkrement verschobene `adapterSeg`-Loch (T4).
- **Vertrag/Spec:** Lastenheft/ADR-0009/Spezifikation/Code durchgängig konsistent — kein Widerspruch; „kategorisch", `role:` > Inferenz, Reihenfolge, Befund-Namen, Rückwärtskompat über alle Strata stimmig.
- **Test:** nach T1–T4 solide; größtes Restrisiko (`hasRule` statt exakter Befundmenge) für den Categorical-Kern geschlossen.
- **Regelwerk:** §3.4/§3.5/§3.6 sauber, `ADR-0009` formgerecht (Schärft aufwärts auf SPEC-*), Anforderungs-Prozess vollständig, Traceability/Gates real.

## Offen / Folge (welle-10b)

`app`-Rolle (Domain/Application-Trennung), `driving`/`driven`-Ports, Namens-Generalisierung von `adapterSeg` (R6/T4), `targetLayer` längster-Präfix-Match; YAML-Merge-Key in Layer-Objekten.

## ADR-0009-Abnahme

[ADR-0009](../plan/adr/0009-rollen-basierter-regel-dispatch.md) ist inhaltlich tragfähig und durch das Dogfooding (`make arch-check` grün, Rückwärtskompat via Inferenz) belegt; **Accepted** nach Schließen von T1–T4/K1 (erledigt).
