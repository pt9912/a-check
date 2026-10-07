# Review-Report: slice-215 — 2026-10-07

**Review-Art:** Harness-/Doku-Review, unabhängiger Lauf. Der Kontext ist ein frischer Subagent
ohne `fork`, und er hat den Gegenstand nicht verfasst. Geprüft wird der Diff gegen den Slice-Plan
(einschließlich der Plan-Änderung in §1), die Hard Rules (`AGENTS.md` §3, besonders §3.7), die
Mess-Regeln und die Register-Form (`v6.13.0` · `regelwerk/modul-06-roadmap.md` §Das
Beobachtungs-Register, Paarung a). Gegen die DoD wird **nicht** geprüft, das ist Aufgabe des
Verifiers (Modul 11).

**Gegenstand:** slice-215, Commit-Range `ed45b49..1a5448b` — Inhalt in `1a5448b`
(docs(harness): vierte Mess-Regel); davor `ed45b49` (reiner `git mv`), `8275329` (Ruhe-Marker),
`315eaf8` (Plan-Änderung).

**Skill:** `.harness/skills/reviewer.md` @ `1a5448b` (sha256 `ec9ebaf2…4eddb92`) ·
**Modell:** `claude-opus-5-5` · **Datum:** 2026-10-07. Der Skill ist selbst Teil des
Gegenstands; angewandt wurde der Stand nach dem Diff, geprüft wurde er wie jeder andere Text.

> **Zitier-Form** *(dieser Block bleibt stehen. Er ist Norm, kein
> Ausfüll-Hinweis)*. Dieser Report friert ein, was er zitiert, bewegt sich
> weiter. Deshalb gilt **Kennung, nicht Adresse**: `slice-<Kennung>` statt seines
> Lifecycle-Pfads, `make <target>` statt eines Links auf die Sensor-Datei, eine
> Baseline-Stelle als **Tag + Pfad in Inline-Code** statt als Link. Ein
> `pfad`-Feld auf den **geprüften Gegenstand** ist davon nicht betroffen. Es
> zitiert den Stand des Laufs und darf ihn festhalten.

**Eingangs-Kontext:**

- Slice-Plan slice-215 (§1 Ziel, Plan-Änderung 2026-10-07, Abgrenzung; §2 DoD-Wortlaut der Regel;
  §3 vier Dateien; §8 Sichtung)
- `BEO-GATE/testbeschreibung-weiter-als-assertion`: `observation.md`, `state.md`, vier
  Evidence-Dateien (slice-210, slice-211, slice-213, slice-214)
- Die Review-Reports der Belege, aus den Archiven entpackt (Scratchpad, Repo unberührt):
  slice-210 (F-6), slice-213 (F-2, F-3, F-4), slice-214 (F-2, F-7, D-3)
- Der Original-Testkommentar von slice-210 in `fe2f8b2` (`internal/cli/`)
- `AGENTS.md` §3, §5, §6 Schritt 7; `CHANGELOG.md`; die Einführung von §6 Schritt 7 in
  `836323f`/`e65f86a` (slice-196) samt Review-Finding F-3 von slice-196 (Archiv)
- `v6.13.0` · `regelwerk/modul-06-roadmap.md` §Das Beobachtungs-Register und §Wellen-Closure-Prozedur
  Schritt 3 (Prosa-Form ausgeschöpft)

**Sonden:**

- `make doc-check`: Exit 0, „699 Datei(en) geprüft, 0 Befund(e)“ — Geltungsbereich: Links und
  Anker; Prosa-Aussagen sieht das Instrument nicht.
- `make verify-observations`: Exit 0, „79 Beobachtung(en) mit nicht leerem evidence/“ —
  Geltungsbereich: Deckung Register ↔ Zitate; Zielort und Herkunfts-Anker (Paarung a) prüft es
  nicht.
