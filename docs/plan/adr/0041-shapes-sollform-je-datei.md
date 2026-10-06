# ADR-0041: Sollform je Datei — Anweisungsvergleich nach Normalisierung statt Verbotsliste

**Status:** Proposed

**Datum:** 2026-10-06

**Autor:** pt9912 (Change Request und Abnahme der Entscheide), ausgeführt im Auftrag

**Bezug:** [AC-FA-RULE-012](../../../spec/lastenheft.md#ac-fa-rule-012),
[AC-FA-RULE-011](../../../spec/lastenheft.md#ac-fa-rule-011--konstrukt-monopol-regel-construct-leak)
(das Gegenstück),
[AC-FA-CONF-001](../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml),
[AC-QA-01](../../../spec/lastenheft.md#ac-qa-01--determinismus),
[AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze),
[ADR-0027](0027-constructs-roh-text-monopol.md) (Roh-Text-Monopol),
[ADR-0034](0034-stripcomments-string-literale.md) (Kommentar-Strip kennt Zeichenketten)

**Schärft:** [SPEC-CONF-001](../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema)
(Schema und fail-closed-Fälle des Blocks `shapes`),
[SPEC-EXTRACT-001](../../../spec/spezifikation.md#spec-extract-001--import-extraktion)
(Normalisierung und Zerlegung, Dialekt `kotlin`),
[SPEC-RULE-001](../../../spec/spezifikation.md#spec-rule-001--regel-auswertung)
(Vergleich und Befundklassen),
[ARC-003](../../../spec/architecture.md#2-komponenten) (die Extraktion liefert die Anweisungen).

**Regeln:** Baseline-Regelwerk `modul-04-adrs.md` §Ziel-Form: ADR (MADR).

---

## Kontext

Ein Adopter will sicherstellen, dass das Fachkern-Modul keine Fremdabhängigkeit und kein Plugin
bekommt. Die Grenze steht in einer Build-Datei (`build.gradle.kts`), nicht in einem Import.
a-check prüft heute zwei Ebenen: extrahierte **Importe** (Schichten, Kanten, `tech`) und
**Roh-Text** gescannter Quellen (`constructs`, [ADR-0027](0027-constructs-roh-text-monopol.md)).
Beide sind **Verbotslisten**: gemeldet wird, was man vorher als Muster aufgeschrieben hat.

Für Build-Manifeste trägt das nicht. Das Werkzeug kennt viele gleichwertige Schreibweisen, eine
Abhängigkeit oder ein Plugin einzutragen — `plugins` mit Zeilenumbruch vor `{`,
`compileClasspath += files(…)`, `libraries.from(…)`, `freeCompilerArgs.add("-Xplugin=…")`,
`add("…", <beliebiger Ausdruck>)`. Eine Verbotsliste solcher Muster wird nie vollständig, und ihre
Lücke ist **Schweigen**, nicht Fehlalarm. Dieselbe Klasse hat a-check an den eigenen Sensoren
beobachtet: ein Prüf-Muster trifft die häufige Schreibweise seines Gegenstands, die seltene
gleichwertige geht still durch.

Nötig ist die umgekehrte Beweislast: In dieser Datei steht nur, was ausdrücklich erlaubt ist.

**Annahmen, an denen die Entscheidung hängt:** (1) Die zu schützenden Dateien sind **klein und
selten geändert** — eine Liste erlaubter Anweisungen ist pflegbar. (2) Ein **textueller**
Vergleich nach Normalisierung genügt, weil die Frage *„steht hier etwas, das nicht auf der Liste
steht?"* lautet, nicht *„was bewirkt es?"*. Kippt (1), ist die Regel Pflegelast; kippt (2), braucht
es eine Prüfung im Build, und die ist ausdrücklich außerhalb.

## Entscheidung

Wir wählen einen **Anweisungsvergleich nach dialekt-spezifischer Normalisierung**, als eigene
Regelart `shapes` neben `constructs`, mit zwei Modi.

1. **Eigene Regelart, kein Ausbau von `constructs`.** `constructs` beantwortet *„wo darf diese
   Form stehen?"* (Zone), `shapes` *„was darf in dieser Datei stehen?"* (Liste). Ein Schalter in
   `constructs`, der die Beweislast umdreht, hätte zwei Semantiken unter einem Schlüssel.
2. **Die Einheit ist die Anweisung auf oberster Ebene, ein Block zählt als Ganzes.** Wer einen
   erlaubten Block als Ganzes vergleicht, kann keine weitere Anweisung in ihn schmuggeln — die
   Liste muss den Blockinhalt vollständig nennen.
3. **Normalisierung je Dialekt, fest im Werkzeug.** Kommentare fallen weg (Zeile, Block,
   verschachtelt), Zeichenketten bleiben **byte-genau** stehen (inklusive Roh-Strings und
   `${…}`-Vorlagen, deren Ende über die Klammer-Tiefe gefunden wird), Leerraum außerhalb von
   Zeichenketten fällt weg — **außer** er trennt zwei Wortzeichen oder zwei Operatorzeichen; dann
   wird er zu genau einem Leerzeichen. Die zweite Hälfte schärft den abgenommenen Entscheid
   „Leerraum zwischen Bezeichner-Zeichen bleibt" um denselben Gedanken für Operatoren
   (`a - -b` ≠ `a--b`); sie macht den Vergleich nur strenger.
4. **Die Anweisungsgrenze ist eine Fortsetzungsregel, keine Grammatik.** Auf oberster Ebene und
   innerhalb von `{…}` trennt `;` immer und ein Zeilenende, sofern keine `(`/`[` offen ist, die
   nächste Nicht-Leer-Zeile nicht mit `{`, `.` oder `?.` beginnt und die vorige nicht auf ein
   Operatorzeichen oder `,` endet. Innerhalb eines Blocks wird jede Grenze zu `;` normalisiert —
   sonst fielen `a()` und `b()` auf zwei Zeilen mit `a()b()` zusammen. **Beide Fehlrichtungen
   sind fail-safe:** zu fein getrennt heißt, jedes Stück muss erlaubt sein; zu grob getrennt
   heißt, das Ganze muss erlaubt sein. Ein Fehler der Regel kann eine Datei rot färben, nie
   eine unerlaubte Anweisung grün.
5. **Literal-Einträge werden wie die Datei normalisiert, Regex-Einträge sind voll verankert.**
   Eine Liste darf darum lesbar, mehrzeilig, mit Leerraum geschrieben werden. Ein Regex trifft
   die **ganze** normalisierte Anweisung (implizit `^…$`); ein unverankertes `dependencies\{.*`
   erlaubte sonst jeden Blockinhalt.
6. **Zwei Modi, eine Normalisierung.** `allow-statements` vergleicht als **Menge** (Reihenfolge
   und Wiederholung gleichgültig); `exact` vergleicht die Anweisungsfolge mit einer
   Sollform-Datei desselben Dialekts und meldet die **erste** Abweichung.
7. **Kein Warn-Level.** Ein `allow`-Eintrag ohne Treffer wird nur auf Wunsch gemeldet
   (`unused: fail`), dann als Befund mit Exit 1 — a-check kennt keinen dritten Schweregrad, und
   ein neuer Default-Befund färbte bestehende Konfigurationen rot.
8. **Fail-closed, wo das Werkzeug nichts prüfen kann.** Ein Glob ohne Treffer, eine Datei, die
   zugleich `exclude` ausnimmt, und eine Datei, die sich nicht zerlegen lässt (offene Klammer,
   Zeichenkette, Kommentar; schließende Klammer ohne Gegenstück), sind Exit 2. Eine nicht
   zerlegbare Datei still grün zu lassen wäre genau die Lücke, gegen die die Regel steht.
9. **Platz im Hexagon:** Lesen, Normalisieren und Zerlegen leistet die **Extraktion**
   ([ARC-003](../../../spec/architecture.md#2-komponenten)) — sie ist Text-Heuristik je Dialekt,
   wie die Sprach-Backends je Sprache, und kennt Dateien. Sie liefert dem Kern je Datei die
   Anweisungen mit Original-Zeile und die normalisierten Literal-Einträge. Der **Kern** vergleicht
   und erzeugt die Befunde — rein, ohne I/O. Ein Zerlegungsfehler ist ein Fehler der Extraktion
   und erreicht die Composition Root als Exit 2, wie die Mehrdeutigkeit der Mehr-Wurzel-Auflösung.
10. **Ein Dialekt: `kotlin`.** Der generische Dialekt (Kommentar-, Zeichenketten- und
    Trennzeichen konfigurierbar) wartet auf einen zweiten Konsumenten; der Schlüssel `dialect`
    ist als geschlossene Menge angelegt und wächst additiv.

## Verglichene Alternativen

| Option | Pro | Contra |
|---|---|---|
| A — nichts tun; Muster in `constructs` nachtragen | kein neuer Vertrag; vorhandene Mechanik | Verbotsliste bleibt unvollständig; die Lücke ist Schweigen. Jede neue Schreibweise des Werkzeugs ist ein stiller Durchlass |
| B — fail-closed **Import-Allowlist** je Schicht | dieselbe Beweislast-Umkehr; im Out-of-Scope der Roh-Text-Anforderung bereits benannt | prüft **Importe** gescannter Quellen — ein Build-Manifest hat keine; der Leitfall ist damit nicht erreichbar. Bleibt ein eigener, gated Faden |
| C — nur `exact` gegen eine Sollform-Datei | kleinster Vertrag; kein Listen-Vokabular | spröde: jede inhaltlich harmlose Umordnung ist rot; keine Mengen-Semantik, keine Regex-Einträge für Versions-Literale |
| D — Build-Werkzeug befragen (Tooling-API, Abhängigkeits-Report) | sieht die wirkliche Semantik, auch was Plugins eintragen | bricht die Hermetik (Netz, Toolchain im Image); a-check prüft fremde Bäume ohne deren Werkzeug. Gehört in eine Prüfung im Build |
| E — vollständiger Kotlin-Parser | exakte Anweisungsgrenzen | widerspricht der Text-Heuristik-Linie; großer, sprachspezifischer Code für eine Frage, die der Vergleich nach Normalisierung fail-safe beantwortet |
| **F — Anweisungsvergleich nach Normalisierung, zwei Modi (gewählt)** | Beweislast umgekehrt; fail-safe in beiden Fehlrichtungen der Zerlegung; hermetisch; Normalisierung macht Formatierung irrelevant | heuristische Anweisungsgrenze kann legitime Dateien rot färben; Liste muss den Blockinhalt vollständig nennen; Pflegelast bei häufig geänderten Dateien |

## Konsequenzen

- **Positiv:** Eine Fremdabhängigkeit im Fachkern-Modul ist rot, gleich in welcher Schreibweise —
  das Werkzeug muss die Schreibweise nicht kennen. Formatierung, Kommentare und Zeilenumbrüche
  ändern das Urteil nicht.
- **Positiv:** Die Regel ist von `layers` und `languages` unabhängig; ein Konsument kann sie
  einsetzen, ohne seine Schicht-Konfiguration anzufassen.
- **Negativ:** Ein Fehlalarm der heuristischen Anweisungsgrenze an einer legitimen Datei kostet
  Vertrauen; eine Regel, die zu oft rot ist, wird abgeschaltet. Die Gegenprobe-Fälle des CR
  werden darum Tests, und jeder gemeldete Fehlalarm ist ein Re-Evaluierungs-Anlass.
- **Negativ:** Die Liste muss einen erlaubten Block vollständig nennen; ein neues erlaubtes Detail
  im Block ist eine Listen-Änderung, keine Zeile mehr.
- **Folgepflicht:** Spezifikation (Schema, Normalisierungsvertrag, Befundklassen), Architektur
  ([ARC-003](../../../spec/architecture.md#2-komponenten)), `--print-config`-Gerüst und Benutzerhandbuch werden nachgezogen; die
  Implementierung folgt in zwei Schritten (`allow-statements` zuerst, dann `exact` und `unused`).

## Fitness Function (falls maschinell prüfbar)

| Tooling | Regel | Make-Target |
|---|---|---|
| Go-Tests | jeder Gegenprobe-Fall der Anforderung als eigener Testfall; Normalisierer-Tests je Lexik-Fall (verschachtelte Kommentare, Roh-Strings, `${…}` mit Klammern, Zeichen-Literal `'{'`) | `make test` |
| a-check selbst | der Kern bleibt rein — die Normalisierung liegt in der Extraktion | `make arch-check` |

## Re-Evaluierungs-Trigger

- Ein **zweiter Konsument** will eine Nicht-Kotlin-Datei prüfen (`go.mod`, `package.json`, …) —
  dann generischer Dialekt, als Folge-ADR, die Punkt 10 ablöst.
- Ein **gemeldeter Fehlalarm** an einer legitimen Datei, den die Fortsetzungsregel (Punkt 4)
  verursacht — dann Grenze schärfen oder ausweisen.
- Ein Adopter braucht die **Semantik** dessen, was ein erlaubtes Plugin einträgt — dann ist die
  Frage eine Build-Prüfung, nicht diese Regel.

## Geschichte

| Datum | Ereignis | Verweis |
|---|---|---|
| 2026-10-06 | Proposed | Change Request „Positivliste von Anweisungen je Datei (Sollform)" (Maintainer) |

Nach `Accepted` wird diese Datei **nicht mehr inhaltlich überschrieben**.
Spätere Korrekturen oder Schärfungen entstehen als neue ADR mit
`Supersedes ADR-NNNN` (Baseline-Regelwerk `modul-04-adrs.md`
§Hard Rule für Accepted-ADRs).
