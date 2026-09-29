#!/usr/bin/env bash
# verify-trigger-audit.sh — die „dass"-Haelfte des Trigger-Audits fuer MR-Einträge
# (AGENTS.md §5; hervorgegangen aus BEO-HARNESS/mr-aufloesungs-trigger-ohne-waechter,
# 3x: slice-170, slice-197, slice-206).
#
# WAS HIER GILT. Jede Closure-Notiz eines Slice ab slice-208 traegt eine Zeile
# "Trigger-Audit der aktiven MR:" — den Beleg, dass die Closure die aktiven
# Eintraege aus harness/conventions.md angesehen hat. Das Feld
# "Aufloesungs-Trigger" selbst ist Prosa und maschinell nicht auswertbar; der
# Sensor prueft die SICHTUNG, nicht das Urteil.
#
# WO DER SINN GRENZE. Der Trigger-Audit des Baseline-Regelwerks (modul-06,
# Closure-Schritt 2) zaehlt Carveout, bootstrap-aware Gate und ADR auf; der
# MR-Eintrag steht nicht darin — darum eigene Schalung in der
# Verifikations-Schicht statt eines CR an das Fremdwerkzeug.
#
# NICHT geprueft (ehrliche Grenze, AC-QA-02):
# - Das URTEIL. Ob "0 offen" stimmt oder die genannten MR-IDs wirklich
#   abgelaufen sind, ist Prosa-Auswertung; der Sensor prueft die Form.
# - Slices VOR slice-208 sind grandfathered — ihre Closures entstanden, bevor
#   die Zusage stand (dieselbe Grenze wie verify-risiko-ausgaenge RISK_FROM).
# - wellenlose Closures in done/wellenlos/ ZAEHLEN mit (der Regelfall dieses
#   Repos); Welle-Closures aus done/welle-*/ werden nicht gescannt — deren
#   Form ist eine andere (modul-06 Closure-Schritt 2).
set -euo pipefail
cd "$(dirname "$0")/.."

DONE_DIR="docs/plan/planning/done"
CONVENTIONS="harness/conventions.md"
AUDIT_FROM=208
PHRASE="Trigger-Audit der aktiven MR:"

aktive_mr() {  # aktive MR-Nummern aus dem Konventions-Speicher, auf stdout
  sed -n '/^### Aktive Adaptionen/,/^### Aufgelöste/p' "$CONVENTIONS" \
    | grep -oE 'MR-[0-9]{3}' | sort -u
}

check_file() {  # $1 = Datei; Befunde auf stdout, Rueckgabe 1 bei Befund
  local f="$1" num
  num="$(basename "$f" | sed -nE 's/^slice-0*([0-9]+)-.*/\1/p')"
  [ -n "$num" ] || return 0
  [ "$num" -ge "$AUDIT_FROM" ] || return 0
  if ! grep -qF "$PHRASE" "$f"; then
    echo "$f: Closure-Notiz ohne $PHRASE-Zeile — die Sichtung der aktiven MR-Eintraege ist nicht belegt"
    return 1
  fi
  return 0
}

# Der Lauf scannt alle done-Slices ab AUDIT_FROM — wellenlose wie
# Welle-Slices in done/ selbst (nicht die Archiv-Stubs tieferer Ebenen).
run() {
  local fail=0 count=0 f
  for f in "$DONE_DIR"/slice-*.md; do
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
    echo "verify-trigger-audit: FAIL — Closure ohne Trigger-Audit-Zeile (AGENTS.md §5; BEO-HARNESS/mr-aufloesungs-trigger-ohne-waechter)." >&2
    exit 1
  fi
  local aktive
  aktive="$(aktive_mr | tr '\n' ' ')"
  # Die Leere Aktiven-Menge ist eine MELDUNG, kein stiller Durchlauf: sie
  # sagt, dass kein MR-Eintrag mehr auf eine Sichtung wartet.
  if [ -z "$(aktive_mr)" ]; then
    printf "verify-trigger-audit ok: %s Closure(s) geprueft — aktive MR-Menge LEER, nichts wartet auf eine Sichtung" "$count"
  else
    printf "verify-trigger-audit ok: %s Closure(s) geprueft — aktive MR: %s" "$count" "$(aktive_mr | tr '\n' ' ')"
  fi
  echo " (Selbsttest gefeuert)."
}

