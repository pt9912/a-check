# Welle-ID: Kurzform vs. a-checks beschreibende Form

**Sub-Area:** Planungs-Harness

`tools/archive-wave` (aus `d-check` übernommen) erkennt in einem `**Welle:**`-Feld nur die
Ziffernfolge (`\bwelle-(\d+)\b`) und erwartet denselben Kurz-Namen als CLI-Argument
(`welle-12`). a-checks eigene Wellen-Kennungen in Roadmap und CHANGELOG tragen dagegen
durchgehend einen beschreibenden Suffix (`welle-12-regelwerk-migration`) — beide Formen sind
in diesem Repo etabliert und werden nicht ineinander aufgelöst.
