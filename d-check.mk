# d-check.mk — Doku-Gate via d-check (Schwester-Tool), zum `include` in das
# Makefile des konsumierenden Repos. Stack-konforme Integration ohne
# Skript-Kopie: nur dieses Fragment + die repo-eigene .d-check.yml.
#
# Erzeugt aus `d-check --print-mk` (DC-FA-CLI-010, v0.24.0) und an die
# Repo-Politik angepasst: Image digest-gepinnt statt Tag (Pin-Politik
# sinngemäß AC-QA-03) sowie `## `-Help-Annotation je Target (make help).
# Refresh: `docker run --rm -v "$PWD:/repo:ro" ghcr.io/pt9912/d-check:vX --print-mk`,
# dann Digest neu pinnen.

# d-check-Image, digest-gepinnt (v0.24.0) — Reproduzierbarkeit (Pin-Politik
# sinngemäß AC-QA-03). `?=` erlaubt Override; für strikte Reproduzierbarkeit
# bleibt der Digest gepinnt.
DCHECK_IMAGE ?= ghcr.io/pt9912/d-check@sha256:1c28a2b7e0e624763577ecba75b027f384692ecaa8a78a6e353a1a0c1889a4f8

# TRACE_FLAGS: optionale Flags für die advisory RTM-Targets (z. B. --json).
TRACE_FLAGS ?=

.PHONY: doc-check
doc-check: ## Doku: Links/Anker/Kennungs-Linkpflicht/Referenzmatrix via d-check (netzlos, read-only).
	docker run --rm --network none -v "$(CURDIR)":/repo:ro $(DCHECK_IMAGE)

# doc-trace: advisory Requirements Traceability Matrix (DC-FA-CLI-009).
.PHONY: doc-trace
doc-trace: ## Doku: advisory Requirements Traceability Matrix via d-check (DC-FA-CLI-009).
	docker run --rm --network none -v "$(CURDIR)":/repo:ro $(DCHECK_IMAGE) --trace $(TRACE_FLAGS)

# doc-complete: Vollständigkeits-Gate — Requirements-Waise ⇒ Exit 1 (DC-FA-CLI-011).
.PHONY: doc-complete
doc-complete: ## Doku: Vollständigkeits-Gate, Requirements-Waise ⇒ Exit 1 (DC-FA-CLI-011).
	docker run --rm --network none -v "$(CURDIR)":/repo:ro $(DCHECK_IMAGE) --trace --require-complete $(TRACE_FLAGS)
