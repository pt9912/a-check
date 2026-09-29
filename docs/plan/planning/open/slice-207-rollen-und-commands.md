# slice-207 — Rollen- und Command-Verdrahtung übernehmen

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt. Er wechselt nur durch `git mv`.

**Welle:** ohne Welle — die Closure-Bedingung wäre die eigene DoD.

**Bezug:** Folge-Slice aus
[slice-205](../done/slice-205-init-tool-vergleich.md) (Entscheidung
„Übernehmen").
[`AC-QA-02`](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze).

**Berührte Spec-Stellen:** — · Der Slice berührt kein Spec-Stratum.

**Verantwortlich:** — bis zur Priorisierung.

**Autor:** Claude. **Datum:** 2026-09-29.

**Lerneintrag — Form:** neuer Sensor.

---

## 1. Ziel und Abgrenzung

**Ziel:** Die Rollen- und Command-Verdrahtung des Generators ist übernommen —
`.claude/agents/` (sechs Rollen-Subagents, die per Zeiger auf die Commands
verweisen) und `.claude/commands/` (`plan-welle`, `implement-slice`,
`close-welle`). **Die Inhalte der Commands werden an a-checks Prozess
angepasst** (`AGENTS.md` §6 8-Schritt-Workflow statt des Generators-Wortlauts;
a-checks Adaptionen sind im AGENTS.md-Träger verankert, den die Commands
referenzieren).

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Inhaltliche Abweichungen vom Modul-8-Prozess.** *Schicht-Abgrenzung*: die
  Commands tragen a-checks Prozess; eigene Erfindungen sind Out-of-Scope.

## 2. Definition of Done

- [ ] `.claude/agents/` (sechs Rollen) und `.claude/commands/` (drei
      Lifecycle-Commands) existieren, an a-checks Prozess angepasst.
- [ ] Die Zeiger-Kette ist geprüft: jeder Agent verweist auf seinen Command,
      jeder Command auf die kanonischen Quellen; `make doc-check` grün.
- [ ] Unabhängiger Review, Report unter [`docs/reviews/`](../../../../docs/reviews/README.md).
- [ ] Closure-Notiz mit Steering-Loop-Lerneintrag; Register fortgeschritten;
      jedes Risiko aus §6 trägt einen Ausgang.

`make gates` und `make verify` grün.

## 3. Plan (vor Code)

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| `.claude/agents/` (6 Dateien) | neu | Rollen-Subagents mit Zeiger auf die Commands |
| `.claude/commands/` (3 Dateien) | neu | plan-welle, implement-slice, close-welle — an `AGENTS.md` §6 angepasst |

## 4. Trigger

**Start** (`open` → `in-progress`): das WIP-Limit ist frei.

**Rückführungen — vorab benennen:**

- `in-progress` → `next` (zu groß): entfällt — drei Datei-Familien mit
  Zeiger-Logik.
- `in-progress` → `open` (blockiert): entfällt — kein Blocker absehbar.

## 5. Closure-Trigger

DoD vollständig, `make gates` und `make verify` grün, Closure-Notiz mit
Lerneintrag.

## 6. Risiken und offene Punkte

- **Die Commands frieren den Prozess** — eine AGENTS.md-Änderung veraltet die
  Command-Kopien. — **Ausgang:** bei Closure (die Commands referenzieren
  AGENTS.md statt es zu kopieren — Zeiger, nicht Kopie).

## 7. Closure-Notiz

*(wird vor dem `git mv` nach `done/` gefüllt)*

## 8. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** GATE (Achsen 1, 2, 3 ✓).

**Vorgelagert — offene Beobachtungen sichten:** das Register wurde durchgegangen
(2026-09-29): keine Treffer in GATE für diesen Vorgang.

**Modus-Begründungsblock:** alle berührte Sub-Areas GF.
