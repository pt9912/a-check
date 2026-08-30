# slice-129 — Risiko-Ausgänge fallen vor dem `mv` auf, nicht danach

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese Datei liegt — eines von
`open/`, `next/`, `in-progress/`, `done/`. Er wechselt nur durch `make slice-mv`
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).

**Bezug:** [`BEO-006`](../observations.md) — bei **3×**, Schwelle überschritten.

**Berührte Spec-Stellen:** — *(keine)*

**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.

**Autor:** Claude (Opus 5), im Auftrag des Maintainers. **Datum:** 2026-08-30.

---

## 1. Ziel

Ein Formfehler in den Risiko-Ausgängen fällt beim Abschluss auf — nicht erst, wenn der **nächste**
Slice `make verify` fährt.

## 2. Definition of Done

- [x] `verify-risiko-ausgaenge` prüft **auch** `in-progress/` — aber nur Slices, deren
      Closure-Notiz **ausgefüllt** ist. Ein Slice mit Vorlagen-Platzhalter ist in Arbeit und darf
      nicht beanstandet werden; das ist die Bedingung, an der die Erweiterung hängt.
- [x] Der Selbsttest deckt **beide** Richtungen der neuen Bedingung: eine ausgefüllte Notiz in
      `in-progress/` mit fehlerhaftem Ausgang wird gemeldet, eine unausgefüllte **nicht**.
- [x] [`AGENTS.md`](../../../../AGENTS.md) §4 und
      [`harness/README.md`](../../../../harness/README.md) §Sensors nennen den erweiterten
      Geltungsbereich; das Workflow-Skelett sagt, was das für Schritt 8 bedeutet.

- [x] `make gates` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.
- [x] Closure-Notiz mit benanntem Lerneintrag geschrieben (§7).
- [x] Beobachtungs-Register fortgeschrieben.
- [x] Jedes Risiko aus §6 trägt genau einen Ausgang.

## 3. Plan (vor Code)

| Datei | Änderungs-Art | Begründung |
|---|---|---|
| `tools/verify-risiko-ausgaenge.sh` | update | Geltungsbereich + Selbsttest |
| [`AGENTS.md`](../../../../AGENTS.md) §4, [`harness/README.md`](../../../../harness/README.md) | update | der Vertrag ändert sich; `doc-targets` erzwingt die Deklaration |
| `.claude/commands/slice.md` | update | Schritt 8 fängt jetzt, was Schritt 9 verschob |

**Auszuführende Gates:** `make verify` (der Gegenstand), `make gates` — tragend `doc-targets`.

### Der Schaden ist zweifach belegt, nicht einfach

`BEO-006` zählt drei Vorfälle, und der zweite Schaden ist der teurere:

1. **Später Befund.** Der Formfehler von
   [slice-121](../done/slice-121-port-richtung-inbound-outbound.md) fiel erst auf, als
   [slice-122](../done/slice-122-planning-modul-wirksam.md) `make verify` fuhr — slice-121 lag da
   schon in `done/`. Derselbe Fehler wiederholte sich an slice-122 selbst.
2. **§3.3-Verstoß als Folge.** Wer nach dem `mv` noch am Inhalt arbeiten muss, erzeugt genau den
   Commit, den die Hard Rule verbietet: der Lifecycle-Commit von slice-122 zeigt **Rename 85 %**
   statt 100 %. Der Sensor zwingt zu einem Regelbruch, um seinen eigenen Befund zu beheben.

### Warum die Bedingung „ausgefüllt" trägt und nicht der Ort

Ein Slice in `in-progress/` ist **in Arbeit** — seine Closure-Notiz trägt am Anfang den
Vorlagen-Platzhalter, und ihn zu beanstanden hieße, jeden laufenden Slice rot zu melden. Der
richtige Auslöser ist deshalb nicht das Verzeichnis, sondern der **Zustand der Notiz**: sobald sie
ausgefüllt ist, ist der Slice abschlussbereit und die Prüfung fällig. Das ist genau der Moment vor
Schritt 9.

### Was NICHT mitwandert

`doc-structure` (Regeln 2 und 3, die Closure-**Struktur**) bleibt auf `done/`. Sein `files`-Glob
kann die Bedingung „nur wenn ausgefüllt" nicht ausdrücken — ein Glob über `in-progress/` würde
jeden laufenden Slice mit Platzhalter beanstanden. Die Hälfte von `BEO-006`, die dort liegt,
bleibt damit offen; das gehört benannt, nicht stillschweigend mitgenommen.

## 4. Trigger

**Start:** eingetreten — `BEO-006` bei 3×.

**Rückführungen:**

- `in-progress` → `open`: falls sich zeigt, dass „ausgefüllt" nicht mechanisch von „Platzhalter"
  zu trennen ist. Dann ist die Bedingung ein Urteil, und ein Sensor ist die falsche Antwort.

## 5. Closure-Trigger

Sensor erweitert, Selbsttest in beide Richtungen, Deklarationen nachgezogen, Gates grün — **und
dieser Slice selbst ist der erste Beleg**: sein eigener Abschluss läuft durch die neue Prüfung,
bevor er wandert.

**Was bewusst nicht getan wird:** `doc-structure` auf `in-progress/` erweitern (siehe §3) und die
Closure-Notiz **inhaltlich** prüfen — das bleibt beim Skill
[`closure-note-reviewer`](../../../../.harness/skills/closure-note-reviewer.md).

## 6. Risiken und offene Punkte

