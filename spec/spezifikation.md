# Spezifikation — a-check

**Version:** 0.31.0

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
constructs:                     # Roh-Text-Monopol: Muster → erlaubte Zone(n) (optional)
  - {pattern: 'dlopen\s*\(', match: regex, adapter: adapters/plugin, composition_root: forbid}
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
  `markers`, `forbidden_constructs`, `constructs`, `resolution`, `exclude`. Fehlt ein Optionalblock, entfällt die
  zugehörige Prüfung — nicht still, sondern bewusst nicht-konfiguriert. Die je
  Block präzisierte Anforderung:
  - `adapter_sink` → gemeinsame Senke aus [AC-FA-RULE-002](lastenheft.md#ac-fa-rule-002--keine-lateralen-adapter-kanten-regel-lateral-adapter); fehlt sie, darf **kein** Adapter einen anderen importieren (strengere Auslegung).
  - `tech` → [AC-FA-RULE-003](lastenheft.md#ac-fa-rule-003--tech-kapselung-regel-tech-leak); fehlt es, entfällt `tech-leak` (gedeckt durch die Boundary von [AC-FA-CONF-001](lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml)). `adapter` ist ein nicht-leeres Pfad-Fragment **oder** eine Pfad-**Liste** (das Symbol ist in **jedem** gelisteten Adapter erlaubt; nicht-leerer Skalar ≡ einelementige Liste, rückwärtskompatibel; eine **leere** Liste, ein leerer Listen-Eintrag, ein **leerer oder fehlender** Skalar → Exit 2 — der leere Adapter war vor 0.14.0 ein stiller Never-Leak-Eintrag, dieselbe fail-closed-Linie wie der leere `resolution`-Root). Der Adapter-Abgleich ist ein **Teilstring**-Vergleich auf dem Dateipfad, nicht segmentgrenzen-bewusst (dokumentierte Grenze, wie vor 0.14.0). Je Eintrag optional `match: substring|regex` (Default `substring`): `substring` = Teilstring-Vergleich, `regex` = **RE2**, unverankerter Suchlauf (`regexp.MatchString`) gegen das extrahierte Symbol ([SPEC-EXTRACT-001](#spec-extract-001--import-extraktion)); und optional `composition_root: allow|forbid` (Default `allow` = Composition-Root-Ausnahme wie bisher; `forbid` schaltet **nur** die `tech-leak`-Ausnahme der Composition Root für diesen Eintrag ab, [SPEC-RULE-001](#spec-rule-001--regel-auswertung)) — anderer Wert → Exit 2. Unbekannter `match`-Wert, leere oder nicht kompilierbare Regex → Exit 2 (strict-decode; ein leeres Regex-Muster würde jeden Import treffen und ist unzulässig). RE2 ist linear und deterministisch ([SPEC-DET-001](#spec-det-001--determinismus-vertrag)).
  - `composition_root` → deklarierte `tech-leak`-Ausnahme ([AC-FA-RULE-003](lastenheft.md#ac-fa-rule-003--tech-kapselung-regel-tech-leak) Boundary).
  - `allow` → konfigurativ erlaubte Sonderkante/Re-Export ([AC-FA-RULE-005](lastenheft.md#ac-fa-rule-005--schicht-richtung-regel-wrong-direction) / [AC-FA-RULE-004](lastenheft.md#ac-fa-rule-004--port-disziplin-regel-port-impurity) Boundary).
  - `markers` → dokumentierte Heuristik-Ausnahme ([AC-QA-02](lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)).
  - `forbidden_constructs` → schichtbezogen verbotene Konstrukte ([AC-FA-RULE-004](lastenheft.md#ac-fa-rule-004--port-disziplin-regel-port-impurity)); als Text-Muster geprüft (siehe [SPEC-EXTRACT-001](#spec-extract-001--import-extraktion)). **Fail-closed beim Laden — jeder Eintrag, der nie melden kann, ist Exit 2:** (a) die genannte Schicht existiert **nicht** in `layers`; (b) ihre **effektive Rolle ist nicht `port`** (explizite `role:` oder Namens-Inferenz, [AC-FA-RULE-006](lastenheft.md#ac-fa-rule-006--schicht-rollen-generische-regel-anwendung)) — die Auswertung ist an die Port-Disziplin gebunden, ein Eintrag für eine andere Rolle wäre ein stiller Never-Match; (c) ein **leeres Muster**; (d) eine **leere Musterliste**. Die Schicht-Schlüssel werden **sortiert** geprüft, damit bei mehreren Fehlern stets derselbe zuerst gemeldet wird ([SPEC-DET-001](#spec-det-001--determinismus-vertrag)). Der Fehlertext nennt `constructs` als **Gegenstück** — nicht als Ersatz: dort ist das Scoping zonen-gebunden und scan-weit, hier schicht-gebunden, und eine Schicht-Blacklist außerhalb der Rolle `port` bietet das Schema nicht an.
  - `constructs` → **Roh-Text-Monopol** je Muster ([AC-FA-RULE-011](lastenheft.md#ac-fa-rule-011--konstrukt-monopol-regel-construct-leak)): Liste von Einträgen `{pattern, adapter}` mit optionalem `match: substring|regex` (Default `substring`; `regex` = **RE2**, unverankerter Suchlauf) und optionalem `composition_root: allow|forbid` (Default `allow`). `adapter` ist die **Zone** — ein nicht-leeres Pfad-Fragment **oder** eine Pfad-**Liste**; das Muster ist in **jeder** gelisteten Zone erlaubt, der Zonen-Abgleich ist derselbe Teilstring-Vergleich auf dem Dateipfad wie bei `tech`. Schlüssel, Defaults und fail-closed-Fälle sind mit `tech` **deckungsgleich** — leeres/fehlendes `pattern`, leeres/fehlendes `adapter` (Skalar wie Liste), leerer Listen-Eintrag, unbekanntes `match`, unbekanntes `composition_root` oder eine nicht kompilierbare Regex → Exit 2. Anders als bei `tech` ist auch bei `match: substring` ein **leeres** `pattern` unzulässig (es wäre ein stiller Never-Match). Unterschied zu `tech`: gematcht wird **Roh-Quelltext** statt eines extrahierten Import-Symbols ([SPEC-EXTRACT-001](#spec-extract-001--import-extraktion)); Unterschied zu `forbidden_constructs`: das Scoping ist **zonen**-gebunden und scan-weit statt schicht-gebunden, der Befund heißt `construct-leak` ([SPEC-RULE-001](#spec-rule-001--regel-auswertung)). Fehlt der Block, entfällt `construct-leak`.
  - `resolution` → Symbol→Layer-Auflösung **je Sprache** (Map Sprache → `{mode, roots, package_base}`); `mode ∈ {path (Default), fixed-root, relative}`, `namespace` **reserviert** → Exit 2. `fixed-root`: `roots` vorangestellt; **bei gesetztem `package_base`** (gepunktete Sprache) zusätzlich Präfix-Strip + `.`→`/` (eine Pfad-Sprache wie C++ behält ihre `.`-Endungen); greift nur, wenn der Paket-Baum den Verzeichnis-Baum spiegelt ([AC-QA-02](lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)). `relative`: Specifier, die `.`/`..` sind oder mit `./`/`../` beginnen, werden lexikalisch gegen das **Verzeichnis der importierenden Datei** normalisiert (`path.Clean`-Semantik); alle anderen Specifier (Bare-Imports) sowie Wurzel-Escapes (führendes `..` **nach** der Normalisierung) liefern eine **leere** Kandidatenmenge — das Roh-Symbol wird nicht als Pfad-Kandidat weitergereicht (kein Geister-Match, [AC-QA-02](lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)); `roots`/`package_base` sind bei `relative` unzulässig → Exit 2; Endungs-Agnostik gilt, solange die `layers`-Glob-Präfixe oberhalb der Dateiebene enden (verzeichnisbasierte Globs). **Mehr-Wurzel-Auflösung (deklarations-bewusst, fail-closed):** `fixed-root` mit **≥ 2** `roots` löst den internen FQN gegen die **real gescannten Dateien** auf. Für ein **deklarations-bewusstes** Backend (Kotlin, [SPEC-EXTRACT-001](#spec-extract-001--import-extraktion)) gilt eine Evidenz-Rangfolge, stärkste zuerst: (1) **deklariert** — eine gescannte Datei im Paket-Verzeichnis `root/pkg/` trägt das Symbol als **Top-Level-Deklaration** (unabhängig vom Dateinamen); (2) **nur-Paketverzeichnis** — `root/pkg/` existiert, aber keine Datei deklariert das Symbol (eine gleichnamige, das Symbol **nicht** deklarierende Datei — und ein **Wildcard-/Paket-Import** `a.b.*` → `a/b/`, der ein ganzes Paket-Verzeichnis trifft — zählt nur hier); (3) **keine** (Phantom, extern). Die echte Deklaration **sticht** damit den bloßen Datei-Namens-Match. Für die übrigen (nicht deklarations-bewussten) Backends gilt unverändert der Datei-Namens-Match (endungs-agnostisch, package==directory: ein Wildcard-/Paket-Import — Symbol mit Trailing-Dot — trifft das Paket-Verzeichnis, ein Symbol dessen Datei mit gestrippter Endung oder sein Paket-Verzeichnis). Die Schicht wird am Pfad des **auflösenden** Kandidaten via [SPEC-RULE-001](#spec-rule-001--regel-auswertung) bestimmt, nicht am Wurzel-Präfix. Es entscheidet die **stärkste vorhandene Evidenzstufe**; auf ihr löst **genau ein** Root (oder alle dieselbe Schicht). ≥ 2 Roots **verschiedener** Schichten auf der stärksten Stufe: auf Stufe **deklariert** (bzw. für nicht deklarations-bewusste Backends: Datei-Match) ⇒ echte Mehrdeutigkeit, **Exit 2 nach dem Scan** (ein FQN muss in höchstens eine Schicht auflösen; `expect`/`actual` same-layer löst sauber); auf Stufe **nur-Paketverzeichnis** (kein Deklarations-Treffer) ⇒ **extern** (fail-open — ohne Deklaration nicht diskriminierbar). Die Stufe *nur-Paketverzeichnis* **löst** damit rückwärtskompatibel, solange sie eindeutig ist; ganz ohne Evidenz ⇒ extern. Grenzen ([AC-QA-02](lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)): ein intern gemeintes Symbol ohne Top-Level-Deklaration in gescanntem Code (verschachtelte Klasse, `object`-Member, generiert, Star-Import) bleibt still extern; die Deklarations-Auflösung ist **Kotlin-only** (übrige Backends: package==directory); datei-tiefe Globs sind eine heuristische Grenze. Fehlt der Block (oder eine Sprache) → Import-als-Pfad. Nutzt Sprache **und Pfad** der Quelldatei ([SPEC-RULE-001](#spec-rule-001--regel-auswertung)/[SPEC-EXTRACT-001](#spec-extract-001--import-extraktion)).
  - `exclude` → **Scan-Scope**: Datei-Globs relativ zur Scan-Wurzel (dieselbe Glob-Semantik wie `layers`/`languages`); eine matchende Datei wird **vor** der Extraktion vollständig vom Scan ausgenommen — sie existiert für keine Prüfung (weder Import- noch `forbidden_constructs`- noch `constructs`-Erkennung, [SPEC-EXTRACT-001](#spec-extract-001--import-extraktion)). Der Ausschluss wirkt auf den **Verzeichnis-Walk**: ein Verzeichnis, dessen **ganzer Teilbaum** von einem rekursiven Muster (`**` oder `<präfix>/**`) gedeckt ist, wird **beschnitten** (nicht betreten) statt durchlaufen und dann datei-weise gefiltert — der Prune greift **vor** dem Lesen des Verzeichnisinhalts, sodass ein ausgeschlossener Teilbaum auch unlesbar oder sehr groß sein darf, ohne den Scan abzubrechen; er ist beweisbar output-äquivalent zum Datei-Ausschluss (nicht-teilbaum-deckende Muster wie `<dir>/*` prunen **nicht**, [SPEC-EXTRACT-001](#spec-extract-001--import-extraktion)). Ein leerer Glob → Exit 2 (der ungültige Fall der ansonsten totalen Glob-Engine). Fehlt der Block, wird jede `languages`-Glob-Datei gescannt — byte-identisch zum bisherigen Verhalten ([AC-FA-CONF-001](lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml)).
- **Schicht-Rollen** ([AC-FA-RULE-006](lastenheft.md#ac-fa-rule-006--schicht-rollen-generische-regel-anwendung)): ein `layers`-Eintrag ist **entweder** eine Glob-Liste (`name: [globs]`) **oder** ein Objekt `{globs: [...], role: domain|app|port|adapter, direction: driving|driven}` (`direction` optional). Fehlt `role`, wird es aus konventionellen Namen abgeleitet (`core`→`domain`, `ports`→`port`, `adapters`→`adapter`, `application`/`app`→`app`); `role:` hat Vorrang. Die Reinheits-Regeln (`core-impurity`/`app-impurity`/`port-impurity`/`lateral-adapter`) greifen über die Rolle, nicht den Namen — fremd benannte Schichten sind damit voll prüfbar. Optional trägt eine `port`-/`adapter`-Schicht zusätzlich `direction` ∈ {`driving`, `driven`} (**orthogonal** zur Rolle, [AC-FA-RULE-008](lastenheft.md#ac-fa-rule-008--driving-driven-port-richtung-regel-port-direction-mismatch)); die Connectivity-Regel `port-direction-mismatch` prüft, dass ein Adapter nur Ports **seiner** Richtung importiert — ohne `direction` keine Prüfung.
- Kein Include/Vererbung zwischen Config-Dateien (Lastenheft-Out-of-Scope).

## SPEC-EXTRACT-001 — Import-Extraktion

Präzisiert [AC-FA-EXTRACT-001](lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion)
und [AC-QA-02](lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze).

Pro Datei, die einem Schicht-Glob entspricht, liefert das über `languages`
gewählte Backend die Menge der importierten Symbole/Module:

1. **Scan-Scope (`exclude`).** Der Verzeichnis-Walk wird an ausgeschlossenen
   Teilbäumen **beschnitten**, nicht nur die Datei-Extraktion gefiltert
   ([SPEC-CONF-001](#spec-conf-001--konfigurationsschema)):
   - Ein **Verzeichnis** wird **nicht betreten** (Prune), wenn ein `exclude`-Glob
     seinen **ganzen Teilbaum** deckt — d. h. der Glob ist `**` oder hat die Form
     `<präfix>/**`, dessen `<präfix>` das Verzeichnis matcht (`.security/**` prunt
     `.security`, `**/node_modules/**` jedes `…/node_modules`, `dist/**` `dist`).
     **Nur** solche rekursiven Teilbaum-Muster prunen: ein Single-Segment-`<dir>/*`,
     ein rand-genaues `<dir>/` oder ein Datei-Muster (`<dir>/*.go`) deckt nur
     **einen Teil** des Teilbaums und prunt **nicht** — sonst fielen Dateien, die
     der Glob gar nicht matcht, still aus dem Scope (False-Green,
     [AC-QA-02](lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)).
     So bleibt der Prune **beweisbar output-äquivalent** zum Datei-Ausschluss:
     jede Datei unter einem geprunten Verzeichnis wäre ohnehin datei-ausgeschlossen.
     Die Scan-Wurzel selbst wird nie geprunet.
   - Eine **Datei**, die einem `exclude`-Glob entspricht, wird **vor** der
     Extraktion vollständig ausgenommen — sie wird nicht gelesen und liefert
     weder Import- noch `forbidden_constructs`- noch `constructs`-Treffer.
   Der Prune greift **vor** dem Lesen des Verzeichnisinhalts: ein unlesbarer
   oder sehr großer ausgeschlossener Teilbaum bricht den Scan nicht ab. Ein
   **nicht** ausgeschlossener unlesbarer Ordner bricht dagegen weiterhin
   fail-closed ab (kein stilles Überspringen —
   [AC-QA-02](lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)).
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
   (z. B. das Glob `"**/*.py"`) darf keine echten Imports verschlucken.

   **Zeichenketten-Literale werden dabei übersprungen**, damit ein `/*` oder `//`
   **darin** keinen Kommentar öffnet: der Scanner kennt die Begrenzer `"`, `'` und
   `` ` `` und übernimmt ihren Inhalt verbatim. Ein Backslash escapt das nächste
   Byte (nicht in Backtick-Literalen); `"`- und `'`-Literale enden **auch am
   Zeilenumbruch** (keines ist mehrzeilig, ein unbalancierter Begrenzer kostet damit
   höchstens seine Zeile), Backtick-Literale laufen bis zum schließenden Begrenzer.
   **Leitplanke:** der Scanner verschluckt im Zweifel **weniger** — ein
   stehengebliebener Kommentar ist ein sichtbares Falsch-Positiv, eine verschluckte
   Zeile ein stilles Falsch-Negativ. **Ausgewiesene Grenzen** hierzu: ein
   unbalanciertes Apostroph (Rust-Lifetime `&'a str`) lässt die Kommentare **seiner
   Zeile** stehen — für die zeilen-verankerten Import-Muster folgenlos, wirksam nur
   für `constructs`/`forbidden_constructs` in derselben Zeile; und Raw-String-Formen
   mit eigener Syntax (C++ `R"(…)"`, Rust `r#"…"#`, Text-Blöcke in Java/Kotlin/C#)
   werden **nicht** erkannt.

   Import-ähnliche Zeilen in **String-Literalen** bleiben eine **ausgewiesene
   Heuristik-Grenze**: sie werden nicht gewertet, weil das Literal übersprungen wird —
   ein Muster darin greift nicht. Wo die Heuristik an
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

**Konstrukt-Treffer (`constructs`).** Ebenfalls text-heuristisch, aber **zonen**-statt
schicht-gebunden ([AC-FA-RULE-011](lastenheft.md#ac-fa-rule-011--konstrukt-monopol-regel-construct-leak)):
je gescannter Datei wird **jeder** `constructs`-Eintrag zeilenweise gegen die **vorbereitete**
Quelle geprüft — **dieselbe** Vorbereitung wie für Importe und `forbidden_constructs`
(Schritt 4), keine eigene. Für die **C-Syntax-Sprachen** sind Kommentare damit entfernt
(Zeilennummern erhalten): ein Treffer, der ausschließlich in einem Kommentar steht, entsteht
dort nicht — eine bewusste, ausgewiesene Divergenz zu einer `grep`-Referenz
([AC-QA-02](lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)). Für **Python**
gilt das **nicht**: seine Quelle wird nach Schritt 4 bewusst **nicht** C-gestrippt, und ein
`#`-Strip wäre hier die schlechtere Wahl — ein `#` in einem String-Literal würde den Zeilenrest
verschlucken und einen echten Treffer **verbergen** (False-Green, die schwerere Fehlerklasse).
Ein `constructs`-Treffer in einem Python-Kommentar wird also gemeldet; das ist die ausgewiesene
Grenze, keine stille Setzung. String-Literale bleiben in **allen** Sprachen die bestehende Grenze. Die Prüfung läuft über **alle** gescannten Dateien,
auch über solche in **keinem** `layers`-Glob (das Monopol ist eine Aussage über den Baum, nicht
über eine Schicht); `exclude` greift wie bisher davor. Je Treffer werden der **Index des
`constructs`-Eintrags** und die Zeile geliefert — nicht bloß der Mustertext, damit die
Regel-Auswertung die Zone am Eintrag selbst entscheidet und zwei Einträge mit gleichem Muster,
aber verschiedenen Zonen nicht verwechselt werden können. Ergebnis je Datei: dedupliziert je
(Eintrag, Zeile) und stabil sortiert nach (Eintrag, Zeile)
([SPEC-DET-001](#spec-det-001--determinismus-vertrag)). Die Treffer speisen die
`construct-leak`-Regel ([SPEC-RULE-001](#spec-rule-001--regel-auswertung)).

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
| `lateral-slice` | Datei mit Rolle `app` importiert ein `app`-Ziel **derselben Schicht** (`tl == f.Layer`) mit **anderer Slice-Identität** — `sliceOf(quelle) ≠ sliceOf(kandidat)`, wobei `sliceOf(p)` = das **längste `app`-Rollen-Glob-Literalpräfix**, das `p` als Segment-Run matcht (kanonisch, modul-präfix-tolerant wie die Ziel-Auflösung). Trägt die `app`-Schicht nur **ein** Glob, sind alle ihre Dateien dieselbe Slice (Regel inert, opt-in). **Getrennte `app`-Schichten** (verschiedene Layer) sind **edge-regiert** (`wrong-direction`/`allow`), **nicht** `lateral-slice`. **Kategorisch innerhalb der Schicht** (nicht über `edges`/`allow` aufhebbar) | [AC-FA-RULE-009](lastenheft.md#ac-fa-rule-009--slice-isolation-regel-lateral-slice) |
| `tech-leak` | ein `tech`-Muster (Substring oder RE2-Regex, je `match`) erscheint außerhalb **aller** seiner zugeordneten Adapter (`adapter` als Pfad oder Pfad-Liste) — und außerhalb `composition_root`, sofern der Eintrag nicht `composition_root: forbid` deklariert (dann prüft `tech-leak` auch dort; die Meldung nennt alle gelisteten Adapter in Deklarationsreihenfolge) | [AC-FA-RULE-003](lastenheft.md#ac-fa-rule-003--tech-kapselung-regel-tech-leak) |
| `port-impurity` | Datei mit Rolle `port` importiert eine `adapter`-Rolle oder ein `tech`-Muster **oder** enthält ein `forbidden_constructs`-Muster (text-heuristisch erkannt). **Kern-Referenzen sind erlaubt** (Ports sprechen die Sprache des Kerns) und werden über `edges`/`allow` regiert — eine undeklarierte `ports → core`-Kante fällt unter `wrong-direction` | [AC-FA-RULE-004](lastenheft.md#ac-fa-rule-004--port-disziplin-regel-port-impurity) |
| `port-direction-mismatch` | Datei mit Rolle `adapter` und Richtung `direction` X importiert eine `port`-Rolle mit Richtung Y (X ≠ Y, **beide gesetzt**) — ein Treiber-Adapter spricht nur `driving`-Ports, ein getriebener nur `driven`-Ports; **orthogonal** zur Rolle, ohne `direction` keine Prüfung. **Kategorisch** (nicht über `edges`/`allow` aufhebbar, wie `lateral-adapter`) | [AC-FA-RULE-008](lastenheft.md#ac-fa-rule-008--driving-driven-port-richtung-regel-port-direction-mismatch) |
| `port-locality` | Datei mit Rolle `app` importiert ein **im Application-Baum geschachteltes** `port`-Ziel, dessen **Scope-Verzeichnis** ihren Pfad nicht enthält — `portScope ≠ "" ∧ appTreeContains(portScope) ∧ segIndex(quelle, portScope) < 0`, wobei `portScope(k)` = das **längste `port`-Rollen-Glob-Literalpräfix**, das `k` matcht, **minus seinem letzten Pfad-Segment** (Port-Ordner-Marker, typisch `ports`): use-case-lokal (`…/createorder`) ⊂ business-area (`…/order`) ⊂ app-weit (`…/application`). `appTreeContains(s)` = `s` ist Vorfahr eines `app`-Glob-Präfixes (der Port liegt *im* App-Baum). **Geschwister-Ports** (klassisch, `…/ports` neben `…/services`) → `appTreeContains` false → **inert**. **Nur `app`-Importeure** (ein Adapter, der einen Port implementiert, ist nicht erfasst). **Kategorisch** (nicht über `edges`/`allow` aufhebbar) | [AC-FA-RULE-010](lastenheft.md#ac-fa-rule-010--port-lokalität-regel-port-locality) |
| `wrong-direction` | ein Import quert eine Schicht-Kante entgegen `edges`/`allow` | [AC-FA-RULE-005](lastenheft.md#ac-fa-rule-005--schicht-richtung-regel-wrong-direction) |
| `construct-leak` | ein `constructs`-Muster (Substring oder RE2, je `match`) erscheint im **Roh-Quelltext** einer Datei außerhalb **aller** seiner Zonen (`adapter` als Pfad oder Pfad-Liste) — und außerhalb `composition_root`, sofern der Eintrag nicht `composition_root: forbid` deklariert. **Nicht** import-, sondern **datei**-bezogen: die Regel wertet die Konstrukt-Treffer je Datei aus (wie der konstrukt-basierte Zweig von `port-impurity`) und gilt **scan-weit**, auch für Dateien in **keinem** `layers`-Glob | [AC-FA-RULE-011](lastenheft.md#ac-fa-rule-011--konstrukt-monopol-regel-construct-leak) |

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
dispatchen über die Rolle, nicht den Namen.

**Rückzug bei unauflösbarem Ziel-Glob.**
Ein Glob mit **Wildcard in der Mitte** (`…/application/**/ports/**`) trägt keinen literalen Präfix
und kann als Import-**Ziel** nie matchen. Der Kandidat darf dann **nicht** auf das nächst-passende,
umschließende Glob zurückfallen — das ergäbe einen Befund über eine Schicht, in der das Ziel
womöglich gar nicht liegt (ein legitimer `adapter → ports`-Import würde als
`wrong-direction: adapter -> application` gemeldet). Stattdessen wird die Zuordnung
**zurückgezogen**: das Ziel gilt als **extern** (fail-open, kein Befund) — dieselbe Linie wie bei
der schwachen Evidenzstufe der Mehr-Wurzel-Auflösung. Der Rückzug greift **nur**, wenn ein Glob
einer **anderen** Schicht alle drei Bedingungen erfüllt: sein literaler **Kopf** (vor dem ersten
Wildcard-Segment) **und** sein literaler **Tail-Marker** (nach dem letzten Wildcard-Segment, z. B.
`ports`) kommen beide als Segment-Run im Kandidaten vor, **und** der Kopf ist mindestens so
spezifisch wie der Präfix der gewählten Schicht. Der Tail-Marker ist tragend: ohne ihn würde
`…/application/**/ports/**` jeden Application-Import zurückziehen und die Schicht still abschalten.
Ein Prefix, das auf dem Wildcard-Segment endet (kein Tail-Marker, z. B. `a/**/**`), zieht nie
zurück. Die Prüfung ist eine reine Existenz-Aussage über die Glob-Menge — reihenfolge-unabhängig
([SPEC-DET-001](#spec-det-001--determinismus-vertrag)). Die **Quell**-Seite (Datei → Schicht) ist
unberührt: sie bewertet den literalen Kopf und löst solche Globs korrekt auf — die Asymmetrie
bleibt, wird aber ausgewiesen statt in einen Fehlbefund übersetzt
([AC-QA-02](lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)). Die **`tech`-Muster** dagegen lösen in
**Deklarationsreihenfolge** auf (**Erst-Treffer** gewinnt, `matchTech`) — uniform für
`match: substring` und `match: regex`; es gibt für `tech` **kein** „längster Präfix"
([AC-FA-RULE-003](lastenheft.md#ac-fa-rule-003--tech-kapselung-regel-tech-leak)).

**Slice-Identität und Port-Scope** ([AC-FA-RULE-009](lastenheft.md#ac-fa-rule-009--slice-isolation-regel-lateral-slice)/[AC-FA-RULE-010](lastenheft.md#ac-fa-rule-010--port-lokalität-regel-port-locality))
nutzen dieselbe **segment-basierte, modul-präfix-tolerante** Kandidaten-Auflösung wie die Ziel-Rolle
(längstes Rollen-Glob-**Literalpräfix** als Segment-Run, `segIndex`): `sliceOf(p)` = das längste `app`-
Glob-Literalpräfix, `portScope(k)` = das längste `port`-Glob-Literalpräfix **minus dem letzten Segment**
(Port-Ordner-Marker). Beide vergleichen **kanonisch** (das Glob-Präfix ist für Quelle wie Kandidat dieselbe
Zeichenkette), sodass der Modul-Präfix eines Kandidaten (`hexslice/example/internal/…`) nicht stört.
`lateral-slice` greift nur **innerhalb einer** `app`-Schicht (`tl == f.Layer`) — Slices sind die
per-Glob-Untereinheiten *einer* Schicht; verschiedene `app`-**Layer** sind edge-regiert. `port-locality`
greift nur, wenn `portScope` im **App-Baum** liegt (`appTreeContains`) — Geschwister-Ports (klassisches
Hexagonal) haben keine per-Slice-Lokalität.
**Voraussetzung** ([AC-QA-02](lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)): die `app`-/`port`-
Globs müssen als Import-**Ziel** auflösen, also ein **sauberes literales Präfix** tragen (per-Slice
`…/createorder/**`, per-Port-Ordner `…/ports/**`); ein Glob mit Wildcard **in der Mitte** (`…/**/ports/**`)
oder Datei-Endung (`…/*.go`) hat kein solches Präfix und löst als Ziel **nicht** auf — dann bleiben die
beiden Regeln (wie schon die bestehende Ziel-Rollen-Auflösung) für dieses Ziel wirkungslos, eine
ausgewiesene Grenze, kein stiller Fehlbefund.

Pro (Datei, Import) gilt **deterministische Erst-Treffer-Reihenfolge** in der
Tabellen-Reihenfolge (`core-impurity` → `app-impurity` → `port-impurity` →
`lateral-adapter` → `lateral-slice` → `tech-leak` → `port-direction-mismatch` →
`port-locality` → `wrong-direction`); ein Import erzeugt höchstens einen Befund.
**`construct-leak` steht außerhalb dieser Kette**: es bewertet keine Importe, sondern die
Roh-Text-Treffer einer Datei ([SPEC-EXTRACT-001](#spec-extract-001--import-extraktion)) — je
Treffer außerhalb seiner Zone genau ein Befund. Treffen zwei `constructs`-Muster **dieselbe
Zeile** und liegen beide außerhalb ihrer Zone, entstehen **zwei** Befunde; ihre Reihenfolge ist
über die Totalordnung aus [SPEC-DET-001](#spec-det-001--determinismus-vertrag) festgelegt.
`port-locality` steht **vor** `wrong-direction`, damit eine erlaubte `app → port`-Kante die
Lokalitäts-Verletzung nicht maskiert (kategorisch).
Dateien unter `composition_root` sind als Verdrahtungspunkt von **allen**
Schicht-Regeln ausgenommen — sie importieren bestimmungsgemäß quer über die
Schichten — und von `tech-leak` bzw. `construct-leak` **je Eintrag**: bei
`composition_root: allow` (Default) wie bisher, bei `composition_root: forbid`
prüft die jeweilige Regel den Eintrag auch dort weiter
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
- **Abdeckungs-Diagnose** (advisory, auf stderr **nach** der Zusammenfassung): Gescannte Dateien,
  die in **keinem** `layers`-Glob liegen, unterliegen keiner Schicht-Regel — eine bewusste
  fail-open-Grenze ([AC-QA-02](lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)),
  die **ausgewiesen** wird: der Scan nennt ihre Anzahl und ihre Pfade. **Der Exit-Code bleibt davon
  unberührt** — ein befundfreier Lauf endet auch mit ungedeckten Dateien mit `0`. Nicht gezählt
  werden Dateien unter `composition_root` (bestimmungsgemäß schichtlos) und `exclude`-Dateien (nie
  im Scan, [SPEC-EXTRACT-001](#spec-extract-001--import-extraktion)). Gezählt wird ausschließlich
  die **Quell**-Seite (die gescannte Datei), **nicht** die Ziel-Seite eines Imports: ein Ziel ohne
  auflösbare Schicht ist von repo-**externem** Code nicht sicher unterscheidbar. Die Pfade sind
  **stabil sortiert** ([SPEC-DET-001](#spec-det-001--determinismus-vertrag)); ab **zehn** Dateien
  wird die Liste gekürzt und die **Restzahl genannt** (keine stille Kappung). Sind alle gescannten
  Dateien gedeckt, entsteht **keine** Ausgabe.
- **Grenz-Diagnose** (advisory, auf stderr **nach** der Abdeckungs-Diagnose): Import-Zeilen, deren
  **Schreibweise** das Werkzeug per Konstruktion nicht zu einer beurteilbaren Kante führt, werden als
  `pfad:zeile: form` ausgewiesen — die Einlösung von
  [AC-QA-02](lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze) am geprüften Baum statt
  nur in der Doku. Zwei Klassen zählen, beide **ohne den Datei-Baum** entscheidbar — aus der
  Schreibweise und der Konfiguration, nie aus dem Index:
  1. **Nicht extrahiert** — die Zeile greift kein Backend-Muster: ein **relativer Python-Import**
     (führender Punkt) und eine **zweite Direktive auf derselben Zeile** (`import a, b`,
     `using A; using B;`); beides benannte Grenzen aus
     [SPEC-EXTRACT-001](#spec-extract-001--import-extraktion).
  2. **Extrahiert, aber durch die Konfiguration unauflösbar** — ein Symbol mit `./`- oder
     `../`-Präfix unter einem Auflösungs-Modus, der **nicht** `relative` ist
     ([SPEC-CONF-001](#spec-conf-001--konfigurationsschema)), **und** kein Glob-Präfix der
     `layers` kommt in einem seiner Auflösungs-Kandidaten segmentweise vor. Beide Hälften zählen:
     die Schreibweise **allein** genügt nicht. Unter `path` wird das Symbol wörtlich zum Kandidaten
     und die Schicht-Zuordnung sucht das Glob-Präfix segmentweise **an beliebiger Stelle** — ein
     `../adapters/db/x.h` gegen `adapters/**` löst also auf und wird **nicht** gemeldet. Unter
     `relative` löst dieselbe Zeile ohnehin auf. Geprüft werden Modus und Globs, beides
     Konfiguration; der Datei-Index bleibt außen vor.

  **Nicht** gemeldet wird ein Symbol, das syntaktisch auflösbar wäre und im Baum nur kein Ziel
  findet — das ist von repo-**externem** Code nicht unterscheidbar, dieselbe Grenze, an der die
  Abdeckungs-Diagnose die Ziel-Seite ausnimmt. **Der Exit-Code bleibt unberührt**; die Meldung
  erscheint auch bei **null** Befunden. Stabil sortiert nach (Pfad, Zeile)
  ([SPEC-DET-001](#spec-det-001--determinismus-vertrag)); ab **zehn** Zeilen wird gekürzt und die
  **Restzahl genannt**. Ein Baum ohne solche Zeilen erzeugt **keine** Ausgabe.
- **Auflösungs-Diagnose** (advisory, auf stderr **nach** der Grenz-Diagnose): Löst im **gesamten
  Scan** kein einziges extrahiertes Symbol auf eine Schicht auf, obwohl Symbole extrahiert wurden,
  nennt der Scan je Schicht mit Symbolen eine Zeile:
  `Schicht <name>: N Datei(en), 0 von M Import-Symbolen lösen auf eine Schicht auf`. Das ist die
  gefährlichste Konfiguration, die dieses Werkzeug zulässt — alle Dateien in Schichten, alle
  Symbole extrahiert, und trotzdem wird keine Kante beurteilt, weil jedes Ziel über den fail-open-Pfad
  ([SPEC-RULE-001](#spec-rule-001--regel-auswertung)) als repo-extern gilt.

  Die Auslöse-Bedingung ist **repo-weit, nicht je Schicht**: eine einzelne Schicht ohne auflösende
  Symbole ist legitim (ein abhängigkeitsfreier Kern importiert nur die Standardbibliothek) und von
  einer falsch konfigurierten Auflösung nicht zu unterscheiden — dieselbe Grenze, an der die
  Abdeckungs-Diagnose die Ziel-Seite ausnimmt. Ein Scan, in dem **nirgends** eine Kante entsteht,
  ist dagegen sicher beurteilbar. Daraus folgt ausdrücklich: der **Teil**ausfall (eine von mehreren
  Schichten falsch aufgelöst) bleibt **still**. Schichten **ohne** extrahierte Symbole werden nicht
  genannt; `composition_root`-Dateien zählen nicht. Stabil nach Schichtnamen sortiert
  ([SPEC-DET-001](#spec-det-001--determinismus-vertrag)); **Exit-Code unberührt**, Meldung auch bei
  **null** Befunden.
- **Read-only:** der geprüfte Baum wird nie beschrieben (Mount `:ro`).
- **No-scan-Modus `--print-graph`** ([AC-FA-CLI-002](lastenheft.md#ac-fa-cli-002--architektur-graph-ausgabe)):
  `a-check --print-graph [pfad]` liest **nur** `pfad/.a-check.yml` und gibt einen Graphen aus
  ([SPEC-CLI-002](#spec-cli-002--graph-renderer-vertrag)) — **kein** Quell-Scan. Es gilt
  **load-time/config-validation-Parität** zum Scan: das strikte Schema-Decoding
  ([SPEC-CONF-001](#spec-conf-001--konfigurationsschema)) **und** die Sprach-Backend-Prüfung
  gegen die Registry ([SPEC-EXTRACT-001](#spec-extract-001--import-extraktion)) laufen als
  **validation-only**-Schritt **ohne** Datei-Walk; jeder ladezeitige Fehler → Exit 2 (unbekannter
  Schlüssel, ungültige `role`/`direction`/`tech`, **unbekannte Sprache** u. a.). Zusätzlich brechen
  ein unbekanntes Flag **und** — nur in diesem Modus — ein weiteres Positionsargument/Flag **nach**
  dem optionalen `pfad` (`--print-graph <pfad> --bogus`) mit Exit 2 (Go-`flag` parst nach dem ersten
  Positionsargument nicht weiter; das darf nicht still als zweites Argument ignoriert werden).
  Bewusst **keine** Parität für scanzeitige, datei-mengenabhängige Fehler (Mehr-Wurzel-Auflösungs-
  Mehrdeutigkeit, [SPEC-CONF-001](#spec-conf-001--konfigurationsschema)) — der Modus liest keine
  Quellen. Erfolgreiche Ausgabe → Exit 0, read-only.

## SPEC-CLI-002 — Graph-Renderer-Vertrag

Präzisiert [AC-FA-CLI-002](lastenheft.md#ac-fa-cli-002--architektur-graph-ausgabe).

Der Renderer bildet das geladene Config-Modell **pur** (kein I/O) auf einen
Mermaid-`flowchart`-String ab; die Composition Root schreibt ihn nach stdout.

- **Knoten:** ein Knoten je `layers`-Eintrag. Layer-Namen sind freie YAML-Map-Keys und taugen
  **nicht** als Mermaid-Knoten-IDs; der Renderer vergibt **stabile interne IDs** (`L0`, `L1`, … in
  Sortier-Reihenfolge) und gibt den Rohnamen **nur als escaptes Label** aus. `composition_root` wird
  als **ein** isolierter Notizknoten (`C0`) mit seinen Globs als Labels gerendert, `adapter_sink` als
  **ein** isolierter Ausnahme-/Notizknoten (`S0`) mit dem Pfadfragment — beide **ohne** gezeichnete
  Kanten und ohne Layer-Klasse; fehlen sie, entsteht **kein** leerer Sonderknoten.
- **Kanten:** eine Kante `Lᵢ --> Lⱼ` je `edges`-Eintrag (interne IDs, nie Rohnamen); je `allow`-Eintrag
  eine **abgesetzte** (gestrichelte, gelabelte) Kante. Ein `edges`/`allow`-Endpunkt, der auf **keinen**
  deklarierten Layer zeigt, wird als deterministischer **Dangling-Knoten** (`D0`, `D1`, …) gerendert —
  der no-scan-Modus führt **keine** neue Config-Strenge ein (kein Exit 2 für einen Endpunkt, den der
  Loader heute akzeptiert, [SPEC-CONF-001](#spec-conf-001--konfigurationsschema)).
- **Rollen-Farbe:** ein `classDef` je **effektiver** Rolle (explizites `role:` oder Kern-Namens-Inferenz
  `core`→`domain`/`ports`→`port`/`adapters`→`adapter`/`application`,`app`→`app`,
  [SPEC-RULE-001](#spec-rule-001--regel-auswertung)); es gilt **derselbe** Resolver wie in der
  Regel-Engine, keine kopierte Inferenz.
- **Richtung:** Layer mit `direction` ∈ {`driving`, `driven`} werden in **stabilen** Mermaid-Subgraphs
  gruppiert; Layer ohne `direction` bleiben außerhalb, behalten aber Rolle und Kanten. Kanten dürfen
  Subgraph-Grenzen queren.
- **Legende:** die impliziten, kategorischen Constraints (`core-impurity`, `lateral-adapter`,
  `lateral-slice`, `port-direction-mismatch`, `port-locality`) erscheinen als Legende/Notiz, **nicht** als gezeichnete Kante
  — ebenso das Roh-Text-Monopol `construct-leak`
  ([AC-FA-RULE-011](lastenheft.md#ac-fa-rule-011--konstrukt-monopol-regel-construct-leak)), das
  gar keine Schicht-Kante hat: die Legende ist der **normative Ort** für Nicht-Kanten-Semantik
  ([AC-QA-02](lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze): ehrliche Ausgabe, keine
  Semantik-Behauptung über den realen Code).
- **Escaping-Vertrag:** jeder nutzergesteuerte Text (Layer-Namen, `composition_root`-Globs,
  `adapter_sink`, Legenden-Text) wird zuerst einzeilig normalisiert (`\r`/`\n`/`\t` → Leerzeichen,
  übrige Steuerzeichen → `?`) und dann in fester Reihenfolge kodiert: `&`→`&amp;`, `"`→`&quot;`,
  `<`→`&lt;`, `>`→`&gt;`, `[`→`&#91;`, `]`→`&#93;`, `|`→`&#124;`, `\`→`&#92;`, `` ` ``→`&#96;`. Die
  Renderer-eigene Struktur (`["…"]`, feste `<br/>`-Trenner) wird erst **danach** um den kodierten Text
  gelegt; nutzergesteuerter Text erzeugt nie Mermaid-Syntax.
- **Determinismus** ([SPEC-DET-001](#spec-det-001--determinismus-vertrag)): Knoten stabil nach
  Layer-Name sortiert, Kanten nach (`from`, dann `to`), `classDef` in fester Reihenfolge — **byte-
  identische** Ausgabe bei identischer Config, unabhängig von der internen Map-/Slice-Ordnung des
  Modells. `tech` wird in v1 nicht gerendert.

## SPEC-DET-001 — Determinismus-Vertrag

Präzisiert [AC-QA-01](lastenheft.md#ac-qa-01--determinismus).

Identische Eingabe (Repo-Stand + `.a-check.yml` + Image-Digest) ⇒
**byte-identische** Ausgabe und identischer Exit-Code. Befunde werden nach
einer Totalordnung sortiert: `pfad`, dann `zeile`, dann `regelname`, dann
**`meldung`**. Der letzte Schlüssel macht die Ordnung erst total: dieselbe Datei/Zeile
kann mehrere Befunde **derselben** Regel tragen (zwei `constructs`-Muster oder zwei
`forbidden_constructs`-Muster in einer Zeile), und ohne ihn hinge deren Reihenfolge an der
internen Eingabe-Ordnung.
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
  mit `A_CHECK_IMAGE`, einem `a-check`-Scan-Target **und** einem
  `a-check-graph`-Target; schreibt nichts. Das `a-check-graph`-Target führt
  `--print-graph` aus ([SPEC-CLI-002](#spec-cli-002--graph-renderer-vertrag):
  Mermaid nach stdout, read-only, kein Scan) über **dasselbe** `A_CHECK_IMAGE`
  und denselben netzlosen read-only-Mount — kein zweiter Digest.
  - `A_CHECK_IMAGE` trägt im **erzeugten** Fragment einen **Platzhalter** statt
    eines konkreten Digests: das Binary kann den Digest des Image, in dem es läuft, nicht kennen — er
    entsteht erst beim Push. Der Platzhalter ist so gewählt, dass eine
    unveränderte Übernahme **sichtbar abbricht** statt ein fremdes Release zu
    ziehen. Die im Repo **committete** `a-check.mk` trägt dagegen den echten
    Digest des aktuellen Release; sie ist das gepinnte Artefakt, das erzeugte
    Fragment die Vorlage dafür.
  - Die Container-Runtime steht als `DOCKER ?= docker` im Fragment und wird als
    `$(DOCKER)` aufgerufen, damit ein Konsument mit eigener Runtime nicht die
    Hälfte seiner Targets anders fährt als die andere.
- Ein unbekanntes Flag ⇒ Exit-Code 2.

## Historie

Regeln dieser Sektion: **kein ADR- und kein Slice-Verweis.** Die Decken-Regel gilt für alle drei
Spec-Straten — welche ADR eine Festlegung schärft, deklariert die ADR aufwärts in ihrem
`Schärft:`-Feld (Baseline-Regelwerk `modul-03-spec.md` §Ziel-Form: Spezifikation).

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
| 0.19.0 | 2026-07-09 | Neu `SPEC-CLI-002` (Graph-Renderer-Vertrag: Config-Modell→Mermaid **pur**; stabile interne IDs + escaptes Label je nutzergesteuertem Text; Kante je `edges`, abgesetzte `allow`-Kante; Dangling-/Composition-Root-/Adapter-Sink-Sonderknoten; `classDef` je effektiver Rolle via geteiltem Resolver; `direction`-Subgraphs; implizite Regeln als Legende; Escaping-Vertrag; Determinismus-Ordnung; `tech` v1 deferred). `SPEC-CLI-001` um den no-scan-`--print-graph`-Modus präzisiert (load-time/config-validation-Parität inkl. unbekannter Sprache; Restargument nach dem Pfad → Exit 2; **keine** scanzeitige Fehler-Parität). Folgt [`AC-FA-CLI-002`](lastenheft.md#ac-fa-cli-002--architektur-graph-ausgabe) 0.19.0. |
| 0.31.0 | 2026-08-15 | `SPEC-EXTRACT-001`: **der Kommentar-Strip überspringt Zeichenketten-Literale** — ein `/*` oder `//` in einem String öffnete bisher einen Phantom-Kommentar bis zum nächsten `*/`, wodurch Importe, `forbidden_constructs`, `constructs` und Deklarationen dahinter **still** verschwanden (sieben C-Syntax-Backends; Python war gegen genau das ausgenommen). Der Scanner kennt jetzt die Begrenzer `"`, `'` und Backtick, mit Backslash-Escape (nicht bei Backtick) und Zeilen-Ende für `"`/`'`. **Leitplanke:** im Zweifel weniger verschlucken — ein stehengebliebener Kommentar ist ein sichtbares Falsch-Positiv, eine verschluckte Zeile ein stilles Falsch-Negativ. Ausgewiesene Grenzen: unbalanciertes Apostroph (Rust-Lifetime) wirkt auf die **nicht** zeilen-verankerten Muster seiner Zeile; Raw-String-Formen (`R"(…)"`, `r#"…"#`, Text-Blöcke) bleiben unerkannt. **Verhaltensänderung:** bei Konsumenten können zuvor verschluckte Verstöße sichtbar werden ([ADR-0034](../docs/plan/adr/0034-stripcomments-string-literale.md)). slice-090. |
| 0.30.0 | 2026-08-09 | `SPEC-CONF-001`: **`forbidden_constructs` fail-closed beim Laden** — der Block wurde bisher ungeprüft durchgereicht und hatte **vier** stille Ausgänge, die alle mit Exit 0 endeten: unbekannte Schicht (Tippfehler), Schicht ohne Rolle `port`, leeres Muster, leere Musterliste. Alle vier sind jetzt **Exit 2**; Schicht-Schlüssel sortiert geprüft ([SPEC-DET-001](#spec-det-001--determinismus-vertrag)). Schließt dieselbe falsch-grüne Klasse, die `constructs`/`tech`/`languages` längst fail-closed behandeln ([AC-QA-02](lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)); die Bindung an `role: port` bleibt, weil sie [AC-FA-RULE-004](lastenheft.md#ac-fa-rule-004--port-disziplin-regel-port-impurity) („Port-Disziplin") **einlöst** — eine Ausweitung wäre eine Lastenheft-Änderung. Der Fehlertext nennt `constructs` als Gegenstück, nicht als Ersatz; die Schicht-Blacklist außerhalb `port` bleibt eine ausgewiesene Angebots-Lücke. Kein Lastenheft-Bump ([ADR-0033](../docs/plan/adr/0033-forbidden-constructs-fail-closed.md)). slice-086. |
| 0.29.0 | 2026-08-09 | `SPEC-CLI-001`: **Auflösungs-Diagnose** (advisory, nach der Grenz-Diagnose) — löst im **gesamten Scan** kein extrahiertes Symbol auf eine Schicht auf, obwohl Symbole extrahiert wurden, nennt der Scan je Schicht mit Symbolen `Schicht <name>: N Datei(en), 0 von M Import-Symbolen lösen auf eine Schicht auf`. Deckt die gefährlichste Konfiguration ab: alle Dateien in Schichten, alle Symbole extrahiert, jedes Ziel per fail-open extern — **vollständig grün, vollständig blind**. Auslösung **repo-weit, nicht je Schicht**: eine einzelne Schicht ohne auflösende Symbole ist legitim (abhängigkeitsfreier Kern) und von kaputter Auflösung nicht unterscheidbar; daraus folgt ausdrücklich, dass der **Teil**ausfall still bleibt. Schichten ohne Symbole werden nicht genannt, `composition_root` zählt nicht, stabil nach Schichtnamen sortiert, **Exit-Code unberührt**, Meldung auch bei null Befunden. Kein Lastenheft-Bump ([ADR-0032](../docs/plan/adr/0032-aufloesungs-diagnose-repoweit.md)). slice-085. |
| 0.28.0 | 2026-08-09 | `SPEC-CLI-001`: **Grenz-Diagnose** (advisory, nach der Abdeckungs-Diagnose) — ein Scan weist als `pfad:zeile: form` die Import-Zeilen aus, deren **Schreibweise** per Konstruktion zu keiner beurteilbaren Kante führt: (1) **nicht extrahiert** (relativer Python-Import, zweite Direktive auf derselben Zeile) und (2) **extrahiert, aber strukturell unauflösbar** (`./`/`../`-Präfix unter einem Modus ≠ `relative`; unter `relative` still). Ein Symbol, das nur im konkreten Baum kein Ziel findet, wird **nicht** gemeldet — von repo-externem Code nicht unterscheidbar. **Exit-Code unberührt**, Meldung auch bei null Befunden, stabil nach (Pfad, Zeile) sortiert, ab zehn Zeilen gekürzt **mit Restzahl**; ein Baum ohne solche Zeilen bleibt still. Löst die Doku-Zusage von [AC-QA-02](lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze) am geprüften Baum ein, ohne eine Grenze zu verschieben; kein Lastenheft-Bump ([ADR-0031](../docs/plan/adr/0031-heuristik-grenzen-diagnose.md)). slice-081. |
| 0.27.0 | 2026-08-09 | `SPEC-DIST-001`: `A_CHECK_IMAGE` trägt im **erzeugten** Fragment einen **Platzhalter** statt eines Digests ([ADR-0030](../docs/plan/adr/0030-kein-digest-im-generierten-fragment.md)) — das Binary kann den Digest seines eigenen Image nicht kennen, der eingebackene Wert nannte immer den Vorgänger. Die committete `a-check.mk` trägt weiterhin den echten Digest. Zusätzlich als Fragment-Bestandteil ausgewiesen: `DOCKER ?= docker` + `$(DOCKER)`-Aufruf (slice-082). Folgt [`AC-QA-03`](lastenheft.md#ac-qa-03--reproduzierbarkeit) 0.23.0. slice-083. |
| 0.20.0 | 2026-07-09 | `SPEC-DIST-001`: das `--print-mk`-Fragment liefert zusätzlich ein `a-check-graph`-Target, das `--print-graph` ([`SPEC-CLI-002`](#spec-cli-002--graph-renderer-vertrag)) über dasselbe digest-gepinnte `A_CHECK_IMAGE` + netzlosen read-only-Mount ausführt (Mermaid→stdout, kein Scan); kein zweiter Digest. Folgt [`AC-FA-DIST-001`](lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk) 0.20.0. |
| 0.26.0 | 2026-07-25 | `SPEC-CLI-001`: **Abdeckungs-Diagnose** (advisory) — ein Scan nennt auf stderr nach der Zusammenfassung die gescannten Dateien, die in **keinem** `layers`-Glob liegen; der **Exit-Code bleibt unberührt**. `composition_root`-Dateien zählen nicht, `exclude`-Dateien sind nie im Scan; gezählt wird nur die **Quell**-Seite (ein Import-Ziel ohne Schicht ist von repo-externem Code nicht unterscheidbar). Pfade stabil sortiert, ab zehn Dateien gekürzt **mit ausgewiesener Restzahl**; vollständige Abdeckung erzeugt keine Ausgabe. Macht die bewusste fail-open-Grenze aus [AC-QA-02](lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze) sichtbar, ohne sie zu verschieben; kein Lastenheft-Bump ([ADR-0029](../docs/plan/adr/0029-abdeckungs-diagnose-advisory.md)). |
| 0.25.0 | 2026-07-25 | `SPEC-RULE-001`: **Rückzug bei unauflösbarem Ziel-Glob** — ein Glob mit Wildcard in der Mitte (`…/application/**/ports/**`) kann als Import-**Ziel** nie matchen; der Kandidat fällt jetzt **nicht mehr** auf das umschließende Glob zurück, sondern gilt als **extern** (fail-open). Der Rückfall erzeugte einen **Fehlbefund** (`wrong-direction` auf eine deklarierte `adapter → ports`-Kante), dessen naheliegende Reparatur ein dauerhaftes Falsch-Negativ nach sich zieht. Rückzug eng gefasst: literaler **Kopf** *und* **Tail-Marker** des anderen Globs müssen im Kandidaten vorkommen, und der Kopf muss mindestens so spezifisch sein wie der gewählte Präfix; ohne Tail-Marker kein Rückzug. Quell-Seite (Datei→Schicht) unberührt; reihenfolge-unabhängig ([AC-QA-01](lastenheft.md#ac-qa-01--determinismus)), kein Lastenheft-Bump ([ADR-0028](../docs/plan/adr/0028-ziel-glob-schattenwurf.md)). |
| 0.24.0 | 2026-07-25 | `SPEC-CONF-001`/`SPEC-EXTRACT-001`/`SPEC-RULE-001`: neuer Optionalblock **`constructs`** — Roh-Text-Muster mit erlaubter **Zone** (`adapter` als Pfad/Pfad-Liste, `match: substring\|regex`, `composition_root: allow\|forbid`; Schlüssel/Defaults/fail-closed-Fälle deckungsgleich mit `tech`, zusätzlich leeres `pattern` auch bei `substring` → Exit 2) und die Regel **`construct-leak`**: jedes Vorkommen außerhalb aller Zonen ist ein Befund. Auswertung **datei**- statt import-bezogen (außerhalb der Erst-Treffer-Kette) und **scan-weit** inkl. Dateien in keinem `layers`-Glob; gematcht wird die **vorbereitete** (kommentar-bereinigte) Quelle — ausgewiesene Divergenz zur `grep`-Referenz ([AC-QA-02](lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)); Treffer tragen den **Eintrags-Index**, nicht den Mustertext. `SPEC-DET-001`: Befund-Totalordnung um **`meldung`** als letzten Schlüssel erweitert (dieselbe Datei/Zeile/Regel kann mehrfach auftreten). `SPEC-CLI-002`: Legende nennt `construct-leak` als Nicht-Kanten-Semantik. Folgt [`AC-FA-RULE-011`](lastenheft.md#ac-fa-rule-011--konstrukt-monopol-regel-construct-leak)/[`AC-FA-CONF-001`](lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml) 0.22.0 ([ADR-0027](../docs/plan/adr/0027-constructs-roh-text-monopol.md)). |
| 0.23.0 | 2026-07-24 | `SPEC-CLI-002`: die Graph-Legende nennt jetzt **alle fünf** kategorischen Regeln — `lateral-slice` und `port-locality` ([AC-FA-RULE-009](lastenheft.md#ac-fa-rule-009--slice-isolation-regel-lateral-slice)/[AC-FA-RULE-010](lastenheft.md#ac-fa-rule-010--port-lokalität-regel-port-locality), 0.22.0) ergänzen `core-impurity`/`lateral-adapter`/`port-direction-mismatch`; reine Legenden-Notiz, keine gezeichnete Kante ([AC-QA-02](lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)). Nachzug zu [`AC-FA-CLI-002`](lastenheft.md#ac-fa-cli-002--architektur-graph-ausgabe). |
| 0.22.0 | 2026-07-24 | `SPEC-RULE-001`: zwei neue kategorische Regeln über das Rollenmodell — **`lateral-slice`** (`app`-Datei importiert `app`-Ziel anderer Slice-Identität; `sliceOf` = längstes `app`-Glob-Literalpräfix; opt-in über per-Slice-Globs) und **`port-locality`** (`app`-Datei importiert `port` außerhalb dessen `portScope` = längstes `port`-Glob-Literalpräfix minus letztem Segment; nur `app`-Importeure). Erst-Treffer-Kette um beide erweitert (`port-locality` vor `wrong-direction`). Voraussetzung: als Import-Ziel auflösbare `app`/`port`-Globs (sauberes literales Präfix; `**/…/**`/`*.go` lösen nicht auf — [AC-QA-02](lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)-Grenze). Folgt [`AC-FA-RULE-009`](lastenheft.md#ac-fa-rule-009--slice-isolation-regel-lateral-slice)/[`AC-FA-RULE-010`](lastenheft.md#ac-fa-rule-010--port-lokalität-regel-port-locality) 0.21.0 ([ADR-0026](../docs/plan/adr/0026-hexslice-vertical-slice-regeln.md)). |
| 0.21.0 | 2026-07-23 | `SPEC-EXTRACT-001`/`SPEC-CONF-001`: **`exclude` beschneidet den Verzeichnis-Walk** (Prune), nicht nur die Datei-Extraktion — ein Verzeichnis wird nicht betreten, wenn ein **rekursives Teilbaum-Muster** (`**` oder `<präfix>/**`) seinen ganzen Teilbaum deckt (`.security/**`, `**/node_modules/**`, `dist/**`); nicht-teilbaum-deckende Muster (`<dir>/*`, `<dir>/`, Datei-Globs) prunen **nicht** (sonst False-Green), der Prune ist so beweisbar output-äquivalent zum Datei-Ausschluss; Scan-Wurzel nie geprunet. Prune **vor** dem Lesen des Verzeichnisinhalts ⇒ ein unlesbarer/sehr großer ausgeschlossener Teilbaum bricht den Scan nicht mehr ab; ein **nicht** ausgeschlossener unlesbarer Ordner bleibt fail-closed ([`AC-QA-02`](lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)). Realisiert die Verzeichnis-Absicht von [ADR-0018](../docs/plan/adr/0018-exclude-scan-scope.md) ([ADR-0025](../docs/plan/adr/0025-exclude-verzeichnis-prune.md)). Folgt [`AC-FA-CONF-001`](lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml) 0.20.0 (kein Lastenheft-Bump — Schärfung des Wie). |
