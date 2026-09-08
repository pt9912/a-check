#!/usr/bin/env bash
# dcheck-phrase-selftest.sh — Kalibrierungs-Selbsttest fuer phrasen-basierte
# d-check-Modul-Konfigurationen.
#
# Antwort auf BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf (3x: slice-120,
# slice-123, slice-165) -- slice-168 loest auf. Ein Pruefer, dessen
# Trigger-Phrase nichts im Repo trifft, meldet gruen, ohne je etwas geprueft
# zu haben (leere Kandidatenmenge). Fuer a-checks EIGENE Bash-Pruefer loest
# tools/verify-risiko-ausgaenge.sh das seit slice-102 mit einem self_test()
# gegen Gut-/Schlecht-Fixtures. Fuer extern konfigurierte d-check-Module
# (reviews-Trigger-Phrase, structure tasks-ignore-pattern) gibt es dafuer
# kein Aequivalent -- dieses Skript ist es: es faehrt den gepinnten
# d-check-Digest gegen eigene Fixtures und prueft, ob die HEUTE in AGENTS.md
# empfohlene Formulierung ("unabhängiger Review") tatsaechlich noch feuert.
#
# ZWEI KONTROLLEN JE MUSTER, wie bei verify-risiko-ausgaenge.sh: POSITIV (die
# empfohlene Phrase loest aus) und NEGATIV (eine Phrase, die es nicht sollte,
# loest NICHT aus) -- ohne die zweite waere ein Muster, das alles durchlaesst,
# von einem korrekten nicht zu unterscheiden.
#
# ZWEI HAELFTEN, seit slice-169. Die WERKZEUG-Seite oben fragt: reagiert
# d-check noch auf die Phrase? Die KORPUS-Seite unten fragt das Gegenstueck:
# traegt a-checks eigener Bestand sie noch? Genau die ist zweimal ausgefallen
# (slice-120, slice-165) -- ein Muster kann tadellos funktionieren und trotzdem
# nichts pruefen, weil die Kandidatenmenge leer geworden ist.
#
# Die Korpus-Seite liest das Muster AUS .d-check.yml, nicht aus einer Kopie
# hier (Review slice-168, F-2: eine Kopie neben dem Original bleibt gruen,
# nachdem das Original gebrochen wurde). Der Auszug ist fail-closed: findet er
# das Feld nicht, bricht der Lauf ab, statt mit leerem Muster durchzulaufen.
#
# NICHT GEPRUEFT (ehrliche Grenzen):
#  - ob HERUNTERGENOMMENE Formulierungen (die der Kurs kuenftig einfuehrt)
#    ebenfalls greifen wuerden -- nur die AKTUELL empfohlene. Ein neuer Nachzug
#    bei jeder Formulierungs-Aenderung bleibt Handarbeit.
#  - die uebrigen ZWOELF phrasen-basierten Felder in .d-check.yml. Geprueft
#    sind die zwei, deren Ausfall BELEGT ist; fuer die anderen gibt es keinen
#    Vorfall, und ein Sensor ohne Anlass ist selbst eine Behauptung.
set -euo pipefail

DCHECK_REF="${DCHECK_REF:?DCHECK_REF muss gesetzt sein (Makefile uebergibt es)}"
DOCKER="${DOCKER:-docker}"

TMP="$(mktemp -d)"
trap 'rm -rf "$TMP"' EXIT

run_dcheck() {  # $1 = Repo-Wurzel, Rest = zusaetzliche d-check-Flags
  local dir="$1"; shift
  "$DOCKER" run --rm --network none -v "$dir:/repo:ro" "$DCHECK_REF" "$@" 2>&1 || true
}

git_init_fixture() {  # $1 = Verzeichnis
  chmod -R a+rX "$1"  # mktemp -d liefert 0700; Docker (Root im Container) braucht Lesezugriff
  git -C "$1" init -q
  git -C "$1" -c user.email=fixture@local -c user.name=fixture add -A
  git -C "$1" -c user.email=fixture@local -c user.name=fixture commit -qm fixture >/dev/null
}

fail=0

