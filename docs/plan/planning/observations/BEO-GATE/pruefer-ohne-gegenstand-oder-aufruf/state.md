**Stand:** verkörpert in [`make dcheck-phrase-selftest`](../../../../../../Makefile)
(`tools/dcheck-phrase-selftest.sh`, im `gates`-Aggregat) — die **Werkzeug**-Hälfte
`seit slice-168`, die **Korpus**-Hälfte `seit slice-169`.

Beide Ausfallarten sind damit gedeckt: dass `d-check` auf die gewählte Formulierung nicht mehr
reagiert (Fixtures, vier Kontrollen), und dass a-checks eigener Bestand sie nicht mehr trägt
(Nichtleerheit gegen `done/`, zwei Kontrollen). Die zweite ist zweimal ausgefallen — slice-120 und
slice-165 — und war bis slice-169 ungedeckt.

Die Korpus-Hälfte liest das Muster **aus `.d-check.yml`** statt aus einer Kopie, fail-closed. Das
schließt die zweite Hälfte des Review-Befunds zu slice-168: Eine Kopie neben dem Original blieb
grün, nachdem das Original gebrochen wurde.

**Der Eintrag bleibt stehen, nicht gestrichen** — die Klasse ist breiter als die zwei gedeckten
Muster. `.d-check.yml` führt **vierzehn** phrasen-basierte Felder; gedeckt sind die zwei mit
belegtem Ausfall. Für die übrigen zwölf gibt es keinen Vorfall, und ein Sensor ohne Anlass ist
selbst eine Behauptung. Fällt eines davon aus, ist es ein neuer Beleg hier — kein neuer Eintrag.

`versions.current-from` ist der nächstliegende Kandidat (Beleg `evidence/slice-173.md`), fällt aber
**laut** aus: `d-check` bricht bei nicht auflösbarem Anker ab, statt grün zu melden. Ihm fehlt
genau die Gefährlichkeit, die diesen Eintrag ausmacht.
