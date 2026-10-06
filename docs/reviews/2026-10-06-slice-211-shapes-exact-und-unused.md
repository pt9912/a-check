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
