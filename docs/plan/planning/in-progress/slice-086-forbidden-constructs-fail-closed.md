# slice-086 — `forbidden_constructs` ohne `port`-Rolle: fail-closed statt still

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Deckt:** `CR-4` des Konsumenten-Einsatzes (2026-08-09);
[SPEC-CONF-001](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema),
[AC-FA-RULE-004](../../../../spec/lastenheft.md#ac-fa-rule-004--port-disziplin-regel-port-impurity).
**Bezug:** Welle [`welle-13-konsumenten-befunde`](../welle-13-konsumenten-befunde.md).

---

## 0. Trigger

**Beginn:** sofort.

**Rückführungen:**

- `in-progress` → `open`: falls der Entscheid aus §3 zur Ausweitung statt zum Fehlerfall geht —
  das wäre eine Verhaltensänderung mit eigener Regel-Zeile in
  [SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung) und damit ein
  anderer Slice.

## 1. Auslöser

**Mechanismus: eine Konfiguration wird streng validiert, angenommen — und nie ausgewertet.**

`forbidden_constructs` wird **je Schicht** konfiguriert, aber nur über die Rolle `port`
ausgewertet. Im Code, [`internal/hexagon/core/rules.go`](../../../../internal/hexagon/core/rules.go)
Zeile 34, hängt die gesamte Auswertung an einer Rollen-Bedingung: nur wenn die Schicht die Rolle
`port` trägt, werden die Treffer zu Befunden.

Ein Eintrag für eine `domain`-, `app`- oder Adapter-Schicht durchläuft die strenge Dekodierung,
wird angenommen — und wirkt nie. Gemessen im Konsumenten-Einsatz: **sechs Schichten** mit einem
Include-Muster → **0 Befunde** bei vorhandenem Verstoß.

**Das widerspricht der eigenen Linie.** Dieses Repo lehnt stille Defaults durchgehend ab: ein
unbekannter `languages`-Schlüssel bricht mit Exit 2
([AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml),
0.9.0 — „schließt die stille Nicht-Extraktion (falsch-grün)"), ein leerer `tech.adapter` ebenso
(0.14.0 — „fail-closed statt vormals stillem Never-Leak-Eintrag"). Hier ist derselbe Fall bisher
still.

### Nachtrag 2026-08-09 (Aufbereitung): der Block hat **drei** stille Fälle, nicht einen

Der Config-Adapter reicht die Map **ungeprüft** durch
([`config.go`](../../../../internal/adapter/driven/config/config.go): `Forbidden: yc.Forbidden`) —
es gibt für `forbidden_constructs` **keine** Validierung. Gemessen gegen `a-check:dev`, gleiche
Fixture, nur der Config-Block variiert:

| Fall | Konfiguration | Ergebnis |
|---|---|---|
| **A** Referenz | `ports: ["impl "]` auf `role: port` | `port-impurity`, **Exit 1** — wirkt |
| **B** falsche Rolle (`CR-4`) | `core: ["impl "]` | **0 Befunde, Exit 0** |
| **C** Schicht existiert nicht | `portz: ["impl "]` (Tippfehler) | **0 Befunde, Exit 0** |
| **D** leeres Muster | `ports: [""]` | **0 Befunde, Exit 0** |

**C und D sind neu** und in `CR-4` nicht enthalten. C ist der schlimmere Fall: ein Tippfehler ist
nicht einmal eine konzeptuelle Verwechslung, und nichts weist darauf hin.

**Das Schwester-Feature weist beide hart ab** — mit Fehlertexten, deren Begründung wörtlich auch
hier gilt:

```text
a-check: constructs-Muster: leeres pattern unzulässig (es würde nie melden)
a-check: constructs-Muster "impl ": leerer adapter unzulässig (war ein stiller Never-Leak-Eintrag)
```

[SPEC-CONF-001](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema) nennt für
`constructs` ausdrücklich: *„Anders als bei `tech` ist auch bei `match: substring` ein **leeres**
`pattern` unzulässig (es wäre ein stiller Never-Match)."* Genau dieselbe Aussage ist für
`forbidden_constructs` nie getroffen worden. Der Slice repariert deshalb **eine** fehlende
Validierung mit drei Ausgängen, nicht drei Baustellen.

## 2. Betroffene Module

Zwei Schichten:

1. **Spec** — [`spec/spezifikation.md`](../../../../spec/spezifikation.md)
   ([SPEC-CONF-001](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema): der
   Fehlerfall gehört ins Schema) und die Exit-2-Liste.
2. **`internal/adapter/driven/config/`** — die Validierung beim Laden.

## 3. Auszuführende Gates

`make gates`, `make image-test`.

**Der Entscheid, der vor dem Bau fällt** — der CR nennt beide Wege:

| Weg | Folge |
|---|---|
| **Fail-closed** (CR-Vorschlag) — Eintrag für eine Schicht ohne Rolle `port` ergibt **Exit 2**, mit Verweis auf `constructs` als zonen-gebundenes Gegenstück | passt zur bestehenden Linie; **bricht** aber Konfigurationen, die den Eintrag heute wirkungslos tragen |
| **Auswertung ausweiten** — der Block gilt für alle Rollen | Verhaltensänderung statt Fehlerfall; braucht eine eigene Regel-Zeile in [SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung) und einen neuen Befund-Namen, weil `port-impurity` dann nicht mehr passt |

### Aufbereitung 2026-08-09 — die drei Fragen, die den Entscheid tragen, sind beantwortet

**1. Bestandsmessung: leer.** Von **sieben** lokalen Konsumenten-Konfigurationen (`a-check`,
`d-check`, `d-migrate`, `b-cad`, `m-trace`, `belief-agent`, `grid-gym`) nutzt **keine einzige**
`forbidden_constructs`. Fail-closed bricht im messbaren Bestand **niemanden**. Der einzige bekannte
Nutzer ist der CR-Melder selbst — und dessen Einträge wirken heute nicht, er verliert also nichts,
was er hätte.

**2. Vertragslage: der Block ist per Lastenheft ein *Port*-Werkzeug.**
[AC-FA-RULE-004](../../../../spec/lastenheft.md#ac-fa-rule-004--port-disziplin-regel-port-impurity)
heißt „Port-Disziplin" und sagt wörtlich: *„Ports … tragen — sprachabhängig konfigurierbar — keine
implementierungs-/dialekt-spezifischen Konstrukte (z. B. Rust `impl`)."* Die Bindung an `role: port`
ist damit **keine Nachlässigkeit, sondern die Einlösung des Vertrags**; auch der Befundname
`port-impurity` folgt daraus.

**Das entscheidet Weg 2 mit:** Die Auswertung auszuweiten wäre keine Spezifikations-Schärfung,
sondern eine **Lastenheft-Änderung** — aus einem Port-Werkzeug würde ein allgemeines
Schicht-Werkzeug. Nach
[MR-001](../../../../harness/conventions.md#mr-001--source-precedence-mit-eigener-spezifikations-schicht)
(„ADR darf Spezifikation schärfen, nicht Lastenheft") bräuchte Weg 2 einen Lastenheft-CR und eine
neue `AC-*`-ID. Das ist ein anderer, größerer Slice — nicht dieser.

**3. Ist `constructs` ein Ausweg? Nur teilweise — und der Fehlertext darf nichts anderes
behaupten.** Die beiden Blöcke sind **komplementär, nicht austauschbar**:

| | `forbidden_constructs` | `constructs` |
|---|---|---|
| Logik | **Blacklist**: Muster in Schicht X verboten | **Whitelist**: Muster nur in Zone Y erlaubt |
| Geltung | schicht-gebunden | scan-weit, auch Dateien ohne Schicht |

Wer „Muster in *einer* Schicht verboten, sonst egal" ausdrücken will, müsste mit `constructs`
**alle übrigen Zonen aufzählen** — und die Liste bei jeder neuen Schicht nachziehen, sonst kippt die
Aussage still. **Folge-Befund:** für die Schicht-Blacklist außerhalb von `port` gibt es heute kein
Werkzeug. Das ist eine Lücke im Angebot, kein Grund gegen fail-closed — aber der Fehlertext muss
`constructs` als *Gegenstück* nennen, nicht als Ersatz, sonst schickt er den Konsumenten in eine
Sackgasse.

**Damit ist der Entscheid vorbereitet, nicht vorweggenommen:** Weg 1 (fail-closed) ist der einzige,
der innerhalb des bestehenden Vertrags bleibt, den messbaren Bestand nicht bricht und der eigenen
fail-closed-Linie folgt. Zu entscheiden bleibt beim Bau der **Fehlertext** — er muss die Lücke aus
Punkt 3 benennen, statt sie zu überspielen.

**Negativ-Proben** (A–D wie in §1 gemessen):

| Probe | Erwartung |
|---|---|
| **B** Block auf einer `domain`-Schicht | **Exit 2** mit Nennung der Rolle — in **keinem** Fall stilles Grün |
| **C** Block auf nicht existierender Schicht | **Exit 2**, nennt den unbekannten Schlüssel |
| **D** leeres Muster | **Exit 2**, Wortlaut deckungsgleich mit `constructs` |
| **A** Block auf einer `port`-Schicht | unverändert `port-impurity`, Exit 1 |
| Konfiguration ohne den Block | byte-identisch |
| Die sieben Bestands-Konfigurationen | unverändert grün |

## 4. Was bewusst nicht getan wird

- **`constructs` anfassen.** Der zonen-gebundene Block
  ([AC-FA-RULE-011](../../../../spec/lastenheft.md#ac-fa-rule-011--konstrukt-monopol-regel-construct-leak))
  ist das funktionierende Gegenstück und bleibt unberührt; er ist im Fehlertext zu nennen.
- **Die Rollen-Inferenz ändern.** Dass die Rolle heuristisch aus dem Schicht-Namen abgeleitet wird,
  ist eine andere Frage.
- **Die Schicht-Blacklist-Lücke schließen.** Aus der Aufbereitung (§3 Punkt 3): für „Muster in
  Schicht X verboten, sonst egal" gibt es außerhalb der Rolle `port` kein Werkzeug. Das zu ändern
  hieße, [AC-FA-RULE-004](../../../../spec/lastenheft.md#ac-fa-rule-004--port-disziplin-regel-port-impurity)
  aufzubohren oder eine neue `AC-*` zu schneiden — ein eigener Slice mit Lastenheft-CR, sobald ein
  Konsument den Bedarf belegt. Dieser Slice **benennt** die Lücke im Fehlertext, er füllt sie nicht.

## 5. DoD

- [ ] Der Entscheid aus §3 ist **getroffen und begründet** in der Closure-Notiz — inklusive des
      Fehlertexts, der `constructs` als Gegenstück und nicht als Ersatz benennt.
- [ ] **Keiner** der drei stillen Fälle (B falsche Rolle, C unbekannte Schicht, D leeres Muster)
      endet mehr in grünem Exit 0. Beleg: die Proben aus §3, die vor dem Bau grün waren.
- [ ] `make gates` und `make image-test` grün — **Ausgabe in eine Datei**, Exit-Code getrennt
      geprüft, nie in eine Pipe.

## 6. Closure-Notiz

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice,
[`AGENTS.md`](../../../../AGENTS.md) §5; `make verify` prüft das.)_

## 7. Sub-Area-Modus

Alle berührten Sub-Areas GF.
