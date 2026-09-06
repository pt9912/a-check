# Review-Report: slice-168 — 2026-09-06

**Review-Art:** Plan + Code — geprüft gegen den Slice-Plan und gegen das
neue Werkzeug selbst (Modul 10 §Drei Review-Arten): die Kern-Behauptung
ist eine **Deckungs-Zusage** („der Prüfer kann nicht mehr still ohne
Gegenstand grün melden"), und die ist nur durch eigenes Ausführen und
Mutieren prüfbar, nicht durch Lesen.

**Gegenstand:** Commit `805ce63` (`tools/dcheck-phrase-selftest.sh`,
`Makefile`, `.claude/hooks/pretooluse-command-guard.sh`, `AGENTS.md` §4,
`harness/README.md` §Sensors) plus der Slice-Plan.

**Skill:** `.harness/skills/reviewer.md` @ Stand `5649d84` (unverändert
seit Anlage) · <!-- d-check:ignore -->
**Modell:** claude-sonnet-5, **getrennter Kontext** (eigener Subagent,
kein `fork` — er hat den Implementierungs-Kontext nicht geerbt) ·
**Datum:** 2026-09-06

**Eingangs-Kontext:**

- `docs/plan/planning/in-progress/slice-168-pruefer-kalibrierungs-selbsttest.md`
- `tools/dcheck-phrase-selftest.sh`, `tools/verify-risiko-ausgaenge.sh`
  (Vorbild-Muster), `Makefile`, `.d-check.yml`
- `docs/plan/planning/observations/BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf/`
- `AGENTS.md` §4/§5, `harness/README.md` §Sensors

**Eigene Läufe:** Selbsttest im Original (Exit 0) · **vier** Mutationen
auf Kopien, je eine pro Kontrolle · Bruch der **echten** `.d-check.yml` ·
`make gates` (Exit 0), `make verify` (Exit 0), `git status` danach leer.
Das ausgelieferte Skript ist unverändert.

---

## Findings

### F-1 — Der Sensor prüft die Werkzeug-Seite, nicht die Korpus-Seite

- `kategorie`: HIGH
- `quelle`: eigener Lauf; Gegenprobe am realen Korpus
- `pfad`: `tools/dcheck-phrase-selftest.sh` (Muster 1),
  `docs/plan/planning/observations/BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf/state.md`
- `befund`: Die Kontrollen fragen, ob **d-check** auf die Phrase reagiert.
  Ausgefallen ist zweimal die **Gegenrichtung**: a-checks eigener Korpus traf
  das Muster nicht mehr (slice-120: der gesuchte Wortlaut kam null Mal vor;
  slice-165: die real verwendete DoD-Formulierung traf die Trigger-Phrase
  nicht). Das Skript liest den echten Korpus nie an — es hardcodet die
  richtige Fixture-Zeile. Ein Autor, der morgen wieder eine nicht-auslösende
  Wortform schreibt, bleibt unbemerkt, und der neue Sensor bleibt grün.
  `state.md`s Zusage „deckt die beiden real aufgetretenen Fälle ab" ist
  dadurch zu stark: gedeckt ist die Werkzeug-Seite beider Fälle, nicht die
  Korpus-Seite. Fehlt: eine Nichtleerheits-/Erwartungszahl-Kontrolle der
  **echten** Kandidatenmenge — Vorbild ist `exempt-expect-count` im
  `structure`-Block der `.d-check.yml`.
- `verifizierbar`: ja — reproduzierbar über eine Wortform-Änderung in einer
  `done/`-Slice-DoD bei anschließend grünem `make dcheck-phrase-selftest`.
- `klasse`: Sensor prüft die Richtung, in die noch nie jemand gefallen ist

**Nachgemessen (Korrektur am Finding):** Der Reviewer nannte slice-161,
-163 und -164 als heute unbewacht. Case-insensitiv nachgezählt trifft das
für **slice-161, -162 und -163** zu (Report vorhanden, Wortform löst
nicht aus); **slice-164 gehört zur Kandidatenmenge**. Die Menge ist heute
also **nicht leer** (slice-164, -165, -166, -167) — der „ohne
Gegenstand"-Ausfall ist **nicht live**, sondern jederzeit wieder
erreichbar. Der Kern des Findings bleibt damit bestehen, seine
Dringlichkeit ist geringer als vom Reviewer angesetzt.

### F-2 — `TASKS_IGNORE_PATTERN` ist eine ungekoppelte Kopie

- `kategorie`: MEDIUM
- `quelle`: eigener Lauf (echte `.d-check.yml` gebrochen)
- `pfad`: `tools/dcheck-phrase-selftest.sh:93` gegen `.d-check.yml:192`
- `befund`: Das Muster steht im Skript als Konstante und in der echten
  Konfiguration ein zweites Mal. Beide sind heute zeichengleich, aber nichts
  hält sie zusammen: Nach einem Bruch der **echten** Konfiguration blieb
  `make dcheck-phrase-selftest` bei **Exit 0**. Dass `doc-structure` den
  Bruch heute auffängt, hängt daran, dass zufällig zwei Slices genau vier
  DoD-Punkte führen — das ist korpus-abhängig, kein Schutz. Ein Sensor, der
  eine Kopie kalibriert statt des Originals, kalibriert nichts.
- `verifizierbar`: ja — Muster in `.d-check.yml` ändern, Selbsttest laufen
  lassen.
- `klasse`: Kalibrierung gegen eine Kopie statt gegen das Original

### F-3 — Ungedeckte phrasen-basierte Konfigurationen als „künftig" etikettiert

- `kategorie`: MEDIUM
- `quelle`: eigene Durchsicht der `.d-check.yml`
- `pfad`: Slice-Plan §2 (Antwort 1) und §3 („Nicht umgesetzt")
- `befund`: Der Plan schreibt die verbleibende Lücke künftigen Mustern zu.
  Es gibt sie aber **heute**: `versions.pin-pattern` (Prüfmenge genau eine
  Fundstelle — eine Umformulierung macht den Prüfer stumm),
  `structure`-Verbotsmuster, `vcs.immutable-when`, `commits.exempt-pattern`,
  `matrix.exclude-sections`. Die Lücke ist verschoben, nicht benannt.
- `verifizierbar`: ja — Sichtprüfung der Konfiguration.
- `klasse`: Bestehende Lücke als künftige deklariert

### F-4 — Dritter Risiko-Ausgang benannt, aber nicht genommen

- `kategorie`: MEDIUM
- `quelle`: Abgleich Slice §7 gegen Modul 5 §Offene Risiken
- `pfad`: Slice-Plan §7, zweites Risiko
- `befund`: Der Ausgang lautete „weiter offen → Beobachtungs-Register (kein
  neuer Eintrag jetzt)". Der dritte Ausgang **ist** das Wandern ins
  Register; ihn zu nennen und nicht zu vollziehen, ist keiner der drei.
  `tools/verify-risiko-ausgaenge.sh` matcht nur die Form und bleibt grün —
  der Sensor deckt das nicht auf.
- `verifizierbar`: ja.
- `klasse`: Ausgangs-Form erfüllt, Ausgangs-Substanz nicht

### F-5 — Statuswert `verkörpert` trägt nur zur Hälfte

- `kategorie`: LOW
- `quelle`: Folge aus F-1
- `pfad`: `…/BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf/state.md`
- `befund`: Für die „ohne Aufruf"-Hälfte ist `verkörpert` verifiziert
  (`doc-complete` liegt tatsächlich im `verify`-Aggregat, `Makefile:181`).
  Für die „ohne Gegenstand"-Hälfte gilt es nur für die Werkzeug-Seite.
- `verifizierbar`: ja.
- `klasse`: Statuswert stärker als der Beleg

### F-6 — DoD ohne Reserve

- `kategorie`: INFO
- `pfad`: Slice-Plan §4
- `befund`: Sieben Punkte, vier ignoriert, **drei von drei** zählend —
  konform, aber exakt an der Grenze. Jede Erweiterung dieses Slice kippt ihn
  über die Größen-Regel.

## Negativbefunde

- geprüft, ohne Befund: **vier** Mutationen (je eine pro Kontrolle) melden
  alle vier korrekt rot, mit richtig zugeordnetem Meldungstext — der Slice
  hatte nur *eine* Richtung von Hand geprüft, die Zusage trägt also stärker
  als im Plan behauptet.
- geprüft, ohne Befund: die Fixture-Muster sind gegenüber `.d-check.yml`
  heute zeichengleich (die fehlende **Kopplung** ist F-2, nicht die
  aktuelle Gleichheit).
- geprüft, ohne Befund: `doc-complete` liegt im `verify`-Aggregat — die
  „ohne Aufruf"-Argumentation des Plans (§2, Antwort 1) trägt.
- geprüft, ohne Befund: `make gates` und `make verify` beide Exit 0,
  Arbeitsbaum danach leer.
- geprüft, ohne Befund: Kopffelder (`Verantwortlich`, `Autor`, `Berührte
  Spec-Stellen`) vollständig, konsistent zu slice-166/slice-167.
- geprüft, ohne Befund: Closure-Notiz mit Lerneintrag-Form „neuer Sensor";
  Risiko 1 sauber „gestrichen mit Begründung".
- geprüft, ohne Befund: `.PHONY`, `GATES`-Liste des Command-Guard,
  `AGENTS.md` §4 und `harness/README.md` §Sensors konsistent
  (`doc-targets` grün).
- geprüft, ohne Befund: `evidence/` trägt genau die drei belegten Vorgänge;
  Ruhe-Marker der Roadmap korrekt entfernt (`doc-planning` grün);
  `state.md`-Link auflösbar.

## Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 1 |
| MEDIUM | 3 |
| LOW | 1 |
| INFO | 1 |

**Finding-Klassen dieses Laufs:** Sensor prüft die Richtung, in die noch
nie jemand gefallen ist · Kalibrierung gegen eine Kopie statt gegen das
Original · Bestehende Lücke als künftige deklariert · Ausgangs-Form
erfüllt, Ausgangs-Substanz nicht · Statuswert stärker als der Beleg

## Verdikt

**Merge-blockierend:** nein. Das ausgelieferte Werkzeug tut, was es tut,
und tut es nachweislich richtig (vier Mutationen, vier rote Meldungen).
Blockierend wären die Findings nur, wenn sie die Substanz träfen — sie
treffen die **Reichweite der Zusage**: F-1/F-5 überziehen sie, F-3 schiebt
eine bestehende Lücke in die Zukunft, F-4 nimmt einen Ausgang nicht.

**Eingearbeitet vor Abschluss:** F-1, F-3, F-4, F-5 als Text-Korrektur in
Slice und `state.md` (die Zusage wird auf das gemessene Maß
zurückgeschrieben, der Beobachtungs-Ausgang wechselt auf `geplant` mit
Kennung). F-1 und F-2 sind substanzielle Arbeit und damit ein
**Folge-Slice** — sie in slice-168 nachzuziehen hieße, ihn über die
Größen-Regel zu heben (F-6).

**Übergabe:** Dieser Report ist Lauf-Beleg, keine Verifikation —
DoD-/Spec-Konformität prüft der Verifier separat (Modul 11).
