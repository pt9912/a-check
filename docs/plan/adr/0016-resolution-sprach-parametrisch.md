# ADR-0016 — Resolution sprach-parametrisch: Sprach-Map + Threading + `mode`-Diskriminator

- **Status:** Accepted
- **Datum:** 2026-07-01
- **Autor:** pt9912
- **Bezug:** [AC-FA-CONF-001](../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml), [AC-FA-EXTRACT-001](../../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion) — **Erweiterung** von [ADR-0014](0014-resolution-roots.md) (Resolution-Roots), nach dem Muster, mit dem ADR-0014 selbst [ADR-0002](0002-text-heuristische-extraktion.md) erweiterte.
- **Schärft:** [SPEC-CONF-001](../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema) + [SPEC-RULE-001](../../../spec/spezifikation.md#spec-rule-001--regel-auswertung) + [SPEC-EXTRACT-001](../../../spec/spezifikation.md#spec-extract-001--import-extraktion).
- **Supersedes:** —

## Kontext

[ADR-0014](0014-resolution-roots.md) (Accepted) skizzierte den `resolution`-Block **flach und global**:
`resolution: {roots: ["src"], package_base: "com.xwal"}`. Das war für den **einen** JVM-Auslöser
gedacht. Der reale **Polyglot-Bestand** (Repos in Go, C++, C#, Python, TypeScript — nicht nur x-wal)
und **Mono-Repos** (mehrere Sprachen in *einem* Repo) brechen diese Annahme in drei Punkten, die
ADR-0014 nicht ausschreibt:

1. **Pro Sprache verschieden:** Go löst über den Modulpfad auf, C++ über den Include-Root, JVM/Python
   über gepunktete Pakete, TypeScript relativ-zum-File, C# über Namespaces. Ein globaler Block kann
   ein Mono-Repo (Go + Kotlin) nicht bedienen.
2. **Import kennt seine Sprache nicht:** `targetLayer` (`rules.go:231`) bekommt heute nur das
   Import-Symbol. Um den richtigen Modus zu wählen, muss es die **Quelldatei-Sprache** wissen —
   `FileImports` (`model.go:17`) trägt sie aber nicht.
3. **Mehrere Modi:** fester-Wurzel, relativ, namespace sind verschiedene Auflösungs-**Signale**, kein
   Parameter *eines* Modus.

## Optionen

| Weg | Idee | Bewertung |
|---|---|---|
| **A — Sprach-Map + `mode` + Threading** | `resolution` als Map Sprache→Config mit `mode`-Diskriminator; `FileImports.Language` bis `targetLayer` durchgereicht. | **Gewählt.** Estate-weit (ein Schema für alle Sprachen/Modi), Mono-Repo-tauglich, additiv erweiterbar; bleibt text-heuristisch (ADR-0002/0014-treu). |
| **B — flacher globaler Block** ([ADR-0014](0014-resolution-roots.md)) | ein `{roots, package_base}` fürs ganze Repo. | Verworfen: kann Mono-Repos nicht bedienen; kein `mode` für relativ/namespace. |
| **C — Sprache aus dem Import raten** | Modus am Symbol erkennen (`.` → dotted, `/` → Pfad). | Verworfen: brüchig (Go-Modulpfade enthalten `.`), rät statt zu wissen; das Threading ist billig und exakt. |

## Entscheidung

**Weg A**, als **Erweiterung** von [ADR-0014](0014-resolution-roots.md) (dessen `roots`/`package_base`
gelten unverändert *innerhalb* einer Sprach-Config):

1. **`resolution` = Map Sprache → Config.** Fehlt eine Sprache (oder der ganze Block) → heutiges
   Verhalten (Import-als-Pfad), rückwärtskompatibel.
2. **`mode`-Diskriminator** je Sprache:
   - `path` — Import *ist* der wurzel-relative Pfad (Go; Default, `== weggelassen`).
   - `fixed-root` — `roots` (Wurzeln) vorangestellt; **bei gesetztem `package_base`** (gepunktete
     Sprache) zusätzlich Präfix-Strip + `.`→`/` (eine Pfad-Sprache wie C++ behält ihre `.`-Endungen).
     Deckt C++-`src`-Root, JVM, Python.
   - `relative`, `namespace` — **reserviert**; bis zum jeweiligen Folge-ADR → **Exit 2** (kein stiller
     No-Op, konsistent mit slice-017/[AC-FA-EXTRACT-001](../../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion)).
3. **Sprach-Threading:** `FileImports.Language` (aus dem Extraktions-Backend) wird über `ruleFor` bis
   `targetLayer` durchgereicht; die Auflösung nutzt den `mode` der **Import-Sprache**.
4. **Grenze ([AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)):**
   `fixed-root` greift nur, wenn der **Paket-/Import-Baum den Verzeichnis-Baum spiegelt**; wo Paket ≠
   Verzeichnis (z. B. flach unter einer Gradle-Modulgrenze), löst der Import nicht auf → kein Ziel-Layer
   → keine schicht-basierte Regel (ausgewiesen, nicht still).

## Konsequenzen

- [ADR-0014](0014-resolution-roots.md) **bleibt gültig** — ADR-0016 verallgemeinert nur den Block
  (global → sprach-parametrisch) und ergänzt das Threading; kein Supersede.
- **Schema** ([AC-FA-CONF-001](../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml)/[SPEC-CONF-001](../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema)):
  `resolution`-Map mit `mode`; strict-decode, unbekannter/reservierter `mode` → Exit 2.
- **Estate-weit erweiterbar:** `relative` (TS, C++-`"…"`) und `namespace` (C#) kommen **additiv** je
  Folge-ADR — dieselbe Map, dasselbe Threading, nur ein neuer `mode`-Wert; kein Re-Architecting.
- Build-Manifest bleibt **optionaler** Resolution-Hint (ADR-0014), nie Regel-Backend.

## Fitness Function

- `make test`: Kotlin/Java-Paket→Layer via `fixed-root`/`package_base` (Paket==Verzeichnis); C++
  `src`-Root; Mono-Repo Go+Kotlin (je eigener `mode`); `path`==weggelassen; reservierter `mode` → Exit 2.
- `make arch-check` (Dogfooding, [AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)):
  unverändert 0 (a-check deklariert kein `resolution` → `path`).

## Re-Evaluierungs-Trigger

- **`relative`-Modus** (TypeScript/C++-`"…"`): eigener ADR, sobald ein Pilot feuert (braucht den
  importierenden Dateipfad, nicht nur die Wurzel).
- **`namespace`-Modus** (C#): eigener ADR (Namespace→Datei-Index; am Rand von ADR-0002).
- **Paket≠Verzeichnis** (x-wal, falls sein Layout die Pakete nicht spiegelt): Paket→Verzeichnis-Map als
  eigenes Inkrement.

## Geschichte

| Datum | Ereignis |
|---|---|
| 2026-07-01 | Proposed → Accepted (Sign-off Auftraggeber: Weg A — Sprach-Map + `mode` + Threading, estate-weit; erweitert ADR-0014, `Supersedes: —`). Umsetzung [slice-015](../planning/done/slice-015-resolution-roots.md). Ab jetzt immutable; Ablösung nur via Folge-ADR mit `Supersedes`. |
