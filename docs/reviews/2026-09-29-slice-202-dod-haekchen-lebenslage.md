# Review-Report: slice-202 — 2026-09-29

**Review-Art:** Adversarisches Code-/Harness-Review — geprüft gegen den
Slice-Plan, die Baseline-Regeln (`v6.13.0` ·
`regelwerk/modul-05-planning-harness.md`, `regelwerk/modul-06-roadmap.md`,
`regelwerk/modul-10-review-harness.md`) und eigene Maschinen-Läufe
(Mutations-Proben gegen Fixtures, Gate-Läufe gegen den echten Bestand).

**Gegenstand:** Die fünf Commits `HEAD~5..HEAD` (`f806ade`, `18884ff`,
`f916634`, `84f3ec3`, `d988fac` — ausschließlich docs(planning)) **plus** der
Verdrahtungs-Commit `08b7754`, der als `HEAD~5` **außerhalb** des Auftrags-Range
liegt und dessen Substanz (`.d-check.yml`, `tools/dcheck-phrase-selftest.sh`)
die Claims des Range trägt. Die Commit-Grenzen sind gegen `git log --oneline -7`
geprüft — Details als F-12 (INFO). Arbeitsverzeichnis `d988fac`, clean.

**Skill:** `.harness/skills/reviewer.md` · **Modell:** glm-5.3-flash ·
**Datum:** 2026-09-29

**Eingangs-Kontext:**

- `docs/plan/planning/in-progress/slice-202-dod-haekchen-lebenslage.md`
  (§1, §2 DoD, §3, §6, §7, §8)
- `BEO-GATE/attestierung-vor-dem-vorgang` (observation.md, state.md,
  evidence/slice-169.md, evidence/slice-197.md, evidence/slice-200.md)
- `AGENTS.md` §3/§5/§6 · `harness/README.md` §Sensors · `harness/conventions.md`
- `Makefile` (Aggregat-Listen), `d-check.mk` (Pin), `.d-check.yml` (Modul
  `structure`), `tools/dcheck-phrase-selftest.sh`
- `v6.13.0` · `regelwerk/modul-05-planning-harness.md` §Trigger je
  Lifecycle-Übergang, §Offene Risiken werden bei Closure aufgelöst
- Image `ghcr.io/pt9912/d-check@sha256:b4b8756b40d3dcd2670a3f83526cb5e5d727d1a850571f73be31edba248abb40`
  (Pin aus `d-check.mk`, lokal vorhanden)

**Eigene Läufe (gemessen, 2026-09-29):**

| Lauf | Exit | Beobachtung |
|---|---|---|
| `make doc-check` | 0 | „609 Datei(en) geprüft, 0 Befund(e)“ — Modul `structure` ist im hermetischen Default-Lauf inert (`.d-check.yml:236–239`) |
| `make doc-structure` | 2 | **2 Befunde** `section-missing` (Details F-2) |
| `make verify` | 2 | `verify-risiko-ausgaenge: FAIL` (Details F-5) **plus** die beiden `doc-structure`-Befunde; `[verify] FAIL` |
| `make dcheck-phrase-selftest` | 0 | Erfolgs-Zeile „4 Werkzeug-Kontrollen (2 Muster × 2 Richtungen)“ (Details F-4) |
| 6 Mutations-Proben (Fixtures, `mktemp`, `git init`, gepinnter Digest) | — | Details F-2, F-6, F-5 |

---

## Findings

