# slice-184 — `d-check` auf `v0.75.0`, und das neue Modul `mentions` schließt eine benannte Grenze

**Welle:** ohne Welle — reaktiv (neue Werkzeug-Version verfügbar); die
Closure-Bedingung wäre die eigene DoD (Baseline-Regelwerk
`modul-06-roadmap.md` §Wann Arbeit eine Welle braucht).

**Bezug:** [`AC-QA-02`](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)
(Modul-Integrität des Doku-Stacks), [`AC-QA-03`](../../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit)
(Digest-Pin). Keine aktive ADR wird berührt.

**Berührte Spec-Stellen:** — (Werkzeug-Pin und Gate-Konfiguration ohne Vertragsberührung).

**Verantwortlich:** Claude — gesetzt beim Übergang nach `in-progress/`.

**Autor:** Claude. **Datum:** 2026-09-08.

**Lerneintrag — Form:** neuer Sensor.

---

## 1. Ziel und Abgrenzung

**Ziel:** Der `d-check`-Pin steht auf `v0.75.0`, und dessen einzige Neuerung —
das Modul **`mentions`** — wächtert die **Gegenrichtung** des Link-Checks: ob
eine *existierende* Datei irgendwo genannt wird.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Andere Doku-Familien als `harness/sensors/` und die ADRs.** Für beide ist
  der Bedarf **benannt** — die Ziel-Form von `harness/README.md` nennt die Lücke
  wörtlich, und `gate-consistency` trägt für die ADRs eine Eigenbau-Prüfung.
  Für Slice-Pläne, Carveouts oder Beobachtungen gibt es keinen solchen Anlass;
  ein Wächter ohne Anlass ist selbst eine Behauptung.
- **Die `citations`- und `sources`-Module.** Sie waren schon in `v0.74.1`
  verfügbar und sind unkonfiguriert — das ist Bestand, der bewusst stehen
  bleibt, und hat mit diesem Sprung nichts zu tun.
- **Produkt-Code.** Der Slice rührt `internal/` nicht an; er arbeitet in der
  Gate-/Werkzeug-Schicht.

## 2. Ausgangsmessung (2026-09-08)

### 2.1 Was der Sprung `v0.74.1 → v0.75.0` bringt

**Gemessen** über `--print-config` beider Digests, Diff über die volle Ausgabe:
**vier Zeilen**, und alle vier sagen dasselbe — die Modul-Liste bekommt
`mentions`. Kein anderer Konfigurations-Block ändert sich.

Der `--print-mk`-Diff zeigt die Folge: **sechs** der zwölf Recipe-Zeilen bekommen
zusätzlich `--disable mentions` — die sechs, die ohnehin eine geschlossene
Modul-Liste führen. Die erste Fassung schrieb hier *„jedes Target"*; das war
nicht gemessen (Review F-2).
**Die Folgerung bleibt trotzdem:** Das Modul ist strikt opt-in — aber sie trägt
auf `modules:` in `.d-check.yml`, nicht auf den `--disable`-Flags.

**Neuer Digest:** `sha256:18e9cd857f8db3569526d1f9a3cbeba8af51e9f2dd84c17a22444028b977c3da`
(alt: `sha256:e31a372b66dbde26305982424854cfce7c9ab7ce555a94debeee7ee26e6d4641`).

### 2.2 Was `mentions` tut — ausprobiert, weil `--print-config` es nicht sagt

Das Modul hat **keinen** dokumentierten Block in `--print-config`; es kommt dort
genau einmal vor, in der Verfügbar-Liste. Ermittelt gegen eine eigene Fixture:

```
art/zwei.md:1   docs/**/*.md   artifact-unmentioned
                kein Vorkommen als "art/zwei.md" in der Ist-Menge
```

- Es braucht **beide** Listen, sonst bricht es ab:
  `das Modul mentions braucht mentions.artifacts UND mentions.documents
  (DC-FA-MENT-001, fail-closed)`.
- Beide Felder sind **Listen** (`["glob"]`), nicht Einzel-Strings — ein String
  bricht mit `cannot unmarshal !!str into []string`.
- Die Erfolgszeile nennt die Deckung: *„1 von 2 Artefakt(en) erwähnt, über 1
  Dokument(e)"* — also **nicht** nur ein Ja/Nein, sondern eine Quote.

### 2.3 Warum a-check das braucht — die Grenze ist benannt, nicht erfunden

`v6.5.0` · `templates/harness/README.template.md` beschreibt sie wörtlich:

