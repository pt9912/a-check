# slice-007 — Release-Pipeline: release.yml + `:latest`-Politik-ADR

**Status:** open (Backlog; wartet auf Trigger/Priorisierung).
**Welle:** welle-05-release (Teil-Lieferung: Pipeline; der erste getaggte
GHCR-Release + die Pilot-Einbindung bleiben Wellen-Restarbeit).
**Bezug:** [`AGENTS.md`](../../../../AGENTS.md) §3.1 (Docker/make-only CI-Pfad);
[AC-FA-DIST-001](../../../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk)
(GHCR-Image/Tagging),
[AC-QA-03](../../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit)
(Digest-Pin-Reproduzierbarkeit),
[AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze);
schärft [SPEC-DIST-001](../../../../spec/spezifikation.md#spec-dist-001--laufzeitform-und-distribution)
über eine neue ADR (`:latest`-Politik). Stack-Vorbild `d-check`
(`.github/workflows/release.yml` + dessen Tag-Politik-ADR).

## Ziel

Die GHCR-Release-Pipeline für a-check anlegen — die computational Hälfte von
`welle-05-release`. Nach diesem Slice ist alles **Lokal-Lieferbare** da; es fehlt
nur noch das, was ein GitHub-Remote/GHCR braucht (erster `v*`-Tag + Pilot-Repo).

Bestandteile:

- **ADR (`:latest`-Politik)** — frische Entscheidung (kein Supersede; a-checks
  [ADR-0004](../../adr/0004-distribution-image-mk.md) trifft keine Tag-Aussage): `:latest` nur für stabile Releases
  (`vX.Y.Z`), Prereleases ohne `:latest`, Konsumenten pinnen verbindlich per
  `@sha256:`-Digest. Schärft die Distributions-Spec (siehe §Bezug). In
  [`releasing.md`](../../../../docs/user/releasing.md) bereits als „entscheidet
  welle-05 per ADR" angekündigt.
- **`.github/workflows/release.yml`** — Trigger `v*`-Tags: SemVer-Validate
  (fail-fast) → `make ci VERSION=…` → GHCR-Login → Tag (`:latest` nur stabil) →
  OCI-Label-Verify (`org.opencontainers.image.version` == Tag) → Push →
  GitHub-Release mit Digest-Pin in den Notes. SHA-gepinnte Actions,
  `permissions:` minimal (`contents/packages: write`).
- **`releasing.md`** — von „Zielbild welle-05" auf „Pipeline existiert" umstellen
  (ehrlich: noch kein getaggter Release), ADR verlinken.

## Definition of Done

- ADR zur `:latest`-Politik angelegt (Status Proposed → Sign-off durch Auftraggeber), im ADR-Index ergänzt.
- `.github/workflows/release.yml` existiert (SHA-gepinnt, minimale `permissions`, SemVer-Validate + `IS_STABLE`-Verzweigung + Label-Verify).
- OCI-Labels lokal belegt (das gebaute Image trägt `org.opencontainers.image.version` = `VERSION`).
- [`releasing.md`](../../../../docs/user/releasing.md) auf den realen Pipeline-Stand aktualisiert + ADR verlinkt.
- `make ci` grün; `make doc-check` grün.

## Offene Fragen

- **Build-Tag für den Release-Push:** `make build` erzeugt `a-check:dev`;
  `release.yml` taggt `a-check:dev → ghcr.io/pt9912/a-check:v<version>`. Reicht,
  oder soll ein eigener `:release`/`:latest`-Build-Tag her? (Tendenz: `:dev`
  wiederverwenden, eine Build-Quelle.)
- **Wellen-Abschluss:** der erste echte `v*`-Tag + die Pilot-Einbindung
  (Konsumenten-Repo) brauchen ein GitHub-Remote/GHCR — bleibt Wellen-Restarbeit
  nach diesem Slice (Folge-Trigger).
