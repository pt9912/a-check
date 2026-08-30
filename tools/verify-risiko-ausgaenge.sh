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
# Stichtag fuer die Risiko-Ausgaenge (slice-102, Baseline v5.12.0 modul-05
# §Offene Risiken werden bei Closure aufgeloest). Bootstrap-aware wie bei
# verify-slice-form: aeltere Notizen entstanden, BEVOR die geschlossene
# Dreier-Menge im Repo galt — slice-092 fuehrt etwa "Maintainer-Entscheidung"
# als Ausgang, damals korrekt, heute ausserhalb der Menge. Rueckwirkend
# umzuschreiben waere Geschichts-Politur.
RISK_FROM=102
# Die drei Ausgaenge als FORM, nicht als Inhalt: eingetreten -> Carveout oder
# Folge-Slice mit ID · entfallen -> gestrichen mit Begruendung · weiter offen ->
# Beobachtungs-Register. Ob der eingetragene Ausgang TRAEGT, bleibt Urteil
# (modul-06: "Mensch urteilt, Maschine prueft Deckung").
AUSGAENGE='(Carveout|Folge-Slice|gestrichen mit Begründung|Beobachtungs-Register|BEO-[0-9]{3})'
# Nur echte Platzhalter-Wendungen — NICHT jede kursive Klammer: slice-040/041
# tragen substanziellen Text in genau dieser Schreibweise, ein breiter Regex
# haette sie faelschlich als Platzhalter gemeldet (Fehlalarm im ersten Lauf).
PLACEHOLDER='(beim Abschluss|_\(folgt\)_|TODO|TBD|noch offen)'
FLOSKELN='^(fertig|erledigt|wie geplant umgesetzt|war ok|alles gut|läuft)\.?$'

# Rumpf eines Closure-Abschnitts: ab der Überschrift bis zur nächsten
# `## `-Überschrift, ohne die Überschrift selbst.
slice_num() {  # $1 = Pfad -> Nummer ohne fuehrende Nullen, leer wenn keine
  basename "$1" | sed -nE 's/^slice-0*([0-9]+)-.*/\1/p'
}

# Risiko-Block als EIGENE Sektion (Ziel-Form v5.12.0: "## 6. Risiken und
# offene Punkte"). Die alte a-check-Form fuehrt ihn stattdessen als Block IN
# der Closure-Notiz; beide werden gelesen, damit der Gliederungs-Wechsel
# (slice-107) den Check nicht still abschaltet — genau die Klasse
# "halluziniertes Gate", gegen die modul-13 steht.
risiko_sektion() {  # $1 = Datei
  awk '
    /^## .*[Rr]isiken/ { inblock=1; next }
    inblock && /^## / { inblock=0 }
    inblock { print }
  ' "$1"
}

# Aufzaehlungspunkte zusammenkleben: ein Ausgang steht oft erst in der zweiten
# Zeile. Fortsetzungszeilen NORMALISIERT ankleben (Einrueckung raus), sonst
# zerreisst ein Umbruch die Wendung — real aufgetreten an slice-102.
#
# GETRENNT von der Extraktion (slice-107): woher die Zeilen kommen — Block IN
# der Closure (alte Form) oder eigene Sektion (Ziel-Form v5.12.0) — entscheiden
# die Funktionen darueber. Diese hier sieht nur noch Zeilen.
punkte_kleben() {  # Zeilen auf stdin
  awk '
    {
      line = $0
      sub(/^[[:space:]]+/, "", line)
      if ($0 ~ /^- /) { if (buf != "") print buf; buf = line }
      else if (buf != "") { buf = buf " " line }
    }
    END { if (buf != "") print buf }
  '
}

