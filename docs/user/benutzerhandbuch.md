# Benutzerhandbuch: a-check

**Handbuch-Version:** 1.0 · **Software-Version:** 0.1.0 (in Entwicklung) ·
**Stand:** 2026-06-21

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

> **Hinweis zum Image.** Im Vorab-Stand (0.1.0) bauen Sie das Image lokal
> (`make build` → `a-check:dev`, siehe [README](../../README.md)). Ab dem
> Release nutzen Sie das digest-gepinnte GHCR-Image. In allen Beispielen steht
> `<a-check-image>` stellvertretend für beides.

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

- **Exit-Code 0** — keine Verstöße. a-check gibt nichts auf der Standardausgabe aus.
- **Exit-Code 1** — mindestens ein Befund. Jeder Befund steht auf der
  Standardausgabe als `pfad:zeile: regel: meldung`; eine Zusammenfassung (Anzahl
  je Regel) erscheint auf der Fehlerausgabe.
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

**Hinweise:** Das Fragment pinnt das Image. Ab dem Release referenziert
`A_CHECK_IMAGE` einen `@sha256:`-Digest; heben Sie den Pin bewusst per Commit an,
damit CI-Läufe reproduzierbar bleiben. Vergleiche das mitgelieferte
[`a-check.mk`](../../a-check.mk) dieses Repos.

### 3.4 Befunde lesen und beheben

Jeder Befund nennt die Regel. Die fünf Regeln und ihre Behebung:

| Regel | Bedeutung | Behebung |
|---|---|---|
| `core-impurity` | Der Kern importiert einen Adapter oder ein Framework/Tech. | Abhängigkeit über einen Port (Schnittstelle) umkehren; Tech nur im Adapter nutzen. |
| `lateral-adapter` | Ein Adapter importiert einen anderen Adapter. | Gemeinsame Logik in die konfigurierte Senke (`adapter_sink`) ziehen oder über einen Port führen. |
| `tech-leak` | Ein Framework/Tech erscheint außerhalb seines Adapters. | Den Tech-Zugriff in den zugeordneten Adapter kapseln. |
| `port-impurity` | Ein Port importiert Kern/Adapter oder enthält ein verbotenes Konstrukt. | Den Port auf reine Abstraktionen reduzieren. |
| `wrong-direction` | Ein Import läuft entgegen einer erlaubten Schicht-Kante. | Die Kante in `edges` aufnehmen (falls legitim) oder den Import umdrehen. |

### 3.5 Heuristik-Ausnahmen konfigurieren

a-check erkennt Importe **text-heuristisch**, nicht über einen vollständigen
Parser. Selten wird ein harmloses Symbol fälschlich erkannt (z. B. ein
framework-fremdes `Queue.h`). In diesem Fall tragen Sie es in die Allowlist ein:

```yaml
markers:
  ignore_symbols: ["Queue.h"]
```

## 4. Konfiguration (`.a-check.yml`)

Die Datei liegt in der Repo-Wurzel und wird **streng** dekodiert. Beispiel:

```yaml
version: 1
languages:
  go: ["**/*.go"]                 # Sprache -> Datei-Globs
layers:
  core:     ["internal/core/**"]  # Schicht -> Pfad-Muster
  adapters: ["internal/adapters/**"]
edges:
  - {from: adapters, to: core}    # erlaubte gerichtete Kante
adapter_sink: driver-common       # gemeinsame Adapter-Senke (optional)
tech:
  - {pattern: "net/http", adapter: http}   # Tech -> Adapter (optional)
composition_root: ["cmd/**"]      # verdrahtet alles, von Regeln ausgenommen (optional)
markers:
  ignore_symbols: []              # Heuristik-Ausnahmen (optional)
```

Pflichtblöcke: `version`, `languages`, `layers`, `edges`. Alle übrigen Blöcke
sind optional; fehlt einer, entfällt die zugehörige Prüfung. Das vollständige
Schema mit allen Schlüsseln und Defaults steht in der
[Spezifikation](../../spec/spezifikation.md).

## 5. Berechtigungen und Sicherheit

a-check kennt keine Benutzerrollen — es ist ein Kommandozeilen-Werkzeug. Statt
Rechten gelten Garantien:

- **Read-only:** a-check schreibt nie in das geprüfte Repository (Mount mit `:ro`).
- **Netzlos:** mit `--network none` öffnet a-check keine Netzverbindungen.
- **Hermetisch:** Das Image ist distroless/static und digest-gepinnt — gleicher
  Lauf, gleiches Ergebnis.

Geben Sie keine Zugangsdaten oder Tokens an a-check — es benötigt keine.

## 6. Fehlerbehebung

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

- **Kern (core):** die reine Domänenlogik ohne I/O oder Framework.
- **Port:** eine Schnittstelle/Abstraktion, über die der Kern mit der Außenwelt spricht.
- **Adapter:** die konkrete Anbindung an Technik (Datenbank, HTTP, UI …).
- **Composition Root:** der Ort, der konkrete Adapter an den Kern verdrahtet (z. B. `main`); von den Schicht-Regeln ausgenommen.
- **`core-impurity` / `lateral-adapter` / `tech-leak` / `port-impurity` / `wrong-direction`:** die fünf geprüften Regeln (Abschnitt 3.4).
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
