#!/usr/bin/env bash
# verify-review-haken.sh — der Review-DoD-Haken ist an die Existenz des Reports
# gebunden (AGENTS.md §5; hervorgegangen aus
# BEO-GATE/attestierung-vor-dem-vorgang, 5x: slice-169, slice-197, slice-200,
# slice-191, slice-205).
#
# WAS HIER GILT. Ein abgehakter DoD-Punkt "Unabhaengiger Review" in einem
# in-progress/-Slice attestiert einen Vorgang, dessen Beleg unter docs/reviews/
# existieren muss — dieselbe Phrase, die das Modul `reviews` als Opt-in erwartet
# (MR-019), in derselben exakten Wortform wie der Gate-Index (structure (1)).
#
# WO DER SINN GRENZE. Das Modul `reviews` scannt genau EIN done-dir
# (d-check v0.79.0: opt-in ueber DoneDir) — die attestierte Gestalt entsteht
# aber im in-progress/-Stand, bevor der Slice nach done/ wandert. Darum eigene
# Schalung in der Verifikations-Schicht statt eines CR an das Fremdwerkzeug
# (slice-204 §6: "forbid-pattern prueft Text, nicht Datei-Existenz"; structure
# hat keine Datei-Existenz-Bedingung).
#
# NICHT geprueft (ehrliche Grenze, AC-QA-02):
# - Der DATEINAME-Match ist Form: ein Report mit der Kennung im Namen kann
#   leer oder zum falschen Lauf sein; ob er TRAEGT, ist Urteil (Modul 10).
# - Nur ABGEHAKTE Zeilen sind eine Zusage; ein unabgehakter DoD-Punkt ist
#   keine (Opt-in pro Slice, MR-019). open//next/-Slices scanned dieses
#   Skript nicht — dort deckt die structure-Bedingung 7 die unchecked-Haelfte
#   (.d-check.yml, seit slice-202), done/ das Modul `reviews` (doc-reviews).
# - Ein Report unter docs/reviews/ kann von einem FRUEHEREN Lauf desselben
#   Slice stammen (Folge-Review); die Kennung im Dateinamen ist die Bindung,
#   nicht das Datum.
set -euo pipefail
cd "$(dirname "$0")/.."

PROGRESS_DIR="docs/plan/planning/in-progress"
REVIEWS_DIR="docs/reviews"
PHRASE="Unabhängiger Review"

slice_num() {  # $1 = Pfad -> Nummer ohne fuehrende Nullen, leer wenn keine
  basename "$1" | sed -nE 's/^slice-0*([0-9]+)-.*/\1/p'
}

hat_report() {  # $1 = Slice-Nummer -> 0, wenn ein Report-Dateiname sie traegt
  local n="$1"
  ls "$REVIEWS_DIR" 2>/dev/null | grep -qEi "slice[-_]?0*${n}([^0-9]|$)"
}

check_file() {  # $1 = Datei; Befunde auf stdout, Rueckgabe 1 bei Befund
  local f="$1" num
  num="$(slice_num "$f")"
  [ -n "$num" ] || return 0
  if grep -E '^[[:space:]]*- \[x\]' "$f" | grep -qF "$PHRASE"; then
    if ! hat_report "$num"; then
      echo "$f: Review-DoD abgehakt ($PHRASE), aber kein Report mit der Kennung slice-$num unter $REVIEWS_DIR/"
      return 1
    fi
  fi
  return 0
}

# Der Lauf scannt NUR in-progress/ — die Abgrenzung nach done/ (Modul reviews)
# und open//next/ (structure (7)) ist die des Scans, nicht einer Ausnahme.
run() {
  local fail=0 count=0 f
  for f in "$PROGRESS_DIR"/slice-*.md; do
    [ -e "$f" ] || continue
    count=$((count + 1))
    if ! out="$(check_file "$f")"; then
      printf '%s\n' "$out" >&2
      fail=1
    elif [ -n "$out" ]; then
      printf '%s\n' "$out" >&2
      fail=1
    fi
  done
  if [ "$fail" -ne 0 ]; then
    echo "verify-review-haken: FAIL — Review-DoD-Haken ohne Report (AGENTS.md §5; BEO-GATE/attestierung-vor-dem-vorgang)." >&2
    exit 1
  fi
  # Null Claims sind ein NORMALZUSTAND: in-progress/ kann leer sein. Gezaehlt
  # und gemeldet wird trotzdem — ein Ueberspringen, das niemand sieht, waere
  # von einer Pruefung, die nichts findet, nicht zu unterscheiden.
  printf "verify-review-haken ok: %s in-progress-Slice(s) geprueft, jeder abgehakte Review-DoD mit Report" "$count"
  echo " (Selbsttest gefeuert)."
}

