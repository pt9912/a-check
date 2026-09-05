# MR-017 — ADR-Vorlage ist die vendored Fassung `v6.0.0` (löst [`MR-013`](../conventions.md#mr-013) ab)

- **Status:** Accepted
- **Datum:** 2026-09-05
- **Geltungsbereich:** [`MR-000`](../conventions.md#mr-000) §ID-Schema-Deklaration, Zeile zu
  `ADR-NNNN`
- **Ersetzt-Baseline-Regel:** — *(keine. Dieser Eintrag korrigiert eine **Repo**-Aussage, keine
  Baseline-Regel — derselbe Rückbau-Kandidat wie sein Vorgänger, siehe unten)*
- **Adaption:** [`MR-000`](../conventions.md#mr-000) deklariert `ADR-NNNN` als vierstellig „gemäß
  Kurs-ADR-Vorlage `v1.3.0`". Maßgeblich ist die vendored Vorlage
  `.harness/baseline/v6.0.0/templates/docs/plan/adr/adr.template.md`. **Das Schema selbst ändert
  sich nicht** — vierstellig, chronologisch über den ADR-Index; abgelöst wird ausschließlich die
  Versions-Referenz.
- **Begründung:** [`MR-000`](../conventions.md#mr-000) wird nicht korrigiert — Einträge werden nie
  überschrieben. Eine aktuell **behauptende** Versionsangabe stehenzulassen wäre dagegen eine
  stille Falschaussage. [`MR-013`](../conventions.md#mr-013) hat dasselbe für `v5.12.0` geleistet
  und ist mit diesem Stand fällig geworden; sein eigener Trigger lautete wörtlich „die nächste
  Baseline-Migration, die ihn ihrerseits ablöst" — die ist mit [slice-135](../../docs/plan/planning/done/slice-135-regelwerk-v600-delta-analyse.md)
  eingetreten. **Gemessen, nicht angenommen:** `docs/plan/adr/adr.template.md` ist zwischen
  `v5.12.0` und `v6.0.0` byte-identisch (`git diff v5.12.0 v6.0.0 -- lab/templates/docs/plan/adr/adr.template.md`,
  leerer Diff) — dieser Eintrag korrigiert ausschließlich die Versions-**Referenz**, keine
  inhaltliche Abweichung.
- **Rückbau-Kandidat, ausgewiesen:** Nach dem Fork-Test der Baseline — *„ein Eintrag, der keine
  benannte Regel ersetzt, ist ein Fork, keine Adaption"* — gehört dieser Eintrag streng genommen
  nicht in den Adaptions-Block. Er steht hier, weil die Alternative eine falsche Versionsangabe
  wäre. Sauber auflösen lässt er sich erst, wenn die ID-Schema-Deklaration in
  [`MR-000`](../conventions.md#mr-000) selbst überarbeitet wird — derselbe Zustand wie bei
  [`MR-013`](../conventions.md#mr-013), unverändert durch diesen dritten Durchlauf. Zwei
  aufeinanderfolgende Baseline-Migrationen haben ihn jetzt beide fällig werden lassen, ohne dass
  die ID-Schema-Überarbeitung eingetreten wäre — ein Muster, das für sich genommen ein
  Beobachtungs-Register-Eintrag wäre, träte es ein drittes Mal auf.
- **Auflösungs-Trigger:** die Überarbeitung der ID-Schema-Deklaration — oder die nächste
  Baseline-Migration, je nachdem was zuerst eintritt.
- **Löst auf:** [`MR-013`](../conventions.md#mr-013)
- **Ausgelöst durch Baseline-Stand:** `v6.0.0`
