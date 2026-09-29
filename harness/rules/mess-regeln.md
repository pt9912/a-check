# Drei Mess-Regeln binden jeden, der einen Beleg schreibt

*Ausgelagert aus [`AGENTS.md`](../../AGENTS.md) §5, Regel 15 — seit slice-201. Der Wortlaut ist unverändert; die Links sind der neuen Tiefe nachgeführt.*

**Drei Mess-Regeln binden jeden, der einen Beleg schreibt** — also auch den
  Implementer- und den Planner-Lauf, nicht nur den Review:
  1. *Geltungsbereich einer Messung* (`seit slice-179`): Wer eine Messung als
     Beleg schreibt — Slice-Plan, Closure-Notiz, Review-Report —, **nennt ihren
     Geltungsbereich** und sagt, ob er den Gegenstand deckt. Nicht *„22 Befunde,
     keine weitere Klasse"*, sondern *„22 Befunde über Markdown-Links; Prosa
     sieht das Instrument nicht"*.
  2. *Eine Mutations-Probe belegt erst, wenn sie rot war* (`seit slice-181`):
     Wer einen Prüfer mit einer Probe belegt, zeigt **beide** Richtungen und
     nennt die **Meldung** der roten, nicht nur den Exit-Code. Grün beweist
     nichts — ein Prüfer, der seinen Gegenstand nicht erreicht, ist grün.

  3. *Wer eine Menge zählt, zählt sie zweimal verschieden* (`seit slice-193`):
     Eine Zählung, die einen Befund oder einen Umfang trägt, wird mit einem
     **zweiten, anders gebauten** Zähler wiederholt. Weichen beide ab, ist der
     Unterschied der Befund — nicht die erste Zahl.

  **Kein Sensor:** alle drei sind ein Urteil über eine Absicht, einen Aufbau
  oder eine Zuordnung (§3.7). Die **Herleitung** und die gemessenen Fälle stehen im Reviewer-Skill
  ([`.harness/skills/reviewer.md`](../../.harness/skills/reviewer.md) §Mess-Regeln) —
  dort urteilt, wer prüft; hier steht der Satz, an den sich bindet, wer
  schreibt.
