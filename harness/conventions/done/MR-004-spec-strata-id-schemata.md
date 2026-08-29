# MR-004 — Spezifikation und Architektur: Strata und ID-Schemata

- **Datum:** 2026-06-21
- **Geltungsbereich:** [`spec/spezifikation.md`](../../../spec/spezifikation.md), [`spec/architecture.md`](../../../spec/architecture.md)
- **Adaption:** Die mit [`MR-001`](../../conventions.md#mr-001)
  angekündigten Ränge 2–3 sind angelegt. Stratum-Platzierung und ID-Schema
  werden hier deklariert (sonst stille Setzung):
  - **Technik-Stratum** = `spec/spezifikation.md`; ID-Präfix
    `SPEC-<BEREICH>-<NNN>` (Bereiche initial `CONF`/`EXTRACT`/`RULE`/`CLI`/
    `DET`/`DIST`). Präzisiert das Lastenheft, erweitert es nie; sprachneutral.
  - **Sicht-Stratum** = `spec/architecture.md`; ID-Präfix `ARC-<NNN>` für
    *Struktur*-Kennungen (Komponenten/Schnittstellen), **keine** eigenen
    Anforderungen; sprach- und meilensteinfrei.
- **Begründung:** Ein Spec-Dokument ohne deklariertes Stratum/ID-Schema ist
  eine stille Setzung (gleiche Harness-Lüge-Klasse wie ein undeklariertes
  Gate) und nicht normativ zitierbar. Die Schemata spiegeln die
  Bereichskürzel des Lastenhefts
  ([`MR-002`](../../conventions.md#mr-002))
  für durchgängige Traceability.
- **Auflösungs-Trigger:** permanent. Die maschinelle Kennungs-Linkpflicht
  (`ids`-Muster für `SPEC-*`/`ARC-*` in [`.d-check.yml`](../../../.d-check.yml))
  folgt mit dem Implementierungs-Slice, der die übrigen `d-check`-Module
  aktiviert (konsistent mit der dortigen Deferred-Politik).
