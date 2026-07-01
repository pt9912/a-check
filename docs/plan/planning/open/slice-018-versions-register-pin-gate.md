# slice-018 — Versions-Register (`version.md`) + Pin-Gate

**Status:** open (Backlog — interne Release-Hygiene, nicht konsumenten-gated).
**Bezug:** schärft [AC-QA-03](../../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit)
(Reproduzierbarkeit/Digest-Pin) + [AC-FA-DIST-001](../../../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk)
(Distribution). [Roadmap](../in-progress/roadmap.md). **Evidenz:** ein **stale README-Pin**
(`v0.2.0` / Digest `13459f44…`) fiel am 2026-07-01 nur per Zufalls-Audit auf — kein Gate fing ihn.

> **Backlog-Stub.** Kein Entwurf zur Abnahme. Übernimmt das `version.md`-Register-Muster des
> Schwester-Tools d-check; wird zum Slice ausgearbeitet, sobald wir die manuellen Pin-Audits durch
> ein Gate ersetzen wollen (natürlicher Trigger: das nächste Release).

## 1. Auslöser

Versions-/Digest-Pins liegen an mehreren Stellen (`README.md`, `docs/user/benutzerhandbuch.md`,
`a-check.mk`, `internal/cli/cli.go`-`aCheckImage`, `CHANGELOG.md`) und driften **still**: der
README-Status stand auf `v0.2.0` mit einem veralteten Digest, während der Release längst `v0.3.0`
war — **kein Gate** hat das gemeldet. Das widerspricht dem „Invarianten als Gate statt Review-Meinung"-
Ethos. Das a-check nun gepinnte **d-check v0.35.0** bringt die Module `versions`/`pins` mit, die genau
diese Klasse erzwingen könnten (in `.d-check.yml` aktuell **nicht** aktiv).

Hinzu kommt `d-check.mk` (Pin des **Schwester-Tools**): heute eine Digest-Koordinate (driftfrei);
nach [slice-019](slice-019-dcheck-mk-print-mk-angleichung.md) **Tag + Digest**, die gegeneinander
driften können. Dessen **Tag↔Digest-Konsistenz** gehört dann in dieses Gate — **oder** wird explizit
exemptiert; hier zu entscheiden, damit slice-019 keine Drift-Quelle schafft, die das Gate nicht sieht.

## 2. Geplanter Umfang

1. **`version.md`** (Repo-Wurzel, Muster von d-check): Versions-Koordinaten (Tag + Datum + Release-Link)
   für `v0.1.0`–`v0.4.0`; **kein** Duplikat der CHANGELOG-Details. Anker `#vX.Y.Z` **nur** auf der
   aktuellen Version (wandert beim Release → stale Pins brechen als `anchor-missing`).
2. **Versions-Erwähnungen umhängen**: README-Status + Handbuch-Software-Version verlinken auf
   `version.md#aktuell` statt Nummer/Digest hart zu setzen.
3. **`versions`/`pins` in `.d-check.yml` scharfschalten** — mit `exempt-paths` für die absichtlich
   historischen Einträge (`CHANGELOG.md`, Handbuch-§10, `docs/reviews/**`), damit sie nicht anschlagen.
4. **Fitness-Function**: ein absichtlich stale gesetzter Pin ⇒ `make doc-check` rot.

## 3. Vor der Umsetzung zu klären

- **`versions`/`pins`-Modul-Vertrag von d-check lesen** (Spezifikation im d-check-Repo) — was genau
  prüfen sie (Versions-Nummern-Referenzen? Digest-Konsistenz? gegen welche Quelle der Wahrheit?), damit die
  `.d-check.yml`-Config stimmt und keine Falsch-Positiven in der Historie entstehen.
- **Digest-Quelle der Wahrheit**: `a-check.mk`-`A_CHECK_IMAGE` und `cli.go`-`aCheckImage` sind die Pins;
  soll das Gate auch deren **Gleichheit** erzwingen (heute manuell)?
- **Abgrenzung zu [slice-017](../done/slice-017-unbekannte-sprache-exit2.md)**: unabhängig (dort Sprach-Backends,
  hier Release-Pins) — nur beide „stille falsch-grün/veraltet"-Klassen.
