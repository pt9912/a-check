# slice-177 — Sensors-Struktur: zwei Tabellen und `harness/sensors/`

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `make slice-mv` ([`AGENTS.md`](../../../../AGENTS.md)
§3.3/§5).

**Welle:** [welle-15](../welle-15-regelwerk-v650-migration.md)

**Bezug:** [slice-174](../done/slice-174-regelwerk-v650-delta-analyse.md)
§3.2, T-2 und T-4 — Etappe **D** des Schnitts in §3.4.

**Berührte Spec-Stellen:** — *(keine)* — Harness-Struktur ohne
Vertragsberührung.

**Verantwortlich:** — *(noch nicht priorisiert)*

**Autor:** Claude (Opus 5), im Auftrag des Maintainers. **Datum:**
2026-09-07.

---

## 1. Ziel und Abgrenzung

Die Sensors-Tabelle wird Index: Nicht-Gates stehen in einer zweiten Tabelle mit `kein Gate` in der Zeile, und ein Gate-Vertrag, der mehr als einen Satz braucht, wandert nach `harness/sensors/<target>.md` — die Target-Zelle wird zum Link darauf.

**Nicht in diesem Slice**, je Punkt mit Grund:

- **Andere Etappen von [welle-15](../welle-15-regelwerk-v650-migration.md)** —
  Schicht-Abgrenzung: jede Etappe misst gegen den vendorten Stand und ist
  einzeln lieferbar.
- **Nachrüsten des Altbestands**, wo die Ziel-Form nur für Neues gilt —
  Bestand bleibt bewusst stehen; ein Sensor gegen unentschiedenen Altbestand
  wäre ein Fehlalarm.

## 2. Analyse (vor der Umsetzung)

**Der Slice ist kein reiner `v6.5.0`-Nachzug.** Die Regel, die a-checks §4-Tabelle
bricht, steht **seit `v5.12.0` unverändert** in
[`AGENTS.template.md`](../../../../.harness/baseline/v6.2.0/templates/AGENTS.template.md)
§4: *„Diese Tabelle **listet auf**; definiert wird hier nichts. Die **Bindung**
eines Targets … steht in `harness/README.md` §Sensors."* Über vier
Baseline-Stände und zwei Adaptions-Durchgänge hinweg nie befolgt, und **keine**
aktive Adaption deckt es. Neu an `v6.5.0` ist nur die *Antwort* — die
`harness/sensors/`-Struktur; die *Regel* ist alt.
Beleg: [`BEO-HARNESS/baseline-regel-nie-erwogen-weil-bestand-sie-verletzt`](../observations/BEO-HARNESS/baseline-regel-nie-erwogen-weil-bestand-sie-verletzt/observation.md)
`evidence/slice-177.md`, damit 2×.

**Ausgangsmessung 2026-09-07** (Zeichen je Vertrags-/Zweck-Zelle, Schwelle 250
als Näherung für „mehr als einen Satz"):

| Ort | über der Schwelle | Spitzenwerte |
|---|---|---|
| [`AGENTS.md`](../../../../AGENTS.md) §4 | **16 von 40** | `doc-check` 1871 · `symlink-check` 1106 · `doc-workflows` 729 |
| [`harness/README.md`](../../../../harness/README.md) §Sensors | **6 von 24** | `doc-check` 639 · `symlink-check` 369 · `gate-consistency` 356 |

**Der eigentliche Fund ist nicht der Überhang, sondern die Doppelung.**
Dieselben Verträge stehen an **zwei** Orten, in AGENTS.md ausführlicher — bei
`doc-check` 1871 gegen 639 Zeichen, bei `symlink-check` 1106 gegen 369. Zwei
Fassungen desselben Vertrags sind zwei Quellen, und Kopien driften.

Die Ziel-Form löst genau das mit: *„Von außen — aus `AGENTS.md`, einer ADR,
einem Slice — wird **diese Datei direkt** adressiert, nicht der Index: sie
wandert nie."* Steht der Vertrag unter `harness/sensors/<target>.md`, tragen
**beide** Tabellen nur noch ihre Index-Zeile, und AGENTS.md §4 verweist auf
dieselbe Datei statt eine zweite Fassung zu führen.

**Zu klären, bevor umgebaut wird:** AGENTS.md §4 führt `| Target | Zweck |` —
die Spalte, in der die Ziel-Form `kein Gate` verlangt, existiert dort nicht
(`harness/README.md` hat sie als *Bindung*). Ob a-check die Spalte ergänzt oder
die Abweichung deklariert, ist die Adaptions-Frage aus
[slice-174](../done/slice-174-regelwerk-v650-delta-analyse.md) §7.

## 3. Umsetzung

*(offen)*

## 4. Definition of Done

- [ ] Nicht-Gates (`slice-mv`, `archive-wave`, `regelwerk-check`) stehen in einer zweiten Tabelle, `kein Gate` in der Zeile statt als Prosa-Vorspann.
- [ ] Für jeden Überhang eine Datei unter `harness/sensors/`, Target-Zelle verlinkt; welche der 16 gemessenen Zellen betroffen sind, ist im Slice begründet — nicht alle 16 pauschal.
- [ ] Unabhängiger Review durchgeführt (Report unter [`docs/reviews/`](../../../reviews/README.md)).
- [ ] `make gates` grün.
- [ ] `make verify` grün.
- [ ] Jedes Risiko trägt einen Ausgang.

## 5. Trigger

**Start** (`open` → `in-progress`): Etappe A (Vendoring) liegt in `done/`,
Maintainer-Freigabe, WIP-Limit frei.

**Rückführungen:** wächst der Umfang über die DoD hinaus, zurück nach `next/`
zur Zerlegung. Ändert sich der adoptierte Stand erneut, zurück nach `open/`.

## 6. Closure-Trigger

DoD vollständig, `make gates`/`make verify` grün, Closure-Notiz geschrieben.
Der Slice trägt ein `**Welle:**`-Feld und archiviert **mit seiner Welle**
([`AGENTS.md`](../../../../AGENTS.md) §6).

## 7. Risiken und offene Punkte

- *Die Ziel-Form wird übernommen, ohne dass a-checks Bestand sie trägt — dann
  steht eine Regel da, die der eigene Bestand bricht* — Ausgang bei Closure;
  Klasse [`BEO-HARNESS/baseline-regel-nie-erwogen-weil-bestand-sie-verletzt`](../observations/BEO-HARNESS/baseline-regel-nie-erwogen-weil-bestand-sie-verletzt/observation.md).

## 8. Closure-Notiz

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice;
Lerneintrag — Form: wird dort benannt.)_

## 9. Sub-Area-Modus

**Vorgelagert — Sub-Area-Wahl prüfen:** entsteht mit dem Übergang nach
`in-progress/`.

**Vorgelagert — offene Beobachtungen sichten:** entsteht mit dem Übergang nach
`in-progress/`; der Register-Stand beim Anlegen ist ein anderer als beim
Beginn der Arbeit.

**Alle berührten Sub-Areas GF** — kein Begründungsblock nötig.
