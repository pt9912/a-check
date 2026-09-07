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

### 3.1 Die Größe täuscht: 8 von 30 Dateien tragen die Substanz

Der Roh-Diff nennt 30 Dateien und `+572`. Davon sind **36 hinzugefügte Zeilen
reine Markdown-Tabellen-Trennzeilen** (`| --- | --- |`) — eine Formatierungs-
Korrektur im Kurs, die keine Regel ändert. Gemessen je Datei
(`git diff … | grep -cE '^\+\| *-{2,}[ |-]*\|? *$'`):

| Klasse | Dateien | substanzielle `+`-Zeilen |
|---|---|---|
| **Substanz** (≥ 18) | **8** | 436 von 536 |
| Wenig (1–11) | 13 | 100 |
| **Null** — nur Trennzeilen | **9** | 0 |

Die neun ohne jede Substanz: `modul-07`, `modul-08`, `modul-10`, `modul-11`,
`modul-12`, `modul-14`, `modul-15`, `grundlagen-klassifikation`,
`grundlagen-durchsetzungsschicht`. **Sie berühren a-check nicht** — nicht weil
sie unwichtig wären, sondern weil sich ihr Normtext nicht geändert hat.

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

**Das ist genau die Klemme aus slice-172/173** — dort gefunden, hier an der
Wurzel gelöst. a-check hat sie mit `exempt-paths` behandelt: fünf
Verzeichnis-Klassen, die der Versions-Sensor überspringt. Der Kurs geht anders
vor: In einem einfrierenden Artefakt entsteht der Link gar nicht erst, dann
braucht es auch keine Ausnahme. **Die beiden Antworten schließen sich nicht
aus** — `exempt-paths` deckt den Bestand, die Zitier-Form den Zuwachs —, aber
sie überschneiden sich, und welche a-check führt, ist eine Entscheidung, keine
Ableitung.

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

**Folge-Slice:** Kopieranleitung nachziehen; ob der Bestand nachgerüstet wird,
ist eine eigene Frage — die Ziel-Form gilt für **neue** Slices.

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

### 3.3 Was nicht berührt

- Die neun Null-Substanz-Dateien (§3.1).
- `grundlagen-begriffe.md` (+42/−40): Umbau der Begriffs-Tabelle; neu benannt
  sind *Auslesestand, kein Artefakt* · *Waisen* · *Pfad* ·
  *Vertrag › Technik › Sicht › ADR › Slice*. Sie schärfen T-5 und die
  Referenz-Richtung begrifflich, ändern aber keine Regel, an der a-check hängt.
