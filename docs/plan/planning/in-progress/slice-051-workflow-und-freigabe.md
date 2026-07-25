# slice-051 — Etappe E (3/3): Workflow-Skelett und Freigabe-Belege

**Status:** in-progress — dritter und letzter Schnitt der **Etappe E (Mechanik)** aus
[slice-048 §5](../done/slice-048-modul-delta-lesen.md).
**Deckt:** Fund **B-16** (dritter Bindepunkt der Durchsetzungsschicht fehlt) und **B-20**
(`releasing.md` ist Prozedur ohne Freigabe-Belege).
[Roadmap](roadmap.md).

---

## 1. B-16 — der fehlende dritte Bindepunkt

Der Regelwerk-Abschnitt `grundlagen-durchsetzungsschicht` nennt drei Bindepunkte an der
Agent-Schleife. a-check hat zwei davon, beide fail-closed verdrahtet:

| Bindepunkt | Wann | Quadrant | in a-check |
|---|---|---|---|
| Tool-Call-Gate | vor jedem Tool-Call | computational feedforward | `pretooluse-command-guard.sh` ✓ |
| Handoff-Gate | vor der „fertig"-Meldung | computational feedback | `stop-require-gates.sh` ✓ |
| **Workflow-Skelett** | beim Start einer Aufgabe | inferential feedforward | **fehlte** |

Das Skelett ist ausdrücklich der **schwächste** der drei — es gibt den Ablauf vor, erzwingt ihn
nicht, und bleibt das einzige, das ein Agent ignorieren kann. Genau deshalb ist es billig zu
liefern und trotzdem wirksam: es hält den 8-Schritt-Pfad aus
[`AGENTS.md`](../../../../AGENTS.md) §6 im Lauf präsent, statt darauf zu hoffen, dass er gelesen
wurde.

**Bewusst keine Automatik.** Der Slash-Command *behauptet* nichts und prüft nichts; täte er es,
wäre er ein Gate ohne Sensor. Er ordnet nur — inklusive der beiden Schritte, die a-check heute
zusätzlich kennt: `make verify` beim Abschluss (slice-050) und der Gate-Lauf **in eine Datei**
statt in eine Pipe.

Letzteres ist die Antwort auf den in [slice-048 §2](../done/slice-048-modul-delta-lesen.md)
dokumentierten Fehler: `make … | tail` liefert den Exit-Code von `tail`, nicht von `make` — vier
Vorfälle an einem Tag, nach der 3×-Regel des Regelwerks längst eine Harness-Lücke. Sie bekommt
hier ihren **ersten Guide**; ein Sensor dafür bleibt offen (§4).

## 2. B-20 — Freigabe ohne Belege ist Bürokratie

[`releasing.md`](../../../../docs/user/releasing.md) beschreibt den Ablauf gut (Tag → Pipeline →
Re-Pin), aber Modul 16 verlangt etwas anderes: eine **Freigabe-Checkliste, in der kein Häkchen
ohne Beleg-Slot existiert**. Gemessen vor diesem Slice: **null** Checklisten-Items, keine
Anti-Item-Liste, keine Incident-Klausel.

Für a-check ist das keine Formalie: das Repo liefert ein **digest-gepinntes Image in fremde
CI-Läufe**. Wenn ein Release kaputt ist, hängt an der Frage „Rollback oder Fix-Forward?" der
Aufwand jedes Konsumenten — und die Antwort gehört *vor* den Vorfall, nicht hinein. Die
Konsumenten-Seite ist dabei asymmetrisch: ein Rollback heißt hier **nicht**, dass a-check etwas
zurücknimmt, sondern dass **jeder Konsument seinen Pin selbst zurückdreht**. Das macht
Fix-Forward zur Normalantwort und schreibt die Ausnahmen auf.

## 3. Betroffene Module

