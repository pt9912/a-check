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

**Realer Vorfall am 2026-08-09, nicht nur Fixture.** Der Push von `6d8bbe7..2e12c58` machte die
CI rot: `commit-scope-check` meldete drei `docs(planning)`-Commits, die zusätzlich
`docs/reviews/` anfassen (CI-Run `31301467076`). Der Sensor hat **korrekt** gemeldet — das ist
der Beleg, dass er bei gültiger Range funktioniert und `F-5` wirklich nur den Fehlerpfad betrifft.

Gleichzeitig zeigt der Vorfall etwas, das keiner der beiden Reviews hatte: **der Sensor läuft zu
spät.** Er hängt weder in `make gates` noch im `commit-msg`-Hook, sondern nur in
[`ci.yml:70`](../../../../.github/workflows/ci.yml). Wer vor jedem Commit alle Gates fährt, die
`gates` kennt, erzeugt trotzdem CI-Rot. Das ist ein eigener Mangel — siehe §4.

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
- ***Wann* `commit-scope-check` läuft** (§1, realer Vorfall) — er hängt nur in der CI, nicht in
  `make gates` und nicht im `commit-msg`-Hook. Das ist ein **anderer** Mechanismus als der hier
  behandelte: nicht „der Sensor verschluckt einen Fehler", sondern „der Sensor greift erst,
  wenn der Fehler schon veröffentlicht ist". Er gehört in einen **eigenen Folge-Slice**; ihn hier
  anzuhängen hieße den Slice zu dehnen — genau das, was der Plan-Review (`R-068-F5`) am
  Vorgänger-Entwurf bemängelt hat.

## 5. DoD

- [x] Ein unauflösbarer Range lässt `commit-scope-check` mit Exit ≠ 0 und Range-Fehler enden.
      Beleg: `RANGE=definitely-not-a-revision` → Exit 2 mit
      `FAIL — Range … ist nicht aufloesbar`; `RANGE=HEAD..HEAD` (leer, aber gültig) → Exit 0.
- [x] Ein Traversierungsfehler in `scan()` lässt `suppression-check` mit Exit ≠ 0 enden, während
      der reguläre Lauf grün bleibt. Beleg: `scan` gegen eine fehlende Wurzel — vorher `rc=0`,
      nachher `rc=1`; `make suppression-check` unverändert Exit 0.
- [x] `make gates` grün — **Ausgabe in eine Datei**, Exit-Code getrennt geprüft, nie in eine Pipe.

## 6. Closure-Notiz

**Geliefert:** [`commit-scope-check.sh`](../../../../tools/commit-scope-check.sh) löst die Range
über `resolve_range()` auf und bricht fail-closed ab, wenn `git rev-list` scheitert;
[`suppression-check.sh`](../../../../tools/suppression-check.sh) prüft jede Scan-Wurzel einzeln
und hat `2>/dev/null || true` verloren. Beide Selbsttests prüfen **beide** Richtungen.

**Lerneintrag — Form: geschärfte Regel.** Als Prüfsatz: *Ein Sensor muss „nichts gefunden" von
„nicht nachgesehen" unterscheiden — und die Stelle, an der beides zusammenfällt, ist fast immer
eine Fehler-Unterdrückung, die dort steht, wo sie nicht hingehört.*

**Die Ursache** ist in beiden Fällen dieselbe Konstruktion, nicht Nachlässigkeit: eine
Command-Substitution trägt nur den Exit-Code ihres **letzten** Befehls. `$(scan ./internal; scan
./cmd)` konnte einen Fehler der ersten Wurzel strukturell nicht melden, und
`for sha in $(git rev-list …)` iteriert bei einem Fehler schlicht null mal. Beide Male sah der
Fehlerfall exakt aus wie der Erfolgsfall — kein `|| true` musste dafür falsch gesetzt sein, die
Struktur allein genügte. Deshalb ist die Korrektur nicht „das `|| true` entfernen", sondern die
Rückgabe an einer Stelle abzufragen, wo sie überhaupt ankommt.

**Gegenprobe gegen Über-Verschärfung** ist Teil beider Selbsttests: leerer Range und
treffer-freie Wurzel müssen grün bleiben. Ohne sie wäre ein Sensor, der immer rot meldet, von
einem korrekten nicht zu unterscheiden.

**Zwei beobachtbare Closure-Kriterien:**

1. `make commit-scope-check RANGE=definitely-not-a-revision` exitet ≠ 0 und nennt den Range;
   `RANGE=HEAD..HEAD` exitet 0 mit „0 Commit(s) geprueft".
2. `scan` gegen eine nicht existierende Wurzel exitet ≠ 0 (vorher `rc=0, hits=''`), während
   `make suppression-check` über den realen Baum unverändert Exit 0 meldet.

**Folge-Slices:** [slice-070](../open/slice-070-grundgesamtheit-messen.md),
[slice-071](../open/slice-071-sensor-scope-vollstaendig.md) aus Gruppe A. Dazu **neu**: *wann*
`commit-scope-check` läuft (§4) — er hängt nur in der CI, greift also erst nach dem Push. Am
2026-08-09 real aufgetreten (CI-Run `31301467076`); noch nicht geschnitten.

## 7. Sub-Area-Modus

Alle berührten Sub-Areas GF.
