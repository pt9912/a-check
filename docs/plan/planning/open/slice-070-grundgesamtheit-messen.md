# slice-070 — „nichts gefunden" ist nicht „nichts da"

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Deckt:** `F-12` aus dem [Review-Report `welle-12`](../../../reviews/2026-08-09-welle-12-unabhaengig.md)
und `R-068-F4` aus dem [Plan-Review](../../../reviews/2026-08-09-slice-068-plan-review.md);
trägt den in `R-068-F1` **zurückgewiesenen** Entscheid als offene Designfrage.
**Bezug:** Neuschnitt des zurückgezogenen Sammel-Entwurfs `4b029e4` nach Fehlermechanismus;
Geschwister [slice-068](../done/slice-068-phony-vollstaendig.md),
[slice-069](../in-progress/slice-069-sensor-fehler-propagierung.md),
[slice-071](../open/slice-071-sensor-scope-vollstaendig.md). Roadmap-Zeile *Aktuelle Welle* in der
[Roadmap](../in-progress/roadmap.md).

---

## 1. Auslöser

**Mechanismus: der Sensor kann „ich habe nichts gefunden" nicht von „es war nichts da" und nicht
von „ich habe es nicht erkannt" unterscheiden.** Alle drei enden heute in derselben grünen Zeile.

Im laufenden Betrieb sichtbar:

```
verify-ac-form ok: 0 neue AC(s) geprueft, 19 bei Einfuehrung bestehende (grandfathered).
```

Null geprüfte Einheiten, Ergebnis „ok" (`F-12`; betrifft alle vier `verify-*`-Sensoren).

Dazu der Spezialfall `R-068-F4` — gemessen an
[`tools/verify-slice-form.sh:32`](../../../../tools/verify-slice-form.sh):

```
slice-068.md                      -> number=''      → grandfathered
slice-068-sensor-false-greens.md  -> number='068'   → geprueft
```

`slice_num()` verlangt einen Bindestrich nach der Nummer. Eine Slice-Datei ohne Kurztitel wird
deshalb nicht etwa als Fehler gemeldet, sondern **als „älter (grandfathered)" mitgezählt** — sie
erscheint in der grünen Meldung als Teil des geprüften Bestands. `applies() == false` bedeutet
heute zweierlei: *zu alt* und *nicht erkannt*. Nur das erste ist legitim.

## 2. Betroffene Module

Eine Schicht: **`tools/`** — [`verify-closure-notes.sh`](../../../../tools/verify-closure-notes.sh),
[`verify-slice-form.sh`](../../../../tools/verify-slice-form.sh),
[`verify-slice-links.sh`](../../../../tools/verify-slice-links.sh),
[`verify-ac-form.sh`](../../../../tools/verify-ac-form.sh).

## 3. Auszuführende Gates

`make verify`; die vier Sensoren zusätzlich einzeln.

**Offene Designfrage — vor dem Bau zu entscheiden.** Der erste Anlauf (`4b029e4`) wollte prüfen,
ob „die Quelle gelesen wurde und nicht leer war". `R-068-F1` hat das zurückgewiesen, mit zwei
Gründen, die beide zutreffen: eine nichtleere Quelle kann null *erkannte* Einträge enthalten
(nichtleeres Lastenheft ohne `### AC-`-Überschriften), und eine **gemeinsame** Bedingung kann es
nicht geben, weil Leermengen je Sensor unterschiedlich legitim sind — bei `verify-ac-form` und
`verify-slice-links` regulär, bei `verify-closure-notes` nicht.

Die Anforderung lautet deshalb **nicht** „eine Bedingung", sondern:

1. **Je Sensor eine deklarierte Erwartung** an seine Grundgesamtheit, im Sensor selbst benannt und
   begründet — nicht als globale Regel von außen.
2. **Trennung von „nicht zutreffend" und „nicht erkannt".** Eine Eingabe, die der Sensor nicht
   parsen kann, ist ein **Befund**, nie eine Grandfathering-Kategorie. Das ist der allgemeine Fall
   hinter `R-068-F4`.
3. **Der Normalfall bleibt grün.** Eine Verschärfung, die legitime Leermengen rot macht, ist kein
   Fortschritt.

Welche Untergrenze je Sensor gilt, ist beim Bau festzulegen und im Sensor zu begründen — dieser
Slice schreibt sie bewusst **nicht** vor, weil genau diese Vorwegnahme im ersten Anlauf falsch war.

**Negativ-Proben:**

| Probe | Erwartung |
|---|---|
| Slice-Datei ohne Kurztitel (`slice-NNN.md`) | rot mit Namens-Befund — **nicht** grandfathered |
| Quelle vorhanden, aber ohne erkennbare Einträge | rot, je nach deklarierter Erwartung des Sensors |
| legitime Leermenge (z. B. null neue ACs) | **grün** — die Probe belegt, dass nicht pauschal verschärft wurde |

## 4. Was bewusst nicht getan wird

- **Pauschales „0 geprüft ⇒ Fehler".** Ausdrücklich verworfen: `verify-ac-form` meldet heute
  korrekt null neue ACs.
- **Die übrigen Mechanismen** — [slice-068](../done/slice-068-phony-vollstaendig.md),
  [slice-069](../in-progress/slice-069-sensor-fehler-propagierung.md),
  [slice-071](../open/slice-071-sensor-scope-vollstaendig.md).
- **Änderung der Grandfathering-*Schwellen*** (`SLICE_FORM_FROM`, AC-Stichtag). Dieser Slice
  ändert, was bei *nicht erkannter* Eingabe passiert — nicht, ab wann geprüft wird.

## 5. DoD

- [ ] Jeder der vier Sensoren deklariert seine erwartete Grundgesamtheit im Skript und meldet rot,
      wenn sie unterschritten wird. Beleg: je Sensor die Probe aus §3, vorher grün, nachher rot.
- [ ] Eine nicht parsebare Eingabe ist ein Befund, keine Grandfathering-Kategorie. Beleg: die
      `slice-NNN.md`-Probe, plus ein grüner Lauf mit legitimer Leermenge.
- [ ] `make verify` und `make gates` grün — **Ausgabe in eine Datei**, Exit-Code getrennt geprüft,
      nie in eine Pipe.

## 6. Closure-Notiz

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice,
[`AGENTS.md`](../../../../AGENTS.md) §5; `make verify` prüft das.)_

## 7. Sub-Area-Modus

Alle berührten Sub-Areas GF.
