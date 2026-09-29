# slice-190 — ID-Schema-Deklaration überarbeiten statt weiter flicken

**Welle:** ohne Welle.

**Bezug:** Ausgang von
[`BEO-HARNESS/adaption-korrigiert-repo-aussage`](../observations/BEO-HARNESS/adaption-korrigiert-repo-aussage/observation.md)
bei **3×** (*geplant*); ausgelöst durch
[slice-187](../done/wellenlos/slice-187-voll-abgleich-erstdurchgang-rest.md) §3.6.
[`AC-QA-02`](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze).

**Berührte Spec-Stellen:** — · Der Slice berührt kein Spec-Stratum.

**Verantwortlich:** Claude — gesetzt beim Übergang nach `in-progress/`.

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

**Vier** aktive `MR`-Einträge tragen `Ersetzt-Baseline-Regel: —`:
[`MR-019`](../../../../harness/conventions/MR-019-review-dod-opt-in.md),
[`MR-020`](../../../../harness/conventions/done/MR-020-adr-vorlage-generisch.md),
[`MR-023`](../../../../harness/conventions/done/MR-023-id-schema-beobachtungs-kennung.md),
[`MR-024`](../../../../harness/conventions/MR-024-historische-kern-drift-deklariert.md).
(zur Laufzeit der Messung standen [MR-020](../../../../harness/conventions/done/MR-020-adr-vorlage-generisch.md)/[MR-023](../../../../harness/conventions/done/MR-023-id-schema-beobachtungs-kennung.md) noch in `harness/conventions/`.)
Davon hängen **zwei** an der ID-Schema-Deklaration: [MR-020](../../../../harness/conventions/done/MR-020-adr-vorlage-generisch.md) und [MR-023](../../../../harness/conventions/done/MR-023-id-schema-beobachtungs-kennung.md)
(gegreppt über `id-schema` in den Eintrags-Dateien — zweimal verschieden
gezählt: Anker-Suche in conventions.md findet [MR-020](../../../../harness/conventions/done/MR-020-adr-vorlage-generisch.md)/[MR-023](../../../../harness/conventions/done/MR-023-id-schema-beobachtungs-kennung.md) als die einzigen
ID-Schema-Korrekturen der aktiven Menge).

