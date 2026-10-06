**Vorgang:** slice-211

**Fund (Review F-5):** Der Kommentar zu `TestShapesUnusedEndToEnd` versprach die Zeile des
Eintrags in der `.a-check.yml`; die Assertion prüfte nur das Präfix `.a-check.yml:`. Behoben, die
Zeilennummer wird jetzt aus der Konfiguration abgeleitet und verglichen.
