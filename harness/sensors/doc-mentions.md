# `make doc-mentions` — Erwähnungs-Deckung: die Gegenrichtung des Link-Checks

## Vertrag

Jede Datei unter `harness/sensors/` ist in [`AGENTS.md`](../../AGENTS.md) §4
**genannt**. Das ist die Gegenrichtung zu `make doc-check`: Ein Link-Sensor
prüft, ob ein **genanntes** Ziel existiert; dieser prüft, ob eine
**existierende** Datei genannt wird.

Die Lücke ist keine Erfindung — `v6.5.0` · `templates/harness/README.template.md`
benennt sie wörtlich: *„Seine Grenze: Er prüft EINE Richtung — ob das Ziel
existiert; **eine Datei ohne Index-Zeile** und eine Zeile auf die falsche Datei
bleiben still grün."*

Der Lauf meldet eine **Quote**, nicht nur ein Ja/Nein:
`mentions: 16 von 16 Artefakt(en) erwähnt, über 1 Dokument(e)` — die 16. ist
diese Datei selbst; der Sensor wächtert sich mit..

## Grenze — was das Grün nicht abdeckt

1. **Nur `AGENTS.md` ist Dokument-Menge.** Das Modul sucht den **vollen
   repo-relativen Pfad** als Zeichenkette (gemessen, slice-184).
   [`harness/README.md`](../README.md) nennt dieselben Dateien
   geschwister-relativ (`sensors/<name>.md`) und **kann das nicht ändern**, ohne
   seine Links zu brechen — Markdown-Links sind dateirelativ. Seine
   Sensor-Tabelle bleibt damit ungewächtert.
2. **Die zweite Richtung fehlt weiter:** eine Index-Zeile, die auf die *falsche*
   Datei zeigt, bleibt still grün. Das prüft `doc-check` (Ziel existiert) nur
   für die Existenz, nicht für die Zuordnung.
3. **Erwähnung ≠ Verlinkung.** Alle drei Formen zählen: Markdown-Link,
   Inline-Code und nackte Prosa. Eine Datei, die nur in einem *Zitat* genannt
   wird, gilt als erwähnt.
4. **Die ADRs sind nicht gedeckt**, obwohl dieselbe Frage dort gilt. Der
   ADR-Index verlinkt geschwister-relativ — **39 von 39** —, und das Modul sähe
   jede einzelne als unerwähnt. Diese Richtung trägt weiter die Eigenbau-Prüfung
   (3) in `tools/gate-consistency.sh`; die Ablösung scheitert an der Pfad-Form,
   nicht am Willen (gemessen, slice-184).

## Sperren

- `artifact-unmentioned` — eine Sensor-Datei wird in `AGENTS.md` §4 nicht
  genannt. Entweder fehlt ihre Zeile, oder die Datei ist verwaist.
- `das Modul mentions braucht mentions.artifacts UND mentions.documents` —
  fehlt eine der beiden Listen, bricht der Lauf ab, statt leer durchzulaufen.
- `mentions.artifacts […] trifft kein Artefakt — eine Deckungs-Aussage ueber
  null Mitglieder ist keine`
- `mentions.documents […] trifft kein Dokument — eine Deckungs-Aussage gegen
  null Dokumente ist keine`

**Alle vier Sperren sind fail-closed, und die letzten beiden sind bemerkenswert:**
Das Modul kann **nicht ohne Gegenstand grün melden**. Es ist von Haus aus gegen
genau die Klasse gehärtet, die a-check fünfmal getroffen hat
([`BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf`](../../docs/plan/planning/observations/BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf/observation.md)) —
und für die `make dcheck-phrase-selftest` eine eigene Korpus-Kontrolle bauen
musste. Hier trägt sie das Werkzeug selbst; gemessen mit zwei Leer-Globs
(slice-184).

## Bindung

`DC-FA-MENT-001` · slice-184 (konfiguriert und ins `gates`-Aggregat; Modul
verfügbar seit `d-check v0.75.0`) · Antwort auf die in der Ziel-Form von
`harness/README.md` benannte Grenze.
