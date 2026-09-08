# slice-187 — Voll-Abgleich: die drei Harness-Deklarations-Paare

**Welle:** ohne Welle.

**Bezug:** Folge-Slice aus
[slice-186](../done/wellenlos/slice-186-voll-abgleich-restliche-paare.md) §7,
Risiko 1 (*eingetreten*).
[`AC-QA-02`](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze).

**Berührte Spec-Stellen:** `spec/lastenheft.md` und `spec/spezifikation.md`
sind unter den Paaren — die Kennungen benennt die Umsetzung.

**Verantwortlich:** Claude — gesetzt beim Übergang nach `in-progress/`.

**Autor:** Claude. **Datum:** 2026-09-08.

**Lerneintrag — Form:** geschärfte Regel.

---

## 1. Ziel und Abgrenzung

**Plan-Änderung vor der Arbeit, nicht danach.** Der Slice wurde als *„der Rest
des Erstdurchgangs (sieben Paare)"* geschnitten. Die erste Handlung nach dem
Übergang war eine **reproduzierbare** Messung (§2) — und die sagt: 247
Kandidaten über sieben Paare, gegen 39 und ~52, die die zwei Vorgänger unter
**demselben** Instrument getragen haben. Sieben Paare sind damit kein Slice,
sondern eine Kampagne.

Der Zuschnitt wird deshalb **vor** der Arbeit korrigiert statt am Ende als
*eingetreten* geschlossen. Das ist die Antwort auf
[`BEO-PLAN/erstdurchgang-als-einmal-slice-geschnitten`](../observations/BEO-PLAN/erstdurchgang-als-einmal-slice-geschnitten/observation.md)
bei 2×: Der Eintrag sagt, ein Bestandsdurchgang gehöre **gemessen und als
Restmenge geführt**, nicht als Slice geschnitten. Genau das passiert hier.

**Ziel:** Die drei **Harness-Deklarations-Paare** sind abgeglichen — je Paar
mit einem der drei Ausgänge: übernommen · bewusst abweichend (mit Begründung) ·
ohne Befund.

| Paar | Kandidaten (§2) |
|---|---|
| [`AGENTS.md`](../../../../AGENTS.md) | 60 |
| [`harness/conventions.md`](../../../../harness/conventions.md) | 38 |
| [`docs/plan/planning/README.md`](../README.md) | 26 |

**Warum diese drei zusammen:** Sie sind dieselbe Artefaktklasse — die
Dokumente, in denen das Repo seine Regeln **deklariert** (Briefing, Konventionen,
Planning-Index). Ein Befund in einem davon ist mit hoher Wahrscheinlichkeit auch
in den anderen zu prüfen; getrennt zu schneiden hieße, dieselbe Frage dreimal zu
stellen.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Die zwei Spec-Straten** ([`spec/lastenheft.md`](../../../../spec/lastenheft.md) 28,
  [`spec/spezifikation.md`](../../../../spec/spezifikation.md) 28).
  **Es wäre ein anderer Vorgang:** Sie sind vertraglich bindend (Rang 1 und 2);
  `lastenheft.md` trägt 19 grandfatherte `AC-*`, `spezifikation.md` sieben
  `SPEC-*` mit ADRs darauf. Eine Übernahme dort ist ein **Change Request**, keine
  Doku-Pflege — andere Rolle, anderer Entscheider. Folge-Slice bei Closure.
- **`.d-check.yml` (42) und
  [`.harness/skills/closure-note-reviewer.md`](../../../../.harness/skills/closure-note-reviewer.md)
  (25).** **Schicht-Abgrenzung:** Gate-Konfiguration und Skill sind keine
  Deklarations-Dokumente; ihre Ziel-Formen prüfen andere Zusagen (Modul-Auswahl
  bzw. Urteilsgrundlage). Folge-Slice bei Closure.
- **Die acht bereits abgeglichenen Paare** (slice-185: fünf, slice-186: drei).
  Erledigt; ein zweiter Durchgang wäre Arbeit ohne Gegenstand.
- **Die elf Instanz-Vorlagen.** Dieselbe Abgrenzung wie in slice-185 §1: kein
  einzelnes Gegenstück, ihre Form prüft `make doc-structure` über Muster.
- **Der *laufende* Abgleich bei künftigen Baseline-Sprüngen.** Der ist seit
  slice-185 in [`harness/conventions.md`](../../../../harness/conventions.md)
  §Baseline verankert und braucht keinen Slice — er ist ein Schritt der
  Migration. **Dieser Slice trägt nur einen Teil des Erstdurchgangs.**
- **Ein Sensor auf Satz-Deckung.** Urteil über zwei Formulierungen, kein Match
  ([`AGENTS.md`](../../../../AGENTS.md) §3.7). Das Instrument in §2 ist eine
  **Vorauswahl**, kein Prüfer, und sagt das von sich.

## 2. Ausgangsmessung (2026-09-08) — mit dem Instrument, nicht nur mit der Zahl

**Warum das hier steht.** Die Zahlen der zwei Vorgänger stammten aus einem Lauf,
den niemand wiederholen kann — zweimal als Finding erhoben, jetzt im Register als
[`BEO-PLAN/messung-ohne-reproduzierbares-instrument`](../observations/BEO-PLAN/messung-ohne-reproduzierbares-instrument/observation.md)
bei 2×. Dessen `state.md` nennt das billige Gegenmittel: **die Parameter in den
Satz**, nicht das Werkzeug ins Repo. Dieser Slice geht einen Schritt weiter und
legt beides in den Plan, der ohnehin mit ihm archiviert wird.

**Die Parameter, vollständig:** Wortfolgen-Fenster **6**, Mindestlänge einer
geprüften Zeile **6 Wörter**. Korpus je Paar sind **genau zwei Dateien**: die
Vorlage unter `.harness/baseline/v6.5.0/templates/` und ihr Gegenstück im Repo.
Normalisiert wird auf Kleinbuchstaben ohne Markup, Links auf ihren Text
reduziert, Platzhalter `<…>` und HTML-Kommentare entfernt. Übersprungen werden
Überschriften, Blockquotes und Tabellen-Trennzeilen — **in `.md`**; in `.yml`
trägt der **Kommentar** den Text, dort wird `#` gestrippt statt übersprungen.
Eine Zeile gilt als gedeckt, sobald **eine** ihrer Wortfolgen im Gegenstück
vorkommt.

