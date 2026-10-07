**Vorgang:** slice-213

**Fund (Review F-2, F-3, F-4):** `TestGomodUnsplittable` nannte sich „alle fail-closed-Fälle der
Spezifikation", ohne den Fall „Backtick mitten im Token" — ohne die Prüfung blieb der Test grün.
`TestEvaluateShapesMessagesOneLine` behauptete einzeilige `shape-differs`-Meldungen und prüfte nur
einen von drei Zweigen. Ein End-to-End-Kommentar sagte „je ein Befund" und prüfte mit `Contains`
nicht die Anzahl. Drei Funde, **ein** Vorgang. Behoben, die zwei Prüf-Lücken mit
Mutations-Gegenprobe (rot aus dem richtigen Grund).
