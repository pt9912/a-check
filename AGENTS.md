# AGENTS.md — Briefing für AI-Coding-Agenten

## 1. Was diese Datei ist

Onboarding-Briefing für jede AI-Session, die in diesem Repo Code oder
Dokumentation ändert. Sie verweist auf die kanonischen Quellen und
formuliert die Hard Rules, die der Implementation-Agent immer einhalten
muss.

**Bei Konflikt zwischen dieser Datei und einer kanonischen Quelle gilt
die kanonische Quelle** (Source Precedence — siehe
[`harness/README.md`](harness/README.md)).

Strukturregeln (ID-Schemata, Verzeichniskonvention, Adaptionen ggü.
Baseline, Modus-Deklarationen pro Sub-Area) leben in
[`harness/conventions.md`](harness/conventions.md).

Das Betriebsregelwerk der adoptierten Baseline in Agenten-Kurzform
einmal pro Session lesen, bevor der Workflow (§6) startet. Lese-Form:
das nach Modulen und Grundlagen-Abschnitten aufgeteilte Release-Bundle
[`lab-regelwerk.zip`](https://github.com/pt9912/ai-harness-course/releases/download/v1.3.0/lab-regelwerk.zip)
(`v1.3.0`) — so lädt ein Agent einen einzelnen Abschnitt, ohne das
gesamte Regelwerk im Kontext zu halten. Das Bundle ist eine derivative
Sicht auf die Quelldatei
[`agents-regelwerk.md`](https://raw.githubusercontent.com/pt9912/ai-harness-course/v1.3.0/kurs/de/agents-regelwerk.md);
bei Konflikt gilt die Quelldatei, über ihr die kanonischen Quellen
(Source Precedence). Der adoptierte Stand steht in
[`harness/conventions.md`](harness/conventions.md) §Baseline.

## 2. Kanonische Quellen (Source Precedence)

In dieser Reihenfolge:

1. [`spec/lastenheft.md`](spec/lastenheft.md) — vertraglich abnahmebindend.
2. [`spec/spezifikation.md`](spec/spezifikation.md) — technisch verbindlich, fortschreibbar.
3. [`spec/architecture.md`](spec/architecture.md) — Komponenten- und Sequenzsicht (sprach-/meilensteinfrei).
4. [`docs/plan/adr/README.md`](docs/plan/adr/README.md) — ADR-Index.
5. [`docs/plan/planning/in-progress/roadmap.md`](docs/plan/planning/in-progress/roadmap.md) — aktuelle Welle.
6. [`README.md`](README.md) — Projekt-Überblick.
7. **AGENTS.md (diese Datei).**
8. [`harness/README.md`](harness/README.md) — Harness-Einstieg.

## 3. Harte Regeln

### 3.1 Docker/make-only

Implementierungssprache ist **Go** (Fundament-ADR, entsteht mit slice-001):
ein statisches, sprach-agnostisches Binary, das *fremde* Quellen
text-heuristisch prüft. Es gilt: **kein Host-Go und keine
Host-Paketmanager** (`go`, `pip`, `npm`, `cargo`, `apt`, `brew`, …). Alle
Checks laufen über `make`; die Go-Toolchain läuft in Docker. Der Host
braucht nur `git`, GNU `make`, `bash` und Docker.

**Falsch:** `go build ./…`, `go test ./…`
**Richtig:** `make gates` (Implementierungs-Gates entstehen mit slice-003)

**Begründung:** Toolchain-Reproduzierbarkeit + Supply-Chain-Defense.

### 3.2 Suppression-Verbot

Inline-Suppressions sind verboten (`//nolint` o. Ä.). Ausnahmen leben
zentral in der Lint-Konfiguration mit Begründung (entsteht mit slice-003).

### 3.3 git mv + Inhaltsänderung = zwei Commits

Datei verschoben **und** Inhalt umgeschrieben: (1) `git mv` als eigener
Commit (Git erkennt R-Rename), (2) Inhalt umschreiben als zweiter Commit.
Sonst fällt die Rename-Detection unter die 50 %-Schwelle und
`git log --follow` wird unzuverlässig.

### 3.4 Architektur sprach-/meilensteinfrei; Spec-Straten nie abwärts

`spec/architecture.md` benennt Schichten und Rollen statt Technologie.
Kein Spec-Stratum (auch `spec/spezifikation.md`) referenziert ADRs,
Wellen, Slices, Commit-Hashes oder Closure-Daten. Die sprachkonkrete
Übersetzung und die Begründungen leben in den ADRs (`Schärft:`-Feld
aufwärts); die zeitliche Schicht in `docs/plan/planning/`.

### 3.5 ADRs sind nach `Accepted` immutable

Eine ADR mit Status `Accepted` wird nicht inhaltlich überschrieben.
Korrekturen entstehen als neue ADR mit `Supersedes ADR-NN`.

### 3.6 Gates dürfen nicht ohne ADR gelockert werden

Jede Schwellen-Senkung (Coverage, Linter-Strenge, Prüfregel) ist ein
ADR, kein PR-Kommentar.

## 4. Quality Gates

Nur hier gelistete Targets existieren im Makefile. Halluzinierte Gates
sind die häufigste Form von Harness-Lüge. `doc-check` (Bootstrap) sowie
`lint`/`test`/`coverage-gate`/`arch-check`/`gates` (slice-003) sind **real**
und grün; jede Gate ist eine Dockerfile-Stage.

| Target | Zweck | Stand |
|---|---|---|
| `make doc-check` | Doku-Links/Anker/Kennungen via `d-check` (Schwester-Tool, digest-gepinnt, netzlos, read-only) | **real** (Bootstrap) |
| `make lint` | golangci-lint mit dem Projekt-Profil (§3.2, [ADR-0005](docs/plan/adr/0005-lint-profil.md)) | **real** (slice-003) |
| `make test` | Akzeptanzkriterien der `AC-FA-*` als Go-Tests | **real** (slice-003) |
| `make coverage-gate` | Gesamt-Coverage ≥ 90 % über `./internal/...` ([ADR-0006](docs/plan/adr/0006-coverage-gate.md)) | **real** (slice-003) |
| `make arch-check` | Eigen-Architektur via `a-check` selbst (Dogfooding) | **real** (slice-003) |
| `make gates` | alle inneren Gates (mandatory vor Handoff) | **real** (slice-003) |

## 5. Dokumentations-Regeln

- Commits/PRs müssen mindestens eine `AC-*`- oder `ADR-*`-ID nennen.
  IDs werden nur beim Spec-/ADR-Schreiben nach dem deklarierten Schema
  vergeben (siehe [`harness/conventions.md`](harness/conventions.md)) —
  nie ad hoc im Commit/PR; Agenten referenzieren IDs, sie erfinden keine.
- Neue oder geänderte `AC-*`-Anforderungen entstehen nur in
  [`spec/lastenheft.md`](spec/lastenheft.md) — nie per ADR (ADRs schärfen
  die Spezifikation, nicht das Lastenheft).
- Neue ADRs müssen den ADR-Index aktualisieren.
- Roadmap/Status-Geschichte lebt in `docs/plan/planning/`, nicht in der
  Architektur-Spec.
- Slice-Lifecycle (`open → next → in-progress → done`) ist reine
  Datei-Bewegung (`git mv`, siehe §3.3).

## 6. Minimal Agent Workflow

Pro Slice:

1. [`harness/README.md`](harness/README.md) lesen.
2. Relevante kanonische Quelle lesen (Source Precedence beachten).
3. Betroffene Requirement-/ADR-IDs identifizieren.
4. Kleinste sinnvolle Änderung planen.
5. Engsten nützlichen Sensor laufen lassen.
6. Repo-weiten Gate-Lauf vor Handoff (`make gates`, sobald slice-003 ihn anlegt).
7. Doku/Indizes aktualisieren, falls ein öffentlicher Vertrag berührt.
8. Ausgeführte Sensors und verbleibende Risiken berichten — keine
   Erfolgsmeldung ohne Gate-Ausführung.
