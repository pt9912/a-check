# slice-183 — `AGENTS.md` verweist, statt Regelwerk-Normtext nachzuschreiben

**Welle:** ohne Welle — die Closure-Bedingung wäre die DoD dieses Slice
(Baseline-Regelwerk `modul-06-roadmap.md` §Wann Arbeit eine Welle braucht).

**Bezug:** [`AC-QA-02`](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze) (Harness-Integrität), keine aktive ADR.

**Berührte Spec-Stellen:** — (der Slice berührt Rang 8 der Source Precedence).

**Verantwortlich:** Claude — gesetzt beim Übergang nach `in-progress/`.

**Autor:** Claude. **Datum:** 2026-09-08.

**Lerneintrag — Form:** geschärfte Regel.

---

## 1. Ziel und Abgrenzung

**Ziel:** [`AGENTS.md`](../../../../AGENTS.md) trägt für jede Regel, die im
vendorten Regelwerk steht, einen **Verweis** plus a-checks eigene Ausprägung —
nicht die nachgeschriebene Begründung der Baseline. Die Datei sagt das über sich
selbst (§1): *„sie trägt Hard Rules und Pointer auf kanonische Quellen, sie
dupliziert deren Inhalt nicht; sonst entsteht Drift."*

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Regeln streichen.** Der Slice kürzt **Herleitungen**, keine Zusagen. Eine
  Regel, die a-check bindet, bleibt — auch wenn sie im Regelwerk steht; was
  geht, ist ihre *Begründung*, wenn die dort schon ausformuliert ist.
- **[`harness/README.md`](../../../../harness/README.md).** Erledigt in
  [slice-182](../done/wellenlos/slice-182-readme-verweist-statt-wiederholt.md);
  ein zweiter Durchgang dort wäre Arbeit ohne Gegenstand.
- **[`harness/conventions.md`](../../../../harness/conventions.md).** Rang und
  Zweck sind andere: Sie ist der Konventionsspeicher, nicht das Briefing. Ein
  eigener Vorgang, wenn die Messung dort etwas zeigt.
- **Ein Sensor.** *„Schreibt dieser Satz das Regelwerk nach?"* ist ein Urteil
  über die Herkunft, kein Match ([`AGENTS.md`](../../../../AGENTS.md) §3.7).

## 2. Ausgangsmessung (2026-09-08)

**Methode** wie in slice-182: Abschnitt von einer `##`-Überschrift bis zur
nächsten, HTML-Kommentare und Tabellenzeilen entfernt, verbleibende Zeichen
gezählt. Gemessen gegen `v6.5.0` · `templates/AGENTS.template.md`.

| Abschnitt | a-check | Ziel-Form | Faktor |
|---|---|---|---|
| §1 Was diese Datei ist | 2211 | 2714 | 0,8 × |
| §2 Kanonische Quellen | 782 | 718 | 1,1 × |
| §3 Harte Regeln | 3760 | 3388 | 1,1 × |
| §4 Quality Gates | 1051 | 458 | 2,3 × |
| **§5 Dokumentations-Regeln** | **16 822** | **726** | **22,9 ×** |
| §6 Minimal Agent Workflow | 2889 | 874 | 3,3 × |

Datei gesamt **35 384** gegen **10 885** — beim Anlegen des Plans waren es
34 690 bzw. 15 464 in §5; **§5 ist zwischen Anlage und Umsetzung um 1358
Zeichen gewachsen**, durch die Regel, die slice-181 dort verkörpert hat. Der
Gegenstand wächst, während man ihn plant. Vier der sechs Abschnitte liegen bei
Faktor ≤ 2,3 — **§5 trägt die Hälfte der ganzen Datei.**

**Der Faktor allein ist kein Befund** (die Lehre aus slice-182). Zwei Dinge
sprechen hier trotzdem dafür, dass etwas dran ist:

1. **Die Ziel-Form §5 ist ausgeschrieben, nicht platzhaltend** — vier
   vollständige Punkte, 726 Zeichen. Bei §Safety in `harness/README.md` war der
   Faktor 22 × ein Artefakt zweier Platzhalter-Zeilen; hier ist er es nicht.
