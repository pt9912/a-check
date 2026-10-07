#!/usr/bin/env bash
# image-multiarch.sh — baut das Release-Bild für linux/amd64 und linux/arm64 als
# EINEN Image-Index (OCI-Archiv) und prüft ihn mit multiarch-check.sh
# (ADR-0043, SPEC-DIST-001). Kein Emulator: die Kompilier-Stufe läuft auf der
# Plattform des Bau-Rechners (Dockerfile, deps/build).
#
# Der Builder ist ein eigener docker-container-Builder: der Standard-Treiber
# `docker` schreibt mit dem klassischen Bildspeicher keinen Mehr-Plattform-Index.
# Sein BuildKit-Image ist digest-gepinnt (AC-QA-03); der Builder-Name trägt den
# Digest-Anfang, damit eine Pin-Hebung einen neuen Builder ergibt, statt still
# den alten weiterzunutzen.
#
# Attestierungen und SBOM sind abgeschaltet: der Index trägt genau die zwei
# Plattform-Bilder (SPEC-DIST-001).
set -euo pipefail

: "${DOCKER:=docker}"
: "${VERSION:?VERSION fehlt}"
: "${BUILDKIT_IMAGE:?BUILDKIT_IMAGE fehlt}"
: "${OUT:?OUT fehlt}"
: "${GO_VERSION:?GO_VERSION fehlt}"
: "${GOLANGCI_LINT_VERSION:?GOLANGCI_LINT_VERSION fehlt}"

digest="${BUILDKIT_IMAGE##*@sha256:}"
[ "$digest" != "$BUILDKIT_IMAGE" ] || { echo "image-multiarch: BUILDKIT_IMAGE ist nicht digest-gepinnt: $BUILDKIT_IMAGE" >&2; exit 2; }
BUILDER="a-check-multiarch-${digest:0:12}"

if ! "$DOCKER" buildx inspect "$BUILDER" >/dev/null 2>&1; then
  "$DOCKER" buildx create --name "$BUILDER" --driver docker-container \
    --driver-opt "image=$BUILDKIT_IMAGE" >/dev/null
fi

mkdir -p "$(dirname "$OUT")"
"$DOCKER" buildx build --builder "$BUILDER" \
  --platform linux/amd64,linux/arm64 \
  --target runtime \
  --provenance=false --sbom=false \
  --build-arg VERSION="$VERSION" \
  --build-arg GO_VERSION="$GO_VERSION" \
  --build-arg GOLANGCI_LINT_VERSION="$GOLANGCI_LINT_VERSION" \
  -o "type=oci,dest=$OUT" .

bash tools/multiarch-check.sh "$OUT" "$VERSION"
