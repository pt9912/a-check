# Lastenheft — a-check

**Version:** 0.1.0

**Status:** Draft

**Autor:** pt9912, **Datum:** 2026-06-20.

---

## 1. Zweck und Geltungsbereich

`a-check` ist ein Kommandozeilen-Tool, das die **hexagonale
Schicht-Architektur** eines Repositories durchsetzt — Kern-Reinheit,
Adapter-Kapselung, Port-Disziplin und die Import-/Schicht-Richtung —
**sprachübergreifend**, gesteuert über eine Konfigurationsdatei. Es
konsolidiert die handgepflegten `arch-check.sh`-Skripte der
Schwester-Repositories, die heute dieselben Hexagon-Regeln je Repo neu
erfinden: C++ über `#include`-Heuristik (`b-cad`), Go über `go list`
(`d-check`), Rust über `use`-Heuristik (`grid-guide`), Kotlin über
Gradle-Modulgrenzen (`d-migrate`) — vier Sprachen, vier Mechanismen,
dieselben fünf Regeln.

Das Tool wird als Docker-Image über GHCR verteilt, per `.a-check.yml`
pro Repo konfiguriert und über ein bereitgestelltes `a-check.mk` als
`make a-check`-Gate eingebunden — ein Image, ein Update-Pfad,
repo-spezifische Schicht-/Tech-Regeln per Config statt per Skript-Kopie.
Es ist das **Architektur-Gegenstück zu `d-check`** (Doku-Referenzen):
dieselbe Gründungslogik (eine Familie driftender Skripte durch ein
Werkzeug ersetzen), eine Abstraktionsebene höher.

**Out of Scope (Produkt):** `a-check` ersetzt keine sprach-eigene,
compile-time durchgesetzte Modulgrenze (z. B. Gradle-Module in
`d-migrate`), sondern ergänzt sie um die *fein­granularen*
Fitness-Functions, die der Compiler nicht abdeckt (laterale
Adapter-Kanten, Port-Disziplin). Es ist eine **Heuristik** auf
Import-Ebene, kein vollständiger Sprach-Parser (siehe `AC-QA-02`).

## 2. Stakeholder

| Stakeholder | Rolle | Erwartung |
|---|---|---|
| Repo-Maintainer (pt9912) | Auftraggeber | Ein gepflegtes Architektur-Gate statt N driftender `arch-check.sh`-Kopien; Regeländerung wirkt überall |
| Hexagon-Repos (`b-cad` C++, `d-check` Go, `grid-guide` Rust, `d-migrate` Kotlin) | Konsument | Ein Docker-Step + `.a-check.yml`, der ihre bestehenden Regeln deterministisch erzwingt |
| CI-Pipelines | Konsument | `make a-check` mit stabilem Exit-Code; netzloser, hermetischer Lauf |
| AI-Agenten (Harness-Sensorik) | Konsument | Maschinenlesbarer Architektur-Sensor als Gate, analog `d-check` |

## 3. Funktionale Anforderungen

> **Schema-Konvention.** Funktionale Anforderungen verwenden Bereichskürzel:
> `AC-FA-<BEREICH>-<NNN>`. Bereiche: `RULE` (Hexagon-Regeln), `EXTRACT`
> (Import-Extraktion je Sprache), `CLI` (Aufruf/Ausgabe), `CONF`
> (Konfiguration), `DIST` (Distribution).

### AC-FA-RULE-001 — Kern-Reinheit (Regel `core-impurity`)

**Beschreibung:** Der Kern (konfigurierte Schicht, z. B. `hexagon/core`)
importiert weder einen Adapter noch ein in der Config als „Framework/Tech"
deklariertes Symbol. Verstoß ⇒ Befund mit Datei, Zeile und verletzter Regel.

**Akzeptanzkriterien:**

- **Happy:** Given ein Kern-Modul, das nur erlaubte Imports nutzt, when `a-check` läuft, then kein Befund für dieses Modul.
- **Boundary:** Given ein Kern-Modul, das einen `driver-common`-artigen, in der Config erlaubten gemeinsamen Port nutzt, when `a-check` läuft, then kein Befund (erlaubte Kante).
- **Negative:** Given ein Kern-Modul, das einen Adapter oder ein Tech-Symbol importiert, when `a-check` läuft, then ein Befund (Grund `core-impurity`) und Exit-Code 1.

**Out-of-Scope:** transitive Import-Analyse über Modulgrenzen hinweg in 0.1.0 (nur direkte Imports).

### AC-FA-RULE-002 — Keine lateralen Adapter-Kanten (Regel `lateral-adapter`)

**Beschreibung:** Ein Adapter importiert keinen anderen Adapter, außer einer
in der Config benannten gemeinsamen Senke (z. B. `driver-common`). Erfasst die
in `d-migrate` real existierende, heute nur per Review erzwungene Regel.

