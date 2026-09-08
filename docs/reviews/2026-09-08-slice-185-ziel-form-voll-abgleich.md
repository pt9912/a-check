# Review-Report: slice-185 — 2026-09-08

**Review-Art:** Plan + Diff — geprüft gegen den Slice-Plan, die Ziel-Formen des
vendorten Stands und die berührten Repo-Artefakte (Modul 10 §Drei Review-Arten).
Kern-Behauptung dieses Slice ist eine **Messung**: welche Ziel-Formen es gibt,
welche ein Gegenstück haben, und was diesem Gegenstück fehlt. Der Review misst
nach.

**Gegenstand:** Commits `43383a8` („docs(planning): slice-185 schneiden") und
`3951644` („feat(harness): slice-185 -- drei HIGH-Kategorien und der
Derivativ-Hinweis"); der dazwischen liegende `392fd59` ist ein reiner
`slice-mv`. **Während des Laufs bewegte sich der Bestand:** `7de5e5d`
(„Zweck-Spalte auf 150 Zeichen begrenzt", ebenfalls `slice-185`) und eine
unversionierte Erweiterung von `slice-186` §2.1 kamen hinzu. Beide sind **nicht**
Teil des geprüften Gegenstands; keiner berührt eine Fundstelle dieses Reports.
Was am Rand auffiel, steht als F-12 mit dieser Einschränkung.

**Skill:** `.harness/skills/reviewer.md` @ Stand `3951644` — **einschließlich der
drei mit diesem Commit ergänzten HIGH-Kategorien**, die damit auf ihren eigenen
Erzeuger angewendet werden.
**Modell:** claude-opus-5 · **Datum:** 2026-09-08

**Eingangs-Kontext:**

- `slice-185` (Plan, `in-progress/`), `slice-186` (Folge-Slice, `open/`)
- `AGENTS.md` §3.1/§3.2/§3.3/§3.4/§3.7, §5, §6
- `harness/conventions.md` §Baseline · `harness/README.md` §Sensors
- `spec/architecture.md`, `docs/plan/adr/README.md`,
  `docs/plan/carveouts/README.md`, `README.md`
- `v6.5.0` · `templates/` (alle 28 Dateien), `templates/README.md`
- `v6.5.0` · `regelwerk/modul-05-planning-harness.md`,
  `regelwerk/modul-06-roadmap.md`, `regelwerk/modul-08-agentenrollen.md`
- Beobachtungs-Register `BEO-HARNESS/hard-rule-37-ohne-sensor`,
  `BEO-HARNESS/chronik-in-gelesenen-dateien`
- frühere Reports zum selben Bereich: `slice-161`, `slice-164`, `slice-167`
  (Baseline-Delta-Analysen)

---

## Findings

### F-1 — `README.md` ist gegen den Verzeichnis-**Index** gemessen, nicht gegen seine Ziel-Form

- `kategorie`: HIGH
- `quelle`: Reviewer-Skill §Klassifikation, „nachweislich falsche
  Tatsachenbehauptung"; `v6.5.0` · `templates/README.md` §Übersicht
  (Zuordnungstabelle: `project-readme.template.md` → „Projekt-Root-`README.md`")
- `pfad`: `docs/plan/planning/in-progress/slice-185-ziel-form-voll-abgleich.md`
  §2.1 (Tabelle „Zwei Gegenstücke sind kürzer als ihre Vorlage");
  fortgeschrieben in
  `docs/plan/planning/open/slice-186-voll-abgleich-restliche-paare.md` §2
- `befund`: Die als Ziel-Form von `README.md` angesetzten **13 433** Zeichen
  gehören zu `v6.5.0` · `templates/README.md` — dem **Index des
  Templates-Verzeichnisses** (13 535 roh), der sich selbst als
  „Referenz-Quelle" für die Vorlagen beschreibt und keine ist. Die Ziel-Form
  eines Projekt-`README.md` ist `project-readme.template.md` mit **2033** roh;
  a-checks `README.md` (9230) ist ihr gegenüber **länger**, nicht kürzer.
  *(Einheit dieser Zeile durchgehend `wc -m`, also Zeichen; in Bytes, `wc -c`,
  lauten dieselben drei Werte 13 707 / 2076 / 9274 — der Vergleich kippt in
  keiner der beiden Einheiten.)* Damit trägt die Aussage „**Zwei** Gegenstücke sind kürzer als ihre
  Vorlage" nur einen Fall, und `slice-186` §2 begründet seine Reihenfolge
  („`README.md` zuerst") auf dem entfallenen zweiten. Konsistent damit führt
  §1 desselben Plans `project-readme.template.md` unter den Ziel-Formen
  **ohne** Gegenstück — obwohl `README.md` Rang 7 der Source Precedence ist
  und im selben Plan als eines der 14 Paare gezählt wird.
- `verifizierbar`: ja — `wc -m` auf die drei Dateien; unter `templates/` liegt
  keine andere Datei im Bereich 12 800–14 200 Zeichen, die als Quelle der
  13 433 in Frage käme.
- `klasse`: Ziel-Form-Paar gegen den Verzeichnis-Index statt gegen die Vorlage
  gemessen

### F-2 — Der Nicht-Befund zu `spec/architecture.md` deckt nur eine von drei Klauseln

- `kategorie`: HIGH
- `quelle`: `AGENTS.md` §3.4; `v6.5.0` ·
  `templates/spec/architecture.template.md` §Hard Rule
- `pfad`: `docs/plan/planning/in-progress/slice-185-ziel-form-voll-abgleich.md`
  §3 („bewusst abweichend") und §8, vierter Spiegelstrich
- `befund`: Die Hard Rule der Ziel-Form verlangt **drei** Dinge: keine Wellen /
  Slices / Commit-Hashes / Closure-Daten · **keine ADR-Bezüge** · **keine
  Historie** („`Letzte Änderung` oben ist ein Frische-Marker, kein Protokoll").
  Der Plan zitiert sie ohne die dritte Klausel und belegt „hält sie ein" mit
  einer einzigen Messung („null ADR-Verweise"). `AGENTS.md` §3.4 nennt die
  Historie-Klausel nicht — es deckt ADRs, Wellen, Slices, Commit-Hashes und
  Closure-Daten —, und `spec/architecture.md` trägt §8 *Historie* als
  vierzeiliges Änderungsprotokoll. Das Kopffeld derselben Datei steht auf
  `Datum: 2026-06-21`, während §8 eine Änderung vom 2026-09-05 verzeichnet und
  `git log` den letzten Commit an der Datei auf 2026-09-05 datiert. Der Satz
  „**Die Differenz ist also keine Lücke**" ist damit für ein Drittel der Hard
  Rule belegt und für den Rest weder gemessen noch abgedeckt.
- `verifizierbar`: ja — `grep -n '^#' spec/architecture.md` zeigt §8 *Historie*;
  `sed -n '/^### 3.4/,/^### 3.5/p' AGENTS.md` zeigt den Umfang der Hard Rule;
  `git log -1 --date=short -- spec/architecture.md` zeigt 2026-09-05.
- `klasse`: Teil-Messung als Deckungs-Nachweis für eine mehrteilige Zusage
  ausgegeben

### F-3 — „a-check kopiert die Ziel-Formen nicht" widerspricht `AGENTS.md` §5

- `kategorie`: HIGH
- `quelle`: `AGENTS.md` §5 („Slice-Form: neue Slices entstehen aus der vendored
  Ziel-Form … **Beim Kopieren anzupassen** — sieben Punkte");
  `docs/plan/carveouts/README.md` §Form eines Carveouts;
  `v6.5.0` · `templates/README.md` §Verwendung, Schritt 5
- `pfad`: `.harness/skills/reviewer.md` Zeilen 48–57 (HIGH-Kategorie „Norm nur
  im Kommentar")
- `befund`: Die neue Kategorie stützt ihre Umdeutung auf den Satz „**a-check
  kopiert die Ziel-Formen nicht** und hat dort keinen Fall". Belegt ist im Repo
  das Gegenteil: `AGENTS.md` §5 und `docs/plan/carveouts/README.md` beschreiben
  beide das Anlegen eines Artefakts als **Kopieren der vendorten Ziel-Form**
  („Beim Kopieren anzupassen"); was a-check nicht führt, ist eine eigene
  *Vorlagen-Datei* (`find . -name '*.template.md'` außerhalb der Baseline:
  leer). Beim Kopieren fallen die `<!-- -->`-Blöcke bestimmungsgemäß weg
  (`templates/README.md` §Verwendung, Schritt 5) — der Fall der Ziel-Form
  existiert also, und er ist besetzt: `v6.5.0` ·
  `templates/docs/reviews/review-report.template.md` trägt im vierten
  Kommentarblock die Norm „DIE KLASSEN-BEZEICHNUNG MUSS ÜBER LÄUFE HINWEG
  STABIL SEIN" ausschließlich dort. Mit der Umdeutung verliert die Kategorie
  genau die Zusage, die sie in der Ziel-Form trägt.
- `verifizierbar`: ja — `grep -n 'Beim Kopieren anzupassen' AGENTS.md
  docs/plan/carveouts/README.md`; die Kommentarblöcke der Ziel-Formen sind mit
  einem Regex über `<!--…-->` auszählbar (`slice.template.md` 9,
  `NNNN-titel.template.md` 6, `carveout.template.md` 5,
  `review-report.template.md` 4).
- `klasse`: Ziel-Form-Regel mit einer falschen Repo-Tatsache entschärft

### F-4 — Im Paar, das als abgeglichen geführt wird, fehlt das Feld `klasse`

- `kategorie`: HIGH
- `quelle`: `v6.5.0` · `templates/.harness/skills/reviewer.template.md`
  §Output-Schema; `v6.5.0` · `regelwerk/modul-05-planning-harness.md`
  §Closure- und Lerneintrag-Regeln (Finding-Klasse als eine der drei
  Zähler-Quellen)
- `pfad`: `.harness/skills/reviewer.md` Zeilen 80–83 (§Output-Schema);
  `docs/plan/planning/in-progress/slice-185-ziel-form-voll-abgleich.md` §3,
  Tabellenzeile `.harness/skills/reviewer.md`, und §4, abgehakter DoD-Punkt
  „**Fünf** Paare sind abgeglichen"
- `befund`: Die Ziel-Form führt **sechs** Felder je Finding — `kategorie` ·
  `quelle` · `pfad` · `befund` · `verifizierbar` · **`klasse`** („stabile
  Kurz-Bezeichnung des Fehlermusters … speist den Steering-Loop-Zähler").
  a-checks Skill führt fünf; `klasse` fehlt, und das Wort kommt in der Datei
  nur als Bestandteil eines Kategorien-Namens vor. Der Bestand praktiziert das
  Feld trotzdem (alle acht flachen Reports unter `docs/reviews/` tragen eine
  `klasse`-Zeile), es steht nur in keiner Regel. Der DoD-Punkt „Fünf Paare sind
  abgeglichen" ist abgehakt und §1 sagt „einmal **vollständig** gegen dieses
  Gegenstück abgeglichen" — für das Paar, das die gesamte Argumentation des
  Slice trägt und dessen Kandidaten-Spalte auf `—` steht, hält das nicht.
- `verifizierbar`: ja — `grep -n -i 'klasse' .harness/skills/reviewer.md`
  liefert eine Zeile (den Kategorien-Namen), `grep -n 'klasse'
  …/templates/docs/reviews/review-report.template.md` liefert das Schema-Feld.
- `klasse`: Vollständigkeits-Haken gesetzt, ohne den Gegenstand zu erschöpfen

### F-5 — Der Herkunfts-Anker `seit slice-185` steht nicht am Zielort

- `kategorie`: MEDIUM
- `quelle`: `v6.5.0` · `regelwerk/modul-06-roadmap.md` §Das
  Beobachtungs-Register (Ausgang *verkörpert*: „Zielort **und**
  Herkunfts-Anker") und §Wellen-Closure-Prozedur, Paarung (a); `AGENTS.md` §5
- `pfad`: `docs/plan/planning/in-progress/slice-185-ziel-form-voll-abgleich.md`
  §8, Steering-Loop-Eintrag („— liegt in `harness/conventions.md` §Baseline
  (`seit slice-185` dort, …)"); Zielort `harness/conventions.md` §Baseline
- `befund`: Die Zeichenfolge `seit slice-185` kommt im Repo ausschließlich in
  der Slice-Datei selbst vor; der genannte Zielort trägt sie nicht — dort steht
  „Gemessen an einem Beispiel (slice-185):". Der Bestand führt die Ankerform
  sonst wörtlich (`seit slice-171` in `MR-022`, `seit slice-179`,
  `seit slice-181`, `seit slice-182`, `seit slice-183` in `AGENTS.md`). Die
  Aussage steht im Indikativ, obwohl die Paarung laut derselben Notiz erst nach
  dem `git mv` fällig ist.
- `verifizierbar`: ja — `grep -rn 'seit slice-185' --include='*.md' .`
- `klasse`: Anker-Paarung im Indikativ behauptet, bevor der Zielort sie trägt

### F-6 — Zwei unvereinbare Zahlen für dieselbe Messung; das Instrument ist nicht im Repo

- `kategorie`: MEDIUM
- `quelle`: `AGENTS.md` §5 („Geltungsbereich einer Messung", `seit slice-179`)
- `pfad`: `docs/plan/planning/in-progress/slice-185-ziel-form-voll-abgleich.md`
  §2.3 („meldet **30** Sätze ohne Entsprechung") gegen
  `docs/plan/planning/open/slice-186-voll-abgleich-restliche-paare.md` §2
  (Tabellenzeile `AGENTS.md` | **22**)
- `befund`: Dieselbe maschinelle Vorauswahl über dasselbe Paar liefert im
  Schneide-Commit 30 und im Umsetzungs-Commit 22; keine der beiden Stellen
  nennt einen Unterschied im Geltungsbereich. Die Gesamtzahl **169** ist aus
  der 22 gebildet (155 aus `slice-186` § 2 plus 14 aus `slice-185` §3 = 169),
  nicht aus der 30 — die 169 selbst ist also intern konsistent, die 30 steht
  daneben. Das Messinstrument („Grob-Abgleich per Wortfolge") ist in keinem der
  beiden Commits enthalten; die Zahlen 26, 66 741, 30, 22, 169 sind damit
  Ergebnisse eines Laufs, den niemand wiederholen kann. Nachgerechnet ließ sich
  66 741 mit keiner der beiden naheliegenden Normalisierungen reproduzieren
  (Rohsumme der 14 Dateien: 80 113 mit dem Index aus F-1, 68 611 mit der
  richtigen Ziel-Form; normalisiert 60 623 bzw. 48 444).
- `verifizierbar`: nein — das Instrument fehlt; genau das ist der Befund.
- `klasse`: Messung als Beleg ohne reproduzierbares Instrument

### F-7 — Die Grundgesamtheit ist um eins zu klein, und `roadmap` ist als Instanz-Vorlage einsortiert

- `kategorie`: MEDIUM
- `quelle`: `v6.5.0` · `templates/README.md` §Ein- vs. wiederkehrende Templates
  (`roadmap` steht dort unter den **Singletons**)
- `pfad`: `docs/plan/planning/in-progress/slice-185-ziel-form-voll-abgleich.md`
  §2.1 („**26**") und §1; übernommen nach `harness/conventions.md` §Baseline
  (neuer Absatz: „die **14** Vorlagen mit genau einem Gegenstück … Die zwölf
  **Instanz**-Vorlagen … fallen heraus")
- `befund`: Das Templates-Verzeichnis trägt 28 Dateien: 25 `.template.md`,
  `Makefile`, `.d-check.yml` und den Verzeichnis-Index `README.md`. Ziel-Formen
  sind damit **27**, nicht 26; die genannte 26 ist exakt das Ergebnis von
  `find … -name '*.md'` — sie zählt den Index mit (siehe F-1) und `Makefile`
  sowie `.d-check.yml` nicht, obwohl beide unter den 14 Paaren geführt werden.
  Weiter zählt der Plan `roadmap.template.md` zu den Instanz-Vorlagen, obwohl
  die Baseline `roadmap` als Singleton führt und a-check mit
  `docs/plan/planning/in-progress/roadmap.md` genau **ein** Gegenstück hat; das
  Paar fehlt in `slice-185` wie in `slice-186`. Die Rechnung geht mit 14 + 12 +
  1 (`reconciliation`) auf 27 auf, nicht auf 26. Die Zahl **14** ist damit
  inzwischen eine Repo-Konvention in `harness/conventions.md`.
- `verifizierbar`: ja — `find …/templates -type f | wc -l` = 28,
  `… -name '*.md' | wc -l` = 26, `… -name '*.template.md' | wc -l` = 25;
  `find docs -name 'roadmap*.md'` liefert genau eine Datei.
- `klasse`: Grundgesamtheit aus einem Glob abgeleitet, dessen Menge von der
  Arbeitsmenge abweicht

### F-8 — `docs/plan/adr/README.md`: übernommen wurde der Bedienhinweis, der Normtext-Kandidat bleibt ohne Ausgang

- `kategorie`: MEDIUM
- `quelle`: `v6.5.0` · `templates/docs/plan/adr/README.template.md`
  (Template-Hinweis-Block gegen §Konventionen);
  Messbasis des Slice selbst: §2.1 „Normtext **ohne Bedienhinweise** und
  Platzhalter", §2.3 „Template-Hinweise … die beim Kopieren bestimmungsgemäß
  verschwinden"
- `pfad`: `docs/plan/planning/in-progress/slice-185-ziel-form-voll-abgleich.md`
  §3, Tabellenzeile `docs/plan/adr/README.md` (2 Kandidaten, „übernommen: der
  *Derivativ*-Hinweis")
- `befund`: Der übernommene *Derivativ*-Satz steht in der Ziel-Form
  **innerhalb** des `> **Template-Hinweis.**`-Blockquotes — also in genau der
  Klasse, die §2.1 aus dem Normtext herausrechnet und §2.3 als Rauschen führt.
  Der Kandidat aus dem Normtext-Teil (§Konventionen: „als Kennung (`SPEC-*`,
  `ARC-*`, …), ersatzweise als Abschnitt, wo die Sektion keine Kennungen
  vergibt. **Prozess-ADRs ohne Spec-Stratum tragen `—`**") fehlt in a-checks
  Index vollständig, obwohl der Bestand die Regel befolgt — `ADR-0001`,
  `ADR-0005`, `ADR-0006` und `ADR-0038` tragen `Schärft: —`. Für diesen zweiten
  Kandidaten nennt der Plan keinen der drei Ausgänge.
- `verifizierbar`: ja — die Blockquote-Zugehörigkeit steht in der Ziel-Form
  sichtbar; `grep -n -i 'Kennung\|Prozess-ADR\|ersatzweise'
  docs/plan/adr/README.md` ist leer.
- `klasse`: Kandidat aus dem Bedienhinweis übernommen, Kandidat aus dem
  Normtext übergangen

### F-9 — §1 (Ziel) und §4 (DoD) stehen nach der Plan-Änderung gegeneinander

- `kategorie`: LOW
- `quelle`: `AGENTS.md` §6, Schritt 4 („er erfindet die Abgrenzung nicht neu und
  darf sie nicht stillschweigend weiten"); `v6.5.0` ·
  `regelwerk/modul-05-planning-harness.md` §Ziel-Form: Slice
- `pfad`: `docs/plan/planning/in-progress/slice-185-ziel-form-voll-abgleich.md`
  §1 gegen §4
- `befund`: §1 sagt weiterhin „**Jede** vendored Ziel-Form mit genau einem
  Gegenstück im Repo ist einmal vollständig … abgeglichen"; geliefert sind
  fünf von 14. Die Plan-Änderung ist im DoD-Punkt ausdrücklich benannt — das
  ist die Hälfte, die die Regel verlangt —, aber der Zielsatz in §1 wurde nicht
  nachgezogen, sodass Ziel und DoD verschiedene Zusagen tragen.
- `verifizierbar`: nein — Lesart, kein Match.
- `klasse`: Zielsatz nach Plan-Änderung nicht nachgezogen

### F-10 — Ausschluss-Gründe in §1 treffen nicht auf die genannten Ziel-Formen zu

- `kategorie`: LOW
- `quelle`: `v6.5.0` · `templates/README.md` §Übersicht (`gate.template.md`:
  „Ein Gate oder Werkzeug … Index in `harness/README.md`")
- `pfad`: `docs/plan/planning/in-progress/slice-185-ziel-form-voll-abgleich.md`
  §1, zweiter Ausschluss-Punkt
- `befund`: Unter „Die Ziel-Formen **ohne Gegenstück**" stehen
  `reconciliation.template.md`, `project-readme.template.md` und
  `gate.template.md`. Nur die erste trifft zu: `gate.template.md` hat in
  a-check **16** Gegenstücke (`harness/sensors/*.md`) und ist damit eine
  Instanz-Vorlage; `project-readme.template.md` hat genau eines (siehe F-1).
  Der Ausgang „nicht abgleichen" ist für `gate.template.md` richtig, die
  angegebene Begründung ist es nicht — und ein Ausschluss ohne tragenden Grund
  ist laut §1-Regel eine Behauptung, keine Grenze.
- `verifizierbar`: ja — `ls harness/sensors/` liefert 16 Dateien.
- `klasse`: Ausschluss mit unzutreffender Begründung

### F-11 — „alle sechs sind Bootstrap-Bedienhinweise" trifft für den `Makefile`-Kopf nicht zu

- `kategorie`: LOW
- `quelle`: `v6.5.0` · `templates/Makefile`, Kopfkommentar
- `pfad`: `docs/plan/planning/in-progress/slice-185-ziel-form-voll-abgleich.md`
  §3, Tabellenzeile `Makefile`
- `befund`: Der Kopfkommentar der Ziel-Form trägt neben Bedienhinweisen einen
  Satz mit Norm-Charakter: „nur behauptete Targets in `AGENTS.md` /
  `harness/README.md` eintragen (keine halluzinierten Gates)". Der Ausgang
  „ohne Befund" ist trotzdem richtig — a-check trägt die Norm in `AGENTS.md`
  §4 und in `harness/README.md` §Nicht-Gates —, aber die Begründung („alle
  sechs sind Bootstrap-Bedienhinweise") deckt diesen Kandidaten nicht.
- `verifizierbar`: ja — `cat …/templates/Makefile`, erste sechs Zeilen.
- `klasse`: Kandidaten-Klassifikation gröber als der Kandidat

### F-12 — Der `hint` der neuen Zellengrenze nennt weiter den alten Wert

- `kategorie`: MEDIUM
- `quelle`: Reviewer-Skill §Klassifikation, neue HIGH-Kategorie „Norm nur im
  Kommentar" (Gegenstand `.d-check.yml`); `AGENTS.md` §3.7
- `pfad`: `.d-check.yml`, Block `files: "AGENTS.md"` / `section: "## 4. Quality
  Gates"` (Regel `cell-max-chars: 150`, `hint` eine Zeile darunter)
- `befund`: `7de5e5d` senkt `cell-max-chars` für die Spalte *Zweck* von 250 auf
  150; der `hint` derselben Regel — der Text, den ein **roter Lauf ausgibt** —
  sagt weiterhin „max 250 Zeichen". Rule und Meldung widersprechen sich, und die
  Meldung ist der einzige Ort, an dem ein Lauf die Schwelle nennt. **Dieser
  Commit gehört nicht zum erklärten Gegenstand** (er landete während des Laufs);
  das Finding steht hier, weil er dieselbe Slice-Kennung trägt und vor der
  Closure liegt — nicht als Ergebnis einer Prüfung von `7de5e5d`, die dieser
  Report nicht geleistet hat.
- `verifizierbar`: nein für die Abweichung selbst — kein Gate hält einen `hint`
  gegen seine Regel; das ist genau die Lage, die die neue HIGH-Kategorie
  beschreibt. Die Beobachtung selbst: `sed -n '378,385p' .d-check.yml`.
- `klasse`: Schwellenwert gesenkt, Meldungstext nicht nachgezogen

## Negativbefunde

- geprüft, ohne Befund: **`make gates`** — Exit 0 (Lauf in eine Datei
  umgeleitet, Exit-Code gelesen, nicht gepipet). **`make verify`** — Exit 0,
  21 Anforderungen, 0 Waisen. Die Zusage der beiden Commit-Messages hält.
- geprüft, ohne Befund: Hard Rule **§3.1** — die beiden Commits ändern
  ausschließlich Markdown; kein `go`/`pip`/`npm`/`cargo`-Aufruf, kein
  Produkt-Code. Hard Rule **§3.2** — keine Suppression berührt.
- geprüft, ohne Befund: Hard Rule **§3.3** — `392fd59` ist ein reiner
  `slice-mv` (`open/` → `in-progress/`), Inhalt und Bewegung sind getrennte
  Commits.
- geprüft, ohne Befund: Hard Rule **§3.7** auf die neuen Absätze selbst
  angewandt — keiner der drei neuen Kategorie-Blöcke, kein Absatz in den
  beiden Index-Dateien und kein Satz des `conventions.md`-Absatzes beschreibt
  eine verworfene Alternative im Konjunktiv oder bricht mitten im Satz ab. Die
  Formulierung „In der Baseline zielt das auf `<!-- -->`-Blöcke … a-check …
  hat dort keinen Fall" ist eine Abgrenzung und damit eine zulässige
  Kommentar-Klasse; ihr Wahrheitsgehalt ist F-3, nicht §3.7.
- geprüft, ohne Befund: die neue HIGH-Kategorie **„Zustandsfeld trägt
  Chronik"** auf den Commit selbst — die berührten Dateien tragen kein
  `Stand`-/`Status`-Feld, und der Roadmap-Drift-Log wurde nicht angefasst.
- geprüft, ohne Befund: die beiden zitierten Register-Einträge —
  `BEO-HARNESS/hard-rule-37-ohne-sensor` trägt `evidence/slice-108.md` und
  `evidence/slice-169.md`, `BEO-HARNESS/chronik-in-gelesenen-dateien` trägt
  `evidence/slice-103.md` und `evidence/slice-182.md`. Beide Zählstände (2×)
  im Plan sind korrekt, beide `state.md` stehen auf `offen`, kein neuer Beleg
  wurde angelegt — konsistent mit §8 und §9 des Plans.
- geprüft, ohne Befund: die Kern-Behauptung von §2.3 — der Arbeitsteilungs-Satz
  „Diese Tabelle **listet auf**; definiert wird hier nichts" steht in den
  historisch vendorten Bäumen `v5.12.0`, `v6.0.0`, `v6.2.0` und im aktuellen
  `v6.5.0` **wortgleich** an derselben Zeilennummer und fehlt in `v3.5.2`.
  „Seit `v5.12.0` unverändert in der Ziel-Form" ist damit gegen die
  Repo-Historie belegt, und die Herleitung des Slice steht.
- geprüft, ohne Befund: die ersten **zwei** der drei neuen HIGH-Kategorien —
  „Kommentar trägt keine der Kommentar-Klassen" und „Zustandsfeld trägt
  Chronik" sind gegenüber `v6.5.0` ·
  `templates/.harness/skills/reviewer.template.md` inhaltlich vollständig
  übersetzt, einschließlich des Nachsatzes „Kein Gate fängt das" und mit
  zusätzlichem Repo-Beleg. Nur die dritte trägt nicht (F-3).
- geprüft, ohne Befund: `docs/plan/carveouts/README.md` — von den Kandidaten
  der Ziel-Form ist der richtige übernommen; die übrigen Normsätze
  („Aufgelöste Carveouts wandern nach `done/`", „Jeder aktive Carveout
  braucht: Trigger, Folge-Slice, letzten Prüf-Termin", „Bei Welle-Closure:
  Carveout-Audit zwingend") sind in §Form eines Carveouts bzw. §Audit
  bereits gedeckt.
- geprüft, ohne Befund: `slice-186` **nimmt den Punkt an**, den `slice-185` §7
  Risiko 3 ihm übergibt — §1 nennt die neun Paare als Ziel, schließt sie nicht
  aus, und der Start-Trigger („`slice-185` liegt in `done/`") ist beobachtbar
  und kein Ergebnis der eigenen Arbeit. Die Kandidaten-Summe stimmt: 155 aus
  seiner Tabelle plus 14 aus `slice-185` §3 ergeben die genannten 169 über 13
  Paare.
- geprüft, ohne Befund: die Folge-Slice-Paarung — `slice-186` existiert als
  Datei im Planning-Lifecycle (`open/`).
- geprüft, ohne Befund: der `Derivativ`-Absatz in `docs/plan/adr/README.md`
  behauptet kein Gate falsch — `make gate-consistency` prüft tatsächlich die
  ADR-Index-Vollständigkeit, und `make doc-check` die Gegenrichtung; beide
  Targets existieren und sind in `AGENTS.md` §4 gelistet.
- geprüft, ohne Befund: Kopf-Metadaten von `slice-185` und `slice-186`
  (`Welle:`, `Bezug:`, `Berührte Spec-Stellen:`, `Verantwortlich:`, `Autor:`,
  `Lerneintrag — Form:`), DoD-Größe, Closure-Struktur und Risiko-Ausgänge —
  strukturell konsistent; jedes der drei Risiken trägt genau einen Ausgang aus
  der geschlossenen Menge.
- **Nicht geprüft** (Geltungsbereich dieses Laufs): die neun Paare, die
  `slice-186` übernimmt, sind nicht abgeglichen worden — dieser Report prüft
  die **fünf** bearbeiteten Paare und die Methode, nicht den Rest der Menge.
  Ebenso ungeprüft: ob die maschinelle Vorauswahl in den fünf bearbeiteten
  Paaren Kandidaten *verfehlt* hat, die a-check in anderen Worten trägt — dazu
  fehlt das Instrument (F-6). F-4 und F-8 sind zwei Fälle, die ein Lesen der
  Ziel-Form gefunden hat; sie belegen, dass die Menge nicht erschöpft ist,
  nicht, dass sie es jetzt wäre.

## Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 4 |
| MEDIUM | 5 |
| LOW | 3 |
| INFO | 0 |

Zwölf Findings; F-12 liegt am Rand des Gegenstands und ist dort als solches
gekennzeichnet.

**Finding-Klassen dieses Laufs:** Ziel-Form-Paar gegen den Verzeichnis-Index
statt gegen die Vorlage gemessen · Teil-Messung als Deckungs-Nachweis für eine
mehrteilige Zusage ausgegeben · Ziel-Form-Regel mit einer falschen
Repo-Tatsache entschärft · Vollständigkeits-Haken gesetzt, ohne den Gegenstand
zu erschöpfen · Anker-Paarung im Indikativ behauptet, bevor der Zielort sie
trägt · Messung als Beleg ohne reproduzierbares Instrument · Grundgesamtheit
aus einem Glob abgeleitet, dessen Menge von der Arbeitsmenge abweicht ·
Kandidat aus dem Bedienhinweis übernommen, Kandidat aus dem Normtext übergangen
· Zielsatz nach Plan-Änderung nicht nachgezogen · Ausschluss mit unzutreffender
Begründung · Kandidaten-Klassifikation gröber als der Kandidat ·
Schwellenwert gesenkt, Meldungstext nicht nachgezogen

**Wiederkehrende Klasse über diesen Lauf hinweg:** sieben der zwölf Findings
(F-1, F-2, F-4, F-6, F-7, F-8, F-11) sind dieselbe Bewegung — **eine Messung
wird weiter getragen, als ihr Gegenstand reicht**. Das ist die Klasse, für die
`AGENTS.md` §5 seit `slice-179` die Geltungsbereichs-Frage führt und
`BEO-PLAN/review-geltungsbereich-zu-eng` den Zähler.

## Verdikt

**Merge-blockierend: ja.** Die vier HIGH betreffen nicht die Idee des Slice —
die trägt, und der Nachweis für sie (§2.3, Arbeitsteilungs-Satz seit `v5.12.0`)
ist gegen die Repo-Historie unabhängig bestätigt. Sie betreffen die
**Ausführung der Messung**, und dieser Slice macht Messgenauigkeit zu seinem
Gegenstand: Er verankert eine Zahl (14 Paare) als Repo-Konvention in
`harness/conventions.md`, verändert mit dem Reviewer-Skill die Urteilsgrundlage
jedes künftigen Reviews, und übergibt seine Menge an `slice-186`. Ein Fehler in
der Grundgesamtheit (F-1, F-7) und eine mit einer falschen Repo-Tatsache
entschärfte Ziel-Form-Regel (F-3) wandern damit weiter, statt hier zu enden.

**Vier Findings haben Reichweite über `slice-185` hinaus:** F-1 und F-7
betreffen die Menge, die `slice-186` laut seinem §1 abarbeitet — die
Ziel-Form-Zuordnung von `README.md` und das fehlende `roadmap`-Paar. F-3 und
F-4 betreffen `.harness/skills/reviewer.md`, also die Datei, gegen die dieser
Report selbst geschrieben wurde.

**Übergabe:** Reviewer → Implementer. Dieser Report ist Lauf-Beleg, keine
Verifikation — DoD-/Spec-Konformität prüft der Verifier separat (Modul 11).
Kein Finding enthält einen Lösungsvorschlag; die Entscheidung, welcher Ausgang
je Finding gilt, liegt beim Implementer und, wo eine Ziel-Form-Regel bewusst
nicht übernommen werden soll, bei einer `MR`-Adaption statt bei einem Nachtrag
(so benennt es `slice-185` §5 selbst).
