# Review-Report: slice-209 — 2026-10-06

**Review-Art:** Plan / Design — unabhängiger Lauf (frischer Subagent-Kontext ohne `fork`, kein
Autor des Gegenstands). Geprüft wird der Vertrags-Diff gegen Slice-Plan, ADRs und Konventionen;
**nicht** gegen die DoD (Verifier, Modul 11).

**Gegenstand:** slice-209, Commit-Range `02fc35e..bab57c6` (ein Commit `bab57c6`)

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

- Slice-Plan slice-209 (§1 mit den zehn Abnahme-Entscheiden, §2, §3, §6), Welle-Plan welle-16
- ADR-0041 (Proposed, Gegenstand), ADR-0027 und ADR-0034 (beide Accepted, im `Bezug:`)
- AC-FA-RULE-012 (neu), AC-FA-RULE-011, AC-FA-CONF-001, AC-QA-01, AC-QA-02;
  SPEC-CONF-001, SPEC-EXTRACT-001, SPEC-RULE-001, SPEC-DET-001; ARC-001, ARC-003
- `AGENTS.md` §3 (Hard Rules, insb. 3.4, 3.5, 3.7), `harness/conventions.md`
  §Anforderungs-Anlege-Prozess
- Eigen-Konfiguration `.a-check.yml` (für die Hexagon-Frage in F-5)
- Folge-Slices slice-210, slice-211 (nur Plan-Bezug: ob sie ohne eigene Vertragsentscheidung
  implementieren können)
- `v6.13.0` · `regelwerk/modul-08-agentenrollen.md` §Konflikt-Pfad als Rollen-Sequenz

---

## Findings

