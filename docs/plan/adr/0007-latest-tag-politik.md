# ADR-0007 — `:latest`-Tag-Politik für stabile Releases

- **Status:** Accepted
- **Datum:** 2026-06-21
- **Autor:** pt9912
- **Bezug:** [AC-FA-DIST-001](../../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk) (GHCR-Image/Tagging), [AC-QA-03](../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit) (Digest-Pin-Reproduzierbarkeit), [AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)
- **Schärft:** [SPEC-DIST-001](../../../spec/spezifikation.md#spec-dist-001--laufzeitform-und-distribution) — macht die Tagging-Politik (`:latest` nur stabil) und die Digest-Pin-Pflicht der Konsumenten verbindlich für [`release.yml`](../../../.github/workflows/release.yml) und [`releasing.md`](../../../docs/user/releasing.md).
- **Supersedes:** —

## Kontext

Die Release-Pipeline (slice-007, [`release.yml`](../../../.github/workflows/release.yml))
pusht das GHCR-Image auf `v*`-Tags. Offen ist die Tagging-Politik: bewegt ein
beweglicher `:latest`-Tag mit? Anders als beim Schwester-Tool `d-check` (dessen
Distributions-ADR ursprünglich „kein `latest`" festschrieb und später ratifiziert
wurde) trifft a-checks [ADR-0004](0004-distribution-image-mk.md) **keine**
Tag-Aussage — dies ist also eine **frische** Entscheidung, kein Supersede.

Seit [ADR-0004](0004-distribution-image-mk.md) trägt der `@sha256:`-Digest die
Reproduzierbarkeit ([AC-QA-03](../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit)):
solange Konsumenten verbindlich per Digest pinnen, gefährdet ein beweglicher
`:latest`-Tag die Reproduzierbarkeit nicht — er ist reiner Komfort-Einstieg.

## Optionen

1. **`:latest` für stabile Releases (gewählt).** `vX.Y.Z` ohne Prerelease-Suffix
   pusht zusätzlich `:latest`; Konsumenten pinnen verbindlich per Digest.
   Trade-off: ein beweglicher Tag existiert — für CI-Pins ungeeignet, dokumentiert.
2. **Kein `:latest`.** Strikt nur volle Semver-Tags. Pro: kein beweglicher Tag;
   Contra: kein Komfort-Einstieg. Verworfen — der Digest trägt die Reproduzierbarkeit ohnehin.
3. **Zusätzlich bewegliche Major-/Minor-Tags** (`:v0`, `:v0.1`). Pro: feinere
   Komfort-Granularität; Contra: größere Drift-Fläche, nicht gefordert. Verworfen.

## Entscheidung

1. **`:latest` nur für stabile Releases.** Für Tags ohne Prerelease-Suffix
   (`vMAJOR.MINOR.PATCH`) wird `:latest` gesetzt und gepusht und zeigt stets auf
   das neueste stabile Release (`release.yml` `IS_STABLE`-Verzweigung).
2. **Prereleases erhalten kein `:latest`** (`vX.Y.Z-rc.1` o. Ä.: nur der
   Versions-Tag).
3. **Verbindlicher Konsum per Digest.** Konsumenten pinnen auf `@sha256:`-Digest
   ([AC-QA-03](../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit),
   [ADR-0004](0004-distribution-image-mk.md)); `:latest` ist Komfort-Einstieg,
   **nicht** für CI-Pipelines.
4. **Keine beweglichen Major-/Minor-Tags** — nur `:latest` plus volle Semver-Tags.

## Konsequenzen

- [`release.yml`](../../../.github/workflows/release.yml) setzt/pusht `:latest`
  nur bei `IS_STABLE=true`; der OCI-Label-Verify erzwingt
  `org.opencontainers.image.version` == Tag-Version (Version-Drift shippt nicht).
- [`releasing.md`](../../../docs/user/releasing.md) dokumentiert die
  Digest-Pin-Pflicht und verweist auf diese ADR.
- `:latest` ist bewusst **nicht** reproduzierbar (beweglich) — die
  Reproduzierbarkeit liegt am Digest
  ([AC-QA-03](../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit)).

## Fitness Function

- [`release.yml`](../../../.github/workflows/release.yml) trägt die
  SemVer-Validate-Stufe (fail-fast) + `IS_STABLE`-Verzweigung + OCI-Label-Verify;
  `make doc-check` hält die ADR-/Doku-Verweise konsistent.
- Kein `make`-Gate prüft die CI-YAML selbst — die strukturelle Kontrolle ist die
  SemVer-Validate-Stufe plus die `IS_STABLE`-Verzweigung.

## Re-Evaluierungs-Trigger

- Bedarf an beweglichen Major-/Minor-Tags → neue ADR.
- Wechsel der Registry- oder Tagging-Mechanik → Pin-Strategie erneut prüfen.

## Geschichte

| Datum | Ereignis |
|---|---|
| 2026-06-21 | Proposed — Release-Pipeline (slice-007); frische Tagging-Politik (a-checks ADR-0004 trifft keine Tag-Aussage). |
| 2026-06-21 | Proposed → Accepted (Sign-off Auftraggeber). Ab jetzt immutable; Ablösung nur via Folge-ADR mit `Supersedes`. |
