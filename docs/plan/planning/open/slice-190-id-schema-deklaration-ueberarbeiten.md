# slice-190 — ID-Schema-Deklaration überarbeiten statt weiter flicken

**Welle:** ohne Welle.

**Bezug:** Ausgang von
[`BEO-HARNESS/adaption-korrigiert-repo-aussage`](../observations/BEO-HARNESS/adaption-korrigiert-repo-aussage/observation.md)
bei **3×** (*geplant*); ausgelöst durch
[slice-187](../done/wellenlos/slice-187-voll-abgleich-erstdurchgang-rest.md) §3.6.
[`AC-QA-02`](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze).

**Berührte Spec-Stellen:** — · Der Slice berührt kein Spec-Stratum.

**Verantwortlich:** —

**Autor:** Claude. **Datum:** 2026-09-08.

**Lerneintrag — Form:** wird bei Closure benannt.

---

## 1. Ziel und Abgrenzung

**Ziel:** Die ID-Schema-Deklaration ist **eine** aktuelle Aussage statt eines
akzeptierten Eintrags plus einer Kette von Korrektur-Einträgen. Danach sind
[`MR-020`](../../../../harness/conventions.md#mr-020) und
[`MR-023`](../../../../harness/conventions.md#mr-023) entbehrlich und können
aufgelöst werden.

**Das Muster, das den Slice auslöst:** Dreimal hat ein Adaptions-Eintrag eine
**Repo**-Aussage korrigiert statt einer Baseline-Regel — `slice-097`,
`slice-162`, `slice-187`. Unter dem Fork-Test ist jeder davon ein
**Rückbau-Kandidat**: Wer a-check forkt, übernimmt eine Adaption, die nichts
adaptiert, sondern nur eine veraltete eigene Zeile repariert. Zwei davon hängen
an derselben Zeile-Familie: der ID-Schema-Deklaration in
[`MR-000`](../../../../harness/conventions.md#mr-000).

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **[`MR-000`](../../../../harness/conventions.md#mr-000) inhaltlich ändern.** *Bestand bleibt bewusst stehen:* Der Eintrag
  ist akzeptiert und damit immutabel
  ([`harness/conventions.md`](../../../../harness/conventions.md) §Disziplin).
  Die Überarbeitung entsteht als **neuer** Eintrag, der die Deklaration als
  Ganzes ablöst — nicht als Edit am alten.
- **Die anderen aktiven `MR`-Einträge.** *Schicht-Abgrenzung:* Sie korrigieren
  Baseline-Regeln, nicht Repo-Aussagen, und fallen nicht unter den Fork-Test.
- **Ein Sensor auf „korrigiert eine Repo-Aussage".** *Es wäre ein anderer
  Vorgang:* Ob ein Eintrag eine Regel oder eine eigene Aussage korrigiert, ist
  ein Urteil über seinen Gegenstand, kein Match
  ([`AGENTS.md`](../../../../AGENTS.md) §3.7). Das Feld
  `Ersetzt-Baseline-Regel` macht es **sichtbar**, nicht prüfbar — und das ist
  bereits der Zustand.

## 2. Ausgangsmessung

*(beim Übergang nach `in-progress/` zu erheben — mindestens: wie viele aktive
`MR`-Einträge tragen `Ersetzt-Baseline-Regel: —`, und welche davon hängen an der
ID-Schema-Deklaration.)*

## 3. Umsetzung

*(entsteht mit der Arbeit)*

## 4. Definition of Done

- [ ] Die ID-Schema-Deklaration steht als **eine** aktuelle Aussage; wer sie
      nachschlägt, findet keinen Verweis auf eine Korrektur-Kette.
- [ ] [`MR-020`](../../../../harness/conventions.md#mr-020) und
      [`MR-023`](../../../../harness/conventions.md#mr-023) sind aufgelöst
      (`git mv` nach `conventions/done/`) oder ihre Fortgeltung ist begründet.
- [ ] [`BEO-HARNESS/adaption-korrigiert-repo-aussage`](../observations/BEO-HARNESS/adaption-korrigiert-repo-aussage/observation.md)
      trägt seinen Ausgang.
- [ ] Unabhängiger Review durchgeführt (Report unter [`docs/reviews/`](../../../reviews/README.md)).
- [ ] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [ ] Beobachtungs-Register fortgeschrieben.
- [ ] Jedes Risiko aus §7 trägt einen Ausgang.

`make gates` und `make verify` grün. Ein öffentlicher Vertrag ist berührt:
`harness/conventions.md` ist Rang 9 und deklariert die IDs, die jeder Commit
nennt.

## 5. Trigger

**Start** (`open` → `in-progress`): das WIP-Limit ist frei. Der Slice hat
**keine** Abhängigkeit von slice-188/189 — er berührt eine andere Datei-Familie.

**Rückführungen — vorab benannt:**

- `in-progress` → `next` (zu groß): Zeigt die Messung in §2 mehr als zwei
  betroffene Einträge, wird nach Datei-Familie zerlegt.
- `in-progress` → `open` (blockiert): Lässt sich die Deklaration nicht ablösen,
  ohne eine Kennung zu ändern, die in Commits steht, ist das ein eigener Vorgang
  — Kennungen sind repo-weit referenziert.

## 6. Closure-Trigger

DoD vollständig, `make gates` und `make verify` grün, Closure-Notiz mit
Lerneintrag.

## 7. Risiken und offene Punkte

- **Eine Ablösung ist selbst ein Eintrag, der eine Repo-Aussage korrigiert.**
  Der Slice kann das Muster reproduzieren, das er beenden soll — der Unterschied
  liegt darin, ob die neue Fassung **generisch** formuliert ist und darum nicht
  wieder veraltet. — **Ausgang:** <offen bis Closure>
- **Die Kennungen stehen in Commit-Messages und in `.d-check.yml`.** Eine
  geänderte Form bräche `make trace-check` rückwirkend.
  — **Ausgang:** <offen bis Closure>

## 8. Closure-Notiz

*(bei Closure auszufüllen)*

## 9. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** *(beim Übergang nach `in-progress/`
auszufüllen — berührt ist `HARNESS`.)*

**Vorgelagert — offene Beobachtungen sichten:** *(ebenso — **zwei** Quellen:
Register und der Review-Report des Vorgänger-Slice; und das Register ist hier
**nach `BEO-HARNESS`** zu lesen, nicht nach `BEO-PLAN`, siehe slice-187 §3.6.)*
