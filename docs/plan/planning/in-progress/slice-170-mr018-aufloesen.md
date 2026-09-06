# slice-170 — `MR-018` auflösen: der Rückbau-Trigger ist eingetreten

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `make slice-mv` ([`AGENTS.md`](../../../../AGENTS.md)
§3.3/§5).

**Welle:** ohne Welle (der Closure-Trigger wäre die eigene DoD — kein
repo-weites Mehr).

**Bezug:** [`MR-018`](../../../../harness/conventions/done/MR-018-review-pflicht-v610-wortlaut.md)
nennt seinen eigenen Rückbau-Trigger; dieser Slice vollzieht ihn.
Beobachtung:
[`BEO-HARNESS/rueckbau-kandidat-ueberlebt-baseline-migration`](../observations/BEO-HARNESS/rueckbau-kandidat-ueberlebt-baseline-migration/observation.md).

**Berührte Spec-Stellen:** — *(keine)* — Harness-Konventionen ohne
Vertragsberührung.

**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim
Maintainer.

**Autor:** Claude (Opus 5), im Auftrag des Maintainers. **Datum:**
2026-09-06.

---

## 1. Ziel

[`MR-018`](../../../../harness/conventions/done/MR-018-review-pflicht-v610-wortlaut.md) auflösen. Die Adaption war eine **Provenienz-Korrektur**: der
Wortlaut des Rollenwechsel-Absatzes in [`AGENTS.md`](../../../../AGENTS.md)
§6 stammte aus einem Kurs-Template, das a-check nicht vendored hatte —
also gab es keinen netzlos auflösbaren Anker, auf den das Feld
`Ersetzt-Baseline-Regel` hätte zeigen können. Dieser Grund ist entfallen.

## 2. Analyse (vor der Umsetzung)

**Gemessen, nicht angenommen** — `diff` der beiden vendorten Stände:

```text
diff .harness/baseline/v6.0.0/templates/AGENTS.template.md \
     .harness/baseline/v6.2.0/templates/AGENTS.template.md
236a237,243
> Dieser Workflow deckt ausschließlich die Implementer-Rolle ab. Schritt 8
> ist der Rollenwechsel, kein Abschluss: Bericht → Handoff an Reviewer …
```

Der Absatz, um den es geht, liegt seit
[`slice-167`](../done/slice-167-etappe-a-vendoring-v620.md) **netzlos im
Repo**. Der Eintrag nennt **zwei** Bedingungen; eingetreten ist genau eine —
und es ist die, die den Rückbau trägt:

| Bedingung aus dem Eintrag | Zustand |
|---|---|
| **Rückbau-Kandidat**: „sobald der Stand vendored ist und auf das dann vendorte Template gezeigt werden kann" | **eingetreten** (slice-167) — die Bedingung ist a-check-lokal und beschreibt genau das Vendoring |
| **Auflösungs-Trigger**: „die nächste Baseline-Migration, die `modul-08` oder den Rollenwechsel-Absatz in `AGENTS.template.md` **inhaltlich ändert**" | **nicht** eingetreten — der Absatz stammt aus dem `v6.1.0`-Zuwachs und steht in `v6.2.0` unverändert; `modul-08` unterscheidet sich nur in der `<!-- Quelle: … -->`-Zeile. Upstream hat am Gegenstand nichts geändert; a-check hat ihn nur **nachvollzogen** |

**Die Auflösung trägt allein aus der ersten Bedingung** — der Eintrag hat
sie selbst als seinen Rückbau-Weg benannt. Der zweite Trigger bleibt
unerfüllt und ist es auch nach der Auflösung: er wäre der Weg für eine
Adaption, die aus einer Baseline-**Änderung** entstanden wäre. Diese
entstand aus einer Baseline-**Lücke** im vendorten Stand, und die hat das
Vendoring geschlossen.

