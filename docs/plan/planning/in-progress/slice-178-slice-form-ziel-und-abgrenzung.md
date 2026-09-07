# slice-178 — Slice-Form: §1 *Ziel und Abgrenzung*, §8-Titel, Schritt 4

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `make slice-mv` ([`AGENTS.md`](../../../../AGENTS.md)
§3.3/§5).

**Welle:** [welle-15](../welle-15-regelwerk-v650-migration.md)

**Bezug:** [slice-174](../done/slice-174-regelwerk-v650-delta-analyse.md)
§3.2 T-3 — Etappe **E** des Schnitts in §3.4.

**Berührte Spec-Stellen:** — *(keine)* — Planungs-Form ohne
Vertragsberührung.

**Verantwortlich:** — *(noch nicht priorisiert)*

**Autor:** Claude (Opus 5), im Auftrag des Maintainers. **Datum:**
2026-09-07.

---

## 1. Ziel und Abgrenzung

Die Kopieranleitung für neue Slices nennt §1 als *Ziel und Abgrenzung* mit
Out-of-Scope je Punkt und Begründung in vier Klassen, §8 unter seinem neuen
Titel *Sub-Area-Prüfungen und Modus-Begründung* — und der Minimal Agent
Workflow trägt die Schritt-Hälfte derselben Regel: die Plan-Ausgabe in
Schritt 4 nennt Out-of-Scope, und was der Lauf darüber hinaus mitnimmt, ist
eine Plan-Änderung vor dem Code.

**Nicht in diesem Slice**, je Punkt mit Grund:

- **Nachrüsten der bestehenden Slices in `open/`** — Bestand bleibt bewusst
  stehen: die Ziel-Form gilt für neue Slices, und ein Sensor gegen
  unentschiedenen Altbestand wäre ein Fehlalarm.
- **Ein Sensor auf die Out-of-Scope-Form** — wäre ein anderer Vorgang
  (Gate-Schicht statt Planungs-Form) und braucht erst einen Bestand, an dem
  er kalibriert werden kann.
- **Andere Etappen von [welle-15](../welle-15-regelwerk-v650-migration.md)** —
  Schicht-Abgrenzung.

## 2. Analyse

### 2.1 Was `v6.5.0` an der Slice-Form ändert

| Stelle | Änderung |
|---|---|
| **§1** | heißt *Ziel und Abgrenzung* und trägt eine Out-of-Scope-Liste **je Punkt mit Begründung** — *„ein Ausschluss ohne Grund ist eine Behauptung, keine Grenze"* |
| **§8** | heißt *Sub-Area-Prüfungen und Modus-Begründung*; der Titel trägt beide Hälften, weil nur die zweite bedingt ist. **Der Abschnitt entfällt nie** |
| `modul-09` | die **Schritt-Hälfte**: die Plan-Ausgabe in Schritt 4 nennt Out-of-Scope, und was der Lauf darüber hinaus mitnimmt, ist eine Plan-Änderung vor dem Code |

Die vier Ausschluss-Klassen sind ausdrücklich ein **Suchraster, keine
Ausfüll-Liste**: *„Ein Slice mit einem echten Ausschluss ist besser als einer
mit vier erfundenen."*

### 2.2 Drei Befunde am eigenen Bestand — gemessen, nicht übernommen

Die Kopieranleitung in [`AGENTS.md`](../../../../AGENTS.md) §5 stimmte an drei
Stellen nicht mehr mit dem überein, was a-check tut:

1. **„die vier Felder streichen … `Welle:`"** — falsch. **Acht** Slices im
   Bestand tragen ein `**Welle:**`-Feld, fünf davon mit
   [welle-15](../welle-15-regelwerk-v650-migration.md), drei mit *„ohne Welle"*.
   Ebenso die *drei Paarungen*: bei wellenlosen Slices trägt die Closure sie,
   bei Slices mit Welle die Welle-Closure — gestrichen ist keines von beiden.
   Zu streichen sind **zwei** Felder, nicht vier.
