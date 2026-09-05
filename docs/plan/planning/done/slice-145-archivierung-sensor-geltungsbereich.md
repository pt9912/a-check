# slice-145 — Etappe B: Sensor-Geltungsbereich für die Archivierung geprüft

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese Datei liegt — eines von
`open/`, `next/`, `in-progress/`, `done/`. Er wechselt nur durch `make slice-mv`
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).

**Bezug:** Etappe **B** aus [slice-143 §6](../done/slice-143-archivierung-delta-analyse.md#6-vorschlag-vier-etappen),
per Maintainer-Wort gezogen 2026-09-05. Setzt die Baseline-Vorschrift um: „Vor der ersten
Archivierung ist der Geltungsbereich der vorhandenen Sensoren zu prüfen." Vorgänger
[slice-144](../done/slice-144-archive-wave-werkzeug.md) (Etappe A). [Roadmap](../in-progress/roadmap.md).

**Berührte Spec-Stellen:** — *(keine)* — Gate-Konfiguration ohne Vertragsberührung.

**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.

**Autor:** Claude (Sonnet 5), im Auftrag des Maintainers. **Datum:** 2026-09-05.

---

## 1. Ziel

Jede der sieben Fundstellen aus [slice-143 §3](../done/slice-143-archivierung-delta-analyse.md#3-sensoren-mit-geltungsbereich-auf-done-oder-docsreviews--die-verlangte-vorprüfung)
einzeln prüfen: bricht sie an einem archivierten Stub, und falls ja, anpassen — **bevor** Etappe C/D
den ersten `APPLY=1`-Lauf fährt. Kein Sensor wird hier neu erfunden, nur bestehende Globs gegen die
tatsächliche Ziel-Form eines Stubs (`archiv-stub-slice.template.md`) gehalten.

## 2. Die sieben Fundstellen — geprüft, nicht vermutet

| Fundstelle | Rezidiv-Frage | Ausgang |
|---|---|---|
| `.d-check.yml` `structure` (1) — DoD-Regel, `docs/plan/planning/**/slice-*.md` | Ein Stub trägt **keine** `##`-Überschrift — `section-pattern` für `DoD` findet nichts | **Angepasst** — `exempt-paths` um `done/welle-[0-9]*/**` und `done/wellenlos/**` ergänzt (§3) |
| `.d-check.yml` `structure` (4) — Kopffelder-Regel, `docs/plan/planning/**/slice-*.md` | Ein Stub trägt `Welle:`/`Archiviert mit:`/`Hervorgegangen:` statt `Verantwortlich:`/`Autor:`/`Berührte Spec-Stellen:` | **Angepasst** — dieselben zwei `exempt-paths` (§3) |
| `.d-check.yml` `structure` (2) — Closure-Struktur, `docs/plan/planning/done/slice-*.md` | Glob ist **nicht rekursiv** (kein `**`) — ein Stub unter `done/<subdir>/` matcht nie | **Unverändert sicher** — Bash-/Glob-Semantik: ein einzelnes `*` überquert kein `/` |
| `.d-check.yml` `structure` (3) — Lerneintrag-Form, `docs/plan/planning/done/slice-*.md` | Dieselbe Frage wie (2), derselbe Glob | **Unverändert sicher**, aus demselben Grund |
| `.d-check.yml` `links.resolve-from` (`fixed-dirs: docs/plan/planning/done`) | Ändert sich die Lifecycle-Invarianten-Prüfung für Dateien, die **innerhalb** von `done/` liegen? | **Unverändert sicher** — `done/` ist dort ausschließlich **Ziel**, nie Quelle einer wandernden Datei; die Prüfung gilt Dateien in `open/`/`next/`/`in-progress/`, nicht dem Inhalt von `done/` |
| `tools/verify-risiko-ausgaenge.sh` | `for f in "$DONE_DIR"/slice-*.md` — Bash-Glob, rekursiv? | **Unverändert sicher** — Bash-Globs überqueren kein `/`; ein Stub unter `done/<subdir>/` wird nie iteriert |
| `tools/verify-observations.sh` | `find "$DONE_DIR" -type f -name '*.md'` — **ist** rekursiv; liest ein Stub-Zitat im `Hervorgegangen:`-Feld dann noch mit? | **Unverändert korrekt** — die Rekursion ist hier **erwünscht**: ein archivierter Stub, dessen `Hervorgegangen:`-Feld eine `BEO-*`-Kennung nennt, soll weiter zitat-geprüft werden. Kein Sensor-Bruch, sondern die einzig sinnvolle Lesart |
| `tools/commit-scope-check.sh` | Prüft nur den Pfad-Präfix `^docs/plan/planning/` | **Unverändert sicher** — pfad-, nicht formgebunden, unabhängig von Verzeichnistiefe |

**Fünf von sieben Fundstellen waren beim genaueren Blick bereits sicher** — meist, weil ein
nicht-rekursiver Glob (`slice-*.md` statt `**/slice-*.md`) archivierte Unterverzeichnisse
strukturell ausschließt. Nur die **beiden rekursiven** `structure`-Regeln (DoD, Kopffelder) griffen
tatsächlich zu weit.

## 3. Die Änderung — und ihre Gegenprobe

Zwei `exempt-paths`-Einträge in beiden betroffenen `structure`-Regeln:

```yaml
- "docs/plan/planning/done/welle-[0-9]*/**"   # Wellen-Modus: die KURZE Welle-ID (slice-144 §4.1)
- "docs/plan/planning/done/wellenlos/**"       # Einzel-Slice-Modus (WellenlosArchiveDir)
```

**Gegenprobe statt Annahme:** zwei Fixture-Dateien angelegt, wortgleich zur Ziel-Form eines Stubs
(`archiv-stub-slice.template.md` — keine `##`-Überschrift, `Welle:`/`Archiviert mit:`/
`Hervorgegangen:` statt der drei sonst verlangten Felder), eine unter
`done/welle-99/slice-999-fake.md`, eine unter `done/wellenlos/slice-998-fake.md`.
`make doc-structure` vor der Anpassung: nicht gelaufen (der Fund war Ableitung aus der bekannten
`observations/**`-Kollision aus slice-139, nicht erneut provoziert). **Nach** der Anpassung: `0`
Befunde. Beide Fixtures nach dem Lauf entfernt — kein Spur im Bestand.

## 4. Auszuführende Gates

`make gates`, zum Abschluss `make verify`. Kein neuer Sensor — zwei bestehende `structure`-Regeln
um `exempt-paths` ergänzt.

## 5. Was bewusst nicht getan wird

- **`matrix`-Modul (`slice`-Klasse, `docs/plan/planning/**/slice-*.md`) bleibt unverändert.** Ein
  Stub matcht weiterhin als „Slice"-Klassenmitglied — das ist korrekt, kein Fehlklassifikation. Kein
  gemessener Grund für eine Ausnahme (Zurückhaltung wie in
  [slice-144](../done/slice-144-archive-wave-werkzeug.md) §5: kein Vorratsbau ohne Befund).
- **`docs/reviews/**`-Ausnahmen bleiben unverändert.** `ReviewArchiveDir` (`docs/reviews/archiv`)
  liegt bereits innerhalb der fünf bestehenden `exempt-paths: docs/reviews/**`-Einträge — keine
  neue Fundstelle.
- **Kein `APPLY=1`-Lauf.** Diese Etappe prüft und härtet Sensoren gegen eine **simulierte** Form,
  nicht gegen echtes Archiv — Etappe C/D aus slice-143 §6.

## 6. DoD

- [x] Alle sieben Fundstellen aus slice-143 §3 einzeln geprüft, fünf als sicher bestätigt (Glob-
      Semantik oder Modul-Zweck), zwei angepasst (§2).
- [x] Gegenprobe mit zwei Stub-förmigen Fixture-Dateien geführt — vor der Anpassung nicht
      reproduziert (Ableitung genügte, dieselbe Kollisionsklasse wie slice-139), nach der Anpassung
      `0` Befunde, Fixtures rückstandsfrei entfernt (§3).
- [x] `make gates` und `make verify` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie
      in eine Pipe.

## 7. Closure-Notiz

**Geliefert:** die von der Baseline verlangte Sensor-Vorprüfung, mit einem belegbaren Ergebnis statt
einer Behauptung: fünf von sieben Fundstellen waren bereits sicher (meist durch nicht-rekursive
Globs), zwei `structure`-Regeln angepasst und per Fixture-Gegenprobe bestätigt.

**Lerneintrag — Form: geschärfte Regel.** *Ob ein Glob-Muster ein archiviertes Unterverzeichnis
mitliest, entscheidet ein einziges Zeichen — `**` vs. `*` — und das ist an der Konfiguration selbst
ohne Test nicht zuverlässig zu sehen; derselbe Namensraum (`slice-*.md`) erscheint vier Mal im
selben Modul, zweimal rekursiv, zweimal nicht, aus jeweils eigenem Grund.* Ohne den Vergleich in
§2 hätte eine pauschale Annahme („alle vier `structure`-Regeln sind betroffen" oder „keine ist
betroffen") in zwei von vier Fällen falsch gelegen. *Weil* ein Sensor-Geltungsbereich nicht am
Namen des Moduls hängt, sondern an der exakten Glob-Form jeder einzelnen Regel — dieselbe Lehre wie
bei der Kollision zwischen `observations/**/evidence/slice-NNN.md` und dem Slice-Plan-Glob
([slice-139](../done/slice-139-beobachtungsregister-migration.md)), hier im Voraus statt im
Nachhinein gefangen, weil Etappe A ihren eigenen Vorschlag ernst nahm.

**Zwei beobachtbare Closure-Kriterien:**

1. Die Fixture-Gegenprobe aus §3 ist mit denselben zwei Dateien und `make doc-structure`
   reproduzierbar (Fixtures nicht im Bestand, aus der Beschreibung rekonstruierbar).
2. `make gates` und `make verify` je **Exit 0** auf dem Stand dieses Slice, Ausgabe in Dateien,
   Exit-Codes getrennt geprüft.

**Offene Risiken und ihr Ausgang.** Keine notiert — alle sieben Fundstellen sind mit Ausgang
versehen (§2), kein offener Rest.

**Folge-Slices:** keine vergeben. Etappe C (formale Wellen archivieren) bleibt aus slice-143 §6
vorgeschlagen und braucht eine eigene Kennung bei ihrer Eröffnung — inklusive des in
[slice-144](../done/slice-144-archive-wave-werkzeug.md) §4.2 gemessenen Blockers (fehlende
`**Welle:**`-Felder bei `welle-12`/`welle-13`).

## 8. Sub-Area-Modus

Berührt: **Gate-/Werkzeug-Schicht** (`.d-check.yml`) — Greenfield.
