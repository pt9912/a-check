**Stand:** verkörpert in [`AGENTS.md`](../../../../../../AGENTS.md) §5 (Regel *Geltungsbereich einer
Messung*) `seit slice-179` — zugewiesen im Lese-Schritt der Closure von
[welle-15](../../../done/welle-15-regelwerk-v650-migration.md).

Drei Instanzen, dieselbe Form, verschiedene Gegenstände: ein Review-Geltungsbereich (slice-105),
eine Sensor-Messung (slice-172), eine Diff-Bereinigung (slice-174). Jedes Mal war die Begrenzung
**begründet** und trotzdem zu eng; jedes Mal fand es jemand anderes als der Messende — zweimal der
Maintainer, einmal der unabhängige Review.

**Warum kein Sensor:** *„Ist dieser Geltungsbereich weit genug?"* ist ein Urteil über eine Absicht,
kein Match. Und ein zweites Muster, das den Bestand nach übersehenen Klassen durchsucht, kann
dieselbe Verengung haben wie das erste.

**Was stattdessen greift** — die billigste Prüfung, die alle drei Fälle gefangen hätte: Wer eine
Messung als Beleg schreibt, nennt ihren **Geltungsbereich** und sagt, ob er den Gegenstand deckt.
Bei slice-174 wäre das gewesen: *„bereinigt um Tabellen-Trennzeilen"* — und die Frage, ob das die
einzige Formatierungs-Klasse ist, hätte sich beim Schreiben gestellt.
