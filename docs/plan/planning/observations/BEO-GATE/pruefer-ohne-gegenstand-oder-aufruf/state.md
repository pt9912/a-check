**Stand:** geplant — [`slice-169`](../../../in-progress/slice-169-korpus-seitige-kalibrierung.md)
schreibt die fehlende Hälfte.

Bereits verkörpert ist die **Werkzeug-Seite**:
[`make dcheck-phrase-selftest`](../../../../../../Makefile)
(`tools/dcheck-phrase-selftest.sh`, im `gates`-Aggregat) `seit slice-168` — vier Kontrollen gegen
eigene Fixtures, die belegen, dass `d-check` auf die gewählten Phrasen (`reviews`-Trigger-Phrase,
`structure` `tasks-ignore-pattern`) noch reagiert.

**Nicht** verkörpert ist die **Korpus-Seite** — ob a-checks eigene Dokumente die Phrase noch
tragen. Genau die ist zweimal ausgefallen (slice-120, slice-165); die Kandidatenmenge des
`reviews`-Moduls ist heute nicht leer, kann es aber jederzeit wieder werden, ohne dass ein Sensor
es sagt.

Der `doc-complete`-Fall (Prüfer ohne **Aufruf**, nicht ohne Gegenstand) ist seit slice-123 durch
die Aufnahme ins `verify`-Aggregat strukturell behoben; eine gesonderte Prüfung braucht er nicht.
