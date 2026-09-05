# slice-087 — Vollständigkeits-Sensor für handgepflegte Indizes

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Deckt:** den Lerneintrag aus [slice-081](welle-13/slice-081-heuristik-diagnose.md), bestätigt durch
den Wiederholungsfall in [slice-085](welle-13/slice-085-schicht-ohne-aufloesung.md).
**Bezug:** dieselbe Fehlerklasse, die [slice-071](welle-12/slice-071-sensor-scope-vollstaendig.md)
für das Regelwerk-Manifest geschlossen hat; Ablösungs-Muster nach
[slice-073](welle-12/slice-073-dcheck-statt-eigenbau.md)/[slice-080](../done/slice-080-verify-abloesung-dcheck.md).

---

## 0. Trigger

**Beginn: sofort.** Der Fehler ist zweimal in zwei Tagen aufgetreten, das zweite Mal, **nachdem**
er diagnostiziert und benannt war. Ein Fehler, den Wissen allein nicht verhindert, wartet auf
nichts.

**Rückführungen:**

- `in-progress` → `open`: falls der Entscheid aus §3 auf den d-check-CR fällt **und** die
  Einreichung Maintainer-Sache bleibt — dann ist der Slice blockiert wie
  [slice-080](../done/slice-080-verify-abloesung-dcheck.md), und das ist der ehrliche Zustand.

## 1. Auslöser

**Mechanismus: `doc-check` prüft, dass jeder Link auflöst — nicht, dass jede Datei verlinkt ist.**
Was nicht verlinkt ist, kann auch nicht ins Leere zeigen. Ein handgepflegter Index sieht darum
**immer** vollständig aus, egal wie viele Einträge fehlen.

**Zwei Vorfälle, beide gemessen:**

| Datum | Fehlend | Entdeckt durch |
|---|---|---|
| 2026-08-09 | [ADR-0030](../../adr/0030-kein-digest-im-generierten-fragment.md) fehlte im ADR-Index | Zufall beim Nachtragen der nächsten ADR |
| 2026-08-09 | [ADR-0031](../../adr/0031-heuristik-grenzen-diagnose.md) fehlte im ADR-Index | Zufall, **einen Slice später** |

Der zweite Fall ist der eigentliche Auslöser: Der Fehler war bekannt, benannt und als Folge-Slice
festgehalten — und trat trotzdem beim allernächsten ADR wieder auf. `make gates` war beide Male
grün.

**Dieselbe Klasse hat dieses Repo schon einmal getroffen.**
[slice-071](welle-12/slice-071-sensor-scope-vollstaendig.md) fand sie bei `regelwerk-check`: geprüft
wurde Manifest → Baum, nicht Baum → Manifest; die Lösung war ein `comm -13` über beide Mengen. Zwei
Vorkommen derselben Asymmetrie in verschiedenen Sensoren machen sie zur **Klasse**, nicht zum
Einzelfall.

**Bestandsmessung (2026-08-09) — der Sensor bekäme heute genau ein Ziel.**

| Kandidat | Ist es ein Datei-Index? | Lücke heute |
|---|---|---|
| [`docs/plan/adr/README.md`](../../adr/README.md) | **ja** — Tabelle, eine Zeile je ADR | keine (beide Nachträge erfolgt) |
| `docs/reviews/README.md` | **nein** — Formvorgabe für Reports, listet sie bewusst nicht | — |
| `docs/plan/planning/README.md` | **nein** — Lifecycle-Prozess | — |
| `docs/plan/planning/next/README.md` | **nein** — Zustandsbeschreibung | — |

Ein Sensor für **einen** Index ist wenig — aber der Index ist der, in dem die immutablen
Entscheidungen dieses Repos stehen, und die Fehlerrate lag bei zwei von drei neuen ADRs.

## 2. Betroffene Module

Abhängig vom Entscheid in §3: entweder ein lokaler Sensor unter `tools/` — eingehängt in ein
bestehendes Meta-Gate, etwa [`gate-consistency.sh`](../../../../tools/gate-consistency.sh) — oder
ein Change Request an d-check ohne lokale Codeänderung.
[`AGENTS.md`](../../../../AGENTS.md) §5 (Gate-Liste) in beiden Fällen, falls ein Target entsteht.

## 3. Auszuführende Gates

`make gates`.

**Der Entscheid, der vor dem Bau fällt: lokaler Sensor oder d-check-CR?** Die Vorarbeit ist
gemacht, die Wahl nicht.

**Gemessen:** Kein d-check-Modul deckt die Richtung **Ziel → Verweis** ab. Geprüft wurden alle 19
(`--print-config`, `--help`): `links`/`anchors`/`codepaths` prüfen Verweis → Ziel; `tracked` den
Git-Status aufgelöster Ziele; `planning` die Roadmap-↔-`in-progress`-Konsistenz — nah dran, aber auf
Slices und Roadmap festgelegt. `--trace --require-complete` findet **Anforderungs**-Waisen
(Requirement ohne ADR/Slice), nicht Datei-Waisen; die RTM listet
[ADR-0031](../../adr/0031-heuristik-grenzen-diagnose.md) korrekt, obwohl sie im Index fehlte.
**Damit ist die Fähigkeit CR-fähig im Sinne von
[slice-073](welle-12/slice-073-dcheck-statt-eigenbau.md)** — generisch und von keinem Modul gedeckt.

