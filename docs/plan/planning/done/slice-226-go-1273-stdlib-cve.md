# slice-226 — Go-Toolchain auf 1.27.3: behebbare HIGH-CVE der Standardbibliothek

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

**Autor:** Claude. **Datum:** 2026-10-22.

**Lerneintrag — Form:** geschärfte Regel.

---

## 1. Ziel und Abgrenzung

**Ziel:** Das Release-Image wird mit Go **1.27.3** gebaut; der behebbare HIGH-Befund aus dem
Nachtlauf `image-scan` vom 2026-10-22 (Lauf 39749572989, auch 2026-10-21) entfällt:
`CVE-2026-111943` (`crypto/x509`, Man-in-the-middle) in `stdlib` `v1.27.2`, behoben in
1.26.10 / 1.27.3.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Das Release.** *Ein anderer Vorgang*: der Fix wirkt beim Konsumenten erst mit einem neuen
  Release (Fix-Forward-Patch-Release nach `releasing.md`); es folgt unmittelbar auf diesen Slice.
- **Die `go`-Direktive in den `go.mod` und der Lint-Pin.** *Ein Folge-Slice übernimmt es*:
  slice-225 führt sie; die Direktive nennt die Mindest-Sprachversion, nicht die Build-Toolchain.
- **`golangci-lint`-Image.** *Bestand bleibt bewusst stehen*: es baut nicht das ausgelieferte Binary.
- **`tzdata` (`DLA-4792-1`, UNKNOWN).** *Bestand bleibt bewusst stehen*: nicht handlungspflichtig
  nach [ADR-0037](../../adr/0037-cve-scan-gegen-das-publizierte-image.md), liegt in der Laufzeit-Basis.

## 2. Definition of Done

- [x] `golang`-Basis auf `1.27.3` mit Index-Digest in `Dockerfile` und `tools/archive-wave/Dockerfile`;
      `GO_VERSION` in beiden Makefiles.
- [x] Beleg vor dem Release: Trivy gegen das lokal gebaute Multi-Arch-Archiv zeigt
      `CVE-2026-111943` nicht mehr, je Plattform mit Architektur-Nachweis.
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
| `Dockerfile`, `Makefile` | update | Toolchain 1.27.3 |
| `tools/archive-wave/Dockerfile`, `tools/archive-wave/Makefile` | update | dasselbe für das Werkzeug |
| `CHANGELOG.md` | update | Sicherheits-Hebung, für Konsumenten sichtbar |

## 4. Trigger

**Start** (`open` → `in-progress`): Nachtlauf `image-scan` rot seit 2026-10-21, Maintainer-Wort
(„ja", 2026-10-22), WIP-Limit frei.

**Rückführungen:** `in-progress` → `open` (blockiert): 1.27.3 bricht einen Gate-Lauf, der nicht in
einer Sitzung zu beheben ist.

## 5. Closure-Trigger

DoD vollständig, Gates grün, Closure-Notiz mit Lerneintrag.

## 6. Risiken und offene Punkte

- **Der Fix wirkt erst mit dem Release:** bis dahin bleibt der Nachtlauf rot. — **Ausgang:** *entfallen* — gestrichen mit Begründung: das Release `v0.23.3` folgt unmittelbar auf diesen Slice, danach ist der Nachtlauf gegen das neue `latest` grün.

## 7. Closure-Notiz

**Lerneintrag — Form: geschärfte Regel.** Dieselbe Lücke wie in slice-224 hat sich wiederholt: die
Basis-Hebung ist ein bewusster Commit ([ADR-0038](../../adr/0038-dependabot-als-hebungskanal.md)),
und der Nachtlauf `image-scan` ist der Kanal, der sie meldet. Der Ad-hoc-Scan gegen das
Multi-Arch-Archiv lief von Beginn an je Plattform über ein Ein-Manifest-Layout und nannte die
Architektur im Bericht (`amd64`, `arm64`); die Regel aus slice-224 hat damit gegriffen, ohne neue
Beobachtung.

**Geliefert:** `golang:1.27.3` per Index-Digest (`sha256:f20d83a8…`) im Release- und im
`archive-wave`-Dockerfile, `GO_VERSION` in beiden Makefiles, CHANGELOG `[Unreleased]` (Security).

**Messung:** Trivy `0.74.0` gegen das lokal gebaute Multi-Arch-Archiv, je Plattform mit
Architektur-Nachweis: `amd64` und `arm64` ohne HIGH/CRITICAL-Befund. Gegenprobe: der Nachtlauf vom
2026-10-22 gegen das publizierte `v0.23.2` meldete je Plattform `CVE-2026-111943`, installiert
`v1.27.2`, behoben in 1.27.3.

**Was hat funktioniert:** das Muster aus slice-224 (Index-Digest, vier Dateien, ein CHANGELOG-Eintrag).

**Was ging anders als geplant:** nichts.

**Steering-Loop-Eintrag:** geschärfte Regel — gezählt, nicht verkörpert.

**Beobachtungs-Register (`../observations/`):** keine Beobachtung angefallen.

**Folge-Slices:** keine.

**Risiken aus §6:** das eine Risiko trägt seinen Ausgang (*entfallen*, mit Begründung).

**Drei Paarungen:** Anker — kein `liegt in`-Feld · Folge-Slice — keiner · Register — keine Zitate.

**Trigger-Audit der aktiven MR:** [`MR-016`](../../../../harness/conventions.md#mr-016) [`MR-019`](../../../../harness/conventions.md#mr-019) [`MR-025`](../../../../harness/conventions.md#mr-025) [`MR-027`](../../../../harness/conventions.md#mr-027) [`MR-028`](../../../../harness/conventions.md#mr-028) [`MR-029`](../../../../harness/conventions.md#mr-029) [`MR-030`](../../../../harness/conventions.md#mr-030) — 0 offen (geprüft 2026-10-22).

## 8. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** `GATE` (Achsen 1, 2, 3 ✓).

**Vorgelagert — offene Beobachtungen sichten** (2026-10-22):
`BEO-GATE/versionsangabe-neben-digest-ungeprueft` (verkörpert) — Tag und Digest zusammen heben,
`make version-coherence` hält `GO_VERSION` zwischen Makefile und Dockerfile;
`BEO-GATE/hebungskanal-haengt-an-repo-externen-schaltern` (1×) — der Nachtlauf trägt wieder den Kanal.

**Modus-Begründungsblock:** alle berührten Sub-Areas GF.
