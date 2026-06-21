# slice-006 — CI: GitHub-Actions-Workflow (gates + image-test + traceability)

**Status:** open (Backlog; wartet auf Trigger/Priorisierung).
**Welle:** welle-08-ci.
**Bezug:** [`AGENTS.md`](../../../../AGENTS.md) §3.1 (Docker/make-only) + §5
(Traceability: `AC-*`/`ADR-*`-ID je Commit);
[AC-FA-DIST-001](../../../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk)
(Image-Akzeptanz), [AC-QA-01](../../../../spec/lastenheft.md#ac-qa-01--determinismus)/[AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)
(Determinismus nativ == Container). Stack-Vorbild `d-check` (`.github/workflows/ci.yml`).

## Ziel

a-check bekommt — wie `d-check` — eine **PR-/Push-CI**, die denselben Vertrag
auf jede Integration zieht (pre-integration statt nur Release). Das schließt
zugleich die im Stop-Hook dokumentierte **Restlücke** „frischer Klon ohne
Gate-State wird durchgewunken — CI ist dort das Netz": dieses Netz existiert
heute nicht und wird mit diesem Slice real.

Bestandteile:

- **`make ci`** = `gates` + **`image-test`** (Image-Integrationstests gegen das
  gebaute Runtime-Image: `--print-mk`/`--print-config`/unbekanntes Flag +
  nativ==Container-Determinismus eines echten Scans).
- **`make trace-check`** — Traceability-Gate: jede Commit-Message im Range nennt
  eine `AC-*`/`ADR-*`/`MR-*`/`slice-NNN`-ID (Selbsttest + HEAD; `RANGE=a..b` für CI).
- **Dockerfile-OCI-Labels** (`org.opencontainers.image.*`) + `VERSION`-Build-Arg —
  Voraussetzung für den späteren `welle-05-release`-Label-Verify, hier vorgezogen.
- **`.github/workflows/ci.yml`** — `pull_request` + `push` (Tags ausgenommen),
  `permissions: {}`, SHA-gepinnte Actions; Schritte: `trace-check` (Commit-Range,
  fail-fast) → `make ci`.

## Definition of Done

- `make image-test` prüft [AC-FA-DIST-001](../../../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk) (Happy/Boundary/Negative) + nativ==Container gegen das gebaute Image.
- `make ci` aggregiert `gates` + `image-test`.
- `make trace-check` mit Selbsttest; `RANGE=` für CI; erkennt fehlende IDs fail-closed.
- Dockerfile trägt `org.opencontainers.image.*`-Labels inkl. `version` aus `VERSION`-Build-Arg.
- `.github/workflows/ci.yml` existiert (SHA-gepinnt, `permissions: {}`, Tags ausgenommen).
- [`AGENTS.md`](../../../../AGENTS.md) §4 + [`harness/README.md`](../../../../harness/README.md) um `ci`/`image-test`/`trace-check` ergänzt.
- Beleg: `make ci` grün; `make trace-check` grün.

## Offene Fragen

- **Lokaler `commit-msg`-Hook:** `trace-check.sh --message` wird mitgeliefert,
  aber der git-Hook (`.githooks` + `make hooks`) wird zunächst **nicht** verdrahtet —
  Folge-Kandidat, oder gleich mit aufnehmen?
- **Release-Workflow (`release.yml`):** bewusst out-of-scope dieses Slices
  (eigene `welle-05-release`); die OCI-Labels werden hier nur **vorbereitet**.
- **Image-Tag:** `make build` erzeugt `a-check:dev`; `image-test`/`ci` laufen
  dagegen. Die GHCR-Tag-Strategie klärt `welle-05-release`.
