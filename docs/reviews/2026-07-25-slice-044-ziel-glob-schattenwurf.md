# Review-Synthese — slice-044: Rückzug bei unauflösbarem Ziel-Glob

**Datum:** 2026-07-25 · **Gegenstand:** [slice-044](../plan/planning/done/slice-044-ziel-glob-schattenwurf.md)
([ADR-0028](../plan/adr/0028-ziel-glob-schattenwurf.md), Spezifikation 0.25.0) ·
**Art:** Code-Review (Regelwerk Modul 10) über `origin/main...HEAD` vor dem Merge.

> **Form.** Wie bei [slice-042](2026-07-25-slice-042-constructs-monopol.md): strukturierter
> Ein-Reviewer-Mehrwinkel-Durchgang, Reviewer und Verifier dieselbe Instanz. Die DoD-/Spec-
> Verifikation (Proben-Tabelle) steht getrennt in
> [slice-044 §7](../plan/planning/done/slice-044-ziel-glob-schattenwurf.md) und ist nach Modul 11
> **kein** Review-Gegenstand. Kategorien nach Regelwerk `v1.3.0` (HIGH/MEDIUM/LOW/INFO).

## Kontext-Eingang

Diff `origin/main...HEAD` (4 Commits, später 6) · [ADR-0028](../plan/adr/0028-ziel-glob-schattenwurf.md)
· berührte Vorentscheidungen [ADR-0013](../plan/adr/0013-layerof-laengster-praefix.md),
[ADR-0023](../plan/adr/0023-deklarations-bewusste-mehr-wurzel-aufloesung.md),
[ADR-0026](../plan/adr/0026-hexslice-vertical-slice-regeln.md) ·
[SPEC-RULE-001](../../spec/spezifikation.md#spec-rule-001--regel-auswertung) ·
[`AGENTS.md`](../../AGENTS.md) §3.4 · [`harness/conventions.md`](../../harness/conventions.md)
[MR-005](../../harness/conventions.md#mr-005--referenzmatrix-intra-spec-richtung--adrslice-disziplin-d-check-angleichung).

## Findings

| # | Kategorie | Quelle | Pfad | Befund | verifizierbar |
|---|---|---|---|---|---|
| F-1 | **HIGH** | [AC-QA-02](../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze) (False-Green-Klasse) | `internal/hexagon/core/rules.go` (`shadowedByWildcardGlob`) | Ein Rückzug allein über den literalen **Kopf** des Globs zöge **jeden** Import in den Teilbaum zurück, nicht nur die Port-Ziele — die umschließende Schicht wäre still ungegatet. **Im Entwurf gefunden**, vor dem ersten Commit geschlossen; die Klasse ist schwerer als der Fehlbefund, den der Slice behebt | **ja** — Gegenprobe: `outbound → application` (Nicht-Port-Ziel) muss weiter melden |
| F-2 | **MEDIUM** | [`AGENTS.md`](../../AGENTS.md) §3.4 / [MR-005](../../harness/conventions.md#mr-005--referenzmatrix-intra-spec-richtung--adrslice-disziplin-d-check-angleichung) | `spec/spezifikation.md` (SPEC-RULE-001-Körper) | Der Spec-Body verwies auf eine ADR. Spec-Straten verweisen nie abwärts; `doc-check`-Modul `matrix` meldet `matrix-forbidden`. Der Commit ging trotzdem durch, weil der Gate-Lauf gepipt war (dieselbe Klasse wie slice-042-F-4) | ja — `make doc-check` |
| F-3 | **LOW** | Maintainability | `internal/hexagon/core/rules.go` (`shadowedByWildcardGlob`) | Kopf und Tail-Marker werden unabhängig als Segment-Run gesucht, ohne Reihenfolge-Bedingung. In einem pathologischen Pfad (Modul-Präfix heißt selbst `ports`) zieht der Guard konservativ zurück, obwohl das Glob nicht matchen könnte. Folge ist fail-open, kein Fehlbefund | ja |
| F-4 | **INFO** | [`AGENTS.md`](../../AGENTS.md) §3.3 / Branch-Disziplin | Prozess (kein Pfad) | Die Code-Änderung lag zunächst direkt auf `main` und im selben Commit wie die Spezifikation — die Spec-first-Reihenfolge war im Verlauf nicht ablesbar | nein (Prozess) |

## Auflösung

| # | Auflösung |
|---|---|
| F-1 | **Tail-Marker-Bedingung** ergänzt: der Rückzug greift nur, wenn auch das literale Segment **nach** dem Wildcard (`ports`) im Kandidaten vorkommt. Vier Gegenproben als Tests; zusätzlich per **Diskriminierungs-Probe** belegt (Guard testweise abgeschaltet ⇒ `TestNestedPortImportNoFalsePositive` fällt) |
| F-2 | ADR-Verweis aus dem Spec-Body entfernt (bleibt in der Historie-Zeile, die `matrix` ausnimmt); Commit per `--amend` repariert |
| F-3 | **Nicht geändert.** Die Bedingung ist in [SPEC-RULE-001](../../spec/spezifikation.md#spec-rule-001--regel-auswertung) genau so beschrieben („beide kommen als Segment-Run vor") — Spec und Code sagen dasselbe, es entsteht keine verdeckte Behauptung. Eine Reihenfolge-Bedingung wäre Zusatzkomplexität für einen pathologischen Pfad |
| F-4 | Commit-Historie neu geschnitten: eigener Branch, Spec-Commit ohne Code, Code-Commit separat |

## Negativbefunde (geprüft, ohne Befund)

- `literalHead`/`literalTail` — Randfälle `a/*b`, `**/foo`, `src/*/ports/x`, `…/**/**` durchgespielt; alle fallen konservativ auf „kein Rückzug", Einheitstests gepinnt.
- Determinismus — der Rückzug ist eine reine Existenz-Aussage über die Glob-Menge, reihenfolge-unabhängig; geprüft, ohne Befund.
- Quell-Seite (`LayerOf`) — unberührt; Port-Dateien unter Innen-Wildcard-Globs bleiben voll geprüft.
- `internal/adapter/driven/{config,extract,graph,report}` — unberührt.
- `tools/image-test.sh` Block (4) — vergleicht stdout **und** stderr byte-identisch; Fixture ist vollständig gedeckt, keine Wechselwirkung.
- Konsumenten-Regression (b-cad, d-check, d-migrate, m-trace, belief-agent, HexSlice-Go-Beispiel) — je `exit=0`, 0 Befunde; kein Konsument nutzt Innen-Wildcard-Globs, der Guard ist für den Bestand inert.

## Nachlauf

F-3 bleibt als dokumentierte Konservativität stehen. F-2/F-4 sind Prozess-Findings derselben
Wurzel wie slice-042-F-4 (Gate-Ausgabe nicht prüfbar gemacht, Disziplin-Schritte übersprungen) —
nach Modul 10 §Pflege beim dritten Auftreten ein Fall für eine Hard-Rule- oder Gate-Ergänzung,
nicht für ein weiteres Einzel-Finding.
