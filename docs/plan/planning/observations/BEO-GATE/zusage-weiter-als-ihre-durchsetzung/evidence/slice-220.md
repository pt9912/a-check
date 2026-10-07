**Vorgang:** slice-220

**Fund (Review F-4):** Der Handbuch-Satz „jedes Release läuft durch den Image-Test, und seine
Ausgabe muss der auf `linux/amd64` gleichen" sagte mehr zu, als `release.yml` vergleicht: geprüft
wird der Hash über Exit-Code, stdout und stderr **eines** Scans im Image-Test. Behoben: der Satz
nennt den Scan einer festen Test-Fixture.
