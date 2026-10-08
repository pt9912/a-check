# slice-223 — README nennt die Plattformen des Images und den Weg auf macOS

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `git mv` (`make slice-mv`).

**Welle:** ohne Welle — die Closure-Bedingung wäre die eigene DoD.

**Bezug:** [AC-FA-DIST-001](../../../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk)
(Plattformen `linux/amd64` und `linux/arm64`, ein Index-Digest).

**Berührte Spec-Stellen:** — · Der Slice berührt kein Spec-Stratum.

**Verantwortlich:** —

**Autor:** Claude. **Datum:** 2026-10-08.

**Lerneintrag — Form:** wird bei Closure benannt.

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

- [ ] `README.md` und `README.de.md`: Distributions-Punkt nennt die Plattformen, den einen Digest und
      verlinkt den macOS-Abschnitt des Handbuchs; CHANGELOG `[Unreleased]`.
- [ ] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [ ] Beobachtungs-Register fortgeschrieben (oder „keine Beobachtung" notiert).
- [ ] Jedes Risiko aus §6 trägt einen Ausgang.

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

- **Keine.** *(Kein Risiko notiert.)*

## 7. Closure-Notiz

*(folgt bei Closure)*

## 8. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** `USER` — die READMEs gehören zur Benutzer-Doku im weiteren
Sinn; `README.md` ist Rang 7 der Source Precedence.

**Vorgelagert — offene Beobachtungen sichten** (2026-10-08):
`BEO-USER/werkzeug-aussage-weiter-als-die-quelle` (1×) — der Satz nennt nur, was belegt ist
(Pipeline-Test auf Linux-arm64, Maintainer-Probe unter Colima); `BEO-USER/handbuch-vokabel-der-adapter-rolle`
(1×) berührt die READMEs nicht.

**Modus-Begründungsblock:** alle berührten Sub-Areas GF.
