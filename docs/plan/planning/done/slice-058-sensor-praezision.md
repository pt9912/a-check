# slice-058 — Sensor-Präzision: drei belegte Ungenauigkeiten in `tools/`

**Status:** der Zustand ist das **Verzeichnis** dieser Datei
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5) — dieses Feld führt ihn bewusst **nicht** doppelt.
Ein Wert hier wäre eine zweite Quelle für dieselbe Wahrheit und stünde nach jedem `git mv` falsch
da (Review 2026-07-26, Etappe B F-5).
**Deckt:** die Findings **F-3**, **F-4**, **F-5** aus
[`docs/reviews/2026-07-26-slice-049-mechanik-sensoren.md`](../../../reviews/2026-07-26-slice-049-mechanik-sensoren.md)
und **F-3** aus
[`docs/reviews/2026-07-26-etappe-d-slice-052-053-054.md`](../../../reviews/2026-07-26-etappe-d-slice-052-053-054.md).
**Bezug:** Review-Serie zur Migrations-Kette B–F (2026-07-26), erster von drei Fix-Schnitten.
[Roadmap](../in-progress/roadmap.md) — der Verweis ist bewusst **zustandsunabhängig** geschrieben
(`../in-progress/…` statt `roadmap.md`): so löst er aus `in-progress/` **und** aus `done/` auf und
überlebt den Lifecycle-`git mv`. Guide-Kandidat 1 aus SL-002.

---

## 1. Auslöser

Drei Befunde an sonst funktionierenden Sensoren, jeder mit einem Lauf belegt:

| Fund | Beobachtung (gemessen) |
|---|---|
| **R-049-F3** | Der Selbsttest von `suppression-check.sh` behauptet Fehlalarm-Freiheit („eine ohne MUSS still bleiben"), prüft sie aber gegen eine Fixture, die das Muster `//[[:space:]]*(nolint\|lint:ignore)` **per Konstruktion** nicht treffen kann (`nolint` ohne `//`-Präfix). Real reproduziert: die Kommentarzeile `// Probe D: … OHNE //nolint — …` wurde als Direktive gemeldet. |
| **R-049-F4** | `regelwerk-check.sh` wählt den zu prüfenden Baseline-Stand per `find … \| sort \| tail -1` — lexikografisch. Belegt: `printf 'v3.10.0\nv3.5.2\n' \| sort \| tail -1` → **`v3.5.2`**. Bei zwei vendored Ständen prüft das Target den falschen und meldet „Integrität ok". |
| **R-052-F3** | `verify-slice-form.sh` deklariert „höchstens 3 **DoD**-Punkte", zählt aber jede Checkbox-Zeile der ganzen Datei. Heute äquivalent (über slice-052…057 geprüft, keine Abweichung), also latent. |

Gemeinsamer Nenner: alle drei sind **fail-closed** — keiner erzeugt ein False-Green, jeder kann
aber Rauschen erzeugen oder am falschen Gegenstand grün melden. Ein Sensor, der rauscht, wird
abgeschaltet statt repariert (slice-057, SL-001).

## 2. Betroffene Module

- [`tools/suppression-check.sh`](../../../../tools/suppression-check.sh) — Erkennungs-Muster,
  Negativ-Fixture, Kommentar-Korrektur (R-049-F3, R-049-F5).
- [`tools/regelwerk-check.sh`](../../../../tools/regelwerk-check.sh) — Stand-Auswahl (R-049-F4).
- `tools/verify-slice-form.sh` — Abschnitts-Schnitt der
  DoD-Zählung (R-052-F3).

**Eine Schicht** (Gate-/Werkzeug-Schicht). Kein `AC-*`, kein `SPEC-*`, keine ADR berührt: die
Entscheidungen stehen bereits ([ADR-0005](../../adr/0005-lint-profil.md),
[MR-006](../../../../harness/conventions.md#mr-006--baseline-committet-vendored-statt-per-url-referenziert)),
korrigiert wird ihre Umsetzung. Keine Lockerung im Sinne von [`AGENTS.md`](../../../../AGENTS.md)
§3.6 — alle drei Änderungen machen die Sensoren **genauer**, nicht nachsichtiger.

## 3. Auszuführende Gates

`make gates` (enthält `suppression-check`) und `make verify` (enthält `verify-slice-form`),
Ausgabe je in eine Datei, Exit-Code getrennt geprüft.

Je Fix eine **Negativ-Probe**, weil ein Sensor ohne Probe, die ihn rot macht, ein toter Sensor ist:

1. **`suppression-check`** — die echte Direktive muss weiter gefunden werden (rot), die
   Kommentarzeile, die `//nolint` nur *erwähnt*, muss ab jetzt schweigen. Beides als
   Selbsttest-Fixture, damit die Probe dauerhaft mitläuft statt einmalig zu belegen.
2. **`regelwerk-check`** — die Auswahl muss bei zwei Ständen den höheren nennen; Probe über eine
   temporäre zweite Verzeichnis-Fixture.
3. **`verify-slice-form`** — eine Datei mit 2 DoD-Punkten und zusätzlichen Checkboxen außerhalb
   des DoD-Abschnitts muss ab jetzt schweigen, eine mit 4 DoD-Punkten weiter rot werden.

## 4. Was bewusst nicht getan wird

- **Keine Erweiterung des `suppression-check`-Gegenstands.** Er scannt weiterhin nur `./internal`
  und `./cmd`; die Grenze ist dokumentiert und nicht Teil dieser Befunde.
- **Kein Umbau der `regelwerk-check`-Integritätshälfte in ein Gate.** Ob sie in `make gates`
  gehört, ist die in [slice-049 §3](../done/slice-049-mechanik-sensoren.md) bewusst vertagte
  eigene Entscheidung.
- **Kein Fix an `make verify`s Befund-Maskierung** und **nichts an der Guard-Gate-Liste** — beides
  ist Durchsetzungs-/Aggregat-Ebene und gehört in den zweiten Fix-Schnitt, sonst überschreitet
  dieser Slice seine eine Schicht.

## 5. DoD

- [x] `suppression-check` meldet die Direktive weiter und die bloße Erwähnung nicht mehr; beide
      Fälle als Selbsttest-Fixture, belegt durch einen Lauf gegen eine Datei, die beide Zeilen
      trägt (R-049-F3, R-049-F5).
- [x] `regelwerk-check` wählt den Stand nach Versions- statt Zeichenordnung; belegt durch eine
      Probe mit zwei vendored Verzeichnissen (R-049-F4).
- [x] `verify-slice-form` zählt nur Checkboxen im DoD-Abschnitt; belegt durch eine Fixture mit
      Checkboxen außerhalb, die schweigen muss, und eine mit 4 DoD-Punkten, die rot bleibt
      (R-052-F3). `make gates` und `make verify` grün — Ausgabe in eine Datei, Exit-Code getrennt
      geprüft, nie in eine Pipe.

## 6. Closure-Notiz

**Geliefert:** drei präzisierte Sensoren in `tools/`. `suppression-check` unterscheidet jetzt eine
Direktive von ihrer bloßen Erwähnung (Auswertung des Textes nach dem **ersten** `//` einer Zeile);
`regelwerk-check` wählt den vendored Stand nach Versionsordnung und macht einen zweiten Stand
sichtbar, statt ihn stillschweigend zu übergehen; `verify-slice-form` zählt DoD-Punkte nur noch im
DoD-Abschnitt. Jede Änderung trägt ihre Negativ-Probe als dauerhafte Selbsttest-Fixture, nicht als
einmaligen Beleg.

**Lerneintrag — Form: geschärfte Regel.**
> **Eine Negativ-Probe, die den Prüfgegenstand gar nicht treffen kann, belegt nichts — sie
> beruhigt nur.** Der Selbsttest von `suppression-check` behauptete Fehlalarm-Freiheit und prüfte
> sie an der Zeile `// ein gewoehnlicher Kommentar ueber nolint-Regeln`: darin fehlt das
> `//`-Präfix, das Muster konnte sie also nie melden. Ihr Schweigen war kein Beleg, sondern eine
> Tautologie — *weil* eine Fixture, die strukturell außerhalb des Musters liegt, dieselbe Antwort
> gibt wie ein kaputtes Muster. Bemerkenswert ist, dass die Regel dagegen bereits existierte:
> [slice-049](../done/slice-049-mechanik-sensoren.md) schrieb „die Probe muss den Fall treffen,
> der wirklich durchrutschen könnte, nicht den bequemsten" — und wandte sie auf den eigenen
> Selbsttest nicht an. Prüfsatz für jede Negativ-Fixture: *kann sie das Muster überhaupt treffen?
> Wenn nein, gehört sie ersetzt durch eine, die es **beinahe** tut.*

**Zwei beobachtbare Closure-Kriterien:**

1. `make gates` und `make verify` grün auf demselben Stand (je Exit 0) — belegt.
2. Jeder der drei Fixes ändert ein **messbares** Ergebnis gegenüber dem Vorzustand, je mit einem
   Lauf belegt: `suppression-check` meldet für eine Datei mit Direktive **und** Erwähnung genau
   **eine** statt zwei Zeilen; `regelwerk-check` wählt bei `v3.5.2` neben `v3.10.0-probe` den
   **höheren** statt des lexikografisch letzten; `verify-slice-form` zählt für eine Datei mit
   Checkboxen außerhalb des DoD **2** statt 4 Punkte.

**Nebenbeobachtung — SL-002 hat mich selbst getroffen, und zwar zweimal:** das erste `make gates`
dieses Slice war **rot** mit acht `target-missing`, weil dieses Dokument die Vorlage aus
`docs/plan/planning/` übernahm und in `in-progress/` eine Verzeichnisebene tiefer liegt. Nach der
Korrektur war der nächste Lauf erneut rot — mit zwei weiteren Verweisen derselben Art, die ich
beim Schreiben **dieser Closure-Notiz** neu eingeführt hatte. Zwei Vorfälle in einem einzigen
Slice, von derselben Instanz, die den Befund kennt und gerade darüber schreibt. Gefangen **vor**
dem Commit,
also billig — genau der Unterschied zu den in
[Etappe B](../../../reviews/2026-07-26-etappe-b-slice-048.md) und
[slice-049](../../../reviews/2026-07-26-slice-049-mechanik-sensoren.md) belegten Fällen, wo derselbe
Fehler erst nach dem Commit auffiel. Der Befund stützt den Sensor-Kandidaten aus
SL-002, und zwar in seiner allgemeineren Form: relative Verweise brechen
nicht nur beim `git mv`, sondern auch beim **Anlegen** aus einer Vorlage anderer Tiefe.

**Folge-Slices:** slice-059 (Durchsetzungs-Lücken: Guard-Gate-Liste, `make verify` ohne
Befund-Maskierung, Workflow-Skelett Schritt 9) und slice-060 (Belegtreue der Planungs-Doku:
`SL-003`, slice-048-Korrekturen, Status-Felder).

## 7. Sub-Area-Modus

Alle berührten Sub-Areas GF.
