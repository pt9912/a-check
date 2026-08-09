# Review-Report: `slice-068` — Plan-Review — 2026-08-09

**Review-Art:** unabhängiger **Plan**-Review (kein Implementierungs-Diff — der Gegenstand ist ein
Slice-Entwurf).

**Unabhängigkeit — ausdrücklich:** **unabhängiger Lauf außerhalb der Claude-Modellfamilie**,
derselbe Reviewer wie beim
[`welle-12`-Report](2026-08-09-welle-12-unabhaengig.md) desselben Tages. Der Entwurf plant die
Korrektur *seiner eigenen* Funde — die Prüfung, ob ein Plan die Funde trifft, liegt damit bei der
Instanz, die sie erhoben hat.

**Gegenstand:** `4b029e4` (Entwurf), Kontext `3b348dd` (Roadmap-Korrektur + Report).

**Skill:** `.harness/skills/reviewer.md` @ `20ee992` · <!-- d-check:ignore (Adopter-spezifischer Skill-Pfad, existiert im Ziel-Repo ggf. nicht) -->
**Modell:** OpenAI Codex, GPT-5 · **Datum:** 2026-08-09 · **Arbeitsbaum:** unverändert

**Formhinweis zur Ablage:** Wortlaut des Reviewers unverändert; ergänzt sind Links auf die
genannten Kennungen (`doc-check`) und der abgegrenzte
[Anhang](#anhang--gegenprobe-des-maintainer-agenten). Der Anhang ist **nicht** Teil des
unabhängigen Laufs.

---

## Findings

### R-068-F1 — „Quelle nicht leer" ist kein belastbarer Ersatz für eine gemessene Grundgesamtheit

- `kategorie`: **HIGH**
- `quelle`: Fund `F-12` aus dem [`welle-12`-Report](2026-08-09-welle-12-unabhaengig.md);
  Roadmap-Trigger; Regelwerk Modul 13
- `pfad`: `slice-068`-Entwurf @ `4b029e4` Zeilen 68 und 85;
  [`tools/verify-ac-form.sh:103`](../../tools/verify-ac-form.sh),
  [`tools/verify-slice-links.sh:106`](../../tools/verify-slice-links.sh)
- `befund`: Eine nichtleere Quelle kann null erkannte Einträge oder nur einen Restbestand
  enthalten; sie bleibt dann nach dem beschriebenen Entscheid grün. Zugleich sind null gefilterte
  Einheiten sowohl bei `verify-ac-form` als auch bei `verify-slice-links` legitime Zustände — der
  Plan definiert daher keine gemeinsame Bedingung, die Quellenverlust erkennt, ohne erlaubte
  Leermengen abzulehnen.
- `verifizierbar`: ja — nichtleeres Lastenheft ohne erkennbare AC-Überschriften beziehungsweise
  nichtleere Lifecycle-Verzeichnisse ohne wandernde Slices.

### R-068-F2 — Die `.PHONY`-Probe kann einen Teil-Fix als vollständig ausweisen

- `kategorie`: **HIGH**
- `quelle`: Fund `F-1`; Harness-Lügen-Schutz; Slice-Vorlage §Negativ-Probe
- `pfad`: `slice-068`-Entwurf @ `4b029e4` Zeilen 65 und 96; [`Makefile:35`](../../Makefile)
- `befund`: Acht neue Targets sind von `F-1` betroffen; im aktuellen Makefile fehlen insgesamt neun
  explizite Targets in `.PHONY`, einschließlich `arch-graph`. Die geplante Fixture benennt nur ein
  Target: Wird ausschließlich dieses deklariert, wird die Probe rot, während alle übrigen Targets
  weiterhin durch gleichnamige Dateien übersprungen werden können.
- `verifizierbar`: ja — Targetmenge gegen `.PHONY` vergleichen und jede fehlende Deklaration
  einzeln durch eine gleichnamige Datei prüfen.

### R-068-F3 — Die Suppression-Probe belegt weder den vollständigen Scope noch fehlgeschlagene Traversierung

- `kategorie`: **HIGH**
- `quelle`: Hard Rule [`AGENTS.md`](../../AGENTS.md) §3.2;
  [ADR-0005](../plan/adr/0005-lint-profil.md); Fund `F-2`
- `pfad`: `slice-068`-Entwurf @ `4b029e4` Zeile 66;
  [`tools/suppression-check.sh:33`](../../tools/suppression-check.sh),
  [`tools/suppression-check.sh:71`](../../tools/suppression-check.sh)
- `befund`: Eine einzelne `.go`-Datei außerhalb von `internal/` und `cmd/` wird bereits von einer
  zusätzlichen fest verdrahteten Wurzel erfasst; andere Repo-Verzeichnisse können weiter ungemessen
  bleiben. Zusätzlich verwirft `scan()` Fehler von `find` über `|| true`, sodass eine nicht lesbare
  oder fehlende Quelle weiterhin zu null Treffern und Exit 0 führt.
- `verifizierbar`: ja — die unveränderte `scan`-Funktion lieferte für `/definitely/not/a/source`
  trotz `find`-Fehler `rc=0, hits=0`.

### R-068-F4 — Weiteres False-Green: neue Slice-Datei kann als „grandfathered" übersprungen werden

- `kategorie`: **HIGH**
- `quelle`: Slice-Dateikonvention; Vertrag von `verify-slice-form`; Auftragsfrage P6
- `pfad`: [`docs/plan/planning/slice.template.md:3`](../plan/planning/slice.template.md),
  [`tools/verify-slice-form.sh:32`](../../tools/verify-slice-form.sh),
  [`tools/verify-slice-form.sh:114`](../../tools/verify-slice-form.sh)
- `befund`: Eine Datei `slice-068.md` trifft den Haupt-Glob `slice-*.md`, aber nicht die
  Nummernextraktion, die einen weiteren Bindestrich und Kurztitel verlangt. `applies()` liefert
  deshalb `false` und der Hauptlauf zählt die neue Datei als „älter (grandfathered)" — obwohl die
  Quelle existiert und nicht leer ist; dieser False-Green fehlt im Plan.
- `verifizierbar`: ja — die aktuellen Funktionen lieferten für `slice-068.md` `number=''` und
  `skipped`.

### R-068-F5 — Die Größenbegründung beantwortet das Ein-Lauf-/Ein-Review-Kriterium nicht

- `kategorie`: **MEDIUM**
- `quelle`: Regelwerk Modul 5; Slice-Vorlage §Größen-Regel
- `pfad`: `slice-068`-Entwurf @ `4b029e4` Zeilen 41 und 91
- `befund`: Der Plan zählt zwei Dateikategorien und drei Checkboxen, schneidet aber fünf
  unabhängige Fehlermechanismen über Makefile und sieben Skripte; DoD 1 bündelt sämtliche
  Mechanismen, während DoD 2 `F-1` nochmals separat wiederholt. Ob dieser Umfang in einem
  Agentenlauf und einer Review-Sitzung abgeschlossen werden kann, bleibt unbeantwortet.
- `verifizierbar`: nein — Plan-Schnitt und Reviewbarkeit sind ein Reviewurteil, kein bestehender
  Tool-Lauf.

## Antworten P1–P6

- **P1 — Deckung:** Die Menge „vier `verify-*`-Sensoren" entspricht exakt `F-12`. Die volle Deckung
  ist dennoch nicht gegeben: `R-068-F1` bis `R-068-F4` zeigen ungemessene Teilmengen beziehungsweise
  einen zusätzlichen False-Green.
- **P2 — `F-12`-Entscheid:** **Nicht tragfähig.** „Quelle gelesen und nicht leer" verwechselt
  Rohbestand mit erkanntem Prüfbestand und berücksichtigt legitime Leermengen nicht
  sensorspezifisch.
- **P3 — Negativ-Proben:** Range-Fehler und zusätzliche reguläre Baseline-Datei treffen ihre
  Befunde. `.PHONY`, Suppression-Scope und die zusammengefasste `verify-*`-Probe können
  unvollständige Korrekturen als vollständig ausweisen.
- **P4 — Schnitt:** Geprüft, ohne separaten Befund. `F-2` stellt den bereits in
  [`AGENTS.md`](../../AGENTS.md) und [ADR-0005](../plan/adr/0005-lint-profil.md) behaupteten
  Repo-Scope her; es ist keine neue Produktsemantik und kein spec-first-Fall. Die Gruppierung nach
  „grün ohne vollständige Messung" ist sachlich vertretbar.
- **P5 — Größe:** Zwei Schichten und drei Checkboxen sind formal eingehalten; die materielle
  Ein-Lauf-/Ein-Review-Grenze ist durch `R-068-F5` nicht belegt.
- **P6 — Vollständigkeit:** Zwei weitere Fälle derselben Klasse wurden gefunden: verschluckte
  Traversierungsfehler im Suppression-Sensor und das falsche Grandfathering nicht parsebarer neuer
  Slice-Dateinamen.

## Negativbefunde

- Geprüft, ohne Befund: `F-12` betrifft im ursprünglichen Report tatsächlich genau
  `verify-closure-notes`, `verify-slice-form`, `verify-slice-links` und `verify-ac-form`.
- Geprüft, ohne Befund: Der ungültige Commit-Range ist eine wirksame Gegenprobe für `F-5`.
- Geprüft, ohne Befund: Eine zusätzliche reguläre Datei außerhalb von `SHA256SUMS` ist eine
  wirksame Gegenprobe für `F-14`.
- Geprüft, ohne Befund: Keiner der fünf geplanten Fixes verlangt eine Änderung der Produkt-Spec
  oder einer `Accepted`-ADR.
- Geprüft, ohne weiteren Befund: Änderungen an Command-Guard, CI-Verkabelung und `.d-check.yml` im
  Migrationsdiff ergaben keine zusätzliche gleichartige Auslassung.

## Kategorie-Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 4 |
| MEDIUM | 1 |
| LOW | 0 |
| INFO | 0 |

## Verdikt

**Plan blockiert.** Der zentrale `F-12`-Entscheid lässt weiterhin ungemessene Bestände grün, zwei
Gegenproben akzeptieren Teil-Fixes, und mindestens zwei weitere False-Greens derselben Klasse
fehlen im Schnitt.

Ausgeführt wurden `make verify`, die vier betroffenen Einzelchecks sowie gezielte Funktionsproben.
Der bekannte ungültige Range meldete weiterhin Exit 0. `make gates` wurde für diesen reinen
Plan-Review ohne Implementierungsdiff nicht ausgeführt.

---

## Anhang — Gegenprobe des Maintainer-Agenten

**Nicht Teil des unabhängigen Laufs.** Vier der fünf Findings sind mechanisch nachprüfbar; alle
vier bestätigt:

| Finding | Gegenprobe |
|---|---|
| `R-068-F4` | bestätigt. `slice_num()` verlangt einen Bindestrich nach der Nummer: `slice-068.md` → `number=''` → grandfathered; `slice-068-sensor-false-greens.md` → `number='068'` → geprüft |
| `R-068-F3` | bestätigt. Isoliert: `scan /definitely/not/a/source` → `bfs: error: … nicht gefunden` auf stderr, aber `rc=0`, `hits=''`. `\|\| true` setzt den Exit |
| `R-068-F2` | bestätigt. Neun Targets fehlen in `.PHONY`; die geplante Probe nannte eines und hätte einen Teil-Fix als vollständig ausgewiesen |
| `R-068-F1` | bestätigt. Der Entscheid trennt Rohbestand nicht vom *erkannten* Prüfbestand, und eine gemeinsame Bedingung kann es nicht geben — Leermengen sind je Sensor unterschiedlich legitim |

**Folge:** Der Entwurf `4b029e4` wird nicht überarbeitet, sondern **zerlegt** — die Größen-Regel
verlangt bei Nichtpassung ausdrücklich Zerlegen statt Dehnen, und aus fünf Funden sind sieben
geworden. Neuer Schnitt nach **Fehlermechanismus**, ein Mechanismus je Slice:
[slice-068](../plan/planning/done/slice-068-phony-vollstaendig.md) (Target läuft gar nicht) ·
[slice-069](../plan/planning/in-progress/slice-069-sensor-fehler-propagierung.md) (Sensor verschluckt
einen Fehler) ·
[slice-070](../plan/planning/open/slice-070-grundgesamtheit-messen.md) („nichts gefunden" ≠
„nichts da") ·
[slice-071](../plan/planning/open/slice-071-sensor-scope-vollstaendig.md) (Sensor misst nur einen
Teil des behaupteten Bestands).
