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

# PENDING-MODUS (commit-msg-Hook, slice-072): der Commit existiert noch NICHT.
# Scope kommt aus der Pending-Message, die Pfade aus dem git-INDEX statt aus
# `git show <sha>`. Ohne diesen Modus greift der Sensor erst in der CI — also
# nach dem Push, wo die Korrektur ein Rebase auf einem veroeffentlichten Branch
# waere. Real aufgetreten am 2026-08-09 (CI-Run 31301467076): drei
# (planning)-Commits, alle mit gruenem `make gates` davor.
# Die Regel-Gueltigkeit wird am ARBEITSBAUM geprueft (nicht an einem Commit) —
# es gibt noch keinen, an dem `git show <sha>:AGENTS.md` haengen koennte.
check_pending() {  # $1 = Message-Datei
  local subj fremd
  if [ ! -f "$1" ]; then
    echo "commit-scope-check: FAIL — Message-Datei '$1' nicht lesbar." >&2
    return 1
  fi
  subj="$(head -n1 "$1")"
  printf '%s' "$subj" | grep -qE "$SCOPE_RE" || return 0
  grep -qF "$MARKER" AGENTS.md || return 0        # Regel gilt hier noch nicht
  fremd="$(git diff --cached --name-only | fremde_pfade)"
  if [ -n "$fremd" ]; then
    echo "commit-scope-check: FAIL — Scope (planning), aber ausserhalb von docs/plan/planning/ gestagt:" >&2
    printf '%s\n' "$fremd" | sed 's/^/    /' >&2
    echo "    -> eigener Commit mit passendem Scope (AGENTS.md §5, SL-003)" >&2
    echo "    Der Commit ist NICHT entstanden; die Aufteilung kostet hier einen git reset," >&2
    echo "    nach einem Push einen Rebase auf veroeffentlichter Historie." >&2
    return 1
  fi
  return 0
}

# Pfadliste (stdin) -> die Pfade ausserhalb des erlaubten Bereichs.
# Eigene Funktion, damit BEIDE Modi (Range und Pending) dieselbe Regel
# anwenden und der Selbsttest sie ohne git pruefen kann (slice-072).
fremde_pfade() {
  grep -v "$ERLAUBT" | grep -v '^$' || true
}

# Befunde eines Commits auf stdout; Rueckgabe 1 bei Befund.
check_commit() {  # $1 = sha
  local sha="$1" subj fremd
  subj="$(git log -1 --format=%s "$sha")"
  printf '%s' "$subj" | grep -qE "$SCOPE_RE" || return 0
  regel_galt "$sha" || return 0          # grandfathered
  fremd="$(git show --stat --format='' --name-only "$sha" | fremde_pfade)"
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
  # Pfad-Filter, beide Richtungen (slice-072): er traegt jetzt BEIDE Modi, also
  # muss er ohne git pruefbar sein — und in beide Richtungen, sonst waere ein
  # Filter, der alles oder nichts durchlaesst, nicht zu unterscheiden.
  if [ -n "$(printf 'docs/plan/planning/open/x.md\n' | fremde_pfade)" ]; then
    echo "commit-scope-check: Selbsttest FEHLGESCHLAGEN — erlaubter Pfad als fremd gemeldet" >&2
    exit 2
  fi
  if [ -z "$(printf 'docs/reviews/y.md\n' | fremde_pfade)" ]; then
    echo "commit-scope-check: Selbsttest FEHLGESCHLAGEN — fremder Pfad nicht erkannt" >&2
    exit 2
  fi
}

self_test

# Pending-Modus vor dem Range-Modus: MSGFILE gewinnt, wie bei trace-check.
if [ -n "${MSGFILE:-}" ]; then
  check_pending "$MSGFILE" || exit 1
  echo "commit-scope-check ok: Pending-Commit gegen den Index geprueft (Selbsttests gefeuert)."
  exit 0
fi

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
