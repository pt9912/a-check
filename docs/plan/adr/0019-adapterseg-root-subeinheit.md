# ADR-0019 — Root-Sub-Einheit: Dateien direkt im Layer-Root bilden die Sub-Einheit „''"

- **Status:** Accepted
- **Datum:** 2026-07-03
- **Autor:** pt9912
- **Bezug:** [AC-FA-RULE-002](../../../spec/lastenheft.md#ac-fa-rule-002--keine-lateralen-adapter-kanten-regel-lateral-adapter), [AC-FA-RULE-006](../../../spec/lastenheft.md#ac-fa-rule-006--schicht-rollen-generische-regel-anwendung) — **Erweiterung** von [ADR-0010](0010-layer-relativer-adapterseg-laengster-praefix.md) (layer-relativer `adapterSeg`), nach dem Muster, mit dem [ADR-0016](0016-resolution-sprach-parametrisch.md) [ADR-0014](0014-resolution-roots.md) erweiterte.
- **Schärft:** [SPEC-RULE-001](../../../spec/spezifikation.md#spec-rule-001--regel-auswertung).
- **Supersedes:** —

## Kontext

[ADR-0010](0010-layer-relativer-adapterseg-laengster-praefix.md) definierte die Adapter-Sub-Einheit
als **erstes Pfadsegment nach dem Schicht-Glob-Präfix**. Für Dateien, die **direkt im
Layer-Root** liegen (kein weiteres Verzeichnis), degeneriert dieses Segment zum
**Dateinamen** — zwei Dateien desselben Adapters gelten dann als verschiedene
Sub-Einheiten, und jedes eigene `.cpp → .h`-Include feuert als `lateral-adapter`.

Der b-cad-Pilot (2026-07-03, verifizierter v0.8.0-Lauf) macht das messbar: die
Richtungs-Modellierung ([AC-FA-RULE-008](../../../spec/lastenheft.md#ac-fa-rule-008--driving-driven-port-richtung-regel-port-direction-mismatch))
verlangt **pro-Adapter-Layer** (eine `direction` je Schicht; b-cads Adapter sind
verschieden gerichtet), und b-cads Adapter halten ihre Dateien flach im
Verzeichnis-Root — Ergebnis: **40 Falsch-Positive** `lateral-adapter` (ausnahmslos
Same-Directory-Includes wie `io/dxf_reader.cpp → "adapters/io/dxf_reader.h"`),
null echte Verstöße. Die Richtungsprüfung ist damit heute nur um den Preis eines
vergifteten Gates aktivierbar.

Der Kern des Problems: **Dateinamen sind keine Architektur-Einheiten.**
Sub-Einheiten eines Adapters sind — wie Schichten selbst — **Verzeichnisse**;
eine Datei im Layer-Root gehört zur „Wurzel-Einheit" des Layers, nicht zu einer
eigenen Einheit namens ihres Dateinamens.

## Optionen

| Weg | Idee | Bewertung |
|---|---|---|
| **A — Root-Sub-Einheit `''`** | Enthält der Pfad-Rest nach dem Glob-Präfix keinen weiteren `/`, ist die Sub-Einheit `''` (Root). Root↔Root = gleiche Einheit; Root↔Unterverzeichnis und Unterverzeichnis↔anderes Unterverzeichnis bleiben lateral. | **Gewählt.** Minimal, verzeichnis-treu (konsistent zur Layer-Definition), tilgt exakt die Falsch-Positiv-Klasse; Cross-Layer-lateral bleibt unberührt kategorisch. |
| **B — `direction` je Sub-Einheit** | Richtung nicht je Layer, sondern je Adapter-Sub-Einheit deklarieren (kein Layer-Split nötig). | Verworfen: neues Config-Konzept + Schema-Umbau für ein Problem, das Weg A mit einer Segment-Regel löst; [ADR-0012](0012-driving-driven-richtung-orthogonale-dimension.md)-Modell (Richtung = Layer-Dimension) bliebe gebrochen. |
| **C — Konsumenten-Pflicht: Unterverzeichnisse** | Repos müssen je Adapter mindestens ein Unterverzeichnis führen. | Verworfen: erzwingt künstliche Struktur in fremden Repos für eine reine Tool-Semantik-Frage; b-cad müsste ~30 Dateien grundlos verschieben. |

## Entscheidung

**Weg A**, als **Erweiterung** von [ADR-0010](0010-layer-relativer-adapterseg-laengster-praefix.md)
(längster-Präfix-Wahl und Segment-Bewusstsein gelten unverändert):

1. `adapterSeg` liefert für einen Pfad, dessen Rest nach dem Schicht-Glob-Präfix
   **keinen weiteren `/`** enthält, die Sub-Einheit nach **Blatt-Klassifikation**
   (Schärfung D, Dogfooding-Fund der Umsetzung): ein **datei-förmiges** Blatt
   (letztes Segment enthält `.`) gehört zur Root-Sub-Einheit **`''`**; ein
   **verzeichnis-förmiges** Blatt (ohne `.`) **ist** die Sub-Einheit — ein
   Go-Paket-Import endet auf dem Paket-**Verzeichnis** (`…/driven/report`), und
   eine Root-Deutung würde die Cross-Paket-Lateral-Erkennung in Go-Repos blenden
   (der Eigen-Paket-Import eines externen Testpakets war der aufdeckende Fall).
2. Die Regel gilt **uniform** für die Datei-Seite und die (gemäß `resolution`
   normalisierte, [ADR-0017](0017-relative-resolution-modus.md)) Import-Kandidaten-Seite.
3. **Wirkung auf `lateral-adapter`** ([AC-FA-RULE-002](../../../spec/lastenheft.md#ac-fa-rule-002--keine-lateralen-adapter-kanten-regel-lateral-adapter)):
   Same-Layer-Imports zwischen Root-Dateien sind **keine** lateralen Kanten mehr;
   Root ↔ Unterverzeichnis sowie verschiedene Unterverzeichnisse bleiben lateral;
   **Cross-Layer**-Adapter-Importe bleiben unverändert kategorisch.
4. **Ehrliche Einordnung:** Für einen Sammel-Layer mit flach gemischten Adaptern
   im Root entfällt damit die (scheinbare) Datei-Ebene-Trennung — sie war nie eine
   Architektur-Aussage, sondern ein Artefakt der Segment-Regel. Wer Sub-Einheiten
   prüfen will, strukturiert sie als Verzeichnisse (wie Schichten).

## Konsequenzen

- [ADR-0010](0010-layer-relativer-adapterseg-laengster-praefix.md) **bleibt gültig** —
  ADR-0019 präzisiert nur den Degenerat-Fall „Rest ohne `/`"; kein Supersede.
- **Richtungs-Modellierung wird verlustfrei:** pro-Adapter-Layer erzeugen keine
  Falsch-Positive mehr; die verifizierte b-cad-Vollrichtungs-Config wird exakt
  befundfrei (Fitness Function).
- [SPEC-RULE-001](../../../spec/spezifikation.md#spec-rule-001--regel-auswertung)
  (`lateral-adapter`-Zeile) nennt die Root-Sub-Einheit ausdrücklich.

## Fitness Function

- `make test`: Root↔Root same-layer (datei-förmig) → kein Befund (die b-cad-Klasse
  `x.cpp → "…/x.h"`); Root↔Unterverzeichnis → lateral; zwei Unterverzeichnisse →
  lateral (Bestand); Cross-Layer → lateral (Bestand, kategorisch); uniform auch
  über `resolution`-Kandidaten (`relative`/`fixed-root`); **Go-Paket-Blätter:**
  Eigen-Paket-Import (Testpaket) → kein Befund, Fremd-Paket-Import → lateral
  (pinnt die Blatt-Klassifikation gegen die Root-Blendung).
- `make arch-check` (Dogfooding): unverändert 0 (a-checks eigene Adapter liegen in
  Unterverzeichnissen).
- **Pilot-Beleg:** verifizierte b-cad-Vollrichtungs-Config (V4, 2026-07-03:
  40 Falsch-Positive, 0 echte) ⇒ nach diesem ADR **0 Befunde**.

## Re-Evaluierungs-Trigger

- **Sub-Einheiten-Bedarf auf Datei-Ebene** (ein Pilot will bewusst Datei-granulare
  Adapter-Trennung): eigenes Inkrement (z. B. Opt-in-Schalter), nicht Default.
- **Endungslose Datei-Specifier** (TypeScript `./b` ohne NodeNext-Endung): textuell
  nicht von einem Verzeichnis-Blatt unterscheidbar → gelten als Sub-Einheit und
  melden aus einer Root-Datei heraus lateral (dokumentierte Grenze,
  [AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze));
  Re-Eval beim realen TypeScript-Mono-Repo-Pilot (m-trace-Klasse).
- **Verzeichnisnamen mit Punkt** (Review-R1: die Gegenrichtung derselben Ambiguität —
  ein Paket-Verzeichnis wie `yaml.v2`/`config.v1` als Blatt gilt datei-förmig → Root,
  Lateral zwischen solchen Sub-Einheiten wird geblendet, falsch-negativ): per Test
  gepinnte, dokumentierte Grenze ([AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze));
  Re-Eval, falls ein Pilot real gepunktete Sub-Einheiten-Verzeichnisse führt.

## Geschichte

| Datum | Ereignis |
|---|---|
| 2026-07-03 | Proposed — Entwurf mit [slice-024](../planning/open/slice-024-adapterseg-root-subeinheit.md); Evidenz b-cad-Pilot (V3b: 39, V4 nach den dortigen Umbauten: 40 Falsch-Positive). |
| 2026-07-03 | Sign-off Auftraggeber für die Entscheide A–C (gemäß Empfehlung, [slice-024](../planning/open/slice-024-adapterseg-root-subeinheit.md) §6). |
| 2026-07-03 | **In der Umsetzung geschärft (Entscheid D, zur Abnahme):** Blatt-Klassifikation datei-förmig (`.`) → Root `''` vs. verzeichnis-förmig → Sub-Einheit — das Dogfooding deckte auf, dass die reine Root-Regel Go-Paket-Blätter (`report_test.go` → Eigen-Paket) falsch meldet und Cross-Paket-Lateral blenden würde; endungslose Specifier als dokumentierte Grenze in die Re-Eval-Trigger. Status bleibt Proposed bis D abgenommen. |
| 2026-07-03 | Proposed → Accepted (Sign-off Auftraggeber: Entscheid D mit dem Merge-Wort bestätigt, [slice-024](../planning/open/slice-024-adapterseg-root-subeinheit.md) §6). Ab jetzt immutable; Ablösung nur via Folge-ADR mit `Supersedes`. |
