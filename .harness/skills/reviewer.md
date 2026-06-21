# Reviewer-Skill — a-check

- **Status:** Accepted
- **Gilt für:** Plan-/Design-/Code-Review der Doku- und (ab slice-003)
  Code-Artefakte dieses Repos.
- **Bezug:** [`AGENTS.md`](../../AGENTS.md) §3 (Hard Rules) + §5 (Traceability);
  Regelwerk v1.3.0 Modul 10. Baseline: [`harness/conventions.md`](../../harness/conventions.md) §Baseline.

Repo-spezifisches „worauf achtest du", damit ein Reviewer-Agent zwischen
Sessions nicht driftet (Regelwerk Modul 10). Diese Datei wird versioniert,
nicht überschrieben (ADR-Hard-Rule, Modul 4).

## Kontext-Eingang (Pflicht)

Bevor der Reviewer den Gegenstand liest:

- der Review-Gegenstand (Diff, ADR-Entwurf, Slice-Plan)
- [`spec/lastenheft.md`](../../spec/lastenheft.md) für referenzierte `AC-*`-IDs
- ADRs aus [`docs/plan/adr/`](../../docs/plan/adr/), deren ID im Gegenstand
  vorkommt — nur aktive (`Proposed`/`Accepted`), nie `Superseded`
- [`AGENTS.md`](../../AGENTS.md) §3 Hard Rules
- frühere Findings am selben Bereich ([`docs/reviews/`](../../docs/reviews/))

## Klassifikation (für dieses Repo)

**HIGH** — blockiert:
- Verstoß gegen eine Hard Rule ([`AGENTS.md`](../../AGENTS.md) §3.1–§3.6)
- Harness-Lüge: behauptetes Gate ohne Make-Target, erfundene ID, stille Setzung
- Spec-Stratum referenziert abwärts (ADR/Slice) — Referenz-Richtung verletzt
- nachweislich falsche Tatsachenbehauptung (gegen ein Repo-Artefakt verifiziert)

**MEDIUM** — vor Merge/Acceptance klären:
- unbelegte Tatsachenbehauptung (nicht gegen ein Repo-Artefakt belegbar)
- `Bezug:`/`Schärft:` unvollständig oder unpassend zur Entscheidung
- fehlende wesentliche Konsequenz/Risiko in einer ADR; fehlender oder vager
  Fitness-Function-Anker

**LOW** — nice-to-fix: Wording, Bezug-Feinheiten, fehlender Querverweis,
Provenance-Platzierung außerhalb der Historie-Zone.

**INFO** — Hinweis ohne erwartete Aktion (Verweis auf zuständige Rolle oder
Folge-Slice).

## Was dieser Skill NICHT macht

- Keine Lösungsvorschläge — Reviewer kategorisiert, Implementer entscheidet.
- Keine Verifikation gegen DoD (Verifier, Modul 11), keine Validation
  (Validator).
- Kein Schreibzugriff auf den Review-Gegenstand.

## Output-Schema

Pro Finding: `kategorie` · `quelle` (AC-/ADR-ID, Hard-Rule, Konvention) ·
`pfad` (Datei:Zeile) · `befund` (1–2 Sätze, beobachtbar, **ohne
Lösungsvorschlag**) · `verifizierbar` (ja/nein — gäbe es einen Gate-/Tool-Lauf,
der es bestätigt?).

Zusätzlich: pro geprüftem Bereich eine **Negativbefund-Zeile** („geprüft, ohne
Befund"), eine **Kategorie-Summary** und ein **Verdikt**. **HIGH-Findings
werden vor Übernahme adversarisch gegen das Repo-Artefakt verifiziert**
(Modul 11). Kontext-Trennung: wer ein Artefakt verfasst hat, reviewt es nicht
im selben Kontextfenster (Modul 8). Report-Ablage: ein Report pro Lauf unter
[`docs/reviews/`](../../docs/reviews/), Folgeläufe als neue Datei.

## Pflege (Steering-Loop)

Bei dreimaligem gleichem Finding: Klassifikation schärfen → Folge-ADR oder
[`AGENTS.md`](../../AGENTS.md)-Eintrag → Fitness Function (Modul 13).
