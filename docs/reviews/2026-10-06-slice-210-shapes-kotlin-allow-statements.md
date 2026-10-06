# Review-Report: slice-210 — 2026-10-06

**Review-Art:** Code — unabhängiger Lauf (frischer Subagent-Kontext ohne `fork`, kein Autor des
Gegenstands). Geprüft wird der Diff gegen Slice-Plan, ADR, Spezifikation und Hard Rules
(Maintainability); **nicht** gegen die DoD (Verifier, Modul 11).

**Gegenstand:** slice-210, Commit-Range `604ed20..fe2f8b2` (ein Commit `fe2f8b2`)

**Skill:** `.harness/skills/reviewer.md` @ `a6d19b6` ·
**Modell:** `claude-opus-5-5` · **Datum:** 2026-10-06

> **Zitier-Form** *(dieser Block bleibt stehen — er ist Norm, kein
> Ausfüll-Hinweis)*. Dieser Report friert ein; was er zitiert, bewegt sich
> weiter. Deshalb: **Kennung, nicht Adresse** — `slice-<Kennung>` statt seines
> Lifecycle-Pfads, `make <target>` statt eines Links auf die Sensor-Datei, eine
> Baseline-Stelle als **Tag + Pfad in Inline-Code** statt als Link. Ein
> `pfad`-Feld auf den **geprüften Gegenstand** ist davon nicht betroffen — es
> zitiert den Stand des Laufs und darf ihn festhalten.

**Eingangs-Kontext:**