| ID | Kategorie | Befund | Quelle | Pfad | Verifizierbar | Klasse |
|---|---|---|---|---|---|---|
| F-1 | HIGH | DoD-Punkt 2 „`make gates` und `make verify` grün“ trägt `[x]`, aber `make verify` ist bei `d988fac` **rot** (Exit 2: `verify-risiko-ausgaenge` FAIL plus zwei `doc-structure`-Befunde). Der grün-Claim war bei jedem Commit des Range unerreichbar, weil der Verdrahtungs-Commit selbst die rote Bedingung einführte — der Slice, der das Attestieren vor dem Vorgang mechanisieren soll, attestiert in der eigenen DoD einen nicht gemessenen Stand. | `BEO-GATE/attestierung-vor-dem-vorgang` (dieselbe Finding-Klasse); `v6.13.0` · `regelwerk/modul-10-review-harness.md` §Klassifikation | `docs/plan/planning/in-progress/slice-202-dod-haekchen-lebenslage.md:51` | ja — `make verify` (Exit 2, gemessen) | Attestierung vor dem Vorgang |
| F-2 | HIGH | Die Verdrahtung ist gegen den **echten Bestand** rot: `make doc-structure` meldet `section-missing` (a) für die `next/`-Regel — `next/` trägt nur die README, die Kandidatenmenge ist leer, das Modul fällt fail-closed („Regel trifft keine Datei“) — und (b) für `open/slice-045-intern-extern-dateimenge.md:1`, das keine DoD-Sektion trägt („DoD (bei Ausarbeitung)“ steht in Prosa, Zeile 241). Beide Zustände sind normale Lifecycle-Zustände, die die Regel (7) nicht unterscheiden kann; gemessen an den Proben p7–p9 (Verzeichnis-Varianten) feuert `section-missing` in allen leeren Besetzungen. Der Selbsttest erhält für Muster 3 **keine** Korpus-Kontrolle (die Korpus-Hälfte des Skripts deckt nur die `reviews`-Phrase) — genau die Kontrolle, die die Röte gegen den Bestand gezeigt hätte. | `v6.13.0` · `regelwerk/modul-06-roadmap.md` §Das Beobachtungs-Register (Beleg-Form: „war sie rot, und woran?“) | `.d-check.yml:419–426`, `tools/dcheck-phrase-selftest.sh:150–207` | ja — `make doc-structure` (Exit 2, gemessen); Proben p7–p9 | Verdrahtung ohne Korpus-Lauf |
| F-3 | HIGH | §1 und DoD-Punkt 1 behaupten, das `forbid-pattern` sei „im `gates`-Aggregat“ wirksam. Das enforcement-Target `doc-structure` hängt nachweislich **nicht** in `gates` (`Makefile:207` führt es nicht; `Makefile:191` begründet den Ausschluss ausdrücklich), sondern in `verify` (`Makefile:196`); im hermetischen Default-`doc-check` ist das Modul inert. Im `gates`-Aggregat erreicht Bedingung 7 nur der Selbsttest — gegen synthetische Fixtures, nie gegen den Bestand. | `v6.13.0` · `regelwerk/modul-10-review-harness.md` §Klassifikation (nachweislich falsche Tatsachenbehauptung) | `docs/plan/planning/in-progress/slice-202-dod-haekchen-lebenslage.md:27` und `:44` | ja — `Makefile:207`, `Makefile:193–196`, `make doc-check` (0 Befunde) | Aggregat-Bindung falsch deklariert |
| F-4 | HIGH | Die Erfolgs-Zeile des Selbsttest lautet gemessen „4 Werkzeug-Kontrollen (2 Muster × 2 Richtungen)“ — der Lauf führt aber **7** Kontrollen über **3** Muster aus (2 + 2 + 3). Zusätzlich beschreiben die Vertrags-Texte den erweiterten Sensor nicht: `harness/sensors/dcheck-phrase-selftest.md:10` („4 — zwei Muster × Positiv/Negativ“) und `:13` („Die Werkzeug-Seite prüft **zwei** Muster“), `harness/README.md:81` („vier Werkzeug-Kontrollen“) sowie der Skript-Kopf (`tools/dcheck-phrase-selftest.sh:16–19` „ZWEI KONTROLLEN JE MUSTER“, `:36–38` „Geprueft sind die zwei“) sind seit der Erweiterung veraltet. | `v6.13.0` · `regelwerk/modul-10-review-harness.md` §Mess-Regeln (Zähler zweimal verschieden) | `tools/dcheck-phrase-selftest.sh:260` | ja — `make dcheck-phrase-selftest` (Exit 0, Zeile gemessen); Zweit-Zähler: 7 Assert-Blöcke im Skript | Zählung deckt ihre Menge nicht |
| F-5 | MEDIUM | §6 notiert als Ausgang „bei Closure (Fixture mit zitiertem Muster in den Selbsttest, wie SL-004 es verlangt)“ — ein Ausgang, der keiner der geschlossenen Dreier-Menge entspricht: `verify-risiko-ausgaenge` fällt genau darüber (gemessen: „Risiko ohne Ausgang aus der geschlossenen Dreier-Menge“), während §7 „jedes mit genau einem Ausgang“ behauptet. Und die versprochene Fixture existiert nicht — der Selbsttest prüft keine Zitat-Kontexte; die Behauptung „empirisch gemessen“ (DoD-Zeile 47–48, Skript-Kommentar `:157–158`) ist im Repo nicht reproduzierbar. Die Substanz stimmt nachweislich — eigene Proben p4–p6: `[x]` in Code-Block und Inline-Code in `next/`- und `open/`-Fixtures bleibt grün —, getragen wird sie im Repo aber von nichts. | `v6.13.0` · `regelwerk/modul-05-planning-harness.md` §Offene Risiken werden bei Closure aufgelöst | `docs/plan/planning/in-progress/slice-202-dod-haekchen-lebenslage.md:81–84`, `:119` | ja — `make verify` (FAIL-Zeile gemessen); `grep` im Selbsttest; Proben p4–p6 | Beleg nicht reproduzierbar |
| F-6 | MEDIUM | Die Scoping-Kontrolle wird als „`in-progress/` bleibt grün“ attested; gemessen ist der Lauf aber **rot** (Exit 1) — die `open/`-Geschwister-Regel der Fixture meldet `section-missing` (Proben p2, p3). Geprüft ist allein die Abwesenheit von `section-forbidden`, nicht ein grüner Lauf; der Geltungsbereich der Messung wird nicht genannt. | `v6.13.0` · `regelwerk/modul-10-review-harness.md` §Mess-Regeln (Geltungsbereich einer Messung) | `docs/plan/planning/in-progress/slice-202-dod-haekchen-lebenslage.md:47` | ja — Proben p2/p3 (rc=1, `section-missing` in der Ausgabe) | Messung ohne Geltungsbereich |
| F-7 | MEDIUM | Die Muster-3-Fixture hält die Muster als **Kopie** (`tools/dcheck-phrase-selftest.sh:164–173`), während der Skript-Kopf die Kopie-Falle selbst benennt: „eine Kopie neben dem Original bleibt gruen, nachdem das Original gebrochen wurde“ (`:27–30`) — und die Korpus-Seite `done-dir` deswegen aus `.d-check.yml` liest (`:229`). Eine Änderung an Bedingung 7 in `.d-check.yml` lässt den Werkzeug-Teil von Muster 3 grün weiterlaufen; Muster 1/2 tragen dieselbe Struktur, aber der Slice hat sie auf ein drittes Muster ausgedehnt, ohne sie zu erwähnen oder abzufedern. | `tools/dcheck-phrase-selftest.sh:27–30` (eigene Grenze des Skripts) | `tools/dcheck-phrase-selftest.sh:164–173` | nein — Struktur-Befund; kein Gate | Konfigurations-Kopie neben dem Original |
| F-8 | MEDIUM | Die Verdrahtung (`08b7754`, 09:35:53) wurde committet, **bevor** der Claim kam (`f806ade`, 09:36:45) und bevor der `git mv` nach `in-progress/` folgte (`18884ff`) — Arbeit vor dem Claim, gegen die Lifecycle-Ordnung (`next→in-progress` vor der Arbeit). Die Closure-Notiz legt den Verstoß offen dar (`:104–108`) — die Darstellung ist mit der Commit-Historie deckungsgleich, der Verstoß selbst bleibt ein solcher. | `v6.13.0` · `regelwerk/modul-05-planning-harness.md` §Trigger je Lifecycle-Übergang und WIP-Limit | `08b7754` vs. `f806ade`/`18884ff` (Commit-Zeiten) | ja — `git log --format='%h %ci %s'` | Arbeit vor dem Claim |
| F-9 | LOW | Die Commit-Message von `18884ff` behauptet „Ruhe-Marker gezogen“, enthält aber nur den `git mv` und die `state.md`-Änderung; der Marker-Zug liegt im Folgesommit `f916634`. | `v6.13.0` · `regelwerk/modul-10-review-harness.md` §Klassifikation | Commit `18884ff` (Message vs. Diff) | ja — `git show 18884ff --stat` | Commit-Message nennt fremde Änderung |
| F-10 | LOW | §7 behauptet für die Anker-Paarung, der Steering-Loop-Eintrag liege in `.d-check.yml` „mit `seit slice-202`“; der Zielort trägt an der Stelle nur die Klammerform „(slice-202)“ (`.d-check.yml:413`), nicht die Anker-Form, die andere Einträge dort tragen („verkörpert seit slice-160“, `.d-check.yml:43`). | `v6.13.0` · `regelwerk/modul-06-roadmap.md` §Wellen-Closure-Prozedur (Anker-Paarung) | `docs/plan/planning/in-progress/slice-202-dod-haekchen-lebenslage.md:121–122` | ja — grep gegen `.d-check.yml` | Anker-Form am Zielort abweichend |
| F-11 | LOW | `d988fac` entfernt aus `state.md` die gültige Verlinkung auf `.d-check.yml` (relativer Pfad, korrekt aufgelöst) mit der Begründung „zitiert Kennung statt Adresse“. Die zitierte Zitier-Form (`AGENTS.md` §5) gilt für einfrierende Artefakte und nennt lebende Dokumente, die den adoptierten Stand nennen, ausdrücklich als „nicht betroffen — dort ist der Link richtig“; `state.md` ist ein lebendes Dokument und Bedingung 7 ist verdrahtet. | `AGENTS.md` §5 (Zitier-Form, Ausnahme-Satz) | `docs/plan/planning/observations/BEO-GATE/attestierung-vor-dem-vorgang/state.md:1` | ja — `git show d988fac` | Zitier-Form auf lebendes Dokument angewandt |
| F-12 | INFO | Der Auftrag-Range `HEAD~5..HEAD` enthält die Verdrahtung nicht: `HEAD~5` **ist** `08b7754`. Der Range trägt nur Lifecycle- und Closure-Commits; die technischen Claims des Range (DoD, Closure-Notiz) verweisen auf den ausgeschlossenen Commit. Dieser Report prüft beide. | `v6.13.0` · `regelwerk/modul-08-agentenrollen.md` §Die neun Übergaben (Implementer→Reviewer trägt den Stand) | `git log --oneline -7` | ja — Commit-Graph | Geprüfter Range ohne Substanz-Commit |

