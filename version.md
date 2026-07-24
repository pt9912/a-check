# a-check — Release-Register

> Kanonischer, **auflösender** Link-Ziel-Ort für Erwähnungen **eigener**
> a-check-Releases — etwa [die jeweils aktuelle Version](#aktuell). a-check pinnt
> per **Digest** (`@sha256:…`,
> [AC-QA-03](spec/lastenheft.md#ac-qa-03--reproduzierbarkeit)), nicht per
> beweglichem Tag; darum trägt die Aktuell-Zeile zusätzlich den **vollen
> Release-Digest** — die **eine Wahrheit**, gegen die
> [`tools/gate-consistency.sh`](tools/gate-consistency.sh) die harten Pins
> ([`a-check.mk`](a-check.mk), [`internal/cli/cli.go`](internal/cli/cli.go) und die
> `docker run`-Beispiele in [`README.md`](README.md) + [`README.de.md`](README.de.md)) auf Gleichheit prüft
> (slice-018). **Kein Duplikat** der Detail-Changes — die stehen im
> [CHANGELOG](CHANGELOG.md); hier nur Versions-Koordinaten (Version, Datum,
> Digest, Tag). **Fremde** Versionen (d-check, Baseline) gehören nicht hierher,
> sondern verlinken auf ihre eigene Quelle.

## Aktuell

Aktuelle Version: [`v0.15.0`](#v0.15.0) — 2026-07-24.
Release-Digest: `ghcr.io/pt9912/a-check@sha256:6425c93a9a4359ef28c4da231a2d1db6f421fdaa8f96877ac89d201827c42d09`.

Aus anderen Dokumenten stabil referenzierbar als `version.md#aktuell` (zeigt immer
hierher, nie auf eine feste Nummer). Pro Release sind genau diese Zeile **und** eine
neue [Verlaufs](#verlauf)-Zeile nachzuziehen **und der `<a id>`-Anker auf die neue
Version zu verschieben** (die bisherige Zeile verliert ihn — sonst bleiben veraltete
feste Pins `version.md#vX.Y.Z` auflösbar und die Anker-Kaskade meldet den vergessenen
Bump nicht). Zusammen mit dem Digest-Gleichheits-Gate ist der Bump damit einpunktig:
`version.md` schreiben, das Gate fängt eine vergessene harte Pin-Stelle.

## Verlauf

| Version                       | Datum      | Release                                                             |
| ----------------------------- | ---------- | ------------------------------------------------------------------ |
| `v0.15.0` <a id="v0.15.0"></a>| 2026-07-24 | [Tag v0.15.0](https://github.com/pt9912/a-check/releases/tag/v0.15.0) |
| `v0.14.0`                     | 2026-07-23 | [Tag v0.14.0](https://github.com/pt9912/a-check/releases/tag/v0.14.0) |
| `v0.13.0`                     | 2026-07-09 | [Tag v0.13.0](https://github.com/pt9912/a-check/releases/tag/v0.13.0) |
| `v0.12.0`                     | 2026-07-06 | [Tag v0.12.0](https://github.com/pt9912/a-check/releases/tag/v0.12.0) |
| `v0.11.0`                     | 2026-07-05 | [Tag v0.11.0](https://github.com/pt9912/a-check/releases/tag/v0.11.0) |
| `v0.10.0`                     | 2026-07-04 | [Tag v0.10.0](https://github.com/pt9912/a-check/releases/tag/v0.10.0) |
| `v0.9.0`                      | 2026-07-04 | [Tag v0.9.0](https://github.com/pt9912/a-check/releases/tag/v0.9.0) |
| `v0.8.0`                      | 2026-07-03 | [Tag v0.8.0](https://github.com/pt9912/a-check/releases/tag/v0.8.0) |
| `v0.7.0`                      | 2026-07-03 | [Tag v0.7.0](https://github.com/pt9912/a-check/releases/tag/v0.7.0) |
| `v0.6.0`                      | 2026-07-02 | [Tag v0.6.0](https://github.com/pt9912/a-check/releases/tag/v0.6.0) |
| `v0.5.0`                      | 2026-07-02 | [Tag v0.5.0](https://github.com/pt9912/a-check/releases/tag/v0.5.0) |
| `v0.4.0`                      | 2026-07-01 | [Tag v0.4.0](https://github.com/pt9912/a-check/releases/tag/v0.4.0) |
| `v0.3.0`                      | 2026-06-23 | [Tag v0.3.0](https://github.com/pt9912/a-check/releases/tag/v0.3.0) |
| `v0.2.0`                      | 2026-06-22 | [Tag v0.2.0](https://github.com/pt9912/a-check/releases/tag/v0.2.0) |
| `v0.1.0`                      | 2026-06-21 | [Tag v0.1.0](https://github.com/pt9912/a-check/releases/tag/v0.1.0) |
