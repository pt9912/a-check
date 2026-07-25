# slice-038 — Layer-Tie-Break: Deklarationsreihenfolge statt alphabetisch (Konformitäts-Fix zu ADR-0013)

**Status:** **done** (2026-07-24) — Code + Tests + `make ci` grün, Binary-verifiziert, Implementierungs-Review bestanden (Befund R-1 behandelt), Entscheid C clean; ff auf `main` (Konformitäts-/Bugfix; **kein** neuer Vertrag, **kein** neuer ADR, **kein** Lastenheft-Bump).
**Auslöser:** Nebenbefund **F-3** aus der HexSlice-Gap-Analyse
([slice-037 §4.0](slice-037-hexslice-gap-analyse.md#40-port-rollen-klassifikation-verschachtelter-ports-verifiziert)).
**Bezug:** stellt die Implementierung auf den bereits akzeptierten Vertrag von
[ADR-0013](../../adr/0013-layerof-laengster-praefix.md) her; präzisierte Regel
[SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung),
betroffener Loader [AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml),
Determinismus [AC-QA-01](../../../../spec/lastenheft.md#ac-qa-01--determinismus). [Roadmap](../in-progress/roadmap.md).

> **Hinweis:** Analyse/Entwurf zur Abnahme. Der Fix ist ein **Konformitäts-Bugfix** (die Impl
> erfüllt eine akzeptierte, immutable ADR nicht) — daher **kein** neuer ADR und **kein**
> Lastenheft-Bump (§4 begründet). Entscheidungen §7 **vor** der Umsetzung.

---

## 1. Ziel

Die Schicht-Auflösung bei **Literal-Präfix-Gleichstand** an ihren dokumentierten Vertrag
angleichen: **„bei Gleichstand die zuerst deklarierte Schicht"**
([ADR-0013](../../adr/0013-layerof-laengster-praefix.md) §Entscheidung). Heute entscheidet
faktisch die **alphabetische** Layer-Reihenfolge — ein stiller Vertragsbruch, den kein Test
fängt.

## 2. Problem & Evidenz (verifiziert)

- **Vertrag:** [ADR-0013](../../adr/0013-layerof-laengster-praefix.md) (Accepted, immutable) und
  [SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung) legen fest:
  spezifischste Schicht = längster literaler Präfix, **bei Gleichstand die zuerst deklarierte**.
  Die ADR-Fitness-Function verlangt sogar explizit einen Test: *„make test: … Gleichstand →
  zuerst deklariert"*.
- **Ist:** `config.go` dekodiert `layers` als **`map[string]yaml.Node`** und baut `m.Layers` via
  `sortedKeys(…)` **alphabetisch** — die YAML-Deklarationsreihenfolge geht beim Decode verloren.
  `LayerOf`/`layerOfCand` iterieren diese alphabetische Slice mit striktem `>`; bei Gleichstand
  gewinnt daher die **alphabetisch erste** Schicht, nicht die zuerst deklarierte.
- **Beobachtet** ([slice-037 §4.0](slice-037-hexslice-gap-analyse.md#40-port-rollen-klassifikation-verschachtelter-ports-verifiziert),
  gegen `a-check:dev`): bei zwei Schichten mit identischem Literal-Präfix
  (`src/hexagon/application/**` vs. `src/hexagon/application/**/ports/**`) gewinnt stets
  `application` (`a` < `p`) — **auch nach YAML-Reihenfolge-Tausch**. Ein tiefen-verschachtelter
  Port wird als Rolle `app` fehlklassifiziert; `ports`-vor-`application`-Deklarieren hilft heute
  **nicht**.
- **Test-Blindstelle:** die von [ADR-0013](../../adr/0013-layerof-laengster-praefix.md) geforderte „zuerst deklariert"-Prüfung existiert nicht
  **end-to-end** (über `config.Load`). Ein Test, der `[]core.Layer` direkt in Wunschreihenfolge
  konstruiert, prüft nur `LayerOf` und **umgeht** genau den Decode, der den Bug trägt — deshalb
  blieb er verborgen.

## 3. Ursache (ein Ort)

`internal/adapter/driven/config/config.go`:
- Feld `Layers map[string]yaml.Node` — eine Go-Map, inhärent ungeordnet.
- `for _, name := range sortedKeys(yc.Layers)` — alphabetische Sortierung beim Aufbau von
  `m.Layers`.

Der Rest der Engine (`LayerOf`, `targetLayer`/`layerOfCand`) ist **korrekt**: er respektiert die
Reihenfolge der übergebenen `[]core.Layer`. Nur die Quelle dieser Reihenfolge ist falsch.

## 4. Fix (zur Abnahme) — order-erhaltender Decode

`layers` **in Dokumentreihenfolge** dekodieren statt alphabetisch:

- Feldtyp `Layers map[string]yaml.Node` → **ein** `yaml.Node` (die ganze Mapping-Node). `yaml.v3`
  erhält in `Node.Content` die **Dokumentreihenfolge** (Schlüssel/Wert-Paare abwechselnd).
- `m.Layers` durch paarweise Iteration über `Content` aufbauen (Key = `Content[i].Value`, Value =
  `Content[i+1]`), `sortedKeys` **für `layers` entfällt**; `decodeLayer` bleibt unverändert
  (validiert jede Schicht **selbst** strikt — `knownLayerKey`/`validRole`/`validDirection` —,
  hängt **nicht** am Top-Level-`KnownFields`, keine Strictness-Regression).
- **Guard 1 — Node-Typ (Review-Befund F-2):** **vor** der paarweisen Iteration
  `node.Kind == yaml.MappingNode` erzwingen, sonst Exit 2. Der bisherige `map[string]yaml.Node`-Decode
  wies eine `layers`-**Sequence** (`[a, b]`) oder einen **Scalar** per Typfehler fail-closed ab; die
  rohe Node tut das nicht mehr — ohne Guard würde die Content-Iteration Sequence-Elemente als
  Name/Wert **fehldeuten** statt sauber abzuweisen ([SPEC-CONF-001](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema) strict-decode).
  *(Der bestehende `TestLayerScalarFailsClosed` prüft nur den **per-Layer**-Scalar `core: "…"` via
  `decodeLayer`, **nicht** den Top-Level-`layers:`-Kind — diese Ebene ist ungedeckt.)*
- **Guard 2 — Duplikat-Schlüssel (Review-Befund F-1):** im Content-Loop doppelte Layer-**Namen**
  explizit abweisen → Exit 2. **Verifiziert:** `yaml.v3` erkennt Duplikate nur beim Decode in Map/Struct
  (`uniqueKeys`, `decode.go` `mapping`/`mappingStruct`); der Decode in eine **rohe Node** umgeht das
  (`decode.go` Node-Sonderpfad `out.Type() == nodeType`) und legt Duplikate **doppelt** in `Content`
  ab. Ohne Guard würde aus einem **heute** abgewiesenen Config-Fehler (doppelter Layer-Name → Exit 2,
  da die äußere `layers`-Map durch `mapping()` läuft) ein **stiller** Doppel-`core.Layer` mit unklarer
  Tie-Break-Semantik — eine fail-closed-Regression
  ([AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)).
- Pflichtblock-Prüfung (`len(...) == 0`) auf die Node-Form nachziehen — der **leere-Mapping**-Fall
  (`layers: {}`, `len(Content) == 0`) bleibt so durch die Pflichtblock-Prüfung gedeckt (Exit 2).
- **Bekannte Grenze (Implementierungs-Review R-1):** ein YAML-**Merge-Key** `<<` im `layers`-Block
  wird beim Node-Decode **nicht** mehr expandiert (der Map-Decode tat es) — er wird als Schichtname
  `<<` gelesen und **fail-closed** abgewiesen (Exit 2, verifiziert; nicht silent-wrong). Praktisch
  irrelevant: die zugehörige Anker-Definition scheitert ohnehin schon an der Top-Level-Strictness
  (unbekannter Schlüssel), und kein Konsument nutzt Merge-Keys in `layers`. Als dokumentierte Grenze
  ausgewiesen ([AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze))
  und per `TestLayersMergeKeyFailsClosed` gepinnt.
- **`languages`/`resolution`/`forbidden_constructs` bleiben Map + `sortedKeys`** — dort trägt die
  Reihenfolge **keine** Regel-Semantik (Dispatch je Sprache bzw. je Layer-**Name**); ihr Tie-Break
  ist ein eigener, hier **nicht** berührter Gegenstand (§7 Entscheid B).

**Warum kein ADR / kein Bump:**
[ADR-0013](../../adr/0013-layerof-laengster-praefix.md) hat „zuerst deklariert" **bereits
entschieden** und ist immutable; die Impl erfüllt sie nur nicht. Ein Fix, der eine akzeptierte
ADR **erfüllt**, ist ein Bugfix, **keine** neue Entscheidung (§7 Entscheid A hält den Grenzfall
offen). [ADR-0013](../../adr/0013-layerof-laengster-praefix.md) selbst hielt für diese Konsistenz-Arbeit **„kein Versions-Bump"** fest. SPEC und
Lastenheft sagen bereits das Richtige — kein Textänderungsbedarf (Decode-Reihenfolge ist ein
Impl-Detail, gehört in Code + Test, nicht in die sprachneutrale Spec).

**Determinismus** ([AC-QA-01](../../../../spec/lastenheft.md#ac-qa-01--determinismus)): Dokument-
reihenfolge ist bei identischer Datei **byte-stabil** — wie die alphabetische zuvor, nur
konform. Kein Zeit-/Zufalls-/Locale-Einfluss.

## 5. Umsetzungsplan

0. **Ist-Verhalten zuerst pinnen (Charakterisierungs-Tests, vor dem Umbau):** doppelter Layer-Name
   → heute Exit 2; Top-Level-`layers:` als Sequence/Scalar → heute Exit 2. So ist das fail-closed-
   Verhalten festgeschrieben, das der Fix **bewahren** muss (Review F-1/F-2).
1. `config.go`: `Layers`-Feld auf order-erhaltende Node umstellen; `m.Layers` in Dokument-
   reihenfolge bauen; **Guard 1** (`Kind == MappingNode`) + **Guard 2** (Duplikat-Namen) im
   Content-Loop; Pflichtblock-Check auf die Node-Form. `decodeLayer`/`sortedKeys`(für andere Blöcke)
   unangetastet.
2. **End-to-end-Test** (`config`-Paket, über `Adapter.Load` auf eine echte YAML): zwei Schichten
   mit **identischem Literal-Präfix**; Reihenfolge-Tausch im YAML **flippt** die Klassifikation
   (A-zuerst → A gewinnt; B-zuerst → B gewinnt). Genau der Decode-Pfad, den ein `LayerOf`-Unit-Test
   nicht abdeckt (§2).
3. Regressions-Beleg: `make arch-check` unverändert **0** (a-checks Eigen-Config hat keine
   Gleichstand-Overlaps → Ergebnis identisch, [ADR-0013](../../adr/0013-layerof-laengster-praefix.md)
   §Konsequenzen).
4. **Vor Merge (Entscheid C):** die realen Konsumenten-Configs (a-check, b-cad, belief-agent,
   d-check, d-migrate, grid-guide) kurz auf Gleichstand-Overlaps sichten — ein Betroffener bekäme
   einen Verhaltenswechsel; Ergebnis im Closure benennen.
5. `make gates` + `make ci`; adversarisches Multi-Linsen-Review (schriftlich); Verifikation;
   Closure (`done/`, Lerneintrag).

## 6. Definition of Done

- [x] `config.go` dekodiert `layers` **in Deklarationsreihenfolge** (`decodeLayers` über die rohe
      `yaml.Node`, Content-Iteration statt `map`+`sortedKeys`); übrige Blöcke unverändert;
      `decodeLayer`-Strictness erhalten.
- [x] **End-to-end-Test über `Adapter.Load`** (`TestLayerTieBreakUsesDeclarationOrder`, nicht
      `LayerOf` direkt): Namen bewusst nicht-alphabetisch (`zeta` vor `alpha`) — Gleichstand liefert
      die zuerst deklarierte, Reihenfolge-Tausch flippt. Erfüllt die bislang offene
      [ADR-0013](../../adr/0013-layerof-laengster-praefix.md)-Fitness-Function.
- [x] **Fail-closed bewahrt (Review F-1/F-2):** `TestDuplicateLayerNameFailsClosed`,
      `TestLayersSequenceFailsClosed`, `TestLayersScalarFailsClosed`, `TestLayersEmptyMappingFailsClosed`
      — je Exit 2 (Guard 1 Kind + Guard 2 Duplikat + Pflichtblock).
- [x] **Vor Merge:** Konsumenten-Configs (a-check, b-cad, d-check, d-migrate lokal) auf
      Gleichstand-Overlaps gesichtet (Entscheid C) — **keiner betroffen** (durchweg längere/disjunkte
      Literal-Präfixe → Longest-Prefix ohne Tie); belief-agent/grid-guide nicht lokal, ungesichtet.
- [x] `make arch-check` **0 Befunde** (Dogfooding unverändert).
- [x] `make ci` grün (gates + image-test: Fragment-Parität, `--print-graph`, Scan — alle Akzeptanzen).
- [x] **Implementierungs-Review** (adversarisch, 6 Linsen) bestanden; ein Befund **R-1** (Merge-Key
      `<<` in `layers` nicht mehr expandiert → fail-closed) behandelt: gepinnt + als Grenze
      dokumentiert (§4). Entscheide **A** (als Konformitäts-Bugfix, kein ADR), **B** (nur `layers`),
      **C** (Konsumenten clean), **D** (Overlap-Ergonomie an slice-037 A′ abgegrenzt) wie empfohlen.
- [x] Folge-Notiz in [slice-037](slice-037-hexslice-gap-analyse.md) §4.0/§4.3: nach dem Fix
      genügt `ports` **vor** `application` zu deklarieren (kein Business-Area-Aufzählen nötig).

## 7. Offen / Entscheidungen zur Abnahme

- **Entscheid A — Konformität vs. neue Entscheidung (ADR ja/nein):** der Fix erfüllt die
  akzeptierte [ADR-0013](../../adr/0013-layerof-laengster-praefix.md) → **Bugfix, kein ADR**
  (Empfehlung). Wertet der Auftraggeber den Wechsel alphabetisch→deklariert als *neue* beobachtbare
  Entscheidung, genügt ein knapper bestätigender ADR (`Schärft:` [ADR-0013](../../adr/0013-layerof-laengster-praefix.md), `Supersedes: —`) —
  §3.5-konform, da [ADR-0013](../../adr/0013-layerof-laengster-praefix.md) unberührt/erfüllt bleibt. *Empfehlung: als Konformitäts-Bugfix führen.*
- **Entscheid B — Scope der Reihenfolge:** **nur `layers`** (trägt die Tie-Break-Semantik;
  Empfehlung) — oder zusätzlich `languages` (Backend-Wahl bei Mehrfach-Match) order-erhaltend? Für
  `languages` ist heute kein Konsument-Bedarf belegt (analog dem Aktiver-Konsument-Gate von
  [slice-013](../open/slice-013-driving-driven-vertiefung.md#6-offen--entscheidungen-zur-abnahme)); der
  `forbidden_constructs`/`resolution`-Dispatch ist namens-/sprach-basiert, ordnungs-invariant.
  *Empfehlung: nur `layers`.*
- **Entscheid C — Regressions-Risiko:** die Änderung wirkt **nur** auf Configs mit
  Gleichstand-Overlap **und** nicht-alphabetischer Wunschreihenfolge. Bestandskonsumenten (a-check
  selbst, b-cad, belief-agent, d-check, d-migrate, grid-guide) haben — soweit bekannt — keine
  solchen Overlaps → kein Verhaltenswechsel. *Vor Merge: die realen Konsumenten-Configs kurz auf
  Gleichstand-Overlaps sichten; falls einer betroffen ist, im Closure benennen.*
- **Entscheid D — tiefere Overlap-Ergonomie (abgrenzen):** ob ein *spezifischerer-Glob-gewinnt-
  über-Overlap*-Mechanismus (tiefen-agnostische Ports ohne Reorder) gebaut wird, ist
  [slice-037 Option A′](slice-037-hexslice-gap-analyse.md#7-optionen--offene-entscheidungen),
  **nicht** dieser Slice. Hier nur der Tie-Break-Konformitäts-Fix; er macht den Tie **steuerbar**
  (Reorder wirkt), löst die Overlap-Ergonomie aber nicht restlos. *Empfehlung: getrennt halten.*

## 8. Closure (Verifikations-Notiz, Umsetzung 2026-07-24)

- **Code:** `internal/adapter/driven/config/config.go` — `Layers` von `map[string]yaml.Node` auf eine
  rohe `yaml.Node` umgestellt; neue `decodeLayers` iteriert `Content` in Dokumentreihenfolge mit
  **Guard 1** (`Kind == MappingNode`) + **Guard 2** (Duplikat-Namen) + Leer-Mapping-Check; `sortedKeys`
  nur noch für `resolution` genutzt.
- **Tests:** fünf neue Tests im `config`-Paket (Tie-Break end-to-end + vier fail-closed-Negativtests);
  `make test` grün (`ok …/config`).
- **Binary-Beleg (gegen frisch gebautes `a-check:dev`):** dieselbe überlappende HexSlice-Fixture, die
  **vor** dem Fix fälschlich 0 lieferte, meldet jetzt bei `ports`-vor-`application` korrekt
  `port-impurity` (Exit 1); bei `application`-vor-`ports` bleibt es 0 — der **Reihenfolge-Tausch
  flippt** das Ergebnis (Tie-Break deklarationsreihenfolge-getrieben, nicht mehr alphabetisch).
- **Gates:** `make gates` grün — `arch-check` 0 (Dogfooding unverändert), `test`/`lint`/`coverage`,
  `doc-check` 0, `gate-consistency`/`guard-selftest`/`record-gates` ok.
- **Entscheid C:** a-check/b-cad/d-check/d-migrate ohne Gleichstand-Overlap → kein Verhaltenswechsel.
- **Implementierungs-Review (6 Linsen):** bestanden; Befund **R-1** (Merge-Key `<<` in `layers` nicht
  mehr expandiert → fail-closed, Exit 2) gepinnt (`TestLayersMergeKeyFailsClosed`) + als Grenze
  dokumentiert. `make ci` grün (inkl. image-test).

**Lerneintrag:** Ein akzeptierter, immutabler ADR-Vertrag („zuerst deklariert") kann still verletzt
sein, wenn die **Fitness-Function auf der falschen Ebene** testet — ein `LayerOf`-Unit-Test mit
handsortierter `[]Layer` umging genau den `config`-Decode, der die Reihenfolge alphabetisch zerstörte.
Lehre: End-to-end **über den echten Loader** testen, wo der Vertrag entsteht. Und: ein
Datentyp-Wechsel (Map → rohe `yaml.Node`) verschiebt still fail-closed-Semantik (Duplikat-Keys,
Nicht-Mapping, Merge-Keys) — jede vom alten Typ gratis gelieferte Prüfung muss bewusst
rekonstruiert und gepinnt werden.
