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

**Verantwortlich:** —

**Autor:** Claude. **Datum:** 2026-10-07.

**Lerneintrag — Form:** wird bei Closure benannt.

---

## 1. Ziel und Abgrenzung

**Ziel:** Die Release-Pipeline baut das Image einmal für beide Plattformen, testet genau dieses
Bild auf `linux/amd64` und `linux/arm64`, veröffentlicht den Index-Digest, prüft das
Versions-Label je Plattform und spiegelt nach Docker Hub mit Gleichheits-Prüfung nach dem
abgenommenen Vertrag; `releasing.md` beschreibt den Ablauf.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Das Release selbst und die Gegenprobe am veröffentlichten Image.** *Ein anderer Vorgang*: sie
  sind der Closure-Trigger von welle-18.
- **Die CI für Pull-Requests auf arm64.** *Bestand bleibt bewusst stehen*: der PR-Lauf prüft
  Quelltext und das amd64-Bild; das arm64-Bild entsteht und wird geprüft, wo es veröffentlicht
  wird.

## 2. Definition of Done

- [ ] `release.yml`: ein Bau für beide Plattformen, Test des gebauten Bilds auf beiden (arm64
      nach ADR), Push des Index, Versions-Label je Plattform, Digest-Pin in Summary und
      GitHub-Release.
- [ ] Spiegel-Schritt prüft die Gleichheit nach
      [AC-FA-DIST-002](../../../../spec/lastenheft.md#ac-fa-dist-002) in der abgenommenen Fassung,
      fail-closed.
- [ ] `docs/user/releasing.md` und Benutzerhandbuch (Plattformen, Pin), CHANGELOG `[Unreleased]`.
- [ ] Unabhängiger Review, Report unter [`docs/reviews/`](../../../reviews/README.md).
- [ ] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [ ] Beobachtungs-Register fortgeschrieben (oder „keine Beobachtung" notiert).
- [ ] Jedes Risiko aus §6 trägt einen Ausgang.

`make gates` grün (`doc-workflows` prüft die `uses:`-Form).

## 3. Plan (vor Code)

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| `.github/workflows/release.yml` | update | Multi-Arch-Release |
| `docs/user/releasing.md`, `docs/user/benutzerhandbuch.md`, `CHANGELOG.md` | update | öffentlicher Vertrag |

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
  1×). — **Ausgang:** *(bei Closure zuzuweisen: eingetreten / entfallen / weiter offen)*
- **Pre-Flight deckt die Pipeline nicht:** `make preflight` läuft die Release-Schritte nicht (vgl.
  `BEO-GATE/preflight-deckt-den-ci-schritt-nicht`, 1×). — **Ausgang:** *(bei Closure zuzuweisen: eingetreten / entfallen / weiter offen)*

## 7. Closure-Notiz

*(folgt bei Closure)*

## 8. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** `GATE` (1, 2, 3 ✓) · `USER` (2, 3 ✓).

**Vorgelagert — offene Beobachtungen sichten** (2026-10-07): die zwei Einträge in §6;
`BEO-GATE/hebungskanal-haengt-an-repo-externen-schaltern` (1×) — der arm64-Runner ist ein
repo-externer Schalter; `releasing.md` §Vorbedingungen nennt ihn. Beim Übergang nach `next/`
erneut sichten.

**Modus-Begründungsblock:** alle berührten Sub-Areas GF.
