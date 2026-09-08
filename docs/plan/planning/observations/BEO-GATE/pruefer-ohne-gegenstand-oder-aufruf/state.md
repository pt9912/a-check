**Stand:** verkörpert in [`make dcheck-phrase-selftest`](../../../../../../Makefile)
(`tools/dcheck-phrase-selftest.sh`, im `gates`-Aggregat) — die **Werkzeug**-Hälfte
`seit slice-168`, die **Korpus**-Hälfte `seit slice-169`.

Gedeckt ist **eine** Konfiguration: die `reviews`-Trigger-Phrase. Für sie prüft der Lauf beide
Ausfallarten — dass `d-check` auf die gewählte Formulierung nicht mehr reagiert (Fixtures, vier
Kontrollen), und dass a-checks eigener Bestand sie nicht mehr trägt (Nichtleerheit gegen die
**Kandidatenmenge des Moduls**: ein DoD-Haken in einem flachen `done/`-Slice). Die zweite ist
zweimal ausgefallen — slice-120 und slice-165 — und war bis slice-169 ungedeckt.

Die Korpus-Hälfte liest ihr Kandidatenverzeichnis **aus `.d-check.yml`** statt aus einer Kopie,
fail-closed. Das schließt die zweite Hälfte des Review-Befunds zu slice-168: Eine Kopie neben dem
Original blieb grün, nachdem das Original gebrochen wurde.

**Der Eintrag bleibt stehen, nicht gestrichen** — die Klasse ist breiter als das eine gedeckte
Muster. Gedeckt ist, was einen **belegten** Ausfall hat; für die übrigen phrasen-basierten Felder
in `.d-check.yml` gibt es keinen Vorfall, und ein Sensor ohne Anlass ist selbst eine Behauptung.
Fällt eines davon aus, ist es ein neuer Beleg hier — kein neuer Eintrag.

**Zwei Kandidaten stehen ausdrücklich nicht drin**, beide weil sie **laut** ausfallen statt still:
`versions.current-from` (Beleg `evidence/slice-173.md`) — `d-check` bricht bei nicht auflösbarem
Anker ab, statt grün zu melden. Und `structure.tasks-ignore-pattern`: trifft es nichts mehr,
zählen die konstanten DoD-Posten mit, und `doc-structure` meldet `section-oversized`. Beiden fehlt
genau die Gefährlichkeit, die diesen Eintrag ausmacht — die leere Prüfmenge, die grün meldet.
