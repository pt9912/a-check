#!/usr/bin/env bash
# slice-mv.sh — Lifecycle-Wechsel eines Slice samt der Verweise AUF ihn
# (AGENTS.md §5), Antwort auf BEO-008 bei 3x (slice-118).
#
# PROBLEM. `git mv` bewegt die Datei; die Verweise auf sie bleiben stehen und
# zeigen ins Leere. verify-slice-links prueft die Gegenrichtung — Verweise IN
# wandernden Dateien — und nimmt done/ als Quelle ausdruecklich aus; gefangen
# hat es darum jedes Mal doc-check, NACH dem Wechsel. Dreimal so passiert
# (slice-093, slice-114, slice-080).
#
# WARUM WERKZEUG UND NICHT GUIDE. Die Schwester-Klasse SL-002 zaehlt neun
# Vorfaelle, ZWEI davon nach dem Guide in Schritt 9 des Workflow-Skeletts. Ein
# Guide, der an derselben Familie schon einmal versagt hat, ist nicht die
# staerkere Antwort (Baseline: ab dem dritten Vorfall Guide ODER Sensor).
#
# WARUM EIN SKRIPT, WO slice-080 GERADE 509 ZEILEN ABBAUTE. Die abgeloesten
# waren PRUEFUNGEN — die kann ein Modul uebernehmen. Dies ist eine BEWEGUNG;
# d-check ist read-only (DC-QA-03) und wird nie eine Datei verschieben.
#
# NICHT BEHANDELT (ehrliche Grenzen, AC-QA-02) — zwei Stueck.
#
# (1) SEMANTIK. Das Werkzeug zieht PFADE nach, keine Aussagen. Die Roadmap-Zeile
# "In Arbeit: <slice>" bleibt nach dem Wechsel nach done/ stehen und muss von Hand
# fallen; ihr Verweis ist dann korrekt und ihre Aussage falsch. Beim ersten Einsatz
# (slice-118 selbst) sofort aufgetreten. Ein Werkzeug, das Zustandssaetze umschreibt,
# muesste raten, welcher Satz einen Zustand behauptet — das ist Urteil, kein Match.
#
# (2) WELLE-PLAN-DATEIEN. Sie wechseln
# beim Closure-mv die Verzeichnis-TIEFE (flach -> done/), und ein Pfad aus
# Tiefe n braucht aus n+1 ein zusaetzliches "../" — eine andere Ersetzung als
# der Verzeichnis-Tausch auf gleicher Ebene (slice-089).
set -euo pipefail
cd "$(dirname "$0")/.."

PLANNING="docs/plan/planning"
LIFECYCLE="open next in-progress done"

usage() {
  cat >&2 <<'USAGE'
Aufruf: make slice-mv SLICE=<slice-NNN[-kurztitel[.md]]> TO=<open|next|in-progress|done>

  Bewegt den Slice per `git mv` und zieht die Verweise AUF ihn repo-weit nach.
  Beide im Bestand vorkommenden Formen werden getroffen:
      ../<verzeichnis>/<datei>            und
      docs/plan/planning/<verzeichnis>/<datei>
USAGE
}

# Die Ersetzung als eigene Funktion — der Selbsttest muss sie pruefen koennen,
# ohne ein git-Repo zu bewegen.
rewrite_file() {  # $1=zieldatei  $2=basename  $3=von  $4=nach
  sed -i \
    -e "s|\.\./$3/$2|../$4/$2|g" \
    -e "s|$PLANNING/$3/$2|$PLANNING/$4/$2|g" \
    "$1"
}

