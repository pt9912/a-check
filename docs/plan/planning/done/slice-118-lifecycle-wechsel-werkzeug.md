# slice-118 — Der Lifecycle-Wechsel zieht seine Verweise selbst nach

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese Datei liegt — eines von
`open/`, `next/`, `in-progress/`, `done/`. Er wechselt nur durch `git mv`
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).

**Bezug:** [`BEO-008`](../../../../docs/plan/planning/observations/BEO-PLAN/verweis-auf-wandernden-slice/observation.md) — bei **3×**, Schwelle überschritten.

**Berührte Spec-Stellen:** — *(keine)*

**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.

**Autor:** Claude (Opus 5), im Auftrag des Maintainers. **Datum:** 2026-08-30.

---

## 1. Ziel

Ein `git mv` durch die Lifecycle-Verzeichnisse zieht die Verweise **auf** die bewegte Datei im
selben Schritt nach, statt sie dem nächsten Gate-Lauf zu überlassen.

## 2. Definition of Done

- [x] Ein Target bewegt einen Slice und zieht **beide** im Bestand vorkommenden Verweis-Formen
      nach — `../<verzeichnis>/<datei>` (26 Fundstellen) und
      `docs/plan/planning/<verzeichnis>/<datei>` (2) —, **repo-weit**, nicht nur unter
      `docs/plan/planning/`: `docs/reviews/` trägt gemessen ebenfalls solche Verweise.
- [x] Ein Selbsttest belegt die Ersetzung in beiden Richtungen: die zu treffende Form wird
      getroffen, eine **zitierte** oder fremde Form nicht.
- [x] Workflow-Skelett (Schritt 9) und [`AGENTS.md`](../../../../AGENTS.md) §5 nennen das Target
      statt der Merkregel.

- [x] `make gates` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.
- [x] Closure-Notiz mit benanntem Lerneintrag geschrieben (§7).
- [x] Beobachtungs-Register fortgeschrieben.
- [x] Jedes Risiko aus §6 trägt genau einen Ausgang.

## 3. Plan (vor Code)

| Datei | Änderungs-Art | Begründung |
|---|---|---|
| `tools/slice-mv.sh` | neu | die Bewegung samt Nachzug, mit Selbsttest |
| [`Makefile`](../../../../Makefile) | update | Target + `.PHONY` |
| [`AGENTS.md`](../../../../AGENTS.md) §4/§5, [`harness/README.md`](../../../../harness/README.md) | update | Deklaration; `doc-targets` erzwingt sie |
| `.claude/commands/slice.md` | update | Schritt 9 nennt das Target |

**Auszuführende Gates:** `make gates` — tragend `doc-targets` (beide Doku-Tabellen) und
`guard-selftest` (die GATES-Liste des Guard). Zum Abschluss `make verify`.

### Warum ein Werkzeug und nicht ein Guide

**Der Guide ist hier gemessen gescheitert.** [`SL-002`](../../../../docs/plan/planning/observations/README.md) zählt neun Vorfälle,
**zwei davon nach** dem Guide in Schritt 9 des Workflow-Skeletts. `BEO-008` ist die
Schwester-Klasse — Verweise **auf** einen wandernden Slice statt **in** ihm — und steht nach
[slice-116](../done/slice-116-nullmengen-haerte-cr.md) bei 3×. Die Baseline verlangt ab dem
dritten Vorfall einen Guide **oder Sensor**; wo ein Guide bereits einmal an derselben Familie
versagt hat, ist er nicht die stärkere Antwort.

### Warum ein Skript, wo slice-080 gerade 509 Zeilen abgebaut hat

Die abgelösten Skripte waren **Prüfungen** — und eine Prüfung kann ein Modul übernehmen. Dies ist
eine **Bewegung**. `d-check` ist read-only (`DC-QA-03`) und wird nie eine Datei verschieben; es
gibt kein Modul, das dieser Aufgabe entgegenkäme. Die Richtung „verteilen statt kopieren" ist
darum nicht verletzt, sondern nicht anwendbar.

## 4. Trigger

**Start:** eingetreten — `BEO-008` bei 3×.

**Rückführungen:**

- `in-progress` → `next`: falls sich zeigt, dass die Ersetzung Verweise trifft, die **nicht** die
  bewegte Datei meinen. Dann ist der Schnitt zu grob und braucht eine engere Bindung.

## 5. Closure-Trigger

Target läuft, Selbsttest feuert in beide Richtungen, Deklarationen nachgezogen, Gates grün.

**Was bewusst nicht getan wird:** die **Welle-Plan-Dateien**. Sie wechseln beim Closure-`mv` die
Verzeichnis**tiefe** (flach → `done/`), und ein Pfad aus Tiefe *n* braucht aus *n+1* ein
zusätzliches `../` — das ist eine andere Ersetzung als der Verzeichnis-Tausch auf gleicher Ebene
([slice-089](../done/slice-089-welle-datei-verweis-invariante.md)). Sie hier mitzunehmen hieße,
zwei Regeln unter einen Namen zu legen.

## 6. Risiken und offene Punkte

