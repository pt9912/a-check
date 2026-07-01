# slice-017 — Unbekannter `languages`-Schlüssel → Exit 2 (falsch-grün-Falle schließen)

**Status:** done (2026-07-01). Umsetzung + `make gates` + adversarisches Multi-Linsen-Review
(3 Linsen) + Delta erledigt; Synthese [`docs/reviews/2026-07-01-slice-017-unbekannte-sprache-exit2.md`](../../../reviews/2026-07-01-slice-017-unbekannte-sprache-exit2.md).
**Bezug:** schärft [AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml)
(Config-Validität) — die zulässige Backend-Menge **besitzt** [AC-FA-EXTRACT-001](../../../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion);
Motiv [AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)
(ehrliche Grenze). Kein neuer ADR (innerhalb [ADR-0003](../../adr/0003-config-modell-a-check-yml.md)/
strict-config). [Roadmap](../in-progress/roadmap.md). **Evidenz:** b-cad-Pilot + Polyglot-Bestand.

## 1. Auslöser

Ein `languages`-Schlüssel außerhalb der unterstützten Backends (`cpp`/`go`/`rust`/`kotlin`/`java`)
wird heute **still ignoriert**: `langFor` (`extract.go:90-102`) gibt den Config-Key **roh** zurück,
`importsFromSource` trifft `default: return nil` (`extract.go:104-118`) — der Config-Decode akzeptiert
den Schlüssel kommentarlos, weil `languages` eine **Map** ist (`config.go:44`) und `KnownFields(true)`
(`config.go:68`) nur **Struct-Felder** prüft, nicht Map-Keys. Folge: `languages: {python: ["**/*.py"]}`
extrahiert **nichts** → `0 Befunde` → **falsch-grün**. Das widerspricht dem „keine stillen Defaults"-
Ethos (ein unbekannter *Struct*-Schlüssel bricht sonst mit Exit 2 ab) und ist gefährlicher als ein
sichtbarer Fehler, weil es Sicherheit vortäuscht.

Die Härtung ist **unabhängig vom Hinzufügen** eines Backends (z. B. Python): auch danach bleibt jeder
Tippfehler (`pythn`) und jede weitere Sprache (`swift`) ein Kandidat für die Falle — nicht durch das
Python-Backend „mit-erledigt".

## 2. Betroffene Artefakte (vor der Implementierung benannt)

- **Slice-ID:** slice-017.
- **AC:** [AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml)
  (neue Negative-AC), Menge aus [AC-FA-EXTRACT-001](../../../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion).
  **Kein ADR** (strict-config, [ADR-0003](../../adr/0003-config-modell-a-check-yml.md)).
- **Spec:** [SPEC-EXTRACT-001](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion)
  (Owner der Backend-Menge, normativ), [SPEC-CONF-001](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema)
  (**verweist** darauf — kein Duplikat).
- **Module:** `internal/adapter/driven/extract` (Backend-**Registry** + Eingabe-Validierung),
  `internal/cli` (mappt Extract-Fehler → Exit 2, **bereits vorhanden**, `cli.go:47-50`).
- **Version:** Lastenheft/Spezifikation **0.8.0 → 0.9.0**.
- **Gates:** `make gates` → `make ci`.

## 3. Umfang (Reihenfolge: Lastenheft → Spec → Code → Tests)

1. **Lastenheft 0.9.0** (Prozess nach [`AGENTS.md` §5](../../../../AGENTS.md) / [`harness/conventions.md`](../../../../harness/conventions.md)):
   [AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml)
   neue **Negative-AC** im Given/When/Then-Stil (ein `languages`-Schlüssel außerhalb der
   Backend-Menge → Exit 2); **Versions-Bump + Historie-Zeile**.
2. **Spezifikation 0.9.0:** [SPEC-EXTRACT-001](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion)
   nennt die Backend-Menge **normativ** (Owner); [SPEC-CONF-001](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema)
   **verweist** darauf (Link statt Kopie — sonst dieselbe Drift, die dieser Slice im Code beseitigt,
   eine Ebene höher); **Historie-Zeile**.
3. **Code:** die `switch`-Dispatch in `extract.go` durch eine **Backend-Registry**
   `map[string]func(src string) []core.Import` ersetzen (in `newAdapter` gebaut). Ihre **Keys sind
   die Menge** (echte Single Source). `importsFromSource` schlägt in der Map nach — **kein `switch`,
   kein stiller `default: return nil`** mehr. `Extract` **validiert** die `m.Languages`-Keys gegen
   die Registry **vor** dem Walk; unbekannt → Fehler (den `cli.go` bereits auf Exit 2 mappt).
