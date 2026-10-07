# Review-Report: slice-212 — 2026-10-07

**Review-Art:** Plan / Design, unabhängiger Lauf. Der Kontext ist ein frischer Subagent ohne
`fork`, und er hat den Gegenstand nicht verfasst. Geprüft wird der Vertrags-Diff gegen
Slice-Plan, Welle-Plan, ADRs, Hard Rules und die Primärquellen der Formate. Gegen die DoD wird
**nicht** geprüft, das ist Aufgabe des Verifiers (Modul 11).

**Gegenstand:** slice-212, Commit-Range `95656bc..f639a2f` (`80f6b58` Messung, `f639a2f` Vertrag)

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

- Slice-Plan slice-212 (§1 mit drei Abnahme-Fragen, §1b Messung, §6), Welle-Plan welle-17
- ADR-0042 (Proposed, Gegenstand), ADR-0041 (Accepted, Teil-Ablösung Punkt 10)
- AC-FA-RULE-012 (Lastenheft 0.29.0), AC-QA-01, AC-QA-02; SPEC-CONF-001, SPEC-EXTRACT-001,
  SPEC-RULE-001 (Spezifikation 0.36.0)
- `AGENTS.md` §3 (insb. 3.4, 3.5, 3.7) und §6 Schritt 4
- Register `BEO-SPEC/lexik-vertrag-ohne-sprach-gegenprobe` (2×) und
  `BEO-SPEC/ausgabe-einzeilig-nicht-zugesichert` (2×)
- Primärquellen, im Lauf abgerufen: Go-Modulreferenz `https://go.dev/ref/mod` §go.mod files
  (Lexical elements, Grammar, alle Direktiven einschließlich `toolchain`/`godebug`/`tool`/`ignore`)
  und RFC 8259. Zusätzlich das Werkzeug selbst, `golang.org/x/mod/modfile/read.go` (Lexer und
  Parser, die das `go`-Kommando nutzt).
- Reale Manifeste, nur gelesen: alle `go.mod` unter `/Development` und die `package.json` von
  m-trace sowie eine Stichprobe aus `node_modules`
- `v6.13.0` · `regelwerk/modul-04-adrs.md` §Ziel-Form: ADR (MADR); `v6.13.0` ·
  `regelwerk/modul-08-agentenrollen.md` §Konflikt-Pfad als Rollen-Sequenz

---

## Findings

