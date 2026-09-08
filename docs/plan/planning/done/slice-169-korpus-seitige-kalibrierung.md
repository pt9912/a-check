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

**Verantwortlich:** Claude — gesetzt beim Übergang nach `in-progress/`.

**Autor:** Claude (Sonnet 5), im Auftrag des Maintainers. **Datum:**
2026-09-06.

---

## 1. Ziel und Abgrenzung

**Ziel:** `slice-168` kalibriert die **Werkzeug-Seite** phrasen-basierter
Prüfer: reagiert `d-check` noch auf die gewählte Phrase? Ausgefallen ist zweimal
die **Korpus-Seite**: a-checks eigene Dokumente trafen das Muster nicht
mehr (slice-120, slice-165). Dieser Slice deckt sie nach.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **`structure.tasks-ignore-pattern` als Korpus-Kontrolle.** Es fällt **laut**
  aus: trifft es nichts mehr, zählen die konstanten DoD-Posten mit und
  `doc-structure` meldet `section-oversized`. Der Eintrag im Register
  beschreibt die **stille** Klasse — grün ohne Prüfgegenstand. Derselbe Maßstab
  schließt `versions.current-from` aus (`d-check` bricht bei nicht auflösbarem
  Anker ab).
- **Die übrigen phrasen-basierten Felder in `.d-check.yml`.** Für sie gibt es
  keinen belegten Ausfall, und ein Sensor ohne Anlass ist selbst eine
  Behauptung. Fällt eines aus, ist das ein neuer Beleg am bestehenden
  Register-Eintrag — kein neuer Sensor.
- **Eine Erwartungszahl statt Nichtleerheit.** Die belegte Ausfallart ist „die
  Menge wird leer"; eine feste Zahl bräche bei jedem neuen Slice, und eine
  Regel, die den Bestand massenhaft bricht, wird abgeschaltet statt befolgt.
- **Produkt-Code.** Der Slice rührt `internal/` nicht an — er arbeitet
  ausschließlich in der Gate-/Werkzeug-Schicht.

*(Der Abschnitt trug bis zur Review-Einarbeitung nur das Ziel. Die Ausschlüsse
galten faktisch — sie standen als Argumente in §2.3 und §6 — und sind hier
nachgezogen; der erste ist mit der Einarbeitung von F-6 zusätzlich
**enger** geworden: die zweite Korpus-Kontrolle ist entfallen, nicht nur
begrenzt.)*

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

**3. Geltungsbereich: die eine Konfiguration mit belegtem *stillem* Ausfall.**
Korpus-seitig gedeckt ist die `reviews`-Trigger-Phrase — zweimal ausgefallen
(slice-120, slice-165), und ihr Ausfall ist **still**: ohne Kandidaten meldet
`doc-reviews` grün, ohne etwas zu prüfen. Für die übrigen phrasen-basierten
Felder gibt es keinen solchen Vorfall — **ein Sensor ohne Anlass ist selbst eine
Behauptung**, und die Prüfmenge zu verbreitern, ohne zu wissen wofür, ist genau
die Bewegung, die dieser Beobachtungs-Eintrag beschreibt.

**Keine Gesamtzahl.** Eine frühere Fassung nannte hier „vierzehn Felder"; die
Zahl hängt an einer Zählregel für *phrasen-basiert*, die nirgends festgelegt ist
— je nach Auslegung zählt `.d-check.yml` deutlich mehr regex-tragende Einträge.
Und die `reviews`-Phrase ist gar kein Feld dort: sie lebt im gepinnten Werkzeug
und steht in der Konfiguration nur im Kommentar (Review slice-169, F-7).

