# Eine Anforderung legt eine Ausgabe fest, ohne den bestehenden Ausgabecode gelesen zu haben

**Sub-Area:** Spec-Straten

Eine neue Anforderung oder ADR erweitert eine vorhandene Meldung (Zeilenform, Reihenfolge, Deckel,
Kopfzeile, Randfälle) und schreibt sie aus der Beschreibung des Wunsches, nicht aus dem Code, der
sie heute erzeugt. Was der Code schon entscheidet — Sortierung, frühe Rückkehr, Entdoppelung,
Rollen-Inferenz —, bleibt unbenannt, und der Umsetzungs-Slice müsste es raten.
