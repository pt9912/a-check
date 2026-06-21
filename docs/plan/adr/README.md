# ADR-Index

Architecture Decision Records dieses Repos. Konventionen:

- **Dateiname:** `<NNNN>-<kurzer-titel-kebab>.md` (vierstellig, zero-padded).
- **Status:** `Proposed` → `Accepted`; danach immutable. Ablösung nur
  via neue ADR mit `Supersedes ADR-NNNN` (Status der alten wird
  `Superseded by ADR-NNNN`).
- Jede ADR deklariert im `**Schärft:**`-Feld aufwärts, welche
  Spec-Stelle sie verbindlich macht (nie das Lastenheft).
- Neue ADRs werden in der Tabelle unten ergänzt.

Die **Fundament-ADRs** (slice-001) legen die technische Basis fest, die das
Technik-Stratum [`spec/spezifikation.md`](../../../spec/spezifikation.md)
(slice-002) formalisiert. Ihr `**Schärft:**`-Feld zeigt nun aufwärts auf die
jeweilige `SPEC-*`-Stelle; [ADR-0001](0001-go-impl-sprache.md) bleibt `—`
(die Spezifikation ist sprachneutral, die Sprachwahl koppelt an keine
Spec-§). Status **`Accepted`** (Acceptance-Sign-off 2026-06-21); ab jetzt
immutable — Korrekturen nur via Folge-ADR mit `Supersedes`.

| ID | Titel | Status | Datum | Bezug |
|---|---|---|---|---|
| [ADR-0001](0001-go-impl-sprache.md) | Go als Implementierungssprache | Accepted | 2026-06-21 | [AC-QA-01](../../../spec/lastenheft.md#ac-qa-01--determinismus), [AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze), [AC-QA-03](../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit), [AC-FA-DIST-001](../../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk) |
| [ADR-0002](0002-text-heuristische-extraktion.md) | Text-heuristische Import-Extraktion | Accepted | 2026-06-21 | [AC-FA-EXTRACT-001](../../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion), [AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze), [AC-QA-01](../../../spec/lastenheft.md#ac-qa-01--determinismus) |
| [ADR-0003](0003-config-modell-a-check-yml.md) | Config-Modell `.a-check.yml` | Accepted | 2026-06-21 | [AC-FA-CONF-001](../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml), [AC-FA-CLI-001](../../../spec/lastenheft.md#ac-fa-cli-001--aufruf-scan-wurzel-und-exit-codes) |
| [ADR-0004](0004-distribution-image-mk.md) | Distribution: Image + `--print-mk`/`a-check.mk` | Accepted | 2026-06-21 | [AC-FA-DIST-001](../../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk), [AC-QA-03](../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit), [AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze) |
| [ADR-0005](0005-lint-profil.md) | Lint-Profil (golangci-lint v2, SOLID-nahe Linter) | Proposed | 2026-06-21 | [`AGENTS.md`](../../../AGENTS.md) §3.2/§3.6 |
| [ADR-0006](0006-coverage-gate.md) | Coverage-Gate (Gesamt-Schwelle 90 %) | Proposed | 2026-06-21 | [`AGENTS.md`](../../../AGENTS.md) §3.6/§4 |
