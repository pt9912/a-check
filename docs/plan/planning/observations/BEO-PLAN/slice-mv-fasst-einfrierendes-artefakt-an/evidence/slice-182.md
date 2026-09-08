**Vorgang:** slice-182

**Fund:** `make slice-mv SLICE=slice-182 TO=done` schrieb **acht** `pfad`-Felder im Review-Report
dieses Slice von `docs/plan/planning/in-progress/…` auf `…/done/…` um — dieselbe Ausfallart wie
bei slice-169, in einem anderen Vorgang und damit ein zweites Auftreten.

**Was das über die erste Behandlung sagt:** Bei slice-169 ist die Änderung von Hand
zurückgenommen worden, und der Eintrag blieb unter der Schwelle. Der Handgriff wiederholt sich
jetzt, weil am Werkzeug nichts geändert wurde — das ist kein Vorwurf, sondern die
Zähler-Mechanik: Ein Eintrag unter 3× wartet, und genau deshalb wird er beim nächsten Mal wieder
gebraucht.

**Wie es auffiel:** `git status` nach dem `slice-mv` — die Datei unter `docs/reviews/` steht
zwischen den Planning-Dateien. Nicht durch einen Lauf: `make doc-check` bleibt grün, weil der
Report seine Pfade in Inline-Code nennt (**null** Markdown-Links auf den Slice, nachgezählt).

**Behandlung:** wie beim ersten Mal mit `git checkout` zurückgenommen.
