# slice-011 — `app`-Rolle + strenge `domain` (welle-10b / b2a)

**Status:** done.
**Welle:** welle-10-regel-engine-generalisierung (Inkrement **b2a**).
**Bezug:** Re-Evaluierungs-Trigger aus [ADR-0009](../../adr/0009-rollen-basierter-regel-dispatch.md)/[ADR-0010](../../adr/0010-layer-relativer-adapterseg-laengster-praefix.md); erweitert den Rollen-Mechanismus [AC-FA-RULE-006](../../../../spec/lastenheft.md#ac-fa-rule-006--schicht-rollen-generische-regel-anwendung) und schärft [AC-FA-RULE-001](../../../../spec/lastenheft.md#ac-fa-rule-001--kern-reinheit-regel-core-impurity); [Roadmap welle-10](../in-progress/roadmap.md).

> **Hinweis:** Umgesetzt, reviewt und abgenommen (2026-06-22). Die in §3 als Vorlage
> gesetzten Anforderungs-/ADR-Texte sind nun verbindlich in
> [`spec/lastenheft.md`](../../../../spec/lastenheft.md) und
> [ADR-0011](../../adr/0011-domain-application-trennung-rolle-app.md). Dieses Dokument
> hält Entwurf + Closure (§7).

---

## 1. Ziel

Das Rollen-Modell {`domain`, `port`, `adapter`} um eine **`app`-Rolle**
(Application/Use-Case-Schicht) erweitern und damit die **Domain/Application-Trennung**
durchsetzen: Use-Cases orchestrieren über Ports, die pure `domain` kennt **keine**
Ports. So modellieren die Konsumenten mit feinerem Hexagon (b-cad/d-migrate:
`domain`/`application`/`port`) ihre Struktur voll, ohne die Engine zu forken.

- **Scope b2a (dieser Slice):** Rolle `app` (darf `domain`+`port`, **nicht**
  Adapter/Tech → neuer Befund `app-impurity`); Rolle `domain` **verschärft**
  (importiert nur `domain` → `domain↛port` wird **kategorisch**, nicht mehr
  kanten-geregelt); Namens-Inferenz `application`/`app` → `app`.
- **Scope b2b (Folge, eigener Slice):** `driving`/`driven`-Port-Subtypen mit
  feineren Kanten; `LayerOf` längster-Präfix.

## 2. Problem

Nach [ADR-0009](../../adr/0009-rollen-basierter-regel-dispatch.md) trägt eine
Schicht eine Rolle ∈ {`domain`, `port`, `adapter`}. Zwei Lücken bleiben:

1. **Keine `app`-Schicht.** Wer Use-Cases von der puren Domäne trennt (eigene
   `application`-Schicht), kann sie nicht als das modellieren, was sie ist: rein
   gegenüber Technik, aber **berechtigt**, Ports zu nutzen. Heute bleibt nur
   `port` (zu eng — App ist kein Port) oder roleless (gar keine Reinheit).
2. **`domain↛port` ist nur kanten-geregelt.** Importiert eine `role: domain`-Schicht
   eine `port`-Schicht, feuert heute höchstens `wrong-direction` — und eine
   deklarierte `domain→port`-Kante **hebt das auf**. Die Invariante „Domain kennt
   keine Ports" lässt sich so aushebeln; sie ist nicht durchgesetzt.

## 3. Entwurf (zur Abnahme)

### 3.1 Neue Anforderung — Domain/Application-Trennung

```text
AC-FA-RULE-007 — Domain/Application-Trennung (Rolle app + strenge domain)
Erweitert: AC-FA-RULE-006 (Rollen-Menge um app). Schärft: AC-FA-RULE-001 (core-impurity).

Beschreibung: Das Rollen-Modell aus AC-FA-RULE-006 wird um die Rolle app
(Application/Use-Case-Schicht) erweitert; die Rolle domain wird verschärft.

- Rollen-Menge: {domain, app, port, adapter}. Namens-Inferenz ergaenzt
  application->app und app->app (zusaetzlich zu core->domain, ports->port,
  adapters->adapter). Explizite role: hat weiter Vorrang.
- Rolle app: darf domain UND port importieren (Use-Cases orchestrieren ueber
  Ports), aber KEINE Adapter-/Tech-Typen. Verstoss => Befund app-impurity (neu).
  Die Schicht-Richtung (app->domain, app->port) bleibt kanten-geregelt
  (wrong-direction); die Reinheit (kein Adapter/Tech) ist KATEGORISCH.
- Rolle domain (verschaerft): die innerste Schicht ist die strengste. Jeder Import
  auf eine app-, port- ODER adapter-Schicht ODER ein Tech-Muster => Befund
  core-impurity, KATEGORISCH (auch bei deklarierter Kante). Rollenlose Ziel-Schichten
  bleiben kanten-geregelt (kein kategorischer Befund). Bisher war domain->port nur
  kanten-geregelt (wrong-direction); jetzt ist die Domain-Reinheit eine harte
  Invariante: "Domain kennt keine Ports."

Rollen->Befund: domain->core-impurity, app->app-impurity (neu),
port->port-impurity, adapter->lateral-adapter.

Akzeptanzkriterien:
- Happy: Given eine Schicht role: app mit deklarierten Kanten app->domain und
  app->port, when sie eine domain- und eine port-Schicht importiert, then KEIN
  Befund.
- Negative-app: Given eine Schicht role: app, when sie eine adapter-Schicht ODER
  ein Tech-Muster importiert, then Befund app-impurity, Exit 1 (kategorisch, auch
  bei deklarierter Kante).
- Negative-domain: Given eine Schicht role: domain, when sie eine port- (oder app-/
  adapter-)Schicht importiert (auch MIT deklarierter Kante), then Befund
  core-impurity, Exit 1.
- Boundary: Given eine Config OHNE role: und OHNE Layer namens application/app
  (klassisch core/ports/adapters), when a-check laeuft, then identisches Verhalten
  wie 0.4.0 (kein domain->port im Eigen-Repo => Dogfooding unveraendert gruen).

Out-of-Scope: driving/driven-Port-Subtypen (welle-10b/b2b); feinere app-interne
Struktur.
```

### 3.2 Schema-Erweiterung (Config)

Die `role`-Whitelist im `layers`-Objekt-Zweig
([AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml),
Strict-Decode, `config.go:127-128`) wird um `app` ergänzt: gültig sind
`{domain, app, port, adapter}`. Unbekannte Rolle ⇒ weiterhin Exit 2 (fail-closed).
Ein **Positiv**-Test belegt die Akzeptanz von `role: app`.

### 3.3 Folge-ADR — Domain/Application-Trennung

```text
ADR-0011 — Domain/Application-Trennung: Rolle app + strenge domain
Status: Proposed (-> Accepted bei Merge).
Bezug: AC-FA-RULE-007 (neu), AC-FA-RULE-006 (Rollen-Mechanismus), AC-FA-RULE-001.
Schaerft: SPEC-RULE-001.

Kontext: ADR-0009 fuehrte Rollen {domain, port, adapter} ein, liess aber die
Domain/Application-Trennung offen (Re-Eval-Trigger). domain->port war nur
kanten-geregelt -- eine deklarierte Kante konnte die Trennung aushebeln.

Decision:
1. Neue Rolle app (kein Adapter/Tech, darf domain+port) -> Befund app-impurity.
   Inferenz application/app -> app; role: app explizit moeglich.
2. domain verschaerft: importiert nur domain (+ stdlib); Import auf app/port/adapter
   ODER Tech -> core-impurity, KATEGORISCH (Kante hebt nicht auf). "Domain kennt
   keine Ports" wird harte Invariante statt Kanten-Konvention.
3. app-Reinheit ebenfalls kategorisch; nur die Richtung bleibt kanten-geregelt.

Consequences: Vier-Schichten-Hexagon (domain <- app <- port <- adapter) voll
pruefbar; b-cad/d-migrate koennen ihre application-Schicht modellieren.
Rueckwaertskompat: ohne app-Layer und ohne domain->port-Import (a-check-Eigen-Repo)
unveraendert gruen. Migration: wer Ports bisher aus einer role: domain-Schicht
importierte, wird rot -- Port-Nutzung in eine role: app-Schicht heben. Ein 6.
Befund-Name (app-impurity) kommt hinzu; bestehende Namen stabil. Lastenheft
0.4.0 -> 0.5.0.

Re-Evaluierungs-Trigger: driving/driven-Port-Subtypen (b2b); LayerOf laengster-Praefix.
```

### 3.4 Versions-Bump

Lastenheft **0.4.0 → 0.5.0** (neue Anforderung *Domain/Application-Trennung*).

### 4.1 Engine + Config

1. `internal/adapter/driven/config/config.go` — die `role`-Whitelist
   (`config.go:127-128`: Prüfung + Fehlertext `(domain|port|adapter)`) um `app`
   erweitern → `(domain|app|port|adapter)`; fail-closed bleibt.
2. `internal/hexagon/core/model.go` (`inferRole`) — `application`→`app`,
   `app`→`app` ergänzen. `roleOf`/Vorrang von `role:` unverändert.
3. `internal/hexagon/core/rules.go` (`ruleFor`) — zwei Zweige, Erst-Treffer-Reihenfolge
   `domain` → `app` → `port` → `adapter`-lateral → `tech-leak` → `wrong-direction`:
   ```
   case src==domain ∧ ( tgt∈{app,port,adapter} ∨ isTech ):  core-impurity   // VERSCHAERFT
   case src==app    ∧ ( tgt==adapter ∨ isTech ):            app-impurity     // NEU
   ```
   Roleless Ziel (`tgt==""`) fällt durch auf die Kanten-Prüfung (kein kategorischer
   Befund). Befund-Namen sonst stabil.

### 4.2 „Fünf → sechs": Doku-/Kommentar-Sweep (F1)

`app-impurity` ist der **erste neue Befund-Name seit 0.1.0** — jede Aufzählung der
Regelmenge wird stale, und `make gates` (Links/Build/Tests) fängt das **nicht**.
**Entscheidung: alle konkreten Zählungen aktualisieren, keine narrativ ausklammern.**
Stellen:

- `internal/hexagon/core/rules.go:9` — Doc-Comment „five hexagon rules" → „six".
- `spec/spezifikation.md:108` „die fünf Hexagon-Regeln" → „sechs"; `:130` die
  normative Erst-Treffer-Reihenfolge ([SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung))
  um `app-impurity` ergänzen (nach `core-impurity`); `:69` Schema-Enum
  `role: domain|port|adapter` → `domain|app|port|adapter`.
- `spec/architecture.md:58` — die Kern-Zeile „die fünf Regeln" → „sechs".
- `docs/user/benutzerhandbuch.md:136` „Die fünf Regeln" → „sechs" + Befund-Tabelle
  (Z. 140-144) um eine `app-impurity`-Zeile (Beschreibung + Behebung); `:271` Namens-Liste.
- `README.md:16-17` „Fünf universelle Regeln, je eine Anforderung" → „Sechs" +
  Befund-Liste; die 1:1-Aussage entschärfen (F7: der neue Befund hängt an der neuen
  Anforderung, nicht an der 001…005-Kontiguität); `:42` „dieselben fünf Regeln" → „sechs".

### 4.3 Tests — kategorisch, nicht tautologisch (F3/F4)

Vorbild `TestRoleCrossLayerLateralCategorical` (`rules_test.go:218`): **Kante deklarieren**
+ `len(fs)==1 && fs[0].Rule=="…"` asserten (nicht bloß „ein Befund existiert" — ohne
Kante feuert ohnehin `wrong-direction`, anderer Name; das „auch mit Kante" wäre sonst
unbewiesen).

- `app` happy: Kanten `app→domain` **und** `app→port` deklariert, Import beider ⇒ **kein** Befund.
- `app→adapter` ⇒ `app-impurity`, kategorisch (Kante gesetzt, `len==1`).
- `app→Tech` ⇒ `app-impurity` (zweiter Arm des Negative-app-AC).
- `domain→port` ⇒ `core-impurity`, kategorisch (Kante gesetzt, `len==1`).
- `domain→app` ⇒ `core-impurity` (die Schärfung deckt `app`, nicht nur `port`).
- `domain→adapter` ⇒ `core-impurity` (Regressions-Pin, bestehendes Verhalten).
- Inferenz: `application`/`app`→`app`; explizite `role:` schlägt Inferenz.
- `config_test.go`: **Positiv**-Test `role: app` akzeptiert (heute nur `domainx`-Negativ, `:89`).
- Boundary: klassische Config (`core`/`ports`/`adapters`) unverändert grün.

### 4.4 Abschluss

Spezifikation **0.5.0**; `CHANGELOG.md` `[Unreleased]` (neuer Befund + Bump, F6);
`make gates` grün; Multi-Linsen-Review; Commit.

## 5. Definition of Done

- [x] Neue Anforderung *Domain/Application-Trennung* in `spec/lastenheft.md` (4 AC + Out-of-Scope, „Erweitert"/„Schärft"-Zeilen), Bump 0.5.0 + Historie.
- [x] Folge-ADR `Proposed → Accepted` (bei Merge) + ADR-Index.
- [x] [SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung) (Rollen-Tabelle: `app` + strenge `domain`; Erst-Treffer-Reihenfolge `:130` um `app-impurity`) + Schema-Enum `:69` (`app`) nachgezogen, Spezifikation 0.5.0.
- [x] **„Fünf → sechs"-Sweep vollständig** (F1): `rules.go:9`-Comment, `spezifikation.md:108`, `architecture.md`-Kern-Zeile, `benutzerhandbuch.md:136`/Tabelle `140-144`/`:271`, `README.md:16-17`/`:42` (inkl. F7-Entschärfung der 1:1-Aussage).
- [x] `config.go`: `role`-Whitelist um `app`; `model.go`: Inferenz `application`/`app`→`app`; `rules.go`: `domain`-Schärfung (`↛app/port/adapter`) + `app-impurity`-Zweig.
- [x] Engine: `app-impurity` (kategorisch) + `domain↛port` (kategorisch); `make arch-check` (Dogfooding) **ohne** Änderung der Eigen-`.a-check.yml` grün.
- [x] Tests (kategorisch, `len==1`+Regelname): `app` happy (Kanten gesetzt), `app→adapter`, `app→Tech`, `domain→port`, `domain→app`, `domain→adapter`-Pin, Inferenz `application`/`app`, `config_test` Positiv-`role: app`, Boundary klassisch.
- [x] `CHANGELOG.md` `[Unreleased]`: neuer Befund `app-impurity` + 0.5.0-Bump (F6).
- [x] Multi-Linsen-Review bestanden.

## 6. Offen / Risiken — Entscheidungen zur Abnahme

- **Entscheid A — AC-Heimat:** eine **neue Regel-Anforderung** (Empfehlung) vs. die
  geshippte [AC-FA-RULE-006](../../../../spec/lastenheft.md#ac-fa-rule-006--schicht-rollen-generische-regel-anwendung)
  erweitern. *Empfehlung neue Anforderung* — analog dazu, wie 006 die Regeln
  001/002/004 generalisierte, statt sie umzuschreiben; in 006s Rollen-Menge kommt
  ein Vorwärts-Verweis auf die neue Anforderung. (Kippbar auf „006 erweitern".)
- **Entscheid B — Umfang der `domain`-Schärfung:** „`domain` importiert **nur**
  `domain`" (↛`app`/`port`/`adapter`/Tech, Empfehlung) vs. minimal „nur `port`
  ergänzen" (Status quo + `domain↛port`). *Empfehlung voll* — die innerste Schicht
  ist die strengste; `app`/`port`/`adapter`-Importe der Domäne sind alle
  Reinheits-Brüche. (Kippbar auf minimal.)
- **Migration:** `domain→port` wird rot (gewollt) — Port-Nutzung in eine
  `role: app`-Schicht heben. `app-impurity` ist ein **6. Befund-Name** —
  Output-Konsumenten (CI-Parser) berücksichtigen.
- **Rückwärtskompat DoD-kritisch:** Eigen-`.a-check.yml` bleibt grün — kein
  `app`-Layer, kein `domain→port`-Import (Kern importiert keine Ports).
- **b2b** (`driving`/`driven`, `LayerOf` längster-Präfix) bleibt **Out-of-Scope**
  → eigener Slice.
- **Review-Integration (slice-011-Review, 2026-06-22):** F1 (HIGH, „fünf→sechs"-Sweep)
  → §4.2 + DoD; F2 (Config-Zeile `config.go:127-128` + Positiv-Test) → §3.2/§4.1/§4.3;
  F3/F4 (kategorische Tests im vollen AC-Umfang) → §4.3; F5 (AC-Wortlaut präzisiert,
  rollenlos bleibt kanten-geregelt) → §3.1; F6 (CHANGELOG) → DoD; F7 (README-1:1-Aussage)
  → §4.2. **F1-Entscheid:** alle konkreten Zählungen aktualisiert, keine ausgeklammert.

## 7. Closure-Notiz (nach `done/`)

**Belege:** `make gates` grün — lint/test/coverage; `arch-check` 0 (Dogfooding
unverändert, **ohne** Änderung der Eigen-`.a-check.yml`); `doc-check` 0/45.
[AC-FA-RULE-007](../../../../spec/lastenheft.md#ac-fa-rule-007--rolle-app-und-strenge-domain)
(Lastenheft 0.5.0) + [ADR-0011](../../adr/0011-domain-application-trennung-rolle-app.md)
`Accepted`; Spezifikation 0.5.0. Engine: `config.go` (`role`-Whitelist um `app`),
`rules.go` (`impurityFinding`-Extraktion: `app`-Zweig + `domain`-Schärfung, gocyclo-bedingt
herausgezogen) + `inferRole` (`application`/`app`). „Fünf→sechs"-Sweep über
Specs/Handbuch/README/CHANGELOG/Code-Comment. Multi-Linsen-Review (4 Linsen) + Delta
bestanden ([Review-Doc](../../../reviews/2026-06-22-slice-011-app-rolle.md)).
Abnahme-Entscheidungen: A = neue Anforderung, B = volle `domain`-Schärfung.

**Lerneintrag:**

- *Eine Schärfung trifft die definierende Anforderung mit:* ein neuer Befund
  (`app-impurity`) macht JEDE Regel-Zählung stale (sechs Stellen) UND die definierende
  Kern-Reinheits-Anforderung intern widersprüchlich (ihr Boundary erlaubte den jetzt
  verbotenen `domain→port`-Import). `make gates` fängt weder Zählung noch Wortlaut — nur
  der getracter Sweep + die Vertrag-Linse fanden beides (eine Zähl-Stelle übersahen sogar
  alle vier Linsen, erst der eigene Grep fand sie).
- *gocyclo erzwingt die ehrlichere Struktur:* die Domain-seitige Reinheits-Dispatch-Tabelle
  als eigener Helfer (`impurityFinding`) ist lint-konform UND die klarere Modellierung.
- *Differenzial schlägt Präsenz (erneut):* der `domain→adapter`-Pin musste die Kante setzen
  und `len==1` asserten, sonst beweist er die Kategorik nicht.

**Folge (welle-10b/b2b):** `driving`/`driven`-Port-Subtypen; `LayerOf` längster-Präfix
(Asymmetrie zu `targetLayer`).
