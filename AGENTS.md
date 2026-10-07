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
[`.harness/baseline/v6.13.0/regelwerk/README.md`](.harness/baseline/v6.13.0/regelwerk/README.md)
ist der Index (17 Module + acht Grundlagen-Abschnitte, eine Datei je
Abschnitt); die Ziel-Formen daneben unter
[`templates/`](.harness/baseline/v6.13.0/templates/README.md). Integrität:
`.harness/baseline/v6.13.0/SHA256SUMS`.

**Breiterer Pflicht-Blick bleibt bei drei Anlässen** — dort genügt der eine
Abschnitt nicht: Bootstrap · jede Änderung an
[`harness/conventions.md`](harness/conventions.md) (Adaptionen `MR-<NNN>`,
Source Precedence, ID-Schema) · Drift-Audit gegen die Baseline
(`modul-02` §Freshness-Audit der vendored Baseline — darunter die Stichprobe
gegen den Bestand, die **auch bei aktuellem Pin** läuft; `make regelwerk-check`
deckt davon nur die Integritäts-Hälfte, siehe §4).

Die vendorten **Ziel-Formen** unter
[`templates/`](.harness/baseline/v6.13.0/templates/README.md) tragen **zwei
Rollen**: als **Referenz-Form**, auf die das Regelwerk mit `../templates/…`
verweist, und als **Vorlage, die beim Anlegen kopiert und ausgefüllt wird statt
frei formuliert** — für ADR, Slice, Welle, Carveout, Review-Report und
Beobachtung. a-check führt **keine eigenen Kopien** davon; was beim Kopieren
anzupassen ist, steht am Ort der jeweiligen Ablage (für Slices:
[`docs/plan/planning/README.md`](docs/plan/planning/README.md) §Beim Kopieren
der Slice-Ziel-Form).

