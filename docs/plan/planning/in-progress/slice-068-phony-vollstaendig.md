# slice-068 — `.PHONY` vollständig, mit Sensor gegen Rückfall

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Deckt:** `F-1` aus dem [Review-Report `welle-12`](../../../reviews/2026-08-09-welle-12-unabhaengig.md),
verschärft durch `R-068-F2` aus dem
[Plan-Review](../../../reviews/2026-08-09-slice-068-plan-review.md).
**Bezug:** Neuschnitt des zurückgezogenen Sammel-Entwurfs `4b029e4` nach Fehlermechanismus;
Geschwister [slice-069](../open/slice-069-sensor-fehler-propagierung.md),
[slice-070](../open/slice-070-grundgesamtheit-messen.md),
[slice-071](../open/slice-071-sensor-scope-vollstaendig.md). Roadmap-Zeile *Aktuelle Welle* in der
[Roadmap](../in-progress/roadmap.md).

---

## 1. Auslöser

**Mechanismus: das Target läuft gar nicht.** [`Makefile:35`](../../../../Makefile) deklariert
`.PHONY` unvollständig — gemessen fehlen **neun** Targets:

```
verify · verify-closure-notes · verify-slice-form · verify-ac-form
verify-slice-links · suppression-check · regelwerk-check · commit-scope-check · arch-graph
```

Acht davon sind Neuzugänge der Regelwerk-Migration, `arch-graph` ist älter (`6803f88`). Existiert
eine gleichnamige Datei im Wurzelverzeichnis, meldet `make` „is up to date" und **Exit 0**, ohne
das Rezept auszuführen. Für `verify` heißt das: die Verifikationsschicht meldet Erfolg, ohne einen
Sensor gestartet zu haben.

Heute greift das nicht — keine Kollision existiert. Es ist latentes Falsch-Grün genau in der
Schicht, die Falsch-Grün verhindern soll.

**Warum ein Sensor und nicht nur eine Korrektur.** Die Liste einmal zu vervollständigen behebt den
Bestand, nicht die Ursache: das nächste neue Target fällt genauso heraus. Genau diese Lehre trägt
[slice-052](../done/slice-052-slice-form.md) („Bekanntheit ist keine Durchsetzung").

## 2. Betroffene Module

Zwei Schichten:

1. **[`Makefile`](../../../../Makefile)** — die `.PHONY`-Deklaration.
2. **[`tools/gate-consistency.sh`](../../../../tools/gate-consistency.sh)** — trägt bereits die
   Invariante „Doku ↔ Makefile konsistent" und ist damit der vorhandene Ort für „jedes Target ist
   `.PHONY`". Kein neues Target, kein neuer Eintrag in der Gate-Liste.

## 3. Auszuführende Gates

`make gate-consistency` (im `gates`-Aggregat) und `make gates`.

**Negativ-Probe — vollständig, nicht exemplarisch.** `R-068-F2` hat die ursprünglich geplante
Ein-Target-Probe verworfen: sie wäre schon rot geworden, wenn nur *ein* Target deklariert ist,
während die übrigen offen bleiben. Die Probe muss deshalb **die Differenzmenge messen**, nicht ein
Beispiel:

| Probe | Erwartung |
|---|---|
| Menge aller Rezept-Targets minus `.PHONY`-Menge | **leer** — sonst rot, mit Nennung jedes fehlenden Targets |
| Selbsttest mit künstlich aus `.PHONY` entferntem Target | rot, und nennt genau dieses Target |

Der Selbsttest ist der Beleg, dass der Sensor rot werden *kann* — ohne ihn wäre er ein toter
Sensor.

## 4. Was bewusst nicht getan wird

- **Die anderen sechs False-Green-Funde** — sie liegen in
  [slice-069](../open/slice-069-sensor-fehler-propagierung.md),
  [slice-070](../open/slice-070-grundgesamtheit-messen.md) und
  [slice-071](../open/slice-071-sensor-scope-vollstaendig.md), je nach Fehlermechanismus.
- **`.PHONY` für aus [`d-check.mk`](../../../../d-check.mk) eingebundene Targets** — das Fragment
  ist Fremdlieferung; ein Befund dort gehört gemeldet, nicht lokal gepatcht.
- **Ein generisches „alle Targets aller `include`-Dateien"** — der Sensor prüft die Targets, die
  dieses Repo selbst definiert. Die Grenze wird im Sensor benannt, nicht verschwiegen.

## 5. DoD

- [ ] `.PHONY` enthält jedes im [`Makefile`](../../../../Makefile) definierte Rezept-Target.
      Beleg: die Differenzmenge aus §3 ist leer, ausgegeben vom Sensor.
- [ ] `gate-consistency` prüft die Vollständigkeit dauerhaft und hat einen Selbsttest, der ihn
      nachweislich rot macht. Beleg: Selbsttest-Ausgabe im Lauf („Selbsttests gefeuert").
- [ ] `make gates` grün — **Ausgabe in eine Datei**, Exit-Code getrennt geprüft, nie in eine Pipe.

## 6. Closure-Notiz

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice,
[`AGENTS.md`](../../../../AGENTS.md) §5; `make verify` prüft das.)_

## 7. Sub-Area-Modus

Alle berührten Sub-Areas GF.