## 3. Umsetzung

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| [`harness/conventions/MR-029-id-schema-deklaration-gesamt.md`](../../../../harness/conventions/MR-029-id-schema-deklaration-gesamt.md) | neu | Die vollständige, generisch formulierte Deklaration — löst [MR-020](../../../../harness/conventions/done/MR-020-adr-vorlage-generisch.md)/[MR-023](../../../../harness/conventions/done/MR-023-id-schema-beobachtungs-kennung.md) ab |
| [`harness/conventions.md`](../../../../harness/conventions.md) | update | Aktive/aufgelöste Tabellen; Zu-[`MR-000`](../../../../harness/conventions.md#mr-000)-Zeiger auf die aktuelle Fassung |
| [`harness/conventions/done/MR-020…`](../../../../harness/conventions/done/MR-020-adr-vorlage-generisch.md), [`MR-023`](../../../../harness/conventions/done/MR-023-id-schema-beobachtungs-kennung.md) | refactor | Reine `git mv`-Moves; interne Links auf die done/-Lage nachgezogen |

## 4. Definition of Done

- [x] Die ID-Schema-Deklaration steht als **eine** aktuelle Aussage; wer sie
      nachschlägt, findet keinen Verweis auf eine Korrektur-Kette.
- [x] [`MR-020`](../../../../harness/conventions.md#mr-020) und
      [`MR-023`](../../../../harness/conventions.md#mr-023) sind aufgelöst
      (`git mv` nach `conventions/done/`) oder ihre Fortgeltung ist begründet.
- [x] [`BEO-HARNESS/adaption-korrigiert-repo-aussage`](../observations/BEO-HARNESS/adaption-korrigiert-repo-aussage/observation.md)
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
  wieder veraltet. — **Ausgang:** *entfallen*, gestrichen mit Begründung:
  [MR-029](../../../../harness/conventions/MR-029-id-schema-deklaration-gesamt.md) deklariert die Formen generisch (versions- und schwellenfrei); die
  Gültigkeit hängt an conventions.md §Baseline, die jede Migration ohnehin
  mitführt. Eine inhaltliche Form-Änderung wäre ein neuer Eintrag — das ist der
  normale Adaptions-Weg, kein Korrektur-Rücklauf.
- **Die Kennungen stehen in Commit-Messages und in `.d-check.yml`.** Eine
  geänderte Form bräche `make trace-check` rückwirkend.
  — **Ausgang:** *entfallen*, gestrichen mit Begründung: die Formen wurden
  nicht geändert — die Konversion konsolidiert die Deklaration, die Kennungen
  (inkl. `slice-NNN` numerisch) bleiben wie deklariert.

## 8. Closure-Notiz

**Lerneintrag — Form: geschärfte Regel.** *Eine Deklaration, die über
Korrektur-Einträge gepflegt wird, wird durch einen Ablöse-Eintrag geschlossen,
der sie als Ganzes und generisch formuliert — die Korrektur-Kette selbst ist
das Symptom, dass die Deklaration ihren Ort gewechselt hat.* Gemessen:
zweimal korrigierte Einträge an einer Zeilen-Familie ([MR-020](../../../../harness/conventions/done/MR-020-adr-vorlage-generisch.md)/[MR-023](../../../../harness/conventions/done/MR-023-id-schema-beobachtungs-kennung.md)), beide
unter dem Fork-Test; [MR-029](../../../../harness/conventions/MR-029-id-schema-deklaration-gesamt.md) fasst sie als Ganzes. Die generische Formulierung
(„die jeweils aktuell vendorte Fassung") statt einer Versionsnummer ist die
Maßnahme gegen die Wiederholung.

**Was hat funktioniert:** die Ausgangsmessung (§2) trennte die ID-Schema-Familie
([MR-020](../../../../harness/conventions/done/MR-020-adr-vorlage-generisch.md)/[MR-023](../../../../harness/conventions/done/MR-023-id-schema-beobachtungs-kennung.md)) von den übrigen Zeiger-freien Einträgen ([MR-019](../../../../harness/conventions/MR-019-review-dod-opt-in.md)/[MR-024](../../../../harness/conventions/MR-024-historische-kern-drift-deklariert.md)) — die
Ablösung blieb auf die Familie beschränkt.

**Was ging anders als geplant:** die Link-Tiefen der bewegten Einträge — drei
Nachzugs-Runden, gefangen von `doc-check` (`repo-escape`, `target-missing`).
Die Tiefe der Auslagerungs-Orte (harness/rules/ bzw. conventions/done/) ist
nicht die Herkunfts-Tiefe; die Regeln stehen im Selbsttest-Flow, nicht in der
Vorab-Prüfung.

**Steering-Loop-Eintrag:** — *(nichts verkörpert; die Deklaration selbst ist
das Artefakt, und der Fork-Test bleibt das Instrument, das künftige
Korrektur-Einträge an ihr messen wird.)*

**Beobachtungs-Register (`../observations/`):**
`BEO-HARNESS/adaption-korrigiert-repo-aussage/` — Ausgang *verkörpert*
(`seit slice-190`, [MR-029](../../../../harness/conventions/MR-029-id-schema-deklaration-gesamt.md) als die Deklaration): die drei Vorfälle
(slice-097, slice-162, slice-187) sind durch die Ablösung gegenstandslos.

**Folge-Slices:** — *(keine.)*

**Risiken aus §6:** jedes mit genau einem Ausgang — siehe §6.

**Drei Paarungen:** Anker — der Steering-Loop-Eintrag liegt in
[`harness/conventions/MR-029-id-schema-deklaration-gesamt.md`](../../../../harness/conventions/MR-029-id-schema-deklaration-gesamt.md) ·
Folge-Slice — keine · Register —
`BEO-HARNESS/adaption-korrigiert-repo-aussage` ist verkörpert (Anker am
state.md), die 3×-Kette ist geschlossen.

## 9. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** *(beim Übergang nach `in-progress/`
auszufüllen — berührt ist `HARNESS`.)*

**Vorgelagert — offene Beobachtungen sichten:** *(ebenso — **zwei** Quellen:
Register und der Review-Report des Vorgänger-Slice; und das Register ist hier
**nach `BEO-HARNESS`** zu lesen, nicht nach `BEO-PLAN`, siehe slice-187 §3.6.)*
