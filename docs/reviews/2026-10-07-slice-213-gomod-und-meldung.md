# Review-Report: slice-213 — 2026-10-07

**Review-Art:** Code, unabhängiger Lauf. Der Kontext ist ein frischer Subagent ohne `fork`, und
er hat den Gegenstand nicht verfasst. Geprüft wird der Diff gegen Slice-Plan, ADR-0042, den
Vertrag (AC-FA-RULE-012, SPEC-EXTRACT-001, SPEC-RULE-001), die Hard Rules und die Go-Modulreferenz
als Primärquelle. Gegen die DoD wird **nicht** geprüft, das ist Aufgabe des Verifiers (Modul 11).

**Gegenstand:** slice-213, Commit-Range `4cdfb62..9a3b623` (`9a3b623` feat(shapes): Dialekt
gomod, einzeilige umkehrbare Meldung)

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

- Slice-Plan slice-213 (§1 mit Abgrenzung: `json` geht an slice-214; §6; §8 mit der Sichtung von
  `BEO-GATE/testbeschreibung-weiter-als-assertion` bei 2×)
- ADR-0042 (Accepted); Lastenheft 0.29.0 AC-FA-RULE-012 (Boundary `gomod`); Spezifikation 0.36.0
  SPEC-CONF-001, SPEC-EXTRACT-001 (Absätze „Dialekt `gomod`“, „`literal`-Einträge“), SPEC-RULE-001
  (Ausgabe-Regel, Zeile `shape-unused`)
- Vorlauf-Report zu slice-212 vom selben Tag (F-10, D-7, Abschnitt „Injektivität der Meldung“)
- `AGENTS.md` §3 (insb. 3.2, 3.7)
- Primärquelle, im Lauf abgerufen: Go-Modulreferenz `https://go.dev/ref/mod` §Lexical elements und
  die Block-Grammatik aller Direktiven
- `v6.13.0` · `regelwerk/modul-08-agentenrollen.md` §Konflikt-Pfad als Rollen-Sequenz

**Sonden, alle in einem Klon außerhalb des Repos** (eigenes Image-Tag, das Repo unberührt):

- `make test`, `make lint`, `make coverage-gate` auf `9a3b623`: je Exit 0, Coverage 96,5 %.
- Das gebaute Image auf **alle 153** `go.mod` unter `/Development` (davon 8 unter `testdata`), je
  mit `allow: [{pattern: '.*', match: regex}]`: **153× Exit 0.** Geltungsbereich: Lexik und
  Zerlegung; ob zwei Dateien dieselbe Normalform haben, deckt diese Messung nicht. Der Bestand
  enthält Kommentare, quotierte Modulpfade, `replace` und `exclude`, aber kein `retract`, `tool`,
  `godebug` und keine handgeschriebene Form wie `require(`.
- Vier Mutationen gegen die Tests, jede durch `git diff` als angewandt belegt; eine Kontroll-Mutation
  (`;`-Prüfung entfernt) war **rot** mit `"module a;b\n": Fehler erwartet` — die Probe erreicht
  ihren Gegenstand. Ergebnisse in F-2, F-3, F-11.
- 16 handgebaute `go.mod`-Eingaben gegen das Image (Ergebnisse in den Negativbefunden).

---

## Findings