# Risiko-Block INNERHALB der Closure-Notiz (a-checks Form bis slice-106).
risiko_aus_closure() {  # Closure-Rumpf auf stdin
  awk '
    /^\*\*Offene Risiken/ { inblock=1; next }
    inblock && /^\*\*/ && !/^\*\*Offene Risiken/ { inblock=0 }
    inblock { print }
  '
}

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
  local f="$1" n body prosa sentences fail=0
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
  # ZITAT-VORFILTER (slice-100, SL-004). Text, der UEBER ein Muster spricht, ist
  # nicht das Muster. Der Sensor hat zweimal an einer Notiz gefeuert, die die
  # ausloesende Wendung in Backticks zitierte. Denselben Vorfilter benutzen die
  # Satzzaehlung unten (seit slice-075) und verify-slice-links — die Placeholder-
  # und Floskel-Pruefung sahen ihn als einzige nicht.
  #
  # KEINE Lockerung: die Platzhalter-Liste bleibt unveraendert. Sie enthaelt eine
  # Wendung, die in einem Risiko-Block auch UNZITIERT legitim vorkommt; sie zu
  # streichen waere eine Schwellen-Senkung und braucht nach AGENTS.md §3.6 eine
  # ADR. Diese Frage steht im Beobachtungs-Register, nicht in diesem Skript.
  prosa="$(printf '%s\n' "$body" \
    | sed -e '/^[[:space:]]*```/,/^[[:space:]]*```/d' -e 's/`[^`]*`//g')"
  if printf '%s\n' "$prosa" | grep -qE "$PLACEHOLDER"; then
    echo "$f: Closure-Abschnitt trägt einen Platzhalter statt eines Lerneintrags"
    fail=1
  fi
  if printf '%s\n' "$prosa" | grep -qiE "$FLOSKELN"; then
    echo "$f: Closure-Abschnitt besteht aus einer Floskel ohne Substanz"
    fail=1
  fi
  # Risiko-Ausgaenge (slice-102). Geprueft wird INNERHALB eines vorhandenen
  # Blocks, nicht seine Existenz: modul-05 bindet die Pflicht an *notierte*
  # Risiken. Wer den Block weglaesst, wird nicht erwischt — benannte Grenze,
  # dieselbe Klasse wie "hoechstens zwei Schichten" in verify-slice-form.
  local num punkt
  num="$(slice_num "$f")"
  if [ -n "$num" ] && [ "$num" -ge "$RISK_FROM" ]; then
    while IFS= read -r punkt; do
      [ -n "$punkt" ] || continue
      if ! printf '%s' "$punkt" | grep -qE "Ausgang:.*$AUSGAENGE"; then
        echo "$f: Risiko ohne Ausgang aus der geschlossenen Dreier-Menge: ${punkt:0:70}…"
        fail=1
      fi
    done <<< "$( { printf '%s\n' "$body" | risiko_aus_closure; risiko_sektion "$f"; } | punkte_kleben )"
  fi
  # Satzzahl (slice-075, Fund F-13): gezaehlt werden Satzzeichen, die tatsaechlich
  # ein Satzende markieren — also von Whitespace oder Zeilenende gefolgt. Die
  # alte Zaehlung nahm JEDES `.`/`!`/`?`; die einzeilige Notiz
  # "Geprueft via foo.go." ergab dadurch zwei Saetze statt einem und bestand die
  # Mindestzahl.
  # Code-Bloecke fallen jetzt VOLLSTAENDIG weg statt nur ihrer Fence-Zeilen: das
  # alte `grep -v '^\s*```'` liess den Inhalt stehen, obwohl der Kommentar
  # "ausserhalb von Code-Zeilen" behauptete. Inline-Code ebenso — beides derselbe
  # Vorfilter wie in verify-slice-links.
  sentences="$(printf '%s\n' "$prosa" | grep -oE '[.!?]([[:space:]]|$)' | wc -l)"
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
  # RISIKO-AUSGAENGE, drei Richtungen (slice-102). Ohne die dritte Fixture
  # waere ein Muster, das alles durchlaesst, von einem korrekten nicht zu
  # unterscheiden — und ohne die zweite eines, das nur "Ausgang:" sucht.
  { echo '## 6. Closure-Notiz'; echo 'Text hier. Zweiter Satz.'; echo '';
    echo '**Offene Risiken und ihr Ausgang:**'; echo '';
    echo '- *A* — Ausgang: **Folge-Slice**, slice-900.';
    echo '- *B* — Ausgang: **gestrichen mit Begründung**, weil weg.';
    echo '- *C* — Ausgang: **weiter offen**, fürs Beobachtungs-Register.'; } > "$tmp/slice-102-gut.md"
  if ! check_file "$tmp/slice-102-gut.md" >/dev/null; then
    echo "verify-closure-notes: Selbsttest FEHLGESCHLAGEN — gueltige Risiko-Ausgaenge beanstandet" >&2
    rm -rf "$tmp"; exit 2
  fi
  # Ausgang ueber einen Zeilenumbruch hinweg — die Probe fuer das Ankleben.
  { echo '## 6. Closure-Notiz'; echo 'Text hier. Zweiter Satz.'; echo '';
    echo '**Offene Risiken und ihr Ausgang:**'; echo '';
    echo '- *A* — Ausgang: **gestrichen mit'; echo '  Begründung**, weil weg.'; } > "$tmp/slice-102-umbruch.md"
  if ! check_file "$tmp/slice-102-umbruch.md" >/dev/null; then
    echo "verify-closure-notes: Selbsttest FEHLGESCHLAGEN — Ausgang ueber Zeilenumbruch nicht erkannt" >&2
    rm -rf "$tmp"; exit 2
  fi
  # NEUE GLIEDERUNG (slice-107): Risiken als eigene Sektion, nicht in der Closure.
  { echo '## 6. Risiken und offene Punkte'; echo '';
    echo '- *A* — **Ausgang:** weiter offen, `BEO-001` im Register.'; echo '';
    echo '## 7. Closure-Notiz'; echo 'Text hier. Zweiter Satz.'; } > "$tmp/slice-102-sektion-gut.md"
  if ! check_file "$tmp/slice-102-sektion-gut.md" >/dev/null; then
    echo "verify-closure-notes: Selbsttest FEHLGESCHLAGEN — gueltiger Ausgang in eigener Sektion beanstandet" >&2
    rm -rf "$tmp"; exit 2
  fi
  { echo '## 6. Risiken und offene Punkte'; echo '';
    echo '- *A* — bleibt halt so.'; echo '';
    echo '## 7. Closure-Notiz'; echo 'Text hier. Zweiter Satz.'; } > "$tmp/slice-102-sektion-ohne.md"
  if check_file "$tmp/slice-102-sektion-ohne.md" >/dev/null; then
    echo "verify-closure-notes: Selbsttest FEHLGESCHLAGEN — Risiko ohne Ausgang in eigener Sektion nicht erkannt" >&2
    rm -rf "$tmp"; exit 2
  fi
  { echo '## 6. Closure-Notiz'; echo 'Text hier. Zweiter Satz.'; echo '';
    echo '**Offene Risiken und ihr Ausgang:**'; echo '';
    echo '- *A* — bleibt halt so.'; } > "$tmp/slice-102-ohne.md"
  if check_file "$tmp/slice-102-ohne.md" >/dev/null; then
    echo "verify-closure-notes: Selbsttest FEHLGESCHLAGEN — Risiko ohne Ausgang nicht erkannt" >&2
    rm -rf "$tmp"; exit 2
  fi
  { echo '## 6. Closure-Notiz'; echo 'Text hier. Zweiter Satz.'; echo '';
    echo '**Offene Risiken und ihr Ausgang:**'; echo '';
    echo '- *A* — Ausgang: **Maintainer-Entscheidung**.'; } > "$tmp/slice-102-fremd.md"
  if check_file "$tmp/slice-102-fremd.md" >/dev/null; then
    echo "verify-closure-notes: Selbsttest FEHLGESCHLAGEN — Ausgang ausserhalb der Dreier-Menge akzeptiert" >&2
    rm -rf "$tmp"; exit 2
  fi
  # Grandfathering: derselbe Block unter dem Stichtag muss schweigen.
  cp "$tmp/slice-102-ohne.md" "$tmp/slice-091-alt.md"
  if ! check_file "$tmp/slice-091-alt.md" >/dev/null; then
    echo "verify-closure-notes: Selbsttest FEHLGESCHLAGEN — Risiko-Regel greift trotz Grandfathering" >&2
    rm -rf "$tmp"; exit 2
  fi
  # ZITAT-VORFILTER, beide Richtungen (slice-100, SL-004). Die erste Fixture ist
  # die eigentliche Probe: dieselbe Wendung, einmal zitiert und einmal nicht.
  # Ohne die zweite waere ein Vorfilter, der alles verschluckt, von einem
  # korrekten nicht zu unterscheiden.
  { echo '## 6. Closure-Notiz';
    echo 'Der Sensor kennt `noch offen` als Platzhalter-Wendung. Das ist ein Zitat.'; } > "$tmp/zitiert.md"
  if ! check_file "$tmp/zitiert.md" >/dev/null; then
    echo "verify-closure-notes: Selbsttest FEHLGESCHLAGEN — zitiertes Muster als Platzhalter gemeldet (SL-004)" >&2
    rm -rf "$tmp"; exit 2
  fi
  { echo '## 6. Closure-Notiz';
    echo 'Der Rest ist noch offen. Zweiter Satz hier.'; } > "$tmp/unzitiert.md"
  if check_file "$tmp/unzitiert.md" >/dev/null; then
    echo "verify-closure-notes: Selbsttest FEHLGESCHLAGEN — unzitierter Platzhalter nicht erkannt (Vorfilter zu breit)" >&2
    rm -rf "$tmp"; exit 2
  fi
  # Satzzaehlung, beide Richtungen (slice-075, Fund F-13). Die erste Fixture ist
  # die eigentliche Probe: ein Punkt im Dateinamen darf kein Satzende sein.
  { echo '## 6. Closure-Notiz'; echo 'Geprueft via foo.go.'; } > "$tmp/ein-satz.md"
  if check_file "$tmp/ein-satz.md" >/dev/null; then
    echo "verify-closure-notes: Selbsttest FEHLGESCHLAGEN — Punkt im Dateinamen als Satzende gezaehlt (F-13)" >&2
    rm -rf "$tmp"; exit 2
  fi
  # Gegenprobe: zwei echte Saetze, einer davon MIT Dateiname — sonst waere eine
  # Zaehlung, die nur nie zaehlt, von einer korrekten nicht zu unterscheiden.
  { echo '## 6. Closure-Notiz'; echo 'Erster Satz via foo.go. Zweiter Satz hier.'; } > "$tmp/zwei-saetze.md"
  if ! check_file "$tmp/zwei-saetze.md" >/dev/null; then
    echo "verify-closure-notes: Selbsttest FEHLGESCHLAGEN — zwei echte Saetze faelschlich beanstandet" >&2
    rm -rf "$tmp"; exit 2
  fi
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
