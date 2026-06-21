# slice-007 — Release-Pipeline: release.yml + `:latest`-Politik-ADR

**Status:** done (Teil-Lieferung von welle-05-release; Welle bleibt offen).
**Welle:** welle-05-release.
**Bezug:** [`AGENTS.md`](../../../../AGENTS.md) §3.1 (Docker/make-only CI-Pfad);
[AC-FA-DIST-001](../../../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk)
(GHCR-Image/Tagging),
[AC-QA-03](../../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit)
(Digest-Pin-Reproduzierbarkeit),
[AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze);
[ADR-0007](../../adr/0007-latest-tag-politik.md) (`:latest`-Politik) schärft
[SPEC-DIST-001](../../../../spec/spezifikation.md#spec-dist-001--laufzeitform-und-distribution).
Stack-Vorbild `d-check` (`.github/workflows/release.yml`).

---

## 1. Ziel

Die GHCR-Release-Pipeline anlegen — die lokal lieferbare Hälfte von
`welle-05-release`. Nach diesem Slice ist die Pipeline real; es fehlt nur noch,
was ein GitHub-Remote/GHCR braucht: der erste `v*`-Tag und die Pilot-Einbindung
in ein Konsumenten-Repo (Wellen-Restarbeit).

## 2. Definition of Done

- [x] [ADR-0007](../../adr/0007-latest-tag-politik.md) (`:latest`-Politik) angelegt (Status **Proposed** → Sign-off durch Auftraggeber ausstehend), im [ADR-Index](../../adr/README.md) ergänzt.
- [x] [`.github/workflows/release.yml`](../../../../.github/workflows/release.yml) existiert (SHA-gepinnt, minimale `permissions`, SemVer-Validate + `IS_STABLE`-Verzweigung + OCI-Label-Verify).
- [x] OCI-Labels lokal belegt (`docker inspect a-check:dev` trägt `org.opencontainers.image.version`).
- [x] [`releasing.md`](../../../../docs/user/releasing.md) auf den realen Pipeline-Stand aktualisiert + auf die ADR verlinkt.
- [x] `make ci` grün; `make doc-check` grün.

## 3. Umsetzung

- [ADR-0007](../../adr/0007-latest-tag-politik.md) — frische Tagging-Entscheidung
  (kein Supersede; a-checks [ADR-0004](../../adr/0004-distribution-image-mk.md)
  trifft keine Tag-Aussage): `:latest` nur stabil, Prereleases ohne, Konsum per
  Digest. Schärft die Distributions-Spec (siehe §Bezug).
- [`.github/workflows/release.yml`](../../../../.github/workflows/release.yml) —
  `v*`-Tags: SemVer-Validate → `make ci VERSION=…` → GHCR-Login → Tag
  (`a-check:dev → ghcr.io/pt9912/a-check:v<version>`, `:latest` nur stabil) →
  OCI-Label-Verify → Push → GitHub-Release mit Digest-Pin in den Notes.
- [`releasing.md`](../../../../docs/user/releasing.md) — von „Zielbild" auf
  „Pipeline existiert" umgestellt; auf die ADR verlinkt.

## 4. Closure-Notiz (nach `done/`)

**Belege:** `make ci` grün (gates + image-test); `make doc-check` grün
(33 Dateien). Die OCI-Labels sind lokal belegt — das gebaute `a-check:dev`
trägt `org.opencontainers.image.version=0.0.0-dev` (Default; die Pipeline setzt
sie aus dem Tag). Der Workflow selbst läuft erst mit GitHub-Remote/GHCR — die
strukturelle Kontrolle ist die SemVer-Validate-Stufe + `IS_STABLE`-Verzweigung
+ Label-Verify (Fitness Function der ADR).

**Lerneintrag (Steering-Loop):**

- *Frische ADR statt Supersede:* a-check trifft die `:latest`-Entscheidung
  erstmalig (diese ADR), während `d-check` sie als Ratifikation einer
  „kein-latest"-Klausel führte — die Baseline wird adaptiert, nicht kopiert.
- *Ehrliche Teil-Lieferung:* welle-05-release **schließt nicht** mit diesem
  Slice — der erste getaggte Release + Pilot brauchen ein GitHub-Remote/GHCR.
  Das ist explizit als Wellen-Restarbeit ausgewiesen statt „done" zu suggerieren
  (vermeidet eine Harness-Lüge auf Wellen-Ebene).
- *Vorgezogene Voraussetzung genutzt:* die OCI-Labels + `VERSION`-Build-Arg aus
  slice-006 werden hier direkt vom Label-Verify konsumiert — kein Dockerfile-Umbau.

**Offene Fragen — aufgelöst:**

- *Build-Tag:* `release.yml` taggt `a-check:dev` (der `make build`-Tag) nach GHCR
  um — eine Build-Quelle, kein Extra-Tag.
- *ADR-Status:* Proposed; Sign-off durch den Auftraggeber ausstehend (Muster
  den Fundament-ADR-Sign-offs).

**Folge-Kandidaten:** Wellen-Restarbeit `welle-05-release` (erster `v*`-Tag nach
GitHub-Remote-Setup + Pilot-Einbindung); Sign-off dieser ADR; optional der lokale
`commit-msg`-Hook (aus slice-006).

## 5. Sub-Area-Modus-Begründung

### Sub-Area: Release-/Distributions-Harness

- **Modus:** GF — Workflow/ADR neu angelegt, kein Bestand zu inventarisieren.
- **Konventionen-Dichte:** hoch (Stack-Vorbild `d-check` release.yml + ADR-Format).
- **Phase-Reife:** Phase 5 — Pipeline + Politik stehen; `make ci` grün.
- **Evidenz-/Diskrepanz-Risiko:** mittel — die CI-YAML selbst läuft lokal nicht;
  Kontrolle über SemVer-Validate/`IS_STABLE`/Label-Verify + lokalen Label-Beleg.
- **Reconciliation-Aufwand:** Wellen-Restarbeit (erster Tag + Pilot), bewusst offen.
