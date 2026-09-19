# Review-Report: slice-195 — 2026-09-19

**Review-Art:** Code — geprüft gegen **Plan + Konventionen** (Modul 10 §Drei
Review-Arten): der Gegenstand sind drei Doku-Dateien, eine Deklarations-Zeile
und ein Register-Eintrag, kein Verhaltenscode.

**Gegenstand:** Commits `7aef989` (Handbuch + CHANGELOG), `282a26a`
(`harness/conventions.md`, neue Modus-Zeile), `ad089b6` (Messung, Closure-Notiz,
Register-Eintrag), Vorlauf `52b0467` (`git mv` + Ruhe-Marker)

**Skill:** `.harness/skills/reviewer.md` @ Stand `ad089b6`, Dateistand unverändert seit `60e7b66` (`slice-193`) · <!-- d-check:ignore (Adopter-spezifischer Skill-Pfad, existiert im Ziel-Repo ggf. nicht) -->
**Modell:** unbekannt (Subagent) · **Datum:** 2026-09-19

> **Zitier-Form** *(dieser Block bleibt stehen — er ist Norm, kein
> Ausfüll-Hinweis)*. Dieser Report friert ein; was er zitiert, bewegt sich
> weiter. Deshalb: **Kennung, nicht Adresse** — `slice-NNN` statt seines
> Lifecycle-Pfads, `make <target>` statt eines Links auf die Sensor-Datei, eine
> Baseline-Stelle als **Tag + Pfad in Inline-Code** statt als Link
> (`` `v<X.Y.Z>` · `regelwerk/<datei>.md` §<Abschnitt> ``). Der vendored Baum
> trägt genau einen Tag; der Sprung löscht den alten, und ein Link darauf färbt
> ein Artefakt rot, das niemand mehr anfassen darf. Ein `pfad`-Feld auf den
> **geprüften Gegenstand** ist davon nicht betroffen.

**Eingangs-Kontext** (die Verträge, gegen die geprüft wurde):

- `docs/plan/planning/in-progress/slice-195-handbuch-adapter-vokabel.md` (Stand `ad089b6`)
- [`ADR-0036`](../../docs/plan/adr/0036-port-richtung-inbound-outbound.md) (`Accepted`; rollen-abhängiges Vokabular)
- `AGENTS.md` §1, §3, §4, §5, §6 (Hard Rules, Doku-Regeln, Workflow)
- `harness/conventions.md` §Modus-Deklaration pro Sub-Area, §Adaptions-Block, §Baseline
- `harness/conventions/MR-023-id-schema-beobachtungs-kennung.md`
- `docs/plan/planning/observations/README.md`, `docs/plan/planning/README.md`
- `v6.6.0` · `regelwerk/modul-05-planning-harness.md` §Ziel-Form: Slice, §Lifecycle als State Machine ·
  `v6.6.0` · `regelwerk/modul-06-roadmap.md` §Das Beobachtungs-Register ·
  `v6.6.0` · `templates/docs/plan/planning/slice.template.md` §5 Closure-Trigger ·
  `v6.6.0` · `templates/docs/reviews/review-report.template.md`
- Slice-Vorlauf-Volltexte: `slice-188`, `slice-193`, `slice-194` (vor ihrer Archivierung) und `slice-190`/`slice-191` (in `open/`)

---

## Findings

### F-1 — Die zugesagte Änderungshistorie-Zeile im Handbuch fehlt, der Versions-Stempel bleibt stehen

