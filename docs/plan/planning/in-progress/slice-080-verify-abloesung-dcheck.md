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

**Beide Hälften sind erfüllt.** `v0.67.0` trägt das Modul `structure` **und** die Option
`links.resolve-from`; der Pin ist mit [slice-115](../done/slice-115-dcheck-pin-v0670.md) gehoben.
CR 1 und CR 2 aus [slice-073 §8](../done/slice-073-dcheck-statt-eigenbau.md) sind damit umgesetzt —
die Vorbedingung außerhalb dieses Repos ist **entfallen**, nicht offen. Beleg: `--print-config` des
Release nennt in beiden Konfigurationsbeispielen a-checks eigene Verzeichnisse und
Abschnittstitel.

**Rückführungen:**

- `in-progress` → `open`: falls die gelieferten Module die Paritäts-Probe aus §3 nicht bestehen.
- Verwerfen: falls d-check CR 3 ablehnt. Dann bleiben die 589 Zeilen dauerhaft lokal, und
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

**Gemessen an `v0.67.0`.** Die vier sind **nicht** gleich weit — der Lauf gegen einen Scratch-Baum
mit echter `.d-check.yml` trennt sie:

| Eigenbau | Deckung durch das Release |
|---|---|
| `verify-slice-links` | Vertrag passt eins zu eins (`dirs` + `fixed-dirs`); Probe steht aus |
| `verify-closure-notes` | strukturelle Hälfte **ja** — `sections: one` fand den Doppeltreffer *Closure-Trigger* / *Closure-Notiz* selbst; zwei Regeln nötig wegen des Abschnitts-Namenswechsels im Bestand. Die Risiko-Ausgangs-Prüfung bleibt lokal: sie vergleicht §6 mit §7, `structure` ist abschnitts-**lokal** |
| `verify-slice-form` | **nein** — `max-tasks: 3` liefert **9** `section-oversized`, acht davon die Slices 108–115 mit je **7** Task-Items und 2–3 Liefer-Punkten. Der Bestand ist regelkonform; der Zähler misst das Falsche |
| `verify-ac-form` | **nein** — die **19** grandfatherten Anforderungen sind Abschnitte **einer** Datei, `exempt-paths` ist datei-granular |

Die zwei Lücken sind **generisch, nicht a-check-eigen**: die Ziel-Form der Baseline liefert selbst
eine DoD mit neun Checkboxen aus, sechs davon pro Slice konstant, und schreibt eine Zeile darüber,
dass Gate-Läufe und Closure-Pflichten nicht mitzählen. Daraus entsteht **CR 3** (§8). Er ist
Vorbedingung für die **vollständige** Ablösung — **nicht** für den Beginn dieses Slice.

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

- **CR 3 einreichen.** Fremdrepo, Maintainer-Sache — wie schon bei CR 1 und CR 2
  ([slice-073 §4](../done/slice-073-dcheck-statt-eigenbau.md)).
- **`gate-consistency` (1)+(2).** Das ist [slice-079](../done/slice-079-gate-consistency-abloesen.md) und
  hängt an keinem CR.
- **Die lokal verbleibenden Prüfungen.** `.d-check.yml`-Modulliste, Pin-Konsistenz, `.PHONY`,
  `suppression-check`, `regelwerk-check` — in
  [slice-073](../done/slice-073-dcheck-statt-eigenbau.md) als lokal belegt, zwei davon, weil sie
  Nicht-Markdown prüfen.

## 5. DoD

- [ ] **CR 3 ist als Text geliefert** (§8): die zwei gemessenen Lücken in `structure`, je mit
      Paritäts-Mutations-Beleg, und ohne Verhaltensänderung, solange die neuen Schlüssel fehlen.
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

## 8. CR-Text für d-check

Dieser Abschnitt **ist** die Lieferung aus §5 DoD 1. Er liegt im Slice, weil §4 das Einreichen
ausdrücklich dem Maintainer überlässt — dieselbe Form wie
[slice-073 §8](../done/slice-073-dcheck-statt-eigenbau.md).

---

### CR 3 — `structure`: die geprüfte Menge deklarierbar machen

**Anlass — gemessen, nicht vermutet.** `structure` wendet `max-tasks` und die Abschnitts-Auswahl
auf die **vorgefundene** Menge an. Das Regelwerk verlangt an zwei Stellen eine **erklärte
Teilmenge**:

1. **Die Größen-Regel zählt nicht alle Task-Items.** Ein Lauf von `max-tasks: 3` gegen die
   Slice-Pläne eines Adopters (`v0.67.0`, `--enable structure`) liefert **9 Befunde**
   `section-oversized`. Acht davon sind die Slices 108–115 — **jeder Slice, der seit Inkrafttreten
   der Konstanten-Regel geschrieben wurde**: je **7** Task-Items, davon **2–3** Liefer-Punkte. Der
   Bestand ist regelkonform; der Zähler misst das Falsche.
