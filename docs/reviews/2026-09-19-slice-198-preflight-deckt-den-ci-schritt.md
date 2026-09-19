# Review-Report: slice-198 — 2026-09-19

**Review-Art:** unabhängiger Lauf — Code-/Doku-Review gegen **Plan, ADR und
Konventionen** (Modul 10 §Drei Review-Arten). Nicht gegen DoD/Spec; das ist die
Verifikation und eine andere Rolle.

**Gegenstand:** Commits `23b3e9d` (Makefile-Target `preflight` + Index-Zeile),
`760ab51` (`docs/user/releasing.md`, Item 5), `ccf4e92` (Closure + Register),
`06c7876` (GATES-Liste des Command-Guard) sowie der Slice-Plan in seinem
in-progress-Stand.

**Skill:** `.harness/skills/reviewer.md` @ `60e7b66` (jüngster Commit an der
Datei) ·
**Modell:** unbekannt (Subagent) · **Datum:** 2026-09-19

**Eingangs-Kontext:**

- `slice-198` (Slice-Plan, in-progress-Stand)
- `AGENTS.md` §3 (Hard Rules), §4 (Quality Gates), §5, §6
- `harness/README.md` §Sensors (Gate-Index), `harness/conventions.md`
  §Modus-Deklaration (u. a. `MR-024`)
- `Makefile:213-229`, `.github/workflows/ci.yml`, `.github/workflows/release.yml`,
  `tools/ci-commit-range.sh`, `.claude/hooks/pretooluse-command-guard.sh`,
  `d-check.mk` (`doc-immutable`)
- `docs/user/releasing.md` §Freigabe-Checkliste
- Baseline `v6.6.0` · `regelwerk/modul-05-planning-harness.md`,
  `modul-06-roadmap.md`, `modul-08-agentenrollen.md`, `modul-10-review-harness.md`

---

## Findings

### F-1 — Beleg ohne Geltungsbereich: die grüne Probe nennt keine Range, und heute ist sie leer

- `kategorie`: MEDIUM
- `quelle`: `AGENTS.md` §5 *Geltungsbereich einer Messung* (Mess-Regel 1);
  Reviewer-Skill §Mess-Regeln; Baseline `v6.6.0` ·
  `regelwerk/modul-06-roadmap.md` §Das Beobachtungs-Register (Beleg-Form)
- `pfad`: `docs/plan/planning/in-progress/slice-198-preflight-deckt-den-ci-schritt.md:73`
  (DoD-Probe) und `docs/user/releasing.md:117` (Beleg-Slot Item 5)
- `befund`: Die DoD nennt die rote Probe mit ihrem Gegenstand (drei
  `core-drift-vcs`-Meldungen; die Evidence-Datei nennt zusätzlich
  `89fc7dc..ed7a3d8`), die **grüne** nur „gegen die **aktuelle** Range" — ohne
  Wert und ohne Aufruf; die Variable `PREFLIGHT_RANGE`, mit der man eine Range
  vorgeben kann, steht nur im `Makefile`. Am 2026-09-19 ist `origin/main` ==
  `HEAD`, derselbe Aufruf fährt damit `06c7876..HEAD` und
  `commit-scope-check` meldet wörtlich „**0** (planning)-Commit(s) in
  `06c7876e…HEAD` geprueft": der Beleg ist grün, ohne einen Commit anzusehen.
- `verifizierbar`: ja — `make preflight` (Exit 0) und die Range-Zeile im Lauf
  belegen es; beide gefahren (siehe §Gate-Läufe 1).
- `klasse`: Beleg ohne Geltungsbereich — grüne Probe ohne Range-Wert

### F-2 — „genau das, was der nächste Push prüfen wird": die CI-Weiche kennt drei Fälle

- `kategorie`: MEDIUM
- `quelle`: `AGENTS.md` §5 (Mess-Regel 1); `tools/ci-commit-range.sh`
  (§DIE DREI LAGEN, selbstgetestet über `make ci-range-selftest`)
