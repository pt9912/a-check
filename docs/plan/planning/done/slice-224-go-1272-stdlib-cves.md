# slice-224 — Go-Toolchain auf 1.27.2: behebbare HIGH-CVEs der Standardbibliothek

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `git mv` (`make slice-mv`).

**Welle:** ohne Welle — die Closure-Bedingung wäre die eigene DoD.

**Bezug:** [ADR-0037](../../adr/0037-cve-scan-gegen-das-publizierte-image.md) (CVE-Scan; über rot
entscheiden behebbare CRITICAL/HIGH), [ADR-0038](../../adr/0038-dependabot-als-hebungskanal.md)
(Hebung als bewusster Commit), [AC-QA-03](../../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit)
(digest-gepinnte Basis).

**Berührte Spec-Stellen:** — · Der Slice berührt kein Spec-Stratum.

**Verantwortlich:** Claude

**Autor:** Claude. **Datum:** 2026-10-10.

**Lerneintrag — Form:** geschärfte Regel.

---

## 1. Ziel und Abgrenzung

**Ziel:** Das Release-Image wird mit Go **1.27.2** gebaut; die behebbaren HIGH-Befunde aus dem
Nachtlauf `image-scan` vom 2026-10-09 (Lauf 37918923676) entfallen:
`CVE-2026-78667` (`net/http`, DoS) und `CVE-2026-97031` (`crypto/tls`, DoS) — der lokale Lauf am
2026-10-10 zeigt dazu `CVE-2026-78669` —, alle in `stdlib`
`v1.27.0`, behoben in 1.26.9 / 1.27.2. Gebaut wird beide Module mit Go 1.27.2 (Maintainer:
„Go ≥ 1.27.2").

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Das Release.** *Ein anderer Vorgang*: der Fix wirkt beim Konsumenten erst mit einem neuen
  Release; nach der Incident-Klausel in `releasing.md` ist das ein Fix-Forward-Patch-Release
  (`v0.23.2`) und folgt unmittelbar auf diesen Slice (Maintainer-Wort 2026-10-10: „Bitte auch ein
  Release erstellen").
- **Die `go`-Direktive in den `go.mod`.** *Bestand bleibt bewusst stehen* (Plan-Änderung vor dem
  Code-Commit, 2026-10-10): sie nennt die Mindest-Sprachversion, nicht die Build-Toolchain — die
  setzt das digest-gepinnte Basis-Image. Gemessen: mit `go 1.27.2` lädt der Lint-Container
  (golangci-lint `v2.13.2` mit Go 1.27.0; auch `v2.14.0` bringt nur 1.27.1) eine Toolchain aus dem
  Netz nach, und die Typprüfung bricht (`export data version 5 is greater than maximum supported
  version 4`). Für den CVE-Fix zählt allein, womit das ausgelieferte Binary gebaut wird.
- **`golangci-lint`-Image.** *Bestand bleibt bewusst stehen*: es baut nicht das ausgelieferte
  Binary; der Scan betrifft das Release-Bild.
- **Warum Dependabot den Hub nicht gemeldet hat.** *Ein anderer Vorgang*: nach dem Release prüfen,
  ob der Kanal aus [ADR-0038](../../adr/0038-dependabot-als-hebungskanal.md) die Basis-Hebung hätte vorschlagen müssen.

## 2. Definition of Done

- [x] `golang`-Basis auf `1.27.2` mit Index-Digest in `Dockerfile` und `tools/archive-wave/Dockerfile`;
      `GO_VERSION` in beiden Makefiles.
- [x] Beleg vor dem Release: Trivy gegen das lokal gebaute Multi-Arch-Archiv (`--input`) zeigt
      die CVEs nicht mehr, je Plattform; Gegenprobe: das publizierte, mit 1.27.0 gebaute Image zeigt
      sie.
- [x] CHANGELOG `[Unreleased]`.
- [x] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [x] Beobachtungs-Register fortgeschrieben (oder „keine Beobachtung" notiert).
- [x] Jedes Risiko aus §6 trägt einen Ausgang.

`make ci`, `make archive-wave-test`, `make gates` und `make verify` grün. Ein unabhängiger Review
ist für diesen Slice nicht vorgesehen ([MR-019](../../../../harness/conventions.md#mr-019): Opt-in) —
eine Versions-Hebung mit Messung vorher und nachher.

## 3. Plan (vor Code)

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| `Dockerfile`, `Makefile` | update | Toolchain 1.27.2 |
| `tools/archive-wave/Dockerfile`, `tools/archive-wave/Makefile` | update | dasselbe für das Werkzeug |
| `CHANGELOG.md` | update | Sicherheits-Hebung, für Konsumenten sichtbar |

## 4. Trigger

**Start** (`open` → `in-progress`): Maintainer-Wort („Go ≥ 1.27.2 — siehe auch github ci",
2026-10-10), WIP-Limit frei.

**Rückführungen:** `in-progress` → `open` (blockiert): 1.27.2 bricht einen Gate-Lauf, der nicht in
einer Sitzung zu beheben ist.

## 5. Closure-Trigger

DoD vollständig, Gates grün, Closure-Notiz mit Lerneintrag.

## 6. Risiken und offene Punkte

- **Der Fix wirkt erst mit dem Release:** bis dahin bleibt der Nachtlauf rot. — **Ausgang:** *entfallen* — gestrichen mit Begründung:
  der Abstand ist kein offenes Risiko, sondern der bekannte Weg; das Patch-Release `v0.23.2` ist vom
  Maintainer beauftragt (2026-10-10) und folgt unmittelbar, der erneute `image-scan`-Lauf gegen das
  publizierte Image ist sein Beleg.

## 7. Closure-Notiz

**Lerneintrag — Form: geschärfte Regel.** Ein CVE-Scan belegt eine Plattform erst, wenn sein Bericht
deren Architektur nennt — auch außerhalb von `make image-scan`. Der Ad-hoc-Scan gegen das lokale
Multi-Arch-Archiv fiel mit `--platform linux/arm64` still auf amd64 zurück; erst ein Layout mit nur
dem einen Manifest lieferte `arm64`. Dieselbe Form wie in slice-221 auf einem zweiten Eingangsweg:
gezählt, nicht verkörpert (2×).

**Geliefert:** `golang:1.27.2` per Index-Digest (`sha256:5bc7f572…`) im Release- und im
`archive-wave`-Dockerfile, `GO_VERSION` in beiden Makefiles, CHANGELOG `[Unreleased]` (Security).

**Messung:** Trivy `0.74.0` gegen das lokal gebaute Multi-Arch-Archiv, je Plattform mit
Architektur-Nachweis: `amd64` und `arm64` ohne `stdlib`-Befund (verbleibend je ein UNKNOWN in
`tzdata` der Basis, nicht handlungspflichtig). Gegenprobe `make image-scan` gegen das publizierte
`v0.23.1`: Exit 2, je Plattform drei behebbare HIGH — `CVE-2026-78667`, `CVE-2026-78669`,
`CVE-2026-97031`, installiert `v1.27.0`, behoben in 1.27.2.

**Was hat funktioniert:** der Nachtlauf als Kanal. Die offene Frage aus §1, warum Dependabot keine
Hebung vorschlug, beantwortet `.github/dependabot.yml` selbst: das Ökosystem `docker` ist bewusst
nicht aufgenommen, eine Basis-Hebung ist ein bewusster Commit. Der Scan hat genau die Lücke
getragen, für die er da ist.

**Was ging anders als geplant:** Die `go`-Direktive auf `1.27.2` brach `make lint` — die
golangci-lint-Images bringen 1.27.0 bzw. 1.27.1 mit und laden sonst eine Toolchain aus dem Netz
nach. Plan-Änderung vor dem Code-Commit: die Direktive bleibt (§1). Und der Nachtlauf nannte zwei
HIGH, der lokale Lauf drei (`CVE-2026-78669` dazu) — alle drei mit derselben Hebung behoben.

**Steering-Loop-Eintrag:** geschärfte Regel — gezählt, nicht verkörpert.

**Beobachtungs-Register (`../observations/`):** neu
`BEO-GATE/scanner-faellt-still-auf-andere-plattform-zurueck` (2×: slice-221, slice-224).

**Folge-Slices:** keine.

**Risiken aus §6:** das eine Risiko trägt seinen Ausgang (*entfallen*, mit Begründung).

**Drei Paarungen:** Anker — kein `liegt in`-Feld · Folge-Slice — keiner · Register — der genannte
Pfad existiert mit nicht leerem `evidence/`.

**Trigger-Audit der aktiven MR:** [`MR-016`](../../../../harness/conventions.md#mr-016) [`MR-019`](../../../../harness/conventions.md#mr-019) [`MR-025`](../../../../harness/conventions.md#mr-025) [`MR-027`](../../../../harness/conventions.md#mr-027) [`MR-028`](../../../../harness/conventions.md#mr-028) [`MR-029`](../../../../harness/conventions.md#mr-029) [`MR-030`](../../../../harness/conventions.md#mr-030) — 0 offen (geprüft 2026-10-10).

## 8. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** `GATE` (Achsen 1, 2, 3 ✓).

**Vorgelagert — offene Beobachtungen sichten** (2026-10-10):
`BEO-GATE/versionsangabe-neben-digest-ungeprueft` (verkörpert) — Tag und Digest müssen zusammen
gehoben werden, `make version-coherence` hält `GO_VERSION` zwischen Makefile und Dockerfile;
`BEO-GATE/hebungskanal-haengt-an-repo-externen-schaltern` (1×) — betrifft die offene Frage, warum
kein Dependabot-PR kam.

**Modus-Begründungsblock:** alle berührten Sub-Areas GF.