self_test() {
  local tmp f
  tmp="$(mktemp -d)"
  f="$tmp/probe.md"
  {
    echo '[a](../open/slice-999-x.md)'                    # trifft
    echo 'siehe `docs/plan/planning/open/slice-999-x.md`' # trifft, auch zitiert
    echo '[b](../open/slice-998-y.md)'                    # ANDERE Datei
    echo '[c](../done/slice-999-x.md)'                    # ANDERES Verzeichnis
    echo 'open/slice-999-x.md'                            # keine der zwei Formen
  } > "$f"
  rewrite_file "$f" "slice-999-x.md" "open" "done"
  # Beide Zielformen sind umgeschrieben ...
  if grep -q '\.\./open/slice-999-x\.md' "$f" \
     || grep -q "$PLANNING/open/slice-999-x\.md" "$f"; then
    echo "slice-mv: Selbsttest FEHLGESCHLAGEN — Zielform nicht ersetzt" >&2
    rm -rf "$tmp"; exit 2
  fi
  if ! grep -q '\.\./done/slice-999-x\.md' "$f" \
     || ! grep -q "$PLANNING/done/slice-999-x\.md" "$f"; then
    echo "slice-mv: Selbsttest FEHLGESCHLAGEN — Ersetzung nicht angekommen" >&2
    rm -rf "$tmp"; exit 2
  fi
  # ... und NUR sie. Ohne diese drei Gegenproben waere eine Ersetzung, die
  # alles umschreibt, von einer korrekten nicht zu unterscheiden.
  if ! grep -q '\.\./open/slice-998-y\.md' "$f"; then
    echo "slice-mv: Selbsttest FEHLGESCHLAGEN — fremde Datei mitgeaendert" >&2
    rm -rf "$tmp"; exit 2
  fi
  if [ "$(grep -c '\.\./done/slice-999-x\.md' "$f")" -ne 2 ]; then
    echo "slice-mv: Selbsttest FEHLGESCHLAGEN — bereits richtiger Verweis verdoppelt/verloren" >&2
    rm -rf "$tmp"; exit 2
  fi
  if ! grep -qx 'open/slice-999-x.md' "$f"; then
    echo "slice-mv: Selbsttest FEHLGESCHLAGEN — praefixlose Nicht-Form mitgeaendert" >&2
    rm -rf "$tmp"; exit 2
  fi
  rm -rf "$tmp"
}

self_test

SLICE="${1:-}"
TO="${2:-}"
[ -n "$SLICE" ] && [ -n "$TO" ] || { usage; exit 2; }

case " $LIFECYCLE " in
  *" $TO "*) ;;
  *) echo "slice-mv: '$TO' ist kein Lifecycle-Verzeichnis ($LIFECYCLE)" >&2; exit 2 ;;
esac

# Quelle finden: Praefix oder voller Dateiname, in genau EINEM Verzeichnis.
found=""
for d in $LIFECYCLE; do
  for f in "$PLANNING/$d/"${SLICE%.md}*.md; do
    [ -e "$f" ] || continue
    if [ -n "$found" ]; then
      echo "slice-mv: '$SLICE' ist mehrdeutig — $found und $f" >&2
      exit 2
    fi
    found="$f"
  done
done
[ -n "$found" ] || { echo "slice-mv: kein Slice '$SLICE' unter $PLANNING/" >&2; exit 2; }

base="$(basename "$found")"
from="$(basename "$(dirname "$found")")"
if [ "$from" = "$TO" ]; then
  echo "slice-mv: '$base' liegt bereits in $TO/" >&2
  exit 2
fi

git mv "$found" "$PLANNING/$TO/$base"

# Verweise repo-weit nachziehen — NICHT nur unter docs/plan/planning/: gemessen
# tragen auch docs/reviews/ solche Verweise. Die vendored Baseline bleibt
# aussen vor (MR-006: unveraenderter Fremdtext).
count=0
while IFS= read -r f; do
  [ -n "$f" ] || continue
  rewrite_file "$f" "$base" "$from" "$TO"
  count=$((count + 1))
done <<< "$(grep -rl -e "\.\./$from/$base" -e "$PLANNING/$from/$base" \
             --include='*.md' . 2>/dev/null | grep -v '^\./\.harness/baseline/' || true)"

echo "slice-mv ok: $base  $from/ -> $TO/, $count Datei(en) mit Verweisen nachgezogen (Selbsttest gefeuert)."
echo "  Beide Aenderungen gehoeren in EINEN Commit: der Rename bleibt bei 100 %, die Verweise"
echo "  liegen in ANDEREN Dateien (AGENTS.md §3.3)."
