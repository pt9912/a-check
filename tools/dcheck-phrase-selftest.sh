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
# MEHRERE KONTROLLEN JE MUSTER (Positiv/Negativ, bei Muster 3 zusaetzlich
# Scoping und Zitat), wie bei verify-risiko-ausgaenge.sh: POSITIV (die
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
#  - die uebrigen phrasen-basierten Felder in .d-check.yml. Geprueft sind die
#    drei, deren Ausfall BELEGT ist oder deren Anlass im Slice-202-Review lag;
#    fuer die anderen gibt es keinen Vorfall, und ein Sensor ohne Anlass ist
#    selbst eine Behauptung.
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

# Die Fixture traegt nur den structure-Block, nicht a-checks volle
# .d-check.yml. Grenze der Fixture: Config anderer Module (z. B. `ids`) zeigt
# auf a-check-eigene Pfade, die es hier nicht gibt — d-check bricht dann mit
# einem Konfigurationsfehler ab, bevor `structure` laeuft.
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

# --- Muster 3: `structure`-Modul, forbid-pattern "DoD-Haekchen in open/" ----

# Slice-202: ein [x] in einem open/-Slice attestiert einen Vorgang, der noch
# aussteht (BEO-GATE/attestierung-vor-dem-vorgang, 3x: slice-169, slice-197,
# slice-200). VIER Kontrollen: POSITIV (ein [x] in open/ feuert) · NEGATIV
# (ein [ ] in open/ feuert nicht) · SCOPING (ein [x] in in-progress/ feuert
# nicht — dort sind Haekchen der Normalfall) · ZITAT (Code-Block- und
# Inline-Code-Nennungen des Musters bleiben gruen, SL-004). GRENZE: next/ ist
# in .d-check.yml nicht gedeckt (leere Kandidatenmenge laeuft fail-closed
# rot, gemessen slice-202) — die Fixture spiegelt genau diesen Zustand.

# Das Muster wird aus .d-check.yml gelesen, nicht kopiert (Review slice-202,
# F-7: eine Kopie neben dem Original bleibt gruen, nachdem das Original
# gebrochen wurde). Der Auszug ist fail-closed.
HAEKCHEN_PATTERN="$(grep -A6 'planning/open/\*\*/slice-' .d-check.yml | grep -m1 'forbid-pattern' | sed -e "s/.*forbid-pattern: '//" -e "s/'.*//")"
if [ -z "$HAEKCHEN_PATTERN" ]; then
  echo "dcheck-phrase-selftest: FAIL — forbid-pattern der open/-Bedingung nicht in .d-check.yml gefunden (fail-closed)." >&2
  exit 2
fi

setup_haekchen_fixture() {  # $1 = Zielverz., $2 = Lebenslage (open|next|in-progress), $3 = DoD-Body-Datei
  local dir="$1" stage="$2" body="$3"
  rm -rf "$dir"
  mkdir -p "$dir/docs/plan/planning/$stage"
  cat > "$dir/.d-check.yml" <<YAML
modules: [structure]
structure:
  - files: "docs/plan/planning/open/**/slice-*.md"
    section-pattern: '^#+ .*(DoD|Definition of Done)'
    forbid-pattern: '${HAEKCHEN_PATTERN}'
YAML
  { echo "# slice-902 fixture"; echo; echo "## DoD"; echo; cat "$body"; } > "$dir/docs/plan/planning/$stage/slice-902-fixture.md"
  git_init_fixture "$dir"
}

printf '%s\n' "- [x] Attestierter Vorgang, der noch aussteht." > "$TMP/body-pos.txt"
setup_haekchen_fixture "$TMP/haekchen-pos" "open" "$TMP/body-pos.txt"
out="$(run_dcheck "$TMP/haekchen-pos")"
if ! printf '%s' "$out" | grep -q "section-forbidden"; then
  echo "dcheck-phrase-selftest: FAIL — ein [x] in open/ feuert das forbid-pattern nicht mehr (Positiv-Kontrolle, attestierung-vor-dem-vorgang 3x)." >&2
  printf '%s\n' "$out" >&2
  fail=1
fi

printf '%s\n' "- [ ] Echte offene Aufgabe." > "$TMP/body-neg.txt"
setup_haekchen_fixture "$TMP/haekchen-neg" "open" "$TMP/body-neg.txt"
out="$(run_dcheck "$TMP/haekchen-neg")"
if printf '%s' "$out" | grep -q "section-forbidden"; then
  echo "dcheck-phrase-selftest: FAIL — das forbid-pattern greift bei offenem Haekchen (Negativ-Kontrolle)." >&2
  fail=1
fi

