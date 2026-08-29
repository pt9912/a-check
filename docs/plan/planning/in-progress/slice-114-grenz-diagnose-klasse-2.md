# slice-114 — R-2: die Grenz-Diagnose meldet auflösende Zeilen

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.
**Autor:** Claude (Opus 5), im Auftrag des Maintainers. **Datum:** 2026-08-29.
**Berührte Spec-Stellen:** [SPEC-CLI-001](../../../../spec/spezifikation.md#spec-cli-001--aufruf-scan-wurzel-und-exit-codes)
— die stderr-Ausgabe eines Scans.
**Deckt:** `R-2` aus [`docs/reviews/2026-08-15-v0170-go-kern.md`](../../../reviews/2026-08-15-v0170-go-kern.md).
**Bezug:** Maintainer-Auftrag 2026-08-15, wörtlich: *„R-2: die Grenz-Diagnose meldet auflösende
Zeilen fälschlich, dafür braucht es eine Folge-ADR zu [`ADR-0031`](../../adr/0031-heuristik-grenzen-diagnose.md)."*
[Roadmap](../in-progress/roadmap.md).

---

## 1. Ziel

Dieselbe Zeile erscheint gleichzeitig als **Befund** und als **unbeurteilt**:

```text
core/model.cpp:1: core-impurity: Kern importiert ../adapters/db/x.h
Hinweis: … core/model.cpp:1: relativer Pfad, den der Auflösungs-Modus "path" nicht auflöst
```

Die Ursache steht **wörtlich in der ADR**, nicht nur im Code.
[`[`ADR-0031`](../../adr/0031-heuristik-grenzen-diagnose.md)`](../../adr/0031-heuristik-grenzen-diagnose.md) Entscheidung 5 behauptet: *„ein
`../`-Pfad gegen einen nicht-`relative`-Modus kann kein Ziel treffen, **egal wie der Baum
aussieht**."* Der Code wiederholt es als Kommentar.

**Die Prämisse ist falsch, und zwar aus der Auflösung selbst ablesbar:** unter Modus `path` gibt
`resolveImport` das Symbol **wörtlich** zurück, Punkte inklusive; `layerOfCand` sucht das
Glob-Präfix mit `segIndex` **segmentweise an beliebiger Stelle**. Für `../adapters/db/x.h` und
`layers: {adapters: ["adapters/**"]}` trifft es hinter den Punkten. Unter `fixed-root` ebenso, dort
mit vorangestellter Wurzel. Relative Includes sind in C++ die Norm — der Fall trifft breit.

Weil die falsche Aussage **dokumentiert** ist und [`ADR-0031`](../../adr/0031-heuristik-grenzen-diagnose.md) `Accepted` und damit immutabel, ist
die Korrektur eine **Folge-ADR mit `Supersedes`**, nicht ein Code-Fix.

## 2. Definition of Done

- [ ] Folge-ADR ist `Accepted`, löst [`ADR-0031`](../../adr/0031-heuristik-grenzen-diagnose.md) im Feld `Supersedes` ab, ist im Index verlinkt und ersetzt
      Entscheidung 5 durch eine Fassung ohne die falsche Prämisse.
- [ ] [SPEC-CLI-001](../../../../spec/spezifikation.md#spec-cli-001--aufruf-scan-wurzel-und-exit-codes)
      beschreibt die Klasse so, wie sie dann gilt.
- [ ] `HeuristicLimits` meldet Klasse 2 nur noch, wenn **kein** Layer-Glob-Präfix im Symbol
      segmentweise vorkommt; Probe: die Zeile aus §1 erscheint als Befund und **nicht** mehr als
      Hinweis, ein echt unauflösbares Symbol weiterhin als Hinweis.

- [ ] `make gates` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.
- [ ] Closure-Notiz mit benanntem Lerneintrag geschrieben (§7).
- [ ] Beobachtungs-Register fortgeschrieben.
- [ ] Jedes Risiko aus §6 trägt genau einen Ausgang.

## 3. Plan (vor Code)

| Datei | Änderungs-Art | Begründung |
|---|---|---|
| `docs/plan/adr/0035-*.md` | neu | Folge-ADR, löst [`ADR-0031`](../../adr/0031-heuristik-grenzen-diagnose.md) ab |
| `docs/plan/adr/README.md` | update | Index-Pflicht |
| `spec/spezifikation.md` | update | [`SPEC-CLI-001`](../../../../spec/spezifikation.md#spec-cli-001--aufruf-scan-wurzel-und-exit-codes), Klasse 2 |
| `internal/hexagon/core/rules.go` | update | `HeuristicLimits` |
| `internal/cli/cli_test.go` | update | Probe beide Richtungen |

**Auszuführende Gates:** `make gates` — tragend `test`, `coverage-gate`, `arch-check` und
`doc-immutable` (die ADR-Immutabilität wird hier bewusst über den zulässigen Weg berührt). Zum
Abschluss `make verify`.

## 4. Trigger

**Start:** unmittelbar — der Auftrag steht seit dem 2026-08-15, der Befund ist verifiziert.

**Rückführungen:** `in-progress` → `next`, falls die Neufassung von Entscheidung 5 mehr als eine
Klasse berührt — dann ist es kein Fix, sondern ein Diagnose-Umbau.

## 5. Closure-Trigger

ADR `Accepted` und indiziert, Spec nachgezogen, Code geändert, beide Proben belegt, Gates grün.

**Was bewusst nicht getan wird:** Klasse 1 (nicht extrahierte Zeilen) bleibt unberührt — sie
braucht keine Konfiguration und ist von diesem Befund nicht betroffen. Und die Diagnose bleibt
**tree-frei**: die neue Prüfung liest Modus **und Globs**, beides Konfiguration, nie den
Datei-Index. Sonst fiele sie in die Klasse, die [`ADR-0031`](../../adr/0031-heuristik-grenzen-diagnose.md) ausdrücklich ausgeschlossen hat und die
[`ADR-0029`](../../adr/0029-abdeckungs-diagnose-advisory.md) schon auf der Ziel-Seite ausgeschlossen hatte.

## 6. Risiken und offene Punkte

- *Die Neufassung könnte Klasse 2 auf null Fälle schrumpfen* — dann wäre sie eine Regel ohne
  Gegenstand. **Ausgang:** weiter offen, `BEO-022`; die Probe in §2 misst es, und ein Fall mit
  echtem Nicht-Treffer bleibt als Fixture stehen.
- *[`ADR-0031`](../../adr/0031-heuristik-grenzen-diagnose.md) bleibt zitiert, obwohl abgelöst* — **Ausgang:** gestrichen mit Begründung: das ist
  der Normalfall der Append-only-Disziplin; die abgelöste ADR bleibt lesbar, der Index führt
  beide, und die Nachfolgerin nennt sie in `Supersedes`.
- *Die Folge-ADR ist Voraussetzung für
  [slice-113](../open/slice-113-steering-loop-ins-register.md)* — **Ausgang:** Folge-Slice; dessen
  §0 nennt genau diese Bedingung.

## 7. Closure-Notiz

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice,
[`AGENTS.md`](../../../../AGENTS.md) §5; `make verify` prüft das.)_

## 8. Sub-Area-Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** berührt werden **Entscheidungen** (`docs/plan/adr/`),
**Spec-Straten** und **Kern und Regeln** (`internal/hexagon/`) — drei Sub-Areas, alle in der
Modus-Deklaration geführt. Das ist eine Schicht mehr als die Größen-Regel erlaubt; sie hängen hier
aber an **einer** Aussage und lassen sich nicht schneiden, ohne die Spec-first-Reihenfolge zu
brechen.

**Vorgelagert — offene Beobachtungen sichten:** keine Treffer für diese Sub-Areas.

Alle berührten Sub-Areas GF.
