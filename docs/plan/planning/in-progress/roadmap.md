# Roadmap

**Status:** Aktiv. **Letzte Änderung:** 2026-06-21.

**Format-Regel:** Die Roadmap ist eine Reihenfolge von **Wellen**, keine
Reihenfolge von Terminen. Termine erscheinen — falls überhaupt — als
Konsequenz der Wellen-Schätzung, nicht als Treiber. Die Roadmap steht
außerhalb der normativen Klammer: sie *orchestriert* Slices und Wellen,
erzeugt aber keine Spezifikation (Regelwerk Modul 6).

> **Hinweis zur Slice-Buchführung.** Die abgeschlossenen Slices liegen als
> Planning-Harness-Dateien unter `done/` (retroaktiv nachgezogen, Regelwerk
> Modul 5) mit Closure-Notiz + Lerneintrag; ab `slice-004` entstehen sie
> regulär über den Lifecycle (`open → next → in-progress → done`).

---

## Aktuelle Welle

**`welle-05-release` teils geliefert — Pipeline steht, wartet auf den ersten Tag.**
Zuletzt abgeschlossen: [slice-007 §4](../done/slice-007-release-pipeline.md#4-closure-notiz-nach-done)
(Release-Pipeline `.github/workflows/release.yml` + [ADR-0007](../../adr/0007-latest-tag-politik.md)
`Proposed`); davor welle-08-ci (`slice-006` — PR-/Push-CI;
[slice-006 §4](../done/slice-006-ci-pipeline.md#4-closure-notiz-nach-done)).
Alle inneren Gates sind real und grün (`make ci` = `gates` + `image-test`;
`make trace-check` grün). Noch **kein getaggtes GHCR-Release** (Status 0.1.0):
der erste `v*`-Tag und die Pilot-Einbindung brauchen ein GitHub-Remote/GHCR und
bleiben offene Restarbeit von `welle-05-release`; offen außerdem der
[ADR-0007](../../adr/0007-latest-tag-politik.md)-Sign-off.

## Nächste Wellen

| Welle | Trigger | Wichtigste Inhalte | Status |
|---|---|---|---|
| welle-05-release | Image-Veröffentlichung | **Pipeline geliefert** ([slice-007](../done/slice-007-release-pipeline.md): `release.yml` + [ADR-0007](../../adr/0007-latest-tag-politik.md)). Offen: erster getaggter GHCR-Release + `@sha256:`-Digest-Pin in `a-check.mk` ([AC-FA-DIST-001](../../../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk), [AC-QA-03](../../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit)); Pilot-Einbindung in ein Konsumenten-Repo | teils geliefert |
| welle-06-sprach-backends | Bedarf | Ausbau/Härtung der Extraktion je Zielsprache; opt-in toolchain-gestützte Backends ([AC-FA-EXTRACT-001](../../../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion) Out-of-Scope-Re-Eval) | offen |

_(Kein fixer Termin — Wellen feuern auf Trigger.)_

## Meilensteine

| Meilenstein | Welle(n) | Status |
|---|---|---|
| M1: Spec-Fundament steht (Lastenheft + Spezifikation + Architektur + Fundament-ADRs) | welle-01/02 | **erreicht** (2026-06-21) |
| M2: Dogfooding — a-check prüft die eigene Architektur grün ([AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)) | welle-03 | **erreicht** (2026-06-21) |
| M3: erstes GHCR-Release + Pilot-Einbindung | welle-05 | offen |

## Abhängigkeitsgraph

```mermaid
flowchart LR
    W0[welle-00-bootstrap]
    W1[welle-01-fundament]
    W2[welle-02-spec]
    W3[welle-03-implementierung]
    W4[welle-04-durchsetzungsschicht]
    W5[welle-05-release]

    W0 --> W1 --> W2 --> W3 --> W4 --> W5
```

## Abgeschlossene Wellen

| Welle | Abschluss | Closure-Beleg |
|---|---|---|
| welle-00-bootstrap | 2026-06-20 | Harness-Trias + Lastenheft 0.1.0 + Doku-Gate `make doc-check` ([CHANGELOG](../../../../CHANGELOG.md)) |
| welle-01-fundament | 2026-06-21 | [slice-001 §7](../done/slice-001-fundament-adrs.md#7-closure-notiz-nach-done) — Fundament-ADRs [ADR-0001](../../adr/0001-go-impl-sprache.md)…[ADR-0004](../../adr/0004-distribution-image-mk.md) `Accepted` |
| welle-02-spec | 2026-06-21 | [slice-002 §7](../done/slice-002-architektur-spezifikation.md#7-closure-notiz-nach-done) — Technik-/Sicht-Stratum (`SPEC-*`/`ARC-*`) |
| welle-03-implementierung | 2026-06-21 | [slice-003 §7](../done/slice-003-implementierung-gates.md#7-closure-notiz-nach-done) — Go-Implementierung + Gates; [ADR-0005](../../adr/0005-lint-profil.md)/[ADR-0006](../../adr/0006-coverage-gate.md) `Accepted` |
| welle-04-durchsetzungsschicht | 2026-06-21 | [slice-004 §4](../done/slice-004-durchsetzungsschicht.md#4-closure-notiz-nach-done) — Meta-Gates `gate-consistency`/`record-gates` + `.claude`-Stop-Hook |
| welle-07-command-guard | 2026-06-21 | [slice-005 §4](../done/slice-005-command-guard.md#4-closure-notiz-nach-done) — PreToolUse-Command-Guard (Tool-Call-Gate); Durchsetzungsschicht vollständig |
| welle-08-ci | 2026-06-21 | [slice-006 §4](../done/slice-006-ci-pipeline.md#4-closure-notiz-nach-done) — PR-/Push-CI (`ci.yml`): `make ci` (+ `image-test`) + `make trace-check`; Dockerfile-OCI-Labels |
