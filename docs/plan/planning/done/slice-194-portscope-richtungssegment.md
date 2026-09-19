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

**Lerneintrag — Form:** benannte Spec-Lücke.

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

## 2. Ausgangsmessung — vorher und nachher

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

**Nachher — dieselben Fixtures, gemessen gegen das released `v0.19.0` und gegen das Image mit
dieser Änderung.** Die zwei Zeilen unterscheiden sich **nur** im gebauten Stand:

| Variante | Config | `v0.19.0` (vorher) | dieser Stand (nachher) |
|---|---|---|---|
| v1 | `…/createorder/ports/**` | `port-locality: 1`, Exit 1 | `port-locality: 1`, Exit 1 — **unverändert** |
| v2 | `…/ports/{inbound,outbound}/**`, **keine** `direction` | `gesamt: 0 Befund(e)`, Exit 0 — **still** | 0 Befunde, Exit 0, **und der Hinweis** nennt beide Globs samt abgeleitetem Scope |
| v4 | getrennte Port-Schichten **mit** `direction` (der Weg, den §4 anweist) | `gesamt: 0 Befund(e)`, Exit 0 — **still** | `port-locality: 1`, Exit 1 — **dieselbe Meldung und derselbe Scope wie v1** |

**v4 ist die tragende Zeile:** Der Verstoß, den v1 bei `…/ports/**` meldet, wird bei getrennten
Port-Schichten **wieder** gemeldet — mit demselben Scope (`…/createorder`) und derselben Meldung.
**v2 ist die zweite:** Wo die Ableitung die Richtung nicht kennen kann, bleibt sie unverändert
inert, aber der Lauf sagt es jetzt. Zwischen „still" und „unbeurteilt" liegt genau dieser Hinweis.

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
| `internal/cli/cli.go` (+ Kern-Anteil) | update | weitere Advisory-Diagnose nach dem Muster von [`ADR-0029`](../../adr/0029-abdeckungs-diagnose-advisory.md), [`ADR-0035`](../../adr/0035-grenz-diagnose-gegen-globs.md), [`ADR-0032`](../../adr/0032-aufloesungs-diagnose-repoweit.md): stderr, exit-neutral, gedeckelt, mit „Abhilfe" |
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

**Die drei Mutations-Proben, mit ihrer roten Meldung** (Mess-Regel 2: grün beweist nichts, und
eine Probe ohne Meldung belegt nur einen Exit-Code):

| Probe | Mutation | rote Meldung |
|---|---|---|
| Kern | Schnitt zurückgenommen (`scopeFor` = nur Marker) | `ein Port-Glob mit Richtungssegment darf port-locality nicht abschalten, got []` |
| Leitplanke 1 | `!nestedInAppTree(m, prefix)` entfernt | `eine Geschwister-Port-Schicht MIT Richtung bleibt inert, got [{other/svc.go 3 port-locality app außerhalb Port-Scope hex: …}]` |
| Leitplanke 2 | `!appTreeContains(m, marker)` entfernt | `der Scope darf nicht ueber den Port-Ordner hinauswachsen, got []` |

Alle drei sind danach **zurückgenommen**; der Stand, über den die Gates liefen, trägt keine
Mutation.

**Auszuführende Gates:** `make gates` (tragend `test`, `coverage-gate`, `arch-check`,
`doc-check`, `doc-structure`), zum Abschluss `make verify`.

## 4. Definition of Done

- [x] Die Spec-Kette trägt die geschärfte Ableitung:
      [`SPEC-RULE-001`](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung)
      nachgezogen, eine neue ADR samt Index-Eintrag.
- [x] Die Engine setzt es durch: `portScope` überspringt das Richtungssegment, und die weitere
      Advisory-Diagnose meldet die Restfälle.
- [x] Tests und Benutzerhandbuch sind nachgezogen; jede Probe ist **rot gewesen** und nennt ihre
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
  klassisch-hexagonalen Repo und wird abgeschaltet statt befolgt. — **Ausgang:** *entfallen*,
  gestrichen mit Begründung: Die Diagnose keilt auf **zwei** Bedingungen (`nestedInAppTree` **und**
  `!appTreeContains`), und der Geschwister-Fall erfüllt schon die erste nicht — sein Verzeichnis
  liegt außerhalb des App-Baums. `TestInertPortScopesSiblingPortsSilent` hält das für die
  **Diagnose** fest, `TestPortLocalitySiblingWithDirectionStaysInert` für die **Ableitung** — beide
  sind in der Mutations-Probe **rot** gewesen (§3).
