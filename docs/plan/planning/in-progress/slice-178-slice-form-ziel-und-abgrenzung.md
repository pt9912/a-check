# slice-178 — Slice-Form: §1 *Ziel und Abgrenzung*, §8-Titel, Schritt 4

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `make slice-mv` ([`AGENTS.md`](../../../../AGENTS.md)
§3.3/§5).

**Welle:** [welle-15](../welle-15-regelwerk-v650-migration.md)

**Bezug:** [slice-174](../done/slice-174-regelwerk-v650-delta-analyse.md)
§3.2 T-3 — Etappe **E** des Schnitts in §3.4.

**Berührte Spec-Stellen:** — *(keine)* — Planungs-Form ohne
Vertragsberührung.

**Verantwortlich:** — *(noch nicht priorisiert)*

**Autor:** Claude (Opus 5), im Auftrag des Maintainers. **Datum:**
2026-09-07.

---

## 1. Ziel und Abgrenzung

Die Kopieranleitung für neue Slices nennt §1 als *Ziel und Abgrenzung* mit
Out-of-Scope je Punkt und Begründung in vier Klassen, §8 unter seinem neuen
Titel *Sub-Area-Prüfungen und Modus-Begründung* — und der Minimal Agent
Workflow trägt die Schritt-Hälfte derselben Regel: die Plan-Ausgabe in
Schritt 4 nennt Out-of-Scope, und was der Lauf darüber hinaus mitnimmt, ist
eine Plan-Änderung vor dem Code.

**Nicht in diesem Slice**, je Punkt mit Grund:

- **Nachrüsten der bestehenden Slices in `open/`** — Bestand bleibt bewusst
  stehen: die Ziel-Form gilt für neue Slices, und ein Sensor gegen
  unentschiedenen Altbestand wäre ein Fehlalarm.
- **Ein Sensor auf die Out-of-Scope-Form** — wäre ein anderer Vorgang
  (Gate-Schicht statt Planungs-Form) und braucht erst einen Bestand, an dem
  er kalibriert werden kann.
- **Andere Etappen von [welle-15](../welle-15-regelwerk-v650-migration.md)** —
  Schicht-Abgrenzung.

## 2. Analyse (vor der Umsetzung)

*(offen — Ausgangslage in
[slice-174](../done/slice-174-regelwerk-v650-delta-analyse.md) §3.2 T-3)*

## 3. Umsetzung

*(offen)*

## 4. Definition of Done

- [ ] [`AGENTS.md`](../../../../AGENTS.md) §5 nennt beim Kopieren der Ziel-Form
      §1 als *Ziel und Abgrenzung* samt der vier Out-of-Scope-Klassen und dem
      neuen §8-Titel.
- [ ] [`AGENTS.md`](../../../../AGENTS.md) §6 trägt die Schritt-Hälfte: die
      Plan-Ausgabe in Schritt 4 nennt Out-of-Scope, Abweichung ist eine
      Plan-Änderung vor dem Code.
- [ ] Unabhängiger Review durchgeführt (Report unter [`docs/reviews/`](../../../reviews/README.md)).
- [ ] `make gates` grün.
- [ ] `make verify` grün.
- [ ] Jedes Risiko trägt einen Ausgang.

## 5. Trigger

**Start** (`open` → `in-progress`): der neue Stand liegt vendored
([slice-175](../done/slice-175-etappe-a-vendoring-v650.md) in `done/`),
Maintainer-Freigabe, WIP-Limit frei.

**Rückführungen:** wächst der Umfang über die zwei Punkte hinaus, zurück nach
`next/`. Ändert sich der adoptierte Stand erneut, zurück nach `open/`.

## 6. Closure-Trigger

DoD vollständig, `make gates`/`make verify` grün, Closure-Notiz geschrieben.
Der Slice trägt ein `**Welle:**`-Feld und archiviert **mit seiner Welle**
([`AGENTS.md`](../../../../AGENTS.md) §6).

## 7. Risiken und offene Punkte

- *Die Kopieranleitung wächst, ohne dass ein Lauf sie liest — sie steht in
  `AGENTS.md` §5, und die Ziel-Form liegt vendored daneben* — Ausgang bei
  Closure.

## 8. Closure-Notiz

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice;
Lerneintrag — Form: wird dort benannt.)_

## 9. Sub-Area-Modus

**Vorgelagert — Sub-Area-Wahl prüfen:** entsteht mit dem Übergang nach
`in-progress/`.

**Vorgelagert — offene Beobachtungen sichten:** entsteht mit dem Übergang nach
`in-progress/`.

**Alle berührten Sub-Areas GF** — kein Begründungsblock nötig.
