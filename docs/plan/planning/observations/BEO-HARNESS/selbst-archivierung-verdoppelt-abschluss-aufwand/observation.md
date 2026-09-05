# Selbst-Archivierung verdoppelt den Abschluss-Aufwand je wellenlosem Slice

**Sub-Area:** Harness-Einstieg

Seit der `AGENTS.md`-§6-Ergänzung archiviert sich jeder wellenlose Slice unmittelbar nach seiner
Closure selbst: zweiter `archive-wave`-Lauf, zweiter `make gates`/`make verify`-Durchgang, zweiter
Commit. Das verhindert den Backlog, den zwei vorherige manuelle Sweeps aufräumen mussten — aber es
verdoppelt auch den Abschluss-Aufwand jedes einzelnen wellenlosen Slice, unabhängig von seiner
Größe. Ob der Mehraufwand sich gegen den vermiedenen Sweep-Aufwand rechnet, zeigt sich erst über
mehrere künftige Closures, nicht an einer einzelnen.
