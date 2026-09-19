# slice-198 — Der lokale Pre-Flight deckt die Workflow-Schritte

**Welle:** ohne Welle.

**Bezug:** Beim Release `v0.20.0` (2026-09-19) war `make ci` lokal grün und der CI-Lauf auf `main`
rot. [`AC-QA-02`](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze).

**Berührte Spec-Stellen:** — (der Slice berührt kein Spec-Stratum; er betrifft den Gate-Bestand).

**Verantwortlich:** Claude — gesetzt beim Übergang nach `in-progress/`.

**Autor:** Claude, im Auftrag des Maintainers. **Datum:** 2026-09-19.

**Lerneintrag — Form:** wird bei Closure benannt.

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
  löst die Range in der CI auf (inkl. Force-Push-Fall); lokal ist sie `origin/main..HEAD`, und das
  ist genau das, was der nächste Push prüfen wird.

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

- [ ] `make preflight` existiert, fährt `ci` **und** die drei Range-Schritte, und steht im
      Gate-Index mit seiner Grenze.
- [ ] [`releasing.md`](../../../../docs/user/releasing.md) nennt es an der Stelle, an der bisher
      `make ci` stand.

- [ ] Unabhängiger Review durchgeführt (Report unter [`docs/reviews/`](../../../reviews/README.md)).
- [ ] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [ ] Beobachtungs-Register fortgeschrieben.
- [ ] Jedes Risiko aus §7 trägt einen Ausgang.

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
  es erweitert es — und es sieht nur den lokalen `git`-Stand. — **Ausgang:** <offen bis Closure>
- **Ein grüner Pre-Flight auf einem veralteten `origin/main`.** Wer nicht `git fetch`t, prüft eine
  falsche Range. — **Ausgang:** <offen bis Closure>
- **Die Klasse dahinter hat jetzt drei Vorfälle.** `BEO-GATE/preflight-deckt-den-ci-schritt-nicht`
  entsteht mit diesem Slice; ob daraus ein Sensor wird („jeder `make`-Aufruf im Workflow ist in
  einem Aggregat oder deklariert"), entscheidet der Lese-Schritt bei Closure. — **Ausgang:** <offen
  bis Closure>

## 8. Closure-Notiz

*(bei Closure auszufüllen)*

## 9. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** berührt sind `GATE` (Makefile, Workflow, Gate-Index) und
`HARNESS` (`releasing.md` liegt unter `docs/user/` — das ist die Sub-Area `USER`, seit slice-195
deklariert).

**Vorgelagert — offene Beobachtungen sichten:** Register über `GATE`, `HARNESS`, `USER` gelesen —
[`BEO-GATE/ungelaufene-mechanik-docker-hub-spiegel`](../observations/BEO-GATE/ungelaufene-mechanik-docker-hub-spiegel/observation.md)
und [`BEO-GATE/weiche-haengt-an-fremdem-event-feld`](../observations/BEO-GATE/weiche-haengt-an-fremdem-event-feld/observation.md)
betreffen CI-Mechanik, aber einen anderen Gegenstand; zu „Pre-Flight ≠ CI-Schritt" kein Treffer —
der Eintrag entsteht mit diesem Slice.

**Modus-Begründungsblock:** alle berührten Sub-Areas GF. Kein Block je Sub-Area nötig.
