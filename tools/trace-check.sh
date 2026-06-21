#!/usr/bin/env bash
# trace-check.sh — Traceability-Gate (AGENTS.md §5; harness/README.md
# §Traceability rules). Verlangt, dass jede Commit-Message mindestens eine
# AC-/ADR-/MR-/slice-ID nennt. EINE Wahrheit für drei Aufrufer (keine
# Logik-Dopplung):
#   - lokaler commit-msg-Hook:   trace-check.sh --message <COMMIT_MSG_FILE>
#   - PR-/Push-CI:               trace-check.sh --range <BASE>..<HEAD>
#   - lokal (make trace-check):  trace-check.sh            (Selbsttest + HEAD)
#   -                            trace-check.sh --self-test (nur Selbsttest)
#
# Ausgenommen (ID-frei erlaubt): Merge- und Revert-Commits (erste Zeile).
# fail-closed: unbekannter Modus / fehlende Argumente / kaputter Range brechen
# mit Exit 2 ab. Stack-Vorbild d-check tools/trace-check.sh.
set -euo pipefail

# Kennungs-Muster: die in .d-check.yml deklarierten ids (ADR/MR/AC) plus
# slice-NNN (Planning-ID). Eine Quelle der ID-Definition (AGENTS §5).
ID_RE='(ADR-[0-9]{4}|MR-[0-9]{3}|AC-(FA-[A-Z]+|QA)-[0-9]+|slice-[0-9]+)'

cd "$(git rev-parse --show-toplevel 2>/dev/null || echo .)"

is_exempt() { head -n1 <<<"$1" | grep -qE '^(Merge |Revert )'; }
has_id()    { grep -qE "$ID_RE" <<<"$1"; }

# Bereinigt eine Commit-Message wie git's strip-Cleanup: alles ab der
# scissors-Zeile (`# … >8 …`, verbose-Diff) weg, dann Kommentarzeilen (`#…`).
# In ALLEN Modi angewandt, damit Hook (--message) und CI (--range) GENAU
# dieselbe (kommentar-bereinigte) Message prüfen — eine ID muss auf einer
# Inhalts-Zeile stehen, nicht in einem #-Kommentar.
clean_message() { sed -e '/^#.*>8/,$d' -e '/^#/d'; }

# Prüft eine einzelne Message; 0 = ok, 1 = ID fehlt.
check_msg() { # $1 message, $2 label
  local msg="$1" label="$2"
  is_exempt "$msg" && return 0
  has_id "$msg" && return 0
  echo "trace-check: FAIL — $label nennt keine AC-/ADR-/MR-/slice-ID" >&2
  printf '  > %s\n' "$(head -n1 <<<"$msg")" >&2
  return 1
}

# Negativ-Selbsttest (analog tools/gate-consistency.sh): beweist bei jedem Lauf,
# dass das Gate eine fehlende ID auch wirklich fängt.
self_test() {
  check_msg "fix(x): siehe ADR-0001" "selftest-id" >/dev/null 2>&1 \
    || { echo "trace-check: Selbsttest FEHLGESCHLAGEN — ID nicht erkannt" >&2; exit 2; }
  if check_msg "chore: ohne bezug" "selftest-noid" >/dev/null 2>&1; then
    echo "trace-check: Selbsttest FEHLGESCHLAGEN — fehlende ID nicht erkannt" >&2; exit 2
  fi
  check_msg "Merge branch 'x'" "selftest-merge" >/dev/null 2>&1 \
    || { echo "trace-check: Selbsttest FEHLGESCHLAGEN — Merge nicht ausgenommen" >&2; exit 2; }
}

mode="${1:-}"
case "$mode" in
  --message)
    [ -n "${2:-}" ] && [ -f "$2" ] \
      || { echo "trace-check: --message braucht eine Message-Datei" >&2; exit 2; }
    self_test
    check_msg "$(clean_message <"$2")" "commit-msg"   # bei Erfolg still (Hook-Hygiene)
    ;;
  --range)
    range="${2:-}"
    [ -n "$range" ] || { echo "trace-check: --range braucht <base>..<head>" >&2; exit 2; }
    self_test
    base="${range%%..*}"
    # fail-closed: eine nicht auflösbare Basis (Zero-SHA bzw. fehlender Commit)
    # würde sonst still nur HEAD prüfen und ID-lose Zwischen-Commits durchlassen.
    # Der CI-Workflow liefert für neue Branches eine gültige Basis
    # (origin/<default-branch>); kommt dennoch keine an, blockieren wir LAUT.
    if [[ "$base" =~ ^0*$ ]] || ! git rev-parse -q --verify "${base}^{commit}" >/dev/null 2>&1; then
      echo "trace-check: FAIL — Range-Basis '$base' nicht auflösbar; der CI-Workflow muss eine gültige Basis liefern (z. B. origin/<default-branch>)." >&2
      exit 2
    fi
    commits="$(git rev-list --no-merges "$range")"
    fail=0 n=0
    while IFS= read -r sha; do
      [ -z "$sha" ] && continue
      n=$((n + 1))
      check_msg "$(git log -1 --format=%B "$sha" | clean_message)" "$sha" || fail=1
    done <<<"$commits"
    [ "$fail" -eq 0 ] && echo "trace-check: $n Commit(s) tragen eine AC-/ADR-/MR-/slice-ID (Selbsttest gefeuert)."
    exit "$fail"
    ;;
  --self-test)
    self_test
    echo "trace-check: Selbsttest grün."
    ;;
  "")
    self_test
    check_msg "$(git log -1 --format=%B HEAD | clean_message)" "HEAD"
    echo "trace-check: HEAD trägt eine AC-/ADR-/MR-/slice-ID (Selbsttest gefeuert)."
    ;;
  *)
    echo "trace-check: unbekannter Modus '$mode' (erwartet --message|--range|--self-test|<leer>)" >&2
    exit 2
    ;;
esac
