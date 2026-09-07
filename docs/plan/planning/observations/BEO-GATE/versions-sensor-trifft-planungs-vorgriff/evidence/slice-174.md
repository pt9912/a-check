**Vorgang:** slice-174
**Fund:** Zweimal im selben Slice — **eine** Gelegenheit, kein zweifaches Auftreten (Baseline
`modul-06`: zwei Funde im selben Vorgang zählen einmal).

1. Die Welle-Datei `welle-15-regelwerk-v650-migration.md` nannte in §1 den Zielpfad als Literal;
   `make doc-check` meldete `version-stale` — **drei Minuten** nach Inbetriebnahme des Sensors in
   slice-173.
2. Der Etappen-Slice `slice-175` hatte ihn in der DoD stehen: *„`.harness/baseline/v6.5.0/` liegt
   mit eigenem `SHA256SUMS`"*. Derselbe Befund.

Beide Stellen waren **sachlich richtig** — sie beschreiben, wohin die Migration hebt. Umformuliert
zu „der neue Stand unter `.harness/baseline/`"; die Aussage bleibt, das Literal fällt weg.

**Was das über den Sensor sagt:** Er wirkt, und zwar sofort und an lebenden Dateien — das ist der
Zweck. Was fehlt, ist eine Unterscheidung, die er strukturell nicht treffen kann. Der Vorfall ist
darum kein Fehlalarm im engeren Sinn, sondern die Kehrseite seiner Schärfe.