2. **Sechs Stichproben, sechs Treffer.** Substanz-Suche (nicht Wortsuche) im
   vendorten Regelwerk findet für jede geprüfte Regel eine Fundstelle:
   Größen-Regel · WIP-Limit · AC-Form · Closure-/Lerneintrag-Pflicht ·
   Lifecycle mit fünf Übergängen · Beobachtungs-Register in Pfadform. **Alle in
   `modul-05` bzw. `modul-06`.**

**Was das noch nicht sagt** — und was der Slice zu klären hat: Eine Regel zu
**nennen** und die a-check-Ausprägung dazuzuschreiben ist genau die Aufgabe
dieser Datei. Befund ist erst die nachgeschriebene **Begründung**. Beispiel für
beides im selben Block: *„WIP-Limit = 1 … Null ist zulässig"* ist a-checks
Ausprägung (die Baseline sagt „pro Rolleninhaber", a-check „pro Repo") — der
Nachsatz *„Bis slice-077 stand hier ‚genau ein'"* ist Chronik (§3.7).

**Die größten Blöcke in §5, als Kandidaten-Liste:**

| Block | Zeichen | erste Einschätzung |
|---|---|---|
| Slice-Form (sieben Kopier-Punkte) | 4285 | a-check-eigen — jeder Punkt **gegen den Bestand gemessen** (174 Dateien, 58 Stellen); steht nirgends sonst |
| Geltungsbereich einer Messung | 2051 | a-check-eigen (`seit slice-179`/`slice-182`) |
| Zitier-Form in einfrierenden Artefakten | 1057 | a-check-eigen (`seit slice-176`) |
| Beobachtungs-Register | 1002 | **Kandidat** — `modul-06` §Das Beobachtungs-Register trägt Form und Zählregel |
| CR-Texte an ein fremdes Werkzeug | 943 | a-check-eigen |
| Commit-Scope `(planning)` | 941 | a-check-eigen (gemessen: 5 von 74) |
| Slice-Lifecycle (fünf Übergänge) | 849 | **Kandidat** — die Tabelle steht in `modul-05` |
| Diskrepanz-Trichter | 627 | **Kandidat** — `modul-07` §Werkzeug-Wahl |

Die Einschätzungen sind **nicht** das Ergebnis — sie sind die Liste, die der
Slice je Block gegen das Regelwerk hält.

## 3. Plan (vor Code) — und was die Prüfung daraus machte

| Datei / Komponente | Änderungs-Art | Ergebnis |
|---|---|---|
| `AGENTS.md` §5, drei Blöcke | refactor | **gekürzt** — Slice-Lifecycle, Beobachtungs-Register, Diskrepanz-Trichter |
| `AGENTS.md` §5, die übrigen 14 Blöcke | *geprüft, unverändert* | a-check-eigen, jeder mit eigener Messung oder eigenem Anlass |
| `AGENTS.md` §6 | *geprüft, unverändert* | wo er die Baseline zitiert, tut er es bereits **mit Verweis** |
| `AGENTS.md` §4 | *geprüft, unverändert* | mit slice-181 bereits angefasst |
| `AGENTS.md` §1–§3 | *geprüft, unverändert* | Faktor ≤ 1,1 × — die Hard Rules sind a-checks eigene |

**Die drei gekürzten Blöcke, je mit dem Beleg:**

| Block | vorher | jetzt | steht im Regelwerk |
|---|---|---|---|
| Slice-Lifecycle | 849 | 449 | `modul-05` §Trigger je Lifecycle-Übergang — die Tabelle der fünf Übergänge wortgleich |
| Beobachtungs-Register | 1002 | 442 | `modul-06` §Das Beobachtungs-Register — Form, Zählregel, drei Ausgänge |
| Diskrepanz-Trichter | 627 | 316 | `modul-07` §Werkzeug-Wahl bei Diskrepanz |

**Was in jedem der drei blieb, weil es *nicht* im Regelwerk steht** — geprüft,
nicht angenommen: `make slice-mv` als Träger des Übergangs · dass der direkte
Weg `open/ → in-progress/` zulässig und `next/` keine Pflichtstation ist
(**null** Treffer im vendorten Regelwerk) · der Ort der Ablage und die Regel,
dass das Sub-Area-Kürzel **nachgeschlagen und nicht erfunden** wird · die
Ablageorte für Carveout und BF-Markierung.

