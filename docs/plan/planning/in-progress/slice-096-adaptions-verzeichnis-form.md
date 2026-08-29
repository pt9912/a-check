# slice-096 — Etappe C1: Adaptions-Speicher in die Verzeichnis-Form

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Deckt:** keine `AC-*`/`ADR-*` — Formwechsel im Konventionsspeicher.
**Bezug:** Etappe **C** aus [slice-092 §6](../done/slice-092-regelwerk-v5120-delta-analyse.md),
erste Hälfte. Vorgänger [slice-095](../done/slice-095-adaptions-durchgang-v5120.md) (Etappe B).
Form-Entscheid des Maintainers 2026-08-29: **Verzeichnis-Form**.
[Roadmap](../in-progress/roadmap.md).

---

## 1. Auslöser

`grundlagen-harness-dateien.md` §Konventionsspeicher macht die Verzeichnis-Form zum Default:
ein Eintrag je Datei, `conventions.md` trägt nur noch den Index. Der Grund ist Kontext, nicht
Ästhetik — *„`conventions.md` liest **jeder** Agentenlauf"*, und aufgelöste Einträge lesen sich
wie geltende. Die Form ist laut Baseline **Wahl**; der Maintainer hat den Default gewählt.

Dieser Slice wechselt **nur die Form**. Die Urteile aus Etappe B werden hier **nicht** ausgeführt —
das ist C2, und die Trennung hat einen Grund: der Formwechsel ist anker-kritisch und muss für sich
prüfbar sein. Wandert er zusammen mit acht neuen Einträgen, sagt kein Gate mehr, woran ein
gebrochener Verweis lag.

## 2. Betroffene Module

- `harness/conventions.md` §Adaptions-Block — wird Index (zwei Tabellen).
- `harness/conventions/` — neun Eintrags-Dateien; `conventions/done/` für die eine bereits
  aufgelöste.

Eine Schicht, ein Konventionsspeicher.

## 3. Die anker-kritische Hälfte

Jeder bestehende Verweis im Repo adressiert eine Adaption über den **Überschriften-Slug** der
Inline-Form (`#mr-006--baseline-committet-vendored-statt-per-url-referenziert`). Verschwinden die
`### MR-NNN — Titel`-Überschriften, verschwinden ihre Auto-Anker mit ihnen.

Die Baseline schreibt die Gegenmaßnahme wörtlich vor: *„Kam dieses Repo von der Inline-Form, trägt
die Zeile den alten Überschriften-Slug als **zweiten** Anker daneben, sonst rotten die bereits
veröffentlichten Verweise."* Jede Index-Zeile bekommt also **zwei** Anker — die stabile Kennung
`mr-NNN` für neue Verweise, den alten Slug für die bestehenden.

Belegen kann das nur `doc-check`: es prüft jeden Link **und** jeden Anker. Der Bestand trägt
**137** solcher Verweise in **48** Dateien — gezählt, nicht geschätzt.

