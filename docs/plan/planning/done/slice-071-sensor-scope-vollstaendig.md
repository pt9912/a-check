# slice-071 — Sensoren messen den ganzen Bestand, den sie behaupten

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Deckt:** `F-2` und `F-14` aus dem
[Review-Report `welle-12`](../../../reviews/2026-08-09-welle-12-unabhaengig.md), verschärft durch
den ersten Teil von `R-068-F3` aus dem
[Plan-Review](../../../reviews/2026-08-09-slice-068-plan-review.md);
[ADR-0005](../../adr/0005-lint-profil.md),
[MR-006](../../../../harness/conventions.md#mr-006--baseline-committet-vendored-statt-per-url-referenziert).
**Bezug:** Neuschnitt des zurückgezogenen Sammel-Entwurfs `4b029e4` nach Fehlermechanismus;
Geschwister [slice-068](../done/slice-068-phony-vollstaendig.md),
[slice-069](../done/slice-069-sensor-fehler-propagierung.md),
[slice-070](../done/slice-070-grundgesamtheit-messen.md). Roadmap-Zeile *Aktuelle Welle* in der
[Roadmap](../in-progress/roadmap.md).

---

## 1. Auslöser

**Mechanismus: der Sensor misst einen Ausschnitt und meldet über das Ganze.** Er läuft korrekt,
sein Ergebnis stimmt für das Gemessene — nur ist das Gemessene kleiner als die Behauptung.

**`F-2` — [`tools/suppression-check.sh:71`](../../../../tools/suppression-check.sh)** scannt
`scan ./internal; scan ./cmd`, während Target-Beschreibung und Hard Rule
[`AGENTS.md`](../../../../AGENTS.md) §3.2 von den „Go-Quellen des Repos" sprechen. Eine `.go`-Datei
im Wurzelverzeichnis mit wirksamem `//nolint` läuft durch.

Heute ist die reale Exposition **null** — außerhalb von `internal/` und `cmd/` existiert keine
`.go`-Datei. Die Lücke ist latent, nicht akut; sie zu schließen ist trotzdem billiger als sie zu
bewachen. `R-068-F3` schärft nach: eine zusätzliche fest verdrahtete Wurzel behebt das nicht,
solange die Wurzelmenge **aufgezählt** statt **abgeleitet** ist.

**`F-14` — [`tools/regelwerk-check.sh:56`](../../../../tools/regelwerk-check.sh)** prüft die
Baseline-Integrität mit `sha256sum -c`, das ausschließlich **gelistete** Einträge bestätigt. Eine
Datei im vendored Baum, die nicht in `SHA256SUMS` steht, bleibt ungemessen — der Lauf meldet
weiterhin „Integritaet ok — 42 Dateien". Die Prüfrichtung ist einseitig: Manifest → Baum, nicht
Baum → Manifest.

## 2. Betroffene Module

Eine Schicht: **`tools/`** — [`suppression-check.sh`](../../../../tools/suppression-check.sh) und
[`regelwerk-check.sh`](../../../../tools/regelwerk-check.sh).

Keine Spec-, ADR- oder Produktcode-Änderung: beide Sensoren sollen den Bestand messen, den die
vorhandene Doku bereits als Gegenstand nennt.

## 3. Auszuführende Gates

`make gates` (enthält `suppression-check`), `make regelwerk-check` (läuft in keinem Aggregat —
gezielt aufrufen).

**Negativ-Proben:**

| Probe | Erwartung |
|---|---|
| `.go`-Datei mit `//nolint` **außerhalb** `internal/`/`cmd/` — im Wurzelverzeichnis **und** in einem dritten Verzeichnis | beide rot; eine Probe an nur einer Stelle würde eine bloß erweiterte Aufzählung als Ableitung ausweisen |
| Datei im Baseline-Baum, die in `SHA256SUMS` fehlt | rot, nennt die Datei |
| unveränderter Bestand | grün — die Erweiterung darf den Normalfall nicht rot machen |

Die zwei Orte in der ersten Zeile sind die Lehre aus `R-068-F2`: eine Ein-Ort-Probe kann einen
Teil-Fix als vollständig ausweisen.

## 4. Was bewusst nicht getan wird

- **Ein Go-Parser für `suppression-check`.** Die Zeichenketten-Grenze (`s := "//nolint"` wird
  gemeldet) bleibt als ehrliche Grenze deklariert — sie ist fail-closed und selten.
- **Die Freshness-Hälfte von `regelwerk-check`.** Ob die vendored Baseline dem Upstream-Stand
  entspricht, ist eine Netzoperation und ausdrücklich kein Gate. Dieser Slice fasst nur die
  Integritäts-Hälfte an.
- **Die übrigen Mechanismen** — [slice-068](../done/slice-068-phony-vollstaendig.md),
  [slice-069](../done/slice-069-sensor-fehler-propagierung.md),
  [slice-070](../done/slice-070-grundgesamtheit-messen.md).
- **Reihenfolge-Hinweis:** [slice-069](../done/slice-069-sensor-fehler-propagierung.md) berührt dieselbe
  Datei `suppression-check.sh` (Fehler-Propagierung). Beide Slices sind **nacheinander** zu
  fahren, nicht parallel.

## 5. DoD

- [x] `suppression-check` misst die Go-Quellen des Repos, und die Wurzelmenge ist **abgeleitet,
      nicht aufgezählt**. Beleg: `.go` mit `//nolint` im Wurzelverzeichnis **und** in einem
      dritten Verzeichnis → Exit 2, beide gemeldet (`./outside.go:3`, `./dritter/tief.go:3`);
      nach Entfernen Exit 0.
- [x] `regelwerk-check` meldet eine Datei im Baseline-Baum, die in `SHA256SUMS` fehlt. Beleg:
      `UNLISTED.md` → Exit 2 mit Nennung der Datei; nach Entfernen Exit 0.
- [x] `make gates` und `make regelwerk-check` grün — **Ausgabe in eine Datei**, Exit-Code getrennt
      geprüft, nie in eine Pipe.

## 6. Closure-Notiz

**Geliefert:** [`suppression-check.sh`](../../../../tools/suppression-check.sh) leitet die
Wurzelmenge aus dem Repo-Baum ab (`SCAN_EXCLUDES` deklariert `.git` und die vendored Baseline);
[`regelwerk-check.sh`](../../../../tools/regelwerk-check.sh) prüft die Baseline-Integrität in
**beide** Richtungen — Manifest → Baum wie bisher, und neu Baum → Manifest.

**Lerneintrag — Form: geschärfte Regel.** Als Prüfsatz: *Wer eine Menge prüft, muss auch prüfen,
dass die Menge vollständig ist — sonst misst der Sensor korrekt und sagt trotzdem die Unwahrheit.*

**Die Ursache** war in beiden Fällen nicht Nachlässigkeit, sondern eine Grenze, die zum Zeitpunkt
ihrer Entstehung **stimmte**: `./internal ./cmd` war die vollständige Liste der Go-Bäume, als sie
geschrieben wurde, und `sha256sum -c` prüft prinzipbedingt nur, was im Manifest steht. Beide
Grenzen verfallen still — die erste, sobald ein Verzeichnis dazukommt, die zweite, sobald eine
Datei dazukommt. Kein Lauf hätte das gemeldet, weil beide Sensoren *innerhalb ihrer Menge* korrekt
arbeiten. Genau deshalb ist eine aufgezählte Wurzelmenge eine Lücke mit Verfallsdatum, und eine
einseitige Prüfrichtung ebenso.

**Zwei beobachtbare Closure-Kriterien:**

1. Eine `.go`-Datei mit `//nolint` an **zwei** Orten außerhalb `internal/`/`cmd/` macht
   `make suppression-check` rot und nennt beide — eine Ein-Ort-Probe hätte eine bloß erweiterte
   Aufzählung als Ableitung ausgewiesen (`R-068-F2`).
2. Eine nicht manifestierte Datei im Baseline-Baum macht `make regelwerk-check` rot und nennt sie;
   der unveränderte Bestand bleibt grün (42 Dateien).

**Folge-Slices:** keine aus diesem Slice. Offen bleiben
[slice-072](../open/slice-072-scope-sensor-praeventiv.md),
[slice-073](../done/slice-073-dcheck-statt-eigenbau.md) und
[slice-074](../done/slice-074-doc-targets-wirksam.md); **Gruppe A des Reviews ist mit diesem Slice
vollständig** (`F-1`, `F-2`, `F-5`, `F-12`, `F-14` plus `R-068-F3`/`R-068-F4`).

## 7. Sub-Area-Modus

Alle berührten Sub-Areas GF.
