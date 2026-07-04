# ADR-0021 — Traceability-Gate via d-check-Modul `commits` (löst `tools/trace-check.sh` ab)

- **Status:** Accepted
- **Datum:** 2026-07-04
- **Autor:** pt9912
- **Bezug:** [`AGENTS.md` §5](../../../AGENTS.md#5-dokumentations-regeln) (Traceability-Regel), [`harness/README.md` §Traceability rules](../../../harness/README.md#traceability-rules); d-check-Modul `commits` (DC-FA-COMMITS-001).
- **Schärft:** — (Harness-Gate, kein Spec-Stratum — wie [ADR-0005](0005-lint-profil.md)/[ADR-0006](0006-coverage-gate.md)).
- **Supersedes:** — (die Skript-Mechanik entstand als Bootstrap-Gate ohne eigenen ADR).

## Kontext

`make trace-check` erzwingt die Traceability-Regel (jede Commit-Message nennt eine
`AC-`/`ADR-`/`MR-`/`slice`-ID, [`AGENTS.md` §5](../../../AGENTS.md#5-dokumentations-regeln)).
Die Mechanik ist heute `tools/trace-check.sh` — **98 Zeilen bash**, drei Aufrufer
(commit-msg-Hook, CI-Range, lokal), ausdrücklich mit „Stack-Vorbild d-check
`tools/trace-check.sh`" beschriftet: **eine Skript-Kopie**.

Genau diese Kopien zu tilgen ist der Stack-Zweck (Konfiguration/Modul statt Fork). Das
Schwester-Repo `d-check` hat sein *eigenes* `tools/trace-check.sh` durch das Modul
`commits` abgelöst; a-checks digest-gepinntes `d-check` (`v0.37.1`) bringt dieses Modul
mit. Eine Vorab-Probe zeigte: das `commits`-Modul liefert mit a-checks
`AC-`/`ADR-`/`MR-`/`slice`-Mustern **dasselbe Verdikt** wie das Skript.

## Optionen

| Weg | Idee | Bewertung |
|---|---|---|
| **A — `commits`-Modul** | Alle drei Aufrufer über `make trace-check` auf das digest-gepinnte `commits`-Modul routen; `tools/trace-check.sh` entfernen. | **Gewählt.** Tilgt die Skript-Kopie (Stack-Zweck); *eine* getestete Wahrheit (d-checks Modul) statt gepflegter bash-Logik; verifiziert verhaltensgleich. Muster deckt sich mit d-checks eigener Ablösung. |
| **B — Skript behalten** | Status quo. | Verworfen: eine gepflegte Skript-Kopie, die von d-checks weiterentwickeltem `commits`-Modul driftet — die Divergenz-Klasse, die der Stack bekämpft. |
| **C — Hybrid (Skript für Hook, Modul für CI)** | Zwei Mechaniken. | Verworfen: doppelte Wahrheit, schlechter als beide reinen Wege. |

## Entscheidung

**Weg A.** `make trace-check` fährt das Modul `commits` (focus-disabled, damit nur
`commits` läuft) über eine Commit-Range; `MSGFILE=<datei>` prüft stattdessen die
**Pending-Message** (`--commit-msg -` via stdin):

- **CI:** `make trace-check RANGE="$RANGE"` (unveränderter Aufruf, jetzt Modul).
- **commit-msg-Hook:** `make trace-check MSGFILE="$1"` (Pending-Message; **braucht Docker**).
- **lokal:** `make trace-check` (Default `HEAD~1..HEAD`).

Die `commits`-Config (`id-patterns` = die vier a-check-Muster, `exempt-pattern` für
`Merge `/`Revert `) lebt in [`.d-check.yml`](../../../.d-check.yml); das Modul ist **nicht**
in der `modules`-Liste (fokussiertes Gate, im Default-`doc-check` inert).
`tools/trace-check.sh` wird entfernt.

## Konsequenzen

- **Der commit-msg-Hook braucht Docker** (Container-Spin-up je Commit). Akzeptiert: der
  Hook ist **opt-in** (`make hooks`), und die klon-unabhängige Kontrolle ist ohnehin der
  **CI-Range-Check** — der Hook ist lokale Bequemlichkeit, kein Verlass-Gate (Muster wie
  d-check).
- **Der bash-Selbsttest entfällt** (das Skript bewies je Lauf, dass es eine fehlende ID
  fängt). Der Ersatz ist d-checks upstream-getestetes Modul plus die Paritäts-Proben unten;
  Rest-Risiko bewusst getragen.
- **Eine Skript-Kopie weniger** (`tools/` von 4 auf 3 Skripte); `--print-mk`-Konsistenz mit
  d-check steigt.
- Die Traceability-**Regel** ([`AGENTS.md` §5](../../../AGENTS.md#5-dokumentations-regeln))
  ist unverändert; nur die Mechanik wechselt.

## Fitness Function

- **Paritäts-Proben** (Umsetzungs-Lauf): Message **mit** ID → 0; **ohne** ID → Exit ≠ 0 +
  `commit-untraceable`; `Merge `/`Revert ` → ausgenommen; Range mit ID-losem Commit → Exit
  ≠ 0; nicht auflösbare/leere Range → fail-closed (nicht still grün). CI-Range-Check + Hook
  + lokal grün gegen denselben Modul-Aufruf.
- `gate-consistency` grün (`trace-check`-Target bleibt in [`AGENTS.md` §4](../../../AGENTS.md#4-quality-gates)).

## Re-Evaluierungs-Trigger

- **`commits`-Modul-Vertrag ändert sich** (neue Flags/Semantik bei einem d-check-Bump):
  die Paritäts-Proben beim nächsten Pin-Bump erneut fahren.
- **Docker-loser Commit-Workflow** gewünscht (Hook ohne Container): dann eine schlanke
  lokale Vorprüfung erwägen — der CI-Range-Check bleibt die Wahrheit.

## Geschichte

| Datum | Ereignis |
|---|---|
| 2026-07-04 | Proposed — Entwurf; Paritäts-Proben gegen das abgelöste Skript bestanden (ID-Pflicht, Merge-/Revert-Ausnahme, fail-closed Range). |
| 2026-07-04 | Proposed → Accepted (Sign-off Auftraggeber per Merge-Wort). Ab jetzt immutable; Ablösung nur via Folge-ADR mit `Supersedes`. |
