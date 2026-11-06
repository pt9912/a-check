# Review — slice-230 Port-Scope-Meldung quittieren (Umsetzung)

- **Review-Art:** unabhängiger Lauf (frischer Kontext, kein `fork`)
- **Skill-Version:** `.harness/skills/reviewer.md` (Status Accepted, Regelwerk v6.13.0)
- **Modell-ID:** claude-sonnet-5-5
- **Datum:** 2026-11-06
- **Gegenstand:** Commit-Range `69ee28c..HEAD` (f20a84e feat, c963cdf docs, 89d5948 planning), slice-230
- **Maßstab:** ADR-0045 (Accepted), `AC-FA-CONF-002`, SPEC-CONF-001/SPEC-CLI-001, `AGENTS.md` §3
- **Läufe:** `make gates` Exit 0 (Ausgabe in Datei, Exit-Code getrennt geprüft). Dieser Lauf prüft keine DoD (Sache des Verifiers).
- **Geltungsbereich der Prüfung:** Diff-Lektüre gegen ADR/AC plus ein Gate-Lauf. YAML-Randfälle (Alias, `!!binary`, Block-Scalar) wurden aus dem Code abgeleitet, **nicht** durch einen Lauf mit Fixture belegt (kein Produktionscode- und kein Testcode-Eingriff erlaubt).

## Findings

### HIGH

