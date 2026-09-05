# slice-124 — CVE-Scan gegen das publizierte Image

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese Datei liegt — eines von
`open/`, `next/`, `in-progress/`, `done/`. Er wechselt nur durch `make slice-mv`
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).

**Bezug:** Maintainer-Auftrag 2026-08-30; Muster aus dem Schwester-Repo (`tools/image-scan.sh`,
`.github/workflows/image-scan.yml`). **Vorbereitung für** den Hebungs-Kanal (Dependabot) — der
Sensor kommt zuerst, sonst hebt der Kanal nichts Gemessenes.

**Berührte Spec-Stellen:** — *(keine)*

**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.

**Autor:** Claude (Opus 5), im Auftrag des Maintainers. **Datum:** 2026-08-30.

---

## 1. Ziel

Bekannte Schwachstellen im **publizierten** Image werden gemeldet, ohne dass ein Commit sie
auslösen muss.

## 2. Definition of Done

- [x] Eine ADR begründet die drei Entscheidungen, die dieser Sensor trifft — **Netz als Zweck**
      (nicht als Zugeständnis), **nicht im `gates`-Aggregat**, und **nur behebbare CRITICAL/HIGH**
      entscheiden über rot. Sie steht im ADR-Index.
- [x] `tools/image-scan.sh` + `make image-scan`: Trivy digest-gepinnt, Auswertung als **netzlos
      prüfbare** Funktion mit `--selftest`, fail-closed bei leerer Prüfmenge. Exit 0/1/2.
- [x] `.github/workflows/image-scan.yml` fährt ihn zeitgesteuert **und** per `workflow_dispatch`;
      der Ausgang wird aus dem **Log** gelesen, nicht aus dem Exit-Code. Deklaration in
      [`AGENTS.md`](../../../../AGENTS.md) §4 und
      [`harness/README.md`](../../../../harness/README.md) §Sensors.

- [x] `make gates` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.
- [x] Closure-Notiz mit benanntem Lerneintrag geschrieben (§7).
- [x] Beobachtungs-Register fortgeschrieben.
- [x] Jedes Risiko aus §6 trägt genau einen Ausgang.

## 3. Plan (vor Code)

| Datei | Änderungs-Art | Begründung |
|---|---|---|
| `docs/plan/adr/00NN-*.md` + [Index](../../adr/README.md) | neu | die drei Entscheidungen; Index-Pflicht (§5) |
| `tools/image-scan.sh` | neu | der Sensor, mit `--selftest` |
| [`Makefile`](../../../../Makefile) | update | Target + `.PHONY`; **nicht** ins `gates`-Aggregat |
| `.github/workflows/image-scan.yml` | neu | der Nachtlauf |
| [`AGENTS.md`](../../../../AGENTS.md) §4, [`harness/README.md`](../../../../harness/README.md) | update | Deklaration; `doc-targets` erzwingt sie |
| `.claude/hooks/pretooluse-command-guard.sh` | update | `NICHT_PRUEFEND` — siehe unten |

**Auszuführende Gates:** `make gates` — tragend `doc-targets` (beide Doku-Tabellen),
`guard-selftest` (Drift-Wächter der Listen) und `gate-consistency`. Zum Abschluss `make verify`.

### Das Risiko ist vorab gemessen

| Prüfung | Ergebnis |
|---|---|
| `ghcr.io/pt9912/a-check:latest` | existiert (v0.17.0) — der Scan hat einen Gegenstand |
| Trivy-Pin `0.74.0` | Digest aus zweiter Quelle bestätigt (eigener Pull == Schwester-Repo) |
| bestehende Netz-Gates in a-check | **keine** — dies wäre das erste |

Die letzte Zeile ist der Grund, warum die ADR nicht per Analogie argumentieren kann: im
Schwester-Repo steht der Scan in einer Klasse mit den Frische-Achsen. Hier gibt es diese Klasse
nicht; die Abgrenzung *„Netz ist der Zweck, nicht ein Zugeständnis"* muss für sich stehen.

### Drei Fallen, im Schwester-Repo gemessen — hier übernommen, nicht neu entdeckt

1. **`make` normalisiert jeden fehlgeschlagenen Recipe auf Exit 2.** Die Drei-Code-Semantik des
   Skripts (0 sauber / 1 behebbare Befunde / 2 gescheitert) überlebt die `make`-Grenze **nicht**.
   Der Workflow liest den Ausgang deshalb aus dem **Log**. Wer `rc` auswertet, meldet gemessene
   Befunde als „Scan gescheitert".
