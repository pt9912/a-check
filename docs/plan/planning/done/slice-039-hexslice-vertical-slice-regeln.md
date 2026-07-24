# slice-039 — HexSlice Vertical-Slice-Regeln: `lateral-slice` + `port-locality`

**Status:** **done** (2026-07-24) — spec-first umgesetzt, adversarisches Review (b-cad-Regression gefunden+gefixt), `make ci` grün, gegen HexSlice-Beispiel + b-cad/d-check/d-migrate verifiziert. Zwei neue Regeln.
**Auslöser:** Nutzer-Wahl [slice-037](../open/slice-037-hexslice-gap-analyse.md) **Option B + C** — die Vertical-Slice-Achse gaten. Evidenz: realer HexSlice-Go-Konsument `hexslice-architecture/lab/examples/go` (mit `.a-check.yml` + `a-check.mk`).
**Bezug:** neue Regeln unter [AC-FA-RULE-006](../../../../spec/lastenheft.md#ac-fa-rule-006--schicht-rollen-generische-regel-anwendung)-Rollenmodell; präzisiert [SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung); Heuristik-Grenze [AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze). [Roadmap](../in-progress/roadmap.md).

> **Hinweis:** Entwurf/Umsetzung. AC-/ADR-IDs werden beim Schreiben in `spec/` vergeben
> (Anlege-Prozess: [conventions](../../../../harness/conventions.md#anforderungs-anlege-prozess)).

---

## 1. Ziel

HexSlices Unterscheidungsmerkmal gegenüber „gewöhnlicher Hexagonal" gaten:

- **B — `lateral-slice`** ([AC-FA-RULE-009](../../../../spec/lastenheft.md#ac-fa-rule-009--slice-isolation-regel-lateral-slice)): eine `app`-Datei importiert eine **fremde Use-Case-Slice** → Befund. Kategorisch (wie `lateral-adapter` für Adapter, [AC-FA-RULE-002](../../../../spec/lastenheft.md#ac-fa-rule-002--keine-lateralen-adapter-kanten-regel-lateral-adapter)).
- **C — `port-locality`** ([AC-FA-RULE-010](../../../../spec/lastenheft.md#ac-fa-rule-010--port-lokalität-regel-port-locality)): eine `app`-Datei importiert einen `port`, dessen **Scope-Verzeichnis** sie nicht enthält → Befund. Deckt HexSlice-Regel 5 „so lokal wie möglich, so gemeinsam wie nötig" (use-case-lokal ⊂ business-area ⊂ app-weit).

## 2. Evidenz & Verifikation (gegen `a-check:dev`, alle real gelaufen)

Der Beispiel-Baum: `internal/hexagon/application/order/{createorder,cancelorder}/` (Slices, je `ports/`),
`.../application/order/ports/` (business-area-shared), `.../domain/order/`, `internal/adapters/{inbound,outbound}/`.

- **Compliant heute:** 0 Befunde.
- **Gap bestätigt:** injizierte `createorder → cancelorder/ports`-Kopplung → weiter 0 (C ungegatet); app→app ebenso (B ungegatet).
- **Prerequisite-Befund (Fundament):** die **Ziel-Auflösung** `layerOfCand` matcht nur **saubere literale Glob-Präfixe** (`segIndex`). Die Beispiel-Globs `…/createorder/*.go` und `application/**/ports/**` haben **kein** sauberes Präfix → app→app/app→port lösen auf **extern** auf → die Slice-Kanten sind **heute gar nicht erzwungen** (Beweis: `app→ports`-Kante entfernt ⇒ 0 statt `wrong-direction`).
- **Weg (Nutzer-Entscheid: Config-Disziplin):** mit **per-Slice-/per-Port-Dir-`**`-Globs** (saubere Präfixe) lösen die Ziele auf — verifiziert: `app→ports`-Kante entfernt ⇒ **4× `wrong-direction`**; volle Config ⇒ 0. Ports schlagen `app` per **längerem** Präfix (kein §4.0-Tie; slice-038 orthogonal). **Kein Engine-Change.** Grenze wird dokumentiert ([AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)); die korrigierte Beispiel-Config wird mitgeliefert.

## 3. Design (entschieden)

Beide Regeln sind **pfad-heuristisch** und nutzen den vorhandenen `segIndex` (modul-präfix-tolerant:
findet einen Segment-Run auch im modul-qualifizierten Kandidaten `hexslice/example/internal/…`).

### 3.1 Slice-Identität (B) — Glob-pro-Slice

`sliceOf(path)` = das **längste `app`-Rollen-Glob-Literalpräfix** `p` mit `segIndex(path, p) ≥ 0`;
Rückgabe `p` (kanonisch, repo-relativ). Zwei `app`-Dateien mit **verschiedenem** `sliceOf` liegen in
verschiedenen Slices.

- **Regel `lateral-slice`:** `srcRole == app && tgtRole == app && sliceOf(srcPath) != sliceOf(cand)` → Befund. **Kategorisch** (nicht über `edges`/`allow` aufhebbar). Kein Sink in v1 (Slices teilen über Ports, nicht über app-Code — geteilter app-Code = Out-of-Scope, künftiges `app_sink`).
- Gleiche Slice (auch Sub-Pakete `…/createorder/sub/x`) → `sliceOf` identisch → kein Befund.
- Ein **einziges** breites `app`-Glob (`application/**`) ⇒ alle app-Dateien dieselbe Slice ⇒ `lateral-slice` inert (rückwärtskompatibel; Slice-Isolation ist opt-in über per-Slice-Globs).

### 3.2 Port-Scope (C) — pfad-basiert

`portScope(cand)` = das **längste `port`-Rollen-Glob-Literalpräfix**, das `cand` matcht, **minus seinem
letzten Pfad-Segment** (dem Port-Dir-Marker, typisch `ports`). Rückgabe repo-relativ, kanonisch.

- `…/createorder/ports/**` → Präfix `…/createorder/ports` → Scope `…/createorder`.
- `…/order/ports/**` → Scope `…/order` (business-area).
- `…/application/ports/**` bzw. `hexagon/ports/**` → Scope `…/application` bzw. `hexagon` (app-weit).
- **Regel `port-locality`:** `srcRole == app && tgtRole == port && segIndex(srcPath, portScope) < 0` → Befund. **Kategorisch.** Nur **app**-Importeure (Adapter→Port = Implementierungs-Beziehung, edge/direction-regiert, **nicht** erfasst). `composition_root` global ausgenommen.

### 3.3 Erst-Treffer-Reihenfolge

`… → lateral-adapter → lateral-slice → tech-leak → port-direction-mismatch → port-locality → wrong-direction`.
`lateral-slice` nach `lateral-adapter` (analoge kategorische Klasse, app-Seite); `port-locality` **vor**
`wrong-direction` (ein erlaubtes `app→port` per Kante darf die Lokalitäts-Verletzung nicht maskieren).

### 3.4 Config-Schema

**Keine Schema-Änderung.** Beide Regeln leiten aus vorhandenen `layers`-Globs + Pfaden ab. Voraussetzung
(dokumentiert): `app`- und `port`-Globs mit **sauberen literalen Präfixen** (per-Slice/per-Port-Dir `**`),
sonst lösen die Ziele nicht auf (Prerequisite-Befund §2, [AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)).

## 4. Umsetzungsplan (spec-first)

1. **Lastenheft:** [AC-FA-RULE-009](../../../../spec/lastenheft.md#ac-fa-rule-009--slice-isolation-regel-lateral-slice) (`lateral-slice`) + [AC-FA-RULE-010](../../../../spec/lastenheft.md#ac-fa-rule-010--port-lokalität-regel-port-locality) (`port-locality`), je 3 AK (Happy/Boundary/Negative) + Out-of-Scope; Bump 0.20.0 → **0.21.0** + Historie.
2. **[ADR-0026](../../adr/0026-hexslice-vertical-slice-regeln.md)** (Accepted, schärft [SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung)): Design-Entscheide — Glob-pro-Slice (B), pfad-basierter Port-Scope (C), Config-Disziplin-Prerequisite (verworfen: `layerOfCand`-Voll-Glob, Modul-Präfix-Bruch-Risiko), Erst-Treffer-Platzierung. ADR-Index.
3. **[SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung):** zwei Regeln + `sliceOf`/`portScope`-Definition + Erst-Treffer-Kette; Bump 0.21.0 → **0.22.0** + Historie.
4. **Code** (`rules.go`): `sliceOf`, `portScope`, zwei Dispatch-Arme in `ruleFor`.
5. **Tests:** AC-Tests je Regel (happy/boundary/negative, kategorisch) + **Integrations-Fixture** aus dem Go-Beispiel (positiv 0 Befunde + injizierte Cross-Slice-/Cross-Scope-Verstöße).
6. **Benutzerhandbuch:** neue Regeln + Config-Disziplin (saubere Präfix-Globs).
7. **Korrigierte Beispiel-Config** für `hexslice-architecture/lab/examples/go` mitliefern (saubere Präfix-Globs) — separater Hinweis, anderes Repo.
8. `make gates` + `make ci`; adversarisches Review; Verifikation gegen das Go-Beispiel; Closure.

## 5. Definition of Done

- [x] Lastenheft [AC-FA-RULE-009](../../../../spec/lastenheft.md#ac-fa-rule-009--slice-isolation-regel-lateral-slice) + [AC-FA-RULE-010](../../../../spec/lastenheft.md#ac-fa-rule-010--port-lokalität-regel-port-locality) (je 3–4 AK + Out-of-Scope) + Bump 0.21.0/Historie.
- [x] [ADR-0026](../../adr/0026-hexslice-vertical-slice-regeln.md) Accepted + Index; [SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung) erweitert (`sliceOf`/`portScope`/Erst-Treffer) + Bump 0.22.0/Historie.
- [x] `rules.go`: `sliceOf`/`portScope`/`lateralSlice`/`portLocality` + Dispatch (`connectivityRule` in `lateralRule`/`directionRule` gesplittet, gocyclo-konform); **kategorisch**.
- [x] AC-Tests grün (11: happy/boundary/negative/kategorisch je Regel, Adapter-nicht-erfasst, **+2 Review-Regressionstests** klassisches Hexagonal); Go-Beispiel end-to-end (positiv 0; injiziert → `lateral-slice`+`port-locality`).
- [x] **Adversarisches Implementierungs-Review — Regression gefunden + gefixt:** die erste Fassung brach **b-cad** mit 27 Falsch-Positiven (sibling-Ports → `port-locality`; getrennte app-Layer `services`/`services_geo` mit Kante → `lateral-slice`). Fix: `lateral-slice` nur **same-app-layer** (`tl == f.Layer`), `port-locality` nur bei **`appTreeContains(portScope)`** (geschachtelte Ports). Verifiziert: b-cad/d-check/d-migrate wieder **0**, HexSlice-Beispiel fängt weiter. Spec/Lastenheft/ADR nachgezogen.
- [x] `make arch-check` **0**; `make gates` + `make ci` grün.
- [x] Benutzerhandbuch: **Aufgabe §3.7** „HexSlice absichern" (Arbeitsanleitung + Beispiel-Config, nach [`benutzerhandbuch-standard.md`](../../../../docs/user/benutzerhandbuch-standard.md)) + §3.4-Regeltabelle (verfeinert) + Config-Disziplin-Kasten + Glossar, 1.31.
- [ ] **Korrigierte Beispiel-Config** (saubere Präfix-Globs) — anderes Repo, dem Nutzer vorgelegt (Anwendung auf Wort).
- [ ] Merge/Closure + Lerneintrag.

## 6. Offen / Grenzen

- **Config-Disziplin ist Voraussetzung** (saubere Präfix-Globs) — als [AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)-Grenze dokumentiert, nicht still.
- **Geteilter app-Code zwischen Slices** (nicht über Ports) — Out-of-Scope v1; künftiges `app_sink` analog `adapter_sink`.
- **Adapter→Port-Lokalität** bewusst **nicht** erfasst (Implementierungs-Beziehung); nur app-Importeure.
- **`layerOfCand`-Voll-Glob-Auflösung** (Original-`*.go`-Config ohne Umschreiben) — verworfen (Modul-Präfix-Bruch), eigener Slice falls je nötig.
- **Graph-Legende (`--print-graph`)** nennt die kategorischen Regeln noch als `core-impurity`/`lateral-adapter`/`port-direction-mismatch` — `lateral-slice`/`port-locality` fehlen dort (Renderer + [SPEC-CLI-002](../../../../spec/spezifikation.md#spec-cli-002--graph-renderer-vertrag) unberührt). **Deferred Folge** (kleiner Renderer-/Test-Nachzug), nicht in slice-039.

## 7. Closure-Notiz (2026-07-24)

- **Umgesetzt:** Lastenheft 0.21.0 ([AC-FA-RULE-009](../../../../spec/lastenheft.md#ac-fa-rule-009--slice-isolation-regel-lateral-slice)/[AC-FA-RULE-010](../../../../spec/lastenheft.md#ac-fa-rule-010--port-lokalität-regel-port-locality)), [ADR-0026](../../adr/0026-hexslice-vertical-slice-regeln.md) Accepted, [SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung) 0.22.0, `rules.go`
  (`sliceOf`/`portScope`/`appTreeContains`/`lateralSlice`/`portLocality` + `connectivityRule`-Split),
  11 `core`-Tests, Benutzerhandbuch 1.31 (Aufgabe §3.7 + Regeltabelle + Config-Kasten + Glossar).
- **Adversarisches Review = echter Fund:** die erste Fassung (kategorisch über alle app-Layer; jede
  app→port-Lokalität) brach **b-cad** mit 27 Falsch-Positiven. Ursache: (1) `lateral-slice` traf getrennte
  app-Layer mit deklarierter Kante (`services`→`services_geo`); (2) `port-locality` traf sibling-Ports
  (klassisch, Ports neben `services`). Fix: `lateral-slice` **same-app-layer**, `port-locality` nur bei
  **`appTreeContains`**. Regressionstests + reale Verifikation (b-cad/d-check/d-migrate = 0).
- **Verifikation:** `make ci` grün, `arch-check` 0; HexSlice-Beispiel (compliant 0; injiziert →
  `lateral-slice`+`port-locality`) gegen `a-check:dev`.
- **Lerneintrag:** eine neue kategorische Regel gegen **nur** die Motiv-Topologie (nested HexSlice) zu
  testen ist eine Falle — sie muss gegen die **Bestands-Topologien** der realen Konsumenten (klassisches
  Hexagonal, sibling-Ports, Sub-Layer mit Kante) gegengeprüft werden, sonst schlägt sie dort als
  Falsch-Positiv zu. Opt-in-Regeln brauchen einen expliziten „greift-nur-wenn"-Guard (`tl == f.Layer`,
  `appTreeContains`), nicht bloß eine „feuert-wenn-anders"-Bedingung.
- **Offen:** Beispiel-Config-Anwendung (anderes Repo, §6); Graph-Legende-Nachzug (deferred, §6).
