# slice-014 — Java-Sprach-Backend (welle-06-sprach-backends)

**Status:** in-progress (Entwurf zur Abnahme).
**Welle:** welle-06-sprach-backends (erster Trigger — Konsumenten-Bedarf).
**Bezug:** erweitert [AC-FA-EXTRACT-001](../../../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion)
um Java; innerhalb [ADR-0002](../../adr/0002-text-heuristische-extraktion.md)
(text-heuristisch, **kein** neuer ADR); schärft
[SPEC-EXTRACT-001](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion).
[Roadmap welle-06](../in-progress/roadmap.md). **Trigger:** `/Development/KI/belief-agent` (Java-Repo
in Spec-Phase) braucht a-check; v0.3.0 kennt nur C++/Go/Rust/Kotlin.

> **Hinweis:** Entwurf zur Abnahme. Die in §3 als Code-Fence gesetzten AC-Texte sind
> unverbindlich — gültig erst nach Freigabe in [`spec/lastenheft.md`](../../../../spec/lastenheft.md).
> DoD §5 offen; Entscheidungen §6 **vor** der Umsetzung zu treffen.

---

## 1. Ziel

Ein **Java**-Backend für die Import-Extraktion, analog Kotlin — damit Java-Repos
(zuerst belief-agent) ihre Hexagon-Architektur über a-check + `.a-check.yml` prüfen
können, ohne dass die Engine sich ändert. Reine Extraktions-Erweiterung; keine neue
Regel, kein neues Modell-Konzept.

## 2. Problem

a-check v0.3.0 wählt das Extraktions-Backend über `languages` ∈ {`go`, `cpp`, `rust`,
`kotlin`} (`extract.go` `importsFromSource`). Eine Java-Datei matcht kein Backend →
ihre Imports werden nicht extrahiert → Java-Repos sind nicht prüfbar. Java nutzt
`import <dotted.path>;` — fast deckungsgleich mit Kotlin, mit **einer** echten
Abweichung: `import static <dotted.path>;` (Member-Import); die Kotlin-Regex griffe
fälschlich `static` als Symbol.

## 3. Entwurf (zur Abnahme)

### 3.1 Anforderungs-Erweiterung — AC-FA-EXTRACT-001 (Java)

```text
AC-FA-EXTRACT-001 (erweitert um Java): die Backend-Liste wird um Java ergänzt —
C++ (#include), Go (import), Rust (use/extern crate), Kotlin (import), Java (import,
inkl. import static). Java teilt die Kotlin-Punkt-Pfad-Form; die Heuristik überspringt
das optionale static-Schlüsselwort und ignoriert das abschliessende ';'.

Neue/ergaenzte Akzeptanzkriterien (zu den bestehenden Happy/Boundary/Negative):
- Happy (Java): Given `import com.foo.Bar;`, when das Java-Backend laeuft, then liefert
  es das Symbol `com.foo.Bar`.
- Boundary (Java static): Given `import static com.foo.Bar.baz;`, when das Java-Backend
  laeuft, then liefert es `com.foo.Bar.baz` (das `static` wird uebersprungen, nicht als
  Symbol gewertet).
- Negative bleibt sprach-agnostisch (import-aehnliche Zeile in Kommentar/String wird
  nicht gewertet — bestehende Heuristik-Grenze AC-QA-02).

Out-of-Scope: Java-Toolchain-Backends (javac/jdeps); Annotations-/Reflection-Importe;
package-Statement-Auswertung; Wildcard-Imports (`import com.foo.*;`) werden wie
gehabt heuristisch gegriffen (Symbol `com.foo.` mit Praefix-Match), nicht expandiert.
```

### 3.2 Versions-Bump

Lastenheft + Spezifikation **0.6.0 → 0.7.0** (neue Sprach-Unterstützung, MINOR; spätere
Software-Version v0.4.0). `vier → fünf Sprachen` über die Doku.

## 4. Umsetzungsplan

1. `internal/adapter/driven/extract/extract.go`: `Adapter`-Feld `javaImp`; Regex
   `^\s*import\s+(?:static\s+)?([A-Za-z_][A-Za-z0-9_.]*)` in `newAdapter`; `case "java":
   return dedupeSort(lineMatches(src, a.javaImp))` in `importsFromSource`.
2. Tests (`extract_test.go`): Java-Happy (dotted), `import static` (static übersprungen),
   `;`-Toleranz, Kommentar-/String-Grenze (sprach-agnostisch) — [AC-FA-EXTRACT-001](../../../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion).
3. Spec: `spec/lastenheft.md` ([AC-FA-EXTRACT-001](../../../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion) + Bump 0.7.0 + Historie),
   `spec/spezifikation.md` ([SPEC-EXTRACT-001](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion):
   Java-Muster + Bump), kein ADR-Index-Eintrag (kein neuer ADR).
4. „vier → fünf Sprachen"-Sweep: README, `spec/architecture.md` ([ARC-003](../../../../spec/architecture.md)),
   Benutzerhandbuch (§1/§4 `languages`-Enum `go/cpp/rust/kotlin/java` + Beispiel + Historie).
   **Nicht** [ADR-0002](../../adr/0002-text-heuristische-extraktion.md): `Accepted` ⇒ immutable
   (AGENTS §3.5); sein „konsolidiert vier" bleibt als Stand-zur-Entscheidungszeit.
5. `make gates`; **4-Linsen-Review** (Code · Vertrag/Spec · Test · Regelwerk) schriftlich
   nach `docs/reviews/`; **Verifikation** gegen DoD/Spec; Commit(s).

## 5. Definition of Done

