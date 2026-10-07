# Fünf Mess-Regeln binden jeden, der einen Beleg oder eine Zusage schreibt

*Ausgelagert aus [`AGENTS.md`](../../AGENTS.md) §5, Regel 15 — seit slice-201; die Regeln 1–3 tragen den Wortlaut von dort, die Links sind der neuen Tiefe nachgeführt. Regel 4 steht seit slice-215, Regel 5 seit slice-219.*

**Fünf Mess-Regeln binden jeden, der einen Beleg oder eine Zusage schreibt** — also auch den
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

  4. *Ein Testkommentar sagt nicht mehr zu, als seine Assertion prüft*
     (`seit slice-215`): Für jede Eigenschaft, die ein Testkommentar nennt,
     steht eine Assertion, die sie prüft — eine Mutation, die genau diese
     Eigenschaft bricht, macht den Test rot. Ein Allquantor („alle Fälle",
     „jede Form") gilt nur über eine Menge, die der Test aufzählt; sonst sagt
     der Kommentar „die aufgeführten".

  5. *Eine Zusage nennt den Lauf, der sie hält* (`seit slice-219`): Wer in Doku,
     Kommentar, Vertrag oder Gate-Index zusagt, dass etwas geprüft, verglichen oder
     durchgesetzt wird **oder dass eine Eigenschaft gilt**, nennt den Lauf, der es
     hält, und dessen Grenze **am selben Satz** — als Text oder als Zeiger auf die
     Stelle, die sie trägt: nicht *„wird geprüft"*, sondern *„geprüft von X; Y nicht,
     weil …"*. Gibt es keinen Lauf, sagt der Satz das — *„geplant mit …"* oder
     *„vom Review getragen, kein Sensor"* —, oder er entfällt. In den Spec-Straten
     ist die Zusage der Vertrag selbst; ihren Lauf nennt der umsetzende Slice, und
     die Regel verlangt dort: keine Eigenschaft in den Vertrag, für die kein Lauf
     geplant ist.

  **Kein Sensor:** alle fünf sind ein Urteil über eine Absicht, einen Aufbau
  oder eine Zuordnung (§3.7). Die **Herleitung** und die gemessenen Fälle stehen im Reviewer-Skill
  ([`.harness/skills/reviewer.md`](../../.harness/skills/reviewer.md) §Mess-Regeln) —
  dort urteilt, wer prüft; hier steht der Satz, an den sich bindet, wer
  schreibt.
