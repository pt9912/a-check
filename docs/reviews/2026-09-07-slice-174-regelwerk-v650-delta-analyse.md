# Review-Report: slice-174 — 2026-09-07

**Review-Art:** Plan — geprüft gegen den Slice-Plan, `AGENTS.md` §3/§4/§5, die
Ziel-Formen der adoptierten Baseline (`v6.2.0` · `regelwerk/modul-05-planning-harness.md`
§Ziel-Form: Slice, `regelwerk/modul-06-roadmap.md` §Roadmap-Regeln und
§Das Beobachtungs-Register) und gegen den **nachgerechneten** Upstream-Diff. Der
Slice ist eine reine Messung; geprüft wurde deshalb vorrangig, ob die Zahlen
reproduzieren und ob die Einordnung trägt.

**Gegenstand:** Commit-Range `61b7661^..HEAD` — `61b7661` (Welle-Eröffnung),
`0518e90` (Lifecycle-`git mv`), `dc0e58b` (Delta-Analyse plus slice-175/176/177
und ein Register-Eintrag).

**Skill:** `.harness/skills/reviewer.md` @ Stand `HEAD` (unverändert seit
slice-159) · <!-- d-check:ignore -->
**Modell:** claude-opus-5[1m] · **Datum:** 2026-09-07

**Review-Unabhängigkeit:** unabhängiger Lauf — eigenes Kontextfenster, die
Analyse wurde nicht von dieser Instanz verfasst. Alle Upstream-Zahlen sind in
einem frischen Klon von `pt9912/ai-harness-course` (Tags `v6.2.0` und `v6.5.0`)
selbst gemessen, nicht dem Slice-Zitat entnommen.

**Eingangs-Kontext:**

- `docs/plan/planning/done/slice-174-regelwerk-v650-delta-analyse.md`
- `docs/plan/planning/welle-15-regelwerk-v650-migration.md`
- `docs/plan/planning/open/slice-175-…`, `…/slice-176-…`, `…/slice-177-…`
- `docs/plan/planning/observations/BEO-GATE/versions-sensor-trifft-planungs-vorgriff/`
- `AGENTS.md` §3 (Hard Rules), §4 (Quality Gates), §5 (Dokumentations-Regeln), §6
- `harness/conventions.md` §Baseline und §Adaptions-Block
- `docs/plan/planning/done/wellenlos/slice-172-…` und `…/slice-173-…`
  (Volltexte aus den Archiven gelesen)
- `.d-check.yml` (Module `versions`, `links`, `ids`, `targets`, `reviews`)
- Kurs-Repo, Tags `v6.2.0` / `v6.5.0`, frischer Klon
- adoptierte Baseline: `v6.2.0` · `regelwerk/modul-05-planning-harness.md`,
  `regelwerk/modul-06-roadmap.md`, `regelwerk/modul-13-quality-gates.md`

**Eigene Läufe (kein Repo-Zustand verändert):** `make doc-check` Exit 0 ·
`make verify` Exit 0 · `make doc-targets` Exit 0 ·
`make commit-scope-check RANGE=61b7661^..HEAD` Exit 0 · `make trace-check`
Exit 2 (siehe F-16) · `git status --porcelain` vor und nach dem Review leer.

---

## Findings

### F-1 — Die Bereinigung erfasst nur eine von zwei Formatierungs-Klassen: 84 der 536 „substanziellen" Zeilen sind reine Whitespace-Umformatierung

