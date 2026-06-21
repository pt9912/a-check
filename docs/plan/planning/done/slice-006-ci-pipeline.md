# slice-006 — CI: GitHub-Actions-Workflow (gates + image-test + traceability)

**Status:** done.
**Welle:** welle-08-ci.
**Bezug:** [`AGENTS.md`](../../../../AGENTS.md) §3.1 (Docker/make-only) + §5
(Traceability: `AC-*`/`ADR-*`-ID je Commit);
[AC-FA-DIST-001](../../../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk)
(Image-Akzeptanz), [AC-QA-01](../../../../spec/lastenheft.md#ac-qa-01--determinismus)/[AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)
(Determinismus nativ == Container). Stack-Vorbild `d-check` (`.github/workflows/ci.yml`).

---

## 1. Ziel

a-check bekommt — wie `d-check` — eine **PR-/Push-CI**, die denselben Vertrag
auf jede Integration zieht (pre-integration statt nur Release) und die im
Stop-Hook dokumentierte **Restlücke** „frischer Klon ohne Gate-State — CI ist
dort das Netz" real macht (vorher: das Netz fehlte, ein dokumentierter Verweis
auf Nicht-Vorhandenes).

## 2. Definition of Done

- [x] `make image-test` prüft [AC-FA-DIST-001](../../../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk) (Happy/Boundary/Negative) + nativ==Container gegen das gebaute Image.
- [x] `make ci` aggregiert `gates` + `image-test`.
- [x] `make trace-check` mit Selbsttest; `RANGE=` für CI; erkennt fehlende IDs fail-closed.
- [x] Dockerfile trägt `org.opencontainers.image.*`-Labels inkl. `version` aus `VERSION`-Build-Arg.
- [x] `.github/workflows/ci.yml` existiert (SHA-gepinnt, `permissions: {}`, Tags ausgenommen).
- [x] [`AGENTS.md`](../../../../AGENTS.md) §4/§5 + [`harness/README.md`](../../../../harness/README.md) um `ci`/`image-test`/`trace-check` ergänzt.
- [x] Beleg: `make ci` grün; `make trace-check` grün.

## 3. Umsetzung

- `tools/image-test.sh` — extrahiert das statische Binary aus dem Image
  (`docker cp`) und vergleicht nativ vs. Container byte-identisch:
  `--print-mk` (Happy), `--print-config` read-only (Boundary), unbekanntes Flag
  → Exit 2 (Negative), Verstoß-Fixture → `core-impurity` + Exit 1 (Scan).
- `tools/trace-check.sh` — eine Wahrheit für drei Aufrufer (`--message`-Hook,
  `--range`-CI, `HEAD`-lokal); Negativ-Selbsttest; ID-Muster `AC-/ADR-/MR-/slice`.
- `Dockerfile` — `VERSION`-Build-Arg + `org.opencontainers.image.*`-Labels
  (runtime-Stage). `Makefile` — `VERSION`, `build --build-arg`, Targets
  `image-test`/`ci`/`trace-check`.
- `.github/workflows/ci.yml` — `pull_request` + `push` (Tags ausgenommen),
  `permissions: {}`, SHA-gepinnte Actions; Schritte `trace-check` (Range,
  fail-fast) → `make ci`.

## 4. Closure-Notiz (nach `done/`)

**Belege:** `make ci` grün — `gates` (lint/test/coverage 92,60 %/arch-check/
doc-check/gate-consistency/guard-selftest/record-gates) + `image-test` (4/4:
Happy/Boundary/Negative/Scan, nativ==Container byte-identisch,
[AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze));
`make trace-check` grün (Selbsttest gefeuert, HEAD trägt eine ID).

**Lerneintrag (Steering-Loop):**

- *Geschlossene Restlücke:* der Stop-Hook verwies auf „CI als Netz" für den
  frischen Klon — dieses Netz existiert jetzt (`ci.yml`). Ein dokumentierter
  Verweis auf Nicht-Vorhandenes (selbst eine milde Harness-Lüge) ist aufgelöst.
- *Neue Sensoren:* `image-test` prüft den Distributions-Vertrag gegen das
  *gebaute Image* (nicht nur als Go-Test) und beweist nativ==Container
  byte-identisch; `trace-check` macht die AGENTS-§5-Commit-Regel maschinell —
  beide mit Negativ-Selbsttest (Muster `gate-consistency`).
- *Vorgezogene Release-Voraussetzung:* die OCI-Labels + `VERSION`-Build-Arg
  liegen jetzt; `welle-05-release` muss sie nur noch aus dem Tag setzen und
  per Label-Verify prüfen — kein Dockerfile-Umbau mehr nötig.

**Offene Fragen — aufgelöst:**

- *Lokaler `commit-msg`-Hook:* `trace-check.sh --message` ist mitgeliefert, der
  git-Hook (`.githooks` + `make hooks`) wird **nicht** verdrahtet — bewusster
  Folge-Kandidat (open), die CI deckt den Pflichtfall bereits ab.
- *`release.yml`:* out-of-scope (eigene `welle-05-release`); hier nur die
  OCI-Labels vorbereitet.
- *Image-Tag:* `image-test`/`ci` laufen gegen `a-check:dev` (der `make build`-Tag);
  die GHCR-Tag-Strategie klärt `welle-05-release`.

**Folge-Kandidaten (open):** `welle-05-release` (GHCR-Release + `release.yml` +
Label-Verify, nutzt die hier gelegten Labels); optional der lokale
`commit-msg`-Hook.

## 5. Sub-Area-Modus-Begründung

### Sub-Area: CI-/Integrations-Harness

- **Modus:** GF — Workflow/Skripte neu angelegt, kein Bestand zu inventarisieren.
- **Konventionen-Dichte:** hoch (Stack-Vorbild `d-check` ci.yml/image-test/trace-check).
- **Phase-Reife:** Phase 5 — Gate erzwingen, `make ci`/`trace-check` grün.
- **Evidenz-/Diskrepanz-Risiko:** niedrig (gespiegelte Logik + Selbsttests + lokal verifiziert).
- **Reconciliation-Aufwand:** keiner; Folge-Welle `welle-05-release`.
