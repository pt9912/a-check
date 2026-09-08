# slice-185 — Voll-Abgleich Ziel-Form ↔ Gegenstück: was vier Delta-Analysen nicht finden konnten

**Welle:** ohne Welle — die Closure-Bedingung wäre die eigene DoD
(Baseline-Regelwerk `modul-06-roadmap.md` §Wann Arbeit eine Welle braucht).

**Bezug:** [`AC-QA-02`](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)
(Harness-Integrität). Keine aktive ADR wird berührt.

**Berührte Spec-Stellen:** — (Harness-Artefakte ohne Vertragsberührung).

**Verantwortlich:** —

**Autor:** Claude. **Datum:** 2026-09-08.

**Lerneintrag — Form:** geschärfte Regel.

---

## 1. Ziel und Abgrenzung

**Ziel:** Jede vendored Ziel-Form mit **genau einem** Gegenstück im Repo ist
einmal vollständig gegen dieses Gegenstück abgeglichen — und der Abgleich ist
als wiederkehrender Schritt beim Baseline-Sprung verankert.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Die zwölf Instanz-Vorlagen** (Slice, ADR, Carveout, Review-Report,
  Beobachtung, Welle, Archiv-Stubs, …). Sie haben kein *einzelnes* Gegenstück,
  sondern viele; ihre Form prüft `make doc-structure` bereits über Muster. Ein
  Abgleich „Vorlage gegen alle Instanzen" wäre ein anderer Vorgang mit einem
  anderen Werkzeug.
- **Die Ziel-Formen ohne Gegenstück** (`reconciliation.template.md`,
  `project-readme.template.md`, `gate.template.md`, …). a-check führt sie
  bewusst nicht — `reconciliation.md` etwa gibt es nicht, weil es keinen
  Brownfield-Bootstrap gab ([`AGENTS.md`](../../../../AGENTS.md) §5,
  Kopieranleitung Punkt 2). Bestand, der bewusst stehen bleibt.
- **Ein Sensor auf Satz-Deckung.** *„Trägt das Gegenstück diese Aussage?"* ist
  ein Urteil über zwei Formulierungen, kein Match — dieselbe Grenze wie
  [`AGENTS.md`](../../../../AGENTS.md) §3.7. Ein Grob-Vergleich per Wortfolge
  liefert überwiegend Rauschen; das ist in §2.3 gemessen, nicht vermutet.
- **Produkt-Code.** Der Slice rührt `internal/` nicht an.

## 2. Ausgangsmessung (2026-09-08)

### 2.1 Der Gegenstand

| | |
|---|---|
| Ziel-Formen unter `.harness/baseline/v6.5.0/templates/` | **26** |
| davon mit **genau einem** Gegenstück im Repo | **14** |
| Normtext dieser 14 (ohne Bedienhinweise und Platzhalter) | **66 741** Zeichen |

**Zwei Gegenstücke sind kürzer als ihre Vorlage** — bei einem ausgefüllten
Dokument der schärfste Einstiegspunkt:

| Paar | Ziel-Form | Gegenstück |
|---|---|---|
| `.harness/skills/reviewer.md` | 5294 | **3835** |
| `README.md` | 13 433 | **9230** |

### 2.2 Der erste Treffer trägt sofort

`.harness/skills/reviewer.md` fehlen **drei HIGH-Kategorien**, die die Ziel-Form
führt:

| HIGH-Kategorie | Vorkommen in a-checks Skill |
|---|---|
| **Norm nur im Template-Kommentar** — eine Regel steht im `<!-- -->`-Block und nirgends sonst | **0** |
| **Kommentar trägt keine der Kommentar-Klassen** — verworfene Alternative, abwesender Text, abgebrochener Satz | **0** |
| **Zustandsfeld trägt Chronik** — eine `Stand`/`Status`-Zelle erzählt die Entstehung | **0** |

Jede trägt in der Ziel-Form denselben Nachsatz: *„Kein Gate fängt das."* Sie
stehen dort, **weil** kein Sensor sie sieht.

**Das ist bereits teuer geworden.** Zwei Register-Einträge zählen genau diese
Klassen —
[`chronik-in-gelesenen-dateien`](../observations/BEO-HARNESS/chronik-in-gelesenen-dateien/observation.md)
(2×) und
[`hard-rule-37-ohne-sensor`](../observations/BEO-HARNESS/hard-rule-37-ohne-sensor/observation.md)
(2×). Gefunden hat sie jedes Mal ein Review; **gesucht hat keines danach**,
weil der Skill die Kategorien nicht führt. Sie fielen als Nebenprodukt an.

