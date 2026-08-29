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

**Der Fall des alten Baums ist gemessen — und die erste Messung war falsch.** Eine Suche nach dem
Textmuster `](…​.harness/baseline/v3.5.2` ergab „kein einziger Link". Das stimmte nicht: eine
Auflösung der Link-**Ziele** statt ihrer Schreibweise findet **genau einen** echten Verweis, in
[`.harness/skills/reviewer.md`](../../../../.harness/skills/reviewer.md) — dort steht er relativ
als `../baseline/v3.5.2/…` und enthält die Zeichenkette gar nicht. Er wird mitgezogen. Die
übrigen 16 Dateien nennen den Pfad nur in Backticks: historische Aussagen in Closure-Notizen und
Review-Reports, die korrekt bleiben und nicht umgeschrieben werden.

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

- [x] `harness/README.md` trägt `## Leseordnung` mit drei bis fünf **geordneten** Zeigern —
      Beleg: Diff; die Pflichtgliederung aus `grundlagen-harness-dateien.md` ist damit vollständig.
- [x] Die fünf Provenienz-Zeiger stehen auf `v5.12.0`, und der Form-Review ist je Klasse belegt
      (§3) statt behauptet — Beleg: Diff und `diff -rq`-Messung.
- [x] `.harness/baseline/v3.5.2/` ist entfernt, `harness/conventions.md` §Baseline sagt es, und
      `make regelwerk-check` meldet **einen** Stand ohne Hinweis-Zeile — Beleg: Target-Ausgabe.

Pflicht, aber **kein** Liefer-Punkt: `make gates` und zum Abschluss `make verify` grün — Ausgabe
in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.

## 7. Closure-Notiz

**Geliefert:** `harness/README.md` erfüllt seine Pflichtgliederung, die Provenienz-Zeiger stehen
auf dem adoptierten Stand, und der alte vendored Baum ist gefallen — 42 Dateien weg, `doc-check`
danach bei 213 Dateien und 0 Befunden. Damit ist **Etappe C abgeschlossen** und von der Migration
nur noch D offen.

**Lerneintrag — Form: geschärfte Regel.** *Ein Verweis wird an seinem **aufgelösten Ziel**
gemessen, nicht an seiner Schreibweise.* Vor dem Löschen habe ich gefragt, ob ein Link auf den
alten Baum zeigt, und per Textmuster `](…​.harness/baseline/v3.5.2` gesucht. Ergebnis: „kein
einziger". **Das war falsch.** In [`.harness/skills/reviewer.md`](../../../../.harness/skills/reviewer.md)
steht der Verweis relativ als `../baseline/v3.5.2/…` und enthält die gesuchte Zeichenkette
überhaupt nicht. Erst eine Auflösung der Ziele — Link relativ zum Ort der Datei normalisieren und
gegen das Verzeichnis prüfen — fand ihn. *Weil* eine Textsuche die Schreibweise trifft und nicht
die Bedeutung, und relative Pfade dieselbe Datei auf beliebig viele Arten schreiben. Das Gate
hätte den Bruch danach gefangen; die falsche Behauptung wäre trotzdem in `done/` gewandert.

**Zwei beobachtbare Closure-Kriterien:**

1. `make regelwerk-check` läuft mit Exit 0 und **ohne** die Hinweis-Zeile über mehrere Stände —
   es gibt wieder genau einen (`v5.12.0`, 51 Dateien gegen `SHA256SUMS`, beide Richtungen).
2. `doc-check` prüft **213** Dateien mit 0 Befunden, **nachdem** 42 gelöscht wurden. Dass die
   Zahl steigt statt zu fallen, ist der Beleg für den `scan.ignore`: der Baum war nie Teil der
   geprüften Menge, und niemand zeigte auf ihn außer der einen mitgezogenen Stelle.

**Offene Risiken und ihr Ausgang:**

- *Die drei neuen Vorlagen (`observations`, `reconciliation`, `welle-results`) sind nicht
  adoptiert* — Ausgang: **Folge-Slice**, Etappe D. `observations` **ist** das
  Beobachtungs-Register.
- *Drei ausgewiesene Auslegungen aus C2 und C3 warten weiter auf einen Ort* („höchstens zwei
  Schichten" ungeprüft · die `done/`-Platzierung des Rückbau-Eintrags · der Rückbau-Kandidat unter
  den Adaptionen) — Ausgang: **Folge-Slice**, Etappe D, mit demselben Register.

- *Die Placeholder-Heuristik von `verify-closure-notes` kollidiert mit dem neuen Risiko-Block* —
  sie kennt `noch offen` als Platzhalter-Wendung, und die neue Baseline verlangt in **jeder**
  Closure-Notiz einen Abschnitt über *offene* Risiken. Der Zusammenstoß ist systematisch, nicht
  zufällig. Ausgang: **weiter offen**, fürs Beobachtungs-Register (D). Ein Sensor-Eingriff wäre
  hier verfrüht — die Steering-Loop-Regel verlangt den zweiten Vorfall für einen Eintrag und den
  dritten für einen Sensor; dies ist der erste, und die Umformulierung war billiger als die Regel.
- *Die vorgeschriebene Reihenfolge kann die Prüfung nicht abnehmen, die sie erfüllen soll* —
  `verify-closure-notes` greift **nur in `done/`**, der Workflow führt `make verify` aber in
  Schritt 8 aus und den `git mv` erst in Schritt 9. Die Closure-Notiz eines Slice wird damit
  frühestens beim **nächsten** Slice geprüft; genau so ist der Befund an slice-098 erst hier
  aufgetaucht. Ausgang: **weiter offen**, fürs Register (D) — die Lücke ist im Ablauf, nicht in
  diesem Slice.

**Folge-Slices:** Etappe D — der letzte Schritt der Migration.

## 8. Sub-Area-Modus

Berührt werden die Harness-Doku und die **Vendored Baseline** (`.harness/baseline/`, laut
Modus-Tabelle ausdrücklich **kein Modus** — externer Fremdtext). Alle berührten Sub-Areas mit
Modus sind GF.
