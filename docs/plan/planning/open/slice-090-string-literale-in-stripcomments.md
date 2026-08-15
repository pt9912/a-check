# slice-090 — `stripComments`: ein `/*` im String verschluckt den Rest der Datei

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Deckt:** `R-1` (HIGH) aus dem [Review-Report vom 2026-08-15](../../../reviews/2026-08-15-v0170-go-kern.md);
[AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze),
[AC-FA-EXTRACT-001](../../../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion).
**Bezug:** dieselbe False-Green-Klasse, gegen die Python in
[slice-020](../done/slice-020-python-backend.md) ausgenommen wurde — für die übrigen sieben
Backends blieb sie offen.

---

## 0. Trigger

**Beginn: sofort.** Der Fehler ist in [`v0.17.0`](../../../../version.md#aktuell) **ausgeliefert**
und erzeugt stille Falsch-Negative bei jedem Konsumenten, dessen Quellcode Glob-Strings oder URLs
enthält.

**Rückführungen:**

- `in-progress` → `open`: falls die Messung zeigt, dass ein zeichenketten-bewusster Scanner ohne
  echten Lexer nicht rauschfrei zu haben ist — dann wäre der Fix selbst eine Heuristik über einer
  Heuristik und braucht erst einen Entscheid.

## 1. Auslöser

**Mechanismus: der Kommentar-Entferner kennt keine Zeichenketten.**
[`stripComments`](../../../../internal/adapter/driven/extract/extract.go) liest rohe Bytes. Jede
Byte-Folge `/*` **in einem String-Literal** öffnet einen Phantom-Blockkommentar, der bis zum
nächsten `*/` konsumiert wird — oft bis Dateiende. Alles dazwischen ist für jede Prüfung
unsichtbar: Importe, `forbidden_constructs`, `constructs`, Kotlin-Deklarationen.

**Gemessen (2026-08-15), beide Richtungen:**

```text
Go-Fixture, Verstoss core -> adapters NACH  var globs = []string{"/**"}
  mit dem String    -> gesamt: 0 Befund(e)                     EXIT=0
  ohne den String   -> core-impurity: Kern importiert …        EXIT=1

TypeScript, tech-Regel auf "deno.land", Import im Kern
  mit    // in der URL -> gesamt: 0 Befund(e)                   EXIT=0
  ohne   // in der URL -> core-impurity                         EXIT=1
```

**Die Klasse ist bekannt — die Reichweite war es nicht.** `prepSource` nimmt Python ausdrücklich
aus, mit exakt dieser Begründung: *„a `/*`-like byte sequence inside a Python string literal (e.g.
the glob `**/*.py`) would otherwise swallow every real import up to the next `*/` — a silent
false-green."* Der Schluss, dass dasselbe für **alle** C-Syntax-Backends gilt, wurde nie gezogen.

**Betroffen sind sieben Backends:** Go, C++, Rust, Kotlin, Java, C#, TypeScript. Glob-Strings sind
in Werkzeug-Code alltäglich, `https://`-URLs in TypeScript/Deno die Norm.

**Selbstbetroffenheit, ehrlich:** `extract.go` trägt selbst einen `"/**"`-String; beim Selbstscan
werden rund 340 Zeilen geblankt. **Realer Schaden gering** — die Go-Importe stehen oberhalb, und
a-check nutzt keinen `constructs`-Block. Geschwächt ist die Dogfooding-**Aussage**, nicht die
Prüfung.

## 2. Betroffene Module

Zwei Schichten:

1. **Spec** — eine ADR nach dem Muster von
   [ADR-0028](../../adr/0028-ziel-glob-schattenwurf.md) (Verhaltensänderung, die bestehende grüne
   Gates rot machen kann) und [`spec/spezifikation.md`](../../../../spec/spezifikation.md)
   ([SPEC-EXTRACT-001](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion): die Vorbereitung der Quelle).
2. **[`internal/adapter/driven/extract/`](../../../../internal/adapter/driven/extract/extract.go)** —
   `stripComments` und die Python-Ausnahme in `prepSource`.

## 3. Auszuführende Gates

`make gates`, `make image-test`.

**Der Entscheid, der vor dem Bau fällt: wie weit geht der Zeichenketten-Automat?** Ein vollständiger
Lexer je Sprache ist nach [ADR-0002](../../adr/0002-text-heuristische-extraktion.md) Out-of-Scope —
die Frage ist, welche Literal-Formen der Automat kennt und welche als ausgewiesene Grenze bleiben.

| Umfang | Deckt | Bleibt offen |
|---|---|---|
| **(a) `"` und `'` mit Backslash-Escape** | den belegten Fall (Glob-Strings, URLs) in allen sieben Backends | Go-Backtick-Raw-Strings, C++ `R"(…)"`, Rust `r#"…"#`, Java/Kotlin/C# Text-Blöcke, TS-Templates |
| **(b) (a) + Backtick** | zusätzlich Go-Raw-Strings und TS-Template-Literale | die übrigen Raw-Formen |
| **(c) je Sprache eigene Literal-Regeln** | nahezu alles | Aufwand; kippt Richtung Lexer, gegen [ADR-0002](../../adr/0002-text-heuristische-extraktion.md) |

**Neigung des Autors: (b).** Backtick-Strings sind in Go alltäglich (genau dort, wo Globs und
Regexe stehen) und in TypeScript die übliche Interpolationsform; sie kosten denselben
Zustandsautomaten. Die verbleibenden Raw-Formen sind selten genug, um sie als Grenze auszuweisen —
und der Automat darf **nie** mehr verschlucken als heute: im Zweifel lieber ein Kommentar zu wenig
entfernt (Falsch-Positiv, sichtbar) als eine Zeile zu viel (Falsch-Negativ, still).

**Das ist eine Verhaltensänderung mit Vertragsbezug.** Bei Konsumenten können Importe sichtbar
werden, die es vorher nicht waren — grüne Gates werden rot. Nach
[`AGENTS.md` §3.6](../../../../AGENTS.md#36-gates-dürfen-nicht-ohne-adr-gelockert-werden) in der
Gegenrichtung (Argumentation aus [ADR-0029](../../adr/0029-abdeckungs-diagnose-advisory.md))
braucht das eine ADR — auch wenn es sachlich eine Fehlerbehebung ist.

**Negativ-Proben:**

| Probe | Erwartung |
|---|---|
| Go-Datei mit `"/**"`-String, Verstoß danach | Befund, Exit 1 |
| TypeScript mit `https://`-URL-Import und `tech`-Regel darauf | `tech-leak`, Exit 1 |
| **echter** Blockkommentar mit Import darin | weiterhin **kein** Befund |
| **echter** Zeilenkommentar mit Import darin | weiterhin **kein** Befund |
| Python-Fixture aus `TestPythonGlobStringKeepsImports` | unverändert grün |
| `make arch-check` (Dogfooding) | unverändert **0** Befunde |

## 4. Was bewusst nicht getan wird

- **Einen Lexer bauen.** [ADR-0002](../../adr/0002-text-heuristische-extraktion.md) legt die
  text-heuristische Extraktion als Fundament fest. Dieser Slice schließt eine **Lücke im
  Kommentar-Strip**, er ersetzt die Heuristik nicht.
- **Die Python-Ausnahme aufheben.** Sie ist eine zweite, unabhängige Absicherung; sie zu entfernen,
  weil der Automat jetzt greift, tauscht einen belegten Schutz gegen eine Annahme.
- **Die übrigen sechs Findings.** Eigene Schnitte — `R-2` ist der nächste.

## 5. DoD

- [ ] Spec-first: die Verhaltensänderung steht in einer ADR mit `Status: Accepted` und in
      [SPEC-EXTRACT-001](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion),
      bevor der Code committet wird.
- [ ] Ein `/*` oder `//` **in einem Zeichenketten-Literal** verschluckt nichts mehr; echte
      Kommentare werden weiterhin entfernt. Beleg: die sechs Proben aus §3, davon vier vor dem Bau
      rot bzw. falsch-grün.
- [ ] `make gates` und `make image-test` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft.

## 6. Closure-Notiz

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice,
[`AGENTS.md`](../../../../AGENTS.md) §5; `make verify` prüft das.)_

## 7. Sub-Area-Modus

Alle berührten Sub-Areas GF.