**Akzeptanzkriterien:**

- **Happy:** Given ein Adapter ohne Fremd-Adapter-Import, when `a-check` läuft, then kein Befund.
- **Boundary:** Given ein Adapter, der die konfigurierte gemeinsame Senke importiert, when `a-check` läuft, then kein Befund.
- **Negative:** Given Adapter A importiert Adapter B (nicht die Senke), when `a-check` läuft, then ein Befund (`lateral-adapter`) und Exit-Code 1.

**Out-of-Scope:** Zyklen-Erkennung über drei oder mehr Adapter (eigenes Re-Eval).

### AC-FA-RULE-003 — Tech-Kapselung (Regel `tech-leak`)

**Beschreibung:** Ein in der Config einem Adapter zugeordnetes Framework/Tech
(z. B. `*.hxx` → Geometrie-Adapter, `sqlite3*` → Persistenz-Adapter, `Qt` →
UI-Adapter, `net/http` → http-Adapter) erscheint **nur** in seinem Adapter (und
ggf. der Composition Root).

**Akzeptanzkriterien:**

- **Happy:** Given ein Tech-Symbol nur in seinem zugeordneten Adapter, when `a-check` läuft, then kein Befund.
- **Boundary:** Given dasselbe Symbol in der konfigurierten Composition Root, when `a-check` läuft, then kein Befund (deklarierte Ausnahme).
- **Negative:** Given das Symbol außerhalb seines Adapters, when `a-check` läuft, then ein Befund (`tech-leak`) und Exit-Code 1.

**Out-of-Scope:** semantische Unterscheidung gleichnamiger, aber framework-fremder Symbole (Heuristik-Grenze, siehe `AC-QA-02`).

### AC-FA-RULE-004 — Port-Disziplin (Regel `port-impurity`)

**Beschreibung:** Ports sind reine Abstraktionen und Dependency-Senke: sie
importieren weder Adapter noch Kern und tragen — sprachabhängig konfigurierbar
— keine implementierungs-/dialekt-spezifischen Konstrukte (z. B. Rust `impl`,
dialekt-typisierte Felder).

**Akzeptanzkriterien:**

- **Happy:** Given ein Port mit nur Abstraktions-Definitionen, when `a-check` läuft, then kein Befund.
- **Boundary:** Given ein Port mit konfigurativ erlaubtem Re-Export, when `a-check` läuft, then kein Befund.
- **Negative:** Given ein Port, der einen Adapter importiert oder ein verbotenes Konstrukt enthält, when `a-check` läuft, then ein Befund (`port-impurity`) und Exit-Code 1.

**Out-of-Scope:** Typ-Inferenz über das deklarierte Pattern hinaus.

### AC-FA-RULE-005 — Schicht-Richtung (Regel `wrong-direction`)

**Beschreibung:** Die in der Config deklarierten Schicht-Kanten
(`core ← ports ← adapters`, ggf. weitere) sind einbahnig; eine Kante entgegen
der Richtung ist ein Befund.

**Akzeptanzkriterien:**

- **Happy:** Given Imports nur entlang der erlaubten Richtung, when `a-check` läuft, then kein Befund.
- **Boundary:** Given eine in der Config explizit erlaubte Sonderkante, when `a-check` läuft, then kein Befund.
- **Negative:** Given eine Kante gegen die deklarierte Richtung, when `a-check` läuft, then ein Befund (`wrong-direction`) und Exit-Code 1.

**Out-of-Scope:** automatische Ableitung der Schichten ohne Config.

### AC-FA-EXTRACT-001 — Sprach-Backends für die Import-Extraktion

**Beschreibung:** Pro Sprache liefert ein Backend die Menge „welche
Symbole/Module importiert diese Datei" — text-heuristisch über konfigurierbare
Muster: C++ (`#include`), Go (`import`), Rust (`use`/`extern crate`), Kotlin
(`import`). Das Backend wird über die Config (Sprache + Datei-Globs) gewählt.

**Akzeptanzkriterien:**

- **Happy:** Given eine Go-Datei mit zwei Imports, when das Go-Backend läuft, then liefert es genau diese zwei Importpfade.
- **Boundary:** Given eine Rust-Alias-Form (`use tauri as t;`), when das Rust-Backend läuft, then wird `tauri` erkannt.
- **Negative:** Given eine in einem Kommentar/String stehende Import-ähnliche Zeile, when das Backend läuft, then wird sie nicht als Import gewertet (oder als bewusste, dokumentierte Heuristik-Grenze gemeldet — `AC-QA-02`).

**Out-of-Scope:** vollständiges AST-Parsing; Toolchain-gestützte Backends (`go list`, Bytecode) sind ein opt-in-Re-Eval, nicht 0.1.0.

### AC-FA-CLI-001 — Aufruf, Scan-Wurzel und Exit-Codes