> *„Seine Grenze: Er prüft EINE Richtung — ob das Ziel existiert; **eine Datei
> ohne Index-Zeile** und eine Zeile auf die falsche Datei bleiben still grün."*

Dieselbe Grenze steht seit [slice-181](../done/wellenlos/slice-181-tabellenzellen-gewaechtert.md)
in `harness/sensors/gate-consistency.md`, und dort trägt eine **Eigenbau-Prüfung**
sie für die ADRs (ADR-Index-Vollständigkeit, slice-087).

**Der Bestand ist heute sauber** — 0 von 15 Sensor-Dateien und 0 von 39 ADRs
sind unverlinkt. Die Lücke ist **latent**, nicht offen. Das ist der richtige
Zeitpunkt für einen Wächter: Er kostet nichts und fängt den ersten Fall, statt
ihn zu finden, wenn er schon steht.

## 3. Umsetzung

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| `d-check.mk` | update | verbatim aus `v0.75.0 --print-mk`, Digest gesetzt |
| `.d-check.yml`, Block `mentions` | neu | `harness/sensors/*.md` gegen `AGENTS.md` |
| `Makefile`, `doc-mentions` | neu | das Fragment liefert kein Target; ins `gates`-Aggregat und `.PHONY` |
| `harness/sensors/doc-mentions.md` | neu | vier Grenzen, alle gemessen |
| `AGENTS.md` §4, `harness/README.md` §Sensors | update | Deklaration, sonst meldet `doc-targets` |
| `.claude/hooks/pretooluse-command-guard.sh` | update | die `GATES`-Liste des Guards |

**Der `guard-selftest` fand die dritte Stelle.** Ein neues Gate muss an **drei**
Orten bekannt sein: Aggregat, `.PHONY` — und der `GATES`-Liste des
Command-Guards. Fehlt die dritte, greift die Pipe-Regel für dieses Target nicht,
und `make doc-mentions | tail` liefe ungehindert durch. Der Selbsttest meldete
das von selbst; ohne ihn wäre es eine stille Lücke gewesen.

**Zwei Mutations-Proben, beide rot** — die zwei Fälle, die die Ziel-Form als
Lücke benennt:

| Mutation | Erwartet | Gemessen |
|---|---|---|
| Index-Zeile einer bestehenden Sensor-Datei entfernt | rot | `artifact-unmentioned` für `harness/sensors/symlink-check.md` |
| neue Sensor-Datei ohne Index-Zeile angelegt | rot | `artifact-unmentioned` für `harness/sensors/probe-ohne-zeile.md` |
| unverändert | grün | `16 von 16 Artefakt(en) erwähnt` |

### 3.1 Liefer-Punkt 3 fällt gemessen **negativ** aus

Die DoD verlangt, *entschieden und gemessen* zu haben, ob `mentions` die
ADR-Index-Eigenbau-Prüfung in `gate-consistency` ablöst. **Sie tut es nicht**,
und der Grund ist die Pfad-Form:

| | |
|---|---|
| Sensor-Dateien mit vollem Pfad in `AGENTS.md` | **15 von 15** |
| ADRs mit vollem Pfad im ADR-Index | **0 von 39** — der Index verlinkt geschwister-relativ (`](0001-….md)`) |

`mentions` sucht den **vollen repo-relativen Pfad**; es meldete alle 39 ADRs als
unerwähnt. Die Ablösung scheitert nicht am Willen, sondern an einer Eigenschaft
des Index, die er aus gutem Grund hat — ein Markdown-Link ist dateirelativ, und
ein voller Pfad wäre dort schlicht falsch.

**Das ist der Unterschied zur Ablösung, die slice-079 für `doc-targets`
gemacht hat:** Dort war die Parität in beiden Richtungen messbar und gegeben.
Hier ist sie messbar und **nicht** gegeben — und genau deshalb steht die
Eigenbau-Prüfung weiter in `tools/gate-consistency.sh`.

## 4. Definition of Done

- [ ] Pin auf `v0.75.0`: `d-check.mk` **verbatim** aus `--print-mk` des neuen
      Digests, `DCHECK_DIGEST` gesetzt, alle Versions-Nennungen nachgezogen.
- [ ] `mentions` konfiguriert für `harness/sensors/**` gegen die zwei
      Index-Tabellen; Mutations-Probe in **beide** Richtungen — eine Datei ohne
      Index-Zeile meldet rot, der unveränderte Bestand grün.
