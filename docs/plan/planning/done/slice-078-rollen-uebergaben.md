# slice-078 — Zwei fehlende Rollen-Übergaben: schließen oder deklarieren

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Deckt:** `F-11` aus dem
[Review-Report `welle-12`](../../../reviews/2026-08-09-welle-12-unabhaengig.md) (Gruppe C).
**Bezug:** [`harness/conventions.md`](../../../../harness/conventions.md) §Adaptions-Block;
[slice-066](../done/slice-066-wellen-closure-und-rollen.md); Roadmap-Zeile *Aktuelle Welle* in der
[Roadmap](../in-progress/roadmap.md).

---

## 1. Auslöser

**Mechanismus: eine bekannte Lücke, die sichtbar geführt, aber nicht deklariert ist.** Sie ist
weder verschwiegen noch falsch gemessen — sie hat nur keinen Ort, an dem sie als bewusste
Abweichung gilt.

[`harness/README.md:141`](../../../../harness/README.md) führt die Rollen-Übergaben und markiert
zwei davon ausdrücklich:

```
| Verifier → Validator | **fehlt** — siehe unten |
| Validator → Planner  | **fehlt** — siehe unten |
```

Regelwerk-Modul 8 verlangt jedes der neun Artefakte. Unmittelbar danach erklärt
[slice-066](../done/slice-066-wellen-closure-und-rollen.md) alle 21 Funde für geschlossen und die
Migration für vollständig — **ohne** dass für die beiden fehlenden Übergaben eine Adaption
deklariert wäre.

Das Repo hat für genau diesen Fall einen Mechanismus: den `MR-*`-Block in
[`harness/conventions.md`](../../../../harness/conventions.md), der neun bewusste Abweichungen von
der Baseline trägt. Eine zehnte fehlt hier — oder die Lücke wird geschlossen.

## 2. Betroffene Module

Zwei Schichten:

1. **[`harness/conventions.md`](../../../../harness/conventions.md)** — der `MR-*`-Block, falls
   die Abweichung deklariert wird.
2. **[`harness/README.md`](../../../../harness/README.md)** — die Übergaben-Tabelle, die dann auf
   die Deklaration zeigt statt auf „siehe unten".

## 3. Auszuführende Gates

`make gates` (enthält `doc-check` und `gate-consistency`), `make verify`.

**Offener Entscheid — vor dem Bau zu treffen.** Die Validator-Rolle ist in a-check nicht besetzt;
ob sie es werden soll, ist **nicht repo-intern beantwortbar** — es hängt daran, ob dieses Repo
jemals eine von der Implementierung getrennte Abnahme-Instanz hat. Drei Ausgänge:

| Ausgang | Folge |
|---|---|
| **Übergaben definieren** | zwei Artefakte beschreiben, ohne dass jemand sie erzeugt — eine Zusage ohne Deckung |
| **Als `MR-*` deklarieren** | die Abweichung wird sichtbar und begründet; Modul 8 gilt dann ausgewiesen eingeschränkt |
| **Rolle besetzen** | echte Änderung am Arbeitsmodell, weit über diesen Slice hinaus |

