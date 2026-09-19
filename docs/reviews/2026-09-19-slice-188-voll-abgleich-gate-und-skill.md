# Review-Report: slice-188 — 2026-09-19

**Review-Art:** **unabhängiger Lauf** (frischer Kontext, nicht der Autor) über Plan + Diff —
geprüft gegen den Slice-Plan, die vendorte Ziel-Form,
den Vorgänger-Report und die berührten Repo-Artefakte (Modul 10 §Drei Review-Arten).
Kern-Behauptungen dieses Slice sind eine **Ausgangsmessung mit zwei Zählern** (§2),
**drei Übernahmen** in `.d-check.yml` (§3.1), ein **Nicht-Befund** für den Closure-Skill
(§3.2) und eine **Gegenprobe** (§3.3). Der Review rechnet die Zahlen nach, prüft jede
Übernahme gegen die Stelle, auf die sie zurückführt, und hält §3.3 gegen die Regel,
die der Slice selbst übernimmt.

**Gegenstand:** Commits `fd301de` (Lifecycle-Move `open` → `in-progress`) und `fde831c`
(„drei Ziel-Form-Uebernahmen in .d-check.yml"); Working Tree sauber, `HEAD` = `fde831c`.

**Skill:** `.harness/skills/reviewer.md` @ Stand `fde831c` — einschließlich §Mess-Regeln
(drei Regeln, `seit slice-193` alle drei verkörpert). Alle drei sind auf die Messungen
dieses Slice angewandt.
**Modell:** deepseek-v4.1-flash:cloud[1m] · **Datum:** 2026-09-19

> **Zitier-Form** *(Norm, nicht Ausfüll-Hinweis)*: Dieser Report friert ein; was er
> zitiert, bewegt sich weiter. Deshalb **Kennung, nicht Adresse** — `slice-NNN` statt
> seines Lifecycle-Pfads, `make <target>` statt eines Links auf die Sensor-Datei, eine
> Baseline-Stelle als Tag + Pfad in Inline-Code (`v6.6.0` ·
> `regelwerk/modul-05-planning-harness.md` §…). Ein `pfad`-Feld auf den **geprüften
> Gegenstand** hält den Stand des Laufs fest und darf das.

**Eingangs-Kontext:**

- `slice-188` (Plan, `in-progress/`, Stand `fde831c`), `slice-187` (Volltext **und**
  Review-Report, beide aus `slice-187-archiv.zip`), `slice-186`, `slice-185` (§Ausgangs-Slice)
- `AGENTS.md` §3.1–§3.7, §4, §5, §6; `harness/conventions.md` §Baseline,
  §Adaptions-Block, §Modus-Deklaration
- `.d-check.yml` vor (`fd301de`) und nach `fde831c`; `Makefile`, `d-check.mk`,
  `tools/slice-mv.sh`
- `v6.6.0` · `templates/.d-check.yml`, `templates/.harness/skills/closure-note-reviewer.template.md`;
  `v6.6.0` · `regelwerk/modul-05-planning-harness.md`, `regelwerk/modul-06-roadmap.md`,
  `regelwerk/modul-13-quality-gates.md` (Modulindex)
- `.harness/skills/closure-note-reviewer.md` (ganz gelesen)
- `v6.5.0` · `templates/.d-check.yml` — aus `git` zurückgeholt (`f49913a`), für die
  Vergleichszahl aus §2
- Beobachtungs-Register, namentlich `BEO-PLAN/kandidaten-klassifikation-groeber-als-der-kandidat`
  (verkörpert), `BEO-GATE/probe-liefert-den-gegenstand-mit` (verkörpert),
  `BEO-GATE/zusage-weiter-als-ihre-durchsetzung` (2×), `BEO-HARNESS/chronik-in-gelesenen-dateien`
  (3×, *geplant*), `BEO-HARNESS/aggregat-aufzaehlung-hinkt-dem-makefile-hinterher` (2×),
  `BEO-GATE/muster-trifft-nur-die-haeufige-schreibweise` (2×),
  `BEO-GATE/versionsangabe-neben-digest-ungeprueft` (verkörpert),
  `BEO-GATE/ruhe-marker-ungewaechtert` (verkörpert, „Ehemals: `BEO-014`")
- früherer Report zum selben Bereich: der zu `slice-187`; seine fünfzehn Findings sind
  hier auf Wiederholung geprüft

**Geltungsbereich der eigenen Messungen** (Skill §Mess-Regeln, erste Regel):

- **Das Instrument aus `slice-187` §2 ist nachgebaut und ausgeführt** (Parameter wie
  deklariert: Fenster 6, Mindestlänge 6 Wörter; Fenster 8 als Gegenlauf), gegen `HEAD`
  und gegen den Stand vor der Arbeit (`fde831c^`). Alle **sechs** Zahlen der §§2/3.4
  reprodzieren exakt; die Vergleichszahl **42** gegen den `v6.5.0`-Stand ebenfalls —
  deren Korpus ist nicht mehr im Baum, aber aus `git` (`f49913a`) rekonstruierbar.
- **Zweite Zähler, anders gebaut:** die Modul-Menge in §3.3 gegen einen
  YAML-Parser-Lauf (Schlüssel und Mengenoperationen statt `grep`), die Abschnittszahl in
  §3.2 gegen zwei verschiedene Muster (`^## ` und `^#`), die Delta-Zuordnung in §2 gegen
  einen Zeilennummern-Vergleich der beiden Vorlagen-Stände. Jede Abweichung steht als
  Finding, jede Übereinstimmung als Negativbefund.
- **Gelesen, nicht gemessen:** die 25 Kandidaten des Skill-Paares und die 58 des
  `.d-check.yml`-Paares sind von mir einzeln gelesen; *„Umformulierung"* ist dabei ein
  Urteil über zwei Formulierungen, kein Lauf.
- **Nicht ausgeführt:** `make ci`, `make doc-immutable`, `make trace-check`,
  `make commit-scope-check`, `make doc-reviews`, `make archive-wave-test`. Die
  Gate-Ergebnisse unten decken `gates` und `verify` und damit die Aggregate;
  die Einzel-Targets darüber hinaus sind nicht belegt.
- **Die Gegenrichtung des Abgleichs** (Repo → Vorlage, also *„trägt das Repo Aussagen,
  die falsch sind"*) habe ich nur punktuell geprüft — das Instrument sieht sie
  prinzipiell nicht (Grenze 3, `slice-187` §2).

---

## Findings

### F-1 — Die übernommene Aktivierungs-Regel steht ausschließlich im Konfigurations-Kommentar

- `kategorie`: HIGH
- `quelle`: Reviewer-Skill §Klassifikation, HIGH *Norm nur im Template-Kommentar*, zweite
  (hiesige) Ausprägung — *„[`.d-check.yml`] trägt die Begründung jeder Regel
  ausschließlich in Kommentaren. Das ist zulässig — aber eine **Zusage** darf dort nicht
  allein stehen; sie gehört in `AGENTS.md` §4 oder eine Sensor-Datei, die der Lauf
  liest."*; die Regel selbst trägt ihren Grund aus `v6.6.0` ·
  `regelwerk/modul-13-quality-gates.md` §Kernidee (Gate-Lüge)
- `pfad`: `.d-check.yml`:9–17
- `befund`: Die übernommene Regel *„AKTIVIEREN HEISST ZWEI SCHRITTE"* steht nach diesem
  Diff nur in `.d-check.yml`; ein repo-weiter `grep` auf die Wendung (ohne
  `.harness/baseline/`) findet sie sonst nur im Slice-Plan. `AGENTS.md` §4 sagt zu
  Modulen nichts, `harness/README.md` §Sensors und `harness/sensors/gate-consistency.md`
  führen nur die `modules`-Liste, und dieser Sensor benennt seine eigene Grenze
  (*„misst Präsenz und Abwesenheit von Namen, nicht Wirkung"*). Der Kommentar selbst
  trägt daneben die Zustandsaussage *„GEMESSEN, slice-188: … kein stummer Block"*, die
  kein Lauf nachhält.
- `verifizierbar`: ja — `grep -rl "ZWEI SCHRITTE" --exclude-dir=.harness .` (ein Treffer).
- `klasse`: Zusage steht nur im Konfigurations-Kommentar

### F-2 — Der neue Kopfkommentar beschreibt den abwesenden Text

- `kategorie`: HIGH
- `quelle`: `AGENTS.md` §3.7 (*„Falsch: abwesenden Text beschreiben … Richtig: die
  geltende Zusage nennen; die vorige hält `git`"*); Reviewer-Skill §Klassifikation, HIGH
  *Kommentar trägt keine der Kommentar-Klassen*; Register-Eintrag
  `BEO-HARNESS/chronik-in-gelesenen-dateien` (3×, Ausgang *geplant*)
- `pfad`: `.d-check.yml`:6–7
- `befund`: Der Satz *„… und diese Zeile trug bis slice-188 einen Stand, der vier
  Migrationen zurueckliegt"* erzählt den früheren Zustand des Textes, in dem er steht —
  im Kommentar einer Datei, die jeder Gate-Lauf mitliest. `git show fde831c` trägt
  dieselbe Auskunft. Die Schreibweise ist eine **vierte** gegenüber den drei Phrasen,
  die `slice-191` als Sensor bauen will (*„Bis slice-NNN stand hier"*, *„bis slice-NNN
  tat das"*, *„Vorher stand hier"*) — benannt, nicht gezählt, weil der Sensor noch
  nicht existiert.
- `verifizierbar`: ja — die Zeile steht als `+`-Zeile in `git show fde831c -- .d-check.yml`.
- `klasse`: Chronik in gelesenen Dateien

### F-3 — Der Kopf sagt „NICHT hier", derselbe Diff schreibt den Stand zweimal hinein

- `kategorie`: MEDIUM
- `quelle`: der Vorgang gegen seinen eigenen Befund 3 (`slice-188` §3.1: *„Die Antwort
  ist kein zweiter Pin, sondern **kein Pin**"*); `harness/conventions.md` §Baseline als
  der eine Ort
- `pfad`: `.d-check.yml`:5–7 gegen `.d-check.yml`:9, :182, :188
- `befund`: Nach dem Diff steht in derselben Datei, der adoptierte Baseline-Stand stehe
  *„NICHT hier"* — und `(Ziel-Form v6.6.0)` steht zweimal darin (:9, :182), dazu eine
  dritte Nennung `(Baseline v6.6.0, …)` aus `slice-193` (:188), die der Slice
  unangetastet lässt. Nach der eigenen Analyse in §3.1 fängt kein Lauf eine nackte
  Stand-Kennung; beim nächsten Sprung altern die drei Stellen so still wie die entfernte
  Zeile.
- `verifizierbar`: ja — `grep -n "v6\.6\.0" .d-check.yml`; `git log -S 'v6.6.0' -- .d-check.yml`
  liefert genau zwei Commits.
- `klasse`: Stand-Kennung dupliziert, ohne Wächter

### F-4 — Die Abschnittszahl des Skill-Vergleichs ist um eins zu hoch

- `kategorie`: MEDIUM
- `quelle`: Reviewer-Skill §Mess-Regeln, *Wer eine Menge zählt, zählt sie zweimal
  verschieden* (`seit slice-193`, verkörpert); Register-Eintrag
  `BEO-PLAN/kandidaten-klassifikation-groeber-als-der-kandidat`, zweite Ausprägung
  (*„die Zahl stammt aus der falschen Quelle"*)
- `pfad`: `docs/plan/planning/in-progress/slice-188-voll-abgleich-gate-und-skill.md`:113–116
- `befund`: §3.2 begründet den einzigen *ohne Befund*-Ausgang des Slice mit *„dieselben
  **sieben** `##`-Abschnitte in derselben Reihenfolge"*; beide Dateien führen **sechs**
  (`grep -c '^## '` → 6 und 6). Sieben ergibt erst das Muster `^#`, das die Titelzeile
  mitzählt (7 und 7). Dieselbe Quellen-Verschiebung wie in `slice-186` (*„alle fünf
  Abschnitte"* gegen sechs), am selben Gegenstand.
- `verifizierbar`: ja — `grep -c '^## '` und `grep -c '^#'` auf beide Dateien.
- `klasse`: Zahl aus der falschen Quelle, nicht gegengezählt

### F-5 — Die 22 zusätzlichen Kandidaten kommen aus zwei neuen Blöcken, nicht aus einem

- `kategorie`: MEDIUM
- `quelle`: Reviewer-Skill §Mess-Regeln, erste Regel (Geltungsbereich) und dritte Regel
  (zweite Zählung); Register-Eintrag `BEO-PLAN/kandidaten-klassifikation-groeber-als-der-kandidat`
- `pfad`: `…slice-188-….md`:74–77 (§2)
- `befund`: §2 schreibt die **22** zusätzlichen Kandidaten dem neuen `targets`-Block der
  Ziel-Form zu („den `slice-193` inhaltlich bereits umgesetzt hat"). Der Delta des
  Baseline-Sprungs besteht aus **zwei** Blöcken — `v6.5.0` → `v6.6.0` der Vorlage fügt
  +24 Zeilen hinzu, Kopf `5–9` und `targets` `20–38`. Nachgezählt liegen **5** der 22 in
  den Kopfzeilen 5–9, **17** im `targets`-Block; die fünf gehören dem Block, den dieser
  Slice als Befund 1 führt.
- `verifizierbar`: ja — Vorlagenstände aus `git` vergleichen und die Kandidatenzeilen
  gegen die hinzugefügten Zeilen halten.
- `klasse`: Zahl aus der falschen Quelle, nicht gegengezählt

### F-6 — Die ausgegebene Gegenprobe ist nicht die, die die übernommene Regel verlangt

- `kategorie`: MEDIUM
- `quelle`: Reviewer-Skill §Mess-Regeln, *Eine Mutations-Probe belegt erst, wenn sie rot
  war* (verkörpert); Register-Eintrag `BEO-GATE/probe-liefert-den-gegenstand-mit`
  (verkörpert); die übernommene Regel selbst (`.d-check.yml`:14–15)
- `pfad`: `…slice-188-….md`:131–141 (§3.3) gegen `.d-check.yml`:14–15
- `befund`: §3.3 sagt *„Die übernommene Regel verlangt sie, also ist sie gefahren"* und
  legt eine Aufzählung vor, welcher Block in `modules` oder in `--enable` steht. Die
  Regel, die derselbe Slice übernimmt, verlangt eine andere Probe: *„einen Verstoss der
  Klasse einbauen und den Befund sehen, nicht den Exit-Code"* — und §7 hält fest, dass
  kein Modul ein-, aus- oder umgeschaltet wurde, die verlangte Probe also keinen
  Gegenstand hätte. Die Substanz von §3.3 hält (siehe Negativbefunde); die als
  *gefordert und gefahren* ausgegebene Hälfte ist eine andere Messung.
- `verifizierbar`: ja — der Regeltext in `.d-check.yml`:14–15 neben der Tabelle in §3.3.
- `klasse`: Probe trifft ihren Gegenstand nicht

### F-7 — §9 trägt weiterhin den Ausfüll-Platzhalter, obwohl der Übergang erfolgt ist

- `kategorie`: MEDIUM
- `quelle`: `v6.6.0` · `regelwerk/modul-05-planning-harness.md` §Zwei Schritte vor der
  Modus-Begründung (*„Sie hängen weder am Modus noch am Slice-Typ und stehen deshalb in
  **jedem** Slice-Plan"*); `regelwerk/modul-06-roadmap.md` §Wann Arbeit eine Welle
  braucht, Tabelle *Träger im Repo ohne Wellen* (Sichtungs-Schritt: Slice-**Planung**)
- `pfad`: `….md`:212–218 (§9)
- `befund`: Beide Pflichtblöcke stehen unausgefüllt (*„beim Übergang nach `in-progress/`
  auszufüllen"*), obwohl der Übergang mit `fd301de` erfolgt und die Arbeit mit `fde831c`
  gelaufen ist; der Slice ist wellenlos, der Sichtungs-Schritt damit der **einzige**
  Leser für alles, was im Register unter der Schwelle steht. `slice-187` §9 hat ihm die
  verschärfte Form ausdrücklich übergeben (Register nach **allen** berührten Kürzeln,
  Satz „steht in §9 der drei Folge-Slices"). Kein Gate prüft §9-Inhalt.
- `verifizierbar`: ja — die Platzhalter sind Dateiinhalt; ob der Schritt *unausgeführt*
  blieb, ist ein Urteil.
- `klasse`: Vorgelagerter Pflicht-Block bleibt Platzhalter

### F-8 — Die Mengen-Aussage in §3.3 ist breiter als die zwei Listen, die sie zeigen

- `kategorie`: LOW
- `quelle`: der Slice selbst (§3.3 beruft sich auf die dritte Mess-Regel)
- `pfad`: `….md`:133–141
- `befund`: *„**Jeder** konfigurierte Top-Level-Block … steht entweder in `modules` oder
  wird per `--enable` geschaltet."* Der YAML-Lauf findet **15** Top-Level-Schlüssel;
  `ignore-refs` (Option des aktiven `links`) steht in keiner der beiden Zeilen, dazu die
  Globals `scan` und `modules`. Und *„entweder … oder"* hat eine Ausnahme, die die
  Tabelle selbst zeigt: `reviews` steht in **beiden** Zeilen — die „zwei verschiedenen
  Quellen" überschneiden sich um ein Element.
- `verifizierbar`: ja — die beiden Zahlenreihen in §3.3 gegen `yaml.safe_load` bzw.
  `grep -oE -- '--enable [a-z-]+'`.
- `klasse`: Sammelaussage breiter als die Menge, die die Tabelle zeigt

### F-9 — Die `Aktiv:`-Zeile nennt fünf der acht Module, die in `modules` stehen

- `kategorie`: LOW
- `quelle`: `harness/conventions.md` §Modus-Deklaration (der Zustand ist das Artefakt,
  nicht die Aufzählung); Register-Kandidat
  `BEO-HARNESS/aggregat-aufzaehlung-hinkt-dem-makefile-hinterher` (2×)
- `pfad`: `.d-check.yml`:19–26 gegen `.d-check.yml`:34
- `befund`: *„Aktiv: links, anchors, ids …, matrix …, spans …"* — `modules:` führt acht
  (`+ hostpaths`, `+ versions`, `+ reviews`), und §3.3 des Plans listet für dieselbe
  Frage acht. Die Zeile ist älter als dieser Diff; der Slice hat den Kopf aber bearbeitet
  und für den Modul-Zustand eine Messung erklärt, ohne die Aufzählung gegen `modules` zu
  halten.
- `verifizierbar`: ja — die Zeile und die `modules:`-Zeile derselben Datei nebeneinander.
- `klasse`: Aufzählung hinkt der Menge hinterher, die sie abschließend nennt

### F-10 — Der „zweite, anders gebaute Zähler" ist derselbe Bau mit anderem Parameter

- `kategorie`: LOW
- `quelle`: Reviewer-Skill §Mess-Regeln, *Wer eine Menge zählt, zählt sie zweimal
  verschieden* (verkörpert)
- `pfad`: `….md`:59–67 (§2)
- `befund`: §2 nennt den Fenster-8-Lauf einen *„zweiten, anders gebauten Zähler"* und
  beschreibt ihn zwei Zeilen tiefer als *„derselbe Lauf mit Fenster 8"*. Es ist dasselbe
  Skript mit einer geänderten Konstanten — die Zeile-weise-`any(...)`-Konstruktion bleibt
  dieselbe. Der Ertrag ist real (er macht halb gedeckte Zeilen sichtbar) und reproduziert
  (s. Negativbefunde), aber es ist keine unabhängige Zählung.
- `verifizierbar`: ja — der Fenster-Parameter gegen den Code des Instruments.
- `klasse`: Zweiter Zähler ist derselbe Bau mit anderem Parameter

### F-11 — Der Lifecycle-Commit schreibt die Streichung des Ruhe-Markers dem Werkzeug zu

- `kategorie`: LOW
- `quelle`: `tools/slice-mv.sh` §NICHT BEHANDELT (1) *SEMANTIK* (*„Das Werkzeug zieht
  PFADE nach, keine Aussagen"*); `AGENTS.md` §5 (Lifecycle); Wiederholung der Klasse aus
  dem `slice-187`-Report (dort F-13)
- `pfad`: Commit `fd301de` (Betreff) gegen `docs/plan/planning/in-progress/roadmap.md`
- `befund`: Der Betreff lautet *„slice-188 open -> in-progress (make slice-mv)"*; der
  Commit streicht neben dem Move den Ruhe-Marker *„Nichts in Arbeit."*. Ein `grep` auf
  `tools/slice-mv.sh` findet weder „Marker" noch den Satz — das Werkzeug fasst ihn nicht
  an. Dieselbe Zuschreibung wie in `slice-187`, dort als F-13 vermerkt; ein
  Register-Eintrag dafür existiert nicht.
- `verifizierbar`: ja — `git show fd301de` gegen `grep -n "Marker" tools/slice-mv.sh`.
- `klasse`: Werkzeug-Zuschreibung für eine Leistung, die das Werkzeug nicht erbringt

### F-12 — Das Zitat der entfernten Kopfzeile trägt eine Form, die sie dort nicht hatte

- `kategorie`: INFO
- `quelle`: `slice-188` §3.1, Zeile 3 der Tabelle
- `pfad`: `….md`:94 gegen `fde831c^:.d-check.yml`:4
- `befund`: Die entfernte Zeile wird als *„Baseline v3.5.2 (vendored, `MR-006` als
  Link)"* zitiert; in der Datei stand `MR-006` **nackt**, ohne Klammer-Ziel. Die
  Link-Form erzwingt die Kennungs-Linkpflicht im Markdown-Dokument — wer die Zeile im
  zitierten Original sucht, findet sie nicht in dieser Gestalt.
- `verifizierbar`: ja — `git show fde831c -- .d-check.yml`, Zeile 4 der `-`-Seite.
- `klasse`: Zitat trägt eine Form, die der zitierte Text nicht hatte

## Negativbefunde

- geprüft, ohne Befund: **Die sechs Zahlen der §§2/3.4 reproduzieren exakt.** Instrument
  aus `slice-187` §2 nachgebaut, unverändert ausgeführt: `.d-check.yml` **64 → 58**
  (Fenster 6), **70 → 67** (Fenster 8); `closure-note-reviewer.md` **25** und **34**
  (vor und nach der Arbeit gleich, die Datei ist nicht im Diff). Die Vergleichszahlen aus
  §2 halten ebenfalls: **42** gegen den `v6.5.0`-Stand (Vorlage aus `f49913a`
  zurückgeholt) und **25** für den Skill — dessen Vorlage hat der Baseline-Sprung
  byte-identisch mitgenommen (`git show 012530f -M` → 0 Zeilen). Kein Kandidat ist
  verschwunden und keiner neu entstanden; die Nettodifferenz **+22** entspricht genau den
  Kandidaten in den hinzugefügten Vorlagenzeilen.
- geprüft, ohne Befund: **Der Kern von §3.3 hält — gegen einen anders gebauten Zähler.**
  YAML-Parser-Lauf über 15 Top-Level-Schlüssel, Mengen gegen `modules` und `--enable`:
  **kein** Modul-Block ist konfiguriert und ungeschaltet; die im Plan genannten neun
  `--enable`-Module stimmen mit `Makefile`/`d-check.mk` überein (`commits`, `mentions`,
  `planning`, `reviews`, `structure`, `targets`, `tracked`, `vcs`, `workflows`), die acht
  `modules`-Einträge ebenfalls. Was die Aussage *breiter* macht, steht in F-8.
- geprüft, ohne Befund: **Übernahme 1 trägt ihren Grund.** Der übernommene Block nennt
  die Gate-Lüge und `Modul 13`; die a-check-Zusätze (der `--enable`-Weg, die zwei
  belegten Fälle `targets/slice-074` und `planning/BEO-014`) sind Ausprägung des Repos,
  keine Nachschrift der Baseline-Begründung. Beide Vorfallangaben sind durch die Datei
  selbst gedeckt (`slice-074` als Stub in `done/welle-12/`, `BEO-014` = *„Ehemals"*-Form
  von `BEO-GATE/ruhe-marker-ungewaechtert`, beide Kommentare im Bestand).
- geprüft, ohne Befund: **Übernahme 2 sitzt an der Stelle, auf die sie zurückführt.**
  Die Regel *„NAMENTLICH, nie als Glob"* samt Grund steht vor `exempt-targets`, und die
  zwei Kandidatenzeilen, die die Zahl 64 → 58 senken, sind genau die beiden
  Vorlagenzeilen dieser Regel (33, 34) plus vier Kopfzeilen.
- geprüft, ohne Befund: **§3.2 hat recht, wo es nicht zählt.** Unabhängig nachgesehen:
  beide Dateien führen dieselben sechs `##`-Abschnitte in derselben Reihenfolge, dieselben
  vier Schweregrade mit denselben Definitionen, die Negativbefund-Pflicht und dasselbe
  fünffeldrige Output-Schema; die Abweichung im Feld `quelle` ist die im Kopf der Datei
  deklarierte (keine Closure-Note-ADR, kein Python-Tooling). Die falsche Zahl steht in F-4.
- geprüft, ohne Befund: **Die 25 Kandidaten sind Umformulierungen — im gelesenen
  Einzelfall.** Jeder der 25 Sätze hat im Repo-Gegenstück eine Trägerstelle (Kontext-Eingang
  an `AGENTS.md` §5 statt an Slice-Template/ADR, HIGH-Beispiele mit einem ausgetauschten,
  MEDIUM-Zusatz *Futur-Aussage*, `AGENTS.md §5` statt `<ADR-NNNN>` im Schema). Das ist
  mein Urteil über zwei Formulierungen, kein Lauf — die Aussage der Kandidaten bleibt
  eine Vorauswahl.
- geprüft, ohne Befund: **Keine Übernahme ist eine Gate-Lockerung** (`AGENTS.md` §3.6).
  `git show fde831c -- .d-check.yml` enthält ausschließlich `#`-Zeilen; kein Modul wurde
  ein-, aus- oder umgeschaltet, `modules:` und `--enable` sind unverändert, `Makefile`
  und `d-check.mk` sind nicht im Diff. Der Risiko-Ausgang in §7 (*„kein aktives Modul
  umgeschaltet"*) ist damit belegt.
- geprüft, ohne Befund: **Referenz-Richtung, Form und Kennungen.** `.d-check.yml` liegt in
  keiner Klasse des Moduls `matrix`; kein Spec-Stratum und keine ADR ist berührt; die
  zitierten Kennungen lösen auf (`slice-074`, `slice-188`, `MR-006`, `BEO-014`,
  `Modul 13` im vendorten Index). Keine erfundene Kennung, kein Abwärtsverweis.
- geprüft, ohne Befund: **Risiko-Ausgänge — Form.** Zwei notierte Risiken, zwei Ausgänge,
  beide aus der geschlossenen Dreier-Menge und beide mit Begründung. Der zweite
  („Closure-Skill unverändert") ist gegen den Diff belegt: `.harness/skills/` kommt in
  `fde831c` nicht vor.
- geprüft, ohne Befund: **Commit-Hygiene.** `fde831c` berührt genau zwei Dateien
  (`.d-check.yml`, der Slice-Plan in `in-progress/`) und nennt `slice-188`; `fd301de` ist
  der Lifecycle-Move samt Verweis-Nachzug auf `slice-189`. Keine ADR angefasst, kein
  Produkt-Code, keine Suppression, keine Host-Toolchain. Zum Betreff von `fd301de`: F-11.
- geprüft, ohne Befund: **Der Ruhe-Marker und die Roadmap.** `in-progress/` trägt genau
  einen Slice (WIP-Limit 1), der Ruhe-Marker ist gefallen, `make doc-planning` grün — die
  Richtung, die `modul-06` §Roadmap-Struktur verlangt.
- geprüft, ohne Befund: **Der Plan-Zuschnitt ist eingehalten.** §1 schließt die zwei
  Spec-Straten, die neuen/geänderten Module und die elf Instanz-Vorlagen mit je einer
  Begründung aus; keine Berührung außerhalb der dort genannten Paare; die DoD trägt
  höchstens drei Liefer-Punkte nach der Ignore-Regel des Struktur-Gates. Einzige
  Plan-Lücke: §9 — F-7.
- **Nicht geprüft, und darum keine Aussage:** Die Gegenrichtung des Abgleichs über die
  volle Länge der zwei Gegenstücke; die Zahlen der §§2/3.4 unter einem *anderen*
  Parameter als 6 und 8 (eine dritte Fenstergröße habe ich nicht gefahren); `make ci`,
  `make doc-immutable`, `make trace-check`, `make commit-scope-check`, `make doc-reviews`;
  die CI-Lage.

## Ergebnis der Gate-Läufe (ausgeführt am 2026-09-19, Stand `fde831c`, Working Tree sauber)

| Target | Exit | Geltungsbereich der Aussage |
|---|---|---|
| `make doc-check` | 0 | 573 Dateien, 0 Befunde — Link-, Anker-, Kennungs- und Matrix-Prüfung der Repo-Doku, hermetisch |
| `make doc-structure` | 0 | Struktur-Invarianten (DoD-Größe, Closure-Struktur, Kopffelder, AC-Form, Zellengrenzen) |
| `make doc-targets` | 0 | Deklarations-Konsistenz Gate-Index ↔ Makefile-Ziele |
| `make doc-planning` | 0 | Roadmap ↔ `in-progress/` (Ruhe-Marker-Hälfte) |
| `make version-coherence` | 0 | 3 gepinnte SHAs mit einheitlichem Tag-Kommentar, 3 doppelt deklarierte Variablen — **Divergenz**, nicht Wahrheit |
| `make verify` | 0 | 21 Anforderungen, 0 Waisen; DoD-/Closure-Schicht |
| `make gates` | 0 | Aggregat (Code-Fragen) inkl. `record-gates` |

Kein Gate hat einen der Findings oben gemeldet; F-1, F-2, F-4 und F-5 sind per Konstruktion
gate-frei (Kommentar-Klassen, Zahl zwischen zwei Dateien).

## Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 2 |
| MEDIUM | 5 |
| LOW | 4 |
| INFO | 1 |

**Finding-Klassen dieses Laufs:** Zusage steht nur im Konfigurations-Kommentar · Chronik in
gelesenen Dateien · Stand-Kennung dupliziert, ohne Wächter · Zahl aus der falschen Quelle,
nicht gegengezählt · Probe trifft ihren Gegenstand nicht · Vorgelagerter Pflicht-Block
bleibt Platzhalter · Sammelaussage breiter als die Menge, die die Tabelle zeigt ·
Aufzählung hinkt der Menge hinterher, die sie abschließend nennt · Zweiter Zähler ist
derselbe Bau mit anderem Parameter · Werkzeug-Zuschreibung für eine Leistung, die das
Werkzeug nicht erbringt · Zitat trägt eine Form, die der zitierte Text nicht hatte

**Zähler-Hinweis, mit der Regel dazu** (drei Klassen, drei verschiedene Lagen):

1. *Chronik in gelesenen Dateien* (F-2) — `BEO-HARNESS/chronik-in-gelesenen-dateien` steht
   bei **3×** mit Ausgang *geplant* (`slice-191`); der Eintrag ist also zugewiesen und
   zählt nicht weiter. Was der neue Fall beiträgt, ist eine **Kalibrierungs-Auskunft**:
   die vierte Schreibweise gehört als *benannt, nicht gezählt* in den Eintrag, denn die
   drei geplanten Phrasen treffen sie nicht.
2. *Zahl aus der falschen Quelle* (F-4, F-5) und *Probe trifft ihren Gegenstand nicht*
   (F-6) sind **verkörperte** Klassen (`BEO-PLAN/kandidaten-klassifikation-groeber-als-der-kandidat`
   und `BEO-GATE/probe-liefert-den-gegenstand-mit`, beide `seit slice-193` bzw.
   `seit slice-181` in Briefing und Skill). Sie bewegen keinen Zähler; sie zeigen, dass
   die verkörperte Regel an **diesem** Diff nicht angewandt wurde — F-4 genau dort, wo
   §2 die Regel selbst zitiert.
3. *Zusage steht nur im Konfigurations-Kommentar* (F-1) fällt unter die HIGH-Kategorie
   des Reviewer-Skills; der Register-Eintrag `BEO-GATE/zusage-weiter-als-ihre-durchsetzung`
   (2×) grenzt laut eigenem `state.md` die **Nachbarklasse** ab („eine Zusage behauptet
   eine Prüfung, und der Prüfer prüft weniger") und sagt selbst, die HIGH-Kategorie decke
   sie mit ab. F-1 zählt den Eintrag damit nicht hoch.
4. *Werkzeug-Zuschreibung* (F-11) hat **keinen** Register-Eintrag, obwohl `slice-187`
   dieselbe Klasse schon einmal ausgewiesen hat (dort F-13) — zweites Auftreten in Folge,
   dritte Gelegenheit wäre eine Lücke.

**Wiederkehrend gegenüber dem Vorgänger-Report:** *Chronik in gelesenen Dateien* (dort
F-3, hier F-2) · *Probe trifft ihren Gegenstand nicht* (dort F-7, hier F-6) ·
*Kandidaten-Klassifikation gröber als der Kandidat* (dort F-12, hier F-4/F-5) ·
*Werkzeug-Zuschreibung* (dort F-13, hier F-11). **Erledigt sind zwei:** die
Ausgangsmessung ist erneut reproduzierbar, und der Zuschnitt ist nicht gewachsen. **Neu
und unangenehm:** die zweite der zwei Warnungen, die `slice-187` §9 diesem Slice
ausdrücklich übergeben hat (*„Zusage in der Doku, Grenze nur im Konfigurations-Kommentar"*),
ist in diesem Diff in der Nachbarform eingetreten (F-1) — die Handover-Senke (§9) ist
dieselbe Stelle, an der dieser Slice seinen Platzhalter stehen ließ (F-7).

## Verdikt

**Merge-blockierend: ja**, für F-1 und F-2 — beide liegen in `.d-check.yml`, beide in
Zeilen, die dieser Diff selbst geschrieben hat, und beide in Klassen, die der Reviewer-Skill
als HIGH führt und die **kein Gate** fängt. Sie sind in wenigen Zeilen behebbar; F-1, weil
die Regel eine zweite Stelle braucht, an der ein Lauf sie liest, F-2, weil der Satz die
Regel nicht trägt, sondern ihre Vorgeschichte.

**Die fünf MEDIUM sind je einzeln klein und einzeln behebbar**, zwei davon sind aber
tragende Belege: F-4 stützt den **einzigen** *ohne Befund*-Ausgang des Slice und ist um
eins falsch; F-6 erklärt eine Probe als gefahren, die die übernommene Regel nicht
verlangt und die ohne Gegenstand geblieben wäre. F-7 ist eine Plan-Lücke, kein
Artefakt-Fehler, und blockiert nur, weil der Sichtungs-Schritt in einem wellenlosen
Repo der einzige Leser unterhalb der Schwelle ist.

**Was ausdrücklich trägt:** Die Ausgangsmessung ist auch in diesem Slice nachrechenbar —
alle sechs Kandidatenzahlen und beide Vergleichszahlen reproduzieren, und die
Nullbewegung beim Skill ist richtig gedeutet, statt als Fortschritt gelesen. Die drei
Übernahmen hängen an echten Lücken: die Aktivierungs-Regel fehlte, die
*„namentlich"*-Regel fehlte, und der Kopf nannte einen Stand, der fünf Migrationen alt
war. Die Abgrenzung von Normtext und Bedienhinweis ist die richtige Prüf-Ebene und hält
für die gelesenen 64 bzw. 25 Kandidaten. Und der Slice hat sich der schärfsten
Selbstprüfung unterzogen, die der Vorgänger hinterlassen hat: zwei Zähler statt einem.

**Übergabe:** an die Implementer-Rolle (Rückkante Review → Plan für F-7, das den Plan
betrifft). Die Finding-Klassen gehen zusätzlich in die Slice-Closure §7 und von dort in
den Zähler. Dieser Report ist **Lauf-Beleg**, keine Verifikation — DoD- und
Spec-Konformität prüft der Verifier separat (Modul 11).
