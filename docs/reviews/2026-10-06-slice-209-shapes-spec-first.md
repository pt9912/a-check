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
