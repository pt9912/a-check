#!/usr/bin/env bash
# verify-ac-form.sh — Form neuer Akzeptanzkriterien (slice-054),
# Fund B-15 aus slice-048.
#
# Die Baseline verlangt je funktionalem AC drei Pfade im Given/When/Then-Stil
# (Happy · Boundary · Negative) plus einen Out-of-Scope-Block.
#
# GRANDFATHERING — namentlich, nicht per Heuristik. Die 19 bei Einfuehrung
# bestehenden AC-IDs stehen unten. Sie werden NICHT umgeschrieben: sie sind
# vertraglich abnahmebindend, ihr Inhalt traegt Rand- und Negativfaelle bereits
# in Prosa, und ein Umbau traefe die Form statt der Substanz. Dieselbe Mechanik
# benutzt die Baseline fuer die Referenz-Richtungs-Regel ("prueft nur ab
# Einfuehrung neu").
#
# Die Liste waechst NICHT mit: jede kuenftige AC-ID faellt automatisch unter die
# Regel. Genau das macht die Liste zum Beleg dafuer, dass die Regel ab jetzt
# greift — eine Datums- oder Nummern-Heuristik haette diese Eigenschaft nicht.
#
# NICHT geprueft (ehrliche Grenze, AC-QA-02): ob der Satz hinter "Boundary:"
# wirklich einen Randfall beschreibt. Das ist semantisch und bleibt Review-Sache;
# ein Regex darueber wuerde Form-Erfuellung mit Inhalt verwechseln.
set -euo pipefail
cd "$(dirname "$0")/.."

SPEC="spec/lastenheft.md"

GRANDFATHERED="
AC-FA-RULE-001 AC-FA-RULE-002 AC-FA-RULE-003 AC-FA-RULE-004 AC-FA-RULE-005
AC-FA-RULE-006 AC-FA-RULE-007 AC-FA-RULE-008 AC-FA-RULE-009 AC-FA-RULE-010
AC-FA-RULE-011 AC-FA-EXTRACT-001 AC-FA-CLI-001 AC-FA-CLI-002 AC-FA-CONF-001
AC-FA-DIST-001 AC-QA-01 AC-QA-02 AC-QA-03
"

is_grandfathered() {  # $1 = AC-ID
  grep -qw -- "$1" <<<"$GRANDFATHERED"
}

# Abschnitt einer AC: ab ihrer Ueberschrift bis zur naechsten ###-Ueberschrift.
ac_body() {  # $1 = Datei, $2 = AC-ID
  awk -v id="$2" '
    $0 ~ "^### " id " " { inblock=1; next }
    inblock && /^### / { inblock=0 }
    inblock { print }
  ' "$1"
}

check_ac() {  # $1 = Datei, $2 = AC-ID; Befunde auf stdout
  local body fail=0 pfad
  body="$(ac_body "$1" "$2")"
  for pfad in "Happy Path" "Boundary" "Negative" "Out-of-Scope"; do
    if ! printf '%s\n' "$body" | grep -qE "^\*\*$pfad:?\*\*"; then
      echo "$1: $2 ohne Pfad '**$pfad:**' (AC-Form, AGENTS.md §5)"
      fail=1
    fi
  done
  return "$fail"
}

self_test() {
  local tmp
  tmp="$(mktemp -d)"
  {
    echo '### AC-FA-TEST-001 — gut'
    echo '**Beschreibung:** x.'
    echo '**Happy Path:** Given a, when b, then c.'
    echo '**Boundary:** Given rand, when b, then d.'
    echo '**Negative:** Given falsch, when b, then Fehler.'
    echo '**Out-of-Scope:** nichts.'
    echo '### AC-FA-TEST-002 — unvollstaendig'
    echo '**Beschreibung:** y.'
    echo '**Happy Path:** Given a, when b, then c.'
    echo '**Out-of-Scope:** nichts.'
  } > "$tmp/spec.md"

  if ! check_ac "$tmp/spec.md" AC-FA-TEST-001 >/dev/null; then
    echo "verify-ac-form: Selbsttest FEHLGESCHLAGEN — vollstaendige AC beanstandet" >&2
    rm -rf "$tmp"; exit 2
  fi
  if check_ac "$tmp/spec.md" AC-FA-TEST-002 >/dev/null; then
    echo "verify-ac-form: Selbsttest FEHLGESCHLAGEN — fehlende Pfade nicht erkannt" >&2
    rm -rf "$tmp"; exit 2
  fi
  # Der Abschnitts-Schnitt muss an der naechsten Ueberschrift enden — sonst
  # wuerde AC-002 die Pfade von AC-001 "erben" und still gruen werden.
  if printf '%s\n' "$(ac_body "$tmp/spec.md" AC-FA-TEST-002)" | grep -q 'Boundary'; then
    echo "verify-ac-form: Selbsttest FEHLGESCHLAGEN — Abschnitts-Schnitt laeuft in die naechste AC" >&2
    rm -rf "$tmp"; exit 2
  fi
  # Grandfathering in beide Richtungen.
  if ! is_grandfathered AC-QA-01; then
    echo "verify-ac-form: Selbsttest FEHLGESCHLAGEN — Bestands-AC nicht grandfathered" >&2
    rm -rf "$tmp"; exit 2
  fi
  if is_grandfathered AC-FA-NEU-001; then
    echo "verify-ac-form: Selbsttest FEHLGESCHLAGEN — neue AC faelschlich grandfathered" >&2
    rm -rf "$tmp"; exit 2
  fi
  rm -rf "$tmp"
}

self_test

fail=0; geprueft=0; alt=0
while IFS= read -r id; do
  [ -z "$id" ] && continue
  if is_grandfathered "$id"; then
    alt=$((alt + 1)); continue
  fi
  geprueft=$((geprueft + 1))
  if ! out="$(check_ac "$SPEC" "$id")"; then
    printf '%s\n' "$out" >&2; fail=1
  fi
done <<<"$(grep -oE '^### AC-[A-Z0-9-]+' "$SPEC" | sed 's/^### //')"

if [ "$fail" -ne 0 ]; then
  echo "verify-ac-form: FAIL — neue Akzeptanzkriterien ohne die drei Pfade (AGENTS.md §5)." >&2
  exit 1
fi
echo "verify-ac-form ok: $geprueft neue AC(s) geprueft, $alt bei Einfuehrung bestehende (grandfathered)."
