# AGENTS.md — Briefing für AI-Coding-Agenten

## 1. Was diese Datei ist

Onboarding-Briefing für jede AI-Session, die in diesem Repo Code oder
Dokumentation ändert. Sie verweist auf die kanonischen Quellen und
formuliert die Hard Rules, die der Implementation-Agent immer einhalten
muss.

Regeln dieser Datei: Baseline-Regelwerk `modul-09-implementierung.md` §Ziel-Form: AGENTS.md —
sie trägt Hard Rules und Pointer auf kanonische Quellen, sie dupliziert deren Inhalt nicht;
sonst entsteht Drift.

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
`grundlagen-bootstrap`; Release → `modul-16`. Es liegt **committet
vendored** im Repo, also netzlos verfügbar:
[`.harness/baseline/v6.5.0/regelwerk/README.md`](.harness/baseline/v6.5.0/regelwerk/README.md)
ist der Index (17 Module + acht Grundlagen-Abschnitte, eine Datei je
Abschnitt); die Ziel-Formen daneben unter
[`templates/`](.harness/baseline/v6.5.0/templates/README.md). Integrität:
`.harness/baseline/v6.5.0/SHA256SUMS`.

Das vendored Regelwerk ist ein **didaktik-freier Extrakt** und trägt keine
eigene Normativität: bei Konflikt gilt der Kurs
([`v6.5.0`](https://github.com/pt9912/ai-harness-course/tree/v6.5.0)), über
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

### 3.7 Ein Kommentar beschreibt, was da ist

Regeln dieser Sektion: Baseline-Regelwerk `grundlagen-harness-dateien.md` §Was ein Kommentar
trägt. Gilt für Code, Konfiguration, Skripte — **und für Zustandsfelder**.

Ein Kommentar trägt genau eine dieser Klassen — **Zusage · Kopplung · Abgrenzung · Rang-Zeiger ·
Grenze** — und schreibt an den, der die Stelle *ändert*, nicht an den, der die Entscheidung
*trifft*.

**Falsch:** Konjunktiv über die verworfene Alternative („ohne dieses Feld behauptete die Ausgabe
eine Verteilung, die nicht stattgefunden hat").
**Richtig:** Indikativ über den Zustand („verteilt ist wahr, wenn die Splitting-Regel angewendet
werden konnte").

**Falsch:** abwesenden Text beschreiben („die frühere Fassung prüfte nur die Länge").
**Richtig:** die geltende Zusage nennen; die vorige hält `git`.

**Zustandsfelder ebenso.** Eine `Stand`-/`Status`-Zelle in Roadmap, Beobachtungs-Register oder
Meilenstein-Tabelle nennt den Zustand und den Beleg als auflösbaren Anker, nicht die Chronik; das
Drift-Log der Roadmap trägt nur Umplanungen, keine Schließungen und keine erreichten Meilensteine.

**Begründung:** Die Abwägung gehört in die ADR, die Historie in `git`, die Herkunft in **ein**
auflösbares Feld. Was daneben steht, liest jeder Lauf mit und bezahlt es mit Kontext.

**Durchsetzung:** keine — die Regel ist **inferentiell**: „ist dieser Satz Chronik?" ist ein
Urteil, kein Match. Sie hängt am Review, nicht an einem Lauf. Diese Grenze ist benannt, nicht
verschwiegen (`modul-13`: einen Sensor zu behaupten, wo keiner steht, ist selbst eine
Harness-Lüge).

## 4. Quality Gates

Regeln dieser Sektion: Nur Targets aufzählen, die im Makefile **existieren** — halluzinierte
Gates sind die häufigste Form von Harness-Lüge (Baseline-Regelwerk `modul-13-quality-gates.md`).

Nur hier gelistete Targets existieren im Makefile. Halluzinierte Gates
sind die häufigste Form von Harness-Lüge; `make gate-consistency` erzwingt
die Übereinstimmung Doku ↔ Makefile mechanisch. Die Code-Gates sind
Dockerfile-Stages, die Meta-Gates laufen als Host-Bash. **Mandatory** ist, was in einem der
beiden Aggregate hängt: `gates` (Code-Fragen) oder `verify` (DoD-/Closure-Fragen). Von den
`doc-*`-Targets sind das `doc-check`, `doc-targets`, `doc-planning`, `doc-workflows` und
`doc-immutable` (in
`gates`) sowie `doc-structure` und `doc-complete` (in `verify`); die übrigen sind **advisory** —
`d-check`-Funktionen, die man aufruft, wenn man sie braucht. Ob ein Gate gerade grün ist, sagt
die CI (Badge im [`README.md`](README.md)), nicht diese Tabelle.

| Target | Zweck |
|---|---|
| [`make doc-check`](harness/sensors/doc-check.md) | Links, Anker, Kennungs-Linkpflicht und Referenzmatrix der Repo-Doku lösen auf; dazu die Lifecycle-Invariante wandernder Slices (`links.resolve-from`) und die Versions-Kohärenz (`versions`). **Vier benannte Grenzen** — Symlinks, die `current-from`-Trägerdatei, die Zeitdokument-Klassen und Digests — stehen in der Sensor-Datei, nicht hier |
| `make doc-trace` | advisory Requirements Traceability Matrix via `d-check` (DC-FA-CLI-009; `TRACE_FLAGS=--json`) |
| `make doc-complete` | Vollständigkeits-Gate: eine Anforderung ohne referenzierenden Slice ⇒ Exit 1 (DC-FA-CLI-011). Seit slice-123 **im `verify`-Aggregat** — davor advisory und damit nie gelaufen; eine Waise fiel erst auf, als jemand das Target von Hand aufrief |
| `make doc-doctor` | erklärende Diagnose mit Fix-Kandidaten (DC-FA-CLI-007) — **advisory** |
| `make doc-repair` | Reparatur-Patch (unified diff) auf stdout, git-apply-rein (DC-FA-CLI-008) |
| `make doc-immutable` | ADR-Immutabilität (§3.5) via git-Diff (Modul `vcs`; `RANGE=`/`STAGED=1`, DC-FA-VCS-001) — **CI-durchgesetzt** über die Commit-Range ([`ci.yml`](.github/workflows/ci.yml)) |
| `make doc-commits` | Commit-Message-Traceability (Modul `commits`; `RANGE=`, DC-FA-COMMITS-001) |
| [`make doc-planning`](harness/sensors/doc-planning.md) | Lifecycle-Konsistenz Roadmap ↔ `in-progress/`: liegt dort ein Slice, benennt ihn die Roadmap-Sektion, statt den Ruhe-Marker zu tragen |
| [`make doc-workflows`](harness/sensors/doc-workflows.md) | Deklarations-Form der `uses:`-Referenzen unter `.github/workflows`: voller SHA plus Tag-Kommentar beim Fremden, existierendes Ziel und gedeckte Rechte beim Lokalen. Prüft die **Form**, nicht die Gültigkeit |
| [`make doc-reviews`](harness/sensors/doc-reviews.md) | Review-Report-Deckung: eine `done/`-Slice-DoD-Zeile mit der Phrase „unabhängiger Review" braucht einen Report gleicher Kennung. **Opt-in pro Slice über die Phrase selbst** |
| `make doc-tracked` | Getrackt-Status auflösbarer Referenz-Ziele (Modul `tracked`, DC-FA-TRK-001) |
| `make doc-targets` | Deklarations-Konsistenz Doku ↔ Build-Targets (Modul `targets`, DC-FA-TGT-001), konfiguriert in [`.d-check.yml`](.d-check.yml) seit slice-074. **Im `gates`-Aggregat seit slice-079** — es hat dort `gate-consistency` (1)+(2) abgelöst, deren Parität in beiden Richtungen gemessen ist (slice-073/079) |
| [`make doc-structure`](harness/sensors/doc-structure.md) | Struktur-Invarianten innerhalb der Dokumente, **fünf Regeln**: Größen-Regel, Closure-Struktur, Lerneintrag-Form, Kopffelder, AC-Form |
| `make doc-usage` | Aufruf und Optionen von d-check selbst (`--help`) — **advisory**, seit dem Pin auf `v0.74.1` von `d-check --print-mk` mit erzeugt |
| `make doc-help` | Liste der `doc-*`-Targets (Utility) |
| `make lint` | golangci-lint mit dem Projekt-Profil (§3.2, [ADR-0005](docs/plan/adr/0005-lint-profil.md)) |
| `make test` | Akzeptanzkriterien der `AC-FA-*` als Go-Tests |
| `make coverage-gate` | Gesamt-Coverage ≥ 90 % über `./internal/...` ([ADR-0006](docs/plan/adr/0006-coverage-gate.md)) |
| `make arch-check` | Eigen-Architektur via `a-check` selbst (Dogfooding) |
| `make gate-consistency` | Meta-Gate: `.d-check.yml`-Module (Harness-Lügen-Schutz) + Pin-Konsistenz (Digest-Gleichheit harte Pins == `version.md#aktuell`, Version == CHANGELOG, `d-check.mk`-Deklaration; slice-018) + ADR-Index-Vollständigkeit (jede ADR-Datei ist im Index verlinkt; slice-087) |
| [`make version-coherence`](harness/sensors/version-coherence.md) | Kohärenz **doppelt deklarierter** Versions-Angaben: ein `uses:`-SHA ⇒ ein Tag-Kommentar; eine Variable in `Makefile` **und** `Dockerfile` ⇒ ein Wert. Prüft **Divergenz, nicht Wahrheit** |
| `make record-gates` | Gate-Nachweis (Working-Tree-Hash) für den Stop-Hook |
| `make suppression-check` | Fitness Function zum Suppression-Verbot (§3.2, [ADR-0005](docs/plan/adr/0005-lint-profil.md)): keine `//nolint`-Direktive in den Go-Quellen — `nolintlint` prüft nur Wohlgeformtheit, nicht Existenz (slice-049) |
| [`make symlink-check`](harness/sensors/symlink-check.md) | Zwei Prüfungen je getracktem Symlink: das Ziel existiert, und ein Ziel unter `.harness/baseline/` trägt den adoptierten Stand |
| [`make dcheck-phrase-selftest`](harness/sensors/dcheck-phrase-selftest.md) | Kalibrierung der phrasen-basierten Modul-Konfigurationen, **beide Hälften**: vier Werkzeug-Kontrollen gegen eigene Fixtures (reagiert `d-check` noch?) und eine **Korpus**-Kontrolle gegen den `done/`-Bestand (trägt er die Phrase noch?) |
| `make guard-selftest` | Selbsttest des PreToolUse-Command-Guard (Tool-Call-Gate §3.1) |
| [`make ci-range-selftest`](harness/sensors/ci-range-selftest.md) | Selbsttest der Commit-Range-Weiche der CI: vier Fälle, darunter der **Force-Push**, bei dem eine im Runner-Klon unerreichbare Basis auf den Default-Branch fällt statt abzubrechen |
| [`make image-scan`](harness/sensors/image-scan.md) | **kein Bestandteil von `gates`** (nicht hermetisch) — CVE-Scan gegen das **publizierte** Image; über rot entscheiden nur **behebbare** CRITICAL/HIGH. Exit-Codes und Sperren: siehe Sensor-Datei |
| [`make regelwerk-check`](harness/sensors/regelwerk-check.md) | **kein Gate** — misst die Integrität der vendored Baseline gegen `SHA256SUMS`, fail-closed. Die Freshness-Hälfte bleibt als Netz-Operation ungeprüft |
| `make slice-mv` | **kein Gate** — ein Werkzeug: Lifecycle-Wechsel eines Slice per `git mv` **samt der Verweise auf ihn**, repo-weit und in beiden im Bestand vorkommenden Formen (`SLICE=<slice-NNN> TO=<open\|next\|in-progress\|done>`). Antwort auf [`BEO-008`](docs/plan/planning/observations/BEO-PLAN/verweis-auf-wandernden-slice/observation.md) bei 3× (slice-118); die Gegenrichtung — Verweise **in** wandernden Dateien — trägt `doc-check` |
| `make gates` | alle inneren Gates (mandatory vor Handoff) |
| [`make verify-risiko-ausgaenge`](harness/sensors/verify-risiko-ausgaenge.md) | Jedes in §6 **notierte** Risiko trägt genau einen Ausgang aus der geschlossenen Dreier-Menge. Geprüft in `done/` **und** in `in-progress/`, sobald die Closure-Notiz ausgefüllt ist |
| [`make verify-observations`](harness/sensors/verify-observations.md) | Deckung des Beobachtungs-Registers: jeder in `done/` zitierte Pfad hat ein Verzeichnis, jedes Verzeichnis ein nicht leeres `evidence/`. Der Zähler wird abgeleitet, nicht geführt |
| `make verify` | **Verifikations-Schicht** (getrennt von `gates`, Regelwerk Modul 11): beantwortet DoD-/Closure-Fragen statt Code-Fragen; vor der „fertig"-Meldung auszuführen |
| `make image-test` | [AC-FA-DIST-001](spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk) + nativ==Container-Akzeptanz + Fragment-Parität (committete [`a-check.mk`](a-check.mk) == `--print-mk`, slice-034) gegen das gebaute Image |
| `make ci` | CI-äquivalent: `gates` + `image-test` (Workflow `.github/workflows/ci.yml`) |
| `make trace-check` | Traceability via Modul `commits` ([ADR-0021](docs/plan/adr/0021-commits-modul-trace-check.md)): `AC-*`/`ADR-*`/`MR-*`/`slice`-ID je Commit (§5; `MSGFILE=` Hook, `RANGE=` CI) |
| `make commit-scope-check` | Commit-Scope `(planning)` berührt nur `docs/plan/planning/` (§5, [`SL-003`](docs/plan/planning/observations/README.md)); misst jeden Commit an der zu seinem Zeitpunkt geltenden Fassung (`RANGE=` wie `trace-check`, slice-062) |
| `make archive-wave-test` | Testsuite von `tools/archive-wave/` (eigenes `go.mod` — **nicht** Teil von `make test`, das nur das Hauptmodul deckt) |
| [`make archive-wave`](harness/sensors/archive-wave.md) | **kein Gate** — bewegt Zeitdokumente ins Archiv und ersetzt Volltexte durch Stubs (`WELLE=`/`SLICE=`, `APPLY=1`). Seit slice-157 Pflichtschritt beim Abschluss eines wellenlosen Slice (§6). Die `WELLE=`-Falle und die Stub-Grenze stehen in der Sensor-Datei |

## 5. Dokumentations-Regeln

- Commits/PRs müssen mindestens eine `AC-*`- oder `ADR-*`-ID nennen
  (auch `MR-*`/`slice-NNN` gelten). Durchgesetzt durch `make trace-check`
  (slice-006; via d-check-Modul `commits`, [ADR-0021](docs/plan/adr/0021-commits-modul-trace-check.md))
  — lokal über `HEAD~1..HEAD`, in der CI über den Commit-Range
  ([`.github/workflows/ci.yml`](.github/workflows/ci.yml)). IDs werden nur
  beim Spec-/ADR-Schreiben nach dem deklarierten Schema vergeben (siehe
  [`harness/conventions.md`](harness/conventions.md)) — nie ad hoc im
  Commit/PR; Agenten referenzieren IDs, sie erfinden keine.
- **Commit-Scope `(planning)`:** ein Commit mit diesem Scope (`docs(planning)`,
  `fix(planning)`, `chore(planning)`) berührt **ausschließlich**
  `docs/plan/planning/`. Wandert Substanz eines anderen Bereichs mit, ist das ein
  eigener Commit mit passendem Scope. Durchgesetzt durch `make commit-scope-check`
  (slice-062); jeder Commit wird an der Fassung gemessen, die zu **seinem**
  Zeitpunkt galt, ältere sind damit grandfathered.
  **Warum nur dieser Scope:** über die gesamte Historie ist die Regel hier
  rauschfrei — fünf Treffer bei 74 Commits, alle fünf echte Diskrepanzen
  ([`SL-003`](docs/plan/planning/observations/README.md)). Für `docs(...)` allgemein wären es 31
  bei 193, weil `docs(spec)` legitim `spec/` und `docs(adr)` legitim ADRs ändert;
  eine Regel, die den Bestand massenhaft bricht, wird abgeschaltet statt befolgt.
  Ein weiterer Scope wird erst geregelt, wenn er auffällt — und dann gemessen,
  nicht geraten.
- **Wer eine Anforderung anlegt, nennt ihre Kennung in der Closure-Notiz.** Im **Plan** kann er es
  nicht: IDs werden referenziert statt erfunden, die neue Kennung existiert dort noch nicht, und
  jede genannte ist linkpflichtig — ein Link ins Leere macht `doc-check` rot. Also umschreibt der
  Plan sie („eine neue `AC-FA-CLI`-Kennung"), und die Requirements-Matrix sieht den Slice **nicht**.
  Bei der Closure ist die Anforderung geschrieben; dort steht die Kennung mit Link. Durchgesetzt
  durch `make doc-complete` im `verify`-Aggregat (slice-123) — eine Anforderung ohne
  referenzierenden Slice ist ab dann abschluss-blockierend.
- Neue oder geänderte `AC-*`-Anforderungen entstehen nur in
  [`spec/lastenheft.md`](spec/lastenheft.md) — nie per ADR (ADRs schärfen
  die Spezifikation, nicht das Lastenheft).
- Neue ADRs müssen den ADR-Index aktualisieren.
- Roadmap/Status-Geschichte lebt in `docs/plan/planning/`, nicht in der
  Architektur-Spec.
- **Slice-Lifecycle** ist reine Datei-Bewegung (`make slice-mv`, das den `git mv` samt der Verweise **auf** die Datei fährt; siehe §3.3) — der
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
- **WIP-Limit = 1.** Es liegt **höchstens ein** Slice in `in-progress/` (die Roadmap
  zählt nicht mit). Das ist eine harte Obergrenze, kein Vorschlag: zwei aktive Slices
  teilen sich einen Gate-Nachweis und eine Closure-Aufmerksamkeit, und beides
  trägt nur einmal. **Null ist zulässig** — nach jedem Abschluss der Normalfall,
  bis der nächste Slice gezogen wird. Bis slice-077 stand hier „genau ein"; das
  machte aus dem Baseline-Maximum zusätzlich ein Minimum und erklärte den
  regulären Leerlauf zum Regelverstoß.
- **AC-Form:** die Pflicht-Bausteine einer Anforderung stehen in
  [`harness/conventions.md`](harness/conventions.md) §Anforderungs-Anlege-Prozess
  — dort seit jeher die drei Pfade (Happy/Boundary/Negative im
  Given/When/Then-Stil) plus Out-of-Scope. Neu ist nur die **Durchsetzung**:
  `make verify` prüft sie ab slice-054 für **neue** `AC-*`; die **19** bei
  Einführung bestehenden sind **grandfathered** (vertraglich bindend, Rand- und
  Negativfälle bereits in Prosa — ein Umbau träfe die Form statt der Substanz),
  und die Grandfather-Liste wächst nicht mit.
- **Diskrepanz-Trichter:** eine Ausnahme von einer Regel oder einem Gate wird **nicht** ad hoc
  gesetzt, sondern über zwei Fragen eingeordnet — Granularität **vor** Temporalität: Cluster im
  selben Geltungsbereich ⇒ BF-Sub-Area-Markierung in
  [`harness/conventions.md`](harness/conventions.md#modus-deklaration-pro-sub-area); einzelne
  Diskrepanz mit erreichbarem Trigger ⇒ **Carveout** unter
  [`docs/plan/carveouts/`](docs/plan/carveouts/README.md); Trigger nie erreichbar ⇒ permanente
  ADR. Bootstrap-aware Gates gehören in keine der drei Klassen — sie stufen die Prüfung, sie
  nehmen keine Diskrepanz aus (slice-065).
- **Beobachtungs-Register:** der Zähler des Steering Loops liegt als **stehende** Ablage
  [`docs/plan/planning/observations/`](docs/plan/planning/observations/README.md) — nicht je Welle, weil
  eine übernommene Sektion an einer ungebrochenen Kette hinge und eine vergessene Übernahme den
  Zähler auf null setzte. Je Beobachtung ein Verzeichnis `BEO-<KUERZEL>/<slug>/` (Kürzel aus
  [`harness/conventions.md`](harness/conventions.md#modus-deklaration-pro-sub-area), seit
  slice-139 — davor Tabellenform, `BEO-NNN`, seit slice-101). **Eingetragen wird bei der
  Slice-Closure**: neues Verzeichnis mit `observation.md`, oder eine weitere Datei in ein
  vorhandenes `evidence/`. Der Zähler wird **abgeleitet** — er ist die Zahl der Evidence-Dateien;
  es gibt kein Feld, das man erhöht.
  Er ist zugleich der dritte Ausgang, den jedes offene Risiko einer Closure nimmt: *eingetreten* ⇒
  Carveout oder Folge-Slice · *entfallen* ⇒ gestrichen **mit Begründung** · *weiter offen* ⇒
  Register (slice-101).
- **Steering-Loop:** wiederkehrende Fehlermuster werden in
  [`docs/plan/steering-loop.md`](docs/plan/planning/observations/README.md) gezählt. Ab dem
  **zweiten** gleichartigen Vorfall entsteht ein Eintrag, ab dem **dritten** ist
  es eine Harness-Lücke und verlangt einen Guide oder Sensor — „besser
  aufpassen" ist keine Antwort. Ein Eintrag ohne Vorfallszahl ist unzulässig:
  die Zahl ist das Einzige, was die Schwelle prüfbar macht.
- **Slice-Form:** neue Slices entstehen aus der **vendored Ziel-Form**
  [`.harness/baseline/v6.5.0/templates/docs/plan/planning/slice.template.md`](.harness/baseline/v6.5.0/templates/docs/plan/planning/slice.template.md) — a-check führt keine eigene Kopie, sie würde gegen die Baseline driften.
  **Beim Kopieren anzupassen** — sieben Punkte, jeder gegen den Bestand gemessen (slice-178):

  1. Die Zeile `Lerneintrag — Form: <…>` **ergänzen** — die Ziel-Form kennt sie nicht als Feld,
     `make verify` verlangt sie.
  2. **Ein** Feld streichen: das *Reconciliation-Register* — a-check hat keinen
     Brownfield-Bootstrap, `reconciliation.md` existiert nicht. **Alles andere bleibt**, auch was
     frühere Fassungen dieser Liste zum Streichen empfahlen: `**Welle:**` führen **174** Dateien,
     den **Herkunfts-Anker** (`— liegt in <Zielort>` plus `seit slice-<NNN>` dort) führt das Repo
     an **58** Stellen, und die *drei Paarungen* stehen in jeder Closure-Notiz.
  3. **Wer die drei Paarungen trägt, entscheidet der Repo-Zustand — nicht das `Welle:`-Feld des
     Slice.** Baseline `modul-06-roadmap.md` §Wann Arbeit eine Welle braucht: *„Wellenlos ist eine
     Eigenschaft des **Repos** … das Kopf-Feld `**Welle:**` sagt nur, ob dieser Slice in ein Bündel
     gehört; daraus folgt für die Vorgänge unten nichts."* Liegt eine **offene Welle** vor, trägt
     ihre Closure die Paarungen — **auch für Slices ohne Wellen-Zugehörigkeit**. Erst ohne
     Wellen-Betrieb trägt die Slice-Closure sie selbst.
  4. **Die Abschnitts-Nummern der Ziel-Form gelten nicht unverändert.** a-check schiebt zwischen
     Ziel und DoD einen **Analyse-Abschnitt** ein (er hält die Messung, auf der die DoD steht) und
     stellt die DoD hinter den Umsetzungs-Abschnitt. Wieviel sich dadurch verschiebt, hängt vom
     Slice ab — der Bestand führt den Sub-Area-Abschnitt als §7 bis §10. **Nicht die Nummer
     kopieren, sondern die Reihenfolge der Ziel-Form einhalten und den Analyse-Abschnitt
     einfügen, wo er gebraucht wird.**
  5. **§1 heißt *Ziel und Abgrenzung*** und trägt beides: das Ziel in einem Satz und die
     Ausschlüsse **je Punkt mit Begründung** — *ein Ausschluss ohne Grund ist eine Behauptung,
     keine Grenze*. **Keine Mindestzahl**; die vier Klassen sind ein **Suchraster**, keine
     Ausfüll-Liste: Was übernimmt ein **Folge-Slice** (mit Kennung, und die Kennung muss den Punkt
     annehmen)? Was bleibt als **Bestand** bewusst stehen? Was wäre ein **anderer Vorgang**?
     Welche **Schicht** rührt der Slice nicht an? Der Abschnitt ist die Grenze, an der ein
     wachsender Slice sich messen lässt: Wer später etwas mitnimmt, das hier ausgeschlossen war,
     hat den Plan **geändert**, nicht ergänzt.
  6. **Der Sub-Area-Abschnitt heißt *Sub-Area-Prüfungen und Modus-Begründung*** — der Titel trägt beide Hälften, weil
     nur die zweite bedingt ist. Die zwei *Vorgelagert*-Blöcke (Sub-Area-Wahl prüfen · offene
     Beobachtungen sichten) laufen in **jedem** Slice-Plan, unabhängig von Modus und Slice-Typ;
     **der Abschnitt entfällt nie**.
  7. Die Ziel-Form führt eine **Review-DoD-Zeile**
     ([`MR-019`](harness/conventions.md#mr-019)) — ihr Wortlaut wird beim Kopieren auf die exakte
     Trigger-Phrase „unabhängiger Review" umgeschrieben, statt den Baseline-Wortlaut unverändert
     zu übernehmen; sonst prüft `make doc-reviews` sie nie (empirisch geprüft, slice-165 §3/§4).

  Die Regel trägt die
  Größen-Regel — **höchstens drei Liefer-Punkte und höchstens zwei Schichten**; passt der Slice
  nicht hinein, wird er **zerlegt, nicht gedehnt** — und verlangt den Lerneintrag in einer von drei
  **benannten** Formen (geschärfte Regel · neuer Sensor · benannte Spec-Lücke). `make verify`
  prüft beides ab slice-052; ältere Slices sind grandfathered.
  **Gezählt wird nur, was mit dem Umfang wächst.** Gate-Läufe, Review-Report, Closure-Notiz,
  Register und Risiko-Ausgänge zählen **nicht** — sie sind pro Slice konstant (Baseline `modul-05`
  §Ziel-Form: Slice, Review-Report seit `v6.2.0`). Der Gate-Lauf steht darum als feste Zeile unter dem DoD; als Checkbox ist er
  ab slice-098 ein Befund. Ab demselben Stichtag trägt der Kopf `Verantwortlich:`, `Autor:` und
  die berührten Spec-Stellen — `—` ist eine gültige Antwort, Schweigen nicht.
- **Zitier-Form in einfrierenden Artefakten** (`v6.5.0`, vier Ziel-Formen:
  Review-Report, Welle-Ergebnisnotiz, beide Archiv-Stubs): Was einfriert,
  zitiert **Kennung statt Adresse** — `slice-NNN` statt seines Lifecycle-Pfads,
  `make <target>` statt eines Links auf die Sensor-Datei, eine Baseline-Stelle
  als **Tag + Pfad in Inline-Code** statt als Link
  (`` `v<X.Y.Z>` · `regelwerk/<datei>.md` §<Abschnitt> ``). Grund: Der vendored
  Baum trägt genau einen Tag, der nächste Sprung löscht den alten, und ein Link
  darauf färbt ein Artefakt rot, das niemand mehr anfassen darf. **Gemessen
  statt behauptet** (slice-176): Nach der Umstellung von 16 Links ließ sich der
  vorige Stand entfernen, ohne ein einziges eingefrorenes Artefakt anzufassen —
  slice-172 hatte an derselben Stelle 22 Nachzüge gebraucht. Verankert im
  Reviewer-Skill; für Archiv-Stubs erzeugt `tools/archive-wave/` den Text, für
  die Ergebnisnotiz gilt sie beim Schreiben.
  **Nicht** betroffen: lebende Dokumente — dort ist der Link richtig, und
  `versions` hält ihn aktuell.
- **Geltungsbereich einer Messung** (`seit slice-179`, Lese-Schritt der
  welle-15-Closure): Wer eine Messung als **Beleg** schreibt — in einem
  Slice-Plan, einer Closure-Notiz, einem Review-Report —, nennt ihren
  **Geltungsbereich** und sagt, ob er den Gegenstand deckt. Nicht *„22 Befunde,
  keine weitere Klasse"*, sondern *„22 Befunde über Markdown-Links; Prosa sieht
  das Instrument nicht"*.
  **Anlass:** [`BEO-PLAN/review-geltungsbereich-zu-eng`](docs/plan/planning/observations/BEO-PLAN/review-geltungsbereich-zu-eng/observation.md)
  bei 3× — dreimal war die Begrenzung begründet und trotzdem zu eng, und
  dreimal fand es jemand anderes als der Messende. **Kein Sensor:** ob ein
  Geltungsbereich weit genug ist, ist ein Urteil über eine Absicht (§3.7); ein
  zweites Muster, das nach übersehenen Klassen sucht, kann dieselbe Verengung
  haben wie das erste. Was greift, ist die Frage beim **Schreiben** — sie kostet
  einen Halbsatz und hätte alle drei Fälle gefangen.
- **CR-Texte an ein fremdes Werkzeug** (bisher vier an `d-check`) leben im Slice, der sie erzeugt,
  und gehen erst nach einem Prüf-Durchgang hinaus: der Skill
  [`.harness/skills/cr-text-reviewer.md`](.harness/skills/cr-text-reviewer.md) markiert jeden Satz,
  der eine **Tatsache** über ein System behauptet — das eigene oder das fremde —, und nennt den
  Handgriff, der ihn belegt. Anlass ist [`BEO-022`](docs/plan/planning/observations/BEO-GATE/cr-text-behauptet-statt-gemessen/observation.md) bei **3×**:
  dreimal stand eine Behauptung als Annahme da, die eine Messung zugelassen hätte. Die Prüf-Frage
  ist **nicht** „hast du gemessen?", sondern „hast du *das* gemessen, worüber du redest?" — die
  zweite Ausprägung misst die eigene Menge und sagt über die fremde aus, sieht dabei aus wie ein
  Beleg und kann sogar zutreffen. **Kein Sensor:** ob ein Satz gemessen wurde, ist ein Urteil über
  seinen Entstehungsweg, kein Match (§3.7).
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
3. Betroffene IDs benennen: Slice-ID, `AC-*`, `ADR-*`, betroffene Module,
   auszuführende Gates.
4. Kleinste sinnvolle Änderung planen — **und die Plan-Ausgabe nennt
   Out-of-Scope.** Das ist die Schritt-Hälfte derselben Regel, deren
   Dokument-Hälfte §1 *Ziel und Abgrenzung* des Slice-Plans ist (§5). Der Lauf
   schreibt fort, was der Plan schon ausschließt; er erfindet die Abgrenzung
   nicht neu und **darf sie nicht stillschweigend weiten**: Nimmt der Lauf etwas
   mit, das §1 ausschließt, ist das eine **Plan-Änderung** und gehört vor den
   Code, nicht in den Bericht danach.
5. Engsten nützlichen Sensor laufen lassen.
6. Repo-weiten Gate-Lauf vor Handoff (`make gates`, sobald slice-003 ihn anlegt).
7. Doku/Indizes aktualisieren, falls ein öffentlicher Vertrag berührt.
8. Ausgeführte Sensors und verbleibende Risiken berichten — keine
   Erfolgsmeldung ohne Gate-Ausführung.

Dieser Workflow deckt ausschließlich die Implementer-Rolle ab. Schritt 8
ist der Rollenwechsel, kein Abschluss: Bericht → Handoff an Reviewer
([`.harness/skills/reviewer.md`](.harness/skills/reviewer.md), siehe
[`harness/README.md`](harness/README.md) §Guides) → Verifier. Kein
Self-Review — anderer Kontext findet andere Findings, derselbe Kontext
dieselben blinden Flecken (Baseline-Regelwerk `modul-08-agentenrollen.md`);
ein `fork`-Subagent erbt den Kontext und zählt darum **nicht** — ein
frischer Subagent-Typ ohne `fork`, briefed mit dem nötigen Kontext, zählt
(Modul 8 §Kontext-Trennung). Ausfallbeleg:
[`BEO-HARNESS/behauptete-vollstaendigkeit-extern-gefangen`](docs/plan/planning/observations/BEO-HARNESS/behauptete-vollstaendigkeit-extern-gefangen/observation.md).

Beim **Abschluss** eines Slice zusätzlich `make verify` (Verifikations-Schicht,
§4): `gates` beantwortet Code-Fragen, `verify` die DoD-/Closure-Fragen. Die
semantische Hälfte — trägt die Notiz ein Lernsignal oder nur eine Floskel? —
leistet der Skill
[`.harness/skills/closure-note-reviewer.md`](.harness/skills/closure-note-reviewer.md).

**Danach, wenn der Slice wellenlos ist** (kein `**Welle:**`-Feld, keine aktive
Welle in `in-progress/`): sofort archivieren —
`make archive-wave SLICE=<slice-id> APPLY=1`, danach `make gates`/`make
verify` auf dem archivierten Stand erneut grün, als **eigener Commit** direkt
im Anschluss an den Closure-Commit (git mv + Content-Rewrite ist ein eigener
Vorgang, §3.3). Kein Backlog, den erst ein späterer Sweep aufräumt —
Baseline-Regelwerk `modul-06-roadmap.md` §Wann Arbeit eine Welle braucht,
Tabelle *Träger im Repo ohne Wellen*: „Zeitdokumente archivieren … Träger:
Slice-Closure". Ein Slice mit echtem `**Welle:**`-Feld archiviert stattdessen
**mit seiner Welle** bei deren Closure (`WELLE=<id>`), nicht einzeln.
