# slice-173 — Baseline-Pins vom `versions`-Modul wächtern lassen

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `make slice-mv` ([`AGENTS.md`](../../../../AGENTS.md)
§3.3/§5).

**Welle:** ohne Welle (der Closure-Trigger wäre die eigene DoD — kein
repo-weites Mehr).

**Bezug:** Fund des unabhängigen Reviews zu
[slice-172](../done/wellenlos/slice-172-baseline-v600-entfernen.md) (F-3,
MEDIUM). Dort §7, Risiko 3: der Wächter für Baseline-Pins ist nicht zu
erfinden, sondern **unkonfiguriert**.

**Berührte Spec-Stellen:** — *(keine)* — Gate-Konfiguration ohne
Vertragsberührung.

**Verantwortlich:** — *(noch nicht priorisiert)*

**Autor:** Claude (Opus 5), im Auftrag des Maintainers. **Datum:**
2026-09-06.

---

## 1. Ziel

Ein nicht nachgezogener `.harness/baseline/<tag>/`-Pfad wird ein Befund,
statt still auf einen alten Stand zu zeigen, solange dieser noch existiert.

## 2. Analyse (vor der Umsetzung)

Die vendored Ziel-Form des Adaptions-Eintrags
([`MR-NNN-titel.template.md`](../../../../.harness/baseline/v6.2.0/templates/harness/conventions/MR-NNN-titel.template.md))
benennt den Wächter ausdrücklich: *„jeder Baseline-Bump entwertet `<tag>` —
der adoptierte Stand steht einmal im Adaptions-Block, ein Versions-Sensor
prüft jeden Pin dagegen; Muster in `.d-check.yml`"*. Die vendored
[`templates/.d-check.yml`](../../../../.harness/baseline/v6.2.0/templates/.d-check.yml)
liefert es fertig aus:

```yaml
# versions:
#   pin-pattern: '\.harness/baseline/(v\d+\.\d+\.\d+)/'
#   current-from: harness/conventions.md#baseline
#   exempt-paths: ["harness/conventions/done/**"]   # aufgeloeste Eintraege sind eingefroren
```

**Ist-Stand gemessen (2026-09-06):** `versions` ist in
[`.d-check.yml`](../../../../.d-check.yml) konfiguriert, trägt aber **ein**
Muster — die Lastenheft-Version. Baseline-Pins deckt es nicht, obwohl **35**
Dateien welche tragen. Ein vorhandener Prüfer ohne Gegenstand, dieselbe
Klasse wie
[`BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf`](../observations/BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf/observation.md).

### 2.1 Die `exempt-paths`-Frage — entschieden durch Messung

[slice-172](../done/wellenlos/slice-172-baseline-v600-entfernen.md) hielt fest,
die Ziel-Form habe hier eine Lücke: ihr `exempt-paths` nimmt
`harness/conventions/done/**` aus, was voraussetze, dass alte Stände
liegenbleiben — wer lösche, breche die Ausnahme. **Das war falsch, und die
Messung zeigt warum.**

Muster ohne `exempt-paths` eingetragen, `make doc-check` gefahren: **18
`version-stale`-Befunde**, verteilt auf genau drei Verzeichnisse —

| Klasse | Befunde | Was dort steht |
|---|---|---|
| `docs/plan/planning/done/` | 8 | geschlossene Slices, Welle-Ergebnisnotiz |
| `docs/reviews/` | 6 | Review-Reports |
| `harness/conventions/done/` | 4 | aufgelöste Adaptionen |

**Kein einziger Befund in einer lebenden Datei.** Beanstandet werden `v3.5.2`,
`v5.12.0`, `v6.0.0`, `v6.1.0` — durchweg Stände, die dort **wahr** zitiert
sind, weil damals gegen sie gemessen wurde.

Damit ist die Arbeitsteilung sichtbar, die slice-172 übersehen hat:

- **`versions`** hält *lebende* Dokumente auf dem adoptierten Stand.
  Zeitdokumente sind ausgenommen — sie nennen, was galt.
- **Die Link-Prüfung** sorgt dafür, dass *jeder* Pfad auflöst, auch in
  Zeitdokumenten. Verschwindet ein Stand, meldet **sie** — genau das ist bei
  slice-172 passiert (22 × `target-missing`).

Der Bump von [`MR-018`](../../../../harness/conventions.md#mr-018) war also
richtig, aber nicht aus Versions-Gründen: sein Link hätte ins Leere gezeigt.
Die Ziel-Form ist konsistent; slice-172 hat die Ursache falsch benannt.
**Entscheidung:** `exempt-paths` deckt alle drei Zeitdokument-Klassen, nicht
nur die eine der Vorlage. Gemessene Wirkung: 18 → **0** Befunde.

**Die Prüfmenge ist gemessen, nicht gerechnet** — Korrektur aus dem Review
(F-1). Zuerst stand hier „15 Dateien", per Hand aus *35 minus drei Klassen*
abgeleitet: eine Schätzung im Gewand einer Messung, ausgerechnet in einem
Slice, der sein Verfahren zur Ablesung erklärt. Verfahren jetzt: **alle**
`v6.2.0`-Pins auf einen Phantom-Stand setzen, `make doc-check` fahren, die
meldenden Dateien zählen, zurücksetzen. Ergebnis **14** —
[`AGENTS.md`](../../../../AGENTS.md),
[`harness/README.md`](../../../../harness/README.md),
[`.harness/skills/reviewer.md`](../../../../.harness/skills/reviewer.md),
[`docs/plan/carveouts/README.md`](../../carveouts/README.md),
[`planning/README.md`](../README.md), dieser Plan, die drei Symlinks unter
`.claude/rules/` und die **fünf aktiven `MR`-Einträge, die einen Pin tragen**.
Aktiv sind **sieben**; [`MR-019`](../../../../harness/conventions.md#mr-019) und [`MR-020`](../../../../harness/conventions.md#mr-020) führen keinen — auch das stand
vorher falsch da.

**Zwei lebende Zeiger fallen still unter die Ausnahmen** (Review, F-5):
[`docs/reviews/README.md`](../../../reviews/README.md) ist ein
**Konventions**-Dokument, kein Report, und zeigt auf die *aktuelle*
Report-Vorlage; ebenso der Vorlagen-Link in
[`MR-018`](../../../../harness/conventions.md#mr-018), den slice-172
nachweislich nachgezogen hat. Beide werden vom Muster nicht mehr gesehen, weil
der Glob ihr Verzeichnis nimmt, nicht ihre Rolle. Hingenommen statt behoben:
eine Regel je Datei-Rolle statt je Verzeichnis wäre die saubere Form, kostet
aber mehr Konfigurations-Fläche, als die zwei Zeiger wert sind. **Benannt,
nicht verschwiegen** — und wenn ein dritter dazukommt, ist es die Klasse und
nicht mehr der Einzelfall.

**Zwei Grenzen, die erst die Messung zeigte:**

1. **Die `current-from`-Trägerdatei prüft sich nicht selbst.** Ein Phantom-Pin
   in [`conventions.md`](../../../../harness/conventions.md) erzeugt nur
   `target-missing`, kein `version-stale` — `d-check` nimmt die Datei aus, aus
   der es den Erwartungswert liest. Betroffen sind dort **10** Pins, darunter
   alle `Ersetzt-Baseline-Regel`-Zeiger.
2. **Gedeckt sind sie trotzdem — durch einen Zufall.** `.claude/rules/conventions.md`
   ist ein Symlink auf dieselbe Datei, und unter *diesem* Pfad prüft `d-check`
   den Inhalt sehr wohl. Das ist ein Nebeneffekt der Kontext-Symlinks, keine
   Zusage: verschwindet der Symlink, verschwindet die Deckung, und kein Sensor
   sagt es.

**Zwei weitere Klassen fand der Sensor selbst** — beim ersten Lauf gegen die
Belege dieses Slice. Das Beobachtungs-Register trägt drei Datei-Rollen mit
**drei Lebensdauern** (Baseline `modul-06` §Das Beobachtungs-Register):
`observation.md` ist *einmal geschrieben*, `evidence/<vorgang>.md` ist
*unveränderlich ab Merge* — beide also Zeitdokumente —, `state.md` dagegen ist
der *veränderliche* Stand. Ausgenommen werden darum genau die ersten zwei;
`state.md` bleibt geprüft.

**Halb gemessen, halb argumentiert** — Präzisierung aus dem Review (F-6): Einen
Befund erzeugt hat nur die `evidence/**`-Klasse (die Beleg-Datei dieses Slice,
die eine Mutation zitiert). Die `observation.md`-Klasse ist **abgeleitet**:
beide betroffenen Dateien pinnen heute `v6.2.0` und melden nichts; ausgenommen
sind sie, weil die Baseline sie als *einmal geschrieben* führt. Und
„`state.md` bleibt geprüft" hat derzeit **null** Gegenstände — keine trägt
einen Pin. Beides ist Vorsorge, und sie ist als solche zu lesen.

### 2.2 Was der Sensor nicht sieht — und was daraus folgt

Ein Symlink hat für `d-check` keinen Linkpfad: es liest den **Zielinhalt**.
Gemessen: `.claude/rules/modul-05-planning-harness.md` auf einen entfernten
Baseline-Stand umgebogen ⇒ `make doc-check` meldet **0 Befunde**, obwohl der
Symlink ins Leere zeigt.

Das ist die **dritte** Instanz von
[`BEO-GATE/symlink-ziel-nach-baseline-bump-ungeprueft`](../observations/BEO-GATE/symlink-ziel-nach-baseline-bump-ungeprueft/observation.md)
(slice-142 · slice-167 · hier) und damit keine Notiz mehr, sondern eine Lücke,
die einen Sensor verlangt. Sie fällt in dieselbe Schicht und denselben
Gegenstand, also in diesen Slice — dritter Liefer-Punkt, eine Schicht, die
Größen-Regel hält.

**Die drei Instanzen haben zwei verschiedene Formen** — Fund des Reviews (F-2),
und er hat den Sensor verändert. Die erste Fassung prüfte nur *„Ziel existiert
nicht"*. Damit hätte sie **slice-167 grün gemeldet**: dort zeigten die Symlinks
auf `v6.0.0`, das während der Migration daneben vendored blieb — sie lösten auf
und waren trotzdem falsch. Ein Sensor, der nur die eine Form fängt, hätte die
Beobachtung als verkörpert ausgewiesen und eine der gezählten Instanzen
durchgelassen. Der Sensor prüft darum **zwei** Dinge:

| Prüfung | Form | Instanz |
|---|---|---|
| Ziel existiert | der Stand wurde entfernt | slice-173 (hier gemessen) |
| Baseline-Ziel trägt den adoptierten Stand | der alte Stand liegt noch daneben | slice-167 |

Den adoptierten Stand liest das Skript aus derselben Prosa-Zeile wie
`versions.current-from` — fail-closed, ohne lesbaren Stand bricht es ab statt
zu raten. Diese Kopplung ist dieselbe, die §7 als Risiko führt.

**Ein zweiter Review-Fund steckte in der Extraktion** (F-3): `git ls-files -s`
plus awk-Feldzuweisung kollabierte mehrfache Leerzeichen im Dateinamen
(`a  b.md` → `a b.md`) und stolperte über C-quotierte Nicht-ASCII-Pfade. Der
Defekt lag **vor** der geteilten Prüf-Funktion, also genau außerhalb des
Codepfads, den die Design-Begründung als geteilt auswies — der Selbsttest
konnte ihn nicht sehen. Ersetzt durch `git ls-files -z` (rohe Pfade,
NUL-getrennt); der Selbsttest führt seither einen Namen mit zwei Leerzeichen
mit.

## 3. Umsetzung

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| [`.d-check.yml`](../../../../.d-check.yml) | update | zweites `versions`-Muster samt `exempt-paths` für **fünf** Zeitdokument-Klassen (§2.1) |
| `tools/symlink-check.sh` | neu | Sensor für §2.2; **eine** Prüf-Funktion für Gate-Lauf und Selbsttest, damit der Test den Codepfad prüft, der im Gate läuft |
| [`Makefile`](../../../../Makefile) | update | Target `symlink-check`, im `gates`-Aggregat, in `.PHONY` |
| `.claude/hooks/pretooluse-command-guard.sh` | update | `symlink-check` in die `GATES`-Liste — sonst liefe `make symlink-check \| tail` ungehindert durch; sein Selbsttest fängt das Fehlen |
| [`AGENTS.md`](../../../../AGENTS.md) §4, [`harness/README.md`](../../../../harness/README.md) | update | Deklaration beider Änderungen; `doc-targets` erzwingt sie mechanisch |

## 4. Definition of Done

- [x] [`.d-check.yml`](../../../../.d-check.yml) trägt ein zweites
      `versions`-Muster für `.harness/baseline/<tag>/`; ein künstlich
      veralteter Pin erzeugt nachweislich einen `version-stale`-Befund, ein
      korrekter nicht — beide Richtungen gemessen (`AGENTS.md:31`
      `v6.1.0 version-stale` bzw. 0 Befunde nach Rücknahme).
- [x] Die `exempt-paths`-Frage aus §2.1 ist entschieden — **fünf**
      Zeitdokument-Klassen statt der einen der Vorlage —, die Begründung ist
      gemessen und steht in [`AGENTS.md`](../../../../AGENTS.md) §4, wo sie
      beim nächsten Baseline-Sprung gelesen wird.
- [x] `make symlink-check` existiert, hängt im `gates`-Aggregat und feuert
      nachweislich in **beiden** Formen der Beobachtung (§2.2): Ziel fehlt,
      und Baseline-Ziel trägt einen anderen als den adoptierten Stand. Beide
      Prüfungen sind mutations-kalibriert — je eine Mutation im Skript lässt
      den Selbsttest fehlschlagen.
- [x] Unabhängiger Review durchgeführt (Report unter [`docs/reviews/`](../../../reviews/README.md)).
- [x] `make gates` grün.
- [x] `make verify` grün.
- [x] Jedes Risiko trägt einen Ausgang.

## 5. Trigger

**Start** (`open` → `in-progress`): Maintainer-Freigabe und WIP-Limit frei.

**Rückführungen:** stellt sich heraus, dass das Muster den Bestand
massenhaft bricht statt einzelne Nachzüge zu melden — zurück nach `next/`;
eine Regel, die den Bestand flächig rot färbt, wird abgeschaltet statt
befolgt ([`AGENTS.md`](../../../../AGENTS.md) §5, Begründung zum
Commit-Scope).

## 6. Closure-Trigger

DoD vollständig, `make gates`/`make verify` grün, Closure-Notiz
geschrieben. Danach Archivierung als wellenloser Slice
([`AGENTS.md`](../../../../AGENTS.md) §6).

## 7. Risiken und offene Punkte

- *Das Muster trifft auch Prosa-Erwähnungen eines Pfades, die bewusst einen
  historischen Stand nennen — dann meldet der Sensor einen Nachzug, der
  falsch wäre* — **Ausgang:** eingetreten, und zwar **18×** beim ersten Lauf
  (§2.1). Aufgefangen durch `exempt-paths` auf die drei Zeitdokument-Klassen;
  kein Folge-Slice nötig, weil die Auffangregel Teil dieses Slice ist. Was
  bleibt: eine historische Nennung **außerhalb** dieser drei Klassen würde
  weiterhin fälschlich gemeldet — dann ist die Klasse zu erweitern, nicht die
  Prosa umzuschreiben.
- *`current-from` liest den Stand aus einer Prosa-Zeile; ändert deren
  Wortlaut, bricht der Sensor an einer Stelle, die niemand mit ihm in
  Verbindung bringt* — **Ausgang:** weiter offen → Beobachtungs-Register,
  [`BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf`](../observations/BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf/observation.md).
  Es ist dieselbe Klasse in ihrer **Korpus**-Ausprägung: eine
  Sensor-Konfiguration hängt an einem Wortlaut in einer *anderen* Datei, und
  ob der noch trägt, sagt kein Lauf. `make dcheck-phrase-selftest` deckt die
  Werkzeug-Seite für zwei Muster ab; `current-from` ist ein drittes und steht
  nicht darin. Der Eintrag ist bereits auf *geplant* →
  [slice-169](../open/slice-169-korpus-seitige-kalibrierung.md), und **dieser
  Slice erweitert seinen Gegenstand**: Beleg `evidence/slice-173.md`. Kein
  vierter Liefer-Punkt hier — das wäre über der Größen-Regel.

## 8. Closure-Notiz

**Lerneintrag — Form: neuer Sensor** (`make symlink-check`), dazu eine
**Korrektur** an einer Aussage des Vorgänger-Slice.

- **Was hat funktioniert:** Erst konfigurieren, dann entscheiden. Die
  `exempt-paths`-Frage stand seit dem Anlegen des Slice als offene
  Grundsatzfrage im Plan; beantwortet hat sie nicht ein Argument, sondern ein
  Lauf — 18 Befunde, verteilt auf drei Verzeichnisse, **null** in lebenden
  Dateien. Danach war die Entscheidung keine Abwägung mehr, sondern eine
  Ablesung.

- **Was ging anders als geplant:** Der Slice sollte eine Lücke der Baseline
  schließen, die
  [slice-172](../done/wellenlos/slice-172-baseline-v600-entfernen.md) benannt
  hatte. Die **Zuschreibung** war falsch: die Ziel-Form nimmt Zeitdokumente
  aus, weil dort der *damalige* Stand wahr ist — nicht, weil sie das
  Fortbestehen alter Stände voraussetzt. Dass Pfade auch in Zeitdokumenten
  auflösen müssen, trägt die **Link**-Prüfung, und genau die hat bei slice-172
  gemeldet (22 × `target-missing`). Zwei Sensoren, zwei Fragen.

  **Die Widerlegung trägt aber nur die halbe Strecke** — Korrektur aus dem
  Review (F-7), und zuerst stand hier das zu bequeme *„es gab keine"*. Was
  slice-172 **beobachtet** hatte, bleibt wahr: verschwindet ein vendorter
  Stand, erzwingt die Link-Prüfung einen Edit an Dateien, die das Repo als
  unveränderlich führt — bei
  [`MR-018`](../../../../harness/conventions.md#mr-018) genau so geschehen.
  Falsch war nur die Ursache: nicht `exempt-paths` setzt liegenbleibende
  Stände voraus, sondern **Link-Auflösbarkeit und Unveränderlichkeit
  kollidieren**, sobald ein Stand geht. Das ist eine Klemme, keine Lücke, und
  sie steht jetzt als solche in
  [`conventions.md`](../../../../harness/conventions.md#baseline) §Baseline —
  zusammen mit der Unterscheidung *aktiver* Eintrag (Zeiger wandert mit,
  maschinell durchgesetzt) gegen *aufgelöster* (eingefroren, vom Sensor
  ausgenommen), die slice-172 dort unausgesprochen gelassen hatte.

- **Steering-Loop-Eintrag — neuer Sensor:** `make symlink-check` — jeder
  getrackte Symlink löst auf — liegt in `Makefile:symlink-check`
  (`tools/symlink-check.sh`, im `gates`-Aggregat).
  Auslöser: [`BEO-GATE/symlink-ziel-nach-baseline-bump-ungeprueft`](../observations/BEO-GATE/symlink-ziel-nach-baseline-bump-ungeprueft/observation.md)
  (slice-142, slice-167, slice-173 — 3×). Die dritte Instanz ist die erste
  **gemessene**: ein umgebogener Symlink ließ `doc-check` bei 0 Befunden.

- **Beobachtungs-Register (`../observations/`):** `evidence/slice-173.md` in
  [`BEO-GATE/symlink-ziel-nach-baseline-bump-ungeprueft`](../observations/BEO-GATE/symlink-ziel-nach-baseline-bump-ungeprueft/observation.md)
  ergänzt — Zähler 3×, Stand jetzt *verkörpert*; und in
  [`BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf`](../observations/BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf/observation.md)
  ergänzt — Zähler 4×, Stand bleibt *geplant* → slice-169, dessen Gegenstand
  damit um ein drittes Muster wächst.

- **Folge-Slices:** keine neuen. [slice-169](../open/slice-169-korpus-seitige-kalibrierung.md)
  bestand bereits und ist eine Datei in `open/`.

- **Risiken aus §7:** zwei, jedes mit genau einem Ausgang — einmal
  *eingetreten* (18×, im selben Slice aufgefangen), einmal *weiter offen* →
  Beobachtungs-Register. Siehe §7.

- **Drei Paarungen** (Repo ohne Wellen-Betrieb, nach dem `git mv` geprüft):
  **Anker** — `liegt in Makefile:symlink-check` vergeben, Ziel existiert und
  trägt `seit slice-173` im `state.md` der auslösenden Beobachtung.
  **Folge-Slice** — keiner neu; `slice-169` liegt in `open/`.
  **Register** — beide zitierten Beobachtungen existieren als Verzeichnis mit
  nicht leerem `evidence/`; `make verify` prüft die maschinelle Hälfte.

## 9. Sub-Area-Modus

**Vorgelagert — Sub-Area-Wahl prüfen:** eine Sub-Area berührt —
**Gate-/Werkzeug-Schicht** (`.d-check.yml`, `Makefile`), Achsen 1,2,3,
deklariert in
[`conventions.md`](../../../../harness/conventions.md#modus-deklaration-pro-sub-area).

**Vorgelagert — offene Beobachtungen sichten:** Register am 2026-09-07
durchgegangen, **17** Einträge in `BEO-GATE`, davon 11 offen. Drei einschlägig:

| Eintrag | Stand vorher | Bezug zu diesem Slice |
|---|---|---|
| [`symlink-ziel-nach-baseline-bump-ungeprueft`](../observations/BEO-GATE/symlink-ziel-nach-baseline-bump-ungeprueft/observation.md) | offen, 2× | **3× erreicht und verkörpert** — die Lücke wurde hier erstmals gemessen (§2.2) statt nur beobachtet; Antwort `make symlink-check` |
| [`pruefer-ohne-gegenstand-oder-aufruf`](../observations/BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf/observation.md) | geplant → slice-169, 3× | **erweitert, nicht neu verplant** — `current-from` ist ein drittes phrasen-basiertes Muster, das `dcheck-phrase-selftest` nicht deckt; Beleg `evidence/slice-173.md` |
| [`versionsangabe-neben-digest-ungeprueft`](../observations/BEO-GATE/versionsangabe-neben-digest-ungeprueft/observation.md) | verkörpert | **berührt, nicht erhöht** — `make version-coherence` prüft *doppelt deklarierte* Angaben auf Divergenz; dieser Slice prüft *einfach* deklarierte gegen **einen** erklärten Stand. Andere Frage, benachbarte Antwort |

**Alle berührten Sub-Areas GF** — kein Begründungsblock nötig.
