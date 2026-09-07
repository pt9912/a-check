# Review-Report: slice-179 — 2026-09-07

**Review-Art:** Plan-Review (unabhängiger Lauf) — geprüft wird der Durchgang
gegen die kanonischen Quellen: die sieben aktiven `MR`-Einträge, die vendored
Baseline `v6.5.0` und `AGENTS.md` §3/§4/§5. Kontext-Trennung eingehalten: dieser
Lauf hat den Gegenstand nicht verfasst und den Verfasser-Kontext nicht geerbt.

**Gegenstand:** slice-179, Commit-Range `HEAD~1..HEAD` (`7b84f82`,
`feat(harness): slice-179 Etappe B`) — zwei Dateien, `+148/−22`.

**Skill:** `.harness/skills/reviewer.md` @ `3fae6d3` (slice-176)
**Modell:** `claude-opus-5[1m]` · **Datum:** 2026-09-07

> **Zitier-Form** *(dieser Block bleibt stehen — er ist Norm, kein
> Ausfüll-Hinweis)*. Dieser Report friert ein; was er zitiert, bewegt sich
> weiter. Deshalb **Kennung, nicht Adresse** — `slice-NNN` statt seines
> Lifecycle-Pfads, `make <target>` statt eines Links auf die Sensor-Datei, eine
> Baseline-Stelle als **Tag + Pfad in Inline-Code** statt als Link
> (`v<X.Y.Z>` · `regelwerk/<datei>.md` §<Abschnitt>). Der vendored Baum trägt
> genau einen Tag; der nächste Sprung löscht den alten, und ein Link darauf
> färbt ein Artefakt rot, das niemand mehr anfassen darf. Das `pfad`-Feld auf
> den **geprüften Gegenstand** ist davon nicht betroffen — es hält den Stand
> des Laufs fest und darf das.

**Eingangs-Kontext** (die Verträge, gegen die geprüft wurde):

- slice-179 (Plan, `in-progress/`), slice-174 §3.4 (die gestellte Frage),
  slice-175 §2.2 (die Datei-Ebene), slice-176 (Zitier-Form), slice-177
- die sieben aktiven Einträge unter `harness/conventions/` und die Tabelle
  §Aktive Adaptionen in `harness/conventions.md`
- `v6.5.0` · `regelwerk/grundlagen-harness-dateien.md`,
  `regelwerk/grundlagen-referenz-richtung.md`,
  `regelwerk/grundlagen-source-precedence.md`,
  `regelwerk/modul-05-planning-harness.md`, `regelwerk/modul-06-roadmap.md`,
  `regelwerk/modul-08-agentenrollen.md`, `regelwerk/modul-15-observability.md`
- `AGENTS.md` §3 (Hard Rules), §4 (Gates), §5 (Doku-Regeln), `.d-check.yml`
- Kurs-Klon (`v6.2.0`/`v6.5.0`) für die Delta-Messungen, die aus dem Repo
  allein nicht reproduzierbar sind (nur `v6.5.0` liegt vendored)
- ADRs: keine berührt (39 im Bestand, keine nennt `exempt-paths` oder
  `version-stale` — nachgezählt, siehe Negativbefunde)

---

## Findings

### F-1 — Ausnahme-Verdikt gilt für vier von fünf Globs, steht aber unbedingt da

