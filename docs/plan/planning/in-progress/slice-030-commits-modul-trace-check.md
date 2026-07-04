# slice-030 — Traceability-Gate via `commits`-Modul (löst `tools/trace-check.sh` ab)

**Status:** in-progress (2026-07-04 — umgesetzt, [ADR-0021](../../adr/0021-commits-modul-trace-check.md) Proposed).
**Typ:** Tool-Ablösung / Harness-Konsolidierung, nicht konsumenten-gated.
**Bezug:** [`AGENTS.md` §5](../../../../AGENTS.md#5-dokumentations-regeln) (Traceability-Regel),
[ADR-0021](../../adr/0021-commits-modul-trace-check.md); Ziel „Skript-Copies verringern".
Folge von [slice-019](../done/slice-019-dcheck-mk-print-mk-angleichung.md) (Pin v0.37.1 →
Modul verfügbar) + [slice-029](../done/slice-029-doc-check-module-hardening.md) (dort als
Konsolidierungs-Kandidat identifiziert). [Roadmap](roadmap.md).

## 1. Auslöser

`tools/trace-check.sh` (98 Zeilen bash) ist eine **Skript-Kopie** — selbst beschriftet mit
„Stack-Vorbild d-check `tools/trace-check.sh`". Genau diese Kopien zu tilgen ist der
Stack-Zweck. d-check hat sein eigenes Skript längst durch das Modul `commits` abgelöst;
a-check pinnt seit slice-019 `v0.37.1`, das das Modul mitbringt, und slice-029 hat die
Verhaltensgleichheit vorab belegt.

## 2. Umgesetzt (ADR-0021 Weg A)

1. **`.d-check.yml`**: `commits`-Config (`id-patterns` = die vier a-check-Muster
   `AC-`/`ADR-`/`MR-`/`slice`, `exempt-pattern` `^(Merge |Revert )`); **nicht** in der
   `modules`-Liste (fokussiertes Gate).
2. **`Makefile` `trace-check`**: fährt das Modul — `MSGFILE=<datei>` → `--commit-msg -`
   (Pending-Message via stdin, Hook); sonst `--enable commits` focus-disabled `--range`
   (`RANGE=` CI, sonst `HEAD~1..HEAD`).
3. **`.githooks/commit-msg`**: `exec make trace-check MSGFILE="$1"` (statt `bash trace-check.sh`).
4. **`tools/trace-check.sh` entfernt.**
5. Doku: [ADR-0021](../../adr/0021-commits-modul-trace-check.md), ADR-Index, AGENTS §4/§5,
   harness/README §Sensors + §Traceability, CHANGELOG-`[Unreleased]`.

## 3. Paritäts-Proben (Fitness Function — bestanden)

| Fall | Erwartet | Beobachtet |
|---|---|---|
| `make trace-check` (Default `HEAD~1..HEAD`) | 0 | ✓ 0 |
| `RANGE=HEAD~3..HEAD` (alle mit ID) | 0 | ✓ 0 |
| `MSGFILE` mit ID | rc 0 (still) | ✓ |
| `MSGFILE` ohne ID | rc ≠ 0 + Meldung | ✓ `commit-untraceable`, Exit 2 |
| `MSGFILE` `Merge …` | rc 0 (ausgenommen) | ✓ |
| `RANGE=deadbeef..HEAD` (kaputte Basis) | fail-closed | ✓ „Range-Basis nicht auflösbar", Exit 2 |

## 4. Trade & Restkost

- **commit-msg-Hook braucht jetzt Docker** (Container je Commit). Akzeptiert: Hook ist
  opt-in (`make hooks`), der CI-Range-Check ist die klon-unabhängige Wahrheit (Muster wie
  d-check). Bewusst getragen ([ADR-0021](../../adr/0021-commits-modul-trace-check.md) §Konsequenzen).
- **bash-Selbsttest entfällt** — Ersatz = d-checks upstream-getestetes Modul + die Proben §3.

## 5. Offen bis Merge

[ADR-0021](../../adr/0021-commits-modul-trace-check.md) Proposed → Accepted (Merge-Wort); Slice → `done/`; `make gates` grün.
