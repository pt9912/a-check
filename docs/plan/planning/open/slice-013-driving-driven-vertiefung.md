# slice-013 — Driving/Driven-Vertiefung (Teil A vertagt, Teil B verworfen)

**Status:** open — **Entscheid 0 abgenommen 2026-07-25** (Maintainer-Wort, auf Basis der
Nachmessung §2a): **Teil B (Port→Port-Richtungsregel) ist verworfen**, **Teil A (Auto-Inferenz)
ist vertagt**. Der Slice bleibt als Entwurf für Teil A offen; die Trigger stehen in §0.
**Bezug:** Carry-forward aus [slice-012 §7](../done/slice-012-driving-driven-layerof.md);
verfeinert [AC-FA-RULE-008](../../../../spec/lastenheft.md#ac-fa-rule-008--richtungs-dimension-regel-port-direction-mismatch);
löst die in [ADR-0012](../../adr/0012-driving-driven-richtung-orthogonale-dimension.md)
als *out-of-scope* gestellten Richtungs-Inkremente. [Roadmap](../in-progress/roadmap.md).
**Evidenz-Kandidat:** externes Repo `x-wal` (lokal) — Kotlin-Hexagon mit `adapters/driving`+`adapters/driven` und `port/input`+`port/output`; **noch ohne `.a-check.yml`** (struktureller Kandidat, kein aktiver Konsument).

> **Hinweis:** Entwurf zur Abnahme. AC-/ADR-Texte in §3 (Code-Fences) sind unverbindlich
> bis Freigabe in `spec/`. DoD §5 offen; Entscheidungen §6 **vor** der Umsetzung.

---

## 0. Entscheid 0 — abgenommen (2026-07-25)

Auf Basis der Nachmessung §2a über **neun** lokale Repos entschieden:

### Teil B — Port→Port-Richtungsregel: **verworfen**

Nicht vertagt, sondern verworfen. Zwei unabhängige Messungen an zwei Repos zeigen, dass die
Kante, die die Regel bewachen soll, in dieser Architektur nicht vorkommt: x-wal (2026-06-23,
19 `port/input`-Dateien, 0 Treffer) und **b-cad** (2026-07-25, 8 `driving`- + 6 `driven`-Port-
Dateien, **0** Port→Port-Includes — die Port-Dateien zeigen nur auf `hexagon/model/**` und die
Standardbibliothek). b-cad ist dabei der **einzige** Konsument mit richtungs-getrennten
Port-Schichten überhaupt; kein zweiter kann die Regel derzeit auch nur auslösen. Eine Regel ohne
Anwendungsfall zu bauen ist genau die Spekulation, gegen die das Konsumenten-Gate erfunden wurde.

**Wiedervorlage-Trigger:** der **erste real gemessene Port→Port-Import über eine
Richtungsgrenze** in einem aktiven Konsumenten. Dann als eigener Slice (slice-013b) neu
aufsetzen — die Entwurfsteile §3.2/§6-C bleiben hier als Vorarbeit lesbar. Der Vertrag ändert
sich nicht: [ADR-0012](../../adr/0012-driving-driven-richtung-orthogonale-dimension.md) führt
Port→Port bereits als out-of-scope, diese Entscheidung bestätigt das nur.

### Teil A — Auto-Inferenz der Richtung: **vertagt**

Das ursprüngliche Vertagungs-Argument („kein aktiver Konsument") ist **verbraucht** — b-cad
aktiviert die Richtung real. Vertagt wird jetzt aus zwei **anderen**, schärferen Gründen:

1. **Kein messbarer Gewinn.** Beim einzigen aktiven Konsumenten spart die Inferenz unter der
   empfohlenen Exact-Segment-Grammatik **0 von 9** Deklarationen (unter Substring 2 von 9). Bei
   belief-agent träfe sie zwar 16 von 19 Schichten, hätte dort aber **keine Prüfwirkung** (keine
   Port-Schicht) und bräuchte eine fremde Vokabel (`inbound`/`outbound`).
2. **Die Inferenz ist nicht verhaltens-neutral** (§2a, Konsequenz 2b.3): eine Schicht mit
   Namens-Hinweis und ohne `direction` bekäme still eine Richtung — und damit kann eine heute
   **inerte kategorische** Regel später scharf werden, ohne dass der Konsument sie aktiviert hat.
   Das ist vor der Umsetzung zu klären (Opt-in-Schalter?), nicht danach.

**Geschärfter Lande-Trigger** (ersetzt den alten „ein Konsument aktiviert die Richtung"):

- ein Konsument trägt Schichten, deren **Namen** die entschiedene Grammatik wirklich treffen
  (Entscheid E) — nicht bloß Richtungs-Struktur, sondern Richtungs-**Namen**; **und**
- die Verhaltens-Neutralität ist geklärt (Opt-in oder Nachweis, dass keine bestehende Config
  dadurch neue Befunde bekommt).

Bis dahin bleiben die Entscheide A–F (§6) unabgenommen; die Umkehr-Last gegenüber
[ADR-0012](../../adr/0012-driving-driven-richtung-orthogonale-dimension.md) trägt weiterhin der
Folge-ADR (§3.3).

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

## 2a. Nachmessung 2026-07-25 — die Evidenzlage hat sich verschoben

Der Entwurf datiert auf 2026-06-23 und stützt sich allein auf `x-wal`. Seither ist **b-cad**
aktiver Konsument geworden. Nachgemessen an allen lokalen `.a-check.yml`-Konsumenten:

| Konsument | `direction`-Deklarationen | Port-Schichten getrennt? |
|---|---|---|
| **b-cad** | **9** (`ports_driving`/`ports_driven` **und** 7 Adapter-Schichten) | **ja** |
| d-check | 0 | nein (eine `ports`-Schicht) |
| d-migrate | 0 | nein (fünf Port-Globs in **einer** Schicht; `adapters/driving/**`+`adapters/driven/**` ebenfalls in **einer**) |
| m-trace | 0 | nein |
| **belief-agent** | 0 — aber **16 von 19** Schichten heißen `inbound_*`/`outbound_*` | **keine Port-Schicht deklariert** |
| HexSlice-Go-Beispiel | 0 (Namen `domain`/`ports`/`app`/`adapters`) | nein |
| x-wal | — (weiterhin **keine** `.a-check.yml`) | n/a |
| u-boot, ai-harness-init | — (keine `.a-check.yml`) | n/a |

**Konsequenz 1 — das Aktiver-Konsument-Gate aus Entscheid 0 ist erfüllt**, aber von **b-cad**,
nicht von x-wal: b-cad aktiviert die Richtung real auf Adapter- **und** Port-Seite. Der
slice-012-Carry-forward („mindestens ein Konsument soll sie real aktivieren") ist damit
eingelöst — der Vertagungsgrund von 2026-06-23 trägt nicht mehr.

**Konsequenz 2 — Teil A gewinnt trotzdem fast nichts.** Unter der empfohlenen
**Exact-Segment**-Grammatik (§6-E) müsste eine Schicht *genau* `driving`/`driven` heißen; b-cads
Schichten heißen `ports_driving`, `ports_driven`, `ui_command`, `ui_view`, `io`, `geometry`,
`persistence`, `pluginhost`, `plugins` ⇒ **0 von 9** Deklarationen eingespart. Unter einer
Substring-/Token-Grammatik wären es **2 von 9** (nur die beiden Port-Schichten); die übrigen
sieben tragen kein Token und blieben explizit. Die Redundanz, die Teil A beseitigen soll,
existiert beim einzigen aktiven Konsumenten also kaum — gegen die Umkehr-Last gegenüber
[ADR-0012](../../adr/0012-driving-driven-richtung-orthogonale-dimension.md) ist das wenig.

**Konsequenz 3 — Teil B hat weiterhin null Anwendungsfälle, jetzt am aktiven Konsumenten
gemessen.** In b-cads Port-Baum (8 `driving`-, 6 `driven`-Dateien) gibt es **keinen einzigen**
Port→Port-Include — weder richtungs-querend noch überhaupt; die Includes der Port-Dateien zeigen
ausschließlich auf `hexagon/model/**` und die Standardbibliothek. Die Port-Symbole werden von
`services`, `adapters/*` und `plugin_api` genutzt, nicht von Ports untereinander. Damit ist die
x-wal-Messung von 2026-06-23 (0 Treffer) an einem **zweiten, aktiven** Konsumenten bestätigt.

> **Mess-Notiz zur Ehrlichkeit:** eine erste, grobe Zählung schien 53 Port→Port-Includes zu
> zeigen — der Zähler traf jedoch das `src/hexagon/ports/…`-Präfix in der **Dateipfad**-Spalte
> der `grep -n`-Ausgabe, nicht das Include-Ziel. Nach Korrektur: **0**. Die erste Zahl war ein
> Messfehler, kein Befund.

**Konsequenz 2b — belief-agent sieht aus wie das Kronzeugen-Repo für Teil A, ist es aber nicht.**
16 von 19 Schichten heißen `inbound_cli` bzw. `outbound_*` — eine Namens-Inferenz mit
Substring-Grammatik träfe sie alle. Drei Dinge stehen dagegen:
1. **Falsche Vokabel (Entscheid B):** die Namen tragen `inbound`/`outbound`, nicht
   `driving`/`driven`. Die empfohlene Vokabel greift dort **nicht**; sie zu erweitern heißt,
   fremde Begriffe zu interpretieren — genau das, was Entscheid B ausschließen wollte.
2. **Keine Wirkung:** belief-agent deklariert **keine Port-Schicht**. `direction` steuert allein
   `port-direction-mismatch` ([AC-FA-RULE-008](../../../../spec/lastenheft.md#ac-fa-rule-008--richtungs-dimension-regel-port-direction-mismatch)),
   und die Regel verlangt **beide** Seiten gesetzt. Die Inferenz erzeugte dort 16 Richtungen
   **ohne jede Prüfwirkung** — Metadaten, die nichts gaten.
3. **Sie wäre nicht rückwärtskompatibel-neutral** (neuer Befund, im Entwurf nicht bedacht):
   §3.1 sagt „rückwärtskompatibel: Schichten **ohne Hinweis** + ohne `direction` unverändert" —
   korrekt, aber Schichten **mit** Hinweis und ohne `direction` ändern ihr Verhalten. Sobald ein
   solcher Konsument später eine Port-Schicht mit Richtung ergänzt, würde eine heute **inerte
   kategorische** Regel still scharf — ein Gate, das der Konsument nie aktiviert hat. Eine
   Verschärfung per Inferenz braucht mindestens einen expliziten Opt-in-Schalter, sonst ist sie
   dieselbe Klasse stiller Setzung, gegen die dieses Repo sonst fail-closed vorgeht.

**Konsequenz 3b — d-migrate ist der Gegenbeweis für die Glob-Variante.** Dort trägt **eine**
Schicht beide Tokens: `adapters: {globs: ["adapters/driven/**", "adapters/driving/**"]}`. Eine
glob-basierte Inferenz stünde genau vor dem Konflikt-Fall aus Entscheid E — und zwar nicht
hypothetisch, sondern in einer real existierenden Konsumenten-Config.

**Konsequenz 4 — die Dogfooding-Gefahr aus Entscheid A ist real:** a-checks Eigen-Config trägt
`adapters: ["internal/adapter/driven/**"]` — eine **Glob**-basierte Inferenz würde dort `driven`
ableiten. Namens-basiert bleibt der Layer-Name `adapters` token-los. Die Empfehlung „Namens-basiert"
steht damit bestätigt, falls Teil A je gezogen wird.

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

### 3.2 Port→Port (Teil B) — **verworfen (§0)**, bleibt als Vorarbeit lesbar

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
(b) Port→Port kategorisch *[nur falls Entscheid 0 ⇒ B]*. Schärft die Spezifikation. Bump Lastenheft beim Landen gegen den dann-aktuellen Stand —
die im Entwurf notierte **0.8.0 ist seither vergeben** (Stand dieser Currency-Notiz 0.20.0
→ nächster Minor). **Currency-Notiz (nach Entwurf):** Der Folge-ADR erhält die nächste
freie Nummer **nach [ADR-0024](../../adr/0024-print-graph-mermaid.md)**; da er **nach
[MR-005](../../../../harness/conventions.md#mr-005--referenzmatrix-intra-spec-richtung--adrslice-disziplin-d-check-angleichung)**
entsteht, ist er slice-token-frei zu argumentieren *oder* trägt den Provenance-Marker (die
`adr → slice`-Disziplin gilt für ADRs ab 0021; im Repo existieren inzwischen echte
Supersede-Präzedenzen).

## 4. Umsetzungsplan

**Rückgrat = Teil A (Auto-Inferenz).** Teil-B-Schritte sind *konditional* hinter Entscheid 0.

1. **Spec (A) zuerst:** [AC-FA-RULE-008](../../../../spec/lastenheft.md#ac-fa-rule-008--richtungs-dimension-regel-port-direction-mismatch)-Out-of-Scope-Zeile **geschärft** (Namens-Inferenz *rein*, Glob-/Pfad-Inferenz bleibt *out*) + **3 neue AC** (Happy/Boundary/Negative) + Bump (nächster Minor, §3.3) + Historie; Spezifikation; Folge-ADR (Re-Eval, §3.3) + Index.
2. `rules.go` `dirOf`: Inferenz-Zweig — `driving`/`driven`-Hinweis im Schicht-**Namen** (nicht Glob; §6-A), Grammatik §6-E, wenn `Direction==""`.
3. Tests (A): Inferenz happy / expliziter Vorrang / kein-Hinweis / **Beide-Hinweise-Konflikt** (§6-E).
4. ~~**[nur falls Entscheid 0 ⇒ B]** `rules.go` **eigener `case`-Arm** (`srcRole=="port" && tgtRole=="port"`) + eigener Befund-Name (§6-C) + Tests (mismatch/kategorisch/boundary).~~ — **entfällt, Teil B ist verworfen (§0).**
5. Doku-Sweep (Benutzerhandbuch §4 `direction`-Inferenz; ggf. README/architecture); `make gates`; 4-Linsen-Review (schriftlich); Verifikation; Closure (`done/`, Lerneintrag).

## 5. Definition of Done

**Rückgrat A:**
- [ ] [AC-FA-RULE-008](../../../../spec/lastenheft.md#ac-fa-rule-008--richtungs-dimension-regel-port-direction-mismatch)-**Out-of-Scope-Zeile geschärft**: Namens-Inferenz zugelassen, **Glob-/Pfad-Inferenz bleibt out-of-scope** (sonst öffnet der Edit zu viel).
- [ ] **Drei neue AC** (Happy/Boundary/Negative) für die Inferenz (Anforderungs-Anlege-Prozess) + Bump (nächster Minor, §3.3) + Historie.
- [ ] Folge-ADR (Re-Evaluierung von [ADR-0012](../../adr/0012-driving-driven-richtung-orthogonale-dimension.md), `Supersedes: —`, §3.3) `Accepted` + Index; Spezifikation (Regel-/Schema-Strata) nachgezogen.
- [ ] `dirOf`-Inferenz in `rules.go` (Grammatik §6-E); Tests (happy/Vorrang/kein-Hinweis/Beide-Hinweise).
- [ ] `make arch-check` **0 am echten a-check-Config**. *(Beleg-Argument, nicht Teil des Hakens: unter Namens-Inferenz trägt der `adapters`-Name kein Token, §6-A.)*

**Konditional B:** ~~Port→Port-Guard + eigener Befund-Name (§6-C); Tests.~~ — **entfällt** (§0:
Teil B verworfen; Wiedervorlage nur bei erstem real gemessenem Port→Port-Crossing).

**Abschluss:**
- [ ] Doku-Sweep; `make gates` grün; 4-Linsen-Review; Verifikation; Closure + Lerneintrag.

## 6. Offen / Entscheidungen zur Abnahme

- **Entscheid 0 — Scope (das Gate *symmetrisch* anlegen):** slice-012s Gate verlangt, dass
  **ein Konsument die Richtung real aktiviert**, bevor a-check-seitige Folge-Ergonomie gebaut wird.
  x-wal hat **keine `.a-check.yml`** → *kein* Teil hat diesen aktiven Konsumenten; damit ist
  **auch Teil A verfrüht** (Ergonomie für ein Feature, das noch keiner aktiviert hat). Zwei Achsen:
  - **Aktiver-Konsument-Gate:** A *und* B fallen durch (x-wal nicht aktiv).
  - **ADR-Präzedenz:** A trägt **zusätzlich** die Umkehr-Last gegen [ADR-0012](../../adr/0012-driving-driven-richtung-orthogonale-dimension.md) (§3.3), B nicht.

  *Empfehlung von 2026-06-23: **beide vertagen**, bis ein Konsument eine `.a-check.yml` mit
  `driving`/`driven`-Adapter- **und** -Port-Schichten trägt und die Redundanz **fühlt** (der Pilot).
  Dann Teil A — mit `Amends`-ADR; Teil B nur bei nachgewiesenem Port→Port-Crossing (slice-013b).*

  **Fortschreibung 2026-07-25 (§2a):** Der Pilot ist da — **b-cad** trägt die Richtung real auf
  beiden Seiten. Das damalige Argument („kein aktiver Konsument") ist damit **verbraucht**; die
  Nachmessung liefert aber zwei neue, schärfere:
  - **Teil A:** der aktive Konsument spart unter der empfohlenen Grammatik **0 von 9**
    Deklarationen (unter Substring 2 von 9). Die Redundanz, die A beseitigen soll, ist bei ihm
    fast nicht vorhanden — zu wenig, um die Umkehr-Last gegen
    [ADR-0012](../../adr/0012-driving-driven-richtung-orthogonale-dimension.md) zu tragen.
    *Neue Empfehlung: vertagen — und zwar bis ein Konsument Schichten trägt, die die Grammatik
    wirklich treffen (Namen `driving`/`driven`), statt bis „irgendein Konsument die Richtung
    aktiviert".*
  - **Teil B:** **0 Port→Port-Kanten** im aktiven Konsumenten (zweite Bestätigung nach x-wal).
    *Neue Empfehlung: nicht vertagen, sondern **verwerfen** — zwei unabhängige Messungen an zwei
    Repos zeigen, dass die Kante, die die Regel bewachen soll, in dieser Architektur gar nicht
    vorkommt. Ein Wiedervorlage-Trigger (erster realer Port→Port-Import) ist billiger als ein
    offener Entwurf, der Aufmerksamkeit bindet.*
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
  von [AC-FA-RULE-008](../../../../spec/lastenheft.md#ac-fa-rule-008--richtungs-dimension-regel-port-direction-mismatch)
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
