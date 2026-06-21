# slice-008 — Lokaler commit-msg-Hook (Traceability opt-in)

**Status:** open (Backlog; wartet auf Trigger/Priorisierung).
**Welle:** welle-09-commit-hook.
**Bezug:** [`AGENTS.md`](../../../../AGENTS.md) §5 (Commit nennt `AC-*`/`ADR-*`-ID);
löst den in [slice-006 §4](../done/slice-006-ci-pipeline.md#4-closure-notiz-nach-done)
benannten Folge-Kandidaten ein. Stack-Vorbild `d-check` (`.githooks/commit-msg`).

## Ziel

Den lokalen `commit-msg`-Hook nachziehen, der `tools/trace-check.sh --message`
aufruft — dieselbe Wahrheit wie `make trace-check` und die CI. So fängt das
Traceability-Gate eine ID-lose Message **vor** dem Commit (statt erst in der CI).
Opt-in pro Klon via `make hooks`; die klon-unabhängige Kontrolle bleibt der
CI-Range-Check.

## Definition of Done

- `.githooks/commit-msg` ruft `tools/trace-check.sh --message "$1"` (eine Wahrheit).
- `make hooks` setzt `core.hooksPath` auf `.githooks`.
- `hooks` ist als Setup-Utility (kein Gate) in `tools/gate-consistency.sh` allowlisted.
- [`harness/README.md`](../../../../harness/README.md) nennt den opt-in-Hook.
- Beleg: `make hooks` + ein Test-Commit ohne ID wird lokal abgelehnt; `make gates` grün.

## Offene Fragen

- Keine — kleiner Nachzug zu slice-006; die CI deckt den Pflichtfall bereits ab,
  der Hook ist reiner lokaler Komfort (Frühwarnung).