- [ ] Entschieden und **gemessen**, ob `mentions` die ADR-Index-Eigenbau-Prüfung
      in `gate-consistency` ablöst — Parität in beiden Richtungen, wie slice-079
      sie für `doc-targets` gemessen hat.
- [ ] Unabhängiger Review durchgeführt (Report unter [`docs/reviews/`](../../../reviews/README.md)).
- [ ] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [ ] Beobachtungs-Register fortgeschrieben.
- [ ] Jedes Risiko aus §7 trägt einen Ausgang.

`make gates` und `make verify` grün. Ein öffentlicher Vertrag ist berührt:
`d-check.mk` und `.d-check.yml` sind Teil des Harness-Stacks, und `AGENTS.md` §4
nennt die Targets.

## 5. Trigger

**Start** (`open` → `in-progress`): WIP-Limit frei. Der Anlass — `v0.75.0`
verfügbar — ist eingetreten.

**Rückführungen — vorab benannt:**

- `in-progress` → `next` (zu groß): Zeigt die Paritäts-Messung, dass die
  Ablösung der Eigenbau-Prüfung einen eigenen Umbau braucht, wird sie
  abgetrennt und der Slice trägt nur Pin und Konfiguration.
- `in-progress` → `open` (blockiert): Bricht `v0.75.0` ein bestehendes Gate,
  bleibt der Pin auf `v0.74.1`, und der Bruch wird als CR-Text an `d-check`
  formuliert ([`AGENTS.md`](../../../../AGENTS.md) §5).

## 6. Closure-Trigger

DoD vollständig, `make gates` und `make verify` grün, Closure-Notiz mit
Lerneintrag. Danach Archivierung als wellenloser Slice
([`AGENTS.md`](../../../../AGENTS.md) §6).

## 7. Risiken und offene Punkte

- **`mentions` ist undokumentiert.** Sein Verhalten ist ausprobiert, nicht
  gelesen — was es bei Sonderfällen tut, war ungeprüft. — **Ausgang:**
  *entfallen*, gestrichen mit Begründung: Die genannten Sonderfälle **sind**
  jetzt geprüft. Inline-Code und nackte Prosa zählen als Erwähnung (Fixture mit
  drei Formen), der Teilpfad **nicht** (nur der volle repo-relative Pfad
  trifft), und beide Befunde stehen als Grenze in
  `harness/sensors/doc-mentions.md`. Was bleibt, ist die Groß-/Kleinschreibung —
  sie ist im Bestand gegenstandslos, weil alle Pfade klein sind.
- **Der Sprung kann ein bestehendes Gate brechen**, ohne dass der Diff der
  Konfiguration es zeigt: Ein Modul kann sein *Verhalten* geändert haben, wo
  seine *Optionen* gleich blieben. — **Ausgang:** *entfallen*, gestrichen mit
  Begründung: `make gates` lief auf dem neuen Pin **vor** jeder weiteren
  Änderung durch, Exit 0. Das ist der Nachweis, den das Risiko verlangt — und
  er ist der Grund, warum der Pin-Sprung ein **eigener Commit** vor der
  Konfiguration war.
- **Die Ablösung der Eigenbau-Prüfung könnte den Geltungsbereich verkleinern** —
  dieselbe Falle, die slice-079 für `doc-targets` gemessen und vermieden hat.
  — **Ausgang:** *entfallen*, gestrichen mit Begründung: Die Ablösung findet
  **nicht statt** (§3.1). Gemessen nennen **0 von 39** ADRs ihren vollen Pfad im
  Index; `mentions` meldete jede einzelne als unerwähnt. Das Risiko war richtig
  gestellt und hat seinen Zweck erfüllt — es hat die Messung erzwungen, die die
  Ablösung verhindert hat.

## 8. Closure-Notiz

**Lerneintrag — Form: neuer Sensor** (`make doc-mentions`).

- **Was hat funktioniert:** Den Sprung **vor** jeder Konfiguration als eigenen
  Commit fahren und `make gates` darauf laufen lassen. Das beantwortet das
  Risiko *„der Sprung bricht ein Gate"* mit einem Beleg statt mit einer
  Vermutung — und es trennt zwei Fehlerquellen, die sonst in einem Diff liegen.

