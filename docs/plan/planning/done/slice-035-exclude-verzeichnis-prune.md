# slice-035 — `exclude` beschneidet den Verzeichnis-Walk (Prune)

**Status:** in-progress — **fertig, `make ci` grün, 3× adversarisch reviewed (F-1 gefunden+gefixt), Merge ausstehend**. Bei Merge → `git mv` nach `done/`. Closure-Notiz: §7.
**Typ:** Defekt-Fix + Spec-Schärfung (Scan-Scope), Folge von slice-023/[ADR-0018](../../adr/0018-exclude-scan-scope.md).
**Bezug:** schärft [SPEC-EXTRACT-001](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion) via
[ADR-0025](../../adr/0025-exclude-verzeichnis-prune.md); folgt
[AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml) (kein Lastenheft-Bump),
fail-closed-Grenze [AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze).
[Roadmap](../in-progress/roadmap.md).

## 1. Motivation

Konsumenten-Meldung (m-trace): a-check stirbt am root-eigenen
`.security/.trivy-cache/fanal` mit Exit 2, **obwohl** `.security/**` in `exclude` steht.
Ursache verifiziert in `internal/adapter/driven/extract/extract.go`: `exclude` wurde
**nur auf Datei-Ebene** geprüft; der Verzeichnis-Walk (`filepath.WalkDir`) stieg in jeden
Ordner ab und rief `ReadDir` auf, bevor der Datei-Glob je griff. Ist ein Ordner **im**
ausgeschlossenen Teilbaum unlesbar, liefert `WalkDir` den Callback mit `walkErr` → der
Walk propagiert den Fehler → `Extract` bricht ab → CLI Exit 2 (`cli.go:82`).

`exclude` = „ignoriere diesen Teilbaum" ist die erwartete Semantik — das Schwester-Tool
`d-check` behandelt seinen `scan.ignore` genau so (`.security/.trivy-cache/**` stolpert
dort nicht). [ADR-0018](../../adr/0018-exclude-scan-scope.md) nannte die repo-weiten
**Verzeichnis**-Klassen (`node_modules/`, `dist/`) sogar als Motiv — der Prune wurde nur
nie umgesetzt.

**Verworfen** (Meldung selbst empfahl es): unlesbare Ordner still mit Warnung überspringen
— versteckt Coverage-Lücken; der **explizite** `exclude`-Prune ist der ehrliche Hebel
([ADR-0025](../../adr/0025-exclude-verzeichnis-prune.md) Option B).

## 2. Design

**Kernänderung (`extract.go`, `Extract`-Walk):** `rel` einmal oben berechnen; im
`d.IsDir()`-Zweig zusätzlich zum `.git`-Skip prunen, wenn ein `exclude`-Glob den
**ganzen Teilbaum** des Verzeichnisses deckt → `filepath.SkipDir` (Helfer `dirExcluded`).
Datei-Ausschluss ([ADR-0018](../../adr/0018-exclude-scan-scope.md)) bleibt unverändert
für Einzeldateien.

**Prune nur bei Teilbaum-deckendem Muster (`**` / `<präfix>/**`):** die negations-freie
Glob-Engine (`core/rules.go:globToRegexp`) übersetzt `**` → `.*` (spannt Grenzen), ein
einzelnes `*` aber → `[^/]*` (ein Segment). `dirExcluded(dir)` prunt genau dann, wenn ein
Glob `**` ist **oder** die Form `<präfix>/**` hat und `<präfix>` `dir` matcht. Damit ist
der Prune **beweisbar output-äquivalent** zum Datei-Ausschluss: jede Datei unter `dir`
matcht `<präfix>/.*` und wäre ohnehin ausgeschlossen — **kein** Coverage-Verlust.

> **Adversarisches Review (vor der Landung):** der erste Entwurf prunte bei bloßem
> Verzeichnis-**Rand**-Treffer (`MatchGlobs(rel + "/", …)`). Das über-prunt: `exclude:
> ["src/*"]` hätte `src/` geprunet und `src/app/x.go` still verloren, obwohl der Glob die
> Datei nicht matcht (False-Green, [AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)).
> Das Multi-Linsen-Review fing den Bug; die Regel ist auf teilbaum-deckende Muster
> verengt (Weg A statt B in [ADR-0025](../../adr/0025-exclude-verzeichnis-prune.md)).

