**Stand:** verkörpert in [`make symlink-check`](../../../../../../Makefile)
(`tools/symlink-check.sh`, im `gates`-Aggregat) `seit slice-173`.

Jeder von git getrackte Symlink muss auflösen; ein Ziel, das nicht existiert, ist Exit 1. Der
Geltungsbereich ist bewusst weiter als der Anlass — die Fehler-Familie ist *„Symlink zeigt ins
Leere"*, nicht *„Baseline-Symlink zeigt ins Leere"*.

**Nicht** geprüft und damit offen geblieben: ob das Ziel das **richtige** ist. Ein Symlink auf ein
existierendes, aber veraltetes Modul bleibt grün — das wäre ein Urteil über Absicht, kein Match.
