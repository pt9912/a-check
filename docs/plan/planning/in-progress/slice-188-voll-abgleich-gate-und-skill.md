# slice-188 — Voll-Abgleich: Gate-Konfiguration und Closure-Skill

**Welle:** ohne Welle.

**Bezug:** Folge-Slice aus
[slice-187](../done/wellenlos/slice-187-voll-abgleich-erstdurchgang-rest.md) §1
(Schicht-Abgrenzung).
[`AC-QA-02`](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze).

**Berührte Spec-Stellen:** — · Der Slice berührt kein Spec-Stratum.

**Verantwortlich:** Claude — gesetzt beim Übergang nach `in-progress/`.

**Autor:** Claude. **Datum:** 2026-09-08.

**Lerneintrag — Form:** wird bei Closure benannt.

---

## 1. Ziel und Abgrenzung

**Ziel:** Die zwei Paare der Gate-/Skill-Schicht sind gegen ihre vendorte
Ziel-Form abgeglichen — je Paar mit einem der drei Ausgänge: übernommen ·
bewusst abweichend (mit Begründung) · ohne Befund.

| Paar | Kandidaten (Instrument: slice-187 §2) |
|---|---|
| [`.d-check.yml`](../../../../.d-check.yml) | 42 |
| [`.harness/skills/closure-note-reviewer.md`](../../../../.harness/skills/closure-note-reviewer.md) | 25 |

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Die zwei Spec-Straten.** *Es wäre ein anderer Vorgang:* Rang 1 und 2,
  Change-Request-Charakter. Eigener Folge-Slice
  ([slice-189](../open/slice-189-voll-abgleich-spec-straten.md)).
- **Die elf bereits abgeglichenen Paare.** Erledigt (slice-185/186/187).
- **Neue oder geänderte `d-check`-Module.** *Schicht-Abgrenzung:* Der Abgleich
  prüft, ob die **Konfiguration** die Ziel-Form trägt — nicht, ob a-check ein
  Modul mehr einschalten sollte. Ein neues Modul ist ein eigener Slice mit
  Pin-Frage.
- **Die elf Instanz-Vorlagen.** Kein einzelnes Gegenstück; ihre Form prüft
  `make doc-structure` über Muster.

## 2. Ausgangsmessung — neu erhoben (2026-09-08)

**Die Zahlen des Plans waren überholt, und das stand angekündigt.**
[slice-192](../done/wellenlos/slice-192-baseline-v660-vendoring.md) §3.5 hat es
festgehalten: `templates/.d-check.yml` liegt mit **+24 Zeilen** im Delta des
Baseline-Sprungs. Die hier ursprünglich genannten *42 Kandidaten* waren gegen den
**alten** Stand gemessen und sind vor der Arbeit ersetzt worden, nicht danach
erklärt.

**Instrument und Parameter:** unverändert aus
[slice-187](../done/wellenlos/slice-187-voll-abgleich-erstdurchgang-rest.md) §2 —
zitiert, nicht wiederholt; der Python-Block liegt im dortigen Archiv. Fenster 6,
Mindestlänge 6 Wörter, Korpus je Paar zwei Dateien, in `.yml` wird `#` gestrippt
statt übersprungen.

**Und ein zweiter, anders gebauter Zähler** — die dritte Mess-Regel
([`AGENTS.md`](../../../../AGENTS.md) §5, `seit slice-193`): derselbe Lauf mit
Fenster **8**. Er ist strenger (eine längere Wortfolge trifft seltener) und muss
darum **mehr** Kandidaten melden; tut er das nicht, stimmt am Zähler etwas nicht.

| Paar | Fenster 6 | Fenster 8 |
|---|---|---|
| [`.d-check.yml`](../../../../.d-check.yml) | **64** | 70 |
| [`.harness/skills/closure-note-reviewer.md`](../../../../.harness/skills/closure-note-reviewer.md) | **25** | 34 |

**Beide Zähler zeigen in dieselbe Richtung**, und der Abstand ist die Auskunft:
Bei `.d-check.yml` sind es 6 Zeilen Unterschied, beim Skill **9** — dort ist mehr
*halb* übernommen, also sinngemäß umformuliert statt wörtlich. Das ist Grenze 2
des Instruments (slice-187 §2), hier sichtbar gemacht statt nur benannt.

**Zum Vergleich:** gegen den alten Stand waren es 42 und 25. Die 22 zusätzlichen
Kandidaten bei `.d-check.yml` sind der neue `targets`-Block der Ziel-Form — den
[slice-193](../done/wellenlos/slice-193-gate-index-steht-einmal.md) inhaltlich
bereits umgesetzt hat.

## 3. Umsetzung

### 3.1 `.d-check.yml` — drei Übernahmen