**Beschreibung:** `a-check [pfad]` prüft das Repo unter `pfad` (Default `/src`
im Container) gegen die `.a-check.yml`. Exit-Codes: `0` kein Befund, `1`
mindestens ein Befund, `2` Nutzungs-/Konfigurationsfehler. Befunde auf stdout,
Zusammenfassung auf stderr (analog `d-check`).

**Akzeptanzkriterien:**

- **Happy:** Given ein konformes Repo, when `a-check` läuft, then Exit-Code 0.
- **Boundary:** Given ein read-only gemountetes Repo, when `a-check` läuft, then vollständige Prüfung ohne Schreibzugriff.
- **Negative:** Given eine fehlende/ungültige `.a-check.yml`, when `a-check` läuft, then Exit-Code 2 mit Zeilenangabe.

**Out-of-Scope:** Auto-Fix/Reparatur von Architekturverstößen (es gibt keinen deterministisch ableitbaren Fix).

### AC-FA-CONF-001 — Konfigurationsdatei `.a-check.yml`

**Beschreibung:** `.a-check.yml` deklariert: die Sprache(n) + Datei-Globs je
Schicht, die Schichten (`core`/`ports`/`adapters`/…) mit Pfad-Mustern, die
erlaubten Kanten, die Tech→Adapter-Zuordnungen und die gemeinsame Adapter-Senke.
Striktes Decoding, fail-closed (Exit 2 bei unbekanntem Schlüssel).

**Akzeptanzkriterien:**

- **Happy:** Given eine gültige `.a-check.yml`, when `a-check` läuft, then werden die deklarierten Regeln angewandt.
- **Boundary:** Given eine Config ohne optionale Tech-Zuordnungen, when `a-check` läuft, then laufen nur die Schicht-/Lateral-Regeln (kein `tech-leak`).
- **Negative:** Given ein Tippfehler im Schlüssel, when `a-check` läuft, then Exit-Code 2 (kein stiller Default).

**Out-of-Scope:** Vererbung/Includes zwischen Config-Dateien.

### AC-FA-DIST-001 — Distribution: Image, `--print-mk`, `a-check.mk`

**Beschreibung:** `a-check` wird als GHCR-Image (distroless/static,
digest-gepinnt) verteilt. `a-check --print-config` gibt ein kommentiertes
`.a-check.yml`-Gerüst aus; `a-check --print-mk` gibt ein `a-check.mk` mit dem
**aktuell digest-gepinnten** Image und einem `a-check`-Target aus. Konsumenten
`include a-check.mk` und liefern `.a-check.yml` — keine Skript-Kopie.

**Akzeptanzkriterien:**

- **Happy:** Given das Image, when `a-check --print-mk` läuft, then ein `include`-bares Makefile-Fragment mit digest-gepinntem `A_CHECK_IMAGE` und `a-check`-Target auf stdout.
- **Boundary:** Given `a-check --print-config`, when es läuft, then ein dekodierbares `.a-check.yml`-Gerüst, **schreibt nichts** (read-only).
- **Negative:** Given `--print-mk` mit einem zusätzlichen unbekannten Flag, when aufgerufen, then Exit-Code 2.

**Out-of-Scope:** Nicht-Docker-Distribution (Binary-Releases) in 0.1.0.

## 4. Nichtfunktionale Anforderungen

### AC-QA-01 — Determinismus

Identische Eingabe (Repo-Stand + `.a-check.yml` + Image-Digest) ⇒
byte-identische Ausgabe und identischer Exit-Code. Befunde sind stabil sortiert.

### AC-QA-02 — Hermetik und ehrliche Heuristik-Grenze

Der Scan ist **text-basiert** (keine Sprach-Toolchain), läuft **netzlos**
(`--network none`) im distroless/static-Image und schreibt nie ins geprüfte
Repo. Die Heuristik-Grenzen (z. B. ein framework-fremdes `Queue.h` unter einem
`Q[A-Za-z]`-Muster) werden **dokumentiert** statt verschwiegen; eine
Allowlist/Marker-Ausnahme ist konfigurierbar.

### AC-QA-03 — Reproduzierbarkeit

Image und ausgelieferte `a-check.mk` referenzieren einen `@sha256:`-Digest;
Pin-Hebung ist ein bewusster Commit (analog der Pin-Politik der
Konsumenten-Repos).

## 7. Historie

| Version | Datum | Änderung |
|---|---|---|
| 0.1.0 | 2026-06-20 | Erstfassung (Bootstrap): Zweck/Inventur, fünf universelle Hexagon-Regeln (`AC-FA-RULE-001…005`), Sprach-Extraktion, CLI, Config, Distribution (`--print-mk`/`a-check.mk`); NFAs Determinismus/Hermetik/Reproduzierbarkeit. |
