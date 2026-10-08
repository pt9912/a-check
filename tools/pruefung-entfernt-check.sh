#!/usr/bin/env bash
# pruefung-entfernt-check.sh — Sensor zur Regel „Entfernte Prüfungen" (AGENTS.md §5,
# harness/rules/entfernte-pruefungen.md), slice-222. Antwort auf
# BEO-GATE/umbau-verliert-pruefung-still bei 3× (slice-217, slice-218, slice-221).
#
# REGEL. Entfernt ein Commit unter tools/ oder .github/workflows/ eine
# FEHLERPUNKT-Zeile — `fail "`, `::error::`, `exit 1`/`exit 2`, `probe "`,
# `assert ` — und fuegt er sie nicht wortgleich (nach Trim) wieder ein, traegt
# seine Message eine Zeile `Entfernte-Pruefungen: <Begruendung>`.
#
# WAS DER SENSOR LEISTET. Er macht den Verlust beim Commit SICHTBAR: er druckt
# die entfernten Zeilen und verlangt eine Begruendung. Ob eine entfernte Pruefung
# VERLOREN oder ERSETZT ist, entscheidet er nicht — das ist ein Urteil des
# Autors und des Reviews. Gemessen im Bestand (slice-222 §1b): 12 von 94
# Werkzeug-Commits haetten die Zeile gebraucht, darunter alle drei Belege.
#
# GRENZEN (am Satz, Mess-Regel 5): geprueft werden nur tools/ und
# .github/workflows/ und nur die fuenf Muster oben; eine Pruefung in anderer
# Form (etwa ein `return 1` ohne Meldung, ein Go-Testfall) sieht er nicht.
# Verschoben heisst: dieselbe Zeile steht im selben Commit wieder da —
# umformuliert zaehlt als entfernt.
#
# GRANDFATHERING OHNE STICHTAGS-HASH. Jeder Commit wird an der Fassung gemessen,
# die zu SEINEM Zeitpunkt galt: traegt AGENTS.md im Commit den Regel-Anker
# noch nicht, wird er uebersprungen (Muster von commit-scope-check.sh).
#
# Aufruf: make pruefung-entfernt-check [MSGFILE=<datei>] [RANGE=<a>..<b>]
#         (Default: HEAD~1..HEAD); --selftest prueft die Auswertung.
set -euo pipefail
cd "$(dirname "$0")/.."

MARKER='Entfernte-Pruefungen:'           # Regel-Anker in AGENTS.md und Trailer
PFADE=(tools .github/workflows)
MUSTER='fail "|::error::|exit [12]|probe "|assert '

# Diff (stdin) -> entfernte, nicht verschobene Fehlerpunkt-Zeilen (getrimmt,
# eindeutig). Nie ein Abbruch: keine Treffer heisst leere Ausgabe.
entfernte_fehlerpunkte() {
  local d rem add
  d="$(cat)"
  rem="$(printf '%s\n' "$d" | grep -E '^-[^-]' | sed 's/^-//; s/^[[:space:]]*//; s/[[:space:]]*$//' \
         | grep -E "$MUSTER" | sort -u || true)"
  [ -n "$rem" ] || return 0
  add="$(printf '%s\n' "$d" | grep -E '^\+[^+]' | sed 's/^+//; s/^[[:space:]]*//; s/[[:space:]]*$//' \
         | sort -u || true)"
  comm -23 <(printf '%s\n' "$rem") <(printf '%s\n' "$add") | grep . || true
}

# Message (stdin) -> 0, wenn sie eine Trailer-Zeile mit nicht leerer Begruendung traegt.
hat_begruendung() {
  grep -qE "^${MARKER}[[:space:]]*[^[:space:]]" || return 1
}

melde() {  # $1 = Bezeichnung, $2 = Zeilen
  echo "$1: entfernt Fehlerpunkt-Zeilen, die im selben Commit nicht wieder stehen:"
  printf '%s\n' "$2" | sed 's/^/    - /'
  echo "    -> Message-Zeile '${MARKER} <Grund>' ergaenzen: ersetzt durch …, oder entfaellt weil …"
  echo "       (Regel: harness/rules/entfernte-pruefungen.md)"
}

regel_galt() {  # $1 = sha
  git show "$1:AGENTS.md" 2>/dev/null | grep -qF "$MARKER"
}

check_commit() {  # $1 = sha
  local sha="$1" weg
  regel_galt "$sha" || return 0          # grandfathered
  weg="$(git show --format='' "$sha" -- "${PFADE[@]}" | entfernte_fehlerpunkte)"
  [ -n "$weg" ] || return 0
  git log -1 --format=%B "$sha" | hat_begruendung && return 0
  melde "$(git log -1 --format='%h %s' "$sha")" "$weg"
  return 1
}

