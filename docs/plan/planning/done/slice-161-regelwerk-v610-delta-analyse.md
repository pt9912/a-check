# slice-161 — Regelwerk-Migration `v6.0.0` → `v6.1.0`: Delta-Analyse

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `make slice-mv` ([`AGENTS.md`](../../../../AGENTS.md)
§3.3/§5).

**Welle:** [welle-14](../welle-14-regelwerk-v610-migration.md).

**Bezug:** Maintainer-Anfrage 2026-09-05 ("Es gibt ein neues Regelwerk
v6.1.0"). Präzedenz: [slice-135](../done/wellenlos/slice-135-regelwerk-v600-delta-analyse.md),
dieselbe Analyse-Form für den vorigen Sprung `v5.12.0` → `v6.0.0`.

**Berührte Spec-Stellen:** — *(keine)* — reine Ist-Messung, keine
Vertragsberührung.

**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim
Maintainer.

**Autor:** Claude (Sonnet 5), im Auftrag des Maintainers. **Datum:**
2026-09-05.

> **Analyse zur Abnahme.** Wie bei slice-135: keine Kennungen vergeben,
> keine Artefakte geändert. Der Etappen-Schnitt (§6) gehört vor der
> Umsetzung abgenommen. Die Analyse **ist** die Lieferung.

**Geltungsbereich der Quelle:** gemessen und gelesen wurde ausschließlich
`lab/regelwerk/` und `lab/templates/` (das vendored Artefakt) aus einem
frischen Klon von `pt9912/ai-harness-course`. Der Kurs unter `kurs/de/` ist
nicht Gegenstand.

---

## 1. Ist-Stand

| | |
|---|---|
| a-check gepinnt auf | **`v6.0.0`** (Kurs-Welle 116, 2026-09-03) |
| aktuelles Release | **`v6.1.0`** (Kurs-Welle 118, 2026-09-05) — am Tag der Analyse |
| dazwischen | zwei Kurs-Wellen (117, 118); Welle 117 berührt laut Commit-Historie nur `docs/` des Kurs-Repos, nicht das Bundle |

Der Pin steht an drei deklarierenden Stellen — `harness/conventions.md`
§Baseline (kanonisch), `AGENTS.md` §1, `harness/README.md` §Guides.

## 2. Umfang des Sprungs (gemessen)

`git diff --stat v6.0.0 v6.1.0 -- lab/regelwerk lab/templates` (Roh-Git-Tags,
unabhängig vom vendored Extrakt aus einem frischen Klon): **6 Dateien,
`+38/−1`**.

| Datei | Zeilen |
|---|---|
| `lab/regelwerk/README.md` | Stand-Zeile (Kurs-Welle 116→118), reine Provenienz |
| `lab/regelwerk/modul-07-carveouts.md` | +3 |
| `lab/regelwerk/modul-10-review-harness.md` | +7 |
| `lab/regelwerk/modul-13-quality-gates.md` | +6 |
| `lab/templates/AGENTS.template.md` | +7 |
| `lab/templates/harness/README.template.md` | +14 (zwei Hunks) |

Zum Vergleich: der letzte Sprung (`v5.12.0` → `v6.0.0`, slice-135) war
`+579/−110` über 26 Dateien. Dieser Sprung ist **rund ein Fünfzehntel des
Volumens** — der bisher kleinste seit slice-092/135.

**Provenienz unabhängig geprüft:** `gh release view v6.1.0 --repo
pt9912/ai-harness-course` bestätigt Tag, Assets (`lab-regelwerk.zip`,
`SHA256SUMS`) und Datum; der Diff oben lief gegen einen frisch geklonten
Git-Tag, nicht gegen das noch nicht vendierte ZIP-Asset — die Byte-Identität
von Klon-Tag und Release-ZIP ist damit **noch offen** (siehe §5) und wird
erst mit der Vendoring-Umsetzung selbst geprüft (`sha256sum -c`, wie bei
jedem `regelwerk-check`-Lauf).

## 3. Was sich **nicht** gedreht hat (Entwarnung — geprüft, nicht angenommen)

- **Review-Deckung (`modul-10`) ist bereits erfüllt — und übertroffen.** Der
  Zuwachs benennt die Unterscheidung Kategorisierung (inferential) vs.
  Deckung (mechanisierbar, Beispiel `d-check`-Modul `reviews`). a-check hat
  das mit slice-159/160 bereits umgesetzt: `make doc-reviews` steht
  ausführlich in [`AGENTS.md`](../../../../AGENTS.md) §4 (Modul `reviews`,
  `DC-FA-RVW-001`, im `gates`-Aggregat seit slice-160) und §6 trägt bereits
  einen expliziten Rollenwechsel-Absatz nach Schritt 8. Kein Handlungsbedarf.
- **`AGENTS.template.md`-Zuwachs (Rollenwechsel statt Abschluss nach Schritt
  8) ist bereits erfüllt — und übertroffen.** a-checks
  [`AGENTS.md`](../../../../AGENTS.md) §6 trägt seit slice-159 (2026-09-05,
  heute) den Absatz "Vor
  dem Abschluss…" mit Kontext-Trennungs-Pflicht (`fork` zählt explizit
  nicht) und Zitat des auslösenden Befunds (23 Slices ohne Review,
  [`BEO-HARNESS/behauptete-vollstaendigkeit-extern-gefangen`](../observations/BEO-HARNESS/behauptete-vollstaendigkeit-extern-gefangen/observation.md)).
  Der neue Kurstext ist knapper als a-checks eigener. Kein Handlungsbedarf.