*Zwei Kandidaten bleiben ausdrücklich draußen, beide weil sie **laut**
ausfallen:* `versions.current-from` (Beleg `evidence/slice-173.md` im Register)
— `d-check` bricht bei nicht auflösbarem Anker ab, statt grün zu melden. Und
`structure`s `tasks-ignore-pattern`: trifft es nichts, zählen die konstanten
DoD-Posten mit und `doc-structure` meldet `section-oversized`. Beiden fehlt die
Gefährlichkeit der leeren Prüfmenge. **Der zweite stand bis zur
Review-Einarbeitung als Korpus-Kontrolle drin** — mit demselben Maßstab
gemessen, mit dem der erste ausgeschlossen wurde, gehört er nicht hinein
(Review slice-169, F-6).

## 3. Umsetzung

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| `tools/dcheck-phrase-selftest.sh` | update | **eine** Korpus-Kontrolle neben die vier Werkzeug-Kontrollen; Kandidatenverzeichnis per fail-closed-Auszug aus `.d-check.yml` statt als Kopie |
| [`AGENTS.md`](../../../../AGENTS.md) §4, [`harness/README.md`](../../../../harness/README.md), `harness/sensors/dcheck-phrase-selftest.md` | update | der Vertrag nennt beide Hälften und den Umfang der zweiten |
| `observations/BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf/state.md` | update | Ausgang von *geplant* auf *verkörpert*, mit benannter Grenze |

**Eine Kontrolle, nicht zwei.** Der Plan sah je Muster eine
Nichtleerheits-Prüfung vor; `structure`s `tasks-ignore-pattern` ist mit der
Review-Einarbeitung entfallen, weil es **laut** ausfällt (§1, Review F-6).

**Gezählt wird die Menge des Moduls.** Das Modul `reviews` sieht einen
**DoD-Haken** in einem **flachen** `done/`-Slice — eine Nennung in Prosa, in
einer Tabelle oder in einer Wellen-Ergebnisnotiz sieht es nicht. Die erste
Fassung zählte per `grep -rli` jede Datei unter `done/`, die die Phrase
irgendwo trug: **5** statt **2**, und der Lauf blieb grün, nachdem beide echten
Kandidaten entwertet waren (Review F-1). Die Zahl steht nur in der
Erfolgs-Zeile, wo sie bei jedem Lauf neu gezählt wird.

**Mutations-belegt in beide Richtungen**, jede Probe auf einer Kopie:

| Mutation | Erwartet | Gemessen |
|---|---|---|
| die zwei DoD-Haken entwertet, Prosa-Nennungen bleiben | rot | Exit 1, *„Kandidatenmenge des reviews-Moduls ist LEER"* |
| dieselbe Mutation gegen die **erste** Fassung | *(sie meldete grün)* | Exit 0, *„3 Slice(s)"* — der Befund, den F-1 belegt |
| `done-dir` umbenannt (unlesbar) | rot, **fail-closed** | Exit 1, *„nicht lesbar oder kein Verzeichnis … abgebrochen"* |
| unverändert | grün | Exit 0, *„2 flache(r) Slice(s)"* |

Die dritte Zeile ist der Beleg für den Review-Befund F-2 zu slice-168: Dort
blieb der Selbsttest nach einem Bruch der **echten** Konfiguration bei Exit 0,
weil er eine Kopie las.

## 4. Definition of Done

- [x] Analyse-Fragen aus §2 beantwortet und belegt.
- [x] Korpus-seitige Kontrolle **der Kandidatenmenge des Moduls** umgesetzt
      und gegen eine absichtlich leere Menge verifiziert — Mutations-Probe auf
      einer Kopie, beide echten Kandidaten entwertet, Lauf rot.
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
- *Die übrigen phrasen-basierten Felder bleiben ungedeckt, und der
  nächste Ausfall trifft eines davon* — **Ausgang:** weiter offen →
  Beobachtungs-Register,
  [`BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf`](../observations/BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf/observation.md).
  Der Eintrag bleibt bestehen, weil die Klasse breiter ist als das eine
  korpus-seitig gedeckte Muster; sein Ausgang wechselt mit diesem Slice von
  *geplant* auf *verkörpert*, aber die Grenze ist benannt statt geschlossen.

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