2. **Die Nummerierung verschiebt sich** — a-check führt zwischen Ziel und DoD
   einen **Analyse-Abschnitt**, den die Ziel-Form nicht kennt. Dadurch DoD
   `§2 → §4` und Sub-Area `§8 → §9`. Stand nirgends; wer die Anleitung wörtlich
   nahm, kopierte falsch.
3. **§9 trug den alten Titel** *Sub-Area-Modus* statt *Sub-Area-Prüfungen und
   Modus-Begründung* — die Verkürzung, gegen die die neue Fassung ausdrücklich
   argumentiert.

**Der Analyse-Abschnitt ist keine Abweichung, die wegzuräumen wäre.** Er hält
die Messung, auf der die DoD steht — dieselbe Sitzung hat viermal gezeigt, was
passiert, wenn eine Behauptung im Plan später als Messung gelesen wird. Er
bleibt und ist jetzt in der Anleitung benannt.

## 3. Umsetzung

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| [`AGENTS.md`](../../../../AGENTS.md) §5 | update | Kopieranleitung als **sechs nummerierte Punkte** statt eines Fließtexts: was zu ergänzen, was zu streichen ist, wie die Nummerierung verschiebt, §1- und §9-Titel samt Suchraster |
| [`AGENTS.md`](../../../../AGENTS.md) §6 Schritt 4 | update | die Schritt-Hälfte: die Plan-Ausgabe nennt Out-of-Scope; wer mehr mitnimmt, ändert den Plan |
| [`harness/README.md`](../../../../harness/README.md) §Minimal agent workflow | update | derselbe Schritt, mit Verweis statt Wiederholung |

**Nicht angefasst:** die bestehenden Slices in `open/`. Ihre §1 tragen kein
*Ziel und Abgrenzung* — Bestand bleibt bewusst stehen, die Form gilt für neue
Slices (§1).

## 4. Definition of Done

- [x] [`AGENTS.md`](../../../../AGENTS.md) §5 nennt beim Kopieren der Ziel-Form
      §1 als *Ziel und Abgrenzung* samt der vier Out-of-Scope-Klassen und dem
      neuen §8-Titel.
- [x] [`AGENTS.md`](../../../../AGENTS.md) §6 trägt die Schritt-Hälfte: die
      Plan-Ausgabe in Schritt 4 nennt Out-of-Scope, Abweichung ist eine
      Plan-Änderung vor dem Code.
- [x] Unabhängiger Review durchgeführt (Report unter [`docs/reviews/`](../../../reviews/README.md)).
- [x] `make gates` grün.
- [x] `make verify` grün.
- [x] Jedes Risiko trägt einen Ausgang.

## 5. Trigger

**Start** (`open` → `in-progress`): der neue Stand liegt vendored
([slice-175](../done/slice-175-etappe-a-vendoring-v650.md) in `done/`),
Maintainer-Freigabe, WIP-Limit frei.

**Rückführungen:** wächst der Umfang über die zwei Punkte hinaus, zurück nach
`next/`. Ändert sich der adoptierte Stand erneut, zurück nach `open/`.

## 6. Closure-Trigger

DoD vollständig, `make gates`/`make verify` grün, Closure-Notiz geschrieben.
Der Slice trägt ein `**Welle:**`-Feld und archiviert **mit seiner Welle**
([`AGENTS.md`](../../../../AGENTS.md) §6).

## 7. Risiken und offene Punkte

