# a-check.mk — Architektur-Gate via a-check, zum `include` in das
# Makefile des konsumierenden Repos. Erzeugt von `a-check --print-mk`.
#
# A_CHECK_IMAGE ist auf den v0.5.0-Release digest-gepinnt (AC-QA-03, ADR-0007);
# Pin-Hebung ist ein bewusster Commit (ADR-0004).
A_CHECK_IMAGE ?= ghcr.io/pt9912/a-check@sha256:81951e61babc2ead2c74d1155ef98470ea2dcb21d25eae3a56fe62cc6534ed64

.PHONY: a-check
a-check: ## Architektur: Hexagon-Regeln via a-check (netzlos, read-only).
	docker run --rm --network none -v "$(CURDIR)":/src:ro $(A_CHECK_IMAGE) /src
