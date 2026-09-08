# slice-193 — Der Gate-Index steht einmal

**Welle:** ohne Welle.

**Bezug:** Folge-Slice aus
[slice-192](../in-progress/slice-192-baseline-v660-vendoring.md) §1 — die eine
inhaltliche Neuerung des Sprungs auf `v6.6.0`.
[`AC-QA-02`](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze).

**Berührte Spec-Stellen:** — · Der Slice berührt kein Spec-Stratum.

**Verantwortlich:** —

**Autor:** Claude. **Datum:** 2026-09-08.

**Lerneintrag — Form:** wird bei Closure benannt.

---

## 1. Ziel und Abgrenzung

**Ziel:** Der Gate-Index steht **einmal** — in
[`harness/README.md`](../../../../harness/README.md) §Sensors.
[`AGENTS.md`](../../../../AGENTS.md) §4 trägt die **Regel** (kein behauptetes
Gate ohne Deckung) und den **Zeiger**, nicht die Liste. Die drei Prüfer, die
heute auf zwei Indizes zeigen, zeigen danach auf einen.

**Der Grund ist nicht Ordnungsliebe** — er steht in `grundlagen-harness-dateien.md`
(`v6.6.0`): Beide Dateien liegen in **jedem** Lauf-Kontext, ein zweiter Index
wird also pro Lauf **zweimal bezahlt**, und er läuft auseinander, weil die
Pflicht, ihn nachzuziehen, nirgends steht.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Neue Gates oder geänderte Verträge.** *Schicht-Abgrenzung:* Der Slice
  bewegt eine Deklaration; welche Zusage ein Target trägt, entscheidet er nicht.
- **Die Aufteilung Gates / Nicht-Gates.** *Bestand bleibt bewusst stehen:*
  `harness/README.md` führt beide Blöcke bereits mit dem Kriterium *urteilen*
  gegen *bewegen · messen · sagen*. Was aus §4 kommt, wird dort **eingeordnet**,
  nicht neu sortiert.
- **Der Voll-Abgleich der Ziel-Formen.** *Ein Folge-Slice übernimmt es:*
  [slice-188](../open/slice-188-voll-abgleich-gate-und-skill.md) /
  [slice-189](../open/slice-189-voll-abgleich-spec-straten.md).

## 2. Ausgangsmessung (2026-09-08)

| Gegenstand | Ist-Stand |
|---|---|
| [`AGENTS.md`](../../../../AGENTS.md) §4 | 7094 Zeichen, **41** Tabellenzeilen |
| [`harness/README.md`](../../../../harness/README.md) §Sensors | 9512 Zeichen, **28** Tabellenzeilen, dazu §Nicht-Gates |
| `targets` in [`.d-check.yml`](../../../../.d-check.yml) | `doc-tables: [AGENTS.md, harness/README.md]`, `authority: AGENTS.md` |
| `mentions` in [`.d-check.yml`](../../../../.d-check.yml) | `documents: ["AGENTS.md"]` — jede Sensor-Datei ist dort genannt |
| `structure`-Zellengrenze | `cell-max-chars` auf der Spalte *Zweck* in `AGENTS.md` §4 |
| Target-Zellen mit Argument in der Code-Span | **null**, in beiden Tabellen |

**Die Differenz 41 gegen 28 ist der eigentliche Umfang:** `AGENTS.md` §4 führt
Targets, die in `harness/README.md` gar nicht stehen. Welche das sind, ist die
erste Handlung des Slice — und sie ist maschinell zu erheben, nicht zu schätzen.

## 3. Umsetzung

*(entsteht mit der Arbeit)*

## 4. Definition of Done

- [ ] [`AGENTS.md`](../../../../AGENTS.md) §4 trägt **keine Tabelle** mehr,
      sondern die Regel und den Zeiger; jedes dort bisher gelistete Target ist
      in [`harness/README.md`](../../../../harness/README.md) eingeordnet —
      **belegt durch eine Differenz-Messung vorher/nachher**, nicht durch
      Augenschein.
- [ ] Die drei Konfigurationen zeigen auf **einen** Index: `targets.doc-tables`,
      `targets.authority`, `mentions.documents`. Die Zellengrenze auf der
      *Zweck*-Spalte entfällt oder wandert mit.
- [ ] Eine **Mutations-Probe** je umgezogenem Prüfer war **rot**, mit genannter
      Meldung: ein Target ohne Eintrag und ein Eintrag ohne Target.
- [ ] Unabhängiger Review durchgeführt (Report unter [`docs/reviews/`](../../../reviews/README.md)).
- [ ] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [ ] Beobachtungs-Register fortgeschrieben.
- [ ] Jedes Risiko aus §7 trägt einen Ausgang.

`make gates` und `make verify` grün. Zwei öffentliche Verträge sind berührt:
`AGENTS.md` ist Rang 8, `harness/README.md` Rang 9.

## 5. Trigger

**Start** (`open` → `in-progress`):
[slice-192](../in-progress/slice-192-baseline-v660-vendoring.md) liegt in `done/` und
das WIP-Limit ist frei.

**Rückführungen — vorab benannt:**

- `in-progress` → `next` (zu groß): Zeigt die Differenz-Messung mehr als ein
  Dutzend Targets ohne Gegenstück, wird der Umzug vom Konfigurations-Umbau
  getrennt.
- `in-progress` → `open` (blockiert): Verlangt `targets` zwei Autoritäts-Dateien
  und kann nur eine, ist die Ablösung ein CR an `d-check` — dann bleibt die
  Doppelung stehen, sichtbar als benannte Abweichung.

## 6. Closure-Trigger

DoD vollständig, `make gates` und `make verify` grün, Closure-Notiz mit
Lerneintrag.

## 7. Risiken und offene Punkte

- **Ein Target kann beim Umzug verschwinden.** 41 gegen 28 Zeilen: Wer die
  Tabelle streicht, ohne die Differenz zu messen, verliert Deklarationen still —
  und `make doc-targets` prüft danach gegen einen Index, in dem sie fehlen.
  — **Ausgang:** <offen bis Closure>
- **`mentions` hängt an `AGENTS.md`.** Zieht die Konfiguration um, ohne dass die
  Sensor-Dateien in `harness/README.md` genannt sind, meldet der Lauf gegen den
  Bestand — oder, schlimmer, er meldet nichts, weil die Prüfmenge leer wird.
  — **Ausgang:** <offen bis Closure>
- **Die Zusage „ein Index" ist selbst eine Deklaration ohne Sensor.** Nichts
  hindert einen künftigen Lauf daran, wieder eine Tabelle in `AGENTS.md` zu
  schreiben. — **Ausgang:** <offen bis Closure>

## 8. Closure-Notiz

*(bei Closure auszufüllen)*

## 9. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** *(beim Übergang nach `in-progress/`
auszufüllen — berührt sind `HARNESS` und `GATE`.)*

**Vorgelagert — offene Beobachtungen sichten:** *(ebenso — **zwei** Quellen:
Register nach **allen** berührten Kürzeln und der Review-Report des
Vorgänger-Slice.)*
