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

**Gemessen an `v0.68.0`** (die Erstmessung lief gegen `v0.67.0`, das CR 3 noch nicht trug — der
Unterschied steht in der letzten Spalte). Die vier sind **nicht** gleich weit:

| Eigenbau | Deckung durch `v0.67.0` | Deckung durch `v0.68.0` |
|---|---|---|
| `verify-slice-links` | Vertrag passt eins zu eins (`dirs` + `fixed-dirs`); Probe stand aus | **abgelöst** — Probe gefahren, Parität in beide Richtungen |
| `verify-closure-notes` | strukturelle Hälfte **ja** — `sections: one` fand den Doppeltreffer *Closure-Trigger* / *Closure-Notiz* selbst | **strukturelle Hälfte abgelöst**; die Risiko-Ausgangs-Prüfung bleibt lokal: sie vergleicht §6 mit §7, `structure` ist abschnitts-**lokal** |
| `verify-slice-form` | **nein** — `max-tasks: 3` lieferte **9** `section-oversized`, acht davon Slices mit 2–3 Liefer-Punkten. Der Bestand war regelkonform; der Zähler maß das Falsche | **abgelöst** — `tasks-ignore-pattern` erklärt die Grundmenge, 229 Dateien / 0 Befunde |
| `verify-ac-form` | **nein** — die **19** grandfatherten Anforderungen sind Abschnitte **einer** Datei, `exempt-paths` ist datei-granular | **weiterhin nein** — `exempt-section-pattern` erreicht sie, aber ein zwanzigstes AC gibt es nicht: die Menge liefe leer, und die Nullmengen-Härte meldet `section-missing`, wo der Sensor grün lässt |

Die zwei Lücken waren **generisch, nicht a-check-eigen**: die Ziel-Form der Baseline liefert selbst
eine DoD mit neun Checkboxen aus, sechs davon pro Slice konstant, und schreibt eine Zeile darüber,
dass Gate-Läufe und Closure-Pflichten nicht mitzählen. Daraus entstand **CR 3** (§8) — angenommen
und in `v0.68.0` umgesetzt.

## 2. Betroffene Module

Drei Schichten:

1. **[`.d-check.yml`](../../../../.d-check.yml)** — Konfiguration der neuen Module.
2. **`tools/`** — zwei `verify-*.sh` entfallen ganz, ein drittes schrumpft auf seinen
   abschnitts-übergreifenden Kern, `verify-ac-form.sh` bleibt (§6). Dazu die Einhängung in
   [`Makefile`](../../../../Makefile), die Zeilen in [`AGENTS.md`](../../../../AGENTS.md) §4 und
   [`harness/README.md`](../../../../harness/README.md) §Sensors sowie die GATES-Liste des Guard.
3. **`d-check.mk`** — der Pin auf `v0.68.0`, ohne den die zwei neuen Schlüssel ins Leere liefen.

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

- [x] **CR 3 ist als Text geliefert** (§8): die zwei gemessenen Lücken in `structure`, je mit
      Paritäts-Mutations-Beleg, und ohne Verhaltensänderung, solange die neuen Schlüssel fehlen.
- [x] Jeder abgelöste Sensor hat einen **Paritäts-Mutations-Beleg** in beide Richtungen. Beleg:
      die Tabelle aus §3, je Zeile ein Lauf gegen Modul **und** (vor dem Entfernen) gegen den
      Sensor.
- [x] Die abgelösten `verify-*.sh` sind entfernt, ihre Einhängung in
      [`Makefile`](../../../../Makefile) und [`AGENTS.md`](../../../../AGENTS.md) §4 nachgezogen.
      Beleg: `make gate-consistency` grün (Doku ↔ Makefile), `make verify` grün.
- [x] `make gates` grün — **Ausgabe in eine Datei**, Exit-Code getrennt geprüft, nie in eine Pipe.

## 6. Closure-Notiz