---

## Negativbefunde

- **Positiv-Probe:** eine `[x]`-Zeile im DoD einer `open/`-Fixture meldet gemessen `section-forbidden` mit „verbotenes Muster trifft: `^\s*- \[x\]`“ (Probe p1) — die im DoD behauptete rote Richtung existiert mit genau der benannten Meldung; der Claim „die Mutations-Probe … war rot“ (§7) ist wahr.
- **Zitat-Kontexte:** eigene Proben bestätigen die strukturelle Behauptung — `[x]` in einem fenced Code-Block (p4, p6) und in Inline-Code (p5) feuert das `forbid-pattern` **nicht**; SL-004 („neuer Doku-Sensor meldet sein eigenes Umfeld“) ist nicht eingetreten. Der Befund betrifft nur den fehlenden Beleg im Repo (F-5), nicht die Substanz.
- **Beobachtungs-Register:** die 3×-Aussage ist derivativ belegt — `evidence/` trägt genau drei Dateien (`slice-169.md`, `slice-197.md`, `slice-200.md`); `observation.md` ist seit der slice-169-Ära nicht angefasst worden (unveränderlich ab Anlage); `state.md` trägt Zustand und Beleg ohne Chronik.
- **Roadmap/Ruhe-Marker:** einmal gezogen (`f916634`), Zustand konsistent mit der `in-progress/`-Besetzung; kein Widerspruch zwischen Marker und Verzeichnis.
- **Commit-Scope/Traceability:** alle sechs Commits des Vorgangs sind planning-only und nennen `slice-202`, `AC-QA-02`, `AC-QA-03` — `commit-scope-check`- und `trace-check`-Sicht ohne Befund.
- **Reihenfolge-Darstellung:** die Closure-Notiz beschreibt die Verdrahtung-vor-Claim-Reihenfolge wahrheitsgetreu (mit der Commit-Historie deckungsgleich, F-8); keine still geschönte Chronik.