2. **Trivys `--exit-code` macht Fehler und Befund ununterscheidbar.** Ein nicht existierendes
   Image quittiert mit demselben Code wie ein Fund. Beide Läufe fahren darum `--exit-code 0`; über
   Befunde entscheidet die **Auswertung**, ein Nicht-Null-Exit heißt eindeutig „gescheitert".
3. **`grep -c` liefert 1, wenn nichts passt.** Ohne `|| true` risse der Zähl-Pfad den Lauf ab —
   und ein **sauberes** Image sähe aus wie ein Fehler.

### Warum das Target in `NICHT_PRUEFEND` gehört

Der PreToolUse-Guard verlangt, dass jedes deklarierte Prüf-Target in seiner GATES-Liste steht
(Drift-Wächter, slice-059). `image-scan` ist ein Prüf-Target — aber sein Ausgang wird
bestimmungsgemäß **weiterverarbeitet** (der Workflow liest das Log). Es gehört damit in dieselbe
Klasse wie `doc-repair`, dessen Ausgabe in `git apply` läuft. Die Entscheidung wird **begründet**
in die Liste geschrieben, nicht stillschweigend.

## 4. Trigger

**Start:** eingetreten — Maintainer-Auftrag, das Image existiert.

**Rückführungen:**

- `in-progress` → `next`: falls der erste Lauf so viele Befunde meldet, dass ihre Behebung ein
  eigener Slice ist. Dann liefert dieser Slice den Sensor, nicht die Sauberkeit.

## 5. Closure-Trigger

ADR steht, Sensor läuft mit Selbsttest, Workflow deklariert, Gates grün.

**Was bewusst nicht getan wird:** der **Hebungs-Kanal** (Dependabot) — das ist der Folge-Slice, und
die Reihenfolge ist Absicht: ein Kanal ohne Sensor hebt nichts Gemessenes. Ebenso **nicht**: den
Scan ins `gates`-Aggregat nehmen (er braucht Netz, `gates` ist hermetisch) und **nicht behebbare**
Befunde rot machen — ein Nachtlauf, der an unbehebbaren Basis-Image-CVEs rot wird, ist in zwei
Wochen ein weggeklicktes Abzeichen.

## 6. Risiken und offene Punkte

- *Der erste Lauf könnte Befunde melden, die dieser Slice nicht behebt* — **Ausgang:**
  eingetreten, **Folge-Slice**. Der Lauf fand **9 behebbare HIGH**, alle in der Go-`stdlib`
  `v1.26.4` des publizierten Image; die OS-Fläche (vier Debian-Pakete) meldete **0**. Dieser Slice
  liefert den Sensor, nicht die Sauberkeit — das stand so in §5, bevor gemessen wurde. Die
  Behebung ist eine Go-Hebung plus Release und damit ein eigener Vorgang.
- *Der Zähl-Pfad hängt an Trivys Template-Feldnamen* — **Ausgang:** weiter offen im
  **Beobachtungs-Register**: der Digest-Pin hält es still, solange er steht, und `--selftest`
  deckt die Auswertung, nicht die Feldnamen. Ein Trivy-Bump muss den Zähl-Pfad gegen einen echten
  Lauf prüfen, nicht gegen den Selbsttest.
- *`schedule` ist eine Abtastung, keine Zusage* — **Ausgang:** weiter offen im
  **Beobachtungs-Register**: GitHub schaltet zeitgesteuerte Workflows in inaktiven Repos nach 60
  Tagen ab, und ein still abgeschalteter Nachtlauf sieht aus wie ein grüner. `workflow_dispatch`
  mildert das (der Lauf ist von Hand auslösbar), hebt es aber nicht auf.

## 7. Closure-Notiz

**Geliefert:** `make image-scan` — Trivy `0.74.0` digest-gepinnt gegen
`ghcr.io/pt9912/a-check:latest`, mit netzlos prüfbarer Auswertung (`--selftest`, sieben
Fixtures), fail-closed bei leerer Prüfmenge, und dem Nachtlauf
[`image-scan.yml`](../../../../.github/workflows/image-scan.yml). Dazu
[ADR-0037](../../adr/0037-cve-scan-gegen-das-publizierte-image.md) für die drei Entscheidungen:
Netz als Zweck, außerhalb von `gates`, nur behebbare CRITICAL/HIGH entscheiden.

