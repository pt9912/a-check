# Der Versions-Sensor trifft den Planungs-Vorgriff auf den Zielstand

**Sub-Area:** Gate-/Werkzeug-Schicht

Das `versions`-Muster für Baseline-Pins (`seit slice-173`) hält jeden Pfad
`.harness/baseline/<tag>/` gegen den **adoptierten** Stand. Ein Planungsdokument einer Migration
nennt aber zwangsläufig den **Ziel**-Stand — die Welle-Datei beschreibt, wohin gehoben wird, der
Etappen-Slice hat ihn in seiner DoD stehen. Beides ist korrekt geschrieben und wird beanstandet.

Der Sensor kann die zwei Fälle nicht unterscheiden: ein **vergessener Nachzug** (das, wogegen er
gebaut ist) und ein **beabsichtigter Vorgriff** sehen als Zeichenfolge gleich aus. Die
Unterscheidung liegt in der Absicht des Dokuments, und die ist kein Match.

**Nicht** dasselbe wie die Zeitdokument-Klassen: Die zitieren einen **vergangenen** Stand und
liegen in Verzeichnissen, die sich als Glob fassen lassen (`done/`, `docs/reviews/`, …). Der
Vorgriff steht in **lebenden** Dateien — `docs/plan/planning/<welle>.md`, `open/`, `in-progress/` —,
also genau dort, wo der Sensor wirken soll. Ein Glob darüber nähme ihm die Prüfmenge.

**Umgangen, nicht gelöst:** Beide Male wurde die Stelle so umformuliert, dass kein Pfad-Literal
dasteht („der neue Stand unter `.harness/baseline/`"). Das ist zumutbar und verliert wenig — der
Pfad existiert zum Schreibzeitpunkt ohnehin nicht, ein Link darauf bräche also ohnehin. Ob es
tragfähig bleibt, wenn eine Migration mehr Planungsdokumente führt, ist offen.
