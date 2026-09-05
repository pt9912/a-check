# Review-Report: slice-162 — 2026-09-05

**Review-Art:** Code — geprüft gegen Plan/Konventionen (Modul 10 §Drei
Review-Arten): der Gegenstand ist ein Doku-Rückschnitt in `AGENTS.md` §6
plus eine neue `MR`-Adaption, kein Verhaltenscode.

**Gegenstand:** Commit `2488705a6e452eacd5f0a1715a1438dc504f778f`
("build(harness): slice-162 schneiden -- Review-Pflicht-Absatz in
AGENTS.md §6 auf v6.1.0-Wortlaut, MR-018")

**Skill:** `.harness/skills/reviewer.md` @ Stand `2488705` (unverändert seit
Anlage) · <!-- d-check:ignore -->
**Modell:** claude-sonnet-5 · **Datum:** 2026-09-05

**Eingangs-Kontext:**

- `docs/plan/planning/done/slice-162-review-pflicht-absatz-v610-wortlaut.md`
- `docs/plan/planning/done/slice-161-regelwerk-v610-delta-analyse.md` §4.4, §6
- `docs/reviews/2026-09-05-slice-161-delta-analyse.md`
- `harness/conventions/MR-016-validator-unbesetzt.md`,
  `MR-017-adr-vorlagen-version.md` (Formvorbild)
- `AGENTS.md` §3 (Hard Rules, insb. §3.7), §6
- `.harness/baseline/v6.0.0/regelwerk/modul-08-agentenrollen.md`
- Kurs-Repo `pt9912/ai-harness-course`, Tags `v6.0.0`/`v6.1.0` (extern geklont
  zur unabhängigen Nachprüfung)

---

## Findings

### F-1 — MR-018-Begründung überzeichnet Deckung eines Details in zwei benannten Dateien

- `kategorie`: MEDIUM
- `quelle`: `harness/conventions/MR-018-review-pflicht-v610-wortlaut.md` (Begründung), Commit-Message
- `pfad`: `harness/conventions/MR-018-review-pflicht-v610-wortlaut.md:16–19`; `.harness/skills/reviewer.md:62–63`; `docs/reviews/README.md` (gesamt)
- `befund`: MR-018 und die Commit-Message behaupten, das entfernte
  „Report-Dateiname-Muster" (`docs/reviews/<YYYY-MM-DD>-<slice-oder-diff-ref>.md`)
  stehe bereits in `docs/reviews/README.md` **und** `.harness/skills/reviewer.md`.
  Geprüft: `.harness/skills/reviewer.md` trägt nur die Substanz („ein Report pro
  Lauf … Folgeläufe als neue Datei", Zeile 63) — das konkrete Namensmuster mit
  Platzhaltern steht wörtlich **nur** in
  `.harness/baseline/v6.0.0/templates/docs/reviews/review-report.template.md`
  (Zeile 5), einer dritten, nicht genannten Datei. `docs/reviews/README.md`
  erwähnt das Muster gar nicht, nur die Kopf-Metadaten-Pflicht aus derselben
  Vorlage. Die Kürzung selbst ist unproblematisch (die Disziplin bleibt
  auffindbar), aber die Tatsachenbehauptung „steht bereits in docs/reviews/README.md
  und .harness/skills/reviewer.md" ist für das Dateiname-Muster nicht zutreffend.
- `verifizierbar`: ja — `grep -n "YYYY-MM-DD.*slice-oder-diff-ref" docs/reviews/README.md .harness/skills/reviewer.md` liefert keinen Treffer, wohl aber in `.harness/baseline/v6.0.0/templates/docs/reviews/review-report.template.md`.
- `klasse`: Duplizierungs-Beleg nennt falsche Fundstelle statt der tatsächlichen

### F-2 — `Ersetzt-Baseline-Regel` von MR-018 dupliziert MR-016 ohne Prüfung der näherliegenden Alternative

- `kategorie`: LOW
- `quelle`: `harness/conventions/MR-018-review-pflicht-v610-wortlaut.md` (Kopf); `harness/conventions/MR-016-validator-unbesetzt.md` (Kopf)
- `pfad`: `harness/conventions/MR-018-review-pflicht-v610-wortlaut.md:6`
- `befund`: MR-018 zeigt exakt denselben Anker wie MR-016
  (`modul-08-agentenrollen.md#die-neun-übergaben-und-ihre-artefakte-modul-8`),
  obwohl der übernommene Text wörtlich aus `AGENTS.template.md` stammt, dessen
  Ziel-Form und Inhalts-Regeln `modul-09-implementierung.md` §Ziel-Form: AGENTS.md
  bzw. §AGENTS.md-Regeln trägt. Modul 8 ist inhaltlich vertretbar (der Absatz
  handelt von Rollen-Trennung/Kontext-Trennung), aber die Alternative wird im
  Eintrag nicht diskutiert — bei einer dritten Wiederholung desselben Anker-Musters
  wäre das ein Kandidat für das Beobachtungs-Register.
- `verifizierbar`: nein — Urteilsfrage, kein Gate-Lauf entscheidet den treffenderen Anker.
- `klasse`: MR-Anker kopiert statt gegen Alternativ-Anker geprüft

## Negativbefunde

- geprüft, ohne Befund: Wortlaut-Vergleich `AGENTS.md` §6 (neuer Absatz) gegen
  `git diff v6.0.0 v6.1.0 -- lab/templates/AGENTS.template.md` — Basistext ist
  wortgleich (Interpunktion/Zeilenumbrüche identisch bis auf Markdown-Linkifizierung
  zweier bereits im Kurs-Text vorhandener Backtick-Pfade, `docs/reviews/README.md`-Link
  ausgenommen, der im Kurs-Text nicht vorkommt); genau zwei zusätzliche Sätze
  (`fork`-Ausschluss, `BEO-HARNESS`-Zitat-Anker) wie behauptet, keine weiteren
  stillen Abweichungen.
- geprüft, ohne Befund: Hard Rule §3.7 am neuen Absatz — die entfernte
  Chronik-Passage ("Diese Zeile ist die Korrektur eines echten, gemessenen
  Ausfalls … weil genau das Fehlen dieser Zeile … der Auslöser war") ist restlos
  weg; was bleibt, ist Indikativ über den geltenden Zustand plus ein einziger
  auflösbarer Rang-Zeiger (`Ausfallbeleg: …`) — die Klasse „Rang-Zeiger" aus §3.7,
  keine Begründungs-Chronik mehr.
- geprüft, ohne Befund: HIGH-Verifikationspflicht — bleibt wörtlich unverändert
  in `.harness/skills/reviewer.md` Zeile 59–61 stehen (nicht in `AGENTS.md`
  dupliziert, korrekt gemäß Behauptung).
- geprüft, ohne Befund: `MR-018`-Formvergleich gegen `MR-016`/`MR-017` — identische
  Feldmenge und -reihenfolge (Status, Datum, Geltungsbereich, Ersetzt-Baseline-Regel,
  Adaption, Begründung, Auflösungs-Trigger, Ausgelöst durch Baseline-Stand); kein
  `Löst auf`-Feld, korrekt, da MR-018 keinen bestehenden Adaptions-Eintrag ablöst;
  Tabellenzeile in `harness/conventions.md` trägt beide Pflicht-Anker
  (`<a id="mr-018">` und Slug-Verlinkung) wie die Nachbarzeilen.
- geprüft, ohne Befund: Relative Markdown-Links in
  `docs/plan/planning/done/slice-162-review-pflicht-absatz-v610-wortlaut.md`
  — `make doc-check` (0 Befunde, 408 Dateien) deckt sie ab, keine toten Anker.
- geprüft, ohne Befund: DoD-/Risiko-Form von slice-162 gegen `.d-check.yml`-Modul
  `structure` — `make doc-structure` und `make verify` beide grün (je 0 Befunde,
  `verify-risiko-ausgaenge` erfasst slice-162 explizit als „abschlussbereit in
  in-progress/" mit beiden Risiken auf geschlossenem Ausgang „gestrichen mit
  Begründung"); die beiden `- [ ]`-Punkte (Beobachtungs-Register, Risiko-Ausgänge)
  sind zulässig offen, da der Slice noch nicht nach `done/` gewandert ist.
- geprüft, ohne Befund: `make gates` — Exit 0, alle Teil-Gates grün, inkl.
  `doc-reviews`, `doc-planning`, `doc-workflows`, `doc-targets`
  (je 408 Dateien, 0 Befunde), `suppression-check`, `guard-selftest`,
  `ci-range-selftest`, `record-gates`.
- geprüft, ohne Befund: Provenienz-Zitat „taggleich mit slice-159 (2026-09-05)" —
  `git log -1 --format=%ad --date=short 7de8569` bestätigt `2026-09-05`; und
  `v6.1.0`-Tag-Metadaten des Kurs-Repos bestätigen „Welle 118" und Datum
  `2026-09-05 18:32:46 +0200`, wie in MR-018 behauptet.
- geprüft, ohne Befund: Commit-Scope — Präfix `build(harness)`, nicht
  `(planning)`, daher greift die Scope-Restriktion aus §5 nicht; `make
  commit-scope-check` bestätigt 0 `(planning)`-Commits im geprüften Range.
- nicht geprüft (Werkzeug-Grenze, nicht Gegenstand des Befunds): `make
  trace-check` scheiterte in dieser Sitzung mit „Range-Basis-Vorfahren nicht
  lesbar" (Docker-Mount-/Range-Eigenheit der Prüfumgebung, kein Repo-Defekt —
  `git cat-file -t` bestätigt beide Commits lokal vorhanden); die
  ID-Traceability-Pflicht (§5) ist am Commit-Message-Text ohnehin visuell
  erfüllt (`slice-162, slice-161`).

## Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 0 |
| MEDIUM | 1 |
| LOW | 1 |
| INFO | 0 |

**Finding-Klassen dieses Laufs:** Duplizierungs-Beleg nennt falsche
Fundstelle statt der tatsächlichen · MR-Anker kopiert statt gegen
Alternativ-Anker geprüft

## Verdikt

**Merge-blockierend:** nein — kein HIGH-Finding, die Wortlaut-Kürzung selbst
ist gegen `v6.1.0` unabhängig nachgeprüft deckungsgleich, §3.7 ist am neuen
Text sauber eingehalten, `make gates`/`make doc-structure`/`make verify` sind
grün. Das MEDIUM-Finding (F-1) betrifft nur die *Begründung* der Kürzung
(welche Datei welches Detail trägt), nicht die Kürzung selbst — der
inhaltliche Verlust, den es beschreibt, ist keiner (die Substanz steht
weiterhin auffindbar in `.harness/skills/reviewer.md`, nur nicht wortgleich
mit der behaupteten Fundstelle). Empfehlung: MR-018s Begründungssatz vor der
nächsten Berührung dieser Datei auf die tatsächliche Fundstelle
(`review-report.template.md`) korrigieren — MR-Einträge sind nach `Accepted`
inhaltlich unveränderlich (§Adaptions-Block-Disziplin), die Korrektur würde
also als eigener Nachtrag-Eintrag laufen, nicht als Edit dieser Datei.

**Übergabe:** Findings gehen an den Implementer. Die beiden Finding-Klassen
gehen in die Slice-Closure §7 von slice-162 (sobald es nach `done/`
übergeht) und von dort in den Zähler des Beobachtungs-Registers. Dieser
Report ist Lauf-Beleg, keine Verifikation — DoD-/Spec-Konformität prüft der
Verifier separat (Modul 11).