## Kategorie-Summary

**4 HIGH, 4 MEDIUM, 3 LOW, 1 INFO.** Die HIGH-Findings konzentrieren sich auf
denselben Sachverhalt aus drei Richtungen: der verdrahtete Sensor ist gegen den
echten Bestand rot (F-2), deshalb war die DoD-Attestierung „gates und verify
grün“ bei jeder Commit-Stufe des Range unerreichbar (F-1), und die deklarierte
Aggregat-Bindung stimmt nicht (F-3) — im `gates`-Aggregat sieht kein Target die
neue Bedingung. Der Sensor ist in seiner Positiv-Wirkung gemessen und die
Zitat-Kontext-Substanz stimmt; was fehlt, ist der Bestands-Lauf, die
Korpus-Kontrolle und die korrigierten Zähl- und Vertrags-Texte.

## Verdikt

**Nicht übernahmefähig in dieser Fassung.** Blockierend sind F-1 bis F-3: ein
Slice, dessen Lerneintrag die Attestierungs-Klasse mechanisch fangen will,
schließt (vor der Closure) mit einem DoD-Häkchen über einen roten `make verify`
und einer falschen Aggregat-Bindung — die DoD-Punkte 1 und 2 sind in der
geltenden Fassung nicht erfüllt, und `verify-risiko-ausgaenge` verweigert dem
§6-Risiko den Ausgang. Die Positiv-Wirkung des Sensors selbst (Probe p1 gegen
den gepinnten Digest) ist belegt; nach Korrektur der Scoping-/Bestands-Fragen
(F-2), der Zähl- und Vertrags-Texte (F-4) und des §6-Ausgangs (F-5) ist der
Slice in einem Folgelauf prüfbar.

