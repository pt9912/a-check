# slice-183 — `AGENTS.md` verweist, statt Regelwerk-Normtext nachzuschreiben

**Welle:** ohne Welle — die Closure-Bedingung wäre die DoD dieses Slice
(Baseline-Regelwerk `modul-06-roadmap.md` §Wann Arbeit eine Welle braucht).

**Bezug:** [`AC-QA-02`](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze) (Harness-Integrität), keine aktive ADR.

**Berührte Spec-Stellen:** — (der Slice berührt Rang 8 der Source Precedence).

**Verantwortlich:** —

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
  [slice-182](../in-progress/slice-182-readme-verweist-statt-wiederholt.md);
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
| **§5 Dokumentations-Regeln** | **15 464** | **726** | **21,3 ×** |
| §6 Minimal Agent Workflow | 2889 | 874 | 3,3 × |

Datei gesamt **34 690** gegen **10 885**. Vier der sechs Abschnitte liegen bei
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

## 3. Plan (vor Code)

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| `AGENTS.md` §5 | refactor | trägt die Hälfte der Datei; je Block Herkunft prüfen |
| `AGENTS.md` §6, §4 | update | Faktor 3,3 × und 2,3 ×, dieselbe Frage im kleineren Maßstab |
| `AGENTS.md` §1–§3 | *(geprüft, voraussichtlich unverändert)* | Faktor ≤ 1,1 × — die Hard Rules sind a-checks eigene |

**Kriterium je Block** — dasselbe wie in slice-182: Steht die **Begründung** im
vendorten Regelwerk, wird sie ein Verweis. Ist der Satz a-checks Ausprägung
(Zahl, Stichtag, Grandfathering, Durchsetzungs-Target), bleibt er. Chronik geht
in jedem Fall (§3.7).

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

**Start** (`open` → `in-progress`): [slice-182](../in-progress/slice-182-readme-verweist-statt-wiederholt.md)
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
  — **Ausgang:** <offen bis Closure>
- **Der Verweis auf einen Regelwerk-Abschnitt altert.** Dieselbe Grenze wie in
  slice-182: Inline-Code-Zeiger sieht kein Sensor.
  — **Ausgang:** <offen bis Closure>
- **§5 ist zu groß für einen Slice.** 15 464 Zeichen in 17 Blöcken; die
  Rückführung nach `next/` ist vorab benannt, aber sie kostet einen Durchgang.
  — **Ausgang:** <offen bis Closure>

## 8. Closure-Notiz

*(bei Closure auszufüllen)*

## 9. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** berührt ist **Harness-Einstieg**
(`HARNESS`; Achsen 1,2,3 laut
[`harness/conventions.md`](../../../../harness/conventions.md#modus-deklaration-pro-sub-area)).
Die Gate-/Werkzeug-Schicht ist **nicht** berührt: kein Target, kein Skript.
`AGENTS.md` §4 nennt Targets, ändert aber keines.

**Vorgelagert — offene Beobachtungen sichten:** *(beim Übergang nach
`in-progress/` auszufüllen — der Register-Stand beim Anlegen ist ein anderer als
beim Beginn der Arbeit; die Sichtung gehört an den späteren Zeitpunkt, nicht als
Ankündigung hierher.)*

**Modus-Begründungsblock:** alle berührten Sub-Areas GF.
