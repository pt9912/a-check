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
- **Stand:** [`v5.12.0`](https://github.com/pt9912/ai-harness-course/releases/tag/v5.12.0)
  (Release-Tag) — **Kurs-Welle 98 · 2026-08-26**, wie im Kopf des vendored
  [`regelwerk/README.md`](../.harness/baseline/v5.12.0/regelwerk/README.md) ausgewiesen;
  **committet vendored** unter
  [`.harness/baseline/v5.12.0/`](../.harness/baseline/v5.12.0/regelwerk/README.md) —
  siehe [`MR-006`](#mr-006--baseline-committet-vendored-statt-per-url-referenziert)
- **Datum der Adoption:** 2026-06-20 (`v1.3.0`). **Stand gehoben auf `v3.5.2` am 2026-07-25**,
  vollständig durchgeführt in `welle-12` (Etappen A–F, Slices 046–078, geschlossen 2026-08-09) —
  aus dieser Migration ist **nichts offen**.
- **Laufende Migration auf `v5.12.0`** (Stand gehoben am 2026-08-29, Analyse
  [slice-092](../docs/plan/planning/done/slice-092-regelwerk-v5120-delta-analyse.md)): geschnitten
  in vier Etappen. **A** (Vendoring + Stand), **B** (Adaptions-Durchgang) und **C** (Form:
  Verzeichnis-Form des Adaptions-Speichers, Ausführung der Urteile, Slice-Form, Rest der Form)
  sind abgeschlossen. Offen ist allein **D** — die neue Mechanik: Risiko-Ausgänge bei Closure und
  das Beobachtungs-Register.
- **Ein Stand liegt vendored**, `v5.12.0`. Der Vorgänger `v3.5.2` lag während der Etappen A–C
  daneben — Vorschrift, nicht Rückstand: er war die Vergleichsgrundlage des Form-Reviews und ist
  mit dessen Abschluss entfallen (`modul-02` §Freshness-Audit; slice-099). `make regelwerk-check`
  meldet seitdem wieder genau einen Stand.

## Adoptierte Konventions-Quellen

- **Vendored Baseline (Regelwerk + Templates) — die Lese-Form:**
  [`.harness/baseline/v5.12.0/regelwerk/README.md`](../.harness/baseline/v5.12.0/regelwerk/README.md)
  (Index) und
  [`.harness/baseline/v5.12.0/templates/README.md`](../.harness/baseline/v5.12.0/templates/README.md),
  materialisiert aus dem self-contained `lab-regelwerk.zip` des Releases; Integrität über
  `.harness/baseline/v5.12.0/SHA256SUMS`. **Netzlos** auf jedem Checkout, pro Abschnitt eine
  Datei — ein Agent lädt den benötigten Abschnitt, nie das ganze Bundle
  ([`MR-006`](#mr-006--baseline-committet-vendored-statt-per-url-referenziert)).
- **Extern (Lehrmaterial, maßgeblich für den Inhalt):**
  [`ai-harness-course@v5.12.0`](https://github.com/pt9912/ai-harness-course/tree/v5.12.0)
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
> [slice-046 §4.2](../docs/plan/planning/done/slice-046-regelwerk-v352-migration-analyse.md)
> als Migrations-Brocken geführte Punkt entfällt damit ersatzlos.

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

Jede Zeile trägt **zwei** Anker: die stabile Kennung `mr-<NNN>` — die Adresse, unter der andere
Dateien referenzieren — und daneben den **alten Überschriften-Slug** aus der Inline-Form. Ohne den
zweiten rotten alle vor slice-096 veröffentlichten Verweise.

Die Spalte *Ersetzt-Baseline-Regel* ist das Pflichtfeld des neuen Stands. Sie steht hier überall
auf `—`, weil sie in einen akzeptierten Eintrag **nicht nachgetragen** werden darf (Einträge
werden nie überschrieben); sie entsteht in den Nachfolge-Einträgen der Etappe C2
([slice-095 §4](../docs/plan/planning/done/slice-095-adaptions-durchgang-v5120.md)).

| MR | Titel | Geltungsbereich | Ersetzt-Baseline-Regel |
|---|---|---|---|
| [MR-011](conventions/MR-011-verfeinerungs-form.md) <a id="mr-011"></a> | Verfeinerungen tragen `SPEC-*` statt der Suffix-Form | [`spec/spezifikation.md`](../spec/spezifikation.md) | [`grundlagen-source-precedence.md` §ID-Schema als Klammer](../.harness/baseline/v5.12.0/regelwerk/grundlagen-source-precedence.md#id-schema-als-klammer) |
| [MR-012](conventions/MR-012-referenzmatrix-grandfathering.md) <a id="mr-012"></a> | Referenz-Richtung maschinell, ADRs 0001–0020 grandfathered | [`.d-check.yml`](../.d-check.yml) (`matrix`), [`docs/plan/adr/`](../docs/plan/adr/) | [`grundlagen-referenz-richtung.md` §Referenz-Richtung (SDP)](../.harness/baseline/v5.12.0/regelwerk/grundlagen-referenz-richtung.md#referenz-richtung-sdp-wer-darf-wen-referenzieren) |
| [MR-013](conventions/MR-013-adr-vorlagen-version.md) <a id="mr-013"></a> | ADR-Vorlage ist die vendored Fassung `v5.12.0` | [`MR-000`](#mr-000) §ID-Schema, Zeile zu `ADR-NNNN` | — *(korrigiert eine Repo-Aussage; Rückbau-Kandidat, im Eintrag benannt)* |
| [MR-014](conventions/MR-014-keine-agenten-telemetrie.md) <a id="mr-014"></a> | Keine Agenten-Telemetrie | gesamtes Repo; Baseline-Modul `modul-15` | [`modul-15-observability.md` §Kernidee](../.harness/baseline/v5.12.0/regelwerk/modul-15-observability.md#kernidee-modul-15) |
| [MR-015](conventions/MR-015-welle-closure-ohne-replay.md) <a id="mr-015"></a> | Welle-Closure ohne Replay-Lauf (`make ci` grün) | [`docs/plan/planning/`](../docs/plan/planning/README.md) | [`modul-06-roadmap.md` §Wellen-Closure-Prozedur](../.harness/baseline/v5.12.0/regelwerk/modul-06-roadmap.md#wellen-closure-prozedur-modul-6) |
| [MR-016](conventions/MR-016-validator-unbesetzt.md) <a id="mr-016"></a> | Validator-Rolle unbesetzt | gesamtes Repo; Baseline-Modul `modul-08` | [`modul-08-agentenrollen.md` §Die neun Übergaben](../.harness/baseline/v5.12.0/regelwerk/modul-08-agentenrollen.md#die-neun-übergaben-und-ihre-artefakte-modul-8) |

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
| [MR-004](conventions/done/MR-004-spec-strata-id-schemata.md) <a id="mr-004"></a><a id="mr-004--spezifikation-und-architektur-strata-und-id-schemata"></a> | [MR-011](conventions/MR-011-verfeinerungs-form.md) |
| [MR-005](conventions/done/MR-005-referenzmatrix.md) <a id="mr-005"></a><a id="mr-005--referenzmatrix-intra-spec-richtung--adrslice-disziplin-d-check-angleichung"></a> | [MR-012](conventions/MR-012-referenzmatrix-grandfathering.md) |
| [MR-006](conventions/done/MR-006-baseline-vendored.md) <a id="mr-006"></a><a id="mr-006--baseline-committet-vendored-statt-per-url-referenziert"></a> | [MR-010](conventions/done/MR-010-rueckbau-drei-adaptionen.md) |
| [MR-007](conventions/done/MR-007-adr-vorlagen-version.md) <a id="mr-007"></a><a id="mr-007--adr-vorlagen-version-v352-statt-v130"></a> | [MR-013](conventions/MR-013-adr-vorlagen-version.md) |
| [MR-008](conventions/done/MR-008-kein-replay.md) <a id="mr-008"></a><a id="mr-008--kein-replay-keine-agenten-telemetrie"></a> | [MR-014](conventions/MR-014-keine-agenten-telemetrie.md) |
| [MR-009](conventions/done/MR-009-validator-unbesetzt.md) <a id="mr-009"></a><a id="mr-009--validator-rolle-unbesetzt-zwei-übergaben-ohne-artefakt"></a> | [MR-016](conventions/MR-016-validator-unbesetzt.md) |
| [MR-010](conventions/done/MR-010-rueckbau-drei-adaptionen.md) <a id="mr-010"></a> | — *(Rückbau-Eintrag; mit seiner Entstehung erledigt, siehe Datei)* |

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
  **Maschinell geprüft ist davon nur, dass die vier Bausteine benannt sind**
  (`make verify-ac-form`, ab slice-054); ob ein Kriterium tatsächlich im
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
[slice-048](../docs/plan/planning/done/slice-048-modul-delta-lesen.md), behoben
in slice-056).

**Qualifikation.** Eine Sektion ist eine Sub-Area, wenn sie mindestens **zwei**
der drei Inklusions-Achsen erfüllt: (1) eine eigene `MR-NNN`-Adaption wäre
plausibel formulierbar, (2) eine eigene Diskrepanz-/Inventur-Zeile ist sinnvoll,
(3) es gibt eine eigene Pfad-/Datei-Familie. Zu grobe Schnitte („das Backend")
bündeln mehrere Sub-Areas und werden ausdifferenziert.

| Sub-Area (Pfad / Modul) | Achsen | Modus | Begründung | Graduation / Folge-Slice |
|---|---|---|---|---|
| **Spec-Straten** — `spec/` | 1,2,3 | Greenfield | Anforderung vor Code, ausnahmslos: jede `AC-*` entstand vor ihrer Implementierung; eigene Adaptionen `MR-001`/`MR-002`/`MR-004` | n/a (GF) |
| **Entscheidungen** — `docs/plan/adr/` | 1,2,3 | Greenfield | ADR vor Code; Immutabilität maschinell durchgesetzt (`doc-immutable`) | n/a (GF) |
| **Kern und Regeln** — `internal/hexagon/` | 1,2,3 | Greenfield | jede Regel hat eine `AC-FA-RULE-*` als Anker; Dogfooding über `arch-check` | n/a (GF) |
| **Adapter** — `internal/adapter/` | 2,3 | Greenfield | Ports vor Adaptern; die Schichtung ist selbst gegatet | n/a (GF) |
| **Planungs-Harness** — `docs/plan/planning/` | 1,2,3 | Greenfield | Form und Größen-Regel stehen in der Vorlage, der Sensor `verify-slice-form` prüft sie — **seit slice-052**; davor war die Praxis unbelegt | n/a (GF), erreicht mit slice-052 |
| **Gate-/Werkzeug-Schicht** — `tools/`, `Makefile`, `Dockerfile`, `.claude/` | 1,2,3 | Greenfield | jedes Target ist in `AGENTS.md` §4 deklariert, bevor es zählt; `gate-consistency` erzwingt das | n/a (GF) |
| **Review-Harness** — `docs/reviews/` | 2,3 | Greenfield | Konvention und Skill vor dem Report | n/a (GF) |
| **Vendored Baseline** — `.harness/baseline/` | 1,3 | **kein Modus** | externer, unveränderter Fremdtext ([`MR-006`](#mr-006--baseline-committet-vendored-statt-per-url-referenziert)); GF/BF beschreiben das Verhältnis *eigener* Doku zu *eigenem* Code und sind hier nicht anwendbar | Aktualisierung nur als Migrations-Slice |

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
