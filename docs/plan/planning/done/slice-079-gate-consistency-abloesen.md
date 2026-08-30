# slice-079 — `gate-consistency` (1)+(2) durch das `targets`-Modul ablösen

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Deckt:** den in [slice-073](../done/slice-073-dcheck-statt-eigenbau.md) als **abgelöst**
gemessenen Anteil; hebt die in [slice-074](../done/slice-074-doc-targets-wirksam.md) bewusst
in Kauf genommene Doppelprüfung auf.
**Bezug:** Roadmap-Zeile *Aktuelle Welle* in der [Roadmap](../in-progress/roadmap.md).

---

## 0. Trigger

**Beginn:** sofort — dieser Slice wartet auf nichts. Das `targets`-Modul ist seit d-check
`v0.38.0` released, seit [slice-074](../done/slice-074-doc-targets-wirksam.md) konfiguriert, und
die Parität ist in [slice-073](../done/slice-073-dcheck-statt-eigenbau.md) mit zwei Fixtures
gemessen. Er lag bisher nur still, weil ihn niemand geschnitten hatte.

**Rückführungen** (Baseline-Vorlage §4, in a-checks Vorlage nicht abgebildet — hier bewusst
mitgeführt, weil die Startbedingung in den Slice gehört und nicht in die Roadmap):

- `in-progress` → `next`: falls sich beim Bau zeigt, dass mehr als (1)+(2) betroffen ist.
- `in-progress` → `open`: falls eine der Negativ-Proben aus §3 die gemessene Parität **widerlegt**
  — dann ist die Ablösung nicht belegt und der Slice braucht eine neue Messung, keinen Umbau.

## 1. Auslöser

**Mechanismus: zwei Prüfer für dieselbe Invariante.** Seit
[slice-074](../done/slice-074-doc-targets-wirksam.md) prüfen `make doc-targets` (d-check-Modul
`targets`) und `gate-consistency` Checks (1)+(2) beides die Deklarations-Konsistenz Doku ↔
Build-Targets. Die Doppelung ist dort **benannt**, aber sie war als Übergang gedacht, nicht als
Zustand.

[slice-073](../done/slice-073-dcheck-statt-eigenbau.md) hat die Parität gemessen — zwei Fixtures,
beide Richtungen, identische Befundmenge:

```
Phantom in AGENTS.md         -> gate-phantom       / "dokumentiert 'make phantom-target'"
undokumentiertes Make-Target -> gate-undocumented  / "reales Target 'neues-gate' fehlt"
```

Dass d-checks Modul als Ablösung **gebaut** wurde, steht in dessen `slice-063`: DoD-Punkt
*„Paritäts-Mutations-Beleg vs. `gate-consistency.sh`"*, Release `v0.38.0`.

**Dieser Slice wartet auf nichts.** Er stand bisher nur still, weil zwei `done/`-Slices im Kreis
auf „dessen Folge-Slice" verwiesen, den niemand geschnitten hatte.

## 2. Betroffene Module

Zwei Schichten:

1. **[`tools/gate-consistency.sh`](../../../../tools/gate-consistency.sh)** — `doc_targets()`,
   `check_documented_exist()`, `UTILITY_TARGETS`, der zugehörige Selbsttest-Teil und die
   Hauptteil-Blöcke (1)+(2) entfallen. Die Checks (3)(4)(5) bleiben — sie sind in
   [slice-073](../done/slice-073-dcheck-statt-eigenbau.md) als **lokal** belegt.
2. **[`Makefile`](../../../../Makefile)** und [`AGENTS.md`](../../../../AGENTS.md) §4 —
   `doc-targets` wandert ins `gates`-Aggregat, die §4-Zeile verliert ihren „nicht im
   Aggregat"-Vermerk.

## 3. Auszuführende Gates

`make gates` (enthält danach `doc-targets`), `make doc-targets`, `make gate-consistency`.

**Negativ-Proben — nach dem Umbau gegen das verbliebene Aggregat:**

| Probe | Erwartung |
|---|---|
| Phantom-Target in einer `AGENTS.md`-Tabellenzeile | `make gates` rot mit `gate-phantom` |
| undokumentiertes Rezept-Target im `Makefile` | `make gates` rot mit `gate-undocumented` |
| unveränderter Bestand | `make gates` grün, und `gate-consistency` meldet nur noch (3)(4)(5) |

Die ersten beiden sind dieselben Fixtures wie in
[slice-074](../done/slice-074-doc-targets-wirksam.md) — **jetzt aber gegen `gates`**, nicht gegen
das Einzel-Target. Ohne diese Verschiebung wäre die Ablösung eine Entfernung ohne Ersatz im
Aggregat.

## 4. Was bewusst nicht getan wird

- **Die Checks (3)(4)(5) anfassen.** `.d-check.yml`-Modulliste, Pin-Konsistenz und
  `.PHONY`-Vollständigkeit sind in [slice-073](../done/slice-073-dcheck-statt-eigenbau.md)
  als lokal belegt — zwei davon, weil sie Nicht-Markdown prüfen.
