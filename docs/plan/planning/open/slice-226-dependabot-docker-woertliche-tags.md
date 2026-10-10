# slice-226 — Dependabot hebt auch die Docker-Basen: wörtliche Tags, ein Ort je Version

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `git mv` (`make slice-mv`).

**Welle:** ohne Welle — die Closure-Bedingung wäre die eigene DoD.

**Bezug:** [ADR-0038](../../adr/0038-dependabot-als-hebungskanal.md) (der Kanal; Entscheidung 3
schloss `docker` aus), [ADR-0037](../../adr/0037-cve-scan-gegen-das-publizierte-image.md) (der
Sensor, dessen Befund keinen Weg nach oben hatte), [ADR-0030](../../adr/0030-kein-digest-im-generierten-fragment.md)
(Digest-Hebung als bewusster Commit), [AC-QA-03](../../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit)
(digest-gepinnte Basis). Folge-ADR: [`ADR-0044`](../../adr/0044-dependabot-docker-woertliche-tags.md) (entsteht in diesem Slice).

**Berührte Spec-Stellen:** — · Betriebs-/Prozess-Entscheidung, kein Spec-Stratum.

**Verantwortlich:** —

**Autor:** Claude. **Datum:** 2026-10-10.

**Lerneintrag — Form:** wird bei Closure benannt.

---

## 1. Ziel und Abgrenzung

**Anlass (gemessen, 2026-10-10):** Go 1.27.2 erschien, ohne dass es jemand mitbekam; erst der
nächtliche `image-scan` wurde rot, nachdem ein verwundbares Image draußen war. [ADR-0038](../../adr/0038-dependabot-als-hebungskanal.md) hatte
`docker` ausgeschlossen („ein Digest-Hoist ist ein bewusster Commit") — der PR-Weg schließt den
bewussten Commit aber nicht aus, er sorgt nur dafür, dass jemand von der Hebung erfährt, und
[ADR-0038](../../adr/0038-dependabot-als-hebungskanal.md) Entscheidung 2 (kein Automerge) gilt unverändert.

**Der Befund, der den Slice größer macht als eine Konfigurationszeile:** Dependabots
`FROM`-Parser (dependabot-core `docker/file_parser.rb`, Stand `dc25739e`, 2026-09-16) braucht einen
**wörtlichen Tag**. Die Regex aus der Quelle, gegen die Dockerfiles dieses Repos gefahren:

| `FROM`-Zeile | Ergebnis |
|---|---|
| `golang:${GO_VERSION}@sha256:…` (heute) | image `golang`, tag `nil` → **verworfen, keine Version** |
| `golangci/golangci-lint:${GOLANGCI_LINT_VERSION}@sha256:…` (heute) | **verworfen** |
| `gcr.io/distroless/static-debian12:nonroot@sha256:…` | Abhängigkeit |
| `golang:1.27.2@sha256:…` (wörtlich) | Abhängigkeit |

Ein `docker`-Eintrag in `dependabot.yml` allein wäre also **gültig und wirkungslos** — er sähe
nur `distroless`, nie die Go-Toolchain und nie das Lint-Image ([AGENTS.md](../../../../AGENTS.md)
§4, „ein konfigurierter Block ist erst aktiv, wenn er wirkt").

**Ziel:** Dependabot öffnet PRs für alle drei Basis-Images beider Dockerfiles, und ein Gate hält
das so.

**Liefer-Punkte (3):**

1. **[ADR-0044](../../adr/0044-dependabot-docker-woertliche-tags.md)** (Folge-ADR, `Supersedes [ADR-0038`](../../adr/0038-dependabot-als-hebungskanal.md)) und die `docker`-Einträge in
   `.github/dependabot.yml` für `/` und `/tools/archive-wave`.
2. **Wörtliche Tags in den `FROM`-Zeilen**, und dafür **ein Ort je Version**: die
   `ARG GO_VERSION`/`ARG GOLANGCI_LINT_VERSION` in den Dockerfiles, die gleichnamigen
   Makefile-Variablen und die `--build-arg`-Weitergabe entfallen. Der Grund ist nicht Geschmack:
   ein Dependabot-PR ändert nur die `FROM`-Zeile; stünde die Version noch im Makefile, wäre jeder
   solche PR an `make version-coherence` rot.
3. **Sensor-Regel (3) in `make version-coherence`:** eine `FROM`-Zeile mit `@sha256:` trägt einen
   wörtlichen Tag (kein `$`). Ohne sie könnte ein künftiger Umbau die Wirkung still abschalten.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Automerge.** *Bestand bleibt bewusst stehen* ([ADR-0038](../../adr/0038-dependabot-als-hebungskanal.md) Entscheidung 2): der PR wird mit
  `make ci` geprüft, der Merge bleibt ein Akt.
- **Der Trivy-Pin (`tools/image-scan.sh`) und der d-check-Pin (`d-check.mk`).** *Bestand bleibt
  bewusst stehen*: sie leben in keinem Manifest, das Dependabot liest; sie bleiben bewusste
  Commits.
- **Die Hebung des Lint-Images und der `go`-Direktive selbst.** *Ein Folge-Slice übernimmt es:*
  slice-225. Er bekommt hier seinen Wächter (der erste Lint-PR ist sein Startsignal); sein
  DoD-Wortlaut wird in diesem Slice nachgezogen, weil `GOLANGCI_LINT_VERSION` nicht mehr existiert.
- **Ob Dependabot den Index-Digest eines Multi-Arch-Tags richtig fortschreibt.** *Ein anderer
  Vorgang:* das zeigt erst der erste echte PR; vorher ist es nur behauptbar. Als Risiko in §6.

## 2. Definition of Done

- [ ] [ADR-0044](../../adr/0044-dependabot-docker-woertliche-tags.md) `Accepted`, [ADR-0038](../../adr/0038-dependabot-als-hebungskanal.md) `Superseded by [ADR-0044`](../../adr/0044-dependabot-docker-woertliche-tags.md), ADR-Index nachgezogen.
- [ ] `.github/dependabot.yml`: `docker` für `/` und `/tools/archive-wave`, Präfix mit der Kennung
      der neuen ADR (Trace-Gate gemessen: ohne Kennung rot, mit Kennung grün).