| ID | Kategorie | Befund | Quelle | Pfad | Verifizierbar | Klasse |
|---|---|---|---|---|---|---|
| F-1 | HIGH | Der CHANGELOG-Eintrag sagt, er „beschreibt den **abgenommenen** Vertrag“. Im Repo ist dieser Vertrag nicht abgenommen: ADR-0041 trägt `Status: Proposed`, welle-16 §4 macht die Abnahme der Anforderung zur Startbedingung von slice-210/211, und laut DoD in slice-209 §2 wird die ADR erst nach der Abnahme durch den Maintainer `Accepted`. Abgenommen sind die zehn Entscheide aus §1, nicht der Vertragstext, und der weicht davon ab (F-2). | Skill §HIGH „nachweislich falsche Tatsachenbehauptung“; ADR-0041 Kopf; welle-16 §4 | `CHANGELOG.md:17` | nein — kein Gate vergleicht CHANGELOG-Prosa mit dem ADR-Status | Zustandsbehauptung greift einem ausstehenden Abnahme-Schritt vor |
| F-2 | MEDIUM | Die Anweisungsgrenze in ADR-0041 Punkt 4 und SPEC-EXTRACT-001 Schritt 3 weicht vom abgenommenen Entscheid 5 ab. Ein Zeilenende trennt dort auch **innerhalb** von `{…}` (geschrieben als `;`), und Bedingung (c) ist neu: Endet eine Zeile auf ein Operatorzeichen, `,` oder `.`, folgt keine Grenze. Entscheid 5 lässt Zeilenenden nur auf oberster Ebene trennen und nur ohne offene `{`. Die ADR kennzeichnet nur die Schärfung von Entscheid 4 (Punkt 3) als Abweichung, diese nicht. Slice-Plan §1 nennt Entscheid 5 weiterhin in der alten Fassung. | slice-209 §1 Entscheid 5; `AGENTS.md` §6 Schritt 4 (Plan-Änderung vor dem Code, nicht still) | `docs/plan/adr/0041-shapes-sollform-je-datei.md:72-79`; `spec/spezifikation.md:284-292`; `docs/plan/planning/in-progress/slice-209-shapes-spec-first.md:54-58` | nein — inferentiell (Entscheid ↔ Vertragstext) | Abweichung vom abgenommenen Entscheid ohne Kennzeichnung |
| F-3 | MEDIUM | Die Lexik soll vollständig sein („Alles andere ist Code“), kennt aber keine Multi-Dollar-Zeichenketten (`$$"…"`, ab Kotlin 2.2 stabil). Darin ist `${` Text, für den Lexer öffnet es eine Vorlage. Konkrete Eingabe: `version = "x" + $$"${" + '"' + '}' + "// "; dependencies.add("implementation", "com.evil:lib:1.0")`. Kotlin führt zwei Anweisungen aus. Der Lexer schließt die äußere Zeichenkette dagegen am `"` vor `//`, liest `// …` als Zeilen-Kommentar und **verwirft** den `dependencies.add`-Aufruf. Übrig bleibt eine Anweisung `version="x"+$$"${"+'"'+'}'+"`. Trifft ein `allow`-Regex sie (etwa `version=".*"`), bleibt die Datei grün und die Fremdabhängigkeit taucht in keiner Ausgabe auf. Das widerspricht der Linie „nie still grün“ in AC-FA-RULE-012 und der ADR-0034-Leitplanke „im Zweifel weniger verschlucken“. | AC-FA-RULE-012 (fail-safe); AC-QA-02 (ausgewiesene Grenze); ADR-0034 | `spec/spezifikation.md:266-275` | ja — sobald slice-210 läuft: ein Testfall mit obiger Eingabe in `make test` | Lexik-Aufzählung behauptet Vollständigkeit, die die Zielsprache nicht trägt |
| F-4 | MEDIUM | Ein voll verankerter Regex mit `.*` innerhalb einer Zeichenkette überspannt Code. Beispiel: `version="1"+run{dependencies.add("implementation","x:y:1")}+""` trifft `^(?:version=".*")$`, und `run{}` wird ausgeführt. Lastenheft-Out-of-Scope und welle-16 §6 führen den Regex-Eintrag als Ersatz für das gestrichene `ignore: [version-literals]` an. Der gestrichene Ausschluss hatte als Grund „ein unscharf umrissener Ausschluss ist die Lücke“, und bei der naheliegenden Schreibweise hat der Ersatz dieselbe Lücke. ADR-0041 Punkt 5 und die Konsequenzen nennen diese Grenze nicht. | AC-QA-02 (ehrliche Heuristik-Grenze); ADR-0041 §Konsequenzen; welle-16 §6 | `spec/lastenheft.md:424`; `docs/plan/adr/0041-shapes-sollform-je-datei.md:80-83` | nein — Grenz-Benennung ist Urteil | fail-safe-Zusage ohne benannte Grenze |
| F-5 | MEDIUM | Ort und Zeitpunkt der Normalisierung von `literal`-Einträgen widersprechen sich. SPEC-CONF-001 sagt „wird beim **Laden** … normalisiert“ und führt Zerlegungsfehler als „Fail-closed beim Laden“ (Konfigurations-Adapter, ARC-004). ADR-0041 Punkt 9 und der Schlusssatz von SPEC-EXTRACT-001 legen dieselbe Arbeit in die **Extraktion** (ARC-003). Ein Konfigurations-Adapter, der den Normalisierer der Extraktion importiert, wäre im Eigen-Dogfooding (`adapter_sink: driver-common`) eine `lateral-adapter`-Kante. Offen bleibt auch, ob der no-scan-Pfad `--print-graph` (Config.Load → Extraction.Validate) einen nicht zerlegbaren `literal`-Eintrag melden muss. | ADR-0041 Punkt 9; SPEC-CONF-001; ARC-003/ARC-004; SPEC-CLI-002 (Ladezeit-Parität) | `spec/spezifikation.md:94`; `spec/spezifikation.md:302`; `docs/plan/adr/0041-shapes-sollform-je-datei.md:94-99` | teilweise — eine Implementierung im Konfigurations-Adapter würde `make arch-check` rot färben; die Mehrdeutigkeit selbst fängt kein Gate | Vertrag legt dieselbe Arbeit an zwei Hexagon-Orte |
| F-6 | MEDIUM | Der Vertrag sagt nicht, ob „nicht kompilierbare Regex“ am Roh-Muster oder an der Hülle `^(?:…)$` geprüft wird. Das Roh-Muster `a)\|(.*` kompiliert allein nicht, umhüllt als `^(?:a)\|(.*)$` aber schon, und dann trifft es **jede** Anweisung. Die Wahl ist eine eigene Vertragsentscheidung in slice-210 mit fail-safe-Folge. | SPEC-CONF-001 (Verankerung, Exit-2-Fälle); slice-209 §1 Ziel („ohne eigene Vertragsentscheidung“) | `spec/spezifikation.md:94` | ja — Testfall in `make test`, sobald slice-210 läuft | Verankerungs-Zusage ohne Kompilier-Reihenfolge |
| F-7 | LOW | AC-FA-RULE-012 nennt zum Dialekt `kotlin` die Endungen „(`*.kt`, `*.kts`)“. SPEC-CONF-001 sagt nicht, ob das eine Bedingung ist, etwa Exit 2 für `files: [go.mod]` mit `dialect: kotlin`, oder nur eine Beschreibung. | AC-FA-RULE-012; SPEC-CONF-001 | `spec/lastenheft.md:400` | nein | Lastenheft-Detail ohne Präzisierung in der Spezifikation |
| F-8 | LOW | Begriffe, an denen Zeilen-Mapping und Faltung hängen, sind unbestimmt: welche Zeichen „Leerraum“ und „Zeilenende“ sind (CR allein, CRLF, Unicode-Leerraum, BOM); was „erstes Code-Zeichen“ bei einer Anweisung heißt, die mit Zeichenkette, Zeichen-Literal oder Backtick-Bezeichner beginnt (laut Lexik sind diese **kein** Code); was „letzte Zeile der Datei“ bei leerer Datei oder abschließendem Zeilenumbruch bedeutet. | slice-209 §1 Entscheid 6; SPEC-EXTRACT-001; SPEC-RULE-001 `shape-differs` | `spec/spezifikation.md:275-292`; `spec/spezifikation.md:326` | nein | Grundbegriff des Normalisierungsvertrags undefiniert |
| F-9 | LOW | Laut Determinismus-Kriterium von AC-FA-RULE-012 haben Befunde die Form `pfad:zeile: <klasse>: <anweisung>`. SPEC-RULE-001 setzt bei `shape-differs` aber Suffixe ` (erwartet: …)` / ` (nicht in der Sollform)` bzw. das Präfix `fehlt: …`. `shape-unused` meldet den Konfig-Eintrag statt einer Anweisung. Die Lastenheft-Beschreibung nimmt nur `shape-unused` aus. | AC-FA-RULE-012 (Determinismus-Kriterium); SPEC-RULE-001 | `spec/lastenheft.md:416`; `spec/spezifikation.md:326-327` | nein | Befund-Form in Lastenheft und Spezifikation unterschiedlich gefasst |
| F-10 | LOW | Trifft ein Glob dieselbe Datei in zwei Einträgen, wird sie laut SPEC-CONF-001 „in beiden geprüft“. Die Meldung nennt den Eintrag nicht, also entstehen byte-gleiche Befundzeilen. SPEC-DET-001 sagt nichts zu Duplikaten. Offen ist auch der Fall, dass zwei Globs **eines** Eintrags dieselbe Datei treffen. | SPEC-CONF-001; SPEC-DET-001 | `spec/spezifikation.md:94`; `spec/spezifikation.md:577` | nein | Duplikat-Semantik bei überlappenden Globs offen |
| F-11 | LOW | Der Kern wertet jetzt neben Importen und Roh-Text auch ein drittes Modell aus (Anweisungsfolge). ARC-001 („wertet die sieben Regeln auf einem abstrakten Import-/Schicht-Modell aus“) und SPEC-RULE-001 („Präzisiert die sieben Hexagon-Regeln“) sind trotzdem nicht nachgezogen, obwohl der Commit beide Dateien ändert. Die Zahl „sieben“ war schon vorher veraltet: Das Lastenheft führt 12 `AC-FA-RULE`-Überschriften. | `AGENTS.md` §4 (abschließende Aufzählung neben wachsender Menge) | `spec/architecture.md:71`; `spec/spezifikation.md:306` | nein | abschließende Aufzählung neben wachsender Menge |
| F-12 | LOW | `expect` ist ein „Pfad relativ zur Scan-Wurzel“. Ob `..` oder absolute Pfade zulässig sind, bleibt offen, und damit auch, ob die Regel außerhalb des Prüfbaums liest. | SPEC-CONF-001; AC-QA-02 (Hermetik) | `spec/spezifikation.md:94` | nein | Pfad-Escape nicht ausgeschlossen |
| F-13 | INFO | Bedingung (c) verbindet auch dann zwei Anweisungen zu einer, wenn eine Zeile auf ein Postfix-`!!`/`++`/`--` oder ein schließendes `>` endet. Beispiel: `val v = findProperty("x")!!`, darunter `dependencies { … }` → eine Anweisung, der erlaubte Block trifft nicht, die Datei wird rot. Das ist fail-safe, aber eine Fehlalarm-Quelle, die zu Risiko 1 in slice-209 §6 passt. Ein Fall für die Testauswahl von slice-210. | slice-209 §6 Risiko 1; ADR-0041 Re-Evaluierungs-Trigger 2 | `spec/spezifikation.md:288-289` | ja — Testfall in `make test` (slice-210) | Fortsetzungsregel erzeugt Fehlalarm an legitimer Schreibweise |
| F-14 | INFO | Der Satz „Noch nicht implementiert … lehnt ihn bis dahin … ab (Exit 2)“ im CHANGELOG beschreibt einen Übergangszustand. Er wird falsch, sobald slice-210 liefert. slice-210 und slice-211 planen eine CHANGELOG-Fortschreibung, nennen dieses Löschen aber nicht ausdrücklich. | `AGENTS.md` §6 Schritt 7 | `CHANGELOG.md:17-19` | nein | Übergangszustand in einem fortlaufenden Dokument |

