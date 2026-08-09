# ADR-0033 — `forbidden_constructs` fail-closed statt still wirkungslos

- **Status:** Accepted
- **Datum:** 2026-08-09
- **Autor:** pt9912
- **Bezug:** [AC-FA-RULE-004](../../../spec/lastenheft.md#ac-fa-rule-004--port-disziplin-regel-port-impurity), [AC-FA-CONF-001](../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml), [AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze), [SPEC-CONF-001](../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema), [SPEC-RULE-001](../../../spec/spezifikation.md#spec-rule-001--regel-auswertung), [ADR-0027](0027-constructs-roh-text-monopol.md)
- **Schärft:** [SPEC-CONF-001](../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema) — die Validierung beim Laden der Konfiguration.
- **Supersedes:** —

## Kontext

`forbidden_constructs` wird **je Schicht** konfiguriert, aber nur über die Rolle `port`
ausgewertet ([SPEC-RULE-001](../../../spec/spezifikation.md#spec-rule-001--regel-auswertung)). Der
Config-Adapter reicht den Block dabei **ungeprüft** durch — es gibt für ihn keinerlei Validierung.
Ein Eintrag durchläuft also die ansonsten strenge Dekodierung, wird angenommen und wirkt nie.

Gemessen gegen dieselbe Fixture, variiert wurde nur der Block:

| Fall | Konfiguration | Ergebnis vorher |
|---|---|---|
| **A** Referenz | `ports: ["impl "]` auf `role: port` | `port-impurity`, **Exit 1** — wirkt |
| **B** falsche Rolle | `core: ["impl "]` | **0 Befunde, Exit 0** |
| **C** unbekannte Schicht | `portz: ["impl "]` (Tippfehler) | **0 Befunde, Exit 0** |
| **D** leeres Muster | `ports: [""]` | **0 Befunde, Exit 0** |
| **E** leere Musterliste | `ports: []` | **0 Befunde, Exit 0** |

Der Auslöser war ein realer Konsumenten-Einsatz (2026-08-09) mit Fall **B**: sechs Schichten mit
einem Include-Muster, **0 Befunde** bei vorhandenem Verstoß. **C** ist der härtere Fall — ein
Tippfehler ist nicht einmal eine konzeptuelle Verwechslung, und nichts weist darauf hin.

**Das widerspricht der eigenen Linie in einem einzigen Feature.** Das Schwester-Feature
`constructs` ([ADR-0027](0027-constructs-roh-text-monopol.md)) weist dieselben Fälle hart ab, mit
Fehlertexten, deren Begründung wörtlich auch hier gilt:

```text
a-check: constructs-Muster: leeres pattern unzulässig (es würde nie melden)
a-check: constructs-Muster "impl ": leerer adapter unzulässig (war ein stiller Never-Leak-Eintrag)
```

Ebenso bricht ein unbekannter `languages`-Schlüssel mit Exit 2, und ein leerer `tech.adapter`
ebenso — beides ausdrücklich, um falsch-grüne Konfigurationen zu verhindern. Nur
`forbidden_constructs` ist bisher still.

## Entscheidung

**Jeder Eintrag, der nachweislich nie melden kann, bricht beim Laden mit Exit 2** — vier Ausgänge
einer einzigen, bisher fehlenden Validierung:

1. Die genannte Schicht **existiert nicht** in `layers`.
2. Die Schicht existiert, ihre **effektive Rolle ist nicht `port`** (explizite `role:` oder
   Namens-Inferenz, [AC-FA-RULE-006](../../../spec/lastenheft.md#ac-fa-rule-006--schicht-rollen-generische-regel-anwendung)).
3. Ein **Muster ist leer** — der Never-Match, den `constructs` bereits abweist.
4. Die **Musterliste ist leer**: eine benannte Schicht ohne jedes Muster ist derselbe
   Never-Match, nur eine Ebene höher.

**Der Fehlertext nennt `constructs` als Gegenstück, nicht als Ersatz.** Das ist Teil der
Entscheidung, nicht Kosmetik: die beiden Blöcke sind **komplementär**. `forbidden_constructs` ist
eine **Blacklist je Schicht**, `constructs` eine **Whitelist-Zone, scan-weit**. Wer „Muster in
*einer* Schicht verboten, sonst egal" ausdrücken will, müsste mit `constructs` alle übrigen Zonen
aufzählen und die Liste bei jeder neuen Schicht nachziehen. Ein Fehlertext, der `constructs` als
Ersatz anbietet, schickt den Konsumenten in eine Sackgasse.

**Die Auswertung wird *nicht* ausgeweitet.** Siehe Alternativen — das wäre eine Lastenheft-Frage.

## Konsequenzen

- **Kein Vertragsschnitt nach oben.** Die Validierung setzt durch, was
  [AC-FA-RULE-004](../../../spec/lastenheft.md#ac-fa-rule-004--port-disziplin-regel-port-impurity)
  ohnehin sagt; kein Lastenheft-Bump, keine neue `AC-*`-ID. Der Exit-Code 2 für Konfigurationsfehler
  steht seit jeher in
  [AC-FA-CLI-001](../../../spec/lastenheft.md#ac-fa-cli-001--aufruf-scan-wurzel-und-exit-codes).
- **Breaking Change, bewusst und gemessen klein.** Eine Konfiguration mit einem dieser Einträge
  bricht künftig mit Exit 2 statt still zu schweigen. Von **sieben** lokalen Konsumenten-Konfigurationen
  nutzt **keine** den Block; der einzige bekannte Nutzer ist der Melder des Change Requests, dessen
  Einträge heute nicht wirken — er verliert nichts, was er hätte, und gewinnt die Information, dass
  er es nie hatte.
- **Die Schicht-Blacklist-Lücke bleibt offen — ausdrücklich.** Für „Muster in Schicht X verboten,
  sonst egal" gibt es außerhalb der Rolle `port` weiterhin kein Werkzeug. Diese ADR **benennt** die
  Lücke (im Fehlertext und hier), sie füllt sie nicht. Sie zu füllen hieße,
  [AC-FA-RULE-004](../../../spec/lastenheft.md#ac-fa-rule-004--port-disziplin-regel-port-impurity)
  aufzubohren oder eine neue Anforderung zu schneiden — ein Lastenheft-CR, sobald ein Konsument den
  Bedarf belegt.
- **Fehlermeldungen sind deterministisch.** Die Schicht-Schlüssel werden sortiert geprüft, damit
  eine Konfiguration mit zwei Fehlern immer denselben zuerst meldet
  ([SPEC-DET-001](../../../spec/spezifikation.md#spec-det-001--determinismus-vertrag)).

## Verworfene Alternativen

- **Die Auswertung auf alle Rollen ausweiten** (der zweite Weg des Change Requests). Verworfen —
  und zwar nicht aus Aufwandsgründen, sondern wegen der Quellenlage:
  [AC-FA-RULE-004](../../../spec/lastenheft.md#ac-fa-rule-004--port-disziplin-regel-port-impurity)
  heißt „Port-Disziplin" und nennt wörtlich, dass **Ports** keine implementierungs-spezifischen
  Konstrukte tragen. Die Bindung an `role: port` ist damit die **Einlösung** des Vertrags, nicht
  dessen Verletzung; auch der Befundname `port-impurity` folgt daraus. Eine Ausweitung wäre eine
  Lastenheft-Änderung mit neuer Anforderung und neuem Befundnamen — nach
  [`MR-001`](../../../harness/conventions.md#mr-001--source-precedence-mit-eigener-spezifikations-schicht)
  („ADR darf Spezifikation schärfen, nicht Lastenheft") kann diese ADR das gar nicht leisten.
- **Advisory-Warnung statt Exit 2**, in der Gestalt der drei stderr-Diagnosen
  ([ADR-0029](0029-abdeckungs-diagnose-advisory.md)/[ADR-0031](0031-heuristik-grenzen-diagnose.md)/[ADR-0032](0032-aufloesungs-diagnose-repoweit.md)).
  Verworfen: jene melden **Grenzen der Heuristik** an gültiger Konfiguration — hier liegt ein
  **Konfigurationsfehler** vor, und für den ist Exit 2 die durchgehende Linie dieses Repos
  (unbekannte Sprache, leerer `tech.adapter`, leeres `constructs`-Muster). Zwei Antworten auf
  dieselbe Fehlerklasse wären die eigentliche Inkonsistenz.
- **Nur Fall B beheben** (der Wortlaut des Change Requests). Verworfen: C, D und E entstammen
  derselben fehlenden Validierung an derselben Stelle. Drei von vier Ausgängen zu schließen macht
  den vierten zum Stolperstein.
- **Status quo mit Dokumentation.** Verworfen — eine angenommene, streng dekodierte und dann
  ignorierte Konfiguration ist genau die falsch-grüne Klasse, gegen die
  [AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze) steht.

## Fitness Function

- `make test`: die vier Fälle B–E enden mit **Exit 2** und einer Meldung, die den Grund nennt;
  Fall A bleibt unverändert `port-impurity` mit Exit 1; eine Konfiguration **ohne** den Block ist
  byte-identisch zu vorher.
- **Bestands-Probe:** die sieben lokalen Konsumenten-Konfigurationen laden unverändert.
- `make arch-check` (Dogfooding) und `make ci` grün.

## Geschichte

| Datum | Ereignis |
|---|---|
| 2026-08-09 | Proposed → Accepted (Sign-off Auftraggeber). Auslöser: `CR-4` eines realen Konsumenten-Einsatzes (Fall B). Die Aufbereitung maß drei weitere stille Ausgänge derselben fehlenden Validierung (C, D, E) und klärte die Quellenlage gegen die im CR ebenfalls angebotene Ausweitung. Ab jetzt immutable; Ablösung nur via Folge-ADR mit `Supersedes`. |
