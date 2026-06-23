# ADR-0012 — Driving/Driven-Richtung als orthogonale Schicht-Dimension

- **Status:** Proposed
- **Datum:** 2026-06-23
- **Autor:** pt9912
- **Bezug:** [AC-FA-RULE-008](../../../spec/lastenheft.md#ac-fa-rule-008--driving-driven-port-richtung-regel-port-direction-mismatch) (neu; Lastenheft 0.5.0→0.6.0), [AC-FA-RULE-006](../../../spec/lastenheft.md#ac-fa-rule-006--schicht-rollen-generische-regel-anwendung) (Rollen-Mechanismus)
- **Schärft:** [SPEC-RULE-001](../../../spec/spezifikation.md#spec-rule-001--regel-auswertung) + [SPEC-CONF-001](../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema) — macht die optionale Richtung und die Regel `port-direction-mismatch` verbindlich.
- **Supersedes:** —

## Kontext

[ADR-0011](0011-domain-application-trennung-rolle-app.md) (welle-10b/b2a) führte die
Rolle `app` ein und ließ die `driving`/`driven`-Port-Unterscheidung als
Re-Evaluierungs-Trigger offen. Nach [AC-FA-RULE-007](../../../spec/lastenheft.md#ac-fa-rule-007--rolle-app-und-strenge-domain)
trägt eine Schicht eine Rolle ∈ {`domain`, `app`, `port`, `adapter`} — aber **alle Ports
sind gleich**. Ein **Treiber**-Adapter (CLI/HTTP), der direkt einen **driven**-Port
(Repository/Filesystem) importiert statt über die `app`-Schicht zu gehen, ist ein
Architektur-Bruch, den **keine** Regel fängt: `lateral-adapter` greift nur
Adapter→Adapter; `wrong-direction` ist kanten-geregelt und per `edges`/`allow` aufhebbar.

Die Konsumenten (b-cad/d-check/d-migrate) modellieren ihre Treiber/Getriebenen-Trennung
mit getrennten `driving`/`driven`-Port-Modulen; ohne eine Richtungs-Dimension können sie
diese Trennung nicht durchsetzen, ohne die Engine zu forken.

## Entscheidung

1. **Optionale Richtung `direction` ∈ {`driving`, `driven`}** auf `port`-/`adapter`-
   Schichten, **orthogonal** zur Rolle. `driving` = primär/inbound (Use-Case-Schnittstelle,
   vom Treiber-Adapter aufgerufen); `driven` = sekundär/outbound (vom Kern/App definiert,
   vom getriebenen Adapter implementiert). Die Reinheits-Regeln (`core-`/`app-`/
   `port-impurity`, `lateral-adapter`) bleiben **rollen-basiert unverändert**.
2. **Neue Regel `port-direction-mismatch`:** ein `role: adapter` mit Richtung X, der eine
   `role: port`-Schicht mit Richtung Y importiert (Y ≠ X, **beide gesetzt**), ist ein
   Befund. In der Erst-Treffer-Kette **vor** `wrong-direction` (die spezifischere Regel
   gewinnt). Schichten **ohne** `direction` unterliegen der Regel **nicht**; die
   `app`-Schicht ist richtungs-agnostisch und wird nicht erfasst.

**Verworfene Alternative — Subtyp-Rollen** (`port_driving`/`port_driven`/…): bläht das
Rollen-Enum und **jede** Reinheits-Prüfung auf (jede Regel müsste alle Subtypen kennen).
Die orthogonale Dimension ist sparsamer — nur eine neue Connectivity-Regel kommt hinzu,
die bestehenden Regeln bleiben unberührt.

## Konsequenzen

- Treiber/Getriebenen-Trennung voll prüfbar; Konsumenten mit getrennten `driving`/
  `driven`-Port-Modulen erzwingen sie per Config statt per Engine-Fork.
- **Richtung hängt an der Schicht, nicht an der Datei** (`model.Layer.Direction`): ein
  Treiber-Adapter spricht nur dann „nur `driving`-Ports", wenn die **Adapter-Schicht
  selbst** `direction: driving` trägt. Konsumenten brauchen also **richtungs-getrennte
  Adapter-Schichten** (spiegelbildlich zu den `driving`/`driven`-Port-Modulen), nicht nur
  getrennte Port-Module — ein realer Config-Aufwand, kein monolithischer `adapters`-Layer.
- **Kategorisch:** `edges`/`allow` heben `port-direction-mismatch` nicht auf (wie
  `lateral-adapter`); die Regel steht in der Erst-Treffer-Kette **vor** `wrong-direction`
  und greift daher auch bei einer deklarierten `allow`-Kante. Einzige Ausnahme:
  `composition_root` (global von allen Schicht-Regeln befreit).
- **Siebter Befund-Name** `port-direction-mismatch` — Output-Konsumenten (CI-Parser, Doku)
  **und der Engine-Doc-String** (`rules.go` „runs the … rules") nachziehen
  („sechs→sieben"-Sweep); bestehende Namen stabil.
- **Rückwärtskompat:** ohne `direction` ändert sich nichts — a-checks Dogfooding
  (Eigen-`.a-check.yml` ohne Richtung) bleibt grün. Die Richtung ist *opt-in und inert*.
- **Config-Schema:** Objekt-Form `{globs, role, direction}`, `direction ∈ {driving, driven}`,
  strict-decode (Exit 2 sonst).
- **`direction` ohne `role: adapter`/`port` ist inert:** die Regel braucht
  `srcRole==adapter`/`tgtRole==port`; auf einer rollenlosen oder `domain`/`app`-Schicht
  hat `direction` keine Wirkung. Der Decoder validiert nur das Enum (`driving|driven`) und
  erzwingt **keine** Rolle — eine Cross-Feld-Pflicht wäre mit der Namens-Inferenz der
  Rolle inkonsistent (die Rolle kann fehlen und inferiert werden).
- Lastenheft 0.5.0 → **0.6.0**.

## Fitness Function

- `make arch-check` (Dogfooding, [AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)):
  unverändert 0 Befunde **ohne** Änderung der Eigen-`.a-check.yml` (a-check hat keine `direction`).
- `make test`: `adapter(driving) → port(driving)` happy; `adapter(driving) → port(driven)`
  ⇒ `port-direction-mismatch` (kategorisch, genau ein Befund — **auch bei deklarierter
  `allow`-Kante**); Boundary ohne `direction` unverändert; fremd benannte Schichten voll
  geprüft.

## Re-Evaluierungs-Trigger

- Richtungs-Regeln zwischen Ports untereinander (Port→Port) — out-of-scope, späteres Inkrement.
- Auto-Inferenz der Richtung aus Namen/Pfad (`driving`/`driven`) — bewusst verworfen
  (explizit deklariert statt geraten).

## Geschichte

| Datum | Ereignis |
|---|---|
| 2026-06-23 | Proposed — welle-10b (b2b); `direction` als orthogonale Dimension + Regel `port-direction-mismatch` (Entscheid A: Attribut, B: minimal `adapter→port`). |
