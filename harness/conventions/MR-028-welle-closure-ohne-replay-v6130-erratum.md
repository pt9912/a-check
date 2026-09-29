# MR-028 — Welle-Closure ohne Replay-Lauf, Zeiger auf `v6.13.0` (löst [`MR-026`](../conventions.md#mr-026) ab, Erratum zur Schritt-Zuordnung)

- **Status:** Accepted
- **Datum:** 2026-09-29
- **Geltungsbereich:** [`docs/plan/planning/`](../../docs/plan/planning/README.md),
  Closure-Kriterien einer Welle
- **Ersetzt-Baseline-Regel:** [`modul-06-roadmap.md` §Wellen-Closure-Prozedur](../../.harness/baseline/v6.13.0/regelwerk/modul-06-roadmap.md#wellen-closure-prozedur-modul-6)
- **Adaption:** unverändert seit [`MR-015`](done/MR-015-welle-closure-ohne-replay.md).
  Die Baseline nennt als beobachtbaren Closure-Trigger einer Welle *„alle Slices in
  `done/` **und** `make gates` grün **und** der Replay-Lauf grün"*. Für a-check tritt an die Stelle
  des Replay-Laufs **`make ci` grün**.
- **Begründung:** dieselbe wie bei [`MR-026`](done/MR-026-welle-closure-ohne-replay-v6130.md) —
  a-check führt kein Golden Set; `make ci` ist der äquivalente repo-weite Beleg.
- **Was dieser Eintrag gegenüber [`MR-026`](done/MR-026-welle-closure-ohne-replay-v6130.md)
  korrigiert:** Jener ließ die gemessene Abweichung des Zielabschnitts in *„die Schritte 3 und 4
  tragen den Planungs-Bestand-Lese-Schritt (Welle 150) und die Prosa-Erschöpfungs-Regel
  (Welle 152)"* enden. Gemessen tragen **beide** neuen Blöcke in Schritt 3 der Prozedur;
  **Schritt 4 ist zwischen `v6.6.0` und `v6.13.0` unverändert** — gefunden im unabhängigen Review
  zu [slice-199](../../docs/plan/planning/done/slice-199-baseline-v6130-migration.md).
  Die tragende Aussage des Eintrags — die Replay-Forderung in Schritt 1 ist wortgleich geblieben,
  der Ausgang *bleibt gültig* — ist von dem Fehler nicht berührt.
- **Gemessene Abweichung des Zielabschnitts** (`v6.6.0` → `v6.13.0`, 10 552 → 13 522 Zeichen, ohne
  die Quelle-Zeile): der Trigger-Audit (Schritt 2) prüft **vier** Artefaktklassen statt dreien —
  neu die **Hard Rule** (Welle 145); **Schritt 3** trägt den Planungs-Bestand-Lese-Schritt
  (Welle 150) und die Prosa-Erschöpfungs-Regel (Welle 152). **Die Replay-Forderung in Schritt 1
  ist wortgleich geblieben** (Zeile 45 beider Fassungen) — der Ausgang ist *bleibt gültig*.
- **Auflösungs-Trigger:** derselbe wie bei [`MR-015`](done/MR-015-welle-closure-ohne-replay.md) —
  sobald a-check eine nicht-deterministische Komponente enthält, entsteht ein Golden Set und der
  Ersatz entfällt.
- **Löst ab:** [`MR-026`](../conventions.md#mr-026)
- **Ausgelöst durch Baseline-Stand:** `v6.13.0`
