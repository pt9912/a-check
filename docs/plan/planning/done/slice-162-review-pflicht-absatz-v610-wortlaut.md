# slice-162 — Review-Pflicht/Rollenwechsel-Absatz in `AGENTS.md` §6 auf `v6.1.0`-Wortlaut zurückschneiden

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `make slice-mv` ([`AGENTS.md`](../../../../AGENTS.md)
§3.3/§5).

**Welle:** [welle-14](../welle-14-regelwerk-v610-migration.md).

**Bezug:** [slice-161](../done/slice-161-regelwerk-v610-delta-analyse.md)
§4.4/§6 (Folge-Slice-Vorschlag), Maintainer-Wort 2026-09-05 ("ja
umsetzen").

**Berührte Spec-Stellen:** — *(keine)* — Harness-/Konventions-Änderung
ohne Vertragsberührung.

**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim
Maintainer.

**Autor:** Claude (Sonnet 5). **Datum:** 2026-09-05.

---

## 1. Ziel

Den in `AGENTS.md` §6 ohne `MR`-Deklaration gewachsenen
Review-Pflicht/Rollenwechsel-Absatz auf den `v6.1.0`-Baseline-Wortlaut
zurückschneiden — nur zwei a-check-spezifische Zusätze bleiben, die der
Baseline-Wortlaut nicht trägt — und die verbleibende Differenz als neuen
`MR`-Eintrag dokumentieren.

## 2. Definition of Done

- [x] `AGENTS.md` §6, Absatz nach Schritt 8: Wortlaut deckungsgleich mit
      `lab/templates/AGENTS.template.md`s `v6.1.0`-Zuwachs (Kurs-Repo
      `pt9912/ai-harness-course`, Tag `v6.1.0`), erweitert um genau zwei
      a-check-spezifische Sätze — den `fork`-Ausschluss (Modul 8
      §Kontext-Trennung) und den Zitat-Anker auf
      `BEO-HARNESS/behauptete-vollstaendigkeit-extern-gefangen` (ohne die
      bisherige Chronik-Prosa, die `git` bereits hält, §3.7).
- [x] Neuer [`MR-018`](../../../../harness/conventions.md#mr-018) in
      `harness/conventions.md` (Adaptions-Block + eigene Datei
      [`harness/conventions/MR-018-review-pflicht-v610-wortlaut.md`](../../../../harness/conventions/MR-018-review-pflicht-v610-wortlaut.md))
      dokumentiert die verbleibende Differenz zum Baseline-Wortlaut, mit
      `Ersetzt-Baseline-Regel` → `.harness/baseline/v6.0.0/regelwerk/modul-08-agentenrollen.md#die-neun-übergaben-und-ihre-artefakte-modul-8`.
- [x] `make gates` grün.
- [x] `make verify` grün.
- [ ] Beobachtungs-Register fortgeschrieben.
- [ ] Jedes Risiko trägt einen Ausgang.

## 3. Plan (vor Code)

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| `AGENTS.md` §6 | update | Absatz kürzen (§1) |
| `harness/conventions/MR-<NNN>-review-pflicht-v610-wortlaut.md` | neu | Adaption dokumentieren |
| `harness/conventions.md` §Aktive Adaptionen | update | Tabellenzeile ergänzen |

## 4. Trigger

**Start** (`next` → `in-progress`): direkt eröffnet — Maintainer-Wort
liegt vor, slice-161 ist in `done/`, WIP-Limit frei.

**Rückführungen — vorab benennen:**

- `in-progress` → `next` (zu groß, zurück zur Zerlegung): falls sich beim
  Formulieren zeigt, dass mehr als drei Liefer-Punkte nötig sind (z. B.
  weil weitere `AGENTS.md`-Stellen dieselbe Drift zeigen).
- `in-progress` → `open` (blockiert): falls der Wortlaut-Vergleich mit
  `v6.1.0` eine inhaltliche Lücke aufdeckt, die eine Maintainer-Klärung
  vor der Umsetzung braucht.

## 5. Closure-Trigger

`AGENTS.md` §6 trägt den gekürzten Absatz, der neue `MR`-Eintrag steht mit
`Ersetzt-Baseline-Regel`, `make gates` und `make verify` je Exit 0
(Ausgabe in Datei, Exit-Code getrennt geprüft), Closure-Notiz geschrieben.

## 6. Risiken und offene Punkte

- *Die Kürzung entfernt mechanische Details (Report-Dateiname-Muster,
  HIGH-Verifikations-Pflicht), die bislang auch in `AGENTS.md` selbst
  standen* — **Ausgang:** gestrichen mit Begründung: das Dateiname-Muster
  (`docs/reviews/<YYYY-MM-DD>-<slice-oder-diff-ref>.md`) steht bereits in
  [`review-report.template.md`](../../../../.harness/baseline/v6.0.0/templates/docs/reviews/review-report.template.md)
  (verlinkt aus
  [`docs/reviews/README.md`](../../../reviews/README.md)), die
  HIGH-Verifikationspflicht in
  [`.harness/skills/reviewer.md`](../../../../.harness/skills/reviewer.md)
  (`HIGH-Findings werden vor Übernahme adversarisch … verifiziert`) —
  reine Dedublizierung, kein Informationsverlust, und im Einklang mit
  `AGENTS.md`s eigener Ziel-Form
  ("sie trägt Hard Rules und Pointer … sie dupliziert deren Inhalt
  nicht", §1).
- *Ein künftiges Re-Vendoring auf `v6.1.0` (Etappe A aus slice-161 §6)
  könnte denselben Absatz nochmals berühren* — **Ausgang:** gestrichen mit
  Begründung: dieser Slice übernimmt bereits den `v6.1.0`-Wortlaut: ein
  künftiges Re-Vendoring vendort denselben Text, den dieser Slice schon
  trägt — kein Widerspruch möglich, also kein Risiko im Sinn der
  Dreier-Menge.

## 7. Closure-Notiz

- **Was hat funktioniert:** die unabhängige Wortlaut-Gegenrechnung gegen
  einen frischen Klon von `pt9912/ai-harness-course` (Tag `v6.1.0`) hat
  den neuen `AGENTS.md`-Absatz als Satz-für-Satz-deckungsgleich mit dem
  Baseline-Zuwachs bestätigt, plus exakt den zwei benannten a-check-Sätzen
  — keine stillen Abweichungen.
- **Was ging anders als geplant:** das Review (F-1) und eine direkte
  Nachfrage des Maintainers deckten eine falsche `Ersetzt-Baseline-Regel`
  auf — `modul-08-agentenrollen.md` ist zwischen `v6.0.0` und `v6.1.0`
  unverändert und war nie der Wortlaut-Treiber; der ist ein **Template**
  (`AGENTS.template.md`, `v6.1.0`, nicht vendored), kein Regelwerk-Modul.
  [`MR-018`](../../../../harness/conventions.md#mr-018) dafür korrigiert
  (jetzt: kein Regel-Ersatz, sondern Provenienz-Korrektur wie
  [`MR-017`](../../../../harness/conventions.md#mr-017)).
- **Lerneintrag — Form: geschärfte Regel.** *Ein `Ersetzt-Baseline-Regel`-Feld
  darf nicht auf ein Regelwerk-Modul zeigen, nur weil es thematisch passt
  und vendored ist — es muss das Modul sein, dessen Text sich zwischen den
  beiden Baseline-Ständen tatsächlich geändert hat. Ändert sich stattdessen
  nur ein Template (oder gar keine Regelwerk-Datei), ist der Eintrag keine
  klassische Regel-Adaption, sondern eine Provenienz-/Wortlaut-Korrektur —
  derselbe Fall wie [`MR-017`](../../../../harness/conventions.md#mr-017) —
  und trägt das ehrlich als
  `Ersetzt-Baseline-Regel: —` mit Begründung, statt einen unveränderten
  Modul-Anker als Beleg vorzutäuschen.* *Weil* zwei unabhängige Prüfungen
  (Review F-2, Maintainer-Nachfrage) genau diese Verwechslung fanden, bevor
  sie in den Adaptions-Block einlief.
- **Beobachtungs-Register (`../observations/`):**
  [`BEO-HARNESS/adaption-korrigiert-repo-aussage`](../observations/BEO-HARNESS/adaption-korrigiert-repo-aussage/observation.md)
  — `evidence/slice-162.md` ergänzt, Zähler steht bei 2× (neben
  slice-097; die verwandte Kette
  [`MR-007`](../../../../harness/conventions/done/MR-007-adr-vorlagen-version.md)→[`013`](../../../../harness/conventions/done/MR-013-adr-vorlagen-version.md)→[`017`](../../../../harness/conventions.md#mr-017)
  steht separat unter
  [`rueckbau-kandidat-ueberlebt-baseline-migration`](../observations/BEO-HARNESS/rueckbau-kandidat-ueberlebt-baseline-migration/observation.md)).
  Unter der 3×-Schwelle, aber näher dran.
- **Folge-Slices:** keine neue ID — die offene Entscheidung ist kein
  Folge-Slice, sondern eine Etappen-Wahl für `welle-14`: Etappe A
  (Vendoring, slice-161 §6) direkt umsetzen, oder zuerst einen
  vollständigen Adaptions-Durchgang (Etappe B: alle 18 aktiven
  `MR`-Einträge gegen `v6.1.0` bewerten, nicht nur die vom Diff berührten)
  fahren — offen für den Maintainer, gestellt am Ende dieser Sitzung.
- **Risiken aus §6:** beide mit Ausgang (beide gestrichen mit Begründung)
  — siehe §6.
- **Drei Paarungen:** verschoben auf die Closure von `welle-14` (Modul 8
  §Rollen-Sequenz für eine Welle) — wie bei slice-161.

## 8. Sub-Area-Modus

**Vorgelagert — Sub-Area-Wahl prüfen:** eine Sub-Area berührt —
**Harness-Einstieg** (`AGENTS.md`) — Greenfield, erfüllt die Schwelle
≥ 2/3 laut
[`harness/conventions.md`](../../../../harness/conventions.md#modus-deklaration-pro-sub-area).

**Vorgelagert — offene Beobachtungen sichten:** Register durchgegangen —
sieben `offen`-Einträge zu Harness-Einstieg (vgl.
[slice-161](../done/slice-161-regelwerk-v610-delta-analyse.md) §10),
keiner davon zum Gegenstand dieses Slice (Review-Pflicht-Wortlaut); die
neue Beobachtung
[`BEO-PLAN/slice-provenienz-aus-gedaechtnis-statt-git-log`](../observations/BEO-PLAN/slice-provenienz-aus-gedaechtnis-statt-git-log/observation.md)
(1×) ist Planungs-Harness, nicht Harness-Einstieg, und ebenfalls nicht
Gegenstand.

**Alle berührten Sub-Areas GF** — kein Begründungsblock nötig.
