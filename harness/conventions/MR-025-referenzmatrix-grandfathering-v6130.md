# MR-025 — Referenz-Richtung maschinell, ADRs 0001–0020 grandfathered, Zeiger auf `v6.13.0` (löst [`MR-012`](../conventions.md#mr-012) ab)

- **Status:** Accepted
- **Datum:** 2026-09-29
- **Geltungsbereich:** [`.d-check.yml`](../../.d-check.yml) (`matrix`-Modul),
  [`docs/plan/adr/`](../../docs/plan/adr/)
- **Ersetzt-Baseline-Regel:** [`grundlagen-referenz-richtung.md` §Referenz-Richtung (SDP)](../../.harness/baseline/v6.13.0/regelwerk/grundlagen-referenz-richtung.md#referenz-richtung-sdp-wer-darf-wen-referenzieren)
- **Adaption:** unverändert seit [`MR-012`](done/MR-012-referenzmatrix-grandfathering.md).
  Die Baseline verlangt die Aufwärts-Richtung für **jede** Kante im bindenden Text.
  a-check nimmt die vor der Übernahme `Accepted`-ADRs **0001–0020** per `exempt-paths` ganz aus:
  sie nennen Slice-Kennungen im Körper als Verifikations-Zeiger.
- **Begründung:** dieselbe wie bei [`MR-012`](done/MR-012-referenzmatrix-grandfathering.md) —
  die zwanzig ADRs sind immutabel ([`AGENTS.md`](../../AGENTS.md) §3.5), und die Ausnahme ist
  maschinell kodiert. Dieser Eintrag existiert, weil der Zielabschnitt im Sprung `v6.6.0` →
  `v6.13.0` **nicht wortgleich** blieb — der Zeiger wandert darum nicht still, sondern über
  diesen Nachfolger.
- **Gemessene Abweichung des Zielabschnitts** (`v6.6.0` → `v6.13.0`, 23 242 → 23 083 Zeichen,
  4 Hunks, ohne die Quelle-Zeile): Beispiels-Kennung `slice-NNN` → `slice-tie-break-determinismus`
  (Welle 130/131, *Kennungen sind Namen*); Provenance-Passage verschärft — die Historie-Tabelle
  nennt dort keine ADR und keinen Slice, sondern beim Vertrag den externen CR (Welle 140);
  Kennungs-Form-Passage gekürzt, die Form-Aufzählung steht seitdem in §ID-Schema als Klammer.
  **Kein Hunk berührt die Grandfathering-Aussage** — der Ausgang ist *bleibt gültig*.
- **Offener Adoption-Aspekt:** Welle 138 (*Die Referenzmatrix mechanisch vollständig*) gibt
  Welle-, Carveout- und Roadmap-Kennungen eigene Matrix-Klassen. Ob die `matrix`-Sektion in
  [`.d-check.yml`](../../.d-check.yml) um die neuen Klassen ergänzt werden muss, ist hier **nicht
  entschieden** — ein Folge-Slice benennt und trägt es.
- **Auflösungs-Trigger:** permanent, wie bei [`MR-012`](done/MR-012-referenzmatrix-grandfathering.md).
  Die Grandfather-Menge ist geschlossen und wächst nicht.
- **Löst ab:** [`MR-012`](../conventions.md#mr-012)
- **Ausgelöst durch Baseline-Stand:** `v6.13.0`
