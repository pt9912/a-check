# slice-077 — Zwei Status-Aussagen, die der Bestand widerlegt

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Deckt:** `F-10` und `F-15` aus dem
[Review-Report `welle-12`](../../../reviews/2026-08-09-welle-12-unabhaengig.md) (Gruppe C).
**Bezug:** Roadmap-Zeile *Aktuelle Welle* in der [Roadmap](../in-progress/roadmap.md).

---

## 1. Auslöser

**Mechanismus: eine Zahl im Text, die niemand nachgezählt hat.** Kein Sensor ist betroffen, kein
Gate wird rot — der Text sagt etwas, das der Bestand daneben widerlegt.

**`F-10` — [`AGENTS.md:194`](../../../../AGENTS.md)** verschärft das Baseline-Limit:

> **WIP-Limit = 1.** Es liegt **genau ein** Slice in `in-progress/` …

Die Baseline (`modul-05:44`) sagt *„WIP-Limit pro Implementer = 1"* — ein **Maximum**. „Genau ein"
macht daraus zusätzlich ein Minimum und verbietet damit den Leerlauf, den das Repo regelmäßig hat:
nach jedem Slice-Abschluss steht `in-progress/` leer, und die Roadmap führt zeitweise „keine
laufende Welle" als ausdrücklich zulässigen Zustand.

**`F-15` — [`Makefile`](../../../../Makefile)** beschreibt im Kommentarblock über `verify`
*„die drei Teil-Sensoren"* und begründet auf *„drei unabhaengige Fragen"*. Das Rezept führt
**vier** aus; die vierte kam mit `verify-slice-links` (slice-060) hinzu, ohne die Mengenangabe
nachzuziehen. Dieselbe Zahl stand auch in der Sitzungsnotiz falsch (dort als „fünf") — die Angabe
war nie nachgezählt, nur weitergeschrieben.

## 2. Betroffene Module

Eine Schicht: **Doku/Deklaration** — [`AGENTS.md`](../../../../AGENTS.md) §5 und der
Kommentarblock über `verify` im [`Makefile`](../../../../Makefile). Kein Sensor, kein Rezept.

## 3. Auszuführende Gates

`make gates` (enthält `doc-check` und `gate-consistency`), `make verify`.

**Negativ-Proben** — hier ist der Bestand selbst die Probe, weil beide Aussagen an einer
**zählbaren** Größe hängen:

| Probe | Erwartung |
|---|---|
| `ls docs/plan/planning/in-progress/` ohne Slice | mit der korrigierten Formulierung **regelkonform**; vorher ein Regelverstoß im Normalbetrieb |
| Anzahl `bash tools/verify-*.sh`-Aufrufe im `verify`-Rezept | Text nennt dieselbe Zahl |

Eine Fixture-Probe gibt es nicht und wäre gelogen: der Fehler ist eine Aussage, kein Verhalten.
Der Beleg ist der Abgleich Text ↔ gezählter Bestand.

## 4. Was bewusst nicht getan wird

- **Einen Sensor für Zahlenangaben in Prosa bauen.** Verlockend nach zwei Fällen, aber ohne
  Messung spekulativ: unbekannt ist, wie viele solcher Angaben das Repo trägt. Wenn dieser Slice
  zeigt, dass es viele sind, ist das ein Folge-Slice mit eigener Zählung.
- **Das WIP-Limit inhaltlich ändern.** Die harte Größe 1 als **Obergrenze** bleibt; korrigiert
  wird nur, dass sie fälschlich auch als Untergrenze formuliert war.
- **`F-7`** (übersprungener Closure-Lauf) und **`F-11`** (fehlende Rollen-Übergaben) — beide
  Gruppe C, aber andere Mechanismen. `F-7` löst der anstehende `welle-12`-Abschluss selbst auf.

## 5. DoD

- [ ] [`AGENTS.md`](../../../../AGENTS.md) §5 formuliert das WIP-Limit als Obergrenze, und ein
      leeres `in-progress/` ist damit regelkonform. Beleg: die Formulierung, geprüft gegen den
      realen Bestand.
- [ ] Der `verify`-Kommentarblock im [`Makefile`](../../../../Makefile) nennt die Zahl, die das
      Rezept ausführt. Beleg: gezählte Aufrufe gegen die Textangabe.
- [ ] `make gates` und `make verify` grün — **Ausgabe in eine Datei**, Exit-Code getrennt geprüft,
      nie in eine Pipe.

## 6. Closure-Notiz

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice,
[`AGENTS.md`](../../../../AGENTS.md) §5; `make verify` prüft das.)_

## 7. Sub-Area-Modus

Alle berührten Sub-Areas GF.
