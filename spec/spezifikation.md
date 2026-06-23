# Spezifikation — a-check

**Version:** 0.6.0

**Status:** Draft

**Stratum:** Technik (verbindlich, fortschreibbar; ADR-Schärfung erlaubt)

**Autor:** pt9912, **Datum:** 2026-06-21.

---

## Zweck und Einordnung

Dieses Dokument ist das **Technik-Stratum**. Es *präzisiert* die
Anforderungen des [Lastenhefts](lastenheft.md) (Vertrag) — es **erweitert
sie nie**; bei Konflikt sticht das Lastenheft. Es ist **sprachneutral und
meilensteinfrei**: die sprachkonkrete Übersetzung und die Begründungen
leben in den ADRs, nicht hier.

`SPEC-<BEREICH>-<NNN>`-Kennungen präzisieren je eine Lastenheft-Anforderung.
Bereiche: `CONF`, `EXTRACT`, `RULE`, `CLI`, `DET`, `DIST`.

## SPEC-CONF-001 — Konfigurationsschema

Präzisiert [AC-FA-CONF-001](lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml).

`.a-check.yml` wird **strikt** dekodiert: jeder unbekannte Schlüssel und
jeder Typfehler ist ein Konfigurationsfehler (Exit-Code 2, siehe
[SPEC-CLI-001](#spec-cli-001--aufruf-scan-wurzel-und-exit-codes)) — kein
stiller Default. Top-Level-Schlüssel:

```yaml
version: 1                      # Schema-Version (Pflicht; unbekannte Version → Exit 2)
languages:                      # Sprache → Datei-Globs (wählt das Extraktions-Backend)
  go:    ["**/*.go"]
  cpp:   ["**/*.h", "**/*.hpp", "**/*.cpp"]
layers:                         # Schicht → Pfad-Muster (Globs, repo-relativ)
  core:     ["hexagon/core/**"]
  ports:    ["hexagon/ports/**"]
  adapters: ["hexagon/adapters/**"]
edges:                          # erlaubte gerichtete Schicht-Kanten (from → to)
  - {from: adapters, to: ports}
  - {from: ports,    to: core}  # Ports dürfen Domänentypen referenzieren (AC-FA-RULE-004)
adapter_sink: driver-common     # gemeinsame Senke, die Adapter importieren dürfen (optional)
tech:                           # Tech-/Framework-Muster → zugeordneter Adapter (optional)
  - {pattern: "net/http", adapter: http}
  - {pattern: "sqlite3*", adapter: persistence}
composition_root: ["hexagon/main/**"]   # deklarierte Ausnahme für tech-leak (optional)
allow:                          # explizit erlaubte Sonderkanten/Re-Exports (optional)
  - {from: ports, to: ports, reason: "Re-Export"}
markers:                        # Heuristik-Grenze: Allowlist/Marker-Ausnahmen (optional)
  ignore_symbols: ["Queue.h"]
forbidden_constructs:           # Schicht → verbotene Text-Muster (Port-Disziplin, optional)
  ports: ["impl "]
```

- **Pflichtblöcke:** `version`, `languages`, `layers`, `edges`.
- **Optionalblöcke:** `adapter_sink`, `tech`, `composition_root`, `allow`,
  `markers`, `forbidden_constructs`. Fehlt ein Optionalblock, entfällt die
  zugehörige Prüfung — nicht still, sondern bewusst nicht-konfiguriert. Die je
  Block präzisierte Anforderung:
  - `adapter_sink` → gemeinsame Senke aus [AC-FA-RULE-002](lastenheft.md#ac-fa-rule-002--keine-lateralen-adapter-kanten-regel-lateral-adapter); fehlt sie, darf **kein** Adapter einen anderen importieren (strengere Auslegung).
  - `tech` → [AC-FA-RULE-003](lastenheft.md#ac-fa-rule-003--tech-kapselung-regel-tech-leak); fehlt es, entfällt `tech-leak` (gedeckt durch die Boundary von [AC-FA-CONF-001](lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml)).
  - `composition_root` → deklarierte `tech-leak`-Ausnahme ([AC-FA-RULE-003](lastenheft.md#ac-fa-rule-003--tech-kapselung-regel-tech-leak) Boundary).
  - `allow` → konfigurativ erlaubte Sonderkante/Re-Export ([AC-FA-RULE-005](lastenheft.md#ac-fa-rule-005--schicht-richtung-regel-wrong-direction) / [AC-FA-RULE-004](lastenheft.md#ac-fa-rule-004--port-disziplin-regel-port-impurity) Boundary).
  - `markers` → dokumentierte Heuristik-Ausnahme ([AC-QA-02](lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)).
  - `forbidden_constructs` → schichtbezogen verbotene Konstrukte ([AC-FA-RULE-004](lastenheft.md#ac-fa-rule-004--port-disziplin-regel-port-impurity)); als Text-Muster geprüft (siehe [SPEC-EXTRACT-001](#spec-extract-001--import-extraktion)).
- **Schicht-Rollen** ([AC-FA-RULE-006](lastenheft.md#ac-fa-rule-006--schicht-rollen-generische-regel-anwendung)): ein `layers`-Eintrag ist **entweder** eine Glob-Liste (`name: [globs]`) **oder** ein Objekt `{globs: [...], role: domain|app|port|adapter, direction: driving|driven}` (`direction` optional). Fehlt `role`, wird es aus konventionellen Namen abgeleitet (`core`→`domain`, `ports`→`port`, `adapters`→`adapter`, `application`/`app`→`app`); `role:` hat Vorrang. Die Reinheits-Regeln (`core-impurity`/`app-impurity`/`port-impurity`/`lateral-adapter`) greifen über die Rolle, nicht den Namen — fremd benannte Schichten sind damit voll prüfbar. Optional trägt eine `port`-/`adapter`-Schicht zusätzlich `direction` ∈ {`driving`, `driven`} (**orthogonal** zur Rolle, [AC-FA-RULE-008](lastenheft.md#ac-fa-rule-008--driving-driven-port-richtung-regel-port-direction-mismatch)); die Connectivity-Regel `port-direction-mismatch` prüft, dass ein Adapter nur Ports **seiner** Richtung importiert — ohne `direction` keine Prüfung.
- Kein Include/Vererbung zwischen Config-Dateien (Lastenheft-Out-of-Scope).

## SPEC-EXTRACT-001 — Import-Extraktion

Präzisiert [AC-FA-EXTRACT-001](lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion)
und [AC-QA-02](lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze).

Pro Datei, die einem Schicht-Glob entspricht, liefert das über `languages`
gewählte Backend die Menge der importierten Symbole/Module:

1. Die Datei wird zeilenweise gelesen.
2. Je Sprache werden konfigurierbare Muster angewandt (Defaults):
   - **C++:** `#include "…"` / `#include <…>`
   - **Go:** `import "…"` sowie Block-Form `import ( … )`
   - **Rust:** `use …;` und `extern crate …;` inkl. Alias-Form (`use x as y;` → `x`)
   - **Kotlin:** `import …`
3. Import-ähnliche Zeilen in Zeilen-/Block-Kommentaren werden **nicht**
   gewertet (`//` und `/* */` werden entfernt). Import-ähnliche Zeilen in
   **String-Literalen** sind eine **ausgewiesene Heuristik-Grenze** (0.1.0:
   reines Kommentar-Stripping, keine String-Awareness). Wo die Heuristik an
   ihre Grenze stößt (z. B. ein framework-fremdes `Queue.h` unter einem
   `Q[A-Za-z]`-Muster oder ein Treffer in einem String), wird die Grenze
   ausgewiesen, nicht verschwiegen; `markers.ignore_symbols` erlaubt eine
   dokumentierte Ausnahme.
4. Ergebnis je Datei: eine **deduplizierte, stabil sortierte** Symbolmenge
   (siehe [SPEC-DET-001](#spec-det-001--determinismus-vertrag)).

Nur direkte Imports (keine transitive Auflösung über Modulgrenzen);
Toolchain-gestützte Backends sind Lastenheft-Out-of-Scope.

Neben Importen erkennt das Backend optionale `forbidden_constructs`-Muster
([SPEC-CONF-001](#spec-conf-001--konfigurationsschema)) text-heuristisch je
Schicht — dieselbe Muster-Mechanik, anderer Treffertyp (Sprachkonstrukt statt
Import); sie speist die `port-impurity`-Regel
([SPEC-RULE-001](#spec-rule-001--regel-auswertung)).

## SPEC-RULE-001 — Regel-Auswertung

Präzisiert die sieben Hexagon-Regeln `AC-FA-RULE-*`; ihre Anwendung über
**Layer-Rollen** statt Namen regelt [AC-FA-RULE-006](lastenheft.md#ac-fa-rule-006--schicht-rollen-generische-regel-anwendung). Eingabe: die
Symbolmengen je Datei ([SPEC-EXTRACT-001](#spec-extract-001--import-extraktion))
und das Schicht-/Kanten-/Tech-Modell ([SPEC-CONF-001](#spec-conf-001--konfigurationsschema)).
Jede verletzende Datei erzeugt einen Befund (Datei, Zeile, Regelname,
Meldung); ≥ 1 Befund ⇒ Exit-Code 1.

| Regelname | Auswertung | präzisiert |
|---|---|---|
| `core-impurity` | Datei mit Rolle `domain` importiert ein Symbol, das auf eine `app`-, `port`- oder `adapter`-Rolle oder ein `tech`-Muster auflöst — `domain` ist die innerste Schicht, **kategorisch** | [AC-FA-RULE-001](lastenheft.md#ac-fa-rule-001--kern-reinheit-regel-core-impurity) |
| `app-impurity` | Datei mit Rolle `app` importiert eine `adapter`-Rolle oder ein `tech`-Muster; `domain`- und `port`-Referenzen sind erlaubt (Richtung edge-regiert) | [AC-FA-RULE-007](lastenheft.md#ac-fa-rule-007--rolle-app-und-strenge-domain) |
| `lateral-adapter` | Datei mit Rolle `adapter` importiert eine *andere* `adapter`-Schicht (Layer-Identität) oder — in derselben Schicht — eine andere Adapter-Sub-Einheit (relativ zum Schicht-Glob-Präfix); nicht `adapter_sink`. **Kategorisch** (nicht über `edges`/`allow` aufhebbar) | [AC-FA-RULE-002](lastenheft.md#ac-fa-rule-002--keine-lateralen-adapter-kanten-regel-lateral-adapter) |
| `tech-leak` | ein `tech`-Muster erscheint außerhalb seines zugeordneten Adapters (und außerhalb `composition_root`, falls konfiguriert) | [AC-FA-RULE-003](lastenheft.md#ac-fa-rule-003--tech-kapselung-regel-tech-leak) |
| `port-impurity` | Datei mit Rolle `port` importiert eine `adapter`-Rolle oder ein `tech`-Muster **oder** enthält ein `forbidden_constructs`-Muster (text-heuristisch erkannt). **Kern-Referenzen sind erlaubt** (Ports sprechen die Sprache des Kerns) und werden über `edges`/`allow` regiert — eine undeklarierte `ports → core`-Kante fällt unter `wrong-direction` | [AC-FA-RULE-004](lastenheft.md#ac-fa-rule-004--port-disziplin-regel-port-impurity) |
| `port-direction-mismatch` | Datei mit Rolle `adapter` und Richtung `direction` X importiert eine `port`-Rolle mit Richtung Y (X ≠ Y, **beide gesetzt**) — ein Treiber-Adapter spricht nur `driving`-Ports, ein getriebener nur `driven`-Ports; **orthogonal** zur Rolle, ohne `direction` keine Prüfung. **Kategorisch** (nicht über `edges`/`allow` aufhebbar, wie `lateral-adapter`) | [AC-FA-RULE-008](lastenheft.md#ac-fa-rule-008--driving-driven-port-richtung-regel-port-direction-mismatch) |
| `wrong-direction` | ein Import quert eine Schicht-Kante entgegen `edges`/`allow` | [AC-FA-RULE-005](lastenheft.md#ac-fa-rule-005--schicht-richtung-regel-wrong-direction) |

Die Schicht einer Datei ergibt sich aus dem **spezifischsten** passenden `layers`-Glob
(längster **literaler** Präfix vor dem ersten Wildcard-Segment, konsistent mit der
Symbol-Auflösung unten; bei Gleichstand die zuerst deklarierte Schicht), ihre
**Rolle** aus `role:` (Vorrang) oder Namens-Inferenz ([AC-FA-RULE-006](lastenheft.md#ac-fa-rule-006--schicht-rollen-generische-regel-anwendung)).
Symbole werden über die `layers`-Globs des Zielpfads bzw. die `tech`-Muster
aufgelöst (**spezifischster/längster** Präfix gewinnt) — die Ziel-Rolle ist die **des aufgelösten Layers**; die Reinheits-Regeln
dispatchen über die Rolle, nicht den Namen.

Pro (Datei, Import) gilt **deterministische Erst-Treffer-Reihenfolge** in der
Tabellen-Reihenfolge (`core-impurity` → `app-impurity` → `port-impurity` →
`lateral-adapter` → `tech-leak` → `port-direction-mismatch` → `wrong-direction`); ein Import erzeugt höchstens einen Befund.
Dateien unter `composition_root` sind als Verdrahtungspunkt von **allen**
Schicht-Regeln **und** `tech-leak` ausgenommen — sie importieren
bestimmungsgemäß quer über die Schichten.

## SPEC-CLI-001 — Aufruf, Scan-Wurzel und Exit-Codes

Präzisiert [AC-FA-CLI-001](lastenheft.md#ac-fa-cli-001--aufruf-scan-wurzel-und-exit-codes).

- Aufruf: `a-check [pfad]`; Default-Scan-Wurzel `/src` (Container-Mount).
- `.a-check.yml` wird aus der Scan-Wurzel gelesen.
- **Exit-Codes:** `0` kein Befund · `1` ≥ 1 Befund · `2` Nutzungs-/
  Konfigurationsfehler (fehlende/ungültige Config, unbekanntes Flag); eine
  ungültige Config wird **mit Zeilenangabe** gemeldet.
- **Befunde** auf stdout, ein Datensatz je Zeile im Format
  `pfad:zeile: regelname: meldung`; **Zusammenfassung** (Anzahl je Regel,
  Gesamtzahl) auf stderr.
- **Read-only:** der geprüfte Baum wird nie beschrieben (Mount `:ro`).

## SPEC-DET-001 — Determinismus-Vertrag

Präzisiert [AC-QA-01](lastenheft.md#ac-qa-01--determinismus).

Identische Eingabe (Repo-Stand + `.a-check.yml` + Image-Digest) ⇒
**byte-identische** Ausgabe und identischer Exit-Code. Befunde werden nach
einer Totalordnung sortiert: `pfad`, dann `zeile`, dann `regelname`.
Extraktions-Symbolmengen werden stabil sortiert. Keine Zeitstempel,
Zufalls- oder locale-abhängige Reihenfolgen in der Ausgabe.

## SPEC-DIST-001 — Laufzeitform und Distribution

Präzisiert [AC-FA-DIST-001](lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk),
[AC-QA-02](lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)
und [AC-QA-03](lastenheft.md#ac-qa-03--reproduzierbarkeit).

- **Laufzeitform:** ein einzelnes, in sich geschlossenes, **statisch
  gelinktes** Artefakt auf einem **distroless/static** Basis-Image; läuft
  **netzlos** (`--network none`) mit read-only gemountetem Prüfbaum. (Diese
  Form ist sprachneutral spezifiziert; die Implementierungssprache, die sie
  realisiert, ist eine ADR-Entscheidung.)
- **Image** ist `@sha256:`-digest-gepinnt; Pin-Hebung ist ein bewusster
  Commit.
- `--print-config`: gibt ein **kommentiertes** `.a-check.yml`-Gerüst auf
  stdout aus; schreibt nichts.
- `--print-mk`: gibt ein include-bares Makefile-Fragment auf stdout aus —
  mit digest-gepinntem `A_CHECK_IMAGE` und einem `a-check`-Target; schreibt
  nichts.
- Ein unbekanntes Flag ⇒ Exit-Code 2.

## Historie

| Version | Datum | Änderung |
|---|---|---|
| 0.1.0 | 2026-06-21 | Erstfassung (Technik-Stratum): `SPEC-CONF/EXTRACT/RULE/CLI/DET/DIST-001` präzisieren die Lastenheft-Verträge (Config-Schema, Extraktions-Algorithmus, Regel-Auswertung, CLI/Exit-Codes, Determinismus, Laufzeit-/Distributionsform). Sprachneutral. |
| 0.2.0 | 2026-06-22 | `SPEC-RULE-001` `port-impurity` nachgezogen: Port-Befund bei Adapter-/`tech`-Import statt bei Kern-Import; `ports → core` ist edge-regiert. Folgt [`AC-FA-RULE-004`](lastenheft.md#ac-fa-rule-004--port-disziplin-regel-port-impurity) 0.2.0. |
| 0.3.0 | 2026-06-22 | `SPEC-RULE-001`/`SPEC-CONF-001` rollen-basiert: die Reinheits-Regeln dispatchen über eine Layer-Rolle (`domain`/`port`/`adapter`, aus `role:` oder Namens-Inferenz); `lateral-adapter` cross-layer + kategorisch; `layers`-Eintrag als Glob-Liste oder `{globs, role}`. Folgt [`AC-FA-RULE-006`](lastenheft.md#ac-fa-rule-006--schicht-rollen-generische-regel-anwendung) 0.3.0. |
| 0.4.0 | 2026-06-22 | `SPEC-RULE-001`: `adapterSeg` layer-relativ (Adapter-Sub-Einheit nach dem Schicht-Glob-Präfix, namensunabhängig) + `targetLayer` längster-Präfix-Auflösung. Folgt [`AC-FA-RULE-006`](lastenheft.md#ac-fa-rule-006--schicht-rollen-generische-regel-anwendung) 0.4.0. |
| 0.5.0 | 2026-06-22 | `SPEC-RULE-001`: neue Rolle `app` (Befund `app-impurity` bei Adapter-/`tech`-Import) + `core-impurity` verschärft (`domain` importiert nur `domain`, kategorisch); Schema-Enum `role` um `app`. Folgt [`AC-FA-RULE-007`](lastenheft.md#ac-fa-rule-007--rolle-app-und-strenge-domain) 0.5.0. |
| 0.6.0 | 2026-06-23 | `SPEC-RULE-001`: neue Regel `port-direction-mismatch` (Adapter-Richtung ≠ Ziel-Port-Richtung, beide gesetzt; in der Erst-Treffer-Kette vor `wrong-direction`) + Schicht-Zuordnung einer Datei auf **spezifischsten/längsten** Glob-Präfix umgestellt (Angleichung an `targetLayer`); `SPEC-CONF-001`-Schema: Objekt-Form um `direction` (und das fehlende `app`) ergänzt. Folgt [`AC-FA-RULE-008`](lastenheft.md#ac-fa-rule-008--driving-driven-port-richtung-regel-port-direction-mismatch) 0.6.0. |
