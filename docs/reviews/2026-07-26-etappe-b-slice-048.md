# Review-Report: Etappe B — slice-048 — 2026-07-26

**Review-Art:** Code-Review — geprüft gegen Plan, `AGENTS.md` §3 Hard Rules und
[`harness/conventions.md`](../../harness/conventions.md). Der Gegenstand ist ein reiner
Analyse-Slice ohne Produkt-Code; die Prüflast liegt entsprechend auf der **Belegbarkeit der
Messbehauptungen**.

**Unabhängigkeit — ausdrücklich:** **Selbst-Review**, kein unabhängiger Lauf. Neues
Kontextfenster (Modul 8 §Kontext-Trennung erfüllt), aber dieselbe Modell-Familie wie die
Autoren-Instanz. Der Wiedereinstiegs-Block der [Roadmap](../plan/planning/in-progress/roadmap.md)
hält für diese Kette einen **unabhängigen** Reviewer für angezeigt; dieser Report ersetzt ihn
nicht.

**Gegenstand:** `main..540f599` (4 Commits: `5fe94e8`, `e2a2c4b`, `327e960`, `540f599`) —
185 eingefügte Zeilen in zwei Dateien.

**Skill:** `.harness/skills/reviewer.md` @ `20ee992` (2026-07-25) · <!-- d-check:ignore (Adopter-spezifischer Skill-Pfad, existiert im Ziel-Repo ggf. nicht) -->
**Modell:** `claude-opus-5[1m]` · **Datum:** 2026-07-26

**Eingangs-Kontext** (die Verträge, gegen die geprüft wurde):

- [slice-048](../plan/planning/done/slice-048-modul-delta-lesen.md) (Gegenstand) und
  [slice-046 §6](../plan/planning/done/slice-046-regelwerk-v352-migration-analyse.md) (Etappen-Schnitt)
- [`AGENTS.md`](../../AGENTS.md) §3 Hard Rules, §4 Quality Gates, §5 Dokumentations-Regeln
- [`harness/conventions.md`](../../harness/conventions.md) — `MR-000` … `MR-006` (Stand `main`)
- vendored Baseline `v3.5.2`: `modul-05`, `modul-07`, `modul-10`, `modul-13`
- frühere Findings am selben Bereich:
  [Etappe-A-Zweit-Review](2026-07-25-slice-047-baseline-vendoring-zweitreview.md)
- keine `AC-*`/`ADR-*` berührt (Analyse-Slice, Produkt-Achse unberührt)

---

## Findings

### F-1 — Closure-Kriterium „Gates grün" hält auf dem Slice-Endstand nicht

- `kategorie`: **HIGH**
- `quelle`: Hard-Rule-Klasse „Harness-Lüge" ([`.harness/skills/reviewer.md`](../../.harness/skills/reviewer.md) §Klassifikation);
  [`AGENTS.md`](../../AGENTS.md) §4 (`doc-check` hängt im `gates`-Aggregat), §6 Schritt 8
- `pfad`: [`docs/plan/planning/done/slice-048-modul-delta-lesen.md:161`](../plan/planning/done/slice-048-modul-delta-lesen.md)
  (Closure-Kriterium 1); Ursache in `docs/plan/planning/in-progress/roadmap.md:48` @ `540f599`
- `befund`: Die Closure-Notiz führt als beobachtbares Kriterium „`make gates` grün auf dem Stand
  des Slice (Exit 0) — belegt". Auf dem Stand, der den Slice abschließt (`540f599`, Datei in
  `done/`), exitet `make doc-check` mit **2**: der im Vor-Commit eingefügte Roadmap-Eintrag
  verlinkt `slice-048-modul-delta-lesen.md` verzeichnisrelativ innerhalb von `in-progress/`,
  während der Abschluss-Commit die Datei nach `done/` verschiebt. Repariert wird der Verweis erst
  im **nächsten** Slice (`b319085`, slice-049).
- `verifizierbar`: **ja — belegt.** `make doc-check` auf einem Worktree bei `540f599`:
  `d-check: 119 Datei(en) geprüft, 1 Befund(e)` ·
  `docs/plan/planning/in-progress/roadmap.md:48  slice-048-modul-delta-lesen.md  target-missing` ·
  `make: *** [d-check.mk:30: doc-check] Fehler 1` · `EXIT=2`.
- `gegenprobe` (Skill §Output-Schema: HIGH adversarisch verifizieren): Zwei Entlastungen wurden
  geprüft und tragen **nicht** vollständig. (1) *„Der Gate-Lauf fand vor dem `git mv` statt"* —
  die Behauptung lautet „auf dem Stand des Slice", und der Zustand eines Slice **ist** laut
  [`AGENTS.md`](../../AGENTS.md) §5 das Verzeichnis; der abschließende Stand ist gerade der rote.
  (2) *„nur ein transienter Zwischenstand, gemergt wird die Spitze"* — trägt für die
  Merge-Fähigkeit der Kette, nicht für die Belegtreue der Notiz. Die Merge-Folge steht im
  Verdikt.

