# slice-108 — `AGENTS.md`: Hard Rule 3.7 und die drei Regelwerk-Zeiger

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.
**Autor:** Claude (Opus 5), im Auftrag des Maintainers. **Datum:** 2026-08-29.
**Berührte Spec-Stellen:** — *(keine)*
**Deckt:** `BEO-011` (die Zeile zu `AGENTS.md` §3.7).
**Bezug:** [slice-105](../done/slice-105-form-review-nachholen.md) §4.
[Roadmap](../in-progress/roadmap.md).

---

## 1. Ziel

`AGENTS.md` fehlt eine **Hard Rule**, und sie ist genau die Regel, die der Maintainer heute
mündlich durchgesetzt hat.

Der Form-Review hat die Zeile als *ungeprüft* markiert — Namensvariante oder echte Lücke? Gelesen:
`3.1`–`3.6` sind Namensvarianten (`Docker-only` ↔ `Docker/make-only`), **`3.7 Ein Kommentar
beschreibt, was da ist` fehlt wirklich.** Sie sagt: ein Kommentar trägt eine von fünf Klassen —
Zusage · Kopplung · Abgrenzung · Rang-Zeiger · Grenze — und schreibt an den, der die Stelle
*ändert*, nicht an den, der die Entscheidung *trifft*. Kein Konjunktiv über verworfene
Alternativen, keine Beschreibung abwesenden Texts. *„Die Abwägung gehört in die ADR, die Historie
in `git` … Was daneben steht, liest jeder Lauf mit und bezahlt es mit Kontext."*

Das ist wörtlich die Vorgabe, unter der slice-103 und slice-106 die Chronik aus `conventions.md`
und der Roadmap genommen haben — nur war sie dort eine Ansage und ist hier eine adoptierte Regel.

## 2. Definition of Done

- [x] `AGENTS.md` trägt `### 3.7`, auf a-check bezogen und mit den fünf Klassen — Beleg: Diff.
- [x] Die drei Regelwerk-Zeiger stehen (Datei-Zeiger in §1, §3.7, §4) — Beleg: Zählung 3/3.
- [x] `BEO-011` verliert die `AGENTS.md`-Hälfte, der Rest der Zeile bleibt — Beleg: Register.

- [x] `make gates` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.
- [x] Closure-Notiz mit benanntem Lerneintrag geschrieben (§7).
- [x] Beobachtungs-Register fortgeschrieben.
- [x] Jedes Risiko aus §6 trägt genau einen Ausgang.

## 3. Plan (vor Code)

| Datei | Änderungs-Art | Begründung |
|---|---|---|
| `AGENTS.md` §1 | update | Datei-Zeiger auf `modul-09` §Ziel-Form: AGENTS.md |
| `AGENTS.md` §3 | neu | Hard Rule 3.7 samt Sektions-Zeiger |
| `AGENTS.md` §4 | update | Sektions-Zeiger auf `modul-13` |
| `docs/plan/planning/observations.md` | update | `BEO-011` verkleinern |

**Auszuführende Gates:** `make gates` (tragend `doc-check` und `doc-targets`, weil §4 angefasst
wird), zum Abschluss `make verify`. **Kein neuer Sensor** — die Regel ist inferentiell, sie prüft
ein Urteil über Kommentar-Text und gehört damit dem Review, nicht einem Zähler.

## 4. Trigger

**Start:** unmittelbar; die Lücke ist gemessen und der Maintainer setzt die Regel bereits durch.

**Rückführungen:** `in-progress` → `open`, falls die Formulierung der Regel eine
Konventions-Entscheidung verlangt, die über das Übernehmen hinausgeht.

## 5. Closure-Trigger

Die Regel steht, die Zeiger sind gezählt, `make gates` und `make verify` sind grün.

**Was bewusst nicht getan wird:** Der Bestand wird **nicht** gegen die neue Regel durchgesehen.
Sie gilt ab jetzt für neue und geänderte Stellen; ein rückwirkender Durchgang über alle Kommentare
im Repo wäre ein eigener Slice und bräuchte eine Messung, wie groß er ist. Ebenso bleibt die
`README.md`-Hälfte von `BEO-011` offen — anderer Sub-Area, anderer Schnitt.

## 6. Risiken und offene Punkte

- *Die Regel ist inferentiell und hat keinen Sensor* — **Ausgang:** weiter offen, `BEO-016` im
  Beobachtungs-Register. `modul-13` nennt genau diesen Fall; ein Zähler über Kommentar-Klassen
  wäre Schein-Genauigkeit.
- *Der Bestand ist ungeprüft* — **Ausgang:** weiter offen, ebenfalls `BEO-016`.

## 7. Closure-Notiz

**Geliefert:** `AGENTS.md` trägt Hard Rule 3.7 und drei Regelwerk-Zeiger (0/3 → 3/3).

**Lerneintrag — Form: benannte Spec-Lücke.** *Eine Regel, die der Maintainer mündlich durchsetzt,
steht in der adoptierten Baseline — und fehlte im Repo.* Die Vorgabe „keine Chronik oder Forensik,
das müllt den Kontext des Code-Agenten zu" kam heute als Ansage; sie ist wörtlich `modul-09`/
`grundlagen-harness-dateien` §Was ein Kommentar trägt, verkörpert als Hard Rule 3.7 der Ziel-Form.
Zwei Slices (103, 106) haben danach gehandelt, **ohne** dass die Regel im Repo stand. *Weil* der
Form-Review sie als „ungeprüft" markiert und niemand nachgelesen hatte — und weil eine mündlich
befolgte Regel sich anfühlt wie eine vorhandene. **Der Prüfsatz:** wer eine Anweisung befolgt,
prüft, ob sie im Repo steht; sonst gilt sie nur, solange derselbe Mensch zusieht.

**Zwei beobachtbare Closure-Kriterien:**

1. `grep` zählt drei Regelwerk-Zeiger in `AGENTS.md`, §3 führt sieben Hard Rules statt sechs.
2. `doc-check` 223 Dateien 0 Befunde und `doc-targets` Exit 0 — der neue §4-Zeiger hat die
   Target-Tabelle nicht verschoben.

**Offene Risiken und ihr Ausgang:**

- *Die Regel ist inferentiell und hat keinen Sensor* — **Ausgang:** weiter offen, `BEO-016`.
  `modul-13` verlangt genau das: einen Sensor zu behaupten, wo keiner steht, wäre selbst eine
  Harness-Lüge. Die Grenze steht in der Regel selbst.
- *Der Bestand ist nicht gegen die Regel durchgesehen* — **Ausgang:** weiter offen, ebenfalls
  `BEO-016`. Ein rückwirkender Durchgang bräuchte erst eine Messung seines Umfangs.

**Beobachtungs-Register:** `BEO-011` auf 2× erhöht und verkleinert — die `AGENTS.md`-Hälfte ist
gelesen und war eine echte Lücke; `BEO-016` neu angelegt.

**Folge-Slices:** `docs/plan/planning/README.md` (vier Sektionen, vier Zeiger), danach die
Spec-Straten.

## 8. Sub-Area-Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** berührt wird **Harness-Einstieg** (`AGENTS.md`), erfüllt
alle drei Inklusions-Achsen (slice-101).

**Vorgelagert — offene Beobachtungen sichten:** `BEO-011` (diese Zeile), `BEO-009` und `BEO-010`
betreffen dieselbe Sub-Area und bleiben offen — sie sind Gegenstand eigener Schnitte.

Alle berührten Sub-Areas GF.