- **Der Diskriminator „deklarierte Richtung" greift zu kurz.** Eine Richtungsebene, die anders
  heißt als ihr `direction`-Wert, bleibt still — dieselbe Klasse, seltener. — **Ausgang:**
  *entfallen*, gestrichen mit Begründung: Genau dafür ist die Diagnose da, und sie greift
  **breiter** als die Ableitung — sie keilt nicht auf dem Vokabular, sondern auf „Verzeichnis im
  App-Baum, Scope nicht". `TestInertPortScopesReportsUndeclaredDirection` fährt den Fall ohne
  deklarierte Richtung als roten Beleg.
- **Die Semantik-Änderung verschiebt bestehende Scopes.** Ein Port-Glob, dessen letztes Segment
  zufällig dem deklarierten Richtungswert entspricht, skopiert danach eine Ebene höher und kann
  neue Befunde erzeugen. — **Ausgang:** *entfallen*, gestrichen mit Begründung: Der zusätzliche
  Schnitt wird nur genommen, wo der Marker-Schnitt den App-Baum **verfehlt** und der tiefere ihn
  **erreicht**; der Scope kann dadurch nie weiter werden als zuvor. Die Umkehr steht als Test
  (`TestPortLocalityDirectionSegmentOwnSliceSilent`), der Aufweichungs-Fall als eigener Test
  (`TestPortLocalityDirectionFoldedKeepsSliceScope`) — beide in der Mutations-Probe **rot** (§3).
- **Der Kommentar im Konsumenten-Repo wird durch diesen Slice falsch.** Er hält die Kopplung
  fest, die es danach nicht mehr gibt. — **Ausgang:** *entfallen*, gestrichen mit Begründung: Der
  Kommentar liegt in einem **fremden** Repo; dieser Slice kann ihn nicht ändern, und ihn hier zu
  einem Folge-Slice zu machen hieße, eine Adresse zu erfinden, die die Sendung nicht annehmen
  kann. Die Meldung dorthin geht mit der Antwort auf die Konsumenten-Meldung hinaus.
- **`BEO-PLAN/messung-ohne-reproduzierbares-instrument` steht bei 2×.** Die Meldung kam mit
  einem externen Nachweis, der im Repo nicht nachvollziehbar war. — **Ausgang:** *entfallen*,
  gestrichen mit Begründung: Das Instrument liegt mit diesem Slice **im** Repo — die beiden
  Glob-Varianten gegen denselben Verstoß sind Regeltests, die Mutations-Probe ist gefahren und
  ihre rote Meldung benannt (§3). Der Beleg ist damit wiederholbar; ein Vorkommen der Klasse war
  das nicht.

## 8. Closure-Notiz

**Lerneintrag — Form: benannte Spec-Lücke.** *Die Spezifikation war vollständig und trotzdem
falsch: „minus seinem letzten Pfad-Segment" beschreibt eine Ableitung, die **genau ein** Segment
zwischen App-Baum und Port-Ordner voraussetzt — und diese Voraussetzung stand nirgends.* Der Satz
war nicht ungenau, er war **unterbestimmt**, und die fehlende Bedingung fiel erst auf, als eine
zweite Dimension ([`ADR-0036`](../../adr/0036-port-richtung-inbound-outbound.md)) denselben
Glob-Präfix benutzte. **Weil** eine Ableitung ihre Annahmen mitträgt, gehört die Annahme **in den
Satz**: Die drei Segmente (App-Baum · Port-Ordner · Richtungssegment) sind jetzt benannt und in
beiden Richtungen geprüft.