| Weg | Dafür | Dagegen |
|---|---|---|
| **Lokaler Sensor** (~15 Zeilen `comm -13`, Muster [slice-071](welle-12/slice-071-sensor-scope-vollstaendig.md)) | schützt **heute**; die Fehlerrate ist belegt | ein weiterer P-Rest, den [slice-079](../done/slice-079-gate-consistency-abloesen.md)/[slice-080](../done/slice-080-verify-abloesung-dcheck.md) später wieder abtragen |
| **Nur d-check-CR** | die richtige Heimat; d-check hat die Datei-Menge bereits | Einreichung ist Maintainer-Sache; bis dahin schützt **nichts**, und der Fehler wiederholt sich messbar |
| **Beides** | Schutz jetzt, Ablösung später mit Trigger wie [slice-080](../done/slice-080-verify-abloesung-dcheck.md) | doppelte Buchführung für 15 Zeilen |

**Neigung des Autors: „beides".** Die Ablösungs-Lektion aus
[slice-073](welle-12/slice-073-dcheck-statt-eigenbau.md) galt 589 Zeilen in vier Sensoren; hier
stehen ~15 Zeilen gegen einen zweimal belegten Fehler. Ein Slice, der auf ein Fremdrepo wartet,
hätte den zweiten Vorfall nicht verhindert. Der Entscheid gehört trotzdem beim Bau begründet und
nicht hier vorweggenommen — insbesondere die Frage, **wo** der Sensor hängt: `gate-consistency`
trägt bereits Vollständigkeits-Prüfungen (`.PHONY`, Pins), ist aber selbst Ablöse-Kandidat in
[slice-079](../done/slice-079-gate-consistency-abloesen.md).

**Negativ-Proben:**

| Probe | Erwartung |
|---|---|
| ADR-Datei ohne Index-Zeile | **rot**, nennt die Datei |
| Index-Zeile ohne ADR-Datei | rot (deckt heute schon `doc-check`; darf nicht schweigen) |
| vollständiger Index (Ist-Zustand) | grün |
| Selbsttest mit Fixture beider Richtungen | feuert, sonst ist der Sensor unbelegt |

## 4. Was bewusst nicht getan wird

- **Den Index automatisch erzeugen.** Die Zeilen tragen Titel, Status und Bezugs-IDs — kuratierte
  Information, die kein Generator aus der Datei zieht. Ein erzeugter Index wäre schlechter als der
  gepflegte plus Sensor.
- **Weitere Indizes erfinden.** Die Messung in §1 hat genau einen gefunden; ein Sensor mit
  konfigurierbarer Liste ist Vorratshaltung.
