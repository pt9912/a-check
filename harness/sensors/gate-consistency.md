# `make gate-consistency` — Meta-Gate gegen Harness-Lügen und Pin-Drift

## Vertrag

Drei Prüfungen, alle drei gegen dieselbe Klasse: eine Deklaration, die etwas
behauptet, was der Bestand nicht trägt.

1. **`.d-check.yml`-Module** — die konfigurierten Module existieren im
   gepinnten Werkzeug. Schutz gegen Gate-Drift: eine Konfiguration, die auf ein
   verschwundenes Modul zeigt, läuft still ins Leere.
2. **Pin-Konsistenz** — Digest-Gleichheit der harten Pins gegen
   `version.md#aktuell`, Version gegen `CHANGELOG`, `d-check.mk`-Tag- und
   Digest-Deklaration. Betroffen sind [`a-check.mk`](../../a-check.mk),
   [`README.md`](../../README.md) und `README.de.md`.
3. **ADR-Index-Vollständigkeit** — jede ADR-Datei ist im Index verlinkt.

## Grenze — was das Grün nicht abdeckt

1. **Prüfung (3) deckt die Gegenrichtung, die `doc-check` per Konstruktion
   nicht sieht.** Ein Link-Sensor prüft, ob ein *genannter* Eintrag existiert;
   ob eine *existierende* Datei genannt wird, sieht er nie. Der Unterschied ist
   der Grund, warum diese Prüfung hier steht und nicht dort.
2. **Prüfung (1) misst Existenz, nicht Wirkung.** Ein Modul kann vorhanden und
   konfiguriert sein und trotzdem nichts prüfen, weil seine Kandidatenmenge
   leer ist. Diese Hälfte trägt `make dcheck-phrase-selftest` — für die zwei
   Muster mit belegtem Ausfall, nicht für alle.
3. **Prüfung (2) vergleicht Deklarationen untereinander, nicht gegen die
   Registry.** Ob der gepinnte Digest im Register noch existiert, sagt erst ein
   Netz-Zugriff; das Gate ist hermetisch und stellt die Frage nicht.

## Bindung

Harness-Prozess ([`AC-QA-02`](../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)
für die Modul-Integrität) · slice-004 (Modul-Prüfung) · slice-018
(Pin-Konsistenz) · slice-087 (ADR-Index) · im `gates`-Aggregat.

Die Teilprüfungen (1)+(2) der **Doku-↔-Makefile-Konsistenz** hat
`make doc-targets` mit slice-079 abgelöst; deren Parität ist in beiden
Richtungen gemessen (slice-073/079).
