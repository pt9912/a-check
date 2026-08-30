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

- [ ] Eine ADR begründet die drei Entscheidungen, die dieser Sensor trifft — **Netz als Zweck**
      (nicht als Zugeständnis), **nicht im `gates`-Aggregat**, und **nur behebbare CRITICAL/HIGH**
      entscheiden über rot. Sie steht im ADR-Index.
- [ ] `tools/image-scan.sh` + `make image-scan`: Trivy digest-gepinnt, Auswertung als **netzlos
      prüfbare** Funktion mit `--selftest`, fail-closed bei leerer Prüfmenge. Exit 0/1/2.
- [ ] `.github/workflows/image-scan.yml` fährt ihn zeitgesteuert **und** per `workflow_dispatch`;
      der Ausgang wird aus dem **Log** gelesen, nicht aus dem Exit-Code. Deklaration in
      [`AGENTS.md`](../../../../AGENTS.md) §4 und
      [`harness/README.md`](../../../../harness/README.md) §Sensors.

- [ ] `make gates` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.
- [ ] Closure-Notiz mit benanntem Lerneintrag geschrieben (§7).
- [ ] Beobachtungs-Register fortgeschrieben.
- [ ] Jedes Risiko aus §6 trägt genau einen Ausgang.

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

- *Der erste Lauf könnte Befunde melden, die dieser Slice nicht behebt* — **Ausgang:** <bei Closure>
- *Der Zähl-Pfad hängt an Trivys Template-Feldnamen; werden sie umbenannt, rendert das Template
  nichts und das sähe aus wie ein sauberes Bild* — der Digest-Pin hält das still, solange er steht.
  **Ausgang:** <bei Closure>
- *`schedule` ist eine Abtastung, keine Zusage — GitHub schaltet es in inaktiven Repos nach 60
  Tagen ab, und ein still abgeschalteter Nachtlauf sieht aus wie ein grüner* —
  **Ausgang:** <bei Closure>

## 7. Closure-Notiz

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice,
[`AGENTS.md`](../../../../AGENTS.md) §5.)_

**Lerneintrag — Form: <geschärfte Regel | neuer Sensor | benannte Spec-Lücke>**

## 8. Sub-Area-Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** berührt werden die **Gate-/Werkzeug-Schicht** (`tools/`,
`Makefile`, `.claude/`), die **CI-Schicht** (`.github/workflows/`) und mit den zwei
Deklarations-Tabellen der **Harness-Einstieg**.

**Vorgelagert — offene Beobachtungen sichten:** [`BEO-023`](../observations.md) (ein Prüfer bleibt
ohne Gegenstand **oder ohne Aufruf** unkalibriert) trifft hier doppelt: der Sensor bekommt einen
`--selftest` **und** einen Workflow, der ihn fährt. Beides ist Antwort auf dieselbe Beobachtung.

Alle berührten Sub-Areas GF.
