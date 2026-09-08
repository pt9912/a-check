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

Neun Ränge inkl. `docs/user`-Stratum (Benutzerhandbuch).

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
| [`.harness/skills/reviewer.md`](../.harness/skills/reviewer.md) | Reviewer-Skill: HIGH-Liste, Kategorien-Regeln, Negativbefund-Pflicht, Output-Schema, §Mess-Regeln (Modul 10) — nächste Rolle nach Schritt 8 des Minimal Agent Workflow, nicht Teil der Implementer-Eingabe |
| [`.harness/baseline/v6.5.0/regelwerk/`](../.harness/baseline/v6.5.0/regelwerk/README.md) | adoptiertes Betriebsregelwerk der Baseline, **committet vendored** (netzlos): 17 Module + acht Grundlagen-Abschnitte, eine Datei je Abschnitt — einmal pro Session den zur Aufgabe gehörenden Abschnitt lesen, nie das ganze Bundle. Ziel-Formen daneben unter [`templates/`](../.harness/baseline/v6.5.0/templates/README.md), Integrität via `SHA256SUMS`. Derivativ (didaktik-freier Extrakt) — Stand und Begründung: [`conventions.md` §Baseline](conventions.md#baseline) / [`MR-006`](conventions.md#mr-006--baseline-committet-vendored-statt-per-url-referenziert) |

## Sensors (Feedback-Gates)

Jedes hier gelistete Target existiert im Makefile — `make doc-targets` erzwingt
das mechanisch, über **diese** Tabelle **und** `AGENTS.md` §4. Die Code-Gates sind
Dockerfile-Stages (Muster d-check/u-boot, digest-gepinnte Bases); die Meta-/
Harness-Gates laufen als Host-Bash. Die Durchsetzungsschicht deckt Tool-Call-,
Handoff- und Meta-Gate ab; die PR-/Push-CI
([`.github/workflows/ci.yml`](../.github/workflows/ci.yml)) zieht `make ci` +
`make trace-check` auf jede Integration und schließt die
Stop-Hook-„frischer-Klon"-Restlücke.

<!--
Drei Spalten — KEIN Lauf-Status (Form des Baseline-Templates
`.harness/baseline/v6.5.0/templates/harness/README.template.md` §Sensors —
vendored und damit netzlos nachschlagbar; Stand siehe conventions.md
§Baseline). Die Bindung-Spalte trägt STRUKTURELLE Referenzen (AC-/ADR-/CO-/
Slice-ID, Schwelle, Image-Hash) — nicht, ob ein Gate gerade grün ist.
Lauf-Wahrheit pro Commit liegt in der CI, nicht in diesem Rang-9-Dokument.
-->

| Target | Vertrag | Bindung |
|---|---|---|
| [`make doc-check`](sensors/doc-check.md) | Links, Anker, Kennungs-Linkpflicht und Referenzmatrix der Repo-Doku lösen auf; dazu die Lifecycle-Invariante wandernder Slices und die Versions-Kohärenz. Grenzen und Ausnahmen: siehe Datei | Harness-Prozess (Doku-Hygiene; Dogfooding des Stacks); Bootstrap-Gate; [`SL-002`](../docs/plan/planning/observations/README.md) seit slice-080 |
| `make lint` | golangci-lint mit Projekt-Profil; Inline-Suppressions verboten | [`ADR-0005`](../docs/plan/adr/0005-lint-profil.md) (Lint-Profil); slice-003 |
| `make test` | Akzeptanzkriterien der bezogenen `AC-FA-*` als Tests; Determinismus-Test | [`AC-QA-01`](../spec/lastenheft.md#ac-qa-01--determinismus) (AC-Bindung); slice-003 |
| `make coverage-gate` | Gesamt-Coverage ≥ Schwelle über `./internal/...` (`-coverpkg`, `tools/coverage-gate.sh`) | Kalibrierungs-Bindung **90 %** seit 2026-06-21 ([`ADR-0006`](../docs/plan/adr/0006-coverage-gate.md); Senkung nur per ADR, [`AGENTS.md` §3.6](../AGENTS.md#36-gates-dürfen-nicht-ohne-adr-gelockert-werden)); slice-003 |
| `make arch-check` | Eigen-Architektur via `a-check` selbst (Dogfooding) | [`AC-QA-02`](../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze) (AC-Bindung); slice-003 |
| [`make gate-consistency`](sensors/gate-consistency.md) | Meta-Gate, drei Prüfungen: `.d-check.yml`-Module, Pin-Konsistenz, ADR-Index-Vollständigkeit. Grenzen und Bindungen: siehe Sensor-Datei | Harness-Prozess ([`AC-QA-02`](../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze) für die Modul-Integrität); slice-004, Pin-Konsistenz slice-018, ADR-Index slice-087 |
| `make doc-targets` | Deklarations-Konsistenz Doku ↔ Build-Targets (`d-check`-Modul `targets`): jedes hier oder in `AGENTS.md` §4 dokumentierte Target existiert real, und jedes reale Gate-Target ist in `AGENTS.md` §4 gelistet | [`DC-FA-TGT-001`](../.d-check.yml); slice-074 (konfiguriert), slice-079 (im `gates`-Aggregat, löst `gate-consistency` (1)+(2) ab) |
| [`make doc-structure`](sensors/doc-structure.md) | Struktur-Invarianten innerhalb der Dokumente (`d-check`-Modul `structure`): Größen-Regel des DoD, Closure-Struktur, Lerneintrag-Form, Kopffelder, AC-Form, Zellengrenzen der Gate-Tabellen — konfiguriert in [`.d-check.yml`](../.d-check.yml) | `DC-FA-STRUCT-001`; slice-115 (Pin), slice-080 (konfiguriert, löst drei Eigenbau-Sensoren ab), slice-120 (`verify-ac-form` als vierter; Pin `v0.69.0` wegen `exempt-expect-count`) |
| `make doc-complete` | Vollständigkeits-Gate (`d-check --trace --require-complete`): eine Anforderung ohne referenzierenden Slice ⇒ Exit 1 | `DC-FA-CLI-011`; slice-123 (ins `verify`-Aggregat — davor advisory und nie gelaufen) |
| `make record-gates` | inhaltsbasierter Working-Tree-Hash-Nachweis für den `.claude`-Stop-Hook (Handoff-Gate) | Harness-Prozess (Durchsetzungsschicht); slice-004 |
| [`make dcheck-phrase-selftest`](sensors/dcheck-phrase-selftest.md) | Kalibrierung phrasen-basierter `d-check`-Konfigurationen in **beiden** Hälften: vier Werkzeug-Kontrollen gegen Fixtures und eine Korpus-Kontrolle gegen den `done/`-Bestand | Harness-Prozess; Antwort auf [`BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf`](../docs/plan/planning/observations/BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf/observation.md) bei 3×; Werkzeug-Hälfte seit slice-168, Korpus-Hälfte seit slice-169 |
| [`make symlink-check`](sensors/symlink-check.md) | Jeder getrackte Symlink löst auf, und ein Baseline-Ziel trägt den adoptierten Stand | Harness-Prozess; Antwort auf [`BEO-GATE/symlink-ziel-nach-baseline-bump-ungeprueft`](../docs/plan/planning/observations/BEO-GATE/symlink-ziel-nach-baseline-bump-ungeprueft/observation.md) bei 3×; slice-173 |
| `make guard-selftest` | Selbsttest des PreToolUse-Command-Guard (`.claude/hooks/`): Host-Toolchain fail-closed geblockt, `make`/`git`/`docker` durchgelassen | Harness-Prozess (Tool-Call-Gate; [`AGENTS.md` §3.1](../AGENTS.md#31-dockermake-only)); slice-005 |
| [`make ci-range-selftest`](sensors/ci-range-selftest.md) | Selbsttest der Commit-Range-Weiche der CI (`tools/ci-commit-range.sh`): vier Fälle, darunter der **Force-Push** — eine im Runner-Klon unerreichbare Basis fällt auf den Default-Branch statt abzubrechen | CI-Schicht; slice-134 |
| [`make image-scan`](sensors/image-scan.md) | CVE-Scan gegen das **publizierte** Image; über rot entscheiden nur **behebbare** CRITICAL/HIGH | [`ADR-0037`](../docs/plan/adr/0037-cve-scan-gegen-das-publizierte-image.md); **nicht** in `gates` (nicht hermetisch); slice-124 |
| [`make doc-planning`](sensors/doc-planning.md) | Äquivalenz Roadmap ↔ `in-progress/` (`d-check`-Modul `planning`): liegt dort ein Slice, fehlt der Ruhe-Marker; ist das Verzeichnis leer, steht er. **Kein Name** wird geprüft — Grenze in der Sensor-Datei | `DC-FA-PLAN-001`; slice-122 (konfiguriert und ins Aggregat — davor ohne Gegenstand, [`BEO-014`](../docs/plan/planning/observations/BEO-GATE/ruhe-marker-ungewaechtert/observation.md)) |
| [`make doc-workflows`](sensors/doc-workflows.md) | Deklarations-Form der `uses:`-Referenzen unter `.github/workflows`: voller SHA plus Tag-Kommentar beim Fremden, existierendes Ziel und gedeckte Rechte beim Lokalen. Prüft die **Form**, nicht die Gültigkeit | `DC-FA-WF-001`; slice-130 (konfiguriert und ins Aggregat; fand den latenten Release-Bruch in `release.yml`) |
| [`make version-coherence`](sensors/version-coherence.md) | Kohärenz doppelt deklarierter Versions-Angaben: ein `uses:`-SHA ⇒ ein Tag-Kommentar; eine Variable in `Makefile` **und** `Dockerfile` ⇒ ein Wert. Prüft **Divergenz**, nicht Wahrheit | slice-131 (Antwort auf [`BEO-026`](../docs/plan/planning/observations/BEO-GATE/versionsangabe-neben-digest-ungeprueft/observation.md) bei 3×) |
| `make suppression-check` | keine `//nolint`-Direktive in den Go-Quellen — `nolintlint` prüft nur Wohlgeformtheit, nicht Existenz | [`ADR-0005`](../docs/plan/adr/0005-lint-profil.md) ([`AGENTS.md` §3.2](../AGENTS.md#32-suppression-verbot)); slice-049 |
| `make gates` | aggregiert die inneren Gates und schließt mit `record-gates`. **Welche genau, sagt das [`Makefile`](../Makefile)** — eine Liste hier wäre eine zweite Quelle | — (Aggregat) |
| `make image-test` | Distributions-Akzeptanz (`--print-mk`/`--print-config`/`--print-graph`/unbekanntes Flag) + Fragment-Parität (committete [`a-check.mk`](../a-check.mk) == `--print-mk`-Output) + nativ==Container-Determinismus eines Scans gegen das gebaute Image | [`AC-FA-DIST-001`](../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk)/[`AC-QA-02`](../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze); slice-006, Fragment-Parität slice-034 |
| `make ci` | CI-äquivalent: `gates` + `image-test` (Engine des Workflows `.github/workflows/ci.yml`) | — (Aggregat) |
| `make trace-check` | Traceability via Modul `commits`: jede Commit-Message nennt `AC-*`/`ADR-*`/`MR-*`/`slice-NNN` (`MSGFILE=` Hook, `RANGE=` CI) | [`ADR-0021`](../docs/plan/adr/0021-commits-modul-trace-check.md); Harness-Prozess ([`AGENTS.md` §5](../AGENTS.md#5-dokumentations-regeln)); slice-006, Modul seit slice-030 |
| [`make verify-observations`](sensors/verify-observations.md) | Deckung des Beobachtungs-Registers: jeder in `done/` zitierte Pfad hat ein Verzeichnis, jedes Verzeichnis ein nicht leeres `evidence/`. Der Zähler wird abgeleitet, nicht geführt | Harness-Prozess (Regelwerk `modul-06` §Das Beobachtungs-Register); slice-102, Verzeichnisform slice-139 |
| [`make doc-mentions`](sensors/doc-mentions.md) | Jede Datei unter `harness/sensors/` ist in [`AGENTS.md`](../AGENTS.md) §4 genannt — die **Gegenrichtung** des Link-Checks | `DC-FA-MENT-001`; slice-184 (Modul seit `d-check v0.75.0`); im `gates`-Aggregat |
| [`make doc-reviews`](sensors/doc-reviews.md) | Review-Report-Deckung: eine `done/`-Slice-DoD-Zeile mit der Phrase „unabhängiger Review" braucht einen Report gleicher Kennung; **Opt-in pro Slice über die Phrase selbst** | `DC-FA-RVW-001`; slice-160 |
| [`make verify-risiko-ausgaenge`](sensors/verify-risiko-ausgaenge.md) | Jedes in §6 notierte Risiko trägt genau einen Ausgang aus der geschlossenen Dreier-Menge; geprüft in `done/` und in abschlussbereiten `in-progress/`-Slices | Harness-Prozess ([`AGENTS.md`](../AGENTS.md) §5); slice-102, slice-129 |
| `make doc-immutable` | ADR-Immutabilität über die Commit-Range (`d-check`-Modul `vcs`; `RANGE=`/`STAGED=1`) | [`AGENTS.md` §3.5](../AGENTS.md#35-adrs-sind-nach-accepted-immutable); slice-029, CI-durchgesetzt |

### Nicht-Gates

**Werkzeuge — genannt, weil der Lauf sie braucht, aber kein Gate**
(`modul-13` §Vorhanden ≠ behauptet). Die Bindung-Spalte trägt `kein Gate` in
der Zeile selbst.

| Target | Was es tut | Bindung |
|---|---|---|
| [`make archive-wave`](sensors/archive-wave.md) | **bewegt** — Zeitdokumente einer Welle oder eines wellenlosen Slice ins Archiv, Volltext wird Stub | `kein Gate`; Pflichtschritt der Slice-Closure seit slice-157 ([`AGENTS.md`](../AGENTS.md) §6) |
| `make slice-mv` | **bewegt** — Lifecycle-Wechsel eines Slice per `git mv` samt der Verweise auf ihn | `kein Gate`; Antwort auf [`BEO-PLAN/verweis-auf-wandernden-slice`](../docs/plan/planning/observations/BEO-PLAN/verweis-auf-wandernden-slice/observation.md) bei 3×; slice-118 |
| [`make regelwerk-check`](sensors/regelwerk-check.md) | **misst** — Integrität der vendored Baseline gegen `SHA256SUMS`, fail-closed; die Freshness-Hälfte bleibt als Netz-Operation ungeprüft | `kein Gate`; Wartung, [`MR-006`](conventions.md#mr-006--baseline-committet-vendored-statt-per-url-referenziert) |
| `make doc-repair` | **sagt**, was ein schreibender Lauf täte — Reparatur-Patch als unified diff auf stdout | `kein Gate`; `DC-FA-CLI-008` |
| `make doc-trace` · `make doc-doctor` · `make doc-usage` · `make doc-help` | **sagen** — Traceability-Matrix, Diagnose mit Fix-Kandidaten, Aufruf-Hilfe, Target-Liste | `kein Gate`; advisory, verfügbar aber nicht als Gate behauptet |

**Nicht hier, obwohl sie in keinem Aggregat hängen:** `make doc-tracked`,
`make doc-commits` und `make image-scan` — sie **urteilen** über einen Zustand
(Getrackt-Status, Commit-Traceability, CVE-Lage) und sind damit Gates, nur nicht
aggregierte. Das Kriterium ist **urteilen** gegen *bewegen · messen · sagen*,
nicht die Aggregat-Zugehörigkeit.

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

Sechs Rollen, neun Übergaben und die Regel *ohne Artefakt kein Rollenwechsel*:
`modul-08` §Die neun Übergaben. **Sieben** der neun sind in a-check verkörpert — unter eigenen
Namen, und diese Zuordnung steht nur hier:

| Übergabe | Artefakt in a-check |
|---|---|
| Planner → Architect | Slice-Plan mit `AC-*`-Bezug (§Deckt / §Betroffene Module) |
| Architect → Planner | ADR unter [`docs/plan/adr/`](../docs/plan/adr/README.md) mit `Schärft:`-Feld |
| Planner → Implementation | die Slice-Datei in `in-progress/` — der `git mv` **ist** die Übergabe |
| Implementation → Reviewer | Commit-Range plus Slice-Plan-Verweis |
| Reviewer → Implementation | Review-Report unter [`docs/reviews/`](../docs/reviews/README.md), Findings HIGH/MEDIUM/LOW/INFO mit Kopf-Metadaten |
| Implementation → Verifier | abgehakte DoD-Punkte plus Sensor-Belege (Gate-Ausgabe mit Exit-Code) |
| Verifier → Planner | `make verify` (Exit-Code) plus Closure-Notiz mit zwei beobachtbaren Kriterien |
| Verifier → Validator | **unverkörpert**, deklariert als [MR-016](conventions/MR-016-validator-unbesetzt.md) |
| Validator → Planner | **unverkörpert**, deklariert als [MR-016](conventions/MR-016-validator-unbesetzt.md) |

**Die beiden Validator-Kanten sind unbesetzt, und das ist benannt statt erfunden:**
[`MR-016`](conventions/MR-016-validator-unbesetzt.md) trägt Begründung und Auflösungs-Trigger.

## Minimal agent workflow

1. Diese Datei lesen.
2. Relevante kanonische Quelle lesen.
3. Betroffene IDs identifizieren.
4. Kleinste Änderung planen — die Plan-Ausgabe nennt Out-of-Scope (siehe [`AGENTS.md`](../AGENTS.md) §6).
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
