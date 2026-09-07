# `make doc-planning` — Lifecycle-Konsistenz Roadmap ↔ `in-progress/`

## Vertrag

Liegt ein Slice in `in-progress/`, muss die Roadmap-Sektion *Offene Wellen* ihn
**benennen**, statt den Ruhe-Marker „Nichts in Arbeit." zu tragen — und
umgekehrt. Beide Richtungen sind derselbe Defekt (Baseline `modul-06`
§Roadmap-Struktur: die Marker-Hälfte ist die deklarierte Redundanz).

## Grenze — was das Grün nicht abdeckt

1. **Die Listen-Hälfte** — die Bijektion „genannte Wellen-Kennungen ↔ flache
   Welle-Dateien" prüft der Lauf nicht. Sie hat eine Vorbedingung, die der
   Sensor nicht kennt: das **Kardinalitäts-Modell**. Ein Wächter, der gegen
   *genau eine* Datei hielte, meldete legitime Zustände als Drift. Heilbar,
   aber nicht durch dieses Modul.
2. **Zwei Modul-Fähigkeiten sind bewusst unkonfiguriert** — `closure:` und
   `waves:`. Die Closure-Struktur prüft `doc-structure`; sie hier zusätzlich
   einzuschalten hieße, dieselbe Frage zweimal zu stellen.

**Vorher lief es ohne Gegenstand und meldete grün** — die Lücke ist als
[`BEO-GATE/ruhe-marker-ungewaechtert`](../../docs/plan/planning/observations/BEO-GATE/ruhe-marker-ungewaechtert/observation.md)
registriert und seit slice-122 geschlossen.

## Bindung

`DC-FA-PLAN-001` · slice-122 (konfiguriert und ins Aggregat) · im
`gates`-Aggregat.
