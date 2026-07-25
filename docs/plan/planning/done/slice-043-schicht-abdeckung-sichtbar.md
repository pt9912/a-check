# slice-043 — Schicht-Abdeckung sichtbar machen (P-Rest-Kandidat 2a)

**Review:** [Review-Synthese](../../../reviews/2026-07-25-slice-043-abdeckungs-diagnose.md).
**Status:** **done (2026-07-25)** — Entscheide abgenommen (§0): Gestalt **(a)** — nur die
Diagnose, `strict_coverage` **vertagt**. Umsetzung, Review und Merge erledigt; Ergebnis §7, Closure §8.
**Auslöser:** Maintainer-Nachfrage vom 2026-07-24 zu
[slice-042 §8](../done/slice-042-constructs-aufruf-monopol.md) („warum können andere Konsumenten das
nicht gebrauchen?"). Die anschließende Nachmessung über **alle fünf** lokalen
`.a-check.yml`-Konsumenten zeigt eine Fehlerklasse, die mit dem b-cad-P-Rest nichts zu tun hat.
**Bezug:** macht die bewusste fail-open-Grenze aus
[AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)
sichtbar (sie wird **nicht** verschoben); Extraktions-/Auflösungs-Vertrag
[SPEC-EXTRACT-001](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion);
Konfigurations-Vertrag
[AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml);
Determinismus [AC-QA-01](../../../../spec/lastenheft.md#ac-qa-01--determinismus).
[Roadmap](../in-progress/roadmap.md).

> **Hinweis:** Die Entscheide aus §3 sind in **§0 abgenommen**; §1–§6 stehen als Entwurfs- und
> Evidenz-Stand unverändert, §2a korrigiert die Evidenz aus §2. Es wurde **keine neue `AC-*`-ID**
> vergeben (Anlege-Prozess:
> [conventions §Anforderungs-Anlege-Prozess](../../../../harness/conventions.md#anforderungs-anlege-prozess)) —
> die Begründung steht in §0.

---

## 0. Abnahme (2026-07-25)

Auf Basis der Nachmessung (§2a) und der Aufwandsschätzung (§2b) entschieden:

| Entscheid | Abgenommen | Begründung |
|---|---|---|
| **Gestalt** | **(a) — nur die Diagnose**, Exit-Code unverändert | (d) kostet das 2–2,5-fache, und der Zuwachs ist Vertrags-Maschinerie für eine Strenge, die noch niemand angefragt hat. Die Diagnose ist der Wirkmechanismus; der Schalter nützt erst, *nachdem* sie den Konsumenten auf die Lücke gestoßen hat |
| **`strict_coverage`** | **vertagt** | eigener Folge-Slice. **Trigger:** ein Konsument, der die Diagnose sieht und sie als Gate will. Bis dahin kein Config-Schlüssel, keine neue AC-ID, keine Exit-Code-Logik |
| **Granularität** | **Datei-Liste**, stabil nach Pfad sortiert; ab **10** Dateien gekürzt mit **ausgewiesener** Restzahl | „Zähler + Zonen" war mit d-migrates angeblichen hunderten Dateien begründet — der Fall existiert nicht (§2a). Die Kürzung ist keine stille Kappung: die Restzahl steht in der Meldung |
| **Zonen-Bildung** | **entfällt** | war nur für die Aggregation nötig, die es jetzt nicht gibt |
| **Abgrenzung** | `composition_root`-Dateien zählen **nicht**; `exclude`-Dateien sind ohnehin nicht im Scan | wie in §3 vorgeschlagen |
| **Wo greift es** | **Quell-Seite** (gescannte Datei ohne Schicht) | die Ziel-Seite ist nicht sicher von repo-**externem** Code trennbar (§3) |
| **§5 Nebenbefund** | **mitgenommen** | der leere Quell-Schicht-Name (`wrong-direction:  -> ui`) ist dasselbe Symptom, unabhängig verifiziert und billig |

**Vertrags-Folge:** kein Lastenheft-Bump und **keine neue `AC-*`-ID** — eine advisory
stderr-Zeile ohne Exit-Code-Wechsel schärft das *Wie* der bestehenden CLI-Ausgabe
([AC-FA-CLI-001](../../../../spec/lastenheft.md#ac-fa-cli-001--aufruf-scan-wurzel-und-exit-codes)
nennt „Zusammenfassung auf stderr"), sie erweitert den Vertrag nicht. Präzedenz:
[ADR-0025](../../adr/0025-exclude-verzeichnis-prune.md) und
[ADR-0028](../../adr/0028-ziel-glob-schattenwurf.md) liefen ebenso spec-only. Es entsteht eine
**ADR** (die Entscheidung „ausweisen statt verschieben" braucht eine zitierbare Stelle) plus eine
Schärfung in [SPEC-CLI-001](../../../../spec/spezifikation.md#spec-cli-001--aufruf-scan-wurzel-und-exit-codes).

## 1. Die Fehlerklasse

Ein Konsument, dessen `languages`-Globs **weiter** reichen als seine `layers`-Globs, hat
gescannte Dateien in **keiner** Schicht. Das ist zulässig und oft gewollt — aber es erzeugt zwei
Symptome, die der Nutzer heute nicht als *eine* Ursache erkennt (Proben aus
[slice-042 §8.1](../done/slice-042-constructs-aufruf-monopol.md)):

| Seite | Verhalten heute | Bewertung |
|---|---|---|
| **Ziel-Seite** — ein Import löst sich auf **keine** Schicht auf | 0 Befunde, still | **fail-open** — dokumentierte Grenze, aber unsichtbar |
| **Quell-Seite** — eine schichtlose Datei importiert eine Schicht | `wrong-direction`, aber mit **leerem** Quell-Schicht-Namen (`wrong-direction:  -> ui`) | korrekt, aber kryptisch |
| Gegenprobe: Zone als Schicht deklariert | regulärer Befund (`core-impurity`) | zwischen stiller Lücke und gefundenem Verstoß steht **nur die Deklaration** |

**Kernpunkt:** Der Konsument kann das heute selbst reparieren — eine Zeile in `layers` (oder in
`exclude`, wenn die Zone nicht dazugehört). Er erfährt nur nirgends, dass die Lücke existiert.
Ein grünes Gate über einem teilweise ungeprüften Baum ist genau die Klasse „stille
Setzung", gegen die dieses Repo sonst fail-closed vorgeht.

## 2. Evidenz (Nachmessung 2026-07-24, alle lokalen Konsumenten)

| Konsument | gescannt, aber in keinem Layer-Glob |
|---|---|
| a-check (Dogfooding) | **0** — Baum vollständig gedeckt |
| b-cad | **0** — Baum vollständig gedeckt |
| d-check | **0** — Baum vollständig gedeckt |
| d-migrate | der gesamte `test/`-Kotlin-Baum (kein `exclude`, `languages: **/*.kt`) |
| m-trace | `apps/api/internal/storage/**`, `apps/api/scripts/coverage-overview/**` |

Der **m-trace-Fall** trägt den Slice: `apps/api/internal/storage` wird von der dortigen Config in
einer `tech`-Regel ausdrücklich als Architektur-Zone geführt (zulässiger `database/sql`-Halter),
hat aber keine Schicht. Ein Import aus `hexagon/domain` dorthin wäre exakt der Verstoß, den das
Gate finden soll — und bliebe still. **Heute latent** (nur `*_test.go` importieren die Zone, und
die sind dort per `exclude` aus dem Scan), also kein akuter Fehlbefund, aber eine unbemerkte
Lücke im laufenden Gate eines realen Konsumenten.

## 2a. Nachmessung 2026-07-25 — die Evidenz aus §2 war zur Hälfte falsch

Vor der Abnahme neu gemessen, diesmal mit einer **exakten Portierung der Glob-Semantik**
(`core.globToRegexp`: `**` → `.*`, `*` → `[^/]*`, verankert) statt einer Endungs-Näherung. Gezählt
wurden Dateien, die einem `languages`-Glob entsprechen, **nicht** von `exclude` erfasst sind und in
**keinem** `layers`-Glob liegen (`composition_root` ausgenommen, §3):

| Konsument | ungedeckt | gegen §2 |
|---|---|---|
| a-check, b-cad, d-check, belief-agent, HexSlice-Go-Beispiel | **0** | bestätigt |
| **d-migrate** | **0** | ✗ **§2 ist falsch** |
| **m-trace** | **2 Dateien** | ✓ bestätigt, aber viel kleiner als „Zonen" nahelegt |

**Korrektur 1 — d-migrate hat sehr wohl ein `exclude`.** §2 behauptet „der gesamte `test/`-Kotlin-Baum
(kein `exclude`, `languages: **/*.kt`)". Tatsächlich enthält die dortige `.a-check.yml` seit ihrem
**ersten** Commit (2026-07-06, unverändert bis heute) den Block
`exclude: ["**/src/test/**", "**/src/testFixtures/**", "test/**", "**/build/**"]` — genau der
behauptete Baum ist ausgeschlossen und wird nie gescannt. Die Aussage war also nicht *veraltet*,
sondern **von Anfang an unzutreffend**; die Erstmessung hat den `exclude`-Block übersehen.

**Korrektur 2 — m-trace sind zwei Dateien, keine „Zonen".**
`apps/api/internal/storage/migrate.go` und `apps/api/scripts/coverage-overview/main.go`. Die
`*_test.go`-Dateien derselben Verzeichnisse sind per `exclude` ohnehin draußen.

**Was von der Evidenz übrig bleibt — und was nicht:**

- Der **scharfe Fall bleibt scharf**: `apps/api/internal/storage` wird in m-traces `tech`-Regel
  ausdrücklich als zulässiger `database/sql`-Halter geführt (verifiziert) und hat trotzdem keine
  Schicht. Ein `hexagon/domain`→dorthin-Import bliebe still. Heute **latent**: die einzigen realen
  Importeure sind `*_test.go`-Dateien, und die sind ausgeschlossen.
- Die Begründung **„flotten-weite Evidenz"**, mit der Kandidat 2a entgated wurde
  ([slice-042 §8](../done/slice-042-constructs-aufruf-monopol.md)), trägt jedoch **nicht**: sie
  stützte sich auf zwei Konsumenten, und einer davon fällt weg. Real ist **ein** Konsument mit
  **zwei** Dateien, von denen eine (`scripts/coverage-overview`) ein Werkzeug-Skript ist und
  architektonisch kaum zählt.

**Der Nebenbefund aus §5 reproduziert unverändert** (Fixture gegen `a-check:dev`, Stand `main`):

```text
tools/gen.go:3: wrong-direction:  -> ui (x/hexagon/ui/y)
```

— an der Stelle des Quell-Schicht-Namens steht nichts. Das ist unabhängig von der Zähl-Frage und
für sich genommen ein klarer, billiger Fix.

## 2b. Aufwandsschätzung für Gestalt (d) — Diagnose **plus** `strict_coverage`

Am Code durchgezählt (Stand `main`), nicht geschätzt aus dem Gefühl.

### Was (d) anfassen muss

| Ort | Änderung | Größe |
|---|---|---|
| `spec/lastenheft.md` | **neue Anforderung** (Bereich `CONF` oder `CLI` — der Entscheid steht noch aus) mit drei AK + Out-of-Scope, Versions-Bump **0.23.0**, Historie | ~30 Zeilen, **neue AC-ID** |
| `docs/plan/adr/` | Folge-ADR: warum **opt-in** statt Default (ein Default-Exit-1 bräche bestehende grüne Gates), warum nur die **Quell**-Seite, wie die Zone gebildet wird | ~90 Zeilen |
| `spec/spezifikation.md` | [SPEC-CONF-001](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema) (neuer Schlüssel), [SPEC-CLI-001](../../../../spec/spezifikation.md#spec-cli-001--aufruf-scan-wurzel-und-exit-codes) (stderr-Ausgabe **und** Exit-Code-Bedingung), [SPEC-DET-001](../../../../spec/spezifikation.md#spec-det-001--determinismus-vertrag) (stabile Zonen-Ordnung) + Historie | ~25 Zeilen, **3 Abschnitte** |
| `internal/adapter/driven/config` | `strict_coverage` im Schema + Model-Feld. Die Negativ-AK („Nicht-Bool ⇒ Exit 2") fällt **gratis** an — der strikte Decoder scheitert schon am Typ | ~5 Zeilen |
| `internal/hexagon/core` | Abdeckungs-Ermittlung + Zonen-Bildung, stabil sortiert. Die Daten liegen bereits vor (`FileImports.Layer == ""`), es braucht keine Extraktions-Änderung | ~60 Zeilen |
| `internal/cli` | Diagnose ausgeben + Exit-Code-Kombination (darf eine 1 nie zu 0 machen) | ~15 Zeilen |
| Tests | AK-Tests (Diagnose da/nicht da, `strict` ⇒ Exit 1, Nicht-Bool ⇒ Exit 2, Determinismus) + Unit-Tests der Zonen-Bildung | ~120 Zeilen |
| Handbuch, CHANGELOG, Roadmap | neuer Abschnitt + Glossar + Historie | ~35 Zeilen |

### Was **nicht** nötig ist (geprüft)

- **Keine Port-/Architektur-Änderung.** Die Diagnose muss nicht durch `ReportPort` — die
  Composition Root schreibt schon heute nach `errw` (`a-check: %v`). Damit entfällt ein
  `spec/architecture.md`-Bump ([ARC-005](../../../../spec/architecture.md)).
- **Kein `image-test`-Bruch.** Block (4) vergleicht stdout **und stderr** byte-identisch
  nativ↔Container; sein Fixture ist aber vollständig gedeckt, es entstünde dort keine Diagnose.
  Die Byte-Identität hielte ohnehin, solange die Zonen-Ordnung deterministisch ist.
- **Keine Extraktions-Änderung.** `LayerOf` liefert schon `""` für ungedeckte Dateien.

### Der Vergleich

| | Gestalt (a): nur Diagnose | Gestalt (d): Diagnose + `strict_coverage` |
|---|---|---|
| Lastenheft | **kein Bump** möglich (advisory stderr-Zeile ohne Exit-Code-Wechsel = Schärfung des *Wie*, Präzedenz [ADR-0025](../../adr/0025-exclude-verzeichnis-prune.md)) | **neue AC-ID** + Bump 0.23.0 |
| ADR | ja (eine Entscheidung: Ausweisen statt Verschieben) | ja, mit **zusätzlicher** Opt-in-/Gate-Begründung |
| Spezifikation | 1 Abschnitt ([SPEC-CLI-001](../../../../spec/spezifikation.md#spec-cli-001--aufruf-scan-wurzel-und-exit-codes)) | 3 Abschnitte |
| Config-Vertrag | **keiner** | neuer Schlüssel + fail-closed |
| Exit-Code-Logik | **keine** | ja (Kombination mit dem Befund-Exit) |
| Code + Tests | ~110 Zeilen | ~200 Zeilen |

**Verhältnis: (d) ist rund das 2–2,5-fache von (a)** — und der Zuwachs ist fast vollständig
**Vertrags-Maschinerie** (neue AC-ID, Config-Schlüssel mit Validierung, Exit-Code-Semantik, eine
ADR-Begründung für eine Strenge, die noch niemand angefragt hat), nicht Erkennungs-Logik. Die
Erkennung selbst ist in beiden Gestalten dieselben ~60 Zeilen.

### Zwei Punkte, die die Schätzung relativieren

1. **Die Zonen-Bildung — der einzige noch offene Entscheid (§3) — verliert mit §2a ihren Anlass.**
   „Zähler + Zonen" war damit begründet, dass d-migrates `test/`-Baum „hunderte Dateien" umfasse
   und eine Datei-Liste die Meldung unlesbar mache. Dieser Fall existiert nicht. Bei **zwei**
   Dateien ist die Datei-Liste die Meldung — kein Aggregations-Entscheid nötig, kein
   Determinismus-Risiko jenseits einer Pfad-Sortierung.
2. **`strict_coverage: true` würde m-traces grünes Gate rot machen** (2 Dateien). Das ist als
   Opt-in zulässig, heißt aber: der Schalter nützt erst, wenn der Konsument die zwei Dateien
   deklariert oder ausschließt — also **nachdem** ihn die Diagnose darauf gestoßen hat. Die
   Diagnose ist der Wirkmechanismus; die Strenge ist die Kür danach.

## 3. Zu entscheiden vor der Umsetzung (ADR-Skizze)

| Frage | Optionen | Erste Neigung |
|---|---|---|
| **Gestalt** | (a) immer eine Diagnose-Zeile auf stderr · (b) Opt-in `strict_coverage: true` ⇒ Exit 1 · (c) eigenes Flag (`--print-coverage`) · (d) (a)+(b) | **(d)**: die Diagnose sieht jeder sofort (das ist der Zweck), die Strenge bleibt Konsumenten-Wahl. Ein Default-Exit-1 würde bestehende grüne Gates brechen — Verschärfung ohne ADR wäre unzulässig ([AGENTS §3.6](../../../../AGENTS.md#36-gates-dürfen-nicht-ohne-adr-gelockert-werden) sinngemäß in der Gegenrichtung) |
| **Granularität** | je Datei · je Verzeichnis-Zone · Zähler + Zonen | **Zähler + Zonen** — d-migrates `test/`-Baum sind hunderte Dateien; eine Datei-Liste macht die Meldung unlesbar und die Diagnose wertlos |
| **Zonen-Bildung** | längster gemeinsamer Verzeichnis-Präfix · erste N Segmente · Verzeichnis der Datei | offen — muss **deterministisch und stabil sortiert** sein ([AC-QA-01](../../../../spec/lastenheft.md#ac-qa-01--determinismus), [SPEC-DET-001](../../../../spec/spezifikation.md#spec-det-001--determinismus-vertrag)) |
| **Abgrenzung** | `composition_root`-Dateien sind **per Definition** schichtlos; `exclude`-Dateien sind gar nicht im Scan ([ADR-0018](../../adr/0018-exclude-scan-scope.md)/[ADR-0025](../../adr/0025-exclude-verzeichnis-prune.md)); repo-**externe** Import-Ziele (Fremdbibliotheken) sind kein Befund | beide ersten Klassen **ausnehmen**; die dritte ist der Grund, warum die **Ziel**-Seite nicht einfach fail-closed werden kann — extern und „schichtlos im eigenen Baum" sind auf Symbol-Ebene nicht sicher unterscheidbar |
| **Wo greift die Strenge** | Quell-Seite (gescannte Datei ohne Schicht) · Ziel-Seite (Import ohne Schicht) | **Quell-Seite** — sie ist die *Ursache* und ohne Fehlalarm-Risiko entscheidbar. Die Ziel-Seite ist die Symptom-Seite und mit Externem vermischt; sie wird über die Ursache mit-repariert (Gegenprobe §1) |
| **Verhältnis zu [AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)** | — | die Grenze wird **nicht verschoben**, sondern *ausgewiesen* — das ist wörtlich der Vertrag („die Grenze wird ausgewiesen, nicht als Vollständigkeit ausgegeben") |

## 4. Anforderungs-Skizze (Bereich offen — `CONF` oder `CLI`, Entscheid im CR)

- **Happy:** Given ein Baum mit gescannten Dateien außerhalb aller Layer-Globs, when a-check
  läuft, then nennt die Ausgabe Anzahl und Zonen dieser Dateien — deterministisch sortiert,
  zusätzlich zu (nicht anstelle von) den regulären Befunden; der Exit-Code bleibt unverändert.
- **Boundary:** Given ein vollständig gedeckter Baum (a-check/b-cad/d-check), when a-check läuft,
  then **keine** Diagnose-Zeile (kein Rauschen im Normalfall). Given `composition_root`-Dateien,
  then zählen sie **nicht** als ungedeckt. Given `strict_coverage: true` und eine ungedeckte
  Zone, then Exit 1.
- **Negative:** Given `strict_coverage` mit einem Nicht-Bool-Wert, when a-check lädt, then Exit 2
  (fail-closed, Muster der bestehenden Config-Härtung).
- **Out-of-Scope:** keine Auto-Inferenz von Schichten aus Verzeichnisnamen; keine Strenge auf der
  **Ziel**-Seite (Externes ist nicht sicher abgrenzbar, §3); keine Änderung der
  Auflösungs-Semantik ([SPEC-EXTRACT-001](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion));
  keine Import-Allowlist (das ist Kandidat 2b, weiter gated —
  [slice-042 §8.2](../done/slice-042-constructs-aufruf-monopol.md)).

## 5. Mitzunehmen: leerer Quell-Schicht-Name

Der in [slice-042 §9](../done/slice-042-constructs-aufruf-monopol.md) notierte Mikro-CR
(`wrong-direction:  -> ui` — Loch an der Stelle des Quell-Schicht-Namens) ist das **Quell-seitige
Symptom derselben Klasse** und gehört hierher: die Meldung soll die Datei als schichtlos
ausweisen statt einen leeren Namen zu rendern. Klein, aber nutzersichtbar; Empfehlung: in diesem
Slice mitnehmen, damit beide Symptome dieselbe Sprache sprechen.

## 6. Betroffene Module und DoD

| Modul | Änderung |
|---|---|
| `internal/adapter/driven/config` | optionaler Schlüssel (`strict_coverage`), fail-closed-Validierung |
| `internal/hexagon/core` | Abdeckungs-Ermittlung + Zonen-Bildung, stabil sortiert; Meldungstext für schichtlose Quelldateien (§5) |
| `internal/cli` | Diagnose-Ausgabe, Exit-Code-Logik ([SPEC-CLI-001](../../../../spec/spezifikation.md#spec-cli-001--aufruf-scan-wurzel-und-exit-codes)) |
| Doku | Lastenheft-CR, ADR, [SPEC-CONF-001](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema)/[SPEC-DET-001](../../../../spec/spezifikation.md#spec-det-001--determinismus-vertrag), Benutzerhandbuch, [CHANGELOG](../../../../CHANGELOG.md), Roadmap |

**DoD** (Gestalt (a), nach §0): Spec-first (**kein** Lastenheft-CR → ADR → Spezifikation → Code →
Tests); **Review-Synthese unter [`docs/reviews/`](../../../reviews/)** nach Regelwerk Modul 10
(kategorisierte Findings + Negativbefund-Zeilen; die DoD-Verifikation bleibt nach Modul 11
davon getrennt); `make gates` **und**
`make ci` grün mit echter Ausgabe; **Regressions-Probe gegen alle fünf lokalen Konsumenten**:
a-check/b-cad/d-check bleiben diagnose-**frei** (sonst ist die Meldung Rauschen), m-trace und
d-migrate melden genau die Zonen aus §2; Gegenprobe: Zone deklarieren ⇒ Diagnose verschwindet und
ein injizierter Verstoß in der Zone wird gefunden; Benutzerhandbuch-Currency. Die Konsumenten
werden **informiert**, nicht umkonfiguriert — die Config-Entscheidung (Schicht vs. `exclude`)
liegt in ihren Repos.

---

## 7. Ergebnis (2026-07-25)

Spec-first geliefert: [ADR-0029](../../adr/0029-abdeckungs-diagnose-advisory.md) `Accepted` →
Spezifikation **0.26.0** ([SPEC-CLI-001](../../../../spec/spezifikation.md#spec-cli-001--aufruf-scan-wurzel-und-exit-codes))
→ Code → Tests; Benutzerhandbuch **1.35**, [CHANGELOG](../../../../CHANGELOG.md) unter
`[Unreleased]`, **Review-Synthese** unter [`docs/reviews/`](../../../reviews/2026-07-25-slice-043-abdeckungs-diagnose.md).
**Kein Lastenheft-Bump, keine neue `AC-*`-ID** (§0).

### 7.1 Verifikation (DoD-/Spec-Abgleich, nach Regelwerk Modul 11 getrennt vom Review)

| DoD-Punkt (§4/§6) | Beleg |
|---|---|
| Diagnose nennt Anzahl **und** Pfade, deterministisch sortiert, zusätzlich zu den Befunden | Test `TestCoverageNotice` + `TestCoverageNoticeDeterministic` (zwei Läufe byte-identisch, Pfade in Pfad-Ordnung) |
| Exit-Code bleibt unverändert | `TestCoverageNotice` pinnt Exit **0** bei ungedeckten Dateien ohne Befund |
| Vollständig gedeckter Baum ⇒ **keine** Diagnose (kein Rauschen) | `TestCoverageNoticeAbsentWhenCovered`; Konsumenten-Probe: a-check, b-cad, d-check, d-migrate, belief-agent, HexSlice-Go-Beispiel **diagnose-frei** |
| `composition_root`-Dateien zählen nicht | `TestUncoveredFiles` + `TestCoverageNotice` (Gegenprobe auf `cmd/main.go`) |
| Kürzung nennt die Restzahl | `TestCoverageNoticeCapNamesRemainder` (13 Dateien ⇒ „… und 3 weitere") |
| §5: leerer Quell-Schicht-Name | `TestWrongDirectionLayerlessSourceLabel` — Meldung beginnt mit `(ohne Schicht) -> ` |
| Regressions-Probe gegen alle lokalen Konsumenten | alle sieben `exit=0`, unveränderte Befundlage; **m-trace** nennt genau seine zwei Dateien |
| `make gates` **und** `make ci` grün | `make ci` grün: lint 0 issues, Coverage **96,20 %**, `arch-check` 0, `doc-check` 0, `image-test` OK |
| Benutzerhandbuch-Currency | 1.35, §2 „Das Ergebnis verstehen" |

**Nicht erfüllt und bewusst so:** der DoD-Punkt „`strict_coverage: true` ⇒ Exit 1" aus dem
Entwurfs-§4 entfällt mit der Abnahme (§0) — er gehört in den vertagten Folge-Slice.

### 7.2 Lerneinträge

1. **Konsumenten-Evidenz gehört gemessen, nicht gelesen.** Die Behauptung „d-migrate hat kein
   `exclude`" stand einen Tag lang in zwei Slices und hätte die Gestalt der Lösung bestimmt. Sie
   fiel erst, als die Glob-Semantik **exakt** nachgebaut wurde — die erste Näherung über
   Datei-Endungen war selbst falsch (sie zählte bei b-cad 41 Dateien, die nie gescannt werden).
   Wer über fremde Configs argumentiert, muss deren Semantik reproduzieren, nicht überfliegen.
2. **Eine korrigierte Evidenz kann den Scope halbieren, ohne den Slice zu erledigen.** Die
   Fehlerklasse blieb real; nur ihre Verbreitung schrumpfte — und damit die angemessene
   Maschinerie. „Kleiner bauen" war hier die richtige Antwort, nicht „vertagen" und nicht
   „wie geplant bauen".
3. **Eine Diagnose ist nur so viel wert, wie sie schweigt.** Dass sechs von sieben Konsumenten
   nichts sehen, ist kein Nebenergebnis, sondern die Bedingung dafür, dass die siebte Meldung
   gelesen wird. Das gehört in die Fitness Function, nicht in die Hoffnung.

## 8. Closure-Notiz (2026-07-25)

**Abgeschlossen und nach `main` gemergt.** DoD erfüllt (§7.1), Review-Synthese vor dem Merge
abgelegt ([`docs/reviews/`](../../../reviews/2026-07-25-slice-043-abdeckungs-diagnose.md)).

**Was der Slice bewusst *nicht* tut:** er gatet nichts. Die Abdeckungs-Lücke bleibt fail-open — sie
wird nur sichtbar. Die Opt-in-Strenge (`strict_coverage` ⇒ Exit 1) ist ein eigener Folge-Slice mit
dem Trigger aus [ADR-0029](../../adr/0029-abdeckungs-diagnose-advisory.md): **ein Konsument, der
die Diagnose sieht und sie als Gate will.** Bis dahin gibt es keinen Config-Schlüssel.

**Nachlauf beim Konsumenten:** m-trace sieht ab dem nächsten Release den Hinweis auf
`apps/api/internal/storage/migrate.go` und `apps/api/scripts/coverage-overview/main.go`. Die
Entscheidung — Schicht deklarieren oder `exclude` — liegt dort, nicht hier. Der erste Fall ist der
architektonisch relevante: `internal/storage` wird in der dortigen `tech`-Regel bereits als Zone
geführt.

**Release-Hinweis:** die Änderung liegt in [CHANGELOG](../../../../CHANGELOG.md) `[Unreleased]`,
zusammen mit `construct-leak` ([slice-042](slice-042-constructs-aufruf-monopol.md)) und dem
Ziel-Glob-Rückzug ([slice-044](slice-044-ziel-glob-schattenwurf.md)) — drei Auslieferungen warten
auf denselben Schnitt.
