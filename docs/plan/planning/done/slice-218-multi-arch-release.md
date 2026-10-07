# slice-218 — Multi-Arch-Image: Release-Pipeline

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `git mv` (`make slice-mv`).

**Welle:** welle-18 — [Welle-Plan](../welle-18-multi-arch-image.md).

**Bezug:** [AC-FA-DIST-001](../../../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk),
[AC-FA-DIST-002](../../../../spec/lastenheft.md#ac-fa-dist-002),
[AC-QA-03](../../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit); die ADR aus slice-216.

**Berührte Spec-Stellen:** `spezifikation.md`
§[SPEC-DIST-001](../../../../spec/spezifikation.md#spec-dist-001--laufzeitform-und-distribution)
(in slice-216 geschrieben, hier umgesetzt).

**Verantwortlich:** Claude — gesetzt beim Übergang nach `next/` (Maintainer: „ok“, 2026-10-07).

**Autor:** Claude. **Datum:** 2026-10-07.

**Lerneintrag — Form:** neuer Sensor.

---

## 1. Ziel und Abgrenzung

**Ziel:** Die Release-Pipeline baut das Image einmal für beide Plattformen, testet genau dieses
Bild auf `linux/amd64` und `linux/arm64`, veröffentlicht den Index-Digest, prüft das
Versions-Label je Plattform und spiegelt nach Docker Hub mit Gleichheits-Prüfung nach dem
abgenommenen Vertrag; `releasing.md` beschreibt den Ablauf.

**Plan-Änderung 2026-10-07 (vor dem Code):** Der eine Bau schreibt **zwei** Ausgaben — das
OCI-Archiv, das `multiarch-check` prüft, und den Upload ohne Tag (`push-by-digest`); das
Make-Target verlangt, dass beide denselben Index-Digest tragen (lokal gemessen: gleich, und der
Upload legt kein Tag an). Der Image-Test prüft dazu das Versions-Label und legt die Scan-Ausgabe
für den Plattform-Vergleich ab. Dafür berührt der Slice `tools/image-multiarch.sh`,
`tools/image-test.sh` und das `Makefile` (Review slice-217 F-11).

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Das Release selbst und die Gegenprobe am veröffentlichten Image.** *Ein anderer Vorgang*: sie
  sind der Closure-Trigger von welle-18.
- **Die CI für Pull-Requests auf arm64.** *Bestand bleibt bewusst stehen*: der PR-Lauf prüft
  Quelltext und das amd64-Bild; das arm64-Bild entsteht und wird geprüft, wo es veröffentlicht
  wird.

## 2. Definition of Done

- [x] `release.yml`: ein Bau für beide Plattformen, Test des gebauten Bilds auf beiden (arm64
      nach ADR) samt Vergleich der Scan-Ausgabe beider Plattformen, Bild-Tag erst danach,
      Versions-Label je Plattform (im Image-Test, der es heute nicht prüft), Digest-Pin in Summary
      und GitHub-Release.
- [x] Spiegel-Schritt prüft die Gleichheit nach
      [AC-FA-DIST-002](../../../../spec/lastenheft.md#ac-fa-dist-002) in der abgenommenen Fassung,
      fail-closed.
- [x] `docs/user/releasing.md` (Ablauf, arm64-Runner unter Vorbedingungen), Benutzerhandbuch
      (Plattformen, Pin), Hub-Seite `packaging/dockerhub/` (der GHCR-Digest löst dort auf;
      Index- statt Config-Digest), CHANGELOG `[Unreleased]`.
- [x] Unabhängiger Review, Report unter [`docs/reviews/`](../../../reviews/README.md).
- [x] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [x] Beobachtungs-Register fortgeschrieben (oder „keine Beobachtung" notiert).
- [x] Jedes Risiko aus §6 trägt einen Ausgang.

`make gates` grün (`doc-workflows` prüft die `uses:`-Form).

## 3. Plan (vor Code)

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| `.github/workflows/release.yml` | update | Multi-Arch-Release |
| `tools/image-multiarch.sh`, `tools/image-test.sh`, `Makefile` | update | Upload ohne Tag mit Digest-Gleichheit; Versions-Label und Scan-Ausgabe im Image-Test |
| `docs/user/releasing.md`, `docs/user/benutzerhandbuch.md`, `CHANGELOG.md` | update | öffentlicher Vertrag |
| `packaging/dockerhub/overview.md`, `packaging/dockerhub/README.md` | update | Hub-Seite sagt Index-Digest und Auflösbarkeit des GHCR-Digests |

## 4. Trigger

**Start** (`next` → `in-progress`): slice-217 in `done/`.

**Rückführungen — vorab benannt:**

- `in-progress` → `next` (zu groß): Test-Job und Spiegel brauchen getrennte Mechaniken, die sich
  nicht in einer Review-Sitzung prüfen lassen — dann teilen.
- `in-progress` → `open` (blockiert): eine Pipeline-Mechanik lässt sich vor dem Tag nicht trocken
  laufen und hat keinen sicheren Probe-Weg (etwa ein Prerelease-Tag).

## 5. Closure-Trigger

DoD vollständig, `make gates` Exit 0, Closure-Notiz mit Lerneintrag.

## 6. Risiken und offene Punkte

- **Erst am Tag geprüft:** Push, Spiegel und arm64-Test laufen nur in der Release-Pipeline; ein
  Fehler zeigt sich am ersten echten Tag (vgl. `BEO-GATE/ungelaufene-mechanik-docker-hub-spiegel`,
  1×). — **Ausgang:** *weiter offen* — der erste Release-Tag der Welle prüft es; das
  Beobachtungs-Register zählt die Klasse unter `BEO-GATE/ungelaufene-mechanik-docker-hub-spiegel`.
- **Pre-Flight deckt die Pipeline nicht:** `make preflight` läuft die Release-Schritte nicht (vgl.
  `BEO-GATE/preflight-deckt-den-ci-schritt-nicht`, 1×). — **Ausgang:** *weiter offen* — lokal
  geprobt wurden die Schritte einzeln (Upload ohne Tag, Riegel, Tag, Spiegel gegen zwei
  `registry:2`), nicht der Workflow; das Beobachtungs-Register führt die Klasse unter
  `BEO-GATE/preflight-deckt-den-ci-schritt-nicht`.

## 7. Closure-Notiz

**Lerneintrag — Form: neuer Sensor.** Die Release-Pipeline trägt drei neue Riegel, die vor dem
Bild-Tag greifen: der hochgeladene Index-Digest muss gleich dem geprüften Archiv sein
(`make image-multiarch PUSH_NAME=…`), die Scan-Ausgabe beider Plattformen muss je Strom gleich
sein (Hash-Vergleich in `publish`), und ein Versions-Tag wird nie auf einen anderen Digest
umgehängt. Dazu prüft `make image-test` das Versions-Label und verlangt mit `IMAGE_REF` ein
ausdrückliches `VERSION`. Lokal je mit Gegenprobe; der Lauf am Tag steht aus.

**Geliefert:** `release.yml` (build → test-amd64/test-arm64 → publish → hub-description),
`tools/image-multiarch.sh` (Upload ohne Tag), `tools/image-test.sh` (Label, Scan-Ausgabe),
`tools/multiarch-check.sh` (fünf OCI-Labels wieder geprüft), `releasing.md`, Handbuch 1.46,
Hub-Seite, CHANGELOG `[Unreleased]`.

**Was hat funktioniert:** Die Pipeline-Schritte einzeln gegen zwei lokale Registries zu proben —
Upload ohne Tag, Tag per `imagetools`, Kopie in eine zweite Registry — hat die Annahmen
(Digest bleibt, kein Tag entsteht) vor dem Tag belegt statt sie zu behaupten.

**Was ging anders als geplant:** Der Umbau verlor die Prüfung von fünf OCI-Labels, ohne dass es
jemand sagte (Review F-4) — dieselbe Klasse wie D-1 in slice-217, jetzt im Register. Die
Umnummerierung in `releasing.md` hätte einen Verweis der unveränderlichen [ADR-0030](../../adr/0030-kein-digest-im-generierten-fragment.md) verschoben
(F-7); der Re-Pin bleibt Schritt 6. Und ein Make-Riegel in der Voraussetzungsliste traf jedes
Target, nicht nur `image-test` (D-2).

**Steering-Loop-Eintrag:** neuer Sensor — liegt in `.github/workflows/release.yml` (Jobs
`test-amd64`, `test-arm64`, `publish`). Auslöser:
[ADR-0043](../../adr/0043-multi-arch-ein-bau-getestet-dann-getaggt.md), keine Register-Schwelle.

**Beobachtungs-Register (`../observations/`):** `BEO-GATE/zusage-weiter-als-ihre-durchsetzung` →
5× (Ausgang beim Lese-Schritt der welle-18-Closure); `BEO-GATE/umbau-verliert-pruefung-still`
neu (2×: slice-217, slice-218).

**Folge-Slices:** keine.

**Risiken aus §6:** beide tragen ihren Ausgang (*weiter offen*, Beobachtungs-Register).

**Drei Paarungen:** getragen von der Closure von welle-18.

**Trigger-Audit der aktiven MR:** [`MR-016`](../../../../harness/conventions.md#mr-016) [`MR-019`](../../../../harness/conventions.md#mr-019) [`MR-025`](../../../../harness/conventions.md#mr-025) [`MR-027`](../../../../harness/conventions.md#mr-027) [`MR-028`](../../../../harness/conventions.md#mr-028) [`MR-029`](../../../../harness/conventions.md#mr-029) [`MR-030`](../../../../harness/conventions.md#mr-030) — 0 offen (geprüft 2026-10-07).

## 8. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** `GATE` (1, 2, 3 ✓) · `USER` (2, 3 ✓).

**Vorgelagert — offene Beobachtungen sichten** (2026-10-07): die zwei Einträge in §6;
`BEO-GATE/hebungskanal-haengt-an-repo-externen-schaltern` (1×) — der arm64-Runner ist ein
repo-externer Schalter; `releasing.md` §Vorbedingungen nennt ihn. Beim Übergang nach `next/`
erneut sichten.

**Modus-Begründungsblock:** alle berührten Sub-Areas GF.