### F-2 — B-17 führt einen bereits aufgelösten `MR` als offenen Migrations-Fund

- `kategorie`: **MEDIUM**
- `quelle`: [`harness/conventions.md` §Adaptions-Block](../../harness/conventions.md#adaptions-block)
  (Disziplin: Aufhebung ist die vorgesehene Form); Skill §MEDIUM „unbelegte Tatsachenbehauptung"
- `pfad`: [`docs/plan/planning/done/slice-048-modul-delta-lesen.md:91`](../plan/planning/done/slice-048-modul-delta-lesen.md) (B-17), Zuweisung in §5
- `befund`: B-17 stellt fest, [`MR-003`](../../harness/conventions.md#mr-003--source-precedence-ohne-docsuser-rang)
  sei „gegenstandslos geworden" und „gehört gestrichen, nicht gepflegt", und weist den Fund
  Etappe C zu. Der Eintrag trägt bei `main` bereits seit 2026-06-21 einen ausdrücklichen
  Abschluss-Block („**Aufgelöst (2026-06-21)**", `harness/conventions.md:142`) mit demselben
  Sachgrund — dem eingefügten `docs/user`-Rang. Der Befund nennt diesen Block nicht.
- `verifizierbar`: ja — `git show main:harness/conventions.md` Zeile 142 ff.

### F-3 — Zahlenpaar „20 Dateien / 2867 Zeilen" ist nicht deckungsgleich

- `kategorie`: **LOW**
- `quelle`: Skill §MEDIUM/§LOW (Beleg-Genauigkeit); der Wert ist als Umfangs-Beleg geführt
- `pfad`: [`docs/plan/planning/done/slice-048-modul-delta-lesen.md:14,25`](../plan/planning/done/slice-048-modul-delta-lesen.md);
  wiederholt in `roadmap.md` (Etappe-B-Absatz)
- `befund`: §1 nennt „20 Dateien, 2867 Zeilen" und zählt dazu `modul-00` … `modul-16` plus die
  drei Grundlagen-Abschnitte auf — das sind 20 Dateien mit zusammen **2778** Zeilen. Auf 2867
  kommt man nur einschließlich der Index-`README.md` (89 Zeilen), die in der Aufzählung nicht
  vorkommt. Die Aussage „vollständig gelesen" ist davon unberührt; die beiden Zahlen beziehen
  sich auf verschiedene Mengen.
- `verifizierbar`: ja — Zeilenzählung über `.harness/baseline/v3.5.2/regelwerk/` mit und ohne
  `README.md`.

### F-4 — B-19 ist INFO klassifiziert, trägt aber einen Handlungsauftrag

- `kategorie`: **LOW**
- `quelle`: `modul-10` §Finding-Kategorien („INFO = Hinweis, keine Aktion erwartet")
- `pfad`: [`docs/plan/planning/done/slice-048-modul-delta-lesen.md:94`](../plan/planning/done/slice-048-modul-delta-lesen.md) (B-19), §5 Etappe-C-Zuweisung
- `befund`: B-19 ist als INFO geführt, formuliert aber eine Bedingung mit Handlungsfolge („dann
  gehört er als **bewusste Abweichung mit `MR-*`** deklariert, nicht stillschweigend
  ausgelassen") und wird in §5 einer Etappe zugewiesen. Kategorie und abgeleitete Handlung
  stimmen nicht überein.
- `verifizierbar`: nein — Klassifikations-Frage, kein Gate-Lauf.

### F-5 — Status-Feld widerspricht dem Verzeichnis

- `kategorie`: **LOW**
- `quelle`: [`AGENTS.md`](../../AGENTS.md) §5 („der Zustand ist das Verzeichnis, kein Feld im Dokument")
- `pfad`: [`docs/plan/planning/done/slice-048-modul-delta-lesen.md:3`](../plan/planning/done/slice-048-modul-delta-lesen.md)
- `befund`: Die Datei liegt in `done/` und trägt im Kopf „**Status:** in-progress". Kein
  Einzelfall: bei `main` tragen `slice-012`, `slice-014`, `slice-018` und `slice-047` in `done/`
  dasselbe Feld mit demselben Widerspruch — fünf gleichartige Vorfälle, also oberhalb der
  3×-Schwelle aus [`AGENTS.md`](../../AGENTS.md) §5 §Steering-Loop, ohne Eintrag im (zum
  Zeitpunkt dieses Slice noch nicht existierenden) Kanal.
- `verifizierbar`: ja — Feld-Abgleich `done/`-Dateien gegen Verzeichnis; heute wäre ein Sensor in
  `make verify` der Ort.

## Negativbefunde

- geprüft, ohne Befund: **B-1 (Slice-Größen-Messung)** — `slice-047` = 7, `slice-046` = 6,
  `slice-044` = 4 DoD-Punkte, exakt wie behauptet, gegen `main` nachgezählt.
- geprüft, ohne Befund: **B-4 (`closure-note-reviewer`-Vorlage)** — `SHA256SUMS` Zeile 35 trägt
  tatsächlich `templates/.harness/skills/closure-note-reviewer.template.md`.
- geprüft, ohne Befund: **B-11 (`nolintlint`)** — bei `main` kein `nolintlint` in `.golangci.yml`;
  die Hard-Rule-Lücke besteht wie beschrieben.
- geprüft, ohne Befund: **B-15 (AC-Form, „16 von 19")** — mit **zwei** Lesarten nachgerechnet:
  „Wort kommt vor" ergibt 17/19, „eigener Out-of-Scope-**Block**" ergibt genau **16/19**
  (`AC-QA-01`/`AC-QA-02` ohne, `AC-QA-03` nur im Fließtext). Unter der im Befund gemeinten
  Block-Lesart trägt die Zahl. Ein zunächst als HIGH notierter Verdacht ist an dieser Gegenprobe
  gefallen.
- geprüft, ohne Befund: **B-16 (`.claude/commands/`)** — bei `main` enthält `.claude/` nur `hooks`
  und `settings.json`.
- geprüft, ohne Befund: **B-18 (`next/`)** — bei `main` existieren unter
  `docs/plan/planning/` nur `done`, `in-progress`, `open`.
- geprüft, ohne Befund: **§2 Steering-Loop-Zählung** — „acht von 48" trifft exakt: Bezug in
  `slice-001` … `slice-008`, danach lückenlos keiner.
- geprüft, ohne Befund: **§2 Modus-Zählung („zwölf von 48")** — eine bewusst zu breite Suche
  (`sub-area|greenfield|brownfield|modus`) ergab 19 Treffer; die sieben Zusatztreffer
  (`slice-015/016/020/021/022/031/045`) betreffen durchweg die **Auflösungs-Modi des Produkts**
  (`fixed-root`/`relative`/`namespace`), keinen Sub-Area-Modus. Die Zahl trägt.
- geprüft, ohne Befund: **Hard Rule §3.3** (`git mv` + Inhalt = zwei Commits) — `540f599` ist ein
  reiner Rename (`1 file changed, 0 insertions, 0 deletions`, R-Erkennung greift).
- geprüft, ohne Befund: **Traceability** ([`AGENTS.md`](../../AGENTS.md) §5) — alle vier Commits
  nennen mindestens eine `slice-NNN`-ID.
- geprüft, ohne Befund: **Negativbefund-Pflicht des Gegenstands selbst** (`modul-10`) — §4 des
  Slice führt zwölf Negativbefunde; Schweigen wird nicht als Konformität gewertet.
- geprüft, ohne Befund: **Eigen-Probe zu B-1** — §6 des Slice hält die gemeldete Regel selbst ein
  (genau drei DoD-Punkte).
- geprüft, ohne Befund: **Gate-Landschaft („13 Targets")** — aus den 16 Makefile-Targets bei
  `main` abzüglich `help`/`compile`/`build` herstellbar; die Zählbasis ist im Slice nicht
  angegeben, die Zahl ist aber nicht widerlegbar.

## Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 1 |
| MEDIUM | 1 |
| LOW | 3 |
| INFO | 0 |

## Verdikt

**Merge-blockierend: nein — mit einer Auflage.**

Begründung, ausdrücklich statt still: F-1 ist nach Kategorie ein HIGH, weil eine Gate-Behauptung
auf dem belegten Stand nicht hält. Es blockiert den Merge **dieser Kette** trotzdem nicht, weil
gemergt wird die Spitze und der rote Verweis bereits im Folge-Commit `b319085` repariert ist —
der Merge-Gegenstand ist nicht der rote Zustand. Was F-1 verlangt, ist die Korrektur der
**Behauptung** (Closure-Kriterium 1) beziehungsweise ein Sensor, der sie künftig trägt: das ist
exakt der offene Punkt **SL-002** im Steering-Loop („brechende
Verweise nach `git mv`", 7 Vorfälle). Dieser Report liefert dafür den ersten Beleg **mit rotem
Gate-Lauf** statt mit Erinnerung.

F-2 ist vor dem Merge zu klären, weil Etappe C den Fund verarbeitet hat — ob dort korrekt, prüft
das Review der Etappe C, nicht dieses.

**Übergabe:** Findings gehen an die Implementation. Der Report ersetzt keine Verifikation —
DoD-/Spec-Konformität prüft `make verify` separat (Modul 11).
