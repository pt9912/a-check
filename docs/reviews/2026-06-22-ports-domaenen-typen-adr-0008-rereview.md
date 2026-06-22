# Re-Review (Delta) — Ports-Fixes nach Multi-Linsen-Review ([ADR-0008](../plan/adr/0008-ports-duerfen-domaenen-typen-referenzieren.md))

**Datum:** 2026-06-22

**Gegenstand:** der Fix-Delta R1–R8 nach dem [Erst-Review](2026-06-22-ports-domaenen-typen-adr-0008.md). Die Fixes waren autor-eingebaut und nur `make gates`-verifiziert, nicht unabhängig reviewt — dieser Delta-Lauf schließt die Lücke. Stand: Branch `feat/ports-domaenen-typen-adr-0008`.

**Methode:** zwei adversarische Linsen (read-only) — Code/Test-Korrektheit (R1) und Spec-/Doku-Konsistenz (R2/R3/R4/R8).

**Verdict:** Delta **sauber**, keine Blocker, keine neuen Widersprüche. Ein konvergenter MINOR → als **R9** ergänzt.

## Befunde

| Fix | Verdict | Beleg |
|---|---|---|
| **R1** Regressionstest | ✅ bestätigt | `TestPortToCoreWithoutEdge` erreicht nachweislich den `wrong-direction`-Zweig (`ruleFor`-Trace); differenzielles Paar mit `TestPortDomainAllowed`; beide Assertions komplementär + nicht-maskierend; keine `testModel`-Kontamination (Factory + lokale Kopie). |
| **R2** Doku-Regel | ✅ bestätigt | Kein Rest von „importiert Kern" / „reine Abstraktionen" (grep-verifiziert); Handbuch 1.4 + Historie akkurat. |
| **R3** `--print-config` | ✅ bestätigt | Gerüst **empirisch** strict-decodebar (durch a-check gelaufen: Exit 0; `bogus` → Exit 2); `forbidden_constructs.ports` jetzt sinnvoll (echter `ports`-Layer). |
| **R4** Architektur | ✅ bestätigt | §3-Satz korrekt; §2-Mermaid-Pfeile gedreht + `adapter→core` ergänzt; alle Knoten definiert, syntaktisch gültig. |
| **R8** Beispiele | ✅ bestätigt | `ports`-Schicht + beide Kanten, konsistent zu Gerüst + Spec-Beispiel. |

## R9 (aus diesem Re-Review)

Konvergent von beiden Linsen: alle Beispiel-Configs (`--print-config`, README, Handbuch, Spezifikation) zeigen das *reine* Hexagon **ohne** `adapters → core`-Kante, während [`architecture.md`](../../spec/architecture.md) §3 und a-checks eigene `.a-check.yml` diese Kante real führen (a-checks Adapter referenzieren Domänentypen direkt). Wer das Gerüst kopiert **und** in Adaptern Domänentypen referenziert, bekommt `wrong-direction`.

→ Technisch korrekt (durch `core ← ports ← adapters`, [AC-FA-RULE-005](../../spec/lastenheft.md#ac-fa-rule-005--schicht-richtung-regel-wrong-direction), gedeckt), aber didaktisch unklar. **R9:** im `--print-config`-Gerüst sowie in den README-/Handbuch-Beispielen eine auskommentierte Optionskante `# - {from: adapters, to: core}` ergänzt — schließt die Lücke, ohne das reine Hexagon als Default aufzugeben.

## Gesamt

Delta freigegeben: **R1–R4/R7/R8 bestätigt geschlossen, R9 ergänzt.** R5 bleibt offen (RULE-005-Nuance), R6 → welle-10 (`adapterSeg`-Generalisierung). [ADR-0008](../plan/adr/0008-ports-duerfen-domaenen-typen-referenzieren.md) `Accepted`.
