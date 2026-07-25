# slice-031 — Deklarations-bewusste Auflösung für Split-Packages (Bug + Feature)

**Status:** **done** (2026-07-06) — spec-first umgesetzt: Lastenheft + Spezifikation **0.18.0** +
[ADR-0023](../../adr/0023-deklarations-bewusste-mehr-wurzel-aufloesung.md) `Accepted` (**Weg A**,
Supersedes [ADR-0022](../../adr/0022-datei-mengen-bewusste-mehr-wurzel-aufloesung.md)) + Code +
Tests (AC1–AC7 + Wildcard-Split) + Benutzerhandbuch 1.27; **`make ci` grün** + real gegen d-migrate
verifiziert (Auslöser-Exit-2 getilgt). Closure §9.
**Typ:** Bug (spurioser Exit 2 / stiller Adoptions-Blocker) + Feature
(Symbol-/Deklarations-Index), konsumenten-getrieben (**d-migrate**).
**Bezug:** schärft [AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml)
(Mehr-Wurzel-Auflösung + fail-closed) und
[AC-FA-EXTRACT-001](../../../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion)
(die Extraktion erhält die Fähigkeit, zusätzlich Top-Level-Deklarationen zu liefern);
bewegt sich an der Heuristik-Grenze
[AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze).
**Stufe 3** zur Mehr-Wurzel-Auflösung: re-evaluiert
[ADR-0022](../../adr/0022-datei-mengen-bewusste-mehr-wurzel-aufloesung.md) (datei-mengen-bewusst,
[slice-027](slice-027-kmp-multimodul-resolution.md)) sowie
[ADR-0016](../../adr/0016-resolution-sprach-parametrisch.md) /
[ADR-0014](../../adr/0014-resolution-roots.md); text-heuristisch nach
[ADR-0002](../../adr/0002-text-heuristische-extraktion.md). Ein Folge-ADR
(**Supersedes [ADR-0022](../../adr/0022-datei-mengen-bewusste-mehr-wurzel-aufloesung.md)**,
Provenance-Marker/slice-token-frei nach
[MR-005](../../../../harness/conventions.md#mr-005--referenzmatrix-intra-spec-richtung--adrslice-disziplin-d-check-angleichung))
entsteht mit der Umsetzung. Löst den reservierten „Symbol-/Namespace-Index" ein
([ADR-0016](../../adr/0016-resolution-sprach-parametrisch.md) `mode`-Diskriminator,
[slice-021](slice-021-csharp-backend.md) Namespace-Index reserviert,
[slice-027](slice-027-kmp-multimodul-resolution.md) §7). [Roadmap](../in-progress/roadmap.md).
**Evidenz:** a-check `v0.11.0` real gegen `pt9912/d-migrate` reproduziert — die einzige
korrekte Vollrichtungs-Config (`fixed-root`, ein Root pro Gradle-Modul) endet in **Exit 2**,
Scan komplett abgebrochen, null Befunde.

## 1. Auslöser

d-migrate ist ein **JVM-Kotlin-Hexagon** (kein KMP) mit **konzept-basierten** Packages
(`dev.dmigrate.{driver,format,server,…}`), die über mehrere Gradle-Module/Schichten spannen —
die Hexagon-Schicht steckt im **Modul-Verzeichnis**, nicht im Package-Namen. Idiomatisches
Kotlin legt dabei **Top-Level-Deklarationen** (Extension-Funktionen, mehrere Klassen je Datei)
in Packages ab, die real über zwei Module verteilt sind (**Split-Package**). Die einzige Config,
die das Schicht-Gate scharf stellt, ist `resolution.kotlin = {mode: fixed-root, package_base:
dev.dmigrate, roots: [<27 Modul-Roots>]}`. Sie wird abgelehnt:

```
a-check: mehrdeutige Mehr-Wurzel-Auflösung: "dev.dmigrate.driver.connection.asJdbc"
 (in adapters/driven/driver-common/.../driver/data/AbstractJdbcDataReader.kt)   ← PRODUKTIV-Importer
 existiert real unter mehreren Wurzeln und fällt in verschiedene Schichten —
 "adapters" (adapters/driven/driver-common/.../driver/connection/asJdbc) vs.
 "ports"    (hexagon/ports-common/.../driver/connection/asJdbc)                 → Exit 2
```

Zwei reale Beispiele (empirisch nachgewiesen), beide **eindeutig in genau einem Modul deklariert** —
die Mehrdeutigkeit ist ein reines Werkzeug-Artefakt:

| Symbol | Import (Produktiv) | Deklariert in (real, genau 1×) | Wahre Schicht |
|---|---|---|---|
| `asJdbc` (Top-Level-Extension-Fun) | `dev.dmigrate.driver.connection.asJdbc` | `adapters/driven/driver-common/.../driver/connection/JdbcDatabaseConnection.kt` | **adapters** |
| `ChunkColumnSchema` (Klasse) | `dev.dmigrate.format.data.ChunkColumnSchema` | `hexagon/ports-common/.../format/data/ChunkSchema.kt` | **ports** |

Das Paket `dev.dmigrate.driver.connection` liegt real unter **beiden** Roots
(`hexagon/ports-common/` = Port-Interfaces/Config, `adapters/driven/driver-common/` = JDBC-Impls) —
eine Hexagon-Grenze **durch ein Paket**. Legales JVM; genau der Fall, der Symbol-Auflösung nötig macht.

**Systemisch, kein Einzelfall:** d-migrate hat **~17** solcher schicht-übergreifenden Split-Packages
und **~32** Top-Level-Symbol-Importe. Der `markers.ignore_symbols`-Stopgap greift zwar vor der
Auflösung (unterdrückt Exit 2), ist aber **bewiesen untauglich**: `asJdbc` ignoriert ⇒ sofort der
nächste Tripwire (`ChunkColumnSchema`) — Whack-a-Mole über genau die Grenz-Symbole, die das Gate
prüfen soll. Der Exit 2 bricht zudem den **ganzen** Scan ab (keine Befunde), also blockiert der
Fall die Adoption **vollständig**, nicht nur punktuell.

## 2. Kern-Erkenntnis (Fehlerquelle im Code)

Die datei-mengen-bewusste Auflösung aus [slice-027](slice-027-kmp-multimodul-resolution.md)
kennt zwei Evidenzstufen (`strength`, `internal/hexagon/core/rules.go`):

- **Stufe 2** — der Kandidat `<root>/<pkg>/<sym>` entspricht einer real gescannten **Datei**
  (endungs-agnostisch): trifft nur, wenn die Datei **wie das Symbol heißt**
  (`ConnectionPool` → `ConnectionPool.kt`).
- **Stufe 1** — nur das **Elternverzeichnis** (`<root>/<pkg>`) existiert.
- **Stufe 0** — Phantom ⇒ bleibt extern.

`filterReal` behält die stärkste vorhandene Stufe über alle Roots (A1-Fix aus slice-027, so gewinnt
eine echte Datei über ein bloßes Elternverzeichnis). Für ein Symbol, dessen **Datei ≠ Symbolname**
(Extension-Fun, zweite Klasse je Datei — idiomatisches Kotlin), gibt es **nirgends** Stufe 2 —
nur Stufe 1 unter *jedem* Root, dessen Paketverzeichnis existiert. Bei einem Split-Package über zwei
Schicht-Roots ⇒ Stufe 1 in beiden ⇒ `targetLayer` sieht zwei reale Kandidaten in verschiedenen
Schichten ⇒ `AmbiguousResolution` ⇒ Exit 2.

**Die fehlende Evidenz ist genau die, die den Fall entscheidet:** *welches Modul das Symbol
tatsächlich deklariert.* Der `fileIndex` (`rules.go`) indiziert bewusst nur **Pfade und
Verzeichnisse**, **keine Deklarationen** — ausgewiesene Heuristik-Grenze
([AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)).
Der Quelltext jeder gescannten Datei ist zur Extraktionszeit bereits gelesen; ein Deklarations-Index
braucht **kein** zusätzliches Datei-I/O — dieselbe Beobachtung, die slice-027 trug. Die
Deklarationen selbst sind jedoch **neue Nutzlast** und müssen durch Extraktions-Port und Core-Modell
fließen (§4.4).

## 3. Design (deklarations-bewusste Auflösung — Stufe 3)

**Leitbegriff: das *deklarations-bewusste Backend*.** slice-027 belegt die Schicht-Zugehörigkeit
eines Symbols über **Pfad**-Evidenz (Datei==Symbolname / Elternverzeichnis). Für Symbole mit
**Datei ≠ Symbolname** ist Pfad-Evidenz zu schwach. Der echte Beweis, dass ein Symbol unter einem
Root lebt, ist, dass eine dort gescannte Datei es **deklariert**. Backends, die Deklarationen
extrahieren (0.18.0: **nur Kotlin**), lösen deklarations-bewusst auf; alle übrigen liefern ein
**leeres** Deklarations-Set (No-op) und bleiben unverändert bei der slice-027-Pfad-Semantik.

**Evidenz-Rangfolge (deklarations-bewusstes Backend), stärkste zuerst — löst [High-2]:**

1. **deklariert** — eine gescannte Datei in `<root>/<pkg>/` trägt eine **Top-Level-Deklaration**
   `<sym>`, unabhängig vom Dateinamen.
2. **nur-Paketverzeichnis** — `<root>/<pkg>/` existiert, aber keine gescannte Datei deklariert
   `<sym>`. **Auch eine bloß gleichnamige Datei `<sym>.kt`, die `<sym>` nicht deklariert, zählt nur
   hier** — für ein deklarations-bewusstes Backend ist der Dateiname **kein** Beweis (die
   slice-027-Stufe 2 wird durch echte Deklarations-Evidenz **abgelöst**).
3. **keine** (0) — Phantom, bleibt extern.

**Kritischer Fall, explizit (der Review-Punkt):** Root A hat `Foo.kt`, deklariert aber `Foo`
**nicht**; Root B deklariert `Foo` in `Types.kt`. → **Root B (deklariert, Stufe 1) sticht Root A
(nur-Paketverzeichnis, Stufe 2)** — echte Deklaration gewinnt, kein Falsch-Root-Match, kein Exit 2.
Für **nicht** deklarations-bewusste Backends bleibt slice-027 unverändert (Datei-Match > Verzeichnis).

**Ablauf:**

1. **Deklarations-Index** (einmal pro Auflösungslauf, I/O-frei aus dem schon gelesenen Quelltext):
   je gescannte Datei die Menge ihrer Top-Level-Deklarationsnamen (Kotlin: `fun` inkl. Extension
   `fun R.name`, `val`/`var`/`const val`, `class`/`data class`/`sealed class`/`enum class`/
   `annotation class`, `object`, `interface`, `sealed interface`, `typealias`). Regex/text-heuristisch
   ([ADR-0002](../../adr/0002-text-heuristische-extraktion.md)); ordnungsfrei ⇒ deterministisch
   ([AC-QA-01](../../../../spec/lastenheft.md#ac-qa-01--determinismus)).
2. **Auflösung — stärkste Evidenzstufe entscheidet.** Bestimme je Root die Stufe (deklariert >
   nur-Paketverzeichnis > keine); nimm die **stärkste über alle Roots vorhandene** Stufe und
   betrachte nur Roots auf dieser Stufe:
   - **Genau ein** Root (oder alle auf **dieselbe** Schicht) ⇒ auflösen auf diese Schicht. **(löst
     `asJdbc`→adapters, `ChunkColumnSchema`→ports über die Stufe *deklariert*)**
   - **≥ 2** Roots in **verschiedenen** Schichten ⇒
     - Stufe **deklariert**: **Exit 2** (echte Mehrdeutigkeit, fail-closed nach dem Scan; im gültigen
       JVM selten — ein FQN wird höchstens einmal deklariert, `expect`/`actual` ist same-layer, §6 AC5).
     - Stufe **nur-Paketverzeichnis**: **extern (fail-open)** — ohne Deklaration nicht diskriminierbar;
       genau der d-migrate-Fall, den slice-027 fälschlich mit Exit 2 quittierte. Neue, engere
       [AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)-Grenze
       (§7), konsistent mit slice-027s „0 reale Treffer ⇒ extern".
   - Stufe **keine** (nirgends Evidenz) ⇒ extern (Phantom).

   **Damit ist die Rolle von *nur-Paketverzeichnis* eindeutig:** ein Kotlin-Symbol **ohne**
   Deklaration, aber mit **genau einem** Paketverzeichnis (oder mehreren derselben Schicht), **löst** —
   die slice-027-Paket-Verzeichnis-Auflösung bleibt als Stufe *nur-Paketverzeichnis*
   **rückwärtskompatibel** erhalten; **extern** wird es nur bei Paketverzeichnissen in **≥ 2
   verschiedenen** Schichten ohne diskriminierende Deklaration.
3. **Subsumiert den slice-027-A1-Fix:** der „Klasse direkt unter `package_base` / Split-Package"-
   Phantomfall löst jetzt über echte Deklaration statt über die Datei-exakt-Stratifizierung.
4. **Rückwärtskompatibilität (beweisbar):** die neue Evidenz greift nur für deklarations-bewusste
   Backends (Kotlin) und nur bei `fixed-root` mit ≥ 2 Roots. Eine bislang **eindeutige** Auflösung
   bleibt eindeutig (im Normalfall Datei==Deklaration); eine bislang **externe** bleibt extern, wenn
   nichts deklariert wird. Nicht-Kotlin, `path`/`relative`/1-Root/kein-`resolution` und der
   belief-agent-KMP-Fall (Dateien nach Klassen benannt ⇒ Datei==Deklaration) laufen unverändert
   (§6 AC6).

**Verankerung: Weg A — entschieden** ([ADR-0023](../../adr/0023-deklarations-bewusste-mehr-wurzel-aufloesung.md) `Accepted`):

- **(A) additive Evidenzstufe in `fixed-root`** *(gewählt)* — kein neuer `mode`; der Deklarations-Index
  wird bei `fixed-root` mit ≥ 2 Roots automatisch konsultiert. Konsumenten-Config unverändert
  (d-migrates ehrliche Config löst dann auf), rückwärtskompatibel wie slice-027, Supersedes
  [ADR-0022](../../adr/0022-datei-mengen-bewusste-mehr-wurzel-aufloesung.md).
- **(B) eigener opt-in `mode`** *(verworfen)* — Konsumenten müssten umstellen und der Modus überlappt
  konzeptionell mit dem für C# reservierten Namespace-Index
  ([slice-021](slice-021-csharp-backend.md)); fragmentiert das Modell.

## 4. Geplanter Umfang (nach Abnahme)

1. **Spec — Lastenheft**
   ([AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml)):
   den Mehr-Wurzel-Absatz + das AK „Negative (Mehr-Wurzel, echte Mehrdeutigkeit)" von der
   package==directory- auf die **deklarations-bewusste** Semantik umschreiben (Exit 2 nur bei
   **Deklaration** in ≥ 2 verschiedenen Schichten; ein in genau einem Root deklariertes Symbol löst;
   kein Deklarations-Treffer ⇒ extern). Neues AK „Happy (Klasse≠Datei / Top-Level-Symbol)". **Out-of-
   Scope** aktualisieren: der Klasse≠Datei-Fall wandert **in Scope**; die neue engere Grenze (kein
   Deklarations-Treffer ⇒ extern) + verschachtelte-Klassen-Importe + Kommentar/String-Fehltreffer
   bleiben [AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze).
   [AC-FA-EXTRACT-001](../../../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion):
   die Extraktion erhält die **Fähigkeit**, Top-Level-Deklarationen zu liefern; **in 0.18.0 befüllt
   nur das Kotlin-Backend sie**, alle übrigen liefern ein **leeres** Deklarations-Set (No-op, kein
   Verhaltenswechsel — ausgewiesene Grenze, §5-Entscheid 3). **Version-Bump 0.17.0 → 0.18.0** +
   Historie-Zeile. Drei AK (Happy/Boundary/Negative) + Out-of-Scope je geänderter Anforderung
   ([conventions §Anforderungs-Anlege-Prozess](../../../../harness/conventions.md#anforderungs-anlege-prozess)).
   **AC-Änderung nur im Lastenheft.**
2. **Spec — Spezifikation**
   ([SPEC-CONF-001](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema) /
   [SPEC-EXTRACT-001](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion) /
   [SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung)): den
   Auflösungs-Algorithmus um die Deklarations-Evidenzstufe (inkl. Rangfolge §3) und den
   Extraktions-Zusatz (Deklarations-Set, Kotlin befüllt / übrige no-op) schärfen.
3. **ADR:** neuer ADR **Supersedes
   [ADR-0022](../../adr/0022-datei-mengen-bewusste-mehr-wurzel-aufloesung.md)**; trägt die Re-Eval-Relation
   zu [ADR-0016](../../adr/0016-resolution-sprach-parametrisch.md) /
   [ADR-0014](../../adr/0014-resolution-roots.md); hält **Weg A** fest (§3); `Schärft:` aufwärts auf
   die `SPEC-*`-Stelle; ADR-Index-Update; Provenance-Marker/slice-token-frei
   ([MR-005](../../../../harness/conventions.md#mr-005--referenzmatrix-intra-spec-richtung--adrslice-disziplin-d-check-angleichung)).
   [ADR-0022](../../adr/0022-datei-mengen-bewusste-mehr-wurzel-aufloesung.md) → Status `Superseded`
   (Header-Pointer + Index) — zweiter Supersede im Repo.
4. **Code (Datenmodell-Plumbing explizit — löst [Medium-1]):** die Deklarationen sind **neue Nutzlast**;
   die slice-027-Behauptung „kein neues Plumbing" gilt **nur** für den Fehler-Kanal, nicht für die Daten.
   - **Core-Modell** (`internal/hexagon/core/model.go`, heute `FileImports` ohne Deklarationen):
     das Extraktions-Ergebnis je Datei trägt zusätzlich ein Deklarations-Feld (`Declarations []string`).
   - **Extraktions-Port + Backend** (`internal/adapter/driven/extract/extract.go`, heute nur Imports):
     der Port liefert Deklarationen **neben** den Importen; das **Kotlin**-Backend füllt sie
     (Top-Level-Decl-Muster), alle anderen Backends geben ein leeres Set zurück (No-op).
   - **Index + Auflösung** (`internal/hexagon/core/rules.go`, heute `fileIndex` nur Pfade/Verzeichnisse):
     `fileIndex` gewinnt eine Deklarations-Menge (`dir → {declNames}`); neue Evidenzstufe/Rangfolge in
     `strength`/`denotes`; `filterReal`/`targetLayer` nutzen sie. Der **Fehler-Kanal**
     `Evaluate → ruleFor → targetLayer` aus slice-027 wird wiederverwendet (kein neues Fehler-Plumbing).
   - **`--print-config`:** das Multi-Modul-`resolution`-Beispiel um den Split-Package-Hinweis ergänzen.
5. **Tests:** E2E für AC1–AC6 (Muster der bestehenden `fixed-root`-E2E); Unit für Deklarations-
   Extraktion (je Decl-Art + Extension-Fun + Kommentar/String-Negativprobe), die Evidenz-Rangfolge
   (kritischer Fall AC3) und die Ambiguität. **Regressionspin:** ein belief-agent-artiges disjunktes
   Fixture bleibt byte-identisch (AC6).
6. **Benutzerhandbuch:** `resolution`-Abschnitt um das **Split-Package**-Rezept ergänzen (Modul-Root je
   Schicht, Deklarations-Auflösung, Kotlin-only-Grenze); Currency-Bump.
7. **Gates:** `make gates` + `make ci` + `make trace-check`; AC1–AC6 = Fitness-Function.

## 5. Entscheide & offene Fragen

**Entschieden (im Slice fixiert, damit ADR/Spec/Tests nicht auseinanderlaufen):**

1. **Verankerung → Weg A** (additive Evidenzstufe in `fixed-root`, kein neuer `mode`,
   rückwärtskompatibel, keine Konsumenten-Umstellung) —
   [ADR-0023](../../adr/0023-deklarations-bewusste-mehr-wurzel-aufloesung.md) `Accepted`.
2. **Rolle *nur-Paketverzeichnis* / Residual** (§3): die Stufe *nur-Paketverzeichnis* **löst
   rückwärtskompatibel**, wenn sie eindeutig ist (genau ein Root bzw. alle dieselbe Schicht); **extern
   (fail-open)** nur bei Paketverzeichnissen in **≥ 2 verschiedenen** Schichten ohne diskriminierende
   Deklaration. *Deklariert* dagegen fail-closed (Exit 2) bei ≥ 2 verschiedenen Schichten — die
   Asymmetrie ist der Kern des Fixes.
3. **Evidenz-Rangfolge** (§3): echte Deklaration sticht den bloßen Datei-Namens-Match.
4. **Backend-Umfang 0.18.0** → **nur Kotlin** ist deklarations-bewusst; alle übrigen liefern ein
   leeres Deklarations-Set (No-op, slice-027-Verhalten unverändert). Java ist syntaxnah und günstig
   als additiver Folge-Slice; der C#-Namespace-Index bleibt eine **eigene** reservierte Achse
   ([slice-021](slice-021-csharp-backend.md)).

**Noch offen (Umsetzung):**

1. **Deklarations-Muster-Präzision** (Kotlin) — welche Formen matchen; Kommentar/String-Stripping vor
   oder nach der Muster-Suche (Python-Lehre aus [slice-020](slice-020-python-backend.md):
   Stripping kann falsch-grün erzeugen — hier umgekehrt ein Kommentar-Treffer falsch-„deklariert").
   Fixierung in Spec + Tests.
2. **Befund-/Fehler-Determinismus** — Ambiguitäts-Zeuge (Symbol, deklarierende Roots, Schichten)
   stabil sortiert, wie der slice-027-C1-Fix.

## 6. Akzeptanzkriterien (Fitness-Function, als Tests)

Referenz-Fixtures (generalisiert aus der d-migrate-Evidenz):

```
port-mod/src/main/kotlin/com/ex/pkg/Types.kt     // package com.ex.pkg; interface Foo   (Datei ≠ Symbol)
adapter-mod/src/main/kotlin/com/ex/pkg/Impl.kt   // package com.ex.pkg; fun Foo.asBar()  (Top-Level-Ext-Fun)
```
`port-mod/**` → Schicht `ports`, `adapter-mod/**` → Schicht `adapters`; geteiltes
`package_base: com.ex`, zwei Roots (Split-Package `com.ex.pkg`).

- **AC1 (der Bug / Top-Level-Fun):** eine Adapter-Datei importiert `com.ex.pkg.asBar` (nur in
  `adapter-mod` deklariert) → löst auf **adapters**, **kein Exit 2**, 0 Befunde. *(heute: Exit 2)*
- **AC2 (Klasse≠Datei):** Import `com.ex.pkg.Foo` (Interface in `Types.kt`, nur `port-mod`) → löst auf
  **ports**. *(heute: Exit 2)*
- **AC3 (Rangfolge — echte Deklaration sticht Dateinamen):** Root A trägt eine Datei `Foo.kt`, die
  `Foo` **nicht** deklariert (z. B. nur `FooBuilder`); Root B deklariert `Foo` in `Types.kt`. Import
  `…​.Foo` → löst auf **Root B** (echte Deklaration), kein Falsch-Match, kein Exit 2.
- **AC4 (verbotene Kante wird gemeldet):** eine `ports`-Datei importiert `com.ex.pkg.asBar`
  (Adapter-Symbol) → da `ports → adapters` verboten: **1 Befund** (`wrong-direction`), Exit 1 —
  beweist, dass die Auflösung die Kanten-Prüfung *ermöglicht*, nicht nur Exit 2 unterdrückt.
- **AC5 (echte Mehrdeutigkeit bleibt):** dasselbe Symbol real **deklariert** in ≥ 2 Roots
  verschiedener Schichten → Exit 2 nach dem Scan (fail-closed, deterministischer Zeuge);
  same-layer (`expect`/`actual`) löst sauber.
- **AC6 (Residual: ≥ 2 verschiedene Schichten ohne Deklaration → extern):** ein Symbol, dessen
  Paket-Verzeichnis unter ≥ 2 Roots **verschiedener Schichten** liegt, das aber in **keiner**
  gescannten Datei deklariert ist → bleibt **extern**, kein Exit 2, kein Geister-Befund (fail-open,
  [AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)).
- **AC7 (nur-Paketverzeichnis eindeutig → löst; Rückwärtskompatibilität):** (a) ein Kotlin-Symbol
  **ohne** Deklaration, dessen Paket-Verzeichnis unter **genau einem** Root (bzw. mehreren derselben
  Schicht) liegt → **löst** auf diese Schicht (die slice-027-Auflösung bleibt erhalten, kein Regress).
  (b) die unveränderten Modi (`path`/`relative`/1-Root/kein-`resolution`) und **nicht-Kotlin**-Backends
  → **byte-identisch** zu `v0.11.0`; das belief-agent-KMP-Kotlin-Fixture (fixed-root ≥ 2, Dateien nach
  Klassen benannt ⇒ Datei==Deklaration, [slice-027](slice-027-kmp-multimodul-resolution.md) §6)
  → **identische Befunde**.

*(Lastenheft-Mapping der drei Pflicht-AK: **Happy** = AC1/AC2/AC3, **Boundary** = AC6/AC7, **Negative** = AC5.)*

## 7. Grenzen / Folge

- **Residual (entschieden, §3/§5):** die Stufe *nur-Paketverzeichnis* **löst** rückwärtskompatibel,
  wenn sie eindeutig ist; **extern (fail-open)** nur bei Paketverzeichnissen in **≥ 2 verschiedenen**
  Schichten ohne diskriminierende Deklaration. Ein interner Import, dessen Symbol in **keiner**
  gescannten Datei als Top-Level-Deklaration auftaucht (verschachtelte Klasse, `object`-Member,
  generierter/nicht gescannter Code, Star-Import) **und** dessen Paketverzeichnisse schicht-uneins
  sind, bleibt still extern — potenzielles Falsch-Negativ, ausgewiesen
  ([AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)).
  Bewusst gegenüber dem Alternativ-Exit-2 gewählt, der den **ganzen** Scan bräche.
- **Wildcard-/Paket-Import** (`a.b.*` → `a/b/`): trifft ein **ganzes** Paket-Verzeichnis, also die
  Stufe *nur-Paketverzeichnis* — bei einem Split über Schicht-Grenzen daher **fail-open (extern)**,
  kein Exit 2 (Empirie: d-migrates `dev.dmigrate.driver.*` unter ports- und adapters-Root). Ein
  Wildcard kann nicht auf **ein** Symbol und damit nicht auf eine Schicht festgelegt werden.
- **Text-Heuristik:** Deklarations-Muster sind Regex, kein Parser — Kommentar-/String-Fehltreffer,
  exotische Formatierung, `expect`/`actual`-Feinheiten bleiben Grenzen
  ([ADR-0002](../../adr/0002-text-heuristische-extraktion.md)).
- **Sprach-Umfang:** nur **Kotlin** ist in 0.18.0 deklarations-bewusst; die übrigen Backends bleiben
  bei der slice-027-Pfad-Semantik (leeres Deklarations-Set). Java ist naher Folgekandidat;
  C#-Namespace-Index bleibt eine eigene reservierte Achse
  ([slice-021](slice-021-csharp-backend.md)). Per-Root `package_base` bleibt out-of-scope
  ([slice-027](slice-027-kmp-multimodul-resolution.md) §5).
- **Konsument:** d-migrate ist der erste reale Nutzer; nach Landung liefert der dortige Pilot die
  Fitness-Probe (Vollrichtungs-Config → 0 Befunde bzw. gemeldete echte Verstöße statt Exit 2).
  Verwandt: die P-Rest-Sammlung [slice-025](slice-025-p-rest-generalisierung.md) nennt d-migrate als
  offenen Pilot.

## 8. Umsetzungs-Notizen (nach Bau)

- **Datenmodell-Plumbing** wie geplant (§4.4): `FileImports.Declarations` (`model.go`), Extraktions-Port
  liefert Deklarationen neben Importen (`extract.go`, Kotlin-Backend füllt, übrige no-op), `fileIndex.decls`
  (`rules.go`). `strength` trägt die Stufe 3; `filterReal` gibt `(cands, weak)` zurück; `targetLayer`
  entscheidet Exit-2 (starke Stufe) vs. extern (schwache Stufe) am Konflikt.
- **Kotlin-Deklarations-Regex** an **Spalte 0** verankert (eingerückte Member zählen nicht),
  Kommentar-gestrippt; mehrzeiliger String bleibt Grenze ([AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)).
  Die zu starke „Kommentar/String"-AK der ersten Fassung ehrlich auf „Kommentar + Spalte-0; String = Grenze"
  korrigiert (Lastenheft/Spez).

## 9. Closure

**done 2026-07-06.** `make ci` grün (arch-check Dogfooding 0, doc-check 0/90, coverage ok, test inkl.
**8 neuer AC-Tests**, image-test OK) + `make gates` grün.
[ADR-0023](../../adr/0023-deklarations-bewusste-mehr-wurzel-aufloesung.md) `Accepted`,
[ADR-0022](../../adr/0022-datei-mengen-bewusste-mehr-wurzel-aufloesung.md) → `Superseded` (Header + Index;
**zweiter Supersede** im Repo). Lastenheft/Spez **0.18.0**, Benutzerhandbuch **1.27**. Zwei adversarische
Review-Runden (High/High/Medium je) eingearbeitet — Residual-Eindeutigkeit, Evidenz-Rangfolge, Datenmodell-
Plumbing, Backend-Scope, Planning-Drift. Slice nach `done/` (untracked → direkt, kein `git mv`-Tanz).

**Lerneintrag:** Das **End-to-End-Fahren gegen den realen Konsumenten** (`make build` + a-check:dev gegen
d-migrate) fand einen Fall, den weder die grünen AC-Tests noch die zwei Doku-Reviews zeigten: nach dem
Symbol-Fix schlug ein **Wildcard-Import** (`dev.dmigrate.driver.*`) als zweiter Exit 2 durch. Ein Wildcard
trifft ein **ganzes** Paket-Verzeichnis — Paket-Verzeichnis-Evidenz (schwach, Stufe 1), nicht Datei-exakt
(Stufe 2) — und muss bei einem Split über Schichten **fail-open** sein, nicht Exit 2. Lehre: die
Fitness-Function ist erst mit dem echten Ziel-Repo vollständig; ein Symbol-Auflösungs-Fix muss **jede
Import-Form** (Einzel-Symbol **und** Wildcard) gegen den realen Baum prüfen. Der Auslöser-Exit-2 (`asJdbc`)
ist verifiziert getilgt; a-check läuft real gegen d-migrate end-to-end durch (`lateral-adapter`-Befunde =
d-migrates eigene Architektur, dort via `adapter_sink` zu regeln).
