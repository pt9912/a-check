# MR-013 — ADR-Vorlage ist die vendored Fassung `v5.12.0` (löst [`MR-007`](../conventions.md#mr-007) ab)

- **Status:** Accepted
- **Datum:** 2026-08-29
- **Geltungsbereich:** [`MR-000`](../conventions.md#mr-000) §ID-Schema-Deklaration, Zeile zu
  `ADR-NNNN`
- **Ersetzt-Baseline-Regel:** — *(keine. Dieser Eintrag korrigiert eine **Repo**-Aussage, keine
  Baseline-Regel — siehe Rückbau-Kandidat unten)*
- **Adaption:** [`MR-000`](../conventions.md#mr-000) deklariert `ADR-NNNN` als vierstellig „gemäß
  Kurs-ADR-Vorlage `v1.3.0`". Maßgeblich ist die vendored Vorlage
  `.harness/baseline/v5.12.0/templates/docs/plan/adr/adr.template.md`. **Das Schema selbst ändert
  sich nicht** — vierstellig, chronologisch über den ADR-Index; abgelöst wird ausschließlich die
  Versions-Referenz.
- **Begründung:** [`MR-000`](../conventions.md#mr-000) wird nicht korrigiert — Einträge werden nie
  überschrieben. Eine aktuell **behauptende** Versionsangabe stehenzulassen wäre dagegen eine
  stille Falschaussage. [`MR-007`](../conventions.md#mr-007) hat dasselbe für `v3.5.2` geleistet und ist mit
  diesem Stand fällig geworden; sein eigener Trigger lautete wörtlich „bis zur nächsten
  Baseline-Migration, die ihn ihrerseits ablöst".
- **Rückbau-Kandidat, ausgewiesen:** Nach dem Fork-Test der neuen Baseline — *„ein Eintrag, der
  keine benannte Regel ersetzt, ist ein Fork, keine Adaption"* — gehört dieser Eintrag streng
  genommen nicht in den Adaptions-Block. Er steht hier, weil die Alternative eine falsche
  Versionsangabe wäre. Sauber auflösen lässt er sich erst, wenn die ID-Schema-Deklaration in
  [`MR-000`](../conventions.md#mr-000) selbst überarbeitet wird; die `v5.12.0`-Vorlage legt das
  nahe, indem sie das Schema dort als Default-Liste führt.
- **Auflösungs-Trigger:** die Überarbeitung der ID-Schema-Deklaration — oder die nächste
  Baseline-Migration, je nachdem was zuerst eintritt.
- **Löst auf:** [`MR-007`](../conventions.md#mr-007)
- **Ausgelöst durch Baseline-Stand:** `v5.12.0`
