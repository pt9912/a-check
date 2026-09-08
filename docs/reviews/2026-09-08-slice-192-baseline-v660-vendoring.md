# Review-Report: slice-192 — 2026-09-08

**Review-Art:** Plan — geprüft gegen den Slice-Plan, gegen die Baseline-Regeln,
die er in Anspruch nimmt, und gegen die Tat selbst (Modul 10 §Drei Review-Arten).
Kern-Behauptungen sind der Bau-Weg des Bundles samt Gegenprobe, die Pin-Deckung,
die Wortgleichheit der fünf `MR`-Zeiger, die Delta-Analyse und die
Sub-Area-Sichtung.

**Gegenstand:** Commit-Range `0b97c5d..HEAD` (`28d647b`) — `012530f`
(„Baseline auf v6.6.0 gehoben") und `28d647b` („slice-192 Umsetzung,
Risiko-Ausgaenge und Closure"); der Übergangs-Commit `0b97c5d` ist die Basis.
Mitgelesen: `d410675` (Zitier-Form für Planungs-Dokumente, `seit slice-192`).

**Skill:** `.harness/skills/reviewer.md` @ `28d647b` (in dieser Range selbst
gebumpt, inhaltlich unverändert) · <!-- d-check:ignore -->
**Modell:** claude-opus-5[1m] · **Datum:** 2026-09-08

**Eingangs-Kontext** (die Verträge, gegen die geprüft wurde):

- slice-192 (Plan, `in-progress/`), slice-193, slice-188, slice-189 (`open/`)
- `harness/conventions.md` §Baseline, §Adaptions-Block, §Modus-Deklaration pro
  Sub-Area; die fünf aktiven `MR`-Dateien mit Baseline-Zeiger und die zwei
  aufgelösten MR-011 / MR-021
- `AGENTS.md` §3 (Hard Rules), §4, §5, §6
- AC-QA-02 und AC-QA-03 als die vom Plan bezogenen Kennungen
- `v6.6.0` · `regelwerk/modul-05-planning-harness.md` §Offene Risiken werden bei
  Closure aufgelöst · `regelwerk/modul-06-roadmap.md` §Das Beobachtungs-Register
  und §Roadmap-Struktur: fünf Abschnitte · `regelwerk/modul-13-quality-gates.md`
  §Hard Rule · `regelwerk/grundlagen-harness-dateien.md`
  §harness/README.md als Einstiegspunkt
- Beobachtungs-Register, gelesen: alle 63 Verzeichnisse (`state.md` und
  Beleg-Zahl), im Detail die sechs in §9 genannten sowie
  `BEO-PLAN/messung-ohne-reproduzierbares-instrument`,
  `BEO-PLAN/vollstaendigkeits-haken-ohne-erschoepften-gegenstand`,
  `BEO-PLAN/zielsatz-nach-plan-aenderung-nicht-nachgezogen`,
  `BEO-PLAN/kandidaten-klassifikation-groeber-als-der-kandidat`,
  `BEO-GATE/attestierung-vor-dem-vorgang`,
  `BEO-HARNESS/harness-einstieg-ohne-modus-zeile`
- Report zu slice-187 (aus `done/slice-187-archiv.zip` gelesen), alle 15
  `klasse`-Felder und §Summary
- Kurs-Repo, Tags `v6.5.0` und `v6.6.0` — **nicht** dem Zitat des Plans
  vertraut: beide Bundles aus `git archive <tag>` in ein Wegwerf-Verzeichnis neu
  gebaut und gegen den vendorten Baum bzw. gegen `git show 0b97c5d:…` verglichen

---

## Findings

### F-1 — Fünf historische Nennungen des alten Standes sind mitgehoben; §2 nennt damit einen leeren Diff als Instrument

- `kategorie`: HIGH
- `quelle`: nachweislich falsche Tatsachenbehauptung (Reviewer-Skill
  §Klassifikation); `AGENTS.md` §5, Zitier-Form-Punkt, Absatz zu
  Planungs-Dokumenten (`seit slice-192`)
- `pfad`: `docs/plan/planning/in-progress/slice-192-baseline-v660-vendoring.md`
  Zeilen 22, 47, 49, 70/71, 165
- `befund`: Bei `0b97c5d` trugen diese fünf Stellen den alten Stand — Zielsatz
  *„`v6.5.0` ist entfernt"*, Abschnittstitel *„Delta-Analyse `v6.5.0` →
  `v6.6.0`"*, der Mess-Aufruf `git diff v6.5.0 v6.6.0 -- lab/regelwerk
  lab/templates`, die Herkunft des von slice-187 übernommenen Satzes und der
  gleichlautende DoD-Punkt; `28d647b` ersetzt sie durch den neuen Stand. §1 und
  der **abgehakte** DoD-Punkt behaupten jetzt, der soeben adoptierte Stand sei
  entfernt worden, der Abschnittstitel vergleicht einen Tag mit sich selbst, und
  der genannte Aufruf liefert leere Ausgabe, während unmittelbar daneben
  *„10 Dateien, +87/−32"* als sein Ergebnis steht.
- `verifizierbar`: ja — `git show 0b97c5d:<plan>` gegen `HEAD`; im Kurs-Klon
  liefert `git diff v6.6.0 v6.6.0 -- lab/regelwerk lab/templates` nichts,
  `git diff v6.5.0 v6.6.0 -- …` genau die genannten 10 Dateien und +87/−32.
  **Kein Gate fängt es:** das `versions`-Muster in `.d-check.yml` bindet an
  `(\.harness|\.\.)/baseline/(v\d+\.\d+\.\d+)/`, also an einen **Pfad** — eine
  nackte Kennung im Fließtext sieht es nicht, und genau die ist hier gehoben
  worden.
- `klasse`: Historische Versions-Nennung beim Pin-Bump mitgehoben

### F-2 — Die Closure-Log-Zeile zu welle-15 behauptet nach dem Bump eine Migration nach `v6.6.0`

- `kategorie`: HIGH
- `quelle`: `v6.6.0` · `regelwerk/modul-06-roadmap.md` §Roadmap-Struktur: fünf
  Abschnitte (*Abgeschlossene Wellen* — „das Closure-Log (ruhender
  Audit-Bestand)"); nachweislich falsche Tatsachenbehauptung
- `pfad`: `docs/plan/planning/in-progress/roadmap.md` Zeile 94
- `befund`: `012530f` ändert die Zelle von *„Regelwerk-Migration
  `v6.2.0`→`v6.5.0`"* auf *„`v6.2.0`→`v6.6.0`"*. Die Zeile widerspricht damit
  der Ergebnisnotiz, auf die sie in derselben Zelle zeigt (`done/welle-15-results.md`
  Zeile 1: *„Regelwerk-Migration `v6.2.0` → `v6.5.0`"*), und der Welle-Kennung
  `welle-15-regelwerk-v650-migration` in der Nachbarzelle; welle-15 schloss am
  2026-09-07, `v6.6.0` erschien am 2026-09-08.
- `verifizierbar`: ja — `git show 0b97c5d:docs/plan/planning/in-progress/roadmap.md`
  gegen `HEAD`, dazu die erste Zeile der Ergebnisnotiz. Kein Gate fängt es aus
  demselben Grund wie F-1 (nackte Kennung, kein Pfad-Pin).
- `klasse`: Historische Versions-Nennung beim Pin-Bump mitgehoben

### F-3 — Die Negativbefund-Zeile der Sichtung erklärt `PLAN` und `REVIEW` für unberührt, während der Diff in beiden arbeitet

- `kategorie`: HIGH
- `quelle`: `v6.6.0` · `regelwerk/modul-05-planning-harness.md` §Zwei Schritte
  vor der Modus-Begründung (*„Steht eine der berührten Sub-Areas dort? Dann
  gehört der Zähler-Stand ins Kriterium Evidenz-/Diskrepanz-Risiko"*)
- `pfad`: `docs/plan/planning/in-progress/slice-192-baseline-v660-vendoring.md`
  §9, Zeilen 293–310 (Sub-Area-Wahl) gegen Zeilen 340–341 (Negativbefund-Zeile)
- `befund`: §9 führt als berührt `HARNESS` und die vendored Baseline, dazu
  `GATE` als Randberührung, und schließt: *„Zu `SPEC`, `PLAN`, `KERN`, `ADAPT`
  und `REVIEW` steht nichts Offenes an, das dieser Slice berührt."* Die Range
  ändert `docs/plan/planning/README.md`, `docs/plan/planning/in-progress/roadmap.md`,
  den Slice-Plan und vier Register-Dateien — alles `PLAN` — sowie
  `docs/reviews/README.md` (`REVIEW`); im Register stehen vier **offene**
  `BEO-PLAN`-Einträge bei je 2×, deren Klassen genau die Arbeitsformen dieses
  Slice sind (`messung-ohne-reproduzierbares-instrument`,
  `vollstaendigkeits-haken-ohne-erschoepften-gegenstand`,
  `zielsatz-nach-plan-aenderung-nicht-nachgezogen`,
  `kandidaten-klassifikation-groeber-als-der-kandidat`), und §9 Block (2) nennt
  zwei davon zwölf Zeilen höher selbst als *„in die Arbeit eingebaut"*.
- `verifizierbar`: teils — dass die Pfade berührt sind und die vier Einträge bei
  2× offen stehen, ist mechanisch (`git diff --name-only`, `ls …/evidence`); ob
  ein Eintrag eine Sub-Area „betrifft", bleibt ein Urteil.
- `klasse`: Sichtungs-Schritt liest eine andere Sub-Area als die deklariert berührte

### F-4 — Drei Pin-Zahlen im selben Vorgang, und keine deckt die Messung

- `kategorie`: MEDIUM
- `quelle`: DoD-Haken gegen den Gegenstand (`AGENTS.md` §5, Closure-Pflicht);
  Reviewer-Skill §Klassifikation, unbelegte Tatsachenbehauptung
- `pfad`: `docs/plan/planning/in-progress/slice-192-baseline-v660-vendoring.md`
  Zeilen 22–24 (§1), 100 (§3.2, Tabellenzeile *Lebende Dateien mit Pin*), 167
  (DoD); dazu die Commit-Message von `012530f`
- `befund`: §1 und der abgehakte DoD-Punkt nennen **46** Vorkommen in **14**
  lebenden Dateien, §3.2 und die Commit-Message nennen **16** Dateien.
  Nachgemessen am Stand `d410675` (unmittelbar vor der Plan-Anlage) tragen **15**
  lebende Dateien zusammen **47** Nennungen des alten Standes, und `012530f`
  bumpt genau diese 15 — dazu vier Symlinks und zwei eingefrorene Einträge.
  Zählt man statt der Nennungen die Pfad-Pins im Sinne des `versions`-Musters,
  sind es **37** in **13** lebenden Dateien. Unter beiden Lesarten trifft weder
  46/14 noch 16, und welche Zählregel gilt, sagt der Slice nicht.
- `verifizierbar`: ja — `grep -o` über den Baum bei `d410675` bzw. `0b97c5d`
  (ohne `.harness/baseline/`, ohne `done/`, `evidence/`, `conventions/done/`) und
  `git show 012530f --numstat`.
- `klasse`: Vollständigkeits-Haken gesetzt, ohne den Gegenstand zu erschöpfen

### F-5 — §3.5 rechnet die zehn Delta-Dateien falsch auf und adressiert die Vorlage, die die Folge-Slices nicht betrifft

- `kategorie`: MEDIUM
- `quelle`: nachweislich falsche Tatsachenbehauptung; `harness/conventions.md`
  §Baseline (*Delta und Voll-Abgleich*)
- `pfad`: `docs/plan/planning/in-progress/slice-192-baseline-v660-vendoring.md`
  Zeilen 154–160
- `befund`: §3.5 zerlegt die zehn Dateien aus §2 in *„neun echte plus eine, die
  sich nur im gepinnten Release-Asset-Link unterscheidet:
  `templates/harness/conventions.template.md`"*. Gemessen kommt
  `conventions.template.md` in jenem `git diff` überhaupt nicht vor — ihr
  Unterschied entsteht erst beim Bundle-Bau durch den Tag-Pin —, nach
  Tag-Normalisierung unterscheiden sich weiterhin **zehn** Dateien, und die
  zehnte ist `regelwerk/README.md` mit der Stand-Zeile *Kurs-Welle 128 → 129*.
  Die abgeleitete Aussage adressiert slice-188/189: deren Paare sind
  `.d-check.yml`, `closure-note-reviewer.md` und die drei Spec-Straten —
  `conventions.template.md` ist keines davon, während `templates/.d-check.yml`
  (+24 Zeilen, slice-188s erstes Paar) in §3.5 nicht vorkommt und slice-188s
  Ausgangsmessung *„42 Kandidaten"* damit überholt ist.
- `verifizierbar`: ja — beide Bundles aus `git archive <tag>` neu bauen, Tags zu
  einem Platzhalter normalisieren, `diff -rq`; dazu die Paar-Tabellen von
  slice-188 §1 und slice-189 §1.
- `klasse`: Ableitung vermischt zwei Grundgesamtheiten

### F-6 — Der Risiko-Ausgang stützt sich auf eine Null-Treffer-Messung ohne Geltungsbereich, die den Fund aus F-1 konstruktionsbedingt nicht sehen kann

- `kategorie`: MEDIUM
- `quelle`: Reviewer-Skill §Mess-Regeln, *Geltungsbereich einer Messung*
  (`seit slice-179`); `v6.6.0` · `regelwerk/modul-05-planning-harness.md`
  §Offene Risiken werden bei Closure aufgelöst
- `pfad`: `docs/plan/planning/in-progress/slice-192-baseline-v660-vendoring.md`
  Zeilen 209–218 (Risiko 2, Ausgang *entfallen*)
- `befund`: Die Begründung lautet *„`grep` auf den alten Stand liefert außerhalb
  der eingefrorenen Artefakte **null** Treffer"*; der Plan selbst ist ein
  lebendes Dokument und trägt in §3.1 und §3.3 zwei solche Treffer (dort
  richtig, als Kennung). Vor allem misst der `grep` nur eine Richtung: Er sucht
  den **alten** Stand und schweigt genau dort, wo eine historische Nennung
  fälschlich auf den neuen gehoben wurde — die Null ist teilweise ein Ergebnis
  des Defekts aus F-1/F-2, nicht ein Beleg gegen ihn.
- `verifizierbar`: ja — `grep -rn "v6\.5\.0"` ohne `.harness/baseline/` listet
  zwei Treffer im Plan; die Gegenrichtung ist `grep` auf den **neuen** Stand in
  Dokumenten, die über die Vergangenheit sprechen.
- `klasse`: Geltungsbereich einer Messung nicht genannt

### F-7 — §9 schreibt dem Report zu slice-187 eine Finding-Klasse zu, die dieser als erledigt führt

- `kategorie`: MEDIUM
- `quelle`: Reviewer-Skill §Output-Schema (*„die Bezeichnung muss über Läufe
  hinweg stabil sein"*, Norm aus `v6.6.0` ·
  `templates/docs/reviews/review-report.template.md`)
- `pfad`: `docs/plan/planning/in-progress/slice-192-baseline-v660-vendoring.md`
  Zeilen 333–338
- `befund`: §9 Block (2) nennt *„zwei seiner Finding-Klassen"*:
  *Vollständigkeits-Haken ohne erschöpften Gegenstand* — die trägt der Report zu
  slice-187 dreimal (F-5, F-6, F-8) — und *Messung ohne reproduzierbares
  Instrument*. Die zweite ist kein `klasse`-Feld jenes Reports; seine §Summary
  führt sie unter *„Erledigt sind zwei: Messung als Beleg ohne reproduzierbares
  Instrument (dort F-6)"*, also als Klasse des **Vorgänger**-Reports, die
  slice-187 aufgelöst hat. Der Zähler-Übergabepunkt zitiert damit einen Lauf,
  der die Klasse nicht führt.
- `verifizierbar`: ja — alle 15 `klasse`-Zeilen und die §Summary des Reports zu
  slice-187 (aus `done/slice-187-archiv.zip`).
- `klasse`: Finding-Klasse dem falschen Lauf zugeschrieben

### F-8 — Der Lerneintrag benennt eine Regel ohne Zielort; die Stelle, deren Reichweite er widerlegt, bleibt unverändert

- `kategorie`: MEDIUM
- `quelle`: `AGENTS.md` §5 (Closure-Pflicht, Lerneintrag-Formen);
  `v6.6.0` · `regelwerk/modul-05-planning-harness.md` §Closure- und
  Lerneintrag-Regeln
- `pfad`: `docs/plan/planning/in-progress/slice-192-baseline-v660-vendoring.md`
  Zeilen 237–252; `harness/conventions.md` §Baseline, Absatz *Aufgelöste
  Einträge sind davon ausgenommen*; `AGENTS.md` §5, Zitier-Form-Punkt, Zeile 313
- `befund`: §8 deklariert die Form *geschärfte Regel* und formuliert sie —
  *„Eine Messung, die eine Kollision für erledigt erklärt, nennt die
  Datei-Klasse, über die sie gemessen hat"* —, ohne sie an einem Ort abzulegen.
  `harness/conventions.md` §Baseline behauptet unverändert weiter, die Kollision
  sei *„bezahlt"*, belegt mit *„null Nachzügen"*, ohne die Datei-Klasse zu
  nennen, über die damals gemessen wurde; und die Aufzählung der einfrierenden
  Ziel-Formen in `AGENTS.md` §5 führt weiterhin genau vier und nicht
  `harness/conventions/done/` — der nächste aufgelöste `MR` darf denselben Link
  wieder tragen, der hier zwei Nachzüge gekostet hat. (Die andere, tatsächlich
  verkörperte Regel dieses Slice — Zitier-Form für Planungs-Dokumente,
  `seit slice-192` — ist davon unberührt und in `AGENTS.md` §5 abgelegt.)
- `verifizierbar`: nein — ob eine Regel „verkörpert" ist, ist ein Urteil; dass
  beide Stellen in der Range unverändert sind, ist mechanisch
  (`git diff 0b97c5d..HEAD -- harness/conventions.md AGENTS.md`).
- `klasse`: Lerneintrag als geschärfte Regel deklariert, ohne Zielort

### F-9 — Der genannte Mess-Ausschluss liegt außerhalb des gemessenen Gegenstands, die Abschnittsgrenze ist undeklariert

- `kategorie`: LOW
- `quelle`: Reviewer-Skill §Mess-Regeln, *Geltungsbereich einer Messung*
- `pfad`: `docs/plan/planning/in-progress/slice-192-baseline-v660-vendoring.md`
  Zeilen 136–146 (§3.4)
- `befund`: §3.4 betont für **beide** Stufen den Ausschluss der
  `<!-- Quelle: -->`-Zeile; diese steht in jeder Regelwerk-Datei auf Zeile 2
  bzw. 3, also vor der ersten `###`-Überschrift, und kann in keinen der fünf
  gemessenen Abschnitte fallen — für die Abschnitts-Stufe ist der Ausschluss
  gegenstandslos. Vier der fünf Zeichenzahlen sind auf ±1 reproduzierbar; die
  fünfte (MR-022, 4593) entspricht dem Abschnitt **ohne** seinen
  Unterabschnitt `#### Vergabe`, während die anderen vier keinen Unterabschnitt
  haben — welche Grenze gilt, sagt §3.4 nicht.
- `verifizierbar`: ja — nachgemessen; am Ergebnis ändert die Grenze nichts:
  alle fünf Abschnitte sind auch **einschließlich** Unterabschnitt wortgleich
  (22 742 / 553 / 10 340 / 1371 / 10 768 Zeichen, beide Stände identisch).
- `klasse`: Mess-Ausschluss ohne Gegenstand

### F-10 — Beleg und Ausgang der 3×-Beobachtung entstanden im Planungs-Commit, §8 stellt sie als Closure-Handlung dar

- `kategorie`: LOW
- `quelle`: `v6.6.0` · `regelwerk/modul-06-roadmap.md` §Das
  Beobachtungs-Register (*„Eingetragen wird bei der Slice-Closure"*; der Beleg
  ist die Kennung eines **abgeschlossenen** Vorgangs)
- `pfad`: `docs/plan/planning/observations/BEO-GATE/versions-sensor-trifft-planungs-vorgriff/state.md`
  und `…/evidence/slice-192.md` (beide angelegt in `0273140`);
  `docs/plan/planning/in-progress/slice-192-baseline-v660-vendoring.md`
  Zeilen 281–286
- `befund`: Beleg und Ausgang wurden geschrieben, als slice-192 noch in `open/`
  lag — zwei Commits vor dem Übergang nach `in-progress/` und drei vor der
  Arbeit; §8 formuliert es dagegen als Vorgang der Closure (*„ist jetzt
  gezogen"*). Die Regel erlaubt das Schreiben vor dem `git mv`, nicht vor dem
  Vorgang.
- `verifizierbar`: ja — `git log --oneline -- <register-verzeichnis>` zeigt
  `0273140` als anlegenden Commit.
- `klasse`: Attestierung vor dem Vorgang

### F-11 — Die Prämisse des vierten Risiko-Ausgangs ist nicht mehr nachprüfbar

- `kategorie`: INFO
- `quelle`: Reviewer-Skill §Klassifikation, unbelegte Tatsachenbehauptung
- `pfad`: `docs/plan/planning/in-progress/slice-192-baseline-v660-vendoring.md`
  Zeilen 226–233 (Risiko 4)
- `befund`: Der Ausgang begründet den Bau-Weg damit, dass der lokale
  Kurs-Checkout eine uncommittete Änderung trage; zum Prüfzeitpunkt ist er
  sauber (`git status --porcelain` leer, HEAD einen Commit **nach** dem Tag).
  Der Ausgang hängt nicht daran — ich habe beide Bundles unabhängig aus
  `git archive <tag>` nachgebaut —, aber die Aussage über ein repo-externes
  System ist retrospektiv nicht mehr belegbar und war es beim Schreiben nur für
  den Schreibenden.
- `verifizierbar`: nein — der Zustand eines fremden Arbeitsbaums zu einem
  vergangenen Zeitpunkt ist im Repo nicht festgehalten.
- `klasse`: Prämisse über ein repo-externes System ohne festgehaltenen Beleg

### F-12 — `.harness/skills/` hat keine Zeile in der Modus-Deklaration, wird aber von diesem Slice geändert

- `kategorie`: INFO
- `quelle`: `harness/conventions.md` §Modus-Deklaration pro Sub-Area
  (Qualifikation über die drei Inklusions-Achsen)
- `pfad`: `harness/conventions.md` §Modus-Deklaration (Tabelle);
  `.harness/skills/reviewer.md` (in `012530f` geändert)
- `befund`: Die Tabelle führt für `HARNESS` die Pfade `AGENTS.md`, `CLAUDE.md`
  und `harness/`; `.harness/skills/` steht in keiner Zeile, obwohl das
  Verzeichnis zwei Skill-Dateien mit eigener Pfad-Familie trägt und in dieser
  Range geändert wurde. Das ist Bestand, nicht durch diesen Slice entstanden —
  aber dieselbe Klasse, die das Register unter
  `BEO-HARNESS/harness-einstieg-ohne-modus-zeile` bereits als *verkörpert*
  führt.
- `verifizierbar`: ja — die Tabelle gegen `git show 012530f --name-only`.
- `klasse`: Harness-Datei ohne Zeile in der Modus-Deklaration

## Negativbefunde

- geprüft, ohne Befund: **Der Bau-Weg aus §3.1, vollständig nachgefahren.**
  `git archive v6.6.0` in ein Wegwerf-Verzeichnis, dort `tools/build-bundle.sh
  <ziel> v6.6.0` — das Ergebnis ist zum vendorten Baum unter dem neuen Tag
  **byte-gleich** (`diff -rq`: einziger Unterschied ist das nur dort liegende
  `SHA256SUMS`), 54 Dateien wie behauptet.
- geprüft, ohne Befund: **Die Gegenprobe am alten Tag trägt.** Derselbe Weg auf
  `v6.5.0` reproduziert den bei `0b97c5d` vendorten Baum byte-gleich — die
  Behauptung „auf demselben Weg entstanden, nicht auf einem ähnlichen" ist damit
  belegt und nicht nur erklärt.
- geprüft, ohne Befund: **Das Manifest-Verfahren.** `find . -type f ! -name
  SHA256SUMS | sort | xargs sha256sum` erzeugt für **beide** Stände eine mit dem
  committeten `SHA256SUMS` byte-gleiche Datei; `sha256sum -c` gegen den
  vendorten Baum ist grün, 54 Einträge, keine unmanifestierte Datei. Präzisions-
  Vorbehalt ohne Folge: `make regelwerk-check` **prüft** mit `sha256sum -c` und
  benutzt die `find`-Pipeline nur für die Gegenrichtung — „derselbe Aufruf" gilt
  für die Datei-Auswahl, nicht für den Prüf-Befehl.
- geprüft, ohne Befund: **Alle fünf aktiven `MR`-Zeiger dürfen wandern.**
  Abschnittsweise gegen beide gebauten Bundles gemessen — `grundlagen-referenz-richtung.md`
  §Referenz-Richtung (SDP), `modul-15-observability.md` §Kernidee,
  `modul-06-roadmap.md` §Wellen-Closure-Prozedur, `modul-08-agentenrollen.md`
  §Die neun Übergaben, `grundlagen-source-precedence.md` §ID-Schema als Klammer:
  fünfmal zeichengleich. Die zweistufige Anlage trägt: MR-014 ist der einzige
  Zeiger in eine geänderte Datei, und die Änderung liegt in
  §Doku-Konsistenz-Drift-Regeln, nicht in §Kernidee.
- geprüft, ohne Befund: **Die Delta-Menge ist vollständig.** `git diff v6.5.0
  v6.6.0 -- lab/regelwerk lab/templates` liefert genau die zehn Dateien der
  §2-Tabelle mit +87/−32; nach Tag-Normalisierung der gebauten Bundles bleiben
  dieselben zehn. Kein weiterer Punkt ist einschlägig: `modul-02` trifft eine
  abgeschlossene Bootstrap-Phase, `modul-09` ist Folge des Gate-Index-Punkts,
  `modul-15` ändert nur den Abschnitt, den MR-014 ausnimmt.
- geprüft, ohne Befund: **Die zwei Einschätzungen zur maschinellen Hälfte
  stimmen.** `exempt-targets` in `.d-check.yml` ist namentlich
  (`[help, build, compile, hooks, arch-graph]`), kein Glob; Target-Zellen mit
  Argument in der Code-Span gibt es in beiden Tabellen **null**, verlinkte
  je **16** — beide Zahlen exakt nachgemessen, und die neue Zellen-Regel nimmt
  verlinkte Zellen ausdrücklich aus.
- geprüft, ohne Befund: **Kein lebendes Dokument nennt den alten Stand als
  Pfad, und kein eingefrorenes wurde fälschlich gebumpt.** Die verbleibenden
  Nennungen liegen in `done/`, `observations/**/evidence/`,
  `conventions/done/` und im Plan selbst, alle als Kennung in Inline-Code, keine
  als Link — der entfernte Baum hinterlässt kein totes Ziel. Umgekehrt trägt
  außer den zwei unter F-1/F-2 genannten Stellen kein Dokument den neuen Stand
  an einer historischen Aussage.
- geprüft, ohne Befund: **Die Umstellung von MR-011 und MR-021 ist eine
  Form-, keine Inhaltsänderung.** `git show` zeigt in beiden Dateien genau eine
  geänderte Zeile: der `Ersetzt-Baseline-Regel`-Link wird zur Kennung, der
  **genannte Stand bleibt `v6.5.0`**, Datei und Abschnitt bleiben wortgleich.
  Die Immutabilitäts-Disziplin („nichts nachträglich **inhaltlich** geändert")
  ist damit gewahrt; die Alternative wäre ein garantiert toter Link gewesen.
  Der Regel-**Widerspruch**, den das aufdeckt, ist unter F-8 gebucht, nicht hier.
- geprüft, ohne Befund: **Symlinks.** Alle sieben getrackten Symlinks lösen auf,
  die vier Baseline-Ziele zeigen auf den adoptierten Stand.
- geprüft, ohne Befund: **`AGENTS.md` §3.3 (git mv + Inhaltsänderung).** Der
  Baum-Wechsel liegt in **einem** Commit; die Rename-Detection greift trotzdem
  für alle **55** bewegten Dateien — 23 × R100, 53 × ≥ 90 %, die niedrigste
  Ähnlichkeit ist R065 (`SHA256SUMS`) —, also durchweg über der 50 %-Schwelle,
  die §3.3 schützt. Vendorter Fremdtext ist zudem nicht der Gegenstand der Regel.
- geprüft, ohne Befund: **Die Closure-Notiz trägt Substanz, nicht Floskel.**
  Sie führt ein Lernsignal mit Ursache, ein konkretes Folge-Slice mit auffindbarer
  Datei (slice-193 liegt in `open/`) und zwei beobachtbare Kriterien, davon eines
  nachrechenbar (die Gegenprobe). Der Maßstab
  `.harness/skills/closure-note-reviewer.md` liefert kein Finding; die Reibung am
  Lerneintrag ist eine Zielort-Frage (F-8), keine Floskel-Frage.
- geprüft, ohne Befund: **Die vier Risiko-Ausgänge sind je einer aus der
  geschlossenen Dreier-Menge**, und der eingetretene trägt eine Folge-Slice-ID,
  die die Sendung annimmt: slice-193 §1 nennt denselben Umbau als Ziel.
  Substanzielle Einwände zu zweien der drei *entfallen* stehen unter F-6 und
  F-11; der dritte (MR-Zeiger) trägt vollständig.
- geprüft, ohne Befund: **Die anderen fünf Register-Zuordnungen aus §9 stimmen.**
  `zwei-baseline-staende-nach-migrationsende`, `rueckbau-kandidat-ueberlebt-baseline-migration`,
  `adaption-korrigiert-repo-aussage`, `baseline-regel-nie-erwogen-weil-bestand-sie-verletzt`
  und `symlink-ziel-nach-baseline-bump-ungeprueft` sind mit dem richtigen Kürzel,
  dem richtigen Zähler-Stand und der richtigen Berührungs-Art geführt; die
  Aussage „kein aktiver `MR` wird durch diesen Sprung entbehrlich" hält gegen
  alle acht aktiven Einträge.
- geprüft, ohne Befund: **Die neue Beobachtung ist formgerecht angelegt** —
  `observation.md` (Identität, Sub-Area geführt), `state.md` (Zustand `offen
  (1×)` ohne Chronik, mit Abgrenzung gegen die nächstähnliche Kennung),
  `evidence/slice-192.md` als einzelner Beleg.
- geprüft, ohne Befund: **DoD-Größenregel.** Sieben Punkte, davon vier pro Slice
  konstant (Review, Closure-Notiz, Register, Risiko-Ausgänge) — drei
  Liefer-Punkte, an der Obergrenze, nicht darüber. Der Review-Punkt steht
  korrekt **ungehakt**.
- geprüft, ohne Befund: **Commit-Disziplin.** `012530f` trägt Scope `harness`
  und darf damit außerhalb von `docs/plan/planning/` arbeiten; `28d647b` trägt
  Scope `planning` und berührt ausschließlich `docs/plan/planning/`. Beide
  Messages nennen die Slice-Kennung.
- geprüft, ohne Befund: **Begleit-Aussagen des Bumps sind nach dem Sprung noch
  wahr** — 17 Module und acht Grundlagen-Abschnitte, 28 Vorlagen, identische
  Dateimenge in beiden Ständen; die „15 Paare mit genau einem Gegenstück" aus
  `harness/conventions.md` §Baseline sind unberührt.
- geprüft, ohne Befund: **Kein Spec-Stratum ist berührt**, keine `AC-*`
  geändert, keine ADR angefasst — die Kopfzeile *Berührte Spec-Stellen: —*
  trägt.

## Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 3 |
| MEDIUM | 5 |
| LOW | 2 |
| INFO | 2 |

**Finding-Klassen dieses Laufs:** Historische Versions-Nennung beim Pin-Bump
mitgehoben (F-1, F-2 — zwei Funde, **ein** Vorgang) · Sichtungs-Schritt liest
eine andere Sub-Area als die deklariert berührte (F-3) · Vollständigkeits-Haken
gesetzt, ohne den Gegenstand zu erschöpfen (F-4) · Ableitung vermischt zwei
Grundgesamtheiten (F-5) · Geltungsbereich einer Messung nicht genannt (F-6) ·
Finding-Klasse dem falschen Lauf zugeschrieben (F-7) · Lerneintrag als
geschärfte Regel deklariert, ohne Zielort (F-8) · Mess-Ausschluss ohne
Gegenstand (F-9) · Attestierung vor dem Vorgang (F-10) · Prämisse über ein
repo-externes System ohne festgehaltenen Beleg (F-11) · Harness-Datei ohne Zeile
in der Modus-Deklaration (F-12).

**Wiederkehrend gegenüber den Vorgänger-Reports:** *Sichtungs-Schritt liest eine
andere Sub-Area als die deklariert berührte* (im Report zu slice-187 F-2, hier
F-3) und *Vollständigkeits-Haken gesetzt, ohne den Gegenstand zu erschöpfen*
(dort F-5/F-6/F-8, hier F-4) — beide unmittelbar nach dem Lauf, der sie
benannt hat. *Attestierung vor dem Vorgang* (F-10) trifft eine Kennung, die im
Register bei 1× steht. Ob daraus jeweils ein Ausgang wird, ist eine
Planner-Entscheidung, kein Reviewer-Vorschlag.

## Verdikt

**Merge-blockierend:** ja — für F-1, F-2 und F-3.

**F-1 und F-2 sind derselbe Handgriff an zwei Orten** und die schwersten
Befunde: Der Slice hat beim Pin-Bump nicht zwischen einem Zeiger auf den
adoptierten Stand und einer Aussage über die Vergangenheit unterschieden. Die
Folge ist, dass der Slice sein eigenes Messinstrument als leeren Diff ausweist,
sein Zielsatz und ein abgehakter DoD-Punkt die Entfernung des gerade adoptierten
Standes behaupten und das Closure-Log der Roadmap welle-15 eine Migration
zuschreibt, die sie nicht durchgeführt hat. Kein Gate fängt das: `versions`
bindet an Pfade, und der einzige Beleg, den §7 dafür anführt, ist eine Suche
nach dem alten Stand — sie kann eine fälschlich gehobene Nennung nicht sehen.
**F-3** ist die Wiederholung des Befundes, aus dem der Slice ausdrücklich gelernt
haben will: Die Sichtung liest diesmal `BEO-HARNESS` und `BEO-GATE` gründlich und
erklärt `PLAN` für frei — während §9 zwölf Zeilen höher eine offene
`BEO-PLAN`-Klasse als in die Arbeit eingebaut führt und vier dortige Einträge bei
2× stehen, deren Klassen dieser Diff trägt.

**Die fünf MEDIUM sind je einzeln in wenigen Zeilen behebbar**, hängen aber an
derselben Wurzel wie F-6: Vier von ihnen (F-4, F-5, F-6, F-7) sind Aussagen über
eine Messung, deren Grundgesamtheit nicht mitgenannt ist.

**Was ausdrücklich trägt:** Der Bau-Weg ist der belastbarste, den diese Kette
bisher hatte. Ich habe beide Bundles unabhängig aus den Tags gebaut, den neuen
Baum byte-gleich reproduziert, die Gegenprobe am alten Tag nachvollzogen und das
Manifest für beide Stände neu erzeugt — alles vier stimmt. Die MR-Zeiger-Messung
ist zweistufig richtig angelegt und in allen fünf Fällen nachgerechnet; die
Auflösung des Regel-Widerspruchs bei MR-011 und MR-021 ist der minimale Eingriff
und wahrt die Immutabilität. Und die Delta-Analyse ist inhaltlich vollständig:
Der eine Punkt, der a-check trifft, ist richtig erkannt, richtig abgegrenzt und
hat eine Adresse, die ihn annimmt.
