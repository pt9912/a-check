# slice-186 — Voll-Abgleich: `AGENTS.md` §5 und drei Ziel-Form-Paare

**Welle:** ohne Welle — die Closure-Bedingung wäre die eigene DoD.

**Bezug:** Folge-Slice aus
[slice-185](../done/wellenlos/slice-185-ziel-form-voll-abgleich.md) §7, Risiko 3
(*eingetreten*: 13 Paare sprengen die Größen-Regel).
[`AC-QA-02`](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze).

**Berührte Spec-Stellen:** — · Der Slice berührt kein Spec-Stratum. Die zwei
Straten (`spec/lastenheft.md`, `spec/spezifikation.md`) stehen unter den sieben
an [slice-187](../open/slice-187-voll-abgleich-erstdurchgang-rest.md)
abgetretenen Paaren; §9 führt `SPEC` folgerichtig als *nicht berührt*.

**Verantwortlich:** Claude — gesetzt beim Übergang nach `in-progress/`.

**Autor:** Claude. **Datum:** 2026-09-08.

**Lerneintrag — Form:** geschärfte Regel.

---

## 1. Ziel und Abgrenzung

**Ziel — geändert während der Arbeit, siehe §3.3.** Ursprünglich: *„die zehn
Ziel-Form-Paare, die
[slice-185](../done/wellenlos/slice-185-ziel-form-voll-abgleich.md) nicht mehr
getragen hat, sind abgeglichen"*. Geliefert: die drei für `AGENTS.md` vorab
gemessenen Befunde (§2.1) sind aufgelöst, und **drei** Paare sind abgeglichen —
je Paar mit einem der drei Ausgänge: übernommen · bewusst abweichend (mit
Begründung) · ohne Befund. Die sieben verbleibenden gehen an
[slice-187](../open/slice-187-voll-abgleich-erstdurchgang-rest.md); der Titel
und dieser Satz sind der Plan-Änderung nachgezogen, statt die alte Zusage
stehen zu lassen.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Die fünf in slice-185 abgeglichenen Paare.** Erledigt; ein zweiter Durchgang
  wäre Arbeit ohne Gegenstand.
- **Die zwölf Instanz-Vorlagen.** Dieselbe Abgrenzung wie in slice-185 §1: kein
  einzelnes Gegenstück, Form geprüft durch `make doc-structure`.
- **Ein Sensor auf Satz-Deckung.** Urteil über zwei Formulierungen, kein Match
  ([`AGENTS.md`](../../../../AGENTS.md) §3.7).

## 2. Ausgangsmessung (2026-09-08, aus slice-185 übernommen)

Kandidaten je Paar aus der maschinellen Vorauswahl — **überwiegend Rauschen**
(Platzhalter, Bedienhinweise), die Zahl ist eine Reihenfolge, kein Befund:

| Paar | Kandidaten |
|---|---|
| `README.md` | *(neu zu erheben, siehe unten)* |
| `AGENTS.md` | 22 |
| `harness/conventions.md` | 18 |
| `spec/lastenheft.md` | 18 |
| `.d-check.yml` | 16 |
| `.harness/skills/closure-note-reviewer.md` | 16 |
| `docs/plan/planning/README.md` | 12 |
| `spec/spezifikation.md` | 12 |
| `harness/README.md` | 9 |
| `docs/plan/planning/in-progress/roadmap.md` | *(neu, siehe unten)* |

**Ein zehntes Paar** (Review zu slice-185, F-7): `roadmap.template.md` hat ein
Gegenstück in [`docs/plan/planning/in-progress/roadmap.md`](../in-progress/roadmap.md).
Die Paarbildung von slice-185 suchte es unter `docs/plan/planning/roadmap.md`
und fand nichts — die Roadmap liegt im Lifecycle-Verzeichnis.

**Reihenfolge:** Die Kandidaten-Zahl ist eine grobe Sortierung, kein Befund.
Der Indikator *„Gegenstück kürzer als Vorlage"*, der in slice-185 beim ersten
Treffer trug, greift hier bei **keinem** Paar — er hatte dort genau einen
echten Fall (den Reviewer-Skill), und der ist erledigt.
**`README.md` ist gegen `project-readme.template.md` zu messen** (2076 Zeichen),
nicht gegen `templates/README.md` — das ist der Index des Vorlagen-Verzeichnisses
und keine Ziel-Form (Review zu slice-185, F-1). a-checks `README.md` ist mit
9274 Zeichen **länger** als seine Vorlage; die 32 Kandidaten der Vorauswahl
stammen aus dem falschen Vergleich und sind neu zu erheben.

### 2.1 Für `AGENTS.md` stehen drei Befunde bereits fest

