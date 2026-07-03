# slice-023 — d-check-Pilot-Deltas: `tech.adapter`-Liste, `composition_root` steuerbar, `exclude`-Globs

**Status:** open (konsumenten-gated — das Schwester-Repo **d-check** wartet: dessen
`arch-check`-Ablösung durch a-check ist dort per Plan-Review auf diese drei Deltas
als **Umstellungs-Vorbedingung** gestellt; erst ein lieferndes a-check-Release +
Pin-Hebung entsperrt den dortigen Umbau).
**Bezug:** Change Request
[AC-FA-RULE-003](../../../../spec/lastenheft.md#ac-fa-rule-003--tech-kapselung-regel-tech-leak)
+ [AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml)
(**Lastenheft 0.14.0**, CR-Zeilen 1/3–3/3 in der Historie); schärft bei Umsetzung
[SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung)
und den Config-Teil der Spezifikation. [Roadmap](../in-progress/roadmap.md)
(welle-11-dcheck-pilot-deltas). **Autor:** pt9912. **Datum:** 2026-07-03.

## 1. Auslöser

d-check will sein `make arch-check`-Gate (Import-Regeln R1–R6 des Hexagon-Schnitts,
heute ein `go list`-Shell-Skript) durch a-check-Konsum ersetzen — der erste
**Go-Pilot** und die symmetrische Schwester-Beziehung (a-check konsumiert d-check
bereits via `d-check.mk`). Der dortige unabhängige Plan-Review (2026-07-03, Repo
d-check: slice-058, dortige ADR 0029) hat drei quellen-belegte Deltas
identifiziert (gegen v0.6.0 erhoben, am v0.7.0-Stand re-verifiziert), die die
Übersetzung der d-check-Regeln in eine `.a-check.yml` blockieren:

1. **Ein-Pattern-ein-Adapter:** `tech.adapter` ist ein einzelner Pfad-Substring
   (Erst-Treffer-Präzedenz) — d-checks R3 erlaubt `yaml` aber in **zwei** Adaptern
   (Config **und** Report). Nicht ausdrückbar, ohne die Kapsel breit zu lockern.
2. **`*_test.go` im Scan-Scope:** der Scanner erfasst jede `languages`-Glob-Datei,
   die Glob-Engine kennt **keine Negation** — Test-Dateien sind nicht
   ausschließbar. d-checks abgelöstes Gate prüfte via `go list` nur
   Nicht-Test-Imports; schon ein `os`-Import in einem Adapter-Test macht den
   sauberen d-check-Baum rot.
3. **Composition-Root-Total-Ausnahme:** `composition_root` ist von **allen**
   Prüfungen ausgenommen, auch von `tech-leak`
   ([SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung))
   — d-check verbietet `net/http`/`yaml` heute aber auch in CLI/`cmd`;
   die Deckung ginge dort still verloren.

## 2. Geplanter Umfang

Die drei CRs sind im Lastenheft **0.14.0** bereits bindend erfasst
([AC-FA-RULE-003](../../../../spec/lastenheft.md#ac-fa-rule-003--tech-kapselung-regel-tech-leak)/[AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml),
Historie 1/3–3/3); dieser Slice liefert die Umsetzung:

1. **`tech.adapter` als Pfad-Liste** (Skalar bleibt gültig; leere Liste → Exit 2;
   Symbol in jedem gelisteten Adapter erlaubt). Alle drei bestehenden
   Semantik-Zusagen halten: Substring/Regex je `match`, Erst-Treffer-Präzedenz,
   Rückwärtskompatibilität byte-identisch.
2. **`tech.composition_root: allow|forbid`** (Default `allow`): `forbid` schaltet
   **nur** die `tech-leak`-Ausnahme der Composition Root für diesen Eintrag ab;
   die Schicht-Regel-Ausnahme (Verdrahtungspunkt darf alle Schichten importieren)
   bleibt unangetastet.
3. **`exclude`-Block** (Top-Level, Datei-Globs relativ zur Scan-Wurzel): matchende
   Dateien fallen **vor** der Extraktion aus dem Scan; ungültiger Glob → Exit 2;
   ohne Block byte-identisch.
4. **Spezifikation** nachziehen ([SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung)
   `tech-leak`-Zeile + Composition-Root-Absatz, Config-Schema, Versions-Gleichlauf
   mit dem Lastenheft) + Benutzerhandbuch; ADR-Bedarf je Delta bei Umsetzung
   entscheiden (Kandidat: `exclude` als Scan-Scope-Entscheidung).
5. **Tests** entlang der neuen Akzeptanzkriterien (Mehr-Adapter,
   Composition-Root-Verbot, `exclude`-Happy/Boundary/Negative, Exit-2-Guards) +
   Dogfooding (`make arch-check` bleibt grün, byte-identisch ohne neue Keys).
6. **Release + Digest-Re-Pin** (`a-check.mk`/`cli.go`-Selbst-Pin), damit d-check
   den neuen Stand digest-gepinnt konsumieren kann — das entsperrt drüben den
   Umbau.

## 3. Vor der Umsetzung zu klären

> **Abnahme (2026-07-03):** Entscheide a–d gemäß Empfehlung bestätigt (Maintainer-Wort),
> inkl. des Review-Entscheids fail-closed für leeren/fehlenden `tech.adapter`;
> [ADR-0018](../../adr/0018-exclude-scan-scope.md) damit `Accepted`.
>
> **Entscheide bei Umsetzung (2026-07-03, abgenommen):** (a) `exclude`-Ort =
> **Top-Level**, per [ADR-0018](../../adr/0018-exclude-scan-scope.md) (Proposed —
> Sign-off ausstehend; m-trace-Evidenz unten). (b) Reihenfolge: `exclude` wirkt als
> **Scan-Scope vor der Extraktion** (Datei existiert für keine Prüfung),
> `composition_root` erst in der Regel-Auswertung — in
> [SPEC-EXTRACT-001](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion)/[SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung)
> dokumentiert; einen „gescannte Dateien"-Zähler gibt der Report heute nicht aus
> (Ausweis des Scan-Scopes = [ADR-0018](../../adr/0018-exclude-scan-scope.md)
> Re-Eval-Trigger). (c) `tech-leak`-Meldetext nennt **alle** gelisteten Adapter in
> Deklarationsreihenfolge (`config|report`); Composition-Root-Verbots-Befunde tragen
> den Zusatz `(composition_root: forbid)`. (d) Strengegefälle (nach Review-R1
> präzisiert): der **nicht-leere Skalar** bleibt byte-identisch; ein **leerer oder
> fehlender** `adapter` bricht mit Exit 2 — die Alt-Semantik war dort ein **stiller
> Never-Leak** (`strings.Contains(pfad, "")` ist immer wahr, das Muster meldete
> nie), also ein Falsch-Grün-Eintrag, den fail-closed ersetzt statt konserviert
> (Ethos-Linie [slice-017](../done/slice-017-unbekannte-sprache-exit2.md)); die **Liste** ist durchgängig
> strikt (leere Liste/leerer Eintrag → Exit 2). YAML-Aliasse auf `adapter` werden
> dereferenziert.

- **`exclude`-Ort im Schema:** Top-Level neben `languages` (ein Scan hat eine
  Ausschluss-Liste) oder je Sprache? Vorschlag: Top-Level — die d-check-Evidenz
  (`**/*_test.go`) ist sprachgebunden nur per Konvention, und ein globaler
  Ausschluss (generierter Code) ist der allgemeinere Fall. Entscheidung ggf. per
  ADR festhalten.
  **m-trace-Evidenz (2026-07-03, Sichtung des Schwester-Repos):** pnpm-TS-Monorepo
  **mit Go-Hexagon-App** (`apps/api`: `hexagon/{domain,application,port}` +
  `adapters/{driving,driven}`, 99 `*_test.go`) neben TS-Workspaces
  (`packages/player-sdk` mit `src/{core,adapters,transport}`, `apps/dashboard`,
  `packages/stream-analyzer`). Die Ausschluss-Klassen sind dort gemischt
  sprachübergreifend und sprachgebunden: `node_modules/` liegt **je Workspace**
  (pnpm — Root plus vier weitere), dazu `dist/`-Generat, `test-results/`,
  `tests/`-Ordner, `*.test.ts` und `*_test.go`. Ohne `exclude` lägen 296 statt
  94 eigene `.ts`-Dateien im Scan-Scope (**zwei Drittel Fremdcode**) — für
  TS-Konsum ist `exclude` damit **Existenzbedingung**, nicht nur Test-Hygiene.
  Zudem matcht der `languages`-Glob `**/*.ts` auch `*.d.ts`-Typ-Generat
  (Suffix-Falle — ins Handbuch-Rezept). Ein **Top-Level**-`exclude` (z. B.
  `["**/node_modules/**", "**/dist/**", "**/*.d.ts", "**/*_test.go",
  "**/*.test.ts", "tests/**", "test-results/**"]`) deckt beide Welten in einer
  Liste; ein je-Sprache-`exclude` müsste die repo-weiten Verzeichnis-Klassen je
  Sprache duplizieren und lückte bei jeder neuen Sprache still. **Die Evidenz
  stützt: Top-Level.** m-trace ist zugleich Kandidat für einen
  Go+TS-Mono-Repo-Pilot (konsumiert bereits d-check mit Digest-Pin v0.2.0 und
  fährt ein eigenes `make arch-check` nur für die Go-API).
- **Wechselwirkung `exclude` ↔ `composition_root`:** beide nehmen Dateien aus,
  aber auf verschiedenen Ebenen (Scan vs. Regel-Auswertung) — Reihenfolge und
  Meldeverhalten (Zähler „gescannte Dateien") dokumentieren.
- **`tech-leak`-Meldetext** bei Adapter-Liste: welcher Adapter wird als
  „erwartet" genannt (alle gelisteten)?
- **Paritäts-Gegenprobe drüben:** nach Release fährt d-check seine Proben-Matrix
  (je Skript-Verbotszweig + Allow-Gegenproben, u. a. `net/url` darf nicht
  flaggen — [AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)-Grenze
  bleibt ehrlich); Rest-Deltas melden die d-check-Belege zurück als Folge-CR.

## 4. Closure-Notiz (nach `done/`)

_(folgt bei Closure — Umsetzung, Gate-Ausgaben, Review, Release + Digest-Re-Pin,
d-check-Entsperrung.)_