Der mittlere Weg entspricht dem, was das Repo mit
[`MR-008`](../../../../harness/conventions.md#mr-008--kein-replay-keine-agenten-telemetrie) für den
fehlenden Replay-Lauf bereits getan hat — **derselbe Fall**: eine Baseline-Erwartung, die dieses
Repo nicht erfüllt, ausgewiesen statt stillschweigend übergangen. Die Entscheidung gehört in die
Closure-Notiz.

**Negativ-Probe:** die Übergaben-Tabelle enthält nach diesem Slice **kein** „fehlt — siehe unten"
mehr, das nicht auf eine Deklaration zeigt. Prüfbar durch Lesen; ein Sensor dafür existiert nicht
und wäre für zwei Zeilen Schein-Genauigkeit.

## 4. Was bewusst nicht getan wird

- **Die Validator-Rolle besetzen.** Das ist eine Entscheidung über das Arbeitsmodell, nicht über
  eine Doku-Zeile.
- **Modul 8 „erfüllen", indem zwei Artefakt-Beschreibungen ohne Erzeuger entstehen.** Eine
  Übergabe, die niemand ausführt, ist die Harness-Lüge, gegen die dieser Slice antritt.
- **Die übrigen sieben Übergaben prüfen.** Sie sind laut
  [`harness/README.md`](../../../../harness/README.md) belegt; das nachzumessen wäre ein eigener
  Schnitt.

## 5. DoD

- [x] Der Entscheid aus §3 ist **getroffen und begründet** in der Closure-Notiz: mittlerer Weg,
      [`MR-009`](../../../../harness/conventions.md#mr-009--validator-rolle-unbesetzt-zwei-übergaben-ohne-artefakt).
- [x] Keine Übergabe ist mehr als „fehlt" geführt, ohne auf eine Deklaration zu zeigen. Beleg:
      `grep '\*\*fehlt\*\*' harness/README.md` liefert nichts; beide Zeilen tragen jetzt
      **unverkörpert** mit Link auf
      [`MR-009`](../../../../harness/conventions.md#mr-009--validator-rolle-unbesetzt-zwei-übergaben-ohne-artefakt).
- [x] `make gates` und `make verify` grün — **Ausgabe in eine Datei**, Exit-Code getrennt geprüft,
      nie in eine Pipe.

## 6. Closure-Notiz

**Geliefert:** [`MR-009`](../../../../harness/conventions.md#mr-009--validator-rolle-unbesetzt-zwei-übergaben-ohne-artefakt)
in [`harness/conventions.md`](../../../../harness/conventions.md); die Übergaben-Tabelle in
[`harness/README.md`](../../../../harness/README.md) führt beide Kanten als **unverkörpert** mit
Verweis auf die Deklaration statt als „fehlt — siehe unten".

**Der Entscheid: der mittlere Weg.** Nicht „Übergaben definieren" — zwei Artefakt-Beschreibungen
ohne Erzeuger wären eine Zusage ohne Deckung, also genau die Harness-Lüge, gegen die dieser Slice
antritt. Nicht „Rolle besetzen" — das ist eine Entscheidung über das Arbeitsmodell, nicht über eine
Doku-Zeile. Bleibt die Deklaration, und dafür stand die Präzedenz im selben Block:
[`MR-008`](../../../../harness/conventions.md#mr-008--kein-replay-keine-agenten-telemetrie)
behandelt den strukturgleichen Fall — eine Baseline-Erwartung, die dieses Repo nicht erfüllt,
ausgewiesen statt stillschweigend übergangen, samt „Folgewirkung, ausdrücklich".

**Lerneintrag — Form: geschärfte Regel.** Als Prüfsatz: *Eine begründete Lücke ist erst dann
deklariert, wenn sie an dem Ort steht, den das Repo für Abweichungen führt — eine Begründung am
Fundort ist eine Fußnote, keine Adaption.*

**Die Ursache** ist die feinste des ganzen Reviews: hier fehlte **nichts Inhaltliches**. Die
Herleitung stand seit slice-066 ausführlich in
[`harness/README.md`](../../../../harness/README.md) — Validation vs. Verifikation, warum die
Kante offen ist, was sie zu schließen bräuchte, sogar der gefährlichste Fall („Verifikation grün,
Validation rot"). Was fehlte, war ausschließlich der **Ort**: der `MR-*`-Block ist die Instanz, an
der dieses Repo Abweichungen führt, und eine Lücke, die dort nicht steht, existiert für jede
Prüfung gegen `modul-08` nicht. Deshalb konnte slice-066 unmittelbar nach der Tabelle „alle 21
Funde geschlossen" schreiben, ohne zu lügen — und trotzdem eine Lücke hinterlassen.

**Ausdrücklich nicht behauptet:** Das Risiko wird durch die Deklaration **nicht kleiner**. Es wird
sichtbar getragen statt unsichtbar. [`MR-009`](../../../../harness/conventions.md#mr-009--validator-rolle-unbesetzt-zwei-übergaben-ohne-artefakt) nennt das in einem eigenen Feld.

**Zwei beobachtbare Closure-Kriterien:**

1. `grep '\*\*fehlt\*\*' harness/README.md` liefert nichts — jede unverkörperte Übergabe zeigt auf
   eine Deklaration.
2. Der `MR-*`-Block trägt zehn Einträge, und [`MR-009`](../../../../harness/conventions.md#mr-009--validator-rolle-unbesetzt-zwei-übergaben-ohne-artefakt) nennt Datum, Geltungsbereich, Adaption,
   Begründung, Folgewirkung und Auflösungs-Trigger — dieselben Pflichtfelder wie die neun
   bestehenden.

**Folge-Slices:** keine. Damit ist Gruppe C abgearbeitet, soweit sie Slices braucht: `F-8` löste
[slice-046](../done/slice-046-regelwerk-v352-migration-analyse.md), `F-10`/`F-15`
[slice-077](../done/slice-077-status-aussagen.md), `F-11` dieser Slice. `F-7` löst der
`welle-12`-Abschluss selbst.

## 7. Sub-Area-Modus

Alle berührten Sub-Areas GF.
