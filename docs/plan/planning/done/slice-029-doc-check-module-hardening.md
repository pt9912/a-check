# slice-029 — doc-check-Modul-Hardening (spans + vcs; codepaths/planning verworfen)

**Status:** done (2026-07-04). **Typ:** Harness-Hardening (Doku-Gate), nicht konsumenten-gated.
**Bezug:** schärft [`AGENTS.md` §3.5](../../../../AGENTS.md#35-adrs-sind-nach-accepted-immutable)
(ADR-Immutabilität) maschinell; [`.d-check.yml`](../../../../.d-check.yml)-Module,
[`ci.yml`](../../../../.github/workflows/ci.yml). Folge von
[slice-019](slice-019-dcheck-mk-print-mk-angleichung.md) (Pin v0.37.1 brachte die
Module) + [slice-028](slice-028-dcheck-matrix-konvergenz.md). [Roadmap](../in-progress/roadmap.md).

## 1. Auslöser

Nach dem v0.37.1-Pin (slice-019) stehen d-checks volle Module offen. a-check fuhr im
mandatory `doc-check` nur `links/anchors/ids/matrix/hostpaths`. Frage: welche der übrigen
lohnen? Geprüft an der realen Wirkung, nicht am Prospekt.

## 2. Umgesetzt

1. **`spans`** → in die `modules`-Liste (mandatory `doc-check`). Hermetisch, 0 Befunde,
   fängt offene Code-Spans / verschachtelte Links. Sauberer Sofort-Gewinn.
2. **`vcs`** (DC-FA-VCS-001) → **ADR-Immutabilität (§3.5) maschinell**. Config-Block in
   `.d-check.yml` (an a-checks `- **Status:** Accepted`-Format angepasst); `make doc-immutable`
   fährt `--enable vcs`. **Die CI ruft es über die PR-/Push-Range** (`ci.yml`, neben
   `trace-check`) — damit ist §3.5 durchgesetzt statt nur konventionell. Verifiziert:
   Core-Drift eines Accepted-ADR → `core-drift-vcs` (Exit 2); saubere Range, neue ADR-Datei
   und Proposed→Accepted-Übergang → 0 (kein Falsch-Positiv).

## 3. Bewusst NICHT adoptiert

- **`codepaths`** (DC-FA-CODE-001) — **Domänen-Misfit.** a-check *ist* ein Tool über
  Import-Auflösung; seine Doku (Spec, Handbuch, Resolution-ADRs, Backend-Slices) ist voller
  **illustrativer Specifier** (`./b`, `../core/model`, `./db`, `../adapters/db` …). Aktivierung
  ergab **~40 `codepath-missing`**, fast alle solche Beispiele (dazu Beispiel-Configs
  `internal/core/**` und das konzeptionelle `tools/arch-check.sh` = das von a-check ersetzte
  Fremd-Skript). codepaths kann Beispiel-Specifier nicht von echten Datei-Verweisen trennen;
  Aktivierung hieße ~40 `d-check:ignore`-Marker + Dauerlast bei jedem neuen Beispiel — Kosten
  über Nutzen (2–3 echte Fälle). Für a-check ungeeignet.
- **`planning`** (DC-FA-PLAN-001) — **vakuum auf a-checks Roadmap.** Das Modul prüft „Ruhe-Marker
  im `## Aktuelle Welle`-Block genau dann, wenn kein `slice-*` in `in-progress/` liegt". a-checks
  Roadmap ist **welle-organisiert** (Prosa unter `## Aktuelle Welle`), nicht slice-listend — Test
  mit einem Dummy-Slice in `in-progress/` ließ das Verdikt bei **0** (das Modul findet kein
  Slice-Listing zum Abgleich). Echter Nutzen erst nach Roadmap-Umbau auf die Marker-Konvention —
  eigene Entscheidung, hier nicht erzwungen.

## 4. Konsolidierungs-Kandidat (eigener Slice)

- **`commits`** (DC-FA-COMMITS-001) — funktioniert mit a-checks `AC-`/`ADR-`/`MR-`/`slice`-Mustern
  (getestet: gleiches Verdikt wie `make trace-check`), **überlappt** aber a-checks *eigenes*
  `trace-check` vollständig. Kandidat, `trace-check` durch das Modul zu **ersetzen** (Skript-Copies
  verringern — d-check tat genau das per eigenem Ablöse-ADR). Eigener Ablöse-Slice + ADR, nicht
  Teil dieses Modul-Hardenings.

## 5. Verifikation & Closure

- **`make gates` grün**; `doc-check` mit `spans` = 0 Befunde. `gate-consistency` grün
  (doc-immutable-Zeile in AGENTS §4 auf `real (slice-029, CI)` gehoben).
- **vcs de-riskt** (§2): Core-Drift-Fang + Kein-Falsch-Positiv auf neuer/akzeptierter ADR-Datei.
- **Netto:** ein hermetischer Modul-Gewinn (`spans`) + eine echte Vertrags-Durchsetzung
  (`vcs`/§3.5, CI). codepaths/planning als Misfit/vakuum dokumentiert, `commits` als
  Konsolidierung vorgemerkt. **Lerneintrag:** Modul-Prospekt ≠ Fit — die Wirkung gegen den
  realen Doku-Bestand entscheidet (codepaths klang top, ist aber für a-checks Domäne Lärm).
