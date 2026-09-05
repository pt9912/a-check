# Verify vor git mv: falsche Reihenfolge

**Sub-Area:** Gate-/Werkzeug-Schicht
**Ehemals:** `BEO-006` (Tabellenform, bis slice-139)

Die vorgeschriebene Reihenfolge kann die Prüfung nicht abnehmen, die sie erfüllen soll: der Closure-/Risiko-Sensor greift nur in `done/`, der Workflow fährt `make verify` in Schritt 8 und den `git mv` erst in Schritt 9