**Prüf-Ebene:** Zeile für Zeile über alle **64** Kandidaten. Die Vorlage ist ein
**Startgerüst mit auskommentierten Blöcken**; der weitaus größte Teil der
Kandidaten sind Bedienhinweise der Form *„aktivieren, sobald …"*, die beim
Ausfüllen bestimmungsgemäß verschwinden. Getrennt wurde je Kandidat: **Normtext**
(Kandidat für einen Befund) gegen **Bedienhinweis** (kein Gegenstand) — der
DoD-Punkt, den dieser Slice trägt.

| # | Ziel-Form-Stelle | Was a-check hatte | Übernommen |
|---|---|---|---|
| 1 | Kopf, *„AKTIVIEREN HEISST ZWEI SCHRITTE"* | nichts — der Kopf zählte auf, was aktiv ist, nicht die Regel dahinter | Block **und** Einschalten; nur der erste Schritt lässt das Modul stumm, samt der geforderten **Gegenprobe** (einen Verstoß einbauen und den Befund sehen) |
| 2 | `exempt-targets`, *„NAMENTLICH, nie als Glob"* | die Liste ist namentlich, die Regel stand nirgends | die Regel **mit ihrem Grund**: Ein neu hinzukommendes Target soll gemeldet werden, statt still durchzulaufen |
| 3 | *(keine Ziel-Form-Stelle — Fund am Rand)* | Kopfzeile *„Baseline v3.5.2 (vendored, [`MR-006`](../../../../harness/conventions.md#mr-006--baseline-committet-vendored-statt-per-url-referenziert))"* | ersetzt durch den Zeiger auf [`harness/conventions.md`](../../../../harness/conventions.md) §Baseline |

**Befund 1 ist der teuerste, und a-check hat ihn zweimal bezahlt**, bevor die
Ziel-Form ihn aufschrieb: `targets` war konfiguriert und nicht eingeschaltet
(slice-074), `planning` ebenso
([`BEO-GATE/ruhe-marker-ungewaechtert`](../observations/BEO-GATE/ruhe-marker-ungewaechtert/observation.md)).
Beide Male meldete der Lauf **0 Befunde, Exit 0**. Die Regel steht jetzt dort, wo
der nächste Block geschrieben wird.

**Befund 3 stand in keiner Ziel-Form** — das Instrument läuft Vorlage → Repo und
konnte ihn nicht melden (Grenze 3, slice-187 §2). Gefunden hat ihn das Lesen des
Kopfes beim Suchen nach Befund 1: Die Zeile nannte einen Baseline-Stand, der
**vier Migrationen** zurückliegt. Kein Lauf fängt das — `versions` bindet an
Pfad-Pins, und dies war eine nackte Kennung. Die Antwort ist kein zweiter Pin,
sondern **kein Pin**: Der Stand steht in `conventions.md` §Baseline, und diese
Datei zeigt dorthin.

### 3.2 `.harness/skills/closure-note-reviewer.md` — ohne Befund, Abweichung bereits deklariert

**Prüf-Ebene:** alle **25** Kandidaten gelesen, dazu ein **Struktur-Vergleich auf
Abschnitts-Ebene**: Beide Dateien führen dieselben **sieben** `##`-Abschnitte in
derselben Reihenfolge, dieselben vier Schweregrade mit denselben Definitionen,
dieselbe **Negativbefund-Pflicht** und dasselbe Output-Schema.

Die 25 Kandidaten sind Umformulierungen — der größere Abstand der beiden Zähler
(25 gegen 34, §2) zeigt genau das. **Ein** Unterschied ist substanziell und
**bereits als bewusste Abweichung deklariert**, im Kopf der Datei: Die Vorlage
bindet an eine Closure-Note-ADR und an `tools/check_closure_notes.py`; beides gibt
es in a-check nicht — dieselbe ADR-Nummer trägt hier etwas anderes, und
Python-Tooling widerspräche [`AGENTS.md`](../../../../AGENTS.md) §3.1. Die Pflicht
hängt stattdessen an §5, das Struktur-Gate ist `make doc-structure`.

**Das ist der Fall, für den §1 die dritte Ausgangs-Kategorie führt:** *bewusst
abweichend, mit Begründung* — und die Begründung stand schon da, geschrieben in
slice-050. Ein Abgleich, der jede Differenz als Lücke liest, hätte hier
stillen Rückbau erzeugt.

### 3.3 Die Gegenprobe zu Befund 1 — kein stummer Block

Die übernommene Regel verlangt sie, also ist sie gefahren: **Jeder** konfigurierte
Top-Level-Block in [`.d-check.yml`](../../../../.d-check.yml) steht entweder in
`modules` oder wird per `--enable` geschaltet.

| Weg | Module |
|---|---|
| `modules:` | `links`, `anchors`, `ids`, `matrix`, `spans`, `hostpaths`, `versions`, `reviews` |
| `--enable` in `Makefile`/`d-check.mk` | `commits`, `mentions`, `planning`, `reviews`, `structure`, `targets`, `tracked`, `vcs`, `workflows` |
| **konfiguriert, aber nirgends geschaltet** | **keins** |

Die Zahlen sind gegen `grep -E "^[a-z-]+:"` und `grep -o -- "--enable [a-z]+"`
erhoben — zwei verschiedene Quellen für dieselbe Frage, wie es die dritte
Mess-Regel verlangt.

### 3.4 Vorher/Nachher, beide Zähler

| Paar | Fenster 6 | Fenster 8 |
|---|---|---|
| `.d-check.yml` | 64 → **58** | 70 → 67 |
| `.harness/skills/closure-note-reviewer.md` | 25 → **25** | 34 → 34 |

Die zweite Zeile bewegt sich nicht, und das ist richtig: Dort war nichts zu
übernehmen. Die erste sinkt um 6 bzw. 3 — die drei Übernahmen sind teils in
eigenen Worten geschrieben, und ein Wortfolgen-Zähler sieht das nicht.
**Wer die Spalte als Fortschritt liest, liest sie falsch** (slice-187 §3.4).

## 4. Definition of Done

- [ ] Beide Paare sind abgeglichen; je Paar steht der Ausgang **mit seiner
      Prüf-Ebene** im Plan.
- [ ] Für `.d-check.yml` ist getrennt, was **Bedienhinweis der Vorlage** und was
      **Normtext** ist — nur Letzteres ist ein Kandidat für einen Befund.
- [ ] Unabhängiger Review durchgeführt (Report unter [`docs/reviews/`](../../../reviews/README.md)).
- [ ] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [ ] Beobachtungs-Register fortgeschrieben.
- [ ] Jedes Risiko aus §7 trägt einen Ausgang.

`make gates` und `make verify` grün.

## 5. Trigger

**Start** (`open` → `in-progress`):
[slice-187](../done/wellenlos/slice-187-voll-abgleich-erstdurchgang-rest.md) liegt
in `done/` und das WIP-Limit ist frei.

**Rückführungen — vorab benannt:**

- `in-progress` → `next` (zu groß): 67 Kandidaten liegen über dem, was
  slice-186 getragen hat (39), aber unter slice-187 (124). Braucht `.d-check.yml`
  allein mehr als einen Durchgang, wird der Skill abgetrennt.
- `in-progress` → `open` (blockiert): Verlangt eine Ziel-Form-Regel ein Modul,
  das a-check nicht gepinnt hat, ist das ein Pin-Vorgang und keine Entscheidung
  dieses Slice.

## 6. Closure-Trigger

DoD vollständig, `make gates` und `make verify` grün, Closure-Notiz mit
Lerneintrag.

## 7. Risiken und offene Punkte

- **`.d-check.yml` ist Konfiguration, keine Prosa.** Ein „Befund" dort kann eine
  Verhaltensänderung eines Gates sein, keine Doku-Pflege — und ein geändertes
  Gate ohne ADR verletzt [`AGENTS.md`](../../../../AGENTS.md) §3.6.
  — **Ausgang:** *entfallen*, gestrichen mit Begründung: Alle drei Übernahmen
  (§3.1) sind **Dokumentations**-Änderungen — Kopfkommentar, Bedienhinweis,
  Regel-Erklärung. Kein aktives Modul wurde ein-, aus- oder umgeschaltet; die
  Gegenprobe (§3.3) bestätigt den Ist-Zustand, ändert ihn nicht. Eine
  Gate-Verhaltensänderung war unter keinem der drei Punkte nötig.
- **Der Closure-Skill steuert jede Closure.** Eine Übernahme dort wirkt sofort
  auf den nächsten Slice; ein Fehler darin fällt erst beim übernächsten auf.
  — **Ausgang:** *entfallen*, gestrichen mit Begründung: §3.2 fand **keine**
  Übernahme — die einzige inhaltliche Abweichung war bereits deklariert, seit
  slice-050. Der Skill ist unverändert; das Risiko hatte keinen Gegenstand.

## 8. Closure-Notiz

*(bei Closure auszufüllen)*

## 9. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** *(beim Übergang nach `in-progress/`
auszufüllen — berührt ist `GATE`, für den Skill zusätzlich `REVIEW`.)*

**Vorgelagert — offene Beobachtungen sichten:** *(ebenso — **zwei** Quellen:
Register und der Review-Report des Vorgänger-Slice, siehe slice-187 §9.)*
