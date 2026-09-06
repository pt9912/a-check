# `AGENTS.md` §5 hinkt einem neuen konstanten Baseline-DoD-Item hinterher

**Sub-Area:** Harness-Einstieg

Die Baseline erweitert die pro Slice konstant-nicht-gezählten DoD-Posten um einen Review-Report
(`modul-05-planning-harness.md` §Ziel-Form: Slice, `slice.template.md`). `AGENTS.md` §5 führt die
eigene Liste dieser Posten in eigenen Worten und die `tasks-ignore-pattern`-Regex in `.d-check.yml`
kennt kein Review-Muster — beide sind mit dem Baseline-Zuwachs nicht mehr deckungsgleich, und die
Frage, ob a-check den Review-DoD-Punkt verpflichtend (wie die Baseline) oder opt-in (wie bisher)
führt, ist eine Architektur-Entscheidung, keine reine Übernahme.
