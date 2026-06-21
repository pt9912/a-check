# slice-005 — Durchsetzungsschicht: PreToolUse-Command-Guard

**Status:** open (Backlog; wartet auf Trigger/Priorisierung).
**Welle:** welle-07-command-guard.
**Bezug:** [`AGENTS.md`](../../../../AGENTS.md) §3.1 (Docker/make-only) — die
Hard Rule, die dieser Guard erzwingt; Begründung Toolchain-Reproduzierbarkeit
([AC-QA-03](../../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit)) +
Hermetik ([AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)).
Folge-Kandidat aus [slice-004 §4](../done/slice-004-durchsetzungsschicht.md#4-closure-notiz-nach-done);
Stack-Vorbild `d-check`.

## Ziel

Den dritten — noch fehlenden — Bindepunkt der Durchsetzungsschicht
(Regelwerk-Grundlagen §Durchsetzungsschicht) schließen: das **Tool-Call-Gate**.
slice-004 hat das Meta-Gate (`gate-consistency`) und das Handoff-Gate
(`record-gates` + Stop-Hook) geliefert; was bleibt, ist der **PreToolUse-Hook**,
der eine Host-Toolchain-Verwendung *vor* der Ausführung blockt statt sie erst
im Nachhinein als roten Gate-Lauf sichtbar zu machen.

- **command-guard** (Tool-Call-Gate, *präventiv*): ein `.claude`-PreToolUse-Hook
  auf das `Bash`-Tool, der Host-Paketmanager/-Toolchains
  (`go`, `pip`, `npm`, `cargo`, `apt`, `brew`, … sowie `go build`/`go test`
  außerhalb von Docker) fail-closed ablehnt. Erlaubt bleiben `make`, `git`,
  `bash`, `docker` und Lese-Werkzeuge — die in [`AGENTS.md`](../../../../AGENTS.md)
  §3.1 als „der Host braucht nur" benannte Menge.

Damit ist die Durchsetzungsschicht für a-check vollständig (Tool-Call-Gate +
Handoff-Gate + Meta-Gate) und auf Parität mit `d-check`.

## Definition of Done

- `.claude/settings.json` trägt einen `PreToolUse`-Hook (Matcher `Bash`), der
  ein Guard-Skript unter `.claude/hooks/` aufruft.
- Das Guard-Skript lehnt verbotene Host-Toolchain-Aufrufe fail-closed ab
  (Exit/`deny`-Entscheidung) und gibt einen deutschen, umsetzbaren Grund zurück
  („nutze `make …`").
- Die Erlaubt-Menge (`make`/`git`/`bash`/`docker` + Lese-Tools) ist explizit und
  bewusst eng; `make`-Aufrufe, die intern `docker run` machen, passieren.
- Selbsttest: ein bekannt-verbotenes Kommando (`go build ./...`) wird im Test
  nachweislich abgelehnt, ein erlaubtes (`make help`) durchgelassen.
- [`AGENTS.md`](../../../../AGENTS.md) §3.1/§4 + [`harness/README.md`](../../../../harness/README.md)
  benennen den Guard (Durchsetzungsschicht vollständig).
- Beleg: `make gates` grün; der Guard ist über das Hook-Log/den Selbsttest belegt.

## Offene Fragen

- **Allowlist vs. Denylist:** enge Allowlist (nur `make`/`git`/`bash`/`docker` …)
  ist strenger und fail-closed-freundlicher, kann aber legitime Ad-hoc-Lesebefehle
  blocken; Denylist ist durchlässiger. Entscheidung bei Umsetzung.
- **Parsing-Tiefe:** nur das erste Token prüfen vs. Pipes/`env VAR=…`/`xargs`
  auflösen — wie weit geht die Umgehungs-Härtung, ohne brüchig zu werden?
- **MR-Eintrag:** Braucht die a-check-spezifische Allowlist einen eigenen
  [`MR-*`](../../../../harness/conventions.md)-Adaptionseintrag, oder ist sie
  Baseline-konform (wie bei slice-004 entschieden)?
- **Reichweite:** nur `Bash` oder auch andere ausführende Tools? (a-check nutzt
  primär `Bash`/`make`.)