## 3.1 Das Ergebnis widerlegt die Vermutung des Plans

§2 nannte zwei Gründe, warum der Faktor **21,3 ×** in §5 diesmal etwas heißen
könnte: die Ziel-Form ist dort ausgeschrieben statt platzhaltend, und sechs
Substanz-Stichproben fanden alle sechs eine Fundstelle im Regelwerk. **Beides
trifft zu und trägt trotzdem nicht.**

§5 schrumpft um **1095** Zeichen — rund **6,5 %**. Die großen Blöcke sind
a-check-eigen und tragen je eine eigene Messung oder einen eigenen Anlass:
die Kopieranleitung für Slices (4285; sieben Punkte, jeder gegen den Bestand
gemessen), *Geltungsbereich einer Messung* (2051), *Zitier-Form in
einfrierenden Artefakten* (1057), *CR-Texte an ein fremdes Werkzeug* (943),
*Commit-Scope `(planning)`* (941, mit eigener Messung: fünf Treffer bei 74
Commits).

**Warum die Stichproben in die Irre führten:** Sie fragten *„steht die Regel im
Regelwerk?"* — und das tut sie fast immer, denn a-check hat die Baseline
adoptiert. Die richtige Frage ist *„steht die **Begründung** dort, und schreibt
a-check sie nach?"*. Eine Regel zu **nennen** und ihre repo-eigene Ausprägung
danebenzustellen ist die Aufgabe dieser Datei; nur die nachgeschriebene
Herleitung ist der Befund.

Das ist die Regel aus slice-182 an ihrem eigenen Gegenstand bestätigt: **Der
Größen-Vergleich mit einer Ziel-Form ist kein Befund.** Diesmal hat sie in die
andere Richtung getragen — dort ließ sie einen Abschnitt mit Faktor 22 ×
stehen, hier einen mit 21,3 ×.

## 4. Definition of Done

- [ ] §5 ist **je Block** gegen das vendorte Regelwerk geprüft; jede Streichung
      und jedes Bleiben ist begründet.
- [ ] §4 und §6 ebenso.
- [ ] Die Messung ist nach der Umsetzung wiederholt und nennt ihren
      **Geltungsbereich** und ihre **Methode**.
- [ ] Unabhängiger Review durchgeführt (Report unter [`docs/reviews/`](../../../reviews/README.md)).
- [ ] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [ ] Beobachtungs-Register fortgeschrieben.
- [ ] Jedes Risiko aus §7 trägt einen Ausgang.

`make gates` und `make verify` grün. Die drei Paarungen trägt die Closure
dieses Slice selbst — wellenloser Betrieb. Ein öffentlicher Vertrag ist
berührt: `AGENTS.md` ist Rang 8, und `make doc-targets` prüft ihre §4 gegen das
Makefile.

## 5. Trigger

**Start** (`open` → `in-progress`): [slice-182](../done/wellenlos/slice-182-readme-verweist-statt-wiederholt.md)
liegt in `done/` (WIP-Limit 1).

**Rückführungen — vorab benannt:**

- `in-progress` → `next` (zu groß): Zeigt die Block-Prüfung, dass §5 allein mehr
  als drei Liefer-Punkte trägt, wird nach Abschnitten zerlegt.
- `in-progress` → `open` (blockiert): Stellt sich heraus, dass eine Regel **nur**
  in `AGENTS.md` steht und im Regelwerk fehlt, ist das eine benannte Spec-Lücke
  — sie wird eingetragen, und der Slice wartet auf ihre Auflösung.

## 6. Closure-Trigger

DoD vollständig, `make gates` und `make verify` grün, Closure-Notiz mit
Lerneintrag — und die **gemessene** Gegenprobe je geänderten Abschnitts.

## 7. Risiken und offene Punkte

