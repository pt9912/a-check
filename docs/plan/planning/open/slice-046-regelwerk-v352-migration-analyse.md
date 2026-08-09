# slice-046 — Regelwerk-Migration v1.3.0 → v3.5.2: Delta-Analyse

**Status:** open — **Analyse zur Abnahme** (keine Spec-/Code-/Harness-Änderung; reine Ist-Messung
+ Etappen-Vorschlag).
**Auslöser:** Maintainer-Frage am 2026-07-25 („welche Regelwerksversion verwenden wir?") und die
anschließende Vorgabe: **vollständige Migration nach `v3.5.2`** — auch dort, wo ein `MR-*` heute
etwas anderes sagt. **Erst analysieren.**
**Bezug:** betrifft die Baseline-Deklaration in
[`harness/conventions.md` §Baseline](../../../../harness/conventions.md#baseline), das Briefing
[`AGENTS.md`](../../../../AGENTS.md) §1 und den Harness-Einstieg
[`harness/README.md`](../../../../harness/README.md). [Roadmap](../in-progress/roadmap.md).

> **Hinweis:** Analyse zur Abnahme. Es werden hier **keine** `AC-*`/`ADR-*`/`MR-*`-IDs vergeben und
> keine Artefakte geändert. Die Etappen (§6) gehören **vor** die Umsetzung abgenommen.

---

## 1. Ist-Stand

| | |
|---|---|
| a-check gepinnt auf | **`v1.3.0`** (Adoption 2026-06-20, „von Beginn an gepinnt") |
| aktuelles Kurs-Release | **`v3.5.2`** (Stand „Kurs-Welle 34", 2026-07-24) |
| dazwischen | 11 Releases, **zwei Major-Sprünge** (v1 → v2 → v3) |
| Flotte | d-check: `v1.4.0` · ai-harness-init: vendored `v3.5.1` · b-cad/d-migrate/m-trace/belief-agent: keine Deklaration |

Der Pin steht an **drei** Stellen: `harness/conventions.md` §Baseline (kanonisch),
`AGENTS.md` §1 (ZIP-URL + Quelldatei-URL), `harness/README.md` §Guides (Tabellenzeile).

## 2. Umfang des Sprungs (gemessen)

`git diff --stat v1.3.0 v3.5.2 -- lab/regelwerk`: **alle 21 Dateien geändert**, **+705 / −1868
Zeilen**. Das Regelwerk wurde **verdichtet**, nicht nur ergänzt — die Annahme „additive Ergänzungen,
unser Stand bleibt gültig" trägt nicht.

**Gelesen wurde das kanonische Artefakt** — `lab-regelwerk.zip` aus dem `v3.5.2`-Release (125 KB),
nicht der Git-Baum: das Bundle ist laut eigenem README die ausgelieferte, self-navigierbare Form
und trägt `regelwerk/` **und** `templates/` **parallel**.

## 3. Was sich inhaltlich **nicht** gedreht hat (Entwarnung)

Geprüft, weil heute frisch darauf gebaut: **Modul 10 (Review-Harness)** führt in `v3.5.2`
unverändert die Kategorien **`HIGH`/`MEDIUM`/`LOW`/`INFO`**, das Output-Schema mit `verifizierbar`
und die **Negativbefund-Zeile** je betrachtetem Bereich. Die drei Review-Synthesen vom 2026-07-25
und [`docs/reviews/README.md`](../../../reviews/README.md) stehen damit **nicht** auf einer
überholten Fassung.
*Neu und schärfer:* die repo-konkrete HIGH-Liste muss **mindestens zwei repo-spezifische Regeln**
nennen — a-checks `.harness/skills/reviewer.md` ist daraufhin zu prüfen.

## 4. Die drei echten Migrations-Brocken

### 4.1 Vendoring-Modell — strukturell neu

`v3.5.2` erwartet die Baseline **committet vendored** unter
`.harness/baseline/<tag>/{regelwerk,templates}/`, materialisiert aus dem ZIP, mit
`SHA256SUMS` als Integritätsmanifest. Der Zweck: die `../templates/…`-Verweise des Regelwerks
lösen **netzlos lokal** auf, und der Nachschlag ist über den `<tag>` reproduzierbar.

a-check referenziert die Baseline heute **nur per URL** (`AGENTS.md` §1) — das v1.3.0-Modell.
Real-Beleg für das neue Modell in der Flotte: `ai-harness-init` führt `.harness/baseline/v3.5.1/`.

**Konsequenz:** Die Migration ist keine Versionsnummer, sondern das **Materialisieren eines
Verzeichnisses** plus Integritätsnachweis — und die Frage, ob dessen Aktualität gegatet wird
(Kandidat für `gate-consistency`, analog zur bestehenden Pin-Konsistenz).

### 4.2 `MR-*`-Kollision — und die Vorgabe, sie aufzulösen

Das `conventions.template.md` von `v3.5.2` **schreibt eine MR-Nummer vor**:

> `### MR-003 — Regelwerk und Templates als vendored, nachschlagbare Baseline` <!-- d-check:ignore (Zitat der Template-Ueberschrift, kein Verweis auf a-checks MR-003) -->

a-checks **[MR-003](../../../../harness/conventions.md#mr-003--source-precedence-ohne-docsuser-rang)** ist belegt — „Source Precedence ohne `docs/user`-Rang", seit 2026-06-21
**aufgelöst**. Damit steht die Migration vor einer Nummern-Kollision. Die Maintainer-Vorgabe ist
eindeutig: **vollständig migrieren, auch gegen bestehende `MR-*`**. Zu entscheiden bleibt das *Wie*
(§6, Etappe C): Umnummerierung des Bestands, Aufgeben der Nummern-Identität oder Übernahme der
Template-Nummer mit dokumentiertem Bruch. Betroffen sind auch die **Verweise** auf [MR-003](../../../../harness/conventions.md#mr-003--source-precedence-ohne-docsuser-rang)
(`harness/README.md` §Source precedence) und die `d-check`-`ids`-Linkpflicht auf `MR-\d{3}`.

Der **gesamte MR-Bestand** ist gegen die neuen Defaults zu prüfen — sechs Einträge:

| MR | Gegenstand | Erste Einschätzung gegen v3.5.2 |
|---|---|---|
| [MR-000](../../../../harness/conventions.md#mr-000--baseline-aussage-inkl-id-schema-deklaration) | Baseline-Aussage inkl. ID-Schema | Template kennt den Abschnitt weiterhin — Inhalt abgleichen |
| [MR-001](../../../../harness/conventions.md#mr-001--source-precedence-mit-eigener-spezifikations-schicht) | Source Precedence mit eigener Spezifikations-Schicht | prüfen, ob v3.5.2 die drei Spec-Straten inzwischen als Default führt (dann entfällt die Adaption) |
| [MR-002](../../../../harness/conventions.md#mr-002--id-schema-mit-bereichskürzeln-ab-initialer-fassung) | ID-Schema mit Bereichskürzeln | dito |
| [MR-003](../../../../harness/conventions.md#mr-003--source-precedence-ohne-docsuser-rang) | Source Precedence ohne `docs/user`-Rang (**aufgelöst**) | **Nummern-Kollision** mit dem Template |
| [MR-004](../../../../harness/conventions.md#mr-004--spezifikation-und-architektur-strata-und-id-schemata) | Spezifikation/Architektur: Strata + ID-Schemata | dito [MR-001](../../../../harness/conventions.md#mr-001--source-precedence-mit-eigener-spezifikations-schicht) |
| [MR-005](../../../../harness/conventions.md#mr-005--referenzmatrix-intra-spec-richtung--adrslice-disziplin-d-check-angleichung) | Referenzmatrix (d-check-Angleichung) | repo-spezifisch, vermutlich unberührt |

### 4.3 Template-Abgleich — 21 Vorlagen gegen handgeschriebene Artefakte

Gemessen (Zeilen Template/a-check):

| Template | a-check | Bemerkung |
|---|---|---|
| `.harness/skills/closure-note-reviewer.template.md` | **fehlt ganz** | 98 Zeilen, kein Gegenstück im Repo |
| `docs/reviews/review-report.template.md` (76) | 34 | a-checks README ist eine Konvention, **kein Report-Template**; die Vorlage verlangt Felder, die heute fehlen: **Review-Art**, **Skill-Version**, **Modell-ID** |
| `harness/conventions.template.md` (223) | 225 | fehlende Abschnitte: `## Glossar (optional)`, [MR-003](../../../../harness/conventions.md#mr-003--source-precedence-ohne-docsuser-rang)-Vendoring (§4.2) |
| `harness/README.template.md` (141) | 129 | strukturell **vollständig** — keine fehlenden Abschnitte |
| `AGENTS.template.md` (179) | 163 | nur Überschriften-Varianten (`3.1 Docker-only` vs. `Docker/make-only`) — inhaltlich zu prüfen |
| `spec/*.template.md` (96–124) | 145–619 | a-checks Spec-Straten sind weit gewachsen; Template-Konformität nur auf **Struktur**, nicht auf Umfang zu prüfen |
| `Makefile`, `.d-check.yml`, `project-readme` | größer | repo-spezifisch gewachsen, Abgleich vermutlich unkritisch |
| `docs/plan/**`-Vorlagen (Slice, Welle, Carveout, ADR, Roadmap) | teils vorhanden | Slice-/Welle-/Carveout-Vorlagen hat a-check **nicht** als Datei — Form ist implizit gelebt |

## 5. Was diese Analyse **nicht** geklärt hat

- **Modul-für-Modul-Delta:** nur Modul 10 wurde inhaltlich gegengelesen (weil heute darauf gebaut).
  Die übrigen 20 Dateien sind nur im Umfang vermessen. Module mit hoher a-check-Relevanz stehen aus:
  **05** (Planning/Slice-Lifecycle), **04** (ADR-Regeln, Immutabilität), **13** (Quality Gates),
  **02** (Bootstrap/Vendoring), **08/11** (Rollen, Verifikation).
- **Ob `v3.5.2`-Defaults bestehende `MR-*` überflüssig machen** (§4.2, Spalte „Erste Einschätzung")
  — das ist Lesearbeit, keine Messung.
- **Ob die Migration Gate-Änderungen erzwingt** (z. B. Vendoring-Aktualität in `gate-consistency`).

## 6. Vorschlag: vier Etappen

| Etappe | Inhalt | Warum getrennt |
|---|---|---|
| **A — Vendoring** | `.harness/baseline/v3.5.2/{regelwerk,templates}/` + `SHA256SUMS` materialisieren; `AGENTS.md` §1, `harness/README.md`, `conventions.md` §Baseline auf den vendored Pfad + neuen Stand umstellen | mechanisch, sofort prüfbar, **Voraussetzung** für alles Weitere (netzloser Nachschlag) |
| **B — Modul-Delta lesen** | die fünf relevanten Module gegenlesen und die Treffer als Findings sammeln | Lesearbeit; Ergebnis bestimmt C und D |
| **C — `MR-*`-Bereinigung** | Kollision [MR-003](../../../../harness/conventions.md#mr-003--source-precedence-ohne-docsuser-rang) auflösen, Bestand gegen die neuen Defaults durchgehen, überflüssige Adaptionen streichen (Vorgabe: Default gewinnt) | berührt die Konventions-Identität des Repos — eigener, reviewbarer Schritt |
| **D — Template-Konformität** | fehlende Artefakte anlegen (`closure-note-reviewer`, Review-Report-Vorlage), Struktur-Abweichungen schließen | breit, aber mechanisch; nach C, weil C die Form mitbestimmt |

**Reihenfolge-Argument:** A zuerst, weil danach jeder folgende Schritt **netzlos** gegen die
vendored Quelle arbeitet statt gegen eine URL — genau die Eigenschaft, die das Vendoring begründet.

## 7. DoD dieser Analyse

- [x] Ist-Stand und Pin-Stellen erhoben (§1).
- [x] Sprung-Umfang gemessen, kanonisches Artefakt (ZIP) verwendet (§2).
- [x] Modul 10 gegengelesen — die frischen Review-Artefakte sind nicht überholt (§3).
- [x] Die drei Migrations-Brocken benannt und belegt (§4), Template-Abgleich tabellarisch.
- [x] Offene Lücken der Analyse ausgewiesen (§5) statt als Vollständigkeit ausgegeben.
- [x] **Abnahme:** Etappen-Schnitt A–D (§6) durch den Maintainer — erteilt am 2026-07-25, belegt
      durch die ausgeführten Etappen (siehe §8). **Retroaktiv festgehalten am 2026-08-09**: der
      Haken blieb offen, während die Arbeit lief, und hielt diesen Slice zwei Wochen in `open/`.

## 8. Closure-Notiz

**Geliefert:** die Delta-Analyse, die die Regelwerk-Migration ausgelöst und geschnitten hat. Der
vorgeschlagene Schnitt wurde angenommen — aber **nicht unverändert**.

**Der abgenommene Schnitt weicht vom vorgeschlagenen ab.** §6 schlug **vier** Etappen vor (A–D);
umgesetzt wurden **sechs**, mit **E vor D**:

| Etappe | Umsetzungs-Slice |
|---|---|
| A — Vendoring | [slice-047](../done/slice-047-baseline-vendoring.md) |
| B — Modul-Delta lesen | [slice-048](../done/slice-048-modul-delta-lesen.md) |
| E — Mechanik/Sensoren *(neu, vorgezogen)* | [slice-049](../done/slice-049-mechanik-sensoren.md) … [slice-051](../done/slice-051-workflow-und-freigabe.md) |
| D — Form/Templates | [slice-052](../done/slice-052-slice-form.md) … [slice-054](../done/slice-054-ac-form.md) |
| C — `MR-*`-Bereinigung | [slice-055](../done/slice-055-mr-bestand.md), [slice-056](../done/slice-056-sub-area-modus.md) |
| F — Steering/Closure *(neu)* | [slice-057](../done/slice-057-steering-loop.md), [slice-065](../done/slice-065-carveout-ort-und-trichter.md), [slice-066](../done/slice-066-wellen-closure-und-rollen.md) |

Der Grund für die Erweiterung steht im Drift-Log der [Roadmap](../in-progress/roadmap.md): die
vollständige Baseline-Lektüre in Etappe B brachte elf zusätzliche Funde, und die Mechanik-Funde
schufen erst die Sensoren, an denen die Form-Funde hingen — deshalb E vor D.

**Die in §5 ausgewiesenen Lücken sind geschlossen**, teils anders als erwartet: das
Modul-für-Modul-Delta lieferte [slice-048](../done/slice-048-modul-delta-lesen.md) mit 21 Funden
statt der erwarteten fünf Module; die `MR-*`-Frage entschied
[slice-055](../done/slice-055-mr-bestand.md) (die vermutete Kollision wurde **widerlegt**); und
Gate-Änderungen erzwang die Migration sehr wohl — `suppression-check`, `verify`,
`commit-scope-check` und der Drift-Wächter entstanden erst dadurch.

**Lerneintrag — Form: geschärfte Regel.** Als Prüfsatz: *Ein DoD-Punkt, der auf eine fremde
Handlung wartet, wird nicht durch das Eintreten der Handlung erfüllt, sondern durch ihr
Festhalten.*

**Die Ursache** ist beobachtbar und war zwei Wochen sichtbar, ohne aufzufallen: die Abnahme
erfolgte mündlich und wurde sofort in Arbeit umgesetzt — die Etappen liefen, der Haken blieb. Kein
Sensor konnte es melden, weil ein offener DoD-Punkt in `open/` der **Normalzustand** ist; erst der
Widerspruch zum Closure-Log der Welle („alle Slices in `done/`") machte ihn sichtbar, und den fand
der unabhängige Review als `F-8`, nicht ein Gate.

**Zwei beobachtbare Closure-Kriterien:**

1. Jede Etappe aus der Tabelle oben zeigt auf mindestens einen Slice in `done/`; keine Zeile ist
   ohne Umsetzung.
2. `docs/plan/planning/open/` enthält nach diesem Slice keinen Slice der Welle `welle-12` mehr —
   damit ist das Closure-Kriterium „alle Slices in `done/`" erstmals tatsächlich erfüllt statt
   behauptet.

**Folge-Slices:** keine. Der Abschluss von `welle-12` nach der Fünf-Schritt-Prozedur ist der
nächste Zug — und deren erste reale Probe.
