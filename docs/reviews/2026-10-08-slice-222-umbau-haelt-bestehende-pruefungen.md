# Review-Report: slice-222 — 2026-10-08

**Review-Art:** unabhängiger Lauf (frischer Kontext, kein `fork`) — Code-Review gegen den
Slice-Plan (§1b Messung M1–M4 und Plan-Änderung), `AGENTS.md` §3/§5 und die Mess-Regeln.

**Gegenstand:** Commit-Range `12114fa..4c07156` — `d831c82` (Messung und Plan-Änderung),
`4c07156` (Sensor `make pruefung-entfernt-check`, Regel 18, Hook, `preflight`, CI-Schritt).

**Skill:** `.harness/skills/reviewer.md` @ Stand `4c07156`
**Modell:** claude-opus-5-5 · **Datum:** 2026-10-08

**Eingangs-Kontext:**

- Slice-Plan slice-222 (`in-progress/`), Diff der Range
- `AGENTS.md` §3, §5; `harness/README.md` §Sensors; `harness/conventions.md`
- `.harness/skills/reviewer.md` inkl. §Mess-Regeln
- `v6.13.0` · `regelwerk/modul-05-planning-harness.md`, `regelwerk/modul-08-agentenrollen.md`

Alle Proben liefen in Wegwerf-Repos bzw. einem Klon im Scratchpad, nicht im Arbeitsbaum.

---

## Findings

### F-1 — Der Selbsttest im `commit-msg`-Hook schreibt in den Index des echten Repos

