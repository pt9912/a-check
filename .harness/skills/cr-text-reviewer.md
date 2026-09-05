# CR-Text-Reviewer-Skill — a-check

* Status: Accepted
* Bezug: [`BEO-022`](../../docs/plan/planning/observations/BEO-GATE/cr-text-behauptet-statt-gemessen/observation.md) (2× bei Verkörperung — die Migration auf die Verzeichnisform korrigierte einen doppelt gezählten Vorgang, siehe dort),
  Regelwerk `modul-07-diskrepanzen.md` (Diskrepanz-Trichter), `modul-13-quality-gates.md`
  §Harness-Lüge
* Gilt für: einen **CR-Text an ein fremdes Werkzeug**, bevor er den Slice verlässt — die vier
  bisherigen entstanden in [slice-073](../../docs/plan/planning/done/welle-12/slice-073-dcheck-statt-eigenbau.md) §8
  und [slice-080](../../docs/plan/planning/done/wellenlos/slice-080-verify-abloesung-dcheck.md) §8,
  [slice-116](../../docs/plan/planning/done/wellenlos/slice-116-nullmengen-haerte-cr.md) §9
* Entstanden: slice-119

> **Warum ein Guide und kein Sensor.** *„Ist diese Behauptung gemessen?"* ist ein Urteil über den
> **Entstehungsweg** eines Satzes, nicht über seinen Text. Dieselbe Grenze benennt
> [`AGENTS.md`](../../AGENTS.md) §3.7 für die Kommentar-Regel. Ein Guide ist hier zudem die
> **erste** Antwort auf die Schwelle — anders als bei
> [`BEO-008`](../../docs/plan/planning/observations/BEO-PLAN/verweis-auf-wandernden-slice/observation.md), wo ein Guide schon einmal gescheitert war
> und darum ein Werkzeug entstand
> ([slice-118](../../docs/plan/planning/done/wellenlos/slice-118-lifecycle-wechsel-werkzeug.md)).

## Kontext-Eingang (Pflicht)

Was der Reviewer *immer* mitbringt, bevor er urteilt:

- den CR-Text selbst, vollständig
- das **Lastenheft des Zielwerkzeugs** — es ist öffentlich; jede Behauptung über sein Verhalten
  ist dort belegbar oder falsch
- den **eigenen Bestand**, auf den sich der Antrag beruft (Datei, Zahl, Konfiguration)

Ohne den zweiten Punkt prüft der Reviewer Plausibilität statt Belegbarkeit — genau der Fehler, den
dieser Skill verhindert.

## Die zwei Ausprägungen

Beide sind belegt, und die zweite ist die gefährlichere.

**(A) Gar nicht gemessen.** Eine Behauptung, die eine Messung zuließe, steht als Annahme da.

**(B) Die falsche Menge gemessen.** Es *wurde* gemessen — nur nicht das, worüber der Satz redet.
Diese Form **sieht aus wie ein Beleg** und kann im Ergebnis sogar zutreffen; sie ist deshalb
schwerer zu fangen als (A). Der Empfänger führt sie bei sich als eigene Klasse: *„gemessen wird
die eigene Menge, ausgesagt wird über die fremde"*.

Die Prüf-Frage ist darum nicht *„hast du gemessen?"*, sondern **„hast du *das* gemessen, worüber
du redest?"**

## Prüf-Auftrag

Lies den CR-Text. Markiere **jeden Satz, der eine Tatsache über ein System behauptet** — das
eigene oder das fremde. Die drei **bisher belegten** Klassen und ihr Handgriff:

| Klasse | Beispiel-Behauptung | Handgriff |
|---|---|---|
| **eigener Bestand** | „hier sind es 19 Anforderungen" | `grep -c` über die Datei |
| **fremder Vertrag** | „das Werkzeug kennt keine nicht-fatalen Befunde" | den Abschnitt im fremden Lastenheft lesen, mit Kennung zitieren |
| **eigener Gate-Pfad** | „der Preis ist nur schwächer" | `grep` über `Makefile`, `.github/workflows/`, `.githooks/` |

**Eine vierte Klasse ist möglich** — die Liste ist gemessen, nicht hergeleitet. Wer eine findet, trägt sie hier nach und legt eine weitere `evidence/<vorgangs-id>.md` unter `BEO-022` an.

**Ein Satz ohne Handgriff ist kein Befund** — es gibt Argumente, die keine Tatsache behaupten
(„das wäre eine Vertragsänderung, größer als dieser Antrag"). Sie bleiben stehen.

## Die vier belegten Fälle als Fixtures

Wer den Skill prüft, prüft ihn gegen diese — drei aus eigenen CRs, einer vom Empfänger:

1. **CR 3, eigener Bestand.** Beantragt, 19 von 19 Abschnitten auszunehmen, ohne zu zählen, dass
   danach keiner übrig bleibt. Ein `grep -c '^### AC-'` hätte es gezeigt. → Klasse *eigener
   Bestand*, Ausprägung **(A)**.
2. **CR 4, erster Entwurf, fremder Vertrag.** Machte eine „nicht-fatale Zeile" zur Bedingung;
   `DC-FA-CLI-003` führt differenzierte Exit-Codes als **Out-of-Scope**. → Klasse *fremder
   Vertrag*, Ausprägung **(A)**.
3. **CR 4, zweiter Entwurf, eigener Gate-Pfad.** Nannte den Preis der `--doctor`-Wahl
   *„schwächer"*; gemessen ist er **null** — kein Gate fährt `doc-doctor`. → Klasse *eigener
   Gate-Pfad*, Ausprägung **(A)**.
4. **Die Antwort des Empfängers, Ausprägung (B).** Eine Tabelle mit vier Messzeilen, deren vierte
   ungemessen war: die eigenen Gate-Dateien wurden gemessen, ausgesagt wurde über das *fremde*
   Fragment. Sie traf im Ergebnis zu und war trotzdem kein Beleg. → Klasse *fremder Vertrag*,
   Ausprägung **(B)**.

**Die Gegenprobe gehört dazu.** Dieselbe Sitzung, die (1)–(3) produziert hat, hat den Pin-Bump,
die Modul-Parität und die Backtick-Falle sauber vorab gemessen. Es ist also kein Aufwandsproblem:
gemessen wurde, was **Arbeit** war; behauptet wurde, was **Argument** war. Der Skill greift
darum genau dort, wo der Text *überzeugen* will.

## Was der Reviewer NICHT tut

- **Den Antrag inhaltlich bewerten.** Ob die beantragte Fähigkeit sinnvoll ist, entscheidet der
  Empfänger; dieser Skill prüft die **Belegbarkeit**, nicht die Sache.
- **Wortlaut glätten.** Ein unbelegter Satz wird belegt oder gestrichen, nicht abgeschwächt —
  *„vermutlich schwächer"* ist derselbe Fehler mit Weichzeichner.
- **Bestehende CR-Texte nachprüfen.** Ihre Fehler sind benannt und beantwortet.

## Ausgabe

Je markiertem Satz eine Zeile:

```
<Abschnitt> :: <zitierter Satz> :: <Klasse> :: <Handgriff, der ihn belegt oder widerlegt>
```

Und ein Schluss-Satz: **wie viele Tatsachen-Behauptungen der Text trägt und wie viele davon
belegt sind.** Eine Zahl ohne die andere ist keine Auskunft.
