# Slice 201: AGENTS.md §5 als Index-Tabelle mit Regel-Auslagerung

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt. Er wechselt nur durch `git mv`.

**Welle:** ohne Welle — die Closure-Bedingung wäre die eigene DoD.

**Bezug:** Folge-Slice aus
[slice-200](../in-progress/slice-200-adoption-v6130-agents-und-matrix.md)
§3.1 (Aspekt 2); Ziel-Form `AGENTS.template.md` §5 (`v6.13.0`, Welle 149).
[`AC-QA-02`](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze).

**Berührte Spec-Stellen:** — · Der Slice berührt kein Spec-Stratum.

**Verantwortlich:** — bis zur Priorisierung.

**Autor:** Claude. **Datum:** 2026-09-29.

**Lerneintrag — Form:** geschärfte Regel.

---

## 1. Ziel und Abgrenzung

**Ziel:** [`AGENTS.md`](../../../../AGENTS.md) §5 trägt die **Index-Tabelle**
(`| # | Regel | Datei |` nach der Ziel-Form `v6.13.0`, Welle 149); jede Regel,
die über einen Satz hinausgeht oder eine eigene Begründung trägt, ist in
`harness/rules/<name>.md` ausgelagert, und die Tabellenzeile ist der
Kurzform-Zeiger darauf. **Substanz geht nicht verloren:** jede Begründung und
Durchsetzung steht nach wie vor im Repo — nur nicht mehr in §5.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Inhaltliche Änderungen an Regeln.** *Schicht-Abgrenzung*: der Slice
  konvertiert Form, er schärft oder streicht keine Regel — eine Aussage, die
  vorher da war, ist nachher woanders, nicht weg. Der Wortfolgen-Vergleich
  (§2) hält das fest.
- **§4 Harte Regeln.** *Bestand bleibt bewusst stehen*: §4 trägt die
  Falsch/Richtig-Lehrform in Prosa, die Welle 149 ausdrücklich in §5-Ort
  belässt; dort ist kein Umstelltisch.
- **Andere Dateien als `AGENTS.md` §5 und `harness/rules/`.** *Schicht-
  Abgrenzung*: Sensor und Makefile rühren sich nicht an.

## 2. Definition of Done

- [x] §5 trägt die Index-Tabelle; jede ausgelagerte Regel verweist in der
      Datei-Spalte auf `harness/rules/<name>.md`, Kurzregeln stehen vollständig
      in der Tabelle (`Datei`-Spalte: `—`).
- [x] **Wortfolgen-Vergleich als Gegenprobe:** die Menge der Aussagen von §5
      (Regeln, Begründungen, Durchsetzungs-Zusagen) ist vor/nach der Konversion
      unverändert — gemessen, nicht angenommen; Abweichungen sind benannt.
- [x] Unabhängiger Review, Report unter [`docs/reviews/`](../../../../docs/reviews/README.md);
      Closure-Notiz mit Lerneintrag; Register fortgeschritten; jedes Risiko
      trägt einen Ausgang.

`make gates` und `make verify` grün.

## 3. Plan (vor Code)

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| [`AGENTS.md`](../../../../AGENTS.md) §5 | refactor | Bullet-Liste → Index-Tabelle; Langregeln ausgelagert |
| `harness/rules/<name>.md` (je Langregel eine) | neu | Volltext der Regel — Wortlaut unbeschädigt, Zielort der Tabellen-Zeile |

## 4. Trigger

**Start** (`open` → `in-progress`): das WIP-Limit ist frei.

**Rückführungen — vorab benannt:**

- `in-progress` → `next` (zu groß): zeigt der Wortfolgen-Vergleich, dass sich
  mehr als drei Regeln nicht verlustfrei tabellieren lassen (Wechselwirkung
  mit Sensor-Zitaten), wird die Auslagerung dieser Regeln abgetrennt.
- `in-progress` → `open` (blockiert): entfällt — kein Blocker absehbar.

## 5. Closure-Trigger

DoD vollständig, `make gates` und `make verify` grün, Closure-Notiz mit
Lerneintrag.

## 6. Risiken und offene Punkte

- **Sensor- und Skill-Zitate treffen die neue Lage nicht** — Stellen, die
  „`AGENTS.md` §5, dritter Punkt" oder Absatz-Muster zitieren, verlieren ihre
  Adresse. — **Ausgang:** bei Closure (Abgleich der zitierenden Stellen vor dem
  `git mv`; `make doc-check` fängt Links, nicht Anker-Sätze).
- **Der Wortfolgen-Vergleich meldet Platzhalter-Rauschen** — Absicht der Kon-
  version ist Form-Wahrung, nicht Text-Gleichheit. — **Ausgang:** bei Closure.

## 7. Closure-Notiz

*(wird vor dem `git mv` nach `done/` gefüllt)*

## 8. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** HARNESS (Achsen 1, 2, 3 ✓).

**Vorgelagert — offene Beobachtungen sichten:** das Register wurde durchgegangen
(2026-09-29):
[`agents-md-hinkt-baseline-dod-item-hinterher`](../observations/BEO-HARNESS/agents-md-hinkt-baseline-dod-item-hinterher/observation.md)
(offen — die Konversion arbeitet genau diese Stelle ab, Befund wird gezählt);
keine weiteren Treffer für diesen Vorgang.

**Modus-Begründungsblock:** alle berührten Sub-Areas GF.
