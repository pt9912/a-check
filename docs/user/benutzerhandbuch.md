# Benutzerhandbuch: a-check

**Handbuch-Version:** 1.5 · **Software-Version:** 0.1.0 · **Stand:** 2026-06-22 ·
**Autor:** pt9912 (Maintainer)

---

## 1. Einleitung

### Zweck der Software

**a-check** prüft, ob ein Repository seine **hexagonale Schicht-Architektur**
einhält — sprachübergreifend (C++, Go, Rust, Kotlin), gesteuert über eine
Konfigurationsdatei. a-check liest Ihren Quellcode, meldet Architektur-Verstöße
mit Datei und Zeile und liefert einen Exit-Code, mit dem Sie es als Gate in CI
oder `make` einsetzen. a-check **repariert nichts** und **schreibt nie** in Ihr
Repository.

### Zielgruppe

Repo-Maintainer und CI-Integratoren, die ein einheitliches Architektur-Gate
über mehrere Sprachen wollen. Vorausgesetzt werden Grundkenntnisse in Git,
Docker und `make`; a-check-Interna müssen Sie nicht kennen.

### Voraussetzungen

- **Docker.** a-check läuft ausschließlich als Container — Sie brauchen kein
  lokales Go und keine Sprach-Toolchain.
- Optional **GNU make**, wenn Sie a-check als `make`-Gate einbinden.
- Ein Repository mit erkennbaren Schichten (z. B. `core`, `ports`, `adapters`).

> **Hinweis zum Image.** Im Vorab-Stand (0.1.0) ist noch **kein GHCR-Image
> veröffentlicht**. Bauen Sie es lokal mit `make build`
> ([README](../../README.md)); das erzeugt den Tag **`a-check:dev`**. Ab dem
> Release nutzen Sie das digest-gepinnte GHCR-Image. In allen Beispielen steht
> `<a-check-image>` stellvertretend für beides — im Vorab-Stand also wörtlich
> `a-check:dev`.

## 2. Erste Schritte

### Schnelltest

1. Wechseln Sie in Ihr Repository.
2. Erzeugen Sie ein Konfigurations-Gerüst:
   ```bash
   docker run --rm <a-check-image> --print-config > .a-check.yml
   ```
3. Passen Sie `.a-check.yml` an Ihre Schichten an (Abschnitt 4).
4. Führen Sie a-check aus:
   ```bash
   docker run --rm --network none -v "$PWD:/src:ro" <a-check-image> /src
   ```

### Das Ergebnis verstehen

- **Exit-Code 0** — keine Verstöße. Die Standardausgabe (stdout) bleibt leer;
  auf der Fehlerausgabe (stderr) steht die Zusammenfassung `gesamt: 0 Befund(e)`.
- **Exit-Code 1** — mindestens ein Befund. Jeder Befund steht auf stdout als
  `pfad:zeile: regel: meldung`; die Zusammenfassung (Anzahl je Regel und
  Gesamtzahl) steht auf stderr.
- **Exit-Code 2** — Nutzungs- oder Konfigurationsfehler (z. B. fehlende oder
  ungültige `.a-check.yml`, unbekannte Option).

## 3. Aufgaben

### 3.1 a-check lokal ausführen

**Voraussetzung:** Docker läuft; im Repository liegt eine `.a-check.yml`.

**Vorgehen:**
1. Wechseln Sie in das zu prüfende Repository.
2. Führen Sie aus:
   ```bash
   docker run --rm --network none -v "$PWD:/src:ro" <a-check-image> /src
   ```

**Ergebnis:** a-check listet alle Verstöße und beendet sich mit 0 (sauber)
oder 1 (Befunde).

**Hinweise:** Der Mount `:ro` (read-only) und `--network none` (netzlos) sind
Absicht — a-check braucht keinen Schreibzugriff und keine Netzverbindung.

### 3.2 Eine `.a-check.yml` erstellen

**Voraussetzung:** Sie kennen die Verzeichnisstruktur Ihrer Schichten.

**Vorgehen:**
1. Erzeugen Sie das kommentierte Gerüst:
   ```bash
   docker run --rm <a-check-image> --print-config > .a-check.yml
   ```
2. Tragen Sie unter `languages` Ihre Sprache(n) und Datei-Globs ein.
3. Beschreiben Sie unter `layers` Ihre Schichten mit Pfad-Mustern.
4. Legen Sie unter `edges` die erlaubten Schicht-Kanten fest.