- **Was ging anders als geplant — Liefer-Punkt 3 fällt negativ aus, und das ist
  ein Ergebnis.** Die DoD verlangte, *gemessen* zu entscheiden, ob `mentions`
  die ADR-Index-Eigenbau-Prüfung ablöst. Gemessen: **0 von 39** ADRs nennen
  ihren vollen Pfad im Index — er verlinkt geschwister-relativ, und das ist
  richtig so. Die Ablösung scheitert an einer Eigenschaft des Gegenstands, nicht
  am Werkzeug. **Der Unterschied zu slice-079**, wo dieselbe Frage positiv
  ausging: Dort war die Parität messbar **und** gegeben, hier messbar und
  **nicht** gegeben.

- **Dieselbe Messung hat die Konfiguration selbst geformt.** `mentions` sucht
  den vollen repo-relativen Pfad; damit fällt auch
  [`harness/README.md`](../../../../harness/README.md) aus der Dokument-Menge —
  es nennt dieselben Dateien geschwister-relativ und **kann das nicht ändern**,
  ohne seine Links zu brechen. Ohne die Vorab-Frage *„welche Schreibweisen hat
  mein Gegenstand?"* wäre die Regel mit 39 Fehlalarmen eingeführt worden.

- **Steering-Loop-Eintrag — neuer Sensor:** `make doc-mentions` wächtert die
  **Gegenrichtung** des Link-Checks: ob eine *existierende* Datei genannt wird.
  Die Lücke ist keine Erfindung — `v6.5.0` ·
  `templates/harness/README.template.md` benennt sie wörtlich (*„eine Datei ohne
  Index-Zeile … bleibt still grün"*), und
  `harness/sensors/gate-consistency.md` führt sie seit slice-181 als eigene
  Grenze. — liegt in `Makefile:doc-mentions` (`.d-check.yml`, Block `mentions`;
  im `gates`-Aggregat).
  Auslöser: die benannte Grenze der Ziel-Form, nicht ein Register-Eintrag — der
  Bedarf stand fest, bevor das Werkzeug ihn decken konnte.

- **Der `guard-selftest` fand die dritte Stelle.** Ein neues Gate muss an drei
  Orten bekannt sein: `gates`-Aggregat, `.PHONY` und die `GATES`-Liste des
  Command-Guards. Die dritte hatte ich vergessen; ohne den Selbsttest wäre
  `make doc-mentions | tail` still durchgelaufen. **Ein Sensor, der einen
  anderen Sensor vollständig hält** — das ist der Fall, für den die
  Durchsetzungsschicht gebaut ist.

- **Was der Review fand (1 HIGH, 2 MEDIUM, 3 LOW).** Sein Verdikt: die Mechanik
  trägt, alle vier `mentions`-Verhaltensaussagen treffen zu, beide
  Mutations-Proben reproduzieren, und **der negative Liefer-Punkt 3 hält
  adversarisch** — er hat gegen `documents: ["**/*.md"]` gegengeprüft, und selbst
  dann bleiben 19 von 39 ADRs unerwähnt.
  **F-1 (HIGH):** Die abschließende Aufzählung in
  [`AGENTS.md`](../../../../AGENTS.md) §4 nannte `doc-mentions` nicht, während
  [`harness/README.md`](../../../../harness/README.md) „im `gates`-Aggregat"
  sagte. Beim Nachmessen kamen zwei ältere Fehler derselben Zeile ans Licht:
  `doc-immutable` stand dort als „in `gates`", obwohl es nur CI-durchgesetzt ist,
  und `doc-reviews` fehlte seit slice-160. Nach der Korrektur sind genannte und
  reale Menge deckungsgleich — acht Targets, Differenz null.
  **F-2/F-3:** zwei Zahlen, die ich nicht gemessen hatte — *„jedes Target
  bekommt `--disable mentions`"* (sechs von zwölf) und *„15 von 15"* (der Lauf
  meldet 16, die Sensor-Datei zählt sich selbst mit).

- **Der Review fand auch etwas, das den Sensor stärker macht als gedacht (F-6):**
  `mentions` hat **vier** fail-closed-Sperren, nicht zwei. Die zwei fehlenden
  sind Leer-Mengen-Sperren: *„eine Deckungs-Aussage über null Mitglieder ist
  keine"*. **Das Modul kann nicht ohne Gegenstand grün melden** — es ist von
  Haus aus gegen die Klasse gehärtet, für die `make dcheck-phrase-selftest` eine
  eigene Korpus-Kontrolle bauen musste.

- **Beobachtungs-Register (`../observations/`):** ein neuer Eintrag,
  [`BEO-HARNESS/aggregat-aufzaehlung-hinkt-dem-makefile-hinterher`](../observations/BEO-HARNESS/aggregat-aufzaehlung-hinkt-dem-makefile-hinterher/observation.md)
  mit **zwei** Belegen (slice-160 nachgetragen, slice-184) — die Klasse der
  abschließenden Aufzählung neben einer maschinenlesbaren Quelle, die kein
  Sensor deckt.
  Die drei einschlägigen `GATE`-Einträge bekommen **keinen** Beleg: Sie haben
  den Slice **geformt** statt von ihm einen zu bekommen (§9) — das ist die
  Wirkung, die der Zähler beabsichtigt.

- **Folge-Slices:** keiner. Die zwei benannten Grenzen —
  `harness/README.md` ungewächtert, ADR-Index nicht ablösbar — sind
  Eigenschaften des Gegenstands, keine offenen Aufgaben; sie stehen in
  `harness/sensors/doc-mentions.md`.

- **Risiken aus §7:** drei, jedes mit genau einem Ausgang — dreimal *entfallen*
  mit Begründung, weil jedes Risiko genau die Messung erzwungen hat, die es
  auflöst.

- **Drei Paarungen** (Repo ohne Wellen-Betrieb) — geprüft **nach** dem `git mv`
  nach `done/`, weil sie dort suchen; alle drei tragen:
  **Anker** — `liegt in Makefile:doc-mentions`; `slice-184` steht dort sowie in
  [`.d-check.yml`](../../../../.d-check.yml) und
  `harness/sensors/doc-mentions.md`.
  **Folge-Slice** — keiner genannt; die zwei benannten Grenzen sind
  Eigenschaften des Gegenstands, keine offenen Aufgaben.
  **Register** — der neue Eintrag
  `aggregat-aufzaehlung-hinkt-dem-makefile-hinterher` existiert mit **zwei**
  Belegen; die drei in §9 zitierten `GATE`-Einträge tragen alle ein nicht leeres
  `evidence/`.

## 9. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** zwei Sub-Areas berührt.
**Gate-/Werkzeug-Schicht** `GATE` (`d-check.mk`, `.d-check.yml`, `Makefile`,
`.claude/hooks/`) und **Harness-Einstieg** `HARNESS`
([`AGENTS.md`](../../../../AGENTS.md) §4,
[`harness/README.md`](../../../../harness/README.md) §Sensors,
`harness/sensors/`) — beide Achsen 1,2,3 laut
[Modus-Deklaration](../../../../harness/conventions.md#modus-deklaration-pro-sub-area).

**Vorgelagert — offene Beobachtungen sichten:** gesichtet am 2026-09-08.
`BEO-GATE/` führt 21 Einträge; drei sind einschlägig, und **alle drei haben
diesen Slice geformt**:

- [`probe-liefert-den-gegenstand-mit`](../observations/BEO-GATE/probe-liefert-den-gegenstand-mit/observation.md)
  — **3×**, seit slice-181 *verkörpert* in
  [`AGENTS.md`](../../../../AGENTS.md) §5. **Kein neuer Beleg**, und das ist
  die Wirkung: Die Regel *„war sie rot, und woran?"* hat hier von Anfang an
  gegriffen — beide Proben sind rot gefahren und mit ihrer **Meldung**
  protokolliert, nicht nur mit dem Exit-Code.
- [`muster-trifft-nur-die-haeufige-schreibweise`](../observations/BEO-GATE/muster-trifft-nur-die-haeufige-schreibweise/observation.md)
  — **2×**, offen. **Kein neuer Beleg, aber knapp:** Die Frage *„welche
  Schreibweisen hat mein Gegenstand?"* ist hier **vor** der Konfiguration
  gestellt worden, und die Antwort hat sie geändert — `mentions` sieht nur den
  vollen Pfad, also fällt `harness/README.md` aus der Dokument-Menge und die
  ADR-Ablösung ganz weg. Ohne diese Frage wäre die Regel mit 39 Fehlalarmen
  eingeführt worden.
- [`pruefer-ohne-gegenstand-oder-aufruf`](../observations/BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf/observation.md)
  — **5×**, *verkörpert*. **Kein Beleg**: Die Kandidatenmenge ist nicht leer
  (16 Artefakte), und der Lauf nennt sie in seiner Erfolgszeile — das Modul
  liefert die Nichtleerheits-Aussage von selbst mit.

**Keine weiteren Treffer** für `GATE` oder `HARNESS`.

**Alle berührten Sub-Areas GF** — kein Begründungsblock nötig.
