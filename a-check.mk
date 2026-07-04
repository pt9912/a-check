# a-check.mk — Architektur-Gate via a-check, zum `include` in das
# Makefile des konsumierenden Repos. Erzeugt von `a-check --print-mk`.
#
# A_CHECK_IMAGE ist auf den v0.10.0-Release digest-gepinnt;
# Pin-Hebung ist ein bewusster Commit.
A_CHECK_IMAGE ?= ghcr.io/pt9912/a-check@sha256:0932cb1dbdfa6ece0a5f9892dbe541cf29618ffe2667feda01c96d3218af2fc9

.PHONY: a-check
a-check: ## Architektur: Hexagon-Regeln via a-check (netzlos, read-only).
	docker run --rm --network none -v "$(CURDIR)":/src:ro $(A_CHECK_IMAGE) /src
