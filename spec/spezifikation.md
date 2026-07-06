# Spezifikation — a-check

**Version:** 0.18.0

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
tech:                           # Tech-/Framework-Muster → zugeordnete(r) Adapter (optional)
  - {pattern: "net/http", adapter: http}
  - {pattern: "sqlite3*", adapter: persistence}
  - {pattern: "Q[A-Za-z]", adapter: ui, match: regex}  # RE2 statt Substring (Default: match: substring)
  - {pattern: "yaml", adapter: [config, report]}       # Pfad-LISTE: in jedem gelisteten Adapter erlaubt
  - {pattern: "net/http", adapter: http, composition_root: forbid}  # tech-leak auch in der Composition Root
exclude:                        # Datei-Globs: vor der Extraktion vom Scan ausgenommen (optional)
  - "**/*_test.go"
  - "**/node_modules/**"
composition_root: ["hexagon/main/**"]   # Ausnahme von Schicht-Regeln + tech-leak (letzteres je tech-Eintrag via composition_root: forbid abschaltbar; optional)
allow:                          # explizit erlaubte Sonderkanten/Re-Exports (optional)
  - {from: ports, to: ports, reason: "Re-Export"}
markers:                        # Heuristik-Grenze: Allowlist/Marker-Ausnahmen (optional)
  ignore_symbols: ["Queue.h"]
forbidden_constructs:           # Schicht → verbotene Text-Muster (Port-Disziplin, optional)
  ports: ["impl "]
resolution:                     # Symbol→Layer-Auflösung je Sprache (optional)
  go:     {mode: path}                          # Default (== weggelassen)
  cpp:    {mode: fixed-root, roots: ["src"]}
  kotlin: {mode: fixed-root, roots: ["src/main/kotlin"], package_base: "com.x"}
  typescript: {mode: relative}                  # datei-relativ (./x, ../y)