## Negativbefunde

| Bereich | Ergebnis |
|---|---|
| Referenz-Richtung der Spec-Straten (`AGENTS.md` §3.4) | geprüft, ohne Befund — Geltungsbereich: alle hinzugefügten Zeilen unter `spec/` im Diff, gesucht nach `ADR`, `slice-`, `welle` und Hash-Mustern. Keine Treffer, auch nicht in den neuen Historie-Zeilen (Lastenheft 0.28.0, Spezifikation 0.33.0). Prosa-Verweise ohne diese Token sieht die Suche nicht; dafür wurde der Text gelesen. |
| Historie-Regel der Spezifikation („kein ADR- und kein Slice-Verweis“) | geprüft, ohne Befund. Die 0.33.0-Zeile steht über 0.32.0, wie im vorhandenen absteigenden Block. |
| AC-Form AC-FA-RULE-012 (`conventions.md` §Anforderungs-Anlege-Prozess) | geprüft, ohne Befund — expliziter Anker `ac-fa-rule-012`, Happy, vier Boundary-Fälle, Negative, Determinismus, Out-of-Scope mit Begründungen; Versions-Bump und Historie-Zeile vorhanden. Die Spalte, ob jedes Kriterium Given/When/Then trägt, wurde gelesen: ja. |
| AC-FA-CONF-001-Erweiterung | geprüft, ohne Befund — Schema-Satz, Exit-2-Verweis und neues Negative-Kriterium. Die zusätzlichen Exit-2-Fälle in SPEC-CONF-001 (`unused` bei `exact`, `literal` mit mehr als einer Anweisung, unbekannter `match`) **präzisieren** das Schema. Sie erweitern es nicht, denn das Lastenheft bindet `unused` an `allow-statements` und `literal` an „Anweisung“. |
| Gegenprobe-Fälle des CR (Boundary „versteckte Anweisung“, „Kommentar und Zeichenkette“, `exact`, `unused`, Negative) gedanklich gegen SPEC-EXTRACT-001 | geprüft, ohne Befund **für die gelisteten Formen**. Weitere Anweisung oben, im Block, `plugins` mit `{` auf der Folgezeile, `compileClasspath += files(…)`, `add("x", run { … })`, „im Kommentar begonnen“ (`/* dependencies { */ implementation("x")`): jeweils `shape-unlisted`. Eine Form mit unbalancierter Klammer nach dem Kommentar ist kein gültiges Kotlin und landet in Exit 2. Ausnahme ist die nicht gelistete Form aus F-3. |
| Kollision zweier Quellen auf dieselbe normalisierte Form (Leerraum-Faltung, Grenzregel) | geprüft, ohne Befund — Geltungsbereich: Faltung (Wort/Wort, Operator/Operator, Kommentar als Leerraum, `val a` gegen `vala`) und Grenzregel (oberste Ebene, `;` im Block, (a)–(c)). Beide entfernen nur Leerraum und Kommentare und fügen Grenzen hinzu. Einer erlaubten Form lässt sich so kein Code-Token hinzufügen. Die Lexik ist davon getrennt geprüft (F-3). |
| Lexik-Fälle Roh-Strings, verschachtelte Block-Kommentare, `${…}` mit Klammern, Zeichen-Literal `'{'`, Backtick-Bezeichner, Shebang | geprüft, ohne Befund außer F-3. Das Ende eines Roh-Strings (Lauf von mindestens drei `"`, überzählige vorne) stimmt mit Kotlin überein. Wird Code fälschlich als Zeichenkette gelesen, bleibt er byte-genau sichtbar. Gefährlich ist nur Code, der als Kommentar verworfen wird. |
| ADR-0041-Form (`modul-04` Ziel-Form) | geprüft, ohne Befund — `Bezug:`, `Schärft:` mit SPEC-/ARC-Kennungen, sechs Alternativen, Konsequenzen positiv/negativ, Fitness Function mit existierenden Targets (`make test`, `make arch-check`), drei Re-Evaluierungs-Trigger, Geschichte. ADR-0027 und ADR-0034 im `Bezug:` sind `Accepted`. |
| ADR-Index | geprüft, ohne Befund — Zeile ADR-0041 vorhanden, Status `Proposed` stimmt mit dem Kopf überein. |
| Kommentar-/Zustandsregel (`AGENTS.md` §3.7) im Diff | geprüft, ohne Befund — die YAML-Kommentare im Schema-Beispiel tragen Zusagen (literal normalisiert, RE2 verankert). Keine Chronik in Zustandsfeldern; der Übergangszustand im CHANGELOG steht als F-14 da, nicht als §3.7-Verstoß (kein Kommentar, kein Zustandsfeld). |
| `spec/architecture.md` ARC-003 sprach-/meilensteinfrei | geprüft, ohne Befund — „je Dialekt“ als Rolle, keine Technik, kein Slice-Bezug. |
| Versions-Kohärenz (Lastenheft 0.28.0, Spezifikation 0.33.0, Architektur 0.5.0, `docs/user/releasing.md`) | geprüft, ohne Befund. |
| Repo-weite Link-/Anker-Hygiene | `make doc-check` Exit 0, „659 Datei(en) geprüft, 0 Befund(e)“ — Geltungsbereich: Links, Anker, Kennungs-Linkpflicht; über die Vertrags-Semantik sagt das nichts. |

## Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 1 |
| MEDIUM | 5 |
| LOW | 6 |
| INFO | 2 |

**Finding-Klassen dieses Laufs:** Zustandsbehauptung greift einem ausstehenden Abnahme-Schritt
vor · Abweichung vom abgenommenen Entscheid ohne Kennzeichnung · Lexik-Aufzählung behauptet
Vollständigkeit, die die Zielsprache nicht trägt · fail-safe-Zusage ohne benannte Grenze ·
Vertrag legt dieselbe Arbeit an zwei Hexagon-Orte · Verankerungs-Zusage ohne
Kompilier-Reihenfolge · Lastenheft-Detail ohne Präzisierung in der Spezifikation ·
Grundbegriff des Normalisierungsvertrags undefiniert · Befund-Form in Lastenheft und
Spezifikation unterschiedlich gefasst · Duplikat-Semantik bei überlappenden Globs offen ·
abschließende Aufzählung neben wachsender Menge · Pfad-Escape nicht ausgeschlossen ·
Fortsetzungsregel erzeugt Fehlalarm an legitimer Schreibweise · Übergangszustand in einem
fortlaufenden Dokument

## Verdikt

**Merge-blockierend:** ja — F-1 (HIGH) und F-2 bis F-6 (MEDIUM). Wiegt nicht gleich schwer: F-1
ist eine Formulierung. F-2 ist der Punkt, den die Maintainer-Abnahme sehen muss, denn ohne
Kennzeichnung nimmt sie einen Entscheid ab, der so nicht im Vertrag steht. F-3 bis F-6 betreffen
die Kernzusage „fail-safe“ und die Implementierbarkeit ohne eigene Vertragsentscheidung, also das
Ziel aus slice-209 §1.

