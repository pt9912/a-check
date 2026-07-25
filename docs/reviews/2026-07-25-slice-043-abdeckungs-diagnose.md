# Review-Synthese — slice-043: Abdeckungs-Diagnose

**Datum:** 2026-07-25 · **Gegenstand:** [slice-043](../plan/planning/in-progress/slice-043-schicht-abdeckung-sichtbar.md)
([ADR-0029](../plan/adr/0029-abdeckungs-diagnose-advisory.md), Spezifikation 0.26.0) ·
**Art:** Plan-Review (Regelwerk Modul 10) **vor** der Umsetzung, danach Code-Review über
`main...HEAD`.

> **Form.** Strukturierter Ein-Reviewer-Mehrwinkel-Durchgang; Reviewer und Verifier dieselbe
> Instanz — die Rollentrennung aus Modul 8/11 ist nicht erfüllt. Die DoD-/Spec-Verifikation steht
> getrennt in [slice-043](../plan/planning/in-progress/slice-043-schicht-abdeckung-sichtbar.md).
> Kategorien nach Regelwerk `v1.3.0` Modul 10.

## Kontext-Eingang

Entwurf slice-043 (Stand 2026-07-24) · [ADR-0029](../plan/adr/0029-abdeckungs-diagnose-advisory.md)
· [AC-QA-02](../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)/[AC-QA-01](../../spec/lastenheft.md#ac-qa-01--determinismus)/[AC-FA-CLI-001](../../spec/lastenheft.md#ac-fa-cli-001--aufruf-scan-wurzel-und-exit-codes)
· [SPEC-CLI-001](../../spec/spezifikation.md#spec-cli-001--aufruf-scan-wurzel-und-exit-codes)
· [`AGENTS.md`](../../AGENTS.md) §3.6 · Konsumenten-Konfigurationen (sieben, lokal).

## Findings — Plan-Review (vor der Umsetzung)

| # | Kategorie | Quelle | Pfad | Befund | verifizierbar |
|---|---|---|---|---|---|
| P-1 | **HIGH** | Evidenz des Entwurfs (§2) | `slice-043` §2, `slice-042` §8.1 | Der Entwurf nennt zwei betroffene Konsumenten und begründet damit „flotten-weite Evidenz". Für einen davon (d-migrate) trifft das nicht zu: dessen `.a-check.yml` enthält seit ihrem **ersten** Commit einen `exclude`-Block, der exakt den behaupteten Baum ausnimmt. Real ungedeckt: **ein** Konsument, **zwei** Dateien | **ja** — Zählung mit exakter Glob-Semantik über sieben Konfigurationen |
| P-2 | **MEDIUM** | Entwurf §3 (Granularität) | `slice-043` §3 | Die Wahl „Zähler + Zonen" statt Datei-Liste ist mit „hunderte Dateien" begründet — eine Menge, die nach P-1 nicht existiert. Der einzige noch offene Entscheid (Zonen-Bildung) hängt an dieser Begründung | ja |
| P-3 | **MEDIUM** | [`AGENTS.md`](../../AGENTS.md) §3.6 (sinngemäß in der Gegenrichtung) | Entwurf §3 (Gestalt (d)) | Gestalt (d) führt einen Config-Schlüssel mit Exit-Code-Wirkung ein. Der Aufwand ist überwiegend Vertrags-Maschinerie (neue AC-ID, fail-closed-Validierung, Exit-Code-Semantik) — gemessen 2–2,5× der Gestalt (a) — für eine Strenge ohne angefragten Bedarf | ja (Aufwandszählung §2b) |
| P-4 | **LOW** | eigene Messmethode | — | Die erste Nachmessung näherte `languages`-Globs über Datei-Endungen und ignorierte deren Pfad-Anker; sie meldete für b-cad 41 ungedeckte Dateien, die a-check nie scannt. Erst die exakte Portierung von `globToRegexp` lieferte belastbare Zahlen | ja |

## Findings — Code-Review (nach der Umsetzung)

| # | Kategorie | Quelle | Pfad | Befund | verifizierbar |
|---|---|---|---|---|---|
| C-1 | **INFO** | [AC-QA-01](../../spec/lastenheft.md#ac-qa-01--determinismus) | `internal/cli/cli.go` (`writeCoverageNotice`) | Die Kürzung ab zehn Dateien ist eine Mengen-Beschränkung der Ausgabe. Sie nennt die Restzahl, ist also nicht still — geprüft gegen die Regelwerks-Linie „keine stillen Kappungen" | ja — Test mit 13 Dateien |
| C-2 | **INFO** | Modul 10 §Rollentrennung | Prozess | Reviewer und Verifier sind dieselbe Instanz; die Findings P-1…P-4 entstanden beim eigenen Nachmessen, nicht durch eine unabhängige Linse. Der Befund ist die **Form**, nicht der Inhalt | nein |

## Auflösung

| # | Auflösung |
|---|---|
| P-1 | Evidenz in [slice-043 §2a](../plan/planning/in-progress/slice-043-schicht-abdeckung-sichtbar.md) korrigiert; die falsche Zeile auch in [slice-042 §8.1/§8.2](../plan/planning/done/slice-042-constructs-aufruf-monopol.md) berichtigt. Der Slice wurde **nicht** verworfen: die Fehlerklasse bleibt real (ein Verzeichnis, das der Konsument in einer `tech`-Regel als Architektur-Zone führt, hat keine Schicht) — nur ihre Verbreitung ist klein |
| P-2 | Granularität auf **Datei-Liste** entschieden; die Zonen-Bildung entfällt ersatzlos (§0) |
| P-3 | Gestalt **(a)** abgenommen, `strict_coverage` vertagt mit Trigger ([ADR-0029](../plan/adr/0029-abdeckungs-diagnose-advisory.md) Entscheidung 2) |
| P-4 | Messmethode ersetzt (exakte Glob-Portierung); die verworfene Näherung ist in §2a als Fehlerquelle benannt |
| C-1 | Kein Delta — Verhalten ist gewollt und getestet |
| C-2 | Kein Delta möglich im Ein-Instanz-Setup; als Form-Einschränkung ausgewiesen |

## Negativbefunde (geprüft, ohne Befund)

- `internal/hexagon/core` (`UncoveredFiles`) — Quell-Seite, `composition_root`-Ausnahme, Sortierung; keine Berührung der Regel-Auswertung, keine neuen Befund-Typen.
- `layerLabel` — nur `wrong-direction` kann eine leere Quell-Schicht sehen (jede andere Regel setzt eine Rolle und damit eine Schicht voraus); geprüft, ohne Befund.
- `internal/adapter/driven/{config,extract,graph,report}` — unberührt; kein Config-Schlüssel, kein Port-Wechsel, keine `ARC-*`-Berührung.
- Exit-Code-Pfad — die Diagnose wird **nach** `Report` geschrieben und verändert den Rückgabewert nicht; Test pinnt Exit 0 bei ungedeckten Dateien.
- Distributions-Akzeptanz — `image-test` vergleicht stderr byte-identisch nativ↔Container; Fixture vollständig gedeckt, keine Wechselwirkung.
- Konsumenten-Probe — sechs von sieben Konfigurationen diagnose-frei, m-trace nennt genau seine zwei Dateien; alle Exit-Codes unverändert.

## Nachlauf

`strict_coverage` bleibt als Folge-Slice offen (Trigger in
[ADR-0029](../plan/adr/0029-abdeckungs-diagnose-advisory.md)). P-1 ist der dritte Fall dieser
Session, in dem eine **Behauptung über Konsumenten-Configs** einer Nachmessung nicht standhielt
(nach der Port→Port-Zählung und der `53`-Fehlzählung) — nach Modul 10 §Pflege ist das die Schwelle,
ab der die Ursache adressiert gehört: Konsumenten-Evidenz gehört gemessen, mit exakter
Glob-Semantik, nicht aus der Config gelesen.