- **Was ging anders als geplant.** Der unabhängige Review
  ([`2026-09-08-slice-169-korpus-seitige-kalibrierung.md`](../../../reviews/2026-09-08-slice-169-korpus-seitige-kalibrierung.md))
  war **merge-blockierend**: 2 HIGH, 8 MEDIUM, 4 LOW, vierzehn Finding-Klassen.
  Zwei davon änderten die Substanz, und beide fand er, nicht die schreibende
  Session. (1) Der Plan sah **zwei** Korpus-Kontrollen vor, je
  Muster eine. `structure`s `tasks-ignore-pattern` fällt aber **laut** aus:
  trifft es nichts, meldet `doc-structure` `section-oversized`. Mit demselben
  Maßstab, mit dem der Plan `versions.current-from` ausschloss, gehört es nicht
  hinein — die Kontrolle ist entfallen statt repariert worden (F-6). (2) Die
  verbliebene Kontrolle zählte eine **Obermenge** ihres Gegenstands: `grep -rli`
  über `done/` fand 5 Dateien, die Kandidatenmenge des Moduls sind 2. Auf einer
  Kopie gemessen blieb der Lauf grün, nachdem beide echten Kandidaten entwertet
  waren (F-1) — **genau die Ausfallart, gegen die dieser Slice antritt.**
  Ein Sensor, der die falsche Menge zählt, ist von einem funktionierenden erst
  durch die Mutations-Probe zu unterscheiden.

- **Steering-Loop-Eintrag — neuer Sensor:** Die Korpus-Hälfte des
  Kalibrierungs-Selbsttests — eine **Nichtleerheits**-Prüfung gegen die
  Kandidatenmenge des `reviews`-Moduls (DoD-Haken in einem flachen
  `done/`-Slice), mit dem Kandidatenverzeichnis aus `.d-check.yml` statt aus
  einer Kopie. — liegt in `Makefile:dcheck-phrase-selftest`
  (`tools/dcheck-phrase-selftest.sh`, im `gates`-Aggregat).
  Auslöser: [`BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf`](../observations/BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf/observation.md)
  (slice-120, slice-123, slice-165, slice-173, slice-181 — 5×), Ausgang war
  seit slice-168 *geplant* auf diesen Slice.

- **Die Lehre, die über diesen Sensor hinausgeht:** Ein Sensor, der die
  **falsche Menge** zählt, ist von einem funktionierenden nicht durch Lesen zu
  unterscheiden — beide melden grün. Nur die Mutations-Probe trennt sie, und
  sie muss gegen die Menge laufen, über die der Sensor eine Aussage macht,
  nicht gegen die, die er zählt. Der Slice hatte die Probe: sie mutierte das
  Muster und traf damit die Zählmenge des Skripts, nicht die Kandidatenmenge
  des Moduls. **Die Probe war grün und der Sensor blind.** Das ist dieselbe
  Klasse wie der Eintrag, den dieser Slice bedient — nur eine Ebene höher: dort
  ein Prüfer ohne Gegenstand, hier eine Probe ohne Gegenstand.

- **Beobachtungs-Register (`../observations/`):** ein neuer Eintrag,
  [`BEO-GATE/attestierung-vor-dem-vorgang`](../observations/BEO-GATE/attestierung-vor-dem-vorgang/observation.md)
  (Beleg `evidence/slice-169.md`, Zähler damit 1×) — der Review-Haken dieses
  Slice stand auf `[x]`, während kein Report existierte, und `make doc-reviews`
  deckt die Attestierung in `in-progress/` nicht ab.
  Ein zweiter Beleg an
  [`BEO-HARNESS/hard-rule-37-ohne-sensor`](../observations/BEO-HARNESS/hard-rule-37-ohne-sensor/observation.md)
  (Zähler 2×) — der Review fand zwei Kommentare im Konjunktiv über die
  verworfene Alternative und wies nach, dass dieselbe Form schon aus slice-168
  im Bestand steht.
  Der **auslösende** Eintrag
  [`BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf`](../observations/BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf/observation.md)
  bekommt **keinen** neuen Beleg — *ein Vorgang zählt einmal*, und dieser Slice
  ist der Vorgang, der ihn verkörpert, nicht einer, der ihn erneut auslöst.
  Sein Ausgang wechselt von *geplant* auf
  **verkörpert** — mit einer benannten Grenze: gedeckt ist **eine**
  Konfiguration, die mit belegtem *stillem* Ausfall; der Eintrag bleibt darum
  stehen statt gestrichen zu werden.