- *Ein `sed` über Dateipfade trifft auch Vorkommen in Code-Blöcken oder Backticks* — **Ausgang:**
  gestrichen mit Begründung: anders als bei einem **Sensor** ist das hier richtig. Ein zitierter
  Pfad, der auf die bewegte Datei zeigt, ist nach dem `mv` genauso falsch wie ein Link; der
  Selbsttest führt die zitierte Form darum als **Treffer**, nicht als Gegenprobe.
- *Das Werkzeug ersetzt die Aufmerksamkeit nicht, es verschiebt sie* — **Ausgang:** weiter offen im
  **Beobachtungs-Register**: wer `git mv` von Hand fährt, hat denselben Bruch wie vorher. Der
  Unterschied ist, dass der richtige Weg jetzt **kürzer** ist als der falsche, nicht länger.

## 7. Closure-Notiz

**Geliefert:** `make slice-mv SLICE=<slice-NNN> TO=<zustand>` — der `git mv` samt Nachzug der
Verweise **auf** die bewegte Datei, repo-weit und in beiden im Bestand gemessenen Formen. Vier
Fehlerfälle greifen (fehlende Argumente, fremdes Zielverzeichnis, unbekannter Slice, schon dort),
der Selbsttest fährt fünf Fixtures. Workflow-Skelett Schritt 9,
[`AGENTS.md`](../../../../AGENTS.md) §4/§5 und die `NICHT_PRUEFEND`-Liste des Guard sind nachgezogen.

**Lerneintrag — Form: geschärfte Regel.** *Wo ein Guide an einer Fehler-Familie schon einmal
versagt hat, ist die Antwort auf den dritten Vorfall ein Werkzeug — kein zweiter Guide.* Die
Baseline lässt bei überschrittener Schwelle beides zu („Guide **oder** Sensor"), und die Wahl
wirkt beliebig, bis man die Trefferquote danebenlegt: [`SL-002`](../../../../docs/plan/planning/observations/README.md) — dieselbe
Familie, nur die Gegenrichtung — zählt neun Vorfälle, **zwei davon nach** dem Guide in Schritt 9.
Ein Guide, der einmal überlesen wurde, wird zweimal überlesen. *Weil* das so ist, ist die Frage
bei der Schwelle nicht „Guide oder Sensor?", sondern „hat diese Familie schon einen Guide, und hat
er getragen?". Die stärkere Antwort ist nicht die aufwendigere, sondern die, die den **richtigen
Weg kürzer macht als den falschen**: `make slice-mv SLICE=… TO=…` ist weniger zu tippen als
`git mv` plus `grep -rl` plus `sed`.

**Drei beobachtbare Closure-Kriterien:**

1. Der Selbsttest fährt **beide** Richtungen: zwei Fixtures müssen ersetzt werden (Link-Form und
   zitierte Form), drei dürfen es **nicht** (andere Datei, anderes Verzeichnis, präfixlose
   Nicht-Form). Ohne die drei Gegenproben wäre eine Ersetzung, die alles umschreibt, von einer
   korrekten nicht zu unterscheiden.
2. Der **Drift-Wächter des Guard hat sofort gegriffen** — `guard-selftest` meldete
   *„Pruef-Target 'slice-mv' fehlt in der GATES-Liste"*, bevor das Target je gelaufen war. Das ist
   sein zweiter echter Einsatz nach slice-060, und beide Male hat er einen Zustand gemeldet, den
   niemand gesucht hat. Die Auflösung war nicht Aufnahme, sondern `NICHT_PRUEFEND` **mit
   Begründung**: der Exit-Code sagt „Bewegung geglückt", nicht „Bestand in Ordnung".
3. Dieser Slice ist selbst mit dem Werkzeug nach `done/` gewandert — der Beleg steht im
   Lifecycle-Commit.

**Offene Risiken und ihr Ausgang:** der erste gestrichen mit Begründung, der zweite weiter offen
im Register.

**Beobachtungs-Register:** [`BEO-008`](../../../../docs/plan/planning/observations/BEO-PLAN/verweis-auf-wandernden-slice/observation.md) ist **verkörpert** in
`make slice-mv` — die Beobachtung bleibt mit ihrem Zähler stehen, ihr Stand nennt jetzt den Ort.
Neu ist nichts: der zweite Risiko-Ausgang („das Werkzeug ersetzt die Aufmerksamkeit nicht") ist
dieselbe Beobachtung, nicht eine weitere.

**Folge-Slices:** [slice-117](../done/slice-117-handbuch-verweis-cli.md) (Handbuch-Verweis) ist
startbar; der Guide für CR-Texte aus [`BEO-022`](../../../../docs/plan/planning/observations/BEO-GATE/cr-text-behauptet-statt-gemessen/observation.md) steht weiter aus.
## 8. Sub-Area-Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** berührt werden die **Gate-/Werkzeug-Schicht** (`tools/`,
`Makefile`) und der **Harness-Einstieg** (zwei Deklarations-Tabellen, Workflow-Skelett).

**Vorgelagert — offene Beobachtungen sichten:** [`BEO-008`](../../../../docs/plan/planning/observations/BEO-PLAN/verweis-auf-wandernden-slice/observation.md) ist der Anlass
(3×). [`BEO-022`](../../../../docs/plan/planning/observations/BEO-GATE/cr-text-behauptet-statt-gemessen/observation.md) liegt in derselben Schicht und bleibt offen — dieser Slice
schreibt keinen CR.

Alle berührten Sub-Areas GF.
