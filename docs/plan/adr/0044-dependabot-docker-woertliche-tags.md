# ADR-0044 — Dependabot hebt auch die Docker-Basen: wörtliche Tags, ein Ort je Version

- **Status:** Accepted
- **Datum:** 2026-10-10
- **Autor:** pt9912
- **Bezug:** [AC-QA-03](../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit); [`AGENTS.md`](../../../AGENTS.md) §4 (ein konfigurierter Block ist erst aktiv, wenn er wirkt), §5 (Traceability-Pflicht je Commit); [ADR-0037](0037-cve-scan-gegen-das-publizierte-image.md) (der Sensor, der meldet); [ADR-0030](0030-kein-digest-im-generierten-fragment.md) (Digest-Hebung als bewusster Commit); [ADR-0043](0043-multi-arch-ein-bau-getestet-dann-getaggt.md) (der Mehr-Plattform-Bau, den die Basis-Images tragen)
- **Schärft:** — (Betriebs-/Prozess-Entscheidung; koppelt an keine Spec-§. Dieselbe Verortung wie [ADR-0038](0038-dependabot-als-hebungskanal.md).)
- **Supersedes:** [ADR-0038](0038-dependabot-als-hebungskanal.md)

## Kontext

[ADR-0038](0038-dependabot-als-hebungskanal.md) hat Dependabot als Hebungs-Kanal eingeführt und
`docker` bewusst ausgeschlossen: „ein Digest-Hoist ist ein bewusster Commit". Zwei Dinge haben
sich seither gezeigt.

**Der Ausschluss ließ eine Befund-Klasse ohne Weg nach oben.** Am 2026-10-09 meldete der nächtliche
`image-scan` (ADR-0037) drei behebbare HIGH in der Go-Standardbibliothek des publizierten Images;
behoben waren sie in Go 1.27.2, das schon erschienen war. Niemand hatte es mitbekommen — der Kanal
sah die Basis-Images nicht. Der Weg „ein Mensch sieht den Scan und hebt" ist der, den ADR-0038
selbst als nicht tragend beschreibt: er verlagert die Zusage vom Werkzeug auf die Aufmerksamkeit.
Dass ein Pin-Wechsel ein bewusster Commit bleibt, widerspricht dem PR-Weg nicht — der PR *ist* der
Commit-Entwurf, und der Merge bleibt ein Akt (ADR-0038 Entscheidung 2).

**Ein `docker`-Eintrag allein wirkte nicht.** Dependabots `FROM`-Parser (dependabot-core,
`docker/lib/dependabot/docker/file_parser.rb` mit `shared/shared_file_parser.rb`, Stand `dc25739e`,
2026-09-16) macht aus einer Zeile nur dann eine Abhängigkeit, wenn sie einen **wörtlichen Tag oder
Digest** trägt (`version_from` liefert `tag || digest`). Die Regex aus der Quelle, wortgleich gegen
die Dockerfiles dieses Repos gefahren:

| `FROM`-Zeile | Ergebnis |
|---|---|
| `golang:${GO_VERSION}@sha256:…` | image `golang`, Tag `nil` → **verworfen** |
| `golangci/golangci-lint:${GOLANGCI_LINT_VERSION}@sha256:…` | **verworfen** |
| `gcr.io/distroless/static-debian12:nonroot@sha256:…` | Abhängigkeit |
| `golang:1.27.2@sha256:…` | Abhängigkeit |

Mit dem heutigen Dockerfile sähe der Kanal nur `distroless` — nie die Go-Toolchain, nie das
Lint-Image. Die Eintrag wäre gültig und wirkungslos.

**Die Version steht heute an zwei Orten.** Makefile (`GO_VERSION`, `GOLANGCI_LINT_VERSION`) und
Dockerfile-`ARG` führen dieselbe Zahl; `make version-coherence` hält sie gleich. Neben dem Digest
ist die Zahl ohnehin nicht wirksam (der Digest sticht den Tag); sie benennt, was im Image liegt.
Ein Dependabot-PR ändert aber nur die `FROM`-Zeile — stünde die Zahl noch im Makefile, wäre jeder
solche PR an diesem Gate rot.

## Entscheidung

**1. Die Entscheidungen 1 und 2 aus ADR-0038 gelten unverändert weiter.** Jeder
`commit-message.prefix` trägt die Kennung eines Entscheids, hier **dieser** (die vorige ist
ersetzt); `commits.exempt-pattern` bleibt unverändert. Kein Automerge: der Kanal öffnet PRs, geprüft
wird mit `make ci`, der Merge bleibt ein Akt.