| ID | Kategorie | Befund | Quelle | Pfad | Verifizierbar | Klasse |
|---|---|---|---|---|---|---|
| F-1 | HIGH | Die Regel-Tabelle in §3.4 des Handbuchs sagt zu `shape-unlisted` allgemein: „Geprüft wird nach Normalisierung (Kommentare, Leerraum und Umbrüche zählen nicht)“. Für `gomod` stimmt das nicht. Dort ist das Zeilenende Grammatik (SPEC-EXTRACT-001 `gomod` Schritt 1/3, ADR-0042 Punkt 3). `TestGomodLineEndIsGrammar` belegt es: `require a/b` + LF + `v1.0.0` ergibt zwei Anweisungen. Der Satz war vor diesem Slice richtig. Falsch wurde er durch den Dialekt, den dieser Slice liefert; §4 desselben Handbuchs sagt jetzt das Gegenteil. | Skill §HIGH „nachweislich falsche Tatsachenbehauptung“; SPEC-EXTRACT-001; ADR-0042 Punkt 3 | `docs/user/benutzerhandbuch.md:260` | nein — kein Gate liest Prosa gegen das Verhalten | Allgemeine Aussage durch neue Variante falsch geworden, nicht nachgezogen |
| F-2 | MEDIUM | Der Kommentar zu `TestGomodUnsplittable` sagt „Alle fail-closed-Fälle der Spezifikation“. Die Spezifikation nennt als Fehler „ein `//`, `"` oder `` ` `` direkt hinter einem Nicht-Leerraum-Zeichen“. Der Fall mit `` ` `` fehlt im Test. Mutation: `` ` `` aus der Prüfung in `gomodTokenEnd` entfernt, dann geht `` module a`b` `` als Token durch, und `make test` bleibt **grün** (Exit 0). | `BEO-GATE/testbeschreibung-weiter-als-assertion`; SPEC-EXTRACT-001 `gomod` Schritt 1 | `internal/adapter/driven/extract/gomod_shape_test.go:68`; Gegenstand `internal/adapter/driven/extract/gomod_shape.go:153` | ja — die Mutation in `make test` | Testbeschreibung sagt mehr zu als die Assertion |
| F-3 | MEDIUM | Der Kommentar zu `TestEvaluateShapesMessagesOneLine` sagt „shape-unlisted und shape-differs schreiben ihre Anweisungen einzeilig“. Geprüft wird bei `shape-differs` nur der Zweig `… (erwartet: …)`. Die Zweige `fehlt: …` und `… (nicht in der Sollform)` prüft kein Test. Mutation: `oneLine` dort jeweils entfernt, dann bleibt `make test` in **beiden** Fällen grün (Exit 0). Die Spezifikation fordert die Schreibweise „in beiden Teilen eines `shape-differs`“, also in allen drei Formen. | `BEO-GATE/testbeschreibung-weiter-als-assertion`; SPEC-RULE-001 Ausgabe-Regel | `internal/hexagon/core/shapes_test.go:270`; Gegenstand `internal/hexagon/core/shapes.go:263,265` | ja — die Mutationen in `make test` | Testbeschreibung sagt mehr zu als die Assertion |
| F-4 | LOW | Der Kommentar zu `TestShapesGomodEndToEnd` sagt, eine zusätzliche `require`-Zeile und eine zusätzliche `replace`-Direktive seien „je ein Befund“. Bei der `require`-Zeile wird die ganze Ausgabe verglichen. Bei `replace` prüft die Assertion nur `strings.Contains`, die Anzahl also nicht. Heute stimmt das Verhalten (eine Zeile); die Assertion belegt es aber nicht. | `BEO-GATE/testbeschreibung-weiter-als-assertion`; AC-FA-RULE-012 Boundary (`gomod`) „je ein Befund“ | `internal/cli/cli_shapes_test.go:309,337` | nein — eine Mutation, die einen zweiten Befund erzeugt, ist nicht naheliegend | Testbeschreibung sagt mehr zu als die Assertion |
| F-5 | LOW | Der Code ist bei einem `(` mitten in einer Zeile strenger, als die Spezifikation aufzählt. Ein einzelnes `(`-Token mitten in einer Zeile, auf oberster Ebene (`module ( x`) oder in einem Eintrag (`a ( b`), ist im Code Exit 2. Gemessen am Image: „eine Klammer an unzulässiger Stelle“. SPEC-EXTRACT-001 Schritt 3 nennt als Fehler nur `(` am Zeilenende in einem Block, `)` außerhalb eines Blocks bzw. neben dem schließenden `)` und `()` an anderer Stelle. Ein `(`-Token mitten in einer Zeile ist nach Schritt 1 Interpunktion und nach Schritt 3 Teil einer Anweisung. Der Vorlauf-Report zu slice-212 (F-10) las den Vertrag ebenso. Gedeckt ist der Code nur über die Generalklausel („Wo die Referenz schweigt … fail-closed“). Die Wirkung ist fail-closed. | SPEC-EXTRACT-001 `gomod` Schritt 1/3; slice-213 §1 („eine Lücke geht als Plan-Änderung vor den Code“) | `internal/adapter/driven/extract/gomod_shape.go:80,95` | ja — Eingabe `module ( x` gegen das Image | Code strenger als die aufgezählten Fehlerfälle des Vertrags |
| F-6 | LOW | Die Exit-2-Liste für `gomod` im Handbuch liest sich abschließend („wenn die Datei etwas enthält, das …: …“). Es fehlen Fälle, die im Code Exit 2 sind und von Hand leicht entstehen: `//` oder ein Anführungszeichen direkt hinter einem Zeichen (`require (// deps` → Exit 2, gemessen), eine `)`-Zeile mit weiteren Tokens, ein `(` mitten in einer Zeile (F-5). | SPEC-EXTRACT-001 `gomod` Schritt 1/3 | `docs/user/benutzerhandbuch.md:650-653` | ja — Eingabe `require (// c` gegen das Image | Benutzer-Doku zählt Fehlerfälle unvollständig auf |
| F-7 | LOW | Das Handbuch sagt nicht, dass die Meldung eines `gomod`-Blocks nicht als `allow`-Eintrag taugt. Die Spezifikation sagt es („bei einem `gomod`-Block auch dann nicht“). Wer die Meldung `require (b v1)` übernimmt, bekommt Exit 2: „eine Klammer steht innerhalb eines Tokens ((b)“, gemessen am Image. Das Beispiel im Handbuch zeigt die mehrzeilige Form, ohne den Grund zu nennen. | SPEC-EXTRACT-001 Absatz „`literal`-Einträge“; ADR-0042 | `docs/user/benutzerhandbuch.md:628-653` | ja — Eingabe gegen das Image | Benutzer-Doku lässt eine Vertrags-Konsequenz aus |
| F-8 | LOW | Der Handbuch-Kopf sagt „Handbuch-Version: 1.44 · … · Stand: 2026-10-06“. Die Historie-Zeile 1.44 trägt 2026-10-07. Bei 1.43 stimmten Kopf und Zeile überein. | Konsistenz Kopf ↔ Historie | `docs/user/benutzerhandbuch.md:3,1068` | nein | Kopf-Stand nicht mit der Historie nachgezogen |
| F-9 | LOW | Grammatikfehler im Testkommentar: „eine rohes Zeilenende“. | Wording | `internal/hexagon/core/shapes_test.go:247-248` | nein | Wording |
| F-10 | INFO | Abweichungen vom Go-Werkzeug, jede fail-closed oder fail-safe und vom Vertrag gedeckt bzw. dort benannt. Exit 2 bei Formen, die `go` annimmt: `require(` (gemessen: „Klammer innerhalb eines Tokens“), `require (// c`, `v1//c`. Angenommen als Bezeichner, obwohl `go` sie ablehnt: `\v`, `\f` und andere Nicht-ASCII-Leerzeichen. Als **ein** Token angenommen: `"a"b`; die Referenz sieht dort zwei Tokens. Damit unterscheidet sich die Normalform von der von `"a" b` (fail-safe). `retract [v1, v2]` geht durch (`[v1,` als Bezeichner, wörtlich nach der Referenz). Keiner dieser Fälle lässt eine Anweisung als erlaubt durchgehen. | Go-Modulreferenz §Lexical elements; SPEC-EXTRACT-001; AC-QA-02 | `internal/adapter/driven/extract/gomod_shape.go:117-221` | ja — Eingaben gegen das Image | Werkzeugverhalten weicht vom Referenz-Text ab |
| F-11 | INFO | Eine Zeile, die nur aus `()` oder `( )` besteht, ist im Code Exit 2 (leerer Kopf). Die Spezifikation nennt „Kopf leer ⇒ Fehler“ nur für den öffnenden Block, den leeren Block nur mit Kopf. Kein Test belegt den Fall. Mutation: Prüfung auf leeren Kopf im `()`-Zweig entfernt, dann bleibt `make test` grün. | SPEC-EXTRACT-001 `gomod` Schritt 3 | `internal/adapter/driven/extract/gomod_shape.go:70-77` | ja — die Mutation in `make test` | Code strenger als die aufgezählten Fehlerfälle des Vertrags |
| F-12 | INFO | `dialect: json` liegt in der geschlossenen Menge von SPEC-CONF-001, ist aber heute Exit 2: „unbekannter dialect "json" (gomod\|kotlin)“, gemessen. Der Zwischenzustand ist als Abgrenzung erklärt (slice-214) und im CHANGELOG benannt („noch nicht implementiert“). | slice-213 §1; SPEC-CONF-001 | `internal/adapter/driven/extract/shapes.go:22`; `CHANGELOG.md:11-15` | ja — Eingabe gegen das Image | Übergangszustand in einem fortlaufenden Dokument |

## Adversariale Konstruktion (Auftrag 2a)

- **(i) Zusätzliche Abhängigkeit als erlaubt durchgelassen:** **kein Fall gefunden.**
  - Verschluckt wird nur `//` am Token-Anfang, beim Go-Werkzeug genauso. An jeder anderen Stelle
    ist `//` Exit 2.
  - `/*`, `;`, verklebtes `=>`, offener Block, Schachtelung und `)` mit weiteren Tokens sind Exit 2,
    am Image gemessen bzw. durch Tests belegt.
  - Eine Datei mit `CR` als einzigem Zeilenende (`module x` CR `require evil v1`) fällt zu
    **einer** Anweisung `module x require evil v1` zusammen. Die ist unerlaubt und wird gemeldet,
    gemessen.
  - Der Vergleich läuft auf dem unveränderten Text (`markAllowed` mit `st.Text`). `oneLine` wird
    nur beim Schreiben der Meldung angewandt.
- **(ii) Zwei verschiedene Quellen, eine Normalform:** **kein Fall gefunden**, außer den gewollten.
  `require ()`, `require ( )` und `require (` LF `)` ergeben alle `require ()` (Vertrag: „leer
  als `<Kopf> ()`“). Tokens enthalten außerhalb von Zeichenketten keinen Leerraum, und Zeichenketten
  sind in einer Zeile geschlossen. Darum ist das Zusammenfügen mit einem Leerzeichen eindeutig.
  `;` ist außerhalb von Zeichenketten verboten, also ist der Block-Trenner eindeutig. Eine Zeile auf
  oberster Ebene kann nicht die Form `<Kopf> (<…>)` haben, denn `(x` ist eine Klammer im Token.
- **(iii) Reale Datei fälschlich Exit 2:** 0 von 153 (Messung oben). Möglich ist es nur bei den
  handgeschriebenen Formen aus F-10. `go mod tidy` und `go mod edit -fmt` erzeugen sie nicht.
- **(iv) Code weicht von der Spezifikation ab:** F-5 und F-11, beide strenger als die Aufzählung.
  Alle anderen Punkte der Schritte 1–3 habe ich Zeichen für Zeichen gegen den Code gelesen,
  Ergebnis in den Negativbefunden.

## Negativbefunde

| Bereich | Ergebnis |
|---|---|
| Lexik `gomod` gegen SPEC-EXTRACT-001 Schritt 1 und die Referenz | geprüft. Folgendes ist umgesetzt: BOM entfällt (`TrimPrefix`, gemessen). Leerraum ist `' '`, `\t`, `\r`; nur `LF` trennt Zeilen. `//` gilt nur am Token-Anfang. `"…"` hat Backslash-Maskierung, `` `…` `` keine. Ein Zeilenende in einer Zeichenkette, auch direkt nach Backslash, ist Exit 2. Eine Zeichenkette mitten im Token ist Exit 2. `/*` ist Exit 2, auch direkt hinter einer Zeichenkette. `;` außerhalb einer Zeichenkette ist Exit 2. Klammer im längeren Token außer `()` ist Exit 2. `=>` als Teil eines Tokens ist Exit 2, auch hinter einer Zeichenkette. Abweichungen: F-5, F-10. |
| Normalisierung (Schritt 2) | geprüft, ohne Befund. Kommentare entfallen, Tokens werden mit genau einem Leerzeichen verbunden, Zeichenketten bleiben byte-genau (`module "a\\b\rc"` erscheint maskiert, aber unverändert verglichen). |
| Zerlegung, Zeilennummer, Zeilenzahl (Schritt 3) | geprüft. Block-Kopf, Einträge, `)` allein, leerer Block am Zeilenende und `<Kopf> (<e>;<e>)` stimmen mit dem Vertrag überein. Die Original-Zeile ist die des ersten Tokens, auch beim Block (`block.line`). Die Zeilenzahl zählt `LF`, dazu eine letzte Zeile ohne `LF`, mindestens 1. Ein einzelnes `CR` zählt nicht. Abweichungen: F-5, F-11. |
| `literal`-Einträge `gomod` | geprüft, ohne Befund im Code. Sie werden über dieselbe Registry zerlegt, genau eine Anweisung, im validierenden Einstieg vor jedem Dateizugriff (`shapeLiterals`). Mehrzeilige YAML-Form funktioniert (E2E-Test). Doku: F-7. |
| Ausgabe-Regel (Auftrag 2b) | geprüft, ohne Befund im Code. `oneLine` maskiert gleichzeitig (`strings.NewReplacer`, ein Durchlauf) `\` → `\\`, `LF` → `\n`, `CR` → `\r`. Das ist ein präfixfreier Code und damit injektiv. Angewandt wird es an allen vier Stellen (`shape-unlisted`, alle drei `shape-differs`-Zweige mit beiden Teilen, `shape-unused`). Das Präfix lautet `literal: ` bzw. `regex: `. Abschließende `\r`/`\n` werden **vor** der Maskierung entfernt. Verglichen wird roh. Die zusammengesetzte `shape-differs`-Meldung bleibt eindeutig: ` (erwartet: ` kann außerhalb einer Zeichenkette in keiner gültigen Normalform der zwei Dialekte stehen (bei `gomod` wäre `(erwartet:` eine Klammer im Token). Dieselbe Herleitung wie im Vorlauf-Report. Test-Deckung: F-3. |
| Hard Rule §3.2 (keine Inline-Suppression) | geprüft, ohne Befund. Kein `nolint` im Diff; `make lint` Exit 0 im Klon. |
| Hard Rule §3.7 (Kommentare) | geprüft, ohne Befund. Die Kommentare in `gomod_shape.go` und `shapes.go` tragen Zusage, Kopplung oder Grenze im Indikativ. Geprüft habe ich auch den Satz zu `checkGomodToken` („would make two sources fall on one normal form“): Für `=>` trifft er unter dem Token-Modell des Vertrags zu, denn dort ist `=>` Interpunktion und `a=>b` fiele auf `a => b`. Er wiederholt ADR-0042 Punkt 2 und keine verworfene Alternative. |
| Schichtung | geprüft, ohne Befund. Der Dialekt liegt im Adapter `extract`, und der Kern kennt keinen Dialekt. `oneLine` liegt im Kern, weil die Meldung im Kern entsteht. |
| Abgrenzung slice-213 §1 | geprüft, ohne Befund. Kein `json`-Code, keine Vertragsdatei unter `spec/` im Diff, die Dateisuche ist unverändert. |
| `--print-config`-Gerüst | geprüft, ohne Befund. Der `gomod`-Eintrag sagt „Zeilenumbrueche zaehlen, Kommentare nicht“. Der Satz „Umbrueche zaehlen nicht“ beim `kotlin`-Eintrag steht in dessen Kommentarblock. |
| CHANGELOG `[Unreleased]` | geprüft, ohne Befund. Der Übergangsvermerk ist für `gomod` und die Meldung umgeschrieben, `json` ist als „noch nicht implementiert“ markiert (F-11 im Vorlauf-Report damit erledigt). |
| Tests für `gomod`, Rest | geprüft, ohne Befund. `TestGomodRealForm`, `TestGomodLines`, `TestGomodEmptyBlock`, `TestGomodLineEndIsGrammar`, `TestGomodNoBlockComment` und der E2E-Test für die `require`-Zeile und `/* */` belegen, was ihre Kommentare sagen. Zeilenangaben 5 und 9 im E2E-Test nachgezählt. Kontroll-Mutation rot (siehe Kopf). |

## Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 1 |
| MEDIUM | 2 |
| LOW | 6 |
| INFO | 3 |

**Finding-Klassen dieses Laufs:** Allgemeine Aussage durch neue Variante falsch geworden, nicht
nachgezogen · Testbeschreibung sagt mehr zu als die Assertion · Code strenger als die
aufgezählten Fehlerfälle des Vertrags · Benutzer-Doku zählt Fehlerfälle unvollständig auf ·
Benutzer-Doku lässt eine Vertrags-Konsequenz aus · Kopf-Stand nicht mit der Historie nachgezogen ·
Wording · Werkzeugverhalten weicht vom Referenz-Text ab · Übergangszustand in einem fortlaufenden
Dokument

**Zum Register-Eintrag (Urteil bei der Closure, nicht hier):** F-2, F-3 und F-4 gehören zur
Klasse `BEO-GATE/testbeschreibung-weiter-als-assertion`. Sie stammen aus **einem** Vorgang
(slice-213) und zählen darum einmal. Mit slice-210 und slice-211 steht der Eintrag damit bei
**3×**. Die Sichtung in slice-213 §8 („jeder Testkommentar wird beim Schreiben gegen seine
Assertion gelesen“) hat den dritten Fall nicht verhindert. Der Ausgang wird bei der Closure
zugewiesen.

## Verdikt

**Merge-blockierend:** ja, für die **Closure** von slice-213. Die Commits liegen schon auf
`main`; das Urteil betrifft den nächsten Lifecycle-Schritt. Blockierend sind F-1 (falsche
Aussage in der Benutzer-Doku, die dieser Slice falsch gemacht hat) sowie F-2 und F-3 (Test-Zusagen
ohne Beleg, Mutationen überleben). Der Code selbst hält den Vertrag. Eine durchgelassene
Abhängigkeit, eine Normalform-Kollision oder ein Fehlalarm im realen Bestand wurde nicht
gefunden. F-4 bis F-12 sind LOW oder INFO; der Implementer kann sie annehmen oder begründen. Einen
Rollen-Widerspruch zu einem HIGH gibt es bisher nicht, die Konflikt-Sequenz greift also nicht.

**Übergabe:** Die Findings gehen an den Implementer. Die Finding-Klassen gehen in die Closure §7
von slice-213 und von dort in den Zähler. Dieser Report ist ein Lauf-Beleg und ersetzt keine
Verifikation gegen die DoD.
