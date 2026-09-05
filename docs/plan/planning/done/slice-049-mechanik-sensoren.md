# slice-049 — Etappe E (1/3): zwei fehlende Sensoren nachrüsten

**Status:** *(der Zustand ist das Verzeichnis dieser Datei, nicht dieses Feld — korrigiert in slice-063)* — erster Schnitt der **Etappe E (Mechanik)** aus
[slice-048 §5](../done/slice-048-modul-delta-lesen.md); Reihenfolge **E vor D** am 2026-07-25 per
Maintainer-Wort abgenommen.
**Welle:** welle-12-regelwerk-migration.
**Deckt:** Fund **B-11** (Suppression-Verbot halb durchgesetzt) und **B-8** (kein
Maintenance-Target für die vendored Baseline).
**Nicht in diesem Slice:** B-3 (`verify` + `check-references`), B-4 (`closure-note-reviewer`),
B-16 (Workflow-Skelett), B-20 (Freigabe-Checkliste) — sie folgen als 2/3 und 3/3, damit jeder
Schnitt die Größen-Regel aus **B-1** einhält (≤ 3 DoD-Punkte, höchstens zwei Schichten).
[Roadmap](../in-progress/roadmap.md).

---

## 1. Warum diese zwei zuerst

Beide sind **Sensor-Funde**: sie fügen Beobachtung hinzu, statt Form umzuschreiben. Genau das ist
die Begründung für *E vor D* — die Form-Funde aus Etappe D hängen an Sensoren, und eine Form ohne
Sensor ist laut [slice-048 §2](../done/slice-048-modul-delta-lesen.md) die Praxis, die in diesem
Repo schon zweimal eingeschlafen ist.

**B-11**: die Hard Rule [`AGENTS.md`](../../../../AGENTS.md) §3.2 existierte **nur** als Prosa und
als Kommentar in [`.golangci.yml`](../../../../.golangci.yml) Zeile 3. Nach Modul 09 ist eine Hard
Rule in nur einem Quadranten *halb durchgesetzt*. Der Wert liegt nicht in einem heutigen Verstoß —
es gibt **null** `//nolint` im Code —, sondern darin, dass ein künftiger unbeanstandet durchliefe.

### 1.1 Der Plan war falsch: `nolintlint` kann die Regel nicht ausdrücken

slice-048 hatte B-11 als „billigsten Fund" geführt: *einen Linter aktivieren*. Die vom DoD
verlangte Negativ-Probe hat das widerlegt — und zwar erst im dritten Anlauf, weil die ersten beiden
kontaminiert waren:

| Probe | Eingriff | `make lint` | Aussage |
|---|---|---|---|
| A | nackter `//nolint` | **rot** | scheinbarer Beleg — aber nur, weil die Direktive *unused* war |
| B | wohlgeformter `//nolint` auf ungenutzter Variable | rot | **unbrauchbar**: rot durch `unused`, der Verstoß selbst war unterdrückt |
| **C** | wohlgeformter `//nolint` auf **genutzter** Variable | **grün (Exit 0)** | `nolintlint` schweigt; die Unterdrückung wirkt |
| **D** | dieselbe Stelle **ohne** `//nolint` | rot (`gochecknoglobals`) | beweist, dass C einen **echten** Verstoß verdeckte |

C und D zusammen sind der Beweis: `nolintlint` prüft die **Wohlgeformtheit** von Direktiven, nicht
ihre **Existenz**. Die Regel „Inline-Suppressions sind verboten" ist damit für diesen Linter
gar nicht formulierbar. Er wurde folglich **nicht** aktiviert — neben einem echten Verbot könnte er
nie feuern und wäre eine tote Regel, also genau der Verfall, den *Entropy Management* beschreibt.

Durchgesetzt wird §3.2 stattdessen von `tools/suppression-check.sh`, das jede `//nolint`- und
`//lint:ignore`-Direktive ablehnt und einen eigenen Selbsttest gegen ein totes Muster mitführt.

