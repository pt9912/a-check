# Review-Report: slice-161 + welle-14 (Regelwerk-Migration `v6.0.0` → `v6.1.0`) — 2026-09-05

**Review-Art:** Plan — geprüft gegen Spec/ADR/Konventionen (Modul 10 §Drei Review-Arten). Gegenstand
ist eine Analyse-zur-Abnahme plus die zugehörige Welle-Eröffnung, nicht Code.

**Gegenstand:** ungetagt, noch nicht committet —
`docs/plan/planning/welle-14-regelwerk-v610-migration.md` (neu),
`docs/plan/planning/in-progress/slice-161-regelwerk-v610-delta-analyse.md` (neu),
`docs/plan/planning/in-progress/roadmap.md` (§Offene Wellen geändert).

**Skill:** `.harness/skills/reviewer.md` @ Stand 2026-09-05 (HEAD `9450268`).
**Modell:** claude-sonnet-5 · **Datum:** 2026-09-05.

**Eingangs-Kontext:**

- `AGENTS.md` §3 (Hard Rules), §5 (Dokumentations-Regeln), §6 (Workflow)
- `harness/conventions.md` §Baseline, §Adaptions-Block, §Modus-Deklaration pro Sub-Area
- Baseline-Regelwerk `modul-05-planning-harness.md`, `modul-06-roadmap.md`
- Ziel-Formen `welle.template.md`, `slice.template.md`
- Präzedenz: `docs/plan/planning/done/wellenlos/slice-135-regelwerk-v600-delta-analyse.md`
  (entpackt aus `slice-135-archiv.zip`), Folge-Slices `slice-136`/`slice-139`/`slice-141`
  (alle `done/wellenlos/`)
- unabhängiger Klon `pt9912/ai-harness-course` (Tags `v6.0.0`/`v6.1.0`) für die Gegenrechnung
- `docs/plan/planning/observations/BEO-HARNESS/`, `BEO-GATE/` (Register-Stichprobe)
- `git log`/`git show` gegen `AGENTS.md` (Provenienz-Prüfung §4.4)

---

## Findings

### F-1 — Falsches Herkunfts-Datum für den Review-Absatz in `AGENTS.md` §6

- `kategorie`: HIGH
- `quelle`: Reviewer-Skill „nachweislich falsche Tatsachenbehauptung, gegen ein Repo-Artefakt
  verifiziert"