4. **Fehlermeldung** parallel zu `validRole`/`validDirection` (`config.go:145,148`):
   `%s: unbekannte Sprache %q (cpp|go|rust|kotlin|java)` — die Auswahl-Liste aus den **Registry-Keys**
   generiert (Menge kommt auch in der Meldung aus der Single Source).
5. **Tests:** `python:` → `cli.Run == 2` + Meldung nennt die Menge; jede unterstützte Sprache
   lädt/extrahiert unverändert; Registry deckt genau `{cpp,go,rust,kotlin,java}` (Determinismus/
   Rückwärtskompat der Extraktion).

## 4. Design-Entscheidungen (aufgelöst)

- **Wo validiert wird → im Extraktions-Adapter** (er besitzt die Registry). **Nicht `config.go`**:
  es dürfte `extract` nicht importieren — beide sind Sub-Einheiten der `adapters`-Schicht →
  `lateral-adapter` (kategorisch, a-check würde sich selbst anmeckern). `cli.go` (Composition Root)
  mappt den Extract-Fehler schon auf Exit 2.
- **Single Source echt** (statt behauptet): die Registry-**Map** ersetzt den `switch`; ein neues
  Backend ist **ein** Map-Eintrag (Name + Extractor-Closure) — Validierungs-Menge und Dispatch sind
  danach **dieselbe** Datenstruktur, kein Drift-Guard-Test als Krücke nötig. (Die Regex-Felder im
  Adapter-Struct bleiben; die Closures referenzieren sie. `--print-config`-`sampleConfig` bleibt ein
  **Beispiel**, keine Mengen-Quelle.)
- **Kein stilles Netz mehr:** mit der Registry entfällt `default: return nil`; eine nicht
  registrierte Sprache erreicht `importsFromSource` nach der Validierung nicht — defense-in-depth
  gegen künftige Bypässe (z. B. programmatisch gebautes `Model` in Tests).
- **Verworfen:** `core.SupportedLanguages` + Validierung in `config.Load` — Domänen-Reinheits-Dehnung
  („welche Backends existieren" ist Adapter-Fähigkeit, keine Domäne-Invariante).

## 5. Definition of Done

- [x] Lastenheft 0.9.0 (neue Negative-AC im G/W/T-Stil + Historie) + Spezifikation 0.9.0
  (Backend-Menge-Owner + Verweis statt Kopie + Historie).
- [x] Code: Backend-Registry-Map (Single Source), Extract-Validierung, kein `default: return nil`;
  Meldung aus den Registry-Keys.
- [x] Tests: unbekannt → Exit 2 (+ exaktes Meldungs-Literal), unterstützte unverändert,
  Registry == `{cpp,go,rust,kotlin,java}`, Mono-Repo-Fälle, Case-Sensitivität, stderr-nicht-stdout.
- [x] `make gates` grün; Multi-Linsen-Review (3 Linsen) + Delta; Synthese unter `docs/reviews/`.
- [ ] **Nicht** beim Python-Slice als „gedeckt" abräumen (§1: unabhängige Härtung) — Dauer-Merker.

## 6. Closure-Notiz

**Gate-Beleg:** `make gates` grün — `arch-check` (Dogfooding) 0, `doc-check` 0, alle Test-Pakete `ok`
(inkl. 7 slice-017-Tests + Härtungen); `record-gates` geschrieben.

**2 beobachtbare Kriterien:**
1. `languages: {python: […]}` (oder Mono-Repo `go`+`typescript`) → **Exit 2**, Meldung
   `unbekannte Sprache "python" (cpp|go|java|kotlin|rust)` auf stderr (`TestUnknownLanguageExit2`,
   `TestMonoRepoMixedUnsupportedExit2`) — statt vorher still `0 Befunde`.
2. `go`+`cpp` (beide unterstützt) → Exit 0 (`TestMonoRepoMultiSupportedRuns`).

**Lerneintrag:**
- Die **Registry-Map** löste zwei Review-Findings in einem: echte Single Source (Dispatch **und**
  Validierung teilen die Keys) **und** Wegfall des stillen `default: return nil` — statt „Menge
  validieren + Switch getrennt pflegen + Drift-Guard-Test".
- Das Review deckte eine **vorbestehende Spec-Überzusage** auf: „Exit 2 **mit Zeilenangabe**" galt
  für mehrere Config-Fehler (Version/Pflichtblöcke) nie — der neue Pfad vergrößerte sie; jetzt auf
  „wo die Fehlerquelle eine Zeile hat" abgeschwächt.
- **Mono-Repo-Nachbarschaft:** die Deklaration mehrerer Sprachen ist abgedeckt (`languages`-Map,
  je-Key-Validierung); die *Auflösung* pro Sprache (Go Modulpfad + TS relativ) bleibt offen und
  gehört in slice-015.
