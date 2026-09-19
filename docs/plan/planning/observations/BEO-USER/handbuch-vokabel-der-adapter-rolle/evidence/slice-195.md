**Vorgang:** slice-195

**Fund:** Abschnitt 3.7 des Benutzerhandbuchs zeigte als Beispielstruktur
`internal/adapters/{inbound,outbound}/…   # Adapter`. Dasselbe Dokument führt ab §4 und in der
Regel-Tabelle seit [ADR-0036](../../../../../adr/0036-port-richtung-inbound-outbound.md):
an `role: adapter` gilt `driving`/`driven`, an `role: port` `inbound`/`outbound` — „ein Port
*treibt* nichts, er wird benutzt; ein Adapter ist nicht *eingehend*". Zwei Vokabulare für dieselbe
Rolle, im selben Dokument.

**Wie es auffiel:** nicht im Repo, sondern beim **Konsumenten** (`hexslice-architecture`), der den
Umbau auf `adapters/driving|driven + ports/inbound|outbound` vollzogen hatte und beim Gegenlesen
des Handbuchs auf die alte Form stieß.

**Gemessen, mit zwei verschieden gebauten Zählern** (Mess-Regel 3): `grep` auf die alte Wendung
liefert **eine** Zeile; ein `awk` über Zeilen, die `dapter` **und** `inbound|outbound` führen,
liefert **sechs** — von denen **eine** fehlerhaft ist. Der zweite Zähler ist der weitere; die
fünf richtigen Zeilen sind die Gegenprobe, die den Befund auf die eine Stelle eingrenzt. Die beiden anderen Dokumente unter `docs/user/` (`benutzerhandbuch-standard.md`,
`releasing.md`) führen das Vokabular überhaupt nicht.

**Behoben** mit `internal/adapters/{driving,driven}/…`; die Sub-Area „Benutzer-Doku" bekam ihre
Zeile in der Modus-Deklaration (`harness/conventions.md`, Kürzel `USER`) — die Bedingung dafür,
dass dieser Beleg überhaupt einen Ort hat.
