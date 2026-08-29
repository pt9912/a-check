# MR-005 — Referenzmatrix: intra-Spec-Richtung + ADR→Slice-Disziplin (d-check-Angleichung)

- **Datum:** 2026-07-04
- **Geltungsbereich:** [`.d-check.yml`](../../.d-check.yml) (`matrix`-Modul),
  [`docs/plan/adr/`](../../docs/plan/adr/)
- **Adaption:** a-check übernimmt die vollständigere Referenzmatrix-Kodierung des
  Schwester-Repos `d-check` (dort `DC-FA-MTX-001/002/003`). Zusätzlich zu den bisherigen
  Cross-Klassen-Regeln (`spec-straten → adr`, `spec-straten → slice`):
  - **intra-Spec-Richtung** (`order` + `direction: no-downward`): Rang
    `lastenheft → spezifikation → architecture` (autoritativste Schicht zuerst); ein
    Abwärtsverweis zwischen Spec-Straten (auch transitiv, z. B. Lastenheft →
    Architektur) ist `matrix-downward`.
  - **`adr → slice`-Regel + Token-Erkennung** (`token: 'slice-\d{3}'`): eine
    Slice-Kennung im **ADR-Körper** ist `matrix-forbidden` — außer per Provenance-Marker
    `<!-- d-check:status-provenance -->` deklariert. ADRs verweisen abwärts **nur** als
    deklarierte Provenance/Verifikations-Zeiger, **nie** als Entscheidungsgrundlage
    (die Argumentation läuft aufwärts über Spec/Verhalten, `Schärft:`-Feld).
- **Grandfathering:** die vor der Übernahme `Accepted`-ADRs (0001–0020) sind immutabel
  ([`AGENTS.md` §3.5](../../AGENTS.md#35-adrs-sind-nach-accepted-immutable)) und nennen
  Slices als legitime Verifikations-Zeiger im Körper — sie werden per `exempt-paths`
  ganz übersprungen; neue ADRs **ab 0021** sind **slice-token-frei** (aus
  AGENTS/Modul/Verhalten argumentiert) **oder** tragen — falls sie einen Slice als
  Provenance/Verifikation nennen — den Provenance-Marker (gelebte Praxis: 0021/0022
  sind slice-token-frei, kein Marker nötig).
- **Begründung:** schärft [`AGENTS.md` §3.4](../../AGENTS.md#34-architektur-sprach-meilensteinfrei-spec-straten-nie-abwärts)
  maschinell (Spec-Straten-Abwärts **und** ADR-aus-Planung-Argumentation). a-checks
  gepinntes `d-check` unterstützt die Schlüssel (verifiziert 2026-07-04 gegen v0.35.0);
  der Pin ist seit slice-019 auf `v0.37.1`, seit slice-036 auf `v0.51.1` gehoben.
- **Auflösungs-Trigger:** permanent.
