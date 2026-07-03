# slice-022 — TypeScript-Sprach-Backend + `relative`-Auflösungs-Modus (welle-06-sprach-backends)

**Status:** Entwurf zur Abnahme (2026-07-03). Entscheide §7 **vor** der Umsetzung zu treffen.
**Welle:** welle-06-sprach-backends (viertes Backend-Inkrement nach
[slice-014](../done/slice-014-java-backend.md)/[slice-020](../done/slice-020-python-backend.md)/[slice-021](../done/slice-021-csharp-backend.md))
**plus** Import-Auflösung (zweiter gelieferter Modus nach [slice-015](../done/slice-015-resolution-roots.md)).
**Bezug:** erweitert [AC-FA-EXTRACT-001](../../../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion)
um TypeScript und [AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml)
um den gültigen `relative`-Modus; **neuer Folge-ADR
[ADR-0017](../../adr/0017-relative-resolution-modus.md)** (Proposed) zu
[ADR-0016](../../adr/0016-resolution-sprach-parametrisch.md); schärft
[SPEC-EXTRACT-001](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion),
[SPEC-CONF-001](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema),
[SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung).
[Roadmap welle-06](../in-progress/roadmap.md). **Trigger:** Polyglot-Bestand (TypeScript-Repos),
Maintainer-Priorität 2; letzter offener welle-06-Kandidat.

> **Hinweis:** Entwurf zur Abnahme. Die in §4 als Code-Fence gesetzten AC-Texte sind
> unverbindlich — gültig erst nach Freigabe in [`spec/lastenheft.md`](../../../../spec/lastenheft.md).
> [ADR-0017](../../adr/0017-relative-resolution-modus.md) ist `Proposed` — Sign-off ausstehend.

---

## 1. Ziel

Ein **TypeScript**-Backend für die Import-Extraktion (ES-Module-Syntax) **plus** der bislang
reservierte **`relative`-Auflösungs-Modus**, damit TypeScript-Repos ihre Hexagon-Architektur
über a-check prüfen können. Anders als Python ([slice-020](../done/slice-020-python-backend.md))
und C# ([slice-021](../done/slice-021-csharp-backend.md)) ist das **kein reines
Extraktions-Inkrement**: TypeScript-Module referenzieren einander mit relativen Specifiern
(`./db`, `../core/model`) — ohne den `relative`-Modus wäre jedes intra-Repo-Symbol unauflösbar
und **keine einzige schicht-basierte Regel** würde greifen (geliefert-aber-nutzlos). `tech`-Regeln
greifen zusätzlich sofort am Roh-Symbol — Bare-Imports (`react`, `pg`, `express`) sind für
`tech-leak` sogar ideal sichtbar.

## 2. Problem

a-check v0.6.0 kennt {`cpp`, `go`, `rust`, `kotlin`, `java`, `python`, `csharp`};
`languages: {typescript: …}` bricht (korrekt, [slice-017](../done/slice-017-unbekannte-sprache-exit2.md))
mit Exit 2, `resolution: {…: {mode: relative}}` ebenso (reserviert,
[slice-015](../done/slice-015-resolution-roots.md)). Zwei Achsen:

**Extraktion** — TypeScript importiert über mehrere Formen mit dem Modul-Specifier als String:

- `import { Db } from '../adapters/db';` — Grundform; der Specifier steht **hinter `from`**.
- `import './polyfill';` — Seiteneffekt-Import ohne Bindung.
- `import type { T } from './ports/repo';` — Typ-Import (zur Laufzeit erased, architektonisch
  eine Kopplung — §7 Entscheid B).
- `export * from './core/model';` / `export { X } from './x';` — **Re-Export**: Barrel-Dateien
  (`index.ts`) sind der Standard-Kanal, über den Schichten leaken (§7 Entscheid C).