**Geliefert:** Zwei Eigenbau-Sensoren sind entfallen (`verify-slice-form.sh`,
`verify-slice-links.sh`), ein dritter ist auf seinen abschnitts-übergreifenden Kern reduziert und
heißt jetzt nach dem, was er tut (`verify-risiko-ausgaenge.sh`, 323 → 187 Zeilen). Zusammen **509
Zeilen Shell weniger, keine Zusage weniger** — jede Befundklasse ist vor dem Entfernen gegen
Sensor **und** Modul gefahren worden, in beide Richtungen. Der Pin steht auf `v0.68.0`, das CR 3
aus §8 umsetzt.

**Lerneintrag — Form: geschärfte Regel.** *Ein `hint` gehört nur an eine Regel, deren modul-eigene
Meldung generisch ist — er gewinnt gegen sie, und wo sie Zahlen oder Namen trägt, kostet er
Diagnose.* Gemessen: die Kopffeld-Regel meldete achtzehnmal `section-marker-missing`, und mein
`hint` sagte dazu „Kopf trägt Verantwortlich/Autor/Spec-Stellen" — welche der drei Marken fehlte,
stand nirgends. Erst nach dem Entfernen des `hint` nannte die modul-eigene Meldung sie:
`geforderte Marke fehlt: Spec-Stellen:` — und damit auch den Fehler, nämlich dass der Bestand
`**Berührte Spec-Stellen:**` schreibt und `require-all` keine freie Substring-Suche ist. Dasselbe
gilt für die Größen-Regel, deren Meldung `trägt 4 Task-Items (3 ignoriert), erlaubt sind 3` drei
Zahlen führt, die kein verfasster Satz ersetzt. *Weil* der `hint` per Vertrag gewinnt, ist er kein
additives Feld: ihn zu setzen heißt, die Meldung des Moduls zu **verwerfen**. Er bleibt darum nur
dort stehen, wo die Regel eine Zusage hütet, die aus dem Grund-Code allein nicht hervorgeht.

**Drei beobachtbare Closure-Kriterien:**

1. Der Paritäts-Mutations-Beleg deckt **beide** Richtungen, je Befundklasse ein Lauf gegen Sensor
   und Modul: vierter Liefer-Punkt, fehlende Lerneintrag-Form, fehlendes Kopffeld, zwei
   Closure-Abschnitte, Rumpf aus einer Floskel, präfixloser Verweis — sechsmal `Sensor=2 Modul=2`,
   und über den unveränderten Bestand `Sensor=0 Modul=0`.
2. `make verify` läuft grün mit `doc-structure` im Aggregat; `make doc-check` grün mit
   `links.resolve-from` — 229 Dateien, 0 Befunde.
3. Die Lifecycle-Invariante hat die **Schicht gewechselt**: sie hing an `verify`, sie hängt jetzt
   an `gates`. Beobachtbar daran, dass `make doc-check` sie fährt und das Workflow-Skelett in
   Schritt 9 darauf zeigt statt auf Schritt 8.

**Was nicht abgelöst werden konnte, und warum das kein Rest ist.** `verify-ac-form` bleibt
vollständig lokal. Das Muster `exempt-section-pattern` erreicht die 19 grandfatherten
Anforderungen sauber — genau das war CR 3 —, aber es gibt kein zwanzigstes AC. Die Regel liefe
leer, und d-check meldet dann per Vertrag `section-missing`, während der Sensor die legitime
Leermenge grün lässt. Das ist der Bruch, den §3 als Abbruchbedingung nennt: *ein Modul, das mehr
rot macht als der Sensor, ist genauso ein Bruch wie eines, das weniger fängt.* Die Nullmengen-Härte
ist dabei richtig — sie verhindert das stille Abschalten einer Regel. Sie kollidiert nur mit einem
Bestand, in dem die geprüfte Menge legitim leer sein darf. Ebenso bleibt die Risiko-Ausgangs-Prüfung
lokal: sie misst §6 an §7, und `structure` prüft abschnitts-**lokal**.