**2. Drei Ökosysteme.** Module, GitHub Actions und `docker` — für das Wurzel-Dockerfile und das
Dockerfile von `tools/archive-wave/`. Der Trivy-Pin (`tools/image-scan.sh`) und der
Schwester-Tool-Pin (`d-check.mk`) bleiben draußen: sie leben in keinem Manifest, das der Kanal
liest, und sind weiter bewusste Commits.

**3. Wörtliche Tags in jeder `FROM`-Zeile mit Digest.** `golang:1.27.2@sha256:…`, nicht
`golang:${GO_VERSION}@sha256:…`. Die Versions-`ARG` der Dockerfiles entfallen.

**4. Ein Ort je Version.** Die gleichnamigen Makefile-Variablen und ihre `--build-arg`-Weitergabe
entfallen; die `FROM`-Zeile ist die einzige Deklaration. Wo eine Zahl doppelt deklariert wird,
hält `make version-coherence` sie weiter gleich — die Regel bleibt, der Bestand, an dem sie heute
greift, verschwindet.

**5. Ein Gate hält die Wirkung.** `make version-coherence` verlangt, dass eine `FROM`-Zeile mit
`@sha256:` einen wörtlichen Tag trägt. Ohne diese Regel könnte ein künftiger Umbau (zurück zu einer
Variable) den Kanal still abschalten, und `dependabot.yml` bliebe grün.

## Verglichene Alternativen

| Option | Pro | Contra |
|---|---|---|
| A — Nichts tun (ADR-0038 gilt) | kein Umbau; Hebung bleibt ein Mensch-Akt | gemessen am 2026-10-09: das verwundbare Image war draußen, bevor jemand hinsah; der Scan meldet nach dem Release, nicht davor |
| B — `docker`-Eintrag ohne Tag-Umbau | eine Konfigurationszeile | wirkungslos für Toolchain und Lint-Image (Tabelle oben); der Eintrag meldete Grün, ohne etwas zu sehen |
| C — eigener Melder (Skript fragt die Registry nach neueren Tags) | unabhängig von Dependabots Parser; ließe die `ARG` stehen | ein zweiter Mechanismus mit eigener Pflege, Netz-Zugriff, eigener PR-Logik; `gates` bliebe hermetisch nur, wenn er außerhalb liegt |
| **D — wörtliche Tags, `docker`-Ökosystem, Gate** | der Standard-Kanal trägt alles; ein Ort je Version | Umbau der Dockerfiles; bis zum ersten echten PR ist nur das **Erkennen** der Zeile belegt, nicht die Fortschreibung des Index-Digests |

## Konsequenzen

- **Belegt ist das Erkennen, nicht das Fortschreiben.** Die Regex-Messung zeigt, dass jede
  Basis-Image-Zeile eine Abhängigkeit ergibt. Dass Dependabot den **Index-Digest** eines
  Multi-Arch-Tags (ADR-0043) richtig mitführt, zeigt erst der erste echte PR. Bis dahin ist das
  eine offene Annahme, keine Zusage.
- **Der Kanal hängt weiter an Schaltern außerhalb des Repos** (`dependabot_security_updates`,
  Alerts — [`docs/user/releasing.md`](../../user/releasing.md) §Vorbedingungen) und jetzt an
  Dependabots Parser. Beides sieht kein Gate des Repos; die Regel aus Entscheidung 5 deckt nur die
  Form der Zeilen.
- **Mehr PRs.** Drei Basis-Images in zwei Dockerfiles; das Limit von fünf offenen PRs je Ökosystem
  aus ADR-0038 bleibt die Gegenmaßnahme gegen Rauschen, nicht die Lösung.
- **Jeder Basis-Wechsel läuft durch `make ci`.** Das ist die Prüfung aus ADR-0030: Digest aus der
  Quelle, Gates gelesen, Wirkung gemessen — der PR ersetzt nur das Heraussuchen.
- **Die Go-Mindestversion im `go.mod` ist kein Teil davon.** Ein Toolchain-PR ändert die
  Build-Basis, nicht die `go`-Direktive; die steht in einem eigenen Manifest und hat eine eigene
  Kopplung zum Lint-Image.

## Fitness Function

Zwei Messungen, beide Richtungen:

1. **Erkennen:** Dependabots `FROM`-Regex (aus der Quelle, Stand oben) gegen beide Dockerfiles ergibt
   für **jede** Zeile mit Digest, die ein Basis-Image nennt, eine Abhängigkeit. Die Gegenprobe —
   dieselbe Zeile mit `${…}` im Tag — ergibt keine.
2. **Halten:** `make version-coherence` meldet eine `FROM`-Zeile mit Digest und Variablen-Tag als
   Befund; die wörtliche Form ist grün. Selbsttest in beiden Richtungen, dazu die Gegenprobe am
   echten Dockerfile.
