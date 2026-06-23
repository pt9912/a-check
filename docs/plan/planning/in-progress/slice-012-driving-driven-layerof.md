# slice-012 — Driving/Driven-Richtung + `LayerOf` längster-Präfix (welle-10b / b2b)

**Status:** in-progress (Entwurf zur Abnahme).
**Welle:** welle-10-regel-engine-generalisierung (Inkrement **b2b**, Abschluss).
**Bezug:** Re-Evaluierungs-Trigger aus [ADR-0010](../../adr/0010-layer-relativer-adapterseg-laengster-praefix.md) (`LayerOf`) und [ADR-0011](../../adr/0011-domain-application-trennung-rolle-app.md) (`driving`/`driven`-Ports); verfeinert den Rollen-Mechanismus [AC-FA-RULE-006](../../../../spec/lastenheft.md#ac-fa-rule-006--schicht-rollen-generische-regel-anwendung); [Roadmap welle-10](roadmap.md).

> **Hinweis:** Dieses Dokument hält den **Entwurf** zur Abnahme. Die in §3 als
> Vorlage in Code-Fences gesetzten Anforderungs-/ADR-Texte sind unverbindlich —
> gültig erst nach Freigabe in [`spec/lastenheft.md`](../../../../spec/lastenheft.md)
> bzw. neuen ADRs. Die DoD-Haken in §5 sind offen; die Beschlüsse in §6 sind
> getroffen (zur Freigabe) — die Umsetzung folgt dem Plan in §4.

---

## 1. Ziel

Den Rollen-Mechanismus um die **primär/sekundär-Unterscheidung** der Hexagon-Ports
(`driving`/`driven`) erweitern und eine über die Wellen gewachsene
Auflösungs-Inkonsistenz schließen. Zwei trennbare Teile:

- **Teil A — `driving`/`driven`-Richtung (groß):** eine optionale Richtung
  `direction ∈ {driving, driven}` auf `port`-/`adapter`-Schichten, **orthogonal** zur
  Rolle. `driving` = primär/inbound (Use-Case-Schnittstelle, vom Treiber-Adapter
  aufgerufen); `driven` = sekundär/outbound (vom Kern/App definiert, vom getriebenen
  Adapter implementiert). Neue Regel: ein Adapter spricht nur Ports **seiner**
  Richtung.
- **Teil B — `LayerOf` längster-Präfix (klein):** die Schicht-Zuordnung einer Datei
  (`LayerOf`) an die `targetLayer`-Auflösung angleichen (spezifischster Glob gewinnt),
  damit verschachtelte Schichten konsistent klassifiziert werden.

So modellieren die Konsumenten (b-cad/d-check/d-migrate mit getrennten
`driving`/`driven`-Port-Modulen) ihre Treiber/Getriebenen-Trennung voll, ohne die
Engine zu forken.

## 2. Problem

1. **Keine Richtungs-Dimension.** Nach [AC-FA-RULE-007](../../../../spec/lastenheft.md#ac-fa-rule-007--rolle-app-und-strenge-domain)
   trägt eine Schicht eine Rolle ∈ {`domain`, `app`, `port`, `adapter`} — aber alle
   Ports sind gleich. Ein **Treiber**-Adapter (CLI/HTTP), der direkt einen
   **driven**-Port (Repository/Filesystem) importiert, statt über die `app`-Schicht
   zu gehen, ist ein Architektur-Bruch, den heute **keine** Regel fängt
   (`lateral-adapter` greift nur Adapter→Adapter; `wrong-direction` ist
   kanten-geregelt und kann per Kante erlaubt werden).
2. **`LayerOf` ≠ `targetLayer`.** `LayerOf` (eigene Schicht einer Datei) nimmt den
   **ersten** passenden Glob (Regex-Match, `rules.go` `LayerOf`); `targetLayer`
   (Import-Ziel) nimmt den **längsten** Präfix ([ADR-0010](../../adr/0010-layer-relativer-adapterseg-laengster-praefix.md)).
   Bei verschachtelten Schicht-Globs (`src/app/**` ⊂ `src/**`) können beide
   abweichen — dieselbe Datei wird als Quelle anders eingeordnet denn als Ziel.

## 3. Entwurf (zur Abnahme)

### 3.1 Neue Anforderung — Driving/Driven-Richtung (Teil A)

```text
AC-FA-RULE-008 — Driving/Driven-Port-Richtung (Regel port-direction-mismatch)
Verfeinert: AC-FA-RULE-006 (Rollen-Mechanismus) um eine ORTHOGONALE Richtungs-Dimension.

Beschreibung: Eine port- oder adapter-Schicht traegt optional eine Richtung
direction in {driving, driven}. driving = primaer/inbound (Use-Case-Schnittstelle,
vom Treiber-Adapter aufgerufen); driven = sekundaer/outbound (vom Kern/App definiert,
vom getriebenen Adapter implementiert). Die Richtung ist ORTHOGONAL zur Rolle: die
Reinheits-Regeln (core-/app-/port-impurity, lateral-adapter) bleiben rollen-basiert
unveraendert.

Neue Regel port-direction-mismatch: ein role: adapter mit direction X, der eine
role: port-Schicht mit direction Y (Y != X, beide gesetzt) importiert, ist ein
Befund -- ein Treiber-Adapter spricht nur driving-Ports, ein getriebener Adapter nur
driven-Ports. Schichten OHNE direction unterliegen der Regel NICHT (Rueckwaerts-
Kompatibilitaet: ohne direction-Deklaration aendert sich nichts). Die app-Schicht ist
richtungs-agnostisch (nutzt driven-Ports, implementiert driving-Ports) und wird nicht
erfasst.

Rollen+Richtung -> Befund (Ergaenzung): adapter(X) -> port(Y!=X) => port-direction-mismatch.

Akzeptanzkriterien:
- Happy: Given ein role: adapter, direction: driving, when er eine role: port,
  direction: driving-Schicht importiert, then kein Befund.
- Negative: Given ein role: adapter, direction: driving, when er eine role: port,
  direction: driven-Schicht importiert, then Befund port-direction-mismatch, Exit 1.
- Boundary: Given Schichten OHNE direction (klassisch role: port/adapter), when
  a-check laeuft, then identisches Verhalten wie 0.5.0.

Out-of-Scope: Auto-Inferenz der Richtung aus Namen (driving/driven im Pfad);
Richtungs-Regeln zwischen Ports untereinander -- spaeteres Inkrement.
```

### 3.2 Schema-Erweiterung (Config, Teil A)

Erweitert [AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml)/[SPEC-CONF-001](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema):
das `layers`-Objekt (`{globs, role}`) erhält einen optionalen Schlüssel `direction`:

```yaml
layers:
  cli:    {globs: ["src/driving/cli/**"],   role: adapter, direction: driving}
  api:    {globs: ["src/driving/http/**"],  role: port,    direction: driving}
  repo:   {globs: ["src/driven/db/**"],     role: adapter, direction: driven}
  store:  {globs: ["src/driven/ports/**"],  role: port,    direction: driven}
```

Strict-Decode bleibt: Whitelist `{globs, role, direction}`, `direction ∈ {driving, driven}`, sonst Exit 2.

### 3.3 Folge-ADRs

```text
ADR-0012 — Driving/Driven-Richtung als orthogonale Schicht-Dimension
Status: Proposed (-> Accepted bei Merge).
Bezug: AC-FA-RULE-008 (neu), AC-FA-RULE-006. Schaerft: SPEC-RULE-001 + SPEC-CONF-001.
Decision: optionale direction in {driving, driven} auf port-/adapter-Schichten,
orthogonal zur Rolle; neue Regel port-direction-mismatch (Adapter-Richtung != Ziel-
Port-Richtung). Ohne direction keine Pruefung (rueckwaertskompatibel). Alternative
subtype-Rollen (port_driving/...) verworfen: blaeht Rollen-Enum + jede Reinheits-
Pruefung auf; die orthogonale Dimension ist sparsamer. Lastenheft 0.5.0 -> 0.6.0.

ADR-0013 — LayerOf laengster-Praefix (Angleichung an targetLayer)
Status: Proposed (-> Accepted bei Merge).
Bezug: AC-QA-01 (Determinismus). Schaerft: SPEC-RULE-001.
Decision: LayerOf nimmt die spezifischste (laengster-Glob-Praefix) passende Schicht
statt des Erst-Treffers, konsistent mit targetLayer; bei Gleichstand die zuerst
deklarierte. Verhaltensaenderung NUR bei verschachtelten Schicht-Globs.
```

### 3.4 Versions-Bump

Lastenheft **0.5.0 → 0.6.0** (Teil A: neue Anforderung *Driving/Driven-Richtung*). Teil B ist eine Engine-/Spec-Konsistenz ohne neue Anforderung.

## 4. Umsetzungsplan

### 4.1 Teil A — Richtung
1. `internal/adapter/driven/config/config.go` (`decodeLayer`) — Whitelist
   `{globs, role, direction}`; `yamlLayer.Direction`; Validierung `direction ∈
   {driving, driven}` (sonst Exit 2); fail-closed. Rückgabe um die Richtung erweitern
   (heute `([]string, string, error)`) → liefert zusätzlich `direction`; der Aufrufer
   setzt `core.Layer{… Direction: direction}`.
2. `internal/hexagon/core/model.go` — `Layer.Direction string`.
3. `internal/hexagon/core/rules.go` (`ruleFor`) — neue `case` in den Switch, **vor**
   `wrong-direction` (so gewinnt die spezifischere Regel, falls die `adapter→port`-Kante
   nicht in `edges`/`allow` steht und sonst `wrong-direction` zöge):
   ```
   srcRole==adapter ∧ tgtRole==port ∧ dirOf(f.Layer)≠"" ∧ dirOf(tl)≠"" ∧ dirOf(f.Layer)≠dirOf(tl)
     ⇒ port-direction-mismatch
   ```
   `dirOf(name) = layerByName(name, m).Direction` — **ohne** Namens-Inferenz (anders als
   `roleOf`; Auto-Inferenz ist out-of-scope, §3.1). Als Prädikat-Helfer
   `directionMismatch(m, f.Layer, tl)` auslagern (gocyclo). Dokumentierte Regel-Kette:
   `… → tech-leak → port-direction-mismatch → wrong-direction` ([SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung)). Befund-
   Namen sonst stabil.
4. Tests: driving→driving happy; driving→driven ⇒ `port-direction-mismatch` (kategorisch,
   `len==1`); Boundary (ohne `direction` unverändert); fremde Namen.

### 4.2 Teil B — `LayerOf`
5. `internal/hexagon/core/rules.go` (`LayerOf`) — Match-Prädikat bleibt `matchesAny`
   (volle Glob-Semantik für echte Pfade, inkl. `**`); **nur die Auswahl** unter mehreren
   Treffern wechselt vom Erst-Treffer auf den längsten `globPrefix` (Helfer `globPrefix`
   wiederverwendet, konsistent mit `targetLayer`/`adapterSeg`). Globs ohne literalen
   Präfix (`**/…`) haben Spezifität 0 und verlieren gegen jeden Präfix-Treffer. Tie-Break:
   zuerst deklariert. Tests: verschachtelte Globs (`src/app/**` ⊂ `src/**`) → Datei landet
   in der spezifischsten Schicht; Gleichstand → zuerst deklariert.

### 4.3 Abschluss
6. `spec/lastenheft.md` (neue Anforderung *Driving/Driven-Richtung*, 0.6.0), `spec/spezifikation.md`
   ([SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung):
   `port-direction-mismatch`-Zeile + Erst-Treffer-Reihenfolge; Schema-Enum `direction`),
   ADR-Index; „sechs→sieben"-Befund-Sweep (analog b2a: `rules.go`-Comment, beide Specs,
   Benutzerhandbuch, README, CHANGELOG); `make gates`; Multi-Linsen-Review; Commit.

## 5. Definition of Done

- [ ] Neue Anforderung *Driving/Driven-Richtung* in `spec/lastenheft.md` (3 AC + Out-of-Scope), Bump 0.6.0 + Historie.
- [ ] Folge-ADRs (Richtung + `LayerOf`) `Proposed → Accepted` + ADR-Index; [SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung)/[SPEC-CONF-001](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema) nachgezogen.
- [ ] `config.go`: `direction`-Whitelist; `model.go`: `Layer.Direction`; `rules.go`: `port-direction-mismatch` + `LayerOf` längster-Präfix.
- [ ] „Sechs→sieben"-Sweep vollständig (Regel-Zählungen/Befund-Listen) — vollständig wie der b2a-Sweep.
- [ ] Engine: `port-direction-mismatch` (kategorisch); `make arch-check` (Dogfooding) **ohne** Änderung der Eigen-`.a-check.yml` grün (a-check hat keine `direction`).
- [ ] Tests: Richtung happy/Mismatch/Boundary; `LayerOf` verschachtelt.
- [ ] Multi-Linsen-Review bestanden.

## 6. Beschlüsse zur Abnahme

Die §6-Entscheidungen sind getroffen (Stand: zur Freigabe). Begründungen knapp:

- **Entscheid 0 — Bedarf (Gate): PROCEED mit Inertheits-Sicherung.** Die Richtung ist
  *opt-in und inert ohne `direction`* (Boundary-AC §3.1 + Rückwärtskompat-DoD §5): ohne
  Deklaration ändert sich für a-check und alle Bestands-Konsumenten **nichts**, das
  „spekulativ"-Risiko ist damit gekappt (geliefert-aber-ungenutzt schadet nicht).
  Begründet über die Roadmap-Einordnung (welle-10b, b2b-Einstiegspunkt) und die
  Konsumenten b-cad/d-check/d-migrate mit getrennten `driving`/`driven`-Port-Modulen (§1).
  ⚠️ *Offen für den Freigebenden:* mindestens ein Konsument soll die Trennung tatsächlich
  aktivieren wollen — sonst bleibt Teil A bewusst inert ausgeliefert.
- **Entscheid A — Modellierung: Attribut `direction`** (orthogonal zur Rolle). Subtyp-
  Rollen (`port_driving`/…) verworfen — blähen Rollen-Enum + jede Reinheits-Prüfung auf;
  die orthogonale Dimension lässt die Reinheits-Regeln unberührt ([ADR-0012](../../adr/0012-driving-driven-richtung-orthogonale-dimension.md)).
- **Entscheid B — Regel-Umfang: minimal `adapter→port`.** Port→Port-Richtungsregeln
  ausdrücklich out-of-scope (späteres Inkrement, §3.1).
- **Entscheid C — Befund-Name: `port-direction-mismatch`** (7. Befund). Doku-Sweep
  „sechs→sieben" zwingend (§4.3); Output-Konsumenten (CI-Parser) im Sweep mitgeführt.
- **Entscheid D — Slice-Schnitt: A+B gebündelt** in slice-012 (Freigabe-Anweisung).
  Teil B bleibt im Code unabhängig (eigene Commits möglich, falls gewünscht).
- **Rückwärtskompat (DoD-kritisch):** ohne `direction` (Eigen-`.a-check.yml`) bleibt
  alles grün; die `LayerOf`-Änderung greift nur bei verschachtelten Globs (a-check:
  keine) → `make arch-check` unverändert.

## 7. Closure-Notiz (nach `done/`)

_(wird beim Abschluss gefüllt: `make gates`-Beleg, `arch-check` 0 unverändert,
Review-Runden, Lerneintrag; schließt welle-10 ab.)_
