# Slice implementieren (Harness)

Argument: $ARGUMENTS

Dieser Command führt die **Implementer**-Rolle (Modul 8/9) für *einen* Slice.

**Der 8-Schritt-Workflow ist in [`AGENTS.md`](../../AGENTS.md) §6 verkörpert und als
[`slice`](slice.md)-Command aufbereitet — lies ihn und folge ihm.** Diese Datei
wiederholt ihn nicht (Zeiger statt Kopie — zweiter Wortlaut driftet); sie trägt nur
die Rollen-Rahmung, die über den Workflow hinausgeht.

**Eingang:** der Slice in `docs/plan/planning/in-progress/` — der Claim (`open →
next → in-progress`) ist per `make slice-mv SLICE=… TO=in-progress` auf dem
Hauptzweig geleistet, **vor** der Arbeit; das Kopf-Feld `Verantwortlich:` nennt den
Lauf. **Der Claim entfernt den Ruhe-Marker aus der Roadmap** — er steht genau
dann, wenn `in-progress/` leer ist, und kommt mit der Archivierung zurück
(`doc-planning` prüft beide Richtungen; `seit slice-207`).

**Ausgang:** Commit-Range plus Slice-Plan-Verweis an den Reviewer — in **frischem
Kontext**, kein Selbst-Review; ein `fork`-Subagent erbt den Kontext und zählt nicht.
Reviewer-Skill: [`.harness/skills/reviewer.md`](../../.harness/skills/reviewer.md).

**Rückführungen sind Disziplin, kein Scheitern** (Modul 5): Slice zu groß →
`make slice-mv … TO=next` (zurück zur Zerlegung); blockiert → `TO=open` (Carveout
prüfen, Modul 7).

**Die Durchsetzungsschicht bindet den Lauf:** der Stop-Hook verlangt einen frischen
`make gates`-Lauf nach jeder Inhaltsänderung (inkl. Commits), und der Commit-Guard
scannt Messages — Commits daher immer via Message-Datei (`git commit -F <datei>`).
Beide ersetzen nicht die Berichtspflicht: ausgeführte Sensors und Restrisiken nennen,
keine Erfolgsmeldung ohne Gate-Ausgabe (§6 Schritt 8).

Hier endet die Implementer-Rolle. Review, Verifikation und Closure laufen in
getrennten Kontexten (Modul 8).
