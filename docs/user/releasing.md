# Releasing — a-check

Release-Prozess für `ghcr.io/pt9912/a-check`
([AC-FA-DIST-001](../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk),
[ADR-0004](../plan/adr/0004-distribution-image-mk.md),
[ADR-0007](../plan/adr/0007-latest-tag-politik.md)). Die Pipeline ist
[`.github/workflows/release.yml`](../../.github/workflows/release.yml) (seit
slice-007). Das **aktuelle Release** samt Versions-Koordinaten (Version, Datum,
Digest, Tag) führt das [Release-Register `version.md#aktuell`](../../version.md#aktuell);
der [Verlauf](../../version.md#verlauf) listet alle Releases ab `v0.1.0`. Der
GitHub-Release-Link und die GHCR-Tags (`+ latest`) stehen dort.

## Aktueller Stand

Die [aktuelle Version](../../version.md#aktuell) ist auf GHCR verfügbar;
Konsumenten pinnen den **Digest** ([Konsum](#konsum-digest-pin)). Das mitgelieferte
[`a-check.mk`](../../a-check.mk) und `a-check --print-mk`
([AC-FA-DIST-001](../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk))
sind auf den im [Release-Register](../../version.md#aktuell) geführten `@sha256:`-Digest
gepinnt — `tools/gate-consistency.sh` prüft die Gleichheit dieser harten Pins mit
`version.md#aktuell` (slice-018). Für lokale Entwicklung gegen ungetaggte Stände
dient weiterhin das lokal gebaute Image:

```sh
make build                               # baut a-check:dev (static/distroless)
make a-check A_CHECK_IMAGE=a-check:dev   # Konsum-Aufruf gegen das lokale Image
```

## Versionsquelle

Versionen folgen SemVer; die menschlich kuratierte Begründung jedes Releases ist
der zugehörige Abschnitt in [`CHANGELOG.md`](../../CHANGELOG.md). Der
`[Unreleased]`-Stand wird **beim Re-Pin nach dem Publish** (Schritt 6) unter die neue
Version geschnitten — **zusammen** mit dem `version.md#aktuell`-Bump: der Pin-Check
(slice-018) verlangt `version.md#aktuell` == aktuellstes `CHANGELOG`-Release; ein
CHANGELOG-Schnitt **vor** dem Tag (ohne version.md-Bump) macht `make ci` in der Pipeline
rot. So bleibt `version.md#aktuell` bis zum Publish beim vorigen Release (immer wahr —
nie ein unveröffentlichter Digest). Das Lastenheft steht bei **0.17.0** (eigene Doku-Achse,
**≠** der Release-Tag: v0.10.0 trägt Lastenheft 0.16.0).

## Release auslösen

```sh
git tag vX.Y.Z          # die neue Version
git push origin vX.Y.Z
```

Die Pipeline ([`release.yml`](../../.github/workflows/release.yml)) läuft bei
jedem `v*`-Tag-Push:

1. **SemVer-Validate** (fail-fast): nur `vMAJOR.MINOR.PATCH` oder
   `…-PRERELEASE`; Build-Metadaten (`+`) werden abgelehnt — vor Login/Build/Push.
2. **`make ci VERSION=<version>`** — alle Gates (`make gates`) **plus**
   `image-test`; baut zugleich das Runtime-Image mit `VERSION` aus dem Tag
   (→ OCI-Label `org.opencontainers.image.version`).
3. **OCI-Label-Verify** — `org.opencontainers.image.version` muss exakt der
   Tag-Version entsprechen (Version-Drift shippt nicht).
4. **Push** nach `ghcr.io/pt9912/a-check:v<version>`; `:latest`
   **ausschließlich** für stabile Releases (kein Prerelease-Suffix) —
   [ADR-0007](../plan/adr/0007-latest-tag-politik.md). Konsumenten pinnen
   Digests, nicht `:latest`.
5. **Digest-Pin** im Job-Summary und in den Notes des angelegten GitHub-Releases;
   danach gibt `a-check --print-mk` ein `a-check.mk` mit dem **aktuell
   digest-gepinnten** `A_CHECK_IMAGE` aus
   ([AC-QA-03](../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit)).
6. **Register-Re-Pin + CHANGELOG-Schnitt** (slice-018): den CHANGELOG `[Unreleased]` →
   `[X.Y.Z] - <Datum>` schneiden **und** — im **selben** Commit, sonst Versions-Drift —
   den neuen Digest + die neue Version an [`version.md#aktuell`](../../version.md#aktuell)
   (Version, Datum, voller `@sha256:`-Digest) plus eine neue
   [Verlaufs](../../version.md#verlauf)-Zeile und den wandernden `<a id>`-Anker nachziehen.
   Die harten Pins ([`a-check.mk`](../../a-check.mk),
   [`internal/cli/cli.go`](../../internal/cli/cli.go)) und die `docker run`-Beispiele in
   **beiden** READMEs ([`README.md`](../../README.md) + [`README.de.md`](../../README.de.md))
   tragen den Digest verbatim; vergisst der Re-Pin eine davon, meldet
   `make gates` (via `tools/gate-consistency.sh`) die Digest- bzw. Versions-Drift
   **rot** — statt sie wie früher nur per Zufalls-Audit aufzudecken. Die Prosa-Erwähnungen
   („Status", „aktuelles Release", Handbuch-Software-Version) verlinken auf
   `version.md#aktuell` und brauchen **keinen** Bump.

## Freigabe-Checkliste

Produktionsreife heißt **belegte** Betriebsfähigkeit, nicht „gebaut". Darum gilt hier:
**kein Häkchen ohne Beleg.** Fehlt der Beleg, ist das Item *nicht* abgehakt — auch wenn es
inhaltlich erfüllt wäre. Die Beleg-Pflicht ist der einzige Schutz davor, dass die Liste zur
Bürokratie verkommt. Ein Punkt je Phase (slice-051):

| # | Item | Beleg-Slot |
|---|---|---|
| 1 | **Spec** — jede neue/geänderte `AC-*` liegt im Lastenheft, keine ADR hat es geschärft | Lastenheft-Version + `SPEC-*`/`ARC-*`-Bumps im Diff |
| 2 | **Architektur** — jede referenzierte ADR ist `Accepted` oder ausgewiesen `Superseded` | `make doc-immutable` über die Release-Range, Exit 0 |
| 3 | **Planung** — die gelieferten Slices liegen in `done/` und tragen eine ausgefüllte Closure-Notiz | `make verify`, Exit 0 |
| 4 | **Agenten** — `AGENTS.md` §4 beschreibt nur real existierende Targets | `make gate-consistency` (in `gates`), Exit 0 |
| 5 | **Qualität** — Gates grün auf **frischem Klon** *und* in der CI mit demselben Image | CI-Run-Link + `make ci`-Ausgabe |
| 6 | **Distribution** — OCI-Label `…image.version` == Tag; `--print-mk` gibt den neuen Digest | Job-Summary der Release-Pipeline + `docker inspect`-Zeile |
| 7 | **Register** — Digest und Version in `version.md#aktuell`, CHANGELOG, beiden READMEs, `a-check.mk`, `cli.go` identisch | `make gates` **nach** dem Re-Pin, Exit 0 |
| 8 | **Betrieb** — die Incident-Klausel unten ist gelesen und gilt unverändert | Verweis auf diesen Abschnitt im Release-Eintrag |

**Freigabe-Eintrag:** Datum, Version, Digest und die acht Belege gehören in
[`version.md#verlauf`](../../version.md#verlauf) bzw. die GitHub-Release-Notes — nicht in eine
separate Datei, die niemand pflegt.

### Ausdrücklich *nicht* Teil der Freigabe (Anti-Items)

Ohne diese Liste wandern sie schleichend in die Pflicht:

- **Replay-Lauf gegen ein Golden Set.** a-check hat keinen; die Determinismus-Aussage tragen die
  Akzeptanztests ([SPEC-DET-001](../../spec/spezifikation.md#spec-det-001--determinismus-vertrag)).
- **100 % Coverage.** Die Schwelle ist 90 % und steht in
  [ADR-0006](../plan/adr/0006-coverage-gate.md); eine Anhebung wäre eine ADR, kein Release-Item.
- **Manuelle Smoke-Tests am Image.** `make image-test` deckt den Vertrag ab; was es nicht deckt,
  gehört als Test dorthin, nicht in eine Handprüfung vor dem Tag.
- **Konsumenten-Pins heben.** Das ist *nach* dem Release Sache der Konsumenten
  ([Konsum](#konsum-digest-pin)), kein Freigabe-Kriterium.

### Incident-Klausel: Rollback, Fix-Forward, Konsumenten-Pin

Wer nach dem Release einen Fehler findet, wählt **nicht im Vorfall** — die Regel steht hier:

| Wenn … | dann | warum |
|---|---|---|
| Der Fehler erzeugt **falsche Befunde** (False-Positive/False-Negative) in Konsumenten-CI | **Fix-Forward**: Patch-Release, danach Pin-Hebung ankündigen | ein „Rollback" beim Digest-Pin-Modell wirkt erst, wenn *jeder* Konsument selbst zurückdreht — das dauert länger als ein Patch |
| Das Image ist **gar nicht lauffähig** (Start bricht ab) | **Fix-Forward sofort**, zusätzlich `:latest` auf das letzte gute Release zurücksetzen | `:latest` ist die einzige Stelle, die a-check selbst zurückdrehen kann ([ADR-0007](../plan/adr/0007-latest-tag-politik.md)) |
| Der Fehler betrifft **nur** die Doku/das Register | kein neues Image; Register und CHANGELOG korrigieren | ein Release ohne Artefakt-Änderung erzeugt nur Pin-Arbeit ohne Gegenwert |
| Ein Konsument ist **blockiert** und der Fix braucht länger | Konsument dreht **seinen** Pin auf das letzte gute Digest zurück | die Pin-Hebung ist laut [ADR-0004](../plan/adr/0004-distribution-image-mk.md) bewusst pro Konsument — das gilt in beide Richtungen |

**Kein Rollback des GHCR-Tags `vX.Y.Z` selbst.** Ein veröffentlichter Versions-Tag wird nie
überschrieben: Konsumenten, die bereits gepinnt haben, würden sonst still ein anderes Artefakt
bekommen — genau das, was das Digest-Pin-Modell verhindern soll.

## Konsum (Digest-Pin)

Konsumenten pinnen auf den Digest, nicht auf bewegliche Tags
([AC-QA-03](../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit), Reproduzierbarkeit;
hermetisch/netzlos nach
[AC-QA-02](../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)):

```sh
docker run --rm --network none -v "$PWD:/src:ro" \
  ghcr.io/pt9912/a-check@sha256:<digest-aus-den-release-notes> /src
```

Die Pin-Hebung ist *manuell pro Konsument* — der akzeptierte Trade-off des
Pin-Modells ([ADR-0004](../plan/adr/0004-distribution-image-mk.md)): Digest
austauschen, Begründung in den Commit-Body. Ein zentral via `--print-mk`
verteiltes `a-check.mk` hält den Hebungs-Aufwand klein (eine Quelle statt N
Skript-Kopien).

## Aufruf-Referenz

Aufruf, Flags und Konfiguration: siehe [Benutzerhandbuch](benutzerhandbuch.md).
