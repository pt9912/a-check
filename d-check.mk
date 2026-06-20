# d-check.mk — Doku-Gate via d-check (Schwester-Tool), zum `include` in das
# Makefile des konsumierenden Repos. Stack-konforme Integration ohne
# Skript-Kopie: nur dieses Fragment + die repo-eigene .d-check.yml.
#
# Interim: bis d-check selbst ein `--print-mk` bereitstellt (analog zu a-checks
# geplantem a-check.mk / AC-FA-DIST-001), wird dieses Fragment von Hand
# gepflegt. Zielbild danach: `docker run ... d-check --print-mk > d-check.mk`.

# d-check-Image, digest-gepinnt (v0.19.0) — Reproduzierbarkeit (Pin-Politik
# sinngemäß AC-QA-03, auf das Werkzeug angewandt). `?=` erlaubt Override.
DCHECK_IMAGE ?= ghcr.io/pt9912/d-check@sha256:6134b8bd963de188858357ba05861a849dfb79dfac774437818f976100909ceb

.PHONY: doc-check
doc-check: ## Doku: Links/Anker/Kennungs-Linkpflicht/Referenzmatrix via d-check (netzlos, read-only).
	docker run --rm --network none -v "$(CURDIR)":/repo:ro $(DCHECK_IMAGE)
