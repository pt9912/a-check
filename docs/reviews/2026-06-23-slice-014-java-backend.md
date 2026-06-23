# Review — slice-014 Java-Sprach-Backend

**Datum:** 2026-06-23 · **Slice:** [slice-014](../plan/planning/done/slice-014-java-backend.md) ·
**Anforderung:** [AC-FA-EXTRACT-001](../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion)
(Lastenheft 0.7.0) · innerhalb [ADR-0002](../plan/adr/0002-text-heuristische-extraktion.md).

**Methode:** vier perspektiven-diverse Linsen — *Vertrag/Spec* + *Regelwerk/Konvention*
(Plan-Review, eigene Linse) sowie *Code-Korrektheit* + *Test-Abdeckung* (unabhängige
adversarische Agent-Linse, empirisch via `make test`). Proportional zum Diff (~3 Code-Zeilen
+ Tests).

**Gesamtbewertung:** **Kein BLOCKER.** Die Java-Regex ist im Kern korrekt und gegen die
offensichtlichen False-Positives robust. Die Test-Linse deckte reale Härtungs-Lücken auf —
alle vor Closure geschlossen.

## Befunde

| # | Linse | Schwere | Befund | Status |
|---|---|---|---|---|
| K1 | Regelwerk | MAJOR (Plan) | Sweep listete [ADR-0002](../plan/adr/0002-text-heuristische-extraktion.md)-Prosa — `Accepted` ⇒ immutable (AGENTS §3.5) | ✅ aus dem Sweep entfernt; „konsolidiert vier" bleibt Stand-zur-Entscheidungszeit |
| K2 | Test | MAJOR | Regex-Mutanten überlebten die Minimal-3-ACs (`static\s*` vs `\s+`; `static`-im-Pfad; Wildcard) | ✅ 3 Mutanten-Tests (`com.static.Foo`, Mehrfach-Whitespace-`static`, Wildcard) |
| K3 | Code | MINOR | Wildcard-Symbol behält Trailing-Dot (`com.foo.`), un-dokumentiert/un-gepinnt | ✅ Spec-Out-of-Scope-Notiz + Test |
| K4 | Code | MINOR | Mehrere `import`/Zeile → nur der erste gegriffen (Java erlaubt es syntaktisch) | ✅ als dokumentierte Heuristik-Grenze belassen ([AC-QA-02](../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)) |

## Sauber bestätigt (knapp)

- **static-Skip:** `import static com.foo.Bar.baz;` → `com.foo.Bar.baz`, `static` nicht als
  Symbol; `import com.static.Foo;` (static im Pfad) korrekt erhalten — das `(?:static\s+)?`
  ankert direkt hinter `import\s+`.
- **False-Positives** abgewehrt: `importance`, `importer`, `import_x()`, `importstatic com.x`
  → leer (das `\s+` nach `import` verlangt eine Wortgrenze).
- **Determinismus** ([AC-QA-01](../../spec/lastenheft.md#ac-qa-01--determinismus)): `dedupeSort`
  stabil; Spec ([SPEC-EXTRACT-001](../../spec/spezifikation.md#spec-extract-001--import-extraktion))
  ↔ Code konsistent, Historie 0.7.0 in beiden Spec-Straten.

## Abnahme

Code gegen Plan/ADR (Review) **und** gegen DoD/Spec (Verifikation) geprüft; `make gates` grün
(Lint, 6 Java-Tests + Coverage, `arch-check` 0, `doc-check`). Closure mit diesem Slice;
welle-06 bleibt offen.
