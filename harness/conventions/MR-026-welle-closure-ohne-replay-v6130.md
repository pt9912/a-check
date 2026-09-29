# MR-026 — Welle-Closure ohne Replay-Lauf, Zeiger auf `v6.13.0` (löst [`MR-015`](../conventions.md#mr-015) ab)

- **Status:** Accepted
- **Datum:** 2026-09-29
- **Geltungsbereich:** [`docs/plan/planning/`](../../docs/plan/planning/README.md),
  Closure-Kriterien einer Welle
- **Ersetzt-Baseline-Regel:** [`modul-06-roadmap.md` §Wellen-Closure-Prozedur](../../.harness/baseline/v6.13.0/regelwerk/modul-06-roadmap.md#wellen-closure-prozedur-modul-6)
- **Adaption:** unverändert seit [`MR-015`](done/MR-015-welle-closure-ohne-replay.md).
  Die Baseline nennt als beobachtbaren Closure-Trigger einer Welle *„alle Slices in
  `done/` **und** `make gates` grün **und** der Replay-Lauf grün"*. Für a-check tritt an die Stelle
  des Replay-Laufs **`make ci` grün**.
- **Begründung:** dieselbe wie bei [`MR-015`](done/MR-015-welle-closure-ohne-replay.md) —
  a-check führt kein Golden Set; `make ci` ist der äquivalente repo-weite Beleg. Dieser Eintrag
  existiert, weil der Zielabschnitt im Sprung `v6.6.0` → `v6.13.0` **nicht wortgleich** blieb —
  der Zeiger wandert darum nicht still, sondern über diesen Nachfolger.
- **Gemessene Abweichung des Zielabschnitts** (`v6.6.0` → `v6.13.0`, 10 552 → 13 522 Zeichen,
  ohne die Quelle-Zeile): der Trigger-Audit prüft **vier** Artefaktklassen statt dreien — neu die
  **Hard Rule** (Welle 145); die Schritte 3 und 4 tragen den Planungs-Bestand-Lese-Schritt
  (Welle 150) und die Prosa-Erschöpfungs-Regel (Welle 152). **Die Replay-Forderung in Schritt 1
  ist wortgleich geblieben** (Zeile 45 beider Fassungen) — der Ausgang ist *bleibt gültig*.
- **Auflösungs-Trigger:** derselbe wie bei [`MR-015`](done/MR-015-welle-closure-ohne-replay.md) —
  sobald a-check eine nicht-deterministische Komponente enthält, entsteht ein Golden Set und der
  Ersatz entfällt.
- **Löst ab:** [`MR-015`](../conventions.md#mr-015)
- **Ausgelöst durch Baseline-Stand:** `v6.13.0`
