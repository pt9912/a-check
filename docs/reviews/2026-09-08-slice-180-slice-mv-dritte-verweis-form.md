# Review-Report: slice-180 — 2026-09-08

**Review-Art:** Code + Plan — geprüft gegen den Slice-Plan und gegen das
Werkzeug selbst (Modul 10 §Drei Review-Arten). Kern-Behauptung ist, dass
`make slice-mv` die dritte Verweis-Form nachzieht; sie ist an ausführbaren
Läufen gegen den **realen historischen Bestand** gemessen, nicht am
Plan-Text.

**Gegenstand:** Commit `577d698` („feat(harness): slice-180 -- slice-mv lernt
die dritte Verweis-Form"), dazu der vorangehende Übergangs-Commit `9124b71`
als Kontext.

**Skill:** `.harness/skills/reviewer.md` @ Stand `3fae6d3` · <!-- d-check:ignore -->
**Modell:** claude-opus-5 · **Datum:** 2026-09-08

**Eingangs-Kontext:**

- `docs/plan/planning/in-progress/slice-180-slice-mv-dritte-verweis-form.md`
- `tools/slice-mv.sh` @ `577d698` und @ `577d698^`
- `AGENTS.md` §3.1/§3.2/§3.7, §5, §6
- Beobachtungs-Register: `BEO-PLAN/verweis-auf-wandernden-slice`,
  `BEO-PLAN/slice-mv-fasst-einfrierendes-artefakt-an`,
  `BEO-GATE/muster-trifft-nur-die-haeufige-schreibweise`
- der reale flache Welle-Plan `welle-15-regelwerk-v650-migration.md` in den
  Ständen `61b7661`, `97472ab`, `179c44b` (der Gegenstand, für den die dritte
  Form eingeführt wurde)
- `v6.5.0` · `regelwerk/modul-06-roadmap.md` §Das Beobachtungs-Register;
  `v6.5.0` · `regelwerk/modul-05-planning-harness.md` §Offene Risiken werden
  bei Closure aufgelöst

**Nicht Gegenstand dieses Laufs:** `make gates` / `make verify` sind nicht
erneut ausgeführt worden — Verifikation ist eine andere Rolle (Modul 11).
Alle Messungen unten liefen auf **Kopien** außerhalb des Repos; am
Review-Gegenstand ist nichts geändert.

---

## Findings

### F-1 — Der Auswahlschritt kennt die dritte Form nicht; der Zielfall bleibt unrepariert

- `kategorie`: HIGH
- `quelle`: `AGENTS.md` §5 (Slice-Lifecycle: „`make slice-mv`, das den `git mv`
  samt der Verweise **auf** die Datei fährt"); DoD-Punkt 1 des Slice-Plans
- `pfad`: `tools/slice-mv.sh:164–165`
- `befund`: Die zwei neuen `sed`-Regeln stehen in `rewrite_file()`, aber die
  **Kandidatenliste** wird eine Zeile weiter unten weiterhin nur über die
  ersten beiden Formen gebildet (`grep -rl -e "\.\./$from/$base" -e
  "$PLANNING/$from/$base"`). Eine Datei, die den Slice ausschließlich in der
  dritten Form nennt — genau der Fall, für den der Slice gebaut ist —, wird nie
  an `rewrite_file` übergeben und bleibt unverändert.
- `verifizierbar`: ja — reproduziert gegen den realen Gegenstand (Handgriff
  unten).
- `klasse`: Ersetzungsregel ohne passenden Auswahlschritt

**Adversarische Verifikation.** Der flache Welle-Plan aus `welle-15` ist die
Datei, an der die Lücke zehnmal auftrat. In allen drei Ständen, in denen ein
`slice-mv` über ihn lief, zählt er **null** Vorkommen von Form 1 und Form 2:

| Stand | Form 1 (`../<dir>/slice-`) | Form 2 (`docs/plan/planning/<dir>/slice-`) | Form 3 (`](<dir>/slice-`) |
|---|---|---|---|
| `61b7661` | 0 | 0 | 1 |
| `97472ab` | 0 | 0 | 6 |
| `179c44b` | 0 | 0 | 6 |

Damit kann die Kandidaten-`grep` ihn nicht finden. Nachgestellt in einem
Wegwerf-Repo mit der **unveränderten** Datei aus `61b7661` und `slice-169` in
`open/`:

```
$ bash tools/slice-mv.sh slice-169 in-progress
slice-mv ok: slice-169-korpus-seitige-kalibrierung.md  open/ -> in-progress/,
  0 Datei(en) mit Verweisen nachgezogen (Selbsttest gefeuert).
$ grep -n 'slice-169' docs/plan/planning/welle-15-regelwerk-v650-migration.md
80:- Blockiert: [slice-169](open/slice-169-korpus-seitige-kalibrierung.md) nicht
```

Exit 0, „0 Datei(en) … nachgezogen", der Verweis zeigt weiter auf `open/`, das
Ziel liegt in `in-progress/` — derselbe `target-missing`, den DoD-Punkt 1
ausschließen soll („**ohne** Handgriff danach").

**Gegenprobe in die andere Richtung:** Dieselbe Datei, ergänzt um *einen*
Verweis in Form 1, wird gefunden — und dann werden alle drei Formen korrekt
umgeschrieben. Die neuen Regeln funktionieren; erreicht werden sie nur über
eine der alten.

### F-2 — Der Usage-Block sagt eine Deckung zu, die gemessen nicht besteht

- `kategorie`: HIGH
- `quelle`: `AGENTS.md` §4 (halluzinierte Zusagen sind die häufigste Form von
  Harness-Lüge); `AC-QA-02` (ehrliche Grenze wird ausgewiesen, nicht
  überschrieben)
- `pfad`: `tools/slice-mv.sh:43–47` (Usage), dazu die Commit-Message von
  `577d698` („End-to-End … alle drei umgeschrieben")
- `befund`: Der Usage-Block behauptet „Alle drei im Bestand vorkommenden Formen
  werden getroffen"; gemessen gilt das nur für Dateien, die zusätzlich Form 1
  oder Form 2 tragen (F-1). Der Kopf-Kommentar führt zugleich einen
  ausdrücklichen Abschnitt „NICHT BEHANDELT (ehrliche Grenzen, `AC-QA-02`) —
  zwei Stueck", in dem diese Bedingung nicht steht.
- `verifizierbar`: ja — derselbe Handgriff wie bei F-1; die Zusage ist im
  Quelltext, das Gegenbeispiel im `git`-Bestand.
- `klasse`: Zusage ohne Deckung im selben Artefakt

### F-3 — Der Geltungsbereich der End-to-End-Probe deckt den Gegenstand nicht und ist nicht als Grenze benannt

- `kategorie`: MEDIUM
- `quelle`: `AGENTS.md` §5 §Geltungsbereich einer Messung (`seit slice-179`) —
  wer eine Messung als Beleg schreibt, nennt ihren Geltungsbereich und sagt, ob
  er den Gegenstand deckt
- `pfad`: `docs/plan/planning/in-progress/slice-180-slice-mv-dritte-verweis-form.md`
  §3, Absatz „End-to-End auf einem Klon"
- `befund`: Die Probe legte eine flache Datei „mit **allen drei** Formen" an;
  der reale Gegenstand — der flache Welle-Plan aus `welle-15` — trägt
  ausschließlich die dritte (Tabelle unter F-1). Genau die zusätzlichen Formen
  der Probe bewirken die Auswahl und machen die Lücke unsichtbar; der Bericht
  nennt diesen Unterschied nicht als Grenze der Messung.
- `verifizierbar`: ja — beide Proben nachgestellt, die mit allen drei Formen
  läuft grün, die mit nur der dritten meldet „0 Datei(en)".
- `klasse`: Probe konstruiert, statt den realen Gegenstand nachzustellen

### F-4 — Der Selbsttest prüft die Ersetzung, nicht den Weg dorthin

- `kategorie`: MEDIUM
- `quelle`: DoD-Punkt 2 des Slice-Plans („Der Selbsttest … deckt alle **drei**
  Formen ab, je Richtung eine Mutations-Probe")
- `pfad`: `tools/slice-mv.sh:68–120` (`self_test`)
- `befund`: `self_test` ruft ausschließlich `rewrite_file` auf einer
  vorbereiteten Datei auf; der Auswahlschritt in Zeile 164 liegt außerhalb
  seines Geltungsbereichs. Eine Mutation, die die Kandidaten-`grep` beschädigt
  oder — wie im Ist-Zustand — unvollständig lässt, lässt den Selbsttest grün.
- `verifizierbar`: ja — die vier in §3 behaupteten Mutationen sind
  nachgestellt und verhalten sich wie beschrieben (Negativbefunde unten); eine
  fünfte, die den Auswahlschritt trifft, bleibt unentdeckt.
- `klasse`: Testgrenze fällt nicht mit der Zusagegrenze zusammen

### F-5 — Die Prognose in §9 zum Register-Eintrag wird von der eigenen DoD widerlegt

- `kategorie`: MEDIUM
- `quelle`: `BEO-PLAN/slice-mv-fasst-einfrierendes-artefakt-an` (`observation.md`
  und beide `evidence/`-Dateien); DoD-Punkt 3 des Slice-Plans
- `pfad`: `docs/plan/planning/in-progress/slice-180-slice-mv-dritte-verweis-form.md`
  §9, zweiter Aufzählungspunkt
- `befund`: §9 stellt fest, der Eintrag „erreicht mit diesem Slice die Schwelle
  **nicht**". Die DoD verlangt einen unabhängigen Review-Report unter
  `docs/reviews/`, dessen `pfad`-Felder den geprüften Stand `in-progress/`
  festhalten; der Closure-`slice-mv` schreibt genau diese Felder um — dieselbe
  Ausfallart und derselbe Dateityp wie in `evidence/slice-169.md` und
  `evidence/slice-182.md`, in einem neuen Vorgang.
- `verifizierbar`: ja — nachgestellt mit einer Datei, die diesen Report
  nachbildet (Handgriff unten).
- `klasse`: Zähler-Prognose ohne den eigenen Folgeschritt gerechnet

**Adversarische Verifikation.**

```
$ cat docs/reviews/2026-09-08-slice-180-x.md
- `pfad`: `docs/plan/planning/in-progress/slice-180-slice-mv-dritte-verweis-form.md` §3
$ bash tools/slice-mv.sh slice-180 done
slice-mv ok: … in-progress/ -> done/, 1 Datei(en) mit Verweisen nachgezogen …
$ cat docs/reviews/2026-09-08-slice-180-x.md
- `pfad`: `docs/plan/planning/done/slice-180-slice-mv-dritte-verweis-form.md` §3
```

Der Report behauptet danach, gegen einen `done/`-Stand geprüft zu haben, den es
beim Lauf nicht gab — wortgleich die Beschreibung in `observation.md`.

### F-6 — Die Begründung „derselbe Vorgang" steht gegen den eigenen Bestand

- `kategorie`: LOW
- `quelle`: `v6.5.0` · `regelwerk/modul-06-roadmap.md` §Das Beobachtungs-Register
  („Zwei Funde **im selben** Vorgang sind … *eine* Gelegenheit")
- `pfad`: `docs/plan/planning/in-progress/slice-180-slice-mv-dritte-verweis-form.md`
  §9, dritter Aufzählungspunkt
- `befund`: §9 begründet, der Zähler von
  `BEO-GATE/muster-trifft-nur-die-haeufige-schreibweise` steige nicht, weil „der
  Fund … derselbe Vorgang" sei, den ein anderer Eintrag schon zählt. Die zitierte
  Regel begrenzt die Zählung **pro Eintrag**, nicht über Einträge hinweg; der
  Bestand führt mit slice-175 einen Vorgang, der in **beiden** Registern eine
  eigene `evidence/`-Datei hat.
- `verifizierbar`: ja — beide `evidence/slice-175.md` liegen im Repo.
- `klasse`: Zählregel über ihre Reichweite hinaus angewandt

### F-7 — Neuer Kommentar im Konjunktiv über die verworfene Alternative

- `kategorie`: LOW
- `quelle`: `AGENTS.md` §3.7 („Falsch: Konjunktiv über die verworfene
  Alternative … Richtig: Indikativ über den Zustand")
- `pfad`: `tools/slice-mv.sh:112–114`
- `befund`: Der neu eingefügte Kommentar lautet „Ohne sie waere eine Ersetzung,
  die jedes `open/` trifft, von einer kontext-gebundenen nicht zu
  unterscheiden" — strukturgleich zum „Falsch"-Beispiel in §3.7. Er setzt eine
  gleichartige Formulierung aus dem Bestand fort (Zeilen 93–94, `seit
  slice-118`), ist aber selbst neu.
- `verifizierbar`: nein — §3.7 ist ausdrücklich inferentiell und hat keinen
  Sensor.
- `klasse`: Kommentar beschreibt die Abwägung statt den Zustand

**Kein HIGH, obwohl Hard Rule:** Die HIGH-Definition des Skills nennt
ausdrücklich §3.1–§3.6; §3.7 ist dort nicht erfasst. Die Einordnung als LOW ist
darum eine Anwendung des Skills, keine Milderung.

### F-8 — Unmaskierte Metazeichen im Dateinamen ändern still den Zielnamen

- `kategorie`: LOW
- `quelle`: eigene Messung an `rewrite_file`
- `pfad`: `tools/slice-mv.sh:60–65`
- `befund`: `$2` geht unmaskiert in alle vier `sed`-Ausdrücke; der Punkt in
  `.md` ist damit ein Metazeichen. Gemessen: `[x](open/slice-999-xYmd)` wird zu
  `[x](done/slice-999-x.md)` — der Dateiname ändert sich still. Die Klasse
  besteht seit `slice-118` in den ersten beiden Regeln; die zwei neuen Regeln
  erweitern ihre Fläche um Link-Ziel und Inline-Code.
- `verifizierbar`: ja — 18-Fall-Probe gegen `rewrite_file`, Fall 15.
- `klasse`: Eingabe ungeschützt im Muster

### F-9 — Risiko-Ausgang zeigt auf einen Eintrag, den der Commit nicht anfasst

- `kategorie`: INFO
- `quelle`: `v6.5.0` · `regelwerk/modul-05-planning-harness.md` §Offene Risiken
  werden bei Closure aufgelöst („der dritte Ausgang hängt das Risiko an den
  Zähler")
- `pfad`: `docs/plan/planning/in-progress/slice-180-slice-mv-dritte-verweis-form.md`
  §7, erster Aufzählungspunkt
- `befund`: Der Ausgang lautet „*weiter offen* → Beobachtungs-Register,
  `BEO-GATE/muster-trifft-nur-die-haeufige-schreibweise`", während §9 für
  denselben Eintrag festhält, dass sein Zähler nicht steigt; der Eintrag ist in
  `577d698` nicht angefasst. `v6.5.0` · `regelwerk/modul-06-roadmap.md` lässt
  „benannt, nicht gezählt" ausdrücklich zu — ob dieser Weg gegangen wird,
  entscheidet sich in der Closure, und §8 ist noch leer. Hinweis an die
  Verifier-Rolle, kein Befund am Commit.
- `verifizierbar`: ja — `make verify-risiko-ausgaenge` prüft die Form des
  Ausgangs, nicht seinen Vollzug.
- `klasse`: Ausgang deklariert, Vollzug offen

## Negativbefunde

- geprüft, ohne Befund: **Über-Treffer der zwei neuen Regeln** — 18 konstruierte
  Fälle gegen `rewrite_file`. Unverändert bleiben: anderer Slice
  (`](open/slice-998-y.md)`), anderes Verzeichnis (`](done/…)`, `](next/…)`),
  fremdes Präfix (`](docs/foo/open/…)`), Großschreibung (`](Open/…)`), nackter
  Pfad allein auf der Zeile und nackter Pfad im Satz. Korrekt getroffen werden
  Anker (`…md#abschnitt`) und Link-Titel (`…md "Titel"`).
- geprüft, ohne Befund: **Idempotenz und Rückweg** — drei aufeinanderfolgende
  Läufe (`open`→`in-progress`→`done`→`in-progress`) auf einer Datei, die den
  bereits richtigen Verweis in beiden neuen Formen zusätzlich trägt: keine
  Verdopplung, kein Verlust, jeder Lauf endet im erwarteten Zustand.
- geprüft, ohne Befund: **die vier Mutations-Proben aus §3** — alle vier
  nachgestellt und wie behauptet: unverändert grün; Link-Regel entfernt →
  „Form 3 als Link nicht ersetzt"; Inline-Code-Regel entfernt → „Form 3 in
  Inline-Code nicht ersetzt"; Verallgemeinerung `s|$3/$2|$4/$2|g` → „nackter
  Pfad mitgeaendert". Die Tabelle in §3 ist reproduzierbar.
- geprüft, ohne Befund: **weitere reale Schreibweisen der dritten Form** —
  `](./<dir>/…)`, Winkelklammer-Ziel `](<…>)`, `<a href="…">` und
  Referenz-Links `[x]: <dir>/…` kommen im eigenen Bestand **null** Mal vor
  (`grep` über alle `*.md` außerhalb von `.harness/baseline/`). Die
  Nicht-Abdeckung dieser Formen ist damit keine übersehene reale Form.
- geprüft, ohne Befund: **neue Fläche auf einfrierenden Artefakten** — unter
  `docs/reviews/` gibt es heute **null** Vorkommen der zwei neuen Anker
  (präfixloses Link-Ziel bzw. Inline-Code mit Slice-Basename); die sieben
  Reports mit Pfad-Nennungen tragen ausschließlich Form 2 und waren schon vor
  `577d698` betroffen. Die zwei neuen Regeln vergrößern die Fläche heute nicht.
- geprüft, ohne Befund: **die Unterscheidung in §9 zum Eintrag
  `slice-mv-fasst-einfrierendes-artefakt-an` trägt.** `9124b71` fasste fünf
  Dateien an, keine unter `docs/reviews/`; der Nachzug in der
  Welle-Ergebnisnotiz betraf zwei Markdown-Links (Form 1), die sonst gebrochen
  wären. `observation.md` benennt als Ausfallart ausdrücklich das
  Inline-Code-`pfad`-Feld eines Review-Reports, und `evidence/slice-169.md`
  führt denselben Ausnahme-Präzedenzfall („dort steht ein echter Markdown-Link,
  der sonst bräche"). Der Übergang ist **kein** dritter Beleg. (Der dritte
  entsteht erst durch den Closure-Schritt — F-5.)
- geprüft, ohne Befund: **§1-Abgrenzung eingehalten.** Beide Ausschlüsse halten
  — kein Sensor auf die Verweis-Form angelegt, keine Änderung an der Behandlung
  der Verweise *in* der wandernden Datei. Der Diff berührt außer dem Slice-Plan
  nur `tools/slice-mv.sh` in den drei in §3 angekündigten Stellen.
- geprüft, ohne Befund: **`AGENTS.md` §3.1/§3.2** — keine Host-Toolchain im
  Diff (`sed`/`grep`/`git`), keine Suppression-Direktive.
- geprüft, ohne Befund: **Zähler-Konsistenz der drei zitierten
  Register-Einträge** — `verweis-auf-wandernden-slice` 7 `evidence/`-Dateien
  gegen „7×", `slice-mv-fasst-einfrierendes-artefakt-an` 2 gegen „2×",
  `muster-trifft-nur-die-haeufige-schreibweise` 1 gegen „1×". Alle drei decken
  sich mit ihrer `state.md`.
- geprüft, ohne Befund: **Trigger-Phrase der Review-DoD-Zeile** — die Zeile
  trägt „Unabhängiger Review"; dieselbe Großschreibung führen bereits
  abgeschlossene Slices im `done/`-Bestand, für die `make doc-reviews` grün
  läuft.
- geprüft, ohne Befund: **§2 als Analyse** — die Begründung, warum die Form vor
  `welle-15` nicht vorkam, ist am Bestand nachvollziehbar; das zitierte
  Selbsttest-Fragment steht wortgleich im Stand `577d698^`.

## Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 2 |
| MEDIUM | 3 |
| LOW | 3 |
| INFO | 1 |

**Finding-Klassen dieses Laufs:** Ersetzungsregel ohne passenden
Auswahlschritt · Zusage ohne Deckung im selben Artefakt · Probe konstruiert
statt realen Gegenstand nachgestellt · Testgrenze fällt nicht mit der
Zusagegrenze zusammen · Zähler-Prognose ohne den eigenen Folgeschritt
gerechnet · Zählregel über ihre Reichweite hinaus angewandt · Kommentar
beschreibt die Abwägung statt den Zustand · Eingabe ungeschützt im Muster ·
Ausgang deklariert, Vollzug offen

**Wiederkehrende Klasse.** F-1, F-3 und F-4 sind dieselbe Klasse in drei
Artefakten: ein Muster ist an den Vorkommen kalibriert, die beim Bauen sichtbar
waren, und die Probe, die es belegen soll, teilt die Verengung. Das ist die
Beschreibung von `BEO-GATE/muster-trifft-nur-die-haeufige-schreibweise` — und
in diesem Lauf trifft sie den Slice, der genau diese Klasse behandelt. Die
Gegenprobe, die `state.md` dieses Eintrags als das billige Mittel nennt („die
Gegenprobe am Fund"), ist hier verfügbar gewesen: Der reale Gegenstand liegt in
`git` und war ohne Klon nachstellbar.

## Verdikt

**Merge-blockierend: ja.** Zwei HIGH stehen offen. Der Slice liefert die
Ersetzungsregeln korrekt — sie sind eng gefasst, treffen kein Gegenbeispiel und
sind idempotent —, aber sie werden auf dem realen Gegenstand nie erreicht: Der
Auswahlschritt sucht weiterhin nur nach den zwei alten Formen (F-1), und
Usage-Block wie Commit-Message sagen eine Deckung zu, die dagegen nicht besteht
(F-2). Der flache Welle-Plan aus `welle-15` — die Datei, an der die Lücke
zehnmal auftrat — trägt in **allen** geprüften Ständen null Vorkommen der zwei
alten Formen; DoD-Punkt 1 („ohne Handgriff danach") ist damit nicht erfüllt.

Die drei MEDIUM erklären, warum das im Lauf nicht auffiel (F-3/F-4), und
benennen eine Prognose, die der eigene Abschluss widerlegt (F-5). Die LOW und
das INFO sind ohne Eile.

**Übergabe:** an die Implementer-Rolle (Modul 8). Dieser Report ist Lauf-Beleg,
keine Verifikation — DoD-/Spec-Konformität prüft der Verifier separat
(Modul 11); `make gates` und `make verify` sind in diesem Lauf **nicht**
ausgeführt worden.