- **Zählstellen, zwei verschieden gebaute Zähler** (Mess-Regel 3), je über das Repo ohne `.git`,
  `done/`, `.harness/baseline/` und `docs/reviews/`:
  Zähler A — jede Zeile mit `Mess-Regel`/`Messregel` (Wortsuche, Groß-/Kleinschreibung egal);
  Zähler B — ein Zahlwort (`zwei|drei|vier|beide|alle drei|alle vier`) bis 40 Zeichen vor
  `urteilsregel|mess|regeln über` (Nachbarschaftssuche, findet auch Stellen ohne das Wort
  „Mess-Regel“). Beide finden dieselben vier Zählstellen, alle nachgezogen: `AGENTS.md:229`,
  `harness/rules/mess-regeln.md:1` und `:5`, `.harness/skills/reviewer.md:114` (nur B findet
  diese; A trifft dort die Überschrift, nicht die Zahl). Keine weitere Zahl der Mess-Regeln
  steht außerhalb von `done/`/Archiv; `kandidaten-klassifikation-groeber-als-der-kandidat/state.md:34`
  („die zwei älteren Mess-Regeln“) zählt relativ zu Regel 3 und bleibt wahr.
  **Verortung**, dritter Zähler (Regex auf `AGENTS…§5` über `observations/`, `harness/`,
  `.harness/skills/`, Planning-READMEs, `open/`, `next/`): Treffer siehe F-7.
- Herkunfts-Anker: `grep -c 'seit slice-1(79|81|93)'` — `harness/rules/mess-regeln.md` 3,
  `.harness/skills/reviewer.md` 3, `AGENTS.md` **0**; `seit slice-215` steht in
  `mess-regeln.md:23` und `reviewer.md:203`.

---

## Findings