- `import fs = require('fs');` — TS-spezifische CommonJS-Interop-Form.
- **Kollisionsgefahr:** dynamisches `import('./lazy')` und `require('x')` sind **Ausdrücke**
  (mitten in der Zeile, in Funktionsrümpfen) — die zeilenverankerte Heuristik greift sie
  bewusst nicht (§7 Entscheid D).

**Auflösung** — der Kern von [ADR-0017](../../adr/0017-relative-resolution-modus.md):
`../core/model` ist ohne den **Pfad der importierenden Datei** mehrdeutig. Das Threading-Muster
existiert seit [ADR-0016](../../adr/0016-resolution-sprach-parametrisch.md) (Sprache), der Pfad
liegt in `FileImports.Path` bereits vor — `ruleFor` reicht ihn nur nicht an `targetLayer` durch.

## 3. Betroffene Artefakte (vor der Implementierung benannt)

- **Slice-ID:** slice-022.
- **ADR:** [ADR-0017](../../adr/0017-relative-resolution-modus.md) (neu; `Proposed` →
  `Accepted` per Sign-off; füllt den reservierten `relative`-Wert aus
  [ADR-0016](../../adr/0016-resolution-sprach-parametrisch.md) additiv).
- **AC:** [AC-FA-EXTRACT-001](../../../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion)
  (TypeScript-Backend), [AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml)
  (`relative` gültig; Grenze [AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)).
- **Spec:** [SPEC-EXTRACT-001](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion)
  (TS-Muster + Backend-Menge), [SPEC-CONF-001](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema)
  (`mode`-Menge + `relative`-Constraints), [SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung)
  (Quellpfad-Threading + relative Normalisierung).
- **Module:** `internal/adapter/driven/extract` (TS-Regexes + Registry-Eintrag),
  `internal/hexagon/core` (`resolveImport`-relative-Zweig, `targetLayer`-Signatur + Quellpfad),
  `internal/adapter/driven/config` (`relative` gültig, fail-closed-Constraints).
- **Version:** Lastenheft + Spezifikation **0.12.0 → 0.13.0**; Benutzerhandbuch-Historie.
- **Gates:** `make gates` → `make ci`; Multi-Linsen-Review (4 Linsen) + ggf. Delta →
  `docs/reviews/`.

## 4. Entwurf (zur Abnahme)

### 4.1 Anforderungs-Erweiterung — AC-FA-EXTRACT-001 (TypeScript)

```text
AC-FA-EXTRACT-001 (erweitert um TypeScript): die Backend-Liste wird um
TypeScript ergänzt (languages-Schlüssel `typescript`) — gewertet wird der
Modul-Specifier (der String hinter `from` bzw. im Import): `import … from '…'`
(inkl. `import type`), Seiteneffekt-Import `import '…'`, Re-Export
`export … from '…'` (inkl. `export * from`/`export type … from`) und die
Interop-Form `import X = require('…')`. Die links von `from` stehenden
Namen/Aliasse werden nie als Symbol geliefert. Dynamisches `import(…)` und
`require(…)` als Ausdruck werden nicht gegriffen (zeilenverankerte Heuristik).

Neue/ergaenzte Akzeptanzkriterien:
- Happy (TS): Given `import { Db } from '../adapters/db';`, when das
  TypeScript-Backend laeuft, then liefert es das Symbol `../adapters/db`.
- Boundary (TS type/Seiteneffekt): Given `import type { Repo } from
  './ports/repo';` und `import './polyfill';`, when das Backend laeuft,
  then liefert es `./ports/repo` bzw. `./polyfill`.
- Boundary (TS Re-Export/require): Given `export * from './core/model';`
  und `import fs = require('fs');`, when das Backend laeuft, then liefert
  es `./core/model` bzw. `fs`.
- Negative (TS Ausdruck): Given `const m = await import('./lazy');` oder
  `const x = require('pg');` (Ausdrucks-Position), when das Backend laeuft,
  then wird KEIN Symbol geliefert (dokumentierte Heuristik-Grenze, AC-QA-02).

Out-of-Scope: dynamisches import()/require() im Ausdruck; import-aehnliche
Zeichenfolgen in Template-Literalen (Backticks — bestehende String-Grenze,
AC-QA-02); JavaScript (.js/.mjs/.cjs) als eigener languages-Schluessel;
tsconfig paths/baseUrl-Aliasse (Re-Evaluierungs-Trigger ADR-0017);
Node-Modul-Aufloesung (Endungen/index-Dateien — der Glob-Praefix-Match ist
endungs-agnostisch, eine Datei-Existenz-Probe findet nicht statt).
```

