# Makefile — a-check
#
# Bootstrap-Stand: nur das Doku-Gate `doc-check`, eingebunden aus d-check.mk
# (Schwester-Tool d-check als Doku-Sensor, Dogfooding des Stacks). Die übrigen
# Gates (lint/test/arch-check und der Aggregat `gates`) entstehen mit slice-003,
# sobald a-check selbst gebaut wird.

include d-check.mk

.DEFAULT_GOAL := help

.PHONY: help
help: ## Diese Hilfe anzeigen.
	@grep -hE '^[a-zA-Z_-]+:.*?## ' $(MAKEFILE_LIST) | sort | awk 'BEGIN{FS=":.*?## "}{printf "  \033[36m%-12s\033[0m %s\n",$$1,$$2}'
