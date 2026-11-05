# slice-229 — Quittieren der Port-Scope-Meldung: Anforderung und Entscheidung

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `git mv` (`make slice-mv`).

**Welle:** ohne Welle — die Closure-Bedingung wäre die eigene DoD.

**Bezug:** GitHub-Issue #7 (Wunsch), GitHub-Issue #4 (Anlass),
[ADR-0040](../../adr/0040-portscope-richtungssegment.md) (Port-Scope-Diagnose),
[AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml) (Konfigurationsdatei, strikt dekodiert),
[AC-FA-RULE-010](../../../../spec/lastenheft.md#ac-fa-rule-010--port-lokalität-regel-port-locality) (`port-locality`).

**Berührte Spec-Stellen:** [AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml) (eine neue AC im Bereich CONF entsteht daneben) ·
[SPEC-CONF-001](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema) und
[SPEC-CLI-001](../../../../spec/spezifikation.md#spec-cli-001--aufruf-scan-wurzel-und-exit-codes) werden vom Folge-Slice nachgezogen, nicht hier.

**Verantwortlich:** —

**Autor:** Claude. **Datum:** 2026-11-05.

**Lerneintrag — Form:** wird bei Closure benannt.

---

## 1. Ziel und Abgrenzung

**Ziel:** Der Wunsch aus Issue #7 — eine Schicht quittiert die Port-Scope-Meldung gezielt — ist als
**Anforderung** im Lastenheft und als **Entscheidung** in einer ADR festgeschrieben, bevor Code
entsteht. Die Maintainer-Entscheidung (2026-11-05): annehmen, eng gefasst. Der Rahmen:

1. **Nur die Port-Scope-Diagnose**, **je Schicht**, mit **Pflicht-Begründung**, in der Konfiguration.
2. **Die Meldung bleibt sichtbar, nur kürzer:** die Zeile nennt die **Schicht** und die Begründung,
   nicht nur eine Zahl — sonst verdeckt die Quittierung, was die Meldung zeigt.
3. **Eine überflüssige Quittierung fällt auf:** löst die Schicht die Meldung gar nicht aus, wird das
   gemeldet, damit eine veraltete Quittierung nicht unbemerkt stehen bleibt.
4. **Exit-Code und Regel unberührt:** `port-locality` prüft weiter genau so viel wie vorher.

Offen, und von der ADR zu entscheiden: der Name und die Form des Schlüssels (er soll sich später
auf andere Diagnosen erweitern lassen, ohne dass das hier zugesagt wird), die Zeilenform der
kurzen Meldung, und wie sich ein Konfigurations-Fehler (Schicht ohne Port-Rolle, leere Begründung)
verhält — Exit 2 wie bei jedem strikten Schema-Verstoß, oder Hinweis.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Die Umsetzung** (Schema-Dekodierung, Kern, Hinweis-Ausgabe, Handbuch). Ein Folge-Slice
  übernimmt sie: [slice-230](../open/slice-230-port-scope-meldung-quittieren-umsetzung.md). Der Schnitt
  nach Lieferwert: hier entsteht die Zusage, dort die Lieferung; Schema, Kern und Ausgabe sind mehr
  als zwei Schichten.
- **Quittieren der anderen drei Diagnosen** (Abdeckung, Grenze, Auflösung). *Bestand bleibt bewusst
  stehen*: bisher meldet ein einziger Konsument das Rauschen einer einzigen Diagnose; ein
  verallgemeinerter Entwurf ohne zweiten Beleg wäre geraten. Der Schlüssel wird nur so benannt,
  dass die Erweiterung möglich bleibt.
- **Eine Änderung von `port-locality` selbst** und von der Scope-Ableitung. *Ein anderer Vorgang*:
  die Quittierung ändert, was gemeldet wird, nicht was geprüft wird.
- **Ein stummes Quittieren und ein Schalter für alle Hinweise.** *Ein anderer Vorgang*, und vom
  Melder selbst ausgeschlossen: ein Schweigen verdeckte genau das, was der Hinweis sichtbar macht.

## 2. Definition of Done

- [ ] Lastenheft: eine neue AC im Bereich CONF für das Quittieren (drei Akzeptanzkriterien
      Happy/Boundary/Negative im Given/When/Then-Stil plus Out-of-Scope), Versions-Bump und
      Historie-Zeile (Change Request, [AGENTS.md](../../../../AGENTS.md) §5 Regel 4).
- [ ] Eine ADR mit der Form des Schlüssels, der Zeilenform und dem Fehlerverhalten; `Schärft:`
      auf die Spezifikation, **nicht** auf das Lastenheft; ADR-Index nachgezogen.
- [ ] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [ ] Beobachtungs-Register fortgeschrieben (oder „keine Beobachtung" notiert).
- [ ] Jedes Risiko aus §6 trägt einen Ausgang.

`make gates` und `make verify` grün. Ein unabhängiger Review ist vorgesehen
([MR-019](../../../../harness/conventions.md#mr-019): Opt-in) — der Slice legt eine
Vertragsänderung fest.

## 3. Plan (vor Code)

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| `spec/lastenheft.md` | update | neue AC im Bereich CONF, Versions-Bump, Historie |
| `docs/plan/adr/NNNN-….md`, `docs/plan/adr/README.md` | neu / update | Entscheidung zur Form; Index |

## 4. Trigger

**Start** (`open` → `next`): Die Maintainer-Entscheidung zu Issue #7 liegt vor — erfüllt am
2026-11-05 (annehmen, eng gefasst). Priorität: niedrig; der Hinweistext sagt seit slice-228, dass
die Meldung für app-weite Ports erwartbar ist. Ein **zweiter** Konsument mit demselben Rauschen
zöge den Entwurf auf alle Diagnosen, bevor der Slice startet.

**Rückführungen:** `in-progress` → `open` (blockiert): die ADR scheitert an einem Schema-Zwang —
eine ältere Version lehnt den neuen Schlüssel mit Exit 2 ab; dann ist die Frage der
Versions-Kopplung vor der ADR zu klären.

## 5. Closure-Trigger

DoD vollständig, Gates grün, Closure-Notiz mit Lerneintrag.

## 6. Risiken und offene Punkte

- **Das strikte Schema lehnt den neuen Schlüssel in älteren Versionen mit Exit 2 ab:** ein
  Konsument, der die Konfiguration anpasst und ein altes Image fährt, bricht. — **Ausgang:**
  *(bei Closure zuzuweisen: eingetreten / entfallen / weiter offen)*
- **Eine Quittierung deckt auch einen später entstandenen echten Fehler zu:** sie gilt für alle
  Port-Globs der Schicht. — **Ausgang:** *(bei Closure zuzuweisen: eingetreten / entfallen /
  weiter offen)*

## 7. Closure-Notiz

*(folgt bei Closure)*

## 8. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** `SPEC` (Achsen 1, 2, 3 ✓) und `ADR` (Achsen 1, 2, 3 ✓).

**Vorgelagert — offene Beobachtungen sichten** (2026-11-05): `BEO-KERN/dirvocab-portfor-auseinander`
(2×, offen) — dieselbe `direction`-Zuordnung steht an drei Stellen; berührt diesen Slice nur,
wenn die ADR ein Richtungs-Vokabular in den Schlüssel aufnimmt. Sonst keine Treffer zu Diagnose-
Rauschen oder Quittierung.

**Modus-Begründungsblock:** alle berührten Sub-Areas GF.
