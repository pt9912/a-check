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

**Breiterer Pflicht-Blick bleibt bei drei Anlässen** — dort genügt der eine
Abschnitt nicht: Bootstrap · jede Änderung an
[`harness/conventions.md`](harness/conventions.md) (Adaptionen `MR-<NNN>`,
Source Precedence, ID-Schema) · Drift-Audit gegen die Baseline
(`modul-02` §Freshness-Audit der vendored Baseline — darunter die Stichprobe
gegen den Bestand, die **auch bei aktuellem Pin** läuft; `make regelwerk-check`
deckt davon nur die Integritäts-Hälfte, siehe §4).

Die vendorten **Ziel-Formen** unter
[`templates/`](.harness/baseline/v6.5.0/templates/README.md) tragen **zwei
Rollen**: als **Referenz-Form**, auf die das Regelwerk mit `../templates/…`
verweist, und als **Vorlage, die beim Anlegen kopiert und ausgefüllt wird statt
frei formuliert** — für ADR, Slice, Welle, Carveout, Review-Report und
Beobachtung. a-check führt **keine eigenen Kopien** davon; was beim Kopieren
anzupassen ist, steht am Ort der jeweiligen Ablage (für Slices:
[`docs/plan/planning/README.md`](docs/plan/planning/README.md) §Beim Kopieren
der Slice-Ziel-Form).

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

Implementierungssprache ist **Go**:
ein statisches, sprach-agnostisches Binary, das *fremde* Quellen
text-heuristisch prüft. Es gilt: **kein Host-Go und keine
Host-Paketmanager** (`go`, `pip`, `npm`, `cargo`, `apt`, `brew`, …). Alle
Checks laufen über `make`; die Go-Toolchain läuft in Docker. Der Host
braucht nur `git`, GNU `make`, `bash` und Docker.

**Falsch:** `go build ./…`, `go test ./…`
**Richtig:** `make gates`

**Begründung:** Toolchain-Reproduzierbarkeit + Supply-Chain-Defense.

**Durchsetzung:** Ein PreToolUse-Command-Guard
(`.claude/hooks/pretooluse-command-guard.sh`) lehnt Host-Toolchain-
und Paketmanager-Aufrufe (`go`/`golangci-lint`/`pip`/`npm`/`cargo`/`apt`/`brew`/…)
**vor** der Ausführung fail-closed ab (Tool-Call-Gate der Durchsetzungsschicht);
`make gates` belegt ihn über `make guard-selftest`.

### 3.2 Suppression-Verbot

Inline-Suppressions sind verboten (`//nolint` o. Ä.). Ausnahmen leben
zentral in der Lint-Konfiguration mit Begründung.

### 3.3 git mv + Inhaltsänderung = zwei Commits

Datei verschoben **und** Inhalt umgeschrieben: zwei Commits, und der
Move-Commit bleibt rein (Git erkennt R-Rename). **Welcher zuerst kommt, sagt
der Vorgang:**

1. **Regelfall:** `git mv` als eigener Commit, dann den Inhalt umschreiben.
2. **Lifecycle-Übergang nach `done/`:** erst der Inhalt (DoD-Häkchen,
   Closure-Notiz), dann der reine `git mv` — die Notiz ist die **Bedingung**
   dafür, dass die Datei nach `done/` darf, nicht ihre Folge. Den Move fährt
   `make slice-mv`; die **Reihenfolge** entscheidet der Lauf.

**Begründung:** Sonst fällt die Rename-Detection unter die 50 %-Schwelle und
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
sind die häufigste Form von Harness-Lüge; `make doc-targets` erzwingt
die Übereinstimmung Doku ↔ Makefile mechanisch — über diese Tabelle **und**
[`harness/README.md`](harness/README.md) §Sensors. Die Code-Gates sind
Dockerfile-Stages, die Meta-Gates laufen als Host-Bash. **Mandatory** ist, was in einem der
beiden Aggregate hängt: `gates` (Code-Fragen) oder `verify` (DoD-/Closure-Fragen). Von den
`doc-*`-Targets sind das `doc-check`, `doc-targets`, `doc-planning`, `doc-workflows`,
`doc-reviews` und `doc-mentions` (in `gates`) sowie `doc-structure` und `doc-complete`
(in `verify`); die übrigen sind **advisory** —
`d-check`-Funktionen, die man aufruft, wenn man sie braucht. Ob ein Gate gerade grün ist, sagt
die CI (Badge im [`README.md`](README.md)), nicht diese Tabelle.

