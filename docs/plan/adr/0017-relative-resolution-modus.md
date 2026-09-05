# ADR-0017 — `relative`-Auflösungs-Modus: Import-Auflösung gegen den Ort der importierenden Datei

- **Status:** Accepted
- **Datum:** 2026-07-03
- **Autor:** pt9912
- **Bezug:** [AC-FA-CONF-001](../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml), [AC-FA-EXTRACT-001](../../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion) — **Erweiterung** von [ADR-0016](0016-resolution-sprach-parametrisch.md) (füllt dessen reservierten `mode`-Wert `relative`), nach dem Muster, mit dem ADR-0016 selbst [ADR-0014](0014-resolution-roots.md) erweiterte.
- **Schärft:** [SPEC-CONF-001](../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema) + [SPEC-RULE-001](../../../spec/spezifikation.md#spec-rule-001--regel-auswertung) + [SPEC-EXTRACT-001](../../../spec/spezifikation.md#spec-extract-001--import-extraktion).
- **Supersedes:** —

## Kontext

[ADR-0016](0016-resolution-sprach-parametrisch.md) (Accepted) lieferte die sprach-parametrische
`resolution`-Map mit `mode`-Diskriminator und **reservierte** `relative` (Exit 2) für Sprachen,
deren Importe gegen den **Ort der importierenden Datei** auflösen statt gegen eine feste Wurzel.
Der dort benannte Re-Evaluierungs-Trigger („eigener ADR, sobald ein Pilot feuert") feuert jetzt:
das **TypeScript-Backend** ([slice-022](../planning/done/welle-06/slice-022-typescript-backend.md),
welle-06). TypeScript-Module importieren einander mit relativen Specifiern (`./db`,
`../core/model`) — die gelieferten Modi können das nicht:

1. **`path`/`fixed-root` kennen nur das Symbol.** `../core/model` ist ohne den Quelldatei-Pfad
   mehrdeutig — aus `src/adapters/db.ts` meint es `src/core/model`, aus `src/app/x/y.ts` dagegen
   `src/app/core/model`.
2. **Das Signal fehlt im Threading.** [ADR-0016](0016-resolution-sprach-parametrisch.md) reichte
   die Quelldatei-**Sprache** bis `targetLayer` durch; der Quelldatei-**Pfad** liegt in
   `FileImports.Path` zwar vor (`ruleFor` hält ihn), wird aber nicht an die Auflösung
   weitergereicht.
3. **Dieselbe Signal-Klasse existiert im Bestand mehrfach:** TypeScript-`./x`-Importe, relative
   Python-Importe (`from .`, dort dokumentierte Extraktions-Grenze) und C++-`"…"`-Includes —
   alle lösen gegen den Ort des Files auf.

## Optionen

| Weg | Idee | Bewertung |
|---|---|---|
| **A — lexikalische Datei-relativ-Auflösung + Quellpfad-Threading** | `./`/`../`-Specifier werden lexikalisch gegen das Verzeichnis der importierenden Datei normalisiert; `targetLayer` bekommt zusätzlich den Quelldatei-Pfad. | **Gewählt.** Exakt (kein Raten), deterministisch, rein text-/pfad-basiert — kein Dateisystem-Zugriff im Kern; dasselbe billige Threading-Muster wie [ADR-0016](0016-resolution-sprach-parametrisch.md). |
| **B — Dateisystem-Probe (Bundler-Semantik)** | Kandidaten (`.ts`/`.tsx`/`index.ts`) gegen das Dateisystem proben, wie Node/TS-Resolver. | Verworfen: braucht FS-Zugriff im Kern (bricht die [ARC-001](../../../spec/architecture.md)-Reinheit), Node-Resolution ist ein Fass ohne Boden ([AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)); für den Glob-Präfix-Match ist die Datei-Endung irrelevant. |
| **C — relative Specifier wie `fixed-root` an der Wurzel verankern** | `./x` als wurzel-relativen Pfad behandeln. | Verworfen: semantisch falsch — `../core` aus `src/adapters/db.ts` meint `src/core`, nicht `core`; genau die Mehrdeutigkeit aus dem Kontext. |

## Entscheidung

**Weg A**, als **Erweiterung** von [ADR-0016](0016-resolution-sprach-parametrisch.md) (dessen
Sprach-Map, Sprach-Threading und `mode`-Diskriminator gelten unverändert):

1. **`mode: relative` wird gültig** (dritter Modus neben `path`/`fixed-root`); `namespace`
   bleibt als einziger Wert reserviert → Exit 2.
2. **Auflösungs-Semantik:** Ein Specifier ist **relativ**, wenn er `.` oder `..` ist oder mit
   `./` bzw. `../` beginnt (Barrel-Import `from '.'` inklusive). Nur relative Specifier werden
   aufgelöst — lexikalische Normalisierung von `dir(Quelldatei) + "/" + Specifier`
   (Slash-Pfade, `path.Clean`-Semantik). Das Ergebnis ist der wurzel-relative Kandidat für den
   bestehenden Glob-Präfix-Match. **Endungs-agnostisch:** der Match läuft über
   Layer-Glob-Präfixe auf Segmentgrenzen; ob `./db` die Datei `db.ts`, `db.js`
   (NodeNext-Schreibweise) oder das Verzeichnis `db/` meint, ändert den Ziel-Layer nicht —
   **solange** der Glob-Präfix oberhalb der Dateiebene endet (verzeichnisbasierte
   `layers`-Globs); bei datei-tiefen Globs kippt eine Specifier-Endung den Match
   (dokumentierte Grenze).
3. **Nicht-relative Specifier** (Bare-Imports wie `react`, `fs`, `@actions/core`) liefern
   unter `mode: relative` die **leere Kandidatenmenge** — das Roh-Symbol wird ausdrücklich
   **nicht** als Pfad-Kandidat weitergereicht (anders als der `path`-Default): sonst matchte
   `@actions/core` auf Segmentgrenze gegen einen `core/**`-Glob — ein Geister-Befund. Kein
   Ziel-Layer → keine schicht-basierte Regel; `tech`-Muster greifen unabhängig am Roh-Symbol.
   tsconfig-`paths`/`baseUrl`-Aliasse: Re-Evaluierungs-Trigger, kein Teil dieses ADR.
4. **Wurzel-Escape:** behält ein Specifier **nach** der Normalisierung ein führendes `..`
   (Beispiel: `../../../x` aus `src/core/`; dagegen normalisiert `../../x` von dort auf `x` —
   exakt Wurzelebene, aufgelöst), liefert er ebenfalls die leere Kandidatenmenge —
   ausgewiesene Grenze
   ([AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)),
   kein Fehler.
5. **Quellpfad-Threading:** `targetLayer` erhält zusätzlich den Pfad der importierenden Datei
   (aus `FileImports.Path` — vorhanden, bislang nicht durchgereicht); die Auflösung nutzt den
   `mode` der Import-Sprache **und** den Quellpfad.
6. **Konfigurations-Disziplin:** `mode: relative` nimmt weder `roots` noch `package_base` —
   deklariert → Exit 2 (fail-closed, kein still ignorierter Schlüssel; dieselbe Ethos-Linie wie
   [ADR-0016](0016-resolution-sprach-parametrisch.md)/[slice-017](../planning/done/wellenlos/slice-017-unbekannte-sprache-exit2.md)).

## Konsequenzen

- [ADR-0016](0016-resolution-sprach-parametrisch.md) **bleibt gültig** — ADR-0017 füllt einen
  reservierten `mode`-Wert additiv (dieselbe Map, dasselbe Threading, ein neuer Wert); kein
  Supersede. Genau das dort versprochene „kein Re-Architecting".
- **Schema** ([AC-FA-CONF-001](../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml)/[SPEC-CONF-001](../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema)):
  `mode ∈ {path, fixed-root, relative}`; nur `namespace` reserviert.
- **Erste Nutzerin:** TypeScript ([slice-022](../planning/done/welle-06/slice-022-typescript-backend.md)).
  Der Modus ist sprach-parametrisch generisch; C++-`"…"`-Includes (datei-relativ) vs.
  `<…>`-Includes (Include-Root) brauchen zusätzlich ein Import-**Kind**-Signal aus der
  Extraktion → Re-Evaluierungs-Trigger.
- `targetLayer` wächst um einen Parameter (Quellpfad; ein Aufrufer) — die Signatur-Erweiterung
  ist der ADR-würdige Kern, analog zum Sprach-Threading von
  [ADR-0016](0016-resolution-sprach-parametrisch.md).

## Fitness Function

- `make test`: `./`/`../`-Auflösung inkl. `..`-Normalisierung über Segmentgrenzen und
  Barrel-`.`; Grenz-Testpaar `../../x` (exakt Wurzel, aufgelöst) vs. `../../../x` (Escape →
  leer); **adversarisch** `@actions/core` gegen Layer-Glob `core/**` → leer, kein
  Geister-Befund (pinnt die Leere-Kandidatenmenge-Semantik gegen Roh-Durchreichung);
  Bare-Import → leer (`tech` greift am Roh-Symbol); Mono-Repo Go+TypeScript (je eigener
  `mode`); `relative` mit `roots`/`package_base` → Exit 2; `namespace` weiterhin Exit 2.
- `make arch-check` (Dogfooding, [AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)):
  unverändert 0 (a-check deklariert kein `resolution` → `path`).

## Re-Evaluierungs-Trigger

- **tsconfig-`paths`/`baseUrl`-Aliasse** (Alias→Wurzel-Map, `mode`-intern additiv): bei realem
  TypeScript-Pilot mit Alias-Layout.
- **C++-quoted-Include-Split** (`"…"` datei-relativ, `<…>` fixed-root): braucht ein
  Import-Kind aus der Extraktion; bei realem C++-Pilot mit datei-relativen Includes.
- **`namespace`-Modus** (C#): unverändert eigener Folge-ADR
  ([ADR-0016](0016-resolution-sprach-parametrisch.md) §Re-Evaluierungs-Trigger).

## Geschichte

| Datum | Ereignis |
|---|---|
| 2026-07-03 | Proposed — Entwurf mit [slice-022](../planning/done/welle-06/slice-022-typescript-backend.md); Sign-off Auftraggeber ausstehend. |
| 2026-07-03 | Entwurf nach adversarischem Review geschärft: leere Kandidatenmenge statt Roh-Durchreichung (Geister-Match `@actions/core`↔`core/**`), Barrel-`.`/`..` als relativ, Escape scharf definiert (führendes `..` nach Normalisierung), Endungs-Agnostik an verzeichnisbasierte Globs gebunden. |
| 2026-07-03 | Proposed → Accepted (Sign-off Auftraggeber: Entscheide A–G gemäß Empfehlung, [slice-022](../planning/done/welle-06/slice-022-typescript-backend.md) §7). Ab jetzt immutable; Ablösung nur via Folge-ADR mit `Supersedes`. |
