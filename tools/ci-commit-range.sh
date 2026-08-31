#!/usr/bin/env bash
# ci-commit-range.sh — ermittelt die Commit-Range, gegen die die CI ihre drei
# Range-Pruefungen faehrt (`trace-check`, `commit-scope-check`, `doc-immutable`).
#
# WARUM ALS SKRIPT UND NICHT INLINE IM WORKFLOW (slice-134): die Weiche traf eine
# Fallunterscheidung, die niemand ausfuehren konnte, ohne einen Push auf GitHub zu
# machen. Sie war dadurch drei Faelle lang unbelegt — und der dritte war falsch.
# Als Skript hat sie einen Selbsttest und laeuft lokal in `make gates`.
#
# DIE DREI LAGEN. Ein `pull_request` bringt seine Basis mit. Ein `push` bringt
# `github.event.before` — und das ist in ZWEI Lagen unbrauchbar, nicht in einer:
#
#   neuer Branch  -> all-zeros                      -> gegen den Default-Branch
#   Force-Push    -> gueltig aussehender SHA, den   -> gegen den Default-Branch
#                    der Runner-Klon NICHT kennt
#   normaler Push -> erreichbarer SHA               -> diese Range benutzen
#
# Der Force-Push-Fall ist der, an dem die Weiche bis slice-134 zerbrach: Dependabot
# rebast seine Branches und schiebt neu; der alte Commit bleibt serverseitig
# aufloesbar, ist aber von keiner Ref mehr erreichbar. `actions/checkout` holt ihn
# darum nicht, auch nicht mit `fetch-depth: 0` — die Tiefe betrifft die Historie
# der geholten Ref, nicht verwaiste Objekte. `d-check` meldete dann
# "Range-Basis nicht aufloesbar", Exit 2, auf jedem Rebase.
#
# GEPRUEFT WIRD ERREICHBARKEIT IM KLON, nicht die Form des Strings: `git cat-file`
# fragt genau das. Der Fallback misst MEHR Commits als der Push brachte, nie
# weniger — die sichere Richtung, und die Absicht der urspruenglichen Weiche
# ("sonst Silent-Green").
#
# NICHT GEPRUEFT (ehrliche Grenze): welchen Wert GitHub in `github.event.before`
# schreibt. Der Selbsttest baut sein eigenes Repo und misst dieses Skript; die
# Event-Semantik des fremden Systems erfaehrt nur ein echter Lauf. Genau diese
# Annahme war der Defekt, den slice-134 behebt — sie steht darum im
# Beobachtungs-Register, nicht in einer Zusage.
set -euo pipefail
cd "$(dirname "$0")/.."

# Erreichbarkeit im Klon — das Merkmal, auf das es ankommt.
resolvable() {  # $1 = SHA
  git cat-file -e "${1}^{commit}" 2>/dev/null
}

# Rein: keine Seiteneffekte, kein Netz. Damit im Selbsttest ausfuehrbar.
compute_range() {  # $1=EVENT $2=PR_BASE $3=PUSH_BEFORE $4=HEAD $5=DEFAULT_BRANCH
  local event="$1" pr_base="$2" before="$3" head="$4" default="$5"
  if [ "$event" = "pull_request" ]; then
    printf '%s..%s' "$pr_base" "$head"
    return
  fi
  if [[ "$before" =~ ^0+$ ]] || ! resolvable "$before"; then
    printf 'origin/%s..%s' "$default" "$head"
    return
  fi
  printf '%s..%s' "$before" "$head"
}

# ----------------------------------------------------------------- Selbsttest --
# Vier Faelle, und der vierte ist der wichtige: er haelt fest, dass eine
# BRAUCHBARE Basis auch benutzt wird. Ohne ihn waere ein Skript, das immer auf
# den Default-Branch faellt, von einem richtigen nicht zu unterscheiden.
self_test() {
  local tmp here fail=0
  here="$(pwd)"
  tmp="$(mktemp -d)"
  (
    cd "$tmp"
    git init -q .
    git config user.email t@t; git config user.name t
    printf 'a\n' > f; git add f; git commit -q -m erst
  )
  local reachable
  reachable="$(cd "$tmp" && git rev-parse HEAD)"
  local ghost="0123456789abcdef0123456789abcdef01234567"

  cd "$tmp"
  local got
  got="$(compute_range pull_request BASE ANY HEADSHA main)"
  [ "$got" = "BASE..HEADSHA" ] || {
    echo "ci-commit-range: Selbsttest FEHLGESCHLAGEN — pull_request nutzt nicht die PR-Basis ($got)" >&2; fail=1; }

  got="$(compute_range push '' 0000000000000000000000000000000000000000 HEADSHA main)"
  [ "$got" = "origin/main..HEADSHA" ] || {
    echo "ci-commit-range: Selbsttest FEHLGESCHLAGEN — neuer Branch faellt nicht auf den Default-Branch ($got)" >&2; fail=1; }

  got="$(compute_range push '' "$ghost" HEADSHA main)"
  [ "$got" = "origin/main..HEADSHA" ] || {
    echo "ci-commit-range: Selbsttest FEHLGESCHLAGEN — unerreichbare Basis (Force-Push) nicht abgefangen ($got)" >&2; fail=1; }

  got="$(compute_range push '' "$reachable" HEADSHA main)"
  [ "$got" = "$reachable..HEADSHA" ] || {
    echo "ci-commit-range: Selbsttest FEHLGESCHLAGEN — erreichbare Basis wird nicht benutzt ($got)" >&2; fail=1; }

  cd "$here"
  rm -rf "$tmp"
  [ "$fail" -eq 0 ] || exit 2
}

self_test

if [ "${1:-}" = "--selftest" ]; then
  echo "ci-range-selftest ok: vier Faelle (pull_request, neuer Branch, Force-Push, normaler Push)."
  echo "  NICHT geprueft: welchen Wert GitHub in github.event.before schreibt — das sagt nur ein echter Lauf."
  exit 0
fi

RANGE="$(compute_range \
  "${EVENT:-}" "${PR_BASE:-}" "${PUSH_BEFORE:-}" "${HEAD_SHA:-}" "${DEFAULT_BRANCH:-main}")"

# Der Default-Branch liegt im Runner-Klon nicht zwingend vor; nur im Fallback holen.
case "$RANGE" in
  origin/*)
    git fetch --no-tags origin \
      "+refs/heads/${DEFAULT_BRANCH:-main}:refs/remotes/origin/${DEFAULT_BRANCH:-main}" \
      >/dev/null 2>&1 || true
    ;;
esac

printf '%s\n' "$RANGE"