**Warum der Fail-closed-Fix „gratis" kommt:** `filepath.WalkDir` ruft den Callback für
ein Verzeichnis **vor** `ReadDir` auf; `SkipDir` überspringt das Lesen des Inhalts. Der
ausgeschlossene Teilbaum wird also nie stat/gelesen → kein `walkErr` → kein Exit 2. Ein
**nicht** ausgeschlossener unlesbarer Ordner läuft weiter in den unveränderten
`return walkErr`-Pfad (fail-closed bleibt).

## 3. Geplanter Umfang (umgesetzt)

1. **Code:** `internal/adapter/driven/extract/extract.go` — Teilbaum-deckender
   Verzeichnis-Prune (`dirAction`/`dirExcluded`). ✅
2. **Spec:** [SPEC-EXTRACT-001](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion)
   Scan-Scope-Schritt (Datei **und** Verzeichnis, nur `…/**`-Muster + Failure-Mode) +
   [SPEC-CONF-001](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema)
   `exclude`-Zeile; Version 0.20.0 → **0.21.0** + Historie-Zeile. ✅
3. **ADR:** [ADR-0025](../../adr/0025-exclude-verzeichnis-prune.md) `Accepted` (erweitert
   [ADR-0018](../../adr/0018-exclude-scan-scope.md), kein Supersede) + ADR-Index-Zeile. ✅
4. **Tests:** in `extract_test.go` — `TestDirActionPrunePredicate` (Tabelle, root-frei,
   diskriminierend), `TestExtractExcludeNoOverPruneOnSingleStar` (F-1-Regression +
   Nicht-Sprach-Datei), `TestExtractExcludePrunesNestedSubtree` (E2E `.security/**`),
   `TestExtractExcludeDoesNotPruneOnFileGlob`. ✅
5. **Handbuch:** `exclude`-Abschnitt + Version 1.29 → 1.30. ✅
6. **Gates:** `make gates` (+ `make ci`) — **ausstehend** (nach der Review-Revision).

## 4. Akzeptanzkriterien (als Test/Gate)

- **Happy (Prune-Entscheidung, diskriminierend):** `dirAction` prunt `.security` /
  `dist` / `a/node_modules` / `**` und prunt `src`(`src/*`) / `sub`(`sub/`) / `core`
  (`**/*_test.go`) **nicht**; `.git` immer geskippt, Scan-Wurzel nie geprunet.
  `TestDirActionPrunePredicate`.
- **Negative (kein Über-Prune, F-1, End-to-End):** Given `exclude: ["src/*"]` mit
  `src/loose.go` (datei-ausgeschlossen) und `src/app/x.go`, when `Extract` läuft, then
  bleibt `src/app/x.go` im Scan (kein stiller Teilbaum-Verlust).
  `TestExtractExcludeNoOverPruneOnSingleStar`.
- **Boundary (E2E realer Fall):** Given `exclude: [".security/**"]` mit
  `.security/cache/x.go` und Geschwister `core.go`, when `Extract` läuft, then nur
  `core.go` bleibt. `TestExtractExcludePrunesNestedSubtree`.
- **Regression:** ohne `exclude` byte-identisch (bestehender
  `TestExtractExcludeSkipsFiles`, zweiter Zweig).

## 5. Grenzen / Folge

- **Der unlesbar-Teilbaum-Fall ist im root-laufenden Test nicht direkt reproduzierbar**
  (die `test`-Stage läuft als root; `chmod 0o000` blockt root nicht). Die Failure-Mode-Grenze
  ist stattdessen durch den `WalkDir`-Stdlib-Vertrag garantiert (Prune vor `ReadDir`) und
  in ADR/Spec dokumentiert; die **Prune-Entscheidung** selbst ist über `dirAction`
  deterministisch und root-frei getestet (Review-Empfehlung umgesetzt).
- **Kein Lastenheft-/Vertragswechsel:** [AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml)
  bleibt; geschärft wird das **Wie** (Walk-Prune) und der Wegfall eines Abbruchs.
- **Nicht-idiomatische `…/**`-Äquivalente** (`**/`, `foo/**/`, `a/**/**`) prunen
  **nicht** (das Prädikat verlangt das kanonische Suffix `/**`) — der Output bleibt
  korrekt (die Datei-Schleife schließt die Dateien aus), nur die Prune-Optimierung und
  die Unlesbar-Robustheit entfallen für diese Schreibweisen. Bewusste Grenze
  ([AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)),
  vom Verifikations-Review als LOW/Nicht-Blocker bestätigt; kanonisch schreibt man `<dir>/**`.