**Konflikt-Pfad:** Ein Rollen-Widerspruch liegt noch nicht vor. Bestreitet der Implementer F-1
oder F-2, läuft der Konflikt nach `v6.13.0` · `regelwerk/modul-08-agentenrollen.md`
§Konflikt-Pfad als Rollen-Sequenz über den Architect. Eine Herabstufung auf Widerspruch hin
findet nicht statt.

**Übergabe:** Die Findings gehen an den Implementer. Bei F-2 läuft die Rückkante Review → Plan,
denn §1 des Slice-Plans und der Vertrag sagen Verschiedenes. Die Finding-Klassen gehen zusätzlich
in die Closure-Notiz §7 von slice-209 und von dort in den Zähler. Dieser Report ist Lauf-Beleg
und ersetzt keine Verifikation.

## Delta-Re-Review (Fix-Range f83a68c..919e3b3)

**Gegenstand:** `6149b53` (docs(spec) Nachlauf) und `919e3b3` (docs(planning)), geprüft gegen
F-1 bis F-14 dieses Reports. Dazu adversarisch: Öffnen die Fixes neue Lücken? **Gleicher Lauf,
gleicher Kontext, gleiches Modell** (`claude-opus-5-5`), Skill @ `a6d19b6`, Datum 2026-10-06.
Die Zeilenangaben beziehen sich auf den Stand `919e3b3`.

### Status je Finding

