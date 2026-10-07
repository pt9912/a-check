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

---

## Delta-Review (Nachlauf 8868422..f778348)

**Review-Art:** Code und Doku, unabhängiger Lauf, frischer Subagent ohne `fork`; der Kontext hat
weder den Gegenstand noch den Erst-Review verfasst. Geprüft gegen Slice-Plan slice-214 (mit der
Plan-Änderung), ADR-0042 Entscheidung 2, SPEC-EXTRACT-001 „Dialekt `json`“ (0.38.0),
AC-FA-RULE-012 (Lastenheft 0.29.0, unverändert), `AGENTS.md` §3 und RFC 8259 als Primärquelle.
Nicht gegen die DoD.

**Gegenstand:** slice-214, Commit-Range `8868422..f778348` — `f25a894` (planning, Plan-Änderung),
`268b185` (Spezifikation 0.38.0), `f778348` (Code, Tests, Handbuch, CHANGELOG).

**Skill:** `.harness/skills/reviewer.md` @ `a6d19b6` (sha256 `4cc75e58…35cf4d9`) ·
**Modell:** `claude-opus-5-5` · **Datum:** 2026-10-07

**Sonden.** `make test` im Repo auf `f778348`: Exit 0. `make doc-check`: Exit 0, 694 Dateien,
0 Befunde. Alle übrigen Sonden in Klonen außerhalb des Repos (`git archive f778348` ins
Scratchpad, eigene Image-Tags, Lauf je über `make test`; kein Host-Go):

- **Einzelproben** gegen `normalizeJSON`: 33 RFC-gültige und 57 RFC-ungültige Eingaben, je mit
  `encoding/json` als zweitem, anders gebautem Zähler (gültig ⇔ `json.Valid` und Objekt-Wurzel).
