# ADR-0029 — Abdeckungs-Diagnose: ungedeckte Dateien ausweisen, nicht gaten

- **Status:** Accepted
- **Datum:** 2026-07-25
- **Autor:** pt9912
- **Bezug:** [AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze), [AC-QA-01](../../../spec/lastenheft.md#ac-qa-01--determinismus), [AC-FA-CLI-001](../../../spec/lastenheft.md#ac-fa-cli-001--aufruf-scan-wurzel-und-exit-codes), [SPEC-CLI-001](../../../spec/spezifikation.md#spec-cli-001--aufruf-scan-wurzel-und-exit-codes), [SPEC-DET-001](../../../spec/spezifikation.md#spec-det-001--determinismus-vertrag)
- **Schärft:** [SPEC-CLI-001](../../../spec/spezifikation.md#spec-cli-001--aufruf-scan-wurzel-und-exit-codes) — die stderr-Ausgabe eines Scans.
- **Supersedes:** —

## Kontext

Ein Konsument, dessen `languages`-Globs weiter reichen als seine `layers`-Globs, hat gescannte
Dateien in **keiner** Schicht. Das ist zulässig und oft gewollt — aber es erzeugt eine stille
Lücke: die Schicht-Regeln greifen für diese Dateien nicht, und ein Import **auf** sie löst auf
keine Schicht auf und bleibt fail-open unbeurteilt
([AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze), bewusste
Grenze). Nichts in der Ausgabe weist darauf hin. Ein grünes Gate über einem teilweise ungeprüften
Baum sieht aus wie ein grünes Gate über einem geprüften.

Der scharfe reale Fall: ein Konsument führt ein Verzeichnis in einer `tech`-Regel ausdrücklich als
Architektur-Zone (zulässiger Datenbank-Halter), hat ihm aber keine Schicht gegeben. Ein Import aus
der Domäne dorthin wäre exakt der Verstoß, den das Gate finden soll — und bliebe still.

**Die Evidenz ist klein.** Eine Nachmessung über sieben lokale Konsumenten-Konfigurationen mit
exakter Glob-Semantik ergab: sechs sind vollständig gedeckt, einer hat **zwei** Dateien. Eine
frühere Messung hatte einen zweiten Konsumenten genannt; dort war ein `exclude`-Block übersehen
worden. Die Fehlerklasse ist real, ihre heutige Verbreitung gering.

## Entscheidung

**Die Grenze wird ausgewiesen, nicht verschoben** — das ist wörtlich der Vertrag von
[AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)
(„die Grenze wird dokumentiert statt verschwiegen").

1. **Advisory-Diagnose auf stderr.** Ein Scan, der Dateien in keiner Schicht findet, nennt sie
   nach der Befund-Zusammenfassung. **Der Exit-Code bleibt unberührt** — ein sauberes Repo mit
   ungedeckten Dateien endet weiter mit 0.
2. **Kein Config-Schlüssel, keine Strenge.** Ein opt-in `strict_coverage` (ungedeckt ⇒ Exit 1)
   wird **nicht** eingeführt. Er nützt erst, *nachdem* die Diagnose den Konsumenten auf die Lücke
   gestoßen hat; die Diagnose ist der Wirkmechanismus, die Strenge die Kür. Ein Default-Exit-1
   wäre ohnehin unzulässig — er bräche bestehende grüne Gates ohne ADR
   ([`AGENTS.md` §3.6](../../../AGENTS.md#36-gates-dürfen-nicht-ohne-adr-gelockert-werden) in der
   Gegenrichtung).
3. **Nur die Quell-Seite.** Gezählt werden **gescannte Dateien** ohne Schicht. Die Ziel-Seite
   (ein Import, der auf keine Schicht auflöst) bleibt außen vor: dort ist „schichtlos im eigenen
   Baum" nicht sicher von **repo-extern** (Fremdbibliothek) trennbar, und eine Meldung darüber
   wäre Rauschen. Die Ursache wird über die Quell-Seite mit repariert.
4. **Abgrenzung.** `composition_root`-Dateien zählen **nicht** — sie sind bestimmungsgemäß
   schichtlos. `exclude`-Dateien erreichen den Scan ohnehin nie
   ([ADR-0018](0018-exclude-scan-scope.md)/[ADR-0025](0025-exclude-verzeichnis-prune.md)).
5. **Datei-Liste, keine Zonen-Aggregation.** Stabil nach Pfad sortiert
   ([SPEC-DET-001](../../../spec/spezifikation.md#spec-det-001--determinismus-vertrag)). Ab **zehn**
   Dateien wird gekürzt und die **Restzahl ausgewiesen** — eine Kürzung, die ihre eigene Größe
   nennt, ist keine stille Kappung. Eine Aggregation zu „Zonen" wurde erwogen und verworfen: sie
   war mit einem Konsumenten mit hunderten Dateien begründet, den es nicht gibt.

## Konsequenzen

- **Kein Vertragsschnitt nach oben.** Kein Lastenheft-Bump, keine neue `AC-*`-ID: eine advisory
  stderr-Zeile ohne Exit-Code-Wechsel schärft das *Wie* der bestehenden Ausgabe
  ([AC-FA-CLI-001](../../../spec/lastenheft.md#ac-fa-cli-001--aufruf-scan-wurzel-und-exit-codes)
  nennt „Zusammenfassung auf stderr"). Präzedenz: [ADR-0025](0025-exclude-verzeichnis-prune.md),
  [ADR-0028](0028-ziel-glob-schattenwurf.md).
- **Nur ein Konsument sieht heute etwas.** Sechs von sieben Konfigurationen bleiben
  diagnose-**frei** — das ist die Bedingung dafür, dass die Meldung Signal und nicht Rauschen ist.
- **Byte-Identität bleibt.** Die Diagnose ist deterministisch sortiert; der nativ↔Container-Vergleich
  der Distributions-Akzeptanz (stdout **und** stderr) hält.
- **Der Konsument behält die Wahl:** Schicht deklarieren *oder* `exclude` ergänzen. a-check
  entscheidet nicht, was zur Architektur gehört — es sagt nur, worüber es nichts aussagt.

## Verworfene Alternativen

- **Opt-in `strict_coverage` mitliefern** (Gestalt (d) des Entwurfs): verworfen für **jetzt**, nicht
  grundsätzlich. Kostet das 2–2,5-fache, fast ausschließlich Vertrags-Maschinerie (neue AC-ID,
  Config-Schlüssel mit fail-closed-Validierung, Exit-Code-Semantik) — für eine Strenge, die noch
  niemand angefragt hat. Folge-Entscheidung, sobald ein Konsument sie will.
- **Ziel-seitige Strenge** (Import ohne auflösbare Schicht ⇒ Befund): verworfen — nicht von
  repo-externem Code trennbar, Geister-Match-Risiko.
- **Zonen-Aggregation statt Datei-Liste:** verworfen, siehe Entscheidung 5.
- **Eigenes Flag (`--print-coverage`):** verworfen — eine Diagnose, die man anfordern muss, sieht
  genau der nicht, der sie braucht.

## Fitness Function

- `make test`: Diagnose erscheint bei ungedeckten Dateien und **fehlt** bei vollständiger
  Abdeckung; `composition_root` zählt nicht; Exit-Code unverändert (auch bei 0 Befunden);
  Kürzung ab zehn Dateien nennt die Restzahl; zwei Läufe byte-identisch.
- **Konsumenten-Probe:** die vollständig gedeckten Konfigurationen bleiben diagnose-frei; der
  betroffene Konsument nennt genau seine ungedeckten Dateien.
- `make arch-check` (Dogfooding): unverändert **0** und diagnose-frei — der Eigen-Baum ist gedeckt.
- `make ci` grün.

## Geschichte

| Datum | Ereignis |
|---|---|
| 2026-07-25 | Proposed → Accepted (Sign-off Auftraggeber; Gestalt „nur Diagnose" gewählt, die Opt-in-Strenge ausdrücklich vertagt, nachdem die Nachmessung die Evidenz auf einen Konsumenten mit zwei Dateien verkleinert hatte). Ab jetzt immutable; Ablösung nur via Folge-ADR mit `Supersedes`. |
