**Vorgang:** slice-192

**Fund:** Der Slice-Plan für die Migration auf den nächsten Baseline-Stand nannte das Zielverzeichnis
als **Pfad** — zweimal, in §1 und in der DoD. Das Modul `versions` meldete beide Stellen als
`version-stale`: Sie sind Pins auf einen Stand, den
[`harness/conventions.md`](../../../../../../../harness/conventions.md) §Baseline (noch) nicht
deklariert. Der Befund ist **richtig**; falsch war die Schreibweise.

**Drittes Vorkommen, und diesmal in der dritten Ausprägung:** slice-174 nannte den **kommenden**
Stand, slice-176 den **gehenden**, dieser Slice den kommenden in einem Plan, der die Migration
selbst ist. Der Titel des Eintrags nennt nur den Vorgriff; die Klasse ist der Bezug auf einen
anderen als den adoptierten Stand aus planerischer Absicht.

**Die Schwelle ist erreicht, und die Entscheidung stand schon im `state.md`:** Planungs-Dokumente
übernehmen die **Zitier-Form** — Kennung statt Adresse. Verkörpert in
[`AGENTS.md`](../../../../../../../AGENTS.md) §5, im vorhandenen Zitier-Form-Punkt, `seit slice-192`.
Der Umweg ist damit keine Umgehung mehr, sondern die Regel.
