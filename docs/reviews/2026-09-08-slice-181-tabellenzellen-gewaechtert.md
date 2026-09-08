# Review-Report: slice-181 — 2026-09-08

**Review-Art:** unabhängiger Lauf — anderes Kontextfenster als die
Implementierung, Gegenstand nicht selbst verfasst (`v6.5.0` ·
`regelwerk/modul-08-agentenrollen.md` §Rollen-Regeln). Geprüft gegen
Slice-Plan, `AGENTS.md` §3/§4/§5/§6 und den gemessenen Repo-Bestand.

**Gegenstand:** Commit `7b56af7` („feat(harness): slice-181 -- Tabellenzellen
der Gate-Tabellen gewaechtert") — `.d-check.yml`, `AGENTS.md`,
`harness/README.md`, neu `harness/sensors/gate-consistency.md`.

**Skill:** `.harness/skills/reviewer.md` @ Stand `7b56af7` (unverändert seit
Anlage) · <!-- d-check:ignore -->
**Modell:** claude-opus-5[1m] · **Datum:** 2026-09-08

**Eingangs-Kontext:**

- `.harness/skills/reviewer.md`; `AGENTS.md` §3, §4, §5, §6
- Slice-Plan slice-181 (in `in-progress/`, Stand `7b56af7`)
- `tools/gate-consistency.sh`, `tools/dcheck-phrase-selftest.sh`, `Makefile`
- `harness/sensors/*.md` (alle 15 Dateien, gezielt fünf davon)
- `v6.5.0` · `templates/harness/README.template.md` §Sensors (Feedback-Gates)
- d-check `v0.74.1` (Digest-Pin), Modul `structure` — 14 eigene Läufe gegen
  Kopien unter `/tmp/.../scratchpad/`, nie gegen das Arbeitsrepo

**Mess-Werkzeug und sein Geltungsbereich:** Zellenlängen wurden zweifach
erhoben — mit einem eigenen `awk`-Splitter, der escapte Pipes (`\|`) vor dem
Zerlegen schützt, und autoritativ mit dem gepinnten d-check selbst über
variierte `cell-max-chars`. Beide Verfahren stimmen exakt überein (Probe: bei
Schwelle 241 genau ein Befund, bei 242 keiner — dieselbe Zelle, die der
`awk`-Lauf mit 242 ausweist). Der Geltungsbereich der Messung sind die **66
Zellen der zwei konfigurierten Spalten** (26 `Vertrag` in `harness/README.md`,
40 `Zweck` in `AGENTS.md` §4); Prosa neben den Tabellen und die dritte Tabelle
in `harness/conventions.md` sind darin **nicht** enthalten und werden unten
getrennt behandelt.

---

## Findings

### F-1 — Die Kalibrierungs-Zahlen in der Konfiguration sind um je eins zu niedrig

- `kategorie`: HIGH
- `quelle`: `AGENTS.md` §5 („die Schwelle ist am Bestand **gemessen**"),
  Reviewer-Skill („nachweislich falsche Tatsachenbehauptung")
- `pfad`: `.d-check.yml`:332–334 (Kommentarblock zu Regel (6)); wortgleich in
  der Commit-Message von `7b56af7`
- `befund`: Der Kommentar sagt „Bei 200 waeren es 17 Befunde gewesen, bei 250
  acht"; gemessen sind es **18** bzw. **9**. Beide Zahlen sind exakt um die
  eine Zelle zu niedrig, deren Übersehen derselbe Commit als Korrektur
  beschreibt (`AGENTS.md`:188, `make slice-mv`, 421 Zeichen, escapte Pipes) —
  die vom Sensor korrigierte Zählung ist in der Erzählung angekommen, in den
  Zahlen nicht.
- `verifizierbar`: ja — d-check `v0.74.1`, Modul `structure`, gegen den Stand
  `7b56af7^` plus der neuen Konfiguration: `cell-max-chars: 250` ⇒ Exit 1, 9
  `section-cell-oversized` (`AGENTS.md`:159/170/178/188/198,
  `harness/README.md`:74/85/87/91); `cell-max-chars: 200` ⇒ Exit 1, 18 Befunde.
  Die neun Fundstellen sind deckungsgleich mit den neun Zellen, die der Commit
  ändert.
- `klasse`: gemessene Zahl durch die vor-korrigierte ersetzt

### F-2 — „genau EINE der 66 Zellen trägt mehr als einen Satz" ist falsch

- `kategorie`: HIGH
- `quelle`: Reviewer-Skill („nachweislich falsche Tatsachenbehauptung");
  `v6.5.0` · `templates/harness/README.template.md` §Sensors (das
  Satz-Kriterium, gegen das hier argumentiert wird)
- `pfad`: `.d-check.yml`:322–324
- `befund`: Der Kommentar trägt diese Zahl als *tragende* Begründung dafür,
  Zeichen statt Sätze zu messen. Gezählt über dieselben 66 Zellen führen
  **15** mehr als einen Satz (Stand `7b56af7^`) bzw. **16** (Stand `7b56af7`)
  — darunter `make doc-workflows`, `make version-coherence`,
  `make verify-observations`, `make doc-complete`, `make regelwerk-check` und
  `make archive-wave` (letztere mit drei Sätzen). Richtig ist am Satz nur der
  Halbsatz danach: die 354-Zeichen-Zelle trägt tatsächlich genau einen.
- `verifizierbar`: ja — Auszählung der Satzenden (`[.!?]` gefolgt von
  Leerzeichen oder Zellenende) über die 66 Zellen; kein Treffer ist eine
  Abkürzung, alle 15 bzw. 16 Fundstellen sind echte Satzgrenzen, jede einzeln
  nachlesbar. Kein Gate prüft das — die Fähigkeit `table.column` kennt laut
  `d-check --print-config` nur `cell-max-chars` und `cell-min-chars`, keine
  Satz-Option.
- `klasse`: Begründung stützt sich auf eine Zahl, die die eigene Menge nicht
  hergibt

### F-3 — Die neue Sensor-Datei nennt drei Prüfungen; das Target führt vier aus

- `kategorie`: HIGH
- `quelle`: `AGENTS.md` §4 („halluzinierte Gates sind die häufigste Form von
  Harness-Lüge"); Reviewer-Skill (Harness-Lüge)
- `pfad`: `harness/sensors/gate-consistency.md`:5–13; gleichlautend
  `AGENTS.md`:178 und `harness/README.md`:74 („Meta-Gate, drei Prüfungen")
- `befund`: `tools/gate-consistency.sh` lässt das Gate an **vier**
  Bedingungen rot werden — (3) `.d-check.yml`-Modulliste, (4)
  Pin-Konsistenz, (5) **`.PHONY`-Vollständigkeit** (`check_phony_complete
  Makefile || fail=1`, Zeile 324, angelegt in slice-068 gegen ein
  False-Green), (6) ADR-Index. Die Sensor-Datei führt (5) nicht auf, weder
  unter „Vertrag" noch unter „Grenze", obwohl sie als geschlossene Aufzählung
  formuliert ist („Drei Prüfungen, alle drei gegen dieselbe Klasse").
- `verifizierbar`: ja — `make gate-consistency` nennt die vierte in seiner
  eigenen Erfolgszeile: „gate-consistency ok: .d-check.yml-Module intakt, Pins
  konsistent, **.PHONY vollstaendig**, ADR-Index vollstaendig" (Exit 0). Kein
  Gate hält die Sensor-Datei gegen das Skript.
- `klasse`: Sensor-Datei als geschlossene Aufzählung, die eine Prüfung ausläßt

### F-4 — Die Sensor-Datei beschreibt Prüfung (1) als etwas, das das Skript nicht tut

- `kategorie`: HIGH
- `quelle`: Reviewer-Skill (Harness-Lüge / falsche Tatsachenbehauptung);
  AC-QA-02 (die Bindung, die das Skript für diese Prüfung selbst nennt)
- `pfad`: `harness/sensors/gate-consistency.md`:8–11, mittelbar :23–26
- `befund`: Die Datei sagt: „die konfigurierten Module **existieren im
  gepinnten Werkzeug** … eine Konfiguration, die auf ein verschwundenes Modul
  zeigt, läuft still ins Leere". Das Skript prüft nichts dergleichen: es
  verlangt, dass die `modules:`-Zeile die vier Namen `links`, `anchors`,
  `ids`, `matrix` **enthält**, und dass sie `external` **nicht** enthält —
  Letzteres ist die Hälfte, die den netzlosen `doc-check` als AC-QA-02-Beleg
  trägt, und sie fehlt in der Beschreibung ganz. Die daran anschließende
  „Grenze" (2) („Prüfung (1) misst Existenz, nicht Wirkung") grenzt damit eine
  Prüfung ab, die es in dieser Form nicht gibt.
- `verifizierbar`: ja — `tools/gate-consistency.sh`:306–318 ist die einzige
  Stelle, an der `modules` vorkommt; ein `grep -n "modules" ` über die Datei
  liefert nur diesen Block.
- `klasse`: Vertrag der Sensor-Datei weicht vom Skript ab

### F-5 — „fünf Regeln" steht an drei Stellen, seit diesem Commit sind es sechs

- `kategorie`: HIGH
- `quelle`: `AGENTS.md` §4 (Deklarations-Pflicht der Gates), Reviewer-Skill
  (falsche Tatsachenbehauptung)
- `pfad`: `AGENTS.md`:171 · `harness/README.md`:76 ·
  `harness/sensors/doc-structure.md`:5
- `befund`: Der Commit hängt zwei Konfigurations-Einträge als Regel (6) in das
  Modul `structure` (`.d-check.yml`:317–369). Alle drei Deklarationen des
  Targets nennen weiterhin „**fünf** Regeln: Größen-Regel, Closure-Struktur,
  Lerneintrag-Form, Kopffelder, AC-Form" und zählen die neue Zellengrenze
  nicht mit — die Sensor-Datei zusätzlich in ihrer „Grenze"-Sektion nicht.
  Damit läuft unter `make verify` eine Prüfregel, die kein Dokument des Repos
  ausweist.
- `verifizierbar`: ja — Abzählen der Regel-Kommentare `# (1)` … `# (6)` und
  der `- files:`-Einträge im `structure`-Block von `.d-check.yml` (sechs
  Regeln, sieben Einträge). Kein Gate prüft die Zahl.
- `klasse`: neue Gate-Regel ohne Nachzug ihrer Deklaration

### F-6 — Die neue `gates`-Zelle behauptet eine Deckung, die die Tabelle nicht hat

- `kategorie`: HIGH
- `quelle`: Reviewer-Skill (falsche Tatsachenbehauptung); `AGENTS.md` §4
- `pfad`: `harness/README.md`:87
- `befund`: Die Zelle ersetzt die frühere (korrekte, aber redundante) Liste
  durch „aggregiert die inneren Gates **dieser Tabelle** und schließt mit
  `record-gates`". `make suppression-check` hängt im `gates`-Aggregat
  (`Makefile`:194), kommt in `harness/README.md` aber an **keiner** Stelle vor
  — weder in der Sensors-Tabelle noch unter `### Nicht-Gates` noch in der
  Prosa darunter, die drei andere Ausnahmen namentlich benennt. Die zweite
  Quelle ist damit nicht abgeschafft, sondern durch eine falsche
  Mengenaussage ersetzt.
- `verifizierbar`: ja — `grep -n "suppression-check" harness/README.md`
  liefert keinen Treffer, `grep -n "^gates:" Makefile` nennt es.
- `klasse`: Zeiger-statt-Liste mit falsch gefasstem Geltungsbereich

### F-7 — Die Umsetzung nimmt mit, was §1 ausschließt, und liefert nicht, was §1 zusagt

- `kategorie`: HIGH
- `quelle`: `AGENTS.md` §6 Schritt 4 („Nimmt der Lauf etwas mit, das §1
  ausschließt, ist das eine **Plan-Änderung** und gehört vor den Code");
  `v6.5.0` · `regelwerk/modul-05-planning-harness.md` §Ziel-Form: Slice
- `pfad`: Slice-Plan slice-181 §1 (Zeilen 30–34 und 38–42), gegen `7b56af7`
- `befund`: §1 sagt zu, `table.column` für `harness/README.md` §Sensors **und
  `harness/conventions.md` §Aktive Adaptionen** zu konfigurieren, und schließt
  `AGENTS.md` §4 ausdrücklich aus („es wäre ein **anderer Vorgang** … Ihre
  Antwort ist Kürzen nach §3.7, nicht eine Grenze"). Der Commit tut das
  Umgekehrte: `AGENTS.md` §4 bekommt die Grenze **und** fünf gekürzte Zellen,
  `harness/conventions.md` kommt in `.d-check.yml` nicht vor. Die dort in §2
  vermessene Zelle steht unverändert bei **331** Zeichen
  (`harness/conventions.md`:128, MR-019). §1 und §3 („*(offen)*") des Plans
  sind im selben Commit unverändert. Dies ist ein Plan-vs-Code-Befund, keine
  DoD-Verifikation (Abgrenzung nach `docs/reviews/README.md`).
- `verifizierbar`: ja — `grep -n "conventions.md" .d-check.yml` trifft nur
  `links`/`versions`-Blöcke, keinen `structure`-Eintrag; die 331 Zeichen sind
  mit demselben Splitter gemessen wie die 66 Zellen.
- `klasse`: stille Weitung der Abgrenzung in beide Richtungen

### F-8 — Die zweite Tabelle liegt im Geltungsbereich, ihre Inhaltsspalte ist trotzdem ungewächtert

- `kategorie`: MEDIUM
- `quelle`: BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf (der Eintrag, den der
  Slice-Kopf als Anlass nennt)
- `pfad`: `.d-check.yml`:350–357; Gegenstand `harness/README.md`:96–108
  (`### Nicht-Gates`)
- `befund`: Der Selektor `section: "## Sensors (Feedback-Gates)"` **erfasst**
  die Untersektion `### Nicht-Gates` — die dort mitkonfigurierte
  `Bindung`-Regel feuert nachweislich auf ihren Zeilen. Die Inhaltsspalte
  dieser Tabelle heißt `Was es tut` und ist nicht adressiert; eine Zelle mit
  400 Zeichen passiert grün, und `section-column-missing` schweigt, weil die
  erste Tabelle der Sektion die Spalte `Vertrag` führt. Weder Konfiguration
  noch Sensor-Doku benennen, dass fünf Zeilen im Geltungsbereich liegen und
  für die Längen-Hälfte trotzdem ungeprüft sind.
- `verifizierbar`: ja — zwei Läufe gegen Kopien: (a) `Bindung`-Zelle in
  `harness/README.md`:105 geleert ⇒ Exit 1, `section-cell-undersized` auf
  Zeile 105; (b) Inhaltszelle derselben Zeile auf 400 Zeichen verlängert ⇒
  Exit 0, 0 Befunde.
- `klasse`: Regel im Geltungsbereich ohne Gegenstand, nicht deklariert

### F-9 — Offene Frage 4 des Plans ist weder beantwortet noch als Abweichung notiert

- `kategorie`: MEDIUM
- `quelle`: Slice-Plan slice-181 §2 Punkt 4 samt „Zu entscheiden (Punkt 4)"
- `pfad`: Slice-Plan slice-181 §2 (Zeilen 103–123), gegen `7b56af7`
- `befund`: Der Plan stellt fest, dass die Ziel-Form die zweite Tabelle mit
  `Tut was` überschreibt, a-check mit `Was es tut`, und entscheidet selbst:
  „der Spaltenname hat keinen [Vorteil] — dort ist Angleichen billiger als
  eine abweichende Konfiguration, die beim nächsten Vorlagen-Vergleich wieder
  auffällt." Der Commit gleicht nicht an, konfiguriert nicht und hinterlässt
  keine Notiz, warum die eigene Empfehlung fallen gelassen wurde. Von den fünf
  offenen Punkten aus §2 sind 1, 2, 3 und 5 beantwortet; 4 ist übergangen.
- `verifizierbar`: ja — `v6.5.0` · `templates/harness/README.template.md`:133
  führt `| Target | Tut was | Bindung |`, `harness/README.md`:102 führt
  `| Target | Was es tut | Bindung |`.
- `klasse`: offene Plan-Frage ohne Ausgang

### F-10 — Die zwei Mediane stehen ohne Zuordnung und in umgekehrter Reihenfolge

- `kategorie`: LOW
- `quelle`: `AGENTS.md` §5 („Geltungsbereich einer Messung")
- `pfad`: `.d-check.yml`:331–332
- `befund`: „Median der zwei Spalten 152 bzw. 158 Zeichen" nennt nicht, welche
  Zahl zu welcher Spalte gehört. Gemessen ist `Vertrag` = 158,5 und `Zweck` =
  152,5 — die Reihenfolge des Satzes ist damit die umgekehrte zu der, in der
  der Block die Spalten einführt (`Vertrag` zuerst).
- `verifizierbar`: ja — Median über die 26 bzw. 40 Zellen des Stands
  `7b56af7^`.
- `klasse`: Messwert ohne Zuordnung

### F-11 — Das Ziel-Form-Zitat ist über zwei Zwischensätze hinweg zusammengezogen

- `kategorie`: LOW
- `quelle`: `v6.5.0` · `templates/harness/README.template.md` §Sensors
  (Feedback-Gates), Zeilen 101–105
- `pfad`: `.d-check.yml`:319–321
- `befund`: Der Kommentar zitiert als zusammenhängenden Satz: „Braucht ein
  Gate mehr als EINEN SATZ — … —, wandert das nach harness/sensors/
  <target>.md". Im Original endet der Satz nach „wandert das nach."; die
  Pfadangabe folgt zwei Sätze später. Die Aussage bleibt getroffen, die Form
  behauptet Wörtlichkeit, die nicht vorliegt.
- `verifizierbar`: ja — Zeilenvergleich gegen die vendored Vorlage.
- `klasse`: elidiertes Zitat ohne Auslassungszeichen

### F-12 — Unter „Grenze — was das Grün nicht abdeckt" steht als Erstes eine Stärke

- `kategorie`: LOW
- `quelle`: `AGENTS.md` §3.7 (ein Kommentar trägt genau eine Klasse — hier
  Grenze vs. Abgrenzung)
- `pfad`: `harness/sensors/gate-consistency.md`:19–22
- `befund`: Punkt 1 der Grenzen-Sektion beschreibt, was Prüfung (3)
  **zusätzlich** leistet („die Gegenrichtung, die `doc-check` per Konstruktion
  nicht sieht"). Das ist eine Abgrenzung gegenüber einem anderen Sensor, keine
  Lücke des eigenen Grüns; die Sektion verspricht Lücken.
- `verifizierbar`: nein — Urteil über die Textklasse, kein Match.
- `klasse`: Grenzen-Sektion trägt eine Nicht-Grenze

### F-13 — Der `hint` benennt beim Feuern die falsche Spalte, die dokumentierte Grenze deckt nur die Befundklasse

- `kategorie`: LOW
- `quelle`: `AGENTS.md` §3.7 (Zusage/Grenze), eigener Lauf
- `pfad`: `.d-check.yml`:338–342 (Grenz-Notiz) gegen :358 und :368 (die zwei
  `hint`-Texte)
- `befund`: Die Notiz hält fest, dass der `hint` auch über
  `section-column-missing` und `section-cell-undersized` steht. Nicht
  festgehalten ist, dass er dabei auch die **Spalte** falsch benennt: über
  einem Befund zur Spalte `Bindung` steht „Zellengrenze der Vertrags-Spalte".
  Zudem ist der in der Notiz zitierte Text („Zelle ueber 250 Zeichen") nicht
  der Wortlaut des `hint`.
- `verifizierbar`: ja — Lauf gegen eine Kopie mit geleerter `Bindung`-Zelle:
  Befund `section-cell-undersized`, Spalte `Bindung`, Hinweistext
  „Zellengrenze der Vertrags-Spalte (max 250 Zeichen, nicht leer) …".
- `klasse`: dokumentierte Grenze schmaler als die gemessene

### F-14 — Bestand: `AGENTS.md` §4 nennt zwei Korpus-Kontrollen, Skript und Sensor-Datei eine

- `kategorie`: INFO
- `quelle`: `tools/dcheck-phrase-selftest.sh`:201
- `pfad`: `AGENTS.md`:183 gegen `harness/README.md`:79 und
  `harness/sensors/dcheck-phrase-selftest.md`:11
- `befund`: Die Zelle in `AGENTS.md` §4 sagt „zwei **Korpus**-Kontrollen gegen
  den `done/`-Bestand"; das Skript meldet „4 Werkzeug-Kontrollen … und 1
  Korpus-Kontrolle", und beide anderen Dokumente sagen eine. Der Commit
  berührt diese Zeile nicht (identisch in `7b56af7^`) — sie liegt aber
  mitten in der Tabelle, die dieser Slice vermessen und gekürzt hat, und
  unterschreitet die neue Grenze, ist also dauerhaft grün.
- `verifizierbar`: ja — Zeilenvergleich; kein Gate deckt die Aussage.
- `klasse`: Bestandsdivergenz in der berührten Tabelle

## Negativbefunde

- geprüft, ohne Befund: **Substanzverlust beim Kürzen** — alle neun geänderten
  Zellen einzeln gegen `7b56af7^` verglichen. Jede entfernte Aussage ist
  wiedergefunden: die vier `doc-check`-Grenzen samt `links.resolve-from` und
  `versions` in `harness/sensors/doc-check.md`:10/11/21/28/59; „Form, nicht
  Gültigkeit **des Tag-Kommentars**" in `harness/sensors/doc-workflows.md`:12;
  `BEO-NNN`/`Ehemals:` und „Lage/Existenz ungeprüft" in
  `harness/sensors/verify-observations.md`:8/16; `WELLE=`-Falle,
  Stub-Grenze und slice-157 in `harness/sensors/archive-wave.md`:12/22/30;
  Pin-Dateien, ADR-Index und die slice-018/087-Bindungen in der neuen
  `harness/sensors/gate-consistency.md`. Die aus `AGENTS.md` entfernte
  `slice-mv`-Aufrufsyntax (`SLICE=`/`TO=`) steht in `Makefile`:159 und läuft
  über `make help`; die entfernte `BEO-PLAN`-Herkunft steht in
  `harness/README.md`:105. Ausgenommen sind die in F-3, F-4 und F-6 genannten
  Punkte — dort ist Substanz nicht verschoben, sondern verloren bzw. falsch
  ersetzt worden.
- geprüft, ohne Befund: **Die vier Mutations-Proben sind reproduzierbar.** Auf
  Kopien unter dem Scratchpad, mit dem gepinnten d-check `v0.74.1`: (1)
  `Vertrag`-Zelle auf 300 Zeichen verlängert ⇒ Exit 1,
  `section-cell-oversized` auf ihrer Zeile; (2) `name: Vertrag` → `name: Tut
  was` ⇒ Exit 1, `section-column-missing` **auf der Abschnitts-Überschrift**
  (Zeile 47), nicht auf einer Tabellenzeile; (3) `Vertrag`-Zelle geleert ⇒
  Exit 1, `section-cell-undersized`; (4) unverändert ⇒ Exit 0, 0 Befunde. Alle
  vier Aussagen des Commits treffen zu.
- geprüft, ohne Befund: **Die Begründung für `cell-min-chars: 1` auf der
  `Bindung`-Spalte trägt.** Gegenprobe in beide Richtungen: eine Tabellenzeile
  ohne dritte Zelle meldet mit dem Selektor `section-column-missing`, ohne ihn
  Exit 0 und 0 Befunde. Die Zusage „eine Zeile ohne die Spalte geht nicht
  stumm durch" ist damit gemessen, nicht behauptet.
- geprüft, ohne Befund: **„max 242 nach der Auflösung"** — d-check meldet bei
  `cell-max-chars: 242` null Befunde und bei 241 genau einen
  (`harness/README.md`:88, `make image-test`).
- geprüft, ohne Befund: **Grundmenge 66** — 26 Datenzeilen in
  `harness/README.md`:69–94 plus 40 in `AGENTS.md` §4.
- geprüft, ohne Befund: **Die Korrektur „in beiden Formen" → „in allen drei
  Formen"** stimmt mit `Makefile`:159 überein („drei Verweis-Formen,
  Auswahl+Ersetzung selbstgetestet seit slice-180").
- geprüft, ohne Befund: **Die Zuordnung „acht plus eine neunte"** — die neun
  Befunde bei Schwelle 250 auf dem Stand `7b56af7^` sind exakt die neun
  Zellen, die der Commit ändert; keine Fundstelle blieb offen, keine wurde
  ohne Befund angefasst. Nur die *Zahl* im Text ist falsch (F-1).
- geprüft, ohne Befund: **`AGENTS.md` §3.1** — der Commit fügt keinen Aufruf
  einer Host-Toolchain hinzu; die neue Regel läuft über den bereits
  digest-gepinnten d-check im Container.
- geprüft, ohne Befund: **`AGENTS.md` §3.2** — keine Suppressions berührt.
- geprüft, ohne Befund: **`AGENTS.md` §3.6** — die Schwelle 250 gegenüber dem
  Vorschlag 200 ist **keine** ADR-pflichtige Lockerung: 200 war ein Wert aus
  einem Maintainer-Hinweis, keine gesetzte Schwelle, und der Plan gibt in §2.1
  ausdrücklich vor, gegen den Bestand zu messen statt zu setzen. Die
  Entscheidungs*form* ist damit regelkonform; belastet ist nur ihr Beleg
  (F-1). Auch die Richtung stimmt: 18 → 9 Befunde ist derselbe Sprung, den die
  falschen Zahlen 17 → 8 behaupten.
- geprüft, ohne Befund: **`AGENTS.md` §3.5 / Referenz-Richtung** — keine ADR
  berührt, kein Spec-Stratum referenziert abwärts, keine Kennung erfunden.
- geprüft, ohne Befund: **Verlinkung der Target-Zelle** — die Ziel-Form
  verlangt beim Auslagern den Link von der Target-Zelle auf die Sensor-Datei;
  `gate-consistency` trägt ihn jetzt in beiden Tabellen.
- geprüft, ohne Befund: **Läufe auf dem Commit-Stand** — `make verify`,
  `make doc-check`, `make doc-targets` und `make gate-consistency` je Exit 0.
  `make gates` in Gänze (Go-Stages) wurde **nicht** gefahren; die Aussage des
  Commits dazu ist insoweit ungeprüft und Sache des Verifiers.
- geprüft, ohne Befund: **Kein Schreibzugriff auf den Gegenstand** — alle 14
  Sensor-Läufe liefen gegen Kopien unter dem Scratchpad; das Arbeitsrepo ist
  unverändert außer um diesen Report.

## Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 7 |
| MEDIUM | 2 |
| LOW | 4 |
| INFO | 1 |

**Finding-Klassen dieses Laufs:** gemessene Zahl durch die vor-korrigierte
ersetzt · Begründung stützt sich auf eine Zahl, die die eigene Menge nicht
hergibt · Sensor-Datei als geschlossene Aufzählung, die eine Prüfung ausläßt ·
Vertrag der Sensor-Datei weicht vom Skript ab · neue Gate-Regel ohne Nachzug
ihrer Deklaration · Zeiger-statt-Liste mit falsch gefasstem Geltungsbereich ·
stille Weitung der Abgrenzung in beide Richtungen · Regel im Geltungsbereich
ohne Gegenstand · offene Plan-Frage ohne Ausgang

**Wiederkehrendes Muster (drei der sieben HIGH):** F-1, F-2 und F-6 sind
dieselbe Klasse — ein Satz, der eine Messung *behauptet*, während die
zugehörige Menge nie ausgezählt wurde oder die Auszählung nicht mehr zum Satz
gehört. Das ist der Gegenstand, den der Slice selbst adressiert („Mit
`cell-max-chars` stellt sich die Frage nicht: Der Lauf zählt") — und der
Zeichenzähler zählt Zeichen, keine Behauptungen.

## Verdikt

**Merge-blockierend: ja.** Die Mechanik trägt: Konfiguration, Schwelle,
Untergrenze und alle vier Mutations-Proben sind gemessen und reproduzierbar,
und keine der neun gekürzten Zellen hat Substanz verloren, die nicht
anderswo steht. Blockierend ist die **Aussagenschicht** darum herum: zwei
Zahlen in der committeten Konfiguration sind falsch (F-1, F-2), die neue
Sensor-Datei beschreibt ihr Target unvollständig und in einem Punkt unrichtig
(F-3, F-4), drei Deklarationen des Moduls `structure` sind seit diesem Commit
überholt (F-5), eine ersetzte Zelle behauptet eine Deckung, die die Tabelle
nicht hat (F-6), und die Umsetzung steht quer zu §1 des eigenen Plans (F-7).
Ein Slice, der Gate-Doku gegen Überladung härtet, darf die Doku daneben nicht
lockern.

**Übergabe:** Findings gehen an die Implementer-Rolle; F-7 ist vor der
Fortschreibung von §3 zu entscheiden (Plan angleichen oder Umsetzung
zurückziehen), nicht danach zu berichten. Dieser Report ist Lauf-Beleg, keine
Verifikation — DoD-/Spec-Konformität prüft der Verifier separat (`v6.5.0` ·
`regelwerk/modul-11-verification.md`).
