# slice-116 — CR 4: die legitime Leermenge einer erklärten Teilmenge

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese Datei liegt — eines von
`open/`, `next/`, `in-progress/`, `done/`. Er wechselt nur durch `git mv`
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).

**Bezug:** [slice-080](../done/slice-080-verify-abloesung-dcheck.md) §6 — der dort gemessene
Paritätsbruch, an dem `verify-ac-form` als einziger Eigenbau-Sensor hängen blieb.

**Berührte Spec-Stellen:** — *(keine)*

**Verantwortlich:** Implementation; Einreichung beim Maintainer.

**Autor:** Claude (Opus 5), im Auftrag des Maintainers. **Datum:** 2026-08-30.

---

## 1. Ziel

Den Antrag formulieren, der `verify-ac-form` ablösbar macht: `structure` soll eine **erklärte,
aufzählende** Ausnahme, die gerade den ganzen Bestand deckt, von einem Konfigurationsdefekt
unterscheiden können.

## 2. Definition of Done

- [x] **CR 4 ist als Text geliefert** (§9): Anlass mit der Messung aus
      [slice-080](../done/slice-080-verify-abloesung-dcheck.md), Vertragsvorschlag, die
      Abgrenzung gegen `exempt-paths`, das vorweggenommene Gegenargument — und die **Wahl** zwischen den drei am Vertrag gangbaren Wegen samt ihrem Preis.
- [x] Der **Trigger** für die spätere Ablösung steht beobachtbar in §4 — ein anderer Mensch kann
      sagen, ob er eingetreten ist, ohne Rückfrage.

- [x] `make gates` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.
- [x] Closure-Notiz mit benanntem Lerneintrag geschrieben (§7).
- [x] Beobachtungs-Register fortgeschrieben.
- [x] Jedes Risiko aus §6 trägt genau einen Ausgang.

## 3. Plan (vor Code)

| Datei | Änderungs-Art | Begründung |
|---|---|---|
| diese Datei, §9 | neu | der CR-Text ist die Lieferung |
| [`observations.md`](../observations.md) | update | die Kollision als Beobachtung |

**Auszuführende Gates:** `make gates`, `make verify`. Kein Code, kein Modul — die Prüfung ist die
Doku-Hygiene des CR-Textes selbst.

## 4. Trigger

**Start:** eingetreten — der Bruch ist in
[slice-080](../done/slice-080-verify-abloesung-dcheck.md) gemessen, nicht vermutet.

**Trigger für die Ablösung von `verify-ac-form`** — zweiteilig, beide Hälften beobachtbar:

1. **Ein d-check-Release trägt `exempt-expect-count`** (§9, Abweichung 1). Prüfbar an
   `--print-config`: die `structure`-Sektion nennt den Schlüssel namentlich.
2. **Der Pin in [`d-check.mk`](../../../../d-check.mk) ist auf dieses Release gehoben.**

Die zweite Hälfte ist die, die man leicht wegläßt — `targets` lief so dreizehn Minor-Versionen ins
Leere ([slice-074](../done/slice-074-doc-targets-wirksam.md)).

**Alternativer Trigger, der denselben Zweck erfüllt:** ein **zwanzigstes** `AC-*` entsteht im
Lastenheft. Dann ist die geprüfte Menge nicht mehr leer, die Nullmengen-Härte kollidiert nicht
mehr, und `verify-ac-form` ist **ohne** CR 4 ablösbar. Welcher der beiden zuerst eintritt, ist
offen — darum stehen beide hier.

**Rückführungen:**

- `in-progress` → `open`: falls sich beim Schreiben zeigt, dass der Vertrag eine
  Verhaltensänderung ohne den neuen Schlüssel verlangt. Dann ist es kein CR, sondern ein
  Bruch, und der gehört anders begründet.

## 5. Closure-Trigger

CR-Text steht, Trigger ist beobachtbar formuliert, Gates grün.

**Was bewusst nicht getan wird:** **einreichen**. Fremdrepo, Maintainer-Sache — dieselbe Grenze
wie bei CR 1 und CR 2 ([slice-073 §4](../done/slice-073-dcheck-statt-eigenbau.md)) und CR 3
([slice-080 §4](../done/slice-080-verify-abloesung-dcheck.md)). Ebenso wenig wird
`verify-ac-form` angefasst: es bleibt, bis der Trigger aus §4 eintritt.

## 6. Risiken und offene Punkte

