# Review-Report: slice-047 — 2026-07-25

**Review-Art:** Code — geprüft gegen Plan
([slice-047](../plan/planning/done/slice-047-baseline-vendoring.md)),
den Etappen-Schnitt aus [slice-046 §6](../plan/planning/open/slice-046-regelwerk-v352-migration-analyse.md)
und die Konventionen ([`AGENTS.md`](../../AGENTS.md) Hard Rules,
[`harness/conventions.md`](../../harness/conventions.md)).

**Gegenstand:** slice-047 · Commit `cc3c225` (die Harness-Änderung; die Range `origin/main...HEAD` umfasst zusätzlich `05ed884` = slice-046-Analyse, **nicht** Gegenstand dieses Reports) · 48 Dateien.

**Skill:** `.harness/skills/reviewer.md` — Bestand **vor** der Baseline-Hebung (noch nicht gegen
`v3.5.2` Modul 10 geprüft, siehe [slice-046 §3](../plan/planning/open/slice-046-regelwerk-v352-migration-analyse.md)); dieser Report folgt der **vendored Vorlage**
`.harness/baseline/v3.5.2/templates/docs/reviews/review-report.template.md` — der erste in diesem
Repo, der das tut.
**Modell:** `claude-opus-5[1m]` · **Datum:** 2026-07-25

**Eingangs-Kontext** (die Verträge, gegen die geprüft wurde):

- Plan: [slice-047](../plan/planning/done/slice-047-baseline-vendoring.md), Etappe A aus [slice-046](../plan/planning/open/slice-046-regelwerk-v352-migration-analyse.md)
- [`harness/conventions.md`](../../harness/conventions.md) §Baseline + Adaptions-Block (`MR-000`…`MR-006`)
- [`AGENTS.md`](../../AGENTS.md) §1 (Regelwerks-Pflicht), §3 Hard Rules, §4 Gates
- [`.d-check.yml`](../../.d-check.yml) (Doku-Gate-Konfiguration)
- vendored Baseline: `regelwerk/modul-02-harness-bootstrap.md` (Vendoring), `modul-10-review-harness.md` (diese Form), `templates/harness/conventions.template.md`
- **Keine** berührten `AC-*`/`ADR-*` — die Änderung liegt außerhalb der Produkt-Achse.

> **Form-Einschränkung.** Reviewer und Verifier sind dieselbe Instanz; die Rollentrennung aus
> Modul 8/11 ist nicht erfüllt. Die DoD-Verifikation steht getrennt in
> [slice-047 §3/§4](../plan/planning/done/slice-047-baseline-vendoring.md).

---

## Findings

### F-1 — Provenienz des vendored Baums ist nicht gegatet

