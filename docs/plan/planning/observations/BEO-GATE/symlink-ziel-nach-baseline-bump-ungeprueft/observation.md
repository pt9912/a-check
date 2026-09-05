# Symlink-Ziel nach Baseline-Bump ungeprüft

**Sub-Area:** Gate-/Werkzeug-Schicht

`.claude/rules/` trägt vier Symlinks auf die vendored Baseline (Claude-Code-nativ, bei jeder Sitzung
automatisch in den Kontext geladen). Beim Fall des alten Baseline-Baums blieben sie auf den alten
Pfad zeigen — kein Schritt der Migration prüfte sie, und kein Gate hätte den Bruch gefangen: die
Migration durchsuchte nur Dateiinhalt (`grep`), ein Symlink-Ziel liegt im Dateisystem, nicht im
Text. `d-check`s Linkpflicht deckt Markdown-Links, keine Dateisystem-Symlinks — die Lücke ist
strukturell, nicht durch ein präziseres Suchmuster behebbar.
