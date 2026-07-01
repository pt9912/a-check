# slice-015 — Resolution-Roots: Import-Auflösung gegen konfigurierbare Wurzeln (sprach-parametrisch)

**Status:** open (Entwurf zur Abnahme — **Abnahme-Gate §5 offen**, blockiert die DoD).
**Bezug:** setzt [ADR-0014](../../adr/0014-resolution-roots.md) um und **erweitert** ihn per
**Folge-ADR** (§3.0; Sprach-Map + Sprach-Threading, `Supersedes: —`); erweitert
[AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml)
(Schema) + [AC-FA-EXTRACT-001](../../../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion);
ehrliche Heuristik-Grenze [AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze).
[Roadmap](../in-progress/roadmap.md). **Evidenz:** der **Polyglot-Bestand** (Repos in Go, C++, C#,
Python, TypeScript) — nicht nur x-wal: b-cad (C++ `src`-Root), x-wal (JVM, gepunktete Pakete), plus
Go/C#/Python/TS-Repos. Die Architektur (Sprach-Map + Threading + `mode`-Diskriminator) trägt **alle**;
slice-015 füllt **einen** Modus, die zwei übrigen kommen additiv (§4).

## 1. Auslöser (Gate)

Die heutige Auflösung nimmt „Import = wurzel-relativer Pfad" an (`targetLayer`, `rules.go:231` —
matcht den Import gegen die `layers`-Glob-Präfixe per `segIndex` auf `/`, `rules.go:256`). Das hält
**nur** für Sprachen, deren Import *ist* der Pfad (Go: Modulpfad) — und bricht in vier Formen:

- **Fester-Wurzel-dotted** (JVM, Python): `com.x.Y` / `a.b.c` sind gepunktet, kein `/` — `segIndex`
  trifft nie. Braucht Wurzel + Separator-Normalisierung.
- **Wurzel ≠ Scan-Wurzel** (C++, nur `<…>`/quoted-als-Pfad): b-cads Includes sind `src/`-gewurzelt;
  `src/` deklarieren statt raten.
- **Relativ zum File** (TypeScript `./x`; C++ **`"…"`-Includes** relativ zum importierenden File):
  löst gegen den *Ort des Files* auf.