- **Kürzen entfernt eine Zusage statt einer Herleitung.** `AGENTS.md` ist Rang 8
  und bindet jeden Lauf; ein verlorener Halbsatz ist teurer als in Rang 9.
  — **Ausgang:** *entfallen*, gestrichen mit Begründung. Für jeden der drei
  Blöcke ist geprüft, was bleibt und warum — und in einem Fall hat die Prüfung
  eine Aussage **gerettet**, die sonst mitgegangen wäre: *„`next/` ist keine
  Pflichtstation"* hat **null** Treffer im vendorten Regelwerk und steht nur
  hier. Das Risiko war der Grund, je Satz zu prüfen statt je Block.
- **Der Verweis auf einen Regelwerk-Abschnitt altert.** Dieselbe Grenze wie in
  slice-182: Inline-Code-Zeiger sieht kein Sensor.
  — **Ausgang:** *weiter offen* → Beobachtungs-Register,
  [`BEO-PLAN/ziel-form-tag-gescopt`](../observations/BEO-PLAN/ziel-form-tag-gescopt/observation.md).
  Der Slice fügt **drei** solche Zeiger hinzu (`modul-05`, `modul-06`,
  `modul-07`, je mit §-Abschnitt in Inline-Code). Sie brechen still, wenn die
  nächste Baseline einen Abschnitt umbenennt — dieselbe benannte Grenze wie bei
  den neun aus slice-182, und kein neuer Eintrag.
- **§5 ist zu groß für einen Slice.** 15 464 Zeichen in 17 Blöcken; die
  Rückführung nach `next/` ist vorab benannt, aber sie kostet einen Durchgang.
  — **Ausgang:** *entfallen*, gestrichen mit Begründung: Das Risiko setzte
  voraus, dass viel zu kürzen ist. Gemessen waren es **drei von 17** Blöcken und
  1095 Zeichen — die Rückführung wurde nie fällig. **Die Annahme hinter dem
  Risiko war der eigentliche Fehler**, nicht seine Größe.

## 8. Closure-Notiz

**Lerneintrag — Form: geschärfte Regel** (dritte Hälfte von *Geltungsbereich
einer Messung*, `AGENTS.md` §5).

- **Was hat funktioniert:** Je Block prüfen statt je Abschnitt. Die Frage
  *„steht die **Begründung** im Regelwerk?"* trennt sauber, wo *„steht die Regel
  dort?"* alles bejaht. Und einmal hat die Prüfung eine Aussage **gerettet**:
  *„`next/` ist keine Pflichtstation"* hat null Treffer im vendorten Regelwerk
  und wäre bei blockweisem Kürzen mitgegangen.

- **Was ging anders als geplant — die Vermutung war falsch, und das ist das
  Ergebnis.** §2 nannte zwei Gründe, warum der Faktor 22,9 × diesmal etwas
  heißen könnte: die Ziel-Form ist dort ausgeschrieben, und sechs von sechs
  Substanz-Stichproben fanden eine Fundstelle. Beides trifft zu. Gemessen sind
  **drei von 17** Blöcken Doppelung, **1095** Zeichen, **6,5 %**. Die großen
  Blöcke tragen je eine eigene Messung: die Kopieranleitung für Slices (4285,
  sieben Punkte gegen den Bestand gemessen), *Geltungsbereich einer Messung*
  (2051), *Zitier-Form* (1057), *CR-Texte* (943), *Commit-Scope* (941, fünf
  Treffer bei 74 Commits).

- **Die Zahl, die das am schärfsten sagt:** §5 ist nach diesem Slice **16 360**
  Zeichen groß — vorher 16 822. Netto **−462**, weil der Slice selbst 633
  Zeichen Regel dort verkörpert hat. **Ein Abschnitt, der die gelernten Regeln
  aufnimmt, wächst schneller als das Aufräumen ihn schrumpfen kann** — und das
  ist keine Schwäche, sondern seine Funktion. Wer §5 klein haben will, muss
  fragen, wo die Regeln sonst hin sollen, nicht wo sie doppelt stehen.

