**Vorgang:** slice-170
**Fund:** Nach dem Abschluss von `welle-14` liegen `v6.0.0` und `v6.2.0` gleichzeitig vendored.
Gemessen über das Feld `Ersetzt-Baseline-Regel` **aller 20** `MR`-Dateien (aktive und aufgelöste):
genau **fünf** tragen dort einen `v6.0.0`-Anker —
[`MR-011`](../../../../../../../harness/conventions/MR-011-verfeinerungs-form.md),
[`MR-012`](../../../../../../../harness/conventions/MR-012-referenzmatrix-grandfathering.md),
[`MR-014`](../../../../../../../harness/conventions/MR-014-keine-agenten-telemetrie.md),
[`MR-015`](../../../../../../../harness/conventions/MR-015-welle-closure-ohne-replay.md),
[`MR-016`](../../../../../../../harness/conventions/MR-016-validator-unbesetzt.md) — **vor wie nach
diesem Slice dieselben fünf**. Daneben halten weitere Links denselben Stand fest: die Spalte
*Ersetzt-Baseline-Regel* der Index-Tabelle in
[`conventions.md`](../../../../../../../harness/conventions.md) sowie je ein Verweis in
[`MR-017`](../../../../../../../harness/conventions/done/MR-017-adr-vorlagen-version.md) und
[`MR-018`](../../../../../../../harness/conventions/done/MR-018-review-pflicht-v610-wortlaut.md).
Die Auflösung von [`MR-018`](../../../../../../../harness/conventions.md#mr-018) hat daran **nichts** geändert — sein Feld trägt seit `541581f` ein
„—", er war nie Teil der Menge.

Damit ist der Befund schärfer als zunächst notiert: es gab keinen Fortschritt, den man
fälschlich für einen halten könnte. `welle-14-results.md` hält für dieselbe Frage fest, es gebe
„kein absehbarer Trigger" — also auch keinen Wächter und keine Frist.