- **Differenzielles Fuzzing** gegen denselben Referenz-Zähler: 3 000 000 Zufalls-Eingaben (ASCII,
  Länge 1–14, Alphabet aus Strukturzeichen, Atom-Zeichen, `"`, `\`, `u`, `x`, Leerraum, U+0001;
  380 angenommen) und 1 000 000 Mutationen (1–3 Einfüge-/Lösch-/Ersetz-Schritte) eines gültigen
  Objekts mit Zahlen, Literalen, Escapes und Verschachtelung (150 699 angenommen): **0
  Abweichungen**. Geltungsbereich: ASCII-Eingaben — ungültiges UTF-8 liegt außerhalb, weil
  `encoding/json` es annimmt (dort gesondert gemessen, unten); Tiefe und Länge sind durch die
  Generatoren begrenzt.
- **Korpus:** dieselben 233 realen `package.json` wie im Erst-Review: 231 zerlegt, 2 Exit 2
  (Zeichenketten-Wurzel), 0 Abweichungen gegen `encoding/json` in Annahme **und**
  Mitgliederzahl. Der strengere Prüfer kostet an realen Dateien nichts.
- **End-to-End** über `cli.Run`: 9 Dateien, 4 `literal`-Einträge.
- **Mutationen** an `json_shape.go`, je mit der vollen Bestands-Suite: die drei aus F-2 (M1–M3),
  drei am neuen Prüfer (M4 Prüfer-Aufruf entfernt, M5 UTF-8-Prüfung aus, M6 Escape-Prüfung aus)
  und drei an seinen Grenzen (V1 Exponent `[+-]*`, V2 `\b` aus der Escape-Menge, V5 `-*` vor der
  Zahl); für die überlebenden je eine Gegenprobe, unmutiert grün, mutiert rot mit Meldung.
- **Tiefe:** Arrays der Tiefe 10^5, 10^6, 3·10^6 im Wert eines Mitglieds.

### Stand der Erst-Findings

| ID | Stand | Messung |
|---|---|---|
| F-1 | erledigt | CHANGELOG sagt jetzt „gültiges JSON nach RFC 8259 … jede andere Abweichung sind Exit 2“. End-to-End: abschließendes Komma in `dependencies`, `+1`, `.5`, `{"s":{nul:1}}` je **Exit 2** (vorher Exit 1). |
| F-2 | erledigt | Kommentar auf „Fehlerfälle der Lexik und der Zerlegung (Schritt 1 und 4)“ verengt. M2 (Steuerzeichen) ist jetzt **rot**: `TestJSONValidity` „Steuerzeichen U+0001 … Fehler erwartet“. M1 (`:` auf Tiefe 1) und M3 (Leerraum nur zwischen gleichen Token-Arten) überleben die Suite weiter, sind aber **äquivalente Mutanten**: Unter beiden lief das Fuzzing mit 0 Abweichungen, weil der Grammatik-Prüfer dieselben Fälle früher abweist. Die Eindeutigkeit der Normalform hängt damit an zwei Schichten, von denen jede für sich getestet ist (M4 rot in 17 Fällen plus `TestJSONLiteralSource`). Siehe D-6. |
| F-3 | erledigt | Alle Formen aus F-3 sind Exit 2, je mit eigener Meldung: `tru`, `eee`, `nulll`, `truefalse`, `01`, `--1`, `1.`, `"\x"`, Backslash vor Steuerzeichen, `{"a":1"b":2}`, `[1,,2]`, `[,]`, `{1}` im Wert. Handbuch-Satz und Kopf-Kommentar sagen jetzt „gültiges JSON“; das Fuzzing deckt die Zusage im genannten Geltungsbereich. |
| F-4 | erledigt (Gegenstand entfallen) | Ein maskiertes `LF` in einer Zeichenkette ist jetzt eine ungültige Escape-Folge, Exit 2 (gemessen mit der F-4-Eingabe). Eine falsche Zeilennummer kann nur noch an Eingaben entstehen, die nicht zerlegt werden. |
| F-5 | erledigt | Spezifikation 0.38.0 Schritt 2 verlangt gültiges JSON; damit folgt der Vertrag ADR-0042 Entscheidung 2. Lastenheft unverändert und widerspruchsfrei (die Boundary `json` nennt nur die Objekt-Wurzel). Restfehler im neuen Satz: D-2. |
| F-6 | erledigt | Handbuch zählt `//`, `/* */`, `#` als Kommentare, JSON5 getrennt. `// c` und `/*` gemessen Exit 2. |
| F-7 | erledigt | Kommentar nennt „im Scan Exit 2“ als Folge in Klammern; die Assertion prüft den Fehler, der Weg zu Exit 2 ist End-to-End belegt (auch für die neuen Fälle, oben). |
| F-8 | erledigt | `"\xff"`, abgeschnittene Folge `\xc3`, überlange Kodierung `\xc0\xaf`, UTF-8-kodiertes Surrogat `\xed\xa0\x80`: je Exit 2 („kein gültiges UTF-8“). `encoding/json` nimmt alle vier an; der Prüfer ist hier strenger als die Referenz und folgt RFC 8259 §8.1. |
| F-9 | erledigt | `{"name":"web""scripts":{"postinstall":"x"}}` End-to-End **Exit 2** („erwartet "}" oder ",", gefunden "scripts"“), vorher Exit 1. |

### Neue Findings

| ID | Kategorie | Befund | Quelle | Pfad | Verifizierbar | Klasse |
|---|---|---|---|---|---|---|
| D-1 | HIGH | Der Kopf des Slice-Plans sagt zu SPEC-EXTRACT-001 „— umgesetzt, nicht geändert“. `268b185` ändert genau diesen Absatz (Spezifikation 0.37.0 → 0.38.0, Schritt 2 „Gültigkeit“ neu); die Plan-Änderung in §1 nennt das, das Kopffeld blieb stehen. Ebenso steht der Ausschluss „Vertragsänderungen … der Vertrag steht“ unverändert neben der Plan-Änderung, die eine Vertragsänderung aufnimmt. Adversarisch gegen `git diff 8868422..f778348 -- spec/` geprüft. | Skill §HIGH „nachweislich falsche Tatsachenbehauptung“; `modul-05` §Ziel-Form: Slice (Kopf nennt die berührten Spec-Stellen) | `docs/plan/planning/in-progress/slice-214-shapes-dialekt-json.md:14`, `:45-46` | nein — kein Gate liest das Kopffeld gegen den Diff | Zustandsangabe im Plan nach Plan-Änderung nicht nachgezogen |
| D-2 | HIGH | Schritt 2 sagt: „Die Fehlerliste unter 1 und 4 sind Fälle dieser Regel“ (der Regel „gültiges JSON nach RFC 8259“). Für den ersten Fehler unter 4, „eine Wurzel, die kein Objekt ist“, stimmt das nicht: RFC 8259 §2 lässt jeden Wert als Wurzel zu (`JSON-text = ws value ws`). Der Code sagt es selbst richtig — `jsonValidate` „accepts any root value — that the root is an object is the split's rule“ —, ebenso das Handbuch („Ebenso Exit 2, wenn die Wurzel kein Objekt ist“). Das Verhalten (Exit 2) ist in Schritt 4 eigens festgelegt; falsch ist die Zuordnung, nicht die Folge. | Skill §HIGH „nachweislich falsche Tatsachenbehauptung“; RFC 8259 §2 | `spec/spezifikation.md:379-380` | nein | Vertrag ordnet eine eigene Regel der Quelle zu |
| D-3 | MEDIUM | Der Kommentar von `TestJSONValidity` sagt „jede Form, die die Quelle verbietet, ist ein Fehler“ und, vor der letzten Assertion, „gültige Grenzfälle bleiben gültig“. Drei Mutationen überleben die ganze Suite (`ok` für alle sechs Pakete): V1 Exponent `[+-]*` nimmt `{"a":1e+-2}` an, V5 `-*` nimmt `{"a":--1}` an, V2 ohne `b` in der Escape-Menge lehnt das gültige `{"a":"\b"}` ab. Gegenprobe unmutiert grün, mutiert je rot („ROT ungueltig angenommen: {"a":1e+-2}“, „… {"a":--1}“, „ROT gueltig abgelehnt: … ungültige Escape-Folge“). Das Verhalten des Codes ist richtig (Fuzzing oben); die Beschreibung des Tests reicht weiter als seine 20 Fälle und den einen Gültig-Fall. Die Zusage „Je Fall eine Zeile“ allein stimmt. | `BEO-GATE/testbeschreibung-weiter-als-assertion`; SPEC-EXTRACT-001 `json` Schritt 2 | `internal/adapter/driven/extract/json_shape_test.go:97-99`, `:127`; Gegenstand `internal/adapter/driven/extract/json_shape.go:237`, `:403` | ja — die drei Mutationen in `make test` | Testbeschreibung sagt mehr zu als die Assertion |
| D-4 | LOW | Der neue Prüfer steigt rekursiv ab (`value` → `array`/`object`); die Zerlegung davor hielt die Tiefe in einem Slice. Ein gültiges Array der Tiefe 3·10^6 (rund 6 MB) beendet den Prozess mit `fatal error: stack overflow` („goroutine stack exceeds 1000000000-byte limit“) statt mit einer Meldung; Tiefe 10^6 läuft in 0,69 s durch. RFC 8259 §9 erlaubt eine Grenze der Schachtelungstiefe; Spezifikation und Handbuch nennen keine. Der Exit-Code eines Go-Laufzeit-Abbruchs ist 2, über `cli.Run` nicht gemessen — fail-closed, aber nicht über den Vertragsweg. | AC-QA-02 (ehrliche Grenze); RFC 8259 §9; SPEC-EXTRACT-001 `json` Schritt 2 | `internal/adapter/driven/extract/json_shape.go:278-345` | ja — Eingabe gegen `normalizeJSON` | Implementierungsgrenze nicht benannt |
| D-5 | LOW | Der Kommentar von `TestJSONValidity` begründet seinen Umfang mit dem Gegenfall im Konjunktiv: „auch die, die sonst fail-safe in eine Anweisung fielen“. Nicht HIGH, weil der Kommentar daneben eine Zusage trägt; der Nebensatz beschreibt, was ohne die Prüfung geschähe. | `AGENTS.md` §3.7 | `internal/adapter/driven/extract/json_shape_test.go:98-99` | nein | Kommentar beschreibt die verworfene Alternative |
| D-6 | INFO | Hinter `jsonValidate` sind die Prüfungen „leeres Mitglied“, „Mitglied beginnt nicht mit einer Zeichenkette“, „kein `:` auf Tiefe 1“ (`jsonMember`) und die Leerraum-Regel (`jsonTokens`) für jede Eingabe redundant: M1 und M3 sind äquivalente Mutanten (oben). Die Kommentare dort bleiben wahr. Für die DoD-Behauptung „Mutations-Gegenprobe“ ist das die Frage des Verifiers, welche Schicht ein Fall erreicht. | `BEO-GATE/probe-liefert-den-gegenstand-mit` (Hinweis, kein Auftreten) | `internal/adapter/driven/extract/json_shape.go:73`, `:203-230` | ja — M1/M3 mit Fuzzing | Prüfung nach vorgelagertem Prüfer unerreichbar |

### Adversariale Konstruktion

- **Von RFC 8259 verbotene Form, die noch angenommen wird:** keine gefunden. Einzelproben (57,
  alle Exit 2): Zahlen `-`, `-01`, `1.e2`, `1e+`, `1e5.5`, `-.5`, `1.5e`, `1-2`, `0x10`; Literale
  `TRUE`, `NaN`, `Infinity`, `nulll`; Escapes `\U0041`, `\'`, `\a`, `\u` mit drei Hex-Ziffern, `\u`
  mit `-`; Struktur `{"a"::1}`, `{"a":1:2}`, `{:1}`, `[1:2]`, `{"b"}` und `{,}` verschachtelt,
  `[,1]`, `{"a":1,"b"}`, Leerraum statt Komma in Array und Objekt; zwei BOM; NUL nach der Wurzel;
  geschütztes Leerzeichen U+00A0 als Leerraum; dazu das Fuzzing.
- **Gültige Form, die fälschlich Exit 2 ist:** keine gefunden, außer D-4 (Tiefe ab 3·10^6). Einzeln
  angenommen und byte-genau normalisiert: `-0`, `1E+2`, `1e-2`, `0.0`, `0e0`, `-0.0E-0`,
  `1e999`, eine 29-stellige Ganzzahl; `\u0000`, ein Surrogat-Paar, ein **einzelnes**
  Surrogat-Escape (RFC 8259 §8.2: von der ABNF erlaubt), groß- und kleingeschriebene Hex-Ziffern,
  ein u-Escape im Schlüssel, `\/`, `\b \f \n \r \t`, `\\`, `\"`, DEL (U+007F) roh, `ä € 𝄞` roh;
  leerer Schlüssel; doppelte Schlüssel (RFC: SHOULD, zwei Anweisungen); `[[],{}]`, zehnfach
  geschachtelte leere Arrays, `{}`, ` {} `; Leerraum aus allen vier Zeichen vor und nach jedem
  Strukturzeichen; ein BOM am Anfang (§8.1 MAY ignore; `encoding/json` lehnt ihn ab, die
  Spezifikation lässt ihn ausdrücklich entfallen).
- **`literal`-Einträge:** `"a": tru` und `"a": 01` Exit 2 beim Laden („lässt sich nicht
  zerlegen“), `"a": -0` gegen `{"a":-0}` Exit 0.

### Negativbefunde

| Bereich | Ergebnis |
|---|---|
| `jsonValidate`/`jsonParser` gegen RFC 8259 §2–§8 | geprüft, ohne Befund außer D-4. Objekt- und Array-Grammatik ohne abschließendes Komma, Zahl-Regex = §6 (`int = zero / digit1-9 *DIGIT`, `frac`, `exp`), Literale exakt, Escape-Menge = §7, UTF-8 nur in Zeichenketten geprüft, außerhalb lässt die Lexik ohnehin nur ASCII zu. Kein Index-Fehler bei Backslash am Zeichenketten-Ende (die Lexik lässt keine solche Zeichenkette durch). |
| Spezifikation 0.38.0 — Referenz-Richtung (`AGENTS.md` §3.4) | geprüft, ohne Befund. Der Diff nennt keine ADR-, Slice-, Wellen-Kennung und keinen Commit. Historie-Zeile beschreibt die Änderung, das ist ihre Aufgabe. |
| Handbuch §4 `json` und Historie 1.45 | geprüft, ohne Befund. Jede genannte Exit-2-Form gemessen; „Ebenso Exit 2, wenn die Wurzel kein Objekt ist“ trennt die eigene Regel richtig von RFC 8259; Handbuch-Version 1.45 im Kopf, in der Historie in place fortgeschrieben (unveröffentlicht). |
| CHANGELOG `[Unreleased]` | geprüft, ohne Befund. „Spezifikation 0.36.0–0.38.0“ deckt die drei berührten Stände. |
| Code-Kommentare in `json_shape.go` | geprüft, ohne Befund. Kopf-, `jsonValidate`-, `jsonCheckString`- und `next`-Kommentare decken sich mit dem gemessenen Verhalten; keine Chronik. |
| Test-Kommentare außer `TestJSONValidity` | geprüft, ohne Befund. `TestJSONUnsplittable` sagt nur noch, was die Schleife prüft. Review-Kennungen in Kommentaren sind im Bestand üblich (`shapes.go`, `shapes_test.go`, `kotlin_shape_test.go`). |
| Reihenfolge Plan → Spezifikation → Code (`AGENTS.md` §6 Schritt 4) | geprüft, ohne Befund. `f25a894` vor `268b185` vor `f778348`; der Planning-Commit berührt nur `docs/plan/planning/`. |
| Hard Rules `AGENTS.md` §3.1–§3.6 | geprüft, ohne Befund. Kein `//nolint`, keine ADR geändert, kein Move, keine Gate-Schwelle. |
| Traceability der drei Commit-Messages | geprüft, ohne Befund. Jede nennt `slice-214`, zwei zusätzlich `AC-FA-RULE-012`/`ADR-0042`. |

### Summary (Delta)

| Kategorie | Anzahl |
|---|---|
| HIGH | 2 |
| MEDIUM | 1 |
| LOW | 2 |
| INFO | 1 |

Erst-Findings: 9 von 9 erledigt (F-4 durch Wegfall des Gegenstands).

**Finding-Klassen dieses Laufs:** Zustandsangabe im Plan nach Plan-Änderung nicht nachgezogen ·
Vertrag ordnet eine eigene Regel der Quelle zu · Testbeschreibung sagt mehr zu als die Assertion ·
Implementierungsgrenze nicht benannt · Kommentar beschreibt die verworfene Alternative · Prüfung
nach vorgelagertem Prüfer unerreichbar

### Verdikt (Delta)

**Merge-blockierend:** ja, wegen D-1, D-2 (HIGH) und D-3 (MEDIUM). Am Verhalten blockiert nichts:
Der Prüfer nimmt in allen Messungen genau gültiges JSON mit Objekt-Wurzel an (4 000 000
Fuzz-Eingaben, 90 Einzelproben, 233 reale Dateien, je gegen `encoding/json`, 0 ungewollte Abweichungen),
strenger als die Referenz nur bei ungültigem UTF-8, nachsichtiger nur beim BOM am Anfang — beides spec-treu. Blockierend sind zwei
Sätze, die nachweislich nicht stimmen (Plan-Kopf, Zuordnung in Schritt 2), und ein Testkommentar,
dessen Zusage drei Mutationen grün überstehen. D-4 betrifft eine nicht benannte Grenze, nicht
einen falschen Befund.

**Übergabe:** Findings an den Implementer. D-3 ist innerhalb desselben Vorgangs (slice-214)
wie F-2 aufgetreten und zählt für `BEO-GATE/testbeschreibung-weiter-als-assertion` nicht
zusätzlich. D-4 gehört, falls eine Grenze in den Vertrag soll, als Frage an den Architect.
Dieser Report ersetzt keine Verifikation.

### Delta-Review 2 (Nachlauf 7876a5c~3..7876a5c)

**Gegenstand:** `7df9980` (planning), `5f4ec43` (Spezifikation), `7876a5c` (Test). Gleicher
Skill-Stand, gleicher frischer Kontext. **Sonden** in Klonen von `7876a5c` über `make test`:
unmutiert Exit 0; V1, V2, V5 wie oben neu angelegt.

| ID | Stand | Messung |
|---|---|---|
| D-1 | erledigt | Kopf nennt „Schritt 2 ‚Gültigkeit‘ geändert durch die Plan-Änderung in §1 (Spezifikation 0.38.0)“; der Ausschluss gilt „über die Plan-Änderung oben hinaus“. Beides deckt sich mit `268b185`. |
| D-2 | erledigt | Schritt 2 ordnet die Liste unter 1 und das fehlende `:` unter 4 der Regel zu, die Objekt-Wurzel als Zusatzbedingung des Dialekts — gegen RFC 8259 §2 richtig. Die übrigen Fehler unter 4 (Inhalt nach der Wurzel, Klammern, leeres Mitglied) bleiben unzugeordnet; der Satz behauptet über sie nichts und ist damit nicht falsch. |
| D-3 | erledigt | Kommentar sagt jetzt „jede aufgeführte … Form“ und „die aufgeführten gültigen Grenzfälle“. V1 rot: „zwei Exponent-Zeichen ("{\"a\":1e+-2}"): Fehler erwartet“; V5 rot: „doppeltes Minus ("{\"a\":--1}"): Fehler erwartet“; V2 rot: der neue Gültig-Fall mit allen Escapes aus §7 meldet „ungültige Escape-Folge“. |
| D-4 | Ausgang angenommen | LOW; der Implementer begründet und trägt den Befund als Beobachtung ins Register (Closure). Die Begründung „eine Tiefengrenze wäre eine Vertragsänderung“ trifft die Benennung einer Grenze; der Befund selbst — eine unbenannte Grenze, gezogen von der Go-Laufzeit — bleibt bis dahin bestehen und ist im Register richtig aufgehoben. |
| D-5 | erledigt | Der Konjunktiv-Nebensatz ist entfernt. |
| D-6 | Ausgang angenommen | INFO, Notiz in der Closure. |

**Neues Finding:**

| ID | Kategorie | Befund | Quelle | Pfad | Verifizierbar | Klasse |
|---|---|---|---|---|---|---|
| D-7 | LOW | Der neue Kommentar vor dem zweiten Gültig-Fall lautet im Quelltext wörtlich „jede Escape-Folge aus §7 und jede Exponent-Form aus §6“ (Bytes geprüft) — gemeint ist „§7“ bzw. „§6“; ein Leser sieht eine Escape-Folge, keinen Abschnitts-Verweis. | Wording | `internal/adapter/driven/extract/json_shape_test.go:131` | nein | Wording |

**Negativbefunde:** Spezifikation-Diff ohne ADR-/Slice-/Wellen-Kennung (`AGENTS.md` §3.4) ·
Planning-Commit berührt nur `docs/plan/planning/` · drei Commit-Messages mit `slice-214` · der
neue Gültig-Fall belegt `1E2` und `2.5e10` zusätzlich zu den Formen der ersten Zeile.

**Summary (Delta 2):** HIGH 0 · MEDIUM 0 · LOW 1 (D-7) · INFO 0.

**Verdikt (Delta 2):** nicht merge-blockierend. D-1, D-2, D-3, D-5 erledigt und gemessen; D-4 und
D-6 haben einen angenommenen Ausgang; D-7 ist Wording.
