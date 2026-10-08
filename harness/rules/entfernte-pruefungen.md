# Ein Commit, der eine Prüfung entfernt, sagt warum

*Aus [`AGENTS.md`](../../AGENTS.md) §5, Regel 18 — seit slice-222.*

**Entfernt ein Commit unter `tools/` oder `.github/workflows/` eine Fehlerpunkt-Zeile** —
`fail "`, `::error::`, `exit 1`/`exit 2`, `probe "`, `assert ` — und steht dieselbe Zeile im
selben Commit nicht wieder da, trägt seine Message eine Zeile

```text
Entfernte-Pruefungen: <ersetzt durch …> | <entfällt, weil …>
```

Geprüft von `make pruefung-entfernt-check`: im `commit-msg`-Hook gegen den Index (vor dem
Commit), in `make preflight` und im CI-Workflow über die Commit-Range. Der Sensor druckt die
entfernten Zeilen; **ob** die Prüfung verloren oder ersetzt ist, entscheidet er nicht — das ist das
Urteil des Autors, und der Review hält es gegen den Diff. Grenze: nur diese zwei Verzeichnisse und
diese fünf Muster; eine umformulierte Zeile zählt als entfernt.

**Warum.** Dreimal verlor ein Umbau eine bestehende Prüfung, ohne dass es jemand sagte — der
Abgleich der Index-Beschriftung (slice-217), die Prüfung von fünf OCI-Labels (slice-218), die
Anker-Probe eines Selbsttests (slice-221). Alle drei fing erst der unabhängige Review; die Tests
blieben grün, weil sie die weggefallene Eigenschaft nie einzeln gebrochen hatten. Gemessen hätte
der Sensor alle drei beim Commit gezeigt
([`BEO-GATE/umbau-verliert-pruefung-still`](../../docs/plan/planning/observations/BEO-GATE/umbau-verliert-pruefung-still/observation.md)).
