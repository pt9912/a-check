# slice-200 — Adoption `v6.13.0`: AGENTS.md-Form und Matrix-Klassen

**Welle:** ohne Welle.

**Bezug:** Folge-Slice aus
[slice-199](../done/wellenlos/slice-199-baseline-v6130-migration.md) §3 (MR-Durchgang,
Schritt 5) und [MR-025](../../../../harness/conventions/MR-025-referenzmatrix-grandfathering-v6130.md)
§Offener Adoption-Aspekt.
[`AC-QA-02`](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze).

**Berührte Spec-Stellen:** — · Der Slice berührt kein Spec-Stratum.

**Verantwortlich:** Claude — gesetzt beim Übergang nach `in-progress/`.

**Autor:** Claude. **Datum:** 2026-09-29.

**Lerneintrag — Form:** wird bei Closure benannt.

---

## 1. Ziel und Abgrenzung

**Ziel:** Die drei Adoption-Aspekte aus dem v6.13.0-Durchgang sind **je entschieden und
umgesetzt** — umsetzen oder als Bestand mit Begründung deklarieren:

1. **WIP-Limit-Formulierung** ([`AGENTS.md`](../../../../AGENTS.md) §5): a-check zählt die
   Menge in `in-progress/`; die Baseline (Welle 148) zählt **den Lauf** des Rolleninhabers.
2. **§5-Form** ([`AGENTS.md`](../../../../AGENTS.md) §5): die Ziel-Form startet §5 als
   Index-Tabelle mit Regel-Auslagerung (Welle 149) — a-checks Liste ist ein Bullet-Bestand
   und längst gewachsen; die Entscheidung folgt der Startwert-Politik der Welle.
3. **Matrix-Klassen** ([`.d-check.yml`](../../../../.d-check.yml) `matrix`): die Baseline
   gibt Welle-, Carveout- und Roadmap-Kennungen eigene Matrix-Klassen (Welle 138); a-checks
   `matrix`-Sektion deckt sie gegenwärtig nicht.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Die Spec-Straten** (Lastenheft-RB-Reihe, Suffix-Form, Architektur-Gliederung).
  *Ein Folge-Slice übernimmt es* mit Kennung:
  [slice-189](../open/slice-189-voll-abgleich-spec-straten.md) trägt den Abgleich ohnehin und läuft
  seit slice-199 gegen `v6.13.0`.
- **slice-190 / slice-191** — *Bestand bleibt bewusst stehen*: eigene Slices, eigener Trigger.

## 2. Definition of Done

- [x] Jeder der drei Aspekte trägt eine Entscheidung mit Beleg — umgesetzt (Diff) oder als
      Bestand mit Begründung gegen die jeweilige Welle.
- [x] `make gates` und `make verify` grün.
- [x] Unabhängiger Review, Report unter [`docs/reviews/`](../../../../docs/reviews/README.md).
- [x] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [x] Beobachtungs-Register fortgeschrieben; jedes Risiko aus §6 trägt einen Ausgang.

## 3. Plan (vor Code)

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| [`AGENTS.md`](../../../../AGENTS.md) | update | WIP-Limit-Formulierung (Aspekt 1), §5-Form (Aspekt 2) |
| [`.d-check.yml`](../../../../.d-check.yml) | update | Matrix-Klassen (Aspekt 3) |

### 3.1 Umsetzung — je Aspekt die Entscheidung

**Aspekt 1 (WIP-Limit) — umgesetzt.** [`AGENTS.md`](../../../../AGENTS.md) §5
zählt jetzt **pro Lauf** (Welle 148: mehrere Läufe derselben Person zählen
einzeln, wenn `Verantwortlich:` sie per Zweig unterscheidet) und hält als
Hauptzweig-Folge fest: **höchstens ein** Slice in `in-progress/`. Damit bleibt
`doc-planning` (DC-FA-PLAN-001), der Verzeichnis-Inhalte zählt, mit §5
kohärent — der Sensor behält seine Form-Aussage. Nachgezogen hat dieselbe
Stelle die **Übergangs-Zahl**: die Baseline (Welle 137/150) kennt **sechs**
Übergänge inklusive `open|next → done` (Konsolidierung), a-check schrieb
bisher „fünf".

**Aspekt 2 (§5-Form) — entschieden, Auslagerung ausgelagert.** Die
Tabellenform mit Regel-Auslagerung (Welle 149) ist **Ziel**: die Liste ist
über die Vier-Regel-Größe der Ziel-Form längst hinausgewachsen. Die Konversion
fasst **17 Regeln** (zweimal verschieden gezählt: Anker-Zähler 17; 13 fett-markierte + 4 unformatierte = 17 — Review slice-200, M2) samt Begründungs- und Durchsetzungstexten und legt bis zu
eine `harness/rules/<name>.md`-Datei je langen Regel an — das ist ein eigener,
reviewbarer Vorgang und geht als [**slice-201**](../open/slice-201-agents-md-index-tabelle.md)
in `open/`, wie der Rückführungs-Vorbehalt in §4 es vorschneidet.

