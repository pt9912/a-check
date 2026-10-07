**Vorgang:** slice-220

**Fund (Review F-1, F-2, F-3; Delta D-1):** Der macOS-Abschnitt sagte für beide Laufzeiten, ein
nicht freigegebener Pfad erscheine im Container leer — die Docker-Doku sagt für Docker Desktop
„Mounts denied", direkt nach dem zitierten Satz. `colima start --mount` stand ohne den Hinweis,
dass es bei laufender VM nichts ändert und die Mount-Liste `$HOME` ersetzt. Die Quellenliste
deckte zwei von sieben Aussagen; ein Satz zum Konfigurationsort folgte einer ungenauen
Review-Formulierung statt dem Quelltext. Vier Funde, **ein** Vorgang. Behoben gegen die Quellen.
