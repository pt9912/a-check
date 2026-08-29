# slice-103 — `harness/conventions.md` trägt den Ist-Zustand, keine Chronik

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.
**Autor:** Claude (Opus 5), im Auftrag des Maintainers.
**Berührte Spec-Stellen:** — *(keine; Harness-Konventionen ohne Vertragsberührung)*
**Deckt:** keine `AC-*`/`ADR-*`.
**Bezug:** Maintainer-Vorgabe 2026-08-29: *„Keine Chronik oder Forensik. Damit wird nur der
Kontext des Code-Agenten zugemüllt."* [Roadmap](../in-progress/roadmap.md).

---

## 1. Auslöser

`harness/conventions.md` liest **jeder** Agentenlauf — das ist der Grund, aus dem der
Adaptions-Speicher in slice-096 die Verzeichnis-Form bekam. Dieselbe Rechnung gilt für den Rest
der Datei, und dort steht Chronik.

**Drei Stellen, gemessen gegen die Pflichtgliederung** in `grundlagen-harness-dateien.md`
§Konventionsspeicher und die Vorlage:

| Stelle | Sollform | Ist |
|---|---|---|
| §Baseline | *„welche Konvention adoptiert, mit Stand/Version"* | zusätzlich die Etappen-Chronik zweier Migrationen |
| §Adoptierte Konventions-Quellen | **drei** Bullets (Extern · Vendored · In-Repo), *„KEINE Wiederholung des Inhalts — nur Verweise"* | **fünf** |
| §Adaptions-Block, Präambel | Index-Regeln | zusätzlich ein Block „Zur Nummern-Frage (geprüft in slice-055, ohne Befund)" |

**Der Beweis, dass Chronik nicht bloß Ballast ist, steht in der Datei selbst.** §Baseline behauptet
heute *„Offen ist allein **D**"* — geschrieben in slice-099, überholt seit slice-102. Eine Sektion,
die Fortschritt erzählt, muss bei jedem Schritt nachgezogen werden und ist zwischendurch falsch.
Der Ist-Zustand dagegen ändert sich nur, wenn er sich ändert.

**Die beiden überzähligen Bullets tragen je einen eigenen Defekt:**

- *Problem-Quellen* nennt die vier `arch-check.sh`-Varianten — und sagt im letzten Satz selbst,
  sie definierten *„die **Anforderung** …, **nicht die Harness-Form**"*. Ein Eintrag unter
  *Konventions-Quellen*, der von sich sagt, keine zu sein, benennt seinen eigenen richtigen Ort.
- *Konventions-Vorbild* sagt, die `d-check`-Muster würden übernommen *„sobald die jeweiligen
  Slices sie anlegen"* — die haben sie mit slice-003 bis slice-005 angelegt. Stehengebliebene
  Zukunftsform über einen seit Monaten erreichten Zustand.

## 2. Betroffene Module

`harness/conventions.md`, drei Abschnitte. Eine Schicht.

## 3. Auszuführende Gates

