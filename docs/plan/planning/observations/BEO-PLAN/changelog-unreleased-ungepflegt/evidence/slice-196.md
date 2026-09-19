**Vorgang:** slice-196

**Fund:** Der Abschnitt `[Unreleased]` trug beim Beginn dieses Slice **drei** Einträge — die von
slice-194 und slice-195 —, während seit `v0.19.0` **265** Commits und **58** Slices gelandet waren.
Gemessen: `git log v0.19.0..HEAD -- internal/` nennt für den ganzen Zeitraum genau **einen** Slice
(194), `spec/` änderte sich in 194 und 154; die übrigen **56** sind Harness-Arbeit.

**Das ist der dritte Vorfall.** Die ersten beiden (slice-127, slice-133) fielen — wie dieser — bei
der **Release-Vorbereitung** auf. Der Eintrag steht damit bei 3×, und sein `state.md` sagt für
diesen Fall: es ist eine Lücke und verlangt **Guide oder Sensor**.

**Warum der Ausgang ein Guide wurde und kein Sensor.** Die Prüfung, die es bräuchte, müsste
entscheiden, ob eine Änderung *eintragspflichtig* ist — eine Änderung an `spec/` ist es, eine an
`internal/hexagon/core/rules.go` auch, ein Sensor-Umbau nicht. Ein Gate auf „hat einen Eintrag"
bräuchte dieses Urteil als Eingabe und wäre entweder falsch-positiv auf jedem Harness-Slice oder
abgeschaltet. Der Fehler ist ohnehin kein Wissens-, sondern ein **Zeitpunkt**-Fehler: Der Eintrag
entsteht bei der Release-Vorbereitung statt im Slice.

**Behoben** mit der geschärften Regel in [`AGENTS.md`](../../../../../../../AGENTS.md) §6 Schritt 7:
Die CHANGELOG-Zeile gehört in den Slice, der den Vertrag berührt — und der Satz nennt die Grenze
gleich mit („kein Gate deckt das"), damit sie nicht für eine Durchsetzung gehalten wird.

**Nachgetragen:** der Abschnitt selbst — 58 Slices, in Gruppen, jede Kennung genannt.
