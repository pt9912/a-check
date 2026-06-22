# Lastenheft — a-check

**Version:** 0.5.0

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
dieselben sechs Regeln.

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

**Beschreibung:** Der Kern (innerste Schicht, z. B. `hexagon/core`) importiert
weder einen Adapter, einen Port oder eine Application-Schicht noch ein in der
Config als „Framework/Tech" deklariertes Symbol — die Domäne kennt nur sich
selbst. Verstoß ⇒ Befund mit Datei, Zeile und verletzter Regel.
[AC-FA-RULE-007](#ac-fa-rule-007--rolle-app-und-strenge-domain) verschärft dies
**kategorisch**: auch eine deklarierte Kante auf einen Port oder eine `app`-Schicht
hebt den Befund nicht auf.

**Akzeptanzkriterien:**

- **Happy:** Given ein Kern-Modul, das nur erlaubte Imports nutzt, when `a-check` läuft, then kein Befund für dieses Modul.
- **Boundary:** Given ein Kern-Modul, das nur andere Kern-/Domänen-Module (gleiche Rolle) und reine Standardbibliothek nutzt, when `a-check` läuft, then kein Befund.
- **Negative:** Given ein Kern-Modul, das einen Adapter, einen Port, eine `app`-Schicht oder ein Tech-Symbol importiert, when `a-check` läuft, then ein Befund (Grund `core-impurity`) und Exit-Code 1.

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

**Beschreibung:** Ports drücken die **Sprache des Kerns** aus und **dürfen
Domänen-/Kern-Typen referenzieren** (über eine deklarierte
`{from: ports, to: core}`-Kante) — das ist erwünscht, nicht nur geduldet, weil
ein Port die Domäne in seiner Signatur spricht. Sie importieren aber **keinen
Adapter** und **kein als Framework/Tech deklariertes Symbol**
(Persistence, Messaging, Vendor-Bibliotheken …) und tragen — sprachabhängig konfigurierbar — keine
implementierungs-/dialekt-spezifischen Konstrukte (z. B. Rust `impl`). *Prüf-Test:*
Ließe sich der Adapter komplett austauschen, ohne Port **und** Domäne zu ändern?
Wenn nein, leakt der Port Infrastruktur. Eine `ports → core`-Kante **ohne**
Deklaration bleibt eine Richtungsverletzung
([AC-FA-RULE-005](#ac-fa-rule-005--schicht-richtung-regel-wrong-direction)); das
Kern-/Adapter-Verbot der Domäne selbst regelt
[AC-FA-RULE-001](#ac-fa-rule-001--kern-reinheit-regel-core-impurity).

**Akzeptanzkriterien:**

- **Happy:** Given ein Port, der nur Domänen-/Kern-Typen referenziert (deklarierte `{from: ports, to: core}`-Kante), when `a-check` läuft, then kein Befund.
- **Boundary:** Given ein Port mit konfigurativ erlaubtem `ports → ports`-Re-Export, when `a-check` läuft, then kein Befund.
- **Negative:** Given ein Port, der einen **Adapter** oder ein **Tech-/Framework-Symbol** importiert oder ein verbotenes Konstrukt enthält, when `a-check` läuft, then ein Befund (`port-impurity`) und Exit-Code 1.

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

### AC-FA-RULE-006 — Schicht-Rollen (generische Regel-Anwendung)

**Generalisiert:** [AC-FA-RULE-001](#ac-fa-rule-001--kern-reinheit-regel-core-impurity) / [AC-FA-RULE-002](#ac-fa-rule-002--keine-lateralen-adapter-kanten-regel-lateral-adapter) / [AC-FA-RULE-004](#ac-fa-rule-004--port-disziplin-regel-port-impurity) (namens- → rollen-basiert).

**Beschreibung:** Die Reinheits-Regeln `core-impurity`, `port-impurity` (import-
**und** konstrukt-basiert) und `lateral-adapter` werden über die **Rolle** einer
Schicht angewandt, nicht über ihren Namen. Eine Schicht trägt optional eine Rolle
∈ {`domain`, `port`, `adapter`} (in [AC-FA-RULE-007](#ac-fa-rule-007--rolle-app-und-strenge-domain) um `app` erweitert); fehlt sie, wird sie aus konventionellen Namen
abgeleitet (`core`→`domain`, `ports`→`port`, `adapters`→`adapter`). Eine explizite
`role:` hat **Vorrang** vor der Inferenz; ein konventionell benannter Layer bekommt
zwangsläufig eine Rolle (Rückwärtskompatibilität). Eine Schicht ohne Rolle (weder
deklariert noch ableitbar) unterliegt nur den kanten-basierten Regeln
(`wrong-direction`/`tech-leak`). Rollen-Mapping: `domain`→`core-impurity`,
`port`→`port-impurity`, `adapter`→`lateral-adapter`. `lateral-adapter` feuert für
Importe zwischen **verschiedenen** `role: adapter`-Schichten (Layer-Identität,
namensunabhängig) und ist **kategorisch** — nur `adapter_sink` hebt auf, nicht
`allow`/`edges`. Innerhalb einer Schicht werden Adapter-Sub-Einheiten relativ zum
Glob-Präfix der Schicht unterschieden (ebenfalls namensunabhängig). Befund-**Namen**
bleiben unverändert.

**Akzeptanzkriterien:**

- **Happy:** Given zwei verschiedene Schichten mit `role: adapter`, when die eine die andere importiert (auch bei deklarierter `allow`-Kante), then ein Befund (`lateral-adapter`) — namensunabhängig und kategorisch.
- **Boundary:** Given eine Config mit klassischen Namen `core`/`ports`/`adapters` **ohne** `role`, when `a-check` läuft, then identisches Verhalten wie 0.2.0 (inkl. konstrukt-basierter `port-impurity` und Intra-`adapters`-Unterscheidung).
- **Negative:** Given (a) ein `role: domain`-Layer importiert einen `role: adapter`-Layer **oder** (b) ein `role: port`-Layer mit fremdem Namen (mit deklarierten `forbidden_constructs`) enthält ein verbotenes Konstrukt, when `a-check` läuft, then ein Befund (a) `core-impurity` bzw. (b) `port-impurity` und Exit-Code 1.

**Out-of-Scope:** `driving`/`driven`-Port-Subtypen; die Rolle `app` ist in [AC-FA-RULE-007](#ac-fa-rule-007--rolle-app-und-strenge-domain) ergänzt.

### AC-FA-RULE-007 — Rolle `app` und strenge `domain`

**Erweitert:** [AC-FA-RULE-006](#ac-fa-rule-006--schicht-rollen-generische-regel-anwendung) (Rollen-Menge um `app`). **Schärft:** [AC-FA-RULE-001](#ac-fa-rule-001--kern-reinheit-regel-core-impurity) (`core-impurity`).

**Beschreibung:** Das Rollen-Modell aus [AC-FA-RULE-006](#ac-fa-rule-006--schicht-rollen-generische-regel-anwendung) wird um die Rolle `app` (Application-/Use-Case-Schicht) erweitert; die Rolle `domain` wird verschärft. Rollen-Menge: {`domain`, `app`, `port`, `adapter`}. Die Namens-Inferenz ergänzt `application`→`app` und `app`→`app`; eine explizite `role:` behält Vorrang.

- **Rolle `app`:** darf `domain` **und** `port` importieren (Use-Cases orchestrieren über Ports), aber **keine** Adapter-/Tech-Typen — Verstoß ⇒ Befund `app-impurity` (neu). Die Schicht-**Richtung** (`app → domain`, `app → port`) bleibt kanten-geregelt (`wrong-direction`); die **Reinheit** ist **kategorisch**.
- **Rolle `domain` (verschärft):** die innerste Schicht ist die strengste — ein Import auf eine `app`-, `port`- oder `adapter`-Schicht **oder** ein `tech`-Muster ist `core-impurity`, **kategorisch** (auch bei deklarierter Kante). Rollenlose Ziel-Schichten bleiben kanten-geregelt. Bisher war `domain → port` nur kanten-geregelt; jetzt gilt die harte Invariante „Domäne kennt keine Ports".

Rollen-Mapping (Ergänzung): `app`→`app-impurity`. Befund-**Namen** der übrigen Regeln bleiben unverändert.

**Akzeptanzkriterien:**

- **Happy:** Given eine `role: app`-Schicht mit deklarierten Kanten `app → domain` und `app → port`, when sie eine `domain`- und eine `port`-Schicht importiert, then kein Befund.
- **Negative (app):** Given eine `role: app`-Schicht, when sie eine `adapter`-Schicht **oder** ein `tech`-Muster importiert (auch bei deklarierter Kante), then ein Befund (`app-impurity`) und Exit-Code 1.
- **Negative (domain):** Given eine `role: domain`-Schicht, when sie eine `port`- (oder `app`-/`adapter`-)Schicht importiert (auch bei deklarierter Kante), then ein Befund (`core-impurity`) und Exit-Code 1.
- **Boundary:** Given eine Config ohne `role:` und ohne Layer `application`/`app` (klassisch `core`/`ports`/`adapters`), when `a-check` läuft, then identisches Verhalten wie 0.4.0.

**Out-of-Scope:** `driving`/`driven`-Port-Subtypen; feinere `app`-interne Struktur.

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
erlaubten Kanten, die Tech→Adapter-Zuordnungen und die gemeinsame Adapter-Senke. Ein `layers`-Eintrag
ist **entweder** eine Glob-Liste (`name: [globs]`, Rolle per Namens-Inferenz)
**oder** ein Objekt `{globs: [...], role: domain|port|adapter}`
([AC-FA-RULE-006](#ac-fa-rule-006--schicht-rollen-generische-regel-anwendung)).
Striktes Decoding, fail-closed (Exit 2 bei unbekanntem Schlüssel — auch im Objekt).

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
| 0.2.0 | 2026-06-22 | `AC-FA-RULE-004` neu gefasst: Ports **dürfen** Domänen-/Kern-Typen referenzieren (Sprache des Kerns; `ports → core` per deklarierter Kante), `port-impurity` trennt scharf gegen Adapter-/Tech-Importe. Motiviert durch die Vier-Repo-Evidenz (b-cad/d-migrate-Ports referenzieren die Domäne). |
| 0.3.0 | 2026-06-22 | Neu `AC-FA-RULE-006` (Schicht-Rollen): die Reinheits-Regeln dispatchen über eine Layer-Rolle (`domain`/`port`/`adapter`, aus `role:` oder Namens-Inferenz) — generalisiert `AC-FA-RULE-001`/`AC-FA-RULE-002`/`AC-FA-RULE-004` namens-unabhängig (welle-10a). `AC-FA-CONF-001`-Schema: `layers`-Eintrag als Glob-Liste **oder** `{globs, role}`. |
| 0.4.0 | 2026-06-22 | `AC-FA-RULE-006`: `lateral-adapter` jetzt **vollständig** namensunabhängig — Adapter-Sub-Einheiten werden relativ zum Schicht-Glob-Präfix unterschieden (statt am Literal `adapters`); `adapterSeg`-Generalisierung aus dem Out-of-Scope eingelöst (welle-10b). |
| 0.5.0 | 2026-06-22 | Neu `AC-FA-RULE-007` (Rolle `app` + strenge `domain`): `app` darf `domain`+`port`, aber keinen Adapter/Tech (neuer Befund `app-impurity`); `domain` verschärft — Import auf `app`/`port`/`adapter`/Tech ist `core-impurity`, kategorisch („Domäne kennt keine Ports"). Erweitert `AC-FA-RULE-006`, schärft `AC-FA-RULE-001` (welle-10b). |
