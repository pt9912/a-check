# slice-027 — Datei-mengen-bewusste KMP-Multi-Modul-Resolution (Bug + Feature)

**Status:** **done** (2026-07-05) — Spec/Lastenheft **0.17.0** + [ADR-0022](../../adr/0022-datei-mengen-bewusste-mehr-wurzel-aufloesung.md)
`Accepted` (Supersedes [ADR-0020](../../adr/0020-mehr-wurzel-phantom-guard.md)) + Code + AC1–AC6 grün +
Benutzerhandbuch 1.26; Maintainer-Review + Umsetzungs-Review (A1-Fix) eingearbeitet; Closure §9.
**Typ:** Bug (stilles Falsch-Negativ) + Feature
(Multi-Modul-Auflösung), konsumenten-getrieben (belief-agent).
**Bezug:** schärft [AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml)
(resolution + fail-closed) und [AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)
(ehrliche Heuristik-Grenze); **Stufe 2** zu [ADR-0020](../../adr/0020-mehr-wurzel-phantom-guard.md)
(dort §Re-Evaluierungs-Trigger namentlich als slice-027 gegatet), re-evaluiert das
Auflösungs-Modell [ADR-0016](../../adr/0016-resolution-sprach-parametrisch.md) /
[ADR-0014](../../adr/0014-resolution-roots.md). Ein Folge-ADR (Supersedes
[ADR-0020](../../adr/0020-mehr-wurzel-phantom-guard.md), Provenance-Marker nach
[MR-005](../../../../harness/conventions.md#mr-005--referenzmatrix-intra-spec-richtung--adrslice-disziplin-d-check-angleichung))
entsteht mit der Umsetzung. [Roadmap](../in-progress/roadmap.md).
**Evidenz:** belief-agent-KMP-Bericht gegen **v0.10.0** — jede `resolution`-Variante
endet in Reject **oder** still falsch-grün.

## 1. Auslöser

Für **KMP-/Gradle-Multi-Modul mit geteiltem `package_base` und disjunkten
Paket-Sub-Namespaces pro Modul** existiert **keine** korrekte `resolution`-Config.
Realer Fall (`pt9912/belief-agent`):

```
hexagon/domain/src/commonMain/kotlin/dev/beliefagent/domain/**        # dev.beliefagent.domain.*  -> "domain"
hexagon/application/src/commonMain/kotlin/dev/beliefagent/application/** # dev.beliefagent.application.* -> "application"
```

Erwartung: `import dev.beliefagent.application.*` **aus** einer domain-Datei muss die
verbotene `domain → application`-Kante melden. Reproduktion — jede Variante scheitert:

| # | `resolution.kotlin` | Ergebnis |
|---|---|---|
| 1 | 2 Roots `.../dev/beliefagent` je Modul, `package_base: dev.beliefagent` | **Exit 2 (Reject)** |
| 2 | 2 Roots tiefer (`.../domain`, `.../application`) | **Exit 2 (Reject)** |
| 3 | 1 Glob-Root `hexagon/*/src/commonMain/kotlin/dev/beliefagent` | akzeptiert, **falsch-grün** |
| 4 | 2 Roots + tiefe paket-spezifische `layers`-Globs | akzeptiert, **falsch-grün** |
| 5 | `mode: path` | akzeptiert, **falsch-grün** (ohne `package_base` keine Intern-Erkennung) |

Beleg der echten Durchsetzung sonst: Tech-Leak `org.koin` im Kern → gefunden;
`domain → application` → **0 Befunde** (die Lücke). Der Reject (1/2) ist der
fail-closed-Guard aus [slice-026](../done/slice-026-kmp-mehr-root-phantom.md) /
[ADR-0020](../../adr/0020-mehr-wurzel-phantom-guard.md), der die *legitime* disjunkte
Config pauschal ablehnt.

## 2. Kern-Erkenntnis (Fehlerquelle im Code)

Die Schicht-Zuordnung eines aufgelösten Imports hängt effektiv am **Root-Präfix**
statt am **Pfad der real existierenden Zieldatei**:

- `resolveImport` (`internal/hexagon/core/rules.go`), `fixed-root`: erzeugt **je Root
  einen Kandidaten** — **datei-mengen-blind**, ohne zu prüfen, ob die Zieldatei real
  in diesem Root liegt.
- `targetLayer`: matcht **alle** Kandidaten gegen alle `layers`-Globs und nimmt den
  **längsten literalen Glob-Präfix über alle Kandidaten**. Bei Layer-Globs flacher als
  die paket-diskriminierende Tiefe entscheidet damit der Root-Präfix — das **Phantom**
  (nicht existierende Datei) gewinnt → falsche Schicht → still grün.
- Der statische Guard `PhantomRootConflictIn` (Config-Ladezeit) lehnt jede
  `fixed-root`-Config mit ≥ 2 Roots ab, deren Roots verschiedene Schichten *erzwingen*
  — genau der disjunkte Multi-Modul-**Normalfall**.

**Schlüssel:** die reale Datei-Menge ist bereits vorhanden — `FileImports.Path` (aus der
Extraktion) lebt im selben repo-relativen Slash-Namensraum wie Layer-Globs und
`fixed-root`-Kandidaten. Der Fix braucht **kein** zusätzliches Datei-I/O.

## 3. Design (datei-mengen-bewusste Auflösung — Stufe 2)

1. **File-Set-Index** (einmal pro Auflösungslauf): endungs-agnostischer,
   package==directory-Match der Kandidaten gegen die reale Datei-Menge (Verzeichnis-
   Existenz / Endung strippen). Ordnungsfrei ⇒ deterministisch (stabil sortiert).
2. **Auflösung datei-mengen-bewusst:** internen FQN über **alle** Roots gegen die reale
   Datei-Menge suchen. Disjunkte Sub-Namespaces matchen in **höchstens einem** Root ⇒
   eindeutig. Ein Phantom (0 reale Treffer) **bleibt extern** — nie einer Schicht
   zugeordnet (Tech-Regeln greifen unverändert am Roh-Symbol).
3. **Schicht-Zuordnung** über die `layers`-Globs am **Pfad des überlebenden realen
   Kandidaten**, nicht am Root-Präfix über Phantomen.
4. **Echte Mehrdeutigkeit → fail-closed (Exit 2, nach dem Scan):** gleicher FQN real in
   ≥ 2 Roots **mit verschiedenen Schichten** (§5 Entscheid 1: distinct-layer). Signalisiert
   als `error` aus der Auflösung, den die CLI wie die bestehenden Lade-/Extraktions-Fehler
   auf Exit 2 abbildet — **neues Plumbing** (`Evaluate → cli.Run`), kein bloßer if-Zweig
   (§4 Punkt 4). Findings unterdrückt, stderr-Meldung, stdout leer.
5. **Statischen Guard entfernen:** die datei-mengen-bewusste Auflösung ist strikt
   präziser — sie fängt den echten Bug (§6 AC2) **und** akzeptiert die legitime Config
   (AC4). Ein reduziert erhaltener Guard bräche nur AC4 wieder. Der Folge-ADR supersedet
   [ADR-0020](../../adr/0020-mehr-wurzel-phantom-guard.md).

**Rückwärtskompatibilität (beweisbar):** nur der `fixed-root`-Pfad mit **≥ 2 Roots**
ändert Verhalten. `path`/`relative`/1-Root/`package_base`-ohne-Roots/kein-`resolution`
werden identisch durchgereicht.

## 4. Geplanter Umfang (nach Abnahme)

1. **Spec — Lastenheft** ([AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml)):
   das bisherige AK „Mehr-Wurzel-Phantom → Exit 2" durch die neue Auflösungs-Semantik
   ersetzen — mit **drei AK** (Happy/Boundary/Negative,
   [conventions §Anforderungs-Anlege-Prozess](../../../../harness/conventions.md#anforderungs-anlege-prozess));
   die bestehende „Happy (Auflösung)" bleibt erhalten:
   - **Happy (Multi-Modul disjunkt):** disjunkte Roots + geteiltes `package_base` → interner FQN
     löst datei-mengen-bewusst auf seine reale Schicht; `domain → application` wird gemeldet.
   - **Boundary (`expect`/`actual`):** gleicher FQN real in ≥ 2 Roots, **gleiche** Schicht → löst
     sauber (kein Exit 2).
   - **Negative (echte Mehrdeutigkeit):** gleicher FQN real in ≥ 2 Roots, **verschiedene**
     Schichten → Exit 2 (nach dem Scan).
   Version-Bump 0.16.0 → 0.17.0 + Historie; die Umschreibung **führt die distinct-layer-
   Verfeinerung explizit mit** (Abweichung vom bisherigen pauschalen Reject). **Out-of-Scope
   ergänzen:** asymmetrisches `expect`/`actual`, verschachtelte-Klassen-Importe, datei-tiefe
   Globs, per-Root `package_base` bleiben ausgewiesene Grenzen
   ([AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)).
   AC-Änderung **nur** im Lastenheft.
2. **Spec — Spezifikation** ([SPEC-CONF-001](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema)):
   den Mehr-Wurzel-Absatz von der Ladezeit-Guard-Form auf die Scan-Zeit-datei-mengen-bewusste
   Auflösung umschreiben.
3. **ADR:** neuer ADR **Supersedes [ADR-0020](../../adr/0020-mehr-wurzel-phantom-guard.md)**;
   trägt zusätzlich die **Re-Evaluierungs-/Erweiterungs-Relation** zu
   [ADR-0016](../../adr/0016-resolution-sprach-parametrisch.md) /
   [ADR-0014](../../adr/0014-resolution-roots.md) (deren fixed-root-Root-Prepend-Semantik wird
   faktisch datei-mengen-bewusst erweitert; 0014/0016 bleiben immutabel). Provenance-Marker,
   slice-token-frei ([MR-005](../../../../harness/conventions.md#mr-005--referenzmatrix-intra-spec-richtung--adrslice-disziplin-d-check-angleichung)).
   Die bisherige Guard-ADR → Status Superseded (Header-Pointer + ADR-Index).
4. **Code:**
   - **Plumbing (F1, neu):** `core.Evaluate(m, files)` → `([]Finding, error)`; die reale
     Datei-Menge (File-Set-Index) fließt `Evaluate → ruleFor → targetLayer`; `targetLayer` gibt
     zusätzlich `error` zurück (echte Mehrdeutigkeit). `cli.Run` (`cli.go:57`) mappt den Fehler
     wie `config.Load`/`Extract` auf Exit 2. Signatur-Ripple über **alle**
     `Evaluate`/`targetLayer`-Aufrufer (v. a. `rules_test.go`).
   - **Auflösung:** File-Set-Index + `denotes` (paket-granular, §5 Entscheid 2) + `filterReal`
     (greift **nur** bei `fixed-root` mit ≥ 2 Roots); Schicht am realen Kandidaten-Pfad.
   - **Guard raus:** `PhantomRootConflictIn` **und der Guard-only-Helfer `rootForcedLayer`**
     (`rules.go:281-319`), Typ `PhantomRootConflict` (`model.go`), Guard-Aufruf in
     `config.resolveAndCheck` (`config.go:136-153`) → `Load` ruft direkt `decodeResolution`.
   - **`--print-config` (F8, Neu-Aufwand):** `sampleConfig` (`cli.go:76-97`) trägt heute
     **keinen** `resolution`-Block → ein Multi-Modul-`resolution`-Beispiel **von Null** ergänzen (AC6).
5. **Tests (F3):**
   - **Neu:** E2E (Muster der bestehenden `fixed-root`-E2E-Tests) für AC1/AC2/AC4/AC5 + Unit für
     Index/`denotes`/`filterReal`/Ambiguität.
   - **Umschreiben/entfernen** (Verhaltenswechsel, ehrlich dokumentiert):
     `TestPhantomFlatGlobsMisresolves` (`rules_test.go:912`, umschreiben → jetzt korrekt),
     `TestPhantomDeepGlobsResolvesCorrectly` (`:928`, Signatur), `TestPhantomRootConflict`
     (`:1325` + Aufrufer `:1394`, entfernen — testet den gelöschten Guard). **Config-Ebene:**
     `TestResolutionMultiRootPhantomFailsClosed` (`config_test.go:434`, invertieren → lädt jetzt;
     das Exit-2 wandert in den E2E-AC5), `…ConflictLanguageDeterministic` (`:462`, entfernen);
     `…DeepGlobsValid` (`:486`) / `…SameLayerValid` (`:505`) **behalten**; Fixtures
     `kmpFlat`/`kmpDeep` anpassen.
6. **Gates:** `make gates` + `make ci` + `make trace-check`; AC1–AC6 = Fitness-Function.

## 5. Abnahme-Entscheide (abgenommen 2026-07-05, Maintainer-Review)

1. **AC5-Schärfe → (a) distinct-layer.** Exit 2 nur bei gleichem FQN real in ≥ 2 Roots mit
   **verschiedenen** Schichten; gleiche Schicht (`expect`/`actual`) löst sauber. Die
   Lastenheft-Umschreibung (§4 Punkt 1) führt die Abweichung vom bisherigen pauschalen
   Reject explizit mit. *(Verworfen: (b) literal — bräche same-layer-`expect`/`actual`.)*
2. **`denotes`-Granularität → (a) paket-granular.** Verzeichnis-Existenz-Zweig deckt Kotlins
   Klasse≠Datei (Top-Level-Deklaration in beliebiger `.kt`); Restgrenze verschachtelte-Klassen-
   Importe ist [AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)-Terrain.
   *(Verworfen: (b) blatt-genau — präziser bei AC5, aber Klasse≠Datei-Falsch-Negativ.)*
3. **Per-Root `package_base` → (a) out-of-scope.** Der disjunkte Fall löst mit geteiltem
   `package_base` (die Sub-Namespaces diskriminieren); additiver Folge-Slice.

## 6. Akzeptanzkriterien (Fitness-Function, als Tests)

Referenz-Fixtures:

```
mod-a/src/commonMain/kotlin/com/ex/domain/A.kt         // package com.ex.domain
mod-b/src/commonMain/kotlin/com/ex/application/B.kt     // package com.ex.application; import com.ex.domain.A (erlaubt)
```

- **AC1 (positiv):** sauber → 0 Befunde, Exit 0.
- **AC2 (der Bug):** `import com.ex.application.B` in `mod-a/.../domain/A.kt` → **1 Befund**
  „domain → application verboten", Exit ≠ 0. *(heute: 0)*
- **AC3:** Tech-Leak (Roh-Symbol) unverändert scharf.
- **AC4:** disjunkte Multi-Root-Config wird **akzeptiert** (kein Exit-2-Reject).
- **AC5:** echt ambige Config — gleicher FQN real in ≥ 2 Roots **und verschiedene Schichten**
  (§5 Entscheid 1: distinct-layer) → Exit 2.
- **AC6:** `--print-config` dokumentiert die Multi-Modul-Resolution.

## 7. Grenzen / Folge

- **`expect`/`actual`** (gleicher FQN legitim in zwei Source-Sets) ist mit §5 Entscheid 1(a)
  same-layer-sauber; die *asymmetrische* Restform und verschachtelte-Klassen-Importe
  bleiben ausgewiesene Heuristik-Grenzen ([AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)).
- **„0 reale Treffer ⇒ extern" (neue, engere Grenze, F9):** ein interner Import, dessen
  Paket-Verzeichnis unter **keiner** konfigurierten Root in der gescannten Datei-Menge liegt
  (fehlkonfigurierte/nicht gelaufene Source-Set), bleibt still extern → potenzielles
  Falsch-Negativ. Gegenüber einer **korrekt** konfigurierten Basis kein neuer Fehler; im Sinne
  von [AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)
  hier ausgewiesen (die Verzeichnis-Existenz-Matchung entschärft den Klasse≠Datei-Fall).
- **Endungs-Agnostik** gilt für verzeichnisbasierte Globs (package==directory); datei-tiefe
  Globs sind eine dokumentierte Grenze.
- **Per-Root `package_base`** und Namespace-Index bleiben additive Folge-Slices.

## 8. Umsetzungs-Review-Härtung (2026-07-05, adversarisches Multi-Linsen-Review)

Drei Linsen (Core-Korrektheit / Spec-Konformität / Test-Härte). Linse B: konform (Guard-
Entfernung ADR-gedeckt + netto präziser, keine Harness-Lüge). Behobene Befunde:

- **A1 (kritisch, reproduziert):** `denotes` akzeptierte im Symbol-Zweig ein **Phantom**, dessen
  Elternverzeichnis nur zufällig existiert (Klasse direkt unter `package_base` / **Split-Package**)
  → **spurioses Exit 2** für legitimen Code. Fix: `strength`-Stufen (Datei-exakt/Paket-Verzeichnis
  = 2 sticht Nur-Elternpaket = 1); `filterReal` behält die stärkste vorhandene Stufe. Neuer E2E
  pinnt es.
- **C1/C6:** Ambiguitäts-Zeugen-Determinismus (`errors.As` + `LayerA`/`LayerB`-Reihenfolge) neu gepinnt.
- **C2:** zwei durch den Guard-Abbau vakuant gewordene Config-Ladetests entfernt.
- **C3:** AC5-stderr-Assertion auf eine **scan-zeit-spezifische** Phrase geschärft (der alte Guard
  erzeugte denselben Kern-Text).
- **C4/C5:** same-layer-`expect`/`actual` und **Wildcard-Import** als E2E ergänzt.
- **B1/B2:** [MR-005](../../../../harness/conventions.md#mr-005--referenzmatrix-intra-spec-richtung--adrslice-disziplin-d-check-angleichung)/`.d-check.yml`-Provenance-Marker-Wortlaut an die gelebte
  slice-token-freie Praxis angeglichen; Slice-Status-Header aktualisiert.
- **Handbuch-Currency (Nachtrag, Maintainer-Frage):** der `resolution`-Abschnitt im
  Benutzerhandbuch kannte nur Ein-Root-Rezepte — das **Multi-Modul (KMP)**-Rezept ergänzt und
  die alte „paket-tiefe Globs"-Guard-Empfehlung getilgt (Handbuch-Version 1.26). §4 hatte den
  Benutzerhandbuch als Deliverable **nicht gelistet** — Scope-Lücke, nachgezogen.

## 9. Closure

**done 2026-07-05.** `make gates` grün (arch-check Dogfooding 0, doc-check 0/87, coverage ok,
gate-consistency ok, record-gates geschrieben) + `make image-test` OK (`--print-config` dekodierbar).
[ADR-0022](../../adr/0022-datei-mengen-bewusste-mehr-wurzel-aufloesung.md) `Accepted`,
[ADR-0020](../../adr/0020-mehr-wurzel-phantom-guard.md) → `Superseded` (Header + Index; **erster
Supersede** im Repo — `matrix`/`status.forbidden` bleibt grün, da adr→adr/slice→adr keine
gate-regierten Kanten sind). Slice nach `done/` (reiner `git mv`,
[AGENTS §3.3](../../../../AGENTS.md#33-git-mv--inhaltsänderung--zwei-commits)); Roadmap-Link + Handbuch nachgezogen.

**Lerneintrag:** Das adversarische Umsetzungs-Review fand einen **echten kritischen Befund (A1)**,
den die grünen AC-Tests **nicht** zeigten: der erste `denotes` akzeptierte im Symbol-Zweig ein Phantom,
dessen Elternverzeichnis nur zufällig existiert (Klasse direkt unter `package_base` / Split-Package) →
**spurioses Exit 2** für legitimen Code. Lehre: bei datei-mengen-bewusster Auflösung ist die **Stärke
der Evidenz** entscheidend (datei-exakt sticht nur-Elternpaket), nicht bloß „existiert irgendwas unter
dem Pfad" — ein Prädikat, das „das Elternverzeichnis existiert" mit „das Symbol existiert" verwechselt,
erzeugt Phantom-Positive. Zudem: eine Verhaltens-Änderung (Guard→Auflösung) macht Bestands-Tests still
**vakuant** — sie müssen als nicht-mehr-pinnend erkannt und ersetzt werden, nicht bloß grün gelassen.
