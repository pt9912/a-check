# awk/sed verliert Ausfuehrungs-Modus

**Sub-Area:** Gate-/Werkzeug-Schicht
**Ehemals:** `BEO-032` (Tabellenform, bis slice-139)

Eine Datei per `awk`/`sed` neu zu schreiben (Ausgabe in eine Temp-Datei, dann `mv`) verliert den **Ausführungs-Modus**: `100755` wird zu `100644`. Kein Gate meldet es — die Targets rufen `bash <skript>`, der Modus ist dort folgenlos