# Scoping spiegelnden Bestand: open/ ist mit einem OFFENEN Slice belegt (die
# open/-Regel hat ihre Kandidatenmenge), der gepruefte Slice liegt in
# in-progress/ mit [x] — legitim, er darf nicht melden.
printf '%s\n' "- [ ] Echte offene Aufgabe." > "$TMP/body-scope-open.txt"
printf '%s\n' "- [x] Mit der Arbeit abgehakt — hier legitim." > "$TMP/body-scope.txt"
mkdir -p "$TMP/haekchen-scope/docs/plan/planning/open" "$TMP/haekchen-scope/docs/plan/planning/in-progress"
cat > "$TMP/haekchen-scope/.d-check.yml" <<YAML
modules: [structure]
structure:
  - files: "docs/plan/planning/open/**/slice-*.md"
    section-pattern: '^#+ .*(DoD|Definition of Done)'
    forbid-pattern: '${HAEKCHEN_PATTERN}'
YAML
{ echo "# slice-901 fixture"; echo; echo "## DoD"; echo; cat "$TMP/body-scope-open.txt"; } > "$TMP/haekchen-scope/docs/plan/planning/open/slice-901-fixture.md"
{ echo "# slice-902 fixture"; echo; echo "## DoD"; echo; cat "$TMP/body-scope.txt"; } > "$TMP/haekchen-scope/docs/plan/planning/in-progress/slice-902-fixture.md"
git_init_fixture "$TMP/haekchen-scope"
out="$(run_dcheck "$TMP/haekchen-scope")"
if printf '%s' "$out" | grep -qE "section-forbidden|section-missing"; then
  echo "dcheck-phrase-selftest: FAIL — Scoping-Kontrolle rot: das forbid-pattern greift in in-progress/ oder die open/-Kandidatenmenge ist leer." >&2
  printf '%s\n' "$out" >&2
  fail=1
fi

cat > "$TMP/body-zitat.txt" <<'BODY'
- Zitat des Musters im Code-Block unten.

```
- [x] zitiertes Muster
```
BODY
setup_haekchen_fixture "$TMP/haekchen-zitat" "open" "$TMP/body-zitat.txt"
out="$(run_dcheck "$TMP/haekchen-zitat")"
if printf '%s' "$out" | grep -q "section-forbidden"; then
  echo "dcheck-phrase-selftest: FAIL — das forbid-pattern trifft das zitierte Muster im Code-Block (Zitat-Kontext, SL-004)." >&2
  fail=1
fi

# --- Muster 4: `structure`-Modul, Chronik-Phrasen in gelesenen Dateien ------

# Slice-191: die haeufige Schreibweise von Chronik in AGENTS.md / harness/
# docs/plan/planning/README.md ist greppbar (BEO-HARNESS/
# chronik-in-gelesenen-dateien, 3x: slice-103, slice-182, slice-187).
# VIER Kontrollen: POSITIV (die Phrase in AGENTS.md feuert) · NEGATIV (eine
# saubere Datei bleibt gruen) · ZITAT (Inline-Code-Nennung bleibt gruen,
# SL-004) · SCOPING (Chronik in einem Slice-Plan bleibt gruen — der
# Geltungsbereich ist AGENTS.md, harness/*.md und Planning-README).
# GRENZE: der Sensor prueft eine Phrase, nicht die Klasse.

CHRONIK_PATTERN="$(grep -B2 'forbid-pattern.*Bis slice' .d-check.yml | grep -m1 'forbid-pattern' | sed -e "s/.*forbid-pattern: '//" -e "s/'.*//")"
if [ -z "$CHRONIK_PATTERN" ]; then
  echo "dcheck-phrase-selftest: FAIL — Chronik-forbid-pattern nicht in .d-check.yml gefunden (fail-closed)." >&2
  exit 2
fi

setup_chronik_fixture() {  # $1 = Zielverz., $2 = Datei-Body
  local dir="$1" body="$2"
  rm -rf "$dir"
  mkdir -p "$dir/harness"
  cat > "$dir/.d-check.yml" <<YAML
modules: [structure]
structure:
  - files: "AGENTS.md"
    section-pattern: '^#'
    sections: each
    forbid-pattern: '${CHRONIK_PATTERN}'
YAML
  cat > "$dir/AGENTS.md" <<EOF
# test

## Sektion

$(cat "$body")
EOF
  git_init_fixture "$dir"
}

printf '%s\n' "Bis slice-99 stand hier etwas." > "$TMP/body-chronik-pos.txt"
setup_chronik_fixture "$TMP/chronik-pos" "$TMP/body-chronik-pos.txt"
out="$(run_dcheck "$TMP/chronik-pos")"
if ! printf '%s' "$out" | grep -q "section-forbidden"; then
  echo "dcheck-phrase-selftest: FAIL — die Chronik-Phrase feuert nicht mehr (Positiv-Kontrolle, chronik-in-gelesenen-dateien 3x)." >&2
  printf '%s\n' "$out" >&2
  fail=1
fi