| ID | Kategorie | Befund | Quelle | Pfad | Verifizierbar | Klasse |
|---|---|---|---|---|---|---|
| F-1 | MEDIUM | Die dritte Ausprägung nennt „slice-213 und slice-214“ mit den zwei Zitaten „alle fail-closed-Fälle der Spezifikation“ und „jede Form, die die Quelle verbietet“ und sagt „je drei Fälle fehlten“. Für slice-213 fehlte **ein** Fall: Evidence slice-213 „ohne den Fall ‚Backtick mitten im Token‘“, Report slice-213 F-2 „Der Fall mit `` ` `` fehlt im Test“. Drei fehlende Fälle trägt nur slice-214 (F-2: M1–M3; D-3: drei Mutationen). **Adversarisch geprüft, nicht auf HIGH gehoben:** Das erste Zitat steht auch in slice-214 (F-2), es gibt also eine Lesart, in der beide Zitate slice-214 gehören und „je drei“ stimmt. In der naheliegenden Zuordnung (erste Kennung ↔ erstes Zitat) ist die Zahl für slice-213 falsch. | Skill §Klassifikation „unbelegte Tatsachenbehauptung“; Mess-Regel 3 (die Sammelaussage deckt ihre Menge nicht) | `.harness/skills/reviewer.md:211-213` | ja — Evidence-Datei slice-213 lesen | Herleitung gibt eine Belegzahl falsch wieder |
| F-2 | LOW | Die Ausprägung heißt „der Allquantor geht über eine **offene** Menge“. Für „alle fail-closed-Fälle der Spezifikation“ ist die Menge geschlossen — die Spezifikation zählt sie auf (Report slice-213 F-2: „ein `//`, `"` oder `` ` `` direkt hinter einem Nicht-Leerraum-Zeichen“); der Test zählte eine kleinere auf. Offen ist die Menge nur beim zweiten Zitat („jede Form, die die Quelle verbietet“). Die Regel-Datei fasst denselben Fall anders und passend („gilt nur über eine Menge, die der Test aufzählt“); Skill und Regel benennen die Ausprägung damit verschieden. | Skill §Output-Schema (Klassen-Bezeichnung stabil); Mess-Regel 3 | `.harness/skills/reviewer.md:211`; `harness/rules/mess-regeln.md:26-28` | nein — Urteil über eine Einordnung | Ausprägungs-Bezeichnung gröber als ihr Beleg |
| F-3 | LOW | Die erste Ausprägung zitiert in Anführungszeichen „verbotenes Muster in Kommentar und Zeichenkette wird mitgeprüft“. Der Wortlaut steht an keiner anderen Stelle im Repo; der Testkommentar (`fe2f8b2`) lautete „ein verbotenes Muster nur im Kommentar oder in der Zeichenkette einer Anweisung ist kein eigener Befund“. „Die Fixture hatte keines“ ist gröber als Report slice-210 F-6: Ein Muster in einer Zeichenkette gab es, aber in einer **unerlaubten** Anweisung. Die Evidence-Datei slice-210 trägt dieselbe Vergröberung; der Skill übernimmt sie in Zitatform. | Skill §Klassifikation „unbelegte Tatsachenbehauptung“ (abgeschwächt: Sinn erhalten, Wortlaut nicht) | `.harness/skills/reviewer.md:207-208` | ja — `git show fe2f8b2` gegen den Zitat-Wortlaut | Paraphrase in Zitatform |
| F-4 | LOW | Regel und Herleitung binden „Testkommentar **oder Testname**“. Der DoD-Wortlaut in Plan §2 nennt nur den Testkommentar, und keiner der vier Belege betrifft einen Testnamen: Evidence slice-213/slice-214 sagen „`TestGomodUnsplittable` nannte sich …“, die Reports zeigen in beiden Fällen den **Kommentar** zu dem Test als Träger. Die Regel reicht damit weiter als Plan und Belege; eine Plan-Änderung dazu steht nicht in §1. | `AGENTS.md` §6 Schritt 4 (Abgrenzung nicht still weiten); Mess-Regel 1 (Geltungsbereich) | `harness/rules/mess-regeln.md:23-24`; `.harness/skills/reviewer.md:204` | nein — Urteil über die Reichweite | Regel reicht weiter als ihre Belege |
| F-5 | LOW | „Jede Eigenschaft, die ein Testkommentar oder ein Testname nennt, belegt eine Assertion“ ist im Deutschen in Subjekt und Objekt nicht entschieden (beide Nominalgruppen Nominativ = Akkusativ); die naheliegende Lesart (Subjekt zuerst) ist die umgekehrte der gemeinten: Die Eigenschaft belegt die Assertion. Der Wortlaut stammt aus Plan §2 und steht so in Regel und Skill. | Reviewer-Skill §LOW (Wording) | `harness/rules/mess-regeln.md:23-25`; `.harness/skills/reviewer.md:203-205` | nein | Wording |
| F-6 | LOW | Zwei neue Sätze tragen Chronik statt Zustand: Skill „Die vierte Auflage kam vor der Verkörperung — die Prosa-Form ist damit nicht ausgeschöpft, sondern **erst jetzt** geschrieben“ (deiktisches „jetzt“, gilt nur zum Schreibzeitpunkt); `state.md` „Der vierte Beleg kam vor der Verkörperung“. Die `Stand:`-Zeile selbst ist korrekt (Ausgang, Zielort, Anker). Die Ausnahme-Begründung gehört inhaltlich zur Grenze der Regel, steht aber als Erzählung. | `AGENTS.md` §3.7 (auch Zustandsfelder; Herkunft in **ein** auflösbares Feld) | `.harness/skills/reviewer.md:219-221`; `docs/plan/planning/observations/BEO-GATE/testbeschreibung-weiter-als-assertion/state.md:5-6` | nein — §3.7 ist inferentiell | Chronik in gelesener Datei |
| F-7 | LOW | Die Plan-Änderung erkennt die alte Verortung „Zusage in `AGENTS.md` §5“ als Fehler, misst sie aber nur im Skill-Kopf. Dieselbe Verortung tragen die `Stand:`-Zeilen der drei Geschwister-Einträge: `BEO-PLAN/review-geltungsbereich-zu-eng` („verkörpert in `AGENTS.md` §5 … `seit slice-179`“), `BEO-GATE/probe-liefert-den-gegenstand-mit` (`seit slice-181`), `BEO-PLAN/kandidaten-klassifikation-groeber-als-der-kandidat` (`seit slice-193`). `AGENTS.md` trägt keinen der drei Anker (0 Treffer); sie stehen in `harness/rules/mess-regeln.md`. Paarung (a) ist für diese drei damit nicht erfüllt. **Bestand seit slice-201, nicht durch diesen Diff entstanden**, außerhalb von Plan §3. | `v6.13.0` · `regelwerk/modul-06-roadmap.md` §Wellen-Closure-Prozedur Schritt 3 (a); Mess-Regel 1 (Geltungsbereich der Nachzug-Messung) | `docs/plan/planning/observations/BEO-PLAN/review-geltungsbereich-zu-eng/state.md:1`; `…/BEO-GATE/probe-liefert-den-gegenstand-mit/state.md:1`; `…/BEO-PLAN/kandidaten-klassifikation-groeber-als-der-kandidat/state.md:1` | ja — `grep 'seit slice-179' AGENTS.md` | Zielort im Register nach Umzug der Regel nicht nachgezogen |
| F-8 | INFO | Die Begründung für den fehlenden CHANGELOG-Eintrag („wie bei den drei Vorgänger-Regeln“) trägt über eine Präzedenz, die die Norm nicht prüfte: Die drei Regeln entstanden am 2026-09-07/08 (`d1e8022`, `b25b121`, `60e7b66`), der Satz in `AGENTS.md` §6 Schritt 7 kam erst am 2026-09-19 (`836323f`, slice-196). Das Ergebnis trägt dennoch über den Wortlaut: „öffentlicher Vertrag“, und „eine Regel“ kam laut Report slice-196 F-3 wegen **konsumenten-sichtbarer** Einträge hinzu (slice-038, slice-041). Ob `harness/rules/*.md` „eine Regel“ im Sinne von §6 Schritt 7 ist, sagt der Satz selbst nicht. | `AGENTS.md` §6 Schritt 7 | `docs/plan/planning/in-progress/slice-215-testkommentar-gegen-assertion.md:36-37`; `AGENTS.md:250-255` | nein | Begründung über Präzedenz statt Norm |
| F-9 | INFO | `state.md` legt eine Lesart der Baseline fest: „ein fünfter [Beleg] nach ihr wäre der Anlass, die Prosa-Form als ausgeschöpft zu werten“. Die Baseline sagt „Erreicht dieselbe Fehlerklasse trotzdem ein viertes Mal die Schwelle“; ob das ein weiteres Auftreten nach der Verkörperung meint oder ein viertes Erreichen der 3×-Schwelle, ist dort offen. Die Lesart ist vertretbar, hat aber keinen Anker, und keinen Wächter, der beim fünften Beleg auf sie schaut. Zuständig: Architect (Interpretation einer Baseline-Regel). | `v6.13.0` · `regelwerk/modul-06-roadmap.md` §Wellen-Closure-Prozedur Schritt 3 („Verkörpert heißt nicht zwangsläufig automatisiert“) | `docs/plan/planning/observations/BEO-GATE/testbeschreibung-weiter-als-assertion/state.md:6` | nein | Baseline-Lesart ohne Anker |

