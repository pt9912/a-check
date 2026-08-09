# a-check.mk — Architektur-Gate via a-check, zum `include` in das
# Makefile des konsumierenden Repos. Erzeugt von `a-check --print-mk`.
#
# A_CHECK_IMAGE wird beim Release auf `@sha256:…` digest-gepinnt.
A_CHECK_IMAGE ?= ghcr.io/pt9912/a-check@sha256:aef28cfe25bb054b1b0eb28420222a45b9f6ce9425b7ffd0f55e6ae56f295b56

# Container-Runtime ueber eine Indirektion, damit ein Repo mit podman/nerdctl
# oder einem docker-Wrapper nicht die Haelfte seiner Targets anders faehrt als
# die andere (slice-082).
#
# REIHENFOLGE ZAEHLT: `?=` setzt nur, wenn DOCKER noch nicht belegt ist.
# Wer eine eigene Runtime nutzt, definiert sie VOR dem `include` — oder
# hart (`DOCKER = podman`). Ein `DOCKER ?= podman` NACH dem
# include greift nicht mehr, weil dieses Fragment die Variable dann schon
# gesetzt hat.
DOCKER ?= docker

.PHONY: a-check a-check-graph
a-check: ## Architektur: Hexagon-Regeln via a-check (netzlos, read-only).
	$(DOCKER) run --rm --network none -v "$(CURDIR)":/src:ro $(A_CHECK_IMAGE) /src

a-check-graph: ## Architektur-Graph (Mermaid) aus .a-check.yml auf stdout (read-only, kein Scan).
	$(DOCKER) run --rm --network none -v "$(CURDIR)":/src:ro $(A_CHECK_IMAGE) --print-graph /src
