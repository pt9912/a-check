# slice-102 — Etappe D2: Risiko-Ausgänge und Register-Deckung als Sensor

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.
**Autor:** Claude (Opus 5), im Auftrag des Maintainers.
**Berührte Spec-Stellen:** — *(keine; Verifikations-Schicht ohne Vertragsberührung)*
**Deckt:** keine `AC-*`/`ADR-*`.
**Bezug:** Etappe **D** aus [slice-092 §6](../done/slice-092-regelwerk-v5120-delta-analyse.md),
zweite Hälfte und **letzter Slice der Migration**. Vorgänger
[slice-101](../done/slice-101-beobachtungs-register.md).
[Roadmap](../in-progress/roadmap.md).

---

## 1. Auslöser

Zwei Pflichten der Baseline stehen seit slice-101 im Repo, aber ohne Sensor:

1. **Risiko-Ausgänge** (`modul-05`): jedes notierte Risiko bekommt beim Übergang nach `done/`
   genau **einen** von drei Ausgängen. *„Urteilsfrei ist, **dass** zu jedem notierten Risiko ein
   Ausgang dasteht und **welcher der drei** es ist: Die drei sind eine geschlossene Menge, kein
   Freitext — ein Risiko ohne Ausgang und ein Ausgang, der keiner der drei ist, sind an der
   **Form** erkennbar, nicht am Inhalt."*
2. **Register-Deckung** (`modul-06`): eine in `done/` zitierte `BEO-NNN` hat eine Registerzeile,
   und jede Registerzeile trägt mindestens einen Beleg — **formgebunden** (`slice-NNN`), in der
   Anzahl, die der Zähler nennt.

`modul-06` sagt zur Arbeitsteilung: *„Mensch urteilt, Maschine prüft Deckung."* Ob ein Ausgang
*trägt*, bleibt Urteil. Dass er **dasteht** und **welcher** es ist, ist Form.

## 2. Betroffene Module

- `tools/verify-closure-notes.sh` — Risiko-Ausgänge.
- `tools/verify-observations.sh` — neu, Register-Deckung; hängt an `verify`, nicht an `gates`.
- `Makefile`, [`AGENTS.md`](../../../../AGENTS.md) §4 und
  [`harness/README.md`](../../../../harness/README.md) §Sensors — Deklaration des neuen Targets.

Eine Schicht: Gate-/Werkzeug-Schicht.

## 3. Wo die Prüfung ihre Grenze hat

