# a-check.mk — Architektur-Gate via a-check, zum `include` in das
# Makefile des konsumierenden Repos. Erzeugt von `a-check --print-mk`.
#
# A_CHECK_IMAGE ist auf den v0.11.0-Release digest-gepinnt;
# Pin-Hebung ist ein bewusster Commit.
A_CHECK_IMAGE ?= ghcr.io/pt9912/a-check@sha256:f53a06fe02973a0cd4771eccb3d22171619a5ddfcbca2fac9922ce3fe7520483

.PHONY: a-check
a-check: ## Architektur: Hexagon-Regeln via a-check (netzlos, read-only).
	docker run --rm --network none -v "$(CURDIR)":/src:ro $(A_CHECK_IMAGE) /src
