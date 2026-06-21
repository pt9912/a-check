# ADR-0005 — Lint-Profil: golangci-lint v2, SOLID-nahe Linter ohne `//nolint`

- **Status:** Proposed
- **Datum:** 2026-06-21
- **Bezug:** [`AGENTS.md`](../../../AGENTS.md) §3.2 (Suppression-Verbot), §3.6 (Prüfregeln/Lockerungen = ADR), §4 (`make lint`); [ADR-0003](0003-config-modell-a-check-yml.md) (Dependency-Fläche YAML)
- **Schärft:** — (Tooling-/Prozess-Entscheidung; koppelt an keine Spec-§. Sie *realisiert* das `make lint`-Gate, das die Spec-Straten nicht selbst beschreiben.)
- **Supersedes:** —

## Kontext

[`AGENTS.md`](../../../AGENTS.md) §4 fordert `make lint` (golangci-lint mit
Projekt-Profil), §3.2 verbietet Inline-Suppressions (`//nolint`), §3.6 verlangt
für jede Prüfregel und jede Gate-Lockerung einen ADR. Das Lint-Gate braucht
also ein **deklariertes** Profil, dessen Ausnahmen begründet und zentral sind.

## Optionen

1. **Minimal-Profil** (nur die fünf Default-Linter). Trade-off: schnell grün,
   fängt aber keine Komplexitäts-, Kontext-, Interface- oder Modul-Smells —
   genau die SOLID-nahen Klassen, die Reviews sonst manuell tragen müssten.
2. **SOLID-nahes Profil** (Default-5 + Komplexitäts-/Kontext-/Interface-/
   Modul-Linter), Ausnahmen **zentral** in `.golangci.yml` mit `Why:` statt
   `//nolint`. Trade-off: strenger, erzwingt disziplinierten Code (keine
   Paket-Globals, Komplexitätsgrenzen) — deckt sich mit dem Stack.
3. **golangci-lint mit `//nolint` ad hoc.** Verworfen — verstößt direkt gegen
   [`AGENTS.md`](../../../AGENTS.md) §3.2.

## Entscheidung

**Option 2.** golangci-lint **v2** (versions-gepinnt `v2.12.2`), Profil in
[`.golangci.yml`](../../../.golangci.yml). Inline-Suppression ist verboten;
Ausnahmen leben zentral unter `exclusions` mit `Why:`-Kommentar (§3.6).
Aktivierte Zusatz-Linter u. a. `cyclop`/`gocognit`/`gocyclo`/`funlen`
(Komplexität), `gochecknoglobals`/`gochecknoinits` (Zustand), `ireturn`
(nur Kern-Ports erlaubt), `forbidigo` (kein direktes `fmt.Print*`),
`gomodguard_v2` (Dependency-Fläche bei `yaml.v3`, [ADR-0003](0003-config-modell-a-check-yml.md)),
`revive`-Regelset, `testpackage`. Deklarierte Ausnahmen: Komplexitäts-Linter
für `_test.go`; `testpackage` für die White-Box-Unit-Tests unter `internal/`
(die Black-Box-Vertragsabdeckung liefern CLI-Akzeptanztests in einem
Folge-Increment).

## Konsequenzen

- **Fitness Function / Gate:** `make lint` (golangci-lint `v2.12.2` in Docker)
  ist das Maintainability-Gate ([`AGENTS.md`](../../../AGENTS.md) §4); Teil von
  `make gates`.
- **Code-Disziplin:** keine Paket-Globals (Regexes leben in den Adaptern),
  Komplexitätsgrenzen erzwingen kleine Funktionen, `ireturn` lässt nur die
  Kern-Port-Schnittstellen ([ARC-002](../../../spec/architecture.md)) als
  Interface-Rückgabe zu.
- **Ehrlichkeit:** jede Ausnahme ist ADR-/`Why:`-dokumentiert (§3.6), nicht
  still — kein `//nolint` im Code.
- Die golangci-lint-Version wird wie die übrigen Images bewusst gehoben
  (Pin-Politik, analog [ADR-0004](0004-distribution-image-mk.md)).
