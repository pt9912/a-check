**Vorgang:** slice-186

**Fund:** Drei Stellen — [`AGENTS.md`](../../../../../../../AGENTS.md) §4, `harness/README.md`
§Sensors und der Vertrag in `harness/sensors/doc-planning.md` — sagten zu, die Roadmap-Sektion
**benenne** den Slice aus `in-progress/`. Geprüft wird nur die Äquivalenz *Slice vorhanden ⟺
Ruhe-Marker fehlt*; der genannte Name geht nicht ein. Die Einschränkung stand ausschließlich als
Kommentar in `.d-check.yml`, ausdrücklich als „BENANNTE GRENZE" — an einem Ort, den niemand liest,
der die Zusage liest.

**Was das gekostet hat:** Die Roadmap nannte **siebzehn Slice-Übergänge lang** einen Slice, der in
`done/` lag, und das Gate blieb grün. Gefunden hat es der unabhängige Review (Report zu slice-186,
F-7), nicht der Lauf.

**Behoben** durch die Grenze am Satz: Die Sensor-Datei führt sie als ersten Punkt, und beide
Deklarations-Stellen sagen zu, was geprüft wird.
