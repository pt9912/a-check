# Lastenheft — a-check

**Version:** 0.24.0

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
dieselben sieben Regeln.

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

Regeln dieser Sektion: ID-Schema `AC-FA-<BEREICH>-<NNN>` mit Bereichssegment, weil dieses Repo
den Zählraum je Bereich führt (Baseline-Regelwerk `grundlagen-source-precedence.md` §Vergabe und
§ID-Schema als Klammer). Jede Anforderung trägt drei Pfade — Happy · Boundary · Negative — plus
Out-of-Scope (`modul-03-spec.md` §Ziel-Form: Akzeptanzkriterium); `make verify` prüft das für neue
Kennungen. Eine Anforderung, deren Bedarf ersatzlos entfällt, wird **nicht gelöscht**, sondern
trägt den Vermerk *zurückgezogen* im Titel; die Nummer bleibt vergeben.

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
Innerhalb einer Schicht werden Adapter-Sub-Einheiten relativ zum
Schicht-Glob-Präfix unterschieden; enthält der Pfad-Rest nach dem Präfix
**kein weiteres Verzeichnis**, entscheidet die **Blatt-Klassifikation**: ein
**datei-förmiges** Blatt (letztes Segment enthält `.`) gehört zur
**Root-Sub-Einheit `''`** — Sub-Einheiten sind, wie Schichten, Verzeichnisse,
keine Dateinamen —, ein **verzeichnis-förmiges** Blatt (ohne `.`, z. B. ein
Go-Paket-Pfad) **ist** die Sub-Einheit. Importe zwischen Root-Dateien
derselben Schicht sind damit keine lateralen Kanten; Root ↔ Unterverzeichnis
und verschiedene Unterverzeichnisse bleiben lateral,
Cross-Layer-Adapter-Importe bleiben kategorisch.

**Akzeptanzkriterien:**

- **Happy:** Given ein Adapter ohne Fremd-Adapter-Import, when `a-check` läuft, then kein Befund.
- **Boundary:** Given ein Adapter, der die konfigurierte gemeinsame Senke importiert, when `a-check` läuft, then kein Befund.
- **Negative:** Given Adapter A importiert Adapter B (nicht die Senke), when `a-check` läuft, then ein Befund (`lateral-adapter`) und Exit-Code 1.
- **Happy (Root-Sub-Einheit):** Given zwei Dateien direkt im Root desselben Adapter-Layers (`x.cpp` importiert `x.h`), when `a-check` läuft, then **kein** `lateral-adapter`.
- **Boundary (Root ↔ Unterverzeichnis):** Given eine Root-Datei, die eine Datei eines Unterverzeichnisses derselben Schicht importiert (oder umgekehrt), when `a-check` läuft, then ein Befund (`lateral-adapter`) — verschiedene Sub-Einheiten.
- **Negative (Bestand):** Given Importe zwischen zwei Unterverzeichnissen derselben Schicht oder zwischen zwei Adapter-Layern, when `a-check` läuft, then `lateral-adapter` wie bisher (kategorisch; `adapter_sink`-Ausnahme unverändert).
- **Boundary (Verzeichnis-Blatt/Go-Paket):** Given ein externes Go-Testpaket, das sein **eigenes** Paket importiert (Kandidaten-Blatt = eigenes Sub-Einheiten-Verzeichnis), when `a-check` läuft, then kein Befund; Given ein Import eines **fremden** Paket-Verzeichnisses derselben Schicht, then `lateral-adapter` (die Blatt-Klassifikation blendet die Cross-Paket-Erkennung nicht).

**Out-of-Scope:** Zyklen-Erkennung über drei oder mehr Adapter (eigenes Re-Eval); datei-granulare Sub-Einheiten als Opt-in (Re-Eval); endungslose Datei-Specifier (z. B. TypeScript `./b` ohne Endung) sind textuell nicht von Verzeichnis-Blättern unterscheidbar und gelten als Sub-Einheit — dokumentierte Heuristik-Grenze (`AC-QA-02`).

### AC-FA-RULE-003 — Tech-Kapselung (Regel `tech-leak`)

