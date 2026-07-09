# ADR-0024 — `--print-graph`: Mermaid als Format und Render-Umfang

- **Status:** Accepted
- **Datum:** 2026-07-09
- **Autor:** pt9912
- **Bezug:** [AC-FA-CLI-002](../../../spec/lastenheft.md#ac-fa-cli-002--architektur-graph-ausgabe), [AC-QA-01](../../../spec/lastenheft.md#ac-qa-01--determinismus), [AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)
- **Schärft:** [SPEC-CLI-002](../../../spec/spezifikation.md#spec-cli-002--graph-renderer-vertrag) — legt Ausgabeformat, Render-Umfang und Determinismus-Ordnung der Graph-Ausgabe fest.
- **Supersedes:** —

## Kontext

`.a-check.yml` **ist** bereits ein Graph: `layers` sind Knoten, `edges`/`allow`
gerichtete Kanten, `role`/`direction` Typ und Gruppierung,
`composition_root`/`adapter_sink` Sonderfälle. Der neue no-scan-Modus
[AC-FA-CLI-002](../../../spec/lastenheft.md#ac-fa-cli-002--architektur-graph-ausgabe)
macht diese **deklarierte Absicht** sichtbar. Damit Change-Request,
Spezifikation und Tests nicht auseinanderlaufen, braucht die Ausgabe eine
festgelegte Format- und Umfangs-Entscheidung mit echten Wahl-Punkten:

- **Format:** Mermaid vs. DOT/Graphviz vs. JSON.
- **Umfang:** nur deklarierte Kanten — oder auch die kategorischen, regel-impliziten
  Constraints?
- **Sonderfälle:** wie `composition_root`/`adapter_sink`, die keine Layer-Kanten sind?
- **Ordnung:** wie wird Determinismus
  ([AC-QA-01](../../../spec/lastenheft.md#ac-qa-01--determinismus)) gegen die interne
  Map-/Slice-Ordnung des Modells garantiert?

## Entscheidung

1. **Format = Mermaid `flowchart`.** Die Repo-Doku (Roadmap-, Architektur-Diagramme)
   nutzt bereits Mermaid; die Ausgabe ist direkt in PRs/Docs einbettbar. Ein zweites
   Format käme später **additiv** über ein separates `--graph-format`-Flag, **nicht**
   durch Umbau von `--print-graph` (bleibt ein `bool`, damit die bare Form gültig
   bleibt).
2. **Render-Umfang v1 = deklarierte Kanten + effektive Rollen + Sonderknoten.**
   Gezeichnet werden `edges` (durchgezogen) und `allow` (abgesetzt) zwischen
   Layer-Knoten, gefärbt nach der **effektiven** Rolle (geteilter Kern-Resolver, keine
   kopierte Inferenz). `composition_root`/`adapter_sink` erscheinen als **isolierte
   Notizknoten** ohne gezeichnete Verdrahtungs-Kanten. `direction` gruppiert in
   stabile Subgraphs. `tech` ist v1 **deferred** (Folge-Inkrement).
3. **Implizite Regeln als Legende, nicht als Kante.** Die kategorischen
   Rollen-/Richtungs-Constraints (`core-impurity`, `lateral-adapter`,
   `port-direction-mismatch`) sind keine Config-Kanten; sie erscheinen als Legende —
   der Graph behauptet keine Semantik über den realen Code
   ([AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)).
4. **Determinismus durch stabile IDs + feste Ordnung.** Layer-Namen (freie
   YAML-Map-Keys) werden nie als Mermaid-ID benutzt: interne IDs `L*`/`D*`/`C0`/`S0`,
   der Rohname nur als **escaptes** Label; Knoten nach Name, Kanten nach
   (`from`, dann `to`), `classDef` in fester Reihenfolge — byte-identische Ausgabe
   ([AC-QA-01](../../../spec/lastenheft.md#ac-qa-01--determinismus)).

## Konsequenzen

- **Neuer Ausgabe-Vertrag, kein neues Scan-Verhalten.** Der Modus ist read-only und
  liest keine Quellen; unbekannte `edges`/`allow`-Endpunkte werden als Dangling-Knoten
  gezeigt statt neu mit Exit 2 abgelehnt (keine neue Config-Strenge).
- **Ehrlichkeitsgrenze gewahrt**
  ([AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)):
  Absicht ≠ Beweis; der Graph ersetzt nicht den Scan.
- **Additive Format-Erweiterung offen** (DOT/Graphviz, Findings-/Verstoß-Graph) —
  eigene Folge-Slices; die `bool`-Form `--print-graph` bleibt dauerhaft gültig.
- **Kein Versions-Bump am Lastenheft durch diesen ADR** — die Anforderung
  ([AC-FA-CLI-002](../../../spec/lastenheft.md#ac-fa-cli-002--architektur-graph-ausgabe))
  entstand im Change-Request; der ADR schärft nur die Spezifikation aufwärts
  ([SPEC-CLI-002](../../../spec/spezifikation.md#spec-cli-002--graph-renderer-vertrag)).

## Fitness Function

- `make test`: die Akzeptanzkriterien aus
  [AC-FA-CLI-002](../../../spec/lastenheft.md#ac-fa-cli-002--architektur-graph-ausgabe) /
  [SPEC-CLI-002](../../../spec/spezifikation.md#spec-cli-002--graph-renderer-vertrag)
  als E2E-/Unit-Tests (Happy/Boundary/Negative, Mermaid-unsafe-Namen, Dangling,
  effektive Rollen, `direction`-Subgraphs, Determinismus).
- `make arch-check` (Dogfooding,
  [AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)):
  der neue `graph`-Präsentationsadapter bleibt schichtkonform (kein lateraler
  Adapter-Import) — 0 Befunde.
- `make image-test`: `--print-graph` gegen das gebaute Image (Dekodierbarkeit +
  read-only + nativ==Container).

## Geschichte

| Datum | Ereignis |
|---|---|
| 2026-07-09 | Proposed — Format Mermaid + Render-Umfang v1 (deklarierte Kanten + effektive Rollen + Sonderknoten; implizite Regeln als Legende; Determinismus-Ordnung via stabile interne IDs); `tech` deferred, additive Format-Erweiterung (`--graph-format`) offen. |
| 2026-07-09 | Proposed → Accepted (Sign-off Auftraggeber; spec-first-Review der Vertrags-/Spezifikations-Schicht bestanden). Ab jetzt immutable; Ablösung nur via Folge-ADR mit `Supersedes`. |
