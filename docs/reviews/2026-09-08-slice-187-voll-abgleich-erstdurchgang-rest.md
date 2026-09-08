# Review-Report: slice-187 — 2026-09-08

**Review-Art:** Plan + Diff — geprüft gegen den Slice-Plan, die Ziel-Formen des
vendorten Stands, das Beobachtungs-Register und die berührten Repo-Artefakte
(Modul 10 §Drei Review-Arten). Die Kern-Behauptungen dieses Slice sind eine
**reproduzierbare Messung** (§2), **acht Übernahmen** (§3.1–§3.3), eine
**Nicht-Erhöhung** von sechs Register-Einträgen (§7, §9) und eine
**Immutabilitäts-Grenze** (§3.5). Der Review rechnet die Messung nach, prüft
jede Übernahme gegen die Vorlage und den Ist-Stand, und hält die
Nicht-Erhöhung gegen das Register.

**Gegenstand:** Commits `bd2ba58` (mv `open`→`in-progress`), `b287f01`
(Neu-Zuschnitt), `4504e40` (`AGENTS.md` + `harness/conventions.md`), `248f69b`
(Planning-Index + §3), `0f1e477` (Immutabilitäts-Grenze), `a0d6c83` (Closure +
zwei Folge-Slices); Range `c937fb1..HEAD`, Working Tree sauber.

**Skill:** `.harness/skills/reviewer.md` @ Stand `a0d6c83` — einschließlich
§Mess-Regeln; beide Regeln sind auf die Messungen dieses Slice angewandt.
**Modell:** claude-opus-5 (1M) · **Datum:** 2026-09-08

**Eingangs-Kontext:**

- `slice-187` (Plan, `in-progress/`), `slice-188` und `slice-189` (`open/`),
  `slice-186` und `slice-185` (Volltext plus Review-Report aus ihren Archiven)
- `AGENTS.md` §1/§3.1–§3.7/§4/§5/§6 — Stand vor und nach dem Diff
- `harness/conventions.md` §Purpose, §Baseline, §Adaptions-Block, `MR-000`,
  §Aktive Adaptionen, §Modus-Deklaration · `harness/conventions/MR-020-…`
- `docs/plan/planning/README.md`, `docs/plan/planning/in-progress/roadmap.md`
- `.d-check.yml` (Module `vcs`, `commits`, `structure`), `Makefile`,
  `tools/slice-mv.sh`
- `v6.5.0` · `templates/AGENTS.template.md`,
  `templates/harness/conventions.template.md`,
  `templates/docs/plan/planning/README.template.md`,
  `templates/spec/architecture.template.md`,
  `templates/project-readme.template.md`
- `v6.5.0` · `regelwerk/modul-02-harness-bootstrap.md`,
  `regelwerk/modul-05-planning-harness.md`, `regelwerk/modul-06-roadmap.md`
- Beobachtungs-Register, gelesen: alle 61 Einträge (Zähler und `state.md`),
  im Detail `BEO-PLAN/erstdurchgang-als-einmal-slice-geschnitten`,
  `BEO-PLAN/messung-ohne-reproduzierbares-instrument`,
  `BEO-PLAN/vollstaendigkeits-haken-ohne-erschoepften-gegenstand`,
  `BEO-PLAN/zielsatz-nach-plan-aenderung-nicht-nachgezogen`,
  `BEO-PLAN/kandidaten-klassifikation-groeber-als-der-kandidat`,
  `BEO-PLAN/form-vergleich-sprachblind`,
  `BEO-PLAN/risiko-ausgang-fuer-gewollte-wirkung`,
  `BEO-PLAN/slice-provenienz-aus-gedaechtnis-statt-git-log`,
  `BEO-HARNESS/chronik-in-gelesenen-dateien`,
  `BEO-HARNESS/hard-rule-37-ohne-sensor`,
  `BEO-HARNESS/adaption-korrigiert-repo-aussage`,
  `BEO-HARNESS/baseline-regel-nie-erwogen-weil-bestand-sie-verletzt`,
  `BEO-HARNESS/aggregat-aufzaehlung-hinkt-dem-makefile-hinterher`,
  `BEO-GATE/zusage-weiter-als-ihre-durchsetzung` (neu),
  `BEO-GATE/probe-liefert-den-gegenstand-mit`
- früherer Report zum selben Bereich: der zu `slice-186` (aus dem Archiv);
  seine zwölf Findings sind hier auf Wiederholung geprüft

**Geltungsbereich der eigenen Messungen** (Skill §Mess-Regeln, erste Regel):

- **Das Instrument aus §2 ist ausgeführt**, unverändert aus dem Plan
  extrahiert (`python3`, Parameter wie deklariert: Fenster 6, Mindestlänge 6),
  über **alle 16** Vorlage/Gegenstück-Kombinationen des Repos — einmal gegen
  `HEAD`, einmal gegen einen `git archive`-Auszug von `c937fb1`. Die Zahlen
  unten sind damit reproduzierbar, unter **diesen** Parametern; ein anderes
  Fenster liefert andere Zahlen.
- Der Ziel-Form-Abgleich der drei Paare ist von mir **auf Satz-Ebene für die
  Kandidatenlisten** nachvollzogen (48 + 29 + 26 Zeilen einzeln gelesen) und
  **auf H2-Ebene** für den Rest der Dateien; er ersetzt kein Wort-für-Wort-Lesen
  der Gegenstücke.
- **Die Gegenrichtung habe ich nur punktuell geprüft** — also: trägt das
  Gegenstück Aussagen, die *falsch* sind, ohne dass die Vorlage etwas dazu sagt?
  Das Instrument sieht diese Richtung prinzipiell nicht (F-5 ist ein Treffer
  aus dieser Stichprobe, kein Vollstand).
- **`make gates` / `make verify` habe ich nicht ausgeführt** (Auftrag: Belege
  liegen vor). Alle Aussagen unten stützen sich auf Dateien, `git` und den
  einen Python-Lauf, nicht auf Gate-Läufe.

---

## Findings

### F-1 — Der von `slice-185` abgetretene offene Punkt zu `spec/architecture.md` hat keine annehmende Adresse mehr

