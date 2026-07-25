# slice-057 — Etappe F (1/3): Steering-Loop-Kanal und der erste Sensor daraus

**Status:** in-progress — erster Schnitt der **Etappe F (Betriebsmodell)** aus
[slice-048 §5](../done/slice-048-modul-delta-lesen.md).
**Deckt:** Fund **B-21** (kein Steering-Loop-Kanal) und den daraus fälligen Sensor zum
Pipe-Fehler, den [slice-051](../done/slice-051-workflow-und-freigabe.md) ausdrücklich offen ließ.
**Nicht hier:** B-13/B-14/B-10/B-9 (Wellen-Closure, Carveouts, Rollen) — 2/3 und 3/3.
[Roadmap](../in-progress/roadmap.md).

---

## 1. Warum der Kanal zuerst kommt

Etappe F hat fünf Funde. Vier davon sind Artefakte, die a-check fehlen; **B-21 ist der Kanal, der
die anderen vier gemeldet hätte**. Die Baseline setzt die Schwelle bei 3×: zweimal ist ein Symptom,
dreimal eine Lücke im Harness. Ohne einen Ort, an dem gezählt wird, bleibt jeder Vorfall ein
Einzelfall — und genau so ist es hier gelaufen: die Praxis existierte in slice-001…008 und schlief
danach **lückenlos** ein, 40 Slices ohne einen Eintrag.

Der Kanal entsteht als `docs/plan/steering-loop.md` und **nicht** in einer Wellen-Closure-Notiz,
wo die Baseline ihn verortet: a-check hat bis heute keine Welle auditierbar geschlossen (B-13,
offen in 3/3). Ein Kanal, der auf ein nicht existierendes Artefakt wartet, sammelt nichts.

## 2. Der erste Eintrag ist überfällig — und wird sofort beantwortet

**SL-001 — Gate-Lauf in einer Pipe verschluckt.** Fünf Vorfälle an einem Tag. `make <gate> | tail`
liefert den Exit-Code von `tail`; ein rotes Gate verschwindet spurlos, und in einem Fall ging ein
Commit mit rotem `doc-check` heraus. slice-051 hat dafür einen **Guide** geschrieben (Schritt 6 des
Workflow-Skeletts) — der fünfte Vorfall passierte **danach**. Nach Modul 09 ist eine Regel mit nur
einem Quadranten halb durchgesetzt; hier wird der zweite nachgereicht.

**Regel 2 im PreToolUse-Guard**: `make <gate> | …` und `make <gate> && git commit` werden
fail-closed abgelehnt, mit einer eigenen Begründung, die den richtigen Aufruf zeigt.

**SL-002 — brechende Verweise beim `git mv`.** Sieben Vorfälle, ebenfalls über der Schwelle.
Erfasst, **nicht** beantwortet: die wirksame Antwort wäre eine Prüfung, die *vor* dem Verschieben
meldet, welche Verweise brechen, und die gehört zu `make verify` statt in diesen Slice. Der Eintrag
hält die Zählung, damit die Lücke nicht ein achtes Mal als Einzelfall durchgeht.

## 3. Der Sensor hat sich beim Bau selbst geprüft

Zweimal feuerte die neue Regel auf ihren eigenen Autor:

1. **Berechtigt** — ein Prüfkommando enthielt eine echte Pipe auf `make gates`.
2. **Als Fehlalarm** — dasselbe Muster stand nur als *Zeichenkette* in einem Argument, nicht als
   Pipeline. Die bestehende Toolchain-Regel löst diesen Fall längst sauber („Toolname nur im
   Arg-String"); meine neue tat es nicht.

Der zweite Fall ist die Ursache der Quote-Behandlung. Ohne ihn wäre die Regel mit einer
Rausch-Quelle in Betrieb gegangen — und ein Sensor, der rauscht, wird abgeschaltet statt repariert.
Beim Nachziehen brach ich den Guard zusätzlich syntaktisch (einfache Anführungszeichen im Regex
beenden den umgebenden Bash-String), was **jeden** Bash-Aufruf blockierte, bis es behoben war.

## 4. Betroffene Module

- `docs/plan/steering-loop.md` — der Kanal (B-21).
- `.claude/hooks/pretooluse-command-guard.sh` — Regel 2 samt eigener Begründung und sieben neuen
  Selbsttest-Fällen (SL-001).
- [`AGENTS.md`](../../../../AGENTS.md) §5 — Verweis auf den Kanal.

Zwei Schichten (Planungs-Doku, Durchsetzungsschicht).

## 5. DoD

- [x] `docs/plan/steering-loop.md` existiert mit Schwellen-Regel, Pflege-Regeln und den zwei
      belegten Einträgen (Vorfallszahlen aus der Commit-Historie, nicht geschätzt) (B-21).
- [x] Der PreToolUse-Guard lehnt `make <gate> | …` und `make <gate> && git commit` ab; belegt durch
      13 Diskriminierungs-Proben (fünf müssen greifen, sechs dürfen nicht, zwei für die unveränderte
      Regel 1) und den erweiterten `--selftest` in `make gates` (SL-001).
- [x] `make gates` und `make verify` grün.

## 6. Closure-Notiz

_(beim Abschluss.)_

## 7. Sub-Area-Modus

Alle berührten Sub-Areas GF.