**H1 — Kommentar beschreibt die verworfene Alternative im Konjunktiv**
- kategorie: HIGH · quelle: `AGENTS.md` §3.7, Skill HIGH-Liste (Kommentar-Klassen)
- pfad: `internal/adapter/driven/config/config.go` (Kommentar über `decodeAcknowledge`, Satz „The map form is checked on the raw node: a decoded map would hide a duplicate key, a null and a non-string value.")
- befund: Der zweite Satz sagt, was eine **nicht gewählte** Form täte („would hide …"); er trägt keine der Klassen Zusage/Kopplung/Abgrenzung/Rang-Zeiger/Grenze im Indikativ über den geltenden Zustand. (Grenzfall: man kann ihn als Kopplung „Roh-Knoten, weil …" lesen; die Form ist trotzdem genau das Beispiel „Falsch" aus §3.7.)
- verifizierbar: nein (inferentiell, kein Gate)
- klasse: Kommentar im Konjunktiv über verworfene Alternative

### MEDIUM

**M1 — Geteilter Glob bei abweichender `direction`: Wortlaut der AC und Umsetzung weichen ab, ungetestet**
- kategorie: MEDIUM · quelle: `AC-FA-CONF-002` Boundary (überflüssig), ADR-0045 Punkt 3
- pfad: `internal/hexagon/core/rules.go` `InertPortScopes` (Schicht-Zuordnung nur bei eigener Auslösung), `PortScopeAcks`
- befund: AC/ADR sagen, ein geteilter Glob werde „**jeder** Schicht zugerechnet, die ihn trägt", und „keine der beiden gilt deshalb als überflüssig". Die Umsetzung (Auslegung d) rechnet ihn nur den Schichten zu, deren **eigene** Ableitung ihn inert lässt. Trägt Schicht B denselben Glob, löst ihn aber mit eigener `direction` auf, gilt B (hat sie nur diesen Glob) als „ohne Anlass", obwohl der Glob in der Diagnose steht (Wortlaut: „keiner ihrer Port-Globs erscheint in der Diagnose"). Zusätzlich: `Scope` und `Cause` der Zeile stammen von der ersten auslösenden Schicht; die Glob-Zeile einer nicht quittierten zweiten Schicht kann einen Scope zeigen, den ihre eigene Ableitung nicht ergibt. Die Auslegung ist fachlich begründbar (B ist real überflüssig), aber weder in der ADR noch im Plan festgelegt noch durch einen Test mit unterschiedlichen Richtungen gehalten (`TestPortScopeAcksSharedGlobCountsForEveryLayer` nutzt zwei Schichten ohne `direction`). SPEC-CLI-001 formuliert „ein geteilter Glob zählt für jede tragende Schicht" und deckt den Fall damit nicht ausdrücklich.
- verifizierbar: ja (Testfall mit zwei Port-Schichten, gleicher Glob, verschiedene `direction`)
- klasse: Auslegungsentscheidung außerhalb des Wortlauts ohne Beleg

### LOW

**L1 — ADR-0045 Punkt 2 widerspricht sich bei der Ursachen-Reihenfolge (Auslegung b)**
- kategorie: LOW · quelle: ADR-0045 Punkt 2, SPEC-CLI-001
- pfad: `internal/cli/cli.go` `writePortScopeNotice`, `internal/hexagon/core/rules.go` `ackedLayerOf`
- befund: „nach Name sortiert" und das Beispiel `direction, App-Glob ist groeber als der Scope` sind unvereinbar (Label-ASCII: „App-Glob" < „direction"; Konstantenname: coarse-app < direction). Die Umsetzung folgt dem Beispiel und der Reihenfolge der Abhilfe-Zeilen und hat das in SPEC-CLI-001 explizit festgeschrieben, ein Test hält es (`TestPortScopeAcksOnlyTriggeringGlobsCount`). Die Entscheidung ist vertretbar; die ADR bleibt immutabel widersprüchlich, die Spec ist die nachrangig tragende Stelle.
- verifizierbar: ja (Test existiert)
- klasse: interner Widerspruch einer Accepted-ADR, durch Spec aufgelöst

**L2 — Handbuch-Kopf: `Stand`-Datum nicht mitgezogen**
- kategorie: LOW · quelle: Konvention Handbuch-Versionierung
- pfad: `docs/user/benutzerhandbuch.md` Kopfzeile (`Handbuch-Version: 1.48 … Stand: 2026-10-07`) vs. Historie-Zeile 1.48 vom 2026-11-06
- befund: Version wurde gehoben, das `Stand`-Datum nicht; Kopf und Historie nennen zwei Daten.
- verifizierbar: nein
- klasse: Versionskopf und Historie auseinander

**L3 — Schwacher Assert im Negativtest „doppelter Schlüssel"; kein Test für doppeltes `acknowledge` am Schicht-Objekt**
- kategorie: LOW · quelle: `AC-FA-CONF-002` Negative, Mess-Regel „Testkommentar sagt nicht mehr zu, als die Assertion prüft"
- pfad: `internal/adapter/driven/config/config_acknowledge_test.go` (`"doppelter Schluessel"` mit `want == ""`)
- befund: `strings.Contains(err.Error(), "")` ist immer wahr; geprüft wird nur „irgendein Fehler mit `Schicht`". Die Mutation „seenKey-Prüfung entfernen" bleibt trotzdem rot (kein Fehler), eine Mutation, die die Fehlermeldung ändert, nicht. Doppeltes `acknowledge` auf Schicht-Ebene (von `yaml.v3` beim Struct-Decode abgelehnt, nicht von der Handprüfung) ist nicht abgedeckt.
- verifizierbar: ja
- klasse: Negativtest mit leerer Erwartung

**L4 — Mutations-Lücke bei der Abhilfe-Zeile der Ursache „App-Glob"**
- kategorie: LOW · quelle: ADR-0045 Punkt 2, `AC-FA-CONF-002` Happy
- pfad: `internal/cli/cli_acknowledge_test.go` (`TestAcknowledgeShortensNotToSilences`, `TestAcknowledgeCapCountsOnlyUnacknowledged`)
- befund: „keine Abhilfe-Zeile für die quittierte Schicht" ist nur für Ursache `direction` geprüft; im Deckel-Test (Ursache App-Glob) fehlt die Assertion auf die `Bei Ursache App-Glob`-Zeile. Eine Mutation, die `coarse` aus `scopes` statt `open` ableitet, überlebt. Ebenso fehlt ein gemischter CLI-Fall (quittierte + nicht quittierte Schicht mit Überlauf), der Restzahl und Abhilfe-Bedingung zugleich hält.
- verifizierbar: ja
- klasse: Test schwächer als die Zusage (Abhilfe je Ursache)

**L5 — Lesbarkeit: `writePortScopeNotice` bündelt jetzt fünf Aufgaben**
- kategorie: LOW · quelle: Wartbarkeit
- pfad: `internal/cli/cli.go` `writePortScopeNotice`
- befund: Kopfzeile, quittierte Zeilen, gedeckelte Glob-Zeilen mit Flag-Sammlung, Restzahl, Abhilfe und Redundanz-Zeilen in einer Funktion mit verschachtelter Schleife; Lint ist grün, die Funktion ist aber die längste der Datei geworden. `decodeAcknowledge` hat sechs Fehlerausgänge, bleibt aber linear lesbar.
- verifizierbar: nein
- klasse: Funktion wächst um Ausgabe-Zweige

### INFO

**I1 — Bewertung Auslegung (a):** „Glob quittiert nur, wenn jede tragende Schicht quittiert; Zählung/Zeile je tragender Schicht" ist konsistent mit ADR-Punkt 2 (Abhilfe „solange mindestens ein nicht quittierter Glob") und mit Punkt 3 (Zuordnung je Schicht). Kein Befund; die Regel steht in SPEC-CLI-001.

