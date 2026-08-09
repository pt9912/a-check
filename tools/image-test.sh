#!/usr/bin/env bash
# image-test.sh — AC-FA-DIST-001 + AC-QA-02-Akzeptanz gegen das lokal gebaute
# Runtime-Image (slice-006). Stack-Vorbild d-check tools/image-test.sh.
#
#   (1) Happy:    `--print-mk` → includebares Fragment (A_CHECK_IMAGE +
#                 a-check/a-check-graph-Targets); nativ == Container == die
#                 committete `a-check.mk` byte-identisch (Fragment-Parität, slice-034).
#   (2) Boundary: `--print-config` → dekodierbares .a-check.yml-Gerüst,
#                 read-only-Mount, Exit 0 (schreibt nichts).
#   (2b) Boundary:`--print-graph` → dekodierbares Mermaid-flowchart aus einer
#                 read-only Config, Exit 0; nativ == Container (AC-FA-CLI-002).
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
ROOT="$(cd "$(dirname "$0")/.." && pwd)"   # Repo-Wurzel (cwd-unabhängig)
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
# Fragment-Paritaet (slice-034) — seit slice-083 OHNE die A_CHECK_IMAGE-Zeile:
# das ERZEUGTE Fragment traegt dort einen Platzhalter, die COMMITTETE
# a-check.mk den echten Release-Digest (ADR-0030). Das Binary kann den Digest
# seines eigenen Image nicht kennen; ein eingebackener Wert naehme immer den des
# Vorgaenger-Release. Alles ausserhalb dieser einen Zeile muss byte-identisch
# bleiben — sonst driftet das gelieferte Fragment vom committeten.
norm_mk() { grep -v '^A_CHECK_IMAGE ?=' "$1"; }
cmp -s <(norm_mk "$WORK/mk.c.out") <(norm_mk "$ROOT/a-check.mk") || fail "committete a-check.mk != --print-mk-Output ausserhalb der Pin-Zeile (Fragment-Paritaet slice-034/slice-083)"
grep -qE '^A_CHECK_IMAGE \?= ghcr\.io/pt9912/a-check@sha256:[0-9a-f]{64}$' "$ROOT/a-check.mk" || fail "committete a-check.mk traegt keinen vollen Release-Digest (AC-QA-03)"
grep -q 'SETZE-HIER-DEN-RELEASE-DIGEST-EIN' "$WORK/mk.c.out" || fail "--print-mk gibt keinen Platzhalter aus, sondern einen konkreten Wert (ADR-0030)"
grep -qE '@sha256:[0-9a-f]{64}' "$WORK/mk.c.out" && fail "--print-mk gibt einen konkreten Digest aus — er waere der des Vorgaenger-Release (ADR-0030)"
grep -q 'A_CHECK_IMAGE' "$WORK/mk.c.out" || fail "--print-mk: A_CHECK_IMAGE fehlt"
grep -qE '^a-check:' "$WORK/mk.c.out" || fail "--print-mk: a-check-Target fehlt"
grep -qE '^a-check-graph:' "$WORK/mk.c.out" || fail "--print-mk: a-check-graph-Target fehlt (AC-FA-DIST-001 0.20.0)"
grep -qF -- '--print-graph /src' "$WORK/mk.c.out" || fail "--print-mk: a-check-graph-Recipe ruft nicht --print-graph auf"
echo "image-test: (1) Happy — --print-mk nativ == Container == committete a-check.mk (Fragment-Parität), A_CHECK_IMAGE + a-check/a-check-graph-Targets vorhanden"

# --- (2) Boundary: --print-config, read-only-Mount → Exit 0 -----------------
mkdir -p "$WORK/ro"
pc_n=0; "$WORK/a-check" --print-config >"$WORK/pc.n.out" 2>/dev/null || pc_n=$?
pc_c=0; docker run --rm --network none -v "$WORK/ro":/src:ro "$IMG" --print-config >"$WORK/pc.c.out" 2>"$WORK/pc.c.err" || pc_c=$?
[ "$pc_n" -eq 0 ] || fail "--print-config nativ: Exit $pc_n, want 0"
[ "$pc_c" -eq 0 ] || fail "--print-config Container (ro): Exit $pc_c, want 0 (stderr: $(cat "$WORK/pc.c.err"))"
cmp -s "$WORK/pc.n.out" "$WORK/pc.c.out" || fail "--print-config stdout nativ vs. Container nicht byte-identisch"
grep -q 'version: 1' "$WORK/pc.c.out" || fail "--print-config: .a-check.yml-Gerüst unerwartet"
echo "image-test: (2) Boundary — --print-config read-only, Exit 0, dekodierbares Gerüst"

# --- (2b) Boundary: --print-graph, read-only-Mount → Exit 0 (AC-FA-CLI-002) -
cat >"$WORK/ro/.a-check.yml" <<'YML'
version: 1
languages:
  go: ["**/*.go"]
layers:
  core: ["internal/core/**"]
  ports: ["internal/ports/**"]
  adapters: ["internal/adapters/**"]
edges:
  - {from: adapters, to: ports}
  - {from: ports, to: core}
allow:
  - {from: ports, to: ports, reason: "Re-Export"}
YML
pg_n=0; "$WORK/a-check" --print-graph "$WORK/ro" >"$WORK/pg.n.out" 2>/dev/null || pg_n=$?
pg_c=0; docker run --rm --network none -v "$WORK/ro":/src:ro "$IMG" --print-graph /src >"$WORK/pg.c.out" 2>"$WORK/pg.c.err" || pg_c=$?
[ "$pg_n" -eq 0 ] || fail "--print-graph nativ: Exit $pg_n, want 0"
[ "$pg_c" -eq 0 ] || fail "--print-graph Container (ro): Exit $pg_c, want 0 (stderr: $(cat "$WORK/pg.c.err"))"
cmp -s "$WORK/pg.n.out" "$WORK/pg.c.out" || fail "--print-graph stdout nativ vs. Container nicht byte-identisch (AC-QA-01)"
grep -q 'flowchart TB' "$WORK/pg.c.out" || fail "--print-graph: Mermaid-flowchart-Kopf fehlt"
grep -qF -- '-.->|allow|' "$WORK/pg.c.out" || fail "--print-graph: gestrichelte allow-Kante fehlt"
echo "image-test: (2b) Boundary — --print-graph read-only, Exit 0, dekodierbares Mermaid"

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

echo "image-test: OK — AC-FA-DIST-001 + AC-FA-CLI-002 + AC-QA-02-Akzeptanz erfüllt"
