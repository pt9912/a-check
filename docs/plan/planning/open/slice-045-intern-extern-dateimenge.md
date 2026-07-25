# slice-045 — Intern vs. extern über die Dateimenge (Ziel-seitige Abdeckung)

**Status:** open — **entschieden 2026-07-25: jetzt nicht bauen** (§0), mit ausgewiesenen Triggern.
Der Slice bleibt als Entwurf stehen: Messung, Entscheide und Umsetzungs-Fallstrick sind
festgehalten, damit ein späterer Lauf nicht bei null anfängt. Keine Spec-/Code-Änderung.
**Auslöser:** Maintainer-Nachfrage am 2026-07-25 im Anschluss an
[slice-043](../done/slice-043-schicht-abdeckung-sichtbar.md): „87 von 180 Include-Zielen in b-cad
lösen auf keine Schicht auf — und das können wir mit `.a-check.yml` nicht abbilden." Richtig — und
die Prüfung dieser Aussage förderte zutage, dass a-check die nötige Information für einen Teil der
Konsumenten **bereits hält**, sie nur nicht nutzt.
**Bezug:** adressiert die **Ziel**-Seite der Abdeckungs-Lücke, die
[ADR-0029](../../adr/0029-abdeckungs-diagnose-advisory.md) ausdrücklich **ausgeklammert** hat;
berührt die Auflösung [SPEC-EXTRACT-001](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion)/[SPEC-CONF-001](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema)
und die Heuristik-Grenze [AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze).
[Roadmap](../in-progress/roadmap.md).

