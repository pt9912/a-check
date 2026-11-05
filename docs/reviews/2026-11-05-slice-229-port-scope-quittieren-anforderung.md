# Review-Report: slice-229 — 2026-11-05

**Review-Art:** Plan (Anforderung und Entscheidung gegen Plan, ADRs, Hard Rules; Zeilenform zusätzlich gegen den Bestandscode gelesen)

**Gegenstand:** slice-229, Commit 1bb7971 (Range 464f1f7..1bb7971 plus die Lifecycle-Commits davor): `spec/lastenheft.md` 0.31.0 (`AC-FA-CONF-002`, Historie-Zeile, Verweis in `AC-FA-CONF-001`), `ADR-0045`, ADR-Index, `docs/user/releasing.md` (Versions-Zahl).

**Skill:** `.harness/skills/reviewer.md` @ 1bb7971 · Status Accepted
**Modell:** claude-sonnet-5-5 · **Datum:** 2026-11-05

**Eingangs-Kontext:**

- slice-229 §1 (Ziel, Abgrenzung), §2 (DoD), §6
- ADR-0045 (Gegenstand), ADR-0040, ADR-0036
- `AC-FA-CONF-001`, `AC-FA-CONF-002`, `AC-FA-RULE-010`, `AC-QA-01`
- `AGENTS.md` (Hard Rules §3.4, §3.5, §3.7; §5 Regeln 4, 5, 9)
- Bestandscode: `writePortScopeNotice` (`internal/cli/cli.go`), `InertPortScope`/`InertPortScopes` (`internal/hexagon/core/rules.go`), `decodeLayerObject`/`knownLayerKey` (`internal/adapter/driven/config/config.go`)
- `v6.13.0` · `regelwerk/modul-05-planning-harness.md` §Ziel-Form: Slice

---

## Findings

