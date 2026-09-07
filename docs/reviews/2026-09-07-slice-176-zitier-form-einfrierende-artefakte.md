# Review — slice-176: Zitier-Form für einfrierende Artefakte

**Review-Art:** unabhängiger Lauf (eigenes Kontextfenster, kein `fork`; die Umsetzung
stammt aus einem anderen Kontext — `AGENTS.md` §6, `v6.5.0` · `regelwerk/modul-08-agentenrollen.md`
§Rollen-Regeln)
**Skill:** `.harness/skills/reviewer.md` @ `3fae6d3` · **Modell:** `claude-opus-5[1m]`
**Datum:** 2026-09-07
**Gegenstand:** Commit-Range `HEAD~2..HEAD` (`47fb99d` Lifecycle-`git mv`, `3fae6d3` Umsetzung)

> **Zitier-Form** *(Norm, bleibt stehen).* Dieser Report friert ein; was er zitiert,
> bewegt sich weiter. Deshalb **Kennung, nicht Adresse** — `slice-NNN` statt seines
> Lifecycle-Pfads, `make <target>` statt eines Links auf die Sensor-Datei, eine
> Baseline-Stelle als `v<X.Y.Z>` · `regelwerk/<datei>.md` §<Abschnitt>. Das `pfad`-Feld
> auf den **geprüften Gegenstand** ist davon nicht betroffen — es hält den Stand des
> Laufs fest und darf das.

**Eingangs-Kontext** (die Verträge, gegen die geprüft wurde):

- Slice-Plan slice-176 (Stand `3fae6d3`), Diff der Range `HEAD~2..HEAD`
- `AGENTS.md` §3 (Hard Rules), §4 (Gate-Tabelle), §5 (Dokumentations-Regeln)
- `harness/conventions.md` §Baseline, §Modus-Deklaration pro Sub-Area
- `.harness/skills/reviewer.md` (Klassifikation, Output-Schema, Zitier-Form)
- `v6.5.0` · `regelwerk/modul-05-planning-harness.md` §Offene Risiken werden bei Closure aufgelöst
- `v6.5.0` · `regelwerk/modul-06-roadmap.md` §Das Beobachtungs-Register
- `v6.5.0` · `templates/docs/reviews/review-report.template.md` §Zitier-Form sowie
  `templates/docs/plan/planning/welle-results.template.md`, `archiv-stub-slice.template.md`,
  `archiv-stub-welle.template.md`
- frühere Findings am selben Bereich: Reports zu slice-174 und slice-175
- `AC-*`: keine berührt (der Slice-Kopf führt `— (keine)`, geprüft und zutreffend)

---

## HIGH

### F-1 — Ein eingefrorener Evidence-Beleg hat durch die Umstellung seine Aussage verloren und behauptet jetzt etwas Falsches

