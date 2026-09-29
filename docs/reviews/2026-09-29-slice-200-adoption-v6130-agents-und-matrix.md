# Review-Report: slice-200 — 2026-09-29

**Review-Art:** Code — adversarisches Diff-Review gegen Slice-Plan und
Baseline-Sollform (Modul 10): je Commit geprüft, die drei Umsetzungs-Aspekte
gegen die Sollform im neuen Stand, Gate-Behauptungen durch eigene Läufe und
Mutations-Proben belegt.

**Gegenstand:** Commits `14cf001..3900731` (vier Commits: `git mv` nach
`in-progress/`; `AGENTS.md` §5 WIP-Limit/Übergangs-Zahl + `.d-check.yml`
Matrix-Klassen; Entscheidungs-Dokumentation + slice-201 geschnitten +
Ruhe-Marker; Risiko-Ausgang + slice-201 Titel).

**Skill:** `.harness/skills/reviewer.md` @ `3900731` · <!-- d-check:ignore -->
**Modell:** glm-5.3-flash · **Datum:** 2026-09-29

**Eingangs-Kontext:**

- `docs/plan/planning/in-progress/slice-200-adoption-v6130-agents-und-matrix.md`
- `docs/plan/planning/open/slice-201-agents-md-index-tabelle.md`
- `harness/conventions/MR-025-referenzmatrix-grandfathering-v6130.md` §Offener Adoption-Aspekt
- `AGENTS.md` §3 Hard Rules, §5
- `v6.13.0` · `regelwerk/modul-05-planning-harness.md` §Trigger je Lifecycle-Übergang und WIP-Limit
- `v6.13.0` · `regelwerk/grundlagen-harness-dateien.md` §harness/conventions.md als Konventionsspeicher
- `v6.13.0` · `templates/.d-check.yml` (Klassen `welle`/`carveout`/`roadmap`, Slug-Slice-Token)
- `v6.13.0` · `templates/AGENTS.template.md` §5 (Index-Tabelle, Regel-Auslagerung)
- `v6.13.0` · `templates/docs/plan/planning/slice.template.md`
- Beobachtung `BEO-HARNESS/agents-md-hinkt-baseline-dod-item-hinterher` (`state.md`: `offen`)

---

## Findings

### F-1 — slice-201: DoD vollständig abgehakt, obwohl der Slice in `open/` liegt und keine der drei Leistungen erbracht ist

- `kategorie`: HIGH
- `quelle`: nachweislich falsche Tatsachenbehauptung (Skill HIGH-Liste); Ziel-Form
  `v6.13.0` · `templates/docs/plan/planning/slice.template.md` §2 — DoD-Items
  stehen als offene `[ ]`-Liste in der Vorlage
- `pfad`: `docs/plan/planning/open/slice-201-agents-md-index-tabelle.md:46-52`
- `befund`: Alle drei DoD-Items tragen `[x]`, während `AGENTS.md` §5 weiterhin die
  17-Bullet-Liste ist (keine Index-Tabelle), `harness/rules/` nicht existiert und
  kein Review-Report zu slice-201 vorliegt — der Slice selbst liegt in `open/`.
  Der DoD-Block behauptet damit einen Abschluss, der nicht stattgefunden hat; unter
  den fünf Slices in `open/`/`next/` ist slice-201 der einzige mit abgehaktem DoD.
- `verifizierbar`: ja — Sichtprüfung gegen die drei benannten Artefakte; ein Gate,
  das DoD-Häkchen in `open/` prüft, existiert nicht.
- `klasse`: DoD vor der Arbeit abgehakt

### F-2 — slice-200 §3.1: Zählung „15 Regeln" gegen `AGENTS.md` §5 (17)

- `kategorie`: HIGH
- `quelle`: nachweislich falsche Tatsachenbehauptung (Skill HIGH-Liste); Mess-Regel
  „Wer eine Menge zählt, zählt sie zweimal verschieden" (`AGENTS.md` §5)