# --- Muster 1: `reviews`-Modul, Trigger-Phrase "unabhängiger Review" -----

setup_reviews_fixture() {  # $1 = Zielverzeichnis, $2 = DoD-Zeilentext
  local dir="$1" line="$2"
  rm -rf "$dir"
  mkdir -p "$dir/docs/plan/planning/done" "$dir/docs/reviews"
  cat > "$dir/.d-check.yml" <<'YAML'
modules: [reviews]
reviews:
  done-dir: docs/plan/planning/done
  reviews-dir: docs/reviews
YAML
  cat > "$dir/docs/plan/planning/done/slice-900-fixture.md" <<EOF
# slice-900 fixture

## DoD

- [x] ${line}
EOF
  git_init_fixture "$dir"
}

setup_reviews_fixture "$TMP/reviews-pos" \
  "Unabhängiger Review durchgeführt, Report unter \`docs/reviews/\` liegt vor."
out="$(run_dcheck "$TMP/reviews-pos")"
if ! printf '%s' "$out" | grep -q "review-missing"; then
  echo "dcheck-phrase-selftest: FAIL — Trigger-Phrase 'unabhängiger Review' loest das reviews-Modul nicht mehr aus (Positiv-Kontrolle)." >&2
  printf '%s\n' "$out" >&2
  fail=1
fi

setup_reviews_fixture "$TMP/reviews-neg" \
  "Unabhängiges Plan-Review über getrennten Kontext durchgeführt."
out="$(run_dcheck "$TMP/reviews-neg")"
if printf '%s' "$out" | grep -q "review-missing"; then
  echo "dcheck-phrase-selftest: FAIL — reviews-Modul loest auf eine Formulierung aus, die es laut AGENTS.md nicht sollte (Negativ-Kontrolle)." >&2
  fail=1
fi

# --- Muster 2: `structure`-Modul, tasks-ignore-pattern "Unabhängiger Review" ---

# Nur der eine relevante Modul-Block, nicht a-checks volle .d-check.yml —
# eine vollstaendige Kopie zieht Config fuer andere Module mit (z. B. `ids`),
# die auf a-check-eigene Pfade verweist und in einem Fixture-Repo ohne diese
# Pfade mit einem Konfigurationsfehler abbricht, noch bevor `structure` laeuft.
TASKS_IGNORE_PATTERN='^( *grün|Unabhängiger Review|Closure-Notiz|Beobachtungs-Register|Jedes Risiko|Reconciliation)'

setup_structure_fixture() {  # $1 = Zielverzeichnis, $2 = vierter DoD-Punkt
  local dir="$1" fourth="$2"
  rm -rf "$dir"
  mkdir -p "$dir/docs/plan/planning"
  cat > "$dir/.d-check.yml" <<YAML
modules: [structure]
structure:
  - files: "docs/plan/planning/**/slice-*.md"
    section-pattern: '^#+ .*(DoD|Definition of Done)'
    max-tasks: 3
    tasks-ignore-pattern: '${TASKS_IGNORE_PATTERN}'
YAML
  cat > "$dir/docs/plan/planning/slice-901-fixture.md" <<EOF
# slice-901 fixture

## DoD

- [x] Erster echter Liefer-Punkt.
- [x] Zweiter echter Liefer-Punkt.
- [x] Dritter echter Liefer-Punkt.
- [x] ${fourth}
EOF
  git_init_fixture "$dir"
}

setup_structure_fixture "$TMP/structure-pos" \
  "Unabhängiger Review durchgeführt, Report unter \`docs/reviews/\` liegt vor."
out="$(run_dcheck "$TMP/structure-pos")"
if printf '%s' "$out" | grep -q "section-oversized"; then
  echo "dcheck-phrase-selftest: FAIL — 'Unabhängiger Review' wird von tasks-ignore-pattern nicht mehr als konstant erkannt (Positiv-Kontrolle, vier Punkte muessten auf drei bereinigt werden)." >&2
  printf '%s\n' "$out" >&2
  fail=1
fi

setup_structure_fixture "$TMP/structure-neg" \
  "Vierter echter Liefer-Punkt, der mitzaehlen muss."
