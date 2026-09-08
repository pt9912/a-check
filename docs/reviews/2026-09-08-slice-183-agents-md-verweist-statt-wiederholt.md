# Review-Report: slice-183 — 2026-09-08

**Review-Art:** Code/Doku-Review gegen Plan und Konventionen (Modul 10 §Drei
Review-Arten) — geprüft wird der Diff gegen §1 des Slice-Plans, gegen die Hard
Rules und gegen den vendorten Regelwerks-Stand, auf den der Diff neu verweist.

**Gegenstand:** Commits `b5a81d5` („drei Bloecke in AGENTS.md Paragraph 5 auf
Verweise") und `9cd11d8` („Pruef-Frage geschaerft, Register bei 3x")

**Skill:** `.harness/skills/reviewer.md` @ Stand `3fae6d3` · <!-- d-check:ignore -->
**Modell:** claude-opus-5[1m] · **Datum:** 2026-09-08

**Review-Art im Sinne der Kontext-Trennung:** **unabhängiger Lauf** — der
Reviewer hat weder Plan noch Diff verfasst, kein `fork`-Kontext.

**Eingangs-Kontext:**

- Slice-Plan `slice-183` (Stand `9cd11d8`)
- `AGENTS.md` §1, §3, §5, §6 in beiden Fassungen (`b5a81d5^` und `HEAD`)
- `v6.5.0` · `regelwerk/modul-05-planning-harness.md`,
  `regelwerk/modul-06-roadmap.md`, `regelwerk/modul-07-carveouts.md`,
  `regelwerk/modul-03-spec.md`
- `v6.5.0` · `templates/AGENTS.template.md` (Ziel-Form der Messung)
- `harness/conventions.md`, `harness/README.md` §Purpose,
  `docs/plan/carveouts/README.md`
- Beobachtungs-Register `BEO-HARNESS/baseline-normtext-nachgeschrieben`
- eigene Gate-Läufe: `make gates` (Exit 0), `make verify` (Exit 0)

