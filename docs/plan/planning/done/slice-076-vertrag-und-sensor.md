# slice-076 — Wo Vertrag und Sensor klaffen, zieht eine Seite nach

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Deckt:** `F-3` und `F-6` aus dem
[Review-Report `welle-12`](../../../reviews/2026-08-09-welle-12-unabhaengig.md) (Gruppe B).
**Bezug:** Roadmap-Zeile *Aktuelle Welle* in der [Roadmap](../in-progress/roadmap.md).

---

## 1. Auslöser

**Mechanismus: eine Doku behauptet eine Prüfung, die der Sensor nicht leistet.** Anders als in
Gruppe A misst hier nichts falsch — die Lücke sitzt zwischen dem, was zugesagt, und dem, was
gemessen wird. Und sie zeigt in **beide Richtungen**:

**`F-6` — der Sensor ist ehrlich, die Doku ist es nicht.**
`tools/verify-slice-form.sh:22` schreibt selbst:

> NICHT geprueft: „hoechstens zwei Schichten" aus B-1 — was eine Schicht ist, ist eine
> Ermessensfrage ueber Modul-Grenzen; ein Zaehler darueber waere Schein-Genauigkeit. Bleibt Sache
> des Reviews.

[`docs/plan/planning/README.md:8`](../README.md) ordnet dieselbe Zeile — *„höchstens drei
DoD-Punkte, höchstens zwei Schichten"* — pauschal `make verify` zu. Die Ausnahme des Sensors ist
sauber begründet; die Doku hat sie nur nie übernommen.

**`F-3` — der Vertrag verlangt mehr, als der Sensor prüft.**
[`AGENTS.md`](../../../../AGENTS.md) §5 und
[`harness/conventions.md`](../../../../harness/conventions.md) §Anforderungs-Anlege-Prozess
verlangen die drei Pfade **im Given/When/Then-Stil**.
`tools/verify-ac-form.sh` prüft, ob vier fettgedruckte
Bezeichner vorkommen — ein AC mit viermal „beliebiger Text" besteht.

**Die zwei Fälle lösen sich gegenläufig auf**, und genau das macht sie zu einem Slice: bei `F-6`
zieht die **Doku** nach, bei `F-3` ist zu **entscheiden**, ob der Sensor die Form messen kann oder
der Vertrag seine Zusage zurücknimmt.

## 2. Betroffene Module

Zwei Schichten:

1. **Doku** — [`docs/plan/planning/README.md`](../README.md) (`F-6`), und je nach Entscheid
   [`AGENTS.md`](../../../../AGENTS.md) §5 / [`harness/conventions.md`](../../../../harness/conventions.md) (`F-3`).
2. **`tools/`** — `verify-ac-form.sh`, falls der Entscheid
   zum Sensor geht.

## 3. Auszuführende Gates

`make verify`, `make gates`.

**Offener Entscheid zu `F-3` — vor dem Bau zu treffen.** Zwei Wege, und der Slice nimmt keinen
vorweg:

| Weg | Was er kostet |
|---|---|
| **Sensor zieht nach** — Given/When/Then im AC-Rumpf prüfen | Given/When/Then ist Prosa; ein Regex darauf ist genau die Schein-Genauigkeit, die `F-6` beim Schichten-Zähler zu Recht vermeidet |
| **Vertrag zieht zurück** — die Form als Review-Sache deklarieren, wie es der Schichten-Fall bereits tut | die Zusage wird kleiner, aber wahr; Präzedenz steht im selben Repo |

Der Schichten-Fall (`F-6`) ist damit **auch die Argumentationsvorlage** für `F-3`: dort wurde
bereits entschieden, dass eine Ermessensfrage nicht maschinell wird. Ob dasselbe für
Given/When/Then gilt, ist die zu treffende Entscheidung — sie gehört in die Closure-Notiz.

**Negativ-Proben:**

| Probe | Erwartung |
|---|---|
| Slice mit drei benannten Schichten | unverändert **grün** — die Doku sagt danach, dass hier nicht geprüft wird |
| neues AC mit viermal „beliebiger Text" | je nach Entscheid rot (Sensor zog nach) **oder** grün bei korrigierter Zusage |
| Bestand | unverändert grün |

## 4. Was bewusst nicht getan wird

