# syntax=docker/dockerfile:1.7
# Dockerfile — a-check (Muster: d-check/u-boot, gleiche Build-Familie).
# Jede Gate ist eine Stage (`docker build --target …`); Bases sind
# digest-gepinnt (AC-QA-03 Reproduzierbarkeit). Das Laufzeit-Image ist
# statisch gelinkt auf distroless/static (AC-QA-02, AC-FA-DIST-001).
ARG GO_VERSION=1.27.2
ARG GOLANGCI_LINT_VERSION=v2.13.2

# ---- deps ------------------------------------------------------------------
# Läuft auf der Plattform des Bau-Rechners ($BUILDPLATFORM), auch wenn das Ziel
# eine andere ist: die build-Stufe kompiliert per GOOS/GOARCH für die
# Ziel-Plattform, darum braucht der Mehr-Plattform-Bau keinen Emulator
# (ADR-0043). Bei einem Bau für die eigene Plattform ist beides dieselbe.
FROM --platform=$BUILDPLATFORM golang:${GO_VERSION}@sha256:5bc7f572bbaa98885a3a1fd9c0aa76b59e3e14e8628bfc316bbfd0c701e4818c AS deps
WORKDIR /src
ENV GOFLAGS="-mod=readonly -buildvcs=false" \
    GOMODCACHE=/go/pkg/mod \
    GOCACHE=/root/.cache/go-build
COPY go.mod ./
COPY go.su[m] ./
RUN mkdir -p "$GOMODCACHE" && go mod download

# ---- compile ---------------------------------------------------------------
FROM deps AS compile
COPY . .
RUN CGO_ENABLED=0 go build -o /tmp/a-check ./cmd/a-check

# ---- lint ------------------------------------------------------------------
FROM golangci/golangci-lint:${GOLANGCI_LINT_VERSION}@sha256:ba07dffad130794ae79ebaa0056809d18c0168f3f846480ffd3eb6c04578b83d AS lint
WORKDIR /src
ENV GOFLAGS="-buildvcs=false"
COPY --from=deps /go/pkg/mod /go/pkg/mod
COPY . .
RUN golangci-lint run ./...

# ---- test ------------------------------------------------------------------
FROM deps AS test
COPY . .
RUN CGO_ENABLED=0 go test ./...

# ---- coverage --------------------------------------------------------------
# Kalibrierungs-Bindung (harness/README.md §Sensors): Schwelle via
# COVERAGE_THRESHOLD; Verfehlung ⇒ Carveout-Pflicht (AGENTS.md §3.6).
# `-coverpkg` misst über die Paketgrenzen von ./internal/... — sonst
# zählt nur paket-lokale Abdeckung (Integrationstests bekommen keine
# Cross-Package-Gutschrift). `pipefail` via SHELL, damit `go test … | tee`
# den Exit-Code nicht maskiert.
FROM deps AS coverage
SHELL ["/bin/bash", "-eo", "pipefail", "-c"]
ARG COVERAGE_THRESHOLD=90
ENV COVERAGE_THRESHOLD=${COVERAGE_THRESHOLD}
COPY . .
RUN mkdir -p /out && \
    COVERPKG=$(go list ./internal/... | tr '\n' ',' | sed 's/,$//') && \
    CGO_ENABLED=0 go test \
        -coverpkg="$COVERPKG" \
        -coverprofile=/out/coverage.out \
        -covermode=atomic \
        ./... && \
    go tool cover -func=/out/coverage.out | tee /out/coverage-func.txt && \
    bash tools/coverage-gate.sh /out/coverage-func.txt "$COVERAGE_THRESHOLD"

# ---- build -----------------------------------------------------------------
# TARGETOS/TARGETARCH setzt BuildKit je Ziel-Plattform; die runtime-Stufe
# darunter führt keinen Befehl aus und bleibt damit emulator-frei.
FROM deps AS build
ARG TARGETOS TARGETARCH
COPY . .
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -trimpath -ldflags="-s -w" -o /out/a-check ./cmd/a-check

# ---- runtime ---------------------------------------------------------------
FROM gcr.io/distroless/static-debian12:nonroot@sha256:d093aa3e30dbadd3efe1310db061a14da60299baff8450a17fe0ccc514a16639 AS runtime
# VERSION fließt ins OCI-Label org.opencontainers.image.version; die spätere
# Release-Pipeline (welle-05-release) setzt sie aus dem Git-Tag und verifiziert
# die Übereinstimmung (make build VERSION=…). Default für lokale Builds.
ARG VERSION=0.0.0-dev
LABEL org.opencontainers.image.source="https://github.com/pt9912/a-check" \
      org.opencontainers.image.description="a-check — sprach-agnostischer Hexagonal-Architektur-Checker (text-heuristisch, netzlos)." \
      org.opencontainers.image.title="a-check" \
      org.opencontainers.image.vendor="pt9912" \
      org.opencontainers.image.licenses="MIT" \
      org.opencontainers.image.version="${VERSION}"
COPY --from=build /out/a-check /a-check
ENTRYPOINT ["/a-check"]