- **`harness/README.template.md`-Sensors-Kommentar zu den Modulen `reviews`/
  `planning` (≥ `d-check` `v0.73.0`) ist bereits erfüllt.** a-check pinnt
  `d-check` auf `v0.74.1` (slice-160); beide Ziele (`doc-reviews`,
  `doc-planning`) existieren im Makefile und stehen im `gates`-Aggregat.
  **Wichtig:** dieser Zuwachs steht komplett **innerhalb** des
  `<!-- … -->`-Kommentarblocks des Templates (Bedienhinweis für
  Template-Füller) — er ist kein Text, der wörtlich in `harness/README.md`
  erscheinen müsste, anders als die Tabellenzeile in §4.1 unten.

## 4. Die echten Brocken

### 4.1 Reviewer-Skill fehlt als Zeile in `harness/README.md` §Guides (echter, kleiner Nachzug)

`lab/templates/harness/README.template.md` bekommt eine neue Tabellenzeile
in der **Guides**-Tabelle (außerhalb jedes Kommentarblocks, also echter
Pflichttext):

> `.harness/skills/reviewer.md` — Reviewer-Skill: HIGH-Liste,
> Kategorien-Regeln, Negativbefund-Pflicht, Output-Schema (Modul 10) —
> nächste Rolle nach Schritt 8 des Minimal Agent Workflow, nicht Teil der
> Implementer-Eingabe.

a-checks [`harness/README.md`](../../../../harness/README.md) §Guides
(Zeilen 40–46) listet Spec, ADRs, Planning, `AGENTS.md`, `conventions.md`
und die vendored Baseline — **keine** Zeile für
[`.harness/skills/reviewer.md`](../../../../.harness/skills/reviewer.md),
obwohl `AGENTS.md` §6 selbst mehrfach darauf verweist. Der Pointer fehlt an
genau der einen Stelle, die "was lenkt den Agenten *vor* der Handlung"
versammelt. Kleiner, konkreter Nachzug.

### 4.2 Carveout-Auflösung: Bindung-Spalten-Rückstellung nicht benannt (Kandidat für Vertagung)

`modul-07-carveouts.md` +3 Zeilen: Auflösung eines Carveouts setzt auch die
Bindung-Spalte in `harness/README.md` §Sensors zurück (`CO-<NNN>` →
Normalzustand) — sonst behauptet die Tabelle einen Carveout, der längst
geschlossen ist.

