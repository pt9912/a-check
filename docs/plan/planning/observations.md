# Beobachtungs-Register

Regeln dieses Registers: Baseline-Regelwerk `modul-06-roadmap.md`
§Das Beobachtungs-Register — wer schreibt, wer liest, wann gestrichen wird, welche Form ein Beleg
hat, und dass eine leere Tabelle `— keine —` trägt statt zu verschwinden.

**Wer schreibt:** die **Slice-Closure** — neue Kennung vergeben oder Zähler erhöhen und Beleg
ergänzen. Damit läuft der Zähler mit jedem geschlossenen Slice und nicht mit der Welle; in diesem
Repo ist das wesentlich, weil gerade keine Welle offen ist.
**Wer liest:** die Welle-Closure, was **3×** erreicht hat — und die Slice-Planung (§8 des Plans),
was darunter steht. Wer nur den ersten Schritt kennt, sieht alles unter 3× nie wieder an.

**Belege sind formgebunden:** Slice-Kennung `slice-<NNN>`, kein Freitext; so viele, wie der Zähler
sagt. Angelegt mit slice-101.

| Kennung | Beobachtung | Sub-Area | Zähler | Belege | Stand |
|---|---|---|---|---|---|
| BEO-001 | Die Harness-Einstiegs- und Konventions-Dateien haben keine Zeile in der Modus-Deklaration pro Sub-Area | Harness-Einstieg | 4× | slice-091, slice-095, slice-096, slice-097 | **verkörpert** in [`harness/conventions.md` §Modus-Deklaration pro Sub-Area](../../../harness/conventions.md#modus-deklaration-pro-sub-area) (`seit slice-101`) — Schwelle beim Erstauftreten im Register bereits überschritten |
| BEO-002 | „Höchstens zwei Schichten" der Größen-Regel bleibt maschinell ungeprüft; was eine Schicht ist, ist Ermessen über Modul-Grenzen | Planungs-Harness | 1× | slice-098 | offen — ein Zähler darüber wäre Schein-Genauigkeit; bleibt Sache des Reviews |
| BEO-003 | Ein Rückbau-Eintrag, der erklärt, dass **nicht** mehr abgewichen wird, liegt von Beginn an in `conventions/done/` statt in der aktiven Tabelle — Auslegung, keine Vorschrift | Harness-Einstieg | 1× | slice-097 | offen — überstimmbar; die Begründung steht in slice-097 §3 |
| BEO-004 | Ein Adaptions-Eintrag korrigiert eine Repo-Aussage statt einer Baseline-Regel und ist damit unter dem Fork-Test ein Rückbau-Kandidat | Harness-Einstieg | 1× | slice-097 | offen — auflösbar erst mit der Überarbeitung der ID-Schema-Deklaration; der Eintrag trägt den Trigger selbst |
| BEO-005 | Die Platzhalter-Liste des Closure-Struktur-Gates enthält eine Wendung, die **unzitiert** legitim in einem Risiko-Block steht | Gate-/Werkzeug-Schicht | 2× | slice-099, slice-100 | offen — Streichen wäre eine Schwellen-Senkung und braucht nach [`AGENTS.md`](../../../AGENTS.md) §3.6 eine ADR; entschieden wird am nächsten unzitierten Vorfall. Die Liste liegt seit slice-080 als `forbid-pattern` in [`.d-check.yml`](../../../.d-check.yml), nicht mehr im Skript |
| BEO-006 | Die vorgeschriebene Reihenfolge kann die Prüfung nicht abnehmen, die sie erfüllen soll: das Closure-Struktur-Gate greift nur in `done/`, der Workflow fährt `make verify` in Schritt 8 und den `git mv` erst in Schritt 9 | Gate-/Werkzeug-Schicht | 1× | slice-099 | offen — die Closure-Notiz eines Slice wird frühestens beim **nächsten** geprüft; die Ablösung durch `doc-structure` (slice-080) ändert daran nichts, die Regel trägt denselben `files`-Glob |
| BEO-007 | Der Risiko-Ausgangs-Sensor prüft **innerhalb** eines vorhandenen Blocks, nicht seine Existenz — wer den Block weglässt, wird nicht erwischt | Gate-/Werkzeug-Schicht | 1× | slice-102 | offen — die Existenz einzufordern verlangte ein Urteil (gab es hier überhaupt Risiken?) und erzeugte auf jedem älteren Slice Rauschen |
| BEO-008 | Ein Verweis **auf** einen wandernden Slice bleibt ungeprüft und bricht, wenn dieser wandert — die Lifecycle-Invariante ist auf wandernde Quellen gerichtet, nicht auf Verweise auf sie | Planungs-Harness | 3× | slice-093, slice-114, slice-080 | **verkörpert** in `make slice-mv` (`tools/slice-mv.sh`, seit slice-118): der `git mv` zieht die Verweise selbst nach. Der Zähler bleibt stehen — die Beobachtung ist nicht falsch geworden, sie hat einen Ort bekommen. Ungedeckt bleibt, wer `git mv` von Hand fährt; der Unterschied ist, dass der richtige Weg jetzt **kürzer** ist als der falsche |
| BEO-009 | Chronik in Dateien, die jeder Agentenlauf liest: eine Sektion, die Fortschritt erzählt, ist zwischen zwei Schritten falsch und kostet bei jedem Lauf Kontext | Harness-Einstieg | 1× | slice-103 | offen — in `conventions.md` behoben; `AGENTS.md` §5 und `harness/README.md` §Sensors tragen je eine gemessene Reststelle |
| BEO-010 | Baseline-Normtext wird im Repo **nachgeschrieben** statt verwiesen — zwei Fassungen driften | Harness-Einstieg | 1× | slice-103 | offen — in `conventions.md` auf den Zeiger gekürzt; `AGENTS.md` §1 trägt dieselbe Klausel weiterhin im Rumpf |
| BEO-012 | Ein Review kann seinen Geltungsbereich begründen und ihn trotzdem zu eng ziehen; beide bisherigen Lücken fand der Maintainer, nicht der Review | Planungs-Harness | 1× | slice-105 | offen — zwei Verfahren mit gegenläufigem Irrtum sind besser als eines, aber kein Beweis |
| BEO-013 | Die Meilenstein-Tabelle der Roadmap führt keine `Trigger`-Spalte — Abweichung schon gegenüber `v3.5.2`, nicht Migrationsfolge | Planungs-Harness | 1× | slice-106 | offen — Füllen hieße, Trigger für drei erreichte Meilensteine zu erfinden |
| BEO-014 | Die Marker-Hälfte von *Offene Wellen* ist ungewächtert: `doc-planning` lief grün, während der Ruhe-Marker bei beanspruchtem `in-progress/` stand | Gate-/Werkzeug-Schicht | 2× | slice-106, slice-115 | offen — das Modul `planning` läuft ohne Konfigurationsblock in `.d-check.yml`; ein aktiviertes Modul ohne Konfiguration meldet grün, statt zu schweigen. Der zweite Fall (`doc-structure` ohne `structure`-Block) ist mit slice-080 **entfallen** — der Block existiert; für `planning` steht die Beobachtung unverändert |
| BEO-015 | Vier Begriffe der Slice-Ziel-Form sind nicht adoptiert: `Welle:`-Feld, Reconciliation-Register, *drei Paarungen*, Herkunfts-Anker | Planungs-Harness | 1× | slice-107 | offen — je eine eigene Entscheidung nötig; vier Mechaniken ungeprüft einzuführen wäre schlimmer |
| BEO-016 | Hard Rule 3.7 ist **inferentiell** und hat keinen Sensor; der Bestand ist nicht gegen sie durchgesehen | Harness-Einstieg | 1× | slice-108 | offen — ein Zähler über Kommentar-Klassen wäre Schein-Genauigkeit; die Regel hängt am Review, und das ist in `AGENTS.md` §3.7 benannt |
| BEO-017 | Vier Sektionen der Spec-Straten fehlen inhaltlich: *Globale Out-of-Scope-Punkte*, *Glossar*, *Externe Abhängigkeiten*, *Fehlermodelle und Resilienz* | Spec-Straten | 1× | slice-110 | offen — **Maintainer-Inhalt**; eine erfundene Vertragsaussage sieht abgenommen aus |
| BEO-018 | `spec/spezifikation.md` gliedert nach Vertrags-Kennung statt nach den sieben Themen der Ziel-Form — **undeklariert**, in einer kanonischen Quelle vom Rang 2 | Spec-Straten | 1× | slice-110 | offen — entweder Adaptions-Eintrag mit benannter ersetzter Regel oder Angleichung; beides eine Entscheidung, keine Messung |
| BEO-019 | Form-Vergleiche gegen die Vorlagen sind sprachblind: `README.md` ist englisch, die Ziel-Form deutsch — beide Verfahren meldeten drei fehlende Sektionen statt einer | Planungs-Harness | 1× | slice-111 | offen — weder exakter Titel- noch Kernbegriff-Vergleich überbrückt einen Sprachwechsel; jeder künftige Form-Review trifft dieselbe Wand |
| BEO-020 | Die Ziel-Form liegt nur noch **tag-gescopt** unter `.harness/baseline/<tag>/templates/` — jeder Baseline-Sprung verschiebt die Verweise darauf | Planungs-Harness | 1× | slice-112 | offen — das ist der Preis, den die co-located-Regel vermeiden wollte; drei Verweise, klein und benannt |
| BEO-021 | Das Modul `sources` ist mit dem Pin `v0.67.0` erreichbar und deckt genau die Asset-Integrität der vendored Baseline, die slice-047 offengelassen hat | Gate-/Werkzeug-Schicht | 1× | slice-115 | offen — die `SHA256SUMS` neben `.harness/baseline/<tag>/` werden von keinem Gate gelesen; ob `sources` das trägt, ist eine eigene Messung und nicht der Pin-Bump |
| BEO-022 | Eine messbare Behauptung in einem CR wird behauptet statt gemessen — gegen den **eigenen Bestand** (CR 3: 19 von 19 ausgenommen, danach bleibt keiner), gegen den **fremden Vertrag** (CR 4/1: nicht-fatale Meldung vorausgesetzt, die `DC-FA-CLI-003` als Out-of-Scope führt) oder gegen den **eigenen Gate-Pfad** (CR 4/2: der Preis der `--doctor`-Wahl „schwächer" genannt, gemessen ist er null). Zweite Ausprägung, beim Empfänger aufgetreten und dort eigener Zähler: die **falsche Menge** gemessen — sieht aus wie ein Beleg und kann zutreffen | Gate-/Werkzeug-Schicht | 3× | slice-080, slice-116, slice-116 | **verkörpert** im Skill `.harness/skills/cr-text-reviewer.md` (seit slice-119), verwiesen aus [`AGENTS.md`](../../../AGENTS.md) §5 und dem Workflow-Skelett. Kein Sensor möglich: ob ein Satz gemessen wurde, ist ein Urteil über seinen Entstehungsweg. Ungedeckt bleibt, wer den Skill nicht aufruft |

## Gestrichene Einträge

Regeln dieser Sektion: Baseline-Regelwerk `modul-06-roadmap.md`
§Das Beobachtungs-Register — wer eine Zeile still löscht, macht sie ununterscheidbar von einer,
die es nie gab.

| Kennung | Beobachtung | Gestrichen am | Warum sie nicht mehr auftreten kann |
|---|---|---|---|
| BEO-011 | Vier Form-Vergleiche ungeprüft | 2026-08-29 | Alle vier sind gelesen (slice-108, slice-110); was dabei als echte Lücke herauskam, steht als `BEO-017` und `BEO-018` mit eigenem Gegenstand. Die Beobachtung „ungeprüft" kann nicht wieder auftreten, weil sie sich auf eine geschlossene Menge von vier Zeilen bezog. |
