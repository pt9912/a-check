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

*(offen — die Messung entsteht mit der Arbeit; Ausgangslage in
[slice-174](../done/slice-174-regelwerk-v650-delta-analyse.md) §3.2)*

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
