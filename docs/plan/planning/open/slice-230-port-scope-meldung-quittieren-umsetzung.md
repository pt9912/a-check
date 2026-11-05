# slice-230 — Quittieren der Port-Scope-Meldung: Umsetzung

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `git mv` (`make slice-mv`).

**Welle:** ohne Welle — die Closure-Bedingung wäre die eigene DoD.

**Bezug:** GitHub-Issue #7, [slice-229](../next/slice-229-port-scope-meldung-quittieren-anforderung.md)
(Anforderung und ADR — die Vorlage dieses Slice),
[ADR-0040](../../adr/0040-portscope-richtungssegment.md) (Port-Scope-Diagnose).

**Berührte Spec-Stellen:** [SPEC-CONF-001](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema) (Schema) und
[SPEC-CLI-001](../../../../spec/spezifikation.md#spec-cli-001--aufruf-scan-wurzel-und-exit-codes) (Port-Scope-Diagnose) — beide werden hier nachgezogen.

**Verantwortlich:** —

**Autor:** Claude. **Datum:** 2026-11-05.

**Lerneintrag — Form:** wird bei Closure benannt.

---

## 1. Ziel und Abgrenzung

**Ziel:** Eine Port-Schicht kann die Port-Scope-Meldung quittieren, wie es die Anforderung und die
ADR aus slice-229 festlegen: Schema-Schlüssel mit Pflicht-Begründung, kurze Meldungszeile mit
Schichtnamen, Hinweis auf eine überflüssige Quittierung, Exit-Code unberührt.

Der Slice ist **nicht planbar, bevor slice-229 in `done/` liegt**: Form, Schlüsselname und
Fehlerverhalten stehen erst dort. §3 und die DoD werden beim `open` → `next` an die ADR
angepasst — das ist die Plan-Änderung, die AGENTS §6 vor den Code stellt.

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

*(wird beim `open` → `next` aus der ADR von slice-229 abgeleitet)*

## 4. Trigger

**Start** (`open` → `next`): [slice-229](../next/slice-229-port-scope-meldung-quittieren-anforderung.md)
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
