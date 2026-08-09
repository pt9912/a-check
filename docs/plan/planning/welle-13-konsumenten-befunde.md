# Welle welle-13-konsumenten-befunde: Was ein realer Einsatz gefunden hat

**Lifecycle:** Die aktive Welle liegt flach unter `docs/plan/planning/`; bei Closure wandert diese
Datei per `git mv` nach `done/` (neben ihre `welle-13-results.md`). Der Zustand ist die
Verzeichnis-Position — kein Status-Feld. Ob eine flache Welle *aktuell* oder *geplant* ist, sagt
die [Roadmap](in-progress/roadmap.md).

**Erste Welle-Plan-Datei dieses Repos.** Die zwölf Wellen davor existieren nur als
Prosa-Überschriften in der Roadmap; die Grandfather-Klausel in [`README.md`](README.md#ab-wann-das-gilt)
sagt dazu: *„Auch Welle-Plan-Dateien entstehen erst künftig."* Das ist hier eingelöst.

**Zielmeilenstein:** kein Meilenstein-Bezug — die Welle hebt die Qualität eines ausgelieferten
Werkzeugs, ohne einen extern beobachtbaren Zustand zu erreichen.

**Verantwortlich:** Maintainer. **Datum:** 2026-08-09.

---

## 1. Welle-Ziel

**a-check sagt selbst, wo es blind ist — statt dass ein Konsument es nachstellen muss.**

Die Befunde aus einem realen Einsatz in einem Fremd-Repo (2026-08-09) haben dasselbe Muster: das
Werkzeug *kennt* seine Grenzen, aber der Konsument erfährt sie nicht dort, wo er arbeitet. Einmal
kostete das einen **realen Fehlpin**.

Spiegelbar an [AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze):
*„Die Heuristik-Grenzen werden **dokumentiert** statt verschwiegen."* Die Welle prüft, ob dieses
Versprechen am **Repo** eingelöst ist — nicht nur im Lastenheft.

## 2. Trigger (Welle startet)

- **Gefeuert am 2026-08-09:** ein realer Konsumenten-Einsatz meldet Befunde — zuerst vier, dann als
  formale Change Requests sechs —, davon einer mit belegtem Schaden (Fehlpin durch den dokumentierten `--print-mk`-Bump-Weg).
- `welle-12-regelwerk-migration` ist geschlossen ([`done/welle-12-results.md`](done/welle-12-results.md))
  — der Harness trägt jetzt, was die Produkt-Achse an Belegen braucht.

## 3. Closure-Trigger (Welle schließt)

- Alle Slices der Welle in `done/` — die Zahl steht in §4, nicht hier (sie ist am 2026-08-09 von vier auf sechs gewachsen).
- **Jeder Befund ist entweder behoben oder als deklarierte Grenze ausgewiesen** — kein
  Befund bleibt unbeantwortet liegen. Das ist die inhaltliche Bedingung; „alle Slices done" allein
  wäre sie nicht, weil ein Slice auch mit dem Ergebnis „bleibt so" schließen kann.
- Für die spec-first-Slices ([slice-081](in-progress/slice-081-heuristik-diagnose.md),
  [slice-083](done/slice-083-print-mk-digest-selbstbezug.md)): die zugehörige ADR trägt
  `Status: Accepted`, und die Lastenheft-Änderung steht **vor** dem Code.
- `make ci` grün (Exit 0 in eine Datei, getrennt geprüft) — die repo-eigene Ersetzung des
  Baseline-Replay-Laufs nach
  [`MR-008`](../../../harness/conventions.md#mr-008--kein-replay-keine-agenten-telemetrie).
- Closure-Notiz `done/welle-13-results.md` nach der Fünf-Schritt-Prozedur
  ([`README.md`](README.md#wellen-closure-prozedur)).

## 4. Slices in dieser Welle

| Slice | CR | Titel | Bezug |
|---|---|---|---|
| [slice-081](in-progress/slice-081-heuristik-diagnose.md) | `CR-1` | Laufzeit-Diagnose für nicht extrahierte Import-Formen | [AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze) |
| [slice-082](done/slice-082-print-mk-docker-indirektion.md) | `CR-6` | `--print-mk`: `$(DOCKER)` statt wörtlichem `docker` | [AC-FA-DIST-001](../../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk) |
| [slice-083](done/slice-083-print-mk-digest-selbstbezug.md) | `CR-5` | `--print-mk` nennt den Digest des Vorgängers | [AC-QA-03](../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit) |
| [slice-084](done/slice-084-handbuch-heuristik-grenzen.md) | `CR-3` | Heuristik-Grenzen dort, wo Konsumenten lesen | [AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze) |
| [slice-085](open/slice-085-schicht-ohne-aufloesung.md) | `CR-2` | Diagnose: Schicht ohne auflösende Importe | [SPEC-RULE-001](../../../spec/spezifikation.md#spec-rule-001--regel-auswertung) |
| [slice-086](open/slice-086-forbidden-constructs-fail-closed.md) | `CR-4` | `forbidden_constructs` ohne `port`-Rolle: fail-closed | [SPEC-CONF-001](../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema) |

**Die Welle ist am 2026-08-09 von vier auf sechs Slices gewachsen.** Der Konsument reichte die
Befunde als sechs formale Change Requests nach; zwei davon (`CR-2`, `CR-4`) waren in der ersten
Meldung nicht enthalten. Die Umplanung steht im
[Drift-Log](in-progress/roadmap.md#historische-trigger-verschiebungen).

**Reihenfolge aus Konsumentensicht** (aus den CRs übernommen): `CR-1` und `CR-2` zuerst — *„Sie
machen aus einem Handgriff, den ich sechsmal wiederholt habe — Verstoß einbauen, messen, notieren —
eine Eigenschaft des Werkzeugs."* Dann `CR-3`/`CR-4` als Ehrlichkeits-Korrekturen, zuletzt
`CR-5`/`CR-6` als Politur.

**`CR-3` ist mit [slice-084](done/slice-084-handbuch-heuristik-grenzen.md) in der
Dokumentations-Variante erfüllt** — der CR lässt beide zu („alle Direktiven extrahieren … oder,
wenn das nicht gewollt ist, die Grenze dokumentieren wie die relative"), sein Akzeptanzkriterium
nennt aber die Fix-Variante. **Offen: ob die Doku-Variante als Erfüllung gilt.** Fällt der Entscheid
auf den Fix, ist das ein eigener Slice — nicht eine Wiedereröffnung von 084.

**Reihenfolge, wo sie zwingend ist:** [slice-082](done/slice-082-print-mk-docker-indirektion.md)
vor [slice-083](done/slice-083-print-mk-digest-selbstbezug.md) — beide ändern `cli.go` **und**
`a-check.mk`, und die Fragment-Parität in `tools/image-test.sh` erzwingt gemeinsames Wandern.

**Ein Slice trägt eine Verwerfungs-Bedingung:**
[slice-084](done/slice-084-handbuch-heuristik-grenzen.md) entfällt, wenn
[slice-081](in-progress/slice-081-heuristik-diagnose.md) zuerst gebaut wird und seine Diagnose die
betroffenen Formen vollständig meldet. Eine Welle, die jeden gelisteten Slice *garantiert* liefert,
wäre hier die falsche Zusage.

## 5. Abhängigkeiten

- **Wird blockiert von:** nichts. Jeder Slice der Welle trägt den Trigger „sofort" in seinem `§0`.
- **Blockiert:** nichts.
- **Ausdrücklich nicht in dieser Welle, obwohl offen:**
  [slice-079](open/slice-079-gate-consistency-abloesen.md) (Harness-Nachlauf aus `welle-12`,
  ohne Welle) und [slice-080](open/slice-080-verify-abloesung-dcheck.md) (wartet auf ein
  d-check-Release). **Ein Slice, der auf ein Fremdrepo wartet, gehört in keine Welle** — er würde
  ihren Closure-Trigger auf unbestimmte Zeit blockieren. Genau dieser Fehler hielt `slice-046`
  zwei Wochen offen.

## 6. Out-of-Scope für diese Welle

- **Die Heuristik-Grenzen schließen.** Ein AST-Backend ist in
  [AC-FA-EXTRACT-001](../../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion)
  Out-of-Scope. Diese Welle macht Grenzen *sichtbar*, sie verschiebt sie nicht.
- **Ein Release.** `[Unreleased]` ist seit `v0.16.0` leer; ob die Ergebnisse dieser Welle ein
  Release rechtfertigen, entscheidet der Maintainer nach der Closure — nicht die Welle selbst.
- **`F-9`** (Freigabe-Belege aus dem `welle-12`-Review). Berührt
  [`releasing.md`](../../user/releasing.md) wie [slice-083](done/slice-083-print-mk-digest-selbstbezug.md),
  ist aber ein Harness-Thema und nicht aus dem Konsumenten-Einsatz gemeldet.

## 7. Closure-Notiz

_(erst nach Welle-Abschluss füllen — Verweis auf `done/welle-13-results.md`.)_
