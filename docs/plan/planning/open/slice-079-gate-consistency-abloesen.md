# slice-079 — `gate-consistency` (1)+(2) durch das `targets`-Modul ablösen

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Deckt:** den in [slice-073](../done/slice-073-dcheck-statt-eigenbau.md) als **abgelöst**
gemessenen Anteil; hebt die in [slice-074](../done/slice-074-doc-targets-wirksam.md) bewusst
in Kauf genommene Doppelprüfung auf.
**Bezug:** Roadmap-Zeile *Aktuelle Welle* in der [Roadmap](../in-progress/roadmap.md).

---

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

- [ ] `gate-consistency` prüft die Doku-↔-Target-Konsistenz nicht mehr; `doc-targets` läuft im
      `gates`-Aggregat. Beleg: die ersten beiden Proben aus §3 gegen `make gates`.
- [ ] Die Zeile in [`AGENTS.md`](../../../../AGENTS.md) §4 nennt den neuen Zustand, und
      `gate-consistency` beschreibt nur noch, was es prüft. Beleg: `make gates` grün, inklusive
      der Doku-↔-Makefile-Prüfung durch `targets` selbst.
- [ ] `make gates` grün — **Ausgabe in eine Datei**, Exit-Code getrennt geprüft, nie in eine Pipe.

## 6. Closure-Notiz

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice,
[`AGENTS.md`](../../../../AGENTS.md) §5; `make verify` prüft das.)_

## 7. Sub-Area-Modus

Alle berührten Sub-Areas GF.