- **Rand-genaue Verzeichnis-Exklusion** (`legacy/` als erstklassige „exclude dieses
  Verzeichnis"-Form ⇒ Prune + Voll-Ausschluss) ist bewusst **nicht** Teil dieses Slice —
  sie wäre eine neue Config-Semantik ([ADR-0025](../../adr/0025-exclude-verzeichnis-prune.md)
  Re-Eval-Trigger). Heute deckt man einen Teilbaum mit `<dir>/**`.
- **Unveröffentlicht** — die Release-Achse (GHCR-Tag) zieht getrennt nach (Re-Pin nach
  releasing.md, falls der Fix in ein Image soll; separat vom Lastenheft-Stand).

## 6. Sub-Area-Modus-Begründung

### Sub-Area: Extraktions-Adapter (`internal/adapter/driven/extract`)

- **Modus:** GF — Spec/Test führt, erzwungen über `make gates`.
- **Konventionen-Dichte:** hoch — `exclude` etabliert (slice-023), Walk-Logik klein und lokal.
- **Phase-Reife:** Phase 4 — Adapter real und gegen die Registry getestet.
- **Evidenz-/Diskrepanz-Risiko:** niedrig — der `dirAction`-Tabellentest isoliert die
  Prune-Entscheidung deterministisch; die einzige nicht-testbare Facette (root/unlesbar) ist stdlib-garantiert.
- **Reconciliation-Aufwand:** keiner erwartet.

## 7. Closure-Notiz (nach `done`)

**Abgeschlossen 2026-07-23** auf Branch `slice-035-exclude-verzeichnis-prune`.

**Geliefert:** Verzeichnis-Prune bei teilbaum-deckendem `exclude`-Muster
(`extract.go` `dirAction`/`dirExcluded`); der Walk beschneidet ausgeschlossene
`…/**`-Teilbäume, statt sie zu durchlaufen — der m-trace-Auslöser (Exit 2 am
ausgeschlossenen unlesbaren `.security/.trivy-cache/fanal`) ist getilgt, weil der
Prune vor `ReadDir` greift. Spec-first: [SPEC-EXTRACT-001](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion)/[SPEC-CONF-001](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema)
0.21.0, [ADR-0025](../../adr/0025-exclude-verzeichnis-prune.md) `Accepted` (erweitert
[ADR-0018](../../adr/0018-exclude-scan-scope.md)), Benutzerhandbuch 1.30. Kein
Lastenheft-Bump (Schärfung des Wie).

**Gate-Evidenz:** `make ci` grün — lint 0, test (inkl. `TestDirActionPrunePredicate` +
F-1-Regression), arch-check 0 (Dogfooding), doc-check 96/0, gate-consistency/guard/record
ok, image-test alle Blöcke (nativ==Container).

**Review:** adversarisches Multi-Linsen-Review (3 Winkel) fand **F-1** — der erste
Rand-Test-Ansatz über-prunte `src/*` und hätte `src/app/x.go` still verloren (False-Green,
[AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)).
Fix: Prune auf teilbaum-deckende Muster verengt (beweisbar output-äquivalent). Ein
Verifikations-Review bestätigte die revidierte Regel per Brute-Force (~150 000
Kombinationen, 0 Invarianten-Verletzung); Restbefund 3× nicht-idiomatische Unter-Prune-
Schreibweise (LOW, kein False-Green, dokumentiert §5).

### Lerneintrag

**Ein „Verzeichnis ausschließen" über Datei-Globs ist eine Teilbaum-Deckungs-Frage, kein
Rand-Match.** Der naheliegende Rand-Test (`dir/` matcht) prunt zu viel: ein
Single-Segment-Glob (`src/*`) matcht den Rand, deckt aber nicht den Teilbaum — pruning
darauf wirft nicht ausgeschlossene Dateien still weg. Lehre: einen Walk nur dort
beschneiden, wo der Prune **beweisbar output-äquivalent** zur Datei-Filterung ist
(hier: nur `<präfix>/**`). Das Multi-Linsen-Review fing den Bug vor der Landung — die
diskriminierende Testtechnik (Glob, der den Rand, aber nicht die Datei matcht) war
zugleich der Bug-Detektor.

**Folge:** rand-genaue Verzeichnis-Exklusion (`legacy/`) + Scan-Scope-Report-Ausweis als
[ADR-0025](../../adr/0025-exclude-verzeichnis-prune.md)-Re-Eval-Trigger vermerkt; d-check-
Pin-Hebung als slice-036 vorgemerkt.
