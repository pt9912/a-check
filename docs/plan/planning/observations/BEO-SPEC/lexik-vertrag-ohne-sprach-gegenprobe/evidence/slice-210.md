**Vorgang:** slice-210

**Fund:** Der unabhängige Review der Implementierung (F-2) fand eine vierte Lexik-Lücke, diesmal
im **Vertrag** selbst, nicht im Code: In Kotlin ist `$` vor einem Backtick-Bezeichner auch in
einer Zeichenkette eine Vorlage; [SPEC-EXTRACT-001](../../../../../../../spec/spezifikation.md#spec-extract-001--import-extraktion)
kannte nur `${`. Ein `"` im Bezeichner beendete die Zeichenkette zu früh, ein späteres `/*`
verschluckte Code, die Datei blieb grün. Behoben über Plan-Änderung, Spezifikation 0.34.0 und
einen Regressionstest mit Mutations-Gegenprobe.

**Was es zeigt:** Die drei Korrekturen aus slice-209 schlossen die gefundenen Fälle, nicht die
Klasse. Die Grammatik der Zielsprache wurde wieder aus dem Gedächtnis aufgezählt, nicht gegen ihre
Lexer-Spezifikation abgeglichen.