> **Hinweis:** Entwurf zur Abnahme. Es werden hier **keine** `AC-*`/`ADR-*`-IDs vergeben
> (Anlege-Prozess:
> [conventions §Anforderungs-Anlege-Prozess](../../../../harness/conventions.md#anforderungs-anlege-prozess)).
> Die Entscheide §4 gehören **vor** die Umsetzung.

---

## 0. Abnahme (2026-07-25) — nicht bauen, mit Trigger

| # | Frage | Entscheidung |
|---|---|---|
| 1 | **Scope** | **(c) — jetzt nicht bauen.** Falls ein Trigger (§0.2) feuert: dann **(a)**, nur `fixed-root` |
| 2 | **`module_base` deklarieren?** | **Nein** |
| 3 | **Wirkung** | weder (a) noch (b) als **eigener** Mechanismus. Falls gebaut: **(a) als Anreicherung der bestehenden [ADR-0029](../../adr/0029-abdeckungs-diagnose-advisory.md)-Zeile**; ein **Befund** gehört, wenn überhaupt, in den `strict_coverage`-Slice |
| 4 | **Verhältnis zur Quell-Diagnose** | **eine** Meldung. Zwei Objekte, **eine** Ursache — zwei Kanäle für eine Ursache sind per Konstruktion Doppelmeldung |

### 0.1 Begründung

**Zu 1.** Die Neigung „(a), weil korrekt und kostenlos" war an der falschen Frage gemessen: sie
beantwortet, ob man **kann**, nicht ob es **nützt**. Nach §5.1 ist die Nachfrage (P2/2b) erodiert,
die Messung zeigt **null** reale Fundstellen, und die Information ist über die Quell-Diagnose
bereits sichtbar. Das ist genau die Lage, in der
[ADR-0029](../../adr/0029-abdeckungs-diagnose-advisory.md) selbst `strict_coverage` vertagt hat
(„für eine Strenge, die noch niemand angefragt hat") — dieselbe Elle hier anzulegen ist
**konsistent, nicht zögerlich**.

Zweites Argument: **a-check könnte das Feature nicht dogfooden.** Die Eigen-Config läuft im
`path`-Modus (§2.3); ein `fixed-root`-only-Mechanismus würde von `make arch-check` **nie**
ausgeführt. Ein Mechanismus, den das eigene Gate nicht fährt, hängt allein an Unit-Tests.

**Zu 2.** Das Argument „Deklaration ≠ gelesenes Manifest" trägt gegenüber
[ADR-0002](../../adr/0002-text-heuristische-extraktion.md)/[ADR-0014](../../adr/0014-resolution-roots.md)
— die Hermetik bliebe gewahrt. Es scheitert an anderem: ein neuer Config-Schlüssel **mit
fail-closed-Validierung** für vier Konsumenten, die dafür heute **null** Konfiguration brauchen und
von denen ihn **keiner** angefragt hat. Und er wäre ein **zweiter Ausdruck für dieselbe Wahrheit**,
die `segIndex` bereits präfix-tolerant löst — zwei Quellen, die auseinanderlaufen können
(Schicht-Zuordnung sagt „intern", Dateimengen-Test sagt „extern", weil `module_base` veraltet ist).
Die Suffix-Heuristik aus 1(b) ist damit ebenfalls erledigt: sie reproduziert genau das
**Geister-Match-Risiko**, das [ADR-0029](../../adr/0029-abdeckungs-diagnose-advisory.md)
§Verworfene Alternativen bereits benannt hat.

**Zu 4 — konkrete Gestalt, falls es je gebaut wird:** die bestehende Diagnose-Zeile bekommt einen
Zusatz **je Datei** („… — importiert aus n Datei(en)"). Das schärft das Signal — *ungedeckt und
erreicht* ist etwas anderes als *ungedeckt und tot* — **ohne** zweiten Ausgabekanal, **ohne** neue
Kürzungslogik und **ohne** zweite Determinismus-Sortierung.

### 0.2 Trigger, die die Entscheidung umdrehen

Einer genügt; dann gilt **(a) + 3(a) + 4 gemeinsam**:

1. Ein Konsument, dessen Abdeckungs-Diagnose feuert, fragt **„welcher Import erreicht die Datei
   eigentlich?"** — dann ist die Priorisierung **angefragt** statt vermutet.
2. Eine Nachmessung zeigt bei einem `fixed-root`-Konsumenten **≥ 1** intern-aber-schichtloses Ziel
   (heute: **0**).
3. **`strict_coverage` wird gezogen** — dann fällt der ziel-seitige Test als **Nebenprodukt** an und
   wird dort mitentschieden, nicht in einem eigenen Slice.

## 1. Die Frage

a-check kennt heute genau **eine** Unterscheidung für ein Import-Ziel: *löst es auf ein
deklariertes `layers`-Glob auf — ja oder nein?* Ein `<algorithm>`, ein Qt-Header und ein
vergessenes eigenes Verzeichnis sehen damit **gleich** aus: alle drei „lösen auf keine Schicht
auf". Deshalb muss die Ziel-Seite fail-open bleiben, deshalb konnte
[ADR-0029](../../adr/0029-abdeckungs-diagnose-advisory.md) nur die **Quell**-Seite melden, und
deshalb bleibt b-cads `arch-check.sh`-Regel P2 unersetzt
([slice-042 §3](../done/slice-042-constructs-aufruf-monopol.md)).

**Die Frage dieses Slices:** Kann a-check „repo-intern, aber schichtlos" von „repo-extern"
unterscheiden — und wenn ja, für welche Konsumenten und zu welchem Preis?

## 2. Messung (2026-07-25)

### 2.1 Die Idee: die gescannte Dateimenge ist der Diskriminator

a-check hält bereits einen Index der real gescannten Dateien (`fileIndex` in
`internal/hexagon/core/rules.go`, seit der deklarations-bewussten Mehr-Wurzel-Auflösung
[ADR-0023](../../adr/0023-deklarations-bewusste-mehr-wurzel-aufloesung.md)). Er wird heute **nur**
im `fixed-root`-Pfad mit ≥ 2 Wurzeln konsultiert. Der Test wäre:

> Löst der Kandidat auf eine Datei auf, die ich **gescannt** habe? **Ja** ⇒ repo-intern, aber
> schichtlos (meldenswert). **Nein** ⇒ repo-extern (still).

### 2.2 Ergebnis für `fixed-root` — sauber

b-cad (`resolution: {cpp: {mode: fixed-root, roots: ["src"]}}`), 180 verschiedene Include-Ziele:

| Klasse | Anzahl |
|---|---|
| repo-intern, lösen auf eine Schicht auf | 93 |
| **schichtlos**, davon: entsprechen einer real gescannten Datei | **0** |
| **schichtlos**, davon: keine reale Datei im Scan (Systemheader, Qt, OCC, SQLite) | **87** |

**87 von 87 korrekt als extern erkannt, keine Fehlklassifikation.** Für diese Konsumenten-Klasse
ist die Unterscheidung ohne **jede** neue Config-Zeile möglich.

### 2.3 Ergebnis für `mode: path` — nicht ohne Zusatz-Heuristik

**Vier von sieben** Konsumenten haben **keinen** `resolution`-Block und laufen damit im
Default-Modus `path` (Import == Pfad):

| Konsument | `resolution` | Dateimengen-Test heute möglich? |
|---|---|---|
| b-cad | `cpp: fixed-root, roots: ["src"]` | **ja** |
| d-migrate | `fixed-root` + `package_base` | ja |
| belief-agent | `fixed-root` + `package_base` | ja |
| a-check (Dogfooding) | — (`path`) | **nein** |
| d-check | — (`path`) | **nein** |
| m-trace | — (`path`) | **nein** |
| HexSlice-Go-Beispiel | — (`path`) | **nein** |
| x-wal, u-boot, ai-harness-init | keine `.a-check.yml` | keine Evidenz (noch keine Konsumenten) |

Der Grund ist der **Modul-Präfix**:

```text
Import im Code:  github.com/pt9912/m-trace/apps/api/hexagon/domain
Datei im Index:  apps/api/hexagon/domain/…
```

Das Nachschlagen in der Dateimenge ist **exakt** (`path.Clean` + Endungs-Strip) — der Modul-Präfix
steht davor, es trifft nie. Dass die **Schicht**-Zuordnung trotzdem funktioniert, liegt allein an
`segIndex`: der sucht das Glob-Präfix an **jeder Segmentgrenze**.

## 3. Was daraus folgt

Der Kandidat ist **halb** tragfähig:

- Für `fixed-root`-Konsumenten (3 von 7) ist er heute schon korrekt und kostenlos.
- Für `path`-Konsumenten (4 von 7, darunter a-check selbst) bräuchte der Dateimengen-Test
  dieselbe Segment-Toleranz wie die Schicht-Zuordnung. Der Diskriminator wäre dann *„irgendein
  Suffix des Kandidaten entspricht einem gescannten Verzeichnis"* — und damit **wieder unscharf**:
  ein fremdes Modul, das zufällig auf `…/internal/storage` endet, träfe ebenfalls. Genau die
  Unschärfe, die [slice-043 §3](../done/slice-043-schicht-abdeckung-sichtbar.md) als „nicht sicher
  abgrenzbar" beschrieben hat — für `fixed-root` ist diese Einschätzung **widerlegt**, für `path`
  **bestätigt**.

## 4. Zu entscheiden vor der Umsetzung — **in §0 entschieden**

> Die Tabelle steht als Entscheidungs-Vorlage samt der damaligen Neigung; **abgenommen** ist §0.
> Bei Neigung 1 wich die Entscheidung ab (dort: (a), entschieden: (c)) — die Begründung steht in
> §0.1.

| # | Frage | Optionen | Erste Neigung |
|---|---|---|---|
| 1 | **Scope** | (a) nur `fixed-root`-Konsumenten bedienen, `path` ausgewiesen auslassen · (b) auf `path` ausdehnen mit Segment-Toleranz · (c) gar nicht | **(a)** — sie ist heute korrekt, ohne neue Unschärfe; die Grenze wird ausgewiesen statt geraten |
| 2 | **Modul-Präfix deklarieren?** | ein optionaler Schlüssel (z. B. `module_base: github.com/pt9912/m-trace`) machte den `path`-Fall exakt statt heuristisch | offen — es wäre **neue Config-Fläche**, aber die ehrlichere Variante gegenüber Suffix-Raten. Abzuwägen gegen [ADR-0002](../../adr/0002-text-heuristische-extraktion.md)/[ADR-0014](../../adr/0014-resolution-roots.md): Build-Manifeste (`go.mod`) zu **lesen** ist dort bewusst ausgeschlossen — eine **Deklaration** ist etwas anderes als ein gelesenes Manifest |
| 3 | **Wirkung** | (a) nur Diagnose (wie [ADR-0029](../../adr/0029-abdeckungs-diagnose-advisory.md)) · (b) Befund | **(a)** — dieselbe Linie: erst sichtbar machen, Strenge später und opt-in |
| 4 | **Verhältnis zur Quell-Seiten-Diagnose** | eine gemeinsame Meldung oder zwei getrennte | offen; die Quell-Seite meldet Dateien, die Ziel-Seite meldet Import-Ziele — verschiedene Objekte |

## 5. Abgrenzung — was dieser Slice **nicht** löst

- **P2 / Kandidat 2b — Einschätzung revidiert (2026-07-25, siehe §5.1).** Frühere Fassung dieses
  Abschnitts sagte „löst P2 nicht". Nach der Messung in §5.1 ist das **nicht mehr richtig**: P2s
  Schutzwirkung zerfällt in drei Teile, von denen keiner eine Beweislast-Umkehr braucht. Dieser
  Slice deckt davon den **dritten** ab (Import auf eine schichtlose, aber real gescannte Datei).
- **Keine Rücknahme von [ADR-0029](../../adr/0029-abdeckungs-diagnose-advisory.md).** Die dortige
  Entscheidung „nur Quell-Seite" war für den damaligen Kenntnisstand richtig; dieser Slice liefert
  den Grund, sie für einen Teil der Konsumenten zu erweitern — als Folge-ADR, nicht als Korrektur.
- **Keine Aussage über Projekte ohne `.a-check.yml`** (x-wal, u-boot, ai-harness-init): sie sind
  keine Konsumenten und liefern weder Bedarf noch Gegenevidenz.

### 5.1 Warum Kandidat 2b seine Begründung verliert

**Maintainer-Einwand (2026-07-25):** dass ein Include gar nicht auflöst, ist **Build-Integrität —
Aufgabe des Compilers**, nicht von a-check. Die Prüfung dieses Einwands am realen b-cad-Baum:

**Bestand:** **alle** Quote-Include-Ziele in `plugins/` + `src/plugin_api/` existieren real unter
`src/`; es gibt **keinen** modul-lokalen Fall; der Baum hat **0** ungedeckte Dateien. Der
P2-Zusatz „auflösbar **oder nicht**" beschreibt damit einen Fall, der im Bestand nicht vorkommt.

**Injektionsproben gegen `a-check:dev`:**

| Fall | Wer fängt es |
|---|---|
| existierende `src/`-Datei außerhalb der Allowlist (`adapters/io/…`) | **a-check**: `lateral-adapter`, Exit 1 |
| **neue** Datei unter `src/` in keinem Layer-Glob, von einem Plugin inkludiert | **a-check**: Abdeckungs-Diagnose nennt die Datei ([ADR-0029](../../adr/0029-abdeckungs-diagnose-advisory.md)) — die *Ursache* wird sichtbar, Fix ist eine Config-Zeile |
| Ziel existiert nicht | **Compiler** — der Build bricht ab |

**Schluss:** P2 zerfällt in drei Teile, die drei verschiedene Mechanismen abdecken — Kanten,
Diagnose, Compiler. **Keiner davon braucht die Beweislast-Umkehr**, die Kandidat 2b vorsah. Damit
ist dessen Begründung ([slice-042 §8.2](../done/slice-042-constructs-aufruf-monopol.md),
[slice-025 §4](../done/slice-025-p-rest-generalisierung.md)) weitgehend **erodiert** — nicht
widerlegt (ein Konsument *könnte* Default-verboten wollen), aber ohne verbleibenden Bedarf.

**Der Rest, den dieser Slice adressiert:** bei Fall 2 meldet a-check die **Datei** (advisory), nicht
den **Import**. Wer die Kante selbst als Befund will, braucht den Dateimengen-Test — und zwar genau
dessen `fixed-root`-Hälfte, in der b-cad liegt (§2.2).

**Folge für den Rückbau beim Konsumenten:** b-cads `arch-check.sh` kann nach einem Release, das
`construct-leak` ausliefert, auf **null** Regeln gehen — nicht, wie in
[slice-042 §10.4](../done/slice-042-constructs-aufruf-monopol.md) angenommen, auf eine. Die
Entscheidung liegt beim Konsumenten-Repo.

## 6. Betroffene Module (bei Umsetzung)

| Modul | Änderung |
|---|---|
| `internal/hexagon/core` | Dateimengen-Test im Auflösungs-Pfad (heute nur `fixed-root` mit ≥ 2 Wurzeln); Ermittlung der intern-aber-schichtlosen Ziele, stabil sortiert |
| `internal/cli` | Ausgabe (analog der Quell-Seiten-Diagnose) |
| `internal/adapter/driven/config` | **nur bei Entscheid 2** ein optionaler Modul-Präfix-Schlüssel + fail-closed-Validierung |
| Doku | Folge-ADR, [SPEC-EXTRACT-001](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion)/[SPEC-CLI-001](../../../../spec/spezifikation.md#spec-cli-001--aufruf-scan-wurzel-und-exit-codes), Benutzerhandbuch, [CHANGELOG](../../../../CHANGELOG.md), Roadmap |

### 6.1 Umsetzungs-Fallstrick — `denotes()` ist als Diskriminator **unbrauchbar**

Die naheliegende Wiederverwendung wäre `fileIndex.denotes()` bzw. `strength(cand) > 0`. **Das wäre
falsch.** `strength` hat drei Stufen, und **Stufe 1 prüft nur, ob das Elternverzeichnis des
Kandidaten existiert**:

```go
if _, ok := idx.dirs[path.Dir(c)]; ok { return 1 }
```

`#include <QWidget>` wird unter b-cads `roots: ["src"]` zum Kandidaten `src/QWidget`; dessen
Elternverzeichnis `src` liegt in `idx.dirs` — Ergebnis **„intern"**. Am realen Baum nachgemessen:

| Diskriminator | wertet die 87 externen Ziele als „intern" |
|---|---|
| `denotes()` / `strength(cand) > 0` (Stufe 1 genügt) | **87 von 87** — komplett falsch |
| `strength(cand) >= 2` (Datei existiert exakt) | **0 von 87** — korrekt |

Die Messung in §2.2 hat exakt gematcht, entspricht also **Stufe 2**. Eine Umsetzung muss den Test
auf **`strength(cand) >= 2`** festnageln — mit Test **und** Kommentar, warum Stufe 1 hier nicht
gilt. Sonst kippt der Mechanismus von „87 korrekt extern" auf „87 falsch intern", und zwar
**still**.

**DoD (bei Ausarbeitung):** spec-first (Folge-ADR → Spezifikation → Code → Tests); **Review-Synthese
unter [`docs/reviews/`](../../../reviews/)** nach Regelwerk Modul 10; `make gates` **und** `make ci`
grün mit echter Ausgabe; **Konsumenten-Probe über alle Konfigurationen** — die
`fixed-root`-Konsumenten dürfen **keine** Fehlklassifikation zeigen (b-cad: 87 extern, 0 intern),
die `path`-Konsumenten müssen sich entsprechend Entscheid 1 verhalten; Benutzerhandbuch-Currency.

## 7. Closure-Notiz

_(beim Abschluss.)_
