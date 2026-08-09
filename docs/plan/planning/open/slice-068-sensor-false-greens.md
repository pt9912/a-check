# slice-068 — Fünf False-Green-Sensoren (Gruppe A des unabhängigen Reviews)

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Deckt:** `F-1`, `F-2`, `F-5`, `F-12`, `F-14` aus dem
[Review-Report 2026-08-09](../../../reviews/2026-08-09-welle-12-unabhaengig.md);
[ADR-0005](../../adr/0005-lint-profil.md) (Scope der Suppression-Regel),
[MR-006](../../../../harness/conventions.md#mr-006--baseline-committet-vendored-statt-per-url-referenziert)
(Baseline-Integrität).
**Bezug:** erster unabhängiger Review von `welle-12`; Roadmap-Zeile *Aktuelle Welle* in der
[Roadmap](../in-progress/roadmap.md).

---

## 1. Auslöser

Der erste Review außerhalb der Claude-Modellfamilie (2026-08-09) hat fünf Sensoren gefunden, die
**grün melden, ohne gemessen zu haben**. Jeder Fund trägt eine Fixture; zwei davon sind in diesem
Repo nachgestellt worden:

- `make commit-scope-check RANGE=definitely-not-a-revision` → **Exit 0**, Ausgabe
  „ok: 0 (planning)-Commit(s) in definitely-not-a-revision geprueft" (`F-5`).
- `make verify` meldet im grünen Lauf „verify-ac-form ok: **0 neue AC(s) geprueft**" — null
  geprüfte Einheiten, Ergebnis „ok" (`F-12`).

Das trifft die zentrale Behauptung der Regelwerk-Migration: dass Regeln nicht nur aufgeschrieben,
sondern gemessen werden. Ein Sensor, der bei leerer oder unauflösbarer Grundgesamtheit „ok" sagt,
ist von einem bestandenen Lauf nicht unterscheidbar.

Die fünf Funde im Einzelnen:

| Fund | Sensor | Defekt |
|---|---|---|
| `F-1` | alle neuen Targets | `.PHONY` fehlt → eine gleichnamige Datei lässt `make` das Rezept überspringen, Exit 0 |
| `F-2` | `suppression-check` | scannt nur `internal/` und `cmd/`, behauptet „Go-Quellen des Repos" |
| `F-5` | `commit-scope-check` | Fehler von `git rev-list` verworfen → unauflösbarer Range meldet „0 geprüft", Exit 0 |
| `F-12` | die vier `verify-*` | leere Grundgesamtheit wird als „ok" gemeldet |
| `F-14` | `regelwerk-check` | `sha256sum -c` prüft nur gelistete Einträge; eine nicht manifestierte Datei im Baseline-Baum bleibt ungemessen |

## 2. Betroffene Module

Zwei Schichten:

1. **`Makefile`** — `.PHONY`-Deklaration (`F-1`).
2. **`tools/`** — `suppression-check.sh` (`F-2`), `commit-scope-check.sh` (`F-5`),
   `verify-closure-notes.sh` / `verify-slice-form.sh` / `verify-slice-links.sh` /
   `verify-ac-form.sh` (`F-12`), `regelwerk-check.sh` (`F-14`).

Keine Spec-, ADR- oder Produktcode-Änderung: die Sensoren sollen tun, was die vorhandene Doku
bereits behauptet. Damit ist dies **kein** spec-first-Fall.

## 3. Auszuführende Gates

`make gates` (enthält `suppression-check`, `gate-consistency`, `guard-selftest`, `record-gates`)
und `make verify`; dazu gezielt `make commit-scope-check` und `make regelwerk-check`, die in
keinem Aggregat laufen.

**Negativ-Proben — der eigentliche Beleg dieses Slice.** Jeder korrigierte Sensor bekommt eine
Probe in seinen Selbsttest, die auf dem Stand **vor** der Korrektur grün gewesen wäre und danach
rot ist:

| Sensor | Negativ-Probe |
|---|---|
| `Makefile` | Datei mit dem Namen eines Prüf-Targets im Repo-Wurzelverzeichnis; das Target läuft trotzdem |
| `suppression-check` | `.go`-Datei mit `//nolint` außerhalb von `internal/`/`cmd/` |
| `commit-scope-check` | nicht auflösbarer Range → Exit ≠ 0 mit Range-Fehler statt „0 geprüft" |
| die vier `verify-*` | Grundgesamtheit künstlich leer (Verzeichnis bzw. Quelldatei ohne Einträge) → Exit ≠ 0 |
| `regelwerk-check` | Datei im Baseline-Baum, die in `SHA256SUMS` fehlt |

Ein Sensor ohne Probe, die ihn nachweislich rot macht, zählt in diesem Slice als nicht behoben.

**Zur Bau-Reihenfolge:** die vier `verify-*`-Sensoren lesen Markdown. Bei Änderungen an ihrer
Muster-Logik sind Zitat-Kontexte (Inline-Code, Code-Blöcke) auszublenden —
[`SL-004`](../../steering-loop.md). Dieser Slice fasst die Muster-Logik nicht an, nur die
Bestands-Prüfung; die Regel gilt trotzdem für die neuen Fixtures.

## 4. Was bewusst nicht getan wird

- **`F-3`, `F-4`, `F-6`, `F-13` (Sensoren messen das Falsche)** — Gruppe B. Sie brauchen je einen
  Entscheid, *was* gemessen werden soll; das ist keine Bestands-Korrektur und gehört nicht in
  denselben Schnitt.
- **`F-7`, `F-8`, `F-10`, `F-11`, `F-15` (Status- und Prozess-Widersprüche)** — Gruppe C.
- **`F-9` (Freigabe-Belege)** — Gruppe D, zusammen mit dem `--print-mk`-Digest-Defekt.
- **Semantik-Verschärfung bei `F-12`:** „0 geprüfte Einheiten" wird **nicht** pauschal zum Fehler.
  Bei `verify-ac-form` ist null neue ACs der reguläre Zustand. Geprüft wird, ob die **Quelle**
  gelesen wurde und nicht leer war — nicht, ob die gefilterte Menge nicht leer ist.
- **Ein Go-Parser für `suppression-check`** — die Zeichenketten-Grenze bleibt bestehen und ist als
  ehrliche Grenze deklariert.

## 5. DoD

- [ ] Alle fünf Sensor-Defekte behoben, **je mit einer Negativ-Probe im Selbsttest des jeweiligen
      Sensors**, die vor der Korrektur grün und danach rot ist. Beleg: Selbsttest-Ausgabe je
      Sensor, plus die reproduzierten Fixtures aus §3.
- [ ] `.PHONY` im [`Makefile`](../../../../Makefile) vollständig — jedes Target ist deklariert.
      Beleg: Fixture-Datei mit einem Target-Namen im Wurzelverzeichnis; das Target läuft trotzdem.
- [ ] `make gates` und `make verify` grün — **Ausgabe in eine Datei**, Exit-Code getrennt geprüft,
      nie in eine Pipe.

## 6. Closure-Notiz

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice,
[`AGENTS.md`](../../../../AGENTS.md) §5; `make verify` prüft das.)_

## 7. Sub-Area-Modus

Alle berührten Sub-Areas GF.