| ID | Status | Beleg |
|---|---|---|
| F-1 | behoben | `CHANGELOG.md:17-19` — „Vorgeschlagen, noch nicht abgenommen und nicht implementiert: die ADR steht auf `Proposed`“. |
| F-2 | behoben | `docs/plan/adr/0041-shapes-sollform-je-datei.md:79-84` — die Schärfung gegenüber Entscheid 5 ist ausgewiesen, der alte Wortlaut wird zitiert. In `docs/plan/planning/in-progress/slice-209-shapes-spec-first.md:72-84` steht ein eigener Block „mit der ADR zur Abnahme vorgelegt, nicht abgenommen“, die zehn Entscheide bleiben im Abnahme-Wortlaut. Die Abnahme bekommt die Abweichung damit zu sehen. |
| F-3 | behoben, mit neuer Restlücke N-1 | `spec/spezifikation.md:266-283` — Dollar-Präfix `n ≥ 2` für beide Zeichenketten-Formen; eine Vorlage öffnet erst ab `n` Zeichen `$` vor `{`, die letzten `n` gehören zur Vorlage. Der F-3-Fall ist durchgespielt (siehe unten). |
| F-4 | behoben als benannte Grenze — trägt, kein HIGH | `spec/spezifikation.md:94` (Grenze der Regex-Einträge, sichere Klasse `"[^"$\\]*"`), `spec/lastenheft.md:424`, ADR-0041 Punkt 5 und §Konsequenzen. Begründung unten. Die Verdrahtung der „sicheren Form“ in den Folge-Slice fehlt: N-3. |
| F-5 | behoben | `spec/spezifikation.md:94` — der Konfigurations-Adapter prüft nur Schlüssel, Werte und Regex. Die Zerlegbarkeit der `literal`-Einträge prüft der validierende Einstieg der Extraktion „vor jedem Dateizugriff“, Exit 2 auch in `--print-graph`. Die Gegenstücke stehen in `spec/spezifikation.md:312-315`, ARC-003 (`spec/architecture.md:73`) und ADR-0041 Punkt 9. Eine laterale Adapter-Kante ist ausgeschlossen. |
| F-6 | behoben, mit Rest N-2 | `spec/spezifikation.md:94` — ein Regex „muss für sich kompilieren“, danach wird er zu `^(?:<pattern>)$`; `a)\|(.*` ist als Beispiel genannt. |
| F-7 | behoben | `spec/spezifikation.md:94` — „der Dialekt wird erklärt, nicht erraten“, die Endung wird nicht geprüft. |
| F-8 | behoben | `spec/spezifikation.md:266-268` (UTF-8, BOM, `LF`/`CRLF`/`CR`, Leerraum-Menge), `spec/spezifikation.md:299-302` (erstes Zeichen, das weder Leerraum noch Kommentar ist; Zeichenkette, Zeichen-Literal und Backtick zählen dazu; Zeilenzahl ≥ 1). |
| F-9 | behoben | `spec/lastenheft.md:397-398` und `:416` — Form `pfad:zeile: <klasse>: <meldung>`, `shape-differs` nennt die Sollform-Anweisung. DoD von slice-209 ist mitgezogen. |
| F-10 | behoben | `spec/spezifikation.md:94` — byte-gleiche Befundzeilen werden einmal ausgegeben. Das deckt auch zwei Globs **eines** Eintrags, die dieselbe Datei treffen, denn die Ausgabe ist dann identisch. |
| F-11 | behoben | `spec/architecture.md:71` (ARC-001 ohne Zahl, drei Modelle), `spec/spezifikation.md:319` („Präzisiert die Regel-Anforderungen“). |
| F-12 | behoben | `spec/spezifikation.md:94` und ADR-0041 Punkt 8 — ein absoluter Pfad oder einer, der nach lexikalischer Normalisierung mit `..` beginnt, ist Exit 2 beim Laden. Symlinks im Baum, die hinauszeigen, behandelt die Regel nicht eigens — das ist dieselbe Frage wie beim bestehenden Walk, kein neuer Befund. |
| F-13 | behoben | `docs/plan/planning/open/slice-210-shapes-kotlin-allow-statements.md:106` — als bekannte Fehlalarm-Quelle und Testfall im Risiko. |
| F-14 | behoben | slice-210 DoD (`…/slice-210-shapes-kotlin-allow-statements.md:59-61`) schreibt den CHANGELOG-Eintrag ausdrücklich um. |

### Adversarisch: Dollar-Präfix-Regel (F-3-Nachspiel)

Gegen `spec/spezifikation.md:266-283` durchgespielt, Kotlin-Semantik der Multi-Dollar-Interpolation
als Vergleich:

- **F-3-Eingabe** `version = "x" + $$"${" + '"' + '}' + "// "; dependencies.add(…)`: Mit `n = 2`
  ist `${` (nur ein `$`) Inhalt, die Zeichenkette schließt am nächsten `"`. Lexer und Kotlin
  bleiben synchron, `dependencies.add(…)` wird eine eigene Anweisung und damit `shape-unlisted`.
  **Geschlossen.**
