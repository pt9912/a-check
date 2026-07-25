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
- **Stand:** [`v3.5.2`](https://github.com/pt9912/ai-harness-course/releases/tag/v3.5.2)
  (Release-Tag), **committet vendored** unter
  [`.harness/baseline/v3.5.2/`](../.harness/baseline/v3.5.2/regelwerk/README.md) —
  siehe [`MR-006`](#mr-006--baseline-committet-vendored-statt-per-url-referenziert)
- **Datum der Adoption:** 2026-06-20 (`v1.3.0`); **Stand gehoben auf `v3.5.2` am 2026-07-25**
  — der Sprung ist in vier Etappen geschnitten, diese ist Etappe A (Vendoring + Stand);
  die inhaltliche Angleichung von Konventionen, `MR-*`-Bestand und Template-Konformität folgt
  in den Etappen B–D. Für den **Migrations-Auftrag** gilt: wo eine Adaption dieses Repos nur
  deshalb existiert, weil der `v1.3.0`-Default anders war, **gewinnt der `v3.5.2`-Default**
  (Maintainer-Vorgabe 2026-07-25) — die Prüfung Adaption-für-Adaption ist Etappe C.
  **Bis diese Prüfung erfolgt ist, bleiben die deklarierten `MR-*` in Kraft**: sie sind teils
  maschinell gegatet (`MR-002`/`MR-004` über die `ids`-Muster in
  [`.d-check.yml`](../.d-check.yml)), und ein pauschaler Vorrang würde für das heute gültige
  ID-Schema zwei einander ausschließende Regeln nebeneinander stellen.

## Adoptierte Konventions-Quellen

- **Vendored Baseline (Regelwerk + Templates) — die Lese-Form:**
  [`.harness/baseline/v3.5.2/regelwerk/README.md`](../.harness/baseline/v3.5.2/regelwerk/README.md)
  (Index) und
  [`.harness/baseline/v3.5.2/templates/README.md`](../.harness/baseline/v3.5.2/templates/README.md),
  materialisiert aus dem self-contained `lab-regelwerk.zip` des Releases; Integrität über
  `.harness/baseline/v3.5.2/SHA256SUMS`. **Netzlos** auf jedem Checkout, pro Abschnitt eine
  Datei — ein Agent lädt den benötigten Abschnitt, nie das ganze Bundle
  ([`MR-006`](#mr-006--baseline-committet-vendored-statt-per-url-referenziert)).
- **Extern (Lehrmaterial, maßgeblich für den Inhalt):**
  [`ai-harness-course@v3.5.2`](https://github.com/pt9912/ai-harness-course/tree/v3.5.2)
  (Kurs unter `kurs/de/`). Das vendored Regelwerk ist ein **didaktik-freier Extrakt** und
  trägt keine eigene Normativität: bei Konflikt gilt der Kurs, über ihm die kanonischen
  Quellen dieses Repos (Source Precedence).
- **Konventions-Vorbild (Harness-Form):**
  [`d-check`](https://github.com/pt9912/d-check) — Schwester-Tool im
  selben Stack; Harness-Form (`AGENTS.md`/`harness/`-Trias),
  Hexagon-Ordnerkonvention, Dockerfile-/Makefile-Muster, Pin-Politik,
  Gate-Nachweis-Mechanik (Working-Tree-Hash, `.claude`-Hooks) werden von
  dort übernommen, sobald die jeweiligen Slices sie anlegen.
- **Problem-Quellen (konsolidierte Vorläufer):** die vier divergenten
  `arch-check.sh`-Varianten, die dieses Tool ablöst — `b-cad` (C++),
  `d-check` (Go), `grid-guide` (Rust), `d-migrate` (Kotlin; dort heute
  nur Review statt Fitness-Function für laterale Adapter/Port-Dialekte).
  Sie definieren die *Anforderung* (siehe
  [`spec/lastenheft.md`](../spec/lastenheft.md) §Zweck), nicht die
  Harness-Form.
- **In-Repo (verkörperte Form):** `AGENTS.md`, `harness/README.md`,
  Verzeichniskonvention `spec/` + `docs/plan/` + `harness/`.

## Adaptions-Block

**Disziplin** (aus dem Konventions-Template der Baseline): Einträge sind
**chronologisch pro Repo** nummeriert und tragen die Pflichtfelder Datum,
Geltungsbereich, Adaption, Begründung, Auflösungs-Trigger. An einem
akzeptierten Eintrag wird **nichts nachträglich inhaltlich geändert** —
Korrekturen entstehen als neuer `MR` oder als ausdrückliche Aufhebung,
analog zur ADR-Immutabilität ([`AGENTS.md`](../AGENTS.md) §3.5).

> **Zur Nummern-Frage (geprüft in slice-055, ohne Befund).** Das
> Konventions-Template druckt konkrete Einträge ab — unter anderem eine
> `MR-003` für das Vendoring, das dieses Repo als
> [`MR-006`](#mr-006--baseline-committet-vendored-statt-per-url-referenziert)
> führt. Das ist **keine** Nummern-Kollision: das Template verlangt
> ausdrücklich *chronologische* Nummerierung, und eine chronologische
> Vergabe kann nicht zugleich pro Titel vorgegeben sein. Die abgedruckten
> Nummern sind Beispiel-Instanzen neben dem generischen `MR-NNN`-Muster;
> normativ sind die Pflichtfelder und diese Disziplin. Der in
> [slice-046 §4.2](../docs/plan/planning/open/slice-046-regelwerk-v352-migration-analyse.md)
> als Migrations-Brocken geführte Punkt entfällt damit ersatzlos.

### MR-000 — Baseline-Aussage (inkl. ID-Schema-Deklaration)

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

### MR-001 — Source Precedence mit eigener Spezifikations-Schicht

- **Datum:** 2026-06-20
- **Geltungsbereich:** [`harness/README.md` §Source precedence](README.md#source-precedence)
- **Adaption:** Die Source-Precedence-Tabelle führt
  `spec/spezifikation.md` als eigenen **Rang 2** zwischen Lastenheft
  (Rang 1) und Architektur (Rang 3). Der Kurs-Default setzt zwei
  Spec-Ränge; dieses Repo nutzt drei. Die Dateien der Ränge 2–3 sind mit
  slice-002 angelegt und in den Tabellen verlinkt; Stratum- und
  ID-Schema-Deklaration: [`MR-004`](#mr-004--spezifikation-und-architektur-strata-und-id-schemata).
- **Begründung:** Spec-Stratifizierung mit drei Spec-Dateien; die
  ADR-Schärfungs-Regel („ADR darf Spezifikation schärfen, nicht
  Lastenheft") soll strukturell sichtbar sein. Konsistent mit dem
  Schwester-Repo `d-check`.
- **Auflösungs-Trigger:** permanent.

### MR-002 — ID-Schema mit Bereichskürzeln ab initialer Fassung

- **Datum:** 2026-06-20
- **Geltungsbereich:** [`spec/lastenheft.md`](../spec/lastenheft.md), alle Traceability-Verweise
- **Adaption:** Funktionale Anforderungen verwenden von Beginn an
  Bereichskürzel: `AC-FA-<BEREICH>-<NNN>` (z. B.
  [`AC-FA-RULE-001`](../spec/lastenheft.md#ac-fa-rule-001--kern-reinheit-regel-core-impurity))
  statt des zweistelligen Kurs-Defaults `<PREFIX>-FA-<NN>`.
  Nichtfunktionale Anforderungen bleiben beim Kurs-Default
  (`AC-QA-<NN>`).
- **Begründung:** Das Lastenheft konsolidiert vier divergente
  Architektur-Checker und trägt von Anfang an mehrere Regel- und
  Funktionsbereiche (`RULE`/`EXTRACT`/`CLI`/`CONF`/`DIST`); eine spätere
  Schema-Migration wäre teurer als ein Bereichsschema ab Welle 1.
- **Auflösungs-Trigger:** permanent.

### MR-003 — Source Precedence ohne `docs/user`-Rang

- **Datum:** 2026-06-20
- **Geltungsbereich:** [`harness/README.md` §Source precedence](README.md#source-precedence),
  [`AGENTS.md` §2](../AGENTS.md#2-kanonische-quellen-source-precedence)
- **Adaption:** Der Template-Default führt neun Ränge inkl. eines Rangs
  `docs/user/*` (Operations, Quality, Releasing); a-check führt acht
  Ränge ohne `docs/user`, weil noch kein Operations-Doku-Stratum
  existiert (CLI-Tool vor dem ersten Release).
- **Begründung:** Ein Rang für nicht existierende Dateien wäre ein
  halluzinierter Eintrag (gleiche Klasse wie ein behauptetes Gate); die
  Rangordnung ist laut Baseline projektspezifische Wahl, die hier
  deklariert wird.
- **Auflösungs-Trigger:** mit der Release-Pipeline entsteht Betriebs-/
  Releasing-Doku; der `docs/user`-Rang wird dann eingefügt und dieser
  Eintrag als aufgelöst markiert.
- **Aufgelöst (2026-06-21):** Mit dem
  [Benutzerhandbuch](../docs/user/benutzerhandbuch.md) existiert Nutzer-/
  Betriebs-Doku; der `docs/user`-Rang ist als Rang 6 (vor `README.md`)
  eingefügt — die Source Precedence führt jetzt neun Ränge.

### MR-004 — Spezifikation und Architektur: Strata und ID-Schemata

- **Datum:** 2026-06-21
- **Geltungsbereich:** [`spec/spezifikation.md`](../spec/spezifikation.md), [`spec/architecture.md`](../spec/architecture.md)
- **Adaption:** Die mit [`MR-001`](#mr-001--source-precedence-mit-eigener-spezifikations-schicht)
  angekündigten Ränge 2–3 sind angelegt. Stratum-Platzierung und ID-Schema
  werden hier deklariert (sonst stille Setzung):
  - **Technik-Stratum** = `spec/spezifikation.md`; ID-Präfix
    `SPEC-<BEREICH>-<NNN>` (Bereiche initial `CONF`/`EXTRACT`/`RULE`/`CLI`/
    `DET`/`DIST`). Präzisiert das Lastenheft, erweitert es nie; sprachneutral.
  - **Sicht-Stratum** = `spec/architecture.md`; ID-Präfix `ARC-<NNN>` für
    *Struktur*-Kennungen (Komponenten/Schnittstellen), **keine** eigenen
    Anforderungen; sprach- und meilensteinfrei.
- **Begründung:** Ein Spec-Dokument ohne deklariertes Stratum/ID-Schema ist
  eine stille Setzung (gleiche Harness-Lüge-Klasse wie ein undeklariertes
  Gate) und nicht normativ zitierbar. Die Schemata spiegeln die
  Bereichskürzel des Lastenhefts
  ([`MR-002`](#mr-002--id-schema-mit-bereichskürzeln-ab-initialer-fassung))
  für durchgängige Traceability.
- **Auflösungs-Trigger:** permanent. Die maschinelle Kennungs-Linkpflicht
  (`ids`-Muster für `SPEC-*`/`ARC-*` in [`.d-check.yml`](../.d-check.yml))
  folgt mit dem Implementierungs-Slice, der die übrigen `d-check`-Module
  aktiviert (konsistent mit der dortigen Deferred-Politik).

### MR-005 — Referenzmatrix: intra-Spec-Richtung + ADR→Slice-Disziplin (d-check-Angleichung)

- **Datum:** 2026-07-04
- **Geltungsbereich:** [`.d-check.yml`](../.d-check.yml) (`matrix`-Modul),
  [`docs/plan/adr/`](../docs/plan/adr/)
- **Adaption:** a-check übernimmt die vollständigere Referenzmatrix-Kodierung des
  Schwester-Repos `d-check` (dort `DC-FA-MTX-001/002/003`). Zusätzlich zu den bisherigen
  Cross-Klassen-Regeln (`spec-straten → adr`, `spec-straten → slice`):
  - **intra-Spec-Richtung** (`order` + `direction: no-downward`): Rang
    `lastenheft → spezifikation → architecture` (autoritativste Schicht zuerst); ein
    Abwärtsverweis zwischen Spec-Straten (auch transitiv, z. B. Lastenheft →
    Architektur) ist `matrix-downward`.
  - **`adr → slice`-Regel + Token-Erkennung** (`token: 'slice-\d{3}'`): eine
    Slice-Kennung im **ADR-Körper** ist `matrix-forbidden` — außer per Provenance-Marker
    `<!-- d-check:status-provenance -->` deklariert. ADRs verweisen abwärts **nur** als
    deklarierte Provenance/Verifikations-Zeiger, **nie** als Entscheidungsgrundlage
    (die Argumentation läuft aufwärts über Spec/Verhalten, `Schärft:`-Feld).
- **Grandfathering:** die vor der Übernahme `Accepted`-ADRs (0001–0020) sind immutabel
  ([`AGENTS.md` §3.5](../AGENTS.md#35-adrs-sind-nach-accepted-immutable)) und nennen
  Slices als legitime Verifikations-Zeiger im Körper — sie werden per `exempt-paths`
  ganz übersprungen; neue ADRs **ab 0021** sind **slice-token-frei** (aus
  AGENTS/Modul/Verhalten argumentiert) **oder** tragen — falls sie einen Slice als
  Provenance/Verifikation nennen — den Provenance-Marker (gelebte Praxis: 0021/0022
  sind slice-token-frei, kein Marker nötig).
- **Begründung:** schärft [`AGENTS.md` §3.4](../AGENTS.md#34-architektur-sprach-meilensteinfrei-spec-straten-nie-abwärts)
  maschinell (Spec-Straten-Abwärts **und** ADR-aus-Planung-Argumentation). a-checks
  gepinntes `d-check` unterstützt die Schlüssel (verifiziert 2026-07-04 gegen v0.35.0);
  der Pin ist seit slice-019 auf `v0.37.1`, seit slice-036 auf `v0.51.1` gehoben.
- **Auflösungs-Trigger:** permanent.

### MR-006 — Baseline committet vendored statt per URL referenziert

- **Datum:** 2026-07-25
- **Geltungsbereich:** [`AGENTS.md`](../AGENTS.md) §1, [§Baseline](#baseline),
  [`harness/README.md` §Guides](README.md#guides-feedforward-quellen)
- **Adaption:** Provenienz/Konkretisierung, **keine** inhaltliche Abweichung vom
  `v3.5.2`-Default — im Gegenteil: sie **stellt ihn her**. Regelwerk *und* Templates der
  Baseline liegen **committet vendored** unter
  `.harness/baseline/v3.5.2/{regelwerk,templates}/`, materialisiert aus dem
  self-contained `lab-regelwerk.zip` des Releases, mit `SHA256SUMS` als
  Integritätsmanifest. Bis dahin referenzierte dieses Repo die Baseline **nur per URL**
  (Modell `v1.3.0`).
- **Begründung:** Weil `regelwerk/` und `templates/` **parallel** vendored liegen, lösen die
  `../templates/…`-Verweise des Regelwerks („so sieht das Artefakt aus") **netzlos lokal**
  auf — ein URL-Verweis kann das nicht. Der Nachschlag wird damit reproduzierbar über den
  `<tag>` und unabhängig von Netz und Login; die Kontext-Hygiene bleibt gewahrt, weil pro
  Abschnitt eine Datei geladen wird statt des ganzen Bundles. Real-Vorbild in der Flotte:
  `ai-harness-init` (dortige `MR-007`, vendored `v3.5.1`).
- **Abgrenzung:** Der vendored Baum ist **externer, unveränderter Fremdtext** mit eigenen
  Platzhaltern und Template-Kennungen — er ist **nicht** a-checks Doku. Darum nimmt
  [`.d-check.yml`](../.d-check.yml) ihn per `scan.ignore` aus dem Doku-Gate: sonst prüfte
  `doc-check` fremde Platzhalter gegen a-checks Kennungs- und Linkregeln.
- **Nummern-Hinweis:** Das `conventions.template.md` der Baseline führt diese Adaption als
  `MR-003`. Diese Nummer ist hier belegt
  ([`MR-003`](#mr-003--source-precedence-ohne-docsuser-rang), aufgelöst 2026-06-21), darum
  die nächste freie — dieselbe Praxis wie in `ai-harness-init` (dort `MR-007`). Ob die
  Nummern-Identität mit dem Template hergestellt wird, entscheidet Etappe C der Migration.
- **Auflösungs-Trigger:** permanent (Provenienz/Baseline-Konformität).

### MR-007 — ADR-Vorlagen-Version: `v3.5.2` statt `v1.3.0`

- **Datum:** 2026-07-25
- **Geltungsbereich:** [`MR-000`](#mr-000--baseline-aussage-inkl-id-schema-deklaration)
  §ID-Schema-Deklaration, Zeile zu `ADR-NNNN`
- **Adaption:** [`MR-000`](#mr-000--baseline-aussage-inkl-id-schema-deklaration) deklariert
  `ADR-NNNN` als vierstellig „gemäß Kurs-ADR-Vorlage `v1.3.0`". Diese Versionsangabe ist seit
  der Migration überholt; maßgeblich ist die vendored Vorlage
  `.harness/baseline/v3.5.2/templates/docs/plan/adr/adr.template.md`. **Das Schema selbst
  ändert sich nicht** — vierstellig, chronologisch über den ADR-Index; abgelöst wird
  ausschließlich die Versions-Referenz.
- **Begründung:** `MR-000` wird **nicht** korrigiert. Die Disziplin des Adaptions-Blocks
  verbietet nachträgliche inhaltliche Änderungen an akzeptierten Einträgen — dieselbe Logik wie
  bei `Accepted`-ADRs ([`AGENTS.md`](../AGENTS.md) §3.5). Eine aktuell **behauptende**
  Versionsaussage stehenzulassen wäre allerdings eine stille Falschaussage; darum dieser
  ablösende Eintrag. Gefunden als B-7 in
  [slice-048](../docs/plan/planning/done/slice-048-modul-delta-lesen.md).
- **Auflösungs-Trigger:** permanent — bis zur nächsten Baseline-Migration, die ihn ihrerseits
  ablöst.

### MR-008 — Kein Replay, keine Agenten-Telemetrie

- **Datum:** 2026-07-25
- **Geltungsbereich:** gesamtes Repo; betrifft die Baseline-Module `modul-12`
  (Replay/Evaluierung) und `modul-15` (Observability)
- **Adaption:** a-check führt **kein** Replay-Manifest (`evals/golden/…`, Golden Sets,
  Drift-Rate) und **keine** Agenten-Telemetrie (Tool-Call-Spans, Token-Attribution pro Rolle,
  Cache-Counter). Beide Module bleiben unverkörpert.
- **Begründung:** Beide adressieren die Nicht-Determinismus- und Kosten-Risiken einer
  **Agenten-Laufzeit**. a-check ist ein deterministisches CLI: gleiche Eingabe, gleiche
  Ausgabe, vertraglich zugesichert
  ([SPEC-DET-001](../spec/spezifikation.md#spec-det-001--determinismus-vertrag)) und durch
  Akzeptanztests belegt. Ein Golden Set über deterministischen Output ist ein zweiter,
  schlechter gepflegter Testbestand; Span-Telemetrie über einen Prozess ohne Modellaufruf misst
  nichts. Die *Reproduzierbarkeits*-Hälfte, die `modul-12` über den `image_hash` einfordert,
  ist ohnehin erfüllt — digest-gepinnte Basis-Images und ein digest-gepinntes Release
  ([AC-QA-03](../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit)).
- **Folgewirkung, ausdrücklich:** das Baseline-Closure-Kriterium einer Welle „Replay-Lauf grün"
  (`modul-06`) ist für dieses Repo **unerfüllbar** und wird durch „`make ci` grün" ersetzt. Ohne
  diese Zeile bliebe ein Wellen-Closure dauerhaft unvollständig, ohne dass jemand sagen könnte,
  warum.
- **Auflösungs-Trigger:** sobald a-check eine nicht-deterministische Komponente enthält (etwa
  ein Modell-gestütztes Heuristik-Modul) oder Agenten-Läufe **im Repo selbst** abrechenbar
  werden. Beides ist heute nicht absehbar; der Eintrag ist bis dahin **nicht** permanent,
  sondern begründet ausgesetzt. Gefunden als B-19 in
  [slice-048](../docs/plan/planning/done/slice-048-modul-delta-lesen.md).

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
- **Drei Akzeptanzkriterien** (Happy/Boundary/Negative im
  Given/When/Then-Stil) plus explizite **Out-of-Scope**-Liste.
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

| Sub-Area (Pfad / Modul) | Modus | Begründung | Graduation-Bedingung / Folge-Slice |
|---|---|---|---|
| `*` (Default für gesamtes Repo) | Greenfield | Projekt startet spec-first; Doc führt, Code folgt | n/a (GF) |