- [ ] Beide Dockerfiles: wörtliche Tags, keine Versions-`ARG`; Makefiles, `tools/image-multiarch.sh`
      und `tools/archive-wave/Makefile` ohne die Variablen; Dockerfile-Zeilen **gebaut** (`make ci`).
- [ ] Gegenprobe der Wirkung: Dependabots Regex gegen die neuen `FROM`-Zeilen ergibt für **alle**
      Basis-Images eine Abhängigkeit (Messung wie in §1, hier mit Ergebnis ins Closure).
- [ ] `make version-coherence`: Regel (3) mit Selbsttest in beiden Richtungen; Gegenprobe am echten
      Dockerfile (`${…}` einbauen, Befund sehen, zurücknehmen).
- [ ] slice-225 §2/§6 nachgezogen (kein `GOLANGCI_LINT_VERSION` mehr; Wächter benannt).
- [ ] Sensor-Datei und Gate-Index-Zeile von `version-coherence` nachgezogen.
- [ ] Closure-Notiz mit Steering-Loop-Lerneintrag; Register fortgeschrieben.
- [ ] Jedes Risiko aus §6 trägt einen Ausgang.

`make ci`, `make archive-wave-test`, `make gates` und `make verify` grün. Ein unabhängiger Review
ist nicht vorgesehen ([MR-019](../../../../harness/conventions.md#mr-019): Opt-in); die Messungen
oben sind der Beleg. Kein CHANGELOG-Eintrag: die Änderung bindet den Harness-Lauf, nicht den
Konsumenten.

## 3. Plan (vor Code)

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| `docs/plan/adr/0044-…md`, `0038-…md`, `adr/README.md` | neu / Status / Index | Folge-ADR vor Code |
| `.github/dependabot.yml` | update | `docker`-Ökosystem, Präfix |
| `Dockerfile`, `tools/archive-wave/Dockerfile` | update | wörtliche Tags, `ARG` raus |
| `Makefile`, `tools/archive-wave/Makefile`, `tools/image-multiarch.sh` | update | Variablen und `--build-arg` raus |
| `tools/verify-versions-kohaerent.sh`, `harness/sensors/version-coherence.md`, `harness/README.md` | update | Regel (3) |
| `docs/user/releasing.md` | update | der Kanal erreicht jetzt auch die Basis-Images |
| `slice-225` | update | DoD-Wortlaut, Wächter |

## 4. Trigger

**Start** (`open` → `in-progress`): Maintainer-Wort („erst pushen und danach einen slice dafür
anlegen und umsetzen", 2026-10-10); WIP-Limit frei.

**Rückführungen:** `in-progress` → `open` (blockiert): die wörtlichen Tags lassen sich nicht ohne
Bruch des Mehr-Plattform-Baus ([ADR-0043](../../adr/0043-multi-arch-ein-bau-getestet-dann-getaggt.md)) einführen.

## 5. Closure-Trigger

DoD vollständig, `make ci`/`gates`/`verify` grün, Closure-Notiz mit Lerneintrag.

## 6. Risiken und offene Punkte

- **Dependabot schreibt den Index-Digest eines Multi-Arch-Tags falsch oder gar nicht fort:** die
  Regex-Messung belegt nur das Erkennen der Zeile, nicht die Fortschreibung. — **Ausgang:** *(bei
  Closure zuzuweisen: eingetreten / entfallen / weiter offen)*
- **PR-Rauschen:** drei Basis-Images in zwei Dockerfiles können wöchentlich bis zu sechs PRs
  öffnen. — **Ausgang:** *(bei Closure zuzuweisen: eingetreten / entfallen / weiter offen)*

## 7. Closure-Notiz

*(folgt bei Closure)*

## 8. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** `GATE` (Achsen 1, 2, 3 ✓), `ADR` (Achsen 1, 2, 3 ✓).

**Vorgelagert — offene Beobachtungen sichten** (2026-10-10):
`BEO-GATE/hebungskanal-haengt-an-repo-externen-schaltern` (1×) — trifft das erste Risiko: der Kanal
hängt an einer Eigenschaft, die das Repo nicht sieht (hier Dependabots Parser); erreicht mit diesem
Slice bei weiter offenem Risiko 2×.
`BEO-GATE/versionsangabe-neben-digest-ungeprueft` (verkörpert) — Regel (2) von `version-coherence`
verliert hier ihren Gegenstand (kein Pin mehr an zwei Orten); die Regel bleibt für künftige
Doppel-Deklarationen und behält ihren Selbsttest.

**Modus-Begründungsblock:** alle berührten Sub-Areas GF.
