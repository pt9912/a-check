#!/usr/bin/env bash
# verify-versions-kohaerent.sh — dieselbe Angabe an zwei Orten (slice-131).
#
# Gate-Schicht, nicht Verifikations-Schicht: die Frage ist eine ueber die
# Build-/CI-Konfiguration, keine ueber DoD oder Closure (Regelwerk Modul 11).
#
# ANLASS: BEO-026 bei 3x. Die Klasse hiess zuerst "der Tag-Kommentar wird nicht
# gegen den Digest geprueft" — das ist eine NETZ-Frage und beschreibt keinen der
# drei Vorfaelle. Jedes Mal lief DIESELBE Angabe an ZWEI Orten auseinander:
#
#   1. image-scan.yml trug `# v5.0.0` an einem Digest, den ci.yml `# v6.0.2` nennt
#   2. release.yml pinnte einen login-action-Digest zweimal, v4.2.0 und v3.6.0
#   3. Makefile fuehrt GO_VERSION 1.26.4, das Dockerfile-ARG 1.27.0
#
# Das ist eine Aussage ueber den Bestand gegen sich selbst und braucht kein
# Netz. Zwei Regeln:
#
#   (1) uses:-Kohaerenz — derselbe 40-stellige SHA traegt ueberall denselben
#       Tag-Kommentar (Vorfall 1+2)
#   (2) Pin-Deklarations-Kohaerenz — eine Versions-Variable, die Makefile UND
#       Dockerfile fuehren, hat an beiden Orten denselben Wert (Vorfall 3)
#
# NICHT GEPRUEFT, und das ist die Grenze: ob eine Angabe der WAHRHEIT
# entspricht. Zwei uebereinstimmend falsche Werte bleiben gruen. Der Sensor
# macht Divergenz sichtbar, nicht Unwahrheit — dafuer muesste er die Registry
# fragen, und `gates` ist hermetisch (ADR-0037 zieht diese Grenze fuer den
# einen Fall, in dem Netz der Zweck ist).
#
# KEINE FUEHRENDE SEITE: der Sensor erklaert weder Makefile noch Dockerfile zur
# Wahrheit. Er verlangt nur, dass sie sich nicht widersprechen — welcher Wert
# der richtige ist, entscheidet, wer die Aenderung macht.
set -euo pipefail
cd "$(dirname "$0")/.."

WF_DIR="${WF_DIR:-.github/workflows}"
MAKEFILE="${MAKEFILE_PATH:-Makefile}"
DOCKERFILE="${DOCKERFILE_PATH:-Dockerfile}"

# ---------------------------------------------------------------- Regel (1) --
# Ausgabe je Fund: "<sha> <tag> <datei>:<zeile>". Der Tag-Kommentar ist das
# erste Wort nach dem `#`; alles dahinter ist Prosa und zaehlt nicht mit.
uses_paare() {
  [ -d "$WF_DIR" ] || return 0
  grep -rHn -E 'uses:[[:space:]]*[^[:space:]@]+@[0-9a-f]{40}[[:space:]]*#' "$WF_DIR" 2>/dev/null \
  | sed -E 's/^([^:]+):([0-9]+):.*@([0-9a-f]{40})[[:space:]]*#[[:space:]]*([^[:space:]]+).*$/\3 \4 \1:\2/' \
  | grep -E '^[0-9a-f]{40} ' || true
}

check_uses() {
  local fail=0 sha tags n
  # Nach SHA gruppieren; mehr als ein DISTINKTER Tag zum selben SHA ist der Befund.
  for sha in $(uses_paare | awk '{print $1}' | sort -u); do
    tags="$(uses_paare | awk -v s="$sha" '$1==s {print $2}' | sort -u)"
    n="$(printf '%s\n' "$tags" | grep -c . || true)"
    [ "$n" -le 1 ] && continue
    echo "versions-kohaerenz: FAIL — ein SHA, mehrere Tag-Kommentare: ${sha:0:12}…" >&2
    uses_paare | awk -v s="$sha" '$1==s {printf "    %s  %s\n", $3, $2}' | sort >&2
    fail=1
  done
  return "$fail"
}

# ---------------------------------------------------------------- Regel (2) --
# Makefile: "NAME ?= wert" (nur diese Form — ":=" ist abgeleitet, nicht
# deklariert, und "=" ohne "?" kommt im Bestand nicht als Pin-Variable vor).
mk_vars() {
  [ -f "$MAKEFILE" ] || return 0
  sed -nE 's/^([A-Z][A-Z0-9_]*)[[:space:]]*\?=[[:space:]]*([^[:space:]#]+).*$/\1 \2/p' "$MAKEFILE" || true
}

# Dockerfile: "ARG NAME=wert". Ein Name kann in mehreren Stages stehen; jeder
# Wert wird verglichen, denn zwei Stages mit verschiedenen Defaults sind
# dieselbe Divergenz, nur eine Datei weiter innen.
df_args() {
  [ -f "$DOCKERFILE" ] || return 0
  sed -nE 's/^[[:space:]]*ARG[[:space:]]+([A-Z][A-Z0-9_]*)=([^[:space:]#]+).*$/\1 \2/p' "$DOCKERFILE" || true
}

