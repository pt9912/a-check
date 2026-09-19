# ADR-0040 — Ein Richtungssegment im Port-Glob schaltet `port-locality` nicht mehr ab

- **Status:** Accepted
- **Datum:** 2026-09-19
- **Autor:** pt9912 (Vorgabe), ausgeführt im Auftrag
- **Bezug:** [AC-FA-RULE-010](../../../spec/lastenheft.md#ac-fa-rule-010--port-lokalität-regel-port-locality),
  [AC-FA-RULE-008](../../../spec/lastenheft.md#ac-fa-rule-008--richtungs-dimension-regel-port-direction-mismatch),
  [ADR-0036](0036-port-richtung-inbound-outbound.md) (die zweite Vokabel),
  [ADR-0029](0029-abdeckungs-diagnose-advisory.md),
  [ADR-0035](0035-grenz-diagnose-gegen-globs.md) und
  [ADR-0032](0032-aufloesungs-diagnose-repoweit.md) — die Diagnosen, in deren Reihe diese tritt
- **Schärft:** [SPEC-RULE-001](../../../spec/spezifikation.md#spec-rule-001--regel-auswertung) und
  [SPEC-CLI-001](../../../spec/spezifikation.md#spec-cli-001--aufruf-scan-wurzel-und-exit-codes) —
  macht die Ableitung des Port-Scopes und diese Diagnose verbindlich.

## Kontext

`port-locality` leitet den **Scope** eines Ports pfad-abgeleitet ab: das längste `port`-Glob-Präfix,
das den Portpfad als Segment-Run trifft, **minus seinem letzten Pfad-Segment** — dem Port-Ordner-
Marker, typisch `ports`. Der Scope gilt nur dann als lokalitätsfähig, wenn er **im App-Baum** liegt
(`appTreeContains`); sonst bleibt die Regel inert, was den klassisch-hexagonalen Geschwister-Ports
ausdrücklich zugesagt ist.

Diese Ableitung setzt **genau ein** Segment zwischen App-Baum und Port-Ordner voraus. Mit
[ADR-0036](0036-port-richtung-inbound-outbound.md) bekam die `port`-Schicht eine zweite Vokabel
(`inbound`/`outbound`), und das Benutzerhandbuch §4 weist an, für `port-direction-mismatch`
**getrennte Port-Schichten** zu deklarieren — trägt eine Schicht genau **eine** Richtung, trägt ihr
Glob naturgemäß das Richtungssegment (`…/createorder/ports/outbound/**`). Damit stehen **zwei**
Segmente zwischen App-Baum und Port-Ordner, und die Ableitung schneidet das falsche ab.

**Gemessen** an einem externen Konsumenten, der diesen Umbau vollzogen hatte, und im Repo
reproduziert — derselbe Baum, derselbe echte Verstoß (eine `app`-Datei der fremden Slice importiert
den slice-lokalen Port), nur das Glob geändert:

| Port-Glob | Ergebnis |
|---|---|
| `…/createorder/ports/**` | `port-locality: 1`, Exit 1 |
| `…/createorder/ports/outbound/**` | `gesamt: 0 Befund(e)`, Exit 0 |
| wie Zeile 2, aber Importziel = fremde **`app`**-Slice (Kontrolle) | `lateral-slice: 1`, Exit 1 |

Die dritte Zeile ist die tragende: Bei **identischer** Config feuert eine andere Regel weiter — die
Config ist nicht tot, sondern `port-locality` ist **still**. Ohne diese Kontrolle wäre „inert" von
„kaputt" nicht zu unterscheiden.

**Die Implementierung ist spec-konform, und das ist der Punkt.** „Minus seinem letzten Pfad-Segment"
stand wörtlich in [SPEC-RULE-001](../../../spec/spezifikation.md#spec-rule-001--regel-auswertung);
die Lücke entstand nicht durch einen Fehler, sondern dadurch, dass eine **zweite** Dimension
denselben Glob-Präfix benutzt. Die beiden Regeln sind einzeln begründet und schalten einander
zusammen ab — die eine lautlos.

## Entscheidung

1. **Die Ableitung zieht ein Richtungssegment zusätzlich ab — wenn es deklariert ist.**
   Endet das Port-Glob auf einem Segment, das dem an **dieser** `port`-Schicht deklarierten
   `direction`-Wert entspricht, fällt dieses Segment zusätzlich zum Port-Ordner-Marker. Das
   Richtungssegment liegt **unter** dem Port-Ordner (`…/ports/outbound/**`), nicht darüber.
   Zwei Leitplanken halten den Schnitt davon ab, den Scope **weiter** zu machen als bisher:
   Er greift nur, wenn der Port **im** App-Baum liegt, und nur dort, wo der Marker-Schnitt
   allein den App-Baum **verfehlt** und der zusätzliche ihn **erreicht**. Eine Aufweichung wäre
   der teurere Fehler — sie machte dieselbe Regel stiller, statt sie zurückzuholen. Diskriminator
   ist der **deklarierte Wert**, nicht eine Wortliste: Die gültige Menge kennt der
   Konfigurations-Validator bereits rollen-abhängig
   ([ADR-0036](0036-port-richtung-inbound-outbound.md)), und eine dritte Quelle für dieselbe
   Menge wäre eine Kopie, die driftet.
2. **Die Restfälle werden laut — mit einem Kriterium, das den legitimen Fall nicht trifft.**
   Eine **weitere** Advisory-Diagnose (stderr, exit-neutral, gedeckelt, mit „Abhilfe") meldet einen
   `port`-Glob, dessen **Verzeichnis im App-Baum liegt**, dessen **abgeleiteter Scope** ihn aber
   nicht erreicht. Das ist die Definition des Defekts: Lokalität wäre hier sinnvoll, und die
   Ableitung hat sie trotzdem verloren. Für Geschwister-Ports (`hex/ports/**` neben
   `hex/services/**`) liegt schon das **Verzeichnis** außerhalb des App-Baums — dort schweigt die
   Diagnose, wie die Regel schweigt. Sie greift damit **breiter** als Punkt 1: auch eine anders
   benannte Richtungsebene und ein Tippfehler im Glob werden sichtbar.
3. **Kein Lastenheft-Bump.** Drei Stellen von
   [AC-FA-RULE-010](../../../spec/lastenheft.md#ac-fa-rule-010--port-lokalität-regel-port-locality)
   sind die mögliche Gegenlesart, und alle drei halten. *„Der Verzeichnis-Teilbaum, der seinen
   Port-Ordner besitzt"* beschreibt weiterhin den Scope. *„Der Scope ist **pfad-abgeleitet** (keine
   Deklaration)"* bleibt wahr: Der Scope entsteht nach wie vor aus dem **Pfad**; `direction` sagt
   nur, wo der Port-Ordner **endet** — die Schicht deklariert ihre Richtung, keine Scope-Grenze.
   Und die out-of-scope gestellte *„erzwungene explizite Scope-Deklaration in der Config"* bleibt
   ausgeschlossen: Wer keine Richtung deklariert, bekommt die alte Ableitung unverändert. Die
   Entscheidung präzisiert eine Ableitung, sie ändert keine Zusage.

**Verworfene Alternative — feste Wortliste (`inbound|outbound`) in der Ableitung.** Sie träfe auch
einen Ordner, der zufällig so heißt, und ließe eine anders benannte Richtungsebene weiter still —
derselbe Fehler in kleiner.

**Verworfene Alternative — Config-Validierung, Exit 2, bei einem Port-Glob auf einem
Richtungssegment.** Sie macht den Konflikt laut, **verbietet** aber die Kombination: Der Konsument
müsste die Richtung an der Port-Schicht wieder aufgeben, und `port-direction-mismatch` bliebe
dauerhaft inert. Das ist genau der Zustand, den die Meldung beklagt — die Regel, die stumm war,
würde durch die Regel ersetzt, die nicht greifen kann.

**Verworfene Alternative — nur dokumentieren.** Sie stellt die angewiesene Konfiguration (getrennte
Port-Schichten) neben eine stumm geschaltete kategorische Regel und nennt das eine Grenze. Die
Grenze wäre echt, aber sie wäre **gewählt**, obwohl die Ableitung sie nicht braucht.

## Konsequenzen

- **Verhaltensänderung an einer kategorischen Regel.** Eine Konfiguration, deren Port-Glob auf einem
  Segment endet, das dem deklarierten `direction`-Wert ihrer Schicht entspricht, kann danach Befunde
  erzeugen, die sie vorher verschwieg. Ein Konsument, der die Kopplung als Grenze dokumentiert hat
  (so geschehen), muss seinen Kommentar nachziehen — die Grenze gibt es nicht mehr.
- **Die Diagnose reiht sich in die bestehenden ein** ([ADR-0029](0029-abdeckungs-diagnose-advisory.md),
  [ADR-0035](0035-grenz-diagnose-gegen-globs.md), [ADR-0032](0032-aufloesungs-diagnose-repoweit.md)).
  Sie erbt deren Hausregeln: eigener Kanal
  (stderr, nach der Zusammenfassung), **Exit-Code unberührt**, Deckel **mit Restzahl**, und ein
  Baum ohne solche Globs bleibt **still** — sonst ist sie Rauschen statt Signal.
- **Der legitime Fall bleibt unberührt, und das ist prüfbar.** Geschwister-Ports sind die eine
  Zusage von [AC-FA-RULE-010](../../../spec/lastenheft.md#ac-fa-rule-010--port-lokalität-regel-port-locality),
  die ausdrücklich inert bleibt; beide Änderungen müssen sie in Ruhe lassen, und ein Test hält das
  fest statt zu hoffen.
- **Die Ableitung liest jetzt zwei Deklarationen** (Glob und `direction`) statt einer. Die Zusage
  ist deshalb nicht „zwei Schnitte in fester Reihenfolge", sondern: Der zusätzliche Schnitt wird
  **nur** genommen, wo er den Scope in den App-Baum bringt. Ein Port-Ordner, der selbst wie eine
  Richtung heißt, ist damit ausgenommen — der Schnitt kann ihn nicht über den App-Baum
  hinausziehen.

## Fitness Function

`make test` — die beiden Glob-Varianten gegen **denselben** Verstoß (rot/grün), die Diagnose mit
ihrer **Umkehr** (Geschwister-Ports bleiben still) und je eine **Mutations-Probe**, die den Befund
zeigt und nicht den Exit-Code. `make gates` (`arch-check` fährt a-checks eigene Konfiguration, die
keine Richtung deklariert — sie belegt die Inertheit der Dimension).
