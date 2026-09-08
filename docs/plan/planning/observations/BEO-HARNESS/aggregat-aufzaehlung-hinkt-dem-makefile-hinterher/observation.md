# Die Aggregat-Aufzählung in `AGENTS.md` §4 hinkt dem Makefile hinterher

**Sub-Area:** Harness-Einstieg

`AGENTS.md` §4 zählt **abschließend** auf, welche `doc-*`-Targets in `gates` oder `verify` hängen,
und erklärt „die übrigen" zu advisory. Die Liste ist eine **zweite Quelle** für eine Menge, die im
`Makefile` steht — und sie altert, sobald ein Target dazukommt oder das Aggregat wechselt.

**Kein Sensor deckt sie.** `make doc-targets` prüft die zwei Gate-**Tabellen** gegen das Makefile
(existiert das Target, ist es deklariert), nicht diesen Prosa-Absatz. `make gate-consistency`
prüft `.PHONY`, Pins und den ADR-Index. Der Absatz steht zwischen beiden Zuständigkeiten.

Die Klasse ist die der **abschließenden Aufzählung neben einer maschinenlesbaren Quelle**: Sie ist
beim Schreiben richtig und wird ohne Vorwarnung falsch.