| ID | Kategorie | Befund | Quelle | Pfad | Verifizierbar | Klasse |
|---|---|---|---|---|---|---|
| F-1 | HIGH | ADR-0042 §Kontext sagt: „Gemessen an **allen** lokalen Repos mit `.a-check.yml`: 6 `go.mod`“. Abhängigkeiten stünden dort „in `require ( … )`-Blöcken“. Das stimmt nicht. a-check selbst trägt `.a-check.yml` und zwei `go.mod`, zusammen sind es 8. `a-check/go.mod` hat eine **einzeilige** Direktive `require gopkg.in/yaml.v3 v3.0.1` ohne Block. Auch slice-212 §1b („je Datei 2× `require`“) stimmt für 2 der 6 genannten Dateien nicht: `d-check/tools/archive-wave/go.mod` und `claude-ai-harness-init/go.mod` haben gar kein `require`. Nachgezählt mit zwei verschiedenen Zählern: `grep` je Direktive und ein Zerleger nach dem SPEC-Modell. Auf die Entscheidung wirkt sich das nicht aus: Eine zusätzliche einzeilige `require` ist eine neue Anweisung und damit rot. Der falsche Satz wird aber mit `Accepted` unveränderlich. | Skill §HIGH „nachweislich falsche Tatsachenbehauptung“; Skill §Mess-Regeln (Geltungsbereich; zweimal zählen) | `docs/plan/adr/0042-shapes-benannte-dialekte-gomod-json.md:33-36`; `docs/plan/planning/in-progress/slice-212-generischer-dialekt-spec-first.md:65,70` | nein — kein Gate vergleicht Mess-Prosa mit dem Bestand | Messbeleg deckt die behauptete Menge nicht |
| F-2 | MEDIUM | ADR-0042 `Bezug:` sagt: „löst Punkt 10 ab … alle übrigen Punkte von ADR-0041 gelten unverändert fort“. Für die neuen Dialekte setzt der Vertrag aber auch andere Punkte außer Kraft. **Punkt 4** von ADR-0041 sagt „Die Anweisungsgrenze ist eine Fortsetzungsregel, keine Grammatik“. SPEC-EXTRACT-001 sagt für `gomod`/`json` das Gegenteil: „vollständige Zerlegung, nicht als Fortsetzungsregel“. **Punkt 3** sagt, Leerraum fällt weg außer zwischen Wort- oder Operatorzeichen. Bei `gomod` steht stattdessen immer genau ein Leerzeichen zwischen Tokens, bei `json` entfällt Leerraum vollständig. Die Konsequenz von ADR-0041 „Zeilenumbrüche ändern das Urteil nicht“ gilt für `gomod` nicht. ADR-0041 selbst bleibt unverändert, ein Verstoß gegen `AGENTS.md` §3.5 liegt also nicht vor. Der Umfang der Ablösung ist aber zu eng angegeben. | ADR-0041 Punkte 3, 4 und §Konsequenzen; Skill §MEDIUM „`Bezug:` unvollständig oder unpassend“ | `docs/plan/adr/0042-shapes-benannte-dialekte-gomod-json.md:10-12`; `spec/spezifikation.md:371-373` | nein — inferentiell | Teil-Ablösung benennt nicht alle berührten Punkte des Vorgängers |
| F-3 | MEDIUM | Boundary (`gomod`) in AC-FA-RULE-012 sagt zu: Unterscheidet sich ein `go.mod` „nur in Leerraum, Zeilen-Kommentaren und Einrückung“, gibt es **keinen** Befund. Laut SPEC-EXTRACT-001 ist `LF` aber Leerraum **und** signifikant. Gegenbeispiel: `require example.com/a v1.0.0` wird umbrochen zu `require example.com/a` + LF + `v1.0.0`. Das ergibt zwei Anweisungen und damit Befunde. ADR-0042 §Konsequenzen nennt nicht, dass Zeilenumbrüche bei `gomod` das Urteil ändern. | AC-FA-RULE-012 Boundary (`gomod`); SPEC-EXTRACT-001 `gomod` Schritt 1/3; Skill §MEDIUM „fehlende wesentliche Konsequenz“ | `spec/lastenheft.md:420`; `spec/spezifikation.md:329-330`; `docs/plan/adr/0042-shapes-benannte-dialekte-gomod-json.md:84-97` | ja — ein Testfall in `make test`, sobald slice-213 läuft | Akzeptanzkriterium sagt mehr zu, als die Zerlegung trägt |
| F-4 | MEDIUM | Die neue einzeilige Schreibweise lässt sich nicht umkehren. Ein Zeilenende wird zu `\n`, ein Backslash in der Meldung wird aber nicht maskiert. Gegenbeispiel `kotlin`: `f("""a\nb""")` (Backslash und `n` als Inhalt) und `f("""a` + LF + `b""")` ergeben dieselbe Meldung. Zwei Folgen: Ein `shape-differs` kann `X (erwartet: X)` mit gleichem Text auf beiden Seiten zeigen. Und zwei verschiedene Anweisungen auf einer Zeile fallen nach der Regel „byte-gleiche `shape-*`-Befundzeilen werden einmal ausgegeben“ zu **einem** Befund zusammen. Beides steht weder in ADR-0042 Punkt 5 noch in den Konsequenzen. Die Frage aus `BEO-SPEC/ausgabe-einzeilig-nicht-zugesichert` ist damit für Einzeiligkeit entschieden, für Eindeutigkeit nicht. | ADR-0042 Punkt 5; SPEC-RULE-001 (Einzeiligkeit, `shape-differs`); SPEC-CONF-001 (Dedup byte-gleicher Zeilen); AC-QA-01 | `spec/spezifikation.md:469`; `spec/spezifikation.md:94` (Dedup-Satz); `docs/plan/adr/0042-shapes-benannte-dialekte-gomod-json.md:71-73` | ja — ein Testfall in `make test` (slice-213) | Ausgabe-Escape nicht umkehrbar, Grenze nicht benannt |
| F-5 | MEDIUM | Für die `shape-unused`-Meldung regeln zwei Stellen derselben Spezifikation abschließende Zeilenenden verschieden. Die Tabellenzeile (seit 0.35.0) sagt: „abschließende Zeilenenden entfallen, jedes **innere** Zeilenende … wird zu `\n`“. Der neue Satz sagt: „**jedes** Zeilenende darin (`LF`, `CRLF`, `CR`) wird als … `\n` geschrieben — … im Eintrag eines `shape-unused`“. Bei einem YAML-Block-Skalar (`\|`) mit abschließendem LF folgt aus der Zeile `foo`, aus dem Satz `foo\n`. | SPEC-RULE-001; `BEO-SPEC/ausgabe-einzeilig-nicht-zugesichert` | `spec/spezifikation.md:398`; `spec/spezifikation.md:469` | ja — der bestehende `shape-unused`-Test in `make test` hält eine der beiden Lesarten fest | Zwei Stellen desselben Vertrags regeln denselben Fall verschieden |
| F-6 | MEDIUM | Bei beiden neuen Dialekten ist die normalisierte Form, also die Meldung, als `literal`-Eintrag nicht wieder zerlegbar. Die Block-Normalform von `gomod`, `require (a v1;b v1)`, enthält `(a`. Das ist „ein `(` innerhalb eines längeren Tokens“, also Exit 2. Die Mitglieds-Normalform von `json`, `"name":"x"`, ist kein Objekt, also Exit 2 („Wurzel, die kein Objekt ist“). Wer eine Meldung in `allow` übernimmt, bekommt deshalb Exit 2. Ein `json`-Literal muss als `{…}` geschrieben werden, ein `gomod`-Block mehrzeilig. So wie ich es lese, ist das bei `kotlin` anders: Dort lässt sich die Normalform erneut zerlegen. Weder ADR-0042 §Konsequenzen noch SPEC-CONF-001 noch das Lastenheft-Kriterium „eine Liste, die die Mitglieder … erlaubt“ sagen, wie ein solcher Eintrag zu schreiben ist. | ADR-0041 Punkt 5 (Literal wie Datei normalisiert); SPEC-CONF-001; AC-FA-RULE-012 Boundary (`json`); Skill §MEDIUM „fehlende wesentliche Konsequenz“ | `spec/spezifikation.md:345-346,363-366,373`; `spec/lastenheft.md:421`; `docs/plan/adr/0042-shapes-benannte-dialekte-gomod-json.md:89-91` | ja — ein Testfall in `make test` (slice-213) | Normalform ist kein Fixpunkt der Normalisierung, Konsequenz nicht benannt |
| F-7 | MEDIUM | Der Welle-Plan ist von der Entscheidung überholt und wurde nicht nachgezogen. welle-17 §1 sagt „ohne für jede Sprache einen eigenen Lexer im Code zu führen“. ADR-0042 wählt Option E, „je Format eigener Code“. welle-17 §6 schließt „Änderungen am Kotlin-Dialekt“ aus, ohne Bedingung („der generische Dialekt steht neben ihm“). ADR-0042 Punkt 5 ändert die `kotlin`-Ausgabe. slice-212 §1 sieht dafür eine „benannte Plan-Änderung“ vor. Im Diff steht sie nur als „Vorschlag“ in §1b. §1 und der Welle-Plan sind unverändert, ebenso die Titel („generischer Dialekt“). | `AGENTS.md` §6 Schritt 4 (Plan-Änderung vor dem Code, nicht still); slice-212 §1 | `docs/plan/planning/welle-17-shapes-generischer-dialekt.md:22,88-89`; `docs/plan/planning/in-progress/slice-212-generischer-dialekt-spec-first.md:48-50,101-105` | nein — inferentiell | Plan-Abgrenzung durch die Entscheidung überholt, Plan nicht nachgezogen |
| F-8 | LOW | Wo die Quelle eine Form **verbietet**, lässt der Vertrag sie zu, und die Normalisierung ist dann nicht injektiv. **`gomod`:** Die Referenz sagt „`/* */` comments are not allowed“, das Werkzeug bricht bei `/*` ab. Der Vertrag behandelt `/` und `*` als gewöhnliche Zeichen, und ADR-0042 §Kontext gibt die Referenz als „kein Kommentar“ wieder. Weitere Kollisionen bei `gomod`: Ein `;` innerhalb eines Bezeichners fällt mit dem Block-Trenner zusammen (`a v1;b v2` als ein Eintrag ≙ zwei Einträge). `a=>b` normalisiert gleich `a => b`, das Werkzeug liest `a=>b` aber als **einen** Bezeichner. **`json`:** Zeichen werden nur gegen eine Menge geprüft, nicht als Token. `nu ll` wird zu `null`, `1 2` zu `12`, `"a" "b"` wird ein Mitglied. In der Wirkung ist alles fail-safe, weil das Werkzeug solche Dateien ablehnt. Als Grenze ist es aber nicht benannt. | AC-QA-02 (ausgewiesene Grenze); ADR-0042 Punkt 2 („abgeleitet“, fail-closed); Go-Modulreferenz §Lexical elements; RFC 8259 §2, §6 | `spec/spezifikation.md:330-339,350-362`; `docs/plan/adr/0042-shapes-benannte-dialekte-gomod-json.md:45-47` | ja — Testfälle in `make test` (slice-213) | Normalisierung nur auf gültigen Eingaben injektiv, Grenze nicht benannt |
| F-9 | LOW | Mehrere Lexik-Fälle lässt der Vertrag offen, slice-213 müsste sie selbst entscheiden. (i) Beginnt ein `"` oder `` ` `` **innerhalb** eines Bezeichners eine Zeichenkette, oder ist es ein Bezeichner-Zeichen? Das Werkzeug liest es als Bezeichner-Zeichen. (ii) Ist ein Backslash vor `LF` in `"…"` eine Maskierung oder „ein Zeilenende in einer Zeichenkette“? Das Werkzeug akzeptiert ihn, und anders als bei `kotlin` regelt die Spezifikation den Fall nicht. (iii) Zählt bei `gomod` ein einzelnes `CR` als Zeile? Die *Zeilenzahl* (Meldung `fehlt:` bei `shape-differs`) ist nur im `kotlin`-Abschnitt definiert. (iv) Wie wird `=>` aus `==>` abgetrennt? Nach meiner Ableitung ist keine Lesart nicht fail-safe. | SPEC-EXTRACT-001 `gomod`/`json`; SPEC-RULE-001 `shape-differs`; slice-212 §1 (Implementierbarkeit) | `spec/spezifikation.md:325-348,350-369` | nein | Grundbegriff des Normalisierungsvertrags undefiniert |
| F-10 | INFO | Abgleich mit dem Werkzeug statt nur mit dem Referenz-Text. `golang.org/x/mod/modfile` weicht von „Punctuation tokens include (, ), and =>“ ab. Ein `//` beendet dort auch **mitten** im Bezeichner das Token und beginnt einen Kommentar; der Vertrag macht daraus Exit 2. `[ ] { } ,` sind Interpunktion, `=>` dagegen ist **kein** eigenes Token. `require ()` auf einer Zeile ist ein leerer Block; der Vertrag macht daraus Exit 2. Ein `(` mitten in der Zeile ist ein gewöhnliches Token, so wie im Vertrag. Jede dieser Abweichungen ist fail-closed oder fail-safe. Messung, Geltungsbereich: 153 `go.mod` unter `/Development` (auch Fremd-Repos unter `sandbox/` und `testdata`) und 3005 `package.json` (5 aus m-trace, 3000 aus `node_modules`), geprüft mit **meinem** Modell von Lexik und Zerlegung, nicht mit a-check. Ergebnis: 0× Exit 2. Ob die Normalform zweier Dateien gleich ist, deckt das Modell nicht. | Go-Modulreferenz; `modfile/read.go` (`readToken`, `parseStmt`, `parseLineBlock`); RFC 8259 | `spec/spezifikation.md:325-373` | ja — die Fälle als Tests in `make test` (slice-213) | Werkzeugverhalten weicht vom Referenz-Text ab |
| F-11 | INFO | Beide CHANGELOG-Einträge unter `[Unreleased]` beschreiben einen vorgeschlagenen und noch nicht implementierten Vertrag („noch nicht implementiert“). Dieser Übergangszustand wird mit slice-213 falsch. Ob slice-213 ihn fortschreibt, ist dort nicht ausdrücklich gesagt. | `AGENTS.md` §6 Schritt 7 | `CHANGELOG.md:9-22` | nein | Übergangszustand in einem fortlaufenden Dokument |