**Offene Risiken und ihr Ausgang:**

- *Regel (3) verlangt die **nummerierte** Closure-Überschrift, weil d-check zwei Regeln gleicher
  Identität abweist; ein künftiger Slice mit unnummeriertem Abschnitt entginge ihr* —
  **Ausgang:** weiter offen im **Beobachtungs-Register**, die Grenze steht im Regel-Kommentar.
- *Die Identitäts-Regel zwingt zu unterschiedlichen Mustern, wo zwei Bedingungen dieselbe
  Abschnitts-Menge treffen, aber verschiedenes Grandfathering brauchen* — **Ausgang:** gestrichen
  mit Begründung: hier war der Unterschied sachlich belegbar (63 von 63 halten die nummerierte
  Form ein), und ein CR wäre erst fällig, wenn ein Fall auftritt, in dem er es nicht ist.
- *Verweise **auf** wandernde Slices bleiben ungeprüft* — **Ausgang:** Folge-Slice; der Zähler von
  `BEO-008` steht mit diesem Slice auf **3×** und hat damit die Schwelle überschritten.

**Beobachtungs-Register:** `BEO-008` auf **3×** erhöht (Beleg slice-080) und als Harness-Lücke
markiert — dreizehn Verweise aus `done/` zeigten nach dem `git mv` ins Leere, gefangen hat es
wieder `doc-check` **nach** dem Wechsel. `BEO-005` und `BEO-006` auf die neue Verortung nachgezogen
(die Platzhalter-Liste lebt als `forbid-pattern`). Bei `BEO-014` ist die zweite Hälfte
(`doc-structure` ohne Konfigurationsblock) **entfallen** — der Block existiert; für `planning`
steht die Beobachtung unverändert.

**Folge-Slices:** ein Sensor für `BEO-008` (Verweise auf wandernde Slices); `verify-ac-form` wird
ablösbar, sobald ein zwanzigstes AC entsteht — dann ist die Menge nicht mehr leer und die
Nullmengen-Härte kollidiert nicht mehr.
## 7. Sub-Area-Modus

Alle berührten Sub-Areas GF.

## 8. CR-Text für d-check

Dieser Abschnitt **ist** die Lieferung aus §5 DoD 1. Er liegt im Slice, weil §4 das Einreichen
ausdrücklich dem Maintainer überlässt — dieselbe Form wie
[slice-073 §8](../done/slice-073-dcheck-statt-eigenbau.md).

---

### CR 3 — `structure`: die geprüfte Menge deklarierbar machen

**Ergebnis: angenommen und umgesetzt in `v0.68.0`** (d-check slice-179,
[dessen ADR-0075](https://github.com/pt9912/d-check/blob/main/docs/plan/adr/0075-erklaerte-teilmenge-in-structure.md)
nennt „ein eingehender CR eines Adopters" als Anlass). Zwei Abweichungen vom Antrag, beide
begründet und beide hier nachgemessen:

1. **Der Schlüssel heißt `exempt-section-pattern`, nicht `exempt-sections`.** In `structure`
   bedeutet das Suffix `-pattern` RE2, und `exclude-sections` ist in zwei anderen Modulen bereits
   als Liste *literaler* Überschriften vergeben.
2. **Das Abschnitts-Muster sieht die ROHE Überschriften-Zeile** samt `#`-Folge — wie
   `section-pattern` daneben. Der Antrag hatte `'^AC-…'` vorgeschlagen; das trifft am realen
   Lastenheft **nichts**, weil dort `### AC-…` steht. Die umgesetzte Fassung nimmt genau die
   Falle heraus, in die der Antrag selbst gelaufen war.

Der Antragstext bleibt darunter unverändert stehen — er ist die Lieferung aus §5 DoD 1 und der
Beleg dafür, was beantragt wurde.

---


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
