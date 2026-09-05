# Hebungskanal haengt an repo-externen Schaltern

**Sub-Area:** Gate-/Werkzeug-Schicht
**Ehemals:** `BEO-030` (Tabellenform, bis slice-139)

Die Wirkung des Hebungs-Kanals hängt an **zwei Schaltern außerhalb des Repos** (`dependabot_security_updates`, Dependabot-Alerts): ohne sie öffnet ein CVE **ohne** neues Upstream-Release keinen PR. Kein Gate kann sagen, ob sie stehen
