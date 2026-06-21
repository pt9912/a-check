# slice-008 — Lokaler commit-msg-Hook (Traceability opt-in)

**Status:** done.
**Welle:** welle-09-commit-hook.
**Bezug:** [`AGENTS.md`](../../../../AGENTS.md) §5 (Commit nennt `AC-*`/`ADR-*`-ID);
löst den in [slice-006 §4](slice-006-ci-pipeline.md#4-closure-notiz-nach-done)
benannten Folge-Kandidaten ein. Stack-Vorbild `d-check` (`.githooks/commit-msg`).

---

## 1. Ziel

Den lokalen `commit-msg`-Hook nachziehen, der `tools/trace-check.sh --message`
aufruft — dieselbe Wahrheit wie `make trace-check` und die CI. Fängt eine
ID-lose Message **vor** dem Commit; opt-in pro Klon via `make hooks`.

## 2. Definition of Done

- [x] `.githooks/commit-msg` ruft `tools/trace-check.sh --message "$1"` (eine Wahrheit).
- [x] `make hooks` setzt `core.hooksPath` auf `.githooks`.
- [x] `hooks` ist als Setup-Utility (kein Gate) in `tools/gate-consistency.sh` allowlisted.
- [x] [`harness/README.md`](../../../../harness/README.md) nennt den opt-in-Hook.
- [x] Beleg: `trace-check --message` lehnt eine ID-lose Message ab (Exit 1), lässt eine mit ID durch (Exit 0); `make gates` grün.

## 3. Umsetzung

- `.githooks/commit-msg` — `exec bash tools/trace-check.sh --message "$1"`.
- `Makefile` — `hooks`-Target (`git config core.hooksPath .githooks`).
- `tools/gate-consistency.sh` — `hooks` in `UTILITY_TARGETS` (Setup, kein Gate).
- `harness/README.md` §Traceability — opt-in-Hook benannt.

## 4. Closure-Notiz (nach `done/`)

**Belege:** `trace-check --message` verifiziert — ID-lose Message → Exit 1
(„nennt keine AC-/ADR-/MR-/slice-ID"), Message mit ID → Exit 0 (still);
`make gates` grün (`gate-consistency` ok mit `hooks`-Allowlist).

**Lerneintrag (Steering-Loop):**

- *Eine Wahrheit, drei Aufrufer:* der Hook dupliziert keine Logik — er `exec`t
  `trace-check.sh --message`, dieselbe Quelle wie `make trace-check` (HEAD) und
  der CI-Range-Check. Hook und CI bewerten identisch (kommentar-bereinigte Message).
- *Utility vs. Gate:* `hooks` ist Setup (kein Quality-Gate) → in
  `gate-consistency` allowlisted statt in die AGENTS-§4-Gate-Tabelle gezwungen;
  die Bidirektionalität bleibt scharf (kein halluziniertes/undokumentiertes Gate).
- *Opt-in mit Netz:* der Hook ist klon-lokaler Komfort (Frühwarnung); die
  verbindliche, klon-unabhängige Kontrolle bleibt der CI-Range-Check (slice-006).

**Offene Fragen:** keine — der slice-006-Folge-Kandidat ist eingelöst.

## 5. Sub-Area-Modus-Begründung

### Sub-Area: Traceability-Harness

- **Modus:** GF — kleiner Nachzug, neu angelegt, kein Bestand zu inventarisieren.
- **Konventionen-Dichte:** hoch (Stack-Vorbild d-check; slice-006-`trace-check`).
- **Phase-Reife:** Phase 5 — Gate/Hook stehen; `make gates` grün.
- **Evidenz-/Diskrepanz-Risiko:** niedrig (eine Skript-Wahrheit + Verhaltenstest).
- **Reconciliation-Aufwand:** keiner.
