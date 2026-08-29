# MR-007 — ADR-Vorlagen-Version: `v3.5.2` statt `v1.3.0`

- **Datum:** 2026-07-25
- **Geltungsbereich:** [`MR-000`](../conventions.md#mr-000)
  §ID-Schema-Deklaration, Zeile zu `ADR-NNNN`
- **Adaption:** [`MR-000`](../conventions.md#mr-000) deklariert
  `ADR-NNNN` als vierstellig „gemäß Kurs-ADR-Vorlage `v1.3.0`". Diese Versionsangabe ist seit
  der Migration überholt; maßgeblich ist die vendored Vorlage
  `.harness/baseline/v3.5.2/templates/docs/plan/adr/adr.template.md`. **Das Schema selbst
  ändert sich nicht** — vierstellig, chronologisch über den ADR-Index; abgelöst wird
  ausschließlich die Versions-Referenz.
- **Begründung:** [`MR-000`](../conventions.md#mr-000) wird **nicht** korrigiert. Die Disziplin des Adaptions-Blocks
  verbietet nachträgliche inhaltliche Änderungen an akzeptierten Einträgen — dieselbe Logik wie
  bei `Accepted`-ADRs ([`AGENTS.md`](../../AGENTS.md) §3.5). Eine aktuell **behauptende**
  Versionsaussage stehenzulassen wäre allerdings eine stille Falschaussage; darum dieser
  ablösende Eintrag. Gefunden als B-7 in
  [slice-048](../../docs/plan/planning/done/slice-048-modul-delta-lesen.md).
- **Auflösungs-Trigger:** permanent — bis zur nächsten Baseline-Migration, die ihn ihrerseits
  ablöst.
