# Review-Report: slice-169 — 2026-09-08

**Review-Art:** Code — geprüft gegen den Slice-Plan, `AGENTS.md` §3/§4/§5 und
die Konfiguration in `.d-check.yml`. Kern-Behauptung des Gegenstands ist eine
**Deckungs-Zusage** („beide Ausfallarten gedeckt"); sie ist gegen das Werkzeug
gemessen worden, nicht gegen den Text gelesen.

**Gegenstand:** Commit `0565ca9` („feat(harness): slice-169 -- Korpus-Haelfte
der phrasen-basierten Kalibrierung") plus der **uncommittete** Arbeitsstand von
`docs/plan/planning/in-progress/slice-169-korpus-seitige-kalibrierung.md`
(dort ist die Review-DoD-Zeile auf `- [ ]` zurückgesetzt).

**Skill:** `.harness/skills/reviewer.md` @ Stand `0565ca9` ·
<!-- d-check:ignore -->
**Modell:** claude-opus-5[1m] · **Datum:** 2026-09-08

**Review-Art nach Rollen-Trennung:** unabhängiger Lauf — eigenes Kontextfenster,
der Gegenstand wurde in dieser Sitzung nicht verfasst
(`v6.5.0` · `regelwerk/modul-08-agentenrollen.md` §Rollen-Regeln).

**Eingangs-Kontext:**

- `.harness/skills/reviewer.md` (Klassifikation, Output-Schema, Zitier-Form)
- `AGENTS.md` §3 (Hard Rules), §4 (Gate-Tabelle), §5 (Dokumentations-Regeln)
- `harness/conventions.md` §Modus-Deklaration pro Sub-Area
- Slice-Plan `slice-169`, Arbeitsstand
- `.d-check.yml` (Module `reviews`, `structure`, `versions`)
- Beobachtungs-Eintrag `BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf`
  (`observation.md`, `state.md`, fünf `evidence/`-Dateien)
- `v6.5.0` · `regelwerk/modul-13-quality-gates.md` §Vorhanden ≠ behauptet
- `v6.5.0` · `regelwerk/modul-05-planning-harness.md` §Ziel-Form: Slice
- frühere Findings am selben Bereich: Review zu `slice-168` (F-1/F-2, im
  Slice-Archiv `done/wellenlos/slice-168-archiv.zip`), Review zu `slice-179`

**Geltungsbereich dieses Laufs** (`AGENTS.md` §5, *Geltungsbereich einer
Messung*): Geprüft sind (a) das Verhalten von `make dcheck-phrase-selftest`
gegen **mutierte Kopien** des `done/`-Bestands, (b) der `sed`-Auszug gegen
sieben YAML-Schreibweisen, (c) die Zahlen-Behauptungen gegen den Bestand,
(d) die Form des Slice-Plans gegen `AGENTS.md` §5. **Nicht** geprüft: die
Innereien des gepinnten `d-check`-Images jenseits seines beobachtbaren
Verhaltens, und die vier Werkzeug-Kontrollen über ihren unveränderten Stand
hinaus. Alle Proben liefen auf Kopien; der Gegenstand wurde nicht angefasst.

---

## Findings

### F-1 — Beide Korpus-Kontrollen melden grün, während die Kandidatenmenge des Moduls leer ist

- `kategorie`: HIGH
- `quelle`: Harness-Lüge (behauptete Deckung ohne Deckung) ·
  `BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf` — genau die Ausfallart, die
  der Sensor zu decken erklärt · `v6.5.0` ·
  `regelwerk/modul-13-quality-gates.md` §Vorhanden ≠ behauptet
- `pfad`: `tools/dcheck-phrase-selftest.sh:171` (reviews) und
  `tools/dcheck-phrase-selftest.sh:182–183` (tasks)
- `befund`: Beide Kontrollen zählen eine **Obermenge** der Menge, über die sie
  eine Aussage machen: `grep -rli` zählt jede Datei unter `done/`, die die
  Phrase irgendwo trägt (rekursiv, ganze Datei, Nicht-Slices eingeschlossen),
  während das Modul `reviews` nur einen **DoD-Haken** in einem flachen
  `done/`-Slice sieht; `grep -rhoE '^- \[[ x]\] .*'` zählt DoD-förmige Zeilen
  in jedem Pfad unter `done/`, während `structure` (1) auf
  `docs/plan/planning/**/slice-*.md` abzüglich `exempt-paths` und nur innerhalb
  der DoD-Sektion keilt. Gemessen auf einer Kopie: Nach Entwerten der beiden
  einzigen echten `reviews`-Kandidaten (`slice-166`, `slice-167`) meldet der
  Lauf **Exit 0** mit „3 Slice(s)"; nach zusätzlicher Verlagerung des einzigen
  `tasks`-Treffers in einen `exempt-paths`-Pfad (`done/wellenlos/`) meldet er
  **Exit 0** mit „1 Slice(s) · 1 DoD-Zeile(n)", obwohl beide realen
  Kandidatenmengen leer sind.
- `verifizierbar`: ja — reproduzierbar über eine Kopie von `.d-check.yml`,
  `tools/` und `docs/plan/planning/done/`; **kein** bestehendes Gate fängt es,
  der Selbsttest ist selbst der Prüfer.
- `klasse`: Sensor misst eine Obermenge seines erklärten Gegenstands

### F-2 — `state.md` schaltet den Register-Eintrag auf *verkörpert* mit einer Deckung, die die Messung nicht trägt

- `kategorie`: HIGH
- `quelle`: `v6.5.0` · `regelwerk/modul-06-roadmap.md` §Das
  Beobachtungs-Register (Ausgang *verkörpert*: „die Regel steht") ·
  `AGENTS.md` §5, Beobachtungs-Register
- `pfad`: `docs/plan/planning/observations/BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf/state.md:5–8`
  und `docs/plan/planning/in-progress/slice-169-korpus-seitige-kalibrierung.md:112–113`
- `befund`: `state.md` sagt „Beide Ausfallarten sind damit gedeckt: … dass
  a-checks eigener Bestand sie nicht mehr trägt (Nichtleerheit gegen `done/`,
  zwei Kontrollen)"; F-1 zeigt gemessen, dass der Lauf grün bleibt, wenn der
  Bestand sie nicht mehr trägt. Der DoD-Punkt „Korpus-seitige Kontrolle **der
  Kandidatenmenge** umgesetzt und gegen eine absichtlich leere Menge
  verifiziert" ist abgehakt; verifiziert wurde die Leerheit der vom Skript
  gezählten Menge (per Muster-Mutation), nicht die der Kandidatenmenge des
  Moduls.
- `verifizierbar`: ja — dieselben Proben wie F-1; `make verify-observations`
  prüft nur die Deckung Pfad ↔ Verzeichnis, nicht die Aussage im `state.md`.
- `klasse`: Register-Ausgang *verkörpert* ohne tragende Verkörperung

### F-3 — Die `tasks`-Korpus-Kontrolle liest Roh-Text, das Modul liest bereinigten Text

- `kategorie`: MEDIUM
- `quelle`: `.d-check.yml`, Kommentar zu `structure` (1): „Das Muster ist
  VERANKERT und sieht den BEREINIGTEN Item-Text: der Gate-Lauf steht als
  `make gates` in Inline-Code, und der ist beim Zaehlen zu Leerzeichen
  bereinigt"
- `pfad`: `tools/dcheck-phrase-selftest.sh:182–183`
- `befund`: Das Skript reicht dem Muster den unbereinigten Item-Text
  (`sed 's/^- \[[ x]\] //'`). Gemessen über `done/`: von den sechs
  Alternativen des Musters feuern nur drei — `Jedes Risiko` (7×),
  `Beobachtungs-Register` (5×), `Unabhängiger Review` (2×); `^ *grün`,
  `Closure-Notiz` und `Reconciliation` feuern **0×**, weil die realen Zeilen
  mit einem Backtick beginnen (`` - [x] `make gates` grün.``). Die
  Nichtleerheit ruht damit auf drei Alternativen; ein Bruch der Alternative,
  um die das Muster gebaut wurde, bleibt für diese Kontrolle unsichtbar.
- `verifizierbar`: ja — `grep -oE '^( *grün|…)'` über
  `docs/plan/planning/done/` reproduziert die Verteilung 7/5/2/0/0/0.
- `klasse`: Selbsttest normalisiert anders als der geprüfte Prüfer

### F-4 — „5 Slice(s)" nennt eine Wellen-Ergebnisnotiz einen Slice

- `kategorie`: MEDIUM
- `quelle`: nachweislich falsche Tatsachenbehauptung, gegen den `done/`-Bestand
  verifiziert
- `pfad`: `tools/dcheck-phrase-selftest.sh:201`;
  `docs/plan/planning/in-progress/slice-169-korpus-seitige-kalibrierung.md:93`
  („heute 5 Slices und 14 DoD-Zeilen")
- `befund`: Die fünf gezählten Dateien sind `slice-164`, `slice-165`,
  `slice-166`, `slice-167` und `welle-15-results.md`; die letzte ist kein
  Slice. Von den vier Slices tragen nur zwei (`slice-166`, `slice-167`) die
  Phrase auf einem DoD-Haken, die anderen zwei nennen sie in Prosa bzw. in
  einer Tabelle. Der Plan begründet die gezählte Zahl damit, sie sage „beim
  Lesen, worüber das Grün eine Aussage macht" — die Zahl 5 sagt es nicht.
- `verifizierbar`: ja — `grep -rli` gegen `grep -lE '^- \[[ x]\] .*[Uu]nabhängiger Review'`
  über `docs/plan/planning/done/`.
- `klasse`: Zählgröße anders benannt als gezählt

### F-5 — Der `reviews`-Fehlerpfad ist unerreichbar: Abbruch vor der eigenen Diagnose

- `kategorie`: MEDIUM
- `quelle`: Maintainability · das Skript deklariert die Meldung, die es nie
  ausgibt
- `pfad`: `tools/dcheck-phrase-selftest.sh:39` (`set -euo pipefail`) in
  Verbindung mit `:171–178`
- `befund`: Bei leerer Treffermenge endet `grep -rli … | wc -l` wegen
  `pipefail` mit Status 1; die Zuweisung schlägt fehl und `set -e` beendet das
  Skript **vor** den vier `echo`-Zeilen und vor der Sammelmeldung. Gemessen auf
  einer Kopie, aus der die Phrase restlos entfernt wurde: **Exit 1 ohne jede
  Ausgabe** — weder „Kandidatenmenge des reviews-Moduls ist LEER" noch
  „mindestens eine Kalibrierungs-Kontrolle ist rot". Das Gate wird rot, die
  Diagnose fehlt; die Sensor-Datei führt unter §Sperren nur die
  Feld-nicht-lesbar-Meldung.
- `verifizierbar`: ja — Kopie mit ersetzter Phrase, Skript direkt aufrufen.
- `klasse`: Fehlerpfad durch `set -e`/`pipefail` unerreichbar

### F-6 — Die zweite Korpus-Kontrolle deckt eine Klasse, die **laut** ausfällt

- `kategorie`: MEDIUM
- `quelle`: das eigene Ausschluss-Argument des Slice gegen
  `versions.current-from` („fällt aber **laut** aus … Ihm fehlt die
  Gefährlichkeit der leeren Prüfmenge")
- `pfad`: `docs/plan/planning/in-progress/slice-169-korpus-seitige-kalibrierung.md:72–83`;
  `harness/sensors/dcheck-phrase-selftest.md:34–40`;
  `tools/dcheck-phrase-selftest.sh:185–188`
- `befund`: `tasks-ignore-pattern` ist ein **Ignore**-Muster in
  `structure` (1) mit `max-tasks: 3`. Trifft es nichts mehr, zählen die
  konstanten Posten mit; die flachen `done/`-Slices tragen 6–7 DoD-Posten, also
  meldet `doc-structure` `section-oversized` — dieselbe Reaktion, die die
  Negativ-Kontrolle des Skripts (`structure-neg`) bereits belegt. Der Ausfall
  ist damit laut, nicht still. Die Meldung des Skripts behauptet dennoch beide
  Richtungen („meldet entweder falsch rot **oder gar nicht mehr, was es
  soll**"). Mit dem Maßstab, mit dem `versions.current-from` ausgeschlossen
  wird, gehört diese Kontrolle nicht in die Klasse, die der
  Beobachtungs-Eintrag beschreibt.
- `verifizierbar`: ja — die vorhandene `structure-neg`-Fixture des Skripts ist
  der Beleg; zusätzlich die gemessenen DoD-Postenzahlen (6–7 je Slice).
- `klasse`: Sensor deckt eine andere Ausfallklasse als die deklarierte

### F-7 — „vierzehn phrasen-basierte Felder in `.d-check.yml`" ist nicht reproduzierbar

- `kategorie`: MEDIUM
- `quelle`: `AGENTS.md` §5, *Geltungsbereich einer Messung* (seit `slice-179`)
  — eine Messung als Beleg nennt ihren Geltungsbereich
- `pfad`: `tools/dcheck-phrase-selftest.sh:36` und `:202`;
  `harness/sensors/dcheck-phrase-selftest.md:34`;
  `docs/plan/planning/observations/BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf/state.md:16–19`;
  `docs/plan/planning/in-progress/slice-169-korpus-seitige-kalibrierung.md:72–78`
- `befund`: Die Zahl steht in vier Artefakten (plus Commit-Message) als
  gemessene Tatsache, ohne die Zählregel zu nennen. `.d-check.yml` führt je
  nach Zählweise deutlich mehr regex-/phrasen-tragende Einträge (allein
  `structure` fünf `section-pattern`, drei `exempt-section-pattern`,
  `tasks-ignore-pattern`, `forbid-pattern`, `require-pattern`, zwei
  `require-all`; dazu `versions` zwei `pin-pattern`, `vcs` drei, `planning`
  zwei, `commits` zwei, `ids` fünf, `matrix` einen). Hinzu kommt ein
  Selbstwiderspruch: eines der beiden „gedeckten" Muster — die
  `reviews`-Trigger-Phrase — ist **kein Feld in `.d-check.yml`**; sie steht
  dort nur in einem Kommentar und lebt im gepinnten Werkzeug. Die Aussage
  „`.d-check.yml` führt vierzehn phrasen-basierte Felder; gedeckt sind die
  zwei" trifft damit auf das eine der zwei nicht zu.
- `verifizierbar`: nein — es gibt keinen Lauf, der die Zählregel prüft; die
  Nicht-Reproduzierbarkeit ist der Befund.
- `klasse`: Zahl als Beleg ohne genannten Geltungsbereich

### F-8 — Slice-Plan §1 trägt keine Abgrenzung; §9 führt den alten Abschnitts-Titel

- `kategorie`: MEDIUM
- `quelle`: `AGENTS.md` §5, Kopieranleitung Punkte 5 und 6 · `v6.5.0` ·
  `regelwerk/modul-05-planning-harness.md` §Ziel-Form: Slice („§1 nennt Ziel
  *und* Abgrenzung", je Punkt mit Begründung) und §Ziel-Form:
  Sub-Area-Modus-Begründung
- `pfad`: `docs/plan/planning/in-progress/slice-169-korpus-seitige-kalibrierung.md:25`
  (§1 „Ziel") und `:201` (§9 „Sub-Area-Modus")
- `befund`: §1 nennt vier Zeilen Ziel und keinen einzigen Ausschluss; §9 heißt
  „Sub-Area-Modus" statt „Sub-Area-Prüfungen und Modus-Begründung". Beide
  Regeln standen zum Commit-Zeitpunkt bereits in `AGENTS.md` (belegt:
  `git show 0565ca9^:AGENTS.md` führt sie in §5 Punkt 5/6), und der Commit hat
  152 Zeilen dieses Plans umgeschrieben, ohne sie nachzuziehen. Folge für
  diesen Review: Die Frage „hat die Umsetzung etwas mitgenommen, das §1
  ausschließt?" ist am Plan **nicht beantwortbar** — es gibt keine Grenze, an
  der sich das Wachstum messen ließe.
- `verifizierbar`: nein — `AGENTS.md` §5 stellt ausdrücklich fest, dass auf den
  Abgrenzungs-Abschnitt kein Sensor läuft.
- `klasse`: Slice-Form der Ziel-Form nicht nachgezogen

### F-9 — Der Sichtungs-Block in §9 ist im `in-progress`-Stand nicht ausgefüllt

- `kategorie`: MEDIUM
- `quelle`: `AGENTS.md` §5, Kopieranleitung Punkt 6 („Die zwei
  *Vorgelagert*-Blöcke … laufen in **jedem** Slice-Plan; der Abschnitt entfällt
  nie") · `v6.5.0` · `regelwerk/modul-05-planning-harness.md` §Zwei Schritte vor
  der Modus-Begründung („Keine Treffer sind ebenfalls eine Antwort und werden
  notiert")
- `pfad`: `docs/plan/planning/in-progress/slice-169-korpus-seitige-kalibrierung.md:207–209`
- `befund`: Der Block sagt „entsteht mit dem Übergang nach `in-progress/`". Der
  Übergang ist mit Commit `3f118e5` erfolgt, der Plan wurde danach mit
  `0565ca9` umgeschrieben, und der Block steht unverändert als Ankündigung da.
  Ein Sichtungs-Ergebnis — auch das leere — ist nicht notiert; im Repo
  ohne Wellen-Betrieb ist dieser Schritt der einzige Leser für alles unter der
  Schwelle.
- `verifizierbar`: nein — kein Sensor prüft den Inhalt der beiden
  Vorgelagert-Blöcke.
- `klasse`: Vorgelagerter Pflicht-Block als Ankündigung stehen geblieben

### F-10 — `versions.current-from` wird dem falschen Slice zugeschrieben

- `kategorie`: MEDIUM
- `quelle`: `evidence/slice-173.md` im selben Beobachtungs-Eintrag („Die
  Korpus-Ausprägung hat ein **drittes** Muster bekommen, und es entstand mit
  **diesem** Slice") · `state.md` desselben Commits nennt `evidence/slice-173.md`
- `pfad`: `harness/sensors/dcheck-phrase-selftest.md:37`
- `befund`: Die Sensor-Datei schreibt „`versions.current-from` wäre der nächste
  Kandidat (slice-179 belegt ihn als drittes Muster)". Der Beleg ist
  `evidence/slice-173.md`; der Volltext von `slice-179` (Archiv
  `done/welle-15/archiv.zip`) enthält keine Aussage über `current-from` als
  drittes Muster. Zwei Artefakte desselben Commits nennen unterschiedliche
  Slices für dieselbe Herkunft.
- `verifizierbar`: ja — `unzip -p` auf das Wellen-Archiv plus Lesen der
  `evidence/`-Datei; kein Gate prüft die Zuordnung.
- `klasse`: Herkunfts-Zuschreibung auf den falschen Vorgang

### F-11 — Sensor-Datei und Plan tragen `slice-169` nicht als Bindung bzw. Liefer-Artefakt

- `kategorie`: LOW
- `quelle`: Provenance-Disziplin der Sensor-Dateien (Muster der übrigen Zeilen
  in `AGENTS.md` §4 und `harness/README.md` §Sensors, die mehrere Slices je
  Ausbaustufe nennen)
- `pfad`: `harness/sensors/dcheck-phrase-selftest.md:48–52` (Bindung:
  „slice-168"); `harness/README.md` Zeile zu `make dcheck-phrase-selftest`
  (Bindung endet auf „slice-168");
  `docs/plan/planning/in-progress/slice-169-korpus-seitige-kalibrierung.md:87–90`
  (§3-Tabelle ohne `harness/sensors/dcheck-phrase-selftest.md`)
- `befund`: Der Commit hat 48 Zeilen der Sensor-Datei umgeschrieben; die
  §3-Umsetzungstabelle des Plans führt nur `tools/…`, `AGENTS.md` und
  `harness/README.md`. Die Bindungs-Angaben beider Doku-Tabellen nennen
  weiterhin allein `slice-168`, obwohl die Korpus-Hälfte aus `slice-169`
  stammt.
- `verifizierbar`: nein — `make doc-targets` prüft Existenz und
  Deklarations-Parität, nicht die Vollständigkeit der Bindungs-Spalte.
- `klasse`: Ausbaustufe ohne Provenance-Nachtrag

### F-12 — `head -1` beim Feld-Auszug wählt bei Namensgleichheit still das erste Feld

- `kategorie`: LOW
- `quelle`: Maintainability · `.d-check.yml` führt `structure` als
  **Regel-Liste**, ein zweites `tasks-ignore-pattern` ist bauartbedingt möglich
- `pfad`: `tools/dcheck-phrase-selftest.sh:160`
- `befund`: Der Auszug nimmt den ersten Treffer im Dokument. Heute existiert
  genau ein `tasks-ignore-pattern`; kommt eine zweite `structure`-Regel mit
  demselben Feld hinzu, prüft der Selbsttest weiter gegen die erste, ohne dass
  eine Meldung darauf hinweist. Der fail-closed-Zweig greift nur bei **null**
  Treffern, nicht bei mehreren.
- `verifizierbar`: ja — zwei gleichnamige Felder in eine Kopie von
  `.d-check.yml` setzen und den Auszug beobachten.
- `klasse`: Mehrdeutigkeit ohne Sperre

### F-13 — `Verantwortlich:` steht auf „noch nicht priorisiert", während der Slice in `in-progress/` liegt

- `kategorie`: LOW
- `quelle`: `v6.5.0` · `regelwerk/modul-05-planning-harness.md` §Lifecycle als
  State Machine („`open → next` setzt den Verantwortlichen … bis zur
  Priorisierung steht dort `—`")
- `pfad`: `docs/plan/planning/in-progress/slice-169-korpus-seitige-kalibrierung.md:18`
- `befund`: Das Kopf-Feld lautet `— *(noch nicht priorisiert)*`, obwohl der
  Slice seit `3f118e5` in `in-progress/` liegt und die Arbeit abgeschlossen
  ist. Die Kopffeld-Regel von `make doc-structure` prüft die **Anwesenheit**
  des Feldes, nicht seine Konsistenz zum Verzeichnis.
- `verifizierbar`: nein — kein Sensor deckt die Konsistenz Feld ↔ Verzeichnis.
- `klasse`: Kopffeld nicht mit dem Lifecycle-Übergang mitgezogen

### F-14 — Kommentar begründet im Konjunktiv über die verworfene Alternative

- `kategorie`: LOW
- `quelle`: `AGENTS.md` §3.7 („Falsch: Konjunktiv über die verworfene
  Alternative … Richtig: Indikativ über den Zustand") · `v6.5.0` ·
  `regelwerk/grundlagen-harness-dateien.md` §Was ein Kommentar trägt
- `pfad`: `tools/dcheck-phrase-selftest.sh:152–155`
- `befund`: Der Block nennt zuerst die geltende Zusage („Gefragt ist
  NICHTLEERHEIT, keine Erwartungszahl") und begründet sie dann im Konjunktiv
  über die nicht gewählte Variante („eine feste Zahl braeche zusaetzlich bei
  jedem neuen Slice"). Der Bestand führt dieselbe Form bereits an
  `:102–105` aus `slice-168`; der Befund ist damit ein Muster, kein Einzelfall.
- `verifizierbar`: nein — `AGENTS.md` §3.7 weist die Regel ausdrücklich als
  inferentiell und sensorlos aus.
- `klasse`: Abwägung im Kommentar statt in der Entscheidung

## Negativbefunde

- geprüft, ohne Befund: **`sed`-Auszug gegen YAML-Schreibweisen.** Sieben
  Varianten auf Kopien probiert — einfache Quotes (Ist-Form) und doppelte
  Quotes werden korrekt gelesen; unquotiert, mit Trailing-Kommentar, mit
  Tab-Einrückung und als Block-Skalar liefern jeweils den **leeren** Wert und
  damit den fail-closed-Abbruch. Eine auskommentierte Zeile vor der echten wird
  übersprungen. Die im Auftrag vermutete stille Leerwert-Falle (leerer Wert ⇒
  Test grün) existiert **nicht**: der leere Wert führt zu `return 1` und
  `fail=1`.
- geprüft, ohne Befund: **`AGENTS.md` §3.1 (Docker/make-only).** Das Skript ruft
  ausschließlich `docker`, `git`, `mktemp`, `chmod`, `sed`, `grep`, `wc`,
  `printf`, `echo` auf; keine Host-Toolchain, kein Paketmanager, kein
  YAML-Werkzeug. Der `d-check`-Lauf geht über den digest-gepinnten Digest mit
  `--network none`.
- geprüft, ohne Befund: **`AGENTS.md` §3.2 (Suppression-Verbot).** Keine
  Suppression-Direktive im geänderten Skript.
- geprüft, ohne Befund: **`AGENTS.md` §3.3 (git mv + Inhalt = zwei Commits).**
  Der Lifecycle-Wechsel liegt als eigener Commit `3f118e5` vor, die
  Inhaltsänderung als `0565ca9`.
- geprüft, ohne Befund: **`AGENTS.md` §3.5/§3.6.** Keine ADR berührt, keine
  Gate-Schwelle gesenkt.
- geprüft, ohne Befund: **Werkzeug-Hälfte (vier Kontrollen).** Unverändert
  gegenüber `slice-168`; Lauf reproduziert. Zusätzlich mit einer eigenen
  Fixture bestätigt, dass das Modul `reviews` nur den DoD-Haken sieht und eine
  reine Prosa-Erwähnung **nicht** auslöst — diese Fixture ist zugleich der
  Beleg für F-1.
- geprüft, ohne Befund: **Zahl „14 DoD-Zeile(n)".** Sie stimmt heute mit der
  Zahl in den flachen `done/slice-*.md` überein (14 = 14); die Divergenz ist
  latent (F-1) bzw. betrifft die Alternativen-Verteilung (F-3), nicht die
  heutige Summe.
- geprüft, ohne Befund: **Plan §2, Messungen zu `d-check`.** `--json` liefert
  in `summary` tatsächlich nur `filesChecked` und `findingCount`;
  `--print-config` gibt tatsächlich das Startgerüst aus. Beide nachvollzogen.
- geprüft, ohne Befund: **Größen-Regel des DoD.** Drei Liefer-Punkte neben vier
  konstanten Posten — innerhalb der Regel „höchstens drei".
- geprüft, ohne Befund: **Risiko-Ausgänge (§7).** Drei Risiken, jedes mit genau
  einem Ausgang aus der geschlossenen Dreier-Menge (2× *gestrichen mit
  Begründung*, 1× *weiter offen → Register*).
- geprüft, ohne Befund: **Closure-Notiz (§8).** Lerneintrag-Form benannt
  („neuer Sensor"), Ursache-Satz vorhanden, drei Paarungen adressiert.
- geprüft, ohne Befund: **Gate-Läufe auf dem Arbeitsstand.**
  `make dcheck-phrase-selftest` Exit 0, `make verify` Exit 0 (21 Anforderungen,
  0 Waisen).
- geprüft, ohne Befund: **Deklarations-Parität Doku ↔ Makefile.** Das Target
  `dcheck-phrase-selftest` existiert, hängt im `gates`-Aggregat und steht in
  beiden Doku-Tabellen; `make doc-targets` ist Teil des grünen `gates`-Laufs.
- geprüft, ohne Befund: **Traceability/Commit-Scope.** `feat(harness)` mit
  Kennung `slice-169`; der Scope `(planning)` ist nicht betroffen.
- geprüft, ohne Befund: **Evidence-Zählstand.** Zum Commit-Zeitpunkt lagen vier
  `evidence/`-Dateien vor; die im Plan genannte Zahl „4×" war korrekt. Die
  fünfte (`slice-181.md`) kam erst mit dem Folge-Commit.

## Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 2 |
| MEDIUM | 8 |
| LOW | 4 |
| INFO | 0 |

**Finding-Klassen dieses Laufs:** Sensor misst eine Obermenge seines erklärten
Gegenstands · Register-Ausgang *verkörpert* ohne tragende Verkörperung ·
Selbsttest normalisiert anders als der geprüfte Prüfer · Zählgröße anders
benannt als gezählt · Fehlerpfad durch `set -e`/`pipefail` unerreichbar · Sensor
deckt eine andere Ausfallklasse als die deklarierte · Zahl als Beleg ohne
genannten Geltungsbereich · Slice-Form der Ziel-Form nicht nachgezogen ·
Vorgelagerter Pflicht-Block als Ankündigung stehen geblieben ·
Herkunfts-Zuschreibung auf den falschen Vorgang · Ausbaustufe ohne
Provenance-Nachtrag · Mehrdeutigkeit ohne Sperre · Kopffeld nicht mit dem
Lifecycle-Übergang mitgezogen · Abwägung im Kommentar statt in der Entscheidung

## Verdikt

**Merge-blockierend: ja.** Zwei HIGH und acht MEDIUM. Der tragende Grund ist
F-1/F-2: Der Slice fügt eine Kontrolle hinzu, die grün meldet, während die
Menge, über die sie eine Aussage macht, leer ist — auf einer Kopie zweimal
gemessen (Exit 0 bei entwerteten `reviews`-Kandidaten; Exit 0 bei beiden leeren
Kandidatenmengen). Das ist dieselbe Klasse, gegen die der Slice antritt, und
`state.md` trägt sie bereits als *verkörpert* ein. Die Werkzeug-Hälfte aus
`slice-168` und der fail-closed-`sed`-Auszug sind davon nicht betroffen; die
Kopplung an `.d-check.yml` (Befund F-2 des `slice-168`-Reviews) ist belastbar
umgesetzt.

**Übergabe:** Findings gehen an den Implementer; die Finding-Klassen gehen
zusätzlich in die Closure-Notiz und von dort in den Zähler. Dieser Report ist
ein **Lauf-Beleg** und ersetzt keine Verifikation — DoD-/Spec-Konformität prüft
der Verifier separat.