- *d-check lehnt ab, weil die Nullmengen-Härte bewusst hart ist* — **Ausgang:** entfallen,
  gestrichen mit Begründung: der Antrag ist **in der Sache angenommen** (§9). Zurückgewiesen wurde
  zweimal nur die Form — erst die vorausgesetzte nicht-fatale Meldung, dann die Erlaubnis anstelle
  einer Zahl. Die vorgesehene ADR über einen dauerhaften Verbleib von `verify-ac-form` wird damit
  nicht gebraucht.
- *Der Antrag wiederholt den Fehler von CR 3 und übersieht wieder eine Konsequenz* —
  **Ausgang:** eingetreten, und zwar **zweimal in diesem Slice**: erst die vorausgesetzte
  nicht-fatale Meldung, dann der ungemessene Preis der `--doctor`-Wahl. Beides ist der Lerneintrag
  (§7) und steht als `BEO-022` im **Beobachtungs-Register**, jetzt bei 3×.

## 7. Closure-Notiz

**Geliefert:** CR 4 als Text (§9), **in der Sache angenommen** — d-check Lastenheft `0.78.0`. Der
Bruch aus [slice-080](../done/slice-080-verify-abloesung-dcheck.md) ist behoben: die
Nullmengen-Härte unterscheidet künftig eine erklärte Teilmenge von einem Selektor-Defekt. §4 nennt
den geschärften Trigger für die Ablösung von `verify-ac-form` — `exempt-expect-count` verfügbar
und Pin gehoben —, daneben steht weiter der alternative Weg über ein zwanzigstes `AC-*`.

**Lerneintrag — Form: geschärfte Regel.** *Jede Behauptung in einem CR, die eine Messung zulässt,
wird gemessen — auch die über die eigenen Kosten.* Der Antrag hat sie **dreimal** nicht gemessen,
und jedes Mal war die Messung zwei Handgriffe entfernt:

1. **Gegen den eigenen Bestand** (CR 3): er beantragte, 19 von 19 Abschnitten auszunehmen, ohne zu
   zählen, dass danach keiner übrig bleibt. Ein `grep -c '^### AC-'` hätte es gezeigt.
