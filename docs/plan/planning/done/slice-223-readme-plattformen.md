# slice-223 — README nennt die Plattformen des Images und den Weg auf macOS

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `git mv` (`make slice-mv`).

**Welle:** ohne Welle — die Closure-Bedingung wäre die eigene DoD.

**Bezug:** [AC-FA-DIST-001](../../../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk)
(Plattformen `linux/amd64` und `linux/arm64`, ein Index-Digest).

**Berührte Spec-Stellen:** — · Der Slice berührt kein Spec-Stratum.

**Verantwortlich:** Claude — gesetzt beim Übergang nach `next/` (Maintainer: „jetzt ein Release“, 2026-10-08).

**Autor:** Claude. **Datum:** 2026-10-08.

**Lerneintrag — Form:** geschärfte Regel.

---

## 1. Ziel und Abgrenzung

**Ziel:** `README.md` und `README.de.md` sagen, dass das Image seit `v0.23.0` für `linux/amd64` und
`linux/arm64` unter **einem** Digest erscheint und auf macOS mit Apple Silicon nativ läuft, mit
Zeiger auf den macOS-Abschnitt des Benutzerhandbuchs. Anlass: die Prüfung vor dem Release
(Maintainer-Frage 2026-10-08, „ist das Handbuch und README.md aktuell") — das Handbuch war es, die
READMEs nicht.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Weitere README-Abschnitte umschreiben.** *Bestand bleibt bewusst stehen*: nur die
  Distributions-Aussage war veraltet.
- **Das Release.** *Ein anderer Vorgang*: folgt nach diesem Slice nach `releasing.md`.

## 2. Definition of Done

- [x] `README.md` und `README.de.md`: Distributions-Punkt nennt die Plattformen, den einen Digest und
      verlinkt den macOS-Abschnitt des Handbuchs; CHANGELOG `[Unreleased]`.
- [x] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [x] Beobachtungs-Register fortgeschrieben (oder „keine Beobachtung" notiert).
- [x] Jedes Risiko aus §6 trägt einen Ausgang.

`make gates` und `make verify` grün. Ein unabhängiger Review ist für diesen Slice nicht vorgesehen
([MR-019](../../../../harness/conventions.md#mr-019): Opt-in) — zwei Sätze Doku je Sprache.

## 3. Plan (vor Code)

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| `README.md`, `README.de.md`, `CHANGELOG.md` | update | Distributions-Aussage |

## 4. Trigger

**Start** (`open` → `in-progress`): Maintainer-Wort („jetzt ein Release"), WIP-Limit frei.

**Rückführungen:** entfallen — zwei Sätze.

## 5. Closure-Trigger

DoD vollständig, `make gates` und `make verify` grün, Closure-Notiz mit Lerneintrag.

## 6. Risiken und offene Punkte

Kein Risiko notiert — zwei Sätze Doku je Sprache, ohne Code.

## 7. Closure-Notiz

**Lerneintrag — Form: geschärfte Regel.** Eine Vertragsänderung, die Nutzer sehen, endet nicht beim
Handbuch: die READMEs tragen dieselbe Aussage in Kurzform. Die DoD von slice-218 zählte Handbuch,
`releasing.md`, Hub-Seite und CHANGELOG auf — die READMEs fehlten, und erst die Frage des
Maintainers vor dem Release fand es. Gezählt, nicht verkörpert: neu im Register (1×).

**Geliefert:** Distributions-Punkt in `README.md` und `README.de.md` (Plattformen, ein Digest,
Verweis auf den macOS-Abschnitt), CHANGELOG `[Unreleased]`.

**Was hat funktioniert:** die Prüfung vor dem Release als Frage, nicht als Annahme.

**Was ging anders als geplant:** nichts im Slice selbst; die Lücke lag im Vorgänger.

**Steering-Loop-Eintrag:** geschärfte Regel — gezählt, nicht verkörpert.

**Beobachtungs-Register (`../observations/`):** neu
`BEO-USER/readme-hinkt-der-vertragsaenderung-nach` (1×).

**Folge-Slices:** keine.

**Risiken aus §6:** kein Risiko notiert.

**Drei Paarungen:** Anker — kein `liegt in`-Feld · Folge-Slice — keiner · Register — der genannte
Pfad existiert mit nicht leerem `evidence/`.

**Trigger-Audit der aktiven MR:** [`MR-016`](../../../../harness/conventions.md#mr-016) [`MR-019`](../../../../harness/conventions.md#mr-019) [`MR-025`](../../../../harness/conventions.md#mr-025) [`MR-027`](../../../../harness/conventions.md#mr-027) [`MR-028`](../../../../harness/conventions.md#mr-028) [`MR-029`](../../../../harness/conventions.md#mr-029) [`MR-030`](../../../../harness/conventions.md#mr-030) — 0 offen (geprüft 2026-10-08).

## 8. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** `USER` — die READMEs gehören zur Benutzer-Doku im weiteren
Sinn; `README.md` ist Rang 7 der Source Precedence.

**Vorgelagert — offene Beobachtungen sichten** (2026-10-08):
`BEO-USER/werkzeug-aussage-weiter-als-die-quelle` (1×) — der Satz nennt nur, was belegt ist
(Pipeline-Test auf Linux-arm64, Maintainer-Probe unter Colima); `BEO-USER/handbuch-vokabel-der-adapter-rolle`
(1×) berührt die READMEs nicht.

**Modus-Begründungsblock:** alle berührten Sub-Areas GF.
