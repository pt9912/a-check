# MR-009 — Validator-Rolle unbesetzt, zwei Übergaben ohne Artefakt

- **Datum:** 2026-08-09
- **Geltungsbereich:** gesamtes Repo; betrifft das Baseline-Modul `modul-08` (Rollen und
  Übergaben)
- **Adaption:** Von den neun Rollen-Übergaben der Baseline sind **sieben** mit einem Artefakt
  belegt; **Verifier → Validator** und **Validator → Planner** bleiben unverkörpert. Die
  Validator-Rolle ist in diesem Repo nicht besetzt.
- **Begründung:** Validation fragt „bauen wir das *Richtige*?" (gegen realen Bedarf),
  Verifikation „bauen wir es *richtig*?" (gegen Plan/DoD). a-check verifiziert maschinell,
  validiert aber nicht über ein Artefakt: Rückmeldung der Konsumenten läuft über Issues und
  über die Adoption eines Releases, nicht über einen Validierungsbeleg. Eine Übergabe zu
  *definieren*, die niemand ausführt, wäre eine Zusage ohne Deckung — genau die Klasse
  Harness-Lüge, gegen die `modul-13` steht. Die ausführliche Herleitung steht seit slice-066 in
  [`harness/README.md`](../../README.md#rollen-und-ihre-übergabe-artefakte); dieser Eintrag macht sie zur
  **deklarierten** Abweichung statt zu einer Fußnote.
- **Folgewirkung, ausdrücklich:** die Vollständigkeits-Erwartung von `modul-08` („jedes der neun
  Artefakte") gilt für dieses Repo **ausgewiesen eingeschränkt** auf sieben. Ohne diese Zeile
  bliebe jede Prüfung gegen `modul-08` dauerhaft unvollständig, ohne dass jemand sagen könnte,
  warum — und genau so hat der unabhängige Review vom 2026-08-09 die Lücke als `F-11` gefunden.
- **Gefährlichster Fall, benannt:** *Verifikation grün, Validation rot* — perfekt das Falsche
  gebaut. Dieses Risiko wird durch die Deklaration **nicht kleiner**; es wird nur sichtbar
  getragen statt unsichtbar.
- **Auflösungs-Trigger:** sobald es einen Abnehmer außerhalb des Repos gibt, dessen Urteil in
  einem Artefakt festgehalten wird — etwa ein Konsument, der die Eignung eines Releases für
  seinen Bedarf schriftlich bestätigt, statt sie durch Adoption stillschweigend zu zeigen.
  Beides ist heute nicht vorhanden; der Eintrag ist bis dahin **nicht** permanent, sondern
  begründet ausgesetzt. Gefunden als `B-9` in
  [slice-048](../../../docs/plan/planning/done/welle-12/slice-048-modul-delta-lesen.md), als offen belegt durch
  den [Review-Report vom 2026-08-09](../../../docs/reviews/2026-08-09-welle-12-unabhaengig.md) (`F-11`).
