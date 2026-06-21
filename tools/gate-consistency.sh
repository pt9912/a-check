#!/usr/bin/env bash
# gate-consistency.sh — Meta-Gate gegen Harness-Lügen (Regelwerk Modul 13 /
# §Durchsetzungsschicht; Stack-Vorbild d-check tools/gate-consistency.sh):
#
#   (1) Jedes in AGENTS.md §4 bzw. harness/README.md §Sensors als
#       Tabellenzeile dokumentierte `make`-Target existiert real (Makefile
#       oder das includebare d-check.mk) — kein halluziniertes Gate.
#   (2) Jedes reale Gate-Target (Makefile + d-check.mk, ohne die Utility-
#       Targets help/build) ist in AGENTS.md §4 gelistet — AGENTS' eigene
#       Zusage „Nur hier gelistete Targets existieren im Makefile".
#   (3) Die modules-Liste der .a-check-Doku-Konfig (.d-check.yml) trägt die
#       aktiven Module (links/anchors/ids/matrix) und NICHT external — sonst
#       verliert der netzlose doc-check still seine Beweis-Aussage (AC-QA-02).
#
# Vor der echten Prüfung läuft ein Selbsttest: ein Phantom-Target muss das
# Gate nachweislich feuern lassen.
set -euo pipefail
cd "$(dirname "$0")/.."

# Utility-Targets: keine Gates, müssen nicht in AGENTS §4 stehen.
UTILITY_TARGETS='help build compile hooks'

# Dokumentierte Targets: alle `make <name>`-Tokens in Tabellenzeilen.
doc_targets() {
  grep -E '^\|' "$1" | grep -oE '`make [a-z][a-z0-9_-]*`' \
    | sed -E 's/`make ([a-z0-9_-]+)`/\1/' | sort -u
}

# Reale Targets aus allen Makefile-Fragmenten (Makefile + includebare *.mk):
# Regelzeilen am Zeilenanfang, auch Mehrfach-Targets (`a b: dep`).
# Zuweisungen (`X := y`, `X ?= y`) und `.PHONY`/`.DEFAULT_GOAL` sind
# ausgeschlossen (führendes `.` bzw. `=` nach dem Doppelpunkt).
makefile_targets() {
  cat "$@" | grep -oE '^[a-zA-Z][a-zA-Z0-9 _-]*:([^=]|$)' \
    | sed 's/:.*//' | tr ' ' '\n' | sed '/^$/d' | sort -u
}

# nutzt globales MK_TARGETS
check_documented_exist() {
  local fail=0 doc t
  for doc in "$@"; do
    while IFS= read -r t; do
      [ -z "$t" ] && continue
      if ! grep -qx "$t" <<<"$MK_TARGETS"; then
        echo "gate-consistency: FAIL — $doc dokumentiert 'make $t', das aber kein reales Target ist" >&2
        fail=1
      fi
    done <<<"$(doc_targets "$doc")"
  done
  return "$fail"
}

self_test() {
  local tmp
  tmp="$(mktemp -d)"
  printf '| `make phantom-target` | x |\n' > "$tmp/doc.md"
  printf 'echtes-target zweites-target: dep\n\ttrue\nVAR := x\n' > "$tmp/Makefile"
  MK_TARGETS="$(makefile_targets "$tmp/Makefile")"
  if check_documented_exist "$tmp/doc.md" 2>/dev/null; then
    echo "gate-consistency: Selbsttest FEHLGESCHLAGEN — Phantom-Target nicht erkannt" >&2
    rm -rf "$tmp"
    exit 2
  fi
  if [ "$(makefile_targets "$tmp/Makefile" | wc -l)" -ne 2 ]; then
    echo "gate-consistency: Selbsttest FEHLGESCHLAGEN — Makefile-Parser (Mehrfach-Targets/Zuweisungen)" >&2
    rm -rf "$tmp"
    exit 2
  fi
  rm -rf "$tmp"
}

self_test
fail=0
MK_TARGETS="$(makefile_targets Makefile d-check.mk)"

# (1) Doku → real
check_documented_exist AGENTS.md harness/README.md || fail=1

# (2) real → AGENTS §4 (ohne Utility-Targets)
agents_targets="$(doc_targets AGENTS.md)"
while IFS= read -r t; do
  [ -z "$t" ] && continue
  if grep -qw "$t" <<<"$UTILITY_TARGETS"; then
    continue
  fi
  if ! grep -qx "$t" <<<"$agents_targets"; then
    echo "gate-consistency: FAIL — reales Target '$t' fehlt in AGENTS.md §4" >&2
    fail=1
  fi
done <<<"$MK_TARGETS"

# (3) .d-check.yml-Modulliste des netzlosen doc-check (AC-QA-02)
modules_line="$(grep -E '^modules:' .d-check.yml || true)"
for m in links anchors ids matrix; do
  if [[ "$modules_line" != *"$m"* ]]; then
    echo "gate-consistency: FAIL — .d-check.yml modules ohne '$m'; der netzlose doc-check beweist AC-QA-02 nur mit den aktiven Modulen" >&2
    fail=1
  fi
done
if [[ "$modules_line" == *external* ]]; then
  echo "gate-consistency: FAIL — .d-check.yml aktiviert external; das doc-check-Gate muss netzlos bleiben (AC-QA-02)" >&2
  fail=1
fi

if [ "$fail" -ne 0 ]; then
  exit 1
fi
echo "gate-consistency ok: Doku ↔ Makefile konsistent, .d-check.yml-Module intakt (Selbsttest gefeuert)."
