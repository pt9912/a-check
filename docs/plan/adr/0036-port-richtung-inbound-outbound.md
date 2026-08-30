# ADR-0036 — Die Richtungs-Dimension trägt je Rolle ihr eigenes Vokabular

- **Status:** Accepted
- **Datum:** 2026-08-30
- **Autor:** pt9912
- **Bezug:** [AC-FA-RULE-008](../../../spec/lastenheft.md#ac-fa-rule-008--richtungs-dimension-regel-port-direction-mismatch) (neu gefasst; Lastenheft 0.24.0→0.25.0), [AC-FA-CONF-001](../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml) (Schema)
- **Schärft:** [SPEC-CONF-001](../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema) + [SPEC-RULE-001](../../../spec/spezifikation.md#spec-rule-001--regel-auswertung) — macht die rollen-abhängige Wertemenge und die Paarungs-Prüfung verbindlich.
- **Supersedes:** ADR-0012

## Kontext

[ADR-0012](0012-driving-driven-richtung-orthogonale-dimension.md) führte `direction` ∈
{`driving`, `driven`} als orthogonale Dimension auf `port`- und `adapter`-Schichten ein.
Sie wägte zwei Alternativen ab — die orthogonale Dimension gegen Subtyp-Rollen
(`port_driving`/`port_driven`) — und entschied für die erste.

**Beide Alternativen trugen dasselbe Vokabular, und genau das war die ungestellte Frage.**
In der Hexagonal-Architektur benennen die zwei Seiten verschiedene Dinge: ein **Port** ist
eingehend oder ausgehend (`inbound`/`outbound`) — er beschreibt, in welcher Richtung die
Schnittstelle steht. Ein **Adapter** ist treibend oder getrieben (`driving`/`driven`) — er
beschreibt eine Rolle im Betrieb. Ein Port *treibt* nichts; er wird benutzt.

Die abgelöste Fassung wusste das und schrieb es hin — *„`driving` = primär/inbound"*,
*„`driven` = sekundär/outbound"* —, ließ aber nur die Adapter-Hälfte als **Wert** zu. Wer die
Standard-Terminologie an einem Port benutzt, bekam `ungültige direction "inbound"
(driving|driven)` und Exit 2.

## Entscheidung

1. **Der Wertebereich von `direction` ist rollen-abhängig:** an `role: port` gilt
   `inbound|outbound`, an `role: adapter` `driving|driven`. Die Dimension bleibt orthogonal
   zur Rolle und opt-in; ohne `direction` ändert sich nichts.
2. **`port-direction-mismatch` prüft eine Paarung statt einer Gleichheit.** `driving` gehört
   zu `inbound`, `driven` zu `outbound`. Die Regel-Aussage ist unverändert: ein treibender
   Adapter spricht nur eingehende Ports. Nur die Mechanik wechselt — bisher genügte der
   Vergleich zweier Strings, weil beide Seiten dasselbe Wort trugen.
3. **Die falsche Vokabel an einer Rolle ist ein Konfigurationsfehler** (Exit 2), kein
   akzeptiertes Alias. Die Meldung nennt Schicht, Wert und die für **diese Rolle** gültige
   Menge.

**Verworfene Alternative — stilles Alias** (`driving` am Port intern auf `inbound` abbilden):
bricht keine bestehende Konfiguration, setzt die Unterscheidung aber nicht durch. a-check kennt
kein Warn-Level ([AC-FA-CLI-001](../../../spec/lastenheft.md#ac-fa-cli-001--aufruf-scan-wurzel-und-exit-codes):
Exit 0/1/2, keine Schwere-Stufen) — die alte Schreibweise bliebe unbemerkt, und die Korrektur
wäre dokumentiert statt gültig. Wer eine Terminologie richtigstellt und die falsche weiter
still annimmt, hat sie nicht richtiggestellt.

## Konsequenzen

- **Breaking für Konfigurationen mit Richtung an Port-Schichten.** Die Umstellung ist eine
  Zeile je Schicht; die Meldung nennt sie. Adapter-Schichten sind unberührt.
- **Die Zuordnung steht an zwei Orten** — in der Validierung (welcher Wert ist an welcher Rolle
  gültig) und in der Paarung (welcher Adapter-Wert passt zu welchem Port-Wert). Sie können
  auseinanderlaufen; ein Test hält beide gegen dieselbe Tabelle.
- **Die Fehlermeldung wird rollen-spezifisch.** Sie kann nicht mehr eine feste Menge nennen,
  sondern muss die Rolle kennen — sonst schickt sie den Leser auf die falsche Hälfte.
- **`app` bleibt richtungs-agnostisch** und wird von der Regel nicht erfasst; die Beschreibung
  nennt jetzt die richtigen Vokabeln (nutzt `outbound`-Ports, implementiert `inbound`-Ports).

## Fitness Function

`make test` — die Tabellen-Tests zur rollen-abhängigen Validierung und zur Paarung, je mit
ihrer **Umkehr**: die gültige Kombination schweigt, die ungültige meldet. `make gates`
(`arch-check` fährt a-checks Eigen-Konfiguration, die keine Richtung deklariert und damit die
Inertheits-Zusage aus Punkt 1 belegt).
