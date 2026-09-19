# Review-Report: slice-194 — 2026-09-19

**Review-Art:** Code — unabhängiger Lauf. Geprüft wird der Diff gegen Plan, ADR-0040 und die
Repo-Konventionen (Modul 10); die DoD-/Spec-Konformität ist Sache der Verifikation und steht
hier nicht.

**Gegenstand:** Commits `f0fe0a6` (Plan), `7e8dc0a` (Lifecycle-Move), `fe01086` (Spezifikation +
ADR-0040 + Index), `3c903a0` (**Engine + Tests**, der Kern), `70f98ee` (Handbuch + CHANGELOG),
`3c624d8` (Plan-Fortschreibung + Register-Beleg). Stand HEAD `3c624d8`, Arbeitsbaum sauber.

**Skill:** `.harness/skills/reviewer.md` @ `60e7b66` (Stand des Laufs) · <!-- d-check:ignore -->
**Modell:** unbekannt (Subagent) — die Lauf-Umgebung nennt `deepseek-v4.1-flash:cloud[1m]` ·
**Datum:** 2026-09-19

**Zitier-Form:** Kennung statt Adresse (`slice-194`, `make <target>`, eine Baseline-Stelle als
`v6.6.0` · `regelwerk/<datei>.md` §<Abschnitt>) — dieser Report friert ein. Das `pfad`-Feld hält
den Stand des Laufs fest und darf das.

**Eingangs-Kontext** (die Verträge, gegen die geprüft wurde):

- `docs/plan/planning/in-progress/slice-194-portscope-richtungssegment.md`
- `docs/plan/adr/0040-portscope-richtungssegment.md`, `0036`, `0029`, `0031`, `0032` (aktive)
- `spec/lastenheft.md` §AC-FA-RULE-010, §AC-FA-RULE-008 · `spec/spezifikation.md` §SPEC-RULE-001,
  §SPEC-CLI-001 und deren `Historie`
- `AGENTS.md` §3.1–§3.7, §4, §5 · `harness/conventions.md` §Modus-Deklaration
- `internal/hexagon/core/rules.go`, `internal/cli/cli.go` und die zugehörigen Tests
- Lab-Fixture des Konsumenten-Repos (`lab/examples/go`, fremd, nur gelesen) — Baumkopien mit
  **einem** injizierten Fremd-Slice-Import
- Beobachtungs-Register: `BEO-HARNESS/zwei-regeln-machen-einander-unmoeglich` (samt Belegen),
  `BEO-PLAN/messung-ohne-reproduzierbares-instrument`, `BEO-KERN/dirvocab-portfor-auserander`

**Instrumente:** `ghcr.io/pt9912/a-check:v0.19.0` (vorher) und das lokal gebaute Dev-Image aus
HEAD (nachher), je `docker run --network none` gegen eine read-only gemountete Baumkopie;
Mutationen in einer **Kopie** des Repos (`make test`/`make build` mit eigenem `IMAGE`), nie im
geprüften Baum.

---

## Findings

### F-1 — Beide Leitplanken von `scopeFor` sind ohne roten Beleg; §7 behauptet dennoch eine Probe

- `kategorie`: HIGH
- `quelle`: `AGENTS.md` §5 (Mess-Regel *Eine Mutations-Probe belegt erst, wenn sie rot war*) ·
  ADR-0040 §Entscheidung 1 / §Konsequenzen / §Fitness Function
