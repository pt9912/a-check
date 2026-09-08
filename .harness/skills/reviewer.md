# Reviewer-Skill — a-check

- **Status:** Accepted
- **Gilt für:** Plan-/Design-/Code-Review der Doku- und (ab slice-003)
  Code-Artefakte dieses Repos.
- **Bezug:** [`AGENTS.md`](../../AGENTS.md) §3 (Hard Rules) + §5 (Traceability);
  Regelwerk **v6.5.0** Modul 10 (vendored: [`.harness/baseline/v6.5.0/regelwerk/modul-10-review-harness.md`](../baseline/v6.5.0/regelwerk/modul-10-review-harness.md)). Baseline: [`harness/conventions.md`](../../harness/conventions.md) §Baseline.

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
- **Kommentar trägt keine der Kommentar-Klassen** — er beschreibt die verworfene
  Alternative („ohne X wäre …"), einen abwesenden Text („früher stand hier …")
  oder bricht mitten im Satz ab, weil eine Teilersetzung den Rest stehen ließ.
  Gilt für Code, Konfiguration, Skripte
  ([`AGENTS.md`](../../AGENTS.md) §3.7). **Kein Gate fängt das** — die Regel ist
  inferentiell und hängt am Review.
  *Im Bestand belegt:* [`BEO-HARNESS/hard-rule-37-ohne-sensor`](../../docs/plan/planning/observations/BEO-HARNESS/hard-rule-37-ohne-sensor/observation.md)
  (slice-108, slice-169).
- **Zustandsfeld trägt Chronik** — eine `Stand`-/`Status`-Zelle (Roadmap,
  Beobachtungs-Register, Meilenstein) erzählt, *wie* der Zustand entstand,
  statt Zustand und Beleg als auflösbaren Anker zu nennen; oder ein Drift-Log
  protokolliert Schließungen und erreichte Meilensteine. **Kein Gate fängt
  das.** Zwei Ausprägungen, beide teuer geworden: ein `state.md`, das nach der
  Behebung den behobenen Zustand weiter behauptet, und eines, dessen
  `Stand:`-Zeile bei 3× keinen der drei Ausgänge trägt.
  *Im Bestand belegt:* [`BEO-HARNESS/chronik-in-gelesenen-dateien`](../../docs/plan/planning/observations/BEO-HARNESS/chronik-in-gelesenen-dateien/observation.md)
  (slice-103, slice-182).
- **Norm nur im Template-Kommentar** — eine Regel steht im `<!-- -->`-Block
  einer vendorten Ziel-Form und nirgends sonst. Sie ist weg, sobald jemand die
  Vorlage kopiert und die Kommentare löscht — und **a-check kopiert sie**
  ([`AGENTS.md`](../../AGENTS.md) §5, *„Beim Kopieren anzupassen"*; eigene
  Vorlagen-Dateien führt das Repo nicht).
  *Besetzter Fall:* `v6.5.0` · `templates/docs/reviews/review-report.template.md`
  trägt im vierten Kommentarblock die Norm *„die Klassen-Bezeichnung muss über
  Läufe hinweg stabil sein"* — sie steuert den Steering-Loop-Zähler und steht
  **nur dort**.
  **Zweite Ausprägung, hiesig:** [`.d-check.yml`](../../.d-check.yml) trägt die
  Begründung jeder Regel ausschließlich in Kommentaren. Das ist zulässig — aber
  eine *Zusage* darf dort nicht allein stehen; sie gehört in
  [`AGENTS.md`](../../AGENTS.md) §4 oder eine Sensor-Datei, die der Lauf liest.
  **Kein Gate fängt beides.**

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
der es bestätigt?) · `klasse` (die Finding-Klasse in einem Halbsatz).

**`klasse` speist den Steering-Loop-Zähler**, und darum gilt für sie die Norm
aus der Ziel-Form des Reports: **die Bezeichnung muss über Läufe hinweg stabil
sein.** Leiten zwei Läufe dieselbe Klasse unterschiedlich ab, zählt das
Register sie getrennt, und keine erreicht je 3×. Alle bestehenden Reports
führen das Feld; im Skill fehlte es bis slice-185.

**Zitier-Form** (Norm, `v6.5.0` · `templates/docs/reviews/review-report.template.md`):
Der Report friert ein; was er zitiert, bewegt sich weiter. Deshalb **Kennung,
nicht Adresse** — `slice-NNN` statt seines Lifecycle-Pfads, `make <target>`
statt eines Links auf die Sensor-Datei, eine Baseline-Stelle als **Tag + Pfad
in Inline-Code** statt als Link: `` `v<X.Y.Z>` · `regelwerk/<datei>.md` §<Abschnitt> ``.
Der vendored Baum trägt genau einen Tag; der nächste Sprung löscht den alten,
und ein Link darauf färbt einen Report rot, den niemand mehr anfassen darf.
Das `pfad`-Feld auf den **geprüften Gegenstand** ist davon nicht betroffen — es
hält den Stand des Laufs fest und darf das.

Zusätzlich: pro geprüftem Bereich eine **Negativbefund-Zeile** („geprüft, ohne
Befund"), eine **Kategorie-Summary** und ein **Verdikt**. **HIGH-Findings
werden vor Übernahme adversarisch gegen das Repo-Artefakt verifiziert**
(Modul 11). Kontext-Trennung: wer ein Artefakt verfasst hat, reviewt es nicht
im selben Kontextfenster (Modul 8). Report-Ablage: ein Report pro Lauf unter
[`docs/reviews/`](../../docs/reviews/), Folgeläufe als neue Datei.

## Pflege (Steering-Loop)

Bei dreimaligem gleichem Finding: Klassifikation schärfen → Folge-ADR oder
[`AGENTS.md`](../../AGENTS.md)-Eintrag → Fitness Function (Modul 13).
