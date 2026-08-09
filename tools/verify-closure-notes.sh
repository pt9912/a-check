#!/usr/bin/env bash
# verify-closure-notes.sh — strukturelle Hälfte der Closure-Pflicht
# (AGENTS.md §5), slice-050 / Fund B-3+B-4 aus slice-048.
#
# Verifikations-Schicht, nicht Gate-Schicht: `make gates` beantwortet
# Code-/Architektur-Fragen, `make verify` DoD-/Closure-Fragen (Regelwerk
# Modul 11). Darum hängt dieses Skript an `verify`, nicht an `gates`.
#
# Geprüft wird ausschliesslich STRUKTUR:
# Die Ueberschrift heisst im Bestand mehrheitlich schlicht "## N. Closure",
# vereinzelt "Closure-Notiz" oder "Verifikation & Closure" — alle drei zaehlen.
# Ausgenommen sind "Closure-Trigger" und "Closure-Kriterien": das sind eigene,
# legitime Abschnitte (slice-001/002/003) und keine Abschluss-Notiz.
#
#   (1) jeder Slice in done/ hat GENAU EINEN Closure-Abschnitt
#       (zwei sind der Platzhalter-neben-Notiz-Fall, real aufgetreten in
#        slice-044 und von diesem Skript gefunden),
#   (2) sein Rumpf ist nicht leer und kein Platzhalter,
#   (3) er trägt mindestens zwei Sätze und keine der bekannten Floskeln.
#
# NICHT geprüft (ehrliche Grenze, AC-QA-02): ob der Inhalt ein echtes
# Lernsignal ist. Das ist semantisch und gehört dem Skill
# .harness/skills/closure-note-reviewer.md — Struktur kann Floskel nur
# listen, nicht erkennen.
#
# EBENFALLS NICHT geprüft: offene DoD-Haken in done/. Eine naive Prüfung
# hätte hier zwei False-Positives: slice-017 führt bewusst einen
# Dauer-Merker als offenen Punkt, slice-039 einen Vorgang in einem fremden
# Repo. Ein belastbarer Check braucht erst einen deklarierten Marker für
# solche Punkte — das ist Form-Arbeit und gehört in Etappe D (Fund B-5),
# nicht in einen Sensor, der sonst sofort Rauschen produziert.
set -euo pipefail
cd "$(dirname "$0")/.."

DONE_DIR="docs/plan/planning/done"
# Nur echte Platzhalter-Wendungen — NICHT jede kursive Klammer: slice-040/041
# tragen substanziellen Text in genau dieser Schreibweise, ein breiter Regex
# haette sie faelschlich als Platzhalter gemeldet (Fehlalarm im ersten Lauf).
PLACEHOLDER='(beim Abschluss|_\(folgt\)_|TODO|TBD|noch offen)'
FLOSKELN='^(fertig|erledigt|wie geplant umgesetzt|war ok|alles gut|läuft)\.?$'

# Rumpf eines Closure-Abschnitts: ab der Überschrift bis zur nächsten
# `## `-Überschrift, ohne die Überschrift selbst.
closure_body() {  # $1 = Datei
  awk '
    /^## .*[Cc]losure/ && !/[Cc]losure-(Trigger|Kriterien)/ { inblock=1; next }
    inblock && /^## / { inblock=0 }
    inblock { print }
  ' "$1"
}

count_closure_headings() {  # $1 = Datei
  grep -E '^## .*[Cc]losure' "$1" | grep -cvE '[Cc]losure-(Trigger|Kriterien)' || true
}

