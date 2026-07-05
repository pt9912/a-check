# slice-018 — Versions-Register (`version.md`) + Pin-Gate

**Status:** in-progress (**in Umsetzung** — Opt 1 + 3 abgenommen 2026-07-05, Modul-Vertrag
gegen d-check-Handbuch `v0.37.1`/1.21 re-verifiziert; Entscheide in §6, Umsetzung in §7).
**Typ:** interne Release-Hygiene, nicht konsumenten-gated.
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

## 5. Abnahme-Entscheide (2026-07-05, Maintainer)

Vorab re-verifiziert gegen das aktuelle d-check-Handbuch (`main` == `v0.37.1`, Handbuch 1.21):
`versions` matcht **nur** `ghcr.io/…:vX.Y.Z`-Tag-URLs gegen `version.md#aktuell`; `pins` prüft
`<!-- dpin: sha256:… -->`-Content-Drift; **kein** Modul prüft Digest-Gleichheit über Dateien
oder Tag↔Digest-Konsistenz. Rest-Audit: die **einzige** Tag-URL im Repo ist `d-check.mk`s
`…/d-check:v0.37.1` — `versions` würde a-checks `version.md#aktuell` (v0.10.0) dagegen als
`version-stale` **fehlzünden**. ⇒ Kern-Erkenntnis §2 bestätigt und verschärft.

- **Opt-Wahl** (§3): **Opt 1 + 3** abgenommen. Opt 2 (`versions`-Modul) verworfen (redundant,
  digest-blind, würde gegen die d-check-Tag-URL fehlzünden).
- **reduce-copies** (Opt 3): **`gate-consistency.sh` erweitern** (kein neues Skript) —
  [[a-check-skript-reduktions-ziel]].
- **`d-check.mk`-Tag↔Digest**: **mitprüfen** — offline-belegbare Deklarations-Konsistenz
  (wohlgeformter Tag aus der `DCHECK_IMAGE`-Zeile + wohlgeformter `DCHECK_DIGEST`; **nur die
  tragenden Zeilen**, damit legitime Kommentar-Versionsnummern keinen Fehlalarm auslösen —
  Review-Befund A2). **Ehrliche Grenze
  ([AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)):** dass `:vX.Y.Z` tatsächlich auf `@sha256:…` auflöst, ist eine Registry-/Netz-
  Eigenschaft — netzlos nicht prüfbar; sie wird beim Re-Pin (online) verifiziert.
- **print-mk-Fragment-Parität**: **separat** (eigener Slice) — das ist Fragment-Text-Parität,
  keine Pin-Drift.
- **Scope-Nachtrag**: der Kommando-Digest in `README.md` (`docker run`) kann nicht verlinken
  ⇒ in die Opt-3-Gleichheitsmenge aufgenommen; der Live-Digest/Version-Verweis in
  `roadmap.md` auf `version.md#aktuell` umgehängt (sonst ungegatete Zweitkopie).

## 6. Umsetzung (slice-018)

1. **`version.md`** (Repo-Wurzel, d-check-Muster + a-check-Digest-Adaption): `## Aktuell` mit
   Version + Datum + **vollem** `@sha256:`-Digest unter `#aktuell`/wanderndem `#vX.Y.Z`-Anker;
   `## Verlauf` (Version/Datum/Release-Link, kein CHANGELOG-Duplikat).
2. **Pin-Konsistenz** in `tools/gate-consistency.sh` (Check (4), Selbsttest `pin_self_test`):
   - **(B) Digest-Gleichheit:** `a-check.mk` == `internal/cli/cli.go` == `README.md`-Kommando
     == `version.md#aktuell` (voller 64-hex).
   - **(A) Versions-Nummer:** `version.md#aktuell` == aktuellstes `CHANGELOG.md`-Release.
   - **(C) d-check.mk:** wohlgeformter Tag (aus der `DCHECK_IMAGE`-Zeile) + wohlgeformter `DCHECK_DIGEST`.
   - **Eindeutigkeit:** jede harte Pin-Datei trägt genau **einen** a-check-Digest; ein zweiter,
     abweichender (Decoy) ⇒ fail-closed (Review-Befund A1).
   - **Fitness-Function:** `pin_self_test` beweist offline für **alle drei** Dimensionen (B/A/C)
     **und** den Decoy-Fall, dass die Drift das Gate rot macht.
