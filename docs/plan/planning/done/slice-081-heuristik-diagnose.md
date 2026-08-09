# slice-081 — Laufzeit-Diagnose für nicht extrahierte Import-Formen

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Deckt:** Konsumenten-Befund vom 2026-08-09 (realer Einsatz in einem Fremd-Repo);
[AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze).
**Bezug:** Vorbild ist die Abdeckungs-Diagnose aus
[slice-043](../done/slice-043-schicht-abdeckung-sichtbar.md) /
[ADR-0029](../../adr/0029-abdeckungs-diagnose-advisory.md).

---

## 0. Trigger

**Beginn:** sofort. Der Befund stammt aus einem realen Einsatz und wartet auf nichts.

**Rückführungen:**

- `in-progress` → `open`: falls die Messung zeigt, dass „nicht extrahiert" ohne AST nicht
  zuverlässig von „keine Import-Zeile" zu trennen ist — dann wäre die Diagnose selbst eine
  Heuristik über einer Heuristik und braucht erst einen Entscheid.

## 1. Auslöser

**Mechanismus: das Werkzeug kennt seine Blindstellen, sagt sie aber nicht am Repo.**

[AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)
verspricht, die Heuristik-Grenzen **offenzulegen statt zu verschweigen**. Eingelöst wird das heute
in der Doku — der Out-of-Scope-Absatz von
[AC-FA-EXTRACT-001](../../../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion) listet
über zwanzig Grenzen. Am **Repo** sagt a-check davon nichts.

Die eine Hälfte existiert bereits: seit
[slice-043](../done/slice-043-schicht-abdeckung-sichtbar.md) meldet ein Scan auf stderr, welche
Dateien in **keinem** `layers`-Glob liegen. Es fehlt das Gegenstück:

> N Datei(en) enthalten Import-Schreibweisen, die dieses Backend nicht extrahiert.

**Belegt aus dem Einsatz** (Fremd-Repo, 2026-08-09): elternrelative C++-Includes, relative
Python-Importe und Mehrfach-Direktiven mussten je Sprach-Skelett von Hand nachgestellt werden, um
zu erkennen, wo das Werkzeug blind ist. Ein grünes Gate über teilweise nicht extrahiertem Code
sieht aus wie ein grünes Gate über geprüftem Code — dieselbe Klasse, die
[slice-043](../done/slice-043-schicht-abdeckung-sichtbar.md) für die Schicht-Seite geschlossen hat.

**Diese Diagnose subsumiert den zweiten Konsumenten-Befund** (Mehrfach-Direktiven, siehe
[slice-084](../done/slice-084-handbuch-heuristik-grenzen.md)): statt jede Grenze einzeln im Handbuch zu
suchen, meldet das Werkzeug, was es in **diesem** Baum nicht gegriffen hat.

## 2. Betroffene Module

Zwei Schichten:

1. **Spec** — [`spec/lastenheft.md`](../../../../spec/lastenheft.md) (neue `AC-*` oder Erweiterung
   von [AC-FA-CLI-001](../../../../spec/lastenheft.md#ac-fa-cli-001--aufruf-scan-wurzel-und-exit-codes)),
   [`spec/spezifikation.md`](../../../../spec/spezifikation.md), ADR nach dem Muster von
   [ADR-0029](../../adr/0029-abdeckungs-diagnose-advisory.md).
2. **`internal/adapter/driven/extract/`** + [`internal/cli/`](../../../../internal/cli/cli.go) —
   die Zählung und ihre Ausgabe.

## 3. Auszuführende Gates

`make gates`, `make image-test` (die Diagnose ist Teil des CLI-Vertrags).

**Der Entscheid, der vor dem Bau fällt: was zählt als „nicht extrahiert"?** Der Slice nimmt ihn
nicht vorweg. Drei Kandidaten, die sich in Präzision und Rauschen unterscheiden:

| Kandidat | Aussage | Risiko |
|---|---|---|
| Zeilen, die ein **Import-Schlüsselwort** tragen, aber von keinem Regex gegriffen werden | präzise am Fund | Kommentare/Strings mit `import` erzeugen Rauschen |
| Dateien **ohne einen einzigen** extrahierten Import | billig | eine Datei ohne Importe ist normal, kein Befund |
| Nur die **benannten** Grenzen (relative Importe, Mehrfach-Direktiven) gezielt erkennen | rauschfrei | wächst mit jeder neuen Grenze mit, also Wartungsposten |

**Advisory, nicht gatend** — wie
[ADR-0029](../../adr/0029-abdeckungs-diagnose-advisory.md) es für die Abdeckungs-Diagnose
entschieden hat: der Exit-Code bleibt unberührt, vollständige Extraktion erzeugt **keine** Ausgabe.
Eine Diagnose, die bei jedem Lauf spricht, wird weggeschaltet.

**Nachtrag 2026-08-09 — `CR-1` präzisiert den Vertrag.** Der Konsument hat den Befund als formalen
Change Request nachgereicht, mit drei gemessenen Fällen (je **0 Befunde** bei vorhandenem Verstoß)
und einem Umsetzungs-Vorschlag:

| Sprache | Schreibweise | Ergebnis |
|---|---|---|
| C++ | `#include "../../adapters/ui/x.h"` | löst gegen `resolution: fixed-root` auf nichts auf |
| Python | `from ..ui import x` | dokumentierte Grenze, aber ohne Signal |
| Python | `from docsearch import ui` | Subpaket-Form |

Erkennung je Backend über ein **zweites, bewusst breiteres Muster** — Python `^\s*from\s+\.`,
C++ `#include\s*"\.\.`; für TypeScript nicht nötig, weil `relative` dort ein regulärer
Auflösungs-Modus ist. Das ist der vierte Kandidat neben den drei in der Tabelle oben und der
konkreteste; er gehört in den Entscheid.

**Zwei Akzeptanzkriterien aus dem CR, die der Slice übernimmt:**

1. **Ohne die neuen Muster byte-identische Befundmenge** — die Diagnose darf die Regel-Auswertung
   nicht anfassen.
2. **Der Hinweis erscheint auch bei `gesamt: 0 Befund(e)`** — *dort ist er am wichtigsten*. Ein
   Hinweis, der nur bei ohnehin roten Läufen erscheint, verfehlt den Zweck.

**Negativ-Proben:**

| Probe | Erwartung |
|---|---|
| Python-Datei mit `from . import x` | gemeldet |
| Python-Datei mit `import a, b` | gemeldet |
| C++-Datei mit `#include "../../x.h"` | gemeldet |
| Baum ohne solche Formen | **keine Ausgabe**, Exit unverändert |
| Baum **mit** solchen Formen und sonst 0 Befunden | Hinweis erscheint, Exit bleibt **0** |

## 4. Was bewusst nicht getan wird

- **Die Grenzen selbst schließen.** Ein AST-Backend ist in
  [AC-FA-EXTRACT-001](../../../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion)
  ausdrücklich Out-of-Scope. Dieser Slice macht die Grenze *sichtbar*, er verschiebt sie nicht.
- **Gatend machen.** Ein `strict_extraction` analog `strict_coverage` ist wie dort vertagt — die
  advisory Form reicht für den belegten Bedarf.
- **Die zwei `--print-mk`-Defekte.** Eigene Slices
  ([slice-082](../done/slice-082-print-mk-docker-indirektion.md),
  [slice-083](../done/slice-083-print-mk-digest-selbstbezug.md)).

## 5. DoD

- [x] Spec-first: die Diagnose steht als Vertrag, geschärft durch eine ADR, bevor Code entsteht.
      Beleg: [ADR-0031](../../adr/0031-heuristik-grenzen-diagnose.md) mit `Status: Accepted`,
      [`spec/spezifikation.md`](../../../../spec/spezifikation.md) 0.28.0
      ([SPEC-CLI-001](../../../../spec/spezifikation.md#spec-cli-001--aufruf-scan-wurzel-und-exit-codes)).
      **Bewusste Abweichung vom Wortlaut: kein Lastenheft-Bump.** Die bestehende
      Abdeckungs-Diagnose steht ebenfalls nicht im Lastenheft — sie kam über
      [ADR-0029](../../adr/0029-abdeckungs-diagnose-advisory.md) in die Spezifikation, weil eine
      advisory stderr-Zeile ohne Exit-Code-Wechsel das *Wie* der bestehenden Ausgabe schärft und
      keinen neuen Vertrag aufmacht. Die Zusage selbst steht längst da:
      [AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)
      verlangt, die Grenzen offenzulegen. Diesem Präzedenzfall zu folgen war richtiger, als für
      eine Ausgabe-Schärfung eine `AC-*`-ID zu erfinden.
- [x] Ein Scan meldet die Formen auf stderr, ohne den Exit-Code zu ändern; ein sauberer Baum
      bleibt still. Beleg: die Proben aus §3, gefahren gegen `a-check:dev` (Closure-Notiz).
- [x] `make gates` und `make image-test` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft.

## 6. Closure-Notiz

**Der Entscheid aus §3 fiel auf Kandidat 4** — gezielte Gegenmuster je Backend, festgehalten in
[ADR-0031](../../adr/0031-heuristik-grenzen-diagnose.md). Ausschlaggebend war nicht Eleganz,
sondern ein gemessenes Rauschprofil: der Restmengen-Ansatz meldet Kommentare und String-Literale,
und genau diese Klasse hat das Repo dreimal getroffen
([`SL-004`](../../steering-loop.md)). Eine Diagnose, die bei jedem Lauf spricht, wird
weggeschaltet — dann ist sie schlechter als keine.

**Beobachtbare Architektur-Aussage: die Config-Kenntnis blieb im Kern.** Beim Bau zerfiel „nicht
extrahiert" in zwei Klassen, die verschiedene Heimaten haben. Klasse 1 (die Zeile greift kein
Muster) ist rein syntaktisch — der Extraktions-Adapter sieht sie, während er die Quelle liest.
Klasse 2 (`../`-Pfad unter einem Modus ≠ `relative`) braucht den `resolution`-Modus, also Config;
sie wird in `core.HeuristicLimits` aus den bereits extrahierten Symbolen abgeleitet. Der Adapter
bleibt damit frei von Config-Wissen, und die eine Stelle, die über Auflösung urteilt, bleibt der
Kern. Die naheliegende Variante — dem Extraktor das Model durchreichen — hätte in einer Stunde
funktioniert und die Schichtung dauerhaft verwischt.

**Lernsignal mit Ursache: [ADR-0030](../../adr/0030-kein-digest-im-generierten-fragment.md) fehlte
im ADR-Index.** slice-083 hat die ADR angelegt und die Index-Zeile in
[`docs/plan/adr/README.md`](../../adr/README.md) nicht nachgezogen; `make gates` blieb grün.
Ursache: `doc-check` prüft, dass jeder Link **auflöst**, nicht dass jede ADR-Datei **verlinkt ist**
— die Gegenrichtung Baum → Index fehlt. Das ist exakt die Klasse, die
[slice-071](../done/slice-071-sensor-scope-vollstaendig.md) für `regelwerk-check` geschlossen hat
(dort per `comm -13`), hier nur an einem anderen Index. Aufgefallen ist es beim Nachtragen von
[ADR-0031](../../adr/0031-heuristik-grenzen-diagnose.md), also durch Zufall, nicht durch einen
Sensor. **Folge-Slice:** ein Vollständigkeits-Sensor für den ADR-Index, in derselben Gestalt wie
die Regelwerk-Gegenrichtung. Der Nachtrag selbst ist als eigener Commit mit slice-083-Bezug
abgelegt.

**Die Proben aus §3, gegen `a-check:dev`:**

Konsumenten-Fixture (relativer Python-Import, `import os, sys`, elternrelativer C++-Include) —
der Fall aus `CR-1`, bei dem der Konsument nur `gesamt: 0 Befund(e)` sah:

```text
gesamt: 0 Befund(e)
Hinweis: 3 Import-Zeile(n) unterliegen einer Heuristik-Grenze und bleiben unbeurteilt:
  app/core/service.py:1: relativer Import — von diesem Backend nicht extrahiert
  app/ui/widget.py:1: zweite Direktive auf derselben Zeile — nur die erste wird extrahiert
  src/core/service.h:1: relativer Pfad, den der Auflösungs-Modus "path" nicht auflöst
  Abhilfe: Schreibweise ändern, resolution-Modus anpassen oder die Grenze bewusst hinnehmen.
```

Exit blieb **0**. Derselbe Baum mit einzeiligen Importen und einem wurzel-relativen Include:
`gesamt: 0 Befund(e)`, **keine** Diagnose. Ein echter Verstoß dort: `wrong-direction`, Exit 1,
ebenfalls **keine** Diagnose — die Zeile löst ja auf. `make arch-check` auf dem Eigen-Baum bleibt
diagnose-frei.

**Ein zweiter Befund fiel beim Fixture-Bauen ab, und er widerlegt nichts an dieser Diagnose,
sondern begrenzt sie:** `#include "adapters/ui/panel.h"` aus `src/core/` bleibt still, obwohl es
ein Verstoß ist — das Symbol trägt das `src/`-Präfix nicht und löst auf keine Schicht auf. Diese
Klasse meldet die Diagnose bewusst **nicht** (von repo-externem Code nicht unterscheidbar,
[ADR-0031](../../adr/0031-heuristik-grenzen-diagnose.md) Entscheidung 5). Der Bedarf für die
Ziel-Seiten-Diagnose aus [slice-085](../done/slice-085-schicht-ohne-aufloesung.md) (`CR-2`) ist
damit an einem laufenden Beispiel belegt, nicht nur behauptet.

**Lerneintrag — Form: neuer Sensor.** Als Prüfsatz: *Ein Index, der von Hand gepflegt wird, braucht
einen Sensor in der Richtung Baum → Index; die Richtung Index → Baum fällt beim Verlinken ohnehin
auf.* `doc-check` deckt die eine Richtung vollständig ab — jeder Link muss auflösen —, und genau
deshalb sah der ADR-Index vollständig aus, obwohl ihm ein Eintrag fehlte: was nicht verlinkt ist,
kann auch nicht ins Leere zeigen. Dasselbe Muster hat
[slice-071](../done/slice-071-sensor-scope-vollstaendig.md) für das Regelwerk-Manifest gefunden;
dass es hier ein zweites Mal auftritt, macht es zur Klasse und nicht zum Einzelfall. Zu prüfen sind
alle handgepflegten Indizes des Repos, nicht nur der ADR-Index.

## 7. Sub-Area-Modus

Alle berührten Sub-Areas GF.