- `kategorie`: HIGH
- `quelle`: eigene Zusage des Slice (§3.1 „Gemessen je Datei", §8
  Steering-Loop-Eintrag „die Bereinigung ist zu messen, nicht zu schätzen");
  Nachrechnung im Kurs-Klon
- `pfad`: `slice-174-…md` §2 (Posten-Tabelle), §3.1 (Klassen-Tabelle), §3.3, §8
- `befund`: Bereinigt wird ausschließlich die Klasse *Tabellen-Trennzeile*
  (`| --- | --- |`, 36 Zeilen, von mir mit derselben Regex reproduziert). Der
  Kurs hat in denselben Dateien zusätzlich die **Zell-Innenabstände** ganzer
  Tabellen normalisiert; `git diff --numstat -w v6.2.0 v6.5.0 -- lab/regelwerk
  lab/templates` liefert **16 Dateien und +452/−29** statt der im Slice
  geführten 21 Dateien mit 536 substanziellen Zeilen — 84 weitere `+`-Zeilen
  sind inhaltsgleich. Konkret: `grundlagen-begriffe.md` steht in §2 und in der
  Substanz-Klasse (§3.1) mit **42** Zeilen und trägt **2** (zwei neue
  Glossar-Zeilen, `RTM` und `harness/sensors/<target>.md`; der Spalten-Diff der
  Begriffs-Spalte zeigt +2/−0, keine Umbenennung). `grundlagen-bootstrap.md`
  (+18), `modul-02` (+1), `modul-04` (+6), `grundlagen-source-precedence` (+3)
  und `grundlagen-referenz-richtung` (+6) haben **null** Inhaltsänderung —
  `git diff -w` liefert für alle fünf eine leere Ausgabe. §3.3 führt genau diese
  fünf als „Tabellen-Ausbau und Verweise auf die neuen Abschnitte"; ausgebaut
  wurde nichts.
- `verifizierbar`: ja — `git diff --numstat -w v6.2.0 v6.5.0 -- lab/regelwerk
  lab/templates` im Kurs-Klon, gegen die Zahlen in §2/§3.1 gehalten. Kein
  Repo-Gate deckt es (Upstream-Messung).
- `klasse`: Bereinigung deckt nur die gesuchte Formatierungs-Klasse

### F-2 — T-1 setzt zwei verschiedene Klemmen gleich und schreibt a-check eine Behandlung zu, die `harness/conventions.md` ausdrücklich verneint

- `kategorie`: HIGH
- `quelle`: `harness/conventions.md` §Baseline; `slice-173` §2.1 (Volltext aus
  `done/wellenlos/slice-173-archiv.zip`)
- `pfad`: `slice-174-…md` §3.2, T-1, Absatz „Das ist genau die Klemme aus
  slice-172/173"
- `befund`: Der Slice schreibt: *„a-check hat sie mit `exempt-paths` behandelt
  … In einem einfrierenden Artefakt entsteht der Link gar nicht erst, dann
  braucht es auch keine Ausnahme."* `slice-173` §2.1 trennt genau diese beiden
  Dinge und hält fest, dass `slice-172` sie zusammengezogen hatte: *„`versions`
  hält lebende Dokumente auf dem adoptierten Stand … Die Link-Prüfung sorgt
  dafür, dass jeder Pfad auflöst, auch in Zeitdokumenten … slice-172 hat die
  Ursache falsch benannt."* `exempt-paths` steht im Modul `versions` und wirkt
  auf die Versions-Staleness, nicht auf die Link-Auflösung; die Link-Hälfte —
  die, auf die sich das zitierte Kurs-Argument („ein Link darauf färbt beim
  nächsten Bump ein Artefakt rot") bezieht — ist in `harness/conventions.md`
  §Baseline als **unbehandelt** deklariert: *„Eine Kollision bleibt und ist
  keine Ausnahme, sondern eine Klemme … Wer löscht, editiert Eingefrorenes; das
  ist der Preis des Löschens."* Die Behauptung, a-check habe *sie* behandelt,
  ist damit gegen zwei Repo-Artefakte falsch; sie trägt die Prämisse von
  `slice-176`s DoD-Punkt („ob `exempt-paths` danach schrumpfen kann").
- `verifizierbar`: nein — Sichtprüfung zweier Repo-Artefakte gegen den
  Slice-Text; kein Gate unterscheidet Modul-Zuständigkeiten.
- `klasse`: Zwei Sensor-Zuständigkeiten in einer Aussage verschmolzen

### F-3 — Eine neue Normregel in `modul-09`, die `AGENTS.md` §6 trifft, steht unter „Was nicht berührt"

- `kategorie`: HIGH
- `quelle`: `AGENTS.md` §6 (Minimal Agent Workflow, 8 Schritte); Upstream-Diff
  `v6.2.0`→`v6.5.0`
- `pfad`: `slice-174-…md` §3.3, Zeile „`modul-09` (+9) … Tabellen-Ausbau und
  Verweise auf die neuen Abschnitte"
- `befund`: `modul-09-implementierung.md` gewinnt in `v6.5.0` keinen
  Tabellen-Ausbau und keinen Verweis, sondern einen neuen Normabsatz zum
  8-Schritt-Workflow: *„Die Plan-Ausgabe in Schritt 4 nennt Out-of-Scope. …
  Nimmt der Lauf etwas mit, das §1 ausschließt, ist das eine Plan-Änderung und
  gehört vor den Code, nicht in den Bericht danach."* `AGENTS.md` §6 führt
  denselben 8-Schritt-Pfad und deklariert in §1 `modul-09` §Ziel-Form: AGENTS.md
  als seine Regelquelle. Die Änderung berührt a-check damit, steht aber in der
  Sektion „Was nicht berührt" und ist keinem der fünf Themen und keiner Etappe
  zugeordnet. DoD-Punkt 1 („Jede der 30 geänderten Dateien ist eingeordnet …
  mit Begründung je Zeile") ist abgehakt.
- `verifizierbar`: ja — `git diff -w v6.2.0 v6.5.0 -- lab/regelwerk/modul-09-implementierung.md`
  gegen `AGENTS.md` §6 gehalten.
- `klasse`: Normänderung als Verweis-Änderung eingeordnet

### F-4 — T-3 bekommt keine Etappe; §3.4 vergibt A–D nur an T-1, T-2, T-4 und T-5

- `kategorie`: MEDIUM
- `quelle`: DoD-Punkt 2 des Slice („Folge-Slice (mit Titel und Etappe) oder
  ausdrückliche Nicht-Handlung mit Grund")
- `pfad`: `slice-174-…md` §3.2 (T-3, Zeile „Folge-Slice: Kopieranleitung
  nachziehen") gegen §3.4 (Etappen-Tabelle)
- `befund`: Die Etappen-Tabelle führt A (Vendoring), B (Adaptions-Durchgang,
  „dabei §3.3-Restposten und T-5"), C (T-1) und D (T-2 + T-4). T-3 — §1 *Ziel
  und Abgrenzung* samt der Kopieranleitung in `AGENTS.md` §5 — erscheint in
  keiner Etappe und ist weder §3.3-Restposten noch T-5. Für T-3 steht damit ein
  Folge-Slice-*Vorschlag* ohne Adresse da; DoD-Punkt 2 ist abgehakt.
- `verifizierbar`: nein — Textabgleich §3.2 gegen §3.4.
- `klasse`: Thema ohne Etappen-Adresse

### F-5 — Zwei der 30 geänderten Dateien sind nirgends eingeordnet

- `kategorie`: MEDIUM
- `quelle`: DoD-Punkt 1 des Slice; Upstream-Diff
- `pfad`: `slice-174-…md` §3.1/§3.3 (vollständige Datei-Zuordnung)
- `befund`: `lab/templates/README.md` (+11/−6, 11 substanzielle Zeilen) und
  `lab/regelwerk/README.md` (+1/−1) kommen im Slice weder als Datei noch in
  einer Klasse vor — die Suche nach `templates/README` und `regelwerk/README`
  im Plan liefert null Treffer, und die Summe 8 (Substanz) + 9 (Null) + 7
  (§3.3-Liste) deckt 24 der 30 Dateien ab; vier weitere (`archiv-stub-slice`,
  `archiv-stub-welle`, `welle-results`, `review-report`) sind über T-1 gedeckt,
  diese zwei nicht. Inhaltlich tragen sie a-check-relevante Aussagen:
  `templates/README.md` benennt die neue Pflicht-Struktur §1 *Ziel und
  Abgrenzung*, die Umbenennung von §8 in *Sub-Area-Prüfungen und
  Modus-Begründung* und die neue Vorlagen-Klasse `harness/sensors/gate.template.md`;
  `regelwerk/README.md` hebt den Stand-Stempel von *Kurs-Welle 119 · 2026-09-05*
  auf *Kurs-Welle 128 · 2026-09-06*, den `harness/conventions.md` §Baseline
  wörtlich führt.
- `verifizierbar`: ja — `git diff --name-only v6.2.0 v6.5.0` gegen `grep` im
  Slice-Plan.
- `klasse`: Vollständigkeits-Zusage ohne Abgleich gegen die Namensliste

### F-6 — Die Sichtungs-Zahl in §9 („31 offene Einträge") reproduziert nicht

- `kategorie`: MEDIUM
- `quelle`: `docs/plan/planning/observations/`; Baseline `v6.2.0` ·
  `regelwerk/modul-06-roadmap.md` §Das Beobachtungs-Register (abgeleiteter
  Zähler)
- `pfad`: `slice-174-…md` §9, Block „Vorgelagert — offene Beobachtungen sichten"
- `befund`: Ausgezählt über `state.md` aller Einträge unter `BEO-GATE`,
  `BEO-HARNESS` und `BEO-PLAN` ergeben sich **33** Einträge mit Stand `offen`
  bzw. `offen (n×)` auf dem Stand vor dieser Range (`2297c3e`) und **34**
  danach; der Plan nennt **31**. Unter keiner mir konstruierbaren Lesart
  („offen" gesamt inkl. `BEO-SPEC`/`BEO-KERN`: 36) ergibt sich 31. Die
  inhaltliche Aussage des Blocks — vier einschlägige Einträge, keiner an der
  Schwelle — habe ich dagegen bestätigt: alle acht Einträge mit ≥ 3
  Evidence-Dateien stehen auf `verkörpert` oder `geplant`.
- `verifizierbar`: ja — Auszählung der `state.md`-Dateien, auch historisch über
  `git ls-tree` / `git show`.
- `klasse`: Zahl als gemessen ausgewiesen, nicht reproduzierbar

### F-7 — Risiko 3 wird als „entfallen" gestrichen, obwohl die zitierte Messung die gestellte Frage nicht beantwortet

- `kategorie`: MEDIUM
- `quelle`: Baseline `v6.2.0` · `regelwerk/modul-05-planning-harness.md`
  §Offene Risiken werden bei Closure aufgelöst (*entfallen* = die Beobachtung
  kann nicht mehr auftreten)
- `pfad`: `slice-174-…md` §7, dritter Spiegelstrich
- `befund`: Das Risiko fragt, ob die `sensors/`-Ziel-Form a-checks
  Tabellen-Praxis *widerspricht*. Belegt wird eine andere Frage — dass die
  Tabellenzeile Default bleibt und die Datei nur den Überhang aufnimmt (im
  Upstream bestätigt). Zur Widerspruchs-Frage steht im Bestand mindestens ein
  offener Punkt: die Ziel-Form verlangt `kein Gate` „in der Spalte, die dort
  Bindung oder Charakter führt", und `AGENTS.md` §4 führt eine Tabelle mit den
  Spalten `| Target | Zweck |` — eine solche Spalte gibt es dort nicht;
  außerdem liegt a-checks autoritative Gate-Tabelle nach `.d-check.yml`
  (`targets.authority: AGENTS.md`) in `AGENTS.md`, während die Ziel-Form den
  Index in `harness/README.md` §Sensors verortet und aus `AGENTS.md` heraus die
  Sensor-**Datei** direkt adressiert sehen will. Dieselbe Risiko-Klasse steht in
  `slice-177` §7 weiterhin als offenes Risiko.
- `verifizierbar`: nein — Vergleich der Ziel-Form gegen zwei Repo-Tabellen.
- `klasse`: Risiko gestrichen auf Basis einer benachbarten Messung

### F-8 — `welle-15` führt zwei Etappen-Nummerierungen; die Aufzählung in §1 kennt die Zitier-Form-Etappe nicht

- `kategorie`: MEDIUM
- `quelle`: Baseline `v6.2.0` · `regelwerk/modul-06-roadmap.md`
  §Roadmap-Struktur (die Welle-Datei trägt Ziel, Trigger, Closure-Kriterien)
- `pfad`: `welle-15-regelwerk-v650-migration.md` §1 gegen §4 und §6
- `befund`: §1 zählt vier Etappen **1–4** auf (Messen · Vendoring ·
  Adaptions-Durchgang · `harness/sensors/`), §4 benennt dieselben Slices als
  Etappe **A**, **C** und **D**, §6 verweist wieder auf „Etappe 4". Die
  Aufzählung in §1 enthält die Zitier-Form (T-1, `slice-176`, Etappe C)
  überhaupt nicht; sie stammt aus der Eröffnung (`61b7661`) und wurde nach der
  Analyse (`dc0e58b`) nicht nachgezogen, obwohl §4 derselben Datei in demselben
  Commit den A–D-Schnitt übernommen hat.
- `verifizierbar`: nein — Textabgleich innerhalb einer Datei.
- `klasse`: Zwei Nummerierungen desselben Schnitts nebeneinander

### F-9 — Der Start-Trigger von `slice-175` ist das Ergebnis von `slice-175`

- `kategorie`: MEDIUM
- `quelle`: Baseline `v6.2.0` · `regelwerk/modul-06-roadmap.md`
  §Roadmap-Regeln („Der Start-Trigger darf kein Ergebnis dieser Welle sein …
  Test: Steht der Trigger in der Slice-Liste dieser Welle, ist er falsch
  platziert")
- `pfad`: `open/slice-175-etappe-a-vendoring-v650.md` §5, Zeile „**Start**
  (`open` → `in-progress`): Etappe A (Vendoring) liegt in `done/`"
- `befund`: `slice-175` **ist** Etappe A (Kopf: „Etappe **A** des Schnitts in
  §3.4"); sein Start-Trigger verlangt, dass Etappe A bereits in `done/` liegt.
  Der Trigger kann nie eintreten, bevor der Slice läuft. In `slice-176` und
  `slice-177` ist derselbe Satz korrekt (dort ist A eine fremde Etappe) — die
  drei §5-Blöcke sind wortgleich.
- `verifizierbar`: nein — Textvergleich Kopf gegen §5.
- `klasse`: Start-Trigger zeigt auf das eigene Ergebnis

### F-10 — Die Kurs-Aussage, das Ausnahme-Ventil sei „eine Gate-Senkung mit eigener Begründungslast", ist nirgends eingeordnet

- `kategorie`: MEDIUM
- `quelle`: `AGENTS.md` §3.6 („Jede Schwellen-Senkung (Coverage, Linter-Strenge,
  Prüfregel) ist ein ADR"); Upstream `grundlagen-harness-dateien.md` in `v6.5.0`
- `pfad`: `slice-174-…md` §3.2, T-1
- `befund`: Derselbe neue Abschnitt, aus dem T-1 zitiert, schließt mit: *„Steht
  die Adresse erst im eingefrorenen Artefakt, bleiben zwei Wege: es doch
  anfassen … oder ein Ausnahme-Ventil im Prüfbereich, also eine Gate-Senkung mit
  eigener Begründungslast."* a-check hat mit `slice-173` genau ein solches
  Ventil gesetzt (fünf `exempt-paths`-Klassen im Modul `versions`); keine der 39
  ADRs nennt `exempt-paths` oder `version-stale`. Ob das unter §3.6 fällt, ist
  ein Urteil — der Satz kommt in keinem der fünf Themen und nicht in §3.3 vor,
  die Frage wird also nicht gestellt.
- `verifizierbar`: ja (Teil) — `grep -rl 'exempt-paths\|version-stale'
  docs/plan/adr/` liefert null Treffer; die Einordnung selbst bleibt Urteil.
- `klasse`: Zitierter Absatz nur zur Hälfte ausgewertet

### F-11 — Zwei Zahlenbasen für dieselben Dateien, ohne die Basis zu nennen

- `kategorie`: LOW
- `quelle`: innere Konsistenz des Slice
- `pfad`: `slice-174-…md` §2 (Posten-Tabelle) gegen §3.2, T-2, erster Absatz
- `befund`: §2 nennt `gate.template.md` mit **+81**,
  `grundlagen-harness-dateien.md` mit **+123** und `harness/README.template.md`
  mit **+33** (roh, `--numstat`); T-2 nennt für dieselben Dateien **+80**,
  **+122** und **+32** (nach Abzug je einer Trennzeile). Beide Zahlen sind für
  sich richtig, die Basis wird an keiner der beiden Stellen genannt.
- `verifizierbar`: ja — `git diff --numstat` gegen die Trennzeilen-Zählung.
- `klasse`: Zwei Messbasen ohne Deklaration

### F-12 — „16 von 40 Gate-Zeilen" zählt Nicht-Gates mit und lässt die zweite Tabelle aus

- `kategorie`: LOW
- `quelle`: T-4 desselben Slice (drei Targets sind ausdrücklich keine Gates);
  `welle-15` §1 Etappe 4 („`AGENTS.md` §4 **und** `harness/README.md` §Sensors")
- `pfad`: `slice-174-…md` §3.2, T-2, Block „Gemessen im Bestand"
- `befund`: Die Zeichenzahlen selbst reproduzieren exakt — `make doc-check`
  1871, `make symlink-check` 1106, genau 16 der 40 Zeilen über 250 Zeichen
  (nachgemessen über die Zweck-Zelle inklusive der umschließenden Leerzeichen).
  Die Grundgesamtheit „40 Gate-Zeilen" enthält jedoch die drei Targets, die T-4
  zwei Absätze später als Nicht-Gates führt, und zwei davon (`make archive-wave`
  727, `make regelwerk-check` 257) liegen unter den 16. Die Sensors-Tabelle in
  `harness/README.md`, die dieselbe Ziel-Form trifft und die `welle-15` §1 als
  betroffen nennt, ist nicht mitgemessen. `slice-177` §4 friert die Zahl 16 als
  Bezugsmenge ein.
- `verifizierbar`: ja — Auszählung der Zweck-Zellen in `AGENTS.md` §4.
- `klasse`: Grundgesamtheit einer Messung nicht gegen die eigene Klassifikation gehalten

### F-13 — Die Nicht-Gate-Menge ist behauptet, nicht gemessen

- `kategorie`: LOW
- `quelle`: `AGENTS.md` §4; Upstream `modul-13-quality-gates.md` §Die dritte
  Lage („etwas, das *sagt*, was ein schreibender Lauf täte")
- `pfad`: `slice-174-…md` §3.2, T-4; `open/slice-177-…md` §4, erster DoD-Punkt
- `befund`: T-4 nennt drei Nicht-Gates (`slice-mv`, `archive-wave`,
  `regelwerk-check`) ohne Messung, während T-2 unmittelbar davor misst. Nach dem
  im selben Absatz zitierten Kriterium fallen mindestens zwei weitere Zeilen
  derselben Tabelle darunter: `make doc-repair` gibt einen unified diff auf
  stdout aus (das Lehrbuchbeispiel der Regel) und `make record-gates` erzeugt
  einen Working-Tree-Hash als Nachweis. `slice-177`s DoD trägt die Dreier-Liste
  als abgeschlossene Menge. Nebenbei: bei `make archive-wave` steht das
  „**kein Gate**" nicht als Vorspann, sondern gegen Ende der Zelle — die
  Beschreibung „fett gesetzter Prosa-Vorspann" trifft nur zwei der drei.
- `verifizierbar`: ja — Sichtprüfung der Zweck-Zellen in `AGENTS.md` §4 gegen
  das Kriterium.
- `klasse`: Menge behauptet, während die Nachbar-Menge gemessen wurde

### F-14 — Zwei zugeordnete Dateien tragen je eine Änderung, die in keinem Thema auftaucht

- `kategorie`: LOW
- `quelle`: DoD-Punkt 1 („mit Begründung je Zeile, nicht als Sammelurteil")
- `pfad`: `slice-174-…md` §3.2, T-3 und T-4
- `befund`: `modul-13` ist über T-4 zugeordnet, gewinnt in `v6.5.0` aber **drei**
  neue Absätze; der dritte („Ein Gate ohne seine Grenze behauptet ebenfalls zu
  viel" — die Grenze gehört mit dem *Kommando* benannt, „nicht mit einer
  eingefrorenen Zahl, die beim nächsten Commit falsch ist") kommt in keinem
  Thema vor, obwohl `AGENTS.md` §4 solche Zahlen führt (`make doc-structure`
  „fünf Regeln", `make dcheck-phrase-selftest` „vier Kontrollen",
  `make ci-range-selftest` „vier Fälle"). `slice.template.md` ist über T-3
  zugeordnet, T-3 nennt aber nur die §1-Änderung; die Umbenennung von §8 in
  *Sub-Area-Prüfungen und Modus-Begründung* bleibt unerwähnt.
- `verifizierbar`: ja — `git diff -w v6.2.0 v6.5.0` je Datei gegen die
  Themen-Texte.
- `klasse`: Datei zugeordnet, Änderung nicht

### F-15 — Zwei Schwächen in den Triggern von `welle-15`

- `kategorie`: LOW
- `quelle`: Baseline `v6.2.0` · `regelwerk/modul-06-roadmap.md`
  §Roadmap-Regeln (beobachtbar = ein anderer Mensch entscheidet ohne Rückfrage)
  und §Wann Arbeit eine Welle braucht (das *Mehr* über die Slice-DoDs)
- `pfad`: `welle-15-regelwerk-v650-migration.md` §2 (zweiter Spiegelstrich) und
  §3 (Einleitungssatz)
- `befund`: Der zweite Start-Trigger belegt sich mit „dieses Gespräch,
  2026-09-07" — ein repo-externes Ereignis, das ein Dritter aus dem Repo nicht
  prüfen kann; der erste Trigger (`v6.5.0` ist veröffentlicht) trägt allein.
  Der Satz, der das *Mehr* der Closure benennt, nennt genau die beiden
  Sensor-Läufe (`doc-check` ohne `version-stale`, `symlink-check` grün), die
  `slice-175` §4 bereits als eigene DoD-Punkte führt; das tatsächliche Mehr —
  `make ci` auf dem finalen Stand, nach allen Etappen — steht erst im vierten
  Aufzählungspunkt darunter.
- `verifizierbar`: nein — Vergleich §3 der Welle-Datei gegen §4 von `slice-175`.
- `klasse`: Trigger-Beleg außerhalb des Repos · benanntes *Mehr* deckt sich mit einer Slice-DoD

### F-16 — `make trace-check` ist in diesem Arbeitsbaum nicht lauffähig (außerhalb der Range)

- `kategorie`: INFO
- `quelle`: `AGENTS.md` §5; `docs/plan/planning/observations/BEO-GATE/trace-check-lokal-nicht-befragbar/`
- `pfad`: `Makefile` `trace-check`; `observations/BEO-GATE/trace-check-lokal-nicht-befragbar/state.md`
- `befund`: `make trace-check` bricht mit `d-check: error:
  Range-Basis-Vorfahren nicht lesbar: object not found` ab — mit
  `RANGE=61b7661^..HEAD`, mit `RANGE=2297c3e..HEAD` und mit der Default-Range;
  `git fsck --connectivity-only` meldet nur unreferenzierte Objekte. Die
  Traceability der drei Commits konnte deshalb nur per Sichtprüfung bestätigt
  werden (alle drei nennen `slice-174`, zwei zusätzlich `welle-15`). Der Ausfall
  ist bereits als Beobachtung geführt; deren `state.md` sagt „offen (1×)",
  während das `evidence/`-Verzeichnis **zwei** Dateien trägt (`slice-170`,
  `slice-172`) — nach `modul-06` ist der Zähler die Zahl der Evidence-Dateien,
  also 2. Beides liegt außerhalb dieser Commit-Range.
- `verifizierbar`: ja — `make trace-check` (Exit 2); `ls …/evidence/` gegen
  `state.md`.
- `klasse`: Zustandsfeld gegen abgeleiteten Zähler veraltet

## Negativbefunde

- geprüft, ohne Befund: **Ausgangsmessung §2.** `git diff --stat v6.2.0 v6.5.0
  -- lab/regelwerk lab/templates` liefert exakt 30 Dateien, `+572/−149`; alle
  acht Zeilen der Posten-Tabelle stimmen zeichengenau mit `--numstat` überein,
  inklusive der als Korrektur nachgetragenen Zeile `grundlagen-harness-dateien.md`
  (+123/−14). Gelöscht wurde nichts, neu ist genau
  `templates/harness/sensors/gate.template.md` (+81).
- geprüft, ohne Befund: **36 Trennzeilen.** Mit der im Slice angegebenen Regex
  je Datei nachgezählt: Summe exakt **36**.
- geprüft, ohne Befund: **Die neun Null-Substanz-Dateien.** `modul-07`,
  `modul-08`, `modul-10`, `modul-11`, `modul-12`, `modul-14`, `modul-15`,
  `grundlagen-klassifikation`, `grundlagen-durchsetzungsschicht` — der
  Zeilen-Diff aller neun besteht ausschließlich aus `|---|---|` → `| --- | --- |`.
  Substanzfrei, wie behauptet.
- geprüft, ohne Befund: **Arithmetik der Klassen-Tabelle §3.1.** 8 + 13 + 9 = 30
  Dateien; 436 + 100 + 0 = 536 = 572 − 36. Die Klassen-*Grenzen* („≥ 18",
  „1–11") passen allerdings nicht zu den Mengen: `grundlagen-bootstrap` und
  `review-report.template` tragen je 18 und stehen in der 1–11-Klasse. Da F-1
  die Grundlage ohnehin verschiebt, ist das hier nur notiert, nicht als
  eigenes Finding geführt.
- geprüft, ohne Befund: **Die Zeichen-Messung in T-2.** 1871 und 1106
  reproduzieren exakt, ebenso „16 von 40" (nächstkleinerer Wert 241 — die
  Schwelle 250 ist eindeutig). Einwand zur Grundgesamtheit siehe F-12.
- geprüft, ohne Befund: **T-1 wörtliches Zitat.** Der Satz „Der vendored Baum
  trägt genau einen Tag …" steht wortgleich in
  `templates/docs/reviews/review-report.template.md` in `v6.5.0`; die vier
  genannten Ziel-Formen tragen den Zitier-Form-Block tatsächlich, drei davon in
  der Kurzfassung, `review-report` in der langen. Der Einwand betrifft die
  Deutung (F-2), nicht das Zitat.
- geprüft, ohne Befund: **T-2 wörtliches Zitat** („Ob der Überhang schon unter
  der Tabelle steht oder in die Zelle gedrängt wurde …") — wortgleich in
  `grundlagen-harness-dateien.md` in `v6.5.0`.
- geprüft, ohne Befund: **T-5.** Der neue Abschnitt *Die zweite Richtung:
  Anforderung → Beleg* trägt den zitierten Satz; die Behauptung, a-check
  deklariere nirgends, welche Verweis-Quellen entlastend gelten, hält stand:
  `.d-check.yml` führt keinen `trace:`-Block, und `doc-trace`/`doc-complete`
  rufen `--trace` ohne Quellen-Deklaration auf.
- geprüft, ohne Befund: **Vergleichszahl des vorigen Sprungs.** `v6.0.0` →
  `v6.2.0` ergibt im Klon 9 Dateien, `+53/−5` — identisch mit `slice-164` §2 und
  mit dem Slice-Text. Auf substanzieller Basis (`-w`) bleibt es bei 53, das
  Verhältnis wäre 452 : 53 ≈ 8,5 statt „rund zehnmal"; die Größenordnungs-Aussage
  trägt in beiden Rechnungen.
- geprüft, ohne Befund: **Beobachtungs-Register-Eintrag.** `BEO-GATE/versions-sensor-trifft-planungs-vorgriff`
  hat die drei Dateien in der vorgeschriebenen Rollenteilung; `observation.md`
  trägt Bezeichnung und Sub-Area (`Gate-/Werkzeug-Schicht`, in
  `harness/conventions.md` mit Kürzel `GATE` deklariert), `state.md` den
  veränderlichen Stand, `evidence/slice-174.md` genau einen Beleg. Die
  Zählregel „zwei Funde, ein Vorgang, ein Beleg" ist korrekt angewandt und
  gegen `modul-06` belegt. Der Text benennt die Grenze des Sensors, statt sie zu
  behaupten.
- geprüft, ohne Befund: **Vier einschlägige Register-Einträge in §9.** Alle vier
  existieren, ihre Stände stimmen, und keiner der acht Einträge mit ≥ 3
  Evidence-Dateien steht auf `offen` — „keiner an der Schwelle" trifft zu. Zur
  Gesamtzahl siehe F-6.
- geprüft, ohne Befund: **Commit-Hygiene.** Alle drei Commits tragen den Scope
  `(planning)` und berühren ausschließlich `docs/plan/planning/`
  (`make commit-scope-check RANGE=61b7661^..HEAD` Exit 0). Der Lifecycle-Commit
  `0518e90` ist ein reiner Rename (`{open => in-progress}/…` bei 0 Zeilen
  Änderung), `AGENTS.md` §3.3 gewahrt; die mitgeführte Roadmap-Zeile ist die von
  `make slice-mv` gezogene Referenz, nicht eine Inhaltsänderung der bewegten
  Datei.
- geprüft, ohne Befund: **Roadmap-Fortschreibung.** Der Welle-Zeiger steht unter
  *Offene Wellen*, der Ruhe-Marker weicht in `0518e90` der Nennung des Slice —
  beide Hälften der Sektion sind unabhängig behandelt, wie `modul-06`
  §Roadmap-Struktur es beschreibt.
- geprüft, ohne Befund: **Slice-Form der drei neuen Slices.** Je zwei
  Liefer-Punkte plus die vier konstanten DoD-Zeilen (Größen-Regel ≤ 3 gewahrt),
  höchstens zwei Schichten, Kopffelder vollständig (`Welle:`, `Bezug:`,
  `Berührte Spec-Stellen:`, `Verantwortlich:`, `Autor:`), Review-DoD-Zeile mit
  der exakten Trigger-Phrase „unabhängiger Review". `make verify` Exit 0. Die
  §1/§7-Blöcke sind allerdings in allen drei Dateien wortgleich; ob eine
  Abgrenzung, die für drei verschiedene Etappen identisch lautet, noch eine
  Abgrenzung ist, ist ein Urteil und hier nur benannt.
- geprüft, ohne Befund: **Sieben aktive `MR`-Einträge.** `welle-15` §6 und
  `slice-174` §3.4 nennen sieben; die Tabelle *Aktive Adaptionen* in
  `harness/conventions.md` führt `MR-011`, `MR-012`, `MR-014`, `MR-015`,
  `MR-016`, `MR-019`, `MR-020` — sieben.
- geprüft, ohne Befund: **Etappe B ohne Slice.** Der fehlende Slice ist an drei
  Stellen ausdrücklich begründet (Slice §3.4/§8, `welle-15` §4) und nicht still
  ausgelassen; die drei angelegten Folge-Slices existieren als Dateien in
  `open/`, die Folge-Slice-Paarung ist damit erfüllbar.
- geprüft, ohne Befund: **Hard Rules `AGENTS.md` §3.** Kein Host-Toolchain-Aufruf
  im Diff, keine Suppression, keine ADR berührt, kein Spec-Stratum referenziert
  abwärts, kein Gate gelockert. Keine erfundene ID: alle im Diff genannten
  `slice-`/`MR-`/`BEO-`-Kennungen lösen auf (`make doc-check` Exit 0).
- geprüft, ohne Befund: **Gate-Lage.** `make doc-check`, `make doc-targets` und
  `make verify` laufen auf dem Stand `HEAD` je mit Exit 0 (eigene Läufe, Ausgabe
  in Dateien, Exit-Code getrennt gelesen). `make gates` und `make image-test`
  habe ich nicht gefahren — DoD-Verifikation ist Verifier-Sache (Modul 11).
- geprüft, **ohne Befund und ohne Repo-Änderung**: Für die Nachrechnung wurde
  ausschließlich ein Klon des Kurs-Repos außerhalb des Arbeitsbaums gelesen;
  `git status --porcelain` ist vor und nach dem Review leer.

## Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 3 |
| MEDIUM | 7 |
| LOW | 5 |
| INFO | 1 |

**Finding-Klassen dieses Laufs:** Bereinigung deckt nur die gesuchte
Formatierungs-Klasse · Zwei Sensor-Zuständigkeiten in einer Aussage verschmolzen ·
Normänderung als Verweis-Änderung eingeordnet · Thema ohne Etappen-Adresse ·
Vollständigkeits-Zusage ohne Abgleich gegen die Namensliste · Zahl als gemessen
ausgewiesen, nicht reproduzierbar · Risiko gestrichen auf Basis einer
benachbarten Messung · Zwei Nummerierungen desselben Schnitts nebeneinander ·
Start-Trigger zeigt auf das eigene Ergebnis · Zitierter Absatz nur zur Hälfte
ausgewertet · Zwei Messbasen ohne Deklaration · Grundgesamtheit einer Messung
nicht gegen die eigene Klassifikation gehalten · Menge behauptet, während die
Nachbar-Menge gemessen wurde · Datei zugeordnet, Änderung nicht · Trigger-Beleg
außerhalb des Repos · Zustandsfeld gegen abgeleiteten Zähler veraltet

## Verdikt

**Merge-blockierend:** ja — drei HIGH und sieben MEDIUM.

Der Slice liefert genau das, wofür er geschnitten ist, und liefert es an den
schwierigsten Stellen sauber: die Roh-Zahlen, die Trennzeilen-Zählung, die
Zeichen-Messung in `AGENTS.md` §4, die neun substanzfreien Dateien und die
Register-Sichtung habe ich nachgerechnet und bestätigt gefunden — einschließlich
der beiden im Plan selbst korrigierten eigenen Fehler. Die Findings treffen
nicht die Sorgfalt, sondern **eine Systematik**: Zwei der drei HIGH und drei der
sieben MEDIUM entstehen daran, dass die Klassifikation *entlang der
Datei-Namensliste* statt entlang der Änderungen läuft. Wo eine Datei einem Thema
zugeordnet ist, gilt sie als abgehandelt (F-14); wo sie in keiner der drei
Klassen steht, fällt sie ganz heraus (F-5); und wo die Bereinigung eine
Formatierungs-Klasse kennt, wird die zweite zu Substanz gezählt (F-1) — was
`grundlagen-begriffe.md` mit 2 realen Zeilen in die Substanz-Klasse hebt und
fünf Dateien ohne jede Inhaltsänderung als „Tabellen-Ausbau" führt.

F-1 verschiebt keine Schlussfolgerung — dass wenige Dateien die Substanz tragen,
wird auf der bereinigten Basis eher deutlicher —, wohl aber die Grundlage, auf
der `slice-177` seinen Umfang schätzt. F-2 und F-3 dagegen wirken in die
Folge-Slices hinein: F-2 trägt die Prämisse eines DoD-Punkts von `slice-176`,
F-3 lässt eine Regeländerung an `AGENTS.md` §6 ohne Adresse, und F-4 lässt das
Thema, zu dem sie gehört, ganz ohne Etappe.

**Übergabe:** Findings gehen an den Implementer; die Finding-Klassen gehören in
die Closure §7 und von dort in den Zähler — insbesondere „Datei zugeordnet,
Änderung nicht" und „Bereinigung deckt nur die gesuchte Formatierungs-Klasse"
sind Kandidaten für einen Register-Eintrag, weil sie in einem Slice viermal
aufgetreten sind (ein Vorgang, ein Beleg). Dieser Report ist Lauf-Beleg, keine
Verifikation — DoD-/Spec-Konformität prüft der Verifier separat (Modul 11).
