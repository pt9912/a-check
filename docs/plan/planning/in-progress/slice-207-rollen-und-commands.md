# slice-207 — Rollen- und Command-Verdrahtung übernehmen

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt. Er wechselt nur durch `git mv`.

**Welle:** ohne Welle — die Closure-Bedingung wäre die eigene DoD.

**Bezug:** Folge-Slice aus
[slice-205](../done/wellenlos/slice-205-init-tool-vergleich.md) (Entscheidung
„Übernehmen").
[`AC-QA-02`](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze).

**Berührte Spec-Stellen:** — · Der Slice berührt kein Spec-Stratum.

**Verantwortlich:** Claude — gesetzt beim Übergang nach `in-progress/`.

**Autor:** Claude. **Datum:** 2026-09-29.

**Lerneintrag — Form:** neuer Sensor.

---

## 1. Ziel und Abgrenzung

**Ziel:** Die Rollen- und Command-Verdrahtung des Generators ist übernommen —
`.claude/agents/` (sechs Rollen-Subagents, die per Zeiger auf ihre
Anweisungsquelle verweisen: Commands, Skill oder Werkzeug) und
`.claude/commands/` (`plan-welle`, `implement-slice`,
`close-welle`). **Die Inhalte der Commands werden an a-checks Prozess
angepasst** (`AGENTS.md` §6 8-Schritt-Workflow statt des Generators-Wortlauts;
a-checks Adaptionen sind im AGENTS.md-Träger verankert, den die Commands
referenzieren).

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Inhaltliche Abweichungen vom Modul-8-Prozess.** *Schicht-Abgrenzung*: die
  Commands tragen a-checks Prozess; eigene Erfindungen sind Out-of-Scope.

## 2. Definition of Done

- [x] `.claude/agents/` (sechs Rollen) und `.claude/commands/` (drei
      Lifecycle-Commands) existieren, an a-checks Prozess angepasst.
- [x] Die Zeiger-Kette ist geprüft: jeder Agent verweist auf seine
      Anweisungsquelle, jeder Command auf die kanonischen Quellen;
      `make doc-check` grün (644/645 Dateien, 0 Befunde).
- [x] Unabhängiger Review, Report unter [`docs/reviews/`](../../../../docs/reviews/README.md).
- [x] Closure-Notiz mit Steering-Loop-Lerneintrag; Register fortgeschritten;
      jedes Risiko aus §6 trägt einen Ausgang.

`make gates` und `make verify` grün.

## 3. Plan (vor Code)

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| `.claude/agents/` (6 Dateien) | neu | Rollen-Subagents mit Zeiger auf die Commands |
| `.claude/commands/` (3 Dateien) | neu | plan-welle, implement-slice, close-welle — an `AGENTS.md` §6 angepasst |

## 4. Trigger

**Start** (`open` → `in-progress`): das WIP-Limit ist frei.

**Rückführungen — vorab benennen:**

- `in-progress` → `next` (zu groß): entfällt — drei Datei-Familien mit
  Zeiger-Logik.
- `in-progress` → `open` (blockiert): entfällt — kein Blocker absehbar.

## 5. Closure-Trigger

DoD vollständig, `make gates` und `make verify` grün, Closure-Notiz mit
Lerneintrag.

## 6. Risiken und offene Punkte

- **Die Commands frieren den Prozess** — eine AGENTS.md-Änderung veraltet die
  Command-Kopien. — **Ausgang:** *entfallen*, gestrichen mit Begründung: die
  Commands referenzieren `AGENTS.md` als Zeiger statt es zu kopieren; veralten
  kann damit keine Kopie.

## 7. Closure-Notiz

**Lerneintrag — Form: geschärfte Regel.** *Der Claim eines Slice entfernt den
Ruhe-Marker aus der Roadmap; er kommt mit der Archivierung zurück.* Verkörpert
im Eingang-Absatz des [`implement-slice`](../../.claude/commands/implement-slice.md)-Commands
(`seit slice-207`). Gemessen: zweimal `doc-planning` rot nach forgetful Claim am
2026-09-29 (slice-205, slice-207) — beide Vorfälle maschinell vom Sensor gefangen,
darum ist die Klasse nicht zusätzlich ins Register gelegt.

**Was hat funktioniert:** die Zeiger-Form — der Generator-Wortlaut (25 Schritte
Implementer, 8 Schritte Closure) bleibt außen vor, `AGENTS.md` §6 und der
`/slice`-Command tragen den Prozess; `doc-check` belegte die Zeiger-Kette in
beiden Richtungen.

**Was ging anders als geplant:** der unabhängige Review fand einen
merge-blockierenden Risiko-Ausgang (F-1 — „bei Closure" ist ein Zeitpunkt, kein
Ausgang aus der geschlossenen Menge) — maschinell von `verify-risiko-ausgaenge`
bestätigt und im Nachlauf korrigiert; daneben drei LOW (Adaptions-Verlust beim
Abschluss-Gate, verallgemeinerte Plan-Aussage, Zähler-Stand der Sichtung),
ebenfalls korrigiert.

**Steering-Loop-Eintrag:** siehe Lerneintrag oben. **Kein neuer Sensor:** die
zwei Findings-Klassen dieses Laufs (Risiko-Ausgang als Zeitpunkt; Ruhe-Marker
vergessen) sind beide von bestehenden Sensoren maschinell gefangen worden
(`verify-risiko-ausgaenge`, `doc-planning`) — ein Zähler für Wächter-bedeckte
Klassen zählt doppelt.

**Beobachtungs-Register (`../observations/`):** keine neue Beobachtung
angefallen — GATE trug 15 offene Einträge bei Sichtung (§8), keiner davon
reicht an diesen Vorgang heran.

**Folge-Slices:** keine.

**Risiken aus §6:** das eine Risiko trägt seinen Ausgang (*entfallen*, mit
Begründung — Zeiger statt Kopie).

**Drei Paarungen:** Anker — nichts verkörpert jenseits des Lerneintrags oben ·
Folge-Slice — keine genannt · Register — keine neue Zeile.

## 8. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** GATE (Achsen 1, 2, 3 ✓).

**Vorgelagert — offene Beobachtungen sichten:** das Register wurde durchgegangen
(2026-09-29): GATE trägt 15 offene Einträge (Zählung über die `state.md`-Köpfe),
keiner betrifft die Sub-Area-Berührung dieses Vorgangs.

**Modus-Begründungsblock:** alle berührte Sub-Areas GF.
