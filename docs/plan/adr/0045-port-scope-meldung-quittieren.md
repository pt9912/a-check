# ADR-0045 — Die Port-Scope-Meldung wird je Schicht quittiert: `acknowledge` am Schicht-Objekt, kürzer statt stumm

- **Status:** Accepted
- **Datum:** 2026-11-05
- **Autor:** pt9912 (Entscheidung zu GitHub-Issue #7), ausgeführt im Auftrag
- **Bezug:** [AC-FA-CONF-002](../../../spec/lastenheft.md#ac-fa-conf-002--port-scope-meldung-je-schicht-quittieren),
  [AC-FA-CONF-001](../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml),
  [AC-FA-RULE-010](../../../spec/lastenheft.md#ac-fa-rule-010--port-lokalität-regel-port-locality),
  [ADR-0040](0040-portscope-richtungssegment.md) (die Diagnose, die quittiert wird),
  [ADR-0036](0036-port-richtung-inbound-outbound.md) (die Rollen-Vokabel am selben Schicht-Objekt)
- **Schärft:** [SPEC-CONF-001](../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema) und
  [SPEC-CLI-001](../../../spec/spezifikation.md#spec-cli-001--aufruf-scan-wurzel-und-exit-codes) —
  macht Schlüssel, Zeilenform und Fehlerverhalten des Quittierens verbindlich.

## Kontext

Die Port-Scope-Diagnose ([ADR-0040](0040-portscope-richtungssegment.md)) meldet jeden Port-Glob,
dessen Verzeichnis im App-Baum liegt, dessen abgeleiteter Scope ihn aber nicht erreicht. Für einen
Konsumenten mit **app-weiten** Ports ist das der Dauerzustand: Der Scope kann dort nicht unterhalb
des App-Präfixes liegen, `port-locality` kann nicht entscheiden, und die Meldung erscheint in jedem
Lauf. Seit der Hinweis seine Ursache nennt, ist sie als erwartbar ausgewiesen — sie bleibt trotzdem,
und eine Meldung, die jeder Lauf wiederholt, verdeckt die **eine** neue.

Der Maintainer hat den Wunsch des Konsumenten eng gefasst angenommen: **nur** diese Meldung,
**je Schicht**, mit **Pflicht-Begründung**; die Meldung bleibt **sichtbar** und nennt die Schicht;
eine **überflüssige** Quittierung fällt auf; Exit-Code und Regel bleiben unberührt. Offen waren der
Name und die Form des Schlüssels, die Zeilenform, das Fehlerverhalten und der Umgang mit der
Versions-Kopplung des strikten Schemas.

## Entscheidung

1. **Der Schlüssel heißt `acknowledge` und sitzt am Schicht-Objekt.** Er ist eine **Map
   Meldung → Begründung**; zugelassen ist genau ein Schlüssel, `port-scope`, und der Wert ist ein
   nicht leerer String:

   ```yaml
   layers:
     ports:
       globs: ["src/app/ports/**"]
       role: port
       acknowledge:
         port-scope: "app-weite Ports: ein Scope unterhalb des App-Präfixes existiert nicht"
   ```

   Die Schicht ist die Einheit, weil ein Konsument das Rauschen **dort** kennt: Die Begründung steht
   neben den Globs, die sie betrifft, und eine Umbenennung der Schicht nimmt sie mit. Die Map-Form
   hält die Erweiterung auf weitere Meldungen offen, **ohne** sie zuzusagen — heute meldet ein
   einziger Konsument das Rauschen einer einzigen Meldung, und jeder weitere Schlüssel ist bis dahin
   ein Exit 2 wie jeder unbekannte.
2. **Die Meldung wird kürzer, nicht stumm.** An die Stelle der Zeilen je Port-Glob einer quittierten
   Schicht tritt **eine** Zeile:

   ```text
     Schicht ports: quittiert — <Begründung> (<n> Port-Glob(s), Ursache: <Ursache>)
   ```

   Sie nennt Schicht, Begründung, Zahl und Ursache — nicht nur eine Zahl. Die Abhilfe-Zeile einer
   Ursache entfällt, wenn **alle** Schichten mit dieser Ursache quittiert sind. Die Zeilen stehen in
   der Reihenfolge der Schichten; jede Zeile zählt gegen den Deckel der Diagnose.
3. **Eine überflüssige Quittierung wird gemeldet, nicht abgelehnt.** Löst die Schicht die Meldung
   nicht aus — keiner ihrer Port-Globs erscheint in der Diagnose, **vor** der Quittierung gemessen —,
   steht eine eigene Zeile:

   ```text
     Schicht ports: Quittierung ohne Anlass — die Schicht löst die Port-Scope-Meldung nicht aus (Begründung: <Begründung>)
   ```

   Sie ist advisory wie die Diagnose selbst: stderr, Exit-Code unberührt.
4. **Konfigurationsfehler sind Exit 2**, wie jeder Verstoß des strikten Schemas
   ([AC-FA-CONF-001](../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml)):
   `acknowledge` an einer Schicht ohne `role: port` oder in der Listen-Form, ein anderer Schlüssel
   als `port-scope`, eine leere oder fehlende (nur Leerraum gilt als leer) Begründung. Die Meldung
   nennt Schicht und Fehler.
5. **Keine Übergangsform für ältere Versionen.** Ein `a-check`, der den Schlüssel nicht kennt, lehnt
   die Konfiguration mit Exit 2 ab — das ist das bestehende Verhalten des strikten Schemas und gilt
   für jeden neuen Schlüssel. Wer das Image pinnt, hebt den Pin **vor** der Konfiguration; das Handbuch
   sagt es dort, wo der Schlüssel steht.
6. **Die Quittierung betrifft nur die Meldung.** Befunde, Exit-Code und die Scope-Ableitung bleiben
   byte-gleich; ein Test hält denselben Baum mit und ohne Quittierung gegeneinander. **Kein
   Lastenheft-Bump jenseits der neuen AC:** [AC-FA-RULE-010](../../../spec/lastenheft.md#ac-fa-rule-010--port-lokalität-regel-port-locality)
   bleibt unverändert, weil sich an dem, was geprüft wird, nichts ändert.

## Verworfene Alternativen

- **Ein Top-Level-Block `acknowledge: [{layer, notice, reason}]`.** Er führt den Schichtnamen ein
  zweites Mal — ein Tippfehler oder eine Umbenennung lässt die Quittierung ins Leere laufen, und das
  fiele erst als „ohne Anlass" auf. Am Objekt ist die Zuordnung strukturell.
- **Ein Zähler „1 quittiert".** Der Melder hat ihn ausdrücklich ausgeschlossen: Eine Zahl verdeckt,
  welche Schicht schweigt und warum, also genau das, was die Meldung zeigen soll.
- **Ein stummes Quittieren oder ein Schalter für alle Hinweise.** Ein Schweigen verdeckt, was der
  Hinweis sichtbar macht, und ein Schalter für alle Hinweise nähme auch die mit, die niemand
  quittieren wollte.
- **Eine Quittierung je Port-Glob.** Feiner, aber ohne Beleg: Das Rauschen entsteht je Schicht, und
  eine Liste von Globs veraltet bei jeder Umbenennung.
- **Eine überflüssige Quittierung als Exit 2.** Sie bricht die CI an einer Stelle, an der die Meldung
  gerade **verschwunden** ist — etwa weil ein Umbau den Scope behoben hat. Das ist eine Nachricht,
  kein Fehler; die Konfiguration wird beim nächsten Anfassen bereinigt, und die Zeile erinnert daran.
- **Eine Übergangsform** (der Schlüssel als Kommentar, ein Warn-Level für unbekannte Schlüssel). a-check
  kennt kein Warn-Level, und die Einführung eines solchen würde das strikte Schema für **alle**
  Schlüssel aufweichen — derselbe Grund, aus dem die falsche Richtungs-Vokabel seit
  [AC-FA-RULE-008](../../../spec/lastenheft.md#ac-fa-rule-008--richtungs-dimension-regel-port-direction-mismatch)
  kein still akzeptiertes Alias ist.

## Konsequenzen

- **Ein neuer Schlüssel im strikten Schema.** Eine Konfiguration mit `acknowledge` läuft erst ab
  der Version, die ihn kennt; ältere melden Exit 2. Das ist die Versions-Kopplung, die der Melder
  mitbedenken sollte, und sie steht im Handbuch.
- **Eine Quittierung kann einen später entstandenen echten Fehler verdecken**, weil sie für **alle**
  Port-Globs der Schicht gilt. Sie verdeckt ihn nur in der Meldung: Die Zeile bleibt sichtbar und
  nennt die Zahl der Port-Globs; ein zusätzlicher Glob ändert die Zahl, und `port-locality` meldet,
  was es entscheiden kann, unverändert.
- **Die Diagnose bekommt eine Schicht-Zuordnung.** Die Meldung kennt bisher Globs; sie muss die
  Schicht nennen, aus der ein Glob stammt. Das ist eine Erweiterung der Daten, die der Kern liefert,
  keine Änderung der Regel.
- **Die Meldung bleibt gedeckelt** und zählt die neuen Zeilen mit; die Restzahl nennt weiter, was
  nicht gezeigt wurde.

## Fitness Function

`make test` — der Baum mit und ohne Quittierung (Befunde und Exit-Code gleich), zwei Schichten mit
einer Quittierung (je Schicht, nicht insgesamt), die überflüssige Quittierung, jeder
Konfigurationsfehler aus Punkt 4 als Exit 2, und die Determinismus-Probe über zwei Läufe; je eine
Mutations-Probe, die den Befund zeigt und nicht den Exit-Code. `make gates` (`arch-check` fährt
a-checks eigene Konfiguration, die keine Quittierung trägt).
