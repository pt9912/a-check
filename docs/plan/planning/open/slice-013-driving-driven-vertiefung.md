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

- **A — Auto-Inferenz der Richtung** aus dem Schicht-**Namen** (`driving`/`driven`-Hinweis):
  fehlt `direction`, wird sie inferiert; explizite `direction:` gewinnt. **[ADR-0012](../../adr/0012-driving-driven-richtung-orthogonale-dimension.md)
  hat diese Inferenz im Re-Eval-Trigger *bewusst verworfen* („explizit deklariert statt
  geraten").** Dieser Slice **kehrt das um** — die Begründungs-Last trägt der Folge-ADR
  (ADR-Supersession-Disziplin, **nicht** ein AGENTS-§-Pin). *Beste Munition:* **deren** eigene
  Inkonsistenz — sie stützt sich bei der **Rolle** auf Namens-Inferenz (`roleOf` `core`→`domain`,
  §Konsequenzen/Decoder), verwirft das **Richtungs**-Analogon aber knapp und ohne Abgleich.
  (Achtung: `roleOf` ist Exact-Match, kein Token — Grammatik §6-E.)
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
  Gate (slice-012 Entscheid-0) erfunden wurde. *(Reproduzierbar: `rg -l -g 'port/input/**'
  'port.out'` im x-wal-Baum (Stand 2026-06-23) → 0 Treffer; Mess-Befehl/Commit beim Abschluss
  zu hinterlegen.)*

## 3. Entwurf (zur Abnahme)

### 3.1 Auto-Inferenz (Teil A)

```text
AC-FA-RULE-008 (erweitert): Fehlt `direction` auf einer port-/adapter-Schicht, wird sie aus
einem `driving`/`driven`-Hinweis im Schicht-NAMEN abgeleitet (NICHT aus Glob/Pfad, §6-A;
Grammatik §6-E). Explizite `direction:` hat Vorrang; kein Hinweis => keine Richtung (inert).
Rueckwaertskompatibel: Schichten ohne Hinweis + ohne direction unveraendert.
```
Code: `dirOf` erhält einen Inferenz-Zweig (heute nur `layerByName(...).Direction`). **Achtung
(Determinismus, [AC-QA-01](../../../../spec/lastenheft.md#ac-qa-01--determinismus)):** anders als
`roleOf` (Exact-Name-Match, `case "adapters"`) ist „Hinweis im Namen" mehrdeutig — die
Inferenz-**Grammatik** (Exact-Segment vs. Substring; Trennzeichen/Case; Konflikt bei *beiden*
Hinweisen; Kollision literal `driven` vs. *enthält* `driven`) ist ein eigener Entscheid (§6-E)
und gehört **vor Code** in Folge-ADR + Spezifikation.

### 3.2 Port→Port (Teil B)

```text
Neue Regel (eigener Befund-Name, Entscheid-C): ein role:port, direction X, der einen
role:port, direction Y (Y!=X, beide gesetzt) importiert, ist ein Befund. Kategorisch
wie der adapter->port-Arm.
```
Code: **eigener `case`-Arm** (`srcRole=="port" && tgtRole=="port"`) in `ruleFor` mit **eigenem
Befund-Namen** (Entscheid-C) — **nicht** nur das `adapter`-Prädikat aufweiten (das würde sonst
`port-direction-mismatch`s Namen/Message wiederverwenden, rules.go:51).

### 3.3 Folge-ADR

Neuer **Folge-ADR** — Beziehung **nach dem [ADR-0014](../../adr/0014-resolution-roots.md)-Muster**
(taggleicher Präzedenzfall): im **Bezug**-Feld als *„Re-Evaluierung von [ADR-0012](../../adr/0012-driving-driven-richtung-orthogonale-dimension.md) (erweitert, kein
Supersede)"* ausschreiben, `Supersedes: —`. **Kein** neues Keyword. AGENTS §3.5 greift hier **nicht**:
sie markiert die Auto-Inferenz in ihrem *eigenen* Re-Evaluierungs-Trigger zur Wiedervorlage —
ein ADR, der einen dokumentierten Re-Eval-Trigger auflöst, ist keine „Korrektur durch Überschreiben";
sie bleibt immutable und im Kern gültig. Inhaltlich **entkräftet** der ADR die „explizit statt
geraten"-Begründung — z. B.: Inferenz nur als Default mit explizitem Vorrang, **Namens-** statt
Pfad-basiert, Determinismus gewahrt. Decision: (a) Richtungs-Namens-Inferenz (Grammatik §6-E),
(b) Port→Port kategorisch *[nur falls Entscheid 0 ⇒ B]*. Schärft die Spezifikation. Bump Lastenheft 0.7.0→0.8.0.

## 4. Umsetzungsplan

**Rückgrat = Teil A (Auto-Inferenz).** Teil-B-Schritte sind *konditional* hinter Entscheid 0.

1. **Spec (A) zuerst:** [AC-FA-RULE-008](../../../../spec/lastenheft.md#ac-fa-rule-008--driving-driven-port-richtung-regel-port-direction-mismatch)-Out-of-Scope-Zeile **geschärft** (Namens-Inferenz *rein*, Glob-/Pfad-Inferenz bleibt *out*) + **3 neue AC** (Happy/Boundary/Negative) + Bump 0.8.0 + Historie; Spezifikation; Folge-ADR (Re-Eval, §3.3) + Index.
2. `rules.go` `dirOf`: Inferenz-Zweig — `driving`/`driven`-Hinweis im Schicht-**Namen** (nicht Glob; §6-A), Grammatik §6-E, wenn `Direction==""`.
3. Tests (A): Inferenz happy / expliziter Vorrang / kein-Hinweis / **Beide-Hinweise-Konflikt** (§6-E).
4. **[nur falls Entscheid 0 ⇒ B]** `rules.go` **eigener `case`-Arm** (`srcRole=="port" && tgtRole=="port"`) + eigener Befund-Name (§6-C) + Tests (mismatch/kategorisch/boundary).
5. Doku-Sweep (Benutzerhandbuch §4 `direction`-Inferenz; ggf. README/architecture); `make gates`; 4-Linsen-Review (schriftlich); Verifikation; Closure (`done/`, Lerneintrag).

## 5. Definition of Done

**Rückgrat A:**
- [ ] [AC-FA-RULE-008](../../../../spec/lastenheft.md#ac-fa-rule-008--driving-driven-port-richtung-regel-port-direction-mismatch)-**Out-of-Scope-Zeile geschärft**: Namens-Inferenz zugelassen, **Glob-/Pfad-Inferenz bleibt out-of-scope** (sonst öffnet der Edit zu viel).
- [ ] **Drei neue AC** (Happy/Boundary/Negative) für die Inferenz (Anforderungs-Anlege-Prozess) + Bump 0.8.0 + Historie.
- [ ] Folge-ADR (Re-Evaluierung von [ADR-0012](../../adr/0012-driving-driven-richtung-orthogonale-dimension.md), `Supersedes: —`, §3.3) `Accepted` + Index; Spezifikation (Regel-/Schema-Strata) nachgezogen.
- [ ] `dirOf`-Inferenz in `rules.go` (Grammatik §6-E); Tests (happy/Vorrang/kein-Hinweis/Beide-Hinweise).
- [ ] `make arch-check` **0 am echten a-check-Config**. *(Beleg-Argument, nicht Teil des Hakens: unter Namens-Inferenz trägt der `adapters`-Name kein Token, §6-A.)*

**Konditional B (nur falls Entscheid 0 ⇒ B):**
- [ ] Port→Port-Guard + eigener Befund-Name (§6-C); Tests (mismatch/kategorisch/boundary).

**Abschluss:**
- [ ] Doku-Sweep; `make gates` grün; 4-Linsen-Review; Verifikation; Closure + Lerneintrag.

## 6. Offen / Entscheidungen zur Abnahme

- **Entscheid 0 — Scope (das Gate *symmetrisch* anlegen):** slice-012s Gate verlangt, dass
  **ein Konsument die Richtung real aktiviert**, bevor a-check-seitige Folge-Ergonomie gebaut wird.
  x-wal hat **keine `.a-check.yml`** → *kein* Teil hat diesen aktiven Konsumenten; damit ist
  **auch Teil A verfrüht** (Ergonomie für ein Feature, das noch keiner aktiviert hat). Zwei Achsen:
  - **Aktiver-Konsument-Gate:** A *und* B fallen durch (x-wal nicht aktiv).
  - **ADR-Präzedenz:** A trägt **zusätzlich** die Umkehr-Last gegen [ADR-0012](../../adr/0012-driving-driven-richtung-orthogonale-dimension.md) (§3.3), B nicht.

  *Empfehlung: **beide vertagen**, bis ein Konsument (x-wal o. a.) eine `.a-check.yml` mit
  `driving`/`driven`-Adapter- **und** -Port-Schichten trägt und die Redundanz **fühlt** (der Pilot).
  Dann Teil A — mit `Amends`-ADR; Teil B nur bei nachgewiesenem Port→Port-Crossing (slice-013b).*
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
- **Entscheid D — Folge-ADR-Beziehung (*gelöst*):** [ADR-0014](../../adr/0014-resolution-roots.md)
  liefert den taggleichen Präzedenzfall — Re-Evaluierung im **Bezug**-Feld ausschreiben,
  `Supersedes: —`; **kein** neues Keyword, **keine** §3.5-Kollision ([ADR-0012](../../adr/0012-driving-driven-richtung-orthogonale-dimension.md) löst seinen eigenen
  Re-Eval-Trigger auf, bleibt immutable). *Empfehlung: dieses Muster.*
- **Entscheid F — neue AC unter RULE-008 vs. eigene Anforderung:** die Inferenz-ACs als Erweiterung
  von [AC-FA-RULE-008](../../../../spec/lastenheft.md#ac-fa-rule-008--driving-driven-port-richtung-regel-port-direction-mismatch)
  (gleiche Richtungs-Anforderung, Versionshistorie der ID) **oder** eine neue eigene
  `AC-FA-RULE`-Anforderung (eigene ID nach dem Konventions-Schema). *Empfehlung: unter RULE-008
  erweitern — dieselbe Anforderung, wie zuvor das Extraktions-Backend um Java erweitert wurde.*
- **Entscheid E — Inferenz-Grammatik (Determinismus, [AC-QA-01](../../../../spec/lastenheft.md#ac-qa-01--determinismus)):**
  Exact-Segment-Match (Schicht heißt *genau* `driving`/`driven`) vs. Substring/Token (Name *enthält*
  `driving`). Zu definieren: Trennzeichen/Case; Konflikt bei **beiden** Hinweisen; Kollision literal
  `driven` vs. *enthält* `driven`. *Empfehlung: Exact-Segment (deterministisch, kollisionsarm);
  „enthält" nur mit klarer Konfliktregel.*

## 7. Closure-Notiz

_(beim Abschluss: `make gates`-Beleg, `arch-check` 0, Review/Verifikation, Lerneintrag.)_
