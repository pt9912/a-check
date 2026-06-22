# ADR-0008 — Ports dürfen Domänen-/Kern-Typen referenzieren

- **Status:** Accepted
- **Datum:** 2026-06-22
- **Autor:** pt9912
- **Bezug:** [AC-FA-RULE-004](../../../spec/lastenheft.md#ac-fa-rule-004--port-disziplin-regel-port-impurity) (Port-Disziplin, Change-Request 0.1.0→0.2.0), [AC-FA-RULE-005](../../../spec/lastenheft.md#ac-fa-rule-005--schicht-richtung-regel-wrong-direction) (Schicht-Richtung, das Edge-Modell), [AC-FA-CONF-001](../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml) (`edges`/`allow`), §1 Zweck (Vier-Repo-Konsolidierung)
- **Schärft:** [SPEC-RULE-001](../../../spec/spezifikation.md#spec-rule-001--regel-auswertung) — macht die `port-impurity`-Auswertung verbindlich: Befund nur bei Adapter-/Tech-Import (oder verbotenem Konstrukt), **nicht** bei Kern-Import; `ports → core` ist edge-regiert.
- **Supersedes:** — (löst die Co-Location-Begründung in [ARC-002](../../../spec/architecture.md#2-komponenten) ab; die Architektur-Sicht wird nachgezogen, ist aber kein ADR-Stratum)

## Kontext

`a-check` konsolidiert die Hexagon-Regeln von vier Schwester-Repos
([Lastenheft §1](../../../spec/lastenheft.md#1-zweck-und-geltungsbereich)).
Bei der Strukturarbeit an a-checks eigener Schichtung trat ein Befund zutage:
**drei der vier Repos haben Ports, die Domänen-/Kern-Typen referenzieren** —
nur `d-check` hat reine Ports (eigene DTOs, importieren nichts):

| Repo | Port referenziert Domäne? | Beleg |
|---|---|---|
| **b-cad** (C++) | ja | `src/hexagon/ports/driven/model_importer_port.h` → `#include "hexagon/model/building.h"` |
| **d-migrate** (Kotlin) | ja | `…/ports/…/PreGenerationValidator.kt` → `import dev.dmigrate.core.model.SchemaDefinition` |
| **a-check** (Go) | ja | `ConfigPort`/`ExtractionPort`/`ReportPort` referenzieren `Model`/`Finding` |
| **d-check** (Go) | nein | `internal/hexagon/port/driven/*.go` definieren eigene DTOs, importieren nichts |

[`AC-FA-RULE-004`](../../../spec/lastenheft.md#ac-fa-rule-004--port-disziplin-regel-port-impurity)
(0.1.0) verlangte „Ports importieren weder Adapter **noch Kern**", und die
Implementierung (`internal/core/rules.go`) verdrahtete das Kern-Verbot **hart**:
`port-impurity` feuerte bei einem Import, der auf die `core`-Schicht auflöst, und
fragte `edges`/`allow` gar nicht ab. Das widersprach (a) dem Beispiel in
[`SPEC-CONF-001`](../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema),
das `{from: ports, to: core}` als **erlaubte** Kante führt, und (b) dem Zweck:
a-check würde auf b-cad, d-migrate und sich selbst **falsche** `port-impurity`-
Befunde werfen — auf genau den Repos, die es bedienen soll.

Das maßgebliche Hexagon-Prinzip (Swap-Test): Ein Port drückt die **Sprache des
Kerns** aus und darf Domänentypen referenzieren; er darf aber keine Adapter-,
Framework-, Persistence-, Messaging- oder Vendor-Typen kennen. *Ließe sich der
Adapter komplett austauschen, ohne Port und Domäne zu ändern? Wenn nein, leakt
der Port Infrastruktur.*

## Optionen

1. **Streng bleiben** (Ports rein, importieren nichts) und a-checks Ports im
   Kern-Paket co-lokieren. *Verworfen:* widerspricht §1 (b-cad/d-migrate sind so
   nicht abbildbar) und zwingt a-check, eine echte Schicht im Kern zu verstecken.
2. **`ports → core` implizit immer erlauben** (Sonderfall außerhalb des
   Edge-Modells). *Verworfen:* bricht a-checks Prinzip „alle Schicht-Richtungen
   werden explizit deklariert" und verschleiert Zyklen.
3. **`ports → core` über das Edge-Modell regieren (gewählt).** `port-impurity`
   verliert das Kern-Verbot und gewinnt das Tech-Verbot (Symmetrie zu
   `core-impurity`); `ports → core` ist erlaubt, **wenn** die Kante deklariert
   ist, sonst `wrong-direction`.

## Entscheidung

1. **`port-impurity` = Port importiert einen Adapter ODER ein Tech-/Framework-
   Symbol ODER enthält ein verbotenes Konstrukt — nicht den Kern.**
2. **`ports → core` ist edge-regiert** ([AC-FA-RULE-005](../../../spec/lastenheft.md#ac-fa-rule-005--schicht-richtung-regel-wrong-direction)):
   deklarierte `{from: ports, to: core}`-Kante ⇒ erlaubt; ohne Deklaration ⇒
   `wrong-direction`.
3. **Symmetrie Kern ↔ Ports:** beide sind die Domänenseite — keine Adapter, kein
   Tech; Domänentypen dürfen sie referenzieren. Im Code:
   `f.Layer == "ports" && (tl == "adapters" || isTech)` (parallel zu
   `core-impurity`).
4. **`ports → adapters` bleibt unbedingt verboten** (Inversionsprinzip, keine
   Edge-Ausnahme) — wie `core → adapters`.

## Konsequenzen

- a-check kann eine **echte** `ports`-Schicht (`internal/hexagon/port/`) führen,
  statt die Port-Interfaces im Kern-Paket zu verstecken; die Co-Location-
  Begründung in [ARC-002](../../../spec/architecture.md#2-komponenten) wird
  abgelöst (Sicht-Stratum nachgezogen).
- **d-check bleibt grün:** reine Ports deklarieren keine `ports → core`-Kante und
  importieren nichts — keiner der geänderten Pfade greift.
- **b-cad/d-migrate** werden mit einer `{from: ports, to: core}`-Kante korrekt
  prüfbar statt fälschlich rot.
- **Verhaltensänderung:** ein Port, der ein Tech-Symbol importiert, meldet jetzt
  `port-impurity` (vorher `tech-leak`) — präziser für die Port-Grenze; die
  Erst-Treffer-Reihenfolge in [SPEC-RULE-001](../../../spec/spezifikation.md#spec-rule-001--regel-auswertung)
  (`port-impurity` vor `tech-leak`) trägt das.
- [Lastenheft](../../../spec/lastenheft.md) 0.1.0 → **0.2.0**
  ([`AC-FA-RULE-004`](../../../spec/lastenheft.md#ac-fa-rule-004--port-disziplin-regel-port-impurity)-Change-Request).

## Fitness Function

- `make arch-check` (Dogfooding, [AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)):
  a-check prüft seine eigene echte `ports`-Schicht — `ports → core` via
  deklarierter Kante grün, `ports → adapter` rot.
- `make test`: `rules_test.go` deckt `ports → core` (clean, Edge vorhanden),
  `ports → adapter` (`port-impurity`) und `ports → tech` (`port-impurity`).

## Re-Evaluierungs-Trigger

- **Generalisierung der Regel-Engine** (Reinheit pro Layer deklarierbar statt an
  die Namen `core`/`ports`/`adapters` gebunden; feine `domain`/`application`-
  Trennung, `driving`/`driven`-Ports mit voller Durchsetzung) → Folge-ADR.
- Ein Konsumenten-Repo, dessen Ports legitim **mehr** als die Domäne brauchen →
  Port-Disziplin erneut prüfen.

## Geschichte

| Datum | Ereignis |
|---|---|
| 2026-06-22 | Proposed — Vier-Repo-Evidenz (b-cad/d-migrate/a-check unrein, d-check rein) + Swap-Test; behebt die Spec-/Impl-Divergenz zu [`AC-FA-RULE-004`](../../../spec/lastenheft.md#ac-fa-rule-004--port-disziplin-regel-port-impurity) und macht eine echte `ports`-Schicht für a-check möglich. |
| 2026-06-22 | Proposed → Accepted (Sign-off Auftraggeber; Multi-Linsen-Review bestanden, R1–R3 geschlossen). Ab jetzt immutable; Ablösung nur via Folge-ADR mit `Supersedes`. |