- `pfad`: `internal/hexagon/core/rules.go:602,606` ·
  `docs/plan/planning/in-progress/slice-194-portscope-richtungssegment.md:186-190` ·
  `docs/plan/adr/0040-portscope-richtungssegment.md` (§Fitness Function, „je eine
  Mutations-Probe")
- `befund`: In einer Arbeitsbaum-Kopie bleibt `make test` **grün** (Exit 0), wenn
  `!nestedInAppTree(m, prefix)` oder `!appTreeContains(m, marker)` aus der Bedingung entfernt
  wird — kein Test hält die beiden Leitplanken, auf denen die ADR ihre zentrale Zusage stützt
  („der Scope kann dadurch nie weiter werden als zuvor"). Die Wirkung der ersten ist mit dem
  mutierten Image messbar: die Geschwister-Form (`hex/ports/inbound/**` mit `direction`, app-Globs
  `hex/services/**` und `other/svc/**`), der AC-FA-RULE-010 Inertheit zusagt, liefert dort
  `other/svc/x.go:3: port-locality: app außerhalb Port-Scope hex`, Exit 1 — der geprüfte Stand
  liefert für dieselbe Config 0 Befunde, Exit 0. §7 nennt für genau diesen Aufweichungs-Fall eine
  „Mutations-Probe"; die einzige belegte Probe (Rücknahme des Schnitts) ist rot und ihre Meldung
  stimmt wörtlich mit dem Register-Beleg überein.
- `verifizierbar`: ja — `make test` in einer Kopie mit je einer entfernten Bedingung (zweimal
  Exit 0) und ein Container-Lauf der mutierten Fassung gegen die Geschwister-Config (Exit 1 gegen
  Exit 0).
- `klasse`: Leitplanke ohne roten Beleg, Probe dennoch behauptet

### F-2 — „dritte Advisory-Diagnose" ist die vierte; die Menge hat bereits drei

- `kategorie`: HIGH
- `quelle`: eigene Nachzählung gegen `internal/cli/cli.go` und `spec/spezifikation.md`
  §SPEC-CLI-001
- `pfad`: `internal/cli/cli.go:193-197` („both advisory diagnoses", `ADR-0029, ADR-0031`) ·
  `docs/plan/adr/0040-portscope-richtungssegment.md:10,13,65,96` · `spec/spezifikation.md:568` ·
  `CHANGELOG.md:26` · Plan `:112` und `:218` ·
  `docs/plan/planning/observations/BEO-HARNESS/zwei-regeln-machen-einander-unmoeglich/evidence/slice-194.md:34`
- `befund`: `cli.go` trägt **vier** Advisory-Ausgaben (Abdeckung, Grenzen, Auflösung, Port-Scope),
  und §SPEC-CLI-001 führt **drei** Diagnose-Absätze; die neue ist die vierte. Die Aufzählung
  `ADR-0029`/`ADR-0031` überspringt die Auflösungs-Diagnose (ADR-0032) an allen genannten Stellen
  — auch in der Aussage, sie sei „die dritte ihrer Art". Die Zahl steht in einer `Accepted`-ADR
  und ist dort ohne Nachfolge-ADR nicht mehr korrigierbar.
- `verifizierbar`: ja — `grep -c 'func write[A-Za-z]*Notice' internal/cli/cli.go` gegen die drei
  Diagnose-Absätze in §SPEC-CLI-001 und gegen die ADR-Aufzählung.
- `klasse`: Repo-eigene Zählung behauptet statt nachgezählt

### F-3 — Die neue Diagnose hat keinen Eintrag in der Sektion, die ihre Vorgänger führen

- `kategorie`: MEDIUM
- `quelle`: ADR-0040 Kopf (`Schärft: … SPEC-CLI-001`) · `AGENTS.md` §5 (Querverweis-Pflicht der
  Historie-Zeile)
- `pfad`: `spec/spezifikation.md:352-430` (Sektion ohne Port-Scope-Absatz) · `:568`
  (Historie-Zeile verlinkt dorthin) · `docs/user/benutzerhandbuch.md:64,83,104` (§2 nennt drei
  Hinweise namentlich)
- `befund`: Der Kopf der ADR nennt `SPEC-CLI-001` als geschärft, und die Historie-Zeile verlinkt für
  die neue Diagnose dorthin — die Sektion beschreibt aber nur die drei älteren Diagnosen; der neue
  Vertrag (Auslöser, Deckel, Schweigen) steht allein in der ADR und im Code. Dasselbe Bild im
  Handbuch: §2 führt die drei älteren Hinweise mit Beispiel aus, den neuen nicht.
- `verifizierbar`: ja — Sichtprüfung der Sektion gegen ihre drei Diagnose-Absätze; `make doc-check`
  fängt es nicht, weil der Link auflöst.
- `klasse`: Neue Diagnose ohne Eintrag in ihrer verbindlichen Sektion

### F-4 — Register: zweiter Beleg committet, die `Stand:`-Zeile bleibt bei 1×

- `kategorie`: MEDIUM
- `quelle`: `AGENTS.md` §5 (Beobachtungs-Register) · Plan §8/§9
- `pfad`:
  `docs/plan/planning/observations/BEO-HARNESS/zwei-regeln-machen-einander-unmoeglich/state.md:1`
  gegen `…/evidence/` (zwei Dateien) · Plan `:221-229`
- `befund`: Das Verzeichnis trägt zwei Belege (`slice-192.md`, `slice-194.md`), die `Stand:`-Zeile
  sagt `offen (1×)`; von den 18 Register-Einträgen mit genannter Zahl ist dieser der **einzige**,
  dessen Ziffer von seiner Belegliste abweicht (zweiter Zähler: Verzeichnis-Inhalt). Plan §8 nennt
  denselben Zähler mit 2×.
- `verifizierbar`: ja — `Stand:`-Ziffer gegen `ls evidence/`; `make verify-observations` prüft
  Deckung, nicht die Ziffer, und bleibt deshalb grün.
- `klasse`: Register-Stand-Zeile nicht gegen die Belegliste nachgezogen

### F-5 — Die No-Bump-Begründung lässt die nächste Gegenlesart des Lastenhefts aus

- `kategorie`: MEDIUM
- `quelle`: `spec/lastenheft.md` §AC-FA-RULE-010 (Beschreibung: „Der Scope ist
  **pfad-abgeleitet** (keine Deklaration)"; Out-of-Scope: „eine erzwungene explizite
  Scope-Deklaration in der Config") · Plan §5 (Rückführungs-Trigger zu genau dieser Frage)
- `pfad`: `docs/plan/adr/0040-portscope-richtungssegment.md` §Entscheidung 3
- `befund`: Die Ableitung verzweigt jetzt auf einen **deklarierten** Wert, während AC-FA-RULE-010
  im selben Absatz „keine Deklaration" führt und eine explizite Scope-Deklaration ausschließt; für
  „kein Lastenheft-Bump" führt die ADR nur das Satzstück „der Verzeichnis-Teilbaum, der seinen
  Port-Ordner besitzt" an und geht auf die beiden anderen Stellen nicht ein.
- `verifizierbar`: ja — Wortlaut von §AC-FA-RULE-010 gegen §Entscheidung 3 der ADR.
- `klasse`: No-Bump-Begründung lässt die nächste Gegenlesart der höheren Quelle aus

### F-6 — Der Deckel der neuen Diagnose hat keinen Test, der der Vorgänger schon

- `kategorie`: LOW
- `quelle`: ADR-0040 §Konsequenzen („Deckel **mit Restzahl**" als geerbte Hausregel)
- `pfad`: `internal/cli/cli.go:152-154` · `internal/cli/cli_test.go` (kein Fall mit mehr als zehn
  Globs; die Deckel-Fälle der Vorgänger stehen bei `:1053` und `:1283`)
- `befund`: `writePortScopeNotice` liegt im Coverage-Bericht des Gates bei **77.8 %**, die drei
  Vorgänger bei 100 % — der Zweig „… und N weitere" wird von keinem Test ausgeführt. Von Hand
  ausgeführt greift er (12 inerte Globs → 10 Zeilen plus „… und 2 weitere", Exit 0).
- `verifizierbar`: ja — Coverage-Zeile des `coverage-gate`-Laufs; Container-Lauf mit zwölf inerten
  Globs.
- `klasse`: Diagnose-Deckel ohne Test, anders als bei den Vorgängern

### F-7 — §7 verweist für die rote Meldung auf §3, wo keine steht

- `kategorie`: LOW
- `quelle`: `AGENTS.md` §5 (Mess-Regel 2: der Beleg nennt die Meldung der roten Probe)
- `pfad`: Plan `:199-200` („ihre rote Meldung benannt (§3)") gegen `:113`
- `befund`: §3 fordert für beide Test-Dateien nur „je mit Mutations-Probe"; die Meldung der roten
  Probe steht allein im Register-Beleg. Die Sache ist belegt, der Zeiger trifft sie nicht.
- `verifizierbar`: ja — beide Stellen im selben Dokument.
- `klasse`: Querverweis auf eine Stelle, die die genannte Meldung nicht trägt

### F-8 — Zwei Fixture-Kommentare begründen sich über den abwesenden Zustand

- `kategorie`: LOW
- `quelle`: `AGENTS.md` §3.7 (Kommentar-Klassen: Zusage · Kopplung · Abgrenzung · Rang-Zeiger ·
  Grenze)
- `pfad`: `internal/hexagon/core/rules_test.go:1503` („the one that used to switch port-locality
  off silently") · `internal/cli/cli_test.go:1479` („the one that used to keep port-locality
  silent")
- `befund`: Beide Kommentare nennen neben der Kopplung (Handbuch §4/ADR-0040) den **früheren**
  Zustand als Begründung — die Klasse, die §3.7 als nicht tragend benennt („früher stand hier …");
  die vorige Fassung hält `git`. Kein Gate fängt das.
- `verifizierbar`: ja — Wortlaut gegen §3.7; die Kommentare stehen in den geänderten Dateien.
- `klasse`: Kommentar nennt den abwesenden früheren Zustand

### F-9 — Der Geschwister-Test erreicht den neuen Zweig nicht

- `kategorie`: LOW
- `quelle`: `spec/lastenheft.md` §AC-FA-RULE-010 (Inertheit der Geschwister-Ports) · Plan §7
  Risiko 1 (dessen Ausgang sich auf genau diesen Test stützt)
- `pfad`: `internal/hexagon/core/rules_test.go:1476-1489` (`classicModel`) und
  `TestInertPortScopesSiblingPortsSilent`
- `befund`: `classicModel` deklariert an beiden Port-Schichten **keine** `direction`; in `scopeFor`
  greift damit der Kurzschluss `dir == ""`, bevor irgendeine der neuen Bedingungen geprüft wird.
  Der Test hält die Geschwister-Inertheit über `nestedInAppTree` allein und wird auch dann grün,
  wenn die Richtungslogik fehlt (in der Kopie gemessen); eine Config mit Richtung an der
  Port-Schicht erreicht er nicht (Wirkung siehe F-1).
- `verifizierbar`: ja — Entfernen des Schnitts in einer Kopie: der Test bleibt grün, `make test`
  wird erst durch `TestPortLocalityDirectionSegmentKeepsRule` rot.
- `klasse`: Test-Fixture erreicht den neuen Zweig nicht

### F-10 — Die v2-Zeile nennt eine Config-Form, die ihre Zahl nicht trägt

- `kategorie`: LOW
- `quelle`: `AGENTS.md` §5 (Mess-Regel 1: Geltungsbereich und Parameter einer Messung)
- `pfad`: Plan §2, Zeile `:92` (Config-Spalte `…/ports/{inbound,outbound}/**`) gegen den Nachbau
- `befund`: Die Config-Spalte schreibt eine Brace-Form, die als Glob nicht literal ist, und lässt
  die Schicht-Aufteilung offen; die behauptete Zahl der genannten Globs („beide") hängt daran. Mit
  per-Slice-Port-Schichten nennt der Hinweis **vier** Globs samt abgeleitetem Scope (Exit 0), nicht
  zwei.
- `verifizierbar`: ja — Container-Lauf mit der nachgebauten Variante.
- `klasse`: Zahl ohne die Parameter, von denen sie abhängt

## Negativbefunde

- geprüft, ohne Befund: **Vorher/Nachher-Tabelle §2** — alle vier Zeilen selbst nachgerechnet
  (v1, v2, v3-Kontrolle, v4) gegen `v0.19.0` und gegen den geprüften Stand; v1 unverändert
  `port-locality: 1`/Exit 1, v3 `lateral-slice: 1`/Exit 1 auf **beiden** Ständen, v4 vorher still
  und nachher `port-locality: 1`/Exit 1 mit demselben Scope (`…/order/createorder`) und derselben
  Meldung wie v1 — deckungsgleich mit dem Plan (Grenze: eigener Nachbau der Config, siehe F-10).
- geprüft, ohne Befund: **Advisory-Eigenschaften** — Hinweis ausschließlich auf stderr (stdout
  0 Byte), Exit-Code 0 auch ohne Befund, Text nach der Zusammenfassung, Deckel mit Restzahl,
  „Abhilfe"-Zeile; Schweigen bei sauberer Config (Repo-eigener Lauf, v4 mit deklarierter Richtung).
- geprüft, ohne Befund: **Schnitt-Reihenfolge** — Spezifikation (§SPEC-RULE-001, Regel-Tabelle und
  Ableitungs-Absatz), ADR (Entscheidung 1), Code (`portScope`/`scopeFor`) und Handbuch §3.7/§4
  sagen übereinstimmend „unter dem Port-Ordner (`…/ports/outbound/**`)"; keine Reste der alten
  Lesart, kein Beleg für das Richtungssegment *über* dem Port-Ordner.
- geprüft, ohne Befund: **Referenz-Richtung / `matrix`** — `make doc-check` Exit 0; kein ADR-Verweis
  in einem Spec-Stratum außerhalb der `Historie`, ADR-0040 ohne `slice-\d{3}`-Token im Körper,
  Index-Zeile vorhanden.
- geprüft, ohne Befund: **Gate-Schwellen** — kein Gate-, Pin- oder Konfigurationsfile im Umfang der
  sechs Commits; keine Schwellen-Senkung (§3.6 nicht berührt).
- geprüft, ohne Befund: **Traceability und Lifecycle** — `make trace-check`,
  `make commit-scope-check` (6 `(planning)`-Commits) und `make doc-immutable` über die Slice-Range:
  Exit 0.
- geprüft, ohne Befund: **Register-Deckung** — `make verify-observations` Exit 0, der neue Beleg
  existiert und ist nicht leer (die *Ziffer* der `Stand:`-Zeile ist F-4).
- geprüft, ohne Befund: **Kommentar-Klassen der neuen Stellen** — `InertPortScopes`, `scopeFor`
  (beide Leitplanken-Begründungen als Grenze), `nestedInAppTree`, `writePortScopeNotice` tragen je
  eine der fünf Klassen; die zwei Ausnahmen sind F-8.
- geprüft, ohne Befund: **Risiko-Ausgänge §7 der Form nach** — fünf Risiken, fünf Ausgänge aus der
  geschlossenen Dreier-Menge, jeder mit Begründung (`make verify-risiko-ausgaenge` Exit 0). Die
  *inhaltliche* Tragfähigkeit ist teils bestritten: Risiko 3 → F-1, Risiko 1 → F-9.
- **Nicht** Gegenstand dieses Laufs: DoD- und Spec-Konformität (Modul 11, Verifikation), die
  Wirkung des Slice auf das **fremde** Konsumenten-Repo, und die Frage, ob `port-direction-mismatch`
  für driven-Adapter je feuern kann (im Plan §1 ausgeschlossen).

## Gate-Läufe (dieser Lauf)

| Lauf | Exit | Geltungsbereich |
|---|---|---|
| `make gates` | 0 | aggregiert über HEAD `3c624d8`, Arbeitsbaum sauber: u. a. `lint`, `test`, `coverage-gate` (**96.10 %** ≥ 90 %), `arch-check` (0 Befunde), `doc-check` (580 Dateien, 0 Befunde), `doc-targets`, `doc-planning`, `doc-workflows`, `doc-reviews`, `doc-mentions`, `gate-consistency`, `version-coherence`, `suppression-check`, `symlink-check`, `dcheck-phrase-selftest`, `guard-selftest`, `ci-range-selftest`, `record-gates` |
| `make verify` | 0 | Verifikations-Schicht: `verify-risiko-ausgaenge`, `verify-observations`, `doc-structure`, `doc-complete` (21 Anforderungen, 0 Waisen) |
| `make trace-check RANGE=f0fe0a6~1..3c624d8` | 0 | Commit-Kennungen der Slice-Range (Modul `commits`) |
| `make commit-scope-check RANGE=f0fe0a6~1..3c624d8` | 0 | 6 `(planning)`-Commits, ausschließlich `docs/plan/planning/` |
| `make doc-immutable RANGE=f0fe0a6~1..3c624d8` | 0 | ADR-Immutabilität über die Slice-Range (Modul `vcs`) |
| `make test IMAGE=<eigene>` in einer Arbeitsbaum-Kopie, Schnitt zurückgenommen | **2** | Rot mit der im Register-Beleg genannten Meldung `ein Port-Glob mit Richtungssegment darf port-locality nicht abschalten, got []` (dazu `TestInertPortScopesResolvedConfigSilent`, `TestPortScopeNoticeSilentWhenDirectionDeclared`) |
| `make test IMAGE=<eigene>` in einer Kopie, Leitplanke 1 entfernt | 0 | **Kein** Test rot (F-1) |
| `make test IMAGE=<eigene>` in einer Kopie, Leitplanke 2 entfernt | 0 | **Kein** Test rot (F-1) |
| Container-Läufe (netzlos, read-only) | 0 / 1 | v1–v4 gegen `v0.19.0` und den geprüften Stand; Zwölf-Glob-Lauf (Deckel, Exit 0); Geschwister-Config gegen den geprüften Stand (Exit 0) und gegen die mutierte Fassung (Exit 1, F-1) |

**Geltungsbereich dieser Läufe:** Go-Fixture Pfad `app`-Importeur → slice-lokaler Port über ein
literales Port-Glob-Präfix; andere Sprach-Backends und andere Auflösungs-Modi wurden nicht
gemessen. Die Mutationsläufe liefen gegen eine **Kopie** des Repos, nicht gegen den geprüften
Baum — dieser blieb unverändert.

## Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 2 |
| MEDIUM | 3 |
| LOW | 5 |
| INFO | 0 |

**Finding-Klassen dieses Laufs:** Leitplanke ohne roten Beleg, Probe dennoch behauptet ·
Repo-eigene Zählung behauptet statt nachgezählt · Neue Diagnose ohne Eintrag in ihrer verbindlichen
Sektion · Register-Stand-Zeile nicht gegen die Belegliste nachgezogen · No-Bump-Begründung lässt
die nächste Gegenlesart der höheren Quelle aus · Diagnose-Deckel ohne Test, anders als bei den
Vorgängern · Querverweis auf eine Stelle, die die genannte Meldung nicht trägt · Kommentar nennt den
abwesenden früheren Zustand · Test-Fixture erreicht den neuen Zweig nicht · Zahl ohne die
Parameter, von denen sie abhängt

## Verdikt

**Merge-blockierend:** ja — F-1 und F-2 sind HIGH. F-2 sitzt in einer `Accepted`-ADR und ist dort
nur über eine Nachfolge-ADR zu korrigieren; F-1 entzieht dem Ausgang von Risiko 3 in §7 die
Grundlage und lässt die in der ADR als „teurerer Fehler" benannte Aufweichung ohne roten Beleg.

**Übergabe:** Findings gehen an den Implementer; F-4 und F-7 sind vor der Closure zu schließen, F-3
berührt Spezifikation und Handbuch, F-5 die ADR-Begründung. Die **Finding-Klassen** gehen in die
Slice-Closure §7 und von dort in den Zähler — die Klasse zu F-2 ist mit der Formulierung aus
`slice-166` identisch gehalten (*Repo-eigene Zählung behauptet statt nachgezählt*). Dieser Report
ist ein **Lauf-Beleg** und ersetzt keine Verifikation (Modul 11; anderer Eingabe-Kontext).