**Was das Instrument nicht kann — drei Grenzen, und die dritte ist die teuerste:**

1. Es findet **wörtliche** Übernahme. Eine sinngemäße Doppelung in anderen Worten
   sieht es nicht, und a-check formuliert fast alles um — die Zahlen sind darum
   eine **Reihenfolge, kein Befund**.
2. Eine Vorlagenzeile gilt als gedeckt, sobald **eine** ihrer Wortfolgen im
   Gegenstück vorkommt (`any(...)`). Eine nur **zur Hälfte** übernommene Zeile
   wird nie Kandidat und ist damit unsichtbar — das ist die Richtung, die zu
   **zu wenigen** Kandidaten führt, und sie ist teurer als das Rauschen.
3. Es läuft nur **Vorlage → Repo**. Eine Aussage, die a-check trägt und die
   Ziel-Form nicht kennt, kann falsch sein, ohne je Kandidat zu werden — §3.6
   führt genau so einen Fall.

Ein Kandidat ist erst ein Befund, wenn das Lesen ihn dazu macht; und das Lesen
findet Dinge, die kein Kandidat waren.

```python
#!/usr/bin/env python3
"""Ziel-Form-Abgleich: welche Zeilen der Vorlage haben im Repo-Gegenstueck
keine Entsprechung?  Parameter unten sind die GANZE Kalibrierung."""
import re, sys

FENSTER = 6          # Wortfolgen-Laenge
MIN_WOERTER = 6      # kuerzere Zeilen werden uebersprungen

def norm(t):
    t = re.sub(r'<!--.*?-->', ' ', t, flags=re.S)          # HTML-Kommentare
    t = re.sub(r'\[([^\]]*)\]\([^)]*\)', r'\1', t)          # Links -> Linktext
    t = re.sub(r'<[^>]{1,80}>', ' ', t)                     # Platzhalter <...>
    t = re.sub(r'[`*_#>|]', ' ', t)                         # Markup
    t = re.sub(r'[^0-9a-zäöüßA-ZÄÖÜ]+', ' ', t)
    return ' ' + ' '.join(t.lower().split()) + ' '

def zeilen(pfad):
    """In .md traegt die Prosa den Text, '#' ist eine Ueberschrift.
    In .yml traegt der KOMMENTAR den Text — dort wird '# ' gestrippt
    statt uebersprungen."""
    yaml = pfad.endswith(('.yml', '.yaml'))
    roh = open(pfad, encoding='utf-8').read()
    roh = re.sub(r'<!--.*?-->', '', roh, flags=re.S)
    for i, z in enumerate(roh.splitlines(), 1):
        s = z.strip()
        if yaml:
            if not s.startswith('#'):
                continue
            s = s.lstrip('#').strip()
        if not s or (not yaml and (s.startswith('#') or s.startswith('>'))):
            continue
        if re.fullmatch(r'[|\-: ]+', s):                    # Tabellen-Trennzeile
            continue
        yield i, s

def offen(vorlage, ziel):
    zieltext = norm(open(ziel, encoding='utf-8').read())
    raus = []
    for nr, z in zeilen(vorlage):
        w = norm(z).split()
        if len(w) < MIN_WOERTER:
            continue
        treffer = any(' ' + ' '.join(w[i:i+FENSTER]) + ' ' in zieltext
                      for i in range(len(w) - FENSTER + 1))
        if not treffer:
            raus.append((nr, z))
    return raus

if __name__ == '__main__':
    for paar in sys.argv[1:]:
        v, z = paar.split('::')
        r = offen(v, z)
        print(f'=== {z}  ({len(r)} Kandidaten gegen {v})')
        for nr, s in r:
            print(f'  {nr:4d}  {s[:150]}')