- **Die vendored Baseline einbeziehen.** Der Teilbaum unter `.harness/baseline/` ist Fremdtext
  ([MR-006](../../../../harness/conventions.md#mr-006--baseline-committet-vendored-statt-per-url-referenziert)),
  nicht a-checks Doku.

## 5. DoD

- [x] Der Entscheid aus §3 ist getroffen und begründet: **beides** — lokaler Sensor als
      `gate-consistency` Check (6), CR-Text für die Ablösung in §8. Begründung in der Closure-Notiz.
- [x] Eine fehlende Index-Zeile macht `make gate-consistency` **rot**. Beleg: die Proben in der
      Closure-Notiz — die Negativ-Probe nutzt den realen Fall (die Index-Zeile von [ADR-0033](../../adr/0033-forbidden-constructs-fail-closed.md) entfernt) und war
      vor dem Bau grün.
- [x] `make gates` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft.

## 6. Closure-Notiz

**Entschieden: beides — lokaler Sensor jetzt, CR für die Ablösung.** Die Neigung aus §3 hat den Bau
überstanden, und zwar aus dem Grund, der dort schon stand: Ein Slice, der auf ein Fremdrepo wartet,
hätte den **zweiten** Vorfall nicht verhindert. Die Kosten sind gemessen klein — 37 Zeilen inklusive
Selbsttest gegen die 589 Zeilen, um die es in
[slice-073](welle-12/slice-073-dcheck-statt-eigenbau.md) ging.

**Der Ort war die eigentliche Frage, und sie hat sich beim Nachsehen aufgelöst.** §3 hielt fest,
`gate-consistency` sei selbst Ablöse-Kandidat. Die Lektüre von
[slice-079](../done/slice-079-gate-consistency-abloesen.md) zeigt: abgelöst werden **nur (1)+(2)**,
die Checks (3)(4)(5) bleiben ausdrücklich stehen. Damit ist der Sensor dort kein Anbau an ein
sterbendes Gate, sondern Check **(6)** in einem, das bleibt. Ohne diese Prüfung wäre der Sensor aus
einer falschen Annahme heraus an einen schlechteren Ort gewandert.

**Die Proben:**

```text
(a) ADR-0033-Zeile aus dem Index entfernt (der reale Fall):
    gate-consistency: FAIL — ADR '0033-forbidden-constructs-fail-closed.md' fehlt im Index …
    EXIT=1
(b) Index unveraendert:
    gate-consistency ok: … ADR-Index vollstaendig (Selbsttests gefeuert).  EXIT=0
```

Dazu ein Selbsttest **beider** Richtungen im Sensor selbst (Lücke feuert, vollständiger Index mit
Anker-Suffix schweigt) — ohne die zweite Hälfte wäre ein immer-rotes Gate von einem korrekten nicht
zu unterscheiden.

**Bewusst eng: nur die fehlende Richtung.** Ein Index-Eintrag **ohne** Datei ist ein toter Link und
gehört `doc-check`. Ihn hier mitzuprüfen wäre genau die Doppelung, die
[slice-079](../done/slice-079-gate-consistency-abloesen.md) an anderer Stelle gerade abträgt.

**Mitgenommener Doku-Drift:** [`harness/README.md`](../../../../harness/README.md) führte `cli.go`
noch als Pin-Stelle. Die Konstante ist seit
[slice-083](welle-13/slice-083-print-mk-digest-selbstbezug.md) entfernt und `cli.go` aus der
Pin-Prüfung genommen; die Zeile beschrieb ein Gate, das es so nicht mehr gibt — dieselbe Klasse
Harness-Lüge, gegen die dieses Gate selbst antritt.

**Lerneintrag — Form: neuer Sensor.** Als Prüfsatz: *Prüft ein Sensor eine Beziehung zwischen zwei
Mengen, gehört die Gegenrichtung mitgeprüft — sonst ist die vollständig aussehende Menge die
ungeprüfte.* Dritter Beleg derselben Asymmetrie:
[slice-071](welle-12/slice-071-sensor-scope-vollstaendig.md) fand sie bei `regelwerk-check`
(Manifest → Baum), hier trifft sie den ADR-Index (Baum → Index), und
[`welle-13`](../done/welle-13-results.md) hat einen dritten Fall an `verify-slice-links` benannt
(prüft nur `slice-*.md`, nicht `welle-*.md`). **Zu prüfen wäre**, ob die übrigen Sensoren dieselbe
Schlagseite haben — `suppression-check` und `commit-scope-check` sind die Kandidaten.

## 7. Sub-Area-Modus

Alle berührten Sub-Areas GF.

## 8. Change Request an d-check (zur Einreichung durch den Maintainer)

Die Ablösung des lokalen Sensors. Einreichung ist ein Akt gegenüber einem Fremdrepo und bleibt
Maintainer-Sache — dieselbe Trennung wie bei CR 1/CR 2 in
[slice-073 §8](welle-12/slice-073-dcheck-statt-eigenbau.md).

> **CR 3 — Modul `index`: Vollständigkeit handgepflegter Datei-Indizes**
>
> **Problem.** Alle 19 Module prüfen die Richtung *Verweis → Ziel*. Ein handgepflegter Index (eine
> Markdown-Datei, die je Datei eines Verzeichnisses eine Zeile trägt) sieht darum **immer**
> vollständig aus: Was nicht verlinkt ist, kann nicht ins Leere zeigen. Gemessen im Repo `a-check`:
> zwei von drei neu angelegten ADRs fehlten im Index, beide bei grünem Lauf, beide nur zufällig
> bemerkt — der zweite, **nachdem** der Fehler bereits diagnostiziert und dokumentiert war.
>
> **Nicht abgedeckt durch:** `links`/`anchors`/`codepaths` (Gegenrichtung), `tracked` (Git-Status
> aufgelöster Ziele), `planning` (auf Roadmap ↔ `in-progress` festgelegt). Auch nicht durch
> `--trace --require-complete`: das findet **Anforderungs**-Waisen (Requirement ohne ADR/Slice),
> nicht **Datei**-Waisen. Beleg: die RTM listete [ADR-0031](../../adr/0031-heuristik-grenzen-diagnose.md) korrekt auf, **während** die Datei im
> Index fehlte.
>
> **Vorschlag.** Ein Modul `index` mit einer Liste von Paaren:
>
> ```yaml
> index:
>   - { dir: "docs/plan/adr", glob: "[0-9]*.md", index: "docs/plan/adr/README.md" }
> ```
>
> Befund `index-orphan` je Datei unter `dir`, die `glob` matcht und aus `index` **nicht** verlinkt
> ist. Ein Link mit Anker-Suffix (`0031-x.md#geschichte`) zählt als Verlinkung. Die Gegenrichtung
> (Eintrag ohne Datei) deckt `links` bereits ab und bleibt ausdrücklich außen vor.
>
> **Warum bei d-check und nicht lokal:** die Fähigkeit ist generisch, und d-check hat die
> Datei-Menge bereits — es scannt die ADR-Dateien für die RTM. Der lokale Sensor in `a-check`
> (`tools/gate-consistency.sh`, Check 6) ist die Überbrückung, nicht das Ziel; er entfällt mit
> diesem Modul.

