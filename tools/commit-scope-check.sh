#!/usr/bin/env bash
# commit-scope-check.sh — Sensor zur Commit-Scope-Regel (AGENTS.md §5), slice-062.
# Antwort auf SL-003: der Betreff bezeichnet nicht die enthaltene Arbeit.
#
# REGEL. Ein Commit mit Scope `(planning)` — `docs(planning)`, `fix(planning)`,
# `chore(planning)` — beruehrt ausschliesslich `docs/plan/planning/`.
#
# WARUM NUR DIESER SCOPE. Gemessen ueber die gesamte Historie: fuenf Treffer bei
# 74 `(planning)`-Commits, alle fuenf echte Diskrepanzen. Fuer `docs(...)`
# allgemein waeren es 31 bei 193 — `docs(spec)` aendert legitim `spec/`,
# `docs(adr)` legitim ADRs. Eine Regel, die den Bestand massenhaft bricht, wird
# abgeschaltet statt befolgt.
#
# GRANDFATHERING OHNE STICHTAGS-HASH. Jeder Commit wird an der Fassung gemessen,
# die zu SEINEM Zeitpunkt galt: der Sensor prueft, ob `AGENTS.md` im jeweiligen
# Commit die Regel bereits traegt, und ueberspringt ihn sonst. Das ist praeziser
# als ein Datum und braucht keinen Hash, den es beim Schreiben noch nicht gibt.
#
# NICHT GEPRUEFT (ehrliche Grenze, AC-QA-02): ob der Betreff die Arbeit inhaltlich
# TREFFEND beschreibt. Der Sensor prueft die Scope/Pfad-Konsistenz, nicht die
# Formulierung — das bleibt Review-Sache.
#
# Aufruf: make commit-scope-check [RANGE=<a>..<b>]   (Default: HEAD~1..HEAD)
set -euo pipefail
cd "$(dirname "$0")/.."

RANGE="${RANGE:-HEAD~1..HEAD}"
MARKER='Commit-Scope `(planning)`'      # Regel-Anker in AGENTS.md
SCOPE_RE='^[a-z]+\(planning\):'
ERLAUBT='^docs/plan/planning/'

# Galt die Regel zum Zeitpunkt dieses Commits?
regel_galt() {  # $1 = sha
  git show "$1:AGENTS.md" 2>/dev/null | grep -qF "$MARKER"
}

# Range aufloesen — FAIL-CLOSED (slice-069, Fund F-5). Frueher stand das
# `git rev-list` direkt in der `for`-Kopf-Substitution: deren Exit-Code ist fuer
# die Schleife bedeutungslos, ein unaufloesbarer Range lieferte einfach null
# Iterationen und der Lauf meldete "0 Commit(s) geprueft", Exit 0.
# LEERER Range bleibt gueltig: `git rev-list` exitet dann 0 ohne Ausgabe, und
# genau das ist der Unterschied zwischen "nichts zu pruefen" und "nicht
# aufloesbar".
resolve_range() {  # $1 = Range -> SHAs auf stdout; !=0 wenn nicht aufloesbar
  git rev-list "$1" 2>/dev/null
}

# Befunde eines Commits auf stdout; Rueckgabe 1 bei Befund.
check_commit() {  # $1 = sha
  local sha="$1" subj fremd
  subj="$(git log -1 --format=%s "$sha")"
  printf '%s' "$subj" | grep -qE "$SCOPE_RE" || return 0
  regel_galt "$sha" || return 0          # grandfathered
  fremd="$(git show --stat --format='' --name-only "$sha" | grep -v "$ERLAUBT" | grep -v '^$' || true)"
  if [ -n "$fremd" ]; then
    echo "$sha: Scope (planning), aber ausserhalb von docs/plan/planning/ geaendert:"
    printf '%s\n' "$fremd" | sed 's/^/    /'
    echo "    -> eigener Commit mit passendem Scope (AGENTS.md §5, SL-003)"
    return 1
  fi
  return 0
}

# Selbsttest gegen ein totes Muster. Die Negativ-Fixture trifft das Muster
# BEINAHE — ein Betreff mit anderem Scope, der dieselben Pfade beruehrt (Lehre
# aus slice-058); zusaetzlich ein Betreff, der "(planning)" nur im Freitext
# fuehrt und darum nicht greifen darf (SL-004: zitiertes Muster ist kein Muster).
self_test() {
  local s
  for s in 'docs(planning): x' 'fix(planning): x' 'chore(planning): x'; do
    printf '%s' "$s" | grep -qE "$SCOPE_RE" || {
      echo "commit-scope-check: Selbsttest FEHLGESCHLAGEN — '$s' nicht als (planning) erkannt" >&2; exit 2; }
  done
  for s in 'feat(harness): x' 'docs(spec): x' 'docs: erwaehnt (planning) im Text'; do
    if printf '%s' "$s" | grep -qE "$SCOPE_RE"; then
      echo "commit-scope-check: Selbsttest FEHLGESCHLAGEN — '$s' faelschlich als (planning) erkannt" >&2; exit 2
    fi
  done
  grep -qF "$MARKER" AGENTS.md || {
    echo "commit-scope-check: Selbsttest FEHLGESCHLAGEN — Regel-Anker '$MARKER' fehlt in AGENTS.md;" >&2
    echo "  der Sensor wuerde jeden Commit als grandfathered ueberspringen (stilles False-Green)." >&2
    exit 2; }
  # Range-Aufloesung, beide Richtungen: Muell muss scheitern, Gueltiges muss
  # durchgehen. Ohne die zweite Haelfte waere ein resolve_range, das IMMER
  # scheitert, von einem korrekten nicht zu unterscheiden.
  if resolve_range 'definitely-not-a-revision' >/dev/null 2>&1; then
    echo "commit-scope-check: Selbsttest FEHLGESCHLAGEN — unaufloesbarer Range nicht erkannt;" >&2
    echo "  der Sensor wuerde '0 Commit(s) geprueft' mit Exit 0 melden (F-5)." >&2
    exit 2
  fi
  resolve_range 'HEAD' >/dev/null 2>&1 || {
    echo "commit-scope-check: Selbsttest FEHLGESCHLAGEN — gueltiger Range faelschlich abgelehnt" >&2
    exit 2; }
}

self_test

if ! shas="$(resolve_range "$RANGE")"; then
  echo "commit-scope-check: FAIL — Range '$RANGE' ist nicht aufloesbar (git rev-list)." >&2
  echo "  Ein nicht aufloesbarer Range ist kein leerer Range: hier wurde nichts geprueft." >&2
  exit 1
fi

fail=0; geprueft=0; alt=0
for sha in $shas; do
  if printf '%s' "$(git log -1 --format=%s "$sha")" | grep -qE "$SCOPE_RE"; then
    if regel_galt "$sha"; then geprueft=$((geprueft + 1)); else alt=$((alt + 1)); continue; fi
  fi
  check_commit "$sha" || fail=1
done

if [ "$fail" -ne 0 ]; then
  echo "commit-scope-check: FAIL — Commit-Scope und geaenderte Pfade passen nicht zusammen (AGENTS.md §5)." >&2
  exit 1
fi
echo "commit-scope-check ok: $geprueft (planning)-Commit(s) in $RANGE geprueft, $alt vor Einfuehrung der Regel (grandfathered)."
