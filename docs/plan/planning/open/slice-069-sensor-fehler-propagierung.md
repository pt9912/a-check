# slice-069 — Sensoren verschlucken keinen Fehler mehr

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Deckt:** `F-5` aus dem [Review-Report `welle-12`](../../../reviews/2026-08-09-welle-12-unabhaengig.md)
und den zweiten Teil von `R-068-F3` aus dem
[Plan-Review](../../../reviews/2026-08-09-slice-068-plan-review.md).
**Bezug:** Neuschnitt des zurückgezogenen Sammel-Entwurfs `4b029e4` nach Fehlermechanismus;
Geschwister [slice-068](../done/slice-068-phony-vollstaendig.md),
[slice-070](../open/slice-070-grundgesamtheit-messen.md),
[slice-071](../open/slice-071-sensor-scope-vollstaendig.md). Roadmap-Zeile *Aktuelle Welle* in der
[Roadmap](../in-progress/roadmap.md).

---

## 1. Auslöser

**Mechanismus: der Sensor läuft, ein Teilschritt scheitert, und der Fehler wird verworfen.** Das
Ergebnis ist nicht „unbekannt", sondern „ok" — die schädlichste Verwechslung, die ein Gate machen
kann.

Zwei Fundstellen, beide reproduziert:

**`F-5` — [`tools/commit-scope-check.sh:77`](../../../../tools/commit-scope-check.sh)** verwirft
den Fehler von `git rev-list`:

```
$ make commit-scope-check RANGE=definitely-not-a-revision
commit-scope-check ok: 0 (planning)-Commit(s) in definitely-not-a-revision geprueft, …
EXIT=0
```

Der Sensor läuft in der CI ([`ci.yml:70`](../../../../.github/workflows/ci.yml)) über eine dort
berechnete Range. Ist sie einmal unauflösbar, meldet das Gate grün, ohne einen Commit gesehen zu
haben.

**`R-068-F3` — [`tools/suppression-check.sh:33`](../../../../tools/suppression-check.sh)** beendet
`scan()` mit `|| true`, was auch Traversierungsfehler schluckt:

```
$ scan /definitely/not/a/source
bfs: error: /definitely/not/a/source: Datei oder Verzeichnis nicht gefunden.
rc=0  hits=''
```

Die Fehlermeldung erscheint auf stderr, der Exit-Code bleibt 0 — in `make gates` rutscht das
durch.

## 2. Betroffene Module

Eine Schicht: **`tools/`** —
[`commit-scope-check.sh`](../../../../tools/commit-scope-check.sh) und
[`suppression-check.sh`](../../../../tools/suppression-check.sh).

Keine Spec-, ADR- oder Produktcode-Änderung: beide Sensoren sollen den Fehler melden, den sie
heute bereits sehen.

## 3. Auszuführende Gates

`make commit-scope-check` (läuft in keinem Aggregat — gezielt aufrufen), `make gates` (enthält
`suppression-check`).

**Negativ-Proben:**

| Probe | Erwartung |
|---|---|
| `make commit-scope-check RANGE=definitely-not-a-revision` | Exit ≠ 0 **mit Range-Fehler**, nicht „0 geprüft" |
| `scan` gegen einen nicht existierenden Pfad | Exit ≠ 0, Meldung nennt den Pfad |
| gültige Range bzw. vorhandene Wurzeln | unverändert grün — die Korrektur darf den Normalfall nicht rot machen |

Die dritte Zeile ist Teil der Probe, nicht Beiwerk: eine Fehler-Propagierung, die auch im
Normalfall auslöst, wäre kein Fortschritt, sondern ein neuer Defekt.

## 4. Was bewusst nicht getan wird

- **Die übrigen False-Green-Funde** — Mechanismen „Target läuft nicht"
  ([slice-068](../done/slice-068-phony-vollstaendig.md)), „leere Menge = ok"
  ([slice-070](../open/slice-070-grundgesamtheit-messen.md)) und „nur ein Teil des Bestands gemessen"
  ([slice-071](../open/slice-071-sensor-scope-vollstaendig.md)).
- **Der Scope von `suppression-check`** — dass er nur `internal/` und `cmd/` sieht, ist `F-2` und
  gehört zu [slice-071](../open/slice-071-sensor-scope-vollstaendig.md). Dieser Slice fasst
  ausschließlich die Fehler-Propagierung an; beide Slices berühren dieselbe Datei und sind
  **nacheinander** zu fahren.
- **Ein repo-weiter Audit aller `|| true` und `2>/dev/null`** — verlockend, aber ohne Messung
  spekulativ. Wenn dieser Slice zeigt, dass das Muster häufig ist, ist das ein Folge-Slice mit
  eigener Zählung.

## 5. DoD

- [ ] Ein unauflösbarer Range lässt `commit-scope-check` mit Exit ≠ 0 und Range-Fehler enden.
      Beleg: die Probe aus §3, vorher grün, nachher rot.
- [ ] Ein Traversierungsfehler in `scan()` lässt `suppression-check` mit Exit ≠ 0 enden, während
      der reguläre Lauf grün bleibt. Beleg: beide Proben aus §3 plus der unveränderte
      `gates`-Lauf.
- [ ] `make gates` grün — **Ausgabe in eine Datei**, Exit-Code getrennt geprüft, nie in eine Pipe.

## 6. Closure-Notiz

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice,
[`AGENTS.md`](../../../../AGENTS.md) §5; `make verify` prüft das.)_

## 7. Sub-Area-Modus

Alle berührten Sub-Areas GF.
