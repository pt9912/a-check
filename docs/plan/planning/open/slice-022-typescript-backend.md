# slice-022 — TypeScript-Sprach-Backend + `relative`-Auflösungs-Modus (welle-06-sprach-backends)

**Status:** in Umsetzung (2026-07-03). **Abnahme erteilt:** Entscheide A–G gemäß Empfehlung
(Maintainer-Wort 2026-07-03, nach adversarischem Review des Entwurfs);
[ADR-0017](../../adr/0017-relative-resolution-modus.md) damit `Accepted`.
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
  `internal/adapter/driven/config` (`relative` gültig, fail-closed-Constraints),
  `internal/cli` (Bestands-Test-Nachzug §5.7 + Integrationstests — `typescript` ist dort
  heute Unsupported-Fixture).
- **Version:** Lastenheft + Spezifikation **0.12.0 → 0.13.0**; Benutzerhandbuch-Historie.
- **Gates:** `make gates` → `make ci`; Multi-Linsen-Review (4 Linsen) + ggf. Delta →
  `docs/reviews/`.

## 4. Entwurf (zur Abnahme)

### 4.1 Anforderungs-Erweiterung — AC-FA-EXTRACT-001 (TypeScript)

```text
AC-FA-EXTRACT-001 (erweitert um TypeScript): die Backend-Liste wird um
TypeScript ergänzt (languages-Schlüssel `typescript`) — gewertet wird der
Modul-Specifier: der String in einfachen ODER doppelten Anführungszeichen
(beide normativ gleichwertig; Backticks nie — Template-Literal-Grenze)
hinter `from` bzw. im Import; das Semikolon ist optional (ASI) und NICHT
Teil des Musters (anders als beim C#-Pflicht-`;`, das dort ein Feature ist).
Formen: `import … from '…'` (inkl. `import type` und Inline-type-Modifier),
Seiteneffekt-Import `import '…'`, Re-Export `export … from '…'` (inkl.
`export * from`, `export * as ns from`, `export type … from`), die
Interop-Form `import X = require('…')` sowie die Fortsetzungszeile
`} from '…'` eines mehrzeilig umbrochenen Imports/Re-Exports (Prettier
bricht named-Import-Listen ab printWidth 80 routinemäßig um — gerade bei
den langen Adapter-Specifiern, die der Check fangen soll). Der Mittelteil
zwischen import/export und from ist auf Import-Clause-Zeichen beschränkt
(Bezeichner, `{ } * ,`, `type`/`as`, Whitespace — kein `=`, `(`, `.`,
keine Quotes), damit Ausdrucks-Zeilen wie `export const q =
knex.from('users')` nie matchen. Die links von `from` stehenden
Namen/Aliasse werden nie als Symbol geliefert. Dynamisches `import(…)` und
`require(…)` als Ausdruck werden nicht gegriffen (zeilenverankerte
Heuristik).

Neue/ergaenzte Akzeptanzkriterien:
- Happy (TS): Given `import { Db } from '../adapters/db';` sowie
  `import { Repo } from "../adapters/db.js"` (Double-Quotes, semikolonfrei,
  .js-Specifier auf .ts-Datei — NodeNext-Pflicht-Schreibweise), when das
  TypeScript-Backend laeuft, then liefert es `../adapters/db` bzw.
  `../adapters/db.js`.
- Boundary (TS type/Seiteneffekt): Given `import type { Repo } from
  './ports/repo';` und `import './polyfill';`, when das Backend laeuft,
  then liefert es `./ports/repo` bzw. `./polyfill`.
- Boundary (TS Re-Export/require/mehrzeilig): Given `export * from
  './core/model';`, `import fs = require('fs');` und ein mehrzeilig
  umbrochener Import, dessen Schlusszeile `} from '../adapters/db';`
  lautet, when das Backend laeuft, then liefert es `./core/model`, `fs`
  bzw. `../adapters/db`.
- Negative (TS Ausdruck): Given `const m = await import('./lazy');`,
  `const x = require('pg');`, `export const q = knex.from('users');` oder
  eine zeilen-anfuehrende Ausdrucks-Zeile `import('./x').then(m => …)`,
  when das Backend laeuft, then wird KEIN Symbol geliefert (dokumentierte
  Heuristik-Grenze, AC-QA-02).

Out-of-Scope: dynamisches import()/require() im Ausdruck; import-aehnliche
Zeichenfolgen in Template-Literalen (Backticks) und in JSX-Textzeilen von
.tsx-Dateien (bestehende String-Grenze, AC-QA-02); Triple-Slash-Direktiven
(`/// <reference path="…" />` — fallen dem C-Kommentar-Strip zu);
JavaScript (.js/.mjs/.cjs) als eigener languages-Schluessel; tsconfig
paths/baseUrl-Aliasse (Re-Evaluierungs-Trigger ADR-0017);
Node-Modul-Aufloesung (Endungen/index-Dateien — keine Datei-Existenz-Probe;
der Glob-Praefix-Match ist endungs-agnostisch, solange die layers-Globs
verzeichnisbasiert sind, siehe AC-FA-CONF-001).
```

### 4.2 Anforderungs-Erweiterung — AC-FA-CONF-001 (`relative`-Modus)

```text
AC-FA-CONF-001 (erweitert): resolution.mode ∈ {path (Default), fixed-root,
relative}; nur `namespace` bleibt reserviert (Exit 2). Ein Specifier ist
RELATIV, wenn er `.` oder `..` ist oder mit `./` bzw. `../` beginnt
(Barrel-Import `from '.'` inklusive); er wird lexikalisch gegen das
Verzeichnis der importierenden Datei normalisiert (path.Clean-Semantik).
Alle anderen Specifier (Bare-Imports) und Specifier mit fuehrendem `..`
NACH der Normalisierung (Wurzel-Escape) liefern eine LEERE Kandidatenmenge
— das Roh-Symbol wird ausdruecklich NICHT als Pfad-Kandidat weitergereicht
(anders als der path-Default; sonst matchte z. B. `@actions/core` auf
Segmentgrenze gegen einen `core/**`-Glob — Geister-Befund). Die
Endungs-Agnostik der Aufloesung gilt, solange der Layer-Glob-Praefix
oberhalb der Dateiebene endet (verzeichnisbasierte Globs wie
`src/adapters/**`); bei datei-tiefen Globs (`src/adapters/db/**`) kippt
eine Specifier-Endung (`db.js`) den Match — dokumentierte Grenze.
`mode: relative` nimmt weder `roots` noch `package_base` (deklariert →
Exit 2).

