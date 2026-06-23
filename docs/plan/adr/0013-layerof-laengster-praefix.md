# ADR-0013 — `LayerOf` längster-Präfix (Angleichung an `targetLayer`)

- **Status:** Proposed
- **Datum:** 2026-06-23
- **Autor:** pt9912
- **Bezug:** [AC-QA-01](../../../spec/lastenheft.md#ac-qa-01--determinismus) (Determinismus) — der eigentliche Treiber ist jedoch die **Quelle↔Ziel-Konsistenz** der Schicht-Auflösung (siehe Kontext); Determinismus ist hier die schwächere Achse, da der Erst-Treffer bereits deterministisch war.
- **Schärft:** [SPEC-RULE-001](../../../spec/spezifikation.md#spec-rule-001--regel-auswertung) — macht die Schicht-Zuordnung einer Datei (`LayerOf`) konsistent zur Symbol-Auflösung (`targetLayer`).
- **Supersedes:** —

## Kontext

Die Engine löst zwei Schicht-Zuordnungen auf, die über die Wellen auseinanderliefen:

- **`LayerOf`** (eigene Schicht einer Datei) nimmt den **ersten** passenden Glob
  (Reihenfolge der Deklaration).
- **`targetLayer`** (Schicht eines Import-Ziels) nimmt den **längsten** passenden
  Präfix ([ADR-0010](0010-layer-relativer-adapterseg-laengster-praefix.md)).

Bei verschachtelten Schicht-Globs (`src/app/**` ⊂ `src/**`) können beide **abweichen**:
dieselbe Datei wird als **Quelle** anders eingeordnet denn als **Ziel** — eine latente
**Konsistenz**-Lücke (der Erst-Treffer ist bereits deterministisch, aber nicht konsistent
zur Ziel-Auflösung). [ADR-0011](0011-domain-application-trennung-rolle-app.md) hielt dies
als Re-Evaluierungs-Trigger fest.

## Entscheidung

`LayerOf` wählt die **spezifischste** passende Schicht (längster Glob-Präfix), konsistent
mit `targetLayer`, statt des Erst-Treffers; bei Gleichstand die **zuerst deklarierte**.

Das **Match-Prädikat bleibt** `matchesAny` (volle Glob-Semantik für echte Pfade, inkl.
`**`); **nur die Auswahl** unter mehreren Treffern wechselt vom Erst-Treffer auf den
längsten `globPrefix` (Helfer `globPrefix` wiederverwendet). Globs ohne literalen Präfix
(`**/…`) haben Spezifität 0 und verlieren gegen jeden Präfix-Treffer.

## Konsequenzen

- **Verhaltensänderung nur bei verschachtelten Schicht-Globs.** Wo höchstens ein Glob je
  Datei matcht (der Normalfall, u. a. a-checks Eigen-`.a-check.yml`), bleibt das Ergebnis
  identisch → `make arch-check` unverändert grün.
- **Für Globs mit literalem Präfix** klassifizieren Quelle (`LayerOf`) und Ziel
  (`targetLayer`) dieselbe Datei konsistent — keine rollen-abhängige Asymmetrie mehr.
- **Bekannte Restdivergenz:** Globs mit Binnen-Wildcard (`src/*/handlers/**`) löst
  `targetLayer` als **Ziel** gar nicht auf — `globPrefix` liefert einen Präfix mit `*`,
  den `segIndex` in echten Pfaden nie findet —, während `LayerOf`s `matchesAny` sie als
  **Quelle** erfasst; dort bleibt die Asymmetrie bestehen (analog zur Spezifität-0-Klausel
  der `**/…`-Globs). a-check nutzt keine Binnen-Wildcards.
- **Art der Änderung:** weil `f.Layer` aus `LayerOf` stammt und `srcRole = roleOf(f.Layer)`,
  verschiebt der längste-Präfix bei verschachtelten Globs die **anzuwendende
  Reinheits-Regel und die `wrong-direction`-Quelle** — nicht nur das Klassifikations-
  Etikett. Das ist der eigentliche Hebel (und das Risiko) der Umstellung.
- Keine neue Anforderung, kein neuer Befund: reine Engine-/Spec-Konsistenz (Teil B von
  slice-012), deshalb **kein** eigener Versions-Bump des Lastenhefts.

## Fitness Function

- `make arch-check` (Dogfooding, [AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)):
  unverändert 0 Befunde.
- `make test`: verschachtelte Globs (`src/app/**` ⊂ `src/**`) → Datei landet in der
  spezifischsten Schicht; Gleichstand → zuerst deklariert; eine **Mehr-Glob-Schicht**
  (Auswahl über den längsten *matchenden* Glob je Layer — spiegelt `targetLayer`s
  Glob-Schleife); nicht-verschachtelte Config unverändert.

## Geschichte

| Datum | Ereignis |
|---|---|
| 2026-06-23 | Proposed — welle-10b (b2b); `LayerOf` längster-Präfix (Angleichung an `targetLayer`, ADR-0010). |
