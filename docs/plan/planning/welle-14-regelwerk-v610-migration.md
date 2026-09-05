# Welle welle-14: Regelwerk-Migration `v6.0.0` → `v6.1.0`

**Lifecycle:** Diese Datei entsteht bei der **Eröffnung** der Welle und liegt
flach unter `docs/plan/planning/`; bei Closure wandert sie per `git mv` nach
`done/` (neben ihre `welle-14-results.md`). Der Zustand ist die
Verzeichnis-Position — kein Status-Feld.

**Zielmeilenstein:** kein Meilenstein-Bezug.

**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.
**Datum:** 2026-09-05.

---

## 1. Welle-Ziel

a-check auf den aktuellen Kurs-Stand `v6.1.0` heben: den Sprung `v6.0.0` →
`v6.1.0` messen und bewerten (slice-161), die dabei gefundenen echten
Nachzüge umsetzen, und die Stand-Deklaration an ihren drei Stellen
(`harness/conventions.md` §Baseline, `AGENTS.md` §1, `harness/README.md`
§Guides) auf `v6.1.0` bringen.

## 2. Trigger (Welle startet)

- `v6.1.0`-Release im Kurs-Repo veröffentlicht — bestätigt:
  `gh release view v6.1.0 --repo pt9912/ai-harness-course` (2026-09-05,
  Kurs-Welle 118).
- Maintainer hat die Migration angewiesen (dieses Gespräch, 2026-09-05).

## 3. Closure-Trigger (Welle schließt)

- Alle Etappen-Slices dieser Welle liegen in `done/`.
- Die Stand-Deklaration nennt an allen drei Stellen `v6.1.0`.
- `make gates` und `make verify` je Exit 0 auf dem finalen Stand — Ausgabe in
  eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.
- Ergebnis-Notiz `done/welle-14-results.md` geschrieben.

## 4. Slices in dieser Welle

| Slice | Titel | Bezug |
|---|---|---|
| slice-161 | Delta-Analyse `v6.0.0` → `v6.1.0` | — (reine Ist-Messung, keine Vertragsberührung) |
| slice-162 | Review-Pflicht/Rollenwechsel-Absatz in `AGENTS.md` §6 auf `v6.1.0`-Wortlaut zurückschneiden | slice-161 §4.4/§6 |

_Weitere Zeilen kommen hinzu, sobald die nächste Etappe eröffnet wird —
dieselbe Reihenfolge wie beim vorigen Sprung (slice-135 schlug vor,
slice-136/139/141/… setzten um)._

**Stand (beide Slices in `done/`) — offene Entscheidung für den nächsten
Schritt, dem Maintainer am Ende der Sitzung vom 2026-09-05 gestellt:**
Etappe A (Vendoring `.harness/baseline/v6.1.0/` + 3-Stellen-Pin-Bump +
Reviewer-Skill-Zeile in `harness/README.md`, slice-161 §6) direkt
eröffnen, oder zuerst einen vollständigen **Adaptions-Durchgang** (Etappe
B: alle 18 aktiven `MR`-Einträge einzeln gegen `v6.1.0` bewerten, nicht
nur die vom Diff berührten) fahren. Anlass: `slice-162` deckte auf, dass
ein diff-only-Vorgehen bestehende Adaptions-Einträge nicht gegen den
neuen Stand prüft — siehe
[`BEO-HARNESS/adaption-korrigiert-repo-aussage`](observations/BEO-HARNESS/adaption-korrigiert-repo-aussage/observation.md)
(2×) und
[`rueckbau-kandidat-ueberlebt-baseline-migration`](observations/BEO-HARNESS/rueckbau-kandidat-ueberlebt-baseline-migration/observation.md)
(1×).

## 5. Abhängigkeiten

- Keine — kein anderer offener Slice und keine andere offene Welle berührt
  `.harness/baseline/` oder die drei Stand-deklarierenden Stellen.

## 6. Out-of-Scope für diese Welle

- Zeitdokumente-Archivierung (Wellen-Closure-Schritt 4 der Baseline) — kein
  akuter Trigger; a-check hat mit `welle-14` die erste offene Welle seit
  `welle-13`, es gibt noch keinen wellenlosen Altbestand, der hier
  nachzuziehen wäre.
- Jede Etappe, die slice-161 nicht vorschlägt oder die der Maintainer nicht
  abnimmt.

## 7. Closure-Notiz

Ergebnis: <folgt bei Closure — Zeiger auf `welle-14-results.md`>
Zähler: <folgt bei Closure — Zeiger auf `../observations/`>
