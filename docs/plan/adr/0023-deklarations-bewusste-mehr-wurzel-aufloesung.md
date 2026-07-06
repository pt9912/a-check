# ADR-0023 — Deklarations-bewusste Mehr-Wurzel-Auflösung für Split-Packages (ersetzt datei-mengen-bewusst)

- **Status:** Accepted
- **Datum:** 2026-07-06
- **Autor:** pt9912
- **Bezug:** [AC-FA-CONF-001](../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml) (Mehr-Wurzel-Auflösung, fail-closed), [AC-FA-EXTRACT-001](../../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion) (Deklarations-Extraktion je Backend), [AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze) (ein stilles Falsch-Negativ ist der teure Vertragsbruch); erweitert das Auflösungs-Modell [ADR-0016](0016-resolution-sprach-parametrisch.md) / [ADR-0014](0014-resolution-roots.md) — der `fixed-root`-Root-Prepend wird deklarations-bewusst; beide bleiben immutabel.
- **Schärft:** [SPEC-CONF-001](../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema), [SPEC-EXTRACT-001](../../../spec/spezifikation.md#spec-extract-001--import-extraktion).
- **Supersedes:** [ADR-0022](0022-datei-mengen-bewusste-mehr-wurzel-aufloesung.md).

## Kontext

[ADR-0022](0022-datei-mengen-bewusste-mehr-wurzel-aufloesung.md) löste das
KMP-disjunkte-Namespace-Falsch-Negativ **datei-mengen-bewusst**: bei `mode: fixed-root` mit
≥ 2 `roots` überlebt ein Kandidat nur, wenn er einer real gescannten Datei entspricht
(endungs-agnostisch, **package==directory**). Das trägt disjunkte Sub-Namespaces, weil das
Paket-Verzeichnis dort in höchstens einem Root existiert.

Ein realer Konsument zeigt die verbleibende Lücke. `d-migrate` ist ein JVM-Kotlin-Hexagon
(kein KMP) mit **konzept-basierten** Packages, die über mehrere Gradle-Module/Schichten
spannen — die Hexagon-Schicht steckt im Modul-**Verzeichnis**, nicht im Package-Namen.
Idiomatisches Kotlin legt **Top-Level-Deklarationen** (Extension-Funktionen, mehrere Klassen je
Datei) in Packages ab, die real über zwei Module verteilt sind (**Split-Package**): dasselbe
Paket liegt unter einem **Port**- *und* einem Driven-**Adapter**-Modul-Root. Für ein
importiertes Symbol, dessen **Datei ≠ Symbolname** ist (`asJdbc` als Extension-Fun in
`JdbcDatabaseConnection.kt`; die Klasse `ChunkColumnSchema` in `ChunkSchema.kt`), erreicht der
package==directory-Filter nur die **Elternpaket-Stufe** — und zwar unter **beiden** Roots →
distinct-layer-Mehrdeutigkeit → **Exit 2**, obwohl das Symbol real **eindeutig in genau einem
Modul deklariert** ist. Der Exit 2 bricht den ganzen Scan ab; die einzige korrekte
Vollrichtungs-Config wird unbrauchbar. Systemisch (~17 solcher Split-Packages), und die
`markers.ignore_symbols`-Umgehung ist Whack-a-Mole über genau die schicht-übergreifenden
Grenz-Symbole, die das Gate prüfen soll.

**Schlüssel:** die fehlende Evidenz ist, *welches Modul das Symbol deklariert.* Der Quelltext
jeder gescannten Datei ist zur Extraktionszeit bereits gelesen — die reale Deklaration ist ohne
zusätzliches Datei-I/O verfügbar, dieselbe Beobachtung, die [ADR-0022](0022-datei-mengen-bewusste-mehr-wurzel-aufloesung.md)
trug.

## Optionen

| Weg | Idee | Bewertung |
|---|---|---|
| **A — additive Deklarations-Evidenz in `fixed-root`** | Ein Deklarations-Index (Top-Level-Deklarationsnamen je Datei) liefert eine Evidenzstufe „deklariert"; sie sticht den bloßen Datei-Namens-Match; automatisch konsultiert bei `mode: fixed-root` mit ≥ 2 Roots, **kein neuer `mode`**. | **Gewählt.** Rückwärtskompatibel wie [ADR-0022](0022-datei-mengen-bewusste-mehr-wurzel-aufloesung.md) (nur `fixed-root` ≥ 2 Roots + deklarations-bewusstes Backend ändert Verhalten), **keine** Konsumenten-Umstellung, subsumiert den ADR-0022-A1-Strength-Fix; die ehrliche Konsumenten-Config löst auf. |
| **B — eigener opt-in `mode`** | Deklarations-Auflösung als separater `resolution.mode` (der reservierte `namespace` oder ein neuer `symbol`). | Verworfen: Konsumenten müssen umstellen; überlappt konzeptionell mit dem für C# reservierten Namespace-Index ([ADR-0016](0016-resolution-sprach-parametrisch.md)); fragmentiert das Auflösungs-Modell. |
| **C — Residual weiter fail-closed** | Kein Deklarations-Treffer, aber Paket-Verzeichnis in ≥ 2 Roots → weiter Exit 2. | Verworfen: das ist genau der Falsch-Positiv-Kern und bräche den ganzen Scan. Fail-open/extern ist [AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)-konsistent (die „0 reale Treffer ⇒ extern"-Linie aus [ADR-0022](0022-datei-mengen-bewusste-mehr-wurzel-aufloesung.md)). |

## Entscheidung

**Weg A.** Bei `mode: fixed-root` mit **≥ 2** `roots` und einem **deklarations-bewussten**
Backend (diese Fassung: **Kotlin**) wird der interne FQN deklarations-bewusst aufgelöst:

1. **Deklarations-Index** (einmal pro Auflösungslauf, I/O-frei aus dem schon gelesenen
   Quelltext): je gescannte Datei die Menge ihrer Top-Level-Deklarationsnamen (Kotlin: `fun`
   inkl. Extension `fun R.name`, `val`/`var`/`const val`, `class`/`data`/`sealed`/`enum`/
   `annotation class`, `object`, `interface`, `sealed interface`, `typealias`), text-heuristisch
   ([ADR-0002](0002-text-heuristische-extraktion.md)). Nicht deklarations-bewusste Backends
   liefern ein leeres Set (No-op).
2. **Evidenz-Rangfolge** (stärkste zuerst): **deklariert** (eine gescannte Datei in
   `root/pkg/` trägt eine Top-Level-Deklaration `sym`, unabhängig vom Dateinamen) >
   **nur-Paketverzeichnis** (`root/pkg/` existiert, aber keine Datei deklariert `sym` — **auch
   eine gleichnamige `sym.kt`, die `sym` nicht deklariert, zählt nur hier**) > **keine**. Für
   deklarations-bewusste Backends **löst die echte Deklaration den bisherigen Datei-Namens-Match
   ab**: trägt Root A eine `Foo.kt` ohne `Foo`-Deklaration und Root B die echte `Foo`-Deklaration
   (in `Types.kt`), gewinnt **Root B**. Nicht deklarations-bewusste Backends behalten die
   [ADR-0022](0022-datei-mengen-bewusste-mehr-wurzel-aufloesung.md)-Semantik (Datei-Match >
   Verzeichnis) unverändert.
3. **Schicht** am Pfad des **auflösenden** Roots
   ([SPEC-RULE-001](../../../spec/spezifikation.md#spec-rule-001--regel-auswertung)-Glob-Auswahl),
   nicht am Wurzel-Präfix. Es entscheidet die **stärkste vorhandene Evidenzstufe**; **genau ein**
   Root auf ihr (oder alle dieselbe Schicht) ⇒ eindeutig. **≥ 2** Roots derselben Stufe in
   **verschiedenen** Schichten: auf Stufe **deklariert** ⇒ echte Mehrdeutigkeit, **Exit 2 nach dem
   Scan**; auf Stufe **nur-Paketverzeichnis** (kein Deklarations-Treffer) ⇒ **extern** (fail-open —
   ohne Deklaration nicht diskriminierbar, genau der zuvor spurios mit Exit 2 quittierte Fall). Die
   Stufe *nur-Paketverzeichnis* **löst** damit rückwärtskompatibel, wenn sie eindeutig ist; ganz ohne
   Evidenz ⇒ **extern** (Phantom). `expect`/`actual` (same-layer) löst sauber; Tech-Regeln greifen
   unverändert am Roh-Symbol.

Der `fixed-root`-Realitäts-Filter aus
[ADR-0022](0022-datei-mengen-bewusste-mehr-wurzel-aufloesung.md) wird damit **um die
Deklarations-Evidenz erweitert**, nicht ersetzt; die Deklarationen sind **neue Nutzlast** und
fließen durch den Extraktions-Port und das Core-Modell (nicht nur der Fehler-Kanal wie bei
ADR-0022).

## Konsequenzen

- **Der Split-Package-Fall löst korrekt:** ein Top-Level-Symbol, dessen Datei ≠ Symbolname ist,
  löst über die reale Deklaration auf sein Modul auf (`asJdbc` → `adapters`, `ChunkColumnSchema`
  → `ports`) — die verbotene Kante wird gemeldet (Exit 1) statt der Scan mit Exit 2 abgebrochen.
- **Rückwärtskompatibel:** nur `fixed-root` mit ≥ 2 Roots **und** einem deklarations-bewussten
  Backend (Kotlin) ändert Verhalten; im Normalfall Datei==Deklaration ⇒ identisches Ergebnis.
  Nicht-Kotlin, `path`/`relative`/1-Root/kein-`resolution` und der bestehende KMP-disjunkte-Fall
  (Dateien nach Klassen benannt) laufen unverändert.
- **Subsumiert den ADR-0022-A1-Fix:** der „Klasse direkt unter `package_base` / Split-Package"-
  Phantomfall, den ADR-0022 per `strength`-Stratifizierung entschärfte, wird jetzt direkt über die
  echte Deklaration entschieden.
- **Ehrliche Grenzen** ([AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)):
  ein intern gemeintes Symbol, das in **keiner** gescannten Datei als Top-Level-Deklaration
  auftaucht (verschachtelte Klasse, `object`-Member, nicht gescannter/generierter Code,
  Star-Import), bleibt still extern; die Deklarations-Auflösung ist in dieser Fassung
  **Kotlin-only** (übrige Backends: package==directory); datei-tiefe Layer-Globs und per-Root
  `package_base` bleiben Grenzen/additive Folgeschritte.

## Fitness Function

- `make test`: Split-Package-Top-Level-Symbol (Datei ≠ Name) genau in **einem** Root deklariert →
  löst dorthin (vor der Änderung: Exit 2); Rangfolge (`Foo.kt` ohne `Foo`-Deklaration in Root A
  vs. echte Deklaration in Root B) → **Root B** gewinnt; verbotene Kante gemeldet (Exit 1); real
  **deklariert** in ≥ 2 Roots verschiedener Schichten → Exit 2 (stderr-Meldung, kein
  stdout-Befund, deterministischer Zeuge); Paket-Verzeichnis in ≥ 2 Roots **ohne** Deklaration →
  extern (kein Exit 2, kein Befund); nicht-Kotlin-Backend → Deklarations-Set leer (No-op);
  KMP-disjunktes Fixture → byte-identisch. Unit: Kotlin-Deklarations-Extraktion (je Decl-Art +
  Extension-Fun + Kommentar/String-Negativprobe), Evidenz-Rangfolge, distinct-layer-Ambiguität.
- `make arch-check` (Dogfooding): unverändert 0 (a-check hat keinen `resolution`-Block; Go ist
  nicht deklarations-bewusst).

## Re-Evaluierungs-Trigger

- **Weitere deklarations-bewusste Backends:** Java ist syntaxnah und ein naher additiver
  Kandidat, sobald ein Konsument es braucht.
- **Verschachtelte-Klassen-/Member-Importe:** falls ein realer Konsument die heute externe Grenze
  braucht, Re-Eval des Deklarations-Prädikats.
- **C#-Namespace-Index:** bleibt die eigene reservierte Achse
  ([ADR-0016](0016-resolution-sprach-parametrisch.md)-`mode`-Diskriminator), unberührt.
- **Per-Root `package_base`:** unverändert offen
  ([ADR-0022](0022-datei-mengen-bewusste-mehr-wurzel-aufloesung.md)).

## Geschichte

| Datum | Ereignis |
|---|---|
| 2026-07-06 | Proposed — d-migrate-Pilot-Evidenz: die einzige korrekte Vollrichtungs-Config (`fixed-root`, ein Root je Gradle-Modul) endet an einem Split-Package-Top-Level-Symbol (`asJdbc`, Datei ≠ Name) in **Exit 2**, real reproduziert gegen `v0.11.0`; der Auflösungs-Entwurf + adversarisches Multi-Linsen-Review (Residual-Eindeutigkeit, Evidenz-Rangfolge, Datenmodell-Plumbing, Backend-Scope) eingearbeitet; Lastenheft-CR 0.18.0. Entscheid **Weg A** (additiv in `fixed-root`) per Auftraggeber-Wort. |
| 2026-07-06 | Umsetzungs-Verifikation: Code + 8 AC-Tests grün (`make ci`); die reale Probe gegen d-migrate (Vollrichtungs-Config) tilgt den `asJdbc`-Exit-2 und deckte einen zweiten Fall auf — ein **Wildcard-/Paket-Import** trifft ein ganzes Paket-Verzeichnis (Stufe *nur-Paketverzeichnis*), löst bei einem Split über Schichten daher **fail-open** statt Exit 2 (`strength`-Wildcard-Zweig auf die schwache Stufe gelegt). |
| 2026-07-06 | Proposed → Accepted (Sign-off Auftraggeber per Merge-Wort). Ab jetzt immutable; Ablösung nur via Folge-ADR mit `Supersedes`. |
