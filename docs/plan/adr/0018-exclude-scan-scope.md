# ADR-0018 — `exclude`-Scan-Scope: Top-Level-Datei-Ausschluss vor der Extraktion

- **Status:** Accepted
- **Datum:** 2026-07-03
- **Autor:** pt9912
- **Bezug:** [AC-FA-CONF-001](../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml) (Lastenheft 0.14.0, CR d-check-Pilot 3/3) — additiv zu [ADR-0003](0003-config-modell-a-check-yml.md) (Config-Modell).
- **Schärft:** [SPEC-CONF-001](../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema) + [SPEC-EXTRACT-001](../../../spec/spezifikation.md#spec-extract-001--import-extraktion).
- **Supersedes:** —

## Kontext

Der Scanner erfasst jede Datei, die einem `languages`-Glob entspricht; die Glob-Engine kennt
**bewusst keine Negation**. Zwei reale Konsumenten zeigen, dass das ohne Ausschluss-Mechanik
nicht trägt ([slice-023](../planning/done/slice-023-dcheck-pilot-deltas.md) §1/§3):

1. **d-check** (Go-Pilot): das abgelöste `go list`-Gate prüfte nur Nicht-Test-Imports — schon
   ein `os`-Import in einem Adapter-**Test** macht den sauberen Baum rot; `**/*_test.go` ist
   mit `languages`-Globs nicht ausschließbar.
2. **m-trace** (pnpm-TS-Monorepo mit Go-Hexagon-App): `node_modules/` liegt **je Workspace**,
   dazu `dist/`-Generat und `test-results/` — ohne Ausschluss lägen dort 296 statt 94 eigene
   `.ts`-Dateien im Scope (zwei Drittel Fremdcode); `**/*.ts` matcht zudem `*.d.ts`-Generat.
   Für TypeScript-Konsum ist der Ausschluss **Existenzbedingung**, nicht Test-Hygiene.

Die Ausschluss-Klassen sind gemischt sprachübergreifend (`node_modules/`, `dist/`,
`test-results/`) und sprachgebunden nur per Konvention (`*_test.go`, `*.test.ts`).

## Optionen

| Weg | Idee | Bewertung |
|---|---|---|
| **A — Top-Level-`exclude`, vor der Extraktion** | Eine Datei-Glob-Liste neben `languages`; matchende Dateien fallen **vor** der Extraktion vollständig aus dem Scan. | **Gewählt.** Ein Scan hat *einen* Scope; die repo-weiten Verzeichnis-Klassen stehen genau einmal; wirkt uniform für Imports **und** `forbidden_constructs`; ohne Block byte-identisch. |
| **B — `exclude` je Sprache** | Ausschluss-Liste in jedem `languages`-Eintrag. | Verworfen: dupliziert `node_modules/`/`dist/` je Sprache und lückt bei jeder neuen Sprache **still** (m-trace-Evidenz); die sprachgebundenen Muster (`*_test.go`) leben in einer Top-Level-Liste genauso präzise. |
| **C — Negation in der Glob-Engine** (`!`-Muster) | `languages`-Globs um Negation erweitern. | Verworfen: bricht die bewusst einfache, deterministische Glob-Semantik (Reihenfolge-Abhängigkeit von Negationen ist eine bekannte Fehlerquelle); `exclude` ist das explizite, getrennte Gegenstück. |

## Entscheidung

**Weg A:**

1. **`exclude`** ist ein optionaler **Top-Level**-Block: eine Liste von Datei-Globs
   **relativ zur Scan-Wurzel** (dieselbe Glob-Semantik wie `layers`/`languages`).
2. **Wirkung vor der Extraktion:** eine matchende Datei wird nicht gelesen — sie existiert
   für keine Regel (weder Import- noch `forbidden_constructs`-Prüfung). Die Reihenfolge ist
   damit scharf: `exclude` (Scan-Scope) → Extraktion → Regel-Auswertung (dort erst
   `composition_root`, das nur Regel-Ausnahmen steuert).
3. **Fail-closed:** ein leerer Glob (`""`) ist der ungültige Fall der totalen Glob-Engine →
   Exit 2 (kein stiller No-Op, Ethos-Linie von
   [slice-017](../planning/done/slice-017-unbekannte-sprache-exit2.md)).
4. **Rückwärtskompatibel:** ohne `exclude`-Block wird jede `languages`-Glob-Datei gescannt —
   byte-identische Ausgabe.
5. **Ehrliche Grenze ([AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)):**
   `exclude` ist Konsumenten-Konfiguration — wer zu breit ausschließt, schwächt sein eigenes
   Gate; a-check weist den Scope nicht still aus, die Config ist die deklarierte Wahrheit
   (wie `markers.ignore_symbols`).

## Konsequenzen

- [ADR-0003](0003-config-modell-a-check-yml.md) bleibt gültig — `exclude` ist ein additiver
  Optionalblock im bestehenden Config-Modell; die Glob-Engine bleibt negations-frei.
- **Schema** ([SPEC-CONF-001](../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema)):
  Optionalblock `exclude: ["**/*_test.go", …]`; **Scan-Scope**
  ([SPEC-EXTRACT-001](../../../spec/spezifikation.md#spec-extract-001--import-extraktion)):
  Ausschluss-Schritt vor der Extraktion.
- Erste Nutzer: d-check (`**/*_test.go`), m-trace-Klasse (`**/node_modules/**`, `**/dist/**`,
  `**/*.d.ts`).

## Fitness Function

- `make test`: Verstoß **nur** in einer ausgeschlossenen Datei → kein Befund/Exit 0; Config
  ohne `exclude` → byte-identische Ausgabe; leerer Glob → Exit 2; `exclude` wirkt auch auf
  `forbidden_constructs`-Treffer (Datei nicht gelesen).
- `make arch-check` (Dogfooding): unverändert 0 (Eigen-Config deklariert kein `exclude`).

## Re-Evaluierungs-Trigger

- **Verzeichnis-Kurzformen** (`node_modules/` statt `**/node_modules/**`), falls Piloten
  über die Glob-Länge stolpern — reine Komfort-Normalisierung, additiv.
- **Ausweis des Scan-Scopes im Report** (Zähler gescannt/ausgeschlossen), falls ein Pilot
  den stillen Ausschluss als Transparenz-Lücke meldet.

## Geschichte

| Datum | Ereignis |
|---|---|
| 2026-07-03 | Proposed — Entwurf mit [slice-023](../planning/done/slice-023-dcheck-pilot-deltas.md); Evidenz d-check-Plan-Review + m-trace-Sichtung. |
| 2026-07-03 | Proposed → Accepted (Sign-off Auftraggeber: Entscheide a–d gemäß Empfehlung inkl. fail-closed für leeren/fehlenden `tech.adapter`, [slice-023](../planning/done/slice-023-dcheck-pilot-deltas.md) §3). Ab jetzt immutable; Ablösung nur via Folge-ADR mit `Supersedes`. |
