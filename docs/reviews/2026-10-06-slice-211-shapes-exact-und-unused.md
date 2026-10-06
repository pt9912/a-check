# Review-Report: slice-211 — 2026-10-06

**Review-Art:** Code — unabhängiger Lauf (frischer Subagent-Kontext ohne `fork`, kein Autor des
Gegenstands). Geprüft wird der Diff gegen Slice-Plan, ADR, Spezifikation und Hard Rules
(Maintainability); **nicht** gegen die DoD (Verifier, Modul 11).

**Gegenstand:** slice-211, Commit-Range `e5c0cd6..adc292d` (ein Commit `adc292d`)

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

- Slice-Plan slice-211 (§1 Ziel und Abgrenzung, §2, §3, §4, §6, §8), Welle welle-16
- ADR-0041 (Accepted), Punkte zu `exact`, `unused` und Wurzel-Grenze von `expect`
- AC-FA-RULE-012 (Beschreibung, Boundary `exact`/`unused`, Negative, Determinismus), AC-FA-CONF-001;
  SPEC-CONF-001 (Bullet `shapes`), SPEC-RULE-001 (Zeilen `shape-differs`, `shape-unused`, Absatz
  „außerhalb der Kette“), SPEC-EXTRACT-001 (Absatz „Sollform-Anweisungen“, nur als Bezug — der
  Normalisierer ist nicht Gegenstand)
- `AGENTS.md` §3 (Hard Rules 3.1–3.7)
- Frühere Reports am selben Bereich: slice-209 und slice-210 — für stabile Klassen-Bezeichnungen
  und um bekannte Fälle (slice-210 F-11: mehrzeilige Meldung aus Roh-Strings) nicht doppelt zu melden

**Messungen dieses Laufs** (Geltungsbereich je Messung genannt):

- `make test`, `make lint`, `make arch-check`, `make coverage-gate` auf `adc292d`: alle Exit 0;
  Coverage gesamt 96,40 % (Schwelle 90 %). Die Funktions-Coverage aus demselben Lauf trägt F-4:
  `expected` 85,7 %, `unusedFindings` 91,7 %, alle übrigen neuen/geänderten Funktionen in
  `core/shapes.go` 100 %. **Geltungsbereich:** Coverage zeigt, welche Zweige **ausgeführt**
  werden, nicht, ob eine Assertion sie prüft.
- Sonden: ein temporärer Test in einer **Kopie** des Repos (Scratch-Verzeichnis, `git archive
  HEAD`, nicht im Repo), ausgeführt über `make test` in der Kopie, über `cli.Run` Ende-zu-Ende.
  14 Fälle: Sollform-Datei als Symlink aus der Wurzel hinaus · Sollform unter `exclude` ·
  Sollform nicht zerlegbar · Sollform ist Verzeichnis · Datei und Sollform leer · Datei leer ·
  Sollform leer · Sollform = geprüfte Datei · dieselbe Datei in zwei `exact`-Einträgen ·
  `unused: fail` mit Doppel-Literal in einer Flow-Zeile, `"q()\r"`, Regex-Eintrag und Treffer in
  der **zweiten** Datei · Reihenfolge `shape-unused`/`shape-unlisted` · `expect` absolut ·
  `fehlt:` an der Zeilenzahl · `--print-graph` mit fehlender Sollform. **Geltungsbereich:**
  belegt ist das Verhalten des Binaries auf diesen Eingaben; die Normalisierung selbst
  (slice-210) ist nicht erneut sondiert.

---

## Findings

