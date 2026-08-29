# slice-101 — Etappe D1: Beobachtungs-Register anlegen

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.
**Autor:** Claude (Opus 5), im Auftrag des Maintainers.
**Berührte Spec-Stellen:** — *(keine; Planungs-Harness ohne Vertragsberührung)*
**Deckt:** keine `AC-*`/`ADR-*`.
**Bezug:** Etappe **D** aus [slice-092 §6](../done/slice-092-regelwerk-v5120-delta-analyse.md),
erste Hälfte. Vorgänger [slice-100](../done/slice-100-closure-sensor-zitat-kontext.md).
[Roadmap](../in-progress/roadmap.md).

---

## 1. Auslöser

`modul-06` §Das Beobachtungs-Register führt eine **stehende** Datei ein:
`docs/plan/planning/observations.md`. Sie ist der Zähler des Steering Loops — und zugleich der
dritte der drei Ausgänge, die `modul-05` für jedes offene Risiko einer Closure verlangt:
*eingetreten* → Carveout oder Folge-Slice · *entfallen* → gestrichen mit Begründung ·
*weiter offen* → **wandert ins Register**.

Genau diesen dritten Ausgang haben die Closure-Notizen von slice-097 bis slice-100 **sechsmal**
benutzt — auf ein Register, das es nicht gibt. Dieser Slice legt es an und trägt sie ein.

**Warum stehend und nicht je Welle:** *„Eine von Closure zu Closure übernommene Sektion hängt an
einer ungebrochenen Kette — vergessene Übernahme setzt den Zähler auf null."* Der feste Ort
streicht das. Für a-check kommt hinzu, dass gerade **keine Welle offen** ist; ein
wellen-getragener Zähler hätte hier überhaupt keinen Träger.

## 2. Betroffene Module

- `docs/plan/planning/observations.md` — neu.
- `harness/conventions.md` §Modus-Deklaration pro Sub-Area — eine fehlende Zeile.
- [`AGENTS.md`](../../../../AGENTS.md) §5 — Register und Eintrags-Zeitpunkt.

Zwei Schichten: Planungs-Harness und Harness-Konventionen.

## 3. Der erste Eintrag erzwingt eine Konventions-Zeile

Die Sub-Area-Spalte des Registers ist nicht frei: *„Steht in der Spalte ein Name, den die
Modus-Deklaration in `harness/conventions.md` nicht führt, ist entweder die Zuordnung falsch oder
die Deklaration unvollständig."*

Vier der sechs Beobachtungen betreffen `AGENTS.md`, `CLAUDE.md` und `harness/` — und **für die
führt die Modus-Deklaration keine Zeile**. Das ist keine Neuigkeit: die Lücke ist seit
[slice-091 §7](../done/slice-091-claude-md-auf-verweis-reduzieren.md) in **vier** Slices benannt
worden, ohne je einen Ort zu haben. Sie ist damit die erste Beobachtung des Registers — und sie
steht bei **4×**, also über der Schwelle von drei. `modul-06` sagt, was dann geschieht: der
Eintrag *„wird zur verkörperten Regel (mit Herkunfts-Anker)"*.

Das Register schließt seine eigene Voraussetzung also im selben Zug. Die neue Sub-Area
**Harness-Einstieg** erfüllt die Qualifikation mit allen drei Inklusions-Achsen: eigene
Adaptions-Linie (die aufgelösten Einträge zur Source Precedence und zum Vendoring lebten genau
dort), eigene Inventur-Linie, eigene Pfad-/Datei-Familie.

## 4. Auszuführende Gates

`make gates` — tragend ist `doc-check`, weil das Register Slice-Kennungen als Belege führt. Zum
Abschluss `make verify`.

**Kein Sensor in diesem Slice.** Die maschinelle Hälfte — jede in `done/` zitierte `BEO-<NNN>` hat
eine Registerzeile, und jede Zeile trägt mindestens einen formgebundenen Beleg — ist **D2**,
zusammen mit den Risiko-Ausgängen. Erst das Register, dann sein Sensor: dieselbe Reihenfolge wie
bei der Closure-Pflicht vor slice-050.

## 5. Was bewusst nicht getan wird

- **Keine erfundenen Belege.** Die Vorlage warnt ausdrücklich davor. Jede der sechs Zeilen nennt
  die Slices, in deren Closure-Notiz die Beobachtung wirklich steht; der Zähler ist deren Anzahl,
  nicht eine Schätzung.
- **`BEO-<NNN>` wird nicht in die ID-Schema-Deklaration nachgetragen.** Sie steht in `MR-000`, und
  Einträge werden nie überschrieben. Die Klasse ist ohnehin Baseline-**Default** (die neue
  `MR-000`-Vorlage führt sie), also keine Adaption. Der Nachtrag gehört in die Überarbeitung der
  ID-Schema-Deklaration — genau der Auflösungs-Trigger, den der Adaptions-Eintrag zur
  ADR-Vorlagen-Version bereits trägt.
- **Keine Antwort auf die eingetragenen Beobachtungen.** Ein Register ist ein Zähler, kein
  Beschluss. Die einzige Ausnahme ist die erste Zeile, weil sie die Schwelle bereits erreicht hat
  (§3).

## 6. DoD

- [ ] `docs/plan/planning/observations.md` steht in der Ziel-Form (beide Tabellen, `— keine —`
      wo leer) und trägt sechs Beobachtungen mit formgebundenen Belegen (`slice-<NNN>`, Anzahl ==
      Zähler) — Beleg: Datei.
- [ ] Die Modus-Deklaration führt die Sub-Area **Harness-Einstieg**; jede Sub-Area-Angabe im
      Register findet sich dort wieder — Beleg: Diff und Abgleich beider Tabellen.
- [ ] [`AGENTS.md`](../../../../AGENTS.md) §5 nennt das Register, seinen Ort und den
      Eintrags-Zeitpunkt (Slice-Closure) — Beleg: Diff.

Pflicht, aber **kein** Liefer-Punkt: `make gates` und zum Abschluss `make verify` grün — Ausgabe
in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.

## 7. Closure-Notiz

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice,
[`AGENTS.md`](../../../../AGENTS.md) §5; `make verify` prüft das.)_

## 8. Sub-Area-Modus

Berührt werden **Planungs-Harness** (`docs/plan/planning/`) und die mit diesem Slice neu
deklarierte **Harness-Einstieg**-Sub-Area. Beide Greenfield.