- `pfad`: `docs/plan/planning/in-progress/slice-200-adoption-v6130-agents-und-matrix.md:74`
- `befund`: „Die Konversion fasst **15 Regeln**" — §5 trägt **17** Top-Level-Bullets,
  gezählt mit zwei unterschiedlich gebauten Zählern: anchorter Count (`^- ` im
  §5-Bereich) = 17; Aufzählung = 13 fett-markierte + 4 unformatierte = 17. Auch die
  Varianten „nur fett-markierte" (13) und „Sub-Regeln" (die drei Mess-Regeln einzeln
  gezählt: 19+) lösen die 15 nicht. Die Größen-Begründung („über die
  Vier-Regel-Größe hinausgewachsen", `AGENTS.md:73`) bleibt a fortiori bestehen —
  die gemessene Zahl stimmt nicht.
- `verifizierbar`: ja — Nachzählen gegen `AGENTS.md` §5; kein Gate (Urteil über die
  Zähl-Konstruktion, §3.7).
- `klasse`: Zählung deckt ihre Menge nicht

### F-3 — slice-201: Gate-Kriterium außerhalb der DoD-Liste

- `kategorie`: LOW
- `quelle`: Ziel-Form `v6.13.0` · `templates/docs/plan/planning/slice.template.md` §2 —
  Gate-Läufe sind DoD-Items
- `pfad`: `docs/plan/planning/open/slice-201-agents-md-index-tabelle.md:56`
- `befund`: „`make gates` und `make verify` grün." steht als Satz ohne Checkbox
  unterhalb der DoD-Liste — das Kriterium ist als DoD-Punkt nicht abhakbar und
  entzieht sich damit der Verifikations-Abnahme.
- `verifizierbar`: nein — Form-Urteil, kein Sensor.
- `klasse`: DoD-Kriterium unverhakt

### F-4 — slice-201: vier Closure-Pflichten in einem DoD-Item zusammengefasst

- `kategorie`: LOW
- `quelle`: Ziel-Form `v6.13.0` · `templates/docs/plan/planning/slice.template.md` §2 —
  Review, Closure-Notiz, Register-Fortschreibung, Risiko-Ausgänge als getrennte Items
- `pfad`: `docs/plan/planning/open/slice-201-agents-md-index-tabelle.md:52-54`
- `befund`: Ein einziges Item bündelt „Unabhängiger Review, Report; Closure-Notiz
  mit Lerneintrag; Register fortgeschritten; jedes Risiko trägt einen Ausgang" —
  damit ist keine der vier Pflichten einzeln abhakbar (und keine davon ist auf dem
  aktuellen Stand erfüllt, siehe F-1).
- `verifizierbar`: nein — Form-Urteil, kein Sensor.
- `klasse`: DoD-Kriterium unverhakt

### F-5 — Wellen-Attributionen (138/148/149/137/150) gegen den vendierten Stand nicht prüfbar

- `kategorie`: INFO
- `quelle`: —
- `pfad`: `docs/plan/planning/in-progress/slice-200-adoption-v6130-agents-und-matrix.md:27-33,67,74`
- `befund`: Die Herkunfts-Wellen der drei Aspekte lassen sich gegen den vendierten
  `v6.13.0`-Baum nicht verifizieren — der didaktik-freie Extrakt trägt keine
  Wellen-Nummern. Die Sachaussagen selbst verifizieren vollständig gegen die
  Ziel-Formen (siehe Negativbefunde); die Wellen-Nummern sind Provenienz-Label mit
  etablierter Praxis (`MR-025` zitiert Wellen-Nummern derselben Herkunft).
- `verifizierbar`: nein — nur gegen das Kurs-Repo (extern).
- `klasse`: Herkunfts-Attribution ohne vendierten Beleg

## Negativbefunde

| Bereich | Ergebnis |
|---|---|
| `AGENTS.md` §5 WIP-Limit-Formulierung gegen `v6.13.0` · `regelwerk/modul-05-planning-harness.md` §Trigger je Lifecycle-Übergang und WIP-Limit | geprüft, ohne Befund — „pro Lauf, nicht pro Rolle" samt per-Zweig-Unterscheidung über `Verantwortlich:` wortgetreu; die Hauptzweig-Folge („höchstens ein Slice in `in-progress/`") ist ausdrücklich gehalten und durch den Buffet-Satz derselben Baseline-Stelle getragen; kein „fünf"-Restbestand in `AGENTS.md`/`harness/README.md` |
| Übergangs-Zahl fünf → sechs gegen denselben Baseline-Abschnitt | geprüft, ohne Befund — die Baseline listet sechs Übergänge inkl. `open\|next→done` (Konsolidierung); `make slice-mv` unterstützt `TO=<open\|next\|in-progress\|done>` und fährt damit auch den sechsten |
| `.d-check.yml` Matrix-Erweiterung gegen `v6.13.0` · `templates/.d-check.yml` | geprüft, ohne Befund — Klassen `welle`/`carveout`/`roadmap` und sechs Regeln (`spec-straten\|adr → welle\|carveout\|roadmap`) exakt wie die Ziel-Form; Kommentar „ohne Token: nur über Links erreichbar" wortgleich |
| Numerischer Slice-Token `slice-\d{3}` gegen `MR-000` | geprüft, ohne Befund — `MR-000` deklariert `slice-NNN`; die Ziel-Form nutzt einen Slug-Token („die Kennung ist ein Slug"), die Abweichung ist schema-begründet; keine Slice-Kennung mit ≥ 4 Ziffern oder Slug-Form in lebenden Dateien (Stichprobe über `docs/` und Repo-Root, `--exclude-dir=.harness`) |
| Gate-Behauptungen des Ranges, eigene Ausführung | belegt — `make doc-check`: EXIT 0, „606 Datei(en) geprüft, 0 Befund(e)"; `make doc-planning`: EXIT 0, 0 Befunde; `make verify`: EXIT 0. **Geltungsbereich:** die vier Mutations-Proben decken `adr→welle`, `adr→carveout`, `adr→roadmap` (jeweils als Token-Fund) und `spec-straten→welle` (als Link-Fund) mit der jeweiligen `matrix-forbidden`-Meldung ab; `spec-straten→carveout`/`→roadmap` sind nicht einzeln probiert — sie teilen denselben Link-Mechanismus, der die `spec-straten`-Seite probe-rot belegt. Probe mit bloßem Token in `spec/spezifikation.md` trat daneben (EXIT 0): die `spec-straten`-Seite erkennt Kanten über Links, nicht über Tokens — d-check-Semantik, unverändert durch diesen Range |
| Ruhe-Marker (roadmap.md, `458f243`) | geprüft, ohne Befund — der Marker „Nichts in Arbeit" steht genau dann, wenn `in-progress/` leer ist; slice-200 liegt dort, der Entzug ist `modul-06`-gerecht, `make doc-planning` grün |
| §1-Abgrenzung und Plan-Treue des Ranges | geprüft, ohne Befund — der Range überschreitet die drei Aspekte nicht; die Auslagerung von Aspekt 2 als slice-201 entspricht dem §4-Rückführungs-Vorbehalt des Plans; alle vier Commits tragen IDs, die beiden `(planning)`-Scopes berühren ausschließlich `docs/plan/planning/` |
| slice-201 Kopf-Felder und Gliederung gegen die Ziel-Form | geprüft, ohne Befund — Lifecycle/Welle/Bezug/Spec-Stellen/Verantwortlich/Autor/Datum vollständig; §1–§8 in der Pflichtgliederung; Lerneintrag-Form benannt („geschärfte Regel"); Befunde nur bei der DoD-Form (F-3, F-4) und den abgehakten Häkchen (F-1) |

## Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 2 |
| MEDIUM | 0 |
| LOW | 2 |
| INFO | 1 |

**Finding-Klassen dieses Laufs:** DoD vor der Arbeit abgehakt · Zählung deckt ihre
Menge nicht · DoD-Kriterium unverhakt · Herkunfts-Attribution ohne vendierten Beleg

## Verdikt

**Merge-blockierend:** ja — für die Closure von slice-200. Beide HIGH-Findings
sind Plan-Defekte (DoD-Vorab-Häkchen im Nachfolge-Slice; falsche Zählung im
Entscheidungs-Beleg), keine Umsetzungs-Defekte: Die Umsetzung der drei Aspekte in
`AGENTS.md` und `.d-check.yml` hält die Sollform des neuen Standes in allen
geprüften Punkten ein, und alle behaupteten Gate-Läufe sind eigenständig grün
bestätigt bei gleichzeitig roter Gegenprobe. F-1 ist vor der Ausführung von
slice-201 aufzulösen (Häkchen öffnen, Ausführung abwarten), F-2 vor dem
Closure-Text von slice-200 (Zahl korrigieren oder Konstruktions-Scope benennen).

**Übergabe:** Findings gehen an den Implementer (Rückkante Review → Plan bei
Plan-Defekt); die **Finding-Klassen** gehen zusätzlich in die Slice-Closure §7
und von dort in den Zähler. Dieser Report selbst ist ein **Lauf-Beleg** (Audit:
dieser Diff, dieser Skill, dieses Modell, dieses Verdikt) — er wird über Läufe
hinweg nicht wieder gelesen.
