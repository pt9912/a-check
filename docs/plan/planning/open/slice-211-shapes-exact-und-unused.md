# slice-211 — `shapes:` mit `mode: exact` und Opt-in-Befund `shape-unused`

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `git mv` (`make slice-mv`).

**Welle:** welle-16 — [Welle-Plan](../welle-16-shapes-sollform.md).

**Bezug:** [AC-FA-RULE-012](../../../../spec/lastenheft.md#ac-fa-rule-012),
[ADR-0041](../../adr/0041-shapes-sollform-je-datei.md) (beide aus slice-209);
[AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml),
[AC-QA-01](../../../../spec/lastenheft.md#ac-qa-01--determinismus).

**Berührte Spec-Stellen:** `spezifikation.md` §[SPEC-CONF-001](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema), §[SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung)
(in slice-209 geschrieben, hier umgesetzt — keine Spec-Änderung).

**Verantwortlich:** Claude — gesetzt beim Übergang nach `next/` (Maintainer: „ja alles machen", 2026-10-06).

**Autor:** Claude. **Datum:** 2026-10-06.

**Lerneintrag — Form:** wird bei Closure benannt.

---

## 1. Ziel und Abgrenzung

**Ziel:** Zwei Ergänzungen auf dem Normalisierer aus slice-210:
`mode: exact` vergleicht die normalisierte Datei mit einer Sollform-Datei
(`expect:`) und meldet die erste Abweichung als `shape-differs`; ein
`allow`-Eintrag, der in keiner Datei trifft, meldet bei `unused: fail`
`shape-unused` — Exit 1, kein Warn-Level.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Ein Diff über alle Abweichungen.** *Bestand bleibt bewusst stehen*: der CR
  verlangt die **erste** Abweichung; ein vollständiger Diff wäre ein neues
  Ausgabeformat ohne belegten Bedarf.
- **`shape-unused` als Default.** *Bestand bleibt bewusst stehen*: ein neuer
  Befund, der bestehende Konfigurationen rot färbt, wäre Breaking — Opt-in
  hält die Kompatibilitätszusage des CR.
- **Normalisierer-Änderungen.** *Schicht-Abgrenzung*: der Normalisierer ist
  slice-210s Liefer-Punkt; wer hier einen Fehler darin findet, schließt ihn als
  Risiko *eingetreten* mit eigener Kennung, nicht still nebenbei.

## 2. Definition of Done

- [ ] `mode: exact` mit `expect:` (Exit 2 bei fehlender Sollform-Datei oder
      `allow` neben `exact`); `shape-differs` für eine Anweisung anders,
      zusätzlich oder fehlend, jeweils mit der ersten Abweichung — als Tests.
- [ ] `unused: fail` und Befund `shape-unused` (Eintrag ohne Treffer über alle
      Dateien des Blocks), deterministisch sortiert — als Tests.
- [ ] `--print-config`, Benutzerhandbuch und CHANGELOG `[Unreleased]`
      nachgezogen.
- [ ] Unabhängiger Review, Report unter [`docs/reviews/`](../../../reviews/README.md).
- [ ] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [ ] Beobachtungs-Register fortgeschrieben (oder „keine Beobachtung" notiert).
- [ ] Jedes Risiko aus §6 trägt einen Ausgang.

`make gates` grün.

## 3. Plan (vor Code)

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| Regel-Auswertung `shapes` in `internal/hexagon/` | update | `exact`-Vergleich, `unused`-Zählung |
| Config-Adapter | update | `expect`, `unused`, Widerspruchs-Fälle |
| Tests (gleich / anders / zusätzlich / fehlend; unused an/aus) | neu | Happy/Boundary/Negative der neuen Regel-Anforderung |
| `docs/user/benutzerhandbuch.md`, `CHANGELOG.md`, `--print-config` | update | öffentlicher Vertrag berührt |

## 4. Trigger

**Start** (`next` → `in-progress`): slice-210 in `done/`, WIP-Limit frei.

**Rückführungen — vorab benannt:**

- `in-progress` → `next` (zu groß): entfällt voraussichtlich — zwei Modi auf
  vorhandenem Normalisierer; träte es ein, ginge `shape-unused` als eigener
  Slice.
- `in-progress` → `open` (blockiert): der Vertrag aus slice-209 lässt offen,
  wie eine `expect:`-Datei selbst normalisiert wird (gleicher Dialekt?) —
  zurück an den Planner.

## 5. Closure-Trigger

DoD vollständig; bewusstes Brechen je Testbehauptung gezeigt; `make gates`
Exit 0; Closure-Notiz mit Lerneintrag.

## 6. Risiken und offene Punkte

- **`exact` ist spröde:** jede Formatierungs-unabhängige, aber inhaltliche
  Änderung (Versions-Bump) macht die Sollform-Datei nötig — Adopter könnten
  `exact` meiden. — **Ausgang:** *(bei Closure zuzuweisen: eingetreten /
  entfallen / weiter offen)*

## 7. Closure-Notiz

*(folgt bei Closure)*

## 8. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** `KERN` (Achsen 1, 2, 3 ✓) ·
`ADAPT` (2, 3 ✓) · `USER` (2, 3 ✓).

**Vorgelagert — offene Beobachtungen sichten** (2026-10-06, gemergter Stand):
dieselben Treffer wie slice-210 §8 (`BEO-KERN/dirvocab-portfor-auseinander`
1×, `BEO-USER/handbuch-vokabel-der-adapter-rolle` 1×); beim Anlegen erneut
sichten, weil slice-210 sie bis dahin bewegt haben kann.

**Modus-Begründungsblock:** alle berührten Sub-Areas GF.
