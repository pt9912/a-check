# slice-164 — Regelwerk-Migration `v6.1.0` → `v6.2.0`: Delta-Analyse (Increment)

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `make slice-mv` ([`AGENTS.md`](../../../../AGENTS.md)
§3.3/§5).

**Welle:** [welle-14](welle-14-regelwerk-v610-migration.md) (auf
`v6.2.0` retargeted, s. Retarget-Hinweis dort).

**Bezug:** [slice-161](../done/slice-161-regelwerk-v610-delta-analyse.md)
(dieselbe Analyse-Form, für `v6.0.0`→`v6.1.0`), Maintainer-Hinweis
2026-09-06 ("Es gibt ein neues Regelwerk Release v6.2.0").

**Berührte Spec-Stellen:** — *(keine)* — reine Ist-Messung, keine
Vertragsberührung.

**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim
Maintainer.

**Autor:** Claude (Sonnet 5), im Auftrag des Maintainers. **Datum:**
2026-09-06.

> **Analyse zur Abnahme.** Wie slice-161/163: keine Kennungen vergeben,
> keine Artefakte geändert. Die zwei Nachzug-Vorschläge (§4) gehören vor
> der Umsetzung abgenommen.

**Geltungsbereich der Quelle:** `lab/regelwerk/` und `lab/templates/` aus
einem frischen Klon von `pt9912/ai-harness-course`, Tags `v6.1.0` und
`v6.2.0`.

---

## 1. Anlass und Provenienz

Der Maintainer meldete `v6.2.0`, während `welle-14` noch bei `v6.1.0`
stand und Etappe A (Vendoring) noch nicht gelaufen war. Unabhängig
bestätigt: `gh release view v6.2.0 --repo pt9912/ai-harness-course` — Tag
`v6.2.0`, veröffentlicht 2026-09-05T18:42:24Z, **2h06m nach**
`v6.1.0` (2026-09-05T16:36:01Z), Assets `lab-regelwerk.zip` +
`SHA256SUMS` vorhanden. Kurs-Welle laut `lab/regelwerk/README.md`: 119
(vorher 118). Ergebnis: `welle-14` wurde auf `v6.2.0` retargeted, bevor
Etappe A ein sofort veraltetes `v6.1.0` vendort hätte.

## 2. Umfang des Increments (gemessen)

`git diff --stat v6.1.0 v6.2.0 -- lab/regelwerk lab/templates`: **4
Dateien, +16/−5**.

| Datei | Zeilen |
|---|---|
| `lab/regelwerk/README.md` | Stand-Zeile (Kurs-Welle 118→119), reine Provenienz |
| `lab/regelwerk/modul-05-planning-harness.md` | +3/−3 |
| `lab/templates/.d-check.yml` | +8 |
| `lab/templates/docs/plan/planning/slice.template.md` | +4/−1 |

Zum Vergleich: `v6.0.0`→`v6.1.0` (slice-161) war `+38/−1` über 6 Dateien.
Dieses Increment ist kleiner. **Gesamtsprung `v6.0.0`→`v6.2.0`** (zur
Einordnung, nicht separat behandelt — slice-161/163 decken den
`v6.0.0`→`v6.1.0`-Teil bereits ab): `git diff --stat v6.0.0 v6.2.0 --
lab/regelwerk lab/templates` zeigt 9 Dateien, `+53/−5` — die Summe der
beiden Increments abzüglich Überschneidung (`README.md`-Stand-Zeile zählt
in beiden Diffs, aber nur einmal im kombinierten).

## 3. Was sich **nicht** gedreht hat

- **`lab/regelwerk/README.md`** — nur die Stand-Zeile. Kein Modul-Inhalt
  betroffen.