- `kategorie`: HIGH
- `quelle`: `harness/sensors/pruefung-entfernt-check.md` §Vertrag („Wegwerf-Repo im
  Temp-Verzeichnis"); `AGENTS.md` §5 Regel 18 (Hook als Aufrufer)
- `pfad`: `tools/pruefung-entfernt-check.sh:113-125` · `Makefile:126` · `.githooks/commit-msg:35`
- `befund`: Bei `git commit -a`, `git commit <pfad>` (Teil-Commit) und `--amend -a` setzt git
  für den Hook ein **absolutes** `GIT_INDEX_FILE` (`…/.git/index.lock` bzw.
  `next-index-*.lock`); das Make-Target fährt vor jedem Lauf `--selftest`, dessen
  `git add`/`git commit` im Wegwerf-Repo dann diesen fremden Index überschreiben. Gemessen im
  Klon mit dem echten Hook: ein `git commit -a`, der nur `README.md` ändert, wird abgelehnt mit
  113 angeblich entfernten Fehlerpunkt-Zeilen (der Index zeigt fast das ganze Repo als
  gelöscht); mit Trailer entsteht der Commit korrekt, aber der Index bleibt zerstört zurück —
  `git status` 1096 Einträge (`D  .a-check.yml`, …), `git fsck`: *invalid sha1 pointer in
  cache-tree of .git/index*, fehlende Blobs. Beim gestagten `git commit` (relatives
  `.git/index`) tritt es nicht auf; nur diesen Weg deckt die Live-Probe des Implementers.
- `verifizierbar`: ja — Wegwerf-Repo mit Hook, `git commit -a` mit und ohne Trailer, danach
  `git status`/`git fsck` (zweimal reproduziert: synthetischer Hook und echter Hook via `make`).
- `klasse`: Selbsttest im Hook-Kontext greift auf den Zustand des aufrufenden Repos zu

### F-2 — „Wortgleich wieder eingefügt" ist eine Mengen-Prüfung über alle Dateien: generische Zeilen maskieren eine entfernte Prüfung

- `kategorie`: MEDIUM
- `quelle`: Mess-Regel 5 (Zusage weiter als ihre Durchsetzung); Plan §1b („nicht verschobene
  Fehlerpunkt-Zeile")
- `pfad`: `tools/pruefung-entfernt-check.sh:41-46`; `harness/sensors/pruefung-entfernt-check.md`
  §Vertrag
- `befund`: Entfernt und hinzugefügt wird je als Menge getrimmter Zeilen über **alle** Dateien
  des Commits verglichen (`sort -u` + `comm`). Eine Prüfung, deren einziger Fehlerpunkt eine
  nackte `exit 1`-/`exit 2`-Zeile ist, gilt als „verschoben", sobald der Commit irgendwo in
  `tools/` oder `.github/workflows/` eine weitere nackte `exit 1` einfügt — genau das Bild
  eines Umbaus. Probe: Block `if b; then echo "b fehlt" >&2; exit 1; fi` aus `tools/x.sh`
  entfernt, in `tools/y.sh` ein neues `exit 1` → Exit 0, keine Meldung. Im Bestand stehen 46
  nackte `exit 1`/`exit 2`-Zeilen in den beiden Verzeichnissen. Vertrag und Regel-Datei nennen
  „dieselbe Zeile steht im selben Commit wieder da" ohne die Grenze, dass Ort und Anzahl nicht
  zählen.
- `verifizierbar`: ja — Wegwerf-Repo, zwei Commits, `RANGE=HEAD~1..HEAD`.
- `klasse`: Verschiebe-Erkennung per Zeilenmenge statt per Ort/Anzahl

### F-3 — Plan-Änderung nennt andere Namen und einen Aufrufer weniger als geliefert

- `kategorie`: MEDIUM
- `quelle`: `AGENTS.md` §6 Schritt 4 (Mitnahme außerhalb §1 ist Plan-Änderung vor dem Code);
  Plan §1b/§3
- `pfad`: Slice-Plan §1b „Plan-Änderung", §3 Tabelle
- `befund`: Der Plan (auch nach der Plan-Änderung in `d831c82`) nennt `make
  verify-pruefung-entfernt`, `tools/verify-pruefung-entfernt.sh` und
  `harness/sensors/verify-pruefung-entfernt.md`; geliefert sind `make pruefung-entfernt-check`,
  `tools/pruefung-entfernt-check.sh`, `harness/sensors/pruefung-entfernt-check.md`. Der Plan
  nennt als Aufrufer `preflight` und CI; geliefert ist zusätzlich der `commit-msg`-Hook
  (`.githooks/commit-msg`, nicht in §3) — der Aufrufer, an dem F-1 hängt. Ebenfalls nicht in §3:
  die Zählstellen-Edits in `docs/user/releasing.md`, `harness/README.md` und einem `state.md`
  des Registers.
- `verifizierbar`: ja — Plan-Text gegen `git show --stat 4c07156`.
- `klasse`: Plan nennt andere Artefakte als der Diff liefert

### F-4 — Lokal still grün bei `color.ui=always`

- `kategorie`: LOW
- `quelle`: Mess-Regel 5 (Grenze am Satz)
- `pfad`: `tools/pruefung-entfernt-check.sh:68,79`; Sensor-Datei §Grenze
- `befund`: `git show`/`git diff --cached` laufen ohne `--no-color`; mit `color.ui=always` in
  der Nutzer-Konfiguration beginnen die Diff-Zeilen mit ANSI-Sequenzen, kein Muster trifft.
  Probe: entfernte `fail "a"`-Zeile → Default Exit 1, mit `color.ui=always` Exit 0. Betroffen
  sind Hook und `preflight`; die CI nicht. Die Sensor-Datei nennt diese Grenze nicht.
- `verifizierbar`: ja — Wegwerf-Repo, `git config color.ui always`.
- `klasse`: Diff-Auswertung abhängig von Nutzer-Git-Konfiguration

### F-5 — Entfernte Zeilen, deren Inhalt mit `-` beginnt, sind unsichtbar

- `kategorie`: LOW
- `quelle`: Mess-Regel 4/5; Selbsttest-Fall „Diff-Kopf --- zaehlt nicht"
- `pfad`: `tools/pruefung-entfernt-check.sh:41,44,100-101`
- `befund`: Der Filter `^-[^-]` schließt nicht nur den Kopf `--- a/…` aus, sondern jede entfernte
  Zeile, deren Inhalt mit `-` beginnt (z. B. YAML-Listeneintrag in Spalte 0
  `- run: echo "::error::kaputt"`; Probe → Exit 0); dasselbe gilt für Zeilen eines
  kombinierten Merge-Diffs (`--`-Präfix). Der Selbsttest benennt nur den Kopf. Im Bestand
  trat der Fall nicht auf (zweiter Zähler ohne diese Einschränkung liefert dieselben 12).
- `verifizierbar`: ja — Wegwerf-Repo.
- `klasse`: Diff-Kopf-Filter breiter als der Diff-Kopf

### F-6 — Selbsttest-Fall „Regel-Anker in AGENTS.md" kann nie FAIL melden

- `kategorie`: LOW
- `quelle`: Mess-Regel 4 (Testbeschreibung weiter als Assertion)
- `pfad`: `tools/pruefung-entfernt-check.sh:109-110`
- `befund`: `grep -qF "$MARKER" AGENTS.md` steht als eigene Zeile unter `set -e`; fehlt der
  Anker, bricht der Selbsttest dort ab, die `t`-Zeile mit `"$?"` wird nie erreicht und kann nur
  `ok` drucken. Mutation (Anker in einer Kopie von `AGENTS.md` entfernt): Exit 1, Ausgabe endet
  nach „kein Trailer" ohne FAIL-Zeile und ohne Zusammenfassung. Rot ist es; die Meldung nennt
  den Grund nicht.
- `verifizierbar`: ja — Mutation im Scratchpad.
- `klasse`: Selbsttest-Fall ohne eigene Fehlermeldung (Abbruch statt FAIL-Zeile)

### F-7 — Hook-Kommentar „WARUM BEIDE HIER" nach Teilersetzung

- `kategorie`: LOW
- `quelle`: `AGENTS.md` §3.7
- `pfad`: `.githooks/commit-msg:13`
- `befund`: Kopf (Z. 2) und Z. 19 wurden auf drei Prüfungen umgestellt, die Überschrift Z. 13
  sagt weiterhin „BEIDE".
- `verifizierbar`: nein — Lesen.
- `klasse`: Zählwort nach Teilersetzung stehen geblieben

### F-8 — Bestandszahl an der falschen Grenze

- `kategorie`: LOW
- `quelle`: Mess-Regel 1 (Geltungsbereich einer Messung)
- `pfad`: `harness/sensors/pruefung-entfernt-check.md` §Grenze Punkt 3
- `befund`: „12 von 94 Werkzeug-Commits" steht an der Grenze *umformuliert zählt als entfernt*;
  die 12 sind alle Treffer (entfernt oder umformuliert), nicht die Umformulierungen. Die
  Ausgänge-Tabelle nennt 0/1/2; ein fehlschlagendes `git diff --cached` (F-1) endet mit 128.
- `verifizierbar`: ja — Zählung (s. Negativbefunde).
- `klasse`: Messwert an eine engere Aussage gehängt

### F-9 — Falsch-rot durch Teiltreffer

- `kategorie`: INFO
- `quelle`: Plan §1b (Kosten „jeder achte Werkzeug-Commit")
- `pfad`: `tools/pruefung-entfernt-check.sh:34`
- `befund`: `exit [12]` ist ein Teilstring-Muster: `exit 10`, `exit 127` und Kommentarzeilen
  wie `# bei Fehler exit 1` lösen aus (Proben: je Exit 1); ebenso eine unverändert aus `tools/`
  herausverschobene Datei. Die gemessenen 12/94 enthalten diese Fälle bereits; das ist Reibung,
  kein Durchrutschen.
- `verifizierbar`: ja — Wegwerf-Repo.
- `klasse`: Muster trifft Nicht-Fehlerpunkte

### F-10 — Pending-Modus sieht bei `--amend` nur das Delta zum alten Commit

- `kategorie`: INFO
- `quelle`: Sensor-Datei §Grenze Punkt 5
- `pfad`: `tools/pruefung-entfernt-check.sh:79`
- `befund`: Ein `git commit --amend -m …`, das den Trailer aus einem Commit mit Entfernung
  streicht, passiert den Hook (Probe G6); der CI-Range-Schritt fängt es. Die Grenze nennt die
  CI als klon-unabhängige Kontrolle, den Amend-Fall nicht.
- `verifizierbar`: ja — Wegwerf-Repo.
- `klasse`: Pending-Modus sieht nur Index gegen HEAD

### F-11 — CHANGELOG-Zeile fehlt

- `kategorie`: INFO
- `quelle`: `AGENTS.md` §6 Schritt 7 („… oder eine Regel")
- `pfad`: `CHANGELOG.md` §[Unreleased]
- `befund`: Regel 18 hat keinen Eintrag; die vier vorigen Commits unter `harness/rules/`
  (slice-215, slice-219) trugen ebenfalls keinen — Repo-Praxis liest „Regel" offenbar als
  Produkt-Regel. Zuständig ist die Einordnung durch Implementer/Planner.
- `verifizierbar`: nein.
- `klasse`: Reichweite von „eine Regel" in §6 Schritt 7 ungeklärt

## Negativbefunde

- geprüft, ohne Befund: Selbsttest — `bash tools/pruefung-entfernt-check.sh --selftest`, 15 Fälle
  `ok`, Exit 0 (selbst gefahren).
- geprüft, ohne Befund: Gegenprobe der drei Belege — Kopie des Skripts mit abgeschaltetem
  Grandfathering: `5131de5` (6 Zeilen, darunter „Config sagt"), `bd68b1c` (5, darunter
  „OCI-Label … ist leer"), `2236547` (5, darunter „Marker nur am Zeilenende") je Exit 1 mit der
  verlorenen Prüfung in der Ausgabe.
- geprüft, ohne Befund: M3 mit zweitem, anders gebautem Zähler (ein `awk`-Durchlauf über
  `git log -p`, Kopfzeilen per `--- `/`+++ ` statt `^-[^-]`): 12 Commits, 40 Zeilen — gleich
  der Sensor-Schleife je Commit. Grundmenge heute 95 statt 94, weil `4c07156` selbst
  hinzukommt; Merge-Commits im Bestand: 0.
- geprüft, ohne Befund: Range über die eigene Lieferung — `make pruefung-entfernt-check
  RANGE=12114fa..4c07156`, Exit 0.
- geprüft, ohne Befund: Pending-Modus beim gestagten Commit (ohne Trailer abgelehnt, mit Trailer
  entstanden) und beim ersten Commit eines leeren Repos (läuft, Exit 0).
- geprüft, ohne Befund: Umbenennung innerhalb `tools/` mit Inhaltsänderung — nur das Delta
  erscheint, kein Fehlalarm; Binärdateien — keine in den beiden Verzeichnissen.
- geprüft, ohne Befund: Grandfathering-Anker — `Entfernte-Pruefungen:` steht genau einmal in
  `AGENTS.md` (Regel-18-Zeile); sein Wegfall macht den Selbsttest und damit das Target rot
  (s. F-6 zur Meldung).
- geprüft, ohne Befund: Trailer-Erkennung — nur am Zeilenanfang, Kommentarzeilen (`#`) und
  Fließtext zählen nicht.
- geprüft, ohne Befund: CI-Workflow — Schritt im Range-Block nach `doc-immutable`, unter
  `set -euo pipefail`, vor `make ci`; `make doc-workflows` Exit 0.
- geprüft, ohne Befund: `make doc-targets`, `make doc-check`, `make doc-mentions`,
  `make doc-structure`, `make guard-selftest` — je Exit 0; Gate-Index-Zeile und
  Nicht-Gates-Absatz konsistent.
- geprüft, ohne Befund: Hard Rules §3.1–§3.6 — keine Inline-Suppression, kein Host-Toolchain-
  Aufruf, kein Spec-Stratum und keine ADR berührt, kein Move im Inhaltscommit.
- geprüft, ohne Befund: Zählstellen „drei Range-Schritte" — alle auf eine zahlfreie Form
  umgestellt (AGENTS §4, Zeiger statt Menge).

## Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 1 |
| MEDIUM | 2 |
| LOW | 5 |
| INFO | 3 |

**Finding-Klassen dieses Laufs:** Selbsttest im Hook-Kontext greift auf den Zustand des
aufrufenden Repos zu · Verschiebe-Erkennung per Zeilenmenge statt per Ort/Anzahl · Plan nennt
andere Artefakte als der Diff liefert · Diff-Auswertung abhängig von Nutzer-Git-Konfiguration ·
Diff-Kopf-Filter breiter als der Diff-Kopf · Selbsttest-Fall ohne eigene Fehlermeldung ·
Zählwort nach Teilersetzung stehen geblieben · Messwert an eine engere Aussage gehängt · Muster
trifft Nicht-Fehlerpunkte · Pending-Modus sieht nur Index gegen HEAD · Reichweite von „eine
Regel" in §6 Schritt 7 ungeklärt

## Verdikt

**Merge-blockierend: ja (F-1).** Der Sensor tut, was die Messung verspricht — die drei Belege
wären rot gewesen, die Bestandszahl hält einem zweiten Zähler stand. Aber der neue Aufrufer im
`commit-msg`-Hook blockiert jeden `git commit -a` und Teil-Commit in einem Klon mit `make hooks`
und hinterlässt, wenn der Commit durchgeht, einen zerstörten Index. F-2 und F-3 sind vor der
Closure zu klären. Kein Rollen-Widerspruch bisher; widerspricht der Implementer F-1, läuft der
Konflikt-Pfad über den Architect (`modul-08`).

**Übergabe:** an den Implementer. Dieser Report ist Lauf-Beleg, keine Verifikation — DoD-/Spec-
Konformität prüft der Verifier separat (Modul 11).
