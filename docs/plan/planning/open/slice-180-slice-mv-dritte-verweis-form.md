# slice-180 — `slice-mv` lernt die dritte Verweis-Form

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `make slice-mv` ([`AGENTS.md`](../../../../AGENTS.md)
§3.3/§5).

**Welle:** ohne Welle (der Closure-Trigger wäre die eigene DoD — kein
repo-weites Mehr).

**Bezug:** Lese-Schritt der Closure von
[welle-15](../done/welle-15-regelwerk-v650-migration.md) —
[`BEO-PLAN/verweis-auf-wandernden-slice`](../observations/BEO-PLAN/verweis-auf-wandernden-slice/observation.md)
steht bei **7×**, und die Verkörperung hat eine Lücke, die allein in dieser
Welle **zehnmal** aufgetreten ist.

**Berührte Spec-Stellen:** — *(keine)* — Werkzeug ohne Vertragsberührung.

**Verantwortlich:** — *(noch nicht priorisiert)*

**Autor:** Claude (Opus 5), im Auftrag des Maintainers. **Datum:**
2026-09-07.

---

## 1. Ziel und Abgrenzung

`make slice-mv` zieht auch die dritte im Bestand vorkommende Verweis-Form nach:
`<lifecycle-verzeichnis>/slice-NNN-….md` **ohne** `../`-Präfix, wie sie eine
flach unter `docs/plan/planning/` liegende Welle-Datei schreibt.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Die Verweise *in* der wandernden Datei** — es wäre ein anderer Vorgang und
  ist bereits gedeckt: `doc-check`s `links.resolve-from` trägt die
  Gegenrichtung.
- **Ein Sensor auf die Verweis-Form selbst** — Bestand bleibt bewusst stehen:
  `doc-check` meldet den toten Verweis bereits nach dem `mv`; was fehlt, ist
  nicht die Meldung, sondern das Nachziehen.

## 2. Analyse (vor der Umsetzung)

**Gemessen über welle-15:** Bei **jedem** der zehn Lifecycle-Wechsel ihrer sechs
Slices ließ das Werkzeug den Verweis in der flach liegenden Welle-Datei stehen —
`[slice-NNN](open/…)` bzw. `(in-progress/…)`. Jedes Mal meldete `make doc-check`
`target-missing`, jedes Mal war die Korrektur ein `sed`.

**Warum die Form vorher nicht vorkam:** `welle-14` lag bei den Übergängen ihrer
Slices bereits in `done/` — ihre Verweise trugen `../done/…`. Wellenlose Slices
haben gar keine Welle-Datei, die auf sie zeigt. Erst eine **offene** Welle-Datei,
die auf Slices in Lifecycle-Verzeichnissen zeigt, erzeugt die dritte Form; sie
ist damit dieselbe Klasse wie
[`BEO-GATE/muster-trifft-nur-die-haeufige-schreibweise`](../observations/BEO-GATE/muster-trifft-nur-die-haeufige-schreibweise/observation.md):
kalibriert an den Vorkommen, die es beim Bauen gab.

**Zu klären vor dem Umbau:** ob `tools/slice-mv.sh` die Form über eine dritte
Regel ergänzt oder ob eine Verallgemeinerung beide bestehenden ablöst — die
zweite Variante ist kürzer und riskanter.

## 3. Umsetzung

*(offen)*

## 4. Definition of Done

- [ ] `make slice-mv` zieht die dritte Form nach; ein Lifecycle-Wechsel eines
      Slice mit `**Welle:**`-Feld lässt `make doc-check` grün, **ohne**
      Handgriff danach.
- [ ] Der Selbsttest des Werkzeugs deckt alle **drei** Formen ab, je Richtung
      eine Mutations-Probe.
- [ ] Unabhängiger Review durchgeführt (Report unter [`docs/reviews/`](../../../reviews/README.md)).
- [ ] `make gates` grün.
- [ ] `make verify` grün.
- [ ] Jedes Risiko trägt einen Ausgang.

## 5. Trigger

**Start** (`open` → `in-progress`): Maintainer-Freigabe und WIP-Limit frei.
Dringlichkeit steigt mit der nächsten offenen Welle — ohne sie kommt die Form
nicht vor.

**Rückführungen:** zeigt sich, dass eine Verallgemeinerung die zwei bestehenden
Formen bricht, zurück nach `next/` und getrennt schneiden.

## 6. Closure-Trigger

DoD vollständig, `make gates`/`make verify` grün, Closure-Notiz geschrieben.
Danach Archivierung als wellenloser Slice ([`AGENTS.md`](../../../../AGENTS.md) §6).

## 7. Risiken und offene Punkte

- *Die dritte Regel deckt eine vierte nicht, die beim nächsten Betriebsmodus
  entsteht — dieselbe Kalibrierungs-Lücke eine Ebene höher* — Ausgang bei
  Closure.

## 8. Closure-Notiz

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice;
Lerneintrag — Form: wird dort benannt.)_

## 9. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** entsteht mit dem Übergang nach
`in-progress/`.

**Vorgelagert — offene Beobachtungen sichten:** entsteht mit dem Übergang nach
`in-progress/`.

**Alle berührten Sub-Areas GF** — kein Begründungsblock nötig.
