# slice-217 — Multi-Arch-Image: Build für `linux/amd64` und `linux/arm64`

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `git mv` (`make slice-mv`).

**Welle:** welle-18 — [Welle-Plan](../welle-18-multi-arch-image.md).

**Bezug:** [AC-FA-DIST-001](../../../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk),
[AC-QA-03](../../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit); die ADR aus slice-216.

**Berührte Spec-Stellen:** `spezifikation.md`
§[SPEC-DIST-001](../../../../spec/spezifikation.md#spec-dist-001--laufzeitform-und-distribution)
(in slice-216 geschrieben, hier umgesetzt).

**Verantwortlich:** Claude — gesetzt beim Übergang nach `next/` (Maintainer: „ok“, 2026-10-07).

**Autor:** Claude. **Datum:** 2026-10-07.

**Lerneintrag — Form:** wird bei Closure benannt.

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

- [ ] Dockerfile: Build-Stufe auf der Plattform des Runners, Ziel-Plattform aus den
      Build-Argumenten; Laufzeit-Stufe ohne `RUN`.
- [ ] Make-Target für den Multi-Arch-Bau nach der ADR aus slice-216, im Gate-Index
      (`harness/README.md` §Sensors oder §Nicht-Gates) eingetragen.
- [ ] Image-Test gegen eine übergebene Bild-Referenz (Tag oder Digest) statt fest
      `$(IMAGE):dev`, Plattform des Hosts geprüft; lokaler Beleg: beide Plattform-Bilder gebaut, das
      arm64-Binary ist ein arm64-ELF, das amd64-Bild besteht den Image-Test — mit Gegenprobe.
- [ ] Unabhängiger Review, Report unter [`docs/reviews/`](../../../reviews/README.md).
- [ ] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [ ] Beobachtungs-Register fortgeschrieben (oder „keine Beobachtung" notiert).
- [ ] Jedes Risiko aus §6 trägt einen Ausgang.

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
  das veröffentlichte Bild. — **Ausgang:** *(bei Closure zuzuweisen: eingetreten / entfallen / weiter offen)*

## 7. Closure-Notiz

*(folgt bei Closure)*

## 8. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** `GATE` (1, 2, 3 ✓).

**Vorgelagert — offene Beobachtungen sichten** (2026-10-07):
`BEO-GATE/preflight-deckt-den-ci-schritt-nicht` (1×) — ein neues Target muss im Pre-Flight und in
der CI dieselbe Menge sein; `BEO-GATE/werkzeug-zugeschriebene-leistung` (2×) — „Multi-Arch" nicht
dem Werkzeug zuschreiben, sondern am Index messen. Beim Übergang nach `next/` erneut sichten.

**Modus-Begründungsblock:** alle berührten Sub-Areas GF.
