# slice-188 — Voll-Abgleich: Gate-Konfiguration und Closure-Skill

**Welle:** ohne Welle.

**Bezug:** Folge-Slice aus
[slice-187](../done/slice-187-voll-abgleich-erstdurchgang-rest.md) §1
(Schicht-Abgrenzung).
[`AC-QA-02`](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze).

**Berührte Spec-Stellen:** — · Der Slice berührt kein Spec-Stratum.

**Verantwortlich:** —

**Autor:** Claude. **Datum:** 2026-09-08.

**Lerneintrag — Form:** wird bei Closure benannt.

---

## 1. Ziel und Abgrenzung

**Ziel:** Die zwei Paare der Gate-/Skill-Schicht sind gegen ihre vendorte
Ziel-Form abgeglichen — je Paar mit einem der drei Ausgänge: übernommen ·
bewusst abweichend (mit Begründung) · ohne Befund.

| Paar | Kandidaten (Instrument: slice-187 §2) |
|---|---|
| [`.d-check.yml`](../../../../.d-check.yml) | 42 |
| [`.harness/skills/closure-note-reviewer.md`](../../../../.harness/skills/closure-note-reviewer.md) | 25 |

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Die zwei Spec-Straten.** *Es wäre ein anderer Vorgang:* Rang 1 und 2,
  Change-Request-Charakter. Eigener Folge-Slice
  ([slice-189](../open/slice-189-voll-abgleich-spec-straten.md)).
- **Die elf bereits abgeglichenen Paare.** Erledigt (slice-185/186/187).
- **Neue oder geänderte `d-check`-Module.** *Schicht-Abgrenzung:* Der Abgleich
  prüft, ob die **Konfiguration** die Ziel-Form trägt — nicht, ob a-check ein
  Modul mehr einschalten sollte. Ein neues Modul ist ein eigener Slice mit
  Pin-Frage.
- **Die elf Instanz-Vorlagen.** Kein einzelnes Gegenstück; ihre Form prüft
  `make doc-structure` über Muster.

## 2. Ausgangsmessung

Instrument, Parameter und Vorgehen stehen in
[slice-187](../done/slice-187-voll-abgleich-erstdurchgang-rest.md) §2 —
sie werden hier **nicht wiederholt**, sondern zitiert; der Aufruf ist derselbe.
Die zwei Zahlen oben stammen aus **demselben** Lauf wie die dortige Tabelle.

**Die Besonderheit dieses Paares:** In `.d-check.yml` trägt der **Kommentar**
den Text, nicht die Prosa — das Instrument strippt dort `#`, statt die Zeile zu
überspringen. 42 Kandidaten sind darum überwiegend Bedienhinweise der Vorlage
(*„aktivieren, sobald …"*), die beim Ausfüllen bestimmungsgemäß verschwinden.
Die Zahl ist eine Reihenfolge, kein Befund.

## 3. Umsetzung

*(entsteht mit der Arbeit)*

## 4. Definition of Done

- [ ] Beide Paare sind abgeglichen; je Paar steht der Ausgang **mit seiner
      Prüf-Ebene** im Plan.
- [ ] Für `.d-check.yml` ist getrennt, was **Bedienhinweis der Vorlage** und was
      **Normtext** ist — nur Letzteres ist ein Kandidat für einen Befund.
- [ ] Unabhängiger Review durchgeführt (Report unter [`docs/reviews/`](../../../reviews/README.md)).
- [ ] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [ ] Beobachtungs-Register fortgeschrieben.
- [ ] Jedes Risiko aus §7 trägt einen Ausgang.

`make gates` und `make verify` grün.

## 5. Trigger

**Start** (`open` → `in-progress`):
[slice-187](../done/slice-187-voll-abgleich-erstdurchgang-rest.md) liegt
in `done/` und das WIP-Limit ist frei.

**Rückführungen — vorab benannt:**

- `in-progress` → `next` (zu groß): 67 Kandidaten liegen über dem, was
  slice-186 getragen hat (39), aber unter slice-187 (124). Braucht `.d-check.yml`
  allein mehr als einen Durchgang, wird der Skill abgetrennt.
- `in-progress` → `open` (blockiert): Verlangt eine Ziel-Form-Regel ein Modul,
  das a-check nicht gepinnt hat, ist das ein Pin-Vorgang und keine Entscheidung
  dieses Slice.

## 6. Closure-Trigger

DoD vollständig, `make gates` und `make verify` grün, Closure-Notiz mit
Lerneintrag.

## 7. Risiken und offene Punkte

- **`.d-check.yml` ist Konfiguration, keine Prosa.** Ein „Befund" dort kann eine
  Verhaltensänderung eines Gates sein, keine Doku-Pflege — und ein geändertes
  Gate ohne ADR verletzt [`AGENTS.md`](../../../../AGENTS.md) §3.6.
  — **Ausgang:** <offen bis Closure>
- **Der Closure-Skill steuert jede Closure.** Eine Übernahme dort wirkt sofort
  auf den nächsten Slice; ein Fehler darin fällt erst beim übernächsten auf.
  — **Ausgang:** <offen bis Closure>

## 8. Closure-Notiz

*(bei Closure auszufüllen)*

## 9. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** *(beim Übergang nach `in-progress/`
auszufüllen — berührt ist `GATE`, für den Skill zusätzlich `REVIEW`.)*

**Vorgelagert — offene Beobachtungen sichten:** *(ebenso — **zwei** Quellen:
Register und der Review-Report des Vorgänger-Slice, siehe slice-187 §9.)*