printf '%s\n' "Eine saubere Aussage ueber den Zustand." > "$TMP/body-chronik-neg.txt"
setup_chronik_fixture "$TMP/chronik-neg" "$TMP/body-chronik-neg.txt"
out="$(run_dcheck "$TMP/chronik-neg")"
if printf '%s' "$out" | grep -q "section-forbidden"; then
  echo "dcheck-phrase-selftest: FAIL — das Chronik-Muster greift bei sauberer Aussage (Negativ-Kontrolle)." >&2
  fail=1
fi

cat > "$TMP/body-chronik-zitat.txt" <<'BODY'
Die Form `- [x]` und der Satz `Bis slice-99 stand hier etwas` sind in
Inline-Code zitiert.
BODY
setup_chronik_fixture "$TMP/chronik-zitat" "$TMP/body-chronik-zitat.txt"
out="$(run_dcheck "$TMP/chronik-zitat")"
if printf '%s' "$out" | grep -q "section-forbidden"; then
  echo "dcheck-phrase-selftest: FAIL — das Chronik-Muster trifft die zitierte Phrase (Zitat-Kontext, SL-004)." >&2
  fail=1
fi

# SCOPING: eine Chronik-Zeile in einem Slice-PLAN bleibt gruen — der
# Geltungsbereich ist auf AGENTS.md, harness/*.md und Planning-README
# beschraenkt; Slice-Plaene erzaehlen Chronik legitime (Idee- und
# Historie-Texte). Die Fixture nutzt die echte Konfigurationsform (nur die
# AGENTS.md-Regel) — der Plan mit Chronik-Zeile liegt ausserhalb des Geltungs-
# bereichs und bleibt gruen.
mkdir -p "$TMP/chronik-scope/docs/plan/planning/open" "$TMP/chronik-scope/harness"
# Die Fixture spiegelt die ECHTEN drei Regeln (AGENTS.md, harness/*.md,
# Planning-README) -- ein Chronik-Slice-Plan bleibt trotzdem gruen: er liegt
# ausserhalb aller drei Geltungsbereiche. Nur so pinnt die Kontrolle den
# echten Geltungsbereich; eine spaeter verbreiterte files:-Glob wuerde rot.
cat > "$TMP/chronik-scope/.d-check.yml" <<YAML
modules: [structure]
structure:
  - files: "AGENTS.md"
    section-pattern: '^#'
    sections: each
    forbid-pattern: '${CHRONIK_PATTERN}'
  - files: "harness/*.md"
    section-pattern: '^#'
    sections: each
    forbid-pattern: '${CHRONIK_PATTERN}'
  - files: "docs/plan/planning/README.md"
    section-pattern: '^#'
    sections: each
    forbid-pattern: '${CHRONIK_PATTERN}'
YAML
cat > "$TMP/chronik-scope/AGENTS.md" <<'BODY'
# test

## Sektion

Sauber.
BODY
mkdir -p "$TMP/chronik-scope/harness" "$TMP/chronik-scope/docs/plan/planning"
cat > "$TMP/chronik-scope/harness/conventions.md" <<'BODY'
# Konventionen (Fixture)

Sauber.
BODY
cat > "$TMP/chronik-scope/docs/plan/planning/README.md" <<'BODY'
# Planning (Fixture)

Sauber.
BODY
cat > "$TMP/chronik-scope/docs/plan/planning/open/slice-903-fixture.md" <<'BODY'
# slice-903 fixture

## Idee

Bis slice-99 stand hier etwas anderes — Chronik in einem Plan ist legitim.
BODY
git_init_fixture "$TMP/chronik-scope"
out="$(run_dcheck "$TMP/chronik-scope")"
if printf '%s' "$out" | grep -qE "section-forbidden|section-missing"; then
  echo "dcheck-phrase-selftest: FAIL — Scoping-Kontrolle rot: Chronik in einem Slice-Plan feuert oder die AGENTS.md-Kandidatenmenge ist leer." >&2
  printf '%s\n' "$out" >&2
  fail=1
fi

# --- KORPUS-SEITE: traegt der echte Bestand die Trigger-Phrase noch? -------
#
# Geprueft ist NICHTLEERHEIT: der Lauf ist rot, wenn die Kandidatenmenge des
# Moduls leer ist, und gruen bei jeder Groesse darueber. Die Zahl in der
# Erfolgs-Zeile wird bei jedem Lauf neu gezaehlt; sie ist Ausgabe, keine Zusage.
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

echo "dcheck-phrase-selftest ok: 12 Werkzeug-Kontrollen (4 Muster) und 1 Korpus-Kontrolle"
echo "  (reviews-Phrase auf einem DoD-Haken: ${korpus_reviews} flache(r) Slice(s) in ${done_dir})."
echo "  NICHT geprueft: die uebrigen phrasen-basierten Felder in .d-check.yml — fuer sie gibt es"
echo "  keinen belegten Ausfall, und ein Sensor ohne Anlass ist selbst eine Behauptung."
