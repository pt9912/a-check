# slice-217 — Multi-Arch-Image: Build für `linux/amd64` und `linux/arm64`

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `git mv` (`make slice-mv`).

**Welle:** welle-18 — [Welle-Plan](welle-18-multi-arch-image.md).

**Bezug:** [AC-FA-DIST-001](../../../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk),
[AC-QA-03](../../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit); die ADR aus slice-216.

**Berührte Spec-Stellen:** `spezifikation.md`
§[SPEC-DIST-001](../../../../spec/spezifikation.md#spec-dist-001--laufzeitform-und-distribution)
(in slice-216 geschrieben, hier umgesetzt).

**Verantwortlich:** Claude — gesetzt beim Übergang nach `next/` (Maintainer: „ok“, 2026-10-07).

**Autor:** Claude. **Datum:** 2026-10-07.

**Lerneintrag — Form:** neuer Sensor.

---

## 1. Ziel und Abgrenzung

**Ziel:** Das Dockerfile baut das Runtime-Image für beide Plattformen per Cross-Compile, und ein
Make-Target erzeugt daraus lokal einen Image-Index; `make ci` bleibt grün, und das arm64-Bild ist
lokal als gebaut und als arm64-Binary belegt.

**Plan-Änderung 2026-10-07 (Review F-4/F-6, vor dem Fix):** Der Test gegen eine Bild-Referenz
läuft über `make image-test IMAGE_REF=…` — so nennt ihn die Fitness Function von
[ADR-0043](../../adr/0043-multi-arch-ein-bau-getestet-dann-getaggt.md); mit gesetztem `IMAGE_REF`
entfällt der Bau. Das zunächst gelieferte eigene Target `image-test-ref` entfällt wieder.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Release-Pipeline, Push, Spiegel, Test auf arm64-Runnern.** *Ein Folge-Slice übernimmt es*:
  slice-218.
- **Die Gate-Stages (lint, test, coverage) für arm64.** *Bestand bleibt bewusst stehen*: sie prüfen
  Go-Quelltext, nicht das Plattform-Bild; Out-of-Scope der Welle.

## 2. Definition of Done

- [x] Dockerfile: Build-Stufe auf der Plattform des Runners, Ziel-Plattform aus den
      Build-Argumenten; Laufzeit-Stufe ohne `RUN`.
- [x] Make-Target für den Multi-Arch-Bau nach der ADR aus slice-216, im Gate-Index
      (`harness/README.md` §Sensors oder §Nicht-Gates) eingetragen.
- [x] Image-Test gegen eine übergebene Bild-Referenz (Tag oder Digest) statt fest
      `$(IMAGE):dev`, Plattform des Hosts geprüft; lokaler Beleg: beide Plattform-Bilder gebaut, das
      arm64-Binary ist ein arm64-ELF, das amd64-Bild besteht den Image-Test — mit Gegenprobe.
- [x] Unabhängiger Review, Report unter [`docs/reviews/`](../../../reviews/README.md).
- [x] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [x] Beobachtungs-Register fortgeschrieben (oder „keine Beobachtung" notiert).
- [x] Jedes Risiko aus §6 trägt einen Ausgang.

`make gates` grün.

## 3. Plan (vor Code)

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| `Dockerfile` | update | Cross-Compile |
| `Makefile`, `harness/README.md` | update | Multi-Arch-Target, Gate-Index |
| `tools/image-test.sh` | update | Bild-Referenz von außen (ADR-Folgepflicht) |
| `tools/image-multiarch.sh`, `tools/multiarch-check.sh` | neu | Bau des Index und Prüfung des Archivs |
| `.claude/hooks/pretooluse-command-guard.sh`, `.gitignore` | update | neues Prüf-Target in der Guard-Liste; Ausgabe-Verzeichnis `.build/` |

## 4. Trigger

**Start** (`next` → `in-progress`): slice-216 in `done/`, Vertrag abgenommen.

**Rückführungen — vorab benannt:**

- `in-progress` → `next` (zu groß): das Make-Target braucht eine neue Werkzeugkette im Repo
  (Builder-Instanz, Emulation) — dann teilen.
  **Bewertung beim Implementieren (Review F-3):** eingetreten ist die Builder-Instanz, nicht die
  Emulation; sie ist ein Aufruf in einem Skript und hielt den Slice in einer Review-Sitzung prüfbar.
  Nicht geteilt — die Bedingung war weiter formuliert als die Größenregel, an der sie hängt.
- `in-progress` → `open` (blockiert): der lokale Docker kann keinen Multi-Arch-Index erzeugen.

## 5. Closure-Trigger

DoD vollständig, `make gates` Exit 0, Closure-Notiz mit Lerneintrag.

## 6. Risiken und offene Punkte

- **Der lokale Build ist nicht der Release-Build:** was hier grün ist, belegt das Dockerfile, nicht
  das veröffentlichte Bild. — **Ausgang:** *weiter offen* — der Release-Bau ruft dasselbe Target
  (slice-218), aber erst der Tag zeigt das veröffentlichte Bild; das Beobachtungs-Register zählt die
  Klasse unter `BEO-GATE/ungelaufene-mechanik-docker-hub-spiegel`.

## 7. Closure-Notiz

**Lerneintrag — Form: neuer Sensor.** `make image-multiarch` baut das Release-Bild für
`linux/amd64` und `linux/arm64` als einen Image-Index und prüft das Archiv mit
`tools/multiarch-check.sh`: genau zwei Einträge (zweimal verschieden gezählt), Plattform aus der
Config, Index-Beschriftung gleich der Config, ELF-Maschinentyp des Binarys, Versions-Label. Jede
dieser Eigenschaften hat eine Gegenprobe, die rot aus dem genannten Grund war: nur amd64,
Attestierung, dritter Eintrag mit Array-Feld, Index nur mit amd64, vertauschte und fremde
Beschriftung, Bau ohne Cross-Compile (amd64-Binary im arm64-Bild), falsche Version. Grenze: kein
Lauf auf arm64 — den trägt der Image-Test auf einem arm64-Rechner (slice-218).

**Geliefert:** Dockerfile per Cross-Compile ([ADR-0043](../../adr/0043-multi-arch-ein-bau-getestet-dann-getaggt.md) Punkt 2), `make image-multiarch`,
`make image-test IMAGE_REF=…` (Bild-Referenz ohne Bau, Host-Plattform per ELF geprüft), Gate-Index,
Guard-Liste; `make ci` unverändert grün.

**Was hat funktioniert:** Die Gegenproben als Archive, die genau eine Eigenschaft brechen. Die
teuerste — ein Bau ohne `GOOS`/`GOARCH` — ist der Fehler, vor dem der Cross-Compile schützt, und
sie wurde rot an der Stelle, die ihn benennt.

**Was ging anders als geplant:** Der Prüfer versprach in zwei Fassungen mehr, als er prüfte: erst
ließ ein Array-Feld eine dritte Plattform durch, dann fiel beim Umbau die Index-Beschriftung weg
(Review F-2, Delta D-1). Der Gate-Index beschrieb eine Pipeline-Verwendung, die es noch nicht gibt
(F-1). Und das eigene Target `image-test-ref` wich von der Fitness Function der ADR ab, bis es
entfiel (F-4). Die BuildKit-Pinnung trägt Tag und Digest nebeneinander; dass beide zusammengehören,
prüft kein Sensor — die bekannte Grenze von `make version-coherence` (Wahrheit, nicht Divergenz,
Review F-10). Der Builder-Container bleibt nach dem Lauf stehen. Benannte Grenze (Delta-Review 2, D2-1): fehlt dem letzten Eintrag die `platform`, liest das
Skript eine Beschriftung, die außerhalb der Manifest-Liste stünde — ein Feld, das das OCI-Schema
dort nicht kennt und BuildKit nicht schreibt.

**Steering-Loop-Eintrag:** neuer Sensor — liegt in `Makefile:image-multiarch` (`tools/multiarch-check.sh`).
Auslöser: [ADR-0043](../../adr/0043-multi-arch-ein-bau-getestet-dann-getaggt.md), keine Register-Schwelle.

**Beobachtungs-Register (`../observations/`):** `BEO-GATE/zusage-weiter-als-ihre-durchsetzung` →
4× (ein Vorgang, drei Funde; Ausgang beim Lese-Schritt der welle-18-Closure);
`BEO-GATE/pipefail-bricht-pruefer-stumm-ab` neu (1×).

**Folge-Slices:** slice-218 (in `open/`) — dort auch das Versions-Label im Image-Test (Review F-11).

**Risiken aus §6:** das eine Risiko trägt seinen Ausgang (*weiter offen*, Beobachtungs-Register).

**Drei Paarungen:** getragen von der Closure von welle-18.

**Trigger-Audit der aktiven MR:** [`MR-016`](../../../../harness/conventions.md#mr-016) [`MR-019`](../../../../harness/conventions.md#mr-019) [`MR-025`](../../../../harness/conventions.md#mr-025) [`MR-027`](../../../../harness/conventions.md#mr-027) [`MR-028`](../../../../harness/conventions.md#mr-028) [`MR-029`](../../../../harness/conventions.md#mr-029) [`MR-030`](../../../../harness/conventions.md#mr-030) — 0 offen (geprüft 2026-10-07).

## 8. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** `GATE` (1, 2, 3 ✓).

**Vorgelagert — offene Beobachtungen sichten** (2026-10-07):
`BEO-GATE/preflight-deckt-den-ci-schritt-nicht` (1×) — ein neues Target muss im Pre-Flight und in
der CI dieselbe Menge sein; `BEO-GATE/werkzeug-zugeschriebene-leistung` (2×) — „Multi-Arch" nicht
dem Werkzeug zuschreiben, sondern am Index messen. Beim Übergang nach `next/` erneut sichten.

**Modus-Begründungsblock:** alle berührten Sub-Areas GF.
