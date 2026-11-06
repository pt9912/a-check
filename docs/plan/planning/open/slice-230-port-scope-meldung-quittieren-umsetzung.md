# slice-230 — Quittieren der Port-Scope-Meldung: Umsetzung

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `git mv` (`make slice-mv`).

**Welle:** ohne Welle — die Closure-Bedingung wäre die eigene DoD.

**Bezug:** GitHub-Issue #7, [slice-229](../done/wellenlos/slice-229-port-scope-meldung-quittieren-anforderung.md)
(Anforderung und ADR — die Vorlage dieses Slice),
[ADR-0040](../../adr/0040-portscope-richtungssegment.md) (Port-Scope-Diagnose).

**Berührte Spec-Stellen:** [SPEC-CONF-001](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema) (Schema) und
[SPEC-CLI-001](../../../../spec/spezifikation.md#spec-cli-001--aufruf-scan-wurzel-und-exit-codes) (Port-Scope-Diagnose) — beide werden hier nachgezogen.

**Verantwortlich:** pt9912

**Autor:** Claude. **Datum:** 2026-11-05.

**Lerneintrag — Form:** wird bei Closure benannt.

---

## 1. Ziel und Abgrenzung

**Ziel:** Eine Port-Schicht kann die Port-Scope-Meldung quittieren, wie es die Anforderung und die
ADR aus slice-229 festlegen: Schema-Schlüssel mit Pflicht-Begründung, kurze Meldungszeile mit
Schichtnamen, Hinweis auf eine überflüssige Quittierung, Exit-Code unberührt.

Form, Schlüsselname und Fehlerverhalten stehen in
[ADR-0045](../../adr/0045-port-scope-meldung-quittieren.md) und
[AC-FA-CONF-002](../../../../spec/lastenheft.md#ac-fa-conf-002--port-scope-meldung-je-schicht-quittieren);
§3 leitet den Plan daraus ab.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Anforderung und ADR.** *Ein Folge-Slice hat sie geliefert:* slice-229; dieser Slice setzt sie
  um und ändert sie nicht.
- **Quittieren anderer Diagnosen.** *Bestand bleibt bewusst stehen*: ohne zweiten Beleg nicht
  zugesagt (slice-229 §1).
- **Ein Release.** *Ein anderer Vorgang*: ein Release folgt, wenn die Änderung ausgeliefert werden
  soll; sie steht bis dahin in `[Unreleased]`.

## 2. Definition of Done

- [ ] Schema-Schlüssel dekodiert und validiert; Verstoß wie in der ADR vorgesehen.
- [ ] Meldungszeile je quittierter Schicht und Hinweis auf überflüssige Quittierung; Exit-Code
      und `port-locality` unverändert (Test).
- [ ] Spezifikation, Benutzerhandbuch und CHANGELOG `[Unreleased]` nachgezogen.
- [ ] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [ ] Beobachtungs-Register fortgeschrieben (oder „keine Beobachtung" notiert).
- [ ] Jedes Risiko aus §6 trägt einen Ausgang.

`make gates` und `make verify` grün. Ein unabhängiger Review ist vorgesehen
([MR-019](../../../../harness/conventions.md#mr-019): Opt-in) — der Slice ändert Schema, Kern und
Ausgabe.

## 3. Plan (vor Code)

Drei Liefer-Punkte, je einer pro Schicht; die Reihenfolge ist Config → Kern → Ausgabe.

1. **Config-Adapter** (`internal/adapter/driven/config/config.go`): `yamlLayer` und `knownLayerKey`
   kennen `acknowledge`; `decodeLayerObject` validiert es gegen die **effektive** Rolle (Decoder-
   Inferenz) und liefert die Begründung. Fehler (alle Exit 2, Meldung nennt Schicht und Fehler):
   Rolle nicht `port`; kein Mapping, leeres Mapping, doppelter Schlüssel, anderer Schlüssel als
   `port-scope`; Begründung fehlt, `null`, kein String (YAML-Tag `!!str`) oder nur Leerraum.
   `core.Layer` bekommt das Feld `AckPortScope string`.
2. **Kern** (`internal/hexagon/core/rules.go`): `InertPortScope` bekommt `Layers []string` — alle
   Schichten, die den Glob tragen, in Deklarationsreihenfolge; die Entdoppelung je Glob bleibt für
   die Zeilen, die Zuordnung läuft vor ihr. Neue reine Funktion `PortScopeAcks(m, scopes)` liefert
   `Acked` (Schicht, Begründung, auslösende Globs, Ursachen sortiert, in Schicht-Reihenfolge),
   `Redundant` (quittierte Schichten ohne auslösenden Glob) und `Open` (Globs, die **nicht jede**
   tragende Schicht quittiert — nach Glob sortiert). `port-locality`, Befunde und Exit-Code werden
   nicht berührt.
3. **Ausgabe** (`internal/cli/cli.go`, `writePortScopeNotice`): Kopfzeile zählt alle
   auslösenden Globs; dann die Zeilen der quittierten Schichten, dann die Glob-Zeilen aus `Open`
   mit Deckel und „… und n weitere“, dann die Abhilfe-Zeilen nur für Ursachen in `Open`; die
   Zeile „Hinweis: Schicht … Quittierung ohne Anlass …“ steht auch dann, wenn sonst keine Meldung
   entsteht (frühe Rückkehr bei null Treffern entfällt dafür). Wortlaut aus ADR-0045 Punkt 2/3.

**Tests (vor/mit dem Code, je Mutations-Probe auf den Befund, nicht den Exit-Code):**
Config: jeder Fehlerfall aus ADR Punkt 4 und der Happy-Fall. Kern: zwei Schichten mit einer
Quittierung; Teil-Auslösung; geteilter Glob (keine Schicht gilt als überflüssig); überflüssig ohne
sonstige Meldung. CLI: Wortlaut, Reihenfolge, Deckel nur für nicht quittierte, Kopfzeile zählt
alles, Abhilfe nur bei offener Ursache; derselbe Baum mit und ohne Quittierung ergibt gleiche
Befunde und gleichen Exit-Code; Determinismus über zwei Läufe.

**Doku:** SPEC-CONF-001 (Schlüssel, Fehler), SPEC-CLI-001 (Zeilenformen), Spez-Bump mit
Historie-Zeile; Benutzerhandbuch dort, wo der Port-Scope-Hinweis steht (inkl. Pin-Hinweis: Pin erst
mit dem Release heben, das den Schlüssel kennt), Handbuch-Historie; CHANGELOG `[Unreleased]`
(`### Added`).

## 4. Trigger

**Start** (`open` → `next`): [slice-229](../done/wellenlos/slice-229-port-scope-meldung-quittieren-anforderung.md)
liegt in `done/`.

**Rückführungen:** `in-progress` → `next` (zu groß): Schema, Kern und Ausgabe sprengen eine
Review-Sitzung — dann an der Schicht-Grenze teilen.

## 5. Closure-Trigger

DoD vollständig, Gates grün, Closure-Notiz mit Lerneintrag.

## 6. Risiken und offene Punkte

- **Der Schnitt trägt nicht:** drei Schichten (Config-Adapter, Kern, Ausgabe) in einem Slice. —
  **Ausgang:** *(bei Closure zuzuweisen: eingetreten / entfallen / weiter offen)*

## 7. Closure-Notiz

*(folgt bei Closure)*

## 8. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** `KERN` (Achsen 1, 2, 3 ✓) und `ADAPT` (Achsen 2, 3 ✓).

**Vorgelagert — offene Beobachtungen sichten** (2026-11-05): `BEO-KERN/dirvocab-portfor-auseinander`
(2×, offen) — berührt diesen Slice, wenn der Schlüssel Richtungen kennt; der Zähler steht bei 2×
und erreichte mit einem dritten Beleg die Schwelle (Lücke, eigener Folge-Slice).

**Modus-Begründungsblock:** alle berührten Sub-Areas GF.
