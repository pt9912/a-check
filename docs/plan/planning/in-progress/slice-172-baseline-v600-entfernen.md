# slice-172 — `v6.0.0` aus dem Vendoring entfernen

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `make slice-mv` ([`AGENTS.md`](../../../../AGENTS.md)
§3.3/§5).

**Welle:** ohne Welle (der Closure-Trigger wäre die eigene DoD — kein
repo-weites Mehr).

**Bezug:** [`BEO-HARNESS/zwei-baseline-staende-nach-migrationsende`](../observations/BEO-HARNESS/zwei-baseline-staende-nach-migrationsende/observation.md)
(1×) — der Eintrag benennt den Zustand und nennt drei mögliche Auflösungen,
ohne eine zu wählen. Der Maintainer hat am 2026-09-06 gewählt: **löschen**,
so dass nur `v6.2.0` vendored bleibt. Vorgeschichte:
[slice-167](../done/slice-167-etappe-a-vendoring-v620.md) §6, das die
Referenz-Klassen aufstellte und `v6.0.0` bewusst liegen ließ.

**Berührte Spec-Stellen:** — *(keine)* — Harness-Konventionen ohne
Vertragsberührung.

**Verantwortlich:** — *(noch nicht priorisiert)*

**Autor:** Claude (Opus 5), im Auftrag des Maintainers. **Datum:**
2026-09-06.

---

## 1. Ziel