**Geprüft wird innerhalb eines vorhandenen Risiko-Blocks**, nicht seine Existenz. Ein Slice, der
keine Risiken notiert, braucht keine Ausgänge — `modul-05` bindet die Pflicht an *notierte*
Risiken. Ein Sensor, der einen Block einforderte, verlangte ein Urteil („gab es hier Risiken?")
und produzierte auf jedem alten Slice Rauschen.

**Damit bleibt eine Lücke, und sie gehört benannt:** wer den Block weglässt, wird nicht erwischt.
Das ist dieselbe Klasse ehrlicher Grenze wie bei `verify-slice-form` („höchstens zwei Schichten"
bleibt Ermessen) — und derselbe Grund, warum die semantische Hälfte beim Skill bleibt.

Ebenso ungeprüft bleibt die **Existenz** einer Beleg-Datei: `modul-06` schließt das ausdrücklich
aus, weil ein Repo Slices führen darf, die es nicht als Plan-Datei ablegt. Ein erfundenes
`slice-999` bliebe unentdeckt.

## 4. Auszuführende Gates

`make verify`, dann `make gates`; `doc-targets` prüft die Deklaration des neuen Targets in
**beiden** Doku-Tabellen.

**Negativ-Proben, je Prüfung und in beide Richtungen** — ohne sie wäre jede der neuen Prüfungen
ein Freibrief: Risiko ohne Ausgang feuert · Ausgang außerhalb der Dreier-Menge feuert · drei
gültige Ausgänge schweigen; zitierte `BEO-NNN` ohne Registerzeile feuert · Registerzeile ohne
Beleg feuert · Beleg-Anzahl ≠ Zähler feuert · das echte Register schweigt.

## 5. Was bewusst nicht getan wird

- **Keine Prüfung der Beleg-*Lage*.** `modul-06` erlaubt sie erst **nach** dem `git mv` — auf dem
  Schreib-Commit läge die Datei noch nicht in `done/`, und ein Sensor dort meldete bei jeder
  korrekten Closure rot. Genau diese Reihenfolge-Falle steht schon als `BEO-006` im Register;
  sie hier zu wiederholen wäre absehbar.
- **Keine Antwort auf die offenen Beobachtungen.** Das Register zählt, dieser Slice prüft seine
  Deckung. Beides ist nicht dasselbe wie sie zu beantworten.

## 6. DoD

- [x] `verify-closure-notes` prüft in `done/` jede Zeile eines vorhandenen Risiko-Blocks auf genau
      einen Ausgang aus der geschlossenen Dreier-Menge — Beleg: Target-Ausgabe und Selbsttest.
- [x] `make verify-observations` prüft Register-Deckung (zitierte `BEO-NNN` ⇒ Zeile; Zeile ⇒
      Beleg; Beleg-Form und Anzahl == Zähler) und hängt in `verify` — Beleg: Target-Ausgabe.
- [x] Das neue Target steht in [`AGENTS.md`](../../../../AGENTS.md) §4 **und**
      [`harness/README.md`](../../../../harness/README.md) §Sensors; `doc-targets` grün — Beleg:
      Target-Ausgabe.

Pflicht, aber **kein** Liefer-Punkt: `make gates` und zum Abschluss `make verify` grün — Ausgabe
in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.

## 7. Closure-Notiz

**Geliefert:** zwei Prüfungen für zwei Pflichten, die seit slice-101 im Repo standen und niemand
maß — Risiko-Ausgänge aus der geschlossenen Dreier-Menge und die Deckung des
Beobachtungs-Registers. **Damit ist die Migration `v3.5.2` → `v5.12.0` abgeschlossen.**

**Lerneintrag — Form: neuer Sensor.** Neu beobachtbar sind zwei Dinge: *ein Risiko, dessen Ausgang
nicht aus der Dreier-Menge stammt*, und *eine Registerzeile ohne formgebundenen Beleg*. Was den
ersten Fall lehrreich macht, steht im eigenen Bestand: slice-092 führt den Ausgang
„Maintainer-Entscheidung" — damals völlig vernünftig, denn die geschlossene Menge galt im Repo
noch nicht. **Genau das ist der Wert einer geschlossenen Menge:** ohne sie erfindet jeder Autor
den Ausgang, der gerade passt, und die Frage, ob ein Risiko wirklich versorgt ist, wird wieder zum
Urteil. *Weil* eine Menge nur dann urteilsfrei prüfbar ist, wenn sie abgeschlossen ist, ist der
vierte Ausgang kein Detail, sondern der Unterschied zwischen Form und Ermessen.

**Zwei beobachtbare Closure-Kriterien:**

1. `make verify-observations` läuft mit Exit 0 über **sieben** Beobachtungen; sein Selbsttest
   fährt vier Register-Fixtures (gut · ohne Beleg · Anzahl ungleich Zähler · Freitext-Beleg) und
   eine Zitat-Fixture. Ohne die erste wäre ein Muster, das alles durchlässt, nicht erkennbar.
2. Der Selbsttest von `verify-closure-notes` fährt **drei** Risiko-Richtungen — gültige Ausgänge
   schweigen, fehlender Ausgang feuert, Ausgang außerhalb der Menge feuert — plus die
   Grandfathering-Probe unter dem Stichtag.

**Offene Risiken und ihr Ausgang:**

- *Ein Slice, der den Risiko-Block ganz weglässt, wird nicht geprüft* — Ausgang: **weiter offen**,
  als `BEO-007` im Beobachtungs-Register.
- *Die Lage und Existenz der Beleg-Datei bleiben ungeprüft* — Ausgang: **weiter offen**, gedeckt
  durch `BEO-006` im Beobachtungs-Register; `modul-06` schließt die Prüfung vor dem `git mv`
  ausdrücklich aus.
- *Vor dem Stichtag stehen Ausgänge außerhalb der Dreier-Menge* — Ausgang: **gestrichen mit
  Begründung**. Sie entstanden, bevor die Menge im Repo galt; rückwirkendes Umschreiben wäre
  Geschichts-Politur, und die Grandfathering-Probe im Selbsttest hält die Stufung in beide
  Richtungen fest.

- *Ein drittes zu weites Muster ist beim Bauen aufgefallen und mitgefixt* — `verify-slice-form`
  suchte `make (gates|verify)` als Präfix und beanstandete damit **diesen** Slice: sein DoD-Punkt
  liefert das Target `make verify-observations`, was eine Lieferung ist und kein Gate-Lauf. Das
  Muster endet jetzt am Target-Namen, mit Gegenprobe im Selbsttest. Ausgang: **gestrichen mit
  Begründung** — behoben, samt Fixture, die den Fall künftig hält.

**Folge-Slices:** keine — die Migration ist abgeschlossen. Was offen bleibt, steht im Register.

## 8. Sub-Area-Modus

Berührt wird die **Gate-/Werkzeug-Schicht** (`tools/`, `Makefile`) — Greenfield.