## Negativbefunde

| Bereich | Ergebnis |
|---|---|
| Ausprägung 2 gegen die Belege | geprüft, ohne Befund. slice-211 „Präfix statt Zeilennummer“ = Evidence slice-211 / Report F-5; slice-213 „`Contains` statt Anzahl“ = Report F-4; slice-214 „Fehler statt Exit-Code“ = Report F-7. |
| Zitate aus slice-213/slice-214 | geprüft, ohne Befund. „alle fail-closed-Fälle der Spezifikation“ und „jede Form, die die Quelle verbietet“ stehen wörtlich in den Reports (F-2 bzw. D-3); „Mutationen überlebten die ganze Suite“ trägt für beide (slice-213 F-2 `make test` grün; slice-214 F-2/D-3 „`ok` für alle sechs Pakete“). Zahl: F-1. |
| Belegzahl und Zähler | geprüft, ohne Befund. Vier Evidence-Dateien, „4×“ und die vier Kennungen stimmen in Skill, `state.md` und Plan überein; drei Funde in slice-213 und slice-214 je als **ein** Vorgang gezählt, wie die Evidence-Dateien sagen. |
| Register-Form `state.md` | geprüft, ohne Befund außer F-6/F-9. Ausgang *verkörpert* aus der geschlossenen Menge; Zielort `harness/rules/mess-regeln.md` Regel 4 als Link (Tiefe 6, löst auf); Herkunfts-Anker `seit slice-215` steht am Zielort (`mess-regeln.md:23`) und in der Herleitung (`reviewer.md:203`) — Paarung (a) erfüllt. `observation.md` unverändert. |
| Zählstellen der Mess-Regeln | geprüft mit zwei Zählern (siehe Sonden), ohne Befund: alle vier Zahl-Stellen nachgezogen, keine weitere außerhalb `done/`/Archiv. Verortung: F-7. |
| Skill-Kopf | geprüft, ohne Befund. „Vier Urteilsregeln“ = vier Bullets im Abschnitt; „Alle vier sagen selbst ‚kein Sensor‘“ — jeder der vier Bullets trägt **Kein Sensor**; Zeiger auf Regel-Datei und `AGENTS.md` §5 stimmen mit dem Bestand seit slice-201. |
| Kopf der Regel-Datei | geprüft, ohne Befund. „die Regeln 1–3 tragen den Wortlaut von dort“ ersetzt das jetzt falsche „Der Wortlaut ist unverändert“; Herkunft je Regel als `seit slice-<Kennung>`. |
| Sensor-Abgrenzung | geprüft, ohne Befund. Kein Sensor, kein Target, kein Produkt-Code — wie Plan §1; die Begründung „Urteil über Text (§3.7)“ deckt sich mit den drei Geschwister-Regeln. |
| Hard Rules `AGENTS.md` §3.1–§3.6 | geprüft, ohne Befund. Kein Go-Code, kein `//nolint`, kein Spec-Stratum, keine ADR, kein Gate gelockert; `ed45b49` ist ein reiner Rename ohne Inhalt, `1a5448b` ändert nur Inhalt (vier Dateien, kein Move). |
| Traceability und Commit-Scope | geprüft, ohne Befund. Alle vier Commits nennen `slice-215`; die drei `(planning)`-Commits berühren nur `docs/plan/planning/`. |
| Plan-Änderung vor der Arbeit | geprüft, ohne Befund. `315eaf8` steht vor `1a5448b` und nennt die zusätzlich berührten Stellen (Zeile 15, Titel, Skill-Kopf); Plan §3 nennt `AGENTS.md` §5 und den Skill-Kopf. Weitung über Plan-Wortlaut: F-4. |
| Spec-Referenz-Richtung | geprüft, ohne Befund. Kein Spec-Stratum berührt. |

## Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 0 |
| MEDIUM | 1 |
| LOW | 6 |
| INFO | 2 |

**Finding-Klassen dieses Laufs:** Herleitung gibt eine Belegzahl falsch wieder ·
Ausprägungs-Bezeichnung gröber als ihr Beleg · Paraphrase in Zitatform · Regel reicht weiter als
ihre Belege · Wording · Chronik in gelesener Datei · Zielort im Register nach Umzug der Regel
nicht nachgezogen · Begründung über Präzedenz statt Norm · Baseline-Lesart ohne Anker

## Verdikt

**Merge-blockierend:** ja, wegen F-1 (MEDIUM): Die Herleitung ist der „gemessene Bestand“, an dem
der Reviewer künftig urteilt, und sie gibt die Belegzahl für slice-213 falsch wieder. Die Regel
selbst (`harness/rules/mess-regeln.md` Regel 4) ist anwendbar: Sie nennt die Prüf-Handlung
(Mutation, die genau die Eigenschaft bricht) und eine Schreibweise für den Allquantor („die
aufgeführten“); ihre Schwächen sind Reichweite (F-4) und Satzbau (F-5), beide LOW. Register-Form
und Zählstellen sind korrekt nachgezogen.

**Übergabe:** Findings gehen an den Implementer. F-7 ist Bestand seit slice-201 außerhalb von
Plan §3; ob er hier mitgeht oder einen eigenen Vorgang bekommt, ist eine Plan-Frage. F-9 ist
eine Lesart-Frage an den Architect. Die Finding-Klassen gehen in die Slice-Closure §7. Dieser
Report ersetzt keine Verifikation; DoD-Konformität prüft der Verifier separat (Modul 11).
