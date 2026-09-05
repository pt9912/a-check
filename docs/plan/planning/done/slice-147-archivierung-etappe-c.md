# slice-147 — Etappe C: `welle-12`/`welle-13` real archiviert

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese Datei liegt — eines von
`open/`, `next/`, `in-progress/`, `done/`. Er wechselt nur durch `make slice-mv`
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).

**Bezug:** Etappe **C** aus [slice-143 §6](../done/slice-143-archivierung-delta-analyse.md#6-vorschlag-vier-etappen),
per Maintainer-Wort gezogen 2026-09-05, nach der Vorbedingung aus
[slice-146](../done/slice-146-welle-feld-nachtrag.md) (Welle-Feld nachgetragen) und der
Sensor-Vorprüfung aus [slice-145](../done/slice-145-archivierung-sensor-geltungsbereich.md).
[Roadmap](../in-progress/roadmap.md).

**Berührte Spec-Stellen:** — *(keine)* — Planungs-Harness-Pflege, kein Vertrags-Fakt.

**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.

**Autor:** Claude (Sonnet 5), im Auftrag des Maintainers. **Datum:** 2026-09-05.

---

## 1. Ziel

`make archive-wave WELLE=welle-12 APPLY=1` und `make archive-wave WELLE=welle-13 APPLY=1`
ausführen: 33 + 6 Slices und 9 Review-Reports in je ein `done/welle-NN/archiv.zip` verdichten,
Stubs schreiben, alte Volltexte löschen, repo-weite Verweise nachziehen — und den Repo-Bestand
danach wieder gate-grün verifizieren.

## 2. Erster Lauf — zwei Sensor-Lücken und ein Code-Bug, real gemessen

Der erste `APPLY=1`-Lauf (beide Wellen) endete jeweils mit Exit 0 und einer plausiblen Zusammen-
fassung — `make gates` danach brach an drei Stellen, keine davon in der Sieben-Fundstellen-Liste
aus [slice-145](../done/slice-145-archivierung-sensor-geltungsbereich.md):

1. **`ReadWelleField`-Bug (Code, `tools/archive-wave/collect.go`).** Die Funktion las das
   `**Welle:**`-Feld bis zur ersten Leerzeile — korrekt für `d-check`s Format, falsch für
   a-checks: dort reihen `**Welle:**`/`**Deckt:**`/`**Bezug:**` als EIN durchgehender Absatz ohne
   Leerzeile dazwischen (gemessen an `slice-069`, `slice-086`). Die Folgefelder landeten
   unverändert im Stub — inklusive darin enthaltener bare Kennungen und eines jetzt archivierten
   Review-Links.
2. **`ids`-Modul-Lücke (`.d-check.yml`).** Das `Hervorgegangen:`-Feld eines Stubs trägt bare
   `AC-*`/`ADR-*`-Kennungen aus dem archivierten Volltext, absichtlich unverlinkt (Baseline-
   Stub-Template) — die `ADR-\d{4}`- und `AC-(FA-[A-Z]+|QA)-\d+`-Regeln kannten keine Ausnahme
   dafür. `d-check` selbst hat diese Ausnahme längst (`.d-check.yml` dort, Modul `ids`,
   `exempt-paths: [..., "docs/plan/planning/done/welle-*/*.md"]`) — die eigene Vorprüfung hatte
   nicht dort nachgesehen.
3. **`ignore-refs` fehlte als Ventil.** Ein Nachtrags-Review
   (`2026-07-26-nachtrag-etappe-c-f2-zurueckgezogen.md`) zitiert einen jetzt archivierten
   anderen Review-Report — ein Review-Report ist Lauf-Beleg und wird nicht mehr editiert
   (dieselbe Begründung wie bei `d-check`s eigenem `ignore-refs`-Abschnitt).

Neuer Steering-Loop-Eintrag [`BEO-GATE`](../observations/BEO-GATE/archiv-sensor-vorpruefung-unvollstaendig/observation.md)
(§9).

## 3. Behoben — drei unabhängige Änderungen

- **`tools/archive-wave/collect.go`:** `ReadWelleField` bricht jetzt zusätzlich an der nächsten
  `**Feld:**`-Zeile ab, nicht nur an der Leerzeile. Test `TestReadWelleField_FolgefeldOhneLeerzeile`
  ergänzt (`tools/archive-wave/collect_test.go`) — reproduziert exakt den a-check-Fund, wäre mit
  der alten Fassung rot. `make archive-wave-test` grün.
- **`.d-check.yml`:** `exempt-paths` um `"docs/plan/planning/done/welle-*/*.md"` ergänzt bei den
  Regeln `ADR-\d{4}` und `AC-(FA-[A-Z]+|QA)-\d+` (Modul `ids`) — wortgleiche Übernahme aus
  `d-check`s eigener Konfiguration, hier für die a-check-eigenen Kennungsklassen. `MR-\d{3}` blieb
  unverändert: kein gemessener Fund einer bare `MR-*`-Kennung in einem Stub.
- **`.d-check.yml` `ignore-refs`:** neuer Abschnitt, ein Eintrag für den konkreten Fund aus §2.3.

## 4. Zweiter Lauf — Arbeitsbaum zurückgesetzt, sauber wiederholt

Vor dem Fix: `git restore --staged --worktree -- .` und `git clean -fd docs/plan/planning/done/`
— der Stand vor dem ersten `APPLY=1`-Lauf war unverändert erreichbar (nichts committet). Nach dem
Fix beide Archivierungen erneut mit `APPLY=1` gefahren, Exit 0 beide Male. Stichprobe
`slice-069`/`slice-086`: Stub trägt jetzt nur noch `Welle:`/`Archiviert mit:`/`Hervorgegangen:` —
keine Folgefeld-Reste mehr.

## 5. Ergebnis

- `docs/plan/planning/done/welle-12/`: `archiv.zip` (33 Slices + 9 Reviews, 42 Dateien, geprüft
  per `unzip -l`), 33 Stubs.
- `docs/plan/planning/done/welle-13/`: `archiv.zip` (6 Slices + 1 Welle-Plan, 7 Dateien), 6 Stubs
  + der Welle-Plan-Stub `welle-13-konsumenten-befunde.md` (Modul 6: die Welle-Plan-Datei wandert
  per `git mv` in ihr `done/<welle-id>/`, neben die Ergebnisnotiz).
- 9 Review-Reports aus `docs/reviews/` gelöscht (kein eigener Stub — Kanon-Vorgabe, sie leben nur
  noch im `welle-12`-Archiv).
- Repo-weite Verweis-Nachzüge in ~24 Dateien (`.harness/skills/`, `done/`-Slices,
  `welle-*-results.md`, `roadmap.md`, `next/README.md`, `harness/conventions.md`,
  drei `MR-*`-Dateien).

## 6. Auszuführende Gates

`make gates` und `make verify`, beide Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie in
eine Pipe.

## 7. Was bewusst nicht getan wird

- **Kein Nachtrag des `<manuell auszufuellen>`-Platzhalters bei `Geschlossen:`.** Präzedenz
  geprüft: `d-check` selbst lässt ihn in 222 Stub-Dateien unausgefüllt stehen — kein Regelverstoß,
  kein Nacharbeitsauftrag.
- **Etappe D (Altbestand + Reviews-Sammelarchiv) bleibt aus** — eigener Vorschlag in
  [slice-143 §6](../done/slice-143-archivierung-delta-analyse.md#6-vorschlag-vier-etappen), noch
  nicht vom Maintainer gezogen.

## 8. DoD

- [x] Beide Wellen real archiviert (`APPLY=1`), Ergebnis stichprobenartig gegen `archiv.zip` und
      Stub-Inhalt verifiziert (§5).
- [x] Zwei Sensor-Lücken und ein Code-Bug behoben, je mit Gegenprobe (Test bzw. Gate-Re-Lauf, §3).
- [x] `make gates` und `make verify` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft.

## 9. Closure-Notiz

**Geliefert:** `welle-12` und `welle-13` sind archiviert — 39 Slices und 9 Review-Reports als
`archiv.zip` verdichtet, 39 + 1 Stubs an ihrer Stelle, alle Verweise nachgezogen, Repo-Bestand
gate- und verify-grün.

**Lerneintrag — Form: neuer Sensor.** *Eine Sensor-Vorprüfung, die aus einer festen Liste
bekannter Fundstellen besteht (slice-145: sieben Stück), ist per Konstruktion unvollständig
gegenüber Modul-Klassen, die zum Prüfzeitpunkt nicht bedacht wurden — `ids` und `ignore-refs`
standen nicht auf der Liste, weil niemand nach ihnen gesucht hatte.* Der neue Sensor ist keine
neue Prüf-Logik, sondern zwei `exempt-paths`-Einträge, die exakt d-checks eigener, bereits
gelöster Konfiguration folgen ([`BEO-GATE`](../observations/BEO-GATE/archiv-sensor-vorpruefung-unvollstaendig/observation.md)).
*Weil* das Vorbild dieselbe Archivierungsform seit vielen Wellen produktiv betreibt, ist sein
`.d-check.yml` die zuverlässigere Quelle für „welche Modul-Klassen kollidieren mit einem Stub" als
eine von Hand geführte Fundstellen-Liste — ein Diff gegen `d-check`s Konfiguration wäre der
robustere erste Schritt einer künftigen Sensor-Vorprüfung, nicht das eigene Nachdenken allein.

**Zwei beobachtbare Closure-Kriterien:**

1. `unzip -l docs/plan/planning/done/welle-12/archiv.zip | tail -1` → `42 files`;
   `unzip -l docs/plan/planning/done/welle-13/archiv.zip | tail -1` → `7 files`.
2. `make gates` und `make verify` je **Exit 0** auf dem Stand dieses Slice, Ausgabe in Dateien,
   Exit-Codes getrennt geprüft.

**Offene Risiken und ihr Ausgang.** Eines notiert, ein Ausgang:
[`BEO-GATE`](../observations/BEO-GATE/archiv-sensor-vorpruefung-unvollstaendig/observation.md) —
**weiter offen** (Beobachtungs-Register, 1. Beleg): die Vorprüfungs-Methode selbst (feste Liste
statt Konfigurations-Diff gegen das Vorbild) ist nicht geändert, nur ihr aktueller Befund
nachgezogen. Ein künftiger `APPLY=1`-Lauf gegen eine neue Sensor-Kombination kann erneut an einer
nicht bedachten Stelle brechen.

**Folge-Slices:** keine vergeben. Etappe D bleibt aus
[slice-143 §6](../done/slice-143-archivierung-delta-analyse.md#6-vorschlag-vier-etappen)
vorgeschlagen und braucht eine eigene Kennung bei ihrer Eröffnung.

## 10. Sub-Area-Modus

Berührt zwei Sub-Areas:

- **Gate-/Werkzeug-Schicht** (`tools/archive-wave/`, `.d-check.yml`) — Greenfield: jede Änderung
  ist getestet (`archive-wave-test`) bzw. per Gate-Re-Lauf verifiziert.
- **Planungs-Harness** (`docs/plan/planning/done/`) — Greenfield: Form und Größenregel stehen in
  der Vorlage, `doc-structure` prüft sie.
