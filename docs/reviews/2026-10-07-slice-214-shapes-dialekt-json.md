# Review-Report: slice-214 — 2026-10-07

**Review-Art:** Code, unabhängiger Lauf. Der Kontext ist ein frischer Subagent ohne `fork`, und
er hat den Gegenstand nicht verfasst. Geprüft wird der Diff gegen Slice-Plan, ADR-0042, den
Vertrag (AC-FA-RULE-012 Boundary `json`, SPEC-EXTRACT-001 Absätze „Dialekt `json`“ und
„`literal`-Einträge“), die Hard Rules und RFC 8259 als Primärquelle. Gegen die DoD wird **nicht**
geprüft, das ist Aufgabe des Verifiers (Modul 11).

**Gegenstand:** slice-214, Commit-Range `0a7792b..7ba590a` (`7ba590a` feat(shapes): Dialekt json)

**Skill:** `.harness/skills/reviewer.md` @ `a6d19b6` (sha256 `4cc75e58…35cf4d9`) ·
**Modell:** `claude-opus-5-5` · **Datum:** 2026-10-07

> **Zitier-Form** *(dieser Block bleibt stehen. Er ist Norm, kein
> Ausfüll-Hinweis)*. Dieser Report friert ein, was er zitiert, bewegt sich
> weiter. Deshalb gilt **Kennung, nicht Adresse**: `slice-<Kennung>` statt seines
> Lifecycle-Pfads, `make <target>` statt eines Links auf die Sensor-Datei, eine
> Baseline-Stelle als **Tag + Pfad in Inline-Code** statt als Link. Ein
> `pfad`-Feld auf den **geprüften Gegenstand** ist davon nicht betroffen. Es
> zitiert den Stand des Laufs und darf ihn festhalten.

**Eingangs-Kontext:**

- Slice-Plan slice-214 (§1 mit Abgrenzung „Vertragsänderungen: Schicht-Abgrenzung“; §6 Risiko
  „Mitglieds-Grenze verrutscht“; §8 mit der Sichtung von
  `BEO-GATE/testbeschreibung-weiter-als-assertion`)
- ADR-0042 (Accepted), Entscheidung 2 („abgeleitet, nicht aufgezählt“), 4 (Einheit `json`) und 6
  (`literal`-Einträge in Quellform)
- Lastenheft 0.29.0 AC-FA-RULE-012 (Boundary `json`, Out-of-Scope); Spezifikation 0.36.0
  SPEC-EXTRACT-001 (Dialekt `json` Schritte 1–3, `literal`-Einträge), SPEC-CONF-001 (`shapes`)
- Vorlauf-Reports zu slice-212 (F-8, D-1, D-3, D-6) und slice-213 (F-2, F-12) vom selben Tag
- `BEO-GATE/testbeschreibung-weiter-als-assertion` (`observation.md`, `state.md`: geplant,
  slice-215; drei Belege)
- `AGENTS.md` §3; `harness/rules/mess-regeln.md`
- Primärquelle, im Lauf abgerufen: RFC 8259 (`https://www.rfc-editor.org/rfc/rfc8259.txt`), §2
  (Leerraum, Strukturzeichen), §3 (Literale), §4 (`member = string name-separator value`, Namen
  SHOULD unique), §6 (`int = zero / ( digit1-9 *DIGIT )`), §7 (geschlossene Escape-Menge,
  `unescaped = %x20-21 / %x23-5B / %x5D-10FFFF`), §8.1 (BOM MAY ignore)

**Sonden, alle in einem Klon außerhalb des Repos** (`git archive HEAD` ins Scratchpad, eigene
Image-Tags; das Repo blieb unberührt, bis auf diesen Report):

