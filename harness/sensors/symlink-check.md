# `make symlink-check` — jeder getrackte Symlink löst auf und zeigt auf den adoptierten Stand

## Vertrag

**Zwei** Prüfungen an jedem von git getrackten Symlink:

1. **Das Ziel existiert.** Ein Symlink ins Leere ist Exit 1.
2. **Ein Ziel unter `.harness/baseline/` trägt den adoptierten Stand.** Der
   Stand wird aus `harness/conventions.md` §Baseline gelesen — fail-closed:
   ohne lesbaren Stand bricht der Lauf ab, statt zu raten.

Beide Prüfungen laufen durch **eine** Funktion, die auch der Selbsttest
aufruft; ein zweiter Codepfad für den Test hieße, dass der Test nicht prüft,
was im Gate läuft.

## Grenze — was das Grün nicht abdeckt

1. **Ob ein Symlink außerhalb der Baseline das inhaltlich richtige Ziel hat** —
   das wäre ein Urteil über Absicht, kein Match. Permanent.
2. **Untrackte Symlinks** — sie sind lokaler Kram und in keinem frischen Klon
   vorhanden. Permanent, und beabsichtigt.

**Warum es zwei Prüfungen sind:** Die drei Vorfälle, die den Sensor ausgelöst
haben, haben zwei Formen. Bei slice-173 zeigte ein Symlink auf einen
*entfernten* Stand; bei slice-167 auf einen *noch vorhandenen alten*, während
zwei Stände nebeneinanderlagen — er löste auf und war trotzdem falsch. Eine
Prüfung allein hätte die zweite Instanz grün gemeldet und den
Beobachtungs-Eintrag fälschlich als verkörpert ausgewiesen.

**Wie groß der Ausschnitt ist, sagt das Kommando:** die Schluss-Zeile nennt die
Zahl der getrackten Symlinks.

## Sperren

- `adoptierter Baseline-Stand nicht lesbar` — `harness/conventions.md` §Baseline
  trägt keine Zeile der Form `- **Stand:** [`vX.Y.Z`](…)` → die Zeile
  wiederherstellen; der Lauf rät nicht.

## Bindung

Harness-Prozess; Antwort auf `BEO-GATE/symlink-ziel-nach-baseline-bump-ungeprueft`
bei 3× · slice-173.