Sie sind nach slice-185 gemessen worden und ersparen dem Paar die Suche. Alle
drei betreffen **§5 Dokumentations-Regeln** (16 740 Zeichen, 18 Blöcke) — mit
Abstand den größten Abschnitt der Datei.

**(a) Drei nachgeschriebene `modul-05`-Sätze im Slice-Form-Block** (**834**
Zeichen — gemessen, 4285 → 3451). Ein Wortfolgen-Test über alle 18 Blöcke gegen das vendorte Regelwerk
findet **fünf** wörtliche Treffer, alle im selben Block; **zwei** davon sind
legitime Zitate mit Anführungszeichen und Quelle (`modul-06` §Wann Arbeit eine
Welle braucht), **drei** stehen ohne beides und stammen aus `modul-05`
§Ziel-Form: Slice — *„ein Ausschluss ohne Grund ist eine Behauptung, keine
Grenze"*, *„wer später etwas mitnimmt … hat den Plan geändert, nicht ergänzt"*,
*„zerlegt, nicht gedehnt … eine von drei benannten Formen"*.
**Die anderen 17 Blöcke: null Treffer.**
*Warum es slice-183 entging:* Der Block wurde dort als a-check-eigen eingestuft
(*„sieben Punkte, jeder gegen den Bestand gemessen"*) — das stimmt für die
sieben Punkte; nachgeschrieben sind die **Rahmen-Sätze** darum. Er ist der
einzige Block, in dem beides gemischt ist.

**(b) Zwei Review-Urteilsregeln am falschen Ort** (**4045** Zeichen):
*Geltungsbereich einer Messung* (2688) und *Eine Mutations-Probe belegt erst,
wenn sie rot war* (1357). Beide sagen ausdrücklich **„Kein Sensor"** — sie sind
inferentielle Urteilsregeln, und `modul-08` §Welche Rolle braucht welche
Artefaktklasse weist genau diesen Fall der **Skill-Datei** zu, nicht dem
Briefing. Ihr Ort wäre
[`.harness/skills/reviewer.md`](../../../../.harness/skills/reviewer.md).

**(c) Die Slice-Form-Kopieranleitung** (**4285** Zeichen, ein Viertel von §5).
Sie ist eine Bedienungsanleitung für die Vorlage und steht in `AGENTS.md`, weil
die vendorte Ziel-Form nicht geändert werden darf und a-check keine eigene
Kopie führt. Ihr natürlicher Ort wäre
[`docs/plan/planning/README.md`](../README.md), wo die Slice-Ablage beschrieben
ist.

**Zusammen 8930 Zeichen**, die §5 auf gut die Hälfte brächten — **ohne dass eine
Regel verlorengeht**, nur an den Ort wandert, an dem sie gelesen wird. (b) und
(c) sind **Umzüge**, kein Rückbau; das ist beim Abgleich sauber zu trennen.

**Geltungsbereich dieser Messung — und was daran nicht trägt.** Der
Wortfolgen-Test sucht **wörtliche** Übernahme; eine sinngemäße Doppelung in
anderen Worten sieht er nicht (slice-185 §2.3). Das war benannt. **Nicht
benannt war der Parameter, und daran hängen die Zahlen:** Fenstergröße, Korpus
und Normalisierung standen nirgends, und das Instrument liegt nicht im Repo —
der Lauf ist damit **nicht wiederholbar**. Der unabhängige Review hat
gegengemessen: Bei 8- bis 10-Wort-Fenstern über `regelwerk/*.md` hält die
Aussage (drei Treffer, alle in Block 13, die anderen Blöcke null). Bei 6- bis
7-Wort-Fenstern sind es sechs bzw. vier in Block 13 **und je einer in Block 18**
(*„nicht ‚fertig', sondern nur ‚weg'"* gegen `modul-01`); mit `templates/` im
Korpus kommen zwei weitere Blöcke dazu. **Die Zahl fünf und der Satz „die
anderen 17 Blöcke: null" gelten also nur unter einem Parameter, den diese
Messung nicht angab.** Der Befund (a) selbst bleibt davon unberührt — die drei
nachgeschriebenen Sätze stehen unter jedem Parameter da. Dieselbe Lücke haben
die Kandidaten-Zahlen (5/4/9 und 20/17/16/16/15/12/11, Summe 107): Sie sind
eine **Reihenfolge**, kein Befund, und stammen aus einem Lauf, den niemand
wiederholen kann. Das ist die eigene Mess-Regel
([`AGENTS.md`](../../../../AGENTS.md) §5, *Geltungsbereich einer Messung*), an
diesem Slice verletzt und vom Review gefangen.

## 3. Umsetzung

### 3.1 Die drei vorab gemessenen `AGENTS.md`-Befunde (§2.1) — aufgelöst

**§5 von 16 740 auf 10 214 Zeichen, 39 % kleiner.** Keine Regel geht verloren;
zwei sind **umgezogen**, eine ist auf Zeiger gekürzt.

**Der Zwischenstand war 9491 (43 %) — der Review hat ihn korrigiert** (F-1):
Beim Umzug der zwei Mess-Regeln ging nicht nur die Herleitung, sondern auch die
**Zusage** mit, und die bindet den *Schreibenden* — Slice-Plan, Closure-Notiz,
Review-Report —, der den Reviewer-Skill laut
[`harness/README.md`](../../../../harness/README.md) §Guides gerade **nicht**
liest. Der Zeiger trägt die zwei Sätze jetzt selbst (1267 statt 430 Zeichen);
die Fälle, die Register-Anker und die Begründung bleiben im Skill. **Zwei
Adressaten, zwei Orte** — das ist die Lehre, nicht die Zeichenzahl.

| Befund | Vorher | Nachher | Wohin |
|---|---|---|---|
| (a) drei nachgeschriebene `modul-05`-Stellen | Slice-Form-Block 4285 | 3451 → dann umgezogen | Zeiger auf `modul-05` §Ziel-Form: Slice |
| (b) zwei Mess-Regeln | 4045 | Zeiger **1267** | [`.harness/skills/reviewer.md`](../../../../.harness/skills/reviewer.md) §Mess-Regeln |
| (c) Slice-Form-Kopieranleitung | 3451 | Zeiger **539** | [`docs/plan/planning/README.md`](../README.md) §Beim Kopieren der Slice-Ziel-Form |

Alle Zahlen sind `wc -m` über den Abschnitt von `## 5.` bis ausschließlich
`## 6.` bzw. über den Bullet-Block — keine Schätzung.

**Bei (a) bleibt die repo-eigene Hälfte stehen:** die **Durchsetzung** —
`make verify` ab slice-052, der Gate-Lauf als feste Zeile statt Checkbox, die
Kopffelder ab slice-098. Die Baseline sagt *was* gilt, a-check *dass* und *wie*
es geprüft wird.

**Bei (b) und (c) ist es ein Umzug, kein Rückbau** — Risiko 2 aus §7. Nach jedem
Schritt trägt der Zielort die Regel **vollständig** (nur die Pfade eine Ebene
tiefer), und §5 einen Zeiger darauf.

### 3.2 Drei Paare abgeglichen

| Paar | Kandidaten | Ausgang |
|---|---|---|
| `README.md` | 5 | **übernommen:** die Rollen-Zeile, in beiden Sprachfassungen |
| `docs/plan/planning/in-progress/roadmap.md` | 4 | **ohne Befund** auf Abschnitts-Ebene: alle **sechs** H2 der Ziel-Form vorhanden (`modul-06` nennt fünf, die Vorlage führt den *Abhängigkeitsgraphen* zusätzlich). **Ein Befund darunter:** siehe §3.4 |
| `harness/README.md` | 9 | **ohne Befund** — die Treffer sind Bedienhinweise (*„Pointer-Artefakt … zuletzt füllen"*), kein Normtext |

**Die Rollen-Zeile mit a-checks Rang.** Die Ziel-Form sagt *„Rang 6 der Source
Precedence"*; a-check hat **neun** Ränge, seit `docs/user` eingefügt wurde
([`MR-003`](../../../../harness/conventions/done/MR-003-source-precedence-ohne-docs-user.md)),
und `README.md` steht auf **7**. Übernommen ist die *Aussage* — verweist auf die
kanonischen Quellen, dupliziert sie nicht —, nicht die Zahl. **Eine Ziel-Form
wörtlich zu kopieren wäre hier falsch gewesen.**

### 3.3 Die sieben verbleibenden Paare — abgetrennt, nicht gedehnt

**107 substanzielle Kandidaten**, Bedienhinweise und Platzhalter
herausgerechnet:

| Paar | substanzielle Kandidaten |
|---|---|
| `AGENTS.md` (Rest nach 3.1) | 20 |
| `spec/lastenheft.md` | 17 |
| `harness/conventions.md` | 16 |
| `.d-check.yml` | 16 |
| `.harness/skills/closure-note-reviewer.md` | 15 |
| `spec/spezifikation.md` | 12 |
| `docs/plan/planning/README.md` | 11 |

Zum Vergleich: slice-185 trug fünf Paare mit 14 Kandidaten, dieser Slice drei
mit 18 plus drei vorab gemessene Befunde. Sieben Paare mit 107 sind der
**sechsfache** Bestand — die Größen-Regel trägt das nicht.

**Adresse:** [slice-187](../open/slice-187-voll-abgleich-erstdurchgang-rest.md),
angelegt mit diesem Slice. Er nimmt den Punkt an: §1 nennt genau diese sieben
Paare, §2 trägt die Messung, §5 die Rückführung, falls auch sieben zu viel sind.

### 3.4 Was der unabhängige Review verändert hat

Der Report
([`2026-09-08-slice-186-…`](../../../reviews/2026-09-08-slice-186-voll-abgleich-restliche-paare.md))
trug **4 HIGH · 3 MEDIUM · 4 LOW · 1 INFO** und war merge-blockierend. Alle
zwölf sind abgearbeitet; die vier HIGH einzeln:

| Befund | Was daran stimmte | Behebung |
|---|---|---|
| **F-1** Mess-Regeln am Ort, den ihr Adressat nicht liest | Beide Regeltexte adressieren den *Schreibenden*; `harness/README.md` §Guides führt den Reviewer-Skill als *„nicht Teil der Implementer-Eingabe"* | §5 trägt die **Zusage** wieder (zwei nummerierte Sätze), der Skill die **Herleitung**; die Trennung ist an beiden Orten benannt |
| **F-2** vier Selbstverweise ins Leere | Der Block war *wortgleich* umgezogen — und genau darin lag der Defekt: zweimal unqualifiziertes `(§3.7)`, *„an §5 dieser Datei"*, *„die Aufgabe dieser Datei"* | alle vier auf `AGENTS.md` re-verankert; im Skill steht jetzt **kein** unqualifiziertes `§` mehr |
| **F-3** Herkunfts-Anker `seit slice-183` nicht am Zielort | `state.md` von [`BEO-HARNESS/baseline-normtext-nachgeschrieben`](../observations/BEO-HARNESS/baseline-normtext-nachgeschrieben/observation.md) nannte `AGENTS.md` §5 | Zielort auf den Skill umgestellt; die zwei anderen `state.md` (`review-geltungsbereich-zu-eng`, `probe-liefert-den-gegenstand-mit`) stimmen durch die F-1-Behebung **wieder von selbst** |
| **F-4** Roadmap sagt *„In Arbeit: slice-169"* | slice-169 liegt seit **siebzehn** Übergängen in `done/`; das Paar war als *ohne Befund* geschlossen, weil nur H2-Überschriften verglichen wurden | Der Name ist **entfernt**, nicht aktualisiert: `modul-06` §Wann Arbeit eine Welle braucht verbietet die zweite Quelle. Es steht ein Zeiger auf `ls in-progress/`, der nicht altern kann |

**Warum F-4 siebzehn Slices überlebt hat, ist der eigentliche Fund** — und er
ist ein Doppelter: `make slice-mv` zog bei jedem Übergang die **Adresse** nach
(der Link zeigte korrekt nach `done/`), nicht die **Aussage** daneben. Und
`make doc-planning` konnte es nicht fangen, weil sein Vertrag mehr zusagte
(*„benennt ihn"*), als das Modul prüft (Äquivalenz *Slice vorhanden ⟺ Marker
fehlt*) — die Einschränkung stand ausschließlich als Kommentar in
[`.d-check.yml`](../../../../.d-check.yml). **Das ist die HIGH-Kategorie *„Norm
nur im Konfigurations-Kommentar"* aus dem Reviewer-Skill, im eigenen Bestand
belegt.** Behoben an beiden Enden: [`harness/sensors/doc-planning.md`](../../../../harness/sensors/doc-planning.md)
führt die Grenze jetzt als ersten Punkt, und die zwei Deklarations-Stellen
([`AGENTS.md`](../../../../AGENTS.md) §4, `harness/README.md` §Sensors) sagen
zu, was geprüft wird.

**Die übrigen acht** in Kürze: `README.de.md` fehlte die Ziel-Form-Sektion *Was
kann ich heute tun?* (F-5, seit slice-111 — die Datei war im Übernahme-, nicht
im Befund-Scope; jetzt tragen beide Fassungen dieselben zehn H2) · der
Wortfolgen-Test nennt seinen Geltungsbereich (F-6, §2.1) · Titel, §1 und
Kopffeld sind der Plan-Änderung nachgezogen (F-8) · Tilde-Werte durch `wc -m`
ersetzt (F-9) · *fünf* → **sechs** Ziel-Form-Abschnitte der Roadmap (F-10) ·
zwei Zeiger auf ausgezogenen Text nachgezogen (F-11) · F-12 ist INFO ohne
Handlungsbedarf (Nachtrag eines Erstauftretens, beide urteilsfreien Prüfungen
bestehen).

## 4. Definition of Done

- [x] **Drei** der zehn Paare sind abgeglichen; je Paar steht der Ausgang im
      Plan (§3.2). **Plan-Änderung, benannt statt stillschweigend:** Die DoD
      verlangte *alle zehn*. Die sieben verbleibenden tragen **107**
      substanzielle Kandidaten (§3.3) — Risiko 1 aus §7 ist eingetreten, und der
      Ausgang ist ein Folge-Slice mit Kennung
      ([slice-187](../open/slice-187-voll-abgleich-erstdurchgang-rest.md)), kein
      gedehnter Slice. **Zum zweiten Mal derselbe Ausgang**; das steht als
      Beobachtung im Register (2×, Schwelle bei 3×).
- [x] Die drei für `AGENTS.md` **vorab gemessenen** Befunde (§2.1) sind
      aufgelöst — die drei `modul-05`-Sätze auf Zitat oder Zeiger, die zwei
      Urteilsregeln in den Reviewer-Skill, die Kopieranleitung nach
      `docs/plan/planning/README.md`. **Umzug, kein Rückbau:** nach dem
      Verschieben trägt der Zielort die Regel vollständig, und §5 einen Zeiger.
      Gegenprobe: §5 gemessen vor und nach dem Umzug.
- [x] Jede Übernahme ist als solche kenntlich; jede bewusste Abweichung trägt
      ihre Begründung — eine Differenz ist nicht automatisch eine Lücke.
- [x] Unabhängiger Review durchgeführt (Report unter [`docs/reviews/`](../../../reviews/README.md)) — 12 Findings, alle abgearbeitet (§3.4).
- [x] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [x] Beobachtungs-Register fortgeschrieben.
- [x] Jedes Risiko aus §7 trägt einen Ausgang.

`make gates` und `make verify` grün.

## 5. Trigger

**Start** (`open` → `in-progress`):
[slice-185](../done/wellenlos/slice-185-ziel-form-voll-abgleich.md) liegt in
`done/` und das WIP-Limit ist frei.

**Rückführungen — vorab benannt:**

- `in-progress` → `next` (zu groß): Braucht ein einzelnes Paar mehr als einen
  Durchgang — `spec/lastenheft.md` mit 84 146 Zeichen ist der Kandidat —, wird
  es abgetrennt.
- `in-progress` → `open` (blockiert): Findet der Abgleich eine Ziel-Form-Regel,
  die a-check bewusst **nicht** übernehmen will, ist das eine `MR`-Adaption und
  keine Nachtrags-Entscheidung dieses Slice.

## 6. Closure-Trigger

DoD vollständig, `make gates` und `make verify` grün, Closure-Notiz mit
Lerneintrag.

## 7. Risiken und offene Punkte

- **Zehn Paare könnten wieder zu groß sein.** slice-185 hat vier geschafft und
  das fünfte ohne Befund geschlossen; zehn ist mehr als das Doppelte — und mit
  §2.1 kommt für `AGENTS.md` ein Umzug von 8930 Zeichen dazu, der für sich
  genommen ein Slice wäre. — **Ausgang:** *eingetreten* → Folge-Slice
  [slice-187](../open/slice-187-voll-abgleich-erstdurchgang-rest.md). Der Slice
  trägt die drei `AGENTS.md`-Befunde **vollständig** und **drei** der zehn
  Paare; die sieben verbleibenden haben **107** substanzielle Kandidaten, gegen
  14 und 18 in den zwei Vorgängern. **Zum zweiten Mal derselbe Ausgang** — das
  ist selbst der Befund, und er steht im Register bei 2×
  ([`BEO-PLAN/erstdurchgang-als-einmal-slice-geschnitten`](../observations/BEO-PLAN/erstdurchgang-als-einmal-slice-geschnitten/observation.md));
  slice-187 wäre der dritte.
- **Ein Umzug kann eine Regel unterwegs verlieren.** (b) und (c) verschieben
  Text zwischen Dokumenten verschiedener Ränge; ein Zeiger, der die Aussage
  nicht trägt, ist schlechter als die Doppelung. — **Ausgang:** *entfallen*,
  gestrichen mit Begründung: Beide Umzüge sind **vollständig** — der Zielort
  trägt den Text unverändert (nur die Pfade eine Ebene tiefer), und §5 trägt
  einen Zeiger, der den Ort **und** den Grund nennt. Bei (a) ist es kein Umzug,
  sondern eine Kürzung auf einen Zeiger — dort bleibt a-checks Hälfte (die
  Durchsetzung) ausdrücklich stehen. Das Risiko war der Grund, je Schritt zu
  prüfen, was am Zielort ankommt.
- **Die maschinelle Vorauswahl ist unscharf in beide Richtungen.** Sie meldet
  Rauschen — und sie kann eine echte Lücke verfehlen, wenn a-check dieselbe
  Aussage in anderen Worten trägt. — **Ausgang:** *eingetreten* → Folge-Slice
  [slice-187](../open/slice-187-voll-abgleich-erstdurchgang-rest.md), dessen §7
  die Unschärfe als eigenes Risiko weiterführt. **Gemessen ist nur die erste
  Richtung:** Von den 18 Kandidaten der drei abgeglichenen Paare trug **einer**
  einen Befund — 94 % Rauschen. Die zweite ist **weder beobachtet noch
  ausgeschlossen**; ein Muster kann seine eigene Verengung nicht messen.
  Getragen hat deshalb nicht die Vorauswahl, sondern das Lesen: je Paar die
  vollständige Ziel-Form gegen den vollständigen Ist-Stand, die Kandidatenliste
  nur als Reihenfolge. Diese Trennung ist die Zusage, die slice-187 mit 107
  Kandidaten übernimmt.

## 8. Closure-Notiz

**Lerneintrag — Form: geschärfte Regel.** *Ein Abschnitt wird nicht groß, weil
er viel regelt, sondern weil er die Herleitung seiner Regeln nachschreibt.*
`AGENTS.md` §5 war mit 16 740 Zeichen der größte Abschnitt der Datei und ist
jetzt 10 214 groß — **39 % kleiner, ohne dass eine Regel verloren geht.** Keiner
der drei Befunde war eine überflüssige Regel: (a) drei Stellen schrieben
`modul-05`-Normtext nach, (b) zwei Mess-Regeln waren Urteilsgrundlage des
Reviewers und standen im Briefing des Implementers, (c) eine Kopieranleitung
gehörte an den Ort, an dem kopiert wird. **Die Regel, die daraus wird:** Steht
Text im falschen Dokument, ist die Antwort *umziehen*, nicht *kürzen* — und die
Gegenprobe ist, ob der Zielort die Aussage danach **vollständig** trägt.

**Und die Schärfung, die erst der Review geliefert hat** (F-1/F-2, §3.4):
*Vollständig* heißt nicht *wortgleich*. Der Umzug der zwei Mess-Regeln war
zeilengenau — genau deshalb war er falsch. Ein Text trägt zwei Dinge, die beim
Verschieben auseinanderfallen: seine **Selbstverweise** (`(§3.7)`, *„dieser
Datei"*, *„§5"* — vier Stück, alle stumm gebrochen) und seinen **Adressaten**.
Die zwei Regeln binden den *Schreibenden*, gelandet sind sie in dem Dokument,
das `harness/README.md` §Guides als *„nicht Teil der Implementer-Eingabe"*
führt. Die Antwort ist keine Rückabwicklung, sondern eine **Teilung**: die
Zusage ins Briefing, die Herleitung in den Skill. Die Prüf-Frage beim nächsten
Umzug lautet darum nicht *„steht der Text vollständig am Zielort?"*, sondern
**„liest ihn dort noch, wen er bindet — und zeigen seine Verweise noch dorthin,
wo sie hinzeigen sollen?"**

**Und die Ausnahme dazu, gemessen:** Nicht jede Doppelung ist eine. Bei (a)
blieb a-checks Hälfte stehen — die **Durchsetzung** (`make verify` ab
slice-052, der Gate-Lauf als feste Zeile, die Kopffelder ab slice-098). Die
Baseline sagt *was* gilt, a-check *dass* und *wie* es geprüft wird; wer das
zusammen streicht, verliert die Hälfte, die kein anderes Dokument trägt. Die
Prüf-Frage aus slice-183 hat genau hier getragen: nicht *„steht die Regel im
Regelwerk?"*, sondern *„schreibt dieser Absatz ihre Begründung nach?"*

**Zwei beobachtbare Closure-Kriterien.** (1) `make gates` und `make verify`
grün auf dem Stand, der nach `done/` geht — mit den drei neuen
Zellengrenzen-Regeln aus [`.d-check.yml`](../../../../.d-check.yml) und dem
`mentions`-Modul, die seit slice-184 mitlaufen. (2) §5 von `AGENTS.md` misst
10 214 Zeichen gegen 16 740 vorher, und die zwei Zielorte
([`.harness/skills/reviewer.md`](../../../../.harness/skills/reviewer.md)
§Mess-Regeln, [`docs/plan/planning/README.md`](../README.md) §Beim Kopieren der
Slice-Ziel-Form) tragen den umgezogenen Text — nachzählbar, nicht behauptet.

**Was nicht geliefert wurde, und warum es benannt statt still ist.** Die DoD
verlangte zehn Paare, geliefert sind drei. Die sieben verbleibenden tragen 107
substanzielle Kandidaten gegen 14 und 18 in den zwei Vorgängern; sie gehen an
[slice-187](../open/slice-187-voll-abgleich-erstdurchgang-rest.md). **Das ist
zum zweiten Mal derselbe Ausgang**, und darin liegt der eigentliche Befund
dieses Slice: Nicht die Slices sind zu groß, sondern der Gegenstand ist keine
Slice-Größe. Ein **Bestand** durchzugehen ist eine Kampagne mit Zähler, eine
**Regel** zu verankern ist ein Slice — die Regel selbst steht seit slice-185 in
[`harness/conventions.md`](../../../../harness/conventions.md) §Baseline und
hängt am Erstdurchgang **nicht**. Neu im Register als
[`BEO-PLAN/erstdurchgang-als-einmal-slice-geschnitten`](../observations/BEO-PLAN/erstdurchgang-als-einmal-slice-geschnitten/observation.md)
bei **2×**; erreicht slice-187 den dritten Beleg, ist die Schwelle da.

**Beobachtungs-Register — was dieser Slice zurückgibt: fünf Einträge, alle bei
2×.** Vier davon sind **neu** und tragen je zwei Belege, weil slice-185 als
Erstauftreten nachgetragen ist:
[`erstdurchgang-als-einmal-slice-geschnitten`](../observations/BEO-PLAN/erstdurchgang-als-einmal-slice-geschnitten/observation.md)
· [`vollstaendigkeits-haken-ohne-erschoepften-gegenstand`](../observations/BEO-PLAN/vollstaendigkeits-haken-ohne-erschoepften-gegenstand/observation.md)
· [`messung-ohne-reproduzierbares-instrument`](../observations/BEO-PLAN/messung-ohne-reproduzierbares-instrument/observation.md)
· [`zielsatz-nach-plan-aenderung-nicht-nachgezogen`](../observations/BEO-PLAN/zielsatz-nach-plan-aenderung-nicht-nachgezogen/observation.md)
· [`kandidaten-klassifikation-groeber-als-der-kandidat`](../observations/BEO-PLAN/kandidaten-klassifikation-groeber-als-der-kandidat/observation.md).
Der fünfte,
[`form-vergleich-sprachblind`](../observations/BEO-PLAN/form-vergleich-sprachblind/observation.md),
steigt von 1× auf 2×.

**Vier auf einmal ist kein Ausrutscher, sondern die Regel:** *Ab dem zweiten
gleichartigen Vorfall entsteht ein Eintrag* ([`AGENTS.md`](../../../../AGENTS.md)
§5). Der Review hat vier Finding-Klassen als **zweite Wiederholung in Folge**
gegenüber slice-185 ausgewiesen; jede davon war damit fällig, und keine hatte
bis dahin eine Kennung. Dass sie erst jetzt entstehen, ist selbst der Befund:
Ein Report zählt Klassen, aber niemand trug sie in den Zähler ein, solange die
Closure nur die *eigenen* Beobachtungen las. **Alle vier stehen jetzt bei 2×** —
erreicht slice-187 den dritten Beleg, sind es vier fällige Ausgänge auf einmal,
und das ist die Warnung, die dieser Absatz trägt.

**Nicht** erhöht wurde
[`BEO-GATE/muster-trifft-nur-die-haeufige-schreibweise`](../observations/BEO-GATE/muster-trifft-nur-die-haeufige-schreibweise/observation.md),
obwohl die Vorauswahl hier versagte: Sie versagte durch **Rauschen**, jene
Klasse durch **Schweigen**. Ein Beleg an der falschen Kennung hätte einen
Zähler auf 3× gehoben und eine Regel ausgelöst, die den Fall nicht trifft — die
Zählregel misst Wiederholung *einer* Klasse, nicht Häufigkeit von Ärger.
**Erhöht** wurde stattdessen
[`BEO-PLAN/form-vergleich-sprachblind`](../observations/BEO-PLAN/form-vergleich-sprachblind/observation.md)
auf 2×: Genau dort trat die Schweige-Richtung auf — der Abgleich las die
englische Fassung und übersah, dass `README.de.md` seit slice-111 eine
Ziel-Form-Sektion fehlt. Der Slice hatte die Datei angefasst und das Paar als
abgeglichen geschlossen; gefunden hat es der unabhängige Review.

**Der Review ist die zweite Hälfte dieser Closure.** Zwölf Findings, vier davon
HIGH, zwei merge-blockierend — an einem Slice, dessen eigene Gate-Läufe grün
waren und dessen Kern-Messung zeichengenau reproduziert. Das ist der Beleg für
die Regel, unter der er läuft
([`AGENTS.md`](../../../../AGENTS.md) §6, *kein Self-Review*): Die vier HIGH
liegen alle an Stellen, die der schreibende Kontext für erledigt hielt — ein
wortgleicher Umzug, ein Anker, ein *ohne Befund* geschlossenes Paar. **Vier
Finding-Klassen wiederholen sich zum zweiten Mal in Folge** gegenüber dem
slice-185-Report, alle vier an derselben Arbeitsform; sie sind in §3.4 benannt
und gehen als Prüfpunkte in slice-187 ein.

## 9. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen.** Berührt sind zwei Sub-Areas, beide in
[`harness/conventions.md`](../../../../harness/conventions.md)
§Modus-Deklaration geführt:

- **`HARNESS`** (`AGENTS.md`, `harness/`) — Achsen 1,2,3. Der Kern des Slice:
  §5 von `AGENTS.md`, `harness/README.md`, und mit (b) der Reviewer-Skill.
- **`PLAN`** (`docs/plan/planning/`) — Achsen 1,2,3. Zielort von (c), und die
  Roadmap ist eines der drei abgeglichenen Paare.

Nicht berührt: `SPEC`. Die zwei Spec-Straten stehen unter den sieben
abgetrennten Paaren — **das** ist der Grund, warum sie in slice-187 als eigene
Sub-Area geführt werden und dort eine eigene Prüfung bekommen; ein Change
Request am Lastenheft ist keine Doku-Pflege.

**Modus:** beide Greenfield, keine Begründungs-Blöcke nötig.

**Vorgelagert — offene Beobachtungen sichten.** Das Register führt zu den zwei
Sub-Areas fünf Einträge, die diesen Slice betreffen:

| Eintrag | Stand vor diesem Slice | Berührung |
|---|---|---|
| [`BEO-PLAN/erstdurchgang-als-einmal-slice-geschnitten`](../observations/BEO-PLAN/erstdurchgang-als-einmal-slice-geschnitten/observation.md) | *neu mit diesem Slice* | **angelegt, 2×** — Erstauftreten slice-185 nachgetragen, zweites Auftreten dieser Slice |
| [`BEO-PLAN/form-vergleich-sprachblind`](../observations/BEO-PLAN/form-vergleich-sprachblind/observation.md) | offen (1×) | **berührt, nicht erhöht** — `README.md` ist englisch, die Ziel-Form deutsch; der Abgleich lief hier über die *Aussage*, nicht über Titel-Gleichheit, und traf |
| [`BEO-GATE/muster-trifft-nur-die-haeufige-schreibweise`](../observations/BEO-GATE/muster-trifft-nur-die-haeufige-schreibweise/observation.md) | offen (2×) | **nicht erhöht** — die Vorauswahl versagte durch *Rauschen*, nicht durch Schweigen; das ist die andere Richtung, und ein falsch gesetzter Beleg hebt den Zähler auf 3× für eine Klasse, die hier nicht vorlag |
| [`BEO-PLAN/vier-form-vergleiche-ungeprueft`](../observations/BEO-PLAN/vier-form-vergleiche-ungeprueft/observation.md) | gestrichen | Vorläufer derselben Frage, geschlossene Menge — kein Gegenstand mehr |
| [`BEO-PLAN/ziel-form-tag-gescopt`](../observations/BEO-PLAN/ziel-form-tag-gescopt/observation.md) | offen | **nicht berührt** — betrifft die Zitier-Form, nicht den Abgleich |

**Keiner erreicht mit diesem Slice 3×.** Der neu angelegte steht bei 2× — slice-187
wäre der dritte Beleg und damit die Schwelle.

**Nachtrag nach dem Review** (§3.4): Die Sichtung oben las das Register auf
*eigene* Beobachtungen. Der Report hat vier Finding-**Klassen** als zweite
Wiederholung in Folge ausgewiesen; für sie gilt dieselbe Schwellenregel, und sie
hatten keine Kennung. Vier Einträge sind nachgetragen, dazu ein fünfter erhöht —
die Liste steht in §8. **Das ist die Lehre für die Sichtung selbst:** Der
Sichtungs-Schritt liest bisher das Register; er muss auch den **Report des
Vorgänger-Slice** lesen, sonst zählt die dritte der drei Quellen
(`modul-05` §Closure- und Lerneintrag-Regeln: *wiederkehrende Finding-Klasse aus
dem Review*) faktisch nicht mit.
