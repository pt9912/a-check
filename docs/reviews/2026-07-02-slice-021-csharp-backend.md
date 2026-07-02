# Review — slice-021 C#-Sprach-Backend

**Datum:** 2026-07-02 · **Slice:** [slice-021](../plan/planning/done/slice-021-csharp-backend.md) ·
**Anforderung:** [AC-FA-EXTRACT-001](../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion)
(Lastenheft/Spezifikation 0.12.0) · innerhalb [ADR-0002](../plan/adr/0002-text-heuristische-extraktion.md)
(Extraktion) und [ADR-0016](../plan/adr/0016-resolution-sprach-parametrisch.md) (Auflösung via
`fixed-root`-Rezept; der reservierte `namespace`-Modus bleibt Exit 2).

**Prozess-Sonderfall:** Die Maintainer-Abnahme der §6-Entscheide A–E lief in ein Timeout; die
Umsetzung folgte den dokumentierten **Empfehlungen**, der Slice-Status deklariert die **Abnahme als
ausstehend** — Closure (`git mv` nach `done/`) und Push sind bis zur Bestätigung zurückgehalten
(von der Regelwerk-Linse als prozesskonform eingedämmt bewertet; einziger Vorbehalt: der
Lastenheft-Bump 0.12.0 liegt vor der Abnahme, ist aber lokal und rückholbar).

**Methode:** vier perspektiven-diverse, adversarische Agent-Linsen parallel (read-only) —
*Code-Korrektheit* · *Vertrag/Spec-Konsistenz* · *Test-Abdeckung* · *Regelwerk/Konvention* — gegen
die Commits 4f4c613 (Entwurf) + 4e27634 (Implementierung); Grundwahrheit war das reale Verhalten
(RE2-Matching der `csUsing`-Regex, `resolveImport`). Danach Fixes + Delta-Verifikation
(`make gates`) + empirische Fixture-Gegenprobe (Container-Scan: 2 × `core-impurity`, Statement- und
`using static`-Zeilen befundfrei, Exit 1).

**Gesamtbewertung:** **Kein BLOCKER.** Die Code-Linse konnte die Regex trotz gezielter Angriffe
(Alias-Gruppe vs. using-Declarations, voll qualifizierte Typen, `global::`, CRLF, Präprozessor)
**nicht widerlegen** — das strukturelle Argument: der Alias-Zweig matcht genau *ein* Token vor `=`,
jede using-Declaration hat *zwei*; das Pflicht-`;` direkt nach dem Namen fängt den Rest. Ein
Test-MAJOR (ungepinnter `^`-Anker) und vier MINOR — alle gefixt.

## Befunde

| # | Linse | Schwere | Befund | Status |
|---|---|---|---|---|
| T-1 | Test | MAJOR | `^`-Anker-Mutant ungepinnt: kein C#-Test mit Inhalt vor `using` — `Log("using Sneaky.Ns;")` würde beim Anker-Verlust fälschlich matchen; zugleich war die String-Hälfte der generischen Negativ-AC für C# ungetestet | ✅ `TestCsharpStringNotCounted` (String-Literal + Mid-Line) |
| T-2 | Test | MINOR | static-Gruppen-Whitespace-Mutant (`static\s+`→`\s*`) überlebte: `using staticData.X;` ergäbe `Data.X` | ✅ `TestCsharpStaticKeywordPrefix` |
| T-3 | Test | MINOR | CLI-Integrationstest tötete den `;`-Kern-Mutanten nicht end-to-end (das Fehlsymbol `var` löste auf keine Schicht auf) | ✅ Fixture-Zeile 3 = using-Declaration mit qualifiziertem Adapter-Typ (`using MyApp.Adapters.Db t = Get();`) — ohne `;`-Anker gäbe es einen dritten Befund |
| V-1 | Vertrag | MINOR | Illustrative dotted-Liste in [SPEC-RULE-001](../../spec/spezifikation.md#spec-rule-001--regel-auswertung) sagte „(JVM/Python)" ohne C# — Parallelitäts-Lücke zum Benutzerhandbuch | ✅ „(JVM/Python/C#)" |
| R-1 | Regelwerk | MINOR | Pre-existing Sweep-Lücke (seit Java übersehen): `harness/README.md` §Safety zählte noch „C++/Go/Rust/Kotlin" | ✅ auf sieben Sprachen gezogen |

## Sauber bestätigt (knapp)

- **Kern-Mutant using-Statements** (Code + Test): `using var f = …;`, `using (…)`,
  `using FileStream f = …;`, `using System.IO.FileStream f = …;` (voll qualifiziert — der
  gefährlichste Fall), `using X y = new X { A = B.C };` matchen alle nicht; der Alias-Zweig feuert
  ausschließlich bei der echten Alias-Direktive `using IDENT = DOTTED;` und liefert das Ziel.
- **`global`/`static`** übersprungen, nie als Symbol; Keyword-als-Segment (`com.static.Foo`)
  erhalten; Generic-Typ-Alias/`namespace`-Deklaration/`extern alias`/`global::` nicht gegriffen.
- **Auflösung** (Rezept `{fixed-root, roots: ["src/MyApp"], package_base: "MyApp"}`):
  `MyApp.Adapters.Db` → adapters-Layer; Fremd-Namespaces (`System.Text`) bleiben unaufgelöst;
  `MyApplication.Foo` wird dank `package_base+"."`-Prefix **nicht** gestript (keine Kollision).
- **Regressionen:** Diff rein additiv (Feld, Compile, Registry-Eintrag) — die sechs
  Bestands-Backends laufen byte-identisch; Determinismus (sortierte Map-Iterationen) unverändert.
- **prepSource-Entscheidung** (Lerneintrag slice-020 angewandt): C-Strip für C# bewusst **an**
  (`// using Evil;` muss neutralisiert werden, test-gepinnt); Zusatzargument der Code-Linse: in
  gültigem C# stehen `using`-Direktiven vor jedem Code/String, die C-Familien-String-Grenze kann
  dort keine Direktiven verschlucken.
- **Regelwerk:** keine Accepted-ADR angefasst; Spec-Straten ohne Abwärts-Referenzen im normativen
  Text; Traceability beider Commits ok; Sweep inkl. [ARC-003](../../spec/architecture.md) (Lerneintrag
  slice-020 angewandt); slice-017-Fixture-Umstellung `csharp`→`fsharp` vollwertig (`fsharp` < `go`).

## Delta-Verifikation

Die Fixes sind rein additiv (drei Tests, zwei Doku-Enum-Zeilen, keine Produktions-Code-Änderung):
`make gates` nach den Fixes grün — lint 0 issues, alle Test-Pakete `ok`, Coverage 96,00 % (≥ 90 %),
`arch-check` (Dogfooding) 0, `doc-check` 66/0. Empirische Gegenprobe am C#-Fixture (Container):
Direktive + Alias-Ziel ⇒ 2 × `core-impurity` (Zeile 1 + 2), `using var`-Statement und
`using static System.Math` befundfrei, Exit 1.

## Offen (vor Closure)

Maintainer-Abnahme der Entscheide A–E (§6 des Slice-Docs); danach Abnahme-Block eintragen,
Closure-Notiz (§7) mit 2 beobachtbaren Kriterien + Lerneintrag, reiner `git mv` nach `done/`,
Roadmap-Nachführung, Push auf Wort.