- **`$$$"…$${…"`** (`n = 3`, zwei `$`): Inhalt — wie Kotlin. **`$$$"…$$$${x}…"`**: vier `$`, das
  letzte Trio öffnet die Vorlage, das erste `$` ist Inhalt — wie Kotlin.
- **`$$"""…"""`** und **`$$$"""…$$${ "}" }…"""`**: Prefix gilt für die Roh-Form, die Vorlage endet
  über die Klammer-Tiefe, die innere `"}"` ist rekursiv gelext — synchron.
- **`$$"a$"`**, **`$$"$name"`**: `$` vor `"` bzw. vor einem Bezeichner ist Inhalt — synchron.
- **Präfix-Erkennung außerhalb von Zeichenketten** („eine Folge von `$` direkt vor `"` … ist immer
  ein Präfix“): `$` kommt in Kotlin-Code sonst nicht vor (Backtick-Bezeichner werden vorher als
  Token gelesen) — keine Fehlerkennung gefunden.
- **Restlücke N-1 — Backslash-Escape gegen die `$`-Folge.** In einer Zeichenkette mit Escape ist
  `\$` in Kotlin ein **wörtliches** `$`: `"\${"` ist der Text `${`, keine Vorlage. Der Vertrag
  formuliert die Vorlage über „eine Folge von mindestens `n` Zeichen `$` vor `{`“ und sagt nicht,
  ob ein escaptes `$` zu dieser Folge zählt. Wer die Folge rückwärts vom `{` aus zählt, öffnet bei
  `"\${"` eine Vorlage, die Kotlin nicht sieht — dieselbe Desynchronisation wie in F-3, mit
  derselben Folge: `version = "x" + "\${" + '"' + '}' + "// "; dependencies.add(…)` lässt den
  `add`-Aufruf als Kommentar verschwinden. Ebenso `$$"\$${x}"` (Kotlin: `$` wörtlich, dann `${`
  mit nur einem `$` → Inhalt). Wer zuerst das Escape verbraucht, liegt richtig — der Vertrag legt
  die Reihenfolge aber nicht fest, und das ist laut ADR-0041 §Konsequenzen „die eine Fehlrichtung,
  die nicht fail-safe ist“. In Roh-Zeichenketten gibt es kein Escape; `"""\${x}"""` ist in Kotlin
  wie im Vertrag eine Vorlage — synchron.

### Beurteilung F-4 (benannte Grenze statt Verbot)

**Trägt — kein HIGH.** Das ursprüngliche Finding war nicht „der Regex ist unsicher“, sondern
„die fail-safe-Zusage nennt ihre Grenze nicht“ (Klasse *fail-safe-Zusage ohne benannte Grenze*).
Die Grenze steht jetzt an allen drei Stellen, an denen der Zweck behauptet wird: Lastenheft
Out-of-Scope, SPEC-CONF-001 und ADR-0041 (Punkt 5 und §Konsequenzen). Ein Verbot wäre nur mit
token-bewusstem Regex durchsetzbar, und das liegt außerhalb der Text-Heuristik
(AC-QA-02). Die dokumentierte Klasse `"[^"$\\]*"` hält gegen die Prüffälle: Sie verlässt die
Zeichenkette nicht (`"` ausgeschlossen), trifft kein `\"` (Backslash ausgeschlossen), keine Vorlage
(`$` ausgeschlossen), keine Roh-Zeichenkette `"""…"""` (die Verankerung scheitert am Rest) und
kein Dollar-Präfix (`$$"…"` hat `$` vor `"`). Was bleibt: Die sichere Form erreicht den Adopter
nur über Handbuch und `--print-config`, und dort ist sie nicht verdrahtet (N-3).

### Adversarisch: literal-Normalisierung im validierenden Einstieg (F-5)

Ohne neue Lücke. Die Normalisierung läuft vor jedem Dateizugriff. Damit hat `--print-graph` die
Ladezeit-Parität aus SPEC-CLI-002, und die Scan-zeitigen Fälle (Glob ohne Treffer,
`exclude`-Widerspruch, nicht zerlegbare Datei) bleiben scanzeitig. Der Kern bekommt die
Einträge fertig und liest nichts. Eine neue Kante Konfigurations-Adapter → Extraktion entsteht
nicht, weil der Konfigurations-Adapter den Normalisierer laut ADR-0041 Punkt 9 nicht kennt.

### Neue Findings

