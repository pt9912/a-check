# ADR-0028 — Unauflösbares Ziel-Glob zieht die Schicht-Zuordnung zurück

- **Status:** Accepted
- **Datum:** 2026-07-25
- **Autor:** pt9912
- **Bezug:** [AC-FA-RULE-005](../../../spec/lastenheft.md#ac-fa-rule-005--schicht-richtung-regel-wrong-direction), [AC-FA-RULE-004](../../../spec/lastenheft.md#ac-fa-rule-004--port-disziplin-regel-port-impurity), [AC-QA-01](../../../spec/lastenheft.md#ac-qa-01--determinismus), [AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze), [SPEC-RULE-001](../../../spec/spezifikation.md#spec-rule-001--regel-auswertung)
- **Schärft:** [SPEC-RULE-001](../../../spec/spezifikation.md#spec-rule-001--regel-auswertung) — die Ziel-Auflösung eines Import-Kandidaten.
- **Supersedes:** —

## Kontext

Die Auflösung eines Import-Ziels auf eine Schicht läuft über den **literalen Präfix** eines
Layer-Globs, gematcht als Segment-Run (modul-präfix-tolerant). Ein Glob mit **Wildcard in der
Mitte** — etwa ein tiefen-agnostisches `…/application/**/ports/**` — hat keinen solchen Präfix:
sein Prefix-String enthält selbst ein Wildcard-Segment und matcht darum **nie** einen realen
Kandidaten. Dass solche Globs als Ziel nicht auflösen, ist seit
[ADR-0026](0026-hexslice-vertical-slice-regeln.md) ausgewiesen.

**Nicht ausgewiesen war die Nebenwirkung.** Der Kandidat verschwindet nicht aus der Auflösung —
er fällt auf das nächst-passende Glob zurück, und das ist typischerweise die **umschließende**
Schicht. Gemessen an einem Fixture (Schicht `ports` mit Glob `…/application/**/ports/**` neben
`application` mit `…/application/**`, deklarierte Kante `outbound → ports`):

```text
src/adapters/outbound/db.go:3: wrong-direction: outbound -> application (…/createorder/ports)
```

Der Adapter importiert einen Port über eine **deklarierte** Kante; gemeldet wird ein Verstoß, den
es nicht gibt. Das wiegt schwerer als eine Lücke: die naheliegende Reparatur ist, die Kante
`outbound → application` zu deklarieren — danach ist das Gate grün und **verdeckt dauerhaft echte**
`outbound → application`-Verstöße. Ein Falsch-Positiv, dessen Behebung ein Falsch-Negativ erzeugt.

Die Quell-Seite (Datei → Schicht) kennt dieses Problem nicht: sie bewertet den literalen **Kopf**
eines Globs und bricht den Gleichstand über die Deklarationsreihenfolge — dort klassifiziert
derselbe verschachtelte Port korrekt. Die beiden Auflösungs-Pfade sind also **asymmetrisch**.

## Entscheidung

Wo ein unauflösbares Glob den Kandidaten **genauso gut decken könnte**, ist seine Schicht nicht
diskriminierbar. Die Zuordnung wird dann **zurückgezogen**: das Ziel gilt als **extern**
(fail-open) — nie als die umschließende Schicht. Dieselbe Linie, die
[ADR-0023](0023-deklarations-bewusste-mehr-wurzel-aufloesung.md) für die schwache Evidenzstufe
zieht: wo die Evidenz die Schicht nicht trennt, wird nicht geraten.

Der Rückzug ist **eng** gefasst. Er greift nur, wenn **alle drei** Bedingungen erfüllt sind —
sonst bleibt die bisherige Zuordnung stehen:

1. Das andere Glob gehört einer **anderen** Schicht als der gewählten.
2. Sein literaler **Kopf** (vor dem ersten Wildcard-Segment) **und** sein literaler **Tail-Marker**
   (nach dem letzten Wildcard-Segment, z. B. `ports`) kommen beide als Segment-Run im Kandidaten
   vor. Der Tail-Marker ist die tragende Bedingung: ohne ihn verschlänge
   `…/application/**/ports/**` **jeden** Application-Import und schaltete die App-Schicht still ab
   — ein weit schwereres False-Green als der Ausgangs-Fehlbefund.
3. Sein Kopf ist **mindestens so spezifisch** wie der Präfix der gewählten Schicht; ein
   spezifischeres literales Glob gewinnt weiterhin.

Ein Glob ohne Tail-Marker (das Prefix endet auf dem Wildcard-Segment, z. B. `a/**/**`) benennt kein
Ziel und zieht nie zurück — nicht-idiomatische Schreibweisen behalten das bisherige Verhalten, wie
schon beim Verzeichnis-Prune ([ADR-0025](0025-exclude-verzeichnis-prune.md)).

**Bewusst nicht entschieden:** die Auflösung *nachzurüsten* (Wildcard-fähiges Ziel-Matching) bleibt
möglich und wäre die vollständigere Lösung. Sie greift jedoch in den Kern-Resolver ein, dessen
Modul-Präfix-Toleranz [ADR-0026](0026-hexslice-vertical-slice-regeln.md) bereits als riskant
verworfen hat; sie braucht ein eigenes Verfahren mit eigener Regressions-Absicherung.

## Konsequenzen

- **Der Fehlbefund verschwindet**, ohne dass eine korrekte Meldung verloren geht: Ziele, die auf
  ein sauberes literales Glob auflösen, sind unberührt; Nicht-Port-Ziele in derselben Schicht
  melden weiter (verifiziert, s. Fitness Function).
- **Eine Kante bleibt ungegatet, aber ehrlich:** Importe in einen Bereich, den nur ein
  Innen-Wildcard-Glob beschreibt, sind extern — kein Befund, keine Falschaussage. Die
  Config-Disziplin aus [ADR-0026](0026-hexslice-vertical-slice-regeln.md) (saubere literale
  Präfixe) bleibt der Weg zum vollen Gate und ist im Benutzerhandbuch ausgewiesen.
- **Inert für den Bestand:** kein lokal geprüfter Konsument nutzt Innen-Wildcard-Globs; für sie
  ändert sich byte-nichts (Regressions-Probe).
- **Kein Lastenheft-Bump:** das *Was* — ein nicht auflösbares Ziel ist extern
  ([AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)) —
  steht bereits; geschärft wird das *Wie*.
- **Determinismus unberührt** ([AC-QA-01](../../../spec/lastenheft.md#ac-qa-01--determinismus)):
  die Prüfung ist eine reine Existenz-Aussage über die Glob-Menge, unabhängig von deren Reihenfolge.

## Verworfene Alternativen

- **Ziel-Matching wildcard-fähig machen:** die vollständigere Lösung, aber ein Eingriff in den
  Segment-Run-Resolver — genau der Umbau, den [ADR-0026](0026-hexslice-vertical-slice-regeln.md)
  als riskant verworfen hat. Bleibt als Folge-Entscheidung offen.
- **Innen-Wildcard-Globs beim Laden verbieten (Exit 2):** fail-closed wäre konsequent, bricht aber
  Konfigurationen, die auf der **Quell**-Seite heute korrekt und nützlich sind (die Datei→Schicht-
  Zuordnung funktioniert mit solchen Globs einwandfrei). Eine Schärfung, die gültige Configs
  ungültig macht, braucht mehr Not als hier vorliegt.
- **Nichts tun, nur dokumentieren:** verlagert die Last auf das Lesen der Doku — und die falsche
  Reparatur (Kante deklarieren) bleibt einen Schritt entfernt und ist unumkehrbar schädlich.

## Fitness Function

- `make test`: der gemessene Fehlbefund verschwindet; vier Gegenproben sichern die Enge des
  Rückzugs — ein Nicht-Port-Ziel derselben Schicht meldet weiter (`wrong-direction`), ein
  kategorischer Befund über dieselbe Schicht meldet weiter (`core-impurity`), ein sauberes
  literales Glob löst weiter auf, und ein spezifischeres literales Glob gewinnt weiter.
- **Konsumenten-Regressions-Probe:** alle lokal geprüften Konsumenten bleiben unverändert.
- `make arch-check` (Dogfooding): unverändert **0** — die Eigen-Config hat keine Innen-Wildcards.
- `make ci` grün.

## Geschichte

| Datum | Ereignis |
|---|---|
| 2026-07-25 | Proposed → Accepted (Sign-off Auftraggeber; Variante „Rückzug statt Nachrüstung" gewählt, nachdem die Messung den Rückfall als Falsch-Positiv mit selbstschädigender Reparatur ausgewiesen hatte). Ab jetzt immutable; Ablösung nur via Folge-ADR mit `Supersedes`. |
