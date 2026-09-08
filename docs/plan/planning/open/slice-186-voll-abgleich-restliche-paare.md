# slice-186 — Voll-Abgleich: die zehn verbleibenden Ziel-Form-Paare

**Welle:** ohne Welle — die Closure-Bedingung wäre die eigene DoD.

**Bezug:** Folge-Slice aus
[slice-185](../done/wellenlos/slice-185-ziel-form-voll-abgleich.md) §7, Risiko 3
(*eingetreten*: 13 Paare sprengen die Größen-Regel).
[`AC-QA-02`](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze).

**Berührte Spec-Stellen:** — *(bei Umsetzung zu füllen: `spec/lastenheft.md`
und `spec/spezifikation.md` sind unter den Paaren.)*

**Verantwortlich:** —

**Autor:** Claude. **Datum:** 2026-09-08.

**Lerneintrag — Form:** wird bei Closure benannt.

---

## 1. Ziel und Abgrenzung

**Ziel:** Die zehn Ziel-Form-Paare, die
[slice-185](../done/wellenlos/slice-185-ziel-form-voll-abgleich.md) nicht mehr
getragen hat, sind abgeglichen — je Paar mit einem der drei Ausgänge:
übernommen · bewusst abweichend (mit Begründung) · ohne Befund.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Die fünf in slice-185 abgeglichenen Paare.** Erledigt; ein zweiter Durchgang
  wäre Arbeit ohne Gegenstand.
- **Die zwölf Instanz-Vorlagen.** Dieselbe Abgrenzung wie in slice-185 §1: kein
  einzelnes Gegenstück, Form geprüft durch `make doc-structure`.
- **Ein Sensor auf Satz-Deckung.** Urteil über zwei Formulierungen, kein Match
  ([`AGENTS.md`](../../../../AGENTS.md) §3.7).

## 2. Ausgangsmessung (2026-09-08, aus slice-185 übernommen)

Kandidaten je Paar aus der maschinellen Vorauswahl — **überwiegend Rauschen**
(Platzhalter, Bedienhinweise), die Zahl ist eine Reihenfolge, kein Befund:

| Paar | Kandidaten |
|---|---|
| `README.md` | *(neu zu erheben, siehe unten)* |
| `AGENTS.md` | 22 |
| `harness/conventions.md` | 18 |
| `spec/lastenheft.md` | 18 |
| `.d-check.yml` | 16 |
| `.harness/skills/closure-note-reviewer.md` | 16 |
| `docs/plan/planning/README.md` | 12 |
| `spec/spezifikation.md` | 12 |
| `harness/README.md` | 9 |
| `docs/plan/planning/in-progress/roadmap.md` | *(neu, siehe unten)* |

**Ein zehntes Paar** (Review zu slice-185, F-7): `roadmap.template.md` hat ein
Gegenstück in [`docs/plan/planning/in-progress/roadmap.md`](../in-progress/roadmap.md).
Die Paarbildung von slice-185 suchte es unter `docs/plan/planning/roadmap.md`
und fand nichts — die Roadmap liegt im Lifecycle-Verzeichnis.

**Reihenfolge:** Die Kandidaten-Zahl ist eine grobe Sortierung, kein Befund.
Der Indikator *„Gegenstück kürzer als Vorlage"*, der in slice-185 beim ersten
Treffer trug, greift hier bei **keinem** Paar — er hatte dort genau einen
echten Fall (den Reviewer-Skill), und der ist erledigt.
**`README.md` ist gegen `project-readme.template.md` zu messen** (2076 Zeichen),
nicht gegen `templates/README.md` — das ist der Index des Vorlagen-Verzeichnisses
und keine Ziel-Form (Review zu slice-185, F-1). a-checks `README.md` ist mit
9274 Zeichen **länger** als seine Vorlage; die 32 Kandidaten der Vorauswahl
stammen aus dem falschen Vergleich und sind neu zu erheben.

### 2.1 Für `AGENTS.md` stehen drei Befunde bereits fest

Sie sind nach slice-185 gemessen worden und ersparen dem Paar die Suche. Alle
drei betreffen **§5 Dokumentations-Regeln** (16 740 Zeichen, 18 Blöcke) — mit
Abstand den größten Abschnitt der Datei.

