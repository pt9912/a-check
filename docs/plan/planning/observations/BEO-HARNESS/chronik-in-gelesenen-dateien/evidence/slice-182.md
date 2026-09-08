**Vorgang:** slice-182

**Fund:** Drei Chronik-Stellen in `harness/README.md`, der Datei, die jeder Lauf als Schritt 1
ganz liest:

- §Sensors, Kopf-Absatz: *„(bis slice-079 tat das `gate-consistency`)"* — die Ablösung hält `git`.
- §Rollen, Schluss-Absatz: *„die Review-Serie vom 2026-07-26 lief in einem anderen
  Kontextfenster … Die Reports weisen sich darum ausdrücklich als Selbst-Review aus"* — ein
  Ereignis von vor sechs Wochen, dessen Aussage in den Reports selbst steht.
- §Rollen, Tabellen-Einleitung: *„(angelegt in slice-066, Fund B-9)"* — Provenienz einer Zuordnung,
  die `git log --follow` beantwortet.

**Warum es zählt:** `AGENTS.md` §3.7 verlangt den Indikativ über den Zustand; die Historie lebt in
`git`. Alle drei Stellen kosten jeden Lauf Kontext und ändern an keiner Entscheidung etwas.

**Wie es auffiel:** beim Vergleich mit `v6.5.0` · `templates/harness/README.template.md`, den der
Maintainer angestoßen hat. Nicht durch einen Lauf — die Regel ist ausdrücklich sensorlos.

**Behandlung:** alle drei gestrichen. Der Register-Eintrag bleibt offen.