| ID | Kategorie | Befund | Quelle | Pfad | Verifizierbar | Klasse |
|---|---|---|---|---|---|---|
| F-1 | LOW | SPEC-RULE-001 setzt für `shape-unused` „Meldung = der Eintrag in seiner deklarierten Form“. Der Code formt diese Meldung dreifach um: abschließende Zeilenenden werden abgeschnitten, innere `LF` als die zwei Zeichen `\n` geschrieben, und ein Regex-Eintrag erhält das Suffix ` (regex)` (Sonde: `.a-check.yml:14: shape-unused: z\(\) (regex)`). Keine der drei Formen steht in SPEC-RULE-001, AC-FA-RULE-012 oder Handbuch §4; die Darstellung eines mehrzeiligen Eintrags (Handbuch-Beispiel `- \|` für `dependencies`) ist aus dem Vertrag nicht ableitbar. Zur Auftragsfrage: Die `\n`-Schreibung ist nicht vertragswidrig, aber auch nicht vertraglich bestimmt — der Vertrag lässt die Form offen. | SPEC-RULE-001 (`shape-unused`); AC-FA-RULE-012 (Determinismus-Kriterium, Form `pfad:zeile: <klasse>: <meldung>`) | `internal/hexagon/core/shapes.go:285-288` | ja — Sonde / Test auf die Meldung | Ausgabeform über den Vertrag hinaus festgelegt |
| F-2 | LOW | Der Kommentar zu `unusedFindings` sagt zu, ein Zeilenende im Eintrag werde als `\n` geschrieben, „so the finding stays one record per line“. Ersetzt wird nur `LF`; ein `CR` — nach SPEC-EXTRACT-001 ebenfalls ein Zeilenende — geht roh in die Ausgabe (Sonde: `allow: ["q()\r"]` ⇒ stdout-Zeile `.a-check.yml:14: shape-unused: q()\r`). Erreichbar nur über ein YAML-Escape in doppelten Anführungszeichen; YAML selbst normalisiert Zeilenumbrüche im Rohtext. | `AGENTS.md` §3.7 (Kommentar trägt eine Zusage); SPEC-EXTRACT-001 Schritt 1 (Definition Zeilenende) | `internal/hexagon/core/shapes.go:271-274`, `:285` | ja — Sonde | Ausgabeform einzeilig nicht zugesichert |
| F-3 | LOW | Die Sollform-Datei wird mit `os.ReadFile` gelesen und folgt damit einem Symlink aus der Scan-Wurzel hinaus (Sonde: `expect: s.kts`, `s.kts` → Datei außerhalb der Wurzel ⇒ Exit 0, Inhalt verglichen; bei Abweichung erschiene ihre erste Anweisung in `(erwartet: …)` auf stdout). Die `files`-Suche nimmt dagegen nur reguläre Dateien (`IsRegular`) und folgt keinem Symlink. SPEC-CONF-001 und ADR-0041 definieren „hinauszeigen“ **lexikalisch** (absolut oder `..`-führend); der Code hält diesen Wortlaut ein, die beiden Pfad-Arten desselben Blocks sind aber ungleich hermetisch. | SPEC-CONF-001 (Wurzel-Grenze `files`/`expect`, Begründung Hermetik); AC-QA-02 | `internal/adapter/driven/extract/shapes.go:111` gegen `:180` | ja — Sonde mit Symlink | Wurzel-Grenze nur lexikalisch, Pfad-Arten ungleich behandelt |
| F-4 | LOW | Vier Vertragsfälle des Diffs haben keinen Test, der bei ihrem Bruch rot würde: (i) nicht zerlegbare Sollform-Datei ⇒ Exit 2 (`expected` 85,7 %, der Zweig läuft in keinem Test); (ii) Regex-Kennzeichnung der `shape-unused`-Meldung (`unusedFindings` 91,7 %; der Regex-Eintrag im Test ist getroffen); (iii) Zählung „über alle Dateien des Eintrags“ — jeder `unused`-Test hat genau eine Datei; (iv) die neuen Gerüst-Zeilen (`unused: fail`, zweiter Eintrag `mode: exact`) — `TestPrintConfigShowsShapes` ist unverändert und prüft nur `allow-statements`. Das Verhalten selbst ist per Sonde richtig. | Slice-Plan slice-211 §3 (Zeile Tests); Mess-Regel „Eine Mutations-Probe belegt erst, wenn sie rot war“ | `internal/adapter/driven/extract/shapes.go:115-118`; `internal/hexagon/core/shapes.go:286-288`; `internal/hexagon/core/shapes_test.go` (`TestEvaluateShapesUnused`); `internal/cli/cli_shapes_test.go:135-144` | ja — `make coverage-gate` (Funktions-Coverage) bzw. Mutation | Vertragsfall ohne Test |
| F-5 | LOW | Der Kommentar zu `TestShapesUnusedEndToEnd` sagt, „der Befund zeigt auf die Zeile des Eintrags in der .a-check.yml“. Die Assertion prüft nur das Präfix `.a-check.yml:`, nicht die Zeilennummer; der Konfigurations-Test prüft `Allow[0].Line != 0`, der Kern-Test setzt die Zeile von Hand. Dass die YAML-Knotenzeile Ende-zu-Ende ankommt, prüft damit kein Test. | AC-FA-RULE-012 (Boundary `unused`); SPEC-RULE-001 (Verortung); `AGENTS.md` §3.7 | `internal/cli/cli_shapes_test.go:177-185`; `internal/adapter/driven/config/config_test.go:732` | ja — Assertion lesen | Testbeschreibung behauptet mehr als die Fixture prüft |
| F-6 | LOW | Die Exit-2-Aufzählung in Handbuch §4 („Exit-Code 2 statt eines stillen Grüns, wenn: …“) ist um die fehlende und die nicht zerlegbare Sollform-Datei ergänzt, nicht aber um die neuen Lade-Fälle des Diffs: `allow` oder `unused` bei `exact`, `expect` fehlt bei `exact` bzw. steht bei `allow-statements`, unbekannter `unused`-Wert, `expect` aus der Wurzel hinaus (die Liste nennt nur „ein Glob“). | SPEC-CONF-001 (Fail-closed beim Laden); `AGENTS.md` §4 (Aufzählung neben ihrer Quelle) | `docs/user/benutzerhandbuch.md:651-655` | nein — Doku-Abgleich | Aufzählung neben ihrer Quelle nicht nachgezogen |
| F-7 | LOW | Der Glossar-Eintrag „Sollform (`shapes`)“ definiert die Sollform weiter nur als Liste mit `shape-unlisted`; die Sollform-**Datei** von `mode: exact` und `shape-differs` fehlen. Der Diff hat die Glossar-Zeile direkt darüber um `shape-differs`/`shape-unused` ergänzt. | Benutzer-Doku; AC-FA-RULE-012 (zwei Modi) | `docs/user/benutzerhandbuch.md:969` | nein — Doku-Abgleich | Aufzählung neben ihrer Quelle nicht nachgezogen |
| F-8 | INFO | SPEC-RULE-001 sagt für `shape-differs` „**Genau ein** Befund je Datei“; SPEC-CONF-001 sagt, eine Datei in zwei Einträgen „wird in beiden geprüft“. Sonde: dieselbe Datei in zwei `exact`-Einträgen mit verschiedenen Sollformen ⇒ **zwei** `shape-differs` auf derselben Zeile. Der Code folgt SPEC-CONF-001; die beiden Sätze lesen sich ohne „je Eintrag“ widersprüchlich. | SPEC-RULE-001 (`shape-differs`) gegen SPEC-CONF-001 (zwei Einträge) | `spec/spezifikation.md:347` gegen `:94` | ja — Sonde | — (Hinweis an Planner, keine Klasse) |
| F-9 | INFO | Eine Sollform-Datei, die zugleich vom eigenen `files`-Glob getroffen wird, kann nie melden (Sonde: `expect: d/b.kts` für `files: [d/b.kts]` mit `evil()` ⇒ Exit 0). Der Vertrag schweigt dazu; für `forbidden_constructs` behandelt das Handbuch einen Eintrag, der „nie melden könnte“, als Exit 2. Ob dieselbe Haltung hier gilt, ist eine Vertragsfrage, keine Code-Abweichung. | AC-FA-RULE-012 („nie ein stilles Grün“); Handbuch §4 `forbidden_constructs` | `docs/user/benutzerhandbuch.md:497`; `internal/adapter/driven/extract/shapes.go:82-89` | ja — Sonde | — (Hinweis an Planner/Architect, keine Klasse) |