# Selbsttest: je eine Fixture pro Richtung muss feuern, die guten schweigen
# (Modul 11: die rote Haelfte zuerst). Die Fixture fuer die Leere
# Aktiven-Menge praefiert die Meldung statt stillen Durchlaufs.
self_test() {
  local tmp DONE_SAVE CONV_SAVE
  tmp="$(mktemp -d)"
  DONE_SAVE="$DONE_DIR"; CONV_SAVE="$CONVENTIONS"

  mkdir -p "$tmp/done"
  # ROT: Closure ab slice-208 ohne Audit-Zeile.
  cat > "$tmp/done/slice-990-rot.md" << 'EOF'
# slice-990
## 7. Closure-Notiz
geliefert.
EOF
  # GRUEN: Closure mit Audit-Zeile.
  cat > "$tmp/done/slice-991-gruen.md" << 'EOF'
# slice-991
## 7. Closure-Notiz
Trigger-Audit der aktiven MR: 0 offen.
EOF
  # GRANDFATHERED: Closure vor slice-208 ohne Audit-Zeile bleibt gruen.
  cat > "$tmp/done/slice-207-alt.md" << 'EOF'
# slice-207
## 7. Closure-Notiz
geliefert.
EOF

  DONE_DIR="$tmp/done"

  if check_file "$tmp/done/slice-990-rot.md" >/dev/null 2>&1; then
    echo "verify-trigger-audit: Selbsttest FEHLGESCHLAGEN — Closure ohne Audit-Zeile blieb gruen" >&2
    DONE_DIR="$PD_SAVE"; rm -rf "$tmp"; exit 2
  fi
  echo "selftest: Closure ohne Audit-Zeile -> rot (gefiret)"

  if ! check_file "$tmp/done/slice-991-gruen.md" >/dev/null 2>&1; then
    echo "verify-trigger-audit: Selbsttest FEHLGESCHLAGEN — Audit-Zeile meldete rot" >&2
    DONE_DIR="$PD_SAVE"; rm -rf "$tmp"; exit 2
  fi

  if ! check_file "$tmp/done/slice-207-alt.md" >/dev/null 2>&1; then
    echo "verify-trigger-audit: Selbsttest FEHLGESCHLAGEN — grandfathered Closure meldete rot" >&2
    DONE_DIR="$PD_SAVE"; rm -rf "$tmp"; exit 2
  fi
  echo "selftest: Grandfathering vor slice-208 -> gruen"

  # GRUEN GEMELDET: leere Aktiven-Menge — gemeldet, nicht still. Der rot
  # gepruefte Fall ist zuvor zu entfernen, sonst meldet run() sein eigenes
  # Fixture statt die leere Menge.
  rm "$tmp/done/slice-990-rot.md"
  CONVENTIONS="$tmp/leer.md"
  : > "$CONVENTIONS"
  out="$(run 2>&1)" || {
    echo "verify-trigger-audit: Selbsttest FEHLGESCHLAGEN — leere Aktiven-Menge meldete rot" >&2
    DONE_DIR="$PD_SAVE"; CONVENTIONS="$CONV_SAVE"; rm -rf "$tmp"; exit 2
  }
  case "$out" in
    *"MR-Menge LEER"*) echo "selftest: leere Aktiven-Menge -> gemeldet" ;;
    *) echo "verify-trigger-audit: Selbsttest FEHLGESCHLAGEN — leere Aktiven-Menge still" >&2
       DONE_DIR="$PD_SAVE"; CONVENTIONS="$CONV_SAVE"; rm -rf "$tmp"; exit 2 ;;
  esac

  DONE_DIR="$DONE_SAVE"; CONVENTIONS="$CONV_SAVE"; rm -rf "$tmp"
}

self_test

run
