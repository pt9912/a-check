# ADR-0034 — Der Kommentar-Strip überspringt Zeichenketten-Literale

- **Status:** Accepted
- **Datum:** 2026-08-15
- **Autor:** pt9912
- **Bezug:** [AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze), [AC-FA-EXTRACT-001](../../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion), [AC-FA-RULE-011](../../../spec/lastenheft.md#ac-fa-rule-011--konstrukt-monopol-regel-construct-leak), [SPEC-EXTRACT-001](../../../spec/spezifikation.md#spec-extract-001--import-extraktion), [ADR-0002](0002-text-heuristische-extraktion.md), [ADR-0027](0027-constructs-roh-text-monopol.md)
- **Schärft:** [SPEC-EXTRACT-001](../../../spec/spezifikation.md#spec-extract-001--import-extraktion) — die Vorbereitung der Quelle vor der Mustersuche.
- **Supersedes:** —

## Kontext

Der Kommentar-Entferner las rohe Bytes ohne Kenntnis von Zeichenketten. Jede Byte-Folge `/*`
**in einem String-Literal** öffnete einen Phantom-Blockkommentar, der bis zum nächsten `*/`
konsumiert wurde — oft bis Dateiende. Alles dazwischen war für **jede** Prüfung unsichtbar:
Importe, `forbidden_constructs`, `constructs`, Kotlin-Deklarationen.

Gemessen (2026-08-15), beide Richtungen:

```text
Go, Verstoss core -> adapters NACH  var globs = []string{"/**"}
  mit dem String  -> gesamt: 0 Befund(e)     EXIT=0
  ohne den String -> core-impurity            EXIT=1

TypeScript, tech-Regel auf "deno.land", Import im Kern
  mit    // in der URL -> gesamt: 0 Befund(e) EXIT=0
  ohne   // in der URL -> core-impurity        EXIT=1
```

Betroffen waren **sieben** Backends: Go, C++, Rust, Kotlin, Java, C#, TypeScript. Glob-Strings sind
in Werkzeug-Code alltäglich, `https://`-URLs in TypeScript/Deno die Norm.

**Die Klasse war bekannt, die Reichweite nicht.** Python ist von der Vorbereitung ausdrücklich
ausgenommen, mit exakt dieser Begründung — *„a `/*`-like byte sequence inside a Python string
literal (e.g. the glob `**/*.py`) would otherwise swallow every real import up to the next `*/` — a
silent false-green"*. Der Schluss, dass dasselbe für alle C-Syntax-Backends gilt, wurde nie
gezogen. Gefunden hat es ein Code-Review, nicht ein Gate.

## Entscheidung

**Ein Zeichenketten-Literal wird verbatim übernommen; ein `/*` oder `//` darin ist Text, kein
Kommentar-Anfang.** Der Scanner kennt drei Begrenzer — `"`, `'` und `` ` `` — mit bewusst groben
Regeln, denn dies ist ein Kommentar-Entferner, kein Lexer
([ADR-0002](0002-text-heuristische-extraktion.md)):

1. **Backslash escapt das nächste Byte**, außer in Backtick-Literalen (Go-Raw-Strings und
   TS-Templates kennen keine Escapes).
2. **`"`- und `'`-Literale enden auch an einem Zeilenumbruch.** Keines von beiden ist in einer
   unterstützten Sprache mehrzeilig; ein unbalanciertes Anführungszeichen kostet damit **höchstens
   seine eigene Zeile**. Ohne diese Regel würde ein einzelnes Apostroph den Rest der Datei
   undurchsichtig machen und das Kommentar-Entfernen ganz aussetzen.
3. **Backtick-Literale sind mehrzeilig** und laufen bis zum schließenden Begrenzer.
4. **Unterminiert am Dateiende:** der Rest wird verbatim kopiert.

**Die Leitplanke, an der alle vier Regeln ausgerichtet sind:** der Scanner darf **nie mehr**
verschlucken als zuvor. Im Zweifel bleibt ein Kommentar stehen — ein **sichtbares** Falsch-Positiv —
statt dass eine Zeile verschwindet, ein **stilles** Falsch-Negativ. Genau diese Asymmetrie war der
Fehler, den die Entscheidung behebt.

## Konsequenzen

- **Verhaltensänderung, die grüne Gates rot machen kann.** Bei Konsumenten werden Importe sichtbar,
  die zuvor verschluckt wurden; wo dort ein echter Verstoß lag, meldet a-check ihn künftig. Das ist
  die **Absicht** — der bisherige Zustand war ein stilles Falsch-Negativ —, aber es ist ein Bruch
  und darum diese ADR ([`AGENTS.md` §3.6](../../../AGENTS.md#36-gates-dürfen-nicht-ohne-adr-gelockert-werden)
  in der Gegenrichtung, Argumentation aus [ADR-0029](0029-abdeckungs-diagnose-advisory.md)).
- **Ausgewiesene Grenze: unbalancierte Apostrophe.** Ein Rust-Lifetime (`&'a str`) ist ein einzelnes
  Apostroph; der Rest **dieser Zeile** gilt dann als Literal, und ihre Kommentare bleiben stehen.
  **Gemessen ist die Reichweite eng:** alle Import-Muster sind zeilenverankert (`^\s*use …`) und
  können von einem Kommentar hinter Code ohnehin nicht ausgelöst werden. Real trifft es nur die
  **nicht** verankerten Muster — `constructs`/`forbidden_constructs` — in derselben Zeile:

  ```text
  fn f(x: &'a str) {} // dlopen ist hier nur erwaehnt
    mit Apostroph  -> construct-leak: Konstrukt dlopen ausserhalb src/plugin   (Falsch-Positiv)
    ohne Apostroph -> gesamt: 0 Befund(e)                                       (korrekt)
  ```

  Die Fehlerrichtung ist sichtbar, nicht still — konform zur Leitplanke. Dasselbe gilt für
  C++-Digit-Separatoren (`1'000'000`), die allerdings paarweise auftreten und sich damit selbst
  ausgleichen.
- **Weiterhin offen, ausdrücklich:** Raw-String-Formen mit eigener Syntax — C++ `R"(…)"`, Rust
  `r#"…"#`, Text-Blöcke in Java/Kotlin/C#. Ein `/*` darin öffnet weiterhin einen Phantom-Kommentar.
  Sie zu decken hieße, je Sprache eigene Literal-Regeln zu führen; das kippt Richtung Lexer und
  damit gegen [ADR-0002](0002-text-heuristische-extraktion.md). Fällt ein realer Fall an, ist das
  eine Folge-Entscheidung.
- **Die Python-Ausnahme bleibt.** Sie ist eine zweite, unabhängige Absicherung; sie zu entfernen,
  weil der Scanner jetzt greift, tauschte einen belegten Schutz gegen eine Annahme.
- **Determinismus unberührt** — der Scanner ist rein textuell und ordnungserhaltend; Zeilennummern
  bleiben stabil, weil Literale mitsamt ihren Zeilenumbrüchen kopiert werden.

## Verworfene Alternativen

- **Nur `"` und `'` behandeln.** Billiger, aber Backtick-Literale sind in Go genau die Stelle, an
  der Globs und Regexe stehen (`` `**/*.go` ``), und in TypeScript die übliche Interpolationsform.
  Derselbe Zustandsautomat deckt sie mit ab; sie auszulassen hieße, den belegten Fall halb zu lösen.
- **Je Sprache eigene Literal-Regeln.** Deckt nahezu alles, kippt aber Richtung Lexer — gegen
  [ADR-0002](0002-text-heuristische-extraktion.md), das die text-heuristische Extraktion als
  Fundament festlegt. Diese Entscheidung schließt eine **Lücke im Kommentar-Strip**, sie ersetzt
  die Heuristik nicht.
- **Alle Sprachen wie Python von der Vorbereitung ausnehmen.** Dann bliebe jeder echte
  Blockkommentar stehen, und ein auskommentierter Import würde als echter gezählt — ein
  Falsch-Positiv in der Breite, das die Regel-Auswertung unbrauchbar machte.

## Fitness Function

- `make test`: ein `/*` bzw. `//` in `"`-, `'`- und Backtick-Literalen verschluckt nichts; **echte**
  Block- und Zeilenkommentare werden weiterhin entfernt; Zeilennummern bleiben stabil.
- **Regressions-Probe:** die Python-Fixture, die den Fehler ursprünglich abfing, bleibt grün.
- `make arch-check` (Dogfooding): unverändert **0** Befunde.
- `make ci` grün.

## Geschichte

| Datum | Ereignis |
|---|---|
| 2026-08-15 | Proposed → Accepted (Sign-off Auftraggeber). Auslöser: Finding `R-1` (HIGH) eines Code-Review-Laufs über den Go-Kern von `v0.17.0`; der Fehler war zu diesem Zeitpunkt **ausgeliefert**. Der Umfang des Scanners wurde auf drei Begrenzer festgelegt und die verbleibenden Raw-Formen als Grenze ausgewiesen, nachdem die Rust-Apostroph-Grenze gemessen und als eng belegt war. Ab jetzt immutable; Ablösung nur via Folge-ADR mit `Supersedes`. |