**Lerneintrag — Form: neuer Sensor.** *Ein Gate, das an einen **Commit** gebunden ist, kann eine
ganze Klasse von Zusagen prinzipiell nicht halten — die, deren Gegenstand sich **ohne** Commit
ändert.* Das war die Prämisse der ADR, und der erste Lauf hat sie **sofort** belegt: neun
behebbare HIGH im publizierten Image, während `make gates` über denselben Baum grün ist. Beides
ist gleichzeitig wahr und widerspricht sich nicht — sie messen verschiedene Gegenstände. *Weil*
a-check bis heute **kein** Netz-Gate hatte, ließ sich diese Lücke auch nicht per Analogie
schließen: die Abgrenzung „Netz ist der Zweck, nicht ein Zugeständnis" musste für sich stehen,
statt sich auf eine bestehende Klasse zu berufen.

**Vier beobachtbare Closure-Kriterien:**

1. **Der Sensor hat beim ersten Lauf gefunden, wofür er gebaut ist:** 9 behebbare HIGH in der
   Go-`stdlib` `v1.26.4`, OS-Fläche 0. Ein Sensor, dessen Erstlauf nichts findet, ist von einem
   toten nicht zu unterscheiden.
2. **Die `make`-Falle ist belegt, nicht behauptet:** das Skript endete mit **1** (behebbare
   Befunde), `make image-scan` mit **2**. Genau deshalb liest der Workflow den Ausgang aus dem
   **Log**; wer `rc` auswertete, meldete neun gemessene Befunde als „Scan gescheitert".
3. Der Selbsttest fährt beide Richtungen — vier Eingaben müssen zählen, eine darf **nicht**
   (`xx FINDING …`, Marker mitten in der Zeile). Ohne diese Probe wäre `grep FINDING` von
   `grep '^FINDING '` nicht zu unterscheiden.
4. Das Target steht in `NICHT_PRUEFEND` des Guard **mit Begründung**: sein Ausgang wird
   bestimmungsgemäß weiterverarbeitet. Ein Pipe-Verbot schützte hier einen Exit-Code, den niemand
   auswerten kann.

**Ein Fund am Rande, der beinahe durchgerutscht wäre:** mein erster Workflow trug
`actions/checkout@de0fac2e… # v5.0.0` — derselbe Digest wie in `ci.yml`, aber mit **falschem**
Tag-Kommentar (dort steht `v6.0.2`). Der Digest ist die Wahrheit, der Kommentar nur Lesehilfe;
falsch ist er trotzdem, und er sieht richtig aus. Gefunden hat es der Abgleich gegen die
bestehenden Workflows, kein Gate — a-check hat keine `workflow_pins`-Prüfung.

**Offene Risiken und ihr Ausgang:** der erste eingetreten (Folge-Slice), die anderen beiden weiter
offen im Register.

**Beobachtungs-Register:** `BEO-026` neu angelegt (CI-Schicht, 1×, Beleg slice-124): a-check prüft
die Form seiner Workflow-`uses:`-Einträge nicht — weder den Pin noch die Übereinstimmung von
Digest und Tag-Kommentar. Der falsche Kommentar oben ist der Beleg.

**Folge-Slices:** die **Go-Hebung auf 1.27** (Maintainer-Wort; `golang:1.27` ist verfügbar,
`go1.27.0`, Digest ermittelt) samt Release — sie behebt die neun Befunde. Danach der
**Hebungs-Kanal** (Dependabot), der künftige Fälle dieser Art ohne Handarbeit meldet.
## 8. Sub-Area-Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** berührt werden die **Gate-/Werkzeug-Schicht** (`tools/`,
`Makefile`, `.claude/`), die **CI-Schicht** (`.github/workflows/`) und mit den zwei
Deklarations-Tabellen der **Harness-Einstieg**.

**Vorgelagert — offene Beobachtungen sichten:** [`BEO-023`](../../../../docs/plan/planning/observations/BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf/observation.md) (ein Prüfer bleibt
ohne Gegenstand **oder ohne Aufruf** unkalibriert) trifft hier doppelt: der Sensor bekommt einen
`--selftest` **und** einen Workflow, der ihn fährt. Beides ist Antwort auf dieselbe Beobachtung.

Alle berührten Sub-Areas GF.
