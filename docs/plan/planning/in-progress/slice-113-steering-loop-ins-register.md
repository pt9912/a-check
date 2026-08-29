# slice-113 — Steering-Loop ins Register, Datei fällt

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Verantwortlich:** — *(bis zur Priorisierung)*
**Autor:** Claude (Opus 5), im Auftrag des Maintainers. **Datum:** 2026-08-29.
**Berührte Spec-Stellen:** — *(keine)*
**Deckt:** keine `AC-*`; hängt an der Folge-ADR zu [`ADR-0031`](../../adr/0031-heuristik-grenzen-diagnose.md) (R-2).
**Bezug:** Maintainer-Entscheidung 2026-08-29. [Roadmap](../in-progress/roadmap.md).

---

## 0. Trigger

**Start:** die Folge-ADR zu [`ADR-0031`](../../adr/0031-heuristik-grenzen-diagnose.md) liegt vor und trägt die `SL-004`-Begründung ohne Verweis
auf `docs/plan/steering-loop.md`. Vorher nicht — solange die immutable ADR auf die Datei zeigt,
bräche ihr Löschen einen Verweis, den kein Slice reparieren darf.

## 1. Ziel

`docs/plan/steering-loop.md` ist seit slice-101 der **zweite** Zähler für dieselbe Sache. Gemessen:
kein Gegenstand steht in beiden, aber beide führen Gegenstand, Vorfallszahl, Schwelle 3× und
Stand — und heute ist derselbe Vorfall in **beide** gelaufen (`SL-004` auf fünf erhöht, `BEO-005`
mit denselben Belegen `slice-099, slice-100`). Ein Vorfall, zwei Bücher, zwei Stände.

Die Ziel-Form kennt keine solche Datei: der Zähler **ist** das Beobachtungs-Register
(`modul-06` §Das Beobachtungs-Register), die Analyse steht in `## Steering-Loop-Einträge` der
Welle-Notiz — beide Welle-Notizen dieses Repos tragen die Sektion bereits.

## 2. Definition of Done

- [ ] Die sechs `SL-*` stehen als Registerzeilen mit Zähler und Stand *verkörpert in `<Ort>`*;
      der alte Name steht in der `Beobachtung`-Spalte, damit die 154 Zitate lesbar bleiben.
- [ ] `docs/plan/steering-loop.md` ist entfernt; die lebenden Verweise
      ([`AGENTS.md`](../../../../AGENTS.md), `.claude/commands/slice.md`,
      `tools/verify-slice-links.sh`) zeigen aufs Register.
- [ ] Die 23 historischen Zeiger in `done/` und `docs/reviews/` sind entlinkt, ihre **Aussage**
      unverändert — wie in slice-112.

- [ ] `make gates` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.
- [ ] Closure-Notiz mit benanntem Lerneintrag geschrieben (§7).
- [ ] Beobachtungs-Register fortgeschrieben.
- [ ] Jedes Risiko aus §6 trägt genau einen Ausgang.

## 3. Plan (vor Code)

| Datei | Änderungs-Art | Begründung |
|---|---|---|
| `docs/plan/planning/observations.md` | update | sechs Zeilen |
| `docs/plan/steering-loop.md` | entfällt | zweiter Zähler |
| `AGENTS.md`, `.claude/commands/slice.md`, `tools/verify-slice-links.sh` | update | lebende Verweise |
| 23 Dateien in `done/` und `docs/reviews/` | update | tote Zeiger entlinken |

**Auszuführende Gates:** `make gates` (tragend `doc-check`), zum Abschluss `make verify`.

## 4. Trigger

**Start:** siehe §0. **Rückführungen:** `in-progress` → `open`, falls die Folge-ADR die
`SL-004`-Begründung doch beibehält — dann bleibt die Datei und braucht stattdessen eine
Archiv-Grenze.

## 5. Closure-Trigger

Sechs Zeilen im Register, Datei weg, kein toter Zeiger, Gates grün.

**Was bewusst nicht getan wird:** Die Welle-Notizen bleiben unberührt. Sie tragen die Analyse an
ihrem vorgesehenen Ort; sie umzuschreiben wäre Nacherzählung.

## 6. Risiken und offene Punkte

- *Die `SL-*`-Kennungen verschwinden als adressierbare Anker* — **Ausgang:** weiter offen,
  `BEO-021` im Beobachtungs-Register; die Namen bleiben in der `Beobachtung`-Spalte lesbar, aber
  ein Link auf `#sl-004` löst nicht mehr auf.
- *Die Folge-ADR könnte die Begründung anders formulieren als erwartet* — **Ausgang:** weiter
  offen, gedeckt durch §0: ohne sie startet dieser Slice nicht.

## 7. Closure-Notiz

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice,
[`AGENTS.md`](../../../../AGENTS.md) §5; `make verify` prüft das.)_

## 8. Sub-Area-Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** berührt wird **Planungs-Harness**; die drei lebenden
Verweise liegen in **Harness-Einstieg** und **Gate-/Werkzeug-Schicht**, tragen aber je eine Zeile.

**Vorgelagert — offene Beobachtungen sichten:** `BEO-005` ist der Zähler, der heute doppelt
geführt wurde, und bleibt bestehen.

Alle berührten Sub-Areas GF.
