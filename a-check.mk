# a-check.mk — Architektur-Gate via a-check, zum `include` in das
# Makefile des konsumierenden Repos. Erzeugt von `a-check --print-mk`.
#
# A_CHECK_IMAGE ist auf den v0.2.0-Release digest-gepinnt (AC-QA-03, ADR-0007);
# Pin-Hebung ist ein bewusster Commit (ADR-0004).
A_CHECK_IMAGE ?= ghcr.io/pt9912/a-check@sha256:4132a7af33eb11fdb3738e1ce1cd5f95b33455a8b4a84f07fcb4c4db1f2e4989

.PHONY: a-check
a-check: ## Architektur: Hexagon-Regeln via a-check (netzlos, read-only).
	docker run --rm --network none -v "$(CURDIR)":/src:ro $(A_CHECK_IMAGE) /src
