# Review-Report: slice-173 — 2026-09-07

**Review-Art:** unabhängiger Lauf (anderes Kontextfenster als die Implementierung,
Modul 8 §Kontext-Trennung) — Code- und Vertrags-Review gegen Slice-Plan,
`AGENTS.md` §3/§4/§5 und die vendored Baseline.

**Gegenstand:** Commit-Range `c034b8a^..HEAD` — `c034b8a`
(`docs(planning)`, Lifecycle-`git mv`) und `c9ee5d8` (`feat(harness)`,
Umsetzung).

**Skill:** `.harness/skills/reviewer.md` @ Stand `92e1f64` (unverändert seit
Etappe A) · <!-- d-check:ignore -->
**Modell:** claude-opus-5 (1M) · **Datum:** 2026-09-07

**Eingangs-Kontext:**

- `docs/plan/planning/done/slice-173-versions-sensor-baseline-pins.md`
- `AGENTS.md` §3 (Hard Rules), §4 (Quality Gates), §5 (Dokumentations-Regeln)
- `harness/README.md` §Sensors, `harness/conventions.md` §Baseline und
  §Aktive Adaptionen
- `.d-check.yml`, `Makefile`, `tools/symlink-check.sh`,
  `.claude/hooks/pretooluse-command-guard.sh`
- Beobachtungs-Register: `BEO-GATE/symlink-ziel-nach-baseline-bump-ungeprueft`
  (`observation.md`, `state.md`, alle drei Belege) und
  `BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf`
- Baseline `modul-05` §Ziel-Form: Slice / §Zwei Schritte vor der
  Modus-Begründung, `modul-06` §Das Beobachtungs-Register
- Voriger Report `docs/reviews/2026-09-06-slice-167-etappe-a-vendoring-v620.md`

**Eigene Läufe** (alle Änderungen anschließend zurückgenommen, `git status`
leer — Nachweis am Ende):

| Lauf | Ergebnis |
|---|---|
| `make doc-check` (Ist-Stand) | Exit 0, 456 Dateien, 0 Befunde |
| `make doc-check` ohne `exempt-paths` | Exit 2, **19** `version-stale` in **4** Pfad-Klassen |
| `make doc-check` mit allen `v6.2.0`-Pins auf `v9.9.9` (außer `harness/conventions.md`) | **11** Dateien mit `version-stale` ⇒ effektive Prüfmenge |
| `make doc-check` mit einem `v9.9.9`-Pin **in** `harness/conventions.md` | nur `target-missing`, **kein** `version-stale` |
| `make symlink-check` | Exit 0, 7 getrackte Symlinks |
| `tools/symlink-check.sh` gegen sechs Fixture-Fälle in einem Wegwerf-Repo außerhalb des Baums | zwei Fehlverhalten, siehe F-3 |
| `make gates` | Exit 0 |
| `make verify` | Exit 0 |
| `make commit-scope-check RANGE=c034b8a^..HEAD` | Exit 0, 1 `(planning)`-Commit geprüft |

---

## Findings

### F-1 — Die tragende Zahl der Closure („Prüfmenge 15 lebende Dateien, darunter alle fünf aktiven MR-Einträge") reproduziert nicht

