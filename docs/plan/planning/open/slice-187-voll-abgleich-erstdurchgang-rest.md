# slice-187 — Voll-Abgleich: der Rest des Erstdurchgangs (sieben Paare)

**Welle:** ohne Welle.

**Bezug:** Folge-Slice aus
[slice-186](../done/slice-186-voll-abgleich-restliche-paare.md) §7,
Risiko 1 (*eingetreten*).
[`AC-QA-02`](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze).

**Berührte Spec-Stellen:** `spec/lastenheft.md` und `spec/spezifikation.md`
sind unter den Paaren — die Kennungen benennt die Umsetzung.

**Verantwortlich:** —

**Autor:** Claude. **Datum:** 2026-09-08.

**Lerneintrag — Form:** wird bei Closure benannt.

---

## 1. Ziel und Abgrenzung

**Ziel:** Die sieben Ziel-Form-Paare, die der Erstdurchgang noch nicht getragen
hat, sind abgeglichen — je Paar mit einem der drei Ausgänge: übernommen ·
bewusst abweichend (mit Begründung) · ohne Befund.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Die acht bereits abgeglichenen Paare** (slice-185: fünf, slice-186: drei).
  Erledigt; ein zweiter Durchgang wäre Arbeit ohne Gegenstand.
- **Die elf Instanz-Vorlagen.** Dieselbe Abgrenzung wie in slice-185 §1.
- **Der *laufende* Abgleich bei künftigen Baseline-Sprüngen.** Der ist seit
  slice-185 in [`harness/conventions.md`](../../../../harness/conventions.md)
  §Baseline verankert und braucht keinen Slice — er ist ein Schritt der
  Migration. **Dieser Slice trägt nur den Erstdurchgang.**

## 2. Ausgangsmessung (2026-09-08)

Substanzielle Kandidaten je Paar — Bedienhinweise und Platzhalter
herausgerechnet, aber **weiterhin verrauscht**: Blockquote-Hinweise rutschen
durch, und die Zahl ist eine Reihenfolge, kein Befund.

| Paar | substanzielle Kandidaten |
|---|---|
| `AGENTS.md` (Rest nach den drei Befunden aus slice-186) | 20 |
| `spec/lastenheft.md` | 17 |
| `harness/conventions.md` | 16 |
| `.d-check.yml` | 16 |
| `.harness/skills/closure-note-reviewer.md` | 15 |
| `spec/spezifikation.md` | 12 |
| `docs/plan/planning/README.md` | 11 |

**107 zusammen.** Zum Vergleich: slice-185 hat fünf Paare mit 14 Kandidaten
getragen, slice-186 drei mit 18 plus drei vorab gemessene Befunde.

**Die zwei Spec-Straten sind die heikelsten.** Sie sind vertraglich bindend
(Rang 1 und 2), und eine Übernahme dort ist keine Doku-Pflege: `lastenheft.md`
trägt 19 grandfatherte `AC-*`, `spezifikation.md` sieben `SPEC-*` mit ADRs, die
darauf zeigen. Der Abgleich kann dort **Vorschläge** liefern; die Entscheidung
ist ein Change Request.

## 3. Umsetzung

*(entsteht mit der Arbeit)*

## 4. Definition of Done

- [ ] Alle sieben Paare sind abgeglichen; je Paar steht der Ausgang im Plan.
- [ ] Für die zwei Spec-Straten ist getrennt, was **Doku-Pflege** und was
      **Change Request** wäre — Letzteres wird benannt, nicht ausgeführt.
- [ ] Unabhängiger Review durchgeführt (Report unter [`docs/reviews/`](../../../reviews/README.md)).
- [ ] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [ ] Beobachtungs-Register fortgeschrieben.
- [ ] Jedes Risiko aus §7 trägt einen Ausgang.

`make gates` und `make verify` grün.

## 5. Trigger

**Start** (`open` → `in-progress`):
[slice-186](../done/slice-186-voll-abgleich-restliche-paare.md) liegt in
`done/` und das WIP-Limit ist frei.

**Rückführungen — vorab benannt:**

- `in-progress` → `next` (zu groß): **Wahrscheinlich.** Zwei Vorgänger haben je
  drei bis fünf Paare getragen; sieben mit 107 Kandidaten sind mehr. Zerlegt
  wird dann nach Rang: erst die Harness-Paare, dann die zwei Spec-Straten.
- `in-progress` → `open` (blockiert): Findet der Abgleich in einem Spec-Stratum
  eine Ziel-Form-Regel, deren Übernahme eine `AC-*` berührt, ist das ein Change
  Request und keine Entscheidung dieses Slice.

## 6. Closure-Trigger

DoD vollständig, `make gates` und `make verify` grün, Closure-Notiz mit
Lerneintrag.

## 7. Risiken und offene Punkte

- **Sieben Paare sind mehr, als die zwei Vorgänger je getragen haben.** Die
  Rückführung ist in §5 benannt und wahrscheinlich.
  — **Ausgang:** <offen bis Closure>
- **Die Spec-Straten könnten den Slice in einen Change Request kippen.** Eine
  Ziel-Form-Regel, die eine `AC-*` berührt, ist keine Doku-Pflege.
  — **Ausgang:** <offen bis Closure>
- **Die Vorauswahl bleibt verrauscht.** 107 Kandidaten sind keine 107 Befunde;
  wer die Zahl für einen Umfang hält, plant falsch.
  — **Ausgang:** <offen bis Closure>

## 8. Closure-Notiz

*(bei Closure auszufüllen)*

## 9. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** *(beim Übergang nach `in-progress/`
auszufüllen — die zwei Spec-Straten machen `SPEC` zur zweiten berührten
Sub-Area, anders als bei slice-185/186.)*

**Vorgelagert — offene Beobachtungen sichten:** *(ebenso.)*