2. **Gegen den fremden Vertrag** (CR 4, erster Entwurf): er machte eine *„nicht-fatale Zeile"* zur
   Bedingung.
   [`DC-FA-CLI-003`](https://github.com/pt9912/d-check/blob/main/spec/lastenheft.md#dc-fa-cli-003--exit-codes)
   führt differenzierte Exit-Codes als **Out-of-Scope**; das Lastenheft ist öffentlich.
3. **Gegen den eigenen Gate-Pfad** (CR 4, zweiter Entwurf): er nannte den Preis der
   `--doctor`-Wahl *„schwächer"*. d-check hat für sein Repo **null** gemessen, und für a-check
   gilt dasselbe — `Makefile`-Aggregate 0, `.github/workflows/` 0, `.githooks/` 0. *„Schwächer"*
   war beschönigend: es war **keine** Sichtbarkeit.

*Weil* der dritte Fall die eigenen Kosten betraf, ist die Regel weiter als „fremde Verträge
prüfen": ein CR argumentiert über zwei Systeme, und **beide** sind messbar. Die Gegenprobe zeigt,
dass es kein Aufwandsproblem ist — dieselbe Sitzung hat den Pin-Bump, die Modul-Parität und die
Backtick-Falle sauber vorab gemessen. Gemessen wurde, was *Arbeit* war; behauptet wurde, was
*Argument* war.

**Zwei beobachtbare Closure-Kriterien:**

1. Jede Zusage in §9, die eine Eigenschaft von d-check behauptet, trägt eine Kennung mit Link —
   [`DC-FA-CLI-003`](https://github.com/pt9912/d-check/blob/main/spec/lastenheft.md#dc-fa-cli-003--exit-codes)
   für die Exit-Code-Binärität,
   [`DC-FA-CLI-007`](https://github.com/pt9912/d-check/blob/main/spec/lastenheft.md#dc-fa-cli-007--diagnose-modus)
   für den Diagnose-Modus,
   [`ADR-0075`](https://github.com/pt9912/d-check/blob/main/docs/plan/adr/0075-erklaerte-teilmenge-in-structure.md)
   für die Sichtbarkeits-Zusage, gegen die abgegrenzt wird.
2. Die Antwort ist im Slice nicht zusammengefasst, sondern **je Abweichung mit ihrer Folge**
   eingetragen: Abweichung 1 schärft den Trigger in §4, Abweichung 2 ist durch eigene Messung
   bestätigt statt übernommen, Abweichung 3 korrigiert eine Zusage des Antragstexts.

**Offene Risiken und ihr Ausgang:** beide aus §6 — der erste entfallen (angenommen), der zweite
eingetreten und als `BEO-022` verbucht.

**Beobachtungs-Register:** `BEO-022` auf **3×** erhöht (Belege slice-080, slice-116 zweimal) und
verallgemeinert: nicht nur „gegen den Bestand oder den fremden Vertrag", sondern **jede messbare
Behauptung**. Damit ist die Schwelle überschritten — ab dem dritten gleichartigen Vorfall ist es
eine Harness-Lücke und verlangt einen Guide oder Sensor, nicht „besser aufpassen"
([`AGENTS.md`](../../../../AGENTS.md) §5).

**Folge-Slices:** ein **Guide für CR-Texte** — die Schwelle von `BEO-022` ist erreicht, und dieser
Slice trägt ihn nicht mehr (er hätte sonst drei Liefer-Punkte statt zwei). Dazu die Ablösung von
`verify-ac-form`, sobald einer der beiden Trigger aus §4 eintritt.
## 8. Sub-Area-Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** berührt wird die **Gate-/Werkzeug-Schicht** (der Antrag
betrifft ein Prüfmodul) und der **Planungs-Harness** (diese Datei, das Register).

**Vorgelagert — offene Beobachtungen sichten:** `BEO-007` (der Risiko-Ausgangs-Sensor prüft
innerhalb eines vorhandenen Blocks, nicht seine Existenz) liegt in derselben Schicht und trägt
dieselbe Form — eine Prüfung, die bei leerer Grundmenge nichts sagt. Sie bleibt offen; dieser
Slice ändert keinen Sensor.

Alle berührten Sub-Areas GF.

## 9. CR-Text für d-check

Dieser Abschnitt **ist** die Lieferung aus §2. Er liegt im Slice, weil §5 das Einreichen
ausdrücklich dem Maintainer überlässt — dieselbe Form wie
[slice-073 §8](../done/slice-073-dcheck-statt-eigenbau.md) und
[slice-080 §8](../done/slice-080-verify-abloesung-dcheck.md).

---

### CR 4 — `structure`: die legitime Leermenge einer erklärten Teilmenge

**Ergebnis: in der Sache angenommen, in der Form nicht** — Lastenheft `0.78.0`,
[ADR-0078](https://github.com/pt9912/d-check/blob/main/docs/plan/adr/0078-erklaerte-leermenge-mit-zahl.md).
Angenommen unverändert: der optionale Schlüssel an derselben Bedingung, byte-identisches Verhalten
ohne ihn, die Abgrenzung gegen `exempt-paths`, kein Schweregrad, Fence-Treue, und das Argument
gegen *„die Regel weglassen"*. Drei Abweichungen:

1. **`exempt-expect-count: <n>` statt `exempt-may-empty: true`** — eine **Deklaration** statt einer
   Erlaubnis. Der Grund ist die Risiko-Aussage dieses Antrags selbst: das Muster *„kann nur
   veralten"*. Genau dieser Fall ist mit einer Erlaubnis **stumm** und mit einer Zahl **laut**.
   Die Prüfung läuft beidseitig und immer, nicht nur bei leerer Restmenge.
2. **Die Sichtbarkeit bleibt im Gate-Lauf**, nicht in `--doctor`. Der Antrag nannte den Preis
   *„schwächer"*; d-check hat ihn für sein Repo gemessen und **null** gefunden — kein Gate fährt
   dort `--doctor`. **Für a-check gilt dasselbe, hier nachgemessen:** `Makefile`-Aggregate 0,
   `.github/workflows/` 0, `.githooks/` 0; `doc-doctor` ist advisory
   ([`AGENTS.md`](../../../../AGENTS.md) §4) und steht in keinem Aggregat. Der angebotene
   Re-Evaluierungs-Trigger — *„fahrt ihr `--doctor` in einem Gate, trägt eure Form"* — ist damit
   **nicht** gezogen.
3. **Es gibt doch einen neuen Grund-Code**, `section-exempt-mismatch`. Der Antrag sagte *„kein
   neuer Grund-Code"*; das galt seiner Form, die nur unterdrückt. Eine **Zahl** kann falsch sein
   und verlangt eine andere Reparatur — *Aufzählung oder Zahl nachziehen* statt *Selektor
   korrigieren* —, und die Befund-Deduplikation läuft über den Grund mit.

**Was daraus für die Ablösung folgt.** Der Trigger aus §4 ist damit präziser: nicht „ein Release
trägt den Schlüssel", sondern **`exempt-expect-count` ist verfügbar und der Pin ist gehoben**. Die
Konfiguration trägt dann `exempt-expect-count: 19` — und der zwanzigste `AC-*` meldet, wenn die
Zahl nicht mitgezogen wird, statt still durchzulaufen.

Der Antragstext bleibt darunter unverändert stehen — er ist die Lieferung aus §2 und der Beleg
dafür, was beantragt wurde.

---


**Anlass — gemessen, und der Messende ist derselbe, der CR 3 gestellt hat.** CR 3 hat geliefert,
was er beantragte: `exempt-section-pattern` erreicht Bestände **innerhalb** einer Datei, die
`exempt-paths` nicht erreicht. Bei der Anwendung stellte sich heraus, dass der Antrag seine
eigene Konsequenz nicht zu Ende gedacht hatte.

Der Adopter grandfathert **19** Anforderungen. Sein Lastenheft trägt genau **19**
`### AC-`-Überschriften. Die Regel

```yaml
- files: "spec/lastenheft.md"
  section-pattern: '^### AC-'
  sections: each
  require-all: [Happy, Boundary, Negative, Out-of-Scope]
  exempt-section-pattern: '^### (AC-FA-RULE-0(0[1-9]|1[01])|AC-FA-EXTRACT-001|…)'
```

meldet darum heute:

```
spec/lastenheft.md:1  …  section-missing
  alle 19 passenden Abschnitte sind von exempt-section-pattern ausgenommen — die Regel liefe leer
```

Das abgelöste Skript meldet an derselben Stelle `0 neue AC(s) geprueft, 19 grandfathered`, Exit 0.
**Das Modul macht mehr rot als der Sensor** — für den Adopter derselbe Bruch wie zu wenig rot, und
der Grund, warum dieser eine Eigenbau-Sensor als einziger nicht abgelöst werden konnte.

**Die Härte ist richtig, ihre Reichweite ist zu weit.** [ADR-0075](https://github.com/pt9912/d-check/blob/main/docs/plan/adr/0075-erklaerte-teilmenge-in-structure.md) begründet sie mit: *„Ohne diese
Antwort schaltete ein zu breites Muster die Regel still ab."* Das trifft ein **generisches**
Muster. Es trifft nicht ein **aufzählendes**: das Muster oben nennt 19 Kennungen einzeln und kann
nicht versehentlich zu breit werden — es kann nur veralten, und dann meldet die Regel wieder.

**Zwei Zustände, die heute denselben Grund-Code teilen und verschiedene Dinge bedeuten:**

| Zustand | Bedeutung | heute |
|---|---|---|
| `section-pattern` trifft **nichts** | Konfigurationsdefekt: falsches Muster, falsche Datei, umbenannter Abschnitt | `section-missing` ✔ richtig |
| Muster trifft, **`exempt-section-pattern` nimmt alle** | Bestandszustand: die erklärte Ausnahme deckt gerade alles — *„es gibt noch nichts Neues zu prüfen"* | `section-missing` ✘ |

**Korrektur des Antrags — der erste Entwurf war am Modell nicht prüfbar.** Er machte die
Sichtbarkeit zur Bedingung („die Regel meldet eine **nicht-fatale** Zeile"). Die gibt es nicht:
`model.Finding` trägt kein Schwere-Feld, und
[`DC-FA-CLI-003`](https://github.com/pt9912/d-check/blob/main/spec/lastenheft.md#dc-fa-cli-003--exit-codes)
ist binär — *ein Befund ist der Exit-1*, und „differenzierte Exit-Codes pro Befund-Kategorie" führt
dieselbe Anforderung ausdrücklich als **Out-of-Scope**. Der Antrag hätte damit stillschweigend
einen Schweregrad eingeführt. Die Fassung unten entscheidet sich stattdessen für einen der drei
gangbaren Wege und benennt, was sie dabei aufgibt.

**Die drei Wege, und warum es der dritte wird.**

| Weg | Was er kostet |
|---|---|
| Befund verschwindet ersatzlos | genau die Sichtbarkeit, die der Anlass verlangt — der Adopter merkt nicht, wenn seine Regel nichts mehr tut |
| Schweregrad in `model.Finding` | Vertragsänderung an `DC-FA-CLI-003`/`DC-FA-CLI-004`, viel größer als dieser Antrag, und ein **eigener** Entscheid des Werkzeugs |
| Sichtbarkeit über `--doctor` | die Zeile steht nicht mehr im Gate-Lauf, sondern im Diagnose-Modus — **schwächer**, aber ohne Vertragsbruch |

**Der dritte Weg ist tragfähig, weil `--doctor` diese Form schon führt.**
[`DC-FA-CLI-007`](https://github.com/pt9912/d-check/blob/main/spec/lastenheft.md#dc-fa-cli-007--diagnose-modus)
verlangt in seinem Boundary-Kriterium: *„Given ein Repo **ohne Befunde**, when `d-check --doctor`
läuft, then Exit-Code 0 und eine Diagnose, die ‚0 Befunde' ausweist."* Der Modus rendert also
bereits heute etwas, das kein Befund ist, und seine Diagnose erscheint *„auf stdout unabhängig vom
Code"*. Eine Zeile über deklarierte Leermengen fügt sich dort ein, ohne `model.Finding`, die
Exit-Codes oder das Zeilenformat zu berühren.

**Vertrag.** Ein optionaler Schlüssel an derselben Bedingung; ohne ihn byte-identisches Verhalten.
Kein neuer Grund-Code, kein neues Befund-Feld, keine Änderung an `DC-FA-CLI-003`/`-004`.

```yaml
structure:
  - files: "spec/lastenheft.md"
    section-pattern: '^### AC-'
    sections: each
    require-all: [Happy, Boundary, Negative, Out-of-Scope]
    exempt-section-pattern: '^### (AC-FA-RULE-…|…)'
    exempt-may-empty: true
    #   NUR fuer exempt-section-pattern: leert die Ausnahme die Menge, entsteht
    #   KEIN Befund — die Regel wartet auf den ersten nicht ausgenommenen
    #   Abschnitt. Ohne den Schluessel: section-missing, heutiges Verhalten.
    #   Greift NICHT fuer exempt-paths und NICHT, wenn schon section-pattern
    #   nichts trifft — das bleibt ein Defekt.
```

**Die Sichtbarkeit wandert, und das ist der Preis — hier benannt statt verschwiegen.** Im
Gate-Lauf ist der Zustand ab dann stumm. `--doctor` führt ihn: eine Zeile je Regel mit gesetztem
Schlüssel, die die Regel-Identität und die Zahl der ausgenommenen Abschnitte nennt — *„alle 19
Abschnitte ausgenommen; greift beim ersten weiteren"* —, in `--doctor --json` als eigenes Feld
neben `findings`, damit sie nicht als Befund fehlgelesen wird.

**Warum der Preis hier vertretbar ist.** Die Sichtbarkeits-Zusage aus
[ADR-0075](https://github.com/pt9912/d-check/blob/main/docs/plan/adr/0075-erklaerte-teilmenge-in-structure.md)
zielt auf ein Muster, das **wirkt, aber danebengreift** — dort ist die stille Fehlkonfiguration
die Gefahr. Hier ist der Fall umgekehrt: das Muster trifft **alles**, und es ist eine namentliche
Aufzählung von 19 Kennungen. Es kann nicht versehentlich zu breit werden; es kann nur veralten,
wenn der Bestand wächst und jemand die Aufzählung mitpflegt. Genau dafür ist `--doctor` der
richtige Ort — ein Modus, den man befragt, wenn man wissen will, was das Werkzeug gerade *nicht*
prüft.

**Warum nicht „die Regel einfach weglassen, bis es etwas zu prüfen gibt".** Das ist die
naheliegende Antwort und die falsche: eine Regel, die man erst einschalten muss, wenn ihr Fall
eintritt, ist kein Gate. Der Sensor existiert genau dafür, beim **ersten** neuen Element zu
greifen — und in diesem Moment denkt niemand an die Konfiguration. Der Adopter hat für dieselbe
Klasse einen eigenen Beleg: das Modul `targets` stand bereit und lief **dreizehn Minor-Versionen**
ins Leere, weil das Target eingebunden, aber nie konfiguriert wurde.

**Abgrenzung gegen `exempt-paths`.** Der Schlüssel gilt bewusst **nicht** dort. Ein Datei-Glob ist
generisch und kann versehentlich einen ganzen Baum verschlucken; eine Abschnitts-Aufzählung
innerhalb *einer* Datei kann das nicht. Wer beides braucht, hat zwei verschiedene Fragen — und die
zweite ist hier nicht beantragt.

**Fence-Treue gilt weiter** — unverändert wie in CR 3.