## Adversariale Konstruktion (Auftrag 2)

Je ein Versuch in jeder der drei Richtungen, gegen Spezifikation 0.36.0 und das Werkzeug:

- **(a) Zusätzliche Abhängigkeit als erlaubt durchgelassen:** **kein Fall gefunden.**
  - `gomod`: Verschluckt wird nur `//` am Token-Anfang, beim Werkzeug genauso. Ein `//` an
    anderer Stelle ist Exit 2, das Werkzeug ist dort nachsichtiger.
  - Zeichenketten dürfen kein `LF` enthalten. Eine `)`-Zeile außerhalb eines Blocks und ein am
    Dateiende offener Block sind Exit 2.
  - `require ( example.com/evil v1` + LF + `)` ist Exit 2 („`)`-Zeile außerhalb“).
  - `CR`-only-Dateien fallen zu **einer** Anweisung zusammen und sind dann rot.
  - Die Normalform-Kollisionen aus F-8 treffen nur Dateien, die das Werkzeug ablehnt.
  - `json`: Jedes neue oder geänderte Wurzel-Mitglied ist eine neue Anweisung. Doppelte Schlüssel
    und `\u`-Schreibweisen von Schlüsseln (`"dependencies"`) sind ebenfalls neue Anweisungen.
    Leerraum zu entfernen ist auf gültigem JSON injektiv.
