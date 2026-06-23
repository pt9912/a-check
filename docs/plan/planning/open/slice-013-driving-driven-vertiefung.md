# slice-013 — Driving/Driven-Vertiefung (Entwurf zur Abnahme)

**Status:** open — **Entwurf zur Abnahme** (`x-wal` ist *struktureller Kandidat*, noch kein aktiver Konsument).
**Bezug:** Carry-forward aus [slice-012 §7](../done/slice-012-driving-driven-layerof.md);
verfeinert [AC-FA-RULE-008](../../../../spec/lastenheft.md#ac-fa-rule-008--driving-driven-port-richtung-regel-port-direction-mismatch);
löst die in [ADR-0012](../../adr/0012-driving-driven-richtung-orthogonale-dimension.md)
als *out-of-scope* gestellten Richtungs-Inkremente. [Roadmap](../in-progress/roadmap.md).
**Evidenz-Kandidat:** externes Repo `x-wal` (lokal) — Kotlin-Hexagon mit `adapters/driving`+`adapters/driven` und `port/input`+`port/output`; **noch ohne `.a-check.yml`** (struktureller Kandidat, kein aktiver Konsument).

> **Hinweis:** Entwurf zur Abnahme. AC-/ADR-Texte in §3 (Code-Fences) sind unverbindlich
> bis Freigabe in `spec/`. DoD §5 offen; Entscheidungen §6 **vor** der Umsetzung.

---

## 1. Ziel

Zwei a-check-seitige Richtungs-Inkremente — **mit unterschiedlichem [ADR-0012](../../adr/0012-driving-driven-richtung-orthogonale-dimension.md)-Status (wichtig):**

- **A — Auto-Inferenz der Richtung** aus dem Schicht-**Namen** (`driving`/`driven`-Token,
  analog `roleOf` `core`→`domain`): fehlt `direction`, wird sie inferiert; explizite
  `direction:` gewinnt. **Die ADR hat diese Inferenz *bewusst verworfen* („explizit
  deklariert statt geraten") — und die Rollen-Inferenz-Analogie dabei erwogen.** Dieser Slice
  **kehrt die Entscheidung um**; der Folge-ADR muss sie aktiv **entkräften** (AGENTS §3.5),
  nicht nur neu begründen.
- **B — Port→Port-Richtungsregel:** Richtungs-Abgleich auch zwischen Ports (heute nur
  `adapter→port`). Dort als **out-of-scope, *späteres Inkrement*** *vertagt* (keine
  Design-Ablehnung).

## 2. Problem & Evidenz (x-wal)

- **A (Auto-Inferenz) — Bedarf nur *teilweise*, ADR-Präzedenz dagegen:** x-wals **Adapter**
  heißen literal `driving`/`driven` → dort spart Inferenz die Deklaration. x-wals **Ports**
  heißen `input`/`output` (kein `driving`/`driven`-Token) → unter Entscheid-B bleibt die
  Inferenz für sie wirkungslos, die Redundanz wird nur **halb** gelöst. Zudem hat x-wal
  **keine `.a-check.yml`** — es *spürt* die Redundanz noch nicht. Gegengewicht: die ADR
  verwirft die Inferenz bewusst (§1-A).
- **B (Port→Port) — Bedarf NICHT belegt:** in x-wal importiert **keines** der 19
  `port/input`-Files ein `port/output`-Symbol (Port-Symbole werden nur aus `application`
  genutzt). Die Regel hätte **null** aktuelle Anwendungsfälle — Spekulation, gegen die das
  Gate (slice-012 Entscheid-0) erfunden wurde.

## 3. Entwurf (zur Abnahme)

### 3.1 Auto-Inferenz (Teil A)

```text
AC-FA-RULE-008 (erweitert): Fehlt `direction` auf einer port-/adapter-Schicht, wird sie
aus dem Namen/Glob inferiert -- ein `driving`-Token => driving, `driven` => driven (analog
roleOf-Namens-Inferenz). Explizite `direction:` hat Vorrang; kein Token => keine Richtung
(Regel inert). Rueckwaertskompatibel: Schichten ohne Token + ohne direction unveraendert.
```
Code: `dirOf` erhält einen Inferenz-Zweig (heute nur `layerByName(...).Direction`).

### 3.2 Port→Port (Teil B)

```text
port-direction-mismatch (erweitert): der Rollen-Guard srcRole==adapter wird um
srcRole==port ergaenzt -- ein role:port, direction X, der einen role:port, direction Y
(Y!=X, beide gesetzt) importiert, ist ein Befund. Kategorisch wie der adapter->port-Arm.
```
Code: Guard in `ruleFor` von `srcRole=="adapter"` auf `srcRole∈{adapter,port}` erweitern.

### 3.3 Folge-ADR

Neuer **Folge-ADR** (neue Nummer, da [ADR-0012](../../adr/0012-driving-driven-richtung-orthogonale-dimension.md)
`Accepted`/immutable). Weil er **deren** bewusste Ablehnung der Inferenz umkehrt, muss er
deren Begründung *„explizit deklariert statt geraten"* **aktiv entkräften** (AGENTS §3.5) —
z. B.: Inferenz nur als Default mit explizitem Vorrang, **Namens-** statt Pfad-basiert,
Determinismus gewahrt. Decision: (a) Richtungs-Namens-Inferenz, (b) Port→Port kategorisch.
Schärft die Spezifikation (Regel-/Schema-Strata). Versions-Bump Lastenheft 0.7.0→0.8.0.

## 4. Umsetzungsplan

1. `rules.go` `dirOf`: Inferenz-Zweig — `driving`/`driven`-Token im Schicht-**Namen** (nicht im Glob; §6-A) wenn `Direction==""`.
2. `rules.go` `ruleFor`/`directionMismatch`: Port→Port-Arm (Guard `srcRole∈{adapter,port}`).
3. Tests: Inferenz happy/Vorrang/kein-Token; Port→Port mismatch/kategorisch/boundary.
4. Spec: [AC-FA-RULE-008](../../../../spec/lastenheft.md#ac-fa-rule-008--driving-driven-port-richtung-regel-port-direction-mismatch)-Erweiterung (0.8.0) + Spezifikation (Regel-/Schema-Strata); neuer Folge-ADR + Index.
5. Doku-Sweep (Benutzerhandbuch §4 `direction`-Inferenz; ggf. README/architecture).
6. `make gates`; 4-Linsen-Review (schriftlich); Verifikation; Closure (`done/`, Lerneintrag).

## 5. Definition of Done

- [ ] [AC-FA-RULE-008](../../../../spec/lastenheft.md#ac-fa-rule-008--driving-driven-port-richtung-regel-port-direction-mismatch)-Erweiterung (Inferenz + Port→Port) + Bump 0.8.0 + Historie; Folge-ADR `Accepted` + Index; Spezifikation nachgezogen.
- [ ] `dirOf`-Inferenz + Port→Port-Guard in `rules.go`; `make arch-check` **0 am echten a-check-Config belegt** — unter **Namens**-Inferenz trägt der `adapters`-Schicht-**Name** kein Token (Inferenz feuert nicht); **Glob**-Inferenz enthielte dagegen `internal/adapter/driven/**`s `driven`-Token (§6-A).
- [ ] Tests: Inferenz (happy/Vorrang/kein-Token), Port→Port (mismatch/kategorisch/boundary).
- [ ] Doku-Sweep; `make gates` grün; 4-Linsen-Review; Verifikation; Closure + Lerneintrag.

## 6. Offen / Entscheidungen zur Abnahme

- **Entscheid 0 — Scope/Schnitt (WICHTIG, *doppelt asymmetrisch*):** Auf der **Evidenz**-Achse
  ist Teil A motiviert (x-wal-Adapter), Teil B nicht (kein Port→Port-Crossing). Auf der
  **ADR-Präzedenz**-Achse ist es **umgekehrt**: Teil A kehrt eine *bewusste* [ADR-0012](../../adr/0012-driving-driven-richtung-orthogonale-dimension.md)-Ablehnung
  um (braucht den entkräftenden Folge-ADR, §3.3), Teil B war nur *vertagt*. *Empfehlung:
  Teil A nur ziehen, wenn der Folge-ADR „explizit statt geraten" sauber entkräftet; Teil B als
  eigener gated Slice zurückstellen (slice-013b).*
- **Entscheid A — Inferenz-Basis Name vs. Glob (*Dogfooding-kritisch*):** **Namens**-basiert
  (a-checks `adapters`-Name trägt kein Token → `arch-check` bleibt 0) vs. **Glob**-basiert
  (a-checks `internal/adapter/driven/**` trägt `driven` → würde `driven` inferieren →
  `arch-check`-0 hinge nur an der token-losen Port-Seite). *Empfehlung: **Namens**-basiert.*
- **Entscheid B — Inferenz-Token-Vokabel:** nur `driving`/`driven` (x-wal-Adapter; x-wal-Ports
  `input`/`output` blieben explizit) **oder** zusätzlich `input`/`output`/`inbound`/`outbound`.
  *Empfehlung: `driving`/`driven` (deckungsgleich mit der Vokabel; `input`/`output` wäre interpretierend).*
- **Entscheid C — Port→Port: kategorisch + Befund-Name:** kategorisch wie der `adapter→port`-Arm
  (Empfehlung); **und** *gleicher* Befund-Name `port-direction-mismatch` vs. **eigener** Name —
  ein `driving`-Port→`driven`-Port ist ein *anderer* Verstoß als ein Adapter am falschen Port
  (CI-Parser/Output-Konsumenten). *Empfehlung: eigener Befund-Name, falls Teil B gezogen wird.*
- **Entscheid D — Folge-ADR:** neuer ADR nötig ([ADR-0012](../../adr/0012-driving-driven-richtung-orthogonale-dimension.md) immutable);
  der Inferenz-Teil **entkräftet** deren Ablehnung (§3.3), der Port→Port-Teil begründet neu. *Empfehlung: ja.*

## 7. Closure-Notiz

_(beim Abschluss: `make gates`-Beleg, `arch-check` 0, Review/Verifikation, Lerneintrag.)_
