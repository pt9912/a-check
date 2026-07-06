# ADR-0022 — Datei-mengen-bewusste Mehr-Wurzel-Auflösung (ersetzt den Phantom-Guard)

- **Status:** Superseded by ADR-0023
- **Datum:** 2026-07-05
- **Autor:** pt9912
- **Bezug:** [AC-FA-CONF-001](../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml) (Mehr-Wurzel-Auflösung, fail-closed), [AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze) (ein stilles Falsch-Negativ ist der teure Vertragsbruch); erweitert das Auflösungs-Modell [ADR-0016](0016-resolution-sprach-parametrisch.md) / [ADR-0014](0014-resolution-roots.md) — der `fixed-root`-Root-Prepend wird datei-mengen-bewusst; beide bleiben immutabel.
- **Schärft:** [SPEC-CONF-001](../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema).
- **Supersedes:** [ADR-0020](0020-mehr-wurzel-phantom-guard.md).

## Kontext

[ADR-0020](0020-mehr-wurzel-phantom-guard.md) fing ein stilles Falsch-Negativ in
Kotlin-Multiplatform per **statischem Guard**: `mode: fixed-root` mit ≥ 2 `roots`, deren
Roots verschiedene Schichten erzwingen, wurde beim Laden mit Exit 2 abgelehnt (Stufe 1,
schnell + laut). Ein realer Konsumenten-Fall (KMP/Gradle-Multi-Modul mit geteiltem
`package_base` und **disjunkten** Paket-Sub-Namespaces je Modul) zeigt die Kehrseite: für
diese **legitime** Struktur existiert **keine** korrekte Config — der Guard rejectet die
disjunkte Multi-Root-Config, und jede Umgehung (ein Glob-Root, `mode: path`, flache/tiefe
Globs) ist **still falsch-grün**: die verbotene `domain → application`-Kante wird nicht
gemeldet ([AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)).

Der Guard war ausdrücklich als **Stufe 1** deklariert; ADR-0020 §Re-Evaluierungs-Trigger
nennt die **datei-mengen-bewusste Auflösung** als Stufe 2. Schlüssel: `Evaluate` hat die
real gescannte Datei-Menge (`FileImports.Path`, repo-relativ, Slash-normalisiert) bereits —
sie lebt im selben Namensraum wie die Layer-Globs und die `fixed-root`-Kandidaten. Kein
zusätzliches Datei-I/O.

## Optionen

| Weg | Idee | Bewertung |
|---|---|---|
| **A — datei-mengen-bewusste Auflösung** | Den `fixed-root`-Kandidaten je Root gegen die reale Datei-Menge filtern (nur existierende Ziele überleben); Schicht am realen Kandidaten-Pfad; echte Mehrdeutigkeit fail-closed nach dem Scan. | **Gewählt.** Fängt den echten Bug **und** akzeptiert die legitime disjunkte Config — strikt präziser als der Guard, keine Schwächung der Durchsetzung. Löst zugleich die von ADR-0020 offen gelassene `expect`/`actual`-Frage (distinct-layer). |
| **B — Guard behalten/reduziert** | Statischen Config-Guard belassen. | Verworfen: der Guard prüft die Config-**Form** (Roots erzwingen verschiedene Schichten) — genau die legitime disjunkte Config; ein reduzierter Guard bräche sie erneut. |
| **C — literal: gleicher FQN in ≥ 2 Roots → Exit 2** | Jeden real doppelt auflösbaren FQN ablehnen. | Verworfen: bräche legitimes same-layer-`expect`/`actual` (gleicher FQN in zwei Source-Sets **derselben** Schicht) mit fälschlichem Exit 2. |

## Entscheidung

**Weg A.** Bei `mode: fixed-root` mit **≥ 2** `roots` wird der interne FQN **datei-mengen-
bewusst** aufgelöst:

1. **Realitäts-Filter:** ein Kandidat `root/…` überlebt nur, wenn er einer **real gescannten
   Datei** entspricht — endungs-agnostisch, package==directory: ein Wildcard-/Paket-Import
   (Symbol mit Trailing-Dot, `a.b.*` → `…/b/`) trifft das Paket-**Verzeichnis**; ein Symbol
   trifft seine Datei (Endung gestrippt, `…/B` ↔ `…/B.kt`) **oder**, bei Kotlins Klasse≠Datei,
   sein Paket-Verzeichnis. Disjunkte Sub-Namespaces matchen so in **höchstens einem** Root.
