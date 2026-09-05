# Pruefer meldet gruen ohne Gegenstand oder Aufruf

**Sub-Area:** Gate-/Werkzeug-Schicht
**Ehemals:** `BEO-023` (Tabellenform, bis slice-139)

Ein Prüfer meldet grün, ohne kalibriert zu sein — in **zwei** Formen: ohne **Gegenstand** (leere Prüfmenge; `verify-ac-form` suchte seit slice-054 `^**Happy Path:**`, ein Wortlaut, den das Repo null mal führt) oder ohne **Aufruf** (`doc-complete` war advisory und lief nie; es hatte einen Gegenstand und hätte gemeldet)
