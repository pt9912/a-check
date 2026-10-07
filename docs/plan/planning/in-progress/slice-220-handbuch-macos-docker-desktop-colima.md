# slice-220 — Benutzerhandbuch: a-check auf macOS mit Docker Desktop und Colima

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `git mv` (`make slice-mv`).

**Welle:** ohne Welle — die Closure-Bedingung wäre die eigene DoD.

**Bezug:** [AC-FA-DIST-001](../../../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk)
(Plattformen `linux/amd64` und `linux/arm64`),
[ADR-0043](../../adr/0043-multi-arch-ein-bau-getestet-dann-getaggt.md).

**Berührte Spec-Stellen:** — · Der Slice berührt kein Spec-Stratum; er beschreibt den Betrieb.

**Verantwortlich:** Claude — gesetzt beim Übergang nach `next/` (Maintainer: „ja, mach den Slice fürs Handbuch“, 2026-10-07).

**Autor:** Claude. **Datum:** 2026-10-07.

**Lerneintrag — Form:** geschärfte Regel.

---

## 1. Ziel und Abgrenzung

**Ziel:** Das Benutzerhandbuch sagt einem macOS-Nutzer, wie a-check unter **Docker Desktop** und
unter **Colima** läuft: welche Plattform gezogen wird, woran er das prüft, und welche Verzeichnisse
die Linux-VM sieht — ein Repo außerhalb der freigegebenen Pfade erscheint unter Colima leer, unter
Docker Desktop scheitert der Mount.
Jede Aussage über ein Werkzeug stützt sich auf dessen Dokumentation; was dieses Projekt selbst
nicht geprüft hat, sagt der Text.

**Quellen (gelesen 2026-10-07):**

- Colima, `embedded/defaults/colima.yaml` (Branch `main`): `arch` „Default: host", „value cannot be
  changed after virtual machine is created"; Mounts „Colima default behaviour: $HOME is mounted as
  writable".
- Docker Desktop, Doku *Settings*, File sharing auf Mac: „By default the `/Users`, `/Volumes`,
  `/private`, `/tmp` and `/var/folders` directory are shared."

**Plan-Änderung 2026-10-07 (Review F-1/F-2/F-3, vor dem Fix):** Das Ziel verallgemeinerte das
leere Verzeichnis auf beide Laufzeiten; die Docker-Doku sagt für Docker Desktop „Mounts denied"
(Satz direkt nach dem zitierten). Die Quellenliste deckte nur zwei Aussagen; ergänzt sind:

- Docker Desktop, *Settings*: „otherwise you may get Mounts denied or cannot start service errors
  at runtime" (Fortsetzung des File-sharing-Absatzes).
- Docker, *Bind mounts*: `-v` legt eine fehlende Quelle als Verzeichnis an.
- Colima-Quelltext (`main`, laut Review): `colima status` gibt `arch:` aus; `--mount`/`-V`
  existiert (Default schreibgeschützt); `colima start` bei laufender VM ändert nichts; eine eigene
  Mount-Liste ersetzt den Standard-Mount von `$HOME`; Konfigurationsort
  `~/.colima/default/colima.yaml`, sofern weder `XDG_CONFIG_HOME` noch `~/.config/colima` greift.
- **Maintainer-Probe:** `v0.23.0` lief beim Maintainer erfolgreich auf macOS unter Colima, und
  `docker image inspect --format '{{.Architecture}}'` zeigte `arm64` (Mitteilung 2026-10-07) —
  eine einmalige Probe, nicht Teil der Release-Pipeline.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Ein Test auf macOS.** *Ein anderer Vorgang*: dieses Projekt hat keinen macOS-Rechner im Lauf;
  geprüft ist das arm64-Bild auf einem nativen Linux-arm64-Runner. Der Text benennt das.
- **Andere Laufzeiten (Podman, Rancher Desktop, OrbStack).** *Bestand bleibt bewusst stehen*: kein
  Bedarf benannt; `DOCKER ?=` im Fragment erlaubt sie, ohne dass das Handbuch sie beschreibt.
- **Code, Fragment, Spezifikation.** *Schicht-Abgrenzung*: reine Benutzer-Doku.

## 2. Definition of Done

- [x] Benutzerhandbuch: Abschnitt „macOS: Docker Desktop und Colima" (Plattform, Prüf-Kommando,
      freigegebene Verzeichnisse, Hinweis auf `colima start --arch`), Versions-Bump und
      Historie-Zeile; CHANGELOG `[Unreleased]`.
