**Vorgang:** slice-217

**Fund (eigene Gegenprobe):** `tools/multiarch-check.sh` endete im ersten Lauf mit Exit 141
(`grep | head -1` und `tar | od -N2` unter pipefail), und die Gegenprobe „Index mit Attestierung"
war rot **ohne** Meldung — ein `grep` ohne Treffer beendete die Zuweisung. Zwei Funde, ein
Vorgang. Behoben: Extraktion über Hilfsfunktionen, die leer statt eines Abbruchs liefern
(`|| true`, `sed -n 1p`), Layer erst in eine Datei; jeder Fehlschlag erreicht eine Prüfung mit
Meldung.
