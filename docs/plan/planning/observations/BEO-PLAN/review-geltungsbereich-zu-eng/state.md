**Stand:** offen (3×) — Schwelle erreicht, Ausgang noch nicht zugewiesen.

Die drei Instanzen teilen die Form, nicht den Gegenstand: ein Review-Geltungsbereich (slice-105),
eine Sensor-Messung (slice-172), eine Diff-Bereinigung (slice-174). Jedes Mal war die Begrenzung
**begründet** und trotzdem zu eng; jedes Mal fand es jemand anderes als der Messende — zweimal der
Maintainer, einmal der unabhängige Review.

**Was ein Ausgang leisten müsste:** Ein Sensor scheidet aus — „ist dieser Geltungsbereich weit
genug?" ist ein Urteil über eine Absicht, kein Match. Bliebe eine geschärfte Regel: die
Prüf-Frage, ob eine Messung **eine** Klasse ihres Gegenstands erfasst hat oder **alle**, an einem
Ort, den der Messende liest. Der Skill
[`cr-text-reviewer.md`](../../../../../../.harness/skills/cr-text-reviewer.md) stellt eine
verwandte Frage („hast du *das* gemessen, worüber du redest?") und wäre der nächstliegende
Träger — er ist heute aber auf CR-Texte an fremde Werkzeuge gescopt.

Zuzuweisen bei der Closure von [welle-15](../../../welle-15-regelwerk-v650-migration.md), deren
Lese-Schritt die 3×-Einträge aufnimmt.
