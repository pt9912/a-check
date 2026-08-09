# ADR-0032 — Auflösungs-Diagnose: repo-weite Auslösung, schicht-genaue Ausgabe

- **Status:** Accepted
- **Datum:** 2026-08-09
- **Autor:** pt9912
- **Bezug:** [AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze), [AC-FA-CLI-001](../../../spec/lastenheft.md#ac-fa-cli-001--aufruf-scan-wurzel-und-exit-codes), [SPEC-CLI-001](../../../spec/spezifikation.md#spec-cli-001--aufruf-scan-wurzel-und-exit-codes), [SPEC-RULE-001](../../../spec/spezifikation.md#spec-rule-001--regel-auswertung), [ADR-0029](0029-abdeckungs-diagnose-advisory.md), [ADR-0031](0031-heuristik-grenzen-diagnose.md)
- **Schärft:** [SPEC-CLI-001](../../../spec/spezifikation.md#spec-cli-001--aufruf-scan-wurzel-und-exit-codes) — die stderr-Ausgabe eines Scans.
- **Supersedes:** —

## Kontext

Die gefährlichste Konfiguration, die dieses Werkzeug erlaubt: **alle Dateien liegen in Schichten,
alle Symbole werden extrahiert — und trotzdem wird keine einzige Kante beurteilt.** Vollständig
grün, vollständig blind.

Gemessen in einem realen Konsumenten-Einsatz (2026-08-09): ein Mono-Scan über sechs Sprach-Skelette
mit sprach-präfixierten Schicht-Globs (`go/internal/ui/**`) meldete **0 Befunde**, obwohl ein
`service → ui`-Verstoß eingebaut war. Go-Importe tragen den Modulpfad
(`example.com/demo/internal/ui`); das Glob-Literalpräfix `go/internal/ui` kommt darin nicht als
Segment-Run vor, also gilt **jedes** Ziel als repo-extern. Der fail-open-Pfad aus
[SPEC-RULE-001](../../../spec/spezifikation.md#spec-rule-001--regel-auswertung) greift wie
vorgesehen und schweigt.

Die beiden bestehenden Diagnosen erreichen diesen Fall nicht:
[ADR-0029](0029-abdeckungs-diagnose-advisory.md) meldet Dateien **ohne Schicht** — hier lagen alle
Dateien korrekt in Schichten. [ADR-0031](0031-heuristik-grenzen-diagnose.md) meldet Zeilen, deren
**Schreibweise** sie disqualifiziert — hier sind die Zeilen tadellos. Die Lücke sitzt bei den
**Zielen**: syntaktisch einwandfrei, extrahiert, und trotzdem auf nichts auflösend.

Der Change Request formulierte die Regel **je Schicht**:

> „Schicht `<name>`: N Datei(en), 0 von M Import-Symbolen lösen auf eine Schicht auf" — wenn in
> einer Schicht kein einziges extrahiertes Symbol auf irgendeine Schicht auflöst, obwohl Symbole
> extrahiert wurden.

**Diese Regel wurde vor der Entscheidung implementiert und am eigenen Baum gemessen. Sie feuerte:**

```text
gesamt: 0 Befund(e)
Hinweis: Schicht core: 3 Datei(en), 0 von 8 Import-Symbolen lösen auf eine Schicht auf
```

Getroffen hat es ausgerechnet `internal/hexagon/core` — die Schicht, die laut
[`ARC-001`](../../../spec/architecture.md) **abhängigkeitsfrei** sein *muss* und nur die
Standardbibliothek importiert. Dass dort kein Symbol auf eine Schicht auflöst, ist der
Architektur-Erfolg, nicht der Verdachtsfall. Ein reiner Domänen-Kern ist per Konstruktion die
Schicht, auf die die Regel als Erstes zeigt.

Damit ist die Einzel-Schicht-Aussage genau das, was
[ADR-0029](0029-abdeckungs-diagnose-advisory.md) (Entscheidung 3) für die Ziel-Seite bereits
festgestellt hat: „löst auf keine Schicht auf" ist nicht von „zeigt legitim nach außen" zu trennen.

## Entscheidung

**Ausgelöst wird repo-weit, ausgegeben wird je Schicht.**

1. **Auslöse-Bedingung: im gesamten Scan löst kein einziges extrahiertes Symbol auf eine Schicht
   auf, obwohl Symbole extrahiert wurden.** Diese Aussage ist sicher, wo die Einzel-Schicht-Aussage
   es nicht ist: ein Repo mit deklarierten Schicht-Kanten, in dem **nirgends** eine Kante entsteht,
   ist falsch konfiguriert — kein legitimer Baum sieht so aus. Trifft es dennoch zu (ein Repo ohne
   jede interne Kante), dann ist a-check dort als Kanten-Gate tatsächlich wirkungslos, und das zu
   sagen ist der Zweck.
2. **Ausgabe im Wortlaut des Change Requests**, eine Zeile je Schicht mit Symbolen:
   `Schicht <name>: N Datei(en), 0 von M Import-Symbolen lösen auf eine Schicht auf`. Der
   Konsument hätte damit genau die Meldung erhalten, die er angefragt hat — die Abweichung liegt
   allein darin, **wann** sie erscheint.
3. **Schichten ohne extrahierte Symbole werden nicht genannt.** Ein Paket ohne Importe ist normal,
   kein Befund; es in die Liste zu nehmen hieße, Stille als Auffälligkeit zu zeigen.
4. **Advisory**, wie beide Vorgänger-Diagnosen: Exit-Code unberührt, Meldung auch bei null
   Befunden, `composition_root`-Dateien ausgenommen.

## Konsequenzen

- **Kein Vertragsschnitt nach oben.** Kein Lastenheft-Bump, keine neue `AC-*`-ID — advisory
  stderr-Ausgabe ohne Exit-Code-Wechsel, wie
  [ADR-0029](0029-abdeckungs-diagnose-advisory.md)/[ADR-0031](0031-heuristik-grenzen-diagnose.md).
- **Die Diagnose erkennt den Totalausfall, nicht den Teilausfall — ausdrücklich.** Ein Repo, in dem
  vier Schichten sauber auflösen und eine fünfte falsch konfiguriert ist, bleibt **still**. Das ist
  der bewusst gezahlte Preis: die Teilausfall-Regel ist nachweislich falsch-positiv (siehe
  Messung), der Teilausfall selbst bisher unbelegt. **Folge-Trigger:** tritt ein Teilausfall real
  auf, braucht er eine eigene Entscheidung — vermutlich über ein Merkmal, das legitime Reinheit von
  kaputter Auflösung trennt (etwa: die Schicht wird von *anderen* Schichten als Ziel erreicht).
- **Drei stderr-Diagnosen nebeneinander.** Feste Reihenfolge von grob nach fein: Abdeckung (Datei
  ohne Schicht), Grenze (Zeile ohne Kante), Auflösung (Ziel ohne Schicht). Jede ist selbsttragend
  formuliert; ein Baum kann alle drei auslösen.
- **Byte-Identität bleibt.** Deterministisch nach Schichtnamen sortiert.

## Verworfene Alternativen

- **Die CR-Regel je Schicht.** Verworfen **auf Messung, nicht auf Vermutung**: sie feuert auf dem
  eigenen abhängigkeitsfreien Kern. Eine Diagnose, die den Musterfall guter Architektur als
  Verdachtsfall meldet, wird beim ersten Blick weggeschaltet.
- **Heuristik „sieht repo-intern aus"** (Symbol trägt ein Präfix, das im Baum vorkommt): verworfen
  — genau die Unterscheidung, die [ADR-0029](0029-abdeckungs-diagnose-advisory.md) als nicht
  sicher treffbar verworfen hat. Sie würde eine unsichere Aussage in eine Meldung gießen, die
  sicher klingt.
- **Gatend machen (Exit 1).** Verworfen wie in beiden Vorgänger-ADRs: ein Repo, das legitim nach
  außen zeigt, darf nicht rot werden. Die Diagnose ist der Wirkmechanismus, die Strenge wäre die
  Kür.
- **Die Auflösung selbst reparieren.** Dass ein Glob-Literalpräfix nicht im Modulpfad vorkommt, ist
  eine Konfigurationsfrage des Konsumenten (`resolution`), keine Werkzeug-Grenze. Die Diagnose
  sagt ihm, **dass** er sie hat.

## Fitness Function

- `make test`: alle Schichten unauflösbar ⇒ je Schicht eine Zeile; **eine** auflösende Schicht ⇒
  **still**; Schichten ohne Symbole ⇒ still; Exit-Code unverändert; zwei Läufe byte-identisch.
- `make arch-check` (Dogfooding): **still** — der reine Kern darf die Diagnose nicht auslösen.
  Diese Probe ist der Grund für die Entscheidung und damit ihr wichtigster Wächter.
- **Konsumenten-Probe:** das sprach-präfixierte Mono-Repo-Muster meldet alle seine Schichten.
- `make ci` grün.

## Geschichte

| Datum | Ereignis |
|---|---|
| 2026-08-09 | Proposed → Accepted (Sign-off Auftraggeber). Auslöser: `CR-2` eines realen Konsumenten-Einsatzes. Die im CR wörtlich verlangte Regel wurde zuerst gebaut und gemessen; sie feuerte auf dem eigenen abhängigkeitsfreien Kern, was die Auslöse-Bedingung von der Schicht auf den Scan hob. Die Ausgabe blieb im CR-Wortlaut. Ab jetzt immutable; Ablösung nur via Folge-ADR mit `Supersedes`. |
