# slice-042 — `constructs`-Aufruf-Monopol (P-Rest-Kandidat 1, entgated)

**Review:** [Review-Synthese](../../../reviews/2026-07-25-slice-042-constructs-monopol.md).
**Status:** **done (2026-07-25)** — Entwurf abgenommen (Maintainer-Wort: alle sieben Entscheide
aus §4 wie empfohlen), spec-first umgesetzt, reviewt und nach `main` gemergt. Ergebnis §10,
Closure §11.
**Auslöser:** Maintainer-Review des b-cad-P-Rests (`tools/arch-check.sh`, out-of-repo) am
2026-07-24 mit anschließendem Umsetzungsauftrag — damit ist der Lande-Trigger
„Kandidat 1 auf Maintainer-Wort" aus [slice-025 §5](slice-025-p-rest-generalisierung.md)
eingetreten.
**Bezug:** führt **Kandidat 1** aus [slice-025](slice-025-p-rest-generalisierung.md) aus;
hebt die Scoping-Mechanik von
[AC-FA-RULE-003](../../../../spec/lastenheft.md#ac-fa-rule-003--tech-kapselung-regel-tech-leak)
(`tech`) von Import-Symbolen auf Roh-Text; grenzt gegen die heute port-gebundenen
`forbidden_constructs` aus
[AC-FA-RULE-004](../../../../spec/lastenheft.md#ac-fa-rule-004--port-disziplin-regel-port-impurity)
ab; bewegt sich an der Heuristik-Grenze
[AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze);
Konfigurations-Vertrag
[AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml).
[Roadmap](../in-progress/roadmap.md).

> **Hinweis:** Der Entwurf vergibt **keine** `AC-*`/`ADR-*`-IDs — das geschieht erst im
> Lastenheft-CR bzw. in der ADR selbst (Anlege-Prozess:
> [conventions §Anforderungs-Anlege-Prozess](../../../../harness/conventions.md#anforderungs-anlege-prozess)).
> Die Entscheide §4 sind **abgenommen** (§4.1) und binden die Umsetzung.

---

## 1. Was dieser Slice gegenüber slice-025 neu hat

[slice-025](slice-025-p-rest-generalisierung.md) argumentiert die drei P-Rest-Muster **aus
dem Code**. Dieser Slice hat sie erstmals **mechanisch vermessen** (§2, Fixture gegen
`a-check:dev`) — mit einem Ergebnis, das den Scope **verengt**: eine der beiden dort
skizzierten Scoping-Gestalten verliert ihre Evidenz (§3). Der Rest ist die Ausführung von
Kandidat 1: eine Regel, ein Befund, ein Rückbau-Schritt beim Konsumenten.

Kandidat 2 (fail-closed Import-Allowlist, [slice-025 §4](slice-025-p-rest-generalisierung.md))
bleibt **gated** — §8.

## 2. Vermessung (Fixture-Probe, 2026-07-24)

**Aufbau:** Minimal-Nachbau der b-cad-Struktur (Schichten `model`/`ui`/`plugin_api`/`plugins`,
`resolution: {cpp: {mode: fixed-root, roots: ["src"]}}`, Kanten `plugins → {plugin_api, model}`),
eine Datei `plugins/example/p.cpp` mit je einem Include; Lauf gegen das lokal gebaute
`a-check:dev` (`make build`, Stand `main`), netzlos, read-only. Der Spalte „b-cad-Skript"
liegt der dortige P-Rest zugrunde (Regeln P1/P2/P2c).

| # | Include in `plugins/example/p.cpp` | b-cad-Skript | a-check | Bewertung |
|---|---|---|---|---|
| 1 | `"adapters/ui/y.h"` | P2 rot | `lateral-adapter`, Exit 1 | doppelt abgedeckt |
| 2 | `"../../src/adapters/ui/y.h"` | P2 rot | `lateral-adapter`, Exit 1 | doppelt abgedeckt — die relative Schreibweise entkommt der Auflösung **nicht** (das Layer-Glob-Präfix wird an jeder Segmentgrenze des Kandidaten gematcht, nicht nur am Anfang) |
| 3 | `<adapters/ui/y.h>` | P2c rot | `lateral-adapter`, Exit 1 | doppelt abgedeckt — **P2c ist hier redundant** |
| 4 | `<hexagon/model/m.h>` | P2c rot | 0 Befunde (erlaubte Kante) | P2c-Residuum = reine **Form**-Konvention |
| 5 | `"hexagon/util/u.h"` (in keinem Layer-Glob) | P2 rot | 0 Befunde | **echtes P2-Residuum** (fail-open by design) |
| 6 | `"helper.h"` (modul-lokal) | P2 rot | 0 Befunde | **echtes P2-Residuum** |

**Lesart:** Die Include-Achse ist zu einem größeren Teil abgedeckt als
[slice-025 §2](slice-025-p-rest-generalisierung.md) annahm. Was a-check strukturell **nicht**
sieht, ist (a) das **Aufrufmuster** (P1 — keine Import-Zeile, gar nicht vermessbar über
Includes) und (b) die **Default-verboten-Semantik** für Ziele, die sich auf keine Schicht
auflösen (Proben 5/6).

## 3. Scope-Verengung gegenüber der slice-025-Skizze

1. **`constructs` braucht nur die Monopol-Gestalt** (`adapter:`-Zone wie bei `tech`). Die
   zweite Gestalt `forbid_in:` je Schicht war in
   [slice-025 §3](slice-025-p-rest-generalisierung.md) **allein** durch P2c motiviert. Probe 3
   zeigt: wo P2c Architektur-Gehalt hat (verbotenes Ziel), trägt a-check den Befund bereits als
   Kante — der C++-Extraktor liest `<…>` und `"…"` mit einem Regex, die Kante entsteht also
   unabhängig von der Klammer-Form. Probe 4 zeigt den Rest: eine Form-Konvention über einem
   **erlaubten** Ziel. Sie hat zudem keine Build-Wirkung, die a-check spiegeln müsste — ein
   `<../adapters/…>` löst der Präprozessor gar nicht auf (die Angle-Suche kennt das Verzeichnis
   der einbindenden Datei nicht). ⇒ **`forbid_in` entfällt aus dem CR.**
2. **Korrektur zu [slice-025 §3 „Wirkung"](slice-025-p-rest-generalisierung.md):** der
   Roadmap-Re-Eval-Trigger „C++-quoted-Include-Split" ist damit **nicht** miterledigt; er
   bleibt offen (und ist nach Probe 3/4 auch nicht dringlich).
3. **Kandidat 2 zerfällt in zwei Teile:** sein Residuum sind exakt die Klassen aus Probe 5
   (unlayered) und 6 (modul-lokal) — nicht „alles Unauflösbare" im Allgemeinen. Die erste ist
   flotten-weit und **entgated**, die zweite bleibt b-cad-spezifisch und gated (§8).
4. **Rückbau beim Konsumenten:** b-cads Skript verliert **P1** (durch diese Regel ersetzt);
   **P2** bleibt (Proben 5/6); **P2c** ist zur Streichung **empfohlen** (Argument: Probe 3/4) —
   die Entscheidung liegt beim Konsumenten-Repo, nicht hier.

## 4. Zu entscheiden vor der Umsetzung (ADR-Skizze)

| Frage | Empfehlung | Begründung |
|---|---|---|
| **Schema-Ort** | eigener Top-Level-Block `constructs`, **keine** Generalisierung von `forbidden_constructs` | Verschiedene Semantik: `forbidden_constructs` ist layer-gebunden und mündet in `port-impurity` ([AC-FA-RULE-004](../../../../spec/lastenheft.md#ac-fa-rule-004--port-disziplin-regel-port-impurity)); das Monopol ist **zonen**-gebunden und **scan-weit**. Eine Zusammenlegung müsste beide Scopings in einen Schlüssel pressen. |
| **Namens-Kollision** | interner Modelltyp umbenennen (heute `Constructs` in `internal/hexagon/core/model.go`, trägt die `forbidden_constructs`-Treffer) | ein YAML-`constructs:`-Block belegt sonst denselben Namen wie ein fremd belegter Modelltyp — Verwechslungsrisiko, das [slice-025 §6](slice-025-p-rest-generalisierung.md) bereits markiert |
| **Befund-Name** | `construct-leak` | Parallele zu `tech-leak` bei identischer Scoping-Mechanik (Zone + `composition_root: allow\|forbid`); die Regel-Reihenfolge ist in [SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung) nachzuziehen |
| **Kommentar-Stripping** | **ja** — auf der bestehenden Quell-Vorbereitung aufsetzen (dieselbe, gegen die `forbidden_constructs` heute matcht) | ein `dlopen`-Treffer in einem Kommentar ist falsch-**rot**. Bewusste, auszuweisende Divergenz zur bash-grep-Referenz (grep sieht Kommentare) — die Paritätsprobe (§7) muss den Kommentar-Fall als Divergenz führen, nicht als Fehler |
| **Scan-Menge** | scan-weit über alle `languages`-Dateien; `exclude` greift davor ([ADR-0018](../../adr/0018-exclude-scan-scope.md)/[ADR-0025](../../adr/0025-exclude-verzeichnis-prune.md)) | das Monopol gilt auch für Dateien **ohne** Layer (Probe-5-Klasse) und für die Composition Root — dort nur bei `composition_root: forbid`. **Code-Hinweis:** die heutige Auswertung überspringt Composition-Root-Dateien früh und prüft Konstrukte nur bei Rolle `port` (`internal/hexagon/core/rules.go`); beide Zweige sind zu erweitern |
| **Muster-Sprache** | RE2 wie bei `tech` ([ADR-0015](../../adr/0015-regex-tech-muster.md)), `match: substring\|regex` | keine zweite Muster-Sprache im Vertrag |
| **Graph-Sichtbarkeit** | **kein** Graph-Knoten, aber eine **Legendenzeile** | `--print-graph` ([AC-FA-CLI-002](../../../../spec/lastenheft.md#ac-fa-cli-002--architektur-graph-ausgabe), [ADR-0024](../../adr/0024-print-graph-mermaid.md)) rendert `layers`/`edges`; eine Roh-Text-Regel hat dort keine Kante. Die Legende ist seit [slice-040](../done/slice-040-graph-legende-vertical-slice-regeln.md)/[slice-041](../done/slice-041-graph-legende-layout.md) der etablierte Ort für Nicht-Kanten-Semantik — damit bekommt die „bewusst nicht im Graph"-Aussage eine zitierbare normative Stelle ([SPEC-CLI-002](../../../../spec/spezifikation.md#spec-cli-002--graph-renderer-vertrag)); Lehre aus [slice-033](slice-033-print-mk-graph-target.md) |

### 4.1 Abnahme (2026-07-25)

Alle sieben Entscheide sind **wie empfohlen** abgenommen. Was daraus für die Umsetzung
festgezurrt ist:

- **Eigener Top-Level-Block `constructs`** — `forbidden_constructs` bleibt unangetastet
  (layer-gebunden, mündet weiter in `port-impurity`).
- **Namens-Kollision:** das Modell-Feld `FileImports.Constructs` (heute die
  `forbidden_constructs`-Treffer) wird umbenannt, bevor der neue Block denselben Namen belegt.
- **Befund `construct-leak`**, Scoping-Mechanik identisch zu `tech` (`adapter` als Skalar oder
  Liste, `match: substring|regex`, `composition_root: allow|forbid`, Default `allow`).
- **Kommentar-Stripping ja** — der Roh-Text-Match läuft auf derselben Quell-Vorbereitung wie
  `forbidden_constructs`; der Kommentar-Fall ist in der Paritätsprobe (§7) als **deklarierte
  Divergenz** zur grep-Referenz zu führen, nicht als Fehler.
- **Scan-weit** inkl. layer-loser Dateien und (nur bei `forbid`) der Composition Root.
- **RE2** wie bei `tech`; keine zweite Muster-Sprache.
- **Graph:** kein Knoten, eine **Legendenzeile**.

Zwei Umsetzungs-Details, die die Entscheide nicht präjudizieren, aber aus ihnen folgen:

1. **Treffer-Identität statt Muster-Lookup.** Ein Roh-Text-Treffer trägt den **Index** seines
   `constructs`-Eintrags, nicht bloß das Muster als Zeichenkette. Sonst müsste die Regel-Engine
   den Eintrag über den Mustertext zurücksuchen — bei zwei Einträgen mit gleichem Muster und
   verschiedenen Zonen wäre das mehrdeutig (still falsch-grün für den zweiten Eintrag) und
   erzwänge entweder eine künstliche Duplikat-Verbotsregel im Schema oder eine
   Erst-Treffer-Präzedenz, die einen echten Verstoß verschluckt.
2. **Totalordnung der Befunde.** Zwei `constructs`-Muster können dieselbe **Zeile** derselben
   Datei treffen ⇒ zwei Befunde mit identischem (Pfad, Zeile, Regel). Die heutige Sortierung
   ([SPEC-DET-001](../../../../spec/spezifikation.md#spec-det-001--determinismus-vertrag)) ist
   auf diesem Tripel **nicht total** und läuft über ein instabiles `sort.Slice` — die Reihenfolge
   solcher Geschwister wäre undefiniert
   ([AC-QA-01](../../../../spec/lastenheft.md#ac-qa-01--determinismus)). Die Ordnung bekommt
   deshalb die **Meldung** als letzten Schlüssel. Latent besteht die Lücke schon heute (zwei
   `forbidden_constructs`-Muster in einer Zeile ⇒ zwei `port-impurity`-Befunde); dieser Slice
   schließt sie mit.

## 5. Schema-Skizze

```yaml
constructs:
  - {pattern: '\bdl(m?open|sym|close)\s*\(', match: regex,
     adapter: adapters/plugin, composition_root: forbid}
```

Semantik: Das Muster darf **nur** in der (oder einer der) genannten Zone(n) auftreten; jeder
Treffer außerhalb ist ein Befund `construct-leak` (Datei, Zeile, Muster, erlaubte Zone[n]),
Exit 1. `adapter` als Skalar **oder** Pfad-Liste, `composition_root: allow|forbid` je Eintrag —
identisch zur `tech`-Mechanik seit
[AC-FA-RULE-003](../../../../spec/lastenheft.md#ac-fa-rule-003--tech-kapselung-regel-tech-leak) 0.14.0.

## 6. Anforderungs-Skizze (neue Anforderung im Bereich `RULE`)

- **Happy:** Given ein `constructs`-Eintrag (`dlopen`-Muster, Zone `adapters/plugin`), when eine
  Datei in `adapters/io` das Muster enthält, then Befund `construct-leak` (Datei, Zeile, Muster,
  erlaubte Zone[n]) und Exit 1; die Zone selbst bleibt befundfrei.
- **Boundary:** Given eine Datei, die in einem `languages`-Glob, aber in **keinem** Layer-Glob
  liegt, when sie das Muster enthält, then Befund (die Roh-Text-Prüfung ist scan-weit, nicht
  layer-gebunden). Given `composition_root: forbid`, when die deklarierte Composition Root das
  Muster enthält, then Befund (mit dem `allow`-Default: kein Befund dort). Given ein Treffer
  ausschließlich in einem Kommentar, then **kein** Befund (§4).
- **Negative:** Given ein Eintrag mit leerem/fehlendem `pattern`/`adapter`, unbekanntem
  Schlüsselwert oder ungültigem Regex, when a-check lädt, then Exit 2 (fail-closed — Muster der
  0.14.0-Härtung von [AC-FA-RULE-003](../../../../spec/lastenheft.md#ac-fa-rule-003--tech-kapselung-regel-tech-leak)).
- **Out-of-Scope:** kein Parser — String-Literale bleiben Text-Heuristik
  ([AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze));
  keine RE2-fremden Features (Lookaround/Backreferences); keine Auto-Inferenz von Zonen; **kein**
  `forbid_in` (§3); keine Import-Allowlist (§8).

## 7. Betroffene Module und Belege

| Modul | Änderung |
|---|---|
| `internal/adapter/driven/config` | neuer `constructs`-Block, fail-closed-Validierung (Exit 2) |
| `internal/hexagon/core` | Modelltyp + Regel `construct-leak`; scan-weite Auswertung inkl. layer-loser Dateien und Composition Root; stabile Sortierung unverändert ([AC-QA-01](../../../../spec/lastenheft.md#ac-qa-01--determinismus), [SPEC-DET-001](../../../../spec/spezifikation.md#spec-det-001--determinismus-vertrag)) |
| `internal/adapter/driven/extract` | Muster-Treffer je Datei scan-weit statt layer-gebunden ([SPEC-EXTRACT-001](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion)) |
| `internal/adapter/driven/graph` | Legendenzeile ([SPEC-CLI-002](../../../../spec/spezifikation.md#spec-cli-002--graph-renderer-vertrag)) |
| Doku | Lastenheft-CR, ADR, [SPEC-CONF-001](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema)/[SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung), Benutzerhandbuch, [CHANGELOG](../../../../CHANGELOG.md), Roadmap |

**DoD:** Spec-first-Reihenfolge (Lastenheft-CR mit neuer `RULE`-ID → ADR → Spezifikation → Code
→ Tests); `make gates` **und** `make ci` grün mit echter Ausgabe; **Paritätsprobe** gegen die
P1-grep-Referenz des Konsumenten (gleiche Treffer, gleiche Nicht-Treffer; Kommentar-Fall als
deklarierte Divergenz ausgewiesen); **Fitness-Probe** (injizierter `dlopen`-Aufruf außerhalb der
Zone ⇒ Befund im a-check-Gate statt im Skript, Gegenprobe innerhalb der Zone ⇒ grün);
Rückbau-Vorschlag an den Konsumenten (P1 streichen, P2c empfohlen — §3.4);
Benutzerhandbuch-Currency.

## 8. Kandidat 2: nicht ein Thema, sondern zwei (Maintainer-Nachfrage 2026-07-24)

Die erste Fassung dieses Entwurfs empfahl, Kandidat 2 gated zu lassen — argumentiert aus der
**Konsumenten-Zahl**. Die Nachfrage „warum können das andere Konsumenten nicht gebrauchen?" hat
eine Nachmessung über **alle fünf lokalen `.a-check.yml`-Konsumenten** ausgelöst (Zählung
gescannter Dateien außerhalb aller Layer-Globs; Proben 7–9 im Fixture). Ergebnis: die Empfehlung
war für die eine Hälfte falsch.

### 8.1 Die Fehlerklasse ist flotten-weit (Proben 7–9)

| # | Fall | Ergebnis |
|---|---|---|
| 7 | schichtlose Datei importiert einen Adapter (**Quell**-Seite) | `wrong-direction`, Exit 1 — **kein** fail-open |
| 8 | `domain`-Datei importiert die schichtlose Zone (**Ziel**-Seite) | 0 Befunde — **fail-open** |
| 9 | Gegenprobe: dieselbe Zone als Schicht deklariert | `core-impurity`, Exit 1 |

Die Lücke sitzt also präzise auf der **Ziel-Seite**: ein Import, dessen Ziel sich auf **keine**
deklarierte Schicht auflöst, bleibt unbeurteilt ([SPEC-EXTRACT-001](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion),
bewusste fail-open-Grenze nach [AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)).
Betroffen ist **jeder** Konsument, dessen `languages`-Globs weiter reichen als seine
`layers`-Globs — und das ist keine C++/Plugin-Eigenheit:

| Konsument | gescannt, aber in keinem Layer-Glob |
|---|---|
| a-check (Dogfooding) | **0** — Baum vollständig gedeckt |
| b-cad | **0** — Baum vollständig gedeckt |
| d-check | **0** — Baum vollständig gedeckt |
| d-migrate | ~~der gesamte `test/`-Kotlin-Baum (kein `exclude`)~~ — **falsch, korrigiert 2026-07-25**: d-migrate hat seit dem ersten Commit ein `exclude` (`test/**`, `**/src/test/**`, …); real **0** ungedeckte Dateien. Nachmessung in [slice-043 §2a](../done/slice-043-schicht-abdeckung-sichtbar.md) |
| m-trace | `apps/api/internal/storage/**` und `apps/api/scripts/coverage-overview/**` |

Der m-trace-Fall ist der scharfe: `apps/api/internal/storage` wird von der dortigen Config in
einer `tech`-Regel **als Architektur-Zone geführt** (zulässiger `database/sql`-Halter), hat aber
**keine** Schicht — ein Import aus `hexagon/domain` dorthin wäre exakt der Verstoß, den das Gate
finden soll, und bliebe still. Heute ist die Lücke **latent** (nur `*_test.go` importieren die
Zone, und die sind per `exclude` ohnehin aus dem Scan) — aber sie ist unbemerkt: nichts in der
Ausgabe weist darauf hin, dass ein gescannter Teilbaum ungeprüft bleibt.

### 8.2 Konsequenz — Kandidat 2 aufteilen

- **2a — Schicht-Abdeckung sichtbar/streng machen (entgated, flotten-weite Evidenz).** Ausweisen
  (advisory) bzw. auf Opt-in fail-closed stellen, wenn gescannte Dateien in keinem Layer-Glob
  liegen bzw. Import-Ziele auf keine Schicht auflösen. Deckt Probe 5 + 8 sowie m-trace (**nicht**
  d-migrate — die Zeile oben ist korrigiert, dortige `exclude`-Deklaration übersehen);
  **kehrt keine Beweislast um** und fasst keinen Roh-Text an — es macht nur eine bereits
  bestehende, bewusste fail-open-Grenze **sichtbar**. Der Fix beim Konsumenten ist eine Zeile
  Config (Probe 9), er muss die Lücke nur erfahren.
- **2b — Präfix-Allowlist je Schicht (weiter gated).** Nur Probe 6 (modul-lokale/unauflösbare
  Specifier) braucht die eigentliche Allowlist mit umgekehrter Beweislast. Dafür steht weiterhin
  **nur** b-cad als Evidenz; die Begründung aus
  [slice-025 §4](slice-025-p-rest-generalisierung.md) bleibt unverändert gültig.

**Empfehlung:** 2a als eigenen Folge-Slice aufsetzen (unabhängig von diesem hier — andere Regel,
andere Evidenz, kein gemeinsamer Code-Pfad) — angelegt als
[slice-043](../done/slice-043-schicht-abdeckung-sichtbar.md); 2b gated lassen. Damit schrumpft b-cads Skript nach
diesem Slice auf **eine** Regel (P2), und die Flotte gewinnt eine Fehlerklasse, die mit dem
b-cad-P-Rest gar nichts zu tun hatte.

## 9. Nebenbefunde (nicht Teil dieses Slices)

**9.1 Leeres `tech.pattern` ist ein stiller Never-Match.** `Tech.matches` verlangt
`Pattern != ""`, ein `tech`-Eintrag mit leerem/fehlendem `pattern` meldet also nie — dieselbe
False-Green-Klasse wie der leere `adapter`, den
[AC-FA-RULE-003](../../../../spec/lastenheft.md#ac-fa-rule-003--tech-kapselung-regel-tech-leak)
in 0.14.0 fail-closed gestellt hat (bei `match: regex` bricht ein leeres Muster bereits ab).
Aufgefallen beim Bau der `constructs`-Validierung, die ein leeres `pattern` fail-closed ablehnt
(§6 Negative). Eigener Mikro-CR — die `tech`-Semantik zu verschärfen ist ein Vertragsschnitt, der
nicht in diesen Slice gehört.

**9.2 Leerer Quell-Schicht-Name.** Probe 7 meldet korrekt, aber mit Loch in der Meldung: `wrong-direction:  -> ui (adapters/ui/y.h)`
— für die schichtlose Quelldatei steht an der Stelle des Quell-Schicht-Namens nichts. Rein
kosmetisch (Befund und Exit-Code stimmen), aber nutzersichtbar; kein Vertrag. Aufgenommen in
[slice-043 §5](../done/slice-043-schicht-abdeckung-sichtbar.md) — es ist das quell-seitige Symptom
derselben Abdeckungs-Klasse.

## 10. Umsetzung und Ergebnis (2026-07-25)

**Spec-first geliefert:** Lastenheft **0.22.0** (neu
[AC-FA-RULE-011](../../../../spec/lastenheft.md#ac-fa-rule-011--konstrukt-monopol-regel-construct-leak),
[AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml) um den
Block erweitert) → [ADR-0027](../../adr/0027-constructs-roh-text-monopol.md) `Accepted` →
Spezifikation **0.24.0** ([SPEC-CONF-001](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema),
[SPEC-EXTRACT-001](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion),
[SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung),
[SPEC-DET-001](../../../../spec/spezifikation.md#spec-det-001--determinismus-vertrag),
[SPEC-CLI-002](../../../../spec/spezifikation.md#spec-cli-002--graph-renderer-vertrag)) → Code →
Tests. Benutzerhandbuch 1.33, [CHANGELOG](../../../../CHANGELOG.md) unter `[Unreleased]`.

**Ein Zusatz gegenüber §4**, aus der Umsetzung erzwungen und in
[SPEC-DET-001](../../../../spec/spezifikation.md#spec-det-001--determinismus-vertrag) nachgezogen: die
Befund-Sortierung bekam die **Meldung** als letzten Schlüssel (§4.1 Punkt 2) — ohne sie ist die Ordnung
auf (Pfad, Zeile, Regel) nicht total, und zwei Treffer in einer Zeile lägen in zufälliger Reihenfolge.

### 10.1 Paritätsprobe gegen die grep-Referenz (realer Konsumenten-Baum)

Kopie des b-cad-Baums (read-only, netzlos, Scratchpad), dessen `.a-check.yml` um den einen
`constructs`-Eintrag ergänzt; Referenz ist die dortige Skript-Regel P1
(`grep -rnE '\bdl(m?open|sym|close)[[:space:]]*\(' src plugins | grep -vE '^src/adapters/plugin/'`).

| Probe | grep-Referenz | a-check | Bewertung |
|---|---|---|---|
| unveränderter Baum | 0 Treffer | 0 Befunde, Exit 0 | Parität |
| injizierter `dlsym(`-Aufruf in `plugins/example/example_plugin.cpp:56` | 1 Treffer | `construct-leak`, gleiche Datei **und Zeile**, Exit 1 | Parität |
| injizierter `dlopen(`-Aufruf in `src/main.cpp:406` (Composition Root) | 1 Treffer | `construct-leak` mit `(composition_root: forbid)` | Parität |
| injizierter `dlclose(`-Aufruf **in** `src/adapters/plugin/…` | kein Treffer (Zone) | kein Befund | Parität (Gegenprobe) |
| Kommentar-Zeile `// … dlopen(path) …` in `src/hexagon/model/material_line.h:41` | **1 Treffer** | **kein** Befund | **deklarierte Divergenz** ([ADR-0027](../../adr/0027-constructs-roh-text-monopol.md)) |

Damit ist die Fitness-Probe der DoD auf dem realen Baum erfüllt: der injizierte Aufruf außerhalb der
Zone wird vom a-check-Gate gefunden, die Gegenprobe innerhalb der Zone bleibt grün. Dieselben Fälle
liegen als Fixture-Tests im Repo (`internal/cli`), inklusive der maschinellen Paritätsprobe gegen eine
nachgebaute grep-Referenz.

### 10.1a Review-Fund F-1: die Kommentar-Ausnahme gilt für Python nicht

Das Review vor dem Merge fand eine **zu weit formulierte Anforderung**: die Boundary „Treffer nur
im Kommentar ⇒ kein Befund" stimmt nur für die **C-Syntax-Sprachen**. Die Konstrukt-Erkennung nutzt
bewusst **dieselbe** Quell-Vorbereitung wie die Import-Extraktion — und die lässt **Python**
ungestrippt (`prepSource`, [SPEC-EXTRACT-001](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion)):
eine `/*`-artige Bytefolge in einem Python-String würde sonst echte Importe verschlucken. Real
gemessen (Fixture, `a-check:dev`): derselbe Kommentar-Treffer meldet in `.py`, schweigt in `.cpp`.

**Entschieden: Doku korrigieren, nicht Code.** Ein `#`-Strip allein für diesen Pfad würde ein `#`
im String-Literal mitverschlucken und könnte einen **echten** Treffer verbergen — False-Green ist
die schwerere Fehlerklasse (Lehre aus [ADR-0025](../../adr/0025-exclude-verzeichnis-prune.md)/F-1
dort), und eine zweite, abweichende Quell-Vorbereitung wäre genau die Drift, die
[ADR-0027](../../adr/0027-constructs-roh-text-monopol.md) §4 vermeiden wollte. Die Grenze steht
jetzt ausgewiesen in
[AC-FA-RULE-011](../../../../spec/lastenheft.md#ac-fa-rule-011--konstrukt-monopol-regel-construct-leak)
(Beschreibung + eigene Boundary-Klausel), in
[SPEC-EXTRACT-001](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion) und im
Benutzerhandbuch — plus ein Test, der beide Sprachen gegeneinander festnagelt. Die ADR bleibt
unberührt: sie sagt „dieselbe vorbereitete Quelle", und genau das ist eingetreten.

### 10.2 Regressions-Probe (alle lokalen Konsumenten)

`a-check:dev` gegen b-cad, d-check, d-migrate und m-trace: **je 0 Befunde, Exit 0** — unverändert.
Die Regel ist opt-in; ohne `constructs`-Block ändert sich nichts.

### 10.3 Gates

`make ci` grün: lint 0 issues · Tests grün · Coverage **96,00 %** (Schwelle 90 %) · `arch-check`
(Dogfooding) 0 Befunde · `doc-check` 106 Dateien, 0 Befunde · `gate-consistency` ok ·
`guard-selftest` ok · `image-test` OK.

### 10.4 Rückbau-Vorschlag an den Konsumenten (b-cad, out-of-repo)

Der Vorschlag steht, die Entscheidung liegt drüben:

1. **P1 streichen** und durch den Config-Eintrag ersetzen (Parität in §10.1 belegt):

   ```yaml
   constructs:
     - {pattern: '\bdl(m?open|sym|close)\s*\(', match: regex,
        adapter: src/adapters/plugin, composition_root: forbid}
   ```

   Voraussetzung: ein Release, das die Regel enthält (der aktuelle Pin hat sie noch nicht).
   Ein Unterschied ist auszuweisen: der `grep` meldet auch Kommentar-Treffer, a-check nicht.
2. **P2c streichen** — Empfehlung unverändert aus §3.4 (Argument: Probe 3/4 in §2).
3. **P2 bleibt** — die Klassen aus Probe 5/6 (unauflösbare bzw. modul-lokale Specifier) deckt a-check
   weiterhin nicht; die unlayered-Hälfte davon adressiert
   [slice-043](../done/slice-043-schicht-abdeckung-sichtbar.md).

Danach trägt b-cads `arch-check.sh` genau **eine** Regel.

## 11. Closure-Notiz (2026-07-25)

**Abgeschlossen.** DoD (§7) vollständig erfüllt: Spec-first-Reihenfolge eingehalten, `make gates`
**und** `make ci` mit echter Ausgabe grün, Paritätsprobe gegen die P1-`grep`-Referenz des
Konsumenten bestanden (Kommentar-Fall als deklarierte Divergenz ausgewiesen, §10.1), Fitness-Probe
am realen Baum bestanden, Rückbau-Vorschlag formuliert (§10.4), Benutzerhandbuch-Currency (1.33).

**Was noch aussteht, aber nicht zu diesem Slice gehört:** ein Release, das die Regel ausliefert —
erst danach kann der Konsument seine `grep`-Regel zurückbauen. Der Eintrag liegt in
[CHANGELOG](../../../../CHANGELOG.md) `[Unreleased]`.

### Lerneinträge

1. **Ein Treffer braucht eine Identität, keinen Namen.** Der erste Entwurf hätte den Mustertext als
   Treffer-Kennung genutzt (wie `forbidden_constructs` es tut). Sobald zwei Einträge dasselbe Muster
   mit verschiedenen Zonen tragen, ist das mehrdeutig — und der Ausweg wäre entweder ein
   künstliches Duplikat-Verbot im Schema oder eine Erst-Treffer-Regel, die einen echten Verstoß
   verschluckt. Der **Index** kostet nichts und macht beide Krücken überflüssig.
2. **Eine neue Regel deckt auf, wo eine alte Ordnung nicht total war.** Die Befund-Sortierung
   (Pfad, Zeile, Regel) war auf einem instabilen `sort.Slice` schon vorher nicht total — sichtbar
   wurde es erst, als eine Regel mehrere Befunde je Zeile erzeugen konnte. Wer eine
   Determinismus-Zusage gibt, muss die Ordnung **total** halten, nicht bloß „meistens eindeutig".
3. **Eine geteilte Quell-Vorbereitung erbt auch ihre Ausnahmen** (§10.1a). „Kommentare zählen
   nicht" stimmte für die C-Syntax-Sprachen und für Python nicht — weil Python dort aus gutem Grund
   ungestrippt bleibt. Die Anforderung war zu weit formuliert; korrigiert wurde die **Aussage**,
   nicht das Verhalten, denn die Alternative (ein `#`-Strip) hätte False-Green riskiert.
4. **Die Paritätsprobe gehört an den realen Baum.** Die Fixture-Tests belegen die Semantik; dass
   Muster, Zone und Zeilennummern *dieselben* Stellen treffen wie die abgelöste `grep`-Regel, zeigt
   erst der Lauf gegen den echten Konsumenten-Baum — inklusive des Falls, in dem beide bewusst
   **auseinandergehen**.
