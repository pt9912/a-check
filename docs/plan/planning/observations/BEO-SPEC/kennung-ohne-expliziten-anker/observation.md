# Kennung ohne expliziten Anker faktisch immutabel

**Sub-Area:** Spec-Straten
**Ehemals:** `BEO-029` (Tabellenform, bis slice-139)

Eine Kennungs-Überschrift ohne **expliziten** Anker ist faktisch immutabel: verlinkt eine `Accepted`-ADR ihren generierten Slug, erzeugt jede Umbenennung einen Widerspruch zwischen zwei mandatory Gates — `doc-check` verlangt auflösende Anker, `doc-immutable` verbietet den Nachzug in der ADR
