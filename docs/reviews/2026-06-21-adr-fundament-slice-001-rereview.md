# Re-Review (Folgelauf) — Fundament-ADRs (slice-001)

- **Review-Art:** Re-Review / Folgelauf nach Behebung (Modul 10: neue Datei statt Überschreibung)
- **Vorlauf:** `2026-06-21-adr-fundament-slice-001.md`
- **Gegenstand:** ADR-0001…0004 + ADR-Index; neu mitgeprüft: `.harness/skills/reviewer.md`
- **Stand:** nach Commit `7a90b8d`
- **Datum:** 2026-06-21 · **Modell:** Opus 4.8
- **Methode:** Drei unabhängige Reviewer-Agenten (frischer Kontext, perspektiven-divers),
  die jede Korrektur **gegen die realen Artefakte** verifizieren (Closure-Check) und auf
  **Regressionen** prüfen; Synthese mit adversarischer Verifikation strittiger Befunde.
- **Skill:** `.harness/skills/reviewer.md` jetzt vorhanden und angewandt.

## Closure-Verifikation der Erstlauf-Befunde

Alle 15 Befunde des Erstlaufs unabhängig gegen die Artefakte bestätigt:

| Erstlauf | Status (gegen Artefakt verifiziert) |
|---|---|
| M1 (d-check `--print-mk` überzeichnet) | **geschlossen** — ADR-0004 gibt `d-check.mk:5-7`-Stand korrekt wieder |
| M2 (strict-decode d-check unbelegt) | **geschlossen** — explizit „nicht aus d-check abgeleitet"; `.d-check.yml` 0 strict-Treffer deckt sich |
| M3 (ADR als „Technik-Quelle") | **geschlossen** — Formulierung entfernt; Wortlaut jetzt regelwerk-gleich („Begründungs-Schicht *unter* den Spec-Straten") |
| M4 (Pin-Staleness fehlt) | **geschlossen** — Trade-off ergänzt, konsistent mit AC-QA-03 + realem Pin-Modell |
| M5 (Index-Bezug verkürzt) | **geschlossen** — Index führt je ADR alle Kopf-ACs, deckungs- und reihenfolgegleich |
| L1, L2, L3, L6 | **geschlossen** (Bezug AC-QA-01; RULE-* als nachgelagert; Rust-Trade-off abgewogen; Querverweis 0001→0004) |
| L4, L5 | **geschlossen (begründet beibehalten)** — Slice-Tokens als Verifikations-Zeiger; CONF-001 primär, CLI-001 als Exit-Code-Vertrag |
| I1, I2, I3 | **geschlossen** (d-check-Konsistenz auf Sprachwahl eingegrenzt; read-only getrennt; static-Build-Anker konkret) |
| I4 | **geschlossen** — `.harness/skills/reviewer.md` angelegt |

## Neue Befunde (Folgelauf)

### LOW

**NL1 — Reviewer-Skill `Status: Proposed`**
- quelle: Konvention / Modul-10-Vorlage
- pfad: `.harness/skills/reviewer.md:3`
- befund: Die Modul-10-Skill-Vorlage trägt ein `Status`-Feld (dort Wert `Accepted`); die Datei nutzt `Proposed`. Form-konform, Wertwahl vertretbar (neue, noch nicht abgenommene Skill), aber nicht deckungsgleich zur Vorlage.
- verifizierbar: ja

**NL2 — ADR-0001 d-check-Go-Beleg nicht in-Repo zitiert**
- quelle: Provenance / Faktentreue
- pfad: `docs/plan/adr/0001-go-impl-sprache.md` (Entscheidung)
- befund: Die Aussage „d-check … in Go geschrieben" ist im Repo durch `spec/lastenheft.md:44` und `harness/conventions.md:49` belegt, wird in der ADR aber nur per externem GitHub-Link gestützt. Faktentreu, aber die nächstliegende In-Repo-Belegstelle bleibt unzitiert.
- verifizierbar: ja

### INFO

- **NI1** — ADR-Index: Bezug-Spalte ist nun mehrwertig ohne deklariertes „Primär-Bezug"-Schema; konsistent (alle ACs gleichrangig gelistet), kosmetisch. (`docs/plan/adr/README.md`)
- **NI2** — read-only-`--print-mk` ist transparent als Design-Verallgemeinerung über die `--print-config`-Boundary-AK hinaus markiert (kein getarntes AC-Zitat). (`0004`)
- **NI3** — die neue Pin-Staleness-Konsequenz hat naturgemäß keinen maschinellen Anker (wie im Erstlauf für M4 „verifizierbar: nein" notiert); kein Widerspruch. (`0004`)
- **NI4** (positiv) — der static-binary-Anker (`file`/`ldd`) ist jetzt entscheidungsspezifisch & maschinell prüfbar (Modul-4-Maßstab erfüllt). (`0001`)

### Verworfen (adversarisch geprüft)

- **Skill „versioniert, nicht überschrieben (ADR-Hard-Rule, Modul 4)"** (von einem Reviewer als LOW gemeldet): steht **wortgleich** in der Modul-10-Vorlage (Schritt 6/Pflege) — template-konform, **kein Befund**.

### Prozess-Hinweis (Steering-Loop)

- Ein Reviewer halluzinierte im eigenen Auftragstext Dateinamen (`0002-cobra-cli…`, `0003-determinismus-strategie…`), die nie existierten, und korrigierte sich selbst (grep: 0 Treffer). Kein Doku-Befund — aber Signal, dass die Reviewer-Skill den konkreten Datei-Kontext fester verankern sollte, falls das Muster wiederkehrt (Modul 10 §Pflege).

## Regressions-Check

Geprüft, **keine Regression**: Die mehrfach verlinkte Index-Bezug-Spalte löst in allen Vorkommen byte-gleich auf (inkl. Umlaut-Anker `…für…` und Doppel-Bindestrich-Anker `…image---print-mk…`); die „nachgelagert"-Formulierungen drehen die Lastenheft-Abhängigkeitsrichtung nicht um; die neuen Konsequenzen (Pin-Staleness, static-Anker) führen keine Fehlbehauptung ein. Beleg: `make doc-check` grün.

## Kategorie-Summary

| Kategorie | Anzahl | IDs |
|---|---|---|
| HIGH | 0 | — |
| MEDIUM | 0 | — |
| LOW | 2 | NL1, NL2 |
| INFO | 4 | NI1–NI4 |

## Verdikt

**Re-Review bestanden.** Alle Erstlauf-Befunde sind gegen die realen Artefakte
verifiziert geschlossen; die Korrekturen haben keine Regression erzeugt
(`make doc-check` grün). Es verbleiben zwei nicht-blockierende LOW (Skill-Status-Wert,
In-Repo-Beleg für die d-check-Go-Aussage) plus INFO. Das ADR-Set ist aus allen drei
Linsen **acceptance-reif** für slice-002.

## Disposition (Implementer, 2026-06-21)

Behebung der Folgelauf-LOW als getrennter Implementer-Schritt. Beleg:
`make doc-check` grün.

| Finding | Aktion |
|---|---|
| NL1 | `.harness/skills/reviewer.md`: `Status` auf `Accepted` gesetzt (deckungsgleich mit der Modul-10-Skill-Vorlage). |
| NL2 | ADR-0001: d-check-Go-Aussage zusätzlich **in-Repo belegt** (`harness/conventions.md` §Adoptierte Konventions-Quellen), nicht nur per externem Link. |
| NI1–NI4 | INFO — keine Aktion (kosmetisch / erwartungsgemäß ohne Maschinen-Anker / positiv). |
