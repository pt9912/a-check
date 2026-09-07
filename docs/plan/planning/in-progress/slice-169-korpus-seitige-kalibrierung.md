# slice-169 — Korpus-seitige Kalibrierung phrasen-basierter Prüfer

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `make slice-mv` ([`AGENTS.md`](../../../../AGENTS.md)
§3.3/§5).

**Welle:** ohne Welle (Trigger: Review-Findings F-1/F-2 aus
[`slice-168`](../done/wellenlos/slice-168-pruefer-kalibrierungs-selbsttest.md) —
kein Mehr über die eigene DoD hinaus).

**Bezug:** [`BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf`](../observations/BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf/observation.md)
— die Korpus-Hälfte, die `slice-168` nicht gedeckt hat.

**Berührte Spec-Stellen:** — *(keine)* — Harness-/Werkzeug-Änderung ohne
Vertragsberührung.

**Verantwortlich:** — *(noch nicht priorisiert)*

**Autor:** Claude (Sonnet 5), im Auftrag des Maintainers. **Datum:**
2026-09-06.

---

## 1. Ziel

`slice-168` kalibriert die **Werkzeug-Seite** phrasen-basierter Prüfer:
reagiert `d-check` noch auf die gewählte Phrase? Ausgefallen ist zweimal
die **Korpus-Seite**: a-checks eigene Dokumente trafen das Muster nicht
mehr (slice-120, slice-165). Dieser Slice deckt sie nach.

## 2. Analyse (vor der Umsetzung)

Zwei Befunde des unabhängigen Reviews zu `slice-168` (Report mit dem Slice
archiviert:
`unzip -p done/wellenlos/slice-168-archiv.zip docs/reviews/2026-09-06-slice-168-pruefer-kalibrierungs-selbsttest.md`):

- **F-1 (HIGH):** Der Selbsttest hardcodet die richtige Fixture-Zeile und
  liest den echten Korpus nie an. Gemessen: die Kandidatenmenge des
  `reviews`-Moduls ist heute **nicht leer** (vier Slices), aber drei
  geschlossene Slices mit vorhandenem Review-Report tragen eine
  nicht-auslösende Wortform — die Menge kann jederzeit wieder leer werden,
  ohne dass ein Sensor es sagt.
- **F-2 (MEDIUM):** `TASKS_IGNORE_PATTERN` steht im Selbsttest als Kopie
  neben dem Original in `.d-check.yml`. Gemessen: nach einem Bruch der
  **echten** Konfiguration blieb der Selbsttest bei Exit 0.

**Die drei Entscheidungen — gemessen, nicht abgewogen:**

**1. Nichtleerheit, keine Erwartungszahl.** Die belegte Ausfallart ist *„die
Menge wird leer"* (slice-120, slice-165). Eine feste Zahl bräche zusätzlich bei
**jedem** neuen Slice — und eine Regel, die den Bestand massenhaft bricht, wird
abgeschaltet statt befolgt ([`AGENTS.md`](../../../../AGENTS.md) §5, Begründung
zum Commit-Scope). `exempt-expect-count` ist die Gegenprobe dazu: Es zählt eine
Menge, die sich **nicht** vergrößern soll (19 grandfatherte `AC-*`), und genau
darum trägt dort eine Zahl.

**2. `sed`-Auszug aus `.d-check.yml`, fail-closed — kein `yq`, keine Kopie.**
Drei Wege geprüft:

| Weg | Befund |
|---|---|
| `d-check --json` | nennt nur `filesChecked` und `findingCount` — **nicht** die Kandidatenmenge je Modul |
| `d-check --print-config` | gibt das **Startgerüst** aus, nicht die aktive Konfiguration des Repos |
| `yq` | neue Abhängigkeit im hermetischen Lauf |

Bleibt der Auszug per `sed`. Seine Fragilität gegenüber YAML-Formatierung ist
gefangen, weil er **fail-closed** ist: Findet er das Feld nicht, bricht der Lauf
ab, statt mit leerem Muster durchzulaufen. Mutations-belegt — ein umbenanntes
Feld ergibt Exit 2 mit der Meldung *„nicht lesbar"*.