Neue/ergaenzte Akzeptanzkriterien:
- Happy (relative): Given `resolution: {typescript: {mode: relative}}` und in
  `src/core/service.ts` der Import `../adapters/db`, when a-check laeuft,
  then loest das Symbol auf `src/adapters/db` auf (Schicht der adapters-Globs)
  — eine Domaenen-Datei mit diesem Import wird core-impurity. Grenzfall
  inklusive: `../../x` aus `src/core/` normalisiert auf `x` (exakt
  Wurzelebene, aufgeloest).
- Boundary (Escape/Bare-Import, adversarisch): Given `../../../x` aus
  `src/core/` (fuehrendes `..` nach der Normalisierung) oder
  `import * as core from '@actions/core'` bei einem Layer-Glob `core/**`,
  when a-check laeuft, then bleibt das Symbol unaufgeloest — leere
  Kandidatenmenge, KEIN Ziel-Layer, kein Geister-Befund; tech-Muster greifen
  unabhaengig am Roh-Symbol (ausgewiesene Grenze, AC-QA-02).
- Negative (relative+roots): Given `{mode: relative, roots: ["src"]}` oder
  `{mode: relative, package_base: "x"}`, when a-check laedt, then Exit 2.
```

### 4.3 Beispiel-Rezept (Benutzerhandbuch)

```yaml
# TypeScript-Hexagon: src/{core,ports,adapters}/…, relative Importe
languages:
  typescript: ["**/*.ts", "**/*.tsx", "**/*.mts", "**/*.cts"]
layers:                       # layers-Globs verzeichnisbasiert halten!
  core:     ["src/core/**"]   # (kein "src/core/**/*.ts" — ein Glob mit
  ports:    ["src/ports/**"]  #  Datei-Endung macht die Symbol-Aufloesung
  adapters: ["src/adapters/**"] # still blind und kippt die Endungs-Agnostik)
edges:
  - {from: adapters, to: ports}
  - {from: ports,    to: core}
resolution:
  typescript: {mode: relative}