## 2. Keine Spec-, keine ADR-Änderung — und warum das kein Schlupfloch ist

Beide Änderungen berühren **keinen** Vertrag der Produkt-Achse: kein `AC-*`, kein `SPEC-*`, kein
`ARC-*`. Nach der Source Precedence ist das reine Harness-Ebene.

- **`suppression-check`** *implementiert* [ADR-0005](../../adr/0005-lint-profil.md), es entscheidet
  nichts Neues. Die ADR trägt die Entscheidung bereits im Titel („SOLID-nahe Linter **ohne**
  `//nolint`") und im Entscheidungs-Satz („Inline-Suppression ist verboten; Ausnahmen leben
  zentral unter `exclusions` mit `Why:`-Kommentar"). Was fehlte, war die Durchsetzung, nicht die
  Entscheidung. Eine neue ADR wäre hier Inflation — und die bestehende ist nach
  [`AGENTS.md`](../../../../AGENTS.md) §3.5 ohnehin immutabel. Dass die Durchsetzung als Skript
  statt als Linter-Zeile entsteht, ändert die Entscheidung nicht, sondern folgt aus §1.1: das von
  der ADR gewählte Werkzeug kann diese eine Regel nicht ausdrücken.
- **`regelwerk-check`** ist **kein Gate**, sondern ein Maintenance-Target. Seine Bindung ist
  [MR-006](../../../../harness/conventions.md#mr-006--baseline-committet-vendored-statt-per-url-referenziert)
  (vendored Baseline) plus der Etappe-A-Fund F-6 aus
  [slice-047](../done/slice-047-baseline-vendoring.md). [`AGENTS.md`](../../../../AGENTS.md) §3.6
  verlangt eine ADR für jede **Lockerung** — hier wird nichts gelockert, sondern Beobachtung
  hinzugefügt.

## 3. `regelwerk-check`: was es prüft — und was ausdrücklich nicht

Modul 13 nennt `regelwerk-check` als das Musterbeispiel für „vorhanden, aber **nicht** als Gate
behauptet". Der Fund F-6 aus Etappe A trennt zwei Hälften, und diese Trennung wird hier
**wörtlich** eingehalten:

| Hälfte | Charakter | in diesem Slice |
|---|---|---|
| **Integrität** — stimmen die in `SHA256SUMS` gelisteten **42** Dateien noch? | hermetisch, deterministisch | **geprüft**, fail-closed |
| **Freshness** — gibt es stromaufwärts ein neueres Release als `v3.5.2`? | Netz-Operation gegen die Release-Liste | **nicht geprüft**, sondern als offene Handlung ausgegeben |

Die zweite Zeile ist der Punkt, an dem Modul 02 der ersten Fassung von slice-047 widersprochen
hatte: Freshness ist Wartung, kein Gate. Das Target **behauptet sie darum nicht** — es druckt den
adoptierten Stand, die Release-URL und den Satz, dass die Freshness in diesem Lauf *nicht* geprüft
wurde. Ein Target, das mehr abzudecken vorgibt, als es tut, wäre nach `grundlagen-durchsetzungsschicht`
selbst eine Harness-Lüge.

**Bewusst offen gelassen:** die Integritäts-Hälfte ist hermetisch und wäre damit *gate-fähig* —
sie wandert hier trotzdem **nicht** in `make gates`. Grund: das wäre eine eigene Entscheidung
(jede beabsichtigte Baseline-Aktualisierung müsste `SHA256SUMS` im selben Commit mitführen), und
sie gehört nicht in einen Slice, der laut B-1 klein bleiben soll. Vermerkt als Kandidat für den
Etappe-E-Rest, nicht vergessen.

## 4. Betroffene Module

- `tools/suppression-check.sh` + [`.golangci.yml`](../../../../.golangci.yml) (erklärender
  Kommentar, warum `nolintlint` die Regel nicht trägt) — B-11.
- `tools/regelwerk-check.sh` — Maintenance-Target (B-8).
- [`Makefile`](../../../../Makefile) — beide Targets, `suppression-check` zusätzlich im
  `gates`-Aggregat.
- [`AGENTS.md`](../../../../AGENTS.md) §4 — Target-Tabelle; `gate-consistency` erzwingt, dass
  jedes reale Nicht-Utility-Target dort steht.

Zwei Schichten (Lint-Konfiguration, Build-/Harness-Targets) — innerhalb der B-1-Grenze.

## 5. Auszuführende Gates

`make gates` (enthält `lint`, `gate-consistency` und neu `suppression-check`), zusätzlich je eine
**Negativ-Probe** pro Sensor: ohne Beweis, dass der Sensor bei echtem Verstoß rot wird, ist er ein
toter Sensor — dieselbe Beweisführung wie bei der Schattenwurf-Diskriminierungsprobe in
[slice-044](../done/slice-044-ziel-glob-schattenwurf.md). §1.1 zeigt, warum das hier keine
Formalie war: die Probe hat den geplanten Ansatz verworfen.

## 6. DoD

- [x] Hard Rule §3.2 maschinell durchgesetzt: `make suppression-check` lehnt jede
      `//nolint`-/`//lint:ignore`-Direktive ab, trägt einen Selbsttest gegen ein totes Muster und
      ist im `gates`-Aggregat. Per Negativ-Probe belegt, dass genau die Direktive, die `make lint`
      grün ließ (§1.1 Probe C), hier **rot** wird und nach Rücknahme wieder grün (B-11).
- [x] `make regelwerk-check` existiert als Maintenance-Target außerhalb des `gates`-Aggregats:
      prüft die Integrität der vendored Baseline gegen `SHA256SUMS` fail-closed und weist die
      Freshness-Hälfte ausdrücklich als **nicht geprüft** aus; per Negativ-Probe belegt, dass eine
      verfälschte Datei erkannt wird (B-8).
- [x] `make gates` grün, inklusive `gate-consistency` mit dem neuen Target in
      [`AGENTS.md`](../../../../AGENTS.md) §4.

## 7. Closure-Notiz

**Geliefert:** `make suppression-check` als Fitness Function zu [`AGENTS.md`](../../../../AGENTS.md)
§3.2 (im `gates`-Aggregat) und `make regelwerk-check` als Wartungs-Target ohne Gate-Anspruch. Beide
mit Selbsttest und Negativ-Probe belegt.

**Lerneintrag — Form: geschärfte Regel.**
> **Der Name eines Werkzeugs ist kein Beleg für seine Semantik.** `nolintlint` klingt nach
> „setzt das `//nolint`-Verbot durch" und prüft in Wahrheit nur die *Wohlgeformtheit* von
> Direktiven. Wäre der Fund ohne Probe umgesetzt worden, hätte das Repo ein Gate **behauptet**,
> das die Hard Rule nicht trägt — eine Harness-Lüge, entstanden aus gutem Willen. Prüfregel für
> jeden künftigen Sensor: *erst die Negativ-Probe, dann die Erfolgsmeldung* — und die Probe muss
> den Fall treffen, der wirklich durchrutschen könnte, nicht den bequemsten. Hier brauchte es
> drei Anläufe, weil die ersten beiden an `unused` scheiterten statt am Prüfgegenstand.

**Zwei beobachtbare Closure-Kriterien:**

1. `make gates` grün mit `suppression-check` im Aggregat (Exit 0) — belegt.
2. Beide Sensoren werden bei echtem Verstoß rot und nach Rücknahme wieder grün — je eine
   Negativ-Probe dokumentiert (§1.1 und DoD).

**Folge-Slice:** [slice-050](slice-050-verify-schicht.md) (Etappe E, 2/3). Offen aus diesem
Slice: ob die hermetische Integritäts-Hälfte von `regelwerk-check` in `make gates` gehört (§3,
bewusst vertagt).
