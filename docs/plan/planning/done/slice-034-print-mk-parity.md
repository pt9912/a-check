# slice-034 — Fragment-Parität: committete `a-check.mk` == `--print-mk`

**Status:** done (**abgeschlossen 2026-07-09** — Gate-Härtung umgesetzt, Negativ-Probe verifiziert, `make ci` grün; noch unveröffentlicht). Closure-Notiz + Lerneintrag: §7.
**Typ:** Gate-Härtung (Durchsetzungsschicht), Folge von slice-033.
**Bezug:** schärft die `image-test`-Akzeptanz zu
[AC-FA-DIST-001](../../../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk)
(das `a-check.mk`-Fragment „Erzeugt von `a-check --print-mk`"); netzlos/hermetisch
([AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)).
Kein neuer Vertrag, kein ADR (Gate-**Verschärfung**, keine Lockerung — [AGENTS §3.6](../../../../AGENTS.md#36-gates-dürfen-nicht-ohne-adr-gelockert-werden)).
[Roadmap](../in-progress/roadmap.md).

## 1. Motivation

Die committete [`a-check.mk`](../../../../a-check.mk) trägt den Header „Erzeugt von `a-check --print-mk`",
ist aber eine **eigene** Datei — kein generiertes Artefakt mit Gate. slice-033 änderte den Generator
(`mkFragment` in `internal/cli/cli.go`, neues `a-check-graph`-Target), ohne dass die committete Datei
nachgezogen wurde; sie driftete still (fehlendes Target **und** ein `v0.11.0`-Alt-Kommentar), erst bei
gezielter Nachfrage entdeckt. `gate-consistency` prüft nur die **Digest**-Gleichheit der harten Pins, **nicht**
die **Byte-Parität** Fragment ↔ committete Datei — genau diese Lücke schließt der Slice, damit die
ausgelieferte Referenz nie wieder still vom `--print-mk`-Vertrag abweicht.

## 2. Design

**Home: `tools/image-test.sh`, Block (1)** — nicht `gate-consistency.sh`. Begründung: `gate-consistency`
ist bewusst **image-frei** (reines host-bash, grep't Digests aus Dateien); die `--print-mk`-Ausgabe braucht
das gebaute Binary/Image. `image-test` **baut das Image und fährt `--print-mk` bereits** (Block 1, `mk.c.out`),
ist also der self-contained Ort. Ergänzung (eine Zeile):

```sh
cmp -s "$WORK/mk.c.out" a-check.mk || fail "committete a-check.mk ≠ --print-mk-Output (Fragment-Parität; regeneriere: a-check --print-mk > a-check.mk)"
```

Die `--print-mk`-Ausgabe pinnt ihren Digest aus der `aCheckImage`-Const (`cli.go`), **nicht** aus dem
Build-`VERSION` — daher ist `mk.c.out` unabhängig vom `image-test`-Build-`VERSION` byte-gleich zur committeten
`a-check.mk` (deren Digest der Re-Pin setzt). Nach einem künftigen `mkFragment`-Change **oder** Re-Pin bleibt
die Parität nur, wenn die committete Datei mitgezogen wird — sonst bricht `image-test` (fail-closed, genau der
slice-033-Fall).

## 3. Geplanter Umfang

1. **`tools/image-test.sh`:** Block (1) um den `cmp mk.c.out ↔ a-check.mk`-Assert erweitern.
2. **Doku:** die `image-test`-Zeile in [AGENTS §4](../../../../AGENTS.md#4-quality-gates) +
   [harness/README.md §Sensors](../../../../harness/README.md#sensors-feedback-gates) um den Paritäts-Aspekt ergänzen.
3. **Verifikation:** `make image-test` grün (Parität hält) **plus** Negativ-Probe (committete `a-check.mk`
   künstlich driften → `image-test` bricht → Gate beißt), dann zurücksetzen.
4. **Gates:** `make ci` + `make trace-check`.

Kein Lastenheft-/SPEC-/ADR-Change: Harness-Gate-Härtung (Präzedenz slice-018 Pin-Gate, slice-029
doc-check-Härtung) — die Parität ist bereits durch den `a-check.mk`-Header „Erzeugt von `--print-mk`"
impliziert, der Slice erzwingt sie nur maschinell.

## 4. Akzeptanzkriterien (als Gate)

- **Happy:** Given committete `a-check.mk` == `--print-mk`-Output, when `make image-test` läuft, then Block (1)
  grün (nativ == Container == committete Datei).
- **Negative:** Given eine gedriftete committete `a-check.mk` (fehlendes Target **oder** Wortlaut-Abweichung),
  when `make image-test` läuft, then **Fehler** mit Regenerier-Hinweis (fail-closed).
- **Boundary:** Given ein Re-Pin (neuer Digest in `cli.go` **und** committeter `a-check.mk` im selben Commit),
  when `make image-test` läuft, then Parität hält (der Digest kommt in beiden aus derselben Quelle).

## 5. Grenzen / Folge

- Prüft **Fragment-Parität**, nicht ob der Digest online auf ein reales Image auflöst (Registry-/Netz-
  Eigenschaft, [AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze);
  beim Re-Pin online verifiziert) — komplementär zum Digest-Gleichheits-Check in `gate-consistency`.
- Deckt nur `a-check.mk`; `--print-config`-Gerüst ↔ eine committete Beispiel-Config gibt es nicht (kein
  committetes Pendant).

## 6. Sub-Area-Modus-Begründung

### Sub-Area: Durchsetzungsschicht (Gate)

- **Modus:** GF — Gate führt, erzwungen über `make ci`.
- **Konventionen-Dichte:** hoch — `image-test`/`gate-consistency` sind etabliert, Gate-Härtung ist gelebte
  Praxis (slice-018/029).
- **Phase-Reife:** Phase 4 — `image-test` real und grün; dieser Slice hängt einen Assert an Block (1).
- **Evidenz-/Diskrepanz-Risiko:** niedrig — die Negativ-Probe belegt die Wirksamkeit direkt.
- **Reconciliation-Aufwand:** keiner erwartet.

## 7. Closure-Notiz (nach `done`)

**Abgeschlossen 2026-07-09** auf Branch `slice-034-print-mk-parity`.

**Geliefert:** `tools/image-test.sh` Block (1) um den `cmp` committete `a-check.mk` ↔ `--print-mk`-Output
erweitert (self-contained: `image-test` baut das Image + fährt `--print-mk` bereits); cwd-unabhängiger
Repo-Root-Anker (`ROOT`); AGENTS §4 + harness/README §Sensors `image-test`-Zeile um den Paritäts-Aspekt
ergänzt. Kein Lastenheft-/SPEC-/ADR-Change (Gate-Verschärfung, Präzedenz slice-018/029).

**Gate-Evidenz:** `make gates` grün (doc-check 94/0, gate-consistency ok); **Negativ-Probe** direkt
verifiziert: gedriftete committete `a-check.mk` → `make image-test` **FAIL** (Exit 2, „committete a-check.mk
!= --print-mk … regeneriere") → nach `git checkout` wieder grün. Das Gate hätte die slice-033-Drift gefangen.

### Lerneintrag

**Ein „generiertes" Artefakt ohne Byte-Gate ist ein un-eingelöstes Versprechen.** Die committete `a-check.mk`
trug den Header „Erzeugt von `--print-mk`", war aber nur **digest**-gegatet — slice-033 änderte den Generator,
die Referenz driftete still. Lehre: sobald ein committetes Artefakt als „aus X erzeugt" deklariert ist, gehört
die **Byte-Gleichheit** unter ein fail-closed-Gate, sonst ist die Deklaration Wunschdenken. **Placement-Lehre:**
ein Paritäts-Check gehört dorthin, wo die *Erzeuger-Ausgabe schon vorliegt* (`image-test` hat Image + `--print-mk`),
nicht ins bewusst image-freie Meta-Gate (`gate-consistency`) — den Ort nach der **Datenverfügbarkeit** wählen,
nicht nach thematischer Nähe (deshalb bewusst von der ursprünglichen „gate-consistency"-Formulierung abgewichen).

**Folge:** kein committetes `--print-config`-Pendant existiert, daher dort keine Parität nötig (§5).
