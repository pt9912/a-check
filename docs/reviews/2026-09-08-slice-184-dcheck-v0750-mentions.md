# Review-Report: slice-184 — 2026-09-08

**Review-Art:** unabhängiger Lauf — der Reviewer hat weder Plan noch Umsetzung
geschrieben; jede zitierte Zahl ist in diesem Lauf neu gemessen, keine aus dem
Gegenstand übernommen (Modul 10 §Drei Review-Arten).

**Gegenstand:** Commits `a43258d` (`d-check`-Pin auf `v0.75.0`), `ec64373`
(`doc-mentions` angelegt) und `98548af` (Slice-Plan: Umsetzung,
Risiko-Ausgänge, Closure-Notiz). Mitgelesen, weil der Plan als Ganzes trägt:
`71939c5` (Schnitt), in dem §2 *Ausgangsmessung* entstand.

**Skill:** `.harness/skills/reviewer.md` @ Stand `3fae6d3` (unverändert seit
slice-176) · <!-- d-check:ignore -->
**Modell:** claude-opus-5[1m] · **Datum:** 2026-09-08

**Eingangs-Kontext:**

- `.harness/skills/reviewer.md`; `AGENTS.md` §3/§4/§5/§6
- slice-184 (Plan in `in-progress/`), slice-181 und slice-079 als zitierte Vorbilder
- `v6.5.0` · `templates/harness/README.template.md` (die zitierte Ziel-Form-Grenze)
- `v6.5.0` · `regelwerk/modul-13-quality-gates.md` §Vorhanden ≠ behauptet
- `harness/conventions.md` §Baseline; `docs/reviews/README.md`
- Eigene Fixtures gegen den Digest `sha256:18e9cd85…` (sieben Stück, siehe
  Negativbefunde) — das Modul `mentions` ist undokumentiert, also wurde jede
  Verhaltens-Aussage des Slice nachgefahren statt geglaubt.

---

## Findings

### F-1 — `doc-mentions` hängt in `gates`, wird von `AGENTS.md` §4 aber als *advisory* geführt

