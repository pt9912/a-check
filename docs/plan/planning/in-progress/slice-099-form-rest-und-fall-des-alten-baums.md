# slice-099 — Etappe C4: Rest der Form und Fall des alten Baums

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.
**Autor:** Claude (Opus 5), im Auftrag des Maintainers.
**Berührte Spec-Stellen:** — *(keine; Harness-Form ohne Vertragsberührung)*
**Deckt:** keine `AC-*`/`ADR-*`.
**Bezug:** Etappe **C** aus [slice-092 §6](../done/slice-092-regelwerk-v5120-delta-analyse.md),
vierte und letzte Hälfte. Vorgänger
[slice-098](../done/slice-098-slice-form-liefer-punkte.md).
[Roadmap](../in-progress/roadmap.md).

---

## 1. Auslöser

Diese Menge war nach C1 aus meiner eigenen Etappen-Aufzählung gefallen und ist in
[slice-098 §5](../done/slice-098-slice-form-liefer-punkte.md) ausdrücklich wieder aufgenommen
worden. Sie besteht aus drei Dingen:

1. **`## Leseordnung`** ist in `v5.12.0` neuer **Pflicht-Abschnitt** von `harness/README.md`.
   Für Singletons gilt laut `modul-02`: neue *optionale* Felder verlangen keine Nacharbeit, neue
   **Pflicht**-Felder und umbenannte Sektionen schon — *„sonst behauptet das Repo eine
   Baseline-Konformität, die seine Artefakte nicht tragen."*
2. **Fünf Provenienz-Zeiger** sagen weiterhin „übersetzt aus `.harness/baseline/v3.5.2/…`".
3. **Der alte vendored Baum** darf erst fallen, *„wenn der Review durch ist"* — und dieser Slice
   ist der Review.

## 2. Betroffene Module

- `harness/README.md` — neuer Abschnitt, plus der Provenienz-Zeiger im Sensors-Kommentar.
- `docs/plan/planning/slice.template.md`, `docs/plan/carveouts/carveout.template.md`,
  `docs/reviews/README.md`, `.harness/skills/reviewer.md` — Provenienz-Zeiger.
- `.harness/baseline/v3.5.2/` — entfällt; `harness/conventions.md` §Baseline sagt es.

Zwei Schichten: Harness-Doku und vendored Baseline.

## 3. Der Form-Review, gemessen

`diff -rq` über beide Template-Stände: **20 Vorlagen geändert, 4 neu**. Der Review ist damit
aber nicht 24 Positionen groß — `modul-02` begrenzt ihn:

| Klasse | Regel | Für a-check |
|---|---|---|
| **Singletons mit Pflichtgliederung** | neue Pflicht-Sektion ⇒ Nacharbeit | `harness/conventions.md` erledigt (C1/C2); `harness/README.md` fehlt **`## Leseordnung`** |
| Singletons ohne Pflichtgliederung | Referenz-Form entscheidet | `AGENTS.md` und Lastenheft — `modul-02` nennt sie ausdrücklich als ohne |
| **Wiederkehrende Vorlagen** | Append-only: neue Instanzen folgen der neuen Form, bestehende nicht rückwirkend | `slice.template.md` erledigt (C3); `carveout.template.md` **strukturell unverändert** — Feld- und Abschnitts-Vergleich zeigt kein Delta |
| Neue Vorlagen | — | `observations`/`reconciliation`/`welle-results` gehören zu Etappe **D** |

**Der Fall des alten Baums ist gefahrlos, und das ist gemessen statt gehofft:** kein einziger
Markdown-**Link** zeigt auf `.harness/baseline/v3.5.2/`. Die 16 Dateien, die den Pfad nennen,
tun das in Backticks — historische Aussagen in Closure-Notizen und Review-Reports, die korrekt
bleiben und nicht umgeschrieben werden.

## 4. Auszuführende Gates

`make regelwerk-check` (muss danach **einen** Stand melden statt zwei), `make gates` mit
`doc-check` als tragendem Teil, zum Abschluss `make verify`.

**Kein neuer Sensor**, also keine Negativ-Probe. Die Probe ist der Bestand: 212 geprüfte Dateien
mit 0 Befunden **nach** dem Löschen von 42 Dateien ist die Aussage, dass niemand auf sie zeigte.

## 5. Was bewusst nicht getan wird

- **Keine historische Aussage wird umgeschrieben.** Closure-Notizen und Review-Reports nennen
  `v3.5.2` als das, was damals galt. Das bleibt.
- **Die drei neuen Vorlagen werden nicht adoptiert.** `observations` ist das Beobachtungs-Register
  — das ist Etappe D, und dort landen auch die drei bisher ausgewiesenen Auslegungen.
- **`AGENTS.md` bekommt keine Pflichtgliederung angelegt.** `modul-02` sagt ausdrücklich, dass es
  für `AGENTS.md` und das Lastenheft keine gibt; eine zu erfinden wäre die stille Setzung, gegen
  die dieses Repo sonst vorgeht.

## 6. DoD

- [ ] `harness/README.md` trägt `## Leseordnung` mit drei bis fünf **geordneten** Zeigern —
      Beleg: Diff; die Pflichtgliederung aus `grundlagen-harness-dateien.md` ist damit vollständig.
- [ ] Die fünf Provenienz-Zeiger stehen auf `v5.12.0`, und der Form-Review ist je Klasse belegt
      (§3) statt behauptet — Beleg: Diff und `diff -rq`-Messung.
- [ ] `.harness/baseline/v3.5.2/` ist entfernt, `harness/conventions.md` §Baseline sagt es, und
      `make regelwerk-check` meldet **einen** Stand ohne Hinweis-Zeile — Beleg: Target-Ausgabe.

Pflicht, aber **kein** Liefer-Punkt: `make gates` und zum Abschluss `make verify` grün — Ausgabe
in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.

## 7. Closure-Notiz

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice,
[`AGENTS.md`](../../../../AGENTS.md) §5; `make verify` prüft das.)_

## 8. Sub-Area-Modus

Berührt werden die Harness-Doku und die **Vendored Baseline** (`.harness/baseline/`, laut
Modus-Tabelle ausdrücklich **kein Modus** — externer Fremdtext). Alle berührten Sub-Areas mit
Modus sind GF.
