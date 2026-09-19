# Einem Werkzeug wird eine Leistung zugeschrieben, die es nicht erbringt

**Sub-Area:** Gate-/Werkzeug-Schicht

Ein Commit-Betreff, eine Prosa-Zeile oder ein Kommentar nennt ein **Werkzeug** als Urheber einer
Wirkung, die der Lauf selbst erbracht hat. Das Werkzeug sagt in seinem eigenen Kopfkommentar, was
es **nicht** tut — die Zuschreibung widerspricht also einer Zusage, die im selben Repo steht.

**Warum das zählt:** Die Zuschreibung ist die einzige Auskunft, die ein späterer Leser über die
Herkunft einer Änderung hat. Nennt sie das falsche Werkzeug, sucht der Leser die Wirkung in dessen
Code, findet sie dort nicht — und schließt, die Wirkung sei undokumentiert. Die Arbeit des
**Laufs** verschwindet dabei hinter dem Namen eines Programms.

**Zwei Ausprägungen, beide am Lifecycle-Move beobachtet:** eine Prosa-Zeile, die dem Werkzeug die
*Reihenfolge* zweier Commits zuschreibt, obwohl es nur den Move fährt · ein Commit-Betreff mit
`(make <werkzeug>)`, dessen Commit daneben einen **Zustandssatz** ändert (den Ruhe-Marker der
Roadmap), den das Werkzeug nicht anfasst.

**Kein Sensor deckt das.** Ob eine Zuschreibung trägt, ist ein Urteil über ihren Gegenstand
([`AGENTS.md`](../../../../../../AGENTS.md) §3.7) — und ein Commit-Betreff ist nach dem Commit
nicht mehr korrigierbar, ohne `git` umzuschreiben. Was greift, ist die Frage beim Schreiben:
*Tut das Genannte das wirklich — oder war ich es?*

**Benannt, nicht gezählt** ist alles, was in einem Vorgang ohne abgeschlossenen Beleg auffiel;
die Zahl dieses Eintrags ist die seiner Belege, nicht die seiner Fundstellen.