- `kategorie`: HIGH
- `quelle`: `AGENTS.md` §5 (Closure-/Lerneintrag-Disziplin), Reviewer-Skill
  §Klassifikation („nachweislich falsche Tatsachenbehauptung, gegen ein
  Repo-Artefakt verifiziert"); `harness/conventions.md` §Aktive Adaptionen
- `pfad`: `docs/plan/planning/done/slice-173-versions-sensor-baseline-pins.md`
  §2.1 (Zeilen 88–90) sowie Commit-Message `c9ee5d8`, Absatz 1
- `befund`: Der Satz „Gemessene Wirkung: 18 → **0** Befunde, Prüfmenge danach
  **15** lebende Dateien (darunter alle fünf aktiven `MR`-Einträge …)" nennt
  zwei Werte, die eine Nachmessung nicht bestätigt. **(a)** Die effektive
  Prüfmenge sind **11** Dateien: `AGENTS.md`, `harness/README.md`,
  `docs/plan/carveouts/README.md`, `docs/plan/planning/README.md`,
  `.harness/skills/reviewer.md`, die fünf aktiven `MR`-Dateien mit Pin und der
  Slice-Plan selbst (gemessen, indem alle `v6.2.0`-Pins auf einen
  Phantom-Stand gesetzt wurden — genau diese elf melden `version-stale`).
  `harness/conventions.md` — mit **13** Pins die pin-reichste Datei des Repos —
  ist nicht darunter (siehe F-4). **(b)** Die aktiven Adaptionen sind
  **sieben** (`MR-011`, `MR-012`, `MR-014`, `MR-015`, `MR-016`, `MR-019`,
  `MR-020`), nicht fünf; `MR-019` und `MR-020` tragen keinen Baseline-Pin und
  sind daher nie in der Prüfmenge. Die Zahl 15 ist per Hand aus
  „35 Dateien mit Pin minus drei Ausnahme-Klassen" abgeleitet, nicht gemessen —
  sie kennt weder die beiden später ergänzten Register-Klassen noch den
  Selbst-Ausschluss aus F-4. Der Slice erklärt sein Verfahren ausdrücklich zur
  Ablesung („nicht argumentiert, sondern abgelesen"); an dieser Stelle ist es
  eine Schätzung im Gewand einer Messung.
- `verifizierbar`: ja — `.d-check.yml`-Muster gegen einen Phantom-Stand fahren
  und die meldenden Dateien zählen; für (b) die Tabelle §Aktive Adaptionen.

### F-2 — `state.md` steht auf *verkörpert*, obwohl der neue Sensor die zweite der drei gezählten Instanzen nicht fängt

- `kategorie`: HIGH
- `quelle`: Baseline `modul-06` §Das Beobachtungs-Register (Ausgang
  *verkörpert* = „die Regel steht"); `AGENTS.md` §4-Disziplin, dass eine
  Sensor-Zeile ihre Grenze benennt
- `pfad`: `docs/plan/planning/observations/BEO-GATE/symlink-ziel-nach-baseline-bump-ungeprueft/state.md`
  Zeile 1–2; `AGENTS.md` §4, Zeile `make symlink-check`; Slice §8
  Steering-Loop-Eintrag
- `befund`: Der Eintrag zählt drei Instanzen, und der Ausgang *verkörpert* wird
  auf alle drei gestützt. `make symlink-check` prüft ausschließlich, ob das
  Ziel **existiert** (`[ -e "$l" ]`). Die zweite Instanz war aber genau der
  Fall, in dem es existierte: `evidence/slice-167.md` hält fest, dass die vier
  `.claude/rules/`-Symlinks beim Vendoring von `v6.2.0` „erneut auf `v6.0.0`
  zeigen" blieben — und `slice-167` §3 (Zeile 92) hält fest, dass `v6.0.0`
  „**vendored liegen**" blieb. Die Symlinks lösten also auf; `make
  symlink-check` hätte Exit 0 gemeldet. Damit deckt der Sensor 2 von 3
  gezählten Instanzen, und zwar nicht die, die zuletzt eintrat. Dieser Fall ist
  auch nicht randständig: `harness/conventions.md` §Baseline erlaubt zwei
  parallel vendored Stände ausdrücklich „während einer Migration" — also genau
  im Zeitfenster, in dem der Beobachtungs-Titel („nach Baseline-Bump") greift.
  `state.md` benennt als offene Grenze „ob das Ziel das **richtige** ist" und
  liest sich damit wie ein Absichts-Urteil; dass darunter eine der drei
  gezählten Instanzen fällt, steht weder dort noch in `AGENTS.md` §4 noch in
  der Closure-Notiz.
- `verifizierbar`: ja — `evidence/slice-167.md` und `slice-167` §3 gegen die
  Prädikat-Semantik des Skripts halten; reproduzierbar mit zwei vendored
  Ständen und einem Symlink auf den alten.

### F-3 — `tools/symlink-check.sh` meldet einen kaputten Symlink als heil (fail-open) und einen heilen als kaputt (Falschalarm); der Selbsttest sieht diesen Codepfad nicht

- `kategorie`: MEDIUM
- `quelle`: `AGENTS.md` §4 (Zusage „Jeder **getrackte** Symlink löst auf");
  Slice §3, Begründungsspalte („**eine** Prüf-Funktion für Gate-Lauf und
  Selbsttest, damit der Test den Codepfad prüft, der im Gate läuft")
- `pfad`: `tools/symlink-check.sh`, Zeile mit
  `git ls-files -s | awk '$1=="120000"{ $1=""; $2=""; $3=""; sub(/^ +/,""); print }'`
- `befund`: Die Pfad-Extraktion baut `$0` über awk-Feldzuweisung neu und
  verliert dabei Information. Zwei Fälle, beide in einem Wegwerf-Repo außerhalb
  des Baums reproduziert: **(a) fail-open** — ein getrackter, ins Leere zeigender
  Symlink namens `a  b.md` (zwei Leerzeichen) wird auf `a b.md` kollabiert;
  existiert diese Datei, meldet das Gate `symlink-check ok: 6 getrackte
  Symlink(s) loesen auf` und Exit 0, während einer von ihnen nicht auflöst.
  **(b) Falschalarm** — git C-quotet Nicht-ASCII-Namen (`"uml\303\244ut.md"`);
  das Skript prüft den gequoteten String, `-e` schlägt fehl, und `make gates`
  wird rot mit der Meldung „Symlink zeigt ins Leere" für einen Symlink, der
  auflöst. Ein einzelnes Leerzeichen im Namen funktioniert korrekt. Der
  Selbsttest kann keinen der beiden Fälle sehen: er reicht zwei selbstgebaute
  Pfade direkt an `kaputte()`, während der Defekt in der Extraktion **vor**
  dieser Funktion liegt — die Design-Begründung „der Test prüft den Codepfad,
  der im Gate läuft" gilt für das Prädikat, nicht für das Gate. Im Repo tritt
  heute keiner der beiden Fälle auf (7 Symlinks, alle ASCII, ohne Leerzeichen);
  die Zusage in `AGENTS.md` §4 ist trotzdem unqualifiziert formuliert, und der
  Skript-Kopf deklariert nur eine andere Grenze („ob das Ziel das richtige ist").
- `verifizierbar`: ja — Fixture-Repo mit `a b.md` (heil) und `a  b.md`
  (kaputt) sowie einem Umlaut-Symlink; das Skript darüber laufen lassen.

### F-4 — `harness/conventions.md` selbst wird vom neuen Muster nicht geprüft; die Grenze ist nirgends deklariert

- `kategorie`: MEDIUM
- `quelle`: `AC-QA-02` (ehrliche Grenze) als Haltung der §4-Tabelle — jede
  andere Sensor-Zeile dort nennt ihr „**Nicht** geprüft"; `harness/conventions.md`
  §Baseline
- `pfad`: `.d-check.yml` §`versions`, `patterns[1]`
  (`current-from: harness/conventions.md#baseline`); Deklaration in `AGENTS.md`
  §4 Zeile `make doc-check` und `harness/README.md` §Sensors
- `befund`: Gemessen: ein auf einen Phantom-Stand gesetzter Pin **in**
  `harness/conventions.md` (Zeile 111, `Ersetzt-Baseline-Regel` von `MR-011` in
  der Tabelle §Aktive Adaptionen) erzeugt `target-missing`, aber **kein**
  `version-stale` — d-check nimmt die Datei, die den `current-from`-Span
  trägt, offenbar ganz aus der Prüfmenge. Betroffen sind 13 Baseline-Pins,
  mehr als in jeder anderen Datei des Repos, darunter die
  `Ersetzt-Baseline-Regel`-Zeiger der Adaptions-Tabelle — also genau die Klasse,
  für die §Baseline zusagt: „Sein Zeiger wandert darum beim Baseline-Wechsel
  mit". Beide Deklarationen (`AGENTS.md` §4, `harness/README.md`) beschreiben
  die Ausnahmen als die fünf Zeitdokument-Klassen und lassen den Leser
  schließen, alles übrige sei erfasst.
- `verifizierbar`: ja — einen Pin in `harness/conventions.md` außerhalb
  §Baseline verfälschen und `make doc-check` fahren.

### F-5 — `exempt-paths` deckt zwei Dateien mit *migrierendem* Zeiger mit ab

- `kategorie`: MEDIUM
- `quelle`: Begründung des Slice selbst („dort ist der GENANNTE Stand wahr,
  weil damals gegen ihn gemessen wurde"); `harness/conventions.md` §Baseline
  („Die Zusage gilt **ohne Ausnahme**")
- `pfad`: `docs/reviews/README.md` Zeile 7;
  `harness/conventions/done/MR-018-review-pflicht-v610-wortlaut.md` Zeile 33
- `befund`: Die Ausnahme ist verzeichnis-, nicht rollenweise geschnitten, und
  zwei der abgedeckten Dateien sind keine Lauf-Belege. **(a)**
  `docs/reviews/README.md` ist das Konventions-Dokument der Review-Ablage und
  zeigt auf die **aktuelle** Report-Vorlage
  (`.harness/baseline/v6.2.0/templates/docs/reviews/review-report.template.md`);
  es fällt unter `docs/reviews/**`. **(b)**
  `harness/conventions/done/MR-018-…` verlinkt in Zeile 33 dieselbe aktuelle
  Vorlage — dieser Zeiger wurde in slice-172 nachweislich **nachgezogen**, ist
  also gerade kein eingefrorenes Zitat; er fällt unter
  `harness/conventions/done/**`. Gemessen: beide auf einen Phantom-Stand
  gesetzt ⇒ **kein** `version-stale`. Nach dem nächsten Bump meldet zu beiden
  nichts, solange der alte Stand noch vendored liegt. Die vier reinen
  Prosa-Zitate in `conventions/done/` (`MR-006`, `MR-007`, `MR-013`, `MR-017`)
  sind dagegen zu Recht ausgenommen — sie sind Backtick-Text, keine Links, und
  nennen wahr, was damals galt.
- `verifizierbar`: ja — beide Zeilen auf einen Phantom-Stand setzen und
  `make doc-check` fahren.

### F-6 — „Dass die Klassenliste unvollständig war, hat kein Argument gezeigt, sondern der Lauf" — für `observation.md` hat kein Lauf etwas gezeigt

- `kategorie`: MEDIUM
- `quelle`: `AGENTS.md` §5, `BEO-GATE/cr-text-behauptet-statt-gemessen`-Disziplin
  („hast du *das* gemessen, worüber du redest?")
- `pfad`: Slice §2.1, letzter Absatz (Zeilen 92–99)
- `befund`: Von den beiden nachgetragenen Register-Klassen hat nur eine einen
  Befund erzeugt: `evidence/slice-173.md` der Symlink-Beobachtung (der 19.
  Befund meines Nachlaufs). Die beiden `observation.md`-Dateien mit
  Baseline-Pin — `BEO-HARNESS/mr-aufloesungs-trigger-ohne-waechter` und
  `BEO-PLAN/archiv-stub-feld-ohne-leser` — pinnen `v6.2.0` und melden nichts;
  ihre Ausnahme ist aus der Drei-Lebensdauern-Regel von `modul-06`
  **argumentiert**, nicht abgelesen. Nebenbefund derselben Stelle: die
  bewusst feine Trennung „`state.md` bleibt geprüft" hat heute **null**
  Gegenstände — keine der 49 `state.md`-Dateien trägt einen Baseline-Pin.
- `verifizierbar`: ja — `make doc-check` ohne `exempt-paths` und die
  meldenden Pfade lesen.

### F-7 — Die Widerlegung der slice-172-„Spec-Lücke" trägt die halbe Strecke

- `kategorie`: MEDIUM
- `quelle`: `harness/conventions.md` §Baseline; Adaptions-Block-Disziplin
  (akzeptierte Einträge werden nicht inhaltlich geändert); `.d-check.yml`
  §`ignore-refs`-Kommentar („Ein Review-Report ist selbst ein LAUF-BELEG … und
  wird nicht mehr editiert")
- `pfad`: Slice §2.1 (Zeilen 76–86) und §8, Absatz „Was ging anders als geplant"
- `befund`: Der Slice erklärt eine Aussage des Vorgängers für falsch und
  begründet das mit einer Zuständigkeits-Teilung: `versions` für lebende
  Dokumente, die Link-Prüfung fürs Auflösen. Der erste Teil trägt — die beiden
  Module beantworten unterschiedliche Fragen, und das ist belegt. Der zweite
  Teil beantwortet nicht, was slice-172 behauptet hatte: dass die Ausnahme das
  *Liegenbleiben* alter Stände voraussetzt. Genau das bleibt so. Verschwindet
  ein Stand, meldet die Link-Prüfung — in Dateien, die dieses Repo als
  unveränderlich führt (Review-Reports, akzeptierte `MR`-Einträge, archivierte
  Slices), und der einzige Ausweg ist, sie doch zu editieren; slice-172 hat
  genau das getan. Die Aussage „Zwei Sensoren, zwei Fragen" beschreibt die
  Aufteilung korrekt und lässt die Kollision aus, die die Frage erst erzeugt
  hat. Verschärfend: `harness/conventions.md` §Baseline hat die Frage
  inzwischen anders entschieden („Die Zusage gilt **ohne Ausnahme**, auch für
  das Feld `Ersetzt-Baseline-Regel` akzeptierter `MR`-Einträge … Sein Zeiger
  wandert darum beim Baseline-Wechsel mit") — die Ausnahme
  `harness/conventions/done/**` nimmt dieser Zusage für aufgelöste Einträge den
  Wächter, ohne die Divergenz zu benennen.
- `verifizierbar`: nein — Urteil über die Reichweite einer Begründung; die
  beiden zitierten Regelstellen sind aber nachlesbar.

### F-8 — Plan §3 und DoD-Punkt 2 sagen „drei Zeitdokument-Klassen", ausgeliefert sind fünf

- `kategorie`: LOW
- `quelle`: `AGENTS.md` §5 (DoD ist der Verifikations-Gegenstand)
- `pfad`: Slice §3, Zeile zu `.d-check.yml`; Slice §4, zweiter DoD-Punkt
- `befund`: Beide Stellen nennen drei Klassen; `.d-check.yml` und die
  Deklarationen in `AGENTS.md` §4 / `harness/README.md` nennen fünf. §2.1 trägt
  die Ergänzung, §3 und §4 sind nicht nachgezogen — ein abgehakter DoD-Punkt
  beschreibt damit einen anderen Stand als das gelieferte Artefakt.
- `verifizierbar`: ja — Textvergleich `.d-check.yml` gegen Slice §3/§4.

### F-9 — ADRs sind weder ausgenommen noch änderbar

- `kategorie`: LOW
- `quelle`: `AGENTS.md` §3.5 (ADR-Immutabilität); `harness/conventions.md`
  §Anforderungs-Anlege-Prozess, wo derselbe Widerspruch für Anker beschrieben
  und über einen expliziten Anker aufgelöst wird
- `pfad`: `.d-check.yml` §`versions`, `patterns[1].exempt-paths`
- `befund`: `docs/plan/adr/**` steht nicht in der Ausnahmeliste. Trägt künftig
  eine `Accepted`-ADR einen `.harness/baseline/<tag>/`-Pfad, verlangt
  `make doc-check` nach dem nächsten Bump einen Nachzug, den `make doc-immutable`
  verbietet — dieselbe Klemme, die §Anforderungs-Anlege-Prozess für
  Überschriften-Anker beschreibt. Heute gegenstandslos: keine der ADRs trägt
  einen Baseline-Pfad (geprüft).
- `verifizierbar`: ja — `grep` über `docs/plan/adr/` heute leer; die Klemme
  entsteht erst mit dem ersten Treffer.

### F-10 — Was `make doc-targets` an dieser Änderung *nicht* erzwingt

- `kategorie`: INFO
- `quelle`: `AGENTS.md` §4, Zeile `make doc-targets`
- `pfad`: `Makefile`, `harness/README.md` §Sensors
- `befund`: Erzwungen ist die Deklarations-Konsistenz in beide Richtungen
  (dokumentiertes Target existiert / reales Gate-Target ist in §4 gelistet).
  **Nicht** erzwungen sind drei Dinge, die dieser Slice mitgeliefert hat und
  die nur durch Sorgfalt stimmen: die Aufnahme ins `gates`-Aggregat, der
  `.PHONY`-Eintrag und die `GATES`-Liste des Command-Guard (letztere fing der
  `guard-selftest`, wie die Commit-Message berichtet — geprüft und
  zutreffend). Ebenfalls ungewächtert ist die *Aufzählung* der
  Aggregat-Mitglieder in `harness/README.md`; sie war vor diesem Slice
  veraltet (`doc-reviews`, `ci-range-selftest` fehlten) und ist hier nebenbei
  vollständig nachgezogen worden — Prosa, die niemand prüft, aber jetzt
  korrekt.
- `verifizierbar`: ja — Prerequisites von `gates` gegen die Aufzählung halten.

### F-11 — 3×-Übertritt in-slice verkörpert statt als Folge-Slice

- `kategorie`: INFO
- `quelle`: Baseline `modul-05` §Zwei Schritte vor der Modus-Begründung
  („erreicht der Eintrag mit diesem Slice 3×, … braucht einen **eigenen
  Folge-Slice**") gegen `modul-06` §Das Beobachtungs-Register (Lese-Schritt
  bei der Slice-Closure, Ausgang *verkörpert*)
- `pfad`: Slice §2.2, letzter Absatz; §9 Sichtungs-Tabelle
- `befund`: Die beiden Regelstellen weisen in unterschiedliche Richtungen; der
  Slice wählt die Verkörperung im selben Lauf und begründet sie (gleiche
  Schicht, gleicher Gegenstand, Größen-Regel gehalten). Das ist vertretbar und
  ausdrücklich benannt, aber es ist eine Auslegung — als `MR` deklariert ist
  sie nicht. Zur Größen-Regel selbst kein Befund: drei Liefer-Punkte, eine
  berührte Schicht, `AGENTS.md` §5 hält.
- `verifizierbar`: nein — Auslegungsfrage zwischen zwei Baseline-Abschnitten.

### F-12 — `make trace-check` läuft in dieser Umgebung auf keiner Range

- `kategorie`: INFO
- `quelle`: `AGENTS.md` §5 (Traceability-Pflicht)
- `pfad`: `Makefile` Zeile 208 (`trace-check`)
- `befund`: Sowohl `RANGE=c034b8a^..HEAD` als auch die Default-Range
  `HEAD~1..HEAD` brechen mit *„Range-Basis-Vorfahren nicht lesbar: object not
  found"* ab (make-Exit 2). Das Repo ist nicht shallow (739 Commits); der
  Ausfall ist **nicht** von diesem Slice verursacht und trifft jede Range
  gleichermaßen. Die inhaltliche Zusage ist erfüllt: beide Commit-Messages
  nennen `slice-173` (von Hand geprüft), und `make commit-scope-check` über
  dieselbe Range ist grün.
- `verifizierbar`: ja — `make trace-check` ohne Argumente reproduziert den
  Abbruch.

---

## Negativbefunde (geprüft, ohne Befund)

- **Mutations-Probe des neuen Musters, beide Richtungen.** Ein verfälschter Pin
  in `AGENTS.md` Zeile 31 erzeugt `version-stale` mit „erwartet v6.2.0"; die
  Rücknahme führt auf 0 Befunde, Exit 0. Die Behauptung der Commit-Message
  reproduziert (die dort genannte „Exit 2" ist der make-Exit; d-check selbst
  endet mit 1 — `AGENTS.md` §4 hält diese Nichtunterscheidbarkeit bereits fest).
- **Wirkung der `exempt-paths` auf die Zeitdokumente.** Ohne Ausnahmen 19
  Befunde, mit Ausnahmen 0 — die Ausnahme greift, ohne Nebenwirkung auf die
  gemeldeten Klassen. Die Glob-Semantik stimmt mit der Absicht überein:
  `docs/reviews/**` fasst auch Dateien direkt in `docs/reviews/`,
  `observations/**/observation.md` und `observations/**/evidence/**` treffen
  genau die 49 `observation.md` und 78 Beleg-Dateien, die 49 `state.md` bleiben
  in der Prüfmenge.
- **Verteilung der 18 gegen 19 Befunde.** Die Differenz zu §2.1 ist genau
  `evidence/slice-173.md` dieses Slice, das beim Erst-Lauf noch nicht existierte;
  die Klassen `done/` (8), `docs/reviews/` (6) und `conventions/done/` (4)
  stimmen exakt mit der Tabelle überein, ebenso die beanstandeten Stände
  `v3.5.2`/`v5.12.0`/`v6.0.0`/`v6.1.0`. Die Kernaussage „kein Befund in einer
  lebenden Datei" trägt auch heute — alle 19 liegen in Zeitdokumenten.
- **`tools/symlink-check.sh`, funktionale Fälle.** Kaputter Symlink ⇒ Exit 1 mit
  Pfad-Ausgabe; Kette Symlink→Symlink→fehlend ⇒ beide Glieder gemeldet; Symlink
  auf ein Verzeichnis ⇒ kein Befund; Lauf außerhalb eines git-Repos ⇒ Exit 128,
  fail-closed; invertiertes Prädikat ⇒ Selbsttest bricht mit Exit 1 ab, bevor
  irgendetwas geprüft wird. `set -euo pipefail` und die Zuweisung aus der
  Pipeline machen einen `git`-Ausfall fail-closed.
- **Symlink-Zahl und Geltungsbereich.** 7 getrackte Symlinks, wie das Target
  meldet; die vier Regelwerk-Verweise plus drei auf eigene Repo-Dateien. Die
  Formulierung „die vier Verweise … auf einzelne Regelwerk-Module" in
  `AGENTS.md` §4 ist korrekt qualifiziert.
- **Deklarations-Vollständigkeit.** `symlink-check` steht in `AGENTS.md` §4, in
  `harness/README.md` §Sensors, in `.PHONY`, in den Prerequisites von `gates`
  und in der `GATES`-Menge des Command-Guard. `make doc-targets` und
  `make gate-consistency` laufen grün.
- **Lifecycle-Commit `c034b8a`.** Reiner `git mv` mit `R100`; die beiden
  mitgeänderten Dateien sind der Ruhe-Marker der Roadmap und ein
  Verweis-Nachzug in einem `state.md`, beide durch `make slice-mv` erzeugt und
  in der Message benannt. Ausschließlich `docs/plan/planning/` berührt —
  `make commit-scope-check` grün. `AGENTS.md` §3.3 gewahrt.
- **Hard Rules.** §3.1: das neue Skript ruft nur `bash`, `git`, `awk`, `sed`,
  `grep`, `ln`, `mktemp` — keine Host-Toolchain, kein Paketmanager. §3.2: keine
  Suppression (`make suppression-check` grün). §3.5: keine ADR berührt. §3.6:
  keine Schwelle gesenkt — die Änderung fügt Prüfung hinzu, sie nimmt keine weg.
- **Slice-Form.** Kopffelder vollständig (`Welle:`, `Bezug:`, `Berührte
  Spec-Stellen:`, `Verantwortlich:`, `Autor:`), Lerneintrag in einer der drei
  benannten Formen („neuer Sensor"), beide Risiken aus §7 mit genau einem
  Ausgang aus der geschlossenen Menge, drei Paarungen benannt. `make verify`
  (inkl. `doc-structure`, `verify-risiko-ausgaenge`, `verify-observations`,
  `doc-complete`) Exit 0.
- **Keine erfundene Kennung.** `slice-169` liegt in `open/`, beide zitierten
  `BEO-GATE`-Verzeichnisse existieren mit nicht leerem `evidence/`, die Belege
  `slice-142`/`slice-167` sind vorhanden, `MR-018` und `MR-006` lösen auf.
- **Aggregat-Läufe.** `make gates` Exit 0, `make verify` Exit 0 — beide von mir
  selbst gefahren, nicht aus der Commit-Message übernommen.
- **Repo-Zustand nach dem Review.** Alle Mess-Eingriffe (`.d-check.yml` ohne
  `exempt-paths`, 28 Dateien mit Phantom-Pin, ein Pin in
  `harness/conventions.md`) wurden mit `git checkout --` zurückgenommen;
  `git status --porcelain` ist danach leer. Einziger Zuwachs im Arbeitsbaum ist
  dieser Report.

---

## Kategorie-Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 2 |
| MEDIUM | 5 |
| LOW | 2 |
| INFO | 3 |

## Verdikt

**Die Mechanik trägt, die Begründungstexte nicht durchgängig.** Beide Sensoren
tun, was ihre Kern-Zusage sagt: das zweite `versions`-Muster fängt einen
veralteten Baseline-Pin in einem lebenden Dokument nachweislich, und
`make symlink-check` fängt einen ins Leere zeigenden getrackten Symlink
nachweislich. Die `exempt-paths`-Entscheidung ist im Grundsatz richtig
begründet und im Bestand rauschfrei.

Blockierend sind zwei Aussagen, die eine Nachmessung nicht bestätigt: die
Prüfmengen-Zahl samt „fünf aktiven `MR`-Einträgen" (F-1) und der Ausgang
*verkörpert* für eine Beobachtung, deren zweite gezählte Instanz der neue
Sensor nicht fängt (F-2). Beide sind Text, nicht Code — aber beide werden
beim nächsten Baseline-Sprung gelesen und geglaubt, und F-2 schließt einen
Steering-Loop-Eintrag, der noch nicht ganz geschlossen ist.

Von den MEDIUMs verdienen zwei besondere Aufmerksamkeit, weil sie die
Reichweite des neuen Wächters betreffen und beide gemessen sind: die
pin-reichste Datei des Repos steht außerhalb seiner Prüfmenge (F-4), und zwei
lebende Zeiger fallen unter eine Verzeichnis-Ausnahme (F-5). Zusammen mit F-1
heißt das: der Sensor prüft weniger, als seine Deklaration nahelegt — und
seine Grenze ist, anders als bei jeder anderen Zeile der §4-Tabelle, nicht
ausgeschrieben.
