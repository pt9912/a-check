# Releasing — a-check

Release-Prozess für `ghcr.io/pt9912/a-check`
([AC-FA-DIST-001](../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk),
[ADR-0004](../plan/adr/0004-distribution-image-mk.md),
[ADR-0007](../plan/adr/0007-latest-tag-politik.md)). Die Pipeline ist
[`.github/workflows/release.yml`](../../.github/workflows/release.yml) (seit
slice-007). **Aktuelles Release: `v0.4.0`** (Vorgänger `v0.3.0`, erstes `v0.1.0`) —
[GitHub-Release](https://github.com/pt9912/a-check/releases/tag/v0.4.0),
GHCR-Tags `v0.4.0` + `latest`.

## Aktueller Stand

`v0.4.0` ist auf GHCR verfügbar; Konsumenten pinnen den **Digest**
([Konsum](#konsum-digest-pin)). Das mitgelieferte
[`a-check.mk`](../../a-check.mk) und `a-check --print-mk`
([AC-FA-DIST-001](../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk))
sind auf `@sha256:b0d6e33c…` digest-gepinnt. Für lokale Entwicklung gegen
ungetaggte Stände dient weiterhin das lokal gebaute Image:

```sh
make build                               # baut a-check:dev (static/distroless)
make a-check A_CHECK_IMAGE=a-check:dev   # Konsum-Aufruf gegen das lokale Image
```

## Versionsquelle

Versionen folgen SemVer; die menschlich kuratierte Begründung jedes Releases ist
der zugehörige Abschnitt in [`CHANGELOG.md`](../../CHANGELOG.md). Vor dem Tag
wird dort der `[Unreleased]`-Stand unter die neue Version geschnitten. Das
Lastenheft steht bei 0.8.0.

## Release auslösen

```sh
git tag v0.4.0
git push origin v0.4.0
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