| Target | Zweck |
|---|---|
| [`make doc-check`](harness/sensors/doc-check.md) | Links, Anker, Kennungs-Linkpflicht und Referenzmatrix lösen auf; dazu Lifecycle-Invariante und Versions-Kohärenz |
| `make doc-trace` | advisory Requirements Traceability Matrix via `d-check` (DC-FA-CLI-009; `TRACE_FLAGS=--json`) |
| `make doc-complete` | Vollständigkeits-Gate: eine Anforderung ohne referenzierenden Slice ⇒ Exit 1 (DC-FA-CLI-011); im `verify`-Aggregat |
| `make doc-doctor` | erklärende Diagnose mit Fix-Kandidaten (DC-FA-CLI-007) — **advisory** |
| `make doc-repair` | Reparatur-Patch (unified diff) auf stdout, git-apply-rein (DC-FA-CLI-008) |
| `make doc-immutable` | ADR-Immutabilität (§3.5) via git-Diff (Modul `vcs`; `RANGE=`/`STAGED=1`) — CI-durchgesetzt über die Commit-Range |
| `make doc-commits` | Commit-Message-Traceability (Modul `commits`; `RANGE=`, DC-FA-COMMITS-001) |
| [`make doc-planning`](harness/sensors/doc-planning.md) | Äquivalenz Roadmap ↔ `in-progress/`: liegt dort ein Slice, fehlt der Ruhe-Marker; ist es leer, steht er. Kein Name wird geprüft |
| [`make doc-workflows`](harness/sensors/doc-workflows.md) | Deklarations-Form der `uses:`-Referenzen unter `.github/workflows`. Prüft die **Form**, nicht die Gültigkeit |
| [`make doc-mentions`](harness/sensors/doc-mentions.md) | Erwähnungs-Deckung, die **Gegenrichtung** des Link-Checks: jede Datei unter `harness/sensors/` ist hier genannt |
| [`make doc-reviews`](harness/sensors/doc-reviews.md) | Review-Report-Deckung für `done/`-Slices mit der DoD-Phrase; **Opt-in pro Slice über die Phrase selbst** |
| `make doc-tracked` | Getrackt-Status auflösbarer Referenz-Ziele (Modul `tracked`, DC-FA-TRK-001) |
| `make doc-targets` | Deklarations-Konsistenz Doku ↔ Build-Targets (Modul `targets`, DC-FA-TGT-001), konfiguriert in [`.d-check.yml`](.d-check.yml); im `gates`-Aggregat |
| [`make doc-structure`](harness/sensors/doc-structure.md) | Struktur-Invarianten innerhalb der Dokumente: Größen-Regel, Closure-Struktur, Lerneintrag-Form, Kopffelder, AC-Form, Zellengrenzen der Gate-Tabellen |
| `make doc-usage` | Aufruf und Optionen von d-check selbst (`--help`) — **advisory**, seit dem Pin auf `v0.75.0` von `d-check --print-mk` mit erzeugt |
| `make doc-help` | Liste der `doc-*`-Targets (Utility) |
| `make lint` | golangci-lint mit dem Projekt-Profil (§3.2, [ADR-0005](docs/plan/adr/0005-lint-profil.md)) |
| `make test` | Akzeptanzkriterien der `AC-FA-*` als Go-Tests |
| `make coverage-gate` | Gesamt-Coverage ≥ 90 % über `./internal/...` ([ADR-0006](docs/plan/adr/0006-coverage-gate.md)) |
| `make arch-check` | Eigen-Architektur via `a-check` selbst (Dogfooding) |
| [`make gate-consistency`](harness/sensors/gate-consistency.md) | Meta-Gate, drei Prüfungen: `.d-check.yml`-Module, Pin-Konsistenz, ADR-Index-Vollständigkeit |
| [`make version-coherence`](harness/sensors/version-coherence.md) | Kohärenz **doppelt deklarierter** Versions-Angaben. Prüft **Divergenz, nicht Wahrheit** |
| `make record-gates` | Gate-Nachweis (Working-Tree-Hash) für den Stop-Hook |
| `make suppression-check` | Fitness Function zum Suppression-Verbot (§3.2): keine `//nolint`-Direktive in den Go-Quellen |
| [`make symlink-check`](harness/sensors/symlink-check.md) | Zwei Prüfungen je getracktem Symlink: das Ziel existiert, und ein Ziel unter `.harness/baseline/` trägt den adoptierten Stand |
| [`make dcheck-phrase-selftest`](harness/sensors/dcheck-phrase-selftest.md) | Kalibrierung der phrasen-basierten Modul-Konfigurationen, **beide Hälften**: Werkzeug und Korpus |
| `make guard-selftest` | Selbsttest des PreToolUse-Command-Guard (Tool-Call-Gate §3.1) |
| [`make ci-range-selftest`](harness/sensors/ci-range-selftest.md) | Selbsttest der Commit-Range-Weiche der CI: vier Fälle, darunter der **Force-Push** |
| [`make image-scan`](harness/sensors/image-scan.md) | **kein Bestandteil von `gates`** (nicht hermetisch) — CVE-Scan gegen das **publizierte** Image |
| [`make regelwerk-check`](harness/sensors/regelwerk-check.md) | **kein Gate** — misst die Integrität der vendored Baseline gegen `SHA256SUMS`, fail-closed. Die Freshness-Hälfte bleibt als Netz-Operation ungeprüft |
| `make slice-mv` | **kein Gate** — ein Werkzeug: Lifecycle-Wechsel eines Slice per `git mv` samt der Verweise auf ihn |
| `make gates` | alle inneren Gates (mandatory vor Handoff) |
| [`make verify-risiko-ausgaenge`](harness/sensors/verify-risiko-ausgaenge.md) | Jedes in §6 **notierte** Risiko trägt genau einen Ausgang aus der geschlossenen Dreier-Menge |
| [`make verify-observations`](harness/sensors/verify-observations.md) | Deckung des Beobachtungs-Registers; der Zähler wird abgeleitet, nicht geführt |
| `make verify` | **Verifikations-Schicht** (getrennt von `gates`): DoD-/Closure-Fragen statt Code-Fragen |
| `make image-test` | [AC-FA-DIST-001](spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk) + nativ==Container + Fragment-Parität |
| `make ci` | CI-äquivalent: `gates` + `image-test` (Workflow `.github/workflows/ci.yml`) |
| `make trace-check` | Traceability via Modul `commits`: `AC-*`/`ADR-*`/`MR-*`/`slice`-ID je Commit (§5) |
| `make commit-scope-check` | Commit-Scope `(planning)` berührt nur `docs/plan/planning/` (§5); misst jeden Commit an der damals geltenden Fassung |
| `make archive-wave-test` | Testsuite von `tools/archive-wave/` (eigenes `go.mod` — **nicht** Teil von `make test`, das nur das Hauptmodul deckt) |
| [`make archive-wave`](harness/sensors/archive-wave.md) | **kein Gate** — bewegt Zeitdokumente ins Archiv, Volltexte werden Stubs (`WELLE=`/`SLICE=`, `APPLY=1`) |

