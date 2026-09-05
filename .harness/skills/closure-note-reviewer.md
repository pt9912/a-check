# Closure-Note-Reviewer-Skill — a-check

* Status: Accepted
* Bezug: [`AGENTS.md`](../../AGENTS.md) §5 (Closure-Pflicht), `make verify` /
  `make doc-structure` (Modul `structure`), Regelwerk `modul-11-verification.md` §Schritt 5,
  `modul-15-observability.md` (Doku-Konsistenz)
* Gilt für: den *inferentiellen* Nachlauf zu `make verify` — greift dort, wo Struktur allein die
  Floskel nicht fängt
* Entstanden: slice-050 (Etappe E der `v3.5.2`-Migration), Fund B-4 aus
  [slice-048](../../docs/plan/planning/done/welle-12/slice-048-modul-delta-lesen.md)

> **Abweichung von der Baseline-Vorlage — bewusst.** Die Vorlage bindet an eine
> Closure-Note-ADR (Nummer gilt im Kurs-Repo) und an `tools/check_closure_notes.py`. Beides
> existiert in a-check nicht: dieselbe Nummer trägt hier
> [ADR-0011](../../docs/plan/adr/0011-domain-application-trennung-rolle-app.md) — die
> Domain/Application-Trennung —, und Python-Tooling widerspräche
> [`AGENTS.md`](../../AGENTS.md) §3.1 (Docker/make-only). Die Pflicht ist darum in `AGENTS.md` §5
> verankert, das Struktur-Gate ist `make doc-structure`. Blind kopierte Bezüge wären
> halluzinierte Anker gewesen.

## Kontext-Eingang (Pflicht)

Was der Reviewer *immer* mitbringt, bevor er urteilt:

- alle Closure-Abschnitte der Slices in `docs/plan/planning/done/`
- [`AGENTS.md`](../../AGENTS.md) §5 — welche drei Inhalte zählen und warum die Pflicht existiert
- das Ergebnis von `make verify` für **denselben** Stand: was das Struktur-Gate schon abgedeckt
  hat, wird **nicht** doppelt gemeldet

Ohne diesen Block prüft der Reviewer Text, aber nicht *gegen die Pflicht-Inhalte*.

## Prüf-Auftrag

Lies den Closure-Abschnitt jedes Slice in `done/`. Markiere alle, die **keinen** der folgenden
Inhalte tragen:

- **(a) ein Lernsignal mit Ursache** — „X, *weil* Y", nicht „X war schwierig"
- **(b) ein konkretes Folge-Slice** — mit ID, und die sollte in `open/` oder `done/` auffindbar sein
- **(c) eine beobachtbare Architektur-Aussage** — eine, die ein anderer am Repo nachprüfen kann

Floskeln ohne einen dieser Inhalte sind ein **HIGH**-Finding. Inferentiell, weil „Inhalt vs.
Floskel" semantisch ist; das Struktur-Gate deckt nur Überschrift, Platzhalter, Floskel-Liste und
Satzzahl ab.

## Klassifikation

**HIGH** — Floskel ohne Substanz: syntaktisch vorhanden (überlebt `make verify`), trägt aber
*keinen* der drei Inhalte. Beispiele: „war ganz okay, läuft jetzt", „wie geplant umgesetzt",
„Gate-Beleg beim Merge" als **einziger** Inhalt.

**MEDIUM** — genau *einer* der Inhalte fehlt oder ist unkonkret:
- Lernsignal ohne das „weil X" (Behauptung statt Ursache)
- Folge-Slice benannt, aber ohne auffindbaren Eintrag
- Architektur-Aussage als Etikett statt als nachprüfbare Beobachtung
- **Aussage im Futur**, die inzwischen überholt ist („folgt mit dem nächsten Release") — real
  aufgetreten in slice-041 und erst beim Bau dieses Sensors bemerkt

**LOW** — alle Inhalte da, aber schwer nachvollziehbar formuliert (Substanz da, Klarheit fehlt).

**INFO** — Hinweis ohne erwartete Aktion.

## Was dieser Skill NICHT macht

- Keine Struktur-Prüfung (Überschrift da? ≥ 2 Sätze? Platzhalter?) — das ist `make verify`;
  nicht doppeln.
- Keine Bewertung, ob der Slice *fachlich* korrekt abgeschlossen wurde — das ist die
  Verifier-/Validator-Frage.
- Keine Umschreibung der Notiz — der Autor formuliert nach, der Reviewer kategorisiert nur.
- Keine Prüfung von Slices ausserhalb `done/` — `open/` und `in-progress/` tragen noch keine
  Pflicht.

## Output-Schema

Jedes Finding:

- `kategorie`: HIGH | MEDIUM | LOW | INFO
- `quelle`: `AGENTS.md §5` | `Closure-Inhaltspflicht (a/b/c)`
- `pfad`: `docs/plan/planning/done/<slice>.md`:&lt;Zeile&gt;
- `befund`: *welcher* Inhalt fehlt, 1–2 Sätze, beobachtbar, ohne Formulierungs-Vorschlag
- `verifizierbar`: **nein** — Floskel-Erkennung ist inferentiell; `make verify` bestätigt nur die
  Struktur, nicht den Inhalt

Zusätzlich am Ende je betrachteter Charge eine Zeile „geprüft, ohne Befund: `done/<Charge>`"
(Negativbefund — macht die Abdeckung sichtbar).

## Pflege (Steering-Loop)

Bei **dreimaligem** HIGH derselben Floskel-Art:

- das Muster in [`AGENTS.md`](../../AGENTS.md) §5 als benanntes Anti-Pattern aufnehmen
- prüfen, ob das Modul `structure` die Floskel *strukturell* fangen könnte
  (`forbid-pattern` in [`.d-check.yml`](../../.d-check.yml) erweitern) → dann ins Struktur-Gate heben; ein deterministischer Marker ist
  billiger als inferentielles Nachlesen
- die Slice-Form schärfen, falls der Abschnitt die Pflicht-Inhalte nicht klar genug abfragt
  (Etappe D, Fund B-5)

Diese Skill-Datei wird **nicht** überschrieben, sondern versioniert.