- **kategorie:** HIGH
- **quelle:** `v6.5.0` · `regelwerk/modul-06-roadmap.md` §Das Beobachtungs-Register
  (`evidence/<vorgangs-id>.md` — *„unveränderlich ab Merge"*; dieselbe Stelle wird im
  Kommentar zu `versions.patterns[1].exempt-paths` in `.d-check.yml` wörtlich als
  Begründung geführt) · slice-176 §2.2 (*„Form-Änderung, keine Inhalts-Änderung"*) ·
  Reviewer-Skill §Klassifikation, *nachweislich falsche Tatsachenbehauptung*
- **pfad:** `docs/plan/planning/observations/BEO-GATE/muster-trifft-nur-die-haeufige-schreibweise/evidence/slice-175.md`:7
- **befund:** Der Beleg dokumentiert den Fund, dass das `versions`-Muster die relative
  Schreibweise `../baseline/<tag>/` nicht traf, und **zitierte das Belegstück wörtlich**:
  *„`.harness/skills/reviewer.md` Zeile 7 verweist als `` `[…](../baseline/v6.2.0/regelwerk/modul-10-review-harness.md)` ``"*.
  Die Umstellung hat genau dieses Belegstück ersetzt; der Satz lautet jetzt *„… verweist
  als ``v6.2.0` · `regelwerk/modul-10-review-harness.md``"*. Damit ist (a) die Zeichenfolge
  verschwunden, um die der ganze Eintrag geht — `../baseline/` ohne `.harness/`-Präfix kommt
  im Satz nicht mehr vor, und der Folgesatz *„das `sed` … suchte dieselbe zu enge
  Zeichenfolge"* hat keinen Bezugspunkt mehr —, und (b) der Satz behauptet über ein
  Repo-Artefakt etwas Unwahres: `.harness/skills/reviewer.md` Zeile 7 ist ein
  Markdown-Link und war nie eine Kennung. Der Vorgang slice-175 ist geschlossen (liegt in
  `done/`), der Beleg damit eingefroren.
- **Gegen-Messung (adversarisch, HIGH-Pflicht):** Die Zeile im Stand vor dem Commit
  wiederhergestellt und `make doc-check` gefahren → **Exit 0, 475 Dateien, 0 Befunde** —
  obwohl der vorige vendorte Stand zu diesem Zeitpunkt bereits gelöscht war. Das Vorkommen
  stand in einem Inline-Code-Span, war also nie ein auflösbarer Link, konnte beim Löschen
  nicht brechen und musste für den Beweis in §2.3 nicht angefasst werden. Danach
  zurückgesetzt, `git status` leer. Folgerichtig zählt die Messung in §2.1 („16 Links")
  ein zitiertes Belegstück als Link mit; die tragende Zahl des Beweises in §2.3 hängt nicht
  daran, die Rechtfertigung in §2.2 („Form-Änderung") für **diese** Stelle schon.
- **verifizierbar:** ja — `make doc-check` mit wiederhergestellter Zeile (Exit 0) belegt,
  dass die Änderung nicht link-getrieben war; die Falschaussage selbst ist gegen
  `.harness/skills/reviewer.md`:7 prüfbar, aber kein Gate misst sie (§3.7: inferentiell).

---

## MEDIUM

### F-2 — `harness/conventions.md` §Baseline behauptet die Klemme weiter im Präsens

- **kategorie:** MEDIUM
- **quelle:** slice-176 §2.3 (*„ist damit bezahlt und nicht mehr fällig"*) ·
  `AGENTS.md` §5 (neuer Zitier-Form-Absatz) · `AGENTS.md` §3.7 *(Zustandsfelder nennen den
  Zustand, nicht die Chronik)* · Source Precedence
- **pfad:** `harness/conventions.md`:41–45
- **befund:** Der Slice erklärt die dort beschriebene Klemme für erledigt und verankert das
  in `AGENTS.md` §5, lässt aber die Quelle unverändert. Dort steht weiterhin im Präsens:
  *„Eine Kollision bleibt und ist keine Ausnahme, sondern eine Klemme: verschwindet ein
  vendorter Stand, bricht ihr Link … Wer löscht, editiert Eingefrorenes; das ist der Preis
  des Löschens."* Gemessen ist das Gegenteil: `git grep -nE '\]\([^)]*baseline/v[0-9]'` über
  `harness/conventions/done/` liefert **kein** Vorkommen mehr. `harness/conventions.md` ist
  ein lebendes Dokument und war ausdrücklich nicht als „nicht anfassbar" ausgenommen — die
  Aussage ist damit stehengelassen, nicht geschützt. Zwei Repo-Stellen sagen jetzt
  Gegensätzliches über denselben Sachverhalt; ein Agent, der die Konventionen liest, bekommt
  die alte Antwort.
- **verifizierbar:** ja — der `git grep` oben; kein Gate deckt den Widerspruch (`versions`
  vergleicht Versionsangaben, nicht Aussagen).

### F-3 — Der als einzige Restlücke benannte Spalt existiert in der genannten Richtung nicht; der tatsächliche bleibt ungenannt

- **kategorie:** MEDIUM
- **quelle:** slice-176 §3 (*„Nicht angefasst: `tools/archive-wave/`. Es erzeugt Archiv-Stubs,
  und deren Zitier-Form ist damit Code-Sache"*), §8 und `state.md` des Eintrags
  `BEO-HARNESS/zwei-baseline-staende-nach-migrationsende` (*„erzeugt Archiv-Stubs im Code und
  kennt sie nicht"*) · `AGENTS.md` §5 (*„für Archiv-Stubs erzeugt `tools/archive-wave/` den Text"*) ·
  Reviewer-Skill §Klassifikation, *unbelegte Tatsachenbehauptung*
- **pfad:** `tools/archive-wave/stub.go`, `tools/archive-wave/archive.go`:118–140;
  `docs/plan/planning/observations/BEO-HARNESS/zwei-baseline-staende-nach-migrationsende/state.md`
- **befund:** `grep -rn "baseline" tools/archive-wave/*.go` liefert **kein** Vorkommen. Die
  vier Stub-Generatoren (`SliceStub`, `SliceStubStandalone`, `ReviewStub`, `WelleStub`)
  erzeugen ausschließlich Titel, Archiv-Zeiger, `Welle:`/`Archiviert:`-Felder und
  `Hervorgegangen:`-Kennungen — eine Baseline-Stelle zitiert der erzeugte Stub überhaupt
  nicht. Der behauptete Restspalt („beim nächsten Sprung schlägt es zu") kann auf diesem Weg
  nicht eintreten. Was der Generator dagegen **erzeugt**, ist die andere Hälfte derselben
  Norm: einen Markdown-Link im `**Welle:**`-Feld und in der Titelzeile, samt der eigens dafür
  gebauten Nachzieh-Mechanik `RewriteFieldForMove` (belegt in `archive_test.go`:66–75, das auf
  `(welle-77-x.md)` im Stub besteht). Genau das verbietet die Zitier-Form mit
  *„`slice-NNN` statt seines Lifecycle-Pfads"*. Der Slice benennt also einen Spalt, den es
  nicht gibt, und übergeht den, den es gibt — die Aussage ist damit nicht gemessen, sondern
  angenommen, und sie ist der einzige Restposten, den `state.md` nach der Umstufung auf
  *verkörpert* noch führt.
- **verifizierbar:** ja — `grep -rn "baseline" tools/archive-wave/*.go` (leer) und
  `make archive-wave-test` gegen `TestFixture_OrtsfesteVerweiseTiefenwechsel`.

### F-4 — Ein Wiederauftreten im eigenen Vorgang ist in der Commit-Message benannt, im Register aber nicht belegt; §9 führt den Eintrag nicht

- **kategorie:** MEDIUM
- **quelle:** `AGENTS.md` §5 (*„Eingetragen wird bei der Slice-Closure … oder eine weitere
  Datei in ein vorhandenes `evidence/`"*) · `v6.5.0` · `regelwerk/modul-05-planning-harness.md`
  §Zwei Schritte vor der Modus-Begründung (Sichtungs-Schritt) ·
  `v6.5.0` · `regelwerk/modul-06-roadmap.md` §Das Beobachtungs-Register
- **pfad:** `docs/plan/planning/observations/BEO-PLAN/verweis-auf-wandernden-slice/evidence/`
  (keine Datei `slice-176.md`); slice-176 §8 und §9
- **befund:** Die Message von `47fb99d` hält fest: *„den Zeiger in der flach liegenden
  Welle-Datei zog slice-mv wieder nicht nach (bekannte Form-Lücke,
  BEO-PLAN/verweis-auf-wandernden-slice)"* — der Diff bestätigt es, `welle-15-…md` wurde von
  Hand nachgezogen. Der Register-Eintrag selbst hatte diesen Vorgang **vorhergesagt**: sein
  Beleg `evidence/slice-175.md` schreibt *„bei den verbleibenden Etappen (C, D, E) kommen je
  zwei dazu"*; slice-176 ist Etappe C. Für die beiden unmittelbar vorangehenden Vorgänge
  liegen Belege (`slice-174.md`, `slice-175.md`), für diesen keiner. §8 erklärt pauschal *„kein
  neuer Eintrag, kein neuer Beleg"*, §9 listet drei einschlägige Einträge und führt diesen
  vierten nicht auf, obwohl er getroffen wurde.
- **verifizierbar:** ja — `make verify-observations` prüft nur die Gegenrichtung (zitierter
  Pfad ⇒ Verzeichnis) und bleibt darum grün; der Befund ist an
  `git show 47fb99d` gegen `ls …/verweis-auf-wandernden-slice/evidence/` ablesbar.

### F-5 — Zwei der drei Risiko-Ausgänge liegen außerhalb der geschlossenen Dreier-Menge und passieren den Sensor über ein negiertes Token

- **kategorie:** MEDIUM
- **quelle:** `v6.5.0` · `regelwerk/modul-05-planning-harness.md` §Offene Risiken werden bei
  Closure aufgelöst (*„eingetreten → Carveout oder Folge-Slice mit ID"*) · `AGENTS.md` §4
  (`make verify-risiko-ausgaenge`: *„genau einen Ausgang aus der geschlossenen Dreier-Menge"*) ·
  slice-176 DoD-Punkt *„Jedes Risiko trägt einen Ausgang"*
- **pfad:** slice-176 §7, Risiken 2 und 3 (Stand `3fae6d3`, Zeilen 168–176)
- **befund:** Beide tragen *„eingetreten"*, aber weder einen Carveout noch eine
  Folge-Slice-ID — sie sagen ausdrücklich *„Kein Folge-Slice"* bzw. *„kein Folge-Slice"* und
  wählen damit einen vierten Ausgang („eingetreten, im Slice aufgefangen"), den die Menge
  nicht kennt. Der Sensor lässt sie durch, weil sein Muster auf die Zeichenfolge
  `Folge-Slice` prüft und diese in der **Verneinung** vorkommt. Mutations-Probe: in beiden
  Zeilen `Folge-Slice` durch `Nachfolge-Vorgang` ersetzt → `make verify-risiko-ausgaenge`
  meldet genau diese zwei als *„Risiko ohne Ausgang aus der geschlossenen Dreier-Menge"*,
  Exit 1. Danach zurückgesetzt, Sensor wieder Exit 0, `git status` leer. Inhaltlich zählt das
  doppelt: Risiko 2 ist das Immutabilitäts-Risiko, und F-1 zeigt, dass es eben nicht
  vollständig „im Slice aufgefangen" wurde.
- **verifizierbar:** ja — die Mutations-Probe oben; der Sensor selbst kann die Verneinung
  nicht sehen.

---

## LOW

### F-6 — Kaputte Inline-Code-Auszeichnung an der umgestellten Stelle

- **kategorie:** LOW
- **quelle:** `v6.5.0` · `templates/docs/reviews/review-report.template.md` §Zitier-Form
  (Form: zwei getrennte Code-Spannen)
- **pfad:** `docs/plan/planning/observations/BEO-GATE/muster-trifft-nur-die-haeufige-schreibweise/evidence/slice-175.md`:7
- **befund:** Geschrieben wurde ``` ``v6.2.0` · `regelwerk/modul-10-review-harness.md`` ```
  — eine Spanne mit **zwei** öffnenden und zwei schließenden Backticks. Das rendert als
  *eine* Code-Spanne, in der die inneren Backticks als Zeichen sichtbar bleiben, statt als
  die zwei getrennten Spannen der Norm. Ursache ist mechanisch: die alten, umschließenden
  Backticks des Inline-Code-Zitats blieben stehen, die Ersetzung brachte eigene mit. Es ist
  das einzige Vorkommen dieser Art im Bestand — eine Suche nach doppelten Backticks über die
  vier einfrierenden Klassen liefert sonst nur Fenced-Block-Anfänge und eine bewusste
  Escape-Schreibweise.
- **verifizierbar:** nein — kein Gate prüft Markdown-Rendering.

### F-7 — Die Verteilungsangabe zur `exempt-paths`-Messung nennt zwei von fünf Klassen nicht

- **kategorie:** LOW
- **quelle:** slice-176 §2.4 (*„28 `version-stale`-Befunde, verteilt auf `done/`-Slices,
  Welle-Dateien und Review-Reports"*)
- **pfad:** slice-176 §2.4
- **befund:** Die **Zahl** ist exakt reproduziert (siehe Negativbefunde), die Aufzählung
  nicht: 13 Befunde in `docs/plan/planning/done/`, 9 in `docs/reviews/`, **4 in
  `harness/conventions/done/`** und **2 in `docs/plan/planning/observations/**/evidence/`**.
  Sechs von 28 liegen also in Klassen, die der Satz nicht nennt — ausgerechnet in den beiden,
  die dieser Slice editiert hat. Die Schlussfolgerung („kann nicht schrumpfen") bleibt davon
  unberührt.
- **verifizierbar:** ja — `exempt-paths` des `versions`-Musters entfernen, `make doc-check`,
  Befundzeilen nach Verzeichnis auszählen.

### F-8 — Klammer-Verschachtelung und doppelte Dateinennung in sieben umgestellten Stellen

- **kategorie:** LOW
- **quelle:** `v6.5.0` · `templates/docs/reviews/review-report.template.md` §Zitier-Form
  (Beispielform ohne zweite Nennung des Dateinamens)
- **pfad:** `docs/plan/planning/done/slice-155-adr-und-carveouts-readme-form.md`:26,33;
  `slice-163`:134,239; `slice-164`:228; `slice-165`:204; `slice-166`:164; `slice-167`:197;
  `docs/plan/planning/observations/BEO-HARNESS/mr-aufloesungs-trigger-ohne-waechter/observation.md`:12
- **befund:** Aus `([`X.template.md`](pfad))` wurde `(`X.template.md` (`v6.2.0` ·
  `templates/…/X.template.md`))` — doppelte Klammern und der Dateiname zweimal in einer Zeile.
  Bei `modul-06`/`modul-08` dasselbe Muster (`(Modul 8 (`v6.2.0` · …))`). Die Norm-Form ist
  `` `v<X.Y.Z>` · `regelwerk/<datei>.md` §<Abschnitt> `` ohne die zusätzliche Hülle. Aussage
  und Auflösbarkeit sind in allen sieben Fällen erhalten; es ist ein Lesbarkeits-Befund, und
  §2.2 nennt nur *zwei* Nachbesserungen von Hand.
- **verifizierbar:** nein — Formfrage ohne Sensor.

---

## INFO

### F-9 — `make trace-check` ist in diesem Arbeitsbaum für jede Range inoperabel

- **kategorie:** INFO
- **quelle:** `AGENTS.md` §5 (Traceability-Durchsetzung) · `harness/README.md` §Sensors
- **pfad:** `Makefile`:208 (Target `trace-check`)
- **befund:** `make trace-check` bricht mit *„Range-Basis-Vorfahren nicht lesbar: object not
  found"* und Exit 2 ab — geprüft mit `RANGE=HEAD~2..HEAD`, `HEAD~1..HEAD` (der dokumentierte
  Default), `HEAD~3..HEAD` und mit expliziten SHAs `96f1511..3fae6d3`. Alle vier gleich. Der
  Klon ist nicht shallow, die Objekte sind lokal auflösbar (`git rev-parse HEAD~3` ok), und
  `make doc-immutable RANGE=HEAD~2..HEAD` läuft über dieselbe Range grün — der Ausfall trifft
  nur das `commits`-Modul. Das ist **nicht** von slice-176 verursacht und hängt in keinem der
  Aggregate (`gates` und `verify` sind grün, ohne ihn zu berühren); es heißt aber, dass die
  Traceability-Zusage lokal nicht gefahren wurde. Für diese Range von Hand geprüft: beide
  Messages nennen `slice-176`, `3fae6d3` zusätzlich `welle-15`. Zuständig: die Gate-Schicht,
  nicht dieser Slice.
- **verifizierbar:** ja — die vier Läufe oben.

---

## Negativbefunde (geprüft, ohne Befund)

- **Vollständigkeit der Umstellung.** `git grep -nE '\]\([^)]*baseline/v[0-9]'` über
  `docs/plan/planning/done/`, `docs/plan/planning/observations/`, `docs/reviews/` und
  `harness/conventions/done/` liefert **kein** Vorkommen. Kein Markdown-Link aus einer
  einfrierenden Klasse zeigt noch in einen Baseline-Pfad. Der in der Vorgeschichte gefundene
  Sonderfall (`../baseline/<tag>/` ohne `.harness/`-Präfix) ist repo-weit ebenfalls leer.
- **Symlinks.** `find` über den Baum: 7 getrackte Symlinks, alle auflösen; die vier
  Baseline-Ziele zeigen auf `v6.5.0`. `make gates` fährt `symlink-check`, dessen Ausgabe das
  bestätigt.
- **Kein zweiter vendorter Stand.** `ls .harness/baseline/` → nur `v6.5.0`.
  `make regelwerk-check` Exit 0: *„54 Datei(en) der Baseline v6.5.0 stimmen mit SHA256SUMS,
  keine unmanifestierte Datei im Baum"*, ein Stand, kein „ungeprüft"-Hinweis — DoD-Punkt 3
  erfüllt.
- **§2.4 reproduziert.** `exempt-paths` des `versions`-Musters aus `.d-check.yml` entfernt,
  `make doc-check` → Exit 2, **exakt 28** `version-stale`-Befunde. Zurückgesetzt,
  `make doc-check` → Exit 0, `git status` leer. Die Zahl in §2.4 stimmt; die Schlussfolgerung
  „zwei Sensoren, zwei Fragen" trägt.
- **Aussage-Erhalt in 13 von 14 Dateien.** Die übrigen Umstellungen (`slice-155`, `161`, `162`,
  `163`, `164`, `165`, `166`, `167`, `evidence/slice-170.md`, `evidence/slice-177.md`, zwei
  `observation.md`, `MR-018`) ersetzen einen Verweis durch dieselbe Angabe in Kennungs-Form;
  der Tag steht danach im Text statt im Pfad. §2.2 trägt für diese Fälle. Einzig F-1 fällt heraus.
- **Aktive `MR-*`-Einträge behalten ihre Links.** Fünf führen `Ersetzt-Baseline-Regel` als
  Link auf `v6.5.0`. Das ist kein Versäumnis: `harness/conventions.md` §Baseline schreibt den
  **mitwandernden** Zeiger für aktive Einträge ausdrücklich vor, und die Zitier-Form gilt
  laut `AGENTS.md` §5 für einfrierende Artefakte.
- **`docs/reviews/` fiel tatsächlich leer aus.** `git grep` gegen `HEAD~2` über `docs/reviews/`
  liefert keinen Baseline-Link — §2.1 Zeile 4 stimmt.
- **Reiner Rename beim Lifecycle-Wechsel.** `git show --name-status -M 47fb99d` weist die
  bewegte Datei als `R100` mit 0 Zeilen Änderung aus; Roadmap und Welle-Datei reisen als
  eigene Dateien mit, wie `AGENTS.md` §3.3 es zulässt. Der Ruhe-Marker der Roadmap ist
  gewechselt (`make doc-planning` grün in `gates`).
- **Commit-Scopes.** `make commit-scope-check RANGE=HEAD~2..HEAD` Exit 0: der eine
  `(planning)`-Commit berührt ausschließlich `docs/plan/planning/`. Der Umsetzungs-Commit
  trägt `feat(harness)` und fällt nicht unter die Scope-Regel.
- **ADR-Immutabilität.** `make doc-immutable RANGE=HEAD~2..HEAD` Exit 0; die Range berührt
  keine ADR — `AGENTS.md` §3.5 unberührt.
- **Größen-Regel.** Drei Liefer-Punkte (Zitier-Form verankert und Bestand umgestellt ·
  `exempt-paths` gemessen · voriger Stand entfernt); Gate-Läufe, Review, Closure-Notiz,
  Register und Risiko-Ausgänge zählen nach `AGENTS.md` §5 nicht mit. Zwei Sub-Areas
  (Planungs-Harness, Harness-Einstieg), beide GF und deklariert — innerhalb *„höchstens zwei
  Schichten"*. `make doc-structure` (in `verify`) grün.
- **Verankerungs-Umfang.** Eine Suche nach der Phrase *Zitier-Form* über `v6.5.0` ·
  `templates/` findet den Block in genau vier Ziel-Formen: Review-Report, Welle-Ergebnisnotiz und beide
  Archiv-Stubs. `AGENTS.md` §5 nennt exakt diese vier — die Aufzählung ist vollständig und
  nicht erfunden. Für die Report-Klasse ist die Form zusätzlich im Reviewer-Skill verankert,
  wo a-check sie beim Schreiben liest (dieser Report folgt ihr).
- **Zitat-Treue.** Das als *wörtlich* ausgewiesene Zitat in §2.2 stimmt zeichengenau mit
  `v6.5.0` · `templates/docs/reviews/review-report.template.md` überein. Die Fassungen in
  `AGENTS.md` §5 und im Reviewer-Skill sind erkennbar eigene Formulierungen und nicht als
  Zitat markiert — korrekt.
- **Abgrenzung „lebende Dokumente".** Der Zusatz in `AGENTS.md` §5 (*„dort ist der Link
  richtig, und `versions` hält ihn aktuell"*) deckt sich mit der Konfiguration: das
  Baseline-Muster in `.d-check.yml` nimmt genau die vier einfrierenden Klassen aus und prüft
  alles übrige, `state.md` eingeschlossen.
- **Selbst-deklarierter Vorgriff.** Dass der Slice-Text in §2.2 eine elidierte Adresse
  (`…/v6.2.0/…`) statt der eigenen Zitier-Form benutzt, ist im Beleg
  `BEO-GATE/versions-sensor-trifft-planungs-vorgriff/evidence/slice-176.md` ausdrücklich als
  Befund notiert — benannt, nicht verschwiegen; kein eigenes Finding.
- **Gate-Läufe.** `make gates` Exit 0 · `make verify` Exit 0 · `make doc-check` Exit 0 ·
  `make regelwerk-check` Exit 0 · `make commit-scope-check RANGE=HEAD~2..HEAD` Exit 0 ·
  `make doc-immutable RANGE=HEAD~2..HEAD` Exit 0 · `make verify-risiko-ausgaenge` Exit 0.
  Alle unpiped in Dateien umgeleitet, Exit-Code separat gelesen.
- **Arbeitsbaum.** Alle vier Mess-Eingriffe (zweimal `.d-check.yml`, einmal der Evidence-Beleg,
  einmal §7 des Slice-Plans) wurden zurückgesetzt; `git status --short` ist am Ende leer, der
  Range-Stand unverändert.

---

## Kategorie-Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 1 |
| MEDIUM | 4 |
| LOW | 3 |
| INFO | 1 |

## Verdikt

**Nicht abschlussreif ohne Klärung von F-1.** Die Kernbehauptung des Slice — der vorige
vendorte Stand ließ sich entfernen, ohne ein eingefrorenes Artefakt anzufassen — hält der
Prüfung stand: kein Markdown-Link aus einer einfrierenden Klasse zeigt noch in einen
Baseline-Pfad, `make regelwerk-check` weist genau einen Stand aus, `make gates` und
`make verify` sind grün, und die Messung in §2.4 ist auf die Zahl genau reproduzierbar. Der
Beweis durch Abwesenheit trägt.

Was ihn nicht trägt, ist die pauschale Rechtfertigung in §2.2. Für 13 der 14 Dateien ist die
Umstellung tatsächlich Form; in einem Fall hat sie das Belegstück eines Fundes durch die
Beschreibung des Fundes ersetzt und den Satz damit unwahr gemacht — und die Gegen-Messung
zeigt, dass gerade diese Stelle nie umgestellt werden musste. Die vier MEDIUM-Befunde teilen
eine Ursache: an drei Stellen (F-2 Konventions-Text, F-3 Restspalt, F-4 Register-Beleg) steht
eine Aussage über den Zustand des Repos, die *nicht* gegen das Repo gemessen wurde, während
die Zahlen, die gemessen wurden, alle stimmen. F-5 zeigt, wie ein solcher Satz einen Sensor
passiert, ohne ihn zu erfüllen.

Die Entscheidung über Übernahme, Rückstellung oder Folge-Slice liegt beim Implementer; dieser
Report kategorisiert und schlägt nichts vor.
