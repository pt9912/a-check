# slice-037 — HexSlice-Architektur: Gate-Lücken-Analyse

**Status:** **done (2026-07-25)** — **alle** Optionen sind entschieden und ihre Umsetzungen
geliefert (§10). Der Slice selbst enthält weiterhin keine Spec-/Code-Änderung; er ist die
Analyse, aus der drei Umsetzungs-Slices und eine Vertagung hervorgingen.

> **Currency 2026-07-25.** Die Analyse stammt vom 2026-07-24; seither ist ein Teil davon
> abgearbeitet:
> - **Option B + C geliefert** — [slice-039](slice-039-hexslice-vertical-slice-regeln.md)
>   hat `lateral-slice` (§4.1) und `port-locality` (§4.2) als
>   [AC-FA-RULE-009](../../../../spec/lastenheft.md#ac-fa-rule-009--slice-isolation-regel-lateral-slice)/[AC-FA-RULE-010](../../../../spec/lastenheft.md#ac-fa-rule-010--port-lokalität-regel-port-locality)
>   umgesetzt ([ADR-0026](../../adr/0026-hexslice-vertical-slice-regeln.md)). Die
>   Vertical-Slice-Achse ist damit gatebar — genau das, was §1 noch als „gar nicht ausdrückbar"
>   verdikt.
> - **Nebenbefund F-3 behoben** — [slice-038](slice-038-layer-tie-break-deklarationsreihenfolge.md)
>   (Tie-Break folgt der Deklarationsreihenfolge).
> - **Option A weitgehend geliefert** — Benutzerhandbuch §3.7 („Eine Vertical-Slice-Architektur
>   (HexSlice) absichern") ist die Anwendungs-Doku; die korrigierte Beispiel-Config liegt im
>   HexSlice-Repo und läuft dort real gegen 0.
> - **§4.3 (Richtung) ist entschieden, nicht mehr offen** — [slice-013 §0](../open/slice-013-driving-driven-vertiefung.md)
>   (2026-07-25): Port→Port verworfen, Auto-Inferenz vertagt. Für HexSlice heißt das: die
>   Richtungs-Prüfung bleibt an richtungs-getrennte Port-Schichten gebunden, dauerhaft.
> - **A′ nachgemessen und geliefert** — §4.0 war nur zur **Hälfte** behoben; die andere Hälfte war
>   kein Loch, sondern ein **Falsch-Positiv** (§4.0a). Behoben in
>   [slice-044](slice-044-ziel-glob-schattenwurf.md) ([ADR-0028](../../adr/0028-ziel-glob-schattenwurf.md)):
>   das Ziel gilt jetzt als extern statt als umschließende Schicht.
> Damit ist **jede** Option entschieden — Closure §10.

**Auslöser:** externe Frage — kann a-check die **HexSlice Architecture** (hexagonal + Vertical Slice) schon durchsetzen/gaten?
**Externe Quelle (out-of-repo):** das Architektur-Dokument `hexslice-architecture.de.md` aus dem
Schwester-Repo `hexslice-architecture` (Stand 2026-07-24; nicht Teil dieses Repos — daher nur
benannt, kein repo-relativer Link und kein Host-Pfad).
**Bezug:** wertet die HexSlice-Regeln gegen das bestehende Rollen-/Kanten-Modell
([AC-FA-RULE-006](../../../../spec/lastenheft.md#ac-fa-rule-006--schicht-rollen-generische-regel-anwendung)
… [AC-FA-RULE-008](../../../../spec/lastenheft.md#ac-fa-rule-008--driving-driven-port-richtung-regel-port-direction-mismatch),
[SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung)) und den Produkt-Scope
([Lastenheft §1](../../../../spec/lastenheft.md#1-zweck-und-geltungsbereich)). [Roadmap](../in-progress/roadmap.md).

> **Hinweis:** Analyse zur Abnahme. Vorgeschlagene Regelnamen/Anforderungen sind
> **unverbindlich** bis Freigabe in `spec/`. Es werden hier **keine** `AC-*`/`ADR-*`-IDs
> vergeben (Anlege-Prozess: [conventions §Anforderungs-Anlege-Prozess](../../../../harness/conventions.md#anforderungs-anlege-prozess)).
> Entscheidungen §7 **vor** jeder Umsetzung.

---

## 1. Ziel

Beantworten, **ob und wieweit** a-check die HexSlice Architecture heute als Gate durchsetzen
kann — und die verbleibenden Lücken so präzise benennen, dass daraus (falls gewünscht) ein
Umsetzungs-Slice mit sauberem Scope entsteht. **Kein** Code, **keine** neue Anforderung in
diesem Slice.

**Verdikt vorweg (nach adversarischem Review + Fixture-Verifikation korrigiert):** Die
**hexagonale Achse** (Domain-/App-Reinheit, Kanten-Richtung, `tech`-Kapselung,
Adapter-Lateralität) ist heute gatebar. Die **Port-Rolle** ist es nur **eingeschränkt**:
tiefen-agnostisch verschachtelte Ports im `application`-Teilbaum (`application/**/ports/**`)
werden von der umschließenden `application`-Schicht **überschattet** (Literal-Präfix-Gleichstand →
`application` gewinnt) und als Rolle `app` fehlklassifiziert — nur ein **business-area-spezifischer**
Glob (längerer Literal-Präfix) oder **app-weite** Ports außerhalb des Teilbaums lösen korrekt auf
(§4.0, gegen `a-check:dev` verifiziert). Die **Vertical-Slice-Achse** (Slice-Isolation,
Port-Lokalität) ist gar nicht ausdrückbar. **a-check gatet den *Hexagon* (Port-Rolle mit
Glob-Vorbehalt), nicht die *Slices*.**

## 2. Was HexSlice fordert

Aus dem Quelldokument (§Abhängigkeitsregeln + §Regeln 1–8):

1. Domain technologieunabhängig.
2. Application enthält die Use Cases.
3. Use Cases als **Vertical Slices** organisiert.
4. Ports vom Application Core definiert.
5. **Ports leben so lokal wie möglich und so gemeinsam wie nötig** (use-case-lokal →
   business-area-shared → application-wide).
6. Inbound Adapter rufen Use Cases auf.
7. Outbound Adapter implementieren Ports.
8. Der Core hängt nicht von Infrastruktur ab.

Erlaubt: `Adapters → Application`, `Application → Domain`, `Adapters → Ports`.
Verboten: `Domain → Application`, `Domain → Adapters`, `Application → Adapters`,
`Application → Infrastructure`.

## 3. Mapping HexSlice → a-check

| # | HexSlice-Forderung | a-check-Mechanik | Regel | Status |
|---|---|---|---|---|
| 1 | Domain technologiefrei | Rolle `domain` | [`core-impurity`](../../../../spec/lastenheft.md#ac-fa-rule-001--kern-reinheit-regel-core-impurity) (kategorisch) | ✅ heute |
| 2 | Application = Use Cases | Rolle `app` | [`app-impurity`](../../../../spec/lastenheft.md#ac-fa-rule-007--rolle-app-und-strenge-domain) | ✅ heute |
| 4 | Ports Besitz der Application (app-weit / business-area-spezifisch) | Rolle `port` | [`port-impurity`](../../../../spec/lastenheft.md#ac-fa-rule-004--port-disziplin-regel-port-impurity) | ⚠️ **eingeschränkt** (§4.0) |
| 4 | Ports **tiefen-agnostisch verschachtelt** (`application/**/ports/**`) | von `application` überschattet → Rolle `app` | *fehlklassifiziert* | ❌ **Lücke** (§4.0) |
| 6 | Inbound Adapter | `adapter` + `direction: driving` | — | ✅ heute |
| 7 | Outbound Adapter impl. Ports | `adapter` + `direction: driven` | [`port-direction-mismatch`](../../../../spec/lastenheft.md#ac-fa-rule-008--driving-driven-port-richtung-regel-port-direction-mismatch) | ⚠️ **nur bei richtungs-getrennten Port-Schichten** (§4.3) |
| — | `Adapters→App`, `App→Domain`, `Adapters→Ports` erlaubt | `edges` | [`wrong-direction`](../../../../spec/lastenheft.md#ac-fa-rule-005--schicht-richtung-regel-wrong-direction) | ✅ heute |
| — | `Domain→App/Adapters` verboten | `core-impurity` kategorisch | — | ✅ heute |
| — | `App→Adapters` verboten | `app-impurity` | — | ✅ heute |
| 8 | `App→Infrastructure` verboten | `tech`-Muster | [`app-impurity`](../../../../spec/lastenheft.md#ac-fa-rule-007--rolle-app-und-strenge-domain) (App zuerst) / [`tech-leak`](../../../../spec/lastenheft.md#ac-fa-rule-003--tech-kapselung-regel-tech-leak) | ✅ heute |
| — | Adapter ruft keinen Adapter (lateral) | `adapter_sink` | [`lateral-adapter`](../../../../spec/lastenheft.md#ac-fa-rule-002--keine-lateralen-adapter-kanten-regel-lateral-adapter) | ✅ heute |
| 3 | **Slice-Isolation** (Use Case A ↛ Interna Use Case B) | — | *keine* | ❌ **Lücke** (§4.1) |
| 5 | **Port-Lokalität** (use-case-lokaler Port nur im eigenen Slice) | alle Ports = 1 Rolle `port` | *keine* | ❌ **Lücke** (§4.2) |

Der HexSlice-Ordnerbaum (`hexagon/domain/**`, `hexagon/application/**`, `.../ports/**`,
`adapters/inbound|outbound/**`) ist reiner Glob-Stoff und mappt 1:1 auf `layers`
([SPEC-CONF-001](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema)).

## 4. Die Lücken (nach Fixture-Verifikation)

### 4.0 Port-Rollen-Klassifikation verschachtelter Ports (verifiziert)

**Befund F-1 des Reviews — die überraschendste Lücke.** HexSlice legt Ports *innerhalb* des
`application`-Teilbaums ab (use-case-lokal `application/<ba>/<uc>/ports/**`, business-area-shared
`application/<ba>/ports/**`). Ein tiefen-agnostischer Glob dafür ist `src/hexagon/application/**/ports/**`
— sein **literaler Präfix** (der Teil vor dem ersten Wildcard-Segment) ist `src/hexagon/application`,
**identisch** zu dem der umschließenden `application`-Schicht `src/hexagon/application/**`. Bei diesem
**Gleichstand** gewinnt `application` → die verschachtelten Ports werden als Rolle `app`
fehlklassifiziert; `port-impurity`/`forbidden_constructs` greifen für sie **nicht**.

*Verifiziert gegen `a-check:dev`:* eine verschachtelte Port-Datei mit verbotenem Konstrukt liefert
unter dem tiefen-agnostischen Glob **0 Befunde** (überschattet) — **auch nach YAML-Reihenfolge-Tausch**;
sie feuert `port-impurity` erst, wenn der Ports-Glob **keinen** Overlap hat (app-weit außerhalb des
Teilbaums) **oder** einen **längeren Literal-Präfix** trägt (business-area-spezifisch, z. B.
`application/orders/**/ports/**` → Präfix `…/application/orders`).

**Nebenbefund F-3 (a-check-intern, out-of-scope hier, aber real):**
[SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung) und der Kommentar
in `LayerOf` versprechen bei Gleichstand die **zuerst deklarierte** Schicht. Tatsächlich wird
`layers` als Go-Map dekodiert und **alphabetisch** sortiert (`sortedKeys`) — die
Deklarationsreihenfolge ist unwiederbringlich, der Tie-Break ist faktisch **alphabetisch**
(`application` < `ports`). Ein YAML-Reorder kann die Überschattung deshalb nie beheben. Latenter
Korrektheitsfehler in a-check; **Kandidat für einen eigenen a-check-Slice** (Spec/Impl-Abgleich:
entweder Deklarationsreihenfolge bewahren *oder* die Spec auf „alphabetisch" korrigieren). HexSlice
ist davon nicht blockiert (der Literal-Präfix-Workaround umgeht den Tie), aber die Ergonomie leidet.

> **Behoben in [slice-038](slice-038-layer-tie-break-deklarationsreihenfolge.md)** (2026-07-24,
> Konformitäts-Bugfix zu [ADR-0013](../../adr/0013-layerof-laengster-praefix.md)): der Tie-Break folgt
> jetzt der **Deklarationsreihenfolge**. Für HexSlice heißt das — ein Konsument deklariert die
> `ports`-Schicht **vor** `application`, und tiefen-agnostisch verschachtelte Ports (`application/**/ports/**`)
> lösen korrekt als `port` auf; das Business-Area-Aufzählen aus §5 entfällt dann. Binär verifiziert.

**Konsequenz für HexSlice:** die Port-Disziplin ist heute nur für **app-weite** Ports und **je
Business-Area einzeln aufgezählte** Port-Globs durchsetzbar — nicht mit *einem* tiefen-agnostischen
Muster. Das ist eine ausgewiesene Grenze
([AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)), keine
stille Vollständigkeit.

### 4.0a Nachmessung 2026-07-25 — F-1 ist halb behoben, der Rest ist ein **Falsch-Positiv**

Gemessen gegen `a-check:dev` (Stand `main`) mit einem HexSlice-Fixture: Schichten `ports`
(**vor** `application` deklariert, Glob **tiefen-agnostisch** `src/hexagon/application/**/ports/**`),
`application`, `domain`, `outbound`; Kanten u. a. `outbound → ports`, **nicht** `outbound → application`.

| Achse | Mechanik | Ergebnis |
|---|---|---|
| **Quell-Seite** (Datei → Schicht, `LayerOf`) | wertet den **literalen Präfix** (`litPrefixLen`) und bricht den Gleichstand nach **Deklarationsreihenfolge** | ✅ **behoben**: die verschachtelte Port-Datei klassifiziert als `port` — sie meldet `port-impurity` für einen Adapter-Import **und** für ein `forbidden_constructs`-Muster |
| **Ziel-Seite** (Import → Schicht, `layerOfCand`) | matcht den **rohen** `globPrefix` als Segment-Run (`segIndex`) | ❌ **offen**: `src/hexagon/application/**/ports` enthält ein Wildcard-Segment und matcht **nie** einen realen Kandidaten — der Import fällt auf das `application`-Glob zurück |

**Der Rest ist kein Loch, sondern ein Fehlbefund.** Der legitime Import eines Outbound-Adapters
auf einen verschachtelten Port meldet:

```text
src/adapters/outbound/db.go:3: wrong-direction: outbound -> application (…/createorder/ports)
```

Die Kante `outbound → ports` **ist** deklariert; a-check sieht sie nur nicht, weil das Ziel als
`application` auflöst. **Gegenprobe** mit business-area-/use-case-spezifischem Glob (sauberer
literaler Präfix): **0 Befunde, Exit 0** — der Workaround aus §5 trägt weiterhin.

**Warum das schwerer wiegt als die ursprüngliche Diagnose (§4.0):** eine *Lücke* lässt einen
Verstoß durch; dieser *Falsch-Positiv* drängt den Nutzer zur falschen Reparatur — er ergänzt
`{from: outbound, to: application}`, um sein Gate grün zu bekommen, und **verdeckt damit dauerhaft
echte** `outbound → application`-Verstöße. Ein Falsch-Positiv, dessen naheliegende Behebung ein
Falsch-Negativ erzeugt, ist die teuerste Sorte.

**Das ist präzise Option A′** — und der Scope schrumpft dadurch: nicht „Overlap-Ergonomie
allgemein", sondern **eine** Asymmetrie zwischen den zwei Auflösungs-Pfaden. `LayerOf` kann
Wildcard-Präfixe verarbeiten, `layerOfCand` nicht. Die dokumentierte Grenze aus
[ADR-0026](../../adr/0026-hexslice-vertical-slice-regeln.md) („Globs mit Wildcard in der Mitte
lösen als Import-**Ziel** nicht auf") ist genau diese Stelle — sie ist als Grenze ausgewiesen,
aber ihre **Nebenwirkung** (Rückfall auf das umschließende Glob statt „extern") ist es nicht.

### 4.1 Slice-Isolation (HexSlice-Regel 3)

Das Herz des Ansatzes: ein Use-Case-Slice koppelt nicht an die Interna eines anderen. a-check
kennt eine analoge Isolation **nur** für die Adapter-Rolle
([`lateral-adapter`](../../../../spec/lastenheft.md#ac-fa-rule-002--keine-lateralen-adapter-kanten-regel-lateral-adapter),
Sub-Einheiten relativ zum Schicht-Glob-Präfix, siehe
[slice-024](slice-024-adapterseg-root-subeinheit.md)). Für die Rolle `app` gibt es
**kein** Pendant — zwei `application/…/use-case-a` und `…/use-case-b` sind derselbe Layer,
Cross-Import erzeugt keinen Befund.

**Technisch bereits vorhanden:** die Sub-Einheiten-Mechanik (Pfad-Rest nach dem Glob-Präfix,
Blatt-Klassifikation) existiert für `lateral-adapter`. Ein `lateral-slice`/`slice-coupling`
wäre strukturell **dieselbe** Logik auf die `app`-Rolle angewandt — kein neuer Extraktions-
oder Auflösungs-Mechanismus, nur ein zweiter Regel-Arm.

### 4.2 Port-Lokalität (HexSlice-Regel 5)

„So lokal wie möglich, so gemeinsam wie nötig" ist ein **Scope**-Konzept über drei Port-Ebenen
(use-case / business-area / application-wide). a-check kollabiert alle drei auf **eine** Rolle
`port` und kann nicht erkennen, dass ein use-case-lokaler Port von einem *fremden* Slice
importiert wird. Das ist die schwierigere Lücke: sie braucht einen **Port-Scope-Begriff**
(welcher Import-Kreis darf einen Port sehen), nicht nur einen zweiten `lateral`-Arm.

Denkbare, aber **noch nicht bewertete** Modellierungen:
- Port-Scope aus der Verschachtelungstiefe relativ zum `application`-Glob ableiten
  (use-case-lokal = tiefster, app-wide = flachster) — heuristisch, package==directory. *Setzt §4.0
  voraus:* solange verschachtelte Ports gar nicht als `port` klassifizieren, ist Lokalität moot.
- Explizite `scope`-Deklaration je Port-Sub-Layer in `layers` — verboser, aber deterministisch.

### 4.3 Richtungs-Prüfung nur bei getrennten Port-Schichten (verifiziert)

**Befund F-2 des Reviews.** `port-direction-mismatch`
([AC-FA-RULE-008](../../../../spec/lastenheft.md#ac-fa-rule-008--driving-driven-port-richtung-regel-port-direction-mismatch))
feuert nur, wenn **beide** Seiten `direction` tragen; `dirOf` inferiert **keine** Richtung aus dem
Namen (verifiziert: `rules.go` `dirOf`). Basis-HexSlice modelliert Ports als *einen* Begriff (nicht
`driving`/`driven`-getrennt) — mit einer undirektionalen `ports`-Schicht ist die Regel **inert**,
und `direction:` auf den Adaptern bleibt wirkungslos. Voll nutzbar erst, wenn ein Konsument die Ports
richtungs-getrennt modelliert (genau das
[slice-013](../open/slice-013-driving-driven-vertiefung.md)-Terrain: Auto-Inferenz/Port→Port, dort *gated*).

## 5. Was heute schon geht — Beispiel-`.a-check.yml` (verifiziert)

Für einen HexSlice-Baum ist die **hexagonale Achse** ohne a-check-Änderung gatebar — mit dem
**Port-Vorbehalt aus §4.0**: verschachtelte Ports brauchen einen **business-area-spezifischen**
Glob (längerer Literal-Präfix als `application/**`), sonst überschattet sie die `application`-Schicht.
Skizze (unverbindlich; gegen `a-check:dev` auf die Klassifikations-Punkte geprüft):

```yaml
version: 1
languages:
  typescript: ["**/*.ts"]
layers:
  domain:      ["src/hexagon/domain/**"]
  application: ["src/hexagon/application/**"]
  # Ports: app-weit (außerhalb des application-Teilbaums) + JE BUSINESS-AREA EINZELN
  # aufgezählt — der Business-Area-Literal (…/orders, …/billing) gibt dem Ports-Glob
  # einen längeren Literal-Präfix als application/** und verhindert die Überschattung (§4.0).
  # Ein tiefen-agnostisches "application/**/ports/**" würde STILL als app fehlklassifizieren.
  ports:
    globs:
      - "src/hexagon/ports/**"                    # application-wide
      - "src/hexagon/application/orders/**/ports/**"
      - "src/hexagon/application/billing/**/ports/**"
    role: port
  inbound:  {globs: ["src/adapters/inbound/**"],  role: adapter}   # direction: siehe §4.3 (in Basis-HexSlice inert)
  outbound: {globs: ["src/adapters/outbound/**"], role: adapter}
edges:
  - {from: inbound,     to: application}
  - {from: application, to: domain}
  - {from: application, to: ports}
  - {from: outbound,    to: ports}
  - {from: ports,       to: domain}
adapter_sink: shared
tech:
  - {pattern: "typeorm", adapter: outbound}
  - {pattern: "express", adapter: inbound}
composition_root: ["src/main/**"]
```

Damit greifen Domain-/App-Reinheit, Kanten-Richtung, `tech`-Kapselung, Adapter-Lateralität **und**
die Port-Disziplin **für die aufgezählten** Port-Orte. **Nicht** gegatet: verschachtelte Ports ohne
eigenen Business-Area-Glob (§4.0), inbound/outbound-**Richtung** ohne getrennte Port-Schichten (§4.3),
Slice-Isolation (§4.1), Port-Lokalität (§4.2). Das gehört in eine Nutzer-Doku **ehrlich ausgewiesen**,
nicht als Vollständigkeit behauptet
([AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)).

## 6. Scope-Einordnung

Die Lücken der Vertical-Slice-Achse (§4.1/§4.2) sind **In-Scope-Kandidaten**, kein Scope-Bruch: der
[Lastenheft-Zweck](../../../../spec/lastenheft.md#1-zweck-und-geltungsbereich) stellt a-check
ausdrücklich als Ergänzung um die *feingranularen* Fitness-Functions dar, „die der Compiler
nicht abdeckt (laterale Adapter-Kanten, Port-Disziplin)". Slice-Isolation und Port-Lokalität
sind exakt diese Klasse. Ihre Umsetzung wäre eine **Regel-Erweiterung** (neue
`AC-FA-RULE-<NNN>`-Anforderung(en) im Lastenheft, je Folge-ADR + Spec-Schärfung), analog wie
[slice-011](slice-011-app-rolle.md) die Rolle `app` und
[slice-012](slice-012-driving-driven-layerof.md) die Richtung ergänzt hat.

Die Klassifikations-Lücke (§4.0) ist **anderer Natur** — kein neues Feature, sondern eine
**bestehende Grenze/Inkonsistenz** (Literal-Präfix-Überschattung + Spec/Impl-Tie-Break-Divergenz,
Nebenbefund F-3). Sie berührt den Kern-Resolver und gehört in einen **eigenen a-check-Slice**
(Spec/Impl-Abgleich), unabhängig von HexSlice — mit HexSlice als einem Motiv unter mehreren.

## 7. Optionen & offene Entscheidungen

- **Option A — nur anwenden (kein a-check-Change):** eine Beispiel-`.a-check.yml` (§5, mit
  business-area-aufgezählten Ports-Globs) + Nutzer-Doku für HexSlice-Bäume liefern; §4.0–§4.3 als
  bekannte Grenzen ausweisen. Kleinster Schritt, sofort lieferbar. *Empfehlung als erster Schritt.*
- **Option A′ — Klassifikations-Grenze (§4.0/F-3) als eigener a-check-Slice:** Spec/Impl-Tie-Break
  angleichen (Deklarationsreihenfolge bewahren *oder* Spec auf „alphabetisch" korrigieren); optional
  ein spezifischerer-Glob-gewinnt-über-Overlap-Mechanismus, damit tiefen-agnostische Ports-Globs
  ohne Business-Area-Aufzählung auflösen. **Unabhängig von HexSlice**, verbessert jeden Konsumenten
  mit geschachtelten Sub-Layern. *Empfehlung: nach A, priorisiert (behebt einen echten Fehler).*
- **Option B — Slice-Isolation nachrüsten:** neue Regel `lateral-slice` (Rolle `app`,
  Sub-Einheiten-Mechanik aus [slice-024](slice-024-adapterseg-root-subeinheit.md)
  wiederverwendet) + neue `AC-FA-RULE-<NNN>` + Folge-ADR + Spec. Mittlerer Aufwand, hoher
  HexSlice-Wert. **Gate wie üblich: erst ein realer Konsument, der die Slice-Kopplung *fühlt*.**
- **Option C — Port-Lokalität nachrüsten:** Port-Scope-Begriff (§4.2). Größter Design-Anteil
  (Scope-Modell zu entscheiden); **setzt §4.0 voraus** (ohne Port-Klassifikation keine Lokalität);
  separat von B halten.

**Entscheid 0 — Konsumenten-Gate (Präzedenz [slice-013 Entscheid 0](../open/slice-013-driving-driven-vertiefung.md#6-offen--entscheidungen-zur-abnahme)):**
Gibt es einen **realen** HexSlice-Konsumenten mit `.a-check.yml`, der die fehlende
Slice-Isolation spürt? Ohne ihn ist B/C verfrühte Ergonomie (dieselbe „aktiver-Konsument"-Linie,
die slice-012/013 setzen). *Empfehlung: A jetzt; B/C erst mit Pilot.*

**Entscheid 1 — Slice-Grenze:** Was ist die „Slice-Einheit" im `application`-Glob? Das erste
Segment nach dem Glob-Präfix (`business-area`) oder das `use-case`-Verzeichnis? HexSlice
verschachtelt `application/<business-area>/<use-case>/` — die Wahl bestimmt, ob business-area-
interne Cross-Use-Case-Importe erlaubt sind.

**Entscheid 2 — Port-Scope-Modell (nur bei C):** tiefen-heuristisch (4.2 erste Variante) vs.
explizite `scope`-Deklaration. Determinismus
([AC-QA-01](../../../../spec/lastenheft.md#ac-qa-01--determinismus)) und package==directory-Grenze
beachten.

## 8. Definition of Done (dieser Analyse-Slice)

- [x] Ist-Mapping HexSlice → a-check (§3) vollständig, jede Regel bewertet.
- [x] Lücken benannt und technisch eingeordnet (§4.0 Klassifikation, §4.1 Slice-Isolation,
      §4.2 Port-Lokalität, §4.3 Richtung).
- [x] **Adversarisches Review + Fixture-Verifikation gegen `a-check:dev`** (F-1 Überschattung
      reproduziert: Config A/B je 0 Befunde, T2/T3 feuern; F-2 Richtung inert am Code; F-3
      alphabetischer Tie-Break in `config.go` `sortedKeys` bestätigt).
- [x] Beispiel-`.a-check.yml` für die heute-gatebare Achse (§5), Port-Vorbehalt eingearbeitet.
- [x] **Abnahme B/C erfolgt** (2026-07-24) und in [slice-039](slice-039-hexslice-vertical-slice-regeln.md)
      umgesetzt; **A** über Benutzerhandbuch §3.7 + korrigierte Beispiel-Config geliefert;
      **§4.3** über [slice-013 §0](../open/slice-013-driving-driven-vertiefung.md) entschieden.
- [x] **Nachmessung 2026-07-25** (§4.0a): F-1 quell-seitig behoben, ziel-seitig als
      **Falsch-Positiv** reproduziert (Fixture + Gegenprobe gegen `a-check:dev`).
- [x] **Abnahme A′ erfolgt (2026-07-25):** Variante 2 („kleiner Fix statt Feature") gewählt — die
      Zuordnung wird zurückgezogen (Ziel = extern) statt nachgerüstet. Umgesetzt in
      [slice-044](slice-044-ziel-glob-schattenwurf.md)
      ([ADR-0028](../../adr/0028-ziel-glob-schattenwurf.md), Spezifikation 0.25.0). Damit sind
      **alle** Optionen dieses Analyse-Slices entschieden; er kann mit slice-044 nach `done/`.

## 9. Closure-Notiz

_(beim Abschluss: gewählte Option + Folge-Slice-Verweis; kein Gate-Beleg nötig — reine Analyse,
kein Code/Vertrag berührt.)_

## 10. Closure-Notiz (2026-07-25)

Der Analyse-Slice ist **erledigt**: jede Option hat eine Entscheidung, und jede Entscheidung ihre
Umsetzung. Wohin das Ergebnis gegangen ist:

| Option / Befund | Entscheidung | Wo es gelandet ist |
|---|---|---|
| **A** — nur anwenden (Beispiel-Config + Nutzer-Doku) | angenommen | Benutzerhandbuch §3.7 („Eine Vertical-Slice-Architektur (HexSlice) absichern") + korrigierte Beispiel-Config im HexSlice-Repo (dort real gegen 0 verifiziert) |
| **B** — Slice-Isolation | angenommen | [slice-039](slice-039-hexslice-vertical-slice-regeln.md): `lateral-slice`, [AC-FA-RULE-009](../../../../spec/lastenheft.md#ac-fa-rule-009--slice-isolation-regel-lateral-slice), [ADR-0026](../../adr/0026-hexslice-vertical-slice-regeln.md) |
| **C** — Port-Lokalität | angenommen | [slice-039](slice-039-hexslice-vertical-slice-regeln.md): `port-locality`, [AC-FA-RULE-010](../../../../spec/lastenheft.md#ac-fa-rule-010--port-lokalität-regel-port-locality) |
| **A′** — Klassifikations-/Overlap-Grenze (§4.0, §4.0a) | angenommen in **zwei** Schritten | quell-seitig [slice-038](slice-038-layer-tie-break-deklarationsreihenfolge.md) (Tie-Break = Deklarationsreihenfolge, Nebenbefund F-3); ziel-seitig [slice-044](slice-044-ziel-glob-schattenwurf.md) ([ADR-0028](../../adr/0028-ziel-glob-schattenwurf.md)) |
| **§4.3** — Richtungs-Prüfung ohne getrennte Port-Schichten | entschieden, **nicht** gebaut | [slice-013 §0](../open/slice-013-driving-driven-vertiefung.md): Port→Port verworfen, Auto-Inferenz vertagt — die Richtung bleibt dauerhaft an richtungs-getrennte Port-Schichten gebunden |

**Verdikt-Korrektur.** §1 urteilte: „a-check gatet den *Hexagon*, nicht die *Slices*." Das gilt
nicht mehr — die Vertical-Slice-Achse ist seit [slice-039](slice-039-hexslice-vertical-slice-regeln.md)
gatebar, und der Port-Vorbehalt aus §4.0 ist in beiden Richtungen abgearbeitet. Was **bleibt**, ist
eine Config-**Disziplin**: Globs brauchen saubere literale Verzeichnispräfixe, sonst sind ihre Ziele
repo-extern und ungegatet (ausgewiesen in Handbuch §3.4/§3.7, nicht mehr als Fehlbefund).

### Lerneinträge

1. **Eine Gap-Analyse altert schnell — und still.** Zwischen Erstfassung und Abschluss lagen ein
   Tag und vier Slices; drei Aussagen im Verdikt waren am Ende überholt. Ein Analyse-Dokument
   braucht darum einen **Currency-Block**, nicht nur eine Closure-Zeile: sonst zitiert jemand ein
   Verdikt weiter, das die eigene Umsetzung längst widerlegt hat.
2. **„Lücke" ist nicht die schlimmste Diagnose.** Die Nachmessung (§4.0a) fand an derselben Stelle
   statt einer Lücke einen **Fehlbefund**, dessen naheliegende Reparatur ein dauerhaftes
   Falsch-Negativ erzeugt hätte. Bei der Priorisierung gehört diese Klasse **über** das bloße
   „fehlt noch".
3. **Zwei Auflösungs-Pfade für dieselbe Frage driften auseinander.** Datei→Schicht und
   Import→Schicht beantworten beide „welche Schicht?", waren aber verschieden implementiert — die
   eine wildcard-tolerant, die andere nicht. Der Unterschied blieb jahrelang unsichtbar, weil kein
   Konsument Globs mit Innen-Wildcard nutzte.