a-check führt **null** aktive Carveouts
([`docs/plan/carveouts/README.md`](../../carveouts/README.md) §Aktueller
Bestand). Kein akuter Trigger — die Regel hat aktuell nichts, worauf sie
wirken könnte. Der Nachzug wäre, den Satz in
`docs/plan/carveouts/README.md` §Form eines Carveouts zu ergänzen, aber das
ist reine Prozess-Dokumentation ohne Wirkung, solange kein Carveout
aufgelöst wird — dieselbe Einordnung wie slice-135 §4.2
(Zeitdokumente-Archivierung: optional, ohne Trigger).

### 4.3 Bootstrap-aware-Gate-Hochschaltung: Hard-Rule-Entfernung nicht als Erledigungs-Pflicht benannt (Kandidat für Vertagung)

`modul-13-quality-gates.md` +6 Zeilen: feuert der Hochschalt-Trigger eines
bootstrap-aware Gates, ist die Entfernung der zugehörigen Hard-Rule-Zeile
aus `AGENTS.md` eine **Erledigungs-Pflicht des auslösenden Slice** (in der
Baseline-Formulierung: "DoD-Punkt") — derselbe Träger,
den der Carveout für seine Bindung-Spalte hat.

**Nicht tief geprüft:** ob a-check aktuell einen bootstrap-aware Gate mit
nahendem Hochschalt-Trigger führt (Kandidaten wären das AC-Grandfathering
oder der `commit-scope-check`-Stichtag, beide in `AGENTS.md` §5 erwähnt) —
dafür wäre jeder einzeln gegen seinen Trigger zu lesen, was diese Analyse
nicht geleistet hat. Ohne diese Prüfung ist unklar, ob der Nachzug akut
ist oder wie §4.2 aufschiebbar. **Grenze der Analyse, nicht Entwarnung.**

### 4.4 a-checks eigene Review-Pflicht/Rollenwechsel-Ergänzung in `AGENTS.md` §6 ist keine deklarierte Adaption

