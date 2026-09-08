# `make doc-planning` — Lifecycle-Konsistenz Roadmap ↔ `in-progress/`

## Vertrag

Geprüft wird eine **Äquivalenz**: Liegt ein Slice in `in-progress/`, trägt die
Roadmap-Sektion *Offene Wellen* den Ruhe-Marker **nicht**; ist das Verzeichnis
leer, trägt sie ihn. Beide Richtungen sind derselbe Defekt (Baseline `modul-06`
§Roadmap-Struktur: die Marker-Hälfte ist die deklarierte Redundanz).

## Grenze — was das Grün nicht abdeckt

1. **Kein Name wird geprüft.** Die Äquivalenz ist marker-seitig; *welchen*
   Slice die Sektion nennt — oder ob sie einen nennt —, sieht der Lauf nicht.
   Bis slice-186 sagten [`AGENTS.md`](../../AGENTS.md) §4, `harness/README.md`
   §Sensors und der Vertrag oben *„benennt ihn"* zu, und die Einschränkung stand
   nur als Kommentar in [`.d-check.yml`](../../.d-check.yml). Gemessen hat das
   getragen: Die Roadmap nannte **siebzehn Slice-Übergänge lang** einen Slice,
   der in `done/` lag, und das Gate blieb grün. Seit slice-186 nennt die Sektion
   gar keinen Namen mehr — der Zustand ist die Verzeichnis-Position —, womit die
   Zusage der Prüfung entspricht.
2. **Die Listen-Hälfte** — die Bijektion „genannte Wellen-Kennungen ↔ flache
   Welle-Dateien" prüft der Lauf nicht. Sie hat eine Vorbedingung, die der
   Sensor nicht kennt: das **Kardinalitäts-Modell**. Ein Wächter, der gegen
   *genau eine* Datei hielte, meldete legitime Zustände als Drift. Heilbar,
   aber nicht durch dieses Modul.
3. **Zwei Modul-Fähigkeiten sind bewusst unkonfiguriert** — `closure:` und
   `waves:`. Die Closure-Struktur prüft `doc-structure`; sie hier zusätzlich
   einzuschalten hieße, dieselbe Frage zweimal zu stellen.

**Vorher lief es ohne Gegenstand und meldete grün** — die Lücke ist als
[`BEO-GATE/ruhe-marker-ungewaechtert`](../../docs/plan/planning/observations/BEO-GATE/ruhe-marker-ungewaechtert/observation.md)
registriert und seit slice-122 geschlossen.

## Bindung

`DC-FA-PLAN-001` · slice-122 (konfiguriert und ins Aggregat) · Vertrag an die
geprüfte Äquivalenz angeglichen slice-186 · im `gates`-Aggregat.