**Der zweite Lerneintrag nimmt die Review-Klasse aus dem Vorgänger-Slice auf:** *Ein grüner Lauf
über einer **stillen** Regel sieht aus wie ein grüner Lauf über einer geprüften.* Die Meldung des
Konsumenten war nur deshalb eine Meldung, weil dort **gemessen** wurde; hier hatte niemand
nachgesehen, und kein Gate fing es — die drei Diagnosen sind die Antwort auf genau diese
Unterscheidung, und die neue reiht sich in die bestehenden ein.

**Steering-Loop-Eintrag:** gezählt, nicht verkörpert.
[`BEO-HARNESS/zwei-regeln-machen-einander-unmoeglich`](../observations/BEO-HARNESS/zwei-regeln-machen-einander-unmoeglich/observation.md)
erreicht mit diesem Slice **2×** (Beleg `evidence/slice-194.md`); die Schwelle ist nicht erreicht,
ein Ausgang ist nicht fällig. Ein neuer Eintrag wurde nicht angelegt.

**Beobachtungs-Register ([`../observations/`](../observations/README.md)):** ein Beleg ergänzt, ein
Eintrag **benannt statt gezählt**.
[`BEO-HARNESS/zwei-regeln-machen-einander-unmoeglich`](../observations/BEO-HARNESS/zwei-regeln-machen-einander-unmoeglich/observation.md)
— `evidence/slice-194.md`, Zähler **2×**; der Beleg führt die **stille** Ausprägung aus, die die
Erkennungs-Sentence des Eintrags bis hierher nicht kannte.
[`BEO-PLAN/messung-ohne-reproduzierbares-instrument`](../observations/BEO-PLAN/messung-ohne-reproduzierbares-instrument/observation.md)
bleibt bei **2×**: Die Ausgangsmessung nennt ihre Parameter (§2), und das Instrument liegt mit
diesem Slice als Regeltest im Repo — die Bedingung des Eintrags ist damit erfüllt, nicht verletzt.
[`BEO-KERN/dirvocab-portfor-auseinander`](../observations/BEO-KERN/dirvocab-portfor-auseinander/observation.md)
bleibt bei 1×: anderer Mechanismus, siehe §9.

**Benannt, nicht gezählt — eine Klasse aus dem Review, die keinen Eintrag bekommt.** F-4 fand die
`Stand:`-Ziffer dieses Eintrags bei `1×`, während seine Belegliste schon zwei Dateien trug: eine
zweite Quelle neben dem abgeleiteten Zähler, und modul-06 verwirft genau die. Es ist der **erste**
Vorfall dieser Art (17 weitere Einträge mit Ziffer stimmen), und die Steering-Loop-Regel setzt den
Eintrag beim **zweiten** an — bis dahin bleibt es hier benannt.

**Folge-Slices:** [slice-195](../open/slice-195-handbuch-adapter-vokabel.md) (Handbuch-Vokabel für
die Adapter-Rolle) — liegt in `open/`.

**Risiken aus §7:** alle fünf *entfallen*, gestrichen mit Begründung — siehe dort.

**Drei Paarungen:** Anker — kein `liegt in`-Feld gesetzt, weil nichts verkörpert wurde; nichts zu
paaren · Folge-Slice getragen (slice-195 existiert in `open/`) · Register getragen (beide genannten
Pfade existieren und tragen Belege).

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
  (1×): die nächstliegende Klasse, und **sie trifft zu**. Ihre Erkennungs-Sentence („an einem
  Sensor, der **rot** wird") beschreibt nur die **laute** Ausprägung; hier schweigt die Regel. Der
  Plan hatte die Zuordnung ausdrücklich offengelassen — **sie ist mit dem Bau entschieden**: Der
  Eintrag wird auf **2×** gehoben, und der Beleg nennt die stille Ausprägung (§8).
- `BEO-PLAN` — [`messung-ohne-reproduzierbares-instrument`](../observations/BEO-PLAN/messung-ohne-reproduzierbares-instrument/observation.md)
  (2×) ist einschlägig und steht als Risiko in §7.
- `BEO-SPEC` — keine Treffer zu diesem Gegenstand.

**Modus-Begründungsblock:** alle berührten Sub-Areas GF (KERN, SPEC; `docs/user/` ist als
Sub-Area unbenannt, nicht BF). Kein Block je Sub-Area nötig.
