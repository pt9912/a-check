# slice-017 — Unbekannter `languages`-Schlüssel → Exit 2 (falsch-grün-Falle schließen)

**Status:** open (Backlog — Härtung).
**Bezug:** schärft [AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml)
(strict-decode) + [AC-FA-EXTRACT-001](../../../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion)
(bekannte Backend-Menge); Motiv [AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)
(ehrliche Grenze). [Roadmap](../in-progress/roadmap.md). **Evidenz:** b-cad-Pilot + Polyglot-Bestand
(Go/Python/C#/TypeScript).

> **Backlog-Stub.** Kleine Härtung, kein Entwurf zur Abnahme. Wird zum Slice ausgearbeitet, sobald
> ein Konsument mit einer noch nicht unterstützten Sprache pilotiert (Trigger §1).

## 1. Auslöser

Ein `languages`-Schlüssel außerhalb der unterstützten Backends (`cpp`/`go`/`rust`/`kotlin`/`java`)
wird heute **still ignoriert**: die Extraktion dispatcht per `switch` und trifft `default: return nil`
(`internal/adapter/driven/extract/extract.go`) — der Config-Decode akzeptiert den Schlüssel
kommentarlos. Folge: `languages: {python: ["**/*.py"]}` extrahiert **nichts** → `0 Befunde` →
**falsch-grün**. Das widerspricht dem strict-decode-/„keine stillen Defaults"-Ethos (jeder unbekannte
*Schlüssel* bricht sonst mit Exit 2 ab) und ist gefährlicher als ein sichtbarer Fehler, weil es
Sicherheit vortäuscht.

## 2. Geplanter Umfang

1. **Bekannte Backend-Menge als Single Source** (heute implizit im `importsFromSource`-`switch`) —
   z. B. eine exportierte `core`-Konstante/Funktion, gegen die validiert wird.
2. **Config-Decode validiert** jeden `languages`-Schlüssel gegen diese Menge; unbekannt → **Exit 2**
   mit klarer Meldung (analog zu ungültiger `role`/`direction`/`match`).
3. [AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml)
   Negative-AC erweitern; [SPEC-CONF-001](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema)
   nennt die zulässige Backend-Menge normativ.
4. Tests: unbekannte Sprache → Exit 2; jede unterstützte lädt.

## 3. Vor der Umsetzung zu klären

- **Ort der Validierung:** der Config-Adapter kennt die Backend-Liste heute nicht (die lebt im
  Extraktions-Adapter). Sauber: die Menge im `core`/`port` deklarieren, damit Config und Extraktion
  **eine** Quelle teilen (sonst driften sie).
- Wechselwirkung mit künftigen Backends: jede neue Sprache ([slice-014](../done/slice-014-java-backend.md)-Muster)
  erweitert die Menge an **einer** Stelle.
