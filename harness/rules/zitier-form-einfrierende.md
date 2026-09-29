# Zitier-Form in einfrierenden Artefakten: Kennung statt Adresse

*Ausgelagert aus [`AGENTS.md`](../../AGENTS.md) §5, Regel 13 — seit slice-201. Der Wortlaut ist unverändert; die Links sind der neuen Tiefe nachgeführt.*

**Zitier-Form in einfrierenden Artefakten** (`v6.6.0`, vier Ziel-Formen:
  Review-Report, Welle-Ergebnisnotiz, beide Archiv-Stubs): Was einfriert,
  zitiert **Kennung statt Adresse** — `slice-NNN` statt seines Lifecycle-Pfads,
  `make <target>` statt eines Links auf die Sensor-Datei, eine Baseline-Stelle
  als **Tag + Pfad in Inline-Code** statt als Link
  (`` `v<X.Y.Z>` · `regelwerk/<datei>.md` §<Abschnitt> ``). Grund: Der vendored
  Baum trägt genau einen Tag, der nächste Sprung löscht den alten, und ein Link
  darauf färbt ein Artefakt rot, das niemand mehr anfassen darf. Verankert im
  Reviewer-Skill; für Archiv-Stubs erzeugt `tools/archive-wave/` den Text, für
  die Ergebnisnotiz gilt sie beim Schreiben.
  **Ebenso ein Planungs-Dokument, das einen anderen als den adoptierten Stand
  nennt** (`seit slice-192`): Ein Slice-Plan, der eine Migration vorbereitet oder
  einen gehenden Stand beschreibt, nennt ihn als **Kennung** — `v<X.Y.Z>` als
  Text, das Verzeichnis beschrieben statt als Pfad geschrieben. Ein Pfad dort
  wäre ein Pin auf einen Stand, der (noch) nicht adoptiert ist, und `versions`
  meldet ihn zu Recht.
  **Nicht** betroffen: lebende Dokumente, die den **adoptierten** Stand nennen —
  dort ist der Link richtig, und `versions` hält ihn aktuell.
