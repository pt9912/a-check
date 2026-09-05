# Makefile — a-check
#
# Docker/make-only (AGENTS.md §3.1): kein Host-Go, keine Host-Paketmanager.
# Jede Gate ist eine Dockerfile-Stage (Muster: d-check/u-boot); der Host
# braucht nur git, make, bash, Docker.

include d-check.mk

# Diese beiden gehen als --build-arg ins Dockerfile und stechen dort den
# ARG-Default. WIRKSAM ist keine von beiden: neben dem Tag steht ein Digest, und
# der sticht den Tag. Die Zahl benennt also, was im gepinnten Image liegt —
# gemessen 1.27.0 bzw. 2.13.2 — und `make version-coherence` haelt sie gegen die
# ARG-Defaults, damit die beiden Orte nicht auseinanderlaufen (slice-131).
GO_VERSION            ?= 1.27.0
GOLANGCI_LINT_VERSION ?= v2.13.2
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
# Container-Runtime ueber eine Indirektion — dieselbe, die das gelieferte
# a-check.mk seit slice-082 fuehrt. Vorher lief `docker build` ueber
# DOCKER_BUILD, `docker run` aber woertlich daneben: die Idee war da, sie war
# nur nie auf `run` ausgedehnt. Dogfooding: was a-check seinen Konsumenten
# liefert, gilt hier auch.
DOCKER ?= docker
DOCKER_BUILD := $(DOCKER) build $(PROGRESS_FLAG) \
    --build-arg GO_VERSION=$(GO_VERSION) \
    --build-arg GOLANGCI_LINT_VERSION=$(GOLANGCI_LINT_VERSION)

# Gate-Stages werden bewusst nicht gecacht (sonst „grün" aus altem Layer).
NO_CACHE_FILTER_LINT := --no-cache-filter lint
NO_CACHE_FILTER_TEST := --no-cache-filter test
NO_CACHE_FILTER_COV  := --no-cache-filter coverage

.DEFAULT_GOAL := help

# Jedes Rezept-Target gehoert hierher: fehlt es, ueberspringt make das Rezept,
# sobald eine gleichnamige Datei existiert — und meldet Exit 0, ohne etwas
# getan zu haben. `gate-consistency` prueft die Vollstaendigkeit (slice-068,
# Fund F-1 des unabhaengigen Reviews).
.PHONY: help compile lint test coverage-gate build arch-check arch-graph \
        gate-consistency guard-selftest ci-range-selftest record-gates gates image-test ci \
        trace-check hooks suppression-check regelwerk-check commit-scope-check \
        verify verify-risiko-ausgaenge verify-observations slice-mv image-scan \
        doc-workflows doc-reviews version-coherence archive-wave-test archive-wave

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
	$(DOCKER) run --rm --network none -v "$(CURDIR)":/src:ro $(IMAGE):dev /src

arch-graph: build ## Architektur-Graph (Mermaid) der eigenen .a-check.yml auf stdout (Dogfooding, netzlos, read-only).
	$(DOCKER) run --rm --network none -v "$(CURDIR)":/src:ro $(IMAGE):dev --print-graph /src

gate-consistency: ## Meta-Gate: dokumentierte Targets ↔ Makefile, .d-check.yml-Module (Harness-Lügen-Schutz).
	@bash tools/gate-consistency.sh

ci-range-selftest: ## Selbsttest der CI-Commit-Range-Weiche (vier Faelle, inkl. Force-Push).
	@bash tools/ci-commit-range.sh --selftest

guard-selftest: ## Selbsttest des PreToolUse-Command-Guard (Denylist greift, Host-Toolchain blockiert).
	@bash .claude/hooks/pretooluse-command-guard.sh --selftest

suppression-check: ## Hard Rule AGENTS §3.2: keine Inline-Suppression (//nolint) in den Go-Quellen.
	@bash tools/suppression-check.sh

regelwerk-check: ## Wartung (KEIN Gate): Integritaet der vendored Baseline gegen SHA256SUMS (MR-006); Freshness bleibt ungeprueft.
	@bash tools/regelwerk-check.sh

verify-risiko-ausgaenge: ## Jedes in §6 notierte Risiko traegt einen Ausgang aus der Dreier-Menge (AGENTS §5, ab slice-102).
	@bash tools/verify-risiko-ausgaenge.sh

verify-observations: ## Deckung des Beobachtungs-Registers: zitierte BEO-Kennung hat eine Zeile, jede Zeile traegt formgebundene Belege.
	@bash tools/verify-observations.sh

commit-scope-check: ## Commit-Scope (planning) beruehrt nur docs/plan/planning/ (AGENTS §5, SL-003). MSGFILE=<datei> (Hook, prueft den Index VOR dem Commit), RANGE=a..b (CI), sonst HEAD~1..HEAD.
	@MSGFILE="$(MSGFILE)" RANGE="$(RANGE)" bash tools/commit-scope-check.sh

