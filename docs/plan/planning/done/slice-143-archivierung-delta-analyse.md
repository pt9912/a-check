# slice-143 — Zeitdokumente-Archivierung: Delta-Analyse

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese Datei liegt — eines von
`open/`, `next/`, `in-progress/`, `done/`. Er wechselt nur durch `make slice-mv`
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).

**Bezug:** Maintainer-Fund 2026-09-05: `docs/reviews/` und `docs/plan/planning/done/` sind nicht
archiviert, obwohl `modul-06-roadmap.md` (`v6.0.0`) das jetzt als Closure-Schritt 4 führt — im
Schwester-Repo `d-check` bereits umgesetzt. Präzedenz für die Analyse-Form:
[slice-135](../done/slice-135-regelwerk-v600-delta-analyse.md), [slice-092](../done/slice-092-regelwerk-v5120-delta-analyse.md).
[Roadmap](../in-progress/roadmap.md).

> **Analyse zur Abnahme.** Wie bei slice-135/092: keine Kennungen vergeben, keine Artefakte
> geändert. Quelle ist die vendorte Baseline selbst — `d-check`s Umsetzung ist Präzedenz-Check,
> nicht Vorlage (`d-check` hat einen eigenen `MR`-Bestand, eigene Sub-Areas, eigene
> Reconciliation-Historie; keiner ihrer fünf `MR-05[9]`/`06[1-4]`-Einträge hat ein Gegenstück hier).

**Berührte Spec-Stellen:** — *(keine)* — reine Ist-Messung, keine Artefakt-Änderung.

**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.

**Autor:** Claude (Sonnet 5), im Auftrag des Maintainers. **Datum:** 2026-09-05.

---

## 1. Was die Baseline verlangt — gelesen, nicht angenommen

Drei Fundstellen in `.harness/baseline/v6.0.0/regelwerk/modul-06-roadmap.md`, wörtlich:

- **Wellen-Closure-Prozedur, Schritt 4:** „Zeitdokumente der Welle archivieren. Die Slice-Dateien,
  die sie einsammelt, ihr eigener Plan und die Review-Reports dieser Slices wandern in ein
  unveränderliches Archiv `done/<welle-id>/archiv.zip`." Slices und Welle-Plan bleiben als
  **gekürzter Stub**; Review-Reports bekommen **keinen** Stub.
- **Wellenlos-Träger-Tabelle** (§Wann Arbeit eine Welle braucht): „**Zeitdokumente archivieren**
  (Closure-Schritt 4) | Slice-Closure | **nach** den Paarungen … Schlüssel ist der Slice:
  `done/slice-<NNN>-archiv.zip`, **flach** neben dem Stub." Kein Optional-Vermerk an dieser Stelle
  — anders als beim Wellen-Fall (siehe unten) ist das für **künftige** wellenlose Slice-Closures
  einfach ein weiterer Schritt der Prozedur.
- **Zwei ausdrückliche Freiheiten für den Altbestand:** „**Kein Zwang zum Nachrüsten — und kein
  Verbot:** Wellen, die vor der Einführung schlossen, müssen nicht archiviert werden; ein Repo
  bleibt ohne das konform. Wer den Altbestand loswerden will, führt die Archivierung als eigenen
  Vorgang aus … Das Repo benennt die Zuordnung — die chronologisch nächste geschlossene Welle oder
  ein einzelnes Sammel-Archiv für den Bestand vor der Einführung." Und: „**Vor der ersten
  Archivierung ist der Geltungsbereich der vorhandenen Sensoren zu prüfen.**"

**Einordnung — kein Adaptions-Fall.** Das ist kein `MR`-Kandidat wie die Replay-Lauf-Ersetzung
([`MR-008`](../../../../harness/conventions/done/MR-008-kein-replay.md)/[`MR-014`](../../../../harness/conventions/MR-014-keine-agenten-telemetrie.md)/[`MR-015`](../../../../harness/conventions/MR-015-welle-closure-ohne-replay.md)):
Es gibt nichts abzulehnen, nur eine Umsetzungsreihenfolge zu wählen. Für
**neue** Slice-Closures ist der Schritt ab jetzt Teil der Prozedur, sobald a-check `v6.0.0` als
Stand führt (seit [slice-136](../done/slice-136-baseline-v600-vendoring.md)) — er wurde in
slice-135/139 fälschlich als „optional, kein Trigger" eingeordnet, weil dort nur der
**Wellen**-Fall gelesen wurde (§4.2 dort), nicht die separate, unbedingte Wellenlos-Zeile.
**Korrektur:** Für den **Altbestand** gilt die zitierte Freiheit uneingeschränkt — nichts davon ist
rückwirkend Pflicht.

