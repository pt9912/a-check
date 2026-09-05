# Als vollständig behauptete Prüfung war unvollständig, extern statt intern gefangen

**Sub-Area:** Harness-Einstieg

Ein Slice erklärt eine eigene Prüfung für vollständig ("gemessen, nicht angenommen") — und die
Prüfung war es nicht. Der Fehler wird nicht durch einen internen Blick derselben Session
gefangen, sondern erst durch den Maintainer, der die Lücke von außen benennt. Drei belegte
Instanzen: [slice-142](../../../done/wellenlos/slice-142-claude-rules-symlinks-repariert.md)
(vier `.claude/rules/`-Symlinks von einer als vollständig behaupteten `grep`-Bereinigung
übersehen — Symlinks folgen keinem `grep`); [slice-135](../../../done/wellenlos/slice-135-regelwerk-v600-delta-analyse.md)/[slice-139](../../../done/wellenlos/slice-139-beobachtungsregister-migration.md)→[slice-143](../../../done/wellenlos/slice-143-archivierung-delta-analyse.md)
(Zeitdokumente-Archivierung fälschlich für optional erklärt — nur die Wellen-Hälfte der
Baseline-Regel gelesen); [slice-150](../../../done/wellenlos/slice-150-roadmap-form-nachgezogen.md)→[slice-151](../../../done/wellenlos/slice-151-offene-slices-ohne-welle-entfernt.md)
("Offene Slices ohne Welle"-Tabelle nach eigener Prüfung behalten, Maintainer korrigiert). Die
Session hat die Parallele zwischen den ersten beiden Fällen selbst benannt
([slice-143](../../../done/wellenlos/slice-143-archivierung-delta-analyse.md) §1 zieht sie
explizit), aber nie einen Register-Eintrag angelegt — dieselbe Lücke, die auch das Fehlen der
unabhängigen Reviewer-Rolle die ganze Session über offen ließ (kein getrennter Kontext, der den
blinden Fleck der schreibenden Session hätte auffangen können).
