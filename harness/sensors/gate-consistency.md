# `make gate-consistency` — Meta-Gate gegen Harness-Lügen und Pin-Drift

## Vertrag

**Vier** Prüfungen, alle gegen dieselbe Klasse: eine Deklaration, die etwas
behauptet, was der Bestand nicht trägt. Die Nummerierung im Skript beginnt bei
(3) — (1) und (2) sind mit slice-079 an `make doc-targets` abgegeben, und die
Nummern bleiben frei, damit die Verweise darauf gültig bleiben.

1. **`.d-check.yml`-Module** — die `modules`-Liste führt die aktiven Module
   (`links`/`anchors`/`ids`/`matrix`) **und nicht `external`**. Die zweite
   Hälfte ist die eigentliche Zusage: `external` ist eine Netzwerk-Tür, und der
   netzlose `doc-check` verlöre mit ihr still seine Beweis-Aussage für
   [`AC-QA-02`](../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze).
2. **Pin-Konsistenz** — Digest-Gleichheit der harten Pins gegen
   `version.md#aktuell`, Version gegen `CHANGELOG`, `d-check.mk`-Tag- und
   Digest-Deklaration.
3. **`.PHONY`-Vollständigkeit** — jedes eigene Rezept-Target ist in einer
   `.PHONY:`-Zeile deklariert. Ohne sie überspringt `make` das Rezept, sobald
   eine gleichnamige Datei existiert, und meldet **Exit 0** — ein Gate, das
   nicht läuft und grün sagt. Bewusst ausgenommen sind Targets, die tatsächlich
   eine gleichnamige Datei erzeugen.
4. **ADR-Index-Vollständigkeit** — jede ADR-Datei ist im Index verlinkt.

Vor der eigentlichen Prüfung laufen **Selbsttests**: ein Phantom-Target muss
das Gate nachweislich feuern lassen, und eine `.PHONY`-Lücke ebenso.

## Grenze — was das Grün nicht abdeckt

1. **Prüfung (4) deckt die Gegenrichtung, die `doc-check` per Konstruktion
   nicht sieht.** Ein Link-Sensor prüft, ob ein *genannter* Eintrag existiert;
   ob eine *existierende* Datei genannt wird, sieht er nie.
2. **Prüfung (1) misst Präsenz und Abwesenheit von Namen, nicht Wirkung.** Ein
   Modul kann in der Liste stehen und nichts prüfen, weil seine Kandidatenmenge
   leer ist. Diese Hälfte trägt `make dcheck-phrase-selftest` — für das eine
   Muster mit belegtem Ausfall, nicht für alle.
3. **Prüfung (2) vergleicht Deklarationen untereinander, nicht gegen die
   Registry.** Ob der gepinnte Digest dort noch existiert, sagt erst ein
   Netz-Zugriff; das Gate ist hermetisch und stellt die Frage nicht.
4. **Die Doku-↔-Makefile-Konsistenz prüft dieses Gate nicht mehr** — sie liegt
   seit slice-079 bei `make doc-targets`, dessen Parität in beiden Richtungen
   gemessen ist (slice-073/079).

## Sperren

- `Target '<name>' fehlt in .PHONY` — eine gleichnamige Datei ließe `make` das
  Rezept überspringen und Exit 0 melden.
- `.d-check.yml modules ohne '<modul>'` — der netzlose `doc-check` beweist
  [`AC-QA-02`](../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze) nur mit den aktiven Modulen.

## Bindung

Harness-Prozess ([`AC-QA-02`](../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)
für die Modul-Integrität) · slice-004 (Modul-Prüfung) · slice-018
(Pin-Konsistenz) · slice-087 (ADR-Index) · im `gates`-Aggregat.