2. **Grandfathering lebt innerhalb einer Datei.** Wer `require-all: [Happy, Boundary, Negative]`
   über `section-pattern` durchsetzt, muss die bei Einführung bestehenden Anforderungen ausnehmen —
   hier **19**. Sie sind Abschnitte **einer** Datei; `exempt-paths` ist datei-granular und kann sie
   nicht erreichen. Die Alternative ist, den Bestand umzuschreiben — das trifft die Form statt der
   Substanz und macht aus vertraglich bindenden Anforderungen Formularübungen.

**Nicht adopter-spezifisch — die Ziel-Form erzeugt den Fehlbefund selbst.** Die Slice-Vorlage der
Baseline liefert eine DoD mit **neun** Checkboxen aus, von denen **sechs** pro Slice konstant sind
(Gate-Lauf, Closure-Notiz, Reconciliation-Register, Beobachtungs-Register, Risiko-Ausgang, drei
Paarungen), und schreibt eine Zeile darüber selbst: *„Gate-Läufe und die vier Closure-Pflichten
darunter zählen nicht mit."* Wer die Vorlage benutzt **und** die Größen-Regel prüft, bekommt heute
zwangsläufig Falsch-Positive — auf **jedem** neuen Slice, während der Altbestand grün bleibt. Der
Sensor wird über die Zeit unbrauchbar, nicht sofort.

**Vertrag.** Zwei Optionen an bestehenden Regeln; ohne sie byte-identisches Verhalten. Keine neuen
Grund-Codes — beide **verkleinern** nur die geprüfte Menge.

```yaml
structure:
  - files: "docs/plan/planning/**/slice-*.md"
    section-pattern: '^## .*(DoD|Definition of Done)'
    max-tasks: 3
    tasks-ignore-pattern: '(make gates|Closure-Notiz|Beobachtungs-Register|Risiko aus)'
    #   Task-Items, die dieses RE2 treffen, zaehlen fuer max-tasks NICHT mit.
    #   Ohne den Schluessel: alle Items zaehlen — heutiges Verhalten.

  - files: "spec/lastenheft.md"
    section-pattern: '^### AC-'
    sections: each
    require-all: [Happy, Boundary, Negative]
    exempt-sections: '^AC-[A-Z]+-0(0[1-9]|1[0-9])\b'
    #   Abschnitte, deren UEBERSCHRIFTSTEXT dieses RE2 trifft, prueft DIESE Regel
    #   nicht — Geschwister von exempt-paths, eine Granularitaetsstufe tiefer.
    #   Der Stichtag steht damit in der Konfiguration statt in einem Skript.
```

**Warum keine eigenen Module.** Beides ist dieselbe Frage mit erklärter Grundmenge, keine neue
Frage. `tasks-ignore-pattern` gehört zu `max-tasks` wie `order-column` zu `order`;
`exempt-sections` ist das Geschwister von `exempt-paths` eine Stufe tiefer. Als Optionen bleibt
der Default unberührt.

**Fence-Treue gilt weiter.** Beide Muster dürfen Code-Blöcke und Inline-Code nicht sehen — sonst
kippt genau die Eigenschaft, die CR 1 gegenüber der Skript-Variante ausgezeichnet hat. Ein
Adopter, der über sein eigenes Regelwerk schreibt, zitiert seine Konstanten-Begriffe ständig in
Backticks.

**Paritäts-Mutations-Beleg.** Die Fixtures liegen vor:

| Probe | Erwartung |
|---|---|
| Slice mit 3 Liefer-Punkten + 4 Konstanten, `max-tasks: 3` **mit** `tasks-ignore-pattern` | **grün** (heute rot — die acht gemessenen Fälle) |
| derselbe Slice mit **vier** Liefer-Punkten | rot |
| `tasks-ignore-pattern` abwesend | rot — heutiges Verhalten unverändert |
| grandfatherte Anforderung ohne `Boundary`-Marke, von `exempt-sections` getroffen | grün |
| neue Anforderung ohne `Boundary`-Marke, nicht getroffen | rot |

**Abgrenzung.** Nicht Teil des Antrags: Grandfathering **ab einer Nummer** als Werkzeug-Begriff —
das ist über Globs bzw. RE2 ausdrückbar. Ebenfalls nicht: abschnitts-**übergreifende** Bedingungen
(„jedes Risiko aus §6 trägt einen Ausgang in §7"). Das ist eine andere Frage und braucht einen
eigenen Antrag, falls er sich lohnt.
