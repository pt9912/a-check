# ADR-0031 — Heuristik-Grenzen-Diagnose: benannte Formen erkennen, nicht Restmengen melden

- **Status:** Accepted
- **Datum:** 2026-08-09
- **Autor:** pt9912
- **Bezug:** [AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze), [AC-FA-EXTRACT-001](../../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion), [AC-FA-CLI-001](../../../spec/lastenheft.md#ac-fa-cli-001--aufruf-scan-wurzel-und-exit-codes), [SPEC-CLI-001](../../../spec/spezifikation.md#spec-cli-001--aufruf-scan-wurzel-und-exit-codes), [SPEC-EXTRACT-001](../../../spec/spezifikation.md#spec-extract-001--import-extraktion), [SPEC-DET-001](../../../spec/spezifikation.md#spec-det-001--determinismus-vertrag), [ADR-0029](0029-abdeckungs-diagnose-advisory.md)
- **Schärft:** [SPEC-CLI-001](../../../spec/spezifikation.md#spec-cli-001--aufruf-scan-wurzel-und-exit-codes) — die stderr-Ausgabe eines Scans.
- **Supersedes:** —

## Kontext

[AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze) verspricht,
die Heuristik-Grenzen **offenzulegen statt zu verschweigen**. Eingelöst wird das heute nur in der
Doku: der Out-of-Scope-Absatz von
[AC-FA-EXTRACT-001](../../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion)
listet die nicht extrahierten Formen auf. **Am geprüften Baum sagt a-check davon nichts.**

[ADR-0029](0029-abdeckungs-diagnose-advisory.md) hat die eine Hälfte der Lücke geschlossen: gescannte
Dateien, die in **keinem** `layers`-Glob liegen, werden ausgewiesen. Die andere Hälfte bleibt still.
Eine Datei kann korrekt in einer Schicht liegen, gescannt werden **und** Import-Zeilen tragen, die
das Backend gar nicht erst extrahiert. Für diese Zeilen entsteht keine Kante, also kein Befund — und
kein Hinweis.

Belegt aus einem realen Konsumenten-Einsatz (2026-08-09), je **0 Befunde** bei tatsächlich
vorhandenem Schichtverstoß:

| Sprache | Schreibweise | Ursache |
|---|---|---|
| C++ | `#include "../../adapters/ui/x.h"` | elternrelativ; das Muster greift, der Pfad löst gegen die Wurzel auf nichts auf |
| Python | `from ..ui import x` | relative Importe werden nicht extrahiert |
| Python | `import a, b` | nur die erste Direktive einer Zeile |
| C# | `using A; using B;` | nur die erste Direktive einer Zeile |

Der Konsument musste je Sprach-Skelett von Hand einen Verstoß in der umgehenden Schreibweise
einbauen, um zu erfahren, wo das Gate blind ist. Genau diesen Handgriff soll das Werkzeug abnehmen.

Der Unterschied zu [ADR-0029](0029-abdeckungs-diagnose-advisory.md) ist der Wirkort: dort ist die
**Datei** ungedeckt und der Konsument sieht sie in seiner Config fehlen; hier ist die **Zeile**
unsichtbar, und in der Config ist nichts falsch. Ohne Meldung ist diese Klasse von außen nicht
auffindbar.

## Entscheidung

**Die Diagnose erkennt gezielt die benannten Grenzen — je Sprach-Backend über ein zweites, bewusst
breiteres Gegenmuster** — statt über eine Restmenge („jede Zeile, die kein Extraktions-Regex
griff").

1. **Advisory auf stderr**, nach der Zusammenfassung, in derselben Gestalt wie
   [ADR-0029](0029-abdeckungs-diagnose-advisory.md): **der Exit-Code bleibt unberührt**. Ein Baum
   ohne solche Zeilen erzeugt **keine** Ausgabe.
2. **Gemeldet wird die Fundstelle mit Grund**, nicht nur eine Zahl: `pfad:zeile` plus die Form, an
   der es lag. Eine Diagnose, die den Konsumenten suchen lässt, spart ihm nichts.
3. **Kein Befund, keine Regel, kein Exit-1** — auch nicht opt-in. Diese Zeilen sind nicht
   *verboten*, sie sind *unbeurteilt*. Aus „ungeprüft" ein „verletzt" zu machen, wäre eine andere
   Aussage, als das Werkzeug treffen kann.
4. **Stabil sortiert** ([SPEC-DET-001](../../../spec/spezifikation.md#spec-det-001--determinismus-vertrag))
   und mit gekappter Liste, die ihre Restzahl nennt — dieselbe Regel wie bei der Abdeckungs-Diagnose;
   zwei Kappungs-Semantiken nebeneinander wären eine Falle.
5. **Form-basiert, nie inhalts-basiert.** Gemeldet wird (a) eine Zeile, die **kein**
   Extraktions-Muster greift, und (b) ein extrahiertes Symbol, das der konfigurierte
   Auflösungs-Modus **per Konstruktion** nicht auflösen kann — ein `../`-Pfad gegen einen nicht
   `relative`-Modus kann kein Ziel treffen, egal wie der Baum aussieht. Beides steht in der
   **Schreibweise** selbst.

   **Nicht** gemeldet wird ein Symbol, das syntaktisch auflösbar wäre und im konkreten Baum nur
   kein Ziel findet. Diese Klasse ist von repo-**externem** Code nicht unterscheidbar — dieselbe
   Grenze, an der [ADR-0029](0029-abdeckungs-diagnose-advisory.md) (Entscheidung 3) die Ziel-Seite
   ausgeschlossen hat. Sie gehört in eine eigene Diagnose mit eigener Evidenz; eine gemeinsame
   Meldung würde eine sichere Aussage mit einer unsicheren mischen.

## Konsequenzen

- **Kein Vertragsschnitt nach oben.** Kein Lastenheft-Bump, keine neue `AC-*`-ID — eine advisory
  stderr-Zeile ohne Exit-Code-Wechsel schärft das *Wie* der bestehenden Ausgabe. Präzedenz:
  [ADR-0029](0029-abdeckungs-diagnose-advisory.md), [ADR-0025](0025-exclude-verzeichnis-prune.md),
  [ADR-0028](0028-ziel-glob-schattenwurf.md).
- **Die Musterliste ist ein Wartungsposten, ausdrücklich.** Jede künftig entdeckte Grenze muss
  bewusst aufgenommen werden, sonst bleibt sie unsichtbar. Das ist der Preis für Rauschfreiheit —
  und ein *ehrlicher* Posten: er zwingt dazu, jede Grenze zu benennen, statt sie in einer Restmenge
  verschwinden zu lassen. Der Out-of-Scope-Absatz von
  [AC-FA-EXTRACT-001](../../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion)
  ist die Liste, an der sich die Muster messen.
- **Deckung ist Teilmenge, nicht Vollständigkeit.** Gemeldet wird, was ein Gegenmuster trifft; für
  alles andere bleibt die Doku die einzige Quelle. Die Diagnose sagt „hier ist eine Grenze", nie
  „hier sind alle".
- **Byte-Identität bleibt.** Deterministisch sortiert; der nativ↔Container-Vergleich der
  Distributions-Akzeptanz (stdout **und** stderr) hält.
- **Zwei Diagnosen auf stderr.** Ein Baum kann beide auslösen. Sie stehen in fester Reihenfolge
  (Abdeckung zuerst, sie ist die gröbere Aussage) und sind je selbsttragend formuliert.

## Verworfene Alternativen

- **Restmengen-Ansatz — jede Zeile melden, die ein Import-Schlüsselwort trägt, aber von keinem
  Extraktions-Regex gegriffen wurde.** Präzise am Fund, rauschanfällig im Umfeld: Kommentare und
  Zeichenketten-Literale mit `import` erzeugen Treffer, ebenso jede Sprach-Syntax, die das Wort
  anders nutzt. Genau diese Klasse hat das Repo schon dreimal getroffen —
  [`SL-004`](../steering-loop.md) hält fest, dass ein neuer Sensor im ersten Lauf sein eigenes
  Umfeld meldete. Eine Diagnose, die bei jedem Lauf spricht, wird weggeschaltet; dann ist sie
  schlechter als keine.
- **Dateien ohne einen einzigen extrahierten Import melden.** Billig und falsch: eine Datei ohne
  Importe ist der Normalfall, kein Befund.
- **Die Grenzen schließen statt melden.** Ein AST-Backend ist in
  [AC-FA-EXTRACT-001](../../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion)
  ausdrücklich Out-of-Scope; die text-heuristische Extraktion ist eine Fundament-Entscheidung, keine
  Bequemlichkeit. Einzelne Grenzen *können* später fallen (Mehrfach-Direktiven wären machbar) — jede
  ist dann eine eigene Verhaltensänderung mit Vertragsbezug. Das ändert nichts daran, dass **jede**
  verbleibende Grenze bis dahin unsichtbar bleibt; die Diagnose ist zu diesen Schritten komplementär,
  nicht ihr Ersatz.
- **Eigenes Flag (`--print-limits`).** Verworfen aus demselben Grund wie in
  [ADR-0029](0029-abdeckungs-diagnose-advisory.md): eine Diagnose, die man anfordern muss, sieht
  genau der nicht, der sie braucht.

## Fitness Function

- `make test`: die belegten Formen beider Klassen — (a) relativer Python-Import, Mehrfach-Direktive;
  (b) elternrelativer C++-Include gegen einen nicht `relative`-Modus — erzeugen je eine Meldung mit
  `pfad:zeile`; derselbe Include **unter** `relative` bleibt **still**, denn dort löst er auf;
  ein Baum ohne solche Formen bleibt **still**;
  der Exit-Code ist unverändert — auch bei **0** Befunden erscheint die Meldung; zwei Läufe
  byte-identisch.
- **Regressions-Probe:** ohne die neuen Muster ist die **Befund**-Menge byte-identisch zu vorher —
  die Diagnose ändert nichts an dem, was gegatet wird.
- `make arch-check` (Dogfooding): unverändert **0** und diagnose-frei.
- `make ci` grün.

## Geschichte

| Datum | Ereignis |
|---|---|
| 2026-08-09 | Proposed → Accepted (Sign-off Auftraggeber). Auslöser: `CR-1` aus einem realen Konsumenten-Einsatz mit vier gemessenen Fällen und dem Vorschlag der Gegenmuster; die Gestalt „gezielte Muster" wurde gegen den Restmengen-Ansatz gewählt, weil dessen Rauschprofil im Repo dreimal belegt ist ([`SL-004`](../steering-loop.md)). Ab jetzt immutable; Ablösung nur via Folge-ADR mit `Supersedes`. |
