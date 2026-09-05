# Guard trifft Erwaehnung statt Aufruf

**Sub-Area:** Gate-/Werkzeug-Schicht
**Ehemals:** `BEO-027` (Tabellenform, bis slice-139)

Der PreToolUse-Command-Guard trifft Text, der die Toolchain **nennt**, nicht nur Kommandos, die sie **aufrufen**: ein Heredoc, das über die Toolchain schreibt, wird abgewiesen
