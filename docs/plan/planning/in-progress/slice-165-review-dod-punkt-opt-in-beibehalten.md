# slice-165 — Review-Checkbox-Punkt: Opt-in beibehalten, Phrase-Konsistenz herstellen

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `make slice-mv` ([`AGENTS.md`](../../../../AGENTS.md)
§3.3/§5).

**Welle:** [welle-14](../welle-14-regelwerk-v610-migration.md).

**Bezug:** [slice-164](../done/slice-164-regelwerk-v620-delta-analyse.md)
§4.2/§5 (Folge-Slice-Vorschlag), Maintainer-Wort 2026-09-06
("Review-DoD-Punkt").

**Berührte Spec-Stellen:** — *(keine)* — Harness-/Konventions-Änderung
ohne Vertragsberührung.

**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim
Maintainer.

**Autor:** Claude (Sonnet 5). **Datum:** 2026-09-06.

---

## 1. Ziel

Entscheiden, ob a-check den in `v6.2.0` neu verpflichtenden Review-DoD-Punkt
übernimmt oder beim bisherigen Opt-in (`make doc-reviews`, `DC-FA-RVW-001`)
bleibt — und dabei einen bei der Prüfung selbst gefundenen echten Defekt
beheben: die Opt-in-Phrase, die a-checks eigene bisherige Slices tatsächlich
verwendet haben, löst den Sensor **nicht** aus.

## 2. Definition of Done

- [x] Empirisch geprüft (isolierter `d-check`-Lauf, nicht nur Doku
      gelesen), welche Wortlaute die `reviews`-Modul-Phrase auslösen und
      welche nicht (§3).
- [x] Entscheidung getroffen und begründet: Opt-in bleibt, verpflichtender
      Baseline-Checkbox-Import wird **nicht** übernommen (§4).
- [x] `AGENTS.md` §5 und `.d-check.yml` (`tasks-ignore-pattern`)
      nachgezogen, neuer `MR`-Eintrag für die Divergenz (§5).
- [x] `make gates` grün.
- [x] `make verify` grün.
- [x] Beobachtungs-Register fortgeschrieben (§9) — dritte Evidenz für
      `BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf` (Schwelle 3× erreicht,
      Ausgang bei der nächsten Welle-Closure).
- [x] Jedes Risiko aus §6 trägt einen Ausgang.

## 3. Empirische Prüfung: welche Phrase löst `make doc-reviews` tatsächlich aus?

`AGENTS.md` §4 dokumentiert den Trigger als *„eine `done/`-Slice-DoD-Zeile
mit der Phrase „unabhängiger Review""*. Statt das zu glauben, isoliert
getestet: eigenes Scratch-Repo mit minimalem `.d-check.yml`
(`modules: [reviews]`) und je einer Test-Slice-Datei pro Wortlaut, `docker
run` gegen den in a-check gepinnten `d-check`-Digest.

| Getesteter Wortlaut (Auszug der DoD-Zeile) | Löst `review-missing` aus? |
|---|---|
| „Review durchgeführt, Report unter `docs/reviews/` liegt vor." (exakter `v6.2.0`-Baseline-Wortlaut) | **Nein** |
| „unabhängiger Review durchgeführt." | Ja |
| „Unabhängiger Review durchgeführt, Report liegt vor." (großes U) | Ja |
| „unabhängiges Review durchgeführt." (Neutrum, ohne „-er") | **Nein** |
| „Unabhängiges Plan-Review über getrennten Kontext durchgeführt." | **Nein** |

