# Review — slice-011: Rolle `app` + strenge `domain` (welle-10b/b2a)

**Datum:** 2026-06-22 · **Gegenstand:** Implementierung von [AC-FA-RULE-007](../../spec/lastenheft.md#ac-fa-rule-007--rolle-app-und-strenge-domain) / [ADR-0011](../plan/adr/0011-domain-application-trennung-rolle-app.md) — neue Rolle `app` (→ Befund `app-impurity`) und Verschärfung von `domain` (`domain↛port` kategorisch). · **Methode:** vier adversarische Linsen (read-only, parallel) + eigener Grep-Sweep für die Vertrag/Spec-Linse (deren Agent leer zurückkam). · **Gate-Stand:** `make gates` grün — lint/test/coverage; `arch-check` 0 (Dogfooding unverändert); `doc-check` 0/44.

## Linsen & Urteile

| Linse | Urteil |
|---|---|
| **Code & Korrektheit** | **0 Befunde, freigabefähig.** `impurityFinding`-Extraktion verhaltensäquivalent + Erst-Treffer-Ordnung erhalten; `domain↛{app,port,adapter}`/Tech kategorisch; `app`-Reinheit kategorisch, Richtung kanten-geregelt; Roleless-Ziel kein False-Positive (stdlib im Kern bleibt grün); Inferenz + `role:`-Vorrang; Konstrukt-`port-impurity` unberührt. |
| **Test-Abdeckung** | 2× MED (fehlende kategorische Pins); die „kategorisch"-Tests sind **echte Differenzial-Tests** (nicht tautologisch — würden bei edge-geregelter Regel rot). |
| **Rückwärtskompat & Doku** | 1× HIGH (`core-impurity`-Anforderung intern widersprüchlich), 2× MED (CHANGELOG/Handbuch-Migration), 2× LOW. Dogfooding-Grün am Code belegt. |
| **Vertrag & Spec** (Agent leer → eigener Grep-Check) | Anchor über alle 7 Verweise konsistent; „fünf→sechs"-Sweep fand **eine** von den Linsen übersehene Stelle. |

## Befunde & Disposition (alle eingearbeitet)

| Schwere | Linse | Stelle | Fix |
|---|---|---|---|
| **HIGH** | Vertrag | Die `core-impurity`-Anforderung blieb im Volltext bei „weder Adapter noch Tech"; ihr Boundary-Kriterium („Kern nutzt einen erlaubten gemeinsamen Port ⇒ kein Befund") **widersprach** b2a direkt. | Beschreibung um Port/`app` + kategorische Schärfung (Verweis auf [AC-FA-RULE-007](../../spec/lastenheft.md#ac-fa-rule-007--rolle-app-und-strenge-domain)) erweitert; Boundary auf „nur Domäne (gleiche Rolle) + stdlib" umgeschrieben. ✓ |
| **HIGH (neu)** | Vertrag | `lastenheft.md:22` „dieselben fünf Regeln" (§1-Narrativ) — von allen vier Linsen übersehen, vom eigenen Sweep gefunden. | → „sechs". ✓ |
| **MED** | Test | Kein `domain→adapter`-Pin *mit Kante* (nur ohne) — der Adapter-Zielarm der Schärfung war offen für stille Regression. | `TestDomainImportsAdapterCategorical` (Kante + `len==1`). ✓ |
| **MED** | Test | `domain→tech` (`isTech`-Arm) ungetestet. | `TestDomainImportsTech` (`core-impurity` **vor** `tech-leak`). ✓ |
| **MED** | Doku | CHANGELOG „Rückwärtskompatibel" verharmloste den Breaking Change für Fremd-Repos. | Auf „**Breaking für geprüfte Repos**" + sichtbaren Migrationsweg umformuliert. ✓ |
| **MED** | Doku | Migration/Invariante fehlte im Handbuch. | Glossar „Kern" (kennt keine Ports) + Historie-Zeile 1.5. ✓ |
| **LOW** | Doku | `architecture.md` nennt die `app`-Rolle nicht. | **Bewusst offen:** beschreibt a-checks *eigene* Architektur (kein `app`-Layer); das Konsumenten-Modell gehört nicht ins Sicht-Stratum. |
| **LOW** | Doku | Handbuch „Software-Version 0.1.0" vs. Lastenheft 0.5.0. | **Bewusst offen:** prä-b2a-Drift; Software-Version = letzter Release, b2a ist `[Unreleased]` — konsistent mit der 1.4-Praxis. |

## Bestätigt (hält der Prüfung stand)

- **Dogfooding grün, belegt:** `internal/hexagon/core/{rules,model}.go` importieren nur stdlib; `port → core` läuft über die inferierte Rolle `port` (nicht `domain`), und die Schärfung trifft ausschließlich `srcRole == "domain"`. Eigen-`.a-check.yml` unverändert.
- **Erst-Treffer-Ordnung** `core-impurity → app-impurity → port-impurity → lateral-adapter → tech-leak → wrong-direction` ([SPEC-RULE-001](../../spec/spezifikation.md#spec-rule-001--regel-auswertung)) deckt sich mit `impurityFinding` (vor dem Rest-`switch`).
- **Roleless-Invariante:** ein `domain`/`app`/`port`-Import auf ein rollenloses/stdlib-Ziel erzeugt kein Impurity (positiver Rollen-Test, kein „≠ domain") — kritisch, damit stdlib im Kern grün bleibt.
- **Anchor** `#ac-fa-rule-007--rolle-app-und-strenge-domain` über alle 7 Verweise identisch.

## Nicht-Befund (geprüft)

- README „je eine Anforderung" / „sechs Regeln": sechs Befund-Namen ↔ je eine `AC-FA-RULE-*`-Anforderung (006 ist der Rollen-Mechanismus, kein Befund) — Aussage hält.
- Übrige `fünf`-Treffer (CHANGELOG-`[0.1.0]`, mehrere Fundament-/Dispatch-ADRs, Review-Artefakte, Lastenheft-`0.1.0`-Historie): historisch bzw. immutable — korrekt unangetastet.
