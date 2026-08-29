# slice-109 — `planning/README.md` auf die Ziel-Form

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.
**Autor:** Claude (Opus 5), im Auftrag des Maintainers. **Datum:** 2026-08-29.
**Berührte Spec-Stellen:** — *(keine)*
**Deckt:** den `planning/README.md`-Anteil aus
[slice-105](../done/slice-105-form-review-nachholen.md) §4.
**Bezug:** [slice-108](../done/slice-108-agents-hard-rule-37.md) (Hard Rule 3.7 gilt seitdem).
[Roadmap](../in-progress/roadmap.md).

---

## 1. Ziel

Die Datei trägt vier Sektionen der Ziel-Form nicht (*Lifecycle-Bedeutungen · Slices vs. Wellen ·
Aktueller Stand · Roadmap*), keinen der vier Regelwerk-Zeiger — **und** zwei Stellen, die seit
slice-108 gegen Hard Rule 3.7 verstoßen.

Die *Aktueller Stand*-Sektion ist die interessanteste: sie existiert in der Ziel-Form, um eine
Tabelle zu **verbieten**. *„Nicht als Snapshot hier eintragen — der Stand ergibt sich aus den
Verzeichnissen, sonst driftet die Tabelle."* Eine fehlende Sektion ist hier also eine fehlende
**Warnung**, nicht bloß eine fehlende Überschrift.

## 2. Definition of Done

- [x] Die vier Sektionen stehen, je mit Regelwerk-Zeiger (0/4 → 4/4) — Beleg: Diff und Zählung.
- [x] Die zwei Chronik-Stellen sind entfernt und die Größen-Regel nennt **Liefer-Punkte** statt
      DoD-Punkte — Beleg: Diff.
- [x] `BEO-011` verliert nichts; die Zeile betrifft `README.md` und die Spec-Straten — Beleg:
      Register unverändert.

- [x] `make gates` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.
- [x] Closure-Notiz mit benanntem Lerneintrag geschrieben (§7).
- [x] Beobachtungs-Register fortgeschrieben.
- [x] Jedes Risiko aus §6 trägt genau einen Ausgang.

## 3. Plan (vor Code)

| Datei | Änderungs-Art | Begründung |
|---|---|---|
| `docs/plan/planning/README.md` | update | vier Sektionen, vier Zeiger, zwei Chronik-Stellen, Größen-Regel |

**Auszuführende Gates:** `make gates` (tragend `doc-check`), zum Abschluss `make verify`. **Kein
neuer Sensor.**

## 4. Trigger

**Start:** unmittelbar; die Lücke ist gemessen.

**Rückführungen:** `in-progress` → `open`, falls die Wellen-Closure-Prozedur beim Einpassen der
Ziel-Form-Sektionen inhaltlich berührt wird — sie ist a-checks eigene Ausarbeitung und gehört
nicht nebenbei umgebaut.

## 5. Closure-Trigger

Vier Zeiger gezählt, vier Sektionen vorhanden, Chronik weg, Gates grün.

**Was bewusst nicht getan wird:** Die **Wellen-Closure-Prozedur** bleibt, wie sie ist. Sie ist
a-checks Ausarbeitung von `modul-06` §Wellen-Closure und trägt Belege, die die Ziel-Form nicht
kennt; sie zu kürzen wäre ein eigener Schnitt mit eigener Begründung. Auch die Zahl *21 Verweise*
bleibt: sie ist **Begründung mit Praxis-Bezug**, nicht Chronik — sie sagt, *warum* die Regel gilt,
nicht, was mit dem Dokument geschah.

## 6. Risiken und offene Punkte

- *Die Grenze zwischen Begründung und Chronik ist ein Urteil* — genau die inferentielle Hälfte von
  Hard Rule 3.7. **Ausgang:** weiter offen, `BEO-016` im Beobachtungs-Register deckt das ab.
- *Vier neue Sektionen können mit `AGENTS.md` §5 doppeln* (dort steht die Lifecycle-Tabelle der
  Übergänge). **Ausgang:** entfallen — gestrichen mit Begründung: die Ziel-Form trennt
  *Bedeutungen* (was ein Verzeichnis heißt) von *Übergängen* (wann gewechselt wird); das eine
  steht hier, das andere dort, und beide verweisen aufeinander.

## 7. Closure-Notiz

**Geliefert:** vier Sektionen der Ziel-Form mit 4/4 Regelwerk-Zeigern, zwei Chronik-Stellen
entfernt, die Größen-Regel auf Liefer-Punkte nachgezogen — und eine dritte Restspur der alten
Roadmap-Benennung in Schritt 5 der Closure-Prozedur mitgenommen.

**Lerneintrag — Form: geschärfte Regel.** *Eine fehlende Sektion kann eine fehlende **Warnung**
sein.* Die Ziel-Form führt *Aktueller Stand* nicht, um einen Stand zu tragen, sondern um eine
Tabelle zu **verbieten**: „nicht als Snapshot hier eintragen — sonst driftet sie." Wer beim
Form-Vergleich nur zählt, welche Überschriften fehlen, liest so etwas als Formalie und lässt es
weg. *Weil* eine Sektion, die nur eine Regel trägt und keinen Inhalt, beim Überschriften-Vergleich
genauso aussieht wie eine überflüssige. **Der Prüfsatz:** bei jeder fehlenden Sektion erst lesen,
*was sie tut*, bevor entschieden wird, ob sie gebraucht wird.

**Zwei beobachtbare Closure-Kriterien:**

1. `grep` zählt vier Regelwerk-Zeiger; die Datei trägt die vier Sektionsnamen der Ziel-Form.
2. `doc-check` 224 Dateien 0 Befunde — die entfernten Chronik-Stellen waren keine Ziele, und die
   vier neuen Sektionen bringen keine toten Verweise.

**Offene Risiken und ihr Ausgang:**

- *Die Grenze zwischen Begründung und Chronik ist ein Urteil* — die Zahl *21 Verweise* bleibt,
  weil sie sagt, **warum** die Regel gilt; „Bis slice-076 las sich diese Zeile…" ging, weil es
  sagt, **was mit dem Dokument geschah**. **Ausgang:** weiter offen, `BEO-016` deckt die
  inferentielle Hälfte von Hard Rule 3.7 ab.
- *Die Wellen-Closure-Prozedur ist unangetastet geblieben* — **Ausgang:** gestrichen mit
  Begründung: sie ist a-checks Ausarbeitung von `modul-06` und trägt Belege, die die Ziel-Form
  nicht kennt; sie nebenbei zu kürzen wäre ein Schnitt ohne eigene Begründung.

**Beobachtungs-Register:** keine Beobachtung angefallen — die zwei Risiken oben sind von
bestehenden Zeilen gedeckt.

**Folge-Slices:** die drei Spec-Straten (`BEO-011`-Rest) und `README.md`.

## 8. Sub-Area-Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** berührt wird **Planungs-Harness** (`docs/plan/planning/`),
in der Modus-Deklaration geführt, alle drei Achsen erfüllt.

**Vorgelagert — offene Beobachtungen sichten:** `BEO-002` (Schichten-Zahl ungeprüft) und `BEO-011`
(ungeprüfte Form-Vergleiche) betreffen dieselbe Sub-Area; beide bleiben offen und sind nicht
Gegenstand dieses Slice.

Alle berührten Sub-Areas GF.