**Beschreibung:** Ein in der Config einem **oder mehreren** Adapter(n)
zugeordnetes Framework/Tech (z. B. `*.hxx` → Geometrie-Adapter, `sqlite3*` →
Persistenz-Adapter, `Qt` → UI-Adapter, `net/http` → http-Adapter, `yaml` →
Config- **und** Report-Adapter) erscheint **nur** in seinem/seinen Adapter(n)
— `adapter` als nicht-leerer Pfad oder Pfad-**Liste** (leere Liste, leerer
Listen-Eintrag sowie leerer/fehlender Skalar → Exit 2; der leere Adapter war
vormals ein stiller Never-Leak-Eintrag), das Symbol ist in **jedem** gelisteten
Adapter erlaubt — und in der Composition Root, sofern der Eintrag deren Ausnahme
nicht per **`composition_root: forbid`** abschaltet (Default `allow` =
bisheriges Verhalten; das Verbot betrifft nur die `tech`-Ausnahme — die
Schicht-Regel-Ausnahme der Composition Root bleibt unberührt). Das Muster matcht
das importierte Symbol als **Substring** (Default) **oder** — via `match: regex`
— als **RE2-Regex** (unverankerter Suchlauf); so wird ein nur als Muster
fassbares Tech wie Qt (`Q[A-Za-z]`) ausdrückbar. Treffen mehrere Muster dasselbe
Symbol, greift der **in Deklarationsreihenfolge erste** (deterministisch; kein
„längster Präfix" für `tech`).

**Akzeptanzkriterien:**

- **Happy:** Given ein Tech-Symbol nur in seinem zugeordneten Adapter, when `a-check` läuft, then kein Befund.
- **Boundary:** Given dasselbe Symbol in der konfigurierten Composition Root und **kein** `composition_root: forbid` am Eintrag, when `a-check` läuft, then kein Befund (deklarierte Ausnahme, Default).
- **Negative:** Given das Symbol außerhalb seines Adapters, when `a-check` läuft, then ein Befund (`tech-leak`) und Exit-Code 1.
- **Regex:** Given ein `match: regex`-Muster, das ein Symbol trifft, das außerhalb seines Adapters liegt, when `a-check` läuft, then ein Befund (`tech-leak`) und Exit-Code 1; liegt es nur im Adapter/der Composition Root, kein Befund.
- **Präzedenz:** Given mehrere `tech`-Muster (substring und/oder regex), die dasselbe Symbol treffen, when `a-check` läuft, then greift der in Deklarationsreihenfolge erste Treffer (deterministisch).
- **Mehr-Adapter:** Given ein `tech`-Eintrag mit Adapter-**Liste** und das Symbol in einem der gelisteten Adapter, when `a-check` läuft, then kein Befund; liegt es außerhalb **aller** gelisteten Adapter, then ein Befund (`tech-leak`) und Exit-Code 1.
- **Composition-Root-Verbot:** Given ein `tech`-Eintrag mit `composition_root: forbid` und das Symbol in der Composition Root, when `a-check` läuft, then ein Befund (`tech-leak`) und Exit-Code 1; die Ausnahme der **Schicht**-Regeln für die Composition Root bleibt unberührt.
- **Rückwärtskompat:** Given ein `tech`-Eintrag **ohne** `match`, **ohne** `composition_root`-Feld und mit **nicht-leerem** Skalar-`adapter`, when `a-check` läuft, then Substring- und Erlaubnis-Semantik wie bisher (byte-identische Ausgabe).
- **Negative (leerer Adapter):** Given ein `tech`-Eintrag mit leerem oder fehlendem `adapter` (Skalar wie Liste), when `a-check` lädt, then Exit-Code 2 — statt des vormals stillen Never-Leak-Eintrags (das Muster meldete nie, falsch-grün).

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

**Out-of-Scope:** feinere `app`-interne Struktur; die `driving`/`driven`-**Richtung** als orthogonales Attribut (kein Port-Subtyp) liefert [AC-FA-RULE-008](#ac-fa-rule-008--driving-driven-port-richtung-regel-port-direction-mismatch).

### AC-FA-RULE-008 — Driving-Driven-Port-Richtung (Regel `port-direction-mismatch`)

**Verfeinert:** [AC-FA-RULE-006](#ac-fa-rule-006--schicht-rollen-generische-regel-anwendung) um eine **orthogonale** Richtungs-Dimension (`direction`).

**Beschreibung:** Eine `port`- oder `adapter`-Schicht trägt optional eine Richtung
`direction` ∈ {`driving`, `driven`}. `driving` = primär/inbound (Use-Case-Schnittstelle,
vom Treiber-Adapter aufgerufen); `driven` = sekundär/outbound (vom Kern/App definiert,
vom getriebenen Adapter implementiert). Die Richtung ist **orthogonal** zur Rolle: die
Reinheits-Regeln (`core-impurity`/`app-impurity`/`port-impurity`/`lateral-adapter`)
bleiben rollen-basiert unverändert. Neue Regel `port-direction-mismatch`: ein
`role: adapter` mit Richtung X, der eine `role: port`-Schicht mit Richtung Y (Y ≠ X,
**beide gesetzt**) importiert, ist ein Befund (**kategorisch** — `edges`/`allow` heben
nicht auf, wie `lateral-adapter`; nur `composition_root` befreit): ein Treiber-Adapter
spricht nur `driving`-Ports, ein getriebener Adapter nur `driven`-Ports. Schichten **ohne**
`direction` unterliegen der Regel **nicht** (Rückwärtskompatibilität: ohne Deklaration
ändert sich nichts). Die `app`-Schicht ist richtungs-agnostisch (nutzt `driven`-Ports,
implementiert `driving`-Ports) und wird nicht erfasst. Befund-**Namen** der übrigen
Regeln bleiben unverändert.

**Akzeptanzkriterien:**

- **Happy:** Given ein `role: adapter`, `direction: driving`, when er eine `role: port`, `direction: driving`-Schicht importiert, then kein Befund.
- **Negative:** Given ein `role: adapter`, `direction: driving`, when er eine `role: port`, `direction: driven`-Schicht importiert, then ein Befund (`port-direction-mismatch`) und Exit-Code 1.
- **Negative (kategorisch):** Given ein `role: adapter`, `direction: driving` **und eine deklarierte `allow`-Kante** auf die `role: port`, `direction: driven`-Schicht, when er sie importiert, then **dennoch** ein Befund (`port-direction-mismatch`) und Exit-Code 1 — die Richtung ist nicht über `edges`/`allow` aufhebbar.
- **Boundary:** Given Schichten **ohne** `direction` (klassisch `role: port`/`adapter`), when `a-check` läuft, then identisches Verhalten wie 0.5.0.

**Out-of-Scope:** Auto-Inferenz der Richtung aus Namen/Pfad (`driving`/`driven` im Pfad); Richtungs-Regeln zwischen Ports untereinander — späteres Inkrement.

### AC-FA-RULE-009 — Slice-Isolation (Regel `lateral-slice`)

**Verfeinert:** [AC-FA-RULE-006](#ac-fa-rule-006--schicht-rollen-generische-regel-anwendung) um die
**Isolation von Use-Case-Slices** innerhalb der `app`-Rolle (Vertical-Slice-Architektur).

**Beschreibung:** Innerhalb der `app`-Rolle bildet jede über ein **eigenes Schicht-Glob** deklarierte
Use-Case-Slice eine Isolationseinheit. Eine `app`-Datei, die eine **andere** Slice **derselben
`app`-Schicht** importiert (ein `app`-Ziel, das über ein *anderes* `app`-Glob **derselben Schicht**
auflöst), ist ein Befund `lateral-slice`. **Kategorisch innerhalb der Schicht** — `edges`/`allow` heben
nicht auf (wie [`lateral-adapter`](#ac-fa-rule-002--keine-lateralen-adapter-kanten-regel-lateral-adapter)
für Adapter); nur `composition_root` befreit. **Getrennte `app`-Schichten** (verschiedene Layer, z. B.
`services` und `services_geo`) sind dagegen **edge-regiert** — ein Import zwischen ihnen ist `wrong-direction`
oder per Kante erlaubt, **nicht** `lateral-slice` (klassisch-hexagonale Sub-Services bleiben legitim).
Use-Case-Slices teilen fachliche Verträge über **Ports**, nicht über direkten Slice-zu-Slice-Code. Trägt die `app`-Rolle **nur ein** Glob (keine per-Slice-Trennung),
unterliegt sie der Regel **nicht** — Slice-Isolation ist opt-in über per-Slice-Globs
(rückwärtskompatibel). Damit ein `app`-Ziel überhaupt als Slice auflöst, muss sein Glob als
**Import-Ziel auflösbar** sein (literales Präfix); die Grenze der Ziel-Auflösung ist
[AC-QA-02](#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze).

**Akzeptanzkriterien:**

- **Happy:** Given zwei Slices mit je eigenem `app`-Glob und **keinen** gegenseitigen Importen, when `a-check` läuft, then kein Befund.
- **Boundary:** Given eine `app`-Rolle mit **einem einzigen** Glob (keine per-Slice-Trennung), when `a-check` läuft, then **kein** `lateral-slice` (identisch zum Vor-Verhalten).
- **Negative:** Given eine `app`-Datei in Slice A, die ein Symbol aus Slice B importiert, when `a-check` läuft, then ein Befund (`lateral-slice`) und Exit-Code 1.
- **Negative (kategorisch):** Given denselben Cross-Slice-Import **und eine deklarierte `allow`-Kante**, when `a-check` läuft, then **dennoch** ein Befund (`lateral-slice`) — nicht über `edges`/`allow` aufhebbar.

**Out-of-Scope:** geteilter `app`-Code zwischen Slices (analog `adapter_sink` ein künftiges `app_sink`); Slice-Isolation für andere Rollen als `app` (Adapter deckt `lateral-adapter`); automatische Slice-Erkennung ohne per-Slice-Glob.

### AC-FA-RULE-010 — Port-Lokalität (Regel `port-locality`)

**Verfeinert:** [AC-FA-RULE-004](#ac-fa-rule-004--port-disziplin-regel-port-impurity) um die
**Lokalitäts-/Sichtbarkeits-Ebene** eines Ports (HexSlice „so lokal wie möglich, so gemeinsam wie nötig").

**Beschreibung:** Ein Port hat ein **Scope-Verzeichnis**, das sich aus seiner Lage ergibt — der
Verzeichnis-Teilbaum, der seinen Port-Ordner besitzt: **use-case-lokal ⊂ business-area ⊂ app-weit**.
Eine `app`-Datei, die einen Port importiert, dessen Scope-Verzeichnis sie **nicht enthält** (die
importierende Datei liegt **außerhalb** des Port-Scope), ist ein Befund `port-locality`. Die Regel
greift **nur für im Application-Baum geschachtelte Ports** — der Port-Scope muss Vorfahr einer
`app`-Slice sein (HexSlice: `ports/` liegt *innerhalb* der Slice/Business-Area). Liegt der Port als
**Geschwister** der `app`-Schicht (klassisches Hexagonal, `hexagon/ports/**` neben `hexagon/services/**`),
gibt es keine per-Slice-Lokalität und die Regel bleibt **inert**. **Kategorisch** — `edges`/`allow` heben
nicht auf; nur `composition_root` befreit. Erfasst werden **nur `app`-Importeure**;
ein Adapter, der einen Port **implementiert** (Implementierungs-Beziehung, richtungs-/edge-regiert,
[AC-FA-RULE-008](#ac-fa-rule-008--driving-driven-port-richtung-regel-port-direction-mismatch)), ist
**nicht** erfasst. Der Scope ist **pfad-abgeleitet** (keine Deklaration); die Ableitung setzt einen als
Import-Ziel auflösbaren Port-Glob voraus (literales Präfix,
[AC-QA-02](#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)).

**Akzeptanzkriterien:**

- **Happy:** Given eine `app`-Datei, die ihren **eigenen** slice-lokalen Port und einen **business-area**-Port importiert, in deren Scope sie liegt, when `a-check` läuft, then kein Befund.
- **Boundary:** Given einen **app-weiten** Port (Scope = die App-Wurzel), when eine beliebige `app`-Datei ihn importiert, then kein Befund.
- **Negative:** Given eine `app`-Datei in Slice A, die den **slice-lokalen** Port einer **fremden** Slice B importiert, when `a-check` läuft, then ein Befund (`port-locality`) und Exit-Code 1.

**Out-of-Scope:** Port-Lokalität für **Adapter**-Importeure (Implementierungs-Beziehung); eine erzwungene explizite Scope-Deklaration in der Config (der Scope ist pfad-abgeleitet); Lokalität über Business-Area-Grenzen als eigene Ebene.

### AC-FA-RULE-011 — Konstrukt-Monopol (Regel `construct-leak`)

**Verallgemeinert:** die Scoping-Mechanik von
[AC-FA-RULE-003](#ac-fa-rule-003--tech-kapselung-regel-tech-leak) (`tech`) von **extrahierten
Import-Symbolen** auf **Roh-Quelltext**. **Abgrenzung:** die schichtgebundenen
`forbidden_constructs` aus [AC-FA-RULE-004](#ac-fa-rule-004--port-disziplin-regel-port-impurity)
(Befund `port-impurity`) bleiben unverändert — anderer Geltungsbereich (Schicht statt Zone),
anderer Befund.

**Beschreibung:** Ein optionaler `constructs`-Block deklariert Roh-Text-Muster, die **nur** in
der/den ihnen zugeordneten **Zone(n)** vorkommen dürfen; jedes Vorkommen außerhalb ist ein Befund
`construct-leak` (Datei, Zeile, Muster, erlaubte Zone[n]) und Exit-Code 1. Damit werden
Konstrukte prüfbar, die **keine Import-Zeile** sind und die die Import-Extraktion darum
grundsätzlich nicht sieht — etwa ein Funktionsaufruf (`dlopen(`), der über einen transitiven
Header oder einen lokalen Prototyp ohne eigenen Include auskommt. Die Zone wird wie bei
[AC-FA-RULE-003](#ac-fa-rule-003--tech-kapselung-regel-tech-leak) deklariert: `adapter` als
nicht-leerer Pfad **oder** Pfad-**Liste** (das Muster ist in **jeder** gelisteten Zone erlaubt),
`match: substring` (Default) **oder** `regex` (RE2), `composition_root: allow` (Default) oder
`forbid` je Eintrag. Die Prüfung ist **scan-weit**: sie gilt für **jede** gescannte Datei — auch
für Dateien, die in **keinem** `layers`-Glob liegen —, während `exclude` wie bisher **vor** dem
Scan greift; die Composition Root ist ausgenommen, sofern der Eintrag nicht `composition_root:
forbid` deklariert. Gematcht wird auf **derselben Quell-Vorbereitung** wie `forbidden_constructs` — keine eigene:
in den **C-Syntax-Sprachen** ist sie kommentar-bereinigt, ein Treffer allein in einem Kommentar
ist dort **kein** Befund (eine bewusste, ausgewiesene Divergenz zu einer `grep`-Referenz, die
Kommentare mitsieht). **Python** wird bewusst nicht C-gestrippt
([AC-FA-EXTRACT-001](#ac-fa-extract-001--sprach-backends-für-die-import-extraktion)); dort meldet
auch ein Treffer im `#`-Kommentar — ausgewiesene Grenze
([AC-QA-02](#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)), weil ein `#`-Strip ein `#` im
String-Literal mitverschlucken und einen echten Treffer **verbergen** würde (False-Green wiegt
schwerer als Falsch-Rot). Text-Heuristik, kein Parser: Treffer in String-Literalen bleiben in
allen Sprachen die dokumentierte Grenze
([AC-QA-02](#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)).

**Akzeptanzkriterien:**

- **Happy:** Given ein `constructs`-Eintrag (Muster `dlopen`, Zone `adapters/plugin`), when eine Datei in `adapters/io` das Muster enthält, then ein Befund (`construct-leak`) mit Datei, Zeile, Muster und erlaubter Zone und Exit-Code 1; eine Datei **in** der Zone bleibt befundfrei.
- **Boundary (schichtlose Datei):** Given eine Datei, die einem `languages`-Glob, aber **keinem** `layers`-Glob entspricht, when sie das Muster außerhalb der Zone enthält, then ein Befund — die Roh-Text-Prüfung ist scan-weit, nicht layer-gebunden.
- **Boundary (Composition Root):** Given das Muster in der deklarierten Composition Root, when der Eintrag `composition_root: forbid` trägt, then ein Befund; mit dem Default `allow` **kein** Befund.
- **Boundary (Kommentar):** Given ein Treffer, der ausschließlich in einem Kommentar einer **C-Syntax-Sprache** (`//`, `/* */`) steht, when `a-check` läuft, then **kein** Befund (deklarierte Divergenz zur `grep`-Referenz); Given denselben Treffer in einem **Python-`#`-Kommentar**, then ein Befund — Python wird nicht C-gestrippt, die Grenze wird ausgewiesen ([AC-QA-02](#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)).
- **Negative:** Given ein Eintrag mit leerem/fehlendem `pattern` oder `adapter`, unbekanntem `match`-/`composition_root`-Wert oder einer als Regex nicht kompilierbaren `pattern`, when `a-check` lädt, then Exit-Code 2 (fail-closed, Muster der 0.14.0-Härtung von [AC-FA-RULE-003](#ac-fa-rule-003--tech-kapselung-regel-tech-leak)).
- **Determinismus:** Given zwei `constructs`-Muster, die **dieselbe Zeile** derselben Datei treffen und beide außerhalb ihrer Zone liegen, when `a-check` läuft, then zwei Befunde in stabiler, byte-identisch reproduzierbarer Reihenfolge ([AC-QA-01](#ac-qa-01--determinismus)).

**Out-of-Scope:** kein Parser — String-Literale bleiben Text-Heuristik ([AC-QA-02](#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)); keine RE2-fremden Features (Lookaround/Backreferences); keine Auto-Inferenz von Zonen; ein **Zonen-Verbot je Schicht** (`forbid_in`: Muster in bestimmten Schichten verboten, anderswo egal) — die Evidenz dafür deckt a-check bereits als Kante ab; eine fail-closed **Import-Allowlist** je Schicht (umgekehrte Beweislast auf Roh-Specifiern) — eigener, weiterhin gated Faden; eine eigene **Graph-Kante** für `constructs` ([AC-FA-CLI-002](#ac-fa-cli-002--architektur-graph-ausgabe) zeigt Nicht-Kanten-Semantik als Legende, nicht als Verbindung).

### AC-FA-EXTRACT-001 — Sprach-Backends für die Import-Extraktion

**Beschreibung:** Pro Sprache liefert ein Backend die Menge „welche
Symbole/Module importiert diese Datei" — text-heuristisch über konfigurierbare
Muster: C++ (`#include`), Go (`import`), Rust (`use`/`extern crate`), Kotlin
(`import`), Java (`import`, inkl. `import static`), Python (`import` und
`from … import`), C# (`using`-Direktiven, Schlüssel `csharp`), TypeScript
(ES-Module-Importe/Re-Exports, Schlüssel `typescript`). Beide
Python-Formen liefern den gepunkteten Modulpfad; ein Alias (`as x`) und die
hinter `from … import` stehenden Namen werden nicht als Symbol gewertet.
C#-`using`-Direktiven liefern den gepunkteten Namespace: `static`/`global`
werden übersprungen, bei der Alias-Form (`using X = Ziel;`) wird das Ziel
(rechte Seite) gewertet; das `;` direkt nach dem Namen ist Pflicht im Muster —
`using`-**Statements** (`using var x = …;`, `using (…)`) sind keine Direktiven
und werden nie gewertet. TypeScript liefert den **Modul-Specifier**: den
String in einfachen **oder** doppelten Anführungszeichen (beide gleichwertig;
Backticks nie — Template-Literal-Grenze) hinter `from` bzw. im Import; das
Semikolon ist optional (ASI) und **nicht** Teil des Musters. Gegriffen werden
`import … from '…'` (inkl. `import type` und Inline-`type`-Modifier), der
Seiteneffekt-Import `import '…'`, Re-Exports `export … from '…'` (inkl.
`export * from`, `export * as ns from`, `export type … from`), die
Interop-Form `import X = require('…')` sowie die Fortsetzungszeile
`} from '…'` eines mehrzeilig umbrochenen Imports/Re-Exports. Der Mittelteil
zwischen `import`/`export` und `from` ist auf Import-Clause-Zeichen
beschränkt (Bezeichner, `{ } * ,`, `type`/`as`, Whitespace — kein `=`, `(`,
`.`, keine Quotes), damit Ausdrucks-Zeilen (`export const q =
knex.from('users')`) nie matchen; die links von `from` stehenden
Namen/Aliasse werden nie als Symbol geliefert. **Zusätzlich** liefert ein
**deklarations-bewusstes** Backend die Menge der **Top-Level-Deklarationen** einer Datei —
Namen von `fun` (inkl. Extension-Funktionen `fun R.name`), `val`/`var`/`const val`, `class`
(inkl. `data`/`sealed`/`enum`/`annotation class`), `object`, `interface`, `sealed interface`
und `typealias`, text-heuristisch — zur Schicht-Auflösung von Symbolen in **Split-Packages**,
deren Datei ≠ Symbolname ist ([AC-FA-CONF-001](#ac-fa-conf-001--konfigurationsdatei-a-checkyml)).
In 0.18.0 ist **Kotlin** das einzige deklarations-bewusste Backend; alle übrigen liefern ein
**leeres** Deklarations-Set (No-op, kein Verhaltenswechsel). Das Backend wird über die
Config (Sprache + Datei-Globs) gewählt.

**Akzeptanzkriterien:**

- **Happy:** Given eine Go-Datei mit zwei Imports, when das Go-Backend läuft, then liefert es genau diese zwei Importpfade.
- **Boundary:** Given eine Rust-Alias-Form (`use tauri as t;`), when das Rust-Backend läuft, then wird `tauri` erkannt.
- **Negative:** Given eine in einem Kommentar/String stehende Import-ähnliche Zeile, when das Backend läuft, then wird sie nicht als Import gewertet (oder als bewusste, dokumentierte Heuristik-Grenze gemeldet — `AC-QA-02`).
- **Happy (Java):** Given `import com.foo.Bar;`, when das Java-Backend läuft, then liefert es das Symbol `com.foo.Bar` (das `;` wird ignoriert).
- **Boundary (Java static):** Given `import static com.foo.Bar.baz;`, when das Java-Backend läuft, then liefert es `com.foo.Bar.baz` — das `static`-Schlüsselwort wird übersprungen, nicht als Symbol gewertet.
- **Happy (Python import):** Given `import myapp.adapters.db`, when das Python-Backend läuft, then liefert es das Symbol `myapp.adapters.db`.
- **Boundary (Python from):** Given `from myapp.adapters import db`, when das Python-Backend läuft, then liefert es `myapp.adapters` — den Modulpfad nach `from`; die importierten Namen werden nicht expandiert.
- **Boundary (Python Alias):** Given `import myapp.adapters as ad`, when das Python-Backend läuft, then liefert es `myapp.adapters` (das `as ad` wird nicht gewertet).
- **Happy (C#):** Given `using MyApp.Adapters.Db;`, when das C#-Backend läuft, then liefert es das Symbol `MyApp.Adapters.Db`.
- **Boundary (C# static/global):** Given `using static System.Math;` bzw. `global using MyApp.Core;`, when das C#-Backend läuft, then liefert es `System.Math` bzw. `MyApp.Core` — die Schlüsselwörter werden übersprungen, nicht als Symbol gewertet.
- **Boundary (C# Alias):** Given `using Db = MyApp.Adapters.Db;`, when das C#-Backend läuft, then liefert es `MyApp.Adapters.Db` (das Ziel; der Alias-Name wird nie geliefert).
- **Negative (C# using-Statement):** Given `using var f = File.Open(p);` oder `using (var f = File.Open(p))`, when das C#-Backend läuft, then wird **kein** Symbol geliefert (Ressourcen-Statement, keine Direktive).
- **Happy (TS):** Given `import { Db } from '../adapters/db';` sowie `import { Repo } from "../adapters/db.js"` (Double-Quotes, semikolonfrei, `.js`-Specifier auf `.ts`-Datei — NodeNext-Schreibweise), when das TypeScript-Backend läuft, then liefert es `../adapters/db` bzw. `../adapters/db.js`.
- **Boundary (TS type/Seiteneffekt):** Given `import type { Repo } from './ports/repo';` und `import './polyfill';`, when das TypeScript-Backend läuft, then liefert es `./ports/repo` bzw. `./polyfill`.
- **Boundary (TS Re-Export/require/mehrzeilig):** Given `export * from './core/model';`, `import fs = require('fs');` und ein mehrzeilig umbrochener Import, dessen Schlusszeile `} from '../adapters/db';` lautet, when das TypeScript-Backend läuft, then liefert es `./core/model`, `fs` bzw. `../adapters/db`.
- **Negative (TS Ausdruck):** Given `const m = await import('./lazy');`, `const x = require('pg');`, `export const q = knex.from('users');` oder eine zeilen-anführende Ausdrucks-Zeile `import('./x').then(m => m.run());`, when das TypeScript-Backend läuft, then wird **kein** Symbol geliefert (dokumentierte Heuristik-Grenze, `AC-QA-02`).
- **Happy (Kotlin-Deklarationen):** Given eine Kotlin-Datei mit `fun Pool.asJdbc()`, `class Foo` und `typealias Bar = X`, when das Kotlin-Backend läuft, then liefert es die Top-Level-Deklarationen `asJdbc`, `Foo`, `Bar` — unabhängig vom Dateinamen.
- **Boundary (Deklaration in Kommentar):** Given eine `fun`-/`class`-ähnliche Zeile in einem Kommentar **oder** eine eingerückte (Member-)Deklaration, when das Kotlin-Backend läuft, then wird sie **nicht** als Top-Level-Deklaration gewertet (Kommentar-Strip + Spalte-0-Verankerung); eine Deklarations-ähnliche Zeile in einem **mehrzeiligen String-Literal** bleibt die ausgewiesene Heuristik-Grenze (`AC-QA-02`, wie bei der Import-Extraktion).
- **Negative (kein Deklarations-Backend):** Given eine Go-/C#-/TypeScript-/…-Datei, when ihr Backend läuft, then ist das Deklarations-Set **leer** (No-op) — die Schicht-Auflösung bleibt bei package==directory.

**Out-of-Scope:** vollständiges AST-Parsing; Toolchain-gestützte Backends (`go list`, `javac`/`jdeps`, Bytecode) sind ein opt-in-Re-Eval, nicht 0.1.0; Java-Wildcard-Imports (`import com.foo.*;`) werden heuristisch gegriffen (Symbol `com.foo.` mit Trailing-Dot), nicht expandiert; mehrere `import`-Statements auf **einer** Zeile werden nur einmal gegriffen (dokumentierte Heuristik-Grenze, `AC-QA-02`); relative Python-Importe (`from .`/`from ..`) werden nicht extrahiert — eine dokumentierte Grenze der Python-Extraktion (`AC-QA-02`), unabhängig vom gültigen `relative`-Auflösungs-Modus ([AC-FA-CONF-001](#ac-fa-conf-001--konfigurationsdatei-a-checkyml)), den Specifier-Sprachen wie TypeScript nutzen; Python-Mehrfach-Import in **einem** Statement (`import a, b`) wird nur als Erst-Treffer (`a`) gegriffen; die Subpaket-Form `from <paket> import <subpaket>` wird nur als `<paket>` gewertet und löst ggf. auf keine Schicht auf (dokumentierte Heuristik-Grenze, `AC-QA-02`); `__init__`-Re-Export-Semantik; import-ähnliche Zeilen in Docstrings (bestehende String-Grenze, `AC-QA-02`); C#-Typ-Aliasse auf generische Typen (`using L = List<int>;` — kein Namespace-Import, nicht gegriffen), `extern alias` und `global::`-qualifizierte Aliasse; C#-Namespace-**Deklarationen** (`namespace X;`/`namespace X { }`) werden nicht als Import gewertet; dynamisches TypeScript-`import()`/`require()` im Ausdruck (zeilenverankerte Heuristik, `AC-QA-02`); import-ähnliche Zeichenfolgen in Template-Literalen (Backticks) und JSX-Textzeilen von `.tsx`-Dateien (bestehende String-Grenze, `AC-QA-02`); Triple-Slash-Direktiven (`/// <reference path="…" />` — fallen dem Kommentar-Strip zu); JavaScript (`.js`/`.mjs`/`.cjs`) als eigener `languages`-Schlüssel; tsconfig-`paths`/`baseUrl`-Aliasse (nicht Teil des `relative`-Modus, [AC-FA-CONF-001](#ac-fa-conf-001--konfigurationsdatei-a-checkyml)); Node-Modul-Auflösung (Endungen/index-Dateien — keine Datei-Existenz-Probe; der Glob-Präfix-Match der Auflösung ist endungs-agnostisch, solange die `layers`-Globs verzeichnisbasiert sind, [AC-FA-CONF-001](#ac-fa-conf-001--konfigurationsdatei-a-checkyml)); TypeScript-Specifier, die `//` enthalten (URL-/Protokoll-Importe wie `https://…`) — sie fallen dem Kommentar-Strip zu (String-/Strip-Grenze, `AC-QA-02`; URL-ESM/Deno nicht unterstützt); kompakt geschriebene Formen ohne Whitespace nach `import`/`export` (`import{A}from'./b'`) werden nicht gegriffen (Formatter-Konvention; das Pflicht-Whitespace schützt den Keyword-Präfix-Ausschluss); als Fortsetzungszeile mehrzeiliger Imports wird nur die `} from '…'`-Form gegriffen — ein nacktes `from '…'` auf eigener Zeile nicht (von Formatern nicht erzeugt; `AC-QA-02`); die Extraktion von **Top-Level-Deklarationen** ist in 0.18.0 auf **Kotlin** begrenzt (übrige Backends: leeres Set); verschachtelte/Member-Deklarationen sowie Deklarationen in nicht gescanntem/generiertem Code werden nicht indiziert (`AC-QA-02`).

### AC-FA-CLI-001 — Aufruf, Scan-Wurzel und Exit-Codes

**Beschreibung:** `a-check [pfad]` prüft das Repo unter `pfad` (Default `/src`
im Container) gegen die `.a-check.yml`. Exit-Codes: `0` kein Befund, `1`
mindestens ein Befund, `2` Nutzungs-/Konfigurationsfehler. Befunde auf stdout,
Zusammenfassung auf stderr (analog `d-check`).

**Akzeptanzkriterien:**

- **Happy:** Given ein konformes Repo, when `a-check` läuft, then Exit-Code 0.
- **Boundary:** Given ein read-only gemountetes Repo, when `a-check` läuft, then vollständige Prüfung ohne Schreibzugriff.
- **Negative:** Given eine fehlende/ungültige `.a-check.yml`, when `a-check` läuft, then Exit-Code 2 (mit Zeilenangabe, wo die Fehlerquelle eine Zeile hat).

**Out-of-Scope:** Auto-Fix/Reparatur von Architekturverstößen (es gibt keinen deterministisch ableitbaren Fix).

### AC-FA-CLI-002 — Architektur-Graph-Ausgabe

**Beschreibung:** `a-check --print-graph [pfad]` gibt die in `pfad/.a-check.yml`
(Default `/src`) **deklarierte** Architektur als **Mermaid-Flowchart** auf stdout
aus: ein Knoten je Schicht, eine Kante je deklarierter `edges`-Kante, eine
abgesetzte Sonderkante je `allow`-Kante, Farbcodierung nach der **effektiven**
Schicht-Rolle (explizites `role:` oder Namens-Inferenz wie
[AC-FA-RULE-006](#ac-fa-rule-006--schicht-rollen-generische-regel-anwendung)). Der
Modus liest **nur** die Konfiguration und **scannt keine Quellen** — er ist
read-only (schreibt nie ins geprüfte Repo, `AC-QA-02`) und deterministisch
(byte-identische Ausgabe bei identischer Config, `AC-QA-01`). Der Graph zeigt die
deklarierte **Absicht**, nicht den realen Code: die kategorischen Rollen-/
Richtungs-Constraints (`core-impurity`, `lateral-adapter`,
`port-direction-mismatch`) sind keine Kanten und erscheinen als **Legende**, nicht
als gezeichnete Verbindung (ehrliche Ausgabe ohne Semantik-Behauptung, `AC-QA-02`).
Dies ist eine eigenständige **Inspektions**-Ausgabe, **keine** Erweiterung der
Distributions-`--print-*`-Familie aus
[AC-FA-DIST-001](#ac-fa-dist-001--distribution-image---print-mk-a-checkmk).
Exit-Codes wie beim Scan-Aufruf
([AC-FA-CLI-001](#ac-fa-cli-001--aufruf-scan-wurzel-und-exit-codes)): `0` bei
erfolgreicher Ausgabe; `2` bei Nutzungs-/Konfigurationsfehler — ein ladezeitiger
Config-Fehler (**inkl. eines unbekannten `languages`-Schlüssels**, identisch zum
Scan), ein unbekanntes Flag **oder** ein zusätzliches Argument/Flag nach dem
optionalen `pfad`. Scanzeitige, datei-mengenabhängige Auflösungs-Fehler treten im
no-scan-Modus nicht auf (`AC-QA-02`).

**Akzeptanzkriterien:**

- **Happy:** Given eine gültige `.a-check.yml` mit `layers`, `edges` und einer `allow`-Kante, when `a-check --print-graph <pfad>` läuft, then eine Mermaid-`flowchart`-Ausgabe mit einem Knoten je Schicht, einer Kante je `edge`, einer abgesetzten `allow`-Kante und rollen-basierter Farbe auf stdout, Exit-Code 0 — und es wird **nichts** ins Repo geschrieben (read-only).
- **Boundary:** Given eine minimale Config **ohne** optionale Blöcke (kein `tech`/`adapter_sink`/`composition_root`/`direction`) **oder** Schicht-Namen mit Mermaid-heiklen Zeichen (Punkte/Slashes/Quotes/`]`/`|`/Backtick/`<`/`>`/`&`), when `a-check --print-graph` läuft, then ein **gültiges** Flowchart ohne Syntax-Ausbruch und ohne leere Sonderknoten; zwei Läufe gegen dieselbe Config sind byte-identisch (`AC-QA-01`).
- **Negative:** Given eine ungültige `.a-check.yml` (unbekannter Schlüssel **oder** unbekannte Sprache `languages: {ruby: …}`), ein unbekanntes Flag **oder** ein zusätzliches Argument nach dem Pfad (`--print-graph <pfad> --bogus`), when es läuft, then Exit-Code 2 (kein halbes Diagramm).

**Out-of-Scope:** Nicht-Mermaid-Formate (DOT/Graphviz) und ein Findings-/Verstoß-Graph (der reale Import-Graph mit markierten Befunden) — eigene Folge-Slices; die visuelle Darstellung von `tech` (Muster/Badge) — Folge-Inkrement; scanzeitige, datei-mengenabhängige Auflösungs-Mehrdeutigkeiten (Mehr-Wurzel) — sie brauchen einen realen Scan und sind kein Vertrag des no-scan-Modus; Auto-Layout-Tuning (Mermaid ordnet selbst; sehr große Configs können unübersichtlich werden — dokumentierte Grenze).

### AC-FA-CLI-003 — Usage-Ausgabe und Handbuch-Verweis

**Beschreibung:** `a-check --help` gibt eine **Usage-Ausgabe** aus, die vier
Bestandteile trägt: eine Kurzbeschreibung dessen, was das Werkzeug tut; die
**Aufruf-Syntax** samt optionalem Pfad-Parameter; einen Hinweis auf die
Konfigurationsdatei `.a-check.yml`; und die **URL des Benutzerhandbuchs**.
Dieselbe URL trägt der **Kopfkommentar** des per `--print-mk` erzeugten
Fragments ([AC-FA-DIST-001](#ac-fa-dist-001--distribution-image---print-mk-a-checkmk)):
das Fragment reist in ein **fremdes** Repo, und sein Kopf ist der einzige Ort,
an dem ein Zeiger auf die Dokumentation dauerhaft mitfährt.

**Zugesichert ist die Anwesenheit der Bestandteile, nicht ihr Wortlaut.** Ein
Test darf prüfen, *dass* Aufruf-Zeile, Konfigurations-Hinweis und URL erscheinen;
er darf sich nicht an ihre Formulierung binden.

**Die URL zeigt auf den Hauptzweig, ohne Versionsangabe.** Das Binary kennt
seinen eigenen Release-Kontext zur Laufzeit nicht — es trägt keine eingebackene
Version —, und ein zur Build-Zeit eingesetzter versionierter Link nennte immer
den **Vorgänger**-Stand: gültig aussehend und falsch. Die tag-freie URL kann
dafür nicht veralten, zeigt aber auch nie auf den Stand, den ein gepinntes Image
fährt; wer die passende Fassung braucht, liest den Software-Versions-Stempel im
Kopf des Handbuchs.

**Akzeptanzkriterien:**

- **Happy:** Given das Image, when `a-check --help` läuft, then eine Usage-Ausgabe mit Kurzbeschreibung, Aufruf-Syntax, Konfigurations-Hinweis und Handbuch-URL, Exit-Code 0 — und es wird **nichts** ins Repo geschrieben (read-only).
- **Boundary:** Given `a-check --print-mk`, when es läuft, then trägt der Kopfkommentar des Fragments dieselbe Handbuch-URL wie die Usage-Ausgabe; zwei Läufe sind byte-identisch (`AC-QA-01`).
- **Negative:** Given ein unbekanntes Flag (`a-check --bogus`), when es läuft, then Exit-Code 2 und die Usage-Ausgabe auf stderr — der Nutzungsfehler bleibt ein Fehler, er wird nicht durch den Hilfetext zu Exit-Code 0.

**Out-of-Scope:** Eine versionierte oder digest-gebundene Handbuch-URL (das Binary kennt seinen Release-Kontext nicht); lokalisierte Usage-Ausgaben (die Ausgabe ist deutsch wie das Handbuch); ein `--version`-Flag und eine Ausgabe der Konfigurations-Fundstelle — eigene Folge-Inkremente; die Erreichbarkeit der URL (das Werkzeug ist hermetisch, `AC-QA-02`, und prüft sie nie).

### AC-FA-CONF-001 — Konfigurationsdatei `.a-check.yml`

**Beschreibung:** `.a-check.yml` deklariert: die Sprache(n) + Datei-Globs je
Schicht, die Schichten (`core`/`ports`/`adapters`/…) mit Pfad-Mustern, die
erlaubten Kanten, die Tech→Adapter-Zuordnungen und die gemeinsame Adapter-Senke. Ein `layers`-Eintrag
ist **entweder** eine Glob-Liste (`name: [globs]`, Rolle per Namens-Inferenz)
**oder** ein Objekt `{globs: [...], role: domain|app|port|adapter, direction: driving|driven}`
([AC-FA-RULE-006](#ac-fa-rule-006--schicht-rollen-generische-regel-anwendung), [AC-FA-RULE-008](#ac-fa-rule-008--driving-driven-port-richtung-regel-port-direction-mismatch)); `direction` ist optional.
Ein `tech`-Eintrag ist `{pattern, adapter}` — `adapter` als Pfad **oder**
Pfad-**Liste** (Symbol in jedem gelisteten Adapter erlaubt; leere Liste
unzulässig) — mit optionalem `match: substring|regex` (Default `substring`;
`regex` = RE2) und optionalem `composition_root: allow|forbid` (Default `allow`,
[AC-FA-RULE-003](#ac-fa-rule-003--tech-kapselung-regel-tech-leak)).
Ein optionaler **`exclude`**-Block (Datei-Globs relativ zur Scan-Wurzel) nimmt
matchende Dateien **vor** der Extraktion vollständig vom Scan aus (z. B.
Test-Dateien wie `**/*_test.go` oder generierter Code — die Glob-Engine der
`layers` kennt bewusst keine Negation, `exclude` ist das explizite Gegenstück);
fehlt er, wird jede `languages`-Glob-Datei gescannt (bisheriges Verhalten).
Ein optionaler **`constructs`**-Block deklariert Roh-Text-Muster mit ihrer erlaubten Zone
(`{pattern, adapter}` — `adapter` als Pfad oder Pfad-Liste — mit optionalem
`match: substring|regex` und `composition_root: allow|forbid`, dieselbe Scoping-Mechanik wie
`tech`, [AC-FA-RULE-011](#ac-fa-rule-011--konstrukt-monopol-regel-construct-leak)); fehlt er,
entfällt die Regel `construct-leak`.
Ein optionaler `resolution`-Block deklariert **je Sprache**, wie Import-Symbole auf Schichten
aufgelöst werden — Map Sprache → `{mode, roots, package_base}`, `mode ∈ {path (Default),
fixed-root, relative}` (`namespace` reserviert); fehlt er (oder eine Sprache darin), gilt
Import-als-Pfad. `relative` löst Specifier, die `.` oder `..` sind oder mit `./` bzw. `../`
beginnen, lexikalisch gegen das **Verzeichnis der importierenden Datei** auf; alle anderen
Specifier (Bare-Imports) und Specifier mit führendem `..` **nach** der Normalisierung
(Wurzel-Escape) liefern eine **leere** Kandidatenmenge — das Roh-Symbol wird nicht als
Pfad-Kandidat weitergereicht (kein Geister-Match, ausgewiesene Grenze `AC-QA-02`);
`mode: relative` nimmt weder `roots` noch `package_base`. Die Endungs-Agnostik der Auflösung
gilt, solange der Layer-Glob-Präfix oberhalb der Dateiebene endet (verzeichnisbasierte Globs);
bei datei-tiefen Globs kippt eine Specifier-Endung den Match (dokumentierte Grenze).
Striktes Decoding, fail-closed (Exit 2 bei unbekanntem Schlüssel, ungültiger `role`/`direction`,
unbekanntem `match`-Wert, einer als Regex nicht kompilierbaren `pattern`, einem `languages`-Schlüssel
außerhalb der unterstützten Backends aus [AC-FA-EXTRACT-001](#ac-fa-extract-001--sprach-backends-für-die-import-extraktion),
einem reservierten/unbekannten `resolution.mode` oder `roots`/`package_base` bei `mode: relative`,
einer **leeren** `tech.adapter`-Liste oder einem leeren/fehlenden `tech.adapter`, einem `composition_root`-Wert außerhalb
`{allow, forbid}`, einem ungültigen `exclude`-Glob oder einem `constructs`-Eintrag mit
leerem/fehlendem `pattern`/`adapter`, unbekanntem `match`/`composition_root` bzw. nicht
kompilierbarer Regex).
Bei `mode: fixed-root` mit **≥ 2** `roots` (geteiltes `package_base`, Paket-Namespaces über
mehrere Module — Gradle-Multi-Modul, auch **Split-Packages**, bei denen dasselbe Paket real
über mehrere Modul-Roots verteilt ist) wird der interne FQN **datei-mengen-bewusst** aufgelöst.
Für ein **deklarations-bewusstes** Backend (0.18.0 nur **Kotlin**, [AC-FA-EXTRACT-001](#ac-fa-extract-001--sprach-backends-für-die-import-extraktion))
belegt die stärkste Evidenz die Schicht: ein Root, in dessen Paket-Verzeichnis eine gescannte
Datei das importierte Symbol als **Top-Level-Deklaration** trägt (unabhängig vom Dateinamen),
sticht einen Root, in dem nur das Paket-Verzeichnis existiert — eine bloß gleichnamige, das
Symbol **nicht** deklarierende Datei zählt nicht als Deklaration. Für die übrigen (nicht
deklarations-bewussten) Backends gilt unverändert der Datei-Namens-Match (endungs-agnostisch,
package==directory-Grenze `AC-QA-02`). Die Schicht wird am Pfad des **auflösenden** Roots
bestimmt, nicht am Wurzel-Präfix. **Fail-closed (Exit 2):** löst derselbe FQN real unter ≥ 2
Roots in **verschiedene** Schichten auf (echte Mehrdeutigkeit — bei deklarations-bewussten
Backends: real **deklariert** in ≥ 2 Roots), bricht `a-check` **nach dem Scan** ab; gleiche
Schicht (z. B. `expect`/`actual`) löst sauber. Die Stufe *nur-Paketverzeichnis* (kein Deklarations-Treffer bzw. ein **Wildcard-/Paket-Import**
`a.b.*`, nur das Paket-Verzeichnis existiert) **löst** rückwärtskompatibel, solange sie **eindeutig** ist (genau ein Root oder alle dieselbe
Schicht); liegen Paket-Verzeichnisse ohne Deklaration unter ≥ 2 Roots in **verschiedenen** Schichten,
bleibt der FQN **extern** (fail-open, kein Geister-Match — anders als *deklariert in ≥ 2
verschiedenen Schichten*, das Exit 2 auslöst). Ganz ohne reale Evidenz bleibt der FQN extern. Dokumentierte Grenzen (`AC-QA-02`): ein
intern gemeintes Symbol, das in keiner gescannten Datei als Top-Level-Deklaration auftaucht
(verschachtelte Klasse, `object`-Member, nicht gescannter/generierter Code, Star-Import),
bleibt still extern; datei-tiefe Globs sind eine heuristische Grenze.

**Akzeptanzkriterien:**

- **Happy:** Given eine gültige `.a-check.yml`, when `a-check` läuft, then werden die deklarierten Regeln angewandt.
- **Boundary:** Given eine Config ohne optionale Tech-Zuordnungen, when `a-check` läuft, then laufen nur die Schicht-/Lateral-Regeln (kein `tech-leak`).
- **Negative:** Given ein Tippfehler im Schlüssel, when `a-check` läuft, then Exit-Code 2 (kein stiller Default).
- **Negative (`match`):** Given ein `tech.match` mit einem anderen Wert als `substring`/`regex` **oder** ein `match: regex` mit leerer bzw. nicht kompilierbarer `pattern`, when `a-check` lädt, then Exit-Code 2.
- **Negative (Sprache):** Given ein `languages`-Schlüssel außerhalb der von [AC-FA-EXTRACT-001](#ac-fa-extract-001--sprach-backends-für-die-import-extraktion) definierten unterstützten Backends, when `a-check` lädt, then Exit-Code 2 — **statt** stiller Nicht-Extraktion (falsch-grün).
- **Happy (Auflösung):** Given ein `resolution` mit `{mode: fixed-root, roots, package_base}` für eine Sprache und ein gepunktetes Symbol dieser Sprache, when `a-check` läuft, then löst das Symbol wurzel-relativ auf seine Schicht auf (statt unaufgelöst zu bleiben) — sofern der Paket-Baum den Verzeichnis-Baum spiegelt (`AC-QA-02`-Grenze).
- **Happy (`relative`):** Given `resolution: {typescript: {mode: relative}}` und in `src/core/service.ts` der Import `../adapters/db`, when `a-check` läuft, then löst das Symbol auf `src/adapters/db` auf (Schicht der `adapters`-Globs) — eine Domänen-Datei mit diesem Import wird `core-impurity`. Grenzfall inklusive: `../../x` aus `src/core/` normalisiert auf `x` (exakt Wurzelebene, aufgelöst).
- **Boundary (`relative` Escape/Bare-Import):** Given `../../../x` aus `src/core/` (führendes `..` nach der Normalisierung) oder `import * as core from '@actions/core'` bei einem Layer-Glob `core/**`, when `a-check` läuft, then bleibt das Symbol unaufgelöst — leere Kandidatenmenge, kein Ziel-Layer, kein Geister-Befund; `tech`-Muster greifen unabhängig am Roh-Symbol (ausgewiesene Grenze, `AC-QA-02`).
- **Negative (`resolution`):** Given ein `resolution.mode` außerhalb `{path, fixed-root, relative}` (inkl. des reservierten `namespace`) **oder** `{mode: relative, roots: […]}` bzw. `{mode: relative, package_base: …}`, when `a-check` lädt, then Exit-Code 2.
- **Happy (`exclude`):** Given `exclude: ["**/*_test.go"]` und ein Tech-/Schicht-Verstoß **nur** in einer Test-Datei, when `a-check` läuft, then kein Befund (die Datei wird nicht gescannt).
- **Boundary (`exclude`):** Given eine Config **ohne** `exclude`, when `a-check` läuft, then byte-identische Ausgabe wie bisher.
- **Negative (neue Schlüssel):** Given eine **leere** `tech.adapter`-Liste, ein leerer/fehlender `tech.adapter`, ein `composition_root` mit einem Wert außerhalb `{allow, forbid}` **oder** ein ungültiger `exclude`-Glob, when `a-check` lädt, then Exit-Code 2.
- **Negative (`constructs`):** Given einen `constructs`-Eintrag mit leerem/fehlendem `pattern` oder `adapter`, unbekanntem `match`-/`composition_root`-Wert **oder** einer als Regex nicht kompilierbaren `pattern`, when `a-check` lädt, then Exit-Code 2 ([AC-FA-RULE-011](#ac-fa-rule-011--konstrukt-monopol-regel-construct-leak)).
- **Happy (Multi-Modul disjunkt):** Given `mode: fixed-root` mit ≥ 2 `roots` + geteiltem `package_base` und disjunkten Paket-Sub-Namespaces je Modul (KMP: `mod-a/…/domain`, `mod-b/…/application` mit flachen Modul-Globs), when eine `domain`-Datei `com.ex.application.B` importiert, then löst der FQN datei-mengen-bewusst auf das reale Modul (Schicht `application`) auf und die verbotene Kante wird gemeldet (Exit 1) — **statt** stiller Fehlklassifikation (vor der datei-mengen-bewussten Auflösung: 0 Befunde, `AC-QA-02`).
- **Happy (Split-Package / Top-Level-Symbol):** Given `mode: fixed-root` mit ≥ 2 `roots`, ein **Split-Package** über zwei Schicht-Roots (`ports`, `adapters`) und ein Kotlin-Top-Level-Symbol, dessen Datei **≠** Symbolname ist (Extension-Fun `asJdbc` bzw. Zweitklasse), **genau in einem** Root deklariert, when eine Datei es importiert, then löst der FQN über die reale Top-Level-Deklaration auf die Schicht dieses Roots auf — **kein Exit 2** (vor der deklarations-bewussten Auflösung: Exit 2). Trägt Root A eine gleichnamige Datei, die das Symbol **nicht** deklariert, und Root B die echte Deklaration, then löst er auf **Root B**.
- **Boundary (Mehr-Wurzel, gleiche Schicht):** Given denselben FQN real unter ≥ 2 Roots, die **dieselbe** Schicht treffen (`expect`/`actual`), when `a-check` läuft, then löst er sauber auf — kein Exit 2.
- **Boundary (Split-Package ohne Deklaration):** Given ein Symbol, dessen Paket-Verzeichnis unter ≥ 2 Roots existiert, das aber in **keiner** gescannten Datei als Top-Level-Deklaration auftaucht, when `a-check` läuft, then bleibt es **extern** — kein Exit 2, kein Befund (fail-open, `AC-QA-02`).
- **Negative (Mehr-Wurzel, echte Mehrdeutigkeit):** Given denselben FQN real **auflösend** unter ≥ 2 Roots in **verschiedenen** Schichten (bei deklarations-bewussten Backends: real **deklariert** in ≥ 2 Roots), when `a-check` läuft, then Exit-Code 2 **nach dem Scan** (fail-closed, ein FQN muss in höchstens eine Schicht auflösen) — stderr nennt die Mehrdeutigkeit, kein Befund auf stdout.

**Out-of-Scope:** Vererbung/Includes zwischen Config-Dateien; per-Root `package_base`; der
reservierte Namespace-Index-Modus; die deklarations-bewusste Auflösung ist auf **Kotlin**
begrenzt (übrige Backends: package==directory, kein Deklarations-Index); die Auflösung
verschachtelter-Klassen-Importe, `object`-Member und datei-tiefer Globs bleiben heuristische
Grenzen (`AC-QA-02`).

### AC-FA-DIST-001 — Distribution: Image, `--print-mk`, `a-check.mk`

**Beschreibung:** `a-check` wird als GHCR-Image (distroless/static,
digest-gepinnt) verteilt. `a-check --print-config` gibt ein kommentiertes
`.a-check.yml`-Gerüst aus; `a-check --print-mk` gibt ein `a-check.mk` mit dem
**aktuell digest-gepinnten** Image, einem `a-check`-Scan-Target **und** einem
`a-check-graph`-Target aus. Konsumenten `include a-check.mk` und liefern
`.a-check.yml` — keine Skript-Kopie. Das `a-check-graph`-Target ruft
`--print-graph` auf ([AC-FA-CLI-002](#ac-fa-cli-002--architektur-graph-ausgabe):
Architektur-Graph als Mermaid nach stdout, read-only, kein Scan) und teilt dieselbe
digest-gepinnte `A_CHECK_IMAGE`-Variable — kein zweiter Digest, keine Skript-Kopie.

**Akzeptanzkriterien:**

- **Happy:** Given das Image, when `a-check --print-mk` läuft, then ein `include`-bares Makefile-Fragment mit digest-gepinntem `A_CHECK_IMAGE` und einem `a-check`-Scan-Target auf stdout.
- **Happy (Graph-Target):** Given das erzeugte `a-check.mk` in einem Konsumenten-`Makefile` eingebunden, when `make a-check-graph` läuft, then ein Mermaid-`flowchart` auf stdout (read-only, kein Scan), Exit-Code 0 — dasselbe `A_CHECK_IMAGE` wie das `a-check`-Scan-Target, kein Schreibzugriff auf den Baum.
- **Boundary:** Given `a-check --print-config`, when es läuft, then ein dekodierbares `.a-check.yml`-Gerüst, **schreibt nichts** (read-only).
- **Negative:** Given `--print-mk` mit einem zusätzlichen unbekannten Flag, when aufgerufen, then Exit-Code 2.

**Out-of-Scope:** Nicht-Docker-Distribution (Binary-Releases) in 0.1.0; ein Datei-schreibendes oder Browser-öffnendes Graph-Target (`a-check-graph` schreibt nur nach stdout, Umleiten ist Nutzer-Sache); weitere Ausgabeformate (`--graph-format` DOT/Graphviz) — eigener Folge-Slice.

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

Das **Image** wird per `@sha256:`-Digest referenziert; die Pin-Hebung ist ein
bewusster Commit (analog der Pin-Politik der Konsumenten-Repos). Die im Repo
**committete** `a-check.mk` trägt den Digest des aktuellen Release
([`version.md#aktuell`](../version.md#aktuell)).

Das von `--print-mk` **erzeugte** Fragment trägt dagegen **keinen konkreten
Digest**, sondern einen **Platzhalter** mit Bezugsquelle. Grund: ein Binary kann
den Digest des Image, in dem es läuft, strukturell nicht kennen — der Digest
entsteht erst beim Push, das Binary ist vorher gebaut. Ein
eingebackener Digest nennt darum **immer den Vorgänger** und sieht dabei
autoritativ aus. Der Platzhalter erzwingt stattdessen genau den bewussten
Commit, den dieser Abschnitt ohnehin verlangt.

- **Happy:** Given ein Release-Image, when `--print-mk` läuft, then enthält die Ausgabe keinen `@sha256:`-Wert eines anderen Release, sondern einen als solchen erkennbaren Platzhalter samt Bezugsquelle.
- **Boundary:** Given die committete `a-check.mk`, when sie gelesen wird, then trägt sie den vollen Digest des aktuellen Release — sie ist das gepinnte Artefakt, das erzeugte Fragment die Vorlage dafür.
- **Negative:** Given ein Konsument übernimmt das Fragment unverändert, when `make a-check` läuft, then bricht der Aufruf sichtbar am Platzhalter ab, statt still ein fremdes Release zu ziehen.
- **Out-of-Scope:** einen korrekten Eigen-Digest zur Laufzeit ermitteln (netzlos nicht möglich; der Host kennt ihn über `docker inspect`, das Binary nicht); die Pin-Hebung beim Konsumenten automatisieren.

## 7. Historie

Regeln dieser Sektion: Ab Status `Accepted` ist **jede** Änderung an diesem Dokument eine
Vertragsänderung — auch das **Hinzufügen** einer Anforderung. Sie entsteht nur aus einem Change
Request, nie aus einer ADR oder einem Slice. Fußabdruck je angenommenem CR: Versions-Bump oben plus
eine Zeile hier (Baseline-Regelwerk `grundlagen-source-precedence.md` §Spec-Stratifizierung).

| Version | Datum | Änderung |
|---|---|---|
| 0.1.0 | 2026-06-20 | Erstfassung (Bootstrap): Zweck/Inventur, fünf universelle Hexagon-Regeln (`AC-FA-RULE-001…005`), Sprach-Extraktion, CLI, Config, Distribution (`--print-mk`/`a-check.mk`); NFAs Determinismus/Hermetik/Reproduzierbarkeit. |
| 0.2.0 | 2026-06-22 | `AC-FA-RULE-004` neu gefasst: Ports **dürfen** Domänen-/Kern-Typen referenzieren (Sprache des Kerns; `ports → core` per deklarierter Kante), `port-impurity` trennt scharf gegen Adapter-/Tech-Importe. Motiviert durch die Vier-Repo-Evidenz (b-cad/d-migrate-Ports referenzieren die Domäne). |
| 0.3.0 | 2026-06-22 | Neu `AC-FA-RULE-006` (Schicht-Rollen): die Reinheits-Regeln dispatchen über eine Layer-Rolle (`domain`/`port`/`adapter`, aus `role:` oder Namens-Inferenz) — generalisiert `AC-FA-RULE-001`/`AC-FA-RULE-002`/`AC-FA-RULE-004` namens-unabhängig (welle-10a). `AC-FA-CONF-001`-Schema: `layers`-Eintrag als Glob-Liste **oder** `{globs, role}`. |
| 0.4.0 | 2026-06-22 | `AC-FA-RULE-006`: `lateral-adapter` jetzt **vollständig** namensunabhängig — Adapter-Sub-Einheiten werden relativ zum Schicht-Glob-Präfix unterschieden (statt am Literal `adapters`); `adapterSeg`-Generalisierung aus dem Out-of-Scope eingelöst (welle-10b). |
| 0.5.0 | 2026-06-22 | Neu `AC-FA-RULE-007` (Rolle `app` + strenge `domain`): `app` darf `domain`+`port`, aber keinen Adapter/Tech (neuer Befund `app-impurity`); `domain` verschärft — Import auf `app`/`port`/`adapter`/Tech ist `core-impurity`, kategorisch („Domäne kennt keine Ports"). Erweitert `AC-FA-RULE-006`, schärft `AC-FA-RULE-001` (welle-10b). |
| 0.6.0 | 2026-06-23 | Neu `AC-FA-RULE-008` (Driving/Driven-Port-Richtung): optionale `direction` ∈ {`driving`, `driven`} auf `port`-/`adapter`-Schichten, **orthogonal** zur Rolle; neuer Befund `port-direction-mismatch` (ein Adapter spricht nur Ports seiner Richtung). Ohne `direction` keine Prüfung (rückwärtskompatibel). `AC-FA-CONF-001`-Schema: Objekt-Form um `direction` (und das in 0.5.0 fehlende `app`) ergänzt. Verfeinert `AC-FA-RULE-006` (welle-10b/b2b). |
| 0.7.0 | 2026-06-23 | `AC-FA-EXTRACT-001` um **Java** erweitert (`import`, inkl. `import static` — das `static`-Schlüsselwort übersprungen, `;` ignoriert) — fünftes Sprach-Backend neben C++/Go/Rust/Kotlin, text-heuristisch (welle-06, slice-014). |
| 0.8.0 | 2026-07-01 | `AC-FA-RULE-003`/`AC-FA-CONF-001`: `tech`-Muster optional als **RE2-Regex** (`match: substring\|regex`, Default `substring`) statt nur Substring — macht ein nur als Muster fassbares Tech wie Qt (`Q[A-Za-z]`) ausdrückbar; Mehrfach-Treffer lösen in Deklarationsreihenfolge (Erst-Treffer, kein „längster Präfix" für `tech`). Unbekanntes `match`/nicht kompilierbare Regex → Exit 2. Rückwärtskompatibel (ohne `match` byte-identisch). welle-05/-06, b-cad-Pilot (Regel E); slice-016. |
| 0.9.0 | 2026-07-01 | `AC-FA-CONF-001`: ein `languages`-Schlüssel außerhalb der unterstützten Backends (`cpp`/`go`/`rust`/`kotlin`/`java`, `AC-FA-EXTRACT-001`) bricht mit **Exit 2** ab — schließt die stille Nicht-Extraktion (falsch-grün) für nicht unterstützte Sprachen. slice-017. |
| 0.10.0 | 2026-07-01 | `AC-FA-CONF-001`: optionaler `resolution`-Block — Map **Sprache → `{mode, roots, package_base}`** (`mode ∈ {path, fixed-root}`, `relative`/`namespace` reserviert → Exit 2), löst gepunktete/wurzel-fremde Importe **pro Sprache** auf ihre Schicht auf (Mono-Repo-tauglich); Default (ohne Block) = Import-als-Pfad, rückwärtskompatibel. Grenze: Paket==Verzeichnis (`AC-QA-02`). welle-06 (Polyglot-Bestand); slice-015. |
| 0.11.0 | 2026-07-02 | `AC-FA-EXTRACT-001` um **Python** erweitert (`import` und `from … import` → gepunkteter Modulpfad; Alias und importierte Namen nicht gewertet; relative Importe nicht extrahiert — Signal des reservierten `relative`-Modus, dokumentierte Grenze `AC-QA-02`) — sechstes Sprach-Backend, text-heuristisch (welle-06, slice-020). |
| 0.12.0 | 2026-07-02 | `AC-FA-EXTRACT-001` um **C#** erweitert (Schlüssel `csharp`; `using`-Direktiven → gepunkteter Namespace, `static`/`global` übersprungen, Alias-**Ziel** gewertet, Pflicht-`;` schließt `using`-Statements aus) — siebtes Sprach-Backend. Schicht-Auflösung über den `fixed-root`-Modus unter Namespace==Verzeichnis (`AC-QA-02`-Grenze); der reservierte `namespace`-Modus bleibt Exit 2 (welle-06, slice-021). |
| 0.13.0 | 2026-07-03 | `AC-FA-EXTRACT-001` um **TypeScript** erweitert (Schlüssel `typescript`; ES-Module-Formen → Modul-Specifier in `'…'`/`"…"`, Semikolon optional: `import … from` inkl. `type`, Seiteneffekt-Import, Re-Exports `export … from`, `import X = require(…)`, Fortsetzungszeile `} from '…'` mehrzeiliger Imports; Mittelteil auf Import-Clause-Zeichen beschränkt — Ausdrucks-`import()`/`require()` nie gegriffen) — achtes Sprach-Backend (welle-06, slice-022). |
| 0.13.0 | 2026-07-03 | `AC-FA-CONF-001`: `resolution.mode` um **`relative`** erweitert — Specifier `.`/`..`/`./…`/`../…` lexikalisch gegen das Verzeichnis der importierenden Datei; Bare-Imports und Wurzel-Escape (führendes `..` nach Normalisierung) → **leere** Kandidatenmenge (kein Geister-Match, `AC-QA-02`); `roots`/`package_base` bei `relative` unzulässig → Exit 2; nur `namespace` bleibt reserviert (welle-06, slice-022). |
| 0.14.0 | 2026-07-03 | **CR d-check-Pilot (1/3)** — `AC-FA-RULE-003`/`AC-FA-CONF-001`: `tech.adapter` auch als Pfad-**Liste** (Symbol in jedem gelisteten Adapter erlaubt; leere Liste/leerer Eintrag → Exit 2); nicht-leerer Skalar bleibt gültig (rückwärtskompatibel), leerer/fehlender `adapter` → Exit 2 (fail-closed statt vormals stillem Never-Leak-Eintrag, Umsetzungs-Review). Anlass: d-check erlaubt `yaml` in **zwei** Adaptern (Config **und** Report) — mit Ein-Pattern-ein-Adapter nicht ausdrückbar (Erst-Treffer-Präzedenz). welle-11, slice-023. |
| 0.14.0 | 2026-07-03 | **CR d-check-Pilot (2/3)** — `AC-FA-RULE-003`/`AC-FA-CONF-001`: die Composition-Root-Ausnahme der `tech`-Kapselung wird **pro Eintrag steuerbar** (`composition_root: allow\|forbid`, Default `allow` = bisheriges Verhalten); `forbid` betrifft nur `tech-leak`, die Schicht-Regel-Ausnahme der Composition Root bleibt. Anlass: d-check verbietet `net/http`/`yaml` auch in CLI/`cmd` — die Total-Ausnahme kostete dort die Deckung. welle-11, slice-023. |
| 0.14.0 | 2026-07-03 | **CR d-check-Pilot (3/3)** — `AC-FA-CONF-001`: optionaler **`exclude`**-Block (Datei-Globs) nimmt Dateien vor der Extraktion vom Scan aus (explizites Gegenstück zur negations-freien Glob-Engine); ohne Block byte-identisch. Anlass: der Scanner erfasst `*_test.go`, d-checks abgelöstes `go list`-Gate prüfte nur Nicht-Test-Imports — ohne Ausschluss ist der saubere Baum dort rot. welle-11, slice-023. |
| 0.15.0 | 2026-07-03 | `AC-FA-RULE-002`: Sub-Einheiten-Grenzfall präzisiert — Blatt-Klassifikation: ein **datei-förmiges** Blatt (`.`) direkt im Layer-Root gehört zur **Root-Sub-Einheit `''`** (Sub-Einheiten sind Verzeichnisse, keine Dateinamen), ein **verzeichnis-förmiges** Blatt (Go-Paket-Pfad) **ist** die Sub-Einheit; Root↔Root same-layer ist kein `lateral-adapter` mehr, Root↔Unterverzeichnis/Cross-Layer/Cross-Paket unverändert. Bewusste Gate-Lockerung per ADR; Anlass: b-cad-Pilot — die Richtungs-Modellierung (pro-Adapter-Layer) erzeugte 40 Falsch-Positive der Klasse `x.cpp → x.h` bei 0 echten Verstößen (welle-05/M3-Pilot, slice-024). |
| 0.16.0 | 2026-07-04 | `AC-FA-CONF-001`: **fail-closed-Guard gegen mehrdeutige Mehr-Wurzel-Auflösung** — `mode: fixed-root` mit ≥ 2 `roots`, von denen zwei Roots je eine andere Schicht erzwingen (die Schicht, in die ein Root allein am Wurzel-Präfix auflöst — längster passender Glob-Präfix, wie die Import-Auflösung), bricht mit Exit 2 statt still falsch-grün (Phantom-Kandidaten über Schicht-Grenzen; die Zuordnung entschiede der längste Präfix statt das Symbol). Anlass: belief-agent-Bericht — KMP (`commonMain`/`jvmMain` teilen `package_base`) bei flachen Source-Set-Globs fing die illegale `core → adapter`-Kante nicht; Rezept: paket-spezifische Globs. Stufe 1 des Fixes (fail-closed); die datei-mengen-bewusste Auflösung folgt gated (slice-027). welle-05-Härtung, slice-026. |
| 0.17.0 | 2026-07-05 | `AC-FA-CONF-001`: **datei-mengen-bewusste Mehr-Wurzel-Auflösung** (Stufe 2) — `mode: fixed-root` mit ≥ 2 `roots` löst den internen FQN gegen die **real gescannten Dateien** auf (endungs-agnostisch, package==directory), statt je Root einen Phantom-Kandidaten am Wurzel-Präfix zu bilden; die disjunkte Multi-Modul-Config (KMP: geteiltes `package_base`, disjunkte Sub-Namespaces) **lädt und löst korrekt** (die verbotene `domain → application`-Kante wird gemeldet), der Ladezeit-Guard aus 0.16.0 entfällt. Echte Mehrdeutigkeit (gleicher FQN real in ≥ 2 Roots, **verschiedene** Schichten) bricht **nach dem Scan** mit Exit 2 (distinct-layer; `expect`/`actual` same-layer löst sauber). Ersetzt Stufe 1 (Supersede-ADR). Anlass: belief-agent-KMP — jede resolution-Variante war Reject oder still falsch-grün. welle-05-Härtung, slice-027. |
| 0.18.0 | 2026-07-06 | `AC-FA-CONF-001`/`AC-FA-EXTRACT-001`: **deklarations-bewusste Auflösung für Split-Packages** (Stufe 3) — bei `mode: fixed-root` mit ≥ 2 `roots` löst ein importiertes Top-Level-Symbol, dessen Datei ≠ Symbolname ist (Kotlin-Extension-Fun, Zweitklasse), über die **reale Top-Level-Deklaration** statt nur über das Paket-Verzeichnis auf: genau ein deklarierender Root ⇒ eindeutig, ≥ 2 deklarierende Roots verschiedener Schichten ⇒ Exit 2, kein Deklarations-Treffer ⇒ **extern** (fail-open). Für deklarations-bewusste Backends sticht die echte Deklaration den bloßen Datei-Namens-Match. `AC-FA-EXTRACT-001`: **Kotlin** liefert zusätzlich Top-Level-Deklarationen; übrige Backends no-op (leeres Set), Verhalten unverändert. Ersetzt Stufe 2 (Supersede-ADR). Anlass: d-migrate-Pilot — die einzige korrekte Vollrichtungs-Config endete an einem Split-Package-Top-Level-Symbol (`asJdbc`) in Exit 2. welle-05-Härtung, slice-031. |
| 0.19.0 | 2026-07-09 | Neu `AC-FA-CLI-002` (Architektur-Graph-Ausgabe): `a-check --print-graph [pfad]` gibt die deklarierte Architektur aus `.a-check.yml` als **Mermaid-Flowchart** auf stdout aus — ein Knoten je Schicht, eine Kante je `edge`, abgesetzte `allow`-Kante, Farbe nach effektiver Rolle; read-only, deterministisch, **kein Scan**. Ladezeitiger Config-Fehler (inkl. unbekannter Sprache), unbekanntes Flag oder Restargument nach dem Pfad → Exit 2; scanzeitige Resolution-Fehler out-of-scope. Eigenständige Inspektions-CLI, **kein** `--print-*`-Ausbau von `AC-FA-DIST-001`. slice-032. |
| 0.20.0 | 2026-07-09 | `AC-FA-DIST-001` erweitert: das `--print-mk`-Fragment `a-check.mk` liefert zusätzlich ein **`a-check-graph`**-Target, das `--print-graph` (`AC-FA-CLI-002`) mit demselben digest-gepinnten `A_CHECK_IMAGE` und netzlosem read-only-Mount aufruft — Mermaid nach stdout, kein Scan, Exit 0; neue AK „Happy (Graph-Target)". Convenience für Konsumenten, die bereits `include a-check.mk` fahren. slice-033. |
| 0.21.0 | 2026-07-24 | Neu **`AC-FA-RULE-009`** (`lateral-slice`: eine `app`-Datei importiert eine fremde Use-Case-Slice — verschiedene `app`-Globs — → kategorischer Befund; opt-in über per-Slice-Globs, ein einziges `app`-Glob inert) und **`AC-FA-RULE-010`** (`port-locality`: eine `app`-Datei importiert einen Port außerhalb dessen pfad-abgeleiteten Scope-Verzeichnisses — use-case-lokal ⊂ business-area ⊂ app-weit — → kategorischer Befund; nur `app`-Importeure, Adapter-Implementierung nicht erfasst). Beide gaten die **Vertical-Slice-Achse** von HexSlice (Doc `hexslice-architecture`); Voraussetzung ist ein als Import-Ziel auflösbarer `app`-/`port`-Glob (literales Präfix, `AC-QA-02`-Grenze). Evidenz: realer HexSlice-Go-Konsument. slice-039. |
| 0.23.0 | 2026-08-09 | **`AC-QA-03` neu gefasst:** das von `--print-mk` **erzeugte** Fragment trägt **keinen konkreten Digest** mehr, sondern einen Platzhalter mit Bezugsquelle ([ADR-0030](../docs/plan/adr/0030-kein-digest-im-generierten-fragment.md)); die im Repo **committete** `a-check.mk` trägt weiterhin den echten Digest. Grund: ein Binary kann den Digest des Image, in dem es läuft, strukturell nicht kennen — er entsteht erst beim Push, das Binary ist vorher gebaut. Der eingebackene Wert nannte darum **immer den Vorgänger** und sah dabei autoritativ aus. Gemessen: `v0.16.0` gab `v0.15.0` aus. **Realer Schaden:** ein Konsument pinnte über den dokumentierten Bump-Weg `v0.15.0` statt `v0.16.0` und vermisste den `constructs`-Block, ohne die Ursache zu sehen. Drei AK ergänzt (Happy/Boundary/Negative). slice-083 (`CR-5`). |
| 0.22.0 | 2026-07-25 | Neu **`AC-FA-RULE-011`** (`construct-leak`: ein optionaler `constructs`-Block hebt die `tech`-Scoping-Mechanik — Zone als Pfad/Pfad-Liste, `match: substring\|regex`, `composition_root: allow\|forbid` — von extrahierten Import-Symbolen auf **Roh-Quelltext**; jedes Vorkommen außerhalb der Zone ist ein Befund, Exit 1). Prüfung **scan-weit** (auch Dateien in keinem `layers`-Glob; `exclude` greift davor), auf derselben Quell-Vorbereitung wie `forbidden_constructs` — in den C-Syntax-Sprachen kommentar-bereinigt (Treffer nur im Kommentar meldet nicht, ausgewiesene Divergenz zur `grep`-Referenz), in Python nicht (`#`-Kommentare bleiben stehen, ausgewiesene Grenze). `AC-FA-CONF-001` um den Block + fail-closed-Decoding erweitert. Damit werden Konstrukte prüfbar, die keine Import-Zeile sind (Aufruf-Monopol `dlopen`); die schichtgebundenen `forbidden_constructs`/`AC-FA-RULE-004` bleiben unberührt. Evidenz: b-cad-P-Rest (Regel P1), Fixture-vermessen. slice-042 (Kandidat 1 aus slice-025). |
| 0.24.0 | 2026-08-30 | Neu **`AC-FA-CLI-003`** (Usage-Ausgabe und Handbuch-Verweis): `--help` trägt Kurzbeschreibung, Aufruf-Syntax, Konfigurations-Hinweis und die **URL des Benutzerhandbuchs**; dieselbe URL steht im Kopfkommentar des per `--print-mk` erzeugten Fragments, das in ein **fremdes** Repo reist. **Zugesichert ist die Anwesenheit der Bestandteile, nicht ihr Wortlaut** — ein Test bindet sich nicht an die Formulierung. Die URL zeigt **tag-frei** auf den Hauptzweig: das Binary kennt seinen Release-Kontext zur Laufzeit nicht, ein zur Build-Zeit eingesetzter versionierter Link nennte immer den Vorgänger — dieselbe Mechanik, die `AC-QA-03` in 0.23.0 für den Image-Digest entschieden hat. Der Preis ist benannt: wer ein altes Release fährt, landet auf dem aktuellen Handbuch und liest dort den Software-Versions-Stempel. slice-117. |