- `kategorie`: HIGH
- `quelle`: `v6.5.0` · `regelwerk/grundlagen-harness-dateien.md`
  §harness/README.md als Einstiegspunkt (*„Die Grenze: Sie gilt für
  einfrierende Artefakte. Ein lebendes … verlinkt weiter, und dort ist der
  Bruch das gewollte Signal"*); `AGENTS.md` §3.6
- `pfad`: `harness/sensors/doc-check.md:34–42`; slice-179 §3.2
- `befund`: Das Verdikt *„Die Ausnahme bildet also ab, worüber die Regel
  spricht — sie senkt keine Schwelle, und `AGENTS.md` §3.6 greift nicht"*
  steht ohne Einschränkung, obwohl der Glob `docs/reviews/**` genau einen
  **lebenden** Zeiger mit ausnimmt: `docs/reviews/README.md:7` trägt
  `.harness/baseline/v6.5.0/templates/docs/reviews/review-report.template.md`,
  und dieser Pin wurde bei jeder Migration nachgezogen (`92e1f64` auf `v6.2.0`,
  `f49913a` auf `v6.5.0`) — für ihn ist der Bruch nach der zitierten Stelle das
  *gewollte* Signal, nicht ein Zeiger, der nicht nachgezogen werden darf. Der
  Satz zwei Zeilen darüber in derselben Datei hält denselben Sachverhalt fest
  (*„Zwei lebende Zeiger fallen dabei mit heraus"*); die vier übrigen Globs
  (`harness/conventions/done/**`, `docs/plan/planning/done/**`,
  `observations/**/observation.md`, `observations/**/evidence/**`) tragen
  nachgeprüft nur eingefrorene Zeiger.
- `verifizierbar`: ja — `grep -rlE '(\.harness|\.\.)/baseline/v[0-9]+\.[0-9]+\.[0-9]+/'`
  über die fünf Globs listet `docs/reviews/README.md` als einzigen lebenden
  Treffer; `git log -S'baseline/v6.5.0' -- docs/reviews/README.md` belegt den
  Nachzug. Ein Gate dafür gibt es nicht.
- `klasse`: „Ausnahme-Verdikt breiter als die geprüfte Menge"

### F-2 — Quelle der zitierten Baseline-Stelle ist ein anderer Abschnitt

- `kategorie`: HIGH
- `quelle`: Reviewer-Skill §Klassifikation, HIGH — *nachweislich falsche
  Tatsachenbehauptung (gegen ein Repo-Artefakt verifiziert)*
- `pfad`: `harness/sensors/doc-check.md:35–39`
- `befund`: Der Text schreibt *„`v6.5.0` ·
  `regelwerk/grundlagen-harness-dateien.md` §Was ein Kommentar trägt zieht die
  Linie selbst"* und zitiert danach *„Die Grenze: Sie gilt für einfrierende
  Artefakte …"*. Dieser Satz steht in der vendorten Datei in Zeile 317–321 und
  damit im Abschnitt §harness/README.md als Einstiegspunkt (Zeilen 165–327);
  §Was ein Kommentar trägt reicht von Zeile 67 bis 163 und enthält weder das
  Zitat noch die Wörter *Ventil*, *Gate-Senkung* oder *nachgezogen*. Das Zitat
  selbst ist wortgetreu — falsch ist die Fundstelle, und sie steht in genau dem
  Absatz, der mit *„geprüft an der Baseline, nicht angenommen"* wirbt.
- `verifizierbar`: ja — `grep -n '^### ' .harness/baseline/v6.5.0/regelwerk/grundlagen-harness-dateien.md`
  gegen die Zeilennummer des Zitats. Kein Gate prüft §-Verweise in Prosa.
- `klasse`: „§-Verweis auf die falsche Abschnitts-Überschrift"

### F-3 — `MR-014` als *permanent* geführt, obwohl sein Eintrag einen erreichbaren Trigger nennt

- `kategorie`: MEDIUM
- `quelle`: `harness/conventions/MR-014-keine-agenten-telemetrie.md`
  §Auflösungs-Trigger; `AGENTS.md` §5 (Diskrepanz-Trichter: *Trigger nie
  erreichbar ⇒ permanente ADR*)
- `pfad`: slice-179 §3.1, Zeile 84 (Spalte *Trigger eingetreten?*)
- `befund`: Die Zeile lautet *„nein (permanent: a-check ruft kein Modell auf)"*.
  `MR-014` deklariert als Auflösungs-Trigger *„sobald Agenten-Läufe **im Repo
  selbst** abrechenbar werden"* — eine erreichbare Bedingung, nicht
  Permanenz; *permanent* führen die Einträge `MR-012` und `MR-020` ausdrücklich
  im Trigger-Feld, `MR-014` nicht. Der Durchgang stuft damit einen bedingten
  Eintrag in die Klasse um, die nach `AGENTS.md` §5 eine ADR verlangt.
- `verifizierbar`: nein — kein Gate hält Durchgangs-Urteil gegen Trigger-Feld.
- `klasse`: „Trigger-Klasse im Durchgang umgeschrieben"

### F-4 — Die von `MR-019` selbst benannte Rückbau-Bedingung ist eingetreten und bleibt unbewertet

- `kategorie`: MEDIUM
- `quelle`: `harness/conventions/MR-019-review-dod-opt-in.md`
  §Ersetzt-Baseline-Regel; `harness/conventions.md` §Aktive Adaptionen, Zeile
  zu `MR-019`
- `pfad`: slice-179 §3.1, Zeile 87; §2, Zeilen 57–61
- `befund`: `MR-019` begründet sein leeres `Ersetzt-Baseline-Regel`-Feld damit,
  der Treiber sei *„eines, das a-check noch nicht vendored hat; kein netzlos
  auflösbarer Anker möglich"*, und benennt als Rückbau-Bedingung *„sobald
  `v6.2.0` vendored ist und auf das dann vendorte Template gezeigt werden
  kann"*. Diese Bedingung ist erfüllt: `.harness/baseline/v6.5.0/templates/docs/plan/planning/slice.template.md`
  trägt die Review-DoD-Zeile in Zeile 83–85 — der Durchgang misst genau diese
  Zeile, zieht die Folge aber nicht, und die lebende Index-Zeile in
  `harness/conventions.md` sagt weiter *„Treiber ist ein noch nicht vendortes
  Template"*. Zugleich führt die Durchgangs-Tabelle für `MR-019` eine
  *Ersetzte Regel* („Review-Report als unbedingter DoD-Checkbox-Punkt"), die
  der Eintrag selbst mit `—` verneint, während sie für `MR-020` das `—`
  übernimmt.
- `verifizierbar`: ja für die Tatsache (`grep -n -i review` auf das vendorte
  Template); nein für die Bewertung.
- `klasse`: „eingetretene Rückbau-Bedingung ohne Ausgang"

### F-5 — Steering-Loop-Eintrag zeigt auf den eigenen Slice als Zielort

- `kategorie`: MEDIUM
- `quelle`: `v6.5.0` · `regelwerk/modul-06-roadmap.md` §Wellen-Closure-Prozedur,
  Schritt 3 (a) Anker-Paarung: *„Wo das Feld steht, existiert der Zielort und
  trägt `seit welle-<NN>` bzw. `seit slice-<NNN>`"*
- `pfad`: slice-179 §8, Zeilen 198–206
- `befund`: Der Eintrag endet mit *„— liegt in
  `docs/plan/planning/done/slice-179-…md` §3.1 als Muster für den nächsten
  Durchgang"*. Das Pflichtfeld `liegt in` löst die Anker-Paarung aus; der
  genannte Zielort ist §3.1 desselben Slice, trägt keinen Anker `seit
  slice-179` und ist selbst ein Zeitdokument, das mit welle-15 zum Stub
  gekürzt wird. Die Commit-Message desselben Commits verwirft diese
  Platzierung für den anderen Lerngegenstand ausdrücklich („nicht in einem
  Slice, den niemand mehr oeffnet"); für die geschärfte Regel selbst gibt es
  keinen Zielort in `AGENTS.md`, einem Gate, einem Skill oder einem
  `MR`-Eintrag.
- `verifizierbar`: ja — die Paarung läuft bei der welle-15-Closure; sie ist im
  Repo kein `make`-Target, sondern Handschritt.
- `klasse`: „Lerneintrag verkörpert im eigenen Zeitdokument"

### F-6 — Einfrier-Zitierform in einem lebenden Dokument nimmt den Pin aus `versions`

- `kategorie`: MEDIUM
- `quelle`: `AGENTS.md` §5, §Zitier-Form in einfrierenden Artefakten:
  *„**Nicht** betroffen: lebende Dokumente — dort ist der Link richtig, und
  `versions` hält ihn aktuell"*; `.d-check.yml`, `versions.pin-pattern`
- `pfad`: `harness/sensors/doc-check.md:35–36`
- `befund`: Die neue Passage nennt die Baseline-Stelle als `` `v6.5.0` ·
  `regelwerk/grundlagen-harness-dateien.md` `` — die Form für einfrierende
  Artefakte. `harness/sensors/doc-check.md` ist ein lebendes Dokument (kein
  `sensors/done/`, kein Status-Feld, wird mit dem Gate fortgeschrieben); es ist
  von keinem `exempt-paths` erfasst. Das Muster
  `(?:\.harness|\.\.)/baseline/(v\d+\.\d+\.\d+)/` kann die Inline-Form nicht
  treffen, also steht die Versionsangabe `v6.5.0` in dieser Datei ab sofort
  außerhalb der `versions`-Deckung — in genau der Datei, die diese Deckung
  beschreibt. Vor diesem Commit trug keine Datei unter `harness/sensors/` die
  Inline-Form.
- `verifizierbar`: ja — `make doc-check` bleibt nach einem künftigen Bump grün,
  während die Angabe veraltet; heute belegbar daran, dass unter `harness/sensors/`
  sonst keine Datei die Inline-Form trägt.
- `klasse`: „Einfrier-Zitierform in lebendem Dokument"

### F-7 — `MR-015`-Beleg zitiert zwei Zeilen außerhalb des Abschnitts, den sein Zeiger nennt

- `kategorie`: LOW
- `quelle`: `harness/conventions/MR-015-welle-closure-ohne-replay.md`
  §Ersetzt-Baseline-Regel (`modul-06-roadmap.md` §Wellen-Closure-Prozedur)
- `pfad`: slice-179 §3.1, Zeile 85
- `befund`: Als Beleg stehen dort *„Zeile 33 und 319 in `modul-06-roadmap.md`,
  beide unverändert"*. Zeile 33 liegt in §Wann Arbeit eine Welle braucht
  (30–54), Zeile 319 in §Regeln gegen typische Fehlannahmen (ab 314); der
  Abschnitt, den `MR-015` als ersetzte Regel benennt, reicht von 151 bis 313
  und trägt die Replay-Zusage in Zeile 195 — die im Beleg fehlt. Das Verdikt
  bleibt davon unberührt (auch 195 ist unverändert); die einzige inhaltliche
  Änderung des Moduls liegt in Eröffnung Schritt 1.
- `verifizierbar`: ja — `grep -n '^### '` plus `git diff -w v6.2.0 v6.5.0` im
  Kurs-Klon.
- `klasse`: „Beleg-Zeile außerhalb des benannten Abschnitts"

### F-8 — „zwei Absätze vorher" und „der Kurs sagt es" für `versions`

- `kategorie`: LOW
- `quelle`: `v6.5.0` · `regelwerk/grundlagen-harness-dateien.md`
  §harness/README.md als Einstiegspunkt
- `pfad`: slice-179 §3.2, Zeilen 116 und 193
- `befund`: Der Grenze-Absatz (Zeilen 317–321) steht unmittelbar vor dem
  Reparatur-Absatz (323–326), nicht zwei Absätze davor. Und der Satz *„Für
  `versions` gilt etwas anderes, und der Kurs sagt es zwei Absätze vorher"*
  legt der Quelle eine Aussage über die Versions-Prüfung bei; die zitierte
  Stelle zieht die Grenze für die Regel *Kennung statt Adresse* — die
  Übertragung auf das `versions`-Muster ist a-checks eigene Folgerung.
- `verifizierbar`: ja — Absatzabstand im vendorten Text.
- `klasse`: „eigene Folgerung als Quellenaussage formuliert"

### F-9 — Der zweite der „zwei lebenden Zeiger" hat seit slice-176 kein Objekt mehr

- `kategorie`: LOW
- `quelle`: Bestandsmessung gegen `harness/conventions/done/MR-018-review-pflicht-v610-wortlaut.md`
- `pfad`: `harness/sensors/doc-check.md:30–32` (der Satz, an den die neue
  Passage anschließt)
- `befund`: Der Satz nennt *„Zwei lebende Zeiger … (`docs/reviews/README.md`,
  der Vorlagen-Link in `MR-018`)"*. `MR-018` trägt seit `3fae6d3` (slice-176)
  keinen `baseline/v…`-Pfad mehr, sondern die Inline-Zitierform; das
  `versions`-Muster hätte dort nichts zu treffen, und `MR-018` ist als
  aufgelöster Eintrag ohnehin eingefroren, nicht lebend. Der Durchgang hat
  genau diesen Absatz erneut geöffnet, den Bestand aber nicht nachgemessen.
- `verifizierbar`: ja — `grep -n 'baseline/v' harness/conventions/done/MR-018-*.md`
  liefert keine Zeile.
- `klasse`: „Bestandsaussage im angefassten Absatz nicht nachgemessen"

### F-10 — Die offene Frage wird unqualifiziert als „`exempt-paths` in `.d-check.yml`" geführt

- `kategorie`: LOW
- `quelle`: slice-174 §3.4 (*„fünf `exempt-paths`-Klassen"*, gesetzt mit
  slice-173); `.d-check.yml`
- `pfad`: slice-179 §2, Zeile 63
- `befund`: `.d-check.yml` führt `exempt-paths` in vier Modulen — fünfmal unter
  `ids`, einmal unter `matrix` (die Grandfather-Menge aus `MR-012`), einmal
  unter `versions` (die fünf Zeitdokument-Klassen) und dreimal unter
  `structure`. Bewertet ist ausschließlich der `versions`-Fall; die
  Überschrift der offenen Frage nennt aber den Schlüssel ohne Modul, und §3.2
  schließt mit *„`exempt-paths` bildet ab, worüber die Regel spricht"*. Die
  `ids`-Vorkommen sind weder von einem `MR`-Eintrag noch von der
  Grandfather-Regel in `AGENTS.md` §5 gedeckt.
- `verifizierbar`: ja — `grep -c 'exempt-paths' .d-check.yml`.
- `klasse`: „Befund über eine Teilmenge auf den Schlüssel verallgemeinert"

### F-11 — `make trace-check` in dieser Umgebung nicht lauffähig

- `kategorie`: INFO
- `quelle`: `AGENTS.md` §5 (Traceability), `ADR-0021`
- `pfad`: `Makefile:208`
- `befund`: `make trace-check` bricht mit *„d-check: error: Range-Basis-Vorfahren
  nicht lesbar: object not found"* ab — auch mit einer Range aus dem Altbestand
  (`f74b7d4..da7695e`) und ohne `shallow`/`alternates` im Klon; die Ursache
  liegt außerhalb dieses Slice. Die Traceability des Commits ist von Hand
  geprüft: die Message nennt `slice-179` und `welle-15`. `make gates`,
  `make verify`, `make doc-structure` und
  `make commit-scope-check RANGE=HEAD~1..HEAD` liefern je Exit 0.
- `verifizierbar`: ja — der Lauf selbst.
- `klasse`: „Sensor in der Review-Umgebung nicht lauffähig"

### F-12 — Der Beleg des einzigen Trigger-Kandidaten ist aus dem Repo allein nicht reproduzierbar

- `kategorie`: INFO
- `quelle`: `harness/conventions.md` §Baseline (genau ein Stand liegt vendored)
- `pfad`: slice-179 §2, Zeilen 57–61
- `befund`: Die Aussage *„`slice.template.md` hat sich in `v6.5.0` geändert
  (`+36/−10`), die Review-DoD-Zeile darin jedoch nicht"* beruht auf
  `git diff -w v6.2.0 v6.5.0` gegen einen Kurs-Klon; `v6.2.0` liegt nicht mehr
  vendored, ein späterer Leser kann sie im Repo nicht nachrechnen. Dieser Lauf
  hat sie gegen den Klon reproduziert und bestätigt: `36/10`, und die Zeile ist
  bytegleich (`v6.2.0` Zeilen 61–63, `v6.5.0` Zeilen 83–85).
- `verifizierbar`: ja, aber nur mit externem Klon — netzlos nicht.
- `klasse`: „Beleg außerhalb des Repos gemessen"

## Negativbefunde

- geprüft, ohne Befund: **Die Zahl sieben.** `harness/conventions/` führt genau
  sieben Dateien (`MR-011`, `MR-012`, `MR-014`, `MR-015`, `MR-016`, `MR-019`,
  `MR-020`), und §Aktive Adaptionen in `harness/conventions.md` trägt genau
  sieben Zeilen — die Mengen decken sich.
- geprüft, ohne Befund: **`MR-011`, `MR-012`, `MR-016` — Aussagen-Ebene.**
  `grundlagen-source-precedence.md`, `grundlagen-referenz-richtung.md` und
  `modul-08-agentenrollen.md` sind zwischen `v6.2.0` und `v6.5.0` inhaltlich
  unverändert (`git diff -w --numstat` nennt sie nicht). Die im Roh-Diff
  auffälligen Zeilen der Matrix-Tabelle (ADR→ADR normativ, Straten-Tabelle)
  sind reine Spalten-Neuausrichtung; die Decken-Regel und die
  ADR→ADR-Aussage standen wortgleich schon in `v6.2.0`.
- geprüft, ohne Befund: **`MR-015` — Aussagen-Ebene.** Die einzige inhaltliche
  Änderung in `modul-06-roadmap.md` liegt in Eröffnung Schritt 1 (Verweis auf
  §1 *Ziel und Abgrenzung*); alle drei Replay-Stellen (33, 195, 319) sind
  unberührt. Die Trennung von Datei- und Aussagen-Ebene, die der Slice als
  Lerneintrag zieht, trägt an diesem Fall.
- geprüft, ohne Befund: **`MR-019` — die übernommene Hälfte.**
  `modul-05-planning-harness.md` Zeile 167 (*„Nicht gezählt: Gate-Läufe,
  Review-Report, …"*) ist in beiden Tags identisch; die Adaption bezieht sich
  also weiter auf einen bestehenden Gegenstand.
- geprüft, ohne Befund: **`MR-020`.** Der Eintrag korrigiert eine Repo-Aussage
  und trägt „permanent" im eigenen Trigger-Feld; die Tabellenzeile gibt beides
  unverändert wieder.
- geprüft, ohne Befund: **Form der Risiko-Ausgänge (§7).** Zwei Risiken, je
  genau ein Ausgang aus der geschlossenen Dreier-Menge (*weiter offen* →
  Register, *entfallen* → gestrichen mit Begründung). Der erste Ausgang ist
  tragfähig: die Klasse `BEO-HARNESS/baseline-regel-nie-erwogen-weil-bestand-sie-verletzt`
  steht bei 2×, und in slice-179 ist kein Auftreten belegbar — der Slice hat
  keine übersehene Regel gefunden, sondern nicht danach gesucht, und ein
  Vorkommen ohne Fund bewegt den Zähler nicht.
- geprüft, ohne Befund: **§9 Sichtung und Sub-Area-Wahl.** Beide einschlägigen
  Register-Einträge sind genannt und stehen bei 2×, keiner erreicht mit diesem
  Slice die Schwelle; `harness/` deckt `harness/sensors/` in der
  Modus-Deklaration, die Zuordnung zu *Harness-Einstieg* folgt der Pfadliste.
- geprüft, ohne Befund: **Größen-Regel und Slice-Form.** Drei gezählte
  Liefer-Punkte, eine Sub-Area, zwei berührte Dateien; `make doc-structure`
  Exit 0 über 494 Dateien (DoD-Größe, Closure-Struktur, Lerneintrag-Form,
  Kopffelder, AC-Form).
- geprüft, ohne Befund: **Commit-Scope und Hard Rules.**
  `make commit-scope-check RANGE=HEAD~1..HEAD` Exit 0 (Scope `harness`, kein
  `(planning)`); keine Host-Toolchain im Diff, keine Suppression, keine ADR
  berührt, kein Spec-Stratum abwärts referenziert.
- geprüft, ohne Befund: **Links, Anker, Kennungs-Linkpflicht,
  Lifecycle-Invariante.** In `make gates` enthalten, Exit 0 — die
  `../done/…`-Verweise des Plans lösen aus jeder Lifecycle-Position auf.
- geprüft, ohne Befund: **ADR-Bestand.** 39 ADRs; keine nennt `exempt-paths`
  oder `version-stale` — die Prämisse aus slice-174 §3.4 stimmt in diesem
  Punkt.
- geprüft, ohne Befund: **Repo-weite Gate-Lage.** `make gates` und
  `make verify` je Exit 0 auf dem geprüften Stand, unabhängig nachgefahren.

## Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 2 |
| MEDIUM | 4 |
| LOW | 4 |
| INFO | 2 |

**Finding-Klassen dieses Laufs:** „Ausnahme-Verdikt breiter als die geprüfte
Menge" · „§-Verweis auf die falsche Abschnitts-Überschrift" · „Trigger-Klasse
im Durchgang umgeschrieben" · „eingetretene Rückbau-Bedingung ohne Ausgang" ·
„Lerneintrag verkörpert im eigenen Zeitdokument" · „Einfrier-Zitierform in
lebendem Dokument" · „Beleg-Zeile außerhalb des benannten Abschnitts" · „eigene
Folgerung als Quellenaussage formuliert" · „Bestandsaussage im angefassten
Absatz nicht nachgemessen" · „Befund über eine Teilmenge auf den Schlüssel
verallgemeinert"

## Verdikt

**Merge-blockierend:** ja — zwei HIGH und vier MEDIUM.

Der Kern des Slice hält: Die sieben Einträge sind vollständig erfasst, sechs
der sieben Verdikte sind an der vendorten Baseline nachgerechnet und richtig,
und der Lerneintrag über die zwei Ebenen trägt an seinem eigenen Beispiel
(`MR-015`). Die Widerlegung der `exempt-paths`-Prämisse aus slice-174 ist
sachlich zutreffend, wo sie den Kurs liest — der Satz von der *„Gate-Senkung
mit eigener Begründungslast"* steht tatsächlich im Absatz über Adressen in
eingefrorenen Artefakten, und die Vermeidung ist tatsächlich die vom Text
empfohlene Antwort.

Blockierend sind zwei Dinge, die beide an derselben Stelle sitzen und beide
denjenigen Fehler wiederholen, den der Slice slice-174 vorwirft. F-1: Das
Verdikt ist für vier der fünf Globs belegt und steht für alle fünf da; der
fünfte Glob nimmt einen lebenden Zeiger mit heraus, den die zitierte Stelle
ausdrücklich ausnimmt, und derselbe Absatz in derselben Datei sagt das zwei
Zeilen weiter oben. F-2: Der Absatz, der sich auf *„geprüft an der Baseline,
nicht angenommen"* beruft, nennt den falschen Abschnitt als Fundstelle.

**Übergabe:** Findings gehen an den Implementer; F-4 hat zusätzlich eine
Rückkante in die Planung (der Eintrag `MR-019` ist nach seiner eigenen
Bedingung fällig, und die lebende Index-Zeile in `harness/conventions.md`
behauptet inzwischen etwas Unwahres). Die **Finding-Klassen** gehen in die
Slice-Closure §7 und von dort in den Zähler. Dieser Report ist ein
Lauf-Beleg und ersetzt keine Verifikation.
