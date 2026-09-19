**Vorgang:** slice-194

**Fund:** Die Meldung eines externen Konsumenten (`hexslice-architecture`, gemessen gegen `v0.19.0`)
und ihre Reproduktion im Repo: Trägt eine `port`-Schicht ihre `direction`, endet ihr Glob auf dem
Richtungssegment (`…/createorder/ports/outbound/**`). Die Scope-Ableitung von `port-locality`
schnitt damals das **letzte** Segment als „Port-Ordner-Marker" ab — das war hier das
Richtungssegment selbst. Der Scope fiel auf `…/ports`, `appTreeContains` wurde false, und die
Regel **schwieg**: kein Befund, Exit 0, bei einem echten slice-übergreifenden Import.

Gemessen, gleicher Baum, gleicher Verstoß, nur das Glob geändert: `…/ports/**` → `port-locality: 1`,
Exit 1 · `…/ports/outbound/**` → `gesamt: 0 Befund(e)`, Exit 0. Kontrolle mit **identischer**
Config und fremder `app`-Slice als Importziel: `lateral-slice: 1`, Exit 1 — die Config war nicht
tot, die Regel war **still**.

**Die zweite Zusage, an der die erste zerbricht,** ist die Anweisung aus dem Benutzerhandbuch §4
und [ADR-0036](../../../../../adr/0036-port-richtung-inbound-outbound.md): Für
`port-direction-mismatch` braucht es **getrennte Port-Schichten** (`inbound`/`outbound`), weil eine
Schicht genau **eine** Richtung trägt. Wer sie befolgt, gibt dem Port-Glob das Richtungssegment —
und befolgt damit die eine Regel, indem er die andere abschaltet. Beide sind für sich begründet,
und zusammen sind sie nicht befolgbar.

**Die stille Ausprägung, die dieser Beleg nachträgt.** Der Eintrag beschreibt seine Erkennung
bislang als *„an einem Sensor, der rot wird, obwohl beide Regeln befolgt wurden"*. Hier wird
**nichts rot**: Die Regel schweigt, der Lauf endet grün, und der Verstoß bleibt unbeurteilt. Kein
Gate fängt das — die Zahl der Befunde sinkt nicht, sie ist null.

**Wie es auffiel:** nicht im Repo, sondern beim **Konsumenten** — und nur, weil dort gemessen
wurde. Der Lab-Baum trägt seit dem Umbau einen Kommentar, der die Kopplung festhält; er war die
einzige Stelle, an der die Grenze überhaupt benannt war.

**Behoben mit [ADR-0040](../../../../../adr/0040-portscope-richtungssegment.md):** Die Ableitung
zieht das Richtungssegment zusätzlich ab — nur wenn die Schicht ihre `direction` trägt, der Port im
App-Baum liegt und der Schnitt den Scope **in** ihn bringt statt über ihn hinaus —, und eine
weitere Advisory-Diagnose macht die Restfälle laut. Der Test
`TestPortLocalityDirectionSegmentKeepsRule` fährt beide Glob-Varianten gegen denselben Verstoß; die
Mutations-Probe war **rot** mit der Meldung *„ein Port-Glob mit Richtungssegment darf
port-locality nicht abschalten, got []"*.