Diese Tabelle **listet auf**; definiert wird hier nichts. Die *Bindung* eines
Targets — welche Anforderung oder Entscheidung es durchsetzt — steht in
[`harness/README.md`](harness/README.md) §Sensors; von dort führt der Weg zur
`AC-*`-ID, zur ADR oder zum Carveout. Wo ein Target mehr braucht als seine
Zelle — Grenze, Ausgänge, Sperren —, steht das in
[`harness/sensors/`](harness/sensors/); `make doc-mentions` hält beide Enden
zusammen.

## 5. Dokumentations-Regeln

- Commits/PRs müssen mindestens eine `AC-*`- oder `ADR-*`-ID nennen
  (auch `MR-*`/`slice-NNN` gelten). Durchgesetzt durch `make trace-check`
  (d-check-Modul `commits`, [ADR-0021](docs/plan/adr/0021-commits-modul-trace-check.md))
  — lokal über `HEAD~1..HEAD`, in der CI über den Commit-Range
  ([`.github/workflows/ci.yml`](.github/workflows/ci.yml)). IDs werden nur
  beim Spec-/ADR-Schreiben nach dem deklarierten Schema vergeben (siehe
  [`harness/conventions.md`](harness/conventions.md)) — nie ad hoc im
  Commit/PR; Agenten referenzieren IDs, sie erfinden keine.
  **Struktur-IDs zählen nicht:** `SPEC-<NNN>` adressiert *innerhalb* der
  Spezifikation und gehört nicht in die Commit-Message — die `id-patterns` in
  [`.d-check.yml`](.d-check.yml) führen sie folgerichtig nicht.
- **Commit-Scope `(planning)`:** ein Commit mit diesem Scope (`docs(planning)`,
  `fix(planning)`, `chore(planning)`) berührt **ausschließlich**
  `docs/plan/planning/`. Wandert Substanz eines anderen Bereichs mit, ist das ein
  eigener Commit mit passendem Scope. Durchgesetzt durch `make commit-scope-check`;
  jeder Commit wird an der Fassung gemessen, die zu **seinem** Zeitpunkt galt.
  **Nur dieser Scope ist geregelt:** Bei `docs(spec)` und `docs(adr)` ist der
  Fremd-Bereich legitim, und eine Regel, die den Bestand massenhaft bricht, wird
  abgeschaltet statt befolgt. Ein weiterer Scope wird erst geregelt, wenn er
  auffällt — und dann gemessen, nicht geraten.