```

Der Warnsatz „layers-Globs verzeichnisbasiert halten" geht auch ins
Benutzerhandbuch-Rezept (ein Glob mit Datei-Endung lässt den Wildcard im
Glob-Präfix stehen → `targetLayer` matcht nie — stiller Blindflug). Hinweis
dort ebenfalls: `.cts`-Dateien importieren typischerweise per
`require()`-Ausdruck und fallen damit flächig unter die dokumentierte
Ausdrucks-Grenze (§4.1 Out-of-Scope).

### 4.4 Versions-Bump + Sweep

Lastenheft + Spezifikation **0.12.0 → 0.13.0** — **zwei Historie-Zeilen** unter dem einen
Bump (eine je erweiterter AC: [AC-FA-EXTRACT-001](../../../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion)
TypeScript-Backend, [AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml)
`relative`-Modus), Kadenz „eine Zeile pro AC-Änderung" bleibt gewahrt.
„sieben → acht Sprachen"-Sweep (Lerneintrag [slice-020](../done/slice-020-python-backend.md):
zählende Stellen explizit benennen): `README.md`, Benutzerhandbuch (§Sprachen + Rezept +
Historie), `spec/architecture.md` ([ARC-003](../../../../spec/architecture.md)-Sprachliste),
`harness/README.md` §Safety-Sprachliste, `CHANGELOG.md` (Unreleased-Eintrag).
**Footgun:** in denselben Sweep-Dateien steht „sieben" auch für die **Regeln**
(`README.md`, `spec/architecture.md` — „dieselben sieben Regeln" in der
[ARC-001](../../../../spec/architecture.md)-Zeile) — diese Stellen bleiben unangetastet; ebenso pinnen historische `CHANGELOG.md`-Einträge alte
Sprachmengen (nicht anfassen). **Nicht**
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
3. **`extract.go`:** **vier** Felder in `newAdapter`, Registry-Eintrag `"typescript"` via
   `lineMatches`: `tsFrom` (`import`/`export` … `from`, `type` optional, **Mittelteil auf
   Import-Clause-Zeichen beschränkt** — kein `=`/`(`/`.`/Quotes, §4.1), `tsSide`
   (Seiteneffekt `import '…'`), `tsRequire` (`import X = require('…')`) und `tsCont`
   (**Fortsetzungszeile** `^\s*\}\s*from\s*['"]…` mehrzeilig umbrochener Imports —
   kollisionsarm, Ausdrucks-Aufrufe wie `db.from('x')` haben nie ein zeilen-anführendes
   `}`). Beide Quote-Arten in allen vier Mustern; Semikolon optional. TypeScript ist
   C-Syntax → `prepSource` lässt das `//`-/`/* */`-Stripping **an** (Lerneintrag
   [slice-020](../done/slice-020-python-backend.md): geteilte Vorverarbeitung mit-reviewen;
   Backtick-Template-Literale/JSX = bestehende String-Grenze der C-Familie).
4. **`core`:** `resolveImport` bekommt den Quellpfad und den `relative`-Zweig: Specifier
   relativ ⇔ `.`/`..`/`./…`/`../…`; lexikalisches `path.Clean` über
   `dir(Quelldatei) + "/" + Specifier`; Bare-Imports und Wurzel-Escape (führendes `..`
   **nach** der Normalisierung) liefern die **leere Kandidatenmenge** — ausdrücklich
   **kein** `[]string{imp}`-Durchreichen wie im `path`-Default (Geister-Match-Gefahr,
   §4.2); `targetLayer`-Signatur + Durchreichung aus `ruleFor` (`f.Path` liegt dort vor).
5. **`config`:** `relative` von reserviert → gültig; Validierung `roots`/`package_base` bei
   `relative` → Exit 2; `namespace` bleibt reserviert; Fehlermeldungs-Nachzug **inkl. der
   `default:`-Enum** `mode %q ungültig (path|fixed-root)` → `(path|fixed-root|relative)`
   (`resolutionEntry`, `config.go`).
