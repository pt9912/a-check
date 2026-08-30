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

- [ ] **CR 4 ist als Text geliefert** (§9): Anlass mit der Messung aus
      [slice-080](../done/slice-080-verify-abloesung-dcheck.md), Vertragsvorschlag, die
      Abgrenzung gegen `exempt-paths` und das vorweggenommene Gegenargument.
- [ ] Der **Trigger** für die spätere Ablösung steht beobachtbar in §4 — ein anderer Mensch kann
      sagen, ob er eingetreten ist, ohne Rückfrage.

- [ ] `make gates` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.
- [ ] Closure-Notiz mit benanntem Lerneintrag geschrieben (§7).
- [ ] Beobachtungs-Register fortgeschrieben.
- [ ] Jedes Risiko aus §6 trägt genau einen Ausgang.

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

1. **Ein d-check-Release trägt die Unterscheidung aus §9.** Prüfbar an `--print-config`: die
   `structure`-Sektion nennt den neuen Schlüssel.
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

- *d-check lehnt ab, weil die Nullmengen-Härte bewusst hart ist* — dann bleibt `verify-ac-form`
  dauerhaft lokal, und **das gehört in eine ADR**, nicht in stilles Liegenlassen.
  **Ausgang:** <bei Closure>
- *Der Antrag wiederholt den Fehler von CR 3 und übersieht wieder eine Konsequenz* —
  **Ausgang:** <bei Closure>

## 7. Closure-Notiz

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice,
[`AGENTS.md`](../../../../AGENTS.md) §5.)_

**Lerneintrag — Form: <geschärfte Regel | neuer Sensor | benannte Spec-Lücke>**

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

**Vertrag.** Ein optionaler Schlüssel an derselben Bedingung; ohne ihn byte-identisches Verhalten.
Kein neuer Grund-Code.

```yaml
structure:
  - files: "spec/lastenheft.md"
    section-pattern: '^### AC-'
    sections: each
    require-all: [Happy, Boundary, Negative, Out-of-Scope]
    exempt-section-pattern: '^### (AC-FA-RULE-…|…)'
    exempt-may-empty: true
    #   NUR fuer exempt-section-pattern: leert die Ausnahme die Menge, ist das
    #   kein Befund, sondern ein Bestandszustand — die Regel wartet auf den
    #   ersten nicht ausgenommenen Abschnitt. Ohne den Schluessel: section-missing,
    #   heutiges Verhalten. Greift NICHT fuer exempt-paths und NICHT, wenn schon
    #   section-pattern nichts trifft — das bleibt ein Defekt.
```

**Die Sichtbarkeit bleibt, und das ist die Bedingung.** [ADR-0075](https://github.com/pt9912/d-check/blob/main/docs/plan/adr/0075-erklaerte-teilmenge-in-structure.md) hält fest, dass die Meldung die
Zahl der ignorierten Items nennt, *„auch bei null"* — eine Zusage, die nicht wirkt, gehört
sichtbar. Dasselbe hier: mit gesetztem Schlüssel meldet die Regel **nicht nichts**, sondern eine
nicht-fatale Zeile — *„alle 19 Abschnitte sind ausgenommen; die Regel greift beim ersten
weiteren"*. Wer die Ausnahme zu breit zieht, sieht es; er sieht es nur nicht mehr als Fehler.

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
