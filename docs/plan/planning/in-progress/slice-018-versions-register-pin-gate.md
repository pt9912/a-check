# slice-018 — Versions-Register (`version.md`) + Pin-Gate

**Status:** open (**Entwurf zur Abnahme** — Modell durchdacht; Digest-vs-Tag-Pin-Frage
geklärt, Optionen + Empfehlung unten). **Typ:** interne Release-Hygiene, nicht konsumenten-gated.
**Bezug:** schärft [AC-QA-03](../../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit)
(Reproduzierbarkeit/Digest-Pin) + [AC-FA-DIST-001](../../../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk)
(Distribution). [Roadmap](../in-progress/roadmap.md). **Evidenz:** ein **stale README-Pin**
(`v0.2.0` / Digest `13459f44…`) fiel am 2026-07-01 nur per Zufalls-Audit auf — kein Gate fing ihn.

## 1. Auslöser

Versions-/Digest-Pins liegen an mehreren Stellen (`README.md`, `docs/user/benutzerhandbuch.md`,
`docs/user/releasing.md`, `a-check.mk`, `internal/cli/cli.go`-`aCheckImage`, `CHANGELOG.md`) und
driften **still**: der README-Status stand auf `v0.2.0` mit veraltetem Digest, während der Release
längst weiter war — **kein Gate** hat das gemeldet. Das widerspricht dem „Invarianten als Gate statt
Review-Meinung"-Ethos. Das seit [slice-019](../done/slice-019-dcheck-mk-print-mk-angleichung.md)
gepinnte **d-check v0.37.1** bringt die Module `versions`/`pins` mit (in `.d-check.yml` **nicht** aktiv).

Hinzu kommt `d-check.mk` (Pin des **Schwester-Tools**): seit slice-019 **Tag + Digest**
(`:v0.37.1` / `@sha256:3bbdb19b…`), die gegeneinander driften können. Dessen **Tag↔Digest-Konsistenz**
gehört in dieses Gate — **oder** wird explizit exemptiert; hier zu entscheiden, damit slice-019 keine
Drift-Quelle schafft, die das Gate nicht sieht.

## 2. Kern-Erkenntnis: d-checks Module passen nicht 1:1 (Digest vs. Tag)

Der Modul-Vertrag (d-check-Handbuch, gelesen 2026-07-04):

- **`versions`** matcht **Versions-Tag-Pins** (`ghcr…:vX.Y.Z`) und prüft die Version gegen
  `version.md#aktuell`. **a-check pinnt aber per Digest** (`@sha256:…`, [AC-QA-03](../../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit)) — die
  Versions-*Nummer* steht nur in **Prosa** (README-Status, Handbuch-Software-Version, releasing.md).
  `versions` greift also **nicht** auf a-checks `@sha256:`-Pins.
- **`pins`** prüft **Link-Content-Drift** via `<!-- dpin: sha256:… -->`-Markern — **nicht**
  Digest-Pin-Aktualität. Passt ebenfalls nicht auf „alle `@sha256:`-Pins == aktueller Release-Digest".

**Zwei zu fangende Drift-Klassen:**
- **(A) Versions-Nummer-Drift** — Prosa `vX.Y.Z` in README/Handbuch/releasing.md ≠ aktuelle Version.
- **(B) Digest-Drift** — `@sha256:…` in `a-check.mk`/`cli.go`/Prosa ≠ aktueller Release-Digest;
  und `a-check.mk`-Digest == `cli.go`-Digest (die zwei autoritativen harten Pins).

## 3. Optionen (zur Abnahme)

| Opt | Idee | Fängt | Kosten |
|---|---|---|---|
| **1 — Prosa→`version.md`-Link** | README-Status/Handbuch-Version/releasing.md verlinken auf `version.md#aktuell` statt Nummer+Digest hart zu setzen. **Eine Quelle** ⇒ strukturell drift-frei. | (A) + Prosa-Hälfte von (B) | Doku-Umbau; `links`/`anchors` (aktiv) fangen einen toten `#aktuell`-Anker automatisch |
| **2 — `versions`-Modul auf Prosa** | `pin-pattern` an a-checks Prosa-Versions-Muster anpassen, `current-from: version.md#aktuell`. | (A) | zusätzliche `.d-check.yml`-Config; deckt (B) NICHT (digest-blind) |
| **3 — Digest-Gleichheits-Check** | Kleiner Check: `a-check.mk`-Digest == `cli.go`-Digest == `version.md`-Digest. | (B) für die harten Pins | **neues Skript** — Spannung zum „Skript-Copies verringern"-Ziel [[a-check-skript-reduktions-ziel]] |

**Empfehlung:** **Opt 1 + 3**. Opt 1 tilgt (A) und die Prosa-(B) *strukturell* (die einzige verbleibende
Wahrheit ist `version.md`; tote Anker fängt das aktive `anchors`-Modul). Opt 3 deckt die zwei **harten**
Pins, die nicht verlinken können (`a-check.mk` ist ein an Konsumenten geliefertes Artefakt, `cli.go`
bettet den Digest ein). Opt 2 (`versions`-Modul) ist redundant zu Opt 1 und digest-blind — **nicht**
adoptieren. Für Opt 3 die reduce-copies-Spannung abwägen: entweder ein schlankes `tools/pin-check.sh`
(wie `gate-consistency.sh`, mit Selbsttest) **oder** die Gleichheit in `gate-consistency.sh` mitprüfen
(kein neues Skript).

## 4. Geplanter Umfang (nach Abnahme der Option)

1. **`version.md`** (Repo-Wurzel, Muster d-check): aktuelle Versions-Koordinate (Version + Datum +
   Digest + Release-Link) unter Anker `#aktuell`; Historie knapp (kein CHANGELOG-Duplikat).
2. **Prosa umhängen** (Opt 1): README-Status, Handbuch-Software-Version, releasing.md „Aktuelles
   Release" → Link auf `version.md#aktuell`.
3. **Digest-Gleichheits-Check** (Opt 3): `a-check.mk` == `cli.go` == `version.md`; in
   `gate-consistency.sh` oder eigenem `pin-check`. In `make gates` aufnehmen.
4. **`d-check.mk`-Tag↔Digest**: mitprüfen (Tag `:vX` und `@sha256:` konsistent) **oder** in §3 bewusst
   exemptieren (dokumentiert).
5. **Release-Prozess** (`releasing.md`): der Re-Pin schreibt künftig `version.md#aktuell` (eine Stelle)
   statt der acht heute manuell nachgezogenen; der Check fängt eine vergessene.
6. **Fitness-Function**: ein absichtlich stale gesetzter Pin/eine falsche Version ⇒ `make gates` rot.

## 5. Vor der Umsetzung noch zu klären

- **Opt-Wahl** (§3) — Empfehlung 1+3; Maintainer-Entscheid.
- **reduce-copies vs. eigenes Skript** (Opt 3): `gate-consistency.sh`-Erweiterung (kein neues Skript)
  bevorzugt?
- **`d-check.mk`-Tag↔Digest**: prüfen oder exemptieren?
- **Abgrenzung**: der bei der Hygiene-Bereinigung notierte Zusatzfall „gedrucktes `--print-mk`-Fragment
  ≠ committete `a-check.mk`" (Wortlaut-Drift) — hier mitnehmen oder separat? (Eher separat: das ist
  Fragment-Text-Parität, keine Pin-Drift.)