- `kategorie`: MEDIUM
- `quelle`: Slice-Plan §3 (Umsetzungstabelle) gegen die eigene Praxis des Vorgänger-Slice
- `pfad`: `docs/plan/planning/in-progress/slice-195-handbuch-adapter-vokabel.md:69` · `docs/user/benutzerhandbuch.md:3` und §10
- `befund`: §3 nennt als Teil der Handbuch-Änderung eine „Änderungshistorie-Zeile"; `7aef989` ändert im Handbuch genau **eine** Zeile (365) und fügt keine Zeile in §10 hinzu. Die jüngste Historie-Zeile bleibt `1.40 / 2026-09-19 / slice-194`, und der Kopf trägt weiter `Handbuch-Version: 1.40` — zwei inhaltlich verschiedene Stände desselben Tages tragen denselben Stempel. Der unmittelbare Vorgänger-Slice hat für seine Handbuch-Änderung genau diese Zeile gesetzt.
- `verifizierbar`: ja — `git show 7aef989 -- docs/user/benutzerhandbuch.md` (ein Hunk, eine Zeile) gegen `sed -n '891,940p' docs/user/benutzerhandbuch.md`
- `klasse`: Zugesagter Teil der Umsetzung nicht geliefert (Historie-Zeile/Version fehlt)

### F-2 — Der Register-Index nennt die neue Familie nicht

