# slice-016 — Regex-fähige `tech`-Muster (`match: substring|regex`)

**Status:** done (2026-07-01). Abnahme (Sign-off), Umsetzung, `make gates`, adversarisches
Multi-Linsen-Review + Delta-Re-Review erledigt (§7/§8; Synthese
[`docs/reviews/2026-07-01-slice-016-regex-tech-muster.md`](../../../reviews/2026-07-01-slice-016-regex-tech-muster.md)).
**Bezug:** setzt [ADR-0015](../../adr/0015-regex-tech-muster.md) um; erweitert
[AC-FA-RULE-003](../../../../spec/lastenheft.md#ac-fa-rule-003--tech-kapselung-regel-tech-leak)
(Regel-Semantik) + [AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml)
(Schema); wahrt [AC-QA-01](../../../../spec/lastenheft.md#ac-qa-01--determinismus) und die
Heuristik-Grenze [AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze).
[Roadmap](../in-progress/roadmap.md). **Evidenz:** b-cad-Pilot (Regel-E-Lücke — Qt als
`Q[A-Za-z]`, mit Substring nicht fassbar).

## 1. Auslöser

Der b-cad-Pilot ersetzt `tools/arch-check.sh` (Regeln A–E) durch `make a-check`. a-check bildet
**A–D** ab, aber **Regel E** (Qt nur in `adapters/ui/` + Composition Root `main.cpp`) ist
arch-check.shs `grep -E 'Q[A-Za-z]'` — mit dem heutigen Substring-`tech.pattern` nicht
ausdrückbar ([ADR-0015](../../adr/0015-regex-tech-muster.md) Kontext). Ohne Regel E bleibt
arch-check.sh nur teilweise ersetzt.

## 2. Betroffene Artefakte (vor der Implementierung benannt)

- **Slice-ID:** slice-016.
- **AC:** [AC-FA-RULE-003](../../../../spec/lastenheft.md#ac-fa-rule-003--tech-kapselung-regel-tech-leak)
  (erweitert), [AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml)
  (Schema erweitert); Rahmen [AC-QA-01](../../../../spec/lastenheft.md#ac-qa-01--determinismus)/[AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze).
- **ADR:** [ADR-0015](../../adr/0015-regex-tech-muster.md) (Proposed → Accepted auf Sign-off);
  Re-Eval von [ADR-0003](../../adr/0003-config-modell-a-check-yml.md).
- **Spec:** [SPEC-CONF-001](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema)
  (Schema), [SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung)
  (`tech-leak`-Zweig), Notiz in [SPEC-EXTRACT-001](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion).
- **Module:** `internal/hexagon/core` (`model.go` — `Tech`; `rules.go` — `matchTech`),
  `internal/adapter/driven/config` (Decode/Validierung), `internal/cli` (`--print-config`-Muster).
- **Nutzer-Doku (Sweep, §3.6):** [`docs/user/benutzerhandbuch.md`](../../../user/benutzerhandbuch.md)
  (`tech`-Block ohne `match`), [`README.md`](../../../../README.md) (Konfigurations-Abschnitt),
  [`CHANGELOG.md`](../../../../CHANGELOG.md) (Unreleased).
- **Gates:** `make gates` (lint/test/coverage-gate/arch-check/doc-check/gate-consistency/guard-selftest),
  danach `make ci` (+ `image-test`).

## 3. Umfang (Reihenfolge: Lastenheft → ADR → Spec → Code → Tests → Doku)

Versions-Bump: **0.8.0** für Lastenheft **und** Spezifikation. `0.8.0` war auch von
[slice-013](../open/slice-013-driving-driven-vertiefung.md) reserviert; slice-016 landete zuerst
und nimmt 0.8.0 — slice-013 hebt bei seinem Merge auf die nächste freie Minor (§5).

1. **Lastenheft** (nächste Minor):
   - [AC-FA-RULE-003](../../../../spec/lastenheft.md#ac-fa-rule-003--tech-kapselung-regel-tech-leak)
     (**Regel-Semantik**): das einem Adapter zugeordnete Muster matcht als **Substring (Default)**
     *oder* als **RE2-Regex** (`match: regex`, unverankert). Neue Verhaltens-AC (§4a).
   - [AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml)
     (**Schema/strict-decode**): `tech`-Eintrag um optionales `match: substring|regex`
     (Default `substring`); unbekannter `match`-Wert **oder** nicht kompilierbare Regex → Exit 2.
     Neue Config-AC (§4b).
2. **[ADR-0015](../../adr/0015-regex-tech-muster.md)** `Proposed → Accepted`, in den
   [ADR-Index](../../adr/README.md) eintragen.
3. **Spezifikation** (nächste Minor):
   - [SPEC-CONF-001](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema):
     `match` im `tech`-Block.
   - [SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung):
     Regex-Zweig (unverankert, gegen das extrahierte Symbol) **plus Präzedenz-Richtigstellung** —
     `tech`-Muster lösen in **Deklarationsreihenfolge, Erst-Treffer** auf (uniform substring/regex),
     **nicht** „längster Präfix"; das gilt nur für `layers`-Globs. Das heutige Wording behauptet
     fälschlich „spezifischster/längster Präfix gewinnt" auch für `tech` — `matchTech`
     (`rules.go`) liefert real schon Erst-Treffer; die Spec wird an den Code angeglichen.
   - Notiz in [SPEC-EXTRACT-001](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion)
     (Regex matcht gegen das extrahierte Symbol).
4. **Code:** `core.Tech` um ein unexportiertes Matcher-Feld `match func(string) bool` (kapselt die
   kompilierte RE2-Regex und hält den Kern `regexp`-import-frei) + Konstruktor `NewTech(pattern,
   adapter, match)`; `matchTech`/`Tech.matches` mit Regex-Zweig (`MatchString`),
   Deklarationsreihenfolge unverändert; ein Literal-`Tech` matcht weiter als Substring
   (rückwärtskompatibel). Config-Decode ruft `NewTech` (strict; Exit 2 bei unbekanntem `match`,
   leerer/ungültiger Regex); `--print-config`-Gerüst um `match` ergänzt.
5. **Tests:** Config-Validierung (unbekannter `match` → Exit 2; ungültige Regex → Exit 2);
   Regex-`tech-leak` Happy/Boundary (Composition-Root)/Negative; **gemischte substring/regex-Treffer
   → Erst-Treffer in Deklarationsreihenfolge**; Substring-Pfad byte-identisch; Dogfooding bleibt 0;
   Determinismus.
6. **Nutzer-Doku-Sweep:** [`benutzerhandbuch.md`](../../../user/benutzerhandbuch.md) (`tech`-Block
   um `match` + Fehlerbehebung), [`README.md`](../../../../README.md) (Konfigurations-Abschnitt),
   [`CHANGELOG.md`](../../../../CHANGELOG.md) (Unreleased-Eintrag).

## 4. Neue/geänderte Akzeptanzkriterien (Entwurf)

### 4a. Verhalten — unter [AC-FA-RULE-003](../../../../spec/lastenheft.md#ac-fa-rule-003--tech-kapselung-regel-tech-leak) (`tech-leak`)

- **Happy:** Given ein `tech`-Eintrag `match: regex` und ein Symbol, das die Regex trifft, nur in
  seinem zugeordneten Adapter, when `a-check` läuft, then kein Befund.
- **Boundary (Composition-Root):** Given dasselbe (Regex-)Symbol in der konfigurierten Composition
  Root, when `a-check` läuft, then kein Befund (deklarierte Ausnahme).
- **Negative:** Given das Regex-Symbol außerhalb seines Adapters, when `a-check` läuft, then ein
  Befund (`tech-leak`) und Exit-Code 1.
- **Präzedenz:** Given mehrere `tech`-Muster (substring und/oder regex), die dasselbe Symbol
  treffen, when `a-check` läuft, then greift der **in Deklarationsreihenfolge erste** Treffer
  (deterministisch; kein „längster Präfix" für `tech`).
- **Rückwärtskompat:** Given ein `tech`-Eintrag **ohne** `match`, when `a-check` läuft, then
  Substring-Semantik wie bisher (byte-identische Ausgabe).

### 4b. Schema — unter [AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml) (strict-decode)

- **Config (gültig):** Given ein `tech`-Eintrag mit `match: substring` oder `match: regex`, when
  `a-check` lädt, then akzeptiert (bei `regex` wird `pattern` beim Laden als RE2 kompiliert).
- **Config-Fehler (Wert):** Given `match` mit einem anderen Wert als `substring`/`regex`, when
  `a-check` lädt, then Exit-Code 2 (kein stiller Default).
- **Config-Fehler (Regex):** Given `match: regex` mit einer als RE2 nicht kompilierbaren `pattern`,
  when `a-check` lädt, then Exit-Code 2.

## 5. Out-of-Scope / Koordination

- **Versions-Kollision mit [slice-013](../open/slice-013-driving-driven-vertiefung.md):** beide reservieren
  0.8.0. Regel: **wer zuerst gemergt wird, nimmt 0.8.0**; der zweite Slice hebt beim Merge auf die
  nächste freie Minor und zieht seine Historien-Einträge (Lastenheft/Spec) entsprechend nach. Kein
  fixer Termin (Roadmap-Wellen feuern auf Trigger).
- **Pro-Sprache** oder **verankerte** Muster, weitere Modi (`glob`) — Enum ist erweiterbar
  ([ADR-0015](../../adr/0015-regex-tech-muster.md) Re-Eval-Trigger).
- Der eigentliche **b-cad-Ersatz** (arch-check.sh raus, `make a-check` rein) ist ein Change *an
  b-cad* und braucht ein **neues a-check-Release + Digest-Re-Pin** (das Feature muss im
  veröffentlichten Image liegen). Dieser Slice liefert nur die a-check-Fähigkeit.

## 6. Gates

`make gates` grün (insb. Dogfooding 0, Determinismus, doc-check-Links) → `make ci` →
Multi-Linsen-Review (Korrektheit · Vertrag/Spec · Test-Abdeckung · Konvention) → Synthese unter
`docs/reviews/` → Merge/Push **erst auf explizites Wort**.

## 7. Definition of Done

- [x] Lastenheft **0.8.0** — [AC-FA-RULE-003](../../../../spec/lastenheft.md#ac-fa-rule-003--tech-kapselung-regel-tech-leak)
  (Verhalten: Substring|Regex, Präzedenz) + [AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml)
  (Schema `match`, Exit 2); Historienzeile.
- [x] [ADR-0015](../../adr/0015-regex-tech-muster.md) `Accepted` + [ADR-Index](../../adr/README.md).
- [x] Spezifikation **0.8.0** — [SPEC-CONF-001](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema)
  (`match`), [SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung)
  (Präzedenz-Richtigstellung: Erst-Treffer statt „längster Präfix" für `tech`), tech-leak-Zeile.
- [x] Code — `Tech.match`-Closure + `NewTech`, `matches`/`matchTech`, Config→Exit 2 (inkl. leeres
  Regex-Pattern fail-closed), `--print-config`.
- [x] Tests — Happy/Boundary/Negative/Präzedenz(beidseitig)/Rückwärtskompat/Modus-Trennung;
  Config-Fehler (unbekannt/leer/ungültig → Exit 2); Exit-Codes via `cli.Run` (regex-tech-leak → 1,
  ungültiges `match` → 2); `ignore_symbols`-Unterdrückung des `Q[A-Za-z]`/`Queue.h`-FP; Determinismus ≥2.
- [x] Nutzer-Doku — [Benutzerhandbuch](../../../user/benutzerhandbuch.md) §3.4/§4/§10 (1.11),
  [README](../../../../README.md), [CHANGELOG](../../../../CHANGELOG.md).
- [x] `make gates` grün; adversarisches Multi-Linsen-Review + Delta-Re-Review; Synthese
  [`docs/reviews/2026-07-01-slice-016-regex-tech-muster.md`](../../../reviews/2026-07-01-slice-016-regex-tech-muster.md).
- [x] Closure: **reiner** `git mv` nach `done/` (getrennt von Inhalts-Edits); **2 beobachtbare
  Kriterien** (§8) + **Lerneintrag** (§8).

## 8. Closure-Notiz

**Gate-Beleg:** `make gates` grün — `arch-check` (Dogfooding) `gesamt: 0 Befund(e)`, `doc-check`
(d-check) `0 Befund(e)`, alle Test-Pakete `ok`, `record-gates` geschrieben.

**2 beobachtbare Kriterien:**
1. Ein Repo mit `tech: [{pattern: "Q[A-Za-z]", adapter: "adapters/ui", match: regex}]` und einem
   Qt-Include außerhalb `adapters/ui` liefert `tech-leak` und Exit-Code 1
   (`TestTechRegexLeakExit1`); dasselbe Muster ohne `match` bleibt Substring und byte-identisch.
2. Ungültiges `match` (`glob`), leere oder nicht kompilierbare Regex → Exit-Code 2
   (`TestTechRegexInvalidMatchExit2`, `TestTechMatch{Empty,Invalid}RegexFailsClosed`).

**Lerneintrag:**
- Das Review deckte eine **latente Spec-Ungenauigkeit** auf:
  [SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung) behauptete
  „längster Präfix" auch für `tech`-Muster, obwohl `matchTech` seit jeher Erst-Treffer in
  Deklarationsreihenfolge liefert. Neue Fähigkeiten machen alte, nie geprüfte Formulierungen sichtbar.
- Ein **leeres Regex-Pattern** hätte via `regexp.Compile("")` jeden Import getroffen (FP-Flut,
  schattet nachfolgende tech-Regeln) — zwei Linsen fanden das unabhängig; Fix: fail-closed statt
  „matcht alles", Parität zur Substring-Leerguard.
- Die **Quellen-Präzedenz-Matrix verbietet Spec→ADR-Verweise**: ADR-Referenzen aus
  Lastenheft/Spezifikation entfernen (der ADR ist downstream; er zitiert die Spec, nicht umgekehrt).
