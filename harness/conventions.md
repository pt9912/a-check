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
- **Stand:** [`v1.3.0`](https://github.com/pt9912/ai-harness-course/releases/tag/v1.3.0)
  (Release-Tag, von Beginn an gepinnt)
- **Datum der Adoption:** 2026-06-20

## Adoptierte Konventions-Quellen

- **Extern (Lehrmaterial):**
  [`ai-harness-course@v1.3.0`](https://github.com/pt9912/ai-harness-course/tree/v1.3.0)
  (Templates: `lab/templates/`, Konventionen:
  `kurs/de/grundlagen/konventionen.md`).
- **Extern (Agenten-Destillat):**
  [`agents-regelwerk.md`](https://raw.githubusercontent.com/pt9912/ai-harness-course/v1.3.0/kurs/de/agents-regelwerk.md)
  — operatives Regelwerk für Code-Agenten ohne Didaktik; derivativ,
  bei Konflikt gilt das Lehrmaterial.
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
  Spec-Ränge; dieses Repo nutzt drei. Die Dateien der Ränge 2–3 entstehen
  mit slice-002; bis dahin sind sie in den Tabellen als „geplant"
  markiert und nicht verlinkt.
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
