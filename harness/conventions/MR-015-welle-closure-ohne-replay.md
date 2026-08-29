# MR-015 — Welle-Closure ohne Replay-Lauf (schärft [`MR-008`](../conventions.md#mr-008))

- **Status:** Accepted
- **Datum:** 2026-08-29
- **Geltungsbereich:** [`docs/plan/planning/`](../../docs/plan/planning/README.md),
  Closure-Kriterien einer Welle
- **Ersetzt-Baseline-Regel:** [`modul-06-roadmap.md` §Wellen-Closure-Prozedur](../../.harness/baseline/v5.12.0/regelwerk/modul-06-roadmap.md#wellen-closure-prozedur-modul-6)
- **Adaption:** Die Baseline nennt als beobachtbaren Closure-Trigger einer Welle *„alle Slices in
  `done/` **und** `make gates` grün **und** der Replay-Lauf grün"*. Für a-check tritt an die Stelle
  des Replay-Laufs **`make ci` grün**.
- **Begründung:** Ein Replay-Lauf setzt ein Golden Set voraus, das a-check nicht führt und nach
  `modul-12` auch nicht führen muss — dessen Regeln sprechen ausdrücklich vom
  *nicht-deterministischen Kern*, den dieses Repo nicht hat. Ohne diese Zeile bliebe **jedes**
  Wellen-Closure dauerhaft unvollständig, ohne dass jemand sagen könnte, warum. `make ci` ist der
  äquivalente repo-weite Beleg: `gates` plus `image-test` gegen das gebaute Image.
- **Auflösungs-Trigger:** derselbe wie bei [`MR-014`](../conventions.md#mr-014) — sobald a-check eine
  nicht-deterministische Komponente enthält, entsteht ein Golden Set und der Ersatz entfällt.
- **Löst auf:** [`MR-008`](../conventions.md#mr-008)
- **Ausgelöst durch Baseline-Stand:** `v5.12.0`