- **Folge-Slices:** keine. Die ungedeckten Muster bekommen einen, wenn eines
  davon ausfällt — nicht vorher.

- **Risiken aus §7:** drei, jedes mit genau einem Ausgang — zweimal *gestrichen
  mit Begründung* (beide Risiken waren der Grund für die getroffene Wahl),
  einmal *weiter offen* → Register.

- **Drei Paarungen** (Repo ohne Wellen-Betrieb) — geprüft **nach** dem
  `git mv` nach `done/`, weil sie dort suchen; eingetragen im dritten
  Closure-Commit:
  **Anker** — `liegt in Makefile:dcheck-phrase-selftest`; der Zielort trägt
  `seit slice-169`, ebenso `harness/README.md` §Sensors,
  `harness/sensors/dcheck-phrase-selftest.md` und das Skript selbst. *(Bis zur
  Prüfung der Closure-Notiz war der Anker nur an zweien der vier Orte — die
  Paarung meldete grün, weil sie **Existenz** prüfte statt **Anker am
  Zielort**.)*
  **Folge-Slice** — keiner genannt.
  **Register** — vier zitierte Einträge, jeder mit nicht leerem `evidence/`:
  `pruefer-ohne-gegenstand-oder-aufruf` (5), `attestierung-vor-dem-vorgang` (1),
  `hard-rule-37-ohne-sensor` (2), `sensor-ohne-dod-phrase-wirkungslos` (1).

## 9. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** zwei Sub-Areas berührt.
**Gate-/Werkzeug-Schicht** `GATE` (`tools/`, `Makefile`, `.d-check.yml`) —
Achsen 1,2,3, Schwelle erfüllt. **Harness-Einstieg** `HARNESS`
([`AGENTS.md`](../../../../AGENTS.md) §4, [`harness/README.md`](../../../../harness/README.md)
§Sensors, `harness/sensors/`) — Achsen 1,2,3; der Slice ändert dort die
Vertrags-Aussage über das Target, das ist eine Berührung und keine
Pfad-Koinzidenz.

**Vorgelagert — offene Beobachtungen sichten:** gesichtet am 2026-09-08 gegen
den gemergten Stand.

- [`BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf`](../observations/BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf/observation.md)
  — **5×**, der Eintrag, den dieser Slice bedient. Er steht auf *verkörpert*;
  dieser Slice schärft die Verkörperung, statt sie neu zu setzen.
- Unter der Schwelle, in `HARNESS` statt `GATE`, aber derselben Mechanik:
  [`sensor-ohne-dod-phrase-wirkungslos`](../observations/BEO-HARNESS/sensor-ohne-dod-phrase-wirkungslos/observation.md)
  (1×) — eine DoD-Phrase, ohne die ein Sensor stumm bleibt. Er bleibt offen;
  dieser Slice erhöht ihn nicht, weil er die Phrase nicht ändert, sondern ihre
  Kandidatenmenge misst.
- **Keine weiteren Treffer** für `GATE`.

*(Der Block trug bis zur Review-Einarbeitung den Satz „entsteht mit dem
Übergang nach `in-progress/`" — der Übergang war zum Zeitpunkt des Commits
vollzogen, die Sichtung fehlte trotzdem.)*

**Alle berührten Sub-Areas GF** — kein Begründungsblock nötig.