**3. Geltungsbereich: zwei von vierzehn — die mit belegtem Ausfall.**
`.d-check.yml` führt **14** phrasen-basierte Felder. Gedeckt sind die zwei, deren
Ausfall dokumentiert ist: die `reviews`-Trigger-Phrase (zweimal ausgefallen) und
`structure`s `tasks-ignore-pattern` (Review-Befund F-2 zu slice-168). Für die
übrigen zwölf gibt es keinen Vorfall — **ein Sensor ohne Anlass ist selbst eine
Behauptung**, und die Prüfmenge zu verdoppeln, ohne zu wissen, wofür, ist genau
die Bewegung, die dieser Beobachtungs-Eintrag beschreibt.

*(`versions.current-from` wäre der nächste Kandidat — slice-179 hat ihn als
drittes Muster belegt. Er bleibt draußen, weil er bei Bruch **laut** ausfällt:
`d-check` bricht bei nicht auflösbarem Anker ab, statt grün zu melden. Die
Gefährlichkeit der leeren Prüfmenge fehlt ihm.)*

## 3. Umsetzung

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| `tools/dcheck-phrase-selftest.sh` | update | zwei **Korpus**-Kontrollen neben die vier Werkzeug-Kontrollen; Muster per fail-closed-Auszug aus `.d-check.yml` statt als Kopie |
| [`AGENTS.md`](../../../../AGENTS.md) §4, [`harness/README.md`](../../../../harness/README.md) | update | der Vertrag nennt jetzt beide Hälften |

**Zwei Kontrollen, nicht mehr:** je Muster eine Nichtleerheits-Prüfung gegen den
`done/`-Bestand. Beide nennen ihre Zahl in der Erfolgs-Zeile (heute 5 Slices und
14 DoD-Zeilen) — eine eingefrorene Zahl stünde falsch, sobald jemand committet,
die **gezählte** sagt beim Lesen, worüber das Grün eine Aussage macht.

**Mutations-belegt in beide Richtungen:**

| Mutation | Erwartet | Gemessen |
|---|---|---|
| `tasks-ignore-pattern` auf ein Muster ohne Treffer | rot | Exit 2, *„trifft im done/-Bestand NICHTS"* |
| Feld umbenannt (unlesbar) | rot, **fail-closed** | Exit 2, *„nicht lesbar … wird nicht geprueft, sondern abgebrochen"* |
| unverändert | grün | Exit 0, beide Mengen genannt |

Die zweite Zeile ist der Beleg für den Review-Befund F-2 zu slice-168: Dort
blieb der Selbsttest nach einem Bruch der **echten** Konfiguration bei Exit 0,
weil er eine Kopie las.

## 4. Definition of Done

- [x] Analyse-Fragen aus §2 beantwortet und belegt.
- [x] Korpus-seitige Kontrolle der Kandidatenmenge umgesetzt und gegen eine
      absichtlich leere Menge verifiziert.
- [x] Muster-Kopie durch eine Kopplung an `.d-check.yml` ersetzt, gegen
      einen Bruch der echten Konfiguration verifiziert.
- [x] Unabhängiger Review durchgeführt (Report unter `docs/reviews/`).
- [x] `make gates` grün.
- [x] `make verify` grün.
- [x] Jedes Risiko trägt einen Ausgang.

## 5. Trigger

**Start** (`open` → `in-progress`): Maintainer-Freigabe und WIP-Limit frei.

**Rückführungen:** wächst der Geltungsbereich über die beiden Befunde
hinaus (Analyse-Frage 3), zurück nach `next/` zur Zerlegung.

## 6. Closure-Trigger

DoD vollständig, `make gates`/`make verify` grün, Closure-Notiz
geschrieben.

## 7. Risiken und offene Punkte

- *Eine Erwartungszahl über den eigenen Korpus altert mit jedem neuen
  Slice und wird dann routinemäßig hochgezählt statt geprüft* — **Ausgang:**
  gestrichen mit Begründung. Das Risiko war der Grund, **keine** Zahl zu nehmen:
  Der Sensor prüft Nichtleerheit (§2.1). Die Zahl erscheint nur in der
  Erfolgs-Zeile, wo sie niemand pflegt — sie wird bei jedem Lauf neu gezählt.
- *Die Kopplung an `.d-check.yml` braucht möglicherweise ein YAML-Werkzeug
  im hermetischen Lauf, das heute nicht da ist* — **Ausgang:** gestrichen mit
  Begründung. Drei Wege gemessen (§2.2); der gewählte `sed`-Auszug braucht
  keines. Was bleibt, ist seine Fragilität gegenüber YAML-Umformatierung — und
  die ist kein offener Punkt, sondern eine benannte **Sperre**: Der Auszug ist
  fail-closed und meldet den Bruch, statt ihn zu überspielen.
