# slice-042 — `constructs`-Aufruf-Monopol (P-Rest-Kandidat 1, entgated)

**Status:** open — **Entwurf zur Abnahme** (spec-first; noch keine Spec-/Code-Änderung).
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

> **Hinweis:** Entwurf zur Abnahme. Es werden hier **keine** `AC-*`/`ADR-*`-IDs vergeben —
> das geschieht erst im Lastenheft-CR bzw. in der ADR selbst (Anlege-Prozess:
> [conventions §Anforderungs-Anlege-Prozess](../../../../harness/conventions.md#anforderungs-anlege-prozess)).
> Die Entscheide §4 gehören **vor** die Umsetzung.

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
| **Graph-Sichtbarkeit** | **kein** Graph-Knoten, aber eine **Legendenzeile** | `--print-graph` ([AC-FA-CLI-002](../../../../spec/lastenheft.md#ac-fa-cli-002--architektur-graph-ausgabe), [ADR-0024](../../adr/0024-print-graph-mermaid.md)) rendert `layers`/`edges`; eine Roh-Text-Regel hat dort keine Kante. Die Legende ist seit [slice-040](../done/slice-040-graph-legende-vertical-slice-regeln.md)/[slice-041](../done/slice-041-graph-legende-layout.md) der etablierte Ort für Nicht-Kanten-Semantik — damit bekommt die „bewusst nicht im Graph"-Aussage eine zitierbare normative Stelle ([SPEC-CLI-002](../../../../spec/spezifikation.md#spec-cli-002--graph-renderer-vertrag)); Lehre aus [slice-033](../done/slice-033-print-mk-graph-target.md) |

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
| d-migrate | der gesamte `test/`-Kotlin-Baum (kein `exclude`) |
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
  liegen bzw. Import-Ziele auf keine Schicht auflösen. Deckt Probe 5 + 8, m-trace und d-migrate;
  **kehrt keine Beweislast um** und fasst keinen Roh-Text an — es macht nur eine bereits
  bestehende, bewusste fail-open-Grenze **sichtbar**. Der Fix beim Konsumenten ist eine Zeile
  Config (Probe 9), er muss die Lücke nur erfahren.
- **2b — Präfix-Allowlist je Schicht (weiter gated).** Nur Probe 6 (modul-lokale/unauflösbare
  Specifier) braucht die eigentliche Allowlist mit umgekehrter Beweislast. Dafür steht weiterhin
  **nur** b-cad als Evidenz; die Begründung aus
  [slice-025 §4](slice-025-p-rest-generalisierung.md) bleibt unverändert gültig.

**Empfehlung:** 2a als eigenen Folge-Slice aufsetzen (unabhängig von diesem hier — andere Regel,
andere Evidenz, kein gemeinsamer Code-Pfad); 2b gated lassen. Damit schrumpft b-cads Skript nach
diesem Slice auf **eine** Regel (P2), und die Flotte gewinnt eine Fehlerklasse, die mit dem
b-cad-P-Rest gar nichts zu tun hatte.

## 9. Nebenbefund (nicht Teil dieses Slices)

Probe 7 meldet korrekt, aber mit Loch in der Meldung: `wrong-direction:  -> ui (adapters/ui/y.h)`
— für die schichtlose Quelldatei steht an der Stelle des Quell-Schicht-Namens nichts. Rein
kosmetisch (Befund und Exit-Code stimmen), aber nutzersichtbar; eigener Mikro-CR, kein Vertrag.
