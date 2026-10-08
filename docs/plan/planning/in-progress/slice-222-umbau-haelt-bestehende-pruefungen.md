# slice-222 — Ein Umbau hält die bestehenden Prüfungen: Regel oder Sensor

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `git mv` (`make slice-mv`).

**Welle:** ohne Welle — die Closure-Bedingung wäre die eigene DoD.

**Bezug:** Lese-Schritt der slice-221-Closure für
[`BEO-GATE/umbau-verliert-pruefung-still`](../observations/BEO-GATE/umbau-verliert-pruefung-still/observation.md)
(slice-217, slice-218, slice-221 — 3×) — Ausgang *geplant* mit diesem Slice.

**Berührte Spec-Stellen:** — · Der Slice berührt kein Spec-Stratum.

**Verantwortlich:** Claude — gesetzt beim Übergang nach `next/` (Maintainer: „leg los“, 2026-10-08).

**Autor:** Claude. **Datum:** 2026-10-08.

**Lerneintrag — Form:** wird bei Closure benannt (erwartet: neuer Sensor oder geschärfte Regel).

---

## 1. Ziel und Abgrenzung

**Ziel:** Die Klasse „ein Umbau eines Prüfers verliert eine bestehende Prüfung, ohne dass es
jemand sagt" wird verkörpert. Zuerst ist zu messen, ob ein **Sensor** möglich ist — anders als bei
Prosa-Klassen ist ein Teil mechanisch fassbar: eine benannte Probe (`probe "<name>"` in einem
`--selftest`, ein Testfall) verschwindet aus einem Diff. Ist ein Sensor möglich, trägt er die
Klasse für diesen Teil; den Rest (eine Prüfung im Ablauf, kein benannter Fall — slice-218: die
OCI-Label-Prüfung) trägt eine Regel.

**Die drei Belege:** slice-217 (der Umbau gegen ein Review-Finding ließ den Abgleich der
Index-Beschriftung fallen), slice-218 (der Pipeline-Umbau ließ die Prüfung von fünf OCI-Labels
fallen), slice-221 (der Wechsel des Entscheidungslaufs auf JSON ließ die Anker-Probe des Selbsttests
fallen). Alle drei fing der unabhängige Review.

## 1b. Messung (2026-10-08)

Muster für einen **Fehlerpunkt** in Shell-Werkzeugen und Workflows: `fail "`, `::error::`,
`exit 1`/`exit 2`, `probe "`, `assert ` — gesucht in Zeilen, die ein Commit unter `tools/` oder
`.github/workflows/` **entfernt** und nicht wortgleich (nach Trim) im selben Commit wieder einfügt.

| # | Frage | Ergebnis | Geltungsbereich |
|---|---|---|---|
| M1 | Fängt eine **Zählung** (Fehlerpunkte vorher/nachher) die drei Belege? | **nein, 1 von 3** — `5131de5` 12→14, `2236547` 10→11 stiegen trotz Verlust; nur `bd68b1c` 11→10 sank | die drei Beleg-Commits, je eine Datei |
| M2 | Fängt die Menge der **entfernten, nicht verschobenen** Fehlerpunkt-Zeilen sie? | **ja, 3 von 3** — darunter je die verlorene Prüfung: „Config sagt" (`5131de5`), „OCI-Label … ist leer" (`bd68b1c`), „Marker nur am Zeilenende" (`2236547`) | die drei Beleg-Commits |
| M3 | Wie oft schlägt M2 im Bestand an? | **12 von 94** Commits mit Änderung an `tools/` oder `.github/workflows/` (40 Zeilen); Gegenzähler `git log -G` (Muster in geänderter Zeile, hinzu oder weg): 62 — eine Obermenge, die 12 liegen darin | ganzer Bestand seit 2026-06-21 |
| M4 | Trifft M2 die unbenannte Hälfte (Risiko §6)? | **ja** — die verlorene OCI-Label-Prüfung (slice-218) ist keine benannte Probe, aber eine `::error::`-Zeile | — |

**Folgerung:** ein **Sensor** trägt die Klasse für alle drei Belege. Er kann nicht entscheiden, ob
eine entfernte Prüfung verloren oder ersetzt ist — das ist ein Urteil —, aber er macht sie beim
Commit **sichtbar** und verlangt eine Begründung: Ein Commit, der eine nicht verschobene
Fehlerpunkt-Zeile entfernt, trägt in seiner Message eine Zeile `Entfernte-Pruefungen: <Grund>`.
Präzision im Bestand: 3 echte Verluste unter 12 Treffern; Kosten: eine Zeile in jedem achten
Werkzeug-Commit.

