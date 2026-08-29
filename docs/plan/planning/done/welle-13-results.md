# welle-13-konsumenten-befunde — Ergebnis-Notiz

**Abschluss:** 2026-08-09. **Zweiter Durchlauf der Fünf-Schritt-Prozedur**
([`planning/README.md`](../README.md)) — und der erste **mit** einer Welle-Plan-Datei:
[`welle-13-konsumenten-befunde.md`](welle-13-konsumenten-befunde.md) ist per `git mv` von flach
nach `done/` gewandert, neben diese Notiz. Bei [`welle-12`](welle-12-results.md) entfiel dieser
Schritt mangels Plan-Datei; damit ist Schritt 3 jetzt **vollständig** belegt, nicht nur zur Hälfte.

---

## Geliefert

**Sechs Change Requests aus einem realen Fremdrepo-Einsatz, sechs Slices, alle in `done/`.**

| Slice | CR | Gegenstand | Vertragsarbeit |
|---|---|---|---|
| [slice-081](slice-081-heuristik-diagnose.md) | `CR-1` | Grenz-Diagnose: Zeilen, die zu keiner Kante führen | [ADR-0031](../../adr/0031-heuristik-grenzen-diagnose.md) |
| [slice-082](slice-082-print-mk-docker-indirektion.md) | `CR-6` | `--print-mk`: `$(DOCKER)` statt wörtlichem `docker` | — |
| [slice-083](slice-083-print-mk-digest-selbstbezug.md) | `CR-5` | `--print-mk` nannte den Digest des **Vorgängers** | [ADR-0030](../../adr/0030-kein-digest-im-generierten-fragment.md) |
| [slice-084](slice-084-handbuch-heuristik-grenzen.md) | `CR-3` | Heuristik-Grenzen dort, wo Konsumenten lesen | — |
| [slice-085](slice-085-schicht-ohne-aufloesung.md) | `CR-2` | Auflösungs-Diagnose: vollständig grün, vollständig blind | [ADR-0032](../../adr/0032-aufloesungs-diagnose-repoweit.md) |
| [slice-086](slice-086-forbidden-constructs-fail-closed.md) | `CR-4` | `forbidden_constructs` fail-closed statt still | [ADR-0033](../../adr/0033-forbidden-constructs-fail-closed.md) |

**Das Welle-Ziel war: „a-check sagt selbst, wo es blind ist — statt dass ein Konsument es
nachstellen muss."** Eingelöst durch drei Diagnosen auf **einer** Achse, grob nach fein, in fester
stderr-Reihenfolge:

| Diagnose | meldet | ADR |
|---|---|---|
| Abdeckung | **Datei** ohne Schicht | [ADR-0029](../../adr/0029-abdeckungs-diagnose-advisory.md) (vor dieser Welle) |
| Grenze | **Zeile** ohne Kante | [ADR-0031](../../adr/0031-heuristik-grenzen-diagnose.md) |
| Auflösung | **Ziel** ohne Schicht | [ADR-0032](../../adr/0032-aufloesungs-diagnose-repoweit.md) |

Alle drei sind advisory, alle drei lassen den Exit-Code unberührt, und alle drei halten an
derselben Grenze: **was von repo-externem Code nicht unterscheidbar ist, wird nicht behauptet.**
Wer eine vierte baut, hat ein Muster statt einer Einzelfallentscheidung.

## Was funktionierte

**Messen vor dem Spezifizieren — zweimal, und beide Male kippte es die Anforderung.**