- **Wer eine Anforderung anlegt, nennt ihre Kennung in der Closure-Notiz.** Im **Plan** kann er es
  nicht: IDs werden referenziert statt erfunden, die neue Kennung existiert dort noch nicht, und
  jede genannte ist linkpflichtig — ein Link ins Leere macht `doc-check` rot. Also umschreibt der
  Plan sie („eine neue `AC-FA-CLI`-Kennung"), und die Requirements-Matrix sieht den Slice **nicht**.
  Bei der Closure ist die Anforderung geschrieben; dort steht die Kennung mit Link. Durchgesetzt
  durch `make doc-complete` im `verify`-Aggregat — eine Anforderung ohne
  referenzierenden Slice ist abschluss-blockierend.
- Neue oder geänderte `AC-*`-Anforderungen entstehen nur in
  [`spec/lastenheft.md`](spec/lastenheft.md) — nie per ADR (ADRs schärfen
  die Spezifikation, nicht das Lastenheft).
- Neue ADRs müssen den ADR-Index aktualisieren.
- Roadmap/Status-Geschichte lebt in `docs/plan/planning/`, nicht in der
  Architektur-Spec.
- **Slice-Lifecycle** ist reine Datei-Bewegung — der Zustand ist das
  Verzeichnis, kein Feld im Dokument. Die **fünf** Übergänge und ihre Trigger
  stehen in `modul-05` §Trigger je Lifecycle-Übergang und WIP-Limit; a-check
  fährt sie mit **`make slice-mv`**, das den `git mv` samt der Verweise **auf**
  die Datei erledigt (§3.3).
  **Repo-eigen daneben:** Der direkte Weg `open/ → in-progress/` ist zulässig;
  `next/` ist ein Ort, keine Pflichtstation
  ([`next/README.md`](docs/plan/planning/next/README.md)).
- **WIP-Limit = 1.** Es liegt **höchstens ein** Slice in `in-progress/` (die Roadmap
  zählt nicht mit). Das ist eine harte Obergrenze, kein Vorschlag: zwei aktive Slices
  teilen sich einen Gate-Nachweis und eine Closure-Aufmerksamkeit, und beides
  trägt nur einmal. **Null ist zulässig** — nach jedem Abschluss der Normalfall,
  bis der nächste Slice gezogen wird; ein Maximum ist kein Minimum.
- **AC-Form:** die Pflicht-Bausteine einer Anforderung stehen in
  [`harness/conventions.md`](harness/conventions.md) §Anforderungs-Anlege-Prozess
  — dort seit jeher die drei Pfade (Happy/Boundary/Negative im
  Given/When/Then-Stil) plus Out-of-Scope. `make verify` prüft die Form für
  **neue** `AC-*`; die **19** grandfatherten sind ausgenommen (vertraglich
  bindend, Rand- und Negativfälle bereits in Prosa — ein Umbau träfe die Form
  statt der Substanz), und die Liste wächst nicht mit.
- **Diskrepanz-Trichter:** eine Ausnahme wird **nicht** ad hoc gesetzt. Die
  Werkzeug-Wahl — BF-Sub-Area-Markierung, Carveout oder permanente ADR — steht
  in `modul-07` §Werkzeug-Wahl bei Diskrepanz; die Ablageorte hier sind
  [`harness/conventions.md`](harness/conventions.md#modus-deklaration-pro-sub-area)
  bzw. [`docs/plan/carveouts/`](docs/plan/carveouts/README.md).
- **Beobachtungs-Register:** Form und Zählregel stehen in `modul-06` §Das
  Beobachtungs-Register; die **drei Ausgänge eines offenen Risikos**
  — *eingetreten* ⇒ Carveout oder Folge-Slice · *entfallen* ⇒ gestrichen **mit
  Begründung** · *weiter offen* ⇒ Register — in `modul-05` §Offene Risiken
  werden bei Closure aufgelöst. Es sind **nicht** dieselben drei wie die
  Register-Ausgänge (*verkörpert · geplant · gestrichen*), und
  `make verify-risiko-ausgaenge` setzt genau die erste Menge durch.
  **Repo-eigen ist der Ort und das Kürzel:** die stehende Ablage
  [`docs/plan/planning/observations/`](docs/plan/planning/observations/README.md),
  je Beobachtung ein Verzeichnis `BEO-<KUERZEL>/<slug>/`, und `<KUERZEL>` wird
  in [`harness/conventions.md`](harness/conventions.md#modus-deklaration-pro-sub-area)
  §Modus-Deklaration **nachgeschlagen, nicht erfunden**.
- **Steering-Loop:** wiederkehrende Fehlermuster werden in
  [`docs/plan/steering-loop.md`](docs/plan/planning/observations/README.md) gezählt. Ab dem
  **zweiten** gleichartigen Vorfall entsteht ein Eintrag, ab dem **dritten** ist
  es eine Harness-Lücke und verlangt einen Guide oder Sensor — „besser
  aufpassen" ist keine Antwort. Ein Eintrag ohne Vorfallszahl ist unzulässig:
  die Zahl ist das Einzige, was die Schwelle prüfbar macht.
- **Zitier-Form in einfrierenden Artefakten** (`v6.5.0`, vier Ziel-Formen:
  Review-Report, Welle-Ergebnisnotiz, beide Archiv-Stubs): Was einfriert,
  zitiert **Kennung statt Adresse** — `slice-NNN` statt seines Lifecycle-Pfads,
  `make <target>` statt eines Links auf die Sensor-Datei, eine Baseline-Stelle
  als **Tag + Pfad in Inline-Code** statt als Link
  (`` `v<X.Y.Z>` · `regelwerk/<datei>.md` §<Abschnitt> ``). Grund: Der vendored
  Baum trägt genau einen Tag, der nächste Sprung löscht den alten, und ein Link
  darauf färbt ein Artefakt rot, das niemand mehr anfassen darf. Verankert im
  Reviewer-Skill; für Archiv-Stubs erzeugt `tools/archive-wave/` den Text, für
  die Ergebnisnotiz gilt sie beim Schreiben.
  **Nicht** betroffen: lebende Dokumente — dort ist der Link richtig, und
  `versions` hält ihn aktuell.
- **Slice-Form:** neue Slices entstehen aus der **vendorten Ziel-Form**
  [`slice.template.md`](.harness/baseline/v6.5.0/templates/docs/plan/planning/slice.template.md) —
  a-check führt keine eigene Kopie, sie würde gegen die Baseline driften. **Was
  beim Kopieren anzupassen ist**, steht in
  [`docs/plan/planning/README.md`](docs/plan/planning/README.md) §Beim Kopieren
  der Slice-Ziel-Form.
- **Zwei Mess-Regeln binden jeden, der einen Beleg schreibt** — also auch den
  Implementer- und den Planner-Lauf, nicht nur den Review:
  1. *Geltungsbereich einer Messung* (`seit slice-179`): Wer eine Messung als
     Beleg schreibt — Slice-Plan, Closure-Notiz, Review-Report —, **nennt ihren
     Geltungsbereich** und sagt, ob er den Gegenstand deckt. Nicht *„22 Befunde,
     keine weitere Klasse"*, sondern *„22 Befunde über Markdown-Links; Prosa
     sieht das Instrument nicht"*.
  2. *Eine Mutations-Probe belegt erst, wenn sie rot war* (`seit slice-181`):
     Wer einen Prüfer mit einer Probe belegt, zeigt **beide** Richtungen und
     nennt die **Meldung** der roten, nicht nur den Exit-Code. Grün beweist
     nichts — ein Prüfer, der seinen Gegenstand nicht erreicht, ist grün.

  **Kein Sensor:** beides ist ein Urteil über eine Absicht bzw. einen Aufbau
  (§3.7). Die **Herleitung** und die gemessenen Fälle stehen im Reviewer-Skill
  ([`.harness/skills/reviewer.md`](.harness/skills/reviewer.md) §Mess-Regeln) —
  dort urteilt, wer prüft; hier steht der Satz, an den sich bindet, wer
  schreibt.
- **CR-Texte an ein fremdes Werkzeug** leben im Slice, der sie erzeugt,
  und gehen erst nach einem Prüf-Durchgang hinaus: der Skill
  [`.harness/skills/cr-text-reviewer.md`](.harness/skills/cr-text-reviewer.md) markiert jeden Satz,
  der eine **Tatsache** über ein System behauptet — das eigene oder das fremde —, und nennt den
  Handgriff, der ihn belegt. Die Prüf-Frage
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
6. Repo-weiten Gate-Lauf vor Handoff (`make gates`).
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
(Modul 8 §Kontext-Trennung).

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
