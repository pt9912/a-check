# slice-221 — `make image-scan` prüft jede Plattform des Index, mit Nachweis der gescannten Architektur

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `git mv` (`make slice-mv`).

**Welle:** ohne Welle — die Closure-Bedingung wäre die eigene DoD.

**Bezug:** [ADR-0037](../../adr/0037-cve-scan-gegen-das-publizierte-image.md) (CVE-Scan gegen
das publizierte Image), [ADR-0043](../../adr/0043-multi-arch-ein-bau-getestet-dann-getaggt.md)
§Konsequenzen (die benannte Lücke: nur die Plattform des scannenden Rechners).

**Berührte Spec-Stellen:** — · Der Slice berührt kein Spec-Stratum.

**Verantwortlich:** Claude — gesetzt beim Übergang nach `next/` (Maintainer: „ja angehen“, 2026-10-08).

**Autor:** Claude. **Datum:** 2026-10-08.

**Lerneintrag — Form:** neuer Sensor.

---

## 1. Ziel und Abgrenzung

**Ziel:** `make image-scan` scannt jede publizierte Referenz **je Plattform** (`linux/amd64`,
`linux/arm64`) und prüft für jeden Lauf, dass Trivy wirklich die verlangte Architektur gescannt
hat — fail-closed. Damit schließt sich die in [ADR-0043](../../adr/0043-multi-arch-ein-bau-getestet-dann-getaggt.md) benannte Lücke, ohne eine neue zu öffnen.

## 1b. Messung (2026-10-08)

Trivy `0.74.0` (digest-gepinnt, wie in `tools/image-scan.sh`) gegen `ghcr.io/pt9912/a-check`:

| # | Lauf | Ergebnis |
|---|---|---|
| M1 | `--platform linux/amd64` gegen `:latest` (Index) | gescannt `architecture: amd64`, ImageID `91900f93…`; 1 Befund (`tzdata`, UNKNOWN) |
| M2 | `--platform linux/arm64` gegen `:latest` | gescannt `architecture: arm64`, ImageID `983372d2…`; derselbe Befund |
| M3 | `--platform linux/arm64` gegen `:v0.22.0` (nur amd64) | **Exit 0, gescannt `architecture: amd64`** — Trivy fällt still auf die vorhandene Plattform zurück |

**Geltungsbereich:** ein Trivy-Stand, eine Registry (GHCR), ein Tag je Fall; der erste M1-Lauf
scheiterte am Download der Vuln-DB (404 beim Mirror) und lief beim zweiten Versuch durch — ein
Lauf-Fehler, den das Skript schon heute als „GESCHEITERT" meldet.

**Folgerung aus M3:** `--platform` allein ist eine Zusage ohne Lauf (Mess-Regel 5) — der Scan
behauptete arm64 und prüfte amd64. Die gescannte Architektur ist darum aus Trivys Ergebnis zu lesen
und gegen die verlangte zu halten; Abweichung ist „GESCHEITERT", nicht grün.

**Plan-Änderung 2026-10-08 (Review F-1/F-2/F-3, vor dem Fix):** Der Entscheidungslauf bleibt beim
**Template**-Format mit den Fixtures, die
[ADR-0037](../../adr/0037-cve-scan-gegen-das-publizierte-image.md) §Fitness Function nennt — der
erste Wurf hatte ihn auf JSON umgestellt und dabei die Anker-Probe still verloren. Den Nachweis der
gescannten Architektur trägt ein **eigener** JSON-Lauf je Plattform; der Abgleich ist eine
Funktion, die `--selftest` deckt (verlangte gegen gescannte Architektur, fehlende Angabe).

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Weitere Plattformen.** *Bestand bleibt bewusst stehen*: das Image hat genau zwei
  ([ADR-0043](../../adr/0043-multi-arch-ein-bau-getestet-dann-getaggt.md)); die Liste steht als
  eine Variable im Skript.
- **Der Release-Workflow scannt.** *Ein anderer Vorgang*: der Scan bleibt der Nachtlauf gegen das
  Publizierte ([ADR-0037](../../adr/0037-cve-scan-gegen-das-publizierte-image.md)); ein Scan im Release wäre eine neue Entscheidung.
- **Behebung des `tzdata`-Befunds.** *Ein anderer Vorgang*: UNKNOWN, nicht handlungspflichtig nach
  [ADR-0037](../../adr/0037-cve-scan-gegen-das-publizierte-image.md) Punkt 3; er gehört zur nächsten Hebung des Basis-Images.

## 2. Definition of Done

