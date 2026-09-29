# slice-208 — Trigger-Audit für MR-Einträge mechanisieren

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt. Er wechselt nur durch `git mv`.

**Welle:** ohne Welle — die Closure-Bedingung wäre die eigene DoD.

**Bezug:** 3. Auflage von
[`BEO-HARNESS/mr-aufloesungs-trigger-ohne-waechter`](../observations/BEO-HARNESS/mr-aufloesungs-trigger-ohne-waechter/observation.md)
(slice-170, slice-197, slice-206).
[`MR-030`](../../../../harness/conventions.md#mr-030).

**Berührte Spec-Stellen:** — · Der Slice berührt kein Spec-Stratum.

**Verantwortlich:** Claude — gesetzt beim Übergang nach `in-progress/`.

**Autor:** Claude. **Datum:** 2026-09-29.

**Lerneintrag — Form:** neuer Sensor.

---

## 1. Ziel und Abgrenzung

**Ziel:** Die „dass"-Hälfte aus dem Register-Eintrag wird mechanisiert: ein
Sensor der Verifikations-Schicht prüft, dass jede Closure die aktiven
`MR`-Einträge auf ihren Auflösungs-Trigger angesehen hat — als
Beleg-Pflicht der Closure-Notiz („Trigger-Audit der aktiven MR: 0 offen"
mit Kennung), nicht als Urteil über die Bedingung selbst.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Die „ob"-Hälfte (stimmt das Urteil?).** *Bestand bleibt bewusst
  stehen*: die Bedingung ist Prosa und maschinell nicht auswertbar — das
  Register-Eintrag benennt die Grenze selbst.
- **Änderungen am `reviews`- oder `structure`-Modul von d-check.**
  *Es wäre ein anderer Vorgang*: die Prüfung läuft über ein
  repo-eigenes `verify`-Skript, kein CR an das Fremdwerkzeug.

## 2. Definition of Done

- [ ] Sensor existiert, läuft im `verify`-Aggregat und meldet rot, wenn
      eine Closure-Notiz ohne Trigger-Audit-Zeile zu den aktiven
      `MR`-Einträgen bleibt.
- [ ] **Gegenprobe in beiden Richtungen** im Selbsttest (mit Zeile grün,
      ohne rot; leere Aktiven-Menge grün gemeldet, nicht stumm).
- [ ] Unabhängiger Review, Report unter [`docs/reviews/`](../../../../docs/reviews/README.md).
- [ ] Closure-Notiz mit Steering-Loop-Lerneintrag; Register fortgeschritten;
      jedes Risiko aus §6 trägt einen Ausgang.

`make gates` und `make verify` grün.

## 3. Plan (vor Code)

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| `tools/verify-trigger-audit.sh` | neu | Scannt Closure-Notizen in `done/` auf die Audit-Zeile, gegen die aktive `MR`-Liste aus `harness/conventions.md` |
| `Makefile` + [`harness/README.md`](../../../../harness/README.md) §Sensors | neu | Target, Aggregat-Anschluss, Gate-Index-Zeile — die vier Orte aus slice-204 |

## 4. Trigger

**Start** (`open` → `in-progress`): das WIP-Limit ist frei.

**Rückführungen:** `in-progress` → `next` (zu groß): entfällt — ein Skript
samt Verdrahtung. `in-progress` → `open` (blockiert): entfällt — kein
Blocker absehbar.

## 5. Closure-Trigger

DoD vollständig, `make gates` und `make verify` grün, Closure-Notiz mit
Lerneintrag.

## 6. Risiken und offene Punkte

- **Ceremonie-Gefahr:** eine immer gleiche „0 offen"-Zeile ohne Blick in die
  Einträge wäre eine Formulierung ohne Beobachtung.
  — **Ausgang:** bei Closure.

## 7. Closure-Notiz

*(wird vor dem `git mv` nach `done/` gefüllt)*

## 8. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** GATE (Achsen 1, 2, 3 ✓).

**Vorgelagert — offene Beobachtungen sichten:** das Register wurde
durchgegangen (2026-09-29): [`BEO-HARNESS/mr-aufloesungs-trigger-ohne-waechter`](../observations/BEO-HARNESS/mr-aufloesungs-trigger-ohne-waechter/observation.md)
ist dieser Slice selbst (3×, Ausgang geplant); keine weitere offene
Beobachtung in HARNESS.

**Modus-Begründungsblock:** alle berührten Sub-Areas GF.
