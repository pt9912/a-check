# MR-016 — Validator-Rolle unbesetzt (schärft [`MR-009`](../conventions.md#mr-009))

- **Status:** Accepted
- **Datum:** 2026-08-29
- **Geltungsbereich:** gesamtes Repo; betrifft das Baseline-Modul `modul-08`
- **Ersetzt-Baseline-Regel:** [`modul-08-agentenrollen.md` §Die neun Übergaben und ihre Artefakte](../../.harness/baseline/v6.0.0/regelwerk/modul-08-agentenrollen.md#die-neun-übergaben-und-ihre-artefakte-modul-8)
- **Adaption:** Die **Validator-Rolle ist in diesem Repo nicht besetzt**. Von den neun
  Rollen-Übergaben sind sieben mit einem Artefakt belegt; *Verifier → Validator* und
  *Validator → Planner* bleiben unverkörpert.
- **Begründung:** Validation fragt „bauen wir das *Richtige*?" gegen den realen Bedarf. a-check
  verifiziert maschinell, validiert aber nicht über ein Artefakt: Rückmeldung der Konsumenten
  läuft über Issues und über die Adoption eines Releases. Eine Übergabe zu *definieren*, die
  niemand ausführt, wäre eine Zusage ohne Deckung.
- **Was mit diesem Stand **enger** wird:** `modul-08` sagt in `v5.12.0` **selbst**, dass die
  beiden Validator-Kanten *„nicht in jeder Sequenz laufen"* und ihr Beleg **repo-extern** ist
  (Artefaktklasse *keins*) — ein Satz, den `v3.5.2` nicht kennt. Damit ist die *Abwesenheit eines
  Artefakts* kein Verstoß mehr. Was bleibt, ist die Rolle selbst: sie ist hier **gar nicht**
  besetzt, auch nicht nach einem MVP-Slice.
- **Gefährlichster Fall, benannt:** *Verifikation grün, Validation rot* — perfekt das Falsche
  gebaut. Die Deklaration macht das Risiko nicht kleiner, nur sichtbar.
- **Auflösungs-Trigger:** sobald ein Abnehmer außerhalb des Repos sein Urteil über die Eignung
  eines Releases in einem Artefakt festhält, statt es durch Adoption stillschweigend zu zeigen.
- **Löst auf:** [`MR-009`](../conventions.md#mr-009)
- **Ausgelöst durch Baseline-Stand:** `v5.12.0`