- **Namespace-entkoppelt** (C#): `using Foo.Bar;` ohne Pfad-Bezug.

Dieser Slice liefert **nur den ersten Modus** (fester-Wurzel-dotted). Die letzten beiden (relativ,
namespace) sind andere Auflösungs-Signale und bekommen je einen eigenen Folge-ADR (§5).

## 2. Betroffene Artefakte (vor der Implementierung benannt)

- **Slice-ID:** slice-015.
- **ADR:** **neu Folge-ADR** — erweitert [ADR-0014](../../adr/0014-resolution-roots.md) (Accepted, immutable)
  um (a) die **Sprach-Map** und (b) das **Sprach-Threading**; nach dem Muster, mit dem der ADR selbst
  [ADR-0002](../../adr/0002-text-heuristische-extraktion.md) erweiterte (`Supersedes: —`).
- **AC:** [AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml)
  (Schema `resolution`), [AC-FA-EXTRACT-001](../../../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion)
  (Symbol→Layer je Sprache).
- **Spec:** [SPEC-CONF-001](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema)
  (`resolution`), [SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung)
  (Symbol→Layer), [SPEC-EXTRACT-001](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion)
  (Sprache je Datei).
- **Module:** `internal/adapter/driven/config` (`resolution`-Decode), `internal/hexagon/core`
  (`FileImports.Language` **neu**; `targetLayer`/`ruleFor` sprach-bewusst), `internal/adapter/driven/extract`
  (setzt `FileImports.Language` aus `langFor`).
- **Version:** Lastenheft/Spezifikation → **nächste freie Minor** (0.10.0, falls slice-015 vor
  [slice-013](slice-013-driving-driven-vertiefung.md) landet — beide sind Entwurf; wer zuerst mergt,
  nimmt 0.10.0).
- **Gates:** `make gates` → `make ci`.

## 3. Umfang (Reihenfolge: ADR → Lastenheft → Spec → Code → Tests)

0. **Folge-ADR** `Proposed → Accepted` (Sign-off): begründet Sprach-Map + Sprach-Threading als
   Erweiterung (`Supersedes: —`); in den [ADR-Index](../../adr/README.md).
1. **`resolution`-Block als Map Sprache → Config** mit **`mode`-Diskriminator** (Mono-Repo- **und**
   estate-tauglich, §4):
   ```yaml
   resolution:                                        # Map Sprache -> Config
     go:     {mode: path}                             # Import = Pfad (Default; == weggelassen)
     cpp:    {mode: fixed-root, roots: ["src"]}       # Include-Root (b-cad), nur <…>-Includes
     kotlin: {mode: fixed-root, roots: ["src/main/kotlin"], package_base: "com.xwal"}  # dotted
     # typescript: {mode: relative}    # reserviert (C++-"…"/TS) -> Folge-ADR
     # csharp:     {mode: namespace}   # reserviert (C#)         -> Folge-ADR
   ```
   strict-decode. **slice-015 implementiert `mode ∈ {path, fixed-root}`**; `relative`/`namespace` sind
   **reserviert** und brechen bis zum jeweiligen Folge-ADR mit **Exit 2** (kein stiller No-Op — konsistent
   mit slice-017). **Fehlt `resolution` (oder eine Sprache darin) → heutiges Verhalten** (Import-als-Pfad),
   rückwärtskompatibel; `go: {mode: path}` und *weggelassen* sind verhaltens-identisch (Testfall §6).
2. **`FileImports.Language`** (core, neu) — von `extract` aus `langFor` gesetzt; **durchgereicht** über
   `ruleFor` (bekommt `f FileImports` bereits) bis `targetLayer`.
3. **Sprach-bewusste `targetLayer`** (`rules.go`): normalisiert den Import gemäß der `resolution` der
   **Import-Sprache** — `package_base`-Präfix strippen, `.`→`/`, wurzel-relativ — **dann** Glob-Präfix-Match
   gegen die (verzeichnisbasierten) `layers`-Globs. Default unverändert.
   **Grenze ([AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)):**
   greift nur, wenn der **Paket-/Import-Baum den Verzeichnis-Baum spiegelt**
   (Kotlin/Java-Konvention: Verzeichnisse folgen Paketen). Wo Paket ≠ Verzeichnis, löst der Import nicht
   auf → kein Ziel-Layer → keine schicht-basierte Regel (ehrlich ausgewiesen, nicht still).
4. **Tests** (dotted-Modus über **Kotlin/Java** — beide vorhanden **und** gepunktet, decken
   „fester-Wurzel-dotted" exakt ab; **Python ist kein Backend**, slice-017 wiese es mit Exit 2 ab):
   Kotlin/Java-Paket→Layer via `package_base` (Paket==Verzeichnis-Layout); C++ `src`-Root; **Mono-Repo
   Go+Kotlin** (eine Go- **und** eine Kotlin-Datei in *einem* Repo, je eigener Modus); Default-Rückwärtskompat
   (Go/C++ ohne `resolution`); `go: {}` == weggelassen.

## 4. Design-Entscheidungen (Entwurf)

- **`resolution` ist eine Map pro Sprache** (nicht global) — ein Mono-Repo (Go + Kotlin/C++) braucht
  mehrere Modi gleichzeitig. Analog zum `languages`-Map-Muster. **Über [ADR-0014](../../adr/0014-resolution-roots.md) hinaus → Folge-ADR.**
- **Sprach-Threading (der Kern):** die Auflösung eines Imports nutzt den Modus der **Quelldatei-Sprache**;
  deshalb trägt `FileImports` die `Language`, durchgereicht bis `targetLayer`. **Über [ADR-0014](../../adr/0014-resolution-roots.md) hinaus → Folge-ADR.** Cross-Sprach-Importe (selten) → §5.
- **`mode`-Diskriminator = estate-tauglich:** die drei Auflösungs-Modi des Bestands leben unter *einem*
  Schema. slice-015 implementiert `path` + `fixed-root` (deckt Go/C++/JVM/Python); **relativ-zum-File**
  (TS, C++-`"…"`) und **Namespace-Index** (C#) kommen je per Folge-ADR **additiv** dazu — dieselbe
  Sprach-Map, dasselbe Threading, nur ein neuer `mode`-Wert; **kein Re-Architecting**. Bis dahin brechen
  die reservierten Modi mit Exit 2 (§3.1). *(Anderes Signal: relativ braucht den importierenden
  Dateipfad, namespace einen Namespace→Datei-Index — deshalb eigene ADRs, nicht dieser Slice.)*
- **Paket-spiegelt-Verzeichnis** ist die ehrliche Auflösungs-Grenze (§3.3), keine offene Frage.
- **Default bleibt** Import-als-Pfad ohne `resolution` — kein Bruch, Dogfooding 0 ([ADR-0014](../../adr/0014-resolution-roots.md) Fitness Function).

## 5. Abnahme-Gate (blockierend — **vor** der DoD zu lösen)

- **Folge-ADR erforderlich** (B1): [ADR-0014](../../adr/0014-resolution-roots.md) ist immutable; die
  Sprach-Map + das Threading dürfen **nicht** als Historiennotiz dort, sondern nur als erweiternder
  Folge-ADR (`Supersedes: —`) fallen. §3.0.
- **x-wal-Grenze** (B3): das reale x-wal-Layout gegen die „Paket==Verzeichnis"-Grenze (§3.3) prüfen —
  spiegeln x-wals Verzeichnisse die `driving`/`driven`-Pakete? Falls **nein**, deckt der fester-Wurzel-Modus
  x-wal **nicht** (eigener Mechanismus, z. B. Paket→Verzeichnis-Map — dann eigener Slice/ADR). Der JVM-Test
  bleibt bewusst innerhalb der Grenze (Paket==Verzeichnis), sonst testet er ein künstliches Layout.
- **Cross-Sprach-Importe**: Default (unaufgelöst → keine Regel) als ehrliche Grenze bestätigen.

## 6. Definition of Done (erst nach §5-Gate)

- [ ] **Folge-ADR** Accepted + Index; §5-Gate gelöst (x-wal-Grenze empirisch geklärt).
- [ ] Lastenheft + Spezifikation (nächste Minor): `resolution`-Map-Schema, Symbol→Layer sprach-bewusst
  **inkl. Paket==Verzeichnis-Grenze**, je Historie-Zeile.
- [ ] Code: `resolution`-Decode; `FileImports.Language` + Threading; `targetLayer` sprach-bewusst; Default unverändert.
- [ ] Tests: Kotlin/Java-dotted→Layer; C++-`src`-Root; **Mono-Repo Go+Kotlin**; `go: {}`==weggelassen;
  Rückwärtskompat; Dogfooding 0.
- [ ] `make gates` + `make ci` grün; Multi-Linsen-Review; Merge auf Wort.