- [slice-085](slice-085-schicht-ohne-aufloesung.md): Die im CR **wörtlich** verlangte Regel („je
  Schicht, wenn kein Symbol auflöst") wurde zuerst gebaut und gegen den eigenen Baum gefahren. Sie
  feuerte auf `internal/hexagon/core` — die Schicht, die laut [ARC-001](../../../../spec/architecture.md) abhängigkeitsfrei sein
  *muss*. Ein reiner Domänenkern ist per Konstruktion das erste Ziel dieser Regel. Wäre sie
  ungeprüft übernommen und **danach** getestet worden, wäre sie mit grünen Tests in die
  Spezifikation gewandert und hätte jeden sauber geschichteten Konsumenten angebellt.
- [slice-086](slice-086-forbidden-constructs-fail-closed.md): Die Aufbereitung fand **drei**
  weitere stille Ausgänge derselben fehlenden Validierung, die der CR nicht kannte; beim Bau kam
  ein **vierter** dazu. Der CR beschrieb einen Fall, das Feature hatte vier.

**Das Dogfooding war nicht Qualitätssicherung nach dem Bau, sondern das Instrument, das den
Entwurf korrigiert hat.** Kosten je Fall: eine Implementierung und ein `make arch-check`.
Vermiedener Schaden: eine `Accepted`-ADR ist immutabel.

**Die Aufbereitung vor dem Bau hat getragen.** Für
[slice-086](slice-086-forbidden-constructs-fail-closed.md) wurden drei Fragen vorab beantwortet —
Bestand (leer: 0 von 7 Konfigurationen), Vertragslage ([AC-FA-RULE-004](../../../../spec/lastenheft.md#ac-fa-rule-004--port-disziplin-regel-port-impurity) ist „Port-Disziplin", die
Rollenbindung **löst ihn ein**), Ausweg (`constructs` ist Komplement, nicht Ersatz). Der Bau musste
danach nur noch den Fehlertext entscheiden. Ohne diese Vorarbeit wäre der naheliegende Weg die
Ausweitung gewesen — und die ist nach [`MR-001`](../../../../harness/conventions.md#mr-001--source-precedence-mit-eigener-spezifikations-schicht)
gar nicht per ADR erreichbar, sondern ein Lastenheft-CR.

**Der Lint hat einen Refactor erzwungen, der überfällig war.** Die eine zusätzliche Validierung hob
`Load` auf zyklomatische Komplexität 16. Da Inline-Suppression eine Hard Rule verletzt, blieb nur
der Schnitt — entlang der Naht, die ohnehin da war: Pflichtblöcke gegen Optionalblöcke. Ein Gate,
das eine Grenze **ohne Ausweg** zieht, produziert den Refactor, den man sonst vertagt.

## Was anders lief

**Die Welle wuchs von vier auf sechs Slices**, weil der Konsument die vier gemeldeten Befunde als
sechs formale Change Requests nachreichte (`CR-2` und `CR-4` waren in der ersten Meldung nicht
enthalten). Eingetragen im Drift-Log der [Roadmap](../in-progress/roadmap.md).

**Die Verwerfungs-Bedingung von [slice-084](slice-084-handbuch-heuristik-grenzen.md) feuerte
nicht.** Sie sah vor, den Handbuch-Nachtrag zu verwerfen, falls
[slice-081](slice-081-heuristik-diagnose.md) zuerst käme und die Formen **vollständig** meldete.
Tatsächlich war 084 zuerst fertig, und 081 meldet **zwei** der vier Formen der Handbuch-Tabelle —
kompaktes TypeScript und import-ähnliche Zeilen in Strings bleiben still. Beide Orte bleiben, jetzt
verzahnt. Dass eine Welle eine solche Bedingung überhaupt trägt, war richtig: sie zwang zur
Nachmessung statt zur Annahme.

**Ein DoD-Wortlaut wurde bewusst nicht erfüllt.**
[slice-081](slice-081-heuristik-diagnose.md) verlangte „die Diagnose steht als Vertrag im
**Lastenheft**". Geliefert wurde die Spezifikations-Schärfung ohne Lastenheft-Bump — dem
Präzedenzfall [ADR-0029](../../adr/0029-abdeckungs-diagnose-advisory.md) folgend, weil eine
advisory stderr-Zeile ohne Exit-Code-Wechsel das *Wie* der bestehenden Ausgabe schärft und
[AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze) die
Offenlegung bereits zusagt. Die Abweichung steht im DoD selbst, nicht im Kleingedruckten.

**Ein Breaking Change ging heraus** ([slice-086](slice-086-forbidden-constructs-fail-closed.md)):
vier Konfigurationsformen enden künftig mit Exit 2 statt mit grünem Exit 0. Gemessen bricht das im
lokalen Bestand niemanden (0 von 7); der einzige bekannte Nutzer ist der CR-Melder, dessen Einträge
heute nicht wirken.

## Steering-Loop-Einträge

Gezogen aus `docs/plan/steering-loop.md` — das Register bleibt der
laufende Zähl-Ort. Hier stehen nur die in dieser Welle real aufgetretenen Vorfälle.

| Eintrag | in dieser Welle |
|---|---|
| `SL-001` — Gate-Lauf in einer Pipe verschluckt | **einmal**: der Guard blockierte `make doc-complete \| tail` bei der d-check-Recherche. Zähler **6 → 7** |
| `SL-002` — relative Verweise brechen beim Verzeichniswechsel | **zweimal**: vier präfixlose Geschwister-Verweise in [slice-087](../done/slice-087-index-vollstaendigkeit.md) (vor dem Commit gefangen) und **21 Verweise dieser Welle-Plan-Datei** beim `git mv` in `done/` — von `doc-check` gefangen, **nicht** vom zuständigen Sensor. Zähler **9 → 11** |
| `SL-003` — Commit-Betreff bezeichnet nicht die Arbeit | **nicht aufgetreten** — zweimal aktiv vermieden: der [ADR-0030](../../adr/0030-kein-digest-im-generierten-fragment.md)-Index-Nachtrag wurde als eigener Commit mit slice-083-Bezug geführt, und der Lerneintrag von [slice-081](slice-081-heuristik-diagnose.md) vom `git mv` getrennt, damit die Lifecycle-Commits `R100` bleiben |
| `SL-004` — neuer Sensor meldet sein eigenes Umfeld | **nicht aufgetreten** — aber als **Entscheidungsgrund** verwendet: [ADR-0031](../../adr/0031-heuristik-grenzen-diagnose.md) verwarf den Restmengen-Ansatz mit dem dreifach belegten Rauschprofil dieses Eintrags |
| `SL-005` — Datei fehlt im handgepflegten Index | **neu, zwei Vorfälle** (siehe unten) |
| `SL-006` — Dateiname/Anker geraten statt nachgesehen | **neu, vier Vorfälle in dieser Welle** (drei weitere in `welle-12`) |

**`SL-004` ist der bemerkenswerteste Eintrag dieser Welle, obwohl er nicht auftrat.** Ein
Register-Eintrag hat eine Architektur-Entscheidung geprägt, bevor der Fehler passieren konnte: Die
naheliegende Gestalt der Grenz-Diagnose (jede nicht gegriffene Import-Zeile melden) wurde verworfen,
weil `SL-004` dreimal belegt, dass ein solcher Sensor sein eigenes Umfeld meldet. Das ist der erste
Fall, in dem der Steering-Loop *feedforward* statt *feedback* gewirkt hat.

**`SL-002` deckte eine Lücke im eigenen Sensor auf — genau in dem Schritt, der ihn hätte finden
sollen.** Der `git mv` dieser Welle-Plan-Datei von flach nach `done/` brach **21** Verweise auf
einen Schlag: sie war präfixlos geschrieben (`done/slice-081-…`, `in-progress/roadmap.md`,
`README.md#…`), weil das aus der flachen Position korrekt war. `verify-slice-links` prüft die
Invariante „ein relativer Verweis löst aus **jedem** Lifecycle-Verzeichnis auf" — aber nur für
`slice-*.md`. **Die Welle-Plan-Datei ist eine neue Gattung und fällt durchs Raster**; sie ist erst
mit dieser Welle entstanden. Gefangen hat es `doc-check`, also **nach** dem `git mv` statt davor —
exakt der Zyklus, den der Sensor abschaffen sollte. Der Sensor-Scope gehört auf
`welle-*.md` erweitert; das ist ein Ein-Zeilen-Fix am Glob, aber er braucht seine eigene
Negativ-Probe und damit einen eigenen Schnitt.

**`SL-005` ist der unangenehmste.** Der ADR-Index-Fehler trat **zweimal an einem Tag** auf — das
zweite Mal, nachdem er diagnostiziert, benannt und als Folge-Slice festgehalten war, beim
allernächsten ADR. Dieselbe Lehre wie bei `SL-001`/`SL-002`, diesmal ohne dass überhaupt ein Guide
dazwischenstand: **Wissen allein verhindert den Fehler nicht.**

## Folge-Slices

- [slice-087](../done/slice-087-index-vollstaendigkeit.md) — Vollständigkeits-Sensor für
  handgepflegte Indizes (`SL-005`). Trigger **sofort**, wartet auf nichts. Die Vorarbeit liegt bei:
  Bestandsmessung (genau ein Index) und d-check-Abdeckung (kein Modul deckt Ziel → Verweis, damit
  CR-fähig). Offen ist der Entscheid lokaler Sensor gegen d-check-CR.

**Benannt, aber ausdrücklich nicht geschnitten** — jeweils mit dem Trigger, der sie fällig macht:

- **`verify-slice-links` auf `welle-*.md` erweitern.** Der Sensor prüft nur `slice-*.md`; die
  Welle-Plan-Datei fiel durchs Raster und brach beim `git mv` mit 21 Verweisen auf einmal. Trigger:
  **sofort** — die nächste Welle-Datei trifft dieselbe Kante. Ein-Zeilen-Fix am Glob, der aber eine
  eigene Negativ-Probe braucht (eine Fixture in `welle-*.md`-Form, die vorher rot ist).

- **Schicht-Blacklist außerhalb `port`.** Für „Muster in Schicht X verboten, sonst egal" gibt es
  kein Werkzeug; `constructs` ist das Komplement, kein Ersatz. Ein **Lastenheft-CR**, sobald ein
  Konsument den Bedarf belegt ([slice-086 §4](slice-086-forbidden-constructs-fail-closed.md)).
- **Weitere Blöcke mit ungeprüfter Auswertungs-Bedingung.** Der Lerneintrag von
  [slice-086](slice-086-forbidden-constructs-fail-closed.md) verallgemeinert: eine Bindung, die nur
  in [SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung) steht und nicht in [SPEC-CONF-001](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema), ist ein stiller Filter statt eines
  Vertrags. `adapter_sink` und `markers.ignore_symbols` sind die Kandidaten für eine Messung.
- **Teilausfall der Auflösung.** [ADR-0032](../../adr/0032-aufloesungs-diagnose-repoweit.md) meldet
  nur den Totalausfall; vier gesunde Schichten und eine kaputte bleiben still. Eigener Entscheid,
  sobald ein Teilausfall real auftritt.

**Unverändert offen aus `welle-12`:** [slice-079](../done/slice-079-gate-consistency-abloesen.md),
[slice-080](../open/slice-080-verify-abloesung-dcheck.md) (wartet auf ein d-check-Release) und
`F-9` (Freigabe-Belege).

## Verifikation (Schritt 1 der Prozedur)

| Prüfung | Ergebnis |
|---|---|
| Alle sechs Slices der Welle in `done/` | ✅ `slice-081` … `slice-086` |
| Jeder Befund behoben **oder** als deklarierte Grenze ausgewiesen | ✅ sechs von sechs behoben; drei Grenzen ausdrücklich ausgewiesen (Teilausfall, Schicht-Blacklist, zwei nicht gemeldete Import-Formen) |
| Für die spec-first-Slices eine ADR mit `Status: Accepted` **vor** dem Code | ✅ [ADR-0030](../../adr/0030-kein-digest-im-generierten-fragment.md), [ADR-0031](../../adr/0031-heuristik-grenzen-diagnose.md), [ADR-0032](../../adr/0032-aufloesungs-diagnose-repoweit.md), [ADR-0033](../../adr/0033-forbidden-constructs-fail-closed.md) — je ein eigener Spec-Commit vor dem Code-Commit |
| `make ci` (Baseline-Ersetzung für den Replay-Lauf, [`MR-008`](../../../../harness/conventions.md#mr-008--kein-replay-keine-agenten-telemetrie)) | ✅ **Exit 0** — „[ci] gates + image-test grün" |
| Carveout-Audit (Schritt 2) | ✅ Bestand **null** — [`carveouts/README.md`](../../carveouts/README.md); das ist eine Aussage, keine Auslassung |
| Konsumenten-Bestandsprobe (neu in dieser Welle) | ✅ alle **sieben** lokalen `.a-check.yml` laden unverändert mit Exit 0 |

**Was diese Verifikation nicht belegt:** dass die drei Diagnosen bei einem **realen** Konsumenten
das Richtige melden. Geprüft wurden Fixtures, die den gemeldeten Fällen nachgebaut sind, und der
eigene Baum. Der ehrliche nächste Schritt wäre eine Rückmeldung des CR-Melders gegen ein Release
mit diesen Änderungen — `[Unreleased]` trägt sie, ein Release ist **nicht** Teil dieser Welle
([`welle-13-konsumenten-befunde.md`](welle-13-konsumenten-befunde.md) §6).

**Ebenfalls nicht belegt:** ein unabhängiges Review dieser Welle. Alle Prüfungen sind Selbst-Reviews
und Gate-Läufe; der `welle-12`-Trigger („ein Lauf außerhalb dieser Modell-Familie") gilt für die
Migration, nicht für diese Welle — er ist hier **nicht** erneut gefeuert worden.
