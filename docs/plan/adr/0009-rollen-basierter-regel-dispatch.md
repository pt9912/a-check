# ADR-0009 — Rollen-basierter Regel-Dispatch

- **Status:** Accepted
- **Datum:** 2026-06-22
- **Autor:** pt9912
- **Bezug:** [AC-FA-RULE-006](../../../spec/lastenheft.md#ac-fa-rule-006--schicht-rollen-generische-regel-anwendung) (neu, Lastenheft 0.2.0→0.3.0), [AC-FA-CONF-001](../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml) (Layer-Schema), §1 Zweck (Vier-Repo-Konsolidierung)
- **Schärft:** [SPEC-RULE-001](../../../spec/spezifikation.md#spec-rule-001--regel-auswertung) + [SPEC-CONF-001](../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema) — macht den rollen-basierten Regel-Dispatch und die Layer-Objektform (`{globs, role}`) verbindlich.
- **Supersedes:** —

## Kontext

Die Engine bindet drei der fünf Regeln an die **Literal-Namen**
`core`/`ports`/`adapters` (`internal/hexagon/core/rules.go`): die import-basierten
`core-impurity`/`port-impurity`/`lateral-adapter` in `ruleFor` **und** die
konstrukt-basierte `port-impurity` in `Evaluate` (`rules.go:24-28`) — vier Stellen.
`wrong-direction`/`tech-leak` sind bereits generisch. Das ist der
Re-Evaluierungs-Trigger aus [ADR-0008](0008-ports-duerfen-domaenen-typen-referenzieren.md).

Die vier Konsumenten-Repos haben **abweichende/feinere** Schicht-Strukturen
(b-cad `src/hexagon/ports`, d-migrate `hexagon/ports-{execute,read,write}`,
d-check `internal/hexagon/port`, grid eigene). Mit fremden Namen bekommen sie heute
nur `wrong-direction`/`tech-leak`, nicht die benannten Reinheits-Invarianten — die
Gründungslogik (ein Tool statt vier `arch-check.sh`) ist damit unvollständig
eingelöst.

`lateral()` unterscheidet zwei Adapter zudem über `adapterSeg` (Pfadsegment nach
dem Literal `"adapters"`), nicht über Layer-Identität — fremd benannte Adapter-Layer
(`src/geometry`, `src/persistence`) liefern beide `""` und werden nicht
unterschieden.

## Optionen

1. **Reine Edge-Modell-Generalisierung** (benannte Regeln streichen, alles über
   `edges`/`tech`). *Verworfen:* verliert die informativen Befund-Namen
   (`core-impurity` etc.), erzwingt feinkörnige Layer + viel Config, und die
   Intra-Layer-`lateral` ist über Kanten nicht abbildbar.
2. **`ports → core` implizit / Namen hart belassen.** *Verworfen:* löst das
   Vier-Repo-Problem nicht.
3. **Rollen-Dispatch (gewählt).** Die Regeln greifen über eine Layer-**Rolle**;
   die Rolle stammt aus `role:` oder aus Namens-Inferenz. Informative Befund-Namen
   bleiben, beliebige Namen werden prüfbar, Bestands-Configs laufen unverändert.

## Entscheidung

1. **Rollen-Dispatch:** Die Reinheits-Regeln dispatchen über eine Layer-Rolle
   ∈ {`domain`, `port`, `adapter`}. Die Rolle stammt aus `role:` (**Vorrang**) oder
   aus Namens-Inferenz (`core`→`domain`, `ports`→`port`, `adapters`→`adapter`).
   Layer ohne Rolle sind nur kanten-geprüft. Rollen-Mapping:
   `domain`→`core-impurity`, `port`→`port-impurity`, `adapter`→`lateral-adapter`.
2. **Import- UND konstrukt-basierte `port-impurity`** werden rollen-basiert (beide
   Stellen, sonst halb-generisch).
3. **`lateral-adapter` cross-layer:** feuert für Importe zwischen verschiedenen
   `role: adapter`-Schichten (Layer-Identität, `tgtLayer ≠ srcLayer`) und ist
   **kategorisch** — nur `adapter_sink` hebt auf, **nicht** `allow`/`edges`
   (konsistent mit der Intra-Schicht-`lateral`; `allow`/`edges` regieren die
   Schicht-*Richtung* via `wrong-direction`, nicht die Lateral-Invariante). Die
   `adapterSeg`-Intra-Unterscheidung bleibt für den klassischen Namen unverändert.
4. **Config:** ein `layers`-Eintrag ist eine Glob-Liste **oder** `{globs, role}`;
   strict-decode bleibt (unbekannter Schlüssel ⇒ Exit 2, auch im Objekt).
5. **Befund-Namen bleiben stabil** (`core-impurity`/`port-impurity`/`lateral-adapter`).

## Konsequenzen

- Beliebige/feinere Layer-**Namen** voll prüfbar — die vier Repos können ihre
  Strukturen ohne Engine-Fork abbilden.
- **Rückwärtskompatibilität 100 %:** klassische Namen (`core`/`ports`/`adapters`)
  werden inferiert; a-checks eigenes Dogfooding bleibt **ohne** Config-Änderung grün
  (ein einziger `adapters`-Layer ⇒ Cross-Layer-`lateral` feuert nie).
- **Migration:** Wer heute zwei verschieden benannte Adapter-Layer per `allow:`
  koppelt und beide auf `role: adapter` hebt, wird rot — die Kopplung muss über
  `adapter_sink` laufen (Folge der kategorischen `lateral`-Invariante).
- **`adapterSeg`-Namens-Generalisierung** der Intra-Schicht-Unterscheidung bleibt
  offen → **welle-10b** (zusammen mit `app`-Rolle und `driving`/`driven`-Ports).
- Lastenheft 0.2.0 → **0.3.0**.

## Fitness Function

- `make arch-check` (Dogfooding, [AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)):
  a-check prüft sich selbst **ohne** Änderung der Eigen-`.a-check.yml` grün (Inferenz).
- `make test`: Rollen-Dispatch mit fremden Namen, Inferenz-Boundary (klassische
  Config inkl. Konstrukt + Intra-`adapterSeg`), Cross-Layer-`lateral` (auch bei
  `allow`), Negative je Rolle (Import **und** Konstrukt).

## Re-Evaluierungs-Trigger

- **welle-10b:** `app`-Rolle (Domain/Application-Trennung), `driving`/`driven`-Ports,
  Namens-Generalisierung von `adapterSeg`.

## Geschichte

| Datum | Ereignis |
|---|---|
| 2026-06-22 | Proposed — welle-10a; generalisiert die vier namensgebundenen Reinheits-Stellen über Layer-Rollen, Namens-Inferenz für Rückwärtskompat. |
| 2026-06-22 | Proposed → Accepted (Sign-off Auftraggeber; Multi-Linsen-Review bestanden, T1–T4/K1 geschlossen). Ab jetzt immutable; Ablösung nur via Folge-ADR mit `Supersedes`. |
