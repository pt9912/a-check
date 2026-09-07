# slice-175 — Etappe A: Baseline `v6.5.0` committet vendoren

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `make slice-mv` ([`AGENTS.md`](../../../../AGENTS.md)
§3.3/§5).

**Welle:** [welle-15](../welle-15-regelwerk-v650-migration.md)

**Bezug:** [slice-174](../done/slice-174-regelwerk-v650-delta-analyse.md)
§3.2, Etappen-Schnitt — Etappe **A** des Schnitts in §3.4.

**Berührte Spec-Stellen:** — *(keine)* — Harness-Struktur ohne
Vertragsberührung.

**Verantwortlich:** — *(noch nicht priorisiert)*

**Autor:** Claude (Opus 5), im Auftrag des Maintainers. **Datum:**
2026-09-07.

---

## 1. Ziel und Abgrenzung

Der Stand `v6.5.0` liegt vendored, die Stand-Deklaration nennt ihn an ihren drei Stellen, und die vier Baseline-Symlinks unter `.claude/rules/` zeigen darauf.

Dazu **die Sensoren, die das prüfen, arbeitsfähig halten**: Der erste echte
Migrations-Lauf des `versions`-Musters hat zwei Fehlkalibrierungen gezeigt (§2),
und ein Sensor, der die Migration nicht sieht, macht die Stand-Deklaration zur
Behauptung. Die Korrektur gehört darum hierher und nicht in einen Folge-Slice.

**Nicht in diesem Slice**, je Punkt mit Grund:

- **Den vorigen vendorten Stand entfernen** — ein Folge-Slice übernimmt es:
  [slice-176](../open/slice-176-zitier-form-einfrierende-artefakte.md), und
  zwar mit Kennung und als DoD-Punkt dort. **16 Zeitdokumente** verweisen noch
  auf ihn; ihre Links brächen beim Löschen, und die Reparatur wäre Handarbeit an
  Dateien, die das Repo als unveränderlich führt — genau die Klemme, die
  [`conventions.md`](../../../../harness/conventions.md#baseline) §Baseline als
  *„Preis des Löschens"* beschreibt. Die Zitier-Form aus Etappe C nimmt diesen
  Preis weg, indem sie die Links durch Kennungen ersetzt. Zwei Stände sind
  solange zulässig: die Zusage nennt den Migrations-Fall ausdrücklich, und
  `make regelwerk-check` weist den ungeprüften namentlich aus.
  *(Entschieden vom Maintainer am 2026-09-07. Bei der Umsetzung war es zunächst
  eine ungefragte Implementer-Entscheidung — hier steht sie, weil ein Ausschluss
  ohne Begründung eine Behauptung ist und keine Grenze.)*
- **Andere Etappen von [welle-15](../welle-15-regelwerk-v650-migration.md)** —
  Schicht-Abgrenzung: jede Etappe misst gegen den vendorten Stand und ist
  einzeln lieferbar.
- **Nachrüsten des Altbestands**, wo die Ziel-Form nur für Neues gilt —
  Bestand bleibt bewusst stehen; ein Sensor gegen unentschiedenen Altbestand
  wäre ein Fehlalarm.

## 2. Analyse

### 2.1 Integrität des Vendorings — vier Kanäle

| Kanal | Ergebnis |
|---|---|
| mitgeliefertes Release-`SHA256SUMS` | `80684c17…d18865` |
| Digest der GitHub-API (`assets[].digest`) | identisch |
| eigener Hash des heruntergeladenen ZIP | identisch |
| entpackter Baum ↔ git-Tag `v6.5.0` | Pfadbaum deckungsgleich |

Inhaltlich weichen **28** Dateien vom git-Tag ab — in allen Zeilenpaaren
ausschließlich die beim ZIP-Bau absolutierten Links (relative Repo-Pfade →
`blob/v6.5.0/…`, zweimal `releases/latest/download` → `releases/download/v6.5.0`).
**Keine inhaltliche Abweichung.** Neu ist genau eine Datei:
`templates/harness/sensors/gate.template.md`.

### 2.2 Wortgleichheit der `Ersetzt-Baseline-Regel`-Ziele — nachgeholt

[`conventions.md`](../../../../harness/conventions.md#baseline) §Baseline macht
das Mitwandern eines Zeigers von einer Bedingung abhängig: *„dass der
referenzierte Abschnitt im neuen Stand wortgleich ist, und das ist zu **messen**,
nicht anzunehmen. Ist er es nicht, trägt die Stelle die Abweichung sichtbar,
statt still umzuziehen."*

**Diese Messung ist bei der Umsetzung unterblieben** — die fünf Zeiger wanderten
per `sed` über alle Baseline-Pfade mit. Der unabhängige Review hat es gefangen
(F-2); nachgeholt am 2026-09-07 mit `git diff -w v6.2.0 v6.5.0` je Zieldatei:

| Eintrag | Ziel-Datei | wortgleich? |
|---|---|---|
| [`MR-011`](../../../../harness/conventions.md#mr-011) | `grundlagen-source-precedence.md` | ✅ unverändert |
| [`MR-012`](../../../../harness/conventions.md#mr-012) | `grundlagen-referenz-richtung.md` | ✅ unverändert |
| [`MR-014`](../../../../harness/conventions.md#mr-014) | `modul-15-observability.md` | ✅ unverändert |
| [`MR-016`](../../../../harness/conventions.md#mr-016) | `modul-08-agentenrollen.md` | ✅ unverändert |
| [`MR-015`](../../../../harness/conventions.md#mr-015) | `modul-06-roadmap.md` | ❌ **+3/−1** |

**Die eine Abweichung, sichtbar getragen statt still umgezogen:** In
`modul-06-roadmap.md` §Wellen-Closure-Prozedur ersetzt `v6.5.0` die Zeile
*„im Slice-Plan (Modul 9)"* durch einen Verweis auf die neue §1-Form
(*Ziel und Abgrenzung*) und auf `modul-09`. Betroffen ist damit die
**Out-of-Scope-Zeile der Wellen-Eröffnung**.

[`MR-015`](../../../../harness/conventions.md#mr-015) ersetzt eine andere Zusage desselben Abschnitts — den Closure-Trigger
*„alle Slices in `done/` und `make gates` grün und der **Replay-Lauf** grün"*,
an dessen Stelle bei a-check `make ci` tritt. Dieser Satz ist in `v6.5.0`
**unverändert**. Der Zeiger durfte also mitwandern; was fehlte, war der Beleg
dafür, und der steht jetzt hier.

## 3. Umsetzung

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| `.harness/baseline/v6.5.0/` | neu | Vendoring, 54 Dateien, eigenes `SHA256SUMS` |
| [`conventions.md`](../../../../harness/conventions.md#baseline) §Baseline | update | Stand `v6.5.0`, Kurs-Welle **119 → 128** (gegen den vendorten Kopf gemessen) |
| [`AGENTS.md`](../../../../AGENTS.md) §1, [`harness/README.md`](../../../../harness/README.md) §Guides | update | die zwei übrigen Stand-Deklarationen |
| 13 lebende Dateien, 4 Symlinks | update | Zeiger auf den neuen Stand |
| [`.d-check.yml`](../../../../.d-check.yml) | update | `versions`-Muster: relative Form `../baseline/<tag>/` ergänzt, Verankerung gegen fremde URLs (§2, Review F-3) |

## 3. Umsetzung

*(offen)*

## 4. Definition of Done

- [x] Der neue Stand liegt unter `.harness/baseline/` mit eigenem `SHA256SUMS`; `make regelwerk-check` meldet ihn als den geprüften. *(Zielpfad bewusst ohne Versions-Literal — er existiert erst nach diesem Slice, und `versions` meldete ihn zu Recht als `version-stale`.)*
- [x] `make doc-check` meldet keinen `version-stale` und `make symlink-check` ist grün — beide Sensoren aus slice-173 in ihrem ersten Migrations-Lauf.
- [x] Die im Lauf gefundenen **zwei** Fehlkalibrierungen des `versions`-Musters
      sind behoben und je Richtung mutations-belegt: die übersehene relative
      Form `../baseline/<tag>/` und das Über-Treffen auf fremde URLs (§2).
- [x] Unabhängiger Review durchgeführt (Report unter [`docs/reviews/`](../../../reviews/README.md)).
- [x] `make gates` grün.
- [x] `make verify` grün.
- [x] Jedes Risiko trägt einen Ausgang.

## 5. Trigger

**Start** (`open` → `in-progress`): [slice-174](../done/slice-174-regelwerk-v650-delta-analyse.md)
liegt in `done/` (die Analyse benennt, was zu vendoren ist), Maintainer-Freigabe,
WIP-Limit frei.

**Rückführungen:** wächst der Umfang über die DoD hinaus, zurück nach `next/`
zur Zerlegung. Ändert sich der adoptierte Stand erneut, zurück nach `open/`.

## 6. Closure-Trigger

DoD vollständig, `make gates`/`make verify` grün, Closure-Notiz geschrieben.
Der Slice trägt ein `**Welle:**`-Feld und archiviert **mit seiner Welle**
([`AGENTS.md`](../../../../AGENTS.md) §6).

## 7. Risiken und offene Punkte

- *Die Ziel-Form wird übernommen, ohne dass a-checks Bestand sie trägt — dann
  steht eine Regel da, die der eigene Bestand bricht* — **Ausgang:** eingetreten,
  und zwar sofort: `AGENTS.md` §4 bricht die Vorlagen-Regel *„diese Tabelle
  listet auf"* seit `v5.12.0`, gefunden auf Maintainer-Nachfrage am selben Tag.
  Folge-Slice [slice-177](../open/slice-177-sensors-struktur-zwei-tabellen.md),
  Beleg in
  [`BEO-HARNESS/baseline-regel-nie-erwogen-weil-bestand-sie-verletzt`](../observations/BEO-HARNESS/baseline-regel-nie-erwogen-weil-bestand-sie-verletzt/observation.md)
  (2×).
- *Der Zeiger einer akzeptierten Adaption wandert mit, ohne dass die
  Wortgleichheit gemessen wurde* — **Ausgang:** eingetreten, im Slice
  aufgefangen (§2.2). Die Messung war unterblieben und wurde nachgeholt; vier
  von fünf Zielen sind unverändert, das fünfte trägt die Abweichung jetzt
  sichtbar. Kein Folge-Slice: die Regel stand bereits, sie wurde nur nicht
  befolgt.
- *Der `versions`-Sensor ist an vorhandenen Vorkommen kalibriert und kennt
  darum nicht jede Schreibweise seines Gegenstands* — **Ausgang:** weiter offen
  → Beobachtungs-Register,
  [`BEO-GATE/muster-trifft-nur-die-haeufige-schreibweise`](../observations/BEO-GATE/muster-trifft-nur-die-haeufige-schreibweise/observation.md)
  (1×). Beide Fehlkalibrierungen dieses Laufs sind behoben, die **Klasse** ist
  es nicht.

## 8. Closure-Notiz

**Lerneintrag — Form: geschärfte Regel** (ein Massen-`sed` ersetzt keine
Bedingungs-Prüfung, die das Repo selbst vorschreibt).

- **Was hat funktioniert:** Die Verifikation über **vier** Kanäle statt einen.
  Fremder Text kommt ins Repo, und ein einzelner Digest belegt nur, dass die
  Datei die ist, die der Server geliefert hat. Erst der Vergleich gegen den
  git-Tag zeigt, *was* darin steht — und dass die 28 Abweichungen mechanisch
  sind, war eine Messung von vier Minuten.

- **Was ging anders als geplant:** Die fünf `Ersetzt-Baseline-Regel`-Zeiger
  wanderten per `sed` über alle Baseline-Pfade mit — **ohne** die
  Wortgleichheits-Messung, die
  [`conventions.md`](../../../../harness/conventions.md#baseline) §Baseline
  ausdrücklich verlangt und die
  [slice-172](../done/wellenlos/slice-172-baseline-v600-entfernen.md) als
  Bedingung eingeführt hatte. Der Review hat es gefangen; nachgeholt ergab die
  Messung, dass **eines von fünf** Zielen nicht wortgleich ist. Das Ergebnis
  trägt (die von [`MR-015`](../../../../harness/conventions.md#mr-015) ersetzte Replay-Zusage ist unverändert) — aber das
  wusste vorher niemand, und genau davor warnt der eigene Satz *„zu messen,
  nicht anzunehmen"*.

  Zweitens hat derselbe `sed` eine Schreibweise verfehlt, die auch der Sensor
  nicht kannte: `../baseline/<tag>/`. Dieselbe zu enge Vorstellung an beiden
  Stellen — die Handarbeit und ihr Wächter teilten den blinden Fleck.

- **Steering-Loop-Eintrag — geschärfte Regel:** Wo das Repo eine **Bedingung**
  für eine mechanische Änderung formuliert, ersetzt das Werkzeug, das die
  Änderung ausführt, die Prüfung nicht. Ein `sed` über alle Vorkommen ist die
  Umsetzung, nicht die Entscheidung; die Bedingung ist je Vorkommen zu messen
  und das Ergebnis aufzuschreiben — auch und gerade, wenn es „passt".
  *(Kein `liegt in`-Feld: mit diesem Slice wurde nichts verkörpert. Der
  nächstliegende Ort wäre die Prosa in `conventions.md` §Baseline, die die
  Bedingung schon nennt — sie zu wiederholen hülfe nicht, sie wurde ja
  gelesen und trotzdem nicht befolgt.)*

- **Beobachtungs-Register (`../observations/`):**
  [`BEO-GATE/muster-trifft-nur-die-haeufige-schreibweise`](../observations/BEO-GATE/muster-trifft-nur-die-haeufige-schreibweise/observation.md)
  **neu angelegt**, Beleg `evidence/slice-175.md`, Stand `offen (1×)`. Kein
  Beleg bei
  [`zwei-baseline-staende-nach-migrationsende`](../observations/BEO-HARNESS/zwei-baseline-staende-nach-migrationsende/observation.md):
  Der Doppelstand ist diesmal entschieden und befristet — beobachtet wird dort
  die *Dauer ohne Entscheidung*. Sein `state.md` ist auf den neuen Zustand
  nachgezogen (Review F-1: es behauptete im Präsens das Gegenteil).

- **Folge-Slices:** keine neuen.
  [slice-176](../open/slice-176-zitier-form-einfrierende-artefakte.md) trägt
  seit heute den Lösch-Schritt für den vorigen Stand,
  [slice-177](../open/slice-177-sensors-struktur-zwei-tabellen.md) den
  `AGENTS.md`-§4-Befund. Beide bestanden bereits.

- **Risiken aus §7:** drei, jedes mit genau einem Ausgang — zweimal
  *eingetreten* (einmal Folge-Slice, einmal im Slice aufgefangen), einmal
  *weiter offen* → Register.

- **Drei Paarungen:** trägt die Welle-Closure, nicht dieser Slice — er hat ein
  `**Welle:**`-Feld.

## 9. Sub-Area-Modus

**Vorgelagert — Sub-Area-Wahl prüfen:** entsteht mit dem Übergang nach
`in-progress/`.

**Vorgelagert — offene Beobachtungen sichten:** entsteht mit dem Übergang nach
`in-progress/`; der Register-Stand beim Anlegen ist ein anderer als beim
Beginn der Arbeit.

**Alle berührten Sub-Areas GF** — kein Begründungsblock nötig.