- `kategorie`: HIGH
- `quelle`: Hard Rule `AGENTS.md` §4 (*„Nur hier gelistete Targets existieren im
  Makefile … Mandatory ist, was in einem der beiden Aggregate hängt"*);
  `v6.5.0` · `regelwerk/modul-13-quality-gates.md` §Vorhanden ≠ behauptet
- `pfad`: `AGENTS.md`:149–154 gegen `Makefile`:207 und `harness/README.md`:93
- `befund`: Der Absatz zählt die mandatorischen `doc-*`-Targets abschließend auf
  — `doc-check`, `doc-targets`, `doc-planning`, `doc-workflows`, `doc-immutable`
  (in `gates`), `doc-structure`, `doc-complete` (in `verify`) — und erklärt
  „die übrigen" zu **advisory**; `doc-mentions` steht nicht darin, hängt aber
  seit `ec64373` im `gates`-Aggregat (`Makefile`:207), und `harness/README.md`:93
  sagt in derselben Zeile „im `gates`-Aggregat". Ein Agent, der den Absatz liest,
  hält ein mandatorisches Gate für eine Abrufoption. Kein Sensor deckt die
  Stelle: `doc-targets` prüft die beiden **Tabellen** (`doc-tables:
  [AGENTS.md, harness/README.md]`), nicht diese Prosa, und `guard-selftest`
  leitet seine Soll-Liste aus `Makefile`/`d-check.mk` ab.
- `verifizierbar`: nein — kein Gate-Lauf bestätigt es; belegt durch Vergleich
  der drei Stellen (siehe Verifikation unten).
- `klasse`: Neues Gate ins Aggregat gehängt, die Mandatory-Deklaration nicht nachgezogen

**Adversarische Verifikation (HIGH-Pflicht).** Drei Handgriffe, alle auf dem
Stand `98548af`:

1. `sed -n '207p' Makefile` → das `gates`-Ziel führt `doc-check doc-targets
   doc-planning doc-workflows doc-reviews doc-mentions`. Sechs `doc-*` hängen im
   Aggregat.
2. `sed -n '149,155p' AGENTS.md` → der Absatz nennt fünf davon nicht:
   `doc-reviews` und `doc-mentions` fehlen, und `doc-immutable` steht dort als
   „in `gates`", obwohl es weder in `gates` (Zeile 207) noch in `verify`
   (Zeilen 192–201) hängt.
3. Gegenprobe, dass der Absatz erschöpfend gemeint ist: der Halbsatz „die
   übrigen sind **advisory** — `d-check`-Funktionen, die man aufruft, wenn man
   sie braucht" lässt keine dritte Klasse offen.

Damit ist die Aussage nicht nur unvollständig, sondern gegen `Makefile`:207
falsch. **Der Fall ist nicht neu:** dieselbe Auslassung entstand mit slice-160
für `doc-reviews` (`git log -S` zeigt: die Prosa-Zeile wurde zuletzt von
slice-130 angefasst, das Aggregat seither zweimal erweitert) — zweites
Vorkommen derselben Klasse.

### F-2 — „Jedes Target bekommt zusätzlich `--disable mentions`" — gemessen 6 von 13

- `kategorie`: MEDIUM
- `quelle`: eigener Vergleich `--print-mk` beider Digests gegen `d-check.mk`
- `pfad`: `docs/plan/planning/in-progress/slice-184-dcheck-v0750-mentions.md`:48–49
  (Satz aus `71939c5`, im Gegenstand unverändert weitergetragen)
- `befund`: Das Fragment definiert 13 Targets; `--disable mentions` steht in
  6 Rezepten (`doc-immutable`, `doc-commits`, `doc-planning`, `doc-tracked`,
  `doc-targets`, `doc-structure`) und fehlt in den sechs übrigen
  `docker run`-Rezepten (`doc-check`, `doc-trace`, `doc-complete`,
  `doc-doctor`, `doc-repair`, `doc-usage`). Die daraus gezogene Folgerung
  („Das Modul ist strikt opt-in") trifft dennoch zu, trägt aber auf einem
  anderen Beleg: `modules:` in `.d-check.yml`:21 führt `mentions` nicht, und
  der Werkzeug-Default `modules: [links, anchors]` ebenso wenig.
- `verifizierbar`: ja — `awk '/^[a-z][a-z-]*:/{t=$0} /docker run/ &&
  !/--disable mentions/{print t}'` über das frische `--print-mk` listet die
  sechs Rezepte ohne die Option.
- `klasse`: Universalaussage über eine Menge, die nicht ausgezählt wurde

### F-3 — Die Deckungs-Quote „15 von 15" steht in zwei lebenden Dokumenten, der Lauf meldet „16 von 16"

- `kategorie`: MEDIUM
- `quelle`: Hard Rule `AGENTS.md` §3.7 (*Ein Kommentar beschreibt, was da ist*
  — „Dieselbe Regel für Zustandsfelder"); eigener Lauf `make doc-mentions`
- `pfad`: `harness/sensors/doc-mentions.md`:16 und `.d-check.yml`:448
- `befund`: Die `## Vertrag`-Sektion gibt die Ausgabe des Laufs als
  `mentions: 15 von 15 Artefakt(en) erwähnt, über 1 Dokument(e)` an, der
  `.d-check.yml`-Kommentar als „dort vollständig (15 von 15 gemessen)". Der
  Lauf meldet auf demselben Commit `16 von 16` — die 16. Datei ist
  `harness/sensors/doc-mentions.md` selbst, angelegt im selben Commit
  `ec64373`, das beide Zahlen schreibt. Die zitierte Zeile hat es in diesem
  Repo nie gegeben.
- `verifizierbar`: ja — `make doc-mentions` gibt `mentions: 16 von 16
  Artefakt(en) erwähnt, über 1 Dokument(e)`, Exit 0.
- `klasse`: Zustandszahl aus der Vor-Messung in die Nach-Zustands-Beschreibung übernommen

### F-4 — „vier Zeilen" ohne Geltungsbereich, und „alle vier sagen dasselbe" trifft auf zwei nicht zu

- `kategorie`: LOW
- `quelle`: `AGENTS.md` §5 *Geltungsbereich einer Messung* (`seit slice-179`)
- `pfad`: `docs/plan/planning/in-progress/slice-184-dcheck-v0750-mentions.md`:44–46
- `befund`: Der Satz nennt die Diff-Form nicht, von der die Zahl abhängt:
  `diff` liefert 4 Zeilen, `diff -u` liefert 11. Von den 4 Zeilen der
  `diff`-Form sind zwei (`13c13` und `---`) Struktur der Ausgabe und sagen
  nichts über die Modul-Liste; inhaltlich ändert sich genau **eine** Zeile,
  die `# Verfügbar:`-Kommentarzeile. Die Aussage in der Sache — nur die
  Modul-Liste ändert sich, kein Konfigurations-Block — ist korrekt.
- `verifizierbar`: ja — `diff` und `diff -u` über die beiden
  `--print-config`-Ausgaben.
- `klasse`: Messzahl ohne benanntes Messverfahren

### F-5 — Das als „wörtlich" ausgewiesene Ziel-Form-Zitat endet vor der dritten genannten Grenze

- `kategorie`: LOW
- `quelle`: `v6.5.0` · `templates/harness/README.template.md`, Zeilen 110–113
- `pfad`: `docs/plan/planning/in-progress/slice-184-dcheck-v0750-mentions.md`:74–79
  und `harness/sensors/doc-mentions.md`:11–14
- `befund`: Das Original lautet „… bleiben still grün, **und geprüft wird nur,
  wo ein Link-Sensor über `harness/` läuft**." Beide Fundstellen setzen nach
  „still grün" einen Punkt und lassen den Nachsatz ohne Auslassungszeichen
  weg; zusätzlich ist „eine Datei ohne Index-Zeile" im Zitat fett, im Original
  nicht, ohne Kennzeichnung als Hervorhebung des Zitierenden.
- `verifizierbar`: ja — `grep -n "Grenze: Er prüft" -A3` in der Ziel-Form.
- `klasse`: Zitat als „wörtlich" ausgewiesen und still gekürzt

### F-6 — `## Sperren` der Sensor-Datei nennt zwei von vier gemessenen Abbruch-Klassen

- `kategorie`: LOW
- `quelle`: eigene Fixtures gegen den gepinnten Digest
- `pfad`: `harness/sensors/doc-mentions.md`:38–44
- `befund`: Gemessen bricht das Modul in vier Fällen fail-closed ab:
  (a) `mentions.artifacts gesetzt, mentions.documents fehlt` (und spiegelbildlich),
  (b) `das Modul mentions braucht mentions.artifacts UND mentions.documents
  (DC-FA-MENT-001, fail-closed)` bei fehlendem oder leerem Block,
  (c) `mentions.artifacts [<glob>] trifft kein Artefakt — eine Deckungs-Aussage
  ueber null Mitglieder ist keine`, (d) dasselbe für `mentions.documents`.
  Die Liste führt nur (b) plus die Befund-Sperre `artifact-unmentioned`. Gerade
  (c) und (d) sind die strukturelle Antwort auf
  `BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf` (5×) — sie stehen weder in der
  Sensor-Datei noch in §9 des Plans, der an dieser Stelle nur die schwächere
  Aussage führt, der Lauf *nenne* die Kandidatenzahl in seiner Erfolgszeile.
- `verifizierbar`: ja — vier Fixtures, jede mit Exit 2 und der zitierten Meldung.
- `klasse`: Gemessene Eigenschaft stärker als die dokumentierte

### F-7 — Zwei Fremdbefunde in derselben Prosa-Stelle wie F-1

- `kategorie`: INFO
- `quelle`: Vergleich `AGENTS.md` §4 gegen `Makefile`
- `pfad`: `AGENTS.md`:151–153 sowie `README.md`:173
- `befund`: Unabhängig von diesem Slice führt derselbe Absatz `doc-immutable`
  als „in `gates`", obwohl es in keinem der beiden Aggregate hängt (es läuft
  in der CI über die Commit-Range), und `doc-reviews` fehlt dort seit
  slice-160. `README.md`:173 führt eine Klammer-Liste „all inner gates
  (lint/test/coverage-gate/arch-check/doc-check/gate-consistency/guard-selftest)",
  die schon vor diesem Slice vier Gates ausließ und damit erkennbar
  beispielhaft ist. Beides liegt außerhalb des Gegenstands und ist hier nur
  notiert, weil F-1 dieselbe Stelle trifft.
- `verifizierbar`: nein — dieselbe ungewächterte Stelle wie F-1.
- `klasse`: Altbestand am Fundort

## Negativbefunde

- geprüft, ohne Befund: **Pin-Vollständigkeit.** `--print-mk` des gepinnten
  Digests gegen die committete `d-check.mk` diffed exakt die ausgewiesene
  Anpassung — sechs eingefügte Kommentarzeilen plus die Ersetzung
  `DCHECK_DIGEST ?=` durch den Digest, zusammen acht Diff-Zeilen. Kein
  weiterer Unterschied; die Behauptung „verbatim" trägt.
- geprüft, ohne Befund: **Digest ↔ Tag.** `docker inspect` auf
  `ghcr.io/pt9912/d-check:v0.75.0` liefert genau
  `sha256:18e9cd857f8db3569526d1f9a3cbeba8af51e9f2dd84c17a22444028b977c3da`;
  die `DCHECK_IMAGE`-Zeile des frischen Fragments nennt denselben Tag.
- geprüft, ohne Befund: **Restliche `v0.74.1`-Nennungen.** Vier Fundstellen
  außerhalb archivierter Slices, keine meint den aktuellen Pin: `.d-check.yml`:23
  („seit d-check v0.74.1" als Herkunft des `reviews`-Moduls), drei Stellen im
  Slice-Plan (Ausgangs-Stand, Rückführungs-Trigger, Bestands-Abgrenzung) und
  eine Evidence-Datei. Die eine Stelle, die den *aktuellen* Pin meinte
  (`AGENTS.md` §4, Zeile `doc-usage`), ist mitgewandert.
- geprüft, ohne Befund: **Erwähnungsformen.** Fixture mit sechs Artefakten und
  einem Dokument: Markdown-Link mit vollem Pfad, Inline-Code und nackte Prosa
  treffen; reiner Basename trifft nicht. Deckt sich mit der Aussage des Slice.
- geprüft, ohne Befund: **Volle Pfad-Form.** Diskriminierende Fixture: ein
  Dokument in `harness/` mit `](sensors/aaa.md)` erzeugt `artifact-unmentioned`
  für `harness/sensors/aaa.md` — das Modul löst Markdown-Links **nicht** relativ
  zum Dokument auf. Umgekehrt trifft `zzz/harness/sensors/ccc.md` als nackte
  Prosa, der Vergleich ist also ein Substring-Test auf dem vollen
  repo-relativen Pfad. Beide Richtungen bestätigen die Begründung, warum
  `harness/README.md` nicht als Dokument-Menge taugt.
- geprüft, ohne Befund: **Fail-closed bei fehlender Liste.** Fünf Fixtures
  (`artifacts` allein, `documents` allein, Block fehlt, Block leer, String
  statt Liste) — alle Exit 2, der String-Fall mit
  `cannot unmarshal !!str … into []string`. Deckt sich mit dem Slice.
- geprüft, ohne Befund: **Keine Konfigurations-Option, die die Pfad-Form
  ändern könnte.** Der YAML-Parser ist strikt: ein erfundener Schlüssel
  `mentions.strip-prefix` bricht mit `field strip-prefix not found`. Das
  Schlüsselpaar `artifacts`/`documents` ist damit vollständig — der negative
  Ausgang von Liefer-Punkt 3 ist nicht an einer übersehenen Option
  vorbeigelaufen.
- geprüft, ohne Befund: **Der negative Ausgang von Liefer-Punkt 3 hält auch
  gegen eine weitere Dokument-Menge.** `artifacts: ["docs/plan/adr/[0-9]*.md"]`
  gegen den ADR-Index gemessen: 39 von 39 `artifact-unmentioned`. Gegen
  `documents: ["**/*.md"]` — also gegen jede Markdown-Datei des Repos, eine
  schwächere Zusage als die des Eigenbau-Checks — bleiben immer noch **19 von
  39** unerwähnt. Es gibt also keine Glob-Kombination, die die
  ADR-Index-Prüfung in `tools/gate-consistency.sh` ablösen würde; der
  Liefer-Punkt ist nicht zu früh aufgegeben worden. **Geltungsbereich:**
  gemessen über Globs und Dokument-Mengen des Moduls, nicht über einen Umbau
  des ADR-Index selbst — würde der Index auf volle Pfade umgestellt, änderte
  das die Antwort, und genau das schließt der Slice mit dem Argument
  „Markdown-Links sind dateirelativ" aus.
- geprüft, ohne Befund: **Beide Mutations-Proben, auf einer Kopie
  nachgefahren.** (1) Zeile für `harness/sensors/symlink-check.md` aus
  `AGENTS.md` entfernt → `artifact-unmentioned`, `15 von 16`, Exit 1.
  (2) Neue Datei `harness/sensors/probe-ohne-zeile.md` → `artifact-unmentioned`,
  `16 von 17`, Exit 1. **Gegenrichtung:** unveränderter Bestand → `16 von 16`,
  0 Befunde, Exit 0, vor und nach beiden Proben. Beide Proben treffen ihren
  Gegenstand (die Probe liefert ihn nicht mit: Probe 1 entfernt eine Zeile aus
  dem Ist-Bestand, Probe 2 legt eine Datei an, die von keiner anderen Regel
  gedeckt ist).
- geprüft, ohne Befund: **Gate-Verdrahtung, alle Orte.** `.PHONY`
  (`Makefile`:50–54), `gates`-Aggregat (`Makefile`:207), `GATES`-Liste des
  Command-Guards (Zeile 79), `AGENTS.md` §4-Tabelle (Zeile 168),
  `harness/README.md` §Sensors (Zeile 93), `make help` und `make doc-help`
  führen das Target. Die CI erreicht es über `make ci` → `gates`.
  Ein vierter, übersehener Ort ist die **Prosa** desselben §4 — als F-1
  geführt, nicht hier.
- geprüft, ohne Befund: **Guard-Wirkung, beide Richtungen.** Der Hook direkt
  mit einem JSON-Payload gefüttert: `make doc-mentions` in einer Pipe liefert
  `"decision": "block"`, dieselbe Ausführung mit Ausgabe-Umleitung läuft
  durch. Die Aufnahme in `GATES` ist also wirksam, nicht nur deklariert.
- geprüft, ohne Befund: **Kein weiterer Registrierungs-Ort nötig.**
  `dcheck-phrase-selftest` kalibriert **phrasen**-basierte Konfigurationen
  (`reviews`-Triggerphrase, `structure`-`tasks-ignore-pattern`); `mentions` ist
  glob-/pfadbasiert und fällt nicht darunter. Die `targets`-Konfiguration
  (`doc-tables`, `authority`, `exempt-targets`) brauchte keine Ergänzung.
  `docs/user/` führt keine Gate-Liste.
- geprüft, ohne Befund: **Groß-/Kleinschreibung.** Der als Rest-Risiko benannte
  Sonderfall ist gemessen: das Modul vergleicht case-sensitiv (`ART/KLEIN.MD`
  deckt `art/klein.md` nicht). Die Einordnung des Slice — im Bestand
  gegenstandslos, weil alle Artefakt-Pfade klein sind — trifft zu.
- geprüft, ohne Befund: **Hard Rules.** §3.1 (jeder Handgriff über
  `make`/`docker`, keine Host-Toolchain), §3.2 (keine Go-Quelle berührt),
  §3.3 (kein `git mv` im Gegenstand), §3.5 (keine ADR angefasst), §3.6 (das
  neue Gate senkt keine Schwelle; die enge Dokument-Menge ist eine
  Einführungs-Grenze, keine Lockerung eines bestehenden Vertrags — die
  ADR-Index-Eigenbau-Prüfung bleibt unverändert in
  `tools/gate-consistency.sh`).
- geprüft, ohne Befund: **§3.7 an den drei neuen Kommentarblöcken.**
  `.d-check.yml`:431–459 (Kommentar 431–456, Konfiguration 457–459), `Makefile`:132–137 und `harness/sensors/doc-mentions.md`
  schreiben im Indikativ über den Ist-Zustand und tragen je eine der erlaubten
  Klassen (Zusage, Abgrenzung, Rang-Zeiger, Grenze). Eine Stelle ist
  grenzwertig und wird darum benannt statt verschwiegen: `doc-mentions.md`:33
  formuliert die ADR-Abgrenzung im Konjunktiv („das Modul sähe jede einzelne
  als unerwähnt"). Sie steht unter der Überschrift *Grenze — was das Grün
  nicht abdeckt* und beschreibt damit die geltende Reichweite, nicht die
  verworfene Alternative; kein Befund.
- geprüft, ohne Befund: **Slice-Form.** Drei Liefer-Punkte (Pin, Konfiguration,
  Ablösungs-Entscheidung), zwei Sub-Areas (`GATE`, `HARNESS`), Kopf mit
  `Verantwortlich:`/`Autor:`/berührten Spec-Stellen, §1 mit drei begründeten
  Ausschlüssen, §8 mit `Lerneintrag — Form: neuer Sensor`, §9 mit beiden
  *Vorgelagert*-Blöcken. Die drei Risiken aus §7 tragen je genau einen Ausgang
  aus der geschlossenen Menge (dreimal *entfallen*, jedes mit Begründung).
- geprüft, ohne Befund: **`DC-FA-MENT-001` ist keine erfundene Kennung** — das
  Werkzeug gibt sie selbst in seinen Abbruch-Meldungen aus.
- geprüft, ohne Befund: **Sensor-Datei-Form.** `## Vertrag` / `## Grenze — was
  das Grün nicht abdeckt` / `## Sperren` / `## Bindung`, identisch zu den zehn
  anderen vierteiligen Sensor-Dateien.
- geprüft, ohne Befund: **`make gates` Exit 0 und `make verify` Exit 0** auf
  dem Stand `98548af`, beide in eine Datei umgeleitet und am Exit-Code gelesen.
  Der `doc-mentions`-Schritt darin meldet `16 von 16`.
- geprüft, ohne Befund: **Commit-Disziplin.** Der Pin-Sprung ist ein eigener
  Commit **vor** der Konfiguration — genau die Trennung, auf der der zweite
  Risiko-Ausgang steht; alle drei Messages nennen `slice-184`, und der
  `(planning)`-Scope von `98548af` berührt nur `docs/plan/planning/`.

## Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 1 |
| MEDIUM | 2 |
| LOW | 3 |
| INFO | 1 |

**Finding-Klassen dieses Laufs:** Neues Gate ins Aggregat gehängt, die
Mandatory-Deklaration nicht nachgezogen (F-1, zweites Vorkommen nach
slice-160) · Universalaussage über eine nicht ausgezählte Menge (F-2) ·
Zustandszahl aus der Vor-Messung übernommen (F-3) · Messzahl ohne benanntes
Verfahren (F-4) · still gekürztes „wörtliches" Zitat (F-5) · gemessene
Eigenschaft stärker als die dokumentierte (F-6).

## Verdikt

**Nicht abnahmefähig ohne F-1.** Der Slice liefert, was er verspricht: der Pin
ist vollständig und verbatim, das Modul verhält sich in allen vier geprüften
Punkten so, wie der Plan es beschreibt, beide Mutations-Proben sind rot und
reproduzierbar, die Gegenrichtung grün, und der negative Ausgang von
Liefer-Punkt 3 hält auch gegen zwei alternative Konfigurationen, die der Plan
selbst nicht gemessen hat — er ist ein Ergebnis, kein Aufgeben. Die
Verdrahtung ist an sechs Orten belegt und in beide Richtungen am Guard geprüft.

Dagegen steht **ein** blockierender Befund: `AGENTS.md` §4 erklärt das neue
Gate in Prosa zu einer advisory-Funktion, während das Makefile es mandatorisch
führt — eine Harness-Lüge derselben Klasse, gegen die dieser Slice antritt, an
einer Stelle, die kein Sensor deckt. F-2 und F-3 sind vor der Closure zu
klären: beides sind Tatsachenbehauptungen, die gegen ein Repo-Artefakt fallen,
in einem Slice, der die Messgenauigkeit zu seinem eigenen Thema macht. F-4 bis
F-6 sind nice-to-fix, F-7 ist Altbestand am selben Fundort.