---

# Nachprüf-Lauf — 2026-09-29, Stand `f6620cb`

**Gegenstand:** die drei Nachzugs-Commits `6da0723` (build/harness),
`af1695e` (docs/harness), `f6620cb` (docs/planning) gegen die Findings des
Erstlaufs. **Eigene Läufe (gemessen):** `make verify` **Exit 0** („Verifikations-Schicht
gruen“, 610 Dateien, 0 Befunde) · `make gates` **Exit 0** (volles Aggregat,
`record-gates` schreibt Nachweis) · Selbsttest Exit 0 · drei Fixtures-Proben
mit der aktuellen Skript-Logik (Extraktion aus `.d-check.yml`, Scoping- und
Zitat-Fixture) · Worktree-Messung des Erstlauf-Stands bei `84f3ec3`.

## Status je Finding

| Finding | Status | Beleg |
|---|---|---|
| F-1 (DoD attestiert grün) | **erledigt** | `make verify` Exit 0 und `make gates` Exit 0, beide gemessen bei `f6620cb`; die DoD-Attestierung ist jetzt durch eigene Messung gedeckt |
| F-2 (Sensor rot gegen Bestand) | **erledigt** | `next/`-Bedingung entfernt; die Grenze (leere Kandidatenmenge läuft fail-closed rot, Wiedereinzug bei Belegung) ist als `GRENZE:` im Kommentar deklariert; `open/slice-045` per `exempt-paths` grandfathered, mit Begründungs-Kommentar; `doc-structure` grün (im `verify`-Lauf enthalten) |
| F-3 (gates-Aggregat-Claim) | **offen (HIGH)** | §1 (Zeile 27) und DoD-Punkt 1 (Zeile 44) tragen weiterhin „im `gates`-Aggregat“ — eine Korrektur auf das `verify`-Aggregat ist im HEAD nicht enthalten; die Wirk-Angabe bleibt nachweislich falsch |
| F-4 (Zählung) | **offen (HIGH)** | Erfolgs-Zeile korrigiert auf „7 Werkzeug-Kontrollen (3 Muster)“ — gemessen laufen aber **8** Kontrollen (2 + 2 + 4; acht `run_dcheck`-Aufrufe, darunter die neue Zitat-Kontrolle); `harness/README.md:81` und die Sensor-Doku übernehmen die falsche Sieben; die Skript-Kopf-Kommentare („ZWEI KONTROLLEN JE MUSTER“, „Geprueft sind die zwei“, `:16–19`/`:36–38`) sind unverändert |
| F-5 (§6-Ausgang, Zitat-Fixture) | **teilweise** | Ausgang trägt jetzt die Dreier-Vokabel (*entfallen/gestrichen mit Begründung*), `verify-risiko-ausgaenge` grün; die neue vierte Kontrolle ist aber wirkungslos konstruiert — siehe N-1 |
| F-6 („bleibt grün“) | **offen (MEDIUM)** | die Scoping-Kontrolle läuft weiterhin rot (Probe a: rc=1, `section-missing` der `open/`-Regel in der `in-progress/`-Fixture); geprüft bleibt nur die Abwesenheit von `section-forbidden`; DoD-Zeile 47 unverändert |
| F-7 (Muster-Kopie) | **erledigt** | das Muster wird fail-closed aus `.d-check.yml` extrahiert (Extraktion gemessen: `^\s*- \[x\]`; leerer Auszug → Exit 2) |
| F-10 (Anker-Form) | **erledigt** | der `.d-check.yml`-Kommentar trägt „(seit slice-202)“ |
| F-11 (state.md-Link) | **erledigt — mit Selbstkorrektur** | der wiederhergestellte Link hat sechs Ebenen und löst auf. **Korrektur zu meinem Erstlauf:** die entfernte Fassung (fünf Ebenen) war **nicht gültig** — die Worktree-Messung bei `84f3ec3` meldet `target-missing` („Linkziel existiert nicht“); mein F-11 irrte hier. `d988fac` reparierte damit einen roten `doc-check`, nennt in der Message aber nur die Zitier-Form-Begründung |
| F-8, F-9, F-12 | unverändert | Historie bzw. Range-Notiz; keine Aktion |

