# slice-211 — `shapes:` mit `mode: exact` und Opt-in-Befund `shape-unused`

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `git mv` (`make slice-mv`).

**Welle:** welle-16 — [Welle-Plan](welle-16-shapes-sollform.md).

**Bezug:** [AC-FA-RULE-012](../../../../spec/lastenheft.md#ac-fa-rule-012),
[ADR-0041](../../adr/0041-shapes-sollform-je-datei.md) (beide aus slice-209);
[AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml),
[AC-QA-01](../../../../spec/lastenheft.md#ac-qa-01--determinismus).

**Berührte Spec-Stellen:** `spezifikation.md` §[SPEC-CONF-001](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema), §[SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung)
(in slice-209 geschrieben, hier umgesetzt). **Plan-Änderung nach dem Review (2026-10-06):** die
Funde zur Sollform-Datei (Symlink, Pfad-Schreibweisen, Hardlink, einzeilige `shape-unused`-Meldung)
waren Vertragslücken; der Slice präzisiert dafür Spezifikation 0.35.0 — keine Zusage des
Lastenhefts ändert sich.

**Verantwortlich:** Claude — gesetzt beim Übergang nach `next/` (Maintainer: „ja alles machen", 2026-10-06).

**Autor:** Claude. **Datum:** 2026-10-06.

**Lerneintrag — Form:** benannte Spec-Lücke.

---

## 1. Ziel und Abgrenzung

**Ziel:** Zwei Ergänzungen auf dem Normalisierer aus slice-210:
`mode: exact` vergleicht die normalisierte Datei mit einer Sollform-Datei
(`expect:`) und meldet die erste Abweichung als `shape-differs`; ein
`allow`-Eintrag, der in keiner Datei trifft, meldet bei `unused: fail`
`shape-unused` — Exit 1, kein Warn-Level.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Ein Diff über alle Abweichungen.** *Bestand bleibt bewusst stehen*: der CR
  verlangt die **erste** Abweichung; ein vollständiger Diff wäre ein neues
  Ausgabeformat ohne belegten Bedarf.
- **`shape-unused` als Default.** *Bestand bleibt bewusst stehen*: ein neuer
  Befund, der bestehende Konfigurationen rot färbt, wäre Breaking — Opt-in
  hält die Kompatibilitätszusage des CR.
- **Normalisierer-Änderungen.** *Schicht-Abgrenzung*: der Normalisierer ist
  slice-210s Liefer-Punkt; wer hier einen Fehler darin findet, schließt ihn als
  Risiko *eingetreten* mit eigener Kennung, nicht still nebenbei.

  **Plan-Änderung 2026-10-06 (Review G-1, vor dem Code):** Die Dateisuche für
  `files` (aus slice-210, nicht der Normalisierer) folgte einem Symlink im
  literalen Präfix eines Globs aus der Scan-Wurzel hinaus und las fremde
  Dateien. Dieselbe Klasse wie der `expect`-Fund N-1 und vor dem Release zu
  schließen: der Slice nimmt die Regel „kein Bestandteil eines geprüften Pfads
  ist ein Symlink" für `files` mit (Spezifikation 0.35.0, Code, Test).

## 2. Definition of Done

- [x] `mode: exact` mit `expect:` (Exit 2 bei fehlender Sollform-Datei oder
      `allow` neben `exact`); `shape-differs` für eine Anweisung anders,
      zusätzlich oder fehlend, jeweils mit der ersten Abweichung — als Tests.
- [x] `unused: fail` und Befund `shape-unused` (Eintrag ohne Treffer über alle
      Dateien des Blocks), deterministisch sortiert — als Tests.
- [x] `--print-config`, Benutzerhandbuch und CHANGELOG `[Unreleased]`
      nachgezogen.
- [x] Unabhängiger Review, Report unter [`docs/reviews/`](../../../reviews/README.md).
- [x] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [x] Beobachtungs-Register fortgeschrieben (oder „keine Beobachtung" notiert).
- [x] Jedes Risiko aus §6 trägt einen Ausgang.

`make gates` grün.

## 3. Plan (vor Code)

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| Regel-Auswertung `shapes` in `internal/hexagon/` | update | `exact`-Vergleich, `unused`-Zählung |
| Config-Adapter | update | `expect`, `unused`, Widerspruchs-Fälle |
| Tests (gleich / anders / zusätzlich / fehlend; unused an/aus) | neu | Happy/Boundary/Negative der neuen Regel-Anforderung |
| `docs/user/benutzerhandbuch.md`, `CHANGELOG.md`, `--print-config` | update | öffentlicher Vertrag berührt |

## 4. Trigger

**Start** (`next` → `in-progress`): slice-210 in `done/`, WIP-Limit frei.

**Rückführungen — vorab benannt:**

- `in-progress` → `next` (zu groß): entfällt voraussichtlich — zwei Modi auf
  vorhandenem Normalisierer; träte es ein, ginge `shape-unused` als eigener
  Slice.
- `in-progress` → `open` (blockiert): der Vertrag aus slice-209 lässt offen,
  wie eine `expect:`-Datei selbst normalisiert wird (gleicher Dialekt?) —
  zurück an den Planner.

## 5. Closure-Trigger

DoD vollständig; bewusstes Brechen je Testbehauptung gezeigt; `make gates`
Exit 0; Closure-Notiz mit Lerneintrag.

## 6. Risiken und offene Punkte

- **`exact` ist spröde:** jede Formatierungs-unabhängige, aber inhaltliche
  Änderung (Versions-Bump) macht die Sollform-Datei nötig — Adopter könnten
  `exact` meiden. — **Ausgang:** *entfallen* — gestrichen mit Begründung: die
  Sprödigkeit ist eine benannte Eigenschaft des Modus (Handbuch §4: „strenger, aber auch spröder“),
  und `allow-statements` mit Regex-Einträgen ist die dokumentierte Alternative für Versions-Literale;
  ob Adopter `exact` meiden, ist eine Wahl, kein Risiko des Vertrags.

## 7. Closure-Notiz

**Lerneintrag — Form: benannte Spec-Lücke.** „Zeigt nicht aus der Scan-Wurzel hinaus“ stand in
[SPEC-CONF-001](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema) nur
**lexikalisch** — kein `/`, kein `..` —, während der Lesezugriff dem Dateisystem folgt. Ein Symlink
im Pfad führte trotzdem hinaus, und fremder Inhalt erschien als Befund. Benannt und geschlossen in
Spezifikation 0.35.0: kein Bestandteil eines `expect`-Pfads und kein Bestandteil des literalen
Präfixes eines `files`-Globs darf ein Symlink sein; die Sollform-Datei darf nicht dieselbe Datei
sein wie eine geprüfte (lexikalisch normalisiert oder als Hardlink).

**Geliefert:** `mode: exact` mit `expect` und Befund `shape-differs` (erste Abweichung: anders,
zusätzlich, fehlend) sowie `unused: fail` mit Befund `shape-unused` an der Zeile des Eintrags in der
`.a-check.yml` ([AC-FA-RULE-012](../../../../spec/lastenheft.md#ac-fa-rule-012),
[ADR-0041](../../adr/0041-shapes-sollform-je-datei.md)); Gerüst, Handbuch 1.43, CHANGELOG. Damit ist
der Vertrag aus slice-209 vollständig implementiert.

**Was hat funktioniert:** Jeder Fix bekam eine Mutations-Gegenprobe, die aus dem richtigen Grund rot
wurde — so trennte der Lauf „Test grün“ von „Test belegt die Zusage“. Und der Reviewer bekam meine
Vermutungen als **Sonden** mit, nicht als Feststellungen: Genau eine davon (die `files`-Suche sei
sicher) war falsch und wurde so gefunden.

**Was ging anders als geplant:** Vier Review-Runden statt einer. Die Hermetik-Lücke zeigte sich an
drei Stellen nacheinander (F-3, N-1, G-1), weil jeder Fix nur die gemeldete Stelle schloss; die
letzte lag in Code aus slice-210 und kam per Plan-Änderung in diesen Slice. Spezifikation 0.35.0
entstand dadurch in drei Schritten.

**Steering-Loop-Eintrag:** benannte Spec-Lücke, siehe Lerneintrag — gezählt, nicht verkörpert.

**Beobachtungs-Register (`../observations/`):** neu
[`BEO-ADAPT/hermetik-lexikalisch-statt-am-dateisystem`](../observations/BEO-ADAPT/hermetik-lexikalisch-statt-am-dateisystem/observation.md)
(1×, drei Funde in einem Vorgang); neu mit zwei Belegen (slice-210, slice-211)
[`BEO-GATE/testbeschreibung-weiter-als-assertion`](../observations/BEO-GATE/testbeschreibung-weiter-als-assertion/observation.md)
und
[`BEO-SPEC/ausgabe-einzeilig-nicht-zugesichert`](../observations/BEO-SPEC/ausgabe-einzeilig-nicht-zugesichert/observation.md)
— beide Klassen hatte der Review von slice-211 als wiederkehrend aus slice-210 benannt.

**Folge-Slices:** keine.

**Risiken aus §6:** das eine Risiko trägt seinen Ausgang (*entfallen* mit Begründung).

**Drei Paarungen:** getragen von der Closure von welle-16.

**Trigger-Audit der aktiven MR:** [`MR-016`](../../../../harness/conventions.md#mr-016) [`MR-019`](../../../../harness/conventions.md#mr-019) [`MR-024`](../../../../harness/conventions.md#mr-024) [`MR-025`](../../../../harness/conventions.md#mr-025) [`MR-027`](../../../../harness/conventions.md#mr-027) [`MR-028`](../../../../harness/conventions.md#mr-028) [`MR-029`](../../../../harness/conventions.md#mr-029) [`MR-030`](../../../../harness/conventions.md#mr-030) — 0 offen (geprüft 2026-10-06; kein Release seit `v0.20.0` — der Auflösungs-Trigger des dritten Eintrags tritt mit dem bevorstehenden Release ein).

## 8. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** `KERN` (Achsen 1, 2, 3 ✓) ·
`ADAPT` (2, 3 ✓) · `USER` (2, 3 ✓).

**Vorgelagert — offene Beobachtungen sichten** (2026-10-06, gemergter Stand):
dieselben Treffer wie slice-210 §8 (`BEO-KERN/dirvocab-portfor-auseinander`
1×, `BEO-USER/handbuch-vokabel-der-adapter-rolle` 1×); beim Anlegen erneut
sichten, weil slice-210 sie bis dahin bewegt haben kann.

**Modus-Begründungsblock:** alle berührten Sub-Areas GF.