## Negativbefunde

| Bereich | Ergebnis |
|---|---|
| (a) `shape-differs` gegen SPEC-RULE-001 — erste Position, Zeile und Meldung in den drei Fällen anders / zusätzlich / fehlend | geprüft, ohne Befund — `firstDifference` bricht an der ersten Position ab; Suffixe ` (erwartet: …)` / ` (nicht in der Sollform)` und Präfix `fehlt: ` wortgleich; „fehlt“ an der Zeilenzahl (Sonde: Datei `a()` + zwei Leerzeilen ⇒ Zeile 3) |
| (a) Randfälle leere Datei / leere Sollform / beides leer | geprüft, ohne Befund — Datei nur Kommentar ⇒ `fehlt: a()` an Zeile 2 (Zeilenzahl); Sollform nur Kommentar ⇒ erste Anweisung `(nicht in der Sollform)`; beides leer ⇒ Exit 0 |
| (a) Datei in zwei Einträgen | geprüft, ohne Befund außer F-8 — byte-gleiche Befunde werden einmal ausgegeben, verschiedene bleiben getrennt |
| (b) `shape-unused`-Semantik „trifft in keiner Datei des Eintrags“ | geprüft, ohne Befund — `hit` je (Eintrag, `allow`-Index), über alle Dateien des Eintrags; Treffer nur in der zweiten Datei zählt (Sonde) |
| (b) Mehrere `allow`-Einträge mit gleichem Muster / Erst-Treffer | geprüft, ohne Befund — `markAllowed` markiert **jeden** passenden Eintrag; der Kern-Test bricht, wenn er beim ersten Treffer abbräche (zweiter, nur per Regex passender Eintrag wäre sonst unbenutzt) |
| (b) Verortung und Sortierung | geprüft, ohne Befund — `Path` = `.a-check.yml`, `Line` = YAML-Knotenzeile; Einordnung in die eine Totalordnung (Sonde: `.a-check.yml:15` vor `d/b.kts:1`); zwei gleiche Muster in einer Flow-Zeile ⇒ ein Befund (byte-gleich, SPEC-CONF-001) |
| (c) Fail-closed beim Laden | geprüft, ohne Befund — `expect` fehlt, `expect` absolut/`..` (Sonde `/etc/passwd` ⇒ Exit 2), `./` ⇒ leer ⇒ Exit 2, `allow`/`unused` bei `exact`, `unused: warn`, `expect` bei `allow-statements` haben je einen Pfad und einen Test |
| (c) Fail-closed beim Scan, vor jeder stdout-Ausgabe | geprüft, ohne Befund — fehlende, nicht zerlegbare und als Verzeichnis angegebene Sollform ⇒ Exit 2, stdout leer, stderr nennt `shapes[i]` und die Datei; Sollform vor den Globs des Eintrags, Einträge in Deklarationsreihenfolge (deterministisch) |
| (c) `--print-graph` (no-scan) | geprüft, ohne Befund — fehlende Sollform ⇒ Exit 0; SPEC-CONF-001 führt die fehlende `expect`-Datei unter „beim Scan“, nicht „beim Laden“ |
| (c) Sollform unter `exclude` | geprüft, ohne Befund — wird gelesen (Sonde Exit 0); der Widerspruchs-Fall des Vertrags gilt nur für geprüfte Dateien |
| (d) Tests auf Tautologie | geprüft, ohne Befund — Kern-Tests prüfen Zeile **und** Meldung je Fall (`Lines: 7` trennt „fehlt“ von einer Anweisungszeile); Ende-zu-Ende-Tests vergleichen stdout byte-genau. Lücken sind F-4 und F-5 |
| (e) CHANGELOG `[Unreleased]` gegen den Code | geprüft, ohne Befund — zwei Modi, drei Befunde, Exit-2-Fälle stimmen mit dem Verhalten |
| (e) Handbuch §3.4-Tabelle, §4 `unused`/`exact`, Historie 1.43 | geprüft, ohne Befund außer F-6 und F-7 — Meldungsformen und Beispiel `.a-check.yml:14: shape-unused: …` passen zum Code |
| (e) `--print-config`-Gerüst | geprüft, ohne Befund in der Form — beide Beispiele sind auskommentiert gültig; Test-Lücke unter F-4 |
| Plan §1 Abgrenzung (kein Diff aller Abweichungen, `unused` nur Opt-in, keine Normalisierer-Änderung) | geprüft, ohne Befund — `kotlin_shape.go`/`kotlin_split.go` nicht im Diff; die Sollform nutzt den Dialekt ihres Eintrags (§4-Rückführungsfrage damit nicht eingetreten) |
| Hard Rules `AGENTS.md` §3.1–§3.6 | geprüft, ohne Befund — kein `//nolint` im Diff; Spec-Straten und ADRs nicht berührt (Dateiliste des Commits); kein Move; Coverage-Gate unverändert 90 % |
| Kommentar-Regel `AGENTS.md` §3.7 (neue Kommentare in Code und Tests) | geprüft, ohne Befund außer F-2 und F-5 — Zusagen und Abgrenzungen; der entfernte `shapeModeKnown`-Kommentar („noch nicht implementiert“) hinterlässt keinen Rest |
| Architektur: Kern rein | geprüft, ohne Befund — `EvaluateShapes` erhält den Konfig-Pfad als Wert und liest nichts; `make arch-check` Exit 0 |

## Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 0 |
| MEDIUM | 0 |
| LOW | 7 |
| INFO | 2 |

**Finding-Klassen dieses Laufs:** Ausgabeform über den Vertrag hinaus festgelegt · Ausgabeform
einzeilig nicht zugesichert (schon bei slice-210 — **anderer** Vorgang, zählt neu) · Wurzel-Grenze
nur lexikalisch, Pfad-Arten ungleich behandelt · Vertragsfall ohne Test · Testbeschreibung behauptet
mehr als die Fixture prüft (schon bei slice-210 — anderer Vorgang, zählt neu) · Aufzählung neben
ihrer Quelle nicht nachgezogen

## Verdikt

**Merge-blockierend:** nein — kein HIGH, kein MEDIUM. Kein Finding ist gegen einen Konflikt mit
dem Implementer zu führen; die Konflikt-Sequenz (Modul 8) greift nicht.

**Übergabe:** F-1 bis F-7 gehen an den Implementer (annehmen oder begründen). F-8 und F-9 sind
Vertragsfragen und gehen an den Planner (F-9 ggf. an den Architect, falls ein neuer Exit-2-Fall
entsteht). Die Finding-Klassen gehen in die Closure-Notiz §7 von slice-211 und von dort in den
Zähler. Dieser Report ist Lauf-Beleg und ersetzt keine Verifikation (DoD, bewusstes Brechen der
Testbehauptungen: Verifier, Modul 11).

## Delta-Re-Review (Nachlauf)

**Gegenstand:** Range `9fc2c4d..5909911`. Darin `b5980ec` (Spezifikation 0.35.0) und `5909911`
(`fix(shapes): slice-211 Nachlauf`); `9fc2c4d` ist dieser Report selbst. Derselbe unabhängige
Kontext, derselbe Skill-Stand `a6d19b6`, Modell `claude-opus-5-5`, Datum 2026-10-06.

**Messungen** (Geltungsbereich je Messung genannt):

- `make test`, `make lint`, `make coverage-gate` auf `5909911`: alle Exit 0; Coverage gesamt
  96,5 %. In den Funktionen, die F-4 betraf: `expected` 92,3 %, `unusedFindings` 100 %.
- Sonden in einer zweiten **Kopie** (`git archive HEAD`, Scratch-Verzeichnis), über `cli.Run`.
  Sieben Fälle, gezielt gegen die beiden neuen Exit-2-Zusagen:
  1. Symlink-**Verzeichnis** im `expect`-Pfad
  2. `expect` gleich einer geprüften Datei in vier Schreibweisen (`d//b.kts`, `x/../d/b.kts`,
     `d/./b.kts`, `./d/b.kts`)
  3. `files` als Glob
  4. Hardlink

  **Geltungsbereich:** belegt ist das Verhalten auf diesen Eingaben. Weitere Pfad-Aliase (z. B.
  Groß-/Kleinschreibung auf nicht unterscheidenden Dateisystemen) sind nicht sondiert.

### Status der Findings F-1 bis F-9

| ID | Status | Beleg |
|---|---|---|
| F-1 | behoben | SPEC-RULE-001 (0.35.0) legt die `shape-unused`-Meldung jetzt fest: einzeilig, Zeilenenden als `\n`, Zusatz ` (regex)` (`spec/spezifikation.md`, Zeile `shape-unused`). Der Code entspricht dem (`internal/hexagon/core/shapes.go:285-289`) |
| F-2 | behoben | `CRLF` und `CR` werden vor der Schreibung zu `LF` (`internal/hexagon/core/shapes.go:285`); der Kommentar `:271-274` nennt alle drei Zeilenenden. Test `TestEvaluateShapesUnusedFormAndFiles` (`internal/hexagon/core/shapes_test.go:225`) |
| F-3 | **teilweise** | Ein Datei-Symlink als letzte Pfadkomponente ist jetzt Exit 2 (`internal/adapter/driven/extract/shapes.go:129-135`, Test `TestShapesExactExit2`). Ein Symlink **weiter vorn** im Pfad wird weiter verfolgt — siehe N-1 |
| F-4 | behoben | Nicht zerlegbare Sollform, Regex-Zusatz, Zählung über zwei Dateien und Gerüst sind jetzt getestet (`internal/cli/cli_shapes_test.go:194`, `:229`; `internal/hexagon/core/shapes_test.go:225`). Rest: `TestPrintConfigShowsExactAndUnused` prüft den Exit-Code nicht (geringfügig, kein eigenes Finding) |
| F-5 | behoben | Die erwartete Zeile wird aus der Konfiguration abgezählt und im Präfix geprüft (`internal/cli/cli_shapes_test.go:186-187`) |
| F-6 | behoben | Die Exit-2-Aufzählung nennt die Lade-Fälle für Schlüssel und Modus sowie die vier Sollform-Fälle (`docs/user/benutzerhandbuch.md:651-658`). Die Aussage zum Symlink darin ist N-1 |
| F-7 | behoben | Der Glossar-Eintrag nennt beide Modi (`docs/user/benutzerhandbuch.md:972`) |
| F-8 | behoben | SPEC-RULE-001: „Genau ein Befund je Eintrag und Datei“ (`spec/spezifikation.md`, Zeile `shape-differs`) |
| F-9 | **teilweise** | Spezifikation (SPEC-CONF-001, 0.35.0) und Handbuch erheben den Fall zu Exit 2. Der Code vergleicht aber nur die **Zeichenkette** (`rel == sh.Expect`, `internal/adapter/driven/extract/shapes.go:108`), und `expect` ist nur von `./` befreit, nicht normalisiert (`internal/hexagon/core/shapes.go:132`) — siehe N-2 |

