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

**Lerneintrag — Form:** wird bei Closure benannt.

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
  nach ADR-0037, liegt in der Laufzeit-Basis.

## 2. Definition of Done

- [ ] `golang`-Basis auf `1.27.3` mit Index-Digest in `Dockerfile` und `tools/archive-wave/Dockerfile`;
      `GO_VERSION` in beiden Makefiles.
- [ ] Beleg vor dem Release: Trivy gegen das lokal gebaute Multi-Arch-Archiv zeigt
      `CVE-2026-111943` nicht mehr, je Plattform mit Architektur-Nachweis.
- [ ] CHANGELOG `[Unreleased]`.
- [ ] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [ ] Beobachtungs-Register fortgeschrieben (oder „keine Beobachtung" notiert).
- [ ] Jedes Risiko aus §6 trägt einen Ausgang.

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

- **Der Fix wirkt erst mit dem Release:** bis dahin bleibt der Nachtlauf rot. — **Ausgang:** *(bei Closure zuzuweisen: eingetreten / entfallen / weiter offen)*

## 7. Closure-Notiz

*(folgt bei Closure)*

## 8. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** `GATE` (Achsen 1, 2, 3 ✓).

**Vorgelagert — offene Beobachtungen sichten** (2026-10-22):
`BEO-GATE/versionsangabe-neben-digest-ungeprueft` (verkörpert) — Tag und Digest zusammen heben,
`make version-coherence` hält `GO_VERSION` zwischen Makefile und Dockerfile;
`BEO-GATE/hebungskanal-haengt-an-repo-externen-schaltern` (1×) — der Nachtlauf trägt wieder den Kanal.

**Modus-Begründungsblock:** alle berührten Sub-Areas GF.