### 4.2 Anforderungs-Erweiterung — AC-FA-CONF-001 (`relative`-Modus)

```text
AC-FA-CONF-001 (erweitert): resolution.mode ∈ {path (Default), fixed-root,
relative}; nur `namespace` bleibt reserviert (Exit 2). `mode: relative` loest
Specifier mit fuehrendem `./`/`../` lexikalisch gegen das Verzeichnis der
importierenden Datei auf; nicht-relative Specifier bleiben unaufgeloest.
`mode: relative` nimmt weder `roots` noch `package_base` (deklariert → Exit 2).

Neue/ergaenzte Akzeptanzkriterien:
- Happy (relative): Given `resolution: {typescript: {mode: relative}}` und in
  `src/core/service.ts` der Import `../adapters/db`, when a-check laeuft,
  then loest das Symbol auf `src/adapters/db` auf (Schicht der adapters-Globs)
  — eine Domaenen-Datei mit diesem Import wird core-impurity.
- Boundary (Wurzel-Escape/Bare-Import): Given ein Specifier, der ueber die
  Scan-Wurzel hinaus normalisiert (z. B. `../../x` nahe der Wurzel), oder ein
  Bare-Import (`react`), when a-check laeuft, then bleibt das Symbol
  unaufgeloest — kein Ziel-Layer, keine schicht-basierte Regel; tech-Muster
  greifen am Roh-Symbol (ausgewiesene Grenze, AC-QA-02).
- Negative (relative+roots): Given `{mode: relative, roots: ["src"]}` oder
  `{mode: relative, package_base: "x"}`, when a-check laedt, then Exit 2.
```

### 4.3 Beispiel-Rezept (Benutzerhandbuch)

```yaml
# TypeScript-Hexagon: src/{core,ports,adapters}/…, relative Importe
languages:
  typescript: ["**/*.ts", "**/*.tsx"]
layers:
  core:     ["src/core/**"]
  ports:    ["src/ports/**"]
  adapters: ["src/adapters/**"]
edges:
  - {from: adapters, to: ports}
  - {from: ports,    to: core}
resolution:
  typescript: {mode: relative}
```

### 4.4 Versions-Bump + Sweep

Lastenheft + Spezifikation **0.12.0 → 0.13.0** (zwei AC-Erweiterungen, eine Historie-Zeile).
„sieben → acht Sprachen"-Sweep (Lerneintrag [slice-020](../done/slice-020-python-backend.md):
zählende Stellen explizit benennen): `README.md`, Benutzerhandbuch (§Sprachen + Rezept +
Historie), `spec/architecture.md` ([ARC-003](../../../../spec/architecture.md)-Sprachliste),
`harness/README.md` §Safety-Sprachliste. **Nicht**
[ADR-0002](../../adr/0002-text-heuristische-extraktion.md)/[ADR-0014](../../adr/0014-resolution-roots.md)/[ADR-0016](../../adr/0016-resolution-sprach-parametrisch.md)
(`Accepted` ⇒ immutable).

## 5. Umsetzungsplan (Reihenfolge: ADR → Lastenheft → Spec → Code → Tests)

0. **[ADR-0017](../../adr/0017-relative-resolution-modus.md)** `Proposed → Accepted`
   (Sign-off = Abnahme §7); [ADR-Index](../../adr/README.md) ist ergänzt.
