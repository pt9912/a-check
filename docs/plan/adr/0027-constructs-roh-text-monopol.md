# ADR-0027 — Roh-Text-Konstrukt-Monopol: eigener `constructs`-Block

- **Status:** Accepted
- **Datum:** 2026-07-25
- **Autor:** pt9912
- **Bezug:** [AC-FA-RULE-011](../../../spec/lastenheft.md#ac-fa-rule-011--konstrukt-monopol-regel-construct-leak), [AC-FA-RULE-003](../../../spec/lastenheft.md#ac-fa-rule-003--tech-kapselung-regel-tech-leak), [AC-FA-RULE-004](../../../spec/lastenheft.md#ac-fa-rule-004--port-disziplin-regel-port-impurity), [AC-FA-CONF-001](../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml), [AC-QA-01](../../../spec/lastenheft.md#ac-qa-01--determinismus), [AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)
- **Schärft:** [SPEC-CONF-001](../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema), [SPEC-RULE-001](../../../spec/spezifikation.md#spec-rule-001--regel-auswertung), [SPEC-EXTRACT-001](../../../spec/spezifikation.md#spec-extract-001--import-extraktion), [SPEC-DET-001](../../../spec/spezifikation.md#spec-det-001--determinismus-vertrag), [SPEC-CLI-002](../../../spec/spezifikation.md#spec-cli-002--graph-renderer-vertrag)
- **Supersedes:** —

## Kontext

a-check beurteilt ausschließlich **extrahierte Import-Symbole**
([SPEC-EXTRACT-001](../../../spec/spezifikation.md#spec-extract-001--import-extraktion)). Ein
Konstrukt, das keine Import-Zeile ist, liegt damit strukturell außerhalb: der reale Auslöser ist
ein **Aufruf-Monopol** in einem C++-Konsumenten — `dlopen`/`dlsym`/`dlclose` dürfen nur im
Plugin-Adapter aufgerufen werden. Die Include-Hälfte (`dlfcn.h`) deckt bereits die `tech`-Regel
ab; der **Aufruf** kann jedoch ohne eigenen Include existieren (transitiver Header, lokaler
Prototyp) und bleibt still. Solange das ungeprüft bleibt, lebt die Regel als handgepflegtes
`grep`-Skript im Konsumenten-Repo weiter — genau die Skript-Kopie, deren Ablösung dieses Tool
begründet.

Zwei benachbarte Mechaniken existieren bereits, beide unpassend:

- **`tech`** ([AC-FA-RULE-003](../../../spec/lastenheft.md#ac-fa-rule-003--tech-kapselung-regel-tech-leak))
  hat genau die richtige **Scoping**-Form (Zone als Pfad oder Pfad-Liste, `substring|regex`,
  `composition_root: allow|forbid`), matcht aber auf der falschen Ebene (Import-Symbol).
- **`forbidden_constructs`** ([AC-FA-RULE-004](../../../spec/lastenheft.md#ac-fa-rule-004--port-disziplin-regel-port-impurity))
  matcht auf der richtigen Ebene (Roh-Text), hat aber das falsche Scoping: **schicht**-gebunden,
  nur für die Rolle `port` ausgewertet, mündet in `port-impurity`.

## Entscheidung

1. **Eigener Top-Level-Block `constructs`**, **keine** Generalisierung von
   `forbidden_constructs`. Beide Blöcke matchen Roh-Text, aber mit gegensätzlicher Quantifizierung:
   `forbidden_constructs` sagt „in **dieser Schicht** verboten, sonst egal", `constructs` sagt
   „**nur** in dieser Zone erlaubt, sonst überall Befund". Eine Zusammenlegung müsste beide
   Scopings in einen Schlüssel pressen und die bestehende `port-impurity`-Semantik umdeuten.
2. **Befund `construct-leak`**, Scoping-Mechanik **identisch** zu `tech-leak` — dieselben
   Schlüssel, dieselben Defaults, dieselbe Muster-Sprache (RE2 via `match: regex`,
   [ADR-0015](0015-regex-tech-muster.md)). Keine zweite Muster-Sprache im Vertrag.
3. **Scan-weite Auswertung.** Die Regel gilt für **jede** gescannte Datei, auch für Dateien in
   **keinem** `layers`-Glob — ein Monopol ist eine Aussage über den ganzen Baum, nicht über eine
   Schicht. Die Composition Root ist ausgenommen, sofern der Eintrag nicht `composition_root:
   forbid` deklariert (identisch zu `tech-leak`). `exclude` greift unverändert davor
   ([ADR-0018](0018-exclude-scan-scope.md)/[ADR-0025](0025-exclude-verzeichnis-prune.md)).
4. **Kommentar-Stripping ja** — gematcht wird auf derselben vorbereiteten Quelle wie
   `forbidden_constructs`. Ein `dlopen`-Vorkommen in einem Kommentar ist falsch-**rot**; die
   Divergenz zur `grep`-Referenz (die Kommentare mitsieht) wird **ausgewiesen**
   ([AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)),
   nicht versteckt. String-Literale bleiben die bestehende Grenze.
5. **Treffer tragen den Eintrags-Index, nicht den Mustertext.** Der Extraktions-Adapter liefert je
   Treffer die **Position im `constructs`-Block**; die Regel-Engine entscheidet die Zone am
   Eintrag selbst. Ein Rück-Lookup über den Mustertext wäre bei zwei Einträgen mit gleichem Muster
   und verschiedenen Zonen mehrdeutig — er erzwänge entweder ein künstliches Duplikat-Verbot im
   Schema oder eine Erst-Treffer-Präzedenz, die einen echten Verstoß verschluckt.
6. **Totalordnung der Befunde.** Zwei Muster können dieselbe Zeile derselben Datei treffen; die
   Sortierschlüssel (Pfad, Zeile, Regel) sind dann gleich. Die Ordnung bekommt die **Meldung** als
   letzten Schlüssel — erst damit ist die Ausgabe unabhängig von der internen Reihenfolge
   byte-identisch ([AC-QA-01](../../../spec/lastenheft.md#ac-qa-01--determinismus)). Die Lücke
   besteht latent schon für zwei `forbidden_constructs`-Muster in einer Zeile und wird mit
   geschlossen.
7. **Graph: Legendenzeile, kein Knoten.** Eine Roh-Text-Regel hat keine Kante zwischen Schichten.
   `--print-graph` weist sie in der Legende aus — dem seit
   [SPEC-CLI-002](../../../spec/spezifikation.md#spec-cli-002--graph-renderer-vertrag)
   etablierten Ort für Nicht-Kanten-Semantik. Damit hat die Aussage „bewusst nicht als Kante" eine
   zitierbare normative Stelle statt einer bloßen Behauptung.
8. **Interne Namens-Entflechtung.** Das Modell-Feld, das bisher die
   `forbidden_constructs`-Treffer trägt, hieß `Constructs` — derselbe Name, den der neue
   YAML-Block belegt. Es wird umbenannt, bevor der Block existiert; ein Name, zwei Bedeutungen ist
   eine Verwechslungsfalle in genau dem Code, der beide Pfade unterscheiden muss.

## Konsequenzen

- **Opt-in / rückwärtskompatibel:** ohne `constructs`-Block ändert sich nichts — kein neuer Befund,
  byte-identische Ausgabe. Bestehende Configs bleiben gültig.
- **Vertrags-Ehrlichkeit:** `tech` = extrahierte Import-Symbole, `constructs` = Roh-Text. Beides
  ausgewiesene Heuristik, keine Semantik-Behauptung über den realen Code.
- **Kosten:** die Roh-Text-Prüfung läuft über alle gescannten Zeilen jeder Datei — pro Muster ein
  Substring- bzw. RE2-Lauf je Zeile. RE2 ist linear
  ([SPEC-DET-001](../../../spec/spezifikation.md#spec-det-001--determinismus-vertrag)); der Scan
  bleibt ohne Backtracking-Risiko.
- **Konsumenten-Wirkung:** das b-cad-Aufrufmuster wird zur Config-Zeile; die dortige `grep`-Regel
  ist ablösbar. Die Streichung liegt beim Konsumenten-Repo, nicht hier.
- **Was NICHT entsteht:** kein Zonen-Verbot je Schicht (`forbid_in`) und keine Import-Allowlist —
  beides bewusst außerhalb dieser Entscheidung
  ([AC-FA-RULE-011](../../../spec/lastenheft.md#ac-fa-rule-011--konstrukt-monopol-regel-construct-leak)
  Out-of-Scope).

## Verworfene Alternativen

- **`forbidden_constructs` generalisieren** (Zone statt Schicht, Rolle egal): verworfen — es müsste
  zwei Scoping-Arten in einem Schlüssel tragen und den etablierten `port-impurity`-Befund
  umdeuten. Ein zweiter Block ist billiger als ein überladener erster.
- **Kein Kommentar-Stripping (volle `grep`-Parität):** verworfen — ein Treffer im Kommentar ist
  falsch-rot, und Falsch-Positive kosten ein Gate schneller sein Vertrauen als eine dokumentierte
  Divergenz. Die Divergenz ist in der Paritätsprobe explizit als solche geführt.
- **Muster nur in Layer-Dateien prüfen** (layer-gebunden wie `forbidden_constructs`): verworfen —
  die Evidenz enthält genau den Fall einer Datei außerhalb aller Layer-Globs (Verdrahtungs-`main`),
  in der das Monopol gelten muss. Layer-Bindung wäre still falsch-grün.
- **Eigener Muster-Dialekt (Glob/Wildcard statt RE2):** verworfen — eine zweite Muster-Sprache im
  Konfigurations-Vertrag ist Lernlast ohne Gegenwert.
- **`constructs` als Graph-Knoten/-Badge rendern:** verworfen — der Graph zeigt deklarierte
  Kanten; ein Roh-Text-Monopol als Knoten zu zeichnen behauptete eine Struktur, die es nicht gibt.

## Fitness Function

- `make test`: die Akzeptanzkriterien von
  [AC-FA-RULE-011](../../../spec/lastenheft.md#ac-fa-rule-011--konstrukt-monopol-regel-construct-leak)
  als Tests (Happy, Boundary schichtlos/Composition-Root/Kommentar, Negative fail-closed,
  Determinismus zweier Treffer in einer Zeile) plus Config-Decode-Tests je Fehlerfall.
- **Paritätsprobe** gegen die `grep`-Referenz des Konsumenten: gleiche Treffer, gleiche
  Nicht-Treffer; der Kommentar-Fall ist als deklarierte Divergenz ausgewiesen, nicht als Fehler.
- **Fitness-Probe:** ein injizierter Aufruf außerhalb der Zone erzeugt einen Befund im a-check-Gate
  (statt im Skript); die Gegenprobe innerhalb der Zone bleibt grün.
- `make arch-check` (Dogfooding): unverändert **0** — die Eigen-Config deklariert keinen
  `constructs`-Block.
- `make ci` (inkl. `image-test`): grün.

## Geschichte

| Datum | Ereignis |
|---|---|
| 2026-07-25 | Proposed → Accepted (Sign-off Auftraggeber; alle sieben Design-Entscheide wie vorgelegt abgenommen: eigener Block, Modell-Umbenennung, Befund `construct-leak`, Kommentar-Stripping, scan-weit, RE2, Legendenzeile). Ab jetzt immutable; Ablösung nur via Folge-ADR mit `Supersedes`. |
