# Review-Report: Etappe F (1/3) — slice-057 — 2026-07-26

**Review-Art:** Code-Review — geprüft gegen den Slice-Plan, `AGENTS.md` §3.1/§4/§5, Regelwerk
`grundlagen-klassifikation` §Steering Loop, `grundlagen-durchsetzungsschicht` (Tool-Call-Gate) und
`modul-09` (zwei Quadranten).

**Unabhängigkeit — ausdrücklich:** **Selbst-Review**, kein unabhängiger Lauf (neues
Kontextfenster, dieselbe Modell-Familie wie die Autoren-Instanz).

**Gegenstand:** `a4632e7..cc7ee9b` (4 Commits) — 328 eingefügte Zeilen, darunter 90 Zeilen
Guard-Logik und der neue Steering-Loop-Kanal.

**Skill:** `.harness/skills/reviewer.md` @ `20ee992` (2026-07-25) · <!-- d-check:ignore (Adopter-spezifischer Skill-Pfad, existiert im Ziel-Repo ggf. nicht) -->
**Modell:** `claude-opus-5[1m]` · **Datum:** 2026-07-26

**Eingangs-Kontext:**

- slice-057 (Gegenstand),
  [slice-048 §5](../plan/planning/done/slice-048-modul-delta-lesen.md) (Fund B-21),
  [slice-051 §4](../plan/planning/done/slice-051-workflow-und-freigabe.md) (offener Pipe-Sensor)
- `docs/plan/steering-loop.md`,
  [`.claude/hooks/pretooluse-command-guard.sh`](../../.claude/hooks/pretooluse-command-guard.sh)
- [`AGENTS.md`](../../AGENTS.md) §3.1 (Host-Minimum), §4 (Target-Tabelle), §5 (Steering-Loop-Regel)
- frühere Findings: [Etappe D](2026-07-26-etappe-d-slice-052-053-054.md) F-2 (drittes Muster ohne
  Eintrag), [Etappe E](2026-07-26-etappe-e-slice-050-051.md) F-2 (Ablauf-Ursache von SL-002)

---

## Findings

### F-1 — Die Gate-Liste des Guards deckt einen Teil der Prüf-Targets nicht ab

- `kategorie`: **MEDIUM**
- `quelle`: [`AGENTS.md`](../../AGENTS.md) §4 (Target-Tabelle);
  Steering-Loop SL-001 (Zweck der Regel)
- `pfad`: [`.claude/hooks/pretooluse-command-guard.sh:68-71`](../../.claude/hooks/pretooluse-command-guard.sh)
  (`const GATES = new Set([...])`)
- `befund`: Die Regel erkennt einen Gate-Lauf über eine **hartcodierte Liste** von 16
  Target-Namen. Nicht enthalten sind unter anderem `doc-immutable` — in `AGENTS.md` §4
  ausdrücklich als **CI-durchgesetzt** geführt (ADR-Immutabilität, §3.5) — sowie `doc-planning`,
  `doc-complete`, `doc-tracked`, `doc-targets` und `doc-trace`. Für diese Targets greift die Regel
  nicht: ihr Exit-Code kann unbemerkt in einer Pipe verschwinden, also genau der Vorgang, gegen den
  SL-001 antritt. Die Liste hat zudem keine Bindung an die Target-Tabelle oder das `Makefile` —
  `gate-consistency` gleicht Doku ↔ Makefile ab, aber nichts gleicht diese Liste damit ab.
- `verifizierbar`: **ja — live belegt.** `make doc-immutable | tail -1` wurde vom Guard **nicht**
  abgelehnt und lief durch. Der Lauf demonstriert den Schaden gleich mit: `make` brach mit
  `Fehler 2` ab (auf stderr, an der Pipe vorbei), während der Exit-Code der Pipeline der von
  `tail` war — der rote Lauf verschwand spurlos, exakt wie in SL-001 beschrieben.
- `gegenprobe`: Geprüft, ob die Auslassung bewusst ist — der Slice benennt in §3 und im
  Steering-Loop **eine** ehrliche Grenze (Sub-Shell-Strings), diese hier nicht. Ebenso geprüft, ob
  die betroffenen Targets bloß advisory sind: `doc-immutable` ist es nicht, es steht in
  `AGENTS.md` §4 als CI-durchgesetzt.

### F-2 — Der Wiedereinstiegs-Block nennt eine Commit-Zahl, die sein eigener Commit überholt

- `kategorie`: **LOW**
- `quelle`: Skill §MEDIUM/§LOW (Beleg-Genauigkeit)
- `pfad`: [`docs/plan/planning/in-progress/roadmap.md:112-115`](../plan/planning/in-progress/roadmap.md)
- `befund`: Der Block beschreibt die Kette als „**eine lineare Kette von 35 Commits**". Zum
  Zeitpunkt des Schreibens traf das zu (`main..c249aeb` = 35); mit dem Commit, der den Block
  einbringt, sind es **36**. Die Zahl ist damit ab ihrer Veröffentlichung falsch. Dies ist
  dieselbe Klasse wie die „Futur-Falle", die slice-050 im Bestand gefunden hat (slice-041, „noch
  unveröffentlicht", real längst ausgeliefert): eine Aussage über einen Stand, den das Dokument
  selbst verändert.
