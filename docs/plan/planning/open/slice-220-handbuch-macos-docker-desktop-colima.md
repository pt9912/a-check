# slice-220 — Benutzerhandbuch: a-check auf macOS mit Docker Desktop und Colima

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `git mv` (`make slice-mv`).

**Welle:** ohne Welle — die Closure-Bedingung wäre die eigene DoD.

**Bezug:** [AC-FA-DIST-001](../../../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk)
(Plattformen `linux/amd64` und `linux/arm64`),
[ADR-0043](../../adr/0043-multi-arch-ein-bau-getestet-dann-getaggt.md).

**Berührte Spec-Stellen:** — · Der Slice berührt kein Spec-Stratum; er beschreibt den Betrieb.

**Verantwortlich:** —

**Autor:** Claude. **Datum:** 2026-10-07.

**Lerneintrag — Form:** wird bei Closure benannt.

---

## 1. Ziel und Abgrenzung

**Ziel:** Das Benutzerhandbuch sagt einem macOS-Nutzer, wie a-check unter **Docker Desktop** und
unter **Colima** läuft: welche Plattform gezogen wird, woran er das prüft, und welche Verzeichnisse
die Linux-VM sieht — ein Repo außerhalb der freigegebenen Pfade erscheint im Container leer.
Jede Aussage über ein Werkzeug stützt sich auf dessen Dokumentation; was dieses Projekt selbst
nicht geprüft hat, sagt der Text.

**Quellen (gelesen 2026-10-07):**

- Colima, `embedded/defaults/colima.yaml` (Branch `main`): `arch` „Default: host", „value cannot be
  changed after virtual machine is created"; Mounts „Colima default behaviour: $HOME is mounted as
  writable".
- Docker Desktop, Doku *Settings*, File sharing auf Mac: „By default the `/Users`, `/Volumes`,
  `/private`, `/tmp` and `/var/folders` directory are shared."

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Ein Test auf macOS.** *Ein anderer Vorgang*: dieses Projekt hat keinen macOS-Rechner im Lauf;
  geprüft ist das arm64-Bild auf einem nativen Linux-arm64-Runner. Der Text benennt das.
- **Andere Laufzeiten (Podman, Rancher Desktop, OrbStack).** *Bestand bleibt bewusst stehen*: kein
  Bedarf benannt; `DOCKER ?=` im Fragment erlaubt sie, ohne dass das Handbuch sie beschreibt.
- **Code, Fragment, Spezifikation.** *Schicht-Abgrenzung*: reine Benutzer-Doku.

## 2. Definition of Done

- [ ] Benutzerhandbuch: Abschnitt „macOS: Docker Desktop und Colima" (Plattform, Prüf-Kommando,
      freigegebene Verzeichnisse, Hinweis auf `colima start --arch`), Versions-Bump und
      Historie-Zeile; CHANGELOG `[Unreleased]`.
- [ ] Unabhängiger Review, Report unter [`docs/reviews/`](../../../reviews/README.md).
- [ ] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [ ] Beobachtungs-Register fortgeschrieben (oder „keine Beobachtung" notiert).
- [ ] Jedes Risiko aus §6 trägt einen Ausgang.

`make gates` und `make verify` grün.

## 3. Plan (vor Code)

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| `docs/user/benutzerhandbuch.md` | update | neuer Abschnitt in §1 |
| `CHANGELOG.md` | update | Benutzer-Doku ist öffentlicher Vertrag (AGENTS §6 Schritt 7) |

## 4. Trigger

**Start** (`open` → `in-progress`): WIP-Limit frei; Maintainer-Wort („ja, mach den Slice fürs
Handbuch", 2026-10-07).

**Rückführungen:** `in-progress` → `next` (zu groß): entfällt — ein Abschnitt.
`in-progress` → `open` (blockiert): eine Quelle widerspricht sich oder ist nicht auffindbar.

## 5. Closure-Trigger

DoD vollständig, `make gates` und `make verify` grün, Closure-Notiz mit Lerneintrag.

## 6. Risiken und offene Punkte

- **Fremde Voreinstellungen ändern sich:** Docker Desktop und Colima können ihre Defaults ändern;
  der Text altert dann still. — **Ausgang:** *(bei Closure zuzuweisen: eingetreten / entfallen / weiter offen)*

## 7. Closure-Notiz

*(folgt bei Closure)*

## 8. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** `USER` (Achsen 2, 3 ✓).

**Vorgelagert — offene Beobachtungen sichten** (2026-10-07):
`BEO-USER/handbuch-vokabel-der-adapter-rolle` (1×) berührt diesen Abschnitt nicht;
`BEO-GATE/zusage-weiter-als-ihre-durchsetzung` (5×, geplant slice-219) — der Abschnitt sagt
darum, was geprüft ist und was nicht.

**Modus-Begründungsblock:** alle berührten Sub-Areas GF.