| ID | Kategorie | Befund | Quelle | Pfad | Verifizierbar | Klasse |
|---|---|---|---|---|---|---|
| N-1 | MEDIUM | Ob ein escaptes `$` (`\$`) zur „Folge von mindestens `n` Zeichen `$` vor `{`“ zählt, legt die Lexik nicht fest. Wer rückwärts zählt, öffnet bei `"\${"` eine Vorlage, die Kotlin nicht sieht. Mit `version = "x" + "\${" + '"' + '}' + "// "; dependencies.add(…)` verschwindet der `add`-Aufruf dann als Kommentar — die nicht fail-safe Richtung, die ADR-0041 §Konsequenzen selbst benennt. | AC-FA-RULE-012 (fail-safe); ADR-0041 §Konsequenzen (Fehl-Lexik) | `spec/spezifikation.md:266-283` | ja — Testfall `"\${"` in `make test`, sobald slice-210 läuft | Lexik-Aufzählung behauptet Vollständigkeit, die die Zielsprache nicht trägt |
| N-2 | LOW | Ein Muster kann für sich kompilieren und umhüllt nicht: `\Qfoo` quotiert bis zum Ende, also scheitert `^(?:\Qfoo)$` an der offenen Gruppe. Der Vertrag nennt den Exit-Code für einen Kompilier-Fehler erst nach dem Umhüllen nicht. Die naheliegende Implementierung endet in Exit 2, aber der Fall ist unbenannt. | SPEC-CONF-001 (Exit-2-Fälle) | `spec/spezifikation.md:94` | ja — Testfall in `make test` | Verankerungs-Zusage ohne Kompilier-Reihenfolge |
| N-3 | LOW | Drei Folgepflichten aus ADR-0041 stehen nicht in der DoD von slice-210: „Benutzerhandbuch und `--print-config` zeigen nur die sichere Form“ (§Konsequenzen), „vollständige Lexik, belegt durch Tests je Form“ und das Dollar-Präfix. Die Lexik-Aufzählung dort nennt Roh-Strings und `${…}`, aber weder Dollar-Präfix noch Escape-gegen-`$`. | ADR-0041 §Konsequenzen; slice-209 §1 Ziel (Folge-Slices ohne eigene Vertragsentscheidung) | `docs/plan/planning/open/slice-210-shapes-kotlin-allow-statements.md:51-61` | nein | ADR-Folgepflicht ohne Träger im Folge-Slice |

### Negativbefunde des Deltas

| Bereich | Ergebnis |
|---|---|
| Spec-Straten-Richtung (`AGENTS.md` §3.4) im Delta | geprüft, ohne Befund — die hinzugefügten Zeilen unter `spec/` nennen kein ADR, keinen Slice, keine Welle. Die Historie-Zeile 0.33.0 ist fortgeschrieben, ohne Verweis nach unten. |
| ADR-0041 noch `Proposed` (§3.5 greift erst bei `Accepted`) | geprüft, ohne Befund — inhaltliche Änderung zulässig. |
| Lastenheft-Text im Delta geändert, Version bleibt 0.28.0 | geprüft, ohne Befund — dieselbe unveröffentlichte CR-Fassung im selben Slice, Status `Draft`. |
| Plan-Änderung slice-209 §1 (Kommentar-/Zustandsregel §3.7) | geprüft, ohne Befund — Zeitdokument, der Block nennt Zustand („vorgelegt, nicht abgenommen“), keine Chronik in einem Zustandsfeld. |
| `make doc-check` | Exit 0, „660 Datei(en) geprüft, 0 Befund(e)“ — Geltungsbereich: Links, Anker, Kennungs-Linkpflicht. |

### Summary des Deltas

| Kategorie | offen aus Erstlauf | neu |
|---|---|---|
| HIGH | 0 (F-1 behoben) | 0 |
| MEDIUM | 0 (F-2 bis F-6 behoben) | 1 (N-1) |
| LOW | 0 (F-7 bis F-12 behoben) | 2 (N-2, N-3) |
| INFO | 0 (F-13, F-14 behoben) | 0 |

**Finding-Klassen des Deltas:** Lexik-Aufzählung behauptet Vollständigkeit, die die Zielsprache
nicht trägt (zweites Auftreten in diesem Slice — **ein** Vorgang, zählt einmal) · Verankerungs-Zusage
ohne Kompilier-Reihenfolge (ebenso) · ADR-Folgepflicht ohne Träger im Folge-Slice

### Verdikt des Deltas

**Merge-blockierend:** ja, wegen N-1 (MEDIUM). Es ist derselbe Mechanismus wie F-3, eine Stufe
tiefer: Die Fehl-Lexik verschluckt Code als Kommentar, und das ist die einzige nicht fail-safe
Fehlrichtung des Vertrags. N-2 und N-3 blockieren nicht. Alle 14 Findings des Erstlaufs sind
behoben; F-4 ist als benannte Grenze tragfähig.
