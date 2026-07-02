# slice-021 — C#-Sprach-Backend (welle-06-sprach-backends)

**Status:** done (2026-07-02). Abnahme erteilt (Entscheide A–E gemäß Empfehlung, §6);
Umsetzung + `make gates`/`make ci` + Multi-Linsen-Review (4 Linsen) + Fixes erledigt, Synthese
[`docs/reviews/2026-07-02-slice-021-csharp-backend.md`](../../../reviews/2026-07-02-slice-021-csharp-backend.md).
**Welle:** welle-06-sprach-backends (drittes Backend-Inkrement nach
[slice-014](../done/slice-014-java-backend.md)/[slice-020](../done/slice-020-python-backend.md)).
**Bezug:** erweitert [AC-FA-EXTRACT-001](../../../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion)
um C#; innerhalb [ADR-0002](../../adr/0002-text-heuristische-extraktion.md)
(text-heuristisch, **kein** neuer ADR); schärft
[SPEC-EXTRACT-001](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion).
**Auflösung:** heute über den gelieferten `fixed-root`-Modus
([ADR-0016](../../adr/0016-resolution-sprach-parametrisch.md), [slice-015](../done/slice-015-resolution-roots.md))
unter der **Namespace==Verzeichnis-Konvention** (.NET-Default); der **allgemeine** Fall
(Namespace frei je Datei, ≠ Ordner) bleibt beim **reservierten `namespace`-Modus**
(Exit 2) und bekommt **einen eigenen Folge-Slice + Folge-ADR**, getriggert durch einen
realen C#-Pilot mit Namespace≠Verzeichnis (§6 Entscheid D).
[Roadmap welle-06](../in-progress/roadmap.md). **Trigger:** Polyglot-Bestand (C#-Repos),
Maintainer-Priorität 2.

> **Hinweis:** Entwurf zur Abnahme. Die in §3 als Code-Fence gesetzten AC-Texte sind
> unverbindlich — gültig erst nach Freigabe in [`spec/lastenheft.md`](../../../../spec/lastenheft.md).
> DoD §5 offen; Entscheidungen §6 **vor** der Umsetzung zu treffen.

---

## 1. Ziel

Ein **C#**-Backend für die Import-Extraktion (`using`-Direktiven), damit C#-Repos ihre
Hexagon-Architektur über a-check prüfen können — `tech`-Regeln sofort (matchen das
Roh-Symbol), Schicht-Regeln über das `fixed-root`-Rezept, wo Namespaces die
Verzeichnisse spiegeln. Reine Extraktions-Erweiterung; kein neuer Auflösungs-Modus.

## 2. Problem

a-check v0.5.0 kennt {`cpp`, `go`, `rust`, `kotlin`, `java`, `python`}; `languages:
{csharp: …}` bricht (korrekt, [slice-017](../done/slice-017-unbekannte-sprache-exit2.md))
mit Exit 2. C# importiert über **`using`-Direktiven** — dotted wie JVM, aber mit
Eigenheiten, die eine eigene Regex brauchen:

- `using System.Text;` — Grundform, `;`-terminiert.
- `using static System.Math;` / `global using MyApp.Core;` — Schlüsselwörter
  `static`/`global` (analog Javas `import static`).
- `using Db = MyApp.Adapters.Db;` — **Alias**: die echte Abhängigkeit steht **rechts**.
- **Kollisionsgefahr:** `using var f = …;` und `using (var f = …)` sind
  **Ressourcen-Statements**, keine Direktiven — sie dürfen **nie** als Import gewertet
  werden (der C#-spezifische Kern-Mutant dieses Slices).

**Auflösung:** C#-Namespaces sind *frei deklarierbar* (Namespace ≠ Ordner möglich) —
das ist das Signal des reservierten `namespace`-Modus (Namespace→Datei-Index,
Cross-File; [slice-015 §4](../done/slice-015-resolution-roots.md)). Die verbreitete
.NET-Konvention ist aber Namespace==Verzeichnis — dort trägt das gelieferte
`fixed-root`-Rezept bereits (dieselbe [AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)-Grenze
„Paket==Verzeichnis" wie bei JVM/Python).

## 3. Entwurf (zur Abnahme)

### 3.1 Anforderungs-Erweiterung — AC-FA-EXTRACT-001 (C#)

```text
AC-FA-EXTRACT-001 (erweitert um C#): die Backend-Liste wird um C# ergänzt
(languages-Schlüssel `csharp`) — using-Direktiven liefern den gepunkteten
Namespace; `static`/`global` werden übersprungen, bei der Alias-Form
(`using X = Ziel;`) wird das Ziel (rechte Seite) gewertet, das `;` ist Pflicht
im Muster. using-STATEMENTS (`using var x = …;`, `using (…)`) sind keine
Direktiven und werden nie gewertet.

Neue/ergaenzte Akzeptanzkriterien:
- Happy (C#): Given `using MyApp.Adapters.Db;`, when das C#-Backend laeuft,
  then liefert es das Symbol `MyApp.Adapters.Db`.
- Boundary (C# static/global): Given `using static System.Math;` und
  `global using MyApp.Core;`, when das Backend laeuft, then liefert es
  `System.Math` bzw. `MyApp.Core` (Schlüsselwörter übersprungen).
- Boundary (C# Alias): Given `using Db = MyApp.Adapters.Db;`, when das Backend
  laeuft, then liefert es `MyApp.Adapters.Db` (das Ziel, nicht den Alias).
- Negative (C# using-Statement): Given `using var f = File.Open(p);` oder
  `using (var f = File.Open(p))`, when das Backend laeuft, then wird KEIN
  Symbol geliefert.

Out-of-Scope: Typ-Aliasse auf generische Typen (`using L = List<int>;` — kein
Namespace-Import, nicht gegriffen); `extern alias`; Namespace-DEKLARATIONEN
(`namespace X;`/`namespace X { }`) werden nicht als Import gewertet;
Toolchain-Backends (Roslyn/MSBuild).
```

### 3.2 Versions-Bump

Lastenheft + Spezifikation **0.11.0 → 0.12.0**. „sechs → sieben Sprachen"-Sweep:
README, Benutzerhandbuch (§1/§4 + Beispiel + Historie),
**`spec/architecture.md`** ([ARC-003](../../../../spec/architecture.md)-Sprachliste —
Lerneintrag slice-020: zählende Stellen explizit benennen).

### 3.3 Auflösung: kein Schema-Delta; Namespace==Verzeichnis-Rezept, Rest ehrlich reserviert

```yaml
# .NET-Konvention: src/MyApp/{Domain,Ports,Adapters}/…, Namespaces MyApp.…
resolution:
  csharp: {mode: fixed-root, roots: ["src/MyApp"], package_base: "MyApp"}
```

Wo ein Repo die Konvention **nicht** hält (frei deklarierte Namespaces), bleibt das
Symbol **unaufgelöst** (keine schicht-basierte Regel — ehrliche
[AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)-Grenze,
identisch zur JVM/Python-Formulierung); `tech`-Regeln greifen unabhängig davon (Roh-Symbol).
Der reservierte `resolution.mode: namespace` (Namespace→Datei-Index) bleibt Exit 2 und
wird **nicht** in diesem Slice gebaut (§6 Entscheid D).

## 4. Umsetzungsplan

1. `internal/adapter/driven/extract/extract.go`: Feld `csUsing`, Regex
   `^\s*(?:global\s+)?using\s+(?:static\s+)?(?:[A-Za-z_][A-Za-z0-9_]*\s*=\s*)?([A-Za-z_][A-Za-z0-9_.]*)\s*;`
   in `newAdapter`; Registry-Eintrag `"csharp"` via `lineMatches`. Das **Pflicht-`;`
   direkt nach dem dotted-Namen** schließt die using-Statements aus (`using var x = …`:
   nach `var` kommt kein `;`; `using (` matcht `[A-Za-z_]` nicht). C# ist C-Syntax →
   `prepSource` lässt das `//`-//`/* */`-Stripping **an** (anders als Python).
2. Tests (`extract_test.go`) — Mutanten-Boundary nach slice-014/020-Lerneintrag:
   Happy dotted; `static`/`global` übersprungen (+ `static` als Namespace-Segment
   `using com.static.X;` bleibt erhalten); Alias → rechte Seite; **`using var`/
   `using (…)`-Statements → kein Match** (Kern-Mutant, pinnt das Pflicht-`;`);
   fehlendes `;` → kein Match; `usings`/`usingFoo` (Keyword-Präfix); `// using Evil;`
   (C-Strip greift); Einrückung/Mehrfach-Whitespace; Underscore/Ziffern
   (`using _Internal.V2;`); generischer Typ-Alias → kein Match.
3. **slice-017-Fixture-Umstellung** (wie `python`→`ruby` in slice-020): `csharp` ist
   dann ein gültiges Backend — `TestCheckLanguagesMixedUnsupported` braucht eine andere
   vor-`go` sortierende Unbekannt-Sprache (z. B. `fsharp`).
4. Resolution-Integrationstest (CLI): C#-Domänen-Datei mit `using MyApp.Adapters.Db;`
   + Rezept (§3.3) → `core-impurity`, Exit 1; inkl. Mehrsegment-Rest (pinnt `.`→`/`,
   Lerneintrag slice-020) und einer `"**/*.cs"`-ähnlichen String-Zeile ist hier **nicht**
   sinnvoll (C-Strip bleibt an — siehe Risiko-Notiz §6).
5. Spec: Lastenheft ([AC-FA-EXTRACT-001](../../../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion)
   + Bump 0.12.0 + Historie); [SPEC-EXTRACT-001](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion)
   (C#-Muster + Backend-Menge `{cpp, go, rust, kotlin, java, python, csharp}` + Bump).
6. „sechs → sieben"-Sweep (§3.2) inkl. Benutzerhandbuch-Rezept-Absatz (C# neben Python).
   **Nicht** [ADR-0002](../../adr/0002-text-heuristische-extraktion.md)/[ADR-0016](../../adr/0016-resolution-sprach-parametrisch.md) (`Accepted` ⇒ immutable).
7. `make gates`/`make ci`; Multi-Linsen-Review (4 Linsen) + ggf. Delta →
   `docs/reviews/`; Verifikation; Closure (reiner `git mv`, 2 Kriterien, Lerneintrag).

## 5. Definition of Done

- [x] [AC-FA-EXTRACT-001](../../../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion)
      um C# erweitert (Happy/`static`+`global`/Alias/using-Statement-Negative +
      Out-of-Scope), Bump 0.12.0 + Historie; [SPEC-EXTRACT-001](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion)
      nachgezogen (Muster + Backend-Menge + [SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung)-dotted-Liste, Review-R1).
- [x] `extract.go`: `csUsing` + Registry-Eintrag; Dogfooding grün (0 Befunde).
- [x] Tests: §4.2-Mutanten (insb. using-Statement-Ausschluss; um Review-R1 erweitert:
      `^`-Anker/String, `staticData`-Keyword-Präfix) + slice-017-Fixture-Umstellung
      + CLI-Resolution-Integration (Mehrsegment; pinnt seit Review-R1 auch den
      `;`-Kern-Mutanten end-to-end).
- [x] „sechs → sieben Sprachen"-Sweep vollständig (README, Benutzerhandbuch,
      [ARC-003](../../../../spec/architecture.md); Review-R1-Nachzug: stale
      Vier-Sprachen-Zählstelle in `harness/README.md` §Safety geheilt).
- [x] `make gates` + `make ci` grün; Multi-Linsen-Review (4 Linsen) + Delta bestanden
      ([Synthese](../../../reviews/2026-07-02-slice-021-csharp-backend.md)).
- [x] **Maintainer-Abnahme der Entscheide A–E (§6)** — erteilt 2026-07-02 (Abnahme-Block §6).
- [x] Closure: reiner `git mv` nach `done/` (AGENTS §3.3); 2 beobachtbare Kriterien + Lerneintrag (§7).

## 6. Offen / Entscheidungen zur Abnahme

> **Abnahme (2026-07-02):** Entscheide A–E gemäß Empfehlung bestätigt (Maintainer-Wort,
> nachträglich zur Empfehlungs-basierten Umsetzung — der Zwischenstand war als
> „Abnahme ausstehend" deklariert, Closure/Push bis hierhin zurückgehalten).

- **Entscheid A — Alias-`using`:** Ziel (rechte Seite) werten (`using Db = MyApp.X.Db;`
  → `MyApp.X.Db`; Empfehlung, Rust-Präzedenz `use x as y` → `x`) vs. Alias-Form gar
  nicht greifen. *Empfehlung: Ziel werten — es ist die echte Abhängigkeit; der Alias
  selbst wird nie als Symbol geliefert.*
- **Entscheid B — `global using` / `using static`:** Schlüsselwörter überspringen
  (Empfehlung, Java-`static`-Präzedenz) vs. nur Grundform. *Empfehlung: überspringen —
  beide Formen sind Alltags-C# (C# 10 `GlobalUsings.cs`).*
- **Entscheid C — using-Statement-Ausschluss über Pflicht-`;` nach dem dotted-Namen:**
  `using var f = …;`/`using (…)` dürfen nie matchen. *Empfehlung: bestätigen — das ist
  der C#-spezifische Kern-Mutant; ohne `;`-Anker wäre jede Ressourcen-Zeile ein
  falsches Import-Symbol (`var`).*
- **Entscheid D — `namespace`-Modus NICHT in diesem Slice:** Extraktion + fixed-root-Rezept
  jetzt; der Namespace→Datei-Index (frei deklarierte Namespaces) bleibt reserviert
  (Exit 2) und bekommt **eigenen Folge-Slice + Folge-ADR zu
  [ADR-0016](../../adr/0016-resolution-sprach-parametrisch.md)**, getriggert durch einen realen
  C#-Pilot mit Namespace≠Verzeichnis. *Empfehlung: bestätigen — anderes Auflösungs-Signal
  (Cross-File-Index ≠ Symbol-Normalisierung), Gating-Doktrin der Welle; die
  Namespace==Verzeichnis-Grenze ist ehrlich dokumentiert, kein stiller No-Op.*
- **Entscheid E — Sprach-Schlüssel `csharp`:** (nicht `c#`/`cs`) — konsistent mit dem
  bisherigen Fixture-Gebrauch und ohne Sonderzeichen im YAML-Schlüssel. *Empfehlung:
  `csharp`.*
- **Risiko/Notiz — geteilte Vorverarbeitung (Lerneintrag slice-020 angewandt):** C# ist
  C-Syntax — `prepSource` lässt das Kommentar-Stripping **an** (korrekt: `// using Evil;`
  muss neutralisiert werden). Damit teilt C# die **bestehende** C-Familien-Grenze: eine
  `/*`-Bytefolge in einem String-Literal (`var g = "**/*.cs";`) kann Folge-Code bis zum
  nächsten `*/` verschlucken — für C++/Go/Java/Kotlin/Rust seit 0.1.0 dokumentierte
  String-Grenze ([AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)),
  kein C#-Spezifikum. String-bewusstes Stripping wäre ein eigenes, sprachübergreifendes
  Inkrement (nicht dieser Slice).
- **Risiko/Notiz — synthetische Verifikation:** noch kein benannter C#-Pilot; Verifikation
  gegen eigene Fixtures (wie Java/Python). Sprache bleibt gated-geliefert.

## 7. Closure-Notiz

**Abschluss (2026-07-02).** slice-021 (welle-06 — drittes Sprach-Backend) umgesetzt, reviewt,
abgenommen und gate-belegt.

- **Gate-Beleg:** `make gates` grün — `lint` 0 issues, alle Test-Pakete `ok`, `coverage-gate`
  96,00 % (≥ 90 %), `arch-check` **0** (Dogfooding), `doc-check` 67/0,
  gate-consistency/guard-selftest/record-gates ok; `make ci` (inkl. `image-test` 4/4) grün.
- **Code:** `csUsing`-Regex + Registry-Eintrag `"csharp"` (rein additiv; sechs Bestands-Backends
  byte-identisch). Der Alias-Zweig matcht strukturell nur die echte Alias-Direktive (ein Token vor
  `=`); das Pflicht-`;` direkt nach dem Namen schließt using-Statements/Declarations kategorisch
  aus. C-Strip bleibt für C# an (`prepSource` — Lerneintrag [slice-020](../done/slice-020-python-backend.md)
  angewandt und adversarisch bestätigt).
- **Review (4 Linsen + Fixes):** Code-Linse ohne Befund; 1 MAJOR (ungepinnter `^`-Anker → 
  String-/Mid-Line-Test) + 4 MINOR gefixt (Details:
  [Synthese](../../../reviews/2026-07-02-slice-021-csharp-backend.md)). Kein BLOCKER.
- **Prozess-Sonderfall dokumentiert:** Umsetzung lief nach dokumentierten Empfehlungen bei
  ausstehender Abnahme (Timeout der Rückfrage); Closure/Push wurden zurückgehalten, der Zustand war
  im Status deklariert; Abnahme nachträglich erteilt (§6-Block).
- **2 beobachtbare Kriterien:** (1) `TestCsharpFixedRootResolution` + Container-Gegenprobe —
  C#-Domäne mit `using MyApp.Adapters.Db;` + Alias `using Db2 = MyApp.Adapters.Db2;` unter dem
  Rezept ⇒ Exit 1 mit genau 2 × `core-impurity` (Zeile 1 + 2), die using-Declaration
  `using MyApp.Adapters.Db t = Get();` bleibt befundfrei (pinnt Alias-Ziel, Mehrsegment-`.`→`/`
  **und** den `;`-Anker end-to-end). (2) `TestCheckLanguagesUnknown` —
  `unbekannte Sprache "ruby" (cpp|csharp|go|java|kotlin|python|rust)`: das gepinnte Meldungsformat
  beweist `csharp` in der Backend-Menge.
- **Lerneintrag (geschärfte Regel):** Ein Integrationstest, der eine Negativ-Grenze „bewacht",
  muss so konstruiert sein, dass der bewachte Mutant ihn **wirklich bricht** — die ursprüngliche
  Fixture-Zeile `using var f = …` ließ den `;`-Kern-Mutanten durch, weil das Fehlsymbol `var` auf
  keine Schicht auflöste (Befund T-3): Negativ-Zeilen so wählen, dass das Fehlsymbol auf eine
  Schicht auflösen *würde*. Zweitens bewährt: bei ausstehender Abnahme ist das **Verzeichnis** die
  autoritative Lifecycle-Quelle (Datei bleibt in `open/`), der Sonderzustand gehört explizit in die
  Status-Zeile, und Closure/Push bleiben kategorisch zurückgehalten.
- welle-06 bleibt **offen** (letzter Kandidat TypeScript — braucht den reservierten
  `relative`-Modus per Folge-ADR); slice-021 ist ihr drittes Inkrement.
