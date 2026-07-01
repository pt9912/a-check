# d-check.mk — Doku-Gate via d-check (Schwester-Tool), zum `include` in das
# Makefile des konsumierenden Repos. Stack-konforme Integration ohne
# Skript-Kopie: nur dieses Fragment + die repo-eigene .d-check.yml.
#
# Erzeugt aus `d-check --print-mk` (DC-FA-CLI-010) und an die Repo-Politik
# angepasst: Image digest-gepinnt statt Tag (Pin-Politik sinngemäß AC-QA-03),
# `## `-Help-Annotation je Target (make help), und bewusst nur die von a-check
# genutzten Targets (doc-check/-trace/-complete). v0.35.0 bietet zusätzlich
# doc-doctor/-repair/-immutable/-commits/-help — hier weggelassen, weil jedes
# reale d-check.mk-Target sonst in AGENTS §4 stehen müsste (gate-consistency).
# Refresh: `docker run --rm ghcr.io/pt9912/d-check:vX --print-mk`, dann Digest neu pinnen.

# d-check-Image, digest-gepinnt (v0.35.0) — Reproduzierbarkeit (Pin-Politik
# sinngemäß AC-QA-03). `?=` erlaubt Override; für strikte Reproduzierbarkeit
# bleibt der Digest gepinnt.
DCHECK_IMAGE ?= ghcr.io/pt9912/d-check@sha256:9d7b23ac82b94a97bc98e36b748e48644dff7adc6702cb608f13402d309e6558

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
