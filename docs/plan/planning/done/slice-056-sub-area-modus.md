# slice-056 — Etappe C (2/2): Modus pro Sub-Area, offene LOW

**Status:** *(der Zustand ist das Verzeichnis dieser Datei, nicht dieses Feld — korrigiert in slice-063)* — letzter Schnitt der **Etappe C**.
**Deckt:** den echten Kern von Fund **B-2** (pauschaler Repo-Modus) und die offenen LOW **F-7**
bis **F-11** aus dem Etappe-A-Zweit-Review.
[Roadmap](../in-progress/roadmap.md).

---

## 1. B-2 — der Modus gehört der Sub-Area, nicht dem Repo

`conventions.md` deklarierte den Modus in **einer** Zeile: `*` (Default für gesamtes Repo) →
Greenfield. Die Baseline nennt genau das ein Anti-Pattern: der Modus ist ein *beobachtbares
Verhältnis zwischen Code und Doku*, kein Etikett auf dem Repo. Wer ihn pauschal setzt, kann nicht
sagen, **wo** Drift entstünde und **wer** sie sähe.

Ersetzt durch acht benannte Sub-Areas, jede gegen die drei Inklusions-Achsen geprüft (Schwelle:
mindestens zwei von *eigene `MR` plausibel* · *eigene Inventur-Linie* · *eigene Pfad-Familie*).

**Das Ergebnis ist unspektakulär, und das ist der Punkt:** alle sieben Sub-Areas mit Modus stehen
auf **Greenfield**. Das Repo ist spec-first gestartet und hat die Reihenfolge über 55 Slices
gehalten — es gibt nichts zu korrigieren. Der Gewinn liegt nicht in einer Modus-Änderung, sondern
in den **Inventur-Linien**: ab jetzt ist pro Sektion sagbar, was driften würde. Dazu die drei
Drift-Anzeichen der Baseline als Beobachtungsauftrag.

Die achte Sub-Area — die vendored Baseline — bekommt **keinen** Modus: GF und BF beschreiben das
Verhältnis *eigener* Doku zu *eigenem* Code; auf unveränderten Fremdtext ist die Frage nicht
anwendbar. Eine Zeile „GF" wäre dort eine Scheinauskunft.

## 2. Die offenen LOW

| # | Befund | Behandlung |
|---|---|---|
| **F-7** | §Baseline nennt nur den Tag, nicht den Stand | **behoben** — „Kurs-Welle 34 · 2026-07-24", übernommen aus dem Kopf des vendored `regelwerk/README.md`, also belegt statt geschätzt |
| **F-8** | Sensors-Kommentar zeigt auf einen Upstream-Pfad, den es im Repo nicht gibt | **behoben** — zeigt jetzt auf die vendored Vorlage, damit netzlos auflösbar |
| **F-9** | Report-Konvention ohne die Kopf-Metadaten der Vorlage | **behoben** — Review-Art, Skill-Version und Modell-ID sind Pflicht; ohne sie ist ein Report nicht einordbar (derselbe Befund wiegt anders, je nachdem ob Selbst- oder Fremd-Review) |
| **F-10** | Lesepflicht ohne Auswahlregel | **behoben** — das Regelwerk ist Nachschlagewerk **pro Entscheidung**, nicht Pro-Session-Lektüre; acht Aufgabe→Modul-Zuordnungen benannt |
| **F-11** | [MR-006](../../../../harness/conventions.md#mr-006--baseline-committet-vendored-statt-per-url-referenziert) §Geltungsbereich nennt `.d-check.yml` nicht | **bewusst nicht behoben**, siehe §3 |

## 3. F-11 wird nicht behoben — und warum das die Regel ist

Der Befund stimmt: [MR-006](../../../../harness/conventions.md#mr-006--baseline-committet-vendored-statt-per-url-referenziert) begründet in seinem Abgrenzungs-Absatz die `scan.ignore`-Änderung an
[`.d-check.yml`](../../../../.d-check.yml), führt die Datei aber nicht im Feld *Geltungsbereich*.

Das Feld nachträglich zu ergänzen wäre eine **inhaltliche Änderung an einem akzeptierten Eintrag**
— genau das, was die Disziplin des Adaptions-Blocks verbietet und was der vorige Slice als
Regel festgeschrieben hat. Für einen LOW-Befund, dessen Sachverhalt im
selben Eintrag bereits *steht*, ist ein ablösender `MR` unverhältnismäßig: er würde die
Nachvollziehbarkeit verschlechtern, nicht verbessern.

Also bleibt der Eintrag, wie er ist, und dieser Abschnitt ist die Antwort. Das ist der Normalfall
bei Immutabilität — nicht jede erkannte Unvollständigkeit wird geheilt; manche wird **erklärt**.

## 4. Betroffene Module

- [`harness/conventions.md`](../../../../harness/conventions.md) — Modus-Tabelle (B-2), Stand-Zeile (F-7).
- [`harness/README.md`](../../../../harness/README.md) (F-8), [`AGENTS.md`](../../../../AGENTS.md) (F-10),
  [`docs/reviews/README.md`](../../../../docs/reviews/README.md) (F-9).

Zwei Schichten (Konventions-/Harness-Doku, Review-Konvention).

## 5. DoD

- [x] `conventions.md` §Modus-Deklaration führt Sub-Areas mit Achsen-Nachweis statt einer
      `*`-Zeile; die Nicht-Anwendbarkeit auf Fremdtext ist benannt (B-2).
- [x] F-7, F-8, F-9 und F-10 behoben; F-11 mit Begründung ausdrücklich offen gelassen (§3).
- [x] `make gates` und `make verify` grün.

## 6. Closure-Notiz

**Geliefert:** acht benannte Sub-Areas mit Achsen-Nachweis statt einer `*`-Zeile, vier behobene
LOW und ein fünftes ausdrücklich nicht behoben. **Etappe C ist damit vollständig.**

**Lerneintrag — Form: geschärfte Regel.**
> **Eine Deklaration, die überall dasselbe sagt, sagt nichts.** Der pauschale `*`-Modus war nicht
> *falsch* — alle sieben Sub-Areas stehen tatsächlich auf Greenfield. Er war **unbrauchbar**:
> aus ihm ließ sich nicht ableiten, wo Drift entstünde oder wer sie sähe. Der Gewinn der neuen
> Tabelle ist keine Korrektur, sondern die Auflösung in Inventur-Linien. Prüfsatz: *wenn eine
> Deklaration für jede Zeile denselben Wert trägt, prüfen, ob die Zeilen falsch geschnitten sind
> — nicht, ob der Wert stimmt.*

**Zwei beobachtbare Closure-Kriterien:**

1. `make gates` und `make verify` grün auf demselben Stand (je Exit 0) — belegt.
2. Jede der acht Zeilen nennt die erfüllten Inklusions-Achsen; keine steht unter der Schwelle von
   zwei, und die eine ohne Modus ist als nicht-anwendbar begründet, nicht leer gelassen.

**Nicht-Behebung als Ergebnis:** F-11 bleibt offen, weil die Heilung teurer wäre als der Befund
(§3). Das ist die erste Stelle im Repo, an der eine Immutabilitäts-Regel gegen eine erkannte
Unvollständigkeit gewinnt — und damit ein Präzedenzfall, der hier bewusst begründet steht statt
still zu passieren.

**Folge-Slices:** keine aus diesem Slice. Offen aus Etappe B: **F** (Betriebsmodell) — dort auch
die zwei Regeln, die heute einen Guide, aber keinen Sensor haben.

## 7. Sub-Area-Modus

Alle berührten Sub-Areas GF.