| ID | Kategorie | Befund | Quelle | Pfad | Verifizierbar | Klasse |
|---|---|---|---|---|---|---|
| F-1 | MEDIUM | Reihenfolge, Deckel, Kopfzeile und Restzahl der neuen Zeilen sind nicht entscheidbar. ADR-0045 Punkt 2 und `AC-FA-CONF-002` (Determinismus) sagen „Reihenfolge der Schichten", der Bestand sortiert nach Glob (`sort.Slice` nach `Glob`, SPEC-CLI-001 „Stabil nach Glob sortiert"). Offen bleibt, wo quittierte und nicht quittierte Zeilen stehen und was die Kopfzeile „%d Port-Glob(s) …" zählt. Offen bleibt auch, ob „… und n weitere" Zeilen oder Globs zählt. Ein Deckel, der eine Quittierungs- oder „ohne Anlass"-Zeile abschneidet, verdeckt genau das, was die Meldung zeigen soll. `writePortScopeNotice` kehrt außerdem bei `len(scopes)==0` sofort zurück, die „ohne Anlass"-Zeile braucht aber gerade diesen Fall. | ADR-0045 Punkt 2, 3 und Konsequenzen; AC-FA-CONF-002 Determinismus; AC-QA-01 | `docs/plan/adr/0045-port-scope-meldung-quittieren.md:50-68,116`; `internal/cli/cli.go:167-183` | ja — `make test` (Ausgabe-Test mit Mischfall über dem Deckel) | Zeilenform unterbestimmt: Ordnung, Deckel, Zähler |
| F-2 | MEDIUM | „Schicht ohne `role: port`" ist mehrdeutig. Das bestehende Decoding prüft die **effektive** Rolle (Namens-Inferenz, `decodeLayerObject` Kommentar zu ADR-0036), `ports: {globs, acknowledge}` ohne `role:` ist dort eine Port-Schicht. AC-FA-CONF-002 und ADR-0045 Punkt 4 fordern wörtlich `role: port`. slice-230 muss raten, ob die inferierte Rolle genügt oder Exit 2 folgt. | AC-FA-CONF-002 Beschreibung/Negative; ADR-0045 Punkt 4; AC-FA-RULE-006 | `spec/lastenheft.md:757,785`; `internal/adapter/driven/config/config.go:545-556` | ja — `make test` (Fall ohne explizite `role:`) | Rollen-Bedingung explizit vs. effektiv nicht festgelegt |
| F-3 | MEDIUM | Zwei Zusagen ohne Lauf. (a) Das Negative-Kriterium verlangt, ein **älterer** `a-check` lehne den Schlüssel mit Exit 2 ab; die jeweils laufende Version kann das nie testen, und der Fitness-Function-Abschnitt der ADR nennt keinen Träger dafür. (b) Die AC schreibt „ein Satz, kein Ticket-Verweis allein"; ADR und Fitness Function prüfen nur „nicht leer / nur Leerraum". Die Einschränkung ist als Vertrag formuliert und von keinem Lauf gehalten. | `v6.13.0` · `regelwerk/modul-05-planning-harness.md` §Ziel-Form: Slice (testbare Kriterien); Skill Mess-Regel „Eine Zusage nennt den Lauf, der sie hält" | `spec/lastenheft.md:760-761,786-789`; ADR-0045 Punkt 5, Fitness Function | nein — kein Lauf könnte (a) bestätigen; (b) nur als Review-Urteil | Zusage im Vertrag ohne benannten Lauf oder „vom Review getragen" |
| F-4 | MEDIUM | Zuordnung Schicht ↔ Meldung ist an drei Stellen offen, die der Bestandscode real macht. (1) Eine Zeile trägt **eine** Ursache; eine Schicht kann Globs mit beiden Ursachen haben (Ursache wird je Glob über `inertCause` bestimmt). (2) Löst nur ein Teil der Globs einer Schicht die Meldung aus, steht die Zahl der auslösenden Globs; die übrigen sind unsichtbar, und „ohne Anlass" greift erst, wenn **keiner** auslöst. (3) `InertPortScopes` dedupliziert über `seen[g]` schichtübergreifend; derselbe Glob in zwei Schichten wird nur der ersten zugeordnet, die zweite mit Quittierung würde fälschlich „ohne Anlass" gemeldet. Zusätzlich: ADR sagt, die Abhilfe-Zeile entfällt, wenn **alle** Schichten dieser Ursache quittiert sind; die AC sagt „Abhilfe-Zeilen entfallen für eine quittierte Schicht" — die Abhilfe ist je Ursache, nicht je Schicht. | ADR-0045 Punkt 2, 3; AC-FA-CONF-002 Beschreibung und Happy/Boundary | `internal/hexagon/core/rules.go:410-431`; `internal/cli/cli.go:172-189`; `spec/lastenheft.md:765-768` | ja — `make test` (Mischfälle) | Mehrdeutigkeit bei Schicht mit gemischten oder teils auslösenden Globs |
| F-5 | LOW | Exit-2-Randfälle des Schlüssels sind nicht aufgezählt: `acknowledge` als Nicht-Map (Liste, Skalar, `null`), leere Map `{}`, `port-scope` mit Nicht-String-Wert oder `null`, doppelter Schlüssel. „Nur Leerraum gilt als leer" steht in der ADR, nicht in der AC; das Kriterium dort sagt nur „leer oder fehlend". | AC-FA-CONF-002 Negative; ADR-0045 Punkt 4; SPEC-CONF-001 (strikt) | `spec/lastenheft.md:785-788`; ADR-0045 Punkt 4 | ja — `make test` | Fail-closed-Aufzählung unvollständig |
| F-6 | LOW | Die Konsequenz „ein zusätzlicher Glob ändert die Zahl … verdeckt ihn nur in der Meldung" mildert das benannte Risiko nur scheinbar: Die Meldung ist für diese Globs das **einzige** Signal (`port-locality` bleibt dort still), und die Zahl hat keinen Soll-Wert, an dem ein Leser die Änderung erkennt. Das Risiko ist in slice-229 §6 und der ADR benannt; die Abschwächung trägt nicht. | slice-229 §6 (zweites Risiko); Skill Mess-Regel „Eine Zusage nennt den Lauf" | ADR-0045 Konsequenzen, zweiter Punkt | nein | Abmilderung ohne Wächter |
| F-7 | LOW | AC-FA-CONF-002 Boundary bündelt zwei Szenarien (zwei Schichten / überflüssige Quittierung) in einem Kriterium; Given/When/Then-Form ist da, die Einzelprüfbarkeit je Kriterium leidet. | AGENTS §5 Regel 9 | `spec/lastenheft.md:779-785` | nein (Stil, Review-Sache laut conventions) | Kriterium bündelt mehrere Szenarien |
| F-8 | INFO | Bis slice-230 widersprechen sich Rang 1 und Rang 2: SPEC-CLI-001 sagt weiter „Ein Weg … zu quittieren, ist nicht zugesagt", das Lastenheft sagt das Gegenteil. Das ist laut slice-229 §1 ausdrücklich der Folge-Slice-Gegenstand (slice-230 nennt beide Stellen); Source Precedence entscheidet bis dahin zugunsten des Lastenhefts. | slice-229 §1 Out-of-Scope; slice-230 Berührte Spec-Stellen | `spec/spezifikation.md:608-611` | nein | Übergangs-Divergenz Spec ↔ Lastenheft, benannt |
| F-9 | INFO | ADR-0045 steht auf `Accepted` vor diesem Review. Der Branch ist lokal (`main` 4 Commits vor `origin/main`), eine Nachbesserung der ADR zu F-1…F-5 ist also in der Range ohne `Supersedes` möglich, danach nicht mehr (AGENTS §3.5). Zeitfenster, keine Regelverletzung. | AGENTS §3.5 | `docs/plan/adr/0045-port-scope-meldung-quittieren.md:3` | ja — `make doc-immutable RANGE=` | ADR vor Review akzeptiert |