### Neue Findings

| ID | Kategorie | Befund | Quelle | Pfad | Verifizierbar | Klasse |
|---|---|---|---|---|---|---|
| N-1 | HIGH | SPEC-CONF-001 (0.35.0) und Handbuch §4 sagen zu, dass bei der Sollform-Datei „ein Symlink nicht verfolgt wird — er könnte aus der Scan-Wurzel hinauszeigen“. `os.Lstat` prüft aber nur die **letzte** Pfadkomponente. Sonde: `expect: link/b.kts` mit `link` → Verzeichnis außerhalb der Wurzel ⇒ die fremde Datei wird gelesen und verglichen. Exit 1 mit `d/b.kts:1: shape-differs: evil() (erwartet: a())`: Der Inhalt der Datei außerhalb erscheint auf stdout. Adversarisch gegen das Binary geprüft. Der Code-Kommentar `:134` trägt dieselbe Zusage. | SPEC-CONF-001 (Fail-closed beim Scan, Symlink-Satz); AC-QA-02 (Hermetik); Reviewer-Skill HIGH „nachweislich falsche Tatsachenbehauptung“ | `internal/adapter/driven/extract/shapes.go:128-135`; `docs/user/benutzerhandbuch.md:657`; `spec/spezifikation.md:94` | ja — Sonde mit Verzeichnis-Symlink | Wurzel-Grenze nur lexikalisch, Pfad-Arten ungleich behandelt |
| N-2 | HIGH | SPEC-CONF-001 (0.35.0) und Handbuch §4 sagen zu, dass eine Sollform-Datei, die von den eigenen `files` getroffen wird, Exit 2 ist („der Eintrag könnte nie melden“). Der Vergleich ist reine Zeichenketten-Gleichheit gegen den nicht normalisierten `expect`-Wert. Sonde, `files: ["d/b.kts"]`, Datei `evil()`: `expect: d//b.kts`, `x/../d/b.kts`, `d/./b.kts` ⇒ jeweils **Exit 0, kein Befund** — genau das stille Grün, das die Zusage ausschließt. `./d/b.kts` und der Glob `d/*.kts` werden richtig abgewiesen. Adversarisch gegen das Binary geprüft. | SPEC-CONF-001 (Fail-closed beim Scan); AC-FA-RULE-012 („nie ein stilles Grün“); Reviewer-Skill HIGH „nachweislich falsche Tatsachenbehauptung“ | `internal/adapter/driven/extract/shapes.go:108`; `internal/hexagon/core/shapes.go:132`; `docs/user/benutzerhandbuch.md:657-658` | ja — Sonde mit drei Schreibweisen | Gleichheit über die Schreibweise statt über das Ziel |
| N-3 | LOW | Der Kopf des Slice-Plans sagt bei den berührten Spec-Stellen „keine Spec-Änderung“. Der Slice ändert die Spezifikation jetzt (0.35.0, `b5980ec`); der Plan ist nicht nachgezogen. | `AGENTS.md` §6 Schritt 4 (Plan-Änderung vor dem Code); Slice-Plan slice-211 Kopf | `docs/plan/planning/in-progress/slice-211-shapes-exact-und-unused.md:14-15` | ja — Plan-Kopf gegen `git log -- spec/` | Plan-Kopf nicht mit dem Vorgang nachgezogen |
| N-4 | LOW | Der CHANGELOG-Eintrag in `[Unreleased]` nennt als Exit-2-Fälle noch „Fehlende Datei oder Sollform-Datei, Widerspruch zu `exclude` und nicht zerlegbare Datei“. Die zwei neuen Vertragsfälle aus 0.35.0 fehlen: Sollform keine reguläre Datei, Sollform von den eigenen `files` getroffen. | `AGENTS.md` §6 Schritt 7 (CHANGELOG trägt die Vertragsänderung) | `CHANGELOG.md:20-21` | nein — Doku-Abgleich | Aufzählung neben ihrer Quelle nicht nachgezogen |
| N-5 | INFO | Ein **Hardlink** der geprüften Datei als Sollform bleibt still grün (Sonde: Exit 0). Er ist ein anderer Pfad, wird also nach dem Wortlaut nicht „von den files getroffen“, und ohne Inode-Vergleich ist er nicht erkennbar. Eine benannte Grenze dazu fehlt in Spezifikation und Handbuch. | AC-FA-RULE-012 („nie ein stilles Grün“); AC-QA-02 (ehrliche Grenze) | `internal/adapter/driven/extract/shapes.go:108` | ja — Sonde | — (Hinweis an Planner, keine Klasse) |

### Negativbefunde (Nachlauf)

