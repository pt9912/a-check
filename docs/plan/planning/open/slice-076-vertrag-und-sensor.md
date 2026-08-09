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
[`tools/verify-slice-form.sh:22`](../../../../tools/verify-slice-form.sh) schreibt selbst:

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
[`tools/verify-ac-form.sh`](../../../../tools/verify-ac-form.sh) prüft, ob vier fettgedruckte
Bezeichner vorkommen — ein AC mit viermal „beliebiger Text" besteht.

**Die zwei Fälle lösen sich gegenläufig auf**, und genau das macht sie zu einem Slice: bei `F-6`
zieht die **Doku** nach, bei `F-3` ist zu **entscheiden**, ob der Sensor die Form messen kann oder
der Vertrag seine Zusage zurücknimmt.

## 2. Betroffene Module

Zwei Schichten:

1. **Doku** — [`docs/plan/planning/README.md`](../README.md) (`F-6`), und je nach Entscheid
   [`AGENTS.md`](../../../../AGENTS.md) §5 / [`harness/conventions.md`](../../../../harness/conventions.md) (`F-3`).
2. **`tools/`** — [`verify-ac-form.sh`](../../../../tools/verify-ac-form.sh), falls der Entscheid
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
  [slice-075](../in-progress/slice-075-sensor-messgroesse.md).
- **Die Grandfathering-Schwelle für ACs verschieben.** Bestand bleibt Bestand.

## 5. DoD

- [ ] Die Zusage in [`docs/plan/planning/README.md`](../README.md) deckt sich mit dem, was
      `make verify` prüft. Beleg: Fixture-Slice mit drei Schichten läuft grün, und die Doku sagt
      genau das.
- [ ] Der `F-3`-Entscheid ist **getroffen und begründet** in der Closure-Notiz, und die gewählte
      Seite ist umgesetzt. Beleg: die zweite Probe aus §3 im gewählten Ausgang.
- [ ] `make gates` und `make verify` grün — **Ausgabe in eine Datei**, Exit-Code getrennt geprüft,
      nie in eine Pipe.

## 6. Closure-Notiz

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice,
[`AGENTS.md`](../../../../AGENTS.md) §5; `make verify` prüft das.)_

## 7. Sub-Area-Modus

Alle berührten Sub-Areas GF.