```

Aufruf: `python3 abgleich.py "<vorlage>::<repo-datei>" …`

**Ergebnis über die 16 gemessenen Paar-Kombinationen** — die bereits
abgeglichenen sind mitgemessen, weil eine Zahl ohne Vergleichsmaßstab keine
Größe ist. `README.md` und `README.de.md` stehen **beide** darin: Die Ziel-Form
führt eine Zeile, a-check zwei Dateien (`BEO-PLAN/form-vergleich-sprachblind`):

| Paar | Kandidaten | Stand |
|---|---|---|
| `AGENTS.md` | **60** | dieser Slice |
| `.d-check.yml` | **42** | abgetrennt (§1) |
| `harness/conventions.md` | **38** | dieser Slice |
| `spec/lastenheft.md` | **28** | abgetrennt (§1) |
| `spec/spezifikation.md` | **28** | abgetrennt (§1) |
| `docs/plan/planning/README.md` | **26** | dieser Slice |
| `.harness/skills/closure-note-reviewer.md` | **25** | abgetrennt (§1) |
| `.harness/skills/reviewer.md` | 40 | slice-185 |
| `spec/architecture.md` | 20 | slice-185 — **eine Klausel blieb offen**, siehe §3.6 |
| `README.md` | 10 | slice-186 |
| `harness/README.md` | 23 | slice-186 |
| `docs/plan/planning/in-progress/roadmap.md` | 9 | slice-186 |
| `README.de.md` | 7 | slice-186 |
| `Makefile` | 4 | slice-185 |
| `docs/plan/adr/README.md` | 4 | slice-185 |
| `docs/plan/carveouts/README.md` | 4 | slice-185 |

**Die Kalibrierung, die den Zuschnitt trägt:** slice-186 hat unter diesem
Instrument **39** Kandidaten getragen (23 + 9 + 7), slice-185 rund **52**. Dieser
Slice trägt **124** — mehr als beide, und das ist eine bewusste Setzung, keine
Nachlässigkeit: Die drei Paare sind dieselbe Artefaktklasse (§1), und `AGENTS.md`
ist mit slice-186 zur Hälfte schon gelesen. Die verbleibenden vier Paare stehen
mit **123** daneben; sie sind die **Restmenge mit Zähler**, nicht ein offener
Rest.

**Der Vergleich mit slice-186 §3.3 hält nicht** — dort standen 107 Kandidaten für
sieben Paare, hier 247 für dieselben sieben. Zwei Instrumente, zwei Zahlen; nur
diese hier ist nachrechenbar. Genau das ist der Grund für den Abschnitt.

## 3. Umsetzung

**Acht Übernahmen über drei Paare.** Je Befund steht unten, *woran* er hängt und
*auf welcher Ebene* er geprüft wurde — die Ebene gehört in den Ausgang
([`BEO-PLAN/vollstaendigkeits-haken-ohne-erschoepften-gegenstand`](../observations/BEO-PLAN/vollstaendigkeits-haken-ohne-erschoepften-gegenstand/observation.md)).

### 3.1 `AGENTS.md` — fünf Übernahmen

**Prüf-Ebene:** Zeile für Zeile über alle **60** Kandidaten, jeder einer
Ziel-Form-Sektion zugeordnet (10 Sektionen), und je Kandidat die Frage *„trägt
a-check die Aussage — irgendwo?"* gegen den Ist-Stand. **Nicht** geprüft: was
die Ziel-Form **nicht** führt (§2, Grenze 3) und Zeilen, deren erste sechs
Wörter treffen (Grenze 2).

| # | Ziel-Form-Stelle | Was a-check hatte | Übernommen |
|---|---|---|---|
| 1 | §4, Satz nach der Tabelle | nichts | *„Diese Tabelle **listet auf**; definiert wird hier nichts"* plus den Weg zur Bindung — `harness/README.md` §Sensors, von dort zur `AC-*`-ID |
| 2 | §3.3, zwei nummerierte Fälle | **nur den Regelfall**, als Reihenfolge statt als Wahl | beide Fälle; der `done/`-Übergang kehrt die Reihenfolge um |
| 3 | §1, *Breiterer Pflicht-Blick* | nur die Auswahlregel *„ein Abschnitt je Aufgabe"* | die drei Anlässe, bei denen das **nicht** reicht |
| 4 | §1, zwei Rollen der Vorlagen | *„die Ziel-Formen daneben unter `templates/`"* | Referenz-Form **und** Kopiervorlage, für sechs Artefaktklassen |
| 5 | §5, Struktur-IDs | nichts | `SPEC-<NNN>` gehört nicht in die Commit-Message |

**Befund 2 ist der schwerste, und er ist gemessen statt vermutet.** Die
Ziel-Form sagt seit jeher, dass die *Reihenfolge vom Vorgang abhängt*: Regelfall
`git mv` zuerst, Lifecycle-Übergang nach `done/` **umgekehrt**, weil die
Closure-Notiz die Bedingung für `done/` ist und nicht ihre Folge. a-checks §3.3
trug nur den Regelfall — **und das Repo fährt den zweiten**: Für slice-183 bis
slice-186 liegt der Closure-Commit ausnahmslos **vor** dem
`… -> done (make slice-mv)`-Commit (gemessen; `make slice-mv` gibt es erst seit
slice-118, für ältere Slices ist der Geltungsbereich nicht erhoben). Das Briefing sagte
also das Gegenteil der geübten Praxis. Nicht die Praxis war falsch, sondern der
Satz, an dem sie gemessen worden wäre.

**Befund 1 ist der älteste.** Der Satz steht seit `v5.12.0` unverändert in der
Ziel-Form, überlebte **vier** Baseline-Deltas und ist genau das Beispiel, mit
dem slice-185 den Voll-Abgleich begründet hat
([`harness/conventions.md`](../../../../harness/conventions.md) §Baseline). Er
war damit **benannt, aber nicht übernommen** — zwei Slices lang.

### 3.2 `harness/conventions.md` — zwei Übernahmen, eine Grenze, ein `MR`

**Prüf-Ebene:** dieselbe — Zeile für Zeile über **38** Kandidaten, fünf
Ziel-Form-Sektionen.

| # | Ziel-Form-Stelle | Übernommen |
|---|---|---|
| 6 | §Adaptions-Block, *Regeln dieser Sektion* | Die Datei trägt den **Index**, nicht die Einträge; der Zustand ist die Verzeichnis-Position, kein Status-Feld; und der Grund — was hier steht, liest **jeder** Lauf |
| 7 | §[`MR-000`](../../../../harness/conventions.md#mr-000), *Bleibt hier* | warum der Eintrag keine eigene Datei bekommt: Adoptions-Erklärung, keine Adaption |

Auch hier war die **Praxis vorhanden und die Regel unausgesprochen**: a-check
führt **seit slice-096** eine Datei je Eintrag mit `conventions/done/` als
zweitem Ort — davor lebten die Einträge als `###`-Abschnitte in
`harness/conventions.md`, und die Datei sagt das selbst. Was fehlte, war der
Satz, der die heutige Ablage zur Regel macht — und mit ihm die Begründung, die
den Schnitt trägt.

**Zwei Ziel-Form-Punkte werden bewusst *nicht* übernommen, beide mit derselben
Ursache:** Das Pflichtfeld *Ersetzt-Baseline-Regel* fehlt in
[`MR-000`](../../../../harness/conventions.md#mr-000), und die
**Beobachtungs-Kennung** steht nicht in dessen ID-Liste. Beides ließe sich nur
durch eine **inhaltliche Änderung an einem akzeptierten Eintrag** beheben, und
die verbietet §Disziplin (analog [`AGENTS.md`](../../../../AGENTS.md) §3.5). Die
Abweichung steht jetzt **im Eintrag selbst als benannte Grenze**, statt als
Leerstelle dazustehen; die Beobachtungs-Kennung ist in §Modus-Deklaration
deklariert, wo auch ihr Kürzel herkommt.

### 3.3 `docs/plan/planning/README.md` — eine Übernahme, ein nachgereichter Befund

**Prüf-Ebene:** Zeile für Zeile über **26** Kandidaten. **Diese Ebene reichte
nicht** — der Review fand in derselben Datei eine gegen den Repo-Bestand
widerlegte Aussage, die **kein** Kandidat war, weil die Ziel-Form den Abschnitt
nicht führt (§3.6, F-5).

| # | Ziel-Form-Stelle | Übernommen |
|---|---|---|
| 8 | Kopf, *Reine `git mv`-Commits beim Wechsel* | der Zeiger auf die Hard Rule — **samt** der umgekehrten Reihenfolge nach `done/` (Befund 2) |

Der Rest des Paares ist **in der Richtung Vorlage → Repo ohne Befund**, und an zwei Stellen ist
a-checks Fassung die schärfere: §Aktueller Stand sagt nicht nur *„nicht als
Snapshot eintragen"*, sondern **warum** (der Snapshot driftet gegen die
Verzeichnisse, und niemand merkt es), und das Beobachtungs-Register ist mit
Verzeichnisform, abgeleitetem Zähler und der Wellen-Unabhängigkeit beschrieben,
wo die Ziel-Form nur die Existenz nennt. Das `reconciliation.md` entfällt
begründet — a-check kam nicht aus einem Brownfield-Bootstrap.

