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

Das Betriebsregelwerk der adoptierten Baseline ist **Nachschlagewerk pro
Entscheidung**, keine Pro-Session-Lektüre: Wird ein Abschnitt gebraucht, wird
**dieser eine** gelesen — nie das ganze Bundle (Kontext-Hygiene). Auswahlregel
nach Aufgabe: Slice schneiden oder schließen → `modul-05`; ADR schreiben →
`modul-04`; Gate anlegen oder ändern → `modul-13`; Review führen → `modul-10`;
DoD/Closure prüfen → `modul-11`; Ausnahme oder Diskrepanz einordnen →
`modul-07`; Modus einer Sub-Area bestimmen → `modul-02` und
`grundlagen-konventionen`; Release → `modul-16`. Es liegt **committet
vendored** im Repo, also netzlos verfügbar:
[`.harness/baseline/v3.5.2/regelwerk/README.md`](.harness/baseline/v3.5.2/regelwerk/README.md)
ist der Index (17 Module + drei Grundlagen-Abschnitte, eine Datei je
Abschnitt); die Ziel-Formen daneben unter
[`templates/`](.harness/baseline/v3.5.2/templates/README.md). Integrität:
`.harness/baseline/v3.5.2/SHA256SUMS`.

Das vendored Regelwerk ist ein **didaktik-freier Extrakt** und trägt keine
eigene Normativität: bei Konflikt gilt der Kurs
([`v3.5.2`](https://github.com/pt9912/ai-harness-course/tree/v3.5.2)), über
ihm die kanonischen Quellen (Source Precedence). Der adoptierte Stand und
die Vendoring-Begründung stehen in
[`harness/conventions.md`](harness/conventions.md) §Baseline bzw.
[`MR-006`](harness/conventions.md#mr-006--baseline-committet-vendored-statt-per-url-referenziert).

## 2. Kanonische Quellen (Source Precedence)

In dieser Reihenfolge:

1. [`spec/lastenheft.md`](spec/lastenheft.md) — vertraglich abnahmebindend.
2. [`spec/spezifikation.md`](spec/spezifikation.md) — technisch verbindlich, fortschreibbar.
3. [`spec/architecture.md`](spec/architecture.md) — Komponenten- und Sequenzsicht (sprach-/meilensteinfrei).
4. [`docs/plan/adr/README.md`](docs/plan/adr/README.md) — ADR-Index.
5. [`docs/plan/planning/in-progress/roadmap.md`](docs/plan/planning/in-progress/roadmap.md) — aktuelle Welle.
6. [`docs/user/`](docs/user/) — Benutzer-/Betriebs-Doku ([Benutzerhandbuch](docs/user/benutzerhandbuch.md)).
7. [`README.md`](README.md) — Projekt-Überblick.
8. **AGENTS.md (diese Datei).**
9. [`harness/README.md`](harness/README.md) — Harness-Einstieg.

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

**Durchsetzung:** Ein PreToolUse-Command-Guard
(`.claude/hooks/pretooluse-command-guard.sh`, slice-005) lehnt Host-Toolchain-
und Paketmanager-Aufrufe (`go`/`golangci-lint`/`pip`/`npm`/`cargo`/`apt`/`brew`/…)
**vor** der Ausführung fail-closed ab (Tool-Call-Gate der Durchsetzungsschicht);
`make gates` belegt ihn über `make guard-selftest`.

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
sind die häufigste Form von Harness-Lüge; `make gate-consistency` erzwingt
die Übereinstimmung Doku ↔ Makefile mechanisch. Die Code-Gates sind
Dockerfile-Stages, die Meta-Gates laufen als Host-Bash. **Mandatory** ist,
was im `gates`-Aggregat hängt; die `doc-*`-Targets jenseits von `doc-check`
und `doc-immutable` sind **advisory** (`d-check`-Funktionen, nicht im
Aggregat). Ob ein Gate gerade grün ist, sagt die CI (Badge im
[`README.md`](README.md)), nicht diese Tabelle.

| Target | Zweck |
|---|---|
| `make doc-check` | Doku-Links/Anker/Kennungen via `d-check` (Schwester-Tool, digest-gepinnt, netzlos, read-only) |
| `make doc-trace` | advisory Requirements Traceability Matrix via `d-check` (DC-FA-CLI-009; `TRACE_FLAGS=--json`) |
| `make doc-complete` | Vollständigkeits-Gate: Requirements-Waise ⇒ Exit 1 (DC-FA-CLI-011) |
| `make doc-doctor` | erklärende Diagnose mit Fix-Kandidaten (DC-FA-CLI-007) |
| `make doc-repair` | Reparatur-Patch (unified diff) auf stdout, git-apply-rein (DC-FA-CLI-008) |
| `make doc-immutable` | ADR-Immutabilität (§3.5) via git-Diff (Modul `vcs`; `RANGE=`/`STAGED=1`, DC-FA-VCS-001) — **CI-durchgesetzt** über die Commit-Range ([`ci.yml`](.github/workflows/ci.yml)) |
| `make doc-commits` | Commit-Message-Traceability (Modul `commits`; `RANGE=`, DC-FA-COMMITS-001) |
| `make doc-planning` | Planning-Lifecycle-Konsistenz Roadmap ↔ `in-progress` (Modul `planning`, DC-FA-PLAN-001) |
| `make doc-tracked` | Getrackt-Status auflösbarer Referenz-Ziele (Modul `tracked`, DC-FA-TRK-001) |
| `make doc-targets` | Deklarations-Konsistenz Doku ↔ Build-Targets (Modul `targets`, DC-FA-TGT-001; neu in v0.51.1) |
| `make doc-help` | Liste der `doc-*`-Targets (Utility) |
| `make lint` | golangci-lint mit dem Projekt-Profil (§3.2, [ADR-0005](docs/plan/adr/0005-lint-profil.md)) |
| `make test` | Akzeptanzkriterien der `AC-FA-*` als Go-Tests |
| `make coverage-gate` | Gesamt-Coverage ≥ 90 % über `./internal/...` ([ADR-0006](docs/plan/adr/0006-coverage-gate.md)) |
| `make arch-check` | Eigen-Architektur via `a-check` selbst (Dogfooding) |
| `make gate-consistency` | Meta-Gate: dokumentierte Targets ↔ Makefile, `.d-check.yml`-Module (Harness-Lügen-Schutz) + Pin-Konsistenz (Digest-Gleichheit harte Pins == `version.md#aktuell`, Version == CHANGELOG, `d-check.mk`-Deklaration; slice-018) |
| `make record-gates` | Gate-Nachweis (Working-Tree-Hash) für den Stop-Hook |
| `make suppression-check` | Fitness Function zum Suppression-Verbot (§3.2, [ADR-0005](docs/plan/adr/0005-lint-profil.md)): keine `//nolint`-Direktive in den Go-Quellen — `nolintlint` prüft nur Wohlgeformtheit, nicht Existenz (slice-049) |
| `make guard-selftest` | Selbsttest des PreToolUse-Command-Guard (Tool-Call-Gate §3.1) |
| `make regelwerk-check` | **kein Gate** — Wartung der vendored Baseline ([MR-006](harness/conventions.md#mr-006--baseline-committet-vendored-statt-per-url-referenziert)): Integrität gegen `SHA256SUMS` fail-closed; die Freshness-Hälfte bleibt als Netz-Operation ausdrücklich ungeprüft |
| `make gates` | alle inneren Gates (mandatory vor Handoff) |
| `make verify-slice-form` | Form der Slice-Pläne ab slice-052: höchstens drei DoD-Punkte, benannte Lerneintrag-Form; ältere grandfathered (bootstrap-aware, slice-052) |
| `make verify-ac-form` | Form neuer `AC-*` (§5): Happy · Boundary · Negative · Out-of-Scope; die 19 bei Einführung bestehenden sind grandfathered (slice-054) |
| `make verify-closure-notes` | Struktur der Closure-Notizen in `done/` (§5): genau eine, ausgefüllt, kein Platzhalter, keine Floskel (slice-050) |
| `make verify` | **Verifikations-Schicht** (getrennt von `gates`, Regelwerk Modul 11): beantwortet DoD-/Closure-Fragen statt Code-Fragen; vor der „fertig"-Meldung auszuführen |
| `make image-test` | [AC-FA-DIST-001](spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk) + nativ==Container-Akzeptanz + Fragment-Parität (committete [`a-check.mk`](a-check.mk) == `--print-mk`, slice-034) gegen das gebaute Image |
| `make ci` | CI-äquivalent: `gates` + `image-test` (Workflow `.github/workflows/ci.yml`) |
| `make trace-check` | Traceability via Modul `commits` ([ADR-0021](docs/plan/adr/0021-commits-modul-trace-check.md)): `AC-*`/`ADR-*`/`MR-*`/`slice`-ID je Commit (§5; `MSGFILE=` Hook, `RANGE=` CI) |

## 5. Dokumentations-Regeln

- Commits/PRs müssen mindestens eine `AC-*`- oder `ADR-*`-ID nennen
  (auch `MR-*`/`slice-NNN` gelten). Durchgesetzt durch `make trace-check`
  (slice-006; via d-check-Modul `commits`, [ADR-0021](docs/plan/adr/0021-commits-modul-trace-check.md))
  — lokal über `HEAD~1..HEAD`, in der CI über den Commit-Range
  ([`.github/workflows/ci.yml`](.github/workflows/ci.yml)). IDs werden nur
  beim Spec-/ADR-Schreiben nach dem deklarierten Schema vergeben (siehe
  [`harness/conventions.md`](harness/conventions.md)) — nie ad hoc im
  Commit/PR; Agenten referenzieren IDs, sie erfinden keine.
- Neue oder geänderte `AC-*`-Anforderungen entstehen nur in
  [`spec/lastenheft.md`](spec/lastenheft.md) — nie per ADR (ADRs schärfen
  die Spezifikation, nicht das Lastenheft).
- Neue ADRs müssen den ADR-Index aktualisieren.
- Roadmap/Status-Geschichte lebt in `docs/plan/planning/`, nicht in der
  Architektur-Spec.
- **Slice-Lifecycle** ist reine Datei-Bewegung (`git mv`, siehe §3.3) — der
  Zustand ist das Verzeichnis, kein Feld im Dokument. **Fünf** Übergänge, drei
  vorwärts und zwei zurück:

  | von → nach | Bedingung |
  |---|---|
  | `open/` → `next/` | für die nächste Welle priorisiert |
  | `next/` → `in-progress/` | Trigger eingetreten und WIP-Limit frei |
  | `in-progress/` → `done/` | DoD erfüllt, Closure-Notiz geschrieben, Gates grün |
  | `in-progress/` → `next/` | **Rückführung:** zu groß — zurück zur Zerlegung, nicht dehnen |
  | `in-progress/` → `open/` | **Rückführung:** blockiert, solange der Blocker steht |

  Der direkte Weg `open/ → in-progress/` bleibt zulässig; `next/` ist ein Ort,
  keine Pflichtstation ([`next/README.md`](docs/plan/planning/next/README.md)).
- **WIP-Limit = 1.** Es liegt **genau ein** Slice in `in-progress/` (die Roadmap
  zählt nicht mit). Das ist eine harte Größe, kein Vorschlag: zwei aktive Slices
  teilen sich einen Gate-Nachweis und eine Closure-Aufmerksamkeit, und beides
  trägt nur einmal.
- **AC-Form:** die Pflicht-Bausteine einer Anforderung stehen in
  [`harness/conventions.md`](harness/conventions.md) §Anforderungs-Anlege-Prozess
  — dort seit jeher die drei Pfade (Happy/Boundary/Negative im
  Given/When/Then-Stil) plus Out-of-Scope. Neu ist nur die **Durchsetzung**:
  `make verify` prüft sie ab slice-054 für **neue** `AC-*`; die **19** bei
  Einführung bestehenden sind **grandfathered** (vertraglich bindend, Rand- und
  Negativfälle bereits in Prosa — ein Umbau träfe die Form statt der Substanz),
  und die Grandfather-Liste wächst nicht mit.
- **Steering-Loop:** wiederkehrende Fehlermuster werden in
  [`docs/plan/steering-loop.md`](docs/plan/steering-loop.md) gezählt. Ab dem
  **zweiten** gleichartigen Vorfall entsteht ein Eintrag, ab dem **dritten** ist
  es eine Harness-Lücke und verlangt einen Guide oder Sensor — „besser
  aufpassen" ist keine Antwort. Ein Eintrag ohne Vorfallszahl ist unzulässig:
  die Zahl ist das Einzige, was die Schwelle prüfbar macht.
- **Slice-Form:** neue Slices entstehen aus
  [`docs/plan/planning/slice.template.md`](docs/plan/planning/slice.template.md). Sie trägt die
  Größen-Regel — **höchstens drei DoD-Punkte und höchstens zwei Schichten**; passt der Slice nicht
  hinein, wird er **zerlegt, nicht gedehnt** — und verlangt den Lerneintrag in einer von drei
  **benannten** Formen (geschärfte Regel · neuer Sensor · benannte Spec-Lücke). `make verify`
  prüft beides ab slice-052; ältere Slices sind grandfathered.
- **Closure-Pflicht:** ein Slice in `done/` trägt **genau einen**
  Closure-Abschnitt, und der ist ausgefüllt — kein Platzhalter, keine
  Floskel. Inhaltlich mindestens eines von dreien: ein **Lernsignal mit
  Ursache** („X, *weil* Y"), ein **konkretes Folge-Slice** oder eine
  **beobachtbare Architektur-Aussage**. Ohne Lerneintrag ist ein Slice
  nicht „fertig", sondern nur „weg". Die *strukturelle* Hälfte prüft
  `make verify` maschinell, die *semantische* (Inhalt vs. Floskel) der
  Skill [`.harness/skills/closure-note-reviewer.md`](.harness/skills/closure-note-reviewer.md)
  (slice-050).

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

Beim **Abschluss** eines Slice zusätzlich `make verify` (Verifikations-Schicht,
§4): `gates` beantwortet Code-Fragen, `verify` die DoD-/Closure-Fragen. Die
semantische Hälfte — trägt die Notiz ein Lernsignal oder nur eine Floskel? —
leistet der Skill
[`.harness/skills/closure-note-reviewer.md`](.harness/skills/closure-note-reviewer.md).
