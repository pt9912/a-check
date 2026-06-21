#!/usr/bin/env bash
# record-gates — Nachweis schreiben, dass `make gates` den aktuellen
# Arbeitsbaum-Inhalt abgedeckt hat. Läuft als letzter gates-Prerequisite
# (nur bei grünen Gates, Reihenfolge via .NOTPARALLEL). Der Stop-Hook
# vergleicht denselben Hash. Regelwerk §Durchsetzungsschicht (Handoff-Gate);
# Stack-Vorbild d-check.
set -euo pipefail
cd "$(git rev-parse --show-toplevel)"

mkdir -p .harness/state
bash tools/harness/working-tree-hash.sh > .harness/state/gates-passed.diffsha
echo "record-gates ok: Nachweis .harness/state/gates-passed.diffsha geschrieben."