### 3.4 Vorher/Nachher, mit demselben Instrument

| Paar | vorher | nachher |
|---|---|---|
| `AGENTS.md` | 60 | **48** |
| `harness/conventions.md` | 38 | **29** |
| `docs/plan/planning/README.md` | 26 | **26** |

**Die dritte Zeile ist der Befund an der Messung selbst.** Übernahme 8 steht im
Dokument, und die Zahl bewegt sich nicht: Das Instrument misst **wörtliche**
Überlappung, und a-check hat den Punkt in eigenen Worten übernommen. Eine
sinkende Zahl belegt eine Übernahme; eine gleichbleibende widerlegt keine. **Wer
diese Spalte als Fortschritt liest, liest sie falsch** — sie ist eine
Reihenfolge, wie §2 sagt, und bleibt es auch hinterher.

### 3.5 Ein Befund neben den Paaren: eine Zusage, die weiter reicht als ihr Prüfer

Der Abgleich stellte die Frage *„was passiert, wenn ich
[`MR-000`](../../../../harness/conventions.md#mr-000) trotzdem ändere?"* — und **maß** statt anzunehmen: `make doc-immutable` über die
Commit-Range bleibt **grün**. Der Grund steht in
[`.d-check.yml`](../../../../.d-check.yml): Das Modul `vcs` führt
`paths: ["docs/plan/adr/[0-9]*.md"]`; `harness/conventions.md` steht dort nicht.

**§Adaptions-Block sagte trotzdem *„analog zur ADR-Immutabilität"* mit Verweis
auf [`AGENTS.md`](../../../../AGENTS.md) §3.5** — und dieser Verweis liest sich
als Verweis auf dieselbe **Durchsetzung**. Für `MR`-Einträge gibt es sie nicht.

Behoben ist es an der Zusage, nicht in der Konfiguration: §Disziplin trägt jetzt
die Grenze im Satz (*analog* meint die Regel, nicht ihren Lauf) samt dem, was ein
künftiger Sensor dort tragen müsste — ein `MR`-Eintrag hat kein `Status:`-Feld,
an dem `immutable-when` greifen könnte. **Neu im Register:**
[`BEO-GATE/zusage-weiter-als-ihre-durchsetzung`](../observations/BEO-GATE/zusage-weiter-als-ihre-durchsetzung/observation.md)
mit **zwei** Belegen — slice-186 (`doc-planning` sagte *„benennt ihn"* zu und
prüft eine Äquivalenz) und dieser Slice.

**Und eine Folge für den Abgleich selbst:** Der Kommentar zum
[Baseline-Eintrag](../../../../harness/conventions.md#mr-000) steht
**über** dem Eintrag, nicht in ihm. Das ist keine Formalie — hätte er drinnen
gestanden, wäre die Übernahme einer Ziel-Form-Regel selbst ein Verstoß gegen die
Regel gewesen, die der Eintrag trägt, und **kein Lauf hätte es gemeldet**.


### 3.6 Was der unabhängige Review verändert hat

Der Report
([`2026-09-08-slice-187-…`](../../../reviews/2026-09-08-slice-187-voll-abgleich-erstdurchgang-rest.md))
trug **5 HIGH · 6 MEDIUM · 3 LOW · 1 INFO**. Er hat die Ausgangsmessung
**vollständig reproduziert** — alle vierzehn Zahlen, mit dem Instrument aus §2
gegen `HEAD` und gegen den Stand davor. Und er hat fünf Dinge gefunden, die der
Abgleich nicht finden konnte oder übersehen hat:

| Befund | Kern | Behebung |
|---|---|---|
| **F-1** Ein liegengebliebener Punkt aus slice-185 | `spec/architecture.md` trägt eine **verletzte** dritte Klausel (Historie-Abschnitt); slice-185 adressierte sie an slice-186, dort kommt sie nicht vor, hier auch nicht — und slice-189 schloss das Paar mit *„bereits abgeglichen"* aus | slice-189 **nimmt** die Klausel: §1 führt sie, die DoD verlangt eine Entscheidung (streichen oder als `MR` deklarieren), §7 trägt das Anker-Risiko |
| **F-2** Der Sichtungs-Schritt las die falsche Sub-Area | §9 nennt `HARNESS` als berührt und sichtet **sechs** `BEO-PLAN`-Einträge; die 15 `BEO-HARNESS`-Einträge blieben ungelesen | §9 sichtet jetzt nach **allen** berührten Kürzeln; zwei `BEO-HARNESS`-Einträge erreichen dadurch 3× |
| **F-3** Chronik in `AGENTS.md` §3.3 | Der Abgleich erklärte die Korrektur **im Regeltext** — genau die Form, die §3.7 verbietet | beide Stellen der Datei entfernt; auf Maintainer-Weisung dazu die Forensik in §1/§4/§5 |
| **F-4** Die Werkzeug-Wahl trug nicht | *„ließe sich **nur** durch eine inhaltliche Änderung beheben"* — §Disziplin nennt zwei Instrumente, und [`MR-020`](../../../../harness/conventions.md#mr-020) ist der Präzedenzfall | [`MR-023`](../../../../harness/conventions.md#mr-023) statt freistehendem Absatz |
| **F-5** *„ohne Befund"* für eine widerlegte Aussage | `docs/plan/planning/README.md` behauptete, a-check habe Wellen-Closure-Schritt 4 nicht adoptiert — 13 `done/welle-*/archiv.zip` und ein Pflichtschritt in [`AGENTS.md`](../../../../AGENTS.md) §6 sagen das Gegenteil | Der Abschnitt führt jetzt **sechs** Schritte, Schritt 4 mit Beleg |

**F-5 ist der Befund über den Abgleich selbst.** Die falsche Aussage war **kein
Kandidat** und konnte keiner sein: Das Instrument läuft nur Vorlage → Repo, und
die Ziel-Form führt den Abschnitt gar nicht. Der Ausgang *„ohne Befund auf
Satz-Ebene"* behauptete trotzdem die Satz-Ebene für das **Paar**. Die Lehre steht
jetzt als dritte Grenze in §2 und als korrigierter Ausgang in §3.3: **Ein
Ziel-Form-Abgleich prüft, ob das Repo die Vorlage trägt — nicht, ob das Repo
recht hat.**

**Die übrigen zehn** in Kürze: die Probe zu `doc-immutable` hatte keine Mutation
vor sich, der tragende Beleg (die `paths`-Liste) steht daneben (F-7) · die
Prüf-Ebene stand nur in einem der drei Ausgänge (F-8, jetzt in allen dreien) ·
zwei Provenienz-Aussagen waren aus dem Gedächtnis, *„seit jeher"* stimmt für die
Verzeichnis-Form erst seit slice-096 (F-9) · der Geltungsbereich des Instruments
nannte nur die billige Richtung (F-10, jetzt drei Grenzen) · die benannte
Reibung am Risiko-Ausgang hatte keine Kennung und wäre mit dem Slice archiviert
worden (F-11 → [`BEO-PLAN/risiko-eingetreten-und-im-slice-aufgeloest`](../observations/BEO-PLAN/risiko-eingetreten-und-im-slice-aufgeloest/observation.md))
· drei Sammelaussagen waren gröber als ihre Menge (F-12) · zweimal war
`make slice-mv` eine Leistung zugeschrieben, die es nicht erbringt (F-13) · eine
Doppelung in derselben Datei (F-14) · F-15 ist INFO.

**Zwei Register-Einträge erreichen durch diesen Slice 3×**, beide unter
`BEO-HARNESS` — also genau dort, wo §9 nicht hingesehen hatte:
[`chronik-in-gelesenen-dateien`](../observations/BEO-HARNESS/chronik-in-gelesenen-dateien/observation.md)
→ *geplant*, [slice-191](../open/slice-191-chronik-phrasen-sensor.md)
(Phrasen-Sensor auf die häufige Schreibweise, mit der Grenze am Sensor) ·
[`adaption-korrigiert-repo-aussage`](../observations/BEO-HARNESS/adaption-korrigiert-repo-aussage/observation.md)
→ *geplant*, [slice-190](../open/slice-190-id-schema-deklaration-ueberarbeiten.md)
(die ID-Schema-Deklaration wird überarbeitet statt weiter geflickt — ihr eigener,
im Eintrag genannter Trigger).


## 4. Definition of Done

- [x] Die **drei** Harness-Deklarations-Paare sind abgeglichen; je Paar steht
      der Ausgang im Plan — übernommen · bewusst abweichend (mit Begründung) ·
      ohne Befund.
- [x] Jede Übernahme ist als solche kenntlich, und **die Ebene steht im
      Ausgang**: auf welcher Auflösung geprüft wurde, nicht nur *ob*
      ([`BEO-PLAN/vollstaendigkeits-haken-ohne-erschoepften-gegenstand`](../observations/BEO-PLAN/vollstaendigkeits-haken-ohne-erschoepften-gegenstand/observation.md),
      2×).
- [x] Die **Restmenge** ist mit Zähler übergeben: vier Paare, 123 Kandidaten,
      zwei Folge-Slices mit Kennung —
      [slice-188](../open/slice-188-voll-abgleich-gate-und-skill.md) (Gate und
      Skill, 67) und
      [slice-189](../open/slice-189-voll-abgleich-spec-straten.md) (die zwei
      Spec-Straten, 56, **plus** die seit slice-185 offene dritte
      Architektur-Klausel, Review F-1). Beide zitieren das Instrument aus §2,
      statt es zu wiederholen.
- [x] Unabhängiger Review durchgeführt (Report unter [`docs/reviews/`](../../../reviews/README.md)) — 15 Findings, alle abgearbeitet (§3.6).
- [x] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [x] Beobachtungs-Register fortgeschrieben.
- [x] Jedes Risiko aus §7 trägt einen Ausgang.

`make gates` und `make verify` grün. Ein öffentlicher Vertrag ist berührt:
`AGENTS.md` ist Rang 8, `harness/conventions.md` Rang 9.

## 5. Trigger

**Start** (`open` → `in-progress`):
[slice-186](../done/wellenlos/slice-186-voll-abgleich-restliche-paare.md) liegt in
`done/` und das WIP-Limit ist frei. **Eingetreten am 2026-09-08.**

**Rückführungen — vorab benannt:**

- `in-progress` → `next` (zu groß): Der Zuschnitt ist **vor** der Arbeit auf drei
  Paare korrigiert worden (§1) — die Rückführung wurde damit vorweggenommen,
  statt sie am Ende als Risiko-Ausgang zu buchen. Sie greift trotzdem, wenn eines
  der drei Paare für sich mehr als einen Durchgang braucht; abgetrennt wird dann
  `AGENTS.md`, das größte.
- `in-progress` → `open` (blockiert): Findet der Abgleich eine Ziel-Form-Regel,
  die a-check bewusst **nicht** übernehmen will, ist das eine `MR`-Adaption und
  keine Nachtrags-Entscheidung dieses Slice.

## 6. Closure-Trigger

DoD vollständig, `make gates` und `make verify` grün, Closure-Notiz mit
Lerneintrag.

## 7. Risiken und offene Punkte

- **Drei Paare mit 124 Kandidaten sind mehr, als beide Vorgänger getragen
  haben** (39 und ~52 unter demselben Instrument). Die Setzung ist begründet
  (§2), aber sie ist eine Wette darauf, dass `AGENTS.md` durch slice-186 zur
  Hälfte gelesen ist. — **Ausgang:** *entfallen*, gestrichen mit Begründung: Die
  Wette hielt. Alle drei Paare sind in **einem** Durchgang abgeglichen, keine
  Rückführung, kein Paar abgegeben; 124 Kandidaten ergaben **acht** Übernahmen
  (6,5 %). Getragen hat nicht die Zahl, sondern dass die Kandidaten dicht
  beieinander lagen — dieselbe Artefaktklasse, wie §1 gesetzt hatte.
- **Vier Register-Einträge stehen bei 2× und treffen genau diese Arbeitsform.**
  Erreicht einer mit diesem Slice 3×, ist er keine Notiz mehr, sondern eine
  Lücke mit fälligem Ausgang. — **Ausgang:** *entfallen*, gestrichen mit
  Begründung: **Keiner** der vier erreicht 3×, und zwar nicht durch Glück — der
  Plan hat ihre Lehre **vorab angewandt** (§9): Zuschnitt vor der Arbeit
  korrigiert, Instrument in den Plan gelegt, Ebene in den Ausgang, Zielsatz
  sofort nachgezogen. Ein Eintrag, dessen Lehre wirkt, zählt nicht weiter.
  **Das Risiko war aber falsch gestellt** — es fragte nach *diesen vier*, und
  §9 sah nur in `BEO-PLAN`. Drei andere sind dennoch bewegt worden:
  [`zusage-weiter-als-ihre-durchsetzung`](../observations/BEO-GATE/zusage-weiter-als-ihre-durchsetzung/observation.md)
  neu bei 2× (§3.5), und **zwei erreichen 3×** —
  [`chronik-in-gelesenen-dateien`](../observations/BEO-HARNESS/chronik-in-gelesenen-dateien/observation.md)
  und [`adaption-korrigiert-repo-aussage`](../observations/BEO-HARNESS/adaption-korrigiert-repo-aussage/observation.md),
  beide unter `BEO-HARNESS`, beide mit Ausgang *geplant* (§3.6). Der Zähler hat
  gearbeitet; die Sichtung hatte nur nicht hingesehen.
- **Die Vorauswahl bleibt eine Reihenfolge, kein Befund.** — **Ausgang:**
  *entfallen*, gestrichen mit Begründung: Die Zahl trägt ihre Warnung jetzt
  selbst. §2 nennt Parameter und Instrument, §3.4 misst das Verhältnis (124
  Kandidaten → 8 Übernahmen) **und** die Gegenrichtung: Das Paar
  `docs/plan/planning/README.md` steht vorher wie nachher bei 26, obwohl eine
  Übernahme dort steht — in eigenen Worten geschrieben, und das sieht ein
  Wortfolgen-Test nicht. Wer die Spalte als Fortschritt liest, liest sie falsch;
  das steht jetzt daneben.
- **`harness/conventions.md` ist Rang 9 und trägt die Adaptions-Einträge.** Eine
  Ziel-Form-Regel, die dort etwas ändert, kann einen akzeptierten `MR` berühren —
  und die sind immutabel. — **Ausgang:** *entfallen*, gestrichen mit Begründung:
  Es **trat ein** und wurde **im Slice aufgelöst**, ohne den Eintrag anzufassen.
  Ein Ziel-Form-Punkt ist über
  [`MR-023`](../../../../harness/conventions.md#mr-023) übernommen — dem
  Instrument, das §Disziplin dafür nennt —, der andere steht als benannte Grenze
  **über** [`MR-000`](../../../../harness/conventions.md#mr-000) (§3.2, §3.5).
  Der Eintrag selbst ist byte-identisch geblieben.
  **Die Reibung hat jetzt eine Kennung:** Die geschlossene Dreier-Menge hat
  keine Kategorie für *„eingetreten und im Slice selbst aufgelöst"* —
  *eingetreten* verlangt einen **künftigen** Adressaten, den es hier nicht
  braucht, und *entfallen* behauptet, das Risiko sei nie eingetreten. Als
  [`BEO-PLAN/risiko-eingetreten-und-im-slice-aufgeloest`](../observations/BEO-PLAN/risiko-eingetreten-und-im-slice-aufgeloest/observation.md)
  bei 1× angelegt, ausdrücklich **getrennt** von
  [`risiko-ausgang-fuer-gewollte-wirkung`](../observations/BEO-PLAN/risiko-ausgang-fuer-gewollte-wirkung/observation.md)
  (dort war der Punkt nie ein Risiko) — ob beide dieselbe Klasse sind, entscheidet
  der Lese-Schritt bei 3× und nicht die Anlage. **Nur benennen hätte nicht
  gereicht:** Der Plan wird archiviert, und was keine Kennung hat, ist beim
  nächsten Vorkommen nicht auffindbar (Review F-11).

## 8. Closure-Notiz

**Lerneintrag — Form: geschärfte Regel.** *Wo ein Repo eine Regel **übt**, ohne
sie zu **schreiben**, sagt sein Briefing im Zweifel das Gegenteil.* Sechs der
acht Übernahmen sind von dieser Art, und zwei davon sind die schwersten:
`AGENTS.md` §3.3 trug nur den Regelfall *„erst `git mv`, dann Inhalt"* — während
jede Slice-Closure dieses Repos die **umgekehrte** Reihenfolge fährt, weil die
Closure-Notiz die Bedingung für `done/` ist. Und der Adaptions-Block beschrieb
seine Disziplin, ohne zu sagen, dass die Datei der **Index** ist und der Zustand
die Verzeichnis-Position. In beiden Fällen war die Praxis richtig und der Satz
falsch oder abwesend. **Der Voll-Abgleich findet genau diese Klasse**, und kein
Sensor kann sie finden: Ein Lauf misst, was dasteht, gegen das, was dasteht —
nicht gegen das, was getan wird.

**Warum das nicht der Delta-Analyse auffällt** — die Frage, die slice-185
aufgeworfen hat, ist hier zum zweiten Mal beantwortet: Der Arbeitsteilungs-Satz
für §4 steht seit `v5.12.0` unverändert in der Ziel-Form, hat **vier**
Baseline-Deltas überlebt und war seit slice-185 **benannt**. Zwei Slices lang
stand er auf einer Liste und nicht im Dokument. Ein Delta findet Änderungen; was
seit der Adoption fehlt, findet nur der Voll-Abgleich — und übernommen ist es
erst, wenn es dasteht.

**Zwei beobachtbare Closure-Kriterien.** (1) `make gates` und `make verify`
grün auf dem Stand, der nach `done/` geht. (2) Die Ausgangsmessung ist
**wiederholbar**: §2 trägt Parameter und Instrument, und derselbe Aufruf liefert
nach der Arbeit 48 / 29 / 26 gegen vorher 60 / 38 / 26. Das ist der erste
Slice dieser Kette, dessen Zahlen jemand nachrechnen kann.

**Und die dritte Zahl ist selbst der Lerneintrag zur Messung.** 26 vor der
Arbeit, 26 danach — mit einer Übernahme dazwischen. Das Instrument misst
**wörtliche** Überlappung; a-check hat den Punkt in eigenen Worten übernommen,
und die Zahl bewegt sich nicht. **Eine sinkende Zahl belegt eine Übernahme;
eine gleichbleibende widerlegt keine.** Wer die Spalte als Fortschritt liest,
hat die Vorauswahl zum Prüfer erklärt.

**Was der Slice nicht getragen hat, und mit welcher Adresse.** Vier Paare
bleiben: `.d-check.yml` (42), `spec/lastenheft.md` (28),
`spec/spezifikation.md` (28), `.harness/skills/closure-note-reviewer.md` (25) —
**123 Kandidaten**, gemessen mit demselben Instrument. Sie sind **keine offene
Aufgabe, sondern eine Restmenge mit Zähler**: §1 nennt für jede den
Ausschlussgrund, §2 die Zahl, und zwei Folge-Slices tragen sie mit Kennung —
[slice-188](../open/slice-188-voll-abgleich-gate-und-skill.md) und
[slice-189](../open/slice-189-voll-abgleich-spec-straten.md), in dieser
Reihenfolge, weil der Spec-Abgleich der heikelste ist und zuletzt kommt. Die zwei Spec-Straten sind dabei ausdrücklich
**ein anderer Vorgang** — dort ist eine Übernahme ein Change Request, kein
Doku-Pflege-Schritt.

**Beobachtungs-Register.** Neu:
[`BEO-GATE/zusage-weiter-als-ihre-durchsetzung`](../observations/BEO-GATE/zusage-weiter-als-ihre-durchsetzung/observation.md)
bei **2×** (§3.5) — eine Zusage, die weiter reicht als der Lauf, der sie
einlösen soll; slice-186 und dieser Slice. **Nicht** erhöht wurden die vier
Einträge aus §9: Ihre Lehre ist in diesen Plan eingebaut, statt ein drittes Mal
belegt zu werden — was der Zweck eines Registers ist und nicht seine Umgehung.

**Der Sichtungs-Schritt hat zum ersten Mal zwei Quellen gelesen** (§9): das
Register **und** den Review-Report des Vorgängers. Die zwei schärfsten Warnungen
für diesen Slice standen nur im Report — *Selbstverweis beim Umzug nicht
re-verankert* und *Zusage in der Doku, Grenze nur im Konfigurations-Kommentar* —,
und die zweite hat unmittelbar getragen: Sie ist §3.5 geworden. `modul-05` nennt
die Finding-Klasse als dritte Zähler-Quelle; ohne diesen Griff zählt sie nicht
mit.

**Und die erste Quelle war zu schmal — das ist der zweite Lerneintrag.** §9 nennt
`HARNESS` als berührte Sub-Area und sichtete dann **nur** `BEO-PLAN`. Die 15
`BEO-HARNESS`-Einträge blieben ungelesen, und **zwei davon treten in genau
diesem Diff ein drittes Mal auf**: Chronik im Briefing und ein Adaptions-Eintrag,
der eine Repo-Aussage korrigiert. Beide haben jetzt ihren Ausgang
([slice-191](../open/slice-191-chronik-phrasen-sensor.md),
[slice-190](../open/slice-190-id-schema-deklaration-ueberarbeiten.md)), aber
gefunden hat sie der Review, nicht die Sichtung. **Die Regel, die daraus wird:**
Das Register wird nach **allen** berührten Kürzeln gelesen — nicht nach dem, in
dessen Verzeichnis man gerade arbeitet. Sie steht in §9 der drei Folge-Slices.

**Der Review hat außerdem den Befund geliefert, den der Abgleich nicht finden
konnte** (§3.6, F-5): `docs/plan/planning/README.md` behauptete, a-check habe
den Archivierungs-Schritt der Wellen-Closure nicht adoptiert — 13 Archive und ein
Pflichtschritt im Briefing sagen das Gegenteil. Die Stelle war **kein Kandidat**
und konnte keiner sein, weil das Instrument nur Vorlage → Repo läuft. **Ein
Ziel-Form-Abgleich prüft, ob das Repo die Vorlage trägt — nicht, ob das Repo
recht hat.** Diese Grenze steht jetzt in §2 und im Ausgang des Paares.

**Und ein Punkt war zwei Slices lang unterwegs verloren** (F-1): slice-185 hatte
für `spec/architecture.md` eine von drei Klauseln als **verletzt** offen gelassen
und an slice-186 adressiert; dort und hier kommt sie nicht vor, und
[slice-189](../open/slice-189-voll-abgleich-spec-straten.md) schloss das Paar mit
*„bereits abgeglichen"* aus. Er nimmt sie jetzt. Die Zusage *„Restmenge mit
Zähler, kein offener Rest"* hat dadurch überhaupt erst gestimmt — sie war beim
Schreiben falsch.

## 9. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen.** Berührt ist **eine** Sub-Area, in
[`harness/conventions.md`](../../../../harness/conventions.md)
§Modus-Deklaration geführt:

- **`HARNESS`** (`AGENTS.md`, `CLAUDE.md`, `harness/`) — Achsen 1,2,3. Zwei der
  drei Paare liegen darin.

**Und eine Zuordnung, die zu prüfen war:** `docs/plan/planning/README.md` liegt
im Pfad der Sub-Area **`PLAN`**. Für die *Berührungs*-Frage zählt aber, wessen
Konventions-Härte der Abgleich betrifft — und das ist hier `HARNESS`: Der Index
deklariert, wie die Slice-Ablage funktioniert, nicht wie ein einzelner Slice
geplant wird. Ein Befund darin ändert eine Deklaration, keinen Plan. **Beide
Sub-Areas sind trotzdem geführt**, damit die Zuordnung nachprüfbar bleibt statt
implizit zu sein.

**Modus:** beide Greenfield; ein Begründungsblock entfällt.

**Vorgelagert — offene Beobachtungen sichten.** **Zwei Quellen, nicht eine** —
das ist die Lehre aus der Closure von slice-186 (§9 dort): Der Sichtungs-Schritt
las bisher nur das Register. `modul-05` §Closure- und Lerneintrag-Regeln nennt
aber **drei** Quellen, und die dritte ist die *wiederkehrende Finding-Klasse aus
dem Review*. Wer den Report des Vorgängers nicht liest, sieht sie nie.

**(1) Register — sechs Einträge betreffen diesen Slice, alle bei 2×:**

| Eintrag | Berührung durch diesen Slice |
|---|---|
| [`erstdurchgang-als-einmal-slice-geschnitten`](../observations/BEO-PLAN/erstdurchgang-als-einmal-slice-geschnitten/observation.md) | **beantwortet, nicht erfüllt** — der Zuschnitt ist vor der Arbeit korrigiert (§1); der Slice erzeugt damit **keinen** dritten Beleg |
| [`messung-ohne-reproduzierbares-instrument`](../observations/BEO-PLAN/messung-ohne-reproduzierbares-instrument/observation.md) | **beantwortet** — §2 trägt Parameter **und** Instrument; die Zahlen sind nachrechenbar |
| [`vollstaendigkeits-haken-ohne-erschoepften-gegenstand`](../observations/BEO-PLAN/vollstaendigkeits-haken-ohne-erschoepften-gegenstand/observation.md) | **als DoD-Punkt verdrahtet** — die Ebene gehört in den Ausgang, nicht nur das Ob |
| [`zielsatz-nach-plan-aenderung-nicht-nachgezogen`](../observations/BEO-PLAN/zielsatz-nach-plan-aenderung-nicht-nachgezogen/observation.md) | **vorweggenommen** — Titel und §1 tragen den korrigierten Umfang schon jetzt |
| [`kandidaten-klassifikation-groeber-als-der-kandidat`](../observations/BEO-PLAN/kandidaten-klassifikation-groeber-als-der-kandidat/observation.md) | **Risiko** — jede Sammelaussage über eine Kandidatenmenge braucht ihre Ausnahme im Satz |
| [`form-vergleich-sprachblind`](../observations/BEO-PLAN/form-vergleich-sprachblind/observation.md) | **kein Gegenstand** — keines der drei Paare hat eine zweite Sprachfassung |

**Keiner erreicht mit diesem Slice 3×** — vier davon, weil der Slice ihre Lehre
**vorab** anwendet statt sie erneut zu belegen. Das ist der Zweck des
Sichtungs-Schritts und nicht seine Umgehung: Ein Eintrag, dessen Lehre wirkt,
zählt nicht weiter.

**(2) Report des Vorgängers** — `docs/reviews/` zu slice-186, mit dem Slice
archiviert (`unzip -p docs/plan/planning/done/wellenlos/slice-186-archiv.zip`).
Er wies **vier** Finding-Klassen als zweite Wiederholung in Folge aus; alle vier
stehen jetzt oben im Register, und alle vier hängen an genau dieser Arbeitsform.
Dazu zwei Klassen, die **nicht** im Register stehen und für diesen Slice die
schärferen sind:

- *Selbstverweis beim Umzug nicht re-verankert* — trifft, sobald ein Befund
  einen Textblock **verschiebt**. Bei `AGENTS.md` ist das wahrscheinlich.
- *Zusage in der Doku, Grenze nur im Konfigurations-Kommentar* — trifft
  `harness/conventions.md` direkt: Es ist das Dokument, das Grenzen deklariert.

**Keine Treffer sind ebenfalls eine Antwort:** Zur Sub-Area `SPEC` steht nichts
Offenes an, das diesen Slice beträfe — sie ist ohnehin abgetrennt (§1).

**Nachtrag nach dem Review — die Sichtung war unvollständig, und das ist ihr
schwerster Befund.** §Sub-Area-Wahl oben nennt `HARNESS` als berührt; gesichtet
wurden **nur** die sechs `BEO-PLAN`-Einträge. Die **15** Einträge unter
`BEO-HARNESS` blieben ungelesen — und zwei davon treten in genau diesem Diff ein
drittes Mal auf (§3.6):

| Eintrag | Stand vorher | durch diesen Slice |
|---|---|---|
| [`BEO-HARNESS/chronik-in-gelesenen-dateien`](../observations/BEO-HARNESS/chronik-in-gelesenen-dateien/observation.md) | offen (2×) | **3×** → *geplant*, [slice-191](../open/slice-191-chronik-phrasen-sensor.md) |
| [`BEO-HARNESS/adaption-korrigiert-repo-aussage`](../observations/BEO-HARNESS/adaption-korrigiert-repo-aussage/observation.md) | offen (2×) | **3×** → *geplant*, [slice-190](../open/slice-190-id-schema-deklaration-ueberarbeiten.md) |
| [`BEO-HARNESS/baseline-regel-nie-erwogen-weil-bestand-sie-verletzt`](../observations/BEO-HARNESS/baseline-regel-nie-erwogen-weil-bestand-sie-verletzt/observation.md) | offen (2×) | **berührt, nicht erhöht** — sein Beleg führt denselben `AGENTS.md`-§4-Satz, den §3.1 als Befund 1 übernimmt; der Eintrag hatte ihn also schon benannt |
| [`BEO-GATE/zusage-weiter-als-ihre-durchsetzung`](../observations/BEO-GATE/zusage-weiter-als-ihre-durchsetzung/observation.md) | — | **neu, 2×** (§3.5) |
| [`BEO-PLAN/risiko-eingetreten-und-im-slice-aufgeloest`](../observations/BEO-PLAN/risiko-eingetreten-und-im-slice-aufgeloest/observation.md) | — | **neu, 1×** (§7, Risiko 4) |

**Die Lehre gehört an den Sichtungs-Schritt, nicht an diesen Slice:** Das
Register wird nach **allen** berührten Kürzeln gelesen, nicht nach dem, in
dessen Verzeichnis man gerade arbeitet. Die drei Folge-Slices (188/189/190/191)
tragen den Satz in ihrem §9 bereits.

**Und die zweite Quelle hat getragen:** Der Report des Vorgängers stand in §9 von
Anfang an und lieferte die zwei schärfsten Warnungen — *Selbstverweis beim Umzug*
und *Zusage mit Grenze nur im Konfigurations-Kommentar*. Die zweite ist §3.5
geworden. Was fehlte, war nicht die Quelle, sondern die **Breite** der ersten.
