# slice-174 — Regelwerk-Migration `v6.2.0` → `v6.5.0`: Delta-Analyse

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `make slice-mv` ([`AGENTS.md`](../../../../AGENTS.md)
§3.3/§5).

**Welle:** [welle-15](../welle-15-regelwerk-v650-migration.md)

**Bezug:** Maintainer-Hinweis 2026-09-07 („Neues Regelwerk-Release
`v6.5.0`"). Vorbilder derselben Form:
[slice-161](../done/slice-161-regelwerk-v610-delta-analyse.md) und
[slice-164](../done/slice-164-regelwerk-v620-delta-analyse.md).

**Berührte Spec-Stellen:** — *(keine)* — reine Ist-Messung ohne
Vertragsberührung.

**Verantwortlich:** — *(noch nicht priorisiert)*

**Autor:** Claude (Opus 5), im Auftrag des Maintainers. **Datum:**
2026-09-07.

---

## 1. Ziel

Das Delta `v6.2.0` → `v6.5.0` ist vermessen und je Änderung beantwortet, ob
sie a-check berührt — und wenn ja, mit welchem Folge-Slice. Kein Vendoring,
kein Nachzug: dieser Slice liefert die Grundlage, auf der die übrigen Etappen
von [welle-15](../welle-15-regelwerk-v650-migration.md) geschnitten werden.

## 2. Ausgangsmessung (vor der Analyse)

Am 2026-09-07 im frischen Klon von `pt9912/ai-harness-course`
(`git diff --stat v6.2.0 v6.5.0 -- lab/regelwerk lab/templates`):

| Kennzahl | Wert |
|---|---|
| Dateien | **30** |
| Zeilen | **+572 / −149** |
| Neu | `lab/templates/harness/sensors/gate.template.md` (+81) |
| Gelöscht | keine |
| Übersprungene Releases | `v6.3.0`, `v6.3.1`, `v6.4.0` |

**Größenordnung, nicht Increment.** Der vorige Sprung `v6.0.0` → `v6.2.0`
umfasste 9 Dateien und `+53/−5`
([slice-164](../done/slice-164-regelwerk-v620-delta-analyse.md) §2). Dieser ist
rund zehnmal so groß; die Analyse ist entsprechend nicht in einem Absatz zu
erledigen, und der Slice liefert bewusst **nur** sie.

Die größten Posten, nach `git diff --numstat` (die Zahlen beim Anlegen dieses
Plans stammten aus `--stat` und zählten `+`/`−` zusammen; **und die größte
Datei fehlte in der Liste** — der erste `--stat`-Aufruf war abgeschnitten):

| Datei | `+`/`−` | warum a-check betroffen ist |
|---|---|---|
| `grundlagen-harness-dateien.md` | **+123/−14** | die Datei, die [`AGENTS.md`](../../../../AGENTS.md) §3.7 zitiert; trägt die neue `harness/sensors/`-Regel |
| `templates/harness/sensors/gate.template.md` | **neu, +81** | Ziel-Form je Sensor, „sobald ein Target mehr braucht als einen Satz" |
| `modul-05-planning-harness.md` | +46/−1 | Slice-§1 trägt Ziel **und Abgrenzung** |
| `grundlagen-begriffe.md` | +42/−40 | Begriffs-Umbau |
| `grundlagen-traceability.md` | +41/−0 | neue zweite Traceability-Richtung (RTM) |
| `modul-13-quality-gates.md` | +40/−10 | Gate-Regeln; a-check führt 40 Targets |
| `templates/…/slice.template.md` | +36/−10 | trifft die Kopieranleitung in [`AGENTS.md`](../../../../AGENTS.md) §5 |
| `templates/harness/README.template.md` | +33/−0 | zwei Tabellen statt einer |

## 3. Analyse

### 3.1 Die Größe täuscht: 16 von 30 Dateien tragen überhaupt Inhalt

Der Roh-Diff nennt 30 Dateien und `+572`. Der Kurs hat in `v6.5.0` zwei
**Formatierungs**-Klassen mitgeführt, die keine Regel ändern: Tabellen-Trennzeilen
(`| --- | --- |`) und normalisierte **Zell-Innenabstände** ganzer Tabellen. Beides
fällt weg, wenn man den Diff whitespace-bereinigt liest:

| Messung | Dateien | `+`/`−` |
|---|---|---|
| roh (`git diff --numstat`) | 30 | +572 / −149 |
| **inhaltlich (`--numstat -w`)** | **16** | **+452 / −29** |

**14 Dateien haben null Inhaltsänderung** — `modul-02`, `modul-04`, `modul-07`,
`modul-08`, `modul-10`, `modul-11`, `modul-12`, `modul-14`, `modul-15`,
`grundlagen-bootstrap`, `grundlagen-durchsetzungsschicht`,
`grundlagen-klassifikation`, `grundlagen-referenz-richtung`,
`grundlagen-source-precedence`. Sie berühren a-check nicht.

**Korrektur aus dem Review (F-1), und sie trifft die Methode dieses Slice.**
Zuerst stand hier eine Bereinigung, die **nur** die 36 Trennzeilen abzog und
daraus „8 Dateien tragen 436 von 536 substanziellen Zeilen, 9 tragen null"
ableitete. Beide Zahlen waren zu hoch, die Klassifikation zu grob: 84 weitere
`+`-Zeilen sind inhaltsgleich, und **fünf** Dateien, die dieser Slice unter
§3.3 als *„Tabellen-Ausbau und Verweise, zu prüfen im Adaptions-Durchgang"*
führte, haben in Wahrheit **gar keine** Inhaltsänderung. Der Lerneintrag dieses
Slice — *den Diff vor der Umfangs-Beurteilung um formatierende Anteile
bereinigen* — war richtig formuliert und **halb ausgeführt**: eine
Formatierungs-Klasse gesehen, die zweite nicht gesucht. Das ist die Klasse
[`BEO-PLAN/review-geltungsbereich-zu-eng`](../observations/BEO-PLAN/review-geltungsbereich-zu-eng/observation.md),
und sie hat mit diesem Slice ihr drittes Auftreten.

**Besonders folgenreich war der Irrtum bei zwei Dateien:**
`grundlagen-referenz-richtung.md` und `grundlagen-source-precedence.md` tragen
die `Ersetzt-Baseline-Regel`-Anker von
[`MR-011`](../../../../harness/conventions.md#mr-011) und
[`MR-012`](../../../../harness/conventions.md#mr-012). Der Slice schrieb, sie
seien im Adaptions-Durchgang zu prüfen — sie sind **unverändert**, und beide
Adaptionen bleiben damit ungebrochen. Und `grundlagen-begriffe.md` steht mit
`+42/−40` im Roh-Diff, trägt aber genau **zwei** neue Zeilen: die Begriffe
*RTM* und `harness/sensors/<target>.md`. Die in einer früheren Fassung dieses
Abschnitts als „neu benannt" geführten *Pfad* und
*Vertrag › Technik › Sicht › ADR › Slice* gibt es dort nicht bzw. schon seit
`v6.2.0` — eine Behauptung aus einem `grep` über den unbereinigten Diff.

### 3.2 Fünf Themen, nach Wirkung auf a-check geordnet

#### T-1 — Zitier-Form in einfrierenden Artefakten *(die Wurzel unserer eigenen Klemme)*

Vier Ziel-Formen tragen neu einen **Zitier-Form**-Block, der ausdrücklich als
Norm markiert ist und beim Kopieren **stehen bleibt**: `archiv-stub-slice`,
`archiv-stub-welle`, `welle-results`, `review-report`. Die Regel: **Kennung,
nicht Adresse** — `slice-NNN` statt Lifecycle-Pfad, `make <target>` statt Link
auf die Sensor-Datei, eine Baseline-Stelle als `v<X.Y.Z>` ·
`regelwerk/<datei>.md` §Abschnitt **statt als Link**. Begründung wörtlich:
*„Der vendored Baum trägt genau einen Tag; der Sprung löscht den alten, und ein
Link darauf färbt beim nächsten Bump ein Artefakt rot, das niemand mehr
anfassen darf."*

**Das trifft die Klemme aus slice-172 — aber nicht die, die `exempt-paths`
behandelt.** Präzisierung aus dem Review (F-2); zuerst stand hier, der Kurs
löse „dieselbe" Klemme an der Wurzel, und das zog zwei Fragen zusammen, die
[slice-173](../done/wellenlos/slice-173-versions-sensor-baseline-pins.md) §2.1
gerade auseinandergehalten hatte:

| Frage | Wer antwortet heute |
|---|---|
| Nennt ein Dokument den **veralteten** Stand? | `versions` — und `exempt-paths` nimmt Zeitdokumente aus |
| **Löst** ein Pfad auf, nachdem ein Stand entfernt wurde? | die Link-Prüfung — und sie erzwingt den Edit an eingefrorenen Dateien |

Die Zitier-Form beantwortet die **zweite**: kein Link, kein Bruch, kein Edit.
`exempt-paths` beantwortet die **erste** und bleibt davon unberührt. Genau die
zweite ist in
[`conventions.md`](../../../../harness/conventions.md#baseline) §Baseline
ausdrücklich als **unbehandelt** deklariert — *„Eine Kollision bleibt und ist
keine Ausnahme, sondern eine Klemme … der Preis des Löschens"*. Die Zitier-Form
ist die erste Antwort darauf, die a-check bekommen könnte; ob `exempt-paths`
danach schrumpft, ist eine **eigene** Messung und keine Folge.

**Folge-Slice:** Zitier-Form übernehmen und gegen den Bestand halten; dabei
prüfen, ob `exempt-paths` danach schrumpfen kann.

#### T-2 — `harness/sensors/<target>.md` *(größter Umfang, 16 Treffer im Bestand)*

Neue Ziel-Form `gate.template.md` (+80) plus die Regel in
`grundlagen-harness-dateien.md` (+122) und `harness/README.template.md` (+32):
Die Sensors-Tabelle bleibt Index, **ein Gate bekommt eine eigene Datei, sobald
sein Vertrag mehr als einen Satz braucht** — Deckungsgrenze,
Ausgabe-Bedeutung, Exit-Codes, Abbruch-Bedingungen. Die Target-Zelle wird zum
**Link** auf die Datei; dieser Link ist die einzige Fassung der Zuordnung, die
ein Sensor prüfen kann.

Der Satz, der a-check direkt trifft: *„Ob der Überhang schon unter der Tabelle
steht oder in die Zelle gedrängt wurde, ist dieselbe Sache — eine Zelle, die
zum Absatz geworden ist, ist der Fund, nicht die Ausnahme."*

**Gemessen im Bestand** (Zeichen je Zweck-Zelle,
[`AGENTS.md`](../../../../AGENTS.md) §4): **16 von 40** Gate-Zeilen sind länger
als 250 Zeichen. Spitzenreiter `make doc-check` mit **1871**, dahinter
`make symlink-check` mit **1106** — angelegt gestern in
[slice-173](../done/wellenlos/slice-173-versions-sensor-baseline-pins.md), also
kein Altbestand, sondern laufende Praxis.

**Folge-Slice:** eigene Etappe, und die größte. Nicht alle 16 auf einmal — die
Vorlage nennt die Datei als Antwort auf einen *Überhang*, nicht als Pflicht je
Target.

#### T-3 — Slice-§1 trägt Ziel **und Abgrenzung**

`modul-05` (+46) und `slice.template.md` (+36): §1 heißt jetzt *Ziel und
Abgrenzung* und verlangt eine Out-of-Scope-Liste **je Punkt mit Begründung**,
in vier benannten Klassen (Folge-Slice mit Kennung · Bestand bleibt bewusst ·
wäre ein anderer Vorgang · Schicht-Abgrenzung). Dazu die Regel, dass ein
genannter Folge-Slice die Sendung auch *annehmen* muss.

a-check führt Out-of-Scope heute auf **Welle**-Ebene
([welle-15](../welle-15-regelwerk-v650-migration.md) §6), nicht je Slice. Die
Kopieranleitung in [`AGENTS.md`](../../../../AGENTS.md) §5 nennt §1 noch
schlicht *Ziel*.

**Und die Regel hat eine zweite Hälfte im Lauf** — Nachtrag aus dem Review
(F-3), zuerst falsch unter §3.3 abgelegt: `modul-09` (+9, inhaltlich) trägt
keinen Tabellen-Ausbau, sondern einen **Normabsatz** zum 8-Schritt-Workflow.
*„Die Plan-Ausgabe in Schritt 4 nennt Out-of-Scope … Nimmt der Lauf etwas mit,
das §1 ausschließt, ist das eine Plan-Änderung und gehört vor den Code, nicht
in den Bericht danach."* [`AGENTS.md`](../../../../AGENTS.md) §6 führt denselben
Workflow und nennt `modul-09` als Regelquelle — die Dokument-Hälfte (§1 des
Plans) und die Schritt-Hälfte (Schritt 4 des Laufs) gehören zusammen und in
**denselben** Folge-Slice.

**Etappe E, [slice-178](../in-progress/slice-178-slice-form-ziel-und-abgrenzung.md):**
Kopieranleitung in [`AGENTS.md`](../../../../AGENTS.md) §5 **und** §6 nachziehen.
Ob der Bestand nachgerüstet wird, ist eine eigene Frage — die Ziel-Form gilt für
**neue** Slices.

#### T-4 — Nicht-Gates gehören in eine **zweite Tabelle**

`modul-13` (+37) und `harness/README.template.md`: Ein Target, das nichts über
den Repo-Zustand *urteilt* — es *bewegt* (Slice-Move), *misst* oder *sagt*, was
ein schreibender Lauf täte —, steht in einer **zweiten Tabelle** und trägt
`kein Gate` **in der Zeile selbst**, nicht in Prosa daneben. Dazu die
Unterscheidung *„Vorhanden ≠ behauptet"*: ein reales Target nicht als Gate zu
führen, ist keine Harness-Lüge.

a-check führt `make slice-mv`, `make archive-wave` und `make regelwerk-check`
in **derselben** Tabelle wie die Gates, mit **kein Gate** als fett gesetztem
Prosa-Vorspann in der Zweck-Zelle — genau die Form, die die neue Regel ablöst.

**Folge-Slice:** zusammen mit T-2, weil beide dieselbe Tabelle umbauen.

#### T-5 — RTM: die zweite Traceability-Richtung

`grundlagen-traceability.md` (+41, neuer Abschnitt *Die zweite Richtung:
Anforderung → Beleg*): Die **Requirements Traceability Matrix** macht *Waisen*
sichtbar — Anforderungen ohne Beleg. Sie **wird erzeugt, nicht gepflegt**; eine
RTM als eigenes Dokument wäre eine Kopie, und Kopien driften.

a-check erfüllt das bereits: `make doc-trace` erzeugt sie (advisory),
`make doc-complete` macht die Waise abschluss-blockierend
([`AGENTS.md`](../../../../AGENTS.md) §4). **Eine Aussage ist neu:** *„Was eine
Anforderung entlastet, ist eine Setzung — und sie gehört aufgeschrieben."*
a-check deklariert nirgends, welche Verweis-Quellen als entlastend gelten.

**Folge-Slice:** klein, eventuell in den Adaptions-Durchgang gefaltet.

### 3.3 Was nicht berührt — und die zwei Nachzüge, die fast durchgerutscht wären

- **Die 14 Null-Inhalt-Dateien** (§3.1). Nichts zu prüfen, auch nicht im
  Adaptions-Durchgang: `grundlagen-referenz-richtung` und
  `grundlagen-source-precedence` tragen zwar die Anker von
  [`MR-011`](../../../../harness/conventions.md#mr-011)/[`MR-012`](../../../../harness/conventions.md#mr-012),
  sind aber **unverändert**.
- `grundlagen-begriffe.md`: zwei neue Begriffe (*RTM*,
  `harness/sensors/<target>.md`), die T-5 und T-2 begrifflich schärfen. Keine
  eigene Handlung.
- `modul-06` (+3): Verweis auf die neue §1-Form. Fällt mit T-3.
- `archiv-stub-slice`, `archiv-stub-welle`, `welle-results` (je +8): die
  Zitier-Form aus T-1. Keine eigene Klasse.

**Zwei Dateien standen in keiner Einordnung** — Review-Befund (F-5), und beide
tragen echte Nachzüge:

| Datei | `+`/`−` (`-w`) | Nachzug |
|---|---|---|
| `regelwerk/README.md` | +1/−1 | Der Kurs-Wellen-Stempel wandert **119 → 128** (2026-09-05 → 2026-09-06). [`conventions.md`](../../../../harness/conventions.md#baseline) §Baseline zitiert ihn **wörtlich** — gehört in Etappe A, sonst steht dort nach dem Vendoring eine falsche Welle-Nummer neben dem richtigen Tag |
| `templates/README.md` | +11/−6 | Index-Zeile für `harness/sensors/gate.template.md` (T-2) und die Beschreibung der neuen `slice.template.md`-Abschnitte §1/§8 (T-3). Kein eigenes Thema, aber der Ort, an dem beide Ziel-Formen erklärt werden |

**§8 der Slice-Ziel-Form heißt neu *Sub-Area-Prüfungen und Modus-Begründung*** —
der Titel trägt jetzt beide Hälften: die zwei *Vorgelagert*-Blöcke sind immer
auszufüllen, der Modus-Begründungsblock bleibt bedingt. a-checks Slices führen
§9 unter dem alten Titel *Sub-Area-Modus*; das fällt mit T-3.

### 3.4 Vorgeschlagener Etappen-Schnitt

| Etappe | Inhalt | Umfang |
|---|---|---|
| **A** — Vendoring | `v6.5.0` vendoren, Stand an drei Stellen, vier Symlinks, `SHA256SUMS` | S |
| **B** — Adaptions-Durchgang | sieben aktive `MR` gegen `v6.5.0`, dabei §3.3-Restposten und T-5 | M |
| **C** — Zitier-Form (T-1) | vier Artefaktklassen; danach `exempt-paths` prüfen | M |
| **D** — Sensors-Struktur (T-2 + T-4) | zwei Tabellen, `harness/sensors/` für den Überhang | **L** |
| **E** — Slice-Form (T-3) | §1 *Ziel und Abgrenzung* und §8-Titel in der Kopieranleitung `AGENTS.md` §5, dazu die Schritt-Hälfte in §6 | S |

**A vor allem anderen** — B, C, D und E messen gegen den vendorten Stand. C vor
D ist keine Pflicht, aber sinnvoll: die Zitier-Form betrifft Verweise auf
Sensor-Dateien, die D erst anlegt.

**Nicht als Etappe geschnitten, aber offen** (Review-Befund F-10): Der
Kurs-Absatz zu `exempt-paths` endet mit *„ein Ausnahme-Ventil … also eine
**Gate-Senkung mit eigener Begründungslast**"*.
[`AGENTS.md`](../../../../AGENTS.md) §3.6 verlangt für eine Prüfregel-Senkung
eine ADR; a-check hat mit
[slice-173](../done/wellenlos/slice-173-versions-sensor-baseline-pins.md) fünf
`exempt-paths`-Klassen gesetzt und **keine** ADR dazu — keine der 39 nennt
`exempt-paths` oder `version-stale`. Das ist kein Nachzug aus `v6.5.0`, sondern
eine offene Frage **an den eigenen Bestand**, die dieser Sprung sichtbar macht.
Sie gehört in Etappe B, wo die Adaptions- und Ausnahme-Lage ohnehin geprüft
wird.

## 4. Definition of Done

- [x] Jede der 30 geänderten Dateien ist eingeordnet: berührt a-check / berührt
      nicht / berührt nur eine Stelle, die a-check ohnehin anders löst — mit
      Begründung je Zeile, nicht als Sammelurteil.
- [x] Für jede Berührung steht ein Vorschlag da: Folge-Slice (mit Titel und
      Etappe) oder ausdrückliche Nicht-Handlung mit Grund.
- [x] Unabhängiger Review durchgeführt (Report unter [`docs/reviews/`](../../../reviews/README.md)).
- [x] `make gates` grün.
- [x] `make verify` grün.
- [x] Jedes Risiko trägt einen Ausgang.

## 5. Trigger

**Start** (`open` → `in-progress`): Maintainer-Freigabe und WIP-Limit frei.

**Rückführungen:** stellt sich heraus, dass die Analyse selbst in Etappen
zerfällt — etwa weil `modul-13` und die `sensors/`-Ziel-Form getrennt zu
bewerten sind —, zurück nach `next/` zur Zerlegung. Wird ein weiteres Release
veröffentlicht, bevor die Analyse steht: zurück nach `open/`, weil dann das
Ziel der Welle neu zu setzen ist (Präzedenz: der Retarget von `welle-14`).

## 6. Closure-Trigger

DoD vollständig, `make gates`/`make verify` grün, Closure-Notiz geschrieben.
Der Slice trägt ein `**Welle:**`-Feld und archiviert **mit seiner Welle**,
nicht einzeln ([`AGENTS.md`](../../../../AGENTS.md) §6).

## 7. Risiken und offene Punkte

- *Die Analyse misst den Diff, nicht die Wirkung — eine Änderung kann klein
  aussehen und einen Vertrag brechen, wie umgekehrt* — **Ausgang:** eingetreten,
  und zwar in beide Richtungen. Neun Dateien sahen nach Änderung aus und tragen
  **null** Substanz (§3.1); umgekehrt steckt in `grundlagen-traceability.md`
  (+41) ein Satz, der eine unentschiedene Setzung in a-check benennt (T-5).
  Aufgefangen durch die Trennung *Trennzeilen vs. Substanz*, die die Analyse
  gemessen statt geschätzt hat. Kein Folge-Slice: die Methode ist im Slice
  beschrieben und beim nächsten Sprung wiederholbar.
- *Ein weiteres Release erscheint während der Welle; `v6.2.0` und `v6.5.0`
  lagen nur zwei Tage auseinander* — **Ausgang:** weiter offen →
  Beobachtungs-Register. Der Präzedenzfall ist der Retarget von `welle-14`
  (`v6.1.0` → `v6.2.0`, zwei Stunden nach der Eröffnung); §5 dieses Slice trägt
  die Rückführung dafür, und die Welle-Datei ist auf den Zielstand hin
  formuliert, nicht auf eine Etappen-Kette.
- *Die neue `sensors/`-Ziel-Form könnte a-checks Tabellen-Praxis nicht nur
  ergänzen, sondern ihr widersprechen — dann ist es keine Nachzugs-, sondern
  eine Adaptions-Frage* — **Ausgang:** eingetreten, Folge-Slice
  [slice-177](../done/slice-177-sensors-struktur-zwei-tabellen.md) (§7 dort
  führt die Frage bereits offen). **Korrektur aus dem Review (F-7):** zuerst
  stand hier *entfallen*, gestützt auf die Default-Frage — die Vorlage nennt
  die Tabellenzeile als Default, also werde a-checks Praxis begrenzt statt
  abgelöst. Das beantwortet aber nicht die **Widerspruchs**-Frage. Gemessen:
  [`AGENTS.md`](../../../../AGENTS.md) §4 führt `| Target | Zweck |` — die
  Spalte, in der die Ziel-Form `kein Gate` verlangt, **existiert dort nicht**;
  [`harness/README.md`](../../../../harness/README.md) hat sie
  (`| Target | Vertrag | Bindung |`), a-checks zweite Sensor-Tabelle also nicht.
  Ob daraus ein Spaltenzuwachs folgt oder eine deklarierte Abweichung, ist
  offen — und damit genau die Adaptions-Frage, die der Ausgang *entfallen*
  verneint hätte.

## 8. Closure-Notiz

**Lerneintrag — Form: geschärfte Regel** (Delta-Analyse trennt Formatierung von
Substanz, bevor sie den Umfang beurteilt).

- **Was hat funktioniert:** Die Größenangabe zu **misstrauen**. `+572` über 30
  Dateien las sich wie eine Generalüberholung; die Zählung der reinen
  Tabellen-Trennzeilen (36) und die Aufteilung je Datei zeigten, dass **8**
  Dateien 436 der 536 substanziellen Zeilen tragen und **9** gar nichts. Ohne
  diesen Schnitt wäre der Etappen-Plan an der falschen Stelle groß geworden.

- **Was ging anders als geplant:** Der Plan nannte beim Anlegen vier größte
  Posten — und die **größte Datei fehlte**. `grundlagen-harness-dateien.md`
  (+123/−14) stand nicht in der Liste, weil der erste `--stat`-Aufruf mit
  `tail -25` abgeschnitten war und ausgerechnet die erste Zeile verschluckte.
  Dieselbe Datei, die [`AGENTS.md`](../../../../AGENTS.md) §3.7 zitiert. Zudem
  waren die Δ-Zahlen `--stat`-Summen (`+`/`−` addiert) und nicht die
  `+`-Zahlen, als die sie dastanden. Beides in §2 korrigiert; die Analyse hat
  danach `--numstat` je Datei verwendet, nicht eine Übersicht mit `head`/`tail`.

- **Steering-Loop-Eintrag — geschärfte Regel:** Ein Roh-Diff wird **vor** der
  Umfangs-Beurteilung um seine formatierenden Anteile bereinigt — **Plural**,
  und die Vollständigkeit der Klassen ist selbst zu prüfen: `git diff -w` ist
  der billigste Test darauf, und er stand von Anfang an zur Verfügung. Dieser
  Slice hat die Regel formuliert und **halb ausgeführt** (§3.1); der Review hat
  es gefangen. *(Kein `liegt in`-Feld: mit diesem Slice
  wurde nichts verkörpert — der Eintrag ist gezählt, nicht verkörpert. Der Ort
  wäre die Delta-Analyse-Form, und die hat a-check nicht als Ziel-Form, sondern
  als Präzedenz in drei Slices.)*

- **Beobachtungs-Register (`../observations/`):**
  [`BEO-GATE/versions-sensor-trifft-planungs-vorgriff`](../observations/BEO-GATE/versions-sensor-trifft-planungs-vorgriff/observation.md)
  **neu angelegt**, Beleg `evidence/slice-174.md`, Stand `offen (1×)` — der
  Sensor aus slice-173 traf zweimal ein Planungsdokument, das den *Ziel*-Stand
  nennt. Zwei Funde, **ein** Vorgang, also ein Beleg.
  Dazu `evidence/slice-174.md` in
  [`BEO-PLAN/review-geltungsbereich-zu-eng`](../observations/BEO-PLAN/review-geltungsbereich-zu-eng/observation.md)
  — **Zähler 3×, Schwelle erreicht**: die Bereinigung dieses Slice sah eine
  Formatierungs-Klasse und suchte die zweite nicht (§3.1).

- **Folge-Slices:** [slice-175](../done/slice-175-etappe-a-vendoring-v650.md)
  (A, Vendoring),
  [slice-176](../done/slice-176-zitier-form-einfrierende-artefakte.md) (C,
  Zitier-Form),
  [slice-177](../done/slice-177-sensors-struktur-zwei-tabellen.md) (D,
  Sensors-Struktur),
  [slice-178](../in-progress/slice-178-slice-form-ziel-und-abgrenzung.md) (E,
  Slice-Form) — alle vier sind Dateien in `open/`. Der vierte entstand erst
  durch den Review (F-3/F-4): T-3 hatte keine Etappe, und die Schritt-Hälfte
  der Regel lag fälschlich unter §3.3. Etappe **B** (Adaptions-Durchgang)
  bekommt ihren Slice, wenn A den Stand gehoben hat: sie misst gegen ihn und
  wäre vorher gegenstandslos.

- **Risiken aus §7:** drei, jedes mit genau einem Ausgang — eingetreten
  (aufgefangen) · weiter offen → Register · gestrichen mit Begründung. Siehe §7.

- **Drei Paarungen:** trägt die Welle-Closure, nicht dieser Slice — er hat ein
  `**Welle:**`-Feld ([`AGENTS.md`](../../../../AGENTS.md) §6, Baseline
  `modul-08` §Rollen-Sequenz für eine Welle).

## 9. Sub-Area-Modus

**Vorgelagert — Sub-Area-Wahl prüfen:** eine Sub-Area berührt —
**Vendored Baseline** (`.harness/baseline/`) führt *keinen* Modus (externer
Fremdtext), und dieser Slice ändert dort nichts. Gemessen und bewertet wird
über sie hinweg; die Sub-Areas, die **Folge**-Slices berühren werden
(`HARNESS`, `GATE`, `PLAN`), gehören in deren §9, nicht hierher.

**Vorgelagert — offene Beobachtungen sichten:** Register am 2026-09-07
durchgegangen, **34** offene Einträge in `BEO-HARNESS`/`BEO-GATE`/`BEO-PLAN`,
keiner an der Schwelle (die Zahl stand zuerst auf 31 — geschätzt statt gezählt,
Review-Befund F-6). Vier einschlägig:

| Eintrag | Stand | Bezug |
|---|---|---|
| [`zwei-baseline-staende-nach-migrationsende`](../observations/BEO-HARNESS/zwei-baseline-staende-nach-migrationsende/observation.md) | offen, 1× | genau dieses Fenster — während einer Migration liegen zwei Stände zulässig nebeneinander; die Welle muss ihn beim Abschluss wieder auf einen bringen |
| [`adaption-korrigiert-repo-aussage`](../observations/BEO-HARNESS/adaption-korrigiert-repo-aussage/observation.md) | offen, 2× | der Adaptions-Durchgang wird `MR`-Einträge anfassen; entsteht dabei ein weiterer Repo-Aussage-Korrektur-Eintrag, ist die Schwelle erreicht |
| [`form-vergleich-sprachblind`](../observations/BEO-PLAN/form-vergleich-sprachblind/observation.md) | offen, 1× | die Analyse ist ein Form-Vergleich gegen Ziel-Formen — genau die Tätigkeit, bei der der Eintrag entstand |
| [`slice-form-vier-begriffe-fehlen`](../observations/BEO-PLAN/slice-form-vier-begriffe-fehlen/observation.md) | offen, 1× | `slice.template.md` ändert sich (+46); die a-check-Kopieranleitung in `AGENTS.md` §5 ist daran zu messen |

**Alle berührten Sub-Areas GF** — kein Begründungsblock nötig.