### 2.3 Warum es vier Migrationen nicht fanden

Jede Baseline-Migration ist eine **Delta**-Analyse: slice-135
(`v5.12.0`→`v6.0.0`), slice-161, slice-164, slice-174 (`v6.2.0`→`v6.5.0`). Ein
Delta findet, was sich **ändert**.

**Gemessen an einem Beispielsatz** — dem Arbeitsteilungs-Satz für `AGENTS.md` §4
(*„Diese Tabelle listet auf; definiert wird hier nichts"*): Er steht **seit
`v5.12.0`** unverändert in der Ziel-Form. In keinem der vier Deltas taucht er
auf, und a-check hat ihn nie übernommen.

**Ein Voll-Abgleich Vorlage-gegen-Gegenstück hat es nie gegeben.** Das ist die
Lücke — nicht ein Versehen einer einzelnen Migration.

**Grenze des maschinellen Vergleichs, gemessen:** Ein Grob-Abgleich per
Wortfolge über `AGENTS.md` meldet **30** Sätze ohne Entsprechung, überwiegend
Rauschen — Template-Hinweise, Platzhalter und Bedienkommentare, die beim
Kopieren bestimmungsgemäß verschwinden. Der Abgleich ist damit **Lese-Arbeit
mit maschineller Vorauswahl**, kein Lauf.

## 3. Umsetzung

*(entsteht mit der Arbeit)*

## 4. Definition of Done

- [ ] Die drei HIGH-Kategorien stehen in
      [`.harness/skills/reviewer.md`](../../../../.harness/skills/reviewer.md),
      in a-checks Sprache und mit den Fundstellen, die sie im Bestand hatten.
- [ ] Die übrigen **13** Paare sind abgeglichen; je Paar steht das Ergebnis —
      übernommen, bewusst abweichend (mit Begründung) oder ohne Befund.
- [ ] Der Abgleich ist als **wiederkehrender Schritt** verankert: Wer die
      Baseline hebt, fährt ihn — nicht nur das Delta.
- [ ] Unabhängiger Review durchgeführt (Report unter [`docs/reviews/`](../../../reviews/README.md)).
- [ ] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [ ] Beobachtungs-Register fortgeschrieben.
- [ ] Jedes Risiko aus §7 trägt einen Ausgang.

`make gates` und `make verify` grün. Ein öffentlicher Vertrag ist berührt:
`AGENTS.md` ist Rang 8, der Reviewer-Skill steuert jeden Review.

## 5. Trigger

**Start** (`open` → `in-progress`): WIP-Limit frei.

**Rückführungen — vorab benannt:**

- `in-progress` → `next` (zu groß): Zeigen die ersten Paare, dass der Abgleich
  je Datei einen eigenen Durchgang braucht, wird nach Dateien zerlegt — der
  Reviewer-Skill zuerst, weil sein Befund schon steht.
- `in-progress` → `open` (blockiert): Findet der Abgleich eine Ziel-Form-Regel,
  die a-check bewusst **nicht** übernehmen will, ist das eine
  `MR`-Adaption und kein Nachtrag; der Slice wartet auf diese Entscheidung.

## 6. Closure-Trigger

DoD vollständig, `make gates` und `make verify` grün, Closure-Notiz mit
Lerneintrag. Danach Archivierung als wellenloser Slice
([`AGENTS.md`](../../../../AGENTS.md) §6).

## 7. Risiken und offene Punkte

- **Der Abgleich ist Lese-Arbeit und damit unvollständig prüfbar.** Was ein
  Leser übersieht, sieht auch der nächste Lauf nicht — dieselbe Grenze, die
  §3.7 für sich benennt. — **Ausgang:** <offen bis Closure>
- **Ein Nachtrag kann eine bewusste Abweichung überschreiben.** a-check hat
  Ziel-Form-Regeln begründet nicht übernommen (`MR`-Adaptionen); ein
  Abgleich, der jede Differenz als Lücke liest, macht daraus stillen Rückbau.
  — **Ausgang:** <offen bis Closure>
- **13 Paare in einem Slice könnten die Größen-Regel sprengen.** Die
  Rückführung ist in §5 benannt, kostet aber einen Durchgang.
  — **Ausgang:** <offen bis Closure>

## 8. Closure-Notiz

*(bei Closure auszufüllen)*

## 9. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** *(beim Übergang nach `in-progress/`
auszufüllen.)*

**Vorgelagert — offene Beobachtungen sichten:** *(ebenso.)*
