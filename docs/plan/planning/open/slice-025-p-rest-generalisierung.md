# slice-025 — P-Rest-Generalisierung: `constructs`-Regel (+ Import-Allowlist, gated)

**Status:** open (Entwurf **abgenommen 2026-07-04** — spec-first; Umsetzung folgt als eigener Slice, noch kein Code).
**Bezug:** erweitert die Scoping-Semantik von
[AC-FA-RULE-003](../../../../spec/lastenheft.md#ac-fa-rule-003--tech-kapselung-regel-tech-leak)
(`tech`-Muster) auf Roh-Text; bewegt sich an der Heuristik-Grenze von
[AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze);
Konfigurations-Vertrag [AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml).
[Roadmap](../in-progress/roadmap.md).

## 1. Auslöser — Maintainer-Ziel: Skript-Copies verringern

Das Gründungsziel von a-check (und d-check) ist die Reduktion der Skript-Kopien in den
Konsumenten-Repos: „Konfiguration statt Fork" (README §Warum). Stand nach dem
b-cad-Pilot-Schnitt (2026-07-04, dortige slice-030):

- **d-check:** 0 Skriptzeilen — `arch-check` läuft vollständig über `.a-check.yml` +
  digest-gepinntes Image (dortige ADR 0029); das dortige `tools/arch-check.sh` ist gelöscht.
- **b-cad:** `tools/arch-check.sh` von voller Hexagon-Durchsetzung auf einen **P-Rest**
  von drei grep-Regeln geschrumpft (75 Zeilen, überwiegend Doku-Kommentar) — Muster,
  die a-check strukturell nicht sieht (§2).
- **grid-guide / d-migrate / m-trace:** noch keine Konsumenten (kein `.a-check.yml`).

Der P-Rest ist damit kein Fork mehr (keine der sieben Regeln dupliziert), aber ein
**Zwischenstand mit Divergenz-Uhr**: sobald ein zweiter Konsument ein P-Rest-artiges
Muster braucht (naheliegend: grid-guide/Rust mit einem `unsafe`-Zonen-Monopol — strukturgleich
mit dem dlopen-Monopol), beginnen die grep-Kopien wieder zu wandern. Genau das Szenario,
das a-check gegründet hat. Dieser Slice sammelt darum die Generalisierung als
**CR-Kandidaten** nach dem Evidenz-Muster des d-check-Delta-CR
([slice-023](../done/slice-023-dcheck-pilot-deltas.md), [ADR-0018](../../adr/0018-exclude-scan-scope.md)).

## 2. Die drei P-Rest-Muster (Evidenz: b-cad slice-030-Skript)

| Muster | b-cad-Regel | Warum a-check es heute nicht sieht |
|---|---|---|
| **Aufruf-Monopol** | P1: `dlopen/dlsym/dlclose`-**Aufruf** nur in `src/adapters/plugin/` (auch nicht in `plugins/`, nicht in `main.cpp`) | Falsche Sprachebene: ein Funktionsaufruf ist keine Import-Zeile; kein Backend matcht Ausdrucks-Zeilen (bewusste Design-Linie, vgl. TypeScript-Extraktion in [AC-FA-EXTRACT-001](../../../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion)). Die Include-Hälfte (`dlfcn.h`) deckt die `tech`-Regel ab; der Aufruf kann aber ohne eigenen Include existieren (transitiv/Prototyp). `forbidden_constructs` ist an die Port-Disziplin gebunden ([AC-FA-RULE-004](../../../../spec/lastenheft.md#ac-fa-rule-004--port-disziplin-regel-port-impurity)) und erreicht Nicht-Port-Schichten sowie Dateien außerhalb aller Layer-Globs nicht. |
| **Fail-closed-Allowlist** | P2: Quote-Includes in `plugins/` + `src/plugin_api/` nur aus drei Präfixen — **alles andere verboten, auflösbar oder nicht** | Inverse Beweislast: a-check beurteilt nur Importe, die sich auf eine deklarierte Schicht **auflösen** ([SPEC-EXTRACT-001](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion)); Unauflösbares bleibt fail-open unbeurteilt (kein Geister-Match, dokumentierte Grenze). Eine Default-verboten-Semantik existiert im Kanten-Modell nicht. |
| **Form-Verbot** | P2c: Projekt-Präfixe nie als **Angle**-Include in `plugins/` + `src/plugin_api/` | Die Form-Dimension ist im Modell gelöscht: der C++-Extraktor liest `<…>` und `"…"` mit einem Regex und verwirft die Klammer-Form (`internal/adapter/driven/extract/extract.go`, `cppInclude`). Der Grund der Invariante ist Build-System-Wissen (Include-Pfad-Layout), das der text-heuristische Vertrag ([ADR-0002](../../adr/0002-text-heuristische-extraktion.md)) bewusst nicht hat. Der „C++-quoted-Include-Split" steht in der [Roadmap](../in-progress/roadmap.md) bereits als Re-Eval-Trigger. |

## 3. CR-Kandidat 1 — `constructs`-Regel: Roh-Text-Muster mit `tech`-Scoping

**Idee:** die bewährte `tech`-Scoping-Mechanik (Adapter-Zone als Skalar **oder Liste**,
`match: substring|regex` mit RE2 ([ADR-0015](../../adr/0015-regex-tech-muster.md)),
`composition_root: allow|forbid`, fail-closed bei leerem/fehlendem Ziel — alles seit
Lastenheft 0.14.0 vorhanden) von **Import-Symbolen** auf **Roh-Quelltext** heben.
Skizze (Schema-Ort in der ADR zu entscheiden, §6):

```yaml
constructs:
  - {pattern: '\bdl(m?open|sym|close)\s*\(', match: regex,
     adapter: adapters/plugin, composition_root: forbid}
  - {pattern: '#include\s*<(adapters/|hexagon/|plugin_api/)', match: regex,
     forbid_in: [plugins, plugin_api]}
```

Zwei Scoping-Gestalten, beide aus der b-cad-Evidenz:

- **Monopol** (`adapter:` wie bei `tech`): Muster nur in der Zone erlaubt, überall sonst
  Befund — geprüft auf **allen** gescannten Dateien (auch außerhalb der Layer-Globs, z. B.
  `main.cpp`; `composition_root: forbid` schließt die Verdrahtungs-Ausnahme). Deckt **P1**.
- **Zonen-Verbot** (`forbid_in:` je Schicht): Muster in den genannten Schichten verboten,
  anderswo egal — die Generalisierung des heute port-gebundenen `forbidden_constructs`
  auf beliebige Schichten. Deckt **P2c**.

**Anforderungs-Skizze** (neue Anforderung im Bereich `RULE`; ID-Vergabe mit
Versions-Bump + Historie-Zeile erst im Lastenheft-CR selbst — Anlege-Prozess
[`harness/conventions.md`](../../../../harness/conventions.md#anforderungs-anlege-prozess)):

- **Happy:** Given ein `constructs`-Monopol-Eintrag (`dlopen`-Muster, Zone
  `adapters/plugin`), when eine Datei in `adapters/io` das Muster enthält, then Befund
  (Datei, Zeile, Muster, erlaubte Zone[n]) und Exit 1; die Zone selbst bleibt befundfrei.
- **Boundary:** Given `composition_root: forbid`, when die deklarierte Composition Root
  das Muster enthält, then Befund (ohne `forbid`: kein Befund dort). Given eine Datei, die
  in einem `languages`-Glob, aber **keinem** Layer-Glob liegt, when sie ein Monopol-Muster
  enthält, then Befund (Roh-Text-Prüfung ist scan-weit, nicht layer-gebunden).
- **Negative:** Given ein Eintrag mit leerem/fehlendem `adapter`/`forbid_in` oder
  unbekanntem Schlüsselwert, when a-check lädt, then Exit 2 (fail-closed — Muster der
  0.14.0-Härtung von [AC-FA-RULE-003](../../../../spec/lastenheft.md#ac-fa-rule-003--tech-kapselung-regel-tech-leak)).
- **Out-of-Scope:** kein Parser — Kommentar-/String-Grenzen bleiben Text-Heuristik
  ([AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze));
  keine RE2-fremden Features (Lookaround/Negation); keine Auto-Inferenz von Zonen.

**Wirkung:** b-cads `arch-check.sh` schrumpft auf **eine** Regel (P2); der
Roadmap-Re-Eval-Trigger „quoted-Include-Split" ist miterledigt (P2c als Muster
ausdrückbar, ohne den Extraktor form-bewusst zu machen).

## 4. CR-Kandidat 2 — fail-closed Import-Allowlist je Schicht (**gated**)

Skizze: optionales `import_allowlist: [präfixe]` je Schicht — jeder extrahierte
Import-Specifier der Schicht muss (roh, vor Auflösung) mit einem Präfix beginnen, sonst
Befund. Deckt **P2**; damit fiele b-cads Skript vollständig (**0 Architektur-Skripte in
der Flotte**).

**Bewusst gated:** die Umkehr der Beweislast (default-verboten auf Roh-Text) steht dem
Kein-Geister-Match-Ethos der Heuristik-Grenze am nächsten — als explizit deklarierte
Opt-in-Allowlist vertretbar, aber erst bei einem **zweiten realen Konsumenten-Bedarf**
auszuarbeiten. Hier kein AK-Entwurf.

## 5. Evidenz-Stand und Lande-Trigger

| Konsument | Stand | P-Rest-Bedarf |
|---|---|---|
| b-cad | Pilot-Schnitt in Arbeit (dortige slice-030) | P1 + P2 + P2c — Skript-Header dokumentiert die drei Muster präzise (Evidenz-Quelle) |
| d-check | umgestellt, 0 Skript | keiner (Go-Regeln waren vollständig kanten-förmig) |
| grid-guide (Rust) | Pilot offen | erwartet: `unsafe`-Zonen-Monopol (Analogon zu P1) — **der natürliche zweite Evidenz-Geber** |
| d-migrate (Kotlin) | Pilot offen; dort heute Review statt Gate | offen |
| m-trace (Go+TS) | Pilot-Kandidat | offen |

**Lande-Trigger:** Kandidat 1 auf Maintainer-Wort oder mit dem zweiten Konsumenten
(grid-guide-Pilot); Kandidat 2 erst bei zweitem realem Allowlist-Bedarf.

## 6. Vor der Umsetzung zu klären (ADR-Skizze)

- **Schema-Ort:** eigener Top-Level-Block `constructs` (Skizze §3) **vs.** Generalisierung
  des bestehenden `forbidden_constructs` (heute Port-gebunden,
  [AC-FA-RULE-004](../../../../spec/lastenheft.md#ac-fa-rule-004--port-disziplin-regel-port-impurity))
  **vs.** beides in einem Block mit zwei Scoping-Arten. Befund-Name (`construct-leak`?)
  und Verhältnis zum bestehenden `port-impurity`-Pfad.
- **Kommentar-Stripping:** Roh-Text-Prüfung vor oder nach `prepSource`? (Python-Lehre aus
  slice-020: C-Stripping kann falsch-grün erzeugen; für Aufrufmuster ist ein Treffer im
  Kommentar umgekehrt falsch-rot — je Sprache entscheiden, Parität zur bash-grep-Referenz
  von b-cad belegen.)
- **Scan-Menge:** Monopol-Muster scan-weit (alle `languages`-Dateien, `exclude` greift
  davor, [ADR-0018](../../adr/0018-exclude-scan-scope.md)) — Determinismus/stabile
  Sortierung der Befunde unverändert ([AC-QA-01](../../../../spec/lastenheft.md#ac-qa-01--determinismus)).
- **Abgrenzung dokumentieren:** `tech` = extrahierte Import-Symbole, `constructs` =
  Roh-Text — beides ausgewiesene Heuristik, keine Semantik-Behauptung
  ([AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)).
- **Rückbau-Pfad:** nach Landung b-cad-Skript auf P2 schrumpfen (Fitness-Probe:
  injizierter `dlopen`-Aufruf außerhalb der Zone ⇒ Befund im a-check-Gate statt im Skript).

## 7. DoD (bei Ausarbeitung)

Spec-first-Reihenfolge (Lastenheft-CR mit neuer `RULE`-ID → ADR → Spezifikation →
Code → Tests); `make gates` grün; Fixture-Paritätsprobe gegen die drei
b-cad-grep-Regeln (gleiche Treffer, gleiche Nicht-Treffer); Benutzerhandbuch-Currency.
