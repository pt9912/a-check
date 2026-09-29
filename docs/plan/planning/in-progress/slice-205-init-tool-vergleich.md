# slice-205 — Init-Tool-Vergleich: was liefert `ai-harness-init`, was davon nimmt a-check?

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt. Er wechselt nur durch `git mv`.

**Welle:** ohne Welle — die Closure-Bedingung wäre die eigene DoD.

**Bezug:** Maintainer-Frage „was können wir aus /tmp/aih-v6.13.0 übernehmen?"
(2026-09-29); [`AC-QA-02`](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze).

**Berührte Spec-Stellen:** — · Der Slice berührt kein Spec-Stratum.

**Verantwortlich:** Claude — gesetzt beim Übergang nach `in-progress/`.

**Autor:** Claude. **Datum:** 2026-09-29.

**Lerneintrag — Form:** geschärfte Regel.

---

## 1. Ziel und Abgrenzung

**Ziel:** Die Ausgabe des `ai-harness-init`-Werkzeugs (Kurs `v6.13.0`) ist
gegen a-checks bestehenden Harness inventarisiert; **je Delta trägt eine
begründete Entscheidung** — übernehmen (→ Folge-Slice mit Kennung) oder
ablehnen (mit Begründung). Der Slice selbst führt **keine** Übernahme aus.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Die Übernahme selbst.** *Es wäre ein anderer Vorgang*: jede Übernahme ist
  ein eigener Slice mit eigenem Review — der Vergleich Slice entscheidet nur.
- **Änderungen am Bestand** (AGENTS.md, Makefile, `.d-check.yml`, `tools/`).
  *Schicht-Abgrenzung*: der Slice ist ein Vergleich — sein einziges Artefakt
  neben diesem Plan ist die Entscheidungsliste.

## 2. Ausgangsmessung

Instrument: `tree -a` über die Generator-Ausgabe, `ls`/`head` über die
Neuheiten, Vergleich gegen a-checks Bestand. Gemessen am 2026-09-29.

## 3. Umsetzung

### 3.1 Inventar — Neuheiten des Generators

