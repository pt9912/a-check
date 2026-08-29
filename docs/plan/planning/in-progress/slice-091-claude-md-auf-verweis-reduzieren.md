# slice-091 — CLAUDE.md auf den reinen Verweis reduzieren

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Deckt:** keine `AC-*`/`ADR-*` — Harness-Änderung ohne Vertragsberührung, wie
[slice-047](../done/slice-047-baseline-vendoring.md).
**Bezug:** Maintainer-Auftrag 2026-08-29 („CLAUDE.md ausdünnen, so dass nur noch `@AGENTS.md`
drin steht"), aus der Prüfung „steht das nicht schon in AGENTS.md?". [Roadmap](../in-progress/roadmap.md).

---

## 1. Auslöser

`CLAUDE.md` trägt 23 Zeilen mit **12 Aussagen**. Gemessen gegen die Originale:

| Herkunft | Anzahl |
|---|---|
| Dopplung aus `AGENTS.md` (§1, §2, §3.1, §6.1, §6.3, §6.6, §6.8) | 10 |
| wörtliche Dopplung aus [`.claude/commands/slice.md`](../../../../.claude/commands/slice.md) Schritt 3 | 1 |
| eigene Substanz („Konflikt melden") | 1 |

**Zwei Dopplungen sind bereits abgedriftet — die Kopie ist schlechter als das Original:**

- `CLAUDE.md:17-18` listet `pip`, `npm`, `cargo`, `apt`, `brew` und lässt **`go`/`golangci-lint`
  weg** — in einem Go-Repo genau die beiden, die der PreToolUse-Guard zusätzlich blockt
  (`AGENTS.md` §3.1).
- `CLAUDE.md:21` verlangt vor dem Abschluss nur `make gates`. `AGENTS.md` §6 verlangt dort
  **zusätzlich `make verify`**. Wer nur die Kopie liest, überspringt die Verifikations-Schicht.

Die Datei wirkt dabei als Regelquelle: der Review-Report zu slice-042 führt sie unter `quelle`
neben `AGENTS.md` §4
([F-4](../../../reviews/2026-07-25-slice-042-constructs-monopol.md)). Sie ist damit ein
**vierter Bindepunkt** neben den dreien der Durchsetzungsschicht — undeklariert, von keinem Gate
erfasst, seit ihrem einzigen Commit (`70f1fd3`) unverändert.

**Warum jetzt und nicht mit der Baseline-Migration:** die vendored Baseline `v3.5.2` sagt zu
`CLAUDE.md` **nichts** — es gibt heute keine Sollform, gegen die die Datei verstößt. Dieser Slice
steht deshalb ausschließlich auf dem gemessenen Duplikations- und Drift-Befund oben, nicht auf
einer Regel aus einem Stand, den das Repo nicht adoptiert hat.

## 2. Betroffene Module

- `CLAUDE.md` — Inhalt wird durch `@AGENTS.md` ersetzt.
- `AGENTS.md` §6 Schritt 3 — nimmt die eine Aussage auf, die sonst ersatzlos verschwände.

Eine Schicht (Harness-Einstiegs-Doku), zwei Dateien.

## 3. Auszuführende Gates

`make gates` — tragend ist `doc-check`: `.d-check.yml` scannt `roots: ["."]` und nimmt nur
`.harness/baseline/**` aus, `CLAUDE.md` läuft also mit. Zum Abschluss `make verify`.

**Kein neuer Sensor**, also keine Negativ-Probe. Was der Slice herstellt, ist ein Datei-Zustand,
den ein Diff belegt — kein Zustand, den ein Gate künftig hält (siehe §4).

## 4. Was bewusst nicht getan wird

- **Kein Sensor, der `CLAUDE.md` klein hält.** Die Datei ist einmal in 91 Slices gewachsen; ein
  Gate gegen ein einmaliges Ereignis ist der Steering-Loop-Schwelle nach verfrüht (Eintrag ab dem
  zweiten Vorfall, Sensor ab dem dritten — `AGENTS.md` §5).
- **`harness/README.md` §Minimal agent workflow bleibt unangetastet**, obwohl Schritt 3 dort
  dieselbe Verkürzung trägt („Betroffene IDs identifizieren"). Diese Datei ist eine bewusst
  kondensierte Zweitfassung des gesamten §6 — sie anzugleichen hieße, alle acht Schritte
  anzugleichen. Eigener Befund, gehört in die Baseline-Migration.
- **Keine Sub-Area-Zeile** für die Harness-Einstiegs-Schicht in
  `harness/conventions.md` §Modus-Deklaration pro Sub-Area (§7 nennt die Lücke).
- **Keine Vorwegnahme der Migration** auf einen neueren Kurs-Stand.

## 5. DoD

- [ ] `CLAUDE.md` enthält genau `@AGENTS.md`; die beiden gemessenen Drifts aus §1 sind damit
      ersatzlos entfallen — Beleg: Datei-Inhalt und `git diff`.
- [ ] `AGENTS.md` §6 Schritt 3 nennt Slice-ID, `AC-*`, `ADR-*`, betroffene Module und
      auszuführende Gates — Beleg: Diff, deckungsgleich mit
      [`.claude/commands/slice.md`](../../../../.claude/commands/slice.md) Schritt 3.
- [ ] `make gates` (und bei Abschluss `make verify`) grün — **Ausgabe in eine Datei**, Exit-Code
      getrennt geprüft, nie in eine Pipe.

## 6. Closure-Notiz

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice,
[`AGENTS.md`](../../../../AGENTS.md) §5; `make verify` prüft das.)_

## 7. Sub-Area-Modus

Berührt werden **Planungs-Harness** (`docs/plan/planning/`, Greenfield) und die
Harness-Einstiegs-Dateien `AGENTS.md`/`CLAUDE.md`. Alle berührten Sub-Areas mit Modus sind GF.

**Benannte Lücke:** für `AGENTS.md`/`CLAUDE.md` gibt es **keine** Zeile in
`harness/conventions.md` §Modus-Deklaration pro Sub-Area — die Tabelle führt `.claude/` unter
„Gate-/Werkzeug-Schicht", aber die Einstiegs- und Briefing-Dateien in der Repo-Wurzel unter
keiner Sub-Area. Nach der Qualifikations-Regel (mindestens zwei der drei Inklusions-Achsen) wäre
sie eine: eigene Datei-Familie (Achse 3) und eine eigene
[MR-Adaption](../../../../harness/conventions.md#mr-006--baseline-committet-vendored-statt-per-url-referenziert)-Linie
ist plausibel (Achse 1). Schließen dieser Lücke ist **nicht** Teil dieses Slice (§4).