`.harness/baseline/v6.0.0/` ist entfernt; genau ein Stand liegt vendored,
wie [`conventions.md`](../../../../harness/conventions.md#baseline)
§Baseline es zusagt. Alle Verweise, die heute dorthin zeigen, lösen danach
auf — ohne dass eine Aussage unwahr wird.

## 2. Analyse (vor der Umsetzung)

### 2.1 Was dem Löschen im Weg steht — gemessen, nicht geschätzt

Löschtest am 2026-09-06: Verzeichnis entfernt, `make doc-check` gelaufen,
Verzeichnis wiederhergestellt. Ergebnis **Exit 2, 22 Befunde**, alle
`target-missing`, keine weitere Befundklasse. Die vier `.claude/rules/`-
Symlinks zeigen bereits auf `v6.2.0` und fallen nicht an — anders als bei
[slice-167](../done/slice-167-etappe-a-vendoring-v620.md), wo genau sie
übersehen wurden.

| Klasse | Links | Ort |
|---|---|---|
| Index-Tabelle *Ersetzt-Baseline-Regel* + Kürzel-Spalte | 6 | [`conventions.md`](../../../../harness/conventions.md) Z. 109–113, 204 |
| Feld `Ersetzt-Baseline-Regel` aktiver Einträge | 5 | [`MR-011`](../../../../harness/conventions.md#mr-011), [`MR-012`](../../../../harness/conventions.md#mr-012), [`MR-014`](../../../../harness/conventions.md#mr-014), [`MR-015`](../../../../harness/conventions.md#mr-015), [`MR-016`](../../../../harness/conventions.md#mr-016) |
| aufgelöster Eintrag → `review-report.template.md` | 1 | [`MR-018`](../../../../harness/conventions.md#mr-018) |
| `done/`-Slices → `modul-08` §Rollen-Sequenz für eine Welle | 6 | [slice-163](../done/slice-163-adaptions-durchgang-v610.md) (2×), [slice-164](../done/slice-164-regelwerk-v620-delta-analyse.md), [slice-165](../done/slice-165-review-dod-punkt-opt-in-beibehalten.md), [slice-166](../done/slice-166-mr020-adr-vorlage-generisch.md), [slice-167](../done/slice-167-etappe-a-vendoring-v620.md) |
| `done/`-Slices → Templates | 3 | [slice-155](../done/slice-155-adr-und-carveouts-readme-form.md) (2×), [slice-162](../done/slice-162-review-pflicht-absatz-v610-wortlaut.md) |
| `done/`-Slice → `AGENTS.template.md` | 1 | [slice-161](../done/slice-161-regelwerk-v610-delta-analyse.md) Z. 154 |

### 2.2 Warum der Bump 21 der 22 Aussagen nicht verändert

Je Linkziel `diff` zwischen beiden vendorten Ständen, ohne die
`<!-- Quelle: … -->`-Zeile, die bei jedem Versionssprung mechanisch
mitläuft ([slice-167](../done/slice-167-etappe-a-vendoring-v620.md) §6 hat
diesen Stempel als Ursache der „31 verschiedenen Dateien" identifiziert):

| Linkziel | abweichende Zeilen |
|---|---|
| `grundlagen-source-precedence.md` · `grundlagen-referenz-richtung.md` · `grundlagen-harness-dateien.md` | 0 · 0 · 0 |
| `modul-06-roadmap.md` · `modul-08-agentenrollen.md` · `modul-15-observability.md` | 0 · 0 · 0 |
| `adr/README.template.md` · `carveouts/README.template.md` · `review-report.template.md` | 0 · 0 · 0 |
| `templates/AGENTS.template.md` | **9** |

Das Feld `Ersetzt-Baseline-Regel` dokumentiert, **welche Regel** ersetzt
wird — nicht, welche Datei-Kopie sie trägt. Steht der Regeltext wortgleich
im verbliebenen Stand, hält der gebumpte Zeiger dieselbe Regel fest; die
Änderung ist eine Pfad-Reparatur nach einem Umzug, wie `make slice-mv` sie
für wandernde Slices fährt, und keine inhaltliche
([`AGENTS.md`](../../../../AGENTS.md) §3.5 analog). **Kein eigener
`MR`-Eintrag** — Maintainer-Entscheidung vom 2026-09-06: die Begründung
trägt diese Sektion, ein Adaptions-Eintrag dafür wäre die vierte Instanz
der Klasse, die
[`BEO-HARNESS/adaption-korrigiert-repo-aussage`](../observations/BEO-HARNESS/adaption-korrigiert-repo-aussage/observation.md)
(2×) bereits als Rückbau-Kandidat führt.

### 2.3 Der eine Zeiger, den der Bump falsch machen würde

[slice-161](../done/slice-161-regelwerk-v610-delta-analyse.md) §4.4 misst
an der verlinkten Vorlage, dass sie *„nach Schritt 8 endet — kein
Rollenwechsel-Absatz"*. Für `v6.0.0` stimmt das; `v6.2.0` trägt den Absatz,
das war der Gegenstand von
[`MR-018`](../../../../harness/conventions.md#mr-018) und
[slice-170](../done/wellenlos/slice-170-mr018-aufloesen.md). Ein stiller
Bump verwandelte einen Messbeleg in eine Falschaussage.

**Entschieden (Maintainer, 2026-09-06):** Zeiger auf `v6.2.0`, unmittelbar
daneben ein Halbsatz, dass die *gemessene* Fassung `v6.0.0` war und die
verlinkte den Absatz bereits trägt. Der Beleg bleibt damit im Repo
nachschlagbar, statt auf eine Netz-URL auszuweichen
([`MR-006`](../../../../harness/conventions.md#mr-006)) oder den Leser auf
das Wort des Slice zu verweisen. Es ist eine inhaltliche Ergänzung an einem
`done/`-Zeitdokument — bewusst, benannt, und der einzige Fall.

## 3. Umsetzung

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| [`conventions.md`](../../../../harness/conventions.md) Z. 109–113, 204 | update | veränderliche Datei, Ziele wortgleich |
| fünf aktive `MR`-Dateien, [`MR-018`](../../../../harness/conventions.md#mr-018) | update | Pfad-Reparatur, §2.2 |
| neun `done/`-Slice-Links | update | Pfad-Reparatur, §2.2 |
| [slice-161](../done/slice-161-regelwerk-v610-delta-analyse.md) Z. 154 | update | Bump **plus** Korrekturhalbsatz, §2.3 |
| [`conventions.md`](../../../../harness/conventions.md#baseline) §Baseline | update | der Absatz begründet heute das Mit-Vendoring; er wird zur Ist-Aussage „genau ein Stand" |
| `.harness/baseline/v6.0.0/` | löschen | Ziel des Slice |

## 4. Definition of Done

- [ ] Alle 22 Zeiger lösen gegen `v6.2.0` auf; der Beleg in
      [slice-161](../done/slice-161-regelwerk-v610-delta-analyse.md) §4.4 nennt die gemessene Fassung weiterhin korrekt.
- [ ] `.harness/baseline/v6.0.0/` ist entfernt, `make regelwerk-check`
      meldet **einen** vendorten Stand ohne „ungeprüft"-Hinweis, und
      [`conventions.md`](../../../../harness/conventions.md#baseline)
      §Baseline beschreibt diesen Zustand statt des alten.
- [ ] Unabhängiger Review durchgeführt (Report unter [`docs/reviews/`](../../../reviews/README.md)).
- [ ] `make gates` grün.
- [ ] `make verify` grün.
- [ ] Jedes Risiko trägt einen Ausgang.

## 5. Trigger

**Start** (`open` → `in-progress`): Maintainer-Freigabe und WIP-Limit frei.

**Rückführungen:** findet der Löschtest nach dem Bump eine Befundklasse,
die §2.1 nicht kennt — etwa eine Prosa-Angabe, die ohne Link falsch wird —,
und wächst der Umfang dadurch über die zwei Liefer-Punkte hinaus: zurück
nach `next/` zur Zerlegung. Stellt sich heraus, dass ein Zielabschnitt doch
nicht wortgleich ist und die Aussage eines akzeptierten Eintrags kippt:
zurück nach `open/`, weil dann die Grundsatzentscheidung aus §2.2 neu zu
treffen ist.

## 6. Closure-Trigger

DoD vollständig, `make gates`/`make verify` grün, Closure-Notiz
geschrieben. Danach Archivierung als wellenloser Slice
([`AGENTS.md`](../../../../AGENTS.md) §6).

## 7. Risiken und offene Punkte

- *Der Bump an akzeptierten Einträgen etabliert eine Präzedenz, die beim
  nächsten Baseline-Sprung ohne die Messung aus §2.2 angewandt wird —
  „Zeiger darf man ja bumpen"* — Ausgang bei Closure.
- *Nach dem Löschen ist die `v6.0.0`-Fassung nur noch über `git` und den
  Kurs-Tag erreichbar; ein künftiger Leser von
  [slice-161](../done/slice-161-regelwerk-v610-delta-analyse.md) §4.4 kann
  den Messbeleg nicht mehr im Arbeitsbaum nachschlagen* — Ausgang bei
  Closure.
- *Die beobachtete Lücke ist mit diesem Slice nicht geschlossen: er trifft
  die Entscheidung einmal, er schafft keinen Wächter, der sie beim nächsten
  Mal einfordert* — Ausgang bei Closure; verwandt
  [`BEO-HARNESS/mr-aufloesungs-trigger-ohne-waechter`](../observations/BEO-HARNESS/mr-aufloesungs-trigger-ohne-waechter/observation.md).

## 8. Closure-Notiz

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice;
Lerneintrag — Form: wird dort benannt.)_

## 9. Sub-Area-Modus

**Vorgelagert — Sub-Area-Wahl prüfen:** zwei Sub-Areas berührt.
**Harness-Einstieg** (`harness/`, `.harness/baseline/`) — Achsen 1,2,3.
**Planungs-Harness** (`docs/plan/planning/`) — Achsen 1,2,3, berührt über
die neun `done/`-Slice-Zeiger und das Beobachtungs-Register. Beide über der
Schwelle ≥ 2/3, beide in
[`conventions.md`](../../../../harness/conventions.md#modus-deklaration-pro-sub-area)
deklariert.

**Vorgelagert — offene Beobachtungen sichten:** Register am 2026-09-06
durchgegangen, 23 offene Einträge in beiden Sub-Areas.

| Eintrag | Stand | Bezug zu diesem Slice |
|---|---|---|
| [`zwei-baseline-staende-nach-migrationsende`](../observations/BEO-HARNESS/zwei-baseline-staende-nach-migrationsende/observation.md) | 1× | **direkt adressiert** — der Slice ist ihre Auflösung; Ausgang bei Closure zu wählen |
| [`adaption-korrigiert-repo-aussage`](../observations/BEO-HARNESS/adaption-korrigiert-repo-aussage/observation.md) | 2× | **nicht erhöht** — genau deshalb entsteht kein `MR`-Eintrag (§2.2); ein vierter Repo-Aussage-Korrektur-Eintrag hätte die Schwelle erreicht |
| [`mr-aufloesungs-trigger-ohne-waechter`](../observations/BEO-HARNESS/mr-aufloesungs-trigger-ohne-waechter/observation.md) | 1× | berührt, **nicht erhöht** — der Slice fasst fünf aktive Einträge an, prüft aber ihre Auflösungs-Trigger nicht; er ändert Zeiger, nicht Trigger |
| [`review-geltungsbereich-zu-eng`](../observations/BEO-PLAN/review-geltungsbereich-zu-eng/observation.md) | 1× | **Zuordnung offen** — der Beleg in [`zwei-baseline-staende…/evidence/slice-170.md`](../observations/BEO-HARNESS/zwei-baseline-staende-nach-migrationsende/evidence/slice-170.md) maß „alle 20 `MR`-Dateien" und übersah die neun `done/`-Slice-Zeiger. Ob eine zu eng gezogene **Messung** dieselbe Beobachtung ist wie ein zu eng gezogener **Review**, ist ein Urteil und fällt bei der Closure, nicht hier |

**Alle berührten Sub-Areas GF** — kein Begründungsblock nötig.
