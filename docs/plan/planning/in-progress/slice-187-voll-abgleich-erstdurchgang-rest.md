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

**Lerneintrag — Form:** wird bei Closure benannt.

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

**Was das Instrument nicht kann, und es sagt es selbst:** Es findet **wörtliche**
Übernahme. Eine sinngemäße Doppelung in anderen Worten sieht es nicht, und a-check
formuliert fast alles um — die Zahlen sind darum eine **Reihenfolge, kein
Befund**. Ein Kandidat ist erst ein Befund, wenn das Lesen ihn dazu macht.

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

**Ergebnis über alle 15 Ziel-Form-Paare** — die acht bereits abgeglichenen sind
mitgemessen, weil eine Zahl ohne Vergleichsmaßstab keine Größe ist:

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

*(entsteht mit der Arbeit)*

## 4. Definition of Done

- [ ] Die **drei** Harness-Deklarations-Paare sind abgeglichen; je Paar steht
      der Ausgang im Plan — übernommen · bewusst abweichend (mit Begründung) ·
      ohne Befund.
- [ ] Jede Übernahme ist als solche kenntlich, und **die Ebene steht im
      Ausgang**: auf welcher Auflösung geprüft wurde, nicht nur *ob*
      ([`BEO-PLAN/vollstaendigkeits-haken-ohne-erschoepften-gegenstand`](../observations/BEO-PLAN/vollstaendigkeits-haken-ohne-erschoepften-gegenstand/observation.md),
      2×).
- [ ] Die **Restmenge** ist mit Zähler übergeben: vier Paare, 123 Kandidaten,
      Folge-Slices mit Kennung.
- [ ] Unabhängiger Review durchgeführt (Report unter [`docs/reviews/`](../../../reviews/README.md)).
- [ ] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [ ] Beobachtungs-Register fortgeschrieben.
- [ ] Jedes Risiko aus §7 trägt einen Ausgang.

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
  Hälfte gelesen ist. — **Ausgang:** <offen bis Closure>
- **Vier Register-Einträge stehen bei 2× und treffen genau diese Arbeitsform.**
  Erreicht einer mit diesem Slice 3×, ist er keine Notiz mehr, sondern eine
  Lücke mit fälligem Ausgang — und vier auf einmal wären ein eigener Vorgang.
  — **Ausgang:** <offen bis Closure>
- **Die Vorauswahl bleibt eine Reihenfolge, kein Befund.** 124 Kandidaten sind
  keine 124 Befunde, und das Instrument sieht sinngemäße Übernahme nicht. Wer
  die Zahl für einen Umfang hält, plant falsch — in beide Richtungen.
  — **Ausgang:** <offen bis Closure>
- **`harness/conventions.md` ist Rang 9 und trägt die Adaptions-Einträge.** Eine
  Ziel-Form-Regel, die dort etwas ändert, kann einen akzeptierten `MR` berühren —
  und die sind immutabel. — **Ausgang:** <offen bis Closure>

## 8. Closure-Notiz

*(bei Closure auszufüllen)*

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