- [x] `tools/image-scan.sh`: je Referenz und Plattform ein Lauf; die gescannte Architektur wird aus
      Trivys Ausgabe gelesen und gegen die verlangte gehalten (Abweichung oder fehlende Angabe ⇒
      „GESCHEITERT", Exit 2); Gegenprobe gegen ein reines amd64-Image (rot mit Meldung).
- [x] `--selftest` deckt die neue Auswertung (Architektur-Abgleich) netzlos ab.
- [x] Sensor-Datei `harness/sensors/image-scan.md`, Gate-Index-Zelle; CHANGELOG `[Unreleased]`.
- [x] Unabhängiger Review, Report unter [`docs/reviews/`](../../../reviews/README.md).
- [x] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [x] Beobachtungs-Register fortgeschrieben (oder „keine Beobachtung" notiert).
- [x] Jedes Risiko aus §6 trägt einen Ausgang.

`make gates` und `make verify` grün; `make image-scan` einmal gegen das Publizierte gelaufen.

## 3. Plan (vor Code)

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| `tools/image-scan.sh` | update | Plattform-Schleife, Architektur-Abgleich, Selbsttest |
| `harness/sensors/image-scan.md`, `harness/README.md` | update | Vertrag und Grenze |
| `CHANGELOG.md` | update | der Sensor ist öffentlich beschrieben (wie seine Einführung) |

## 4. Trigger

**Start** (`open` → `in-progress`): Maintainer-Wort („ja angehen", 2026-10-08), WIP-Limit frei.

**Rückführungen:** `in-progress` → `next` (zu groß): entfällt — ein Skript.
`in-progress` → `open` (blockiert): Trivys Ausgabe trägt die Architektur nicht verlässlich.

## 5. Closure-Trigger

DoD vollständig, `make gates` und `make verify` grün, Closure-Notiz mit Lerneintrag.

## 6. Risiken und offene Punkte

- **Doppelte Laufzeit des Nachtlaufs:** vier Plattform-Läufe statt zwei Referenz-Läufe, je mit
  Vollbericht und Entscheidungslauf. — **Ausgang:** *entfallen* — gestrichen mit Begründung: die
  Zahl der Läufe ist gestiegen (je Referenz und Plattform drei: Architektur, Vollbericht,
  Entscheidung), der befürchtete Schaden nicht — der Lauf gegen das Publizierte brauchte lokal
  wenige Minuten bei einer Vuln-DB im Cache, das Zeitlimit des Workflows sind 30 Minuten.

## 7. Closure-Notiz

**Lerneintrag — Form: neuer Sensor.** `make image-scan` scannt jede publizierte Referenz je
Plattform und weist vor dem Entscheidungslauf die **gescannte Architektur** nach
(`architektur_ok`, gedeckt von `--selftest`). Gemessen vor dem Code: Trivy fällt bei
`--platform` still auf die vorhandene Plattform zurück — ein `--platform` allein hätte einen
arm64-Scan behauptet und amd64 geprüft. Gegenprobe am reinen amd64-Image `v0.22.0`: rot mit
„Trivy hat die Architektur 'amd64' gescannt, verlangt war 'arm64'"; ohne den Abgleich grün.

**Geliefert:** `tools/image-scan.sh` (Plattform-Schleife, Architektur-Lauf, Selbsttest mit
sieben Zähl- und fünf Architektur-Fällen), Sensor-Datei, Gate-Index, CHANGELOG `[Unreleased]`.
Lauf gegen das Publizierte: vier Scans sauber, gescannt amd64 bzw. arm64.

**Was hat funktioniert:** Die Messung vor dem Code (M3) — sie hat die Falle gefunden, bevor ein
Satz über „arm64 gescannt" geschrieben war. Mess-Regel 5 hat hier vorab gegriffen.

**Was ging anders als geplant:** Der erste Wurf stellte die Entscheidung auf JSON um und verlor
dabei still die Anker-Probe des Selbsttests — der dritte Beleg für
`BEO-GATE/umbau-verliert-pruefung-still`. Die Antwort war Zurücksetzen und Ergänzen statt Flicken;
der Entscheidungslauf bleibt beim Template aus [ADR-0037](../../adr/0037-cve-scan-gegen-das-publizierte-image.md).

**Steering-Loop-Eintrag:** neuer Sensor — liegt in `Makefile:image-scan` (`tools/image-scan.sh`,
Architektur-Abgleich). Auslöser: [ADR-0043](../../adr/0043-multi-arch-ein-bau-getestet-dann-getaggt.md)
§Konsequenzen (benannte Lücke), keine Register-Schwelle.

**Lese-Schritt (wellenlos, vor dem `git mv`):** `BEO-GATE/umbau-verliert-pruefung-still` erreicht
mit diesem Slice 3× — Ausgang *geplant*: slice-222 (in `open/`) misst, ob ein Sensor auf das
Verschwinden benannter Proben möglich ist, und verkörpert die Klasse.

**Beobachtungs-Register (`../observations/`):** `BEO-GATE/umbau-verliert-pruefung-still` → 3×,
*geplant* (slice-222).

**Folge-Slices:** slice-222 (Umbau hält bestehende Prüfungen) — in `open/`.

**Risiken aus §6:** das eine Risiko trägt seinen Ausgang (*entfallen*, mit Begründung).

**Drei Paarungen:** Anker — der Zielort `Makefile:image-scan` existiert, kein `seit`-Anker verlangt
(Auslöser ist eine ADR, keine Register-Schwelle) · Folge-Slice — slice-222 existiert in `open/` ·
Register — der genannte Pfad existiert mit nicht leerem `evidence/`.

**Trigger-Audit der aktiven MR:** [`MR-016`](../../../../harness/conventions.md#mr-016) [`MR-019`](../../../../harness/conventions.md#mr-019) [`MR-025`](../../../../harness/conventions.md#mr-025) [`MR-027`](../../../../harness/conventions.md#mr-027) [`MR-028`](../../../../harness/conventions.md#mr-028) [`MR-029`](../../../../harness/conventions.md#mr-029) [`MR-030`](../../../../harness/conventions.md#mr-030) — 0 offen (geprüft 2026-10-08).

## 8. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** `GATE` (Achsen 1, 2, 3 ✓).

**Vorgelagert — offene Beobachtungen sichten** (2026-10-08):
`BEO-GATE/pipefail-bricht-pruefer-stumm-ab` (1×) — das Skript läuft mit `set -uo pipefail`;
Extraktion ohne stummen Abbruch. `BEO-GATE/umbau-verliert-pruefung-still` (2×) — beim Umbau der
Schleife darf keine bestehende Prüfung (Vollbericht, Entscheidungslauf, leere Prüfmenge, Zähler)
wegfallen; ein dritter Beleg wäre eine Lücke. `BEO-GATE/werkzeug-zugeschriebene-leistung` (2×) —
M3 ist genau dieser Fall, vor dem Code gemessen.

**Modus-Begründungsblock:** alle berührten Sub-Areas GF.
