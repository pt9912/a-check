# Slice-Lifecycle ist reine Datei-Bewegung — sechs Übergänge, gefahren mit `make slice-mv`

*Ausgelagert aus [`AGENTS.md`](../../AGENTS.md) §5, Regel 7 — seit slice-201. Der Wortlaut ist unverändert; die Links sind der neuen Tiefe nachgeführt.*

**Slice-Lifecycle** ist reine Datei-Bewegung — der Zustand ist das
  Verzeichnis, kein Feld im Dokument. Die **sechs** Übergänge und ihre Trigger
  stehen in `modul-05` §Trigger je Lifecycle-Übergang und WIP-Limit; a-check
  fährt sie mit **`make slice-mv`**, das den `git mv` samt der Verweise **auf**
  die Datei erledigt (§3.3).
  **Repo-eigen daneben:** Der direkte Weg `open/ → in-progress/` ist zulässig;
  `next/` ist ein Ort, keine Pflichtstation
  ([`next/README.md`](../../docs/plan/planning/next/README.md)).
