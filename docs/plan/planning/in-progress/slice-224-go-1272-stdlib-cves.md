# slice-224 — Go-Toolchain auf 1.27.2: zwei behebbare HIGH-CVEs der Standardbibliothek

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

**Lerneintrag — Form:** wird bei Closure benannt.

---

## 1. Ziel und Abgrenzung

**Ziel:** Das Release-Image wird mit Go **1.27.2** gebaut; die zwei behebbaren HIGH-Befunde aus dem
Nachtlauf `image-scan` vom 2026-10-09 (Lauf 37918923676) entfallen:
`CVE-2026-78667` (`net/http`, DoS) und `CVE-2026-97031` (`crypto/tls`, DoS), beide in `stdlib`
`v1.27.0`, behoben in 1.26.9 / 1.27.2. Die Mindestversion im `go.mod` beider Module steht auf
`1.27.2` (Maintainer: „Go ≥ 1.27.2").

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Das Release.** *Ein anderer Vorgang*: der Fix wirkt beim Konsumenten erst mit einem neuen
  Release; nach der Incident-Klausel in `releasing.md` ist das ein Fix-Forward-Patch-Release
  (`v0.23.2`) und folgt unmittelbar auf diesen Slice (Maintainer-Wort 2026-10-10: „Bitte auch ein
  Release erstellen").
- **`golangci-lint`-Image.** *Bestand bleibt bewusst stehen*: es baut nicht das ausgelieferte
  Binary; der Scan betrifft das Release-Bild.
- **Warum Dependabot den Hub nicht gemeldet hat.** *Ein anderer Vorgang*: nach dem Release prüfen,
  ob der Kanal aus [ADR-0038](../../adr/0038-dependabot-als-hebungskanal.md) die Basis-Hebung hätte vorschlagen müssen.

## 2. Definition of Done

- [ ] `golang`-Basis auf `1.27.2` mit Index-Digest in `Dockerfile` und `tools/archive-wave/Dockerfile`;
      `GO_VERSION` in beiden Makefiles; `go`-Direktive `1.27.2` in beiden `go.mod`.
- [ ] Beleg vor dem Release: Trivy gegen das lokal gebaute Multi-Arch-Archiv (`--input`) zeigt
      beide CVEs nicht mehr, je Plattform; Gegenprobe: das Archiv mit 1.27.0 zeigt sie.
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
| `Dockerfile`, `Makefile`, `go.mod` | update | Toolchain 1.27.2 |
| `tools/archive-wave/Dockerfile`, `tools/archive-wave/Makefile`, `tools/archive-wave/go.mod` | update | dasselbe für das Werkzeug |
| `CHANGELOG.md` | update | Sicherheits-Hebung, für Konsumenten sichtbar |

## 4. Trigger

**Start** (`open` → `in-progress`): Maintainer-Wort („Go ≥ 1.27.2 — siehe auch github ci",
2026-10-10), WIP-Limit frei.

**Rückführungen:** `in-progress` → `open` (blockiert): 1.27.2 bricht einen Gate-Lauf, der nicht in
einer Sitzung zu beheben ist.

## 5. Closure-Trigger

DoD vollständig, Gates grün, Closure-Notiz mit Lerneintrag.

## 6. Risiken und offene Punkte

- **Der Fix wirkt erst mit dem Release:** bis dahin bleibt der Nachtlauf rot. — **Ausgang:** *(bei Closure zuzuweisen: eingetreten / entfallen / weiter offen)*

## 7. Closure-Notiz

*(folgt bei Closure)*

## 8. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** `GATE` (Achsen 1, 2, 3 ✓).

**Vorgelagert — offene Beobachtungen sichten** (2026-10-10):
`BEO-GATE/versionsangabe-neben-digest-ungeprueft` (verkörpert) — Tag und Digest müssen zusammen
gehoben werden, `make version-coherence` hält `GO_VERSION` zwischen Makefile und Dockerfile;
`BEO-GATE/hebungskanal-haengt-an-repo-externen-schaltern` (1×) — betrifft die offene Frage, warum
kein Dependabot-PR kam.

**Modus-Begründungsblock:** alle berührten Sub-Areas GF.
