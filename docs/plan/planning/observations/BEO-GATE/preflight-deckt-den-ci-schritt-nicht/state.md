**Stand:** offen (1×)

Unterhalb der Schwelle. Der belegte Fall ist mit `slice-198` geheilt: `make preflight` fährt `ci`
**plus** die drei Range-Schritte über `origin/main..HEAD` — die Range, die der nächste Push **enthält**.
Eine **leere** Range meldet er als WARNUNG (direkt nach einem Push ist sie leer und
prüft damit nichts).

**Was beim zweiten Auftreten zu prüfen wäre:** ob sich die Differenz **mechanisch** fassen lässt.
Der naheliegende Sensor — *„jeder `make`-Aufruf in einem Workflow ist in einem Aggregat oder
deklariert"* — wäre prüfbar, verlangt aber eine gepflegte Zuordnung Workflow-Schritt ↔ Target; ob
die trägt, ist eine andere Frage als ihre Formulierung.