- **Den Schichten-Zähler doch bauen.** Die Begründung im Sensor steht seit slice-052 und ist nicht
  widerlegt; dieser Slice korrigiert die Doku, nicht die Entscheidung.
- **`F-4`/`F-13`** — dort misst der Sensor eine Ersatzgröße; das ist
  [slice-075](../done/slice-075-sensor-messgroesse.md).
- **Die Grandfathering-Schwelle für ACs verschieben.** Bestand bleibt Bestand.

## 5. DoD

- [x] Die Zusage in [`docs/plan/planning/README.md`](../README.md) deckt sich mit dem, was
      `make verify` prüft. Beleg: Fixture-Slice mit **drei** benannten Schichten → Exit 0, kein
      Befund — und die Doku nennt „höchstens zwei Schichten" jetzt ausdrücklich als Review-Sache.
- [x] Der `F-3`-Entscheid ist **getroffen und begründet** in der Closure-Notiz, und die gewählte
      Seite ist umgesetzt. Beleg: `verify-ac-form` unverändert (Exit 0), die Zusage in
      [`harness/conventions.md`](../../../../harness/conventions.md) nennt jetzt, was maschinell
      fällt und was nicht.
- [x] `make gates` und `make verify` grün — **Ausgabe in eine Datei**, Exit-Code getrennt geprüft,
      nie in eine Pipe.

## 6. Closure-Notiz

**Geliefert:** zwei Zusagen decken sich jetzt mit dem, was gemessen wird —
[`docs/plan/planning/README.md`](../README.md) für die Slice-Form,
[`harness/conventions.md`](../../../../harness/conventions.md) für die AC-Form. Kein Sensor wurde
geändert; beide trugen die Wahrheit bereits im Kopfkommentar.

**Der `F-3`-Entscheid: der Vertrag zieht zurück, nicht der Sensor.** Given/When/Then ist Prosa.
Ein Regex auf die drei Wörter prüfte ihr *Vorkommen*, nicht die Form — „Given/When/Then" als bloße
Aufzählung bestünde, ein sauber formulierter Satz ohne die Reizwörter fiele durch. Das ist
dieselbe Schein-Genauigkeit, die
`tools/verify-slice-form.sh` beim Schichten-Zähler
ausdrücklich ablehnt. Die Entscheidung war damit **nicht offen im luftleeren Raum**: das Repo hatte
sie für den strukturell gleichen Fall schon getroffen, nur nie auf den zweiten übertragen.

Der Sensor deklarierte seine Grenze bereits — aber nur für den *Inhalt* („ob der Satz hinter
Boundary: wirklich einen Randfall beschreibt"). Die *Stil*-Grenze fehlte; sie ist jetzt als zweite
Stufe daneben benannt.

**Lerneintrag — Form: geschärfte Regel.** Als Prüfsatz: *Klaffen Zusage und Sensor, ist zuerst zu
fragen, welche Seite recht hat — nicht, wie der Sensor die Zusage einholt.*

**Die Ursache** ist bei beiden Findings dieselbe und läuft der Intuition zuwider: **beide Sensoren
waren ehrlich, beide Dokus zu großzügig.** `verify-slice-form` schreibt seit slice-052 selbst
„NICHT geprueft: höchstens zwei Schichten … bleibt Sache des Reviews"; `verify-ac-form` seit
slice-054 eine analoge Grenze. Die Dokus haben diese Ausnahmen nie übernommen — sie entstanden
zuerst, als Beschreibung der *Absicht*, und wurden nicht nachgezogen, als der Bau die Absicht
bewusst beschnitt. Ein Review, das nur den Sensor liest, findet nichts; erst der Abgleich mit der
Zusage zeigt die Lücke — und genau so hat der unabhängige Lauf sie gefunden.

**Zwei beobachtbare Closure-Kriterien:**

1. Ein Slice mit drei benannten Schichten läuft grün, und `planning/README.md` sagt genau das —
   vorher behauptete die Zeile ein Gate, das es nie gab.
2. Kein Sensor-Verhalten hat sich geändert: `make verify` meldet dieselben Zahlen wie vorher. Die
   Korrektur war ausschließlich eine Rücknahme von Zusagen auf das gemessene Maß.

**Folge-Slices:** keine. Damit ist Gruppe B des Reviews vollständig (`F-3`, `F-4`, `F-6`, `F-13`).

## 7. Sub-Area-Modus

Alle berührten Sub-Areas GF.
