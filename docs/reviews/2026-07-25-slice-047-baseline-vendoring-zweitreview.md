# Review-Report: slice-047 — 2026-07-25 (unabhängiges Zweit-Review)

**Review-Art:** Code — geprüft gegen Plan
([slice-047](../plan/planning/done/slice-047-baseline-vendoring.md)), den Etappen-Schnitt
([slice-046 §6](../plan/planning/open/slice-046-regelwerk-v352-migration-analyse.md)), die
Konventionen ([`AGENTS.md`](../../AGENTS.md), [`harness/conventions.md`](../../harness/conventions.md),
[`CLAUDE.md`](../../CLAUDE.md)) und die **vendored Baseline-Erwartung**
(`modul-02-harness-bootstrap.md`, `conventions.template.md`, `modul-10-review-harness.md`).

**Gegenstand:** Branch `slice-047-baseline-vendoring`, `origin/main...HEAD` = 3 Commits
(`05ed884`, `cc3c225`, `4d2d8e3`), 51 Dateien / +5226 / −27.

**Reviewer:** **unabhängige zweite Instanz** (separater Agenten-Lauf, read-only: keine Dateiänderung,
keine schreibenden `make`-Targets). Damit ist die Rollentrennung aus Modul 8/10 für diesen Lauf
**erfüllt** — anders als beim Erst-Report
([`…-baseline-vendoring.md`](2026-07-25-slice-047-baseline-vendoring.md)), der sie als Einschränkung
ausweisen musste.
**Modell:** `claude-opus-5[1m]` · **Datum:** 2026-07-25

**Eingangs-Kontext:** Plan + Etappen-Schnitt · `harness/conventions.md` (§Baseline, `MR-000`…`MR-006`)
· `AGENTS.md` §1/§3/§4 · `CLAUDE.md` · `harness/README.md` · `.d-check.yml` ·
vendored `regelwerk/modul-02` (Vendoring/Freshness), `modul-10` (Form),
`templates/harness/conventions.template.md`, `templates/docs/reviews/review-report.template.md` ·
**keine** berührten `AC-*`/`ADR-*`.

---

## Findings

### F-1 — Reviewer-Skill behauptet weiter Regelwerk `v1.3.0`

- `kategorie`: MEDIUM
- `quelle`: Hard Rule „stille Setzung / falsche Tatsachenbehauptung" / `MR-006`
- `pfad`: `.harness/skills/reviewer.md:7`
- `befund`: Die Zeile nennt als Autorität „Regelwerk v1.3.0 Modul 10" und verweist auf denselben §Baseline-Abschnitt, den dieser Diff auf `v3.5.2` hebt. Skill und der mitgelieferte Report nennen damit unterschiedliche Regelwerksstände.
- `verifizierbar`: nein durch bestehende Gates (kein Modul prüft Versions-Strings); reproduzierbar per `grep -rn "v1\.3\.0"`.

### F-2 — `.d-check.yml` deklariert im Kopf weiter „Baseline v1.3.0"

- `kategorie`: MEDIUM
- `quelle`: Maintainability / `MR-006` §Abgrenzung
- `pfad`: `.d-check.yml:4`
- `befund`: Der Kopfkommentar endet mit „Baseline v1.3.0."; zwölf Zeilen darunter fügt derselbe Commit den `scan.ignore`-Kommentar mit Bezug auf `MR-006`/vendored `v3.5.2` ein. Dieselbe Datei nennt zwei verschiedene Baseline-Stände.
- `verifizierbar`: nein

### F-3 — Der Dateizahl-Beleg widerspricht sich über drei Artefakte

- `kategorie`: MEDIUM
- `quelle`: Hard Rule „Kein Erfolg ohne echte Gate-Ausgabe"
- `pfad`: `docs/plan/planning/in-progress/slice-047-baseline-vendoring.md:51`
- `befund`: Der Plan nennt „115 Dateien … (sonst wären es 158)", die Commit-Message und der Erst-Report nennen 116. Real: 115 bei `05ed884`, 116 bei `cc3c225`, 117 auf `HEAD`. Die Zahl ist die **einzige** vorgelegte Evidenz dafür, dass der `scan.ignore` greift.
- `verifizierbar`: ja — `make doc-check` mit ausgewerteter Dateizahl.

