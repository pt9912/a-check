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
# .github/workflows/ und nur die fuenf Muster oben, ohne Kommentarzeilen; eine
# Pruefung in anderer Form (etwa ein `return 1` ohne Meldung, ein Go-Testfall)
# sieht er nicht. Verschoben heisst: dieselbe Zeile steht im selben Commit in
# DERSELBEN Datei wieder da (gezaehlt als Multimenge) — umformuliert oder in eine
# andere Datei verschoben zaehlt als entfernt. Ein nacktes `exit 1`, das in
# derselben Datei an anderer Stelle neu entsteht, gleicht ein entferntes aus.
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
MUSTER='fail "|::error::|exit [12]([^0-9]|$)|probe "|assert '

# Unified Diff (stdin) -> entfernte, nicht verschobene Fehlerpunkt-Zeilen
# (getrimmt), je Vorkommen eine Zeile. Gelesen wird hunkweise: Dateikoepfe
# (`diff --git`, `--- a/…`, `+++ b/…`) stehen VOR dem ersten `@@` und zaehlen nie;
# im Hunk ist `-` entfernt und `+` hinzugefuegt, auch wenn der Inhalt selbst mit
# `-` beginnt. Gegengerechnet wird je Datei und Zeile (Multimenge). Nie ein
# Abbruch: keine Treffer heisst leere Ausgabe.
entfernte_fehlerpunkte() {
  awk -v pat="$MUSTER" '
    function trim(s) { sub(/^[ \t]+/, "", s); sub(/[ \t]+$/, "", s); return s }
    /^diff --git / { inh = 0; f = ""; next }
    !inh && /^--- / { g = substr($0, 5); sub(/^a\//, "", g); if (g != "/dev/null") f = g; next }
    !inh && /^\+\+\+ / { g = substr($0, 5); sub(/^b\//, "", g); if (g != "/dev/null") f = g; next }
    /^@@/ { inh = 1; next }
    inh && /^-/ { l = trim(substr($0, 2)); rem[f SUBSEP l]++; next }
    inh && /^\+/ { l = trim(substr($0, 2)); add[f SUBSEP l]++; next }
    END {
      for (k in rem) {
        split(k, a, SUBSEP); l = a[2]
        if (l ~ /^#/ || l !~ pat) continue
        for (i = add[k]; i < rem[k]; i++) print l
      }
    }'
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
  weg="$(git show --no-color --format='' "$sha" -- "${PFADE[@]}" | entfernte_fehlerpunkte | sort -u)"
  [ -n "$weg" ] || return 0
  git log -1 --format=%B "$sha" | hat_begruendung && return 0
  melde "$(git log -1 --format='%h %s' "$sha")" "$weg"
  return 1
}

check_pending() {  # $1 = Message-Datei
  local weg
  [ -f "$1" ] || { echo "pruefung-entfernt-check: FAIL — Message-Datei '$1' nicht lesbar." >&2; return 1; }
  grep -qF "$MARKER" AGENTS.md || return 0
  local diff
  if ! diff="$(git diff --cached --no-color -- "${PFADE[@]}")"; then
    echo "pruefung-entfernt-check: FAIL — git diff --cached gescheitert; der Index ist nicht pruefbar." >&2
    return 2
  fi
  weg="$(printf '%s\n' "$diff" | entfernte_fehlerpunkte | sort -u)"
  [ -n "$weg" ] || return 0
  hat_begruendung <"$1" && return 0
  melde "Pending-Commit" "$weg" >&2
  echo "    Der Commit ist NICHT entstanden." >&2
  return 1
}

self_test() {
  local fails=0
  t() {  # name erwartet ist
    if [ "$2" = "$3" ]; then printf '  ok   %-46s %s\n' "$1" "$2"
    else printf '  FAIL %-46s erwartet %s, war: %s\n' "$1" "$2" "$3"; fails=$((fails + 1)); fi
  }
  n() { entfernte_fehlerpunkte | grep -c . || true; }
  # Diff-Bausteine: Kopf einer Datei plus Hunk.
  kopf() { printf 'diff --git a/%s b/%s\n--- a/%s\n+++ b/%s\n@@ -1,2 +1,2 @@\n' "$1" "$1" "$1" "$1"; }
  t "Fehlerpunkt entfernt"             1 "$( { kopf tools/x.sh; printf -- '-  [ x ] || fail "kaputt"\n'; } | n)"
  t "verschoben in derselben Datei"    0 "$( { kopf tools/x.sh; printf -- '-  [ x ] || fail "k"\n+      [ x ] || fail "k"\n'; } | n)"
  t "in andere Datei verschoben"       1 "$( { kopf tools/x.sh; printf -- '-  fail "k"\n'; kopf tools/y.sh; printf -- '+  fail "k"\n'; } | n)"
  t "zwei weg, eine wieder da"         1 "$( { kopf tools/x.sh; printf -- '-exit 1\n-exit 1\n+exit 1\n'; } | n)"
  t "nacktes exit 1 anderswo neu"      1 "$( { kopf tools/x.sh; printf -- '-exit 1\n'; kopf tools/y.sh; printf -- '+exit 1\n'; } | n)"
  t "umformuliert zaehlt als entfernt" 1 "$( { kopf tools/x.sh; printf -- '-  fail "alt"\n+  fail "neu"\n'; } | n)"
  t "keine Fehlerpunkt-Zeile"          0 "$( { kopf tools/x.sh; printf -- '-  echo hallo\n'; } | n)"
  t "::error:: und exit 2"             2 "$( { kopf tools/x.sh; printf -- '-  echo "::error::x"\n-  exit 2\n'; } | n)"
  t "probe-Zeile"                      1 "$( { kopf tools/x.sh; printf -- '-  probe "Marker" x 0\n'; } | n)"
  t "Inhalt beginnt mit -"             1 "$( { kopf .github/workflows/w.yml; printf -- '-- run: fail "x"\n'; } | n)"
  t "Diff-Koepfe zaehlen nicht"        0 "$( { kopf tools/x.sh; } | n)"
  t "exit 0 und exit 10 zaehlen nicht" 0 "$( { kopf tools/x.sh; printf -- '-  exit 0\n-  exit 10\n'; } | n)"
  t "Kommentar zaehlt nicht"           0 "$( { kopf tools/x.sh; printf -- '-# dann fail "x"\n'; } | n)"
  b() { if printf '%s\n' "$1" | hat_begruendung; then echo ja; else echo nein; fi; }
  t "Trailer mit Grund"                ja   "$(b "$(printf 'x\n\nEntfernte-Pruefungen: ersetzt durch y\n')")"
  t "Trailer leer"                     nein "$(b "$(printf 'x\n\nEntfernte-Pruefungen:   \n')")"
  t "Trailer nur im Fliesstext"        nein "$(b 'siehe Entfernte-Pruefungen: im Text')"
  t "kein Trailer"                     nein "$(b 'feat: x')"
  local anker=1; grep -qF "$MARKER" AGENTS.md && anker=0
  t "Regel-Anker in AGENTS.md"         0 "$anker"

  # Integration gegen ein Wegwerf-Repo: Range-Modus, beide Richtungen. Die
  # git-Umgebung des Aufrufers wird entfernt — im commit-msg-Hook setzt git bei
  # `commit -a` einen ABSOLUTEN GIT_INDEX_FILE, und ein `git add` im Wegwerf-Repo
  # schriebe sonst in den Index des echten Repos (Review slice-222 F-1).
  local res
  res="$(
    unset GIT_INDEX_FILE GIT_DIR GIT_WORK_TREE GIT_PREFIX GIT_OBJECT_DIRECTORY \
          GIT_ALTERNATE_OBJECT_DIRECTORIES GIT_COMMON_DIR
    repo="$(mktemp -d)"; trap 'rm -rf "$repo"' EXIT
    cd "$repo"; git init -q; git config user.email t@t; git config user.name t
    mkdir tools; printf 'Regel %s\n' "$MARKER" >AGENTS.md
    printf '[ x ] || fail "a"\n[ y ] || fail "b"\n' >tools/x.sh
    git add -A; git commit -qm init
    printf '[ x ] || fail "a"\n' >tools/x.sh; git commit -qam 'ohne Grund'
    printf '' >tools/x.sh; git commit -qam "$(printf 'mit Grund\n\nEntfernte-Pruefungen: a entfaellt, weil z\n')"
    r1=0; check_commit HEAD~1 >/dev/null || r1=1
    r2=0; check_commit HEAD >/dev/null || r2=1
    echo "$r1 $r2"
  )"
  t "Repo: entfernt ohne Grund -> rot"  1 "${res%% *}"
  t "Repo: entfernt mit Grund -> gruen" 0 "${res##* }"
  echo "== Fehlschlaege: $fails"
  [ "$fails" -eq 0 ]
}

if [ "${1:-}" = "--selftest" ]; then
  echo "pruefung-entfernt-check: Selbsttest"
  self_test; exit $?
fi

if [ -n "${MSGFILE:-}" ]; then
  rc=0; check_pending "$MSGFILE" || rc=$?
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
