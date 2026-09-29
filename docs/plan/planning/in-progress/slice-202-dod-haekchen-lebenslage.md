# slice-202 — Attestierungs-Sensor für vorgezogene Häkchen

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt. Er wechselt nur durch `git mv`.

**Welle:** ohne Welle — die Closure-Bedingung wäre die eigene DoD.

**Bezug:** Ausgang von
[`BEO-GATE/attestierung-vor-dem-vorgang`](../observations/BEO-GATE/attestierung-vor-dem-vorgang/observation.md)
bei **3×** (*geplant*; slice-169, slice-197, slice-200).
[`AC-QA-02`](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze).

**Berührte Spec-Stellen:** — · Der Slice berührt kein Spec-Stratum.

**Verantwortlich:** Claude — gesetzt beim Übergang nach `in-progress/`.

**Autor:** Claude. **Datum:** 2026-09-29.

**Lerneintrag — Form:** neuer Sensor.

---

## 1. Ziel und Abgrenzung

**Ziel:** Ein DoD-Häkchen in einem Slice-Plan, der in `open/` oder `next/`
liegt, ist mechanisch gefangen — ein `forbid-pattern` (Modul `structure`,
konfiguriert in [`.d-check.yml`](../../../../.d-check.yml), im `gates`-Aggregat)
meldet `- [x]` auf `docs/plan/planning/{open,next}/**/slice-*.md`. Damit ist
die Attestierungs-Lücke (drei Auftreten: slice-169, slice-197, slice-200)
geschlossen: das Häkchen entsteht nicht mehr davor, ohne dass ein Lauf es
sieht.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Häkchen in `in-progress/`** — *Bestand bleibt bewusst stehen*: dort sind
  Häkchen der Normalfall (DoD wächst mit der Arbeit); der Geltungsbereich ist
  bewusst `open/` und `next/`.
- **Prosa-Attestierungen außerhalb von DoD-Listen** („nach dem `git mv`
  geprüft" im Fließtext). *Es wäre ein anderer Vorgang*: die Klasse ist
  inferentiell; dieser Slice fängt die formbare Hälfte.

## 2. Definition of Done

- [x] Das `forbid-pattern` ist konfiguriert und im `gates`-Aggregat wirksam;
      die Gegenprobe ist in **beiden** Richtungen belegt — Positiv (eine
      `[x]`-Zeile in `open/` meldet `section-forbidden`) und Negativ (offenes
      Häkchen meldet nicht) — plus Scoping-Kontrolle (`in-progress/` bleibt
      grün) und Zitat-Kontexte nach SL-004 (Inline-Code und Code-Block bleiben
      grün, empirisch gemessen); alles verankert in `dcheck-phrase-selftest`
      als Muster 3.
- [x] `make gates` und `make verify` grün.
- [ ] Unabhängiger Review, Report unter [`docs/reviews/`](../../../../docs/reviews/README.md).
- [ ] Closure-Notiz mit Steering-Loop-Lerneintrag; Register fortgeschritten;
      jedes Risiko aus §6 trägt einen Ausgang.

## 3. Plan (vor Code)

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| [`.d-check.yml`](../../../../.d-check.yml) (Modul `structure`) | update | `forbid-pattern` auf `open/`- und `next/`-Slices; `dcheck-phrase-selftest` um die Fixture ergänzen |

## 4. Trigger

**Start** (`open` → `in-progress`): das WIP-Limit ist frei.

**Rückführungen — vorab benannt:**

- `in-progress` → `next` (zu groß): entfällt — der Slice ist ein Sensor plus
  Fixture.
- `in-progress` → `open` (blockiert): kann das Modul `structure` den
  Pfad-Scope nicht einschränken, ist die Alternativ-Wiring (eigenes Target)
  ein eigener Entscheid — zurück, um ihn zu fällen.

## 5. Closure-Trigger

DoD vollständig, `make gates` und `make verify` grün, Closure-Notiz mit
Lerneintrag.

## 6. Risiken und offene Punkte

- **Das Muster trifft legitime Häkchen** — z. B. Plan-Stellen, die eine
  `[x]`-Form *zitieren*. — **Ausgang:** bei Closure (Fixture mit zitiertem
  Muster in den Selbsttest, wie [`SL-004`](../observations/README.md) es
  verlangt).

## 7. Closure-Notiz

**Lerneintrag — Form: neuer Sensor.** *Häkchen an die Lebenslage binden:* ein
`- [x]` in einem Slice in `open/`/`next/` attestiert einen Vorgang, der noch
aussteht — seit slice-169 drei Mal aufgetreten und hier mechanisch gefangen
(struktur-Bedingung 7, `forbid-pattern`, Scoping auf `open/`/`next/`; die
Häufig-Form in `in-progress/`/`done/` bleibt legitim).
liegt in `.d-check.yml` (structure, Bedingung 7) plus
`tools/dcheck-phrase-selftest.sh` (Muster 3, drei Kontrollen).
Auslöser: `BEO-GATE/attestierung-vor-dem-vorgang`
(slice-169, slice-197, slice-200 — 3×).

**Was hat funktioniert:** die Mutations-Probe vor der Selbsttest-Erweiterung
(der rote Lauf mit `section-forbidden` als Meldung) — sie lieferte den
Grund-Code, bevor die Asserts geschrieben wurden, und die Zitat-Kontext-Probe
(SL-004) zeigte, dass Inline-Code und Code-Blöcke natürliche Fixture-Fälle
sind, die grün bleiben.

**Was ging anders als geplant:** die Verdrahtung lief **vor dem Claim** — der
`git mv` nach `in-progress/` folgte erst auf die fertige Umsetzung. Die
Reihenfolge ist nachgezogen (Claim → Closure-Arbeit → done); die Lektion
steht hier, statt still verschwiegen zu werden: der Claim ist die
Arbeits-Reihenfolge, nicht ihre Nachschrift.

**Steering-Loop-Eintrag:** Sensor neu: DoD-Häkchen an die Lebenslage binden
— liegt in `.d-check.yml` (structure, Bedingung 7). Auslöser:
`BEO-GATE/attestierung-vor-dem-vorgang` (slice-169, slice-197, slice-200 — 3×).

**Beobachtungs-Register (`../observations/`):** `BEO-GATE/attestierung-vor-dem-vorgang/`
— Ausgang *verkörpert* (seit slice-202, Sensor Bedingung 7).

**Folge-Slices:** — *(keine.)*

**Risiken aus §6:** jedes mit genau einem Ausgang — siehe §6.

**Drei Paarungen:** Anker — der Steering-Loop-Eintrag liegt in
`.d-check.yml` mit `seit slice-202` · Folge-Slice — keine · Register —
attestierung-vor-dem-vorgang ist verkörpert, der Stand trägt den Anker.

## 8. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** GATE (Achsen 1, 2, 3 ✓).

**Vorgelagert — offene Beobachtungen sichten:** das Register wurde durchgegangen
(2026-09-29):
[`attestierung-vor-dem-vorgang`](../observations/BEO-GATE/attestierung-vor-dem-vorgang/observation.md)
— **dieser Slice ist ihr Ausgang** (3×, *geplant*); keine weiteren Treffer für
diesen Vorgang.

**Modus-Begründungsblock:** alle berührten Sub-Areas GF.
