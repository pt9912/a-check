# Makefile — a-check
#
# Docker/make-only (AGENTS.md §3.1): kein Host-Go, keine Host-Paketmanager.
# Jede Gate ist eine Dockerfile-Stage (Muster: d-check/u-boot); der Host
# braucht nur git, make, bash, Docker.

include d-check.mk

GO_VERSION            ?= 1.26.4
GOLANGCI_LINT_VERSION ?= v2.12.2
IMAGE                 ?= a-check

# Kalibrierungs-Bindung (harness/README.md §Sensors): 90 % seit
# 2026-06-21 (Bootstrap-Kalibrierung). Override: `make coverage-gate
# THRESHOLD=…`; Senkung nur per ADR (AGENTS.md §3.6, ADR-0006).
THRESHOLD ?= 90

PROGRESS_FLAG ?=
DOCKER_BUILD := docker build $(PROGRESS_FLAG) \
    --build-arg GO_VERSION=$(GO_VERSION) \
    --build-arg GOLANGCI_LINT_VERSION=$(GOLANGCI_LINT_VERSION)

# Gate-Stages werden bewusst nicht gecacht (sonst „grün" aus altem Layer).
NO_CACHE_FILTER_LINT := --no-cache-filter lint
NO_CACHE_FILTER_TEST := --no-cache-filter test
NO_CACHE_FILTER_COV  := --no-cache-filter coverage

.DEFAULT_GOAL := help

.PHONY: help compile lint test coverage-gate build arch-check gate-consistency record-gates gates

# Gates seriell: unter `make -j` liefen die Sub-Gates sonst parallel und die
# Reihenfolge/der Abbruch bei rotem Gate wären nicht garantiert.
.NOTPARALLEL:

help: ## Diese Hilfe anzeigen.
	@grep -hE '^[a-zA-Z_-]+:.*?## ' $(MAKEFILE_LIST) | sort | awk 'BEGIN{FS=":.*?## "}{printf "  \033[36m%-14s\033[0m %s\n",$$1,$$2}'

compile: ## Schnelles Compile-Feedback (ohne Tests/Lint).
	$(DOCKER_BUILD) --target compile -t $(IMAGE):compile .

lint: ## golangci-lint mit dem Projekt-Profil (§3.2, ADR-0005).
	$(DOCKER_BUILD) $(NO_CACHE_FILTER_LINT) --target lint -t $(IMAGE):lint .

test: ## go test ./... in Docker (AC-FA-*/AC-QA-01).
	$(DOCKER_BUILD) $(NO_CACHE_FILTER_TEST) --target test -t $(IMAGE):test .

coverage-gate: ## Coverage-Schwelle (Kalibrierungs-Bindung: $(THRESHOLD) %, ADR-0006).
	$(DOCKER_BUILD) $(NO_CACHE_FILTER_COV) \
	    --build-arg COVERAGE_THRESHOLD=$(THRESHOLD) \
	    --target coverage -t $(IMAGE):coverage .

build: ## a-check-Image bauen (static/distroless, digest-gepinnte Bases).
	$(DOCKER_BUILD) -t $(IMAGE):dev .

arch-check: build ## Eigen-Architektur via a-check selbst (Dogfooding, AC-QA-02).
	docker run --rm --network none -v "$(CURDIR)":/src:ro $(IMAGE):dev /src

gate-consistency: ## Meta-Gate: dokumentierte Targets ↔ Makefile, .d-check.yml-Module (Harness-Lügen-Schutz).
	@bash tools/gate-consistency.sh

record-gates: ## Gate-Nachweis (Working-Tree-Hash) für den Stop-Hook schreiben.
	@bash tools/harness/record-gates.sh

gates: lint test coverage-gate arch-check doc-check gate-consistency record-gates ## alle inneren Gates (mandatory vor Handoff).
