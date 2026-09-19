# Der lokale Pre-Flight und der CI-Lauf sind nicht dieselbe Menge

**Sub-Area:** Gate-/Werkzeug-Schicht

Ein Repo hat einen lokalen Aggregat-Lauf („fahre das, was die CI fährt") und einen Workflow, der
**mehr** fährt. Wer den Aggregat-Lauf als Pre-Flight benutzt, bekommt lokal grün und in der CI rot —
und der Unterschied ist keine Nachlässigkeit des Aufrufers, sondern eine **Differenz der Mengen**,
die niemand nennt.

**Woran es auffällt:** am ersten Push, der die fehlenden Schritte verletzt. Bis dahin ist der
Pre-Flight ein Beleg, der mehr zu decken scheint, als er deckt.

**Der belegte Fall:** `make ci` ist `gates` + `image-test`; der Workflow
`.github/workflows/ci.yml` fährt zusätzlich `trace-check`, `commit-scope-check` und
`doc-immutable` über die **Commit-Range**. Am 2026-09-19 war `make ci` lokal grün, während der
CI-Lauf desselben Stands rot war.

**Kein Sensor deckt die Differenz.** Beide Mengen sind für sich korrekt deklariert; dass die eine
die andere nicht enthält, sagt keine der beiden Deklarationen.