**Ergebnis:** Eine gültige `.a-check.yml` in der Repo-Wurzel. Details zu jedem
Schlüssel: Abschnitt 4.

**Hinweise:** Ein unbekannter Schlüssel oder Tippfehler führt zu Exit-Code 2 —
a-check prüft nie mit geratenen Standardwerten.

### 3.3 a-check als `make`- oder CI-Gate einbinden

**Voraussetzung:** Ihr Repository nutzt `make` oder eine CI-Pipeline.

**Vorgehen:**
1. Erzeugen Sie das einbindbare Makefile-Fragment:
   ```bash
   docker run --rm <a-check-image> --print-mk > a-check.mk
   ```
2. Binden Sie es in Ihr `Makefile` ein:
   ```makefile
   include a-check.mk
   ```
3. Rufen Sie das Gate auf:
   ```bash
   make a-check
   ```

**Ergebnis:** `make a-check` prüft das Repository netzlos und read-only und
schlägt bei Befunden fehl (Exit-Code 1).

**Hinweise:** Das Fragment pinnt das Image über `A_CHECK_IMAGE`. **Im
Vorab-Stand** (noch kein GHCR-Image) bauen Sie zuerst `make build` und
überschreiben das Image beim Aufruf:
```bash
make a-check A_CHECK_IMAGE=a-check:dev
```
Ab dem Release referenziert `A_CHECK_IMAGE` einen `@sha256:`-Digest; heben Sie
den Pin bewusst per Commit an, damit CI-Läufe reproduzierbar bleiben. Vergleiche
das mitgelieferte [`a-check.mk`](../../a-check.mk) dieses Repos. Den
Release-Prozess (Tagging, Digest-Pin, GHCR) beschreibt [`releasing.md`](releasing.md).

### 3.4 Befunde lesen und beheben

Jeder Befund nennt die Regel. Die sechs Regeln und ihre Behebung:

| Regel | Bedeutung | Behebung |
|---|---|---|
| `core-impurity` | Der Kern (`role: domain`) importiert einen Port, eine `app`- oder Adapter-Schicht oder ein Framework/Tech — die Domäne ist die innerste Schicht. | Domäne rein halten; Port-/Use-Case-Orchestrierung in eine `app`-Schicht, Tech nur im Adapter. |
| `app-impurity` | Die Application-Schicht (`role: app`) importiert einen Adapter oder ein Framework/Tech (Domäne + Ports darf sie nutzen). | Tech/Adapter hinter einen Port legen; die App spricht nur Domäne + Ports. |
| `lateral-adapter` | Ein Adapter importiert einen anderen Adapter. | Gemeinsame Logik in die konfigurierte Senke (`adapter_sink`) ziehen oder über einen Port führen. |
| `tech-leak` | Ein Framework/Tech erscheint außerhalb seines Adapters. | Den Tech-Zugriff in den zugeordneten Adapter kapseln. |
| `port-impurity` | Ein Port importiert einen Adapter oder ein Framework/Tech, oder enthält ein per `forbidden_constructs` (Abschnitt 4) verbotenes Konstrukt. Domänentypen des Kerns darf ein Port referenzieren. | Den Port von Adapter-/Tech-Importen befreien (Kern-Referenzen sind erlaubt). |
| `wrong-direction` | Ein Import läuft entgegen einer erlaubten Schicht-Kante. | Die Kante in `edges` aufnehmen (falls legitim) oder den Import umdrehen. |

### 3.5 Heuristik-Ausnahmen konfigurieren

a-check erkennt Importe **text-heuristisch**, nicht über einen vollständigen
Parser. Selten wird ein harmloses Symbol fälschlich erkannt (z. B. ein
framework-fremdes `Queue.h`). In diesem Fall tragen Sie es in die Allowlist ein:

```yaml
markers:
  ignore_symbols: ["Queue.h"]
```

`ignore_symbols` wirkt auf erkannte **Importe** (z. B. falsch-positive
`core-impurity`/`tech-leak`); ein per `forbidden_constructs` verbotenes Konstrukt
wird davon nicht erfasst.

## 4. Konfiguration (`.a-check.yml`)

Die Datei liegt in der Repo-Wurzel und wird **streng** dekodiert. Beispiel:

```yaml
version: 1
languages:
  go: ["**/*.go"]                 # Sprache -> Datei-Globs
layers:
  core:     ["internal/core/**"]  # Schicht -> Pfad-Muster
  ports:    ["internal/ports/**"]
  adapters: ["internal/adapters/**"]
edges:
  - {from: adapters, to: ports}   # erlaubte gerichtete Kante
  - {from: ports,    to: core}    # Ports dürfen Domänentypen referenzieren
  # - {from: adapters, to: core}  # falls Adapter Domänentypen direkt referenzieren
adapter_sink: driver-common       # gemeinsame Adapter-Senke (optional)
tech:
  - {pattern: "net/http", adapter: http}   # Tech -> Adapter (optional)
composition_root: ["cmd/**"]      # verdrahtet alles, von Regeln ausgenommen (optional)
allow:                            # explizit erlaubte Sonderkanten/Re-Exports (optional)
  - {from: ports, to: ports}
forbidden_constructs:             # Schicht -> verbotene Text-Muster (Port-Disziplin, optional)
  ports: ["impl "]
markers:
  ignore_symbols: []              # Heuristik-Ausnahmen (optional)
```

**Pflichtblöcke:** `version`, `languages`, `layers`, `edges`.
**Gültige `languages`-Schlüssel:** genau `go`, `cpp`, `rust`, `kotlin` — exakt so
zu schreiben (z. B. `cpp`, **nicht** `c++`); andere Schlüssel werden ignoriert
(keine Extraktion). Jeder Schlüssel bildet auf eine Liste von Datei-Globs ab,
z. B. `cpp: ["**/*.h", "**/*.cpp"]`, `rust: ["**/*.rs"]`, `kotlin: ["**/*.kt"]`.
**Optionalblöcke:** `adapter_sink`, `tech`, `composition_root`, `allow`,
`forbidden_constructs`, `markers`. Fehlt ein Optionalblock, entfällt die
zugehörige Prüfung (kein stiller Standardwert) — fehlt z. B. `adapter_sink`,
darf **kein** Adapter einen anderen importieren (strengere Auslegung). Das
vollständige Schema steht in der [Spezifikation](../../spec/spezifikation.md).

## 5. Berechtigungen und Sicherheit

a-check kennt keine Benutzerrollen — es ist ein Kommandozeilen-Werkzeug. Statt
Rechten gelten Garantien:

- **Read-only:** a-check schreibt nie in das geprüfte Repository (Mount mit `:ro`).
- **Netzlos:** mit `--network none` öffnet a-check keine Netzverbindungen.
- **Hermetisch:** Das Image ist distroless/static und digest-gepinnt — gleicher
  Lauf, gleiches Ergebnis.

Geben Sie keine Zugangsdaten oder Tokens an a-check — es benötigt keine.

## 6. Fehlerbehebung

### Fehler: Docker findet das Image nicht (`Unable to find image` / `pull access denied`)

**Ursache:** Im Vorab-Stand (0.1.0) ist kein GHCR-Image veröffentlicht;
`ghcr.io/pt9912/a-check:0.1.0` ist nicht abrufbar.

**Lösung:** Bauen Sie das Image lokal mit `make build` und verwenden Sie
`a-check:dev` — in `docker run`-Aufrufen als `<a-check-image>`, im Gate über
`make a-check A_CHECK_IMAGE=a-check:dev`.

### Fehler: a-check bricht mit Exit-Code 2 ab

**Ursache:** Die `.a-check.yml` fehlt, ist ungültig oder enthält einen
unbekannten Schlüssel; oder es wurde eine unbekannte Option übergeben.

**Lösung:**
1. Prüfen Sie, ob `.a-check.yml` in der Scan-Wurzel liegt.
2. Lesen Sie die Fehlermeldung auf der Fehlerausgabe (sie nennt die Zeile).
3. Vergleichen Sie mit dem Gerüst aus `--print-config`.

### Fehler: a-check findet nichts, obwohl Verstöße erwartet werden

**Ursache:** Die `layers`- oder `languages`-Globs passen nicht auf Ihre Pfade.

**Lösung:**
1. Prüfen Sie, ob die Globs (z. B. `internal/core/**`) Ihre echten Verzeichnisse treffen.
2. Prüfen Sie, ob die Datei-Endung unter `languages` erfasst ist.

### Fehler: ein `tech-leak`/`core-impurity`-Befund ist falsch-positiv