- **Alle fünf Regelwerk-Module, die von aktiven Adaptionen referenziert
  werden** (`grundlagen-source-precedence.md`, `grundlagen-referenz-richtung.md`,
  `modul-06-roadmap.md`, `modul-08-agentenrollen.md`,
  `modul-15-observability.md` — s. slice-163 §3) — keines davon liegt im
  4-Datei-Diff dieses Increments. Der `slice-163`-Befund
  ([MR-017](../../../../harness/conventions.md#mr-017)s
  Auflösungs-Trigger) bleibt unverändert: die Trigger-Analyse hängt am
  Migrations-*Ereignis*, nicht am Wortlaut, und `v6.2.0` ändert daran
  nichts zusätzlich (kein neues Trigger-Ereignis über das bereits
  gefundene hinaus, da beide Tags Teil derselben Retarget-Entscheidung
  sind).
- **`modul-07`/`modul-10`/`modul-13`** (die drei echten Brocken aus
  slice-161 §4) — unverändert zwischen `v6.1.0` und `v6.2.0`, bleiben bei
  ihrem bisherigen Stand (Vertagt / Vertagt / echter Nachzug, s. slice-161
  §4.1–§4.3).

## 4. Die zwei echten Brocken

### 4.1 `modul-05-planning-harness.md`: „Review-Report" wird zum vierten konstant-nicht-gezählten Posten

Diff (§Ziel-Form: Slice, Größen-Regel):

```diff
- dieses Slice). Nicht gezählt: Gate-Läufe, Closure-Notiz, Register,
- Risiko-Ausgänge, die drei Paarungen; sie sind pro Slice konstant und sagen
- über die Größe nichts · mehrere
+ dieses Slice). Nicht gezählt: Gate-Läufe, Review-Report, Closure-Notiz,
+ Register, Risiko-Ausgänge, die drei Paarungen; sie sind pro Slice konstant
+ und sagen über die Größe nichts · mehrere
```

a-checks eigenes [`AGENTS.md`](../../../../AGENTS.md) §5 (Slice-Form)
führt dieselbe Liste in eigenen Worten: *„Gezählt wird nur, was mit dem
Umfang wächst. Gate-Läufe, Closure-Notiz, Register und Risiko-Ausgänge
zählen **nicht**"* — ohne „Review-Report". **Zusätzlich mechanisch
relevant:** die `tasks-ignore-pattern`-Regex in
[`.d-check.yml`](../../../../.d-check.yml) §`structure` (1) —
`'^( *grün|Closure-Notiz|Beobachtungs-Register|Jedes Risiko|Reconciliation)'`
— kennt kein Muster für einen Review-Eintrag. Gemessen an `slice-163`s
eigener Erfahrung mit diesem Slice selbst: dessen erste Fassung musste ein
DoD-Item „Unabhängiges Plan-Review über getrennten Kontext durchgeführt"
als **echten** Liefer-Punkt zählen (`doc-structure` meldete
`section-oversized` bei vier realen Punkten, s. `slice-163`s
Closure-Historie) — mit dem `v6.2.0`-Zuwachs stellt sich die Frage neu, ob
das richtig war oder ob ein Review-Eintrag künftig konstant sein sollte.

**Echter Nachzug (Text), aber inhaltliche Entscheidung, keine reine
Übernahme:** ob a-check `AGENTS.md` §5 einfach um „Review-Report" ergänzt
**und** die `tasks-ignore-pattern`-Regex passend erweitert, hängt an §4.2
unten — beide Fragen sind gekoppelt, nicht unabhängig lösbar.

### 4.2 `slice.template.md`: neuer Pflicht-Checkbox-Punkt „Review durchgeführt, Report liegt vor"

Diff:

```diff
- Gate-Läufe und die vier Closure-Pflichten darunter zählen nicht mit.
+ Gate-Läufe und die fünf Closure-Pflichten darunter zählen nicht mit.

  - [ ] LH-FA-<NN> erfüllt, Test referenziert.
  - [ ] LH-QA-<NN> erfüllt, Messung dokumentiert.
  - [ ] `make gates` grün.
+ - [ ] Review durchgeführt, Report unter `docs/reviews/` liegt vor
+       (`.harness/skills/reviewer.md`) — Rollenwechsel nach Schritt 8 des
+       Minimal Agent Workflow (`AGENTS.md` §6), kein Self-Review (Modul 8).
  - [ ] Doku-Update für <Schnittstelle X> falls öffentlicher Vertrag berührt.
```

Die Baseline macht den Review-Report damit zu einem **unbedingten**
DoD-Punkt jedes Slice-Plans — nicht mehr optional. a-check hat mit
`slice-159`/`slice-160` bereits einen Review-Mechanismus: `make
doc-reviews` (`DC-FA-RVW-001`, im `gates`-Aggregat) verlangt einen Report
unter `docs/reviews/`, aber **opt-in pro Slice** über die exakte
DoD-Phrase „unabhängiger Review" (`AGENTS.md` §4, Zeile zu
`make doc-reviews`) — ein Slice ohne diese Phrase hat eine leere
Kandidatenmenge und wird nicht geprüft. Das ist eine **andere Form**
derselben Absicht: Baseline macht das Item verpflichtend-konstant,
a-check macht die Prüfung opt-in-scharf.

