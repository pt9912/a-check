# slice-053 — Etappe D (2/3): Lifecycle vervollständigen, Roadmap-Drift sichtbar machen

**Status:** *(der Zustand ist das Verzeichnis dieser Datei, nicht dieses Feld — korrigiert in slice-063)* — zweiter Schnitt der **Etappe D (Form)** aus
[slice-048 §5](../done/slice-048-modul-delta-lesen.md).
**Deckt:** **B-6** (Lifecycle ohne Rückführungen, WIP-Limit undeklariert), **B-18**
(`next/` existiert nicht) und **B-12** (Roadmap ohne Drift-Log).
**Nicht hier:** B-15 (AC-Form) — 3/3. [Roadmap](../in-progress/roadmap.md).

---

## 1. Auslöser

Drei Form-Lücken, die zusammengehören, weil sie dieselbe Frage betreffen: *woran sieht ein
Außenstehender, wie geplant wurde?*

- **B-6:** die Baseline führt **fünf** Lifecycle-Übergänge, darunter zwei **Rückführungen** —
  `in-progress → next` (zu groß) und `in-progress → open` (blockiert).
  [`AGENTS.md`](../../../../AGENTS.md) §5 nennt nur die Vorwärtskette. Für die Rückwege gibt es
  weder Regel noch Präzedenz; das **WIP-Limit = 1** („harte Größe, kein Vorschlag") ist nirgends
  deklariert, wird aber faktisch eingehalten.
- **B-18:** `docs/plan/planning/next/` existiert **gar nicht**, obwohl §5 den Zustand nennt. In
  53 Slices wurde er genau einmal benutzt (slice-009, `9cb5ffa`) und danach mitsamt Verzeichnis
  aufgelöst. Ein deklarierter Zustand ohne Ort ist eine stille Setzung.
- **B-12:** die Roadmap trägt vier der fünf Pflicht-Abschnitte; es fehlt **Historische
  Trigger-Verschiebungen**. Ohne dieses Drift-Log ist jede Umplanung still — und die Roadmap
  wirkt im Rückblick, als wäre sie immer so geplant gewesen.

## 2. Betroffene Module

- [`AGENTS.md`](../../../../AGENTS.md) §5 — Rückführungen und WIP-Limit (B-6).
- `docs/plan/planning/next/README.md` — der fehlende Ort (B-18).
- [`roadmap.md`](../in-progress/roadmap.md) — Drift-Log-Abschnitt, aus dem belegbaren Verlauf befüllt (B-12).

Zwei Schichten (Agenten-Briefing, Planungs-Doku).

## 3. Auszuführende Gates

`make gates` und `make verify`. Kein neuer Sensor — die drei Punkte sind **Form**, und für Form
ohne Ermessensspielraum gibt es hier bereits Prüfer: `doc-planning` hält Roadmap ↔ `in-progress`
konsistent, `verify-slice-form` die Slice-Größe. Ein zusätzlicher Zähler über Roadmap-Abschnitte
wäre Prüfung um der Prüfung willen.

## 4. Was bewusst nicht getan wird

- **Kein Rückwirken auf die Historie.** Das Drift-Log beginnt mit den Verschiebungen, die im Repo
  **belegt** sind (Slice-Dokumente, Roadmap-Text). Ältere Umplanungen, für die es keinen Beleg
  gibt, werden nicht rekonstruiert — ein erfundenes Drift-Log wäre schlimmer als keines.
- **Kein Zwang zu `next/`.** Das Verzeichnis existiert wieder und ist beschrieben; ob ein Slice
  ihn durchläuft, bleibt Planungs-Entscheidung. Die Baseline verlangt den *Zustand*, nicht seine
  Benutzung in jedem Fall.

## 5. DoD

- [x] [`AGENTS.md`](../../../../AGENTS.md) §5 führt alle **fünf** Übergänge inklusive der zwei
      Rückführungen mit ihrer je eigenen Bedingung und deklariert das **WIP-Limit = 1** (B-6).
- [x] `docs/plan/planning/next/` existiert mit einer `README.md`, die den Zustand und seine
      Ein-/Ausgänge beschreibt (B-18).
- [x] [`roadmap.md`](../in-progress/roadmap.md) trägt den Abschnitt **Historische Trigger-Verschiebungen**,
      befüllt aus belegten Umplanungen; `make gates` und `make verify` grün (B-12).

## 6. Closure-Notiz

**Geliefert:** [`AGENTS.md`](../../../../AGENTS.md) §5 führt alle fünf Lifecycle-Übergänge und das
WIP-Limit, `docs/plan/planning/next/` existiert wieder mit erklärender `README.md`, und die
Roadmap trägt das Drift-Log mit sieben belegten Verschiebungen.

**Lerneintrag — Form: benannte Spec-Lücke.**
> **Ein deklarierter Zustand ohne Ort ist eine stille Setzung — und sie hält sich lange.**
> `AGENTS.md` nannte `next/` über 44 Slices hinweg, während das Verzeichnis nicht existierte;
> niemand stolperte darüber, *weil* der Normalweg daran vorbeiführt. Dieselbe Klasse traf in
> slice-050 die Closure-Pflicht (gelebt, nie geschrieben) — hier ist es umgekehrt (geschrieben,
> nie vorhanden). Prüfsatz: *jede Regel, die einen Ort nennt, gegen die Existenz des Ortes
> prüfen — und jede gelebte Praxis gegen die Existenz einer Regel.* Beide Richtungen erzeugen
> stille Setzungen.

**Zwei beobachtbare Closure-Kriterien:**

1. `make gates` und `make verify` grün auf demselben Stand (je Exit 0) — belegt.
2. Jede Zeile des Drift-Logs nennt einen Beleg im Repo (Slice-Abschnitt oder ADR); keine Zeile
   ist rekonstruiert.

**Beobachtung zum Drift-Log:** sieben belegbare Verschiebungen an **einem** Tag, davon fünf
Verwerfungen oder Vertagungen. Das ist kein Planungs-Mangel, sondern das sichtbare Ergebnis davon,
dass Entscheidungen an Messungen hingen — vorher stand das verstreut in Slice-Dokumenten und
damit faktisch nirgends.

**Folge-Slices:** [slice-054](slice-054-ac-form.md) (Etappe D, 3/3 — B-15, AC-Form).

## 7. Sub-Area-Modus

Alle berührten Sub-Areas GF.