check_file() {  # $1 = Datei; gibt Befunde auf stdout, Rückgabe 1 bei Befund
  local f="$1" n body sentences fail=0
  n="$(count_closure_headings "$f")"
  if [ "$n" -eq 0 ]; then
    echo "$f: kein Closure-Abschnitt (AGENTS.md §5)"
    return 1
  fi
  if [ "$n" -gt 1 ]; then
    echo "$f: $n Closure-Abschnitte — genau einer erwartet; ein zweiter ist typischerweise ein stehengebliebener Platzhalter"
    fail=1
  fi
  body="$(closure_body "$f" | sed '/^[[:space:]]*$/d' | sed '/^---$/d')"
  if [ -z "$body" ]; then
    echo "$f: Closure-Abschnitt ist leer"
    return 1
  fi
  if printf '%s\n' "$body" | grep -qE "$PLACEHOLDER"; then
    echo "$f: Closure-Abschnitt trägt einen Platzhalter statt eines Lerneintrags"
    fail=1
  fi
  if printf '%s\n' "$body" | grep -qiE "$FLOSKELN"; then
    echo "$f: Closure-Abschnitt besteht aus einer Floskel ohne Substanz"
    fail=1
  fi
  # Satzzahl: Punkte/Ausrufe/Fragezeichen ausserhalb von Code-Zeilen.
  sentences="$(printf '%s\n' "$body" | grep -v '^\s*```' | grep -oE '[.!?]' | wc -l)"
  if [ "$sentences" -lt 2 ]; then
    echo "$f: Closure-Abschnitt trägt weniger als zwei Sätze"
    fail=1
  fi
  return "$fail"
}

# Selbsttest: je eine Fixture pro Befundklasse muss feuern, eine gute muss
# schweigen. Ohne ihn wäre ein totes awk-Muster ein False-Green.
self_test() {
  local tmp
  tmp="$(mktemp -d)"
  printf '# s\n\n## 6. Closure-Notiz\n\nDer Test war rot, weil das Muster den Rand nicht traf. Folge-Slice slice-999.\n' > "$tmp/gut.md"
  printf '# s\n\n## 6. Closure-Notiz\n\n_(beim Abschluss.)_\n' > "$tmp/platzhalter.md"
  printf '# s\n\n## 6. Closure-Notiz\n\nFertig.\n' > "$tmp/floskel.md"
  printf '# s\n\n## 5. DoD\n\n- [x] x\n' > "$tmp/ohne.md"
  printf '# s\n\n## 6. Closure-Notiz\n\nEcht, weil belegt. Zweiter Satz.\n\n## 8. Closure-Notiz\n\nZweite. Notiz.\n' > "$tmp/doppelt.md"

  if ! check_file "$tmp/gut.md" >/dev/null; then
    echo "verify-closure-notes: Selbsttest FEHLGESCHLAGEN — gute Notiz faelschlich beanstandet" >&2
    rm -rf "$tmp"; exit 2
  fi
  local bad
  for bad in platzhalter floskel ohne doppelt; do
    if check_file "$tmp/$bad.md" >/dev/null; then
      echo "verify-closure-notes: Selbsttest FEHLGESCHLAGEN — '$bad' nicht erkannt (Pruefung tot)" >&2
      rm -rf "$tmp"; exit 2
    fi
  done
  rm -rf "$tmp"
}

self_test

fail=0
count=0
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
  echo "verify-closure-notes: FAIL — Closure-Pflicht verletzt (AGENTS.md §5)." >&2
  exit 1
fi

# ERWARTETE GRUNDGESAMTHEIT (slice-070, Fund F-12): > 0.
# `done/` ist der Endzustand jedes je abgeschlossenen Slice — dort null Dateien
# zu finden heisst nicht "nichts zu pruefen", sondern "der Pruefbestand ist
# verschwunden". Ohne diese Grenze meldete der Sensor "ok: 0 Slice(s)" mit
# Exit 0 und war damit von einem bestandenen Lauf nicht zu unterscheiden.
if [ "$count" -eq 0 ]; then
  echo "verify-closure-notes: FAIL — keine Slice-Datei in $DONE_DIR gefunden." >&2
  echo "  Erwartet wird ein nicht leerer done/-Bestand; null Dateien sind Bestandsverlust," >&2
  echo "  nicht 'nichts zu pruefen'." >&2
  exit 1
fi
echo "verify-closure-notes ok: $count Slice(s) in done/ mit ausgefuellter Closure-Notiz (Selbsttest gefeuert)."
