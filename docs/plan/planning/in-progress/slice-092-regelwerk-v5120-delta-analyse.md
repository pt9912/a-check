# slice-092 — Regelwerk-Migration `v3.5.2` → `v5.12.0`: Delta-Analyse

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Deckt:** keine `AC-*`/`ADR-*` — reine Ist-Messung plus Etappen-Vorschlag, keine Artefakt-Änderung.
**Bezug:** Maintainer-Auftrag 2026-08-29 („Können wir auf die neueste Version umstellen?" — „ja
machen wir"). Präzedenz: [slice-046](../done/slice-046-regelwerk-v352-migration-analyse.md), die
dieselbe Analyse für den Sprung `v1.3.0` → `v3.5.2` geleistet hat.
[Roadmap](../in-progress/roadmap.md).

> **Analyse zur Abnahme.** Es werden hier **keine** Kennungen vergeben und keine Artefakte
> geändert. Der Etappen-Schnitt (§6) gehört **vor** die Umsetzung abgenommen. Anders als bei
> slice-046 hält dieser Slice sich nicht selbst offen, bis die Abnahme kommt — das Warten auf das
> Maintainer-Wort hielt slice-046 zwei Wochen in `open/` und ist dort als Fehler notiert. Die
> Analyse **ist** die Lieferung; die Abnahme steuert die Folge-Slices.

**Geltungsbereich der Quelle:** gemessen und gelesen wurde ausschließlich `lab/regelwerk/` und
`lab/templates/` — das Artefakt, das dieses Repo vendored führt. Der Kurs unter `kurs/de/` ist
nicht Gegenstand.

---

## 1. Ist-Stand

| | |
|---|---|
| a-check gepinnt auf | **`v3.5.2`** (Kurs-Welle 34, 2026-07-24) |
| aktuelles Release | **`v5.12.0`** (Kurs-Welle 98, 2026-08-26) — drei Tage alt |
| dazwischen | **20 Releases**, **zwei Major-Sprünge** (v3 → v4 → v5), alle innerhalb eines Monats |

Der Pin steht an drei deklarierenden Stellen — `harness/conventions.md` §Baseline (kanonisch),
`AGENTS.md` §1, `harness/README.md` §Guides — und zusätzlich in **jedem** Pfad-Verweis auf
`.harness/baseline/v3.5.2/…` sowie im `scan.ignore` der `.d-check.yml`.

## 2. Umfang des Sprungs (gemessen)

`git diff --stat v3.5.2 v5.12.0`:

| Bereich | Dateien | Zeilen | Bestand |
|---|---|---|---|
| `lab/regelwerk/` | 29 geändert | **+2932 / −1012** | 21 → 26 Dateien |
| `lab/templates/` | 24 geändert | **+1073 / −359** | 21 → 25 Dateien |

Zum Vergleich: der letzte Sprung (`v1.3.0` → `v3.5.2`) war **+705 / −1868** über 21 Dateien. Dieser
ist der größere, und er ist **additiv** statt verdichtend — die umgekehrte Bewegung.

**Methodischer Nachtrag zu slice-046 §2 — belegt statt angenommen.** slice-046 hat ausdrücklich
das ZIP gelesen, „nicht den Git-Baum". Diese Messung läuft über den Git-Baum, deshalb wurde die
Gleichwertigkeit geprüft: der vendored Extrakt unter `.harness/baseline/v3.5.2/regelwerk/` weicht
vom Git-Baum des Tags in **allen 21 Dateien** ab, aber **jede einzelne** der 74 abweichenden
Zeilen ist eine umgeschriebene Link-Adresse (relativ → auf den Tag gepinnte GitHub-URL,
`tools/rewrite-doc-links.py` beim Release). **Null inhaltliche Abweichung.** Die
ZIP-Vorsicht gilt damit der *Provenienz*, nicht der *Delta-Messung* — für die ist der Baum
äquivalent.

## 3. Was sich **nicht** gedreht hat (Entwarnung)

Geprüft, weil dieses Repo frisch darauf gebaut hat und Sensoren daran hängen:

- **Slice-Lerneintrag:** die drei Formen (geschärfte Regel · neuer Sensor · benannte Spec-Lücke)
  stehen unverändert. `make verify` prüft sie zu Recht weiter.
- **WIP-Limit = 1** unverändert, inklusive der Lesart „harte Obergrenze".
- **Review-Harness:** `HIGH`/`MEDIUM`/`LOW`/`INFO`, die Negativbefund-Zeile je Bereich und das
  Feld `verifizierbar` stehen unverändert. Die Review-Reports dieses Repos sind **nicht** überholt.
- **Vendoring-Modell im Prinzip:** weiterhin `.harness/baseline/<tag>/{regelwerk,templates}/` aus
  dem self-contained Bundle. [`MR-006`](../../../../harness/conventions.md#mr-006--baseline-committet-vendored-statt-per-url-referenziert)
  behält seinen Gegenstand.

## 4. Die echten Brocken

### 4.1 Das Regelwerk schreibt die Migration jetzt selbst vor

`modul-02` §Freshness-Audit ist von drei auf **sieben Eigenschaften** gewachsen und enthält die
Prozedur, die dieses Repo bisher selbst erfinden musste. Normativ und für uns bindend:

- **Fünf Ausgänge je Adaptions-Eintrag** — *gegenstandslos* (Rückbau) · *bleibt gültig* ·
  *teilweise überholt* (Ablösung durch engeren Nachfolger) · *Bezug entfallen* · *widerspricht*.
  Nur der letzte lässt dem Repo eine Wahl; er verlangt eine **Entscheidung** in der Begründung.
- **Rückbau ist ein neuer Eintrag, kein Edit** — Append-only wie bei ADRs.
- **Das alte Baseline-Verzeichnis bleibt stehen**, bis der Review durch ist; die Form-Prüfung
  läuft als `diff -r .harness/baseline/<alt>/templates .harness/baseline/<neu>/templates`.
- **Auch `permanent`-Einträge werden mitgeprüft.** War eine Adaption eine *Lockerung* und die neue
  Baseline verschärft, ist die Antwort ein **Carveout mit Trigger**, keine stille Dauer-Adaption.

Das ist zugleich die Antwort auf die aus slice-047 §3 offen gebliebene Frage nach dem
**Freshness-Auslöser** — sie ist jetzt geregelt, nicht mehr uns überlassen.

### 4.2 Der Adaptions-Speicher ist strukturell neu

Einträge leben künftig **einzeln** unter `harness/conventions/MR-<NNN>-<titel>.md`;
`harness/conventions.md` trägt nur noch einen **Index** mit einer Zeile je Eintrag. Aufgelöste
Einträge wandern per `git mv` nach `conventions/done/` — der Zustand ist die Verzeichnis-Position,
kein Status-Feld, dieselbe Lifecycle-Form wie bei Slices.

Drei Konsequenzen, die dieses Repo direkt treffen:

1. **Neues Pflichtfeld `Ersetzt-Baseline-Regel`** — genau *eine* Baseline-Regel, als Link mit
   Abschnitts-Anker in die vendored Fassung. Die Baseline ist hier unmissverständlich: *„Ein
   Eintrag, der keine benannte Regel ersetzt, ist ein **Fork**, keine Adaption."* **Keiner** der
   zehn Einträge dieses Repos trägt das Feld heute.
2. **Anker-Bruch mit Vorschrift.** Die Index-Zeile braucht ein explizites `<a id="mr-NNN">`, und
   beim Umzug von der Inline- in die Verzeichnis-Form muss sie den **alten Überschriften-Slug
   zusätzlich** tragen — sonst rottet jeder veröffentlichte Verweis. Dieses Repo verlinkt seine
   Adaptionen genau über solche Slugs, und die Linkpflicht ist über `.d-check.yml` gegatet.
3. **Die Form bleibt Wahl**, der Default ist die Verzeichnis-Form. Bei zehn Einträgen ist das eine
   echte Entscheidung, keine Formalie.

### 4.3 Die gegatete Slice-Form driftet

Aus „höchstens drei **DoD**-Punkte" ist „höchstens drei **Liefer**-Punkte" geworden — mit
expliziter Nicht-Zähl-Liste: Gate-Läufe, Closure-Notiz, Register und Risiko-Ausgänge zählen
**nicht** mehr mit. Der Slice-Kopf trägt neu `Verantwortlich:`, `Autor:` und ein Feld für die
berührten Spec-Stellen. `make verify-slice-form` misst heute die alte Größe.

### 4.4 Neue Pflicht-Mechanik ohne Gegenstück im Repo

- **Risiko-Ausgänge bei Closure:** jedes notierte Risiko bekommt beim Übergang nach `done/` genau
  einen von drei Ausgängen (Carveout/Folge-Slice · gestrichen *mit Begründung* · Beobachtungs-Register).
  Geschlossene Menge, an der **Form** prüfbar — also sensor-fähig.
- **Beobachtungs-Register** in der Roadmap (`modul-06`, +188 Zeilen) — existiert hier nicht.
- **`## Leseordnung`** ist neuer Pflicht-Abschnitt von `harness/README.md`; die Datei hat ihn
  nicht (und trägt umgekehrt einen repo-eigenen Abschnitt zu Rollen-Übergaben).
- Drei neue Vorlagen: `observations`, `reconciliation`, `welle-results`.

### 4.5 Zwei Aussagen im Repo sind schon heute falsch

- **`AGENTS.md` §1 nennt `grundlagen-konventionen`** als nachzuschlagenden Abschnitt. Die Datei ist
  in `v5.12.0` **gelöscht** und auf sechs neue `grundlagen-*`-Dateien aufgeteilt. Die Auswahlregel
  bricht mit dem Re-Vendoring.
- **`harness/conventions.md` §Baseline ist veraltet** — der Text behauptet „der Sprung ist in vier
  Etappen geschnitten, diese ist Etappe A" und „bis diese Prüfung erfolgt ist, bleiben die
  deklarierten Adaptionen in Kraft". Tatsächlich hat `welle-12` die Etappen **A–F** am 2026-08-09
  abgeschlossen. Der Abschnitt ist auf dem Stand von slice-047 stehengeblieben.

## 5. Was diese Analyse **nicht** geklärt hat

- **Gelesen wurden fünf Quellen:** `modul-02` (Freshness/Vendoring), `grundlagen-harness-dateien`
  (Pflichtgliederungen), die neue Adaptions-Vorlage, der Regelwerk-Index und Teile von `modul-05`.
  **Nur im Umfang vermessen, nicht gelesen:** `modul-06` (+188), `modul-12` (+223), `modul-08`
  (+153), `modul-09` (+76), `modul-13`, `modul-03`, `modul-04` und fünf der sechs neuen
  `grundlagen-*`-Dateien.
- **Das Urteil je Adaption** (welcher der fünf Ausgänge für welchen der zehn Einträge) ist
  ausdrücklich **nicht** gefällt — das ist die Arbeit, nicht die Analyse.
- **Ob das Release weiterhin ein `lab-regelwerk.zip` als Asset trägt.** Im Baum existiert der
  Bau-Weg; das Asset selbst wurde nicht abgerufen.
- **Ob d-checks `sources`-Modul die Integritäts-Hälfte jetzt abdecken kann** — die Baseline nennt
  es neuerdings namentlich.

## 6. Vorschlag: vier Etappen

Die Reihenfolge folgt der Prozedur aus §4.1, nicht eigener Erfindung.

| Etappe | Inhalt | Warum getrennt |
|---|---|---|
| **A — Re-Vendoring** | `.harness/baseline/v5.12.0/` **neben** `v3.5.2` anlegen (die Baseline verlangt das alte Verzeichnis bis zum Ende des Reviews), `SHA256SUMS`, Pin an den drei Stellen, Modulnamen in `AGENTS.md` §1 korrigieren, §Baseline-Text auf den Ist-Stand bringen (§4.5) | mechanisch, sofort prüfbar, Voraussetzung für den netzlosen Vergleich beider Formen |
| **B — Adaptions-Durchgang** | die zehn Einträge einzeln durch die fünf Ausgänge, jeweils mit nachgetragenem `Ersetzt-Baseline-Regel`; Ergebnis entscheidet C | berührt die Konventions-Identität; das Urteil je Eintrag ist Lesearbeit, kein Diff |
| **C — Form** | Verzeichnis-Form der Adaptionen inkl. Index-Anker und Alt-Slugs, `## Leseordnung`, Slice-Kopffelder und die Metrik von `verify-slice-form` | erst nach B, weil B den Bestand bestimmt, der die Form füllt |
| **D — Neue Mechanik** | Risiko-Ausgänge bei Closure (sensor-fähig), Beobachtungs-Register, ggf. `sources` für die Asset-Integrität | echte Funktionalität, kein Formabgleich — eigener Nachweis je Sensor |

Danach fällt `.harness/baseline/v3.5.2/`.

**Nicht in diesem Vorschlag:** ob der Sprung überhaupt jetzt kommt. Bei 20 Releases im Monat ist
die Alternative, den Auslöser aus §4.1 zu deklarieren und den Pin bewusst stehen zu lassen — das
wäre ebenfalls eine begründete Antwort und ist billiger als vier Etappen.

## 7. DoD

- [x] Sprung-Umfang gemessen und die Gleichwertigkeit Baum ↔ vendored Extrakt **belegt** (§2).
- [x] Die Brocken benannt und je Brocken die betroffene Repo-Stelle genannt (§4), die Lücken der
      Analyse ausgewiesen statt als Vollständigkeit ausgegeben (§5).
- [x] Etappen-Vorschlag steht (§6); `make gates` und `make verify` grün — **Ausgabe in eine
      Datei**, Exit-Code getrennt geprüft, nie in eine Pipe.

## 8. Closure-Notiz

**Geliefert:** die Messung des Sprungs `v3.5.2` → `v5.12.0` über den vendored Artefakt-Umfang, vier
benannte Migrations-Brocken mit je einer betroffenen Repo-Stelle, drei Entwarnungen, zwei bereits
heute falsche Aussagen im Repo — und ein Etappen-Vorschlag, der der Prozedur folgt, die die neue
Baseline inzwischen selbst vorschreibt.

**Lerneintrag — Form: geschärfte Regel.** *Eine methodische Vorsicht, die aus einem früheren Slice
übernommen wird, ist eine Aussage über die damalige Lage — sie wird gemessen, bevor sie befolgt
oder verworfen wird.* Konkret: slice-046 §2 hält fest, gelesen werde „das kanonische Artefakt …
nicht der Git-Baum". Beide naheliegenden Umgänge damit wären falsch gewesen — blind befolgen hätte
diese Analyse an ein Release-Asset gebunden, das netzlos nicht vorliegt; blind übergehen hätte eine
unbelegte Messgrundlage geliefert. Ein `diff -r` plus ein Zeilen-Filter hat die Frage in zwei
Befehlen entschieden: 74 abweichende Zeilen, **alle** Link-Umschreibungen, null inhaltliche
Abweichung. Aus der geerbten Vorsicht wurde damit eine **belegte Äquivalenz** — die Vorsicht gilt
der Provenienz, nicht der Delta-Messung. *Weil* eine übernommene Vorsicht sonst entweder zur
unbegründeten Blockade oder zur stillen Nachlässigkeit wird, und beides sich gleich anfühlt.

**Zwei beobachtbare Closure-Kriterien:**

1. Die Zahlen in §2 sind mit `git diff --stat v3.5.2 v5.12.0 -- lab/regelwerk` nachrechenbar, und
   §5 nennt die **nicht** gelesenen Module namentlich statt Vollständigkeit zu behaupten.
2. `make gates` und `make verify` je **Exit 0** auf dem Stand dieses Slice, Ausgabe in Dateien,
   Exit-Codes getrennt geprüft.

**Offene Risiken und ihr Ausgang.** Die neue Baseline verlangt das (§4.4); adoptiert ist die Regel
noch nicht, befolgt wird sie hier trotzdem:

- *Das Urteil je Adaption steht aus* — Ausgang: **Folge-Slice**, Etappe B aus §6.
- *Sieben Module sind nur vermessen, nicht gelesen* — Ausgang: **Folge-Slice**, Etappe B; die
  Lücke ist in §5 namentlich ausgewiesen, nicht stillgelegt.
- *Ob der Sprung überhaupt jetzt kommt, ist offen* — Ausgang: **Maintainer-Entscheidung**, in §6
  ausdrücklich als Alternative formuliert (Auslöser deklarieren, Pin bewusst stehen lassen).

**Folge-Slices:** keine vergeben. Vier Etappen A–D sind in §6 vorgeschlagen und brauchen die
Abnahme, bevor sie IDs bekommen — Slice-IDs werden referenziert, nicht auf Vorrat erfunden.

## 9. Sub-Area-Modus

Berührt wird ausschließlich **Planungs-Harness** (`docs/plan/planning/`) — Greenfield. Alle
berührten Sub-Areas GF.