- `pfad`: `docs/plan/planning/in-progress/slice-161-regelwerk-v610-delta-analyse.md:155-156`
- `befund`: Der Satz „a-checks eigenes `AGENTS.md` §6 trägt **seit slice-050** einen deutlich
  umfangreicheren Absatz danach ('Vor dem Abschluss, bei jedem Slice mit Code- oder
  Vertragsänderung: Review …')" ist falsch. `git log -p --follow -- AGENTS.md` zeigt genau **eine**
  Einführung dieser Zeichenkette: Commit `7de8569` — „build(harness): slice-159 schneiden --
  Reviewer-Rolle in AGENTS.md §6 verankert (slice-159, slice-158)", datiert **2026-09-05** (heute).
  `git show 0f868d7 -- AGENTS.md` (der slice-050-Commit, 2026-07-25) ändert `AGENTS.md` an keiner
  Stelle. Der von slice-161 zwei Sätze später selbst zitierte Beleg
  (`BEO-HARNESS/behauptete-vollstaendigkeit-extern-gefangen/state.md`) trägt korrekt „seit
  slice-159" — der Widerspruch stand also bereits im eigenen Zitat. Die Kern-Empfehlung von §4.4
  (Rückschnitt auf den `v6.1.0`-Wortlaut + neuer `MR-<NNN>`) bleibt inhaltlich unberührt, aber die
  historische Einordnung („seit sechs Wochen unbelegte Drift" statt „taggleiche Selbst-Ergänzung")
  ist eine andere Tatsache und sollte vor der Abnahme korrigiert werden.
- `verifizierbar`: ja — `git log -p --follow -S"Vor dem Abschluss, bei jedem Slice mit Code" --
  AGENTS.md`.
- `klasse`: Slice-Provenienz aus Gedächtnis statt `git log` behauptet

### F-2 — `make doc-structure` läuft auf dem aktuellen Stand rot (Heading-Kollision im eigenen Slice-Plan)

- `kategorie`: MEDIUM
- `quelle`: `d-check`-Modul `structure` (`DC-FA-STRUCT-001`, Teil des `verify`-Aggregats,
  `AGENTS.md` §4)
- `pfad`: `docs/plan/planning/in-progress/slice-161-regelwerk-v610-delta-analyse.md:134` (Kollision
  mit `:231`)
- `befund`: Die Heading-Regex des Größen-Sensors (`^#+ .*(DoD|Definition of Done)`) trifft nicht
  nur auf „## 7. DoD", sondern auch auf „### 4.3 … nicht als **DoD**-Punkt benannt …" — direkter
  Lauf: `d-check: 402 Datei(en) geprüft, 1 Befund(e)` /
  `slice-161-regelwerk-v610-delta-analyse.md:231 … section-ambiguous … erwartet ist genau einer`,
  `make doc-structure` beendet mit Exit 1. Da `doc-structure` im `verify`-Aggregat hängt, würde ein
  `make verify`-Lauf vor einer etwaigen Closure daran scheitern, solange die Überschrift von §4.3
  das Wort „DoD" enthält.
- `verifizierbar`: ja — bereits verifiziert (`make doc-structure`, Exit 1, s. o.).
- `klasse`: Eigener Slice-Text kollidiert mit der eigenen Struktur-Heuristik

### F-3 — Wellen-Konstrukt `welle-14` gegen den eigenen Präzedenzfall nicht klar begründet

- `kategorie`: MEDIUM
- `quelle`: Baseline-Regelwerk `modul-06-roadmap.md` §Wann Arbeit eine Welle braucht („Ein Trigger,
  der nichts beobachtet, was die Slices nicht ohnehin belegen, ist Zeremonie.")
- `pfad`: `docs/plan/planning/welle-14-regelwerk-v610-migration.md:30-36` vs.
  `slice-161…md:206-229`
- `befund`: Der Closure-Trigger von `welle-14` (Stand-Deklaration an drei Stellen + `make
  gates`/`make verify` grün) ist genau die Art „Mehr", die der strukturell identische vorige Sprung
  (`v5.12.0`→`v6.0.0`) vollständig **wellenlos** trug — `slice-135/136/139/141` liegen alle unter
  `done/wellenlos/`, keine Welle-Datei existiert dafür. `slice-161` §6 selbst schlägt nur **eine**
  Etappe plus einen unabhängigen Folge-Slice vor — dieselbe Größenordnung wie beim Vorbild. Weder
  `welle-14` noch `slice-161` benennt, warum diesmal ein Wellen-Konstrukt nötig ist statt wellenlos
  wie beim ausdrücklich zitierten Vorbild.
- `verifizierbar`: teilweise — die Faktenlage (Vorbild wellenlos) ist verifiziert; ob die Wahl
  „Zeremonie" ist, bleibt Urteil.
- `klasse`: Wellen-Trigger ohne Abgrenzung zum wellenlosen Vorbild

### F-4 — Sichtungs-Schritt (§8) zählt fünf Treffer, das Register trägt mehr passende Einträge

- `kategorie`: MEDIUM
- `quelle`: Baseline-Regelwerk `modul-05-planning-harness.md` §Zwei Schritte vor der
  Modus-Begründung („Vorgelagert — offene Beobachtungen sichten … entfällt nie")
- `pfad`: `docs/plan/planning/in-progress/slice-161-regelwerk-v610-delta-analyse.md:277-286`
- `befund`: `grep -rn "Sub-Area:** Harness-Einstieg" docs/plan/planning/observations/*/*/observation.md`
  findet 7 `offen`-Einträge (2 weitere `verkörpert`, korrekt ausgeschlossen). Vier davon fehlen in
  §8: `BEO-HARNESS/adaption-korrigiert-repo-aussage`, `.../hard-rule-37-ohne-sensor`,
  `.../rueckbau-eintrag-ablage-auslegung`, `.../selbst-archivierung-verdoppelt-abschluss-aufwand`
  — alle bei 1×, alle mit `state.md: **Stand:** offen`. Die Schwellen-Aussage („alle unter 3×")
  bleibt richtig, die Zähl-Aussage („fünf Treffer") ist gegen das Register nicht belegt.
- `verifizierbar`: ja — s. o. (`grep` + `state.md`-Lektüre je Verzeichnis).
- `klasse`: Register-Sichtung unvollständig trotz unveränderter Schwellen-Aussage

### F-5 — Zitierte `BEO-GATE`-Einträge gehören zu einer nicht deklarierten dritten Sub-Area

- `kategorie`: MEDIUM
- `quelle`: Baseline-Regelwerk `modul-05-planning-harness.md` §Zwei Schritte vor der
  Modus-Begründung („Sub-Area-Wahl prüfen … entfällt nie")
- `pfad`: `docs/plan/planning/in-progress/slice-161-regelwerk-v610-delta-analyse.md:270-283`
- `befund`: §8 deklariert zwei berührte Sub-Areas (Vendored Baseline, Harness-Einstieg), zitiert im
  Sichtungs-Schritt aber `BEO-GATE/symlink-ziel-nach-baseline-bump-ungeprueft` und
  `BEO-GATE/sources-modul-asset-integritaet` — beide tragen laut eigenem `observation.md`
  `Sub-Area: Gate-/Werkzeug-Schicht`. Diese dritte Sub-Area ist im „Vorgelagert — Sub-Area-Wahl
  prüfen"-Schritt nicht genannt, obwohl die vorgeschlagene Etappe (§6, Liefer-Punkt 1) `make
  regelwerk-check` berührt — ein Artefakt genau dieser Sub-Area laut
  `harness/conventions.md`-Tabelle.
- `verifizierbar`: ja — `grep "Sub-Area" docs/plan/planning/observations/BEO-GATE/{symlink-ziel-nach-baseline-bump-ungeprueft,sources-modul-asset-integritaet}/observation.md`.
- `klasse`: Register-Sichtung unvollständig trotz unveränderter Schwellen-Aussage (dieselbe Klasse
  wie F-4, andere Facette — Sub-Area-Wahl statt Zähl-Vollständigkeit)

### F-6 — Abschnitts-Nummerierung nicht monoton

- `kategorie`: LOW
- `quelle`: Maintainability / Konsistenz zum Präzedenzfall
- `pfad`: `docs/plan/planning/in-progress/slice-161-regelwerk-v610-delta-analyse.md:231,246,264,268`
- `befund`: Dateireihenfolge der Überschriften ist `1,2,3,4,5,6,7` („## 7. DoD"), dann `6a`
  („Risiken"), `7a` („Closure-Notiz"), `8` — §7 steht im Text vor §6a/§7a. Der zitierte
  Präzedenzfall `slice-135` nummeriert dieselbe Analyse-Form durchgehend `1`–`9` ohne
  Buchstaben-Suffixe.
- `verifizierbar`: ja — Sichtprüfung der Überschriften-Reihenfolge.
- `klasse`: Nicht-monotone Abschnitts-Nummerierung in Analyse-Slice

### F-7 — Risiko-Ausgänge in §6a vor dem Übergang nach `done/` bereits festgeschrieben

- `kategorie`: LOW
- `quelle`: Baseline-Regelwerk `modul-05-planning-harness.md` §Offene Risiken werden bei Closure
  aufgelöst („bekommt beim Übergang nach `done/` genau einen Ausgang")
- `pfad`: `docs/plan/planning/in-progress/slice-161-regelwerk-v610-delta-analyse.md:246-262` vs.
  `:264-266` (§7a, explizit noch offen)
- `befund`: §6a weist bereits Ausgänge zu (u. a. „Folge-Slice, noch keine ID"), während §7a
  (Closure-Notiz) ausdrücklich als „wird beim Übergang nach `done/` ausgefüllt" markiert ist. Die
  Zuweisung geschieht damit vor dem Übergang, den die Regel als Zeitpunkt nennt. Folgenlos für
  diese Analyse (sie enthält keine Artefaktänderung, die davon abhängt), aber eine Abweichung vom
  Wortlaut.
- `verifizierbar`: ja — Sichtprüfung der beiden Abschnitte.
- `klasse`: Risiko-Ausgang vor dem Lifecycle-Übergang zugewiesen

### F-8 — Register führt eine nicht deklarierte Sub-Area-Bezeichnung

- `kategorie`: INFO
- `quelle`: `harness/conventions.md` §Modus-Deklaration pro Sub-Area
- `pfad`: `docs/plan/planning/observations/BEO-HARNESS/sensor-ohne-dod-phrase-wirkungslos/observation.md:3`
- `befund`: Der Eintrag trägt `Sub-Area: Harness-Tooling` — eine Bezeichnung, die die
  Modus-Deklarationstabelle in `harness/conventions.md` nicht führt (dort nur „Harness-Einstieg").
  Vorbestehende Register-Inkonsistenz, nicht durch `slice-161`/`welle-14` verursacht; sie erklärt
  aber, warum eine rein verzeichnisbasierte Sichtung (F-4/F-5) nicht mechanisch zuverlässig ist —
  das Urteil bleibt beim Menschen (Modul 6), macht Vollständigkeits-Prüfungen aber fehleranfällig.
  Kein Handlungsbedarf innerhalb dieses Reviews; Hinweis für einen künftigen Register-Aufräum-Slice.
- `verifizierbar`: ja — Dateiinhalt wie zitiert.
- `klasse`: —

## Negativbefunde

- geprüft, ohne Befund: Kern-Zahlen §2 (6 Dateien, `+38/−1`) — unabhängig gegen einen frischen Klon
  von `pt9912/ai-harness-course` nachgerechnet (`git diff --stat v6.0.0 v6.1.0 -- lab/regelwerk
  lab/templates`), exakte Übereinstimmung.
- geprüft, ohne Befund: wörtliches Zitat §4.4 (neuer `AGENTS.template.md`-Absatz) und Tabellenzeile
  §4.1 (`harness/README.template.md`) — Zeichen für Zeichen deckungsgleich mit `git diff v6.0.0
  v6.1.0 -- lab/templates/…`.
- geprüft, ohne Befund: Behauptung „`AGENTS.template.md` endet nach Schritt 8" (§4.4) —
  `.harness/baseline/v6.0.0/templates/AGENTS.template.md` endet Zeile 236/237 nach Schritt 8, kein
  Rollenwechsel-Absatz in der vendorten `v6.0.0`-Fassung.
- geprüft, ohne Befund: Behauptung „kein `MR`-Eintrag für den Rollenwechsel-Absatz" (§4.4) —
  `harness/conventions.md` Adaptions-Block enthält keinen solchen Eintrag (grep über
  „Rollenwechsel"/„Review-Pflicht"/„Code- oder Vertragsänderung").
- geprüft, ohne Befund: Behauptung „Welle 117 berührt nur `docs/`" (§1) — `git log --name-only
  v6.0.0..v6.1.0` zeigt, dass Commit `1ab2657` (Welle 117) ausschließlich `CHANGELOG.md`,
  `docs/steering-loop-team.md`, `docs/zeitdokument-archiv.md` ändert.
- geprüft, ohne Befund: Beobachtungs-Register-Existenz und Zähler-Stand der fünf **zitierten**
  Pfade (§8) — alle fünf Verzeichnisse existieren, je genau eine `evidence/`-Datei, `state.md:
  offen` (1×, unter der Schwelle).
- geprüft, ohne Befund: relative Markdown-Links in beiden neuen Dateien — `make doc-check`: `402
  Datei(en) geprüft, 0 Befund(e)`.
- geprüft, ohne Befund: `make doc-planning` — Ruhe-Marker korrekt entfernt, da `in-progress/` nicht
  leer ist; Modul `planning` meldet 0 Befunde.
- geprüft, ohne Befund: WIP-Limit = 1 — `in-progress/` enthält neben `roadmap.md` nur
  `slice-161-regelwerk-v610-delta-analyse.md`.
- geprüft, ohne Befund: Größen-Regel des §6-Vorschlags (3 Liefer-Punkte, Vendored-Baseline- +
  Harness-Einstieg-Schicht) und die Begründung „§4.4 würde die Grenze sprengen" (3 bestehende
  Punkte + Textänderung `AGENTS.md` + neuer `MR`-Eintrag = 5 > 3) — rechnerisch korrekt.
- geprüft, ohne Befund: Start-Trigger von `welle-14` — extern (Release-Tag + Maintainer-Weisung),
  kein Ergebnis der Welle selbst, nicht zirkulär.
- geprüft, ohne Befund: Kopf-Felder beider neuer Dateien (Lifecycle/Welle/Bezug/Berührte
  Spec-Stellen/Verantwortlich/Autor bzw. Zielmeilenstein/Verantwortlich/Datum) — vollständig,
  Reihenfolge entspricht der Ziel-Form.
- geprüft, ohne Befund: Hard Rules `AGENTS.md` §3.1–§3.7 — keine Verstöße; keine Harness-Lüge
  (kein erfundenes Gate-Target, keine erfundene ID) in beiden neuen Dateien.

## Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 1 |
| MEDIUM | 4 |
| LOW | 2 |
| INFO | 1 |

**Finding-Klassen dieses Laufs:** Slice-Provenienz aus Gedächtnis statt `git log` behauptet ·
Eigener Slice-Text kollidiert mit der eigenen Struktur-Heuristik · Wellen-Trigger ohne Abgrenzung
zum wellenlosen Vorbild · Register-Sichtung unvollständig trotz unveränderter Schwellen-Aussage ·
Nicht-monotone Abschnitts-Nummerierung in Analyse-Slice · Risiko-Ausgang vor dem
Lifecycle-Übergang zugewiesen

## Verdikt

**Merge-blockierend:** ja, für §4.4 der Analyse in der aktuellen Fassung (F-1: falsche
Provenienz-Behauptung, vor Abnahme zu korrigieren) und für eine etwaige Closure von `slice-161`
selbst (F-2: `make doc-structure` läuft rot). Der eigentliche Liefer-Vorschlag — eine Etappe
gemäß §6 plus ein Folge-Slice für §4.4 — ist davon **inhaltlich unberührt** und umsetzungsreif,
sobald F-1/F-2 behoben sind. F-3–F-5 sollten vor der Maintainer-Abnahme geklärt werden (Wellen-Wahl
bzw. Vollständigkeit der Register-Sichtung), blockieren aber nicht die Kern-Entscheidung, welche
Etappe umgesetzt wird.

**Übergabe:** Findings gehen an die Implementer-Session zurück (Rückkante Review → Plan). F-1 und
F-2 sind vor jeder Abnahme-Entscheidung zu beheben; F-3 verlangt eine explizite
Maintainer-Entscheidung (Wellen-Konstrukt behalten oder auf wellenlos zurückbauen wie beim
Vorbild) statt einer stillen Übernahme. Die Finding-Klassen gehen bei Closure von `slice-161` in
das Beobachtungs-Register. Dieser Report ist Lauf-Beleg und ersetzt keine Verifikation (Modul 11).
