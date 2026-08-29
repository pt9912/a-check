# ADR-0035 — Die Grenz-Diagnose prüft gegen die Globs, nicht gegen die Schreibweise allein

- **Status:** Accepted
- **Datum:** 2026-08-29
- **Autor:** pt9912
- **Bezug:** [AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze), [AC-FA-CLI-001](../../../spec/lastenheft.md#ac-fa-cli-001--aufruf-scan-wurzel-und-exit-codes), [SPEC-CLI-001](../../../spec/spezifikation.md#spec-cli-001--aufruf-scan-wurzel-und-exit-codes), [ADR-0031](0031-heuristik-grenzen-diagnose.md), [ADR-0029](0029-abdeckungs-diagnose-advisory.md)
- **Schärft:** [SPEC-CLI-001](../../../spec/spezifikation.md#spec-cli-001--aufruf-scan-wurzel-und-exit-codes) — die stderr-Ausgabe eines Scans.
- **Supersedes:** [ADR-0031](0031-heuristik-grenzen-diagnose.md)

## Kontext

[ADR-0031](0031-heuristik-grenzen-diagnose.md) hat die Grenz-Diagnose auf zwei Formen festgelegt.
Ihre **Entscheidung 5** trägt für Klasse 2 eine Begründung, die nicht zutrifft:

> ein `../`-Pfad gegen einen nicht `relative`-Modus kann kein Ziel treffen, egal wie der Baum
> aussieht.

Das ist aus der Auflösung selbst widerlegbar, ohne einen Baum zu betrachten:

- Modus `path` gibt das Symbol **wörtlich** als Kandidat zurück, Punkte inklusive.
- Die Schicht-Zuordnung sucht das Glob-Präfix **segmentweise an beliebiger Stelle** des
  Kandidaten, nicht ab Position 0.

Für `../adapters/db/x.h` und `layers: {adapters: ["adapters/**"]}` trifft das Präfix hinter den
Punkten; die Kante **wird** beurteilt. Unter `fixed-root` gilt dasselbe mit vorangestellter Wurzel.

**Beobachtbare Folge:** dieselbe Zeile erscheint als Befund **und** als unbeurteilt.

```text
core/model.cpp:1: core-impurity: Kern importiert ../adapters/db/x.h
Hinweis: … core/model.cpp:1: relativer Pfad, den der Auflösungs-Modus "path" nicht auflöst
```

Relative Includes sind in C++ die Norm — der Fall trifft breit, nicht am Rand.

## Entscheidung

Die vier übrigen Entscheidungen von [ADR-0031](0031-heuristik-grenzen-diagnose.md) — advisory auf
stderr, Fundstelle mit Grund, kein Exit-Wechsel, stabile Sortierung — gelten unverändert weiter.
**Neu gefasst wird allein Entscheidung 5:**

1. **Klasse 2 wird gegen die konfigurierten Globs geprüft, nicht gegen die Schreibweise allein.**
   Gemeldet wird ein extrahiertes Symbol genau dann, wenn **kein** Layer-Glob-Präfix in einem
   seiner Auflösungs-Kandidaten segmentweise vorkommt. Ist eines dabei, ist die Kante beurteilbar
   und gehört nicht in die Diagnose.

2. **Die Diagnose bleibt tree-frei.** Sie liest Auflösungs-Modus **und** Globs — beides
   Konfiguration — und **nie** den Datei-Index. Damit bleibt die Grenze aus
   [ADR-0031](0031-heuristik-grenzen-diagnose.md) und
   [ADR-0029](0029-abdeckungs-diagnose-advisory.md) unangetastet: ein Symbol, das syntaktisch
   auflösbar wäre und im konkreten Baum nur kein Ziel findet, wird weiterhin **nicht** gemeldet —
   es ist von repo-externem Code nicht unterscheidbar.

3. **„Form-basiert" heißt ab jetzt: aus Symbol und Konfiguration entscheidbar** — nicht: aus der
   Zeichenkette allein. Das ist die eigentliche Korrektur. Die alte Fassung hat *Schreibweise* mit
   *Entscheidbarkeit ohne Baum* gleichgesetzt; nur das zweite trägt die Diagnose.

## Verglichene Alternativen

**Klasse 2 ersatzlos streichen.** Beseitigt den Fehler sicher und kostet die einzige Aussage, die
die Diagnose über *konfigurierte* Unauflösbarkeit machen kann. Ein Symbol, das unter dem gewählten
Modus wirklich nie ein Ziel treffen kann, bliebe dann unsichtbar — und genau das ist die Klasse,
für die ein Nutzer die Diagnose liest. Verworfen: der Fehler liegt in der Begründung, nicht im
Gegenstand.

**Den Datei-Index hinzunehmen** und melden, was real kein Ziel findet. Das wäre präzise, aber es
ist die Klasse, die [ADR-0029](0029-abdeckungs-diagnose-advisory.md) auf der Ziel-Seite
ausgeschlossen hat: von repo-externem Code nicht unterscheidbar. Eine sichere Aussage mit einer
unsicheren zu mischen kostet der Diagnose ihre Glaubwürdigkeit. Verworfen.

**Nur die Prosa korrigieren, den Code lassen.** Die Diagnose meldete weiter auflösende Zeilen; die
ADR beschriebe das Verhalten dann korrekt als falsch. Verworfen — eine Dokumentation, die einen
Defekt genau beschreibt, ist keine Entscheidung.

## Konsequenzen

- Die Zeile aus dem Kontext erscheint künftig als Befund und **nicht** mehr als Hinweis.
- Klasse 2 schrumpft auf die Fälle, in denen die Konfiguration die Auflösung wirklich
  ausschließt. Ob dabei Fälle übrig bleiben, ist eine Messung und keine Annahme — sie steht in der
  Fitness Function.
- Der Kommentar an `HeuristicLimits`, der die alte Begründung wiederholt, wird mit ersetzt. Eine
  Begründung, die an zwei Orten steht, driftet an einem davon.

## Fitness Function

Zwei Proben in beide Richtungen, beide in `internal/cli/cli_test.go`:

- Ein `../`-Symbol, dessen Ziel-Segment von einem Glob-Präfix getroffen wird, erzeugt einen
  **Befund** und **keinen** Hinweis.
- Ein `../`-Symbol, dessen Segmente von keinem Glob-Präfix getroffen werden, erzeugt weiterhin
  einen **Hinweis**.

Die zweite Probe ist die wichtigere: ohne sie wäre eine Klasse, die nie feuert, von einer
korrekten nicht zu unterscheiden.

## Re-Evaluierungs-Trigger

Ein Auflösungs-Modus, dessen Kandidaten nicht mehr aus Symbol und Konfiguration allein bestimmbar
sind — etwa eine Namespace-Auflösung, die den Baum braucht. Dann trägt Entscheidung 2 nicht mehr
und die Diagnose braucht eine eigene Antwort.

## Geschichte

| Datum | Änderung |
|---|---|
| 2026-08-29 | Proposed → Accepted (Sign-off Auftraggeber). Auslöser: `R-2` aus dem Review vom 2026-08-15, verifiziert am Code. Löst [ADR-0031](0031-heuristik-grenzen-diagnose.md) ab, deren Entscheidung 5 eine nicht zutreffende Begründung trägt; die übrigen vier Entscheidungen werden unverändert übernommen. Ab jetzt immutable; Ablösung nur via Folge-ADR mit `Supersedes`. |
