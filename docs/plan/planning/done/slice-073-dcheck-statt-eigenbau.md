# slice-073 — d-check statt Eigenbau: Analyse und CR (zur Abnahme)

**Status:** open — **Analyse zur Abnahme** (Messung + CR-Formulierung; **keine** Ablösung im
Code — die ist Folge-Slice). Der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Deckt:** Maintainer-Frage vom 2026-08-09 — *„Warum können wir d-check nicht anstatt der
`verify-*.sh` verwenden?"*; Bezug zu
[MR-006](../../../../harness/conventions.md#mr-006--baseline-committet-vendored-statt-per-url-referenziert)
und [ADR-0021](../../adr/0021-commits-modul-trace-check.md).
**Bezug:** Roadmap-Zeile *Aktuelle Welle* in der [Roadmap](../in-progress/roadmap.md).

---

## 1. Auslöser

**897 Zeilen Shell** prüfen in `tools/` Invarianten über Markdown — während das Repo mit d-check
ein gepinntes, hermetisches Doku-Gate mit **19 Regelmodulen** einbindet.

Das ist a-checks eigene Doktrin, angewandt auf a-check: der Maßstab dieses Projekts ist
**Skript-Reduktion über die Konsumenten-Flotte**. Ein Shell-Skript im Konsumenten-Repo, das eine
**generische** Invariante prüft, ist ein CR-Kandidat für das Werkzeug — nicht dauerhaft
Konsumenten-Sache. Hier ist a-check der Konsument, und `tools/` ist sein P-Rest.

Die Frage wurde nie gestellt. [slice-050](../done/slice-050-verify-schicht.md) hat die
Doppelung nur für die *Referenzmatrix* geprüft („ein zweiter Prüfer wäre Doppelung") und die
Struktur-Sensoren nicht betrachtet.

**Erste Messung (Handbuch `d-check` §6 Regelmodule, Stand v0.51.1):**

| Eigenbau | Zeilen | d-check-Pendant | Befund |
|---|---|---|---|
| `gate-consistency` Checks (1)+(2) | ~45 von 308 | Modul **`targets`** (`gate-phantom`/`gate-undocumented`) | **vollständige Doppelung** |
| `verify-slice-links` | 146 | `links` prüft nur den Ist-Ort | Verallgemeinerung fehlt |
| `verify-closure-notes` | 146 | — | kein Struktur-Modul |
| `verify-slice-form` | 166 | — | kein Struktur-Modul |
| `verify-ac-form` | 131 | — | kein Struktur-Modul |

**Eine dritte Achse kam am selben Tag dazu:** das CLI-Werkzeug `mq` (Markdown → Node-Baum,
jq-artige Abfragen; als Skill `.claude/skills/markdown-mq/` hinterlegt). Alle fünf Eigenbau-Prüfer
bauen Markdown-Struktur mit **zeilenbasierten** Mitteln nach — `sed` schneidet Code-Fences heraus,
`grep -oE` extrahiert Links, Checkboxen werden über Abschnitts-Grenzen hinweg gezählt. Genau
daraus entstanden `F-4` und `R-068-F4`.

Gemessen an einer Fixture mit Inline-Link, Referenz-Link und einem Link im Code-Block:

| Weg | Inline | Referenz (`F-4`) | im Code-Block |
|---|---|---|---|
| heutiges `links_of()` (`sed`+`grep`) | gefunden | **übersehen** | korrekt ignoriert (per `sed`-Hack) |
| `mq '.link'` | gefunden | **übersehen** | korrekt ignoriert (strukturell) |
| `mq '.link'` + `.link_ref` + `.definition` | gefunden | **gefunden** | korrekt ignoriert (strukturell) |

`mq` löst `F-4` also **nicht** durch einen Selektor — die Auflösung `link_ref` → `definition` →
Ziel bleibt eigene Logik. Was es liefert, ist die *Struktur*: Code-Blöcke fallen ohne
`sed`-Vorfilter weg, und ein Referenz-Link ist ein eigener Node-Typ statt eines nicht getroffenen
Regex-Falls.

**Der Haken ist die Hermetik.** `mq` ist Node-basiert und ein Host-Werkzeug;
[`AGENTS.md`](../../../../AGENTS.md) §3.1 verlangt Docker/make-only. Als Analyse-Werkzeug für
Menschen und Agenten ist es unbedenklich, **in einem Gate** wäre es eine Host-Abhängigkeit — dort
nur über ein Image. Diese Unterscheidung gehört in die Messung, nicht in eine Fußnote.

Der harte Fund ist die erste Zeile der Tabelle oben. `targets` ist konfiguriert als

```yaml
targets:
  makefiles: [Makefile]     # Regelnamen-Quelle
  doc-tables: [AGENTS.md]   # Richtung 1 => gate-phantom
  authority: AGENTS.md      # Richtung 2 => gate-undocumented
  exempt-targets: []        # Utility ohne Doku-Pflicht
```

— das ist Zeile für Zeile, was `gate-consistency.sh` in `doc_targets()`,
`check_documented_exist()` und `UTILITY_TARGETS` selbst gebaut hat. `make doc-targets` **existiert
und ist in [`AGENTS.md`](../../../../AGENTS.md) §4 dokumentiert**, läuft aber in keinem Aggregat;
`gates` fährt stattdessen den Eigenbau.

**Der CR-Weg ist zweifach eingefahren — und trägt einen Namen.** In d-check existiert **kein**
offener CR für ein Struktur-Modul (`open/` und `next/` sind leer, kein Eintrag in Lastenheft oder
Spezifikation). Aber genau dieses Muster ist dort zweimal gelaufen:

| d-check-Modul | abgelöstes Skript | Anlass laut Lastenheft |
|---|---|---|
| `vcs` (`DC-FA-VCS-001`) | `adr-immutable-check.sh` | *„vollständig mechanisieren (Copy-Drift über die Repo-Familie) … die verteilte git-Garantie für Konsumenten ohne Skript-Kopie (**„verteilen statt kopieren"**)"* |
| `targets` (`DC-FA-TGT-001`) | `gate-consistency.sh` | DoD-Punkt *„Paritäts-Mutations-Beleg vs. `gate-consistency.sh`"* |

Beide Male: ein Shell-Skript aus der Repo-Familie wird zum Modul, mit **Paritätsbeleg gegen das
abgelöste Skript**. Der CR-Text dieses Slice hat damit zwei Präzedenzfälle, an deren Form er sich
zu halten hat — einschließlich der Erwartung, dass ein Paritäts-Mutations-Beleg dazugehört und
nicht nachgereicht wird.

## 2. Betroffene Module

Dieser Slice ändert **keinen** Sensor. Er liefert zwei Dokumente:

1. eine **Abdeckungs-Messung** je Eigenbau-Prüfung gegen die 19 d-check-Module,
2. einen **CR-Text** für d-check zu dem, was generisch ist.

Berührt werden damit nur `docs/plan/planning/` und (für den CR) `docs/`.

## 3. Auszuführende Gates

`make doc-check`, `make verify` — beide unverändert grün, weil nichts am Code geändert wird.

**Die Messung ist die Beleg-Arbeit dieses Slice**, nicht ein Gate-Lauf. Sie ist erst gültig, wenn
je Eigenbau-Prüfung eine der vier Aussagen belegt ist:

| Aussage | Beleg |
|---|---|
| **abgelöst** — ein d-check-Modul prüft dasselbe | Lauf beider Prüfer gegen dieselbe Negativ-Fixture, gleiche Befund-Menge |
| **CR-fähig** — generisch, aber kein Modul deckt es | benannte Abstraktion, die nicht a-check-spezifisch ist |
| **bleibt lokal, strukturell** — repo-spezifisch, aber die Markdown-Struktur trägt die Prüfung | der Node-Typ, der die zeilenbasierte Heuristik ersetzt, **plus** der Weg, wie er hermetisch verfügbar wird |
| **bleibt lokal, zeilenbasiert** — weder generisch noch strukturell fassbar | die Repo-Eigenheit, an der es hängt, ist benannt |

Ohne die erste Zeile bliebe der Slice eine Vermutung — genau der Fehler, an dem der
Vorgänger-Entwurf `4b029e4` gescheitert ist.

## 4. Was bewusst nicht getan wird

- **Die Ablösung selbst.** `doc-targets` in `gates` zu hängen und die 45 Zeilen zu entfernen ist
  ein **Folge-Slice** mit eigener Negativ-Probe. Analyse und Eingriff im selben Schnitt wäre
  genau das Dehnen, das `R-068-F5` am Vorgänger-Entwurf bemängelt hat.
- **Den CR an d-check stellen.** Dieser Slice **formuliert** ihn; ihn einzureichen ist ein Akt
  gegenüber einem Fremdrepo und gehört dem Maintainer.
- **`suppression-check` und `regelwerk-check`.** Sie prüfen Go-Quellen bzw. Datei-Hashes — nicht
  d-checks Domäne. Ohne Messung sind sie hier nicht einmal Kandidat.
- **Ein Urteil über die verbleibenden `gate-consistency`-Checks.** Pin-Konsistenz (4),
  `.d-check.yml`-Modulliste (3) und `.PHONY`-Vollständigkeit (5, aus
  [slice-068](../done/slice-068-phony-vollstaendig.md)) werden **mitgemessen**, aber dieser Slice
  behauptet über sie nichts vorab.

## 5. DoD

- [x] Abdeckungs-Messung liegt vor: je Eigenbau-Prüfung eine der vier Aussagen aus §3, **mit dem
      dort geforderten Beleg**. Beleg: die Tabelle in §6, mit Fixture-Läufen für die
      „abgelöst"-Zeile.
- [x] CR-Text für d-check liegt vor, mit einer Abstraktion, die **nicht** a-check-spezifisch ist —
      formuliert als Modul-Vertrag (was es prüft, welche Grund-Codes, welche Konfiguration), nicht
      als Wunschliste, **und mit dem Paritäts-Mutations-Beleg als benanntem DoD-Punkt**, wie ihn
      `vcs` und `targets` beide tragen. Beleg: §8, gegen die Form der 19 bestehenden Module und
      gegen die zwei Präzedenz-CRs aus §1 gehalten.
- [x] `make doc-check` und `make verify` grün — **Ausgabe in eine Datei**, Exit-Code getrennt
      geprüft, nie in eine Pipe.

## 6. Closure-Notiz

**Geliefert:** die Abdeckungs-Messung (unten) und der CR-Text in §8 — zwei Anträge an d-check,
formuliert als Modul-Vertrag mit Grund-Codes, Konfigurations-Schema und Paritäts-Mutations-Beleg.
**Keine** Ablösung im Code; die ist Folge-Slice.

**Die Messung, je Prüfung eine Aussage:**

| Eigenbau | Zeilen | Aussage | Beleg |
|---|---|---|---|
| `gate-consistency` (1)+(2) | ~45 | **abgelöst** | Modul `targets`; zwei Fixtures, beide Richtungen, identische Befundmenge — in [slice-074](../done/slice-074-doc-targets-wirksam.md) gefahren und dort konfiguriert |
| `verify-closure-notes` | 146 | **CR-fähig** | keins der 19 Module prüft Abschnitts-Struktur; Abstraktion in §8 CR 1 |
| `verify-slice-form` | 166 | **CR-fähig** | dito; `max-tasks` abschnitts-treu statt dateiweit |
| `verify-ac-form` | 131 | **CR-fähig** | dito; `section-pattern` + `require-strong` |
| `verify-slice-links` | 146 | **CR-fähig** | `links` prüft nur den Ist-Ort; Erweiterung in §8 CR 2 |
| `gate-consistency` (3) `.d-check.yml`-Module | ~20 | **bleibt lokal** | YAML, und **zirkulär**: d-check würde prüfen, ob d-check richtig konfiguriert ist |
| `gate-consistency` (4) Pin-Konsistenz | ~215 | **bleibt lokal** | **zwei der fünf Pin-Stellen sind kein Markdown** — siehe Nachtrag unten |
| `gate-consistency` (5) `.PHONY` | ~28 | **bleibt lokal** | Makefile ↔ Makefile, **keine Doku-Beziehung**; jeder `targets`-Befund hängt dagegen an einer Doku-Zeile |
| `suppression-check`, `regelwerk-check` | 153 | **bleibt lokal** | Go-Quellen bzw. Datei-Hashes; nicht d-checks Gegenstand |

**Nachtrag 2026-08-09 — Korrektur einer unbelegten Zeile.** Die ursprüngliche Fassung fasste
(3)(4)(5) in **einer** Zeile zusammen mit der Begründung *„keine Doku-Referenz-Invarianten"*. Das
war ein Lektüre-Urteil, kein Beleg — genau die Sorte, die dieselbe Tabelle für alle anderen Zeilen
ausschließt. Auf Nachfrage des Maintainers nachgemessen:

`versions` ist **nicht** auf Tags festgelegt, wie der Kommentar in
[`gate-consistency.sh`](../../../../tools/gate-consistency.sh) seit slice-018 behauptet
(*„d-checks tag-basierte Module"*) — `pin-pattern` ist ein konfigurierbarer Regex. Mit
`'a-check@(sha256:[0-9a-f]{64})'` greift das Modul den Digest tatsächlich und meldet ihn.

Der Grund ist ein anderer und härter: von den fünf harten Pin-Stellen liegen **zwei außerhalb von
Markdown** — [`a-check.mk`](../../../../a-check.mk) und
[`internal/cli/cli.go`](../../../../internal/cli/cli.go). Der Testlauf fand die Pins in
`README.md:86` und `README.de.md:91`, die beiden anderen sieht d-check nicht; sein Lastenheft führt
*„Prüfung von Nicht-Markdown-Formaten … als eigenständige Dateien"* ausdrücklich unter Out of
Scope. Ein Gleichheits-Gate, das zwei von fünf Stellen nicht sieht, kann Gleichheit nicht
behaupten.

Die Ausnahme bestätigt es: `targets` liest mit `makefiles:` sehr wohl Nicht-Markdown — aber nur
als *Vergleichsquelle*; jeder seiner Befunde (`gate-phantom`, `gate-undocumented`) hängt an einer
Doku-Zeile. Genau das fehlt der `.PHONY`-Prüfung, die Makefile gegen Makefile hält.

**634 von 897 Zeilen sind ablösbar**, davon 45 sofort. Keine Prüfung fiel in die vierte Kategorie
(*bleibt lokal, zeilenbasiert*) — was zeilenbasiert bleibt, bleibt es nicht mangels Abstraktion,
sondern weil es kein Markdown prüft.

**Lerneintrag — Form: benannte Spec-Lücke.** d-checks Modulsatz deckt **Referenz**-Invarianten
lückenlos ab (Ziel existiert, Kennung verlinkt, Richtung erlaubt, Target deklariert, Core
unverändert) — aber **keine Struktur-Invariante** innerhalb eines Dokuments. Das ist keine
Nachlässigkeit, sondern eine Grenze, die nie ausgesprochen wurde: die Module sind entlang der
Frage *„zeigt dieses Dokument korrekt auf andere?"* gewachsen, nie entlang *„ist dieses Dokument
selbst richtig gebaut?"*.

**Die Ursache**, warum das erst jetzt auffällt: Jeder Adopter füllt die Lücke lokal mit einem
Skript, und ein lokales Skript sieht wie Repo-Eigenheit aus, nicht wie fehlende Werkzeug-Fähigkeit.
Genau deshalb ist die Frage des Maintainers — *„brauchen wir die verify-Skripte überhaupt?"* — der
Auslöser gewesen und nicht ein Sensor: **kein Gate kann melden, dass ein anderes Gate hätte
existieren können.**

**Zwei beobachtbare Closure-Kriterien:**

1. §8 nennt für beide Anträge Konfigurations-Schema, Grund-Codes und Paritäts-Mutations-Beleg —
   dieselbe Form, die `DC-FA-VCS-001` und `DC-FA-TGT-001` tragen.
2. Jede der sieben Zeilen der Messtabelle trägt eine der vier Aussagen aus §3 mit ihrem Beleg;
   die einzige „abgelöst"-Zeile ist durch Fixture-Läufe gedeckt, nicht durch Lektüre.

**Folge-Slices:** die Ablösung von `gate-consistency` (1)+(2) samt Einhängung von `doc-targets` in
`gates` — noch nicht geschnitten, Voraussetzung ist nichts weiter als eine Entscheidung. Die
Einreichung der beiden CRs bei d-check bleibt Maintainer-Sache (§4).

## 7. Sub-Area-Modus

Alle berührten Sub-Areas GF.

## 8. CR-Text für d-check

Dieser Abschnitt **ist** die Lieferung aus §5 DoD 2. Er liegt im Slice, weil §4 das Einreichen
ausdrücklich dem Maintainer überlässt — ein separates Dokument wäre eine neue Doku-Gattung ohne
deklarierten Ort.

Zwei Anträge, unabhängig voneinander umsetzbar.

---

### CR 1 — Neues Modul `structure`: Struktur-Invarianten über Dokumentklassen

**Anlass.** Adopter-Repos dieses Regelwerks bauen Struktur-Regeln über ihre eigenen Dokumentklassen
mit zeilenbasierten Mitteln nach — `sed` schneidet Code-Fences, `grep` zählt Checkboxen über
Abschnittsgrenzen hinweg. In a-check sind das **589 Zeilen Shell** in vier Skripten. Dieselbe
Copy-Drift-Klasse, die bereits `vcs` (aus `adr-immutable-check.sh`) und `targets` (aus
`gate-consistency.sh`) ausgelöst hat: *verteilen statt kopieren*.

**Nicht adopter-spezifisch.** Das Regelwerk schreibt allen Adoptern Dokumentklassen mit
Pflichtabschnitten vor — Slice-Plan (DoD, Closure-Notiz), Anforderung (Happy/Boundary/Negative),
ADR, Review-Report. Wer diese Formen durchsetzen will, braucht heute ein eigenes Skript.

**Vertrag.** Prüft, dass Dokumente einer Klasse (Glob) einen benannten Abschnitt tragen und darin
eine deklarierte Bedingung erfüllen. Hermetisch, kein git, opt-in.

```yaml
structure:
  - files: "docs/plan/planning/done/slice-*.md"   # Dokumentklasse
    require-section: "6. Closure-Notiz"           # muss existieren (Heading-Klartext)
    non-empty: true                               # und Inhalt tragen
    min-sentences: 2                              # Satzzeichen AUSSERHALB von Fences
    forbid-pattern: '^(Lief gut|Alles ok)\.?$'    # Floskel-Filter
  - files: "docs/plan/planning/**/slice-*.md"
    section: "5. DoD"
    max-tasks: 3                                  # Task-Items IM Abschnitt, nicht dateiweit
  - files: "spec/lastenheft.md"
    section-pattern: '^### AC-'                   # Abschnittsklasse per Regex statt Klartext
    require-strong: [Happy, Boundary, Negative]   # benannte Bausteine im Abschnitt
```

| Grund-Code | Bedeutung |
|---|---|
| `section-missing` | geforderter Abschnitt fehlt |
| `section-empty` | Abschnitt vorhanden, ohne Inhalt |
| `section-constraint` | Bedingung im Abschnitt verletzt (Anzahl, Muster, Bausteine) |

**Zwei Eigenschaften, die den Wert ausmachen** — beide sind der Grund, warum die Skript-Variante
fehlerhaft war:

1. **Abschnitts-treu.** `max-tasks` zählt Task-Items *im* Abschnitt. Die Skript-Variante zählte
   dateiweit und musste den Abschnitts-Schnitt selbst nachbauen.
2. **Fence-treu.** `min-sentences` und `forbid-pattern` sehen Code-Blöcke und Inline-Code nicht.
   Genau daran sind zwei Adopter-Sensoren gescheitert (Satzzeichen im Dateinamen als Satzende;
   ein zitiertes Muster als Treffer).

**Paritäts-Mutations-Beleg** — als DoD-Punkt, wie bei `vcs` und `targets`: jede Fixture, die die
vier abgelösten Skripte rot macht, macht auch das Modul rot; jede, die sie grün lässt, lässt auch
das Modul grün. Die Fixture-Menge liegt in a-check vor und kann beigestellt werden.

**Abgrenzung.** Nicht Teil des Antrags: Grandfathering ab einer Slice-Nummer. Das ist über
`exempt-paths` ausdrückbar, wie es `ids`, `matrix` und `codepaths` bereits führen.

---

### CR 2 — `links`: Auflösung unabhängig vom Lifecycle-Verzeichnis

**Anlass.** Das Regelwerk führt den Slice-Lifecycle als Zustandsmaschine über Verzeichnisse
(`open/` → `next/` → `in-progress/` → `done/`) und verlangt, dass der Wechsel ein `git mv`
**ohne Inhaltsänderung** ist. Beides zusammen erzeugt eine Invariante, die `links` heute nicht
prüfen kann: ein relativer Verweis muss aus **jedem** dieser Verzeichnisse auflösen, weil er beim
Wechsel nicht mitgeändert werden darf.

Heute prüft `links` die Auflösung vom **Ist-Ort**. Ein präfixloser Nachbar-Verweis (`roadmap.md`)
ist dort grün und bricht beim nächsten `git mv` — sichtbar erst nach dem Wechsel, an dem man ihn
nicht mehr reparieren darf, ohne die Regel zu verletzen.

**Vertrag.** Eine Option an `links`; ohne sie byte-identisches Verhalten.

```yaml
links:
  resolve-from:                                   # zusätzliche hypothetische Quellorte
    - files: "docs/plan/planning/*/slice-*.md"    # für diese Quelldateien …
      dirs: ["docs/plan/planning/open", "docs/plan/planning/next",
             "docs/plan/planning/in-progress", "docs/plan/planning/done"]
```

| Grund-Code | Bedeutung |
|---|---|
| `link-position-dependent` | Ziel löst vom Ist-Ort auf, aber nicht aus jedem deklarierten Verzeichnis |

**Warum kein eigenes Modul.** Es ist dieselbe Prüfung mit erweiterter Quellort-Menge, nicht eine
neue Frage. Als Option bleibt der Default unberührt.

**Paritäts-Mutations-Beleg:** ein Slice mit präfixlosem Nachbar-Verweis muss rot werden, derselbe
mit zustandsunabhängigem Verweis grün — beide Fixtures liegen in a-check vor.