```

- **Pflichtblöcke:** `version`, `languages`, `layers`, `edges`. Die `languages`-Schlüssel müssen aus
  der Backend-Menge von [SPEC-EXTRACT-001](#spec-extract-001--import-extraktion) stammen; ein
  unbekannter Schlüssel → Exit 2 (die Menge steht **normativ nur** dort, hier bloß verwiesen —
  **kein Duplikat**).
- **Optionalblöcke:** `adapter_sink`, `tech`, `composition_root`, `allow`,
  `markers`, `forbidden_constructs`, `resolution`, `exclude`. Fehlt ein Optionalblock, entfällt die
  zugehörige Prüfung — nicht still, sondern bewusst nicht-konfiguriert. Die je
  Block präzisierte Anforderung:
  - `adapter_sink` → gemeinsame Senke aus [AC-FA-RULE-002](lastenheft.md#ac-fa-rule-002--keine-lateralen-adapter-kanten-regel-lateral-adapter); fehlt sie, darf **kein** Adapter einen anderen importieren (strengere Auslegung).
  - `tech` → [AC-FA-RULE-003](lastenheft.md#ac-fa-rule-003--tech-kapselung-regel-tech-leak); fehlt es, entfällt `tech-leak` (gedeckt durch die Boundary von [AC-FA-CONF-001](lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml)). `adapter` ist ein nicht-leeres Pfad-Fragment **oder** eine Pfad-**Liste** (das Symbol ist in **jedem** gelisteten Adapter erlaubt; nicht-leerer Skalar ≡ einelementige Liste, rückwärtskompatibel; eine **leere** Liste, ein leerer Listen-Eintrag, ein **leerer oder fehlender** Skalar → Exit 2 — der leere Adapter war vor 0.14.0 ein stiller Never-Leak-Eintrag, dieselbe fail-closed-Linie wie der leere `resolution`-Root). Der Adapter-Abgleich ist ein **Teilstring**-Vergleich auf dem Dateipfad, nicht segmentgrenzen-bewusst (dokumentierte Grenze, wie vor 0.14.0). Je Eintrag optional `match: substring|regex` (Default `substring`): `substring` = Teilstring-Vergleich, `regex` = **RE2**, unverankerter Suchlauf (`regexp.MatchString`) gegen das extrahierte Symbol ([SPEC-EXTRACT-001](#spec-extract-001--import-extraktion)); und optional `composition_root: allow|forbid` (Default `allow` = Composition-Root-Ausnahme wie bisher; `forbid` schaltet **nur** die `tech-leak`-Ausnahme der Composition Root für diesen Eintrag ab, [SPEC-RULE-001](#spec-rule-001--regel-auswertung)) — anderer Wert → Exit 2. Unbekannter `match`-Wert, leere oder nicht kompilierbare Regex → Exit 2 (strict-decode; ein leeres Regex-Muster würde jeden Import treffen und ist unzulässig). RE2 ist linear und deterministisch ([SPEC-DET-001](#spec-det-001--determinismus-vertrag)).
  - `composition_root` → deklarierte `tech-leak`-Ausnahme ([AC-FA-RULE-003](lastenheft.md#ac-fa-rule-003--tech-kapselung-regel-tech-leak) Boundary).
  - `allow` → konfigurativ erlaubte Sonderkante/Re-Export ([AC-FA-RULE-005](lastenheft.md#ac-fa-rule-005--schicht-richtung-regel-wrong-direction) / [AC-FA-RULE-004](lastenheft.md#ac-fa-rule-004--port-disziplin-regel-port-impurity) Boundary).
  - `markers` → dokumentierte Heuristik-Ausnahme ([AC-QA-02](lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)).
  - `forbidden_constructs` → schichtbezogen verbotene Konstrukte ([AC-FA-RULE-004](lastenheft.md#ac-fa-rule-004--port-disziplin-regel-port-impurity)); als Text-Muster geprüft (siehe [SPEC-EXTRACT-001](#spec-extract-001--import-extraktion)).
  - `resolution` → Symbol→Layer-Auflösung **je Sprache** (Map Sprache → `{mode, roots, package_base}`); `mode ∈ {path (Default), fixed-root, relative}`, `namespace` **reserviert** → Exit 2. `fixed-root`: `roots` vorangestellt; **bei gesetztem `package_base`** (gepunktete Sprache) zusätzlich Präfix-Strip + `.`→`/` (eine Pfad-Sprache wie C++ behält ihre `.`-Endungen); greift nur, wenn der Paket-Baum den Verzeichnis-Baum spiegelt ([AC-QA-02](lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)). `relative`: Specifier, die `.`/`..` sind oder mit `./`/`../` beginnen, werden lexikalisch gegen das **Verzeichnis der importierenden Datei** normalisiert (`path.Clean`-Semantik); alle anderen Specifier (Bare-Imports) sowie Wurzel-Escapes (führendes `..` **nach** der Normalisierung) liefern eine **leere** Kandidatenmenge — das Roh-Symbol wird nicht als Pfad-Kandidat weitergereicht (kein Geister-Match, [AC-QA-02](lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)); `roots`/`package_base` sind bei `relative` unzulässig → Exit 2; Endungs-Agnostik gilt, solange die `layers`-Glob-Präfixe oberhalb der Dateiebene enden (verzeichnisbasierte Globs). **Mehr-Wurzel-Auflösung (deklarations-bewusst, fail-closed):** `fixed-root` mit **≥ 2** `roots` löst den internen FQN gegen die **real gescannten Dateien** auf. Für ein **deklarations-bewusstes** Backend (Kotlin, [SPEC-EXTRACT-001](#spec-extract-001--import-extraktion)) gilt eine Evidenz-Rangfolge, stärkste zuerst: (1) **deklariert** — eine gescannte Datei im Paket-Verzeichnis `root/pkg/` trägt das Symbol als **Top-Level-Deklaration** (unabhängig vom Dateinamen); (2) **nur-Paketverzeichnis** — `root/pkg/` existiert, aber keine Datei deklariert das Symbol (eine gleichnamige, das Symbol **nicht** deklarierende Datei — und ein **Wildcard-/Paket-Import** `a.b.*` → `a/b/`, der ein ganzes Paket-Verzeichnis trifft — zählt nur hier); (3) **keine** (Phantom, extern). Die echte Deklaration **sticht** damit den bloßen Datei-Namens-Match. Für die übrigen (nicht deklarations-bewussten) Backends gilt unverändert der Datei-Namens-Match (endungs-agnostisch, package==directory: ein Wildcard-/Paket-Import — Symbol mit Trailing-Dot — trifft das Paket-Verzeichnis, ein Symbol dessen Datei mit gestrippter Endung oder sein Paket-Verzeichnis). Die Schicht wird am Pfad des **auflösenden** Kandidaten via [SPEC-RULE-001](#spec-rule-001--regel-auswertung) bestimmt, nicht am Wurzel-Präfix. Es entscheidet die **stärkste vorhandene Evidenzstufe**; auf ihr löst **genau ein** Root (oder alle dieselbe Schicht). ≥ 2 Roots **verschiedener** Schichten auf der stärksten Stufe: auf Stufe **deklariert** (bzw. für nicht deklarations-bewusste Backends: Datei-Match) ⇒ echte Mehrdeutigkeit, **Exit 2 nach dem Scan** (ein FQN muss in höchstens eine Schicht auflösen; `expect`/`actual` same-layer löst sauber); auf Stufe **nur-Paketverzeichnis** (kein Deklarations-Treffer) ⇒ **extern** (fail-open — ohne Deklaration nicht diskriminierbar). Die Stufe *nur-Paketverzeichnis* **löst** damit rückwärtskompatibel, solange sie eindeutig ist; ganz ohne Evidenz ⇒ extern. Grenzen ([AC-QA-02](lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)): ein intern gemeintes Symbol ohne Top-Level-Deklaration in gescanntem Code (verschachtelte Klasse, `object`-Member, generiert, Star-Import) bleibt still extern; die Deklarations-Auflösung ist **Kotlin-only** (übrige Backends: package==directory); datei-tiefe Globs sind eine heuristische Grenze. Fehlt der Block (oder eine Sprache) → Import-als-Pfad. Nutzt Sprache **und Pfad** der Quelldatei ([SPEC-RULE-001](#spec-rule-001--regel-auswertung)/[SPEC-EXTRACT-001](#spec-extract-001--import-extraktion)).
  - `exclude` → **Scan-Scope**: Datei-Globs relativ zur Scan-Wurzel (dieselbe Glob-Semantik wie `layers`/`languages`); eine matchende Datei wird **vor** der Extraktion vollständig vom Scan ausgenommen — sie existiert für keine Prüfung (weder Import- noch `forbidden_constructs`-Erkennung, [SPEC-EXTRACT-001](#spec-extract-001--import-extraktion)). Ein leerer Glob → Exit 2 (der ungültige Fall der ansonsten totalen Glob-Engine). Fehlt der Block, wird jede `languages`-Glob-Datei gescannt — byte-identisch zum bisherigen Verhalten ([AC-FA-CONF-001](lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml)).
- **Schicht-Rollen** ([AC-FA-RULE-006](lastenheft.md#ac-fa-rule-006--schicht-rollen-generische-regel-anwendung)): ein `layers`-Eintrag ist **entweder** eine Glob-Liste (`name: [globs]`) **oder** ein Objekt `{globs: [...], role: domain|app|port|adapter, direction: driving|driven}` (`direction` optional). Fehlt `role`, wird es aus konventionellen Namen abgeleitet (`core`→`domain`, `ports`→`port`, `adapters`→`adapter`, `application`/`app`→`app`); `role:` hat Vorrang. Die Reinheits-Regeln (`core-impurity`/`app-impurity`/`port-impurity`/`lateral-adapter`) greifen über die Rolle, nicht den Namen — fremd benannte Schichten sind damit voll prüfbar. Optional trägt eine `port`-/`adapter`-Schicht zusätzlich `direction` ∈ {`driving`, `driven`} (**orthogonal** zur Rolle, [AC-FA-RULE-008](lastenheft.md#ac-fa-rule-008--driving-driven-port-richtung-regel-port-direction-mismatch)); die Connectivity-Regel `port-direction-mismatch` prüft, dass ein Adapter nur Ports **seiner** Richtung importiert — ohne `direction` keine Prüfung.
- Kein Include/Vererbung zwischen Config-Dateien (Lastenheft-Out-of-Scope).

## SPEC-EXTRACT-001 — Import-Extraktion

Präzisiert [AC-FA-EXTRACT-001](lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion)
und [AC-QA-02](lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze).

Pro Datei, die einem Schicht-Glob entspricht, liefert das über `languages`
gewählte Backend die Menge der importierten Symbole/Module:

1. Dateien, die einem `exclude`-Glob entsprechen
   ([SPEC-CONF-001](#spec-conf-001--konfigurationsschema)), werden **vor** der
   Extraktion vollständig ausgenommen — sie werden nicht gelesen und liefern
   weder Import- noch `forbidden_constructs`-Treffer.
2. Die Datei wird zeilenweise gelesen.
3. Je Sprache werden konfigurierbare Muster angewandt (Defaults):
   - **C++:** `#include "…"` / `#include <…>`
   - **Go:** `import "…"` sowie Block-Form `import ( … )`
   - **Rust:** `use …;` und `extern crate …;` inkl. Alias-Form (`use x as y;` → `x`)
   - **Kotlin:** `import …`
   - **Java:** `import …;` inkl. `import static …;` (das `static` wird übersprungen, das `;` ignoriert)
   - **Python:** `import a.b.c` (inkl. Alias-Form `import a.b as x` → `a.b`) sowie
     `from a.b import c` → Modulpfad `a.b` (die importierten Namen werden nicht
     expandiert); **relative** Importe (führender Punkt: `from . import x`,
     `from ..pkg import y`) werden **nicht** gegriffen — eine dokumentierte
     Grenze der Python-Extraktion
     ([AC-QA-02](lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)),
     unabhängig vom `relative`-Auflösungs-Modus
     ([SPEC-CONF-001](#spec-conf-001--konfigurationsschema)), den
     Specifier-Sprachen wie TypeScript nutzen
   - **C#** (Schlüssel `csharp`): `using`-**Direktiven** → gepunkteter Namespace —
     `using A.B;` (inkl. `global using …;` und `using static …;`, Schlüsselwörter
     übersprungen) sowie Alias-Form `using X = A.B;` → `A.B` (das Ziel; der
     Alias-Name wird nie geliefert). Das `;` **direkt nach** dem gepunkteten Namen
     ist Teil des Musters — `using`-**Statements** (`using var x = …;`,
     `using (…)`, `using T x = …;`) matchen dadurch nie. Typ-Aliasse auf
     generische Typen, `extern alias`, `global::`-Qualifizierer und
     `namespace`-Deklarationen werden nicht gegriffen
   - **TypeScript** (Schlüssel `typescript`): ES-Module-Formen → **Modul-Specifier**
     (String in `'…'` **oder** `"…"`, gleichwertig; Semikolon optional/ASI, nicht
     Teil des Musters): `import … from '…'` (inkl. `import type` und
     Inline-`type`-Modifier), Seiteneffekt `import '…'`, Re-Export
     `export … from '…'` (inkl. `export * from`, `export * as ns from`,
     `export type … from`), Interop `import X = require('…')` sowie die
     Fortsetzungszeile `} from '…'` mehrzeilig umbrochener Imports/Re-Exports.
     Der Mittelteil zwischen `import`/`export` und `from` ist auf
     Import-Clause-Zeichen beschränkt (Bezeichner, `{ } * ,`, Whitespace — kein
     `=`, `(`, `.`, keine Quotes) — Ausdrucks-Zeilen (`export const q =
     knex.from('users')`, dynamisches `import(…)`/`require(…)`) matchen nie;
     Triple-Slash-Direktiven fallen dem Kommentar-Strip zu; Template-Literale
     und JSX-Textzeilen sind die bestehende String-Grenze
     ([AC-QA-02](lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)).
     Weitere dokumentierte Grenzen: Specifier mit `//` (URL-Importe) fallen dem
     Kommentar-Strip zu; kompakte Formen ohne Whitespace nach `import`/`export`
     sowie ein nacktes `from '…'` auf eigener (nicht `}`-geführter) Zeile
     werden nicht gegriffen
4. Import-ähnliche Zeilen in Zeilen-/Block-Kommentaren werden **nicht**
   gewertet (`//` und `/* */` werden entfernt). Das C-artige Kommentar-Stripping
   gilt **nur für die C-Syntax-Sprachen**; **Python** wird nicht C-gestrippt —
   seine `#`-Kommentarzeilen werden von den zeilen-verankerten Mustern nie
   gewertet, und eine `/*`-artige Bytefolge in einem Python-String-Literal
   (z. B. das Glob `"**/*.py"`) darf keine echten Imports verschlucken. Import-ähnliche Zeilen in
   **String-Literalen** sind eine **ausgewiesene Heuristik-Grenze** (0.1.0:
   reines Kommentar-Stripping, keine String-Awareness). Wo die Heuristik an
   ihre Grenze stößt (z. B. ein framework-fremdes `Queue.h` unter einem
   `Q[A-Za-z]`-Muster oder ein Treffer in einem String), wird die Grenze
   ausgewiesen, nicht verschwiegen; `markers.ignore_symbols` erlaubt eine
   dokumentierte Ausnahme.
5. Ergebnis je Datei: eine **deduplizierte, stabil sortierte** Symbolmenge
   (siehe [SPEC-DET-001](#spec-det-001--determinismus-vertrag)).

**Deklarations-Extraktion (deklarations-bewusste Backends):** zusätzlich zur Import-Menge
liefert ein deklarations-bewusstes Backend die Menge der **Top-Level-Deklarationsnamen** einer
Datei — für **Kotlin**: `fun` (inkl. Extension `fun R.name`, gewertet wird der Deklarationsname
`name`), `val`/`var`/`const val`, `class` (inkl. `data`/`sealed`/`enum`/`annotation class`),
`object`, `interface`, `sealed interface`, `typealias`; an **Spalte 0** verankert (eingerückte
Member zählen nicht) und **Kommentar-gestrippt** wie die Import-Muster (eine Deklarations-ähnliche
Zeile in einem Kommentar zählt **nicht**) — eine solche Zeile in einem **mehrzeiligen
String-Literal** bleibt die ausgewiesene Heuristik-Grenze
([AC-QA-02](lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)). Nur
**Top-Level**-Deklarationen (verschachtelte/Member-Deklarationen und Deklarationen in nicht
gescanntem/generiertem Code werden nicht indiziert). Die Menge — dedupliziert und stabil sortiert
([SPEC-DET-001](#spec-det-001--determinismus-vertrag)) — speist die **deklarations-bewusste
Mehr-Wurzel-Auflösung** ([SPEC-CONF-001](#spec-conf-001--konfigurationsschema)). In dieser Fassung
ist **Kotlin** das einzige deklarations-bewusste Backend; alle übrigen liefern ein **leeres**
Deklarations-Set (No-op, kein Verhaltenswechsel).

Nur direkte Imports (keine transitive Auflösung über Modulgrenzen);
Toolchain-gestützte Backends sind Lastenheft-Out-of-Scope.

**Zulässige Backend-Menge** (normativ, Owner dieser Spec): genau `{cpp, go, rust, kotlin, java, python, csharp, typescript}`.
Ein `languages`-Schlüssel außerhalb dieser Menge ist ein **Konfigurationsfehler** (Exit 2,
[SPEC-CLI-001](#spec-cli-001--aufruf-scan-wurzel-und-exit-codes)) — der Extraktions-Adapter
validiert die `languages`-Schlüssel gegen seine Backend-Registry, **bevor** er Dateien liest
(kein stiller No-Op / falsch-grün für nicht unterstützte Sprachen). Ein neues Backend erweitert
die Registry — die Menge lebt an **einer** Stelle.

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
| `lateral-adapter` | Datei mit Rolle `adapter` importiert eine *andere* `adapter`-Schicht (Layer-Identität) oder — in derselben Schicht — eine andere Adapter-Sub-Einheit (relativ zum Schicht-Glob-Präfix; bei einem Pfad-Rest **ohne weiteres Verzeichnis** entscheidet die Blatt-Klassifikation: datei-förmig (`.`) → **Root-Sub-Einheit `''`** — Sub-Einheiten sind Verzeichnisse, keine Dateinamen —, verzeichnis-förmig (z. B. Go-Paket-Pfad) → das Blatt **ist** die Sub-Einheit); nicht `adapter_sink`. Sub-Einheit **und** `adapter_sink`-Ausnahme werden auf dem gemäß `resolution` normalisierten Ziel-**Kandidaten** geprüft, nicht am Roh-Symbol (ein relativer Specifier trägt den Schicht-Präfix nie; im `path`-Modus sind beide identisch). **Kategorisch** (nicht über `edges`/`allow` aufhebbar) | [AC-FA-RULE-002](lastenheft.md#ac-fa-rule-002--keine-lateralen-adapter-kanten-regel-lateral-adapter) |
| `tech-leak` | ein `tech`-Muster (Substring oder RE2-Regex, je `match`) erscheint außerhalb **aller** seiner zugeordneten Adapter (`adapter` als Pfad oder Pfad-Liste) — und außerhalb `composition_root`, sofern der Eintrag nicht `composition_root: forbid` deklariert (dann prüft `tech-leak` auch dort; die Meldung nennt alle gelisteten Adapter in Deklarationsreihenfolge) | [AC-FA-RULE-003](lastenheft.md#ac-fa-rule-003--tech-kapselung-regel-tech-leak) |
| `port-impurity` | Datei mit Rolle `port` importiert eine `adapter`-Rolle oder ein `tech`-Muster **oder** enthält ein `forbidden_constructs`-Muster (text-heuristisch erkannt). **Kern-Referenzen sind erlaubt** (Ports sprechen die Sprache des Kerns) und werden über `edges`/`allow` regiert — eine undeklarierte `ports → core`-Kante fällt unter `wrong-direction` | [AC-FA-RULE-004](lastenheft.md#ac-fa-rule-004--port-disziplin-regel-port-impurity) |
| `port-direction-mismatch` | Datei mit Rolle `adapter` und Richtung `direction` X importiert eine `port`-Rolle mit Richtung Y (X ≠ Y, **beide gesetzt**) — ein Treiber-Adapter spricht nur `driving`-Ports, ein getriebener nur `driven`-Ports; **orthogonal** zur Rolle, ohne `direction` keine Prüfung. **Kategorisch** (nicht über `edges`/`allow` aufhebbar, wie `lateral-adapter`) | [AC-FA-RULE-008](lastenheft.md#ac-fa-rule-008--driving-driven-port-richtung-regel-port-direction-mismatch) |
| `wrong-direction` | ein Import quert eine Schicht-Kante entgegen `edges`/`allow` | [AC-FA-RULE-005](lastenheft.md#ac-fa-rule-005--schicht-richtung-regel-wrong-direction) |

Die Schicht einer Datei ergibt sich aus dem **spezifischsten** passenden `layers`-Glob
(längster **literaler** Präfix vor dem ersten Wildcard-Segment, konsistent mit der
Symbol-Auflösung unten; bei Gleichstand die zuerst deklarierte Schicht), ihre
**Rolle** aus `role:` (Vorrang) oder Namens-Inferenz ([AC-FA-RULE-006](lastenheft.md#ac-fa-rule-006--schicht-rollen-generische-regel-anwendung)).
Vor der Auflösung wird ein Import-Symbol gemäß dem `resolution`-`mode` seiner **Quelldatei-Sprache**
normalisiert — die Auflösung kennt dazu Sprache **und Pfad** der importierenden Datei
(`fixed-root`: `roots` voran; bei gesetztem `package_base` zusätzlich Präfix-Strip + `.`→`/`;
`relative`: relative Specifier (`.`/`..`/`./…`/`../…`) lexikalisch gegen das Verzeichnis der
importierenden Datei, nicht-relative Specifier und Wurzel-Escapes → **leere** Kandidatenmenge;
`path`/Default: unverändert) — so lösen gepunktete (JVM/Python/C#), `src`-gewurzelte (C++) oder
datei-relative (TypeScript) Importe auf, sofern der Paket-Baum den Verzeichnis-Baum spiegelt
bzw. die `layers`-Globs verzeichnisbasiert sind. Dann werden Symbole über die `layers`-Globs des Zielpfads aufgelöst
(**spezifischster/längster** literaler Präfix gewinnt) — die Ziel-Rolle ist die **des aufgelösten Layers**; die Reinheits-Regeln
dispatchen über die Rolle, nicht den Namen. Die **`tech`-Muster** dagegen lösen in
**Deklarationsreihenfolge** auf (**Erst-Treffer** gewinnt, `matchTech`) — uniform für
`match: substring` und `match: regex`; es gibt für `tech` **kein** „längster Präfix"
([AC-FA-RULE-003](lastenheft.md#ac-fa-rule-003--tech-kapselung-regel-tech-leak)).

Pro (Datei, Import) gilt **deterministische Erst-Treffer-Reihenfolge** in der
Tabellen-Reihenfolge (`core-impurity` → `app-impurity` → `port-impurity` →
`lateral-adapter` → `tech-leak` → `port-direction-mismatch` → `wrong-direction`); ein Import erzeugt höchstens einen Befund.
Dateien unter `composition_root` sind als Verdrahtungspunkt von **allen**
Schicht-Regeln ausgenommen — sie importieren bestimmungsgemäß quer über die
Schichten — und von `tech-leak` **je `tech`-Eintrag**: bei
`composition_root: allow` (Default) wie bisher, bei `composition_root: forbid`
prüft `tech-leak` den Eintrag auch dort weiter
([SPEC-CONF-001](#spec-conf-001--konfigurationsschema)); die
Schicht-Regel-Ausnahme bleibt davon unberührt. `exclude`-Dateien erreichen die
Regel-Auswertung nie (Scan-Scope, [SPEC-EXTRACT-001](#spec-extract-001--import-extraktion)).

## SPEC-CLI-001 — Aufruf, Scan-Wurzel und Exit-Codes

Präzisiert [AC-FA-CLI-001](lastenheft.md#ac-fa-cli-001--aufruf-scan-wurzel-und-exit-codes).

- Aufruf: `a-check [pfad]`; Default-Scan-Wurzel `/src` (Container-Mount).
- `.a-check.yml` wird aus der Scan-Wurzel gelesen.
- **Exit-Codes:** `0` kein Befund · `1` ≥ 1 Befund · `2` Nutzungs-/
  Konfigurationsfehler (fehlende/ungültige Config, unbekanntes Flag); eine
  ungültige Config wird gemeldet — **mit Zeilenangabe, wo die Fehlerquelle eine
  Zeile hat** (Schema-/Sprach-Validierungen ohne Positionsbezug melden ohne Zeile).
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
| 0.7.0 | 2026-06-23 | `SPEC-EXTRACT-001`: Java-Muster (`import …;`, inkl. `import static …;` — `static` übersprungen) als fünftes Backend. Folgt [`AC-FA-EXTRACT-001`](lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion) 0.7.0. |
| 0.8.0 | 2026-07-01 | `SPEC-CONF-001`: `tech`-Eintrag um optionales `match: substring\|regex` (Default `substring`; `regex` = RE2, unverankert, gegen das extrahierte Symbol); unbekannter Wert/nicht kompilierbare Regex → Exit 2. `SPEC-RULE-001`: Präzedenz der `tech`-Muster als **Deklarationsreihenfolge/Erst-Treffer** richtiggestellt (kein „längster Präfix" für `tech` — das gilt nur für `layers`-Globs; `matchTech` liefert real schon Erst-Treffer). Folgt [`AC-FA-RULE-003`](lastenheft.md#ac-fa-rule-003--tech-kapselung-regel-tech-leak)/[`AC-FA-CONF-001`](lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml) 0.8.0. |
| 0.9.0 | 2026-07-01 | `SPEC-EXTRACT-001` nennt die zulässige Backend-Menge `{cpp,go,rust,kotlin,java}` **normativ** (Owner) + Validierung der `languages`-Schlüssel gegen die Backend-Registry (unbekannt → Exit 2, kein stiller No-Op); `SPEC-CONF-001` **verweist** darauf (kein Duplikat). Folgt [`AC-FA-CONF-001`](lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml) 0.9.0. |
| 0.10.0 | 2026-07-01 | `SPEC-CONF-001`: `resolution`-Block (Map Sprache → `{mode, roots, package_base}`; `mode ∈ {path, fixed-root}`, `relative`/`namespace` reserviert → Exit 2). `SPEC-RULE-001`: Import-Symbol wird **vor** der Layer-Auflösung gemäß dem `mode` seiner **Quelldatei-Sprache** normalisiert (`fixed-root`: `roots` voran, bei gesetztem `package_base` zusätzlich Strip + `.`→`/`) — Grenze Paket==Verzeichnis. Folgt [`AC-FA-CONF-001`](lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml) 0.10.0. |
| 0.11.0 | 2026-07-02 | `SPEC-EXTRACT-001`: Python-Muster (`import a.b.c` inkl. Alias, `from a.b import c` → Modulpfad; relative Importe nicht gegriffen — reservierter `relative`-Modus) als sechstes Backend; Backend-Menge → `{cpp,go,rust,kotlin,java,python}`. Folgt [`AC-FA-EXTRACT-001`](lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion) 0.11.0. |
| 0.12.0 | 2026-07-02 | `SPEC-EXTRACT-001`: C#-Muster (`using`-Direktiven inkl. `global`/`static`/Alias-Ziel; Pflicht-`;` nach dem Namen schließt `using`-Statements aus) als siebtes Backend; Backend-Menge → `{cpp,go,rust,kotlin,java,python,csharp}`. Folgt [`AC-FA-EXTRACT-001`](lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion) 0.12.0. |
| 0.13.0 | 2026-07-03 | `SPEC-EXTRACT-001`: TypeScript-Muster (ES-Module-Formen → Modul-Specifier in `'…'`/`"…"`, Semikolon optional; `import … from` inkl. `type`, Seiteneffekt, Re-Exports, `import X = require(…)`, Fortsetzungszeile `} from '…'`; Mittelteil auf Import-Clause-Zeichen beschränkt) als achtes Backend; Backend-Menge → `{cpp,go,rust,kotlin,java,python,csharp,typescript}`. Folgt [`AC-FA-EXTRACT-001`](lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion) 0.13.0. |
| 0.13.0 | 2026-07-03 | `SPEC-CONF-001`/`SPEC-RULE-001`: `mode: relative` gültig (Specifier `.`/`..`/`./…`/`../…` lexikalisch gegen das Verzeichnis der importierenden Datei; Bare-Imports/Wurzel-Escape → **leere** Kandidatenmenge; `roots`/`package_base` unzulässig → Exit 2; nur `namespace` reserviert); die Symbol-Normalisierung kennt Sprache **und Pfad** der Quelldatei (Quellpfad-Threading). Folgt [`AC-FA-CONF-001`](lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml) 0.13.0. |
| 0.14.0 | 2026-07-03 | `SPEC-CONF-001`/`SPEC-RULE-001`: `tech.adapter` auch als Pfad-**Liste** (Symbol in jedem gelisteten Adapter erlaubt; leere Liste/leerer Eintrag → Exit 2; Meldung nennt alle Adapter) + `tech.composition_root: allow\|forbid` je Eintrag (Default `allow`; `forbid` schaltet nur die `tech-leak`-Ausnahme der Composition Root ab, Schicht-Regel-Ausnahme unberührt). Folgt [`AC-FA-RULE-003`](lastenheft.md#ac-fa-rule-003--tech-kapselung-regel-tech-leak)/[`AC-FA-CONF-001`](lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml) 0.14.0. |
| 0.14.0 | 2026-07-03 | `SPEC-CONF-001`/`SPEC-EXTRACT-001`: optionaler **`exclude`**-Block (Datei-Globs relativ zur Scan-Wurzel) nimmt matchende Dateien **vor** der Extraktion vollständig vom Scan aus (auch `forbidden_constructs`); leerer Glob → Exit 2; ohne Block byte-identisch. Folgt [`AC-FA-CONF-001`](lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml) 0.14.0. |
| 0.15.0 | 2026-07-03 | `SPEC-RULE-001`: `lateral-adapter`-Sub-Einheiten — Blatt-Klassifikation für Pfad-Reste ohne weiteres Verzeichnis: datei-förmig (`.`) → **Root-Sub-Einheit `''`** (Verzeichnisse, keine Dateinamen), verzeichnis-förmig (Go-Paket-Pfad) → das Blatt ist die Sub-Einheit; Root↔Root same-layer kein Befund mehr, Root↔Unterverzeichnis/Cross-Paket/Cross-Layer unverändert. Folgt [`AC-FA-RULE-002`](lastenheft.md#ac-fa-rule-002--keine-lateralen-adapter-kanten-regel-lateral-adapter) 0.15.0. |
| 0.16.0 | 2026-07-04 | `SPEC-CONF-001`: **Mehr-Wurzel-Guard (fail-closed)** — `mode: fixed-root` mit ≥ 2 `roots`, von denen zwei Roots je eine andere Schicht erzwingen (längster passender Glob-Präfix am Root, `segIndex >= 0`, wie die Import-Auflösung), bricht mit Exit 2 statt still Phantom-Kandidaten fehlzuklassifizieren. Folgt [`AC-FA-CONF-001`](lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml) 0.16.0. |
| 0.17.0 | 2026-07-05 | `SPEC-CONF-001`: **datei-mengen-bewusste Mehr-Wurzel-Auflösung** (Stufe 2) — `fixed-root` mit ≥ 2 `roots` löst den FQN gegen die real gescannten Dateien auf (endungs-agnostisch, package==directory); Schicht am realen Kandidaten-Pfad, Phantom bleibt extern; der Ladezeit-Guard aus 0.16.0 entfällt. Gleicher FQN real in ≥ 2 Roots + **verschiedene** Schichten → Exit 2 **nach dem Scan** (distinct-layer; `expect`/`actual` same-layer löst sauber). Folgt [`AC-FA-CONF-001`](lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml) 0.17.0. |
| 0.18.0 | 2026-07-06 | `SPEC-CONF-001`/`SPEC-EXTRACT-001`: **deklarations-bewusste Mehr-Wurzel-Auflösung** (Stufe 3) — bei `fixed-root` mit ≥ 2 `roots` gewinnt für ein deklarations-bewusstes Backend (Kotlin) die **reale Top-Level-Deklaration** über den bloßen Datei-Namens-Match (Evidenz-Rangfolge deklariert > Paketverzeichnis > keine); genau ein deklarierender Root ⇒ eindeutig, ≥ 2 deklarierende Roots verschiedener Schichten ⇒ Exit 2, kein Treffer ⇒ extern (fail-open). `SPEC-EXTRACT-001`: **Kotlin** liefert zusätzlich Top-Level-Deklarationen (`fun`/Extension/`val`/`class`/`object`/`interface`/`typealias`), übrige Backends no-op (leeres Set). Folgt [`AC-FA-CONF-001`](lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml)/[`AC-FA-EXTRACT-001`](lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion) 0.18.0. |
