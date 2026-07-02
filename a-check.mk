# a-check.mk — Architektur-Gate via a-check, zum `include` in das
# Makefile des konsumierenden Repos. Erzeugt von `a-check --print-mk`.
#
# A_CHECK_IMAGE ist auf den v0.6.0-Release digest-gepinnt (AC-QA-03, ADR-0007);
# Pin-Hebung ist ein bewusster Commit (ADR-0004).
A_CHECK_IMAGE ?= ghcr.io/pt9912/a-check@sha256:b349a15087b61ba9c3a90bcd16049dddff4efcbcd746768fb6c7e20fb8fe9afd

.PHONY: a-check
a-check: ## Architektur: Hexagon-Regeln via a-check (netzlos, read-only).
	docker run --rm --network none -v "$(CURDIR)":/src:ro $(A_CHECK_IMAGE) /src
