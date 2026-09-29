# slice-204 — Review-Haken an Report-Existenz binden (in-progress)

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt. Er wechselt nur durch `git mv`.

**Welle:** ohne Welle — die Closure-Bedingung wäre die eigene DoD.

**Bezug:** 4. Auflage von
[`BEO-GATE/attestierung-vor-dem-vorgang`](../observations/BEO-GATE/attestierung-vor-dem-vorgang/observation.md)
(slice-169, slice-197, slice-200, slice-191 — Beleg im evidence-Verzeichnis).
[`AC-QA-02`](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze).

**Berührte Spec-Stellen:** — · Der Slice berührt kein Spec-Stratum.

**Verantwortlich:** — bis zur Priorisierung.

**Autor:** Claude. **Datum:** 2026-09-29.

**Lerneintrag — Form:** neuer Sensor.

---

## 1. Ziel und Abgrenzung

**Ziel:** Der DoD-Punkt „Unabhängiger Review durchgeführt" wird in
`in-progress/` an die Existenz des Reports gebunden — der vierte Vorfall
(slice-191: der Closure-Commit attestierte den Review, der Report entstand
erst danach) zeigte, dass Bedingung 7 (Häkchen in `open/`/`next/`) diese
Gestalt nicht fängt: sie entsteht im `in-progress/`-Stand.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Änderungen am `reviews`-Modul oder an MR-019.** *Es wäre ein anderer
  Vorgang*: MR-019 (Opt-in) und der done/-Geltungsbereich von `doc-reviews`
  sind deklarierte Entscheidungen; die Bindung hier läuft über das
  `structure`-Modul, nicht über `doc-reviews`.

## 2. Definition of Done

- [ ] Die `structure`-Bedingung ist konfiguriert: der DoD-Punkt „Unabhängiger
      Review durchgeführt" in einem `in-progress/`-Slice ohne vorhandenen
      Report meldet `section-forbidden` bzw. den passenden Grund-Code.
- [ ] **Gegenprobe in beiden Richtungen** im Selbsttest: mit Report bleibt
      die Stelle grün, ohne Report rot; die `doc-reviews`-Bedeutung
      (done/-Geltungsbereich) bleibt unberührt.
- [ ] Unabhängiger Review, Report unter [`docs/reviews/`](../../../../docs/reviews/README.md).
- [ ] Closure-Notiz mit Steering-Loop-Lerneintrag; Register fortgeschritten;
      jedes Risiko aus §6 trägt einen Ausgang.

`make gates` und `make verify` grün.

## 3. Plan (vor Code)

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| [`.d-check.yml`](../../../../.d-check.yml) (structure, Bedingung 9) | neu | Bedingung auf `in-progress/**/slice-*.md` mit Abschnitts-Selektor und Report-Existenz-Prüfung |
| `tools/dcheck-phrase-selftest.sh` (Muster 5) | neu | Positiv (ohne Report rot) / Negativ (mit Report grün) |

## 4. Trigger

**Start** (`open` → `in-progress`): das WIP-Limit ist frei.

**Rückführungen:** `in-progress` → `next` (zu groß): entfällt — zwei
Konfigurations-Blöcke. `in-progress` → `open` (blockiert): bietet das Modul
keine Existenz-Prüfung für Dateien außerhalb des Repos... (entfällt, das
Repo ist die Wurzel).

## 5. Closure-Trigger

DoD vollständig, `make gates` und `make verify` grün, Closure-Notiz mit
Lerneintrag.

## 6. Risiken und offene Punkte

- **Die Existenz-Prüfung braucht mehr als ein Pattern** — `forbid-pattern`
  prüft Text, nicht Datei-Existenz; der Weg führt über das `reviews`-Modul
  mit `done-dir`-Äquivalent für in-progress oder eine eigene Schalung.
  — **Ausgang:** bei Closure (Umsetzungs-Entscheid mit Beleg).

## 7. Closure-Notiz

*(wird vor dem `git mv` nach `done/` gefüllt)*

## 8. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** GATE (Achsen 1, 2, 3 ✓).

**Vorgelagert — offene Beobachtungen sichten:** das Register wurde durchgegangen
(2026-09-29):
[`attestierung-vor-dem-vorgang`](../observations/BEO-GATE/attestierung-vor-dem-vorgang/observation.md)
— **dieser Slice ist ihr Ausgang** (4×, die Erweiterung um in-progress).

**Modus-Begründungsblock:** alle berührten Sub-Areas GF.
