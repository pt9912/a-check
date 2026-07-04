# ADR-0020 — Fail-closed-Guard gegen mehrdeutige Mehr-Wurzel-Auflösung (Phantom-Kandidaten)

- **Status:** Proposed
- **Datum:** 2026-07-04
- **Autor:** pt9912
- **Bezug:** [AC-FA-CONF-001](../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml) (fail-closed-Validierung), [AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze) (ein stilles Falsch-Negativ ist der teure Vertragsbruch) — betrifft das Auflösungs-Modell [ADR-0016](0016-resolution-sprach-parametrisch.md)/[ADR-0014](0014-resolution-roots.md).
- **Schärft:** [SPEC-CONF-001](../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema).
- **Supersedes:** —

## Kontext

Ein belief-agent-Bericht (2026-07-04, gegen `v0.9.0`) zeigt ein **stilles
Falsch-Negativ** in Kotlin-Multiplatform: reiner Kern in `commonMain`, Adapter in
`jvmMain`, beide teilen `package_base: myapp`. Bei flachen Source-Set-Globs
(`core: [src/commonMain/**]`, `adapters: [src/jvmMain/**]`) und
`roots: [src/commonMain/kotlin/myapp, src/jvmMain/kotlin/myapp]` wird die illegale
Kante `core → adapter` **nicht** gefangen (`0 Befunde, Exit 0`); mit paket-tiefen
Globs schon.

**Verifizierter Mechanismus** (reproduziert + Gegentest, nicht die Bericht-Hypothese):
`resolveImport` bildet im `fixed-root`-Modus aus `import myapp.adapters.Foo` den
paket-relativen Pfad `adapters/Foo` und prependet **jeden** Root — **datei-mengen-blind**,
ohne zu prüfen, ob die Zieldatei in diesem Root liegt:

1. `src/commonMain/kotlin/myapp/adapters/Foo` ← **Phantom** (existiert nicht)
2. `src/jvmMain/kotlin/myapp/adapters/Foo` ← real

`targetLayer` matcht alle Kandidaten gegen alle Layer-Globs und nimmt den **längsten
literalen Glob-Präfix** ([ADR-0010](0010-layer-relativer-adapterseg-laengster-praefix.md)/[ADR-0013](0013-layerof-laengster-praefix.md)).
Bei flachen Globs matcht das Phantom `core: src/commonMain/**` (Präfix `src/commonMain`),
der reale Kandidat `adapters: src/jvmMain/**` (Präfix `src/jvmMain`). `src/commonMain`
ist **länger** ⇒ das Phantom gewinnt ⇒ der Adapter-Import löst auf `core` auf ⇒
`core → core` ⇒ sauber ⇒ Falsch-Negativ. Weil die Auswahl über Präfix-**Länge** läuft,
ist die `roots`-**Reihenfolge irrelevant** (beide Ordnungen ⇒ 0). Der Gegentest —
flache Globs, aber Adapter-Ordner länger benannt — kippt zu korrektem `core-impurity`
und belegt den Längen-Mechanismus.

Bei **einem** Root existiert kein Phantom (Go-Ein-Baum, Kotlin-Ein-Baum sauber); der
Fehler ist an **≥ 2 Roots mit geteiltem `package_base` UND Layer-Globs flacher als die
paket-diskriminierende Tiefe** gebunden.

## Optionen

| Weg | Idee | Bewertung |
|---|---|---|
| **A — fail-closed-Guard (Config-Validierung)** | Die Phantom-fähige Config-Form beim Laden erkennen und mit Exit 2 ablehnen (Verweis auf paket-tiefe Globs). Keine Änderung der Auflösungs-Semantik. | **Gewählt (Stufe 1).** Tilgt das *stille* Grün sofort, konsistent mit der belegten Linie (unbekannte Sprache → Exit 2 [slice-017], leerer `tech.adapter` → Exit 2 [slice-023]); falsch-positiv-frei gegen die reale Flotte (Scan 2026-07-04: keine Config mit ≥ 2 Roots). |
| **B — datei-mengen-bewusste Auflösung** | Phantom-Kandidaten gegen die real gescannte Dateimenge filtern (nur existierende Pfade überleben). | **Aufgeschoben (Stufe 2, gated — slice-027).** Macht die flache Config *korrekt* statt nur laut, aber drei ungelöste Fragen: `expect`/`actual` (derselbe FQN legitim in zwei Source-Sets), endungsloses Matching (`adapters/Foo` ↔ `Foo.kt`), Determinismus der Filterung. Eigener ADR bei Lande-Trigger. |
| **C — nur Warnung (stderr), Exit 0/1** | Die Config bestehen lassen, aber warnen. | Verworfen: a-check hat keinen von Befunden getrennten Warn-Kanal (eigenes Feature); und eine Warnung bei einem *Gate* lässt das falsch-grün-Ergebnis bestehen — das widerspricht dem fail-closed-Ethos. |
| **D — reine Doku (KMP-Rezept)** | Nur Benutzerhandbuch, kein Signal. | Verworfen: macht das stille Grün nicht sichtbar; wer die Falle nicht kennt, tappt hinein. |

