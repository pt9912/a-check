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
[`.harness/baseline/v6.0.0/regelwerk/README.md`](.harness/baseline/v6.0.0/regelwerk/README.md)
ist der Index (17 Module + acht Grundlagen-Abschnitte, eine Datei je
Abschnitt); die Ziel-Formen daneben unter
[`templates/`](.harness/baseline/v6.0.0/templates/README.md). Integrität:
`.harness/baseline/v6.0.0/SHA256SUMS`.

Das vendored Regelwerk ist ein **didaktik-freier Extrakt** und trägt keine
eigene Normativität: bei Konflikt gilt der Kurs
([`v6.0.0`](https://github.com/pt9912/ai-harness-course/tree/v6.0.0)), über
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
| `make doc-check` | Doku-Links/Anker/Kennungen via `d-check` (Schwester-Tool, digest-gepinnt, netzlos, read-only); seit slice-080 zusätzlich die **Lifecycle-Invariante** ([`SL-002`](docs/plan/planning/observations/README.md)) über `links.resolve-from` — sie hat `verify-slice-links` abgelöst und liegt damit in der Gate- statt der Verifikations-Schicht; seit slice-133 zusätzlich die **Versions-Kohärenz zwischen Dokumenten** über `versions` — eine Prosa-Angabe über die Version eines *anderen* eigenen Dokuments wird gegen dessen Kopf gehalten (`version-stale`). **Nicht** Digests: der Erwartungswert kommt versions-förmig aus dem `current-from`-Span, ein `sha256:` bricht dort fail-closed ab — die Digest-Gleichheit bleibt bei `gate-consistency` |
| `make doc-trace` | advisory Requirements Traceability Matrix via `d-check` (DC-FA-CLI-009; `TRACE_FLAGS=--json`) |
| `make doc-complete` | Vollständigkeits-Gate: eine Anforderung ohne referenzierenden Slice ⇒ Exit 1 (DC-FA-CLI-011). Seit slice-123 **im `verify`-Aggregat** — davor advisory und damit nie gelaufen; eine Waise fiel erst auf, als jemand das Target von Hand aufrief |
| `make doc-doctor` | erklärende Diagnose mit Fix-Kandidaten (DC-FA-CLI-007) — **advisory** |
| `make doc-repair` | Reparatur-Patch (unified diff) auf stdout, git-apply-rein (DC-FA-CLI-008) |
| `make doc-immutable` | ADR-Immutabilität (§3.5) via git-Diff (Modul `vcs`; `RANGE=`/`STAGED=1`, DC-FA-VCS-001) — **CI-durchgesetzt** über die Commit-Range ([`ci.yml`](.github/workflows/ci.yml)) |
| `make doc-commits` | Commit-Message-Traceability (Modul `commits`; `RANGE=`, DC-FA-COMMITS-001) |
| `make doc-planning` | Planning-Lifecycle-Konsistenz Roadmap ↔ `in-progress` (Modul `planning`, DC-FA-PLAN-001): liegt ein Slice in `in-progress/`, muss die Roadmap-Sektion ihn **benennen** statt den Ruhe-Marker zu tragen. Seit slice-122 in [`.d-check.yml`](.d-check.yml) konfiguriert und **im `gates`-Aggregat** — davor lief es ohne Gegenstand und meldete grün ([`BEO-014`](docs/plan/planning/observations/BEO-GATE/ruhe-marker-ungewaechtert/observation.md)). Die zweite und dritte Modul-Fähigkeit (`closure:`, `waves:`) bleiben bewusst unkonfiguriert: die Closure-Struktur prüft `doc-structure` |
| `make doc-workflows` | Deklarations-Form der `uses:`-Referenzen unter `.github/workflows` (Modul `workflows`): eine **fremde** Referenz nennt einen vollen 40-stelligen SHA mit Tag-Kommentar dahinter, eine **lokale** (`./…`) braucht keinen Pin, dafür ein existierendes Ziel und einen Aufrufer-Job, der die verlangten Rechte führt. Konfiguriert und **im `gates`-Aggregat seit slice-130**; eigenes `doc-*`-Target, weil [`d-check.mk`](d-check.mk) für dieses Modul keines erzeugt. Geprüft wird die **Form**, nicht die **Gültigkeit**: ob der Tag-Kommentar zum SHA passt, sagt es nicht — das wäre Netz, und der Widerspruch aus [`BEO-026`](docs/plan/planning/observations/BEO-GATE/versionsangabe-neben-digest-ungeprueft/observation.md) bleibt damit ungedeckt |
| `make doc-reviews` | Review-Report-Deckung (Modul `reviews`, `DC-FA-RVW-001`): eine `done/`-Slice-DoD-Zeile mit der Phrase „unabhängiger Review" braucht mindestens einen Report unter [`docs/reviews/`](docs/reviews/README.md) mit derselben `slice-<NNN>`-Kennung im Dateinamen (1:N zulässig, Substring-Match, nicht rekursiv — archivierte Stubs tragen keine DoD mehr und fallen natürlich aus der Kandidatenmenge). **Opt-in pro Slice über die DoD-Phrase selbst** — ohne sie ist die Kandidatenmenge für diesen Slice leer, kein Fehlalarm. Konfiguriert in [`.d-check.yml`](.d-check.yml) und **im `gates`-Aggregat seit slice-160**; eigenes `doc-*`-Target aus demselben Grund wie `doc-workflows` |
| `make doc-tracked` | Getrackt-Status auflösbarer Referenz-Ziele (Modul `tracked`, DC-FA-TRK-001) |
| `make doc-targets` | Deklarations-Konsistenz Doku ↔ Build-Targets (Modul `targets`, DC-FA-TGT-001), konfiguriert in [`.d-check.yml`](.d-check.yml) seit slice-074. **Im `gates`-Aggregat seit slice-079** — es hat dort `gate-consistency` (1)+(2) abgelöst, deren Parität in beiden Richtungen gemessen ist (slice-073/079) |
| `make doc-structure` | Struktur-Invarianten innerhalb der Dokumente (Modul `structure`, DC-FA-STRUCT-001), **fünf Regeln**: Größen-Regel, Closure-Struktur, Lerneintrag-Form, Kopffelder, **AC-Form**. Konfiguriert in [`.d-check.yml`](.d-check.yml) und **im `verify`-Aggregat** — es hat alle vier Eigenbau-Sensoren aus slice-080 abgelöst (`verify-slice-form`, die strukturelle Hälfte von `verify-closure-notes` und seit slice-120 `verify-ac-form`), Parität je Befundklasse gemessen. Braucht den Pin `v0.69.0` (`tasks-ignore-pattern`, `exempt-section-pattern`, `exempt-expect-count`) |
| `make doc-usage` | Aufruf und Optionen von d-check selbst (`--help`) — **advisory**, seit dem Pin auf `v0.74.1` von `d-check --print-mk` mit erzeugt |
| `make doc-help` | Liste der `doc-*`-Targets (Utility) |
| `make lint` | golangci-lint mit dem Projekt-Profil (§3.2, [ADR-0005](docs/plan/adr/0005-lint-profil.md)) |
| `make test` | Akzeptanzkriterien der `AC-FA-*` als Go-Tests |
| `make coverage-gate` | Gesamt-Coverage ≥ 90 % über `./internal/...` ([ADR-0006](docs/plan/adr/0006-coverage-gate.md)) |
| `make arch-check` | Eigen-Architektur via `a-check` selbst (Dogfooding) |
| `make gate-consistency` | Meta-Gate: `.d-check.yml`-Module (Harness-Lügen-Schutz) + Pin-Konsistenz (Digest-Gleichheit harte Pins == `version.md#aktuell`, Version == CHANGELOG, `d-check.mk`-Deklaration; slice-018) + ADR-Index-Vollständigkeit (jede ADR-Datei ist im Index verlinkt; slice-087) |
| `make version-coherence` | Kohärenz **doppelt deklarierter** Versions-Angaben (slice-131, Antwort auf [`BEO-026`](docs/plan/planning/observations/BEO-GATE/versionsangabe-neben-digest-ungeprueft/observation.md) bei 3×): derselbe `uses:`-SHA trägt unter `.github/workflows/` überall denselben Tag-Kommentar, und eine Versions-Variable, die [`Makefile`](Makefile) **und** [`Dockerfile`](Dockerfile) führen, hat an beiden Orten denselben Wert. Geprüft wird **Divergenz, nicht Unwahrheit** — zwei übereinstimmend falsche Angaben bleiben grün; die Registry zu fragen wäre Netz, und `gates` ist hermetisch. Der Sensor erklärt **keine** Seite zur führenden |
| `make record-gates` | Gate-Nachweis (Working-Tree-Hash) für den Stop-Hook |
| `make suppression-check` | Fitness Function zum Suppression-Verbot (§3.2, [ADR-0005](docs/plan/adr/0005-lint-profil.md)): keine `//nolint`-Direktive in den Go-Quellen — `nolintlint` prüft nur Wohlgeformtheit, nicht Existenz (slice-049) |
| `make guard-selftest` | Selbsttest des PreToolUse-Command-Guard (Tool-Call-Gate §3.1) |
| `make ci-range-selftest` | Selbsttest der Commit-Range-Weiche der CI ([`tools/ci-commit-range.sh`](tools/ci-commit-range.sh), slice-134): vier Fälle — `pull_request`, neuer Branch, **Force-Push**, normaler Push. Der dritte war der Defekt: `github.event.before` trägt nach einem Rebase einen gültig aussehenden SHA, den der Runner-Klon nicht kennt (`fetch-depth: 0` holt keine verwaisten Objekte), und die Range-Prüfungen brachen mit *„Range-Basis nicht auflösbar"* ab. Der vierte Fall ist die Gegenrichtung: eine **brauchbare** Basis wird auch benutzt. **Nicht** geprüft: welchen Wert GitHub in das Feld schreibt — das sagt nur ein echter Lauf |
| `make image-scan` | **kein Bestandteil von `gates`** — CVE-Scan gegen das **publizierte** Image (`ghcr.io/pt9912/a-check:latest`) mit digest-gepinntem Trivy ([ADR-0037](docs/plan/adr/0037-cve-scan-gegen-das-publizierte-image.md)). Netz ist hier der **Zweck**: eine gepinnte Vuln-DB fände nur die CVEs von gestern. Über rot entscheiden **nur behebbare** CRITICAL/HIGH; der Vollbericht fällt nie. Skript-Exit 0/1/2 — über `make` sind 1 und 2 **nicht** unterscheidbar (make normalisiert auf 2), wer den Ausgang braucht, liest die Ausgabe. Selbsttest der Auswertung: `bash tools/image-scan.sh --selftest` (netzlos). Zeitgesteuert in [`image-scan.yml`](.github/workflows/image-scan.yml) |
| `make regelwerk-check` | **kein Gate** — Wartung der vendored Baseline ([MR-006](harness/conventions.md#mr-006--baseline-committet-vendored-statt-per-url-referenziert)): Integrität gegen `SHA256SUMS` fail-closed; die Freshness-Hälfte bleibt als Netz-Operation ausdrücklich ungeprüft |
| `make slice-mv` | **kein Gate** — ein Werkzeug: Lifecycle-Wechsel eines Slice per `git mv` **samt der Verweise auf ihn**, repo-weit und in beiden im Bestand vorkommenden Formen (`SLICE=<slice-NNN> TO=<open\|next\|in-progress\|done>`). Antwort auf [`BEO-008`](docs/plan/planning/observations/BEO-PLAN/verweis-auf-wandernden-slice/observation.md) bei 3× (slice-118); die Gegenrichtung — Verweise **in** wandernden Dateien — trägt `doc-check` |
| `make gates` | alle inneren Gates (mandatory vor Handoff) |
| `make verify-risiko-ausgaenge` | Jedes in §6 **notierte** Risiko trägt genau einen Ausgang aus der geschlossenen Dreier-Menge (§5, ab slice-102). Geprüft in `done/` **und** in `in-progress/`, sobald dort die Closure-Notiz **ausgefüllt** ist — der Auslöser ist ihr Zustand, nicht das Verzeichnis (slice-129); ein Slice mit Vorlagen-Platzhalter ist in Arbeit und wird sichtbar übersprungen. Bleibt lokal, weil die Prüfung §6 mit §7 vergleicht und `structure` abschnitts-**lokal** ist. **Nicht** geprüft: die Existenz eines Risiko-Blocks |
| `make verify-observations` | Deckung des Beobachtungs-Registers ([`observations/`](docs/plan/planning/observations/README.md), Verzeichnisform seit slice-139): jeder in `done/` zitierte Pfad (neue Form `BEO-<KUERZEL>/<slug>` oder alte Form `BEO-NNN` über die `Ehemals:`-Zeile) hat ein Verzeichnis, jedes Verzeichnis trägt ein nicht leeres `evidence/`. Der Zähler wird nicht mehr geführt, sondern ist die Zahl der Evidence-Dateien — die alte Zähler-Prüfung entfällt strukturell. **Nicht** geprüft: Lage und Existenz der Beleg-Datei (`modul-06`) und die Umkehrung „jedes Verzeichnis ist zitiert" — die meisten stehen unter der Schwelle (slice-102, Verzeichnisform slice-139) |
| `make verify` | **Verifikations-Schicht** (getrennt von `gates`, Regelwerk Modul 11): beantwortet DoD-/Closure-Fragen statt Code-Fragen; vor der „fertig"-Meldung auszuführen |
| `make image-test` | [AC-FA-DIST-001](spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk) + nativ==Container-Akzeptanz + Fragment-Parität (committete [`a-check.mk`](a-check.mk) == `--print-mk`, slice-034) gegen das gebaute Image |
| `make ci` | CI-äquivalent: `gates` + `image-test` (Workflow `.github/workflows/ci.yml`) |
| `make trace-check` | Traceability via Modul `commits` ([ADR-0021](docs/plan/adr/0021-commits-modul-trace-check.md)): `AC-*`/`ADR-*`/`MR-*`/`slice`-ID je Commit (§5; `MSGFILE=` Hook, `RANGE=` CI) |
| `make commit-scope-check` | Commit-Scope `(planning)` berührt nur `docs/plan/planning/` (§5, [`SL-003`](docs/plan/planning/observations/README.md)); misst jeden Commit an der zu seinem Zeitpunkt geltenden Fassung (`RANGE=` wie `trace-check`, slice-062) |
| `make archive-wave-test` | Testsuite von `tools/archive-wave/` (eigenes `go.mod` — **nicht** Teil von `make test`, das nur das Hauptmodul deckt) |
| `make archive-wave` | Setzt Baseline-Regelwerk `modul-06-roadmap.md` §Wellen-Closure-Prozedur Schritt 4 um (`WELLE=<id>` sammelt die Slices einer Welle über ihr `**Welle:**`-Feld und deren Review-Reports, baut `done/<welle-id>/archiv.zip`, ersetzt Volltexte durch Stubs, zieht Verweise nach; `SLICE=<id>`/`REVIEW=<datei>` für wellenlose Einzel-Fälle), `APPLY=1` optional. **`WELLE=` nennt die kurze numerische Form** (`welle-12`, nicht `welle-12-regelwerk-migration`) — das Werkzeug matcht nur die Ziffernfolge im `**Welle:**`-Feld, a-checks eigene Roadmap-IDs tragen einen beschreibenden Suffix, den es nicht kennt (slice-144, per Dry-Run gemessen). Unverändert aus `d-check` übernommen (eigenständig, `tools/archive-wave/README.md`); **kein Gate**
(kein `make gates`/`make ci`-Bestandteil), aber seit slice-157 **Pflichtschritt beim Abschluss
eines wellenlosen Slice** (§6) — kein von Hand ausgelöster Vorgang mehr, sondern Teil des
Workflows. Sicherer Default ohne `APPLY=1` (nichts wird geschrieben) |

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
  [`.harness/baseline/v6.0.0/templates/docs/plan/planning/slice.template.md`](.harness/baseline/v6.0.0/templates/docs/plan/planning/slice.template.md) — a-check führt keine eigene Kopie, sie würde gegen die Baseline driften.
  **Beim Kopieren anzupassen:** die Zeile `Lerneintrag — Form: <…>` ergänzen (die Ziel-Form kennt
  sie nicht als Feld, `make verify` verlangt sie) und die vier Felder streichen, die a-check nicht
  führt — `Welle:`, Reconciliation-Register, *drei Paarungen*, Herkunfts-Anker. Die Regel trägt die
  Größen-Regel — **höchstens drei Liefer-Punkte und höchstens zwei Schichten**; passt der Slice
  nicht hinein, wird er **zerlegt, nicht gedehnt** — und verlangt den Lerneintrag in einer von drei
  **benannten** Formen (geschärfte Regel · neuer Sensor · benannte Spec-Lücke). `make verify`
  prüft beides ab slice-052; ältere Slices sind grandfathered.
  **Gezählt wird nur, was mit dem Umfang wächst.** Gate-Läufe, Closure-Notiz, Register und
  Risiko-Ausgänge zählen **nicht** — sie sind pro Slice konstant (Baseline `modul-05`
  §Ziel-Form: Slice). Der Gate-Lauf steht darum als feste Zeile unter dem DoD; als Checkbox ist er
  ab slice-098 ein Befund. Ab demselben Stichtag trägt der Kopf `Verantwortlich:`, `Autor:` und
  die berührten Spec-Stellen — `—` ist eine gültige Antwort, Schweigen nicht.
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
4. Kleinste sinnvolle Änderung planen.
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