2. **Schicht** am Pfad des **realen** Kandidaten
   ([SPEC-RULE-001](../../../spec/spezifikation.md#spec-rule-001--regel-auswertung)-Glob-Auswahl),
   nicht am Wurzel-Präfix. Ein Kandidat ohne reale Datei (Phantom) trifft keine Schicht und
   bleibt **extern** (Tech-Regeln greifen unverändert am Roh-Symbol).
3. **Fail-closed (Exit 2, nach dem Scan):** löst derselbe FQN real unter ≥ 2 Roots in
   **verschiedene** Schichten auf, ist das echte Mehrdeutigkeit — `a-check` bricht mit Exit 2
   ab (ein FQN muss in höchstens eine Schicht auflösen). `Evaluate` gibt dafür einen Fehler
   zurück, den die CLI wie einen Lade-/Extraktions-Fehler auf Exit 2 abbildet (Findings
   unterdrückt). **Gleiche** Schicht (`expect`/`actual`) löst sauber — die **distinct-layer**-
   Verfeinerung der von ADR-0020 offen gelassenen Residual-Frage.

Der statische Guard (`PhantomRootConflictIn`/`rootForcedLayer` + Config-Ladezeit-Reject)
**entfällt**.

## Konsequenzen

- **Die legitime disjunkte Multi-Modul-Config lädt und löst korrekt** — die verbotene
  `domain → application`-Kante wird gemeldet (Exit 1) statt still übersehen.
- **Rückwärtskompatibel:** nur der `fixed-root`-Pfad mit **≥ 2 Roots** ändert Verhalten;
  `path`/`relative`/1-Root/`package_base`-ohne-Roots/kein-`resolution` werden identisch
  durchgereicht (der Realitäts-Filter greift nur dort).
- **Exit 2 kann jetzt nach dem Scan entstehen** (bisher nur zur Ladezeit): der Exit-Code-
  Vertrag ([AC-FA-CLI-001](../../../spec/lastenheft.md#ac-fa-cli-001--aufruf-scan-wurzel-und-exit-codes))
  bleibt (2 = Nutzungs-/Konfigurationsfehler), die echte Mehrdeutigkeit ist ein
  Config-Fehler gegen den realen Baum.
- **Ehrliche Grenzen** ([AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)):
  ein Import, dessen Paket-Verzeichnis unter **keiner** Root real ist (fehlkonfigurierte/
  nicht gelaufene Source-Set), bleibt still extern; verschachtelte-Klassen-Importe (das
  letzte Segment ist eine Klasse in einer Klasse) lösen ein Verzeichnis zu tief auf und
  bleiben extern; datei-tiefe Layer-Globs sind eine dokumentierte Grenze. Per-Root
  `package_base` und der Namespace-Index-Modus bleiben additive Folgeschritte.

## Fitness Function

- `make test`: sauberes disjunktes Multi-Modul → Exit 0; `domain → application` → 1 Befund,
  Exit 1 (vor der Änderung: 0); Tech-Leak am Roh-Symbol unverändert scharf; disjunkte
  Multi-Root-Config **lädt** (kein Ladezeit-Reject); gleicher FQN real in ≥ 2 Roots +
  verschiedene Schichten → Exit 2 (stderr-Meldung, kein stdout-Befund); same-layer
  `expect`/`actual` → sauber; `--print-config` dokumentiert die Multi-Modul-Resolution.
  Unit: `denotes` (Datei-/Verzeichnis-/Wildcard-/Phantom-Fälle), distinct-layer-Ambiguität.
- `make arch-check` (Dogfooding): unverändert 0 (a-check hat keinen `resolution`-Block).

## Re-Evaluierungs-Trigger

- **Per-Root `package_base`** (`roots: [{path, package_base}]`): additive Erweiterung, sobald
  ein Konsument disjunkte Module mit **verschiedenen** Paket-Basen führt.
- **Verschachtelte-Klassen-Importe / datei-tiefe Globs:** falls ein realer Konsument die
  heute externe Grenze braucht, Re-Eval des `denotes`-Prädikats.
- **Legitimes same-FQN-verschiedene-Schicht:** falls je ein Konsument dieselbe voll
  qualifizierte Kennung in zwei Schichten *korrekt* meint (heute: keine bekannt), Re-Eval des
  distinct-layer-Fail-closed.

## Geschichte

| Datum | Ereignis |
|---|---|
| 2026-07-05 | Proposed — Entwurf; belief-agent-KMP-Evidenz (jede `resolution`-Variante Reject **oder** still falsch-grün); datei-mengen-bewusste Auflösung reproduziert, AC1–AC6 als Tests grün; Multi-Linsen-Review des Entwurfs eingearbeitet. |
| 2026-07-05 | Umsetzungs-Review (3 Linsen) eingearbeitet: **A1** (spurioses Exit 2 durch ein Phantom, dessen Elternverzeichnis nur zufällig existiert) per `strength`-Stufen gefixt; Test-Härtung (Determinismus-Zeuge, same-layer-/Wildcard-E2E). |
| 2026-07-05 | Proposed → Accepted (Sign-off Auftraggeber per Merge-Wort). Ab jetzt immutable; Ablösung nur via Folge-ADR mit `Supersedes`. |
