# a-check

Sprachübergreifender Hexagon-Architektur-Checker — deterministisch,
seiteneffektfrei, text-heuristisch, ausgeliefert als Container-Image.

**Status: in Entwicklung (0.1.0).** Lastenheft, Spezifikation, Architektur
und die Go-Implementierung stehen; alle inneren Gates sind grün. Ein
getaggtes GHCR-Release folgt. Verbindlich ist das
[Lastenheft](spec/lastenheft.md); die Versionshistorie führt die
[CHANGELOG.md](CHANGELOG.md).

## Was ist a-check?

**a-check** erzwingt die hexagonale Schicht-Architektur eines Repositories
**sprachübergreifend**, gesteuert über eine Konfigurationsdatei. Fünf
universelle Regeln, je eine Anforderung im [Lastenheft](spec/lastenheft.md):

- `core-impurity` — der Kern importiert weder Adapter noch Framework/Tech
  ([AC-FA-RULE-001](spec/lastenheft.md#ac-fa-rule-001--kern-reinheit-regel-core-impurity))
- `lateral-adapter` — ein Adapter importiert keinen anderen Adapter (außer der gemeinsamen Senke)
  ([AC-FA-RULE-002](spec/lastenheft.md#ac-fa-rule-002--keine-lateralen-adapter-kanten-regel-lateral-adapter))
- `tech-leak` — ein Framework/Tech erscheint nur in seinem Adapter
  ([AC-FA-RULE-003](spec/lastenheft.md#ac-fa-rule-003--tech-kapselung-regel-tech-leak))
- `port-impurity` — Ports sind reine Abstraktionen
  ([AC-FA-RULE-004](spec/lastenheft.md#ac-fa-rule-004--port-disziplin-regel-port-impurity))
- `wrong-direction` — Schicht-Kanten sind einbahnig
  ([AC-FA-RULE-005](spec/lastenheft.md#ac-fa-rule-005--schicht-richtung-regel-wrong-direction))

Die Imports werden **text-heuristisch** je Sprache (C++/Go/Rust/Kotlin)
extrahiert ([AC-FA-EXTRACT-001](spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion)).
Jeder Befund nennt Datei, Zeile, Regel und Grund; Exit-Codes: `0` sauber,
`1` Befunde, `2` Nutzungs-/Konfigurationsfehler
([AC-FA-CLI-001](spec/lastenheft.md#ac-fa-cli-001--aufruf-scan-wurzel-und-exit-codes)).

## Warum a-check?

Vier funktional überlappende `arch-check.sh`-Varianten sind in den
Schwester-Repositories gewachsen — C++ über `#include`-Heuristik (`b-cad`),
Go über `go list` (`d-check`), Rust über `use`-Heuristik (`grid-guide`),
Kotlin über Gradle-Modulgrenzen (`d-migrate`): vier Sprachen, vier
Mechanismen, dieselben fünf Regeln. a-check ersetzt sie durch **ein** Tool:

- **Konfiguration statt Fork:** repo-spezifische Schicht-/Tech-Regeln leben
  deklarativ in `.a-check.yml`
  ([AC-FA-CONF-001](spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml)),
  nicht in kopierten Skripten.
- **Ein Distributionsweg:** digest-gepinntes Container-Image plus
  mitgeliefertes `a-check.mk` statt n gepflegter Kopien
  ([AC-FA-DIST-001](spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk)).

Es ist das **Architektur-Gegenstück zu `d-check`** (Doku-Referenzen):
dieselbe Gründungslogik (eine Familie driftender Skripte durch ein Werkzeug
ersetzen), eine Abstraktionsebene höher.

## Kerngedanke

**Architektur ist ein Import-Graph mit prüfbaren Invarianten.** Ob der Kern
rein bleibt, ein Adapter lateral importiert oder eine Schicht-Kante gegen
die Richtung läuft, ist maschinell entscheidbar — a-check macht diese
Invarianten zum Gate statt zur Review-Meinung.

Dabei gilt **berichten, nie reparieren**: a-check ist ein reines Lese-Tool.
Die **Heuristik-Grenze** (text-basiert, kein vollständiger Parser je Sprache)
wird ausgewiesen, nicht verschwiegen; eine Allowlist/Marker-Ausnahme ist
konfigurierbar
([AC-QA-02](spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)).

## Was macht es vertrauenswürdig?

- **Determinismus:** identische Eingabe ⇒ byte-identische, stabil sortierte
  Ausgabe ([AC-QA-01](spec/lastenheft.md#ac-qa-01--determinismus)).
- **Hermetisch & netzlos:** schreibt nie ins geprüfte Repo, läuft mit
  `--network none` auf distroless/static
  ([AC-QA-02](spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)).
- **Keine stillen Defaults:** jede ungültige `.a-check.yml` bricht mit Exit 2
  ab (strict decode,
  [AC-FA-CONF-001](spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml)).
- **Reproduzierbar:** Image und `a-check.mk` referenzieren einen
  `@sha256:`-Digest ([AC-QA-03](spec/lastenheft.md#ac-qa-03--reproduzierbarkeit)).
- **Dogfooding:** a-check prüft seine eigene Hexagon-Architektur bei jedem
  `make arch-check` — mit der [Selbstkonfiguration](.a-check.yml), 0 Befunde.

## Nutzung

Lokal (Dogfooding, läuft heute):

```bash
make build        # static/distroless Image bauen
make arch-check   # a-check prüft sich selbst (netzlos, read-only)
```

Konsumenten binden a-check als `make a-check`-Gate ein — **ohne
Skript-Kopie**: das mitgelieferte [`a-check.mk`](a-check.mk) (von
`a-check --print-mk` erzeugt) wird `include`-t, dazu ein `.a-check.yml`.
Ab dem GHCR-Release wird `A_CHECK_IMAGE` per `@sha256:`-Digest gepinnt
([AC-FA-DIST-001](spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk),
[AC-QA-03](spec/lastenheft.md#ac-qa-03--reproduzierbarkeit)).

## Konfiguration (`.a-check.yml`)

In der Repo-Wurzel; strikt dekodiert (unbekannter Schlüssel ⇒ Exit 2):

```yaml
version: 1
languages:
  go: ["**/*.go"]
layers:
  core:     ["internal/core/**"]
  adapters: ["internal/adapters/**"]
edges:
  - {from: adapters, to: core}
```

Das vollständige Schema steht in der
[Spezifikation §SPEC-CONF-001](spec/spezifikation.md#spec-conf-001--konfigurationsschema);
ein lebendes Beispiel ist die [Selbstkonfiguration dieses Repos](.a-check.yml).
`a-check --print-config` gibt ein kommentiertes Gerüst aus.

## Einstieg

| Dokument | Inhalt |
|---|---|
| [`spec/lastenheft.md`](spec/lastenheft.md) | Anforderungen (`AC-FA-*`, `AC-QA-*`), Akzeptanzkriterien |
| [`spec/spezifikation.md`](spec/spezifikation.md) | Algorithmen, `.a-check.yml`-Schema, Exit-Codes (`SPEC-*`) |
| [`spec/architecture.md`](spec/architecture.md) | Hexagon-Komponenten und Rollen (`ARC-*`) |
| [`docs/plan/adr/README.md`](docs/plan/adr/README.md) | Architekturentscheidungen (ADR-Index) |
| [`harness/README.md`](harness/README.md) | Harness-Einstieg: Source Precedence, Guides, Sensors |
| [`AGENTS.md`](AGENTS.md) | Briefing für AI-Coding-Agenten, Hard Rules |
| [`CHANGELOG.md`](CHANGELOG.md) | Änderungshistorie |

## Entwicklung

Der Host braucht nur `git`, GNU `make`, `bash` und Docker
([`AGENTS.md`](AGENTS.md) §3.1).

```bash
make help     # verfügbare Targets
make gates    # alle inneren Gates (lint/test/coverage-gate/arch-check/doc-check)
```

## Lizenz

Dieses Projekt steht unter der [MIT-Lizenz](LICENSE).
