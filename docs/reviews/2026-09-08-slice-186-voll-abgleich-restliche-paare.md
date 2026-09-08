# Review-Report: slice-186 — 2026-09-08

**Review-Art:** Plan + Diff — geprüft gegen den Slice-Plan, die Ziel-Formen des
vendorten Stands und die berührten Repo-Artefakte (Modul 10 §Drei Review-Arten).
Kern-Behauptungen dieses Slice sind ein **Umzug** („der Zielort trägt die Regel
vollständig") und mehrere **Messungen** (Zeichenzahlen, Kandidaten-Zahlen,
Wortfolgen-Treffer). Der Review zieht den Umzug nach und misst nach.

**Gegenstand:** Commits `1eba737` („AGENTS.md Paragraph 5 -- drei Befunde
aufgeloest"), `3c089c4` („Rollen-Zeile aus der Ziel-Form uebernommen") und
`a32d654` („slice-186 abgeglichen, Closure und Folge-Slice"); Range
`origin/main..HEAD`, Working Tree sauber.

**Skill:** `.harness/skills/reviewer.md` @ Stand `1eba737` — **einschließlich
des §Mess-Regeln-Abschnitts, den dieser Slice dort angelegt hat**; die beiden
Regeln werden damit auf ihren eigenen Umzug angewendet.
**Modell:** claude-opus-5 (1M) · **Datum:** 2026-09-08

**Eingangs-Kontext:**

- `slice-186` (Plan, `in-progress/`), `slice-187` (Folge-Slice, `open/`),
  `slice-185` (Stub plus Volltext aus `slice-185-archiv.zip`)
- `AGENTS.md` §3.1/§3.2/§3.3/§3.5/§3.7, §4, §5, §6 — Stand vor und nach dem Diff
- `harness/README.md` §Source precedence, §Guides, §Sensors ·
  `harness/conventions.md` §Baseline, §Modus-Deklaration
- `.harness/skills/reviewer.md`, `.harness/skills/closure-note-reviewer.md`
- `harness/sensors/doc-planning.md`, `.d-check.yml` (Module `planning`, `ids`)
- `v6.5.0` · `templates/project-readme.template.md`,
  `templates/docs/plan/planning/roadmap.template.md`,
  `templates/harness/README.template.md`,
  `templates/docs/plan/planning/slice.template.md`
- `v6.5.0` · `regelwerk/modul-05-planning-harness.md`,
  `regelwerk/modul-06-roadmap.md`, `regelwerk/modul-08-agentenrollen.md`
- Beobachtungs-Register: `BEO-PLAN/erstdurchgang-als-einmal-slice-geschnitten`,
  `BEO-PLAN/form-vergleich-sprachblind`, `BEO-PLAN/review-geltungsbereich-zu-eng`,
  `BEO-GATE/muster-trifft-nur-die-haeufige-schreibweise`,
  `BEO-GATE/probe-liefert-den-gegenstand-mit`,
  `BEO-HARNESS/baseline-normtext-nachgeschrieben`
- früherer Report zum selben Bereich: `slice-186`s direkter Vorgänger `slice-185`
  (Findings F-4, F-6, F-9, F-11 sind hier auf Wiederholung geprüft)

**Geltungsbereich der eigenen Messungen** (Skill §Mess-Regeln, erste Regel):

- Zeichenzahlen sind **Zeichen**, nicht Bytes (`wc -m`, UTF-8), Abschnittsgrenze
  von der Überschrift `## 5.` bis ausschließlich `## 6.`.
- Der eigene Wortfolgen-Test lief über `regelwerk/*.md` des vendorten Stands,
  normalisiert (Links auf ihren Text reduziert, Markup und Anführungszeichen
  entfernt), Fenstergrößen 6/7/8/10 Wörter. **Sinngemäße Übernahme in anderen
  Worten sieht er nicht** — dieselbe Grenze, die der Slice in §2.1 für seinen
  eigenen Test benennt.
- Der Ziel-Form-Abgleich der drei Paare ist auf **H2-Ebene** vollständig und
  darunter stichprobenweise; er ersetzt kein Wort-für-Wort-Lesen.
- **`make gates` / `make verify` habe ich nicht ausgeführt** (Auftrag: Belege
  liegen vor). Alle Aussagen unten stützen sich auf Dateien, nicht auf
  Gate-Läufe.

---

## Findings

### F-1 — Die zwei Mess-Regeln stehen jetzt in einem Dokument, das ihre Adressaten laut Repo-Aussage nicht lesen

- `kategorie`: HIGH
- `quelle`: `harness/README.md` §Guides, Zeile 44; `v6.5.0` ·
  `regelwerk/modul-08-agentenrollen.md` §Welche Rolle braucht welche
  Artefaktklasse; `AGENTS.md` §3.7
- `pfad`: `.harness/skills/reviewer.md`:112–172 (neuer Abschnitt §Mess-Regeln);
  `AGENTS.md`:306–313 (Zeiger); Begründung in
  `docs/plan/planning/in-progress/slice-186-voll-abgleich-restliche-paare.md`
  §2.1 (b)
- `befund`: Beide Regeltexte nennen ihren Adressaten selbst — *„Wer eine Messung
  als Beleg schreibt — in einem Slice-Plan, einer Closure-Notiz, einem
  Review-Report"* und zweimal *„Was greift, ist die Frage beim **Schreiben**"*.
  Zwei der drei genannten Artefakte schreibt nicht der Reviewer. Der Zielort ist
  aber genau das Dokument, das `harness/README.md` §Guides in derselben Zeile als
  *„nicht Teil der Implementer-Eingabe"* führt; `AGENTS.md` §6 lädt es erst nach
  Schritt 8. Der Zeiger in §5 nennt die zwei Regeln beim Namen, trägt aber ihren
  Satz nicht mehr — nach dem Umzug steht die Zusage *„nenne den Geltungsbereich
  deiner Messung"* in keinem Dokument, das ein Implementer- oder Planner-Lauf
  liest. Die Begründung des Slice (*„`modul-08` weist genau diesen Fall der
  Skill-Datei zu"*) hält daran nicht: `modul-08` ordnet die Skill-Datei dem
  **Reviewer-Urteil** zu, nicht jedem inferentiellen Satz — und das Repo
  widerlegt die Ableitung selbst, weil `AGENTS.md` §3.7 ausdrücklich
  *„inferentiell"* ist und trotzdem im Briefing steht.
- `verifizierbar`: nein — Rollen-Zuordnung ist ein Urteil; die drei Einzelfakten
  (Adressaten-Satz, `nicht Teil der Implementer-Eingabe`, §3.7 bleibt) sind es je
  einzeln.
- `klasse`: Regel an einen Ort umgezogen, den ihr Adressat nicht liest

### F-2 — Vier Selbstverweise im umgezogenen Block zeigen nach dem Umzug ins Leere oder auf die falsche Datei

- `kategorie`: HIGH
- `quelle`: Reviewer-Skill §Klassifikation, „nachweislich falsche
  Tatsachenbehauptung"; `slice-186` §7, Risiko 2, Ausgang *entfallen*
  (*„der Zielort trägt den Text unverändert"*)
- `pfad`: `.harness/skills/reviewer.md`:126, :145, :150, :152, :169
- `befund`: Der Block ist wortgleich verschoben — genau das ist der Defekt: Er
  trug vier Verweise, die sich auf **`AGENTS.md`** bezogen. `(§3.7)` steht jetzt
  zweimal unqualifiziert in einer Datei ohne nummerierte Abschnitte, während
  jeder andere §-Verweis derselben Datei mit `AGENTS.md` qualifiziert ist
  (:6, :21, :27, :35, :51, :60). *„an §5 **dieser Datei** gemessen"* (:145) und
  *„drei von 18 Blöcken in §5"* (:152) benennen eine Messung an `AGENTS.md` §5,
  behaupten sie aber über den Reviewer-Skill, der kein §5 hat. *„ist die Aufgabe
  **dieser Datei**"* (:150) sagte über das Briefing etwas, das über die
  Skill-Datei nicht stimmt. Der Risiko-Ausgang *entfallen* und das zweite
  Closure-Kriterium (*„tragen den umgezogenen Text — nachzählbar, nicht
  behauptet"*) sind auf **Textgleichheit** geprüft; die Aussagen-Gleichheit, um
  die es im Risiko ging, ist es nicht.
- `verifizierbar`: nein — kein Gate prüft, ob ein `§`-Verweis im Fließtext
  auflöst; `make doc-check` sieht nur Links und Anker.
- `klasse`: Selbstverweis beim Umzug nicht re-verankert

### F-3 — Der Herkunfts-Anker `seit slice-183` steht nach dem Umzug nicht mehr am deklarierten Zielort

- `kategorie`: HIGH
- `quelle`: `v6.5.0` · `regelwerk/modul-06-roadmap.md` §Das
  Beobachtungs-Register (Ausgang *verkörpert*: „Zielort **und**
  Herkunfts-Anker") und §Wellen-Closure-Prozedur, Paarung (a)
- `pfad`:
  `docs/plan/planning/observations/BEO-HARNESS/baseline-normtext-nachgeschrieben/state.md`:1
  und :16–17; Zielort `AGENTS.md` §5
- `befund`: Das `state.md` führt den Stand *„verkörpert in `AGENTS.md` §5
  (Geltungsbereich einer Messung, dritte Hälfte) `seit slice-183`"* und
  wiederholt ihn unten. Nach dem Umzug kommt die Zeichenfolge `seit slice-183`
  in `AGENTS.md` nicht mehr vor — sie steht nur noch in
  `.harness/skills/reviewer.md`:145 und im `state.md` selbst. Der neue Zeiger in
  §5 trägt `seit slice-179` und `seit slice-181`, nicht `seit slice-183`. Zwei
  weitere Register-Stände (`BEO-PLAN/review-geltungsbereich-zu-eng`,
  `BEO-GATE/probe-liefert-den-gegenstand-mit`) nennen ebenfalls `AGENTS.md` §5
  als Zielort; dort steht die **Regel** nicht mehr, nur ihr Name. `state.md` ist
  laut derselben Regelwerk-Stelle der *veränderliche* Teil des Eintrags — ein
  Umzug des Zielorts gehört hinein.
- `verifizierbar`: ja — `grep -n "seit slice-183" AGENTS.md` liefert nichts,
  `grep -rn "seit slice-183" --include='*.md' .` liefert zwei Treffer, keiner in
  `AGENTS.md`. `make verify-observations` prüft Verzeichnis- und
  Evidence-Deckung, nicht den Zielort.
- `klasse`: Herkunfts-Anker nach Umzug nicht nachgezogen

### F-4 — Das als „ohne Befund" geschlossene Roadmap-Paar trägt einen falschen Zustandssatz

- `kategorie`: HIGH
- `quelle`: `v6.5.0` · `templates/docs/plan/planning/roadmap.template.md`
  §Offene Wellen (*„woran gearbeitet wird, sagt das `Welle:`-Feld der Slices in
  `in-progress/`"*); `v6.5.0` · `regelwerk/modul-06-roadmap.md`
  §Roadmap-Struktur
- `pfad`: `docs/plan/planning/in-progress/roadmap.md`:24; Ausgang in
  `docs/plan/planning/in-progress/slice-186-voll-abgleich-restliche-paare.md`:140
- `befund`: Die Sektion *Offene Wellen* sagt im Präsens *„In Arbeit: slice-169 —
  wellenlos"*. `slice-169` liegt in `done/wellenlos/` (der Link im Satz zeigt
  selbst dorthin); in `in-progress/` liegt `slice-186`. Die Zeile stammt aus dem
  Übergangs-Commit `3f118e5` und ist seither durch rund siebzehn Slice-Übergänge
  unverändert mitgereist, obwohl `make slice-mv` die Datei jedes Mal anfasst.
  Der Slice führt genau dieses Paar als **ohne Befund** und begründet das mit
  *„alle fünf Abschnitte der Ziel-Form vorhanden"* — ein Vorhandensein von
  Überschriften, das über den Inhalt einer Sektion nichts sagt. Der falsche Satz
  steht in der Sektion, deren Ziel-Form-Regel er verletzt.
- `verifizierbar`: ja — `ls docs/plan/planning/in-progress/` gegen den Satz;
  `git log -S "In Arbeit: [slice-169]"` liefert genau `3f118e5`.
- `klasse`: Vollständigkeits-Haken gesetzt, ohne den Gegenstand zu erschöpfen

### F-5 — `README.de.md` fehlt eine Sektion der Ziel-Form; das Paar ist mit einem einzeiligen Ausgang geschlossen

- `kategorie`: MEDIUM
- `quelle`: `v6.5.0` · `templates/project-readme.template.md` (Sektion *Was kann
  ich heute tun?* mit eigener Regel-Zeile); `BEO-PLAN/form-vergleich-sprachblind`
- `pfad`: `README.de.md` (Sektionsfolge) gegen `README.md`:40;
  `slice-186` §3.2 und §9, Zeile `form-vergleich-sprachblind`
- `befund`: Die Ziel-Form führt fünf inhaltliche Sektionen. `README.md` trägt
  alle fünf; `README.de.md` springt von *Was ist a-check?* direkt zu *Warum
  a-check?* — die Sektion *Was kann ich heute tun?* fehlt dort seit `slice-111`,
  das sie nur der englischen Fassung hinzufügte. Der Slice hat `README.de.md`
  angefasst und die Übernahme ausdrücklich *„in beiden Sprachfassungen"*
  ausgeführt, den Befund-Scope aber auf die englische Fassung begrenzt, ohne das
  zu sagen. §9 führt `form-vergleich-sprachblind` als *„berührt, nicht erhöht …
  und traf"* — für die deutsche Zwillingsdatei traf er nicht; das ist der zweite
  Beleg derselben Klasse, und der Eintrag steht weiter bei 1×. Der Repo-Bestand
  behandelt beide READMEs als zu pflegendes Paar (`docs/user/releasing.md`,
  `version.md`), eine Teil-Übersetzung ist nirgends deklariert.
- `verifizierbar`: ja — `grep -n "^## " README.md README.de.md`.
- `klasse`: zweite Sprachfassung im Übernahme-Scope, nicht im Befund-Scope

### F-6 — Der Wortfolgen-Test nennt seinen Geltungsbereich nicht; seine Zahlen reproduzieren nicht

- `kategorie`: MEDIUM
- `quelle`: Reviewer-Skill §Mess-Regeln, *Geltungsbereich einer Messung* — die
  Regel, die dieser Slice umzieht; Wiederholung der Klasse aus dem
  `slice-185`-Report (F-6)
- `pfad`: `slice-186` §2.1 (a) (*„findet **fünf** wörtliche Treffer, alle im
  selben Block … **Die anderen 17 Blöcke: null Treffer**"*)
- `befund`: Weder Fenstergröße noch Korpus noch Normalisierung sind genannt, und
  das Instrument liegt nicht im Repo. Nachgemessen über `regelwerk/*.md` liefert
  ein 8- oder 10-Wort-Fenster **drei** Treffer, alle in Block 13 — die Aussage
  *„17 Blöcke null"* hält dort. Ein 6- oder 7-Wort-Fenster liefert **sechs** bzw.
  **vier** Treffer in Block 13 **und je einen in Block 18** (*„nicht ‚fertig',
  sondern nur ‚weg'"* gegen `modul-01-entwicklungszyklus.md`). Die genannte Zahl
  fünf und die Aussage über die 17 anderen Blöcke gelten also nur unter einem
  Parameter, den der Slice nicht angibt; mit dem Korpus `templates/` erweitert
  kommen zwei weitere Blöcke dazu. Dieselbe Lücke trifft die Kandidaten-Zahlen
  (5 / 4 / 9 / 20 / 17 / 16 / 16 / 15 / 12 / 11 und die Summe 107) — sie stammen
  aus einem Lauf, den niemand wiederholen kann.
- `verifizierbar`: nein — das Instrument fehlt; genau das ist der Befund. Die
  Gegenmessung ist reproduzierbar, aber nur unter *ihren* Parametern.
- `klasse`: Messung als Beleg ohne reproduzierbares Instrument

### F-7 — Der dokumentierte `doc-planning`-Vertrag sagt mehr zu, als das Modul prüft; die Grenze steht nur im Konfigurations-Kommentar

- `kategorie`: MEDIUM
- `quelle`: Reviewer-Skill §Klassifikation, HIGH-Kategorie „Norm nur im
  Template-Kommentar", zweite Ausprägung (`.d-check.yml`); `AGENTS.md` §4
- `pfad`: `AGENTS.md` §4 (Zeile `make doc-planning`); `harness/README.md`
  §Sensors (Zeile `make doc-planning`); `harness/sensors/doc-planning.md`
  §Vertrag und §Grenze; `.d-check.yml`, Block `planning:`
- `befund`: Alle drei gelesenen Stellen sagen zu, dass die Roadmap-Sektion den
  Slice aus `in-progress/` **benennt**. Der Kommentar in `.d-check.yml` sagt das
  Gegenteil und nennt es „BENANNTE GRENZE": geprüft ist nur die Äquivalenz
  *Slice vorhanden ⟺ Ruhe-Marker steht nicht*, nicht ob der genannte Slice der
  vorhandene ist. Die Sensor-Datei führt einen eigenen Abschnitt §Grenze und
  listet dort zwei andere Lücken — diese nicht. Damit steht eine **Zusage** in
  der Doku, deren Einschränkung ausschließlich in einem Konfigurations-Kommentar
  lebt; F-4 ist die Folge, die siebzehn Slices lang grün blieb. Der Befund liegt
  am Rand des Diffs: Der Slice hat ihn nicht erzeugt, aber das Paar
  `harness/README.md` als *ohne Befund* geschlossen.
- `verifizierbar`: ja — die vier Textstellen nebeneinander; ein Lauf bestätigt
  zusätzlich, dass der Zustand aus F-4 grün ist.
- `klasse`: Zusage in der Doku, Grenze nur im Konfigurations-Kommentar

### F-8 — Titel, §1 und das Kopffeld sind der Plan-Änderung nicht gefolgt

- `kategorie`: LOW
- `quelle`: `AGENTS.md` §6, Schritt 4; Wiederholung der Klasse aus dem
  `slice-185`-Report (F-9) im unmittelbar folgenden Slice
- `pfad`: `slice-186`:1 (Titel), :23 (§1 Ziel), :10 (Kopffeld *Berührte
  Spec-Stellen*)
- `befund`: Der Titel sagt *„die zehn verbleibenden Ziel-Form-Paare"*, §1 sagt
  *„Die zehn Ziel-Form-Paare … sind abgeglichen"*; geliefert sind drei, und der
  DoD-Punkt benennt die Plan-Änderung ausdrücklich. Ziel und DoD tragen damit
  verschiedene Zusagen — derselbe Befund, den der Vorgänger-Report eine Woche
  vorher an derselben Stelle erhoben hat. Das Kopffeld trägt zusätzlich einen
  Bedienhinweis, der bei Closure nicht mehr zutrifft: *„bei Umsetzung zu füllen:
  `spec/lastenheft.md` und `spec/spezifikation.md` sind unter den Paaren"* —
  §9 stellt fest, dass `SPEC` gerade **nicht** berührt ist und die zwei Straten
  zu `slice-187` gehen.
- `verifizierbar`: nein — Lesart, kein Match.
- `klasse`: Zielsatz nach Plan-Änderung nicht nachgezogen

### F-9 — Geschätzte Zahlen stehen unmarkiert neben exakt reproduzierbaren

- `kategorie`: LOW
- `quelle`: Reviewer-Skill §Mess-Regeln, *Geltungsbereich einer Messung*
- `pfad`: `slice-186` §3.1, Tabelle (Spalte *Nachher*: „Zeiger (~430)",
  „Zeiger (~440)"); §2.1 (a) („~600 Zeichen")
- `befund`: Die Kopfzahlen der Tabelle stimmen zeichengenau (siehe
  Negativbefunde). Die drei Tilde-Werte tun das nicht: Der Mess-Regel-Zeiger
  misst **543**, der Slice-Form-Zeiger **539** Zeichen, und (a) hat nicht ~600,
  sondern **834** Zeichen entfernt (4285 → 3451). Die Abweichung ist folgenlos —
  die Gesamtrechnung 8410 + 539 + 543 geht auf —, aber im selben Absatz stehen
  Schätzung und Messung ununterscheidbar nebeneinander.
- `verifizierbar`: ja — Abschnitts-Extraktion und `wc -m`.
- `klasse`: Schätzung und Messung im selben Beleg nicht unterschieden

### F-10 — „alle fünf Abschnitte der Ziel-Form" — die Ziel-Form führt sechs

- `kategorie`: LOW
- `quelle`: `v6.5.0` · `templates/docs/plan/planning/roadmap.template.md`;
  Wiederholung der Klasse aus dem `slice-185`-Report (F-11)
- `pfad`: `slice-186`:140 (Tabellenzeile `roadmap.md`)
- `befund`: Die Zahl fünf stammt aus `regelwerk/modul-06-roadmap.md`
  §Roadmap-Struktur; die **Ziel-Form** trägt sechs H2-Abschnitte — der
  *Abhängigkeitsgraph* kommt hinzu. Der Ausgang *ohne Befund* ist trotzdem
  richtig (`roadmap.md` trägt alle sechs), die Begründung deckt den Gegenstand
  aber nicht: Sie zählt am Regelwerk, wo sie an der Vorlage zählen müsste.
- `verifizierbar`: ja — `grep -n "^## "` auf beide Dateien.
- `klasse`: Kandidaten-Klassifikation gröber als der Kandidat

### F-11 — Zwei Zeiger auf `AGENTS.md` §5 verweisen auf ausgezogenen Text

- `kategorie`: LOW
- `quelle`: `AGENTS.md` §1 (Pointer-Disziplin)
- `pfad`: `.harness/skills/reviewer.md`:51; `harness/README.md` §Guides, Zeile 44
- `befund`: Die HIGH-Kategorie „Norm nur im Template-Kommentar" belegt sich mit
  *„`AGENTS.md` §5, ‚Beim Kopieren anzupassen'"* — die sieben Punkte stehen seit
  diesem Commit in `docs/plan/planning/README.md`; §5 trägt nur noch den Zeiger.
  Die Aussage bleibt über den Zeiger auflösbar, die Fundstelle stimmt nicht mehr.
  Dieselbe Zeile in `harness/README.md` §Guides beschreibt den Reviewer-Skill
  weiter als *„HIGH-Liste, Kategorien-Regeln, Negativbefund-Pflicht,
  Output-Schema"* — der mit diesem Commit hinzugekommene §Mess-Regeln fehlt in
  der Beschreibung.
- `verifizierbar`: ja — beide Zeilen gegen den neuen `AGENTS.md`-§5.
- `klasse`: Zeiger auf eine ausgezogene Stelle

### F-12 — Erstauftreten für einen bereits geschlossenen Vorgang nachgetragen

- `kategorie`: INFO
- `quelle`: `v6.5.0` · `regelwerk/modul-06-roadmap.md` §Das
  Beobachtungs-Register
- `pfad`:
  `docs/plan/planning/observations/BEO-PLAN/erstdurchgang-als-einmal-slice-geschnitten/evidence/slice-185.md`
- `befund`: Der Beleg für `slice-185` ist bei der Closure von `slice-186`
  entstanden — das Register kennt nur *„Erstauftreten benennt und vergibt die
  Kennung"*, nicht den Nachtrag. Beide urteilsfreien Prüfungen bestehen (der
  Dateiname ist die Kennung eines abgeschlossenen Vorgangs; `slice-185` liegt in
  `done/`), der abgeleitete Zähler bleibt korrekt, und die Datei sagt selbst,
  dass sie nachgetragen ist. Kein Handlungsbedarf; genannt, weil es der erste
  Fall im Repo ist und ein Nachtrag den Zähler in einer Closure um zwei hebt.
- `verifizierbar`: ja — `git log` auf die Datei gegen die Lage von `slice-185`.
- `klasse`: Beleg für einen abgeschlossenen Vorgang nachgetragen

## Negativbefunde

- geprüft, ohne Befund: **Die Kern-Messung reproduziert zeichengenau.**
  `AGENTS.md` §5 misst vorher **16 740**, nachher **9491** Zeichen — 43,3 %,
  exakt die genannten Werte. Ebenso exakt: der Slice-Form-Block **4285**, die
  zwei Mess-Regeln **4045** (2688 + 1357), die 18 Blöcke in §5. Die Rechnung
  geht auf: 16 740 − 4285 − 4045 = 8410 Rest, plus 539 + 543 Zeiger = 9492
  gegen gemessene 9491 (eine Zeilenumbruch-Differenz).
- geprüft, ohne Befund: **Block (b) ist wortgleich umgezogen.** Ein Zeilen-Diff
  zwischen dem alten `AGENTS.md`-Text und `.harness/skills/reviewer.md`
  §Mess-Regeln zeigt genau drei Unterschiede, alle drei Link-Pfade eine Ebene
  tiefer. Kein Satz fehlt, keiner ist hinzugekommen. (Was daran trotzdem ein
  Befund ist, steht in F-2.)
- geprüft, ohne Befund: **Block (c) ist wortgleich umgezogen und trägt keine
  hängenden Selbstverweise.** Der Block misst am Zielort 3469 Zeichen gegen 3451
  vorher — die Differenz von 18 sind genau die zwei Pfade, die drei Ebenen tiefer
  zeigen. Die §-Verweise darin (`§1`, `§7 bis §10`, `slice-165 §3/§4`) meinen
  Slice-Plan-Abschnitte bzw. sind qualifiziert und bleiben gültig.
- geprüft, ohne Befund: **Die Kürzung (a) verliert keine repo-eigene Zusage.**
  Was gestrichen ist, steht in `regelwerk/modul-05-planning-harness.md`
  §Ziel-Form: Slice — einschließlich der Nicht-Zählung des Review-Reports. Was
  stehen bleibt, ist die Durchsetzung (`make verify` ab `slice-052`, Gate-Lauf
  als feste Zeile, Kopffelder ab `slice-098`); dafür führt das Regelwerk kein
  Gegenstück.
- geprüft, ohne Befund: **Risiko-Ausgänge — Form.** Drei notierte Risiken, drei
  Ausgänge, alle aus der geschlossenen Dreier-Menge
  (`modul-05-planning-harness.md` §Offene Risiken werden bei Closure aufgelöst).
  *entfallen* trägt eine Begründung. (Ob der zweite Ausgang inhaltlich trägt:
  F-2.)
- geprüft, ohne Befund: **Die Adresse nimmt die Sendung an.** `slice-187` §1
  nennt genau die sieben abgetrennten Paare als Ziel, schließt sie nicht aus,
  trägt sie in §2 mit Messung und in §5 mit einer eigenen Rückführung. Sein
  Start-Trigger (`slice-186` in `done/`) ist kein Ergebnis seiner selbst.
- geprüft, ohne Befund: **Die DoD-Umschreibung ist keine Erfindung dieses
  Slice.** `slice-185` führt dieselbe Form wörtlich (*„Plan-Änderung, benannt
  statt stillschweigend"*, Volltext im Archiv) — dieser Slice ist der zweite Fall
  einer im Repo bereits praktizierten Offenlegung, und der Ausgang *eingetreten →
  Folge-Slice mit Kennung* ist der von `modul-05` vorgesehene. Die Kritik
  beschränkt sich auf F-8 (Titel/§1/Kopffeld).
- geprüft, ohne Befund: **Register-Form.** Der neue Eintrag hat die drei
  Bestandteile mit den drei Lebensdauern; der Zähler ist abgeleitet (zwei
  Evidence-Dateien), es gibt kein gespeichertes Zähler-Feld. `state.md` nennt
  Stand und Beleg statt Chronik und deklariert von sich aus den Geltungsbereich
  der Zählung („eine Untergrenze, kein Vollstand") — das ist die Mess-Regel auf
  den eigenen Eintrag angewandt.
- geprüft, ohne Befund: **Die Nicht-Erhöhung von
  `BEO-GATE/muster-trifft-nur-die-haeufige-schreibweise` trägt.** Deren
  `observation.md` definiert die Klasse ausdrücklich über *Schweigen* („erzeugt
  keinen Fehlalarm, sondern Schweigen"); die hiesige Fehlrichtung ist Rauschen.
  Ein Beleg dort hätte einen Zähler auf 3× gehoben, dessen Regel den Fall nicht
  trifft. Die richtige Kennung für die tatsächlich aufgetretene Schweige-Richtung
  ist `form-vergleich-sprachblind` — siehe F-5.
- geprüft, ohne Befund: **Die Rang-Abweichung in der Rollen-Zeile ist korrekt.**
  Die Ziel-Form sagt Rang 6, a-check trägt neun Ränge und `README.md` steht auf
  **7** — bestätigt durch `harness/README.md` §Source precedence und
  `AGENTS.md` §2. Der Anker `#source-precedence` löst auf. Die Aussage der
  Ziel-Form (verweist, dupliziert nicht) ist übernommen, die Zahl richtig
  angepasst.
- geprüft, ohne Befund: **Sektions-Deckung der zwei „ohne Befund"-Paare auf
  H2-Ebene.** `roadmap.md` trägt alle sechs Abschnitte der Ziel-Form,
  `harness/README.md` alle acht plus einen eigenen H2 (*Rollen und ihre
  Übergabe-Artefakte*) und einen eigenen H3 (*Nicht-Gates*). (Was der Inhalt
  einer dieser Sektionen sagt: F-4.)
- geprüft, ohne Befund: **Closure-Notiz, semantisch** (Maßstab
  `.harness/skills/closure-note-reviewer.md`). Alle drei Inhalte sind da:
  Lernsignal mit Ursache (*„Ein Abschnitt wird nicht groß, weil er viel regelt,
  sondern weil er die Herleitung seiner Regeln nachschreibt"*; *„Nicht die Slices
  sind zu groß, sondern der Gegenstand ist keine Slice-Größe"*), ein konkretes
  Folge-Slice mit auffindbarer Datei (`slice-187` in `open/`), und eine
  nachprüfbare Beobachtung (die 9491 gegen 16 740, oben nachgerechnet). Keine
  Floskel, kein überholtes Futur. Die Notiz nennt auch, was **nicht** geliefert
  wurde — die Hälfte, die die Pflicht verlangt.
- geprüft, ohne Befund: **Commit-Hygiene.** Drei Commits, jeder nennt
  `slice-186`; `docs(planning)` berührt ausschließlich `docs/plan/planning/`;
  kein `git mv` mit gleichzeitiger Inhaltsänderung (der Lifecycle-Wechsel lag
  vorher); keine ADR angefasst; kein Produkt-Code, keine Suppression, keine
  Host-Toolchain.
- geprüft, ohne Befund: **Sub-Area-Prüfungen (§9).** Beide berührten Sub-Areas
  (`HARNESS`, `PLAN`) sind in `harness/conventions.md` §Modus-Deklaration
  geführt, beide Greenfield, der Begründungsblock entfällt zu Recht; der
  Sichtungs-Schritt ist ausgeführt und nennt fünf Einträge samt Nicht-Treffern.
  Die Abgrenzung gegen `SPEC` ist begründet.
- **Nicht geprüft, und darum keine Aussage:** `make gates` / `make verify` /
  `make doc-check` habe ich nicht ausgeführt. Ebenso ungeprüft bleibt, ob die
  drei Paare unterhalb der H2-Ebene wort-für-wort gegen ihre Ziel-Form gelesen
  wurden — F-4 und F-5 zeigen die zwei Stellen, an denen ich unterhalb der
  Überschriften nachgesehen habe, nicht mehr.

## Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 4 |
| MEDIUM | 3 |
| LOW | 4 |
| INFO | 1 |

**Finding-Klassen dieses Laufs:** Regel an einen Ort umgezogen, den ihr
Adressat nicht liest · Selbstverweis beim Umzug nicht re-verankert ·
Herkunfts-Anker nach Umzug nicht nachgezogen · Vollständigkeits-Haken gesetzt,
ohne den Gegenstand zu erschöpfen · zweite Sprachfassung im Übernahme-Scope,
nicht im Befund-Scope · Messung als Beleg ohne reproduzierbares Instrument ·
Zusage in der Doku, Grenze nur im Konfigurations-Kommentar · Zielsatz nach
Plan-Änderung nicht nachgezogen · Schätzung und Messung im selben Beleg nicht
unterschieden · Kandidaten-Klassifikation gröber als der Kandidat · Zeiger auf
eine ausgezogene Stelle · Beleg für einen abgeschlossenen Vorgang nachgetragen

**Wiederkehrend gegenüber dem Vorgänger-Report:** *Vollständigkeits-Haken
gesetzt, ohne den Gegenstand zu erschöpfen* (dort F-4, hier F-4/F-5), *Messung
als Beleg ohne reproduzierbares Instrument* (dort F-6, hier F-6), *Zielsatz nach
Plan-Änderung nicht nachgezogen* (dort F-9, hier F-8) und *Kandidaten-
Klassifikation gröber als der Kandidat* (dort F-11, hier F-10) — vier Klassen
zum **zweiten Mal in Folge**, alle vier an derselben Arbeitsform: ein Bestand
wird paarweise durchgegangen, und der Ausgang je Paar wird knapper notiert, als
der Gegenstand ist. Die zwei mit Substanz — *Vollständigkeits-Haken* und
*Messung ohne Instrument* — sind Kandidaten für das Beobachtungs-Register, bevor
`slice-187` denselben Vorgang ein drittes Mal fährt.

## Verdikt

**Merge-blockierend: ja**, für F-2 und F-3. Beide sind mechanische Folgen
desselben Umzugs und in wenigen Zeilen behebbar; sie stehen aber genau gegen die
Zusage, auf der Risiko 2 als *entfallen* geschlossen wurde — solange sie stehen,
trägt dieser Ausgang nicht. F-4 ist ebenfalls HIGH und in einer Zeile behebbar,
liegt aber außerhalb des Diffs; blockierend ist daran nur, dass das Paar als
*ohne Befund* geschlossen wurde.

**F-1 ist die inhaltlich schwerste Frage und keine Redaktions-Sache.** Ob die
zwei Mess-Regeln in die Skill-Datei gehören, ist eine Rollen-Entscheidung; sie
gegen `harness/README.md` §Guides zu treffen, ohne diese Zeile zu berühren, ist
es nicht. Der Weg dahin ist eine Entscheidung des Architect-Rolleninhabers, kein
Reviewer-Vorschlag — dieser Report kategorisiert nur.

**Übergabe:** an die Implementer-Rolle. Dieser Report ist Lauf-Beleg, keine
Verifikation — DoD- und Spec-Konformität prüft der Verifier separat (Modul 11).
