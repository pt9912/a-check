# Eine neue Regel gilt für den Bestand, den sie beim Entstehen sieht

**Sub-Area:** Harness-Einstieg

Eine Zitier-, Form- oder Struktur-Regel wird **für den Neubestand** entschieden — und der
Altbestand, der ihre Klasse schon trägt, bleibt stehen. Er fällt dann nicht durch ein Gate auf,
sondern durch ein **fremdes**: Der Nachzug des Altbestands ist niemandes Aufgabe, und die Regel
sieht aus, als wäre sie durchgesetzt.

**Der belegte Fall:** Seit `slice-176` zitieren einfrierende Artefakte **Kennung statt Adresse**.
Drei `Accepted`-ADRs entstanden **davor** und zitierten Adressen; der Zeitdokument-Sweep zog die
Adressen nach und erzeugte damit drei Kern-Drift-Befunde, die der Immutabilitäts-Sensor bei
**jedem** Release meldet — obwohl keine Entscheidung sich geändert hat.

**Woran es auffällt:** an einem Sensor, der für eine **andere** Klasse zuständig ist. Die eigene
Regel hat kein Gate für den Altbestand, weil sie ihn beim Entstehen nicht sah.

**Abgrenzung zu [`BEO-HARNESS/massen-ersetzung-trifft-die-historische-aussage`](../massen-ersetzung-trifft-die-historische-aussage/observation.md):**
Dort trifft eine Ersetzung eine Aussage, die **damals wahr war**; hier fehlt einer Regel der
Nachzug in den Bestand, den sie hätte treffen sollen.
