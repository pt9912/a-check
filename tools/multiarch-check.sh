#!/usr/bin/env bash
# multiarch-check.sh — prüft ein OCI-Archiv des Release-Bilds gegen
# SPEC-DIST-001 (ADR-0043):
#   (1) der Index nennt GENAU linux/amd64 und linux/arm64 — keine Attestierung,
#       keine weitere Plattform;
#   (2) je Plattform trägt die Config dieselbe Architektur und os=linux;
#   (3) je Plattform ist das Binary ein ELF der Plattform (e_machine 62 bzw. 183);
#   (4) je Plattform steht org.opencontainers.image.version auf der erwarteten Version.
# Grenze: geprüft wird das ARCHIV, nicht ein Lauf — ob das arm64-Binary auf arm64
# ausführt, belegt der Image-Test auf einem arm64-Rechner, nicht dieses Skript.
# Nur Host-Werkzeuge (bash, tar, od, sed, grep), kein JSON-Parser: das Archiv
# stammt aus BuildKit, dessen JSON kompakt ist; liefert ein Muster nichts, ist das
# ein Befund (fail-closed), kein leerer Erfolg. `sed -n 1p` statt `head -1`: `head`
# beendet den Schreiber vor ihm mit SIGPIPE, und pipefail macht daraus einen Abbruch.
#
# Aufruf: multiarch-check.sh <oci-archiv.tar> <version>
set -euo pipefail

ARCHIV="${1:?Aufruf: multiarch-check.sh <oci-archiv.tar> <version>}"
WANT_VERSION="${2:?Aufruf: multiarch-check.sh <oci-archiv.tar> <version>}"
WANT_PLATFORMS="linux/amd64 linux/arm64"

WORK="$(mktemp -d)"
trap 'rm -rf "$WORK"' EXIT
fail() { echo "multiarch-check: FAIL — $1" >&2; exit 1; }

tar -xf "$ARCHIV" -C "$WORK"
blob() { tr -d '\n\t ' <"$WORK/blobs/sha256/${1#sha256:}"; }
# feld <json> <muster> — erster Treffer oder leer, nie ein Abbruch: ein grep ohne
# Treffer endet mit 1, und das beendet unter pipefail und set -e das Skript ohne
# Meldung. Leer erreicht die Prüfung darunter, und die nennt den Befund.
feld() { printf '%s' "$1" | grep -oE "$2" | sed -n 1p || true; }
wert() { feld "$1" "\"$2\":\"[^\"]*\"" | sed 's/.*:"\(.*\)"/\1/'; }
digest_in() { feld "$1" '"digest":"sha256:[0-9a-f]{64}"' | sed 's/.*"\(sha256:[0-9a-f]*\)"/\1/'; }

# index.json des Layouts zeigt auf genau einen Image-Index.
top="$(digest_in "$(tr -d '\n\t ' <"$WORK/index.json")")"
[ -n "$top" ] || fail "index.json nennt keinen Digest"
liste="$(blob "$top")"
case "$liste" in *'"mediaType":"application/vnd.oci.image.index.v1+json"'*) ;; *) fail "$top ist kein OCI-Image-Index" ;; esac

# Je Eintrag: Digest und Plattform. Ein Eintrag ohne Plattform (Attestierung) zählt
# als "unknown/unknown" und macht (1) rot.
eintraege="$(printf '%s' "$liste" | sed 's/^.*"manifests":\[//; s/\].*$//' | sed 's/},{"mediaType"/}\n{"mediaType"/g')"
[ -n "$eintraege" ] || fail "der Index führt keine Manifeste"
gefunden=""
while IFS= read -r e; do
  d="$(digest_in "$e")"
  os="$(wert "$e" os)"
  arch="$(wert "$e" architecture)"
  [ -n "$d" ] || fail "Index-Eintrag ohne Digest: $e"
  p="${os:-unknown}/${arch:-unknown}"
  gefunden="$gefunden $p"
  case " $WANT_PLATFORMS " in
    *" $p "*) ;;
    *) fail "der Index führt '$p' (Digest $d) — erwartet genau '$WANT_PLATFORMS'" ;;
  esac

  man="$(blob "$d")"
  cfg_d="$(feld "$(printf '%s' "$man" | sed 's/.*"config":{\([^}]*\)}.*/\1/')" 'sha256:[0-9a-f]{64}')"
  [ -n "$cfg_d" ] || fail "$p: Manifest ohne Config"
  cfg="$(blob "$cfg_d")"
  c_arch="$(wert "$cfg" architecture)"
  c_os="$(wert "$cfg" os)"
  [ "$c_os/$c_arch" = "$p" ] || fail "$p: Config sagt $c_os/$c_arch"
  ver="$(wert "$cfg" 'org\.opencontainers\.image\.version')"
  [ "$ver" = "$WANT_VERSION" ] || fail "$p: Versions-Label '$ver', erwartet '$WANT_VERSION'"

  # Das Binary liegt im letzten Layer (COPY in der runtime-Stufe).
  layer="$(printf '%s' "$man" | sed 's/.*"layers":\[//' | { grep -oE 'sha256:[0-9a-f]{64}' || true; } | tail -1)"
  [ -n "$layer" ] || fail "$p: Manifest ohne Layer"
  # Erst in eine Datei, dann lesen: ein `od -N2` am Ende einer Pipe schlösse sie
  # nach zwei Bytes, und `tar` stürbe unter pipefail an SIGPIPE.
  bin="$WORK/bin-${p//\//-}"
  tar -xzOf "$WORK/blobs/sha256/${layer#sha256:}" a-check >"$bin"
  machine="$(od -An -tu2 -j18 -N2 "$bin" | tr -d ' ')"
  case "$p" in
    linux/amd64) want=62 ;;
    linux/arm64) want=183 ;;
    *) want="" ;;
  esac
  [ -n "$want" ] && [ "$machine" = "$want" ] || fail "$p: Binary hat ELF e_machine '$machine', erwartet '$want'"
  echo "multiarch-check: $p — Config, Versions-Label $ver, ELF e_machine $machine"
done <<<"$eintraege"

got="$(printf '%s\n' $gefunden | sort | tr '\n' ' ' | sed 's/ $//')"
[ "$got" = "$WANT_PLATFORMS" ] || fail "Plattformen '$got', erwartet genau '$WANT_PLATFORMS'"
echo "multiarch-check: ok — Index $top mit genau: $got"
