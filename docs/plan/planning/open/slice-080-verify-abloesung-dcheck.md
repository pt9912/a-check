# slice-080 — Die vier `verify-*` durch d-check-Module ablösen

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Deckt:** den in [slice-073](../done/slice-073-dcheck-statt-eigenbau.md) als **CR-fähig**
gemessenen Anteil — 589 der 897 Zeilen `tools/`.
**Bezug:** CR 1 und CR 2 aus [slice-073 §8](../done/slice-073-dcheck-statt-eigenbau.md);
[MR-006](../../../../harness/conventions.md#mr-006--baseline-committet-vendored-statt-per-url-referenziert).

---

## 0. Trigger

**Beginn — beobachtbar, zweiteilig.** Beide Hälften müssen erfüllt sein:

1. **Ein d-check-Release trägt Modul `structure` und die Option `links.resolve-from`.**
   Prüfbar ohne Rückfrage an der Modulliste (`d-check --help`, Handbuch §6 Regelmodule) —
   ein anderer Mensch kann sagen, ob es eingetreten ist.
2. **Der Pin in [`d-check.mk`](../../../../d-check.mk) ist auf dieses Release gehoben.**

Die zweite Hälfte ist die, die man leicht wegläßt — und ihr Fehlen hat in diesem Repo schon einmal
gekostet: das Modul `targets` stand seit `v0.38.0` bereit und lief **dreizehn Minor-Versionen**
ins Leere, weil zwar das Target eingebunden, aber nie konfiguriert wurde
([slice-074](../done/slice-074-doc-targets-wirksam.md)). Ein Modul in einem Release, das a-check
nicht zieht, ändert hier nichts.

**Vorbedingung außerhalb dieses Repos:** CR 1 und CR 2 müssen bei d-check **eingereicht** sein.
Das ist ein Akt gegenüber einem Fremdrepo und bleibt Maintainer-Sache
([slice-073 §4](../done/slice-073-dcheck-statt-eigenbau.md)). Solange die Einreichung nicht
erfolgt ist, kann Trigger-Hälfte 1 nie eintreten — **dieser Slice ist dann dauerhaft blockiert,
und das ist der ehrliche Zustand**, kein Versäumnis.

**Rückführungen:**

- `in-progress` → `open`: falls die gelieferten Module die Paritäts-Probe aus §3 nicht bestehen.
- Verwerfen: falls d-check die CRs ablehnt. Dann bleiben die 589 Zeilen dauerhaft lokal, und
  **diese Entscheidung gehört in eine ADR**, nicht in ein stilles Liegenlassen.

## 1. Auslöser

**Mechanismus: eine gemessene Ablösbarkeit ohne Werkzeug.**
[slice-073](../done/slice-073-dcheck-statt-eigenbau.md) hat je Prüfung belegt, was d-check heute
abdeckt und was nicht. Vier Sensoren fielen in die Kategorie **CR-fähig** — generisch, aber von
keinem der 19 Module gedeckt:

| Eigenbau | Zeilen | fehlende Fähigkeit |
|---|---|---|
| `verify-closure-notes` | 146 | Abschnitts-Struktur (`structure`) |
| `verify-slice-form` | 166 | dito, `max-tasks` abschnitts-treu |
| `verify-ac-form` | 131 | dito, `section-pattern` + `require-strong` |
| `verify-slice-links` | 146 | Lifecycle-feste Auflösung (`links.resolve-from`) |

Das ist a-checks eigene Doktrin auf a-check angewandt: ein Shell-Skript im Konsumenten-Repo, das
eine **generische** Invariante prüft, ist ein CR-Kandidat für das Werkzeug — *verteilen statt
kopieren*. d-check hat dieses Muster zweimal gefahren (`vcs` aus `adr-immutable-check.sh`,
`targets` aus `gate-consistency.sh`), beide Male mit Paritätsbeleg gegen das abgelöste Skript.

## 2. Betroffene Module

Zwei Schichten:

1. **[`.d-check.yml`](../../../../.d-check.yml)** — Konfiguration der neuen Module.
2. **`tools/`** — die vier `verify-*.sh` entfallen, dazu ihre Einhängung in
   [`Makefile`](../../../../Makefile) und die Zeilen in [`AGENTS.md`](../../../../AGENTS.md) §4.

## 3. Auszuführende Gates

`make verify`, `make gates`.

**Paritäts-Mutations-Beleg — die Beleg-Arbeit dieses Slice.** Die Fixture-Mengen liegen bereits
vor: jede Probe, die in
[slice-070](../done/slice-070-grundgesamtheit-messen.md),
[slice-075](../done/slice-075-sensor-messgroesse.md) und
[slice-076](../done/slice-076-vertrag-und-sensor.md) einen Eigenbau-Sensor rot gemacht hat, muss
auch das d-check-Modul rot machen — und jede, die ihn grün ließ, auch grün.

| Probe (aus den genannten Slices) | Erwartung am Modul |
|---|---|
| Slice-Datei ohne Kurztitel | rot |
| leeres `done/`, Lastenheft ohne AC-Überschrift | rot |
| Closure-Notiz „Geprueft via foo.go." | rot |
| präfixloser Referenz-Link | rot |
| legitime Leermenge (0 neue ACs), zwei echte Sätze, zustandsunabhängiger Verweis | **grün** |

Erst wenn beide Richtungen decken, darf ein Eigenbau-Sensor entfallen. Ein Modul, das mehr rot
macht als der Sensor, ist genauso ein Bruch wie eines, das weniger fängt.

## 4. Was bewusst nicht getan wird

- **Die CRs einreichen.** Fremdrepo, Maintainer-Sache — siehe §0.
- **`gate-consistency` (1)+(2).** Das ist [slice-079](../done/slice-079-gate-consistency-abloesen.md) und
  hängt an keinem CR.
- **Die lokal verbleibenden Prüfungen.** `.d-check.yml`-Modulliste, Pin-Konsistenz, `.PHONY`,
  `suppression-check`, `regelwerk-check` — in
  [slice-073](../done/slice-073-dcheck-statt-eigenbau.md) als lokal belegt, zwei davon, weil sie
  Nicht-Markdown prüfen.

## 5. DoD

- [ ] Jeder abgelöste Sensor hat einen **Paritäts-Mutations-Beleg** in beide Richtungen. Beleg:
      die Tabelle aus §3, je Zeile ein Lauf gegen Modul **und** (vor dem Entfernen) gegen den
      Sensor.
- [ ] Die abgelösten `verify-*.sh` sind entfernt, ihre Einhängung in
      [`Makefile`](../../../../Makefile) und [`AGENTS.md`](../../../../AGENTS.md) §4 nachgezogen.
      Beleg: `make gate-consistency` grün (Doku ↔ Makefile), `make verify` grün.
- [ ] `make gates` grün — **Ausgabe in eine Datei**, Exit-Code getrennt geprüft, nie in eine Pipe.

## 6. Closure-Notiz

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice,
[`AGENTS.md`](../../../../AGENTS.md) §5; `make verify` prüft das.)_

## 7. Sub-Area-Modus

Alle berührten Sub-Areas GF.