**Ursache:** Ein gleichnamiges, aber framework-fremdes Symbol (Heuristik-Grenze).

**Lösung:** Tragen Sie das Symbol in `markers.ignore_symbols` ein (Abschnitt 3.5).

## 7. FAQ

**Brauche ich Go installiert?** Nein. a-check läuft als Container; Docker genügt.

**Verändert a-check meinen Code?** Nein. a-check ist read-only und meldet nur.

**Warum hat a-check eine Heuristik-Grenze?** Es liest Importe text-basiert (kein
vollständiger Parser je Sprache) — das hält den Lauf hermetisch und schnell. Die
Grenze ist dokumentiert; Ausnahmen sind konfigurierbar.

**Kann ich mehrere Sprachen in einem Repo prüfen?** Ja — tragen Sie mehrere
Einträge unter `languages` ein.

## 8. Glossar

- **Kern (core):** die reine Domänenlogik ohne I/O, Framework oder Ports (innerste Schicht — kennt nur sich selbst).
- **Port:** eine Schnittstelle/Abstraktion, über die der Kern mit der Außenwelt spricht.
- **Adapter:** die konkrete Anbindung an Technik (Datenbank, HTTP, UI …).
- **Composition Root:** der Ort, der konkrete Adapter an den Kern verdrahtet (z. B. `main`); von den Schicht-Regeln ausgenommen.
- **Schicht:** eine über Pfad-Muster (`layers`) definierte Datei-Gruppe (z. B. `core`, `ports`, `adapters`).
- **Kante (`edges`):** eine erlaubte gerichtete Abhängigkeit zwischen zwei Schichten (`from` → `to`).
- **`adapter_sink`:** eine gemeinsame Senke, die alle Adapter importieren dürfen (Ausnahme von `lateral-adapter`).
- **`forbidden_constructs`:** je Schicht konfigurierte verbotene Text-Muster (für `port-impurity`).
- **Befund:** eine gemeldete Regelverletzung (Datei, Zeile, Regel, Meldung).
- **`core-impurity` / `app-impurity` / `lateral-adapter` / `tech-leak` / `port-impurity` / `wrong-direction`:** die sechs geprüften Regeln (Abschnitt 3.4).
- **Heuristik-Grenze:** a-check erkennt Importe per Textmuster, nicht per Parser; seltene Fehltreffer sind konfigurierbar ausnehmbar.
- **Digest-Pin:** ein `@sha256:`-Verweis auf eine exakte Image-Version für reproduzierbare Läufe.

## 9. Support und Kontakt

Quellcode, Issues und Releases: das Projekt-Repository `pt9912/a-check`.
Verbindlich für das Verhalten sind das [Lastenheft](../../spec/lastenheft.md)
und die [Spezifikation](../../spec/spezifikation.md); ein Überblick steht in der
[README](../../README.md).

## 10. Änderungshistorie

| Handbuch-Version | Stand | Änderung |
|---|---|---|
| 1.0 | 2026-06-21 | Erstfassung zur Software-Version 0.1.0. |
| 1.1 | 2026-06-21 | Review-Einarbeitung: Vorab-Image-Pfad fürs make-Gate (`A_CHECK_IMAGE=a-check:dev`), Config-Schlüssel `allow`/`forbidden_constructs`, Exit-0-stderr-Klarstellung, Image-Fehlerfall, Glossar, Autor. |
| 1.2 | 2026-06-21 | Quer-Verweis aus §3.3 auf den neuen Release-Leitfaden [`releasing.md`](releasing.md). |
| 1.3 | 2026-06-21 | §4: die vier gültigen `languages`-Schlüssel (`go`/`cpp`/`rust`/`kotlin`) explizit gelistet; Software-Version 0.1.0 veröffentlicht. |
| 1.4 | 2026-06-22 | §3.4/§4 an Lastenheft 0.2.0 angeglichen: `port-impurity` — Ports dürfen Domänentypen des Kerns referenzieren (verboten bleiben Adapter/Tech); `ports`-Schicht + `ports → core`-Kante im Beispiel. |
| 1.5 | 2026-06-22 | §3.4/Glossar an Lastenheft 0.5.0 angeglichen: neue Regel `app-impurity` (Rolle `app`); `core-impurity` verschärft — die Domäne kennt keine Ports (`domain↛port` kategorisch); sechs Regeln. |