## Negativbefunde

| Bereich | Ergebnis |
|---|---|
| AC-Form (drei Pfade Happy/Boundary/Negative + Out-of-Scope, Given/When/Then, stabiler Anker `ac-fa-conf-002`, Versions-Bump 0.30.0 → 0.31.0, Historie-Zeile, Verweis in AC-FA-CONF-001) | geprüft, ohne Befund bis auf F-3/F-5/F-7; vierter Pfad „Determinismus" ist zulässig |
| Widerspruch ADR ↔ AC bei Exit-Code, Schlüsselname `acknowledge`/`port-scope`, Quittierung nur Meldung, überflüssige Quittierung advisory, Out-of-Scope-Liste | geprüft, ohne Befund (Abweichungen nur wie in F-1, F-4) |
| `Schärft:` der ADR (nur SPEC-CONF-001 und SPEC-CLI-001, nicht das Lastenheft); AC-Anlage nur im Lastenheft; ADR-Index nachgezogen | geprüft, ohne Befund (AGENTS §5 Regel 4/5, §3.4) |
| Spec-Straten-Richtung: neue Lastenheft-Zeilen enthalten keine ADR-, Slice- oder Commit-Kennung (Issue #7 steht nur in der Historie-Zeile; „ADR" kommt im Text nicht vor) | geprüft, ohne Befund |
| Realisierbarkeit „Diagnose bekommt eine Schicht-Zuordnung": `core.Layer.Name` existiert, `InertPortScope` trägt nur Glob/Scope/Cause und lässt sich erweitern | geprüft, ohne Befund (Einschränkung F-4 (3)) |
| Realisierbarkeit „Redundanz vor der Quittierung, ungedeckelt": `InertPortScopes` liefert ungedeckelt, der Deckel sitzt erst in `writePortScopeNotice` | geprüft, ohne Befund |
| Schema-Decoding: `knownLayerKey` lehnt `acknowledge` heute mit Exit 2 ab (Aussage der ADR zur Versions-Kopplung stimmt); Erweiterung um einen Schlüssel ist an einer Stelle möglich | geprüft, ohne Befund |
| Out-of-Scope slice-229 §1: keine Umsetzung, `spec/spezifikation.md` unberührt, andere Diagnosen nicht quittierbar, `port-locality` unberührt | geprüft, ohne Befund |
| Kommentar-/Zustandsfeld-Regel (§3.7): Historie-Zeile und ADR in Indikativ, keine Chronik in Zustandsfeldern | geprüft, ohne Befund |
| `docs/user/releasing.md` (Versions-Zahl 0.31.0), Index-Zeile ADR-0045 | geprüft, ohne Befund |

## Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 0 |
| MEDIUM | 4 |
| LOW | 3 |
| INFO | 2 |

**Finding-Klassen dieses Laufs:** Zeilenform unterbestimmt: Ordnung, Deckel, Zähler · Rollen-Bedingung explizit vs. effektiv nicht festgelegt · Zusage im Vertrag ohne benannten Lauf oder „vom Review getragen" · Mehrdeutigkeit bei Schicht mit gemischten oder teils auslösenden Globs · Fail-closed-Aufzählung unvollständig · Abmilderung ohne Wächter · Kriterium bündelt mehrere Szenarien · Übergangs-Divergenz Spec ↔ Lastenheft, benannt · ADR vor Review akzeptiert

## Verdikt

**Merge-blockierend:** ja — vier MEDIUM (F-1 bis F-4) sind vor dem Push zu klären, solange die ADR in der Range noch ohne `Supersedes` änderbar ist (F-9). HIGH: keines.

**Übergabe:** Findings gehen an den Implementer; die Finding-Klassen zusätzlich in die Slice-Closure §7. Kein Befund wurde auf Widerspruch hin herabgestuft; der Report ersetzt keine Verifikation gegen die DoD.
