# MR-012 — Referenz-Richtung maschinell, ADRs 0001–0020 grandfathered (schärft [`MR-005`](../conventions.md#mr-005))

- **Status:** Accepted
- **Datum:** 2026-08-29
- **Geltungsbereich:** [`.d-check.yml`](../../.d-check.yml) (`matrix`-Modul),
  [`docs/plan/adr/`](../../docs/plan/adr/)
- **Ersetzt-Baseline-Regel:** [`grundlagen-referenz-richtung.md` §Referenz-Richtung (SDP)](../../.harness/baseline/v6.0.0/regelwerk/grundlagen-referenz-richtung.md#referenz-richtung-sdp-wer-darf-wen-referenzieren)
- **Adaption:** Die Baseline verlangt die Aufwärts-Richtung für **jede** Kante im bindenden Text.
  a-check nimmt die vor der Übernahme `Accepted`-ADRs **0001–0020** per `exempt-paths` ganz aus:
  sie nennen Slice-Kennungen im Körper als Verifikations-Zeiger.
- **Begründung:** Diese ADRs sind immutabel ([`AGENTS.md`](../../AGENTS.md) §3.5) — die Regel
  ließe sich auf sie nur durch einen Verstoß gegen eine andere Regel anwenden. Ab [`ADR-0021`](../../docs/plan/adr/0021-commits-modul-trace-check.md) gilt
  die Baseline-Richtung ungemildert: slice-token-frei oder mit deklariertem Provenance-Marker.
  Die Richtungs-Regel selbst ist mit `v5.12.0` Default geworden und wird hier **nicht** mehr
  behauptet — was bleibt, ist die Ausnahme und ihre maschinelle Kodierung.
- **Auflösungs-Trigger:** permanent. Die Grandfather-Menge ist geschlossen und wächst nicht;
  sie verschwindet erst, wenn keine der zwanzig ADRs mehr `Accepted` ist.
- **Löst auf:** [`MR-005`](../conventions.md#mr-005)
- **Ausgelöst durch Baseline-Stand:** `v5.12.0`
