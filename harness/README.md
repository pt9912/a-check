# Harness

## Purpose

Dieser Harness verbindet Spezifikationen, ADRs, Planning-Dokumente und
Gates dieses Repos. Er ist **kein Ersatz** für `spec/` oder `docs/`,
sondern ein **Einstiegspunkt** für Menschen und AI-Code-Agenten.

Wenn diese Datei einer kanonischen Quelle widerspricht, **gewinnt die
kanonische Quelle**, und diese Datei wird angepasst.

Strukturregeln (Verzeichniskonvention, ID-Schemata, Modus-Deklarationen
pro Sub-Area, Zusatzklassen für Sensors-Bindung) sowie Adaptionen ggü.
der adoptierten Baseline leben in [`conventions.md`](conventions.md).
Diese Datei dupliziert sie nicht.

## Source precedence

| Rang | Datei | Charakter |
|---|---|---|
| 1 | [`spec/lastenheft.md`](../spec/lastenheft.md) | vertraglich abnahmebindend |
| 2 | `spec/spezifikation.md` | technisch fortschreibbar (geplant, entsteht mit slice-002) |
| 3 | `spec/architecture.md` | Komponenten/Sequenzen, meilensteinfrei (geplant, entsteht mit slice-002) |
| 4 | [`docs/plan/adr/`](../docs/plan/adr/) | Architekturentscheidungen |
| 5 | `docs/plan/planning/in-progress/roadmap.md` | aktuelle Welle (geplant, entsteht mit slice-001) |
| 6 | `README.md` | Projekt-Überblick (geplant) |
| 7 | [`AGENTS.md`](../AGENTS.md) | Agent-Briefing |
| 8 | diese Datei | Harness-Einstieg |

Acht Ränge ohne `docs/user`-Stratum (Pre-Release-Tool ohne Betriebs-Doku)
— deklariert als [`MR-003`](conventions.md#mr-003--source-precedence-ohne-docsuser-rang).

## Guides (Feedforward-Quellen)

| Quelle | Inhalt |
|---|---|
| [`spec/lastenheft.md`](../spec/lastenheft.md) | Anforderungen (`AC-FA-*`, `AC-QA-*`), Akzeptanzkriterien |
| `spec/spezifikation.md` | Algorithmen, Schemas (`.a-check.yml`, `--json`), Defaults, Exit-Codes (geplant) |
| `spec/architecture.md` | Hexagon-Schnitt (Rollen), Zugriffs-Constraints, Sequenzen (geplant) |
| [`docs/plan/adr/`](../docs/plan/adr/) | Architekturentscheidungen |
| [`docs/plan/planning/`](../docs/plan/planning/) | Slice-Pläne und Roadmap |
| [`AGENTS.md`](../AGENTS.md) | Hard Rules, Source Precedence, Workflow |
| [`conventions.md`](conventions.md) | repo-lokale Strukturregeln, Adaptions-Block (`MR-*`), Modus-Deklarationen |
| [`agents-regelwerk.md`](https://raw.githubusercontent.com/pt9912/ai-harness-course/v1.3.0/kurs/de/agents-regelwerk.md) | adoptiertes Betriebsregelwerk der Baseline in Agenten-Kurzform; derivativ — Stand siehe [`conventions.md` §Baseline](conventions.md#baseline) |

## Sensors (Feedback-Gates)

Nur Targets, die im Makefile **existieren**, dürfen hier als real gelten.
**Stand Increment 1: das Makefile existiert noch nicht** — die folgende
Tabelle ist vollständig **geplant** (entsteht mit slice-003) und bindet
die jeweils genannten Anforderungen. Bis dahin wird kein Gate als
ausgeführt behauptet.

| Target | Vertrag | Bindung | Stand |
|---|---|---|---|
| `make lint` | golangci-lint mit Projekt-Profil; Inline-Suppressions verboten | Lint-Profil-ADR (geplant) | geplant (slice-003) |
| `make test` | Akzeptanzkriterien der bezogenen `AC-FA-*` als Tests; Determinismus-Test | [`AC-QA-01`](../spec/lastenheft.md#ac-qa-01--determinismus) (AC-Bindung) | geplant (slice-003) |
| `make arch-check` | Eigen-Architektur via `a-check` selbst (Dogfooding) | [`AC-QA-02`](../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze) (AC-Bindung) | geplant |
| `make doc-check` | Links/Anker/Kennungen der Repo-Doku via `d-check` (stack-konform, `--network none`) | (AC-Bindung mit slice-003) | geplant |
| `make gates` | aggregiert die inneren Gates; `record-gates` als letzter Schritt | — | geplant (slice-003) |

**Aktueller Lauf-Status:** keiner — Bootstrap (Increment 1). Es existiert
noch kein Makefile, kein Gate-Lauf.
**Rote Gates:** keine (keine existieren).
**Nicht behauptet:** alle obigen Targets (geplant, noch nicht angelegt).

## Traceability rules

- PRs/Commits **müssen** mindestens eine `AC-*`- oder `ADR-*`-ID nennen.
- Neue oder geänderte Anforderungen brauchen einen Beleg: Test, Gate, Demo oder ADR.
- Neue ADRs müssen im [ADR-Index](../docs/plan/adr/README.md) ergänzt werden.
- Änderungen an Planning-Dokumenten folgen den Lifecycle-Regeln
  (`open → next → in-progress → done`; reine `git mv`-Commits, siehe
  [`AGENTS.md` §3.3](../AGENTS.md#33-git-mv--inhaltsänderung--zwei-commits)).

## Safety and scope boundaries

- `a-check` ist ein **Lese-Tool**: Es schreibt nie in das geprüfte
  Repository — es liest *fremde* Quellbäume (C++/Go/Rust/Kotlin) und
  meldet Architektur-Verstöße (Kernvertrag
  [`AC-QA-02`](../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)).
- **Ehrliche Heuristik-Grenze:** die Extraktion ist text-/regex-basiert,
  kein vollständiger Parser je Sprache; die Grenze wird ausgewiesen, nicht
  als Vollständigkeit ausgegeben
  ([`AC-QA-02`](../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)).
- Determinismus ist Kernvertrag
  ([`AC-QA-01`](../spec/lastenheft.md#ac-qa-01--determinismus)):
  identische Eingabe ⇒ identische Ausgabe, stabil sortiert.
- Hermetik: der Scan läuft ohne Netz (`--network none`), distroless
  Runtime; Images sind digest-gepinnt
  ([`AC-QA-03`](../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit)).
- Dieses Repo ist kein produktiver Service; das Produkt ist ein
  CLI-Tool/Container-Image plus mitgelieferte `a-check.mk`.

## Minimal agent workflow

1. Diese Datei lesen.
2. Relevante kanonische Quelle lesen.
3. Betroffene IDs identifizieren.
4. Kleinste Änderung planen.
5. Engsten nützlichen Sensor laufen lassen (sobald Gates existieren).
6. Repo-weiten Gate-Lauf vor Handoff (`make gates`, sobald slice-003 ihn anlegt).
7. Doku/Indizes aktualisieren, falls ein öffentlicher Vertrag berührt.
8. Ausgeführte Sensors und verbleibende Risiken berichten.
