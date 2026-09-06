# Review-Report: slice-166 — 2026-09-06

**Review-Art:** Plan — geprüft gegen den Slice-Plan und die Konventions-Änderung
selbst (Modul 10 §Drei Review-Arten): der Gegenstand ist eine
Harness-Konventions-Änderung (kein Verhaltenscode), Kern-Behauptungen sind
eine Zählung repo-weiter Referenzen und die Einhaltung von `AGENTS.md`
§3.3 (Move + Inhaltsänderung = zwei Commits).

**Gegenstand:** Commits `c201058` ("docs(harness): slice-166 schneiden --
MR-017 durch MR-020 abgeloest") und `adb9534` ("MR-017s interne Links nach
dem Umzug ... korrigiert")

**Skill:** `.harness/skills/reviewer.md` @ Stand `adb9534` (unverändert
seit Anlage) · <!-- d-check:ignore -->
**Modell:** claude-sonnet-5 · **Datum:** 2026-09-06

**Eingangs-Kontext:**

- `docs/plan/planning/done/slice-166-mr020-adr-vorlage-generisch.md`
- `harness/conventions/MR-020-adr-vorlage-generisch.md`
- `harness/conventions/done/MR-017-adr-vorlagen-version.md`
- `harness/conventions.md` (beide Adaptions-Tabellen)
- `docs/plan/planning/observations/BEO-HARNESS/rueckbau-kandidat-ueberlebt-baseline-migration/`
- `git log`/`git show` gegen `c201058`, `adb9534` und deren Elternstände

---

## Findings

### F-1 — Kern-Zahl „zwölf Dateien" ist falsch, tatsächlich zehn

- `kategorie`: HIGH
- `quelle`: eigene Nachzählung (`git grep -l
  "conventions/MR-017-adr-vorlagen-version.md" c201058^ -- '*.md'`, außerhalb
  `.harness/baseline/`)
- `pfad`: Commit-Message `c201058` sowie
  `docs/plan/planning/done/slice-166-mr020-adr-vorlage-generisch.md` §3, §6,
  §7 (mehrfache Wiederholung derselben falschen Zahl, erste Fassung)
- `befund`: die Behauptung „zwölf Dateien" (als „gemessen, nicht
  angenommen" formuliert) ist falsch — die tatsächliche, nachgezählte Zahl
  ist **zehn**: `slice-163`, `slice-164`, `slice-165`,
  `slice-141-mr013-nachfolger`, zwei Beobachtungs-Evidence-Dateien, ein
  Beobachtungs-`observation.md`, `welle-14-…md`, ein Review-Report,
  `harness/conventions.md`. Nach dem Fix ist repo-weit tatsächlich kein
  alter Pfad mehr verlinkt (das Ziel selbst wurde erreicht), aber die
  Zähl-Behauptung verletzt dieselbe Sorgfalt, die `AGENTS.md` für CR-Texte
  einfordert — hier bei einer repo-eigenen Zählung.
- `verifizierbar`: ja — `git grep`-Befehl oben, exakt reproduzierbar.
- `klasse`: Repo-eigene Zählung behauptet statt nachgezählt

## Negativbefunde

- geprüft, ohne Befund: `AGENTS.md` §3.3 (Move + Inhaltsänderung = zwei
  Commits) eingehalten — `c201058` ist ein reiner 100 %-Rename
  (`{ => done}/MR-017-…md | 0`, keine Inhaltsänderung im selben Commit),
  der Link-Fix folgt als eigener Commit `adb9534`.
- geprüft, ohne Befund: Inhalt/Aussage von `MR-017` vor/nach dem Umzug
  Wort für Wort identisch — der Diff zeigt ausschließlich
  `../conventions.md` → `../../conventions.md` und einen zusätzlichen
  `../` beim `slice-135`-Link; keine inhaltliche Verletzung der
  ADR-/MR-Immutabilität.
- geprüft, ohne Befund: alle relativen Links in `MR-017` lösen korrekt
  auf — `../../conventions.md#mr-013`/`#mr-000` von
  `harness/conventions/done/MR-017-…md` zeigt auf `harness/conventions.md`
  (existiert), der `slice-135`-Link zeigt auf die existierende Datei.
- geprüft, ohne Befund: `harness/conventions.md`s beide Tabellen korrekt
  gepflegt — `MR-017` nur noch in „Aufgelöste Adaptionen" (aufgelöst durch
  `MR-020`), `MR-020` korrekt in „Aktive Adaptionen"; beide Anker
  (`id="mr-017"`, `id="mr-020"`) genau einmal vorhanden, keine Duplikate.
- geprüft, ohne Befund: `MR-020`s Pflichtfelder vollständig und
  formkonsistent zu `MR-017`/`MR-018`/`MR-019`; das Fehlen von
  „Rückbau-Kandidat, ausgewiesen" ist korrekt, da `MR-020` explizit als
  „permanent, kein Rückbau-Kandidat" deklariert ist.
- geprüft, ohne Befund:
  `BEO-HARNESS/rueckbau-kandidat-ueberlebt-baseline-migration/state.md`
  trägt jetzt `verkörpert` mit auflösbarem Link auf `MR-020` und Anker
  `seit slice-166`; Form entspricht bereits zweimal verwendeten
  Präzedenzmustern im selben Register.
- geprüft, ohne Befund: Kopf-Metadaten, DoD-Größe, Trigger-/
  Closure-Trigger-/Risiko-Abschnitte und Closure-Notiz-Form strukturell
  konsistent zu `slice-164`/`slice-165`; beide Risiken in §6 tragen
  gültige Ausgänge aus der geschlossenen Dreier-Menge. Die Behauptung „9
  offen vor diesem Slice" in §8 bestätigt (8 offen danach + 1 gerade
  verkörpert = 9 vorher).

## Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 1 |
| MEDIUM | 0 |
| LOW | 0 |
| INFO | 0 |

**Finding-Klassen dieses Laufs:** Repo-eigene Zählung behauptet statt
nachgezählt

## Verdikt

**Merge-blockierend:** nein mehr — F-1 wurde vor Abschluss von `slice-166`
im Slice-Text korrigiert (zehn statt zwölf). Die Commit-Message von
`c201058` trägt die falsche Zahl weiter fort — kein Amend bereits
erstellter Commits in diesem Repo ohne expliziten Anlass; die Korrektur
steht stattdessen sichtbar in der Closure-Notiz (§7). Das eigentliche
Ziel des Slice (repo-weiter Referenz-Nachzug vollständig, `§3.3` gewahrt,
Adaptions-Tabellen korrekt) ist unabhängig von der falschen Zahl erreicht.

**Übergabe:** Findings sind bereits in `slice-166`s Closure eingearbeitet.
Dieser Report ist Lauf-Beleg, keine Verifikation — DoD-/Spec-Konformität
prüft der Verifier separat (Modul 11).
