# ADR-0015 — Regex-fähige `tech`-Muster (RE2, opt-in via `match`)

- **Status:** Accepted
- **Datum:** 2026-07-01
- **Autor:** pt9912
- **Bezug:** [AC-FA-RULE-003](../../../spec/lastenheft.md#ac-fa-rule-003--tech-kapselung-regel-tech-leak), [AC-FA-CONF-001](../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml) — **Re-Evaluierung** von [ADR-0003](0003-config-modell-a-check-yml.md) (Config-Modell; erweitert, kein Supersede).
- **Schärft:** [SPEC-CONF-001](../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema) + [SPEC-RULE-001](../../../spec/spezifikation.md#spec-rule-001--regel-auswertung) — die `tech`-Muster-Matching-Semantik.
- **Supersedes:** —

## Kontext

Die Regel `tech-leak` ([AC-FA-RULE-003](../../../spec/lastenheft.md#ac-fa-rule-003--tech-kapselung-regel-tech-leak))
ordnet ein Framework/Tech-Muster einem Adapter zu; das Muster matcht heute als **Substring**
(`strings.Contains(imp.Symbol, t.Pattern)`, `internal/hexagon/core/rules.go`). [AC-FA-RULE-003](../../../spec/lastenheft.md#ac-fa-rule-003--tech-kapselung-regel-tech-leak)
nennt selbst `Qt → UI-Adapter` als Beispiel — aber **Qt ist als Substring nicht fassbar:** die
Header heißen `QWidget`, `QString`, `QtGui/…` — es gibt kein gemeinsames Substring, nur das
**Muster** `Q[A-Za-z]`.

Belegt durch den **b-cad-Pilot** (welle-05/-06, [Roadmap](../planning/in-progress/roadmap.md)):
Ziel ist, b-cads `tools/arch-check.sh` (Regeln A–E) durch das `make a-check`-Gate zu ersetzen.
a-check bildet **A–D** heute ab — `core-impurity` (Regel A), `lateral-adapter` (Regel B),
`tech` `.hxx → geometry` (Regel C), `tech` `sqlite3 → persistence` (Regel D). **Regel E**
(Qt nur in `adapters/ui/` + der Composition Root `main.cpp`) ist in arch-check.sh ein
unverankertes `grep -E 'Q[A-Za-z]'` — mit Substring **nicht** ausdrückbar. Ohne Regel E bliebe
arch-check.sh nur **teilweise** ersetzt; das Pilot-Ziel „ein Tool statt vier Skripte" verfehlt.

## Optionen

| Weg | Idee | Bewertung |
|---|---|---|
| **A — `pattern` global auf Regex** | Die bestehende `tech.pattern`-Semantik von Substring auf Regex umstellen. | **Verworfen.** Bricht bestehende Configs **still**: `sqlite3*`, `.hxx`, `net/http`, `gopkg.in/yaml` bekommen Regex-Bedeutung (`.` = beliebig, `*` = Quantor) → Determinismus-/Rückwärtskompat-Bruch, gegen „keine stillen Defaults" ([AC-QA-01](../../../spec/lastenheft.md#ac-qa-01--determinismus)). |
| **B — opt-in `match: substring\|regex`** | Je `tech`-Eintrag ein optionales `match` (Default `substring`); `regex` = RE2. | **Gewählt.** Additiv, Substring bleibt byte-identisch, RE2 deterministisch; Enum erweiterbar. Folgt dem [ADR-0014](0014-resolution-roots.md)-Muster (Re-Eval im *Bezug*, `Supersedes: —`, ein einziges Config-Feld). |
| **C — Konsument listet Qt-Header** | b-cad trägt jeden konkreten Qt-Header als Substring ein. | **Verworfen.** Brüchig (jeder neue Qt-Header = Config-Edit), reproduziert die `Q[A-Za-z]`-Semantik nicht, ist kein echter Ersatz. |

Innerhalb von Weg B wurden drei Oberflächen erwogen — eigener Schlüssel `regex:` (XOR zu
`pattern`), Boolean `regex: true`, und **Enum `match: substring|regex`**. Gewählt ist das **Enum**:
explizit, ein einzelnes Default-behaftetes Feld (wie das optionale `direction`/`role`), und um
weitere Modi (z. B. `glob`) erweiterbar, ohne ein zweites XOR-Feld einzuführen.

## Entscheidung

**Weg B, Enum-Form.**

1. Ein `tech`-Eintrag ist `{pattern, adapter, match?}` mit `match ∈ {substring, regex}`,
   **Default `substring`**.
2. `substring` = heutiges `strings.Contains` (**unverändert**). `regex` = RE2 (Go `regexp`),
   **unverankerter** Suchlauf (`MatchString`) gegen das extrahierte Import-Symbol
   ([SPEC-EXTRACT-001](../../../spec/spezifikation.md#spec-extract-001--import-extraktion)) —
   deckt sich mit arch-check.shs unverankertem `grep -E`.
3. **RE2**: linear, deterministisch, kein katastrophales Backtracking → wahrt
   [AC-QA-01](../../../spec/lastenheft.md#ac-qa-01--determinismus) und die Hermetik.
4. **Ungültige Regex** oder **unbekannter `match`-Wert** → **Exit 2** beim Config-Laden
   (strict-decode, kein stiller Default); die Regex wird beim Laden **einmal** kompiliert.
5. Die Heuristik-Grenze bleibt unverändert: dieselbe `Q[A-Za-z]`-vs-`Queue.h`-False-Positive-
   Semantik wie arch-check.sh; `markers.ignore_symbols` bleibt die dokumentierte Ausnahme
   ([AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)).
6. **Präzedenz bei Mehrfach-Treffern:** treffen mehrere `tech`-Muster (substring und/oder regex)
   dasselbe Symbol, greift der **in Deklarationsreihenfolge erste** Treffer — uniform für beide
   Modi, deterministisch. Für `tech` gibt es **kein** „längster Präfix" (das gilt nur für
   `layers`-Globs); [SPEC-RULE-001](../../../spec/spezifikation.md#spec-rule-001--regel-auswertung)
   wird entsprechend richtiggestellt (heutiges `matchTech` liefert bereits Erst-Treffer — die Spec
   wird an den Code angeglichen, kein Verhaltensbruch).

## Konsequenzen

- [ADR-0003](0003-config-modell-a-check-yml.md) **bleibt gültig** — ADR-0015 erweitert nur die
  Muster-Semantik (Substring → Substring | Regex), kein Supersede.
- **Schema** ([AC-FA-CONF-001](../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml)/[SPEC-CONF-001](../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema)):
  optionales `match` je `tech`-Eintrag; strict-decode (Exit 2 bei unbekanntem Wert/ungültiger Regex).
- Die `tech-leak`-Meldung (`rules.go`) nennt heute `tech.Pattern`; für `match: regex` bleibt das
  der Muster-String — korrekt, kein Sonderfall.
- **Bestehende Spec-Ungenauigkeit richtiggestellt:** [SPEC-RULE-001](../../../spec/spezifikation.md#spec-rule-001--regel-auswertung)
  schreibt „spezifischster/längster Präfix gewinnt" auch für `tech`-Muster; das trifft schon heute
  nicht zu (`matchTech` = Erst-Treffer in Deklarationsreihenfolge). Dieser ADR präzisiert die
  Präzedenz explizit (Entscheidung §6), damit Regex + gemischte Treffer nicht undefiniert bleiben.
- **Nutzer-Doku** (`docs/user/benutzerhandbuch.md`, `README.md`, `CHANGELOG.md`) zieht den
  `tech`-`match`-Zusatz nach — Teil des Slice-Sweeps, nicht des ADR.
- **b-cads Regel E wird ausdrückbar** (`{pattern: "Q[A-Za-z]", adapter: "src/adapters/ui", match: regex}`
  + `composition_root: ["src/main.cpp"]`) → `arch-check.sh` **vollständig** ersetzbar.
- **Bestehende Configs unverändert** — vier Konsumenten-Configs + [Dogfooding](../../../.a-check.yml)
  tragen kein `match`, also Substring, also byte-identisch.

## Fitness Function

- `make test`: Regex-`tech-leak` Happy/Boundary(Composition-Root)/Negative; ungültige Regex und
  unbekannter `match`-Wert → Exit 2; Substring-Pfad unverändert.
- `make arch-check` (Dogfooding, [AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)):
  unverändert **0** (a-check nutzt kein `match: regex`).

## Re-Evaluierungs-Trigger

- Falls je **pro-Sprache** oder **verankerte** Muster nötig werden → eigenes Inkrement
  (das Enum/Schema ist erweiterbar).
- Falls `match: regex` breite Adoption findet und der Substring-Default verwirrt → Re-Eval der
  Default-Wahl.

## Geschichte

| Datum | Ereignis |
|---|---|
| 2026-07-01 | Proposed — welle-05/-06, aus dem b-cad-Pilot (Regel-E-Lücke: Qt nicht als Substring fassbar). Weg B (opt-in `match: substring\|regex`, RE2) gegen Global-Regex (Weg A) und Header-Listen (Weg C); Enum-Oberfläche gegen `regex:`-Schlüssel/`regex: true`-Flag. |
| 2026-07-01 | Proposed → Accepted (Sign-off Auftraggeber: Weg B, Enum `match: substring\|regex`, RE2 unverankert, ungültig → Exit 2, Substring-Default unverändert; Präzedenz = Erst-Treffer in Deklarationsreihenfolge, Richtigstellung des `tech`-„längster Präfix"-Wordings). Ab jetzt immutable; Umsetzung [slice-016](../planning/done/slice-016-regex-tech-muster.md). Ablösung nur via Folge-ADR mit `Supersedes`. |
