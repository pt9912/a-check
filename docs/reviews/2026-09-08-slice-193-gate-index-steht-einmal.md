# Review-Report: slice-193 — 2026-09-08

**Review-Art:** Plan — geprüft gegen den Slice-Plan, gegen die Baseline-Regeln,
die er in Anspruch nimmt, und gegen die Tat selbst (Modul 10 §Drei Review-Arten).
Kern-Behauptungen sind die Differenz-Messung („nichts verloren"), die drei
Mutations-Proben, der `mentions`-Umweg, der Fund an der entfallenen
`structure`-Regel, die drei Risiko-Ausgänge und die Register-Sichtung.

**Gegenstand:** Commit-Range `45d91ac..HEAD` (`0ed8644`) — `498e7bc`
(„der Gate-Index steht einmal") und `0ed8644` („Umsetzung, Risiko-Ausgaenge und
Closure"); der Übergangs-Commit `45d91ac` ist die Basis.

**Skill:** `.harness/skills/reviewer.md` @ `0ed8644` · <!-- d-check:ignore -->
**Modell:** claude-opus-5[1m] · **Datum:** 2026-09-08

**Eingangs-Kontext** (die Verträge, gegen die geprüft wurde):

- slice-193 (Plan, `in-progress/`); slice-192 (Vorgänger, Plan und Report aus
  `done/slice-192-archiv.zip` gelesen), slice-188, slice-189 (`open/`)
- `AGENTS.md` §3 (Hard Rules, insbesondere §3.7), §4, §5, §6
- `harness/README.md` §Sensors und §Nicht-Gates; `harness/conventions.md`
  §Modus-Deklaration pro Sub-Area; `harness/sensors/doc-mentions.md`,
  `harness/sensors/doc-planning.md`
- `AC-QA-02` als die vom Plan bezogene Kennung
- `v6.6.0` · `templates/AGENTS.template.md` §4 · `templates/harness/README.template.md`
  §Sensors · `templates/.d-check.yml` (targets-Block) ·
  `regelwerk/grundlagen-harness-dateien.md` §harness/README.md als Einstiegspunkt ·
  `regelwerk/modul-13-quality-gates.md` §Hard Rule ·
  `regelwerk/modul-05-planning-harness.md` §Offene Risiken werden bei Closure
  aufgelöst · `regelwerk/modul-06-roadmap.md` §Das Beobachtungs-Register
- Beobachtungs-Register: alle 64 Verzeichnisse (`state.md`-Kopfzeile und
  Beleg-Zahl), im Detail die acht aus §9 sowie
  `BEO-HARNESS/aggregat-aufzaehlung-hinkt-dem-makefile-hinterher`,
  `BEO-PLAN/risiko-eingetreten-und-im-slice-aufgeloest`,
  `BEO-GATE/zusage-weiter-als-ihre-durchsetzung`
- `.harness/skills/closure-note-reviewer.md` als Maßstab für §8
- **Eigene Läufe statt zitierter Ergebnisse:** `make doc-targets`,
  `make doc-mentions`, `make doc-check`, `make doc-structure`, `make verify` —
  je einzeln in eine Datei umgeleitet und am Exit-Code gemessen; dazu **sechs**
  eigene Mutations-Proben (drei nachgefahrene, drei neue). Alle Proben
  zurückgesetzt, `git status` leer.

---

## Findings

### F-1 — Der Fund aus §3.4/§8 ist falsch: die entfallene `structure`-Regel meldete nicht grün, sondern `section-column-missing`

- `kategorie`: HIGH
- `quelle`: nachweislich falsche Tatsachenbehauptung (Reviewer-Skill
  §Klassifikation); `v6.6.0` · `regelwerk/modul-13-quality-gates.md` §Hard Rule
  (maschinelle Hälfte)
- `pfad`: `docs/plan/planning/in-progress/slice-193-gate-index-steht-einmal.md`
  §3.4 (Z. 124–132), §8 (Z. 253–256), §9 (Z. 292) · `.d-check.yml` Z. 383–389
- `befund`: Der Plan schreibt, die Zellengrenze auf der *Zweck*-Spalte sei nach
  dem Wegfall der Tabelle **grün** geblieben, „statt fail-closed zu melden".
  Nachgefahren mit der **wortgleichen** Original-Regel aus `45d91ac:.d-check.yml`
  gegen den heutigen `AGENTS.md`-Stand meldet `make doc-structure`
  `AGENTS.md:166 :: ## 4. Quality Gates :: Spalte Zweck section-column-missing`,
  Exit 2 — die Regel ist fail-closed. Vier Zeilen über der gestrichenen Stelle
  führt dieselbe Datei `section-column-missing` bereits als Befund-Code dieses
  Moduls („gemessen, slice-181"). Was grün blieb, war ein Lauf ohne diesen
  Prüfer: `doc-structure` hängt in `verify`, nicht in `gates`; §3.4 nennt den
  Aufruf nicht. Die Folge trägt weiter: §9 bucht den Fall als „bestätigt,
  verkörpert" auf `BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf`, und der Satz
  steht als dauerhafter Kommentar in `.d-check.yml`.
- `verifizierbar`: ja — Original-Regel aus `git show 45d91ac:.d-check.yml`
  wieder einsetzen, `make doc-structure`; Exit 2 mit dem genannten Grund-Code.
  **Kein Gate fängt es:** dass eine *Begründung* für eine Streichung stimmt,
  prüft kein Lauf.
- `klasse`: Messung ohne reproduzierbares Instrument

### F-2 — Die neu geschriebene Grenze 4 in `doc-mentions.md` nennt eine Link-Form, die aus dem ADR-Index nicht auflöst

- `kategorie`: HIGH
- `quelle`: nachweislich falsche Tatsachenbehauptung; `AGENTS.md` §3.7 (ein
  Kommentar/eine Grenze beschreibt, was da ist)
- `pfad`: `harness/sensors/doc-mentions.md` Z. 39–42
- `befund`: Der Slice ersetzt die alte Grenze 4 durch die Aussage, der Umweg
  „stünde dort ebenso offen (`../../docs/plan/adr/<datei>.md` löst aus dem
  ADR-Index auf und trägt den vollen Pfad)". Aus `docs/plan/adr/README.md`
  löst `../../docs/plan/adr/X.md` nach `docs/docs/plan/adr/X.md` auf;
  `make doc-check` meldet `target-missing`. Die tragende Form wäre
  `../../../docs/plan/adr/X.md` — drei Ebenen, nicht zwei. Der Satz, der die
  verbleibende Hürde für eine Abwägung erklärt, nennt damit eine Form, die die
  Abwägung gar nicht zulässt.
- `verifizierbar`: ja — eine ADR-Index-Zeile testweise auf `../../…` umstellen,
  `make doc-check`: Exit 2, `target-missing`; mit `../../../…` Exit 0.
- `klasse`: Pfad-Form behauptet statt gemessen

### F-3 — `.d-check.yml` trägt zwei überholte Begründungsblöcke unmittelbar an den Zeilen, die sie widerlegen

- `kategorie`: HIGH
- `quelle`: `AGENTS.md` §3.7 (Kommentar beschreibt, was da ist — nicht den
  abwesenden Text); Reviewer-Skill §Klassifikation, HIGH-Punkt *Kommentar trägt
  keine der Kommentar-Klassen*
- `pfad`: `.d-check.yml` Z. 171–174 und Z. 449–460
- `befund`: Über `doc-tables: [harness/README.md]` steht unverändert *„BEIDE
  Doku-Tabellen … Ohne die zweite wäre die Ablösung eine Scope-Verkleinerung"* —
  eine Begründung für eine zweielementige Liste, die jetzt ein Element hat.
  Über `documents: ["harness/README.md"]` steht unverändert *„WARUM NUR
  AGENTS.md ALS DOKUMENT: harness/README.md … kann das nicht ändern, ohne seine
  Links zu brechen … Der volle Pfad steht nur in AGENTS.md Paragraph 4 … Die
  zweite Tabelle bleibt damit ungewächtert"* und *„WARUM NICHT DIE ADRs …
  scheitert nicht am Willen, sondern an der Pfad-Form"* — zwölf Zeilen über dem
  neuen Block, der genau das Gegenteil misst und schreibt. Beide Stellen liegen
  im Diff dieses Slice; ein Lauf, der die Datei von oben liest, liest zuerst die
  widerlegte Fassung.
- `verifizierbar`: nein — kein Sensor prüft, ob ein Kommentar noch zu seiner
  Zeile passt (`AGENTS.md` §3.7 nennt die Regel ausdrücklich inferentiell).
- `klasse`: Kommentar beschreibt den abgelösten Zustand

### F-4 — `Makefile` behauptet weiter `AGENTS.md` §4 als Erwähnungs-Ort, in Kommentar und in der `make help`-Zeile

- `kategorie`: HIGH
- `quelle`: nachweislich falsche Tatsachenbehauptung; `AGENTS.md` §3.7
- `pfad`: `Makefile` Z. 132–133 und Z. 139
- `befund`: Der Kommentar über dem Target sagt *„Jede Datei unter
  harness/sensors/ muss in AGENTS.md Paragraph 4 genannt sein"*, die
  `##`-Hilfezeile *„jede harness/sensors/-Datei ist in AGENTS.md genannt
  (Modul mentions)"*. Beides widerspricht `.d-check.yml`
  (`mentions.documents: ["harness/README.md"]`) im selben Commit; eine
  Sensor-Datei, die nur in `AGENTS.md` §4 stünde, meldet heute
  `artifact-unmentioned`. Die `##`-Zeile ist zudem die Ausgabe von `make help`,
  also die Fassung, die ein Lauf ohne Dateilektüre sieht.
- `verifizierbar`: ja — `make help` gegen `.d-check.yml` halten; die
  Gegenprobe (einen Link auf die Geschwister-Form zurücksetzen) meldet
  `artifact-unmentioned` gegen `harness/README.md`, nicht gegen `AGENTS.md`.
- `klasse`: Verweis auf den abgelösten Gate-Index nicht nachgezogen

### F-5 — Zwei lebende Dokumente führen `AGENTS.md` §4 weiter als Deklarations-Ort der Targets

- `kategorie`: HIGH
- `quelle`: nachweislich falsche Tatsachenbehauptung; `AGENTS.md` §4 (neue
  Fassung: „Diese Datei führt die Liste nicht")
- `pfad`: `harness/conventions.md` Z. 278 (Zeile *Gate-/Werkzeug-Schicht* der
  Modus-Deklaration) · `docs/user/releasing.md` Z. 116 (Freigabe-Item 4)
- `befund`: `harness/conventions.md` begründet den GF-Modus der Sub-Area `GATE`
  mit *„jedes Target ist in `AGENTS.md` §4 deklariert, bevor es zählt;
  `gate-consistency` erzwingt das"*; `docs/user/releasing.md` führt als
  Freigabe-Item *„`AGENTS.md` §4 beschreibt nur real existierende Targets"* mit
  Beleg-Slot `make gate-consistency`. Die erste Hälfte beider Sätze ist mit
  diesem Slice falsch geworden. Die zweite Hälfte war schon vorher falsch —
  `gate-consistency` (1)+(2) ist seit slice-079 abgelöst, was der Skriptkopf
  selbst dokumentiert —, sodass beide Sätze jetzt in beiden Hälften unzutreffend
  sind. §9 des Plans bucht dazu *„die drei `AGENTS.md`-§4-Verweise außerhalb
  wurden einzeln gelesen"*; geändert wurden vier Dateien, stehengeblieben sind
  in lebenden Dokumenten mindestens sieben Nennungen (F-4, F-5, F-7).
- `verifizierbar`: ja — `grep -rn "AGENTS.md.*§ *4"` über die lebenden
  Dokumente gegen den neuen §4-Text und gegen `tools/gate-consistency.sh` Z. 5–8.
- `klasse`: Verweis auf den abgelösten Gate-Index nicht nachgezogen

### F-6 — Der Register-Eintrag steht bei 3× ohne einen der drei Ausgänge; der Lese-Schritt ist in diesem Repo diese Closure

- `kategorie`: HIGH
- `quelle`: `v6.6.0` · `regelwerk/modul-06-roadmap.md` §Das
  Beobachtungs-Register (*„Nicht zulässig ist ein Eintrag, der eine Closure ohne
  Ausgang übersteht"*; Tabelle *Träger im Repo ohne Wellen*: Lese-Schritt =
  Slice-Closure §7, vor dem `git mv`); Reviewer-Skill §Klassifikation,
  HIGH-Punkt *Zustandsfeld trägt Chronik*, zweite Ausprägung
- `pfad`: `docs/plan/planning/observations/BEO-PLAN/kandidaten-klassifikation-groeber-als-der-kandidat/state.md`
  Z. 1 · Plan §8 (Z. 258–265) und §9 (Z. 290)
- `befund`: `state.md` trägt `**Stand:** offen (3×) — Ausgang beim Lese-Schritt
  fällig`, der Plan verschiebt die Zuweisung ausdrücklich „beim nächsten
  Lese-Schritt". a-check fährt wellenlos; damit ist der Lese-Schritt die
  Slice-Closure selbst, und ein anderer Träger existiert nicht. Von den drei
  Ausgängen wäre *geplant* offen — er verlangt eine Kennung, und der Plan nennt
  keine Slice-Datei, die die Schreibregel schreibt. Der Eintrag übersteht damit
  eine Closure ohne Ausgang. `make verify` bleibt grün: `verify-observations`
  prüft Deckung (Verzeichnis vorhanden, `evidence/` nicht leer), nicht den Stand.
- `verifizierbar`: nein — die Ausgangs-Pflicht ab 3× hat in diesem Repo keinen
  Sensor; `make verify` deckt nur die Deckung.
- `klasse`: Zustandsfeld bei 3× ohne einen der drei Ausgänge

### F-7 — Drei weitere Zeiger auf den alten Index sind stehengeblieben

- `kategorie`: MEDIUM
- `quelle`: `AGENTS.md` §3.7; Reviewer-Skill §Klassifikation (LOW/MEDIUM:
  fehlender bzw. falscher Querverweis)
- `pfad`: `.claude/hooks/pretooluse-command-guard.sh` Z. 68 ·
  `tools/gate-consistency.sh` Z. 8 · `.harness/skills/reviewer.md` Z. 60
- `befund`: Der Guard-Kommentar begründet seine `GATES`-Liste mit *„Vollstaendig
  gegen die deklarierten Pruef-Targets (AGENTS.md §4)"*; `gate-consistency.sh`
  verortet die freigebliebenen Nummern (1)+(2) *„in den Closure-Notizen … und in
  AGENTS §4"*; der Reviewer-Skill verlangt, eine Zusage aus einem
  `.d-check.yml`-Kommentar gehöre *„in `AGENTS.md` §4 oder eine Sensor-Datei"*.
  Alle drei zeigen auf einen Ort, der die Liste nicht mehr führt. Anders als
  F-4/F-5 sind das Zeiger und Provenienz-Angaben, keine Regel-Zusagen — die
  Aussage ist irreführend, nicht unmittelbar handlungsleitend.
- `verifizierbar`: ja — dieselbe `grep`-Menge wie F-5.
- `klasse`: Verweis auf den abgelösten Gate-Index nicht nachgezogen

### F-8 — Die Register-Sichtung in §9 übergeht zwei Einträge der deklariert berührten Sub-Areas

- `kategorie`: MEDIUM
- `quelle`: `v6.6.0` · `regelwerk/modul-05-planning-harness.md` §Zwei Schritte
  vor der Modus-Begründung (*„Offene Beobachtungen sichten"*); `AGENTS.md` §5
- `pfad`: `docs/plan/planning/in-progress/slice-193-gate-index-steht-einmal.md`
  §9 (Z. 285–300)
- `befund`: §9 sagt zu, das Register *„nach **allen** drei Kürzeln"* durchgegangen
  zu sein, und listet acht Zeilen. Nicht darunter:
  (a) `BEO-HARNESS/aggregat-aufzaehlung-hinkt-dem-makefile-hinterher` (2×) —
  Titel und Gegenstand sind wörtlich *„Die Aggregat-Aufzählung in `AGENTS.md` §4
  hinkt dem Makefile hinterher"*, und ihr `state.md` nennt als Abhilfe *„der
  Absatz müsste auf einen Zeiger schrumpfen (`welche genau, sagt das Makefile`)
  … wäre beim dritten Auftreten die naheliegende"* — genau das, was §3.2 dieses
  Slice ausgeführt hat. Der Eintrag steht seither unverändert bei „offen (2×)"
  und beschreibt eine bereits erfolgte Abhilfe als künftige.
  (b) `BEO-PLAN/risiko-eingetreten-und-im-slice-aufgeloest` (2×) — ihr
  Gegenstand ist der Ausgang, den §7 Risiko 1 wählt (siehe F-9).
  Beide liegen in Sub-Areas, die §9 selbst als berührt führt.
- `verifizierbar`: nein — welche Einträge einschlägig sind, ist ein Urteil über
  Zuordnung; die Deckungs-Hälfte prüft `make verify-observations` und ist grün.
- `klasse`: Register-Sichtung übergeht Einträge der berührten Sub-Areas

### F-9 — Der Ausgang von Risiko 1 ist keiner der drei aus der geschlossenen Menge

- `kategorie`: MEDIUM
- `quelle`: `v6.6.0` · `regelwerk/modul-05-planning-harness.md` §Offene Risiken
  werden bei Closure aufgelöst (*eingetreten ⇒ Carveout oder Folge-Slice mit
  ID*); `AGENTS.md` §5
- `pfad`: `docs/plan/planning/in-progress/slice-193-gate-index-steht-einmal.md`
  §7 (Z. 190–197)
- `befund`: Der Ausgang lautet *„eingetreten → im Slice aufgelöst, und die Klasse
  liegt im Beobachtungs-Register"*. Der Zweig *eingetreten* verlangt einen
  Carveout oder eine Folge-Slice-ID; beides fehlt. Für den Fall, dass eine
  **Klasse** übrigbleibt, hat das Register die Lesart bereits geschärft
  (`risiko-eingetreten-und-im-slice-aufgeloest`, Beleg slice-192: *„für die gibt
  es einen passenden Ausgang: weiter offen → Register"*) — das wäre hier
  *weiter offen*, nicht *eingetreten*. Damit ist dies das dritte Vorkommen der
  offenen Klasse, ungezählt und in §9 nicht gesichtet.
  `make verify-risiko-ausgaenge` bleibt grün: es prüft die Form (ein Ausgang je
  Risiko), nicht die Zuordnung zu einem der drei Zweige.
- `verifizierbar`: ja/teilweise — `make verify-risiko-ausgaenge` bestätigt nur
  die Existenz eines Ausgangs; die Zweig-Zuordnung ist Urteil.
- `klasse`: Risiko-Ausgang außerhalb der geschlossenen Dreier-Menge

### F-10 — Der neue `AGENTS.md` §4 sagt Prosa-Deckung zu und nennt im selben Absatz einen Prüfer, der Prosa nicht sieht

- `kategorie`: MEDIUM
- `quelle`: `v6.6.0` · `regelwerk/modul-13-quality-gates.md` §Hard Rule
  (maschinelle Hälfte); `BEO-GATE/zusage-weiter-als-ihre-durchsetzung`
- `pfad`: `AGENTS.md` Z. 177–182
- `befund`: *„Kein Target nennen, das im Makefile nicht existiert — auch nicht
  in Prosa"* steht unmittelbar vor *„Die maschinelle Hälfte dieser Regel ist
  `make doc-targets`"*, ohne benannte Grenze. Gemessen: ein Phantom-Target, nur
  in der Prosa des Gate-Index genannt, lässt `make doc-targets` mit Exit 0
  durchlaufen; umgekehrt meldet der Sensor `gate-undocumented`, obwohl das
  Target zwei Zeilen tiefer in der Prosa steht. Der Slice weiß das — §3.1 stützt
  seine Fünfer-Zahl genau darauf (*„für den `targets`-Sensor ist Prosa kein
  Eintrag"*) —, schreibt die Grenze aber nicht an die Zusage. §7 Risiko 3 lehnt
  einen Beleg an eben diesem Register-Eintrag ausdrücklich ab, weil hier *„keine
  Zusage weiter reicht als ihr Prüfer"*.
- `verifizierbar`: ja — Phantom nur in die Prosa von `harness/README.md`
  §Sensors einfügen, `make doc-targets`: Exit 0.
- `klasse`: Zusage reicht weiter als ihre Durchsetzung

### F-11 — Der als wertvoller bezeichnete Lerneintrag hat keinen Zielort und keine Kennung

- `kategorie`: MEDIUM
- `quelle`: `AGENTS.md` §5 (Lerneintrag-Formen: geschärfte Regel · neuer Sensor ·
  benannte Spec-Lücke); `v6.6.0` · `regelwerk/modul-06-roadmap.md` §Das
  Beobachtungs-Register (Ausgang *verkörpert* verlangt Zielort und
  Herkunfts-Anker)
- `pfad`: `docs/plan/planning/in-progress/slice-193-gate-index-steht-einmal.md`
  §8 (Z. 232–244)
- `befund`: *„Eine für unmöglich erklärte Kopplung gehört gemessen, bevor sie zur
  Grenze wird"* wird als *„der zweite Lerneintrag ist der wertvollere"* geführt.
  Er steht danach in keiner Datei, die ein Lauf liest — weder in `AGENTS.md` §5
  noch in den Mess-Regeln des Reviewer-Skills —, und keine Register-Kennung führt
  ihn, sodass ihn auch kein Zähler aufnimmt. Der erste Lerneintrag hat seinen
  Zielort (der Umzug selbst); dieser hat keinen. Dieselbe Klasse hat der Report
  zu slice-192 als F-8 gebucht.
- `verifizierbar`: nein — inferentiell; kein Sensor prüft, ob eine geschärfte
  Regel irgendwo steht.
- `klasse`: Lerneintrag als geschärfte Regel deklariert, ohne Zielort

### F-12 — §9 zitiert die Finding-Klasse des Vorgänger-Reports umformuliert

- `kategorie`: LOW
- `quelle`: `v6.6.0` · `templates/docs/reviews/review-report.template.md`
  (Norm: die Klassen-Bezeichnung muss über Läufe hinweg stabil sein);
  Reviewer-Skill §Output-Schema
- `pfad`: `docs/plan/planning/in-progress/slice-193-gate-index-steht-einmal.md`
  §9 (Z. 302–306)
- `befund`: §9 nennt als schärfste Klasse des slice-192-Reports *„eine Regel zu
  zitieren ist nicht, sie anzuwenden"*. Der Report führt diese Formulierung
  nicht; seine zwölf `klasse`-Felder nennen an dieser Stelle *„Sichtungs-Schritt
  liest eine andere Sub-Area als die deklariert berührte"* (F-3). Eine
  umformulierte Klasse zählt der Steering-Loop getrennt — die Norm existiert
  genau deshalb.
- `verifizierbar`: ja — `klasse`-Felder des Reports zu slice-192 gegen §9.
- `klasse`: Finding-Klasse umformuliert statt zitiert

### F-13 — Ein berührtes Verzeichnis hat keine Zeile in der Modus-Deklaration

- `kategorie`: LOW
- `quelle`: `harness/conventions.md` §Modus-Deklaration pro Sub-Area;
  `v6.6.0` · `grundlagen-bootstrap.md` §Was ist eine Sub-Area?
- `pfad`: `docs/plan/carveouts/README.md` Z. 72 · Plan §9 (Z. 280–281)
- `befund`: §9 leitet die Sub-Areas aus `git diff --name-only` ab und beantwortet
  zu `docs/plan/carveouts/` nur, dass es **nicht** `ADR` ist. Welche Sub-Area es
  dann ist, bleibt offen: die Modus-Deklaration führt für `docs/plan/carveouts/`
  keine Zeile. Dieselbe Klasse hat der Report zu slice-192 als F-12 gebucht,
  dort für `.harness/skills/` — die dortige Lücke besteht ebenfalls fort.
- `verifizierbar`: ja — Pfadliste der Modus-Deklaration gegen
  `git diff --name-only 45d91ac..HEAD`.
- `klasse`: Berührtes Verzeichnis ohne Zeile in der Modus-Deklaration

### F-14 — Die erweiterte Aufzählung „in keinem Aggregat" ist wieder unvollständig

- `kategorie`: LOW
- `quelle`: `BEO-HARNESS/aggregat-aufzaehlung-hinkt-dem-makefile-hinterher`
  (*„abschließende Aufzählung neben einer maschinenlesbaren Quelle"*)
- `pfad`: `harness/README.md` Z. 118–123
- `befund`: Der Absatz zählt jetzt fünf Targets auf, die *„in keinem Aggregat
  hängen"* und trotzdem in der Gate-Tabelle stehen. `make trace-check` und
  `make doc-immutable` erfüllen dasselbe Kriterium — beide stehen oben in der
  Tabelle, beide hängen weder in `gates` noch in `verify` — und fehlen. Der
  Slice hat die Aufzählung von drei auf fünf erweitert und dabei die
  Vollständigkeit nicht hergestellt; §1 schließt das Umsortieren aus, das
  Erweitern der Liste aber nicht.
- `verifizierbar`: ja — Zielliste von `gates`/`verify`/`ci` im `Makefile` gegen
  die Target-Zellen der Sensors-Tabelle.
- `klasse`: Abschließende Aufzählung neben einer maschinenlesbaren Quelle

### F-15 — Eine historische Aussage in `doc-planning.md` wurde auf den neuen Zustand eingeebnet

- `kategorie`: LOW
- `quelle`: `AGENTS.md` §3.7 (Zustandsfelder/Kommentare tragen keine
  umgeschriebene Chronik); `BEO-HARNESS/massen-ersetzung-trifft-die-historische-aussage`
- `pfad`: `harness/sensors/doc-planning.md` Z. 14
- `befund`: Aus *„Bis slice-186 sagten `AGENTS.md` §4, `harness/README.md`
  §Sensors und der Vertrag oben *benennt ihn* zu"* wird *„Bis slice-186 sagten
  der Gate-Index und der Vertrag oben … zu"*. Zu jenem Zeitpunkt gab es zwei
  Indizes, und dass die Zusage an **beiden** Stellen stand, war der Punkt der
  Aussage. §9 bucht die Klasse als *„bedient — die drei `AGENTS.md`-§4-Verweise
  außerhalb wurden einzeln gelesen, nicht ersetzt"*; dieser Verweis war
  historisch und wurde ersetzt.
- `verifizierbar`: ja — `git show 45d91ac:harness/sensors/doc-planning.md` gegen
  `HEAD`.
- `klasse`: Historische Aussage auf den neuen Zustand umgeschrieben

### F-16 — Die tragende Link-Form steht ohne Kennzeichnung neben zwei anderen Formen in derselben Datei

- `kategorie`: INFO
- `quelle`: Maintainability; `harness/sensors/doc-mentions.md` §Grenze 1
  (der Preis ist dort benannt)
- `pfad`: `harness/README.md` §Sensors und §Nicht-Gates
- `befund`: Die 16 Sensor-Links tragen jetzt `../harness/sensors/<name>.md`,
  während Nachbarzellen derselben Tabellen `conventions.md` und
  `../docs/plan/…` verwenden. Für einen Leser sieht der Umweg wie ein Tippfehler
  aus; ihn zu „korrigieren" macht `make doc-mentions` blind, ohne dass ein Link
  bricht. Der Preis ist in der Sensor-Datei benannt — an der Stelle, an der die
  Links stehen, steht er nicht. Kein erwarteter Handgriff; der Hinweis gehört
  der Rolle, die über einen Sensor auf die Form entscheidet.
- `verifizierbar`: ja — eine Zelle auf die Geschwister-Form zurücksetzen:
  `make doc-check` Exit 0, `make doc-mentions` Exit 2 (`15 von 16`).
- `klasse`: Tragende Form ohne Kennzeichnung an ihrer Stelle

## Negativbefunde

- geprüft, ohne Befund: **Es ist nichts verloren gegangen — eigenständig
  gemessen.** Target-Namen je Tabellen-**Zelle** (alle `make X` pro Zelle, nicht
  nur das erste) und getrennt aus der Prosa erhoben, aus `45d91ac:AGENTS.md` §4,
  `45d91ac:harness/README.md` §Sensors+§Nicht-Gates und dem heutigen Stand:
  41 Zellen-Targets in §4, 36 Zellen- plus 3 Prosa-Nennungen im alten Index,
  heute **41** Zellen-Targets im Index. `AGENTS.md`-Zellen ohne Zeile im neuen
  Index: **null**; neue Zellen: genau die fünf des Plans (`archive-wave-test`,
  `commit-scope-check`, `doc-commits`, `doc-tracked`, `verify`); entfernte
  Zellen: **null**. Die Zahlen aus §3.1 (3 / 2 / 5 / 0) und die Kennzahlen aus
  §2/§3.2 (7094→1170 Zeichen, 41 bzw. 28 Zeilen) treffen exakt, wenn die
  Abschnitts-Überschrift mitgezählt wird. **Geltungsbereich:** gemessen wurden
  Vorkommen der Form `make <name>` in den beiden Abschnitten; ein Target, das
  ohne `make`-Präfix genannt wäre, sähe dieses Instrument nicht — im Bestand
  gibt es keins.
- geprüft, ohne Befund: **Die drei Mutations-Proben aus §3.5 sind echt und
  treffen ihren Gegenstand.** Alle drei nachgefahren, alle drei rot, mit den
  zitierten Meldungen: `Makefile:113 commit-scope-check gate-undocumented …
  Autoritäts-Doku harness/README.md`; `harness/README.md:99 gibtsnicht
  gate-phantom`; `harness/sensors/doc-planning.md:1 artifact-unmentioned` bei
  `mentions: 15 von 16`. Die erste Probe belegt zusätzlich mehr, als der Plan
  daraus zieht: die Prosa-Nennung desselben Targets zwei Zeilen tiefer verhindert
  den Befund **nicht** — damit ist die Begründung der Fünfer-Zahl aus §3.1
  unabhängig belegt.
- geprüft, ohne Befund: **Der `mentions`-Umweg trägt in beiden Richtungen.**
  `make doc-check` ist mit der Form `../harness/sensors/<name>.md` grün (die
  Links lösen auf), `make doc-mentions` meldet `16 von 16` über ein Dokument,
  und die Rückstellung einer einzigen Zelle auf die Geschwister-Form macht den
  Sensor blind, ohne einen Link zu brechen. Die Prämisse ist ebenfalls belegt:
  `mentions.resolve-from` bricht mit `field resolve-from not found`.
  `links.resolve-from` in `.d-check.yml` bindet nur `docs/plan/planning/*` und
  wird vom Umweg nicht berührt; die Ziel-Form erlaubt den Link auf der
  Target-Zelle ausdrücklich (*„Ein LINK auf der Zelle schadet nicht"*), und
  Target-Zellen mit Argument in der Code-Span gibt es weiterhin **null**.
- geprüft, ohne Befund: **`AGENTS.md` §4 ist die Ziel-Form, Satz für Satz.**
  Beide Normsätze aus `v6.6.0` · `templates/AGENTS.template.md` §4 stehen
  vollständig da (der eine Index samt Bindungs-Hinweis und *„Diese Datei führt
  die Liste nicht"*; *„Kein Target nennen … auch nicht in Prosa"*). Ergänzt sind
  drei repo-eigene Aussagen: der Zeiger auf `harness/sensors/`, die maschinelle
  Hälfte und die Mandatory-Definition. Die ersten beiden sind Baseline-Sätze mit
  repo-eigener Bindung (kein nachgeschriebener Normtext im Sinne der Mess-Regel
  *Geltungsbereich einer Messung*, dritte Hälfte); der Vorbehalt zur zweiten
  steht unter F-10. Die Mandatory-Aussage ist eigenständig und ersetzt eine
  abschließende Aufzählung durch einen Zeiger aufs `Makefile` — inhaltlich die
  Abhilfe, die `aggregat-aufzaehlung-hinkt-dem-makefile-hinterher` vorschlägt
  (Buchungs-Vorbehalt unter F-8).
- geprüft, ohne Befund: **Die drei Konfigurationen zeigen auf einen Index.**
  `targets.doc-tables`, `targets.authority` und `mentions.documents` nennen alle
  `harness/README.md`; `exempt-targets` bleibt namentlich
  (`[help, build, compile, hooks, arch-graph]`), kein Glob. Die
  Zellengrenzen-Regel auf `harness/README.md` §Sensors deckt Vertrags- und
  Bindung-Spalte weiter ab (Vorbehalt zur *Begründung* ihrer Nachbarin: F-1).
- geprüft, ohne Befund: **Gate-Lage auf dem geprüften Stand.** `make doc-check`,
  `make doc-targets`, `make doc-mentions`, `make doc-structure` und
  `make verify` je einzeln, alle Exit 0 (`verify`: 21 Anforderungen, 0 Waisen).
  `make gates` wurde **nicht** als Ganzes gefahren — die vier darin enthaltenen
  Doku-Sensoren und `verify` einzeln; die Code-Gates (`lint`, `test`,
  `coverage-gate`, `arch-check`) sind vom Diff nicht berührt.
- geprüft, ohne Befund: **Kein Spec-Stratum berührt.** Keine `AC-*` geändert,
  keine ADR angefasst, keine Datei unter `spec/` oder `docs/plan/adr/` im Diff;
  die Kopfzeile *Berührte Spec-Stellen: —* trägt. Die ADR-Nennungen von
  `AGENTS.md` §4 (ADR-0005, ADR-0006, ADR-0021, ADR-0037, ADR-0038, ADR-Index)
  sind korrekt **nicht** angefasst worden — sie sind `Accepted` und historisch.
- geprüft, ohne Befund: **Commit-Disziplin.** `498e7bc` trägt Scope `harness`
  und darf außerhalb von `docs/plan/planning/` arbeiten; `0ed8644` trägt Scope
  `planning` und berührt ausschließlich `docs/plan/planning/`. Beide Messages
  nennen die Slice-Kennung; der Übergangs-Commit `45d91ac` liegt vor der Arbeit
  auf dem Hauptzweig.
- geprüft, ohne Befund: **DoD-Größenregel.** Sieben Punkte, davon vier pro Slice
  konstant (Review, Closure-Notiz, Register, Risiko-Ausgänge) — drei
  Liefer-Punkte, an der Obergrenze. Der Review-Punkt steht korrekt **ungehakt**.
- geprüft, ohne Befund: **Der Beleg im Register ist formgerecht.**
  `evidence/slice-193.md` ist eine einzelne Datei mit der Vorgangs-Kennung als
  Namen, der Vorgang liegt (nach dem `git mv`) in `done/`; `observation.md`
  bleibt unverändert. Der Vorbehalt betrifft ausschließlich `state.md` (F-6).
- geprüft, ohne Befund: **Die Closure-Notiz trägt Substanz, nicht Floskel.**
  Maßstab `.harness/skills/closure-note-reviewer.md`: (a) ein Lernsignal mit
  Ursache (*„sie kostet die Stelle, plus jeden Lauf, der beide liest"*) und
  (c) eine nachprüfbare Aussage (*„die Differenz-Messung ist beidseitig null"*,
  hier unabhängig nachgerechnet) sind beide da; ein Folge-Slice ist nicht nötig,
  weil (a) und (c) tragen. Der dritte Absatz der Notiz ist inhaltlich falsch —
  das ist F-1, keine Floskel-Frage.
- geprüft, ohne Befund: **Die Sub-Area-Ableitung aus dem Diff ist vollständig.**
  Die in §9 genannten sieben Verzeichnisse decken `git diff --name-only
  45d91ac..HEAD` restlos; die Zuordnung zu `HARNESS`, `GATE` und `PLAN` trägt.
  Der Vorbehalt betrifft nur das eine unzugeordnete Verzeichnis (F-13).
- geprüft, ohne Befund: **Die übrigen sechs Register-Zeilen aus §9** —
  `zusage-weiter-als-ihre-durchsetzung`, `muster-trifft-nur-die-haeufige-schreibweise`,
  `vollstaendigkeits-haken-ohne-erschoepften-gegenstand`,
  `messung-ohne-reproduzierbares-instrument`, `chronik-in-gelesenen-dateien`,
  `massen-ersetzung-trifft-die-historische-aussage` — sind mit dem richtigen
  Kürzel und dem richtigen Zähler-Stand geführt. Drei der Berührungs-Arten halten
  der Gegenprobe nicht stand (F-1 gegen *messung-ohne-reproduzierbares-instrument*
  und *pruefer-ohne-gegenstand-oder-aufruf*, F-10 gegen
  *zusage-weiter-als-ihre-durchsetzung*, F-15 gegen
  *massen-ersetzung-trifft-die-historische-aussage*); die Zuordnung als solche
  ist richtig.
- geprüft, ohne Befund: **Risiko 2 trägt.** *Entfallen* mit Begründung, und die
  Begründung ist nachprüfbar: das Modul bricht fail-closed ab, wenn eine der
  beiden Listen fehlt, und die Prüfmenge ist mit `16 von 16` nicht leer.
- geprüft, ohne Befund: **Alle 16 Sensor-Dateien sind gedeckt** und keine
  Sensor-Datei nennt den alten Index noch als ihren Erwähnungs-Ort; die
  verbleibenden `AGENTS.md`-Nennungen in `harness/sensors/` betreffen §3.6, §5
  und §6 und sind unverändert richtig.

## Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 6 |
| MEDIUM | 5 |
| LOW | 4 |
| INFO | 1 |

**Finding-Klassen dieses Laufs:** Messung ohne reproduzierbares Instrument
(F-1) · Pfad-Form behauptet statt gemessen (F-2) · Kommentar beschreibt den
abgelösten Zustand (F-3) · Verweis auf den abgelösten Gate-Index nicht
nachgezogen (F-4, F-5, F-7 — drei Funde, **ein** Vorgang) · Zustandsfeld bei 3×
ohne einen der drei Ausgänge (F-6) · Register-Sichtung übergeht Einträge der
berührten Sub-Areas (F-8) · Risiko-Ausgang außerhalb der geschlossenen
Dreier-Menge (F-9) · Zusage reicht weiter als ihre Durchsetzung (F-10) ·
Lerneintrag als geschärfte Regel deklariert, ohne Zielort (F-11) ·
Finding-Klasse umformuliert statt zitiert (F-12) · Berührtes Verzeichnis ohne
Zeile in der Modus-Deklaration (F-13) · Abschließende Aufzählung neben einer
maschinenlesbaren Quelle (F-14) · Historische Aussage auf den neuen Zustand
umgeschrieben (F-15) · Tragende Form ohne Kennzeichnung an ihrer Stelle (F-16).

**Wiederkehrend gegenüber dem Vorgänger-Report:** *Lerneintrag als geschärfte
Regel deklariert, ohne Zielort* (dort F-8, hier F-11) und *Berührtes Verzeichnis
ohne Zeile in der Modus-Deklaration* (dort F-12 für `.harness/skills/`, hier
F-13 für `docs/plan/carveouts/`) — beide unmittelbar nach dem Lauf, der sie
benannt hat. F-1 trifft eine Kennung, die im Register bei 2× steht
(`messung-ohne-reproduzierbares-instrument`), F-9 eine zweite
(`risiko-eingetreten-und-im-slice-aufgeloest`), F-10 eine dritte
(`zusage-weiter-als-ihre-durchsetzung`) — in allen drei Fällen führt der Plan
den Eintrag als *bedient* oder nimmt ihn ausdrücklich vom Beleg aus. Ob daraus
je ein Ausgang wird, ist eine Planner-Entscheidung, kein Reviewer-Vorschlag.

## Verdikt

**Merge-blockierend:** ja — für F-1, F-2, F-4 und F-6.

**F-1 ist der schwerste Befund**, weil er nicht eine Nebenaussage trifft,
sondern den einen Fund, den der Slice als seine unerwartete Entdeckung führt:
Die entfallene Zellengrenzen-Regel meldete nach dem Wegfall der Tabelle
**nicht** grün, sondern `section-column-missing` mit Exit 2 — sie ist genau das
fail-closed, dessen Fehlen der Plan beklagt. Grün war ein Lauf ohne diesen
Prüfer, denn `doc-structure` hängt in `verify`, nicht in `gates`, und §3.4 nennt
seinen Aufruf nicht. Der falsche Satz steht dreifach: im Plan, in der
Closure-Notiz und dauerhaft als Kommentar in `.d-check.yml`; §9 stützt darauf
eine Register-Buchung. Vier Zeilen über der gestrichenen Stelle führte dieselbe
Datei den Befund-Code bereits — die Gegenprobe hätte einen Handgriff gekostet.

**F-2 und F-4 sind derselbe Vorgang an zwei Enden.** Der Slice hat den Index
bewegt und die Konfiguration nachgezogen, aber nicht die Sätze, die auf ihn
zeigen: `Makefile`-Kommentar und `make help`-Zeile behaupten weiter `AGENTS.md`
§4 als Erwähnungs-Ort, während `.d-check.yml` im selben Commit
`harness/README.md` sagt — und die neu geschriebene Grenze 4 nennt für den ADR-
Index eine Umweg-Form, die dort nicht auflöst (`../../` statt `../../../`,
`target-missing`). Beide Male ist es dieselbe Bewegung wie die, aus der der
Slice ausdrücklich gelernt haben will: eine Pfad-Aussage hinschreiben, statt sie
zu fahren. Der Geltungsbereich der Differenz-Messung war die Ursache — sie hielt
die zwei Indizes gegeneinander und sah die übrigen Verweisstellen im Repo
konstruktionsbedingt nicht.

**F-6** blockiert unabhängig davon: Der Register-Eintrag erreicht mit diesem
Slice 3× und verlässt die Closure ohne einen der drei Ausgänge. In einem Repo
ohne Wellen gibt es keinen „nächsten Lese-Schritt" außer der nächsten
Slice-Closure — und die liest einen Eintrag, dem der Ausgang dann seit einem
Vorgang fehlt.

**Was trägt:** Der Umzug selbst ist sauber und vollständig — beidseitig null,
unabhängig nachgerechnet —, die drei zitierten Mutations-Proben sind echt und
treffen ihren Gegenstand, und der `mentions`-Umweg ist in beide Richtungen
belegt, inklusive des Preises, den die Sensor-Datei benennt. Der neue
`AGENTS.md` §4 ist die Ziel-Form, ohne Auslassung.