**Nicht Gegenstand dieser Analyse:** ob a-check von opt-in auf
verpflichtend umstellen soll. Das ist eine echte Architektur-Entscheidung
(Modul 8, Architect-Rolle) mit Rückwirkung auf jeden künftigen Slice-Plan
— sie gehört einem eigenen Folge-Slice, nicht dieser Ist-Messung.

## 5. Vorschlag: ein gebündelter Folge-Slice für §4, unabhängig von den bereits offenen zwei

`welle-14` trägt bereits zwei offene Folge-Slice-Vorschläge (aus
`slice-163`, s. `welle-14` §„Stand"). Dieser Slice fügt einen **dritten**
hinzu, der beide Funde aus §4 bündelt (sie sind inhaltlich gekoppelt, s.
§4.1 Ende) — noch keine Slice-ID vergeben:

- Entscheiden, ob a-check den Review-DoD-Punkt verpflichtend macht (wie
  `v6.2.0`) oder beim Opt-in bleibt (§4.2).
- Je nach Entscheidung: `AGENTS.md` §5 um „Review-Report" ergänzen **und**
  `tasks-ignore-pattern` in `.d-check.yml` passend erweitern (§4.1) — oder
  explizit begründen, warum a-check die abweichende Form beibehält (dann
  als `MR-<NNN>`-Eintrag in `harness/conventions.md`, da es eine Adaption
  ggü. dem `v6.2.0`-Baseline-Default wäre).

Reihenfolge der drei offenen Folge-Slices (dieser, plus die zwei aus
`slice-163`) ist Maintainer-Entscheidung — keiner blockiert die anderen
(§6).

## 6. Risiken und offene Punkte

- *Der vorgeschlagene Folge-Slice (§5) wird nicht gezogen, `AGENTS.md` §5
  und die `tasks-ignore-pattern`-Regex bleiben inkonsistent mit dem
  `v6.2.0`-Baseline-Wortlaut* — **Ausgang:** weiter offen →
  Beobachtungs-Register (kein bestehender Eintrag trifft genau diesen
  Fall — neu angelegt, s. §9).
- *Ob a-check den Review-DoD-Punkt verpflichtend machen sollte, ist eine
  Architektur-Entscheidung, keine reine Übernahme* — **Ausgang:**
  gestrichen mit Begründung: kein Risiko im Sinn der Dreier-Menge, sondern
  die Kernfrage, die dieser Analyse laut Kopf-Vermerk zur Abnahme
  vorgelegt wird.

## 7. DoD

- [x] Increment `v6.1.0`→`v6.2.0` gemessen (4 Dateien, `+16/−5`), jede
      Datei einzeln geprüft, Provenienz unabhängig bestätigt (`gh release
      view`), gegen den slice-161-Präzedenzfall eingeordnet (§1–§3).
- [x] Beide echten Funde benannt, ihre Kopplung erklärt, ein gebündelter
      Folge-Slice statt zwei unabhängiger vorgeschlagen (§4/§5).
- [x] Unabhängiges Plan-Review über getrennten Kontext durchgeführt; die
      dabei gefundene fehlende zweite Sub-Area (§9, Gate-/Werkzeug-Schicht)
      vor Abnahme ergänzt.
- [x] `make gates` grün.
- [x] `make verify` grün.
- [x] Beobachtungs-Register fortgeschrieben (§9).
- [x] Jedes Risiko aus §6 trägt einen Ausgang.

## 8. Closure-Notiz

- **Was hat funktioniert:** das slice-161-Verfahren (frischer Klon,
  dateiweiser `git diff`, Provenienz per `gh release view` unabhängig
  bestätigt) skaliert unverändert auf ein kleineres Increment — kein
  Verfahrens-Nachzug nötig.
- **Was ging anders als geplant:** der Anlass war nicht ein geplanter
  nächster Schritt, sondern eine unangekündigte zweite Release-Meldung
  mitten in einer laufenden Welle — die Reaktion (Welle retargeten statt
  ein zweites, paralleles Migrations-Vorhaben zu eröffnen) war nicht
  vorab geplant, sondern eine Ad-hoc-Entscheidung des Maintainers
  (bestätigt vor diesem Slice).
- **Lerneintrag — Form: geschärfte Regel.** *Erscheint ein neues
  Baseline-Release, während eine Migrations-Welle offensteht und ihre
  Vendoring-Etappe (Etappe A) noch nicht gelaufen ist, wird die Welle auf
  den neuen Stand retargeted, statt das ältere Release zu vendoren und
  sofort erneut migrieren zu müssen. Ist die Vendoring-Etappe bereits
  gelaufen, gilt das nicht mehr — dann ist das neue Release ein eigener,
  neuer Sprung mit eigener Welle/eigenem Slice, kein Retarget.*
- **Beobachtungs-Register (`../observations/`):**
  [`BEO-HARNESS/agents-md-hinkt-baseline-dod-item-hinterher`](../observations/BEO-HARNESS/agents-md-hinkt-baseline-dod-item-hinterher/observation.md)
  neu angelegt, Beleg `evidence/slice-164.md` — Zähler steht bei 1×.
- **Folge-Slices:** noch keine ID vergeben — ein gebündelter Folge-Slice
  für §4/§5, zusätzlich zu dem bereits aus `slice-163` offenen
  ([MR-017](../../../../harness/conventions.md#mr-017)/[MR-000](../../../../harness/conventions.md#mr-000),
  s. `welle-14` §„Stand").
- **Risiken aus §6:** beide mit Ausgang — siehe §6.
- **Drei Paarungen:** verschoben auf die Closure von `welle-14` (dieser
  Slice trägt ein `**Welle:**`-Feld, [Modul 8](../../../../.harness/baseline/v6.0.0/regelwerk/modul-08-agentenrollen.md#rollen-sequenz-für-eine-welle)).

## 9. Sub-Area-Modus

**Vorgelagert — Sub-Area-Wahl prüfen:** zwei Sub-Areas berührt —
**Harness-Einstieg** (`AGENTS.md` §5) und **Gate-/Werkzeug-Schicht**
(`.d-check.yml`s `tasks-ignore-pattern` — Präzedenz für diese Zuordnung:
`BEO-GATE/archiv-sensor-vorpruefung-unvollstaendig` klassifiziert einen
`.d-check.yml`-Fund ebenso unter Gate-/Werkzeug-Schicht, nicht
Harness-Einstieg). Unabhängiges Review fand diese zweite Sub-Area in der
ersten Fassung fehlend (nur eine deklariert). Beide Greenfield, Schwelle ≥
2/3 erfüllt (`harness/conventions.md` §Modus-Deklaration pro Sub-Area).

**Vorgelagert — offene Beobachtungen sichten:** Register für **beide**
Sub-Areas über die Verzeichnisliste geprüft (nicht per Namensabgleich, s.
slice-163 §10-Korrektur). **Harness-Einstieg** (`BEO-HARNESS/`): 8 `offen`
vor diesem Slice, keiner erreicht 3×; mit dem neuen Eintrag dieses Slice
(s. §8) jetzt 9 `offen`, weiterhin keiner bei 3×. **Gate-/Werkzeug-Schicht**
(`BEO-GATE/`): 11 `offen` (5 weitere `verkörpert`), keiner davon trifft
denselben Fund wie dieser Slice (kein Duplikat), keiner erreicht 3×.

**Alle berührten Sub-Areas GF** — kein Begründungsblock nötig.
