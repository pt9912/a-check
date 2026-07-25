# Review-Synthese — slice-042: `constructs`-Roh-Text-Monopol

**Datum:** 2026-07-25 · **Gegenstand:** [slice-042](../plan/planning/done/slice-042-constructs-aufruf-monopol.md)
([ADR-0027](../plan/adr/0027-constructs-roh-text-monopol.md),
[AC-FA-RULE-011](../../spec/lastenheft.md#ac-fa-rule-011--konstrukt-monopol-regel-construct-leak)
0.22.0, Spezifikation 0.24.0) · **Art:** Code-Review (Regelwerk Modul 10) über `main...HEAD`
vor dem Merge.

> **Form — ehrlich ausgewiesen.** Dieses Review lief als **strukturierter
> Ein-Reviewer-Mehrwinkel-Durchgang**, nicht als parallele Multi-Agenten-Linsen wie in den
> Dokumenten bis slice-026. Reviewer und Verifier waren **dieselbe Instanz** — die Rollentrennung
> aus Regelwerk Modul 8/11 ist damit nicht erfüllt; die DoD-/Spec-Verifikation steht getrennt in
> [slice-042 §10](../plan/planning/done/slice-042-constructs-aufruf-monopol.md) und ist **kein**
> Bestandteil dieses Reviews (Modul 11: DoD-Verletzungen sind keine Review-Findings).
>
> **Kategorien-Drift (INFO, harness-relevant).** Die Dokumente bis slice-026 nutzen
> BLOCKER/MINOR/NIT; Regelwerk `v1.3.0` Modul 10 nennt **HIGH/MEDIUM/LOW/INFO**. Dieses Dokument
> folgt dem Regelwerk. Die Abweichung der Altdokumente ist in
> [`harness/conventions.md`](../../harness/conventions.md) **nicht** als `MR-*` deklariert —
> Kandidat für eine Adaptions-Erklärung oder eine Angleichung.

## Kontext-Eingang

Diff `main...HEAD` (6 Commits) · [AC-FA-RULE-011](../../spec/lastenheft.md#ac-fa-rule-011--konstrukt-monopol-regel-construct-leak)/[AC-FA-CONF-001](../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml)
· [ADR-0027](../plan/adr/0027-constructs-roh-text-monopol.md) · [ADR-0015](../plan/adr/0015-regex-tech-muster.md)/[ADR-0018](../plan/adr/0018-exclude-scan-scope.md)/[ADR-0025](../plan/adr/0025-exclude-verzeichnis-prune.md)
(berührte Vorentscheidungen) · [`AGENTS.md`](../../AGENTS.md) Hard Rules ·
[`harness/conventions.md`](../../harness/conventions.md).

## Findings

| # | Kategorie | Quelle | Pfad | Befund | verifizierbar |
|---|---|---|---|---|---|
| F-1 | **MEDIUM** | [AC-FA-RULE-011](../../spec/lastenheft.md#ac-fa-rule-011--konstrukt-monopol-regel-construct-leak) (Boundary-AK) | `spec/lastenheft.md` · `internal/adapter/driven/extract/extract.go` (`prepSource`) | Die Anforderung sagt „Treffer ausschließlich in einem Kommentar ⇒ kein Befund". Für **Python** trifft das nicht zu: `prepSource` lässt Python bewusst ungestrippt, ein Treffer im `#`-Kommentar meldet. Anforderungstext und Verhalten weichen ab | **ja** — Fixture: gleiche Kommentar-Zeile in `.py` meldet, in `.cpp` nicht |
| F-2 | **LOW** | [AC-FA-RULE-003](../../spec/lastenheft.md#ac-fa-rule-003--tech-kapselung-regel-tech-leak) | `internal/hexagon/core/rules.go` (`Tech.matches`) | Ein `tech`-Eintrag mit leerem `pattern` meldet nie (`Pattern != ""` im Matcher) — dieselbe False-Green-Klasse, die 0.14.0 für den leeren `adapter` fail-closed gestellt hat. Bestandsverhalten, nicht durch diesen Diff eingeführt; fiel beim Bau der `constructs`-Validierung auf | ja |
| F-3 | **INFO** | [AC-QA-01](../../spec/lastenheft.md#ac-qa-01--determinismus) / [SPEC-DET-001](../../spec/spezifikation.md#spec-det-001--determinismus-vertrag) | `internal/hexagon/core/rules.go` (`sortFindings`) | Die Sortierschlüssel (Pfad, Zeile, Regel) bilden über einem instabilen `sort.Slice` keine Totalordnung, sobald eine Regel mehrere Befunde je Zeile erzeugt. **Im Entwurf gefunden**, nicht im Review — vor dem ersten Commit geschlossen; hier als Beleg der Kette dokumentiert | ja |
| F-4 | **MEDIUM** | [`AGENTS.md`](../../AGENTS.md) §4 / [`CLAUDE.md`](../../CLAUDE.md) „Kein Erfolg ohne echte Gate-Ausgabe" | Prozess (kein Pfad) | Ein Gate-Lauf wurde als `make … \| tail` ausgeführt und der Commit daran gekettet. Der Exit-Code einer Pipeline ist der des letzten Glieds — ein roter `doc-check` blieb dadurch folgenlos und ein Commit ging mit rotem Gate durch | ja |

## Auflösung

| # | Auflösung |
|---|---|
| F-1 | **Doku korrigiert, nicht Code.** Ein `#`-Strip allein für den Konstrukt-Pfad verschlänge ein `#` im String-Literal und könnte einen echten Treffer verbergen (False-Green wiegt schwerer). Anforderung + [SPEC-EXTRACT-001](../../spec/spezifikation.md#spec-extract-001--import-extraktion) + Handbuch weisen die Grenze jetzt aus; `TestConstructHitsPythonCommentBoundary` nagelt beide Sprachen gegeneinander fest |
| F-2 | **Nicht behoben** — eigener Mikro-CR, notiert in [slice-042 §9.1](../plan/planning/done/slice-042-constructs-aufruf-monopol.md). Die `tech`-Semantik zu verschärfen ist ein Vertragsschnitt außerhalb dieses Slices |
| F-3 | `sortFindings` um die **Meldung** als letzten Schlüssel erweitert; [SPEC-DET-001](../../spec/spezifikation.md#spec-det-001--determinismus-vertrag) nachgezogen |
| F-4 | Commit per `--amend` repariert; Gate-Läufe seitdem ohne Pipe mit geprüftem Exit-Code. Wiederholte sich später erneut ⇒ als dauerhafte Arbeitsregel festgehalten |

## Negativbefunde (geprüft, ohne Befund)

- `internal/adapter/driven/config` — Decode-Pfad, Alias-Behandlung, fail-closed-Fälle; `decodeZoneList`-Generalisierung lässt die `tech`-Meldungen byte-gleich.
- `internal/adapter/driven/graph` — nur die Legendenzeile; Escaping-Vertrag und Determinismus unberührt.
- `internal/adapter/driven/report` — unberührt.
- `internal/cli` — nur das kommentierte `constructs`-Beispiel im `--print-config`-Gerüst; `a-check.mk`/Distributionspfad unberührt (Fragment-Parität in `image-test` grün).
- `internal/hexagon/port` — Port-Schnittstellen unverändert, keine `ARC-*`-Berührung.
- Determinismus der Treffer-Reihenfolge (Eintrags-Index statt Mustertext) — geprüft, ohne Befund.
- Konsumenten-Regression (b-cad, d-check, d-migrate, m-trace) — je 0 Befunde, unverändert.

## Nachlauf

F-2 ist offen (Mikro-CR). F-4 hat einen Steering-Loop ausgelöst: nach dem **zweiten** Auftreten
derselben Klasse wurde die Arbeitsregel „Gate-Läufe nie pipen" dauerhaft festgehalten — Modul 10
§Pflege sieht bei Wiederholung genau das vor.
