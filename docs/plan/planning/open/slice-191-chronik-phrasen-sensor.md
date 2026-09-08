# slice-191 — Phrasen-Sensor gegen Chronik in lebenden Dateien

**Welle:** ohne Welle.

**Bezug:** Ausgang von
[`BEO-HARNESS/chronik-in-gelesenen-dateien`](../observations/BEO-HARNESS/chronik-in-gelesenen-dateien/observation.md)
bei **3×** (*geplant*); ausgelöst durch
[slice-187](../in-progress/slice-187-voll-abgleich-erstdurchgang-rest.md) §3.6.
[`AC-QA-02`](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze).

**Berührte Spec-Stellen:** — · Der Slice berührt kein Spec-Stratum.

**Verantwortlich:** —

**Autor:** Claude. **Datum:** 2026-09-08.

**Lerneintrag — Form:** wird bei Closure benannt.

---

## 1. Ziel und Abgrenzung

**Ziel:** Die **häufige Schreibweise** von Chronik in lebenden Dateien ist
maschinell gefangen — ein `forbid-pattern` im Modul `structure`, konfiguriert in
[`.d-check.yml`](../../../../.d-check.yml), im `gates`-Aggregat. Gegenstand sind
die Dateien, die **jeder Lauf liest**: `AGENTS.md`, `harness/*.md`,
`docs/plan/planning/README.md`.

**Die Grenze gehört in denselben Slice wie der Sensor.** Er prüft eine
**Phrase**, nicht die Klasse *Chronik*; ob ein Satz Chronik ist, bleibt ein
Urteil ([`AGENTS.md`](../../../../AGENTS.md) §3.7). Die Sensor-Datei unter
[`harness/sensors/`](../../../../harness/sensors/) sagt das, bevor jemand ihn
für den ganzen Wächter hält — dieselbe Falle wie bei
[`BEO-GATE/muster-trifft-nur-die-haeufige-schreibweise`](../observations/BEO-GATE/muster-trifft-nur-die-haeufige-schreibweise/observation.md).

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Zeitdokumente** (`docs/plan/planning/**`, `docs/reviews/**`,
  `docs/plan/adr/**`). *Schicht-Abgrenzung:* Dort ist Chronik **richtig** — ein
  Slice-Plan, ein Review-Report und eine ADR sind Belege eines Vorgangs. Ein
  Muster, das dort feuert, wäre abgeschaltet, bevor es zweimal gelaufen ist.
- **Eine Erweiterung auf sinngemäße Chronik.** *Es wäre ein anderer Vorgang:*
  Das ist die Klasse, nicht ihre Schreibweise, und für sie gibt es keinen
  Sensor.
- **Rückbau von Chronik im Bestand.** *Bestand bleibt bewusst stehen:* Die
  bekannten Stellen sind mit slice-103, slice-182 und slice-187 behoben; was der
  neue Lauf zusätzlich findet, ist Gegenstand **dieses** Slice — was er nicht
  findet, ist keine Aufgabe, die er erzeugt.

## 2. Ausgangsmessung

*(beim Übergang nach `in-progress/` zu erheben. Erwartung aus slice-187:
`grep -n "Bis slice-" AGENTS.md` liefert **null** Treffer — der Bestand ist
sauber, und ein Sensor, der auf null Treffern eingeschaltet wird, braucht eine
**Mutations-Probe**, sonst ist er grün ohne Gegenstand.)*

## 3. Umsetzung

*(entsteht mit der Arbeit)*

## 4. Definition of Done

- [ ] Das Muster ist in [`.d-check.yml`](../../../../.d-check.yml) konfiguriert
      und hängt im `gates`-Aggregat.
- [ ] Die **Mutations-Probe** ist gefahren und war **rot**, mit genannter
      Meldung: eine eingefügte Chronik-Zeile in einer lebenden Datei wird
      gemeldet, dieselbe Zeile in einem Slice-Plan **nicht**.
- [ ] Die Sensor-Datei unter [`harness/sensors/`](../../../../harness/sensors/)
      nennt die Grenze: Phrase statt Klasse, und welche Umformulierung entkommt.
- [ ] Unabhängiger Review durchgeführt (Report unter [`docs/reviews/`](../../../reviews/README.md)).
- [ ] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [ ] Beobachtungs-Register fortgeschrieben.
- [ ] Jedes Risiko aus §7 trägt einen Ausgang.

`make gates` und `make verify` grün.

## 5. Trigger

**Start** (`open` → `in-progress`): das WIP-Limit ist frei. Keine Abhängigkeit
von slice-188/189/190.

**Rückführungen — vorab benannt:**

- `in-progress` → `next` (zu groß): Zeigt die Probe, dass das Muster im Bestand
  massenhaft feuert, ist der Rückbau ein eigener Slice und das Muster wartet auf
  ihn — ein Gate, das den Bestand bricht, wird abgeschaltet statt befolgt.
- `in-progress` → `open` (blockiert): Kann `structure` kein `forbid-pattern` auf
  eine ganze Datei (statt auf eine Sektion) legen, ist es ein CR an `d-check`
  und keine Konfiguration.

## 6. Closure-Trigger

DoD vollständig, `make gates` und `make verify` grün, Closure-Notiz mit
Lerneintrag.

## 7. Risiken und offene Punkte

- **Ein Phrasen-Sensor kann für den Wächter der Regel gehalten werden.** Genau
  das ist die Klasse, die er nur zur Hälfte fängt — und die Verwechslung ist im
  Repo belegt. — **Ausgang:** <offen bis Closure>
- **Das Muster kann in Zitaten feuern.** Ein Slice-Plan, der die Regel
  *erklärt*, zitiert die verbotene Wendung; liegt er im Geltungsbereich, meldet
  der Lauf gegen einen legitimen Satz. — **Ausgang:** <offen bis Closure>
- **Der Bestand ist sauber, die Prüfmenge also leer.** Ein Prüfer mit leerer
  Prüfmenge meldet grün und ist damit nicht „unbenutzt", sondern unkalibriert —
  deshalb die Mutations-Probe in der DoD.
  — **Ausgang:** <offen bis Closure>

## 8. Closure-Notiz

*(bei Closure auszufüllen)*

## 9. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** *(beim Übergang nach `in-progress/`
auszufüllen — berührt sind `GATE` (die Konfiguration) und `HARNESS` (die
geprüften Dateien).)*

**Vorgelagert — offene Beobachtungen sichten:** *(ebenso — **zwei** Quellen:
Register und der Review-Report des Vorgänger-Slice; das Register ist **nach
allen berührten Kürzeln** zu lesen, siehe slice-187 §3.6.)*