| Bereich | Ergebnis |
|---|---|
| `b5980ec`: Spec-Stratum, Referenz-Richtung | geprüft, ohne Befund — keine ADR-, Slice- oder Review-Kennung im Spec-Text oder in der Historien-Zeile 0.35.0; Lastenheft unberührt |
| Aufteilung `Shapes` → `entryFiles`, Fehlerreihenfolge | geprüft, ohne Befund — Sollform vor den Globs, Einträge in Deklarationsreihenfolge, stdout leer bei Exit 2 |
| Kommentar-Regel `AGENTS.md` §3.7 (neue Kommentare) | geprüft, ohne Befund außer N-1 (Zusage `:134` zu weit) |
| Neue Tests auf Tautologie | geprüft, ohne Befund — Symlink-Test baut einen echten Datei-Symlink; der Zeilen-Test zählt die Zeile aus der Konfiguration statt sie zu setzen; der Mehrdatei-Test bräche, wenn `b()` nur in der ersten Datei gezählt würde |
| Hard Rules §3.1–§3.6 | geprüft, ohne Befund — kein `//nolint`, keine ADR berührt, Coverage-Schwelle unverändert |

### Summary (Nachlauf)

| Kategorie | Anzahl |
|---|---|
| HIGH | 2 |
| MEDIUM | 0 |
| LOW | 2 |
| INFO | 1 |

**Finding-Klassen des Nachlaufs:** Wurzel-Grenze nur lexikalisch, Pfad-Arten ungleich behandelt
(Wiederholung von F-3 im selben Vorgang — zählt nicht neu) · Gleichheit über die Schreibweise statt
über das Ziel · Plan-Kopf nicht mit dem Vorgang nachgezogen · Aufzählung neben ihrer Quelle nicht
nachgezogen (wie F-6/F-7, selber Vorgang)

### Verdikt (Nachlauf)

**Merge-blockierend:** ja — N-1 und N-2 (HIGH). Beide sind adversarisch gegen das Binary
verifiziert, nicht nur gelesen. In beiden Fällen sagt der neu geschriebene Vertrag (Spezifikation
0.35.0, Handbuch §4) mehr zu, als der Code einlöst. Der Eingriff ist klein und auf die zwei Zeilen
begrenzt, die die Pfade vergleichen bzw. prüfen. F-1, F-2 und F-4 bis F-8 sind behoben, F-3 und F-9
nur teilweise; die Restlücken sind N-1 und N-2. N-3 und N-4 blockieren für sich nicht.

**Übergabe:** N-1 bis N-4 an den Implementer; N-5 an den Planner (benannte Grenze oder nicht).
Ein Rollen-Widerspruch liegt nicht vor; die Konflikt-Sequenz (Modul 8) greift erst, wenn der
Implementer einem der HIGH-Findings widerspricht.

## Kurz-Gegenprüfung

**Gegenstand:** `beb04a8` (Plan-Kopf), `6bf7634` (Spezifikation 0.35.0, Wortlaut) und `00b87db`
(Code, Tests, Handbuch, CHANGELOG). Derselbe unabhängige Kontext, Skill-Stand `a6d19b6`, Modell
`claude-opus-5-5`, Datum 2026-10-06.

**Messungen** (Geltungsbereich je Messung genannt):

- `make test`, `make lint`, `make arch-check`, `make coverage-gate` auf `00b87db`: alle Exit 0;
  Coverage gesamt 96,4 % (`regularNoSymlink` 91,7 %, `expected` 90,0 %, `sameFile` 100 %).
- Zehn Sonden in einer dritten **Kopie** (`git archive HEAD`, Scratch-Verzeichnis), über `cli.Run`:
  - Symlink im `expect`-Pfad an erster und an mittlerer Stelle
  - `x/../d/b.kts` und `d//b.kts`
  - Hardlink
  - Scan-Wurzel selbst ein Symlink
  - vier `files`-Globs über ein Symlink-Verzeichnis: tiefer Literal-Präfix `link/sub/*.kts`,
    vollständig literal `link/sub/geheim.kts`, `link/**` und `*/sub/*.kts`

  **Geltungsbereich:** belegt ist das Verhalten auf diesen Eingaben. Wettläufe zwischen Prüfung und
  Lesen sind auftragsgemäß ausgenommen.

### Status

| ID | Status | Beleg |
|---|---|---|
| N-1 / F-3-Rest | behoben | `regularNoSymlink` prüft jeden Pfadbestandteil unterhalb der Wurzel per `Lstat` (`internal/adapter/driven/extract/shapes.go:99-120`), aufgerufen in `expected` (`:152`). Sonde: Symlink an erster (`link/b.kts`) und mittlerer Stelle (`s/link/b.kts`) ⇒ Exit 2, stdout leer, stderr nennt den Bestandteil. Test `TestShapesExactPathForms` (`internal/cli/cli_shapes_test.go:267-279`) |
| N-2 / F-9-Rest | behoben | `expect` wird nach der Wurzel-Prüfung mit `path.Clean` normalisiert (`internal/hexagon/core/shapes.go:137`), der Vergleich läuft über `sameFile` (`internal/adapter/driven/extract/shapes.go:131`, `:169-177`). Sonde: `x/../d/b.kts`, `d//b.kts` ⇒ Exit 2. Test `internal/cli/cli_shapes_test.go:248-254` |
| N-3 | behoben | Der Plan-Kopf nennt die Präzisierung der Spezifikation (`docs/plan/planning/in-progress/slice-211-shapes-exact-und-unused.md:15-18`) |
| N-4 | behoben | Der CHANGELOG nennt beide neuen Exit-2-Fälle (`CHANGELOG.md:20-22`) |
| N-5 | behoben (über den Hinweis hinaus) | `os.SameFile` erkennt den Hardlink (`internal/adapter/driven/extract/shapes.go:169-177`). Sonde und Test (`internal/cli/cli_shapes_test.go:255-266`) ⇒ Exit 2. Spezifikation 0.35.0 nennt den Hardlink-Fall ausdrücklich |
| F-4-Rest | behoben | `TestPrintConfigShowsExactAndUnused` prüft den Exit-Code (`internal/cli/cli_shapes_test.go:231-233`) |