- `kategorie`: MEDIUM
- `quelle`: `AGENTS.md` §4 („Dieselbe Regel gilt für jede abschließende Aufzählung neben einer maschinenlesbaren Quelle — `Makefile`, `modules:` in `.d-check.yml`, **ein Verzeichnis**") · `v6.6.0` · `regelwerk/modul-06-roadmap.md` §Das Beobachtungs-Register
- `pfad`: `docs/plan/planning/observations/README.md:30-34`
- `befund`: Der Register-Index zählt abschließend auf — „Aktuell **33** aktive Beobachtungen unter `BEO-GATE/`, `BEO-HARNESS/`, `BEO-PLAN/`, `BEO-SPEC/` und `BEO-KERN/`" —, und mit `BEO-USER/handbuch-vokabel-der-adapter-rolle/` führt der Slice die **erste neue Familie seit der Migration** (slice-139) ein, ohne sie dort zu nennen; sechs Familien stehen gegen fünf genannte. Die Zahl in derselben Zeile ist ebenfalls falsch (66 Verzeichnisse, 65 ohne `gestrichen` im `Stand:`) — diese Hälfte driftet **vorbestehend** seit der Migration und ist nicht durch diesen Slice entstanden.
- `verifizierbar`: ja — `ls docs/plan/planning/observations/` gegen die Zeile; `make verify-observations` meldet 66 Verzeichnisse
- `klasse`: Abschließende Aufzählung neben einer maschinenlesbaren Quelle

### F-3 — Die neue Modus-Zeile trägt eine Chronik-Wendung in der Form, die der geplante Sensor trifft

- `kategorie`: MEDIUM
- `quelle`: `AGENTS.md` §3.7 (abwesenden Text beschreiben) · `BEO-HARNESS/chronik-in-gelesenen-dateien` (3×, Ausgang *geplant* → slice-191)
- `pfad`: `harness/conventions.md:282`
- `befund`: Die Begründungs-Zelle endet mit „**bis slice-195 ohne Zeile**, weshalb ein Fund in dieser Familie keinen Register-Ort hatte" — „bis slice-NNN …" ist die im Register namentlich benannte, greppbare Schreibweise der Chronik-Klasse, und slice-191 plant ein `forbid-pattern`-Modul genau für die lebenden Dateien `AGENTS.md`, `harness/*.md`, `docs/plan/planning/README.md`. Dieselbe Datei trägt die Wendung bereits einmal (Zeile 230, aus slice-076); die Nachbarzeile der Sub-Area `PLAN` und die Zeile `HARNESS` formulieren die verwandte Aussage paraphrasiert („davor war die Praxis unbelegt", „hatte bis dahin keinen Ort").
- `verifizierbar`: ja — `grep -n -i 'bis slice-' harness/conventions.md`
- `klasse`: Chronik-Phrase in einer gelesenen Datei

### F-4 — Der Slice-Plan lässt den Abschnitt *Closure-Trigger* der Ziel-Form aus

- `kategorie`: MEDIUM
- `quelle`: `v6.6.0` · `templates/docs/plan/planning/slice.template.md` §5 Closure-Trigger · `AGENTS.md` §5 (Slice-Form) · `docs/plan/planning/README.md` §Beim Kopieren der Slice-Ziel-Form
- `pfad`: `docs/plan/planning/in-progress/slice-195-handbuch-adapter-vokabel.md` (Abschnittsfolge §5 → §6)
- `befund`: Der Plan führt §5 *Trigger* (Start und Rückführungen) und springt danach zu §6 *Risiken*; der Abschnitt *Closure-Trigger*, den die Ziel-Form führt und den die fünf nächsten Slices (`slice-188`, `slice-190`, `slice-191`, `slice-193`, `slice-194`) ausnahmslos tragen, fehlt. Die beobachtbaren Kriterien reisen in den DoD-Punkten von §4 mit, und der Gate-Lauf steht als feste Zeile darunter — der Abschnitt selbst fehlt. Kein Gate deckt die Lücke: die `structure`-Regeln prüfen DoD-Größe, Closure-Struktur in `done/`, Lerneintrag-Form und Kopffelder; nach der Archivierung (Stub ohne Abschnittsüberschriften) wäre sie nicht mehr sichtbar.
- `verifizierbar`: ja — Abschnitts-`grep` des Plans gegen die Vorlage und die fünf Nachbar-Slices
- `klasse`: Abschnitt der Ziel-Form beim Kopieren ausgelassen

### F-5 — Die Zeilenangaben in §1 lösen gegen den Stand der Datei nicht mehr auf

- `kategorie`: LOW
- `quelle`: Slice-Plan §1 „Ausgangslage (gemessen)" · Register-Klasse `BEO-PLAN/dateiinhalt-aus-gedaechtnis-zitiert`
- `pfad`: `docs/plan/planning/in-progress/slice-195-handbuch-adapter-vokabel.md:26-28`
- `befund`: §1 zitiert `:350` (Beispielstruktur) und `:726` (§4 „Richtung"); dieselben Stellen liegen nach der Handbuch-Änderung von slice-194 (+15 Zeilen) bei **365** und **751/754**. §2 desselben Dokuments nennt für die erste Stelle korrekt 365. Beide Zahlen waren bei der Anlage von slice-195 (`590fe8b`) zutreffend und sind beim Abschluss nicht nachgezogen.
- `verifizierbar`: ja — `grep -n` gegen `docs/user/benutzerhandbuch.md`, und `git show 590fe8b:docs/user/benutzerhandbuch.md`
- `klasse`: Dateiinhalt aus Gedächtnis zitiert statt Datei gelesen

### F-6 — Das „Ergebnis" des zweiten Zählers ist die klassifizierte Teilmenge, nicht die Rohzahl

- `kategorie`: LOW
- `quelle`: Slice-Plan §2 · `AGENTS.md` §5 („Wer eine Menge zählt, zählt sie zweimal verschieden") · `BEO-PLAN/kandidaten-klassifikation-groeber-als-der-kandidat`
- `pfad`: `docs/plan/planning/in-progress/slice-195-handbuch-adapter-vokabel.md:56`
- `befund`: Die Zeile meldet für Zähler 2 („`awk` über Zeilen, die `dapter` **und** `inbound\|outbound` führen") **1 Treffer**; derselbe Zähler liefert über die drei Dokumente **6** Zeilen (257, 365, 753, 759, 762, 935). Die Reduktion auf 1 geschieht durch die Klassifikation „führt das Vokabular richtig" im Satz darunter, nicht durch den Zähler; der Register-Beleg übernimmt die Formulierung („beide nennen **eine** Stelle"). Der Befund selbst — genau **eine** fehlerhafte Stelle — hält beiden Nachzählungen stand.
- `verifizierbar`: ja — `awk '/dapter/ && /inbound|outbound/'` über die drei Dokumente im Stand vor `7aef989`
- `klasse`: Zähler-Ergebnis ist die klassifizierte Teilmenge statt der Rohzahl

### F-7 — Der Ausschluss verweist auf eine Adresse, die den Punkt selbst ausschließt

- `kategorie`: LOW
- `quelle`: `v6.6.0` · `regelwerk/modul-05-planning-harness.md` §Ziel-Form: Slice (Klasse 1: „Ein Folge-Slice, der den verwiesenen Punkt selbst ausschließt oder vor dem verweisenden schließt, ist keine [Adresse]")
- `pfad`: `docs/plan/planning/in-progress/slice-195-handbuch-adapter-vokabel.md` §1, erster Ausschluss
- `befund`: „Ein Nachzug der Lab-Config des Konsumenten … er ist Gegenstand von slice-194" — slice-194 führt denselben Punkt in §1 als eigenen Ausschluss („ein anderer Vorgang in einem anderen Repo") und schließt ihn in §6 mit *entfallen* („ihn hier zu einem Folge-Slice zu machen hieße, eine Adresse zu erfinden, die die Sendung nicht annehmen kann"); der Slice lag bei der Anlage von slice-195 bereits in `done/`.
- `verifizierbar`: ja — der archivierte Volltext von slice-194 (`git show 83368ef^:docs/plan/planning/done/slice-194-portscope-richtungssegment.md`)
- `klasse`: Ausschluss verweist auf eine Adresse, die den Punkt selbst ausschließt

### F-8 — `Verantwortlich:` steht auf `—`, obwohl der Slice in `in-progress/` liegt

- `kategorie`: LOW
- `quelle`: `v6.6.0` · `regelwerk/modul-05-planning-harness.md` §Lifecycle als State Machine („bis zur Priorisierung steht dort `—`")
- `pfad`: `docs/plan/planning/in-progress/slice-195-handbuch-adapter-vokabel.md:12`
- `befund`: Der Kopf trägt weiter „**Verantwortlich:** — (bis zur Priorisierung)", während der Slice seit `52b0467` in `in-progress/` liegt und das WIP-Limit belegt. `slice-188` und `slice-193` tragen an derselben Stelle „Claude — gesetzt beim Übergang nach `in-progress/`", der ebenfalls wellenlose `slice-194` dagegen wie hier `—`. Kein Gate prüft das Feld.
- `verifizierbar`: ja — Kopffeld-Vergleich der Nachbar-Slices (vor ihrer Archivierung)
- `klasse`: Kopf-Feld `Verantwortlich:` beim Übergang nicht gesetzt

### INFO-1 — Die Test-Fixtures tragen `adapters/outbound` weiter (außerhalb des erklärten Geltungsbereichs)

- `kategorie`: INFO
- `quelle`: Slice-Plan §2 (Geltungsbereich: Prosa der drei Dokumente unter `docs/user/`)
- `pfad`: `internal/hexagon/core/rules_test.go:1816` u. a. (Schicht-**Name** `outbound`, `Role: "adapter"`)
- `befund`: Dort ist `outbound` der Name einer Adapter-Schicht, nicht der `direction`-Wert der Rolle — ADR-0036 regelt den Wertebereich der Dimension, nicht Schicht-Namen. Der Fund liegt außerhalb des erklärten Geltungsbereichs und ist kein zweites Vorkommen derselben Sache; dieselbe Schreibweise steht auch in `docs/plan/adr/0028-ziel-glob-schattenwurf.md` (Beispiel-Ausgabe). Ein späterer Sweep sollte sie nicht als dasselbe Muster lesen.
- `verifizierbar`: ja — `grep -n 'adapters/outbound' internal/hexagon/core/rules_test.go`
- `klasse`: Gleiche Schreibweise, anderer Gegenstand (Schicht-Name gegen `direction`-Wert)

## Gefahrene Gates

| Lauf | Exit | Geltungsbereich |
|---|---|---|
| `make gates` | **0** | ganzer Arbeitsbaum `ad089b6`: `lint`, `test`, `coverage-gate` (**96,40 %** ≥ 90), `arch-check`, `doc-check`, `doc-targets`, `doc-planning`, `doc-workflows`, `doc-reviews`, `doc-mentions`, `gate-consistency`, `version-coherence`, `suppression-check`, `symlink-check`, `dcheck-phrase-selftest`, `guard-selftest`, `ci-range-selftest`, `record-gates`; alle `d-check`-Läufe **583 Dateien / 0 Befunde** |
| `make verify` | **0** | Verifikations-Schicht: `verify-risiko-ausgaenge` (11 Slices), `verify-observations` (**66** Verzeichnisse), `doc-structure` (0 Befunde), `doc-complete` (21 Anforderungen, 0 Waisen) |

Beide Läufe als eigener Aufruf, Ausgabe in eine Datei, Exit-Code getrennt geprüft (keine Pipe).

**Grenzen der Gate-Läufe** (Geltungsbereich, Mess-Regel 1): `doc-check` sieht Markdown-Links,
Anker und Kennungen — Prosa sieht es nicht, und die Widersprüche dieses Slice sind Prosa. **Kein
Gate** deckt: den Register-Index in `observations/README.md` (keine Aufzählung wird geprüft), die
Abschnittsfolge eines Slice-Plans, das Kopffeld `Verantwortlich:`, die Handbuch-Änderungshistorie
und die Zählung in §2 des Plans. Die beiden Zählungen dieses Reports sind darum **eigene** Zähler
(kein Gate): `grep` auf die alte Wendung und `awk` auf `dapter` + `inbound|outbound` über die drei
Dokumente im Stand `52b0467`, dazu `find`/`head -1` über die Register-Verzeichnisse gegen
`make verify-observations`.

## Negativbefunde

- geprüft, ohne Befund: `docs/user/benutzerhandbuch.md` — die geänderte Stelle ist inhaltlich richtig (`driving`/`driven` ist an `role: adapter` der seit ADR-0036 gültige Wertebereich) und deckungsgleich mit §4 (ab 751) und der Regel-Tabelle (ab 249).
- geprüft, ohne Befund: `docs/user/benutzerhandbuch-standard.md` und `docs/user/releasing.md` — zweiter Zähler: **kein** `adapter`, kein `driving|driven`, kein `inbound|outbound`; die Aussage „führen das Vokabular nicht" ist bestätigt (Gegenprobe des ersten Zählers).
- geprüft, ohne Befund: der neue Register-Eintrag (`observation.md`, `state.md`, `evidence/slice-195.md`) — Bezeichnung und `**Sub-Area:**` vorhanden, `Stand: offen (1×)` deckt sich mit der einen Belegdatei, der Dateiname ist ein gültiger Vorgang (`slice-195`), `Ehemals:` fehlt zu Recht (Erstauftreten), Slug ist lowercase Kebab-Case.
- geprüft, ohne Befund: die Kürzel-Vergabe `USER` — kollidiert mit keinem der acht geführten Kürzel und mit keinem `id-patterns`-Präfix; MR-023 erklärt die Kürzel-Liste ausdrücklich für generisch („wird ohnehin mit jeder neuen Sub-Area gepflegt"), also ist **kein** Adaptions-Eintrag nachzuziehen. Die Zeile ist formal wie ihre acht Geschwister (sechs Spalten, Achsen `2,3`, Modus `Greenfield`, `n/a (GF)`), und die Sub-Area erfüllt die Schwelle ≥ 2 von 3 Achsen über die Achsen 2 (eigene Diskrepanz-Zeile) und 3 (eigene Pfad-Familie `docs/user/`).
- geprüft, ohne Befund: `CHANGELOG.md` — Eintrag unter `[Unreleased]` / `### Fixed`, ADR-Link relativ korrekt; `CHANGELOG.md` ist von der `ids`-Linkpflicht exempt, die Aussage deckt sich mit der Nachzählung.
- geprüft, ohne Befund: Referenz-Richtung und Form — kein Spec-Stratum berührt, keine Abwärtsverweise, keine erfundenen Kennungen, `slice-\d{3}` in keiner ADR (nicht berührt); `matrix`-Modul grün.
- geprüft, ohne Befund: Lifecycle — `52b0467` ist ein reiner `git mv` (similarity 100 %); die Inhaltsänderung (Ruhe-Marker der Roadmap) reist in derselben Form wie in den Vorläufer-Moves; WIP-Limit = 1 (kein zweiter Slice in `in-progress/`); `doc-planning` grün; alle Verweise auf den wandernden Slice lösen auf.
- geprüft, ohne Befund: `harness/conventions.md` als Ganzes (Pflicht-Blick nach `AGENTS.md` §1) — Source Precedence unverändert, ID-Schema-Deklaration unberührt (die `MR-000`-Liste führt keine Kürzel), Adaptions-Block unberührt. Die Zahl „über **55** Slices" im Absatz unter der Tabelle ist gealtert, als Untergrenze aber nicht falsch.
- geprüft, ohne Befund: `spec/`, `docs/plan/adr/` (ADR-0036 gelesen), `tools/`, `.d-check.yml` — nicht berührt.

## Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 0 |
| MEDIUM | 4 |
| LOW | 4 |
| INFO | 1 |

**Finding-Klassen dieses Laufs:** Abschließende Aufzählung neben einer maschinenlesbaren Quelle ·
Chronik-Phrase in einer gelesenen Datei · Zugesagter Teil der Umsetzung nicht geliefert ·
Abschnitt der Ziel-Form beim Kopieren ausgelassen · Dateiinhalt aus Gedächtnis zitiert statt Datei
gelesen · Zähler-Ergebnis ist die klassifizierte Teilmenge statt der Rohzahl · Ausschluss verweist
auf eine Adresse, die den Punkt selbst ausschließt · Kopf-Feld `Verantwortlich:` beim Übergang
nicht gesetzt

## Verdikt

**Merge-blockierend:** ja — vier MEDIUM: die unerfüllte Zusage aus §3 (F-1), der Register-Index,
der die neue Familie nicht führt (F-2), die Chronik-Wendung in der Datei, die der geplante Sensor
durchsucht (F-3), und der fehlende Abschnitt *Closure-Trigger* der Ziel-Form (F-4). Alle vier sind
Doku-Änderungen, keine Verhaltensänderung; F-1 bis F-3 kosten je eine Zeile, F-4 einen Abschnitt,
der aus dem Inhalt von §4/§5 des Plans gebildet werden kann.

**Kein HIGH**, und das ist geprüft statt geschätzt: die HIGH-Liste des Skills habe ich Punkt für
Punkt gegen die Artefakte gehalten. Kein Verstoß gegen §3.1–§3.6, keine erfundene Kennung, kein
abwärts gerichteter Verweis, keine Behauptung eines Gates ohne Make-Target. Zwei Lesarten kamen in
Frage und tragen nicht: F-1 ist eine **Zusage über die eigene Umsetzung**, keine Tatsachenbehauptung
über eine Regel; und die „33" in F-2 ist eine nachweislich falsche Zahl — aber eine **vorbestehende**
(seit slice-139 nicht nachgezogen), die dieser Diff nicht erzeugt hat; neu erzeugt hat er die
fehlende Familie.

**Übergabe:** Findings gehen an den Implementer; ein Plan-Defekt im Sinne der Rückkante
(Review → Plan) liegt nur in F-4 und F-7 vor, beides Form/Adresse, nicht Ziel. Die
Finding-Klassen gehen zusätzlich in die Slice-Closure §7 und von dort in den Zähler — für
`Abschließende Aufzählung neben einer maschinenlesbaren Quelle` ist die Klasse in
`AGENTS.md` §4 bereits verkörpert (`seit slice-188`), für `Chronik-Phrase in einer gelesenen Datei`
steht der Ausgang als *geplant* → `slice-191`; für die übrigen sechs Klassen gibt es im Register
**keinen** Eintrag. F-1 betrifft den Stand des Handbuchs bei der Closure und ist dort mit einer
Zeile zu erledigen. Dieser Report ist **Lauf-Beleg** (dieser Diff, dieser Skill, dieses Modell,
dieses Verdikt) und ersetzt keine Verifikation — DoD-/Spec-Konformität prüft der Verifier separat
(Modul 11).
