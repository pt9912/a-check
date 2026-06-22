# Review — Ports dürfen Domänen-Typen referenzieren ([ADR-0008](../plan/adr/0008-ports-duerfen-domaenen-typen-referenzieren.md), Lastenheft 0.2.0)

**Datum:** 2026-06-22

**Gegenstand:** uncommitteter Stand der [AC-FA-RULE-004](../../spec/lastenheft.md#ac-fa-rule-004--port-disziplin-regel-port-impurity)-Neufassung (Lastenheft 0.1.0→0.2.0) + Reorg zur echten `ports`-Schicht; [ADR-0008](../plan/adr/0008-ports-duerfen-domaenen-typen-referenzieren.md) Status `Proposed`.

**Methode:** vier perspektiven-diverse, adversarische Linsen (read-only, je ein Reviewer; Regelwerk Modul 10) — Vertrag/Lastenheft · Spec-Konsistenz · Code-Korrektheit · Regelwerk/Konvention.

**Gesamtbewertung:** Der Kern ist sauber, symmetrisch erzählt und über das Dogfooding (`make arch-check` grün auf der echten `ports`-Schicht) belegt. **Kein Blocker.** Vor Commit + ADR-Abnahme empfohlen: eine Test-Lücke schließen und drei nicht nachgezogene Nachbarstellen angleichen.

## Konsolidierte Befunde (nach Priorität)

| ID | Schwere | Linse(n) | Ort | Befund | Empfehlung |
|---|---|---|---|---|---|
| R1 | MAJOR | Code + Vertrag | `internal/hexagon/core/rules_test.go` | Kein Test für `ports→core` **ohne** deklarierte Kante → `wrong-direction`. Der Kern-Claim „edge-regiert" ist nur im Positivfall getestet; eine Regression, die `ports→core` bedingungslos durchwinkt, bliebe grün (`arch-check` deckt es nicht, weil das Eigen-`.a-check.yml` die Kante hat). | Regressionstest ergänzen: `testModel()` ohne `{ports,core}`-Edge, `ports`-Datei importiert `core/x` ⇒ `wrong-direction`. |
| R2 | MAJOR | Spec-Konsistenz | `docs/user/benutzerhandbuch.md:143`; `README.md:25` | Rule-Beschreibungen behaupten weiter das **alte** Verbot („Port importiert Kern" / „Ports sind reine Abstraktionen") — direkter Widerspruch zur 0.2.0-Regel. | Beschreibungen umstellen: Port verboten ⇒ Adapter/Tech; Kern-Referenz erlaubt. |
| R3 | MAJOR | Spec-Konsistenz + Code | `internal/cli/cli.go` (`sampleConfig`, `--print-config`) | Das vom Tool selbst emittierte `.a-check.yml`-Gerüst zeigt das alte 2-Schichten-Modell (kein `ports`-Layer, keine `ports→core`-Kante) — Dogfooding-Widerspruch zur nun kanonischen Struktur. | Gerüst auf das 0.2.0-Modell heben (`ports`-Layer + `{from: ports, to: core}`). |
| R4 | MINOR | Spec-Konsistenz | `spec/architecture.md:72` (§3) | „Adapter hängen **ausschließlich** von Ports ab" ist jetzt falsch: Adapter importieren real auch `core` (seit `port` ≠ `core`-Paket). §2-Mermaid blendet den `adapter→core`-Pfeil aus. | „von Ports **und** Domänentypen des Kerns"; §2 optional um `CFG/EXT/REP --> CORE` ergänzen. |
| R5 | MINOR | Vertrag | `AC-FA-RULE-004` Boundary-AC | Das Boundary-AC testet `ports→ports`-Re-Export ([AC-FA-RULE-005](../../spec/lastenheft.md#ac-fa-rule-005--schicht-richtung-regel-wrong-direction)/`allow`-Mechanik), nicht den Vertragskern; die Boundary-Verengung steht nicht in der Historie. | Optional: Boundary auf den Kern ausrichten **oder** Historie ergänzen. (Anm.: `ports→core`-ohne-Kante ist sachlich `wrong-direction`/RULE-005-Domäne.) |
| R6 | MINOR | Code | `rules.go` `adapterSeg` + `.a-check.yml` | `adapterSeg` sucht das Literal-Segment `"adapters"`; die neuen Adapter liegen unter `adapter/driven/` → `lateral-adapter`-Unterscheidung und `adapter_sink` greifen im Eigen-Layout nicht mehr. **Latent** (a-checks Adapter importieren einander nicht; bestehender Defekt, nicht durch den Rename eingeführt). | Known Limitation; gehört in welle-10 (`adapterSeg`-Generalisierung). |
| R7 | NIT | Vertrag | `AC-FA-RULE-004` Beschreibung | „Vendor" als Begriff ohne Schema-Verankerung (Config kennt nur `tech:`/Framework). | streichen oder klar als Beispiel an `tech:` binden. |
| R8 | NIT | Spec-Konsistenz | `README.md:117`; `benutzerhandbuch.md:170,235` | Beispiel-Globs zeigen altes `internal/core`/`internal/adapters` (nicht falsch, aber ohne `ports`-Layer). | Beispiele auf die neue Struktur heben. |

## Pro Linse (Kurzfazit)

- **Vertrag/Lastenheft:** Vertraglich sauber gemacht — direkter Change-Request (nicht per ADR), korrekter minor-Bump, widerspruchsfreie Regel-Zuständigkeiten, Anker valide. Schwäche: die zwei *eigentlichen* Vertragswechsel (Kern-Import erlaubt; Port-Tech jetzt `port-impurity` statt `tech-leak`) leben primär in der Prosa (R5/R7).
- **Spec-Konsistenz:** Die ports↔core-Wende ist über Lastenheft/Spezifikation/Architektur/Code/Config/ADR durchgängig und symmetrisch; §2-Pfeilrichtung korrekt gedreht, Strata-Disziplin gewahrt, Versions-Bumps konsistent. Drei Nachbarstellen tragen noch die alte Geschichte (R2/R3/R4).
- **Code-Korrektheit:** `port-impurity`-Fix semantisch korrekt (alle vier Port-Fälle durch `ruleFor` verifiziert, Erst-Treffer-Reihenfolge intakt), Reorg im Produktionscode vollständig, `port.go`-Signaturen passen, kein toter Code. Eine Test-Lücke (R1), zwei unabhängige MINOR (R3/R6).
- **Regelwerk/Konvention:** **Konform, keine Blocker/Major.** Regel-Lockerung durch [ADR-0008](../plan/adr/0008-ports-duerfen-domaenen-typen-referenzieren.md) gedeckt (§3.6), kein Spec-Stratum referenziert ADRs (§3.4), keine Accepted-ADR berührt (§3.5), IDs real (§5), Anforderungsprozess vollständig, `.golangci.yml` ist Pfad-Nachzug (keine Lockerung). §3.3: alle Renames >50 % Ähnlichkeit (niedrigste `model.go` 69 %) ⇒ `git log --follow` bleibt zuverlässig, **ein Commit vertretbar**.

## ADR-0008-Abnahme

[ADR-0008](../plan/adr/0008-ports-duerfen-domaenen-typen-referenzieren.md) ist inhaltlich tragfähig und durch das Dogfooding belegt; [SPEC-RULE-001](../../spec/spezifikation.md#spec-rule-001--regel-auswertung)-Schärfung korrekt aufwärts. **Empfehlung:** auf `Accepted` setzen **nach** Schließen von R1–R3 (Test + Doku-/Gerüst-Angleich); R4 mitnehmen; R5/R7/R8 optional; R6 in welle-10 verschoben.

## Resolution (2026-06-22)

Nach dem Review umgesetzt und `make gates` grün: **R1** (`TestPortToCoreWithoutEdge` — `ports→core` ohne Kante ⇒ `wrong-direction`), **R2** (`benutzerhandbuch.md`/`README.md` `port-impurity`-Beschreibung auf 0.2.0), **R3** (`--print-config`-Gerüst mit `ports`-Schicht + `ports → core`), **R4** (`architecture.md` §3-Bullet + §2-Mermaid `adapter→core`), **R7/R8** (Begriff „Vendor" entschärft, Beispiel-Globs gehoben). **R5** bewusst offen (RULE-005-Nuance). **R6** → welle-10 (`adapterSeg`-Generalisierung). [ADR-0008](../plan/adr/0008-ports-duerfen-domaenen-typen-referenzieren.md) auf `Accepted` gesetzt.
