# slice-036 — d-check-Pin-Hebung v0.37.1 → v0.51.1

**Status:** done (**abgeschlossen 2026-07-23** — `make ci` grün, mechanischer Pin-Bump, Fragment-Parität verifiziert). Closure-Notiz: §5.
**Typ:** Tooling-Pin-Hebung (Harness, `doc-check`), Folge von [slice-019](../done/slice-019-dcheck-mk-print-mk-angleichung.md).
**Bezug:** Harness-Prozess (`make doc-check` via `d-check`, [AGENTS §4](../../../../AGENTS.md#4-quality-gates)); **kein a-check-Vertrag/ADR** — der Pin betrifft das Schwester-Tool, nicht das a-check-Image. [Roadmap](../in-progress/roadmap.md).

## 1. Motivation

Der d-check-Pin steht seit [slice-019](../done/slice-019-dcheck-mk-print-mk-angleichung.md) auf
`v0.37.1`; Maintainer-Wunsch: Hebung auf **v0.51.1** (14 Minor-Versionen). d-check ist
digest-gepinnt (Digest sticht Tag, sinngemäß
[AC-QA-03](../../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit)); die Hebung ist
eine bewusste Pin-Bump-Änderung an `d-check.mk` + Doku, wie slice-019.

## 2. Design / Umfang (umgesetzt)

1. **Kompatibilität verifiziert:** `doc-check` gegen `d-check@sha256:fede3d02…` (v0.51.1)
   auf dem eigenen Baum → **0 Befunde** (96 Dateien); die `.d-check.yml` ist unverändert
   kompatibel. ✅
2. **`d-check.mk` regeneriert:** verbatim aus `d-check:v0.51.1 --print-mk`, nur die eine
   a-check-Anpassung (DCHECK_DIGEST-Pin + Anpassungs-Kommentar). Neu im Fragment:
   Target **`doc-targets`** (Modul `targets`, DC-FA-TGT-001) + drei neue Module
   `targets`/`citations`/`sources` in den `--disable`-Listen der Einzel-Modul-Targets
   (saubere Obermenge, kein Bruch). `DCHECK_IMAGE`-Tag v0.37.1 → **v0.51.1**,
   `DCHECK_DIGEST` → `sha256:fede3d02…`. ✅
3. **Doku:** [AGENTS §4](../../../../AGENTS.md#4-quality-gates) — Pin-Version v0.37.1 →
   v0.51.1 (8 advisory-Zeilen) + neue `doc-targets`-Zeile (gate-consistency Check 2:
   jedes reale `d-check.mk`-Target muss in AGENTS §4 stehen);
   [conventions.md](../../../../harness/conventions.md#mr-005--referenzmatrix-intra-spec-richtung--adrslice-disziplin-d-check-angleichung) [MR-005](../../../../harness/conventions.md#mr-005--referenzmatrix-intra-spec-richtung--adrslice-disziplin-d-check-angleichung)-Pin-Notiz. ✅
4. **Gates:** `make gates` (doc-check jetzt via v0.51.1) — grün. ✅

## 3. Akzeptanzkriterien

- **Happy:** `make doc-check` läuft über `d-check@…fede3d02…` (v0.51.1) → 0 Befunde;
  `make gates` grün (gate-consistency: `d-check.mk` wohlgeformter Tag + Digest, alle realen
  Targets in AGENTS §4).
- **Boundary:** `d-check.mk` == v0.51.1-`--print-mk` **bis auf** die eine a-check-Anpassung
  (DCHECK_DIGEST-Pin + Kommentar); der `doc-targets`-Zusatz ist dokumentiert.
- **Negative:** ein fehlerhafter DCHECK_DIGEST (nicht `sha256:<64hex>`) → gate-consistency
  Check (C) **rot** (bestehender Selbsttest).

## 4. Grenzen / Folge

- **Kein a-check-Vertrag/ADR/Release:** der Pin ist Harness-Tooling; das a-check-Image ist
  unberührt. Kein Lastenheft-/Spec-/Handbuch-Bump.
- Das neue Modul `targets` (DC-FA-TGT-001) ist als advisory-Target verdrahtet, aber **nicht**
  ins mandatory `doc-check`-Aggregat aufgenommen (wie die übrigen `doc-*`-advisory-Module);
  eine Aufnahme wäre ein eigener Härtungs-Slice (Präzedenz slice-029).

## 5. Closure-Notiz (nach `done`)

**Abgeschlossen 2026-07-23** auf Branch `slice-036-dcheck-v0.51.1`.

**Geliefert:** d-check-Pin `v0.37.1` → **`v0.51.1`** (`d-check.mk` verbatim aus
`d-check:v0.51.1 --print-mk`, einzige a-check-Anpassung = DCHECK_DIGEST-Pin
`sha256:fede3d02…`); neues Fragment-Target `doc-targets` (Modul `targets`,
DC-FA-TGT-001) + drei neue `--disable`-Module (`targets`/`citations`/`sources`) —
saubere Obermenge. Doku: AGENTS §4 (Pin-Version + `doc-targets`-Zeile),
`conventions.md` [MR-005](../../../../harness/conventions.md#mr-005--referenzmatrix-intra-spec-richtung--adrslice-disziplin-d-check-angleichung)-Notiz. Kein a-check-Vertrag/ADR/Handbuch/Release-Bump.

**Gate-Evidenz:** `make ci` grün — doc-check jetzt über `d-check@…fede3d02…` (v0.51.1):
**97 Dateien, 0 Befunde**; gate-consistency ok (`d-check.mk` wohlgeformter Tag + Digest,
`doc-targets` in AGENTS §4). Kompatibilität über 14 Minors verifiziert.

**Lernnotiz:** die v0.51.1-`doc-check` ist strenger (Modul `targets`) und fing beim ersten
Lauf zwei tote `slice-019`-Relativlinks (Ziel liegt in `done/`, nicht `in-progress/`) + eine
unverlinkte [MR-005](../../../../harness/conventions.md#mr-005--referenzmatrix-intra-spec-richtung--adrslice-disziplin-d-check-angleichung) im neuen Slice-Doc — der Pin-Bump härtet das Doku-Gate mit.

**Folge:** Modul `targets` als advisory verdrahtet, nicht ins mandatory `doc-check` (eigener
Härtungs-Slice bei Bedarf, Präzedenz slice-029). d-check-`citations`/`sources`-Module noch
ungenutzt.