out="$(run_dcheck "$TMP/structure-neg")"
if ! printf '%s' "$out" | grep -q "section-oversized"; then
  echo "dcheck-phrase-selftest: FAIL — tasks-ignore-pattern ignoriert einen Punkt, der eigentlich mitzaehlen sollte (Negativ-Kontrolle, vier echte Punkte muessten rot melden)." >&2
  fail=1
fi

# --- KORPUS-SEITE: traegt der echte Bestand die Trigger-Phrase noch? -------
#
# Gefragt ist NICHTLEERHEIT, keine Erwartungszahl. Die belegte Ausfallart ist
# "die Kandidatenmenge wird leer"; eine feste Zahl braeche zusaetzlich bei
# jedem neuen Slice.
#
# GEZAEHLT WIRD DIE MENGE DES MODULS, nicht eine Obermenge davon. Das Modul
# reviews sieht einen DoD-Haken in einem FLACHEN done/-Slice; eine Nennung in
# Prosa, in einer Tabelle oder in einer Wellen-Ergebnisnotiz sieht es nicht.
# Wer weiter zaehlt, meldet gruen, waehrend die echte Menge leer ist -- genau
# die Ausfallart, gegen die diese Kontrolle steht (Review slice-169, F-1:
# gemessen auf einer Kopie meldete die vorige Fassung Exit 0, nachdem beide
# echten Kandidaten entwertet waren).
#
# done-dir kommt aus .d-check.yml, nicht aus einer Kopie hier: eine Kopie
# neben dem Original bleibt gruen, nachdem das Original umgezogen ist
# (Review slice-168, F-2).

REVIEW_PHRASE="unabhängiger Review"

done_dir="$(sed -n "s/^ *done-dir: *\(.*[^ ]\) *$/\1/p" .d-check.yml | head -1)"
if [ -z "$done_dir" ] || [ ! -d "$done_dir" ]; then
  echo "dcheck-phrase-selftest: FAIL — done-dir aus .d-check.yml nicht lesbar oder kein Verzeichnis." >&2
  echo "  Der Auszug ist fail-closed: ohne Kandidatenverzeichnis wird nicht geprueft, sondern abgebrochen." >&2
  exit 1
fi

# Flach, nicht rekursiv -- archivierte Stubs unter done/<welle>/ und
# done/wellenlos/ tragen keine DoD mehr und sind fuer das Modul keine
# Kandidaten.
korpus_reviews=0
for f in "$done_dir"/slice-*.md; do
  [ -f "$f" ] || continue
  if grep -qE '^- \[[ x]\] .*[Uu]nabhängiger Review' "$f"; then
    korpus_reviews=$((korpus_reviews + 1))
  fi
done

if [ "$korpus_reviews" -eq 0 ]; then
  echo "dcheck-phrase-selftest: FAIL — Kandidatenmenge des reviews-Moduls ist LEER." >&2
  echo "  Kein flacher Slice in $done_dir traegt \"$REVIEW_PHRASE\" auf einem DoD-Haken;" >&2
  echo "  make doc-reviews meldet damit gruen, ohne etwas zu pruefen. Die Phrase gehoert in" >&2
  echo "  die DoD neuer Slices (AGENTS.md §5, Kopieranleitung Punkt 7)." >&2
  fail=1
fi

if [ "$fail" -ne 0 ]; then
  echo "dcheck-phrase-selftest: FAIL — mindestens eine Kalibrierungs-Kontrolle ist rot." >&2
  exit 1
fi

echo "dcheck-phrase-selftest ok: 4 Werkzeug-Kontrollen (2 Muster × 2 Richtungen) und 1 Korpus-Kontrolle"
echo "  (reviews-Phrase auf einem DoD-Haken: ${korpus_reviews} flache(r) Slice(s) in ${done_dir})."
echo "  NICHT geprueft: die uebrigen phrasen-basierten Felder in .d-check.yml — fuer sie gibt es"
echo "  keinen belegten Ausfall, und ein Sensor ohne Anlass ist selbst eine Behauptung."
