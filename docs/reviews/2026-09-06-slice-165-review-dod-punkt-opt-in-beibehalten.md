# Review-Report: slice-165 — 2026-09-06

**Review-Art:** Plan — geprüft gegen den Slice-Plan selbst (Modul 10 §Drei
Review-Arten): der Gegenstand ist eine empirische Analyse plus eine
Architektur-Entscheidung (Opt-in vs. verpflichtend), kein Verhaltenscode.

**Gegenstand:** Commit `930f6bf` ("build(harness): slice-165 schneiden --
Review-Checkbox-Punkt bleibt Opt-in, MR-019"), Endstand nach `ef4f7e8`
("fehlende Reports für slice-163/164 nachgetragen, slice-165-Fehler
korrigiert")

**Nachtrag (2026-09-06):** dieser Report wurde nachträglich unter
`docs/reviews/` abgelegt — dieselbe Lücke, die dieser Slice selbst
diagnostiziert (§4: ein Review fand statt, wurde aber nicht als Datei
persistiert), traf ironischerweise den Slice, der sie beschreibt. Erst
`archive-wave`s Dry-Run für `welle-14` deckte das Fehlen dieses Reports
auf.

**Skill:** `.harness/skills/reviewer.md` @ Stand `ef4f7e8` (unverändert
seit Anlage) · <!-- d-check:ignore -->
**Modell:** claude-sonnet-5 · **Datum:** 2026-09-06

**Eingangs-Kontext:**

- `docs/plan/planning/done/slice-165-review-dod-punkt-opt-in-beibehalten.md`
- `docs/plan/planning/done/slice-163-adaptions-durchgang-v610.md`
- `docs/plan/planning/done/slice-164-regelwerk-v620-delta-analyse.md`
- `AGENTS.md` §5, §6
- `.d-check.yml` (`structure`-Modul, `tasks-ignore-pattern`)
- `harness/conventions/MR-019-review-dod-opt-in.md`
- isoliertes Scratch-Repo gegen den gepinnten `d-check`-Digest

---

## Findings

### F-1 — Entscheidungsbegründung zitierte eine bereits überholte `AGENTS.md`-Fassung

- `kategorie`: HIGH
- `quelle`: `git log -p --follow` gegen `AGENTS.md` §6, Commit `2488705` (slice-162, selber Tag)
- `pfad`: `docs/plan/planning/done/slice-165-review-dod-punkt-opt-in-beibehalten.md` §5 (erste
  Fassung)
- `befund`: §5 begründete die Opt-in-Entscheidung mit dem Zitat „`AGENTS.md` §6 verlangt Review
  bereits präzise ‚bei jedem Slice mit Code- oder Vertragsänderung' — nicht bei jedem Slice."
  Diese Formulierung existierte in `AGENTS.md` §6, wurde aber durch `slice-162` (selber Tag)
  ersatzlos gestrichen und durch den `v6.1.0`-Baseline-Wortlaut ersetzt, der keine Einschränkung
  auf Code-/Vertragsänderungen mehr trägt. Das Zitat stammte aus dem Gesprächskontext, nicht aus
  einer `git log`/Datei-Prüfung — dieselbe Fehlerklasse, die
  [`BEO-PLAN/slice-provenienz-aus-gedaechtnis-statt-git-log`](../plan/planning/observations/BEO-PLAN/slice-provenienz-aus-gedaechtnis-statt-git-log/observation.md)
  bereits für Provenienz-Behauptungen registriert.
- `verifizierbar`: ja — `git log -p --follow -- AGENTS.md` zeigt die Streichung in `2488705`.
- `klasse`: Aktueller Dateiinhalt aus Gedächtnis zitiert statt Datei gelesen

### F-2 — Zwei Review-Reports fehlten trotz real durchgeführter Reviews

- `kategorie`: HIGH
- `quelle`: Verzeichnislisting `docs/reviews/` vor diesem Slice
- `pfad`: `docs/plan/planning/done/slice-163-adaptions-durchgang-v610.md`,
  `docs/plan/planning/done/slice-164-regelwerk-v620-delta-analyse.md`
- `befund`: Beide Slices führten in ihrer DoD die Formulierung „Unabhängiges Plan-Review über
  getrennten Kontext durchgeführt" — real durchgeführt über getrennte Subagenten-Kontexte, mit
  echten Findings —, aber ohne persistierten Report unter `docs/reviews/`. `make doc-reviews`
  hatte das nicht gefangen, weil die Formulierung die Trigger-Phrase „unabhängiger Review" nicht
  trifft (s. §3/§4 dieses Slice) — eine leere Kandidatenmenge meldet grün, ohne etwas geprüft zu
  haben.
- `verifizierbar`: ja — `ls docs/reviews/` vor der Korrektur zeigte keine Dateien für diese
  beiden Slices.
- `klasse`: Prüfer ohne Gegenstand (dritte Instanz desselben Registereintrags)

## Negativbefunde

- geprüft, ohne Befund: die empirische Tabelle in §3 (welche DoD-Formulierung `make doc-reviews`
  auslöst) wurde in einem eigenen, isolierten Scratch-Repo gegen denselben gepinnten
  `d-check`-Digest nachgestellt — alle fünf Zeilen reproduzierbar.
- geprüft, ohne Befund: `MR-019`s Pflichtfelder vollständig und formkonsistent zu
  `MR-017`/`MR-018`.
- geprüft, ohne Befund: `tasks-ignore-pattern` in `.d-check.yml` enthält nach der Änderung
  tatsächlich `Unabhängiger Review` als Alternative.
- geprüft, ohne Befund: DoD-Größe (3 echte Liefer-Punkte nach Bereinigung), Kopf-Metadaten,
  Closure-Notiz-Form strukturell konsistent zu `slice-163`/`slice-164`.

## Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 2 |
| MEDIUM | 0 |
| LOW | 0 |
| INFO | 0 |

**Finding-Klassen dieses Laufs:** Aktueller Dateiinhalt aus Gedächtnis
zitiert statt Datei gelesen · Prüfer ohne Gegenstand (dritte Instanz)

## Verdikt

**Merge-blockierend:** nein mehr — beide Findings wurden vor Abschluss von
`slice-165` korrigiert (Commit `ef4f7e8`): die Begründung in §5 wurde durch
einen tragfähigen, aktuell zutreffenden Verweis ersetzt
(`.d-check.yml`s eigener `reviews`-Modul-Kommentar), und die beiden
fehlenden Reports für `slice-163`/`slice-164` wurden aus den bereits
durchgeführten Reviews rekonstruiert und abgelegt. F-2 erreichte damit die
3×-Schwelle für
[`BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf`](../plan/planning/observations/BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf/observation.md)
(zusammen mit `slice-120`/`slice-123`) — Ausgang: geplant in `slice-168`
(Kalibrierungs-Selbsttest für phrasen-/muster-basierte Prüfer), zugewiesen
beim Lese-Schritt der `welle-14`-Closure.

**Übergabe:** Findings sind bereits in `slice-165`s Closure eingearbeitet.
Dieser Report ist Lauf-Beleg, keine Verifikation — DoD-/Spec-Konformität
prüft der Verifier separat (Modul 11).