- *Die Kopieranleitung wächst, ohne dass ein Lauf sie liest — sie steht in
  `AGENTS.md` §5, und die Ziel-Form liegt vendored daneben* — **Ausgang:**
  weiter offen → Beobachtungs-Register,
  [`BEO-HARNESS/baseline-regel-nie-erwogen-weil-bestand-sie-verletzt`](../observations/BEO-HARNESS/baseline-regel-nie-erwogen-weil-bestand-sie-verletzt/observation.md)
  (2×). Das Risiko ist **eingetreten**, nur nicht bei der Anleitung selbst,
  sondern an dem, was sie beschreibt: Drei ihrer Aussagen stimmten nicht mehr
  mit dem Bestand überein (§2.2), und aufgefallen ist das erst, als jemand sie
  gegen den Bestand hielt. Kein Sensor sieht das — eine Anleitung, die falsch
  beschreibt, was ohnehin gemacht wird, bricht kein Gate.

## 8. Closure-Notiz

**Lerneintrag — Form: geschärfte Regel** (die Kopieranleitung beschreibt jetzt,
was a-check *tut*, nicht was es einmal vorhatte).

- **Was hat funktioniert:** Die Anleitung **gegen den Bestand gehalten**, statt
  nur die neue Ziel-Form einzuarbeiten. Der Auftrag war, zwei Titel und einen
  Workflow-Schritt nachzuziehen; gefunden wurden drei Stellen, an denen die
  Anleitung seit Längerem etwas anderes sagte als die Praxis — allen voran
  „`Welle:` streichen", während acht Slices im Bestand es tragen.

- **Was ging anders als geplant:** Ich hatte erwartet, einen Abschnitt zu
  ergänzen. Herausgekommen ist eine Umschreibung: aus einem Fließtext mit
  „fünftens" wurden sechs nummerierte Punkte. Die alte Form konnte nicht
  falsch *auffallen* — man liest sie als Absatz, nicht als Liste, und ein
  Absatz hat keine Zeile, die man gegen den Bestand halten kann.

- **Steering-Loop-Eintrag — geschärfte Regel:** Eine Kopieranleitung ist eine
  **Aussage über den eigenen Bestand** und altert wie jede andere. Sie gehört
  bei jeder Baseline-Migration nicht nur *ergänzt*, sondern **gegen den Bestand
  geprüft** — Punkt für Punkt, mit einem Kommando je Punkt. — liegt in
  `AGENTS.md §5` (die sechs Punkte, jeder gegen den Bestand gemessen).

- **Beobachtungs-Register (`../observations/`):** keine neue Beobachtung.
  [`baseline-regel-nie-erwogen-weil-bestand-sie-verletzt`](../observations/BEO-HARNESS/baseline-regel-nie-erwogen-weil-bestand-sie-verletzt/observation.md)
  bleibt bei 2× — die Klasse passt („die Regel fällt auf, wenn jemand sie
  liest"), aber hier war es keine *Baseline*-Regel, sondern a-checks eigene
  Anleitung. Ein Beleg dort würde die Klasse verwässern; der Fall steht in §2.2
  dieses Slice.

- **Folge-Slices:** keine.

- **Risiken aus §7:** eines, mit genau einem Ausgang — *weiter offen* →
  Register.

- **Drei Paarungen:** trägt die Welle-Closure — der Slice hat ein
  `**Welle:**`-Feld.

## 9. Sub-Area-Modus

**Vorgelagert — Sub-Area-Wahl prüfen:** eine Sub-Area berührt —
**Harness-Einstieg** (`AGENTS.md`, `harness/README.md`), Achsen 1,2,3,
deklariert in
[`conventions.md`](../../../../harness/conventions.md#modus-deklaration-pro-sub-area).
Das **Planungs-Harness** ist Gegenstand der Anleitung, aber nicht geändert:
kein Slice-Plan wurde angefasst.

**Vorgelagert — offene Beobachtungen sichten:** Register am 2026-09-07
durchgegangen. Einer einschlägig:
[`baseline-regel-nie-erwogen-weil-bestand-sie-verletzt`](../observations/BEO-HARNESS/baseline-regel-nie-erwogen-weil-bestand-sie-verletzt/observation.md)
(2×) — mit Ausgang in §7, ohne Beleg (Begründung in §8).

**Alle berührten Sub-Areas GF** — kein Begründungsblock nötig.