### F-4 — `slice-047` liegt in `in-progress/`, wird in der Roadmap aber nicht genannt

- `kategorie`: MEDIUM
- `quelle`: `DC-FA-PLAN-001` / [`harness/README.md` §Traceability rules](../../harness/README.md#traceability-rules)
- `pfad`: `docs/plan/planning/in-progress/roadmap.md:39`
- `befund`: Der Roadmap-Diff trägt nur `slice-046` nach; `slice-047` ist die einzige Datei neben der Roadmap in `in-progress/` und dort nirgends erwähnt.
- `verifizierbar`: ja — `make doc-planning`. **Nachgeprüft: läuft grün (0 Befunde)** — die Verifikations-Behauptung trägt also nicht, der Beobachtungsteil bleibt gültig.

### F-5 — Pauschale „Default gewinnt"-Klausel kollidiert mit als permanent deklarierten `MR-*`

- `kategorie`: MEDIUM
- `quelle`: `MR-002`, `MR-004` (beide „Auflösungs-Trigger: permanent") / `harness/conventions.md` §Purpose
- `pfad`: `harness/conventions.md:29-31`
- `befund`: §Baseline setzt ohne Geltungsbereichs-Einschränkung „bei Konflikt … gewinnt der Default", während `MR-002`/`MR-004` als permanente Adaptionen daneben stehen und über die `ids`-Muster in `.d-check.yml` **maschinell gegatet** sind. Für das heute gültige ID-Schema nennt die Datei zwei einander ausschließende Regeln.
- `verifizierbar`: nein (Widerspruch zweier normativer Sätze)

### F-6 — Der vom Baseline-Default geforderte Freshness-Audit fehlt; die Etappen-Zuordnung widerspricht ihm

- `kategorie`: MEDIUM
- `quelle`: `modul-02-harness-bootstrap.md` §Freshness-Audit
- `pfad`: `docs/plan/planning/in-progress/slice-047-baseline-vendoring.md:57-61`
- `befund`: Modul 02 führt den Audit als „**Netz-Operation, außerhalb der Gates** … Wartung, kein Feedback-Gate" und verlangt „Release-*Liste* prüfen, nicht das Asset". Der Plan ordnet genau diese Prüfung als „Kandidat für `gate-consistency`" (netzlos) der Etappe D zu; ein Auslöser ist nirgends deklariert.
- `verifizierbar`: nein (Lesevergleich)

### F-7 … F-13 (LOW)

| # | Kurztitel | `pfad` | `befund` (gekürzt) |
|---|---|---|---|
| F-7 | „adoptierter Stand" (Kurs-Welle) nicht notiert | `harness/conventions.md:23-42` | Das Template verlangt neben dem Tag die Stand-Zeile („Kurs-Welle 34 · 2026-07-24"); §Baseline nennt nur den Tag |
| F-8 | Sensors-Kommentar zeigt auf Upstream-Pfad | `harness/README.md:62` | Begründet die Spaltenform mit `lab/templates/harness/README.template.md` — diesen Pfad gibt es im Repo nicht; die Datei liegt jetzt vendored |
| F-9 | Report-Konvention nicht nachgezogen | `docs/reviews/README.md:3-13` | Beschreibt das Output-Schema ohne die Kopf-Metadaten der Vorlage (Review-Art, Skill-Version, Modell) |
| F-10 | Lesepflicht ohne Auswahlregel | `AGENTS.md:18-26` | „den zur Aufgabe gehörenden Abschnitt" ohne Zuordnung Aufgabe→Abschnitt; die Baseline rahmt das Regelwerk als Nachschlag **pro Entscheidung**, nicht als Pro-Session-Pflicht |
| F-11 | `MR-006` §Geltungsbereich nennt `.d-check.yml` nicht | `harness/conventions.md:199-217` | Die Adaption begründet dort die `scan.ignore`-Änderung, führt die Datei aber nicht im Geltungsbereich |
| F-12 | Erst-Report benennt die Diff-Range falsch | `docs/reviews/2026-07-25-slice-047-baseline-vendoring.md:9` | „`origin/main...HEAD` (1 Commit)" — Range und Commit-Zahl bezeichnen verschiedene Mengen |
| F-13 | Zeilenverweis im Erst-Report zeigt nicht auf den Inhalt | `…-baseline-vendoring.md:54` | `pfad` war `.d-check.yml:6` (Modul-Aufzählung); der beschriebene Kommentar steht in 16–18 |

## Negativbefunde

- geprüft, ohne Befund: **Integritätsmanifest** — `sha256sum -c --quiet SHA256SUMS` grün; 42 Einträge, 42 Inhaltsdateien, Pfadmengen identisch (kein Extra-, kein Fehl-Eintrag).
- geprüft, ohne Befund: **Struktur** — 21 `regelwerk/`-Abschnitte (README + 3 Grundlagen + `modul-00`…`modul-16`) und 21 Templates; die Angaben „43 Dateien" und „17 Module + drei Grundlagen" stimmen mit dem Baum.
- geprüft, ohne Befund: **Netzlose Auflösung** — alle relativen Links in `regelwerk/**` inkl. sämtlicher `../templates/…`-Verweise lösen lokal auf. Die 26 nicht auflösenden liegen ausschließlich in `templates/**` und sind Platzhalter — damit ist der Kern-Zweck von `MR-006` belegt **und** der `scan.ignore` sachlich unterlegt (das `links`-Modul würde genau diese 26 melden).
- geprüft, ohne Befund: **Breite des `scan.ignore`** — auf `.harness/baseline/**` begrenzt, nicht `.harness/**`; `.harness/skills/**`, `docs/**`, `spec/**`, `harness/**` bleiben im Scan. Kein eigenes Dokument liegt unter `.harness/baseline/`.
- geprüft, ohne Befund: **Verweis-Ziele der drei Pin-Stellen** — alle neuen Links existieren; der `MR-006`-Anker leitet sich korrekt aus der Überschrift ab.
- geprüft, ohne Befund: **`MR-006`-Pflichtfelder** — Datum, Geltungsbereich, Adaption, Begründung, Auflösungs-Trigger vorhanden; Inhalt deckungsgleich mit dem Template-Default (`MR-003`), Abweichung ist ausschließlich die Nummer und ist begründet, nicht still gesetzt.
- geprüft, ohne Befund: **Source-Precedence-Konsistenz** — `harness/README.md` und `AGENTS.md` §2 identisch (9 Ränge); `MR-003`-Verweis intakt; die Baseline führt als Default zwei Spec-Ränge, `MR-001` bleibt eine echte Adaption.
- geprüft, ohne Befund: **`v1.3.0`-Reste** — historisch legitim in `slice-046/047`, `roadmap.md`, `conventions.md` (Adoptions-Datum, `MR-000`, `MR-006`-Vorzustand), älteren Reviews. Aktuell-behauptend und als Finding erfasst: nur `reviewer.md:7` (F-1) und `.d-check.yml:4` (F-2).
- geprüft, ohne Befund: **Produkt-Achse** — kein Diff in `spec/**`, `internal/**`, `cmd/**`, `Makefile`, `tools/**`, `.github/workflows/**`, `a-check.mk`, `version.md`, `CHANGELOG.md`; keine neue ID, kein Bump.
- geprüft, ohne Befund: **Nachweis-Mechanik** — `tools/harness/working-tree-hash.sh` hasht ohne Pfad-Ausnahme; der vendored Baum geht in `gates-passed.diffsha` ein (INFO des Erst-Reports mechanisch korrekt).
- geprüft, ohne Befund: **Modul-10-Konformität des Erst-Reports** — Kopf-Metadaten, Feld-Schema, Negativbefunde, Summary und begründetes Verdikt vorhanden; Summary-Zahlen stimmen mit den Findings.
- nicht verifizierbar, daher kein Befund: Erst-Report-F-2 („roter `doc-check` committet") — Prozess-Aussage über den Lauf-Verlauf, aus dem Repo-Zustand weder bestätigbar noch widerlegbar; nicht bestritten.

## Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 0 |
| MEDIUM | 6 |
| LOW | 7 |
| INFO | 0 |

## Verdikt

**Merge-blockierend: ja** (Stand des Reviews). Sechs MEDIUM offen:

- **Ohne Diskussion blockierend:** F-1, F-2, F-3, F-4 — Sync-Drift bzw. Evidenz-Fehler mit reinem
  Textumfang, ohne Abhängigkeit von einer späteren Etappe. F-3 trifft den Kern der Repo-Norm „kein
  Erfolg ohne echte Gate-Ausgabe": die Zahl, die als *einziger* Beleg für den `scan.ignore` dient,
  ist in drei Artefakten desselben Branches unterschiedlich.
- **Blockierend, weil undeklariert:** F-5, F-6 — als *benannt vertagte* Punkte wären sie akzeptabel;
  heute stehen sie als Widerspruch (F-5) bzw. mit einer der Baseline widersprechenden Zuordnung
  (F-6) im Repo.
- **Kein HIGH:** Die Substanz von Etappe A ist in Ordnung — Manifest grün und vollständig, Baum wie
  deklariert, netzlose Auflösung nachgewiesen, `MR-006` sauber, `scan.ignore` nicht zu breit,
  Produkt-Achse unberührt.

## Auflösung (2026-07-25, nach dem Review)

| # | Auflösung |
|---|---|
| F-1 | `.harness/skills/reviewer.md:7` auf `v3.5.2` + vendored Modul-Pfad umgestellt |
| F-2 | `.d-check.yml`-Kopf auf „Baseline v3.5.2 (vendored, `MR-006`)" |
| F-3 | Plan-§3 neu formuliert: keine absolute Zahl als Beleg, sondern die Aussage „die Dateizahl wächst nur um die neuen **Eigen**-Dokumente, nicht um die 43 vendored" |
| F-4 | Roadmap nennt `slice-047` jetzt als „In Arbeit". **Zusatzbefund:** `make doc-planning` läuft auch **ohne** diesen Eintrag grün — das Gate deckt den Fall nicht ab, die Behauptung „verifizierbar: ja" war zu stark |
| F-5 | Klausel eingegrenzt: der Default gewinnt für Adaptionen, die nur wegen des alten Defaults existieren; **bis zur Prüfung in Etappe C bleiben die deklarierten `MR-*` in Kraft** (mit Hinweis auf die `ids`-Gates) |
| F-6 | Plan-§3 übernimmt die Trennung aus Modul 02: **Freshness = Netz-Wartung außerhalb der Gates** (Release-Liste, nicht Asset), **Asset-Integrität** = Gate-Kandidat via d-check-`sources`; Auslöser-Deklaration nach Etappe C/D verwiesen |
| F-7…F-11 | **offen** — nach Etappe C/D verwiesen (Stand-Zeile, Upstream-Pfad im Kommentar, `docs/reviews/README.md`-Form, Lesepflicht-Auswahlregel, `MR-006`-Geltungsbereich); LOW, nicht blockierend |
| F-12/F-13 | im Erst-Report korrigiert (Erratum, siehe dessen Nachtrag) |

**Damit sind alle sechs MEDIUM abgearbeitet**; die sieben LOW sind benannt und zugeordnet. Das
blockierende Verdikt ist damit aufgehoben — nicht durch Widerspruch, sondern durch Behebung.
