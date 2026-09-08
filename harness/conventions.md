# Harness-Konventionen

## Purpose

Diese Datei deklariert die *repo-lokalen* Strukturregeln dieses Repos
gegenüber der adoptierten Harnesskonvention (Baseline):

- **Adaptionen** ggü. der Baseline (mit Begründung und Auflösungs-Trigger).
- **ID-Schema-Deklaration** — welches Präfix-Schema dieses Repo nutzt
  ([`MR-000`](#mr-000--baseline-aussage-inkl-id-schema-deklaration)).
- **Zusatzklassen-Deklarationen** für die Sensors-Bindung.
- **Modus-Deklarationen** pro Sub-Area (Greenfield / Brownfield /
  Hybrid) inklusive Konvergenz-Auftrag bei BF.

Bei Konflikt zwischen dieser Datei und einer kanonischen Quelle gilt
die kanonische Quelle (Source Precedence, siehe
[`README.md`](README.md)). Diese Datei ist konformitätsbringend für
*Form*-Fragen, nicht autoritativ über Inhalt.

## Baseline

- **Konvention:** AI-Harness-Kurs
- **Stand:** [`v6.5.0`](https://github.com/pt9912/ai-harness-course/releases/tag/v6.5.0)
  (Release-Tag) — **Kurs-Welle 128 · 2026-09-06**, wie im Kopf des vendored
  [`regelwerk/README.md`](../.harness/baseline/v6.5.0/regelwerk/README.md) ausgewiesen.
- **Ort:** **committet vendored** unter
  [`.harness/baseline/v6.5.0/`](../.harness/baseline/v6.5.0/regelwerk/README.md), Integrität
  über `SHA256SUMS`, geprüft mit `make regelwerk-check`
  ([`MR-006`](#mr-006--baseline-committet-vendored-statt-per-url-referenziert)). Genau **ein**
  Stand liegt vendored; mehrere sind nur während einer Migration zulässig, und das Target weist
  den ungeprüften dann namentlich aus. Die Zusage gilt **ohne Ausnahme**, auch für das Feld
  `Ersetzt-Baseline-Regel` **aktiver** Einträge: es nennt eine **Regel**, keine
  Datei-Kopie. Sein Zeiger wandert darum beim Baseline-Wechsel mit — Bedingung ist, dass der
  referenzierte Abschnitt im neuen Stand wortgleich ist, und das ist zu **messen**, nicht
  anzunehmen ([slice-172](../docs/plan/planning/done/wellenlos/slice-172-baseline-v600-entfernen.md)
  §2.2). Ist er es nicht, trägt die Stelle die Abweichung sichtbar, statt still umzuziehen.
  Durchgesetzt seit [slice-173](../docs/plan/planning/done/wellenlos/slice-173-versions-sensor-baseline-pins.md)
  durch das `versions`-Muster in [`.d-check.yml`](../.d-check.yml).
- **Aufgelöste Einträge sind davon ausgenommen** ([`conventions/done/`](conventions/done/)): sie
  sind eingefroren, ihr Zeiger bleibt auf dem Stand, gegen den sie damals formuliert wurden. Der
  Sensor nimmt sie aus. **Die Kollision, die hier stand, ist bezahlt:** Solange ein eingefrorenes
  Artefakt die Baseline als **Link** zitierte, erzwang das Entfernen eines Stands einen Edit an
  ihm — bei [`MR-018`](#mr-018) mit slice-172 geschehen, 22 Nachzüge insgesamt. Seit slice-176
  zitieren einfrierende Artefakte als **Kennung** statt als Adresse
  ([`AGENTS.md`](../AGENTS.md) §5), und derselbe Vorgang kostete **null** Nachzüge — gemessen, nicht
  angenommen. Was bleibt, ist die *Versions*-Frage: Ein Zeitdokument nennt den alten Stand
  weiterhin im Text, weil damals er galt, und dafür ist `exempt-paths` da. Zwei Fragen, zwei
  Antworten.
- **Adoptiert seit:** 2026-06-20.
- **Beim Heben des Stands: Delta *und* Voll-Abgleich** (`seit slice-185`). Die Delta-Analyse zwischen zwei Ständen
  findet, was sich **ändert** — nicht, was seit der Adoption fehlt. Gemessen an einem Beispiel
  (slice-185): Der Arbeitsteilungs-Satz für [`AGENTS.md`](../AGENTS.md) §4 steht seit `v5.12.0`
  unverändert in der Ziel-Form, tauchte in **keinem** der vier Deltas auf und war nie übernommen.
  Darum gehört zu jedem Sprung ein **Voll-Abgleich** der Ziel-Formen gegen ihr Gegenstück im
  Repo: die **15** Vorlagen mit genau einem Gegenstück, Abschnitt für Abschnitt. Er ist
  **Lese-Arbeit mit maschineller Vorauswahl**, kein Lauf — ein Wortfolgen-Vergleich meldet
  überwiegend Platzhalter und Bedienhinweise, die beim Kopieren bestimmungsgemäß verschwinden.
  Die elf **Instanz**-Vorlagen (Slice, ADR, Report, …) fallen heraus, ebenso
  `templates/README.md` — der **Index** des Vorlagen-Verzeichnisses ist keine Ziel-Form; ihre Form prüft
  `make doc-structure` über Muster.

Wann welcher Stand gehoben wurde und in welchen Etappen, steht in
[`docs/plan/planning/done/`](../docs/plan/planning/done/) — nicht hier. Diese Datei trägt den
Ist-Zustand.

## Adoptierte Konventions-Quellen

Pointer, keine Wiederholung des Inhalts.

- **Vendored Baseline (Regelwerk + Templates) — die Lese-Form:**
  [`.harness/baseline/v6.5.0/regelwerk/README.md`](../.harness/baseline/v6.5.0/regelwerk/README.md)
  (Index) und
  [`.harness/baseline/v6.5.0/templates/README.md`](../.harness/baseline/v6.5.0/templates/README.md).
  **Netzlos** auf jedem Checkout, pro Abschnitt eine Datei — ein Agent lädt den benötigten
  Abschnitt, nie das ganze Bundle
  ([`MR-006`](#mr-006--baseline-committet-vendored-statt-per-url-referenziert)).
- **In-Repo (verkörperte Form):** [`AGENTS.md`](../AGENTS.md),
  [`harness/README.md`](README.md), diese Datei und die Vorlagen unter
  [`docs/plan/`](../docs/plan/planning/README.md). Referenz-Form sind die vendored
  `templates/`; die eigenen Dateien sind daraus ausgefüllt.

## Adaptions-Block

**Disziplin** (aus dem Konventions-Template der Baseline): Einträge sind
**chronologisch pro Repo** nummeriert und tragen die Pflichtfelder Datum,
Geltungsbereich, Adaption, Begründung, Auflösungs-Trigger. An einem
akzeptierten Eintrag wird **nichts nachträglich inhaltlich geändert** —
Korrekturen entstehen als neuer `MR` oder als ausdrückliche Aufhebung,
analog zur ADR-Immutabilität ([`AGENTS.md`](../AGENTS.md) §3.5).


### MR-000 — Baseline-Aussage (inkl. ID-Schema-Deklaration)

<a id="mr-000"></a>

- **Datum:** 2026-06-20
- **Geltungsbereich:** gesamtes Repo
- **Adaption:** keine inhaltlichen Adaptionen ggü. Baseline-Default für
  Verzeichniskonvention, Lifecycle-Regeln, Carveout-Disziplin. Spätere
  Adaptionen werden als `MR-<NNN>` nachgetragen.
- **ID-Schema-Deklaration** (vom Konventions-Template als Teil der
  Baseline-Aussage vorgesehen — hier von Beginn an gesetzt):
  - Funktionale Anforderungen: `AC-FA-<BEREICH>-<NNN>` (Bereichskürzel,
    siehe [`MR-002`](#mr-002--id-schema-mit-bereichskürzeln-ab-initialer-fassung));
    Bereiche initial `RULE`/`EXTRACT`/`CLI`/`CONF`/`DIST`.
  - Nichtfunktionale Anforderungen: `AC-QA-<NN>`.
  - ADRs: `ADR-NNNN` (vierstellig, gemäß Kurs-ADR-Vorlage `v1.3.0`).
  - Konventions-Adaptionen: `MR-NNN`. Carveouts: `CO-NNN` (bisher
    ungenutzt). Slices: `slice-NNN`.
- **Begründung:** Initial-Setzung. Eine undeklarierte ID-Systematik wäre
  eine stille Setzung (gleiche Harness-Lüge-Klasse wie ein undeklariertes
  Gate); deshalb steht das Schema von Anfang an hier — gelernt aus dem
  Schwester-Repo `d-check`, wo es als Nachtrag (`MR-008`) ergänzt werden
  musste.
- **Auflösungs-Trigger:** permanent.

### Aktive Adaptionen

Eine Zeile je Datei in [`harness/conventions/`](conventions/). Geltungsbereich und ersetzte
Baseline-Regel stehen hier, damit ein Agent **ohne Öffnen** entscheiden kann, ob der Eintrag ihn
betrifft.

Jede Zeile trägt die stabile Kennung `mr-<NNN>` als Anker — die Adresse, unter der andere Dateien
referenzieren. Wo ein Eintrag vor slice-096 unter seinem Überschriften-Slug veröffentlicht wurde,
steht dieser **zusätzlich** daneben; sonst rotten die damaligen Verweise. Gemessen tragen das
heute die Zeilen der **aufgelösten** Tabelle weiter unten, keine der aktiven — dort ist kein
Eintrag alt genug.

Die Spalte *Ersetzt-Baseline-Regel* ist das Pflichtfeld des neuen Stands. Sie kann in einen
akzeptierten Eintrag **nicht nachgetragen** werden (Einträge werden nie überschrieben); sie
entsteht in den Nachfolge-Einträgen. Gemessen tragen heute **fünf** der sieben aktiven Zeilen
einen Zeiger, zwei ein `—` mit Begründung in der Zelle.

| MR | Titel | Geltungsbereich | Ersetzt-Baseline-Regel |
|---|---|---|---|
| [MR-012](conventions/MR-012-referenzmatrix-grandfathering.md) <a id="mr-012"></a> | Referenz-Richtung maschinell, ADRs 0001–0020 grandfathered | [`.d-check.yml`](../.d-check.yml) (`matrix`), [`docs/plan/adr/`](../docs/plan/adr/) | [`grundlagen-referenz-richtung.md` §Referenz-Richtung (SDP)](../.harness/baseline/v6.5.0/regelwerk/grundlagen-referenz-richtung.md#referenz-richtung-sdp-wer-darf-wen-referenzieren) |
| [MR-014](conventions/MR-014-keine-agenten-telemetrie.md) <a id="mr-014"></a> | Keine Agenten-Telemetrie | gesamtes Repo; Baseline-Modul `modul-15` | [`modul-15-observability.md` §Kernidee](../.harness/baseline/v6.5.0/regelwerk/modul-15-observability.md#kernidee-modul-15) |
| [MR-015](conventions/MR-015-welle-closure-ohne-replay.md) <a id="mr-015"></a> | Welle-Closure ohne Replay-Lauf (`make ci` grün) | [`docs/plan/planning/`](../docs/plan/planning/README.md) | [`modul-06-roadmap.md` §Wellen-Closure-Prozedur](../.harness/baseline/v6.5.0/regelwerk/modul-06-roadmap.md#wellen-closure-prozedur-modul-6) |
| [MR-016](conventions/MR-016-validator-unbesetzt.md) <a id="mr-016"></a> | Validator-Rolle unbesetzt | gesamtes Repo; Baseline-Modul `modul-08` | [`modul-08-agentenrollen.md` §Die neun Übergaben](../.harness/baseline/v6.5.0/regelwerk/modul-08-agentenrollen.md#die-neun-übergaben-und-ihre-artefakte-modul-8) |
| [MR-019](conventions/MR-019-review-dod-opt-in.md) <a id="mr-019"></a> | Review-DoD-Punkt bleibt Opt-in statt verpflichtend | [`AGENTS.md`](../AGENTS.md) §5, [`.d-check.yml`](../.d-check.yml) | — *(kein Baseline-Regel-Ersatz — der Treiber ist ein Template; Begründung und Rückbau-Bedingung stehen im Eintrag)* |
| [MR-020](conventions/MR-020-adr-vorlage-generisch.md) <a id="mr-020"></a> | ADR-Vorlagen-Referenz zeigt generisch auf den vendorten Stand | [`MR-000`](#mr-000) §ID-Schema, Zeile zu `ADR-NNNN` | — *(korrigiert eine Repo-Aussage, kein Baseline-Regel-Ersatz; permanent, kein Rückbau-Kandidat)* |
| [MR-022](conventions/MR-022-verfeinerungs-form.md) <a id="mr-022"></a> | Verfeinerungen tragen `SPEC-*` statt der Suffix-Form (Begründung gemessen, mit Zählregel) | [`spec/spezifikation.md`](../spec/spezifikation.md) | [`grundlagen-source-precedence.md` §ID-Schema als Klammer](../.harness/baseline/v6.5.0/regelwerk/grundlagen-source-precedence.md#id-schema-als-klammer) |

### Aufgelöste Adaptionen

Eine Zeile je Datei in [`harness/conventions/done/`](conventions/done/) — nur Kennung und
Auflösung, damit die Kette auffindbar bleibt, ohne gelesen zu werden. Der Anker zieht aus der
Tabelle oben mit um; **er** ist der Grund, warum ein Verweis auf eine aufgelöste Adaption nicht
bricht.

| MR | aufgelöst durch |
|---|---|
| [MR-001](conventions/done/MR-001-spezifikations-schicht.md) <a id="mr-001"></a><a id="mr-001--source-precedence-mit-eigener-spezifikations-schicht"></a> | [MR-010](conventions/done/MR-010-rueckbau-drei-adaptionen.md) |
| [MR-002](conventions/done/MR-002-id-schema-bereichskuerzel.md) <a id="mr-002"></a><a id="mr-002--id-schema-mit-bereichskürzeln-ab-initialer-fassung"></a> | [MR-010](conventions/done/MR-010-rueckbau-drei-adaptionen.md) |
| [MR-003](conventions/done/MR-003-source-precedence-ohne-docs-user.md) <a id="mr-003"></a><a id="mr-003--source-precedence-ohne-docsuser-rang"></a> | Ereignis am 2026-06-21 ([Benutzerhandbuch](../docs/user/benutzerhandbuch.md) angelegt), **kein** Nachfolge-Eintrag — die Auflösung datiert vor der Verzeichnis-Form |
| [MR-004](conventions/done/MR-004-spec-strata-id-schemata.md) <a id="mr-004"></a><a id="mr-004--spezifikation-und-architektur-strata-und-id-schemata"></a> | [MR-011](conventions/done/MR-011-verfeinerungs-form.md) |
| [MR-005](conventions/done/MR-005-referenzmatrix.md) <a id="mr-005"></a><a id="mr-005--referenzmatrix-intra-spec-richtung--adrslice-disziplin-d-check-angleichung"></a> | [MR-012](conventions/MR-012-referenzmatrix-grandfathering.md) |
| [MR-006](conventions/done/MR-006-baseline-vendored.md) <a id="mr-006"></a><a id="mr-006--baseline-committet-vendored-statt-per-url-referenziert"></a> | [MR-010](conventions/done/MR-010-rueckbau-drei-adaptionen.md) |
| [MR-007](conventions/done/MR-007-adr-vorlagen-version.md) <a id="mr-007"></a><a id="mr-007--adr-vorlagen-version-v352-statt-v130"></a> | [MR-013](conventions/done/MR-013-adr-vorlagen-version.md) |
| [MR-008](conventions/done/MR-008-kein-replay.md) <a id="mr-008"></a><a id="mr-008--kein-replay-keine-agenten-telemetrie"></a> | [MR-014](conventions/MR-014-keine-agenten-telemetrie.md) |
| [MR-009](conventions/done/MR-009-validator-unbesetzt.md) <a id="mr-009"></a><a id="mr-009--validator-rolle-unbesetzt-zwei-übergaben-ohne-artefakt"></a> | [MR-016](conventions/MR-016-validator-unbesetzt.md) |
| [MR-010](conventions/done/MR-010-rueckbau-drei-adaptionen.md) <a id="mr-010"></a> | — *(Rückbau-Eintrag; mit seiner Entstehung erledigt, siehe Datei)* |
| [MR-011](conventions/done/MR-011-verfeinerungs-form.md) <a id="mr-011"></a> | [MR-021](conventions/done/MR-021-verfeinerungs-form.md) |
| [MR-021](conventions/done/MR-021-verfeinerungs-form.md) <a id="mr-021"></a> | [MR-022](conventions/MR-022-verfeinerungs-form.md) |
| [MR-013](conventions/done/MR-013-adr-vorlagen-version.md) <a id="mr-013"></a> | [MR-017](conventions/done/MR-017-adr-vorlagen-version.md) |
| [MR-017](conventions/done/MR-017-adr-vorlagen-version.md) <a id="mr-017"></a> | [MR-020](conventions/MR-020-adr-vorlage-generisch.md) |
| [MR-018](conventions/done/MR-018-review-pflicht-v610-wortlaut.md) <a id="mr-018"></a> | Ereignis am 2026-09-06 (`v6.2.0` vendored — das Kurs-Template trägt den Rollenwechsel-Absatz jetzt selbst, [slice-170](../docs/plan/planning/done/wellenlos/slice-170-mr018-aufloesen.md)), **kein** Nachfolge-Eintrag |

## Anforderungs-Anlege-Prozess

Neue oder geänderte `AC-*`-Anforderungen entstehen **nur** in
[`spec/lastenheft.md`](../spec/lastenheft.md) (vertraglich,
Change-Request-Charakter — Baseline-Regel der Spec-Stratifizierung;
Rang-Struktur dieses Repos: [`MR-001`](#mr-001--source-precedence-mit-eigener-spezifikations-schicht)).
Pflicht-Bausteine pro Anforderung:

- **ID gemäß Schema-Konvention** im Lastenheft §3
  (`AC-FA-<BEREICH>-<NNN>`, siehe
  [`MR-002`](#mr-002--id-schema-mit-bereichskürzeln-ab-initialer-fassung));
  ein neues Bereichskürzel wird dort in der Schema-Konvention
  deklariert. Nichtfunktionale Anforderungen: `AC-QA-<NN>`.
- **Stabiler Anker.** Die Überschrift trägt darunter einen expliziten
  `<a id="ac-fa-…">` mit der **Kennung** — sonst ist ihr Wortlaut faktisch
  immutabel: verlinkt eine `Accepted`-ADR den generierten Slug, bricht jede
  Umbenennung ihn, und der Nachzug in der ADR verletzt die
  Immutabilität ([`AGENTS.md`](../AGENTS.md) §3.5). Gemessen an
  [`AC-FA-RULE-008`](../spec/lastenheft.md#ac-fa-rule-008): eine Umbenennung erzeugte einen Widerspruch zwischen
  `make doc-check` (Anker müssen auflösen) und `make doc-immutable` (ADRs sind
  unantastbar), den **kein** Weg auflöst — außer diesem. Wird eine Überschrift
  dennoch umbenannt, bleibt der **alte** Slug als zweiter Anker stehen; dieselbe
  Doppelung führen die `MR-`-Kennungen weiter unten seit ihrer Umbenennung.
- **Drei Akzeptanzkriterien** (Happy/Boundary/Negative im
  Given/When/Then-Stil) plus explizite **Out-of-Scope**-Liste.
  **Maschinell geprüft ist davon nur, dass die vier Bausteine benannt sind**
  (`make doc-structure`, ab slice-054, seit slice-120 im Modul `structure`); ob ein Kriterium tatsächlich im
  Given/When/Then-Stil formuliert ist, bleibt **Review-Sache** — dieselbe Grenze
  wie bei „höchstens zwei Schichten" der Slice-Form. Bis slice-076 las sich diese
  Zeile, als wäre auch der Stil durchgesetzt.
- **Versions-Bump + Historie-Zeile** im Lastenheft.
- **Schärfungs-Richtung:** ADRs dürfen die Spezifikation schärfen,
  nie das Lastenheft (siehe `MR-001`-Begründung); wer das Lastenheft
  ändern will, ändert es direkt — als Change Request, nicht per ADR.
- **Beleg-Pflicht:** Test, Gate, Demo oder ADR folgt mit dem
  umsetzenden Slice
  ([`harness/README.md` §Traceability rules](README.md#traceability-rules)).

## Zusatzklassen-Deklaration für Sensors-Bindung

Zusätzlich zu den vier kanonischen Klassen (ADR, Carveout, Schwelle,
Reproduzierbarkeit):

| Klasse | Form | Bedeutung | Beispiel |
|---|---|---|---|
| AC-Bindung | `AC-…` | Gate prüft eine konkrete Lastenheft-Anforderung | [`AC-QA-01`](../spec/lastenheft.md#ac-qa-01--determinismus) für den Determinismus-Test in `make test`; [`AC-QA-02`](../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze) für `make arch-check` (Dogfooding) |

## Modus-Deklaration pro Sub-Area

Der Modus ist **pro Sub-Area** zu deklarieren, nicht pro Repo: „ein Repo hat
einen Bootstrap-Modus" ist laut Baseline ausdrücklich ein Anti-Pattern, weil
der Modus ein *beobachtbares Verhältnis zwischen Code und Doku* ist und nicht
ein Etikett. Bis 2026-07-25 stand hier eine einzige `*`-Zeile; das war genau
diese Verkürzung (Fund **B-2** aus
[slice-048](../docs/plan/planning/done/welle-12/slice-048-modul-delta-lesen.md), behoben
in slice-056).

**Qualifikation.** Eine Sektion ist eine Sub-Area, wenn sie mindestens **zwei**
der drei Inklusions-Achsen erfüllt: (1) eine eigene `MR-NNN`-Adaption wäre
plausibel formulierbar, (2) eine eigene Diskrepanz-/Inventur-Zeile ist sinnvoll,
(3) es gibt eine eigene Pfad-/Datei-Familie. Zu grobe Schnitte („das Backend")
bündeln mehrere Sub-Areas und werden ausdifferenziert.

Die **Kürzel**-Spalte trägt seit `v6.0.0` jedes Repo, dessen Kennungen ein Bereichssegment führen
([`grundlagen-harness-dateien.md` §Konventionsspeicher](../.harness/baseline/v6.5.0/regelwerk/grundlagen-harness-dateien.md#harnessconventionsmd-als-konventionsspeicher)) —
seit die Beobachtungs-Kennung selbst der Pfad `BEO-<KUERZEL>/<slug>` ist, trifft das zu
([slice-138](../docs/plan/planning/done/wellenlos/slice-138-sub-area-kuerzel.md)). Kurz, GROSS, ohne
Leerzeichen, ab Vergabe unveränderlich.

| Sub-Area (Pfad / Modul) | Kürzel | Achsen | Modus | Begründung | Graduation / Folge-Slice |
|---|---|---|---|---|---|
| **Spec-Straten** — `spec/` | `SPEC` | 1,2,3 | Greenfield | Anforderung vor Code, ausnahmslos: jede `AC-*` entstand vor ihrer Implementierung; eigene Adaptionen `MR-001`/`MR-002`/`MR-004` | n/a (GF) |
| **Entscheidungen** — `docs/plan/adr/` | `ADR` | 1,2,3 | Greenfield | ADR vor Code; Immutabilität maschinell durchgesetzt (`doc-immutable`) | n/a (GF) |
| **Kern und Regeln** — `internal/hexagon/` | `KERN` | 1,2,3 | Greenfield | jede Regel hat eine `AC-FA-RULE-*` als Anker; Dogfooding über `arch-check` | n/a (GF) |
| **Adapter** — `internal/adapter/` | `ADAPT` | 2,3 | Greenfield | Ports vor Adaptern; die Schichtung ist selbst gegatet | n/a (GF) |
| **Planungs-Harness** — `docs/plan/planning/` | `PLAN` | 1,2,3 | Greenfield | Form und Größen-Regel stehen in der Vorlage, der Sensor `doc-structure` prüft sie — **seit slice-052**; davor war die Praxis unbelegt | n/a (GF), erreicht mit slice-052 |
| **Gate-/Werkzeug-Schicht** — `tools/`, `Makefile`, `Dockerfile`, `.claude/`, `.github/workflows/` | `GATE` | 1,2,3 | Greenfield | jedes Target ist in `AGENTS.md` §4 deklariert, bevor es zählt; `gate-consistency` erzwingt das. Pfadliste um `.github/workflows/` ergänzt mit [slice-138](../docs/plan/planning/done/wellenlos/slice-138-sub-area-kuerzel.md) — deckt, was zuvor lose als „CI-Schicht"/„CI-/Build-Schicht"/„Durchsetzungsschicht" firmierte | n/a (GF) |
| **Review-Harness** — `docs/reviews/` | `REVIEW` | 2,3 | Greenfield | Konvention und Skill vor dem Report | n/a (GF) |
| **Harness-Einstieg** — `AGENTS.md`, `CLAUDE.md`, `harness/` | `HARNESS` | 1,2,3 | Greenfield | Briefing und Konventionen entstehen vor der Regel, die sie beschreiben; die aufgelösten Adaptionen zur Source Precedence und zum Vendoring lebten genau hier. **Nachgetragen mit slice-101**: die Lücke war seit slice-091 in vier Slices benannt (`BEO-001` im [Beobachtungs-Register](../docs/plan/planning/observations/README.md)) und hatte bis dahin keinen Ort | n/a (GF) |
| **Vendored Baseline** — `.harness/baseline/` | *(keins)* | 1,3 | **kein Modus** | externer, unveränderter Fremdtext ([`MR-006`](#mr-006--baseline-committet-vendored-statt-per-url-referenziert)); GF/BF beschreiben das Verhältnis *eigener* Doku zu *eigenem* Code und sind hier nicht anwendbar; kein Kürzel, weil hier nie eine Beobachtung über a-checks eigene Konvergenz entstehen kann | Aktualisierung nur als Migrations-Slice |

**Alle Sub-Areas mit Modus stehen auf Greenfield.** Das ist kein Zufall und
keine Beschönigung: das Repo ist als Greenfield gestartet und hat die
Doc-vor-Code-Reihenfolge über 55 Slices gehalten. Der Gewinn dieser Tabelle
liegt darum **nicht** in einer Modus-Korrektur, sondern in den benannten
**Inventur-Linien** — ab jetzt ist pro Sektion sagbar, *was* driften würde und
*wer* es sehen müsste. Ein `*` konnte das nicht.

**Drift-Anzeichen, auf die zu achten ist** (Baseline, drei Sensoren): häufen
sich Diskrepanzen in einer GF-Sub-Area, übertrifft der Test-Bestand die
Spec-Anker, oder muss die Spec regelmäßig dem Code nachgezogen werden statt
umgekehrt — dann ist die Sub-Area faktisch nach Brownfield gedriftet und die
Zeile wird geändert, nicht die Beobachtung.
