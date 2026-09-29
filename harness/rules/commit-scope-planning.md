# Commit-Scope `(planning)` berührt ausschließlich `docs/plan/planning/`

*Ausgelagert aus [`AGENTS.md`](../../AGENTS.md) §5, Regel 2 — seit slice-201. Der Wortlaut ist unverändert; die Links sind der neuen Tiefe nachgeführt.*

**Commit-Scope `(planning)`:** ein Commit mit diesem Scope (`docs(planning)`,
  `fix(planning)`, `chore(planning)`) berührt **ausschließlich**
  `docs/plan/planning/`. Wandert Substanz eines anderen Bereichs mit, ist das ein
  eigener Commit mit passendem Scope. Durchgesetzt durch `make commit-scope-check`;
  jeder Commit wird an der Fassung gemessen, die zu **seinem** Zeitpunkt galt.
  **Nur dieser Scope ist geregelt:** Bei `docs(spec)` und `docs(adr)` ist der
  Fremd-Bereich legitim, und eine Regel, die den Bestand massenhaft bricht, wird
  abgeschaltet statt befolgt. Ein weiterer Scope wird erst geregelt, wenn er
  auffällt — und dann gemessen, nicht geraten.
