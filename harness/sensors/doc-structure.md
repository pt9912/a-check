# `make doc-structure` — Struktur-Invarianten innerhalb der Dokumente

## Vertrag

**Fünf Regeln** des Moduls `structure`, konfiguriert in `.d-check.yml`:
Größen-Regel des DoD · Closure-Struktur · Lerneintrag-Form · Kopffelder ·
AC-Form.

Es hat alle vier Eigenbau-Sensoren aus slice-080 abgelöst
(`verify-slice-form`, die strukturelle Hälfte von `verify-closure-notes` und
seit slice-120 `verify-ac-form`), Parität je Befundklasse gemessen.

## Grenze — was das Grün nicht abdeckt

1. **Der Stil einer Anforderung** — dass die drei Akzeptanzkriterien im
   Given/When/Then-Stil formuliert sind, prüft der Lauf nicht; er prüft, dass
   die vier Bausteine **benannt** sind. Review-Sache, permanent.
2. **„Höchstens zwei Schichten"** — dieselbe Grenze: zählbar sind Liefer-Punkte,
   nicht Schichten. Permanent.
3. **`in-progress/`** — der `files`-Glob deckt `done/`; „nur wenn ausgefüllt"
   kann er nicht ausdrücken. Die Gegenstück-Prüfung in `in-progress/` leistet
   `make verify-risiko-ausgaenge` für ihren Ausschnitt. Heilbar erst, wenn das
   Modul einen Zustands-Filter kennt.

## Sperren

- **Braucht den Pin `v0.69.0`** oder neuer: `tasks-ignore-pattern`,
  `exempt-section-pattern` und `exempt-expect-count` sind erst dort vorhanden.
  Ein älterer Pin lässt die Konfiguration fail-closed abbrechen.

## Bindung

`DC-FA-STRUCT-001` · slice-115 (Pin), slice-080 (konfiguriert), slice-120
(vierter Eigenbau abgelöst) · im `verify`-Aggregat.