- **Steering-Loop-Eintrag — geschärfte Regel:** *„Steht die Regel im Regelwerk?"
  ist in einem Repo, das eine Baseline adoptiert hat, die falsche Frage — sie
  wird fast immer bejaht. Trägt nur: „schreibt dieser Absatz ihre Begründung
  nach?"* — liegt in [`AGENTS.md`](../../../../AGENTS.md) §5, *Geltungsbereich
  einer Messung*, dritte Hälfte (`seit slice-183` dort).
  Auslöser: [`BEO-HARNESS/baseline-normtext-nachgeschrieben`](../observations/BEO-HARNESS/baseline-normtext-nachgeschrieben/observation.md)
  (slice-103, slice-182, slice-183 — **3×**).

- **Beobachtungs-Register (`../observations/`):** ein Beleg,
  `baseline-normtext-nachgeschrieben` erreicht **3×**. Sein Ausgang ist
  ungewöhnlich und steht im `state.md`: **kein Sensor und keine neue Regel** —
  die Regel existiert bereits in beiden betroffenen Dateien selbst
  ([`AGENTS.md`](../../../../AGENTS.md) §1,
  [`harness/README.md`](../../../../harness/README.md) §Purpose). Was fehlte,
  war ihre Anwendung, und die ist ein Urteil (§3.7). Verkörpert ist stattdessen
  die **Prüf-Frage**, die die Klasse überhaupt erst findet.
  [`chronik-in-gelesenen-dateien`](../observations/BEO-HARNESS/chronik-in-gelesenen-dateien/observation.md)
  bekommt **keinen** Beleg — die drei Blöcke trugen Regel-Wiedergabe, nicht
  Chronik.

- **Folge-Slices:** keiner. Der Plan hatte
  [`harness/conventions.md`](../../../../harness/conventions.md) als möglichen
  eigenen Vorgang genannt; nach diesem Ergebnis ist er **nicht** angezeigt —
  dieselbe Messung dort würde denselben Anteil finden, und ein Slice mit 6,5 %
  Ertrag rechtfertigt sich nicht von selbst.

- **Risiken aus §7:** drei, jedes mit genau einem Ausgang — zweimal *entfallen*
  mit Begründung, einmal *weiter offen* → Register.

- **Drei Paarungen** (Repo ohne Wellen-Betrieb) — geprüft **nach** dem `git mv`
  nach `done/`, weil sie dort suchen; eingetragen im dritten Closure-Commit.

## 9. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** berührt ist **Harness-Einstieg**
(`HARNESS`; Achsen 1,2,3 laut
[`harness/conventions.md`](../../../../harness/conventions.md#modus-deklaration-pro-sub-area)).
Die Gate-/Werkzeug-Schicht ist **nicht** berührt: kein Target, kein Skript.
`AGENTS.md` §4 nennt Targets, ändert aber keines.

**Vorgelagert — offene Beobachtungen sichten:** gesichtet am 2026-09-08.
`BEO-HARNESS/` führt **14** Einträge, 10 davon `offen`. Zwei sind einschlägig:

- [`BEO-HARNESS/baseline-normtext-nachgeschrieben`](../observations/BEO-HARNESS/baseline-normtext-nachgeschrieben/observation.md)
  — **2×**, der Eintrag, den dieser Slice bedient. Er wird auf **3×** gehoben:
  drei Blöcke in `AGENTS.md` §5 gaben Baseline-Normtext wieder, und der Eintrag
  nennt `AGENTS.md` §1 seit slice-103 namentlich als offene Stelle
  (*„`AGENTS.md` §1 trägt dieselbe Klausel weiterhin im Rumpf"*).
- [`BEO-HARNESS/chronik-in-gelesenen-dateien`](../observations/BEO-HARNESS/chronik-in-gelesenen-dateien/observation.md)
  — **2×**, **kein Beleg**: Die drei gekürzten Blöcke trugen Regel-Wiedergabe,
  nicht Chronik. Der Register-Eintrag nennt `AGENTS.md` §5 zwar als Reststelle;
  die betrifft aber eine andere Stelle als die hier angefassten. Keine Treffer
  sind auch eine Antwort.

**Die übrigen zwölf** sind durchgegangen und nicht einschlägig: keiner betrifft
die Frage *„schreibt dieser Absatz das Regelwerk nach?"*.

**Alle berührten Sub-Areas GF** — kein Begründungsblock nötig.