- [x] [AC-FA-EXTRACT-001](../../../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion) um Java erweitert (Happy + Boundary-static AC + Out-of-Scope),
      Bump 0.7.0 + Historie; [SPEC-EXTRACT-001](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion) nachgezogen.
- [x] `extract.go`: `javaImp` + `case "java"`; `make arch-check` (Dogfooding) grün.
- [x] Tests: Java-Happy / `import static` / `;`-Toleranz / Kommentar-Grenze.
- [x] „vier → fünf Sprachen"-Sweep vollständig (README, architecture, Benutzerhandbuch — **ohne** [ADR-0002](../../adr/0002-text-heuristische-extraktion.md), immutable).
- [x] `make gates` grün; 4-Linsen-Review bestanden (schriftlich → `docs/reviews/`); Verifikation gegen DoD/Spec.
- [x] Closure: **reiner** `git mv` nach `done/` (§3.3, getrennt von Inhalts-Edits); **2 beobachtbare Kriterien** (`make gates`-Beleg + Java-Happy/`static`-Tests im §7 verlinkt) + **Lerneintrag**.

## 6. Offen / Entscheidungen zur Abnahme

- **Entscheid A — `import static`-Handling:** Regex mit `(?:static\s+)?` (Empfehlung,
  robust) vs. Java exakt = Kotlin (greift `static` fälschlich). *Empfehlung: static
  überspringen.*
- **Entscheid B — neue Anforderung vs. Erweiterung:** [AC-FA-EXTRACT-001](../../../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion) **erweitern**
  (Empfehlung — Java ist ein weiteres Backend derselben Anforderung) vs. eine separate
  neue Anforderung im `EXTRACT`-Bereich. *Empfehlung: erweitern.*
- **Entscheid C — ADR?** Kein neuer ADR (innerhalb [ADR-0002](../../adr/0002-text-heuristische-extraktion.md)). *Empfehlung: bestätigt — Java fügt keine neue Architektur-Entscheidung hinzu.*
- **Risiko/Notiz — Dotted-Import-Auflösung (pre-existing, mit Kotlin geteilt):** Java liefert
  punkt-getrennte Symbole (`com.foo.Bar`) wie Kotlin; ob/wie diese gegen pfad-basierte
  Layer-Globs (`targetLayer`/`segIndex` auf `/`) auflösen, ist eine **bestehende** Frage —
  slice-014 fügt **nur Extraktion** hinzu (konsistent mit dem Kotlin-Backend), kein
  Resolutions-Blocker. Relevant für belief-agents künftige `.a-check.yml`, nicht für diese Slice.
- **Risiko/Notiz — synthetische Verifikation:** belief-agent hat noch keinen Java-Code; die
  DoD-Verifikation läuft gegen a-checks **eigene** Java-Fixtures, nicht gegen ein reales
  Konsumenten-Repo (Bedarf per CR bestätigt, Aktivierung folgt).
- **Entscheid D — Wildcard-Imports:** `import com.foo.*;` heuristisch greifen (Symbol mit
  Trailing-Dot, Präfix-Match genügt) statt expandieren. *Empfehlung: so lassen, in
  Out-of-Scope vermerkt.*

## 7. Closure-Notiz

**Abschluss (2026-06-23).** slice-014 (welle-06 — erstes Sprach-Backend nach dem Bootstrap)
umgesetzt und gate-belegt.

- **Gate-Beleg:** `make gates` grün — Lint (gofmt-Alignment), `make test` (6 Java-Tests +
  Coverage), `make arch-check` **0** (Dogfooding), `doc-check` 51/0,
  gate-consistency/guard-selftest/record-gates ok.
- **Code:** `javaImp`-Regex `^\s*import\s+(?:static\s+)?(…)` + `case "java"` in
  `importsFromSource`; Java teilt die Kotlin-Punkt-Pfad-Form, überspringt das optionale `static`.
- **Review (4 Linsen):** Plan-Review (Vertrag/Spec + Regelwerk — die Accepted-ADR-Immutabilität
  im Sweep gefangen) + **unabhängige** adversarische Code-Linse (empirisch via `make test`).
  Befunde: **MAJOR** Test-Abdeckung (Regex-Mutanten überlebten) → 3 Mutanten-Tests
  (`com.static.Foo`, Mehrfach-Whitespace-`static`, Wildcard); **MINOR** Wildcard-Trailing-Dot
  → Spec-Notiz + Test; **MINOR** Multi-Import/Zeile → dokumentierte Heuristik-Grenze.
  Doc: [Review](../../../reviews/2026-06-23-slice-014-java-backend.md). Kein BLOCKER.
- **Verifikation (gegen DoD/Spec):** alle DoD-Haken erfüllt; die Java-Akzeptanzkriterien
  (Happy/Boundary-`static`/Negative) sind durch Tests gepinnt; Lastenheft/Spezifikation 0.7.0 konsistent.
- **2 beobachtbare Kriterien:** (1) `make gates`-Beleg grün; (2) die Java-Tests
  `TestJavaImport`/`…StaticImport`/`…CommentNotCounted`/`…StaticInPath`/`…StaticMultiWhitespace`/`…Wildcard`
  in `extract_test.go`.
- **Lerneintrag (geschärfte Regel):** Die Minimal-3-Akzeptanzkriterien ließen mehrere
  Regex-Mutanten durch (`static\s*` vs `\s+`, `static`-im-Pfad); erst die **unabhängige**
  Review-Linse deckte es auf. Konvention geschärft: **Regex-Backends brauchen Boundary-Tests
  gegen die spezifischen Mutanten** (Whitespace-Klasse, Keyword-als-Pfad-Segment), nicht nur
  den Happy-Path.
- welle-06 bleibt **offen** (weitere Sprach-Härtung/Toolchain-Backends); slice-014 ist ihr
  erstes Inkrement.
