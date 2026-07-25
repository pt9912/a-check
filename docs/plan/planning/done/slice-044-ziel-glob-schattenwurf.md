# slice-044 — Unauflösbares Ziel-Glob: Zuordnung zurückziehen statt zurückfallen

**Status:** in-progress (Ergebnis §7, Abschluss mit dem Merge) — Umsetzung von **Option A′** aus
[slice-037 §4.0a](../open/slice-037-hexslice-gap-analyse.md); Variante **2** („kleiner Fix statt
Feature") am 2026-07-25 per Maintainer-Wort abgenommen.
**Auslöser:** die Nachmessung in [slice-037 §4.0a](../open/slice-037-hexslice-gap-analyse.md) —
ein legitimer Adapter→Port-Import meldet `wrong-direction`, weil das Ziel auf die *umschließende*
Schicht zurückfällt.
**Bezug:** schärft die Ziel-Auflösung aus
[SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung); betrifft
[AC-FA-RULE-005](../../../../spec/lastenheft.md#ac-fa-rule-005--schicht-richtung-regel-wrong-direction)
(`wrong-direction`) und bewegt sich an der ausgewiesenen Grenze
[AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze).
[Roadmap](roadmap.md).

Lifecycle-Notiz: der Slice startet direkt in `in-progress/` — Analyse **und** Abnahme liegen
bereits in [slice-037](../open/slice-037-hexslice-gap-analyse.md); ein Entwurfs-Durchlauf in
`open/` wäre eine leere Station.

---

## 1. Das Problem in einem Satz

Ein Layer-Glob mit **Wildcard in der Mitte** (`…/application/**/ports/**`) kann als Import-**Ziel**
nicht auflösen — und statt „unbekannt" liefert die Auflösung die **umschließende** Schicht
(`application`). Das ist kein Loch, sondern ein **Fehlbefund**.

## 2. Messung (slice-037 §4.0a, reproduziert)

Fixture: Schichten `ports` (Glob `src/hexagon/application/**/ports/**`, **vor** `application`
deklariert), `application`, `domain`, `outbound`, `inbound`; Kanten u. a. `outbound → ports`,
`inbound → application`, **nicht** `outbound → application`.

| Achse | Mechanik | vorher |
|---|---|---|
| Quell-Seite (Datei→Schicht, `LayerOf`) | literaler Präfix + Deklarationsreihenfolge | ✅ klassifiziert korrekt als `port` |
| Ziel-Seite (Import→Schicht, `layerOfCand`) | roher `globPrefix` + `segIndex` | ❌ fällt auf `application` zurück ⇒ `wrong-direction: outbound -> application` |

**Warum das teuer ist:** die naheliegende Nutzer-Reparatur — die Kante `outbound → application`
deklarieren — macht das Gate grün und **verdeckt dauerhaft echte** `outbound → application`-Verstöße.
Ein Falsch-Positiv, dessen Behebung ein Falsch-Negativ erzeugt.

## 3. Entscheidung (Variante 2)

**Nicht** die Auflösung nachrüsten (Variante 1: `layerOfCand` wildcard-fähig machen — Eingriff in
den Kern-Resolver mit dem Regressions-Risiko, das die 27 Falsch-Positive in
[slice-039](../done/slice-039-hexslice-vertical-slice-regeln.md) gezeigt haben), sondern die
**Zuordnung zurückziehen**: wo ein unauflösbares Glob den Kandidaten genauso gut decken könnte,
ist die Schicht **nicht diskriminierbar** ⇒ **extern** (fail-open), niemals die umschließende
Schicht. Dieselbe Linie, die
[ADR-0023](../../adr/0023-deklarations-bewusste-mehr-wurzel-aufloesung.md) für die schwache
Evidenzstufe zieht: nicht raten.

Der Rückzug ist **eng** gefasst (drei Bedingungen, sonst bleibt die Zuordnung stehen) — siehe
[ADR-0028](../../adr/0028-ziel-glob-schattenwurf.md). Entscheidend ist die **Tail-Marker**-Bedingung:
ohne sie verschlänge `…/application/**/ports/**` **jeden** Application-Import und schaltete die
App-Schicht still ab — ein weit schlimmeres False-Green als der Ausgangs-Fehlbefund.

## 4. Betroffene Module

| Modul | Änderung |
|---|---|
| `internal/hexagon/core` | `layerOfCand`-Rückzug (`shadowedByWildcardGlob`), Helfer `literalHead`/`literalTail` (der erste aus `litPrefixLen` extrahiert) |
| Doku | [ADR-0028](../../adr/0028-ziel-glob-schattenwurf.md), [SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung), Benutzerhandbuch, [CHANGELOG](../../../../CHANGELOG.md), Roadmap |

**Kein Lastenheft-Bump:** das *Was* (ein nicht auflösbares Import-Ziel ist extern,
[AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)) steht
schon; geschärft wird das *Wie*. Präzedenz: die `exclude`-Prune-Schärfung
([ADR-0025](../../adr/0025-exclude-verzeichnis-prune.md)) lief genauso.

## 5. DoD

- [x] [ADR-0028](../../adr/0028-ziel-glob-schattenwurf.md) `Accepted` + Index; Spezifikation nachgezogen (Minor-Bump).
- [x] Code + Tests: Fehlbefund weg **und** vier Gegenproben grün (Nicht-Port-Ziel meldet weiter, sauberes Glob löst weiter auf, spezifischeres Literal-Glob gewinnt weiter, Glob ohne Tail-Marker unverändert).
- [x] **Regressions-Probe über alle lokalen Konsumenten** — keiner nutzt Innen-Wildcards, also muss jeder byte-identisch bleiben.
- [x] `make gates` **und** `make ci` grün mit echter Ausgabe; Benutzerhandbuch-Currency.

## 6. Closure-Notiz

_(beim Abschluss.)_

---

## 7. Ergebnis (2026-07-25)

Spec-first geliefert: [ADR-0028](../../adr/0028-ziel-glob-schattenwurf.md) `Accepted` →
Spezifikation **0.25.0** ([SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung))
→ Code → Tests; Benutzerhandbuch **1.34**, [CHANGELOG](../../../../CHANGELOG.md) unter `[Unreleased]`.

**Verifikation:**

| Probe | Ergebnis |
|---|---|
| Fehlbefund (`outbound` → verschachtelter Port über deklarierte Kante) | **weg** — 0 Befunde |
| **Diskriminierungs-Probe** (Guard testweise abgeschaltet) | `TestNestedPortImportNoFalsePositive` **fällt** — der Test prüft wirklich den Fix |
| Ziel **ohne** Tail-Marker (`…/createorder`) | `wrong-direction` meldet weiter |
| `domain` → `application` | `core-impurity` meldet weiter (kategorisch) |
| sauberes Literal-Glob | löst weiter auf; echter Verstoß meldet weiter |
| Glob ohne Tail-Marker (`…/**/**`) | Verhalten unverändert |
| kurzer Wildcard-Kopf (`src/**/ports/**`) gegen spezifischeres Literal-Glob | Zuordnung bleibt |
| **Konsumenten-Regression** (b-cad, d-check, d-migrate, m-trace, belief-agent, HexSlice-Go-Beispiel) | je `exit=0`, 0 Befunde — unverändert |
| `make ci` | grün (lint, Tests, Coverage, `arch-check` 0, `doc-check` 0, `image-test` OK) |

**Lerneinträge**

1. **Eine ausgewiesene Grenze deckt ihre Nebenwirkung nicht mit.** Dass Globs mit Innen-Wildcard
   als Ziel nicht auflösen, stand seit [ADR-0026](../../adr/0026-hexslice-vertical-slice-regeln.md)
   in der Doku. Was dort **nicht** stand: was stattdessen passiert. Der Rückfall auf das
   umschließende Glob war die eigentliche Fehlerquelle — eine Grenze zu dokumentieren heißt auch,
   das Verhalten **an** der Grenze zu benennen.
2. **Ein Falsch-Positiv, dessen Reparatur ein Falsch-Negativ ist, wiegt schwerer als eine Lücke.**
   Diese Klasse gehört bei der Priorisierung über die bloße „fehlt noch"-Lücke gestellt.
3. **Über-Unterdrückung ist die naheliegende Falle beim Zurücknehmen.** Der erste Entwurf des
   Rückzugs (nur Kopf-Vergleich) hätte **jeden** Application-Import verschluckt und die Schicht
   still abgeschaltet — schlimmer als der Ausgangsfehler. Der Tail-Marker macht den Rückzug eng;
   die vier Gegenproben sichern ihn.
4. **Die Diskriminierungs-Probe gehört zum Test, nicht zum Gefühl.** Der Guard wurde testweise
   abgeschaltet, um zu belegen, dass der neue Test ohne den Fix wirklich fällt — dieselbe Lehre wie
   die False-Green-Regression aus [slice-035](../done/slice-035-exclude-verzeichnis-prune.md).