check_pins() {
  local fail=0 name mkval dfvals v
  while read -r name mkval; do
    [ -n "$name" ] || continue
    dfvals="$(df_args | awk -v n="$name" '$1==n {print $2}' | sort -u)"
    [ -n "$dfvals" ] || continue          # nur im Makefile: keine zweite Deklaration, nichts zu vergleichen
    while read -r v; do
      [ -n "$v" ] || continue
      [ "$v" = "$mkval" ] && continue
      echo "versions-kohaerenz: FAIL — $name: $MAKEFILE sagt '$mkval', $DOCKERFILE sagt '$v'" >&2
      fail=1
    done <<< "$dfvals"
  done <<< "$(mk_vars)"
  return "$fail"
}

# ----------------------------------------------------------------- Selbsttest --
# Beide Richtungen JEDER Regel. Die gute Fixture ist die wichtigere: ein Sensor,
# der immer feuert, ist von einem korrekten nicht zu unterscheiden — dieselbe
# Ueberlegung wie in verify-risiko-ausgaenge.sh (slice-129).
self_test() {
  local tmp WF_SAVE MK_SAVE DF_SAVE
  tmp="$(mktemp -d)"
  WF_SAVE="$WF_DIR"; MK_SAVE="$MAKEFILE"; DF_SAVE="$DOCKERFILE"
  local sha_a='650006c6eb7dba73a995cc03b0b2d7f5ca915bee'
  local sha_b='de0fac2e4500dabe0009e67214ff5f5447ce83dd0'

  mkdir -p "$tmp/wf"
  WF_DIR="$tmp/wf"

  # (1a) ein SHA, zwei Tags -> muss feuern
  printf '        uses: a/b@%s # v4.2.0\n        uses: a/b@%s # v3.6.0\n' "$sha_a" "$sha_a" > "$tmp/wf/x.yml"
  if check_uses >/dev/null 2>&1; then
    echo "verify-versions-kohaerent: Selbsttest FEHLGESCHLAGEN — ein SHA mit zwei Tags nicht erkannt" >&2
    WF_DIR="$WF_SAVE"; rm -rf "$tmp"; exit 2
  fi

  # (1b) ein SHA, ein Tag, mehrfach + ein zweiter SHA -> muss schweigen
  printf '        uses: a/b@%s # v4.2.0\n        uses: a/b@%s # v4.2.0\n        uses: c/d@%s # v6.0.2\n' \
    "$sha_a" "$sha_a" "$sha_b" > "$tmp/wf/x.yml"
  if ! check_uses >/dev/null 2>&1; then
    echo "verify-versions-kohaerent: Selbsttest FEHLGESCHLAGEN — uebereinstimmende Tags als Befund gemeldet" >&2
    WF_DIR="$WF_SAVE"; rm -rf "$tmp"; exit 2
  fi
  WF_DIR="$WF_SAVE"

  MAKEFILE="$tmp/Makefile"; DOCKERFILE="$tmp/Dockerfile"

  # (2a) derselbe Name, zwei Werte -> muss feuern
  printf 'GO_VERSION ?= 1.26.4\nIMAGE ?= a-check\n' > "$tmp/Makefile"
  printf 'ARG GO_VERSION=1.27.0\n' > "$tmp/Dockerfile"
  if check_pins >/dev/null 2>&1; then
    echo "verify-versions-kohaerent: Selbsttest FEHLGESCHLAGEN — divergente Pin-Deklaration nicht erkannt" >&2
    MAKEFILE="$MK_SAVE"; DOCKERFILE="$DF_SAVE"; rm -rf "$tmp"; exit 2
  fi

  # (2b) gleicher Wert, plus eine Variable OHNE Gegenstueck -> muss schweigen
  printf 'GO_VERSION ?= 1.27.0\nTHRESHOLD ?= 90\n' > "$tmp/Makefile"
  printf 'ARG GO_VERSION=1.27.0\n' > "$tmp/Dockerfile"
  if ! check_pins >/dev/null 2>&1; then
    echo "verify-versions-kohaerent: Selbsttest FEHLGESCHLAGEN — uebereinstimmende Pins als Befund gemeldet" >&2
    MAKEFILE="$MK_SAVE"; DOCKERFILE="$DF_SAVE"; rm -rf "$tmp"; exit 2
  fi

  MAKEFILE="$MK_SAVE"; DOCKERFILE="$DF_SAVE"
  rm -rf "$tmp"
}

self_test

fail=0
check_uses || fail=1
check_pins || fail=1
if [ "$fail" -ne 0 ]; then
  echo "versions-kohaerenz: FEHLGESCHLAGEN — dieselbe Angabe steht an zwei Orten verschieden (BEO-026)." >&2
  exit 1
fi

n_uses="$(uses_paare | awk '{print $1}' | sort -u | grep -c . || true)"
n_pins="$(mk_vars | while read -r n _; do df_args | awk -v x="$n" '$1==x {print x}'; done | sort -u | grep -c . || true)"
echo "versions-kohaerenz ok: $n_uses gepinnte SHA(s) mit einheitlichem Tag-Kommentar, $n_pins doppelt deklarierte Pin-Variable(n) einig (Selbsttest gefeuert)."
echo "  NICHT geprueft: ob eine Angabe WAHR ist — zwei gleich falsche Werte bleiben gruen (Netz, ADR-0037)."
