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

- [x] Sensor existiert, läuft im `verify`-Aggregat und meldet rot, wenn
      eine Closure-Notiz ohne Trigger-Audit-Zeile zu den aktiven
      `MR`-Einträgen bleibt.
- [x] **Gegenprobe in beiden Richtungen** im Selbsttest (mit Zeile grün,
      ohne rot; leere Aktiven-Menge grün gemeldet, nicht stumm).
- [x] Unabhängiger Review, Report unter [`docs/reviews/`](../../../../docs/reviews/README.md).
- [x] Closure-Notiz mit Steering-Loop-Lerneintrag; Register fortgeschritten;
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
  — **Ausgang:** *weiter offen* → Beobachtungs-Register:
  [`BEO-HARNESS/trigger-audit-ceremonie`](../observations/BEO-HARNESS/trigger-audit-ceremonie/observation.md)
  (neu, 1× — die Grenze ist in der Sensor-Datei deklariert, das Urteil bleibt
  beim Menschen).

## 7. Closure-Notiz

**Lerneintrag — Form: neuer Sensor.** [`tools/verify-trigger-audit.sh`](../../../../tools/verify-trigger-audit.sh)
(im `verify`-Aggregat; `seit slice-208`) hält jede Closure ab slice-208 an die
Sichtungs-Zeile „Trigger-Audit der aktiven MR:" — die wächterlose Auflösung
bekommt einen formulierten Beleg pro Closure.

**Was hat funktioniert:** das Vorgänger-Muster aus slice-204 — Sensor-Skript,
vier Orte der Verdrahtung (Target, `.PHONY`, GATES-Liste, Gate-Index) und
Selbsttest in beiden Richtungen samt Grandfathering und leeren
Aktiven-Mengen-Fall, bevor die Übergabe.

**Was ging anders als geplant:** der unabhängige Review fand die Ceremonie-
Gefahr als nicht aufgelöstes Risiko (der §6-Ausgang „bei Closure" war ein
Zeitpunkt statt eines Ausgangs) — sie geht als Restrisiko ins Register und
bleibt deklarierte Grenze des Sensors.

**Steering-Loop-Eintrag:** siehe Lerneintrag oben. Die Beobachtungs-Klasse
[`BEO-HARNESS/mr-aufloesungs-trigger-ohne-waechter`](../observations/BEO-HARNESS/mr-aufloesungs-trigger-ohne-waechter/observation.md)
erreichte **3×** — Ausgang *geplant* →
slice-208, durch diesen
Slice **verkörpert** (`seit slice-208`).

**Beobachtungs-Register (`../observations/`):**
`mr-aufloesungs-trigger-ohne-waechter` → Ausgang *verkörpert*
([evidence/slice-206.md](../observations/BEO-HARNESS/mr-aufloesungs-trigger-ohne-waechter/evidence/slice-206.md)
war die 3. Auflage); `trigger-audit-ceremonie` → neu, 1×
([evidence/slice-208.md](../observations/BEO-HARNESS/trigger-audit-ceremonie/evidence/slice-208.md)).

**Folge-Slices:** keine.

**Risiken aus §6:** das eine Risiko trägt seinen Ausgang (*weiter offen* →
Register — Ceremonie-Gefahr, neu 1×).

**Drei Paarungen:** Anker — verkörpert (Sensor, `seit slice-208`) ·
Folge-Slice — keine genannt · Register —
mr-aufloesungs-trigger (3×, verkörpert) · trigger-audit-ceremonie (1×).

**Trigger-Audit der aktiven MR:** MR-016 MR-019 MR-024 MR-025 MR-027
MR-028 MR-029 MR-030 — 0 offen (geprüft 2026-09-29, Konventions-Stand
dieses Laufs).

## 8. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** GATE (Achsen 1, 2, 3 ✓).

**Vorgelagert — offene Beobachtungen sichten:** das Register wurde
durchgegangen (2026-09-29): [`BEO-HARNESS/mr-aufloesungs-trigger-ohne-waechter`](../observations/BEO-HARNESS/mr-aufloesungs-trigger-ohne-waechter/observation.md)
ist dieser Slice selbst (3×, Ausgang geplant); keine weitere offene
Beobachtung in HARNESS.

**Modus-Begründungsblock:** alle berührten Sub-Areas GF.
