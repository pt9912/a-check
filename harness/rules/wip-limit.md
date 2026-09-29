# WIP-Limit = 1 pro Lauf; auf dem Hauptzweig: höchstens ein Slice in `in-progress/`

*Ausgelagert aus [`AGENTS.md`](../../AGENTS.md) §5, Regel 8 — seit slice-201. Der Wortlaut ist unverändert; die Links sind der neuen Tiefe nachgeführt.*

**WIP-Limit = 1 pro Lauf.** Die Baseline zählt pro Rolleninhaber und **Lauf**,
  nicht pro Rolle — mehrere Läufe derselben Person zählen einzeln, wenn
  `Verantwortlich:` sie per Zweig unterscheidet (`modul-05` §Trigger je
  Lifecycle-Übergang und WIP-Limit). Auf dem Hauptzweig, wo der `git mv` den
  Anspruch sichtbar macht, bedeutet das: **höchstens ein** Slice in
  `in-progress/` (die Roadmap zählt nicht mit). Das ist eine harte Obergrenze,
  kein Vorschlag: zwei aktive Slices teilen sich einen Gate-Nachweis und eine
  Closure-Aufmerksamkeit, und beides trägt nur einmal. **Null ist zulässig** —
  nach jedem Abschluss der Normalfall, bis der nächste Slice gezogen wird;
  ein Maximum ist kein Minimum.