- `modul-09` (+9), `modul-02` (+1), `modul-04` (+6), `modul-06` (+3),
  `grundlagen-source-precedence` (+3), `grundlagen-referenz-richtung` (+6),
  `grundlagen-bootstrap` (+18): Tabellen-Ausbau und Verweise auf die neuen
  Abschnitte. **Zu prüfen im Adaptions-Durchgang**, nicht als eigener Slice —
  `grundlagen-referenz-richtung` und `grundlagen-source-precedence` tragen die
  `Ersetzt-Baseline-Regel`-Anker von
  [`MR-011`](../../../../harness/conventions.md#mr-011) und
  [`MR-012`](../../../../harness/conventions.md#mr-012).

### 3.4 Vorgeschlagener Etappen-Schnitt

| Etappe | Inhalt | Umfang |
|---|---|---|
| **A** — Vendoring | `v6.5.0` vendoren, Stand an drei Stellen, vier Symlinks, `SHA256SUMS` | S |
| **B** — Adaptions-Durchgang | sieben aktive `MR` gegen `v6.5.0`, dabei §3.3-Restposten und T-5 | M |
| **C** — Zitier-Form (T-1) | vier Artefaktklassen; danach `exempt-paths` prüfen | M |
| **D** — Sensors-Struktur (T-2 + T-4) | zwei Tabellen, `harness/sensors/` für den Überhang | **L** |

**A vor allem anderen** — B, C und D messen gegen den vendorten Stand. C vor D
ist keine Pflicht, aber sinnvoll: die Zitier-Form betrifft Verweise auf
Sensor-Dateien, die D erst anlegt.

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
  eine Adaptions-Frage* — **Ausgang:** entfallen, gestrichen mit Begründung.
  Gemessen: die Vorlage nennt die Tabellenzeile ausdrücklich als **Default**
  und die Datei als Antwort auf den *Überhang*. a-checks Praxis wird damit
  nicht abgelöst, sondern begrenzt — 16 von 40 Zeilen liegen über der Grenze,
  24 darunter und bleiben, wie sie sind. Kein Widerspruch, also keine
  Adaption; Etappe D ist ein Nachzug.

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
  Umfangs-Beurteilung um seine formatierenden Anteile bereinigt; die Bereinigung
  ist zu messen, nicht zu schätzen. *(Kein `liegt in`-Feld: mit diesem Slice
  wurde nichts verkörpert — der Eintrag ist gezählt, nicht verkörpert. Der Ort
  wäre die Delta-Analyse-Form, und die hat a-check nicht als Ziel-Form, sondern
  als Präzedenz in drei Slices.)*

- **Beobachtungs-Register (`../observations/`):**
  [`BEO-GATE/versions-sensor-trifft-planungs-vorgriff`](../observations/BEO-GATE/versions-sensor-trifft-planungs-vorgriff/observation.md)
  **neu angelegt**, Beleg `evidence/slice-174.md`, Stand `offen (1×)` — der
  Sensor aus slice-173 traf zweimal ein Planungsdokument, das den *Ziel*-Stand
  nennt. Zwei Funde, **ein** Vorgang, also ein Beleg.

- **Folge-Slices:** [slice-175](../open/slice-175-etappe-a-vendoring-v650.md)
  (Etappe A, Vendoring),
  [slice-176](../open/slice-176-zitier-form-einfrierende-artefakte.md) (Etappe C,
  Zitier-Form),
  [slice-177](../open/slice-177-sensors-struktur-zwei-tabellen.md) (Etappe D,
  Sensors-Struktur) — alle drei sind Dateien in `open/`. Etappe B
  (Adaptions-Durchgang) bekommt ihren Slice, wenn Etappe A den Stand gehoben
  hat: sie misst gegen ihn und wäre vorher gegenstandslos.

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
durchgegangen, **31** offene Einträge in `BEO-HARNESS`/`BEO-GATE`/`BEO-PLAN`,
keiner an der Schwelle. Vier einschlägig:

| Eintrag | Stand | Bezug |
|---|---|---|
| [`zwei-baseline-staende-nach-migrationsende`](../observations/BEO-HARNESS/zwei-baseline-staende-nach-migrationsende/observation.md) | offen, 1× | genau dieses Fenster — während einer Migration liegen zwei Stände zulässig nebeneinander; die Welle muss ihn beim Abschluss wieder auf einen bringen |
| [`adaption-korrigiert-repo-aussage`](../observations/BEO-HARNESS/adaption-korrigiert-repo-aussage/observation.md) | offen, 2× | der Adaptions-Durchgang wird `MR`-Einträge anfassen; entsteht dabei ein weiterer Repo-Aussage-Korrektur-Eintrag, ist die Schwelle erreicht |
| [`form-vergleich-sprachblind`](../observations/BEO-PLAN/form-vergleich-sprachblind/observation.md) | offen, 1× | die Analyse ist ein Form-Vergleich gegen Ziel-Formen — genau die Tätigkeit, bei der der Eintrag entstand |
| [`slice-form-vier-begriffe-fehlen`](../observations/BEO-PLAN/slice-form-vier-begriffe-fehlen/observation.md) | offen, 1× | `slice.template.md` ändert sich (+46); die a-check-Kopieranleitung in `AGENTS.md` §5 ist daran zu messen |

**Alle berührten Sub-Areas GF** — kein Begründungsblock nötig.