**Ergebnis:** Die Phrase ist wortform-scharf — exakt „unabhängig**er** Review"
(Groß-/Kleinschreibung des ersten Buchstabens ist egal, die Endung „-er" und
das stehende Wort „Review" **ohne** Compositum sind es nicht). Weder die
`v6.2.0`-Baseline-Formulierung noch a-checks **eigene, tatsächlich
verwendete** Formulierung („Unabhängiges Plan-Review …") lösen den Sensor
aus.

## 4. Der eigentliche Fund: `make doc-reviews` hat vier eigene Slices nie geprüft

[`slice-161`](../done/slice-161-regelwerk-v610-delta-analyse.md),
[`slice-162`](../done/slice-162-review-pflicht-absatz-v610-wortlaut.md),
[`slice-163`](../done/slice-163-adaptions-durchgang-v610.md) und
[`slice-164`](../done/slice-164-regelwerk-v620-delta-analyse.md) führen in
ihrer DoD je eine Zeile „Unabhängiges Plan-Review über getrennten Kontext
durchgeführt" (bzw. Variationen davon) — **keine** trifft die exakte
Trigger-Phrase aus §3. `make doc-reviews` hat für alle vier Slices **nichts
geprüft**: leere Kandidatenmenge, grün ohne Gegenstand — obwohl für jeden
dieser vier Slices tatsächlich ein unabhängiges Review mit eigenem Report
unter `docs/reviews/` durchgeführt wurde (real, nicht behauptet — die
Reports existieren). Das Review-**Ergebnis** war also jedes Mal echt; nur
die **mechanische Prüfung**, dass es stattgefunden hat, lief nie.

Das ist dieselbe Fund-Klasse wie
[`BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf`](../observations/BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf/observation.md)
(bisher 2×: `verify-ac-form`/slice-120, `doc-complete`/slice-123) — ein
Prüfer meldet grün, weil seine Prüfmenge leer ist, nicht weil geprüft wurde.
Mit diesem Fund erreicht der Eintrag **3×** (§9) — die Schwelle, ab der es
laut `AGENTS.md` §5 „eine Harness-Lücke [ist und] einen Guide oder Sensor
verlangt". Die Zuweisung des Ausgangs (verkörpert/geplant/gestrichen) ist
Aufgabe des **Lese-Schritts bei der nächsten Welle-Closure** (`welle-14`,
Modul 8 §Rollen-Sequenz für eine Welle) — nicht dieses Slice, der nur die
dritte Evidenz beisteuert.

**Nicht rückwirkend behoben:** die DoD-Formulierung in `slice-161`–`164`
selbst wird nicht geändert — geschlossene Slices werden nicht nachträglich
umgeschrieben (dieselbe Zurückhaltung wie bei ADRs, `AGENTS.md` §3.5, auch
wenn Slices dort nicht explizit genannt sind). Der Fund wirkt nur nach vorn.

## 5. Entscheidung und Umsetzung

**Entscheidung:** a-check bleibt beim **Opt-in** (`make doc-reviews`,
Phrase-getriggert) statt den `v6.2.0`-Baseline-Checkbox unbedingt zu
übernehmen. Begründung: `AGENTS.md` §6 verlangt Review bereits präzise
*„bei jedem Slice mit Code- oder Vertragsänderung"* — nicht bei jedem
Slice. Ein unbedingter Checkbox-Punkt (wie `v6.2.0`) würde entweder bei
reinen Analyse-Slices (wie diesem hier) unnötig auslösen oder, wenn generisch
formuliert, wie in §3 gezeigt **gar nicht** auslösen — beides schlechter
als der bestehende, aber korrekt formulierte Opt-in.

Drei Liefer-Punkte:

1. **`AGENTS.md` §5** — „Review-Report" zur Liste der pro Slice konstanten,
   nicht gezählten Posten ergänzt (deckungsgleich mit `v6.2.0`s
   `modul-05`-Zuwachs); neuer fünfter Kopier-Hinweis: sobald die vendored
   Ziel-Form eine Review-DoD-Zeile führt (ab `v6.2.0`, noch nicht vendored),
   wird ihr Wortlaut beim Kopieren auf die exakte Trigger-Phrase
   „unabhängiger Review" umgeschrieben (§3) statt den Baseline-Wortlaut
   unverändert zu übernehmen.
2. **`.d-check.yml`** (`structure`, Größen-Regel) — `tasks-ignore-pattern`
   um die korrekte Phrasenform ergänzt, damit eine künftig richtig
   formulierte Review-Zeile nicht als Liefer-Punkt zählt.
3. **Neuer `MR`-Eintrag** in `harness/conventions.md` — dokumentiert die
   Opt-in-vs-verpflichtend-Divergenz ggü. dem `v6.2.0`-Baseline-Default,
   analog zu [`MR-017`](../../../../harness/conventions/MR-017-adr-vorlagen-version.md)/[`MR-018`](../../../../harness/conventions/MR-018-review-pflicht-v610-wortlaut.md):
   kein vendorter Anker verfügbar (Treiber ist das noch nicht vendorte
   `v6.2.0`-Template), Rückbau-Kandidat, sobald `v6.2.0` vendored ist.

## 6. Risiken und offene Punkte

- *Künftige Slices formulieren die Review-DoD-Zeile weiterhin falsch (nicht
  „unabhängiger Review" wortgleich) und `make doc-reviews` bleibt weiter
  wirkungslos für sie* — **Ausgang:** weiter offen → Beobachtungs-Register
  (`BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf`, jetzt bei 3×, Ausgang
  folgt bei `welle-14`-Closure).
- *Ob a-check dauerhaft beim Opt-in bleiben soll oder mittelfristig auf
  einen unbedingten Checkbox-Punkt umstellt, ist eine Grundsatzfrage, die
  mit diesem Slice nur für den aktuellen `v6.2.0`-Stand beantwortet wird*
  — **Ausgang:** gestrichen mit Begründung: kein Risiko im Sinn der
  Dreier-Menge, sondern die Kernfrage, die dieser Slice laut Auftrag
  beantwortet; eine erneute Grundsatzdebatte bräuchte einen neuen Anlass
  (z. B. eine weitere Baseline-Migration), keinen offenen Punkt hier.

## 7. Closure-Notiz

- **Was hat funktioniert:** die empirische Prüfung gegen den echten
  `d-check`-Digest statt gegen die Doku-Beschreibung hat einen Fund
  aufgedeckt, den reine Lektüre nie gefunden hätte — `AGENTS.md` §4s
  eigene Beschreibung der Phrase ist korrekt, nur a-checks eigene
  **Praxis** wich davon ab.
- **Was ging anders als geplant:** der Auftrag war „Review-DoD-Punkt"
  (v6.2.0-Konvergenz einordnen); tatsächlich lag der größere Fund nicht im
  Baseline-Vergleich, sondern in einem repo-eigenen, vier Slices alten
  Bestandsfehler, der erst durch den empirischen Test sichtbar wurde.
- **Lerneintrag — Form: geschärfte Regel.** *Eine dokumentierte
  Trigger-Phrase für einen phrase-basierten Sensor wird vor Verwendung in
  einem neuen Slice-DoD **empirisch gegen den gepinnten Sensor-Digest**
  geprüft (isoliertes Scratch-Repo, nicht nur die Dokumentationszeile
  gelesen) — ein Prüfer mit leerer Kandidatenmenge meldet grün, ohne
  irgendetwas geprüft zu haben, und eine nahe, aber falsche Formulierung
  ist von der korrekten am gemeldeten Ergebnis nicht zu unterscheiden.*
- **Beobachtungs-Register (`../observations/`):**
  `BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf` um `evidence/slice-165.md`
  ergänzt — Zähler steht jetzt bei **3×** (vorher slice-120, slice-123).
  Ausgang wird beim Lese-Schritt der nächsten Welle-Closure (`welle-14`)
  zugewiesen (Modul 8 §Rollen-Sequenz für eine Welle).
- **Folge-Slices:** noch keine ID vergeben — abhängig vom Ausgang des
  3×-Treffers oben (§9); zusätzlich bleibt der aus `slice-163` offene
  Folge-Slice
  ([MR-017](../../../../harness/conventions/MR-017-adr-vorlagen-version.md)/[MR-000](../../../../harness/conventions.md#mr-000))
  unverändert offen.
- **Risiken aus §6:** beide mit Ausgang — siehe §6.
- **Drei Paarungen:** verschoben auf die Closure von `welle-14` (dieser
  Slice trägt ein `**Welle:**`-Feld, [Modul 8](../../../../.harness/baseline/v6.0.0/regelwerk/modul-08-agentenrollen.md#rollen-sequenz-für-eine-welle)).

## 8. Sub-Area-Modus

**Vorgelagert — Sub-Area-Wahl prüfen:** zwei Sub-Areas berührt —
**Harness-Einstieg** (`AGENTS.md` §5, neuer `MR`-Eintrag) und
**Gate-/Werkzeug-Schicht** (`.d-check.yml`s `tasks-ignore-pattern`), beide
Greenfield, Schwelle ≥ 2/3 erfüllt.

**Vorgelagert — offene Beobachtungen sichten:** beide Register über die
Verzeichnisliste geprüft. **Harness-Einstieg** (`BEO-HARNESS/`): 9 `offen`
(unverändert seit slice-164), keiner erreicht 3×. **Gate-/Werkzeug-Schicht**
(`BEO-GATE/`): 11 `offen` vor diesem Slice — darunter
`pruefer-ohne-gegenstand-oder-aufruf` bei 2×, mit diesem Slice auf 3×
(s. §4/§9-Closure-Notiz); die anderen zehn unverändert, keiner sonst
erreicht 3×.

### Sub-Area: Harness-Einstieg

- **Modus:** Greenfield
- **Konventionen-Dichte:** `AGENTS.md` §5 (Slice-Form) und der
  Adaptions-Block in `harness/conventions.md` sind beide bereits etablierte
  Konventionsträger für genau diese Art Änderung (Ergänzung der
  konstanten DoD-Posten-Liste, neuer `MR`-Eintrag).
- **Phase-Reife:** Phase 5 (etablierte, mehrfach genutzte Konvention).
- **Evidenz-/Diskrepanz-Risiko:** niedrig — die Änderung ist additiv, kein
  bestehender Text wird widersprüchlich.
- **Reconciliation-Aufwand:** keiner; kein Brownfield-Bestand betroffen.

### Sub-Area: Gate-/Werkzeug-Schicht

- **Modus:** Greenfield
- **Konventionen-Dichte:** `.d-check.yml` ist vollständig durch
  `AGENTS.md` §4/`harness/README.md` §Sensors dokumentiert; jede Modul-Regel
  hat eine Bindung.
- **Phase-Reife:** Phase 5.
- **Evidenz-/Diskrepanz-Risiko:** mittel — dieser Slice selbst **ist** der
  Beleg, dass eine Regex-Anpassung ohne empirischen Test unbemerkt driften
  kann (§3/§4); das Risiko wird durch den empirischen Test in diesem Slice
  selbst gemindert, nicht eliminiert (künftige Änderungen könnten denselben
  Fehler wiederholen).
- **Reconciliation-Aufwand:** keiner; kein Brownfield-Bestand betroffen.
