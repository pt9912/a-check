# Welle welle-14: Regelwerk-Migration `v6.0.0` → `v6.2.0`

> **Retarget 2026-09-06:** ursprünglich auf `v6.1.0` eröffnet (Dateiname
> trägt das noch — stabile Kennung, kein Nachzug nötig, dieselbe Praxis wie
> bei Slice-Dateinamen). Noch bevor Etappe A vendorte, erschien `v6.2.0`
> (Kurs-Welle 119, 2026-09-05, 2h06m nach `v6.1.0`) — Vendoring von
> `v6.1.0` jetzt hätte es Stunden später durch `v6.2.0` ersetzen müssen.
> Ziel und Trigger unten auf `v6.2.0` gehoben, bevor Etappe A beginnt.

**Lifecycle:** Diese Datei entsteht bei der **Eröffnung** der Welle und liegt
flach unter `docs/plan/planning/`; bei Closure wandert sie per `git mv` nach
`done/` (neben ihre `welle-14-results.md`). Der Zustand ist die
Verzeichnis-Position — kein Status-Feld.

**Zielmeilenstein:** kein Meilenstein-Bezug.

**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.
**Datum:** 2026-09-05.

---

## 1. Welle-Ziel

a-check auf den aktuellen Kurs-Stand `v6.2.0` heben: den Sprung `v6.0.0` →
`v6.1.0` messen und bewerten (slice-161), den Adaptions-Bestand gegen
`v6.1.0` durchgehen (slice-163), das Increment `v6.1.0` → `v6.2.0`
zusätzlich messen (Folge-Slice, s. §4), die dabei gefundenen echten
Nachzüge umsetzen, und die Stand-Deklaration an ihren drei Stellen
(`harness/conventions.md` §Baseline, `AGENTS.md` §1, `harness/README.md`
§Guides) auf `v6.2.0` bringen.

## 2. Trigger (Welle startet)

- `v6.1.0`-Release im Kurs-Repo veröffentlicht — bestätigt:
  `gh release view v6.1.0 --repo pt9912/ai-harness-course` (2026-09-05,
  Kurs-Welle 118).
- `v6.2.0`-Release im Kurs-Repo veröffentlicht — bestätigt:
  `gh release view v6.2.0 --repo pt9912/ai-harness-course` (2026-09-05,
  Kurs-Welle 119, 2h06m nach `v6.1.0`).
- Maintainer hat die Migration angewiesen (dieses Gespräch, 2026-09-05);
  Retarget auf `v6.2.0` ebenso vom Maintainer angewiesen (2026-09-06).

## 3. Closure-Trigger (Welle schließt)

- Alle Etappen-Slices dieser Welle liegen in `done/`.
- Die Stand-Deklaration nennt an allen drei Stellen `v6.2.0`.
- `make gates` und `make verify` je Exit 0 auf dem finalen Stand — Ausgabe in
  eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.
- Ergebnis-Notiz `done/welle-14-results.md` geschrieben.

## 4. Slices in dieser Welle

| Slice | Titel | Bezug |
|---|---|---|
| slice-161 | Delta-Analyse `v6.0.0` → `v6.1.0` | — (reine Ist-Messung, keine Vertragsberührung) |
| slice-162 | Review-Pflicht/Rollenwechsel-Absatz in `AGENTS.md` §6 auf `v6.1.0`-Wortlaut zurückschneiden | slice-161 §4.4/§6 |
| slice-163 | Adaptions-Durchgang (Etappe B): alle 18 MR-Dateien gegen `v6.1.0` geprüft | Maintainer-Entscheidung 2026-09-06 ("Etappe B zuerst") |
| slice-164 | Delta-Analyse `v6.1.0` → `v6.2.0` (Increment) | Maintainer-Hinweis 2026-09-06 ("neues Release v6.2.0") |
| slice-165 | Review-Checkbox-Punkt bleibt Opt-in, [MR-019](../../../../harness/conventions/MR-019-review-dod-opt-in.md) | slice-164 §4.2/§5 (Folge-Slice-Vorschlag), Maintainer-Wort 2026-09-06 |
| slice-166 | [MR-017](../../../../harness/conventions/done/MR-017-adr-vorlagen-version.md) durch generische ADR-Vorlagen-Referenz abgelöst, [MR-020](../../../../harness/conventions/MR-020-adr-vorlage-generisch.md) | slice-163 §4/§6 (Folge-Slice-Vorschlag), Maintainer-Wort 2026-09-06 |

**Stand (alle sechs Slices in `done/`, Welle auf `v6.2.0` retargeted) —
keine offenen Folge-Slices mehr:** beide aus `slice-163`/`slice-164`
vorgeschlagenen Folge-Slices sind erledigt
([MR-020](../../../../harness/conventions/MR-020-adr-vorlage-generisch.md)
bzw. [MR-019](../../../../harness/conventions/MR-019-review-dod-opt-in.md)).
Nächster und letzter Schritt vor der Welle-Closure: **Etappe A**
(Vendoring `.harness/baseline/v6.2.0/` + 3-Stellen-Pin-Bump +
Reviewer-Skill-Zeile in `harness/README.md`, slice-161 §6, Ziel-Version auf
`v6.2.0` gehoben) — noch keine Slice-ID vergeben.

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