**Plan-Änderung 2026-10-08 (nach der Messung, vor dem Code):** Verkörpert wird als **Sensor**
`make verify-pruefung-entfernt` (Range-Prüfung wie `commit-scope-check`, im `preflight` und im
CI-Workflow über die Commit-Range) plus die Regel-Datei, die den Trailer beschreibt. Eine
Prosa-Regel daneben entfällt — M4 zeigt, dass der Sensor auch die unbenannte Hälfte trifft.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Nachprüfung früherer Umbauten.** *Ein anderer Vorgang*: die drei Fälle sind behoben; der Sensor
  prüft nur Commits in der Range, ältere bleiben ungeprüft.
- **Go-Tests.** *Bestand bleibt bewusst stehen*: keiner der drei Belege betraf Go-Code; ein
  entfernter Testfall dort ist `make test` und dem Review überlassen.
- **Eine Pflicht, jeden Umbau gegen die alte Fassung zu mutieren.** *Bestand bleibt bewusst
  stehen*: die Mutations-Gegenprobe (Mess-Regel 2) gilt schon; der Slice sucht den Zeitpunkt, an
  dem der Verlust **auffällt**, nicht einen zweiten Prüfschritt für dieselbe Probe.

## 2. Definition of Done

- [ ] Messung: welche Prüf-Orte im Repo benannte Fälle tragen (Selbsttests in `tools/*.sh`,
      Go-Tests), und ob ihr Verschwinden über eine Commit-Range mechanisch erkennbar ist — mit
      Geltungsbereich und Gegenprobe an den drei Belegen.
- [ ] Verkörperung: Sensor (wenn messbar) und/oder Regel mit Herkunfts-Anker `seit slice-222`;
      Register-Stand *verkörpert*.
- [ ] Unabhängiger Review, Report unter [`docs/reviews/`](../../../reviews/README.md).
- [ ] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [ ] Beobachtungs-Register fortgeschrieben (oder „keine Beobachtung" notiert).
- [ ] Jedes Risiko aus §6 trägt einen Ausgang.

`make gates` und `make verify` grün.

## 3. Plan (vor Code)

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| `tools/verify-pruefung-entfernt.sh` | neu | Sensor mit Selbsttest (synthetisches Repo) |
| `Makefile`, `.github/workflows/ci.yml`, `.claude/hooks/pretooluse-command-guard.sh` | update | Target, Range-Schritt in CI und `preflight`, Guard-Liste |
| `harness/rules/entfernte-pruefungen.md`, `AGENTS.md` §5, `harness/README.md`, `harness/sensors/verify-pruefung-entfernt.md` | neu / update | Regel, Zeiger, Gate-Index, Sensor-Datei |

## 4. Trigger

**Start** (`open` → `in-progress`): WIP-Limit frei; Maintainer-Wort.

**Rückführungen:** `in-progress` → `next` (zu groß): Sensor und Regel zugleich lassen sich nicht in
einer Review-Sitzung prüfen — dann teilen. `in-progress` → `open` (blockiert): entfällt.

## 5. Closure-Trigger

DoD vollständig, `make gates` und `make verify` grün, Closure-Notiz mit Lerneintrag.

## 6. Risiken und offene Punkte

- **Ein Sensor auf Proben-Namen trifft nur die benannte Hälfte:** der Fall slice-218 (eine Prüfung
  im Ablauf, keine benannte Probe) bliebe unsichtbar. — **Ausgang:** *(bei Closure zuzuweisen: eingetreten / entfallen / weiter offen)*

## 7. Closure-Notiz

*(folgt bei Closure)*

## 8. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** `GATE` (Achsen 1, 2, 3 ✓).

**Vorgelagert — offene Beobachtungen sichten** (2026-10-08):
`BEO-GATE/umbau-verliert-pruefung-still` (3×) ist dieser Slice selbst;
`BEO-GATE/pipefail-bricht-pruefer-stumm-ab` (1×) — ein neuer Shell-Sensor trägt dieselbe Gefahr.

**Modus-Begründungsblock:** alle berührten Sub-Areas GF.
