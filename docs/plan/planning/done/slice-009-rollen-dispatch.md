# slice-009 — Rollen-Dispatch (welle-10a)

**Status:** done.
**Welle:** welle-10-regel-engine-generalisierung (Inkrement **a**).
**Bezug:** Re-Evaluierungs-Trigger aus [ADR-0008](../../adr/0008-ports-duerfen-domaenen-typen-referenzieren.md); generalisiert die Anwendung von [AC-FA-RULE-001](../../../../spec/lastenheft.md#ac-fa-rule-001--kern-reinheit-regel-core-impurity)/[AC-FA-RULE-002](../../../../spec/lastenheft.md#ac-fa-rule-002--keine-lateralen-adapter-kanten-regel-lateral-adapter)/[AC-FA-RULE-004](../../../../spec/lastenheft.md#ac-fa-rule-004--port-disziplin-regel-port-impurity) und erweitert [AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml); [Roadmap welle-10](../in-progress/roadmap.md).

> **Hinweis:** Dieses Dokument hält den **Entwurf** zur Abnahme (zwei Review-Runden
> eingearbeitet, 2026-06-22). Die in Code-Fences gesetzten Anforderungs-/ADR-Texte
> sind Vorlagen — verbindlich erst nach Freigabe in
> [`spec/lastenheft.md`](../../../../spec/lastenheft.md) bzw. einer neuen ADR.

---

## 1. Ziel

Die benannten Reinheits-Regeln (`core-impurity`, `port-impurity`,
`lateral-adapter`) vom Layer-**Namen** entkoppeln: Dispatch über eine
Layer-**Rolle** ∈ {`domain`, `port`, `adapter`}, explizit (`role:`) oder per
Namens-Inferenz. Damit prüfen die vier Konsumenten-Repos
(b-cad/d-check/grid/d-migrate) ihre je eigenen Strukturen voll, ohne die Engine
zu forken.

- **Scope a (dieser Slice):** Rollen {`domain`, `port`, `adapter`} + Namens-Inferenz;
  import- **und** konstrukt-basierte `port-impurity` rollen-basiert; `lateral-adapter`
  für Importe zwischen **verschiedenen** `role: adapter`-Schichten (Layer-Identität,
  namensunabhängig, **kategorisch** — nur `adapter_sink` hebt auf).
- **Scope b (Folge, welle-10b):** `app`-Rolle (Domain/Application-Trennung),
  `driving`/`driven`-Ports, und die **Namens-Generalisierung von `adapterSeg`**
  für die *Intra*-Schicht-Unterscheidung (Review-Finding **R6**).

## 2. Problem

**Vier** namensgebundene Stellen in `internal/hexagon/core/rules.go` — nicht drei:

```go
// (1-3) ruleFor — import-basiert:
case f.Layer == "core"     && (tl == "adapters" || isTech):   // core-impurity
case f.Layer == "ports"    && (tl == "adapters" || isTech):   // port-impurity
case f.Layer == "adapters" && tl == "adapters" && lateral():  // lateral-adapter

// (4) Evaluate — konstrukt-basierte port-impurity (rules.go:24-28):
if f.Layer == "ports" {
    for _, c := range f.Constructs { fs = append(fs, Finding{Rule: "port-impurity", ...}) }
}
```

`wrong-direction`/`tech-leak` sind bereits generisch. **Achtung:** `lateral()`
unterscheidet zwei Adapter über `adapterSeg` (Pfadsegment nach dem Literal
`"adapters"`), **nicht** über Layer-Identität. Selbst nach Rollen-Generalisierung
des äußeren Guards (Eintrittsbedingung) passiert dieser zwar — doch dann liefert
`adapterSeg` für `src/geometry/**` / `src/persistence/**` beide Male `""` ⇒ kein
Befund. Eine namensunabhängige `lateral-adapter`-Prüfung braucht daher in Scope a
den **Cross-Layer-Zweig** (Layer-Identität); die `adapterSeg`-Generalisierung der
Intra-Schicht-Unterscheidung bleibt b.

## 3. Entwurf (zur Abnahme)

### 3.1 Neue Anforderung — Schicht-Rollen (generische Regel-Anwendung)

```text
AC-FA-RULE-006 — Schicht-Rollen (generische Regel-Anwendung)
Generalisiert: AC-FA-RULE-001 / AC-FA-RULE-002 / AC-FA-RULE-004 (namens- -> rollen-basiert).

Beschreibung: Die Reinheits-Regeln core-impurity, port-impurity (import- UND
konstrukt-basiert) und lateral-adapter werden über die ROLLE einer Schicht
angewandt, nicht über ihren Namen. Eine Schicht trägt optional eine Rolle in
{domain, port, adapter}; fehlt sie, wird sie aus konventionellen Namen abgeleitet
(core->domain, ports->port, adapters->adapter). Eine explizite role: hat VORRANG
vor der Namens-Inferenz (ein Layer "core" mit role: adapter ist Adapter). In
Scope a ist die Inferenz nicht abschaltbar: ein konventionell benannter Layer
bekommt zwangsläufig eine Rolle (Rückwärtskompat-Garantie). Eine Schicht ohne
Rolle (weder deklariert noch ableitbar) unterliegt nur den kanten-basierten
Regeln (wrong-direction/tech-leak). Rollen->Regel: domain->core-impurity,
port->port-impurity, adapter->lateral-adapter.

lateral-adapter feuert in diesem Inkrement für Importe zwischen VERSCHIEDENEN
role:adapter-Schichten (Layer-Identität, namensunabhängig). Es ist KATEGORISCH:
nur adapter_sink hebt auf, NICHT allow/edges (analog zur Intra-Schicht-lateral).
Die Intra-Schicht-Unterscheidung über adapterSeg bleibt unverändert (klassischer
"adapters"-Name) und wird in welle-10b namens-generalisiert. Befund-NAMEN bleiben
unverändert.

Akzeptanzkriterien:
- Happy: Given zwei VERSCHIEDENE Schichten "geometry" und "persistence", beide
  role: adapter, when "geometry" "persistence" importiert (auch bei einer
  deklarierten allow-Kante zwischen ihnen), then Befund lateral-adapter
  (namensunabhängig, Cross-Layer, kategorisch — nur adapter_sink hebt auf).
- Boundary: Given eine Config mit klassischen Namen core/ports/adapters OHNE role,
  when a-check läuft, then identisches Verhalten wie 0.2.0 — inklusive der
  konstrukt-basierten port-impurity und der Intra-adapters-adapterSeg-Prüfung.
- Negative: Given (a) ein role: domain-Layer importiert einen role: adapter-Layer,
  ODER (b) ein role: port-Layer mit FREMDEM Namen, für den forbidden_constructs
  deklariert ist, enthält ein verbotenes Konstrukt, when a-check läuft, then Befund
  (a) core-impurity bzw. (b) port-impurity, Exit 1.

Out-of-Scope: app-Rolle, driving/driven-Ports; adapterSeg-Namens-Generalisierung
für die Intra-Schicht-Unterscheidung — welle-10b.
```

### 3.2 Schema-Erweiterung (Config)

Erweitert [AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml)
ohne neue AC: ein `layers`-Eintrag ist **entweder** eine Glob-Liste **oder** ein
Objekt mit `role`:

```yaml
layers:
  core:    ["internal/core/**"]                          # Kurzform, Rolle per Inferenz
  geometry: {globs: ["src/geometry/**"], role: adapter}  # Objektform, explizite Rolle
```

Strict-decode bleibt: unbekannter Schlüssel im Objekt ⇒ Exit 2 (siehe
Implementierungs-Gotcha §4 Schritt 1).

### 3.3 Folge-ADR — Rollen-basierter Regel-Dispatch

```text
ADR-0009 — Rollen-basierter Regel-Dispatch
Status: Proposed (-> Accepted bei Merge).
Bezug: AC-FA-RULE-006 (neu), AC-FA-CONF-001 (Schema), §1 Zweck.
Schärft: SPEC-RULE-001 + SPEC-CONF-001.

Kontext: Die Engine bindet core-/port-/lateral-Reinheit an die Literal-Namen
core/ports/adapters (vier Stellen, inkl. konstrukt-basierter port-impurity in
Evaluate) — Re-Evaluierungs-Trigger aus ADR-0008.

Decision: Die Regeln dispatchen über eine Layer-Rolle (domain/port/adapter), aus
role: (Vorrang) oder Namens-Inferenz. Import- UND konstrukt-basierte port-impurity
werden rollen-basiert. lateral-adapter feuert cross-layer über Layer-Identität
(beide role:adapter, Ziel-Layer != Quell-Layer) und ist KATEGORISCH — nur
adapter_sink hebt auf, nicht allow/edges (konsistent mit der Intra-Schicht-lateral;
allow/edges regieren die Schicht-Richtung via wrong-direction, nicht die
Lateral-Invariante). Die adapterSeg-Intra-Unterscheidung bleibt für den klassischen
Namen und wird später namens-generalisiert. Layer ohne Rolle sind nur kanten-geprüft.
Befund-Namen stabil.

Consequences: Beliebige/feinere Layer-Namen voll prüfbar; Rückwärtskompat 100%
(a-check-Dogfooding unverändert grün; ein einziger adapters-Layer => Cross-Layer-
Zweig feuert nie). Migration: wer zwei verschieden benannte Adapter-Layer per allow:
koppelt und beide auf role: adapter hebt, wird rot — Kopplung muss über adapter_sink
laufen. adapterSeg-Generalisierung -> welle-10b. Lastenheft 0.2.0 -> 0.3.0.
```

### 3.4 Versions-Bump

Lastenheft **0.2.0 → 0.3.0** (neue Regel-Anforderung *Schicht-Rollen* + CONF-Schema-Notiz).

## 4. Umsetzungsplan

1. `internal/adapter/driven/config/config.go` — `layers`-Decode: Glob-Liste **oder**
   `{globs, role}`. **Gotcha:** `dec.KnownFields(true)` (config.go:58) ist eine
   *Decoder*-Eigenschaft und wird von `yaml.Node.Decode(&x)` **nicht** geerbt — der
   Objekt-Zweig muss explizit strict geprüft werden (z. B. erst in
   `map[string]yaml.Node`, Schlüssel `{globs, role}` whitelisten, sonst Exit 2),
   sonst kippt die Strict-Decode-Garantie still. `role:` hat Vorrang vor der Inferenz.
2. `internal/hexagon/core/model.go` — `Layer.Role`-Feld; Rollen-Auflösung eines
   Ziel-Imports. Die Ziel-Rolle ist die Rolle **genau des von `targetLayer`
   aufgelösten Layers** (nicht erneut aus dem Import-String inferieren — sonst
   Abweichung bei Glob-Überlappung, da `targetLayer` per Substring-Präfix matcht).
3. `internal/hexagon/core/rules.go` — (a) die drei `ruleFor`-Namens-Checks →
   Rollen-Checks; (b) die **konstrukt-basierte** `port-impurity` in `Evaluate`
   (`rules.go:24-28`) ebenso rollen-basiert; (c) `lateral()` um den **Cross-Layer-Zweig**
   erweitern. Zielprädikat (beide Zweige ver-ODER-t):
   ```
   role(src)==adapter ∧ role(tgt)==adapter ∧ ¬sink(imp)
     ∧ ( tgtLayer ≠ srcLayer  ∨  adapterSeg(src) ≠ adapterSeg(imp) )
   ```
   Der Ein-`adapters`-Fall (`tgtLayer == srcLayer`, Intra-`adapterSeg`) läuft damit
   unverändert. Befund-Namen stabil.
   **Unangetastet:** `internal/adapter/driven/extract/extract.go` — die
   Konstrukt-Extraktion (`m.Forbidden[fi.Layer]`, extract.go:~76) ist **bereits
   namens-agnostisch** (greift für jeden Layer mit `forbidden_constructs`, nicht für
   das Literal `"ports"`). Nur das **Gate** in `rules.go:24` wird rollen-basiert;
   Extraktion und Gate bleiben entkoppelt.
4. Tests — Cross-Layer-`lateral` mit fremden Namen (inkl. „auch bei `allow`-Kante
   ⇒ trotzdem `lateral`", kategorisch); Inferenz-Boundary (klassische Config
   unverändert grün, **inkl.** Konstrukt-Prüfung + Intra-`adapterSeg`); Negative je
   Rolle (Import **und** Konstrukt für fremd-benannten `role: port`-Layer mit
   deklarierten `forbidden_constructs`).
5. Spezifikation/Architektur nachziehen; `make gates` grün; Multi-Linsen-Review; Commit.

## 5. Definition of Done

- [x] Neue Regel-Anforderung *Schicht-Rollen* in `spec/lastenheft.md` (3 AC + Out-of-Scope, „Generalisiert"-Zeile, kategorisches Cross-Layer-`lateral`), Bump 0.3.0 + Historie.
- [x] Folge-ADR `Accepted` + ADR-Index.
- [x] [SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung)/[SPEC-CONF-001](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema) rollen-basiert nachgezogen.
- [x] Engine dispatcht über Rolle (inkl. Konstrukt-`port-impurity` + kategorisches Cross-Layer-`lateral`); `make arch-check` (Dogfooding) **ohne** Änderung der Eigen-`.a-check.yml` grün.
- [x] Tests: fremd-benannte Rollen, Inferenz-Boundary (inkl. Konstrukt + Intra-`adapterSeg`), Negative je Rolle, Cross-Layer-`lateral` trotz `allow`.
- [x] Multi-Linsen-Review bestanden.

## 6. Offen / Risiken

- **Design-Entscheid (Review 2, Finding 1):** Cross-Layer-`lateral` ist
  **kategorisch** (nur `adapter_sink` hebt auf, nicht `allow`/`edges`). Begründung:
  `lateral-adapter` ist eine kategorische Invariante; `allow`/`edges` regieren die
  Schicht-Richtung (`wrong-direction`), nicht die Lateral-Kopplung. **Migration:**
  eine zuvor per `allow:` grüne Adapter→Adapter-Kante wird nach `role: adapter`-Optin
  rot — Kopplung über `adapter_sink` führen. *(Beim Abnehmen auf „kanten-regiert" kippbar.)*
- **Blocker-Fix (Review 1):** `lateral-adapter` Cross-Layer (Layer-Identität) gehört
  in Scope a — sonst Happy-AC unerfüllbar; nur die `adapterSeg`-Namens-Generalisierung
  (Intra) ist R6 → 10b.
- **Vierte Namens-Bindung (Review 1):** konstrukt-basierte `port-impurity`
  (`Evaluate`) mitgeneralisiert; `extract.go` bleibt unangetastet (§4 Schritt 3).
- **Rückwärtskompat DoD-kritisch:** Eigen-`.a-check.yml` (`core`/`ports`/`adapters`)
  bleibt **ohne** Änderung grün (Inferenz; ein `adapters`-Layer ⇒ Cross-Layer feuert nie).
- **Beobachtung:** `adapterSeg` ist im Eigen-Repo aktuell inert (Pfad
  `adapter/driven`, nicht `adapters`) — die Intra-Prüfung ist dort ein No-Op (10b).
- **Befund-Namen bleiben stabil** (`core-impurity` etc.) — keine Output-Brüche.
- **R5** (Boundary-AC-Nuance aus dem Erst-Review zu (I)) optional mitnehmen.

## 7. Closure-Notiz (nach `done/`)

**Belege:** `make gates` grün — `lint`/`test`/`coverage-gate`; `arch-check` 0 Befunde
(a-check prüft sich selbst rollen-basiert via Inferenz, **ohne** Änderung der
Eigen-`.a-check.yml`); `doc-check` 0. [AC-FA-RULE-006](../../../../spec/lastenheft.md#ac-fa-rule-006--schicht-rollen-generische-regel-anwendung)
(Lastenheft 0.3.0) + [ADR-0009](../../adr/0009-rollen-basierter-regel-dispatch.md)
`Accepted`; Spezifikation 0.3.0 rollen-basiert. Engine: `config.go` (Union-Decode
+ Rolle, yaml-Gotcha), `model.go` (`Layer.Role`), `rules.go` (Rollen-Dispatch +
kategorisches Cross-Layer-`lateral` + Konstrukt-`port-impurity`). Multi-Linsen-Review
(4 Linsen) bestanden, T1–T4/K1 geschlossen.

**Lerneintrag:**

- *Inferenz in `roleOf`, nicht im Config-Adapter:* direkt gebaute Test-Modelle (ohne
  Decode) und Bestands-Configs greifen gleichermaßen — Rückwärtskompat ohne Sonderfall.
- *Differenzialer Test schlägt Präsenz-Test:* der erste „kategorisch"-Test war
  tautologisch (`hasRule`); erst `len==1` + `allow`/`edge`-Tabelle beweist die Invariante.
- *yaml-Gotcha:* `KnownFields(true)` erbt nicht durch `yaml.Node.Decode` — strikte
  Schlüsselprüfung im Objekt-Zweig von Hand.

**Folge (späteres Inkrement):** `app`-Rolle, `driving`/`driven`-Ports,
`adapterSeg`-Namens-Generalisierung (R6), `targetLayer` längster-Präfix-Match.