## Entscheidung

**Weg A (Stufe 1).** Der Config-Adapter lehnt beim Laden fail-closed (Exit 2) ab, wenn:

- eine Sprache `mode: fixed-root` mit **≥ 2** `roots` hat, **und**
- **≥ 2 verschiedene** Roots je **vollständig in einer Schicht enthalten** sind, **und**
- diese Schichten **verschieden** sind.

**„Root R vollständig in Schicht L enthalten"** := ein `layers`-Glob-Präfix `P` von L
ist **Vorfahr-oder-gleich** von R — `segIndex(R, P) == 0` **und** `len(P) ≤ len(R)`.
Dann matcht **jeder** Kandidat `R/…` die Schicht L allein am Wurzel-Präfix, unabhängig
vom Paketpfad. Liegen zwei verschiedene Roots so in zwei verschiedenen Schichten, kann
ein Import je Schicht ein Phantom erzeugen und die Zuordnung wird vom längsten Präfix
statt vom Symbol entschieden — genau die Falsch-Negativ-Bedingung.

Die Meldung nennt die zwei Roots und ihre Schichten und verweist auf das Rezept
(paket-spezifische Globs, tiefer als die Roots — dann diskriminiert der Paketpfad).

**Abgrenzung:** Der Guard prüft nur `fixed-root` mit ≥ 2 Roots. Ein Root, ein
`path`-/`relative`-Modus, oder ≥ 2 Roots mit paket-tiefen Globs (kein Root in einer
Schicht enthalten) laden unverändert. Damit ist die Änderung eine **Verschärfung** einer
bisher stillen Lücke, kein Verhaltenswechsel für bestehende gültige Configs.

## Konsequenzen

- **Die reale Flotte ist unberührt:** Scan 2026-07-04 — a-check/d-check (kein
  `resolution`), b-cad (1 Root) fallen nicht unter den Guard; kein Bestands-Test
  konstruiert ein Modell mit ≥ 2 Roots in ≥ 2 Schichten.
- **Der KMP-Fall wird laut:** die belief-agent-Config bricht mit Exit 2 + Rezept-Hinweis,
  statt still grün zu geben. Der Konsument nutzt paket-tiefe Globs (sofort korrekt) und
  gewinnt später mit Stufe 2 die flache Form zurück.
- [SPEC-CONF-001](../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema)
  nennt den Guard in der fail-closed-Liste.
- **Grenze bleibt ehrlich** ([AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)):
  der Guard verhindert eine *bekannte* Fehlklassifikations-Form; er macht die
  Mehr-Wurzel-Auflösung nicht vollständig (das ist Stufe 2).

## Fitness Function

- `make test`: Config X (KMP flach, 2 Roots in 2 Schichten) → Exit 2; Config Y (dieselben
  Roots, paket-tiefe Globs) → lädt, `core → adapter` = `core-impurity` (Exit 1);
  b-cad-artiger 1-Root → lädt; Grenzform (ein Root **exakt** = einem Glob-Präfix) →
  deterministisch klassifiziert.
- `make arch-check` (Dogfooding): unverändert 0 (a-checks Config hat keinen
  `resolution`-Block).

## Re-Evaluierungs-Trigger

- **Stufe 2 (slice-027):** sobald datei-mengen-bewusste Auflösung gebaut wird, wird der
  Guard von „ablehnen" ggf. zu „auflösen" — der Guard bleibt aber als schnelle,
  semantik-freie Absicherung sinnvoll, bis Stufe 2 die `expect`/`actual`-Mehrdeutigkeit
  entschieden hat.
- **Legitime ≥ 2-Root-in-≥ 2-Schichten-Config:** falls je ein Konsument eine solche Form
  *korrekt* meint (heute: keine bekannt), Re-Eval des Prädikats.

## Geschichte

| Datum | Ereignis |
|---|---|
| 2026-07-04 | Proposed — Entwurf mit [slice-026](../planning/open/slice-026-kmp-mehr-root-phantom.md); Mechanismus reproduziert + per Gegentest belegt; Falsch-Positiv-Scan der realen Flotte (keine ≥ 2-Root-Config). |
