# Dockerfile — a-check (Muster: d-check/u-boot, gleiche Build-Familie).
# Jede Gate ist eine Stage (`docker build --target …`); Bases sind
# digest-gepinnt (AC-QA-03 Reproduzierbarkeit). Das Laufzeit-Image ist
# statisch gelinkt auf distroless/static (AC-QA-02, AC-FA-DIST-001).
ARG GO_VERSION=1.26.4
ARG GOLANGCI_LINT_VERSION=v2.12.2

# ---- deps ------------------------------------------------------------------
FROM golang:${GO_VERSION}@sha256:792443b89f65105abba56b9bd5e97f680a80074ac62fc844a584212f8c8102c3 AS deps
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
FROM golangci/golangci-lint:${GOLANGCI_LINT_VERSION}@sha256:5cceeef04e53efe1470638d4b4b4f5ceefd574955ab3941b2d9a68a8c9ad5240 AS lint
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
FROM deps AS build
COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/a-check ./cmd/a-check

# ---- runtime ---------------------------------------------------------------
FROM gcr.io/distroless/static-debian12:nonroot@sha256:d093aa3e30dbadd3efe1310db061a14da60299baff8450a17fe0ccc514a16639 AS runtime
COPY --from=build /out/a-check /a-check
ENTRYPOINT ["/a-check"]
