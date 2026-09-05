# Weiche haengt an fremdem Event-Feld

**Sub-Area:** Gate-/Werkzeug-Schicht
**Ehemals:** `BEO-033` (Tabellenform, bis slice-139)

Eine Weiche, deren Verhalten an der Semantik eines **fremden** Event-Feldes hängt (`github.event.before`), ist lokal nicht prüfbar: der Selbsttest misst das eigene Skript, die Annahme über das fremde Feld bleibt unbelegt, bis ein echter Lauf sie widerlegt
