# slice-225 — Lint-Pin auf ein golangci-lint mit Go ≥ 1.27.2, dann `go 1.27.2` im `go.mod`

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `git mv` (`make slice-mv`).

**Welle:** ohne Welle — die Closure-Bedingung wäre die eigene DoD.

**Bezug:** [ADR-0005](../../adr/0005-lint-profil.md) (Lint-Profil),
[ADR-0038](../../adr/0038-dependabot-als-hebungskanal.md) (Hebung als bewusster Commit),
[AC-QA-03](../../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit) (digest-gepinnte Basis).

**Berührte Spec-Stellen:** — · Der Slice berührt kein Spec-Stratum.

**Verantwortlich:** —

**Autor:** Claude. **Datum:** 2026-10-10.

**Lerneintrag — Form:** wird bei Closure benannt.

---

## 1. Ziel und Abgrenzung

**Ziel:** Die `go`-Direktive in `go.mod` und `tools/archive-wave/go.mod` steht auf `1.27.2`
(Maintainer: „Go ≥ 1.27.2"), und `make lint` bleibt dabei hermetisch und grün. Voraussetzung ist
ein golangci-lint-Image, dessen **Programm** mit Go ≥ 1.27.2 gebaut ist: slice-224 hat gemessen,
dass `go 1.27.2` unter dem gepinnten `v2.13.2` (Programm und Image mit Go 1.27.0) eine Toolchain
aus dem Netz nachlädt und die Typprüfung bricht (`export data version 5 is greater than maximum
supported version 4`). Auch `v2.14.0`, Stand 2026-10-10 das neueste Tag, bringt nur Image-Go
1.27.1 und ein mit 1.27.0 gebautes Programm.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Die Build-Toolchain.** *Bestand bleibt bewusst stehen*: `golang:1.27.2` ist seit slice-224
  gepinnt; die CVE-Frage ist damit beantwortet, hier geht es nur um Lint und Mindestversion.
- **Ein Release.** *Ein anderer Vorgang*: Lint-Pin und `go`-Direktive ändern das ausgelieferte
  Image nicht; ein Release folgt mit der nächsten Produktänderung.
- **Neue Linter oder Profil-Änderungen**, die ein neues golangci-lint mitbringt. *Ein anderer
  Vorgang*: meldet das neue Release neue Befunde, werden sie behoben, nicht das Profil gelockert —
  eine Lockerung wäre eine ADR ([AGENTS.md](../../../../AGENTS.md) §3.6).

## 2. Definition of Done

- [ ] `GOLANGCI_LINT_VERSION` in `Makefile` und `Dockerfile` auf das neue Tag, Digest im `FROM`
      gehoben; gemessen: Programm **und** Image-Go ≥ 1.27.2.
- [ ] `go 1.27.2` in `go.mod` und `tools/archive-wave/go.mod`; `make lint` lädt keine Toolchain
      nach (Lint-Lauf ohne `toolchain`-Download in der Ausgabe).
- [ ] CHANGELOG `[Unreleased]`, falls die Mindestversion für Konsumenten sichtbar ist (Bau aus
      Quellen).
- [ ] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [ ] Beobachtungs-Register fortgeschrieben (oder „keine Beobachtung" notiert).
- [ ] Jedes Risiko aus §6 trägt einen Ausgang.

`make ci`, `make archive-wave-test`, `make gates` und `make verify` grün. Ein unabhängiger Review
ist nicht vorgesehen ([MR-019](../../../../harness/conventions.md#mr-019): Opt-in) — eine
Versions-Hebung mit Messung vorher und nachher.

## 3. Plan (vor Code)

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| `Makefile`, `Dockerfile` | update | Lint-Pin (Tag + Digest) |
| `go.mod`, `tools/archive-wave/go.mod` | update | `go 1.27.2` |
| `CHANGELOG.md` | update | nur falls für Konsumenten sichtbar |

## 4. Trigger

**Start** (`open` → `next`): ein golangci-lint-Release ist auf Docker Hub, dessen Programm mit
Go ≥ 1.27.2 gebaut ist — prüfbar mit `golangci-lint version` im Image („built with go1.27.2" oder
höher). Stand 2026-10-10 nicht erfüllt (`v2.14.0`: built with go1.27.0).

**Rückführungen:** `in-progress` → `open` (blockiert): das neue golangci-lint meldet Befunde, die
nicht in einer Sitzung zu beheben sind.

## 5. Closure-Trigger

DoD vollständig, Gates grün, Closure-Notiz mit Lerneintrag.

## 6. Risiken und offene Punkte

- **Das neue golangci-lint meldet neue Befunde:** ein Minor-Sprung bringt oft neue Prüfungen in
  aktivierten Linter mit. — **Ausgang:** *(bei Closure zuzuweisen: eingetreten / entfallen / weiter offen)*
- **Der Trigger hat keinen Wächter:** Dependabot führt das Ökosystem `docker` bewusst nicht; ohne
  Nachsehen bleibt der Slice liegen. — **Ausgang:** *(bei Closure zuzuweisen: eingetreten / entfallen / weiter offen)*

## 7. Closure-Notiz

*(folgt bei Closure)*

## 8. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** `GATE` (Achsen 1, 2, 3 ✓).

**Vorgelagert — offene Beobachtungen sichten** (2026-10-10):
`BEO-GATE/versionsangabe-neben-digest-ungeprueft` (verkörpert) — Tag und Digest zusammen heben,
`make version-coherence` hält `GOLANGCI_LINT_VERSION` zwischen Makefile und Dockerfile;
`BEO-GATE/hebungskanal-haengt-an-repo-externen-schaltern` (1×) — berührt das zweite Risiko.

**Modus-Begründungsblock:** alle berührten Sub-Areas GF.
