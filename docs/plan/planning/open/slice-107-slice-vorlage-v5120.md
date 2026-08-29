# slice-107 — Slice-Vorlage auf die `v5.12.0`-Form, samt beider Sensoren

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.
**Autor:** Claude (Opus 5), im Auftrag des Maintainers.
**Berührte Spec-Stellen:** — *(keine)*
**Deckt:** keine `AC-*`/`ADR-*`.
**Bezug:** Befund §5 aus [slice-105](../done/slice-105-form-review-nachholen.md); korrigiert
[slice-098](../done/slice-098-slice-form-liefer-punkte.md). [Roadmap](../in-progress/roadmap.md).

---

## 1. Auslöser — und eine Korrektur an mir selbst

Die vendored Vorlage führt acht Abschnitte statt sieben. Beim Lesen des **vollen** Textes — nicht
nur der Überschriften — zeigt sich ein Widerspruch zu Arbeit von heute:

**Die Baseline behält den Gate-Lauf als DoD-Checkbox.** Ihr §2 listet wörtlich
`- [ ] make gates grün.` und schreibt daneben: *„Gezählt wird nur, was mit dem Umfang wächst — die
Gate-Läufe und die vier Closure-Pflichten darunter zählen **nicht** mit."*

[slice-098](../done/slice-098-slice-form-liefer-punkte.md) hat das anders gelöst: der Gate-Lauf
wurde aus dem DoD **herausgenommen**, damit der Zähler nur Zählbares vorfindet, und Prüfung (3)
meldet seither einen Gate-Lauf im DoD als Befund. Das war **meine** Konstruktion, nicht die der
Baseline — und sie würde **jeden** Slice beanstanden, der aus der Ziel-Form entsteht.

Damit ist auch meine Größenschätzung von vorhin falsch, und zwar in beide Richtungen: erst hatte
ich „größerer Brocken" behauptet, dann auf Nachfrage „kleiner als die Roadmap" — beides ohne die
Vorlage gelesen zu haben. Sie ist **größer als beides**, weil sie Begriffe mitbringt, die a-check
nicht führt.

## 2. Betroffene Module

- `docs/plan/planning/slice.template.md` — Übersetzung.
- `tools/verify-slice-form.sh` — DoD-Erkennung, Rückbau von Prüfung (3).
- `tools/verify-closure-notes.sh` — Risiko-Block als eigene Sektion.

Zwei Schichten: Planungs-Harness und Gate-/Werkzeug-Schicht.

## 3. Der Rückbau von Prüfung (3) ist keine Lockerung

Er sieht aus wie eine — eine Prüfregel verschwindet, und
[`AGENTS.md`](../../../../AGENTS.md) §3.6 verlangt für Lockerungen eine ADR. Er ist aber etwas
anderes: die Regel prüfte eine **Bedingung, die die Baseline nicht stellt**. Ein Sensor, der die
Ziel-Form beanstandet, ist kein strengerer Wächter, sondern ein falscher.

Was an ihre Stelle tritt, ist die Regel der Baseline selbst: der Zähler ignoriert die konstanten
Punkte, statt sie zu verbieten. **Das ist strenger, nicht schwächer** — er zählt dann auch in
Slices richtig, die die konstanten Punkte mitführen, was Prüfung (3) gar nicht erst zuließ.

## 4. Auszuführende Gates

`make verify` (beide Sensoren hängen dort), dann `make gates`.

**Negativ-Proben, je Richtung:** neue Gliederung mit vier Liefer-Punkten muss feuern · dieselbe
mit dreien plus Gate-Lauf und Closure-Pflichten muss schweigen · alte Gliederung unter dem
Stichtag unverändert · Risiko ohne Ausgang in der **eigenen Sektion** muss feuern · gültige
Ausgänge dort müssen schweigen.

## 5. Was bewusst nicht getan wird

- **Kein bestehender Slice wird umgeschrieben.** `modul-02` für wiederkehrende Vorlagen:
  *„Neue Instanzen folgen der neuen Form, bestehende werden nicht rückwirkend umgeschrieben."*
  Die 54 Slices ab slice-052 bleiben in ihrer Form, geschützt durch einen dritten Stichtag.
- **Vier Begriffe der Vorlage werden nicht adoptiert:** das **Reconciliation-Register** (a-check
  hat keinen Brownfield-Bootstrap, die Vorlage nennt den Fall selbst), die **drei Paarungen**, der
  **Herkunfts-Anker** und das **`Welle:`-Feld**. Sie brauchen je eine eigene Entscheidung; sie hier
  mitzunehmen hieße, vier Mechaniken ungeprüft einzuführen. Als Beobachtung ausgewiesen.
- **Die drei Lerneintrag-Formen bleiben.** Die vendored §7 kennt sie nicht als Feld, `modul-05`
  nennt sie unverändert — und `verify-closure-notes` prüft sie seit slice-050. Die Vorlage
  ergänzt hier, sie ersetzt nicht.

## 6. DoD

- [ ] `slice.template.md` trägt die acht Abschnitte der Ziel-Form, mit Regelwerk-Zeiger je
      Sektion; a-checks eigene Regeln bleiben, wo die Baseline schweigt.
- [ ] `verify-slice-form` erkennt beide Gliederungen (Stichtag), zählt Liefer-Punkte **ohne** die
      konstanten Posten, und Prüfung (3) ist zurückgebaut.
- [ ] `verify-closure-notes` findet den Risiko-Block auch als eigene Sektion; alle Fixtures aus §4
      verhalten sich wie beschrieben.

Pflicht, aber **kein** Liefer-Punkt: `make gates` und zum Abschluss `make verify` grün — Ausgabe
in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.

## 7. Closure-Notiz

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice,
[`AGENTS.md`](../../../../AGENTS.md) §5; `make verify` prüft das.)_

## 8. Sub-Area-Modus

Berührt werden **Planungs-Harness** und **Gate-/Werkzeug-Schicht** — beide Greenfield.