Bei der Nachfrage zu §3 (Entwarnung `AGENTS.template.md`) aufgefallen, nicht
Gegenstand des ursprünglichen Diffs: die vendored `v6.0.0`-Vorlage
[`AGENTS.template.md`](../../../../.harness/baseline/v6.0.0/templates/AGENTS.template.md)
endet nach Schritt 8 des Minimal Agent Workflow (Zeile 236) — kein
Rollenwechsel-Absatz, keine Review-Pflicht-Formulierung. a-checks eigenes
[`AGENTS.md`](../../../../AGENTS.md) §6 trägt **seit heute** (Commit
`7de8569`, slice-159, 2026-09-05 — `git log -p --follow -- AGENTS.md`
zeigt genau diese eine Einführung) einen deutlich umfangreicheren Absatz
danach ("Vor dem Abschluss, bei jedem Slice mit Code- oder
Vertragsänderung: Review …"). **Korrektur ggü. einer früheren Fassung
dieses Punkts:** nicht seit slice-050 (2026-07-25) — jener Commit ändert
`AGENTS.md` an keiner Stelle; die eigene Zitat-Quelle zwei Sätze weiter
unten trug „seit slice-159" bereits richtig, der Widerspruch wurde erst im
Review sichtbar. Das ist also keine sechs Wochen alte, unbelegte Drift,
sondern eine **taggleiche** Selbst-Ergänzung — Kern-Befund unverändert
(kein `MR-<NNN>`-Eintrag), nur die historische Einordnung war falsch.

Trotzdem eine echte a-check-eigene Ergänzung ggü. der Baseline — **ohne**
`MR-<NNN>`-Eintrag in
[`harness/conventions.md`](../../../../harness/conventions.md#adaptions-block)
deklariert (`grep` über den Adaptions-Block: kein Treffer für
"Rollenwechsel"/"Review-Pflicht"/"Code- oder Vertragsänderung"). Weil sie
taggleich mit diesem Slice entstand, ist es noch keine Drift, sondern eine
Lücke, die sich schließen lässt, bevor sie eine wird.

`v6.1.0` fügt jetzt selbst einen Rollenwechsel-Absatz in
`lab/templates/AGENTS.template.md` ein (§2 oben, +7 Zeilen), inhaltlich
deckungsgleich mit a-checks Praxis, aber knapper:

> Dieser Workflow deckt ausschließlich die Implementer-Rolle ab. Schritt 8
> ist der Rollenwechsel, kein Abschluss: Bericht → Handoff an Reviewer
> (`.harness/skills/reviewer.md`, siehe `harness/README.md` §Guides) →
> Verifier. Kein Self-Review — anderer Kontext findet andere Findings,
> derselbe Kontext dieselben blinden Flecken (Baseline-Regelwerk
> `modul-08-agentenrollen.md`).

**Maintainer-Vorgabe für den Nachzug:** so nah wie möglich an diesem
`v6.1.0`-Wortlaut bleiben, statt die eigene, bereits gewachsene
Formulierung fortzuschreiben. Das heißt für die Umsetzung: `AGENTS.md` §6
auf einen Absatz **in dieser Kürze** zurückschneiden und nur die zwei
a-check-spezifischen Zusätze behalten, die der Baseline-Wortlaut nicht
trägt und die ein echter, gemessener Ausfall erzwungen hat — den
`fork`-Ausschluss (Modul 8 §Kontext-Trennung) und den Zitat-Anker auf
[`BEO-HARNESS/behauptete-vollstaendigkeit-extern-gefangen`](../observations/BEO-HARNESS/behauptete-vollstaendigkeit-extern-gefangen/observation.md).
Ein neuer `MR-<NNN>` dokumentiert dann diese **verbleibende** Differenz zum
Baseline-Wortlaut — nicht mehr die heutige, umfangreiche Fassung.

**Nicht Gegenstand dieses Punkts:** ob `make doc-reviews`
(`DC-FA-RVW-001`) selbst korrekt an diesen Absatz bindet — das prüft §3
oben bereits als erfüllt.

## 5. Was diese Analyse **nicht** geklärt hat

- **Byte-Identität Release-ZIP ↔ geklonter Tag** ist — anders als bei
  slice-135 — noch nicht geprüft; sie folgt erst mit der Vendoring-Umsetzung
  selbst (`sha256sum -c` gegen das von `gh release` gemeldete Asset-Digest).
- **§4.3** (bootstrap-aware Gates mit nahendem Hochschalt-Trigger) ist nicht
  einzeln durchgeprüft — s. dort.
- **§4.4** (Review-Pflicht/Rollenwechsel-Absatz auf `v6.1.0`-Wortlaut
  zurückschneiden) ist als Nachzug benannt, aber nicht selbst umgesetzt —
  das wäre eine Textänderung an `AGENTS.md` und damit außerhalb einer
  "Analyse zur Abnahme".

## 6. Vorschlag: eine Etappe plus ein benannter Folge-Slice

Anders als beim letzten Sprung (drei Etappen, slice-135 §6) reicht für
`.harness/baseline/` + die drei Stand-Stellen **eine einzige Etappe** — der
Umfang ist klein genug, um innerhalb der Drei-Liefer-Punkte-Grenze
(Modul 5) zu bleiben:

| Liefer-Punkt | Inhalt |
|---|---|
| 1 | `.harness/baseline/v6.1.0/{regelwerk,templates}/` + `SHA256SUMS` **neben** `v6.0.0` anlegen, `make regelwerk-check` grün |
| 2 | Stand-Deklaration an den drei Stellen (`conventions.md` §Baseline, `AGENTS.md` §1, `harness/README.md` §Guides) auf `v6.1.0` |
| 3 | `harness/README.md` §Guides: Zeile für `.harness/skills/reviewer.md` ergänzen (§4.1) |

§4.2 und §4.3 bleiben außerhalb dieser Etappe (Vertagung mangels akutem
Trigger bzw. ungeklärt).

**§4.4 braucht einen eigenen Folge-Slice**, nicht Platz in dieser Etappe —
er würde die Drei-Liefer-Punkte-Grenze sprengen (Textänderung `AGENTS.md`
§6 **und** neuer `MR-<NNN>`-Eintrag sind zwei weitere Punkte) und ist
inhaltlich unabhängig vom Vendoring selbst: Der Rückschnitt auf den
`v6.1.0`-Wortlaut (Maintainer-Vorgabe, §4.4) braucht kein neues
`.harness/baseline/v6.1.0/` — er kann auch **vor** Etappe A laufen. Noch
keine Slice-ID vergeben, wie bei jedem Folge-Slice-Vorschlag dieser
Analyse.

**Warum ein Wellen-Konstrukt (`welle-14`), wo der Präzedenzfall wellenlos
lief:** Der strukturell identische vorige Sprung (`v5.12.0`→`v6.0.0`,
slice-135/136/139/141) trug keine Welle — nach Buchstaben des
"Zeremonie"-Tests aus `modul-06-roadmap.md` §Wann Arbeit eine Welle
braucht wäre das auch hier die naheliegende Wahl gewesen, denn keiner der
drei Liefer-Punkte oben ist ein "Mehr" über die eigene Slice-DoD hinaus.
`welle-14` entstand stattdessen auf **explizite Maintainer-Weisung**
("Bitte eine Welle dazu anlegen") — eine legitime, aber bewusst
abweichende Entscheidung, keine aus Modul 6 selbst abgeleitete Pflicht.
Der eine inhaltliche Unterschied zum Vorbild, der ein "Mehr" liefert: die
Ergebnis-Notiz `welle-14-results.md` fasst **zwei** unabhängig eröffnete
Slices zusammen (die Etappe hier **und** den §4.4-Folge-Slice) — beim
Vorbild entstand eine vergleichbare Zusammenschau nur retroactiv, über
den externen Multi-Linsen-Review
[`docs/reviews/2026-09-05-slice135-157-multi-linsen-review.md`](../../../reviews/2026-09-05-slice135-157-multi-linsen-review.md),
nachdem 23 Slices ohne diese Klammer gelaufen waren. Damit ist die
Wellen-Wahl **beantwortet, nicht offen** — siehe §7.

## 7. Risiken und offene Punkte

- *Das Urteil über §4.2 (Carveout-Bindung-Rückstellung) steht aus* —
  **Ausgang:** weiter offen → Beobachtungs-Register, solange kein Carveout
  aufgelöst wird.
- *§4.3 (bootstrap-aware Gate, Hochschalt-Trigger) ist ungeprüft* —
  **Ausgang:** weiter offen → Beobachtungs-Register, bis eine künftige
  Berührung eines dieser Gates die Prüfung erzwingt.
- *§4.4 (Review-Pflicht/Rollenwechsel-Absatz in `AGENTS.md` §6 ohne
  `MR`-Deklaration, jetzt auf `v6.1.0`-Wortlaut zurückzuschneiden) ist
  benannt, aber nicht umgesetzt* — **Ausgang:** Folge-Slice (noch keine ID,
  §6).
- *Ob `welle-14` gegenüber dem wellenlosen Vorbild gerechtfertigt ist*
  (Review-Finding F-3) — **Ausgang:** gestrichen mit Begründung: kein
  Risiko im Sinn der Dreier-Menge, sondern eine bereits getroffene
  Maintainer-Entscheidung mit benannter Zusatz-Begründung (§6); kein
  Zustand, der noch "eintreten" könnte.
- *Ob die eine vorgeschlagene Etappe (§6) so umgesetzt wird, ist
  Maintainer-Entscheidung* — **Ausgang:** gestrichen mit Begründung: kein
  Risiko im Sinn der Dreier-Menge, sondern die Kernfrage, die diese Analyse
  laut ihrem eigenen Kopf-Vermerk ("Analyse zur Abnahme") beantwortet
  bekommt, statt sie vorwegzunehmen.

## 8. DoD

- [x] Sprung-Umfang gemessen, Herkunft des Release-Tags unabhängig
      bestätigt, Brocken benannt und je Brocken der betroffene
      a-check-Bestand **selbst geprüft**, Entwarnungen verifiziert, Grenzen
      der Analyse ausgewiesen statt Vollständigkeit behauptet (§2–§5).
- [x] Etappen-Vorschlag steht, Wellen-Wahl begründet (§6/§7).
- [x] Unabhängiges Plan-Review über getrennten Kontext durchgeführt
      (`docs/reviews/2026-09-05-slice-161-delta-analyse.md`); HIGH- und
      gate-verifizierte Findings (F-1, F-2) in dieser Fassung behoben.
- [x] `make gates` grün.
- [x] `make verify` grün.
- [x] Beobachtungs-Register fortgeschrieben (§9 Closure-Notiz).
- [x] Jedes Risiko aus §7 trägt einen Ausgang.

## 9. Closure-Notiz

- **Was hat funktioniert:** die unabhängige Gegenrechnung der Kern-Zahlen
  (Diff-Stat, wörtliche Zitate, `AGENTS.template.md`-Ende) über einen
  frischen Klon des Kurs-Repos hat sich bewährt — der Reviewer konnte jede
  Messung reproduzieren, ohne sie zu übernehmen (Negativbefunde in
  [`docs/reviews/2026-09-05-slice-161-delta-analyse.md`](../../../reviews/2026-09-05-slice-161-delta-analyse.md)).
- **Was ging anders als geplant:** das erste Review (kein Fork, getrennter
  Kontext) fand eine faktisch falsche Provenienz-Behauptung (§4.4, „seit
  slice-050" statt tatsächlich „seit slice-159, heute") und einen echten
  `make doc-structure`-Ausfall (Heading-Kollision mit dem DoD-Größen-Sensor)
  — beide vor Abschluss korrigiert. Zusätzlich war die
  Beobachtungs-Register-Sichtung in §10 zunächst unvollständig (5 statt 18
  Treffer über zwei statt drei Sub-Areas).
- **Steering-Loop-Eintrag:** geschärfte Regel — eine Provenienz-Behauptung
  über eine andere Repo-Datei ("seit slice-NNN") wird gegen
  `git log -p --follow -S"<Textfragment>"` verifiziert, bevor sie in einen
  Slice-Plan geschrieben wird, statt aus dem Gesprächskontext übernommen zu
  werden — dieselbe Disziplin, die `cr-text-behauptet-statt-gemessen`
  bereits für CR-Texte an ein fremdes Werkzeug verlangt, jetzt auch für
  repo-eigene Historie.
- **Beobachtungs-Register (`../observations/`):**
  [`BEO-PLAN/slice-provenienz-aus-gedaechtnis-statt-git-log`](../observations/BEO-PLAN/slice-provenienz-aus-gedaechtnis-statt-git-log/observation.md)
  neu angelegt, Beleg `evidence/slice-161.md` — Zähler steht bei 1×.
- **Folge-Slices:** noch keine ID vergeben — zwei vorgeschlagen (§6):
  die Vendoring-Etappe (`.harness/baseline/v6.1.0/` + 3-Stellen-Pin-Bump +
  Reviewer-Skill-Zeile) und der §4.4-Folge-Slice (Review-Pflicht-Absatz in
  `AGENTS.md` §6 auf `v6.1.0`-Wortlaut zurückschneiden, neuer `MR`-Eintrag —
  Kennung wird bei dessen eigener Anlage vergeben, nicht hier).
- **Risiken aus §7:** alle vier mit Ausgang (zwei weiter offen →
  Beobachtungs-Register, zwei gestrichen mit Begründung) — siehe §7.
- **Drei Paarungen:** verschoben auf die Closure von `welle-14` — dieser
  Slice trägt ein `**Welle:**`-Feld, die Paarungs-Prüfung ist damit nach
  Modul 8 §Rollen-Sequenz für eine Welle Sache der Wellen-Closure, nicht
  der Einzel-Slice-Closure.

## 10. Sub-Area-Modus

**Vorgelagert — Sub-Area-Wahl prüfen:** drei Sub-Areas berührt —
**Vendored Baseline** (`.harness/baseline/`, kein Modus, externer Fremdtext,
[`MR-006`](../../../../harness/conventions.md#mr-006--baseline-committet-vendored-statt-per-url-referenziert)),
**Harness-Einstieg** (`AGENTS.md`, `harness/README.md`) und
**Gate-/Werkzeug-Schicht** (`tools/regelwerk-check.sh`, von
Liefer-Punkt 1 aus §6 aufgerufen) — alle drei Greenfield, alle erfüllen
die Schwelle ≥ 2/3 laut
[`harness/conventions.md`](../../../../harness/conventions.md#modus-deklaration-pro-sub-area).
Die dritte Sub-Area fehlte in einer früheren Fassung dieses Abschnitts —
Review-Finding F-5 (zwei zitierte `BEO-GATE`-Einträge gehörten zu einer
nicht deklarierten Sub-Area).

**Vorgelagert — offene Beobachtungen sichten:** Register vollständig
durchgegangen (`grep` über alle `state.md`, nicht nur die zunächst
erinnerten Pfade — Review-Finding F-4 hatte die erste Fassung als
unvollständig markiert):

- **Harness-Einstieg** — 7 `offen` (2 weitere `verkörpert`, korrekt
  ausgeschlossen): [`BEO-HARNESS/adaption-korrigiert-repo-aussage`](../observations/BEO-HARNESS/adaption-korrigiert-repo-aussage/observation.md),
  [`baseline-normtext-nachgeschrieben`](../observations/BEO-HARNESS/baseline-normtext-nachgeschrieben/observation.md),
  [`chronik-in-gelesenen-dateien`](../observations/BEO-HARNESS/chronik-in-gelesenen-dateien/observation.md),
  [`hard-rule-37-ohne-sensor`](../observations/BEO-HARNESS/hard-rule-37-ohne-sensor/observation.md),
  [`rueckbau-eintrag-ablage-auslegung`](../observations/BEO-HARNESS/rueckbau-eintrag-ablage-auslegung/observation.md),
  [`rueckbau-kandidat-ueberlebt-baseline-migration`](../observations/BEO-HARNESS/rueckbau-kandidat-ueberlebt-baseline-migration/observation.md),
  [`selbst-archivierung-verdoppelt-abschluss-aufwand`](../observations/BEO-HARNESS/selbst-archivierung-verdoppelt-abschluss-aufwand/observation.md)
  — alle bei 1×.
- **Gate-/Werkzeug-Schicht** — 11 `offen` (5 weitere `verkörpert`, korrekt
  ausgeschlossen): [`BEO-GATE/awk-sed-verliert-ausfuehrungsmodus`](../observations/BEO-GATE/awk-sed-verliert-ausfuehrungsmodus/observation.md),
  [`closure-platzhalter-legitimer-treffer`](../observations/BEO-GATE/closure-platzhalter-legitimer-treffer/observation.md),
  [`gate-beleg-trennung-nur-ueberflogen`](../observations/BEO-GATE/gate-beleg-trennung-nur-ueberflogen/observation.md),
  [`guard-trifft-erwaehnung-statt-aufruf`](../observations/BEO-GATE/guard-trifft-erwaehnung-statt-aufruf/observation.md),
  [`hebungskanal-haengt-an-repo-externen-schaltern`](../observations/BEO-GATE/hebungskanal-haengt-an-repo-externen-schaltern/observation.md),
  [`pruefer-ohne-gegenstand-oder-aufruf`](../observations/BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf/observation.md),
  [`risiko-block-existenz-ungeprueft`](../observations/BEO-GATE/risiko-block-existenz-ungeprueft/observation.md),
  [`sources-modul-asset-integritaet`](../observations/BEO-GATE/sources-modul-asset-integritaet/observation.md),
  [`symlink-ziel-nach-baseline-bump-ungeprueft`](../observations/BEO-GATE/symlink-ziel-nach-baseline-bump-ungeprueft/observation.md),
  [`ungelaufene-mechanik-docker-hub-spiegel`](../observations/BEO-GATE/ungelaufene-mechanik-docker-hub-spiegel/observation.md),
  [`weiche-haengt-an-fremdem-event-feld`](../observations/BEO-GATE/weiche-haengt-an-fremdem-event-feld/observation.md)
  — alle bei 1×.

**Keiner der 18 Treffer erreicht mit dieser Analyse 3×** — die
Schwellen-Aussage der ersten Fassung war richtig, nur die Zähl- und die
Sub-Area-Zuordnung waren es nicht (F-4/F-5). Eine mögliche zweite Berührung
entsteht erst mit der Umsetzungs-Etappe (§6), nicht mit der Analyse selbst.

**Alle berührten Sub-Areas GF** — kein Begründungsblock nötig
(Vendored Baseline führt ohnehin keinen Modus).
