**Vorgang:** slice-170
**Fund:** `v6.2.0` · `regelwerk/grundlagen-harness-dateien.md`
verlangt, dass ein Verweis aus `AGENTS.md`, einem Slice oder einer ADR auf eine Adaption die
**Index-Form** `harness/conventions.md#mr-<NNN>` nimmt — ausdrücklich **nicht** den Pfad auf die
Eintrags-Datei, weil ein Pfad-Link „genau in dem Moment bricht, in dem die Adaption sich auflöst".

Genau das trat in diesem Slice ein: acht Dateien mussten nachgezogen werden, weil [`MR-018`](../../../../../../../harness/conventions.md#mr-018) nach
`conventions/done/` wanderte. Die Regel wurde beim Nachziehen **nicht erwogen** — der Bestand
führte 129 Pfad-Links in 19 Dateien, und der Vorgänger-Slice `slice-166` hatte für [`MR-017`](../../../../../../../harness/conventions.md#mr-017)
denselben Nachzug schon einmal so gefahren. Aufgefallen ist sie erst dem unabhängigen Review,
das die Baseline gelesen hat statt den Bestand.