- *Zwölf der vierzehn phrasen-basierten Felder bleiben ungedeckt, und der
  nächste Ausfall trifft eines davon* — **Ausgang:** weiter offen →
  Beobachtungs-Register,
  [`BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf`](../observations/BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf/observation.md).
  Der Eintrag bleibt bestehen, weil die Klasse breiter ist als die zwei
  gedeckten Muster; sein Ausgang wechselt mit diesem Slice von *geplant* auf
  *verkörpert*, aber die Grenze ist benannt statt geschlossen.

## 8. Closure-Notiz

**Lerneintrag — Form: neuer Sensor** (Korpus-Hälfte der phrasen-basierten
Kalibrierung).

- **Was hat funktioniert:** Die drei offenen Entscheidungen aus §2 **gemessen**
  statt abgewogen. Bei Entscheidung 2 hieß das, drei Wege auszuprobieren, bevor
  einer gewählt wurde — `--json` nennt die Kandidatenmenge nicht,
  `--print-config` gibt das Startgerüst statt der aktiven Konfiguration, `yq`
  wäre eine neue Abhängigkeit. Ohne die Messung wäre der fragile `sed`-Auszug
  eine Notlösung gewesen; mit ihr ist er der einzige Weg, der ohne neue
  Abhängigkeit trägt.

- **Was ging anders als geplant:** Der Plan nannte **vier** weitere
  phrasen-basierte Konfigurationen als möglichen Geltungsbereich. Gezählt sind
  es **vierzehn** — und die Antwort darauf ist nicht „alle decken", sondern die
  Gegenfrage: Für zwölf davon gibt es **keinen belegten Ausfall**. Ein Sensor
  ohne Anlass ist selbst eine Behauptung, und die Prüfmenge zu verdoppeln, ohne
  zu wissen wofür, wäre genau die Bewegung, die dieser Beobachtungs-Eintrag
  beschreibt.

- **Steering-Loop-Eintrag — neuer Sensor:** Die Korpus-Hälfte des
  Kalibrierungs-Selbsttests — je Muster eine **Nichtleerheits**-Prüfung gegen
  den `done/`-Bestand, mit dem Muster aus `.d-check.yml` statt aus einer Kopie.
  — liegt in `Makefile:dcheck-phrase-selftest`
  (`tools/dcheck-phrase-selftest.sh`, im `gates`-Aggregat).
  Auslöser: [`BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf`](../observations/BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf/observation.md)
  (slice-120, slice-123, slice-165, slice-173 — 4×), Ausgang war seit slice-168
  *geplant* auf diesen Slice.

- **Beobachtungs-Register (`../observations/`):** kein neuer Eintrag, kein
  neuer Beleg. Der auslösende Eintrag wechselt von *geplant* auf **verkörpert**
  — mit einer benannten Grenze: zwölf der vierzehn Muster bleiben ungedeckt,
  und der Eintrag bleibt darum stehen statt gestrichen zu werden.

- **Folge-Slices:** keine. Die zwölf ungedeckten Muster bekommen einen, wenn
  eines davon ausfällt — nicht vorher.

- **Risiken aus §7:** drei, jedes mit genau einem Ausgang — zweimal *gestrichen
  mit Begründung* (beide Risiken waren der Grund für die getroffene Wahl),
  einmal *weiter offen* → Register.

- **Drei Paarungen** (Repo ohne Wellen-Betrieb, nach dem `git mv` geprüft):
  **Anker** — `liegt in Makefile:dcheck-phrase-selftest`, Ziel existiert.
  **Folge-Slice** — keiner genannt.
  **Register** — der zitierte Eintrag existiert mit nicht leerem `evidence/`.

## 9. Sub-Area-Modus

**Vorgelagert — Sub-Area-Wahl prüfen:** eine Sub-Area berührt —
**Gate-/Werkzeug-Schicht** (`tools/`, `Makefile`, `.d-check.yml`),
Greenfield, Schwelle ≥ 2/3 erfüllt.

**Vorgelagert — offene Beobachtungen sichten:** entsteht mit dem Übergang
nach `in-progress/` (der Register-Stand ist beim Anlegen ein anderer als
beim Beginn der Arbeit).

**Alle berührten Sub-Areas GF** — kein Begründungsblock nötig.
