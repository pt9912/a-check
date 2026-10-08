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

**Lerneintrag — Form:** neuer Sensor.

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
`make pruefung-entfernt-check` (Range-Prüfung wie `commit-scope-check`, im `commit-msg`-Hook gegen
den Index, im `preflight` und im CI-Workflow über die Commit-Range) plus die Regel-Datei, die den Trailer beschreibt. Eine
Prosa-Regel daneben entfällt — M4 zeigt, dass der Sensor auch die unbenannte Hälfte trifft.

**Plan-Änderung 2026-10-08 (Review F-1/F-2/F-3, vor dem Fix):** Der Selbsttest, der vor jedem
Lauf startet, darf den Index des aufrufenden Repos nicht berühren — im Hook setzt git bei
`commit -a`/`commit <pfad>` einen absoluten `GIT_INDEX_FILE`; das Wegwerf-Repo läuft darum ohne
die git-Umgebung des Aufrufers. „Verschoben" heißt **dieselbe Zeile in derselben Datei**, gezählt
als Multimenge, nicht als Menge über den ganzen Commit. Die Zählstellen der Range-Schritte in
`releasing.md`, `harness/README.md` und einem `state.md` gehören dazu (AGENTS §4).

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Nachprüfung früherer Umbauten.** *Ein anderer Vorgang*: die drei Fälle sind behoben; der Sensor
  prüft nur Commits in der Range, ältere bleiben ungeprüft.
- **Go-Tests.** *Bestand bleibt bewusst stehen*: keiner der drei Belege betraf Go-Code; ein
  entfernter Testfall dort ist `make test` und dem Review überlassen.
- **Eine Pflicht, jeden Umbau gegen die alte Fassung zu mutieren.** *Bestand bleibt bewusst
  stehen*: die Mutations-Gegenprobe (Mess-Regel 2) gilt schon; der Slice sucht den Zeitpunkt, an
  dem der Verlust **auffällt**, nicht einen zweiten Prüfschritt für dieselbe Probe.

## 2. Definition of Done

- [x] Messung: welche Prüf-Orte im Repo benannte Fälle tragen (Selbsttests in `tools/*.sh`,
      Go-Tests), und ob ihr Verschwinden über eine Commit-Range mechanisch erkennbar ist — mit
      Geltungsbereich und Gegenprobe an den drei Belegen.
- [x] Verkörperung: Sensor (wenn messbar) und/oder Regel mit Herkunfts-Anker `seit slice-222`;
      Register-Stand *verkörpert*.