# Selbsttest: je eine Fixture pro Richtung muss feuern, die guten schweigen
# (Modul 11: eine Probe belegt erst, wenn sie ROT war — die rote Haelfte steht
# zuerst). Die done/-Fixture praefiert die Abgrenzung: ein Haken dort ist
# Sache des Moduls `reviews`, nicht dieses Skripts.
self_test() {
  local tmp PD_SAVE RD_SAVE
  tmp="$(mktemp -d)"
  PD_SAVE="$PROGRESS_DIR"; RD_SAVE="$REVIEWS_DIR"

  mkdir -p "$tmp/pd" "$tmp/rd" "$tmp/done"
  printf '# slice-999\n- [x] Unabhängiger Review, Report unter docs/reviews/.\n' \
    > "$tmp/pd/slice-999-rot.md"
  printf '# slice-998\n- [x] Unabhängiger Review, Report unter docs/reviews/.\n' \
    > "$tmp/pd/slice-998-gruen.md"
  printf 'report\n' > "$tmp/rd/2026-01-01-slice-998-gruen.md"
  printf '# slice-997\n- [ ] Unabhängiger Review, Report unter docs/reviews/.\n' \
    > "$tmp/pd/slice-997-opt-in.md"
  printf '# slice-205\n- [x] Unabhängiger Review, Report unter docs/reviews/.\n' \
    > "$tmp/pd/slice-205-teils.md"
  printf 'report\n' > "$tmp/rd/2026-01-01-slice-2050-falsch.md"
  printf '# slice-996\n- [x] Unabhängiger Review, Report unter docs/reviews/.\n' \
    > "$tmp/done/slice-996-done.md"

  PROGRESS_DIR="$tmp/pd"; REVIEWS_DIR="$tmp/rd"

  # ROT: Haken ohne Report.
  if check_file "$tmp/pd/slice-999-rot.md" >/dev/null 2>&1; then
    echo "verify-review-haken: Selbsttest FEHLGESCHLAGEN — Haken ohne Report blieb gruen" >&2
    PROGRESS_DIR="$PD_SAVE"; REVIEWS_DIR="$RD_SAVE"; rm -rf "$tmp"; exit 2
  fi
  echo "selftest: Haken ohne Report -> rot (gefiret)"

  # GRUEN: Haken mit Report.
  if ! check_file "$tmp/pd/slice-998-gruen.md" >/dev/null 2>&1; then
    echo "verify-review-haken: Selbsttest FEHLGESCHLAGEN — Haken mit Report meldete rot" >&2
    PROGRESS_DIR="$PD_SAVE"; REVIEWS_DIR="$RD_SAVE"; rm -rf "$tmp"; exit 2
  fi

  # GRUEN: unabgehakter Punkt ist keine Zusage (Opt-in, MR-019).
  if ! check_file "$tmp/pd/slice-997-opt-in.md" >/dev/null 2>&1; then
    echo "verify-review-haken: Selbsttest FEHLGESCHLAGEN — unabgehakter Punkt meldete rot" >&2
    PROGRESS_DIR="$PD_SAVE"; REVIEWS_DIR="$RD_SAVE"; rm -rf "$tmp"; exit 2
  fi

  # ROT: slice-2050 ist NICHT slice-205 (Kennungs-Grenze des Dateinamens).
  if check_file "$tmp/pd/slice-205-teils.md" >/dev/null 2>&1; then
    echo "verify-review-haken: Selbsttest FEHLGESCHLAGEN — slice-2050 bediente slice-205" >&2
    PROGRESS_DIR="$PD_SAVE"; REVIEWS_DIR="$RD_SAVE"; rm -rf "$tmp"; exit 2
  fi
  echo "selftest: Kennungs-Grenze slice-205 vs slice-2050 -> getrennt"

  # ABGRENZUNG: ein Haken in done/ ist hier unsichtbar (Modul reviews).
  PROGRESS_DIR="$tmp/leer"; mkdir -p "$PROGRESS_DIR"
  PROGRESS_DIR="$tmp/leer"
  if ! run >/dev/null 2>&1; then
    echo "verify-review-haken: Selbsttest FEHLGESCHLAGEN — leerer in-progress-Stand meldete rot" >&2
    PROGRESS_DIR="$PD_SAVE"; REVIEWS_DIR="$RD_SAVE"; rm -rf "$tmp"; exit 2
  fi
  PROGRESS_DIR="$PD_SAVE"; REVIEWS_DIR="$RD_SAVE"; rm -rf "$tmp"
}

self_test

run
