# `make archive-wave` — Zeitdokumente archivieren (kein Gate)

## Vertrag

**Kein Gate** — das Target urteilt nicht über den Zustand des Repos, es
**bewegt**: Es sammelt die Slices einer Welle über ihr `**Welle:**`-Feld samt
ihrer Review-Reports, baut `done/<welle-id>/archiv.zip`, ersetzt die Volltexte
durch Stubs und zieht die Verweise nach. `SLICE=<id>`/`REVIEW=<datei>` für
wellenlose Einzelfälle.

Setzt Baseline-Regelwerk `modul-06-roadmap.md` §Wellen-Closure-Prozedur
Schritt 4 um. Seit slice-157 **Pflichtschritt beim Abschluss eines wellenlosen
Slice** (`AGENTS.md` §6) — kein von Hand ausgelöster Vorgang mehr.

Unverändert aus `d-check` übernommen; eigenständig, eigenes `go.mod`.

## Grenze — was das Grün nicht abdeckt

1. **Vollständigkeit des Archivs** — bezeugt nur der Archivierungs-Commit.
   Deshalb gehört die Operation in ein Werkzeug und nicht in Handarbeit;
   permanent.
2. **Die Zitier-Form im erzeugten Stub** — das `**Welle:**`-Feld des Stubs
   trägt einen Markdown-**Link**, den eine eigene Mechanik
   (`RewriteFieldForMove`) beim Umzug nachzieht. Die Zitier-Form
   (`AGENTS.md` §5) verlangt dort `slice-NNN` statt eines Lifecycle-Pfads.
   Heilbar; benannt in slice-176 §3, Umbau zieht `archive_test.go` nach.

## Sperren

- `WELLE=` nennt die **kurze numerische Form** (`welle-12`, nicht
  `welle-12-regelwerk-migration`) → das Werkzeug matcht nur die Ziffernfolge im
  `**Welle:**`-Feld, und a-checks Roadmap-IDs tragen einen beschreibenden
  Suffix, den es nicht kennt. Per Dry-Run gemessen, slice-144.
- ohne `APPLY=1` wird **nichts** geschrieben — sicherer Default.

## Bindung

**kein Gate** (kein `gates`/`ci`-Bestandteil), aber Pflichtschritt der
Slice-Closure seit slice-157 · Regelwerk `modul-06` §Wellen-Closure-Prozedur
Schritt 4 · Testsuite über `make archive-wave-test` (eigenes `go.mod`, **nicht**
Teil von `make test`).