- `.claude/commands/slice.md` — Workflow-Skelett (B-16).
- [`docs/user/releasing.md`](../../../../docs/user/releasing.md) — Freigabe-Checkliste mit
  Beleg-Slots, Anti-Items, Incident-Klausel (B-20).

Zwei Schichten (Durchsetzungsschicht-Artefakte, Betriebs-Doku) — innerhalb der B-1-Grenze.

## 4. Was offen bleibt

- **Kein Sensor für den Pipe-Fehler.** Der Guide aus §1 ist inferential; die computational Hälfte
  (etwa ein Guard, der `make …` in einer Pipe ablehnt) ist nicht gebaut. Nach Modul 09 ist die
  Regel damit **halb durchgesetzt** — dieselbe Klasse wie B-11 vor slice-049. Ausdrücklich
  benannt statt stillschweigend ausgelassen; Kandidat für Etappe F.
- **Die Freigabe-Checkliste ist nicht maschinell geprüft.** Sie ist ein Beleg-Formular, kein Gate.
  Ob jeder Beleg-Slot beim Release wirklich gefüllt wurde, sieht heute nur ein Mensch.

## 5. DoD

- [x] `.claude/commands/slice.md` existiert und gibt den 8-Schritt-Pfad plus a-checks zwei
      Zusatz-Schritte (`make verify`, Gate-Lauf in Datei) vor — ohne etwas zu behaupten, das kein
      Sensor deckt (B-16).
- [x] [`releasing.md`](../../../../docs/user/releasing.md) trägt eine Freigabe-Checkliste, in der
      **jedes** Item einen Beleg-Slot hat, plus Anti-Item-Liste und Incident-Klausel
      (Rollback · Fix-Forward · Konsumenten-Pin) (B-20).
- [x] `make gates` und `make verify` grün.

## 6. Closure-Notiz

**Geliefert:** der dritte Bindepunkt der Durchsetzungsschicht als `.claude/commands/slice.md` und
eine Freigabe-Checkliste mit acht Beleg-Slots, Anti-Item-Liste und Incident-Klausel in
[`releasing.md`](../../../../docs/user/releasing.md). Damit ist **Etappe E vollständig**.

**Lerneintrag — Form: benannte Spec-Lücke.**
> **Der Pipe-Fehler hat jetzt einen Guide, aber weiterhin keinen Sensor.** `make … | tail` liefert
> den Exit-Code von `tail`; vier Vorfälle an einem Tag machen daraus nach der 3×-Regel des
> Regelwerks eine Harness-Lücke. Dieser Slice schließt die *inferentielle* Hälfte (Schritt 6 des
> Skeletts sagt es ausdrücklich) — die *computational* fehlt. Nach Modul 09 ist die Regel damit
> halb durchgesetzt, dieselbe Klasse wie B-11 vor slice-049. Der Unterschied zu damals: es steht
> hier, statt unbemerkt zu bleiben.

**Zwei beobachtbare Closure-Kriterien:**

1. `make gates` und `make verify` grün auf demselben Stand (je Exit 0) — belegt.
2. Jedes der acht Freigabe-Items nennt einen konkreten Beleg (Gate-Exit, CI-Link, Job-Summary);
   kein Item ist ohne Slot abhakbar.

**Beobachtung zur Incident-Klausel:** die Regel fiel nicht aus dem Regelwerk, sondern aus a-checks
Verteilmodell. Weil Konsumenten **Digests** pinnen, wirkt ein Rollback erst, wenn jeder von ihnen
selbst zurückdreht — Fix-Forward ist deshalb hier die Normalantwort, nicht die Ausnahme. Ein
allgemeines Runbook hätte das Gegenteil nahegelegt.

**Folge-Slices:** Etappe **D** (Form: B-1, B-5, B-6, B-12, B-15, B-18), Etappe **C**
(`MR-*`-Bereinigung), Etappe **F** (Betriebsmodell) — dort gehört auch der fehlende Pipe-Sensor
hin.
