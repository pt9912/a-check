# MR-002 — ID-Schema mit Bereichskürzeln ab initialer Fassung

- **Datum:** 2026-06-20
- **Geltungsbereich:** [`spec/lastenheft.md`](../../../spec/lastenheft.md), alle Traceability-Verweise
- **Adaption:** Funktionale Anforderungen verwenden von Beginn an
  Bereichskürzel: `AC-FA-<BEREICH>-<NNN>` (z. B.
  [`AC-FA-RULE-001`](../../../spec/lastenheft.md#ac-fa-rule-001--kern-reinheit-regel-core-impurity))
  statt des zweistelligen Kurs-Defaults `<PREFIX>-FA-<NN>`.
  Nichtfunktionale Anforderungen bleiben beim Kurs-Default
  (`AC-QA-<NN>`).
- **Begründung:** Das Lastenheft konsolidiert vier divergente
  Architektur-Checker und trägt von Anfang an mehrere Regel- und
  Funktionsbereiche (`RULE`/`EXTRACT`/`CLI`/`CONF`/`DIST`); eine spätere
  Schema-Migration wäre teurer als ein Bereichsschema ab Welle 1.
- **Auflösungs-Trigger:** permanent.
