# slice-098 — Etappe C3: Slice-Form auf Liefer-Punkte und Kopffelder

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.
**Autor:** Claude (Opus 5), im Auftrag des Maintainers.
**Berührte Spec-Stellen:** — *(keine; Harness-/Sensor-Änderung ohne Vertragsberührung)*
**Deckt:** keine `AC-*`/`ADR-*`.
**Bezug:** Etappe **C** aus [slice-092 §6](../done/slice-092-regelwerk-v5120-delta-analyse.md),
dritte Hälfte. Vorgänger [slice-097](../done/slice-097-adaptions-urteile-ausfuehren.md).
[Roadmap](../in-progress/roadmap.md).

---

## 1. Auslöser

`modul-05` §Ziel-Form: Slice hat die Größen-Metrik ausgetauscht. Aus *„höchstens drei
**DoD**-Punkte"* ist *„höchstens drei **Liefer**-Punkte"* geworden, mit einer expliziten
Nicht-Zähl-Liste:

> gezählt wird nur, was mit dem Umfang wächst (die Artefakte und Akzeptanzkriterien dieses Slice).
> Nicht gezählt: Gate-Läufe, Closure-Notiz, Register, Risiko-Ausgänge, die drei Paarungen; sie sind
> pro Slice konstant und sagen über die Größe nichts.

`make verify-slice-form` zählt heute Checkboxen im DoD-Abschnitt. **Jeder** Slice seit slice-052
verbrennt davon einen auf `make gates` grün — einen Punkt, der laut neuer Regel gar nicht zählt.
Der Sensor misst damit systematisch ein Drittel zu streng.

Dazu verlangt der Kopf drei neue Felder: `Verantwortlich:`, `Autor:` und die berührten
Spec-Stellen. Dieser Slice trägt sie bereits — er ist der erste unter der neuen Form.

## 2. Betroffene Module

- [`docs/plan/planning/slice.template.md`](../slice.template.md) — Kopffelder, Größen-Regel,
  DoD-Form.
- `tools/verify-slice-form.sh` — zwei neue Prüfungen, je mit Negativ-Probe.
- [`AGENTS.md`](../../../../AGENTS.md) §4 (Target-Zweck) und §5 (Slice-Form-Regel).

Zwei Schichten: Planungs-Vorlage und Gate-/Werkzeug-Schicht.

## 3. Wie die neue Metrik prüfbar wird

Die Nicht-Zähl-Liste ist **keine** Ermessensfrage, wenn die Vorlage die konstanten Punkte gar
nicht erst als Checkbox führt. Der Gate-Lauf wird deshalb aus dem DoD herausgenommen und steht als
feste Zeile darunter — er ist Pflicht, aber kein Liefer-Punkt.

Damit bleibt die Zähl-Logik des Sensors unverändert richtig und misst ab sofort das Richtige. Neu
hinzu kommen zwei Prüfungen:

1. **Kein Gate-Lauf als DoD-Punkt.** Eine Checkbox im DoD, die `make gates` oder `make verify`
   nennt, ist ein Befund — sonst kehrt die alte Gewohnheit zurück und frisst weiter einen Slot.
2. **Kopffelder vorhanden.** `Verantwortlich:`, `Autor:` und das Spec-Stellen-Feld müssen im Kopf
   stehen; `—` ist eine gültige Antwort, Schweigen nicht.

Beide gelten **ab slice-098** — die 46 Slices ab slice-052 tragen alle einen Gate-Lauf im DoD und
keine Kopffelder. Sie rückwirkend umzuschreiben wäre Geschichts-Politur; dieselbe
Grandfathering-Mechanik nutzt der Sensor seit slice-052.

## 4. Auszuführende Gates

`make verify-slice-form` (mit erweitertem Selbsttest), dann `make gates`, zum Abschluss
`make verify`. `doc-targets` prüft die Übereinstimmung `AGENTS.md` §4 ↔ Makefile mit.

**Negativ-Proben, ohne die die neuen Prüfungen tot wären:** je eine Fixture mit Gate-Lauf im DoD
und eine ohne Kopffelder müssen den Sensor **rot** machen, und eine konforme Fixture muss
schweigen. Der bestehende Selbsttest hat diese Struktur bereits je Befundklasse; die zwei neuen
reihen sich ein.

## 5. Was bewusst nicht getan wird

- **Die 46 bestehenden Slices bleiben unangetastet.** Grandfathering ist die Antwort, nicht ein
  Massen-Umbau.
- **„Höchstens zwei Schichten" bleibt ungeprüft.** Was eine Schicht ist, ist Ermessen über
  Modul-Grenzen; ein Zähler darüber wäre Schein-Genauigkeit. Das stand schon vor diesem Slice so
  im Sensor und ändert sich nicht.
- **Der Rest der Form ist nicht hier.** `## Leseordnung` in `harness/README.md`, die fünf
  Template-Provenienz-Zeiger und der Fall des alten vendored Baums sind **C4** — sie waren in
  meiner Etappen-Aufzählung nach C1 versehentlich untergegangen und werden hier ausdrücklich
  wieder aufgenommen, statt still zu verschwinden.

## 6. DoD

- [ ] Vorlage trägt die drei Kopffelder und die Größen-Regel in Liefer-Punkten; der Gate-Lauf
      steht als feste Zeile statt als Checkbox — Beleg: Diff.
- [ ] `verify-slice-form` prüft beides ab slice-098, der Selbsttest feuert für **beide** neuen
      Befundklassen und schweigt für die konforme Fixture — Beleg: Target-Ausgabe.
- [ ] [`AGENTS.md`](../../../../AGENTS.md) §4 und §5 nennen die neue Metrik — Beleg: Diff.

Pflicht, aber kein Liefer-Punkt: `make gates` und zum Abschluss `make verify` grün, **Ausgabe in
eine Datei**, Exit-Code getrennt geprüft, nie in eine Pipe.

## 7. Closure-Notiz

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice,
[`AGENTS.md`](../../../../AGENTS.md) §5; `make verify` prüft das.)_

## 8. Sub-Area-Modus

Berührt werden **Planungs-Harness** (`docs/plan/planning/`) und die **Gate-/Werkzeug-Schicht**
(`tools/`) — beide in der Modus-Deklaration pro Sub-Area als Greenfield geführt.