**Aspekt 3 (Matrix-Klassen) — umgesetzt.** [`.d-check.yml`](../../../../.d-check.yml)
führt die Klassen `welle` (Token `welle-[a-z0-9]+…`), `carveout`
(`CO-\d{3}`) und `roadmap` (ohne Token — nur über Links erreichbar) und die
sechs Regeln `spec-straten|adr → welle|carveout|roadmap` (Welle 138). Der
Slice-Token bleibt numerisch (`slice-\d{3}`) — a-checks Schema deklariert
`slice-NNN` ([`MR-000`](../../../../harness/conventions.md#mr-000)); der
Slug-Token der Vorlage träfe hier keine Kennung. `make doc-check` mit den
sechs neuen Regeln: **0 Befunde** — die Kanten waren ungeprüft, nicht
verletzt.

## 4. Trigger

**Start** (`open` → `next` → `in-progress`): das WIP-Limit ist frei.

**Rückführungen — vorab benannt:**

- `in-progress` → `next` (zu groß): wählt ein Aspekt die Auslagerung von §5-Regeln
  (Welle 149), wird die Auslagerung als eigener Slice geschnitten — die Form-Entscheidung
  bleibt hier.
- `in-progress` → `open` (blockiert): entfällt — kein Blocker absehbar.

## 5. Closure-Trigger

DoD vollständig, `make gates` und `make verify` grün, Closure-Notiz mit Lerneintrag.

## 6. Risiken und offene Punkte

- **Die WIP-Limit-Umformulierung ändert die Semantik eines lebenden Sensors** — `doc-planning`
  zählt Verzeichnis-Inhalte; zählt das Limit künftig den Lauf, kann der Sensor nicht mehr die
  Form-Aussage von [`AGENTS.md`](../../../../AGENTS.md) §5 sein. — **Ausgang:** *entfallen*,
  gestrichen mit Begründung: die neue Formulierung hält die **Hauptzweig-Folge** („höchstens
  ein Slice in `in-progress/`") ausdrücklich fest — der Sensor und §5 bleiben kohärent (§3.1).

## 7. Closure-Notiz

**Lerneintrag — Form: geschärfte Regel.** *Ein DoD-Häkchen in `open/` attestiert
Zukunft — zum dritten Mal gefangen, weil kein Sensor die Lebenslage des Häkchens
prüft.* Der von mir angelegte Nachfolge-Slice slice-201 trug alle drei `[x]`,
bevor ein einziges Kriterium erfüllt war — dieselbe Klasse wie slice-169 und
slice-197, diesmal an dem Lauf, der die Attestierungs-Falle aus slice-192
schon kannte. Ausgang: **geplant →
[slice-202](../open/slice-202-dod-haekchen-lebenslage.md)** (Sensor,
`forbid-pattern` auf `open/`- und `next/`-Slices).

**Was hat funktioniert:** die drei Aspekte waren klein und unabhängig; die
Sollform lag im neuen Stand vor (WIP-Limit-Text, Template-§5, Matrix-Block),
so dass jedes „umsetzen oder Bestand" gegen einen Wortlaut messen konnte.
Der unabhängige Review bestätigte die Umsetzung in allen Prüfpunkten — inklusive
vier Mutations-Proben gegen die neuen Matrix-Regeln (rot mit
`matrix-forbidden`).

**Was ging anders als geplant:** der Review fand die Defekte **in meinen
eigenen neuen Plänen**, nicht in der Umsetzung — H1 (abgehaktes DoD in
slice-201) traf die Attestierungs-Falle beim dritten Mal, H2 eine
Kopfsummen-Zahl („15 Regeln", richtig: 17, zweimal verschieden nachgemessen).
Beide sind Plan-Defekte und waren vor der Closure heilbar.

**Steering-Loop-Eintrag:** Attestierungs-Lücke — Sensor geplant: DoD-Häkchen
an die Lebenslage binden, Auslöser
[`BEO-GATE/attestierung-vor-dem-vorgang`](../observations/BEO-GATE/attestierung-vor-dem-vorgang/observation.md)
(slice-169, slice-197, slice-200 — 3×). *(nichts verkörpert — der Eintrag ist
gezählt; der Sensor ist slice-202.)*

**Beobachtungs-Register (`../observations/`):** `evidence/slice-200.md` in
`BEO-GATE/attestierung-vor-dem-vorgang/` ergänzt — Zähler 3×, Ausgang
*geplant* → slice-202.

**Folge-Slices:** [slice-201](../open/slice-201-agents-md-index-tabelle.md)
(§5-Index-Tabelle + Auslagerung) und
[slice-202](../open/slice-202-dod-haekchen-lebenslage.md) (Attestierungs-Sensor)
— sind Dateien in `open/`.

**Risiken aus §6:** jedes mit genau einem Ausgang — siehe §6.

**Drei Paarungen:** Anker — kein Steering-Loop-Eintrag verkörpert, kein
`liegt in`-Feld · Folge-Slice — slice-201 und slice-202 existieren als Dateien
in `open/` · Register — attestierung-vor-dem-vorgang trägt den dritten Beleg
und den Ausgang.

## 8. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** HARNESS (Achsen 1, 2, 3 ✓) und GATE
(Achsen 1, 2, 3 ✓).

**Vorgelagert — offene Beobachtungen sichten:** das Register wurde durchgegangen
(2026-09-29):
[`agents-md-hinkt-baseline-dod-item-hinterher`](../observations/BEO-HARNESS/agents-md-hinkt-baseline-dod-item-hinterher/observation.md)
(offen — der Slice arbeitet genau die Nennungen in [`AGENTS.md`](../../../../AGENTS.md) ab,
Befund wird gezählt) · keine weiteren Treffer in HARNESS/GATE für diesen Vorgang.

**Modus-Begründungsblock:** alle berührten Sub-Areas GF.