- [x] Unabhängiger Review, Report unter [`docs/reviews/`](../../../reviews/README.md).
- [x] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [x] Beobachtungs-Register fortgeschrieben (oder „keine Beobachtung" notiert).
- [x] Jedes Risiko aus §6 trägt einen Ausgang.

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
  der Text altert dann still. — **Ausgang:** *weiter offen* — der Abschnitt nennt seinen Stand
  (2026-10-07); das Beobachtungs-Register führt die Klasse unter
  `BEO-USER/werkzeug-aussage-weiter-als-die-quelle`.

## 7. Closure-Notiz

**Lerneintrag — Form: geschärfte Regel.** Eine Aussage über ein fremdes Werkzeug im Handbuch steht
erst, wenn sie Satz für Satz an dessen Quelle gelesen ist — einschließlich der Sätze **nach** dem
zitierten: der Fehler dieses Slice (leeres Verzeichnis auch unter Docker Desktop) stand in der
Docker-Doku einen Satz weiter. Die Quellen stehen im Plan neben der Aussage, die sie tragen, und
der Text sagt, was das Projekt selbst geprüft hat und was nicht. Verkörpert ist die Regel hier im
Plan und im Abschnitt selbst; allgemein trägt sie slice-219 (Zusage nennt Lauf und Grenze).

**Geliefert:** Benutzerhandbuch 1.47, §1 „macOS: Docker Desktop und Colima" — gezogene Plattform
und Prüf-Kommando, Prüfstand (Linux-arm64-Runner je Release; einmalige Maintainer-Probe unter
Colima mit `arm64`), freigegebene Verzeichnisse, Verhalten außerhalb (Docker Desktop: Mount-Fehler;
Colima: leerer Baum), Colima-Mounts und -Architektur; CHANGELOG `[Unreleased]`.

**Was hat funktioniert:** Die Gegenprobe des Reviewers am Quelltext von Colima, nicht an meinen
Zitaten — sie fand zwei Bedienfallen (Mount bei laufender VM, Mount-Liste ersetzt `$HOME`), die
keine Dokumentationsseite nennt.

**Was ging anders als geplant:** Ich hatte die eine Quelle gelesen und den nächsten Satz nicht —
die Aussage war für Colima richtig und für Docker Desktop falsch. Und die Arbeitsdateien des
Reviewers lagen kurz im Repo und machten `make lint` rot (Colima-Quelltext als `*.go`).

**Steering-Loop-Eintrag:** geschärfte Regel — gezählt, nicht verkörpert; die allgemeine Form
trägt slice-219.

**Beobachtungs-Register (`../observations/`):**
`BEO-USER/werkzeug-aussage-weiter-als-die-quelle` neu (1×);
`BEO-GATE/zusage-weiter-als-ihre-durchsetzung` → 6× (Ausgang bleibt *geplant*, slice-219).

**Folge-Slices:** keine neuen; slice-219 liegt in `open/`.

**Risiken aus §6:** das eine Risiko trägt seinen Ausgang (*weiter offen*, Beobachtungs-Register).

**Drei Paarungen:** Anker — kein `liegt in`-Feld, nichts zu paaren · Folge-Slice — slice-219
existiert in `open/` · Register — beide genannten Pfade existieren mit nicht leerem `evidence/`.

**Trigger-Audit der aktiven MR:** [`MR-016`](../../../../harness/conventions.md#mr-016) [`MR-019`](../../../../harness/conventions.md#mr-019) [`MR-025`](../../../../harness/conventions.md#mr-025) [`MR-027`](../../../../harness/conventions.md#mr-027) [`MR-028`](../../../../harness/conventions.md#mr-028) [`MR-029`](../../../../harness/conventions.md#mr-029) [`MR-030`](../../../../harness/conventions.md#mr-030) — 0 offen (geprüft 2026-10-07).

## 8. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** `USER` (Achsen 2, 3 ✓).

**Vorgelagert — offene Beobachtungen sichten** (2026-10-07):
`BEO-USER/handbuch-vokabel-der-adapter-rolle` (1×) berührt diesen Abschnitt nicht;
`BEO-GATE/zusage-weiter-als-ihre-durchsetzung` (5×, geplant slice-219) — der Abschnitt sagt
darum, was geprüft ist und was nicht.

**Modus-Begründungsblock:** alle berührten Sub-Areas GF.
