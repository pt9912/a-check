# `make doc-structure` — Struktur-Invarianten innerhalb der Dokumente

## Vertrag

Die Regeln des Moduls `structure`, konfiguriert in `.d-check.yml`:
Größen-Regel des DoD · Closure-Struktur · Lerneintrag-Form · Kopffelder ·
AC-Form · Zellengrenzen der beiden Gate-Tabellen. **Keine Zahl hier** — sie
stand an drei Stellen und hinkte der Konfiguration hinterher, sobald eine Regel
dazukam (`seit slice-181`).

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
4. **Sätze in Tabellenzellen** — die Zellengrenze misst **Zeichen**. Sätze wären
   das Kriterium der Ziel-Form, aber `table.column` kann sie nicht zählen;
   gemessen tragen **16** der 66 Vertrags-/Zweck-Zellen mehr als einen Satz, die
   meisten davon unter der Schwelle. Diese Hälfte hängt am Review.
5. **Die Inhaltsspalte der Nicht-Gates-Tabelle** (`Was es tut`) ist
   **unadressiert** — dort greift nur die Untergrenze über `Bindung`. Eine
   400-Zeichen-Zelle passiert grün. Benannt, nicht gedeckt.

## Sperren

- **Braucht den Pin `v0.69.0`** oder neuer: `tasks-ignore-pattern`,
  `exempt-section-pattern` und `exempt-expect-count` sind erst dort vorhanden.
  Ein älterer Pin lässt die Konfiguration fail-closed abbrechen.

## Bindung

`DC-FA-STRUCT-001` · slice-115 (Pin), slice-080 (konfiguriert), slice-120
(vierter Eigenbau abgelöst) · im `verify`-Aggregat.
