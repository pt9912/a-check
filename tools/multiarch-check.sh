#!/usr/bin/env bash
# multiarch-check.sh — prüft ein OCI-Archiv des Release-Bilds gegen
# SPEC-DIST-001 (ADR-0043):
#   (1) der Index führt GENAU zwei Einträge: linux/amd64 und linux/arm64 — keine
#       Attestierung, keine weitere Plattform;
#   (2) je Eintrag nennt die Config os/architecture, das Paar ist eine der zwei, und
#       die platform-Beschriftung des Index-Eintrags sagt dasselbe — nach ihr wählt
#       Docker beim Ziehen;
#   (3) je Eintrag ist das Binary ein ELF dieser Plattform (e_machine 62 bzw. 183);
#   (4) je Eintrag steht org.opencontainers.image.version auf der erwarteten Version.
# Grenze: geprüft wird das ARCHIV, nicht ein Lauf — ob das arm64-Binary auf arm64
# ausführt, belegt der Image-Test auf einem arm64-Rechner, nicht dieses Skript.
#
# Nur Host-Werkzeuge (bash, tar, od, sed, grep), kein JSON-Parser. Darum wird die
# Manifest-Liste NICHT in Einträge zerschnitten (ein Array-Feld in einem Eintrag
# schnitte sie falsch), sondern zweimal verschieden gezählt: die Digests unter dem
# Schlüssel "digest" und die Manifest-Medientypen. Beide Zahlen müssen 2 sein. Die
# Plattform eines Eintrags liest das Skript aus seiner Config und hält die
# Beschriftung im Index dagegen. Die Beschriftung sucht es zwischen dem Digest des
# Eintrags und dem nächsten Digest — BuildKit schreibt `platform` nach `digest`;
# eine andere Schlüssel-Reihenfolge findet die falsche oder keine und ist rot.
# Liefert ein Muster nichts, erreicht der leere Wert eine Prüfung mit Meldung.
#
# Aufruf: multiarch-check.sh <oci-archiv.tar> <version>
set -euo pipefail

ARCHIV="${1:?Aufruf: multiarch-check.sh <oci-archiv.tar> <version>}"
WANT_VERSION="${2:?Aufruf: multiarch-check.sh <oci-archiv.tar> <version>}"
WANT_PLATFORMS="linux/amd64 linux/arm64"

WORK="$(mktemp -d)"
trap 'rm -rf "$WORK"' EXIT
fail() { echo "multiarch-check: FAIL — $1" >&2; exit 1; }

tar -xf "$ARCHIV" -C "$WORK" 2>/dev/null || fail "$ARCHIV ist kein lesbares tar-Archiv"
blob() {
  local f="$WORK/blobs/sha256/${1#sha256:}"
  [ -f "$f" ] || fail "Blob $1 fehlt im Archiv"
  tr -d '\n\t ' <"$f"
}
# feld <json> <muster> — erster Treffer oder leer, nie ein Abbruch: ein grep ohne
# Treffer endet mit 1, und das beendet unter pipefail und set -e das Skript ohne
# Meldung. Leer erreicht die Prüfung darunter, und die nennt den Befund.
feld() { printf '%s' "$1" | grep -oE "$2" | sed -n 1p || true; }
alle() { printf '%s' "$1" | grep -oE "$2" || true; }
wert() { feld "$1" "\"$2\":\"[^\"]*\"" | sed 's/.*:"\(.*\)"/\1/'; }
digest_in() { feld "$1" '"digest":"sha256:[0-9a-f]{64}"' | sed 's/.*"\(sha256:[0-9a-f]*\)"/\1/'; }

# index.json des Layouts zeigt auf genau einen Image-Index.
[ -f "$WORK/index.json" ] || fail "$ARCHIV hat kein index.json (kein OCI-Layout)"
top="$(digest_in "$(tr -d '\n\t ' <"$WORK/index.json")")"
[ -n "$top" ] || fail "index.json nennt keinen Digest"
liste="$(blob "$top")"
case "$liste" in *'"mediaType":"application/vnd.oci.image.index.v1+json"'*) ;; *) fail "$top ist kein OCI-Image-Index" ;; esac