- `make test` im Repo auf `7ba590a`: Exit 0.
- **Reale `package.json` (Auftrag iii).** Alle Dateien `package.json` unter `/Development` ohne
  `node_modules`-Pfade, gefunden mit `find`: **233**. Zähler 1: `normalizeJSON` — 231 zerlegt,
  **2** Exit 2 („die Wurzel ist kein Objekt“). Zähler 2, anders gebaut: Go-`encoding/json` — 233
  gültig; die 2 Exit-2-Dateien haben eine **Zeichenketten-Wurzel** („These examples have been moved
  to https://github.com/grpc/g…“), also vertragsgemäß Exit 2; bei allen 231 Objekt-Wurzeln stimmt
  die Mitgliederzahl von `encoding/json` mit der Anweisungszahl überein (keine Abweichung
  gemeldet). Geltungsbereich: Zerlegbarkeit und Mitglieds-Grenze an realen Dateien; ob ein
  Mitglied byte-genau richtig geschnitten ist, deckt der Zähler-Vergleich nur über die Anzahl.
  Die Zahlen decken sich mit der Commit-Message (233/231/2).
- **Lexik- und Zerlegungsproben** gegen `normalizeJSON` (32 Eingaben, Ergebnis je Probe in den
  Abschnitten unten) und **End-to-End** über `cli.Run` (7 Dateien).
- **Drei Mutationen** an `json_shape.go`, je mit voller Suite (`-skip Probe`, also nur die
  Bestands-Tests) und einer Gegenprobe; Ergebnis in F-2.

---

## Findings

| ID | Kategorie | Befund | Quelle | Pfad | Verifizierbar | Klasse |
|---|---|---|---|---|---|---|
| F-1 | HIGH | Der CHANGELOG sagt allgemein: „Kommentare, JSON5-Schreibweisen, eine Wurzel ohne Objekt und nur durch Leerraum getrennte Werte sind Exit 2.“ Für vier JSON5-Schreibweisen stimmt das nicht, End-to-End gemessen: abschließendes Komma in einem Wert-Objekt `{"name":"web","dependencies":{"react":"^18.2.0",}}`, `+1`, `.5` und ein unquotierter Schlüssel aus Atom-Buchstaben `{"s":{nul:1}}` ergeben **Exit 1** (`shape-unlisted`), nicht Exit 2. Gegenprobe im selben Lauf: `// c` ergibt Exit 2 („Zeichen '/' … nicht zulässig“). Die Wirkung ist fail-safe (Befund statt Grün), die Zusage aber falsch. | Skill §HIGH „nachweislich falsche Tatsachenbehauptung“; SPEC-EXTRACT-001 `json` Schritt 1/3 (nennt JSON5 nicht) | `CHANGELOG.md:17-18` | nein — kein Gate liest Prosa gegen das Verhalten | Doku-Zusage weiter als das Verhalten |
| F-2 | MEDIUM | Der Kommentar zu `TestJSONUnsplittable` sagt „Alle fail-closed-Fälle der Spezifikation“. Drei Fälle des Vertrags prüft der Test nicht, je durch eine Mutation belegt, unter der die **ganze Bestands-Suite grün** bleibt (`ok` für alle sechs Pakete) und die Gegenprobe rot wird: (M1) `colon = colon \|\| depth == 0` → `colon = true` — „kein `:` auf Tiefe 1“; Gegenprobe `{"a"{"b":1}}` rot mit „ROT … kein Fehler“, unmutiert „ein Mitglied trägt kein `:`“. (M2) `case c < 0x20` → `case c == '\n' \|\| c == '\t'` — Steuerzeichen U+0000–U+001F außer LF/TAB; Gegenprobe `"x\x01y"` rot. (M3) Leerraum-Prüfung nur noch zwischen Tokens **gleicher** Art — Leerraum zwischen Zeichenkette und Zahl; Gegenproben `{"a":1 "b":2}` und `{"a":"x" 1}` rot. Unter M3 fallen `{"a":1 "b":2}` und `{"a":1"b":2}` auf dieselbe Normalform `"a":1"b":2`: die Eindeutigkeit, die die Spezifikation mit dieser Regel sichert, hat keinen Test, der sie hält. | `BEO-GATE/testbeschreibung-weiter-als-assertion`; SPEC-EXTRACT-001 `json` Schritt 1 und 3 | `internal/adapter/driven/extract/json_shape_test.go:39`; Gegenstand `internal/adapter/driven/extract/json_shape.go:68,116,215` | ja — die drei Mutationen in `make test` | Testbeschreibung sagt mehr zu als die Assertion |
| F-3 | MEDIUM | Zwei Zusagen über die Lexik reichen weiter als der Code. Handbuch: „Die Lexik folgt RFC 8259“. Code-Kommentar am Datei-Kopf: „…or the input is no JSON at all, the file is unsplittable (exit 2)“. Gemessen werden ohne Fehler zerlegt: `tru`, `eee`, `nul`, `truefalse`, `01`, `--1`, `1.` (RFC 8259 §3, §6), `"\x"` und Backslash vor einem Steuerzeichen (§7, geschlossene Escape-Menge), `{"a":1"b":2}` als **ein** Mitglied (§4, fehlender `value-separator`), `[1,,2]`, `[,]`, `{1}` in einem Wert. Die Spezifikation verspricht das nicht (sie lässt die Zeichen zu, nicht die Tokens); der Handbuch-Satz und der Kommentar tun es. | SPEC-EXTRACT-001 `json` Schritt 1; RFC 8259 §3, §4, §6, §7; `AGENTS.md` §3.7 (Kommentar trägt eine Zusage, die gelten muss) | `docs/user/benutzerhandbuch.md:659`; `internal/adapter/driven/extract/json_shape.go:12-14` | ja — Eingaben gegen `normalizeJSON` | Doku-Zusage weiter als das Verhalten |
| F-4 | LOW | Ein von einem Backslash maskiertes `LF` in einer Zeichenkette zählt der Lexer nicht als Zeile (`jsonString` überspringt das Zeichen mit `l.i++`, ohne `lineEnd`). Die Spezifikation definiert `LF` als Zeilenende für die Zeilennummer und lässt die Maskierung zu („ein Backslash maskiert das nächste Zeichen“). Gemessen: `{"a":"x\` LF `y",` LF `"b":1}` meldet `"b"` in Zeile 2, es steht in Zeile 3; die Zeilenzahl (`countLines`) zählt richtig 3. Die Eingabe ist kein gültiges JSON (RFC 8259 §7), die Wirkung betrifft nur den Ort einer Meldung. | SPEC-EXTRACT-001 `json` Schritt 1 und 3 („Original-Zeile“) | `internal/adapter/driven/extract/json_shape.go:111-112` | ja — Eingabe gegen `normalizeJSON` | Code weicht vom Vertrag in der Zeilenzählung ab |
| F-5 | LOW | Wo RFC 8259 eine Form verbietet, lässt der Vertrag sie zu, ohne die Grenze zu benennen (siehe die Liste in F-3; dazu abschließende Kommas in verschachtelten Objekten, JSON5-Zahlen aus F-1). ADR-0042 Entscheidung 2 sagt: wo die Quelle eine Form **verbietet**, fail-closed. Der Code folgt der Spezifikation, nicht der ADR-Formulierung; eine Normalform-Kollision entsteht dabei nicht (Probe-Abschnitt ii), alle Fälle sind fail-safe. Der Vorlauf-Report zu slice-212 (F-8, D-6) hatte die Klasse für Leerraum benannt; dieser Rest ist weder als Grenze (AC-QA-02) noch als Abweichung genannt. Kein Code-Defekt dieses Slice — slice-214 §1 schließt Vertragsänderungen aus; die Frage gehört an den Architect. | ADR-0042 Entscheidung 2; SPEC-EXTRACT-001 `json`; RFC 8259 §4–§7; AC-QA-02 | `spec/spezifikation.md:362-383`; `docs/plan/adr/0042-shapes-benannte-dialekte-gomod-json.md:63-70` | ja — Eingaben gegen `normalizeJSON` | Vertrag lässt eine von der Quelle verbotene Form zu, Grenze nicht benannt |
| F-6 | LOW | Das Handbuch zählt `#` zu den „JSON5-Schreibweisen (`//`, `#`, `'`)“. JSON5 kennt keine `#`-Kommentare (nur `//` und `/* */`); `/*` fehlt in der Aufzählung. Das Verhalten ist für alle vier Zeichen Exit 2. | Wording | `docs/user/benutzerhandbuch.md:678` | nein | Wording |
| F-7 | INFO | Der Kommentar zu `TestJSONUnsplittable` sagt „Exit 2“; die Assertion prüft nur, dass `normalizeJSON` einen Fehler liefert. Der Weg Fehler → Exit 2 ist in `TestShapesJSONEndToEnd` für die Wurzel ohne Objekt belegt und gilt für jeden Normalisierer gleich. | `BEO-GATE/testbeschreibung-weiter-als-assertion` | `internal/adapter/driven/extract/json_shape_test.go:39` | nein | Testbeschreibung sagt mehr zu als die Assertion |
| F-8 | INFO | Ungültiges UTF-8 in einer Zeichenkette (`"\xff"`) wird angenommen. Die Spezifikation setzt UTF-8 voraus, nennt ungültige Bytes aber nicht als Fehler; die Zeichenkette bleibt byte-genau, eine Kollision entsteht nicht. | SPEC-EXTRACT-001 `json` Schritt 1; RFC 8259 §8.1 | `internal/adapter/driven/extract/json_shape.go:107-121` | ja — Eingabe gegen `normalizeJSON` | Vertrag lässt eine von der Quelle verbotene Form zu, Grenze nicht benannt |
| F-9 | INFO | Ergebnis von Auftrag (i): Zwei Mitglieder ohne trennendes Komma (`{"name":"web""scripts":{"postinstall":"x"}}`) werden **eine** Anweisung `"name":"web""scripts":{…}`, End-to-End Exit 1. Gegen `literal`-Einträge und gegen die dokumentierte Regex-Form `"[^"$\\]*"` (voll verankert) ist das fail-safe; ein Regex `"name":.*` ließe die zusammengezogene Anweisung durch — das liegt in der benannten Regex-Grenze von SPEC-CONF-001, und ein JSON-Parser (npm) lehnt die Datei ab. | SPEC-CONF-001 Grenze der Regex-Einträge; AC-QA-02 | `internal/adapter/driven/extract/json_shape.go:157-170` | ja — Eingabe gegen das CLI | Mitglieds-Grenze bei ungültigem JSON verschoben, fail-safe |

## Adversariale Konstruktion (Auftrag 2a)

- **(i) Zusätzliche Abhängigkeit oder zusätzliches Mitglied als erlaubt durchgelassen:** kein
  Fall mit `literal`-Einträgen und der dokumentierten Regex-Form gefunden.
  - `TestShapesJSONEndToEnd` vergleicht die ganze Ausgabe für zusätzliche Abhängigkeit und
    zusätzliches `scripts`-Mitglied (je genau eine Zeile, Zeile 5).
  - Doppelte Schlüssel `{"a":1,"a":2}` ergeben zwei Anweisungen; jede muss für sich erlaubt sein.
    Ein doppeltes `dependencies` mit unerlaubtem erstem Wert wird gemeldet.
  - Ein zusammengezogenes Mitglied ohne Komma: F-9.
  - `literal`-Einträge, die mehr als ein Mitglied einschmuggeln wollen: `"a":1,"b":2` → „ergibt 2
    Anweisungen“ (Exit 2, `shapes.go:52`); `"a":1}{"b":2`, `"a":1} , {"b":2` → „Inhalt nach dem
    Ende des Wurzel-Objekts“; `"a":{` → „nicht geschlossen“; `"a":1} //` → Zeichen `/`; ` ` →
    null Anweisungen → Exit 2. `{"private": true}` als Eintrag → Exit 2 (E2E-Test, gemessen).
- **(ii) Zwei verschiedene Quellen, eine Normalform:** kein Fall gefunden außer dem gewollten BOM
  (§8.1 MAY ignore; ein zweiter BOM ist Exit 2, gemessen). Begründung: Leerraum fällt nur neben
  einem Strukturzeichen weg, und Strukturzeichen sind Ein-Zeichen-Tokens; zwischen zwei
  Nicht-Strukturzeichen-Tokens ist er Exit 2 (gemessen für Atom/Atom, String/String,
  Zahl/String, String/Zahl). Zeichenketten bleiben byte-genau, Atome wörtlich. `eee` und `tru`
  sind Atome (Zeichen aus der Menge), keine Kollision: Sie bleiben als eigener Text stehen und
  sind nur ungültiges JSON (F-3, F-5). Die Eindeutigkeit hängt an einer Bedingung, die kein Test
  hält (F-2, M3).
- **(iii) Reale `package.json` fälschlich Exit 2:** 0 von 233. Die 2 Exit-2-Dateien haben eine
  Zeichenketten-Wurzel, Exit 2 ist dort vertragsgemäß (Messung oben).
- **(iv) Code weicht von der Spezifikation ab:** F-4 (Zeilennummer nach maskiertem `LF`). Schritt
  für Schritt sonst deckungsgleich: Backslash vor Steuerzeichen maskiert (Spec-Wortlaut, RFC
  verbietet es: F-5); Atom-Zeichen genau `0-9 + - . e E` und die Buchstaben `t r u f a l s n` —
  `e` und `E` überdecken sich mit `true`/`false`, die Menge ist richtig; Duplikat-Schlüssel ohne
  Sonderregel (Spec schweigt, RFC: SHOULD); BOM nur am Anfang; Zeilenzählung LF/CRLF/CR über
  `lineEnd` und `countLines` (`TestJSONLines` mit CRLF und einzelnem CR); `{}` und ` { } ` → null
  Anweisungen.

## Negativbefunde

| Bereich | Ergebnis |
|---|---|
| Registry und `literalSource` (`shapes.go`) | geprüft, ohne Befund. Ein Eintrag in `dialects()`; Einschluss nur bei `json` (Test prüft `gomod`/`kotlin` unverändert); „genau eine Anweisung“ bleibt in `shapeLiterals` und gilt auch für `json`. |
| `jsonMembers`/`take`/`jsonStructure` | geprüft, ohne Befund. Inhalt nach der Wurzel, unpaarige und falsch gepaarte Klammern, leeres und abschließendes Mitglied sind Exit 2 und je getestet; Original-Zeile = Zeile des Schlüssel-Tokens (`TestJSONLines` würde eine Verwechslung mit der Trenner-Zeile fangen: Mitglied `"b"` in Zeile 4, `}` in Zeile 6). |
| `TestJSONRealForm`, `TestJSONStringsAndNesting`, `TestJSONLines`, `TestJSONLiteralSource`, `TestShapesJSONEndToEnd` — Kommentar gegen Assertion | geprüft, ohne Befund. Jede Zusage des Kommentars hat eine Assertion; das E2E vergleicht die Ausgabe vollständig. Einzige Lücke: F-2, F-7. |
| Leerraum-Menge RFC 8259 §2 | geprüft, ohne Befund. Seitenvorschub ist Exit 2 und getestet. |
| `--print-config`-Gerüst (`internal/cli/cli.go`) | geprüft, ohne Befund. Beispiel deckt sich mit Handbuch und Vertrag (Einträge ohne äußere Klammern). |
| Handbuch §4 übrige Aussagen | geprüft, ohne Befund. Beispiel-Konfiguration = Konfiguration des E2E-Tests (grün gemessen); „nach dem Entmaskieren (`\\` → `\`) übernehmbar“ deckt sich mit SPEC-EXTRACT-001 und dem Vorlauf-Befund D-1 zu slice-212; Exit-2-Fall `{ }`-Eintrag belegt. Historie-Zeile 1.45 und Kopf stimmen überein. |
| CHANGELOG übrige Aussagen | geprüft, ohne Befund außer F-1. Der Vermerk „noch nicht implementiert“ ist entfernt. |
| Hard Rules `AGENTS.md` §3.1–§3.6 | geprüft, ohne Befund. Kein `//nolint`; kein Spec-Stratum berührt; keine ADR geändert; kein Move im Commit; keine Gate-Schwelle berührt. |
| `AGENTS.md` §3.7 in Code-Kommentaren | geprüft; keine Chronik, keine verworfene Alternative. Einzige zu weite Zusage: F-3. |
| Traceability der Commit-Message | geprüft, ohne Befund. `AC-FA-RULE-012`, `ADR-0042`, `slice-214`. |
| Slice-§1-Abgrenzung | geprüft, ohne Befund. Ausgabe-Regel und Vertrag unberührt. |

## Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 1 |
| MEDIUM | 2 |
| LOW | 3 |
| INFO | 3 |

**Finding-Klassen dieses Laufs:** Doku-Zusage weiter als das Verhalten · Testbeschreibung sagt
mehr zu als die Assertion · Code weicht vom Vertrag in der Zeilenzählung ab · Vertrag lässt eine
von der Quelle verbotene Form zu, Grenze nicht benannt · Mitglieds-Grenze bei ungültigem JSON
verschoben, fail-safe · Wording

## Verdikt

**Merge-blockierend:** ja — F-1 (HIGH) und F-2, F-3 (MEDIUM). Am Code selbst blockiert keines
dieser Findings ein Verhalten: Lexik und Zerlegung sind spec-treu bis auf F-4, keine Probe ließ
eine Anweisung als erlaubt durch, und keine reale Datei bekam fälschlich Exit 2. Blockierend sind
die Zusagen in CHANGELOG, Handbuch und Kommentaren und der Testkommentar, dessen Lücken drei
Mutationen grün überstehen. F-5 ist eine Vertragsfrage an den Architect (slice-214 §1 schließt
Vertragsänderungen aus) und blockiert diesen Slice nicht.

**Übergabe:** Findings gehen an den Implementer; F-5 zusätzlich als Frage an den Architect. Die
Finding-Klassen gehen in die Slice-Closure §7. F-2 und F-7 sind ein weiteres Auftreten von
`BEO-GATE/testbeschreibung-weiter-als-assertion` (Stand: geplant, slice-215). Dieser Report
ersetzt keine Verifikation.
