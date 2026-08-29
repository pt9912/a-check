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
| 2 | [`spec/spezifikation.md`](../spec/spezifikation.md) | technisch verbindlich, fortschreibbar |
| 3 | [`spec/architecture.md`](../spec/architecture.md) | Komponenten/Sequenzen, meilensteinfrei |
| 4 | [`docs/plan/adr/`](../docs/plan/adr/) | Architekturentscheidungen |
| 5 | [`docs/plan/planning/in-progress/roadmap.md`](../docs/plan/planning/in-progress/roadmap.md) | aktuelle Welle |
| 6 | [`docs/user/`](../docs/user/) | Benutzer-/Betriebs-Doku ([Benutzerhandbuch](../docs/user/benutzerhandbuch.md)) |
| 7 | [`README.md`](../README.md) | Projekt-Überblick |
| 8 | [`AGENTS.md`](../AGENTS.md) | Agent-Briefing |
| 9 | diese Datei | Harness-Einstieg |

Neun Ränge inkl. `docs/user`-Stratum (Benutzerhandbuch); der zuvor ausgelassene
Rang ist mit [`MR-003`](conventions.md#mr-003--source-precedence-ohne-docsuser-rang)
(aufgelöst) eingefügt.

## Guides (Feedforward-Quellen)

| Quelle | Inhalt |
|---|---|
| [`spec/lastenheft.md`](../spec/lastenheft.md) | Anforderungen (`AC-FA-*`, `AC-QA-*`), Akzeptanzkriterien |
| [`spec/spezifikation.md`](../spec/spezifikation.md) | `.a-check.yml`-Schema, Extraktions-Algorithmus, Regel-Semantik, Defaults, Exit-Codes (`SPEC-*`) |
| [`spec/architecture.md`](../spec/architecture.md) | Hexagon-Komponenten/Rollen, Zugriffs-Constraints, Scan-Sequenz (`ARC-*`) |
| [`docs/plan/adr/`](../docs/plan/adr/) | Architekturentscheidungen |
| [`docs/plan/planning/`](../docs/plan/planning/) | Slice-Pläne und Roadmap |
| [`AGENTS.md`](../AGENTS.md) | Hard Rules, Source Precedence, Workflow |
| [`conventions.md`](conventions.md) | repo-lokale Strukturregeln, Adaptions-Block (`MR-*`), Modus-Deklarationen |
| [`.harness/baseline/v5.12.0/regelwerk/`](../.harness/baseline/v5.12.0/regelwerk/README.md) | adoptiertes Betriebsregelwerk der Baseline, **committet vendored** (netzlos): 17 Module + acht Grundlagen-Abschnitte, eine Datei je Abschnitt — einmal pro Session den zur Aufgabe gehörenden Abschnitt lesen, nie das ganze Bundle. Ziel-Formen daneben unter [`templates/`](../.harness/baseline/v5.12.0/templates/README.md), Integrität via `SHA256SUMS`. Derivativ (didaktik-freier Extrakt) — Stand und Begründung: [`conventions.md` §Baseline](conventions.md#baseline) / [`MR-006`](conventions.md#mr-006--baseline-committet-vendored-statt-per-url-referenziert) |

## Sensors (Feedback-Gates)

Nur Targets, die im Makefile **existieren**, dürfen hier stehen — halluzinierte
Gates sind die häufigste Form von Harness-Lüge; die Übereinstimmung Doku ↔
Makefile erzwingt `make doc-targets` mechanisch — über **diese** Tabelle **und**
`AGENTS.md` §4 (bis slice-079 tat das `gate-consistency`). Die Code-Gates sind
Dockerfile-Stages (Muster d-check/u-boot, digest-gepinnte Bases); die Meta-/
Harness-Gates laufen als Host-Bash. Die Durchsetzungsschicht deckt Tool-Call-,
Handoff- und Meta-Gate ab; die PR-/Push-CI
([`.github/workflows/ci.yml`](../.github/workflows/ci.yml)) zieht `make ci` +
`make trace-check` auf jede Integration und schließt die
Stop-Hook-„frischer-Klon"-Restlücke.

<!--
Drei Spalten — KEIN Lauf-Status (Form des Baseline-Templates
`.harness/baseline/v5.12.0/templates/harness/README.template.md` §Sensors —
vendored und damit netzlos nachschlagbar; Stand siehe conventions.md
§Baseline). Die Bindung-Spalte trägt STRUKTURELLE Referenzen (AC-/ADR-/CO-/
Slice-ID, Schwelle, Image-Hash) — nicht, ob ein Gate gerade grün ist.
Lauf-Wahrheit pro Commit liegt in der CI, nicht in diesem Rang-9-Dokument.
-->

| Target | Vertrag | Bindung |
|---|---|---|
| `make doc-check` | Links/Anker/Kennungs-Linkpflicht/Referenzmatrix der Repo-Doku via `d-check` (Schwester-Tool, digest-gepinnt, `--network none`, read-only) | Harness-Prozess (Doku-Hygiene; Dogfooding des Stacks); Bootstrap-Gate |
| `make lint` | golangci-lint mit Projekt-Profil; Inline-Suppressions verboten | [`ADR-0005`](../docs/plan/adr/0005-lint-profil.md) (Lint-Profil); slice-003 |
| `make test` | Akzeptanzkriterien der bezogenen `AC-FA-*` als Tests; Determinismus-Test | [`AC-QA-01`](../spec/lastenheft.md#ac-qa-01--determinismus) (AC-Bindung); slice-003 |
| `make coverage-gate` | Gesamt-Coverage ≥ Schwelle über `./internal/...` (`-coverpkg`, `tools/coverage-gate.sh`) | Kalibrierungs-Bindung **90 %** seit 2026-06-21 ([`ADR-0006`](../docs/plan/adr/0006-coverage-gate.md); Senkung nur per ADR, [`AGENTS.md` §3.6](../AGENTS.md#36-gates-dürfen-nicht-ohne-adr-gelockert-werden)); slice-003 |
| `make arch-check` | Eigen-Architektur via `a-check` selbst (Dogfooding) | [`AC-QA-02`](../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze) (AC-Bindung); slice-003 |
| `make gate-consistency` | Meta-Gate: `.d-check.yml`-Module (Schutz gegen Gate-Drift) + Pin-Konsistenz (Digest-Gleichheit `a-check.mk`/`README.md`+`README.de.md` == `version.md#aktuell`, Version ↔ CHANGELOG, `d-check.mk`-Tag/Digest-Deklaration) + **ADR-Index-Vollständigkeit** (jede ADR-Datei ist im Index verlinkt — die Gegenrichtung, die `doc-check` per Konstruktion nicht sieht) | Harness-Prozess ([`AC-QA-02`](../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze) für die Modul-Integrität); slice-004, Pin-Konsistenz slice-018, ADR-Index slice-087 |
| `make doc-targets` | Deklarations-Konsistenz Doku ↔ Build-Targets (`d-check`-Modul `targets`): jedes hier oder in `AGENTS.md` §4 dokumentierte Target existiert real, und jedes reale Gate-Target ist in `AGENTS.md` §4 gelistet | [`DC-FA-TGT-001`](../.d-check.yml); slice-074 (konfiguriert), slice-079 (im `gates`-Aggregat, löst `gate-consistency` (1)+(2) ab) |
| `make record-gates` | inhaltsbasierter Working-Tree-Hash-Nachweis für den `.claude`-Stop-Hook (Handoff-Gate) | Harness-Prozess (Durchsetzungsschicht); slice-004 |
| `make guard-selftest` | Selbsttest des PreToolUse-Command-Guard (`.claude/hooks/`): Host-Toolchain fail-closed geblockt, `make`/`git`/`docker` durchgelassen | Harness-Prozess (Tool-Call-Gate; [`AGENTS.md` §3.1](../AGENTS.md#31-dockermake-only)); slice-005 |
| `make gates` | aggregiert die inneren Gates (lint/test/coverage-gate/arch-check/doc-check/doc-targets/gate-consistency/suppression-check/guard-selftest) + `record-gates` als letzter Schritt | — (Aggregat) |
| `make image-test` | Distributions-Akzeptanz (`--print-mk`/`--print-config`/`--print-graph`/unbekanntes Flag) + Fragment-Parität (committete [`a-check.mk`](../a-check.mk) == `--print-mk`-Output) + nativ==Container-Determinismus eines Scans gegen das gebaute Image | [`AC-FA-DIST-001`](../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk)/[`AC-QA-02`](../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze); slice-006, Fragment-Parität slice-034 |
| `make ci` | CI-äquivalent: `gates` + `image-test` (Engine des Workflows `.github/workflows/ci.yml`) | — (Aggregat) |
| `make trace-check` | Traceability via Modul `commits`: jede Commit-Message nennt `AC-*`/`ADR-*`/`MR-*`/`slice-NNN` (`MSGFILE=` Hook, `RANGE=` CI) | [`ADR-0021`](../docs/plan/adr/0021-commits-modul-trace-check.md); Harness-Prozess ([`AGENTS.md` §5](../AGENTS.md#5-dokumentations-regeln)); slice-006, Modul seit slice-030 |
| `make verify-observations` | Deckung des Beobachtungs-Registers: zitierte `BEO-NNN` hat eine Zeile, jede Zeile trägt formgebundene Belege (`slice-NNN`, Anzahl == Zähler); Lage/Existenz der Beleg-Datei ausdrücklich ungeprüft | Harness-Prozess (Regelwerk `modul-06` §Das Beobachtungs-Register); slice-102 |
| `make doc-immutable` | ADR-Immutabilität über die Commit-Range (`d-check`-Modul `vcs`; `RANGE=`/`STAGED=1`) | [`AGENTS.md` §3.5](../AGENTS.md#35-adrs-sind-nach-accepted-immutable); slice-029, CI-durchgesetzt |

**Aktueller Lauf-Status:** CI-Badge im [`README.md`](../README.md) bzw. lokal
`make help` / `make gates`.
**Rote Gates:** Begründung im verlinkten `CO-<NNN>` (siehe Bindung-Spalte).

## Traceability rules

- PRs/Commits **müssen** mindestens eine `AC-*`- oder `ADR-*`-ID nennen
  (`MR-*`/`slice-NNN` gelten ebenso) — erzwungen durch `make trace-check`
  (lokal `HEAD`, CI über den Commit-Range, slice-006). Optional pro Klon:
  `make hooks` installiert den lokalen `commit-msg`-Hook (`.githooks`,
  slice-008), der dieselbe Prüfung schon vor dem Commit feuert.
- Neue oder geänderte Anforderungen brauchen einen Beleg: Test, Gate, Demo oder ADR.
- Neue ADRs müssen im [ADR-Index](../docs/plan/adr/README.md) ergänzt werden.
- Änderungen an Planning-Dokumenten folgen den Lifecycle-Regeln
  (`open → next → in-progress → done`; reine `git mv`-Commits, siehe
  [`AGENTS.md` §3.3](../AGENTS.md#33-git-mv--inhaltsänderung--zwei-commits)).

## Safety and scope boundaries

- `a-check` ist ein **Lese-Tool**: Es schreibt nie in das geprüfte
  Repository — es liest *fremde* Quellbäume (C++/Go/Rust/Kotlin/Java/Python/C#/TypeScript) und
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

## Rollen und ihre Übergabe-Artefakte

`modul-08` nennt sechs Rollen (Planner → Architect → Implementation → Reviewer → Verifier →
Validator) und **neun Übergaben**. Die Regel dahinter: *ohne Artefakt gibt es keinen Rollenwechsel,
nur einen Kontext-Switch* — und Rollen-Trennung ist **Kontext**-Trennung, nicht
Personen-Trennung. Eine Instanz darf mehrere Rollen spielen, aber nicht im selben Kontextfenster;
sonst wiederholen sich die blinden Flecken.

Sieben der neun Artefakte existieren in a-check bereits — unter eigenen Namen. Die Zuordnung
(angelegt in slice-066, Fund **B-9**):

| Übergabe | Artefakt in a-check |
|---|---|
| Planner → Architect | Slice-Plan mit `AC-*`-Bezug (§Deckt / §Betroffene Module) |
| Architect → Planner | ADR unter [`docs/plan/adr/`](../docs/plan/adr/README.md) mit `Schärft:`-Feld |
| Planner → Implementation | die Slice-Datei in `in-progress/` — der `git mv` **ist** die Übergabe |
| Implementation → Reviewer | Commit-Range plus Slice-Plan-Verweis |
| Reviewer → Implementation | Review-Report unter [`docs/reviews/`](../docs/reviews/README.md), Findings HIGH/MEDIUM/LOW/INFO mit Kopf-Metadaten |
| Implementation → Verifier | abgehakte DoD-Punkte plus Sensor-Belege (Gate-Ausgabe mit Exit-Code) |
| Verifier → Planner | `make verify` (Exit-Code) plus Closure-Notiz mit zwei beobachtbaren Kriterien |
| Verifier → Validator | **unverkörpert**, deklariert als [MR-009](conventions.md#mr-009--validator-rolle-unbesetzt-zwei-übergaben-ohne-artefakt) |
| Validator → Planner | **unverkörpert**, deklariert als [MR-009](conventions.md#mr-009--validator-rolle-unbesetzt-zwei-übergaben-ohne-artefakt) |

**Die Validator-Kante fehlt, und das ist benannt statt erfunden.** Validation fragt „bauen wir das
Richtige?" (gegen realen Bedarf), Verifikation „bauen wir es richtig?" (gegen Plan/DoD). a-check
verifiziert maschinell, validiert aber nicht über ein Artefakt: Rückmeldung der Konsumenten läuft
heute über Issues und über die Adoption eines Releases, nicht über einen Validierungsbeleg. Der
gefährlichste Fall wäre *Verifikation grün, Validation rot* — perfekt das Falsche gebaut. Wer die
Kante schließen will, braucht einen Abnehmer außerhalb des Repos und ein Artefakt, das dessen
Urteil festhält; beides ist heute nicht vorhanden.

**Kontext-Trennung, real angewandt:** die Review-Serie vom 2026-07-26 lief in einem **anderen
Kontextfenster** als die Implementierung — formal eine Rollen-Trennung, aber in derselben
Modell-Familie. Die Reports weisen sich darum ausdrücklich als *Selbst-Review* aus; ein
unabhängiger Lauf bleibt eine eigene Übergabe.

## Minimal agent workflow

1. Diese Datei lesen.
2. Relevante kanonische Quelle lesen.
3. Betroffene IDs identifizieren.
4. Kleinste Änderung planen.
5. Engsten nützlichen Sensor laufen lassen (sobald Gates existieren).
6. Repo-weiten Gate-Lauf vor Handoff (`make gates`, sobald slice-003 ihn anlegt).
7. Doku/Indizes aktualisieren, falls ein öffentlicher Vertrag berührt.
8. Ausgeführte Sensors und verbleibende Risiken berichten.

## Leseordnung

Für den neuen Menschen — **geordnet**, nicht vollständig. Eine Leseordnung, die alles nennt, ist
keine (Baseline-Regelwerk `grundlagen-harness-dateien.md` §harness/README.md als Einstiegspunkt).

1. [`README.md`](../README.md) — was `a-check` ist und wofür, in fünf Minuten.
2. [`AGENTS.md` §3 Harte Regeln](../AGENTS.md#3-harte-regeln) — was hier **nie** getan wird.
   Wer nur einen Abschnitt liest, liest diesen.
3. [`spec/lastenheft.md`](../spec/lastenheft.md) — was vertraglich zugesagt ist; alles Weitere
   präzisiert das nur.
4. Bei Bedarf: [`conventions.md`](conventions.md) — ID-Schemata, Adaptionen ggü. der Baseline,
   Modus je Sub-Area. Nachschlagewerk, keine Vorab-Lektüre.

Für **Agenten** gilt stattdessen der Pfad oben unter §Minimal agent workflow; er beginnt bei
dieser Datei, nicht beim Projekt-Überblick.
