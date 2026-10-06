# Ein Lexik-Vertrag wird geschrieben, ohne gegen die Sprache gegengeprobt zu werden

**Sub-Area:** Spec-Straten

Ein Normalisierungs- oder Lexik-Vertrag in der Spezifikation beschreibt, wie Zeichenketten,
Vorlagen und Kommentare einer fremden Sprache erkannt werden — und wird aus dem Gedächtnis der
Sprache geschrieben statt gegen ihre Regeln und Zeichen für Zeichen geprüft. Jede Lücke dort ist
die **nicht fail-safe** Fehlrichtung: Code, den der Lexer für Kommentar hält, wird verschluckt.
