# MR-014 — Keine Agenten-Telemetrie (schärft [`MR-008`](../conventions.md#mr-008))

- **Status:** Accepted
- **Datum:** 2026-08-29
- **Geltungsbereich:** gesamtes Repo; betrifft das Baseline-Modul `modul-15`
- **Ersetzt-Baseline-Regel:** [`modul-15-observability.md` §Kernidee](../../.harness/baseline/v5.12.0/regelwerk/modul-15-observability.md#kernidee-modul-15)
- **Adaption:** a-check führt **keine** Agenten-Telemetrie — keine Tool-Call-Spans, keine
  Token-Attribution pro Rolle, keine Cache-Counter. Das Modul bleibt unverkörpert.
- **Begründung:** Span-Telemetrie über einen Prozess ohne Modellaufruf misst nichts. a-check ist
  ein deterministisches CLI ([SPEC-DET-001](../../spec/spezifikation.md#spec-det-001--determinismus-vertrag)),
  die Reproduzierbarkeits-Hälfte ist über digest-gepinnte Images erfüllt
  ([AC-QA-03](../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit)).
- **Was mit diesem Stand **entfällt**:** die Replay-Hälfte von [`MR-008`](../conventions.md#mr-008).
  `modul-12` grenzt sich in `v5.12.0` **selbst** auf den *nicht-deterministischen Kern* ein — ein
  Satz, den `v3.5.2` nicht kennt. Wo kein solcher Kern ist, greift das Modul nicht, und es braucht
  dafür keine Adaption mehr.
- **Auflösungs-Trigger:** sobald Agenten-Läufe **im Repo selbst** abrechenbar werden.
- **Löst auf:** [`MR-008`](../conventions.md#mr-008)
- **Ausgelöst durch Baseline-Stand:** `v5.12.0`
