# Ein neuer Sensor prüft nur, was die DoD-Phrase explizit trägt — ohne sie bleibt er stumm

**Sub-Area:** Harness-Tooling

Ein Gate, das opt-in pro Slice über eine exakte DoD-Phrase wirkt (hier: `make doc-reviews`,
Modul `reviews`, ausgelöst durch die Phrase „unabhängiger Review" in einer DoD-Zeile), ist ab dem
Moment seiner Aktivierung technisch scharf, aber praktisch wirkungslos, solange kein Slice diese
Phrase führt. Der Unterschied zur reinen Prosa-Regel aus
[slice-159](../../../done/wellenlos/slice-159-reviewer-rolle-in-agents-md-verankert.md) ist real
(ein zukünftiger Verstoß wird gefunden, sobald die Phrase auftaucht), aber die Lücke bis dahin ist
strukturell dieselbe: eine Zusage, die niemand macht, wird nie geprüft.
