# Review-Report — slice-002 (Spezifikation + Architektur)

- **Review-Art:** Design-/Plan-Review der neuen Spec-Straten (vor Acceptance)
- **Gegenstand:** `spec/spezifikation.md` (Technik), `spec/architecture.md` (Sicht), `harness/conventions.md` (MR-004), Source-Precedence-Updates (`AGENTS.md`/`harness/README.md`), ADR-`Schärft`-Felder
- **Datum:** 2026-06-21 · **Modell:** Opus 4.8 · **Skill:** `.harness/skills/reviewer.md`
- **Methode:** Drei unabhängige Reviewer-Agenten, frischer Kontext, perspektiven-divers (Stratum-Disziplin/Referenz-Richtung · Lastenheft-Treue · Konsistenz/Querverweise); HIGH-Kandidaten adversarisch gegen das repo-eigene Gate verifiziert.

Befund-Schema (Modul 10): `kategorie` · `quelle` · `pfad` · `befund` · `verifizierbar`.

## Befunde

### MEDIUM (alle behoben)

| ID | quelle | pfad | befund |
|---|---|---|---|
| M1 | AC-FA-RULE-004 | `spec/spezifikation.md` SPEC-CONF/RULE/EXTRACT | `port-impurity` verlangt „verbotene Konstrukte", aber SPEC-CONF hatte keinen Config-Schlüssel dafür und SPEC-EXTRACT liefert nur Imports — Lastenheft-Anforderung ohne Spec-Abdeckung. |
| M2 | AC-FA-CONF-001 | `spec/spezifikation.md` SPEC-CONF | `adapter_sink`/`tech` als optional markiert; `tech`-Optionalität ist AC-Boundary-gedeckt, `adapter_sink` war ungeerdet. |
| M3 | AC-FA-CONF-001 | `spec/spezifikation.md` SPEC-CONF | neue Schlüssel `allow`/`markers` nicht an die referenzierte AC gebunden (durch RULE-004/005 bzw. QA-02 gedeckt, aber nicht ausgewiesen). |
| M4 | SPEC-CLI-001 / ARC-005 | `spec/architecture.md` | ARC-005 „bestimmt den Exit-Code" (impliziert 0/1/2) widerspricht der eigenen Sequenz (Exit 2 beim Config-Adapter/CLI, Report nur 0/1). |

### LOW (behoben)

| ID | pfad | befund |
|---|---|---|
| L1 | `spec/spezifikation.md` SPEC-RULE (`tech-leak`) | Composition-Root-Ausnahme absoluter formuliert als das „ggf." der AC-FA-RULE-003. |
| L2 | `docs/plan/adr/0001-…` | `Schärft: —` referenziert im Begründungstext SPEC-DIST-001, das ADR-0004 formal schärft (Doppel-Bezug). |

### INFO

| ID | befund | Aktion |
|---|---|---|
| I1 | SPEC-EXTRACT Rust-Alias korrekt (`use x as y;` → `x`); vier Sprachen vollständig. | Bestätigung, keine. |
| I2 | AC-FA-CONF-001 Negative fordert „Exit 2 **mit Zeilenangabe**"; SPEC-CLI nannte die Zeilenangabe nicht. | behoben. |
| I3 | Begriffs-Kollision `composition_root` (Config, geprüftes Repo) ↔ ARC-006 „Composition Root" (a-check selbst). | behoben (Disambiguierung in ARC-006). |
| I4 | Source-Precedence-Rangfolge weicht vom Regelwerk-Default ab — deklariert via MR-001/MR-003. | keine (deklariert). |
| I5 | `ids`-Linkpflicht für `SPEC-*`/`ARC-*` deferred — in MR-004 transparent. | keine (deklariert). |

## Negativbefunde (geprüft, ohne Befund)

- **Stratum-Disziplin (Reviewer A, 0 Befunde):** `spezifikation.md` präzisiert ausschließlich (erweitert nie); `architecture.md` trägt keine eigenen Anforderungen; beide sprach-/meilensteinfrei; kein Abwärts-Verweis auf ADR/Slice (adversarisch gegen `.d-check.yml` `matrix`/`ids` bestätigt). ADR-`Schärft` strikt aufwärts; ADR-0001 `—` korrekt sprachneutral begründet. MR-004 deklariert Strata + ID-Schemata vollständig.
- **Lastenheft-Treue (Kernblöcke):** SPEC-CONF Pflichtblöcke/strict-decode/Exit 2, SPEC-RULE-001/002/005, SPEC-DET (Determinismus), SPEC-DIST (statisch/distroless/netzlos/Digest-Pin/`--print-*`) decken ihre AC vollständig und widerspruchsfrei; kein erfundenes Format/Flag.
- **Konsistenz/Dogfooding:** Komponenten ↔ SPEC widerspruchsfrei (ConfigPort↔SPEC-CONF, ExtractionPort↔SPEC-EXTRACT, Kern↔SPEC-RULE, ReportPort↔SPEC-CLI); Schicht-Richtung `core ← ports ← adapters` deckungsgleich über architecture/SPEC-CONF/AC-FA-RULE-005; Dogfooding-Behauptung (AC-QA-02) plausibel; alle vier Zielsprachen abgedeckt.

## Kategorie-Summary

| Kategorie | Anzahl | IDs |
|---|---|---|
| HIGH | 0 | — |
| MEDIUM | 4 | M1–M4 |
| LOW | 2 | L1, L2 |
| INFO | 5 | I1–I5 |

## Verdikt

Kein blockierender (HIGH) Befund. Die zwei neuen Strata sind stratum-diszipliniert, Lastenheft-treu und intern konsistent. Die vier MEDIUM (Spec-Lücke `port-impurity`, Erdung der optionalen/neuen Config-Blöcke, Exit-Code-Widerspruch) waren vor Acceptance zu klären — und sind behoben.

## Disposition (Implementer, 2026-06-21)

Behebung als getrennter Implementer-Schritt nach dem Review. Beleg: `make doc-check` grün.

| Finding | Aktion |
|---|---|
| M1 | SPEC-CONF um `forbidden_constructs` (Schicht → Text-Muster) erweitert, an AC-FA-RULE-004 geerdet; SPEC-EXTRACT erkennt diese Muster text-heuristisch und speist `port-impurity`; SPEC-RULE-`port-impurity` entsprechend präzisiert. |
| M2 | Optionalität je Block mit definierter Semantik + AC-Erdung ausgewiesen (`adapter_sink`→RULE-002 mit Fehlt-Semantik; `tech`→RULE-003/CONF-Boundary). |
| M3 | `allow`→RULE-004/005-Boundary, `markers`→QA-02 als Bezug ausgewiesen. |
| M4 | ARC-005 = Befund-Exit-Code `0`/`1`; ARC-006 meldet `2` (Nutzungs-/Konfigurationsfehler) — konsistent mit der Sequenz. |
| L1 | `tech-leak`: „außerhalb `composition_root`, falls konfiguriert" (matcht das „ggf." der AC). |
| L2 | ADR-0001: SPEC-DIST-001 jetzt als „durch ADR-0004 verbindlich gemacht" ausgewiesen (realisiert ≠ schärft). |
| I2 | SPEC-CLI: „ungültige Config wird mit Zeilenangabe gemeldet". |
| I3 | ARC-006: Disambiguierung gegen den Config-Schlüssel `composition_root`. |
| I1, I4, I5 | keine Aktion (Bestätigung / deklarierte Adaption / deklariert-deferred). |