| Feature | Was es ist | Entscheidung | Begründung |
|---|---|---|---|
| **Erfassungsschicht** (`.claude/hooks/span-emit.sh` + Traeger + `erfassung-feldliste.md`) | Je Werkzeug-Aufruf eines Agenten-Laufs eine JSON-Zeile in einen gitignorierten Zustands-Bereich; der Wrapper hält keinen Tool-Call auf (Exit 0 in jedem Zweig) | **Übernehmen — Folge-Slice slice-206**, mit [`MR-014`](../../../../harness/conventions.md#mr-014)-Abgrenzung | Die Erfassung ist **lokal** (gitignoriert, kein Abfluss nach draußen); der geltende Wortlaut von [`MR-014`](../../../../harness/conventions.md#mr-014) verbietet Tool-Call-Spans aber **uneingeschränkt**. Die Übernahme braucht darum einen **Nachfolge-Eintrag**, der die Lokal/Draußen-Abgrenzung explizit macht — **Maintainer-Vorbehalt** |
| **`.claude/agents/`** — sechs Rollen (Planner/Architect/Implementer/Reviewer/Verifier/Validator) | Subagent-Definitionen, die per Zeiger auf die Commands verweisen („Dein Anweisungssatz steht in implement-slice — lies ihn") | **Übernehmen — Folge-Slice slice-207** | Macht die Rollen aus Modul 8 maschinell ansprechbar; die Commands tragen den Inhalt, die Agents den Zeiger — Form folgt Ziel-Form `v6.13.0` |
| **`.claude/commands/`** — `plan-welle`, `implement-slice`, `close-welle` | Die Lifecycle-Commands (Modul 9 8-Schritt-Flow als Slash-Command) | **Übernehmen — in slice-207 gebündelt** | Der 8-Schritt-Flow ist in `AGENTS.md` §6 vorhanden; die Commands sind die maschinenlesbare Hülle dafür |
| **`harness/mk/`** — modularisierte Make-Includes (10 Module) | Makefile in Module geschnitten | **Abgelehnt** | a-checks monolithisches Makefile ist etabliert (259 Zeilen, gewachsen); der Umzug ist Struktur-Kosmetik mit Migrations-Risiko und ohne Funktions-Zuwachs |
| **`tools/harness/`-Namensraum** | Scripts unter `tools/harness/` statt `tools/` | **Abgelehnt** | a-checks `tools/`-Namen sind repo-weit referenziert (AGENTS.md, harness/README.md, CI); Renames brechen Referenzen ohne Zuwachs |
| **Generiertes Spec-Stratum** (Platzhalter-Specs) | `<Projektname>`-/`YYYY-MM-DD`-Platzhalter | **Abgelehnt** | a-checks Straten sind 190 Slices weit über die Platzhalter hinaus — ein Re-Bootstrap wäre ein Reset |
| **`file`-Modul** (d-check `v0.79.0`, hier nicht aktiv) | `max-lines`/`max-bytes` pro Datei | **Abgelehnt (vorerst)** | Ohne beobachteten Vorfall — `AGENTS.md` ist konvertiert und kurz; beim ersten Wachstums-Signal neu bewerten |

### 3.2 Die MR-014-Abgrenzung (Erfassungsschicht)

[`MR-014`](../../../../harness/conventions.md#mr-014)
(„Keine Agenten-Telemetrie") verbietet in seinem geltenden Wortlaut
Tool-Call-Spans **uneingeschränkt** — und genau solche erzeugt die
Erfassungsschicht („je Werkzeug-Aufruf … eine JSON-Zeile"). Der
Auflösungs-Trigger des Eintrags („sobald Agenten-Läufe **im Repo selbst**
abrechenbar werden") ist mit der Übernahme eingetreten, nicht entlastet. Die
Übernahme ist darum nur mit einem **Nachfolge-Eintrag** zu [`MR-014`](../../../../harness/conventions.md#mr-014)
sauber zu führen, der die Lokal/Draußen-Abgrenzung explizit macht.
**Maintainer-Vorbehalt:** die Entscheidung über die Erfassung selbst ist die
des Maintainers — der Start-Trigger von slice-206 hängt an ihr.

## 4. Definition of Done

- [x] Je Neuheit des Generators trägt die Inventar-Tabelle eine begründete
      Entscheidung (übernehmen mit Folge-Slice-Kennung / abgelehnt).
- [x] Die [`MR-014`](../../../../harness/conventions.md#mr-014)-Abgrenzung für die Erfassungsschicht ist formuliert mit
      Maintainer-Vorbehalt.
- [x] Unabhängiger Review, Report unter [`docs/reviews/`](../../../reviews/README.md).
- [x] Closure-Notiz mit Steering-Loop-Lerneintrag; Register fortgeschritten;
      jedes Risiko aus §6 trägt einen Ausgang.

`make gates` und `make verify` grün.

## 5. Trigger

**Start** (`open` → `in-progress`): das WIP-Limit ist frei.

**Rückführungen — vorab benennen:**

- `in-progress` → `next` (zu groß): entfällt — ein Vergleich, keine Konversion.
- `in-progress` → `open` (blockiert): ist `/tmp/aih-v6.13.0` nicht mehr
  lesbar, bricht der Vergleich — das Init-Werkzeug ist Host-Binary
  (`ai-harness-init` v0.2.5, linuxbrew) und liegt **nicht** im Repo-Bestand;
  die Ausgabe wäre erst durch einen erneuten Lauf des Werkzeugs zu
  reproduzieren.

## 6. Closure-Trigger

DoD vollständig, `make gates` und `make verify` grün, Closure-Notiz mit
Lerneintrag.

## 7. Risiken und offene Punkte

- **Die Generator-Ausgabe liegt nur in `/tmp`** — ein Reboot des Rechners
  löscht sie; das Init-Werkzeug ist Host-Binary (v0.2.5) und liegt nicht im
  Repo-Bestand.
  — **Ausgang:** *entfallen*, gestrichen mit Begründung: der Vergleich ist
  mit §3.1 abgeschlossen und belegt; die /tmp-Ausgabe wird für die
  Entscheidung nicht mehr gebraucht. Die Wiederholung bei Neu-Eintritt
  trägt [`BEO-GATE/init-tool-drift`](../observations/BEO-GATE/init-tool-drift/observation.md).
- **Der Vergleich vergreist sich an einer Snapshot-Version** — der Generator
  entwickelt sich mit dem Kurs weiter. — **Ausgang:** *weiter offen* →
  Beobachtung: beim nächsten Init-Tool-Release neu vergleichen
  ([`BEO-GATE/init-tool-drift`](../observations/BEO-GATE/init-tool-drift/observation.md),
  neu, 1×).

## 8. Closure-Notiz

**Lerneintrag — Form: geschärfte Regel.** *Bootstrap-Vergleiche führen, nicht
durchführen: je Delta eine begründete Entscheidung (Übernahme als Folge-Slice
mit Kennung oder Ablehnung mit Begründung), niemals die Übernahme selbst in
den Vergleich ziehen — der Vergleich ist ein Urteil über eine Menge, und
Urteile sind nach [`AC-QA-02`](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze) vom Aufwand getrennt.* Gemessen: 7 Deltas, 3
Übernahmen (als Folge-Slices geschnitten), 4 Ablehnungen mit Begründung.

**Was hat funktioniert:** die Inventar-Tabelle direkt im Plan (§3.1) statt
eine separate Datei — der Plan ist das einzige Artefakt neben den
Entscheidungen.

**Was ging anders als geplant:** der unabhängige Review fand vier
HIGH-Befunde — die MR-014-Lesart (F-1), zwei falsche Tatsachenbehauptungen
(Werkzeug-Pfad F-2, Makefile-Zahl F-3) und den vorweg angehakten
Review-DoD-Haken (F-4). Alle im Nachlauf korrigiert, bevor die Closure
geschrieben wurde.

**Steering-Loop-Eintrag:** — *(nichts verkörpert; die Entscheidungen sind die
Arbeit, und die Folge-Slices tragen die Umsetzung.)*

**Beobachtungs-Register (`../observations/`):**
`BEO-GATE/init-tool-drift/` neu angelegt — 1× (der Generator driftet mit dem
Kurs; der Vergleich wird bei dessen Release wiederholt).
`BEO-GATE/attestierung-vor-dem-vorgang/` — 5. Auftreten: der Review-DoD-Haken
war vor der Existenz des Reports angehakt
([evidence/slice-205.md](../observations/BEO-GATE/attestierung-vor-dem-vorgang/evidence/slice-205.md)).

**Folge-Slices:** [slice-206](../open/slice-206-erfassungsschicht.md)
(Erfassungsschicht mit [`MR-014`](../../../../harness/conventions.md#mr-014)-Abgrenzung),
[slice-207](../open/slice-207-rollen-und-commands.md) (Rollen + Commands) —
sind Dateien in `open/`.

**Risiken aus §7:** jedes mit genau einem Ausgang — siehe §7.

**Drei Paarungen:** Anker — nichts verkörpert · Folge-Slice — slice-206 und
slice-207 existieren als Dateien in `open/` · Register — init-tool-drift
neu (1×) · attestierung-vor-dem-vorgang (5×, Beleg ergänzt).

## 9. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** GATE (Achsen 1, 2, 3 ✓).

**Vorgelagert — offene Beobachtungen sichten:** das Register wurde durchgegangen
(2026-09-29): keine Treffer in GATE für diesen Vorgang.

**Modus-Begründungsblock:** alle berührten Sub-Areas GF.
