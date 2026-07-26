# Makefile — a-check
#
# Docker/make-only (AGENTS.md §3.1): kein Host-Go, keine Host-Paketmanager.
# Jede Gate ist eine Dockerfile-Stage (Muster: d-check/u-boot); der Host
# braucht nur git, make, bash, Docker.

include d-check.mk

GO_VERSION            ?= 1.26.4
GOLANGCI_LINT_VERSION ?= v2.12.2
IMAGE                 ?= a-check

# VERSION fließt ins OCI-Label org.opencontainers.image.version (Dockerfile
# runtime-Stage). Die Release-Pipeline (welle-05) setzt sie aus dem Git-Tag
# (make ci VERSION=…); lokal der dev-Default.
VERSION               ?= 0.0.0-dev

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

.PHONY: help compile lint test coverage-gate build arch-check gate-consistency guard-selftest record-gates gates image-test ci trace-check hooks

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
	$(DOCKER_BUILD) --build-arg VERSION=$(VERSION) -t $(IMAGE):dev .

arch-check: build ## Eigen-Architektur via a-check selbst (Dogfooding, AC-QA-02).
	docker run --rm --network none -v "$(CURDIR)":/src:ro $(IMAGE):dev /src

arch-graph: build ## Architektur-Graph (Mermaid) der eigenen .a-check.yml auf stdout (Dogfooding, netzlos, read-only).
	docker run --rm --network none -v "$(CURDIR)":/src:ro $(IMAGE):dev --print-graph /src

gate-consistency: ## Meta-Gate: dokumentierte Targets ↔ Makefile, .d-check.yml-Module (Harness-Lügen-Schutz).
	@bash tools/gate-consistency.sh

guard-selftest: ## Selbsttest des PreToolUse-Command-Guard (Denylist greift, Host-Toolchain blockiert).
	@bash .claude/hooks/pretooluse-command-guard.sh --selftest

suppression-check: ## Hard Rule AGENTS §3.2: keine Inline-Suppression (//nolint) in den Go-Quellen.
	@bash tools/suppression-check.sh

regelwerk-check: ## Wartung (KEIN Gate): Integritaet der vendored Baseline gegen SHA256SUMS (MR-006); Freshness bleibt ungeprueft.
	@bash tools/regelwerk-check.sh

verify-closure-notes: ## Struktur der Closure-Notizen in done/ (AGENTS §5): genau eine, ausgefuellt, kein Platzhalter.
	@bash tools/verify-closure-notes.sh

verify-slice-form: ## Form der Slice-Plaene ab slice-052 (max. 3 DoD-Punkte, benannte Lerneintrag-Form); aeltere grandfathered.
	@bash tools/verify-slice-form.sh

verify-ac-form: ## Form neuer Akzeptanzkriterien (Happy/Boundary/Negative + Out-of-Scope); die 19 bestehenden grandfathered.
	@bash tools/verify-ac-form.sh

verify-slice-links: ## Verweise wandernder Slices ueberleben den Lifecycle-Wechsel (SL-002); done/ ist Endzustand und ausgenommen.
	@bash tools/verify-slice-links.sh

# Die drei Teil-Sensoren laufen als Sequenz im selben Rezept, NICHT als
# Prerequisites: make bricht sonst beim ersten roten Target ab, und wer zwei
# Verstoesse in verschiedenen Bereichen hat, sieht nur den ersten (Review
# 2026-07-26, R-052-F4). Eine Verifikations-Schicht beantwortet drei
# unabhaengige Fragen — sie soll alle drei beantworten, auch wenn eine mit Nein
# ausgeht. Die Einzel-Targets bleiben fuer den gezielten Aufruf bestehen.
verify: ## Verifikations-Schicht: DoD-/Closure-Fragen (vor der "fertig"-Meldung; getrennt von gates).
	@fail=0; \
	bash tools/verify-closure-notes.sh || fail=1; \
	bash tools/verify-slice-form.sh    || fail=1; \
	bash tools/verify-ac-form.sh       || fail=1; \
	bash tools/verify-slice-links.sh   || fail=1; \
	if [ "$$fail" -ne 0 ]; then \
	  echo "[verify] FAIL — mindestens eine Verifikations-Frage ist offen; alle Befunde stehen oben." >&2; \
	  exit 1; \
	fi; \
	echo "[verify] Verifikations-Schicht gruen"

record-gates: ## Gate-Nachweis (Working-Tree-Hash) für den Stop-Hook schreiben.
	@bash tools/harness/record-gates.sh

gates: lint test coverage-gate arch-check doc-check gate-consistency suppression-check guard-selftest record-gates ## alle inneren Gates (mandatory vor Handoff).

image-test: build ## AC-FA-DIST-001 + nativ==Container-Akzeptanz gegen das gebaute Image.
	@IMAGE=$(IMAGE) bash tools/image-test.sh

ci: gates image-test ## CI-äquivalenter Lauf: gates + image-test (AC-FA-DIST-001).
	@echo "[ci] gates + image-test grün"

# FOCUS_COMMITS wählt alle übrigen .d-check.yml-Module ab, sodass nur `commits`
# läuft (sonst feuern die Datei-Module auf den Arbeitsbaum). ADR-0021.
DCHECK_RUN_I := docker run --rm -i --network none -v "$(CURDIR):/repo:ro" $(DCHECK_REF)
FOCUS_COMMITS := --enable commits --disable links --disable anchors --disable ids --disable matrix --disable external --disable codepaths --disable spans --disable hostpaths --disable diagrams --disable versions --disable pins --disable immutable --disable vcs --disable planning --disable tracked

trace-check: ## Traceability via Modul commits (ADR-0021): AC-/ADR-/MR-/slice-ID je Commit. MSGFILE=<datei> (Hook), RANGE=a..b (CI), sonst HEAD~1..HEAD. AGENTS §5.
	@$(if $(MSGFILE),$(DCHECK_RUN_I) --commit-msg - < $(MSGFILE),docker run --rm --network none -v "$(CURDIR):/repo:ro" $(DCHECK_REF) $(FOCUS_COMMITS) --range $(if $(RANGE),$(RANGE),HEAD~1..HEAD))

hooks: ## git-Hooks installieren (core.hooksPath -> .githooks; commit-msg Traceability). AGENTS §5.
	@git config core.hooksPath .githooks
	@echo "[hooks] core.hooksPath=.githooks — commit-msg Traceability-Gate aktiv"