## Neue Findings

| ID | Kategorie | Befund | Quelle | Pfad | Verifizierbar | Klasse |
|---|---|---|---|---|---|---|
| N-1 | MEDIUM | Die neue vierte Kontrolle (Zitat-Kontext) kann per Konstruktion **nicht** feuern: die `\n`-Sequenzen im Fixture-Argument werden im Heredoc nicht interpretiert — die Fixture-Datei enthält eine einzige Zeile mit literalen `\n` (mit `cat -A` gemessen), es entsteht weder ein fenced Code-Block noch eine `[x]`-Zeile am Zeilenanfang, und der Lauf meldet 0 Befunde, gleichgültig wie `d-check` Code-Blöcke behandelt. Die §6-Ausgangs-Begründung („gemessen … der erste Lauf der Fixture meldete 0 Befunde“) stützt sich damit auf einen Lauf ohne Prüf-Gegenstand; das eigentliche SL-004-Szenario (echter Code-Block, Probe c) bleibt zwar grün, wird aber von der Kontrolle nicht getragen. | `BEO-GATE/probe-liefert-den-gegenstand-mit` | `tools/dcheck-phrase-selftest.sh:216–219` | ja — Fixture-Replikation (`cat -A`) und Lauf rc=0; Gegenprobe (c) mit echtem Code-Block ebenfalls rc=0 | Probe ohne Gegenstand |
| N-2 | LOW | Die `GRENZE:`-Deklaration in Bedingung 7 nennt nur `next/`; die `open/`-Regel erbt dasselbe fail-closed-Verhalten, sobald `open/` leer wird oder nur `slice-045` trägt (nach Abzug von `exempt-paths` ist die Kandidatenmenge dann leer — gemessen an Probe p8: fehlendes Verzeichnis ⇒ `section-missing`). | `v6.13.0` · `regelwerk/modul-06-roadmap.md` §Das Beobachtungs-Register | `.d-check.yml:419–429` | ja — Probe p8 | Grenze halbseitig deklariert |

## Nachlauf-Summary und Verdikt

**Erledigt:** F-1, F-2, F-7, F-10, F-11. **Offen:** F-3 (HIGH — falsche
Aggregat-Angabe in §1/DoD unverändert), F-4 (HIGH — die korrigierte Zählung 7
ist erneut falsch, gemessen 8; Skript-Kopf unangetastet), F-6 (MEDIUM), N-1
(MEDIUM), N-2 (LOW). **Unverändert:** F-8, F-9, F-12.

