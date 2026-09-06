**Stand:** verkörpert in [`make dcheck-phrase-selftest`](../../../../../../Makefile) (`tools/dcheck-phrase-selftest.sh`,
im `gates`-Aggregat) `seit slice-168`

Deckt die beiden real aufgetretenen phrasen-basierten Fälle (`reviews`-Trigger-Phrase, `structure`
`tasks-ignore-pattern`) mit je einer Positiv-/Negativ-Kontrolle gegen eigene Fixtures ab. **Nicht**
gedeckt: der `doc-complete`-Fall (Prüfer ohne Aufruf, nicht ohne Gegenstand) — der ist bereits seit
slice-123 durch die Aufnahme ins `verify`-Aggregat strukturell behoben, keine gesonderte Prüfung
nötig.
