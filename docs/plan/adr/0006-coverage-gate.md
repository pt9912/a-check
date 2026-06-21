# ADR-0006 — Coverage-Gate: Gesamt-Schwelle 90 % (Kalibrierungs-Bindung)

- **Status:** Accepted
- **Datum:** 2026-06-21
- **Bezug:** [`AGENTS.md`](../../../AGENTS.md) §3.6 (Schwellen = ADR), §4 (`make gates`); [ADR-0005](0005-lint-profil.md) (Schwester-Gate Maintainability)
- **Schärft:** — (Gate-/Tooling-Entscheidung; koppelt an keine Spec-§. Sie *realisiert* das Coverage-Gate, das die Spec-Straten nicht selbst beschreiben.)
- **Supersedes:** —

## Kontext

[`AGENTS.md`](../../../AGENTS.md) §3.6 verlangt für jede Schwelle (Coverage,
Linter-Strenge, Prüfregel) einen ADR. Ein Coverage-Gate braucht also eine
**deklarierte** Schwelle samt Reifegrad. Regelwerk Modul 13 stellt die Wahl
zwischen *bootstrap-aware* (weiche Frühphase, Hochschalt-Trigger) und einer
festen Schwelle.

## Optionen

1. **Keine Coverage-Schwelle.** Verworfen — Regressionen in der Testdichte
   blieben unbemerkt; ein Gate, das nichts erzwingt, ist ein Vorschlag.
2. **Bootstrap-aware Ramp** (z. B. 80 % → 90 % bei einem Meilenstein).
   Sinnvoll, *wenn* der Bestand die Zielschwelle noch nicht trägt.
3. **Feste 90 %-Schwelle ab jetzt.** Der Bestand trägt sie bereits
   (Ist **92,60 %** am 2026-06-21 über `./internal/...`), eine weiche
   Frühphase wäre unehrlich.

## Entscheidung

**Option 3.** Feste Gesamt-Schwelle **90 %** über `./internal/...`
(**Kalibrierungs-Bindung**; Ist 92,60 % am 2026-06-21). Mechanik wie im
Stack (d-check/u-boot): eine Dockerfile-`coverage`-Stage misst mit
`go test -coverpkg=$(go list ./internal/...) -covermode=atomic` und ruft
[`tools/coverage-gate.sh`](../../../tools/coverage-gate.sh) mit der über die
`THRESHOLD`-Make-Variable gereichten Schwelle. `-coverpkg` gibt den
Integrationstests (CLI-Akzeptanztests) Cross-Package-Gutschrift.

- Das untestbare `cmd/`-`os.Exit`-Wrapper ist aus der Messung ausgenommen;
  die CLI-Logik liegt testbar in `internal/cli` (Black-Box-`cli_test`) — das
  erfüllt zugleich die in [ADR-0005](0005-lint-profil.md) angekündigte
  Black-Box-Vertragsabdeckung.
- **Override** nur via `make coverage-gate THRESHOLD=…`; eine **Senkung** ist
  ein neuer ADR ([`AGENTS.md`](../../../AGENTS.md) §3.6), kein PR-Kommentar.
- Schwelle und Historie leben in [`harness/README.md`](../../../harness/README.md)
  §Sensors; eine Verfehlung ⇒ Carveout-Pflicht.

## Konsequenzen

- **Fitness Function / Gate:** `make coverage-gate` (Dockerfile-`coverage`-Stage,
  Teil von `make gates`).
- Neue Code-Pfade brauchen Tests, sonst fällt das Gate unter 90 %.
- Die Schwelle ist eine bewegliche **Kalibrierungs-Bindung** (Sensors-Tabelle),
  keine stille Konstante — Anhebung jederzeit, Senkung nur per ADR.
