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

- [ ] Der Entscheid aus §3 ist **getroffen und begründet** in der Closure-Notiz.
- [ ] Keine Übergabe ist mehr als „fehlt" geführt, ohne auf eine Deklaration zu zeigen. Beleg: die
      Tabelle in [`harness/README.md`](../../../../harness/README.md) und, bei Deklaration, der
      neue `MR-*`-Eintrag.
- [ ] `make gates` und `make verify` grün — **Ausgabe in eine Datei**, Exit-Code getrennt geprüft,
      nie in eine Pipe.

## 6. Closure-Notiz

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice,
[`AGENTS.md`](../../../../AGENTS.md) §5; `make verify` prüft das.)_

## 7. Sub-Area-Modus

Alle berührten Sub-Areas GF.
