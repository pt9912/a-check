**Vorgang:** slice-197

**Fund:** `make doc-immutable RANGE=v0.19.0..HEAD` meldet drei `Accepted`-ADRs mit geändertem Kern
([`ADR-0017`](../../../../../adr/0017-relative-resolution-modus.md), [`ADR-0018`](../../../../../adr/0018-exclude-scan-scope.md), [`ADR-0038`](../../../../../adr/0038-dependabot-als-hebungskanal.md)) — **Exit 2**. Der Befund ist vorbestehend (dieselbe Messung
ohne die Commits dieser Sitzung: ebenfalls 3, Exit 2).

**Die Ursache ist kein Fehler, sondern eine Regel-Lücke:** Die drei ADRs zitieren ihre
Zeitdokumente als **Adresse**; der Zeitdokument-Sweep hat sie in Unterverzeichnisse bewegt und die
Verweise **auf** sie nachgezogen — auch in den ADR-Körpern. Seit `slice-176` gilt für einfrierende
Artefakte **Kennung statt Adresse**; die drei entstanden davor, und niemand hatte den Nachzug als
Aufgabe.

**Wie es auffiel:** an einem Sensor, der für eine **andere** Klasse zuständig ist — dem
Immutabilitäts-Sensor, bei der Vorbereitung eines Releases. Die eigene Regel hat kein Gate für den
Altbestand, weil sie ihn beim Entstehen nicht sah.

**Nachgezogen** mit `slice-197` (die drei ADRs zitieren jetzt Kennungen); die historische Hälfte
ist in [`MR-024`](../../../../../../../harness/conventions.md#mr-024) deklariert und löst sich mit dem nächsten Release selbst ein.