6. **Tests** (Mutanten-Boundary nach [slice-021](../done/slice-021-csharp-backend.md)-Lerneintrag —
   Negativ-Zeilen so wählen, dass das Fehlsymbol auf eine Schicht auflösen *würde*):
   - **Extraktion, positiv:** Grundform Single-/Double-Quote (± Semikolon); Default-,
     Namespace- (`* as ns`), gemischte (`A, { B }`) Import-Clauses; `import type` +
     Inline-`type`-Modifier; Seiteneffekt; `export * from`/`export { X } from`/
     `export * as ns from`/`export type … from`; `import X = require('…')`;
     **Fortsetzungszeile** `} from '../adapters/db';`; `.js`-Specifier.
   - **Extraktion, negativ:** `// import …`/`/* … */` gestrippt; Triple-Slash-Direktive;
     `const m = await import('./lazy')`; zeilen-anführendes `import('./x').then(…)`
     (schärfster Mutant gegen `tsSide`); `const x = require('pg')`;
     `export const q = knex.from('users')` (Mittelteil-Beschränkung); `export {}` ohne
     `from`; `declare module '…' {`; `importX`/`exportX`-Keyword-Präfix; nacktes `}`
     ohne `from`.
   - **Auflösungs-Unit:** `./`-Nachbar; `../`-Eltern über Segmentgrenzen; Mehrfach-`..`;
     Barrel `.`/`..`; **Grenz-Testpaar** `../../x` aus `src/core/` (= exakt Wurzel,
     aufgelöst) vs. `../../../x` (Escape → leer); **adversarisch** `@actions/core` bzw.
     `@x/core` gegen Layer-Glob `core/**` → leer, kein Geister-Befund (pinnt die
     Leere-Kandidatenmenge-Semantik); `.js`-Endung gegen verzeichnisbasierten Glob.
   - **Config-Fails:** `relative`+`roots`, `relative`+`package_base`, `namespace` weiter
     reserviert; Fehlermeldungs-Enum gepinnt.
   - **CLI-Integration:** TS-Fixture — Domäne importiert `../adapters/db` ⇒
     `core-impurity`, Exit 1 (Negativ-Zeile im Fixture so, dass ihr Fehlsymbol auflösen
     *würde*); **Mono-Repo Go+TypeScript** (je eigener Modus).
7. **Bestands-Test-/Fixture-/Meldungs-Nachzug** (Analogon zur `python`→`ruby`-Umstellung in
   [slice-020](../done/slice-020-python-backend.md)):
   - gepinnte Backend-Menge → `cpp|csharp|go|java|kotlin|python|rust|typescript`:
     `TestCheckLanguagesUnknown` (Pipe-Format) **und** `TestBackendRegistrySet`
     (Komma-Liste, `extract_test.go`);
   - `TestCheckLanguagesMixedUnsupported` (`extract_test.go`) nutzt `typescript` als
     Nach-`go`-Unsupported-Fixture → auf `ruby` umbasen (der Kommentar dort dokumentiert
     exakt diesen Präzedenzfall aus [slice-021](../done/slice-021-csharp-backend.md));
   - `TestMonoRepoMixedUnsupportedExit2` (`internal/cli/cli_test.go`) erwartet für
     `go`+`typescript` Exit 2 → Unsupported-Fixture ersetzen (z. B. `ruby`);
   - `TestResolutionReservedModeFailsClosed` (`config_test.go`) nutzt
     `typescript: {mode: relative}` als Reserviert-Beispiel → auf `namespace` umstellen.
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
- [ ] **Maintainer-Abnahme der Entscheide A–G (§7).**
- [ ] Closure: reiner `git mv` nach `done/` (AGENTS §3.3); 2 beobachtbare Kriterien + Lerneintrag.

## 7. Offen / Entscheidungen zur Abnahme

> **Abnahme (2026-07-03):** Entscheide A–G gemäß Empfehlung bestätigt (Maintainer-Wort),
> nach Einarbeitung des adversarischen Entwurfs-Reviews (1 BLOCKER → Entscheid G, 6 MAJOR,
> 8 MINOR, 5 NIT).

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
- **Entscheid G — Mehrzeilen-Imports per Fortsetzungs-Regex (`tsCont`):** Prettier-umbrochene
  Import-Listen (Schlusszeile `} from '…'`) werden gegriffen — Weg (a) des Reviews — statt
  sie als Out-of-Scope-Grenze zu deklarieren (Weg (b)). Kollisionsarm: Ausdrucks-Aufrufe wie
  `db.from('x')` haben nie ein zeilen-anführendes `}`. *Empfehlung: (a) — eine Regex, größter
  realer Nutzen; ohne sie wäre genau der lange Adapter-Specifier, den der Check fangen soll,
  ein stilles Falsch-Grün der core-impurity.*
- **Risiko/Notiz — geteilte Vorverarbeitung:** TypeScript ist C-Syntax → C-Strip bleibt an
  (`// import …` muss neutralisiert werden); eine `/*`-Bytefolge in einem Template-Literal
  teilt die **bestehende**, seit 0.1.0 dokumentierte String-Grenze der C-Familie
  ([AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)) —
  kein TS-Spezifikum; dieselbe Klasse gilt für JSX-Textzeilen in `.tsx`, die mit
  `import … from` beginnen (String-/Markup-Grenze, nicht TS-spezifisch behandelt).
- **Risiko/Notiz — synthetische Verifikation:** noch kein benannter TypeScript-Pilot;
  Verifikation gegen eigene Fixtures (wie Java/Python/C#). Sprache bleibt gated-geliefert.
