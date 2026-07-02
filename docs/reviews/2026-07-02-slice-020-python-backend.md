# Review — slice-020 Python-Sprach-Backend

**Datum:** 2026-07-02 · **Slice:** [slice-020](../plan/planning/done/slice-020-python-backend.md) ·
**Anforderung:** [AC-FA-EXTRACT-001](../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion)
(Lastenheft/Spezifikation 0.11.0) · innerhalb [ADR-0002](../plan/adr/0002-text-heuristische-extraktion.md)
(Extraktion) und [ADR-0016](../plan/adr/0016-resolution-sprach-parametrisch.md) (Auflösung, `fixed-root`-Rezept).

**Methode:** vier perspektiven-diverse, adversarische Agent-Linsen parallel (read-only) —
*Code-Korrektheit* · *Vertrag/Spec-Konsistenz* · *Test-Abdeckung* · *Regelwerk/Konvention* — gegen die
Commits 107aa66 (Entwurf) + 6a3fd64 (Implementierung); Grundwahrheit war das reale Verhalten
(Regex-Semantik, `resolveImport`, `stripComments`). Der Code-MAJOR wurde **empirisch reproduziert**
(Fixture-Repo gegen `a-check:dev`: Exit 0 statt 1). Danach Fixes + **Delta-Re-Review** (eigene
adversarische Linse auf dem eingefrorenen Fix-Diff) + `make gates` + empirische Gegenprobe (Exit 1,
2 Befunde). Der Plan hatte zuvor eine Maintainer-Abnahme (Entscheide A–D) durchlaufen.

**Gesamtbewertung:** **Kein BLOCKER.** Ein substanzieller Code-MAJOR (C-Kommentar-Stripping fraß
Python-Imports → falsch-grün), drei weitere MAJOR (ungepinnte `.`→`/`-Konvertierung, Sweep-Lücke
ARC-003, staler Handbuch-Header) — alle vor Closure gefixt und im Delta-Re-Review bestätigt.

## Befunde

| # | Linse | Schwere | Befund | Status |
|---|---|---|---|---|
| C-1 | Code | MAJOR | `stripComments` (C-artig) frisst in Python nach `/*`-Bytefolge im String-Literal (z. B. `"**/*.py"`) alle Folge-Imports → falsch-grün; empirisch reproduziert (Exit 0 statt 1) | ✅ `prepSource(lang, raw)`: Python wird nicht C-gestrippt (Nicht-Python byte-identisch); Regressionstest + CLI-Fixture-Zeile; [SPEC-EXTRACT-001](../../spec/spezifikation.md#spec-extract-001--import-extraktion) präzisiert |
| T-1 | Test | MAJOR | CLI-Integrationstest pinnte die `.`→`/`-Konvertierung nicht (Einzelsegment-Rest `adapters` = ReplaceAll-No-Op) | ✅ Fixture um `import myapp.adapters.db` (Mehrsegment) erweitert; erwartet 2 Befunde auf Zeile 2+3 |
| T-2 | Test | MAJOR | `pyFrom`-Anker `^` ungepinnt (Kommentar-/Mid-Line-`from…import`) | ✅ `TestPythonFromCommentAndMidline` |
| R-1 | Regelwerk | MAJOR | „fünf → sechs"-Sweep verfehlte `spec/architecture.md` ARC-003 (Rang 3 stand falsch, Rang 6/7 richtig); slice-014-Präzedenz pflegt genau diese Zeile | ✅ ARC-003 auf sechs Sprachen (reiner Enum-Sweep, wie slice-014 ohne Versions-Bump) |
| V-1 | Vertrag | MAJOR | Benutzerhandbuch-Header 1.12/2026-07-01 stale ggü. eigener Historie-Zeile 1.13 | ✅ Header 1.13/2026-07-02 |
| C-2 | Code | MINOR | Rezept-Namenskollision: fremdes Top-Level-Modul mit Schicht-Verzeichnisnamen (`import adapters.db`) löst fälschlich auf → Fehlbefund (Grenze der Paket==Verzeichnis-Voraussetzung, slice-015-Semantik) | ✅ als Grenze im Benutzerhandbuch dokumentiert; Delta-Nachschärfung: `ignore_symbols` ist Substring — Caveat ergänzt (D-1) |
| C-3 | Code | MINOR | Subpaket-Form `from myapp import adapters` liefert nur `myapp` → Verstoß in dieser Form unsichtbar; Konsequenz aus Entscheid A, war aber nicht als Grenze dokumentiert | ✅ Lastenheft Out-of-Scope + Benutzerhandbuch (kanten-relevante Importe voll qualifizieren) |
| T-3…T-5 | Test | MINOR | Zeichenklassen-Mutanten ungepinnt (`_thread`/`__future__`/snake_case/Ziffern); eingerückter `from`-Import; `from a.bimport x` | ✅ je Test ergänzt/erweitert |
| D-1 | Delta | MINOR | `ignore_symbols`-Ausweg ohne Substring-Caveat könnte selbst falsch-grün maskieren | ✅ Handbuch-Passage präzisiert |

## Sauber bestätigt (knapp)

- **Regex-Kontrakte** (Code + Test): Alias (`as`), relative Importe (führender Punkt), `#`-Kommentar
  (Anker), Keyword-Präfixe (`importlib`/`important`/`fromage`), `import`-Wortgrenze beidseitig,
  Erst-Treffer bei `import a, b`; `pyImp`/`pyFrom` gegenseitig exklusiv (kein Doppel-Match).
- **Determinismus** ([SPEC-DET-001](../../spec/spezifikation.md#spec-det-001--determinismus-vertrag)):
  alle Map-Iterationen sortiert, `dedupeSort`/`Extract` stabil — keine Leaks.
- **Vertragskette:** Lastenheft-ACs ↔ SPEC-EXTRACT-001 ↔ Code deckungsgleich; Backend-Menge lebt
  einzig in der Registry (TestBackendRegistrySet pinnt sechs); SPEC-CONF-001 verweist ohne Duplikat;
  `--print-config`-Gerüst ohne stale Aufzählung; Rezept gegen `resolveImport` nachgerechnet.
- **Regelwerk:** keine Accepted-ADR angefasst (§3.5); Spec-Straten referenzieren nicht abwärts (§3.4,
  Historie-Zeilen nach etablierter Praxis); Traceability beider Commits ok; Lifecycle (Datei in `open/`
  bis Closure-`git mv`) konsistent mit slice-015/-017; ID-Linkpflicht erfüllt (CHANGELOG/Fences exempt).
- **Unbekannt-Sprache-Härtung** (slice-017) nach `python`→`ruby`-Umstellung vollwertig; das gepinnte
  Meldungsformat beweist zugleich `python` in der Backend-Menge.

## Delta-Re-Review

Eigenständige adversarische Linse auf dem eingefrorenen Fix-Diff (245 Zeilen): **alle R1-Befunde
adressiert, Root-Cause statt Symptom**; Nicht-Python-Extraktionspfade byte-identisch (prepSource =
reines Delegat); CLI-Zählung/Zeilenanker gegen `report.go` verifiziert; keine tautologischen Tests;
ein MINOR (D-1, `ignore_symbols`-Caveat) — behoben. `make gates` nach den Fixes grün (Coverage 95,90 %
≥ 90 %, Dogfooding 0, doc-check 0); empirische Gegenprobe am Fixture: vorher Exit 0 (Import
verschluckt), nachher Exit 1 mit 2 × `core-impurity` (Zeile 2 `from`-Form, Zeile 3 Mehrsegment-dotted).
