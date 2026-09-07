# `make ci-range-selftest` — Selbsttest der Commit-Range-Weiche der CI

## Vertrag

Vier Fälle gegen `tools/ci-commit-range.sh`: `pull_request`, neuer Branch,
**Force-Push** und normaler Push. Der dritte war der Defekt, den der Sensor
festhält: `github.event.before` trägt nach einem Rebase einen gültig
aussehenden SHA, den der Runner-Klon nicht kennt (`fetch-depth: 0` holt keine
verwaisten Objekte) — die Range-Prüfungen brachen mit *„Range-Basis nicht
auflösbar"* ab. Der vierte Fall ist die Gegenrichtung: Eine **brauchbare** Basis
wird auch benutzt, die Weiche fällt nicht pauschal auf den Default-Branch.

## Grenze — was das Grün nicht abdeckt

1. **Welchen Wert GitHub in `github.event.before` schreibt** — das sagt nur ein
   echter Lauf. Der Selbsttest prüft, was die Weiche mit einem gegebenen Wert
   tut, nicht ob der Wert kommt. Permanent.

## Bindung

CI-Schicht · slice-134 · im `gates`-Aggregat.
