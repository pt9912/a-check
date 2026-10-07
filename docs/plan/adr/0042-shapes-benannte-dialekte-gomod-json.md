# ADR-0042: Benannte `shapes`-Dialekte `gomod` und `json` statt konfigurierbarer Lexik

**Status:** Proposed

**Datum:** 2026-10-07

**Autor:** pt9912 (Anweisung und Abnahme), ausgeführt im Auftrag

**Bezug:** [AC-FA-RULE-012](../../../spec/lastenheft.md#ac-fa-rule-012),
[ADR-0041](0041-shapes-sollform-je-datei.md) — **löst deren Punkt 10 ab** („Ein Dialekt:
`kotlin`; der generische Dialekt wartet auf einen zweiten Konsumenten") und **ersetzt für die
Dialekte `gomod` und `json` deren Punkte 3 und 4** (Normalisierung; Fortsetzungsregel) durch die
Punkte 3 und 4 unten. Für `kotlin` gelten 3 und 4 unverändert, die übrigen Punkte für alle
Dialekte; ADR-0041 selbst bleibt unverändert `Accepted`,
[AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)

**Schärft:** [SPEC-CONF-001](../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema)
(Dialekt-Menge),
[SPEC-EXTRACT-001](../../../spec/spezifikation.md#spec-extract-001--import-extraktion)
(Lexik und Zerlegung `gomod`, `json`),
[SPEC-RULE-001](../../../spec/spezifikation.md#spec-rule-001--regel-auswertung)
(einzeilige Meldung aller `shape-*`-Befunde).

**Regeln:** Baseline-Regelwerk `modul-04-adrs.md` §Ziel-Form: ADR (MADR).

---

## Kontext

[ADR-0041](0041-shapes-sollform-je-datei.md) hat `shapes` mit genau einem Dialekt (`kotlin`)
entschieden und einen „generischen Dialekt mit konfigurierbaren Kommentar-, Zeichenketten- und
Trennzeichen" als Folge benannt. Ausgelöst wird diese Folge durch Maintainer-Anweisung, nicht
durch einen belegten zweiten Konsumenten.

**Gemessen** an allen lokalen Repos mit `.a-check.yml`, a-check selbst eingeschlossen: 8 `go.mod`
und 5 `package.json`.

- `go.mod` nutzt nur `module`, `go`, `require`. In 4 Dateien stehen die Abhängigkeiten in
  `require ( … )`-Blöcken, **eine je Zeile**; eine Datei hat eine einzeilige `require`-Direktive,
  drei haben keine. Alle 85 Kommentare sind `// indirect`; keine Backquote-Zeichenketten.
- `package.json` ist **ein** Wurzel-Objekt; Abhängigkeiten stehen in den Mitgliedern
  `dependencies`, `devDependencies`, `peerDependencies`; daneben führt `scripts` beliebige Befehle aus.
- Nebenbei: der ausgelieferte Kotlin-Dialekt zerlegte 82 reale `build.gradle.kts` ohne Fehler.

Zwei Erfahrungen mit dem Kotlin-Dialekt stehen im Beobachtungs-Register: Eine aus dem Gedächtnis
aufgezählte Lexik verschluckte zweimal Code als Kommentar (die eine nicht fail-safe Fehlrichtung),
und die Einzeiligkeit der Befundzeile blieb zweimal unentschieden.

Die Kotlin-Regeln passen auf `go.mod` **nicht**: Dort ist `/* */` **kein** Kommentar, und ein
Zeilenende ist ein **signifikantes Token** (Go-Modulreferenz, Abschnitt *Lexical elements* der
`go.mod`-Dateien) — die Kotlin-Regel „innerhalb von `(` ist ein Zeilenende Leerraum" faltete einen
`require`-Block zu einer Wortfolge.

## Entscheidung

Wir wählen **benannte Dialekte mit fest verankerter, aus der Grammatik abgeleiteter Lexik**:
`gomod` und `json`, neben `kotlin`.

1. **Keine konfigurierbare Lexik.** Die Lexik eines Formats ist eine Aussage über dieses Format,
   nicht über den Konsumenten. Konfigurierbar wäre sie dort zu prüfen, wo sie am wenigsten
   geprüft wird — beim Konsumenten —, und die nicht fail-safe Fehlrichtung (Code als Kommentar)
   wanderte mit.
2. **Abgeleitet, nicht aufgezählt.** Jede Regel der Lexik nennt ihre Quelle in der Grammatik
   (Go-Modulreferenz für `gomod`, RFC 8259 für `json`). Wo die Quelle schweigt oder wo ein
   Werkzeugverhalten nur erinnert wäre, ist der Fall **fail-closed** (Exit 2), nicht geraten —
   ebenso, wo die Quelle eine Form **verbietet** (`/*` in `go.mod`) oder wo zwei verschiedene
   Quellen sonst auf dieselbe Normalform fielen (ein verklebtes `=>`, `;` außerhalb einer
   Zeichenkette, Leerraum zwischen zwei JSON-Zahl-Zeichen). Das Go-Werkzeug selbst ist in
   Einzelheiten großzügiger als die Referenz; der Vertrag folgt der Referenz und bleibt dort, wo
   sie abweichen, fail-closed oder fail-safe.
3. **Einheit `gomod`:** eine Direktive auf oberster Ebene; ein Block `Kopf ( Zeilen )` ist **eine**
   Anweisung, seine Zeilen darin durch `;` getrennt — dieselbe Form wie ein Kotlin-Block. Eine
   zusätzliche Abhängigkeit ändert den Block und ist rot. Anders als bei `kotlin` ist das
   **Zeilenende Grammatik**: eine auf zwei Zeilen verteilte Direktive ist eine andere Folge.
4. **Einheit `json`:** ein **Mitglied des Wurzel-Objekts** (Schlüssel und ganzer Wert). Eine
   zusätzliche Abhängigkeit ändert das Mitglied `dependencies` und ist rot; ein zusätzliches
   `scripts`-Mitglied ebenso. Eine Wurzel, die kein Objekt ist, ist Exit 2. Leerraum außerhalb
   von Zeichenketten entfällt vollständig — JSON braucht ihn nirgends zur Trennung.
5. **Einzeilige, umkehrbare Meldung für alle `shape-*`-Befunde:** Backslash als `\\`, `LF` als
   `\n`, `CR` als `\r`; verglichen wird die unveränderte Anweisung. Umkehrbar muss sie sein, weil
   sonst zwei verschiedene Anweisungen dieselbe Meldung trügen — `shape-differs` zeigte `X
   (erwartet: X)`, und die Zusammenfassung byte-gleicher Befunde verschluckte einen. Aus demselben
   Grund trägt `shape-unused` die Art des Eintrags als Präfix (`literal: ` bzw. `regex: `) statt
   eines Zusatzes, den ein Literal selbst enthalten könnte. Das gilt auch
   für `kotlin`: betroffen sind Anweisungen mit Backslash oder mehrzeiligen Roh-Zeichenketten.
6. **`literal`-Einträge in Quellform:** ein `json`-Eintrag ist ein Mitglied und wird vor dem
   Zerlegen in `{ }` eingeschlossen — so ist die Meldung eines `json`-Befunds **nach dem
   Entmaskieren** als Eintrag übernehmbar; ein `gomod`-Block wird mehrzeilig geschrieben.

## Verglichene Alternativen

| Option | Pro | Contra |
|---|---|---|
| A — nichts tun (nur `kotlin`) | kein neuer Vertrag | keine Prüfung der Go- und Node-Manifeste; der Anlass bleibt offen |
| B — konfigurierbare Lexik (`dialect: generic` mit Zeichen-Schlüsseln) | ein Mechanismus für viele Formate | verlagert die Lexik-Korrektheit zum Konsumenten; die gemessene Fehlerklasse (Code als Kommentar) würde konfigurierbar; `go.mod`-Zeilenenden als Tokens und JSON-Wurzel-Mitglieder ließen sich mit Zeichen-Schlüsseln allein nicht ausdrücken |
| C — Kotlin-Lexik für `go.mod` wiederverwenden | kein neuer Code | falsch: `/* */` ist dort kein Kommentar, Zeilenenden in Blöcken sind Trenner — die Kotlin-Regel faltet einen `require`-Block |
| D — echter JSON-Parser mit Pfad-Ausdrücken (`$.dependencies.*`) | feinere Einheit (einzelne Abhängigkeit) | neues Ausdrucks-Vokabular, eine zweite Zerlegungs-Semantik neben der Anweisung; der gemessene Fall braucht die feinere Einheit nicht |
| **E — benannte Dialekte `gomod`, `json` (gewählt)** | Lexik aus der Grammatik, fail-closed wo sie schweigt; dieselbe Einheit „Anweisung" wie bei `kotlin` | je Format eigener Code; ein weiteres Format braucht eine eigene Entscheidung |

## Konsequenzen

- **Positiv:** Eine zusätzliche Go- oder npm-Abhängigkeit ist rot, gleich in welcher Formatierung,
  und die Lexik ist an einer zitierbaren Quelle prüfbar statt an einer Erinnerung.
- **Negativ:** Ein erlaubter `require`-Block bzw. ein erlaubtes `dependencies`-Mitglied muss
  vollständig genannt werden; jede Versions-Hebung ist eine Listen-Änderung — oder ein Regex-Eintrag
  mit der sicheren Klasse für Zeichenketten-Inhalt.
- **Negativ:** Die einzeilige, umkehrbare Meldung ändert die Ausgabe bestehender `kotlin`-Befunde,
  deren Anweisung einen Backslash oder eine mehrzeilige Roh-Zeichenkette trägt. Wer Ausgaben
  byte-genau vergleicht, sieht die Änderung.
- **Negativ:** In `gomod` ist das Zeilenende Grammatik; ein anders umbrochenes `go.mod` ist eine
  andere Anweisungsfolge — anders als bei `kotlin`, wo Umbrüche das Urteil nicht ändern.
- **Folgepflicht:** Spezifikation, Implementierung, Handbuch und `--print-config`.

## Fitness Function (falls maschinell prüfbar)

| Tooling | Regel | Make-Target |
|---|---|---|
| Go-Tests | je Lexik-Regel ein Test mit Zitat der Grammatik-Stelle; die fail-closed-Fälle je ein Test; Mutations-Gegenprobe je Fix | `make test` |

## Re-Evaluierungs-Trigger

- Ein Konsument braucht ein weiteres Format (XML, TOML) — eigene Entscheidung je Format.
- Ein Konsument braucht eine feinere JSON-Einheit als das Wurzel-Mitglied — dann Option D neu
  bewerten.
- Die Go-Modulreferenz ändert die Lexik der `go.mod`-Dateien.

## Geschichte

| Datum | Ereignis | Verweis |
|---|---|---|
| 2026-10-07 | Proposed | Maintainer-Anweisung, Messung an realen Konsumenten-Manifesten |

Nach `Accepted` wird diese Datei **nicht mehr inhaltlich überschrieben**.
Spätere Korrekturen oder Schärfungen entstehen als neue ADR mit
`Supersedes ADR-NNNN` (Baseline-Regelwerk `modul-04-adrs.md`
§Hard Rule für Accepted-ADRs).