### Adversariale Fragen

| Frage | Ergebnis |
|---|---|
| Ist die Scan-Wurzel selbst ein Symlink? | ohne Befund — `regularNoSymlink` prüft nur Bestandteile **unterhalb** der Wurzel. Die Wurzel ist Eingabe des Aufrufers und wird aufgelöst. Sonde: Wurzel als Symlink ⇒ regulärer Lauf, Exit 1 mit dem erwarteten `shape-differs` |
| Folgt ein `files`-Glob über ein Symlink-Verzeichnis im Präfix hinaus? | **ja, wenn der Symlink nicht der letzte Bestandteil des Literal-Präfixes ist** — siehe G-1. Ist der Symlink der ganze Präfix (`link/**`) oder liegt er unter einem Wildcard-Segment (`*/sub/*.kts`), wird er nicht verfolgt (Exit 2, „trifft keine Datei“): `WalkDir` liest seinen Start mit `Lstat`, und ein Symlink als Start ist kein Verzeichnis |

### Neues Finding

| ID | Kategorie | Befund | Quelle | Pfad | Verifizierbar | Klasse |
|---|---|---|---|---|---|---|
| G-1 | MEDIUM | `globFiles` startet den Walk am Literal-Präfix des Globs. Liegt ein Symlink-Verzeichnis **vor** dessen letztem Bestandteil (`files: ["link/sub/*.kts"]` oder `["link/sub/geheim.kts"]`, `link` → Verzeichnis außerhalb der Wurzel), löst das Betriebssystem den Symlink beim Pfad-Zugriff auf. Der Walk läuft dann im fremden Baum, und die Datei dort wird gelesen und bewertet. Sonde: Exit 1, `link/sub/geheim.kts:1: shape-unlisted: secretToken()` — fremder Inhalt auf stdout. Für `expect` schließt 0.35.0 genau diesen Weg mit der Begründung „er könnte aus der Scan-Wurzel hinauszeigen“; für `files` gilt dieselbe Begründung, aber die Spezifikation sagt nur die **lexikalische** Grenze zu. Der Code stammt aus slice-210, nicht aus diesem Diff. Der Negativbefund „Hermetik der Dateisuche“ im slice-210-Report („Symlinks werden weder als Start-Verzeichnis verfolgt …“) gilt damit nur für den Symlink als ganzen Start, nicht für einen Symlink vor dem letzten Bestandteil. | SPEC-CONF-001 (Wurzel-Grenze für `files` und `expect`, Begründung Hermetik); AC-QA-02 | `internal/adapter/driven/extract/shapes.go` (`globFiles`: `os.Stat` auf den Präfix, `filepath.WalkDir` ab dort) | ja — Sonde mit Symlink-Verzeichnis vor dem letzten Präfix-Bestandteil | Wurzel-Grenze nur lexikalisch, Pfad-Arten ungleich behandelt |

### Negativbefunde (Kurz-Gegenprüfung)

| Bereich | Ergebnis |
|---|---|
| `6bf7634`: Spec-Stratum | geprüft, ohne Befund — keine Kennung von Slice, Review oder ADR im Spec-Text. Die Zeile 0.35.0 ist innerhalb desselben Slice nachgeschärft, ohne Versions-Sprung; kein Release dazwischen |
| Neue Kommentare (`AGENTS.md` §3.7) | geprüft, ohne Befund — Zusagen. Der Zusatz „(Review slice-211 N-…)“ folgt dem Bestand (`kotlin_shape.go:272`, „Review slice-210 F-2“) |
| Neue Tests auf Tautologie | geprüft, ohne Befund — Schreibweisen, Hardlink und Symlink-Verzeichnis bauen echte Dateisystem-Zustände; die Assertions prüfen Exit-Code, leeres stdout und die Ursache auf stderr |
| Reihenfolge der Prüfungen | geprüft, ohne Befund — `expected` (Symlink-Prüfung) läuft vor `entryFiles`; `sameFile` folgt mit `os.Stat` nur noch einem Pfad, der als symlinkfrei geprüft ist |

### Verdikt (Kurz-Gegenprüfung)

**Merge-blockierend für slice-211:** nein. N-1 bis N-5 sind behoben, ebenso die Reste von F-3,
F-4 und F-9; alle sind per Sonde gegen das Binary nachgeprüft. G-1 ist MEDIUM: Es verlangt vor der
Closure einen **Ausgang** — Aufnahme in diesen Slice oder ein Folge-Slice mit Kennung —, aber keine
Änderung am Diff dieses Slice, aus dem es nicht stammt. Bleibt es ohne Ausgang, wäre die
Hermetik-Zusage für `expect` strenger als für `files`, ohne dass die Asymmetrie benannt ist.

**Übergabe:** G-1 an den Planner (Ausgang zuweisen). Die Klasse von G-1 gehört zu F-3 und N-1 —
derselbe Vorgang slice-211, zählt im Register einmal.

## G-1-Bestätigung