1. **Lastenheft:** §4.1 + §4.2 einarbeiten, Bump **0.13.0** + Historie-Zeile.
2. **Spezifikation:** [SPEC-EXTRACT-001](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion)
   (TS-Muster + Backend-Menge `{cpp, go, rust, kotlin, java, python, csharp, typescript}`),
   [SPEC-CONF-001](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema)
   (`mode`-Menge `{path, fixed-root, relative}`, nur `namespace` reserviert;
   `relative`-Constraints), [SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung)
   (Normalisierung nutzt Quelldatei-Sprache **und** -Pfad; relative Semantik + Wurzel-Escape-Grenze);
   Bump + Historie.
3. **`extract.go`:** Felder `tsFrom` (`import`/`export` … `from '…'`, `type` optional),
   `tsSide` (Seiteneffekt `import '…'`), `tsRequire` (`import X = require('…')`) in
   `newAdapter`; Registry-Eintrag `"typescript"` via `lineMatches`. TypeScript ist C-Syntax →
   `prepSource` lässt das `//`-/`/* */`-Stripping **an** (Lerneintrag
   [slice-020](../done/slice-020-python-backend.md): geteilte Vorverarbeitung mit-reviewen;
   Backtick-Template-Literale = bestehende String-Grenze der C-Familie).
4. **`core`:** `resolveImport` bekommt den Quellpfad und den `relative`-Zweig (lexikalisches
   `path.Clean` über `dir(Quelldatei) + "/" + Specifier`; nur `./`/`../`-Specifier; Escape →
   unaufgelöst); `targetLayer`-Signatur + Durchreichung aus `ruleFor` (`f.Path` liegt dort vor).
5. **`config`:** `relative` von reserviert → gültig; Validierung `roots`/`package_base` bei
   `relative` → Exit 2; `namespace` bleibt reserviert; Fehlermeldungs-Nachzug.
6. **Tests** (Mutanten-Boundary nach [slice-021](../done/slice-021-csharp-backend.md)-Lerneintrag —
   Negativ-Zeilen so wählen, dass das Fehlsymbol auf eine Schicht auflösen *würde*):
   Extraktion (alle §4.1-Formen; `// import …` gestrippt; `import`/`export` mitten in der Zeile
   bzw. Ausdrucks-`import()`/`require()` → kein Match; `importX`/`exportX`-Keyword-Präfix);
   Auflösungs-Unit (`./`-Nachbar, `../`-Eltern über Segmentgrenzen, Mehrfach-`..`,
   Wurzel-Escape, Bare-Import); Config-Fails (`relative`+`roots`/`package_base`,
   `namespace` weiter reserviert); **CLI-Integration:** TS-Fixture — Domäne importiert
   `../adapters/db` ⇒ `core-impurity`, Exit 1; **Mono-Repo Go+TypeScript** (je eigener Modus).
7. **Fixture-/Meldungs-Nachzug:** gepinnte Backend-Menge wird
   `cpp|csharp|go|java|kotlin|python|rust|typescript` (`TestCheckLanguagesUnknown` u. a.);
   `TestResolutionReservedModeFailsClosed` nutzt heute `typescript: {mode: relative}` als
   Reserviert-Beispiel → auf `namespace` umstellen (Analogon zur `python`→`ruby`-Umstellung
   in [slice-020](../done/slice-020-python-backend.md)).
8. **Sweep §4.4** inkl. Benutzerhandbuch-Rezept (§4.3).
9. `make gates`/`make ci`; Multi-Linsen-Review (4 Linsen) + ggf. Delta → `docs/reviews/`;
   Verifikation; Closure (reiner `git mv`, 2 Kriterien, Lerneintrag).

## 6. Definition of Done

- [ ] [ADR-0017](../../adr/0017-relative-resolution-modus.md) **Accepted** (Sign-off) +
      [ADR-Index](../../adr/README.md).