**[`MR-000`](../../../../harness/conventions.md#mr-000) bleibt inline.** Die Vorlage führt ihn ausdrücklich weiter im Index — er ist die
Adoptions-Erklärung, keine Adaption, und gilt für jeden Lauf. Sein Anker bleibt damit unberührt.

## 4. Auszuführende Gates

`make gates` — tragend ist `doc-check` (Links, Anker, Kennungs-Linkpflicht) und
`gate-consistency`. Zum Abschluss `make verify`.

**Kein neuer Sensor**, also keine Negativ-Probe. Die Probe ist der Bestand selbst: 137 vorhandene
Verweise auf die alten Slugs sind die schärfste Prüfung, die dieser Formwechsel bekommen kann —
sie waren vorher grün und müssen es danach sein.

## 5. Was bewusst nicht getan wird

- **Die B-Urteile werden nicht ausgeführt.** Kein Rückbau, keine Nachfolge-Einträge, keine
  Bewegung nach `done/` außer der einen, die schon vor diesem Slice aufgelöst war. Das ist C2.
- **Das Pflichtfeld `Ersetzt-Baseline-Regel` wird nicht in die Bestandseinträge geschrieben** —
  „Einträge werden nie überschrieben"
  ([slice-095 §4](../done/slice-095-adaptions-durchgang-v5120.md)). Der Index trägt die Spalte;
  wo der Bestandseintrag sie nicht hat, steht dort ausdrücklich `—` mit Verweis auf C2, statt eine
  Regel zu erfinden.
- **`.d-check.yml` bleibt unangetastet.** Das `ids`-Muster zeigt auf `harness/conventions.md`, und
  genau dort bleibt der Index — die Linkpflicht trifft weiterhin die richtige Datei.

## 6. DoD

- [x] Neun Eintrags-Dateien unter `harness/conventions/` (die aufgelöste unter `done/`),
      Inhalt unverändert übernommen; `conventions.md` trägt Präambel, [`MR-000`](../../../../harness/conventions.md#mr-000) inline und **zwei**
      Index-Tabellen — Beleg: Diff und Verzeichnis-Zustand.
- [x] Jede Index-Zeile trägt **beide** Anker (Kennung + alter Überschriften-Slug); `doc-check`
      meldet **0** Befunde, obwohl **137** Verweise auf die alten Slugs im Bestand stehen —
      Beleg: Target-Ausgabe mit Exit-Code.
- [x] `make gates` (und bei Abschluss `make verify`) grün — **Ausgabe in eine Datei**, Exit-Code
      getrennt geprüft, nie in eine Pipe.

## 7. Closure-Notiz

**Geliefert:** neun Eintrags-Dateien unter `harness/conventions/`, ein Index aus zwei Tabellen mit
Doppelankern, und **137 bestehende Verweise, die den Formwechsel überlebt haben**, ohne dass einer
von Hand angefasst wurde.

**Lerneintrag — Form: geschärfte Regel.** *Eine Kennung in ihrer eigenen Zieldatei ist von der
Linkpflicht befreit; verlässt sie diese Datei, wird die Pflicht scharf.* `.d-check.yml` bindet
`MR-\d{3}` an `harness/conventions.md` als `target` — solange die Einträge **in** dieser Datei
standen, brauchte keine der dortigen Erwähnungen einen Link. Der Umzug in eigene Dateien hat aus
sieben stillen Erwähnungen sofort sieben `id-unlinked` gemacht, *weil* dieselbe Kennung an einem
anderen Ort eine andere Regel trifft. Zwei davon waren **fremde** Kennungen (die des
Baseline-Templates und die von `ai-harness-init`) — sie zu verlinken wäre schlimmer als sie
stehenzulassen, denn ein Link hätte sie zu a-checks Adaptionen erklärt. Richtig ist dort der
`d-check:ignore`-Marker mit Grund, nicht der Link. **Die Prüfsatz-Form:** *Wandert Text aus der
Zieldatei einer `ids`-Regel heraus, ist vor dem Verlinken zu fragen, wessen Kennung das ist.*

**Zwei beobachtbare Closure-Kriterien:**

1. `doc-check` prüft 203 Dateien mit **0** Befunden; die Zahl der Verweise auf alte
   Überschriften-Slugs ist mit `grep` als **137** in **48** Dateien belegt. Beide Anker sind
   also nachweislich in Gebrauch — `doc-check` meldete `anchor-missing`, solange einer fehlte.
2. `harness/conventions.md` trägt keine `### MR-NNN`-Überschrift mehr außer der von
   [`MR-000`](../../../../harness/conventions.md#mr-000); jeder Agentenlauf liest ab jetzt eine
   Tabellenzeile je aktiver Adaption statt acht Volltexte.

**Offene Risiken und ihr Ausgang:**

- *Die B-Urteile sind weiter nicht ausgeführt* — Ausgang: **Folge-Slice**, Etappe C2.
- *Die Spalte `Ersetzt-Baseline-Regel` steht überall auf `—`* — Ausgang: **Folge-Slice**, C2; sie
  darf per Append-only-Regel nur in Nachfolge-Einträgen entstehen.
- *Die Schätzung „über 30 Verweise" im geschnittenen Slice war um Faktor vier zu niedrig* —
  Ausgang: **gestrichen mit Begründung**; die Zahl ist vor dem Abschluss gemessen und im Text
  ersetzt worden, das Risiko hat sich damit erledigt statt fortzubestehen.

**Folge-Slices:** Etappe C2 (Urteils-Ausführung), danach C3 und D.

## 8. Sub-Area-Modus

Berührt wird der Konventionsspeicher — weiterhin ohne eigene Zeile in der Modus-Deklaration pro
Sub-Area ([slice-091 §7](../done/slice-091-claude-md-auf-verweis-reduzieren.md)). Alle berührten
Sub-Areas mit Modus sind GF.