- `pfad`: `docs/plan/planning/in-progress/slice-198-preflight-deckt-den-ci-schritt.md:39-40`,
  `Makefile:219` (Kommentar), `…/BEO-GATE/preflight-deckt-den-ci-schritt-nicht/state.md:1`
- `befund`: Die Gleichsetzung „lokal ist sie `origin/main..HEAD`" — „und das ist
  **genau** das, was der nächste Push prüfen wird" — steht dreimal im
  Gegenstand. `compute_range` liefert `origin/main..HEAD` nur in zwei der vier
  Selbsttest-Fälle (neuer Branch, unerreichbare Basis/Force-Push); ein
  **normaler Push auf einen bestehenden Branch** prüft `before..HEAD`. Die
  Abweichung geht in die sichere Richtung (lokal ist die Menge dann größer),
  die Aussage „genau" trägt aber nicht.
- `verifizierbar`: ja — `make ci-range-selftest` (Exit 0) plus Lesen von
  `compute_range` in `tools/ci-commit-range.sh`.
- `klasse`: Mengen-Identität behauptet, wo die Weiche drei Fälle kennt

### F-3 — Der Gate-Index nennt seine Mengen an zwei Stellen unvollständig

- `kategorie`: MEDIUM
- `quelle`: `AGENTS.md` §4 (der Index ist die Autoritäts-Doku, es gibt genau
  eine); Baseline `v6.6.0` · `regelwerk/grundlagen-harness-dateien.md`
  §harness/README.md als Einstiegspunkt
- `pfad`: `harness/README.md:55` und `harness/README.md:119-122`
- `befund`: Der Satz „die PR-/Push-CI zieht `make ci` + `make trace-check` auf
  jede Integration" nennt einen Teil der zusätzlichen Schritte, während die
  neue Zeile 37 Zeilen darunter drei nennt (`trace-check`,
  `commit-scope-check`, `doc-immutable`) — zwei Aussagen über dieselbe Menge im
  selben Abschnitt. Und die Aufzählung „Nicht hier, obwohl sie in keinem
  Aggregat hängen" führt fünf Gates, ohne `make preflight` zu nennen, das
  ebenfalls in keinem Aggregat hängt und dessen Zeile darum weder
  `— (Aggregat)` noch „nicht in `gates`" trägt (gemessen: `grep -c preflight`
  über die Aufzählung = 0).
- `verifizierbar`: ja — Lesen beider Stellen gegen `ci.yml`/`.github/workflows/`;
  kein Gate deckt Sätze.
- `klasse`: Mengen-Aussage im Gate-Index nicht mitgezogen

### F-4 — Rang-Zeiger ohne Fundstelle (`AGENTS.md` §6 Schritt 6)

- `kategorie`: LOW
- `quelle`: `AGENTS.md` §3.7 (ein Kommentar trägt genau eine der fünf Klassen —
  ein Rang-Zeiger zeigt auf den Ort der Regel); Hard Rule §3.7