**Warum kein Nachfolge-Eintrag.** Was nach dem Rückbau von der Abweichung
bliebe, sind zwei a-check-eigene Sätze (`fork`-Ausschluss, `BEO`-Anker).
Die **schärfen** `modul-08` §Kontext-Trennung und ersetzen dort nichts;
eine ausgefüllte Ziel-Form, die mehr sagt als ihre Vorlage, ist keine
Adaption — sonst bräuchte jede gefüllte Vorlage einen Eintrag. Präzedenz
für eine Auflösung **durch Ereignis** statt durch Nachfolger ist
[`MR-003`](../../../../harness/conventions.md#mr-003).

**Was die Auflösung *nicht* bewirkt.** Gemessen über das Feld
`Ersetzt-Baseline-Regel` aller 20 `MR`-Dateien: genau **fünf** tragen dort
einen `v6.0.0`-Anker ([`MR-011`](../../../../harness/conventions.md#mr-011), [`MR-012`](../../../../harness/conventions.md#mr-012), [`MR-014`](../../../../harness/conventions.md#mr-014), [`MR-015`](../../../../harness/conventions.md#mr-015), [`MR-016`](../../../../harness/conventions.md#mr-016)) —
vor wie nach diesem Slice dieselben fünf. [`MR-018`](../../../../harness/conventions.md#mr-018) trug dort seit
`541581f` ein „—" und war nie Teil der Menge; sein einziger
`v6.0.0`-Verweis steht im Begründungsfeld und zieht mit nach
`conventions/done/` um. Die erste Fassung dieses Plans behauptete „einen
von sechs entfernt" — die Zahl stammt ungeprüft aus
[`slice-167`](../done/slice-167-etappe-a-vendoring-v620.md) §3 und
`welle-14-results.md`, die [`MR-018`](../../../../harness/conventions.md#mr-018) mitzählen, obwohl sein Feld schon damals
leer war. Der unabhängige Review hat das gefangen.

**Was der Fund über den Prozess sagt.** `welle-14-results.md` schreibt,
der Eintrag löse sich „erst mit der jeweils nächsten Baseline-Migration" —
notiert **nach** slice-167, also nachdem diese Migration den Trigger
gezogen hatte. Die Ursache ist strukturell: der **Trigger-Audit** der
Wellen-Closure (`modul-06`, Schritt 2) prüft *Carveout · bootstrap-aware
Gate · ADR*; `MR`-Einträge stehen nicht darin. Ein `MR` mit
Auflösungs-Trigger hat damit keinen Wächter und wird nur gefunden, wenn
jemand ihn von Hand liest — hier: der Maintainer, im Durchgang durch alle
aktiven Einträge.

## 3. Umsetzung

1. [`MR-018`](../../../../harness/conventions/done/MR-018-review-pflicht-v610-wortlaut.md) per `git mv` nach
   [`harness/conventions/done/`](../../../../harness/conventions/done/) —
   Datei-Inhalt **unverändert** (Adaptions-Block-Disziplin: an einem
   akzeptierten Eintrag wird nichts nachträglich geändert).
2. [`harness/conventions.md`](../../../../harness/conventions.md): Zeile
   verlässt *Aktive Adaptionen*, tritt in *Aufgelöste Adaptionen* mit dem
   Auflösungs-Ereignis; der Anker `mr-018` zieht mit um und hält alle
   bestehenden Verweise gültig.
3. Verweise auf den alten Dateipfad repo-weit nachgezogen. **Nicht**
   angefasst: die Review-Reports — sie nennen den Pfad als Inline-Code in
   ihrer Eingangs-Kontext-Liste, nicht als Link, und sind Lauf-Belege eines
   vergangenen Standes.

## 4. Definition of Done

- [x] [`MR-018`](../../../../harness/conventions/done/MR-018-review-pflicht-v610-wortlaut.md) liegt in `conventions/done/`, Inhalt unverändert; beide
      Tabellen in `conventions.md` nachgezogen, Anker `mr-018` erhalten.
- [x] Jeder **verlinkende** Verweis auf den alten Pfad zeigt auf den neuen;
      `make doc-check` belegt es.
- [x] Unabhängiger Review durchgeführt (Report unter `docs/reviews/`).
- [x] `make gates` grün.
- [x] `make verify` grün.
- [x] Jedes Risiko trägt einen Ausgang.

## 5. Trigger

**Start** (`open` → `in-progress`): Maintainer-Anweisung im Durchgang durch
die aktiven Adaptionen (2026-09-06), WIP-Limit frei.

**Rückführungen:** keine vorgesehen — zwei Liefer-Punkte, eine Schicht.

## 6. Closure-Trigger

DoD vollständig, `make gates`/`make verify` grün, Closure-Notiz
geschrieben.

## 7. Risiken und offene Punkte

- *Der Anker `mr-018` wird beim Umzug in die zweite Tabelle vergessen —
  dann brechen die Verweise aus [`MR-019`](../../../../harness/conventions/MR-019-review-dod-opt-in.md) und `slice-162`, die auf
  `conventions.md#mr-018` zeigen* — **Ausgang:** gestrichen mit Begründung:
  der Anker ist mit umgezogen und beide Verweise lösen auf, belegt durch
  `make doc-check` Exit 0 (das Risiko kann in dieser Form nicht mehr
  eintreten, die Bewegung ist vollzogen).
- *Die Auflösung bringt `.harness/baseline/v6.0.0/` seiner Ablösung nicht
  näher; wer einen Schritt für Fortschritt hält, verwechselt ihn mit dem
  Ziel* — **Ausgang:** weiter offen → Beobachtungs-Register, neuer Eintrag
  [`BEO-HARNESS/zwei-baseline-staende-nach-migrationsende`](../observations/BEO-HARNESS/zwei-baseline-staende-nach-migrationsende/observation.md).
  Die Auflösung verlangt eine Entscheidung über alle fünf verbliebenen
  Einträge zugleich; sie in diesem Slice zu treffen hieße, sie nebenbei zu
  treffen.

## 8. Closure-Notiz

- **Was hat funktioniert:** die Auflösung stand als Frage im Eintrag selbst
  („Rückbau-Kandidat, sobald …") — sie musste nicht erfunden, nur gemessen
  werden. Der `diff` der beiden vendorten `AGENTS.template.md` beantwortet
  sie in vier Zeilen.
- **Was ging anders als geplant:** der Slice war als Aufräumarbeit gedacht
  und hat einen Prozess-Befund geliefert. Der Trigger war **innerhalb** der
  Migration eingetreten, die den Eintrag geschrieben hat; die Closure-Notiz
  derselben Welle behauptet das Gegenteil. Nicht Unaufmerksamkeit, sondern
  ein fehlender Wächter.
- **Lerneintrag — Form: benannte Spec-Lücke.** *Der Trigger-Audit — in
  `modul-06` Closure-Schritt 2, im wellenlosen Betrieb getragen von der
  Slice-Closure — kommt in a-checks eigenem Harness **null Mal** vor:
  `AGENTS.md`, `harness/README.md` und `docs/plan/planning/README.md` nennen
  ihn nicht. Damit ist kein Trigger gewächtert, weder der eines Carveouts
  noch der eines bootstrap-aware Gate noch der einer ADR — und `MR`-Einträge
  zählt schon die Baseline nicht mit, obwohl sie dasselbe Pflichtfeld
  führen. Belegt an diesem Slice: der Rückbau-Trigger von [`MR-018`](../../../../harness/conventions.md#mr-018) war seit
  slice-167 eingetreten und wurde erst beim Durchgang von Hand gefunden.*
- **Beobachtungs-Register (`../observations/`):** **zwei** neue Einträge, je
  1× mit Beleg `slice-170` —
  [`BEO-HARNESS/mr-aufloesungs-trigger-ohne-waechter`](../observations/BEO-HARNESS/mr-aufloesungs-trigger-ohne-waechter/observation.md)
  (der Prozess-Befund; **nicht** in
  [`rueckbau-kandidat-ueberlebt-baseline-migration`](../observations/BEO-HARNESS/rueckbau-kandidat-ueberlebt-baseline-migration/observation.md)
  eingetragen, Begründung in §9) und
  [`BEO-HARNESS/zwei-baseline-staende-nach-migrationsende`](../observations/BEO-HARNESS/zwei-baseline-staende-nach-migrationsende/observation.md)
  (der Ausgang des zweiten Risikos).
- **Folge-Slices:** keine. Die Querschnitts-Frage zu den fünf verbliebenen
  `v6.0.0`-Ankern hängt jetzt am Register statt an einem Satz in einer
  Closure-Notiz.
- **Risiken aus §7:** beide mit Ausgang — siehe §7.
- **Drei Paarungen:** entfällt — dieser Slice ist wellenlos (`**Welle:**
  ohne Welle`).

## 9. Sub-Area-Modus

**Vorgelagert — Sub-Area-Wahl prüfen:** eine Sub-Area berührt —
**Harness-Einstieg** (`harness/`), Greenfield, Schwelle ≥ 2/3 erfüllt.

**Vorgelagert — offene Beobachtungen sichten:** `BEO-HARNESS/` über die
Verzeichnisliste geprüft. Naheliegend war
[`rueckbau-kandidat-ueberlebt-baseline-migration`](../observations/BEO-HARNESS/rueckbau-kandidat-ueberlebt-baseline-migration/observation.md)
— die Überschrift passt, die **Ursache nicht**: dort wird ein Trigger
*gewählt* (der billigere Ersatz statt der sauberen Auflösung), hier wird
keiner gewählt, weil niemand hinsieht. Das Gegenmittel dieses Eintrags
(generischer Verweis statt Versionsnummer,
[`MR-020`](../../../../harness/conventions/MR-020-adr-vorlage-generisch.md))
greift für den vorliegenden Fall nicht. Darum ein **eigener** Eintrag
[`mr-aufloesungs-trigger-ohne-waechter`](../observations/BEO-HARNESS/mr-aufloesungs-trigger-ohne-waechter/observation.md)
(1×) statt einer verdünnten Zählung. Kein weiterer Eintrag unter
`BEO-HARNESS/` betrifft diesen Slice.

**Alle berührten Sub-Areas GF** — kein Begründungsblock nötig.