**Gegenstand:** `abc4f50` (Plan-Änderung), `55dbfd2` (Spezifikation 0.35.0, `files`-Satz) und
`1d63905` (Code, Test, Handbuch, CHANGELOG). Derselbe unabhängige Kontext, Skill-Stand `a6d19b6`,
Modell `claude-opus-5-5`, Datum 2026-10-06.

**Messungen** (Geltungsbereich je Messung genannt):

- `make test`, `make lint`, `make coverage-gate` auf `1d63905`: alle Exit 0; Coverage 96,4 %.
- **Mutations-Gegenprobe** in einer vierten Kopie: `dirNoSymlink` gibt immer `true, nil` zurück ⇒
  `TestShapesFilesGlobSymlinkPrefix` ist **rot**, und zwar aus dem richtigen Grund. Meldung
  `link/sub/geheim.kts:1: shape-unlisted: secretToken()` für die Globs `link/sub/*.kts` und
  `link/sub/geheim.kts`. Der dritte Glob `link/**` bleibt unter der Mutation grün — er war nie
  verwundbar.
- Elf Sonden in derselben Kopie, über `cli.Run`. Jede prüft Exit-Code und ob der fremde Inhalt
  (`secretToken`) auf stdout oder stderr erscheint. **Geltungsbereich:** Pfade über
  `files`/`expect`; Wettläufe zwischen Prüfung und Lesen sind auftragsgemäß ausgenommen.

### Status G-1

**Behoben.**

- `globFiles` prüft den literalen Präfix vor dem Walk mit `dirNoSymlink`
  (`internal/adapter/driven/extract/shapes.go:211-221`).
- `dirNoSymlink` liest jeden Bestandteil per `Lstat`; ein Symlink ist ein Fehler
  (`internal/adapter/driven/extract/shapes.go:253-272`).
- Unterhalb des Präfixes folgt `WalkDir` keinem Symlink und nimmt nur reguläre Dateien (`:241`).
- Test: `TestShapesFilesGlobSymlinkPrefix` (`internal/cli/cli_shapes_test.go:285`).
- Spezifikation 0.35.0 (SPEC-CONF-001) und Handbuch §4 (`docs/user/benutzerhandbuch.md:652`)
  nennen den Fall.

### Adversariale Sonden

| Weg | Ergebnis |
|---|---|
| Glob ohne Wildcard, Symlink-Datei an der Wurzel (`link.kts` → Datei außerhalb) | Exit 2 („trifft keine Datei“), kein Austritt |
| `..` im Präfix (`d/../link/sub/*.kts`) | Exit 2 — `filepath.Join` löst `..` auf, `link` wird als Symlink erkannt |
| `.` im Präfix (`x/./link/sub/*.kts`) | Exit 2, Symlink `x/link` erkannt |
| `**` am Anfang (`**/geheim.kts`, `**/*.kts`) | kein Austritt — Start an der Wurzel, `WalkDir` folgt keinem Symlink; Exit 2 bzw. Exit 0 je nach übrigen Treffern |
| Symlink-Verzeichnis unterhalb des Präfixes (`x/**`, `x/link` → außerhalb) | kein Austritt, Exit 0 über die reguläre Datei |
| Symlink-**Datei** unterhalb des Präfixes (`x/*.kts`, `x/l.kts` → Datei außerhalb) | kein Austritt; die Datei wird **still übergangen**, siehe H-1 |
| Klammer- und Mengen-Syntax im Präfix (`lin[k]/…`, `{link,y}/…`) | Exit 2 („trifft keine Datei“) — der Präfix wird wörtlich gesucht und fehlt; fail-closed |
| `exclude` | ohne Dateisystem-Zugriff (reiner Glob-Vergleich auf `rel`), kein Weg hinaus |
| `expect` mit `..` vor dem Symlink (`d/../link/b.kts`) | Exit 2 — `path.Clean` ergibt `link/b.kts`, Symlink erkannt |

In keiner der elf Sonden erscheint der fremde Inhalt auf stdout oder stderr.

### Neues Finding

| ID | Kategorie | Befund | Quelle | Pfad | Verifizierbar | Klasse |
|---|---|---|---|---|---|---|
| H-1 | INFO | Eine **Symlink-Datei**, die ein `files`-Glob unterhalb des Präfixes trifft, wird nicht geprüft und nicht gemeldet. Sonde: `x/*.kts` mit `x/l.kts` als Symlink ⇒ Exit 0 über die übrigen Treffer. Ist sie der einzige Treffer, kommt Exit 2 „trifft keine Datei“, obwohl eine Datei dieses Namens existiert (Sonde `link.kts`). SPEC-CONF-001 (0.35.0) benennt die Grenze („nimmt nur reguläre Dateien“); das Handbuch nennt sie nur für den Präfix, nicht für den Treffer selbst. Kein Hermetik-Bruch — fail-closed bzw. benannt —, aber ein Glob kann eine Build-Datei, die als Symlink abgelegt ist, still auslassen. | SPEC-CONF-001 (0.35.0, `files`-Satz); AC-FA-RULE-012 („nie ein stilles Grün“) | `internal/adapter/driven/extract/shapes.go:241`; `docs/user/benutzerhandbuch.md:651-653` | ja — Sonde | — (Hinweis an Planner, keine Klasse) |

### Verdikt (G-1-Bestätigung)

**Merge-blockierend:** nein. G-1 ist behoben; der Test ist per Mutation rot aus dem richtigen
Grund. Über `files` und `expect` führt in keiner sondierten Form ein Weg aus der Scan-Wurzel.
H-1 ist ein Hinweis ohne erwartete Aktion im Code. Ob das Handbuch die im Vertrag schon benannte
Grenze auch für den Treffer selbst nennt, entscheidet der Planner.