check_pending() {  # $1 = Message-Datei
  local weg
  [ -f "$1" ] || { echo "pruefung-entfernt-check: FAIL — Message-Datei '$1' nicht lesbar." >&2; return 1; }
  grep -qF "$MARKER" AGENTS.md || return 0
  weg="$(git diff --cached -- "${PFADE[@]}" | entfernte_fehlerpunkte)"
  [ -n "$weg" ] || return 0
  hat_begruendung <"$1" && return 0
  melde "Pending-Commit" "$weg" >&2
  echo "    Der Commit ist NICHT entstanden." >&2
  return 1
}

self_test() {
  local fails=0
  t() {  # name erwartet ist
    if [ "$2" = "$3" ]; then printf '  ok   %-44s %s\n' "$1" "$2"
    else printf '  FAIL %-44s erwartet %s, war: %s\n' "$1" "$2" "$3"; fails=$((fails + 1)); fi
  }
  n() { entfernte_fehlerpunkte | grep -c . || true; }
  t "Fehlerpunkt entfernt"            1 "$(printf -- '-  [ x ] || fail "kaputt"\n' | n)"
  t "verschoben (wortgleich wieder da)" 0 "$(printf -- '-  [ x ] || fail "kaputt"\n+      [ x ] || fail "kaputt"\n' | n)"
  t "umformuliert zaehlt als entfernt" 1 "$(printf -- '-  fail "alt"\n+  fail "neu"\n' | n)"
  t "keine Fehlerpunkt-Zeile"         0 "$(printf -- '-  echo hallo\n' | n)"
  t "::error:: und exit 1"            2 "$(printf -- '-  echo "::error::x"\n-  exit 1\n' | n)"
  t "probe-Zeile"                     1 "$(printf -- '-  probe "Marker" x 0\n' | n)"
  # Die Kopfzeile eines Diffs (---) ist keine entfernte Zeile.
  t "Diff-Kopf --- zaehlt nicht"      0 "$(printf -- '--- a/tools/x.sh\n+++ b/tools/x.sh\n' | n)"
  # Exit 0 ist kein Fehlerpunkt.
  t "exit 0 zaehlt nicht"             0 "$(printf -- '-  exit 0\n' | n)"
  b() { if printf '%s\n' "$1" | hat_begruendung; then echo ja; else echo nein; fi; }
  t "Trailer mit Grund"               ja   "$(b "$(printf 'x\n\nEntfernte-Pruefungen: ersetzt durch y\n')")"
  t "Trailer leer"                    nein "$(b "$(printf 'x\n\nEntfernte-Pruefungen:   \n')")"
  t "Trailer nur im Fliesstext"       nein "$(b 'siehe Entfernte-Pruefungen: im Text')"
  t "kein Trailer"                    nein "$(b 'feat: x')"
  grep -qF "$MARKER" AGENTS.md
  t "Regel-Anker in AGENTS.md"        0 "$?"

  # Integration gegen ein Wegwerf-Repo: Range-Modus, beide Richtungen.
  local repo; repo="$(mktemp -d)"
  (
    set -e
    cd "$repo"; git init -q; git config user.email t@t; git config user.name t
    mkdir tools; printf 'Regel %s\n' "$MARKER" >AGENTS.md
    printf '[ x ] || fail "a"\n[ y ] || fail "b"\n' >tools/x.sh
    git add -A; git commit -qm init
    printf '[ x ] || fail "a"\n' >tools/x.sh; git commit -qam 'ohne Grund'
    printf '' >tools/x.sh; git commit -qam "$(printf 'mit Grund\n\nEntfernte-Pruefungen: a entfaellt, weil z\n')"
  ) >/dev/null
  local rc1 rc2
  (cd "$repo" && check_commit HEAD~1 >/dev/null) && rc1=0 || rc1=1
  (cd "$repo" && check_commit HEAD >/dev/null) && rc2=0 || rc2=1
  rm -rf "$repo"
  t "Repo: entfernt ohne Grund -> rot" 1 "$rc1"
  t "Repo: entfernt mit Grund -> gruen" 0 "$rc2"
  echo "== Fehlschlaege: $fails"
  [ "$fails" -eq 0 ]
}

if [ "${1:-}" = "--selftest" ]; then
  echo "pruefung-entfernt-check: Selbsttest"
  self_test; exit $?
fi

if [ -n "${MSGFILE:-}" ]; then
  check_pending "$MSGFILE"; rc=$?
  [ "$rc" -eq 0 ] && echo "pruefung-entfernt-check ok: Pending-Commit gegen den Index geprueft."
  exit "$rc"
fi

RANGE="${RANGE:-HEAD~1..HEAD}"
if ! shas="$(git rev-list "$RANGE" 2>/dev/null)"; then
  echo "pruefung-entfernt-check: FAIL — Range '$RANGE' nicht aufloesbar." >&2
  exit 2
fi
fail=0; n=0
for sha in $shas; do
  n=$((n + 1))
  check_commit "$sha" || fail=1
done
if [ "$fail" -ne 0 ]; then
  echo "pruefung-entfernt-check: FAIL — mindestens ein Commit entfernt Fehlerpunkte ohne Begruendung." >&2
  exit 1
fi
echo "pruefung-entfernt-check ok: ${n} Commit(s) in ${RANGE} geprueft."