- `kategorie`: HIGH
- `quelle`: `v6.5.0` · `regelwerk/modul-05-planning-harness.md` §Ziel-Form:
  Slice (*„Die Adresse muss die Sendung annehmen: Ein Folge-Slice, der den
  verwiesenen Punkt selbst ausschließt … ist keine"*); `v6.5.0` ·
  `templates/spec/architecture.template.md` (*„keine Historie: `Letzte
  Änderung` oben ist ein Frische-Marker, kein Protokoll"*)
- `pfad`: `docs/plan/planning/in-progress/slice-187-voll-abgleich-erstdurchgang-rest.md`
  §2 (Tabelle), §8 (*„Vier Paare bleiben"*), DoD-Punkt 3;
  `docs/plan/planning/open/slice-189-voll-abgleich-spec-straten.md` §1, vierter
  Ausschlusspunkt; `spec/architecture.md`:181–187
- `befund`: `slice-185` schloss das Paar `spec/architecture.md` **nicht**
  vollständig: Die dritte Klausel der Ziel-Form (keine Historie) ist als
  *verletzt* ausgewiesen — *„Die dritte ist ein offener Punkt und geht an
  `slice-186`"*. `slice-186` erwähnt `spec/architecture.md` in seinem gesamten
  Plan **nullmal** (`grep -ci`). Dieser Slice führt das Paar weder in der
  Tabelle in §2 noch in §1 noch in der Restmenge, und
  `slice-189` schließt es aus mit der Begründung *„in slice-185 bereits
  abgeglichen"*. Der Punkt existiert weiter: `spec/architecture.md` §8
  *Historie* trägt vier Versionszeilen. Die Zusage *„keine offene Aufgabe,
  sondern eine Restmenge mit Zähler"* (§8) deckt ihn damit nicht ab — er ist
  weder in den 123 Kandidaten enthalten noch einem Folge-Slice zugewiesen.
- `verifizierbar`: ja — `grep -c architecture` auf den `slice-186`-Volltext
  (0), `grep -n "^## " spec/architecture.md` (§8 Historie vorhanden), und die
  Ausschlusszeile in `slice-189` §1.
- `klasse`: Adresse nimmt die Sendung nicht an

### F-2 — Der Sichtungs-Schritt sichtet eine andere Sub-Area als die, die §9 selbst als berührt führt

- `kategorie`: HIGH
- `quelle`: `v6.5.0` · `regelwerk/modul-05-planning-harness.md` §Zwei Schritte
  vor der Modus-Begründung (*„Steht eine der berührten Sub-Areas dort? Dann
  gehört der Zähler-Stand ins Kriterium Evidenz-/Diskrepanz-Risiko"*)
- `pfad`: `docs/plan/planning/in-progress/slice-187-voll-abgleich-erstdurchgang-rest.md`
  §9, Block *(1) Register* (Zeilen 481–495) gegen §9, Block *Sub-Area-Wahl*
  (Zeilen 458–471)
- `befund`: §9 stellt fest, berührt sei **`HARNESS`** — ausdrücklich auch für
  `docs/plan/planning/README.md`, dessen Pfad in `PLAN` liegt. Die anschließende
  Sichtung nennt **sechs** Einträge, und alle sechs liegen unter `BEO-PLAN`;
  aus `BEO-HARNESS` ist keiner genannt, aus `BEO-GATE` ebenfalls keiner —
  obwohl §3.5 einen neuen `BEO-GATE`-Eintrag anlegt. Ungesichtet bleiben fünf
  offene `BEO-HARNESS`-Einträge bei 2×, davon drei unmittelbar einschlägig:
  `chronik-in-gelesenen-dateien` (dessen `state.md` `AGENTS.md` namentlich als
  Reststelle führt — siehe F-3), `adaption-korrigiert-repo-aussage` (dessen
  `state.md` als Auflösungs-Trigger *„die Überarbeitung der
  ID-Schema-Deklaration"* nennt, also genau das, was §3.2 anfasst — siehe F-4)
  und `baseline-regel-nie-erwogen-weil-bestand-sie-verletzt`, dessen
  `evidence/slice-177.md` **denselben** `AGENTS.md`-§4-Satz mit Messung führt,
  den §3.1 als Befund 1 einführt. Die Negativbefund-Zeile am Ende von §9 nennt
  `SPEC` — eine Sub-Area, die der Slice selbst als nicht berührt ausweist.
- `verifizierbar`: nein — welche Einträge eine Sub-Area „betreffen", ist ein
  Urteil. Die zwei Einzelfakten sind es (die sechs genannten Kennungen liegen
  alle unter `BEO-PLAN`; `ls docs/plan/planning/observations/BEO-HARNESS/`
  liefert 15 Einträge).
- `klasse`: Sichtungs-Schritt liest eine andere Sub-Area als die deklariert berührte

### F-3 — `AGENTS.md` §3.3 trägt jetzt Chronik über abwesenden Text

- `kategorie`: HIGH
- `quelle`: `AGENTS.md` §3.7 (*„Falsch: abwesenden Text beschreiben … Richtig:
  die geltende Zusage nennen; die vorige hält `git`"*); Reviewer-Skill
  §Klassifikation, HIGH-Kategorie *Kommentar trägt keine der
  Kommentar-Klassen*; `BEO-HARNESS/chronik-in-gelesenen-dateien` (2×)
- `pfad`: `AGENTS.md`:118–120
- `befund`: Der Abgleich hängt an die neu übernommene Regel drei Sätze an:
  *„Bis slice-187 stand hier nur der Regelfall — als Reihenfolge, nicht als
  Wahl. … und das Briefing sagte das Gegenteil."* Das ist die Form, die
  `evidence/slice-182.md` desselben Register-Eintrags als Beispiel führt
  (*„(bis slice-079 tat das `gate-consistency`)"* — dort gestrichen); der
  Eintrag steht bei 2× und sein `state.md` sagt zusätzlich *„`AGENTS.md` §5
  trägt weiterhin eine gemessene Reststelle"*. Mit diesem Absatz sind es zwei
  Stellen in derselben Datei (`AGENTS.md`:118 und :289), und die neue ist in
  diesem Diff entstanden. §7 und §9 behaupten, kein Register-Eintrag erreiche
  mit diesem Slice 3× — für diesen Eintrag ist die Frage nicht gestellt worden
  (F-2). Die Aussage der drei Sätze steht vollständig in `git` (`git log -S`
  auf den §3.3-Text) und in der Closure-Notiz des Slice.
- `verifizierbar`: nein — kein Gate fängt §3.7; der Eintrag sagt das selbst
  (*„ein Zähler über Kommentar-Klassen wäre Schein-Genauigkeit"*). Die
  Fundstelle ist es: `grep -n "Bis slice-" AGENTS.md` liefert zwei Treffer,
  einer davon aus `4504e40`.
- `klasse`: Chronik in gelesenen Dateien

### F-4 — Die Grenze über `MR-000` umgeht das Werkzeug, das §Disziplin für Korrekturen selbst nennt

- `kategorie`: HIGH
- `quelle`: `harness/conventions.md` §Adaptions-Block, §Disziplin
  (*„Korrekturen entstehen als neuer `MR` oder als ausdrückliche Aufhebung"*);
  `harness/conventions/MR-020-adr-vorlage-generisch.md` (Geltungsbereich:
  `MR-000` §ID-Schema-Deklaration); Reviewer-Skill §Klassifikation,
  *nachweislich falsche Tatsachenbehauptung*
- `pfad`: `harness/conventions.md`:112–124 (neuer Absatz);
  `docs/plan/planning/in-progress/slice-187-voll-abgleich-erstdurchgang-rest.md`
  §3.2, Absatz *„Zwei Ziel-Form-Punkte werden bewusst nicht übernommen"*, und
  §7, vierter Risiko-Ausgang
- `befund`: Der Plan begründet den freistehenden Absatz mit: *„Beides ließe
  sich nur durch eine **inhaltliche Änderung an einem akzeptierten Eintrag**
  beheben, und die verbietet §Disziplin."* Die Datei selbst widerlegt das
  *nur*: §Disziplin nennt zwei Instrumente für eine Korrektur — neuer `MR` oder
  ausdrückliche Aufhebung —, und `MR-020` ist der lebende Präzedenzfall, ein
  Eintrag, dessen Geltungsbereich wörtlich *„`MR-000` §ID-Schema-Deklaration,
  Zeile zu `ADR-NNNN`"* lautet und der genau denselben Defekt (eine veraltete
  Zeile im ID-Schema) so behoben hat. Der zweite der beiden Punkte — die
  Beobachtungs-Kennung fehlt im ID-Schema — ist damit derselbe Fall wie
  `MR-020`. Gewählt ist stattdessen ein drittes, in keiner Quelle vorgesehenes
  Werkzeug: ein Absatz oberhalb des Eintrags. Der Risiko-Ausgang in §7 stützt
  sich auf dieselbe Prämisse (*„Damit kann das Risiko für diesen Slice nicht
  mehr eintreten"*) und trägt darum ebenfalls nicht: das Risiko ist
  eingetreten und unaufgelöst, nicht *entfallen*.
- `verifizierbar`: ja — die §Disziplin-Zeile, `MR-020`s Geltungsbereich-Feld
  und die Zeile für `MR-020` in §Aktive Adaptionen nebeneinander.
- `klasse`: Korrektur am Bestand ohne das dafür deklarierte Werkzeug

### F-5 — Das als „ohne Befund auf Satz-Ebene" geschlossene Paar trägt eine gegen das Repo widerlegte Aussage

- `kategorie`: HIGH
- `quelle`: `v6.5.0` · `regelwerk/modul-06-roadmap.md` §Wellen-Closure-Prozedur,
  Schritt 4; `AGENTS.md` §6 (Archivierung als Pflichtschritt der
  Slice-Closure); `BEO-PLAN/vollstaendigkeits-haken-ohne-erschoepften-gegenstand`
  (2×)
- `pfad`: `docs/plan/planning/README.md`:123–129; Ausgang in
  `docs/plan/planning/in-progress/slice-187-voll-abgleich-erstdurchgang-rest.md`:253
- `befund`: Die Datei sagt, die Baseline führe seit `v6.0.0` einen sechsten
  Wellen-Closure-Schritt (Zeitdokumente archivieren), und: *„a-check hat ihn
  nicht adoptiert, mangels Trigger: keine offene Welle, kein Bedarf, alten
  Bestand loszuwerden. Diese fünf Schritte bleiben darum a-checks vollständige,
  gelebte Prozedur."* Das Repo widerlegt beide Hälften: `done/welle-01/` bis
  `done/welle-15/` tragen je ein `archiv.zip`, und der Commit `b94070c` heißt
  *„welle-15 archivieren (Closure-Schritt 4)"*. `AGENTS.md` §4 und §6 sowie
  `harness/README.md` §Nicht-Gates führen `make archive-wave` als Pflichtschritt
  seit `slice-157`. Der Satz stammt aus `e7b6f16` (`slice-139`) und ist seither
  unverändert mitgereist. Der Slice schließt den Rest dieses Paares als *ohne
  Befund auf Satz-Ebene* — die Angabe der Ebene ist der DoD-Punkt, der aus
  genau dieser Beobachtungsklasse stammt. **Einschränkung, die zum Befund
  gehört:** Das Instrument aus §2 läuft nur in der Richtung Vorlage→Repo und
  hätte diese Stelle nie als Kandidat gemeldet; die Ziel-Form führt den
  Abschnitt nicht. Der Ausgangssatz sagt das nicht, sondern behauptet die
  Satz-Ebene für das Paar.
- `verifizierbar`: ja — `ls docs/plan/planning/done/welle-*/archiv.zip` und
  `git log --diff-filter=A -- 'docs/plan/planning/done/welle-*/archiv.zip'`
  gegen den Satz.
- `klasse`: Vollständigkeits-Haken gesetzt, ohne den Gegenstand zu erschöpfen

### F-6 — Die Tabelle „über alle 15 Ziel-Form-Paare" führt 14 Zeilen, und eine davon ist das falsche Gegenstück

- `kategorie`: MEDIUM
- `quelle`: `harness/conventions.md` §Baseline (*„die **15** Vorlagen mit genau
  einem Gegenstück"*); der Slice selbst, §1 (*„slice-185: fünf, slice-186:
  drei"*)
- `pfad`: `docs/plan/planning/in-progress/slice-187-voll-abgleich-erstdurchgang-rest.md`:161–179
- `befund`: Die Überschrift sagt *„Ergebnis über alle 15 Ziel-Form-Paare — die
  acht bereits abgeglichenen sind mitgemessen"*; die Tabelle trägt **14**
  Zeilen, davon **sieben** aus dem Bestand der acht. Es fehlt
  `spec/architecture.md` (von `slice-185` abgeglichen), und statt `README.md`
  — dem Paar, das `slice-186` §3.2 mit 5 Kandidaten führt — steht
  `README.de.md`. Mit dem Instrument des Slice nachgemessen: `spec/architecture.md`
  **20**, `README.md` **10**, `README.de.md` **7**. Auf die drei tragenden
  Summen wirkt sich das nicht aus (247 = die sieben verbleibenden Paare, 124 =
  dieser Slice, 123 = Restmenge — alle drei reproduzieren exakt); die
  Bestandsaussage der Tabelle ist trotzdem nicht die, die ihre Überschrift
  behauptet, und die Auslassung ist dieselbe Datei, um die es in F-1 geht.
- `verifizierbar`: ja — das Instrument aus §2 über alle 16 Paar-Kombinationen
  laufen lassen und die Zeilen zählen.
- `klasse`: Vollständigkeits-Haken gesetzt, ohne den Gegenstand zu erschöpfen

### F-7 — Die Probe zu `make doc-immutable` hat ihren Gegenstand nicht berührt

- `kategorie`: MEDIUM
- `quelle`: Reviewer-Skill §Mess-Regeln, *Eine Mutations-Probe belegt erst,
  wenn sie rot war* (`seit slice-181`, verkörpert in `AGENTS.md` §5);
  `BEO-GATE/probe-liefert-den-gegenstand-mit` (3×, verkörpert)
- `pfad`: `docs/plan/planning/in-progress/slice-187-voll-abgleich-erstdurchgang-rest.md`:278–283;
  `docs/plan/planning/observations/BEO-GATE/zusage-weiter-als-ihre-durchsetzung/evidence/slice-187.md`
- `befund`: §3.5 stellt die Frage *„was passiert, wenn ich `MR-000` trotzdem
  ändere?"* und antwortet: *„und **maß** statt anzunehmen: `make doc-immutable`
  über die Commit-Range bleibt **grün**"*. Die Evidence-Datei ergänzt: *„Die
  Änderung ist trotzdem unterblieben."* Im gesamten Range ist der
  `MR-000`-Block byte-identisch (1239 Zeichen vorher wie nachher) — der Lauf
  hatte also keine Mutation des Gegenstands vor sich, und ein grüner Lauf ohne
  Mutation belegt die Nicht-Reichweite nicht. Die Meldung, die die Mess-Regel
  verlangt (*„war sie rot, und woran?"* bzw. hier: was genau war mutiert), ist
  nicht genannt. Der **tragende** Beleg steht daneben und ist gültig: die
  `paths`-Liste des Moduls `vcs` in `.d-check.yml` führt ausschließlich
  `docs/plan/adr/[0-9]*.md`. Der Befund betrifft nur die als Messung
  ausgegebene Hälfte.
- `verifizierbar`: ja — `git diff c937fb1..HEAD` auf den `MR-000`-Abschnitt
  liefert null Änderungen; `grep -A2 '^vcs:' .d-check.yml` liefert den
  tragenden Beleg.
- `klasse`: Probe trifft ihren Gegenstand nicht

### F-8 — DoD-Punkt 2 verlangt die Prüf-Ebene im Ausgang; zwei der drei Paare tragen sie nicht

- `kategorie`: MEDIUM
- `quelle`: `BEO-PLAN/vollstaendigkeits-haken-ohne-erschoepften-gegenstand`
  (2×; der Slice zitiert ihn selbst als Quelle des DoD-Punkts)
- `pfad`: `docs/plan/planning/in-progress/slice-187-voll-abgleich-erstdurchgang-rest.md`:308–311
  (DoD, `[x]`) gegen §3.1 (Tabelle, Zeilen 201–208) und §3.2 (Tabelle, Zeilen
  227–231)
- `befund`: §3 kündigt an: *„Je Befund steht unten, woran er hängt und **auf
  welcher Ebene** er geprüft wurde — die Ebene gehört in den Ausgang."* Die
  Tabellen in §3.1 und §3.2 führen die Spalten *Ziel-Form-Stelle · Was a-check
  hatte · Übernommen*; eine Ebenen-Angabe steht in keiner Zelle und in keinem
  der begleitenden Absätze. Genannt ist die Ebene **einmal** — in §3.3, für den
  *ohne Befund*-Ausgang des dritten Paares (*„auf Satz-Ebene"*). Der DoD-Punkt
  ist trotzdem abgehakt. Er ist damit auf derselben Ebene gesetzt, gegen die er
  sich richtet.
- `verifizierbar`: nein — ob ein Absatz eine Prüf-Ebene nennt, ist eine Lesart;
  die Spaltenüberschriften sind es.
- `klasse`: Vollständigkeits-Haken gesetzt, ohne den Gegenstand zu erschöpfen

### F-9 — Zwei Provenienz-Aussagen sind aus dem Gedächtnis geschrieben; eine ist gegen `git` falsch

- `kategorie`: MEDIUM
- `quelle`: `BEO-PLAN/slice-provenienz-aus-gedaechtnis-statt-git-log` (1×);
  Reviewer-Skill §Mess-Regeln, *Geltungsbereich einer Messung*
- `pfad`: `docs/plan/planning/in-progress/slice-187-voll-abgleich-erstdurchgang-rest.md`:232–234
  und :212–214
- `befund`: (a) *„a-check führt **seit jeher** eine Datei je Eintrag mit
  `conventions/done/` als zweitem Ort."* Der erste Commit, der eine Datei unter
  `harness/conventions/` anlegt, ist `647f230` (`slice-096`, 2026-08-29,
  *„Adaptions-Speicher in die Verzeichnis-Form"*) — 70 Tage nach der Adoption.
  Davor lebten die Einträge als `###`-Abschnitte in `harness/conventions.md`;
  die Datei sagt das selbst (*„Wo ein Eintrag vor slice-096 unter seinem
  Überschriften-Slug veröffentlicht wurde"*). Die Aussage trägt den
  Closure-Lerneintrag mit (*„die Praxis war richtig und der Satz falsch oder
  abwesend"*) und ist für diesen Eintrag nicht haltbar: bis `slice-096` war auch
  die Praxis eine andere. (b) *„und das Repo fährt seit jeher den zweiten
  [Fall]"* (§3.1, Befund 2): Nachgeprüft trifft das für `slice-183`–`slice-186`
  zu (Closure-Commit jeweils **vor** dem `… -> done (make slice-mv)`-Commit);
  für „seit jeher" ist kein Geltungsbereich genannt, und `make slice-mv`
  existiert erst seit `slice-118`.
- `verifizierbar`: ja — `git log --diff-filter=A -- harness/conventions/` für
  (a); `git log --oneline` um die vier Closures für (b).
- `klasse`: Slice-Provenienz aus Gedächtnis statt aus `git log` behauptet

### F-10 — Der Geltungsbereich des Instruments ist nur in eine Richtung genannt; die teurere Grenze fehlt

- `kategorie`: MEDIUM
- `quelle`: Reviewer-Skill §Mess-Regeln, *Geltungsbereich einer Messung*
  (`BEO-PLAN/review-geltungsbereich-zu-eng`, verkörpert)
- `pfad`: `docs/plan/planning/in-progress/slice-187-voll-abgleich-erstdurchgang-rest.md`:96–99;
  `AGENTS.md`:18–20 gegen `v6.5.0` · `templates/AGENTS.template.md`:30–33
- `befund`: §2 nennt **eine** Grenze: *„Es findet wörtliche Übernahme. Eine
  sinngemäße Doppelung in anderen Worten sieht es nicht"* — also die Richtung,
  die zu **zu vielen** Kandidaten führt. Die gegenläufige, teurere ist nicht
  genannt: Eine Vorlagenzeile gilt als gedeckt, sobald **ein** 6-Wort-Fenster im
  Gegenstück vorkommt (`treffer = any(...)`), also auch dann, wenn die Zeile nur
  zur Hälfte übernommen wurde. Solche Zeilen werden nie Kandidat und sind für
  den Abgleich unsichtbar. Konkreter Treffer aus meiner Stichprobe: Die
  Ziel-Form führt *„Strukturregeln (ID-Schemata, Verzeichniskonvention,
  Adaptionen ggü. Baseline, Modus-Deklarationen pro Sub-Area, **Zusatzklassen
  für Sensors-Bindung**) leben in `harness/conventions.md`"*; `AGENTS.md` §1
  führt vier der fünf Punkte, `harness/README.md` und
  `harness/conventions.md` §Purpose führen den fünften. Die Zeile ist kein
  Kandidat, weil ihre ersten sechs Wörter treffen.
- `verifizierbar`: ja — `grep -n -A4 "^Strukturregeln" AGENTS.md
  .harness/baseline/v6.5.0/templates/AGENTS.template.md harness/README.md`.
- `klasse`: Geltungsbereich einer Messung nur in eine Richtung genannt

### F-11 — Die benannte Reibung am Risiko-Ausgang wird weder gebucht noch als eigene Kennung angelegt

- `kategorie`: MEDIUM
- `quelle`: `v6.5.0` · `regelwerk/modul-06-roadmap.md` §Das
  Beobachtungs-Register (*„Ein Vorkommen **ohne** abgeschlossenen Vorgang
  bekommt keinen Beleg … es gehört trotzdem in den Eintrag — benannt, nicht
  gezählt"*; *„Der Pfad ersetzt die Namens-Disziplin"*)
- `pfad`: `docs/plan/planning/in-progress/slice-187-voll-abgleich-erstdurchgang-rest.md`:385–392;
  `docs/plan/planning/observations/BEO-PLAN/risiko-ausgang-fuer-gewollte-wirkung/observation.md`
- `befund`: §7 stellt fest, die geschlossene Dreier-Menge habe keine Kategorie
  für *„eingetreten und im Slice selbst aufgelöst"*, grenzt das gegen
  `risiko-ausgang-fuer-gewollte-wirkung` ab und schließt: *„deshalb **kein
  Beleg** dort"*. Die Abgrenzung trägt am Titel des Eintrags (dort:
  *gewollte Wirkung*), nicht an seinem Rumpf, der die Klasse über das Symptom
  definiert — *„musste beim Abschluss zu `gestrichen mit Begründung`
  umformuliert werden"*, und genau das ist hier geschehen. Ist die Abgrenzung
  richtig, ist es eine **neue** Beobachtung mit einem abgeschlossenen Vorgang
  dahinter, und die Register-Regel sieht dafür eine eigene Kennung vor; ist sie
  falsch, gehört der Beleg in den vorhandenen Eintrag. Gewählt ist keine der
  beiden Möglichkeiten: Die Beobachtung lebt nur im Plan, der mit dem Slice
  archiviert wird, und ist beim nächsten Vorkommen nicht auffindbar. Dass
  derselbe Slice für eine andere Klasse (§3.5) eine neue Kennung anlegt, macht
  die Ungleichbehandlung sichtbar.
- `verifizierbar`: nein — ob zwei Beobachtungen dieselbe Klasse sind, ist das
  Urteil, das die Register-Regel ausdrücklich dem Menschen zuweist.
- `klasse`: Reibung benannt, aber weder gebucht noch als Kennung angelegt

### F-12 — Drei Sammelaussagen sind gröber als die Menge, über die sie sprechen

- `kategorie`: LOW
- `quelle`: `BEO-PLAN/kandidaten-klassifikation-groeber-als-der-kandidat` (2×)
- `pfad`: `docs/plan/planning/in-progress/slice-187-voll-abgleich-erstdurchgang-rest.md`:359,
  :444–446, :253–258
- `befund`: (a) §7 sagt *„**Vier** Register-Einträge stehen bei 2× und treffen
  genau diese Arbeitsform"*, §9 führt **sechs**, alle bei 2×; die
  Closure-Notiz schreibt *„Nicht erhöht wurden die vier Einträge aus §9"* — §9
  hat keine vier. (b) §3.3 sagt, das Beobachtungs-Register sei in der Ziel-Form
  *„nur [mit] die[r] Existenz"* genannt; die Vorlage führt zusätzlich
  Steering-Loop-Zähler, Fortschreibung bei jeder Slice-Closure und *„überlebt
  jede Welle"* — a-checks Zusatz ist Verzeichnisform und abgeleiteter Zähler,
  nicht die Wellen-Unabhängigkeit. (c) §3.3 sagt, a-check nenne bei §Aktueller
  Stand *„nicht nur ‚nicht als Snapshot eintragen', sondern **warum**"*; die
  Vorlage nennt den Grund ebenfalls (*„sonst driftet die Tabelle"*), nur kürzer.
- `verifizierbar`: ja — Zeilen zählen bzw. die drei Vorlagenstellen lesen.
- `klasse`: Kandidaten-Klassifikation gröber als der Kandidat

### F-13 — Zwei Stellen schreiben `make slice-mv` eine Leistung zu, die das Werkzeug ausdrücklich nicht erbringt

- `kategorie`: LOW
- `quelle`: `tools/slice-mv.sh` §NICHT BEHANDELT (1) *SEMANTIK* (*„Das Werkzeug
  zieht PFADE nach, keine Aussagen"*); `AGENTS.md` §4, Zeile `make slice-mv`
- `pfad`: `AGENTS.md`:119; Commit `bd2ba58` (Betreff und Inhalt)
- `befund`: (a) `AGENTS.md` §3.3 sagt zum zweiten Fall (*erst der Inhalt, dann
  der `git mv`*): *„`make slice-mv` fährt genau ihn"*. Das Werkzeug führt nur
  den Move aus; die Reihenfolge ist eine Entscheidung des Laufs, keine des
  Werkzeugs. (b) Der Commit `bd2ba58` trägt den Betreff *„(make slice-mv)"* und
  enthält neben dem reinen Move die Streichung des Ruhe-Markers *„Nichts in
  Arbeit"* in `docs/plan/planning/in-progress/roadmap.md` — ein Zustandssatz,
  von dem das Skript in seinem Kopfkommentar ausdrücklich sagt, dass es ihn
  nicht anfasst. Die Streichung selbst ist richtig (`in-progress/` trägt seit
  dem Move einen Slice); zugeschrieben ist sie dem Werkzeug.
- `verifizierbar`: ja — `git show bd2ba58` gegen `sed -n '22,27p' tools/slice-mv.sh`.
- `klasse`: Werkzeug-Zuschreibung für eine Leistung, die das Werkzeug nicht erbringt

### F-14 — Die Hälfte der „benannten Grenze" steht bereits 50 Zeilen tiefer in derselben Datei

- `kategorie`: LOW
- `quelle`: `AGENTS.md` §1 / `v6.5.0` · `regelwerk/modul-09-implementierung.md`
  §Ziel-Form: AGENTS.md (Pointer statt Duplikat); der neue Absatz selbst
  (*„Was hier steht, liest **jeder** Agentenlauf"*)
- `pfad`: `harness/conventions.md`:112–124 gegen `harness/conventions.md`:161–164
- `befund`: Der neue Absatz führt als ersten der zwei Punkte: *„Das Pflichtfeld
  *Ersetzt-Baseline-Regel* fehlt, weil es nach seiner Annahme entstand."*
  §Aktive Adaptionen sagt dasselbe seit Längerem: *„Die Spalte
  *Ersetzt-Baseline-Regel* ist das Pflichtfeld des neuen Stands. Sie kann in
  einen akzeptierten Eintrag **nicht nachgetragen** werden … sie entsteht in den
  Nachfolge-Einträgen."* Zwei Stellen derselben Datei tragen jetzt dieselbe
  Aussage; ein Zeiger hätte gereicht. Daneben stehen an :110–111 zwei
  aufeinanderfolgende Leerzeilen aus demselben Commit.
- `verifizierbar`: ja — beide Absätze nebeneinander lesen.
- `klasse`: dieselbe Zusage zweimal in derselben Datei

### F-15 — Erstauftreten für einen abgeschlossenen Vorgang nachgetragen, diesmal ohne Selbstauskunft

- `kategorie`: INFO
- `quelle`: `v6.5.0` · `regelwerk/modul-06-roadmap.md` §Das
  Beobachtungs-Register; Wiederholung der Klasse aus dem `slice-186`-Report
  (dort F-12)
- `pfad`: `docs/plan/planning/observations/BEO-GATE/zusage-weiter-als-ihre-durchsetzung/evidence/slice-186.md`
- `befund`: Der Beleg für `slice-186` ist bei der Closure von `slice-187`
  entstanden (Commit `a0d6c83`, gemeinsam mit `observation.md`, `state.md` und
  `evidence/slice-187.md`). Beide urteilsfreien Prüfungen bestehen — der
  Dateiname ist die Kennung eines abgeschlossenen Vorgangs, `slice-186` liegt in
  `done/` —, und der abgeleitete Zähler stimmt. Anders als im ersten Fall sagt
  die Datei nicht, dass sie nachgetragen ist; im `slice-186`-Report war genau
  das der Grund, warum der Fall folgenlos blieb. Kein Handlungsbedarf; genannt,
  weil es der zweite Fall ist und die Selbstauskunft diesmal fehlt.
- `verifizierbar`: ja — `git log --diff-filter=A` auf die Datei gegen die Lage
  von `slice-186`.
- `klasse`: Beleg für einen abgeschlossenen Vorgang nachgetragen

## Negativbefunde

- geprüft, ohne Befund: **Die Ausgangsmessung reproduziert exakt.** Das
  Instrument aus §2 ist unverändert aus dem Plan extrahiert und ausgeführt
  worden. Gegen `HEAD`: `AGENTS.md` **48**, `harness/conventions.md` **29**,
  `docs/plan/planning/README.md` **26** — die Zahlen aus §3.4 und der
  Closure-Notiz. Gegen `c937fb1`: **60 / 38 / 26** — die Zahlen aus §1 und §2.
  Alle elf übrigen Zeilen der Tabelle reproduzieren ebenfalls zeichengenau
  (42, 28, 28, 25, 40, 23, 9, 7, 4, 4, 4). Die abgeleiteten Summen stimmen:
  247 über die sieben verbleibenden Paare, 124 für diesen Slice, 123 für die
  Restmenge. **Das ist der erste Slice dieser Kette, dessen Zahlen ein anderer
  Lauf nachrechnen kann** — die Antwort auf
  `BEO-PLAN/messung-ohne-reproduzierbares-instrument` trägt.
- geprüft, ohne Befund: **Die deklarierten Parameter tun, was der Text sagt.**
  Fenster 6 und Mindestlänge 6 stehen als Konstanten im Code; die
  YAML-Sonderbehandlung ist da und richtig herum (`pfad.endswith(('.yml',
  '.yaml'))` → nur Kommentarzeilen, `#` gestrippt statt übersprungen), und im
  Gegenstück wird `#` durch die Markup-Klasse in `norm()` mitentfernt.
  HTML-Kommentare werden in beiden Funktionen entfernt, Links auf ihren Text
  reduziert, Platzhalter `<…>` gestrichen, Überschriften/Blockquotes/
  Tabellen-Trennzeilen übersprungen. **Eine Nebenwirkung ohne Folge für die
  Zahlen:** Das Entfernen mehrzeiliger HTML-Kommentare in `zeilen()` geschieht
  vor dem Zeilen-Split, wodurch die ausgegebenen Zeilennummern hinter einem
  solchen Block verschoben sind; auf die Kandidatenmenge wirkt sich das nicht
  aus.
- geprüft, ohne Befund: **`MR-000` ist unangetastet.** Der Abschnitt von
  `### MR-000` bis `### Aktive Adaptionen` ist im Range byte-identisch (1239
  Zeichen vorher wie nachher). Die *Trennung* Kommentar-über-Eintrag hält
  formal; was daran trotzdem ein Befund ist, steht in F-4 und betrifft die
  Werkzeug-Wahl, nicht die Immutabilität.
- geprüft, ohne Befund: **Der Befund in §3.5 stimmt.** `.d-check.yml`, Modul
  `vcs`, führt `paths: ["docs/plan/adr/[0-9]*.md"]`; `harness/conventions.md`
  steht dort nicht, und `make doc-immutable` kann die Datei damit nicht
  erreichen. Auch die zweite Hälfte trägt: ein `MR`-Eintrag hat kein
  `Status:`-Feld, an dem `immutable-when` greifen könnte. (Die als *Messung*
  ausgegebene Probe daneben: F-7.)
- geprüft, ohne Befund: **Die Übernahmen 1, 3, 4 und 5 sind echte Lücken und
  korrekt übersetzt.** Der §4-Arbeitsteilungssatz stand in `AGENTS.md` nicht
  (a-check ersetzt `LH-*` durch `AC-*` — richtig für das hiesige ID-Schema);
  die drei Anlässe des breiteren Pflicht-Blicks und die Stichprobe *„auch bei
  aktuellem Pin"* stehen wörtlich in `v6.5.0` ·
  `regelwerk/modul-02-harness-bootstrap.md` §Freshness-Audit; die Zwei-Rollen-
  Aussage zu `templates/` fehlte; und `SPEC-<NNN>` ist in den `id-patterns` des
  Moduls `commits` tatsächlich nicht geführt.
- geprüft, ohne Befund: **Der neu übernommene Satz „definiert wird hier nichts"
  widerspricht der Tabelle nicht.** `evidence/slice-177.md` hatte für
  `AGENTS.md` §4 **16 von 40** Zellen über 250 Zeichen gemessen (Spitze 1871).
  Nachgemessen im aktuellen Stand: **41** Zeilen, längste Zweck-Zelle **148**
  Zeichen, **null** über 250 — die Zellengrenze aus `slice-185` hat den Bestand
  eingeholt, und der Satz ist im Moment seiner Übernahme wahr.
- geprüft, ohne Befund: **Übernahme 2 ist gegen die Praxis belegt.** Für
  `slice-183`, `slice-184`, `slice-185` und `slice-186` liegt der
  Closure-Commit jeweils **vor** dem `… -> done (make slice-mv)`-Commit; die
  Behauptung, das Briefing habe das Gegenteil der geübten Praxis gesagt, hält
  für den geprüften Ausschnitt. (Zur unbelegten Ausdehnung *„seit jeher"*: F-9.)
- geprüft, ohne Befund: **Der Neu-Zuschnitt ist eine Plan-Änderung, keine
  Zähler-Umgehung.** Der Zuschnitt ist vor der Arbeit korrigiert (`b287f01`,
  vor den drei Arbeits-Commits), die Restmenge ist gemessen und mit Kennungen
  übergeben, und Titel und §1 tragen den korrigierten Umfang — das ist genau
  der Ausgang, den `state.md` von
  `BEO-PLAN/erstdurchgang-als-einmal-slice-geschnitten` als Lehre nennt
  (*„eine Restmenge mit benanntem Umfang statt einer offenen Aufgabe"*), und
  `BEO-PLAN/zielsatz-nach-plan-aenderung-nicht-nachgezogen` ist damit
  vorweggenommen. **Die Einschränkung ist F-1 und F-6:** Die Restmenge, auf die
  sich das stützt, ist nicht vollständig — der abgetretene Punkt zu
  `spec/architecture.md` steht in keiner der Zahlen. Als Zähler-Umgehung lese
  ich das nicht: Die Auslassung ist ein Fehler in der Bestandsaufnahme, kein
  Verzicht auf eine fällige Buchung.
- geprüft, ohne Befund: **Risiko-Ausgänge — Form.** Vier notierte Risiken, vier
  Ausgänge, alle aus der geschlossenen Dreier-Menge, jeder mit Begründung
  (`modul-05` §Offene Risiken werden bei Closure aufgelöst). Die Ausgänge 1, 2
  und 3 tragen auch inhaltlich: Die Wette auf die Artefaktklasse ist
  aufgegangen (ein Durchgang, keine Rückführung), die Vorauswahl trägt ihre
  Warnung jetzt im Text, und die Nicht-Erhöhung der vier genannten Einträge ist
  für **diese vier** richtig begründet. Ausgang 4 trägt nicht — F-4.
- geprüft, ohne Befund: **Register-Form des neuen Eintrags.** Drei Dateien mit
  den drei Lebensdauern; `observation.md` definiert die Klasse und grenzt sie
  gegen `BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf` ab; `state.md` nennt
  Stand und Ausgangs-Kandidat ohne Chronik; zwei Evidence-Dateien, Zähler
  abgeleitet, kein gespeichertes Zähler-Feld. Die Sub-Area *Gate-/Werkzeug-
  Schicht* ist in `harness/conventions.md` §Modus-Deklaration geführt. Der
  `slice-186`-Beleg trägt: der `slice-186`-Report führt den Fall als F-7 mit
  eigener Klasse. (Zum Nachtrag: F-15.)
- geprüft, ohne Befund: **Closure-Notiz, semantisch** (Maßstab
  `.harness/skills/closure-note-reviewer.md`). Alle drei Inhalte sind da:
  Lernsignal mit Ursache (*„Wo ein Repo eine Regel übt, ohne sie zu schreiben,
  sagt sein Briefing im Zweifel das Gegenteil"* — mit zwei belegten Fällen),
  zwei konkrete Folge-Slices mit auffindbaren Dateien in `open/`, und eine
  nachprüfbare Beobachtung (48/29/26 gegen 60/38/26, oben nachgerechnet).
  Keine Floskel, kein überholtes Futur. Die Notiz nennt auch, was **nicht**
  geliefert wurde. Der Absatz zur dritten Zahl (26 vorher wie nachher) ist der
  stärkste Teil: Er widerlegt die eigene Messung als Fortschrittsmaß, statt sie
  zu verteidigen.
- geprüft, ohne Befund: **Die zwei Folge-Slices sind regelkonform und nehmen
  ihre Punkte an.** Beide tragen die Kopffelder (`Verantwortlich: —` ist für
  `open/` die richtige Antwort, `modul-05` §Lifecycle), einen §1 mit
  Ausschlusspunkten samt Begründung je Punkt, nicht-zirkuläre Start-Trigger
  (`slice-187` bzw. `slice-188` in `done/`), vorab benannte Rückführungen und
  eine DoD, die die Prüf-Ebene ausdrücklich verlangt. `slice-188` nimmt die
  zwei in §1 abgetretenen Paare namentlich auf, `slice-189` die zwei
  Spec-Straten; beide zitieren das Instrument statt es zu wiederholen. Die
  Zählung *„elf bereits abgeglichene Paare"* in `slice-188` §1 ist mit 5+3+3
  korrekt. (Der eine Ausschlusspunkt, der nicht trägt: F-1.)
- geprüft, ohne Befund: **Commit-Hygiene.** Sechs Commits, jeder nennt
  `slice-187`; alle vier `docs(planning)`-Commits berühren ausschließlich
  `docs/plan/planning/`, die zwei `docs(harness)`-Commits nur `AGENTS.md` und
  `harness/conventions.md`. Der Lifecycle-Move liegt in einem eigenen Commit
  ohne Inhaltsänderung an der bewegten Datei; keine ADR angefasst; kein
  Produkt-Code, keine Suppression, keine Host-Toolchain.
- geprüft, ohne Befund: **Der Ruhe-Marker steht richtig.** `in-progress/` trägt
  seit `bd2ba58` einen Slice, und *„Nichts in Arbeit"* ist im selben Commit
  gefallen — die Richtung, die `modul-06` §Roadmap-Struktur verlangt. (Wem die
  Streichung zugeschrieben ist: F-13.)
- geprüft, ohne Befund: **Sub-Area-Wahl (§9).** `HARNESS` und `PLAN` sind beide
  in `harness/conventions.md` §Modus-Deklaration geführt, beide Greenfield, der
  Begründungsblock entfällt zu Recht. Die Begründung, warum
  `docs/plan/planning/README.md` trotz seines Pfads unter `HARNESS` zählt,
  folgt der Berührungs-Frage aus `grundlagen-bootstrap.md` und trägt. (Was der
  Sichtungs-Schritt daraus **nicht** gemacht hat: F-2.)
- geprüft, ohne Befund: **Der zweite Sichtungs-Kanal ist neu und trägt.** §9
  liest zum ersten Mal den Review-Report des Vorgängers und benennt zwei
  Klassen, die nicht im Register stehen; eine davon (*Zusage in der Doku,
  Grenze nur im Konfigurations-Kommentar*) ist unmittelbar zu §3.5 geworden.
  `modul-05` §Closure- und Lerneintrag-Regeln nennt die Finding-Klasse als
  dritte Zähler-Quelle; ohne diesen Griff zählt sie nicht mit.
- **Nicht geprüft, und darum keine Aussage:** `make gates`, `make verify`,
  `make doc-check`, `make doc-immutable` habe ich nicht ausgeführt. Ebenso
  ungeprüft bleibt die Gegenrichtung des Abgleichs über die volle Länge der drei
  Gegenstücke — ich habe sie stichprobenartig angesehen, und F-5 ist ein
  Treffer aus dieser Stichprobe; weitere Stellen derselben Art sind damit nicht
  ausgeschlossen.

## Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 5 |
| MEDIUM | 6 |
| LOW | 3 |
| INFO | 1 |

**Finding-Klassen dieses Laufs:** Adresse nimmt die Sendung nicht an ·
Sichtungs-Schritt liest eine andere Sub-Area als die deklariert berührte ·
Chronik in gelesenen Dateien · Korrektur am Bestand ohne das dafür deklarierte
Werkzeug · Vollständigkeits-Haken gesetzt, ohne den Gegenstand zu erschöpfen ·
Probe trifft ihren Gegenstand nicht · Slice-Provenienz aus Gedächtnis statt aus
`git log` behauptet · Geltungsbereich einer Messung nur in eine Richtung
genannt · Reibung benannt, aber weder gebucht noch als Kennung angelegt ·
Kandidaten-Klassifikation gröber als der Kandidat · Werkzeug-Zuschreibung für
eine Leistung, die das Werkzeug nicht erbringt · dieselbe Zusage zweimal in
derselben Datei · Beleg für einen abgeschlossenen Vorgang nachgetragen

**Zähler-Hinweis, mit der Regel dazu:** F-5, F-6 und F-8 tragen dieselbe
Klasse. `modul-06` §Das Beobachtungs-Register sagt, zwei Funde im selben
Vorgang seien *eine* Gelegenheit — sie ergeben also **einen** Beleg, nicht
drei. Dieser eine Beleg hebt
`BEO-PLAN/vollstaendigkeits-haken-ohne-erschoepften-gegenstand` von 2× auf
**3×**, und `state.md` sagt für diesen Fall: *„Beim dritten Mal ist es eine
Lücke."* Dasselbe gilt für F-3 gegenüber
`BEO-HARNESS/chronik-in-gelesenen-dateien` (2× → 3×). Die Aussage in §7 und §9,
kein Eintrag erreiche mit diesem Slice 3×, ist für die sechs **gesichteten**
Einträge richtig und für die ungesichteten nicht geprüft (F-2).

**Wiederkehrend gegenüber dem Vorgänger-Report:** *Vollständigkeits-Haken
gesetzt, ohne den Gegenstand zu erschöpfen* (dort F-4/F-5, hier F-5/F-6/F-8) und
*Kandidaten-Klassifikation gröber als der Kandidat* (dort F-10, hier F-12) —
beide zum **dritten Mal in Folge** an derselben Arbeitsform. *Beleg für einen
abgeschlossenen Vorgang nachgetragen* (dort F-12, hier F-15) zum zweiten Mal.
**Erledigt sind zwei:** *Messung als Beleg ohne reproduzierbares Instrument*
(dort F-6) — §2 löst sie vollständig auf, nachgerechnet; und *Zielsatz nach
Plan-Änderung nicht nachgezogen* (dort F-8) — Titel und §1 tragen den
korrigierten Umfang.

## Verdikt

**Merge-blockierend: ja**, für F-1 und F-4.

**F-1** ist der schwerste: Ein offener Punkt, den `slice-185` ausdrücklich an
`slice-186` adressiert hat, ist von keinem der drei Folge-Slices angenommen
worden und wird von `slice-189` mit einer gegen `slice-185` falschen Begründung
ausgeschlossen. Solange er nirgends steht, trägt die Zusage *„keine offene
Aufgabe, sondern eine Restmenge mit Zähler"* nicht — und das ist die Zusage, auf
der der Neu-Zuschnitt und zwei Risiko-Ausgänge ruhen. **F-4** steht gegen die
Prämisse, mit der der Risiko-Ausgang 4 als *entfallen* geschlossen ist: Die
Datei nennt das Werkzeug für eine Korrektur, und `MR-020` hat es für denselben
Eintrag schon einmal benutzt.

**F-2, F-3 und F-5 sind je einzeln in wenigen Zeilen behebbar**, aber sie hängen
zusammen: Der Sichtungs-Schritt hat die Sub-Area nicht gelesen, in der dieser
Slice arbeitet, und beide Einträge, die er dort gefunden hätte, sind in diesem
Diff ein drittes Mal aufgetreten. Ob daraus eine Lücke mit Ausgang wird, ist
eine Planner-Entscheidung, kein Reviewer-Vorschlag.

**Was ausdrücklich trägt:** Die Ausgangsmessung ist der erste reproduzierbare
Beleg dieser Kette — ich habe sie mit dem Instrument aus dem Plan nachgerechnet
und alle vierzehn Zahlen bestätigt. Der Neu-Zuschnitt vor der Arbeit ist eine
ehrliche Plan-Korrektur, keine Umgehung. Und die Closure-Notiz widerlegt ihre
eigene Messung als Fortschrittsmaß, statt sie zu verteidigen — das ist der
Lerneintrag, den die Pflicht meint.

**Übergabe:** an die Implementer-Rolle. Dieser Report ist Lauf-Beleg, keine
Verifikation — DoD- und Spec-Konformität prüft der Verifier separat (Modul 11).
