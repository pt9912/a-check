# slice-198 — Der lokale Pre-Flight deckt die Workflow-Schritte

**Welle:** ohne Welle.

**Bezug:** Beim Release `v0.20.0` (2026-09-19) war `make ci` lokal grün und der CI-Lauf auf `main`
rot. [`AC-QA-02`](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze).

**Berührte Spec-Stellen:** — (der Slice berührt kein Spec-Stratum; er betrifft den Gate-Bestand).

**Verantwortlich:** Claude — gesetzt beim Übergang nach `in-progress/`.

**Autor:** Claude, im Auftrag des Maintainers. **Datum:** 2026-09-19.

**Lerneintrag — Form:** geschärfte Regel.

---

## 1. Ziel und Abgrenzung

**Ziel:** Wer vor dem Push einen lokalen Pre-Flight fährt, fährt **das, was die CI auf diesem Push
fahren wird** — ein grüner Pre-Flight und ein roter CI-Lauf schließen einander damit aus.

**Ausgangslage (gemessen):** `make ci` ist `gates` + `image-test`. Der Workflow
[`.github/workflows/ci.yml`](../../../../.github/workflows/ci.yml) fährt **darüber hinaus** drei
Schritte als eigenen Block über die **Push-Range**: `make trace-check`, `make
commit-scope-check` und `make doc-immutable`. Am 2026-09-19 war `make ci` lokal grün; der CI-Lauf
desselben Stands war rot — `doc-immutable` meldete drei Kern-Drift-Befunde an `Accepted`-ADRs, die
in der gepushten Range lagen.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **`make ci` selbst erweitern.** *Schicht-Abgrenzung:* `ci` läuft auch in der **Release**-Pipeline
  auf einem Tag aus; dort wäre „die Push-Range" die **Release**-Range, und `doc-immutable` über sie
  ist rot (drei historische Befunde, deklariert in [`MR-024`](../../../../harness/conventions.md#mr-024)).
  Eine Erweiterung von `ci` machte jedes Release rot.
- **Den Sensor ändern** (`vcs`-Modul, `exempt-paths`). *Es wäre ein anderer Vorgang:* Gegenstand von
  [slice-197](../done/wellenlos/slice-197-adr-kern-drift.md) und [`MR-024`](../../../../harness/conventions.md#mr-024).
- **Die Range-Bestimmung neu erfinden.** *Bestand bleibt bewusst stehen:* `tools/ci-commit-range.sh`
  löst die Range in der CI auf (inkl. Force-Push-Fall); lokal ist sie `origin/main..HEAD` — gleich
  oder größer als das, was der nächste Push enthält.

## 2. Ausgangsmessung

| Messung | Bau | Ergebnis |
|---|---|---|
| Was `make ci` fährt | `grep -n "^ci:" Makefile` | `gates` + `image-test` |
| Was der Workflow zusätzlich fährt | `grep -nE "make (trace-check\|commit-scope-check\|doc-immutable)" .github/workflows/ci.yml` | drei Schritte über `RANGE` |
| Der belegte Fall | CI-Lauf `35456322975`, Schritt „Traceability + ADR-Immutabilität" | **Exit 2**, drei Befunde |

**Geltungsbereich:** die Schritte des `ci`-Workflows und des `ci`-Targets. Der `release`-Workflow
fährt nur `make ci` und ist **nicht** betroffen — gemessen am selben Tag (Lauf `35456335601`, grün).

## 3. Umsetzung

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| [`Makefile`](../../../../Makefile) | update | ein Target `preflight`: `ci` **plus** die drei Range-Schritte über `origin/main..HEAD` |
| [`harness/README.md`](../../../../harness/README.md) | update | Zeile im Gate-Index (§Sensors), mit Bindung und Grenze |
| [`docs/user/releasing.md`](../../../../docs/user/releasing.md) | update | Schritt „vor dem Tag": `make preflight` statt `make ci` — mit dem Grund |
| [`harness/conventions.md`](../../../../harness/conventions.md) oder der Guide in [`AGENTS.md`](../../../../AGENTS.md) | update | falls der dritte Vorfall einen Ausgang verlangt (siehe §7) |

**Die Grenze des neuen Targets.** Es kann nur so viel sehen wie das lokale `git`: Ist `origin/main`
nicht geholt, ist die Range zu groß oder falsch. Das Target nennt das, statt es zu verschweigen —
dieselbe Ehrlichkeit wie bei `tools/ci-commit-range.sh`.

**Auszuführende Gates:** `make doc-targets` (Index ↔ Makefile, in beiden Richtungen), `make gates`,
zum Abschluss `make verify`.

## 4. Definition of Done

- [x] `make preflight` existiert, fährt `ci` **und** die drei Range-Schritte, und steht im
      Gate-Index mit seiner Grenze. Die Proben: `PREFLIGHT_RANGE=89fc7dc..ed7a3d8` → **Exit 2**
      mit drei `core-drift-vcs`-Meldungen; `PREFLIGHT_RANGE=ed7a3d8..06c7876` → **Exit 0** über
      vier Commits. Eine **leere** Range meldet der Lauf als WARNUNG, weil sie nichts prüft.
- [x] [`releasing.md`](../../../../docs/user/releasing.md) nennt es an der Stelle, an der bisher
      `make ci` stand — im Beleg-Slot von Item 5.

- [x] Unabhängiger Review durchgeführt (Report unter [`docs/reviews/`](../../../reviews/README.md)).
- [x] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [x] Beobachtungs-Register fortgeschrieben.
- [x] Jedes Risiko aus §7 trägt einen Ausgang.

`make gates` und `make verify` grün.

## 5. Trigger

**Start** (`open` → `in-progress`): das WIP-Limit ist frei.

**Rückführungen — vorab benannt:**

- `in-progress` → `next` (zu groß): Erweist sich die Range-Bestimmung als eigenes Problem (der
  Verweis auf `origin/main` ist in einem flachen Klon falsch), sind Target und Range-Auflösung zwei
  Vorgänge.
- `in-progress` → `open` (blockiert): Verlangt `harness/README.md` §Sensors für ein Nicht-Aggregat
  eine eigene Sensor-Datei, ist das ein Schreibvorgang und keine Entscheidung dieses Slice.

## 6. Closure-Trigger

DoD vollständig, `make gates` und `make verify` grün, Closure-Notiz mit Lerneintrag — und
`make preflight` einmal **rot** gefahren gegen den Stand, der die CI rot machte, als Beleg, dass es
den Gegenstand trifft.

## 7. Risiken und offene Punkte

- **Das Target wird für den Pre-Flight gehalten, den es nicht leistet.** Es ersetzt `make ci` nicht,
  es erweitert es — und es sieht nur den lokalen `git`-Stand. — **Ausgang:** *entfallen*, gestrichen
  mit Begründung: Die Grenze steht im Gate-Index **und** im Rezept; der Pre-Flight führt `ci` als
  Voraussetzung, er kann es also gar nicht ersetzen.
- **Ein grüner Pre-Flight auf einem veralteten `origin/main`.** Wer nicht `git fetch`t, prüft eine
  falsche Range. — **Ausgang:** *entfallen*, gestrichen mit Begründung: Fällt `origin/main` nicht
  auf, bricht das Target mit Exit 2 und **nennt** die Abhilfe (`git fetch origin`), statt eine
  Range zu raten. Nicht gedeckt bleibt der Fall „`origin/main` ist da, aber alt" — benannt, nicht
  verschwiegen.
- **Die Klasse dahinter ist bei 1×, nicht bei 3×.** `BEO-GATE/preflight-deckt-den-ci-schritt-nicht`
  entsteht mit diesem Slice; ein Sensor („jeder `make`-Aufruf im Workflow ist in einem Aggregat
  oder deklariert") wäre prüfbar, ist aber an der Schwelle noch nicht fällig. — **Ausgang:**
  *weiter offen* → **Beobachtungs-Register**
  ([`BEO-GATE/preflight-deckt-den-ci-schritt-nicht`](../observations/BEO-GATE/preflight-deckt-den-ci-schritt-nicht/observation.md)).

## 8. Closure-Notiz

**Lerneintrag — Form: geschärfte Regel.** *„Was die CI fährt" ist nicht dasselbe wie ein
Aggregat-Target, das so heißt.* `make ci` deckt `gates` + `image-test`; der Workflow fährt darüber
hinaus drei Schritte über die **Commit-Range**. **Weil** beide Mengen für sich korrekt deklariert
waren, nannte keine von beiden die Differenz — und ein Pre-Flight, der sie nicht kennt, belegt mehr,
als er deckt. Die Regel steht jetzt als `make preflight` samt Grenze im Gate-Index und als
Beleg-Slot in der Freigabe-Checkliste.

**Der zweite Lerneintrag ist der teurere und kam aus der Sache selbst.** Die drei Sensoren des Repos
haben den neuen Bestand **Schritt für Schritt** nachgezogen: `doc-targets` verlangte die
Index-Zeile, `gate-consistency` den `.PHONY`-Eintrag („eine gleichnamige Datei ließe make das
Rezept überspringen und Exit 0 melden"), `guard-selftest` den Eintrag in der GATES-Liste des
Command-Guard. Drei Fehlschläge, drei benannte Ursachen, kein Rätselraten — der Bestand hat den
Beitragenden geführt.

**Ein Befund des Reviews, der eine Zusage dieses Slice widerlegt hat.** Die DoD nannte für die
**grüne** Probe nur „die aktuelle Range" — und die war am Tag des Reviews **leer** (`origin/main`
== `HEAD`, nachdem der Maintainer gepusht hatte): Der Lauf meldete grün über **null** Commits.
Genau die Klasse der ersten Mess-Regel ([`AGENTS.md`](../../../../AGENTS.md) §5): ein Beleg, der
seinen Geltungsbereich nicht nennt, deckt seinen Gegenstand nicht. Das Rezept **zählt** jetzt die
Commits und meldet eine leere Range als WARNUNG; die Checkliste nennt den Fall.

**Steering-Loop-Eintrag:** gezählt, nicht verkörpert.
[`BEO-GATE/preflight-deckt-den-ci-schritt-nicht`](../observations/BEO-GATE/preflight-deckt-den-ci-schritt-nicht/observation.md)
ist **neu angelegt** (1×); die Schwelle ist nicht erreicht, ein Ausgang nicht fällig.

**Beobachtungs-Register ([`../observations/`](../observations/README.md)):** ein Verzeichnis **neu
angelegt** — `BEO-GATE/preflight-deckt-den-ci-schritt-nicht/`, Beleg `evidence/slice-198.md`, Zähler
**1×**.

**Folge-Slices:** keine.

**Risiken aus §7:** alle drei *entfallen* oder sind weitergegeben — siehe dort.

**Drei Paarungen:** Anker — kein `liegt in`-Feld gesetzt, weil nichts verkörpert wurde · Folge-Slice
— keiner genannt · Register getragen (der genannte Pfad existiert und trägt einen Beleg).

## 9. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** berührt sind `GATE` (Makefile, Workflow, Gate-Index,
Command-Guard) und `USER` ([`docs/user/releasing.md`](../../../../docs/user/releasing.md),
seit slice-195 deklariert).

**Vorgelagert — offene Beobachtungen sichten:** Register über `GATE`, `HARNESS`, `USER` gelesen —
[`BEO-GATE/ungelaufene-mechanik-docker-hub-spiegel`](../observations/BEO-GATE/ungelaufene-mechanik-docker-hub-spiegel/observation.md)
und [`BEO-GATE/weiche-haengt-an-fremdem-event-feld`](../observations/BEO-GATE/weiche-haengt-an-fremdem-event-feld/observation.md)
betreffen CI-Mechanik, aber einen anderen Gegenstand; zu „Pre-Flight ≠ CI-Schritt" kein Treffer —
der Eintrag entsteht mit diesem Slice.

**Modus-Begründungsblock:** alle berührten Sub-Areas GF. Kein Block je Sub-Area nötig.
