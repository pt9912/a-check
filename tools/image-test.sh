#!/usr/bin/env bash
# image-test.sh — AC-FA-DIST-001 + AC-QA-02-Akzeptanz gegen das lokal gebaute
# Runtime-Image (slice-006). Stack-Vorbild d-check tools/image-test.sh.
#
#   (1) Happy:    `--print-mk` → includebares Fragment (A_CHECK_IMAGE +
#                 a-check-Target); nativ == Container byte-identisch.
#   (2) Boundary: `--print-config` → dekodierbares .a-check.yml-Gerüst,
#                 read-only-Mount, Exit 0 (schreibt nichts).
#   (3) Negative: `--print-mk --bogus` → Exit 2 (unbekanntes Flag).
#   (4) Scan:     Verstoß-Fixture → Befund + Exit 1; stdout/stderr/Exit
#                 nativ == Container byte-identisch (AC-QA-01/AC-QA-02).
#
# „Nativ" in einem Docker-only-Repo: das statische Binary wird aus dem Image
# extrahiert (docker cp) und direkt ausgeführt — kein Host-Go (AGENTS §3.1).
# Annahme: Host-Arch = Image-Arch (amd64); auf abweichenden Hosts bricht der
# Nativ-Lauf laut ab.
set -euo pipefail

IMG="${IMAGE:-a-check}:dev"
WORK="$(mktemp -d)"
trap 'rm -rf "$WORK"' EXIT

fail() { echo "image-test: FAIL — $1" >&2; exit 1; }

# Binary aus dem Runtime-Image extrahieren (identisches Artefakt).
cid="$(docker create "$IMG")"
docker cp -q "$cid":/a-check "$WORK/a-check"
docker rm "$cid" >/dev/null
chmod +x "$WORK/a-check"

# --- (1) Happy: --print-mk nativ vs. Container ------------------------------
mk_n=0; "$WORK/a-check" --print-mk >"$WORK/mk.n.out" 2>"$WORK/mk.n.err" || mk_n=$?
mk_c=0; docker run --rm --network none "$IMG" --print-mk >"$WORK/mk.c.out" 2>"$WORK/mk.c.err" || mk_c=$?
[ "$mk_n" -eq 0 ] || fail "--print-mk nativ: Exit $mk_n, want 0"
[ "$mk_c" -eq 0 ] || fail "--print-mk Container: Exit $mk_c, want 0"
cmp -s "$WORK/mk.n.out" "$WORK/mk.c.out" || fail "--print-mk stdout nativ vs. Container nicht byte-identisch (AC-QA-02)"
grep -q 'A_CHECK_IMAGE' "$WORK/mk.c.out" || fail "--print-mk: A_CHECK_IMAGE fehlt"
grep -qE '^a-check:' "$WORK/mk.c.out" || fail "--print-mk: a-check-Target fehlt"
echo "image-test: (1) Happy — --print-mk nativ == Container, A_CHECK_IMAGE + Target vorhanden"

# --- (2) Boundary: --print-config, read-only-Mount → Exit 0 -----------------
mkdir -p "$WORK/ro"
pc_n=0; "$WORK/a-check" --print-config >"$WORK/pc.n.out" 2>/dev/null || pc_n=$?
pc_c=0; docker run --rm --network none -v "$WORK/ro":/src:ro "$IMG" --print-config >"$WORK/pc.c.out" 2>"$WORK/pc.c.err" || pc_c=$?
[ "$pc_n" -eq 0 ] || fail "--print-config nativ: Exit $pc_n, want 0"
[ "$pc_c" -eq 0 ] || fail "--print-config Container (ro): Exit $pc_c, want 0 (stderr: $(cat "$WORK/pc.c.err"))"
cmp -s "$WORK/pc.n.out" "$WORK/pc.c.out" || fail "--print-config stdout nativ vs. Container nicht byte-identisch"
grep -q 'version: 1' "$WORK/pc.c.out" || fail "--print-config: .a-check.yml-Gerüst unerwartet"
echo "image-test: (2) Boundary — --print-config read-only, Exit 0, dekodierbares Gerüst"

# --- (3) Negative: unbekanntes Flag → Exit 2 -------------------------------
neg=0; docker run --rm --network none "$IMG" --print-mk --bogus >/dev/null 2>"$WORK/neg.err" || neg=$?
[ "$neg" -eq 2 ] || fail "unbekanntes Flag: Exit $neg, want 2"
echo "image-test: (3) Negative — unbekanntes Flag, Exit 2"

# --- (4) Scan: Verstoß-Fixture, nativ == Container, Exit 1 ------------------
mkdir -p "$WORK/fix/internal/core" "$WORK/fix/internal/adapters/svc"
cat >"$WORK/fix/.a-check.yml" <<'YML'
version: 1
languages:
  go: ["**/*.go"]
layers:
  core: ["internal/core/**"]
  adapters: ["internal/adapters/**"]
edges:
  - {from: adapters, to: core}
composition_root: ["cmd/**"]
YML
cat >"$WORK/fix/internal/core/x.go" <<'GO'
package core

import _ "fix/internal/adapters/svc"
GO
printf 'package svc\n' >"$WORK/fix/internal/adapters/svc/svc.go"

sc_n=0; "$WORK/a-check" "$WORK/fix" >"$WORK/sc.n.out" 2>"$WORK/sc.n.err" || sc_n=$?
sc_c=0; docker run --rm --network none -v "$WORK/fix":/src:ro "$IMG" /src >"$WORK/sc.c.out" 2>"$WORK/sc.c.err" || sc_c=$?
[ "$sc_n" -eq 1 ] || fail "Scan nativ: Exit $sc_n, want 1 (stderr: $(cat "$WORK/sc.n.err"))"
[ "$sc_c" -eq 1 ] || fail "Scan Container: Exit $sc_c, want 1"
cmp -s "$WORK/sc.n.out" "$WORK/sc.c.out" || fail "Scan stdout nativ vs. Container nicht byte-identisch (AC-QA-02)"
cmp -s "$WORK/sc.n.err" "$WORK/sc.c.err" || fail "Scan stderr nativ vs. Container nicht byte-identisch"
grep -q 'core-impurity' "$WORK/sc.c.out" || fail "Scan: erwarteter core-impurity-Befund fehlt"
echo "image-test: (4) Scan — Verstoß erkannt, nativ == Container, Exit 1"

echo "image-test: OK — AC-FA-DIST-001 + AC-QA-02-Akzeptanz erfüllt"