- Slice-Plan slice-210 (§1 Ziel und Abgrenzung, §2, §3, §6, §8), Welle welle-16
- ADR-0041 (Accepted); im `Bezug:` ADR-0027, ADR-0034
- AC-FA-RULE-012, AC-FA-CONF-001, AC-FA-DIST-001, AC-QA-01, AC-QA-02;
  SPEC-CONF-001 (Bullet `shapes`), SPEC-EXTRACT-001 (Absatz „Sollform-Anweisungen"),
  SPEC-RULE-001 (Zeile `shape-unlisted`), SPEC-DET-001
- `AGENTS.md` §3 (Hard Rules 3.1–3.7) und §4
- Früherer Report am selben Bereich: slice-209 (`2026-10-06-slice-209-shapes-spec-first.md`) — für
  stabile Klassen-Bezeichnungen und um bekannte Fälle (F-3/N-1/K-1, F-13) nicht doppelt zu melden

**Messungen dieses Laufs** (Geltungsbereich je Messung genannt):

- `make test`, `make lint`, `make arch-check`, `make coverage-gate` auf `fe2f8b2`: alle Exit 0;
  Coverage 96,30 % (Schwelle 90 %). Deckt Kompilat, Lint-Profil, Eigen-Architektur und Testlage —
  **nicht** die Lexik-Fragen unten.
- Sonden: ein temporärer Test in einer **Kopie** des Repos (Scratch-Verzeichnis, nicht im Repo),
  ausgeführt über `make test` in der Kopie. Die Sonden protokollieren die Ausgabe von
  `normalizeKotlin`, `shapePaths` und `EvaluateShapes` für konkrete Eingaben. **Geltungsbereich:**
  belegt wird nur, was **a-check** aus der Eingabe macht. Wie **Kotlin** dieselbe Eingabe liest, ist
  aus der Lexer-Grammatik von Kotlin hergeleitet und **nicht** ausgeführt (kein Kotlin-Werkzeug im
  Lauf).
- Regel-Zählung in Handbuch §3.4 mit zwei verschieden gebauten Zählern: `grep -c` über die
  Tabellenzeilen (11) und die Namensliste der ersten Spalte (11 Namen).

---

## Findings

| ID | Kategorie | Befund | Quelle | Pfad | Verifizierbar | Klasse |
|---|---|---|---|---|---|---|
| F-1 | HIGH | Das Glossar nennt die zehn Regeln von `core-impurity` bis `wrong-direction` „die zehn geprüften Regeln (Abschnitt 3.4)“. Seit diesem Commit führt §3.4 **elf** Regeln, `shape-unlisted` eingeschlossen (zwei Zähler: 11 Tabellenzeilen, 11 Namen); die Aussage ist damit falsch, und der Diff hat das Glossar direkt darunter angefasst. | Reviewer-Skill HIGH „nachweislich falsche Tatsachenbehauptung“; `AGENTS.md` §4 (abschließende Aufzählung neben ihrer Quelle) | `docs/user/benutzerhandbuch.md:934` (gegen `:249-262`) | ja — Zählung der Tabellenzeilen in §3.4; kein Gate prüft das | Zählaussage nicht mit der gezählten Menge nachgezogen |
| F-2 | MEDIUM | In Kotlin ist `$` gefolgt von einem Backtick-Bezeichner eine **Kurz-Vorlage**, auch in einer Zeichenkette (`"$`a"b`"`). Der Lexer behandelt den Backtick in der Zeichenkette als Inhalt, deshalb schließt ein `"` im Bezeichner die Zeichenkette zu früh. Danach sind Zeichenkette und Code vertauscht, und ein `/*` aus Kotlin-Zeichenketten-Inhalt öffnet einen Kommentar, der echten Code **verschluckt**. Sonde mit vier Zeilen (`` val `a"b` = "" `` · `val s = "$`a"b`" + "`/*"` · `dependencies { implementation("com.evil:lib:1.0") }` · `val u = "*/" // "`): `normalizeKotlin` liefert drei Anweisungen, die Abhängigkeits-Zeile ist nicht dabei, und es gibt keinen Fehler. Das ist die Fehlrichtung, die ADR-0041 als einzige nicht fail-safe benennt. Der Code folgt SPEC-EXTRACT-001 Schritt 1 („ebenso `$` vor einem Bezeichner“ ist Inhalt; Backtick-Bezeichner nur als Code-Token): Die Lücke liegt im Vertrag. Begrenzt wird sie dadurch, dass die sichtbaren Rest-Anweisungen unerlaubt bleiben; still grün wird die Datei erst, wenn `allow` diese Reste deckt. | ADR-0041 §Konsequenzen (Fehl-Lexik verschluckt Code; „vollständige Lexik, belegt durch Tests je Form“); SPEC-EXTRACT-001 Schritt 1; slice-210 §1 (Vertragslücke ⇒ Plan-Änderung) | `internal/adapter/driven/extract/kotlin_shape.go:240-244` und `:264-271` | ja für die a-check-Seite (Sonde über `make test` in der Kopie); die Kotlin-Seite ist hergeleitet, nicht ausgeführt | Lexik-Aufzählung behauptet Vollständigkeit, die die Zielsprache nicht trägt |
| F-3 | LOW | Laut SPEC-CONF-001 nennt stderr bei den Exit-2-Fällen des Scans „Eintrag und Glob bzw. Datei“. Die drei Meldungen nennen Glob bzw. Datei, aber **nicht** den Eintrag (`shapes[i]`). Die Lade-Meldungen tragen ihn. | SPEC-CONF-001 (Fail-closed beim Scan) | `internal/adapter/driven/extract/shapes.go:93`, `:114`, `:118` | ja — Test auf den stderr-Text; kein Gate | Fehlermeldung trägt weniger als der Vertrag zusagt |
| F-4 | LOW | `EvaluateShapes` fasst **alle** byte-gleichen Befunde zusammen, auch zwei gleiche unerlaubte Anweisungen auf derselben Zeile im **selben** Eintrag (Sonde: zwei `evil()`-Anweisungen in Zeile 1 ⇒ ein Befund). SPEC-RULE-001 sagt „je solcher Anweisung ein Befund“; zugesagt ist das Zusammenfassen nur für dieselbe Datei in **zwei** Einträgen (SPEC-CONF-001). | SPEC-RULE-001 (`shape-unlisted`), SPEC-CONF-001 | `internal/hexagon/core/shapes.go:192-193`, `:221-229` | ja — Kern-Test mit zwei gleichen Anweisungen einer Zeile | Zusammenfassung breiter als ihr Vertrag |
| F-5 | LOW | §6 Risiko 2 verlangt, als Testfall aufzunehmen, dass ein Zeilenende nach Postfix `!!`/`++`/`--` oder nach `>` zwei Anweisungen verbindet. Getestet ist nur `!!`. Die Sonde zeigt dasselbe Verbinden für `x++⏎y()` ⇒ `x++y()` und `val a = b >⏎c()` ⇒ `val a=b>c()`. | Slice-Plan slice-210 §6 | `internal/adapter/driven/extract/kotlin_shape_test.go:81-83` | ja — `make test` mit den fehlenden Fällen | Plan-Zusage nur teilweise eingelöst |
| F-6 | LOW | Der Kommentar zu `TestShapesGreenExceptListedNote` sagt, „ein verbotenes Muster nur im Kommentar oder in der Zeichenkette einer Anweisung ist kein eigener Befund“. In der Fixture steht in **keinem** Kommentar ein Abhängigkeits-Muster, und das Muster in der Zeichenkette liegt in einer **unerlaubten** Anweisung. Der Boundary-Fall von AC-FA-RULE-012 (Muster in der Zeichenkette einer **erlaubten** Anweisung ⇒ kein Befund) läuft damit nicht durch die CLI. | AC-FA-RULE-012 (Boundary „Kommentar und Zeichenkette“); `AGENTS.md` §3.7 | `internal/cli/cli_shapes_test.go:28-41`, `:59-72` | ja — Fixture-Inhalt lesen | Testbeschreibung behauptet mehr als die Fixture prüft |
| F-7 | LOW | `InsideRoot` nimmt `./a` als gültig an (im Test als gültig festgehalten). Ein `files`-Glob dieser Form trifft aber nie eine Datei, weil die verglichenen Pfade kein `./` tragen. Sonde: `./kotlin_shape.go` ⇒ Exit-2-Meldung „trifft keine Datei“ bei vorhandener Datei. Fail-closed, aber die Meldung nennt die falsche Ursache. | SPEC-CONF-001 (`files`: „dieselbe Glob-Semantik wie `layers`“) | `internal/hexagon/core/shapes.go:142-151`; `internal/hexagon/core/shapes_test.go:101`; `internal/adapter/driven/extract/shapes.go:159` | ja — Sonde/Test über `shapePaths` | Validierung akzeptiert eine Form, die der Matcher nie trifft |
| F-8 | LOW | Im `--print-config`-Gerüst wandert der Kommentar der Zeile `# resolution:` eine Spalte nach links (Spalte 35 → 34). Die Nachbarzeilen bleiben auf 35. Die Änderung gehört nicht zu `shapes` und steht nicht in §3 des Plans. | Slice-Plan slice-210 §3; `AGENTS.md` §6 Schritt 4 (Abgrenzung nicht still weiten) | `internal/cli/cli.go:340` | ja — `git show fe2f8b2 -- internal/cli/cli.go` | Nebenänderung außerhalb des Plans |
| F-9 | LOW | Im Zeichen-Literal überspringt ein Backslash auch ein Zeilenende; das Literal läuft dann über die Zeile weiter, und `l.line` bleibt stehen. Sonde: `val c='\⏎x()'` ⇒ eine Anweisung, kein Fehler. Bei `"…"` ist dieselbe Lage „nicht geschlossen“ (Exit 2). Text verschwindet nicht (das Literal bleibt wörtlich erhalten), aber die beiden Literal-Formen sind uneinheitlich fail-closed. | SPEC-EXTRACT-001 Schritt 1/4 | `internal/adapter/driven/extract/kotlin_shape.go:303-304` | ja — Sonde/Test über `normalizeKotlin` | Lexik-Fall uneinheitlich zwischen Literal-Formen |
| F-10 | LOW | Laut Handbuch §4 „spielen Leerraum und Zeilenumbrüche keine Rolle“. Ein Zeilenende ist aber eine Anweisungsgrenze, und nach `!!`/`++`/`>` verbindet es zwei Anweisungen (Fehlalarm, ADR-0041 §Konsequenzen „Negativ“). §4 nennt diese Fehlalarm-Quelle nicht. | ADR-0041 §Konsequenzen; Benutzer-Doku | `docs/user/benutzerhandbuch.md` §4 „Sollform je Datei“, Punkt „Normalisiert.“ | nein — Doku-Urteil | Handbuch-Zusage breiter als das Verhalten |
| F-11 | INFO | Die Meldung ist die normalisierte Anweisung, byte-genau, Roh-Strings eingeschlossen. Eine Anweisung mit mehrzeiligem `"""…"""` erzeugt darum eine **mehrzeilige** Befundzeile auf stdout (Sonde: Meldung enthält `\r\n`). Bei CRLF-Checkout ist ein Roh-String mit `\r\n` nie gleich einem `literal`-Eintrag aus YAML, der `\n` trägt: ein fail-safe Fehlalarm. Das ist spec-konform; dass die Ausgabe eine Zeile je Befund ist, sagt der Vertrag nicht zu. | AC-FA-RULE-012 (Determinismus-Kriterium, Form `pfad:zeile: <klasse>: <meldung>`); SPEC-RULE-001 | `internal/hexagon/core/shapes.go:188`; `internal/adapter/driven/report/report.go:28` | ja — Sonde | Ausgabeform einzeilig nicht zugesichert |
| F-12 | INFO | Der `dialect` wird in der Extraktion (`Validate`) geprüft, nicht beim Laden; so legt es ADR-0041 Punkt 9 fest. Folge: Hat Eintrag 0 einen unbekannten `dialect` und Eintrag 1 einen ungültigen `mode`, meldet der Lauf zuerst Eintrag 1. Deterministisch, beobachtbar Exit 2 in beiden Pfaden; die Reihenfolge folgt aber nicht der Deklaration. | SPEC-CONF-001 (Deklarationsreihenfolge), ADR-0041 Punkt 9 | `internal/adapter/driven/extract/shapes.go:36-42`; `internal/adapter/driven/config/config.go` (`decodeShapes`) | ja — Test mit zwei fehlerhaften Einträgen | — (Hinweis, keine Klasse) |
| F-13 | INFO | Kein Test belegt, dass `--print-config` den `shapes`-Block zeigt; für `resolution` gibt es dafür einen eigenen Test. Ob der DoD-Punkt belegt ist, ist eine Frage an den Verifier. | Slice-Plan slice-210 §2 (Verweis auf die zuständige Rolle) | `internal/cli/cli_test.go:509` (Vorbild); `internal/cli/cli.go:331-337` | ja — `make test` | — (Hinweis, keine Klasse) |

## Negativbefunde

| Bereich | Ergebnis |
|---|---|
| Hard Rules `AGENTS.md` §3.1–§3.6 | geprüft, ohne Befund — kein `//nolint` unter `internal/`; Spec-Straten und ADRs nicht berührt (Dateiliste des Commits); kein Move; keine Schwelle gesenkt (Coverage-Gate unverändert 90 %, Ist 96,30 %) |
| Kommentar-Regel `AGENTS.md` §3.7 (alle neuen Kommentare in Code und Tests) | geprüft, ohne Befund — Zusage, Abgrenzung oder Grenze; keine verworfene Alternative im Konjunktiv, kein abwesender Text. Der Kommentar der Test-Fixture ist als F-6 erfasst (Abdeckung, nicht Kommentar-Klasse) |
| Lexer: verschachtelte Block-Kommentare, `//` in Block-Kommentar, Shebang, BOM, `LF`/`CRLF`/`CR`, Roh-String-Ende (Lauf ≥ 3, überzählige vorne), Dollar-Präfix, vorwärts gezählte `$`-Folge, maskiertes `\$`, Vorlagen mit eingebetteter Zeichenkette, Kommentar mit `}` und Zeichen-Literal `'}'` (Sonden P6/P7) | geprüft, ohne Befund außer F-2 und F-9 — die Fälle F-3/N-1/K-1 aus slice-209 halten, auch Zeichen für Zeichen durchgespielt |
| Normalisierung: Kollision zweier verschiedener Quellen | geprüft, ohne Befund außer über F-2 — die Klassen Wort/Operator/Rest trennen `a - -b`/`a--b`, `1 L`/`1L`; was zusammenfällt (`a .b`/`a.b`, `x "a"`/`x"a"`), ist in Kotlin gleichbedeutend oder ungültig |
| Fortsetzungsregel gegen SPEC-EXTRACT-001 (a)/(b)/(c), `;` in `(`/`[`, zusammenfallende Grenzen, Grenze nach `{`/vor `}`, Original-Zeile, Zeilenzahl | geprüft, ohne Befund — Abweichungen nur in der fail-safe Richtung (`?` mit Leerraum vor `.` verbindet); die bekannten Fehlalarme sind F-5 und F-10 |
| Exit-2-Fälle beim Laden (SPEC-CONF-001) | geprüft, ohne Befund — jeder aufgezählte Fall hat einen Pfad: Schlüssel strikt, `allow`-Objekt von Hand strikt, `files` leer/leerer Glob/außerhalb der Wurzel, `expect` außerhalb, `mode`/`match`/leeres `pattern`, Regex für sich **und** umhüllt; `exact`/`unused` bewusst Exit 2 (§1 ⇒ slice-211) |
| Exit-2-Fälle beim Scan, vor jeder stdout-Ausgabe | geprüft, ohne Befund außer F-3 — Glob ohne Treffer, `exclude`-Widerspruch (ohne Prune gesucht), nicht zerlegbare Datei; Einträge in Deklarationsreihenfolge, Dateien sortiert |
| Sortierung und Determinismus (SPEC-DET-001) | geprüft, ohne Befund außer F-4 — Import- und Shapes-Befunde werden zusammengeführt und **einmal** total sortiert |
| Unabhängigkeit von `layers`/`languages`/`composition_root` | geprüft, ohne Befund — `EvaluateShapes` und `Shapes` lesen keines der drei |
| `--print-graph` | geprüft, ohne Befund — `Validate` prüft `dialect` und `literal`-Einträge ohne Walk; Graph unverändert (§1, Bestand bleibt) |
| Hermetik der Dateisuche | geprüft, ohne Befund — absolute und `..`-führende Globs abgewiesen; Symlinks werden weder als Start-Verzeichnis verfolgt noch als Datei gelesen (`IsRegular`) |
| Architektur: Kern rein, keine laterale Adapter-Kante, Port-Erweiterung | geprüft, ohne Befund — `core/shapes.go` importiert nur `fmt`/`path`/`regexp`/`strings`; der Konfigurations-Adapter importiert die Extraktion nicht; `ExtractionPort` hat eine Methode mehr, die CLI ruft sie über den Port-Typ; `make arch-check` Exit 0 |
| Beobachtung `BEO-KERN/dirvocab-portfor-auseinander` (§8 des Plans: beim Review prüfen) | geprüft, ohne Zuwachs — das `mode`-Vokabular steht nur im Kern, das `dialect`-Vokabular nur in der Extraktion; keines der beiden ist doppelt geführt. Abweichend von §8 („liegen nur im Adapter“) liegt `mode` im Kern — das ist dasselbe Muster wie bei den `constructs`-Werten |
| Regex-Einträge | geprüft, ohne Befund — für sich und verankert kompiliert; die sichere Klasse ist gegen Code, Vorlage und `.*` getestet |
| CHANGELOG `[Unreleased]` gegen den Code | geprüft, ohne Befund — der Eintrag aus slice-209 ist umgeschrieben und nennt nur, was implementiert ist (`allow-statements`, literal/RE2, drei Exit-2-Fälle) |
| Handbuch §3.4-Tabelle, §4-Optionalblöcke, §4-Abschnitt, Historie 1.42 | geprüft, ohne Befund außer F-1 und F-10 — das Beispiel zeigt nur die sichere Klasse `"[^"$\\]*"` |
| Tests auf Tautologie | geprüft, ohne Befund — kein Test prüft nur, was er selbst baut; Lücken sind F-5, F-6 und F-13 |

## Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 1 |
| MEDIUM | 1 |
| LOW | 8 |
| INFO | 3 |

**Finding-Klassen dieses Laufs:** Zählaussage nicht mit der gezählten Menge nachgezogen ·
Lexik-Aufzählung behauptet Vollständigkeit, die die Zielsprache nicht trägt (schon bei slice-209 —
**anderer** Vorgang, zählt neu) · Fehlermeldung trägt weniger als der Vertrag zusagt ·
Zusammenfassung breiter als ihr Vertrag · Plan-Zusage nur teilweise eingelöst · Testbeschreibung
behauptet mehr als die Fixture prüft · Validierung akzeptiert eine Form, die der Matcher nie trifft ·
Nebenänderung außerhalb des Plans · Lexik-Fall uneinheitlich zwischen Literal-Formen ·
Handbuch-Zusage breiter als das Verhalten · Ausgabeform einzeilig nicht zugesichert

## Verdikt

**Merge-blockierend:** ja — F-1 (HIGH) und F-2 (MEDIUM).

F-1 ist gegen das Repo-Artefakt adversarisch geprüft, mit zwei Zählern über die Regel-Tabelle.
F-2 ist auf der a-check-Seite durch eine Sonde belegt; die Kotlin-Seite ist hergeleitet. Weil
der Code dem Vertrag folgt, liegt die Lücke in SPEC-EXTRACT-001. Nach §1 des Slice-Plans
(„wer hier eine Vertragslücke findet, hält an und geht über eine Plan-Änderung“) ist das keine
Sache eines still erweiterten Code-Pfads. Die LOW-Findings blockieren für sich nicht.

**Übergabe:** Findings gehen an den Implementer; F-2 betrifft den Vertrag und geht über den in
§1 benannten Weg (Plan-Änderung, Planner/Architect). Die Finding-Klassen gehen zusätzlich in die
Closure-Notiz §7 von slice-210 und von dort in den Zähler. Dieser Report ist Lauf-Beleg und
ersetzt keine Verifikation (DoD, bewusstes Brechen der Gegenprobe-Tests: Verifier, Modul 11).