**Reproduktion der Messung.** Die Methode aus §2 des Plans („Abschnitt von
`##`-Überschrift bis zur nächsten, HTML-Kommentare und Tabellenzeilen entfernt,
Zeichen gezählt") wurde nachgebaut und reproduziert **alle** Zahlen der
Ausgangsmessung exakt: §1 2211 · §2 782 · §3 3760 · §4 1051 · §5 16 822 · §6
2889, Ziel-Form 2714 / 718 / 3388 / 458 / 726 / 874, und die acht Blockgrößen
der Kandidaten-Liste (4285 · 2051 · 1057 · 1002 · 943 · 941 · 849 · 627). Die
Findings unten stützen sich auf **dasselbe** Messverfahren, nicht auf ein
anderes.

---

## Findings

### F-1 — Die drei „jetzt"-Werte der Kürzungs-Tabelle reproduzieren nicht und widersprechen der Gesamtzahl im selben Dokument

- `kategorie`: HIGH
- `quelle`: Selbstwiderspruch gegen die im selben Dokument genannte
  Gesamtkürzung (1095); Messmethode des Slice selbst (§2)
- `pfad`: `docs/plan/planning/in-progress/slice-183-agents-md-verweist-statt-wiederholt.md:110-112`;
  Commit-Message `b5a81d5`
- `befund`: Die Tabelle nennt „849 → 449", „1002 → 442", „627 → 316"; mit
  derselben Methode gemessen sind die neuen Blöcke **514 / 495 / 374** Zeichen
  lang. Die Differenzen summieren sich auf 1271, während das Dokument zweimal
  1095 als Gesamtkürzung von §5 nennt — 2478 − 1383 = 1095 geht auf, 2478 −
  1207 = 1271 nicht.
- `verifizierbar`: ja — `git show b5a81d5:AGENTS.md`, Block von der
  Bullet-Zeile bis zur nächsten `- `-Zeile, `wc -m` (dieselbe Zählung, die
  849/1002/627 auf die Zeichen genau trifft).
- `klasse`: Nachher-Wert geschätzt statt gemessen, während der Vorher-Wert
  gemessen war

### F-2 — „drei von 17 Blöcken" nennt den Nenner des Vorzustands; gemessen waren es 18

- `kategorie`: HIGH
- `quelle`: nachweislich falsche Tatsachenbehauptung, gegen `AGENTS.md` selbst
  verifiziert
- `pfad`: `AGENTS.md:376-377`; Slice-Plan Zeilen 203, 222;
  `.../baseline-normtext-nachgeschrieben/evidence/slice-183.md:17`
- `befund`: §5 trug zum Messzeitpunkt (`b5a81d5^`, 16 822 Zeichen) **18**
  Top-Level-Blöcke; 17 war der Stand bei `b25b121^` (15 464 Zeichen). Derselbe
  Slice weist die Differenz in §2 aus („§5 ist zwischen Anlage und Umsetzung um
  1358 Zeichen gewachsen") und hat die Zeichenzahl nachgezogen, den Nenner
  aber nicht — die Zahl steht jetzt als Regel-Beleg in `AGENTS.md` §5.
- `verifizierbar`: ja — `git show b25b121^:AGENTS.md` und
  `git show b5a81d5^:AGENTS.md`, §5 extrahieren, `grep -c '^- '`: 17 gegen 18;
  die zugehörigen Zeichenzahlen 15 464 und 16 822 belegen die Zuordnung.
- `klasse`: Bezugsgröße aus dem Vorzustand übernommen, während die Messgröße
  aktualisiert wurde

### F-3 — Größen-Faktor „22,9 ×" rechnet sich aus den eigenen Zahlen nicht

- `kategorie`: HIGH
- `quelle`: Selbstwiderspruch gegen die Messtabelle in §2 desselben Dokuments
- `pfad`: Slice-Plan Zeilen 52, 219; `AGENTS.md:377`;
  `.../baseline-normtext-nachgeschrieben/state.md:13`;
  `.../evidence/slice-183.md:13`; Commit-Message `9cd11d8`
- `befund`: Die Tabelle nennt 16 822 (a-check) gegen 726 (Ziel-Form) und daraus
  „22,9 ×"; der Quotient ist **23,17**. Die fünf übrigen Faktoren derselben
  Tabelle (0,8 / 1,1 / 1,1 / 2,3 / 3,3) rechnen sich korrekt, dieser eine nicht,
  und er ist mit `9cd11d8` in `AGENTS.md` §5 eingegangen.
- `verifizierbar`: ja — 16822 ÷ 726 = 23,17; die beiden Operanden sind mit der
  Methode des Slice exakt reproduzierbar (§Reproduktion der Messung oben).
- `klasse`: Quotient nicht nachgerechnet

### F-4 — Der neue Zeiger führt für die drei Risiko-Ausgänge auf einen Abschnitt, der drei *andere* Ausgänge trägt

- `kategorie`: HIGH
- `quelle`: `v6.5.0` · `regelwerk/modul-06-roadmap.md` §Das Beobachtungs-Register
  gegen `regelwerk/modul-05-planning-harness.md` §Offene Risiken werden bei
  Closure aufgelöst
- `pfad`: `AGENTS.md:264-265`
- `befund`: Die gekürzte Fassung sagt, „Form, Zählregel und die drei
  **Risiko-Ausgänge**" stünden in `modul-06` §Das Beobachtungs-Register; dieser
  Abschnitt führt die drei *Register*-Ausgänge **verkörpert / geplant /
  gestrichen**, die Risiko-Ausgänge **eingetreten / entfallen / weiter offen**
  stehen in `modul-05` §Offene Risiken werden bei Closure aufgelöst. Nach der
  Kürzung nennt `AGENTS.md` die drei Risiko-Ausgänge an keiner Stelle mehr,
  obwohl `make verify-risiko-ausgaenge` genau diese geschlossene Menge
  durchsetzt.
- `verifizierbar`: ja — `grep -n "eingetreten\|entfallen\|weiter offen"` gegen
  `modul-06-roadmap.md` (kein Treffer in §Das Beobachtungs-Register), gegen
  `modul-05-planning-harness.md` (Treffer, Zeilen 140–142) und gegen `AGENTS.md`
  (kein Treffer).
- `klasse`: Zeiger benennt den falschen Abschnitt; die verwiesene Zusage steht
  dort nicht

### F-5 — „die Tabelle der fünf Übergänge wortgleich" trifft auf den verwiesenen Abschnitt nicht zu

- `kategorie`: HIGH
- `quelle`: `v6.5.0` · `regelwerk/modul-05-planning-harness.md` §Trigger je
  Lifecycle-Übergang und WIP-Limit
- `pfad`: Slice-Plan Zeile 110; Commit-Message `b5a81d5`;
  `.../baseline-normtext-nachgeschrieben/evidence/slice-183.md:4`
- `befund`: Der verwiesene Abschnitt führt die fünf Übergänge als
  **Aufzählung**, nicht als Tabelle, und keine der fünf Bedingungen ist
  wortgleich — „für die nächste Welle priorisiert" gegen „priorisiert/
  eingeplant, `Verantwortlich:` gesetzt", „DoD erfüllt, Closure-Notiz
  geschrieben, Gates grün" gegen „Closure-Kriterien erfüllt, Lerneintrag
  geschrieben, jedes Risiko aus dem Slice-Plan mit Ausgang". Die nächstliegenden
  Formulierungen (zwei von fünf) stehen im Mermaid-Diagramm eines **anderen**
  Abschnitts (§Lifecycle als State Machine). Die Behauptung steht in der
  `evidence/`-Datei, die ab Merge unveränderlich ist.
- `verifizierbar`: ja — `sed -n '89,105p'` auf
  `regelwerk/modul-05-planning-harness.md` gegen `git show b5a81d5^:AGENTS.md`
  Zeilen 236–248, Zeile für Zeile.
- `klasse`: Deckungs-Behauptung stärker formuliert als gemessen
  („wortgleich" statt „inhaltlich gedeckt")

### F-6 — Innerer Widerspruch 21,3 × / 22,9 × zwischen §2, §3.1 und §8 sowie zwischen den beiden Commit-Messages

- `kategorie`: MEDIUM
- `quelle`: Selbstwiderspruch im Slice-Plan
- `pfad`: Slice-Plan Zeilen 52 und 219 („22,9 ×") gegen 123 und 146 („21,3 ×");
  Commit-Messages `b5a81d5` („21,3x") gegen `9cd11d8` („22,9x")
- `befund`: §3.1 schreibt „§2 nannte … den Faktor **21,3 ×**" — §2 nennt 22,9 ×.
  21,3 × ist der korrekt gerechnete Faktor des **Plan**-Zustands (15 464 ÷ 726 =
  21,30), wird aber an beiden Stellen als der Faktor des gemessenen Zustands
  ausgegeben; dieselbe Größe trägt damit in einem Dokument zwei Werte, und die
  zwei Commits des Slice nennen je einen anderen.
- `verifizierbar`: ja — `grep -n "21,3\|22,9"` auf der Slice-Datei und
  `git log -2 --format=%B`.
- `klasse`: Zwei Messzeitpunkte in einem Dokument nicht auseinandergehalten

### F-7 — „Sechs Stichproben … alle in `modul-05` bzw. `modul-06`" trifft für die AC-Form nicht zu

- `kategorie`: MEDIUM
- `quelle`: `v6.5.0` · `regelwerk/modul-03-spec.md`
- `pfad`: Slice-Plan Zeilen 67–71
- `befund`: Von den sechs genannten Stichproben liegt die AC-Form
  (Happy/Boundary/Negative plus Out-of-Scope) nicht in `modul-05`/`modul-06`,
  sondern in `modul-03`; a-check verankert sie zudem in `harness/conventions.md`
  §Anforderungs-Anlege-Prozess, worauf `AGENTS.md` §5 bereits verweist. Die
  Trefferzahl „sechs von sechs" bleibt richtig, die Verortungs-Aussage nicht.
- `verifizierbar`: ja — `grep -rln "Boundary" .harness/baseline/v6.5.0/regelwerk/`
  liefert `modul-03-spec.md` und `modul-12-replay-evaluierung.md`, nicht
  `modul-05`/`modul-06`.
- `klasse`: Fundort einer Stichprobe pauschal zugeordnet

### F-8 — Der Ausgang des Register-Eintrags widerspricht sich selbst: `Stand: offen` bei 3×, „keine neue Regel" neben 633 Zeichen neuer Regel

- `kategorie`: MEDIUM
- `quelle`: `v6.5.0` · `regelwerk/modul-06-roadmap.md` §Das
  Beobachtungs-Register (drei Ausgänge, geschlossene Menge)
- `pfad`: `.../baseline-normtext-nachgeschrieben/state.md:1`, `:6`, `:16-17`,
  `:19-20`
- `befund`: Zeile 1 trägt `offen (3×)` — keiner der drei Ausgänge —, während
  Zeile 16 den Zielort mit Herkunfts-Anker nennt (`AGENTS.md` §5,
  `seit slice-183`), also genau die Form des Ausgangs *verkörpert*; die
  Begründung in Zeile 19 stellt `offen` gegen `gestrichen` und lässt
  *verkörpert* unerwähnt. Zeile 6 sagt „kein Sensor, **keine neue Regel**",
  während `9cd11d8` 633 Zeichen normativen Text in `AGENTS.md` §5 einfügt, den
  Zeile 16 selbst als „Schärfung" zitiert.
- `verifizierbar`: teilweise — die Zeichenzahl ja (16 360 − 15 727 = 633,
  `git show 9cd11d8 -- AGENTS.md`); ob `offen` oder `verkörpert` der richtige
  Ausgang ist, prüft kein Sensor: `make verify-observations` misst nur die
  Deckung, nicht den Stand.
- `klasse`: Ausgang im Fließtext vergeben, im Stand-Feld nicht nachgezogen

### F-9 — Der Folge-Slice für `harness/conventions.md` wird mit einer Vorhersage gestrichen, die §1 an eine Messung geknüpft hatte

- `kategorie`: MEDIUM
- `quelle`: §1 des Slice-Plans („Ein eigener Vorgang, **wenn die Messung dort
  etwas zeigt**"); `AGENTS.md` §5 §Geltungsbereich einer Messung
- `pfad`: Slice-Plan Zeilen 255–259 (Closure-Notiz, *Folge-Slices*)
- `befund`: Die Closure begründet „keiner" mit „dieselbe Messung dort würde
  denselben Anteil finden" — in `harness/conventions.md` wurde nicht gemessen,
  und der Slice hat für §5 selbst gezeigt, dass die Vorab-Einschätzung um eine
  Größenordnung danebenlag. Damit entscheidet dieselbe Art Vermutung über das
  Ausbleiben des Folge-Vorgangs, die der Slice als Lernsignal verwirft.
- `verifizierbar`: nein — es gibt keinen Lauf, der eine unterlassene Messung
  meldet.
- `klasse`: Abgrenzung mit Prognose statt mit Messung geschlossen

### F-10 — Ergebniszeile zu §4 beantwortet die Frage nicht, die die DoD stellt

- `kategorie`: MEDIUM
- `quelle`: DoD-Punkt 2 des Slice-Plans („§4 und §6 ebenso [je Block gegen das
  vendorte Regelwerk geprüft]")
- `pfad`: Slice-Plan Zeile 103 (§3-Tabelle, Zeile `AGENTS.md` §4)
- `befund`: Die Ergebnis-Spalte für §4 lautet „mit slice-181 bereits angefasst" —
  das ist eine Aussage über die Bearbeitungs-Historie, keine über die Deckung
  gegen das Regelwerk; für §5, §6 und §1–§3 nennt dieselbe Tabelle jeweils ein
  Deckungs- oder Faktor-Argument.
- `verifizierbar`: nein — ob ein Abschnitt geprüft wurde, ist kein Match.
- `klasse`: Prüf-Behauptung mit sachfremder Begründung belegt

### F-11 — Datei-Gesamtzahlen sind Bytes, alle Abschnittszahlen daneben sind Zeichen

- `kategorie`: LOW
- `quelle`: Messmethode des Slice selbst (§2: „verbleibende **Zeichen**
  gezählt")
- `pfad`: Slice-Plan Zeilen 55–56
- `befund`: „Datei gesamt **35 384** gegen **10 885** … beim Anlegen des Plans
  34 690" sind `wc -c`-Werte; nach der Methode des Slice (`wc -m`, Kommentare
  und Tabellenzeilen entfernt) sind es 34 714 bzw. 10 701. Die Abschnittszeilen
  derselben Tabelle sind dagegen durchgängig Zeichen.
- `verifizierbar`: ja — `wc -c` gegen `wc -m` auf `b5a81d5^:AGENTS.md` und
  `templates/AGENTS.template.md`.
- `klasse`: Einheit innerhalb einer Messtabelle gewechselt

### F-12 — Der gekürzte Trichter-Absatz nennt drei Werkzeuge, aber nur zwei Ablageorte

- `kategorie`: LOW
- `quelle`: Aufbau des Absatzes selbst („die Ablageorte hier sind …")
- `pfad`: `AGENTS.md:259-263`
- `befund`: Der Absatz erklärt die Ablageorte zum repo-eigenen Anteil, den der
  Verweis nicht ersetzt, und listet dann `harness/conventions.md` und
  `docs/plan/carveouts/`; für die dritte Wahl (permanente ADR) bleibt der Ort
  offen, obwohl `docs/plan/adr/` an anderen Stellen derselben Datei verlinkt ist.
- `verifizierbar`: nein — Vollständigkeit einer Aufzählung ist kein Match.
- `klasse`: Aufzählung im Zeiger-Absatz unvollständig gegen den eigenen Anspruch

### F-13 — `AGENTS.md` §5 verlinkt eine Datei, die es nicht gibt (Bestand, im Prüfumfang)

- `kategorie`: INFO
- `quelle`: `AGENTS.md` §5, Block *Steering-Loop*
- `pfad`: `AGENTS.md:272`
- `befund`: Der Linktext lautet `docs/plan/steering-loop.md` — die Datei
  existiert nicht —, während das Ziel auf
  `docs/plan/planning/observations/README.md` zeigt; `make doc-check` bleibt
  grün, weil es das Ziel prüft, nicht den Text. Der Block gehörte zur Menge der
  „je Block geprüften" 17/18 und blieb unverändert; §1 deckt Bestand
  ausdrücklich, deshalb kein Befund gegen die Umsetzung.
- `verifizierbar`: ja — `ls docs/plan/steering-loop.md` (nicht vorhanden),
  `make doc-check` (grün).
- `klasse`: Linktext benennt ein anderes Artefakt als das Linkziel

---

## Negativbefunde

- **geprüft, ohne Befund: Substanz-Verlust im Block *Slice-Lifecycle*.** Für
  jeden entfernten Satz ist die Aussage im verwiesenen Abschnitt vorhanden
  (fünf Übergänge samt Bedingungen; der Baseline-Abschnitt trägt sogar zwei
  Bedingungen mehr, `Verantwortlich:` gesetzt und „der `git mv` landet auf dem
  Hauptzweig"). Die Wortlaut-Behauptung dazu ist F-5, der Inhalt ist gedeckt.
- **geprüft, ohne Befund: „`next/` ist keine Pflichtstation" ist repo-eigen.**
  `grep -rn "Pflichtstation"` und die Suche nach dem direkten Weg
  `open → in-progress` über `.harness/baseline/v6.5.0/` (Regelwerk **und**
  Templates) liefern **null** Treffer; die Aussage steht zu Recht in `AGENTS.md`.
- **geprüft, ohne Befund: Substanz-Verlust im Block *Diskrepanz-Trichter*.**
  Zwei Fragen, Granularität vor Temporalität, die drei Werkzeuge und die
  Abgrenzung der bootstrap-aware Gates stehen im verwiesenen Abschnitt; die
  a-check-eigene Ausprägung dazu (die drei Beispiel-Stufungen) steht unberührt
  in `docs/plan/carveouts/README.md`.
- **geprüft, ohne Befund: Substanz-Verlust im Block *Beobachtungs-Register*,
  außer den Risiko-Ausgängen.** „Warum stehend", „Eingetragen wird bei der
  Slice-Closure" und „Der Zähler wird abgeleitet" stehen im verwiesenen
  Abschnitt; der Ort, die Pfadform und „nachgeschlagen, nicht erfunden" sind in
  `AGENTS.md` geblieben. Die Ausnahme ist F-4.
- **geprüft, ohne Befund: Herkunfts-Anker.** Die entfernten Anker
  („seit slice-139", „seit slice-101", „(slice-065)") werden von keinem
  `state.md` und keiner Closure-Notiz auf `AGENTS.md` §5 als Zielort gezeigt;
  `harness-einstieg-ohne-modus-zeile` trägt `seit slice-101` auf
  `harness/conventions.md`, unberührt. Keine Anker-Paarung gebrochen.
- **geprüft, ohne Befund: Auflösung der drei neuen Zeiger.** Alle drei
  Abschnitte existieren im vendorten Stand — `modul-05` §Trigger je
  Lifecycle-Übergang und WIP-Limit, `modul-06` §Das Beobachtungs-Register,
  `modul-07` §Werkzeug-Wahl bei Diskrepanz (dort als Präfix der vollen
  Überschrift „… : Carveout, BF-Markierung oder ADR"). Das gilt für die
  **Existenz**; ob der Abschnitt trägt, was der Zeiger behauptet, ist F-4.
- **geprüft, ohne Befund: Kern-Arithmetik der Gesamtaussage.** 16 822 − 15 727 =
  **1095**; 1095 ÷ 16 822 = **6,51 %**; 16 360 − 16 822 = **−462**; 16 360 −
  15 727 = **633**; 93,5 % als Komplement zu 6,5 %; „um 1358 Zeichen gewachsen"
  = 16 822 − 15 464. Alle sechs reproduzieren exakt.
- **geprüft, ohne Befund: §1-Abgrenzung eingehalten.** Beide Commits fassen
  ausschließlich `AGENTS.md`, die Slice-Datei und den Register-Eintrag an —
  `harness/README.md` und `harness/conventions.md` sind unberührt, es entstand
  kein Sensor. Kein Punkt aus §1 wurde stillschweigend mitgenommen.
- **geprüft, ohne Befund: Hard Rules §3.1–§3.6.** Kein Host-Toolchain-Aufruf,
  keine Suppression, keine ADR berührt, keine Gate-Schwelle gesenkt; der
  Lifecycle-Wechsel `a38baae` ist der reine `make slice-mv`-Commit vor der
  Inhaltsänderung (§3.3).
- **geprüft, ohne Befund: §3.7 in den drei neuen Absätzen.** Kein Konjunktiv
  über Verworfenes, keine Chronik; die Chronik-Nachsätze („davor Tabellenform",
  „(slice-065)") sind gerade entfallen.
- **geprüft, ohne Befund: Traceability und Commit-Scope.** Beide Commits nennen
  `slice-183`; `fix(harness)`/`docs(harness)` unterliegen nicht der
  `(planning)`-Scope-Regel. `make trace-check` und `make commit-scope-check`
  laufen im grünen `gates`-Aggregat mit.
- **geprüft, ohne Befund: Gate-Behauptungen der Commit-Messages.** Eigener Lauf
  auf dem Stand `9cd11d8`: `make gates` Exit 0, `make verify` Exit 0 — beide
  Zusagen treffen zu.
- **geprüft, ohne Befund: die Regel, auf die der Register-Ausgang sich beruft.**
  Sie steht wörtlich an beiden genannten Stellen: `AGENTS.md` §1 („sie
  dupliziert deren Inhalt nicht; sonst entsteht Drift", Zeile 11) und
  `harness/README.md` §Purpose („Diese Datei dupliziert sie nicht.", Zeile 15).
  Der Ausgang trägt in der Sache; beanstandet ist nur seine Verbuchung (F-8).
- **geprüft, ohne Befund: die Kernthese des Slice.** Dass §5 zu ~93,5 %
  a-check-eigen ist, hält der Nachprüfung stand — die fünf größten Blöcke
  (4285 / 2051 / 1057 / 943 / 941) tragen je eine eigene Messung oder einen
  eigenen Anlass und haben im vendorten Regelwerk keine Entsprechung. Die
  Widerlegung der Ausgangsvermutung ist belastbar; falsch sind die Zahlen, mit
  denen sie ausgedrückt wird.

---

## Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 5 |
| MEDIUM | 5 |
| LOW | 2 |
| INFO | 1 |

**Finding-Klassen dieses Laufs:** Nachher-Wert geschätzt statt gemessen ·
Bezugsgröße aus dem Vorzustand übernommen · Quotient nicht nachgerechnet ·
Zeiger benennt den falschen Abschnitt · Deckungs-Behauptung stärker formuliert
als gemessen · Zwei Messzeitpunkte in einem Dokument nicht auseinandergehalten ·
Fundort einer Stichprobe pauschal zugeordnet · Ausgang im Stand-Feld nicht
nachgezogen · Abgrenzung mit Prognose statt mit Messung geschlossen ·
Prüf-Behauptung mit sachfremder Begründung belegt · Einheit innerhalb einer
Messtabelle gewechselt · Aufzählung im Zeiger-Absatz unvollständig ·
Linktext benennt ein anderes Artefakt als das Linkziel

**Beobachtung zur Klassen-Verteilung** (kein eigenes Finding): Sieben der
dreizehn Findings — F-1, F-2, F-3, F-5, F-6, F-7, F-11 — sind Zahl- oder
Deckungs-Aussagen, deren *Vorher*-Hälfte gemessen und deren *Nachher*- bzw.
Zuordnungs-Hälfte nicht gemessen wurde. Das ist dieselbe Asymmetrie, die der
Slice als Lernsignal formuliert, nur eine Ebene tiefer.

---

## Verdikt

**Merge-blockierend: ja.** Fünf HIGH, fünf MEDIUM.

Zur ausdrücklich gestellten Frage nach der Closure-Aussage *„netto −462, und
das ist keine Schwäche, sondern seine Funktion"*: Sie ist **tragend, nicht
Selbstrechtfertigung**. Die Zahl ist nachgerechnet richtig (16 360 − 16 822 =
−462; die 633 Zeichen des zweiten Commits sind belegt), und die daraus gezogene
Aussage — ein Abschnitt, der gelernte Regeln aufnimmt, wächst schneller als
Aufräumen ihn schrumpfen kann — ist gegen den Bestand prüfbar und nicht der
Versuch, ein verfehltes Ziel umzudeuten: Das Ziel des Slice ist laut §1 nicht
Verkleinerung, sondern *Verweis statt nachgeschriebener Begründung*. Auch die
Widerlegung der Ausgangsvermutung ist ein echtes Ergebnis und in der Sache
belegt.

**Was blockiert, ist nicht die These, sondern ihre Belege.** Der Slice erhebt
Messgenauigkeit zu seinem Gegenstand und trägt die geschärfte Prüf-Frage in eine
Rang-8-Quelle ein; in derselben Bewegung gehen dorthin ein falscher Nenner
(F-2), ein falscher Quotient (F-3) und ein Zeiger auf den falschen Abschnitt
(F-4), und in ein ab Merge unveränderliches `evidence/`-Dokument eine
Wortlaut-Behauptung, die nicht zutrifft (F-5). F-1 ist innerhalb des Dokuments
ohne Zusatzwissen erkennbar: 1271 gegen 1095.

**Übergabe:** Findings an den Implementer; die Finding-Klassen zusätzlich in die
Closure §7 und von dort in den Zähler. Dieser Report ist Lauf-Beleg und ersetzt
keine Verifikation — DoD-/Spec-Konformität prüft der Verifier separat (Modul 11).