- [ ] Lastenheft + Spezifikation **0.13.0**: TS-Backend + `relative`-Modus (§4.1/§4.2) mit
      Historie-Zeilen; [SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung)-Threading
      um den Quellpfad ergänzt.
- [ ] `extract.go`: TS-Regexes + Registry-Eintrag `"typescript"`; Dogfooding grün (0 Befunde).
- [ ] `core`/`config`: `relative`-Auflösung + Quellpfad-Threading; fail-closed-Constraints;
      Default (`path`/ohne `resolution`) byte-identisch unverändert.
- [ ] Tests: §5.6-Umfang inkl. CLI-Integration + Mono-Repo; Fixture-/Meldungs-Nachzug §5.7.
- [ ] „sieben → acht Sprachen"-Sweep vollständig (§4.4).
- [ ] `make gates` + `make ci` grün; Multi-Linsen-Review (4 Linsen) + ggf. Delta bestanden.
- [ ] **Maintainer-Abnahme der Entscheide A–F (§7).**
- [ ] Closure: reiner `git mv` nach `done/` (AGENTS §3.3); 2 beobachtbare Kriterien + Lerneintrag.

## 7. Offen / Entscheidungen zur Abnahme

- **Entscheid A — Sprach-Schlüssel `typescript`:** (nicht `ts`) — konsistent mit `csharp`
  ([slice-021](../done/slice-021-csharp-backend.md) Entscheid E); `.tsx` läuft über die
  Datei-Globs der Config. *Empfehlung: `typescript`.*
- **Entscheid B — `import type` wird gewertet:** Typ-Importe sind zur Laufzeit erased, aber
  architektonisch eine Kopplung — eine Domäne, die Adapter-**Typen** importiert, ist genau der
  Befund, den ein TS-Hexagon braucht. *Empfehlung: werten (Präzedenz: Java-`static`/C#-`global`
  werden übersprungen, das Symbol zählt).*
- **Entscheid C — Re-Exports (`export … from`) werden gewertet:** Barrel-Dateien (`index.ts`)
  sind der Standard-Leak-Kanal. *Empfehlung: werten — ein Re-Export ist eine echte
  Abhängigkeits-Kante.*
- **Entscheid D — dynamisches `import()`/`require()` out of scope:** Ausdrucks-Position,
  zeilenverankerte Heuristik greift nicht (dokumentierte Grenze
  [AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)).
  *Empfehlung: bestätigen — kein stiller Anspruch auf Vollständigkeit.*
- **Entscheid E — nicht-relative Specifier bleiben unaufgelöst; kein `roots`-Fallback:**
  Bare-Imports sind externe Pakete (`tech` greift); tsconfig-`paths`/`baseUrl`-Aliasse sind ein
  **additives** Folge-Inkrement ([ADR-0017](../../adr/0017-relative-resolution-modus.md)
  Re-Evaluierungs-Trigger), getriggert durch einen realen TS-Pilot mit Alias-Layout.
  *Empfehlung: bestätigen — Gating-Doktrin der Welle.*
- **Entscheid F — `relative` nimmt weder `roots` noch `package_base` (Exit 2):** fail-closed
  statt still ignorierter Schlüssel. *Empfehlung: bestätigen — Ethos-Linie
  [slice-015](../done/slice-015-resolution-roots.md)/[slice-017](../done/slice-017-unbekannte-sprache-exit2.md).*
- **Risiko/Notiz — geteilte Vorverarbeitung:** TypeScript ist C-Syntax → C-Strip bleibt an
  (`// import …` muss neutralisiert werden); eine `/*`-Bytefolge in einem Template-Literal
  teilt die **bestehende**, seit 0.1.0 dokumentierte String-Grenze der C-Familie
  ([AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)) —
  kein TS-Spezifikum.
- **Risiko/Notiz — synthetische Verifikation:** noch kein benannter TypeScript-Pilot;
  Verifikation gegen eigene Fixtures (wie Java/Python/C#). Sprache bleibt gated-geliefert.