**I2 — Bewertung Auslegung (c):** `oneLine` (Leerraum zusammenziehen) ist ADR-/AC-lückenfüllend, aber in SPEC-CLI-001 festgeschrieben und schützt die Zeilenform bei Block-Scalar-Begründungen. Es gibt keinen Test mit mehrzeiliger Begründung.

**I3 — Randfälle YAML (aus dem Code abgeleitet, nicht per Fixture belegt):** Alias als Wert, `!!binary`, Zahl/Bool/Timestamp und `null` werden über `Kind != ScalarNode || Tag != "!!str"` mit Exit 2 abgelehnt; ein Alias auf einen **gültigen** String wird ebenfalls abgelehnt (fail-closed, Meldung „erwarte eine Begründung als String" wäre dort irreführend). Explizites `!!str` und Block-Scalar sind zulässig. Die „Listen-Form" der AC kann konstruktionsbedingt kein `acknowledge` tragen (Sequenz hat keine Schlüssel) und ist deshalb ohne eigenen Test/Code. Spezifikations-Historie nennt ADR-0045: Bestand hat in Historie-Zeilen Präzedenz, nicht neu und anders.

## Negativbefunde (geprüft, ohne Befund)

- **ADR-Punkte 1–6, Code:** alle sechs umgesetzt (Schlüssel/Map-Form, eine Zeile, Reihenfolge der Ausgabe, Kopfzeile zählt alle, Deckel nur für offene Globs, frühe Rückkehr entfällt, Exit-2-Fälle, keine Übergangsform, Befunde unberührt).
- **Unberührtheit (ADR 6/Prüfpunkt 4):** `AckPortScope` wird ausschließlich in `PortScopeAcks`/`ackedLayerOf` und im Decoder gelesen; keine Regel, kein Scope-Ableitungspfad (`scopeFor`, `inertCause`) liest es. Test vergleicht Exit-Code und stdout mit/ohne Quittierung.
- **Determinismus:** `ackedBy` ist eine Map nur für Lookups; Ausgabe läuft über Schicht-Reihenfolge und die nach Glob sortierte Diagnose; `Causes` mit fester Ordnung. Keine Map-Iteration im Ausgabepfad.
- **Konfigurationsfehler Exit 2:** Rolle ≠ port (explizit/inferiert), anderer Schlüssel, keine/leere Map, doppelter Schlüssel, `null`/Nicht-String/Leerraum je ein Config-Testfall; CLI-Exit-2 mit Schicht-Nennung an zwei Fällen belegt.
- **Hard Rules:** keine Inline-Suppression (`make gates` grün, `suppression-check`); keine Spec-Referenz nach unten außer der Historie-Zeile mit Präzedenz; Handbuch 1.48 und CHANGELOG `[Unreleased]` vorhanden; Commit-Scope `(planning)` nur Planning-Dateien.
- **Übrige Kommentare/Docs (§3.7):** Indikativ, ohne Chronik (bis auf H1).
- **Gates:** `make gates` Exit 0.

## Kategorie-Summary

HIGH 1 · MEDIUM 1 · LOW 5 · INFO 3

## Verdikt

Nicht freigegeben, solange H1 offen ist (Skill: Kommentar-Verstoß blockiert) und M1 nicht per Entscheidung oder Test geklärt ist. Kein Befund gegen Code-Korrektheit der ADR-Punkte 1–6. Konfliktpfad: H1 ist ein Rollen-Befund ohne bisherigen Widerspruch; bei Widerspruch des Implementers läuft die Sequenz über den Architect, nicht über Herabstufung.
