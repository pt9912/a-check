# slice-194 — portScope überspringt das Richtungssegment

**Welle:** ohne Welle.

**Bezug:** Externe Meldung des Konsumenten `hexslice-architecture` (2026-09-19, gegen
`v0.19.0` gemessen), im Repo reproduziert (§2). Schärft
[`SPEC-RULE-001`](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung); die
Zusage von
[`AC-FA-RULE-010`](../../../../spec/lastenheft.md#ac-fa-rule-010--port-lokalität-regel-port-locality)
bleibt unverändert — der Slice macht sie wieder erreichbar.

**Berührte Spec-Stellen:**
[`AC-FA-RULE-010`](../../../../spec/lastenheft.md#ac-fa-rule-010--port-lokalität-regel-port-locality) ·
[`AC-FA-RULE-008`](../../../../spec/lastenheft.md#ac-fa-rule-008--richtungs-dimension-regel-port-direction-mismatch) ·
[`AC-FA-CONF-001`](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml) ·
[`SPEC-RULE-001`](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung) ·
[`SPEC-CLI-001`](../../../../spec/spezifikation.md#spec-cli-001--aufruf-scan-wurzel-und-exit-codes)

**Verantwortlich:** — (bis zur Priorisierung).

**Autor:** Claude, im Auftrag des Maintainers. **Datum:** 2026-09-19.

**Lerneintrag — Form:** wird bei Closure benannt (eine der drei Formen der Ziel-Form).

---

## 1. Ziel und Abgrenzung

**Ziel:** Trägt eine Port-Schicht ihre `direction` — der von
[`ADR-0036`](../../adr/0036-port-richtung-inbound-outbound.md) vorgesehene und im Benutzerhandbuch
§4 angewiesene Weg, `port-direction-mismatch` überhaupt erst scharf zu machen —, dann bleibt
`port-locality` **wirksam**. Die dokumentierte Aktivierung der einen Regel darf die andere nicht
lautlos abschalten.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Die Richtungs-Dimension selbst** (Wertebereich, Validierung, Paarung) — *Schicht-Abgrenzung:*
  [`ADR-0036`](../../adr/0036-port-richtung-inbound-outbound.md) ist `Accepted` und wird nicht
  angefasst; Gegenstand hier ist die **Scope-Ableitung**, nicht die Richtung.
- **Ein Deklarations-Feld für den Port-Ordner-Marker** (etwa `port_marker: ports`) — *es wäre ein
  anderer Vorgang:* das ersetzt eine Ableitung durch eine Deklaration, und
  [`AC-FA-RULE-010`](../../../../spec/lastenheft.md#ac-fa-rule-010--port-lokalität-regel-port-locality)
  stellt „der Scope ist **pfad-abgeleitet** (keine Deklaration)" ausdrücklich out-of-scope.
- **Die Config des Konsumenten `hexslice-architecture`** — *ein anderer Vorgang in einem anderen
  Repo:* dort steht seit dem Umbau ein Kommentar, der genau diese Kopplung festhält und nach
  diesem Slice **falsch** ist (Risiko in §7). Das ist eine Meldung an den Konsumenten, kein Edit
  hier.
- **`port-direction-mismatch` für driven-Adapter** — *Bestand bleibt bewusst stehen:* in Go
  implementieren driven-Adapter Outbound-Ports strukturell (kein Import), die Regel kann dort
  nicht feuern. Das ist die zweite Hälfte der Meldung und ein eigener Gegenstand, kein Nachzug.

## 2. Ausgangsmessung (beim Übergang nach `in-progress/` zu wiederholen)

**Instrument:** `a-check:dev` — der Regelcode-Stand ist seit
[slice-121](../done/wellenlos/slice-121-port-richtung-inbound-outbound.md) unverändert und liegt
damit auch in `v0.19.0`. Baumkopie des Lab-Fixtures (`hexslice-architecture/lab/examples/go`) mit
**einem** injizierten Fremd-Slice-Import (`cancelorder` importiert den slice-lokalen Port von
`createorder`), gelaufen mit `docker run --rm --network none -v <kopie>:/src:ro a-check:dev /src`.

| Lauf | Port-Glob | Importziel | Ausgabe | Exit |
|---|---|---|---|---|
| v1 | `…/createorder/ports/**` | fremder slice-lokaler Port | `port-locality: 1` — „app außerhalb Port-Scope …/createorder" | **1** |
| v2 | `…/createorder/ports/{inbound,outbound}/**` | **derselbe** Import | `gesamt: 0 Befund(e)` | **0** |
| v3 (Kontrolle) | wie v2 | fremde **App**-Slice statt Port | `lateral-slice: 1` | **1** |

**v3 ist die tragende Zeile:** Bei identischer Config feuert eine andere Regel weiter — die
Config ist **nicht** tot, sondern `port-locality` ist **still**. Ohne diese Kontrolle wäre
„inert" von „Config kaputt" nicht zu unterscheiden.

**Geltungsbereich:** gemessen für den Go-Pfad `app`-Importeur → slice-lokaler Port über einen
literalen Port-Glob-Präfix. Andere Sprach-Backends und andere Auflösungs-Modi wurden **nicht**
gemessen; die Ableitung ist sprach-agnostisch, die Aussage ist es nicht.

**Warum das kein Implementierungsfehler ist:** „**minus seinem letzten Pfad-Segment**
(Port-Ordner-Marker, typisch `ports`)" steht wörtlich in
[`SPEC-RULE-001`](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung)
(spezifikation.md:268 und :313). Die Implementierung ist **spec-konform**; die Lücke entstand
erst, als [`ADR-0036`](../../adr/0036-port-richtung-inbound-outbound.md) eine zweite Vokabel
einführte, die denselben Glob-Präfix benutzt. Der Slice ist darum eine **Spec-Änderung**, kein
Bugfix — und die neue ADR trägt sie.

**Entscheidung des Maintainers (2026-09-19):** Route **(a)+(d)** — die Semantik wird korrigiert
**und** die Restfälle werden laut. Route (b) (Validierung, Exit 2) und (c) (nur als Grenze
dokumentieren) sind verworfen.

## 3. Umsetzung

**Spec-first** ([`AGENTS.md`](../../../../AGENTS.md) §5): Spezifikation → ADR → Code → Tests →
Doku. Das Lastenheft steht **nicht** voran: die Zusage von
[`AC-FA-RULE-010`](../../../../spec/lastenheft.md#ac-fa-rule-010--port-lokalität-regel-port-locality)
ist unverändert, nur ihre Ableitung war unterbestimmt.

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| [`spec/spezifikation.md`](../../../../spec/spezifikation.md) | update | [`SPEC-RULE-001`](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung): die Ableitung überspringt ein abschließendes Richtungssegment; Versions-Bump + Historie |
| eine neue ADR unter [`docs/plan/adr/`](../../adr/README.md) + Index | neu / update | Nummer vergibt der schreibende Lauf nach dem Index — im Plan **nicht** vorweggenommen. **Kein** `slice-\d{3}`-Token im ADR-Körper ([`.d-check.yml`](../../../../.d-check.yml), `matrix`) |
| [`internal/hexagon/core/rules.go`](../../../../internal/hexagon/core/rules.go) | update | `portScope` überspringt das Richtungssegment; Diskriminator ist der **deklarierte** `direction`-Wert der gewinnenden Schicht, **keine** hartkodierte Wortliste |
| `internal/cli/cli.go` (+ Kern-Anteil) | update | dritte Advisory-Diagnose nach dem Muster von [`ADR-0029`](../../adr/0029-abdeckungs-diagnose-advisory.md)/[`ADR-0031`](../../adr/0031-heuristik-grenzen-diagnose.md): stderr, exit-neutral, gedeckelt, mit „Abhilfe" |
| `internal/hexagon/core/rules_test.go`, `internal/cli/cli_test.go` | update | beide Glob-Varianten gegen denselben Verstoß (rot/grün) und die Diagnose — je mit **Mutations-Probe** |
| [`docs/user/benutzerhandbuch.md`](../../../../docs/user/benutzerhandbuch.md) | update | §3.7 dritter Fallstrick + §4-Absatz „Richtung": die Kopplung benennen, den Geschwister-Fall abgrenzen |
| [`CHANGELOG.md`](../../../../CHANGELOG.md) | update | Verhaltens-Änderung an einer Regel und eine neue Diagnose |

**Warum der Diskriminator der deklarierte Wert ist.** Eine feste Wortliste (`inbound|outbound`)
träfe auch einen Ordner, der zufällig so heißt, und ließe eine anders benannte Richtungsebene
weiter still. Der Validator kennt die gültige Menge bereits rollen-abhängig
([`ADR-0036`](../../adr/0036-port-richtung-inbound-outbound.md) §Konsequenzen: „die Zuordnung steht
an zwei Orten") — die Ableitung liest dieselbe Quelle, statt eine dritte zu erfinden.

**Warum die Diagnose den legitimen Fall nicht treffen darf.** Geschwister-Ports (klassisches
Hexagonal, `hex/ports/**` neben `hex/services/**`) sind **dokumentiert inert**
([`AC-FA-RULE-010`](../../../../spec/lastenheft.md#ac-fa-rule-010--port-lokalität-regel-port-locality),
gemessen in `TestPortLocalitySiblingPortsNoFinding`). Eine Diagnose, die dort meldet, wird auf
jedem klassisch-hexagonalen Repo zum Rauschen und damit abgeschaltet. Der Unterschied ist
prüfbar: **hier** liegt der Port im App-Baum und nur der Glob endet eine Ebene zu tief, **dort**
liegt er außerhalb.

**Auszuführende Gates:** `make gates` (tragend `test`, `coverage-gate`, `arch-check`,
`doc-check`, `doc-structure`), zum Abschluss `make verify`.

## 4. Definition of Done

- [ ] Die Spec-Kette trägt die geschärfte Ableitung:
      [`SPEC-RULE-001`](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung)
      nachgezogen, eine neue ADR samt Index-Eintrag.
- [ ] Die Engine setzt es durch: `portScope` überspringt das Richtungssegment, und die dritte
      Advisory-Diagnose meldet die Restfälle.
- [ ] Tests und Benutzerhandbuch sind nachgezogen; jede Probe ist **rot gewesen** und nennt ihre
      Meldung.

- [ ] Unabhängiger Review durchgeführt (Report unter [`docs/reviews/`](../../../reviews/README.md)).
- [ ] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [ ] Beobachtungs-Register fortgeschrieben.
- [ ] Jedes Risiko aus §7 trägt einen Ausgang.

`make gates` und `make verify` grün.

## 5. Trigger

**Start** (`open` → `in-progress`): das WIP-Limit ist frei. Keine Abhängigkeit von
[slice-188](../done/wellenlos/slice-188-voll-abgleich-gate-und-skill.md)/[slice-189](../open/slice-189-voll-abgleich-spec-straten.md)/[slice-190](../open/slice-190-id-schema-deklaration-ueberarbeiten.md)/[slice-191](../open/slice-191-chronik-phrasen-sensor.md) —
der Slice berührt keine ihrer Dateien.

**Rückführungen — vorab benannt:**

- `in-progress` → `next` (zu groß): Zeigt die Umsetzung, dass die Diagnose ohne Änderung an
  `appTreeContains`/`portScope` nicht sauber vom Geschwister-Fall zu trennen ist, ist **(a)** der
  Slice und **(d)** ein eigener — die Semantik-Korrektur wartet nicht auf die Diagnose.
- `in-progress` → `open` (blockiert): Widerspricht die ADR dem Wortlaut von
  [`AC-FA-RULE-010`](../../../../spec/lastenheft.md#ac-fa-rule-010--port-lokalität-regel-port-locality)
  statt ihn zu präzisieren, ist das eine Lastenheft-Änderung (Change Request) und keine
  Spec-Schärfung — dann zuerst dorthin zurück.

## 6. Closure-Trigger

DoD vollständig, `make gates` und `make verify` grün, Closure-Notiz mit Lerneintrag.

## 7. Risiken und offene Punkte

- **Die Diagnose trifft den legitimen Geschwister-Fall.** Dann meldet sie auf jedem
  klassisch-hexagonalen Repo und wird abgeschaltet statt befolgt. — **Ausgang:** <offen bis
  Closure>
- **Der Diskriminator „deklarierte Richtung" greift zu kurz.** Eine Richtungsebene, die anders
  heißt als ihr `direction`-Wert, bleibt still — dieselbe Klasse, seltener. — **Ausgang:** <offen
  bis Closure>
- **Die Semantik-Änderung verschiebt bestehende Scopes.** Ein Port-Glob, dessen letztes Segment
  zufällig dem deklarierten Richtungswert entspricht, skopiert danach eine Ebene höher und kann
  neue Befunde erzeugen. — **Ausgang:** <offen bis Closure>
- **Der Kommentar im Konsumenten-Repo wird durch diesen Slice falsch.** Er hält die Kopplung
  fest, die es danach nicht mehr gibt. Ein fremdes Repo, kein Edit hier; die Meldung dorthin
  gehört zum Abschluss. — **Ausgang:** <offen bis Closure>
- **`BEO-PLAN/messung-ohne-reproduzierbares-instrument` steht bei 2×.** Die Meldung kam mit
  einem externen Nachweis, der im Repo nicht nachvollziehbar war. Dieser Slice liefert das
  Instrument **im** Repo — ob das den Eintrag auflöst, entscheidet der Lese-Schritt bei Closure,
  nicht dieser Plan. — **Ausgang:** <offen bis Closure>

## 8. Closure-Notiz

*(bei Closure auszufüllen)*

## 9. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** Berührt sind `KERN` (Scope-Ableitung und Diagnose),
`SPEC` (Spezifikation + ADR) und die **Handbuch-Familie** `docs/user/`. Die ersten zwei erfüllen
die Schwelle (≥ 2 von 3 Achsen) unbestreitbar. Die dritte tut es **auch** — eigene
Pfad-Familie, eigene Diskrepanz-Zeile sinnvoll — und hat trotzdem **keine Zeile** in
[`harness/conventions.md`](../../../../harness/conventions.md) §Modus-Deklaration. Der Befund
gehört hierher, die Korrektur nicht: er trägt den Folge-Slice
[slice-195](../open/slice-195-handbuch-adapter-vokabel.md).

**Vorgelagert — offene Beobachtungen sichten:** Register über die berührten Kürzel gelesen:

- `BEO-KERN` — **ein** Treffer:
  [`dirvocab-portfor-auseinander`](../observations/BEO-KERN/dirvocab-portfor-auseinander/observation.md)
  (1×). **Nicht dieselbe Beobachtung:** dort driften zwei Pakete über eine Paketgrenze, hier
  schweigt eine Regel wegen einer Glob-Form. Der Zähler bleibt unberührt.
- `BEO-HARNESS` —
  [`zwei-regeln-machen-einander-unmoeglich`](../observations/BEO-HARNESS/zwei-regeln-machen-einander-unmoeglich/observation.md)
  (1×) ist die **nächstliegende** Klasse, trifft aber nicht: ihre Erkennungs-Sentence lautet „an
  einem Sensor, der **rot** wird" — hier wird nichts rot, die Regel **schweigt**. Ob das eine
  eigene Beobachtung ist oder die Klasse weiter gefasst werden muss, entscheidet der Schreibende
  bei Closure; der Plan nimmt die Antwort **nicht** vorweg und zählt den Eintrag darum nicht hoch.
- `BEO-PLAN` — [`messung-ohne-reproduzierbares-instrument`](../observations/BEO-PLAN/messung-ohne-reproduzierbares-instrument/observation.md)
  (2×) ist einschlägig und steht als Risiko in §7.
- `BEO-SPEC` — keine Treffer zu diesem Gegenstand.

**Modus-Begründungsblock:** alle berührten Sub-Areas GF (KERN, SPEC; `docs/user/` ist als
Sub-Area unbenannt, nicht BF). Kein Block je Sub-Area nötig.