`make gates` — tragend ist `doc-check`: die Kürzungen entfernen Verweise, und
[`MR-006`](../../../../harness/conventions.md#mr-006--baseline-committet-vendored-statt-per-url-referenziert)
adressiert `§Baseline` per Anker. Zum Abschluss `make verify`.

**Kein neuer Sensor.** Die Probe ist der Bestand: **niemand** verweist auf die drei Passagen —
gemessen, nicht angenommen; der einzige Fremdtreffer ist eine beiläufige Erwähnung des Wortes
„Konventions-Vorbild" in einer Closure-Notiz, kein Link.

## 4. Was bewusst nicht getan wird

- **Nichts wird „gerettet".** Die Problem-Quellen wandern **nicht** ins Lastenheft — Maintainer-
  Vorgabe: keine Chronik, auch nicht anderswo. `spec/lastenheft.md` §Zweck trägt den Zweck ohnehin;
  der gestrichene Bullet verweist selbst dorthin.
- **Begründungen bleiben.** Eine Adaption ohne ihr *Warum* wäre eine stille Setzung. Gestrichen
  wird, was erzählt, **was mit dem Dokument geschah** — nicht, warum eine Regel gilt.
- **Kein anderer Abschnitt.** §Anforderungs-Anlege-Prozess, §Zusatzklassen und
  §Modus-Deklaration tragen Regeln, keine Chronik; sie bleiben unberührt.

## 5. DoD

- [x] §Baseline nennt Konvention, Stand, vendored Ort und Integrität — ohne Etappen-Erzählung;
      die überholte Aussage „Offen ist allein D" ist damit weg — Beleg: Diff.
- [x] §Adoptierte Konventions-Quellen trägt die drei Bullets der Sollform; *Problem-Quellen* und
      *Konventions-Vorbild* sind ersatzlos gestrichen, und das *Extern*-Bullet ist auf den Zeiger
      gekürzt — der nachgeschriebene Baseline-Normtext steht im Vorlagen-**Kommentar**, nicht im
      Rumpf — Beleg: Diff.
- [x] Die Nummern-Frage-Forensik in der Adaptions-Präambel ist weg; `doc-check` bleibt bei 0
      Befunden — Beleg: Target-Ausgabe.

Pflicht, aber **kein** Liefer-Punkt: `make gates` und zum Abschluss `make verify` grün — Ausgabe
in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.

## 6. Closure-Notiz

**Geliefert:** `harness/conventions.md` von 240 auf 215 Zeilen; §Baseline nennt den Ist-Zustand
statt zweier Migrations-Erzählungen, §Adoptierte Konventions-Quellen trägt die drei Bullets der
Sollform statt fünf, und die Nummern-Frage-Forensik ist weg.

**Lerneintrag — Form: geschärfte Regel.** *Eine Sektion, die Fortschritt erzählt, ist zwischen
zwei Schritten immer falsch.* §Baseline behauptete bis zu diesem Slice „Offen ist allein **D**" —
geschrieben in slice-099, überholt seit slice-102, also **drei Slices später in derselben
Sitzung**. Niemand hatte es nachgezogen, und niemand hätte es gemerkt: die Aussage ist von keinem
Gate erfasst. *Weil* eine Fortschritts-Aussage bei jedem Schritt gepflegt werden muss, während der
**Ist-Zustand** sich nur ändert, wenn er sich ändert. **Der Prüfsatz:** in eine Datei, die jeder
Lauf liest, gehört, was **gilt** — nicht, was geschah; das Gewordene steht in `done/` und wird dort
gesucht, wenn jemand danach fragt.

**Zwei beobachtbare Closure-Kriterien:**

1. 240 → 215 Zeilen, und §Adoptierte Konventions-Quellen trägt genau die drei Bullets, die die
   Vorlage vorsieht (Extern · Vendored · In-Repo) — nachzählbar gegen
   `.harness/baseline/v5.12.0/templates/harness/conventions.template.md`.
2. `doc-check` prüft 218 Dateien mit 0 Befunden, **nachdem** Verweise entfernt wurden — belegt,
   dass keine der drei Passagen ein Ziel war. Vorab gemessen: kein Dokument verlinkt sie.

**Offene Risiken und ihr Ausgang:**

- *Dieselbe Chronik steht an je einer Stelle in `AGENTS.md` und `harness/README.md`* — gemessen,
  nicht vermutet: „Bis slice-077 stand hier …" und „bis slice-079 tat das `gate-consistency`".
  Ausgang: **weiter offen**, als `BEO-009` im Beobachtungs-Register.
- *Baseline-Normtext wird an einer weiteren Stelle nachgeschrieben* — `AGENTS.md` §1 trägt
  dieselbe Normativitäts-Klausel im Rumpf, die hier auf den Zeiger gekürzt wurde. Sie steht als
  Rumpftext **nur** im vendored `regelwerk/README.md`; in der Vorlage lebt sie in einem
  HTML-Kommentar, der beim Adoptieren wegfällt. Ausgang: **weiter offen**, als `BEO-010` im
  Beobachtungs-Register.
- *Gestrichenes ist nur in der Historie auffindbar* — Ausgang: **gestrichen mit Begründung**. Die
  drei Passagen sind Chronik über das Dokument selbst; ihr Ort ist `git log` und `done/`, und
  genau dorthin verweist §Baseline jetzt ausdrücklich.

**Folge-Slices:** keine — der Rest hängt an `BEO-009`.

## 7. Sub-Area-Modus

Berührt wird **Harness-Einstieg** — die Sub-Area, die slice-101 deklariert hat. Greenfield.