**(a) Drei nachgeschriebene `modul-05`-Sätze im Slice-Form-Block** (~600
Zeichen). Ein Wortfolgen-Test über alle 18 Blöcke gegen das vendorte Regelwerk
findet **fünf** wörtliche Treffer, alle im selben Block; **zwei** davon sind
legitime Zitate mit Anführungszeichen und Quelle (`modul-06` §Wann Arbeit eine
Welle braucht), **drei** stehen ohne beides und stammen aus `modul-05`
§Ziel-Form: Slice — *„ein Ausschluss ohne Grund ist eine Behauptung, keine
Grenze"*, *„wer später etwas mitnimmt … hat den Plan geändert, nicht ergänzt"*,
*„zerlegt, nicht gedehnt … eine von drei benannten Formen"*.
**Die anderen 17 Blöcke: null Treffer.**
*Warum es slice-183 entging:* Der Block wurde dort als a-check-eigen eingestuft
(*„sieben Punkte, jeder gegen den Bestand gemessen"*) — das stimmt für die
sieben Punkte; nachgeschrieben sind die **Rahmen-Sätze** darum. Er ist der
einzige Block, in dem beides gemischt ist.

**(b) Zwei Review-Urteilsregeln am falschen Ort** (**4045** Zeichen):
*Geltungsbereich einer Messung* (2688) und *Eine Mutations-Probe belegt erst,
wenn sie rot war* (1357). Beide sagen ausdrücklich **„Kein Sensor"** — sie sind
inferentielle Urteilsregeln, und `modul-08` §Welche Rolle braucht welche
Artefaktklasse weist genau diesen Fall der **Skill-Datei** zu, nicht dem
Briefing. Ihr Ort wäre
[`.harness/skills/reviewer.md`](../../../../.harness/skills/reviewer.md).

**(c) Die Slice-Form-Kopieranleitung** (**4285** Zeichen, ein Viertel von §5).
Sie ist eine Bedienungsanleitung für die Vorlage und steht in `AGENTS.md`, weil
die vendorte Ziel-Form nicht geändert werden darf und a-check keine eigene
Kopie führt. Ihr natürlicher Ort wäre
[`docs/plan/planning/README.md`](../README.md), wo die Slice-Ablage beschrieben
ist.

**Zusammen 8930 Zeichen**, die §5 auf gut die Hälfte brächten — **ohne dass eine
Regel verlorengeht**, nur an den Ort wandert, an dem sie gelesen wird. (b) und
(c) sind **Umzüge**, kein Rückbau; das ist beim Abgleich sauber zu trennen.

**Grenze dieser Messung:** Der Wortfolgen-Test sucht **wörtliche** Übernahme.
Eine sinngemäße Doppelung in anderen Worten sieht er nicht — das bleibt
Lese-Arbeit (slice-185 §2.3).

## 3. Umsetzung

*(entsteht mit der Arbeit)*

## 4. Definition of Done

- [ ] Alle zehn Paare sind abgeglichen; je Paar steht der Ausgang im Plan.
- [ ] Die drei für `AGENTS.md` **vorab gemessenen** Befunde (§2.1) sind
      aufgelöst — die drei `modul-05`-Sätze auf Zitat oder Zeiger, die zwei
      Urteilsregeln in den Reviewer-Skill, die Kopieranleitung nach
      `docs/plan/planning/README.md`. **Umzug, kein Rückbau:** nach dem
      Verschieben trägt der Zielort die Regel vollständig, und §5 einen Zeiger.
      Gegenprobe: §5 gemessen vor und nach dem Umzug.
- [ ] Jede Übernahme ist als solche kenntlich; jede bewusste Abweichung trägt
      ihre Begründung — eine Differenz ist nicht automatisch eine Lücke.
- [ ] Unabhängiger Review durchgeführt (Report unter [`docs/reviews/`](../../../reviews/README.md)).
- [ ] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [ ] Beobachtungs-Register fortgeschrieben.
- [ ] Jedes Risiko aus §7 trägt einen Ausgang.

`make gates` und `make verify` grün.

## 5. Trigger

**Start** (`open` → `in-progress`):
[slice-185](../done/wellenlos/slice-185-ziel-form-voll-abgleich.md) liegt in
`done/` und das WIP-Limit ist frei.

**Rückführungen — vorab benannt:**

- `in-progress` → `next` (zu groß): Braucht ein einzelnes Paar mehr als einen
  Durchgang — `spec/lastenheft.md` mit 84 146 Zeichen ist der Kandidat —, wird
  es abgetrennt.
- `in-progress` → `open` (blockiert): Findet der Abgleich eine Ziel-Form-Regel,
  die a-check bewusst **nicht** übernehmen will, ist das eine `MR`-Adaption und
  keine Nachtrags-Entscheidung dieses Slice.

## 6. Closure-Trigger

DoD vollständig, `make gates` und `make verify` grün, Closure-Notiz mit
Lerneintrag.

## 7. Risiken und offene Punkte

- **Zehn Paare könnten wieder zu groß sein.** slice-185 hat vier geschafft und
  das fünfte ohne Befund geschlossen; zehn ist mehr als das Doppelte — und mit
  §2.1 kommt für `AGENTS.md` ein Umzug von 8930 Zeichen dazu, der für sich
  genommen ein Slice wäre. — **Ausgang:** <offen bis Closure>
- **Ein Umzug kann eine Regel unterwegs verlieren.** (b) und (c) verschieben
  Text zwischen Dokumenten verschiedener Ränge; ein Zeiger, der die Aussage
  nicht trägt, ist schlechter als die Doppelung.
  — **Ausgang:** <offen bis Closure>
- **Die maschinelle Vorauswahl ist unscharf in beide Richtungen.** Sie meldet
  Rauschen — und sie kann eine echte Lücke verfehlen, wenn a-check dieselbe
  Aussage in anderen Worten trägt. — **Ausgang:** <offen bis Closure>

## 8. Closure-Notiz

*(bei Closure auszufüllen)*

## 9. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** *(beim Übergang nach `in-progress/`
auszufüllen.)*

**Vorgelagert — offene Beobachtungen sichten:** *(ebenso.)*
