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

`LayerOf` wählt die **spezifischste** passende Schicht (längster **literaler** Präfix),
konsistent mit `targetLayer`, statt des Erst-Treffers; bei Gleichstand die **zuerst
deklarierte**.

Das **Match-Prädikat bleibt** `matchesAny` (volle Glob-Semantik für echte Pfade, inkl.
`**`); **nur die Auswahl** unter mehreren Treffern wechselt vom Erst-Treffer auf den
längsten literalen Präfix. Die Spezifität misst `litPrefixLen` — die **literale
Pfad-Tiefe vor dem ersten Wildcard-Segment**, NICHT die rohe Glob-Stringlänge: ein
Wildcard-Präfix wie `src/*/x` zählt als sein literaler Kopf `src`, ein `**/…`-Glob hat
Spezifität 0. Das spiegelt `targetLayer`, das via `segIndex` ohnehin nur literale
Präfixe auflöst — so kann ein Wildcard-Präfix einen tieferen literalen nie überstimmen.

## Konsequenzen

- **Verhaltensänderung nur bei verschachtelten Schicht-Globs.** Wo höchstens ein Glob je
  Datei matcht (der Normalfall, u. a. a-checks Eigen-`.a-check.yml`), bleibt das Ergebnis
  identisch → `make arch-check` unverändert grün.
- **Spezifität literal-segment-basiert** (`litPrefixLen`): Quelle (`LayerOf`) und Ziel
  (`targetLayer`) gewichten denselben literalen Präfix; ein Wildcard-Präfix kann einen
  tieferen literalen Präfix **nicht** mehr überstimmen (rohe Stringlänge zählt nicht).
  Keine rollen-abhängige Fehlordnung **innerhalb** `LayerOf` mehr.
- **Verbleibende Asymmetrie (nur Match-Prädikat):** `LayerOf` matcht via `matchesAny`
  (volle Glob-Regex) auch Globs mit Wildcard-Präfix als **Quelle**, während `targetLayer`
  solche Schichten via `segIndex` als **Ziel** gar nicht auflöst. Trägt eine Datei *nur*
  ein Wildcard-Präfix-Glob (z. B. `src/*/x/**`), kann sie als Quelle einer Schicht
  zugeordnet sein, die als Importziel nie resolvt — die Spezifität (literaler Kopf)
  gewichtet beide aber gleich. a-check nutzt nur literale `<pfad>/**`-Globs; dort sind
  Quelle und Ziel deckungsgleich.
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
