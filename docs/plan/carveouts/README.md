# Carveouts — temporäre Ausnahmen mit Auflösungs-Plan

**Aktueller Bestand: keiner.** Das ist eine Messung, keine Nachlässigkeit — siehe
[§Warum leer](#warum-dieses-verzeichnis-heute-leer-ist).

Ein **Carveout** ist eine *temporäre* Ausnahme von einer Regel oder einem Gate, die einen
**Auflösungs-Plan** trägt. Die Kernidee der Baseline (`modul-07`): *ein Carveout ohne
Auflösungs-Trigger ist ein permanenter Carveout, der lügt.* Angelegt in slice-065; die ID-Reihe
`CO-NNN` ist seit
[MR-000](../../../harness/conventions.md#mr-000--baseline-aussage-inkl-id-schema-deklaration)
deklariert.

## Vor dem Anlegen: der Trichter

Nicht jede Diskrepanz ist ein Carveout. **Zwei Fragen, in dieser Reihenfolge** — Granularität
**vor** Temporalität. Die Reihenfolge ist der Punkt: sie verhindert den Reflex, jede Ausnahme als
Carveout zu führen.

**Frage 1 — einzelne Diskrepanz oder Cluster?**
Mehrere Ausnahmen im selben Geltungsbereich, oder ein systemisches „Code existiert vor Doku"-Muster
⇒ **BF-Sub-Area-Markierung** in der Modus-Tabelle von
[`harness/conventions.md`](../../../harness/conventions.md#modus-deklaration-pro-sub-area), mit
Graduation-Trigger. Frage 2 entfällt. Einzelne Diskrepanz ⇒ Frage 2.
*(Kein harter Schwellwert für „Cluster" — Faustregel ist der gemeinsame Geltungsbereich, keine
Zahl.)*

**Frage 2 — ist der Trigger ernsthaft zu erreichen?**
Ja, mit absehbarem Aufwand ⇒ **Carveout** (hier, Form unten).
Nein, „das werden wir nie tun" ⇒ **permanente ADR** unter
[`docs/plan/adr/`](../adr/README.md); der Carveout-Entwurf wird nicht weggeworfen, sondern
überführt (Trigger fällt weg, Checkliste reduziert sich auf die Architektur-Folgen).

| Wahl | Symptom | Träger | Ort |
|---|---|---|---|
| **Carveout** | eine konkrete Gate-/Regelausnahme, abgrenzbar, mit Folge-Slice und erreichbarem Trigger | einzelne Diskrepanz | dieses Verzeichnis |
| **BF-Sub-Area-Markierung** | Diskrepanz-Cluster im selben Geltungsbereich | ganze Sub-Area | Modus-Tabelle in [`conventions.md`](../../../harness/conventions.md#modus-deklaration-pro-sub-area) |
| **permanente ADR** | der Trigger ist ehrlich nie zu erreichen | dauerhafte Regel | [`docs/plan/adr/`](../adr/README.md) |

**Nicht hierher gehören bootstrap-aware Gates.** Eine Reifestufung des *Gates selbst* (die Größen-Regel in `doc-structure`
ab slice-052, das AC-Grandfathering, der Stichtag in `commit-scope-check`) ist kein Carveout: sie
stuft die Prüfung, sie nimmt keine Diskrepanz aus. Die Achsen zu verwechseln erzeugt das
„Bootstrap-Schlupfloch" — Stufung ohne Trigger ist Carveout-Wildwuchs, eine Carveout-Kaskade ohne
BF-Markierung ist verschleierte Sub-Area-Brownfield.

## Form eines Carveouts

Datei: `docs/plan/carveouts/CO-<NNN>-<kurztitel>.md`, aus der vendored Ziel-Form
[`.harness/baseline/v6.0.0/templates/docs/plan/carveouts/carveout.template.md`](../../../.harness/baseline/v6.0.0/templates/docs/plan/carveouts/carveout.template.md)
— a-check führt keine eigene Kopie, sie würde gegen die Baseline driften (dieselbe Begründung wie
bei der Slice-Vorlage, [`AGENTS.md`](../../../AGENTS.md) §5, „Slice-Form"; a-checks frühere lokale
Kopie hatte genau das getan — ein Provenienz-Zeiger auf `v6.0.0` bumpte, während der Aufbau noch
auf einem älteren Stand blieb, gefunden und entfernt in slice-149). **Sechs Pflicht-Header-Felder:**
Status · Datum angelegt · Letzte Prüfung · betroffenes Gate · Geltungsbereich · **Folge-Slice**.

**Beim Kopieren anzupassen:** die Verweise course-relativ → a-checks eigene Pfade (`AGENTS.md`
statt Kurs-Konventionen); den Abschnitt `## Geschichte` streichen — er trägt Audit-Einträge einer
Wellen-Closure-Prozedur, die a-check noch nicht durchläuft (Fund **B-13**, siehe
[§Audit](#audit)); das Feld `Betroffenes Gate` auf eines der in
[`AGENTS.md`](../../../AGENTS.md) §4 deklarierten Targets binden, kein erfundenes.

- **Ohne Folge-Slice** ist der Carveout de facto permanent — dann gehört er über den Trichter in
  eine ADR, nicht hierher.
- **Der Auflösungs-Trigger ist eine messbare Bedingung**, die ein anderer Mensch ohne Rückfrage
  als erreicht beurteilen kann — nicht „sobald Zeit ist".
- **Die Gate-Konfiguration zeigt per `# CO-<NNN>`-Kommentar hierher.** Sonst ist die Ausnahme im
  `make gates`-Output eine stille Senkung ohne Begründung.
- **Auflösung ist ein `git mv` nach `done/`**, zusammen mit dem Entfernen der Gate-Ausnahme und
  einem grünen `make gates` ohne sie. Auflösen ohne Verschieben wäre eine zweite Lüge.

## Warum dieses Verzeichnis heute leer ist

Der Bestand wurde in slice-065 gegen den Trichter geprüft. Jede bestehende Ausnahme fällt
begründet woanders hin:

| Ausnahme | Ergebnis |
|---|---|
| `.golangci.yml` §exclusions | permanent, getragen von [ADR-0005](../adr/0005-lint-profil.md) |
| `.d-check.yml` `scan.ignore` für die vendored Baseline | permanent, getragen von [MR-006](../../../harness/conventions.md#mr-006--baseline-committet-vendored-statt-per-url-referenziert) |
| `.d-check.yml` `exempt-paths` für ADR 0001–0020 | permanent, getragen von [MR-005](../../../harness/conventions.md#mr-005--referenzmatrix-intra-spec-richtung--adrslice-disziplin-d-check-angleichung) |
| Grandfathering in `doc-structure` (`exempt-paths`), `verify-ac-form`, `commit-scope-check` | bootstrap-aware Gate, kein Carveout (siehe oben) |
| `regelwerk-check`-Integrität außerhalb des `gates`-Aggregats | kein Gate ausgenommen; `modul-13` führt es als Muster „vorhanden, aber nicht als Gate behauptet" |

Ein leerer Ort ist ehrlich. Eine deklarierte ID-Reihe **ohne** Ort war es nicht — genau der Fund
**B-14**, und dieselbe Klasse wie `next/` vor slice-053.

## Audit

`modul-07` verlangt pro Wellen-Closure einen Audit-Slice `SL-CO-AUDIT-<welle>`, der jeden aktiven
Carveout auf ein aktuelles `Letzte Prüfung:`-Datum und einen noch gültigen Trigger prüft. a-check
schließt bis heute keine Welle auditierbar (Fund **B-13**, offen) — der Audit entsteht mit der
Wellen-Closure-Prozedur in Etappe F 3/3. Bis dahin ist er **nicht** vorhanden, und das steht hier,
statt behauptet zu werden.