## 2. Ist-Stand — gemessen

| | |
|---|---|
| Slice-Dateien in `done/` | **140** (`slice-001`…`slice-142`, mit Lücken) |
| Review-Reports in `docs/reviews/` | **34** |
| Wellen mit formaler Closure-Notiz (`welle-<NN>-results.md`) | **2** — `welle-12` (33 Slices,
  `slice-046`…`078`), `welle-13` (6 Slices, `slice-081`…`086`) |
| Slices **ohne** formale Wellen-Closure-Notiz | **101** — die in der Roadmap als „retroaktiv
  nachgezogen" deklarierten frühen Wellen (`welle-00`…`welle-11`, vor der reguläreren
  Lifecycle-Praxis „ab `slice-004`") plus alle wellenlosen Slices seit `welle-13` (2026-08-09) |

**Der Altbestand ist nicht einheitlich.** Nur `welle-12` und `welle-13` haben durchlaufen, was
Modul 6 als Wellen-Closure-Prozedur meint (Schritt 3: Ergebnisnotiz **und** `git mv` der
Welle-Plan-Datei — für `welle-13` laut Roadmap „erster Durchlauf mit Welle-Plan-Datei"). Die elf
früheren „Wellen" sind Prosa-Überschriften der Roadmap, keine durchlaufene Prozedur — dieselbe
Feststellung, die `docs/plan/planning/README.md` §Ab wann das gilt bereits für die
**Sechs-Schritte**-Form selbst trifft: „Die zwölf bestehenden Wellen … existieren ausschließlich
als Prosa-Überschriften … sie bekommen **keine** rückwirkenden Ergebnis-Notizen."

## 3. Sensoren mit Geltungsbereich auf `done/` oder `docs/reviews/` — die verlangte Vorprüfung

Gemessen mit `grep`, nicht angenommen:

| Fundstelle | Glob / Pfad | Betroffen? |
|---|---|---|
| `.d-check.yml` `structure` (Closure-Struktur, Lerneintrag-Form) | `docs/plan/planning/done/slice-*.md` | **Ja** — ein Stub hat keine `## Closure`-Überschrift mehr; ein Sensor, der sie verlangt, meldete jeden archivierten Slice als Formfehler, sobald er ihn sieht |
| `.d-check.yml` `structure` (Kopffelder, aus slice-139 um `observations/**` ergänzt) | `docs/plan/planning/**/slice-*.md` | **Ja** — dieselbe Kollisionsklasse wie bei `observations/**/evidence/` (slice-139 §… ), diesmal mit Stub-Dateien statt Evidence-Dateien |
| `.d-check.yml` `planning` (Modul, Zeile 140) | `docs/plan/planning/done` | **Zu prüfen** — welche der drei Modul-Fähigkeiten (`closure:`, `waves:`, `drift`) diese Zeile trägt, ist hier nicht ausgewertet |
| `tools/verify-risiko-ausgaenge.sh` | `docs/plan/planning/done` (+ `in-progress/`) | **Ja** — liest den Risiko-Block jedes Slice; ein Stub hat keinen mehr |
| `tools/verify-observations.sh` | `docs/plan/planning/done` (Zitat-Scan) | **Ja** — durchsucht `done/` nach `BEO-*`-Zitaten; ein Stub muss die Zitate seines Volltexts weiterhin sichtbar tragen oder der Sensor verliert sie |
| `tools/commit-scope-check.sh` | `docs/plan/planning/` (Scope-Grenze `(planning)`) | **Zu prüfen** — ein Archivierungs-Commit ist reine Planning-Bewegung, sollte in den Scope fallen |
| `tools/slice-mv.sh` | Verweis-Nachzug repo-weit | **Ja, aber anders** — kein Archivierungs-Ziel selbst, aber sein Verweis-Nachzug muss weiter funktionieren, wenn ein Stub an der Stelle liegt, auf die jemand verweist |
| `.d-check.yml` (fünf Immutabilitäts-/Matrix-Regeln) | `exempt-paths: […, "docs/reviews/**"]` | **Neutral** — Reviews sind bereits von der Referenzmatrix ausgenommen; ein Archiv ändert daran nichts, solange keine Stub-Form für Reviews existiert (die Baseline sieht explizit keine vor) |

**Sieben Fundstellen, davon fünf mit wahrscheinlichem Anpassungsbedarf.** Keine ist hier bereits
geändert — das ist die Aufgabe einer Umsetzungs-Etappe, nicht dieser Analyse.

## 4. Was `d-check` tat — Präzedenz, nicht Vorlage

Gelesen, um die eigene Schätzung zu kalibrieren, nicht um sie zu ersetzen:

- Ein eigenes Go-Werkzeug `tools/archive-wave/` (273 Zeilen + 278 Zeilen Test), das Sammeln, Zip-Bau,
  Stub-Schreiben und Verweis-Nachzug in einem Lauf bündelt.
- **Fünf** eigene Wellen (`welle-86`…`welle-90`): Werkzeug bauen und auf die letzte offene Welle
  anwenden → auf ältere geschlossene Wellen rückwirkend anwenden → wellenloser
  Einzel-Slice-Modus nachrüsten → eigenständige Review-Archivierung (Reviews ohne zugehörigen
  archivierten Slice) → verbliebene Nachzügler.
- **Fünf** eigene `MR`-Einträge (`059`, `061`–`064`) für Repo-spezifische Präzisierungen
  (Bündelungs-Modus, Register-Migrations-Move, wellenloser Slice-Move, eigenständiger
  Review-Move) — keiner davon ist ohne Prüfung auf a-check übertragbar, weil a-checks
  `MR`-Bestand, Sub-Area-Kürzel und Reconciliation-Historie eigenständig gewachsen sind
  (dieselbe Vorsicht wie in [slice-135 §8](../done/slice-135-regelwerk-v600-delta-analyse.md#8-closure-notiz)
  gegenüber der `welle-88`-Analogie).

**Übertragbar ist die Reihenfolge der Fragen**, nicht die Lösung: Werkzeug vor Rückwirkung, Wellen
vor wellenlos, verbundene Reviews vor eigenständigen. Die konkreten Zahlen, Bündelungs-Entscheidungen
und `MR`-Texte sind a-check-eigen zu treffen.

## 5. Was diese Analyse **nicht** geklärt hat

- **Ob a-checks 101 unformalisierte Slices als ein Sammel-Archiv oder einzeln wellenlos
  archiviert werden.** Die Baseline erlaubt beides ausdrücklich (§1); die Wahl ist Repo-Entscheidung,
  hier nicht getroffen.
- **Welche der 34 Reviews zu welchem Slice gehören** — nicht gezählt, nur als offene Frage benannt.
  Die meisten Dateinamen tragen eine `slice-NNN`-Kennung, drei nicht (`2026-06-21-benutzerhandbuch.md`,
  `2026-07-26-etappe-*-slice-*.md`-Sammel-Reviews, `2026-08-09-welle-12-unabhaengig.md`) — deren
  Zuordnung ist Lesearbeit, keine Messung.
- **Werkzeug-Wahl.** Ob ein eigenes Go-Werkzeug (wie `d-check`) angemessen ist oder a-checks
  kleinerer Bestand (140 vs. `d-check`s deutlich größerer, nicht gemessener Bestand) ein
  Bash-Skript rechtfertigt, ist nicht entschieden.
- **Die Reihenfolge der fünf betroffenen Sensor-Fundstellen aus §3** — welche zuerst, welche
  gemeinsam mit dem ersten Archivierungs-Lauf.

## 6. Vorschlag: vier Etappen

| Etappe | Inhalt | Warum getrennt |
|---|---|---|
| **A — Werkzeug** | Archivierungs-Mechanik bauen (Sammeln, Zip, Stubs, Verweis-Nachzug), gegen ein Test-Fixture, nicht gegen den echten Bestand | isoliert prüfbar, kein Risiko für `done/` |
| **B — Sensor-Geltungsbereich** | die fünf Fundstellen aus §3 anpassen oder ihre Nicht-Geltung für Stubs explizit dokumentieren | muss **vor** dem ersten echten Lauf stehen (Baseline-Vorschrift) |
| **C — `welle-12`/`welle-13` archivieren** | die beiden einzigen Wellen mit formaler Closure-Notiz — kleinster, klarster Fall (39 Slices) | Wellen-Fall zuerst, weil er die einzige Zuordnung ohne Entscheidungsbedarf ist |
| **D — Altbestand (101 Slices) + 34 Reviews** | Sammel-Archiv oder wellenlos, je nach Abnahme aus §5 | größte Etappe, hängt an einer noch offenen Entscheidung |

**Ab Etappe D+1 (künftig):** jede neue Slice-Closure archiviert sich selbst nach der
Wellenlos-Träger-Zeile — kein weiterer Nachzug, sondern Teil des Workflows.

## 7. DoD

- [x] Baseline-Anforderung wörtlich zitiert und von der eigenen Fehleinordnung in slice-135/139
      abgegrenzt (§1).
- [x] Ist-Stand gemessen (§2), Sensor-Geltungsbereich mit `grep`-Fundstellen belegt (§3),
      `d-check`s Umsetzung als Präzedenz gelesen und explizit von einer Vorlage abgegrenzt (§4).
- [x] Lücken benannt statt Vollständigkeit behauptet (§5); Etappen-Vorschlag steht (§6). `make
      gates` und `make verify` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie in
      eine Pipe.

## 8. Closure-Notiz

**Geliefert:** die Baseline-Anforderung zur Zeitdokumente-Archivierung wörtlich gelesen (nicht aus
`d-check`s Beispiel erschlossen), der eigene Fehler aus slice-135 §4.2/slice-139 §5 benannt (der
Wellenlos-Fall wurde übersehen, obwohl unbedingt formuliert), der Ist-Stand vermessen (140 Slices,
34 Reviews, nur 2 von 13 Wellen mit formaler Closure-Notiz), fünf von sieben geprüften
Sensor-Fundstellen als wahrscheinlich betroffen markiert, und ein Vier-Etappen-Vorschlag, der
`d-check`s Reihenfolge übernimmt, nicht dessen Lösungen.

**Lerneintrag — Form: geschärfte Regel.** *Eine Baseline-Regel mit zwei Trägern — Wellen-Fall und
Wellenlos-Fall — wird an beiden gelesen, bevor „optional" behauptet wird; die Freiheits-Klausel des
einen Trägers gilt nicht automatisch für den anderen.* slice-135 §4.2 und slice-139 §5 lasen nur
den Wellen-Fall („Kein Zwang zum Nachrüsten") und schlossen daraus „kein Trigger, weil keine offene
Welle" — ein Fehlschluss: Die Wellenlos-Träger-Tabelle (§1 hier) formuliert **denselben Schritt**
für Slice-Closures ohne die Freiheits-Klausel. Der Fehler fiel nicht bei der eigenen Prüfung auf,
sondern beim Maintainer, der auf `d-check`s bereits umgesetzten Stand verwies — derselbe Musterfall
wie bei `.claude/rules/` in [slice-142](../done/slice-142-claude-rules-symlinks-repariert.md): eine
Lücke, die keine Suchmethode und kein Gate von sich aus fängt, weil sie eine **Lese**-Lücke ist,
keine Mess-Lücke. *Weil* eine Regel mit zwei Trägern zwei Lesevorgänge braucht, und ein Delta,
das nur einen davon zitiert, den anderen unsichtbar macht, ohne dass ein Sensor das je bemerken
könnte — dieselbe Grenze wie beim `cr-text-reviewer`-Skill (`BEO-GATE/cr-text-behauptet-statt-gemessen`):
kein Sensor prüft, ob *alle* einschlägigen Stellen gelesen wurden, nur ob eine Behauptung zu einer
gelesenen Stelle passt.

**Zwei beobachtbare Closure-Kriterien:**

1. Die Zahlen in §2/§3 sind mit den genannten `grep`-/`find`-Befehlen gegen den Stand dieses Slice
   nachrechenbar.
2. `make gates` und `make verify` je **Exit 0** auf dem Stand dieses Slice, Ausgabe in Dateien,
   Exit-Codes getrennt geprüft.

**Offene Risiken und ihr Ausgang.**

- *Die eigene Fehleinordnung des Wellenlos-Falls in slice-135/139 ist unkorrigiert in `done/`
  stehen geblieben* — Ausgang: **gestrichen mit Begründung**: kein Risiko im Sinn der
  Dreier-Menge — `done/`-Prosa wird nicht rückwirkend korrigiert (dieselbe Disziplin wie überall
  in dieser Migration); die Korrektur steht hier, wo sie auffiel, nicht dort.
- *Die Sammel-Archiv-vs.-wellenlos-Entscheidung für 101 Slices steht aus* — Ausgang:
  **Folge-Slice** — Etappe D aus §6, braucht die Abnahme aus §5 zuerst.
- *Werkzeug-Wahl (Go vs. Bash) ungeklärt* — Ausgang: **Folge-Slice** — Etappe A aus §6.

**Folge-Slices:** keine vergeben. Vier Etappen A–D sind in §6 vorgeschlagen und brauchen die
Abnahme, bevor sie IDs bekommen.

## 9. Sub-Area-Modus

Berührt: **Planungs-Harness** (`docs/plan/planning/`), **Review-Harness** (`docs/reviews/`) —
Greenfield.
