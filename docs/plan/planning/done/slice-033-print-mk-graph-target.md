# slice-033 — `a-check.mk` liefert ein `a-check-graph`-Target

**Status:** done (**abgeschlossen 2026-07-09** — spec-first CR umgesetzt, `make ci` grün, `make a-check-graph` funktional verifiziert; noch unveröffentlicht). Closure-Notiz + Lerneintrag: §8.
**Typ:** CR (Distribution/Convenience), Folge von slice-032.
**Bezug:** erweitert
[AC-FA-DIST-001](../../../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk)
(Distribution: Image/`--print-mk`/`a-check.mk`) um ein Convenience-Target für den mit slice-032
gelieferten no-scan-Modus
[AC-FA-CLI-002](../../../../spec/lastenheft.md#ac-fa-cli-002--architektur-graph-ausgabe)
(`--print-graph`, [SPEC-CLI-002](../../../../spec/spezifikation.md#spec-cli-002--graph-renderer-vertrag)).
Präzisiert wird [SPEC-DIST-001](../../../../spec/spezifikation.md#spec-dist-001--laufzeitform-und-distribution).
Kernverträge: [AC-QA-01](../../../../spec/lastenheft.md#ac-qa-01--determinismus) (deterministische Ausgabe),
[AC-QA-03](../../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit) (digest-gepinntes Image im Fragment).
[Roadmap](../in-progress/roadmap.md).

## 1. Motivation

`--print-graph` ([AC-FA-CLI-002](../../../../spec/lastenheft.md#ac-fa-cli-002--architektur-graph-ausgabe))
ist heute nur als direkter `docker run … --print-graph /src`-Einzeiler nutzbar (Benutzerhandbuch §3.6).
Konsumenten binden `a-check` ohnehin über `include a-check.mk` ein (das `--print-mk` erzeugt) und fahren
`make a-check` als Gate. Ein paralleles **`a-check-graph`-Target** im selben Fragment macht die
Graph-Ausgabe **entdeckbar** und tippfrei (`make a-check-graph > architektur.mmd`), ohne die
Image-/Digest-/Mount-Details zu wiederholen. Es ist **kein** Gate: das Target schreibt Mermaid nach
stdout (read-only, kein Scan, Exit 0), analog zum bestehenden `a-check`-Scan-Target.

## 2. Design

Der `--print-mk`-Output (`internal/cli/cli.go`, `mkFragment`) erhält **zusätzlich** zum bestehenden
`a-check`-Scan-Target ein `a-check-graph`-Target — dieselbe `A_CHECK_IMAGE`-Variable, derselbe
netzlose read-only-Mount:

```makefile
.PHONY: a-check-graph
a-check-graph: ## Architektur-Graph (Mermaid) aus .a-check.yml auf stdout (read-only, kein Scan).
	docker run --rm --network none -v "$(CURDIR)":/src:ro $(A_CHECK_IMAGE) --print-graph /src
```

`A_CHECK_IMAGE` bleibt die eine digest-gepinnte Quelle
([AC-QA-03](../../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit)); das Target dupliziert keinen
Digest. Ausgabe geht nach stdout (Umleiten in eine Datei ist Sache des Nutzers) — konsistent mit
`--print-graph` und dem `a-check`-Scan-Target.

## 3. Geplanter Umfang

1. **Lastenheft:** [AC-FA-DIST-001](../../../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk)
   erweitern (das `a-check.mk`-Fragment liefert zusätzlich ein `a-check-graph`-Target) + neue AK;
   Version-Bump **0.19.0 → 0.20.0** + Historie. **AC-Änderung nur im Lastenheft.**
2. **Spezifikation:** [SPEC-DIST-001](../../../../spec/spezifikation.md#spec-dist-001--laufzeitform-und-distribution)
   um das `a-check-graph`-Target präzisieren; Version-Bump + Historie.
3. **Code:** `mkFragment` in `internal/cli/cli.go` um das Target ergänzen.
4. **Tests:** `TestPrintMk` (Target im Fragment) + `image-test.sh` Block (1) um eine
   `a-check-graph:`-Assertion am `--print-mk`-Output erweitern.
5. **Benutzerhandbuch:** §3.6 um die `make a-check-graph`-Variante ergänzen; Currency-Bump.
6. **Gates:** `make gates` + `make ci` + `make trace-check`.

Kein Architektur-Update: der Fragment-**Inhalt** ist ein Distributions-Detail von
[ARC-006](../../../../spec/architecture.md) (Composition Root, bedient `--print-mk`) — keine neue
Komponente/Port, kein Schichtwechsel.

## 4. Entscheide

- **Kein ADR.** Ein mechanisches Convenience-Target im `--print-mk`-Fragment hat keine echten
  Wahl-Punkte (kein neues Format, keine Semantik) — Format/Umfang der Graph-Ausgabe selbst regelt
  bereits [ADR-0024](../../adr/0024-print-graph-mermaid.md), die Distributionsform
  [ADR-0004](../../adr/0004-distribution-image-mk.md). Ein weiterer ADR wäre Zeremonie ohne Entscheidung.
- **Target-Name `a-check-graph`** (nicht `graph`): namespaced wie `a-check`, kollidiert nicht mit
  konsumenten-eigenen Targets.
- **stdout, kein Datei-Schreiben** — read-only bleibt Kernvertrag
  ([AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)); Umleiten
  ist Nutzer-Sache.

## 5. Akzeptanzkriterien (als Tests)

- **Happy:** Given das Image, when `a-check --print-mk` läuft, then enthält das Fragment neben dem
  `a-check`-Scan-Target ein `a-check-graph`-Target, das `--print-graph` mit digest-gepinntem
  `A_CHECK_IMAGE` und netzlosem read-only-Mount aufruft.
- **Boundary:** Given das erzeugte `a-check.mk` inklusive im Konsumenten-`Makefile`, when
  `make a-check-graph` läuft, then Mermaid-`flowchart` auf stdout, Exit 0, kein Schreibzugriff auf den
  Baum (read-only-Mount).
- **Negative:** Given `--print-mk` mit einem zusätzlichen unbekannten Flag, when aufgerufen, then Exit 2
  (unverändert; das neue Target ändert die Flag-Behandlung nicht).

## 6. Grenzen / Folge

- **Nur ein stdout-Target**, kein Auto-Schreiben in eine Datei und kein „öffne im Browser"-Convenience.
- Weitere Ausgabeformate (`--graph-format` DOT/Graphviz) bleiben eigener Folge-Slice (aus slice-032 §7).

## 7. Sub-Area-Modus-Begründung

### Sub-Area: Spec + Distribution + Doku

- **Modus:** GF — Doc führt, Code folgt über Tests/Gates.
- **Konventionen-Dichte:** hoch — Anforderungs-Anlege-Prozess, `--print-mk`-Muster und `image-test`
  sind etabliert.
- **Phase-Reife:** Phase 4 — `--print-mk`/`--print-graph` implementiert; dieser Slice fügt nur ein
  Fragment-Target hinzu.
- **Evidenz-/Diskrepanz-Risiko:** niedrig — die AK binden Fragment-Inhalt (`TestPrintMk`) und Lauf
  (`image-test`) an Tests.
- **Reconciliation-Aufwand:** keiner erwartet.

## 8. Closure-Notiz (nach `done`)

**Abgeschlossen 2026-07-09** auf Branch `slice-033-print-mk-graph-target`.

**Geliefert:** [AC-FA-DIST-001](../../../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk)
erweitert (Lastenheft **0.20.0**) + [SPEC-DIST-001](../../../../spec/spezifikation.md#spec-dist-001--laufzeitform-und-distribution)
(Spez **0.20.0**); `mkFragment` (`internal/cli/cli.go`) um das `a-check-graph`-Target; `TestPrintMk` +
`image-test.sh`-Block (1) um die Target-Assertion; Benutzerhandbuch §3.6 `make`-Variante (Currency **1.29**);
CHANGELOG `[Unreleased]`. **Kein ADR** (mechanisches Convenience-Target ohne Wahl-Punkte). **Kein
Architektur-Update** — der Fragment-Inhalt ist ein Distributions-Detail von
[ARC-006](../../../../spec/architecture.md) (Composition Root, bedient `--print-mk`).

**Gate-Evidenz:** `make ci` grün (coverage 97,10 %, arch-check 0, doc-check 93/0, `image-test` inkl.
`a-check-graph`-Target in Block (1)); `make trace-check` über die Range grün. **Funktionale Probe:** das
erzeugte `a-check.mk` als echtes Makefile eingebunden, `make a-check-graph A_CHECK_IMAGE=a-check:dev` erzeugt
den Mermaid-Graphen (Exit 0) — die Recipe-TAB-Einrückung stimmt.

### Lerneintrag

Ein `make`-Target für einen stdout-dumpenden Modus ist legitim, wenn es die Image-/Digest-/Mount-Details
**kapselt** — aber nur **additiv** zur `A_CHECK_IMAGE`-Einzelquelle, nie mit einem zweiten Digest. Die Grenze
Inspektion (`--print-graph`) ↔ Distribution (`a-check.mk`) ist eine **Verortungs**-Frage (welche Anforderung
trägt es), **kein Verbot**: das Target lebt in
[AC-FA-DIST-001](../../../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk),
weil es das gelieferte Fragment betrifft — der no-scan-Modus selbst bleibt
[AC-FA-CLI-002](../../../../spec/lastenheft.md#ac-fa-cli-002--architektur-graph-ausgabe). **Prozess-Lehre
(Anlass dieses CR):** eine Behauptung „bewusst nicht erweitert" braucht eine **zitierbare normative Stelle**;
fehlt sie, ist es „**nicht im Scope**", nicht „verboten" — die genaue Wortwahl entscheidet, ob ein Folge-CR
nötig ist oder nicht.

**Folge:** weitere Ausgabeformate (`--graph-format` DOT/Graphviz) bleiben eigener Folge-Slice.