- [x] Unabhängiger Review, Report unter [`docs/reviews/`](../../../reviews/README.md).
- [x] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [x] Beobachtungs-Register fortgeschrieben (oder „keine Beobachtung" notiert).
- [x] Jedes Risiko aus §6 trägt einen Ausgang.

`make gates` und `make verify` grün.

## 3. Plan (vor Code)

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| `tools/pruefung-entfernt-check.sh` | neu | Sensor mit Selbsttest (Wegwerf-Repo ohne git-Umgebung des Aufrufers) |
| `Makefile`, `.github/workflows/ci.yml`, `.githooks/commit-msg`, `.claude/hooks/pretooluse-command-guard.sh` | update | Target, Range-Schritt in CI und `preflight`, Hook, Guard-Liste |
| `harness/rules/entfernte-pruefungen.md`, `AGENTS.md` §5, `harness/README.md`, `harness/sensors/pruefung-entfernt-check.md` | neu / update | Regel, Zeiger, Gate-Index, Sensor-Datei |
| `docs/user/releasing.md`, `harness/README.md`, `BEO-GATE/preflight-deckt-den-ci-schritt-nicht/state.md` | update | Zählstellen der Range-Schritte ohne Zahl |

## 4. Trigger

**Start** (`open` → `in-progress`): WIP-Limit frei; Maintainer-Wort.

**Rückführungen:** `in-progress` → `next` (zu groß): Sensor und Regel zugleich lassen sich nicht in
einer Review-Sitzung prüfen — dann teilen. `in-progress` → `open` (blockiert): entfällt.

## 5. Closure-Trigger

DoD vollständig, `make gates` und `make verify` grün, Closure-Notiz mit Lerneintrag.

## 6. Risiken und offene Punkte

- **Ein Sensor auf Proben-Namen trifft nur die benannte Hälfte:** der Fall slice-218 (eine Prüfung
  im Ablauf, keine benannte Probe) bliebe unsichtbar. — **Ausgang:** *entfallen* — gestrichen mit
  Begründung: der Sensor sucht nicht Proben-Namen, sondern Fehlerpunkt-Zeilen; die verlorene
  OCI-Label-Prüfung aus slice-218 ist eine `::error::`-Zeile und fällt darunter (Messung M4, mit
  abgeschaltetem Grandfathering an `bd68b1c` rot).

## 7. Closure-Notiz

**Lerneintrag — Form: neuer Sensor.** `make pruefung-entfernt-check` macht den Verlust einer Prüfung
beim Commit sichtbar: entfernt ein Commit unter `tools/` oder `.github/workflows/` eine
Fehlerpunkt-Zeile, die er nicht in derselben Datei wieder einfügt, verlangt er
`Entfernte-Pruefungen: <Grund>` — im `commit-msg`-Hook vor dem Commit, in `make preflight` und im
CI-Workflow über die Range. Alle drei Belege (slice-217, slice-218, slice-221) wären rot gewesen,
mit der verlorenen Prüfung in der Ausgabe. Regel 18 in `AGENTS.md` §5.

**Geliefert:** Sensor mit Selbsttest (20 Fälle, Wegwerf-Repo ohne die git-Umgebung des Aufrufers),
Regel-Datei, Sensor-Datei, Gate-Index, Hook, `preflight`, CI-Schritt; Zählstellen der
Range-Schritte ohne Zahl.

**Was hat funktioniert:** Messen, bevor entschieden wurde, ob ein Sensor möglich ist — die
naheliegende Zählung (vorher/nachher) fing 1 von 3, die entfernten Zeilen 3 von 3. Und der Sensor
hat seinen eigenen Fix-Commit gefangen: die ersetzten Selbsttest-Fälle trugen die Muster als Text;
der Commit kam mit Begründung durch.

**Was ging anders als geplant:** Der erste Wurf beschädigte bei `git commit -a` den Index des echten
Repos (Review F-1, HIGH): git setzt im Hook einen absoluten `GIT_INDEX_FILE`, und das Wegwerf-Repo
des Selbsttests erbte ihn. Meine Live-Probe hatte nur den gestagten Weg gedeckt. Dazu „verschoben"
als Menge über den ganzen Commit (F-2) und ein Range-Modus, der ein gescheitertes `git show` als
grün las (D-1). Offen als benannte Grenze: ein Amend, der nur die Begründung streicht, passiert den
Hook — die CI-Range fängt es (F-10). Kein CHANGELOG-Eintrag: die Regel bindet den Harness-Lauf,
nicht den Konsumenten (F-11, wie bei den Mess-Regeln).

**Steering-Loop-Eintrag:** neuer Sensor — liegt in `Makefile:pruefung-entfernt-check`.
Auslöser: `BEO-GATE/umbau-verliert-pruefung-still` (slice-217, slice-218, slice-221 — 3×).

**Beobachtungs-Register (`../observations/`):** `BEO-GATE/umbau-verliert-pruefung-still` →
*verkörpert* (`seit slice-222`). Eine neue Beobachtung: keine — der Index-Schaden ist ein Fall von
„Selbsttest berührt den Aufrufer" und steht hier als benannter Fund, nicht gezählt (ein Vorgang).

**Folge-Slices:** keine.

**Risiken aus §6:** das eine Risiko trägt seinen Ausgang (*entfallen*, mit Begründung).

**Drei Paarungen:** Anker — `seit slice-222` steht in der Regel-Datei (Zielort) · Folge-Slice —
keiner · Register — der genannte Pfad existiert mit nicht leerem `evidence/`.

**Trigger-Audit der aktiven MR:** [`MR-016`](../../../../harness/conventions.md#mr-016) [`MR-019`](../../../../harness/conventions.md#mr-019) [`MR-025`](../../../../harness/conventions.md#mr-025) [`MR-027`](../../../../harness/conventions.md#mr-027) [`MR-028`](../../../../harness/conventions.md#mr-028) [`MR-029`](../../../../harness/conventions.md#mr-029) [`MR-030`](../../../../harness/conventions.md#mr-030) — 0 offen (geprüft 2026-10-08).

## 8. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** `GATE` (Achsen 1, 2, 3 ✓).

**Vorgelagert — offene Beobachtungen sichten** (2026-10-08):
`BEO-GATE/umbau-verliert-pruefung-still` (3×) ist dieser Slice selbst;
`BEO-GATE/pipefail-bricht-pruefer-stumm-ab` (1×) — ein neuer Shell-Sensor trägt dieselbe Gefahr.

**Modus-Begründungsblock:** alle berührten Sub-Areas GF.