- **(b) Gültige reale Datei mit Exit 2:** **kein Fall** unter den genannten Manifesten
  (`d-check`, `m-trace/apps/api`, `pg-change-feed`, `pgwire-recorder`, `m-trace/**/package.json`)
  und im erweiterten Bestand (F-10). Theoretisch möglich, im Bestand nicht vorhanden:
  - `x//c` ohne Leerraum davor
  - `require(` ohne Leerzeichen
  - `require ()` auf einer Zeile
  - `(// c`
- **(c) Zwei Quellen, eine Normalform:** gefunden, siehe F-8 und, für die Ausgabe, F-4.

## Negativbefunde

| Bereich | Ergebnis |
|---|---|
| Referenz-Richtung der Spec-Straten (`AGENTS.md` §3.4) | geprüft, ohne Befund. Gelesen wurden alle hinzugefügten Zeilen unter `spec/` im Diff, gesucht nach `ADR`, `slice-`, `welle` und Hash-Mustern. Kein Treffer, auch nicht in den Historie-Zeilen 0.29.0 und 0.36.0. Die externen Quellen (Go-Modulreferenz, RFC 8259) sind keine Repo-Straten. |
| ADR-Immutabilität (`AGENTS.md` §3.5) | geprüft, ohne Befund. `git diff --stat 80f6b58^ f639a2f` berührt ADR-0041 nicht. ADR-0041 sieht diese Ablösung selbst vor (Re-Evaluierungs-Trigger 1: „Folge-ADR, die Punkt 10 ablöst“). Den Umfang der Ablösung betrifft F-2. |
| ADR-Form ADR-0042 (`v6.13.0` · `regelwerk/modul-04-adrs.md` §Ziel-Form: ADR) | geprüft, ohne Befund. Kopffelder Status/Datum/Autor/Bezug/Schärft/Regeln vorhanden, fünf Alternativen mit Trade-off, Fitness Function mit `make test`, Re-Evaluierungs-Trigger, Geschichte. Index-Zeile in `docs/plan/adr/README.md` vorhanden. |
| AC-Form AC-FA-RULE-012 | geprüft, ohne Befund in der **Form**: Happy, Boundary (jetzt sechs), Negative, Determinismus, Out-of-Scope; Versions-Bump 0.29.0 mit Historie-Zeile. Given/When/Then gelesen. Was die gomod-Boundary inhaltlich zusagt, betrifft F-3. |
| Lexik `gomod` gegen die Go-Modulreferenz §Lexical elements | geprüft. `CR` ist Leerraum, nur `LF` signifikant ✓. `//` bis Zeilenende ✓, die Einschränkung auf den Token-Anfang ist eine eigene fail-closed-Wahl, im Werkzeug ist es anders (F-10). Zeichenketten `"…"` mit Backslash-Escape und `` `…` `` ohne ✓. Interpunktion `(`, `)`, `=>` entspricht dem Wortlaut der Referenz (das Werkzeug weicht ab, F-10). Block-Grammatik `"(" newline { Spec } ")" newline` für alle Block-Direktiven (`require`, `exclude`, `replace`, `retract`, `tool`, `ignore`, `godebug`, Klammer-Form von `module`) deckt sich mit „Zeile endet auf `(`“ und „Zeile nur `)`“ ✓. |
| Lexik `json` gegen RFC 8259 | geprüft. Leerraum §2 ✓; Zeichenkette mit Escape und unmaskierten Steuerzeichen U+0000–U+001F §7 ✓; keine Kommentare ✓; BOM tolerieren §8.1 (MAY ignore) ✓. Dass Wert-Tokens nicht validiert werden, betrifft F-8. |
| Abnahme-Fragen §1 (konfigurierbar/benannt; JSON-Tiefe; Einzeiligkeit) | geprüft. Alle drei sind in Lastenheft, ADR und Spezifikation beantwortet und stimmen überein (benannt; Wurzel-Mitglied; `\n` für alle `shape-*`). Restpunkte: F-4, F-5. |
| Gate-Lauf | `make doc-check` Exit 0 („686 Datei(en) geprüft, 0 Befund(e)“). Geltungsbereich: Links, Anker, Kennungs-Linkpflicht. Prosa-Inhalt sieht das Gate nicht. |
| Kommentar-/Zustandsfeld-Regel (`AGENTS.md` §3.7) | geprüft, ohne Befund in den berührten Dateien. Die Historie-Zeilen tragen Zustand und Anlass, keine verworfene Alternative als Konjunktiv. |

## Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 1 |
| MEDIUM | 6 |
| LOW | 2 |
| INFO | 2 |

**Finding-Klassen dieses Laufs:** Messbeleg deckt die behauptete Menge nicht · Teil-Ablösung
benennt nicht alle berührten Punkte des Vorgängers · Akzeptanzkriterium sagt mehr zu, als die
Zerlegung trägt · Ausgabe-Escape nicht umkehrbar, Grenze nicht benannt · Zwei Stellen desselben
Vertrags regeln denselben Fall verschieden · Normalform ist kein Fixpunkt der Normalisierung,
Konsequenz nicht benannt · Plan-Abgrenzung durch die Entscheidung überholt, Plan nicht
nachgezogen · Normalisierung nur auf gültigen Eingaben injektiv, Grenze nicht benannt ·
Grundbegriff des Normalisierungsvertrags undefiniert · Werkzeugverhalten weicht vom
Referenz-Text ab · Übergangszustand in einem fortlaufenden Dokument

**Zu den beiden Register-Einträgen (Urteil bei der Closure, nicht hier):**

- *lexik-vertrag-ohne-sprach-gegenprobe*: Diesmal wurde gegen die Quelle abgeleitet. Ein **nicht
  fail-safe** Lexik-Fall (Code als Kommentar) wurde nicht gefunden. F-8, F-9 und F-10 sind
  fail-safe oder fail-closed.
- *ausgabe-einzeilig-nicht-zugesichert*: Die Einzeiligkeit ist jetzt zugesichert. F-4 und F-5
  liegen im selben Bereich.

## Verdikt

**Merge-blockierend:** ja, für die **Abnahme** (`Accepted`) von ADR-0042 und den Start von
slice-213. F-1 ist eine falsche Tatsachenbehauptung in einer ADR, die mit `Accepted` unveränderlich
wird. F-2 bis F-7 sind vor der Abnahme zu klären, weil slice-213 sonst Vertragsentscheidungen selbst
treffen müsste oder auf einem Plan aufsetzt, der der Entscheidung widerspricht. Die Commits liegen
bereits auf `main`; das Urteil betrifft den nächsten Lifecycle-Schritt, nicht einen offenen Merge.
Einen Rollen-Widerspruch zu einem HIGH gibt es bisher nicht, die Konflikt-Sequenz greift also
nicht.

**Übergabe:** Die Findings gehen an den Implementer (bei F-7 Rückkante Review → Plan). Die
Finding-Klassen gehen in die Closure §7 von slice-212. Dieser Report ist ein Lauf-Beleg und ersetzt
keine Verifikation gegen die DoD.
