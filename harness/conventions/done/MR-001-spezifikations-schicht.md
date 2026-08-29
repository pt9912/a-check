# MR-001 — Source Precedence mit eigener Spezifikations-Schicht

- **Datum:** 2026-06-20
- **Geltungsbereich:** [`harness/README.md` §Source precedence](../../README.md#source-precedence)
- **Adaption:** Die Source-Precedence-Tabelle führt
  `spec/spezifikation.md` als eigenen **Rang 2** zwischen Lastenheft
  (Rang 1) und Architektur (Rang 3). Der Kurs-Default setzt zwei
  Spec-Ränge; dieses Repo nutzt drei. Die Dateien der Ränge 2–3 sind mit
  slice-002 angelegt und in den Tabellen verlinkt; Stratum- und
  ID-Schema-Deklaration: [`MR-004`](../../conventions.md#mr-004).
- **Begründung:** Spec-Stratifizierung mit drei Spec-Dateien; die
  ADR-Schärfungs-Regel („ADR darf Spezifikation schärfen, nicht
  Lastenheft") soll strukturell sichtbar sein. Konsistent mit dem
  Schwester-Repo `d-check`.
- **Auflösungs-Trigger:** permanent.
