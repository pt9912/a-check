# a-check.mk — Architektur-Gate via a-check, zum `include` in das
# Makefile des konsumierenden Repos. Erzeugt von `a-check --print-mk`.
#
# A_CHECK_IMAGE ist auf den v0.11.0-Release digest-gepinnt;
# Pin-Hebung ist ein bewusster Commit.
A_CHECK_IMAGE ?= ghcr.io/pt9912/a-check@sha256:24939a6b47bfbfcb9845813cd781a70fba161fb12c6593e188ac51ede5267a9c

.PHONY: a-check
a-check: ## Architektur: Hexagon-Regeln via a-check (netzlos, read-only).
	docker run --rm --network none -v "$(CURDIR)":/src:ro $(A_CHECK_IMAGE) /src