- *Die Platzhalter-Erkennung könnte eine ausgefüllte Notiz für unausgefüllt halten und still
  durchlassen* — **Ausgang:** gestrichen mit Begründung: der Selbsttest führt **beide**
  Richtungen als eigene Fixtures — Platzhalter muss als „in Arbeit" erkannt werden, ausgefüllte
  Notiz **nicht**. Die zweite ist die wichtigere: eine Erkennung, die alles für offen hält, würde
  jede Prüfung still überspringen und wäre von einer korrekten nicht zu unterscheiden.
- *Die Prüfung läuft ab jetzt zweimal über denselben Slice (in `in-progress/` und in `done/`)* —
  **Ausgang:** gestrichen mit Begründung: das ist keine Doppelung, sondern der Zweck. Der erste
  Lauf fängt den Fehler **vor** dem `mv`, der zweite hält den Bestand. Kosten: ein Dateilauf pro
  Slice, einmalig.

## 7. Closure-Notiz

**Geliefert:** `verify-risiko-ausgaenge` prüft `in-progress/` mit, sobald die Closure-Notiz dort
ausgefüllt ist. Die Meldung nennt die übersprungenen sichtbar (*„1 in Arbeit übersprungen"*), der
Selbsttest deckt beide Richtungen der neuen Bedingung, und
[`AGENTS.md`](../../../../AGENTS.md) §4 sowie das Workflow-Skelett sagen, was das für Schritt 8
bedeutet.

**Lerneintrag — Form: geschärfte Regel.** *Ein Sensor, dessen Geltungsbereich am **Ort** hängt,
prüft zum falschen Zeitpunkt, wenn der Ort erst durch die Arbeit entsteht.* Der Sensor sah
`done/` — und dorthin kommt ein Slice per `git mv`, also **nach** dem Moment, in dem seine
Closure-Notiz fertig ist. Der richtige Auslöser ist nicht das Verzeichnis, sondern der **Zustand**
der Notiz: sie trägt den Vorlagen-Platzhalter, solange gearbeitet wird, und ist ausgefüllt, wenn
der Slice abschlussbereit ist. *Weil* dieser Zustand mechanisch erkennbar ist, konnte der Sensor
den Ort verlassen; wäre er es nicht, wäre ein Guide die einzige Antwort gewesen.

**Der zweite Schaden war der teurere, und er ist der eigentliche Grund.** Ein später Befund kostet
eine Korrektur. Aber wer nach dem `mv` noch am Inhalt arbeiten muss, erzeugt genau den Commit, den
[`AGENTS.md`](../../../../AGENTS.md) §3.3 verbietet — der Lifecycle-Commit von
[slice-122](../done/slice-122-planning-modul-wirksam.md) zeigt **Rename 85 %** statt 100 %. Der
Sensor zwang zu einem Regelbruch, um seinen eigenen Befund zu beheben. Das ist keine Unbequemlichkeit,
sondern ein Widerspruch im Regelwerk selbst.

**Drei beobachtbare Closure-Kriterien:**

1. **Die Probe ist an diesem Slice selbst gefahren:** ein fehlerhafter Ausgang in seiner eigenen
   §6 — bei ausgefüllter Notiz, noch in `in-progress/` — wurde gemeldet (Exit 1, Datei genannt);
   zurückgebaut ⇒ Exit 0. Vor slice-129 wäre derselbe Fehler erst nach dem `mv` aufgefallen.
2. Das Überspringen ist **sichtbar**: die Schluss-Meldung nennt die Zahl der in Arbeit
   befindlichen Slices. Ein Übersprung, den niemand sieht, wäre von einer Prüfung, die nichts
   findet, nicht zu unterscheiden — dieselbe Klasse wie
   [`BEO-023`](../observations.md).
3. Die Grundgesamtheits-Grenze zählt **beide** Verzeichnisse: 126 geprüft, 1 übersprungen.

**Was offen bleibt, und warum es nicht mitwandert:** `doc-structure` (die Closure-**Struktur**)
prüft weiter nur `done/`. Sein `files`-Glob kann die Bedingung „nur wenn ausgefüllt" nicht
ausdrücken — ein Glob über `in-progress/` beanstandete jeden laufenden Slice. Diese Hälfte von
[`BEO-006`](../observations.md) bleibt bestehen; sie wäre erst mit einer Modul-Fähigkeit lösbar,
die es heute nicht gibt.

**Offene Risiken und ihr Ausgang:** beide gestrichen mit Begründung.

**Beobachtungs-Register:** [`BEO-006`](../observations.md) ist **verkörpert** für die
Risiko-Ausgänge; der Zähler bleibt bei 3×, sein Stand nennt jetzt den Ort **und** die verbliebene
Hälfte (`doc-structure`).

**Folge-Slices:** keiner zwingend. Der nächste naheliegende Schritt ist
[`BEO-026`](../observations.md) — das Modul `workflows` ist seit `v0.67.0` im Pin verfügbar und
nie konfiguriert; seit [slice-128](../done/slice-128-dependabot-hebungskanal.md) erzeugt der
Hebungs-Kanal genau die `uses:`-Hebungen, deren Form dort ungeprüft bleibt.
## 8. Sub-Area-Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** berührt wird die **Gate-/Werkzeug-Schicht** (`tools/`) und
mit den zwei Deklarations-Tabellen und dem Workflow-Skelett der **Harness-Einstieg**.

**Vorgelagert — offene Beobachtungen sichten:** [`BEO-006`](../observations.md) ist der Anlass
(3×). [`BEO-025`](../observations.md) (die Dreier-Menge kennt kein „eingetreten und so gewollt")
liegt daneben und bleibt offen — dieser Slice ändert **wann** geprüft wird, nicht **was**.

Alle berührten Sub-Areas GF.
