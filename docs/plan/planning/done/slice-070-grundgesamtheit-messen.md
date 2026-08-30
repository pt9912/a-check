# slice-070 — „nichts gefunden" ist nicht „nichts da"

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Deckt:** `F-12` aus dem [Review-Report `welle-12`](../../../reviews/2026-08-09-welle-12-unabhaengig.md)
und `R-068-F4` aus dem [Plan-Review](../../../reviews/2026-08-09-slice-068-plan-review.md);
trägt den in `R-068-F1` **zurückgewiesenen** Entscheid als offene Designfrage.
**Bezug:** Neuschnitt des zurückgezogenen Sammel-Entwurfs `4b029e4` nach Fehlermechanismus;
Geschwister [slice-068](../done/slice-068-phony-vollstaendig.md),
[slice-069](../done/slice-069-sensor-fehler-propagierung.md),
[slice-071](../done/slice-071-sensor-scope-vollstaendig.md). Roadmap-Zeile *Aktuelle Welle* in der
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
`tools/verify-slice-form.sh:32`:

```
slice-068.md                      -> number=''      → grandfathered
slice-068-sensor-false-greens.md  -> number='068'   → geprueft
```

`slice_num()` verlangt einen Bindestrich nach der Nummer. Eine Slice-Datei ohne Kurztitel wird
deshalb nicht etwa als Fehler gemeldet, sondern **als „älter (grandfathered)" mitgezählt** — sie
erscheint in der grünen Meldung als Teil des geprüften Bestands. `applies() == false` bedeutet
heute zweierlei: *zu alt* und *nicht erkannt*. Nur das erste ist legitim.

## 2. Betroffene Module

Eine Schicht: **`tools/`** — `verify-closure-notes.sh`,
`verify-slice-form.sh`,
`verify-slice-links.sh`,
`verify-ac-form.sh`.

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
  [slice-069](../done/slice-069-sensor-fehler-propagierung.md),
  [slice-071](../done/slice-071-sensor-scope-vollstaendig.md).
- **Änderung der Grandfathering-*Schwellen*** (`SLICE_FORM_FROM`, AC-Stichtag). Dieser Slice
  ändert, was bei *nicht erkannter* Eingabe passiert — nicht, ab wann geprüft wird.

## 5. DoD

- [x] Jeder der vier Sensoren deklariert seine erwartete Grundgesamtheit im Skript und meldet rot,
      wenn sie unterschritten wird. Beleg: vier isolierte Proben — leeres `done/` → Exit 1,
      Lastenheft ohne `### AC-`-Überschrift → Exit 1, fehlendes Lifecycle-Verzeichnis → Exit 1,
      nicht parsebarer Slice-Name → Exit 2.
- [x] Eine nicht parsebare Eingabe ist ein Befund, keine Grandfathering-Kategorie. Beleg:
      `slice-999.md` im echten Baum → `Dateiname gibt keine Slice-Nummer her`, Exit 2; nach
      Entfernen wieder Exit 0. Legitime Leermenge bleibt grün: `0 neue AC(s) geprueft`.
- [x] `make verify` und `make gates` grün — **Ausgabe in eine Datei**, Exit-Code getrennt geprüft,
      nie in eine Pipe.

## 6. Closure-Notiz

**Geliefert:** Die vier `verify-*`-Sensoren deklarieren ihre erwartete Grundgesamtheit im Skript
und melden rot, wenn sie unterschritten wird — **je Sensor eine eigene Grenze**, keine gemeinsame
Bedingung. `verify-slice-form` unterscheidet zusätzlich „nicht erkannt" von „zu alt": `applies()`
hat einen dritten Ausgang, ein nicht parsebarer Dateiname ist ein Befund.

**Lerneintrag — Form: geschärfte Regel.** Als Prüfsatz: *Die Untergrenze eines Sensors gehört auf
die Größe, die nicht null sein darf — und welche das ist, entscheidet der Sensor, nicht der
Harness.*

**Die Ursache**, warum der erste Anlauf (`4b029e4`) daran scheiterte: Ich habe eine *gemeinsame*
Bedingung gesucht („Quelle gelesen und nicht leer"), weil vier gleich aussehende Sensoren eine
gleich aussehende Regel nahelegen. Sie sind aber nicht gleich — bei `verify-ac-form` ist null
geprüfte Einheiten der **reguläre** Zustand dieses Repos, bei `verify-closure-notes` wäre er
Bestandsverlust. Eine Bedingung über alle vier hätte entweder den Normalfall rot gemacht oder das
Loch offen gelassen. Der Plan-Review hat genau das benannt (`R-068-F1`), bevor eine Zeile Code
entstand.

Konkret liegen die Grenzen deshalb auf drei verschiedenen Größen: auf der **Dateizahl**
(`done/` > 0), auf der **Summe aus geprüft und grandfathered** (Slice-Form, AC-Form — nicht auf
der gefilterten Menge), und auf der **Existenz der Quelle** (Lifecycle-Verzeichnisse — dort ist
die Leermenge legitim).

**Zwei beobachtbare Closure-Kriterien:**

1. Vier isolierte Fixtures machen je einen Sensor rot (leeres `done/`, Lastenheft ohne ACs,
   fehlendes `next/`, `slice-999.md`), während `make verify` über den realen Baum Exit 0 meldet —
   einschließlich `verify-ac-form ok: 0 neue AC(s) geprueft` als belegter legitimer Leerfall.
2. `verify-slice-links` gibt bei fehlendem Verzeichnis **nicht** mehr die `SL-002`-Meldung aus,
   sondern eine, die zum Befund passt.

**Folge-Slices:** [slice-071](../done/slice-071-sensor-scope-vollstaendig.md),
[slice-072](../done/slice-072-scope-sensor-praeventiv.md).

## 7. Sub-Area-Modus

Alle berührten Sub-Areas GF.
