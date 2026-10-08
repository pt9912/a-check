**Stand:** verkörpert — liegt in `Makefile:pruefung-entfernt-check` (`tools/pruefung-entfernt-check.sh`, Regel [`harness/rules/entfernte-pruefungen.md`](../../../../../../harness/rules/entfernte-pruefungen.md), `seit slice-222`).

Ein Sensor, keine Prosa: Ein Commit, der unter `tools/` oder `.github/workflows/` eine
Fehlerpunkt-Zeile entfernt und nicht in derselben Datei wieder einfügt, braucht
`Entfernte-Pruefungen: <Grund>` — im `commit-msg`-Hook, in `make preflight` und im CI-Workflow.
Alle Belege hätte er beim Commit gezeigt. Ob eine entfernte Prüfung verloren oder ersetzt ist, bleibt
ein Urteil; der Sensor macht sie sichtbar.
