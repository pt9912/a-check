# slice-171 — Zwei falsche Aussagen im Konventions-Bestand korrigieren

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `make slice-mv` ([`AGENTS.md`](../../../../AGENTS.md)
§3.3/§5).

**Welle:** ohne Welle (der Closure-Trigger wäre die eigene DoD — kein
repo-weites Mehr).

**Bezug:** Ergebnis des Durchgangs durch alle aktiven Adaptionen am
2026-09-06 (Maintainer-geführt, Eintrag für Eintrag). Der Durchgang selbst
hat [`MR-018`](../../../../harness/conventions.md#mr-018) aufgelöst
([slice-170](../done/wellenlos/slice-170-mr018-aufloesen.md)); die beiden
hier notierten Funde blieben liegen.

**Berührte Spec-Stellen:** — *(keine)* — Harness-Konventionen ohne
Vertragsberührung.

**Verantwortlich:** Claude — gesetzt beim Übergang nach `in-progress/`.

**Autor:** Claude (Opus 5), im Auftrag des Maintainers. **Datum:**
2026-09-06.

---

## 1. Ziel und Abgrenzung

**Ziel:** 
Zwei Aussagen im Konventions-Bestand stimmen nicht mehr mit dem Repo
überein. Beide lösen auf, beide sind grün, keine wird von einem Gate
gesehen — sie fallen nur auf, wenn jemand sie liest.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Die Suffix-Form einführen.** Der Eintrag korrigiert seine **Begründung**,
  nicht seine Entscheidung; ein Umbau träfe alle sieben `SPEC-*` samt ihrer
  Traceability-Verweise, ohne eine Aussage zu ändern. Das ist der Inhalt des
  neuen Eintrags, nicht sein Gegenteil.
- **Die zwei Funde aus §2.3** — sie brauchen je eine eigene Entscheidung und
  stehen dort notiert, damit sie nicht verloren gehen. Ein Slice, der sie
  mitnimmt, hätte vier Liefer-Punkte statt zwei.
- **Ein Sensor auf überholte Aussagen.** *„Stimmt diese Begründung noch?"* ist
  ein Urteil, kein Match ([`AGENTS.md`](../../../../AGENTS.md) §3.7) — genau
  deshalb fiel beides erst beim Lesen auf.

**Zwei benannte Plan-Änderungen** ([`AGENTS.md`](../../../../AGENTS.md) §6):

1. **Ein dritter `MR`-Eintrag.** Der Review fand in
   [`MR-021`](../../../../harness/conventions/done/MR-021-verfeinerungs-form.md)
   eine falsche Zahl (F-1) — denselben Fehlertyp, den der Eintrag behob. Er war
   zu diesem Zeitpunkt bereits gepusst und damit im Bestand; ein akzeptierter
   Eintrag wird nicht überschrieben. Daher die Kette
   [`MR-011`](../../../../harness/conventions/done/MR-011-verfeinerungs-form.md) → [`MR-021`](../../../../harness/conventions/done/MR-021-verfeinerungs-form.md) → [`MR-022`](../../../../harness/conventions/MR-022-verfeinerungs-form.md)
   statt einer stillen Korrektur.
2. **Zwei weitere falsche Aussagen im selben Abschnitt** (Review F-6): Der
   Prosa-Rahmen über der Adaptions-Tabelle behauptete *„jede Zeile trägt zwei
   Anker"* (gemessen: null der sieben) und *„die Spalte steht hier überall auf
   `—`"* (gemessen: fünf von sieben tragen einen Zeiger). **Damit ist die
   Rückführungs-Bedingung aus §5 wörtlich eingetreten** — *„weitere überholte
   Aussagen im selben Abschnitt"*. Sie wird trotzdem nicht gezogen: Es sind zwei
   Prosa-Zeilen in einer Datei, die der Slice ohnehin ändert, und nichts daran
   wächst mit dem Umfang (Größen-Regel). Die Entscheidung steht hier, damit sie
   prüfbar ist statt stillschweigend.

## 2. Analyse (vor der Umsetzung)

### 2.1 [`MR-011`](../../../../harness/conventions.md#mr-011) begründet sich mit einem Feld, das die genannte Beziehung nicht trägt

Der Eintrag rechtfertigt, warum a-check die Suffix-Form der Baseline nicht
übernimmt, unter anderem so: *„die Verfeinerungs-Beziehung steht in a-check
ohnehin explizit im `Schärft:`-Feld und nicht in der Kennung."*

**Gemessen (2026-09-06):**

| Frage | Ergebnis |
|---|---|
| Existiert `**Schärft:**` als Feld? | ja — **41** Vorkommen in [`docs/plan/adr/`](../../adr/README.md), eines in [`spec/lastenheft.md`](../../../../spec/lastenheft.md) |
| Welche Richtung trägt es dort? | ADR → `SPEC-*` (bzw. im Lastenheft `AC-*` → `AC-*`) |
| Trägt ein `SPEC-*`-Eintrag es? | **nein** — 7 `SPEC-*`-Überschriften in [`spec/spezifikation.md`](../../../../spec/spezifikation.md), **0** `**Schärft:**`-Felder |
| Wie steht die Verfeinerung dort wirklich? | als Prosasatz „Präzisiert [`AC-…`]" |

Die **Entscheidung** bleibt unberührt: der Maintainer hat am 2026-09-06
entschieden, **nicht** auf die Suffix-Form zu migrieren — ein Umbau träfe
alle bestehenden `SPEC-*` samt ihrer Traceability-Verweise, ohne eine
Aussage zu ändern. Falsch ist nur die **Begründung**, und die trägt der
Eintrag als Beleg.

**Der tragfähige Grund steht daneben und ist stärker:** die `SPEC-*`-Anker
werden aus `Accepted`-ADRs referenziert, und die sind immutabel
([`AGENTS.md`](../../../../AGENTS.md) §3.5). Eine Umbenennung erzeugte
denselben unauflösbaren Widerspruch zwischen `make doc-check` und
`make doc-immutable`, den der
[Anforderungs-Anlege-Prozess](../../../../harness/conventions.md#anforderungs-anlege-prozess)
für `AC-*`-Überschriften bereits beschreibt.

**Weg:** ein akzeptierter Eintrag wird nicht überschrieben
([`conventions.md`](../../../../harness/conventions.md) §Adaptions-Block).
Also ein **neuer** `MR`-Eintrag, der [`MR-011`](../../../../harness/conventions.md#mr-011) ablöst und die Begründung auf
den gemessenen Stand stellt; [`MR-011`](../../../../harness/conventions.md#mr-011) wandert nach `conventions/done/`.
Präzedenz: [`MR-017`](../../../../harness/conventions.md#mr-017) →
[`MR-020`](../../../../harness/conventions.md#mr-020)
([slice-166](../done/slice-166-mr020-adr-vorlage-generisch.md)). Die
Kennung des neuen Eintrags wird beim Schreiben vergeben, nicht hier — IDs
werden referenziert, nicht erfunden
([`AGENTS.md`](../../../../AGENTS.md) §5).

### 2.2 [`harness/README.md`](../../../../harness/README.md) deklariert eine aufgelöste Adaption

Die Rollen-Tabelle führt beide Validator-Kanten als *„unverkörpert,
deklariert als [`MR-009`](../../../../harness/conventions.md#mr-009)"*. [`MR-009`](../../../../harness/conventions.md#mr-009)
ist seit [slice-163](../done/slice-163-adaptions-durchgang-v610.md)
**aufgelöst** und liegt in `conventions/done/`; aktiv ist
[`MR-016`](../../../../harness/conventions.md#mr-016).

Kein Gate fängt das: der Link löst auf, weil die Doppel-Anker-Disziplin den
alten Slug am Leben hält. Genau dafür ist sie da — sie schützt den Verweis,
nicht seine Richtigkeit.

**Beides ist zwischen Anlage und Umsetzung dieses Slice erledigt worden** —
gemessen am 2026-09-08, vor dem Übergang nach `in-progress/`:
[`harness/README.md`](../../../../harness/README.md) nennt [`MR-009`](../../../../harness/conventions/done/MR-009-validator-unbesetzt.md) **null**
Mal und den Kontext-Trennungs-Absatz ebenso. Beides fiel bei
[slice-182](../done/wellenlos/slice-182-readme-verweist-statt-wiederholt.md)
an: Die Rollen-Tabelle wurde auf
[`MR-016`](../../../../harness/conventions/MR-016-validator-unbesetzt.md)
umgestellt (dessen Review-Finding F-10), der Chronik-Absatz gestrichen (§3.7).
**Der DoD-Punkt bleibt stehen und ist abgehakt** — er beschreibt einen Zustand,
nicht eine Tätigkeit; wer ihn streicht, verliert die Aussage, dass er geprüft
wurde.

**Was §2.2 vor dieser Erledigung notierte:** der Absatz *Kontext-Trennung, real
angewandt* beschreibt die Review-Serie vom 2026-07-26 als Selbst-Review und
schließt mit *„ein unabhängiger Lauf bleibt eine eigene Übergabe"*. Seit
`welle-14` ist der unabhängige Lauf die Regel (slice-161…slice-170, je ein
Report aus getrenntem Kontext). Der Satz beschreibt seine Serie korrekt,
liest sich aber als Gegenwart — zu entscheiden ist, ob er bleibt,
umformuliert wird oder um den neuen Stand ergänzt.

### 2.3 Was der Durchgang **nicht** erfasst hat — hier notiert, damit es nicht verloren geht

Keiner der beiden Punkte gehört in die DoD dieses Slice; beide brauchen eine
eigene Entscheidung, und beide standen bis hierher nur im Gesprächsprotokoll.

- **Der Durchgang deckte sechs von acht aktiven Einträgen.** Geprüft wurden
  [`MR-011`](../../../../harness/conventions.md#mr-011),
  [`MR-012`](../../../../harness/conventions.md#mr-012),
  [`MR-014`](../../../../harness/conventions.md#mr-014),
  [`MR-015`](../../../../harness/conventions.md#mr-015),
  [`MR-016`](../../../../harness/conventions.md#mr-016) und
  [`MR-018`](../../../../harness/conventions.md#mr-018).
  **Nicht geprüft:** [`MR-019`](../../../../harness/conventions.md#mr-019) und
  [`MR-020`](../../../../harness/conventions.md#mr-020) — beide entstanden
  erst während `welle-14` und wirkten darum frisch.
  [`MR-019`](../../../../harness/conventions.md#mr-019) nennt in seinem
  Text [`MR-017`](../../../../harness/conventions.md#mr-017) und
  [`MR-018`](../../../../harness/conventions.md#mr-018) als
  Rückbau-Kandidaten; **beide sind inzwischen aufgelöst**, der Eintrag
  beschreibt insoweit einen Stand, den es nicht mehr gibt.
- **Ein Beobachtungs-Eintrag trägt einen nicht deklarierten Sub-Area-Namen.**
  [`BEO-HARNESS/sensor-ohne-dod-phrase-wirkungslos`](../observations/BEO-HARNESS/sensor-ohne-dod-phrase-wirkungslos/observation.md)
  führt „Sub-Area: **Harness-Tooling**" — gemessen der einzige von 48
  Einträgen mit einem Namen, den die
  [Modus-Deklaration](../../../../harness/conventions.md#modus-deklaration-pro-sub-area)
  nicht kennt. `modul-06` sagt dazu: dann ist *„entweder die Zuordnung falsch
  oder die Deklaration unvollständig"* — welches von beidem, ist die
  Entscheidung. Der Eintrag selbst ist ab Anlage unveränderlich, eine
  Korrektur an Ort und Stelle also nicht der Weg.

## 3. Umsetzung

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| [`MR-021`](../../../../harness/conventions/done/MR-021-verfeinerungs-form.md) | neu | löst [`MR-011`](../../../../harness/conventions/done/MR-011-verfeinerungs-form.md) ab, Begründung auf den gemessenen Stand |
| [`MR-011`](../../../../harness/conventions/done/MR-011-verfeinerungs-form.md) nach `conventions/done/` | `git mv` | akzeptierter Eintrag wird nicht überschrieben |
| `harness/conventions.md`, beide Tabellen | update | Zeile wandert von *Aktive* nach *Aufgelöste*, beide Anker bleiben |
| `harness/README.md` §Rollen | — | **bereits erledigt** mit slice-182 |

**Zwei Commits statt einem** ([`AGENTS.md`](../../../../AGENTS.md) §3.3): erst
die Bewegung samt der Pfade, die sie erzwingt — die Datei liegt eine Ebene
tiefer, also brauchen ihre **fünf** relativen Links ein `../` mehr —, dann der
ablösende Eintrag. Rename-Erkennung nach dem ersten Commit: **72 %**, also weit
über der 50-%-Schwelle, `git log --follow` bleibt zuverlässig.

**Die Messung, auf der [`MR-021`](../../../../harness/conventions/done/MR-021-verfeinerungs-form.md) steht** (2026-09-08):

| Frage | Ergebnis |
|---|---|
| `SPEC-*`-Abschnitte in `spec/spezifikation.md` | **7** |
| davon mit `**Schärft:**`-Feld | **0** |
| Wie steht die Verfeinerung dort? | Prosasatz *„Präzisiert [`AC-…`]"* |
| ADR-Dateien mit `Schärft:`-Feld | **40** |
| Vorkommen des Worts in `spezifikation.md` | **1** — im Kopf der Historie-Sektion, und dort steht, dass **die ADR** es aufwärts deklariert |

Die letzte Zeile ist die Gegenprobe: Sie sieht zunächst aus wie ein Gegenbeleg
und ist bei genauem Lesen der stärkste Beleg.

## 4. Definition of Done

- [x] [`MR-011`](../../../../harness/conventions.md#mr-011) durch
      [`MR-021`](../../../../harness/conventions/done/MR-021-verfeinerungs-form.md)
      abgelöst, dessen Begründung den gemessenen Stand trägt;
      [`MR-011`](../../../../harness/conventions.md#mr-011) in
      `conventions/done/`, beide Tabellen nachgezogen, Anker `mr-011` erhalten.
- [x] `harness/README.md` §Rollen nennt
      [`MR-016`](../../../../harness/conventions.md#mr-016) statt [`MR-009`](../../../../harness/conventions/done/MR-009-validator-unbesetzt.md); der
      Absatz zur Kontext-Trennung ist entschieden — **beides mit slice-182
      erledigt**, hier nur noch gemessen (§2.2).
- [ ] Unabhängiger Review durchgeführt (Report unter `docs/reviews/`).
- [ ] Jedes Risiko trägt einen Ausgang.

`make gates` und `make verify` grün.

## 5. Trigger

**Start** (`open` → `in-progress`): Maintainer-Freigabe und WIP-Limit frei.

**Rückführungen:** wächst der Umfang über die zwei Korrekturen hinaus —
etwa weil der Durchgang weitere überholte Aussagen im selben Abschnitt
findet — zurück nach `next/` zur Zerlegung.

## 6. Closure-Trigger

DoD vollständig, `make gates`/`make verify` grün, Closure-Notiz
geschrieben.

## 7. Risiken und offene Punkte

- *Der neue Eintrag wiederholt den Fehler von [`MR-011`](../../../../harness/conventions/done/MR-011-verfeinerungs-form.md) in anderer Form —
  eine Begründung, die plausibel klingt und nicht gemessen ist* — **Ausgang:**
  *entfallen*, gestrichen mit Begründung. Beide tragenden Gründe in
  [`MR-021`](../../../../harness/conventions/done/MR-021-verfeinerungs-form.md)
  nennen ihre Zahl und ihre Quelle (7 `SPEC-*`-Abschnitte, 0 mit Feld, 40
  ADR-Dateien mit Feld), und die Gegenprobe — das eine Vorkommen des Worts in
  der Spezifikation — steht ausdrücklich dabei. Das Risiko war der Grund, die
  Messung in den Eintrag selbst zu schreiben statt nur in den Slice.
- *Die Ablösung von [`MR-011`](../../../../harness/conventions/done/MR-011-verfeinerungs-form.md) entfernt einen der fünf `v6.0.0`-Anker im Feld
  `Ersetzt-Baseline-Regel`; vier bleiben, und die Querschnitts-Entscheidung
  über sie ist damit weiterhin offen* — **Ausgang:** *weiter offen* →
  Beobachtungs-Register,
  [`BEO-HARNESS/zwei-baseline-staende-nach-migrationsende`](../observations/BEO-HARNESS/zwei-baseline-staende-nach-migrationsende/observation.md).
  **Gemessen und dabei korrigiert:** Der Slice-Plan sprach von `v6.0.0`-Ankern;
  im heutigen Bestand zeigt kein `Ersetzt-Baseline-Regel`-Feld mehr dorthin —
  slice-172 hat den Stand entfernt, und die Zeiger sind mit der
  Baseline-Migration auf `v6.5.0` mitgewandert. Der neue Eintrag trägt denselben
  Zeiger, nur auf den aktuellen Stand.
  **Zum Ausgang selbst** (Review F-4): Der genannte Eintrag steht auf
  *verkörpert*, und dieser Slice legt **keinen** Beleg an — der Zähler bewegt
  sich also nicht. Das ist beabsichtigt und der Grund steht in §9: Der Slice
  berührt ein Feld, das der Eintrag beschreibt, aber er **löst die Klasse nicht
  aus** — alle fünf Zeiger stehen bereits auf dem aktuellen Stand. *Weiter
  offen* heißt hier: Die Frage kehrt beim nächsten Baseline-Sprung wieder, nicht
  dass dieser Slice sie erneut gestellt hätte.

## 8. Closure-Notiz

**Lerneintrag — Form: geschärfte Regel** (die Messung wandert in den
`MR`-Eintrag selbst, nicht nur in den Slice).

- **Was hat funktioniert:** Die Begründung **im Eintrag** zu messen statt im
  Slice-Plan. [`MR-011`](../../../../harness/conventions/done/MR-011-verfeinerungs-form.md) fiel auf, weil ein Leser seinen dritten Beleg nachprüfte
  — der Eintrag selbst nannte keine Zahl. [`MR-021`](../../../../harness/conventions/done/MR-021-verfeinerungs-form.md) nennt vier: sieben
  `SPEC-*`-Abschnitte, null mit Feld, 40 ADR-Dateien mit Feld, ein Vorkommen
  des Worts in der Spezifikation. Wer ihn in einem Jahr prüft, braucht den
  Slice-Plan nicht.

- **Was ging anders als geplant — der halbe Slice war beim Beginn schon
  erledigt.** DoD-Punkt 2 (`harness/README.md` nennt [`MR-009`](../../../../harness/conventions/done/MR-009-validator-unbesetzt.md), Kontext-Absatz)
  fiel bei slice-182 an, dessen Review ihn als F-10 fand. Gemessen vor dem
  Übergang: null Vorkommen beider Stellen. **Der DoD-Punkt bleibt stehen und
  ist abgehakt** — er beschreibt einen Zustand, nicht eine Tätigkeit; wer ihn
  streicht, verliert die Aussage, dass er geprüft wurde.

- **Was der Review fand — und warum es diesmal besonders weh tut.** 2 HIGH,
  4 MEDIUM, 3 LOW. **F-1: [`MR-021`](../../../../harness/conventions/done/MR-021-verfeinerungs-form.md) nannte eine falsche Zahl** — *„40
  ADR-Dateien tragen ein `Schärft:`-Feld, das auf solche Anker zeigt"*. Die 40
  zählt Dateien mit dem *String*, den Index eingeschlossen; über Felder, die
  zeigen, sagt sie nichts (gemessen: 34, davon 30 `Accepted`). **Das ist
  derselbe Fehlertyp, den derselbe Eintrag behob.** Ein Beleg, der plausibel
  klingt und die gestellte Frage nicht misst — zweimal hintereinander, in
  derselben Sache.
  **F-2: *„keine weiteren Treffer für `HARNESS`"* war falsch** — 14 Einträge,
  10 offen, zwei genannt. Vier Wörter statt vierzehn Zeilen.

- **Steering-Loop-Eintrag — geschärfte Regel:** *Eine Begründung, die einen
  Beleg über das Repo führt, nennt ihn mit Zahl und Quelle — im Eintrag, nicht
  im Slice.* Ein `MR`-Eintrag ist immutabel und wird Jahre später gelesen; ein
  Beleg wie *„steht ohnehin im `Schärft:`-Feld"* ist dann nicht mehr prüfbar,
  ohne die Messung neu zu erfinden. — liegt in
  [`MR-022`](../../../../harness/conventions/MR-022-verfeinerungs-form.md)
  §Begründung, *Zählregel und Geltungsbereich* (`seit slice-171` dort).
  **Kein Sensor:** Ob eine Begründung gemessen ist, ist ein Urteil über ihren
  Entstehungsweg ([`AGENTS.md`](../../../../AGENTS.md) §3.7) — dieselbe Grenze,
  die §5 für CR-Texte an ein fremdes Werkzeug bereits benennt.

- **Die zweite Lehre, teurer bezahlt:** *Ein `MR`-Eintrag ist ab seinem
  Anlage-Commit im Bestand — der Review kommt danach.* [`MR-021`](../../../../harness/conventions/done/MR-021-verfeinerungs-form.md) war beim
  Finding bereits gepusht; die Immutabilität ließ nur die Ablösung. Daraus die
  Kette [`MR-011`](../../../../harness/conventions/done/MR-011-verfeinerungs-form.md) → [`MR-021`](../../../../harness/conventions/done/MR-021-verfeinerungs-form.md) → [`MR-022`](../../../../harness/conventions/MR-022-verfeinerungs-form.md) für **eine Zahl**. Wer das vermeiden
  will, hat zwei Wege: den Review **vor** den Anlage-Commit ziehen, oder den
  Eintrag zunächst als `Proposed` führen und erst mit der Closure auf
  `Accepted` setzen. **Der Bestand kennt keinen `Proposed`-`MR`** — gemessen:
  alle **12** tragen `Accepted`. Das ist die benannte Spec-Lücke dieses Slice,
  kein Vorschlag: Ob der Weg beschritten wird, entscheidet ein eigener Vorgang.

- **Beobachtungs-Register (`../observations/`):** kein neuer Eintrag, kein
  neuer Beleg — beide gesichteten Einträge sind in §9 begründet ausgeschlossen.
  Keine Beobachtung angefallen ist ebenfalls eine Antwort.

- **Folge-Slices:** keiner neu. Die zwei Funde aus §2.3 bleiben dort notiert;
  sie brauchen je eine eigene Entscheidung und sind in §1 ausgeschlossen.

- **Risiken aus §7:** zwei, jedes mit genau einem Ausgang — einmal *entfallen*
  mit Begründung, einmal *weiter offen* → Register.

- **Drei Paarungen** (Repo ohne Wellen-Betrieb) — geprüft **nach** dem `git mv`
  nach `done/`, weil sie dort suchen; alle drei tragen:
  **Anker** — `liegt in` [`MR-022`](../../../../harness/conventions/MR-022-verfeinerungs-form.md) §Begründung; der Zielort führt `seit slice-171`.
  *(Er fehlte beim ersten Prüfen und zeigte davor auf den inzwischen abgelösten
  [`MR-021`](../../../../harness/conventions/done/MR-021-verfeinerungs-form.md) — nachgetragen, solange der Eintrag noch unveröffentlicht war.)*
  **Folge-Slice** — keiner genannt.
  **Register** — kein Eintrag zitiert, der einen Beleg bekäme: §9 geht alle
  **14** `HARNESS`-Einträge durch und begründet je Zeile, warum keiner ausgelöst
  ist. Die Paarung prüft Deckung, und ungedeckt ist nichts.

## 9. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** eine Sub-Area berührt —
**Harness-Einstieg** `HARNESS` (`harness/conventions.md`, `harness/conventions/`),
Achsen 1,2,3 laut
[Modus-Deklaration](../../../../harness/conventions.md#modus-deklaration-pro-sub-area).
Die Spec-Straten sind **nicht** berührt: Der Slice **misst**
[`spec/spezifikation.md`](../../../../spec/spezifikation.md), ändert dort aber
keine Zeile — eine Lese-Berührung ist keine.

**Vorgelagert — offene Beobachtungen sichten:** gesichtet am 2026-09-08 gegen
den gemergten Stand. Zwei Treffer, beide vorab richtig vermutet:

- [`BEO-HARNESS/adaption-korrigiert-repo-aussage`](../observations/BEO-HARNESS/adaption-korrigiert-repo-aussage/observation.md)
  — **2×**. Dieser Slice ist ein drittes Vorkommen der Klasse, aber **kein
  neuer Beleg**: [`MR-021`](../../../../harness/conventions/done/MR-021-verfeinerungs-form.md) korrigiert die **eigene Begründung** eines
  Konventions-Eintrags, nicht eine Repo-Aussage außerhalb davon. Der Eintrag
  beschreibt den Fall, dass eine *Adaption* eine Aussage über das Repo
  korrigiert; hier korrigiert sie sich selbst. Die Unterscheidung ist knapp —
  sie steht hier, damit ein späterer Leser sie nachrechnen kann statt sie zu
  raten.
- [`BEO-HARNESS/zwei-baseline-staende-nach-migrationsende`](../observations/BEO-HARNESS/zwei-baseline-staende-nach-migrationsende/observation.md)
  — **1×**, *verkörpert*. Der Slice berührt genau ein
  `Ersetzt-Baseline-Regel`-Feld; gemessen tragen alle **fünf** aktiven Felder
  heute `v6.5.0`, keines mehr einen alten Stand. Der Eintrag bleibt, weil die
  Querschnitts-Frage beim **nächsten** Sprung wiederkehrt.

**Die übrigen zwölf `HARNESS`-Einträge, einzeln durchgegangen** — der Review
fand die erste Fassung *„keine weiteren Treffer"* gegen das Register falsch
(F-2), und sie war es: `BEO-HARNESS/` führt **14** Einträge, **10** davon
`offen`.

| Eintrag | Stand | berührt dieser Slice ihn? |
|---|---|---|
| [`hard-rule-37-ohne-sensor`](../observations/BEO-HARNESS/hard-rule-37-ohne-sensor/observation.md) | offen 2× | **nein, und das ist die knappste Entscheidung.** Der Slice *beruft* sich zweimal auf §3.7 — das ist Anwendung der Regel, kein Vorkommen der Lücke. Ein Beleg entstünde bei einem **Verstoß**, der ungeprüft bliebe (so bei slice-180: zwei Konjunktiv-Kommentare). Hier ist keiner gefunden. |
| [`mr-aufloesungs-trigger-ohne-waechter`](../observations/BEO-HARNESS/mr-aufloesungs-trigger-ohne-waechter/observation.md) | offen 1× | **nein.** Der Slice legt einen `MR` mit Auflösungs-Trigger an — der Trigger ist unverändert von [`MR-011`](../../../../harness/conventions/done/MR-011-verfeinerungs-form.md) übernommen und bekommt keinen neuen Wächter, aber er **verliert** auch keinen. Der Eintrag beschreibt eine stehende Lücke, kein Ereignis. |
| [`chronik-in-gelesenen-dateien`](../observations/BEO-HARNESS/chronik-in-gelesenen-dateien/observation.md) | offen 2× | **nein.** Der Eintrag nennt `conventions.md` namentlich; der Slice hat dort zwei falsche **Tatsachen**-Aussagen korrigiert (F-6), keine Chronik entfernt. Andere Klasse. |
| [`baseline-normtext-nachgeschrieben`](../observations/BEO-HARNESS/baseline-normtext-nachgeschrieben/observation.md) | offen 2× | nein — kein Baseline-Normtext berührt. |
| [`adaption-korrigiert-repo-aussage`](../observations/BEO-HARNESS/adaption-korrigiert-repo-aussage/observation.md) | offen 2× | nein, siehe oben. |
| [`baseline-regel-nie-erwogen-weil-bestand-sie-verletzt`](../observations/BEO-HARNESS/baseline-regel-nie-erwogen-weil-bestand-sie-verletzt/observation.md) | offen 2× | nein — die Baseline-Regel (Suffix-Form) ist ausdrücklich erwogen und begründet abgelehnt; das ist das Gegenteil des Eintrags. |
| [`agents-md-hinkt-baseline-dod-item-hinterher`](../observations/BEO-HARNESS/agents-md-hinkt-baseline-dod-item-hinterher/observation.md) | offen 1× | nein — `AGENTS.md` nicht berührt. |
| [`rueckbau-eintrag-ablage-auslegung`](../observations/BEO-HARNESS/rueckbau-eintrag-ablage-auslegung/observation.md) | offen 1× | nein — kein Rückbau-Eintrag; [`MR-022`](../../../../harness/conventions/MR-022-verfeinerungs-form.md) ist eine Ablösung mit fortbestehender Adaption. |
| [`selbst-archivierung-verdoppelt-abschluss-aufwand`](../observations/BEO-HARNESS/selbst-archivierung-verdoppelt-abschluss-aufwand/observation.md) | offen 1× | nein — betrifft die Closure-Mechanik, nicht diesen Gegenstand. |
| [`sensor-ohne-dod-phrase-wirkungslos`](../observations/BEO-HARNESS/sensor-ohne-dod-phrase-wirkungslos/observation.md) | offen 1× | nein — keine DoD-Phrase geändert. |
| [`behauptete-vollstaendigkeit-extern-gefangen`](../observations/BEO-HARNESS/behauptete-vollstaendigkeit-extern-gefangen/observation.md) | verkörpert 4× | **berührt, ohne Beleg.** Die erste Fassung dieser Sichtung behauptete Vollständigkeit (*„keine weiteren Treffer"*) und war unvollständig — extern gefangen, vom Review. Die Regel dagegen ist seit slice-159 in `AGENTS.md` §6 verkörpert und hat **funktioniert**: Der unabhängige Review fand es. Ein Beleg zählt ein *Versagen* der Deckung; hier hat sie getragen. |
| [`harness-einstieg-ohne-modus-zeile`](../observations/BEO-HARNESS/harness-einstieg-ohne-modus-zeile/observation.md) · [`rueckbau-kandidat-ueberlebt-baseline-migration`](../observations/BEO-HARNESS/rueckbau-kandidat-ueberlebt-baseline-migration/observation.md) · [`zwei-baseline-staende-nach-migrationsende`](../observations/BEO-HARNESS/zwei-baseline-staende-nach-migrationsende/observation.md) | verkörpert | nein bzw. oben behandelt. |

**Die Lehre aus F-2 steckt in der Tabelle selbst:** *„keine weiteren Treffer"*
ist eine Vollständigkeits-Behauptung und braucht dieselbe Sorgfalt wie eine
Messung. Vier Wörter sind billiger als vierzehn Zeilen — und genau deshalb
falsch.

**Alle berührten Sub-Areas GF** — kein Begründungsblock nötig.
