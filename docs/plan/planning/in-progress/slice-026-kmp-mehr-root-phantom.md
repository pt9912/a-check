# slice-026 — Mehr-Root-Phantom: fail-closed-Guard (+ datei-mengen-bewusste Auflösung, gated)

**Status:** in-progress (abgenommen + **Stufe 1 umgesetzt** 2026-07-04 — fail-closed-Guard, spec-first; Multi-Linsen-Review + Fixes eingearbeitet, [Review-Synthese](../../../reviews/2026-07-04-slice-026-phantom-guard.md). Stufe 2 gated als slice-027).
**Bezug:** schärft [AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)
(ehrliche Heuristik-Grenze — ein stilles Falsch-Negativ ist der teure Bruch dieses
Vertrags) + [AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml)
(Konfigurations-Validierung, fail-closed); Auflösungs-Modell
[ADR-0016](../../adr/0016-resolution-sprach-parametrisch.md) (Re-Eval),
Wurzel-Auflösung [ADR-0014](../../adr/0014-resolution-roots.md);
Determinismus [AC-QA-01](../../../../spec/lastenheft.md#ac-qa-01--determinismus).
[Roadmap](roadmap.md). **Evidenz:** Fehlerbericht belief-agent gegen
`v0.9.0` (2026-07-04), hier reproduziert (§2, §6).

## 1. Auslöser — stilles Falsch-Negativ bei KMP-Mehr-Source-Set

Ein Kotlin-Multiplatform-Repo (reiner Kern in `commonMain`, Adapter in `jvmMain`,
beide teilen `package_base: myapp`) meldet: die **illegale Kante `core → adapter`
wird mit flachen Source-Set-`layers`-Globs NICHT gefangen** (`0 Befunde, Exit 0`),
mit tiefen paket-spezifischen Globs schon (`1 Befund core-impurity`). Beide
`roots`-Reihenfolgen ergeben 0. Reproduziert gegen `a-check:dev` (identisch mit dem
v0.9.0-Digest `@sha256:0378211f…`), Config aus dem Bericht.

Das ist die **teuerste Fehlerklasse dieses Projekts**: ein Gate, das still grün gibt,
wo es rot geben müsste — der direkte Bruch des „Invarianten als Gate statt
Review-Meinung"-Ethos ([AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)).

## 2. Mechanismus (verifiziert, nicht vermutet)

Die Bericht-Hypothese („Kandidat wird dem falschen Source-Set zugeordnet") stimmt in
der Richtung, aber der genaue Weg ist ein anderer. `resolveImport`
(`internal/hexagon/core/rules.go`, `mode: fixed-root`) bildet aus
`import myapp.adapters.Foo`:

1. Präfix `myapp.` streichen, Punkte → Slashes ⇒ paket-relativ `adapters/Foo`.
2. **Pro Root ein Kandidat** — ohne Prüfung, ob die Zieldatei in diesem Root liegt:
   - `src/commonMain/kotlin/myapp/adapters/Foo` ← **Phantom** (existiert nicht)
   - `src/jvmMain/kotlin/myapp/adapters/Foo` ← real

`targetLayer` matcht **alle** Kandidaten gegen **alle** Layer-Globs und nimmt den
**längsten literalen Glob-Präfix** über die Gesamtmenge
([ADR-0010](../../adr/0010-layer-relativer-adapterseg-laengster-praefix.md)/[ADR-0013](../../adr/0013-layerof-laengster-praefix.md)). Bei
flachen Globs (Config X):

- Phantom matcht `core: src/commonMain/**` (Präfix `src/commonMain`)
- Real matcht `adapters: src/jvmMain/**` (Präfix `src/jvmMain`)

`src/commonMain` ist **länger** als `src/jvmMain` ⇒ das Phantom gewinnt ⇒ der
Adapter-Import löst auf `core` auf ⇒ `core → core` ⇒ sauber ⇒ Falsch-Negativ. Weil
die Auswahl über Präfix-**Länge** läuft (strikt-länger ersetzt), ist die
`roots`-**Reihenfolge irrelevant** — beide Ordnungen ergeben 0.

**Entscheidender Gegentest** (belegt den Längen-Mechanismus, schließt „Reihenfolge"
und „Datei-Existenz" aus):

| Konfiguration | Erwartet | Beobachtet |
|---|---|---|
| Config X (flach, `commonMain`-Präfix länger) | 1 `core-impurity` | **0** ✗ (Falsch-Negativ) |
| Config Y (tiefe Globs) | 1 `core-impurity` | 1 ✓ |
| Gegentest (flach, **Adapter**-Ordner länger benannt) | 1 `core-impurity` | **1** ✓ (Bug kippt) |

Config Y wirkt, weil das paket-diskriminierende Segment (`core`/`adapters`) **im**
Glob-Präfix steckt: das Phantom `…/commonMain/…/adapters/Foo` matcht den Core-Glob
`…/commonMain/…/core` nicht mehr und wird verworfen.

**Kernursache:** Mehr-Root-`fixed-root`-Auflösung ist **datei-mengen-blind** — ein
kartesisches Produkt Root × Paketpfad, von dem bei N Roots N−1 Kandidaten Phantome
sind. Bei einem Root gibt es kein Phantom (⇒ Go-Ein-Baum und Kotlin-Ein-Baum sauber,
Kontroll-Positive des Berichts). Gebunden an: **mehrere Roots mit geteiltem
`package_base` UND Layer-Globs flacher als die paket-diskriminierende Tiefe.** KMP ist
der häufigste Auslöser dieser Config-Form, nicht die einzige.

## 3. Scope slice-026 (Stufe 1) — fail-closed-Guard gegen Phantom-fähige Configs

**Ziel:** das *stille* Falsch-Negativ sofort tilgen — laut statt blind-grün, ohne die
Auflösungs-Semantik anzufassen. Konsistent mit der belegten Linie (unbekannte Sprache
→ Exit 2 [slice-017], leerer `tech.adapter` → Exit 2 [slice-023]).

**Detektierbare Bedingung** (bei `mode: fixed-root`, ≥ 2 `roots`): Ein Root R heißt
*in einer Schicht enthalten*, wenn ein Layer-Glob-Präfix P Vorfahr-oder-gleich von R
ist (`segIndex(R, P) == 0`, `len(P) ≤ len(R)`) — dann matcht **jeder** Kandidat aus R
diese Schicht allein am Root-Präfix, unabhängig vom Paketpfad. Liegen **≥ 2 Roots in
≥ 2 verschiedenen Schichten**, kann ein Import je Schicht ein Phantom erzeugen und die
Schicht-Zuordnung wird beliebig (längster-Präfix). ⇒ **Diese Config ist
unterspezifiziert.** (Config X trifft zu: root1 in `core`, root2 in `adapters`.
Config Y trifft nicht zu: die Layer-Globs sind *tiefer* als die Roots, kein Root ist in
einer Schicht enthalten.)

**Reaktion:** Exit 2 mit erklärender Meldung, die auf das tiefe-Globs-Rezept verweist
(empfohlen — fail-closed, Projekt-Ethos). Weichere Alternative aus dem Bericht: Warnung
auf `stderr` bei Exit 0/1 — verlangt aber einen Warn-Kanal, den a-check heute nicht hat
(eigene Design-Entscheidung, §5). Reine Doku ohne Signal wäre unzureichend, weil sie das
stille Grün nicht sichtbar macht.

**Anforderungs-Skizze** (Schärfung von
[AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml);
ID-/Versions-Bump erst im Lastenheft-CR selbst —
[`harness/conventions.md`](../../../../harness/conventions.md#anforderungs-anlege-prozess)):

- **Happy:** Given `fixed-root`, 1 Root (oder ≥ 2 Roots, aber kein Root in einer
  Schicht enthalten — Config Y), when a-check lädt, then Exit 0/1 wie bisher
  (byte-identisch — kein bestehender Konsument bricht).
- **Boundary:** Given ≥ 2 Roots, davon genau die Grenzform (ein Root exakt = einem
  Layer-Glob-Präfix), when a-check lädt, then greift die Bedingung deterministisch
  (Prädikat-Kante spezifizieren).
- **Negative:** Given Config X (≥ 2 Roots in ≥ 2 Schichten enthalten), when a-check
  lädt, then Exit 2 mit Meldung „Resolution-Root … liegt vollständig in Schicht … —
  Phantom-Kandidaten über Layer-Grenzen möglich; tiefe paket-spezifische Globs nutzen"
  (Rezept-Verweis).
- **Out-of-Scope:** die Auflösung *korrekt* machen (Stufe 2, §4); Nicht-`fixed-root`-Modi.

## 4. Folge-Slice slice-027 (Stufe 2, **gated**) — datei-mengen-bewusste Auflösung

**Idee:** Mehr-Root-Kandidaten gegen die real gescannte Dateimenge filtern — ein
Kandidat überlebt nur, wenn ihm eine gescannte Datei entspricht (Phantom
`…/commonMain/…/adapters/Foo` fällt, weil keine solche Datei existiert). a-check *hat*
die Dateimenge (`files []FileImports`); heute kennt die Auflösung aber nur Globs, nicht
die Datei-Pfade. Macht die flache KMP-Config *korrekt* statt nur laut.

**Bewusst gated** — drei ungelöste Fragen, die einen eigenen Slice + ADR rechtfertigen:
1. **`expect`/`actual`:** derselbe FQN `myapp.Foo` kann legitim in *beiden* Source-Sets
   liegen (commonMain `expect`, jvmMain `actual`) ⇒ echt mehrdeutige Auflösung. Welche
   Schicht? (Beide prüfen? Strengste gewinnt? Kante gegen beide?)
2. **Endungsloses Matching:** `adapters/Foo` ↔ Datei `Foo.kt`, bzw. Paket-Import →
   Verzeichnis (kein `.kt`) — Stamm-/Präfix-Match, Spez-Entscheidung.
3. **Determinismus** ([AC-QA-01](../../../../spec/lastenheft.md#ac-qa-01--determinismus)):
   die Kandidaten-Filterung muss stabil sortiert und reproduzierbar bleiben.

**Lande-Trigger:** Maintainer-Wort oder zweiter Konsument mit Mehr-Source-Set-Bedarf.
Bis dahin schützt Stufe 1 vor dem stillen Grün.

## 5. Vor der Umsetzung zu klären (ADR-Skizze, Stufe 1)

- **Exit 2 vs. Warn-Kanal:** a-check hat heute keinen `stderr`-Warn-Kanal getrennt von
  Befunden. Exit 2 (fail-closed) braucht keinen neuen Kanal und passt zum Ethos; ein
  Warn-Kanal wäre ein eigenständiges Feature (auch für andere künftige Heuristik-Grenzen
  nützlich, aber Scope-Ausweitung). **Empfehlung: Exit 2.** Maintainer-Entscheid.
- **Prädikat-Präzision:** die „Root in Schicht enthalten"-Kante exakt fassen (Root ==
  Glob-Präfix; Root mit mehreren Layer-Globs; `**`-nur-Globs mit Präfix `""`). Testbar
  gegen die drei §2-Configs plus die Grenzform aus §3-Boundary.
- **Falsch-Positiv-Risiko:** gibt es eine *legitime* ≥2-Root-in-≥2-Schichten-Config mit
  flachen Globs? (Vermutung: nein — wer Roots pro Schicht setzt, soll Globs pro Schicht
  setzen; der Guard erzwingt genau das.) Vor Exit 2 belegen, dass kein bestehender
  Konsument (d-check/b-cad/x-wal) darunterfällt.
- **Determinismus/Reihenfolge** der Meldung bei mehreren verletzenden Roots.

## 6. Evidenz-Fixture (bei Umsetzung)

Das reproduzierte Minimal-Repro als `make test`-Fixture: `commonMain/…/core/Bar.kt`
(importiert `myapp.adapters.Foo`) + `jvmMain/…/adapters/Foo.kt`, drei Configs
(X flach → nach Fix Exit 2; Y tief → 1 Befund; Gegentest flach-Adapter-länger → 1
Befund). Der Gegentest bleibt als Regressions-Anker für den Längen-Mechanismus.

## 7. DoD (bei Ausarbeitung)

Spec-first-Reihenfolge (Lastenheft-CR mit geschärfter
[AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml)
→ ADR für die Detektions-Regel → Spezifikation
[SPEC-CONF-001](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema)
→ Code → Tests); Multi-Linsen-Review vor Merge; `make gates` grün; Benutzerhandbuch:
KMP-Rezept (tiefe paket-spezifische Globs) + Guard-Meldung dokumentiert.