**Verdikt:** die tragenden Blocker des Erstlaufs sind getragen — `make gates`
und `make verify` sind grün, der Sensor läuft gegen den Bestand, die
Dokumentation beschreibt den erweiterten Sensor. Blockierend für die Closure
bleiben **F-3** (die Wirk-Angabe „gates-Aggregat“ steht noch in §1 und DoD) und
**F-4** (die Zählung ist in der zweiten Korrektur wieder falsch — 8 gemessen, 7
behauptet, in Erfolgs-Zeile, Sensor-Doku und Gate-Index). N-1 ist vor der
Closure mitzunehmen: eine Kontrolle, die nicht fehl schlagen kann, trägt die
§6-Ausgangs-Begründung nicht.

---

# Verifikations-Lauf 2 — 2026-09-29, Stand `bf68069`

**Gegenstand:** die drei Nachzugs-Commits `fb128dd` (build/harness),
`0855afd` (docs/planning), `bf68069` (docs/harness; nimmt zugleich die
Erstfassung dieses Reports in den Bestand). **Eigene Läufe (gemessen):**
`make gates` **Exit 0** · `make verify` **Exit 0** · `make doc-structure`
Exit 0 (610 Dateien, 0 Befunde) · Selbsttest Exit 0 („8 Werkzeug-Kontrollen
(3 Muster)“) · acht Kontroll-Matrix-Proben gegen den gepinnten Digest
(`m1–m6`, `d1`, `d2`, unten).

## Status je Finding

| Finding | Status | Beleg |
|---|---|---|
| F-1 (DoD attestiert grün) | **erledigt (bleibt)** | `make gates` Exit 0, `make verify` Exit 0, beide gemessen bei `bf68069` |
| F-2 (Sensor rot gegen Bestand) | **erledigt (bleibt)** | `make doc-structure` Exit 0, 0 Befunde; GRENZE-Kommentar jetzt auch für den `open/`-Fall (Rest: F-N-2-Wording) |
| F-3 (gates-Aggregat-Claim) | **erledigt** | §1 (Zeile 27) und DoD-Punkt 1 (Zeile 44) nennen jetzt das `verify`-Aggregat — deckungsgleich mit `Makefile:196` |
| F-4 (Zählung) | **erledigt, Rest LOW (F-N-4)** | Erfolgs-Zeile „8 Werkzeug-Kontrollen (3 Muster)“ stimmt mit den acht `run_dcheck`-Aufrufen überein; Sensor-Doku-Tabelle „8 — drei Muster“, `harness/README.md:81` „acht Werkzeug-Kontrollen (drei Muster)“, Skript-Kopf „MEHRERE KONTROLLEN JE MUSTER“ nachgezogen; Rest: die Prosa-Sätze in der Sensor-Doku (Zeile 13 f.) sagen weiterhin „zwei Muster“ |
| F-5 / N-1 (Zitat-Fixture) | **erledigt** | die Fixture wird jetzt aus einer Body-Datei mit echten Newlines gebaut (`tools/dcheck-phrase-selftest.sh:216–224`) — der Code-Block mit `[x]` am Zeilenanfang ist real; Lauf rc=0, 0 Befunde gemessen; die Sensitivität ist strukturell über die Positiv-Kontrolle verankert (dieselbe Zeile außerhalb des Blocks feuert, Probe p1/m1) |
| F-6 („bleibt grün“) | **offen (MEDIUM)** | die Scoping-Fixture erzeugt weiterhin nur `in-progress/`; die `open/`-Regel matcht keine Datei → gemessen rc=1 mit `section-missing` (Probe a). Die im Nachzug angekündigte Bestands-Spiegelung („`open/` mit offenem Slice belegt“) ist im Skript nicht enthalten; DoD-Zeile 47 unverändert |
| F-7 / F-10 / F-11 | **erledigt (bleibt)** | unverändert auf `bf68069` nachgemessen (Extraktion, Anker-Form, Linktiefe) |
| F-8, F-9, F-12 | unverändert | historisch bzw. Range-Notiz |

## Neues Finding

