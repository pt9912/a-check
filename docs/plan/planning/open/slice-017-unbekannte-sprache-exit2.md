# slice-017 — Unbekannter `languages`-Schlüssel → Exit 2 (falsch-grün-Falle schließen)

**Status:** open (Entwurf zur Abnahme).
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

- [ ] Lastenheft 0.9.0 (neue Negative-AC im G/W/T-Stil + Historie) + Spezifikation 0.9.0
  (Backend-Menge-Owner + Verweis statt Kopie + Historie) — Details §3.1/§3.2.
- [ ] Code: Backend-Registry-Map (Single Source), Extract-Validierung, kein `default: return nil`;
  Meldung aus den Registry-Keys.
- [ ] Tests: unbekannt → Exit 2 (+ Meldung), unterstützte unverändert, Registry == `{cpp,go,rust,kotlin,java}`.
- [ ] `make gates` + `make ci` grün; Multi-Linsen-Review (proportional); Merge auf Wort.
- [ ] **Nicht** beim Python-Slice als „gedeckt" abräumen (§1: unabhängige Härtung).
