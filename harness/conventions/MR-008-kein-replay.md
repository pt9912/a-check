# MR-008 — Kein Replay, keine Agenten-Telemetrie

- **Datum:** 2026-07-25
- **Geltungsbereich:** gesamtes Repo; betrifft die Baseline-Module `modul-12`
  (Replay/Evaluierung) und `modul-15` (Observability)
- **Adaption:** a-check führt **kein** Replay-Manifest (`evals/golden/…`, Golden Sets,
  Drift-Rate) und **keine** Agenten-Telemetrie (Tool-Call-Spans, Token-Attribution pro Rolle,
  Cache-Counter). Beide Module bleiben unverkörpert.
- **Begründung:** Beide adressieren die Nicht-Determinismus- und Kosten-Risiken einer
  **Agenten-Laufzeit**. a-check ist ein deterministisches CLI: gleiche Eingabe, gleiche
  Ausgabe, vertraglich zugesichert
  ([SPEC-DET-001](../../spec/spezifikation.md#spec-det-001--determinismus-vertrag)) und durch
  Akzeptanztests belegt. Ein Golden Set über deterministischen Output ist ein zweiter,
  schlechter gepflegter Testbestand; Span-Telemetrie über einen Prozess ohne Modellaufruf misst
  nichts. Die *Reproduzierbarkeits*-Hälfte, die `modul-12` über den `image_hash` einfordert,
  ist ohnehin erfüllt — digest-gepinnte Basis-Images und ein digest-gepinntes Release
  ([AC-QA-03](../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit)).
- **Folgewirkung, ausdrücklich:** das Baseline-Closure-Kriterium einer Welle „Replay-Lauf grün"
  (`modul-06`) ist für dieses Repo **unerfüllbar** und wird durch „`make ci` grün" ersetzt. Ohne
  diese Zeile bliebe ein Wellen-Closure dauerhaft unvollständig, ohne dass jemand sagen könnte,
  warum.
- **Auflösungs-Trigger:** sobald a-check eine nicht-deterministische Komponente enthält (etwa
  ein Modell-gestütztes Heuristik-Modul) oder Agenten-Läufe **im Repo selbst** abrechenbar
  werden. Beides ist heute nicht absehbar; der Eintrag ist bis dahin **nicht** permanent,
  sondern begründet ausgesetzt. Gefunden als B-19 in
  [slice-048](../../docs/plan/planning/done/slice-048-modul-delta-lesen.md).
