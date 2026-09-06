# `make trace-check` ist im lokalen Klon nicht befragbar

**Sub-Area:** Gate-/Werkzeug-Schicht

[`AGENTS.md`](../../../../../../AGENTS.md) §5 benennt `make trace-check` als die Durchsetzung der
Commit-Traceability — lokal über `HEAD~1..HEAD`, in der CI über den Commit-Range. Lokal bricht das
Target mit `d-check: error: Range-Basis-Vorfahren nicht lesbar: object not found` ab, für jeden
geprüften Range.

Die Objekte lösen für `git` selbst auf (`git cat-file -t HEAD~2` → `commit`), der Klon ist nicht
shallow. Auffällig ist die Objekt-Ablage: `.git/objects/pack/` enthält neben dem kanonischen
`pack-<sha>.pack` ein `loose-<sha>.pack`. Ein Leser, der nur `pack-*.idx` einliest, sieht dessen
Objekte nicht.

**Gefährlich ist nicht der Ausfall, sondern seine Form:** das Target meldet Exit 2, also *rot* —
nicht *grün ohne Gegenstand*. Es ist damit kein stiller Ausfall; wer es aufruft, sieht ihn. Was
fehlt, ist der Aufruf: `trace-check` hängt in keinem Aggregat, der `commit-msg`-Hook ist Opt-in
(`make hooks`), und in der CI läuft es gegen einen frisch geklonten Baum ohne dieses Pack.
