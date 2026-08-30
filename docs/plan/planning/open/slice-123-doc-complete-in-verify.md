# slice-123 — `doc-complete` prüft vor dem Abschluss, nicht danach

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese Datei liegt — eines von
`open/`, `next/`, `in-progress/`, `done/`. Er wechselt nur durch `make slice-mv`
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).

**Bezug:** Maintainer-Wort 2026-08-30. Anlass ist eine gemessene Requirements-Waise:
[AC-FA-CLI-003](../../../../spec/lastenheft.md#ac-fa-cli-003--usage-ausgabe-und-handbuch-verweis)
wird von **keinem** Slice genannt.

**Berührte Spec-Stellen:** — *(keine)*

**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.

**Autor:** Claude (Opus 5), im Auftrag des Maintainers. **Datum:** 2026-08-30.

---

## 1. Ziel

Eine Anforderung ohne referenzierenden Slice fällt beim Abschluss auf, nicht Wochen später: das
Vollständigkeits-Gate hängt an `verify`.

## 2. Definition of Done

- [ ] Die bestehende Waise ist behoben — der Slice, der die Anforderung umgesetzt hat, nennt ihre
      Kennung in der Closure-Notiz. Beleg: `make doc-complete` meldet **0 Waisen**, Exit 0.
- [ ] `doc-complete` hängt im `verify`-Aggregat; [`AGENTS.md`](../../../../AGENTS.md) §4 und
      [`harness/README.md`](../../../../harness/README.md) §Sensors führen es nicht mehr als
      **advisory**, sondern mit seiner Bindung.
- [ ] Die **Ursache** ist als Regel benannt, nicht nur der Einzelfall: wer eine Anforderung
      anlegt, kann ihre Kennung im Plan nicht verlinken (sie existiert noch nicht) — **in der
      Closure-Notiz existiert sie**. Das steht im Workflow-Skelett und in
      [`AGENTS.md`](../../../../AGENTS.md) §5.

- [ ] `make gates` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.
- [ ] Closure-Notiz mit benanntem Lerneintrag geschrieben (§7).
- [ ] Beobachtungs-Register fortgeschrieben.
- [ ] Jedes Risiko aus §6 trägt genau einen Ausgang.

## 3. Plan (vor Code)

| Datei | Änderungs-Art | Begründung |
|---|---|---|
| `done/slice-117-…md` | update | die Kennung in der Closure-Notiz — behebt die Waise |
| [`Makefile`](../../../../Makefile) | update | `doc-complete` ins `verify`-Aggregat |
| [`AGENTS.md`](../../../../AGENTS.md) §4/§5, [`harness/README.md`](../../../../harness/README.md) | update | Bindung statt „advisory"; die Regel |
| `.claude/commands/slice.md` | update | Schritt 8 nennt den Zeitpunkt |

**Auszuführende Gates:** `make verify` (neu mit `doc-complete`), `make gates` — tragend
`doc-targets` (beide Doku-Tabellen). 

### Der Zielkonflikt, aus dem die Waise entstand

Er ist strukturell und trifft **jeden** Slice, der eine Anforderung neu anlegt:

1. IDs werden **referenziert, nicht erfunden** ([`AGENTS.md`](../../../../AGENTS.md) §5) — beim
   Planen existiert die neue Kennung noch nicht.
2. Jede genannte Kennung ist **linkpflichtig** (`ids`, `link-policy: always`) — ein Link auf eine
   Anforderung, die es nicht gibt, macht `doc-check` rot.
3. Also umschreibt der Plan sie („eine neue `AC-FA-CLI`-Kennung"), und die RTM sieht den Slice
   nicht.

Der Ausweg braucht keine Regel-Änderung, nur den richtigen **Zeitpunkt**: bei der Closure ist die
Anforderung geschrieben. Dort kann und muss die Kennung stehen.

## 4. Trigger

**Start:** eingetreten — die Waise ist gemessen (`make doc-complete`, Exit 1).

**Rückführungen:**

- `in-progress` → `open`: falls `doc-complete` über den Bestand weitere Waisen meldet, die eine
  inhaltliche Entscheidung verlangen statt einer Referenz.

## 5. Closure-Trigger

Waise behoben, Gate gebunden, Regel benannt, Gates grün.

**Was bewusst nicht getan wird:** die **Linkpflicht lockern** oder Planning-Dateien von `ids`
ausnehmen. Beides löste den Zielkonflikt, indem es eine Zusage aufgibt — die Kennungs-Linkpflicht
ist genau die Regel, die verhindert, dass Kennungen ins Leere zeigen.

## 6. Risiken und offene Punkte

- *Die Regel verlagert Arbeit in die Closure, die dort vergessen werden kann* — **Ausgang:** <bei Closure>
- *`doc-complete` im `verify`-Aggregat macht jede neue Anforderung sofort abschluss-blockierend* —
  **Ausgang:** <bei Closure>

## 7. Closure-Notiz

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice,
[`AGENTS.md`](../../../../AGENTS.md) §5.)_

**Lerneintrag — Form: <geschärfte Regel | neuer Sensor | benannte Spec-Lücke>**

## 8. Sub-Area-Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** berührt wird die **Gate-/Werkzeug-Schicht** (`Makefile`)
und mit den Deklarations-Tabellen und dem Workflow-Skelett der **Harness-Einstieg**.

**Vorgelagert — offene Beobachtungen sichten:** [`BEO-023`](../observations.md) (ein Prüfer ohne
Gegenstand bleibt unkalibriert) beschreibt dieselbe Familie — hier war der Prüfer nicht ohne
Gegenstand, sondern **ohne Aufruf**: `doc-complete` ist advisory und lief nie.

Alle berührten Sub-Areas GF.