# Eigene doc-*-Targets: das erzeugte Fragment d-check.mk fuehrt fuer diese
# Module keines, obwohl der Pin sie kennt (sie stehen in seinen
# --disable-Listen). Das Fragment bleibt darum unberuehrt (slice-130, slice-160).
#
# DISABLED WIRD DIE modules:-LISTE, NICHT DIE MODUL-MENGE: die Fragment-Targets
# zaehlen alle Module auf, weil sie generisch erzeugt sind und die Konfiguration
# nicht kennen. Hier waere das eine geschlossene Liste gegen eine offene Menge —
# genau die Falle aus slice-115. Die Namen unten sind die aktive
# modules:-Liste der .d-check.yml; sie steht im selben Repo und faellt beim
# Aendern auf.
doc-workflows: ## Deklarations-Form der uses:-Referenzen unter .github/workflows (Modul workflows).
	$(DOCKER) run --rm --network none -v "$(CURDIR):/repo:ro" $(DCHECK_REF) \
	  --enable workflows \
	  --disable links --disable anchors --disable ids \
	  --disable matrix --disable spans --disable hostpaths --disable reviews

# Review-Report-Deckung (DC-FA-RVW-001, Modul reviews, seit slice-160): eine
# DoD-Zeile mit der Phrase "unabhaengiger Review" braucht mindestens einen
# Report unter docs/reviews/ mit derselben slice-<NNN>-Kennung im Dateinamen.
# Opt-in PRO SLICE ueber die DoD-Phrase selbst — ohne sie ist die
# Kandidatenmenge fuer diesen Slice leer, kein Fehlalarm.
doc-reviews: ## Review-Report-Deckung fuer Slices mit DoD-Zusage "unabhaengiger Review" (Modul reviews).
	$(DOCKER) run --rm --network none -v "$(CURDIR):/repo:ro" $(DCHECK_REF) \
	  --enable reviews \
	  --disable links --disable anchors --disable ids \
	  --disable matrix --disable spans --disable hostpaths

# Antwort auf BEO-026 bei 3x (slice-131). Es prueft KOHAERENZ, nicht Wahrheit:
# derselbe uses:-SHA traegt ueberall denselben Tag-Kommentar, und eine
# Versions-Variable, die Makefile UND Dockerfile fuehren, hat an beiden Orten
# denselben Wert. Zwei uebereinstimmend falsche Angaben bleiben gruen — die
# Registry zu fragen waere Netz, und `gates` ist hermetisch.
version-coherence: ## Kohaerenz doppelt deklarierter Versions-Angaben (uses:-SHA <-> Tag-Kommentar, Makefile <-> Dockerfile-ARG).
	@bash tools/verify-versions-kohaerent.sh

# KEIN Bestandteil von `gates` — der Scan braucht NETZ, und das ist hier der Zweck,
# nicht ein Zugestaendnis (ADR-0037). `gates` bleibt hermetisch.
image-scan: ## CVE-Scan gegen das PUBLIZIERTE Image (Netz, NICHT in gates, Trivy digest-gepinnt; ADR-0037). Skript-Exit 1 = behebbare CRITICAL/HIGH, 2 = gescheitert — ueber make nicht unterscheidbar.
	@bash tools/image-scan.sh

# KEIN Gate — ein WERKZEUG. Es prueft nichts, es bewegt: `git mv` plus den Nachzug
# der Verweise AUF die bewegte Datei (BEO-008, slice-118). Die Gegenrichtung —
# Verweise IN wandernden Dateien — traegt das Modul `links` (doc-check).
slice-mv: ## Lifecycle-Wechsel eines Slice samt der Verweise auf ihn (AGENTS §5). SLICE=<slice-NNN> TO=<open|next|in-progress|done>
	@bash tools/slice-mv.sh "$(SLICE)" "$(TO)"