- `kategorie`: MEDIUM
- `quelle`: [AC-QA-03](../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit) (Reproduzierbarkeit, sinngemäß auf die Baseline übertragen) / `MR-006`
- `pfad`: `.harness/baseline/v3.5.2/SHA256SUMS`
- `befund`: Das Manifest ist im Repo erzeugt worden und belegt nur die **innere** Konsistenz des Baums. Dass dieser Baum das Release `v3.5.2` ist — die Aussage, die `conventions.md` §Baseline trifft —, prüft kein Gate; ein vertauschter oder editierter Baum bliebe unentdeckt.
- `verifizierbar`: teilweise — die **Asset**-Integrität ist gate-fähig (d-checks `sources`-Modul, `source-pin` auf den `sha256`); die **Freshness** („gibt es einen neueren Tag?") ist laut `modul-02-harness-bootstrap.md` §Freshness-Audit ausdrücklich eine **Netz-Operation außerhalb der Gates** und gehört nicht in `gate-consistency`. *(Korrigiert nach dem Zweit-Review, siehe Nachtrag.)*

### F-2 — Gate-Lauf an den Commit gekettet, roter `doc-check` committet

- `kategorie`: MEDIUM
- `quelle`: Hard Rule „Kein Erfolg ohne echte Gate-Ausgabe" ([`AGENTS.md` §6](../../AGENTS.md#6-minimal-agent-workflow), [`CLAUDE.md`](../../CLAUDE.md))
- `pfad`: Prozess (kein Repo-Pfad)
- `befund`: Der Slice-Commit entstand im selben Kommando wie der `doc-check`-Lauf und wurde ausgeführt, obwohl der Lauf mit Exit 2 endete (zwei unverlinkte `MR-*`-Kennungen). Es ist das **dritte** Auftreten derselben Klasse an diesem Tag; die zuvor daraus abgeleitete Arbeitsregel („Lauf und Commit nie in einem Kommando") wurde nicht eingehalten.
- `verifizierbar`: ja — `make doc-check` mit ausgewertetem Exit-Code.

### F-3 — Konfigurations-Kommentar hängt an einer Nummer, die Etappe C ändern kann

- `kategorie`: LOW
- `quelle`: Maintainability
- `pfad`: `.d-check.yml:16-18`
- `befund`: Der Kommentar zum `scan.ignore` begründet den Ausschluss mit `MR-006`. Etappe C ist ausdrücklich vorgesehen, die MR-Nummerierung an das Baseline-Template anzugleichen; danach zeigt der Kommentar auf eine Nummer, die es nicht mehr gibt. Kein Gate deckt Kommentar-Text ab.
- `verifizierbar`: nein

### F-4 — Plan und Ausführung liegen in einem Commit *(Präferenz, kein Normverstoß)*

- `kategorie`: INFO
- `quelle`: Maintainability — **keine** normative Stelle verlangt die Trennung Plan-/Umsetzungs-Commit (§3.3 gilt nur für `git mv` + Inhaltsänderung). Als Finding daher zurückgestuft; ein Befund ohne zitierbare Stelle ist selbst eine Beweislast-Umkehr.
- `pfad`: Commit `cc3c225`
- `befund`: Slice-Dokument und Harness-Änderung sind in einem Commit vereint (Folge eines `--amend`, der mehr aufnahm als beabsichtigt). Der Verlauf zeigt damit nicht, was geplant und was ausgeführt wurde; die Commit-Nachricht ist nachträglich auf den vollen Inhalt korrigiert.
- `verifizierbar`: nein

### F-5 — Die Quelldatei-Kette ist aus dem Briefing verschwunden

- `kategorie`: LOW *(hochgestuft nach dem Zweit-Review: betrifft eine Präzedenz-Aussage, nicht bloß einen Hinweis)*
- `quelle`: Maintainability
- `pfad`: `AGENTS.md:18-33`
- `befund`: Vor der Änderung nannte `AGENTS.md` §1 die Upstream-Quelldatei (`agents-regelwerk.md`) und die Regel „bei Konflikt gilt die Quelldatei". Jetzt verweist §1 auf den Kurs-Tree; die Derivativ-Kette läuft nur noch über `conventions.md` §Adoptierte Konventions-Quellen.
- `verifizierbar`: nein

### F-6 — Der Gate-Nachweis umfasst jetzt 43 Fremddateien

- `kategorie`: INFO
- `quelle`: Maintainability
- `pfad`: `.harness/state/gates-passed.diffsha` (Mechanik), `.harness/baseline/v3.5.2/**`
- `befund`: `record-gates` bildet den Working-Tree-Hash; der vendored Baum liegt darin. Ein Baseline-Austausch erscheint damit als gate-relevanter Diff und erzwingt einen neuen Gate-Lauf. Das ist konsistent, aber neu — der Nachweis mischt jetzt Eigen- und Fremdinhalt.
- `verifizierbar`: ja — `make record-gates` vor/nach einer Baseline-Änderung.

## Negativbefunde

- geprüft, ohne Befund: **Provenienz-Inhalt** — der vendored Baum wurde gegen ein **frisch nachgeladenes** Release-ZIP verglichen (`diff -r` über `regelwerk/` und `templates/`): **identisch**, keine Modifikation beim Vendoring. (Behebt F-1 *nicht* — eine einmalige Handprobe ist kein Gate.)
- geprüft, ohne Befund: `spec/**` — kein Stratum berührt, kein Versions-Bump, keine `AC-*`-ID.
- geprüft, ohne Befund: `internal/**`, `cmd/**` — kein Go-Code im Diff; `lint`, `test`, `coverage-gate` (96,20 %), `arch-check` (0 Befunde) unverändert grün.
- geprüft, ohne Befund: `version.md`, `a-check.mk`, `README*.md` — Release-Pins unberührt, `gate-consistency` Pin-Gleichheit grün.
- geprüft, ohne Befund: `Makefile`, `tools/**`, `.github/workflows/**` — keine Gate-Definition geändert; `gate-consistency` Doku↔Makefile grün.
- geprüft, ohne Befund: Doku-Integrität — `doc-check` **0 Befunde**; die geprüfte Dateizahl wächst nur um die neuen Eigen-Dokumente dieses Branches (116 nach dem Vendoring-Commit), **nicht** um die 43 vendored Dateien — ohne `scan.ignore` wären es rund 158.
- geprüft, ohne Befund: `MR-000`…`MR-005` — inhaltlich **nicht** angetastet (bewusst Etappe C); nur `MR-006` neu.

## Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 0 |
| MEDIUM | 2 |
| LOW | 2 |
| INFO | 2 |

## Verdikt

**Merge-blockierend: nein** für die *hier* gelisteten Findings — begründet statt still entschieden:
F-1 ist im Plan als offen ausgewiesen und einer benannten Etappe zugeordnet (Verklemmung sonst),
F-2 betrifft den Weg, nicht den grünen Zustand.

> **Dieses Verdikt trägt nicht allein.** Ein **unabhängiges Zweit-Review** (eigener Report, s.
> Nachtrag) fand **sechs** MEDIUM statt zwei und urteilte **merge-blockierend: ja** — zu Recht: vier
> der zusätzlichen MEDIUM hatten weder eine benannte Etappe noch ein Verklemmungs-Argument.

**Übergabe:** F-1 → Etappe D (Asset-Teil; Freshness bleibt Netz-Wartung). F-3 → Etappe C.
F-2 → Durchsetzungsschicht, außerhalb dieses Slices. F-4/F-5/F-6 → keine Aktion.
Der Report ersetzt keine Verifikation — die DoD-Konformität steht in
[slice-047 §3/§4](../plan/planning/done/slice-047-baseline-vendoring.md).

---

## Nachtrag (2026-07-25) — was das unabhängige Zweit-Review korrigiert hat

Report des Zweit-Reviews: [`2026-07-25-slice-047-baseline-vendoring-zweitreview.md`](2026-07-25-slice-047-baseline-vendoring-zweitreview.md).
Es fand **zehn Findings, die hier fehlten** — Ursache war eine **Abdeckungslücke**, die der
Negativbefund-Block korrekt *nicht* behauptete: keine Zeile für `.harness/skills/**`, für
`docs/plan/planning/**`/`roadmap.md`, für `docs/reviews/README.md` und keine für den
`.d-check.yml`-**Kopf**. Genau dort lagen die Befunde.

Drei **Fehler in diesem Report** sind oben korrigiert (Erratum, kein Umschreiben des Urteils):

| Korrektur | vorher | jetzt |
|---|---|---|
| `pfad` in F-3 | `.d-check.yml:6` (dort steht die Modul-Aufzählung) | `.d-check.yml:16-18` |
| `Gegenstand` | „Diff `origin/main...HEAD` (1 Commit)" — Range und Commit-Zahl bezeichneten verschiedene Mengen | Commit `cc3c225`, Range-Umfang ausgewiesen |
| F-1 `verifizierbar` | „in `gate-consistency`" | nach Modul 02 getrennt: Asset gate-fähig, **Freshness außerhalb der Gates** |

Zwei **Kategorie-Korrekturen**: F-4 → **INFO** (das Zweit-Review bestätigt es nicht als Finding —
es gibt keine normative Stelle für die Trennung Plan-/Umsetzungs-Commit; ein Befund ohne zitierbare
Stelle ist selbst eine Beweislast-Umkehr). F-5 → **LOW** (Präzedenz-Aussage in einem
Rang-8-Dokument, nicht bloß ein Hinweis).

Ein **Negativbefund war zu stark**: „`MR-000`…`MR-005` inhaltlich nicht angetastet" ist textuell
richtig, verschwieg aber, dass der neue §Baseline-Satz „bei Konflikt gewinnt der Default" ihren
**normativen Status** änderte. Die Klausel ist daraufhin eingegrenzt worden — die deklarierten
`MR-*` bleiben bis zur Prüfung in Etappe C ausdrücklich in Kraft.
