# Releasing — a-check

Release-Prozess für `ghcr.io/pt9912/a-check`
([AC-FA-DIST-001](../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk),
[ADR-0004](../plan/adr/0004-distribution-image-mk.md)). Das Distributionsmodell
ist verbindlich; die **automatisierte** Pipeline ist dagegen das Zielbild von
`welle-05-release` und existiert noch **nicht** — siehe
[Roadmap](../plan/planning/in-progress/roadmap.md). Diese Datei beschreibt den
Prozess so, wie er heute (vor dem ersten Release) manuell läuft, und markiert,
was die Release-Welle automatisiert.

## Stand heute (vor dem ersten Release)

Es gibt **kein getaggtes GHCR-Release** und **keine** `.github/`-Pipeline. Wer
heute baut/prüft, nutzt das lokal gebaute Image:

```sh
make build                       # baut a-check:dev (static/distroless, digest-gepinnte Bases)
make a-check A_CHECK_IMAGE=a-check:dev   # Konsum-Aufruf gegen das lokale Image
```

Das `a-check`-Target und das einbindbare Fragment liefert
`a-check --print-mk` ([AC-FA-DIST-001](../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk)),
siehe [`a-check.mk`](../../a-check.mk) dieses Repos. Solange kein GHCR-Image
veröffentlicht ist, wird `A_CHECK_IMAGE` beim Aufruf überschrieben
([Benutzerhandbuch §3.3](benutzerhandbuch.md#33-a-check-als-make--oder-ci-gate-einbinden)).

## Versionsquelle

Versionen folgen SemVer; die menschlich kuratierte Begründung jedes Releases ist
der zugehörige Abschnitt in [`CHANGELOG.md`](../../CHANGELOG.md). Vor dem Tag
wird dort der `[Unreleased]`-Stand unter die neue Version geschnitten. Das
Lastenheft steht bei 0.1.0.

## Release auslösen (Zielbild `welle-05-release`)

Bis die Pipeline (`welle-05`) steht, ist der Release manuell; die Schritte sind
dieselben, die die Welle automatisieren wird:

```sh
git tag v0.1.0
git push origin v0.1.0
```

1. **SemVer-Validate** (fail-fast): nur `vMAJOR.MINOR.PATCH` oder
   `…-PRERELEASE`; Build-Metadaten (`+`) werden abgelehnt.
2. **`make gates`** — alle inneren Gates grün (lint/test/coverage-gate/
   arch-check/doc-check/gate-consistency); Runtime-Image über `make build` mit
   `VERSION` aus dem Tag.
3. **OCI-Label-Pin** — `org.opencontainers.image.version` muss exakt der
   Tag-Version entsprechen (Version-Drift shippt nicht).
4. **Push** nach `ghcr.io/pt9912/a-check:v<version>`. Ob `:latest` für stabile
   Releases (ohne Prerelease-Suffix) mitgeschoben wird, entscheidet `welle-05`
   per ADR (bewegliche Tags sind für CI-Pins ungeeignet).
5. **Digest-Pin** in die GitHub-Release-Notes; danach gibt
   `a-check --print-mk` ein `a-check.mk` mit dem **aktuell digest-gepinnten**
   `A_CHECK_IMAGE` aus
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