- `pfad`: `Makefile:220` (Kommentar) und `Makefile:223` (Hilfezeile)
- `befund`: Beide Zeilen verweisen auf `AGENTS §6 Schritt 6`; dort steht
  „Repo-weiten Gate-Lauf vor Handoff (`make gates`)". `AGENTS.md` nennt `make
  preflight` an keiner Stelle (gemessen: `grep -c preflight AGENTS.md` = 0).
  Wer dem Zeiger folgt, findet den Gegenstand dort nicht.
- `verifizierbar`: ja — der `grep` oben.
- `klasse`: Rang-Zeiger ohne Fundstelle

### F-5 — Geführter Zähler neben der abgeleiteten Belegliste

- `kategorie`: LOW
- `quelle`: Baseline `v6.6.0` · `regelwerk/modul-06-roadmap.md`
  §Das Beobachtungs-Register („es gibt kein Feld, in das man ihn schreibt");
  `docs/plan/planning/observations/README.md` (wörtlich dieselbe Zusage)
- `pfad`: `docs/plan/planning/observations/BEO-GATE/preflight-deckt-den-ci-schritt-nicht/state.md:1`
- `befund`: „**Stand:** offen (1×)" schreibt den Zähler in ein Feld, obwohl er
  abgeleitet ist (Zahl der `evidence/`-Dateien). Die Klasse ist vorbestehend
  breit im Bestand: **21** von **68** `state.md` führen einen Multiplikator.
  Der neue Eintrag setzt sie fort.
- `verifizierbar`: ja — die beiden `grep`-Zählungen; kein Sensor.
- `klasse`: Geführter Zähler neben der abgeleiteten Belegliste

### F-6 — Sub-Area-Liste und ihre Begründung fallen auseinander

- `kategorie`: LOW
- `quelle`: Baseline `v6.6.0` · `regelwerk/modul-05-planning-harness.md`
  §Zwei Schritte vor der Modus-Begründung (jede berührte Sub-Area muss das
  Inklusionskriterium erfüllen)
- `pfad`: `docs/plan/planning/in-progress/slice-198-preflight-deckt-den-ci-schritt.md:152-153`
  (und `:156`)
- `befund`: §9 führt als berührt `GATE` und `HARNESS`, begründet `HARNESS` aber
  mit „`releasing.md` liegt unter `docs/user/` — das ist die Sub-Area `USER`".
  Damit fehlt `USER` — die über `docs/user/releasing.md` tatsächlich berührte
  Sub-Area — in der Liste, während die Zeile darunter das Register „über
  `GATE`, `HARNESS`, `USER`" liest und `HARNESS` (berührt über
  `harness/README.md`) ohne eigenen Grund dasteht.
- `verifizierbar`: nein — Formfrage, kein Gate; `make verify` prüft die
  Modus-Blöcke nicht auf ihre Sub-Area-Zuordnung.
- `klasse`: Sub-Area-Zuordnung und Begründung fallen auseinander

### F-7 — Geltungsbereich-Aussage über `release.yml` weiter als ihr Gegenstand

- `kategorie`: LOW
- `quelle`: `AGENTS.md` §5 (Mess-Regel 1)
- `pfad`: `docs/plan/planning/in-progress/slice-198-preflight-deckt-den-ci-schritt.md:50-51`
- `befund`: „Der `release`-Workflow fährt nur `make ci`" ist für
  make-Targets gemessen richtig (`grep` über die `run:`-Zeilen findet allein
  `make ci VERSION="$VERSION"`), für den Workflow als Ganzes nicht: er führt
  neun weitere Schritte, darunter „Verify OCI labels" — Item **6** der
  Freigabe-Checkliste. Der Satz steht hinter einem Geltungsbereich, der auf
  „die Schritte des `ci`-Workflows und des `ci`-Targets" einschränkt; ohne
  diese Einschränkung im selben Satz liest er sich als Aussage über die
  Release-Pipeline.
- `verifizierbar`: ja — `grep -nE 'run: make' .github/workflows/release.yml`
  und Lesen des Workflows.
- `klasse`: Geltungsbereich-Aussage weiter als der Gegenstand

### F-8 — Der Vorbedingungs-Wächter läuft nach dem teuersten Schritt

- `kategorie`: LOW
- `quelle`: `AGENTS.md` §5/§6 (Pre-Flight vor dem Handoff); Plan §7
  (Risiko-Ausgang 2 behauptet „bricht das Target mit Exit 2")
- `pfad`: `Makefile:223-224`
- `befund`: `preflight: ci` — der Range-Wächter ist die erste *Rezept*-Zeile,
  läuft also erst nach `gates` + `image-test`. Fehlt `origin/main`, wird der
  volle Gate-/Image-Lauf (Läufe dieses Reports: mehrere Minuten Docker-Build)
  bezahlt und **danach** mit Exit 2 abgebrochen. Zudem ist `2` der allgemeine
  GNU-make-Fehler-Exit eines jeden roten Rezepts (im roten Lauf unten ebenfalls
  `Fehler 2`) — unterscheidbar ist die Meldung, nicht der Code.
- `verifizierbar`: ja — Rezept-Reihenfolge; der Wächter isoliert gefahren
  (§Gate-Läufe 8).
- `klasse`: Vorbedingungs-Wächter hinter dem teuersten Schritt

## Negativbefunde

- geprüft, ohne Befund: das Rezept selbst (`Makefile:223-229`) — `preflight`
  fährt `ci` und ruft alle drei Range-Schritte als Sub-`make` mit
  `RANGE=$(PREFLIGHT_RANGE)` auf; die Sub-Aufrufe erreichen ihre Ziele
  (`trace-check` → `--range`, `commit-scope-check` → `RANGE=`, `doc-immutable` →
  `--range`), `.PHONY` ist gesetzt.
- geprüft, ohne Befund: **die Mengen-Behauptung** — `ci.yml` fährt zusätzlich
  **genau** die drei genannten Schritte (Schritt „Traceability +
  ADR-Immutabilität": `make trace-check`, `make commit-scope-check`,
  `make doc-immutable`, je mit `RANGE="$RANGE"` aus `tools/ci-commit-range.sh`;
  danach `make ci`). Kein weiterer make-Schritt im Workflow; `release.yml`
  fährt genau ein make-Target (`ci`, siehe F-7).
- geprüft, ohne Befund: **die rote Probe** — `89fc7dc..ed7a3d8` liefert über
  `make preflight` Exit 2 mit drei `core-drift-vcs`-Meldungen an `ADR-0017`,
  `ADR-0018`, `ADR-0038`, wörtlich wie im Plan und in `MR-024` deklariert.
- geprüft, ohne Befund: **die drei Folge-Einträge** — jeder einzeln als nötig
  belegt (Mutations-Proben in §Gate-Läufe 4, 6, 7: je rot ohne den Eintrag, grün
  mit ihm). Der Guard-Eintrag steht in der richtigen Liste: `preflight`
  **urteilt** (sein Exit-Code trägt einen Befund), gehört damit nach `GATES` und
  nicht nach `NICHT_PRUEFEND`.
- geprüft, ohne Befund: **Einsortierung im Gate-Index** — die Zeile steht in der
  Tabelle *Sensors*, nicht unter *Nicht-Gates*; nach dem Kriterium des README
  (*urteilen* gegen *bewegen · messen · sagen*) trägt die Klassifikation.
- geprüft, ohne Befund: **`§3.7`** auf den neuen Texten — der `Makefile`-Block
  trägt Abgrenzung („eigenes Target, keine Erweiterung von `ci`"), Zusage und
  Grenze und schreibt an den Ändernden; die Index-Zeile trägt Zusage + Grenze;
  die Freigabe-Checkliste einen Beleg-Slot. Keine Chronik, kein beschriebener
  abwesender Zustand, kein Satz über die verworfene Alternative als Zustand.
- geprüft, ohne Befund: **Referenz-Richtung / Source Precedence** — kein
  Spec-Stratum berührt; `docs/user/releasing.md` verweist aufwärts (ADR,
  `harness/conventions.md`), `harness/README.md` (Rang 9) auf `AC-QA-02`
  (Rang 1); keine abwärtsgerichtete Referenz.
- geprüft, ohne Befund: **Register-Form** (außer F-5) — drei Dateien,
  `evidence/slice-198.md` trägt die Kennung eines abgeschlossenen Vorgangs,
  Sub-Area `GATE` ist in `harness/conventions.md` deklariert, der Pfad ist
  nicht neu-erfunden.
- geprüft, ohne Befund: **Kennungs-Linkpflicht** der berührten Dateien —
  `make doc-check` über den Arbeitsbaum, grün im Lauf `make preflight`
  (§Gate-Läufe 1); `MR-024` löst im Konventions-Block auf, die
  `AC-QA-02`-Anker in Index-Zeile und Plan lösen auf.

## Gate-Läufe

Alle Aufrufe mit `> /tmp/<datei> 2>&1; echo "exit=$?"` (nie in eine Pipe, nie
im selben Aufruf wie ein Commit).

| # | Aufruf | Exit | Geltungsbereich |
|---|---|---|---|
| 1 | `make preflight` (Default-Range `06c7876..HEAD`) | **0** | `ci` (`gates` + `image-test`) über den Arbeitsbaum **plus** die drei Range-Schritte über eine **leere** Range (0 Commits) — grün, aber ohne einen Commit zu prüfen (F-1). **Kein** Nachweis der grünen DoD-Probe des Plans |
| 2 | `PREFLIGHT_RANGE=89fc7dc..ed7a3d8 make preflight` | **2** | reproduziert die rote Probe: `make trace-check` und `make commit-scope-check` grün (21 planning-Commits), `make doc-immutable` rot mit drei `core-drift-vcs`-Meldungen (`ADR-0017`, `ADR-0018`, `ADR-0038`) |
| 3 | `make verify` | **0** | Verifikations-Schicht (DoD-/Closure-Fragen), 21 Anforderungen, 0 Waisen |
| 4 | `make doc-targets` | **0** | Gate-Index ↔ Makefile, beide Richtungen. Mutation: Index-Zeile in einer Repo-Kopie unter `/tmp` entfernt → **Exit 1**, `Makefile:223 preflight gate-undocumented` |
| 5 | `make doc-structure` | **0** | Struktur-Invarianten (u. a. Gate-Tabellen-Zellen, Closure-Form) |
| 6 | `bash <Kopie des Guard>.sh --selftest` mit `preflight` aus der `GATES`-Liste entfernt | **1** | „Pruef-Target `preflight` fehlt in der GATES-Liste (Regel 2 greift dort nicht)" — die Gegenrichtung (`guard-selftest ok`) steht in Lauf 1 |
| 7 | `check_phony_complete` (Funktionen aus `tools/gate-consistency.sh`) gegen eine `Makefile`-Kopie ohne den `.PHONY`-Eintrag | **1** | „Target `preflight` fehlt in .PHONY … laesst make das Rezept ueberspringen und Exit 0 melden"; gegen das Original **Exit 0** |
| 8 | Range-Wächter isoliert (die `PREFLIGHT_RANGE`- und die `case`-Zeile wörtlich in eine Kopie ohne `origin/main`) | **2** | „origin/main nicht aufloesbar — erst git fetch origin fahren." |

## Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 0 |
| MEDIUM | 3 |
| LOW | 5 |
| INFO | 0 |

**Finding-Klassen dieses Laufs:** Beleg ohne Geltungsbereich (grüne Probe ohne
Range-Wert) · Mengen-Identität behauptet, wo die Weiche drei Fälle kennt ·
Mengen-Aussage im Gate-Index nicht mitgezogen · Rang-Zeiger ohne Fundstelle ·
Geführter Zähler neben der abgeleiteten Belegliste · Sub-Area-Zuordnung und
Begründung fallen auseinander · Geltungsbereich-Aussage weiter als der
Gegenstand · Vorbedingungs-Wächter hinter dem teuersten Schritt

## Verdikt

**Merge-blockierend:** ja — kein HIGH, aber drei MEDIUM, und alle drei sind
Präzisierungen an Artefakten **dieses** Slice: die Probe in §4/§6 des Plans (samt
dem Beleg-Slot in `docs/user/releasing.md:117`), die Range-Gleichsetzung an drei
Stellen, und zwei Mengen-Aussagen im Gate-Index, den der Slice selbst erweitert.
Kein Befund berührt die Substanz des Targets: `make preflight` fährt die
richtige Menge, trifft den Gegenstand (rote Probe reproduziert) und ist an allen
drei Stellen seines Bestands belegt.

**Übergabe:** Findings gehen an den Implementer (Rückkante Review → Plan, wo
§1/§2/§4/§9 betroffen sind); die **Finding-Klassen** gehen in die Closure §7 und
von dort in den Zähler — die ersten drei gehören zu den in `AGENTS.md` §5
gebündelten Mess-Regeln, die der Reviewer-Skill unter §Mess-Regeln führt, nicht
zu einer neuen Klasse. Der Report ersetzt keine Verifikation (Modul 11).