Das vendored Regelwerk ist ein **didaktik-freier Extrakt** und trägt keine
eigene Normativität: bei Konflikt gilt der Kurs
([`v6.13.0`](https://github.com/pt9912/ai-harness-course/tree/v6.13.0)), über
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

Regeln dieser Sektion: Baseline-Regelwerk `grundlagen-harness-dateien.md`
§harness/README.md als Einstiegspunkt.

**Der Gate-Index steht einmal**, in [`harness/README.md`](harness/README.md)
§Sensors — dort steht auch die *Bindung* jedes Targets, und von dort führt der
Weg zur `AC-*`-ID, zur ADR oder zum Carveout. Wo ein Target mehr braucht als
seine Zelle (Grenze, Ausgänge, Sperren), steht das in
[`harness/sensors/`](harness/sensors/). **Diese Datei führt die Liste nicht.**

**Dieselbe Regel gilt für jede abschließende Aufzählung neben einer
maschinenlesbaren Quelle** — `Makefile`, `modules:` in
[`.d-check.yml`](.d-check.yml), ein Verzeichnis. Sie ist beim Schreiben richtig
und wird ohne Vorwarnung falsch; sie schrumpft auf einen **Zeiger** („welche
genau, sagt das `Makefile`"), statt die Menge ein zweites Mal zu nennen.

**Kein Target nennen, das im Makefile nicht existiert** — auch nicht in Prosa.
Halluzinierte Gates sind die häufigste Form von Harness-Lüge. Die maschinelle
Hälfte dieser Regel ist `make doc-targets`: Es hält den Index gegen die
`Makefile`-Regeln, in **beiden** Richtungen — kein behauptetes Target ohne
Regel, keine Regel ohne Eintrag im Index. Die Autoritäts-Doku ist
`harness/README.md`, und es gibt genau eine.

**Ein konfigurierter Block ist erst aktiv, wenn er geschaltet ist — zwei
Schritte.** Ein Block in [`.d-check.yml`](.d-check.yml) wirkt nur, wenn sein
Modul zusätzlich in `modules:` steht oder über ein `--enable <modul>` in
`Makefile`/`d-check.mk` geladen wird. Nur der erste Schritt lässt das Modul
**gültig und wirkungslos**: Der Lauf meldet grün, ohne geprüft zu haben —
dieselbe Klasse wie ein behauptetes Target ohne Makefile-Regel, und derselbe
Grund, warum `make doc-targets` beide Richtungen prüft. Nach dem Aktivieren
gehört die **Gegenprobe** dazu: einen Verstoß der Klasse einbauen und den
**Befund** sehen, nicht den Exit-Code.

**Mandatory** ist, was in einem der beiden Aggregate hängt: `gates` (Code-Fragen)
oder `verify` (DoD-/Closure-Fragen). Welche Targets das sind, sagt das
[`Makefile`](Makefile); ob eines gerade grün ist, sagt die CI (Badge im
[`README.md`](README.md)).

## 5. Dokumentations-Regeln

<!-- Index-Tabelle nach der Ziel-Form v6.13.0 (Welle 149): kurze Regeln
     stehen vollstaendig in der Tabelle (Datei-Spalte "-"); waechst eine
     Regel ueber einen Satz hinaus, wandert ihr Volltext nach
     harness/rules/<name>.md und die Tabellenzeile bleibt der
     Kurzform-Zeiger. Konversion: slice-201. -->

| # | Regel | Datei |
|---|---|---|
| 1 | Commit/PR-Traceability: jede Message nennt eine Kennung — Volltext: [`harness/rules/commits-pr-ids.md`](harness/rules/commits-pr-ids.md) | `harness/rules/commits-pr-ids.md` |
| 2 | Commit-Scope `(planning)` berührt ausschließlich `docs/plan/planning/` — Volltext: [`harness/rules/commit-scope-planning.md`](harness/rules/commit-scope-planning.md) | `harness/rules/commit-scope-planning.md` |
| 3 | Die neue Kennung steht in der Closure-Notiz — Volltext: [`harness/rules/anforderungs-kennung-closure.md`](harness/rules/anforderungs-kennung-closure.md) | `harness/rules/anforderungs-kennung-closure.md` |
| 4 | Neue oder geänderte `AC-*`-Anforderungen entstehen nur in [`spec/lastenheft.md`](spec/lastenheft.md) — nie per ADR (ADRs schärfen die Spezifikation, nicht das Lastenheft). | — |
| 5 | Neue ADRs müssen den ADR-Index aktualisieren. | — |
| 6 | Roadmap/Status-Geschichte lebt in `docs/plan/planning/`, nicht in der Architektur-Spec. | — |
| 7 | Slice-Lifecycle ist reine Datei-Bewegung — sechs Übergänge, gefahren mit `make slice-mv` — Volltext: [`harness/rules/slice-lifecycle.md`](harness/rules/slice-lifecycle.md) | `harness/rules/slice-lifecycle.md` |
| 8 | WIP-Limit = 1 pro Lauf; auf dem Hauptzweig: höchstens ein Slice in `in-progress/` — Volltext: [`harness/rules/wip-limit.md`](harness/rules/wip-limit.md) | `harness/rules/wip-limit.md` |
| 9 | AC-Form: drei Pfade plus Out-of-Scope, geprüft für neue `AC-*` — Volltext: [`harness/rules/ac-form.md`](harness/rules/ac-form.md) | `harness/rules/ac-form.md` |
| 10 | Diskrepanz-Trichter: die Werkzeug-Wahl steht in `modul-07` — Volltext: [`harness/rules/diskrepanz-trichter.md`](harness/rules/diskrepanz-trichter.md) | `harness/rules/diskrepanz-trichter.md` |
| 11 | Beobachtungs-Register: Form, Zählregel und die drei Risiko-Ausgänge — Volltext: [`harness/rules/beobachtungs-register.md`](harness/rules/beobachtungs-register.md) | `harness/rules/beobachtungs-register.md` |
| 12 | Steering-Loop: ab dem zweiten Vorfall Eintrag, ab dem dritten Lücke — Volltext: [`harness/rules/steering-loop.md`](harness/rules/steering-loop.md) | `harness/rules/steering-loop.md` |
| 13 | Zitier-Form in einfrierenden Artefakten: Kennung statt Adresse — Volltext: [`harness/rules/zitier-form-einfrierende.md`](harness/rules/zitier-form-einfrierende.md) | `harness/rules/zitier-form-einfrierende.md` |
| 14 | Slice-Form: aus der vendorten Ziel-Form, Anpassungen laut Planning-README — Volltext: [`harness/rules/slice-form.md`](harness/rules/slice-form.md) | `harness/rules/slice-form.md` |
| 15 | Fünf Mess-Regeln binden jeden, der einen Beleg schreibt — Volltext: [`harness/rules/mess-regeln.md`](harness/rules/mess-regeln.md) | `harness/rules/mess-regeln.md` |
| 16 | CR-Texte an ein fremdes Werkzeug gehen erst nach einem Prüf-Durchgang hinaus — Volltext: [`harness/rules/cr-texte-fremdwerkzeug.md`](harness/rules/cr-texte-fremdwerkzeug.md) | `harness/rules/cr-texte-fremdwerkzeug.md` |
| 17 | Closure-Pflicht: genau ein Closure-Abschnitt, ausgefüllt, mit Lerneintrag — Volltext: [`harness/rules/closure-pflicht.md`](harness/rules/closure-pflicht.md) | `harness/rules/closure-pflicht.md` |

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
7. Doku/Indizes aktualisieren, falls ein öffentlicher Vertrag berührt. **Die CHANGELOG-Zeile
   gehört hierher — nicht in die Release-Vorbereitung.** Der CHANGELOG ist die *kuratierte
   Begründung* eines Releases; wer ihn beim Taggen schreibt, rekonstruiert aus `git`, statt zu
   bezeugen. Ein Slice, der einen **öffentlichen Vertrag** ändert — Lastenheft, Spezifikation,
   Benutzerhandbuch oder eine Regel —, trägt seinen Eintrag in `[Unreleased]` **in sich**. Kein Gate deckt das: `gate-consistency` vergleicht nur
   Versions-Nummern, nicht ob eine Änderung einen Eintrag hat.
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
