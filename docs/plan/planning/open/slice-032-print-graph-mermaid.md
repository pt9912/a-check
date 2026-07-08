# slice-032 — `--print-graph`: Architektur-Graph (Mermaid) aus `.a-check.yml`

**Status:** open (Entwurf **abgenommen 2026-07-08** — spec-first; Umsetzung folgt als eigener Slice, noch kein Code).
**Typ:** Feature (Inspektions-/Visualisierungs-Ausgabe), maintainer-initiiert (UX/Onboarding).
**Bezug:** neue Anforderung im Bereich `CLI` (Schema-Konvention, siehe
[MR-002](../../../../harness/conventions.md#mr-002--id-schema-mit-bereichskürzeln-ab-initialer-fassung));
Geschwister der `--print-*`-Familie aus
[AC-FA-DIST-001](../../../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk)
und der Aufruf-/Exit-Code-Anforderung
[AC-FA-CLI-001](../../../../spec/lastenheft.md#ac-fa-cli-001--aufruf-scan-wurzel-und-exit-codes);
liest das Schema aus [SPEC-CONF-001](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema).
Kernverträge: [AC-QA-01](../../../../spec/lastenheft.md#ac-qa-01--determinismus) (byte-identische Ausgabe)
und [AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)
(read-only; der Graph zeigt die **deklarierte** Config, keine Semantik-Behauptung).
[Roadmap](../in-progress/roadmap.md).

## 1. Motivation

`.a-check.yml` **ist bereits ein Graph**: `layers` sind Knoten, `edges` gerichtete Kanten,
`role`/`direction` Typ/Gruppierung, `composition_root`/`adapter_sink`/`tech` Sonderknoten bzw.
Annotationen. Heute muss man das YAML im Kopf zu einem Bild zusammensetzen. Ein
`a-check --print-graph` rendert es als **Mermaid-Flowchart** auf stdout:

- **Onboarding/Review:** die deklarierte Architektur-Absicht auf einen Blick; ein PR, der `edges` ändert,
  zeigt die Absicht visuell.
- **Drift sichtbar:** die gezeichnete Absicht neben den realen Befunden macht Config-Fehler auffällig.
- **Dogfooding + Docs:** a-check zeichnet seine eigene Architektur; die Repo-Doku nutzt bereits Mermaid
  (Roadmap-Abhängigkeitsgraph) — die Ausgabe ist direkt einbettbar.
- **Kein Scan nötig:** rein aus `.a-check.yml`, deterministisch, read-only — passt exakt zur bestehenden
  `--print-config`/`--print-mk`-Familie (`internal/cli/cli.go`, vor dem Scan dispatcht).

## 2. Die Abbildung (Config → Mermaid)

| `.a-check.yml` | Mermaid |
|---|---|
| `layers` (Name) | Knoten mit **interner ID** (`L0`…); der Rohname nur als **maskiertes Label** (§3 Punkt 5), optional Glob als Unterzeile |
| `role` domain/app/port/adapter | `classDef`-Farbe je **effektiver** Rolle: explizites `role:` oder dieselbe Namens-Inferenz wie die Regel-Engine (§3 Punkt 7) |
| `edges` `{from, to}` | `Lᵢ --> Lⱼ` (interne Layer-IDs, §3 Punkt 5 — nie Roh-Namen in Kanten); unbekannte Endpunkte als Dangling-Knoten `Dᵢ` (§3 Punkt 6) |
| `allow` `{from, to}` | gestrichelte Kante `Lᵢ -.-> Lⱼ` mit `allow`-Label (abgesetzte deklarierte Sonderkante); unbekannte Endpunkte ebenfalls `Dᵢ` |
| `composition_root` | ein isolierter Notiz-/Legenden-Knoten `C0` mit allen Globs als maskierte Labels, **ohne** gezeichnete Verdrahtungs-Kanten (§3 Punkt 8) |
| `adapter_sink` | isolierter Ausnahme-/Notizknoten `S0` mit maskiertem Pfadfragment, **keine** Layer-Klasse und **keine** gezeichneten Kanten (§3 Punkt 9) |
| `direction` driving/driven | **§5-Entscheid:** Subgraph-Gruppierung in V1 |
| `tech` `{pattern, adapter}` | **§5-Entscheid:** v2+ |
| implizite Regeln (Kern-Reinheit, `lateral-adapter`, `port-direction`) | **Legende/Notiz**, keine Kante |

Der Graph zeigt die **deklarierten Kanten** (`edges`/`allow`) + Rollen. Diese Kanten sind
Config-Deklarationen, kein Beweis, dass ein realer Import im Scan erlaubt wäre: kategorische
Rollen-/Richtungs-Constraints (`core-impurity`, `lateral-adapter`,
`port-direction-mismatch`) bleiben vorrangig und können eine gezeichnete Deklaration trotzdem
rot machen. Diese Constraints sind keine Kanten — sie erscheinen als Legende, nicht als Kluft im
Bild ([AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze):
ehrliche Ausgabe, keine Semantik-Behauptung über den realen Code).

## 3. Design

**Neuer CLI-Modus `a-check --print-graph [pfad]`** (Default-`pfad` `/src`, wie der Scan):

1. **Dispatch** in `internal/cli/cli.go` als `fs.Bool("print-graph", …)`, **vor** dem Scan (Muster von
   `--print-config`/`--print-mk`). Für `--print-graph` gilt nach `fs.Parse`: höchstens **ein**
   Positionsargument (`pfad`); zusätzliche Restargumente im neuen Modus sind Nutzungsfehler
   (**Exit 2**). Das ist wichtig, weil Go-`flag` nach dem ersten Positionsargument nicht weiter als
   Flag parst:
   `a-check --print-graph <pfad> --bogus` darf nicht still als zweites Positionsargument ignoriert
   werden. Der Modus liest **nur** `pfad/.a-check.yml` — **kein** Quell-Scan. Eine globale
   Restargument-Härtung für bestehende Scan-/Print-Modi ist out-of-scope dieses Slices.
2. **Config laden + validieren:** `config.New().Load()` (Schema) **und** die Sprach-Backend-Prüfung des
   Extraktions-Adapters (`languages`-Keys gegen die Registry,
   [SPEC-EXTRACT-001](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion)) als
   **validation-only**-Aufruf **ohne** Datei-Walk. Jeder Fehler — unbekannter Schlüssel, ungültige
   `role`/`direction`, ungültige `tech`-/`resolution`-/`exclude`-Deklaration, **unbekannte Sprache**
   (`languages: {ruby: …}`) — bricht mit **Exit 2** ab. Der Modus garantiert damit
   **load-time/config-validation-Parität** zum Scan
   ([SPEC-CLI-001](../../../../spec/spezifikation.md#spec-cli-001--aufruf-scan-wurzel-und-exit-codes)),
   aber bewusst **keine** Parität für scanzeitige Fehler, die reale Dateien oder Deklarationen brauchen
   (z. B. die mehrwurzelige, dateimengen-/deklarationsabhängige Resolution-Mehrdeutigkeit aus
   [SPEC-CONF-001](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema)). Solche
   Fehler bleiben Aufgabe des echten Scans; `--print-graph` liest keine Quellen und behauptet sie nicht.
   **[High-Fix]:** die Sprach-Validierung liegt heute im **Extraktor** (`extract.checkLanguages`; die
   Backend-Registry ist dort die Single Source,
   [SPEC-EXTRACT-001](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion)), **nicht**
   in `config.Load` — `config.Load` allein ließe `languages: {ruby: …}` durch und lieferte ein Diagramm,
   wo der Scan Exit 2 gibt. Sie **in** `config.Load` zu zentralisieren wäre in a-checks eigenem Hexagon
   eine **laterale Adapter-Kante** (config → extract), die `arch-check` (zu Recht) meldet. Darum bekommt
   der Extraktions-Adapter einen **validation-only**-Einstieg. **Verankerung [Medium-Port-Fix]:** das
   driven-Port-Interface `port.ExtractionPort` (heute **nur** `Extract`, in
   `internal/hexagon/port/port.go`) wird um `Validate(m core.Model) error` erweitert; der Adapter exponiert sein bestehendes `checkLanguages`
   als `Validate`, `Extract` ruft es intern (keine Doppel-Logik), und `cli.go` programmiert weiter
   gegen das **Port-Interface** (kein Umstieg auf den konkreten Adapter-Typ). Der Scan-Pfad nutzt
   denselben Check via `Extract`.
3. **Rendern:** ein neuer Präsentations-Renderer bildet `core.Model` → Mermaid-String ab, **pur**
   (kein I/O), und schreibt ihn nach stdout. **Read-only**, Exit 0.
4. **Determinismus** ([AC-QA-01](../../../../spec/lastenheft.md#ac-qa-01--determinismus)): Knoten
   **stabil sortiert** (nach Layer-Name), Kanten stabil sortiert (nach `from`, dann `to`), `classDef`
   in fester Reihenfolge — **byte-identische** Ausgabe bei identischer Config, unabhängig von der
   internen Map-/Slice-Ordnung des Modells.
5. **Mermaid-Sicherheit [Medium-Fix]:** Layer-Namen sind **freie YAML-Map-Keys** (`config.go` — Punkte,
   Slashes, Bindestriche, Quotes möglich) und taugen **nicht** als Mermaid-Knoten-IDs. Der Renderer
   erzeugt **stabile interne IDs** (`L0`, `L1`, … in der Sortier-Reihenfolge) und gibt den Rohnamen
   **nur als escaped Label** aus (`L0["…"]`); Kanten referenzieren die internen IDs.
   Konkreter Escaping-Vertrag: nutzergesteuerter Text wird vor der Ausgabe einzeilig normalisiert
   (`\r`, `\n`, `\t` → Leerzeichen; übrige Steuerzeichen → `?`) und dann in fester Reihenfolge für
   Mermaid-/HTML-Delimiter kodiert: `&`→`&amp;`, `"`→`&quot;`, `<`→`&lt;`, `>`→`&gt;`,
   `[`→`&#91;`, `]`→`&#93;`, `|`→`&#124;`, `\`→`&#92;`, `` ` ``→`&#96;`. Renderer-eigene
   Struktur (`["…"]`, feste `<br/>`-Trenner für Unterzeilen) wird erst danach um den kodierten
   Nutztext gelegt; Nutztext erzeugt nie Mermaid-Syntax. Diese Regel gilt für **alle**
   nutzergesteuerten Labels, nicht nur für Layer: `composition_root`-Globs, `adapter_sink`,
   Legenden-/Notiztexte werden ebenfalls nur escaped ausgegeben und nie als Mermaid-ID verwendet.
   `tech` wird in V1 bewusst nicht gerendert.
6. **Dangling-Edge-Policy [Medium-Dangling-Fix]:** `config.Load` übernimmt `edges`/`allow` heute ohne
   Endpunkt-Prüfung gegen `layers`; der Scan behandelt solche Kanten faktisch als inert. `--print-graph`
   führt deshalb **keine neue** Config-Strenge ein: ein unbekannter `from`/`to`-Endpunkt wird als
   deterministischer Dangling-Knoten gerendert (`D0`, `D1`, …, sortiert nach Rohname, eigenes
   `classDef dangling`). Das Diagramm bleibt ehrlich über die deklarierte Config und macht den Tippfehler
   sichtbar, ohne einen Scan-kompatiblen Config-Stand plötzlich mit Exit 2 abzulehnen. Leere
   Endpunkt-Namen werden genauso als Dangling-Labels sichtbar, solange sie vom bestehenden Loader
   akzeptiert werden.
7. **Effektive Rollen [Medium-Role-Fix]:** die Rollen-Farbe folgt nicht nur dem rohen `Layer.Role`, sondern
   der **effektiven** Regel-Engine-Rolle: explizites `role:` hat Vorrang, sonst gilt dieselbe Inferenz wie
   im Kern (`core`→`domain`, `ports`→`port`, `adapters`→`adapter`, `application`/`app`→`app`). Damit
   klassische Configs ohne `role:` — inklusive a-checks eigener `.a-check.yml` — korrekt gefärbt werden.
   Zur Drift-Vermeidung wird der Resolver aus dem Kern exportiert/geteilt (z. B. `core.EffectiveRole`)
   und von Regel-Engine **und** Renderer genutzt; der Graph-Adapter kopiert die Inferenz nicht.
8. **Composition-Root-Darstellung [Medium-Composition-Fix]:** `composition_root` ist eine Glob-Liste, kein
   Layer und kein wirkliches Edge-Set. v1 rendert deshalb genau **einen** isolierten Notizknoten `C0`
   (`classDef exempt`, gestrichelter Rahmen) mit den escaped Globs in deklarierter Reihenfolge und
   zeichnet **keine** Kanten von/zu allen Layern. Die Verdrahtungs-Ausnahme wird in der Legende benannt,
   nicht als tatsächliche Mermaid-Verbindung überzeichnet.
9. **Adapter-Sink-Darstellung [Medium-Sink-Fix]:** `adapter_sink` ist in der Regel-Engine ein
   Pfadfragment für die kategorische `lateral-adapter`-Ausnahme, **kein** Layer und keine deklarierte
   Kante. v1 rendert deshalb höchstens einen isolierten Ausnahme-/Notizknoten `S0` (`classDef exempt`)
   mit dem escaped Pfadfragment und zeichnet **keine** Kanten zu Adapter-Layern. Die Legende benennt:
   „Adapter-Sink: Imports, deren aufgelöster Kandidat dieses Fragment enthält, sind von
   `lateral-adapter` ausgenommen." Fehlt `adapter_sink`, entsteht kein leerer Sonderknoten.

**Modul-Platzierung (a-checks eigenes Hexagon, `arch-check` muss grün bleiben):** entschieden ist ein
eigener Präsentationsadapter `internal/adapter/driven/graph/`, der ein neues driven Port-Interface
`port.GraphPort` implementiert (z. B. `Render(m core.Model) string`). Der Renderer ist pur (kein I/O),
importiert nur `core`, und `cli.go` (Composition Root,
[ARC-006](../../../../spec/architecture.md)) programmiert gegen das Port-Interface:
`config.Load → extract.Validate → graph.Render → stdout`. Der bestehende `report`-Adapter bleibt für
Befund-Reports zuständig; kein direkter Umstieg auf einen konkreten Graph-Adapter-Typ im CLI-Pfad.

**Format:** Mermaid (`flowchart TB`), die Repo-Konvention. Andere Formate (DOT/Graphviz) sind
out-of-scope (§7). `--print-graph` bleibt ein **bool** (`fs.Bool`); ein weiteres Format käme später
**additiv über ein separates `--graph-format`-Flag**, **nicht** durch `--print-graph=<fmt>` [Medium-Flag-Fix]:
ein Wechsel auf `flag.String` bräche die heutige **bare** Form `--print-graph` (Go-`flag` verlangt dann
einen Wert). So bleibt die v1-Form dauerhaft gültig.

## 4. Geplanter Umfang (nach Abnahme)

1. **Spec — Lastenheft:** eine **neue `CLI`-Anforderung** (Arbeitstitel „Architektur-Graph-Ausgabe";
   ID-Vergabe mit Versions-Bump + Historie erst im Lastenheft-CR) — **entschieden** als eigene
   `CLI`-Anforderung, **nicht** als `--print-*`-Erweiterung von
   [AC-FA-DIST-001](../../../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk)
   (Inspektion ≠ Distribution; löst [Medium-AC-Verortung], konsistent mit dem Bezug-Kopf):
   `--print-graph` gibt aus `.a-check.yml` ein Mermaid-Flowchart aus (read-only,
   deterministisch); ladezeitige Config-Fehler (**inkl. unbekannter Sprache**)/unbekanntes Flag → Exit 2;
   scanzeitige, dateimengenabhängige Resolution-Fehler sind out-of-scope des no-scan-Modus. Drei AK (Happy/Boundary/Negative) +
   Out-of-Scope, Version-Bump + Historie
   ([conventions §Anforderungs-Anlege-Prozess](../../../../harness/conventions.md#anforderungs-anlege-prozess)).
   **AC-Änderung nur im Lastenheft.**
2. **ADR (bei Abnahme entschieden: ja, klein):** Format-/Umfang-Entscheidung (Mermaid; deklarierte Kanten +
   Rollen + Sonderknoten; implizite Regeln als Legende; Determinismus-Ordnung). `Schärft:` aufwärts auf
   die einschlägige `SPEC-*`-Stelle; ADR-Index; slice-token-frei
   ([MR-005](../../../../harness/conventions.md#mr-005--referenzmatrix-intra-spec-richtung--adrslice-disziplin-d-check-angleichung)).
3. **Spec — Spezifikation:** [SPEC-CLI-001](../../../../spec/spezifikation.md#spec-cli-001--aufruf-scan-wurzel-und-exit-codes) (Exit-Codes, der neue Modus) + der Renderer-Vertrag
   (Model → Mermaid, stabile Ordnung) präzisieren.
4. **Spec — Architektur-Sicht:** [spec/architecture.md](../../../../spec/architecture.md) nachziehen
   (Version + Historie): `ExtractionPort`/[ARC-003](../../../../spec/architecture.md) um den validation-only Einstieg ergänzen;
   `GraphPort` in [ARC-002](../../../../spec/architecture.md) aufnehmen; den neuen `graph`-Adapter als eigene
   Präsentationskomponente mit eigener Architektur-Kennung sichtbar machen; die Sequenz um den
   no-scan-`--print-graph`-Pfad erweitern. Die Architektur bleibt dabei sprach-/meilensteinfrei und
   referenziert keine Slice-ID.
5. **Code:** `port.ExtractionPort` um `Validate(m core.Model) error` erweitert; `port.GraphPort`
   (`Render(m core.Model) string`) ergänzt; der Extraktions-Adapter
   exponiert `checkLanguages` als `Validate` (ohne Walk) und ruft es in `Extract` (keine Doppel-Logik);
   `cli.go` Flag (`fs.Bool("print-graph")`) + Dispatch (`config.Load` → `Validate` → `GraphPort.Render`);
   neuer `graph`-Adapter mit **stabilen internen IDs + escaped Labels**;
   `direction`-Gruppierung als stabile Mermaid-Subgraphs für `driving`/`driven` (Layer ohne
   `direction` bleiben außerhalb dieser Subgraphs, behalten aber Rolle/Kanten);
   gemeinsam genutzter Resolver für effektive Rollen (`core`/`ports`/`adapters`/`application`-Inferenz);
   `--help`/Usage ergänzt; Restargumente nach dem optionalen `pfad` brechen im `--print-graph`-Modus
   mit Exit 2.
6. **Tests:** E2E (`--print-graph <fixture>` → erwartetes Mermaid, Exit 0; ungültige Config → Exit 2;
   **unbekannte Sprache** `languages: {ruby: …}` → Exit 2 [High-Fix]; unbekanntes Flag → Exit 2;
   **zusätzliches Restargument/Flag nach dem Pfad** (`--print-graph <fixture> --bogus`) → Exit 2;
   **read-only**: schreibt nichts) + **Mermaid-unsafe Layer-Namen** (Punkte/Slashes/Bindestriche/Quotes
   plus adversarisch `]`, `|`, `\`, Backtick, `<`, `>`, `&`, CR/LF/TAB) → gültiges Mermaid
   (interne IDs, escaped Labels, kein Syntax-Ausbruch) [Medium-Fix] + **Mermaid-unsafe Sonderlabels**
   (`composition_root`, `adapter_sink` mit denselben adversarischen Zeichen) → ebenfalls gültiges Mermaid +
   **`allow`-Kante** → gestrichelte, gelabelte Kante im erwarteten Output +
   **`direction`-Gruppierung** → `driving`-/`driven`-Layer erscheinen in stabil sortierten Mermaid-Subgraphs,
   Layer ohne `direction` bleiben außerhalb der Direction-Subgraphs; Kanten dürfen Subgraph-Grenzen queren +
   **klassische Layer-Namen ohne `role:`** → effektive Rollen-Farben via Kern-Inferenz +
   **unbekannte `edges`/`allow`-Endpunkte** → deterministische Dangling-Knoten statt Absturz/Exit 2
   [Medium-Dangling-Fix] + Determinismus (zwei Läufe byte-identisch) + Unit für den Renderer
   (effektive Rollen-Farben, `direction`-Subgraphs, `allow`-Kanten, Composition-Root-Notizknoten,
   Dangling-Knoten, ID/Label-Trennung, Sortier-Stabilität).
   Dogfooding-Probe: `--print-graph` auf a-checks eigene `.a-check.yml` liefert stabil den erwarteten
   Graph.
7. **Benutzerhandbuch:** neuer Abschnitt „Architektur visualisieren (`--print-graph`)" mit Beispiel-Output;
   Currency-Bump. `image-test` deckt den neuen Modus (analog `--print-config`-Dekodierbarkeit).
8. **Gates:** `make gates` + `make ci` + `make trace-check`.

## 5. Entscheide & offene Fragen

**Entschieden (im Slice fixiert, damit CR/Spec/Tests nicht auseinanderlaufen):**

1. **AC-Verortung → neue `CLI`-Anforderung** (nicht [AC-FA-DIST-001](../../../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk) erweitern) — Inspektion ≠
   Distribution; konsistent mit dem Bezug-Kopf (löst [Medium-AC-Verortung]).
2. **Validierungs-Konsistenz → Extraktor-`Validate` (nicht zentralisieren)** — `config.Load` + der
   validation-only Sprach-Check des Extraktors, verankert als Erweiterung von `port.ExtractionPort` um
   `Validate(m core.Model) error` (§3 Punkt 2, [High-Fix]/[Medium-Port-Fix]); Zentralisierung in
   `config.Load` wäre eine laterale Adapter-Kante. Diese Parität ist ausdrücklich auf ladezeitige
   Config-Validierung begrenzt; scanzeitige, dateimengenabhängige Resolution-Mehrdeutigkeiten bleiben
   dem echten Scan vorbehalten.
3. **Mermaid-Sicherheit → interne IDs + escaped Labels für alle nutzergesteuerten Labels**
   (§3 Punkt 5, [Medium-Fix]).
4. **Unbekannte Edge-Endpunkte → Dangling-Knoten, kein neuer Config-Exit-2**
   (§3 Punkt 6, [Medium-Dangling-Fix]).
5. **Rollen-Farben → effektive Rollen** (explizites `role:` oder Kern-Inferenz; geteilter Resolver,
   §3 Punkt 7, [Medium-Role-Fix]).
6. **Composition Root → isolierter Notizknoten, keine All-to-All-Kanten**
   (§3 Punkt 8, [Medium-Composition-Fix]).
7. **Adapter Sink → isolierter Ausnahme-/Notizknoten, keine Layer-/Kanten-Semantik**
   (§3 Punkt 9, [Medium-Sink-Fix]).
8. **Modul-Platzierung → eigener `graph`-Adapter + neues `port.GraphPort`**
   (§3 Modul-Platzierung; `report` bleibt Befund-Report).
9. **Render-Umfang v1 → `direction` aktiv, `tech` deferred**
   (`direction` (`driving`/`driven`) wird als strukturelles Feature in v1 gerendert; `tech` bleibt
   bewusst v1-scope außen vor).

**Bei Abnahme entschieden (2026-07-08):**

1. **ADR ja/nein → ja, klein.** Der umsetzende Slice legt einen kleinen ADR an (Format Mermaid;
   Render-Umfang: deklarierte Kanten + Rollen + Sonderknoten; „implizite Regeln als Legende";
   Determinismus-Ordnung) — neuer Ausgabe-Vertrag mit echten Wahl-Punkten, konsistent mit der
   ADR-je-Feature-Praxis des Repos. `Schärft:` aufwärts auf die einschlägige `SPEC-*`-Stelle,
   ADR-Index, slice-token-frei (§4 Punkt 2).
2. **Render-Umfang v2+ → als Folge-Slice bestätigt.** `tech` als visuelle Notiz/Badge
   (inkl. `tech.pattern`) samt zusätzlicher Testerweiterung bleibt bewusst außerhalb v1 und wird
   in einem eigenen Folge-Slice behandelt.
3. **Format v1 → Mermaid als einziges Format bestätigt.** Ein späteres Format kommt **additiv über
   ein separates `--graph-format`-Flag**, nicht durch Änderung von `--print-graph` (bleibt `fs.Bool`,
   §3, [Medium-Flag-Fix]).

## 6. Akzeptanzkriterien (Fitness-Function, als Tests)

Referenz-Fixture (`.a-check.yml`): `core`/`ports`/`adapters` mit `role`s, `edges`
`adapters→ports`, `adapters→core`, `ports→core`, ein `allow`-Re-Export und `composition_root`.

- **Happy:** Given diese `.a-check.yml`, when `a-check --print-graph <pfad>` läuft, then eine
  Mermaid-`flowchart`-Ausgabe mit **einem Knoten je Layer**, **einer Kante je `edge`**, **einer
  gestrichelten `allow`-Kante**, Rollen-`classDef`, Exit 0 — und es wird **nichts** ins Repo
  geschrieben (read-only).
- **Boundary:** Given eine minimale Config **ohne** optionale Blöcke (kein `adapter_sink`/`tech`/
  `composition_root`/`direction`), when es läuft, then ein gültiges Flowchart nur aus Layern + `edges`
  (kein Absturz, keine leeren Sonderknoten).
- **Boundary (effektive Rollen):** Given eine klassische Config ohne explizite `role:`-Felder
  (`core`/`ports`/`adapters`), when es läuft, then nutzt der Graph dieselben effektiven Rollen wie
  die Regel-Engine (Domain/Port/Adapter-Farben) — keine ungefüllten Rollen durch fehlendes `role:`.
- **Boundary (Mermaid-unsafe Namen):** Given eine **gültige** Config mit Layer-Namen, die Mermaid-heikle
  Zeichen tragen (`driven-adapters`, `a.b`, `x/y`, `"q"`, `a]b`, `a|b`, `a\b`, ``a`b``,
  `<a>&b` und CR/LF/TAB), when es läuft, then **gültiges** Mermaid — interne Knoten-IDs, die
  Rohnamen nur als escaped Labels nach §3 Punkt 5 (kein Syntaxfehler, keine Mehrdeutigkeit, kein
  Knoten-/Kanten-Ausbruch) [Medium-Fix].
- **Boundary (Sonderlabels + Dangling):** Given `composition_root`/`adapter_sink` mit Mermaid-heiklen
  Zeichen aus dem adversarischen Set (`]`, `|`, `\`, Backtick, `<`, `>`, `&`, Quotes und
  CR/LF/TAB; `tech` wird v2+ gesondert geprüft) sowie eine `edges`-/`allow`-Kante auf einen
  nicht deklarierten Layer-Namen, when es läuft, then bleibt das
  Mermaid gültig; die Sonderwerte sind escaped Labels, `composition_root` erscheint als isolierter
  Notizknoten ohne All-to-All-Kanten, `adapter_sink` als isolierter Ausnahme-/Notizknoten ohne
  Layer-Klasse und ohne gezeichnete Adapter-Kanten, und der unbekannte Endpunkt erscheint als stabiler
  Dangling-Knoten (kein neuer Config-Exit-2).
- **Negative:** Given eine **ungültige** `.a-check.yml` (unbekannter Schlüssel **oder** eine unbekannte
  Sprache `languages: {ruby: …}` [High-Fix]), ein unbekanntes Flag vor dem Pfad **oder** im neuen Modus
  ein zusätzliches Restargument/Flag nach dem Pfad (`--print-graph <pfad> --bogus`), when es läuft, then
  **Exit 2** (kein halbes Diagramm). Ladezeitige Config-Fehler sind identisch zum Scan; die zusätzliche
  Restargument-Regel gehört zum `--print-graph`-Usage-Vertrag. Scanzeitige, dateimengenabhängige
  Resolution-Mehrdeutigkeiten werden hier nicht erwartet, weil `--print-graph` keine Quellen liest.
- **Determinismus** ([AC-QA-01](../../../../spec/lastenheft.md#ac-qa-01--determinismus)): zwei Läufe
  gegen dieselbe Config → **byte-identische** Ausgabe (stabile Knoten-/Kanten-Sortierung).

*(Lastenheft-Mapping der drei Pflicht-AK: **Happy** = Happy, **Boundary** = minimale Config/effektive Rollen/Mermaid-unsafe/Dangling/Determinismus, **Negative** = Negative.)*

## 7. Grenzen / Folge

- **Kein Semantik-Beweis** ([AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)):
  der Graph zeichnet die **deklarierte** Config, nicht den realen Code — er ersetzt nicht den Scan.
  Die impliziten Verbote sind Legende, keine gezeichneten Kanten. Scanzeitige Fehler, die reale
  Datei-/Deklarationsmengen brauchen (insbesondere Mehr-Wurzel-Resolution-Mehrdeutigkeit), sind kein
  Vertrag dieses no-scan-Modus.
- **Nur Mermaid** in v1; DOT/Graphviz und ein Findings-/Verstoß-Graph (den realen Import-Graph mit
  markierten Befunden zeichnen) sind eigene Folge-Slices.
- **Kein Auto-Layout-Tuning** — Mermaid ordnet selbst; sehr große Flotten-Configs können unübersichtlich
  werden (dokumentierte Grenze).
- **Konsument:** maintainer-initiiert; ein realer Nutzen-Beleg ist die Dogfooding-Probe (a-checks eigener
  Graph) + Einbettung in Docs/PRs. `tech`-Vollausbau folgt bei Bedarf.

## 8. Sub-Area-Modus-Begründung

### Sub-Area: Spec-Schreibung (Lastenheft + Spezifikation)

- **Modus:** GF
- **Konventionen-Dichte:** hoch — Source Precedence, Anforderungs-Anlege-Prozess und
  ID-Schema sind in [harness/conventions.md](../../../../harness/conventions.md#anforderungs-anlege-prozess)
  verankert.
- **Phase-Reife:** Phase 4 — Lastenheft/Spezifikation führen, Code folgt über Tests und Gates.
- **Evidenz-/Diskrepanz-Risiko:** niedrig bis mittel — der neue Modus erweitert den öffentlichen
  CLI-Vertrag, aber die Slice-AKs binden Happy/Boundary/Negative explizit an Tests.
- **Reconciliation-Aufwand:** keiner erwartet; neue `CLI`-Anforderung, Versions-/Historie-Bumps und
  Spezifikationsschärfung entstehen im umsetzenden Slice.

### Sub-Area: Architektur-Sicht und Ports

- **Modus:** GF
- **Konventionen-Dichte:** hoch — [spec/architecture.md](../../../../spec/architecture.md) definiert
  ARC-Komponenten/Ports sprach- und meilensteinfrei; `arch-check` dogfoodet die Schichtgrenzen.
- **Phase-Reife:** Phase 4 — bestehende ARC-Struktur ist kanonisch, die neue Graph-Komponente wird
  additiv eingeordnet.
- **Evidenz-/Diskrepanz-Risiko:** mittel — `GraphPort` und `ExtractionPort.Validate` berühren
  a-checks Eigen-Hexagon; `arch-check` ist der direkte Sensor gegen laterale Adapter-Kanten.
- **Reconciliation-Aufwand:** keiner erwartet; Architektur-Update + `make arch-check` im Slice.

### Sub-Area: CLI/Adapter-Implementierung

- **Modus:** GF
- **Konventionen-Dichte:** hoch — Docker/make-only, Port-Programmierung und read-only CLI-Verhalten sind
  über AGENTS, Spec und bestehende `--print-*`-Tests vorgegeben.
- **Phase-Reife:** Phase 4 — bestehende `--print-config`/`--print-mk`-Muster sind implementiert und
  getestet; `--print-graph` ergänzt sie ohne Host-Toolchain-Wechsel.
- **Evidenz-/Diskrepanz-Risiko:** mittel — Go-`flag`-Restargumente und validation-only Sprachprüfung
  sind bekannte Randstellen; die Slice-AKs pinnen beide.
- **Reconciliation-Aufwand:** keiner erwartet; Umsetzung in CLI/Ports/Graph-Adapter plus gezielte Tests.

### Sub-Area: Tests und Gates

- **Modus:** GF
- **Konventionen-Dichte:** hoch — reale Targets sind in AGENTS/harness gelistet; `image-test`,
  `make gates`, `make ci` und `make trace-check` sind verpflichtend vor Handoff.
- **Phase-Reife:** Phase 4 — bestehende Gate-Landschaft ist real und grün; neue Tests hängen an die
  vorhandenen Targets.
- **Evidenz-/Diskrepanz-Risiko:** niedrig — die Fitness-Function ist im Slice als E2E- und Unit-Set
  ausgeschrieben.
- **Reconciliation-Aufwand:** keiner erwartet; `image-test` wird um den neuen Print-Modus erweitert.

### Sub-Area: Benutzer-Doku

- **Modus:** GF
- **Konventionen-Dichte:** mittel bis hoch — `docs/user/` ist kanonisches Nutzer-/Betriebs-Doku-Stratum,
  Currency-Bumps sind gelebte Praxis.
- **Phase-Reife:** Phase 4 — Handbuch und README existieren; der Slice ergänzt einen neuen
  Visualisierungsabschnitt.
- **Evidenz-/Diskrepanz-Risiko:** niedrig — Beispiel-Output und Dogfooding-Probe decken die
  Nutzererwartung.
- **Reconciliation-Aufwand:** keiner erwartet; Benutzerhandbuch-Bump im umsetzenden Slice.
