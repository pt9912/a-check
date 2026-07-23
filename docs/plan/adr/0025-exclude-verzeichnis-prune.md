# ADR-0025 — `exclude` beschneidet den Verzeichnis-Walk (Prune), nicht nur die Datei-Extraktion

- **Status:** Accepted
- **Datum:** 2026-07-23
- **Autor:** pt9912
- **Bezug:** [AC-FA-CONF-001](../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml), [AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)
- **Schärft:** [SPEC-EXTRACT-001](../../../spec/spezifikation.md#spec-extract-001--import-extraktion) + [SPEC-CONF-001](../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema) — Scan-Scope: Datei- **und** Verzeichnis-Prune vor der Extraktion.
- **Supersedes:** — (erweitert [ADR-0018](0018-exclude-scan-scope.md), ersetzt es nicht).

## Kontext

[ADR-0018](0018-exclude-scan-scope.md) führte `exclude` als „Scan-Scope vor der
Extraktion" ein; seine gewählte Option A nennt ausdrücklich die **repo-weiten
Verzeichnis-Klassen** (`node_modules/`, `dist/`) als Motiv. Die Umsetzung filterte
`exclude` jedoch **nur auf Datei-Ebene**: der Verzeichnis-Walk stieg in jeden Ordner
ab und rief `ReadDir` auf, bevor der Datei-Glob je einen Treffer prüfte. Zwei reale
Folgen (Konsumenten-Meldung m-trace):

1. **Fail-closed-Abbruch auf ausgeschlossenem Teilbaum.** Ist ein Ordner **innerhalb**
   eines ausgeschlossenen Teilbaums unlesbar (z. B. root-eigener
   `.security/.trivy-cache/fanal`), liefert `filepath.WalkDir` den Callback mit
   `walkErr != nil`, der Walk propagiert den Fehler → Exit 2 — **obwohl** `.security/**`
   in `exclude` steht. Der Ausschluss verhindert das Betreten nicht.
2. **Unnötiger Durchlauf.** Große Fremdcode-Teilbäume (`node_modules/`, `dist/`) werden
   voll durchlaufen und je Datei gefiltert, statt am Rand beschnitten.

Das Schwester-Tool `d-check` behandelt seinen `scan.ignore` genau als Teilbaum-Prune
(`.security/.trivy-cache/**` stolpert dort nicht) — dieselbe Werkzeug-Familie, dieselbe
erwartete Semantik: `exclude` = „ignoriere diesen Teilbaum".

**Die entscheidende Feinheit:** Die negations-freie Glob-Engine
([SPEC-CONF-001](../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema))
übersetzt `**` nach `.*` (spannt Verzeichnis-Grenzen), aber ein einzelnes `*` nach
`[^/]*` (**ein** Segment). Nur `**`-terminierte Muster decken einen ganzen Teilbaum;
`<dir>/*`, `<dir>/` oder `<dir>/*.go` decken nur einen Teil. Ein Verzeichnis darf also
**nur** dann geprunet werden, wenn ein Muster seinen **ganzen** Teilbaum deckt — sonst
fielen Dateien, die der Glob gar nicht matcht, still aus dem Scope (False-Green).

## Optionen

| Weg | Idee | Bewertung |
|---|---|---|
| **A — Prune nur bei Teilbaum-deckendem Muster (`**` / `<präfix>/**`)** | Ein Verzeichnis `dir` wird beschnitten, wenn ein `exclude`-Glob `**` ist **oder** die Form `<präfix>/**` hat, dessen `<präfix>` `dir` matcht. Datei-Ausschluss ([ADR-0018](0018-exclude-scan-scope.md)) bleibt zusätzlich für Einzeldateien. | **Gewählt.** Realisiert die Verzeichnis-Absicht von ADR-0018; **beweisbar output-äquivalent** zum Datei-Ausschluss (jede Datei unter `dir` matcht `<präfix>/.*` → wäre ohnehin ausgeschlossen), also **kein** Coverage-Verlust. Prune vor `ReadDir` ⇒ unlesbarer/großer Teilbaum bricht nicht mehr ab. |
| **B — Prune bei Verzeichnis-Rand-Treffer (`MatchGlobs(dir + "/", …)`)** | Prunen, sobald der Verzeichnis-Rand `dir/` einen Glob matcht. | **Verworfen (Über-Prune, False-Green).** Ein Single-Segment-`src/*` → `^src/[^/]*$` matcht den Rand `src/`, aber **nicht** tiefere Dateien; der Prune verlöre `src/app/x.go`, das der Nutzer **nicht** ausgeschlossen hat — eine stille Coverage-Lücke (verletzt [AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)). Nur für `…/**`-Formen wäre der Rand-Test korrekt — genau die deckt Weg A präzise ab. |
| **C — unlesbare Ordner generell mit Warnung überspringen (fail-open)** | Jeder `walkErr` auf einem Verzeichnis wird zur Warnung, der Scan läuft weiter. | **Verworfen.** Versteckt Coverage-Lücken **still** — ein versehentlich unlesbarer *echter* Quellordner würde ungeprüft grün. Bricht die fail-closed-Linie (kein stiller No-Op, ein unauflösbarer Zustand endet in Exit 2). Der **explizite** `exclude`-Prune ist der ehrliche Hebel. |
| **D — Verzeichnis-`rel` direkt gegen `exclude` matchen** | Ohne jede Rand-/Präfix-Behandlung das slash-lose Verzeichnis prüfen. | **Verworfen.** `…/**`-Globs matchen den slash-losen Pfad nie → No-Op; träfe genau die realen Muster (`node_modules/**`, `.security/**`) nicht. |

## Entscheidung

**Weg A:**

1. **Verzeichnis-Prune bei Teilbaum-Deckung.** Bei `d.IsDir()` wird der Teilbaum
   **nicht betreten** (`filepath.SkipDir`), wenn ein `exclude`-Glob `**` ist **oder**
   die Form `<präfix>/**` hat und `<präfix>` das (slash-normalisierte) Verzeichnis
   matcht. Der Datei-Ausschluss vor der Extraktion ([ADR-0018](0018-exclude-scan-scope.md))
   bleibt zusätzlich für einzelne Dateien.
2. **Output-Äquivalenz.** Weil nur `<präfix>/**` prunt, ist jede Datei unter dem
   geprunten Verzeichnis von `<präfix>/.*` gedeckt und wäre ohnehin datei-ausgeschlossen
   — die Menge der gescannten Dateien ist **identisch** zur reinen Datei-Filterung.
   Nicht-teilbaum-deckende Muster (`<dir>/*`, `<dir>/`, `<dir>/*.go`) prunen **nicht**.
   Die **Scan-Wurzel** (`rel == "."`) wird nie geprunet.
3. **Failure-Mode.** Der Prune greift **vor** `ReadDir` (Vertrag von
   `filepath.WalkDir`: der Verzeichnis-Callback läuft vor dem Lesen des Inhalts): ein
   unlesbarer oder sehr großer **ausgeschlossener** Teilbaum erzeugt keinen Scan-Fehler
   mehr (vorher Exit 2). Ein **nicht** ausgeschlossener unlesbarer Ordner descendet, sein
   `ReadDir` scheitert, `WalkDir` ruft den Callback ein zweites Mal mit `walkErr` → der
   Scan bricht **weiterhin** fail-closed ab (kein stilles Überspringen —
   [AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)).
4. **Rückwärtskompatibel.** Ohne `exclude` byte-identisch. Mit `exclude` ist die Menge
   der gescannten Dateien identisch zur bisherigen Datei-Filterung (Output-Äquivalenz,
   Punkt 2) — der einzige beobachtbare Unterschied ist der **wegfallende Abbruch** im
   Unlesbar-Fall und die **Laufzeit** (nicht betretene Teilbäume).

## Konsequenzen

- [ADR-0018](0018-exclude-scan-scope.md) bleibt gültig; dies ist die
  **Verzeichnis-Realisierung** seiner erklärten Absicht — **kein** Supersede.
- **Schema/Scan-Scope** ([SPEC-EXTRACT-001](../../../spec/spezifikation.md#spec-extract-001--import-extraktion),
  [SPEC-CONF-001](../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema)):
  Scan-Scope-Schritt nennt Datei- **und** Verzeichnis-Prune (nur Teilbaum-deckende
  Muster) + die Failure-Mode-Grenze.
- **Kein Lastenheft-Bump:** [AC-FA-CONF-001](../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml)
  ist unberührt (der Vertrag „`exclude` nimmt matchende Pfade aus dem Scope" gilt schon;
  geschärft wird das **Wie** und der Wegfall eines Abbruchs — Schärfung aufwärts über die
  Spezifikation).
- Erste Nutzer: m-trace (`.security/**` mit unlesbarem `.trivy-cache/fanal`),
  `**/node_modules/**`/`**/dist/**` (jetzt am Rand beschnitten statt durchlaufen).

## Fitness Function

- `make test`:
  - **Prune-Prädikat** (`dirAction`, tabellarisch, root-frei): `.security/**` / `dist/**` /
    `**/node_modules/**` / `**` prunen; `src/*` / `sub/` / `<dir>/*.go` / Datei-Globs
    prunen **nicht**; `.git` immer geskippt; Scan-Wurzel nie geprunet.
  - **Kein Über-Prune (F-1-Regression, End-to-End):** `exclude: ["src/*"]` mit
    `src/loose.go` (datei-ausgeschlossen) und `src/app/x.go` (bleibt) → `x.go` im Scan.
  - verschachtelter ausgeschlossener Teilbaum (`.security/**`) fällt vollständig aus,
    Geschwister-Dateien bleiben.
  - Datei-Glob (`**/*_test.go`) prunt das enthaltende Verzeichnis **nicht**.
  - ohne `exclude` byte-identisch (bestehender Test).
- `make arch-check` (Dogfooding,
  [AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)):
  die Eigen-Config deklariert kein `exclude` → unverändert 0 Befunde.

## Re-Evaluierungs-Trigger

- **Verzeichnis-Kurzformen** (`node_modules/` statt `**/node_modules/**`) — vom
  [ADR-0018](0018-exclude-scan-scope.md)-Trigger geerbt; eine solche Kurzform als
  erstklassige Verzeichnis-Exklusion (⇒ Prune + Voll-Ausschluss des Teilbaums) wäre
  additiv, braucht aber eine eigene Config-Semantik-Entscheidung.
- **Ausweis des Scan-Scopes im Report** (gescannt/ausgeschlossen/geprunet), falls ein
  Pilot den stillen Prune als Transparenz-Lücke meldet.

## Geschichte

| Datum | Ereignis |
|---|---|
| 2026-07-23 | Proposed — Verzeichnis-Prune nur bei Teilbaum-deckendem Muster (`**` / `<präfix>/**`), output-äquivalent zum Datei-Ausschluss; Rand-Treffer-Prune (B) verworfen (Über-Prune `src/*` → False-Green), fail-open (C) verworfen (versteckt Coverage-Lücken), direkter Verzeichnis-Match (D) verworfen (No-Op gegen `…/**`); Failure-Mode: ausgeschlossener unlesbarer Teilbaum bricht nicht mehr ab, nicht-ausgeschlossener bleibt fail-closed. Evidenz: m-trace-Meldung (`.security/.trivy-cache/fanal`), d-check-`scan.ignore`-Parität; adversarisches Multi-Linsen-Review (Über-Prune-Bug B vor der Landung gefangen). |
| 2026-07-23 | Proposed → Accepted (Sign-off Auftraggeber; spec-first-Review bestanden). Ab jetzt immutable; Ablösung nur via Folge-ADR mit `Supersedes`. |