| ID | Kategorie | Befund | Quelle | Pfad | Verifizierbar | Klasse |
|---|---|---|---|---|---|---|
| N-3 | HIGH | Das `forbid-pattern` feuert nur, wenn die `[x]`-Zeile die **erste Inhaltszeile** der DoD-Sektion ist — gemessen an acht Varianten: allein (m1), zwei `[x]`-Zeilen (m6) und `[x]`-Zeile mit nachfolgender Prosa (d1) feuern; `[x]` nach Prosa (m3), nach Listitem (m4), nach unchecked `- [ ]` (m5) und im Code-Block (m2) bleiben grün. Damit fängt der Sensor die beobachtete Gestalt („alle Häkchen gesetzt“, slice-200-Klasse), aber **nicht** die gemischte DoD-Gestalt — ein Slice in `open/` mit einem einzigen unchecked Punkt vor dem attestierenden Häkchen läuft durch; genau diese Gestalt trägt der Plan des slice-202 selbst. §1/DoD und Lerneintrag behaupten den Fang ungequalifiziert („Ein DoD-Häkchen … ist mechanisch gefangen“). | `v6.13.0` · `regelwerk/modul-10-review-harness.md` §Klassifikation; `AC-QA-02` (ehrliche Heuristik-Grenze — auszuweisen, nicht als Vollständigkeit auszugeben) | `.d-check.yml:421–426`, `docs/plan/planning/in-progress/slice-202-dod-haekchen-lebenslage.md:25–31` | ja — Kontroll-Matrix m1–m6/d1/d2 gegen den gepinnten Digest; Positiv/Negativ-Wechsel allein über die Vor-Zeile | Geltungsbereich enger als die Claim |

Zusatz-Befund zur Einordnung (kein eigenes Finding): die Zitat-Kontrolle kann
dieses Verhalten nicht testen — ein `[x]` **im** Code-Block steht stets hinter
der Fence-Zeile, ist also unter der „erste Inhaltszeile“-Semantik nie ein
Feuerfall; die Kontrolle bleibt als Regressions-Anker gültig, ihr Szenario ist
aber vom selben Verhalten überdeckt.

## Rest-Befunde (LOW)

- **F-N-4 (LOW):** die Prosa-Sätze der Sensor-Doku (Zeile 13 f.: „Die
  Werkzeug-Seite prüft **zwei** Muster …“) widersprechen der korrigierten
  Tabelle im selben Dokument („8 — drei Muster“) — Rest aus F-4.
- **F-N-2-Wording (LOW):** die neue GRENZE-Formulierung „dann entfällt diese
  Bedingung mit“ beschreibt nicht das gemessene Verhalten — bei leerem
  `open/` (oder nur `slice-045` nach Abzug von `exempt-paths`) läuft die
  Bedingung fail-closed **rot** (`section-missing`, Probe c: rc=1, „kein
  Abschnitt passt auf den Selektor“).

## Schluss-Summary und Verdikt

**Erledigt (gemessen):** F-1, F-2, F-3, F-4, F-5/N-1, F-7, F-10, F-11.
**Offen:** N-3 (HIGH), F-6 (MEDIUM), F-8 (MEDIUM, historisch-disclosed),
F-N-4/F-N-2-Wording/F-9 (LOW), F-12 (INFO).

**Verdikt:** die Form-Bedenken des Erst- und Nachprüf-Laufs sind abgetragen —
gates und verify sind grün, die Aggregat-Angabe stimmt, die Zählung stimmt,
die Zitat-Kontrolle prüft ihr Szenario. Blockierend bleibt **N-3**: der
Geltungsbereich des Sensors ist enger als die Claim in §1, DoD und
Lerneintrag — er fängt die All-haken-Gestalt, nicht jede vorgezogene
Attestierung. Die ehrliche Heuristik-Grenze (`AC-QA-02`) verlangt, die Grenze
auszuweisen: die Claim in Plan und Closure-Notiz auf die gemessene Gestalt
einschränken (und die gemischte Gestalt als benannte Grenze führen),
oder — als Modul-Erweiterung im d-check — die Sektion-zentrierte Auswertung
zu schärfen. F-6 ist vor der Closure zu klären (Fixture-Spiegelung oder
DoD-Formulierung).