- **Auf CR 1/CR 2 warten.** Diese Ablösung hängt an keinem der beiden Anträge; das `targets`-Modul
  ist seit `v0.38.0` released und seit
  [slice-074](../done/slice-074-doc-targets-wirksam.md) konfiguriert.
- **Die vier `verify-*`-Sensoren.** Deren Ablösung hängt sehr wohl an den CRs und steht als
  Welle mit Trigger in der [Roadmap](../in-progress/roadmap.md).

## 5. DoD

- [x] `gate-consistency` prüft die Doku-↔-Target-Konsistenz nicht mehr; `doc-targets` läuft im
      `gates`-Aggregat. Beleg: beide Proben gegen `make gates`, je **EXIT=2** mit `gate-phantom`
      bzw. `gate-undocumented`.
- [x] [`AGENTS.md`](../../../../AGENTS.md) §4, [`harness/README.md`](../../../../harness/README.md)
      und die Schlussmeldung von `gate-consistency` nennen den neuen Zustand — **inklusive** der
      Scope-Korrektur an [`.d-check.yml`](../../../../.d-check.yml), ohne die die Ablösung eine
      Verkleinerung gewesen wäre (Closure-Notiz).
- [x] `make gates` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft.

## 6. Closure-Notiz

**Der Slice hätte als stille Verschlechterung enden können.** Die Paritäts-Messung aus
[slice-073](../done/slice-073-dcheck-statt-eigenbau.md) galt zwei Fixtures — Phantom-Target und
undokumentiertes Target — und beide bestätigten sich hier erneut, in beiden Richtungen. **Sie
deckten aber nicht den Scope ab:** der abgelöste Check (1) las `check_documented_exist AGENTS.md
harness/README.md`, also **zwei** Dokumente; `targets` war mit `doc-tables: [AGENTS.md]`
konfiguriert, also auf **eines**.

Gemessen, bevor der alte Code fiel:

```text
Phantom-Target NUR in harness/README.md
  vorher (Check 1)      -> gefangen
  doc-targets, alt      -> EXIT=0, 0 Treffer     <- die Regression
  doc-targets, nach Fix -> EXIT=2  harness/README.md:77  phantom-nur-hier  gate-phantom
```

Eine Ablösung vergleicht nicht nur **ob** der Ersatz greift, sondern **worüber**. Die
Paritäts-Fixtures beantworten die erste Frage; die zweite steht in der Konfiguration und wäre
beinahe mit dem alten Code verschwunden — dieselbe Klasse wie
[slice-071](../done/slice-071-sensor-scope-vollstaendig.md), nur beim Abbauen statt beim Bauen.

**Die drei Proben aus §3, gegen das Aggregat:**

```text
(1) Phantom-Target in AGENTS.md §4     -> make gates EXIT=2, AGENTS.md:135 gate-phantom
(2) undokumentiertes Rezept-Target     -> make gates EXIT=2, Makefile:157 gate-undocumented
(3) unveraenderter Bestand             -> make gates EXIT=0
```

**Beobachtbare Architektur-Aussage: der Sensor ist um seine Fremd-Aufgabe leichter geworden, nicht
um seine eigene.** `gate-consistency` schrumpft von 391 auf 333 Zeilen und behält genau das, was
[slice-073](../done/slice-073-dcheck-statt-eigenbau.md) als **lokal** belegt hat: `.d-check.yml`-
Modulliste, Pin-Konsistenz, `.PHONY`-Vollständigkeit, ADR-Index — zwei davon prüfen
Nicht-Markdown, zwei eine Gegenrichtung, die kein d-check-Modul kennt. Der Makefile-Parser bleibt
selbstgetestet (`parser_self_test`), weil `check_phony_complete` ihn weiterhin trägt; ihn mit
Check (1) zu entfernen hätte einen ungetesteten Parser hinterlassen.

**Die Schlussmeldung war eine Harness-Lüge — im Sensor gegen Harness-Lügen.** Sie sagte weiterhin
„Doku ↔ Makefile konsistent", nachdem die Prüfung entfernt war. Jetzt nennt sie, was sie prüft, und
verweist für den Rest auf `doc-targets`. Dasselbe galt für
[`harness/README.md`](../../../../harness/README.md) §Sensors, das die Übereinstimmung ausdrücklich
`gate-consistency` zuschrieb.

**Lerneintrag — Form: geschärfte Regel.** Als Prüfsatz: *Vor einer Ablösung wird nicht nur die
Befund-Parität gemessen, sondern der **Geltungsbereich** — welche Dateien, welche Verzeichnisse,
welche Ausnahmen der alte Sensor kannte und ob der neue dieselben kennt.* Die Parität ist die
leichtere Hälfte: sie zeigt sich an einer Fixture. Der Scope steht in Konfiguration und
Argumentlisten und verschwindet lautlos mit dem Code, der ihn nannte. **Zu prüfen wäre** das für
[slice-080](../done/slice-080-verify-abloesung-dcheck.md) vorab: die vier `verify-*`-Sensoren
tragen Grundgesamtheiten (Lifecycle-Verzeichnisse, Grandfathering-Schwellen, `done/`-Ausnahme), die
in keinem CR-Text stehen.

## 7. Sub-Area-Modus

Alle berührten Sub-Areas GF.
