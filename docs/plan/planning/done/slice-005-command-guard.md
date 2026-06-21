# slice-005 — Durchsetzungsschicht: PreToolUse-Command-Guard

**Status:** done.
**Welle:** welle-07-command-guard.
**Bezug:** [`AGENTS.md`](../../../../AGENTS.md) §3.1 (Docker/make-only) — die
Hard Rule, die dieser Guard erzwingt; Begründung Toolchain-Reproduzierbarkeit
([AC-QA-03](../../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit)) +
Hermetik ([AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)).
Folge-Kandidat aus [slice-004 §4](slice-004-durchsetzungsschicht.md#4-closure-notiz-nach-done);
Stack-Vorbild `d-check`.

---

## 1. Ziel

Den dritten — zuvor fehlenden — Bindepunkt der Durchsetzungsschicht
(Regelwerk-Grundlagen §Durchsetzungsschicht) schließen: das **Tool-Call-Gate**.
slice-004 lieferte Meta-Gate (`gate-consistency`) und Handoff-Gate
(`record-gates` + Stop-Hook); slice-005 ergänzt den **PreToolUse-Hook**, der
Host-Toolchain-Nutzung *vor* der Ausführung blockt statt sie erst als roten
Gate-Lauf sichtbar zu machen. Damit ist die Durchsetzungsschicht vollständig
(Tool-Call- + Handoff- + Meta-Gate) und auf Parität mit `d-check`.

## 2. Definition of Done

- [x] `.claude/settings.json` trägt einen `PreToolUse`-Hook (Matcher `Bash`), der `.claude/hooks/pretooluse-command-guard.sh` aufruft.
- [x] Das Guard-Skript lehnt verbotene Host-Toolchain-Aufrufe fail-closed ab (`{"decision":"block"}` + deutscher, umsetzbarer Grund); Pass-Fall ohne Ausgabe.
- [x] Denylist eng auf §3.1 geschnitten; `make`/`git`/`docker` sowie Toolnamen, die nur als Argument/String vorkommen, passieren.
- [x] Selbsttest: `go build ./...`/`sudo apt-get`/`bash -lc "npm…"`/abs. Pfad werden geblockt, `make help`/`docker run … npm`/`git commit -m "…pip…"` durchgelassen.
- [x] [`AGENTS.md`](../../../../AGENTS.md) §3.1/§4 + [`harness/README.md`](../../../../harness/README.md) §Sensors benennen den Guard; `make guard-selftest` in `make gates`.
- [x] `make gates` grün inkl. `guard-selftest`.

## 3. Umsetzung

- `.claude/hooks/pretooluse-command-guard.sh` — node-gestützte Denylist
  (Host-Paketmanager + Host-Go); segmentweise Befehlsposition, rekursive
  Sub-Shell-Auflösung (`-lc`/`-ec`, Tiefe ≤ 3 → fail-closed); eingebauter
  `--selftest`. Logik gespiegelt von `d-check`, IDs/Grund auf a-check adaptiert.
- `.claude/settings.json` — `PreToolUse`/`Bash`-Verdrahtung (neben dem Stop-Hook).
- `Makefile` — Target `guard-selftest`, eingehängt in `make gates` vor `record-gates`.

## 4. Closure-Notiz (nach `done/`)

**Belege:** `make gates` grün — `guard-selftest` ok (Denylist greift: Host-go/
apt/pip/npm/cargo/golangci-lint geblockt, `make`/`git`/`docker` erlaubt), übrige
Gates grün; `make doc-check` grün. Normalmodus verifiziert: `go test ./...`
→ `block`-JSON, `make gates` → keine Ausgabe (Pass).

**Lerneintrag (Steering-Loop):**

- *Durchsetzungsschicht vollständig:* alle drei Regelwerk-Bindepunkte sind nun
  real — Tool-Call-Gate (dieser Guard, präventiv), Handoff-Gate
  (`record-gates` + Stop-Hook) und Meta-Gate (`gate-consistency`). Die §3.1-Regel
  ist nicht mehr nur Text, sondern fail-closed durchgesetzt.
- *Geschärfte Regel:* der Guard prüft die **Befehlsposition** je Segment (nicht
  bloße Substring-Suche) und löst Sub-Shells rekursiv auf — Toolnamen in
  Argumenten/Strings (`git commit -m "…pip…"`, `docker run img npm test`) sind
  bewusst erlaubt; das hält den Stolperdraht treffsicher statt brüchig.
- *Dokumentierte Restlücke:* andere Interpreter (`python -c`, `node -e`,
  `find -exec`) und tief obfuskierte Aufrufe werden nicht gefangen — der Guard
  ist ein Stolperdraht, keine Sandbox (analog zur Stop-Hook-Restlücke; CI/Review
  ist das Netz).

**Offene Fragen — aufgelöst:**

- *Allowlist vs. Denylist:* **Denylist** — präzise auf §3.1 (Host-Toolchains)
  geschnitten; eine enge Allowlist hätte den Lese-/`git`-/`make`-Workflow
  blockiert, ohne mehr Schutz für genau das geschützte Gut zu liefern.
- *Parsing-Tiefe:* segmentweise Kommando-Position + rekursive Sub-Shell-Auflösung
  (Tiefe ≤ 3, dann fail-closed); tiefere Obfuskation bleibt bewusste Restlücke.
- *MR-Eintrag:* **keiner** — der Guard ist Baseline-Konformität (Regelwerk
  §Durchsetzungsschicht), keine Adaption ggü. der Baseline; konsistent mit der
  slice-004-Entscheidung für die gesamte Durchsetzungsschicht.
- *Reichweite:* nur `Bash` — a-check führt ausschließlich darüber aus.

**Folge-Kandidaten:** keine offenen Durchsetzungsschicht-Bindepunkte mehr; nächste
Welle ist `welle-05-release` (siehe [Roadmap](../in-progress/roadmap.md)).

## 5. Sub-Area-Modus-Begründung

### Sub-Area: Harness-Durchsetzungsschicht

- **Modus:** GF — die Mechanik wird neu angelegt (Skript/Doc führt), kein Bestand zu inventarisieren.
- **Konventionen-Dichte:** hoch (Regelwerk §Durchsetzungsschicht + `d-check`-Vorbild).
- **Phase-Reife:** Phase 5 — Gate erzwingen, `make gates` grün.
- **Evidenz-/Diskrepanz-Risiko:** niedrig (Greenfield, gespiegelte Logik + Selbsttest).
- **Reconciliation-Aufwand:** keiner; Durchsetzungsschicht abgeschlossen.
