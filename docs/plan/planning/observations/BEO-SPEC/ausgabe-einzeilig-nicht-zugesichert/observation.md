# Die Einzeiligkeit einer Befundzeile ist nicht zugesichert

**Sub-Area:** Spec-Straten

Das Befund-Format `pfad:zeile: regel: meldung` setzt **einen** Datensatz je Zeile voraus. Enthält
die Meldung Text aus der geprüften Datei oder der Konfiguration, kann ein Zeilenende darin den
Datensatz zerreißen — und der Vertrag sagt nicht, wie die Meldung damit umgeht.