# Die Teil-Sensoren laufen als Sequenz im selben Rezept, NICHT als
# Prerequisites: make bricht sonst beim ersten roten Target ab, und wer zwei
# Verstoesse in verschiedenen Bereichen hat, sieht nur den ersten (Review
# 2026-07-26, R-052-F4). Eine Verifikations-Schicht beantwortet mehrere
# unabhaengige Fragen — sie soll ALLE beantworten, auch wenn eine mit Nein
# ausgeht. Die Einzel-Targets bleiben fuer den gezielten Aufruf bestehen.
#
# OHNE ZAHLENANGABE, bewusst (slice-077, Fund F-15): hier stand einmal eine
# Zahl, waehrend das Rezept eine andere Menge ausfuehrte — ein hinzugekommener
# Sensor zog sie nicht nach. Die Zahl war nie die Aussage; die Aussage ist
# "alle, nicht nur die erste". Eine Angabe, die bei jeder Erweiterung
# mitgepflegt werden muss und keinen Inhalt traegt, ist eine Drift-Quelle
# ohne Gegenwert.
#
# Eine der Fragen ist seit slice-080 ein MODUL-Target statt eines Skripts
# (doc-structure). Es haengt aus demselben Grund hier und nicht an gates:
# es beantwortet Form- und Closure-Fragen, keine Code-Fragen.
verify: ## Verifikations-Schicht: DoD-/Closure-Fragen (vor der "fertig"-Meldung; getrennt von gates).
	@fail=0; \
	bash tools/verify-risiko-ausgaenge.sh || fail=1; \
	bash tools/verify-observations.sh  || fail=1; \
	$(MAKE) --no-print-directory doc-structure || fail=1; \
	$(MAKE) --no-print-directory doc-complete  || fail=1; \
	if [ "$$fail" -ne 0 ]; then \
	  echo "[verify] FAIL — mindestens eine Verifikations-Frage ist offen; alle Befunde stehen oben." >&2; \
	  exit 1; \
	fi; \
	echo "[verify] Verifikations-Schicht gruen"

record-gates: ## Gate-Nachweis (Working-Tree-Hash) für den Stop-Hook schreiben.
	@bash tools/harness/record-gates.sh

gates: lint test coverage-gate arch-check doc-check doc-targets doc-planning doc-workflows doc-reviews gate-consistency version-coherence suppression-check guard-selftest ci-range-selftest record-gates ## alle inneren Gates (mandatory vor Handoff).

image-test: build ## AC-FA-DIST-001 + nativ==Container-Akzeptanz gegen das gebaute Image.
	@IMAGE=$(IMAGE) bash tools/image-test.sh

ci: gates image-test ## CI-äquivalenter Lauf: gates + image-test (AC-FA-DIST-001).
	@echo "[ci] gates + image-test grün"

# FOCUS_COMMITS wählt alle übrigen .d-check.yml-Module ab, sodass nur `commits`
# läuft (sonst feuern die Datei-Module auf den Arbeitsbaum). ADR-0021.
DCHECK_RUN_I := $(DOCKER) run --rm -i --network none -v "$(CURDIR):/repo:ro" $(DCHECK_REF)
FOCUS_COMMITS := --enable commits --disable links --disable anchors --disable ids --disable matrix --disable external --disable codepaths --disable spans --disable hostpaths --disable diagrams --disable versions --disable pins --disable immutable --disable vcs --disable planning --disable tracked

trace-check: ## Traceability via Modul commits (ADR-0021): AC-/ADR-/MR-/slice-ID je Commit. MSGFILE=<datei> (Hook), RANGE=a..b (CI), sonst HEAD~1..HEAD. AGENTS §5.
	@$(if $(MSGFILE),$(DCHECK_RUN_I) --commit-msg - < $(MSGFILE),$(DOCKER) run --rm --network none -v "$(CURDIR):/repo:ro" $(DCHECK_REF) $(FOCUS_COMMITS) --range $(if $(RANGE),$(RANGE),HEAD~1..HEAD))

hooks: ## git-Hooks installieren (core.hooksPath -> .githooks; commit-msg Traceability). AGENTS §5.
	@git config core.hooksPath .githooks
	@echo "[hooks] core.hooksPath=.githooks — commit-msg Traceability-Gate aktiv"

# archive-wave (Baseline-Regelwerk modul-06-roadmap.md §Wellen-Closure-
# Prozedur, Schritt 4): eigenstaendiges Werkzeug unter tools/archive-wave/,
# unveraendert aus d-check uebernommen (dasselbe Planning-Layout, portabel
# per README-Zusage, kein Import aus einem d-check-internen Paket). Eigenes
# go.mod, darum eigenes Test-Target statt Teil von `make test`.
archive-wave-test: ## archive-wave-Testsuite (eigenes go.mod, nicht Teil von `make test`).
	$(MAKE) -C tools/archive-wave test GO_VERSION=$(GO_VERSION) PROGRESS_FLAG='$(PROGRESS_FLAG)'

archive-wave: ## Welle oder wellenlosen Slice archivieren: make archive-wave WELLE=welle-NN|SLICE=slice-NNN|REVIEW=datei.md [APPLY=1].
	$(MAKE) -C tools/archive-wave run WELLE=$(WELLE) SLICE=$(SLICE) REVIEW=$(REVIEW) APPLY=$(APPLY) ROOT=$(CURDIR) \
	    GO_VERSION=$(GO_VERSION) PROGRESS_FLAG='$(PROGRESS_FLAG)'
