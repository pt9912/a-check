# slice-004 — Durchsetzungsschicht: gate-consistency + record-gates

**Status:** done.
**Welle:** welle-04-durchsetzungsschicht.
**Bezug:** [`AGENTS.md`](../../../../AGENTS.md) §3.6 (Gate-Disziplin) + §4 (`make gates`);
[ADR-0005](../../adr/0005-lint-profil.md)/[ADR-0006](../../adr/0006-coverage-gate.md)
(bestehende Gates); Stack-Vorbild `d-check`.
**Autor:** pt9912. **Datum:** 2026-06-21.

---

## 1. Ziel

Die zwei computational-Bindepunkte der Durchsetzungsschicht (Regelwerk-Grundlagen
§Durchsetzungsschicht), die a-check ggü. `d-check` fehlten, sind geschlossen:
`gate-consistency` (Doku ↔ Makefile) und `record-gates` (inhaltsbasierter
Working-Tree-Hash-Nachweis) inklusive `.claude`-Stop-Hook (Handoff-Gate).

## 2. Definition of Done

- [x] `make gate-consistency` existiert, bricht bei Drift ab, mit Selbsttest (Phantom-Target feuert).
- [x] `make record-gates` schreibt den Nachweis; in `make gates` als letzter Schritt (Reihenfolge via `.NOTPARALLEL`).
- [x] Nachweis-Ablage `.harness/state/` (in `.gitignore`).
- [x] `.claude`-Stop-Hook (`stop-require-gates.sh` + `settings.json`) erzwingt den Nachweis vor „fertig" (loop-guarded, bootstrap-aware).
- [x] [`AGENTS.md`](../../../../AGENTS.md) §4 + [`harness/README.md`](../../../../harness/README.md) §Sensors um beide Gates ergänzt.
- [x] `make gates` grün inkl. der neuen Gates.

## 3. Umsetzung

- `tools/gate-consistency.sh` — Doku ↔ Makefile bidirektional (inkl. `d-check.mk`); Utility-Allowlist `help`/`build`/`compile`; `.d-check.yml`-Modulprüfung ([AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)); Selbsttest.
- `tools/harness/working-tree-hash.sh` + `tools/harness/record-gates.sh` — inhaltsbasierter Nachweis (gemeinsame Hash-Quelle).
- `.claude/hooks/stop-require-gates.sh` + `.claude/settings.json` — Stop-Hook (Handoff-Gate).
- `Makefile` — zwei Targets + `make gates`-Erweiterung (`record-gates` zuletzt).

## 4. Closure-Notiz (nach `done/`)

**Belege:** `make gates` grün — `gate-consistency` ok (Selbsttest gefeuert),
`record-gates` Nachweis geschrieben, übrige Gates grün; `make doc-check` grün.

**Lerneintrag (Steering-Loop):**

- *Neuer Sensor:* `gate-consistency` schließt die im
  [slice-003-Review](../../../reviews/2026-06-21-slice-003-impl-gates.md)
  (Befund B-M4) benannte Drift-Klasse maschinell — genau die „real/geplant"-
  Doku-Drift und die mehrfach halluzinierte Datei-/Target-Existenz werden jetzt
  fail-closed gefangen.
- *Geschärfte Regel:* `AGENTS.md` §4 ist die **vollständige** Gate-Liste;
  Utility-Targets (`help`/`build`/`compile`) sind explizit als Nicht-Gates
  allowlisted — die Bidirektionalität (kein halluziniertes *und* kein
  undokumentiertes Gate) ist erzwungen.
- *Handoff-Gate:* `record-gates` + Stop-Hook machen „ich hab die Gates laufen
  lassen" inhaltsbasiert prüfbar (Regelwerk §Durchsetzungsschicht: fail-closed,
  Inhalts-Nachweis, Loop-Guard, bootstrap-aware).

**Offene Fragen — aufgelöst:**

- Eigener `MR-*`-Eintrag? **Nein** — die Mechanik ist Baseline-Konformität
  (Regelwerk §Durchsetzungsschicht), keine Adaption ggü. der Baseline; keine
  stille Setzung.
- Skript vs. Hook-Reihenfolge: beide angelegt — das Skript schreibt, der Hook
  prüft denselben Hash (eine Quelle).

**Folge-Kandidat (open):** der dritte Bindepunkt — der **PreToolUse-
Command-Guard** (Docker-only Tool-Call-Gate) — ist noch nicht adoptiert.

## 5. Sub-Area-Modus-Begründung

### Sub-Area: Harness-Durchsetzungsschicht

- **Modus:** GF — die Mechanik wird neu angelegt (Skript/Doc führt), kein Bestand zu inventarisieren.
- **Konventionen-Dichte:** hoch (Regelwerk §Durchsetzungsschicht + `d-check`-Vorbild).
- **Phase-Reife:** Phase 5 — Gates erzwingen, `make gates` grün.
- **Evidenz-/Diskrepanz-Risiko:** niedrig (Greenfield).
- **Reconciliation-Aufwand:** keiner; Folge-Option Command-Guard.