3. **Prosa umgehängt (Opt 1):** README-Status/Tags, `releasing.md` „Aktuelles Release"/
   „Aktueller Stand", `benutzerhandbuch.md`-Software-Version → Link auf `version.md#aktuell`;
   `roadmap.md`-Live-Verweis ebenso. Historische Vorkommen (CHANGELOG, Handbuch-Historie)
   bleiben literal. `git tag`-Beispiel als Platzhalter `vX.Y.Z`.
4. **Release-Prozess** (`releasing.md` §5, neuer Schritt 6): der Re-Pin schreibt `version.md#aktuell`
   (eine Stelle) + die harten Pins; das Gate fängt eine vergessene.
5. **Contract-Nachzug:** `gate-consistency`-Beschreibung in `AGENTS.md` §4 + `harness/README.md`
   §Sensors um die Pin-Konsistenz erweitert.

## 7. Offene Grenze / Folge

- **Tag↔Digest-Auflösung** von `d-check.mk` (und der Prosa-Versions-Nummer gegen den Registry-
  Digest) ist netzlos nicht beweisbar (§5) — dokumentierte Heuristik-Grenze, verifiziert beim Re-Pin.
- **print-mk-Fragment-Parität** (gedrucktes `--print-mk` == committete `a-check.mk`) als eigener
  Slice vorgemerkt.
- **Kommentar-Versionsnummern** (`internal/cli/cli.go`-Kopfkommentar **und** `a-check.mk`-
  Kopfkommentar, je `v0.10.0` in Prosa) bleiben ungegatet (beim Re-Pin nachgezogen) — bewusst,
  um brüchiges Kommentar-Parsing zu vermeiden; der reproduzierbarkeits-relevante **Digest** ist
  gegatet. Der `a-check.mk`-Kommentar-Wortlaut fällt zudem unter die separate print-mk-Fragment-
  Parität (der Generator `mkFragment` führt keine Versions-Nummer). Beide sind die einzigen
  verbliebenen ungegateten Live-Versions-Literale (Review-Befund C2/C3, bewusst abgegrenzt).

## 8. Review-Härtung (2026-07-05, adversarisches Multi-Linsen-Review)

Drei unabhängige Linsen (Bash-Korrektheit / Spec-Konformität / Doku-Drift). Linse B: konform.
Behobene Befunde:

- **A1 (kritisch):** Digest-Extraktion war unverankert (`head -1`) — ein Decoy-Zweitdigest hätte
  echte Drift maskiert. Fix: jede harte Pin-Datei muss genau **einen** a-check-Digest tragen
  (`sort -u`, Zähl-Guard), version.md-Digest auf den `## Aktuell`-Abschnitt verankert.
- **A2:** d-check-Versions-Zählung über die ganze Datei ⇒ Fehlalarm bei legitimem Kommentar.
  Fix: Tag nur aus der `DCHECK_IMAGE`-Zeile.
- **A3:** `set -e` + fehlschlagende Command-Substitution — Extraktoren mit `|| true` abgesichert.
- **A5:** `pin_self_test` deckte nur (B) ab — jetzt (B/A/C) + Decoy-Regression.
- **C1:** `harness/README.md`-„Aktueller Lauf-Status" trug ein ungegatetes `v0.10.0` — entfernt.
- **C2/C3:** `a-check.mk`- und `cli.go`-Kommentar-Versionen in §7 vollständig verbucht.