# Zähler 1: Digests unter dem Schlüssel "digest" (Annotationen tragen Digests unter
# anderen Schlüsseln und zählen nicht). Zähler 2: Manifest-Medientypen.
digests="$(alle "$liste" '"digest":"sha256:[0-9a-f]{64}"' | sed 's/.*"\(sha256:[0-9a-f]*\)"/\1/')"
n_digest="$(printf '%s\n' "$digests" | grep -c . || true)"
n_typ="$(alle "$liste" '"mediaType":"application/vnd\.(oci\.image\.manifest\.v1|docker\.distribution\.manifest\.v2)\+json"' | grep -c . || true)"
[ "$n_digest" = 2 ] && [ "$n_typ" = 2 ] \
  || fail "der Index führt $n_digest Digest(s) und $n_typ Manifest-Medientyp(en) — erwartet genau 2 und 2 ('$WANT_PLATFORMS')"

gefunden=""
for d in $digests; do
  man="$(blob "$d")"
  cfg_d="$(feld "$(printf '%s' "$man" | sed 's/.*"config":{\([^}]*\)}.*/\1/')" 'sha256:[0-9a-f]{64}')"
  [ -n "$cfg_d" ] || fail "Manifest $d ohne Config"
  cfg="$(blob "$cfg_d")"
  p="$(wert "$cfg" os)/$(wert "$cfg" architecture)"
  case " $WANT_PLATFORMS " in
    *" $p "*) ;;
    *) fail "Manifest $d hat die Plattform '$p' — erwartet ist eine aus '$WANT_PLATFORMS'" ;;
  esac
  gefunden="$gefunden $p"

  seg="${liste#*\"digest\":\"$d\"}"
  seg="${seg%%\"digest\":\"sha256:*}"
  plat="$(feld "$seg" '"platform":\{[^}]*\}')"
  label="$(wert "$plat" os)/$(wert "$plat" architecture)"
  [ "$label" = "$p" ] || fail "der Index beschriftet $d als '$label', die Config sagt '$p'"

  ver="$(wert "$cfg" 'org\.opencontainers\.image\.version')"
  [ "$ver" = "$WANT_VERSION" ] || fail "$p: Versions-Label '$ver', erwartet '$WANT_VERSION'"

  # Das Binary liegt im letzten Layer (COPY in der runtime-Stufe).
  layer="$(alle "$(printf '%s' "$man" | sed 's/.*"layers":\[//')" 'sha256:[0-9a-f]{64}' | tail -1)"
  [ -n "$layer" ] || fail "$p: Manifest ohne Layer"
  # Erst in eine Datei, dann lesen: ein `od -N2` am Ende einer Pipe schlösse sie
  # nach zwei Bytes, und `tar` stürbe unter pipefail an SIGPIPE.
  bin="$WORK/bin-${p//\//-}"
  tar -xzOf "$WORK/blobs/sha256/${layer#sha256:}" a-check >"$bin" 2>/dev/null \
    || fail "$p: der letzte Layer ist kein gzip-tar mit /a-check"
  machine="$(od -An -tu2 -j18 -N2 "$bin" | tr -d ' ')"
  case "$p" in
    linux/amd64) want=62 ;;
    linux/arm64) want=183 ;;
  esac
  [ "$machine" = "$want" ] || fail "$p: Binary hat ELF e_machine '$machine', erwartet '$want'"
  echo "multiarch-check: $p — Config, Versions-Label $ver, ELF e_machine $machine"
done

got="$(printf '%s\n' $gefunden | sort | tr '\n' ' ' | sed 's/ $//')"
[ "$got" = "$WANT_PLATFORMS" ] || fail "Plattformen '$got', erwartet genau '$WANT_PLATFORMS'"
echo "multiarch-check: ok — Index $top mit genau: $got"
