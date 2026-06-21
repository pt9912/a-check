# ADR-0003 — Config-Modell `.a-check.yml` (deklarativ, strict-decode, fail-closed)

- **Status:** Proposed
- **Datum:** 2026-06-21
- **Bezug:** [AC-FA-CONF-001](../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml), [AC-FA-CLI-001](../../../spec/lastenheft.md#ac-fa-cli-001--aufruf-scan-wurzel-und-exit-codes)
- **Schärft:** [SPEC-CONF-001](../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema) — macht strict-decode/fail-closed des YAML-Schemas verbindlich.
- **Supersedes:** —

## Kontext

Die Gründungslogik ist „eine Familie driftender Skripte durch ein Werkzeug
ersetzen" (Lastenheft §1): repo-spezifische Schicht-/Tech-Regeln werden
**per Config statt per Skript-Kopie** ausgedrückt
([AC-FA-CONF-001](../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml)).
Die fünf — config-parametrierten — Hexagon-Regeln (`AC-FA-RULE-*`) sind
nachgelagerte Konsumenten dieser Config (Schichten mit Pfad-Mustern, erlaubte
Kanten, Tech→Adapter-Zuordnungen, gemeinsame Adapter-Senke); die begründende
Anforderung dieser ADR ist
[AC-FA-CONF-001](../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml).

## Optionen

1. **YAML, striktes Decoding, fail-closed** (Exit-Code 2 bei unbekanntem
   Schlüssel). Trade-off: streng (kein stiller Default) — aber genau das
   verhindert, dass eine Regel *still nicht angewandt* wird. Den deklarativen
   YAML-Config-Ansatz teilt der Stack bereits (`d-check` wird über
   `.d-check.yml` konfiguriert); strict-decode/fail-closed ist hier `a-check`s
   bewusste Wahl gegen stille Defaults, nicht aus d-check abgeleitet.
2. **YAML mit laxem Decoding** (unbekannte Keys ignoriert). Trade-off:
   bequemer, aber ein Tippfehler im Key deaktiviert *still* eine Regel —
   dieselbe Harness-Lüge-Klasse wie ein undeklariertes Gate (Regelwerk:
   stille Setzung). Verworfen.
3. **Eingebettete DSL/Code-Config** (z. B. Starlark). Trade-off: mächtiger,
   aber overkill und gegen Determinismus/Einfachheit; nicht-deklarativ.

## Entscheidung

**Option 1 — deklaratives YAML `.a-check.yml`, striktes Decoding,
fail-closed mit Exit-Code 2**
([AC-FA-CONF-001](../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml),
[AC-FA-CLI-001](../../../spec/lastenheft.md#ac-fa-cli-001--aufruf-scan-wurzel-und-exit-codes)).
Deklariert werden: Sprache(n) + Datei-Globs je Schicht, die Schichten mit
Pfad-Mustern, die erlaubten Kanten, die Tech→Adapter-Zuordnungen und die
gemeinsame Adapter-Senke. Kein Include/Vererbung in 0.1.0
([AC-FA-CONF-001](../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml)
Out-of-Scope).

## Konsequenzen

- `a-check --print-config` liefert ein kommentiertes `.a-check.yml`-Gerüst
  (read-only; siehe [ADR-0004](0004-distribution-image-mk.md)).
- **Fitness Function / Gate** (slice-003): Negative-Test „Tippfehler im Key
  ⇒ Exit-Code 2", Boundary-Test „fehlende Tech-Zuordnung ⇒ nur
  Schicht-/Lateral-Regeln, kein `tech-leak`"
  ([AC-FA-CONF-001](../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml)).
- Strict-decode *ist* die maschinelle Durchsetzung gegen stille Defaults —
  die Config trägt damit selbst Harness-Disziplin.