- `verifizierbar`: **ja — belegt.** `git rev-list --count main..cc7ee9b` → **36**;
  `main..35674e3` → 34.

### F-3 — Die Beleg-Liste von SL-002 stellt ungleiche Commits als gleichartig dar

- `kategorie`: **LOW**
- `quelle`: `docs/plan/steering-loop.md` §Pflege („ein Eintrag ohne
  Vorfallszahl ist unzulässig: die Zahl ist das Einzige, was die Schwelle prüfbar macht")
- `pfad`: `docs/plan/steering-loop.md:60-61` (SL-002, Vorfälle)
- `befund`: SL-002 belegt „sieben" Vorfälle mit sieben Commit-SHAs und beschreibt sie als „jedes
  Mal einzeln nachgezogen". Fünf davon sind tatsächlich Ein- bis Zwei-Zeilen-Reparaturen; **zwei**
  nicht: `f57289d` (**200** eingefügte Zeilen) und `d436da9` (**115**) tragen jeweils die Substanz
  eines Folge-Slice und reparieren den Verweis nur nebenbei. Die **Zählung** der Vorfälle bleibt
  richtig — der Verweis brach jedes Mal —, aber die Beleg-Liste verdeckt, dass in denselben
  Commits ein zweites Muster steckt (Betreff ≠ Substanz, siehe
  [Etappe D](2026-07-26-etappe-d-slice-052-053-054.md) F-2). Der Autor hatte die Commits beim
  Schreiben des Eintrags vor sich.
- `verifizierbar`: **ja — belegt.** `git show --stat` über die sieben genannten SHAs.

### F-4 — Der Guard setzt `node` voraus; `AGENTS.md` §3.1 nennt es nicht im Host-Minimum

- `kategorie`: **INFO** (Rollen-Verweis: gehört zu Etappe F 2/3 oder einem eigenen Slice)
- `quelle`: [`AGENTS.md`](../../AGENTS.md) §3.1 („Der Host braucht nur `git`, GNU `make`, `bash`
  und Docker.")
- `pfad`: [`.claude/hooks/pretooluse-command-guard.sh:24-31`](../../.claude/hooks/pretooluse-command-guard.sh)
- `befund`: Der Guard führt seine Prüflogik über `node -e` aus und blockiert **fail-closed**, wenn
  `node` fehlt. Auf einem Host mit exakt der in §3.1 dokumentierten Mindestausstattung wäre damit
  jeder Bash-Tool-Call blockiert. **Nicht diesem Diff anzulasten**: die Abhängigkeit stammt aus
  `6ed6397` (slice-005); dieser Slice weitet die node-abhängige Logik allerdings um 90 Zeilen aus.
  Als INFO gemeldet, weil die Auflösung (Deklaration ergänzen oder Interpreter wechseln) außerhalb
  des Review-Gegenstands liegt.
- `verifizierbar`: ja — `git show a4632e7:.claude/hooks/pretooluse-command-guard.sh` enthält
  `node` bereits fünfmal.

## Negativbefunde

- geprüft, ohne Befund: **Guard-Regel 2 greift real** — **live in dieser Session provoziert**,
  nicht aus dem Slice übernommen. `make guard-selftest | tail -1` → **abgelehnt**;
  `make guard-selftest && git commit …` → **abgelehnt**; beide mit der erklärenden Meldung samt
  richtiger Alternative. Der Sensor wirkt in der echten Umgebung, nicht nur im Selbsttest.
- geprüft, ohne Befund: **kein Fehlalarm auf die vorgeschriebene Form** —
  `make guard-selftest > datei 2>&1; echo "EXIT=$?"` lief durch, **EXIT=0**. Die Regel blockiert
  den Weg, den sie selbst vorschreibt, nicht.
- geprüft, ohne Befund: **die dokumentierte Grenze ist exakt so groß wie behauptet** —
  `bash -c "make guard-selftest | tail -1"` entkommt tatsächlich. Der Slice beschreibt genau das
  („quote-bewusst … Stolperdraht, keine Sandbox"). Die Grenze ist damit **nicht schlimmer als
  dokumentiert** — der wichtigste Test an einer selbst erklärten Einschränkung.
- geprüft, ohne Befund: **Vorfallszahl-Konsistenz SL-001** — slice-051 nennt „vier Vorfälle",
  SL-001 nennt „fünf". Kein Widerspruch: SL-001 schlüsselt auf („vier beim Doku-Schreiben, einer
  beim Roadmap-Nachzug"), und der fünfte fiel nach slice-051 an. Genau dieser fünfte trägt den
  Lerneintrag des Slice.
- geprüft, ohne Befund: **Merge-Spitze grün** — `make verify` auf `cc7ee9b`: **EXIT=0**
  (`54 Slice(s) in done/`, `6 Slice(s) ab slice-52 geprueft, 51 aelter`, `19 grandfathered`);
  `make doc-check`: EXIT=0.
- geprüft, ohne Befund: **Wiedereinstiegs-Block, Ketten-Aussagen** — `git merge-base --is-ancestor`
  bestätigt für alle geprüften Wegmarken (048, 049, 050, 052, 055), dass die Spitze sie enthält;
  `slice-031-deklarations-index-split-package` ist tatsächlich in `main`. Beide Aussagen tragen.
- geprüft, ohne Befund: **Selbstanwendung ehrlich dokumentiert** — der Slice hält fest, dass die
  Regel während ihrer Entwicklung zweimal ihren eigenen Autor blockierte und ein Syntax-Fehler
  kurzzeitig **jeden** Bash-Aufruf lahmlegte, statt das zu glätten. Für einen fail-closed Wächter
  ist das korrektes Verhalten, und der Fehlalarm ist als Ursache der Quote-Behandlung benannt —
  ein Sensor, der rauscht, wird abgeschaltet statt repariert.
- geprüft, ohne Befund: **Kanal-Ort begründet** — die Baseline verortet Steering-Loop-Einträge in
  `done/welle-NN-results.md`; da a-check keine Welle auditierbar schließt (B-13, offen), wäre das
  ein Kanal, der auf ein nicht existierendes Artefakt wartet. Der Zwischenschritt ist als solcher
  deklariert, samt Wanderungs-Bedingung.
- geprüft, ohne Befund: **Pflege-Regeln des Kanals** — Eintrag ab dem **zweiten** Vorfall (nicht
  erst dem dritten, sonst fehlt beim Schwellenwert die Zählung), Eintrag ohne Antwort zulässig,
  Eintrag ohne Vorfallszahl nicht, nichts wird gelöscht. Die Regeln adressieren die drei Wege, auf
  denen ein solches Register typischerweise verfällt.
- geprüft, ohne Befund: **SL-002 bleibt ehrlich offen** — „Antwort: **offen**", zwei benannte
  Kandidaten, keiner in diesem Slice gebaut, mit Begründung warum („die Antwort auf ein Muster
  gehört nicht in denselben Slice wie seine Erfassung, wenn sie ein eigenes Werkzeug braucht").
- geprüft, ohne Befund: **Proben-Arithmetik** — „13 Diskriminierungs-Proben (fünf müssen greifen,
  sechs dürfen nicht, zwei für die unveränderte Regel 1)": 5 + 6 + 2 = 13, konsistent.
- geprüft, ohne Befund: **Lerneintrag-Form** — „neuer Sensor", eine der drei benannten Formen, mit
  Ursache und empirischem Beleg (der fünfte Vorfall geschah nach dem Guide, von derselben Instanz,
  die ihn verfasst hatte).

## Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 0 |
| MEDIUM | 1 |
| LOW | 2 |
| INFO | 1 |

## Verdikt

**Merge-blockierend: nein.**

Die zweite Etappe ohne HIGH — und die einzige, deren zentrale Lieferung ich **im laufenden Betrieb
gegen mich selbst** testen konnte: der Guard blockierte in dieser Review-Session zwei echte
Kommandos, ließ die vorgeschriebene Form durch und verhielt sich an seiner erklärten Grenze exakt
wie dokumentiert. Das ist der stärkste Beleg, den ein Sensor-Slice liefern kann. Der Lerneintrag
(„Ein Guide, der nach dem vierten Vorfall geschrieben wird, verhindert den fünften nicht") ist der
inhaltlich beste der ganzen Kette, weil er am eigenen Rückfall belegt ist statt an einer
Regel-Zitation.

F-1 ist die eine Auflage: die Regel deckt ihren Gegenstand nur teilweise ab, und anders als bei
der Sub-Shell-Grenze ist diese Lücke **nicht** benannt. Sie ist billig zu schließen und sollte es
werden, bevor jemand aus dem grünen Guard schließt, jeder Gate-Lauf sei geschützt — `doc-immutable`
ist es nicht, wie der Live-Lauf zeigt.

Der Kanal selbst erfüllt seinen Zweck bereits: SL-001 und SL-002 sind belegt gezählt. Was fehlt,
ist der **dritte** Eintrag — das in den Etappen D und E belegte Muster „Commit-Betreff und
Traceability-ID bezeichnen nicht die enthaltene Arbeit", mit drei Vorfällen ebenfalls über der
Schwelle. F-3 zeigt, wie nah der Eintrag lag: zwei der Commits, die SL-002 bereits als Beleg
führt, sind genau die Ausreißer.

**Übergabe:** Findings gehen an die Implementation. Der Report ersetzt keine Verifikation —
DoD-/Spec-Konformität prüft `make verify` separat (Modul 11).
